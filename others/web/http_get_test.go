package http_get

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	resp, _ := http.Get("http://example.com")
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	resp.Body.Close()
}
