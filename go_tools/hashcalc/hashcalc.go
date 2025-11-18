package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hashcalc <file>")
		return
	}
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	md5h := md5.New()
	sha1h := sha1.New()
	sha256h := sha256.New()

	io.Copy(io.MultiWriter(md5h, sha1h, sha256h), file)

	fmt.Printf("MD5: %x\nSHA1: %x\nSHA256: %x\n", md5h.Sum(nil), sha1h.Sum(nil), sha256h.Sum(nil))
}
