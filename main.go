package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting aws-health-check")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "healthy\n")
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("successfully responded to health check from %s", r.RemoteAddr)
	})

	port := flag.Int64("port", 27015, "The port to host the http server on")
	flag.Parse()

	log.Printf("Listening for HTTP health checks on port %d", *port)

	var err = http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
