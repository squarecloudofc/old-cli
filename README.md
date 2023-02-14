## Square Cloud CLI
Square Cloud CLI is a command line application that provides a variety of commands to interact with the Square Cloud API

## Table of Contents
- [Installation](#installation)
  - [Windows (PowerShell)](#windows-powershell)
  - [Ubuntu/Debian](#ubuntudebian)
  - [MacOS (Homebrew)](#macos-homebrew)
  - [NodeJS](#nodejs)

## Installation
Below are listed some installation methods for Windows, Linux and MacOS. 
If you don't want to use any of the methods below, download a release and add it to `PATH` manually.

### Windows (PowerShell)
`Square Cloud CLI` can be installed using our [Powershell Script](https://github.com/richaardev/squarecloud-cli/master/windows_install.ps1)
```shell
iwr -useb https://raw.githubusercontent.com/richaardev/squarecloud-cli/master/windows_install.ps1 | iex
```
### Ubuntu/Debian
```shell
VERSION=$(wget -q "https://api.github.com/repos/richaardev/squarecloud-cli/releases/latest" -O - | grep -Po '"tag_name": "v\K[^"]*')
wget https://github.com/richaardev/squarecloud-cli/releases/download/v${VERSION}/squarecloud_${VERSION}_linux_amd64.deb
sudo apt install ./squarecloud_${VERSION}_linux_amd64.deb
```

### MacOS (Homebrew)
```shell
brew install richaardev/squarecloud/squarecloud
```

### NodeJS
`Square Cloud CLI` can be installed via NPM
```shell
npm i -g squarecloud
```

## Contributing
We are currently accepting all kinds of suggestions and contributions. All you have to do is open an Issue or a Pull Request!
