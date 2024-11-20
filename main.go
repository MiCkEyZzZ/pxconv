package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv/pxconv"
)

func main() {
	metric := pxconv.NewMetric(1.3333, 2.6667)

	// Преобразование с плавающей точкой
	dp := pxconv.Dp(7.5)
	sp := pxconv.Sp(3.25)
	fmt.Printf("%v dp = %v px\n", dp, metric.DpToPx(dp))
	fmt.Printf("%v sp = %v px\n", sp, metric.SpToPx(sp))
}
