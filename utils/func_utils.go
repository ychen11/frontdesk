package utils

import (
	"time"
)

type function func()

func TimeFuncs(fn function) int64 {
	start := time.Now().UnixNano()
	fn()
	end := time.Now().UnixNano()

	return end - start
}
