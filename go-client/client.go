package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const frequency = 10 * time.Minute

func postIp(serverIp string, serverPort string, externalIp string) {
	form := make(url.Values)
	form.Set("ip", externalIp)
	resp, err := http.PostForm("http://"+serverIp+":"+serverPort+"/ip", form)
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Post request with application/x-www-form-urlencoded result: ", string(body))
}

func getExternalIp() string {

	responseClient, errClient := http.Get("http://myexternalip.com/raw") // get external IP
	if errClient != nil {
		log.Println("get external ip failed, check the net pls")
		panic(errClient)
	}

	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	externalIp := fmt.Sprintf("%s", string(body))
	log.Println("External IP: ", externalIp)
	return externalIp
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start main")

	serverIp := os.Getenv("SERVER_IP")
	if serverIp == "" {
		serverIp = "127.0.0.1"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8137"
	}

	ticker := time.NewTicker(frequency)
	for _ = range ticker.C {
		externalIp := getExternalIp()
		postIp(serverIp, serverPort, externalIp)
	}
}
