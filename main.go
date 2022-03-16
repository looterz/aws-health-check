package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting aws-health-check")

	addr := flag.String("address", "0.0.0.0", "local network address to listen on")
	port := flag.Int64("port", 27015, "The port to host the http server on")
	path := flag.String("path", "/health", "url path to respond to health checks on")
	flag.Parse()

	http.HandleFunc(fmt.Sprintf("%s", *path), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(""))
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("successfully responded to health check from %s", r.RemoteAddr)
	})

	log.Printf("Listening for HTTP health checks on endpoint %s:%d%s", *addr, *port, *path)

	var err = http.ListenAndServe(fmt.Sprintf("%s:%d", *addr, *port), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
