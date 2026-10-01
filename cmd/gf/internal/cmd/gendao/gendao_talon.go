// This file validates the explicit Talon database path used by DAO generation.

package gendao

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/database/gdb"
)

// talonDriverType identifies the native GoFrame database adapter.
const talonDriverType = "talon"

// newTalonDatabase creates the DAO generator's native database connection.
func newTalonDatabase(path, link string) (gdb.DB, error) {
	if link != "" {
		return nil, fmt.Errorf("talonPath cannot be combined with link")
	}
	if !filepath.IsAbs(path) {
		return nil, fmt.Errorf("talonPath must be an absolute database path: %q", path)
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("talonPath %q is not accessible: %w", path, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("talonPath %q must be a database directory", path)
	}
	if !talonDriverEnabled() {
		return nil, fmt.Errorf("Talon driver is unavailable; rebuild gf with CGO_ENABLED=1 and -tags talon")
	}
	return gdb.New(gdb.ConfigNode{Type: talonDriverType, Name: path})
}
