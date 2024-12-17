package pxconv

import "testing"

// BenchmarkDpToPx бенчмарк для метода DpToPx.
func BenchmarkDpToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToPx(Dp(10))
	}
}

// BenchmarkSpToPx бенчмарк для метода SpToPx.
func BenchmarkSpToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.SpToPx(Sp(15))
	}
}

// BenchmarkPxToDp бенчмарк для метода PxToDp.
func BenchmarkPxToDp(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToDp(50)
	}
}

// BenchmarkDpToSp бенчмарк для метода DpToSp.
func BenchmarkDpToSp(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToSp(Dp(10))
	}
}

// BenchmarkInchToPx бенчмарк для метода InchToPx.
func BenchmarkInchToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.InchToPx(Inch(1.0))
	}
}

// BenchmarkMmToPx бенчмарк для метода MmToPx.
func BenchmarkMmToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.MmToPx(Mm(25.4))
	}
}

// BenchmarkPxToInch бенчмарк для метода PxToInch.
func BenchmarkPxToInch(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToInch(96)
	}
}

// BenchmarkPxToMm бенчмарк для метода PxToMm.
func BenchmarkPxToMm(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToMm(96)
	}
}
