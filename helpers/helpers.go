package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	// kita akan gunakan built in package math dengan function Intn yang menerima argument pool number
	value := rand.Intn(n)
	return value
}