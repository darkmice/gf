//go:build !talon || !cgo

// This file reports that the optional native Talon driver is absent.

package gendao

// talonDriverEnabled reports whether the tagged native driver is registered.
func talonDriverEnabled() bool { return false }
