package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	writeByte()
	writeBytes()
	writeString()
	writeRune()
	writeTo()
}

func newBuffer() {
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))

	fmt.Println(buf1)
	fmt.Println(buf2)
}

// Write—- func (b *Buffer) Write(p []byte) (n int, err error)
func writeBytes() {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	s := []byte(" world")
	buf.Write(s)
	fmt.Println(buf.String())
}

// WriteString—- func (b *Buffer) WriteString(s string) (n int, err error)
func writeString() {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	s := " world"
	buf.WriteString(s)
	fmt.Println(buf.String())
}

// WriteByte—- func (b *Buffer) WriteByte(c byte) error
func writeByte() {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	var s byte = '!'
	buf.WriteByte(s)
	fmt.Println(buf.String())
}

// WriteRune—- func (b *Buffer) WriteRune(r rune) (n int, err error)
func writeRune() {
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	var s rune = '好'
	buf.WriteRune(s)
	fmt.Println(buf.String())
}

// WriteTo—- func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
func writeTo() {
	buf := bytes.NewBufferString("hello")
	buf.WriteTo(os.Stdin)
	fmt.Fprintln(os.Stdin, buf.String())
}
