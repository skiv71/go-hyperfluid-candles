package candle

import (
	"fmt"
	"strconv"
	"time"
)

type Object struct {
	Coin      string `json:"coin"`
	Interval  int    `json:"interval"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Start    int64  `json:"t"`
	End      int64  `json:"T"`
	Coin     string `json:"s"`
	Interval string `json:"i"`
	Open     string `json:"o"`
	High     string `json:"h"`
	Low      string `json:"l"`
	Close    string `json:"c"`
	Volume   string `json:"v"`
}

type Subscription struct {
	Type     string `json:"type"`
	Coin     string `json:"coin"`
	Interval string `json:"interval"`
}

func parseInterval(s *string) int {
	v := 0
	suffix := map[string]int{
		"m": 60,
		"h": 60 * 60,
	}
	for _, r := range *s {
		c := string(r)
		i, err := strconv.ParseInt(c, 10, 64)
		if err == nil {
			v += int(i)
		} else {
			m, ok := suffix[c]
			if ok {
				return v * m
			}
		}
	}
	return v
}

func parseFloat(v *string) string {
	f, err := strconv.ParseFloat(*v, 32)
	if err != nil {
		panic("Invalid value!")
	}
	s := fmt.Sprintf("%.2f", f)
	return s
}

func parseTimestamp(v int64) string {
	t := time.UnixMilli(v)
	return t.Format(time.DateTime)
}

func Parse(data *Data) Object {
	return Object{
		Coin:      data.Coin,
		Interval:  parseInterval(&data.Interval),
		Open:      parseFloat(&data.Open),
		High:      parseFloat(&data.High),
		Low:       parseFloat(&data.Low),
		Close:     parseFloat(&data.Close),
		Volume:    parseFloat(&data.Volume),
		Timestamp: parseTimestamp(data.Start),
	}
}
