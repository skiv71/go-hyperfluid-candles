package config

import "flag"

type Config struct {
	Coin     string
	Interval string
	Stream   string
}

const DEFAULT_STREAM = `candles`

const DEFAULT_COIN = `xyz:SP500`

const DEFAULT_INTERVAL = `5m`

func Get() Config {
	stream := flag.String("stream", DEFAULT_STREAM, "-stream=<candles>")
	coin := flag.String("coin", DEFAULT_COIN, "-coin=<coin>")
	interval := flag.String("interval", DEFAULT_INTERVAL, "-interval=<1m|5m|15m|1h>")
	flag.Parse()
	return Config{
		Coin:     *coin,
		Interval: *interval,
		Stream:   *stream,
	}
}
