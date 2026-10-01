## Decision

`gf gen dao --talonPath /absolute/database/directory` constructs `gdb.ConfigNode{Type: "talon", Name: path}`. A `talon && cgo` source file imports the released adapter. A fallback file reports that Talon support is unavailable in ordinary or cgo-disabled builds. The command validates the absolute, accessible path and forbids `link` in the same input before generation.

## Rationale

The adapter already owns native schema discovery, runtime verification, and SQL behavior. The CLI should only select its connection node. The tagged import avoids making the default binary load a native adapter.

## Risks and Verification

The native runtime supports only the platforms published by the pinned module. Verify both standard and tagged builds, focused path tests, real schema discovery and generation, generated-code compilation, and one DAO read. Keep generated files outside the repository.
