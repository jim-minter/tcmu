package main

import (
	"net/http"
	"net/http/httputil"
	"os"
)

type debugRoundTripper struct {
	http.RoundTripper
}

func (drt *debugRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	b, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return nil, err
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		return nil, err
	}

	rt := drt.RoundTripper
	if rt == nil {
		rt = http.DefaultTransport
	}

	resp, err := rt.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	b, err = httputil.DumpResponse(resp, false)
	if err != nil {
		return resp, err
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		return nil, err
	}

	return resp, err
}
