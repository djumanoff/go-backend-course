package main

import (
	"net/http"
	"testing"
)

type testCase struct {
	method string
	url string
	n int
	backends []string
	resultHost string
}

var rqstr = &Requester{
	&http.Client{},
}

var tests = []testCase{
	{
		"GET",
		"/",
		0,
		[]string{"https://google.com", "https://facebook.com"},
		"https://google.com",
	},
	{
		"GET",
		"/",
		1,
		[]string{"https://google.com", "https://facebook.com"},
		"https://facebook.com",
	},
}

func TestRoundRobbinRequester(t *testing.T) {
	for _, test := range tests {
		backend, _, err := rqstr.RoundRobbinRequester(test.method, test.url, test.backends, test.n)
		if err != nil {
			t.Fatal(err)
		}
		if backend != test.resultHost {
			t.Fail()
		}
	}
}
