package src

/*
Page definition with its keys and values
For root and non-leaf nodes we use all the 7 values.
For leaf nodes we are only gonna use the first 6 values.
*/
type Page struct {
	Name   string // the name of the file the page is stored in on disk
	Keys   [6]uint64
	Values [7]uint64
}
