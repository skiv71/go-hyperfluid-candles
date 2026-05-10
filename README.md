# go-hyperfluid-candles

### Description

This is a short [Go](https://go.dev/) project for acquiring [hyperfluid](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/websocket) candles via [websocket](https://github.com/gorilla/websocket).

The data is parsed and written to a local [json](https://github.com/sdomino/scribble) database.

### Configuration

The default config is: -

```go
COIN = "xyz:SP500"
INTERVAL = 300 // (5m)
```

This can be changed in the ./main.go file

### Usage

To build / run

```shell
go [build / run] main.go
```
