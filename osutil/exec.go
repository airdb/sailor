package osutil

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Exec(bin string, args []string) {
	binPath, err := exec.LookPath(bin)
	if err != nil {
		return
	}

	cmd := exec.Command(binPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func ExecCommand(bin string, args []string) (string, error) {
	binPath, err := exec.LookPath(bin)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(binPath, args...)

	var out strings.Builder

	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.Trim(out.String(), "\n"), nil
}

// const CommandTimeout = 10

func CommandContext(timeout int, binPath string, args []string) (string, error) {
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmd := exec.CommandContext(ctx, binPath, args...)

	// This time we can simply use Output() to get the result.
	out, err := cmd.Output()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Command timed out")
		return "", ctx.Err()
	}

	if err != nil {
		log.Println("Non-zero exit code:", err)
	}

	return strings.Trim(string(out), "\n"), nil
}

// Install command binary.
func Install(tmpPath, binPath string) error {
	defer os.Remove(tmpPath)

	err := exec.CommandContext(context.Background(), "/usr/bin/install", tmpPath, binPath).Run()
	if err != nil {
		log.Println(err)

		return err
	}

	return nil
}

func SendReloadSignal() error {
	ppid := strconv.Itoa(os.Getppid())
	err := exec.CommandContext(context.Background(), "/bin/kill", "-HUP", ppid).Run()

	return err
}
