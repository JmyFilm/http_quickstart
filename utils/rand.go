package utils

import (
	"math/rand"
)

func NeedByPercent(percent float64) bool {
	if percent <= 0 {
		return false
	} else if percent >= 100 {
		return true
	}
	return rand.Float64()*100 <= percent
}
