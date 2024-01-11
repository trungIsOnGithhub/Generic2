package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type nullReader struct {}

func (nullReader) Read(p []byte) (int, error) {
	return len(p), nil
}

func sandbox_jail(command string, sandbox_path string) (jailed_command string, err error) {
	// search command executeable by looking at PATH
	src_exe_path, err := exec.LookPath(command)

	if err != nil { return "", err }

	jailed_command_path := filepath.Join(sandbox_path, src_exe_path)

	err = os.MkdirAll(filepath.Dir(jailed_command_path), 0755)

	if err != nil { return "", err }

	err = copy_file_to_jail(srcCommandPath, sandboxedCommandPath)

	if err != nil { return "", err }

	return srcCommandPath, nil
}

func run_container(command string, args []string, sandbox_path string) {
	cmd := exec.Command(command, args...)

	cmd.Stdin = nullReader{}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	err := syscall.Chroot(sandbox_path)

	if err != nil {
		fmt.Printf("Unable to sandbox command")
		os.Exit(1)
	}

	err = syscall.Chdir("/")

	if err != nil {
		fmt.Printf("Unable to chdir to sandbox root!\n")
		os.Exit(1)
	}

	err = cmd.Run()

	if err != nil {
		fmt.Printf("Error Running Container!\n")
		os.Exit(cmd.ProcessState.ExitCode())
	}
}

// Usage: my_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	// args := os.Args[4:len(os.Args)]

	// fmt.Printf("argument 3: %s\n", command)
	// fmt.Printf("remained arguments : %s\n", args)

	jail_dir, err := os.MkdirTemp("/tmp", "my-docker")

	if err != nil {
		fmt.Printf("Error Creating Jail Director!\n")
		os.Exit(1)
	}

	jailed_command, err := sandbox_jail(command, jail_dir)

	if err != nil {
		fmt.Printf("Error jailing command!\n")
		os.Exit(1)
	}

	defer os.RemoveAll(jail_dir)
}
