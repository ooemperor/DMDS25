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
	Pin(fileID string) (uint64, error)

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
	PageMap      map[uint64]uint64
}

func CreateNewBufferManager(dir string, memory uint64) (*BufferManager, error) {
	mapping := make(map[uint64]uint64)
	return &(BufferManager{dir: dir, memory: memory, PageMap: mapping}), nil
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
	err = bm.Open(fileID)
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
			page, err := bm.deserialize(pageInFile)
			if err != nil {
				return 0, err
			}
			bm.Pages[i] = page
			_ = bm.Close()

			// adding the page to the mapping
			bm.PageMap[pageInFile] = i

			return i, nil
		}
	}
	return 0, errors.New("no file to pin")
}

func (bm *BufferManager) Unpin(pageID uint64) error {

	if reflect.DeepEqual(bm.Pages[pageID], Page{}) {
		return errors.New("there is no page to depin at this Id")
	}
	bm.Pages[pageID] = Page{}
	err := bm.RemoveMapEntryByValue(pageID)
	if err != nil {
		return err
	}

	return nil
}

func (bm *BufferManager) RemoveMapEntryByValue(value uint64) error {
	for key, val := range bm.PageMap {
		if val == value {
			delete(bm.PageMap, key)
			return nil
		}
	}
	return errors.New("no page with this Id found to remove from map")
}

/*
deserialize the byte values currently present in the tmpFileData or throw an error
*/
func (bm *BufferManager) deserialize(pageInFile uint64) (Page, error) {

	pageRowString := strings.Split(string(bm.tmpFileData), "\n")[pageInFile]
	stringArray := strings.Split(pageRowString, ";")

	if bm.tmpFileData == nil || bm.openFileName == "" {
		if len(stringArray) != 13 {
			return Page{}, errors.New("deserialization failed, there is no open file")
		}
	}

	if len(stringArray) != 13 {
		return Page{}, errors.New("deserialization failed, the row does not contain 13 elements")
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
	return Page{Keys: keys, Values: values, pageId: pageInFile, Name: bm.openFileName}, nil
}

/*
serialize the given page and write it to disk
*/
func (bm *BufferManager) serialize(pageID uint64) error {
	page := bm.Pages[pageID]
	_ = bm.Open(bm.dir + page.Name)
	var outputString string = ""

	pageRowStrings := strings.Split(string(bm.tmpFileData), "\n")

	for rowId := uint64(0); rowId < uint64(len(pageRowStrings)); rowId++ {
		if pageID == rowId {
			for i := 0; i < len(page.Keys); i++ {
				if tmpKey := page.Keys[i]; tmpKey != 0 {
					outputString = outputString + strconv.FormatUint(uint64(page.Keys[i]), 10)
				}
				outputString = outputString + ";"
			}

			for i := 0; i < len(page.Values); i++ {
				if tmpValue := page.Values[i]; tmpValue != 0 {
					outputString = outputString + strconv.FormatUint(uint64(page.Values[i]), 10)
				}
				outputString = outputString + ";"
			}
			outputString = outputString + "\n"
		} else {
			outputString = outputString + pageRowStrings[rowId] + "\n"
		}
	}
	_ = bm.Close()

	file, err := os.OpenFile(bm.dir+page.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 777)

	if err != nil {
		log.Fatalf("An error occured while opening the file: %s", err)
	}

	err = os.Truncate(bm.dir+page.Name, 0)
	if err != nil {
		log.Fatalf("An error occured while truncating file: %s", err)
	}
	_, err = file.Write([]byte(outputString))
	if err != nil {
		log.Fatalf("An error occured while writing to file: %s", err)
	}

	_ = bm.Close()

	_ = file.Close()
	return nil
}
