FROM golang:latest

WORKDIR /

RUN wget https://bitcoin.org/bin/bitcoin-core-0.19.1/bitcoin-0.19.1-x86_64-linux-gnu.tar.gz
RUN tar xzf bitcoin-0.19.1-x86_64-linux-gnu.tar.gz
RUN install -m 0755 -o root -g root -t /usr/local/bin bitcoin-0.19.1/bin/*
RUN bitcoind -daemon
RUN bitcoind -daemon -testnet

COPY . /app

WORKDIR /app
RUN go build

ENTRYPOINT ["/app/bitcarve"]
