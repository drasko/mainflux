package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/mainflux/mainflux/tools/mqtt-bench"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

type mqttBrokerCfg struct {
	url string `toml:"url"`
}

type mqttMessageCfg struct {
	size   int    `toml:"size"`
	format string `toml:"format"`
	qos    int    `toml:"qos"`
	retain bool   `toml:"retain"`
}

type mqttTLSCfg struct {
	mtls       bool   `toml:"mtls"`
	skipTLSVer bool   `toml:"skiptlsver"`
	ca         string `toml:"ca"`
}

type mqttCfg struct {
	broker  mqttBrokerCfg  `toml:"broker"`
	message mqttMessageCfg `toml:"message"`
	tls     mqttTLSCfg     `toml:"tls"`
}

type testCfg struct {
	count int `toml:"count"`
	pubs  int `toml:"pubs"`
	subs  int `toml:"subs"`
}

type logCfg struct {
	quiet bool `toml:"quiet"`
}

type mainfluxCfg struct {
	channelID string `toml:"channelID"`
	thingID   string `toml:"thingID"`
	thingKey  string `toml:"thingKey"`
	mtlsCert  string `toml:"mtlsCert"`
	mtlsKey   string `toml:"mtlsKey"`
}

type config struct {
	mqtt mqttCfg       `toml:"mqtt"`
	test testCfg       `toml:"test"`
	log  logCfg        `toml:"log"`
	mf   []mainfluxCfg `toml:"mainflux"`
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
		runBench()
	},
}

func init() {

	cobra.OnInitialize(initConfig)

	// MQTT Broker
	benchCmd.PersistentFlags().StringVarP(&cfg.mqtt.broker.url, "broker", "b", "tcp://localhost:1883",
		"address for mqtt broker, for secure use tcps and 8883")

	// MQTT Message
	benchCmd.PersistentFlags().IntVarP(&cfg.mqtt.message.size, "size", "s", 100, "Size of message payload bytes")
	benchCmd.PersistentFlags().StringVarP(&cfg.mqtt.message.format, "format", "f", "text", "Output format: text|json")
	benchCmd.PersistentFlags().IntVarP(&cfg.mqtt.message.qos, "qos", "q", 0, "QoS for published messages, values 0 1 2")
	benchCmd.PersistentFlags().BoolVarP(&cfg.mqtt.message.retain, "retain", "r", false, "Retain mqtt messages")

	// MQTT TLS
	benchCmd.PersistentFlags().BoolVarP(&cfg.mqtt.tls.mtls, "mtls", "m", false, "Use mtls for connection")
	benchCmd.PersistentFlags().BoolVarP(&cfg.mqtt.tls.skipTLSVer, "skipTLSVer", "t", false, "Skip tls verification")
	benchCmd.PersistentFlags().StringVarP(&cfg.mqtt.tls.ca, "ca", "", "ca.crt", "CA file")

	// Test params
	benchCmd.PersistentFlags().IntVarP(&cfg.test.count, "count", "n", 100, "Number of messages sent per publisher")
	benchCmd.PersistentFlags().IntVarP(&cfg.test.subs, "subs", "", 10, "Number of subscribers")
	benchCmd.PersistentFlags().IntVarP(&cfg.test.pubs, "pubs", "", 10, "Number of publishers")
	benchCmd.PersistentFlags().StringVarP(&cfgFile, "config", "g", "config.toml", "config file default is config.toml")

	// Log params
	benchCmd.PersistentFlags().BoolVarP(&cfg.log.quiet, "quiet", "", false, "Supress messages")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Failed to load config - %s", err.Error())
	}
}

