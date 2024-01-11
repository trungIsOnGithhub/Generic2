package main

import (
	// 	"fmt"
	"os"
	"os/exec"
)

// Usage: my_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	// args := os.Args[4:len(os.Args)]

	// fmt.Printf("argument 3: %s\n", command)
	// fmt.Printf("remained arguments : %s\n", args)
	
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		// fmt.Printf("Err: %v", err)
		os.Exit(cmd.ProcessState.ExitCode())
	}
	
	// fmt.Println(string(output))
}
