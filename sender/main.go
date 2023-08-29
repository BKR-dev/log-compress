package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// dankStuffFoundOnline()
	// readFromStdOut()
	anotherDangThing()
}

func readFromStdOut() []byte {
	buf := make([]byte, 1024)
	os.Stdout.Seek(1024, 2)
	_, err := io.ReadFull(os.Stdout, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)
	return buf
}

func dankStuffFoundOnline() {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = oldOut
	out := <-outC
	fmt.Println(out)

}

func anotherDangThing() {
	stdOutReader := os.Stdout
	scanner := bufio.NewScanner(stdOutReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Bytes())
		}
		done <- true
	}()

	fmt.Println("is done: ", <-done)
}

func somethingElse(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
