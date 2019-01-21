package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

var req *http.Request
var requestDump []byte

func catchHandler(w http.ResponseWriter, r *http.Request) {
	req = r
	requestDump, _ = httputil.DumpRequest(req, true)
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
	if req != nil {
		req.ParseForm()
		fmt.Fprintf(w, "Form Values: %#v\n\n", req.Form)
		fmt.Fprintf(w, "Remote Address: %v\n\n", req.RemoteAddr)
	}
	fmt.Fprintf(w, string(requestDump))
}

func main() {
	catchMux := http.NewServeMux()
	catchMux.HandleFunc("/", catchHandler)
	go http.ListenAndServe(":10101", catchMux)

	inspectMux := http.NewServeMux()
	inspectMux.HandleFunc("/", inspectHandler)
	http.ListenAndServe(":10102", inspectMux)
}
