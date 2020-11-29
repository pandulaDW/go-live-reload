package main

import (
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
		if _, ok := configs.excludeDirMap[parentDir]; ok {
			return nil
		}

		go watcher(path)
		return err
	}
	filepath.Walk(configs.Directory, walkFunction)
}
