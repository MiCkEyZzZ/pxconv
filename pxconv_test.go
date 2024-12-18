package pxconv

import (
	"sync"
	"testing"
)

// TestDpToPx проверяет, правильно ли округляется значение dp в пиксели.
func TestDpToPx(t *testing.T) {
	m := Metric{PxPerDp: 2}
	tests := []struct {
		dp       Dp
		expected int
	}{
		{1, 2},
		{0, 0},
		{2, 4},
	}

	for _, test := range tests {
		res := m.DpToPx(test.dp)
		if res != test.expected {
			t.Errorf("DpToPx(%v) = %v; expected %v", test.dp, res, test.expected)
		}
	}
}

// TestSpToPx проверяет, правильно ли округляется значение sp в пиксели.
func TestSpToPx(t *testing.T) {
	m := Metric{PxPerSp: 4}
	tests := []struct {
		sp       Sp
		expected int
	}{
		{1, 4},
		{0, 0},
		{2, 8},
	}

	for _, test := range tests {
		res := m.SpToPx(test.sp)
		if res != test.expected {
			t.Errorf("SpToPx(%v) = %v; expected %v", test.sp, res, test.expected)
		}
	}
}

// TestDpToSp проверяет корректность преобразования из dp в sp.
func TestDpToSp(t *testing.T) {
	m := Metric{PxPerDp: 2, PxPerSp: 4}
	tests := []struct {
		dp       Dp
		expected Sp
	}{
		{1, 0.5},
		{2, 1},
		{0, 0},
	}

	for _, test := range tests {
		res := m.DpToSp(test.dp)
		if res != test.expected {
			t.Errorf("DpToSp(%v) = %v; expected %v", test.dp, res, test.expected)
		}
	}
}

// TestSpToDp проверяет корректность преобразования из sp в dp.
func TestSpToDp(t *testing.T) {
	m := Metric{PxPerDp: 2, PxPerSp: 4}
	tests := []struct {
		sp       Sp
		expected Dp
	}{
		{0.5, 1},
		{1, 2},
		{0, 0},
	}

	for _, test := range tests {
		res := m.SpToDp(test.sp)
		if res != test.expected {
			t.Errorf("SpToDp(%v) = %v; expected %v", test.sp, res, test.expected)
		}
	}
}

// TestPxToDp проверяет корректность преобразования из px в dp.
func TestPxToDp(t *testing.T) {
	m := Metric{PxPerDp: 2}
	tests := []struct {
		px       int
		expected Dp
	}{
		{2, 1},
		{0, 0},
		{4, 2},
	}

	for _, test := range tests {
		res := m.PxToDp(test.px)
		if res != test.expected {
			t.Errorf("PxToDp(%v) = %v; expected %v", test.px, res, test.expected)
		}
	}
}

// TestPxToSp проверяет корректность преобразования из px в sp.
func TestPxToSp(t *testing.T) {
	m := Metric{PxPerSp: 4}
	tests := []struct {
		px       int
		expected Sp
	}{
		{4, 1},
		{0, 0},
		{8, 2},
	}

	for _, test := range tests {
		res := m.PxToSp(test.px)
		if res != test.expected {
			t.Errorf("PxToSp(%v) = %v; expected %v", test.px, res, test.expected)
		}
	}
}

// TestEnsurePositiveDefault проверяет, что значения по умолчанию применяются корректно.
func TestEnsurePositiveDefault(t *testing.T) {
	m := Metric{PxPerDp: 0, PxPerSp: -5}
	if res := m.DpToPx(2); res != 2 {
		t.Errorf("DpToPx(2) with PxPerDp=0 = %v; expected 2", res)
	}
	if res := m.SpToPx(2); res != 2 {
		t.Errorf("SpToPx(2) with PxPerSp=5 = %v; expected 2", res)
	}
}

// TestRounding проверяет корректность округления значений.
func TestRounding(t *testing.T) {
	m := Metric{PxPerDp: 2, PxPerSp: 3}
	tests := []struct {
		in       float32
		expected int
	}{
		{1.4, 3},
		{1.5, 3},
		{1.6, 3},
	}

	for _, test := range tests {
		res := m.DpToPx(Dp(test.in))
		if res != test.expected {
			t.Errorf("DpToPx(%v) = %v; expected %v", test.in, res, test.expected)
		}
	}
}

