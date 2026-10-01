// This file tests validation and construction of the DAO generator's Talon connection.

package gendao

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

// TestNewTalonDatabase validates Talon input before touching the native runtime.
func TestNewTalonDatabase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := t.TempDir()
		_, err := newTalonDatabase(path, "mysql:example")
		t.Assert(err != nil, true)
		t.Assert(strings.Contains(err.Error(), "cannot be combined"), true)

		_, err = newTalonDatabase("relative/path", "")
		t.Assert(err != nil, true)
		t.Assert(strings.Contains(err.Error(), "absolute"), true)

		_, err = newTalonDatabase(filepath.Join(path, "missing"), "")
		t.Assert(err != nil, true)
		t.Assert(strings.Contains(err.Error(), "not accessible"), true)
		filePath := filepath.Join(path, "file")
		err = os.WriteFile(filePath, nil, 0600)
		t.AssertNil(err)
		_, err = newTalonDatabase(filePath, "")
		t.Assert(err != nil, true)
		t.Assert(strings.Contains(err.Error(), "database directory"), true)

		db, err := newTalonDatabase(path, "")
		if !talonDriverEnabled() {
			t.Assert(err != nil, true)
			t.Assert(strings.Contains(err.Error(), "-tags talon"), true)
			return
		}
		t.AssertNil(err)
		t.Assert(db.GetConfig().Type, talonDriverType)
		t.Assert(db.GetConfig().Name, path)
		t.Assert(db.GetConfig().Link, "")
	})
}
