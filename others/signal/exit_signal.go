package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create a sig channel to catch system signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	// change http service to async
	go http.ListenAndServe(":8080", mux)

	// this one will wait for sig
	fmt.Println(<-sig)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