func runBench() {
	var wg sync.WaitGroup
	var err error
	var caByte []byte

	subTimes := make(bench.SubTimes)

	if cfg.test.pubs < 1 && cfg.test.subs < 1 {
		log.Fatal("Invalid arguments")
	}

	if cfg.mqtt.tls.mtls {
		caFile, err := os.Open(cfg.mqtt.tls.ca)
		defer caFile.Close()

		if err != nil {
			fmt.Println(err)
		}

		caByte, _ = ioutil.ReadAll(caFile)
	}

	resCh := make(chan *bench.RunResults)
	done := make(chan bool)

	start := time.Now()
	n := len(cfg.mf)
	cert := tls.Certificate{}

	for i := 0; i < cfg.test.subs; i++ {
		mf := cfg.mf[i%n]

		if cfg.mqtt.tls.mtls {
			cert, err = tls.X509KeyPair([]byte(mf.mtlsCert), []byte(mf.mtlsKey))
			if err != nil {
				log.Fatal(err)
			}
		}

		c := &bench.Client{
			ID:         strconv.Itoa(i),
			BrokerURL:  cfg.mqtt.broker.url,
			BrokerUser: mf.thingID,
			BrokerPass: mf.thingKey,
			MsgTopic:   fmt.Sprintf("channels/%s/messages/test", mf.channelID),
			MsgSize:    cfg.mqtt.message.size,
			MsgCount:   cfg.test.count,
			MsgQoS:     byte(cfg.mqtt.message.qos),
			Quiet:      cfg.log.quiet,
			Mtls:       cfg.mqtt.tls.mtls,
			SkipTLSVer: cfg.mqtt.tls.skipTLSVer,
			CA:         caByte,
			ClientCert: cert,
			Retain:     cfg.mqtt.message.retain,
		}
		wg.Add(1)
		go c.RunSubscriber(&wg, &subTimes, &done, cfg.mqtt.tls.mtls)
	}
	wg.Wait()

	for i := 0; i < cfg.test.pubs; i++ {
		mf := cfg.mf[i%n]

		if cfg.mqtt.tls.mtls {
			cert, err = tls.X509KeyPair([]byte(mf.mtlsCert), []byte(mf.mtlsKey))
			if err != nil {
				log.Fatal(err)
			}
		}

		c := &bench.Client{
			ID:         strconv.Itoa(i),
			BrokerURL:  cfg.mqtt.broker.url,
			BrokerUser: mf.thingID,
			BrokerPass: mf.thingKey,
			MsgTopic:   fmt.Sprintf("channels/%s/messages/test", mf.channelID),
			MsgSize:    cfg.mqtt.message.size,
			MsgCount:   cfg.test.count,
			MsgQoS:     byte(cfg.mqtt.message.qos),
			Quiet:      cfg.log.quiet,
			Mtls:       cfg.mqtt.tls.mtls,
			SkipTLSVer: cfg.mqtt.tls.skipTLSVer,
			CA:         caByte,
			ClientCert: cert,
			Retain:     cfg.mqtt.message.retain,
		}

		go c.RunPublisher(resCh, cfg.mqtt.tls.mtls)
	}

	// Collect the results
	var results []*bench.RunResults
	if cfg.test.pubs > 0 {
		results = make([]*bench.RunResults, cfg.test.pubs)
	}

	for i := 0; i < cfg.test.pubs; i++ {
		results[i] = <-resCh
	}

	totalTime := time.Now().Sub(start)
	totals := calculateTotalResults(results, totalTime, &subTimes)
	if totals == nil {
		return
	}

	// Print sats
	printResults(results, totals, cfg.mqtt.message.format, cfg.log.quiet)
}

func calculateTotalResults(results []*bench.RunResults, totalTime time.Duration, subTimes *bench.SubTimes) *bench.TotalResults {
	if results == nil || len(results) < 1 {
		return nil
	}
	totals := new(bench.TotalResults)
	totals.TotalRunTime = totalTime.Seconds()
	var subTimeRunResults bench.RunResults
	msgTimeMeans := make([]float64, len(results))
	msgTimeMeansDelivered := make([]float64, len(results))
	msgsPerSecs := make([]float64, len(results))
	runTimes := make([]float64, len(results))
	bws := make([]float64, len(results))

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
			log.Printf("Failed to prepare results for printing - %s", err.Error())
		}
		var out bytes.Buffer
		json.Indent(&out, data, "", "\t")

		fmt.Println(string(out.Bytes()))
	default:
		if !quiet {
			for _, res := range results {
				fmt.Printf("======= CLIENT %s =======\n", res.ID)
				fmt.Printf("Ratio:               %.3f (%d/%d)\n",
					float64(res.Successes)/float64(res.Successes+res.Failures), res.Successes, res.Successes+res.Failures)
				fmt.Printf("Runtime (s):         %.3f\n", res.RunTime)
				fmt.Printf("Msg time min (us):   %.3f\n", res.MsgTimeMin)
				fmt.Printf("Msg time max (us):   %.3f\n", res.MsgTimeMax)
				fmt.Printf("Msg time mean (us):  %.3f\n", res.MsgTimeMean)
				fmt.Printf("Msg time std (us):   %.3f\n", res.MsgTimeStd)

				fmt.Printf("Bandwidth (msg/sec): %.3f\n\n", res.MsgsPerSec)
			}
		}
		fmt.Printf("========= TOTAL (%d) =========\n", len(results))
		fmt.Printf("Total Ratio:                 %.3f (%d/%d)\n",
			totals.Ratio, totals.Successes, totals.Successes+totals.Failures)
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

func main() {
	if err := benchCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
