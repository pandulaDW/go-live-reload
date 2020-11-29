package main

// Config defines all the configuration settings of the application
type Config struct {
	directory  string
	excludeDir map[string]int // set as a map for quick lookups
}

// initConfig function will populate the configuration settings based on the config-go-load.json file.
// if the file is not present in the current directory it will use the default settings
func initConfig() {
	configs = Config{
		directory: ".",
	}

	// default excludeDirs, users will have to explicitly specify others
	excludeDirs := []string{".git"}

	// populate the map
	configs.excludeDir = map[string]int{}
	for _, val := range excludeDirs {
		configs.excludeDir[val] = 0
	}
}
