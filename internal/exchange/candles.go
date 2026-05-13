package exchange

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/skiv71/hyperliquid/internal/candle"
)

type subscribe struct {
	Method       string              `json:"method"`
	Subscription candle.Subscription `json:"subscription"`
}

func subscribeMessage(coin *string, interval *string) []byte {

	d := subscribe{
		Method: `subscribe`,
		Subscription: candle.Subscription{
			Type:     `candle`,
			Coin:     *coin,
			Interval: *interval,
		},
	}

	j, e := json.Marshal(d)

	if e != nil {
		panic(e)
	}

	return j

}

type message struct {
	Channel string      `json:"channel"`
	Data    candle.Data `json:"data"`
}

func Candles(coin string, interval string, channel chan candle.Object) {

	u := url.URL{
		Scheme: "wss",
		Host:   "api.hyperliquid.xyz",
		Path:   "/ws",
	}

	c, _, err := websocket.DefaultDialer.Dial(
		u.String(),
		nil,
	)

	if err != nil {
		panic(err)
	}

	subscribe := subscribeMessage(&coin, &interval)

	log.Print("send: ", string(subscribe))

	if err := c.WriteMessage(1, subscribe); err != nil {
		panic(err)
	}

	for {
		_, data, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
			continue
		}

		m := message{}

		if err := json.Unmarshal(data, &m); err != nil {
			log.Println(err)
		}

		if m.Channel == `candle` {
			channel <- candle.Parse(&m.Data)
		} else {
			log.Print("receive: ", string(data))
		}
	}

}
