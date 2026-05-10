package main

import (
	"log"
	"strings"
	"sync"

	"github.com/sdomino/scribble"
	"github.com/skiv71/go-shit/internal/candle"
	"github.com/skiv71/go-shit/internal/exchange"
)

func main() {

	collection := "candles"

	db, err := scribble.New("./data/", nil)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	feed := make(chan candle.Object)

	wg.Go(func() {
		exchange.Feed(feed)
	})

	for {
		data := <-feed
		log.Print("candle: ", data)
		file := strings.Replace(data.Timestamp, " ", "_", 1)
		if err := db.Write(collection, file, &data); err != nil {
			log.Fatal(err)
		}
	}

}
