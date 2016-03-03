package goan

import (
	"net/http"
	"net/http/httptest"
)

//performRequest performs an HTTP request for recording, used in testing
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	var req *http.Request
    if method == "POST" {
        req, _ = http.NewRequest("POST", path, nil)
    } else if method == "GET" {
        req, _ = http.NewRequest("GET", path, nil)
    }
    
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}