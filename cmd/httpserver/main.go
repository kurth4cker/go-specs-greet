// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main

import (
	"log"
	"net/http"

	gospecsgreet "github.com/kurth4cker/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(gospecsgreet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
