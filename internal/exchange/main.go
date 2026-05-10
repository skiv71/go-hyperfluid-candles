package exchange

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/skiv71/go-shit/internal/candle"
)

type Subscribe struct {
	Method       string              `json:"method"`
	Subscription candle.Subscription `json:"subscription"`
}

func subscribeMessage() []byte {

	d := Subscribe{
		Method: `subscribe`,
		Subscription: candle.Subscription{
			Type:     `candle`,
			Coin:     `xyz:SP500`,
			Interval: `5m`,
		},
	}

	j, e := json.Marshal(d)

	if e != nil {
		panic(e)
	}

	return j

}

type Message struct {
	Channel string      `json:"channel"`
	Data    candle.Data `json:"data"`
}

func Feed(feed chan candle.Object) {

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

	subscribe := subscribeMessage()

	log.Print("send: ", string(subscribe))

	if err := c.WriteMessage(1, subscribe); err != nil {
		panic(err)
	}

	for {
		_, data, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
		}

		m := Message{}

		if err := json.Unmarshal(data, &m); err != nil {
			log.Println(err)
		}

		if m.Channel == `candle` {
			feed <- candle.Parse(&m.Data)
		} else {
			log.Print("receive: ", string(data))
		}
	}

}
