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

/*
BTree is the type definition of a BTree and implements the get, push and getRange method of the above Interface
*/
type BTree struct {
	Name       string //defines the filename of the BTree for loading
	RootPageId uint64
	Manager    *BufferManager
}

func (bm *BTree) get(key uint64) (uint64, error) {
	return 0, nil
}

func (bm *BTree) push(key uint64, value uint64) error {
	return nil
}

func (bm *BTree) getRange(low uint64, high uint64) (map[uint64]uint64, error) {
	return nil, nil
}
