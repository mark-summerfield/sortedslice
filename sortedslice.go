// Copyright © 2024 Mark Summerfield. All rights reserved.
// License: GPL-3

package sortedslice

import (
    "fmt"
    _ "embed"
    )

//go:embed Version.dat
var Version string

func Hello() string {
    return fmt.Sprintf("Hello sortedslice v%s", Version)
}
