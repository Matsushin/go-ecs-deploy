package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepage(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(mainHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
