package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mainflux/mainflux/tools/mqtt-bench"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

// Keep struct names exported, otherwise Viper unmarshaling won't work
type mqttBrokerCfg struct {
	URL string `toml:"url" mapstructure:"url"`
}

type mqttMessageCfg struct {
	Size   int    `toml:"size" mapstructure:"size"`
	Format string `toml:"format" mapstructure:"format"`
	QoS    int    `toml:"qos" mapstructure:"qos"`
	Retain bool   `toml:"retain" mapstructure:"retain"`
}

type mqttTLSCfg struct {
	MTLS       bool   `toml:"mtls" mapstructure:"mtls"`
	SkipTLSVer bool   `toml:"skiptlsver" mapstructure:"skiptlsver"`
	CA         string `toml:"ca" mapstructure:"ca"`
}

type mqttCfg struct {
	Broker  mqttBrokerCfg  `toml:"broker" mapstructure:"broker"`
	Message mqttMessageCfg `toml:"message" mapstructure:"message"`
	TLS     mqttTLSCfg     `toml:"tls" mapstructure:"tls"`
}

type testCfg struct {
	Count int `toml:"count" mapstructure:"count"`
	Pubs  int `toml:"pubs" mapstructure:"pubs"`
	Subs  int `toml:"subs" mapstructure:"subs"`
}

type logCfg struct {
	Quiet bool `toml:"quiet" mapstructure:"quiet"`
}

type mainflux struct {
	ConnFile string `toml:"connections_file" mapstructure:"connections_file"`
}

type mfConn struct {
	ChannelID string `toml:"channelID" mapstructure:"channelID"`
	ThingID   string `toml:"thingID" mapstructure:"thingID"`
	ThingKey  string `toml:"thingKey" mapstructure:"thingKey"`
	MTLSCert  string `toml:"mtlsCert" mapstructure:"mtlsCert"`
	MTLSKey   string `toml:"mtlsKey" mapstructure:"mtlsKey"`
}

type config struct {
	MQTT mqttCfg  `toml:"mqtt" mapstructure:"mqtt"`
	Test testCfg  `toml:"test" mapstructure:"test"`
	Log  logCfg   `toml:"log" mapstructure:"log"`
	Mf   mainflux `toml:"mainflux" mapstructure:"mainflux"`
}

// JSONResults are used to export results as a JSON document
type JSONResults struct {
	Runs   []*bench.RunResults `json:"runs"`
	Totals *bench.TotalResults `json:"totals"`
}

var (
	cfg     config
	cfgFile string
)

var benchCmd = &cobra.Command{
	Use:   "mqtt-bench",
	Short: "mqtt-bench is MQTT benchmark tool for Mainflux",
	Long: `Tool for exctensive load and benchmarking of MQTT brokers used withing Mainflux platform.
        Complete documentation is available at https://mainflux.readthedocs.io`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
		}

		// Set config
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
			if err := viper.ReadInConfig(); err != nil {
				log.Printf("Failed to load config - %s", err.Error())
			}
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			log.Printf("Unable to decode into struct, %v", err)
		}

		fmt.Println("CFG: ", cfg)

		runBench()
	},
}

func main() {
	if err := benchCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}

