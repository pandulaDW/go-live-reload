package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config defines all the configuration settings of the application
type Config struct {
	Directory     string
	ExcludeDir    []string // set as a map for quick lookups
	excludeDirMap map[string]int
}

// initConfig function will populate the configuration settings based on the config-go-load.json file.
// if the file is not present in the current directory it will use the default settings
func initConfig() {
	// setting default configs
	configs = Config{
		Directory:  ".",
		ExcludeDir: []string{".git"},
	}

	// load the config file
	content, err := ioutil.ReadFile("config-go-load.json")

	// panic if there's an error reading the file
	if !os.IsNotExist(err) && err != nil {
		panic(err)
	}

	// if the file exists, populate the config using its values
	if content != nil {
		if err := json.Unmarshal(content, &configs); err != nil {
			panic(err)
		}
	}

	// populate the exclude dir map
	populateExcludeDir(&configs)
}

// helper functions
func populateExcludeDir(configs *Config) {
	configs.excludeDirMap = map[string]int{}

	for _, val := range configs.ExcludeDir {
		configs.excludeDirMap[val] = 0
	}
}
