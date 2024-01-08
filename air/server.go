package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world wiz a!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	server := &http.Server{
		Handler: mux,
		Addr:    ":4000",
	}
	log.Fatal(server.ListenAndServe())
}
