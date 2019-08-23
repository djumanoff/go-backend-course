package main

import (
	"net/http"
)

type Requester struct {
	clt *http.Client
}

func (reqs *Requester) RoundRobbinRequester(method, url string, backends []string, n int) (string, *http.Response, error) {
	backend := backends[n % len(backends)]
	req, err := http.NewRequest(method, backend + url, nil)
	if err != nil {
		return "", nil, err
	}
	resp, err := reqs.clt.Do(req)

	return backend, resp, err
}
