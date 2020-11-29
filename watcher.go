package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"
)

func watcher(path string) {
	fmt.Printf("started watching %s\n", path)
	var checksum []byte
	h := md5.New()
	initialLoad := true

	for {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		time.Sleep(1 * time.Second)

		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}

		newSum := h.Sum(nil)
		h.Reset()

		if bytes.Compare(checksum, newSum) != 0 {
			checksum = newSum
			if !initialLoad {
				ch <- path
			}
			initialLoad = false
		}
	}
}
