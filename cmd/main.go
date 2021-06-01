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

	streamPos := 0x7f
	chunk, err := readFromFile(f, streamPos, 4) //read 4 bytes in the file from byte 0x7f

	var size int32
	buf := bytes.NewReader(chunk)
	err = binary.Read(buf, binary.BigEndian, &size)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(size)
}

func readFromFile(file *os.File, offset, size int) ([]byte, error) {
	res := make([]byte, size)
	if _, err := file.ReadAt(res, int64(offset)); err != nil {
		return nil, err
	}
	return res, nil
}
