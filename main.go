package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "stuff"

	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	time.Sleep(100 * time.Microsecond)
}
