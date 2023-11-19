package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type data struct {
	Container string
	RemoteAdd string
	Method    string
	Url       string
	Query     map[string][]string
	Headers   map[string][]string
}

func handleRoot(respWriter http.ResponseWriter, req *http.Request) {

	result := data{
		Container: os.Getenv("CONTAINER"),
		RemoteAdd: req.RemoteAddr,
		Method:    req.Method,
		Url:       req.URL.String(),
		Query:     req.URL.Query(),
		Headers:   req.Header,
	}

	err := json.NewEncoder(respWriter).Encode(result)
	if err != nil {
		println(err)
	}

}

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":80", nil)
}
