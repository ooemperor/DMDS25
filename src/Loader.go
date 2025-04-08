package src

type Loader struct{}

/*
Load loads the initial root node of a BTree and returns it.
Each page has a seperate file with an id as the name
*/
func (l *Loader) Load(name string, manager *BufferManager) (*BTree, error) {
	id, err := manager.Pin(name, 0)
	if err != nil {
		return nil, err
	}
	return &BTree{Name: name, RootPageId: id, Manager: manager}, nil
}
