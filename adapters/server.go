// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package adapters

import (
	"fmt"
	"net"
	"os/exec"
	"testing"
	"time"
)

func RunServer(t testing.TB, port string, path string) {
	t.Helper()
	url := fmt.Sprintf("localhost:%s", port)
	runCmd, err := compileAndRunServer(path)
	if err != nil {
		t.Fatal("could not start server:", err)
	}
	if err := waitServer(url, 5*time.Second); err != nil {
		t.Fatal("server did not start in time:", err)
	}
	t.Cleanup(func() {
		runCmd.Process.Kill()
	})
}

func compileAndRunServer(directory string) (*exec.Cmd, error) {
	err := exec.Command("go", "build", "-o", "server", directory).Run()
	if err != nil {
		return nil, err
	}
	defer exec.Command("rm", "-f", "server").Run()
	runCmd := exec.Command("./server")
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
