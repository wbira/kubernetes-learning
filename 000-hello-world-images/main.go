package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	data := []byte("OK")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/healthy", healthCheck)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		log.Printf("%s - [%s %s %s] 200 [%s]", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, userAgent)
		w.Header().Set("Server", "Hello/v1")
		fmt.Fprintf(w, "Hello, World!\n")
	})

	log.Println("Starting Hello service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
