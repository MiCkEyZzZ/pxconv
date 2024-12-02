package pxconv

import (
	"math"
)

// Dp represents device-independent units for measuring distances on the screen.
type Dp float32

// Sp represents device-independent units for measuring font sizes.
type Sp float32

// Metric is used to convert device-independent screen units (dp, sp) into pixels (px).
type Metric struct {
	// PxPerDp is the number of pixels per one dp.
	PxPerDp float32
	// PxPerSp is the number of pixels per one sp.
	PxPerSp float32
}

// NewMetric creates a new instance of Metric, validating the input values.
// If zero or negative values are provided, they will be replaced with 1.
func NewMetric(pxPerDp, pxPerSp float32) Metric {
	return Metric{
		PxPerDp: ensurePositive(pxPerDp),
		PxPerSp: ensurePositive(pxPerSp),
	}
}

// DpToPx converts a dp value into pixels (px), rounding to the nearest integer.
func (c Metric) DpToPx(value Dp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerDp)) * float64(value)))
}

// SpToPx converts an sp value into pixels (px), rounding to the nearest integer.
func (c Metric) SpToPx(value Sp) int {
	return int(math.Round(float64(ensurePositive(c.PxPerSp)) * float64(value)))
}

// DpToSp converts a dp value to sp using the respective densities for dp and sp.
func (c Metric) DpToSp(value Dp) Sp {
	return Sp(float32(value) * ensurePositive(c.PxPerDp) / ensurePositive(c.PxPerSp))
}

// SpToDp converts an sp value to dp using the respective densities for sp and dp.
func (c Metric) SpToDp(v Sp) Dp {
	return Dp(float32(v) * ensurePositive(c.PxPerSp) / ensurePositive(c.PxPerDp))
}

// PxToDp converts a pixel (px) value to dp.
func (c Metric) PxToDp(v int) Dp {
	return Dp(float32(v) / ensurePositive(c.PxPerDp))
}

// PxToSp converts a pixel (px) value to sp.
func (c Metric) PxToSp(v int) Sp {
	return Sp(float32(v) / ensurePositive(c.PxPerSp))
}

// GetDensity returns the current density values (PxPerDp and PxPerSp).
// This is useful for checking or debugging the current density ratios.
func (c Metric) GetDensity() (float32, float32) {
	return c.PxPerDp, c.PxPerSp
}

// ensurePositive returns a positive value. If the input value is 0 or less, it returns 1.
func ensurePositive(value float32) float32 {
	if value <= 0 {
		return 1
	}
	return value
}
