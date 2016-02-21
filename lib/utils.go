package goan

import (
	"net/http"
	"net/http/httptest"
)

//performRequest is used mostly by tests to make requests
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
