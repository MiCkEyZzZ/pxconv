package pxconv

import (
	"math"
)

const (
	DefaultDpi = 96   // Стандартное значение DPI для большинства дисплеев.
	MmPerInch  = 25.4 // Количество миллиметров в одном дюйме, константа для вычислений.
)

// Dp — единицы, независимые от устройства, для измерения расстояний на экране.
type Dp float32

// Sp — единицы, независимые от устройства, для измерения шрифтов.
type Sp float32

// Inch — единица измерения для дюймов
type Inch float32

// Mm — единица измерения для миллиметров.
type Mm float32

// Metric используется для конвертации независимых экранных единиц (dp, sp) в пиксели (px).
type Metric struct {
	// PxPerDp - количество пикселей на один dp.
	PxPerDp float32
	// PxPerSp - количество пикселей на один sp.
	PxPerSp float32
	// Dpi - количество пикселей на дюйм.
	Dpi float32
}

// NewMetric создаёт новый экземпляр Metric, проверяя входные значения на корректность.
// Если переданы нулевые или отрицательные значения, они будут заменены на 1.
func NewMetric(pxPerDp, pxPerSp, dpi float32) Metric {
	if dpi <= 0 {
		dpi = DefaultDpi
	}
	return Metric{
		PxPerDp: ensurePositive(pxPerDp),
		PxPerSp: ensurePositive(pxPerSp),
		Dpi:     dpi,
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
func (c Metric) SpToDp(value Sp) Dp {
	return Dp(float32(value) * ensurePositive(c.PxPerSp) / ensurePositive(c.PxPerDp))
}

// PxToDp м-д конвертирует значение пикселей (px) в dp.
func (c Metric) PxToDp(value int) Dp {
	return Dp(float32(value) / ensurePositive(c.PxPerDp))
}

// PxToSp м-д конвертирует значение пикселей (px) в sp.
func (c Metric) PxToSp(value int) Sp {
	return Sp(float32(value) / ensurePositive(c.PxPerSp))
}

// InchToPx м-д конвертирует значение инчи (inch) в пиксели (px), используя DPI
// Например, при DPI = 96, InchToPx(1) вернёт 96.
func (c Metric) InchToPx(value Inch) int {
	return int(math.Round(float64(value) * float64(c.Dpi)))
}

// MmToPx м-д конвертирует значение миллиметры (mm) в пиксели (px), используя DPI.
// Например, при DPI = 96 и MmToPx(25.4), результат будет равен 96.
func (c Metric) MmToPx(value Mm) int {
	return int(math.Round(float64(value) * float64(c.Dpi) / MmPerInch))
}

// PxToInch м-д конвертирует значение пикселей (px) в инчи (inch), используя DPI.
// Например, при DPI = 96 и значении 96 пикселей, результат будет равен 1 дюйму.
func (c Metric) PxToInch(value int) Inch {
	return Inch(float32(value) / c.Dpi)
}

// PxToMm м-д конвертирует значение пикселей (px) в миллиметры (mm), используя DPI
// Например, при DPI = 96 и значении 96 пикселей, результат будет равен 25.4 мм.
func (c Metric) PxToMm(value int) Mm {
	return Mm(float32(value) * MmPerInch / c.Dpi)
}

// GetDensity м-д возвращает текущие значения плотности (PxPerDp и PxPerSp).
// Используется для проверки или отладки текущих коэффициентов плотности.
// Не включает DPI.
func (c Metric) GetDensity() (float32, float32) {
	return c.PxPerDp, c.PxPerSp
}

// ScaleByDpi масштабирует текущие плотности (PxPerDp, PxPerSp и Dpi)
// на указанный коэффициент. Изменяет состояние объекта Metric.
func (c *Metric) ScaleByDpi(scale float32) {
	if scale <= 0 {
		scale = 1 // Предотвращение ошибок при некорректном масштабировании.
	}
	c.PxPerDp *= scale
	c.PxPerSp *= scale
	c.Dpi *= scale
}

// ensurePositive ф-я возвращает положительное значение. Если входное
// значение 0 или меньше, возвращает 1.
func ensurePositive(value float32) float32 {
	if value <= 0 {
		return 1
	}
	return value
}
