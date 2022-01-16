package internal

import (
	"bufio"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// Code inspired from https://www.yellowduck.be/posts/reading-command-output-line-by-line/
func shell(command string) error {
	cmd := exec.Command("sh", "-c", command)

	r, _ := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	done := make(chan struct{})

	scanner := bufio.NewScanner(r)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			log.Debug(line)
		}
		done <- struct{}{}
	}()

	err := cmd.Start()

	if err != nil {
		return err
	}

	cmd.Wait()

	return nil
}
