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

// copy file to sandbox keeping permission
func copy_file_to_jail(src_exe_path string, jailed_command_path string) error {
	src_command_fd, err := os.Open(src_exe_path)

	if err != nil { return err }

	defer src_command_fd.Close()

	jailed_command_fd, err := os.Create(jailed_command_path)

	if err != nil { return err }

	defer jailed_command_fd.Close()

	src_command_fd_info, err := src_command_fd.Stat()

	if err != nil { return err }

	// change jailed command to exactly the mod of
	jailed_command_fd.Chmod(src_command_fd_info.Mode())

	_, err = io.Copy(jailed_command_fd, src_command_fd)

	return err
}

func sandbox_jail(command string, sandbox_path string) (jailed_command string, err error) {
	// search command executeable by looking at PATH
	src_exe_path, err := exec.LookPath(command)
	if err != nil { return "", err }

	// get command path
	jailed_command_path := filepath.Join(sandbox_path, src_exe_path)

	// make a copy dir in snadbox
	err = os.MkdirAll(filepath.Dir(jailed_command_path), 0755)
	if err != nil { return "", err }

	// copy executable to sandbox
	err = copy_file_to_jail(src_exe_path, jailed_command_path)

	if err != nil { return "", err }

	return src_exe_path, nil
}

func run_container(command string, args []string, sandbox_path string) {
	cmd := exec.Command(command, args...)

	cmd.Stdin = nullReader{}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Printf("Unable to Start Command1\n")
	// }

	err := syscall.Chroot(sandbox_path)
	if err != nil {
		fmt.Printf("Unable to sandbox command!\n")
		os.Exit(1)
	}

	err = syscall.Chdir("/")
	if err != nil {
		fmt.Printf("Unable to chdir to sandbox root!\n")
		os.Exit(1)
	}

	// process isolation
	cmd.SysProcAttr = &syscall.SysProcAttr {
		// flag for clone call(Linux only)
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
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