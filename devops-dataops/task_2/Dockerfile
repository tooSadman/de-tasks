FROM golang:1.18.0-alpine3.15 AS gobuild
RUN apk add --no-cache build-base
COPY go.* /tmp/go-build/
RUN cd /tmp/go-build/ \
	&& go mod download
COPY main.go /tmp/go-build/
RUN cd /tmp/go-build \
	&& go build -a -tags musl main.go

FROM alpine:3.15.4 AS base
COPY --from=gobuild /tmp/go-build/main /usr/local/bin/task2-server
CMD ["/usr/local/bin/task2-server"]
