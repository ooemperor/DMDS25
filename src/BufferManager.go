package src

import (
	"errors"
	"log"
	"os"
	"reflect"
)

/*
IBufferManager Interface for the definition of behaviour for Buffer Managers
*/
type IBufferManager interface {
	/*
		create a new BufferManager in the given directory dir of size memory
	*/
	//create(dir string, memory uint64) (BufferManager, error)

	/*
		open the specific b tree
	*/
	Open(fileID string) error

	/*
		delete the file of the btree from disk
	*/
	Delete(fileID string) error

	/*
		close the opened b tree file
	*/
	Close() error

	/*
		pin a given page from the btree to memory
	*/
	Pin(fileID string, pageInFile uint64) (uint64, error)

	/*
		unpin a given page from the btree to memory
	*/
	Unpin(pageID uint64) error
}

type BufferManager struct {
	Pages       [10]Page
	dir         string
	memory      uint64
	tmpFileData []byte
}

func CreateNewBufferManager(dir string, memory uint64) (IBufferManager, error) {
	return &(BufferManager{dir: dir, memory: memory}), nil
}

func (bm *BufferManager) Open(fileID string) error {
	dat, err := os.ReadFile(bm.dir + fileID)
	bm.tmpFileData = dat
	return err
}

func (bm *BufferManager) Close() error {
	if bm.tmpFileData == nil {
		return errors.New("no file to close")

	} else {
		bm.tmpFileData = nil
		return nil
	}
}

func (bm *BufferManager) Delete(fileID string) error {
	dat, _ := os.ReadFile(bm.dir + fileID)

	if dat != nil {
		return os.Remove(bm.dir + fileID)
	}
	return errors.New("no file to delete")
}

func (bm *BufferManager) Pin(fileID string, pageInFile uint64) (uint64, error) {
	err := bm.Open(fileID)
	if err != nil {
		return 0, err
	}
	if bm.tmpFileData == nil {
		return 0, errors.New("no tmpFileData found")
	}

	for i := uint64(0); i < uint64(len(bm.Pages)); i++ {
		if !reflect.DeepEqual(bm.Pages[i], Page{}) {
			continue
		} else {
			//TODO deserialize the page from disk
			page, err := bm.deserialize()
			if err != nil {
				log.Fatalf("An error occured while deserializing: %v", err)
			}
			bm.Pages[i] = page
			return i, nil
		}
	}
	return 0, errors.New("no file to pin")
}

func (bm *BufferManager) Unpin(pageID uint64) error {
	if bm.tmpFileData == nil {
		return errors.New("no page to depin")
	}
	bm.Pages[pageID] = Page{}
	return nil
}

/*
deserialize the byte values currently present in the tmpFileData or throw an error
*/
func (bm *BufferManager) deserialize() (Page, error) {
	//TODO implement the deserialize part
	return Page{}, nil
}

/*
deserialize the byte values currently present in the tmpFileData or throw an error
*/
func (bm *BufferManager) serialize(pageID uint64) error {
	//TODO implement the serialize part
	return nil
}
