<div align="center">

  [![en](https://img.shields.io/badge/lang-en-green.svg)](https://github.com/DenisKozarezov/interpreter/blob/master/README.md)
  [![ru](https://img.shields.io/badge/lang-ru-red.svg)](https://github.com/DenisKozarezov/interpreter/blob/master/README-ru.md)

  [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/DenisKozarezov/interpreter.svg)](https://github.com/DenisKozarezov)

  <h1>Scripting Language Interpreter</h1>

</div>

## Introduction

> [!WARNING]
> This repository is a pet project and is purely for academic interest. The scripting language being developed
> here (the so-called **guest language**) is not a real programming language. The goals of the project are
> to study the topic of source program recognition and its subsequent lexing, parsing and execution of an abstract
> syntax tree (AST).

## Installation

### Install and Set Up `ipret` on Linux <img src="https://logo.svgcdn.com/d/linux-original.png" width=25 height=25>

#### Install `ipret` binary via wget from *GitHub*

1. Open the [Releases](https://github.com/DenisKozarezov/interpreter/releases) section and select the needed version of the `ipret` utility.
2. Download the Linux platform archive using the `wget` command for the required architecture:
```shell
wget https://github.com/DenisKozarezov/interpreter/releases/latest/ipret-linux-arm64.tar.gz
```
3. Unpack the archive and move the downloaded binary file to the current user's executable directory:
```shell
mkdir ipret-temp && tar -xvzf ipret-linux-amd64.tar.gz -C ipret-temp
sudo mv ipret-temp/ipret /usr/local/bin
sudo rm -rf ipret-temp
```
4. Test to ensure the version of `ipret` is the same as downloaded:
```shell
ipret --version
```

### Install and Set Up `ipret` on Windows <img src="https://logo.svgcdn.com/l/microsoft-windows-icon.png" width=25 height=25>

#### Install `ipret` binary via direct download from *GitHub*

1. Open the [Releases](https://github.com/DenisKozarezov/interpreter/releases) section and select the needed version of the `ipret` utility.
2. Download the binary file for the Windows platform for the required architecture:
- `ipret-windows-amd64.zip`
- `ipret-windows-386.zip`
3. Unzip the zip archive using any available archiver.
4. Test to ensure the version of `ipret` is the same as downloaded:
```shell
ipret.exe --version
```

## Usage

The `--help` (`-h`) command is used to display a prompt:
```shell
# Using --help in the root command
ipret --help

# Using --help in the interpreter run command
ipret run --help

# General form of the --help command
ipret <comand> <subcommand> --help
```

```shell
# Evaluate a text file with short flag
ipret run -f ./someFile.txt

# Evaluate a text file with long flag and benchmark
ipret run --filename=someFile.irt --bench
```