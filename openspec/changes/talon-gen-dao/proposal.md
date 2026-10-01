## Why

The standalone `gf` CLI does not inherit driver registrations from an application. Its `gen dao -l` option creates a server-style Link configuration, while the Talon GoFrame adapter requires a native database directory in `ConfigNode.Name`.

## What Changes

- Add an explicit `talonPath` generator input and reject incompatible `link` input.
- Register the released Talon GoFrame adapter only in an opt-in `talon` and cgo build.
- Document the build, invocation, trust requirements, and supported platforms in both CLI readmes.
- Verify path handling, generated code, and a representative DAO query against the signed embedded runtime.

## Impact

The ordinary CLI build and existing database generator paths keep their dependencies and behavior. The opt-in binary depends on `github.com/darkmice/talon-sdk-go` v0.7.5 and its signed native runtime.
