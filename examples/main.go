package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	metric := pxconv.NewMetric(2, 1.5, 96)

	// Конвертация между единицами.

	// Примеры конвертации.
	fmt.Printf("10dp = %dpx\n", metric.DpToPx(pxconv.Dp(10)))     // 10dp = 20px
	fmt.Printf("15sp = %dpx\n", metric.SpToPx(pxconv.Sp(15)))     // 15sp = 23px
	fmt.Printf("2inch = %dpx\n", metric.InchToPx(pxconv.Inch(2))) // 2inch = 192px
	fmt.Printf("50mm = %dpx\n", metric.MmToPx(pxconv.Mm(50)))     // 50mm = 189px
	fmt.Printf("10pt = %dpx\n", metric.PtToPx(pxconv.Pt(10)))     // 10pt = 13px

	fmt.Println()

	// Обратные преобразования.
	fmt.Printf("192px = %.2fdp\n", metric.PxToDp(192))     // 192px = 96.00dp
	fmt.Printf("128px = %.2fsp\n", metric.PxToSp(128))     // 128px = 85.33sp
	fmt.Printf("192px = %.2finch\n", metric.PxToInch(192)) // 192px = 2.00inch
	fmt.Printf("189px = %.2fmm\n", metric.PxToMm(189))     // 189px = 50.01mm
	fmt.Printf("10px = %.2fpt\n", metric.PxToMm(10))       // 10px = 2.65pt

	fmt.Println()

	// Масштабирование для устройства с нестандартным DPI.

	// Конвертация с текущими параметрами.
	fmt.Printf("10dp = %dpx\n", metric.DpToPx(pxconv.Dp(10)))

	// Масштабирование DPI.
	fmt.Printf("10dp после масштабирования (DPI = %.2f): %dpx\n", metric.Dpi, metric.DpToPx(pxconv.Dp(10)))

	fmt.Println()

	// Настройка шрифтов.

	// Размеры шрифта в sp и их отображение в px.
	fontSizes := []pxconv.Sp{12, 14, 16, 18, 24}
	for _, size := range fontSizes {
		fmt.Printf("Размер шрифта %.1fsp = %dpx\n", size, metric.SpToPx(size))
	}

	fmt.Println()

	// Подбор размеров элементов для разных устройств.

	// Эмуляция устройств с разным DPI
	devices := []struct {
		Name string
		Dpi  float32
	}{
		{"Стандартный дисплей", 96},
		{"HD-дисплей", 160},
		{"Дисплей 4К", 320},
	}

	baseMetric := pxconv.NewMetric(1, 1, pxconv.DefaultDpi)
	elementSize := pxconv.Dp(48) // Базовый размер элемента в dp

	for _, device := range devices {
		metric := pxconv.NewMetric(baseMetric.PxPerDp, baseMetric.PxPerSp, device.Dpi)
		pxSize := metric.DpToPx(elementSize)
		fmt.Printf("%s (%ddpi): %dpx\n", device.Name, int(device.Dpi), pxSize)
	}

	fmt.Println()

	// Использование в реальном приложении.

	buttonWidth := pxconv.Dp(200)
	buttonHeight := pxconv.Dp(50)

	fmt.Printf("Размеры кнопки: %dx%d пикселей\n",
		metric.DpToPx(buttonWidth),
		metric.DpToPx(buttonHeight),
	)

	fmt.Println()

	// Расчёт шрифтов.
	baseFontSize := pxconv.Sp(16)
	fmt.Printf("Базовый размер шрифта: %d пикселей\n", metric.SpToPx(baseFontSize))

	fmt.Println()

	// Конвертация пунктов в пиксели.

	// Конвертация пунктов в пиксели.
	pointValues := []pxconv.Pt{10, 12, 14, 16, 18, 24}
	for _, pt := range pointValues {
		px := metric.PtToPx(pt)
		fmt.Printf("%.1fpt = %dpx\n", pt, px)
	}

	// Обратная конвертация.
	pixels := []int{10, 72, 96, 144}
	for _, px := range pixels {
		pt := metric.PxToPt(px)
		fmt.Printf("%dpx = %.2fpt\n", px, pt)
	}

	fmt.Println()

	// Использование пунктов для текстовых размеров

	// Размеры текста в пунктах
	fontSizs := []pxconv.Pt{9, 10, 12, 14, 16, 18, 24}
	fmt.Println("Размеры шрифтов в пикселях:")
	for _, pt := range fontSizs {
		fmt.Printf("%.1fpt = %dpx\n", pt, metric.PtToPx(pt))
	}

	fmt.Println()

	// Конвертация между пунктами и другими единицами

	// Пример: Конвертация 72 пунктов в другие единицы
	pt := pxconv.Pt(72)

	px := metric.PtToPx(pt)
	inch := metric.PxToInch(px)
	mm := metric.PxToMm(px)

	fmt.Printf("%.1fpt = %dpx\n", pt, px)
	fmt.Printf("%dpx = %.2finch\n", px, inch)
	fmt.Printf("%dpx = %.2fmm\n", px, mm)

	fmt.Println()

	// Настройка шрифтов для разных DPI

	// Разные значения DPI для разных устройств
	dpis := []float32{96, 120, 160, 240}
	fontSize := pxconv.Pt(12) // Базовый размер шрифта в пунктах

	for _, dpi := range dpis {
		metric := pxconv.NewMetric(2, 1.5, dpi)
		px := metric.PtToPx(fontSize)
		fmt.Printf("DPI: %.0f, %.1fpt = %dpx\n", dpi, fontSize, px)
	}

	fmt.Println()

	// Подбор размеров UI с использованием пунктов

	// Элементы интерфейса
	buttonFont := pxconv.Pt(14)
	labelFont := pxconv.Pt(12)
	headerFont := pxconv.Pt(18)

	fmt.Printf("Размер шрифта кнопки: %.1fpt = %dpx\n", buttonFont, metric.PtToPx(buttonFont))
	fmt.Printf("Размер шрифта этикетки: %.1fpt = %dpx\n", labelFont, metric.PtToPx(labelFont))
	fmt.Printf("Размер шрифта заголовка: %.1fpt = %dpx\n", headerFont, metric.PtToPx(headerFont))
}
