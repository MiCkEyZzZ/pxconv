package pxconv

import "math"

// Dp представляет независимые от устройства пиксели.
type Dp float32

// Sp представляет независимые от устройства единицы для шрифтов.
type Sp float32

// Metric конвертирует значения в зависимости от устройства, такие как dp и
// sp, в пиксели (px). Используется для расчета размеров элементов интерфейса
// с учетом плотности экрана.
type Metric struct {
	// PxPerDp - количество пикселей на один dp.
	PxPerDp float32
	// PxPerSp - количество пикселей на один sp.
	PxPerSp float32
}

// DpToPx конвертирует значение dp в пиксели (px), округляя до ближайшего целого числа.
func (c Metric) DpToPx(value Dp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerDp)) * float64(value)))
}

// SpToPx конвертирует значение sp в пиксели (px), округляя до ближайшего целого числа.
func (c Metric) SpToPx(value Sp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerSp)) * float64(value)))
}

// DpToSp м-д конвертирует значение dp в sp, используя плотности для dp и sp.
func (c Metric) DpToSp(value Dp) Sp {
	return Sp(float32(value) * c.ensurePositivePxPerDp() / c.ensurePositivePxPerSp())
}

// SpToDp м-д конвертирует значение sp в dp, используя плотности для dp и sp.
func (c Metric) SpToDp(v Sp) Dp {
	return Dp(float32(v) * c.ensurePositivePxPerSp() / c.ensurePositivePxPerDp())
}

// PxToDp м-д конвертирует значение пикселей (px) в dp.
func (c Metric) PxToDp(v int) Dp {
	return Dp(float32(v) / c.ensurePositivePxPerDp())
}

// PxToSp м-д конвертирует значение пикселей (px) в sp.
func (c Metric) PxToSp(v int) Sp {
	return Sp(float32(v) / c.ensurePositivePxPerSp())
}

// ensurePositivePxPerDp возвращает положительное значение PxPerDp,
// если значение равно 0 или отрицательное, используется значение по
// умолчанию 1.
func (c Metric) ensurePositivePxPerDp() float32 {
	return ensurePositive(c.PxPerDp)
}

// ensurePositivePxPerSp возвращает положительное значение PxPerSp,
// если значение равно 0 или отрицательное, используется значение по
// умолчанию 1.
func (c Metric) ensurePositivePxPerSp() float32 {
	return ensurePositive(c.PxPerSp)
}

// ensurePositive возвращает положительное значение. Если входное
// значение 0 или меньше, возвращает 1.
func ensurePositive(value float32) float32 {
	if value <= 0 {
		return 1
	}
	return value
}
