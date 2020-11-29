package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// registerTextFiles function will register each file in the directory tree
// recursively and will start a watcher for each file
func walkFunction(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	// if, directory exclude
	if info.IsDir() {
		return nil
	}

	// if, included in exclude directories, exclude
	parentDir := strings.Split(path, string(os.PathSeparator))[0]
	if _, ok := configs.excludeDirMap[parentDir]; ok {
		return nil
	}

	// if already included in the map, exclude
	if _, ok := configs.fileMap[path]; ok {
		return nil
	}

	go watcher(path)
	configs.fileMap[path] = 0
	return err
}

// registerFiles will register the file at the startup
func registerFiles() {
	filepath.Walk(configs.Directory, walkFunction)
}

// registerNewFiles will register the new files which are created after the program has started
func registerNewFiles() {
	for {
		time.Sleep(5 * time.Second)
		registerFiles()
	}
}
