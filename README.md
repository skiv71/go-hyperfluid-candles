# go-hyperfluid-candles

### Description

This is a short [Go](https://go.dev/) project for acquiring [hyperfluid](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/websocket) candles via [websocket](https://github.com/gorilla/websocket).

The data is parsed and written to a local [json](https://github.com/sdomino/scribble) database.

### Usage

To build / run

```shell
go [build / run] main.go -coin=<coin> -interval=<1m|5m|15m|1h>

// (defaults)
// coin: xyz:SP500
// interval: 5m
```
