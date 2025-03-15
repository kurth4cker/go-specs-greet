// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package specifications

type CurseAdapter func(name string) string

func (c CurseAdapter) Curse(name string) (string, error) {
	return c(name), nil
}

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}
