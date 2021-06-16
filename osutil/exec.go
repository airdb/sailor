package osutil

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
