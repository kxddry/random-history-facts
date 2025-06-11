package main

import (
	"log"
	"time"
)

const baseDelay = 1
const exp = 2
const maxDelay = 15

func retry(attempts int, fn func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i == attempts-1 {
			break
		}
		err = fn()
		if err == nil {
			return
		}
		log.Printf("migrator: retrying attempt #%d: %v", i, err)
		tm := time.Duration(min(maxDelay, baseDelay+(exp*i))) * time.Second
		time.Sleep(tm)
	}
	return err
}
