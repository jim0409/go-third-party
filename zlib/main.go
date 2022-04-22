package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main() {
	var s = "hello, world\n"
	var bs = []byte(s)
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	// w.Write([]byte(s))
	w.Write(bs)
	w.Close()

	r, err := zlib.NewReader(&b)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, r)
	r.Close()

	fmt.Printf("s: %T, %d\n", s, unsafe.Sizeof(s))
	fmt.Printf("bs: %T, %d\n", s, unsafe.Sizeof(bs))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("w: %T, %d\n", w, unsafe.Sizeof(w))
	fmt.Printf("r: %T, %d\n", r, unsafe.Sizeof(r))
}
