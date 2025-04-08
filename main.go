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

	res, _ := btree.Get(11)
	fmt.Printf("result is: %v\n", res)
	err = btree.Push(13, 14)
	res, err = btree.Get(13)
	fmt.Printf("result is: %v\n", res)
	_ = btree.Manager.Flush()
}
