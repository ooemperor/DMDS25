package src

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

type MockBufferManager struct {
	IBufferManager
	Pages       [3][]byte
	dir         string
	memory      uint64
	tmpFileData []byte
}

func (bm *MockBufferManager) create(dir string, memory uint64) (MockBufferManager, error) {
	return MockBufferManager{dir: dir, memory: memory}, nil
}

func (bm *MockBufferManager) Open(fileID string) error {
	dat, err := os.ReadFile(bm.dir + fileID)
	bm.tmpFileData = dat
	return err
}

func (bm *MockBufferManager) Close() error {
	if bm.tmpFileData == nil {
		return errors.New("no file to close")

	} else {
		bm.tmpFileData = nil
		return nil
	}
}

func (bm *MockBufferManager) Delete(fileID string) error {
	dat, _ := os.ReadFile(bm.dir + fileID)

	if dat != nil {
		return os.Remove(bm.dir + fileID)
	}
	return errors.New("no file to delete")
}

func (bm *MockBufferManager) Pin(fileID string, pageInFile uint64) (uint64, error) {
	err := bm.Open(fileID)
	if err != nil {
		return 0, err
	}
	if bm.tmpFileData == nil {
		return 0, errors.New("no tmpFileData found")
	}

	for i := uint64(0); i < uint64(len(bm.Pages)); i++ {
		if !bytes.Equal(bm.Pages[i], make([]byte, 0)) {
			continue
		} else {
			bm.Pages[i] = bm.tmpFileData
			return i, nil
		}
	}
	return 0, errors.New("no file to pin")
}

func (bm *MockBufferManager) Unpin(pageID uint64) error {
	if bm.tmpFileData == nil {
		return errors.New("no page to depin")
	}
	bm.Pages[pageID] = make([]byte, 0)
	return nil
}

func TestIBufferManagerSetup(t *testing.T) {
	var _, err = (&MockBufferManager{}).create("./", uint64(1024))
	if err != nil {
		t.Fatal(err)
	}
}

/*
TestIBufferManagerOpen Checks if we can open a file that exists without an error
*/
func TestIBufferManagerOpen(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
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
TestIBufferManagerOpenWithError checks that an error is raised when a non existant file is tried to open
*/
func TestIBufferManagerOpenWithError(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	err := myBuffer.Open("testFileNonExistent")
	if err == nil {
		t.Errorf("Open should return an error but does not")
	}
}

/*
TestIBufferManagerClose tests if we can close an open file properly
*/
func TestIBufferManagerClose(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
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
TestIBufferManagerDelete tests if we receive an expected error when closing a page that is not open
*/
func TestIBufferManagerDelete(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
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
TestIBufferManagerDeleteWithError tests if we receive an expected error when closing a page that is not open
*/
func TestIBufferManagerDeleteWithError(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	err := myBuffer.Delete("testFileForDeleteNonExistent")
	if err == nil {
		t.Fatal("Delete should return an error but does not")
	}
}

func TestIBufferManagerPin(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	file, _ := os.Create("./testFileForPin")
	_, _ = file.Write([]byte("test"))
	_ = file.Close()
	id, err := myBuffer.Pin("testFileForPin", uint64(0))
	if err != nil {
		t.Fatal(err)
	}

	if id < 0 || id >= uint64(len(myBuffer.Pages)) {
		t.Fatal("Id is invalid")
	}

	_ = os.Remove("./testFileForPin")
}

func TestIBufferManagerPinWithError(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	_, err := myBuffer.Pin("testFileForPinNonExistent", uint64(0))
	if err == nil {
		t.Fatal("This should return an error but does not")
	}
}

func TestIBufferManagerUnpin(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	file, _ := os.Create("./testFileForUnPin")
	_ = file.Close()

	id, err := myBuffer.Pin("testFileForUnPin", uint64(0))
	if err != nil {
		t.Errorf("Error occured when trying to pin: %s", err)
	}
	_ = os.Remove("./testFileForUnPin")

	err = myBuffer.Unpin(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIBufferManagerUnpinWithError(t *testing.T) {
	var myBuffer, _ = (&MockBufferManager{}).create("./", uint64(1024))
	id, _ := myBuffer.Pin("testFileForUnPinNonExistent", uint64(0))

	err := myBuffer.Unpin(id)
	if err == nil {
		t.Fatal("this should have returned an error but does not")
	}
}
