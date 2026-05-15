package pxconv

import (
	"math"
	"testing"

	"pgregory.net/rapid"
)

// genPositiveFloat32 generates a positive float32 in range (0, 1000].
func genPositiveFloat32(t *rapid.T, label string) float32 {
	return rapid.Float32Range(0.01, 1000.0).Draw(t, label)
}

// genPositiveDpi generates a realistic DPI value in range [72, 600].
func genPositiveDpi(t *rapid.T) float32 {
	return rapid.Float32Range(72.0, 600.0).Draw(t, "dpi")
}

// TestPropDpToPxRoundtrip checks that DpToPx → PxToDp is identity
// for integer dp values (rounding makes non-integer inputs lossy).
func TestPropDpToPxRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		pxPerDp := genPositiveFloat32(t, "pxPerDp")
		dp := rapid.Float32Range(0, 10000).Draw(t, "dp")

		m := NewMetric(pxPerDp, 1.0, 96)

		px := m.DpToPx(Dp(dp))
		recovered := m.PxToDp(px)

		// tolerance in dp space due to integer rounding
		tolerance := 1.0 / float64(pxPerDp)

		if math.Abs(float64(recovered)-float64(dp)) > tolerance+1e-4 {
			t.Fatalf("Dp roundtrip failed: dp=%v px=%v recovered=%v", dp, px, recovered)
		}
	})
}

// TestPropSpToPxRoundtrip checks that SpToPx → PxToSp is identity
// within rounding tolerance.
func TestPropSpToPxRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		pxPerSp := genPositiveFloat32(t, "pxPerSp")
		sp := rapid.Float32Range(0, 10000).Draw(t, "sp")

		m := NewMetric(1.0, pxPerSp, 96)

		px := m.SpToPx(Sp(sp))
		recovered := m.PxToSp(px)

		tolerance := 1.0 / float64(pxPerSp)

		if math.Abs(float64(recovered)-float64(sp)) > tolerance+1e-4 {
			t.Fatalf("Sp roundtrip failed: sp=%v px=%v recovered=%v", sp, px, recovered)
		}
	})
}

// TestPropInchToPxEqualsDpi checks that InchToPx(1) == round(Dpi)
// for any positive Dpi.
func TestPropInchToPxEqualsDpi(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dpi := genPositiveDpi(t)
		m := NewMetric(1.0, 1.0, dpi)

		got := m.InchToPx(Inch(1))
		want := int(math.Round(float64(dpi)))

		if got != want {
			t.Fatalf("InchToPx(1) = %v; want round(Dpi)=%v (Dpi=%v)", got, want, dpi)
		}
	})
}

// TestPropMmPerInchEqualsDpi checks that MmToPx(25.4) == InchToPx(1)
// for any positive Dpi (physical invariant: 25.4 mm == 1 inch).
func TestPropMmPerInchEqualsDpi(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dpi := genPositiveDpi(t)

		m := NewMetric(1.0, 1.0, dpi)

		// базовое значение 1 inch в mm
		mm := Mm(25.4)
		inch := Inch(1)

		mmPx := m.MmToPx(mm)
		inchPx := m.InchToPx(inch)

		// сравнение должно быть через tolerance из-за double rounding
		if math.Abs(float64(mmPx-inchPx)) > 1 {
			t.Fatalf(
				"mmPx=%v inchPx=%v dpi=%v",
				mmPx, inchPx, dpi,
			)
		}
	})
}

func TestPropInchMmConsistency(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dpi := genPositiveDpi(t)
		m := NewMetric(1.0, 1.0, dpi)

		inchPx := m.InchToPx(1)
		mmPx := m.MmToPx(25.4)

		// must be same after rounding (allow 1px drift)
		if math.Abs(float64(inchPx-mmPx)) > 1 {
			t.Fatalf("physical mismatch: inchPx=%v mmPx=%v dpi=%v", inchPx, mmPx, dpi)
		}
	})
}

func TestPropPtEqualsInch(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dpi := genPositiveDpi(t)
		m := NewMetric(1.0, 1.0, dpi)

		ptPx := m.PtToPx(72)
		inchPx := m.InchToPx(1)

		if math.Abs(float64(ptPx-inchPx)) > 1 {
			t.Fatalf("pt vs inch mismatch: ptPx=%v inchPx=%v dpi=%v", ptPx, inchPx, dpi)
		}
	})
}

// TestPropPtPerInchEqualsDpi checks that PtToPx(72) == InchToPx(1)
// for any positive Dpi (typographic invariant: 72 pt == 1 inch).
func TestPropPtPerInchEqualsDpi(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dpi := genPositiveDpi(t)
		m := NewMetric(1.0, 1.0, dpi)

		ptPx := m.PtToPx(Pt(72))
		inchPx := m.InchToPx(Inch(1))

		if ptPx != inchPx {
			t.Fatalf("PtToPx(72)=%v != InchToPx(1)=%v at Dpi=%v", ptPx, inchPx, dpi)
		}
	})
}

// TestPropDpToSpIdentity checks that DpToSp → SpToDp is identity
// when PxPerDp == PxPerSp.
func TestPropDpToSpIdentity(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		density := genPositiveFloat32(t, "density")
		dp := rapid.Float32Range(0, 10000).Draw(t, "dp")

		m := NewMetric(density, density, 96)

		sp := m.DpToSp(Dp(dp))
		recovered := m.SpToDp(sp)

		if math.Abs(float64(recovered)-float64(dp)) > 1e-3 {
			t.Fatalf("Dp<->Sp identity failed: dp=%v sp=%v recovered=%v", dp, sp, recovered)
		}
	})
}

// TestPropDpToSpRoundtrip checks that DpToSp → SpToDp roundtrip
// holds for arbitrary different densities.
func TestPropDpToSpRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		pxPerDp := genPositiveFloat32(t, "pxPerDp")
		pxPerSp := genPositiveFloat32(t, "pxPerSp")
		dp := rapid.Float32Range(0, 10000).Draw(t, "dp")

		m := NewMetric(pxPerDp, pxPerSp, 96)

		sp := m.DpToSp(Dp(dp))
		recovered := m.SpToDp(sp)

		if math.Abs(float64(recovered)-float64(dp)) > 1e-3 {
			t.Fatalf("Dp<->Sp roundtrip failed: dp=%v sp=%v recovered=%v", dp, sp, recovered)
		}
	})
}
