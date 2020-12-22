package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"github.com/atotto/clipboard"
)
var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func main() {
	size := flag.Int("size", 20, "an int")
	flag.Parse()
	s := randSeq(*size)
	fmt.Println(s)
	clipboard.WriteAll(s)
}