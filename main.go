package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("hello")
	for i := 0; i < 1000; i++ {
		go GenerateStream()
	}
	for {
	}
}

func GenerateStream() {
	for {
		RandomString(25)
		time.Sleep(1 * time.Millisecond)
	}
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
