package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Starting Forgeware...")

	pythonScript := "ui/forgeware_ui.py"
	pythonExec := "python3"
	if runtime.GOOS == "windows" {
		pythonExec = "python"
	}

	cmd := exec.Command(pythonExec, pythonScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Error launching UI:", err)
	}
}
