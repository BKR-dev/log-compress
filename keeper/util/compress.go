package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	gzip "github.com/klauspost/pgzip"
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

func doCompress() {

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

func bigOldFunc() {
	// 1024 * 1024 * 1024 (1024^3)
	// = 1 Gigglybitse
	partSize := 1_073_741_824

	PrintMemUsage()
	// displayReadBytes()
	now := time.Now()
	fmt.Printf("Starting at: %s\n", now)
	// file names
	filename := "test.log"

	// get all needed information
	partCount, lastPartSize, err := partInfo(filename, partSize)
	if err != nil {
		fmt.Println(err)
	}

	// compress files
	compressFile(filename, partCount, partSize, lastPartSize)
	fmt.Println("Finished compression")

	PrintMemUsage()
	duration := time.Since(now)
	fmt.Printf("Finished in: %s\n", duration)
}

// Returns the total count of Parts and the Size of the remainding Part
func partInfo(filename string, partSize int) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	// get filesize
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, 0, err
	}
	fileSize := int(fileInfo.Size())

	// adding the last part as it is possibly > partSize
	partCount := (fileSize / partSize)
	lastPartSize := fileSize % partSize
	if lastPartSize > 0 {
		partCount++
	}
	return partCount, lastPartSize, nil
}

func compressFile(filename string, partCount, partSize, lastPartSize int) {

	// buffer size of 1G
	buf1G := make([]byte, partSize)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	PrintMemUsage()

	for i := 0; i < partCount; i++ {

		// create compressed file
		compfileName := compFileNameParts(filename, i+1)
		compFile, err := os.Create(compfileName)
		if err != nil {
			fmt.Println(err)
		}
		defer compFile.Close()

		gWriter := gzip.NewWriter(compFile)

		// there could be a better way of doing this
		offset, err := file.Seek(int64(partSize)*int64((i+1)), 0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Set offset relative to file size: %d\n", offset)

		_, err = io.ReadFull(file, buf1G)
		if err != nil {
			fmt.Println(err)
		}

		bytesWritten, err := gWriter.Write(buf1G)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("bytes compressed: %v\n", bytesWritten)

		defer gWriter.Close()
	}

	PrintMemUsage()

}

func compFileNameParts(filename string, part int) string {
	file, _, _ := strings.Cut(filename, ".")
	return file + "." + strconv.Itoa(part) + ".gz"
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v GiB", bToGb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v GiB", bToGb(m.TotalAlloc))
	fmt.Printf("\tSys = %v GiB", bToGb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToGb(b uint64) uint64 {
	return b / 1024 / 1024 / 1024
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// TODO: Reverse it
func decompressFile(compFilename, decompressFilename string) {

}
