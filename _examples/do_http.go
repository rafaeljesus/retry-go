package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rafaeljesus/retry-go"
)

var (
	attempts  = 3
	sleepTime = time.Second * 2
)

func main() {
	_, err := retry.DoHTTP(func() (*http.Response, error) {
		return fn()
	}, attempts, sleepTime)
	if err != nil {
		log.Print("retry.DoHTTP Failed")
		return
	}

	log.Print("retry.DoHTTP OK")
}

func fn() (*http.Response, error) {
	res := &http.Response{}
	return res, nil
}
