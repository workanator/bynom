package state

// Bits provides functionality to manipulate bits in int.
// Methods used to manipulate and test bits conforms types bynom.ModifyState and bynom.TestState.
type Bits int64

// NewBits create a new Bits instance with all bits set to 0.
func NewBits() *Bits {
	return new(Bits)
}

// Replace replaces all bits with the new value v.
func (bits *Bits) Replace(v int64) error {
	*bits = Bits(v)
	return nil
}

// Set sets bits to 1 which are 1 in the value v.
func (bits *Bits) Set(v int64) error {
	*bits = *bits | Bits(v)
	return nil
}

// Reset sets bits to 0 which are 1 in the value v.
func (bits *Bits) Reset(v int64) error {
	*bits = *bits & Bits(^v)
	return nil
}

// AllSet tests if the instance has all bits v set to 1.
func (bits *Bits) AllSet(v int64) bool {
	return *bits&Bits(v) == Bits(v)
}

// AnySet tests if the instance has at least on bit from bits v set to 1.
func (bits *Bits) AnySet(v int64) bool {
	return *bits&Bits(v) > 0
}

// NothingSet tests if the instance has all bits v set to 0.
func (bits *Bits) NothingSet(v int64) bool {
	return *bits&Bits(v) == 0
}

// Equal tests if the instance equals to the value v.
func (bits *Bits) Equal(v int64) bool {
	return *bits == Bits(v)
}

// Int64 returns the instance as int value.
func (bits *Bits) Int64() int64 {
	return int64(*bits)
}
