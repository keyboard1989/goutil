package goutil

import (
	"math/rand"
	"time"
)

func GetRandInt32() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()
}

func GetRandInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func GetResidueRandom(num int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % num
}

func GetRanddomString(num int) string {
	return ""
}
