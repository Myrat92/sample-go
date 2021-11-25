package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start main")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Receive")
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Fprint(w, "error ip")
		}
		fmt.Fprint(w, addrs)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Handling HTTP requests on %s.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
