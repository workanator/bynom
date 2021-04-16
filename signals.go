package bynom

// ResetSignals sets all signal variables in dst to false.
// The function does no affect the plate read position.
func ResetSignals(dst ...*bool) Nom {
	return func(p Plate) error {
		for _, signal := range dst {
			*signal = false
		}
		return nil
	}
}

// SetSignal sets the signal variable to true.
// The function does no affect the plate read position.
func SetSignal(dst *bool) Nom {
	return func(p Plate) error {
		*dst = true
		return nil
	}
}
