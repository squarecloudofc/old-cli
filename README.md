## Square Cloud CLI

Square Cloud CLI is a command line application that provides a variety of commands to interact with the Square Cloud API

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
  - [Windows - Winget](#using-winget)
  - [Windows - PowerShell](#using-powershell)
  - [Linux - Snap](#using-snap)
  - [Linux - Bash](#using-bash)
  - [MacOS - Homebrew](#homebrew-linuxmacos)
  - [NodeJS](#nodejs-windowslinuxmacos)
- [Contributing](#contributing)

## Usage
```shell
squarego is Square Cloud on the command line. Which allows you to manage all your applications from the command line!

Usage:
  squarecloud [flags]
  squarecloud [command]

Available Commands:
  aboutme     View your Square Cloud user information
  apps        See all your active apps on Square Cloud
  backup      Make a backup of your application
  check       Check your configuration file
  commit      Update your application
  delete      delete your application
  help        Help about any command
  init        initialize squarecloud config
  login       Log in for full SquareCloud CLI access
  logs        See the most recent logs of your application
  restart     Restart your application
  start       Start running your application
  status      Get your application status
  stop        Stop running your application
  upload      Upload a new app to SquareCloud
  version     print cli version

Flags:
  -h, --help      help for squarecloud
  -v, --version   version for squarecloud

Use "squarecloud [command] --help" for more information about a command.
```

## Installation

Below are listed some installation methods for Windows, Linux and MacOS.
If you don't want to use any of the methods below, download a release and add it to `PATH` manually.

## Installation in Windows

### Using Winget

I'ts recommended you install `Square Cloud CLI` using `winget`

```shell
winget install SquareCloud.cli
```

### Using PowerShell

`Square Cloud CLI` can be installed using our [Powershell Script](https://github.com/richaardev/squarecloud-cli/master/windows_install.ps1)

```shell
iwr -useb https://raw.githubusercontent.com/richaardev/squarecloud-cli/master/windows_install.ps1 | iex
```

## Installation in Linux

### Using Snap

```shell
sudo snap install squarecloud
```

### Using Bash

```shell
VERSION=$(wget -q "https://api.github.com/repos/richaardev/squarecloud-cli/releases/latest" -O - | grep -Po '"tag_name": "v\K[^"]*')
wget https://github.com/richaardev/squarecloud-cli/releases/download/v${VERSION}/squarecloud_${VERSION}_linux_amd64.deb
sudo apt install ./squarecloud_${VERSION}_linux_amd64.deb
```

## Other Installations

### Homebrew (Linux/MacOS)

```shell
brew tap richaardev/squarecloud
brew install squarecloud
```

### NodeJS (Windows/Linux/MacOS)

`Square Cloud CLI` can be installed via NPM

```shell
npm i -g squarecloud
```

## Contributing

We are currently accepting all kinds of suggestions and contributions. All you have to do is open an Issue or a Pull Request!
