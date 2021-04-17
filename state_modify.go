package bynom

// ReplaceState replaces states dst with the new state s.
func ReplaceState(dst ...*State) ModifyState {
	return func(s State) error {
		for _, v := range dst {
			*v = s
		}
		return nil
	}
}

// SetStateBits sets bits in states dst to 1 which are 1 in the state s.
func SetStateBits(dst ...*State) ModifyState {
	return func(s State) error {
		for _, v := range dst {
			*v = *v | s
		}
		return nil
	}
}

// ResetStateBits sets bits in states dst to 0 which are 1 in the state s.
func ResetStateBits(dst ...*State) ModifyState {
	return func(s State) error {
		for _, v := range dst {
			*v = *v & ^s
		}
		return nil
	}
}