func init() {
	// MQTT Broker
	benchCmd.PersistentFlags().StringVarP(&cfg.MQTT.Broker.URL, "broker", "b", "tcp://localhost:1883",
		"address for mqtt broker, for secure use tcps and 8883")

	// MQTT Message
	benchCmd.PersistentFlags().IntVarP(&cfg.MQTT.Message.Size, "size", "z", 100, "Size of message payload bytes")
	benchCmd.PersistentFlags().StringVarP(&cfg.MQTT.Message.Format, "format", "f", "text", "Output format: text|json")
	benchCmd.PersistentFlags().IntVarP(&cfg.MQTT.Message.QoS, "qos", "q", 0, "QoS for published messages, values 0 1 2")
	benchCmd.PersistentFlags().BoolVarP(&cfg.MQTT.Message.Retain, "retain", "r", false, "Retain mqtt messages")

	// MQTT TLS
	benchCmd.PersistentFlags().BoolVarP(&cfg.MQTT.TLS.MTLS, "mtls", "m", false, "Use mtls for connection")
	benchCmd.PersistentFlags().BoolVarP(&cfg.MQTT.TLS.SkipTLSVer, "skipTLSVer", "t", false, "Skip tls verification")
	benchCmd.PersistentFlags().StringVarP(&cfg.MQTT.TLS.CA, "ca", "", "ca.crt", "CA file")

	// Test params
	benchCmd.PersistentFlags().IntVarP(&cfg.Test.Count, "count", "n", 100, "Number of messages sent per publisher")
	benchCmd.PersistentFlags().IntVarP(&cfg.Test.Subs, "subs", "s", 10, "Number of subscribers")
	benchCmd.PersistentFlags().IntVarP(&cfg.Test.Pubs, "pubs", "p", 10, "Number of publishers")

	// Log params
	benchCmd.PersistentFlags().BoolVarP(&cfg.Log.Quiet, "quiet", "", false, "Supress messages")

	// Config file
	benchCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file")
	benchCmd.PersistentFlags().StringVarP(&cfg.Mf.ConnFile, "connections", "", "connections.toml", "config file for Mainflux connections")
}

func runBench() {
	var wg sync.WaitGroup
	var err error

	checkConnection(cfg.MQTT.Broker.URL, 1)
	subTimes := make(bench.SubTimes)
	var caByte []byte
	if cfg.MQTT.TLS.MTLS {
		caFile, err := os.Open(cfg.MQTT.TLS.CA)
		defer caFile.Close()
		if err != nil {
			fmt.Println(err)
		}

		caByte, _ = ioutil.ReadAll(caFile)
	}

	payload := string(make([]byte, cfg.MQTT.Message.Size))

	mfConns := []mfConn{}
	if _, err := toml.DecodeFile(cfg.Mf.ConnFile, &mfConns); err != nil {
		log.Fatalf("Cannot load Mainflux connections config %s \nuse tools/provision to create file", cfg.Mf.ConnFile)
	}

	resCh := make(chan *bench.RunResults)
	done := make(chan bool)

	start := time.Now()
	n := len(mfConns)
	var cert tls.Certificate
	for i := 0; i < cfg.Test.Subs; i++ {

		mfConn := mfConns[i%n]

		if cfg.MQTT.TLS.MTLS {
			cert, err = tls.X509KeyPair([]byte(mfConn.MTLSCert), []byte(mfConn.MTLSKey))
			if err != nil {
				log.Fatal(err)
			}
		}

		c := &bench.Client{
			ID:         strconv.Itoa(i),
			BrokerURL:  cfg.MQTT.Broker.URL,
			BrokerUser: mfConn.ThingID,
			BrokerPass: mfConn.ThingKey,
			MsgTopic:   fmt.Sprintf("channels/%s/messages/test", mfConn.ChannelID),
			MsgSize:    cfg.MQTT.Message.Size,
			MsgCount:   cfg.Test.Count,
			MsgQoS:     byte(cfg.MQTT.Message.QoS),
			Quiet:      cfg.Log.Quiet,
			Mtls:       cfg.MQTT.TLS.MTLS,
			SkipTLSVer: cfg.MQTT.TLS.SkipTLSVer,
			CA:         caByte,
			ClientCert: cert,
			Retain:     cfg.MQTT.Message.Retain,
			Message:    payload,
		}

		wg.Add(1)

		go c.RunSubscriber(&wg, &subTimes, &done)
	}

	wg.Wait()

	for i := 0; i < cfg.Test.Pubs; i++ {
		mfConn := mfConns[i%n]

		if cfg.MQTT.TLS.MTLS {
			cert, err = tls.X509KeyPair([]byte(mfConn.MTLSCert), []byte(mfConn.MTLSKey))
			if err != nil {
				log.Fatal(err)
			}
		}

		c := &bench.Client{
			ID:         strconv.Itoa(i),
			BrokerURL:  cfg.MQTT.Broker.URL,
			BrokerUser: mfConn.ThingID,
			BrokerPass: mfConn.ThingKey,
			MsgTopic:   fmt.Sprintf("channels/%s/messages/test", mfConn.ChannelID),
			MsgSize:    cfg.MQTT.Message.Size,
			MsgCount:   cfg.Test.Count,
			MsgQoS:     byte(cfg.MQTT.Message.QoS),
			Quiet:      cfg.Log.Quiet,
			Mtls:       cfg.MQTT.TLS.MTLS,
			SkipTLSVer: cfg.MQTT.TLS.SkipTLSVer,
			CA:         caByte,
			ClientCert: cert,
			Retain:     cfg.MQTT.Message.Retain,
			Message:    payload,
		}

		go c.RunPublisher(resCh)
	}

	// Collect the results
	var results []*bench.RunResults
	if cfg.Test.Pubs > 0 {
		results = make([]*bench.RunResults, cfg.Test.Pubs)
	}

	for i := 0; i < cfg.Test.Pubs; i++ {
		results[i] = <-resCh
	}

	totalTime := time.Now().Sub(start)
	totals := calculateTotalResults(results, totalTime, &subTimes)
	if totals == nil {
		return
	}

	// Print sats
	printResults(results, totals, cfg.MQTT.Message.Format, cfg.Log.Quiet)
}

