package main

import (
	"bytes"
	"crypto/md5"
	"io"
	"log"
	"os"
	"time"
)

func reader(ch chan string, path string) {
	var checksum []byte
	h := md5.New()

	for {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		newSum := h.Sum(nil)
		h.Reset()

		if bytes.Compare(checksum, newSum) != 0 {
			checksum = newSum
			ch <- path
		}
		defer f.Close()
	}
}
