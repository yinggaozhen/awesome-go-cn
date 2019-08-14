# Awesome Go

[Gold]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Gold.svg "star > 5000"
[Silver]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Silver.svg "star > 1000"
[Bronze]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Bronze.svg "star > 100"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "最近一周有更新"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "最近一年没有更新"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "包含中文文档"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "项目已归档"

**此项目是 [awesome-go](https://awesome-go.com/) 中文版，最后一次同步时间 : 2019-08-14 09:54:26(每隔1天同步一次)**

[![chinese](https://yinggaozhen.github.io/docs/english.svg)](README_EN.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) 为Awesome Go打赏~

精选了一系列很棒的Go框架、库和软件。灵感来自于[awesome-python](https://github.com/vinta/awesome-python)。

**小图标说明** :


小图标 | 说明  
:-:|-
![star > 5000][Gold] | star数量 > 5000 的项目
![star > 1000][Silver] | star数量 > 1000 的项目
![star > 100][Bronze] | star数量 > 100 的项目
![最近一个周有更新][Green] | 最近一周有更新。可以基本判断当前库处于积极维护状态。
![最近一年未更新][Yellow] | 最近一年没有更新。反应了此库的维护积极性不高，使用时需谨慎。
![归档项目][Archived] | 此项目已归档，不再更新，使用时需谨慎。
![此项目有中文文档][CN] | 此项目有中文文档。

### 说明

[中文](README.md)  | [English](README_EN.md) 
 

### 贡献

你可以快速浏览贡献者名单[contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md). 感谢所有为此项目付出的同学[contributors](https://github.com/avelino/awesome-go/graphs/contributors); 你真棒!

如果您在看到一个包或项目不再维护或不适合，请往awesome-go提交[Pull Request](https://github.com/avelino/awesome-go/pulls)，本项目每隔一天与英文文档同步。感谢!

### 内容

- [Awesome Go](#awesome-go)
    - [音频和音乐](#音频和音乐)
    - [身份验证和OAuth](#身份验证和oauth)
    - [Bot建设](#bot建设)
    - [命令行](#命令行)
    - [配置](#配置)
    - [持续集成](#持续集成)
    - [CSS预处理器](#css预处理器)
    - [数据结构](#数据结构)
    - [数据库](#数据库)
    - [数据库驱动程序](#数据库驱动程序)
    - [日期和时间](#日期和时间)
    - [分布式系统](#分布式系统)
    - [电子邮件](#电子邮件)
    - [可嵌入的脚本语言](#可嵌入的脚本语言)
    - [错误处理](#错误处理)
    - [文件](#文件)
    - [金融](#金融)
    - [表单](#表单)
    - [方法](#方法)
    - [游戏开发](#游戏开发)
    - [代码生成与泛型](#代码生成与泛型)
    - [地理](#地理)
    - [Go 编译器](#go-编译器)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [硬件](#硬件)
    - [图片](#图片)
    - [物联网](#物联网)
    - [作业调度器](#作业调度器)
    - [JSON](#json)
    - [日志记录](#日志记录)
    - [机器学习](#机器学习)
    - [消息](#消息)
    - [微软办公软件](#微软办公软件)
        - [Microsoft Excel](#microsoft-excel)
    - [杂项](#杂项)
        - [依赖注入](#依赖注入)
        - [项目布局](#项目布局)
        - [字符串](#字符串)
    - [自然语言处理](#自然语言处理)
    - [网络](#网络)
        - [HTTP客户端](#http客户端)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [包管理](#包管理)
    - [性能](#性能)
    - [查询语言](#查询语言)
    - [嵌入的资源](#嵌入的资源)
    - [科学与数据分析](#科学与数据分析)
    - [安全](#安全)
    - [序列化](#序列化)
    - [流处理](#流处理)
    - [模板引擎](#模板引擎)
    - [测试](#测试)
    - [文本处理](#文本处理)
    - [Third-party APIs](#third-party-apis)
    - [公用事业公司](#公用事业公司)
    - [UUID](#uuid)
    - [验证](#验证)
    - [版本控制](#版本控制)
    - [视频](#视频)
    - [Web框架](#web框架)
        - [中间件](#中间件)
            - [仿真中间件](#仿真中间件)
            - [用于创建HTTP中间件的库](#用于创建http中间件的库)
        - [路由器](#路由器)
    - [Windows](#windows)
    - [XML](#xml)

- [工具](#工具)
    - [代码分析](#代码分析)
    - [编辑器插件](#编辑器插件)
    - [Go 生成工具](#go-生成工具)
    - [Go 工具](#go-工具)
    - [软件包](#软件包)
        - [DevOps 工具](#devops-工具)
        - [其他软件](#其他软件)

- [服务器应用程序](#服务器应用程序)

- [资源](#资源)
    - [基准](#基准)
    - [会议](#会议)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [聚会](#聚会)
    - [Twitter](#twitter)
    - [网站](#网站)
        - [教程](#教程)

## 音频和音乐

*用于操作音频的库。 (翻译出错了? 试试 [英文版](README_EN.md#audio-and-music) 吧~)*

* [Oto](https://github.com/hajimehoshi/oto) **star:440** 多平台的 low-level 声音播放库。   ![star > 100][Bronze]
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:303** 基于 Go 的PortAudio audio I/O库。   ![star > 100][Bronze]
* [music-theory](https://github.com/go-music-theory/music-theory) **star:257** 基于 Go 的音乐理论模型。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [waveform](https://github.com/mdlayher/waveform) **star:248** 通过音频流生成波形图像的包。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [portmidi](https://github.com/rakyll/portmidi) **star:207** PortMidi的 Go 语言实现接口.   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [id3v2](https://github.com/bogem/id3v2) **star:110** 快速稳定的 ID3 解析及写入Go库。   ![star > 100][Bronze]
* [flac](https://github.com/mewkiz/flac) **star:102** 原生 Go FLAC编码器/解码器，支持FLAC流。   ![star > 100][Bronze]
* [mix](https://github.com/go-mix/mix) **star:98** 基于序列的 Go 原生音乐混音器。   ![最近一年没有更新][Yellow]
* [mp3](https://github.com/tcolgate/mp3) **star:96** 原生 Go MP3解码器。   ![最近一年没有更新][Yellow]
* [go-sox](https://github.com/krig/go-sox) **star:93** libsox 的 Go 语言实现接口。   ![最近一年没有更新][Yellow]
* [flac](https://github.com/eaburns/flac) **star:83** 原生 Go FLAC解码器，将FLAC文件解码为字节片。   ![最近一年没有更新][Yellow]
* [malgo](https://github.com/gen2brain/malgo) **star:71** 迷你音频库。
* [taglib](https://github.com/wtolson/go-taglib) **star:66** taglib 的 Go 语言实现接口.   ![最近一年没有更新][Yellow]
* [gaad](https://github.com/Comcast/gaad) **star:56** 原生 Go AAC位流解析器。   ![最近一年没有更新][Yellow]
* [minimp3](https://github.com/tosone/minimp3) **star:25** 轻量级MP3解码器库。
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:24** libmediainfo 的 Go 语言实现接口。   ![最近一年没有更新][Yellow]
* [vorbis](https://github.com/mccoyst/vorbis) **star:22** “原生” Go Vorbis解码器(使用CGO，但没有依赖关系)。
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:22** EasyMidi是一个简单可靠的库，用于处理标准midi文件(SMF)。   ![最近一年没有更新][Yellow]
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:8** libsamplerate 的 Go 语言实现接口。   ![最近一年没有更新][Yellow]

## 身份验证和OAuth

*用于实现验证方案的库。 (翻译出错了? 试试 [英文版](README_EN.md#authentication-and-oauth) 吧~)*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:6041** JSON Web令牌(JWT)。   ![star > 5000][Gold]
* [casbin](https://github.com/hsluoyz/casbin) **star:4993** 支持ACL、RBAC、ABAC等访问控制模型的授权库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [oauth2](https://github.com/golang/oauth2) **star:2387** goauth2的继任者。通用OAuth 2.0包，附带JWT、谷歌api、计算引擎和应用程序引擎支持。   ![star > 1000][Silver]
* [goth](https://github.com/markbates/goth) **star:2283** 提供了 OAuth 和 OAuth2 的简单清晰易用的方法。可开箱即用处理多个提供程序。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [authboss](https://github.com/volatiletech/authboss) **star:1938** web模块化认证系统。它试图删除尽可能多的模板文件和硬编码，以便每次新建一个新的web项目时，您都可以插入、配置并开始构建您的应用程序，而不必每次都构建一个身份验证系统。   ![star > 1000][Silver]
* [osin](https://github.com/openshift/osin) **star:1547** OAuth2服务器库。   ![star > 1000][Silver]
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1288** 用 Golang 编写的独立且符合规范的OAuth2服务器。   ![star > 1000][Silver]
* [go-jose](https://github.com/square/go-jose) **star:1135** 相当完整地实现了JOSE工作组的JSON Web令牌、JSON Web签名和JSON Web加密规范。   ![star > 1000][Silver]
* [gologin](https://github.com/dghubble/gologin) **star:1044** 用于使用OAuth1和OAuth2身份验证提供者登录的可链处理程序。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gorbac](https://github.com/mikespook/gorbac) **star:915** 轻量级的基于角色的访问控制(RBAC)实现。   ![star > 100][Bronze]
* [loginsrv](https://github.com/tarent/loginsrv) **star:822** JWT登录微服务带有可插拔的后端服务，如OAuth2 (Github)、htpasswd、osiam。   ![star > 100][Bronze]
* [scs](https://github.com/alexedwards/scs) **star:530** HTTP服务器的会话管理器。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [permissions2](https://github.com/xyproto/permissions2) **star:352** 用于跟踪用户、登录状态和权限的库。依赖于cookie安全和bcrypt。   ![star > 100][Bronze]
* [paseto](https://github.com/o1egl/paseto) **star:242** 平台无关的安全令牌(PASETO)。   ![star > 100][Bronze]
* [httpauth](https://github.com/goji/httpauth) **star:181** HTTP身份验证中间件。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:153** JWT中间件，可用于Golang http服务器，提供了许多配置选项。   ![star > 100][Bronze]
* [jwt](https://github.com/pascaldekloe/jwt) **star:94** 轻量级JSON Web令牌库。
* [session](https://github.com/icza/session) **star:91** web服务器会话管理(包括支持谷歌应用程序引擎 - GAE)。
* [branca](https://github.com/hako/branca) **star:77** 基于 Go 实现Branca令牌。   ![最近一年没有更新][Yellow]
* [jwt](https://github.com/robbert229/jwt) **star:71** 简单易用的JSON Web令牌实现(JWT)。
* [sessions](https://github.com/adam-hanna/sessions) **star:47** 非常简单，高性能，可深度定制的会话服务，主要用于的 go http 服务器。
* [sjwt](https://github.com/brianvoe/sjwt) **star:33** 简单的jwt生成器和解析器。
* [securecookie](https://github.com/chmike/securecookie) **star:32** 高效安全的cookie编码/解码。
* [rbac](https://github.com/zpatrick/rbac) **star:27** 最小的RBAC包。
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** 使用SessionGate Redis模块进行会话管理。
* [signedvalue](https://github.com/sashka/signedvalue) **star:7** 与[Tornado's](https://github.com/tornadooweb/tornado) 完全兼容的签名和时间戳字符串实现。
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** 提供cookie .txt文件格式的解析器。   ![最近一年没有更新][Yellow]

## Bot建设

*用于构建和使用机器人的库。 (翻译出错了? 试试 [英文版](README_EN.md#bot-building) 吧~)*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1634** 简单轻量级的Telegram bot客户端。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [telebot](https://github.com/tucnak/telebot) **star:956** 用Go编写的Telegram bot框架。   ![star > 100][Bronze]
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:470** 用 Go 编写的IRC, Slack和电报机器人。   ![star > 100][Bronze]
* [slacker](https://github.com/shomali11/slacker) **star:317** 可简单创建Slack机器人的框架。   ![star > 100][Bronze]
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:223** 基于控制台的，用于加密货币交易所的的交易机器人。   ![star > 100][Bronze]
* [tbot](https://github.com/yanzay/tbot) **star:220** 带有类似于net/http API的Telegram bot服务器。   ![star > 100][Bronze]
* [Tenyks](https://github.com/kyleterry/tenyks) **star:168** 面向服务的IRC bot，使用Redis和JSON进行消息传递。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Kelp](https://github.com/stellar/kelp) **star:162** 官方交易和做市机器人为[Stellar](https://www.stellar.org/) DEX。开箱即用的作品，用 Golang 编写，兼容集中交易和定制交易策略。   ![star > 100][Bronze]
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:137** 此框架提供了聊天机器人相关的服务，包括LINE、Slack、Gitter等。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [hanu](https://github.com/sbstjn/hanu) **star:109** 用于编写Slack机器人的框架。   ![star > 100][Bronze]
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:86** 由swagger文件、基于会话的路由器和中间件生成的纯Golang Telegram Bot API包装器。   ![最近一年没有更新][Yellow]
* [margelet](https://github.com/zhulik/margelet) **star:56** 构建电报机器人的框架。   ![最近一年没有更新][Yellow]
* [govkbot](https://github.com/nikepan/govkbot) **star:25** 简单的Go [VK](https://vk.com) bot库。
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:12** 另一个构建Slack机器人的框架。
* [micha](https://github.com/onrik/micha) **star:10** 基于 GO 实现的Telegram 机器人API库。   ![最近一年没有更新][Yellow]

## 命令行

### 标准CLI

*用于构建标准或基本命令行应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#standard-cli) 吧~)*

* [cobra](https://github.com/spf13/cobra) **star:13425** 现代Go CLI命令行交互工具。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [urfave/cli](https://github.com/urfave/cli) **star:11499** 可让你简单、快速和愉快的构建命令行应用(之前是codegangsta/cli)。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [kingpin](https://github.com/alecthomas/kingpin) **star:2555** 支持子命令的命令行和标志解析器。   ![star > 1000][Silver]
* [go-flags](https://github.com/jessevdk/go-flags) **star:1512**  Go 命令行选项解析器。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [readline](https://github.com/chzyer/readline) **star:1373** 纯golang实现，在MIT许可下提供了GNU-Readline的大部分特性。   ![star > 1000][Silver]
* [docopt.go](https://github.com/docopt/docopt.go) **star:1177** 会让你满意的命令行参数解析器。   ![star > 1000][Silver]
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1001** 用于实现命令行接口的Go库。   ![star > 1000][Silver]
* [cli-init](https://github.com/tcnksm/gcli) **star:870** 一个简单就可开启构建Golang命令行的应用程序。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [pflag](https://github.com/spf13/pflag) **star:754** 基于POSIX/GNU-style --flags实现的包，主要用于替换Go的falg包。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-arg](https://github.com/alexflint/go-arg) **star:740** 基于结构的参数解析。   ![star > 100][Bronze]
* [mow.cli](https://github.com/jawher/mow.cli) **star:624** 用于构建具有复杂标志和参数解析和验证的CLI应用程序。   ![star > 100][Bronze]
* [complete](https://github.com/posener/complete) **star:624** 使用 Go 语言编写的 bash 命令补全工具以及 Go 命令补全工具.   ![star > 100][Bronze]
* [liner](https://github.com/peterh/liner) **star:591** 类似readline-like的命令行接口库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [cli](https://github.com/mkideal/cli) **star:480** 基于golang结构标签，功能丰富易于使用的命令行包。   ![star > 100][Bronze]
* [flaggy](https://github.com/integrii/flaggy) **star:454** 一个健壮的、易用的标志包，具有出色的子命令支持。   ![star > 100][Bronze]
* [ops](https://github.com/nanovms/ops) **star:267** Unikernel 构建器/协调器。   ![star > 100][Bronze]
* [argparse](https://github.com/akamensky/argparse) **star:109** 命令行参数分析器，灵感来自Python的argparse模块。   ![star > 100][Bronze]
* [flag](https://github.com/cosiner/flag) **star:101** 简单但功能强大的命令行选项解析库，用于支持Go子命令。   ![star > 100][Bronze]
* [ukautz/clif](https://github.com/ukautz/clif) **star:97** 简小的命令行接口框架。
* [commandeer](https://github.com/jaffee/commandeer) **star:93** 开发友好的CLI应用程序。
* [sflags](https://github.com/octago/sflags) **star:91** 基于结构的flag生成器，用于flag、urfave/cli、pflag、cobra、kingpin和其他库。
* [wmenu](https://github.com/dixonwille/wmenu) **star:86** 为cli程序提供了简单上手的菜单，可提示用户作出选择。
* [cli](https://github.com/teris-io/cli) **star:57** 为 Go 构建命令接口提供简单而完整的API。
* [job](https://github.com/liujianping/job) **star:51** 工作，把你的短期指令当作长期任务。
* [env](https://github.com/codingconcepts/env) **star:42** 基于标记的结构化的环境配置。
* [wlog](https://github.com/dixonwille/wlog) **star:35** 支持跨平台和并发的简单日志记录接口。   ![最近一年没有更新][Yellow]
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  具有自动配置和依赖注入的cli应用程序框架。
* [gocmd](https://github.com/devfacet/gocmd) **star:33** 用于构建命令行应用程序。
* [flagvar](https://github.com/sgreben/flagvar) **star:31** 符合 Go 标准的“flag”标志参数类型包。
* [strumt](https://github.com/antham/strumt) **star:27** 用于创建提示链。
* [argv](https://github.com/cosiner/argv) **star:17** 基于Base 语法，用于分隔命令行字符串并将其作为参数的 Go 语言库，
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** 用于简化CLI工作流的 Go 库。
* [cmdr](https://github.com/hedzr/cmdr) **star:15** 一个POSIX/GNU风格的、类似getopt的命令行UI Go库。   ![最近一周有更新][Green]
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:9**  Go 选择解析器，借鉴于Perl灵活性的GetOpt::Long。   ![最近一周有更新][Green]
* [sand](https://github.com/Zaba505/sand) **star:7** 用于创建解释器等的简单API。
* [ts](https://github.com/liujianping/ts) **star:6** 时间戳转换和比较工具。

### 高级控制台用户界面

*用于构建控制台应用程序和控制台用户界面的库。 (翻译出错了? 试试 [英文版](README_EN.md#advanced-console-uis) 吧~)*

* [termui](https://github.com/gizak/termui) **star:8955** 此库是基于**termbox-go**实现的，借鉴于[blessed-contrib](https://github.com/yaronn/blessed-contrib)。   ![star > 5000][Gold]
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  更换终端文本样式。
* [gocui](https://github.com/jroimartin/gocui) **star:5408** 旨在创建控制台用户界面的极简Go库。   ![star > 5000][Gold]
* [termbox-go](https://github.com/nsf/termbox-go) **star:3495** 基于文本的跨平台接口库。   ![star > 1000][Silver]
* [color](https://github.com/fatih/color) **star:3028** 多功能包装，彩色终端输出。   ![star > 1000][Silver]   ![归档项目][Archived]
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2346** 构建一个强大的交互式提示，借鉴于[python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit)   ![star > 1000][Silver]
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1544** 在终端呈现进度条，可灵活配置的。   ![star > 1000][Silver]
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1148** 在命令行中构建轻量级ASCII线图╭┈╯，应用程序中没有其他依赖项。   ![star > 1000][Silver]
* [uilive](https://github.com/gosuri/uilive) **star:834** 用于实时更新终端输出的库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [termtables](https://github.com/apcera/termtables)  使用Ruby库[terminal-tables](https://github.com/tj/terminal-table)的端口生成简单的ASCII表，并提供标记和HTML输出。
* [termdash](https://github.com/mum4k/termdash) **star:812** 此库是基于**termbox-go**实现的，借鉴于[termui](https://github.com/gizak/termui)。   ![star > 100][Bronze]
* [mpb](https://github.com/vbauerster/mpb) **star:708** 可在终端显示多进度条。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [aurora](https://github.com/logrusorgru/aurora) **star:638** 支持fmt.Printf/Sprintf的ANSI终端颜色。   ![star > 100][Bronze]
* [progressbar](https://github.com/schollz/progressbar) **star:578** 基本线程安全的进度条，在每个操作系统工作。   ![star > 100][Bronze]
* [uitable](https://github.com/gosuri/uitable) **star:505** 改善终端应用程序中表格数据的可读性。   ![star > 100][Bronze]
* [go-colorable](https://github.com/mattn/go-colorable) **star:378** 适用于windows的颜色编写器。   ![star > 100][Bronze]
* [go-isatty](https://github.com/mattn/go-isatty) **star:346** Go 实现的 isatty。   ![star > 100][Bronze]
* [chalk](https://github.com/ttacon/chalk) **star:305** 美化终端/控制台输出。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [tabby](https://github.com/cheynewallace/tabby) **star:248** 一个可在终端生成一个极简Golang表格轻量级库   ![star > 100][Bronze]
* [gookit/color](https://github.com/gookit/color) **star:207** 终端显色工具库，支持16种颜色，256种颜色，RGB显色输出，兼容Windows。   ![star > 100][Bronze]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:196** 在终端中使用彩色文字。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [simpletable](https://github.com/alexeyco/simpletable) **star:169** 可在终端显示简易表格。   ![star > 100][Bronze]
* [cfmt](https://github.com/mingrammer/cfmt) **star:67** 提供上下文的fmt，灵感来自于bootstrap color classes。
* [tabular](https://github.com/InVisionApp/tabular) **star:29** 不需要向API传递大量参数就可从命令行实用程序中打印ASCII表。   ![最近一年没有更新][Yellow]
* [colourize](https://github.com/TreyBastian/colourize) **star:16** 在终端提供ANSI彩色文本。   ![最近一年没有更新][Yellow]
* [ctc](https://github.com/wzshiming/ctc) **star:10** 不需要Print方法的非侵入性跨平台终端颜色库。   ![包含中文文档][CN]
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** 在终端提供ANSI彩色文本模板。   ![最近一年没有更新][Yellow]

## 配置

*配置解析的库。 (翻译出错了? 试试 [英文版](README_EN.md#configuration) 吧~)*

* [viper](https://github.com/spf13/viper) **star:9402** 配置管理。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2418** 管理来自环境变量的配置数据。   ![star > 1000][Silver]
* [godotenv](https://github.com/joho/godotenv) **star:2139** Ruby 的 dotenv 库的 Go移植版(从.env文件加载环境变量)。   ![star > 1000][Silver]
* [ini](https://github.com/go-ini/ini) **star:1603**  读和写INI文件。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [env](https://github.com/caarlos0/env) **star:876** 解析环境变量并赋值到struct中(默认值)。   ![star > 100][Bronze]
* [konfig](https://github.com/lalamove/konfig) **star:514** 可组合、可观察和高性能的分布式配置管理。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [store](https://github.com/tucnak/store) **star:240** 轻量级配置管理器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [confita](https://github.com/heetch/confita) **star:238** 从多个后端级联加载配置到struct中。   ![star > 100][Bronze]
* [config](https://github.com/olebedev/config) **star:210** 带有环境变量和标记解析的JSON或YAML配置包装器。   ![star > 100][Bronze]
* [joshbetz/config](https://github.com/joshbetz/config) **star:192** 一个可解析环境变量、JSON文件小巧的配置库，在SIGHUP时会自动重新加载。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [config](https://github.com/JeremyLoy/config) **star:190** 云本地应用程序配置。将ENV绑定到结构体仅需两行代码。   ![star > 100][Bronze]
* [hjson](https://github.com/hjson/hjson-go) **star:172** 更加人性化的JSON配置。轻松的语法，更少的错误，更多的注释。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [envconfig](https://github.com/vrischmann/envconfig) **star:144** 从环境变量中读取配置。   ![star > 100][Bronze]
* [gcfg](https://github.com/go-gcfg/gcfg) **star:117** 将ini的配置文件读入 Go structs中;支持用户定义的类型和子选项。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [goConfig](https://github.com/crgimenes/goConfig) **star:110** 将结构体解析为输入，并用来自命令行、环境变量和配置文件的参数填充该结构体的字段。   ![star > 100][Bronze]
* [envh](https://github.com/antham/envh) **star:92** 协助管理环境变量的Helpers。
* [envcfg](https://github.com/tomazk/envcfg) **star:89** 对环境变量进行解析，并赋值到struct。   ![最近一年没有更新][Yellow]
* [koanf](https://github.com/knadh/koanf) **star:87** 轻量级可扩展库，用于读取Go应用程序中的配置。内置支持JSON, TOML, YAML, env，命令行。
* [gookit/config](https://github.com/gookit/config) **star:86** 程序配置管理(load,get,set)。支持JSON, YAML, TOML, INI, HCL。支持多文件加载，数据覆盖合并。   ![包含中文文档][CN]
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  模块化的JSON配置。保持配置结构及其配置的代码，并将解析委托给子模块，而不牺牲配置的完整序列化。
* [gofigure](https://github.com/ian-kent/gofigure) **star:57** 让程序配置变得简单。   ![最近一年没有更新][Yellow]
* [configure](https://github.com/paked/configure) **star:49** 通过多个源提供配置，包括JSON、flags和环境变量。
* [harvester](https://github.com/beatlabs/harvester) **star:40** 一个易于使用的静态和动态配置包
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:37** 遵循[XDG标准](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)的跨平台包。
* [ingo](https://github.com/schachmat/ingo) **star:24** flag保存在类ini的配置文件中。   ![最近一年没有更新][Yellow]
* [go-up](https://github.com/ufoscout/go-up) **star:24** 一个简单的配置库，具有递归占位符解析功能。
* [mini](https://github.com/sasbury/mini) **star:19** 用于解析ini类型的配置文件。
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:8** 合并来自任意url的多个JSON/YAML/TOML文件、针对JSON模式的验证以及模式中定义的默认值的应用程序。
* [envconf](https://github.com/ian-kent/envconf) **star:7** 从环境配置中读取配置。   ![最近一年没有更新][Yellow]
* [sprbox](https://github.com/oblq/sprbox) **star:3** 支持构建环境的工具箱工厂和其他不确定的配置解析器(如YAML、TOML、JSON和环境vars)。

## 持续集成

*用于帮助进行持续集成的工具。 (翻译出错了? 试试 [英文版](README_EN.md#continuous-integration) 吧~)*

* [drone](https://github.com/drone/drone) **star:19093** Drone 是一个基于 Docker 的持续集成平台，用 Go 编写。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [goveralls](https://github.com/mattn/goveralls) **star:579** Coveralls.io 是一个用 Go 编写，可持续对代码覆盖率进行检测的系统。   ![star > 100][Bronze]
* [overalls](https://github.com/go-playground/overalls) **star:98** 针对多package 的 Go 语言项目，可为类似 goveralls 这样的工具生成覆盖率报告。
* [duci](https://github.com/duck8823/duci) **star:44** 一个简单的 ci 服务。   ![最近一周有更新][Green]
* [gomason](https://github.com/nikogura/gomason) **star:33** 在一个干净的工作区中对你的 Go 二进制文件进行测试、构建、签名和发布。
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** 递归覆盖测试工具。   ![最近一年没有更新][Yellow]

## CSS预处理器

*用于预处理CSS文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#css-preprocessors) 吧~)*

* [gcss](https://github.com/yosssi/gcss) **star:423** 纯Go编写的 CSS 预处理器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-libsass](https://github.com/wellington/go-libsass) **star:131**  采用 Go封装，100% 与 Sass 兼容的 libsass 项目。   ![star > 100][Bronze]

## 数据结构

*用 Go 实现的通用的数据结构和算法。 (翻译出错了? 试试 [英文版](README_EN.md#data-structures) 吧~)*

* [gods](https://github.com/emirpasic/gods) **star:6427** 数据结构。容器、集合、列表、堆栈、地图、BidiMaps、树、HashSet等。   ![star > 5000][Gold]
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5141** 可靠的、高性能的和线程安全的数据结构的集合。   ![star > 5000][Gold]
* [golang-set](https://github.com/deckarep/golang-set) **star:1177** 线程安全和非线程安全的高性能集。   ![star > 1000][Silver]
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1163** 用于处理连续的概率数据结构。   ![star > 1000][Silver]
* [gota](https://github.com/kniren/gota) **star:885** 实现了数据帧，序列以及数据噪音。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [willf/bloom](https://github.com/willf/bloom) **star:669** 实现Bloom过滤器。   ![star > 100][Bronze]
* [roaring](https://github.com/RoaringBitmap/roaring) **star:667** 实现了压缩 bitsets 的Go包。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:662** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   ![star > 100][Bronze]
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:515** 布谷鸟过滤器:一个用Go实现，可替代计数 bloom 过滤器。   ![star > 100][Bronze]
* [bitset](https://github.com/willf/bitset) **star:481** 实现了 bitsets 的 Go 包。   ![star > 100][Bronze]
* [trie](https://github.com/derekparker/trie) **star:424** 在Go中实现Trie。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:313** 基于内存的地理索引。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [mafsa](https://github.com/smartystreets/mafsa) **star:273** 实现了 MA-FSA ，包含最小完美哈希。   ![star > 100][Bronze]   ![归档项目][Archived]
* [algorithms](https://github.com/shady831213/algorithms) **star:242** 算法和数据结构。来源于CLRS。   ![star > 100][Bronze]
* [goskiplist](https://github.com/ryszard/goskiplist) **star:194** 基于 Go 的跳表实现。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [hilbert](https://github.com/google/hilbert) **star:182** 用于映射空间填充曲线（例如 Hilbert 曲线和 Peano 曲线）和数值。   ![star > 100][Bronze]
* [merkletree](https://github.com/cbergoon/merkletree) **star:146** 实现了merkle树，提供了对数据结构内容的有效和安全的验证。   ![star > 100][Bronze]
* [bloom](https://github.com/zhenjl/bloom) **star:128** 在Go中实现了Bloom过滤器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:126** 帮助用户构建自定义二进制流的二进制封装器和解包器   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [skiplist](https://github.com/MauriceGit/skiplist) **star:101** 高性能的 Go 跳表实现。   ![star > 100][Bronze]
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:100** 区域四叉树具有高效的点定位和邻域查找功能。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:99** 基于内存的LRU算法实现。
* [encoding](https://github.com/zhenjl/encoding) **star:94** 整形压缩库。   ![最近一年没有更新][Yellow]
* [ring](https://github.com/TheTannerRyan/ring) **star:90** 高性能、线程安全的bloom过滤器。
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:89** 自适应基数树。
* [conjungo](https://github.com/InVisionApp/conjungo) **star:80** 一个小型、强大和灵活的合并库。
* [deque](https://github.com/gammazero/deque) **star:67** 快速环缓冲区deque(双端队列)。
* [skiplist](https://github.com/gansidui/skiplist) **star:64** 在Go中实现了跳表。   ![最近一年没有更新][Yellow]
* [bit](https://github.com/yourbasic/bit) **star:55** Go 语言集合数据结构。提供了额外的位操作功能。   ![最近一年没有更新][Yellow]
* [levenshtein](https://github.com/agnivade/levenshtein) **star:55** 实现在Go中计算levenshtein距离。
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:43** Go实现Count-Min-log sketch的功能 : 使用近似计数器进行近似计数(类似Count-Min sketch，但使用更少内存)。   ![最近一年没有更新][Yellow]
* [bloom](https://github.com/yourbasic/bloom) **star:39** Golang Bloom过滤器的实现。   ![最近一年没有更新][Yellow]
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:35** 基于内存的实现了高性能的key:value存储库。指针缓存。
* [levenshtein](https://github.com/agext/levenshtein) **star:33** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:29** 并行FIFO队列。
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:24** 具备高并发能力，可替换 bufio.Writer。   ![最近一年没有更新][Yellow]
* [crunch](https://github.com/superwhiskers/crunch) **star:19** 打包实现缓冲区，以便轻松处理各种数据类型。   ![最近一周有更新][Green]
* [goset](https://github.com/zoumo/goset) **star:16** 一个有用的Go集合实现。   ![最近一周有更新][Green]
* [pipeline](https://github.com/hyfather/pipeline) **star:15** 实现了 fan-in 和 fan-out 的管道功能。
* [go-ef](https://github.com/amallia/go-ef) **star:11** 实现了 Elias-Fano 编码。   ![最近一年没有更新][Yellow]
* [typ](https://github.com/gurukami/typ) **star:9** 从复杂结构中获取值，支持空类型、安全的类型转换。
* [dict](https://github.com/srfrog/dict) **star:9** 实现类似 python dict的功能(dict)。
* [mspm](https://github.com/BlackRabbitt/mspm) **star:7** 用于信息检索的多字符串模式匹配算法。   ![最近一年没有更新][Yellow]
* [hide](https://github.com/emvi/hide) **star:7** ID type with marshalling to/from hash to prevent sending IDs to clients.
* [treap](https://github.com/perdata/treap) **star:7** 使用树堆进行持久、快速有序的映射。
* [deque](https://github.com/edwingeng/deque) **star:7** 高度优化的双端队列。
* [set](https://github.com/StudioSol/set) **star:6** 使用LinkedHashMap在Go中实现简单的set数据结构。
* [null](https://github.com/emvi/null) **star:5** 对空的 Go 数据类型也可以进行JSON进行解析/反解析。
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** 用于解析类似json的日志的工具，用于收集惟一的字段和事件。
* [ptrie](https://github.com/viant/ptrie) **star:1** 前缀树的一种实现。
* [timedmap](https://github.com/zekroTJA/timedmap) **star:1** 实现了有生命周期的键值对Map。

## 数据库

*数据库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [prometheus](https://github.com/prometheus/prometheus) **star:25692** 用于监控系统和时序的数据库。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [tidb](https://github.com/pingcap/tidb) **star:20143** TiDB是一个分布式SQL数据库。灵感来自谷歌F1的设计。   ![star > 5000][Gold]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [influxdb](https://github.com/influxdb/influxdb) **star:17155** 可伸缩的数据存储，用于指标衡量、事件和实时分析。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [cockroach](https://github.com/cockroachdb/cockroach) **star:16813** 可伸缩、区域备份、事务性数据存储。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [dgraph](https://github.com/dgraph-io/dgraph) **star:10508** 可伸缩、分布式、低延迟、高吞吐量的图形数据库。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [bolt](https://github.com/boltdb/bolt) **star:10005** K/V 数据库。   ![star > 5000][Gold]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [groupcache](https://github.com/golang/groupcache) **star:7674** Groupcache是一个缓存和缓存填充库，在许多情况下，它是memcached的替代品。   ![star > 5000][Gold]
* [badger](https://github.com/dgraph-io/badger) **star:6304** 快速 K/V 存储。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [rqlite](https://github.com/rqlite/rqlite) **star:4710** 基于SQLite的轻量级分布式关系数据库。   ![star > 1000][Silver]
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3174** 在Go中实现[LevelDB](https://github.com/google/leveldb) key/value数据库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3082** Ledisdb是一种高性能的NoSQL，类似于基于LevelDB的Redis。   ![star > 1000][Silver]
* [go-cache](https://github.com/pmylund/go-cache) **star:2923** 基于内存的 K/V 存储/缓存 : (类似于Memcached)，适用于单机应用程序。   ![star > 1000][Silver]
* [BigCache](https://github.com/allegro/bigcache) **star:2466** 高效的键/值缓存为千兆字节的数据。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [buntdb](https://github.com/tidwall/buntdb) **star:2444** 基于内存的K/V，快速，可嵌入的数据库，可自定义索引和空间支持。   ![star > 1000][Silver]
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2369** 属于你的NoSQL数据库。   ![star > 1000][Silver]
* [cache2go](https://github.com/muesli/cache2go) **star:1052** 基于内存的 K/V 缓存，支持超时的自动失效。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1026** 开源，快速，可伸缩的时间序列数据库。支持PromQL。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [GCache](https://github.com/bluele/gcache) **star:906** 支持过期缓存、LFU、LRU和ARC的缓存库。   ![star > 100][Bronze]
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:891** Nutsdb是一个用纯Go编写的简单、快速、可嵌入、持久的键/值存储。它支持完全序列化的事务和许多数据结构，如列表、集合、排序集。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:878** 区块链领域的一个SQL数据库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [diskv](https://github.com/peterbourgon/diskv) **star:755** 支持磁盘备份的可持久化 K/V 存储。   ![star > 100][Bronze]
* [moss](https://github.com/couchbase/moss) **star:718** Moss是一个用100% Go编写的简单LSM键值存储引擎。   ![star > 100][Bronze]
* [eliasdb](https://github.com/krotik/eliasdb) **star:532** 无其他依赖项，支持REST API，短语搜索和sql类似的查询语言的事务图数据库。   ![star > 100][Bronze]
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:491** 基于内存的快速线程安全的缓存，可缓存大量的条目。最大限度地减少GC开销。   ![star > 100][Bronze]
* [levigo](https://github.com/jmhodges/levigo) **star:364** 实现了对LevelDB封装。   ![star > 100][Bronze]
* [pudge](https://github.com/recoilme/pudge) **star:218** 使用Go的标准库编写的快速和简单的键/值存储。   ![star > 100][Bronze]
* [piladb](https://github.com/fern4lvarez/piladb) **star:171** 基于堆栈数据结构的轻量级RESTful数据库引擎。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Vasto](https://github.com/chrislusf/vasto) **star:151** 分布式高性能键值存储。可做磁盘备份。最终一致。高可用。能够在不中断服务的情况下增长或收缩。   ![star > 100][Bronze]
* [slowpoke](https://github.com/recoilme/slowpoke) **star:86** 具有持久性的键值存储。
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:61** 小型平面文件JSON存储。
* [couchcache](https://github.com/codingsince1985/couchcache) **star:40** 由 Couchbase服务 支持的RESTful缓存微服务。
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:30** BigCache 支持集群和独立且生命周期存储项。   ![最近一年没有更新][Yellow]
* [bcache](https://github.com/iwanbk/bcache) **star:29** 基于内存的最终一致的分布式缓存。
* [cache](https://github.com/akyoto/cache) **star:16** 基于内存的 K/V 存储:带生命周期的值存储，0个依赖项，<100 LoC, 100%覆盖率。
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:13** 用于临时数据存放的 K/V 存储。   ![最近一年没有更新][Yellow]
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:8** 用 Go 对[RocksDB](https://rocksdb.org)实现了封装。   ![最近一年没有更新][Yellow]

*数据库迁移。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [migrate](https://github.com/golang-migrate/migrate) **star:2779** 基于CLI的数据库迁移库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1416** 数据库迁移工具。允许使用go-bindata将迁移嵌入到应用程序中。   ![star > 1000][Silver]
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:335** 面向Gorm ORM的数据库 schema 迁移辅助程序。   ![star > 100][Bronze]
* [goose](https://github.com/steinbacher/goose) **star:119** 数据库迁移工具。您可以通过创建增量SQL或Go脚本来管理数据库的升级。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [darwin](https://github.com/GuiaBolso/darwin) **star:83** 用于数据库 schema 升级的库。
* [migrator](https://github.com/lopezator/migrator) **star:32** 非常简单的 Go 数据库迁移库。   ![最近一周有更新][Green]
* [gondolier](https://github.com/emvi/gondolier) **star:26** 使用结构修饰的数据库迁移库。
* [pravasan](https://github.com/pravasan/pravasan) **star:24** 简易的迁移工具-目前只支持MySQL，但计划很快支持Postgres, SQLite, MongoDB等。
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:23** 用Go -pg/pg编写的迁移包。
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:20** 类似 Django fixture，用于 Go 建立内置数据库/sql库。
* [avro](https://github.com/khezen/avro) **star:6** 发现SQL schemas并将其转换为AVRO schemas。   ![最近一周有更新][Green]

*数据库工具。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [vitess](https://github.com/youtube/vitess) **star:8476** vitess提供了可以为大规模web服务扩展MySQL数据库提供便利的服务和工具。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [pgweb](https://github.com/sosedoff/pgweb) **star:5998** 基于web的PostgreSQL数据库浏览器。   ![star > 5000][Gold]
* [kingshard](https://github.com/flike/kingshard) **star:4624** kingshard 是基于 Golang 的MySQL高性能代理。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [orchestrator](https://github.com/github/orchestrator) **star:3049** MySQL复制拓扑管理器和可视化工具。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2410** 自动将MySQL数据同步到Elasticsearch中。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [pREST](https://github.com/nuveo/prest) **star:2092** 基于PostgreSQL database的RESTful API服务。   ![star > 1000][Silver]
* [go-mysql](https://github.com/siddontang/go-mysql) **star:1898**  Go 工具集，用于处理MySQL协议和复制。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [chproxy](https://github.com/Vertamedia/chproxy) **star:309** ClickHouse数据库的HTTP代理。   ![star > 100][Bronze]
* [myreplication](https://github.com/2tvenom/myreplication) **star:141** MySql二进制日志复制监听器。支持基于语句和行的复制。   ![star > 100][Bronze]
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:139** 收集小的 insterts 并向 ClickHouse 服务器发送大请求。   ![star > 100][Bronze]
* [octillery](https://github.com/knocknote/octillery) **star:54** 用于数据库分表(支持每个ORM或原生SQL)。
* [dbbench](https://github.com/sj14/dbbench) **star:30** 数据库基准测试工具，支持多个数据库和脚本。
* [prep](https://github.com/hexdigest/prep) **star:24** 在不更改代码的情况下使用准备好的SQL语句。   ![最近一年没有更新][Yellow]
* [rwdb](https://github.com/andizzle/rwdb) **star:10** rwdb为多个数据库服务器的设置提供读取副本功能。   ![最近一年没有更新][Yellow]
* [datagen](https://github.com/codingconcepts/datagen) **star:8** 一个快速的数据生成器，支持多表感知和多行DML。

*SQL查询生成器，用于构建和使用SQL的库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2332** 帮助您构建SQL查询的Go库。   ![star > 1000][Silver]
* [xo](https://github.com/knq/xo) **star:2189** 基于现有的schema定义和自定义查询生成 Go 代码，基于支持PostgreSQL、MySQL、SQLite、Oracle和Microsoft SQL Server。   ![star > 1000][Silver]
* [gendry](https://github.com/didi/gendry) **star:767** 非入侵的SQL构建器和强大的数据绑定器。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [goqu](https://github.com/doug-martin/goqu) **star:639** 常用的SQL生成器和查询库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Dotsql](https://github.com/gchaincl/dotsql) **star:442** Go library帮助您将sql文件保存在一个地方，并轻松地使用它们。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:435** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   ![star > 100][Bronze]
* [sqrl](https://github.com/elgris/sqrl) **star:179** SQL查询生成器，从Squirrel fork而来，并再此基础上对性能做了优化。   ![star > 100][Bronze]
* [Squalus](https://gitlab.com/qosenergy/squalus)  Go SQL中间层，能使得执行查询更加容易。
* [scaneo](https://github.com/variadico/scaneo) **star:149** 生成用于将数据库行转换为任意结构体的 Go 代码。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [igor](https://github.com/galeone/igor) **star:78** PostgreSQL的抽象层，支持高级功能和类似gorm的语法。   ![最近一年没有更新][Yellow]
* [ormlite](https://github.com/pupizoid/ormlite)  包含一些类似orm的特性和sqlite数据库的辅助程序的轻量级包
* [godbal](https://github.com/xujiajun/godbal) **star:50** 数据库抽象层(dbal)。支持SQL builder，轻松获取结果。
* [dbq](https://github.com/rocketlaunchr/dbq) **star:45** Go的数据库操作没有样板文件。

## 数据库驱动程序

*用于连接和操作数据库的库。 (翻译出错了? 试试 [英文版](README_EN.md#database-drivers) 吧~)*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:8130** MySQL驱动程序。   ![star > 5000][Gold]
    * [pq](https://github.com/lib/pq) **star:5203** 纯 Go 的Postgres驱动。   ![star > 5000][Gold]   ![最近一周有更新][Green]
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3432** SQLite3驱动程序。   ![star > 1000][Silver]
    * [pgx](https://github.com/jackc/pgx) **star:1969** PostgreSQL驱动，支持比现有database/sql更多的特性。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1025** 微软MSSQL驱动程序。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [go-oci8](https://github.com/mattn/go-oci8) **star:406** Oracle 驱动程序。   ![star > 100][Bronze]
    * [goracle](https://github.com/go-goracle/goracle) **star:240** 基于 ODPI-C 的 Oracle 驱动程序   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:104** Firebird RDBMS SQL驱动程序。   ![star > 100][Bronze]
    * [go-adodb](https://github.com/mattn/go-adodb) **star:91** Microsoft ActiveX对象数据库驱动程序。
    * [gofreetds](https://github.com/minus5/gofreetds) **star:90** 基于[FreeTDS](http://www.freetds.org)封装的微软MSSQL Go 驱动。
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:35** Apache Avatica/Phoenix SQL驱动程序。   ![最近一周有更新][Green]
    * [bgc](https://github.com/viant/bgc) **star:12** BigQuery 的数据存储连接。   ![最近一周有更新][Green]

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:6553** 基于 Go 的 Redis 客户端。   ![star > 5000][Gold]   ![最近一周有更新][Green]
    * [redigo](https://github.com/gomodule/redigo) **star:6315** Redigo 是基于 Go 的Redis 客户端。   ![star > 5000][Gold]
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:3142** 官方的 MongoDB 驱动。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [mgo](https://github.com/globalsign/mgo) **star:1649** (已停止维护) MongoDB驱动。   ![star > 1000][Silver]
    * [gorethink](https://github.com/dancannon/gorethink) **star:1462** RethinkDB 驱动。   ![star > 1000][Silver]
    * [neoism](https://github.com/jmcvetta/neoism) **star:357** Golang 的 Neo4j 客户端。   ![star > 100][Bronze]
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:305** Aerospike 客户端。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [gocb](https://github.com/couchbase/gocb) **star:293** 官方Couchbase Go SDK。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [gocql](http://gocql.github.io)  Apache Cassandra 的 Go 驱动。
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:292** Couchbase客户端。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [redeo](https://github.com/bsm/redeo) **star:261** 与 redis 协议兼容的 TCP 服务器/服务。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:90** 实现了基于 Redigo 客户端的redislabs' ReJSON 模块。可简单地将结构体存储为JSON对象并对其进行操作。
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:72** Neo4j REST 客户端。   ![最近一年没有更新][Yellow]
    * [arangolite](https://github.com/solher/arangolite) **star:65** 轻量级的 ArangoDB 驱动。
    * [dynago](https://github.com/underarmour/dynago) **star:65** 满足 principle of least surprise 的 DynamoDB 客户端。   ![最近一年没有更新][Yellow]
    * [godis](https://github.com/piaohao/godis) **star:62** 由GoLang实现的redis客户端，灵感来自jedis。
    * [go-couchdb](https://github.com/fjl/go-couchdb) **star:51** 基于 Go 的 Yet another CouchDB 的 Http 接口封装。
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:31**  Pilosa客户端。
    * [forestdb](https://github.com/couchbase/goforestdb) **star:29** 基于 Go 的 ForestDB 接口实现。   ![最近一年没有更新][Yellow]
    * [goriak](https://github.com/zegl/goriak) **star:24**  Riak KV 驱动。
    * [neo4j](https://github.com/cihangir/neo4j) **star:24** Neo4j Rest API实现。   ![最近一年没有更新][Yellow]
    * [xredis](https://github.com/shomali11/xredis) **star:9** 类型安全，可定制，清晰和易用的Redis客户端。
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache客户端库。
    * [godscache](https://github.com/defcronyke/godscache) **star:6** 谷歌云平台Go Datastore包的封装，它采用了memcached添加缓存。
    * [asc](https://github.com/viant/asc) **star:4** Aerospike 的数据存储连接器。

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:5858** 基于 Go 的现代文本索引库。   ![star > 5000][Gold]   ![最近一周有更新][Green]
    * [riot](https://github.com/go-ego/riot) **star:4703** 基于 Go 的 开源、分布式、简单高效的搜索引擎。   ![star > 1000][Silver]   ![最近一周有更新][Green]   ![包含中文文档][CN]
    * [elastic](https://github.com/olivere/elastic) **star:4170** Elasticsearch 客户端。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:1591** 官方 Elasticsearch 客户端。   ![star > 1000][Silver]
    * [elastigo](https://github.com/mattbaird/elastigo) **star:951** Elasticsearch 客户端。   ![star > 100][Bronze]
    * [elasticsql](https://github.com/cch123/elasticsql) **star:398** 将 SQL 转换为 elasticsearch dsl。   ![star > 100][Bronze]
    * [skizze](https://github.com/seiflotfy/skizze) **star:68** 面向概率数据结构的服务和存储。   ![最近一年没有更新][Yellow]
    * [goes](https://github.com/OwnLocal/goes) **star:24** 实现了与 Elasticsearch 交互的库。   ![最近一年没有更新][Yellow]

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:12694** 图形数据库，支持多个后端。   ![star > 5000][Gold]
    * [cachego](https://github.com/fabiorphp/cachego) **star:110** 基于多个驱动程序的缓存组件。   ![star > 100][Bronze]
    * [gokv](https://github.com/philippgille/gokv) **star:86** 可扩展的简单的 K/V 存储(Redis、Consul、etcd、bbolt、BadgerDB、LevelDB、Memcached、DynamoDB、S3、PostgreSQL、MongoDB、CockroachDB等等)。
    * [dsc](https://github.com/viant/dsc) **star:14** 面向 SQL、NoSQL、结构化文件的数据存储连接。

## 日期和时间

*用于处理日期和时间的库。 (翻译出错了? 试试 [英文版](README_EN.md#date-and-time) 吧~)*

* [now](https://github.com/jinzhu/now) **star:2187** now 是时间有关的工具类。   ![star > 1000][Silver]
* [dateparse](https://github.com/araddon/dateparse) **star:897** 可以解析很多格式不固定的日期字符串。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [carbon](https://github.com/uniplaces/carbon) **star:340** 简单的时间扩展，包含了许多使用方法，从 PHP Carbon 库移植的。   ![star > 100][Bronze]
* [durafmt](https://github.com/hako/durafmt) **star:241** 轻量级、可让time.Duration更加易读的库。   ![star > 100][Bronze]
* [timeutil](https://github.com/leekchan/timeutil) **star:169** 面向 Golang 的时间库，集成了很多有用的扩展(Timedelta, Strftime, ...)。   ![star > 100][Bronze]
* [iso8601](https://github.com/relvacode/iso8601) **star:68** 不用正则表达式有效解析 ISO8601 日期时间。
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:65** 太阳历实现。
* [timespan](https://github.com/SaidinWoT/timespan) **star:60** 用于处理时间间隔相关的库，可定义开始时间和持续时间。
* [date](https://github.com/rickb777/date) **star:28** 增加处理日期、日期范围、时间跨度、时间段和time-of-day。
* [feiertage](https://github.com/wlbr/feiertage) **star:22** 用于计算德国公共假期的函数，比如复活节、感恩节等
* [goweek](https://github.com/grsmv/goweek) **star:18** 用于处理以星期实体单位的库。   ![最近一年没有更新][Yellow]
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:13** 计算指定位置的日出和日落时间。   ![最近一年没有更新][Yellow]
* [kair](https://github.com/GuilhermeCaruso/kair) **star:11** 用于处理日期和时间的 golang 库。
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:9** 可以允许 time.Time 为null。   ![最近一年没有更新][Yellow]
* [tuesday](https://github.com/osteele/tuesday) **star:7** Ruby-compatible Strftime function。   ![最近一年没有更新][Yellow]
* [strftime](https://github.com/awoodbeck/strftime) **star:5** C99-compatible strftime formatter。   ![最近一年没有更新][Yellow]

## 分布式系统

*协助构建分布式系统的包。 (翻译出错了? 试试 [英文版](README_EN.md#distributed-systems) 吧~)*

* [go-kit](https://github.com/go-kit/kit) **star:14507** 支持服务发现、负载平衡、插件式传输、请求跟踪等功能的Microservice toolkit。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [grpc-go](https://github.com/grpc/grpc-go) **star:9163** gRPC的Go语言实现。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [micro](https://github.com/micro/micro) **star:6615** 可插拔的微服务 toolkit 和分布式系统平台。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [NATS](https://github.com/nats-io/gnatsd) **star:6365** 轻量级、高性能消息传递系统，可用于微服务、物联网(IoT)和云。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [rpcx](https://github.com/smallnest/rpcx) **star:3824** 分布式可插拔的RPC服务框架，如阿里巴巴Dubbo。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [tendermint](https://github.com/tendermint/tendermint) **star:3160** 一个高性能中间件，可将任何语言的状态机转换为 Byzantine Fault 状态机。使用 Tendermint 一致性及区块链协议。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [torrent](https://github.com/anacrolix/torrent) **star:2872** BitTorrent 客户端。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Raft consensus协议的实现。 by CoreOS。
* [raft](https://github.com/hashicorp/raft) **star:2851** Raft consensus协议的实现。 by HashiCorp。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [dragonboat](https://github.com/lni/dragonboat) **star:2549** 一个功能齐全，高性能的库集。   ![star > 1000][Silver]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [glow](https://github.com/chrislusf/glow) **star:2535** 全部用 Go 实现，易用、可伸缩，可用于分布式大数据处理，Map-Reduce, DAG执行。   ![star > 1000][Silver]
* [gleam](https://github.com/chrislusf/gleam) **star:2106** 使用纯Go和Luajit编写的快速、可伸缩的分布式map/reduce系统，结合了Go的高并发性和Luajit的高性能，可以独立运行或分布式运行。   ![star > 1000][Silver]
* [emitter-io](https://github.com/emitter-io/emitter) **star:1941** 高性能、分布式、安全和低延迟的发布-订阅平台，使用MQTT、Websockets和love构建。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [KrakenD](https://github.com/devopsfaith/krakend) **star:1772** 具有中间件的高性能API网关框架。   ![star > 1000][Silver]
* [hprose](https://github.com/hprose/hprose-golang) **star:1009** 支持25+种语言RPC库。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [ringpop-go](https://github.com/uber/ringpop-go) **star:575** 可伸缩的，容错、应用分层的的Go应用程序。   ![star > 100][Bronze]
* [gorpc](https://github.com/valyala/gorpc) **star:554** 简单、快速和可伸缩的RPC库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-health](https://github.com/InVisionApp/go-health) **star:482** 用于在服务中启用异步依赖项健康检查的库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [digota](https://github.com/digota/digota) **star:301** 基于 grpc 的电子商务微服务。   ![star > 100][Bronze]
* [dot](https://github.com/dotchain/dot/)  基于 transformation/OT 的分布式同步。
* [sleuth](https://github.com/ursiform/sleuth) **star:300** 用于HTTP服务之间进行无中心p2p自动发现和RPC通信的库(使用[ZeroMQ](https://github.com/zeromq/libzmq))。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-jump](https://github.com/dgryski/go-jump) **star:254** 提供了谷歌的 “Jump” 一致哈希函数接口。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [consistent](https://github.com/buraksezer/consistent) **star:194** Consistent hashing with bounded loads。   ![star > 100][Bronze]
* [resgate](https://resgate.io/)  用于构建REST、实时和RPC API的实时API网关，其中所有客户端都是无缝同步的。
* [redis-lock](https://github.com/bsm/redis-lock) **star:147** 基于redis的分布式锁简易实现。   ![star > 100][Bronze]   ![归档项目][Archived]
* [dht](https://github.com/anacrolix/dht) **star:130** BitTorrent Kademlia DHT的实现。   ![star > 100][Bronze]
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:113** jsonrpc 包，实现了 JSON-RPC 2.0。   ![star > 100][Bronze]
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:102** JSON-RPC 2.0 HTTP客户端。   ![star > 100][Bronze]
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:52** 用于对 Celery worker、任务、事件进行交互和监控的库。
* [doublejump](https://github.com/edwingeng/doublejump) **star:40** 实现了谷歌的jump consistent hash。
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed 分布式锁定实现。
* [drmaa](https://github.com/dgruber/drmaa) **star:25** 符合 DRMAA 标准的集群调度程序作业提交库。   ![最近一年没有更新][Yellow]
* [flowgraph](https://github.com/vectaport/flowgraph) **star:18** flow-based programming package。
* [dynatomic](https://github.com/tylfin/dynatomic) **star:8** 基于 DynamoDB 的 原子计数器的库。
* [pglock](https://cirello.io/pglock)  postgresql 的分布式锁定实现。
* [outboxer](https://github.com/italolelis/outboxer) **star:5** 实现了 outbox pattern Go 库。

## 电子邮件

*实现了电子邮件创建和发送。 (翻译出错了? 试试 [英文版](README_EN.md#email) 吧~)*

* [MailHog](https://github.com/mailhog/MailHog) **star:5184** 电子邮件和SMTP测试工具，对外提供了 web 和 API 接口。   ![star > 5000][Gold]
* [chasquid](https://blitiri.com.ar/p/chasquid)  用Go编写的SMTP服务器。
* [hermes](https://github.com/matcornic/hermes) **star:1621** 可生成干净的、响应式的HTML电子邮件。   ![star > 1000][Silver]
* [email](https://github.com/jordan-wright/email) **star:1095** 一个强大和灵活的电子邮件库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-imap](https://github.com/emersion/go-imap) **star:735** 用于客户端和服务器的IMAP库。   ![star > 100][Bronze]
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:521** SendGrid 的 Go语言库，用于发送电子邮件。   ![star > 100][Bronze]
* [Hectane](https://github.com/hectane/hectane) **star:167** 轻量级的SMTP客户机，提供了HTTP API。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [douceur](https://github.com/aymerick/douceur) **star:162** 在HTML邮件中支持CSS内联。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-message](https://github.com/emersion/go-message) **star:113** 用于Internet消息格式化和邮件消息的流媒体库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [smtp](https://github.com/mailhog/smtp) **star:51** SMTP服务器协议状态机。   ![最近一年没有更新][Yellow]
* [go-dkim](https://github.com/toorop/go-dkim) **star:48** DKIM库，用于签署 & 验证电子邮件。
* [go-premailer](https://github.com/vanng822/go-premailer) **star:35** 在HTML邮件中支持CSS内联。
* [Gomail](https://github.com/go-gomail/gomail/)  一个非常简单和强大的邮件发送库。

## 可嵌入的脚本语言

*在go代码中嵌入其他语言。 (翻译出错了? 试试 [英文版](README_EN.md#embeddable-scripting-languages) 吧~)*

* [otto](https://github.com/robertkrimen/otto) **star:4741** 用 Go 编写的 JavaScript 解释器。   ![star > 1000][Silver]
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:2969** 用 Go 实现的 Lua 5.1 虚拟机和编译器。   ![star > 1000][Silver]
* [go-lua](https://github.com/Shopify/go-lua) **star:1675** 用 Go 实现的 Lua 5.2 VM接口。   ![star > 1000][Silver]
* [tengo](https://github.com/d5/tengo) **star:1314** 字节码编译的脚本语言。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [anko](https://github.com/mattn/anko) **star:923** 用Go编写的解释器。   ![star > 100][Bronze]
* [go-python](https://github.com/sbinet/go-python) **star:906** CPython C-API 的 Go 接口。   ![star > 100][Bronze]
* [expr](https://github.com/antonmedv/expr) **star:706** 一个可以计算表达式的引擎。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-php](https://github.com/deuill/go-php) **star:685** PHP 的 Go 接口。   ![star > 100][Bronze]
* [go-duktape](https://github.com/olebedev/go-duktape) **star:653** 支持 Duktape JavaScript 引擎。   ![star > 100][Bronze]
* [golua](https://github.com/aarzilli/golua) **star:443** Lua C 的 Go 接口。   ![star > 100][Bronze]
* [gisp](https://github.com/jcla1/gisp) **star:428** LISP 的 Go 接口。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [agora](https://github.com/PuerkitoBio/agora) **star:322** 基于 Go 的动态类型，可嵌入的编程语言。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [gval](https://github.com/PaesslerAG/gval) **star:139** 一种用Go编写的高度可定制的表达式语言。   ![star > 100][Bronze]
* [binder](https://github.com/alexeyco/binder) **star:31** Lua接口，基于[gopher-lua](https://github.com/yuin/gopher-lua)。   ![最近一年没有更新][Yellow]
* [gentee](https://github.com/gentee/gentee) **star:27** 嵌入式脚本编程语言。   ![最近一周有更新][Green]
* [purl](https://github.com/ian-kent/purl) **star:27** 嵌入 Go 的 Perl 5.18.2。   ![最近一年没有更新][Yellow]
* [ngaro](https://github.com/db47h/ngaro) **star:19** 嵌入式 Ngaro VM实现，支持在 Retro 中运行脚本。   ![最近一年没有更新][Yellow]

## 错误处理

*处理错误的库。 (翻译出错了? 试试 [英文版](README_EN.md#error-handling) 吧~)*

* [errors](https://github.com/pkg/errors) **star:4914** 可让你很简单的进行错误处理。   ![star > 1000][Silver]
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:732** 可将一系列的错误作为一个整体来显示。   ![star > 100][Bronze]
* [errorx](https://github.com/joomcode/errorx) **star:556** 一个功能丰富的错误包，可进行堆栈跟踪、组装异常信息以及其他。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [tracerr](https://github.com/ztrue/tracerr) **star:496** 可展示错误的堆栈跟踪信息和源码片段。   ![star > 100][Bronze]
* [errlog](https://github.com/snwfdhmp/errlog) **star:185** 用于定位抛出错误的源代码(以及一些其他快速调试特性)。可插入到任何 logger 的位置。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [werr](https://github.com/txgruppi/werr) **star:11** 对错误异常进行了捕获封装，封装信息包含了调用它的文件、行和堆栈。   ![最近一年没有更新][Yellow]

## 文件

*处理文件和文件系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#files) 吧~)*

* [afero](https://github.com/spf13/afero) **star:2249** 文件系统的抽象系统。   ![star > 1000][Silver]
* [pdfcpu](https://github.com/hhrutter/pdfcpu) **star:951** PDF处理器。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [notify](https://github.com/rjeczalik/notify) **star:497** 文件系统事件通知库，具有类似于os/signal的简单API，。   ![star > 100][Bronze]
* [opc](https://github.com/qmuntal/opc) **star:57** 加载Open Packaging Conventions (OPC)文件。
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:48** 使用 tag 加载 csv 文件。
* [skywalker](https://github.com/dixonwille/skywalker) **star:48** 可以轻松地并发地遍历文件系统。   ![最近一年没有更新][Yellow]
* [stl](https://gitlab.com/russoj88/stl)  采用并行读取算法的进行读取和写入STL(立体光刻)文件的模块。
* [tarfs](https://github.com/posener/tarfs) **star:36** tar文件的实现[ FileSystem 接口](https://godoc.org/github.com/kr/fs#FileSystem)。   ![最近一年没有更新][Yellow]
* [vfs](https://github.com/C2FO/vfs) **star:24** 一组可插拔的、可扩展的和自定义的文件系统功能，用于跨越许多文件系统类型，如os、S3和GCS。   ![最近一周有更新][Green]
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:15** 加载gtfs文件。
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:11** 拷贝文件。
* [flop](https://github.com/homedepot/flop) **star:9** 文件操作库，是[GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invoc.html)的镜像。
* [checksum](https://github.com/codingsince1985/checksum) **star:6** 生成大型文件的消息摘要(如 MD5 和 SHA256)。
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:4** ExifTool 的 Go 实现，这个著名的库用于从文件(图片、PDF、office，…)中提取尽可能多的元数据(EXIF、IPTC，…)。

## 金融

*会计和财务软件包。 (翻译出错了? 试试 [英文版](README_EN.md#financial) 吧~)*

* [decimal](https://github.com/shopspring/decimal) **star:1609** 任意精度定点的十进制数。   ![star > 1000][Silver]
* [go-money](https://github.com/rhymond/go-money) **star:622** Fowler 货币模式的实现。   ![star > 100][Bronze]
* [go-finance](https://github.com/FlashBoys/go-finance) **star:536** 综合金融市场数据。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [accounting](https://github.com/leekchan/accounting) **star:488** 货币和货币格式。   ![star > 100][Bronze]
* [techan](https://github.com/sdcoffey/techan) **star:161** 拥有先进的市场分析和交易策略的技术分析库。   ![star > 100][Bronze]
* [orderbook](https://github.com/i25959341/orderbook) **star:70** 限购单匹配引擎。
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:62** 查询 OFX 服务器和/或解析响应。
* [vat](https://github.com/dannyvankooten/vat) **star:60** 增值税编号验证 & 欧盟增值税税率。
* [transaction](https://github.com/claygod/transaction) **star:55** 嵌入式事务数据库，可多线程模式运行。
* [go-finance](https://github.com/alpeb/go-finance) **star:42** 用于货币时间价值(年金)、现金流、利率转换、债券和折旧计算的金融函数库。   ![最近一年没有更新][Yellow]
* [currency](https://github.com/bnkamalesh/currency) **star:9** 高性能、准确的货币计算软件包。

## 表单

*用于处理表单的库。 (翻译出错了? 试试 [英文版](README_EN.md#forms) 吧~)*

* [nosurf](https://github.com/justinas/nosurf) **star:973** CSRF保护中间件。   ![star > 100][Bronze]
* [binding](https://github.com/mholt/binding) **star:754** 将来自 net/HTTP 请求的表单、JSON 数据绑定到结构体。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:440** 用于Go web应用程序和服务的CSRF保护。   ![star > 100][Bronze]
* [form](https://github.com/go-playground/form) **star:350**  将 url 中的数据解析到 Go 变量中，以及将 Go 语言变量编码进 url。支持 Dual Array 及 Full map。   ![star > 100][Bronze]
* [conform](https://github.com/leebenson/conform) **star:174** 控制用户输入。基于struct tags可对数据进行修剪、清理和擦除。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [formam](https://github.com/monoculum/formam) **star:128** 将表单的值解码为 struct。   ![star > 100][Bronze]
* [forms](https://github.com/albrow/forms) **star:105** 与框架无关的库，用于解析和验证支持多部分表单和文件的表单/JSON数据。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [bind](https://github.com/robfig/bind) **star:23** 将表单数据与任意 Go 变量进行绑定。   ![最近一年没有更新][Yellow]

## 方法

*在Go中支持函数式编程的包。 (翻译出错了? 试试 [英文版](README_EN.md#functional) 吧~)*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1069** 常用辅助方法集合。   ![star > 1000][Silver]
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:107** 提供函数式编程功能。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [fuego](https://github.com/seborama/fuego) **star:39** Functional Experiment in Go。

## 游戏开发

*很棒的游戏开发库。 (翻译出错了? 试试 [英文版](README_EN.md#game-development) 吧~)*

* [Leaf](https://github.com/name5566/leaf) **star:3072** 轻量级游戏服务器框架。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [Pixel](https://github.com/faiface/pixel) **star:2452** 手工制作的 2D 游戏库。   ![star > 1000][Silver]
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:1860** 很简单的 2D 游戏库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [goworld](https://github.com/xiaonanln/goworld) **star:1202** 可伸缩的游戏服务器引擎，具有 space-entity 框架和 hot-swapping 功能。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1160** 实现了[Simple DirectMedia Layer](https://www.libsdl.org/)。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [engo](https://github.com/EngoEngine/engo) **star:1083** 开源 2D 游戏引擎。它遵循 Entity-Component-System 范式。   ![star > 1000][Silver]
* [gonet](https://github.com/xtaci/gonet) **star:1051** 实现了游戏服务器骨架。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [termloop](https://github.com/JoelOtter/termloop) **star:1026** 基于终端的 Go 游戏引擎，建立在 Termbox 之上。   ![star > 1000][Silver]
* [nano](https://github.com/lonng/nano) **star:1002** 轻量级、方便、高性能的基于golang的游戏服务器框架。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [g3n](https://github.com/g3n/engine) **star:758**  3D游戏引擎。   ![star > 100][Bronze]
* [Oak](https://github.com/oakmound/oak) **star:628** 纯 Go 实现的游戏引擎。   ![star > 100][Bronze]
* [Azul3D](https://github.com/azul3d/engine) **star:426** 用Go编写的 3D 游戏引擎。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:383** 实现了 [raylib](http://www.raylib.com/)，一个简单易用的库，用于学习视频游戏编程。   ![star > 100][Bronze]
* [go-astar](https://github.com/beefsack/go-astar) **star:327** 实现了A\*路径查找算法。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:313** 用 OpenGL 编写的 2D 游戏引擎。   ![star > 100][Bronze]
* [Pitaya](https://github.com/topfreegames/pitaya) **star:301** 可伸缩的游戏服务器框架，支持集群和客户端库的iOS, Android, Unity。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go3d](https://github.com/ungerik/go3d) **star:164** 以性能为主的2D/3D数学相关包。   ![star > 100][Bronze]
* [glop](https://github.com/runningwild/glop) **star:77** Glop (Game Library Of Power) 是一个相当简单的跨平台游戏库。   ![最近一年没有更新][Yellow]
* [go-collada](https://github.com/GlenKelley/go-collada) **star:12** 处理Collada文件。   ![最近一年没有更新][Yellow]

## 代码生成与泛型

*增强语言的工具，例如通过代码生成支持泛型。 (翻译出错了? 试试 [英文版](README_EN.md#generation-and-generics) 吧~)*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1800** 提供类似 .NET LINQ 的查询方法。   ![star > 1000][Silver]
* [jennifer](https://github.com/dave/jennifer) **star:1278** 不使用模板生成任意 Go 代码。   ![star > 1000][Silver]
* [gen](https://github.com/clipperhouse/gen) **star:1039** 用于生成泛型等类似方法的功能代码生成工具。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [goderive](https://github.com/awalterschulze/goderive) **star:747** 从输入类型来派生函数。   ![star > 100][Bronze]
* [GoWrap](https://github.com/hexdigest/gowrap) **star:266** 使用简单模板为 Go 接口生成装饰器。   ![star > 100][Bronze]
* [interfaces](https://github.com/rjeczalik/interfaces) **star:188** 用于生成接口定义的命令行工具。   ![star > 100][Bronze]
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:87** 用于包作用域反射的 Go 预处理器。   ![最近一年没有更新][Yellow]
* [go-enum](https://github.com/abice/go-enum) **star:86** 从代码注释中生成枚举。
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:43** 代码生成工具，可以不通过内存分配就可以高效的将interface{}转换为不可变类型，。   ![最近一年没有更新][Yellow]
* [gotype](https://github.com/wzshiming/gotype) **star:23** Golang 源码解析，用法类似reflect(反射)。   ![包含中文文档][CN]
* [generis](https://github.com/senselogic/GENERIS) **star:18** 提供泛型、free-form 宏、条件编译和HTML模板的代码生成工具。

## 地理

*地理工具和服务器 (翻译出错了? 试试 [英文版](README_EN.md#geographic) 吧~)*

* [Tile38](https://github.com/tidwall/tile38) **star:6343** 具有空间索引和实时地理定位功能的地理定位数据库。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [S2 geometry](https://github.com/golang/geo) **star:891** S2 geometry 库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [geocache](https://github.com/melihmucuk/geocache) **star:111** 基于内存缓存的的地理位置的应用程序。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [osm](https://github.com/paulmach/osm) **star:68** 用于读取、写入和处理 OpenStreetMap 数据和 APIs。
* [geoserver](https://github.com/hishamkaram/geoserver) **star:25** 基于geoserver REST API的 geoserver 实例。
* [gismanager](https://github.com/hishamkaram/gismanager) **star:19** 将你的 GIS 数据(矢量数据)发布到 PostGIS 和 Geoserver。
* [pbf](https://github.com/maguro/pbf) **star:16** 基于Golang 的 OpenStreetMap PBF 编码器/解码器。

## Go 编译器

*可将 Go 转换为其他语言的编译工具。 (翻译出错了? 试试 [英文版](README_EN.md#go-compilers) 吧~)*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:8557** 将 Go 编译成 JavaScript。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [llgo](https://github.com/go-llvm/llgo) **star:990** 基于 llvm 的编译器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:393** Golang 转换为 Haxe，再转换为 CPP/CSharp/Java/JavaScript 的编译器.   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [c4go](https://github.com/Konstantin8105/c4go) **star:157** 将 C 代码转换为 Go 代码。   ![star > 100][Bronze]
* [f4go](https://github.com/Konstantin8105/f4go) **star:11** 将 FORTRAN 77 转换为 Go代码。

## Goroutines

*管理和处理 Goroutines 的工具。 (翻译出错了? 试试 [英文版](README_EN.md#goroutines) 吧~)*

* [goworker](https://github.com/benmanns/goworker) **star:2254** 基于 go 的后台 worker。   ![star > 1000][Silver]
* [ants](https://github.com/panjf2000/ants) **star:1906** 一个高性能的协程池。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [tunny](https://github.com/Jeffail/tunny) **star:1354** golang 的协程池。   ![star > 1000][Silver]
* [grpool](https://github.com/ivpusic/grpool) **star:500** 轻量级协程池。   ![star > 100][Bronze]
* [pool](https://github.com/go-playground/pool) **star:482** 有限消费者协程或无限协程池，可用于更加简单的处理和取消协程   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-floc](https://github.com/workanator/go-floc) **star:167** 轻松编排 goroutines。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [workerpool](https://github.com/gammazero/workerpool) **star:138** 限制任务执行并发数，而不是队列中的任务数量的协程池，。   ![star > 100][Bronze]
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:105** 控制 goroutines 的执行顺序。   ![star > 100][Bronze]
* [semaphore](https://github.com/kamilsk/semaphore) **star:75** 信号量模式实现，可根据通道和上下文进行具备超时功能的锁定/解锁操作。   ![最近一周有更新][Green]
* [GoSlaves](https://github.com/themester/GoSlaves) **star:75** 简单异步的协程池。
* [semaphore](https://github.com/marusama/semaphore) **star:72** 基于 CAS 的可快速调整的信号量实现(比基于通道的信号量实现更快)。
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:56** manages a resizeable pool of context-aware goroutines to bound concurrency
* [worker-pool](https://github.com/vardius/worker-pool) **star:45** 一个简单的 Go 异步工作池。
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:36** 基于 Go 的 CyclicBarrier 实现。
* [breaker](https://github.com/kamilsk/breaker) **star:34** 灵活的机制，可以使执行流可中断。   ![最近一周有更新][Green]
* [gollback](https://github.com/vardius/gollback) **star:27** 异步简单的函数实用程序，用于管理闭包和回调的执行。
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:25** 并行运行函数。   ![最近一年没有更新][Yellow]
* [async](https://github.com/studiosol/async) **star:23** 一种异步执行函数的安全方法，在出现 panic 时恢复它们。
* [threadpool](https://github.com/shettyh/threadpool) **star:19** Golang 的 threadpool 实现。
* [Hunch](https://github.com/AaronJan/Hunch) **star:13** Hunch 提供了诸如 All、First、Retry、Waterfall 等功能，这使得异步流控制更加直观。
* [oversight](https://cirello.io/oversight)  完整的实现了Erlang supervision trees。
* [artifex](https://github.com/borderstech/artifex) **star:12** 简单的内存作业队列。
* [stl](https://github.com/ssgreg/stl) **star:8** 基于软件事务内存(STM)并发控制机制的软件事务锁。
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** 轻量级的协程池库，可以通过简单的API来管理。
* [go-trylock](https://github.com/subchen/go-trylock) **star:4** 支持 read-write 锁。   ![最近一年没有更新][Yellow]
* [routine](https://github.com/x-mod/routine) **star:3** go routine control with context, support: Main, Go, Pool and some useful Executors.
* [queue](https://github.com/AnikHasibul/queue) **star:2** 提供类似队列组可访问性 sync.WaitGroup 的功能。帮助你节流和限制协程、等待所有协程结束以及更多功能。

## GUI

*用于构建GUI应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

*工具包 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [ui](https://github.com/andlabs/ui) **star:6976** 跨平台的 Platform-native GUI 库。   ![star > 5000][Gold]
* [Wails](https://wails.app)  Mac, Windows, Linux桌面应用程序，主要基于含有内置的OS HTML渲染器的HTML UI。
* [fyne](https://github.com/fyne-io/fyne) **star:6319** 为 Go 而设计的跨平台的本地GUIs，使用EFL呈现。支持 : Linux, macOS, Windows。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [qt](https://github.com/therecipe/qt) **star:6083** 实现了 Qt 的 Go接口(支持Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi)。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [webview](https://github.com/zserge/webview) **star:4676** 跨平台webview窗口，具有简单的双向JavaScript绑定(Windows / macOS / Linux)。   ![star > 1000][Silver]
* [walk](https://github.com/lxn/walk) **star:3706** Windows应用程序库。   ![star > 1000][Silver]
* [app](https://github.com/murlokswarm/app) **star:2960** 用于创建包含了 GO, HTML 和 CSS 的应用程序。支持 MacOS, Windows 正在开发中。   ![star > 1000][Silver]
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:2679** 使用 GO 和 HTML/JS/CSS (电子驱动)进行构建跨平台 GUI 应用程序。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-gtk](http://mattn.github.io/go-gtk/)  实现了 GTK 的 Go接口。
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1449** 实现了 Sciter 的 Go 接口 : 用于现代桌面 UI 开发的可嵌入HTML/CSS/脚本引擎。可跨平台。   ![star > 1000][Silver]
* [gotk3](https://github.com/gotk3/gotk3) **star:773** 实现了 GTK3 的 Go接口。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [gowd](https://github.com/dtylman/gowd) **star:211** 跨平台、快速、简单的桌面UI开发，采用了GO, HTML, CSS和NW.js实现。   ![star > 100][Bronze]

*交互 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4437** 实现跨平台的GUI系统自动化。包含了控制鼠标、键盘等功能。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [systray](https://github.com/getlantern/systray) **star:789** 跨平台 Go 库，可在通知区放置图标和菜单。   ![star > 100][Bronze]
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:494** OSX 桌面通知库。   ![star > 100][Bronze]
* [trayhost](https://github.com/shurcooL/trayhost) **star:160** 跨平台 Go 库，可用于在主机操作系统的任务栏中放置图标。   ![star > 100][Bronze]
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:2** 实现了 libappindicator3 C库 的 Go接口。
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:1** 用于通知计算机上的任何(可插入的)活动的 OSX 库。
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier)  OSX 睡眠/唤醒通知。


## 硬件

*硬件交互相关的库、工具和教程。 (翻译出错了? 试试 [英文版](README_EN.md#hardware) 吧~)*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## 图片

*图像处理相关的库。 (翻译出错了? 试试 [英文版](README_EN.md#images) 吧~)*

* [imaginary](https://github.com/h2non/imaginary) **star:2595** 用于图像大小调整的快速、简单的HTTP微服务。   ![star > 1000][Silver]
* [bild](https://github.com/anthonynsimon/bild) **star:2552** 纯Go语言实现的图像处理算法合集。   ![star > 1000][Silver]
* [imaging](https://github.com/disintegration/imaging) **star:2537** 简单的Go图像处理包。   ![star > 1000][Silver]
* [gocv](https://github.com/hybridgroup/gocv) **star:2467** 使用OpenCV 3.3+实现的计算机视觉(ComputerVision)的Go语言包。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [ln](https://github.com/fogleman/ln) **star:2463** Go实现的3D线艺术（3D Line Art）渲染。   ![star > 1000][Silver]
* [resize](https://github.com/nfnt/resize) **star:2132** Go实现的使用常用的插值法（interpolation methods）调整图像大小的库。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [gg](https://github.com/fogleman/gg) **star:1919** 纯Go语言实现的2D渲染。   ![star > 1000][Silver]
* [pt](https://github.com/fogleman/pt) **star:1775** Go实现的路径跟踪（path tracing）引擎。   ![star > 1000][Silver]
* [svgo](https://github.com/ajstarks/svgo) **star:1335**  Go实现的SVG生成库。   ![star > 1000][Silver]
* [smartcrop](https://github.com/muesli/smartcrop) **star:1262** 为任意图片寻找合适的位置进行图片裁剪。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gift](https://github.com/disintegration/gift) **star:1212** 图像处理包。   ![star > 1000][Silver]
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1094** OpenCV库的Go bindings。   ![star > 1000][Silver]
* [picfit](https://github.com/thoas/picfit) **star:1072** Go实现的图像调整服务器。   ![star > 1000][Silver]
* [geopattern](https://github.com/pravj/geopattern) **star:1011** 由字符串创建漂亮图案的图片生成器。   ![star > 1000][Silver]
* [imagick](https://github.com/gographics/imagick) **star:981**  ImageMagick下MagickWand的C API的Go binding。   ![star > 100][Bronze]
* [bimg](https://github.com/h2non/bimg) **star:799** 使用libvips实现的快速高效的图像处理包。   ![star > 100][Bronze]
* [stegify](https://github.com/DimitarPetrov/stegify) **star:507**  Go实现的LSB隐写（LSB steganography），能够隐藏任何文件到一个图像中。   ![star > 100][Bronze]
* [mort](https://github.com/aldor007/mort) **star:365** Go语言实现的图像存储和处理服务器。   ![star > 100][Bronze]
* [govatar](https://github.com/o1egl/govatar) **star:310** 生成有趣头像的库和CMD工具。   ![star > 100][Bronze]
* [image2ascii](https://github.com/qeesung/image2ascii) **star:293** 将图像转换为ASCII码。   ![star > 100][Bronze]
* [go-nude](https://github.com/koyachi/go-nude) **star:286** Go语言实现的裸照检测工具。   ![star > 100][Bronze]
* [goimagehash](https://github.com/corona10/goimagehash) **star:215**  图像哈希处理的Go语言包。   ![star > 100][Bronze]
* [rez](https://github.com/bamiaux/rez) **star:189** 纯Go语言和SIMD实现的图像大小调整。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [img](https://github.com/hawx/img) **star:129** 图像处理工具的集合。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-cairo](https://github.com/ungerik/go-cairo) **star:85**  cairo图形库的Go binding。
* [mergi](https://github.com/noelyahan/mergi) **star:71** 用于图像处理(合并、裁剪、调整大小、水印、动画)的工具和Go库。
* [darkroom](https://github.com/gojek/darkroom) **star:61** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   ![最近一周有更新][Green]
* [go-gd](https://github.com/bolknote/go-gd) **star:50**  GD库的Go binding。   ![最近一年没有更新][Yellow]
* [gltf](https://github.com/qmuntal/gltf) **star:37** 一个高效、健壮的glTF 2.0文件读取、写入和验证器。
* [cameron](https://github.com/aofei/cameron) **star:31** 一个Go语言的头像生成器。
* [goimghdr](https://github.com/corona10/goimghdr) **star:27** Go语言实现的imghdr模块用于确定文件的图像类型。
* [steganography](https://github.com/auyer/steganography) **star:26** 纯Go实现的LSB隐写（LSB steganography）的库。
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Python下webcolors库的Go语言调用。   ![最近一年没有更新][Yellow]
* [tga](https://github.com/ftrvxmtrx/tga) **star:24** tga包是一个TARGA图像的解码、编码器。   ![最近一年没有更新][Yellow]
* [mpo](https://github.com/donatj/mpo) **star:6** MPO三维照片的解码和转换工具。

## 物联网

*物联网设备编程库。 (翻译出错了? 试试 [英文版](README_EN.md#iot-(internet-of-things)) 吧~)*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1134** Flogo是一个面向物联网边缘应用和集成的开源框架。   ![star > 1000][Silver]
* [gatt](https://github.com/paypal/gatt) **star:814** Gatt是一个用于构建低能耗蓝牙外围设备的Go语言包。   ![star > 100][Bronze]
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot是一个用于机器人、物理计算和物联网的框架。
* [mainflux](https://github.com/Mainflux/mainflux) **star:594** 工业物联网消息和设备管理服务器。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [periph](https://periph.io/)  外围设备I/O与低级板(low-level board)设备接口。
* [devices](https://github.com/goiot/devices) **star:225** 一套用于物联网设备的库，实验性地用于x/exp/io。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:180** 轻量级物联网流处理引擎。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [connectordb](https://github.com/connectordb/connectordb) **star:169** 自我量化和物联网的开源平台。   ![star > 100][Bronze]
* [huego](https://github.com/amimof/huego) **star:110** 一个包含广泛的Philips Hue客户端的Go语言库。   ![star > 100][Bronze]
* [iot](https://github.com/vaelen/iot/)  IoT是一个实现谷歌物联网核心设备的简单框架。
* [eywa](https://github.com/xcodersun/eywa) **star:36** Eywa是一个用于跟踪连接的设备连接管理器。   ![最近一年没有更新][Yellow]

## 作业调度器

*用于作业调度的库。 (翻译出错了? 试试 [英文版](README_EN.md#job-scheduler) 吧~)*

* [gron](https://github.com/roylee0704/gron) **star:636** 使用简单的Go API定义基于时间的任务。 之后Gron的调度程序将运行它们。   ![star > 100][Bronze]
* [JobRunner](https://github.com/bamzi/jobrunner) **star:573** 智能和功能丰富的cron作业调度程序（包含任务队列和实时监控）。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [jobs](https://github.com/albrow/jobs) **star:452** 持久和灵活的后台作业库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [scheduler](https://github.com/carlescere/scheduler) **star:295** Cronjobs让调度变得很简单。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [clockwerk](http://github.com/onatm/clockwerk)  使用简单、流畅的语法调度作业的Go语言库。
* [go-cron](https://github.com/rk/go-cron) **star:178** 一个Go实现的简单的定时任务库。可以在不同的时间间隔（每秒一次到在每年在特定的日期执行）执行闭包或函数。主要用于web应用和长时间运行的守护进程。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [clockwork](https://github.com/whiteShtef/clockwork) **star:77** Go实现的简单直观的作业调度库。
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:36** 支持webhook、crons和经典调度的作业调度程序。

## JSON

*用于JSON处理的库。 (翻译出错了? 试试 [英文版](README_EN.md#json) 吧~)*

* [GJSON](https://github.com/tidwall/gjson) **star:4928** 使用一行代码获取JSON的值。   ![star > 1000][Silver]
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2032** 从JSON自动生成Go的结构（struct）定义。   ![star > 1000][Silver]
* [gojq](https://github.com/elgs/gojq) **star:140** Go语言实现的JSON请求。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [kazaam](https://github.com/Qntfy/kazaam) **star:133** 用于任意JSON文档转换的API。   ![star > 100][Bronze]
* [jsongo](https://github.com/ricardolonga/jsongo) **star:92** 使用Fluent API来更容易地创建Json对象。   ![最近一年没有更新][Yellow]
* [gjo](https://github.com/skanehira/gjo) **star:62** 用于创建JSON对象的小工具。
* [jsonf](https://github.com/miolini/jsonf) **star:55** 用于高亮展示和查询JSON的控制台工具。   ![最近一年没有更新][Yellow]
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:40** 用Go编写的JSON比对工具。
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  将JSON转换为Go的结构（struct）。
* [mp](https://github.com/sanbornm/mp) **star:33** 简单的cli电子邮件解析器。它目前接受stdin并输出JSON。   ![最近一年没有更新][Yellow]
* [go-respond](https://github.com/nicklaw5/go-respond) **star:23** 用于处理常见的HTTP JSON响应的Go语言库。
* [ajson](https://github.com/spyzhov/ajson) **star:13** 为Go语言开发的包含JSONPath支持的抽象JSON。
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** 让自定义结构（struct）转化为JSON兼容的HAL（Hypertext Application Language）返回数据的简单Go包。
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:8** Go-JsonError允许我们轻松创建符合JsonApi规范的json响应错误。
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:5** 基于JSON API错误引用的Go bindings。   ![最近一年没有更新][Yellow]

## 日志记录

*用于生成和处理日志文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#logging) 吧~)*

* [logrus](https://github.com/Sirupsen/logrus) **star:12007** Go的结构化日志操作 。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [zap](https://github.com/uber-go/zap) **star:7486** 快速、结构化、多等级的日志记录。   ![star > 5000][Gold]
* [spew](https://github.com/davecgh/go-spew) **star:3312** 为Go数据结构实现一个漂亮的printer用于帮助调试。   ![star > 1000][Silver]
* [glog](https://github.com/golang/glog) **star:2307** 为Go提供了多等级日志记录。   ![star > 1000][Silver]
* [zerolog](https://github.com/rs/zerolog) **star:2241** Zero-allocation JSON日志记录。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [tail](https://github.com/hpcloud/tail) **star:1542** 努力模拟实现BSD的tail的特性的Go包。   ![star > 1000][Silver]
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1455** 简单的滚动日志，io.WriteCloser的实现。   ![star > 1000][Silver]
* [seelog](https://github.com/cihub/seelog) **star:1361** 具有灵活调度、过滤和格式化的日志功能。   ![star > 1000][Silver]
* [log15](https://github.com/inconshreveable/log15) **star:914** 简单、强大的日志操作。   ![star > 100][Bronze]
* [log](https://github.com/apex/log) **star:733** Go的结构化日志包。   ![star > 100][Bronze]
* [logxi](https://github.com/mgutz/logxi) **star:333** 12-factor app的日志程序，快速且让人高兴地使用。   ![star > 100][Bronze]
* [onelog](https://github.com/francoispqt/onelog) **star:332** Onelog是一个非常简单但非常高效的JSON日志程序。它是所有场景中速度最快的JSON日志程序。而且，它是配置要求最低的日志记录器之一。   ![star > 100][Bronze]
* [log](https://github.com/go-playground/log) **star:267** Go的简单、可配置和可伸缩的结构化日志。   ![star > 100][Bronze]
* [logutils](https://github.com/hashicorp/logutils) **star:248** Go的用于更好地进行日志操作的实用程序，继承了标准日志库。   ![star > 100][Bronze]
* [go-logger](https://github.com/apsdehal/go-logger) **star:234** 简单的日志程序的 Go 程序，与级别处理程序。   ![star > 100][Bronze]
* [logger](https://github.com/azer/logger) **star:135** Go的精简日志库。   ![star > 100][Bronze]
* [xlog](https://github.com/rs/xlog) **star:128** 针对'net/context`实现的结构化的记录器，用于HTTP处理程序。   ![star > 100][Bronze]
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:108** 支持日志多等级、分类和过滤的高性能日志记录。可以发送过滤后的日志消息到各种目标(如控制台，网络，邮件)。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![包含中文文档][CN]
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:97** RollingWriter是一个自动循环的io.Writer的实现,带有多种策略以提供日志文件循环(rotation)。
* [log-voyage](https://github.com/firstrow/logvoyage) **star:82** 用Go编写的功能齐全的日志写入saas。   ![最近一年没有更新][Yellow]
* [glg](https://github.com/kpango/glg) **star:51** glg是一个简单而快速的Go日志库。   ![最近一周有更新][Green]
* [stdlog](https://github.com/alexcesaro/log) **star:43** Stdlog是一个面向对象的库，提供了多等级日志记录。它对cron任务非常有用。   ![最近一年没有更新][Yellow]
* [gologger](https://github.com/sadlil/gologger) **star:38** 为Go提供方便简单的日志操作; 在彩色控制台，简单控制台，文件或Elasticsearch上。   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j的Go语言。   ![最近一年没有更新][Yellow]
* [logex](https://github.com/chzyer/logex) **star:32** 由标准日志库封装的Go日志库，支持跟踪和多等级。   ![最近一年没有更新][Yellow]
* [logrusly](https://github.com/sebest/logrusly) **star:26** [logrus](https://github.com/sirupsen/logrus)的插件，将错误信息发送到[Loggly](https://www.loggly.com/)。   ![最近一年没有更新][Yellow]
* [go-log](https://github.com/siddontang/go-log) **star:23** 支持多等级和多处理程序的日志库。
* [log](https://github.com/teris-io/log) **star:22** Go的结构化日志接口，清晰地将日志facade与其实现(implementation)分离开来。   ![最近一年没有更新][Yellow]
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** 简单的go日志模块，有5个级别，可选循环(rotation)日志文件记录功能和stdout/stderr输出。   ![最近一年没有更新][Yellow]
* [journald](https://github.com/ssgreg/journald) **star:19**  Go实现systemd Journal的原生API用于日志记录。
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:19** 基于当前日期和时间的自动日志文件写入工具，类似cronolog。
* [distillog](https://github.com/amoghe/distillog) **star:19** distilled日志记录(可以将其视为stdlib +日志)。   ![最近一年没有更新][Yellow]
* [gone/log](https://github.com/One-com/gone/tree/master/log)  快速、可扩展、功能齐全、std-lib源兼容的日志库。
* [gomol](https://github.com/aphistic/gomol) **star:15** 为Go实现可多方式输出、结构化日志, 并可扩展日志输出方式。
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** 用于多等级级日志记录的包。   ![最近一年没有更新][Yellow]
* [go-log](https://github.com/subchen/go-log) **star:9** Go实现的简单且可配置的日志工具，并带有多等级、日志格式化和日志写入工具。   ![最近一年没有更新][Yellow]
* [glo](https://github.com/lajosbencz/glo) **star:8** 参照PHP的Monolog实现的具有相同日志等级的Go日志库。
* [xlog](https://github.com/xfxdev/xlog) **star:7** 插件架构和灵活的日志系统，带有多级别、多日志目标和自定义日志格式。
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Go的彩色日志记录器，带有可配置的日志级别。
* [logo](https://github.com/mbndr/logo) **star:5** Go的日志工具，可配置的日志写入器。   ![最近一年没有更新][Yellow]
* [log](https://github.com/aerogo/log) **star:4** 一个O(1)日志系统，允许您将一个日志连接到多个日志写入(例如stdout、文件和TCP连接)。

## 机器学习

*机器学习库。 (翻译出错了? 试试 [英文版](README_EN.md#machine-learning) 吧~)*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:6656** 通用机器学习库。   ![star > 5000][Gold]   ![包含中文文档][CN]
* [gorgonia](https://github.com/chewxy/gorgonia) **star:2699** 基于图形（graph-based）的计算库，如Theano：它为构建各种机器学习和神经网络算法提供了基本框架。   ![star > 1000][Silver]
* [tfgo](https://github.com/galeone/tfgo) **star:1191** 易于使用的Tensorflow bindings:简化了官方Tensorflow Go bindings的使用。在Go中定义计算图形，在Python中加载和执行训练的模型。   ![star > 1000][Silver]
* [goml](https://github.com/cdipaolo/goml) **star:1011** 在线机器学习。   ![star > 1000][Silver]
* [gosseract](https://github.com/otiai10/gosseract) **star:879** 使用c++的Tesseract库实现的OCR。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:644** 快速、灵活、多线程集成的决策树，用于机器学习。   ![star > 100][Bronze]
* [bayesian](https://github.com/jbrukh/bayesian) **star:630** Go的朴素贝叶斯分类。   ![star > 100][Bronze]
* [eaopt](https://github.com/MaxHalford/eaopt) **star:624** 一个进化优化（evolutionary optimization）库。   ![star > 100][Bronze]
* [gorse](https://github.com/zhenghaoz/gorse) **star:537** 基于协同过滤（Collaborative Filtering ）的高性能推荐系统包。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [gobrain](https://github.com/goml/gobrain) **star:388** 用 Go 编写的神经网络库。   ![star > 100][Bronze]
* [regommend](https://github.com/muesli/regommend) **star:249** 推荐和协同过滤引擎。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [ocrserver](https://github.com/otiai10/ocrserver) **star:225** 一个简单的OCR API服务器，非常容易地使用Docker和Heroku部署。   ![star > 100][Bronze]
* [go-deep](https://github.com/patrikeh/go-deep) **star:220** 一个功能丰富的神经网络库 。   ![star > 100][Bronze]
* [go-galib](https://github.com/thoj/go-galib) **star:172** 用Go编写的遗传算法库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:150** Go Interface， 用于开放式神经网络交换(Open Neural Network Exchange)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:144** 用Go编写的推荐算法库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [shield](https://github.com/eaigner/shield) **star:124** 贝叶斯文本分类器，具有灵活的tokenizers和存储后端。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-fann](https://github.com/white-pony/go-fann) **star:99** 快速人工神经网络(FANN)库的Go bindings。   ![最近一年没有更新][Yellow]
* [Goptuna](https://github.com/c-bata/goptuna) **star:80** 黑盒函数的贝叶斯优化框架。一切都会优化。   ![最近一周有更新][Green]
* [goga](https://github.com/tomcraven/goga) **star:78** Go的遗传算法库。   ![最近一年没有更新][Yellow]
* [libsvm](https://github.com/datastream/libsvm) **star:63** 基于LIBSVM 3.14实现。   ![最近一年没有更新][Yellow]
* [neural-go](https://github.com/schuyler/neural-go) **star:61** 多层感知器网络在Go中的实现，使用反向传播算法进行训练。   ![最近一年没有更新][Yellow]
* [go-pr](https://github.com/daviddengcn/go-pr) **star:57** Go编写的模式识别包。   ![最近一年没有更新][Yellow]
* [neat](https://github.com/jinyeom/neat) **star:55** 即插即用的并行Go框架，用于增强拓扑的神经进化(NeuroEvolution of Augmenting Topologies)。   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [golinear](https://github.com/danieldk/golinear) **star:39**  Go实现的liblinear bindings。
* [goscore](https://github.com/asafschers/goscore) **star:37**  为预言模型标记语言（PMML）实现的评分API。
* [fonet](https://github.com/Fontinalis/fonet) **star:33** 一个用Go编写的深度神经网络库。   ![最近一周有更新][Green]
* [godist](https://github.com/e-dard/godist) **star:24** 各种概率分布，以及相关的method。   ![最近一年没有更新][Yellow]
* [Varis](https://github.com/Xamber/Varis) **star:23** Go实现的神经网络。   ![最近一年没有更新][Yellow]
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:21** Go实现的k-modes和k-prototypes聚类算法。   ![最近一年没有更新][Yellow]
* [probab](https://github.com/ThePaw/probab) **star:10** 概率分布函数。贝叶斯推理。使用Go写的。   ![最近一年没有更新][Yellow]
* [evoli](https://github.com/khezen/evoli) **star:8** 遗传算法（Genetic Algorithm）和粒子群优化（Particle Swarm Optimization）库。
* [GoMind](https://github.com/surenderthakran/gomind) **star:6** 一个简单的神经网络库。   ![最近一年没有更新][Yellow]

## 消息

*实现消息传递系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#messaging) 吧~)*

* [sarama](https://github.com/Shopify/sarama) **star:4657**  Apache Kafka的Go库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gorush](https://github.com/appleboy/gorush) **star:3722** 使用[APNs2](https://github.com/sideshow/apns2)和谷歌[GCM](https://github.com/google/go-gcm)推送通知服务器。   ![star > 1000][Silver]
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:3692** 实时消息(Websockets或SockJS)服务器。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [machinery](https://github.com/RichardKnop/machinery) **star:3374** 基于分布式消息传递的异步任务/作业队列。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:2892** go的socket.io库，一个实时应用程序框架。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [NATS Go Client](https://github.com/nats-io/nats) **star:2406** 轻量级和高性能的发布-订阅(publish-subscribe)和分布式队列消息传递系统——这是一个Go库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [APNs2](https://github.com/sideshow/apns2) **star:2046** HTTP / 2苹果消息推送provider——发送推送消息到iOS, tvOS, Safari和OSX应用。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Benthos](https://github.com/Jeffail/benthos) **star:2015** 一系列协议之间的消息流桥接。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1839** gopush-cluster是一个gopush服务器集群。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]   ![包含中文文档][CN]
* [melody](https://github.com/olahol/melody) **star:1568** 处理websocket session的极简框架，包括广播和自动ping/pong处理。   ![star > 1000][Silver]
* [mangos](https://github.com/go-mangos/mangos) **star:1535** Nanomsg(“可伸缩协议”)的纯go实现,具有传输互操作性。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Mercure](https://github.com/dunglas/mercure) **star:1514** 使用Mercure协议分派服务器发送(server-sent)更新的服务器和库(构建在服务器发送事件之上)。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1460** NSQ的官方Go包。   ![star > 1000][Silver]
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1098** Redis支持的统一推送服务, 用于服务端向移动设备的消息通知。   ![star > 1000][Silver]
* [zmq4](https://github.com/pebbe/zmq4) **star:775** ZeroMQ的Go interface第4版。也可用于[第3版](https://github.com/pebbe/zmq3)和[第2版](https://github.com/pebbe/zmq2)。   ![star > 100][Bronze]
* [Gollum](https://github.com/trivago/gollum) **star:767** 一个n:m多路复用器(n:m multiplexer)，收集不同来源的消息并将其广播到一组目的地。   ![star > 100][Bronze]
* [Beaver](https://github.com/Clivern/Beaver) **star:728** 一个实时消息服务器，可用于在web和手机app端构建一个可伸缩的应用内通知，多人游戏，聊天应用。   ![star > 100][Bronze]
* [EventBus](https://github.com/asaskevich/EventBus) **star:563** 具有异步兼容性的轻量级事件总线。   ![star > 100][Bronze]
* [golongpoll](https://github.com/jcuga/golongpoll) **star:429** HTTP longpoll服务器库，使web发布-订阅变得简单。   ![star > 100][Bronze]
* [dbus](https://github.com/godbus/dbus) **star:360** D-Bus的Go bindings。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Glue](https://github.com/desertbit/glue) **star:320** 健壮的Go和Javascript Socket库(替代Socket.io)。   ![star > 100][Bronze]
* [emitter](https://github.com/olebedev/emitter) **star:310** 使用Go的方式发出事件, 带有通配符、谓词、取消可能性和许多其他优点。   ![star > 100][Bronze]
* [pubsub](https://github.com/tuxychandru/pubsub) **star:279** 简单的pubsub的go包。   ![star > 100][Bronze]
* [guble](https://github.com/smancke/guble) **star:139** 使用推送通知服务(谷歌Firebase云消息、苹果推送通知服务、SMS)的消息服务器, 支持websockets,REST API，并具有分布式操作和消息持久性。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Bus](https://github.com/mustafaturan/bus) **star:115** 内部通信的最小消息总线实现。   ![star > 100][Bronze]
* [oplog](https://github.com/dailymotion/oplog) **star:94** 用于REST api的通用oplog/replication系统。   ![最近一年没有更新][Yellow]
* [rabtap](https://github.com/jandelgado/rabtap) **star:72** RabbitMQ的瑞士军刀cli应用。   ![最近一周有更新][Green]
* [messagebus](https://github.com/vardius/message-bus) **star:68** messagebus是一个Go的简单异步消息总线，非常适合在执行事件sourcing、CQRS和DDD时用作事件总线。
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:62** amqp exchanges和队列上的一个小工具。
* [drone-line](https://github.com/appleboy/drone-line) **star:60** 使用二进制、docker或Drone CI发送[Line](https://at.line.me/en)通知。
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:55** RapidMQ是用于管理本地消息队列的轻量且可靠的库。   ![最近一年没有更新][Yellow]
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:51** 一个围绕NSQ topic和channel的小工具。   ![最近一年没有更新][Yellow]
* [go-notify](https://github.com/TheCreeper/go-notify) **star:47** 原生的freedesktop通知规范实现。
* [goose](https://github.com/ian-kent/goose) **star:36** 服务器在Go中发送事件。   ![最近一年没有更新][Yellow]
* [event](https://github.com/agoalofalife/event) **star:27** 观察者模式的实现。   ![最近一年没有更新][Yellow]
* [hub](https://github.com/leandro-lugaresi/hub) **star:25** 适用于Go应用程序的消息/事件中心，使用publish/subscribe模式，并支持别名(类似rabbitMQ exchanges)。   ![最近一年没有更新][Yellow]
* [Commander](https://github.com/jeroenrinzema/commander) **star:21** 高级事件驱动的消费者/生产者(consumer/producer)，支持各种“方言”，如Apache Kafka。   ![最近一周有更新][Green]
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:11** 用于Viessmann Vitotrol web服务的客户端库。
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** 用Go编写的Gaurun客户端。
* [jazz](https://github.com/socifi/jazz) **star:6** 一个简单的RabbitMQ抽象层，用于队列管理和消息的发布和消费。
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:2** 基于Redis streams的队列的生产者和消费者。
* [rmqconn](https://github.com/sbabiv/rmqconn)  RabbitMQ重新连接。amqp.Connection和amqp.Dial之上的Wrapper。允许在强制调用Close()方法关闭连接之前重新连接。

## 微软办公软件

* [unioffice](https://github.com/unidoc/unioffice) **star:1734** 用于创建和处理Office Word (.docx)、Excel (.xlsx)和Powerpoint (.pptx)文档的纯go库。   ![star > 1000][Silver]

### Microsoft Excel

*用于操作Microsoft Excel的库。 (翻译出错了? 试试 [英文版](README_EN.md#microsoft-excel) 吧~)*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:4500** 用于读写Microsoft Excel™(XLSX)文件的Go语言库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [xlsx](https://github.com/tealeg/xlsx) **star:3442** 用以简化在Go程序中读取使用最新版本Microsoft Excel的XML格式文件的库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [xlsx](https://github.com/plandem/xlsx) **star:74** 在Go程序以快速和安全的方式读取/更新现有的Microsoft Excel文件。
* [go-excel](https://github.com/szyhf/go-excel) **star:46** 一个简单轻便的阅读器，可以将类关系型数据库(relate-db-like)的excel作为表来读取。
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:12** libxlsxwriter的Go binding, 用于编写XLSX (Microsoft Excel)文件。   ![最近一年没有更新][Yellow]

## 杂项

### 依赖注入

*用于处理依赖项注入的库。 (翻译出错了? 试试 [英文版](README_EN.md#dependency-injection) 吧~)*

* [dig](https://github.com/uber-go/dig) **star:918** 一个基于反射的Go依赖注入工具包。   ![star > 100][Bronze]
* [fx](https://github.com/uber-go/fx) **star:756** 基于依赖注入的Go应用程序框架(构建在dig之上)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [alice](https://github.com/magic003/alice) **star:34** Go的外挂的依赖注入容器。   ![最近一年没有更新][Yellow]
* [inject](https://github.com/defval/inject) **star:28** 一个基于反射的依赖注入容器，具有简单的接口。
* [wire](https://github.com/Fs02/wire) **star:19** 适用于Go的严格运行时依赖注入(Strict Runtime Dependency Injection)。
* [gocontainer](https://github.com/vardius/gocontainer) **star:9** 简单的依赖注入容器。
* [linker](https://github.com/logrange/linker) **star:9** A reflection based dependency injection and inversion of control library with components lifecycle support.

### 项目布局

*用于组织项目的非正式模式集。 (翻译出错了? 试试 [英文版](README_EN.md#project-layout) 吧~)*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:9246** Go生态系统中历史和新兴的项目布局模式集合。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [scaffold](https://github.com/catchplay/scaffold) **star:26** 快速生成Go项目布局的脚手架。让您专注于已实现的业务逻辑。
* [go-sample](https://github.com/zitryss/go-sample) **star:25** 带有实际代码的Go应用程序项目的示例布局。

### 字符串

*处理字符串的库。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [xstrings](https://github.com/huandu/xstrings) **star:621** 从其他语言移植的有用字符串函数合集。   ![star > 100][Bronze]
* [strutil](https://github.com/ozgio/strutil) **star:64** 字符串处理工具。

*这些库之所以放在这里，是因为不适合放在其他分类。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [gopsutil](https://github.com/shirou/gopsutil) **star:3947** 用于检索进程和系统利用率(CPU、内存、磁盘等)的跨平台的库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [archiver](https://github.com/mholt/archiver) **star:2495** 用于生成和解压.zip和.tar.gz文档的库和命令。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gosms](https://github.com/haxpax/gosms) **star:1227** Go编写的私人的本地短信网关，可以用来发送短信。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:860**  Go语言弹性模式(resiliency pattern)。   ![star > 100][Bronze]
* [go-openapi](https://github.com/go-openapi)  用于解析和使用开放api模式(open-api schemas)的包的集合。
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:674** Go语言的通用对象池。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:630** base64Captcha支持数字，字母，算术，音频和混合模式的验证码。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [shortid](https://github.com/teris-io/shortid) **star:451** 分布式地生成超短、唯一、非顺序、URL友好的id。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:428** 用go编写的随机数据生成器。   ![star > 100][Bronze]
* [llvm](https://github.com/llir/llvm) **star:412** 用于在纯Go中与LLVM IR交互的库。   ![star > 100][Bronze]
* [health](https://github.com/dimiro1/health) **star:362** 易于使用，可扩展的健康检查库。   ![star > 100][Bronze]
* [conv](https://github.com/cstockton/go-conv) **star:342** conv包提供了跨Go类型(Go types)的快速和直观的转换。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [banner](https://github.com/dimiro1/banner) **star:235** 在Go应用程序中添加漂亮的横幅(banner)。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gountries](https://github.com/pariz/gountries) **star:211** 获取国家和细节数据的包。   ![star > 100][Bronze]
* [antch](https://github.com/antchfx/antch) **star:142** 一个快速、强大和可扩展的web爬虫框架。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![包含中文文档][CN]
* [battery](https://github.com/distatus/battery) **star:135** 跨平台、标准化的电池信息库。   ![star > 100][Bronze]
* [ffmt](https://github.com/go-ffmt/ffmt) **star:127** 美化数据,使其更适合人查看。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [stats](https://github.com/go-playground/stats) **star:121** Monitors Go MemStats + 诸如如内存，Swap和CPU的系统状态统计，并通过UDP发送到任何你想记录的地方   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [lk](https://github.com/hyperboloide/lk) **star:121** 一个简单的版权许可证库。   ![star > 100][Bronze]
* [bitio](https://github.com/icza/bitio) **star:97** 高度优化的位级读写器。   ![最近一年没有更新][Yellow]
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:85** 用于RESTful服务的强制的(opinionated)并发健康检查HTTP处理程序。
* [turtle](https://github.com/hackebrot/turtle) **star:77** Go的Emojis库。   ![最近一年没有更新][Yellow]
* [gommit](https://github.com/antham/gommit) **star:75** 分析git提交消息，确保它们遵循已定义的格式。
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:68** 用于RAR、TAR、ZIP和7z文件的解压缩库。
* [indigo](https://github.com/osamingo/indigo) **star:52** 分布式唯一ID生成器, 使用Sonyflake并由Base58编码。
* [morse](https://github.com/alwindoss/morse) **star:50** 实现字符串与摩尔斯电码转换的库。
* [captcha](https://github.com/steambap/captcha) **star:43** captcha包为验证码生成提供了一个易于使用的、unopinionated的API。
* [xkg](https://github.com/go-xkg/xkg) **star:40** X Keyboard Grabber(键盘事件捕获)   ![最近一年没有更新][Yellow]
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:34** 通过Json请求生成PDF的HTTP服务。   ![最近一年没有更新][Yellow]
* [persian](https://github.com/mavihq/persian) **star:32** 一些适用于波斯语的Go工具库。   ![最近一年没有更新][Yellow]
* [ghorg](https://github.com/gabrie30/ghorg) **star:31** GitHub一个组织中所有的仓库复制到一个目录中。   ![最近一周有更新][Green]
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:30** 用于[Browser Capabilities Project](http://browscap.org/)的Go库。
* [datacounter](https://github.com/miolini/datacounter) **star:28** 用于readers/writer/http.ResponseWriter的计数器。
* [autoflags](https://github.com/artyom/autoflags) **star:24** 从struct字段自动定义命令行flag的Go包。
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg)规范在Go中的实现。
* [gosh](https://github.com/osamingo/gosh) **star:17** 提供Go统计处理程序，结构和测量方法。
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  用于生成http输入和输出处理模板。
* [url-shortener](https://github.com/pantrif/url-shortener) **star:17** 一个现代的、强大的、健壮的URL转短链接微服务，带有mysql支持。   ![最近一年没有更新][Yellow]
* [gotoprom](https://github.com/cabify/gotoprom) **star:14** 为Prometheus客户端提供类型安全的指标(metric)构建工具库。
* [sandid](https://github.com/aofei/sandid) **star:12** 能沟让地球上的每一粒沙子都有自己的ID。
* [anagent](https://github.com/mudler/anagent) **star:11** Go语言的最小化，可插入的evloop/timer处理程序, 带有依赖注入。   ![最近一年没有更新][Yellow]
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:9** 根据威尔逊得分排序算法(Wilson Score Equation)计算平均分和评分。   ![最近一年没有更新][Yellow]
* [hostutils](https://github.com/Wing924/hostutils) **star:8** 一个用于打包和解包FQDNs列表的go语言库。
* [shellwords](https://github.com/Wing924/shellwords) **star:7** 一个Go库，实现了依照UNIX Bourne shell的关键词解析规则进行字符串操纵。   ![最近一年没有更新][Yellow]
* [metrics](https://github.com/pascaldekloe/metrics) **star:4** 用于instrumentation和Prometheus exposition的库。
* [numa](https://github.com/lrita/numa) **star:2** NUMA是一个用go编写的实用程序库。它可以帮助我们编写一些NUMA-AWARED代码。

## 自然语言处理

*用于处理人类语言的库。 (翻译出错了? 试试 [英文版](README_EN.md#natural-language-processing) 吧~)*

* [prose](https://github.com/jdkato/prose) **star:2050** 用于支持标记化、词性标记、名称实体提取等文本处理的库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gse](https://github.com/go-ego/gse) **star:1078** 高效的文本分割;支持英语、汉语、日语等。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [when](https://github.com/olebedev/when) **star:927** 带有可插入规则的自然EN和RU语言日期/时间解析器。   ![star > 100][Bronze]
* [gojieba](https://github.com/yanyiwu/gojieba) **star:823** 这是一个Go实现的[jieba](https://github.com/fxsjy/jieba)，这是一个中文分词算法。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:526** 中文汉字到汉语拼音的转换。   ![star > 100][Bronze]
* [kagome](https://github.com/ikawaha/kagome) **star:417** JP形态学分析仪编写的纯Go。   ![star > 100][Bronze]
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:352** Go的自然语言检测包。支持84种语言和24种脚本(书写系统，如拉丁语、西里尔语等)。   ![star > 100][Bronze]
* [nlp](https://github.com/Shixzie/nlp) **star:352** 从字符串中提取值并用nlp填充结构。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [sentences](https://github.com/neurosnap/sentences) **star:262** 句子标记器:将文本转换为句子列表。   ![star > 100][Bronze]
* [nlp](https://github.com/james-bowman/nlp) **star:216** 支持LSA(潜在语义分析)的Go自然语言处理库。   ![star > 100][Bronze]
* [go-nlp](https://github.com/nuance/go-nlp) **star:79** 用于处理离散概率分布的实用程序和用于进行NLP工作的其他工具。   ![最近一年没有更新][Yellow]
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  软件包和用于处理本地化文本的附带工具。
* [getlang](https://github.com/rylans/getlang) **star:74** 快速自然语言检测包。
* [gounidecode](https://github.com/fiam/gounidecode) **star:67** 用于Go的Unicode音译器(也称为unidecode)。   ![最近一年没有更新][Yellow]
* [textcat](https://github.com/pebbe/textcat) **star:61** Go package支持基于n-gram的文本分类，支持utf-8和原始文本。   ![最近一年没有更新][Yellow]
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** 这是一个 Go 实现的[MMSEG](http://technology.chtsai.org/mmseg/)，这是一个中文分词算法。   ![最近一年没有更新][Yellow]
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:56** Unicode文本的ASCII音译。
* [go-stem](https://github.com/agonopol/go-stem) **star:52** 波特词干算法的实现。   ![最近一年没有更新][Yellow]
* [stemmer](https://github.com/dchest/stemmer) **star:47** 用于Go编程语言的Stemmer包。包括英语和德语词根。   ![最近一年没有更新][Yellow]
* [segment](https://github.com/blevesearch/segment) **star:46** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   ![最近一年没有更新][Yellow]
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:45** 快速自动关键字提取算法(RAKE)的Go端口。
* [porter2](https://github.com/zhenjl/porter2) **star:33** 非常快的波特2史坦默。   ![最近一年没有更新][Yellow]
* [go2vec](https://github.com/danieldk/go2vec) **star:30** 用于word2vec嵌入式的阅读器和实用程序函数。
* [paicehusk](https://github.com/rookii/paicehusk) **star:25** Golang实现了Paice/外壳阻塞算法。   ![最近一年没有更新][Yellow]
* [snowball](https://github.com/goodsign/snowball) **star:24** 滚雪球柄端口(cgo包装)为 Go 。提供词干提取功能[Snowball native](http://snowball.tartarus.org/)。   ![最近一年没有更新][Yellow]
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo绑定到Yandex。Mystem -俄罗斯形态学分析仪。   ![最近一年没有更新][Yellow]
* [petrovich](https://github.com/striker2000/petrovich) **star:22** 彼得罗维奇是一个图书馆，它把俄语名字的词形变化成特定的语法格。
* [icu](https://github.com/goodsign/icu) **star:19** Cgo绑定用于icu4c C库的检测和转换功能。保证与版本50.1兼容。   ![最近一年没有更新][Yellow]
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:15**  Go 绑定斯诺鲍libstemmer库，包括波特2。   ![最近一年没有更新][Yellow]
* [shamoji](https://github.com/osamingo/shamoji) **star:10** shamoji是用Go编写的word过滤包。
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** 用于libtextcat C库的Cgo绑定。保证与版本2.2兼容。   ![最近一年没有更新][Yellow]
* [porter](https://github.com/a2800276/porter) **star:8** 这是Martin Porter在C语言中实现的Porter词干分析算法的一个相当简单的移植。   ![最近一年没有更新][Yellow]
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** 一种基于字典和双字母格朗语言模型的记号赋予器。(现在只支持中文分割)

## 网络

*用于处理各种网络层的库。 (翻译出错了? 试试 [英文版](README_EN.md#networking) 吧~)*

* [kcptun](https://github.com/xtaci/kcptun) **star:10693** 基于KCP协议的非常简单和快速udp隧道。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [fasthttp](https://github.com/valyala/fasthttp) **star:9427** 一个快速HTTP实现，比net/http快10倍。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [dns](https://github.com/miekg/dns) **star:3827** 用于 DNS 的库。   ![star > 1000][Silver]
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3415** HTTPLabs 允许你检查 HTTP 请求和伪造响应。   ![star > 1000][Silver]
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:2934** 在纯Go中实现了QUIC协议。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gopacket](https://github.com/google/gopacket) **star:2891** Go library for packet processing with libpcap bindings.   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [webrtc](https://github.com/pions/webrtc) **star:2289** WebRTC API的纯Go实现。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2250** 快速可靠的ARQ协议。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gobgp](https://github.com/osrg/gobgp) **star:1695** 基于 Go 的 BGP 实现。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [ssh](https://github.com/gliderlabs/ssh) **star:1114** 用于构建SSH服务器的高级API(封装密码/ SSH)。   ![star > 1000][Silver]
* [fortio](https://github.com/fortio/fortio) **star:889** 负载测试库和命令行工具，高级的echo服务器和web UI。允许指定一组每秒查询的负载，并记录延迟直方图和其他有用的统计数据，并将它们作图。支持Tcp、Http、gRPC。   ![star > 100][Bronze]
* [water](https://github.com/songgao/water) **star:843** 简单TUN / TAP图书馆。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [sftp](https://github.com/pkg/sftp) **star:743** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-getter](https://github.com/hashicorp/go-getter) **star:729**  通过URL来下载文件或目录。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [NFF-Go](https://github.com/intel-go/nff-go) **star:666** 用于快速开发云计算和裸机网络功能的框架(原YANFF)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  Paho Go客户端提供了一个 MQTT 客户端库，用于通过TCP、TLS或WebSockets连接到MQTT代理。
* [mdns](https://github.com/hashicorp/mdns) **star:553** 简单mDNS(Multicast DNS)客户端/服务器库。   ![star > 100][Bronze]
* [grab](https://github.com/cavaliercoder/grab) **star:546**  用于管理文件下载。   ![star > 100][Bronze]
* [ftp](https://github.com/jlaffaye/ftp) **star:530** 实现了[RFC 959](http://tools.ietf.org/html/rfc959)中描述的ftp客户端。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [lhttp](https://github.com/fanux/lhttp) **star:513** 强大的websocket框架，可以让你更容易的构建IM服务器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![包含中文文档][CN]
* [gosnmp](https://github.com/soniah/gosnmp) **star:440** 用于执行 SNMP 操作的原生 Go 库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [gotcp](https://github.com/gansidui/gotcp) **star:418** 用于快速编写 tcp 应用程序。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [cidranger](https://github.com/yl2chen/cidranger) **star:388** 快速得 IP 到 CIDR 查找。   ![star > 100][Bronze]
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:365** 基于UDP组播的跨平台本地对等点发现库。   ![star > 100][Bronze]
* [gopcap](https://github.com/akrennmair/gopcap) **star:354**  用 Go 实现了对 libpcap 的封装。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-stun](https://github.com/ccding/go-stun) **star:333** 实现了 STUN 客户端(RFC 3489和RFC 5389)。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [raw](https://github.com/mdlayher/raw) **star:304** Package raw支持在设备驱动程序级别读取和写入网络接口的数据。   ![star > 100][Bronze]
* [stun](https://github.com/go-rtc/stun) **star:287** Go实现的RFC 5389 STUN协议。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [tcp_server](https://github.com/firstrow/tcp_server) **star:285**  Go 图书馆建设tcp服务器更快。   ![star > 100][Bronze]
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:232** 通过TCP传输协议缓冲区数据。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [winrm](https://github.com/masterzen/winrm) **star:215**  Go WinRM客户端远程执行Windows机器上的命令。   ![star > 100][Bronze]
* [arp](https://github.com/mdlayher/arp) **star:197** 实现了arp协议，如RFC 826中所述。   ![star > 100][Bronze]
* [ethernet](https://github.com/mdlayher/ethernet) **star:185** 实现了对IEEE 802.3以太网II帧和IEEE 802.1Q VLAN标签的编组和解组。   ![star > 100][Bronze]
* [utp](https://github.com/anacrolix/utp) **star:149** Go uTP微传输协议的实现。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [canopus](https://github.com/zubairhamed/canopus) **star:135** CoAP客户端/服务器实现(RFC 7252)。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [jazigo](https://github.com/udhos/jazigo) **star:126** Jazigo是一个用Go编写的工具，用于检索多个网络设备的配置。   ![star > 100][Bronze]
* [sslb](https://github.com/eduardonunesp/sslb) **star:115** 它是一个超级简单的负载平衡器，只是一个实现某种性能的小项目。   ![star > 100][Bronze]
* [gNxI](https://github.com/google/gnxi) **star:104** 一组基于 gNMI 和 gNOI 协议的网络管理工具。   ![star > 100][Bronze]
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:92** Gmqtt是一个灵活、高性能的MQTT代理库，它完全实现了MQTT协议V3.1.1。   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [xtcp](https://github.com/xfxdev/xtcp) **star:84** TCP服务器框架具有同时全双工通信，优雅关机，自定义协议。
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:63** 实现了一个DHCPv6服务器，如RFC 3315所述。
* [ether](https://github.com/songgao/ether) **star:62** 一个用于发送和接收以太网帧的跨平台 Go 库。   ![最近一年没有更新][Yellow]
* [linkio](https://github.com/ian-kent/linkio) **star:44** 网络链路速度模拟，主要用于接口的读/写。   ![最近一年没有更新][Yellow]
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   ![最近一年没有更新][Yellow]
* [packet](https://github.com/aerogo/packet) **star:29** 通过TCP和UDP发送数据包。它可以缓冲消息和热交换连接。
* [iplib](https://github.com/c-robinson/iplib) **star:24** 用于处理IP地址的库(net)。借鉴于python 的 [ipaddress](https://docy-doc.org/3/library/ipaddress.html)和ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)
* [graval](https://github.com/koofr/graval) **star:24** FTP服务器框架。   ![最近一年没有更新][Yellow]
* [publicip](https://github.com/polera/publicip) **star:18** 包publicip返回面向公共的IPv4地址(internet出口)。   ![最近一年没有更新][Yellow]
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:15** 用于解码 pcap 文件和分析解剖数据。   ![最近一年没有更新][Yellow]
* [goshark](https://github.com/sunwxg/goshark) **star:9** 用于解码IP包，创建用于分析的数据结构包。   ![最近一年没有更新][Yellow]
* [llb](https://github.com/kirillDanshin/llb) **star:8** 一个非常简单、快速的代理服务器后端。可用于快速重定向到预定义域，具有零内存分配和快速响应。   ![最近一年没有更新][Yellow]
* [tspool](https://github.com/two/tspool) **star:5** TCP库使用工作池来提高性能并保护服务器。

### HTTP客户端

*用于发出HTTP请求的库。 (翻译出错了? 试试 [英文版](README_EN.md#http-clients) 吧~)*

* [grequests](https://github.com/levigross/grequests) **star:1418** 一个 Go “克隆”的伟大和著名的请求库。   ![star > 1000][Silver]
* [heimdall](https://github.com/gojektech/heimdall) **star:1084** 具有重试和hystrix功能的增强http客户机。   ![star > 1000][Silver]
* [sling](https://github.com/dghubble/sling) **star:1009** Sling是一个用于创建和发送API请求的Go HTTP客户端库。   ![star > 1000][Silver]
* [gentleman](https://github.com/h2non/gentleman) **star:678** 功能齐全的插件驱动HTTP客户端库。   ![star > 100][Bronze]
* [pester](https://github.com/sethgrid/pester) **star:332** 使用重试、后退和并发执行HTTP客户机调用。   ![star > 100][Bronze]
* [goreq](https://github.com/smallnest/goreq) **star:98** 基于gorequest的增强简化HTTP客户机。   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [rq](https://github.com/ddo/rq) **star:27** golang stdlib HTTP客户端更好的接口。   ![最近一周有更新][Green]

## OpenGL

*用于在Go中使用OpenGL的库。 (翻译出错了? 试试 [英文版](README_EN.md#opengl) 吧~)*

* [glfw](https://github.com/go-gl/glfw) **star:736** GLFW 3 的 Go 接口实现。   ![star > 100][Bronze]
* [gl](https://github.com/go-gl/gl) **star:644** OpenGL 的 Go 接口实现(通过glow生成)。   ![star > 100][Bronze]
* [mathgl](https://github.com/go-gl/mathgl) **star:291** 完全基于 Go 实现的数学软件包，专门用于处理三维数学。借鉴于 GLM。   ![star > 100][Bronze]
* [goxjs/gl](https://github.com/goxjs/gl) **star:129** 跨平台的OpenGL 接口实现(OS X, Linux, Windows，浏览器，iOS, Android)。   ![star > 100][Bronze]
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:57** 跨平台 glfw 库，可用于创建 OpenGL 上下文并接收事件。

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques. (翻译出错了? 试试 [英文版](README_EN.md#orm) 吧~)*

* [GORM](https://github.com/jinzhu/gorm) **star:14764** 一个出色的 ORM 库。主要目标是对开发人员友好。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Xorm](https://github.com/go-xorm/xorm) **star:5229** 基于 Go 的简单而强大的ORM。   ![star > 5000][Gold]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [gorp](https://github.com/go-gorp/gorp) **star:3081** 基于 Go 的关系持久性 ORM-ish 库。   ![star > 1000][Silver]
* [go-pg](https://github.com/go-pg/pg) **star:3022** 用于 PostgreSQL 的ORM。侧重于 PostgreSQL 的特性和性能。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2278** ORM 生成器。根据数据库 schema 生成一个功能强大且运行速度快的ORM。   ![star > 1000][Silver]
* [upper.io/db](https://github.com/upper/db) **star:1857** 对外提供统一的接口用于访问不同的存储介质，例如PostgreSQL, MySQL, SQLite, MSSQL, QL、MongoDB.。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [reform](https://github.com/go-reform/reform) **star:804** 基于非空接口和代码生成的 ORM。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [pop/soda](https://github.com/gobuffalo/pop) **star:682** 数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [QBS](https://github.com/coocood/qbs) **star:541** Stands for Query By Struct. A Go ORM.   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![包含中文文档][CN]
* [go-queryset](https://github.com/jirfag/go-queryset) **star:449** 基于 GORM 100% 类型安全的 ORM。可支持 MySQL, PostgreSQL, Sqlite3, SQL Server。   ![star > 100][Bronze]
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  强大的orm框架。支持: pq/mysql/sqlite3。
* [Zoom](https://github.com/albrow/zoom) **star:240** 基于 Redis 的快速数据存储和查询引擎。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:237** 一个灵活而强大的SQL字符串构建器库。   ![star > 100][Bronze]
* [grimoire](https://github.com/Fs02/grimoire) **star:112** 基于 golang 的数据库访问层和验证库。(支持: MySQL, PostgreSQL和SQLite3)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-store](https://github.com/gosuri/go-store) **star:92** 简单且快速的 Redis 键值存储库。   ![最近一年没有更新][Yellow]
* [Marlow](https://github.com/dadleyy/marlow) **star:67** 从项目结构生成ORM。
* [lore](https://github.com/abrahambotros/lore) **star:4** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   ![最近一年没有更新][Yellow]
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:1** 一个轻量级的ORM。用于Google/Firebase Cloud Firestore。

## 包管理

*用于管理依赖和包的官方工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules 是源码的版本控制和交换的单位。go命令直接支持处理模块，包括记录和解决对其他模块的依赖关系。

*包管理的官方实验工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [dep](https://github.com/golang/dep) **star:12592**  Go 的依赖管理工具，需要 Go 1.9+   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [vgo](https://go.googlesource.com/vgo/)  Go 命令版本管理

*用于包和依赖项管理的非官方库。 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [glide](https://github.com/Masterminds/glide) **star:7788** 轻松管理您的 golang 第三方包。受Maven、Bundler和Pip等工具的启发。   ![star > 5000][Gold]
* [godep](https://github.com/tools/godep) **star:5651** godep是go的依赖工具，它通过修复包的依赖关系来帮助构建可重复的包。   ![star > 5000][Gold]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [govendor](https://github.com/kardianos/govendor) **star:4750** 包管理器。使用 vendor 文件的 Go vendor 工具。   ![star > 1000][Silver]
* [gopm](https://github.com/gpmgo/gopm) **star:2364** 包管理器。   ![star > 1000][Silver]   ![归档项目][Archived]
* [gom](https://github.com/mattn/gom) **star:1352** Go Manager - bundle for Go。   ![star > 1000][Silver]
* [gpm](https://github.com/pote/gpm) **star:1205** 基本的 Go 依赖管理器。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [goop](https://github.com/nitrous-io/goop) **star:776** Go 的简单依赖管理器，灵感来自Bundler。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [nut](https://github.com/jingweno/nut) **star:245** vendor 依赖。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** 使用Git的最小依赖版本。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gigo](https://github.com/LyricalSecurity/gigo) **star:196** 类似pip的golang依赖工具，支持私有存储库和散列。   ![star > 100][Bronze]
* [VenGO](https://github.com/DamnWidget/VenGO) **star:115** 创建和管理可导出的隔离go虚拟环境。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:87** 插件，为自动加载Golang SDK，依赖关系管理和启动Maven项目基础设施中的构建环境提供了方法。   ![最近一周有更新][Green]
* [gop](https://github.com/lunny/gop) **star:50** 在GOPATH之外构建和管理Go应用程序。   ![归档项目][Archived]   ![包含中文文档][CN]

## 性能

* [jaeger](https://github.com/jaegertracing/jaeger) **star:8714** 分布式跟踪系统。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [profile](https://github.com/pkg/profile) **star:1012** Go的简单分析支持包。   ![star > 1000][Silver]
* [tracer](https://github.com/kamilsk/tracer) **star:9** 简单、轻量级的跟踪。   ![最近一周有更新][Green]

## 查询语言

* [graphql-go](https://github.com/graphql-go/graphql) **star:5257** 为Go实现GraphQL。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [graphql](https://github.com/neelance/graphql-go) **star:2804** 关注易用性的GraphQL服务器。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:865** 一个用来查询JSON数据的Go包。   ![star > 100][Bronze]
* [jsonql](https://github.com/elgs/jsonql) **star:201** Golang中的JSON查询表达式库。   ![star > 100][Bronze]
* [rql](https://github.com/a8m/rql) **star:111** 用于REST API的资源查询语言。   ![star > 100][Bronze]
* [graphql](https://github.com/tmc/graphql) **star:51** graphql解析器+工具集   ![最近一年没有更新][Yellow]
* [jsonslice](https://github.com/bhmj/jsonslice) **star:24** 使用高级过滤器查询Jsonpath。

## 嵌入的资源

* [packr](https://github.com/gobuffalo/packr) **star:2143** 将静态文件嵌入到Go二进制文件中的简单方法。   ![star > 1000][Silver]
* [statik](https://github.com/rakyll/statik) **star:2123** 将静态文件嵌入到Go可执行文件中。   ![star > 1000][Silver]
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1659** go.rice 是一个Go包，它使处理html、js、css、图像和模板等资源变得非常容易。   ![star > 1000][Silver]
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:663** 生成一个vfsdata。静态实现给定虚拟文件系统的go文件。   ![star > 100][Bronze]
* [esc](https://github.com/mjibson/esc) **star:473** 将文件嵌入到Go程序中并提供http文件系统接口。   ![star > 100][Bronze]
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:425** 一个可定制的工具用来在Go中嵌入文件   ![star > 100][Bronze]
* [go-resources](https://github.com/omeid/go-resources) **star:155** 嵌入到Go中的普通资源。   ![star > 100][Bronze]
* [statics](https://github.com/go-playground/statics) **star:53** 将静态资源嵌入到go文件中，用于单个二进制编译+使用http。文件系统+符号链接。   ![最近一年没有更新][Yellow]
* [templify](https://github.com/wlbr/templify) **star:20** 将外部模板文件嵌入到Go代码中，以创建单个文件二进制文件。
* [go-embed](https://github.com/pyros2097/go-embed) **star:16** 生成go代码，将资源文件嵌入到库或可执行文件中。

## 科学与数据分析

*用于科学计算和数据分析的库。 (翻译出错了? 试试 [英文版](README_EN.md#science-and-data-analysis) 吧~)*

* [gonum](https://github.com/gonum/gonum) **star:2971** Gonum是一组用于Go编程语言的数字库。它包含用于矩阵、统计、优化等的库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [stats](https://github.com/montanaflynn/stats) **star:1345** 包含Golang标准库中缺少的公共函数的统计软件包。   ![star > 1000][Silver]
* [streamtools](https://github.com/nytlabs/streamtools) **star:1314** 通用图形工具，用于处理数据流。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [gosl](https://github.com/cpmech/gosl) **star:1309** 提供线性代数，FFT，几何，NURBS，数值方法，概率，优化，微分方程，等等。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [gonum/plot](https://github.com/gonum/plot) **star:1220** gonum/plot提供了一个API，用于在Go中构建和绘制绘图。   ![star > 1000][Silver]
* [go-dsp](https://github.com/mjibson/go-dsp) **star:627** Go数字信号处理。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [goraph](https://github.com/gyuho/goraph) **star:600** 纯Go图论库(数据结构，算法可视化)。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [chart](https://github.com/vdobler/chart) **star:581** 简单的图表绘制库。支持多种图形类型。   ![star > 100][Bronze]
* [ewma](https://github.com/VividCortex/ewma) **star:266** 提供指数加权移动平均算法。   ![star > 100][Bronze]
* [graph](https://github.com/yourbasic/graph) **star:236** 基本图形算法库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [orb](https://github.com/paulmach/orb) **star:199** 2D几何类型，支持剪切、GeoJSON和Mapbox矢量平铺。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:127** 数据流的近似直方图。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:78** 用于机器学习和统计的数据模型(类似于 pandas)。
* [sparse](https://github.com/james-bowman/sparse) **star:69**  Go 稀疏矩阵格式的线性代数支持科学和机器学习应用程序，兼容gonum矩阵库。
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:66** TextRank在Golang中的实现，支持扩展特性(摘要、加权、短语提取)和多线程(goroutine)。
* [pagerank](https://github.com/alixaxel/pagerank) **star:49** 加权 PageRank 算法在Go中的实现。
* [geom](https://github.com/skelterjohn/geom) **star:40**  Go 的二维几何。   ![最近一年没有更新][Yellow]
* [evaler](https://github.com/soniah/evaler) **star:40** 简单的浮点算术表达式求值器。   ![最近一年没有更新][Yellow]
* [goent](https://github.com/kzahedi/goent) **star:13**  Go 实现熵度量。
* [triangolatte](https://github.com/tchayen/triangolatte) **star:11** 二维三角库。允许将线和多边形(都基于点)转换为gpu语言。
* [ode](https://github.com/ChristopherRabotin/ode) **star:10** 常微分方程(ODE)求解器，支持扩展状态和基于信道的迭代停止条件。   ![最近一年没有更新][Yellow]
* [PiHex](https://github.com/claygod/PiHex) **star:9** 实现了针对16进制数 Pi 的“bailee - borwein - plouffe”算法。
* [GoStats](https://github.com/OGFris/GoStats) **star:9** GoStats是一个开放源码的GoLang库，主要用于机器学习领域的数学统计，它涵盖了大多数统计度量函数。
* [go-gt](https://github.com/ThePaw/go-gt) **star:6** 用“Go”语言编写的图论算法。   ![最近一年没有更新][Yellow]
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:5** 微型线性插值库。
* [bradleyterry](https://github.com/seanhagen/bradleyterry)  为成对比较提供了一个 Bradley-Terry 模型。
* [assocentity](https://github.com/ndabAP/assocentity) **star:4** assocentity 返回单词到给定实体的平均距离。   ![最近一周有更新][Green]
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** 二次函数求根算法库。

## 安全

*用于帮助您的应用程序更安全的库。 (翻译出错了? 试试 [英文版](README_EN.md#security) 吧~)*

* [lego](https://github.com/xenolf/lego) **star:3535** 纯 Go ACME 客户端库及命令行工具   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:1827** 工具和库，以远程入侵RTSP流从监控摄像头。   ![star > 1000][Silver]
* [acmetool](https://github.com/hlandau/acme) **star:1695** ACME(让我们用自动更新加密)客户端工具。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [memguard](https://github.com/awnumar/memguard) **star:1527** 一个用于处理内存中敏感值的纯Go库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [secure](https://github.com/unrolled/secure) **star:1212** Go 语言 HTTP 中间件，为 Go 提供了一些安全功能   ![star > 1000][Silver]
* [acra](https://github.com/cossacklabs/acra) **star:452** 网络加密代理保护基于数据库的应用程序免受数据泄漏:强选择性加密，SQL注入预防，入侵检测系统。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [nacl](https://github.com/kevinburke/nacl) **star:451**  Go 实现NaCL API的集合。   ![star > 100][Bronze]
* [BadActor](https://github.com/jaredfolkins/badactor) **star:247** 一个驻留在内存中的，应用驱动的监控程序，受 fail2ban 的启发   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [passlib](https://github.com/hlandau/passlib) **star:226** 不过时的密码哈希库。   ![star > 100][Bronze]
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:196** 使用ssh密钥加密/解密。   ![star > 100][Bronze]
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:155** Scrypt 库，具有简单、易懂的 API，同时具有内置的自动校准功能   ![star > 100][Bronze]
* [go-yara](https://github.com/hillu/go-yara) **star:133** YARA的 Go 语言接口，号称是 “对于恶意软件研究者（以及其他人）来说是模式匹配的瑞士军刀”   ![star > 100][Bronze]
* [argon2pw](https://github.com/raja/argon2pw) **star:76** 使用常量时间密码比较生成Argon2密码散列。
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  让我们加密证书并启动TLS服务器。
* [Interpol](https://bitbucket.org/vahidi/interpol)  基于规则的数据生成器，用于模糊和渗透测试。
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:30** 一个安全哈希和加密密码的偏执包。
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2密码散列和验证设计为与现有Python和PHP实现兼容。
* [certificates](https://github.com/mvmaasakkers/certificates) **star:9** 用于生成tls证书的自定义工具。
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:7** 使用围绕acme/autocert的高级包装器，SSL证书变得很容易。
* [jwc](https://github.com/khezen/jwc) **star:5** JSON Web加密库。

## 序列化

*用于二进制序列化的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#serialization) 吧~)*

* [jsoniter](https://github.com/json-iterator/go) **star:5567** 高性能，100% 兼容的“encoding/json” 替代品   ![star > 5000][Gold]
* [goprotobuf](https://github.com/golang/protobuf) **star:5116** 通过库和协议编译器插件使 Go 语言支持 Google的 protocol buffers.   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:2978** Go 语言的 Protocol Buffer 库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:2463** 用于对原生键值对进行解码生成 Go 语言结构体   ![star > 1000][Silver]
* [go-codec](https://github.com/ugorji/go) **star:1239** 高性能、多功能、规范化编码解码以及 rpc 库， 用于 msgpack, cbor 和 json，支持基于运行时的 OR 码生成   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [colfer](https://github.com/pascaldekloe/colfer) **star:471** 为Colfer二进制格式生成代码。   ![star > 100][Bronze]
* [csvutil](https://github.com/jszwec/csvutil) **star:306** 高性能、惯用的CSV记录编码和解码到本机Go结构。   ![star > 100][Bronze]
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:272** Go 语言用的 Cap'n Proto 库及解析器   ![star > 100][Bronze]
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:119** 用于协同 PHP session 格式数据和 PHP 序列化／反序列化函数工作的go语言库   ![star > 100][Bronze]
* [structomap](https://github.com/tuvistavie/structomap) **star:96** 用于从静态结构体简单、动态的生成键值对的库
* [bambam](https://github.com/glycerine/bambam) **star:60** 用于 Go 语言生成 Cap'n Proto schemas 的生成器   ![最近一年没有更新][Yellow]
* [asn1](https://github.com/PromonLogicalis/asn1) **star:40** 面向golang的BER和DER编码库。   ![归档项目][Archived]
* [binstruct](https://github.com/ghostiam/binstruct) **star:7** 用于将数据映射到结构中的Golang二进制解码器。
* [fwencoder](https://github.com/o1egl/fwencoder) **star:6** 用于Go的固定宽度文件解析器(编码和解码库)。   ![最近一年没有更新][Yellow]
* [bel](https://github.com/32leaves/bel) **star:5** 从Go structs/interface生成TypeScript接口。对JSON RPC很有用。

## 服务器应用程序

* [etcd](https://github.com/coreos/etcd) **star:26624** 为共享配置和服务发现提供高可用的键值存储。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Caddy](https://github.com/mholt/caddy) **star:23232** Caddy是另一种HTTP/2 web服务器，易于配置和使用。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [consul](https://www.consul.io/)  Consul 是一个用于服务发现、监控和配置的工具
* [minio](https://github.com/minio/minio) **star:17571** Minio是一个分布式对象存储服务器。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [RoadRunner](https://github.com/spiral/roadrunner) **star:3312** 高性能PHP应用服务器，负载平衡器和进程管理器。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [devd](https://github.com/cortesi/devd) **star:2810** 为开发人员提供本地web服务器。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [algernon](https://github.com/xyproto/algernon) **star:1591** 内置支持Lua、Markdown、GCSS和Amber的HTTP/2 web服务器。   ![star > 1000][Silver]
* [flipt](https://github.com/markphelps/flipt) **star:992** 一个用Go和Vue.js编写的自包含特性标志解决方案   ![star > 100][Bronze]
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:884** 功能齐全，高度可配置的SFTP服务器软件。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  小型化、网络化、基于内存的键值存储
* [Flagr](https://github.com/checkr/flagr) **star:829** Flagr是一个开源特性标记和A/B测试服务。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Fider](https://github.com/getfider/fider) **star:798** Fider是一个收集和组织客户反馈的开放平台。   ![star > 100][Bronze]
* [jackal](https://github.com/ortuman/jackal) **star:719** 用Go编写的XMPP服务器。   ![star > 100][Bronze]
* [discovery](https://github.com/Bilibili/discovery) **star:678** 用于弹性中间层负载平衡和故障转移的注册表。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:6** 从PostgreSQL到Kafka的流数据库事件。
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  传递到负载平衡Riemann事件并/或将其转换为 Carbon。
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:5** Nginx日志解析器和Prometheus 导出。
* [nsq](http://nsq.io/)  一个实时分布式消息平台。

## 流处理

*用于流处理和反应性编程的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#stream-processing) 吧~)*

* [go-streams](https://github.com/reugn/go-streams) **star:188** 去流处理库。   ![star > 100][Bronze]

## 模板引擎

*用于模板和词法分析的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#template-engines) 吧~)*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3087** PDF 文档生成器，支持文本，绘图和图片   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [pongo2](https://github.com/flosch/pongo2) **star:1491** 类似 DjanGo 的模板引擎   ![star > 1000][Silver]
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1412** 快速、强大且易用的模板引擎。将模板转化为 Go 语言并进行编译   ![star > 1000][Silver]
* [hero](https://github.com/shiyanhui/hero) **star:1206** Hero是一个方便、快速和强大的go模板引擎。   ![star > 1000][Silver]   ![包含中文文档][CN]
* [mustache](https://github.com/hoisie/mustache) **star:969** Go 语言实现的 Mustache 模板语言   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [amber](https://github.com/eknkc/amber) **star:824** Amber是一个优雅的Go编程语言模板引擎，它的灵感来自HAML和Jade。   ![star > 100][Bronze]
* [ace](https://github.com/yosssi/ace) **star:763** Ace 是一个 Go 语言的 HTML 模板引擎，受到了 Slim 和 Jade 的启发。 Ace 是对Gold的一种改进。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Razor](https://github.com/sipin/gorazor) **star:702** Go 语言的 Razor 视图引擎   ![star > 100][Bronze]
* [jet](https://github.com/CloudyKit/jet) **star:583** Jet模板引擎。   ![star > 100][Bronze]
* [ego](https://github.com/benbjohnson/ego) **star:417** 轻量级模板语言，允许您在Go中编写模板。模板被翻译成Go并编译。   ![star > 100][Bronze]
* [raymond](https://github.com/aymerick/raymond) **star:342** 使用 Go 语言实现的完整的 handlebars   ![star > 100][Bronze]
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:299** 简单快速的模板引擎。进行模板元素替换时，速度是比[text/template](http://golang.org/pkg/text/template/)快10倍。   ![star > 100][Bronze]
* [Soy](https://github.com/robfig/soy) **star:144** Go 语言实现的谷歌闭包模板(也就是 Soy templates) ,遵循[官方规范](https://developer.google.com/closure/templates/)。   ![star > 100][Bronze]
* [liquid](https://github.com/osteele/liquid) **star:84** Go 语言实现的 Shopify Liquid 模板.
* [kasia.go](https://github.com/ziutek/kasia.go) **star:70** 一个用于HTML 和其他文本文件的模板系统，使用go语言实现   ![最近一年没有更新][Yellow]
* [velvet](https://github.com/gobuffalo/velvet) **star:64** 使用 Go 语言实现的完整的 handlebars   ![最近一年没有更新][Yellow]
* [goview](https://github.com/foolin/goview) **star:46** Goview是一个轻量级、极简的模板库，基于golang html/template构建Go web应用程序。
* [damsel](https://github.com/dskinner/damsel) **star:20** 标记语言，通过css选择器实现了 html 框架 ，并可以通过 pkg html/template 等进行扩展   ![最近一年没有更新][Yellow]
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:13**  对 html/template 进行了简单的封装，支持基于文件的模板可以利用其他模板文件进行扩展   ![最近一年没有更新][Yellow]

## 测试

*用于测试代码库和生成测试数据的库。 (翻译出错了? 试试 [英文版](README_EN.md#testing) 吧~)*

* Testing Frameworks
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  将markdown代码段转换为可测试的go代码。
    * [Testify](https://github.com/stretchr/testify) **star:8188** 对标准测试包的扩展。   ![star > 5000][Gold]
    * [go-cmp](https://github.com/google/go-cmp) **star:1178** 用于比较测试中的Go值的包。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1142** 简洁的、声明式的、易用的端到端HTTP 及 REST API 测试   ![star > 1000][Silver]
    * [godog](https://github.com/DATA-DOG/godog) **star:762** 类似 Cucumber 或 Behat 的 BDD 框架   ![star > 100][Bronze]
    * [baloo](https://github.com/h2non/baloo) **star:648** 表达性强、多功能的、端到端的HTTP API 测试工具   ![star > 100][Bronze]
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD 风格的测试框架，具有 web 界面和计时刷新功能
    * [gocheck](http://labix.org/gocheck)  更加高级的测试框架，用于替换 Gotest
    * [goblin](https://github.com/franela/goblin) **star:626** 类似Mocha的测试框架。   ![star > 100][Bronze]
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:332** 类似 Rails 的测试工具，用于测试数据库应用   ![star > 100][Bronze]
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:332** 记录并回放HTTP交互，以便进行快速、确定和准确的测试。   ![star > 100][Bronze]
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:294** 变异测试的Go源代码。   ![star > 100][Bronze]
    * [gofight](https://github.com/appleboy/gofight) **star:258** 对 Go 语言的路由框架进行 API 测试   ![star > 100][Bronze]
    * [ginkgo](http://onsi.github.io/ginkgo/)  Go的BDD测试框架。
    * [frisby](https://github.com/verdverm/frisby) **star:248** REST API测试框架。   ![star > 100][Bronze]
    * [go-carpet](https://github.com/msoap/go-carpet) **star:195** 在终端中查看测试覆盖率的工具。   ![star > 100][Bronze]
    * [charlatan](https://github.com/percolate/charlatan) **star:190** 为测试生成假接口实现的工具。   ![star > 100][Bronze]
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:122** 一组包，用于增强go测试包并支持公共模式。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [GoSpec](https://github.com/orfjackal/gospec) **star:111** 用于 Go 编程语言的bdd风格的测试框架。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
    * [endly](https://github.com/viant/endly) **star:94** 声明性端到端功能测试。   ![最近一周有更新][Green]
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:87** 清空数据库用于测试，受到database_cleaner 的启发
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:85** 测试框架的简单快照测试插件。
    * [apitest](https://apitest.dev)  支持模拟外部HTTP调用和呈现序列图的基于REST的服务或HTTP处理程序的简单且可扩展的行为测试库。
    * [wstest](https://github.com/posener/wstest) **star:62** 用于单元测试Websocket http.Handler的Websocket客户机。   ![最近一年没有更新][Yellow]
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:55** 极具灵活性的golang深度比较，扩展了go测试包。
    * [gospecify](https://github.com/stesla/gospecify) **star:51** 支持 BDD 语法 。对于任何使用过 rspec 等库的人来说应该非常熟悉。   ![最近一年没有更新][Yellow]
    * [restit](https://github.com/yookoala/restit) **star:49** 帮助编写 RESTful API 集成测试的 Go 语言微型框架.。   ![最近一年没有更新][Yellow]
    * [commander](https://github.com/SimonBaeumer/commander) **star:32** 用于在windows、linux和osx上测试cli应用程序的工具。   ![最近一周有更新][Green]
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:30** 为针对模式测试JSON而创建的库。
    * [gomega](http://onsi.github.io/gomega/)  类似 Rspec 的 matcher/assertion 库
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:26** 用于声明性 Matcher 对象的连贯框架，当将其应用于输入值时，将产生自描述结果。   ![最近一年没有更新][Yellow]
    * [dsunit](https://github.com/viant/dsunit) **star:25** 用于SQL、NoSQL、结构化文件的数据存储测试。
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:22** 用于验证JSON有效负载已正确序列化的包。
    * [assert](https://github.com/go-playground/assert) **star:14** 基础断言库，用于对 Go 语言程序进行测试，提供了一些用于自定义断言的代码块   ![最近一年没有更新][Yellow]
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** 轻量级测试套，为 Go1.7's Subtests 带来了setup/teardown 功能   ![最近一年没有更新][Yellow]
    * [testcase](https://github.com/adamluzsi/testcase) **star:9** 行为驱动开发的惯用测试框架。
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Go 语言 testing/iotest 包的扩展。   ![最近一年没有更新][Yellow]
    * [gocrest](https://github.com/corbym/gocrest) **star:8** 用于 Go 断言的可组合的类似 hamcrest 的 matchers。   ![最近一年没有更新][Yellow]
    * [gogiven](https://github.com/corbym/gogiven) **star:7** 类似于 YATSPEC 的Go BDD测试框架。   ![最近一年没有更新][Yellow]
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** 在测试前从SQL文件生成测试数据，并在测试完成后清除数据。
    * [biff](https://github.com/fulldump/biff) **star:6** 分支测试框架，BDD兼容。   ![最近一年没有更新][Yellow]
    * [Tt](https://github.com/vcaesar/tt) **star:5** 简单而丰富多彩的测试工具。   ![最近一周有更新][Green]
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:2** HTTP客户端测试框架。

* Mock
    * [gomock](https://github.com/golang/mock) **star:2868** 用于Go编程语言的mock框架。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:1757** Mock SQL ，用于测试数据库交互   ![star > 1000][Silver]
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1443** 使用可扩展中间件和易于使用的CLI记录和模拟REST/SOAP api的HTTP(S)代理。   ![star > 1000][Silver]
    * [gock](https://github.com/h2non/gock) **star:824** 多功能、易用 HTTP mock   ![star > 100][Bronze]
    * [httpmock](https://github.com/jarcoal/httpmock) **star:587** 轻松模拟来自外部资源的HTTP响应。   ![star > 100][Bronze]
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:361** 用于生成自包含 mock 对象的工具   ![star > 100][Bronze]
    * [minimock](https://github.com/gojuno/minimock) **star:264** Go接口的模拟生成器。   ![star > 100][Bronze]
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:164** 基于单事务的数据库驱动，主要用于测试目的   ![star > 100][Bronze]
    * [govcr](https://github.com/seborama/govcr) **star:82** HTTP mock : 离线测试时记录和重放浏览器的动作
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Go http.ResponseWriter的模拟对象。   ![最近一年没有更新][Yellow]

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:2892** 随机测试系统。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [gofuzz](https://github.com/google/gofuzz) **star:533** 用于生成随机值来初始化 Go 语言对象的库   ![star > 100][Bronze]
    * [Tavor](https://github.com/zimmski/tavor) **star:212** 通用模糊测试框架   ![star > 100][Bronze]

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:3610** 用于驱动和测试 Chrome, Safari, Edge, Android Webviews, 以及其他支持 Chrome 调试协议的产品   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [selenoid](https://github.com/aerokube/selenoid) **star:1238** Selenium hub 服务器的替代品，在容器中启动浏览器   ![star > 1000][Silver]
    * [cdp](https://github.com/mafredri/cdp) **star:354** 用于Chrome调试协议的类型安全绑定，可与实现该协议的浏览器或其他调试目标一起使用。   ![star > 100][Bronze]
    * [ggr](https://github.com/aerokube/ggr) **star:210** 一个轻量级服务器，可以将 Selenium Wedriver 的请求路由或代理到多个 Selenium hubs   ![star > 100][Bronze]

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:391** 为Golang实现[failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail)。   ![star > 100][Bronze]

## 文本处理

*用于解析和操作文本的库。 (翻译出错了? 试试 [英文版](README_EN.md#text-processing) 吧~)*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:8474** 快速和优雅的 Scraping 框架。   ![star > 5000][Gold]   ![最近一周有更新][Green]
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:7617** GoQuery 为 Go 语言带来了一组类似 jQuery 的语法和功能   ![star > 5000][Gold]
    * [blackfriday](https://github.com/russross/blackfriday) **star:3907** Markdown 解析器   ![star > 1000][Silver]
    * [toml](https://github.com/BurntSushi/toml) **star:2803** TOML配置格式(带反射的编码器/解码器)。   ![star > 1000][Silver]
    * [sh](https://github.com/mvdan/sh) **star:2018** Shell解析器和格式化工具。   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [go-humanize](https://github.com/dustin/go-humanize) **star:1905** 格式化程序，用于将时间、数字和内存大小转换为可读格式。   ![star > 1000][Silver]
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1255** HTML 清理工具   ![star > 1000][Silver]   ![最近一周有更新][Green]
    * [inject](https://github.com/facebookgo/inject) **star:1142** 包注入提供了一个基于反射的注入器。   ![star > 1000][Silver]   ![归档项目][Archived]
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1104** 在Go中解析RSS和Atom feeds。   ![star > 1000][Silver]
    * [go-toml](https://github.com/pelletier/go-toml) **star:609** 使用带有查询支持和方便的cli工具的TOML格式库。   ![star > 100][Bronze]
    * [commonregex](https://github.com/mingrammer/commonregex) **star:553** 一组用于Go的公共正则表达式。   ![star > 100][Bronze]
    * [slug](https://github.com/gosimple/slug) **star:380** URL 友好的 slug 化工具，支持多种语言   ![star > 100][Bronze]
    * [mxj](https://github.com/clbanning/mxj) **star:332** 将XML编码/解码为JSON或map[string]接口{};使用点符号路径和通配符提取值。替换x2j和j2x包。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  格式化二进制为字符串。
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:302** 解析Graphviz DOT语言。   ![star > 100][Bronze]
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:292** Web抓取框架将网站转换为结构化数据。   ![star > 100][Bronze]
    * [gotext](https://github.com/leonelquinteros/gotext) **star:231** GNU gettext 工具   ![star > 100][Bronze]
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:210** 函数获取字符或字符串的固定宽度。   ![star > 100][Bronze]
    * [goq](https://github.com/andrewstuart/goq) **star:147**  声明式 HTML 编组，使用结构标签和 jQuery 语法 (使用 GoQuery).   ![star > 100][Bronze]
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:131** 用于HTML的XPath查询包，允许您通过XPath表达式从HTML文档中提取数据或求值。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:100** 用于Go语言的NMEA解析器库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
    * [sdp](https://github.com/gortc/sdp) **star:70** SDP:会话描述协议[[RFC 4566](https://tools.ietf.org/html/rfc4566)]。
    * [align](https://github.com/Guitarbum722/align) **star:59** 对文本进行对齐的通用应用程序。   ![最近一年没有更新][Yellow]
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:54** 生成漂亮的固定链接地址（slug），支持多种语言   ![最近一年没有更新][Yellow]
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub 风格的 Markdown 渲染器 (使用 blackfriday) ，支持代码块高亮以及可点击的锚点
    * [genex](https://github.com/alixaxel/genex) **star:53** 将正则表达式计数并展开为所有匹配的字符串。
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:44** 通过一个 unicode 文本来猜测该文本使用的语言   ![最近一年没有更新][Yellow]
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:41** 用于Go的零宽度字符检测和删除。
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:37** Editorconfig文件解析器和Go操作器。   ![最近一周有更新][Green]
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:36** 根据正则表达式生成随机字符串   ![最近一年没有更新][Yellow]
    * [allot](https://github.com/sbstjn/allot) **star:34** 用于CLI工具和机器人的占位符和通配符文本解析。
    * [gonameparts](https://github.com/polera/gonameparts) **star:29** 将人名解析为单独的名称部分。   ![最近一周有更新][Green]
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:26** 固定宽度的文本格式(带反射的编码器/解码器)。
    * [Slugify](https://github.com/avelino/slugify) **star:26** 字符串 slug 化的工具。   ![最近一年没有更新][Yellow]
    * [go-vcard](https://github.com/emersion/go-vcard) **star:25** 解析和格式化vCard。
    * [did](https://github.com/ockam-network/did) **star:24** DID(分散标识符)解析器和Stringer。
    * [enca](https://github.com/endeveit/enca) **star:8** [libenca](http://cihar.com/software/enca/)的最小cgo绑定。   ![最近一年没有更新][Yellow]
    * [codetree](https://github.com/aerogo/codetree) **star:7** 解析缩进代码(python、pixy、scarlet等)并返回树结构。
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** 将bbCode转换为HTML，使您可以添加对自定义bbCode标记的支持。   ![最近一年没有更新][Yellow]
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:5** Atom 1.0和RSS 2.0的联合提要。   ![最近一年没有更新][Yellow]
    * [doi](https://github.com/hscells/doi) **star:4** 文档对象标识符(doi)解析器。   ![最近一年没有更新][Yellow]
    * [encdec](https://github.com/mickep76/encdec) **star:3** 软件包为编码器和解码器提供了通用接口。
    * [ltsv](https://github.com/Wing924/ltsv) **star:2** 用于Go的高性能[LTSV(标签为Tab Separeted Value)](http://ltsv.org/)阅读器。
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:461** 从文本中提取url。   ![star > 100][Bronze]
    * [gotabulate](https://github.com/bndr/gotabulate) **star:199** 使用 Go 语言简单、美观的打印表格数据   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
    * [radix](https://github.com/yourbasic/radix) **star:144** 快速字符串排序算法。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
    * [parth](https://github.com/codemodus/parth) **star:31** URL路径分割解析。
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:17** 将xml或json转换为struct。
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:15** 一个基于 sanitization 的 Go 敏感词过滤器。
    * [kace](https://github.com/codemodus/kace) **star:12** 通用大小写转换工具
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:6** 字符串参数解析器，能够理解引用及反斜杠。   ![最近一年没有更新][Yellow]
    * [TySug](https://github.com/Dynom/TySug) **star:3** 关于键盘布局的其他建议。
    * [Tagify](https://github.com/zoomio/tagify) **star:1** 从给定源生成一组标记。

## 第三方api

*用于访问第三方api的库。 (翻译出错了? 试试 [英文版](README_EN.md#third-party-apis) 吧~)*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5045** AWS 提供的官方go语言 SDK   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [github](https://github.com/google/go-github) **star:4771** 访问GitHub REST API v3的库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [slack](https://github.com/nlopes/slack) **star:2428** Slack API。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [google](https://github.com/google/google-api-go-client) **star:1930** 为Go自动生成谷歌api。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:1796** 谷歌云api Go 客户端库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:991**  Twitter 1.1 API 的 go 语言客户端   ![star > 100][Bronze]
* [discordgo](https://github.com/bwmarrin/discordgo) **star:971**  Discord Chat API的客户端。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [stripe](https://github.com/stripe/stripe-go) **star:945**  Stripe API 的 Go 语言客户端   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [facebook](https://github.com/huandu/facebook) **star:774** 支持 Facebook Graph API 的库   ![star > 100][Bronze]
* [minio-go](https://github.com/minio/minio-go) **star:723** 用于Amazon S3兼容云存储的Minio Go库。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-twitter](https://github.com/dghubble/go-twitter) **star:715**  Twitter v1.1 api 的客户端.   ![star > 100][Bronze]
* [go-jira](https://github.com/andygrunwald/go-jira) **star:577**  Go [Atlassian JIRA](https://www.atlassian.com/software/jira)的客户端库   ![star > 100][Bronze]
* [githubql](https://github.com/shurcooL/githubql) **star:504** 访问GitHub GraphQL API v4的库。   ![star > 100][Bronze]
* [webhooks](https://github.com/go-playground/webhooks) **star:357** GitHub 和 Bitbucket 的Webhook接收器。   ![star > 100][Bronze]
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:307** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   ![star > 100][Bronze]
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:300** PayPal支付API的包装器。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-marathon](https://github.com/gambol99/go-marathon) **star:189**  用于和 Mesosphere's Marathon PAAS 交互的 Go 语言库   ![star > 100][Bronze]
* [ethrpc](https://github.com/onrik/ethrpc) **star:169**  Ethereum JSON RPC API的客户端。   ![star > 100][Bronze]
* [gostorm](https://github.com/jsgilmore/gostorm) **star:118** GoStorm是一个Go库，它实现了编写Storm spout和bolt所需的通信协议，这些协议用于与Storm shell通信。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Medium](https://github.com/Medium/medium-sdk-go) **star:116** Medium的OAuth2 API 客户端。   ![star > 100][Bronze]
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:113** 一个用于通过XMPP与HipChat通信的golang包。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [hipchat](https://github.com/andybons/hipchat) **star:110** 这个项目为Hipchat API实现了一个golang客户端库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Trello](https://github.com/adlio/trello) **star:103**  Trello API的 Go 语言封装。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-trending](https://github.com/andygrunwald/go-trending) **star:101** 在Github上访问[trends repository](https://github.com/trends)和[developers](https://github.com/trending/developers)的库。   ![star > 100][Bronze]
* [cachet](https://github.com/andygrunwald/cachet) **star:66** 使用客户端库[Cachet(开源状态页系统)](https://cachethq.io/)。   ![最近一年没有更新][Yellow]
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Clarifai API的客户端。   ![最近一年没有更新][Yellow]
* [pushover](https://github.com/gregdel/pushover) **star:57**  Go 包装的 Pushover API。
* [megos](https://github.com/andygrunwald/megos) **star:56** 用于访问[Apache Mesos](http://mesos.apache.org/)集群的客户端库。   ![最近一年没有更新][Yellow]
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:52** [Internet Game Database API](https://api.igdb.com/) 客户端。
* [wit-go](https://github.com/wit-ai/wit-go) **star:47** wit.ai HTTP API 客户端。
* [gads](https://github.com/emiddleton/gads) **star:43**  Google Adwords 非官方 API
* [circleci](https://github.com/jszwedko/go-circleci) **star:41** CircleCI的API的客户端
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:39**  xkcd API 的客户端。   ![最近一年没有更新][Yellow]
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:39** [Amazon Product Advertising API](https://program.amazon.com/gp/advertising/api/detail/main.html)的客户端库。   ![最近一年没有更新][Yellow]
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:36**  Go MusicBrainz WS2客户端库。
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:35** 运行时Robot v2 API 的 Go 语言封装和命令行客户端。   ![最近一周有更新][Green]
* [fcm](https://github.com/maddevsio/fcm) **star:32**  Firebase Cloud Messaging 的 Go 语言库
* [golyrics](https://github.com/mamal72/golyrics) **star:31** Golyrics是一个从Wikia网站获取音乐歌词数据的Go库。   ![最近一年没有更新][Yellow]
* [mixpanel](https://github.com/dukex/mixpanel) **star:30** Mixpanel是一个库，用于跟踪事件并将Mixpanel概要文件更新从go应用程序发送到Mixpanel。
* [gcm](https://github.com/Aorioli/gcm) **star:29**  Google Cloud Messaging 库   ![最近一年没有更新][Yellow]
* [translate](https://github.com/poorny/translate) **star:26**  Go 在线翻译包。   ![最近一年没有更新][Yellow]
* [gami](https://github.com/bit4bit/gami) **star:26**  Asterisk Manager Interface 的 Go 语言库   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:20** 使用[Unsplash.com](https://unsplash.com) API的客户端库。
* [shopify](https://github.com/rapito/go-shopify) **star:19** 一个用于通过 Shopify API 进行增删改查的 Go 语言库。   ![最近一年没有更新][Yellow]
* [patreon-go](https://github.com/mxpv/patreon-go) **star:18**  Go Patreon API库。
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:17**  Twitch v3 API 的客户端。   ![最近一年没有更新][Yellow]
* [spotify](https://github.com/rapito/go-spotify) **star:17**  用于接入 Spotify WEB API 的 Go 语言库   ![最近一年没有更新][Yellow]
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:16** 用于访问 BreweryDB API的 Go 语言库   ![最近一年没有更新][Yellow]
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:16** [MyAnimeList API](http://myanimelist.net/modules.php?go=api)的客户端库。   ![最近一年没有更新][Yellow]
* [codeship-go](https://github.com/codeship/codeship-go) **star:15**  Codeship API v2的客户端。
* [steam](https://github.com/sostronk/go-steam) **star:14**  用于与Steam服务器进行交互的库。   ![最近一年没有更新][Yellow]
* [textbelt](https://github.com/dietsche/textbelt) **star:14** textbelt.com txt messaging API 的go语言客户端。   ![最近一年没有更新][Yellow]
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:13** Go [imgur](https://imgur.com)的客户端库
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:13**  YNAB API 的 Go 语言封装。
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:12** 简单的包装，方便的谷歌分析报告。   ![最近一年没有更新][Yellow]
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Coinpaprika API的客户端。
* [smite](https://github.com/sergiotapia/smitego) **star:10** 对 Smite game API 的封装。   ![最近一年没有更新][Yellow]
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:9** HackerNews API的小型Go客户端。   ![最近一年没有更新][Yellow]
* [simples3](https://github.com/rhnvrm/simples3) **star:9** 使用REST和用Go编写的V4签名的AWS S3库非常简单。
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:8**  SPTrans Olho Vivo API 的客户端。   ![最近一年没有更新][Yellow]
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph 发布平台 API 客户端。
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** 用于接入 statdns.com API 的库——RRDA API。通过HTTP协议进行 DNS查询。   ![最近一年没有更新][Yellow]
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6**  Tumblr v2 API 的 Go 语言封装。   ![最近一年没有更新][Yellow]
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** [谷歌G Suite Email Audit API](https://developer.google.com/admin-sdk/email-audit/) 的客户端。   ![最近一年没有更新][Yellow]
* [zooz](https://github.com/gojuno/go-zooz) **star:5**  Zooz API 的 Go 语言客户端。
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** 无任何依赖的[Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/pdfs/documentation/utmonaws/sophos-ut-restful-api.pdf?la=en)客户端
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** 用于与[Chronos](https://mesos.github.io/chronos/)作业调度程序进行交互的Go库   ![最近一年没有更新][Yellow]
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1**  TripAdvisor API 的 Go 语言封装。   ![最近一周有更新][Green]
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** Playlyfe Rest API Go SDK。   ![最近一年没有更新][Yellow]
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:1** Go library MalShare API [malshare.com](http://www.malshare.com/)

## 公用事业公司

*可以让你的生活变得更简单的实用工具.。 (翻译出错了? 试试 [英文版](README_EN.md#utilities) 吧~)*

* [fzf](https://github.com/junegunn/fzf) **star:23280** 用Go编写的命令行模糊查找器。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [hub](https://github.com/github/hub) **star:17079** 封装了 git 命令，提供了额外的功能用于在终端中和 Github 进行交互。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [delve](https://github.com/derekparker/delve) **star:12078** Go 语言调试器   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [ctop](https://github.com/bcicen/ctop) **star:8839** [Top-like](http://ctop.sh)接口(例如htop)， 用于容器数据收集。   ![star > 5000][Gold]
* [wuzz](https://github.com/asciimoo/wuzz) **star:8273** 用于HTTP检查的交互式cli工具。   ![star > 5000][Gold]
* [sqlx](https://github.com/jmoiron/sqlx) **star:6834** 为内建的数据库/sql 软件包提供一组扩展。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [peco](https://github.com/peco/peco) **star:5479** 简单的交互过滤工具。   ![star > 5000][Gold]
* [usql](https://github.com/knq/usql) **star:4681** usql 是一个通用的命令行接口，用于操作 sql 数据库。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:4522** 尽可能快速的发布 Go 语言二进制文件。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [godropbox](https://github.com/dropbox/godropbox) **star:3745** 用于编写 Go 语言服务／应用的库，来自 Dropbox.。   ![star > 1000][Silver]
* [realize](https://github.com/tockins/realize) **star:3170** Go 语言构建系统，可以监控文件变化并重新加载。运行，构建，监控文件并支持自定义路径。   ![star > 1000][Silver]
* [goreporter](https://github.com/wgliang/goreporter) **star:2485** 进行代码静态分析，单元测试，代码检视并生成代码质量报告的工具   ![star > 1000][Silver]
* [panicparse](https://github.com/maruel/panicparse) **star:2145** 将类似的协程分组并对调用栈进行着色   ![star > 1000][Silver]
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2027** 实现 Hystrix 风格的、程序员预定义的 fallback 机制（熔断。   ![star > 1000][Silver]
* [resty](https://github.com/go-resty/resty) **star:1998** 简单的 HTTP 和 REST 客户端，受到 Ruby rest-client 的启发。   ![star > 1000][Silver]
* [Task](https://github.com/go-task/task) **star:1946** 简单的“Go”的选择。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [minify](https://github.com/tdewolff/minify) **star:1864** 用于HTML、CSS、JS、XML、JSON和SVG文件格式的快速缩小器。   ![star > 1000][Silver]
* [mmake](https://github.com/tj/mmake) **star:1446** 现代 Make 工具   ![star > 1000][Silver]
* [Storm](https://github.com/asdine/storm) **star:1359** 一个简单又强大的用于 BoltDB 的工具   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [mole](https://github.com/davrodpin/mole) **star:1300** cli应用程序可以轻松创建ssh隧道。   ![star > 1000][Silver]
* [go-funk](https://github.com/thoas/go-funk) **star:1198** 现代 Go 语言工具库，提供了很多有用的工具 (map, find, contains, filter, chunk, reverse, ...)   ![star > 1000][Silver]
* [mc](https://github.com/minio/mc) **star:1122** Minio Client 提供了一组工具，用于操作 Amazon S3 兼容云存储和文件系统。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [filetype](https://github.com/h2non/filetype) **star:955** 通过数字签名来推测文件类型。   ![star > 100][Bronze]
* [boilr](https://github.com/tmrts/boilr) **star:943** 非常快的CLI工具，用于从样板模板创建项目。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [mergo](https://github.com/imdario/mergo) **star:853** 用于将结构体和map合并进 Go 语言的工具。对于配置默认值，避免杂乱的if语句很有帮助。   ![star > 100][Bronze]
* [spinner](https://github.com/briandowns/spinner) **star:800**  一个 Go 语言软件包，提供多种选项，方便在终端中创建加载动画。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:796** 接通断路器。   ![star > 100][Bronze]
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:720** git-time-metric - 。   ![star > 100][Bronze]
* [jump](https://github.com/gsamokovarov/jump) **star:661** 通过学习你的习惯，可以帮助你更快地导航。   ![star > 100][Bronze]
* [immortal](https://github.com/immortal/immortal) **star:606** \*nix 跨平台 (与操作系统无关的)监控程序。   ![star > 100][Bronze]
* [htcat](https://github.com/htcat/htcat) **star:482** 并行及流水线的 HTTP GET 工具。   ![star > 100][Bronze]
* [go-dry](https://github.com/ungerik/go-dry) **star:433** DRY(don't repeat yourself)库。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gopencils](https://github.com/bndr/gopencils) **star:423** 小而简单的包，可以轻松地使用REST api。   ![star > 100][Bronze]
* [godaemon](https://github.com/VividCortex/godaemon) **star:405** 用于编写守护进程的工具   ![star > 100][Bronze]
* [request](https://github.com/mozillazg/request) **star:355** Go 语言版的 HTTP Requests for Humans™.。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [circuit](https://github.com/cep21/circuit) **star:330** 一个高效和功能齐全的 类似 Hystrix Go 实现断路器模式。   ![star > 100][Bronze]
* [ergo](https://github.com/cristianoliveira/ergo) **star:314** 管理运行在不同端口上的多个本地服务变得很容易。   ![star > 100][Bronze]
* [koazee](https://github.com/wesovilabs/koazee) **star:296** 库的灵感来自于延迟计算和函数式编程，从而减少了使用数组的麻烦。   ![star > 100][Bronze]
* [go-rate](https://github.com/beefsack/go-rate) **star:292**  Go 限速器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gohper](https://github.com/cosiner/gohper) **star:247** 多种能够帮助你进行软件开发的工具和模块。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [clockwork](https://github.com/jonboulle/clockwork) **star:221** 一个简单的假 clock 。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Deepcopier](https://github.com/ulule/deepcopier) **star:209** 结构体拷贝   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [serve](https://github.com/syntaqx/serve) **star:191** 任何您需要的静态http服务器。   ![star > 100][Bronze]
* [go-trigger](https://github.com/sadlil/go-trigger) **star:182** Go 语言全局事件触发器，通过 id 和触发器，在程序的任何地方注册事件。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [retry](https://github.com/kamilsk/retry) **star:161** 基于上下文的功能机制，反复执行命令直到成功。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:160** Go:generate 工具，用于构建 Go 语言插件(1.8 only)，并对导出的符号进行包装。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [rerun](https://github.com/ivpusic/rerun) **star:153** 当源代码发生更改时，重新编译和重新运行go应用程序。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [moldova](https://github.com/StabbyCutyou/moldova) **star:148** 基于输入目标生成随机数据的工具   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gotenv](https://github.com/subosito/gotenv) **star:143** 从 `.env` 或者任何 `io.Reader`。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [gubrak](https://github.com/novalagung/gubrak) **star:142** 带有语法糖的Golang实用工具，就像lodash。   ![star > 100][Bronze]
* [robustly](https://github.com/VividCortex/robustly) **star:135** 有弹性的执行函数，遇到错误时捕获并重新运行。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [util](https://github.com/shomali11/util) **star:134** 有用实用函数的集合。(字符串，并发，操作，…)   ![star > 100][Bronze]
* [Death](https://github.com/vrecan/death) **star:133** 利用信号管理应用程序的关闭。   ![star > 100][Bronze]
* [apm](https://github.com/topfreegames/apm) **star:130** Go 语言进程管理工具具有HTTP API.。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:124** 用于基于神奇数字的MIME类型检测的包。   ![star > 100][Bronze]
* [chyle](https://github.com/antham/chyle) **star:109** 使用具有多种配置可能性的git存储库生成变更日志。   ![star > 100][Bronze]
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:105** 用Go编写的XML站点地图生成器。   ![star > 100][Bronze]
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:100** LiveReload 服务器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [onecache](https://github.com/adelowo/onecache) **star:98** 支持多个后端存储(Redis、Memcached、文件系统等)的缓存库。
* [toolbox](https://github.com/viant/toolbox) **star:96** 切片, map, multimap, 结构体, 函数,数据转换工具。服务路由，宏求值和标记器。
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:83** 纯Go bsdiff和bspatch库和CLI工具。
* [pm](https://github.com/VividCortex/pm) **star:71** 进程(即goroutine)管理器与HTTP API。
* [UNIS](https://github.com/esemplastic/unis) **star:69** Go 语言字符串处理函数的通用架构 。   ![最近一年没有更新][Yellow]
* [xferspdy](https://github.com/monmohan/xferspdy) **star:68** Xferspdy在golang中提供二进制diff和补丁库。   ![最近一年没有更新][Yellow]
* [go-health](https://github.com/Talento90/go-health) **star:63** 健康包简化了向服务中添加健康检查的方式。   ![最近一年没有更新][Yellow]
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:58** 数据库客户端，用于主-从 (或主-主) 数据库，集成了简单的、轻量级的轮询调度负载均衡。
* [multitick](https://github.com/VividCortex/multitick) **star:58** 用于 aligned tickers 的多路复用   ![最近一年没有更新][Yellow]
* [repeat](https://github.com/ssgreg/repeat) **star:56** 执行不同的后 backoff 策略，这对重新尝试操作和心跳非常有用。
* [abutil](https://github.com/bahlo/abutil) **star:51** 常用 Go 语言工具的集合。   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [minquery](https://github.com/icza/minquery) **star:50** MongoDB / mgo.v2, 支持高效分页查询(用于继续列出我们停止的文档的游标)。
* [handy](https://github.com/miguelpragier/handy) **star:45** 许多实用程序和帮助程序，如字符串处理程序/格式化程序和验证器。
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:45** 解析你 Go 语言代码中的 TODOs（待办事项）。
* [golog](https://github.com/mlimaloureiro/golog) **star:42** 简单、轻量级的命令后工具，用于对你的计划任务进行跟踪。
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:42** 纯粹 Go 超性能MIME嗅探库/实用程序。
* [goback](https://github.com/carlescere/goback) **star:40**  一个 Go 语言的简单的指数补偿包。   ![最近一年没有更新][Yellow]
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:39** 不需要编写任何汇编代码就能使用 x86 SIMD。   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [gaper](https://github.com/maxcnunes/gaper) **star:39** 当Go项目崩溃或一些人看到文件更改时，构建并重新启动该项目。
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:38** 通用多工作站剪切板，使用类似 S3 的后端作为存储。   ![最近一周有更新][Green]
* [golarm](https://github.com/msempere/golarm) **star:34** 告警（支持系统事件）。   ![最近一年没有更新][Yellow]
* [retry](https://github.com/thedevsaddam/retry) **star:34** 简单易用的重试机制包，为 Go 。   ![最近一年没有更新][Yellow]
* [myhttp](https://github.com/inancgumus/myhttp) **star:34** 简单的API，使HTTP GET请求与超时支持。   ![最近一年没有更新][Yellow]
* [goreadability](https://github.com/philipjkim/goreadability) **star:32** 网页摘要提取器使用Facebook开放图形和arc90的可读性。
* [gpath](https://github.com/tenntenn/gpath) **star:31**  用于简化结构体域访问的库。   ![最近一年没有更新][Yellow]
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:29** 对 Go 来说，重试变得简单而容易。
* [rclient](https://github.com/zpatrick/rclient) **star:27** 可读、灵活、易于使用的REST api客户端。
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:27** conseilSeaweedFS 客户端，几乎具有全部的特性。
* [pgo](https://github.com/arthurkushman/pgo) **star:24** 用于PHP社区的 Convenient 函数。
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:22** 一个小巧的 Go 语言库用于生成占位图片。   ![最近一年没有更新][Yellow]
* [ugo](https://github.com/alxrm/ugo) **star:20** uGo 是一个切片工具箱，有着和 Go 语言一致的语法法。   ![最近一年没有更新][Yellow]
* [generate](https://github.com/go-playground/generate) **star:19** 针对一个路径或环境变量，递归的执行 Go generate，可以通过正则表达式来进行过滤。   ![最近一年没有更新][Yellow]
* [gostrutils](https://github.com/ik5/gostrutils) **star:16** 字符串操作和转换函数的集合。
* [evaluator](https://github.com/nullne/evaluator) **star:16** 基于 s-expression。它很简单，很容易扩展。   ![最近一年没有更新][Yellow]
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** 编译时控制的日志，让你的 release 包变得更小而不需移除 debug 调用。   ![最近一年没有更新][Yellow]
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:14**  用于将结构体编码进 http 头的 Go 语言库
* [okrun](https://github.com/xta/okrun) **star:14**  Go 运行错误 steamroller。   ![最近一年没有更新][Yellow]
* [filler](https://github.com/yaronsumel/filler) **star:14** 使用“fill”标签填充结构的小工具。   ![最近一年没有更新][Yellow]
* [filter](https://github.com/gookit/filter) **star:13** 提供Go数据的过滤、清理和转换。
* [dbt](https://github.com/nikogura/dbt) **star:13** 用于从中心可信存储库运行自更新签名二进制文件的框架。
* [structs](https://github.com/PumpkinSeed/structs) **star:12** 简单来讲就是 "Make" 的替代品。   ![最近一年没有更新][Yellow]
* [scan](https://github.com/blockloop/scan) **star:12** 扫描golang的sql。行直接指向结构、片或基本类型。
* [ghokin](https://github.com/antham/ghokin) **star:12** 没有外部依赖的gherkin (cucumber, behat…)并行格式化程序。
* [rerate](https://github.com/abo/rerate) **star:12** 基于 Redis 的速率计数器和限速器   ![最近一年没有更新][Yellow]
* [retry](https://github.com/shafreeck/retry) **star:10** 一个相当简单的库，以确保您的工作可以完成。
* [command](https://github.com/txgruppi/command) **star:9** 命令模式，支持线程安全的串行、并行调度。   ![最近一年没有更新][Yellow]
* [backscanner](https://github.com/icza/backscanner) **star:8** 类似 bufio 的扫描器，但它以相反的顺序读取和返回行，从给定的位置开始，然后返回。
* [ctxutil](https://github.com/posener/ctxutil) **star:7** 上下文工具。
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:7** 一个用于Go的MIME类型嗅探器。
* [sslice](https://github.com/yaa110/sslice) **star:5** 创建一个总是排序的切片。
* [silk](https://github.com/chrispassas/silk) **star:4** 阅读silk netflow文件。   ![最近一周有更新][Green]
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:3** 基本类型之间的片转换。
* [slicer](https://github.com/leaanthony/slicer) **star:3** 使处理切片更容易。
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:3** 打包处理问题细节。
* [retry](https://github.com/percolate/retry) **star:2** 一个简单但高度可配置的Go重试包。
* [blank](https://github.com/Henry-Sarabia/blank) **star:1** 验证或删除字符串中的空白。
* [olaf](https://github.com/btnguyen2k/olaf) **star:1** Twitter Snowflake 在Go中实现。

## UUID

*用于处理uuid的库。 (翻译出错了? 试试 [英文版](README_EN.md#uuid) 吧~)*

* [ulid](https://github.com/oklog/ulid) **star:1683** 实现了ULID(普遍唯一的词典分类标识符)。   ![star > 1000][Silver]
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  没有麻烦，安全，快速的唯一标识符与命令。
* [uuid](https://github.com/gofrs/uuid) **star:569** 通用唯一标识符(UUID)的实现。支持uuid的创建和解析。积极维护satori uuid的fork。   ![star > 100][Bronze]
* [wuid](https://github.com/edwingeng/wuid) **star:288** 一个非常快的唯一数字生成器，比UUID快10-135倍。   ![star > 100][Bronze]
* [goid](https://github.com/jakehl/goid) **star:21** 生成和解析RFC4122兼容的V4 uuid。
* [sno](https://github.com/muyo/sno) **star:16** 使用嵌入元数据的紧凑、可排序和快速的惟一id。
* [uuid](https://github.com/agext/uuid) **star:10** 使用快速或加密质量的随机节点标识符生成、编码和解码UUIDs v1。
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:4** 一个小而有效的Go唯一字符串ID生成器。

## 验证

*库进行验证。 (翻译出错了? 试试 [英文版](README_EN.md#validation) 吧~)*

* [govalidator](https://github.com/asaskevich/govalidator) **star:3560** 用于字符串，数字，切片和结构的验证器和sanitizers。   ![star > 1000][Silver]
* [validator](https://github.com/go-playground/validator) **star:3542**  Go 结构体及域验证，包括：跨域、跨结构体, Map, 切片和数组。   ![star > 1000][Silver]
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1054** 支持各种数据类型(结构、字符串、映射、片等)的验证，使用可配置和可扩展的验证规则，这些规则在通常的代码构造中指定，而不是在结构标签中指定。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:717** 用简单的规则验证Golang请求数据。深受Laravel请求验证的启发。   ![star > 100][Bronze]
* [validate](https://github.com/gookit/validate) **star:98**  Go 封装数据验证和过滤。支持验证映射、结构、请求(表单、JSON、url)。值，上载文件)数据和更多特性。   ![包含中文文档][CN]
* [checkdigit](https://github.com/osamingo/checkdigit) **star:45** 提供校验数字算法(Luhn, Verhoeff, Damm)和计算器(ISBN, EAN, JAN, UPC等)。
* [jio](https://github.com/faceair/jio) **star:21** jio是一个json模式验证器，类似于[joi](https://github.com/hapijs/joi)。   ![包含中文文档][CN]
* [validate](https://github.com/gobuffalo/validate) **star:19** 这个包提供了一个框架，用于为Go应用程序编写验证。
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:11** 一种规范和约定验证器。   ![最近一周有更新][Green]

## 版本控制

*用于版本控制的库。 (翻译出错了? 试试 [英文版](README_EN.md#version-control) 吧~)*

* [go-git](https://github.com/src-d/go-git) **star:4300** 纯Go中高度可扩展的Git实现。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [git2go](https://github.com/libgit2/git2go) **star:1366**  libgit2 的 Go 语言接口。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [hercules](https://github.com/src-d/hercules) **star:542** 从Git存储库历史中获得高级见解。   ![star > 100][Bronze]
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:71** 在Go中操作和检查VCS存储库。   ![最近一周有更新][Green]
* [gh](https://github.com/rjeczalik/gh) **star:69** 用于GitHub webhook的可编写脚本的服务器和net/http中间件。
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo是一组Go包的集合，提供对本地Mercurial存储库的读取访问。   ![最近一年没有更新][Yellow]

## 视频

*用于操作视频的库。 (翻译出错了? 试试 [英文版](README_EN.md#video) 吧~)*

* [goav](https://github.com/giorgisio/goav) **star:794** FFmpeg的Comphrensive。   ![star > 100][Bronze]
* [gmf](https://github.com/3d0c/gmf) **star:533**  FFmpeg av\* 库的 Go 语言接口。   ![star > 100][Bronze]
* [go-astits](https://github.com/asticode/go-astits) **star:260** 在GO中解析和演示MPEG传输流(.ts)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [go-astisub](https://github.com/asticode/go-astisub) **star:168** 使用 Go 语言操作字幕(.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.)。   ![star > 100][Bronze]
* [gst](https://github.com/ziutek/gst) **star:152**  GStreamer的Go工具。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:66** Go绑定libvlc 2.X/3.X/4。X(由VLC媒体播放器使用)。
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:39** 苹果m3u8播放列表的解析器和生成器库。
* [v4l](https://github.com/korandiz/v4l) **star:28** 用于Linux的视频捕捉库，用Go编写。   ![最近一年没有更新][Yellow]
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:11** 字幕格式支持 .srt、.ttml和.ass。

## Web框架

*全栈 web 框架。 (翻译出错了? 试试 [英文版](README_EN.md#web-frameworks) 吧~)*

* [Gin](https://github.com/gin-gonic/gin) **star:30176** Gin是一个用Go编写的web框架!它具有一个类似于martini的API，性能更好，速度快40倍。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Beego](https://github.com/astaxie/beego) **star:21536** beego是一种用于 Go 编程语言的开源高性能web框架。   ![star > 5000][Gold]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [Buffalo](http://gobuffalo.io)  为 Go 语言带来堪比 Rails 的高生产效率!
* [Echo](https://github.com/labstack/echo) **star:14740** 高性能、极简的Go web框架。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Revel](https://github.com/revel/revel) **star:11256** 用于Go语言的高效web框架。   ![star > 5000][Gold]
* [Goa](https://github.com/goadesign/goa) **star:3513** Goa为在Go中开发远程api和微服务提供了一种全面的方法。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3335** 设置RESTful JSON API的快速简便方法。   ![star > 1000][Silver]
* [Gizmo](https://github.com/NYTimes/gizmo) **star:2852** 《纽约时报》使用的微服务工具包。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Macaron](https://github.com/go-macaron/macaron) **star:2817** Macaron 是一个高效的模块化设计的web框架   ![star > 1000][Silver]
* [utron](https://github.com/gernest/utron) **star:2134** Go(Golang)的轻量级MVC框架。   ![star > 1000][Silver]
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:996** 用于构建 JSON web 服务的 Go 语言框架，受到 Dropwizard 的启发。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [tango](https://github.com/lunny/tango) **star:818** 微型的、支持插件的 web 框架。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [traffic](https://github.com/pilu/traffic) **star:518** Sinatra启发了regexp/pattern mux和用于Go的web框架。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gongular](https://github.com/mustafaakin/gongular) **star:415**  快速 Go web 框架，支持输入映射／验证以及依赖注入。   ![star > 100][Bronze]
* [neo](https://github.com/ivpusic/neo) **star:393** Neo是一个非常简单且快速的Web框架API。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [mango](https://github.com/paulbellamy/mango) **star:340** ManGo 是一个模块化 web 应用框架，受到 Rack 和 PEP333 的启发。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Air](https://github.com/aofei/air) **star:332** 一个理想的精细化的Go web框架。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Gondola](https://github.com/rainycape/gondola) **star:314** web框架写的网站越快越好。   ![star > 100][Bronze]
* [Golf](https://github.com/dinever/golf) **star:236** Golf 是一个快速、简单、轻量级的 Go 语言微型 web 框架。具有强大的功能且没有标准库以外的依赖。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Aero](https://github.com/aerogo/aero) **star:158** 高性能的Go web框架，在Lighthouse中达到最高分。   ![star > 100][Bronze]
* [Gem](https://github.com/go-gem/gem) **star:153** 简单快速的web框架，对REST API友好。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-rest](https://github.com/ungerik/go-rest) **star:116** 小型的 REST 框架。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [hiboot](https://github.com/hidevopsio/hiboot) **star:87** hiboot是一个高性能的web应用程序框架，支持自动配置和依赖注入。   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [WebGo](https://github.com/bnkamalesh/webgo) **star:75** 构建web应用程序的微框架;处理程序链接、中间件和上下文注入。与标准库兼容的HTTP处理程序(即http.HandlerFunc)。
* [Golax](https://github.com/fulldump/golax) **star:71** 一个非Sinatra快速HTTP框架，支持谷歌自定义方法、深度拦截器、递归等。   ![最近一年没有更新][Yellow]
* [Microservice](https://github.com/claygod/microservice) **star:57** 创建微服务的框架，用Golang编写。
* [uAdmin](https://github.com/uadmin/uadmin) **star:54** 受到 Sinatra 启发的 Go 语言 web 框架。
* [YARF](https://github.com/yarf-framework/yarf) **star:50** 快速微框架，旨在以快速和简单的方式构建REST api和web服务。
* [Fireball](https://github.com/zpatrick/fireball) **star:49** 感觉更加自然的 web 框架。
* [vox](https://github.com/aisk/vox) **star:40** 一个面向人类的golang web框架，深受Koa的启发。
* [patron](https://github.com/beatlabs/patron) **star:35** Patron是一个遵循最佳云实践的微服务框架，专注于提升开发效率。   ![最近一周有更新][Green]
* [REST Layer](http://rest-layer.io)  框架，用于在数据库之上构建REST/GraphQL API，主要是通过代码进行配置。
* [Resoursea](https://github.com/resoursea/api) **star:29** 用于快速编写基于资源的服务的REST框架。   ![最近一年没有更新][Yellow]
* [aah](https://aahframework.org)  可伸缩、高性能、快速开发的Go Web框架。
* [rex](https://github.com/goanywhere/rex) **star:27**  Rex 是一个用于进行模块化开发的库，基于Gorilla/mux 完全兼容大多数的 net/HTTP.   ![最近一年没有更新][Yellow]
* [nio](https://github.com/go-nio/nio) **star:21** 现代的、最小的和高效的Go HTTP框架。   ![最近一周有更新][Green]   ![归档项目][Archived]
* [rux](https://github.com/gookit/rux) **star:9** 用于构建golang HTTP应用程序的简单而快速的web框架。   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [Banjo](https://github.com/nsheremet/banjo) **star:8** 非常简单和快速的网络框架 Go 。   ![最近一年没有更新][Yellow]

### 中间件

#### 仿真中间件

* [Tollbooth](https://github.com/didip/tollbooth) **star:1252** 限制速率的 HTTP 请求处理程序。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [CORS](https://github.com/rs/cors) **star:1224** 轻松地向API添加CORS功能。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Limiter](https://github.com/ulule/limiter) **star:793** 简单的速度限制中间件。   ![star > 100][Bronze]
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:748** 添加/解析Server-Timing头。   ![star > 100][Bronze]
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:89** 使用Lightning Network(比特币)实现基于每个请求的api货币化中间件。
* [XFF](https://github.com/sebest/xff) **star:72** 处理 X-Forwarded-For 头的中间件。
* [formjson](https://github.com/rs/formjson) **star:33** 透明地将JSON输入作为标准表单POST处理。   ![最近一年没有更新][Yellow]
* [client-timing](https://github.com/posener/client-timing) **star:13** 用于服务器定时报头的HTTP客户机。

#### 用于创建HTTP中间件的库

* [negroni](https://github.com/urfave/negroni) **star:6344** 符合语言习惯的 HTTP 中间件库。   ![star > 5000][Gold]   ![包含中文文档][CN]
* [alice](https://github.com/justinas/alice) **star:1825** 用于连接中间件的库，简单无痛苦。   ![star > 1000][Silver]
* [render](https://github.com/unrolled/render) **star:1269** Go package用于方便地呈现JSON、XML和HTML模板响应。   ![star > 1000][Silver]
* [stats](https://github.com/thoas/stats) **star:536** 使用中间件来存储关于web应用程序的各种信息。   ![star > 100][Bronze]
* [interpose](https://github.com/carbocation/interpose) **star:289** golang的极简网络/http中间件。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [muxchain](https://github.com/stephens2424/muxchain) **star:208** 用于net/http的轻量级中间件。   ![star > 100][Bronze]
* [renderer](https://github.com/thedevsaddam/renderer) **star:169** 简单、轻量级和更快的响应(JSON、JSONP、XML、YAML、HTML、文件)。   ![star > 100][Bronze]
* [rye](https://github.com/InVisionApp/rye) **star:94** 支持JWT、CORS、Statsd和Go 1.7上下文的小型Go中间件库(带有罐装中间件)。
* [gores](https://github.com/alioygur/gores) **star:82** 处理HTML、JSON、XML等响应的Go包。对于RESTful api非常有用。
* [chain](https://github.com/codemodus/chain) **star:63** 带有范围数据的处理程序包装器链接(基于网络/上下文的“中间件”)。
* [go-wrap](https://github.com/go-on/wrap) **star:59** net/http的小型中间件包。   ![最近一年没有更新][Yellow]
* [catena](https://github.com/codemodus/catena) **star:7** HTTP.Handler wrapper catenation (和chain具有相同的 API ).。

### 路由器

* [httprouter](https://github.com/julienschmidt/httprouter) **star:9749** 高性能路由。使用这个库和标准http处理工具可以构建一个非常高性能大web框架。   ![star > 5000][Gold]
* [mux](https://github.com/gorilla/mux) **star:9734** 强大的URL路由器和调度器为golang。   ![star > 5000][Gold]
* [chi](https://github.com/go-chi/chi) **star:6072** 小巧、快速、具有丰富表达力的 HTTP 路由，基于net/context.。   ![star > 5000][Gold]
* [gocraft/web](https://github.com/gocraft/web) **star:1398** Mux和中间件包在Go中。   ![star > 1000][Silver]
* [Bone](https://github.com/go-zoo/bone) **star:1220** 闪电快速HTTP多路复用器。   ![star > 1000][Silver]
* [Goji](https://github.com/goji/goji) **star:768** 枸杞是一种简约的和灵活的与支持'net/context` HTTP请求多路复用器。   ![star > 100][Bronze]
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:751** 高性能路由器分叉从`httprouter`。第一个路由器适合`fasthttp`。   ![star > 100][Bronze]
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:452** 一个简单和快速的HTTP路由器 Go 。   ![star > 100][Bronze]
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:385** 高速，灵活，基于树的 HTTP 路由。受到了 httprouter 的启发。   ![star > 100][Bronze]
* [lars](https://github.com/go-playground/lars) **star:374** 是一个轻量级、快速、可扩展、零分配的HTTP路由，用于创建定制化的框架。   ![star > 100][Bronze]
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:359** 一个非常快的Go (golang) HTTP路由器，支持正则表达式路由匹配。完全支持构建RESTful api。   ![star > 100][Bronze]
* [Siesta](https://github.com/VividCortex/siesta) **star:349** 编写中间件和处理程序的可组合框架。   ![star > 100][Bronze]
* [vestigo](https://github.com/husobee/vestigo) **star:250** 高性能，独立，HTTP兼容的URL路由器的go web应用程序。   ![star > 100][Bronze]
* [gowww/router](https://github.com/gowww/router) **star:157** 超快的HTTP 路由，完全兼容 net/HTTP.Handler 接口.。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [alien](https://github.com/gernest/alien) **star:106** 轻量级和快速http路由器从外层空间。   ![star > 100][Bronze]
* [violetear](https://github.com/nbari/violetear) **star:95** HTTP路由器。
* [Bxog](https://github.com/claygod/Bxog) **star:94** 简单和快速的HTTP路由器 Go 。它可以处理不同难度、长度和嵌套的路径。他还知道如何根据接收到的参数创建URL。
* [xmux](https://github.com/rs/xmux) **star:88** 高性能mux基于httprouter 'net/context`支持。   ![最近一年没有更新][Yellow]
* [pure](https://github.com/go-playground/pure) **star:86** 是一个轻量级HTTP路由器，它坚持net/ HTTP“实现”的std。
* [GoRouter](https://github.com/vardius/gorouter) **star:49** GoRouter 是一个服务器/API 微型框架、HTTP 请求路由 router, 数据分选器，提供了支持net/context的中间件。
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:38** 一个简单的Go HTTP路由器。
* [FastRouter](https://github.com/razonyang/fastrouter) **star:18** 一个快速，灵活的HTTP路由器写在Go。   ![最近一年没有更新][Yellow]

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:551** 为 Go 语言实现的 Win32 OLE。   ![star > 100][Bronze]
* [d3d9](https://github.com/gonutz/d3d9) **star:86**  Direct3D9 的 Go 语言接口。
* [gosddl](https://github.com/MonaxGT/gosddl) **star:1** 从SDDL-string到用户友好的JSON的转换器。SDDL由四个部分组成:所有者、主群、DACL、SACL。

## XML

*用于操作XML的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#xml) 吧~)*

* [zek](https://github.com/miku/zek) **star:256** 从XML生成Go结构。   ![star > 100][Bronze]
* [xpath](https://github.com/antchfx/xpath) **star:170** Go的XPath包。   ![star > 100][Bronze]
* [xquery](https://github.com/antchfx/xquery) **star:145** XQuery允许您使用XPath表达式从HTML/XML文档中提取数据。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** 简单的命令行XML比较器，生成文件夹、文件和标记的差异。   ![最近一年没有更新][Yellow]
* [xml2map](https://github.com/sbabiv/xml2map) **star:15** XML来映射转换器编写的Golang。
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:6** 基于libxml2的xmlwriter模块的过程性XML生成API。

# 工具

* Go 软件和插件。 (翻译出错了? 试试 [英文版](README_EN.md#tools) 吧~)*

## 代码分析

* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple 是 Go 源代码的linter，专门用于简化代码。
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  添加 zero 返回声明，以匹配 func 返回类型。
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [GoLint](https://github.com/golang/lint) **star:3163** Go 源码的 linter。   ![star > 1000][Silver]
* [errcheck](https://github.com/kisielk/errcheck) **star:1322** Errcheck是一个用于检查Go程序中未检查错误的程序。   ![star > 1000][Silver]
* [gcvis](https://github.com/davecheney/gcvis) **star:920** 实时可视化跟踪 GC 数据。   ![star > 100][Bronze]
* [php-parser](https://github.com/z7zmey/php-parser) **star:643** 用 Go 编写的 PHP 解析器。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  用于大量静态分析检查，您可能已经从 c# 的 ReSharper 等工具中习惯了这些检查。
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  在源码中寻找没有直接单元测试的函数和方法。
* [go-critic](https://github.com/go-critic/go-critic) **star:579** 源代码检查工具。   ![star > 100][Bronze]
* [GolangCI](https://golangci.com/)  GolangCI 是一个针对 GitHub pull 请求的自动代码审查服务。服务是开源的，对开源项目是免费的。
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  来修复(添加，删除) Go 中自动导入的工具。
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:377** 基于 Web 的 Golang AST 可视化工具。   ![star > 100][Bronze]
* [GoCover.io](http://gocover.io/)  GoCover.io 提供了任意 golang 包的代码覆盖率服务。
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:282** go-cleanarch 的创建是为了验证 Clean 体系结构规则，比如 Go 项目中的依赖关系。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [unconvert](https://github.com/mdempsky/unconvert) **star:258** 在源码中删除不必要的类型转换。   ![star > 100][Bronze]
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  对未使用的常量、变量、函数和类型的代码进行检查。
* [gostatus](https://github.com/shurcooL/gostatus) **star:240** 用于显示包含 Go 包的存储库的状态的命令行工具，。   ![star > 100][Bronze]
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:187** 找出项目中过期的依赖项。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [dupl](https://github.com/mibk/dupl) **star:176** 用于代码克隆检测的工具。   ![star > 100][Bronze]
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:166** 检查 Go 项目最近的向下不兼容修改。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:95** checkstyle是一个类似于java checkstyle的检查工具。
* [lint](https://github.com/surullabs/lint) **star:63** 将 linters 作为测试的一部分。
* [validate](https://github.com/mccoyst/validate) **star:62** 使用 tags 自动验证结构体字段。   ![最近一年没有更新][Yellow]
* [go-outdated](https://github.com/firstrow/go-outdated) **star:44** 显示过期包的终端应用。   ![归档项目][Archived]

## 编辑器插件

* [vim-go](https://github.com/fatih/vim-go) **star:10875** Go 开发会用到的 Vim 插件。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [vscode-go](https://github.com/Microsoft/vscode-go) **star:5140** Visual Studio代码的扩展(VS代码)，它提供了对Go语言的支持。   ![star > 5000][Gold]
* [gocode](https://github.com/nsf/gocode) **star:4746** Autocompletion daemon for the Go programming language.   ![star > 1000][Silver]
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  在 VS Code 中支持 Go 的基准分析。
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3240** 包含了可为文本编辑器 SublimeText 3 提供代码自动填充和其他类似IDE的功能的 Golang IDE 插件集合。   ![star > 1000][Silver]
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1485** 在 Atom 中添加自动完成，格式化，语法检查，高亮和审查。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-mode](https://github.com/dominikh/go-mode.el) **star:959** 在 GNU/Emacs 支持 GO。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Watch](https://github.com/eaburns/Watch) **star:169** Runs a command in an acme win on file changes.   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:81** 在保存时突出显示语法错误的 Vim 插件。   ![最近一年没有更新][Yellow]
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  JetBrains IDEs 的 Go 插件。
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:30** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:17** 基于函数或方法的签名生成Go测试的Vim插件。
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:12** 在 Theia IDE中支持 Go。

## Go 生成工具

* [gotests](https://github.com/cweill/gotests) **star:2211** 从源代码生成测试用例。   ![star > 1000][Silver]
* [genny](https://github.com/cheekybits/genny) **star:975** 优雅的 Go 泛型。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [re2dfa](https://github.com/opennota/re2dfa) **star:168** 将正则表达式转换为有限状态机，并输出 Go 源代码。   ![star > 100][Bronze]
* [TOML-to-Go](https://xuri.me/toml-to-go)  在浏览器中将 TOML 转换为 Go 类型。
* [gocontracts](https://github.com/Parquery/gocontracts) **star:51** 通过同步代码和文档来实现 design-by-contract 设计。
* [gonerics](http://github.com/bouk/gonerics)  Go中的易用的泛型。
* [gounit](https://github.com/hexdigest/gounit) **star:29** 使用您自己的模板生成Go测试用例。   ![最近一年没有更新][Yellow]
* [generic](https://github.com/usk81/generic) **star:28** 灵活的 Go 数据类型。
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:12** 可生成用于切片的 Haskell。

## Go 工具

* [go-swagger](https://github.com/go-swagger/go-swagger) **star:4007** 基于 Go 的Swagger 2.0实现。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:1998** 使用 dot format 可视化 Go 程序的调用图。   ![star > 1000][Silver]
* [richgo](https://github.com/kyoh86/richgo) **star:390** 用文本装饰丰富 go test 的输出。   ![star > 100][Bronze]
* [depth](https://github.com/KyleBanks/depth) **star:381** 通过分析导入，将包依赖关系树可视化输出。   ![star > 100][Bronze]
* [gb](https://getgb.io/)  一个基于项目的易用的构建工具。
* [rts](https://github.com/galeone/rts) **star:184** 从服务器响应生成Go结构。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [godbg](https://github.com/tylerwince/godbg) **star:156** 实现了 Rusts 的 dbg! 宏，可以方便的在开发过程中快速、容易地调试。   ![star > 100][Bronze]
* [OctoLinker](https://github.com/OctoLinker/browser-extension)  借助的 OctoLinker 浏览器扩展，可以高效的地浏览  GitHub go文件。
* [colorgo](https://github.com/songgao/colorgo) **star:98** 将 go 命令包装成彩色的 go build 输出。   ![最近一年没有更新][Yellow]
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:37** Bash completion for go and wgo。   ![最近一年没有更新][Yellow]
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:13** 一个[Yeoman](http://yeoman.io)生成器，用于启动新的 Go 项目。
* [gilbert](https://go-gilbert.github.io)  一个为 Go 项目而生的构建系统和任务运行器。

## 软件包

*用Go编写的软件。 (翻译出错了? 试试 [英文版](README_EN.md#software-packages) 吧~)*

### DevOps 工具

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:56593** 来自谷歌的容器集群管理器。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Moby](https://github.com/moby/moby) **star:54518** Collaborative project for the container ecosystem to assemble container-based systems.   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [traefik](https://github.com/containous/traefik) **star:23864** 反向代理和负载均衡器，支持多个后端。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [Gitea](https://github.com/go-gitea/gitea) **star:15405** 从 Gogs fork，完全由社区驱动。   ![star > 5000][Gold]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  将所有GitHub repositories、issues、milestones 和 labels 都迁移到 Gitea。
* [Vegeta](https://github.com/tsenart/vegeta) **star:12239** HTTP负载测试工具和库。超过9000 !   ![star > 5000][Gold]
* [Hey](https://github.com/rakyll/hey) **star:6371** 压力测试工具，可用来代替 ApacheBench (ab)。   ![star > 5000][Gold]
* [GVM](https://github.com/moovweb/gvm) **star:4488** GVM 提供了一个接口来管理 Go 版本。   ![star > 1000][Silver]
* [Wide](https://wide.b3log.org/login)  为使用 Golang 的团队提供基于 web 的 IDE。
* [webhook](https://github.com/adnanh/webhook) **star:4102** 允许用户创建在服务器上执行命令的 HTTP hooks。   ![star > 1000][Silver]
* [gaia](https://github.com/gaia-pipeline/gaia) **star:3752** 可用于任何编程语言来构建强大的管道。   ![star > 1000][Silver]
* [gox](https://github.com/mitchellh/gox) **star:3367** 非常简单，没有多余的跨平台编译工具。   ![star > 1000][Silver]
* [bosun](https://github.com/bosun-monitor/bosun) **star:2856** 按照时间轴发出告警的框架。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [bombardier](https://github.com/codesenberg/bombardier) **star:1755** 快速跨平台 HTTP 基准测试工具。   ![star > 1000][Silver]
* [goxc](https://github.com/laher/goxc) **star:1625** 专注于跨平台编译和打包的 Go 构建工具。   ![star > 1000][Silver]   ![最近一年没有更新][Yellow]
* [fac](https://github.com/mkchoi212/fac) **star:1614** 修复 git 合并冲突。   ![star > 1000][Silver]
* [kala](https://github.com/ajvb/kala) **star:1361** 简单、现代和高性能的作业调度程序。   ![star > 1000][Silver]
* [StatusOK](https://github.com/sanathp/statusok) **star:1164** 监视您的网站和REST api。当服务器宕机或响应时间超过预期时，通过Slack、电子邮件获得通知。   ![star > 1000][Silver]
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:1000** 小型实用程序/库，针对大型对象在Amazon S3中的高速传输进行了优化。   ![star > 1000][Silver]
* [script](https://github.com/bitfield/script) **star:931** 让DevOps编写类shell和系统管理任务变得更加容易。   ![star > 100][Bronze]
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:675** 允许你的 Go应用程序 进行自我更新。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [skm](https://github.com/TimothyYe/skm) **star:550** SKM是一个简单而强大的SSH密钥管理器，它可以帮助您轻松地管理多个SSH密钥!   ![star > 100][Bronze]
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:539** 从命令行管理 BareMetal 服务器(与使用Docker一样容易)。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Pomerium](https://github.com/pomerium/pomerium) **star:523** Pomerium是一个可识别身份的访问代理。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [aurora](https://github.com/xuri/aurora) **star:404** 基于web的跨平台 Beanstalkd 队列服务器控制台。   ![star > 100][Bronze]
* [gonative](https://github.com/inconshreveable/gonative) **star:312** 用原生 Go 创建一个跨平台的 Go 工具链。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [govvv](https://github.com/ahmetalpbalkan/govvv)  可轻松地添加版本信息到 Go 二进制文件。
* [Mora](https://github.com/emicklei/mora) **star:267** 用于访问 MongoDB 文档和元数据的 REST 服务器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Gogs](https://gogs.io/)  自托管的Git服务。
* [godbg](https://github.com/sirnewton01/godbg) **star:219** 基于 web 的 gdb 前端应用程序。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [lstags](https://github.com/ivanilves/lstags) **star:219** 提供了工具和API，可用来同步不同注册中心的Docker图像。   ![star > 100][Bronze]
* [dogo](https://github.com/liudng/dogo) **star:216** 监视源文件中的更改并自动编译和运行(restart)。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [manssh](https://github.com/xwjdsh/manssh) **star:205** manssh是一个命令行工具，可以方便地管理ssh别名配置。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Pewpew](https://github.com/bengadbois/pewpew) **star:203** 灵活的 HTTP 命令行压测工具。   ![star > 100][Bronze]
* [aptly](https://github.com/smira/aptly)  Debian存储库管理工具。
* [gobrew](https://github.com/cryptojuice/gobrew) **star:175** gobrew 允许您轻松地在 go 的多个版本之间切换。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Blast](https://github.com/dave/blast) **star:167** 一个用于API负载测试和批处理作业的简单工具。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [ostent](https://github.com/ostrost/ostent) **star:164** 收集和显示系统指标，并可选 Graphite and/or fluxdb作为依赖。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Packer](https://github.com/mitchellh/packer)  用于从一个源配置为多个平台创建相同的机器图像。
* [grapes](https://github.com/yaronsumel/grapes) **star:135** 旨在轻松地通过ssh分发命令的轻量级工具。   ![star > 100][Bronze]
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:101** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   ![star > 100][Bronze]
* [kcli](https://github.com/cswank/kcli) **star:79** 用于检查kafka主题/分区/消息的命令行工具。
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:71** 用Go编写的托管解决方案，可轻松地在AWS、GCP或DigitalOcean上部署应用程序。   ![最近一周有更新][Green]
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:67** 在Windows机器上远程执行命令的Cli工具。
* [drone-scp](https://github.com/appleboy/drone-scp) **star:55** 通过 SSH 进行文件拷贝。其中 SSH 通过二进制文件、docker 或 Drone CI触发。
* [Dropship](https://github.com/chrismckenzie/dropship) **star:46** 通过 cdn 部署代码的工具。   ![最近一年没有更新][Yellow]
* [Rodent](https://github.com/alouche/rodent) **star:30** 管理Go版本、项目和跟踪依赖项。   ![最近一年没有更新][Yellow]
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:24** 使用二进制文件、docker或 Drone CI 来触发下游Jenkins作业。
* [awsenv](https://github.com/soniah/awsenv) **star:21** 可以为 profile 加载Amazon (AWS)环境变量的轻量二进制文件。   ![最近一年没有更新][Yellow]
* [DepCharge](https://github.com/centerorbit/depcharge) **star:9** Helps orchestrating the execution of commands across the many dependencies in larger projects.
* [lwc](https://github.com/timdp/lwc) **star:8** UNIX wc命令的实时更新版本。   ![最近一年没有更新][Yellow]
* [sg](https://github.com/ChristopherRabotin/sg) **star:5** 可测试一组HTTP极限(如ab)，可以在每个调用之间使用响应代码和数据，根据之前的响应来确定特定的服务器压力。   ![最近一年没有更新][Yellow]

### 其他软件

* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:8284** 快速、简单、可伸缩的分布式文件系统，采用了O(1)磁盘查找。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [rkt](https://github.com/coreos/rkt)  一个应用容器，与其他容器格式(如Docker)兼容，并支持其他执行引擎(如KVM)。
* [restic](https://github.com/restic/restic) **star:7528** 消除重复项备份程序。   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [confd](https://github.com/kelseyhightower/confd) **star:6439** 使用 etcd 或 consul 的模板和数据管理本地应用程序配置文件。   ![star > 5000][Gold]
* [Comcast](https://github.com/tylertreat/Comcast) **star:6181** 模拟坏的网络连接。   ![star > 5000][Gold]
* [LiteIDE](https://github.com/visualfc/liteide) **star:5511** 简单的、开源的、跨平台的Go IDE。   ![star > 5000][Gold]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [drive](https://github.com/odeke-em/drive) **star:4976** 基于命令行的谷歌驱动器客户端。   ![star > 1000][Silver]
* [orange-cat](https://github.com/noraesae/orange-cat)  用Go编写的Markdown预览器。
* [nes](https://github.com/fogleman/nes) **star:4135** 任天堂娱乐系统(NES)模拟器。   ![star > 1000][Silver]
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:3956** 为自动化测试模拟网络和系统条件的代理。   ![star > 1000][Silver]
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [Pipe](https://github.com/b3log/pipe) **star:3026** 一个小巧漂亮的博客平台。   ![star > 1000][Silver]   ![最近一周有更新][Green]
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:2697** 跨平台网络和云备份工具。   ![star > 1000][Silver]
* [myLG](https://github.com/mehrdadrad/mylg) **star:2201** 命令行网络诊断工具。   ![star > 1000][Silver]
* [GoBoy](https://github.com/Humpheh/goboy) **star:2109** 用 Go 编写的任天堂Game Boy彩色模拟器。   ![star > 1000][Silver]
* [syncthing](https://syncthing.net/)  开放，分散的文件同步工具和协议。
* [Stack Up](https://github.com/pressly/sup) **star:1996** Stack Up 是一个超级简单的部署工具 — 只面向Unix。   ![star > 1000][Silver]
* [limetext](http://limetext.org/)  一个强大而优雅的文本编辑器。
* [lgo](https://github.com/yunabe/lgo) **star:1805** 与 Jupyter 可进行交互 Go 程序。它支持代码完成、代码检查以及与Go 100% 兼容性。   ![star > 1000][Silver]
* [snap](https://github.com/intelsdi-x/snap) **star:1803** 强大的遥测框架。   ![star > 1000][Silver]
* [Circuit](https://github.com/gocircuit/circuit) **star:1787** Circuit 是一个可编程平台即服务(PaaS)和/或基础设施即服务(IaaS)，用于管理、发现、同步和编排包含云应用程序的服务和主机。   ![star > 1000][Silver]
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:878** App that displays updates for the Go packages in your GOPATH.   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Documize](https://github.com/documize/community) **star:833** 集成了SaaS工具数据的现代wiki软件。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [scc](https://github.com/boyter/scc) **star:775** 一个非常快速准确的代码计数器，采用了复杂的计算和 COCOMO 预估。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [Leaps](https://github.com/jeffail/leaps) **star:640** 使用操作转换的成对编程服务。   ![star > 100][Bronze]
* [peg](https://github.com/pointlander/peg) **star:615** 解析表达式语法，是Packrat解析器生成器的实现。   ![star > 100][Bronze]
* [vFlow](https://github.com/VerizonDigital/vflow) **star:600** 高性能、可伸缩和可靠的 IPFIX、sFlow和 Netflow 收集器。   ![star > 100][Bronze]
* [gfile](https://github.com/Antonito/gfile) **star:502** 通过WebRTC在两台计算机之间安全地传输文件，不需要任何第三方依赖。   ![star > 100][Bronze]
* [GoDNS](https://github.com/timothyye/godns) **star:437** 一个动态DNS客户端工具，支持DNSPod & HE.net。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [shell2http](https://github.com/msoap/shell2http) **star:421** 通过http服务器执行shell命令(用于原型或远程控制)。   ![star > 100][Bronze]
* [mockingjay](https://github.com/quii/mockingjay-server) **star:415** 一份配置文件中便可伪造HTTP服务器与用户之间的行为。您还可以使服务器随机宕机，以帮助进行更实际的性能测试。   ![star > 100][Bronze]
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:375** 视频流 torrent 客户端。   ![star > 100][Bronze]   ![最近一周有更新][Green]
* [gocc](https://github.com/goccmack/gocc) **star:344** Gocc是一个用Go编写的编译器工具包。   ![star > 100][Bronze]
* [wellington](https://github.com/wellington/wellington) **star:290** Sass 项目管理工具，使用sprite函数(如Compass)扩展语言。   ![star > 100][Bronze]
* [ipe](https://github.com/dimiro1/ipe) **star:277** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   ![star > 100][Bronze]
* [ide](https://github.com/thestrukture/ide) **star:252** 基于浏览器的IDE   ![star > 100][Bronze]
* [Cherry](https://github.com/rafael-santiago/cherry) **star:192** 微型网络聊天服务器。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Orbit](https://github.com/gulien/orbit) **star:128** 一个根据模板来运行命令和生成文件的简单小工具。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [Juju](https://jujucharms.com/)  Cloud-agnostic的服务部署和编制 —— 支持EC2、Azure、Openstack、MAAS等。
* [joincap](https://github.com/assafmo/joincap) **star:122** 用于合并多个pcap文件的命令行实用程序。   ![star > 100][Bronze]
* [Docker](http://www.docker.com/)  面向开发人员和系统管理员的分布式应用程序的开放平台。
* [DDNS](https://github.com/skibish/ddns) **star:98** 个人 DDNS 客户端。
* [boxed](https://github.com/tejo/boxed) **star:72** 基于Dropbox的博客引擎。   ![最近一年没有更新][Yellow]
* [borg](https://github.com/crufter/borg)  基于终端的bash代码段搜索引擎。
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** 基于加密管的简单的NaCL EC25519工具。
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** 测试你的终端。
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** 当有人通过 Tsuru 部署任何应用程序时，会通知您的团队以及其他工具。   ![最近一年没有更新][Yellow]
* [GoLand](https://jetbrains.com/go)  功能齐全的跨平台 Go IDE。
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:12** 包含了 Go 使用手册文档的 Chrome 扩展。   ![最近一年没有更新][Yellow]
* [hugo](http://gohugo.io/)  快速、现代的静态网站引擎。
* [Gor](https://github.com/buger/gor)  Http 流量复制工具，用于实时回放从生产环境到阶段/开发环境的流量。

# 资源

*在哪里可以找到新的Go库。 (翻译出错了? 试试 [英文版](README_EN.md#resources) 吧~)*

## 基准

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1261** HTTP请求路由器基准测试和比较。   ![star > 1000][Silver]
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1008** web框架基准测试。   ![star > 1000][Silver]
* [skynet](https://github.com/atemerev/skynet) **star:915** 天网 1M 线程微基准测试。   ![star > 100][Bronze]
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:875** Go序列化方法的基准测试。   ![star > 100][Bronze]
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Go语言常用基本操作的基准测试。
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:172** 对比各种图像大小调整算法性能。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:123** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:93** 计算语言结构的速度的微观基准测试。
* [autobench](https://github.com/davecheney/autobench) **star:89** 用来来比较不同Go版本之间的性能的框架。   ![最近一年没有更新][Yellow]
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:49** 为流行的 Go 数据库/SQL实用程序收集基准测试。   ![最近一年没有更新][Yellow]
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:19** 强大的HTTP基准测试工具，包含了Аb，Wrk，Siege工具。收集统计和各种参数指标，并比较相关结果。   ![最近一年没有更新][Yellow]
* [kvbench](https://github.com/jimrobinson/kvbench) **star:14** K / V 类型数据库基准测试。   ![最近一年没有更新][Yellow]

## 会议

* [Capital Go](http://www.capitalgolang.com)  华盛顿特区。,美国。
* [GopherCon Europe](https://gophercon.is/)  雷克雅未克,冰岛。
* [GothamGo](http://gothamgo.com/)  美国纽约市。
* [GopherCon Vietnam](https://gophercon.vn/)  越南胡志明市。
* [GopherCon Singapore](https://gophercon.sg)  新加坡枫树商贸城。
* [GopherCon Russia](https://www.gophercon-russia.ru)  莫斯科,俄罗斯。
* [GopherCon Israel](https://www.gophercon.org.il/)  特拉维夫,以色列。
* [GopherCon India](https://www.gophercon.in/)  印度浦那。
* [GopherCon Brazil](https://gopherconbr.org)  弗洛,BR。
* [dotGo](http://www.dotgo.eu)  巴黎,法国。
* [GopherCon Australia](https://gophercon.com.au/)  澳大利亚悉尼。
* [GopherCon](http://www.gophercon.com/)  美国丹佛。
* [GopherChina](http://gopherchina.org)  上海,中国。
* [GolangUK](http://golanguk.com/)  伦敦,英国。
* [GoLab](http://golab.io/)  佛罗伦萨,意大利。
* [GoDays](https://www.godays.io/)  德国柏林。
* [GoCon](http://gocon.connpass.com/)  东京,日本。
* [GoWayFest](https://goway.io/)  白俄罗斯明斯克。

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org)  一本关注 Go 语法/语义和各种细节的书。
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) **star:6824** 一份精选的 Go 书籍清单。   ![star > 5000][Gold]
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:10** in Persian.
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1545** 由 Maria Letta 提供的与 Gopher 有关的图片包，其中包含了插图,表情文字。   ![star > 1000][Silver]
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:65** 可爱的 gopher 标识。   ![最近一年没有更新][Yellow]
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:31** 与 Go gopher 相关的媒介数据[。ai . svg)。   ![最近一年没有更新][Yellow]
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:1859** 阿什莉·麦克纳马拉的歌斐艺术品。   ![star > 1000][Silver]
* [gophers](https://github.com/egonelbre/gophers) **star:1605** Free gophers.   ![star > 1000][Silver]
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:315** Gopherize自己。   ![star > 100][Bronze]
* [gophers](https://github.com/rogeralsing/gophers) **star:50** 随机gopher图形。   ![最近一年没有更新][Yellow]
* [gophers](https://github.com/sillecelik/go-gopher) **star:40** Gopher amigurumi玩具图案。

## 聚会

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
* [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
* [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
* [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
* [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
* [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Казань](https://www.meetup.com/GolangKazan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*在这里添加您所在城市/国家的群组(发送**PR**) (翻译出错了? 试试 [英文版](README_EN.md#meetups) 吧~)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## 网站

* [Go Report Card](https://goreportcard.com)  为你的 Go 包生成一份报告单。
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:24670** 其他 awesome 系列的列表。   ![star > 5000][Gold]
* [CodinGame](https://www.codingame.com/)  以小游戏互动完成任务的形式来学习 Go。
* [Go Blog](http://blog.golang.org)  官方 Go 博客。
* [Go Challenge](http://golang-challenge.org/)  通过解决问题并从 Go 专家那里得到反馈来学习 Go。
* [Go Community on Hashnode](https://hashnode.com/n/go)  Hashnode上的Go社区。
* [Go Forum](https://forum.golangbridge.org)  讨论 Go 的论坛。
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  wiki上的 Go 社区项目列表。
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:14578** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.   ![star > 5000][Gold]   ![最近一周有更新][Green]
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:142** 收藏的 Go 图像，图形和艺术作品。   ![star > 100][Bronze]   ![最近一年没有更新][Yellow]   ![归档项目][Archived]
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  加入我们为Gophers设立的全新Slack社区([了解它是如何产生的](https://blog.gopheracademy.com/gophers-slack-community/))。
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  Google+社区 golang爱好者聚集地。
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go 邮件列表。
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  寻找最新的 Go库 的好地方。
* [justforfunc](https://www.youtube.com/c/justforfunc)  致力于传授 Go 编程语言技巧和技巧的Youtube 频道，由Francesc Campoy [@francesc](https://twitter.com/francesc)主办。
* [r/Golang](https://www.reddit.com/r/golang)  与 Go 有关的新闻。
* [studygolang](https://studygolang.com)  Go语言中文网
* [gowalker.org](https://gowalker.org)   Go API 文档。
* [Gophercises](https://gophercises.com/)  为 Go 初学者提供免费的代码训练。
* [Awesome Go @LibHunt](https://go.libhunt.com)  属于你的 Go 工具箱。
* [Golang News](https://golangnews.com)  关于 Go 编程的链接和新闻。
* [Golang Flow](https://golangflow.io)  发布更新、新闻、包等等。
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:36** 专门收集需要帮助的Go项目，这是你开启开源之路的好去处。   ![最近一年没有更新][Yellow]
* [godoc.org](https://godoc.org/)  开源 Go 包的文档。
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### 教程

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:31641** Golang电子书，主要讲述如何用 Golang 建立一个web应用程序。   ![star > 5000][Gold]   ![包含中文文档][CN]
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  如何缓存数据库的慢查询。
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  如何取消MySQL查询。
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:4047**  Go's reference card。   ![star > 1000][Silver]
* [Go database/sql tutorial](http://go-database-sql.org/)  数据库概论/ sql。
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  在你的移动设备上编辑你编辑和运行你的 Go 代码。
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:457** 一本讲述如何用 Go 进行以太开发的小册。   ![star > 100][Bronze]   ![包含中文文档][CN]
* [Games With Go](http://gameswithgo.org/)  关于编程和游戏开发系列教学视频。
* [Go By Example](https://gobyexample.com/)  手把手教你 如何在 Go 应用程序中使用注释。
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Go 初学者经常遇到的陷阱、误区、错误
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  教你如何用 Go 搭建一个电商平台 (包括demo)。
* [A Tour of Go](http://tour.golang.org/)  互动的 Go 之旅。
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:4593** 学习使用测试驱动开发。   ![star > 1000][Silver]   ![最近一周有更新][Green]   ![包含中文文档][CN]
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  面向 Golang 初学者教程。
* [package main](https://www.youtube.com/packagemain)  关于 Go 编程的YouTube频道。
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:699** 引入示例讲述 Golang 与Node.js在学习上的差异。   ![star > 100][Bronze]
* [Golangbot](https://golangbot.com/learn-golang-series/)  Go 编程教程。
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Go社区投票选举出来的最好的在线 Go 教程。
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  快速使用Godog —— 一个行为驱动开发的构建和测试应用程序框架。
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1133** 由一个经验丰富的Go程序员群体编写的一系列Go学习范例。   ![star > 1000][Silver]
* [Your basic Go](http://yourbasic.org/golang)  如何收集大量的教程。

