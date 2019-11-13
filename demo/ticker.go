package main

import "time"

func main() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			println("test")
		}
	}()

	time.Sleep(time.Minute)
}
