package src

import (
	"errors"
)

/*
IBTree Interface that defines the basic functionality of the B tree
*/
type IBTree interface {
	// Get Method to retreive a value for the given key
	Get(key uint64) (uint64, error)
	// Push insert a new key value pair
	Push(key uint64, value uint64) error
	// GetRange Retreive a the values for a given range of keys
	GetRange(lowLimit uint64, highLimit uint64) (map[uint64]uint64, error)
}

/*
BTree is the type definition of a BTree and implements the get, push and getRange method of the above Interface
*/
type BTree struct {
	Name       string //defines the filename of the BTree for loading
	RootPageId uint64
	Manager    *BufferManager
}

/*
Get fetches the value out of the index
*/
func (bm *BTree) Get(key uint64) (uint64, error) {
	return bm.traverse(key, 0, bm.RootPageId)
}

func (bm *BTree) traverse(key uint64, currentLevel int, nextPageId uint64) (uint64, error) {
	id, err := bm.Manager.Pin(bm.Name, nextPageId)

	if err != nil {
		return 0, err
	}

	page := bm.Manager.Pages[id]

	for i := 0; i < len(page.Keys); i++ {
		// i limit the pages to 1 level deep for the sake of simplicity
		if currentLevel == 1 {
			// we are on leave level so we can start to look for exact key
			if key == page.Keys[i] {
				return page.Values[i], nil
			} else if key > page.Keys[i] {
				continue
			} else {
				return 0, errors.New("key not found on leave level")
			}
		} else if i == len(page.Keys)-1 {
			// we have reached the end of the keys, take the right most path down the tree
			return bm.traverse(key, currentLevel+1, page.Values[i+1])
		} else if key > page.Keys[i] && page.Keys[i] != 0 {
			// go one key to the right since we have not reached the end yet
			continue
		} else {
			// traverse into the next page
			return bm.traverse(key, currentLevel+1, page.Values[i])
		}
	}
	return 0, errors.New("error in traversing")
}

func (bm *BTree) Push(key uint64, value uint64) error {
	return nil
}

func (bm *BTree) GetRange(low uint64, high uint64) (map[uint64]uint64, error) {
	return nil, nil
}
