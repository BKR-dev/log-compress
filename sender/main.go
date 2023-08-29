package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	dankStuffFoundOnline()
	// readFromStdOut()

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
