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
