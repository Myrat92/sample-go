package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start main")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8137"
	}

	http.HandleFunc("/ip", ip)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func ip(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	externalIp, err := request.Form["ip"]
	if !err {
		log.Fatal("receive ip error")
	}

	log.Println(externalIp[0])
}
