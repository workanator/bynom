package bynom

// StateHasAllBits tests if all states dst have all bits s set to 1.
func StateHasAllBits(dst ...*State) TestState {
	return func(s State) bool {
		for _, v := range dst {
			if *v&s != s {
				return false
			}
		}

		return true
	}
}

// StateHasAnyBits tests if all states dst have at least on bit from bits s set to 1.
func StateHasAnyBits(dst ...*State) TestState {
	return func(s State) bool {
		for _, v := range dst {
			if *v&s == 0 {
				return false
			}
		}

		return true
	}
}
