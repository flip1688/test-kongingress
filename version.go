package main

import (
	"fmt"
)

var (
	// COMMIT git commit hash
	COMMIT = "UNKNOW"
)

// Version print version info
func Version() string {
	return fmt.Sprintf(`
	----------
	Commit: %v
	`, COMMIT)
}
