package vegeta

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 4 * time.Second
	targeter := NewCustomTargeter2()

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
		// fmt.Println(string(res.Body))

	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)

}

func NewCustomTargeter() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Set("Content-type", "application/x-www-form-urlencoded")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "https://gamebo-stage.gball8.com/api/gridz/gameBo/gameLog/get"

		form := url.Values{
			"api_key":   []string{"uC26PnoP1BiRQQHWHZK7wUTUYPjKlXBj"},
			"product":   []string{"n88club"},
			"start":     []string{"0"},
			"isLocal":   []string{"1"},
			"length":    []string{"100"},
			"end_at":    []string{"2022-06-01T00:00:00Z to 2022-08-01T23:59:59Z"},
			"player_id": []string{"30127"},
		}
		tgt.Body = []byte(form.Encode())
		return nil
	}
}

func NewCustomTargeter2() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Set("api_key", "zl9oo0vtrtbih1fufvq8pk8stvqprad8")
		header.Set("X-Lang", "en")
		header.Set("Content-Type", "application/json; charset=UTF-8")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "https://player-stage.gball8.com/api/nv88_api/v1/static_api/getGameList"

		var jsonStr = []byte(`{"currency": "VND", "game_category": [ "slots"], "limit": 0,  "offset": 0}`)	

		tgt.Body =  jsonStr
		return nil
	}
}
