package src_test

import (
	"DMDS25/src"
	"testing"
)

type MockBTree struct {
	src.IBTree
	data map[uint64]uint64
}

func (b MockBTree) Get(key uint64) (uint64, error) {
	return b.data[key], nil
}

func (b MockBTree) Push(key uint64, value uint64) error {
	b.data[key] = value
	return nil
}

func (b MockBTree) GetRange(lowLimit uint64, highLimit uint64) (map[uint64]uint64, error) {
	myMap := make(map[uint64]uint64)
	for i := lowLimit; i <= highLimit; i++ {
		myMap[i] = b.data[i]
	}
	return myMap, nil
}

var MyBuffer1, _ = src.CreateNewBufferManager("./testFiles/", uint64(1024))
var myLoader1 = src.Loader{}
var tree1, _ = myLoader1.Load("tree1", MyBuffer1)

var MyBuffer2, _ = src.CreateNewBufferManager("./testFiles/", uint64(1024))
var myLoader2 = src.Loader{}
var tree2, _ = myLoader2.Load("tree2", MyBuffer2)

func TestBTreeSetup(t *testing.T) {
	MyBuffer, err := src.CreateNewBufferManager("./testFiles/", uint64(1024))
	if err != nil {
		t.Fatalf("error while initializing BufferManager: %v", err)
	}
	myLoader := src.Loader{}
	_, err = myLoader.Load("tree1", MyBuffer)

	if err != nil {
		t.Fatalf("error while loading tree with Loader: %v", err)
	}
}

func TestBTree1Get1(t *testing.T) {
	result, err := tree1.Get(1)
	if err != nil {
		t.Errorf("tree1.get(1) return error %d", err)
	}
	if result != 2 {
		t.Errorf("tree1.get(1) returned %d instead of 2", result)
	}
}

func TestBTree1Get11(t *testing.T) {
	result, err := tree1.Get(11)
	if err != nil {
		t.Errorf("tree1.get(11) return error %d", err)
	}
	if result != 12 {
		t.Errorf("tree1.get(11) returned %d instead of 12", result)
	}
}

func TestBTree2Get1(t *testing.T) {
	result, err := tree2.Get(1)
	if err != nil {
		t.Errorf("tree2.get(1) return error %d", err)
	}
	if result != 2 {
		t.Errorf("tree2.get(1) returned %d instead of 2", result)
	}
}

func TestBTree2Get11(t *testing.T) {
	result, err := tree2.Get(11)
	if err != nil {
		t.Errorf("tree2.get(11) return error %d", err)
	}
	if result != 12 {
		t.Errorf("tree2.get(11) returned %d instead of 12", result)
	}
}

func TestBTree2Get15(t *testing.T) {
	result, err := tree2.Get(15)
	if err != nil {
		t.Errorf("tree2.get(15) return error %d", err)
	}
	if result != 16 {
		t.Errorf("tree2.get(15) returned %d instead of 16", result)
	}
}

func TestBTree2Get21(t *testing.T) {
	result, err := tree2.Get(21)
	if err != nil {
		t.Errorf("tree2.get(21) return error %d", err)
	}
	if result != 22 {
		t.Errorf("tree2.get(21) returned %d instead of 22", result)
	}
}

func TestBTree2Get24(t *testing.T) {
	result, err := tree2.Get(24)
	if err != nil {
		t.Errorf("tree2.get(24) return error %d", err)
	}
	if result != 25 {
		t.Errorf("tree2.get(24) returned %d instead of 25", result)
	}
}

func TestIBTreeInsertAndFetch(t *testing.T) {
	tree := &MockBTree{data: map[uint64]uint64{}}
	err := tree.Push(1, 3)

	if err != nil {
		t.Errorf("tree.push(1, 3) return error %d", err)
	}

	result, err := tree.Get(1)

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
		err := tree.Push(uint64(i), uint64(i+1))
		if err != nil {
			t.Errorf("tree.push(%d, %d) returned error %d", uint64(i), uint64(i+1), err)
		}
	}

	result, err := tree.Get(1)

	if err != nil {
		t.Errorf("tree.Get(1) return error %d", err)
	}

	if result != 2 {
		t.Errorf("tree.Get(1) returned %d instead of 3", result)
	}

	rangeResult, rangeErr := tree.GetRange(1, 2)
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
