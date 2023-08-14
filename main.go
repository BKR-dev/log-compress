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
	// = Gigglybitse
	partSize := 1_073_741_824

	buf := make([]byte, partSize)

	file, _ := os.Open(filename)
	defer file.Close()

	fI, _ := os.Stat(filename)

	fileSize := int(fI.Size())
	d1Size, _ := io.ReadFull(file, buf)
	fileSize -= d1Size

	fmt.Printf("fileSize: %d\npartSize: %d\nbufSize: %d\n", fileSize, d1Size, len(buf))

	os.Exit(1)

	compFile, err := os.Create(compFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer compFile.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	gWriter := pgzip.NewWriter(compFile)
	defer gWriter.Close()

	_, err = gWriter.Write(data)
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
