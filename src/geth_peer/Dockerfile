FROM golang:alpine

#curl jsonrpc-glib
RUN apk add --no-cache bash make 
WORKDIR /app
COPY . .

RUN go build -o geth_peer_exporter main.go rpc.go

ENTRYPOINT [ "./geth_peer_exporter" ]
