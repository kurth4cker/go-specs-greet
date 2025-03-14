// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/kurth4cker/go-specs-greet/adapters"
	"github.com/kurth4cker/go-specs-greet/adapters/httpserver"
	"github.com/kurth4cker/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port    = "8080"
		path    = "."
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver  = httpserver.Driver{
			BaseURL: baseURL,
			Client: &http.Client{
				Timeout: 1 * time.Second,
			},
		}
	)

	adapters.RunServer(t, port, path)
	specifications.GreetSpecification(t, driver)
}
