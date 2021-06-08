package main

// Let @crossm0d know how it's going, and add in date/time tags

// [00 00 00 07] 'comment' 08 [00 00 00 57] 'Use this to diffuse the signal. In combination with delays, this gives you nice reverbs'
// [00 00 00 01]
// [00 00 00 07] 'creator' 08 [00 00 00 08] 'Polarity'
// [00 00 00 01]
// [00 00 00 0F] 'device_category' 08 [00 00 00 09] 'Container'

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

	var streamPos int32 = 0x36
	var size int32

	for {
		if size, streamPos, err = readKeyAndValue(f, streamPos); err != nil {
			return err
		}
		if size == 0 {
			break
		}
	}
	return nil
}

func readKeyAndValue(f *os.File, streamPos int32) (int32, int32, error) {
	var size int32
	var text string
	var err error

	streamPos = streamPos + getSkipSize(f, streamPos)
	if streamPos, size, text, err = readNextSizeAndChunk(f, streamPos); err == nil {
		if size == 0 {
			return 0, 0, nil
		}
		printOutput(size, streamPos, text)

		streamPos = streamPos + getSkipSize(f, streamPos)

		if streamPos, size, text, err = readNextSizeAndChunk(f, streamPos); err == nil {
			printOutput(size, streamPos, text)

			fmt.Println()

			return size, streamPos, nil
		}
	}
	return 0, 0, err
}

func getSkipSize(f *os.File, streamPos int32) int32 {
	_, bytes, _ := readFromFile(f, streamPos, 32, false)
	for _, value := range bytes {
		fmt.Printf("%02x ", value)
	}
	fmt.Printf("\n")
	for _, value := range bytes {
		if value >= 0x41 {
			fmt.Printf(".%c.", value)
		} else {
			fmt.Printf("   ")
		}
	}
	fmt.Printf("\n")
	if bytes[4] == 0 {
		return 4
	}
	return 1
}

func printOutput(size int32, streamPos int32, text string) {
	fmt.Printf("size: %x\n", size)
	fmt.Printf("stringPos: %x\n", streamPos)
	fmt.Println("text: ", text)
}

func readNextSizeAndChunk(f *os.File, streamPos int32) (int32, int32, string, error) {
	var err error
	if streamPos, size, err := readIntChunk(f, streamPos); err == nil {
		if size == 0 {
			return streamPos, 0, "", nil
		}
		if streamPos, size, text, err := readTextChunk(f, streamPos, size); err == nil {
			return streamPos, size, text, nil
		}
	}
	return 0, 0, "", err
}

func readIntChunk(f *os.File, streamPos int32) (int32, int32, error) {
	var chunk []byte
	var err error
	if streamPos, chunk, err = readFromFile(f, streamPos, 4, true); err != nil {
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
	if streamPos, chunk, err = readFromFile(f, streamPos, size, true); err != nil {
		return 0, 0, "", err
	}
	return streamPos, size, string(chunk), nil
}

func readFromFile(file *os.File, streamPos, size int32, advance bool) (int32, []byte, error) {
	res := make([]byte, size)
	if _, err := file.ReadAt(res, int64(streamPos)); err != nil {
		return 0, nil, err
	}
	newStreamPos := streamPos
	if advance {
		newStreamPos = streamPos + size
	}
	return newStreamPos, res, nil
}
