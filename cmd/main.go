package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/carlca/gowig/output"
)

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer f.Close()

	// fmt.Println(len(data))
	// if err != nil {
	// 	fmt.Println("File printing error", err)
	// }

	p := output.Param{"device", "Chain"}
	fmt.Println(p.Key)
	fmt.Println(p.Value)
	

}

func readFromFile(file *os.File, offset, size int) ([]byte, error) {
	res := make([]byte, size)
	if _, err := file.ReadAt(res, int64(offset)); err != nil {
		return nil, err
	}
	return res, nil
}
