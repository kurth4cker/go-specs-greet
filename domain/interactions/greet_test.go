// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package interactions_test

import (
	"testing"

	"github.com/kurth4cker/go-specs-greet/domain/interactions"
	"github.com/kurth4cker/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
