package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	base := pxconv.NewMetric(2, 1.5, 96)

	fmt.Printf("10dp = %dpx\n", base.DpToPx(pxconv.Dp(10)))
	fmt.Printf("15sp = %dpx\n", base.SpToPx(pxconv.Sp(15)))
	fmt.Printf("2inch = %dpx\n", base.InchToPx(pxconv.Inch(2)))
	fmt.Printf("50mm = %dpx\n", base.MmToPx(pxconv.Mm(50)))
	fmt.Printf("10pt = %dpx\n", base.PtToPx(pxconv.Pt(10)))

	fmt.Println()

	fmt.Printf("192px = %.2fdp\n", base.PxToDp(192))
	fmt.Printf("128px = %.2fsp\n", base.PxToSp(128))
	fmt.Printf("192px = %.2finch\n", base.PxToInch(192))
	fmt.Printf("189px = %.2fmm\n", base.PxToMm(189))
	fmt.Printf("10px = %.2fpt\n", base.PxToPt(10))

	fmt.Println()

	fmt.Printf("10dp = %dpx\n", base.DpToPx(pxconv.Dp(10)))
	fmt.Printf("10dp after scaling (DPI = %.2f): %dpx\n", base.Dpi, base.DpToPx(pxconv.Dp(10)))

	fmt.Println()

	fontSizes := []pxconv.Sp{12, 14, 16, 18, 24}
	for _, size := range fontSizes {
		fmt.Printf("Font size %.1fsp = %dpx\n", size, base.SpToPx(size))
	}

	fmt.Println()

	devices := []struct {
		Name string
		Dpi  float32
	}{
		{"Standard Display", 96},
		{"HD Display", 160},
		{"4K Display", 320},
	}

	for _, device := range devices {
		m := pxconv.NewMetric(base.PxPerDp, base.PxPerSp, device.Dpi)
		pxSize := m.DpToPx(pxconv.Dp(48))
		fmt.Printf("%s (%ddpi): %dpx\n", device.Name, int(device.Dpi), pxSize)
	}

	fmt.Println()

	buttonWidth := pxconv.Dp(200)
	buttonHeight := pxconv.Dp(50)

	fmt.Printf("Button size: %dx%d pixels\n",
		base.DpToPx(buttonWidth),
		base.DpToPx(buttonHeight),
	)

	fmt.Println()

	baseFontSize := pxconv.Sp(16)
	fmt.Printf("Base font size: %d pixels\n", base.SpToPx(baseFontSize))

	fmt.Println()

	pointValues := []pxconv.Pt{10, 12, 14, 16, 18, 24}
	for _, pt := range pointValues {
		fmt.Printf("%.1fpt = %dpx\n", pt, base.PtToPx(pt))
	}

	pixels := []int{10, 72, 96, 144}
	for _, px := range pixels {
		pt := base.PxToPt(px)
		fmt.Printf("%dpx = %.2fpt\n", px, pt)
	}

	fmt.Println()

	fontSizesPt := []pxconv.Pt{9, 10, 12, 14, 16, 18, 24}
	fmt.Println("Font sizes in pixels:")
	for _, pt := range fontSizesPt {
		fmt.Printf("%.1fpt = %dpx\n", pt, base.PtToPx(pt))
	}

	fmt.Println()

	pt := pxconv.Pt(72)
	px := base.PtToPx(pt)
	inch := base.PxToInch(px)
	mm := base.PxToMm(px)

	fmt.Printf("%.1fpt = %dpx\n", pt, px)
	fmt.Printf("%dpx = %.2finch\n", px, inch)
	fmt.Printf("%dpx = %.2fmm\n", px, mm)

	fmt.Println()

	dpis := []float32{96, 120, 160, 240}
	fontSize := pxconv.Pt(12)

	for _, dpi := range dpis {
		m := pxconv.NewMetric(2, 1.5, dpi)
		pxVal := m.PtToPx(fontSize)
		fmt.Printf("DPI: %.0f, %.1fpt = %dpx\n", dpi, fontSize, pxVal)
	}

	fmt.Println()

	buttonFont := pxconv.Pt(14)
	labelFont := pxconv.Pt(12)
	headerFont := pxconv.Pt(18)

	fmt.Printf("Button font size: %.1fpt = %dpx\n", buttonFont, base.PtToPx(buttonFont))
	fmt.Printf("Label font size: %.1fpt = %dpx\n", labelFont, base.PtToPx(labelFont))
	fmt.Printf("Header font size: %.1fpt = %dpx\n", headerFont, base.PtToPx(headerFont))
}
