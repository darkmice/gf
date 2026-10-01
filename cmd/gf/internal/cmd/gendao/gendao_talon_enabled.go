//go:build talon && cgo

// This file registers the native Talon driver in an explicitly tagged CLI build.

package gendao

import _ "github.com/darkmice/talon-sdk-go/goframe"

// talonDriverEnabled reports whether the tagged native driver is registered.
func talonDriverEnabled() bool { return true }
