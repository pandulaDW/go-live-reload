package main

import (
	"bufio"
	"log"
	"os/exec"
)

func readStdInLinux() {
	cmd := exec.Command("python", "program.py")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Println(in.Text())
	}

	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
}
