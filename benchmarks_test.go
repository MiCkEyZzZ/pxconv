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

// BenchmarkPxToMm бенчмарк для метода PxToMm
func BenchmarkPxToMm(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToMm(96)
	}
}

// BenchmarkDpToPxLargeValue бенчмарк для больших значений.
func BenchmarkDpToPxLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToPx(Dp(1_000_000))
	}
}

// BenchmarkMethodsInSequence бенчмарк последовательного вызова метода.
func BenchmarkMethodsInSequence(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = metric.DpToPx(Dp(10))
		_ = metric.PxToDp(50)
		_ = metric.SpToPx(Sp(15))
	}
}

// BenchmarkDpToPxParallel бенчмарк тестирует многопоточность.
func BenchmarkDpToPxParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.DpToPx(Dp(10))
		}
	})
}

// BenchmarkPtToPx бенчмарк для метода PtToPx.
func BenchmarkPtToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PtToPx(Pt(12))
	}
}

// BenchmarkPxToPt бенчмарк для метода PxToPt.
func BenchmarkPxToPt(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToPt(96)
	}
}

// BenchmarkPtToPxLargeValue бенчмарк для метода PxToPx с большим значением.
func BenchmarkPtToPxLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PtToPx(Pt(1_000))
	}
}

// BenchmarkPxToPtLargeValue бенчмарк для метода PxToPt с большими значениями.
func BenchmarkPxToPtLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToPt(10_000)
	}
}

// BenchmarkPtToPxParallel бенчмарк тестирует многопоточность метода PtToPx.
func BenchmarkPtToPxParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.PtToPx(Pt(12))
		}
	})
}

// BenchmarkPxToPtParallel бенчмарк тестирует многопоточность метода PxToPt.
func BenchmarkPxToPtParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.PxToPt(96)
		}
	})
}
