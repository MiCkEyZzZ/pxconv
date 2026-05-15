package density

// EnsurePositive returns a positive value.
// If the input is zero or negative, it returns 1.
func EnsurePositive(value float32) float32 {
	if value <= 0 {
		return 1
	}
	return value
}
