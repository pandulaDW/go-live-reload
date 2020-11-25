package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "program.py")

	file, _ := os.OpenFile("logs.out", os.O_CREATE|os.O_APPEND, 0666)

	cmd.Stderr = file

	cmd.Run()
}
