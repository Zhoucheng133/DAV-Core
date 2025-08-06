# DAV Core

## 简介

![License](https://img.shields.io/badge/License-MIT-dark_green)

[DAV Server](https://github.com/Zhoucheng133/DAV-Server)的核心组件，由go语言编写

> [!WARNING]
> 如果你要在macOS上使用，在运行前先执行`chmod +x <可执行文件位置>`

## 属性

`-port` [必须]，服务端口  
`-path` [必须]，分享路径  
`-u` [如果没有-p属性可以省略]，需要登录的用户名  
`-p` [如果没有-u属性可以省略]，需要登录的密码


## 构建

### 准备

你需要在你的设备上安装配置好这些东西：
- go

### 生成二进制文件

1. 你需要先克隆或者下载本仓库
2. 生成二进制文件
   ```bash
   go run . #运行程序
   go build #打包
   ```

### 生成动态库供[DAV Server](https://github.com/Zhoucheng133/DAV-Server)使用或二次开发

1. 你需要先克隆或者下载本仓库
2. 生成动态库
   ```bash
   go build -buildmode=c-shared -o server.dll .    # 生成Windows动态库
   go build -buildmode=c-shared -o server.dylib .  # 生成macOS动态库
   ```
