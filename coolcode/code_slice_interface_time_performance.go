package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir1 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))
}
