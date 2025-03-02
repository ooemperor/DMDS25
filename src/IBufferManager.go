package src

/*
IBufferManager Interface for the definition of behaviour for Buffer Managers
*/
type IBufferManager interface {
	create(dir string, memory uint64) (IBufferManager, error)
	open() error
	delete() error
	close() error

	pin() error
	unpin() error
}