// TestExtremeValues проверяет работу с экстремальными значениями.
func TestExtremeValues(t *testing.T) {
	m := Metric{PxPerDp: 1000, PxPerSp: 0.001}
	tests := []struct {
		dp       Dp
		expected int
	}{
		{dp: Dp(1000000), expected: 1000000000}, // 1000000 * 1000
		{dp: Dp(0.000001), expected: 0},         // 0.000001 * 1000 округляется до 0
	}

	for _, test := range tests {
		res := m.DpToPx(test.dp)
		if res != test.expected {
			t.Errorf("DpToPx(%v) = %v; expected %v", test.dp, res, test.expected)
		}
	}
}

// TestConsistency проверяет согласованность преобразований DpToSp и SpToDp.
func TestConsistency(t *testing.T) {
	m := Metric{PxPerDp: 2, PxPerSp: 4}
	dp := Dp(2)
	sp := m.DpToSp(dp)
	resDp := m.SpToDp(sp)

	if resDp != dp {
		t.Errorf("SpToDp(DpToSp(%v)) = %v; expected %v", dp, resDp, dp)
	}
}

// TestBoundaryRounding проверяет округление значений на гарнице.
func TestBoundaryRounding(t *testing.T) {
	m := Metric{PxPerDp: 1.0}
	tests := []struct {
		dp       Dp
		expected int
	}{
		{0.49999, 0},
		{0.50001, 1},
	}

	for _, test := range tests {
		res := m.DpToPx(test.dp)
		if res != test.expected {
			t.Errorf("DpToPx(%v) = %v; expected %v", test.dp, res, test.expected)
		}
	}
}

// TestNegativeValues проверяет, что при передаче отрицательных значений в NewMetric, они заменяются на 1.
func TestNegativeValues(t *testing.T) {
	m := NewMetric(-10, -5, 96)

	if m.PxPerDp != 1 {
		t.Errorf("NewMetric(-10, -5) PxPerDp = %v; expected 1", m.PxPerDp)
	}
	if m.PxPerSp != 1 {
		t.Errorf("NewMetric(-10, -5) PxPerSp = %v; expected 1", m.PxPerSp)
	}
}

// TestZeroValues проверяет, что нулевые значения заменяются на 1 в NewMetric.
func TestZeroValues(t *testing.T) {
	m := NewMetric(0, 0, 96)

	if m.PxPerDp != 1 {
		t.Errorf("NewMetric(0, 0) PxPerDp = %v; expected 1", m.PxPerDp)
	}
	if m.PxPerSp != 1 {
		t.Errorf("NewMetric(0, 0) PxPerSp = %v; expected 1", m.PxPerSp)
	}
}

// TestConstructorWithValidValues проверяет, что корректные значения не изменяются в NewMetric.
func TestConstructorWithValidValues(t *testing.T) {
	m := NewMetric(2, 3, 96)

	if m.PxPerDp != 2 {
		t.Errorf("NewMetric(2, 3) PxPerDp = %v; expected 2", m.PxPerDp)
	}
	if m.PxPerSp != 3 {
		t.Errorf("NewMetric(2, 3) PxPerSp = %v; expected 3", m.PxPerSp)
	}
}

// TestInchToPx проверяет преобразование дюймов в пиксели.
func TestInchToPx(t *testing.T) {
	m := Metric{Dpi: 96}
	tests := []struct {
		inch     Inch
		expected int
	}{
		{1, 96},
		{0, 0},
		{2.5, 240},
	}

	for _, test := range tests {
		res := m.InchToPx(test.inch)
		if res != test.expected {
			t.Errorf("InchToPx(%v) = %v; expected %v", test.inch, res, test.expected)
		}
	}
}

// TestMmToPx проверяет преобразование миллиметров в пиксели.
func TestMmToPx(t *testing.T) {
	m := Metric{Dpi: 96}
	tests := []struct {
		mm       Mm
		expected int
	}{
		{MmPerInch, 96}, // 25.4 мм = 1 дюйм = 96 пикселей
		{0, 0},
		{50.8, 192}, // 50.8 мм = 2 дюйма
	}

	for _, test := range tests {
		res := m.MmToPx(test.mm)
		if res != test.expected {
			t.Errorf("MmToPx(%v) = %v; expected %v", test.mm, res, test.expected)
		}
	}
}

