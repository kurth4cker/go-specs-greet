// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main_test

import (
	"fmt"
	"testing"

	"github.com/kurth4cker/go-specs-greet/adapters"
	"github.com/kurth4cker/go-specs-greet/adapters/grpcserver"
	"github.com/kurth4cker/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "50051"
		path   = "."
		driver = grpcserver.Driver{
			Addr: fmt.Sprintf("localhost:%s", port),
		}
	)

	adapters.RunServer(t, port, path)
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