func calculateTotalResults(results []*bench.RunResults, totalTime time.Duration, subTimes *bench.SubTimes) *bench.TotalResults {
	if results == nil || len(results) < 1 {
		return nil
	}
	totals := new(bench.TotalResults)
	subTimeRunResults := bench.RunResults{}
	msgTimeMeans := make([]float64, len(results))
	msgTimeMeansDelivered := make([]float64, len(results))
	msgsPerSecs := make([]float64, len(results))
	runTimes := make([]float64, len(results))
	bws := make([]float64, len(results))

	totals.TotalRunTime = totalTime.Seconds()

	totals.MsgTimeMin = results[0].MsgTimeMin
	for i, res := range results {

		if len(*subTimes) > 0 {
			times := mat.NewDense(1, len((*subTimes)[res.ID]), (*subTimes)[res.ID])

			subTimeRunResults.MsgTimeMin = mat.Min(times)
			subTimeRunResults.MsgTimeMax = mat.Max(times)
			subTimeRunResults.MsgTimeMean = stat.Mean((*subTimes)[res.ID], nil)
			subTimeRunResults.MsgTimeStd = stat.StdDev((*subTimes)[res.ID], nil)

		}
		res.MsgDelTimeMin = subTimeRunResults.MsgTimeMin
		res.MsgDelTimeMax = subTimeRunResults.MsgTimeMax
		res.MsgDelTimeMean = subTimeRunResults.MsgTimeMean
		res.MsgDelTimeStd = subTimeRunResults.MsgTimeStd

		totals.Successes += res.Successes
		totals.Failures += res.Failures
		totals.TotalMsgsPerSec += res.MsgsPerSec

		if res.MsgTimeMin < totals.MsgTimeMin {
			totals.MsgTimeMin = res.MsgTimeMin
		}

		if res.MsgTimeMax > totals.MsgTimeMax {
			totals.MsgTimeMax = res.MsgTimeMax
		}

		if subTimeRunResults.MsgTimeMin < totals.MsgDelTimeMin {
			totals.MsgDelTimeMin = subTimeRunResults.MsgTimeMin
		}

		if subTimeRunResults.MsgTimeMax > totals.MsgDelTimeMax {
			totals.MsgDelTimeMax = subTimeRunResults.MsgTimeMax
		}

		msgTimeMeansDelivered[i] = subTimeRunResults.MsgTimeMean
		msgTimeMeans[i] = res.MsgTimeMean
		msgsPerSecs[i] = res.MsgsPerSec
		runTimes[i] = res.RunTime
		bws[i] = res.MsgsPerSec
	}

	totals.Ratio = float64(totals.Successes) / float64(totals.Successes+totals.Failures)
	totals.AvgMsgsPerSec = stat.Mean(msgsPerSecs, nil)
	totals.AvgRunTime = stat.Mean(runTimes, nil)
	totals.MsgDelTimeMeanAvg = stat.Mean(msgTimeMeansDelivered, nil)
	totals.MsgDelTimeMeanStd = stat.StdDev(msgTimeMeansDelivered, nil)
	totals.MsgTimeMeanAvg = stat.Mean(msgTimeMeans, nil)
	totals.MsgTimeMeanStd = stat.StdDev(msgTimeMeans, nil)

	return totals
}

