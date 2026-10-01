## ADDED Requirements

### Requirement: Explicit native Talon source

The DAO generator SHALL accept an absolute `talonPath` input and construct a Talon database configuration using `Type` and `Name`, without interpreting it as a Link connection string.

#### Scenario: Valid Talon path

- **WHEN** an existing absolute Talon database directory is provided to a Talon-enabled CLI
- **THEN** the generator SHALL discover tables and fields and generate DAO, DO, and entity code

#### Scenario: Invalid or conflicting input

- **WHEN** `talonPath` is relative, inaccessible, or combined with `link`
- **THEN** the CLI SHALL report a specific input error before generating code

### Requirement: Opt-in driver

The standard CLI SHALL retain its existing driver set. A CLI built with cgo and the `talon` tag SHALL register the released Talon GoFrame adapter.

#### Scenario: Standard binary

- **WHEN** a standard CLI receives `talonPath`
- **THEN** it SHALL report how to build the Talon-enabled CLI
