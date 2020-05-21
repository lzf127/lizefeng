package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/tiao2", tiao2)
	mux.HandleFunc("/tiao3", tiao3)
	mux.HandleFunc("/over", over)

	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}
	server.ListenAndServe()
}
