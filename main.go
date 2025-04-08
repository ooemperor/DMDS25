package main

import (
	"DMDS25/src"
	"fmt"
)

func main() {
	var myBuffer, _ = src.CreateNewBufferManager("./", uint64(1024))
	myLoader := src.Loader{}
	btree, err := myLoader.Load("startTree", myBuffer)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(btree)
	fmt.Println(myBuffer.Pages)
	fmt.Println(myBuffer.Pages[btree.RootPageId])
}
