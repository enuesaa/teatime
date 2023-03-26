//go:build ignore || mage
// +build ignore mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func Gen() error {
	os.RemoveAll("./buf/gen")
	os.RemoveAll("../frontend/gen")

	cmd := exec.Command("buf", "generate")
	cmd.Dir = "./buf"
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf(stderr.String())
	}
	return err
}

func Lint() error {
	cmd := exec.Command("revive")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	fmt.Printf(stdout.String())
	return err
}
