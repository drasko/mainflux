FROM golang:1.10-alpine AS builder
ARG svc
ARG goarch
ARG goarm
ENV GOARCH=$goarch
ENV GOARM=$goarm

WORKDIR /go/src/github.com/mainflux/mainflux
COPY . .
RUN apk update \
    && apk add make\
    && make $svc \
    && mv build/mainflux-$svc /exe

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]
