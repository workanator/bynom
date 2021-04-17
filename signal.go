package bynom

// State keeps the current state which can be modified with signals.
type State int

// ModifyState modifies state on signal.
type ModifyState func(State) error

// TestState tests if state conforms condition.
type TestState func(State) bool

// Signal invokes all modification handlers fns passing them the value v.
// The function does no affect the plate read position.
// If any signal handler from fns returns non-nil error the function fails with that error.
func Signal(v State, fns ...ModifyState) Nom {
	return func(Plate) (err error) {
		for _, mod := range fns {
			if err = mod(v); err != nil {
				break
			}
		}
		return
	}
}

// RequireSignal runs state tests fns for against the value v.
func RequireSignal(v State, fns ...TestState) Nom {
	return func(Plate) error {
		for _, test := range fns {
			if !test(v) {
				return ErrStateTestFailed{
					Assert: v,
				}
			}
		}

		return nil
	}
}
