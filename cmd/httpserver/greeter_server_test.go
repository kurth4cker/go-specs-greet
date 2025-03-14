// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main_test

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"testing"
	"time"

	gospecsgreet "github.com/kurth4cker/go-specs-greet"
	"github.com/kurth4cker/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	// ctx := context.Background()

	runCmd, err := compileAndRunServer(".")
	if err != nil {
		t.Fatal("could not start server:", err)
	}
	if err := waitServer("localhost:8080", 5*time.Second); err != nil {
		t.Fatal("server did not start in time:", err)
	}
	t.Cleanup(func() {
		runCmd.Process.Kill()
	})

	driver := gospecsgreet.Driver{
		BaseURL: "http://localhost:8080",
		Client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
	specifications.GreetSpecification(t, driver)
}

func compileAndRunServer(directory string) (*exec.Cmd, error) {
	buildCmd := exec.Command("go", "build", "-o", "httpserver", directory)
	if err := buildCmd.Run(); err != nil {
		return nil, err
	}
	defer exec.Command("go", "clean", ".").Run()

	runCmd := exec.Command("./httpserver")
	if err := runCmd.Start(); err != nil {
		return nil, err
	}
	return runCmd, nil
}

func waitServer(address string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.Dial("tcp", address)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("server did not start within %v", timeout)
}
