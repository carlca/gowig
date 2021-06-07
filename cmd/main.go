package main

// Let @crossm0d know how it's going, and add in date/time tags

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		GenerateDummyOutput()
	} else {
		ProcessPreset(os.Args[1])
	}
}

func GenerateDummyOutput() {
	fmt.Println()
}

func ProcessPreset(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var streamPos int32 = 0x7f
	var size int32
	var text string

	if streamPos, size, text, err = readNextSizeAndChunk(f, streamPos); err != nil {
		return err
	}

	fmt.Println("size: ", size)
	fmt.Printf("stringPos: %X\n", streamPos)
	fmt.Println("text: ", text)

	streamPos++

	if streamPos, size, text, err = readNextSizeAndChunk(f, streamPos); err != nil {
		return err
	}

	fmt.Println("size: ", size)
	fmt.Printf("stringPos: %X\n", streamPos)
	fmt.Println("text: ", text)

	streamPos++
	streamPos++
	streamPos++
	streamPos++

	if streamPos, size, text, err = readNextSizeAndChunk(f, streamPos); err != nil {
		return err
	}

	fmt.Println("size: ", size)
	fmt.Printf("stringPos: %X\n", streamPos)
	fmt.Println("text: ", text)

	return nil
}

func readNextSizeAndChunk(f *os.File, streamPos int32) (int32, int32, string, error) {
	var size int32
	var text string
	var err error
	if streamPos, size, err = readIntChunk(f, streamPos); err != nil {
		return 0, 0, "", err
	}
	if streamPos, size, text, err = readTextChunk(f, streamPos, size); err != nil {
		return 0, 0, "", err
	}
	return streamPos, size, text, nil
}

func readIntChunk(f *os.File, streamPos int32) (int32, int32, error) {
	var chunk []byte
	var err error
	if streamPos, chunk, err = readFromFile(f, streamPos, 4); err != nil {
		return 0, 0, err
	}
	var size int32
	buf := bytes.NewReader(chunk)
	if binary.Read(buf, binary.BigEndian, &size); err != nil {
		return 0, 0, err
	}
	return streamPos, size, nil
}

func readTextChunk(f *os.File, streamPos, size int32) (int32, int32, string, error) {
	var chunk []byte
	var err error
	if streamPos, chunk, err = readFromFile(f, streamPos, size); err != nil {
		return 0, 0, "", err
	}
	return streamPos, size, string(chunk), nil
}

func readFromFile(file *os.File, offset, size int32) (int32, []byte, error) {
	res := make([]byte, size)
	if _, err := file.ReadAt(res, int64(offset)); err != nil {
		return 0, nil, err
	}
	newStreamPos := offset + size
	return newStreamPos, res, nil
}
