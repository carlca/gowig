package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/carlca/gowig/output"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
    if err != nil {
		fmt.Println("File reading error", err)
        return
    }
	fmt.Println(len(data))
	if err != nil {
		fmt.Println("File printing error", err)
	}

	p := output.Param{"device", "Chain"}
	// streamPos := 0x
	fmt.Println(p.Key)
	fmt.Println(p.Value)
}

func GetChunk(streamPos int64, data []byte) int64 {
	return 0
}
