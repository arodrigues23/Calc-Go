package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func openServer() {

	servidor := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      FunctionHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(servidor.ListenAndServe())

}

type FunctionHandler struct{}

func (f FunctionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if !ValidServer(req) {
		res.Write(MessageToJson(TableMessage[404]))
		return
	}

	Method[req.Method][req.URL.Path](res, req)
}

type HTTPMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var TableMessage = map[int]HTTPMessage{
	404: {404, "Not Found"},
	500: {500, "Internal Server Error"},
}

func MessageToJson(m HTTPMessage) []byte {
	json, _ := json.Marshal(m)
	return json
}

func ValidServer(req *http.Request) bool {
	return Method[req.Method][req.URL.Path] != nil
}
