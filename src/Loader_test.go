package src

import (
	"os"
	"reflect"
	"testing"
)

/*
TestLoader tests if we receive an expected error when closing a page that is not open
*/
func TestLoaderSetup(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, _ := os.Create("./testFileForLoader")
	_, _ = file.Write([]byte("1;2;3;4;5;6;a;b;c;d;e;f;g"))
	_ = file.Close()
	loader := Loader{}

	defer func() {
		_ = os.Remove("./testFileForLoader")
	}()
	tree, err := loader.Load("./testFileForLoader", myBuffer)
	if err != nil {
		t.Errorf("Error loading file: %v", err)
	}
	expectedBTree := &BTree{name: "./testFileForLoader", rootPageId: 0}
	if !reflect.DeepEqual(tree, expectedBTree) {
		t.Errorf("loaded tree does not match expected tree")
	}
}
