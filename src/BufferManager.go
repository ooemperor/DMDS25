package src

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
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
	Pages        [10]Page
	dir          string
	memory       uint64
	tmpFileData  []byte
	openFileName string
}

func CreateNewBufferManager(dir string, memory uint64) (*BufferManager, error) {
	return &(BufferManager{dir: dir, memory: memory}), nil
}

func (bm *BufferManager) Open(fileID string) error {
	dat, err := os.ReadFile(bm.dir + fileID)
	bm.tmpFileData = dat
	bm.openFileName = fileID
	return err
}

func (bm *BufferManager) Close() error {
	if bm.tmpFileData == nil {
		return errors.New("no file to close")

	} else {
		bm.tmpFileData = nil
		bm.openFileName = ""
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

	stringArray := strings.Split(string(bm.tmpFileData), ";")

	if bm.tmpFileData == nil || bm.openFileName == "" {
		if len(stringArray) != 13 {
			return Page{}, errors.New("deserialization failed, there is no open file")
		}
	}

	if len(stringArray) != 13 {
		return Page{}, errors.New("deserialization failed, the file does not contain 13 elements")
	}

	// init the tmp arrays
	keys := [6]uint64{}
	values := [7]uint64{}

	// read the keys
	for i := 0; i < 6; i++ {
		if stringArray[i] != "" {
			keys[i], _ = strconv.ParseUint(stringArray[i], 10, 64)
		} else {
			values[i] = 0
		}
	}

	// read the values
	offset := 6
	for i := 0; i < 7; i++ {
		index := offset + i
		if stringArray[index] != "" {
			values[i], _ = strconv.ParseUint(stringArray[index], 10, 64)
		} else {
			values[i] = 0
		}
	}
	return Page{Keys: keys, Values: values, Name: bm.openFileName}, nil
}

/*
deserialize the byte values currently present in the tmpFileData or throw an error
*/
func (bm *BufferManager) serialize(pageID uint64) error {
	page := bm.Pages[pageID]

	var outputString string = ""

	for i := 0; i < len(page.Keys); i++ {
		if tmpKey := page.Keys[i]; tmpKey != 0 {
			outputString = outputString + strconv.FormatUint(uint64(page.Keys[i]), 10)
		}
		outputString = outputString + ";"
	}

	for i := 0; i < len(page.Values); i++ {
		if tmpValue := page.Values[i]; tmpValue != 0 {
			outputString = outputString + strconv.FormatUint(uint64(page.Keys[i]), 10)
		}
		outputString = outputString + ";"
	}
	file, err := os.Open(bm.dir + page.Name)
	if err != nil {
		log.Fatalf("An error occured while opening the file: %s", err)
	}
	err = file.Truncate(0)
	if err != nil {
		log.Fatalf("An error occured while truncating file: %s", err)
	}
	_, err = file.Write([]byte(outputString))
	if err != nil {
		log.Fatalf("An error occured while writing to file: %s", err)
	}

	_ = file.Close()
	return nil
}
