package src

/*
IBTree Interface that defines the basic functionality of the B tree
*/
type IBTree interface {
	// get Method to retreive a value for the given key
	get(key uint64) (uint64, error)
	// push insert a new key value pair
	push(key uint64, value uint64) error
	// getRange Retreive a the values for a given range of keys
	getRange(lowLimit uint64, highLimit uint64) (map[uint64]uint64, error)
}
