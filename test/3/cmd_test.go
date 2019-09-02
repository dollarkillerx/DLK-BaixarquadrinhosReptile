package cmd_test

import (
	"log"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd, err := exec.Command("bash", "echo","hello").Output()
	if err != nil {
		panic(err)
	}
	log.Print(cmd)
}
