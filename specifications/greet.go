// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("kurth4cker")
	assert.NoError(t, err)
	assert.Equal(t, "hello kurth4cker", got)
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("Chris")
	assert.NoError(t, err)
	assert.Equal(t, "go to hell, Chris", got)
}
