package pxconv

import "testing"

// BenchmarkDpToPx benchmarks the DpToPx method.
func BenchmarkDpToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToPx(Dp(10))
	}
}

// BenchmarkSpToPx benchmarks the SpToPx method.
func BenchmarkSpToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.SpToPx(Sp(15))
	}
}

// BenchmarkPxToDp benchmarks the PxToDp method.
func BenchmarkPxToDp(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToDp(50)
	}
}

// BenchmarkDpToSp benchmarks the DpToSp method.
func BenchmarkDpToSp(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToSp(Dp(10))
	}
}

// BenchmarkInchToPx benchmarks the InchToPx method.
func BenchmarkInchToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.InchToPx(Inch(1.0))
	}
}

// BenchmarkMmToPx benchmarks the MmToPx method.
func BenchmarkMmToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.MmToPx(Mm(25.4))
	}
}

// BenchmarkPxToInch benchmarks the PxToInch method.
func BenchmarkPxToInch(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToInch(96)
	}
}

// BenchmarkPxToMm benchmarks the PxToMm method.
func BenchmarkPxToMm(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToMm(96)
	}
}

// BenchmarkDpToPxLargeValue benchmarks DpToPx with a large value.
func BenchmarkDpToPxLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.DpToPx(Dp(1_000_000))
	}
}

// BenchmarkMethodsInSequence benchmarks a sequence of method calls.
func BenchmarkMethodsInSequence(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = metric.DpToPx(Dp(10))
		_ = metric.PxToDp(50)
		_ = metric.SpToPx(Sp(15))
	}
}

// BenchmarkDpToPxParallel benchmarks DpToPx in a parallel context.
func BenchmarkDpToPxParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.DpToPx(Dp(10))
		}
	})
}

// BenchmarkPtToPx benchmarks the PtToPx method.
func BenchmarkPtToPx(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PtToPx(Pt(12))
	}
}

// BenchmarkPxToPt benchmarks the PxToPt method.
func BenchmarkPxToPt(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToPt(96)
	}
}

// BenchmarkPtToPxLargeValue benchmarks PtToPx with a large value.
func BenchmarkPtToPxLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PtToPx(Pt(1_000))
	}
}

// BenchmarkPxToPtLargeValue benchmarks PxToPt with a large value.
func BenchmarkPxToPtLargeValue(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	for i := 0; i < b.N; i++ {
		metric.PxToPt(10_000)
	}
}

// BenchmarkPtToPxParallel benchmarks PtToPx in a parallel context.
func BenchmarkPtToPxParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.PtToPx(Pt(12))
		}
	})
}

// BenchmarkPxToPtParallel benchmarks PxToPt in a parallel context.
func BenchmarkPxToPtParallel(b *testing.B) {
	metric := NewMetric(2.0, 2.0, 96)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metric.PxToPt(96)
		}
	})
}
