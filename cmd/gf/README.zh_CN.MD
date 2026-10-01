[English](README.MD) | 简体中文

# gf

`gf`是一个强大的 CLI 工具，用于便捷地构建 [GoFrame](https://goframe.org) 应用程序。

## 1. 安装

## 1）预编译二进制文件

您也可以使用预构建的二进制文件安装`gf`工具：<https://github.com/gogf/gf/releases>

1. `Mac` & `Linux`

   ```shell
    wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
   ```

   > 如果您使用`zsh`，您可能需要通过命令`alias gf=gf`重命名别名以解决`gf`和`git fetch`之间的冲突。

2. `Windows`
   手动下载，在命令行中执行，然后按照说明操作。

3. 数据库支持

   | 数据库 | 内置支持 | 说明 |
   | --- | --- | --- |
   | `mysql` | 是 | - |
   | `mariadb` | 是 | - |
   | `tidb` | 是 | - |
   | `mssql` | 是 | - |
   | `oracle` | 是 | - |
   | `pgsql` | 是 | - |
   | `sqlite` | 是 | - |
   | `sqlitecgo` | 否 | 32 位`sqlite`：在`cmd_gen_dao.go`中导入后重新构建。 |
   | `clickhouse` | 是 | - |
   | `dm` | 否 | 在`cmd_gen_dao.go`中导入后重新构建。 |
   | `talon` | 此 fork 可选支持 | 启用`talon`构建标签和 cgo；见下文。 |

## 2）手动安装

```shell
go install github.com/gogf/gf/cmd/gf/v2@latest # 最新版本
go install github.com/gogf/gf/cmd/gf/v2@v2.5.5 # 特定版本(应该 >= v2.5.5)
```

## 此 fork 的 Talon DAO 生成

此 fork 提供独立的`gf-talon`二进制文件，可从内嵌式 Talon 数据库生成 DAO、DO 和 entity 代码。在本仓库的`cmd/gf`目录下启用 cgo 构建：

```shell
CGO_ENABLED=1 go build -tags talon -o gf-talon .
```

指定 Talon 数据库目录的绝对路径，以及现有 GoFrame 应用的根目录：

```shell
./gf-talon gen dao --talonPath /absolute/path/to/talon-data -p /absolute/path/to/app
```

也可以在应用的 CLI 配置中将`gfcli.gen.dao.talonPath`设为数据库绝对路径。不要将`talonPath`与`-l`／`--link`同时使用；Talon 使用目录路径，而非数据库连接字符串。普通`gf`构建保留现有驱动与行为。

使用生成的 DAO 查询时，应用还须导入`github.com/darkmice/talon-sdk-go/goframe`以注册驱动，并在 GoFrame 数据库分组中配置`Type: "talon"`和`Name: "/absolute/path/to/talon-data"`。CLI 进程中的驱动注册不会传递到应用进程。

此构建使用 Talon SDK `v0.7.5`及其固定版本的签名原生运行时依赖`v0.1.54`。内嵌签名运行时无需设置`TALON_NATIVE_*`。若设置任何`TALON_NATIVE_*`变量，SDK 要求完整的显式信任策略；不完整时会拒绝加载。详见[Talon SDK 文档](https://github.com/darkmice/talon-sdk-go/blob/v0.7.5/README.md)。原生库契约覆盖`darwin/amd64`、`darwin/arm64`、`linux/amd64`和`linux/arm64`；其他平台不受支持。

## 2. 命令

```shell
$ gf -h
用法
    gf 命令 [选项]

命令
    up         升级项目中的 GoFrame 版本/工具到最新版本
    env        显示当前 Golang 环境变量
    fix        升级到新 GoFrame 版本后自动修复代码
    run        运行 go 代码，具有热编译功能
    gen        自动生成 dao/do/entity/pb/pbentity 的 go 文件
    tpl        模板解析和构建命令
    init       创建并初始化一个空的 GoFrame 项目
    pack       将任何文件/目录打包到资源文件或 go 文件
    build      为多个平台交叉编译 go 项目
    docker     为当前 GoFrame 项目构建 docker 镜像
    install    将 gf 二进制文件安装到系统(可能需要 root/admin 权限)
    version    显示当前二进制文件的版本信息
    doc        下载 https://pages.goframe.org/ 本地运行
    agents     将本地 AI 编码助手路径软链到`.agents`和`AGENTS.md`

选项
    -y, --yes       对所有命令都使用 yes，不再提示
    -v, --version   显示当前二进制文件的版本信息
    -d, --debug     显示内部详细的调试信息
    -h, --help      显示此命令的更多信息

附加信息
    使用 "gf 命令 -h" 获取有关命令的详细信息。
```

## 3. 常见问题

### 1）．命令`gf run`返回`pipe: too many open files`

请使用`ulimit -n 65535`扩大系统配置以增加当前终端 shell 会话的最大打开文件数，然后再运行`gf run`。
