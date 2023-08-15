package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	pgzip "github.com/klauspost/pgzip"
)

func main() {

	PrintMemUsage()
	// displayReadBytes()
	now := time.Now()
	fmt.Printf("Starting at: %s\n", now)
	filename := "test.log"
	compFilename := "test.gz"
	decompressFilename := "decompressed.test"
	compressFile(filename, compFilename)
	fmt.Println("Finished compression")
	// 30 475 158
	decompressFile(compFilename, decompressFilename)
	duration := time.Since(now)
	fmt.Printf("Finished in: %s\n", duration)
	PrintMemUsage()
}

func compressFile(filename, compFilename string) {

	// 1024 * 1024 * 1024 (1024^3)
	// = 1 Gigglybitse
	partSize := 1_073_741_824

	// buffer size of 1G
	buf1G := make([]byte, partSize)

	// open log file
	file, _ := os.Open(filename)
	defer file.Close()
	// create compressed file
	compFile, err := os.Create(compFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer compFile.Close()

	// get filesize
	fI, _ := os.Stat(filename)
	fileSize := int(fI.Size())
	// read "first" 1G into buffer
	d1Size, _ := io.ReadFull(file, buf1G)
	// subtract copied bytes from filesize
	fileSize -= d1Size
	// making it easy for me to keep track
	fmt.Printf("fileSize: %d\npartSize: %d\nbufSize: %d\n", fileSize, d1Size, len(buf1G))

	PrintMemUsage()

	gWriter := pgzip.NewWriter(compFile)
	defer gWriter.Close()

	_, err = gWriter.Write(buf1G)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = file.Seek(int64(partSize), 0)

	d1Size, _ = io.ReadFull(file, buf1G)
	// subtract copied bytes from filesize
	fileSize -= d1Size
	// making it easy for me to keep track
	fmt.Printf("fileSize: %d\npartSize: %d\nbufSize: %d\n", fileSize, d1Size, len(buf1G))

	PrintMemUsage()

	compFilename = "test2.gz"

	compFile, err = os.Create(compFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer compFile.Close()
	gWriter = pgzip.NewWriter(compFile)
	defer gWriter.Close()

	_, err = gWriter.Write(buf1G)
	if err != nil {
		fmt.Println(err)
	}

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
