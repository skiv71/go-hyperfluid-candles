package candle

import (
	"encoding/json"
	"testing"

	"github.com/skiv71/hyperliquid/internal/config"
)

func TestParse(t *testing.T) {
	data := Data{
		Start:    10e6,
		End:      10e12,
		Coin:     config.DEFAULT_COIN,
		Interval: config.DEFAULT_INTERVAL,
		Open:     "3",
		High:     "4",
		Low:      "1",
		Close:    "2",
		Volume:   "10",
	}
	raw, err := json.Marshal(data)
	if err != nil {
		t.Error("Failed to marshal candle data!")
	}
	o := Data{}
	if err := json.Unmarshal(raw, &o); err != nil {
		t.Error("Failed to unmarshall candle data!")
	}
	c := Parse(&o)
	if o.Coin != data.Coin {
		t.Errorf("Invalid *.Coin: %s; expected: %s", o.Coin, data.Coin)
	}
	if o.Interval != data.Interval {
		t.Errorf("Invalid *.Interval: %s; expected: %s", o.Interval, data.Interval)
	}
	if parseFloat(&o.Open) != c.Open {
		t.Errorf("Invalid Candle.Open: %s; expected: %s", parseFloat(&o.Open), c.Open)
	}
	if parseFloat(&o.High) != c.High {
		t.Errorf("Invalid Candle.High: %s; expected: %s", parseFloat(&o.High), c.High)
	}
	if parseFloat(&o.Low) != c.Low {
		t.Errorf("Invalid Candle.Low: %s; expected: %s", parseFloat(&o.Low), c.Low)
	}
	if parseFloat(&o.Close) != c.Close {
		t.Errorf("Invalid Candle.Close: %s; expected: %s", parseFloat(&o.Close), c.Close)
	}
	if parseFloat(&o.Volume) != c.Volume {
		t.Errorf("Invalid Candle.Volume: %s; expected: %s", parseFloat(&o.Volume), c.Volume)
	}
	if parseTimestamp(o.Start) != c.Timestamp {
		t.Errorf("Invalid Candle.Timestamp: %s; expected: %s", parseTimestamp(o.Start), c.Timestamp)
	}
}
