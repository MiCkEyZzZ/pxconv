package pxconv

import (
	"math"

	"github.com/MiCkEyZzZ/pxconv/internal/consts"
	"github.com/MiCkEyZzZ/pxconv/internal/density"
)

// Dp represents device-independent pixels used for measuring distances on the screen.
type Dp float32

// Sp represents scale-independent pixels used for measuring font sizes.
type Sp float32

// Inch represents inches as a unit of measurement.
type Inch float32

// Mm represents millimeters as a unit of measurement.
type Mm float32

// Pt represents points as a typographic unit.
type Pt float32

// Metric is used to convert screen-independent units (dp, sp) to physical pixels (px).
type Metric struct {
	// PxPerDp - number of pixels per dp unit.
	PxPerDp float32
	// PxPerSp - number of pixels per sp unit.
	PxPerSp float32
	// Dpi - screen density in dots per inch.
	Dpi float32
}

// NewMetric creates a new Metric instance, validating input values.
// If any of the values are zero or negative, they are replaced with 1.
func NewMetric(pxPerDp, pxPerSp, dpi float32) Metric {
	if dpi <= 0 {
		dpi = consts.DefaultDpi
	}
	return Metric{
		PxPerDp: density.EnsurePositive(pxPerDp),
		PxPerSp: density.EnsurePositive(pxPerSp),
		Dpi:     dpi,
	}
}

// DpToPx converts a dp value to pixels, rounding to the nearest integer.
func (c Metric) DpToPx(value Dp) int {
	return int(math.Round(float64(density.EnsurePositive(c.PxPerDp)) * float64(value)))
}

// SpToPx converts an sp value to pixels, rounding to the nearest integer.
func (c Metric) SpToPx(value Sp) int {
	return int(math.Round(float64(density.EnsurePositive(c.PxPerSp)) * float64(value)))
}

// DpToSp converts a dp value to sp, using the current density values.
func (c Metric) DpToSp(value Dp) Sp {
	return Sp(float32(value) * density.EnsurePositive(c.PxPerDp) / density.EnsurePositive(c.PxPerSp))
}

// SpToDp converts an sp value to dp, using the current density values.
func (c Metric) SpToDp(value Sp) Dp {
	return Dp(float32(value) * density.EnsurePositive(c.PxPerSp) / density.EnsurePositive(c.PxPerDp))
}

// PxToDp converts a pixel value to dp.
func (c Metric) PxToDp(value int) Dp {
	return Dp(float32(value) / density.EnsurePositive(c.PxPerDp))
}

// PxToSp converts a pixel value to sp.
func (c Metric) PxToSp(value int) Sp {
	return Sp(float32(value) / density.EnsurePositive(c.PxPerSp))
}

// InchToPx converts inches to pixels using the current DPI.
// For example, with DPI = 96, InchToPx(1) returns 96.
func (c Metric) InchToPx(value Inch) int {
	return int(math.Round(float64(value) * float64(c.Dpi)))
}

// MmToPx converts millimeters to pixels using the current DPI.
// For example, with DPI = 96 and MmToPx(25.4), the result is 96.
func (c Metric) MmToPx(value Mm) int {
	return int(math.Round(float64(value) * float64(c.Dpi) / consts.MmPerInch))
}

// PxToInch converts pixels to inches using the current DPI.
// For example, with DPI = 96 and 96 pixels, the result is 1 inch.
func (c Metric) PxToInch(value int) Inch {
	return Inch(float32(value) / c.Dpi)
}

// PxToMm converts pixels to millimeters using the current DPI.
// For example, with DPI = 96 and 96 pixels, the result is 25.4 mm.
func (c Metric) PxToMm(value int) Mm {
	return Mm(float32(value) * consts.MmPerInch / c.Dpi)
}

// PtToPx converts points to pixels using the current DPI.
// For example, with DPI = 96, PtToPx(72) returns 96.
func (c Metric) PtToPx(value Pt) int {
	return int(math.Round(float64(value) * float64(c.Dpi) / consts.PointsPerInch))
}

// PxToPt converts pixels to points using the current DPI.
func (c Metric) PxToPt(value int) Pt {
	return Pt(float32(value) * consts.PointsPerInch / c.Dpi)
}

// GetDensity returns the current density values (PxPerDp and PxPerSp).
// Useful for inspecting or debugging density coefficients. DPI is not included.
func (c Metric) GetDensity() (float32, float32) {
	return c.PxPerDp, c.PxPerSp
}

// ScaleByDpi scales the current densities (PxPerDp, PxPerSp, and Dpi)
// by the given factor. Modifies the Metric instance in place.
//
// Deprecated: use Scale instead, which returns a new Metric without mutation.
func (c *Metric) ScaleByDpi(scale float32) {
	if scale <= 0 {
		scale = 1 // Prevents errors from invalid scaling factors.
	}
	c.PxPerDp *= scale
	c.PxPerSp *= scale
	c.Dpi *= scale
}
