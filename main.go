package main

import (
	"fmt"
	"os/exec"
)

// define a global variables for the application
var configs Config
var ch chan string

func main() {

	// initialize the channel
	ch = make(chan string)

	// populate the config
	initConfig()

	// register the initial file handlers synchronously
	registerFiles()

	// register new file handlers, asynchronously
	go registerNewFiles()

	// define the command
	cmd := exec.Command("node", "app/app.js")

	// start the command
	cmd.Start()

	// wait for file changes
	for {
		changedFile := <-ch
		fmt.Printf("change detected in %s\nProgram restarting...\n", changedFile)

		// // Restart the program
		// cmd.Process.Kill()
		// time.Sleep(300 * time.Millisecond)
		// cmd.Start()
	}
}
