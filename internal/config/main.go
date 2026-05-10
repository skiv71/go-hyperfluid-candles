package config

import "flag"

type Config struct {
	Coin     string
	Interval string
}

func Get() Config {
	coin := flag.String("coin", "xyz:SP500", "-coin=<coin>")
	interval := flag.String("interval", "5m", "-interval=<1m|5m|15m|1h>")
	flag.Parse()
	return Config{
		Coin:     *coin,
		Interval: *interval,
	}
}
