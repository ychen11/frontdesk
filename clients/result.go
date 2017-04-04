package clients

import (
	"log"
)

type Result struct {
	Response string
	Latency  uint64
}

func buildResult(res string, lat uint64) *Result {
	res := Result{res, lat}
	return res
}
