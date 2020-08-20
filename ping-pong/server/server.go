package main

import (
	"fmt"
	"net/http"
)

type msg string

func main() {
http.HandleFunc("/play",handler)
fmt.Println("Listening on port 8181")
http.ListenAndServe(":8181",nil)
}
func handler(resp http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		msg := req.URL.Query().Get("game")
		if msg == "ping" {
			fmt.Fprintln(resp, "pong")
		} else {
			fmt.Println("error happened")
		}
	default:
		fmt.Fprintln(resp,"Not a GET method")
	}
}
