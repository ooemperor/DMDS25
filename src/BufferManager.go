package src

/*
IBufferManager Interface for the definition of behaviour for Buffer Managers
*/
type IBufferManager interface {
	/*
		create a new BufferManager in the given directory dir of size memory
	*/
	create(dir string, memory uint64) (IBufferManager, error)

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
	Pin(fileID string, pageInFile uint64) (uint, error)

	/*
		unpin a given page from the btree to memory
	*/
	Unpin(pageID uint64) error
}
