package main

import (
	"log"
	"strings"
	"sync"

	"github.com/sdomino/scribble"
	"github.com/skiv71/hyperliquid/internal/candle"
	"github.com/skiv71/hyperliquid/internal/config"
	"github.com/skiv71/hyperliquid/internal/exchange"
)

func streamCandles(cfg *config.Config) {

	db, err := scribble.New("./data/", nil)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	channel := make(chan candle.Object)

	wg.Go(func() {
		exchange.Candles(cfg.Coin, cfg.Interval, channel)
	})

	for {
		data := <-channel
		log.Print("candle: ", data)
		file := strings.Replace(data.Timestamp, " ", "_", 1)
		if err := db.Write(cfg.Stream, file, &data); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {

	cfg := config.Get()

	log.Print("config: ", cfg)

	switch cfg.Stream {
	case "candles":
		streamCandles(&cfg)
	default:
		log.Fatalf("Invalid stream: %s", cfg.Stream)
	}

}
