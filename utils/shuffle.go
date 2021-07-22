package utils

import (
	"math/rand"
	"time"
)

// shuffle slice的具体实现
func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

// ShuffleSlice 打乱slice。
func ShuffleSlice(list []int) {
	shuffle(list)
}