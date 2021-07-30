package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("https://www.jianshu.com/p/cc1ffa5a3f4d", qrcode.Medium, 256, "./golang_qrcode.png")
}
