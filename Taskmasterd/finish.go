package main

import (
	"time"
)

func finish(fifi float64, cc, exe int) (a, b bool, g time.Time, d float64, y bool, i int) {
	return false, true, time.Now(), fifi, fifi >= float64(time.Second*time.Duration(cc)), exe + 1
}
