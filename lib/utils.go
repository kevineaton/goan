package goan

import (
    "log"
    "os"
	"net/http"
	"net/http/httptest"
)

var (
    //LogInfo logs at the Info level (default Stdout)
    LogInfo *log.Logger
    //LogWarning logs at the Warning level (default Stdout)
    LogWarning *log.Logger
    //LogError logs at the Error level (default Stderr)
    LogError *log.Logger
)

//SetupLogger sets up logging for the application
func SetupLogger () {
    LogInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    LogWarning = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
    LogError = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

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