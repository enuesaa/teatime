//go:build ignore || mage
// +build ignore mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Gen() error {
	os.RemoveAll("./buf/gen")
	os.RemoveAll("../frontend/src/gen")

	cmd := exec.Command("buf", "generate")
	cmd.Dir = "./buf"
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf(stderr.String())
	}
	removeOmitEmpty()

	return err
}

func removeOmitEmpty() {
	entries, _ := os.ReadDir("./buf/gen/v1")
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), "pb.go") {
			fileInfo, _ := entry.Info()
			filepath := fmt.Sprintf("./buf/gen/v1/%s", entry.Name())
			file, _ := os.ReadFile(filepath)
			trimed := strings.Replace(string(file), ",omitempty", "", -1)
			
			os.WriteFile(filepath, []byte(trimed), fileInfo.Mode())
		}
	}
}

func Lint() error {
	cmd := exec.Command("revive")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	fmt.Printf(stdout.String())
	return err
}
