package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// registerTextFiles function will register each file in the directory tree
// recursively and will start a watcher for each file
func registerTextFiles() {
	walkFunction := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		parentDir := strings.Split(path, string(os.PathSeparator))[0]
		if _, ok := configs.excludeDir[parentDir]; ok {
			return nil
		}

		fmt.Printf("started watching %s\n", path)
		return err
	}
	filepath.Walk(configs.directory, walkFunction)
}
