package sailor

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
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
		log.Println("exec failed, err: ", err)

		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		}
	}
}

func ExecCommand(bin string, args []string) (string, error) {
	binPath, err := exec.LookPath(bin)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(binPath, args...)

	var out bytes.Buffer

	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.Trim(out.String(), "\n"), nil
}
