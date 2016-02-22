package goan

import (
	"net/http"
	"net/http/httptest"
)
/*
//performRequest is used mostly by tests to make requests
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
*/
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