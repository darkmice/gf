English | [简体中文](README.zh_CN.MD)

# gf

`gf` is a powerful CLI tool for building [GoFrame](https://goframe.org) application with convenience.

## 1. Install

## 1) PreCompiled Binary

You can also install `gf` tool using pre-built binaries: <https://github.com/gogf/gf/releases>

1. `Mac` & `Linux`

   ```shell
    wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
   ```

   > If you're using `zsh`, you might need rename your alias by command `alias gf=gf` to resolve the conflicts between `gf` and `git fetch`.

2. `Windows`
   Manually download, execute in command line it and then follow the instruction.

3. Database support

   | DB | builtin support | remarks |
   | --- | --- | --- |
   | `mysql` | yes | - |
   | `mariadb` | yes | - |
   | `tidb` | yes | - |
   | `mssql` | yes | - |
   | `oracle` | yes | - |
   | `pgsql` | yes | - |
   | `sqlite` | yes | - |
   | `sqlitecgo` | no | 32-bit `sqlite`: import in `cmd_gen_dao.go` and rebuild. |
   | `clickhouse` | yes | - |
   | `dm` | no | Import in `cmd_gen_dao.go` and rebuild. |
   | `talon` | optional in this fork | Build with the `talon` tag and cgo; see below. |

## 2) Manually Install

```shell
go install github.com/gogf/gf/cmd/gf/v2@latest # latest version
go install github.com/gogf/gf/cmd/gf/v2@v2.5.5 # certain version(should be >= v2.5.5)
```

## Talon DAO generation in this fork

This fork provides a separate `gf-talon` binary for generating DAO, DO, and entity code from an embedded Talon database. From this repository's `cmd/gf` directory, build it with cgo enabled:

```shell
CGO_ENABLED=1 go build -tags talon -o gf-talon .
```

Use an absolute Talon database directory and an existing GoFrame application root:

```shell
./gf-talon gen dao --talonPath /absolute/path/to/talon-data -p /absolute/path/to/app
```

Alternatively, set `gfcli.gen.dao.talonPath` to the absolute database path in the application's CLI configuration. Do not combine `talonPath` with `-l`/`--link`; Talon uses a directory path rather than a database connection link. The ordinary `gf` build retains its existing drivers and behavior.

To query through the generated DAO, the application must also import `github.com/darkmice/talon-sdk-go/goframe` for driver registration and configure its GoFrame database group with `Type: "talon"` and `Name: "/absolute/path/to/talon-data"`. The CLI's driver registration does not carry into the application process.

This build uses Talon SDK `v0.7.5` and its pinned signed native runtime dependency `v0.1.54`. The signed runtime works without `TALON_NATIVE_*` settings. If any `TALON_NATIVE_*` setting is supplied, the SDK requires a complete explicit trust policy and fails closed on a partial policy. See the [Talon SDK documentation](https://github.com/darkmice/talon-sdk-go/blob/v0.7.5/README.md). The native bundle contract covers `darwin/amd64`, `darwin/arm64`, `linux/amd64`, and `linux/arm64`; other platforms are unsupported.

## 2. Commands

```shell
$ gf -h
USAGE
    gf COMMAND [OPTION]

COMMAND
    up         upgrade GoFrame version/tool to latest one in current project
    env        show current Golang environment variables
    fix        auto fixing codes after upgrading to new GoFrame version
    run        running go codes with hot-compiled-like feature
    gen        automatically generate go files for dao/do/entity/pb/pbentity
    tpl        template parsing and building commands
    init       create and initialize an empty GoFrame project
    pack       packing any file/directory to a resource file, or a go file
    build      cross-building go project for lots of platforms
    docker     build docker image for current GoFrame project
    install    install gf binary to system (might need root/admin permission)
    version    show version information of current binary
    doc        download https://pages.goframe.org/ to run locally
    agents     link local AI coding-agent paths to .agents and AGENTS.md

OPTION
    -y, --yes       all yes for all command without prompt ask
    -v, --version   show version information of current binary
    -d, --debug     show internal detailed debugging information
    -h, --help      more information about this command

ADDITIONAL
    Use "gf COMMAND -h" for details about a command.
```

## 3. FAQ

### 1). Command `gf run` returns `pipe: too many open files`

Please use `ulimit -n 65535` to enlarge your system configuration for max open files for current terminal shell session, and then `gf run`.
