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