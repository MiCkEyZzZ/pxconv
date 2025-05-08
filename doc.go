// Package pxconv provides tools for working with user interface measurement units:
// density-independent pixels (dp), scale-independent pixels (sp), points (pt),
// and their conversion to physical pixels (px) based on screen density.
//
// # Key Concepts
//
// Dp (Density-independent Pixels) are device-independent units used for specifying
// UI element sizes (e.g., buttons or paddings). They ensure consistent sizing
// across devices with varying screen densities.
//
// Sp (Scale-independent Pixels) are similar to Dp but additionally respect
// the user's font size preferences. They are used for text sizing.
//
// Pt (Points) are traditional typographic units. There are 72 points in an inch.
// They are used for font sizes and other graphical design elements.
//
// Pixels (px) are the physical dots on the screen. Their quantity per length unit
// depends on screen density (DPI – dots per inch).
//
// # Metric Structure
//
// The core `Metric` struct represents parameters used for conversion between
// Dp, Sp, Pt, and pixels. It accounts for screen density and user preferences.
//
// Fields of the struct:
//   - PxPerDp: Number of pixels per Dp.
//   - PxPerSp: Number of pixels per Sp.
//   - Dpi: Screen density in dots per inch.
//
// # Creating a Metric Instance
//
// A `Metric` instance can be created using the `NewMetric` constructor,
// which takes density values for Dp, Sp, and DPI. If invalid values (zero or negative)
// are passed, they are replaced with defaults.
//
// Example:
//
//	metric := pxconv.NewMetric(2.0, 1.5, 96) // Densities: 2 px/dp, 1.5 px/sp, DPI 96.
//
// Alternatively, you can manually create a Metric instance, but be aware that
// invalid values may cause calculation errors.
//
// Example:
//
//	metric := pxconv.Metric{PxPerDp: 2.0, PxPerSp: 1.5, Dpi: 96}
//
// # Unit Conversion
//
// The `Metric` methods allow converting between different measurement units:
//
//   - DpToPx: Converts Dp to pixels (px), rounding to the nearest integer.
//   - SpToPx: Converts Sp to pixels (px), rounding to the nearest integer.
//   - PtToPx: Converts points (pt) to pixels (px).
//   - PxToDp: Converts pixels (px) to Dp.
//   - PxToSp: Converts pixels (px) to Sp.
//   - PxToPt: Converts pixels (px) to points (pt).
//   - DpToSp: Converts Dp to Sp.
//   - SpToDp: Converts Sp to Dp.
//   - InchToPx: Converts inches to pixels (px).
//   - MmToPx: Converts millimeters to pixels (px).
//   - PxToInch: Converts pixels (px) to inches.
//   - PxToMm: Converts pixels (px) to millimeters.
//
// Example:
//
//	metric := pxconv.NewMetric(2.0, 1.5, 96) // Screen densities
//	pxFromDp := metric.DpToPx(10)           // Result: 20 px
//	pxFromSp := metric.SpToPx(10)           // Result: 15 px
//	pxFromPt := metric.PtToPx(12)           // Result: 16 px (at DPI 96)
//	dpFromPx := metric.PxToDp(20)           // Result: 10 dp
//	spFromPx := metric.PxToSp(15)           // Result: 10 sp
//	ptFromPx := metric.PxToPt(16)           // Result: 12 pt (at DPI 96)
//
// # Features
//
// The pxconv package accounts for screen density and user preferences,
// making it suitable for adaptive UI design. If input values are invalid
// (zero or negative), methods fall back to default values to avoid errors.
//
// Example:
//
//	metric := pxconv.NewMetric(0, -5, 0) // PxPerDp = 1, PxPerSp = 1, Dpi = 96.
//
// # Usage
//
// The pxconv package is useful for developing scalable UIs for mobile devices,
// desktops, and other graphical systems. Unit conversion helps create interfaces
// that render correctly on screens with varying densities.
package pxconv
