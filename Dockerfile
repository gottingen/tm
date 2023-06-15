FROM golang:1.20-alpine as builder

RUN apk add --no-cache \
    make \
    git \
    bash \
    curl \
    gcc \
    g++ \
    binutils-gold

# Install jq for tm-ctl
RUN cd / && \
    wget https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 -O jq && \
    chmod +x jq

RUN mkdir -p /go/src/github.com/gottingen/tm
WORKDIR /go/src/github.com/gottingen/tm

# Cache dependencies
COPY go.mod .
COPY go.sum .

RUN GO111MODULE=on go mod download

COPY . .

RUN make

FROM alpine:3.17

COPY --from=builder /go/src/github.com/gottingen/tm/bin/tm-server /tm-server
COPY --from=builder /go/src/github.com/gottingen/tm/bin/tm-ctl /tm-ctl
COPY --from=builder /go/src/github.com/gottingen/tm/bin/tm-recover /tm-recover
COPY --from=builder /jq /usr/local/bin/jq

RUN apk add --no-cache \
    curl

EXPOSE 2379 2380

ENTRYPOINT ["/tm-server"]
