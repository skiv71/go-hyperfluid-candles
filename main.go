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

const COLLECTION = "candles"

func main() {

	cfg := config.Get()

	log.Print("config: ", cfg)

	db, err := scribble.New("./data/", nil)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	feed := make(chan candle.Object)

	wg.Go(func() {
		exchange.Feed(cfg.Coin, cfg.Interval, feed)
	})

	for {
		data := <-feed
		log.Print("candle: ", data)
		file := strings.Replace(data.Timestamp, " ", "_", 1)
		if err := db.Write(COLLECTION, file, &data); err != nil {
			log.Fatal(err)
		}
	}

}
