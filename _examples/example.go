package main

import (
	"log"
	"time"

	retry "github.com/hellofresh/retry-go"
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
		log.Print("fn failed")
		return
	}

	log.Print("fn ok")
}

func fn() error {
	return nil
}
