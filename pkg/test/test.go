package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// NewRequest helps with testing handler requests
func NewRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	return req
}

// ServeRequest help with testing serving handler requests
func ServeRequest(h http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)

	handler.ServeHTTP(rr, req)

	return rr
}
