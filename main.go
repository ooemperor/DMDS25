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

	res, _ := btree.Get(11)
	fmt.Printf("result is: %v\n", res)
	_ = btree.Push(13, 14)

	res, _ = btree.Get(13)
	fmt.Printf("result is: %v\n", res)

	fmt.Println(btree.Manager.Pages)
	_ = btree.Manager.Flush()
}
