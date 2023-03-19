// +build ignore mage

package main

import (
	"os"
	"os/exec"
	"fmt"
	"bytes"
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
