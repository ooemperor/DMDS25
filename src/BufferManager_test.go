package src

import (
	"os"
	"testing"
)

/*
TestBufferManagerSetup tests if the setup function works as expected or not.
*/
func TestBufferManagerSetup(t *testing.T) {
	var _, err = CreateNewBufferManager("./", uint64(1024))
	if err != nil {
		t.Fatal(err)
	}
}

/*
TestBufferManagerOpen Checks if we can open a file that exists without an error
*/
func TestBufferManagerOpen(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, err := os.Create("./testFile")
	_ = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = myBuffer.Open("testFile")
	if err != nil {
		t.Fatal(err)
	}
	_ = os.Remove("./testFile")
}

/*
TestBufferManagerOpenWithError checks that an error is raised when a non existant file is tried to open
*/
func TestBufferManagerOpenWithError(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	err := myBuffer.Open("testFileNonExistent")
	if err == nil {
		t.Errorf("Open should return an error but does not")
	}
}

/*
TestBufferManagerClose tests if we can close an open file properly
*/
func TestBufferManagerClose(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, err := os.Create("./testFileForClose")
	_ = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = myBuffer.Open("testFileForClose")
	if err != nil {
		t.Fatal(err)
	}
	_ = os.Remove("./testFileForClose")
}

/*
TestBufferManagerDelete tests if we receive an expected error when closing a page that is not open
*/
func TestBufferManagerDelete(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, err := os.Create("./testFileForDelete")
	_ = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = myBuffer.Delete("testFileForDelete")
	if err != nil {
		t.Fatal(err)
	}
}

/*
TestBufferManagerDeleteWithError tests if we receive an expected error when closing a page that is not open
*/
func TestBufferManagerDeleteWithError(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	err := myBuffer.Delete("testFileForDeleteNonExistent")
	if err == nil {
		t.Fatal("Delete should return an error but does not")
	}
}

/*
TestBufferManagerPin tests if the pinning of a page works correctly or not
*/
func TestBufferManagerPin(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, _ := os.Create("./testFileForPin")
	_, _ = file.Write([]byte("1;2;3;4;5;6;a;b;c;d;e;f;g"))
	_ = file.Close()
	id, err := myBuffer.Pin("testFileForPin", uint64(0))
	if err != nil {
		t.Fatal(err)
	}

	if id < 0 || id >= uint64(len(myBuffer.Pages)) {
		t.Fatal("Id is invalid")
	}

	if myBuffer.PageMap[uint64(0)] != id {
		t.Fatal("PageMap has not been updated properly")
	}

	_ = os.Remove("./testFileForPin")
}

/*
TestBufferManagerPinWithError tests that the ping successfully return error when expected
*/
func TestBufferManagerPinWithError(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	_, err := myBuffer.Pin("testFileForPinNonExistent", uint64(0))
	if err == nil {
		t.Fatal("This should return an error but does not")
	}
}

/*
TestBufferManagerPin tests if the unpinning of a page works correctly or not
*/
func TestBufferManagerUnpin(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, _ := os.Create("./testFileForUnPin")
	_, _ = file.Write([]byte("1;2;3;4;5;6;a;b;c;d;e;f;g"))
	_ = file.Close()

	id, err := myBuffer.Pin("testFileForUnPin", uint64(0))

	if err != nil {
		t.Errorf("Error occured when trying to pin: %s", err)
	}

	if myBuffer.PageMap[uint64(0)] != id {
		t.Fatalf("PageMap has not been updated properly, expected %v got %v", id, myBuffer.PageMap[uint64(0)])
	}
	_ = os.Remove("./testFileForUnPin")

	err = myBuffer.Unpin(id)

	if err != nil {
		t.Fatal(err)
	}

	if len(myBuffer.PageMap) != 0 {
		t.Fatalf("PageMap has not been updated properly with removed values, lenght is %v", len(myBuffer.PageMap))
	}
}

func TestBufferManagerUnpinWithError(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	id, _ := myBuffer.Pin("testFileForUnPinNonExistent", uint64(0))

	err := myBuffer.Unpin(id)
	if err == nil {
		t.Fatal("this should have returned an error but does not")
	}
}

func TestBufferManagerSerializeRootOnly(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, _ := os.Create("./testFileForSerialize")
	_, _ = file.Write([]byte("1;2;3;4;5;6;a;b;c;d;e;f;g"))
	_ = file.Close()

	defer func() {
		_ = os.Remove("./testFileForSerialize")
	}()

	keys := [6]uint64{uint64(11), uint64(12), uint64(13), uint64(14), uint64(15), uint64(16)}
	values := [7]uint64{uint64(21), uint64(22), uint64(23), uint64(24), uint64(25), uint64(26), uint64(27)}
	page := Page{pageId: 0, Name: "testFileForSerialize", Keys: keys, Values: values}

	myBuffer.Pages[0] = page
	myBuffer.PageMap[0] = 0

	err := myBuffer.serialize(0)

	if err != nil {
		t.Fatalf("error while deserializing: %v", err)
	}
}

func TestBufferManagerSerializeTwoNodes(t *testing.T) {
	var myBuffer, _ = CreateNewBufferManager("./", uint64(1024))
	file, _ := os.Create("./testFileForSerializeTwoNodes")
	_, _ = file.Write([]byte("1;2;3;4;5;6;a;b;c;d;e;f;g\n1;2;3;4;5;6;a;b;c;d;e;f;g"))
	_ = file.Close()

	defer func() {
		_ = os.Remove("./testFileForSerializeTwoNodes")
	}()

	keys := [6]uint64{uint64(11), uint64(12), uint64(13), uint64(14), uint64(15), uint64(16)}
	values := [7]uint64{uint64(21), uint64(22), uint64(23), uint64(24), uint64(25), uint64(26), uint64(27)}
	page := Page{pageId: 0, Name: "testFileForSerializeTwoNodes", Keys: keys, Values: values}
	page1 := Page{pageId: 1, Name: "testFileForSerializeTwoNodes", Keys: keys, Values: values}

	myBuffer.Pages[0] = page
	myBuffer.Pages[1] = page1
	myBuffer.PageMap[0] = 0
	myBuffer.PageMap[1] = 1

	err := myBuffer.serialize(0)

	if err != nil {
		t.Fatalf("error while deserializing page 0: %v", err)
	}

	err = myBuffer.serialize(1)
	if err != nil {
		t.Fatalf("error while deserializing page1: %v", err)
	}
}
