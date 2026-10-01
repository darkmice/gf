## Implementation

- [x] Add explicit Talon source validation and opt-in driver registration.
- [x] Add focused unit tests and validate ordinary and tagged CLI builds.
- [x] Generate from a signed Talon database, compile generated code, and run one DAO query.
- [x] Update English and Chinese CLI documentation and review the full change.

## Verification

- `go test ./internal/cmd/gendao -count=1` and the same test with `CGO_ENABLED=1 go test -tags talon` passed.
- The ordinary CLI and `CGO_ENABLED=1 go build -tags talon` builds passed. `CGO_ENABLED=0 go build -tags talon` retains the explicit unavailable-driver error.
- The Talon-enabled CLI generated DAO, DO, and entity code from a database opened through the pinned signed runtime. An example GoFrame application built and its generated DAO returned `id=1 name=ready`.
- `go test ./internal/cmd -run '^Test_Gen_Dao_Sqlite3$' -count=1` passed.
- A full staged-file review found no outstanding code or specification issues. `make lint` could not run because `golangci-lint` is not installed.
