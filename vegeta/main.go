package vegeta

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Run() {
	rate := vegeta.Rate{Freq: 1200, Per: time.Second}
	duration := 6 * time.Second
	targeter := NewCustomTargeter8()

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	count := 0
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
		//fmt.Println(string(res.Body))
		// fmt.Println(count)
		count++
	}
	metrics.Close()

	reporter := vegeta.NewTextReporter(&metrics)
	reporter(os.Stdout)

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)

}

func NewCustomTargeter3() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		tgt.Method = "GET"
		tgt.URL = "https://sbo-stage.gball8.com"
		return nil
	}
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

		tgt.URL = "http://nv88-gbo-stage.n88club.local/api/gridz/gameBo/gameLog/get"

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

		tgt.Body = jsonStr
		return nil
	}
}

func NewCustomTargeter4() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Set("api_key", "zl9oo0vtrtbih1fufvq8pk8stvqprad8")
		header.Set("Content-Type", "application/json")
		header.Set("accept", "application/json")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "http://10.11.101.80/api/nv88_api/v1/player_api/getProviderAndGameTypeCommission"

		var jsonStr = []byte(`{"currency": "VND"}`)

		tgt.Body = jsonStr
		return nil
	}
}

func NewCustomTargeter5() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		header.Set("accept", "application/x-www-form-urlencoded")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "https://gamebo-next.gball8.com/api/gridz/player/queryBalance/new"

		// var jsonStr = []byte(`{"api_key": "uC26PnoP1BiRQQHWHZK7wUTUYPjKlXBj", "product": "n88club", "player_name": "crkevinp1", "player_id": "30178", "currency": "VND",  "agent_username": "kevinagt1","member_type": "CREDIT", "platform_code": "ADVP",  "isLocal": "true" }`)

		// tgt.Body = jsonStr

		requestData := url.Values{
			"api_key":        {"uC26PnoP1BiRQQHWHZK7wUTUYPjKlXBj"},
			"product":        {"n88club"},
			"player_name":    {"crkevinp1"},
			"player_id":      {"30178"},
			"currency":       {"VND"},
			"agent_username": {"kevinagt1"},
			"member_type":    {"CREDIT"},
			"platform_code":  {"ADVP"},
			"isLocal":        {"true"},
		}
		tgt.Body = []byte(requestData.Encode())

		return nil
	}
}

func NewCustomTargeter6() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Add("api_key", "zl9oo0vtrtbih1fufvq8pk8stvqprad8")
		header.Add("X-Lang", "en")
		header.Add("Content-Type", "application/json")
		header.Add("Accept", "application/json")
		header.Add("Authorization", "Bearer 0af82939f4a346070b1149570ad04dd0")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "http://10.11.101.80/api/nv88_api/v1/player_api/getPlayerBalance"

		var jsonStr = []byte(`{"username": "testken1"}`)

		tgt.Body = jsonStr
		return nil
	}
}

func NewCustomTargeter7() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		token := "c401dd8f9818f1edf43672c2cdb98e68"

		header := http.Header{}
		header.Add("Content-Type", "application/json")
		header.Add("Accept", "application/json")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "https://gamebo-next.gball8.com/api/JILI/cancelBet"

		var jsonStr = []byte(`{"reqId":"f8566411-2276-40cf-9e12-4edd0f70a4f3","token":"` + token + `","currency":"VND","game":10,"round":1660550052000061010,"betAmount":2000,"winloseAmount":400,"userId":"N88cremember02"}`)

		tgt.Body = jsonStr

		return nil
	}
}

func NewCustomTargeter8() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}

		header := http.Header{}
		header.Set("api_key", "zl9oo0vtrtbih1fufvq8pk8stvqprad8")
		header.Set("Content-Type", "application/json")
		// header.Set("accept", "application/x-www-form-urlencoded")
		tgt.Header = header

		tgt.Method = "POST"

		tgt.URL = "http://seamless.n88club.local/api/NESP/balance?authToken=ebe4448e95480931a82a16b5f1ecc73d"

		// var jsonStr = []byte(`{"api_key": "uC26PnoP1BiRQQHWHZK7wUTUYPjKlXBj", "product": "n88club", "player_name": "crkevinp1", "player_id": "30178", "currency": "VND",  "agent_username": "kevinagt1","member_type": "CREDIT", "platform_code": "ADVP",  "isLocal": "true" }`)

		// tgt.Body = jsonStr

		jsonStr := []byte(`{"sid":"9755cb02fcc0530f904ccd308087833c","userId":"N88TECHTEST22","game":{"type":"slots","details":{"table":{"id":null,"vid":null}}},"currency":"VND","uuid":"b4b5baaa-c9fb-42dc-b975-7f2640e1d483"}`)

		tgt.Body = jsonStr
		return nil

	}
}
