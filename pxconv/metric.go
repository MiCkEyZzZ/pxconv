package pxconv

import (
	"math"
)

// Dp — единицы, независимые от устройства, для измерения расстояний на экране.
type Dp float32

// Sp — единицы, независимые от устройства, для измерения шрифтов.
type Sp float32

// Metric используется для конвертации независимых экранных единиц (dp, sp) в пиксели (px).
type Metric struct {
	// PxPerDp - количество пикселей на один dp.
	PxPerDp float32
	// PxPerSp - количество пикселей на один sp.
	PxPerSp float32
}

// NewMetric создаёт новый экземпляр Metric, проверяя входные значения на корректность.
// Если переданы нулевые или отрицательные значения, они будут заменены на 1.
func NewMetric(pxPerDp, pxPerSp float32) Metric {
	return Metric{
		PxPerDp: ensurePositive(pxPerDp),
		PxPerSp: ensurePositive(pxPerSp),
	}
}

// DpToPx м-д конвертирует значение dp в пиксели (px), округляя до ближайшего целого числа.
func (c Metric) DpToPx(value Dp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerDp)) * float64(value)))
}

// SpToPx м-д конвертирует значение sp в пиксели (px), округляя до ближайшего целого числа.
func (c Metric) SpToPx(value Sp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerSp)) * float64(value)))
}

// DpToSp м-д конвертирует значение dp в sp, используя плотности для dp и sp.
func (c Metric) DpToSp(value Dp) Sp {
	return Sp(float32(value) * ensurePositive(c.PxPerDp) / ensurePositive(c.PxPerSp))
}

// SpToDp м-д конвертирует значение sp в dp, используя плотности для dp и sp.
func (c Metric) SpToDp(v Sp) Dp {
	return Dp(float32(v) * ensurePositive(c.PxPerSp) / ensurePositive(c.PxPerDp))
}

// PxToDp м-д конвертирует значение пикселей (px) в dp.
func (c Metric) PxToDp(v int) Dp {
	return Dp(float32(v) / ensurePositive(c.PxPerDp))
}

// PxToSp м-д конвертирует значение пикселей (px) в sp.
func (c Metric) PxToSp(v int) Sp {
	return Sp(float32(v) / ensurePositive(c.PxPerSp))
}

// GetDensity возвращает текущие значения плотности (PxPerDp и PxPerSp).
// Используется для проверки или отладки текущих коэффициентов плотности.
func (c Metric) GetDensity() (float32, float32) {
	return c.PxPerDp, c.PxPerSp
}

// ensurePositive ф-я возвращает положительное значение. Если входное
// значение 0 или меньше, возвращает 1.
func ensurePositive(value float32) float32 {
	if value <= 0 {
		return 1
	}
	return value
}
