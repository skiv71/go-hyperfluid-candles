# go-hyperfluid-candles

### Description

This is a small [Go](https://go.dev/) test project designed to retrieve candlestick data from [hyperfluid](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/websocket).

It does this via a [websocket](https://github.com/gorilla/websocket) connection and persists data in a local [json](https://github.com/sdomino/scribble) database.

### Test
```shell
... // coming soon!
```

### Build
```shell
go build
```

### Usage

```shell
./hyperfluid -coin=<coin> -interval=<1m|5m|15m|1h>

// (defaults)
// coin: xyz:SP500
// interval: 5m
```
