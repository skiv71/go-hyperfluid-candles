package config

import "flag"

type Config struct {
	Coin     string
	Interval string
}

const DEFAULT_COIN = `xyz:SP500`

const DEFAULT_INTERVAL = `5m`

func Get() Config {
	coin := flag.String("coin", DEFAULT_COIN, "-coin=<coin>")
	interval := flag.String("interval", DEFAULT_INTERVAL, "-interval=<1m|5m|15m|1h>")
	flag.Parse()
	return Config{
		Coin:     *coin,
		Interval: *interval,
	}
}
