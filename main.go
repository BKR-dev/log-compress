package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

const (
	chunkSize int = 1024
	partSize  int = 100_000_000
)

var (
	part     []byte
	count    int
	fileSize int64
)

type CountingBytesRead struct {
	io.Reader
	bytesRead int64
}

func (cB *CountingBytesRead) Read(b []byte) (int, error) {
	n, err := cB.Reader.Read(b)
	cB.bytesRead += int64(n)

	if err == nil {
		fmt.Println("Read", n, "bytes for a total of", cB.bytesRead)
	}
	return n, err
}

func openFile(name string) (byteCount int64, buffer *bytes.Buffer) {
	// read file
	data, err := os.Open(name)
	if errors.Is(err, &fs.PathError{}) {
		fmt.Println(err)
	}
	// close file
	defer func(data *os.File) {
		err = data.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(data)

	// read buffer
	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunkSize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if !errors.Is(err, io.EOF) {
		fmt.Println(err)
	} else {
		err = nil
	}
	byteCount = int64(buffer.Len())
	return
}

func main() {

	fileSize, _ := openFile("small.log")
	fmt.Printf(
		"small log:\nTotal bytes: %d\nSize in kb: %dkb\nSize in mb: %dmb\n",
		fileSize, fileSize/1000, fileSize/1_000_000)

	fileSize, _ = openFile("big.log")
	fmt.Printf(
		"big log:\nTotal bytes: %d\nSize in kb: %dkb\nSize in mb: %dmb\n",
		fileSize, fileSize/1000, fileSize/1_000_000)

	displayReadBytes()
}

func displayReadBytes() {

	file, _ := os.Open("big.log")
	defer file.Close()

	b := make([]byte, partSize)
	fmt.Println("partSize: ", partSize)

	for {
		counter := 1

		n, err := file.Read(b)
		if err == io.EOF {
			break
			// nothing more to read
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n <= partSize {
			// call gzip writer
			fmt.Println(n)
			counter++
		}

	}

}

func compressFile(data []byte, part int) (string, error) {

}
