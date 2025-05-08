package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	metric := pxconv.NewMetric(2, 1.5, 96)

	// Conversion between units.

	// Conversion examples.
	fmt.Printf("10dp = %dpx\n", metric.DpToPx(pxconv.Dp(10)))     // 10dp = 20px
	fmt.Printf("15sp = %dpx\n", metric.SpToPx(pxconv.Sp(15)))     // 15sp = 23px
	fmt.Printf("2inch = %dpx\n", metric.InchToPx(pxconv.Inch(2))) // 2inch = 192px
	fmt.Printf("50mm = %dpx\n", metric.MmToPx(pxconv.Mm(50)))     // 50mm = 189px
	fmt.Printf("10pt = %dpx\n", metric.PtToPx(pxconv.Pt(10)))     // 10pt = 13px

	fmt.Println()

	// Inverse conversions.
	fmt.Printf("192px = %.2fdp\n", metric.PxToDp(192))     // 192px = 96.00dp
	fmt.Printf("128px = %.2fsp\n", metric.PxToSp(128))     // 128px = 85.33sp
	fmt.Printf("192px = %.2finch\n", metric.PxToInch(192)) // 192px = 2.00inch
	fmt.Printf("189px = %.2fmm\n", metric.PxToMm(189))     // 189px = 50.01mm
	fmt.Printf("10px = %.2fpt\n", metric.PxToMm(10))       // 10px = 2.65pt

	fmt.Println()

	// Scaling for a device with a non-standard DPI.

	// Conversion with current parameters.
	fmt.Printf("10dp = %dpx\n", metric.DpToPx(pxconv.Dp(10)))

	// DPI scaling.
	fmt.Printf("10dp after scaling (DPI = %.2f): %dpx\n", metric.Dpi, metric.DpToPx(pxconv.Dp(10)))

	fmt.Println()

	// Font setup.

	// Font sizes in sp and their pixel equivalents.
	fontSizes := []pxconv.Sp{12, 14, 16, 18, 24}
	for _, size := range fontSizes {
		fmt.Printf("Font size %.1fsp = %dpx\n", size, metric.SpToPx(size))
	}

	fmt.Println()

	// Choosing element sizes for different devices.

	// Simulate devices with different DPIs
	devices := []struct {
		Name string
		Dpi  float32
	}{
		{"Standard Display", 96},
		{"HD Display", 160},
		{"4K Display", 320},
	}

	baseMetric := pxconv.NewMetric(1, 1, pxconv.DefaultDpi)
	elementSize := pxconv.Dp(48) // Base element size in dp

	for _, device := range devices {
		metric := pxconv.NewMetric(baseMetric.PxPerDp, baseMetric.PxPerSp, device.Dpi)
		pxSize := metric.DpToPx(elementSize)
		fmt.Printf("%s (%ddpi): %dpx\n", device.Name, int(device.Dpi), pxSize)
	}

	fmt.Println()

	// Usage in a real application.

	buttonWidth := pxconv.Dp(200)
	buttonHeight := pxconv.Dp(50)

	fmt.Printf("Button size: %dx%d pixels\n",
		metric.DpToPx(buttonWidth),
		metric.DpToPx(buttonHeight),
	)

	fmt.Println()

	// Font size calculation.
	baseFontSize := pxconv.Sp(16)
	fmt.Printf("Base font size: %d pixels\n", metric.SpToPx(baseFontSize))

	fmt.Println()

	// Convert points to pixels.

	pointValues := []pxconv.Pt{10, 12, 14, 16, 18, 24}
	for _, pt := range pointValues {
		px := metric.PtToPx(pt)
		fmt.Printf("%.1fpt = %dpx\n", pt, px)
	}

	// Reverse conversion.
	pixels := []int{10, 72, 96, 144}
	for _, px := range pixels {
		pt := metric.PxToPt(px)
		fmt.Printf("%dpx = %.2fpt\n", px, pt)
	}

	fmt.Println()

	// Using points for font sizes.

	// Размеры текста в пунктах
	fontSizs := []pxconv.Pt{9, 10, 12, 14, 16, 18, 24}
	fmt.Println("Font sizes in pixels:")
	for _, pt := range fontSizs {
		fmt.Printf("%.1fpt = %dpx\n", pt, metric.PtToPx(pt))
	}

	fmt.Println()

	// Converting between points and other units.

	// Example: Convert 72 points to other units.
	pt := pxconv.Pt(72)

	px := metric.PtToPx(pt)
	inch := metric.PxToInch(px)
	mm := metric.PxToMm(px)

	fmt.Printf("%.1fpt = %dpx\n", pt, px)
	fmt.Printf("%dpx = %.2finch\n", px, inch)
	fmt.Printf("%dpx = %.2fmm\n", px, mm)

	fmt.Println()

	// Font setup for different DPIs.

	// Different DPI values for various devices.
	dpis := []float32{96, 120, 160, 240}
	fontSize := pxconv.Pt(12) // Base font size in points

	for _, dpi := range dpis {
		metric := pxconv.NewMetric(2, 1.5, dpi)
		px := metric.PtToPx(fontSize)
		fmt.Printf("DPI: %.0f, %.1fpt = %dpx\n", dpi, fontSize, px)
	}

	fmt.Println()

	// UI sizing using points.

	// UI elements.
	buttonFont := pxconv.Pt(14)
	labelFont := pxconv.Pt(12)
	headerFont := pxconv.Pt(18)

	fmt.Printf("Button font size: %.1fpt = %dpx\n", buttonFont, metric.PtToPx(buttonFont))
	fmt.Printf("Label font size: %.1fpt = %dpx\n", labelFont, metric.PtToPx(labelFont))
	fmt.Printf("Header font size: %.1fpt = %dpx\n", headerFont, metric.PtToPx(headerFont))
}
