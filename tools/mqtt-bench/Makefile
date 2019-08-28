# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

SOURCES = $(wildcard *.go) cmd/main.go

all: mqtt-bench

.PHONY: all clean

mqtt-bench: $(SOURCES)
	go build -ldflags "-s -w" -o $@ cmd/main.go

clean:
	rm -rf mqtt-bench
