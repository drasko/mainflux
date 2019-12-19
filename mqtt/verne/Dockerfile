# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

FROM erlang:22.2-alpine AS builder
WORKDIR /mainflux

COPY ./mqtt/verne .
RUN apk add --no-cache git && \
  ./rebar3 compile

FROM vernemq/vernemq:1.10.1-alpine
WORKDIR /mainflux
COPY --from=builder --chown=10000:10000 /mainflux .
