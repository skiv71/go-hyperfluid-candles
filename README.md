# hyperliquid-go

### Description

This is a small [Go](https://go.dev/) test project to interact with [hyperliquid](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api).

At the moment, live candles can streamed via a [websocket](https://github.com/gorilla/websocket) connection and persisted in a local [json](https://github.com/sdomino/scribble) database.

### Test
```shell
go test ./...
```

### Build
```shell
go build
```

### Usage

```shell
./hyperliquid -stream=candles -coin=<coin> -interval=<1m|5m|15m|1h>

// (defaults)
// stream: candles
// coin: xyz:SP500
// interval: 5m
```
