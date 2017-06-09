package main

import (
	"log"
	"time"

	"github.com/rafaeljesus/retry-go"
)

var (
	attempts  = 3
	sleepTime = time.Second * 2
)

func main() {
	if err := retry.Do(func() error {
		if err := fn(); err != nil {
			return err
		}
		return nil
	}, attempts, sleepTime); err != nil {
		log.Print("retry.Do Failed")
		return
	}

	log.Print("retry.Do OK")
}

func fn() error {
	return nil
}
