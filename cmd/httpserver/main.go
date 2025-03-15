// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main

import (
	"log"
	"net/http"

	"github.com/kurth4cker/go-specs-greet/adapters/httpserver"
)

func main() {
	http.HandleFunc("/curse", httpserver.CurseHandler)
	http.HandleFunc("/greet", httpserver.GreetHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