// TestPxToInch проверяет преобразование пикселей в дюймы.
func TestPxToInch(t *testing.T) {
	m := Metric{Dpi: 96}
	tests := []struct {
		px       int
		expected Inch
	}{
		{96, 1},
		{0, 0},
		{192, 2},
	}

	for _, test := range tests {
		res := m.PxToInch(test.px)
		if res != test.expected {
			t.Errorf("PxToInch(%v) = %v; expected %v", test.px, res, test.expected)
		}
	}
}

// TestPxToMm проверяет преобразование пикселей в миллиметры.
func TestPxToMm(t *testing.T) {
	m := Metric{Dpi: 96}
	tests := []struct {
		px       int
		expected Mm
	}{
		{96, MmPerInch}, // 96 пикселей = 25.4 мм
		{0, 0},
		{192, 50.8}, // 192 пикселя = 50.8 мм
	}

	for _, test := range tests {
		res := m.PxToMm(test.px)
		if res != test.expected {
			t.Errorf("PxToMm(%v) = %v; expected %v", test.px, res, test.expected)
		}
	}
}

// TestDpToPxMaxValue стресс-тесты.
func TestDpToPxMaxValue(t *testing.T) {
	m := Metric{PxPerDp: 2}
	maxValue := Dp(1e6)
	expected := 2 * int(maxValue)
	result := m.DpToPx(maxValue)

	if result != expected {
		t.Errorf("DpToPx(%v) = %v; expected %v", maxValue, result, expected)
	}
}

// TestPxToDpMaxValue стресс-тесты.
func TestPxToDpMaxValue(t *testing.T) {
	m := Metric{PxPerDp: 2}
	maxValue := 1e6
	expected := Dp(maxValue / 2)
	result := m.PxToDp(int(maxValue))

	if result != expected {
		t.Errorf("PxToDp(%v) = %v; expected %v", maxValue, result, expected)
	}
}

// TestDpToPxConcurrency тесты на конкурентность.
func TestDpToPxConcurrency(t *testing.T) {
	m := Metric{PxPerDp: 2}
	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(dp Dp) {
			defer wg.Done()
			res := m.DpToPx(dp)
			if res != int(dp)*2 {
				t.Errorf("Concurrency error: DpToPx(%v) = %v; expected %v", dp, res, int(dp)*2)
			}
		}(Dp(i))
	}
	wg.Wait()
}

// TestDpToPxLoad нагрузочное тестирование.
func TestDpToPxLoad(t *testing.T) {
	m := Metric{PxPerDp: 2}
	iterations := 1e6

	for i := 0; i < int(iterations); i++ {
		res := m.DpToPx(Dp(i))
		if res != i*2 {
			t.Errorf("Load test error: DpToPx(%v) = %v; expected %v", i, res, i*2)
		}
	}
}

// TestDpToPxToDpConsistency на обратимость.
func TestDpToPxToDpConsistency(t *testing.T) {
	m := Metric{PxPerDp: 2}
	tests := []Dp{0, 1, 2, 100, 1e6}

	for _, dp := range tests {
		px := m.DpToPx(dp)
		resDp := m.PxToDp(px)
		if resDp != dp {
			t.Errorf("Inconsistency error: PxToDp(DpToPx(%v)) = %v; expected %v", dp, resDp, dp)
		}
	}
}

// TestDpToPxRounding на округление.
func TestDpToPxRounding(t *testing.T) {
	m := Metric{PxPerDp: 2}
	tests := []struct {
		dp       Dp
		expected int
	}{
		{0.4, 1},
		{0.5, 1},
		{0.6, 1},
		{1.4, 3},
		{1.5, 3},
	}

	for _, test := range tests {
		res := m.DpToPx(test.dp)
		if res != test.expected {
			t.Errorf("DpToPx(%v) = %v; expected %v", test.dp, res, test.expected)
		}
	}
}