func printResults(results []*bench.RunResults, totals *bench.TotalResults, format string, quiet bool) {
	switch format {
	case "json":
		jr := JSONResults{
			Runs:   results,
			Totals: totals,
		}
		data, err := json.Marshal(jr)
		if err != nil {
			log.Printf("Failed to prepare results for printing - %s\n", err.Error())
		}
		var out bytes.Buffer
		json.Indent(&out, data, "", "\t")

		fmt.Println(string(out.Bytes()))
	default:
		if !quiet {
			for _, res := range results {
				fmt.Printf("======= CLIENT %s =======\n", res.ID)
				fmt.Printf("Ratio:               %.3f (%d/%d)\n", float64(res.Successes)/float64(res.Successes+res.Failures), res.Successes, res.Successes+res.Failures)
				fmt.Printf("Runtime (s):         %.3f\n", res.RunTime)
				fmt.Printf("Msg time min (us):   %.3f\n", res.MsgTimeMin)
				fmt.Printf("Msg time max (us):   %.3f\n", res.MsgTimeMax)
				fmt.Printf("Msg time mean (us):  %.3f\n", res.MsgTimeMean)
				fmt.Printf("Msg time std (us):   %.3f\n", res.MsgTimeStd)

				fmt.Printf("Bandwidth (msg/sec): %.3f\n\n", res.MsgsPerSec)
			}
		}
		fmt.Printf("========= TOTAL (%d) =========\n", len(results))
		fmt.Printf("Total Ratio:                 %.3f (%d/%d)\n", totals.Ratio, totals.Successes, totals.Successes+totals.Failures)
		fmt.Printf("Total Runtime (sec):         %.3f\n", totals.TotalRunTime)
		fmt.Printf("Average Runtime (sec):       %.3f\n", totals.AvgRunTime)
		fmt.Printf("Msg time min (us):           %.3f\n", totals.MsgTimeMin)
		fmt.Printf("Msg time max (us):           %.3f\n", totals.MsgTimeMax)
		fmt.Printf("Msg time mean mean (us):     %.3f\n", totals.MsgTimeMeanAvg)
		fmt.Printf("Msg time mean std (us):      %.3f\n", totals.MsgTimeMeanStd)

		fmt.Printf("Average Bandwidth (msg/sec): %.3f\n", totals.AvgMsgsPerSec)
		fmt.Printf("Total Bandwidth (msg/sec):   %.3f\n", totals.TotalMsgsPerSec)
	}
	return
}

func loadChansConfig(path *string, mfc *mfConn) {

	if _, err := toml.DecodeFile(*path, mfc); err != nil {
		log.Fatalf("cannot load channels config %s \nuse tools/provision to create file", *path)
	}
}

func checkConnection(broker string, timeoutSecs int) {
	s := strings.Split(broker, ":")
	if len(s) != 3 {
		log.Fatalf("wrong host address format")
	}

	network := s[0]
	host := strings.Trim(s[1], "/")
	port := s[2]

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Duration(timeoutSecs)*time.Second)
	conClose := func() {
		if conn != nil {
			log.Println("closing connection")
			conn.Close()
		}
	}

	defer conClose()
	if err, ok := err.(*net.OpError); ok && err.Timeout() {
		fmt.Printf("Timeout error: %s\n", err.Error())
		log.Fatalf("Timeout error: %s\n", err.Error())
		return
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		log.Fatalf("Error: %s\n", err)
		return
	}

	log.Printf("Connection to %s://%s:%s looks ok\n", network, host, port)
}
