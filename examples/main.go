package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	metric := pxconv.Metric{PxPerDp: 2, PxPerSp: 1.5}
	px := metric.DpToPx(pxconv.Dp(10))
	fmt.Printf("10dp = %dpx\n", px)
}
