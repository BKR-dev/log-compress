package main

import (
	"bufio"
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

	file, _ := os.Open(filename)
	defer file.Close()

	compFile, err := os.Create(compFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer compFile.Close()

	bR := bufio.NewReader(file)
	data, err := io.ReadAll(bR)
	if err != nil {
		fmt.Println(err)
	}

	b := make([]byte, len(data))

	_, err = bR.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	gWriter := pgzip.NewWriter(compFile)
	defer gWriter.Close()
	PrintMemUsage()
	_, err = gWriter.Write(b)
	if err != nil {
		fmt.Println(err)
	}

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// TODO: Reverse it
func decompressFile(compFilename, decompressFilename string) {

}
