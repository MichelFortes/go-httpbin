package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type data struct {
	Container string
	Method    string
	Url       string
	Query     map[string][]string
	Headers   map[string][]string
}

var port = os.Getenv("PORT")

func handleRoot(respWriter http.ResponseWriter, req *http.Request) {

	result := data{
		Container: os.Getenv("CONTAINER"),
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

	if port == "" {
		panic("Environment variable PORT not found")
	}

	http.HandleFunc("/", handleRoot)

	http.ListenAndServe(":"+port, nil)
}
