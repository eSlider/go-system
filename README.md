# go-system

[![Go Reference](https://pkg.go.dev/badge/github.com/eslider/go-system.svg)](https://pkg.go.dev/github.com/eslider/go-system)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://go.dev)
[![Latest Release](https://img.shields.io/github/v/tag/eSlider/go-system?sort=semver&label=release)](https://github.com/eSlider/go-system/releases)
[![GitHub Stars](https://img.shields.io/github/stars/eSlider/go-system?style=social)](https://github.com/eSlider/go-system/stargazers)

Go library providing OS-level utilities: environment/config loading, shell command execution, file checksums, and debugger detection.

## Installation

```bash
go get github.com/eslider/go-system
```

## Features

- Load `.env` files and YAML configs with environment-based sections
- Execute shell commands with captured stdout/stderr and exit codes
- SHA-256 file checksums for multipart uploads
- Parse file paths into components (name, basename, extension, directory)
- Detect if running under the Delve debugger

## Quick Start

### Load Environment Variables

```go
system.LoadEnvs("/path/to/config")
dbURL := os.Getenv("DATABASE_URL")
```

### Read YAML Config

```go
type AppConfig struct {
    Port     int    `mapstructure:"port"`
    Database string `mapstructure:"database"`
}

var cfg AppConfig
system.ReadConfig("config.yml", "production", &cfg)
```

### Execute Shell Commands

```go
result, err := system.Exec("git", "status", "--short")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.StdOut)
```

### File Path Info

```go
info := system.GetFileInfo("/data/reports/q4-2025.csv")
// info.Name     = "q4-2025.csv"
// info.BaseName = "q4-2025"
// info.Ext      = ".csv"
// info.Dir      = "/data/reports/"
```

## API

| Function | Description |
|---|---|
| `LoadEnvs(path)` | Load `.env` and `.env.default` files |
| `ReadConfig(path, env, ptr)` | Read YAML config for environment |
| `Exec(args...)` | Run shell command, capture output |
| `CheckSum(file)` | SHA-256 hex digest of multipart file |
| `GetFileInfo(path)` | Parse file path components |
| `IsFileExists(path)` | Check if file exists (not directory) |
| `IsLaunchedByDebugger()` | Detect Delve debugger |

## License

[MIT](LICENSE)
