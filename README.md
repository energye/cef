<p align="center">
   <span style="color: #2ba9f1;font-size: 24px;font-weight: bold;">Go CEF</span>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        是Go基于 LCL & CEF 构建桌面应用的框架
    </strong>
</p>

---
![go-version](https://img.shields.io/github/go-mod/go-version/energye/cef?logo=git&logoColor=green)
[![github](https://img.shields.io/github/last-commit/energye/cef/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/cef)
[![release](https://img.shields.io/github/v/release/energye/cef?logo=git&logoColor=green)](https://github.com/energye/cef/releases)
![repo](https://img.shields.io/github/repo-size/energye/cef.svg?logo=github&logoColor=green&label=repo-size)
[![Go Report](https://goreportcard.com/badge/github.com/energye/cef)](https://goreportcard.com/report/github.com/energye/cef)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/cef)](https://pkg.go.dev/github.com/energye/cef)
[![license](https://img.shields.io/github/license/energye/cef.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### 项目简介

> [Go CEF](https://github.com/energye/cef) 
> 是 Go 基于
> [LCL](https://www.lazarus-ide.org/)、
> [CEF](https://bitbucket.org/chromiumembedded/cef)
> 开发的框架
>
>> LCL - 基础库, 图形用户界面(GUI)组件库, 提供了非常丰富的系统原生控件
>>
>> CEF - 浏览器组件库 [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi), 在LCL基础上封装的CEF3库
>> 
>> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows平台桌面应用。
>>
>> 将web内容无缝集成到应用程序中，并自定义内容交互以满足应用程序的需求。
> 
> 构建&使用
> 
>> LCL 单独使用, 开发原生图形用户界面(GUI)应用. 轻量级, 丰富的系统原生控件
>
>> LCL + CEF 混合使用, 开发原生图形用户界面(GUI)和浏览器应用. 重量级, 全量chromium API



### 特点

> - 依赖 `CEF二进制框架` 环境
> - 具有完整的 CEF API 和 LCL 系统原生控件
> - 开发环境简单, 编译速度快
> - 前端技术: 支持主流前端框架。例如：Vue、React、Angular 和 原生HTML+CSS+JS等
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端快速调用及数据交互
> - 资源加载: 可无需http服务支撑，直接读取本地资源或内置到执行文件的资源, 也支持http服务加载资源

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/lcl)
- [![WebView4Delphi](https://img.shields.io/badge/Webview2%20-green)](https://github.com/salvadordf/WebView4Delphi)

#### 基本需求

> - Golang >= 1.20
> - 动态链接库 `CEF二进制框架` `liblcl.dll`

#### [示例](https://github.com/energye/examples/tree/main/cef)

#### 开发环境

1. 安装 [Golang](https://golang.google.cn/dl/), Windows版本, 仅支持intel架构 [https://golang.google.cn/dl](https://golang.google.cn/dl)
2. 下载 `CEF` 和 `LCL` 控件库动态链接库
- LCL+CEF:

  - [Windows32](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_windows32.zip/download)
  - [Windows64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_windows64.zip/download)
  - [MacOSx64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_macosx64.zip/download)
  - [MacOSARM64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_macosarm64.zip/download)
  - [Linux64 GTK3](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_linux64.zip/download)
  - [LinuxARM64 GTK3](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_cef_binary_linuxarm64.zip/download)

- CEF 118.7.1:

  - [Windows32](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_windows32_client.tar.bz2)
  - [Windows64](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_windows64_client.tar.bz2)
  - [MacOSx64](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_macosx64_minimal.tar.bz2)
  - [MacOSARM64](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_macosarm64_minimal.tar.bz2)
  - [Linux64](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_linux64_minimal.tar.bz2)
  - [LinuxARM64](https://cef-builds.spotifycdn.com/cef_binary_118.7.1%2Bg99817d2%2Bchromium-118.0.5993.119_linuxarm64_minimal.tar.bz2)
  
  `GTK >= 3.24.24 and Glib2.0 >= 2.66`

- Linux GTK2 LCL + CEF(106)
`To be added`

3. 将CEF解压到指定目录，且将liblcl放置到该目录，并配置环境变量 `ENERGY_HOME` 指向CEF目录
4. 创建Go项目开始使用` CEF` 和 `LCL` 构建桌面应用, 参考 `Go CEF` 示例 [CEF examples](https://github.com/energye/examples/tree/main/cef)

### 相关项目
* [Go LCL](https://github.com/energye/lcl)
* [Go Webview2](https://github.com/energye/wv)
* [Go CEF](https://github.com/energye/cef)
* [Go Webkit](https://github.com/energye/wk)
* [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi)
* [CEF](https://github.com/chromiumembedded/cef)
* [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
* [CefSharp](https://github.com/cefsharp/CefSharp)
* [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
* [cefpython](https://github.com/cztomczak/cefpython)
* [Chromium](https://chromium.googlesource.com/chromium/src/)

---

### 欢迎加入
energy底层由多个项目模块组成, 因过于复杂扔处于建设的过程中，有很多的事情无法独自完成，如果有感兴趣的同学想参与energy的实现或学习，可通过微信或QQ联系我。

如果你觉得此项目对你有帮助，请点亮 Star

---

### ENERGY QQ交流群 & 微信

<p align="center">
    <img src="https://assets.yanghy.cn/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://assets.yanghy.cn/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
</p>

---

### 鸣谢 Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

---

### 项目截图
##### Windows-10
<img src="https://assets.yanghy.cn/CEF-simple.png">

----

### 开源协议

[![license](https://img.shields.io/github/license/energye/cef.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)
