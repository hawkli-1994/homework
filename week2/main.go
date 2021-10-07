package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Server() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		for key, values := range r.Header {
			for _, v := range values {
				fmt.Printf("%s, %s", key, v)
				w.Header().Add(key, v)
			}
		}
		version := os.Getenv("VERSION")
		w.Header().Add("version", version)
		status := http.StatusOK
		requestIP := r.RemoteAddr
		defer func() {
			fmt.Printf("ip: %s\n", requestIP)
			fmt.Printf("statusCode: %d\n", status)
		}()
		w.WriteHeader(status)
		io.WriteString(w, "success")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "health")
	})
	http.ListenAndServe(":8000", nil)
}

func main() {
	Server()
}
