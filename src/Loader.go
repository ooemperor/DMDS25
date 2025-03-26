package src

import "log"

type Loader struct{}

/*
Load loads the initial root node of a BTree and returns it.
Each page has a seperate file with an id as the name
*/
func (l *Loader) Load(name string, manager *BufferManager) (*BTree, error) {
	id, err := manager.Pin(name, 0)
	if err != nil {
		log.Fatal(err)
	}
	return &BTree{name: name, rootPageId: id}, nil
}
