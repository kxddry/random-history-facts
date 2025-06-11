package main

import "time"

const baseDelay = time.Second
const exp = 2

func retry(attempts int, fn func() error) (err error) {
	for i := 0; i < attempts; i++ {
		if i == attempts-1 {
			break
		}
		err = fn()
		if err == nil {
			return
		}
		time.Sleep(baseDelay + time.Duration(exp*i)*time.Second)
	}
	return err
}
