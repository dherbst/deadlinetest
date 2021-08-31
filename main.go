package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func nowait(urls []string) {

	transport := &http.Transport{
		IdleConnTimeout:       5 * time.Second,
		MaxIdleConns:          100,
		ExpectContinueTimeout: 10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:  1 * time.Second,
			Deadline: time.Now().Add(20 * time.Second),
		}).DialContext,
	}
	client := &http.Client{Transport: transport}

	for i, v := range urls {
		u := "https://" + v + "/"
		fmt.Println(i, u)
		_, err := client.Get(u)
		if err != nil {
			log.Print("Error Get ", u, err)
		}

	}
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal("Error closing file", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	nowait(urls)
}
