// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package interactions

import "fmt"

func Curse(name string) string {
	return fmt.Sprintf("go to hell, %s", name)
}

func Greet(name string) string {
	return fmt.Sprintf("hello %s", name)
}
