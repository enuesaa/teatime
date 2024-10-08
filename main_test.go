package main

import (
	"log"
	"os/exec"
	"testing"
	"time"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router"
)

func Testmain(m *testing.M) {
	cmd := exec.Command("docker", "compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	repos := repository.New()
	if err := repos.Startup(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	app := router.New(repos)
	go app.Start(":3000")
	time.Sleep(2 * time.Second)
	
	m.Run()

	repos.End()
}
