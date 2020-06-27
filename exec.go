package sailor

import (
	"log"
	"os"
	"os/exec"
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

		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}
}
