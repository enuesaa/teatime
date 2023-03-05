// +build ignore mage

package main

import (
	"os"
	"os/exec"
)

func Gen() error {
	os.RemoveAll("./buf/gen")
	os.RemoveAll("../frontend/gen")

	cmd := exec.Command("buf", "generate")
    cmd.Dir = "./buf"
	return cmd.Run()
}
