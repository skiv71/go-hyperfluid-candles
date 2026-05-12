package config

import "testing"

func TestGet(t *testing.T) {
	cfg := Get()
	if cfg.Coin != DEFAULT_COIN {
		t.Errorf("Invalid Config.Coin: %s; expected: %s", cfg.Coin, DEFAULT_COIN)
	}
	if cfg.Interval != DEFAULT_INTERVAL {
		t.Errorf("Invalid Config.Interval: %s; expected: %s", cfg.Interval, DEFAULT_INTERVAL)
	}
}
