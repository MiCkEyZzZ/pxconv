# pxconv — Screen Unit Converter (dp, sp, px, inch, mm, pt)

[![godoc](https://godoc.org/github.com/MiCkEyZzZ/pxconv?status.svg)](https://pkg.go.dev/github.com/MiCkEyZzZ/pxconv?tab=doc)

`pxconv` is a Go package for working with screen units of measurement (dp, sp, px, inch, mm, pt), considering screen density.


## Features

- **Support for main units**: `dp`, `sp`, `px`, `inch`, `mm`, `pt`.
- **Customizable screen density**: parameters `PxPerDp`, `PxPerSp`, и `Dpi`.
- **Unit conversion**: convenient methods for converting between all supported units.
- **Handling of invalid values**: default replacement (1 by default) to prevent errors.

## Installation

To install the package, use the following command:

```zsh
go get github.com/MiCkEyZzZ/pxconv
```

## Usage Example

Here is an example of how to use the package:

```go
package main

import (
	"fmt"

	"github.com/MiCkEyZzZ/pxconv"
)

func main() {
	// Create an instance of Metric with custom density settings
	m := pxconv.Metric{PxPerDp: 2, PxPerSp: 1.5, Dpi: 96}

	// Convert 10 dp to pixels
	px := m.DpToPx(pxconv.Dp(10))
	fmt.Println(px) // Output: 20
}
```

## API

The package provides the following key methods for unit conversions:

- DpToPx — Convert dp to px.
- PxToDp — Convert px to dp.
- SpToPx — Convert sp to px.
- PxToSp — Convert px to sp.
- InchToPx — Convert inches to px.
- PxToInch — Convert px to inches.
- MmToPx — Convert millimeters to px.
- PxToMm — Convert px to millimeters.
- PtToPx — Convert points to px.
- PxToPt — Convert px to points.

A full list of methods and their descriptions can be found in the [documentation](https://pkg.go.dev/github.com/MiCkEyZzZ/pxconv).

## License

This package is distributed under the MIT License. A full copy of the license is available in the [License](./LICENSE) file.
