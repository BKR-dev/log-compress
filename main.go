package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// 1024 * 1024 * 1024 (1024^3)
	// = 1 Gigglybitse
	partSize := 1_073_741_824

	PrintMemUsage()
	// displayReadBytes()
	now := time.Now()
	fmt.Printf("Starting at: %s\n", now)
	filename := "test.log"
	compFilename := "test.gz"

	partCount, _, _ := partInfo(filename, partSize)
	fmt.Println("Partcount: ", partCount)
	duration := time.Since(now)
	fmt.Printf("Finished in: %s\n", duration)
	os.Exit(1)

	decompressFilename := "decompressed.test"
	// compressFile(filename, compFilename)
	fmt.Println("Finished compression")
	// 30 475 158
	decompressFile(compFilename, decompressFilename)
	PrintMemUsage()
}

func fileInformation(filename string) (int, error) {
	// open log file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	// get filesize
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}

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

	partCount := fileSize / partSize
	lastPartSize := fileSize % partSize
	fmt.Printf("Dividing fileSize %d by partSize %d: %d\n", fileSize, partSize, partCount)
	fmt.Println("Last Part is of Size ", lastPartSize)

	return partCount, lastPartSize, nil
}

// func compressFile(filename, compFilename string) {

// 	// buffer size of 1G
// 	buf1G := make([]byte, partSize)

// 	// create compressed file
// 	compFile, err := os.Create(compFilename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer compFile.Close()

// 	// read "first" 1G into buffer
// 	d1Size, _ := io.ReadFull(file, buf1G)
// 	// subtract copied bytes from filesize
// 	fileSize -= d1Size
// 	// making it easy for me to keep track
// 	// fmt.Printf("fileSize: %d\npartSize: %d\nbufSize: %d\n", fileSize, d1Size, len(buf1G))

// 	PrintMemUsage()

// 	gWriter := pgzip.NewWriter(compFile)
// 	defer gWriter.Close()

// 	_, err = gWriter.Write(buf1G)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// _, _ = file.Seek(int64(partSize), 0)

// 	// d1Size, _ = io.ReadFull(file, buf1G)
// 	// subtract copied bytes from filesize
// 	// fileSize -= d1Size
// 	// making it easy for me to keep track
// 	// fmt.Printf("fileSize: %d\npartSize: %d\nbufSize: %d\n", fileSize, d1Size, len(buf1G))

// 	PrintMemUsage()

// }

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
