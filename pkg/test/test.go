package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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
func ServeRequest(routerURL string, h http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(routerURL, h)
	router.ServeHTTP(rr, req)

	return rr
}
