package src_test

import (
	"DMDS25/src"
	"testing"
)

type MockBTree struct {
	src.IBTree
	data map[uint64]uint64
}

func (b MockBTree) get(key uint64) (uint64, error) {
	return b.data[key], nil
}

func (b MockBTree) push(key uint64, value uint64) error {
	b.data[key] = value
	return nil
}

func (b MockBTree) getRange(lowLimit uint64, highLimit uint64) (map[uint64]uint64, error) {
	myMap := make(map[uint64]uint64)
	for i := lowLimit; i <= highLimit; i++ {
		myMap[i] = b.data[i]
	}
	return myMap, nil
}

func TestIBTreeEmpty(t *testing.T) {
	tree := &MockBTree{}

	result, err := tree.get(1)
	if err != nil {
		t.Errorf("tree.get(1) return error %d", err)
	}
	if result != 0 {
		t.Errorf("tree.get(1) returned %d instead of 0", result)
	}
}

func TestIBTreeNonEmpty(t *testing.T) {
	tree := &MockBTree{data: map[uint64]uint64{}}
	tree.data[1] = 1

	result, err := tree.get(1)
	if err != nil {
		t.Errorf("tree.get(1) return error %d", err)
	}
	if result != 1 {
		t.Errorf("tree.get(1) returned %d instead of 0", result)
	}
}

func TestIBTreeInsertAndFetch(t *testing.T) {
	tree := &MockBTree{data: map[uint64]uint64{}}
	err := tree.push(1, 3)

	if err != nil {
		t.Errorf("tree.push(1, 3) return error %d", err)
	}

	result, err := tree.get(1)

	if err != nil {
		t.Errorf("tree.get(1) return error %d", err)
	}

	if result != 3 {
		t.Errorf("tree.get(1) returned %d instead of 3", result)
	}
}

func TestIBTreeInsertAndFetchRange(t *testing.T) {
	tree := &MockBTree{data: map[uint64]uint64{}}

	for i := 0; i < 5; i++ {
		err := tree.push(uint64(i), uint64(i+1))
		if err != nil {
			t.Errorf("tree.push(%d, %d) returned error %d", uint64(i), uint64(i+1), err)
		}
	}

	result, err := tree.get(1)

	if err != nil {
		t.Errorf("tree.get(1) return error %d", err)
	}

	if result != 2 {
		t.Errorf("tree.get(1) returned %d instead of 3", result)
	}

	rangeResult, rangeErr := tree.getRange(1, 2)
	if rangeErr != nil {
		t.Errorf("tree.getRange(1, 2) return error %d", rangeErr)
	}
	expectedMap := map[uint64]uint64{uint64(1): uint64(2), uint64(2): uint64(3)}

	for k, v := range expectedMap {
		if rangeResult[k] != v {
			t.Errorf("tree.getRange(1, 2) returned %#v instead of %#v", rangeResult, expectedMap)
		}
	}
}
