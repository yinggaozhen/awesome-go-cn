# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "最近一周有更新"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "最近一年没有更新"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "包含中文文档"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "项目已归档"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc文档地址"

**此项目是 [awesome-go](https://awesome-go.com/) 中文版，最后一次同步时间 : 2020-02-25 11:48:27(每隔1天同步一次)**

[![chinese](https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/english.svg)](README_EN.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) 为Awesome Go打赏~

精选了一系列很棒的Go框架、库和软件。灵感来自于[awesome-python](https://github.com/vinta/awesome-python)。

**小图标说明** :


小图标 | 说明  
:-:|-
![awesome][Awesome] | star > 2000
![最近一个周有更新][Green] | 最近一周有更新。可以基本判断当前库处于积极维护状态。
![最近一年未更新][Yellow] | 最近一年没有更新。反应了此库的维护积极性不高，使用时需谨慎。
![归档项目][Archived] | 此项目已归档，不再更新，使用时需谨慎。
![此项目有中文文档][CN] | 此项目有中文文档。
![godoc文档][GoDoc] | godoc文档地址。

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
    - [动态域名](#动态域名)
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
    - [服务器应用程序](#服务器应用程序)
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

* [PortAudio](https://github.com/gordonklaus/portaudio)  基于 Go 的PortAudio audio I/O库。
* [Oto](https://github.com/hajimehoshi/oto) **star:537** 多平台的 low-level 声音播放库。   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:283** 基于 Go 的音乐理论模型。   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:274** 通过音频流生成波形图像的包。   [![最近一年没有更新][Yellow]](https://github.com/mdlayher/waveform)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:225** PortMidi的 Go 语言实现接口.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [mix](https://github.com/go-mix/mix) **star:112** 基于序列的 Go 原生音乐混音器。   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:105** 原生 Go MP3解码器。   [![最近一年没有更新][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:104** libsox 的 Go 语言实现接口。   [![最近一年没有更新][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [taglib](https://github.com/wtolson/go-taglib) **star:72** taglib 的 Go 语言实现接口.   [![最近一年没有更新][Yellow]](https://github.com/wtolson/go-taglib)   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:67** 原生 Go AAC位流解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:37** 轻量级MP3解码器库。
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:28** libmediainfo 的 Go 语言实现接口。   [![最近一年没有更新][Yellow]](https://github.com/zhulik/go_mediainfo)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/go_mediainfo)
* [vorbis](https://github.com/mccoyst/vorbis) **star:25** “原生” Go Vorbis解码器(使用CGO，但没有依赖关系)。   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [id3v2](https://github.com/bogem/id3v2)  快速稳定的 ID3 解析及写入Go库。
* [malgo](https://github.com/gen2brain/malgo)  迷你音频库。
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:9** libsamplerate 的 Go 语言实现接口。   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)
* [flac](https://github.com/mewkiz/flac)  原生 Go FLAC编码器/解码器，支持FLAC流。
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI)  EasyMidi是一个简单可靠的库，用于处理标准midi文件(SMF)。

## 身份验证和OAuth

*用于实现验证方案的库。 (翻译出错了? 试试 [英文版](README_EN.md#authentication-and-oauth) 吧~)*

* [loginsrv](https://github.com/tarent/loginsrv)  JWT登录微服务带有可插拔的后端服务，如OAuth2 (Github)、htpasswd、osiam。
* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:7098** JSON Web令牌(JWT)。   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![最近一周有更新][Green]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [rbac](https://github.com/zpatrick/rbac)  最小的RBAC包。
* [permissions2](https://github.com/xyproto/permissions2)  用于跟踪用户、登录状态和权限的库。依赖于cookie安全和bcrypt。
* [paseto](https://github.com/o1egl/paseto)  平台无关的安全令牌(PASETO)。
* [osin](https://github.com/openshift/osin)  OAuth2服务器库。
* [oauth2](https://github.com/golang/oauth2) **star:2739** goauth2的继任者。通用OAuth 2.0包，附带JWT、谷歌api、计算引擎和应用程序引擎支持。   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2566** 提供了 OAuth 和 OAuth2 的简单清晰易用的方法。可开箱即用处理多个提供程序。   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![最近一周有更新][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [gologin](https://github.com/dghubble/gologin) **star:1162** 用于使用OAuth1和OAuth2身份验证提供者登录的可链处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:986** 轻量级的基于角色的访问控制(RBAC)实现。   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [scs](https://github.com/alexedwards/scs) **star:675** HTTP服务器的会话管理器。   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [branca](https://github.com/hako/branca)  基于 Go 实现Branca令牌。
* [authboss](https://github.com/volatiletech/authboss)  web模块化认证系统。它试图删除尽可能多的模板文件和硬编码，以便每次新建一个新的web项目时，您都可以插入、配置并开始构建您的应用程序，而不必每次都构建一个身份验证系统。
* [casbin](https://github.com/hsluoyz/casbin)  支持ACL、RBAC、ABAC等访问控制模型的授权库。
* [httpauth](https://github.com/goji/httpauth) **star:190** HTTP身份验证中间件。   [![最近一年没有更新][Yellow]](https://github.com/goji/httpauth)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:189** 简单、灵活、安全和惯用的web会话管理，具有可配置化的后端。   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:169** JWT中间件，可用于Golang http服务器，提供了许多配置选项。   [![最近一年没有更新][Yellow]](https://github.com/adam-hanna/jwt-auth)   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [session](https://github.com/icza/session) **star:98** web服务器会话管理(包括支持谷歌应用程序引擎 - GAE)。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [jwt](https://github.com/robbert229/jwt) **star:78** 简单易用的JSON Web令牌实现(JWT)。   [![最近一年没有更新][Yellow]](https://github.com/robbert229/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [jwt](https://github.com/pascaldekloe/jwt)  轻量级JSON Web令牌库。
* [sessionup](https://github.com/swithek/sessionup) **star:77** 简单但有效的HTTP会话管理和标识包。   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [securecookie](https://github.com/chmike/securecookie) **star:33** 高效安全的cookie编码/解码。   [![最近一周有更新][Green]](https://github.com/chmike/securecookie)   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [signedvalue](https://github.com/sashka/signedvalue) **star:8** 与[Tornado's](https://github.com/tornadooweb/tornado) 完全兼容的签名和时间戳字符串实现。   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)
* [sessions](https://github.com/adam-hanna/sessions)  非常简单，高性能，可深度定制的会话服务，主要用于的 go http 服务器。
* [sjwt](https://github.com/brianvoe/sjwt)  简单的jwt生成器和解析器。
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** 使用SessionGate Redis模块进行会话管理。   [![最近一年没有更新][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [scope](https://github.com/SonicRoshan/scope) **star:4** 在Go中轻松管理OAuth2范围。   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** 提供cookie .txt文件格式的解析器。   [![最近一年没有更新][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)
* [go-jose](https://github.com/square/go-jose)  相当完整地实现了JOSE工作组的JSON Web令牌、JSON Web签名和JSON Web加密规范。
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server)  用 Golang 编写的独立且符合规范的OAuth2服务器。

## Bot建设

*用于构建和使用机器人的库。 (翻译出错了? 试试 [英文版](README_EN.md#bot-building) 吧~)*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api)  简单轻量级的Telegram bot客户端。
* [telebot](https://github.com/tucnak/telebot) **star:1125** 用Go编写的Telegram bot框架。   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [slacker](https://github.com/shomali11/slacker) **star:356** 可简单创建Slack机器人的框架。   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:287** 基于控制台的，用于加密货币交易所的的交易机器人。   [![最近一年没有更新][Yellow]](https://github.com/saniales/golang-crypto-trading-bot)   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [Kelp](https://github.com/stellar/kelp) **star:265** 官方交易和做市机器人为[Stellar](https://www.stellar.org/) DEX。开箱即用的作品，用 Golang 编写，兼容集中交易和定制交易策略。   [![最近一周有更新][Green]](https://github.com/stellar/kelp)   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:168** 面向服务的IRC bot，使用Redis和JSON进行消息传递。   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:93** Libary为twitch编写机器人程序。电视聊天   [![godoc][GoDoc]](https://godoc.org/github.com/gempir/go-twitch-irc)
* [margelet](https://github.com/zhulik/margelet) **star:60** 构建电报机器人的框架。   [![最近一年没有更新][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:30** 简单的Go [VK](https://vk.com) bot库。   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [hanu](https://github.com/sbstjn/hanu)  用于编写Slack机器人的框架。
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:24** 另一个构建Slack机器人的框架。   [![最近一周有更新][Green]](https://github.com/alexandre-normand/slackscot)   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [tbot](https://github.com/yanzay/tbot)  带有类似于net/http API的Telegram bot服务器。
* [micha](https://github.com/onrik/micha) **star:12** 基于 GO 实现的Telegram 机器人API库。   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)
* [go-joe](https://joe-bot.net)  一个通用的机器人库的灵感来自于Hubot，但写在围棋。
* [go-tgbot](https://github.com/olebedev/go-tgbot)  由swagger文件、基于会话的路由器和中间件生成的纯Golang Telegram Bot API包装器。
* [go-sarah](https://github.com/oklahomer/go-sarah)  此框架提供了聊天机器人相关的服务，包括LINE、Slack、Gitter等。
* [go-chat-bot](https://github.com/go-chat-bot/bot)  用 Go 编写的IRC, Slack和电报机器人。

## 命令行

### 标准CLI

*用于构建标准或基本命令行应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#standard-cli) 吧~)*

* [cobra](https://github.com/spf13/cobra) **star:15875** 现代Go CLI命令行交互工具。   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![最近一周有更新][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [wlog](https://github.com/dixonwille/wlog)  支持跨平台和并发的简单日志记录接口。
* [urfave/cli](https://github.com/urfave/cli) **star:13159** 可让你简单、快速和愉快的构建命令行应用(之前是codegangsta/cli)。   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [liner](https://github.com/peterh/liner)  类似readline-like的命令行接口库。
* [kingpin](https://github.com/alecthomas/kingpin) **star:2811** 支持子命令的命令行和标志解析器。   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![最近一周有更新][Green]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1636**  Go 命令行选项解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [Dnote](https://github.com/dnote/dnote) **star:1498** 一个简单的命令行笔记本与多设备同步。   [![最近一周有更新][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [readline](https://github.com/chzyer/readline) **star:1463** 纯golang实现，在MIT许可下提供了GNU-Readline的大部分特性。   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [sand](https://github.com/Zaba505/sand)  用于创建解释器等的简单API。
* [docopt.go](https://github.com/docopt/docopt.go) **star:1223** 会让你满意的命令行参数解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mow.cli](https://github.com/jawher/mow.cli)  用于构建具有复杂标志和参数解析和验证的CLI应用程序。
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1092** 用于实现命令行接口的Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [pflag](https://github.com/spf13/pflag) **star:997** 基于POSIX/GNU-style --flags实现的包，主要用于替换Go的falg包。   [![最近一周有更新][Green]](https://github.com/spf13/pflag)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [cli-init](https://github.com/tcnksm/gcli) **star:893** 一个简单就可开启构建Golang命令行的应用程序。   [![最近一年没有更新][Yellow]](https://github.com/tcnksm/gcli)   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [go-arg](https://github.com/alexflint/go-arg) **star:835** 基于结构的参数解析。   [![最近一周有更新][Green]](https://github.com/alexflint/go-arg)   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:665** 使用 Go 语言编写的 bash 命令补全工具以及 Go 命令补全工具.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [flaggy](https://github.com/integrii/flaggy) **star:626** 一个健壮的、易用的标志包，具有出色的子命令支持。   [![最近一周有更新][Green]](https://github.com/integrii/flaggy)   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:513** 基于golang结构标签，功能丰富易于使用的命令行包。   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:380** Unikernel 构建器/协调器。   [![最近一周有更新][Green]](https://github.com/nanovms/ops)   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:164** 命令行参数分析器，灵感来自Python的argparse模块。   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [commandeer](https://github.com/jaffee/commandeer) **star:107** 开发友好的CLI应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [sflags](https://github.com/octago/sflags) **star:104** 基于结构的flag生成器，用于flag、urfave/cli、pflag、cobra、kingpin和其他库。   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [flag](https://github.com/cosiner/flag) **star:103** 简单但功能强大的命令行选项解析库，用于支持Go子命令。   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:101** 简小的命令行接口框架。   [![最近一年没有更新][Yellow]](https://github.com/ukautz/clif)   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [wmenu](https://github.com/dixonwille/wmenu) **star:98** 为cli程序提供了简单上手的菜单，可提示用户作出选择。   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [job](https://github.com/liujianping/job) **star:68** 工作，把你的短期指令当作长期任务。   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [cli](https://github.com/teris-io/cli) **star:64** 为 Go 构建命令接口提供简单而完整的API。   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [1build](https://github.com/gopinath-langote/1build) **star:49** 命令行工具，以无摩擦地管理项目特定的命令。   [![最近一周有更新][Green]](https://github.com/gopinath-langote/1build)   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [env](https://github.com/codingconcepts/env) **star:48** 基于标记的结构化的环境配置。   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  具有自动配置和依赖注入的cli应用程序框架。
* [gocmd](https://github.com/devfacet/gocmd) **star:37** 用于构建命令行应用程序。   [![最近一年没有更新][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [strumt](https://github.com/antham/strumt) **star:34** 用于创建提示链。   [![最近一周有更新][Green]](https://github.com/antham/strumt)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [ts](https://github.com/liujianping/ts)  时间戳转换和比较工具。
* [flagvar](https://github.com/sgreben/flagvar) **star:33** 符合 Go 标准的“flag”标志参数类型包。   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [cmdr](https://github.com/hedzr/cmdr) **star:30** 一个POSIX/GNU风格的、类似getopt的命令行UI Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:19**  Go 选择解析器，借鉴于Perl灵活性的GetOpt::Long。   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [argv](https://github.com/cosiner/argv) **star:18** 基于Base 语法，用于分隔命令行字符串并将其作为参数的 Go 语言库，   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [clîr](https://github.com/leaanthony/clir) **star:17** 一个简单而清晰的CLI库。依赖免费的。   [![最近一周有更新][Green]](https://github.com/leaanthony/clir)   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** 用于简化CLI工作流的 Go 库。   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [cmd](https://github.com/posener/cmd) **star:3** 扩展标准的' flag '包，以支持子命令和更多的idomatic方式。   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)

### 高级控制台用户界面

*用于构建控制台应用程序和控制台用户界面的库。 (翻译出错了? 试试 [英文版](README_EN.md#advanced-console-uis) 吧~)*

* [termui](https://github.com/gizak/termui) **star:9532** 此库是基于**termbox-go**实现的，借鉴于[blessed-contrib](https://github.com/yaronn/blessed-contrib)。   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![最近一周有更新][Green]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  更换终端文本样式。
* [gocui](https://github.com/jroimartin/gocui) **star:5985** 旨在创建控制台用户界面的极简Go库。   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1640** 在终端呈现进度条，可灵活配置的。   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [uitable](https://github.com/gosuri/uitable)  改善终端应用程序中表格数据的可读性。
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1308** 在命令行中构建轻量级ASCII线图╭┈╯，应用程序中没有其他依赖项。   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:1094** 用于实时更新终端输出的库。   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [cfmt](https://github.com/mingrammer/cfmt)  提供上下文的fmt，灵感来自于bootstrap color classes。
* [aurora](https://github.com/logrusorgru/aurora) **star:777** 支持fmt.Printf/Sprintf的ANSI终端颜色。   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [go-colorable](https://github.com/mattn/go-colorable) **star:414** 适用于windows的颜色编写器。   [![最近一周有更新][Green]](https://github.com/mattn/go-colorable)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:398** Go 实现的 isatty。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [go-prompt](https://github.com/c-bata/go-prompt)  构建一个强大的交互式提示，借鉴于[python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit)
* [chalk](https://github.com/ttacon/chalk) **star:327** 美化终端/控制台输出。   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [progressbar](https://github.com/schollz/progressbar)  基本线程安全的进度条，在每个操作系统工作。
* [mpb](https://github.com/vbauerster/mpb)  可在终端显示多进度条。
* [gookit/color](https://github.com/gookit/color) **star:319** 终端显色工具库，支持16种颜色，256种颜色，RGB显色输出，兼容Windows。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![包含中文文档][CN]](https://github.com/gookit/color)
* [simpletable](https://github.com/alexeyco/simpletable) **star:207** 可在终端显示简易表格。   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [tabby](https://github.com/cheynewallace/tabby)  一个可在终端生成一个极简Golang表格轻量级库
* [tabular](https://github.com/InVisionApp/tabular)  不需要向API传递大量参数就可从命令行实用程序中打印ASCII表。
* [termbox-go](https://github.com/nsf/termbox-go)  基于文本的跨平台接口库。
* [termdash](https://github.com/mum4k/termdash)  此库是基于**termbox-go**实现的，借鉴于[termui](https://github.com/gizak/termui)。
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:201** 在终端中使用彩色文字。   [![最近一年没有更新][Yellow]](https://github.com/daviddengcn/go-colortext)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [colourize](https://github.com/TreyBastian/colourize) **star:17** 在终端提供ANSI彩色文本。   [![最近一年没有更新][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [ctc](https://github.com/wzshiming/ctc) **star:16** 不需要Print方法的非侵入性跨平台终端颜色库。   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   [![包含中文文档][CN]](https://github.com/wzshiming/ctc)
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** 在终端提供ANSI彩色文本模板。   [![最近一年没有更新][Yellow]](https://github.com/workanator/go-ataman)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)

## 配置

*配置解析的库。 (翻译出错了? 试试 [英文版](README_EN.md#configuration) 吧~)*

* [godotenv](https://github.com/joho/godotenv) **star:2620** Ruby 的 dotenv 库的 Go移植版(从.env文件加载环境变量)。   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [config](https://github.com/JeremyLoy/config) **star:228** 云本地应用程序配置。将ENV绑定到结构体仅需两行代码。   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [config](https://github.com/olebedev/config)  带有环境变量和标记解析的JSON或YAML配置包装器。
* [xdg](https://github.com/OpenPeeDeeP/xdg)  遵循[XDG标准](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)的跨平台包。
* [koanf](https://github.com/knadh/koanf)  轻量级可扩展库，用于读取Go应用程序中的配置。内置支持JSON, TOML, YAML, env，命令行。
* [hjson](https://github.com/hjson/hjson-go)  更加人性化的JSON配置。轻松的语法，更少的错误，更多的注释。
* [ingo](https://github.com/schachmat/ingo)  flag保存在类ini的配置文件中。
* [ini](https://github.com/go-ini/ini)   读和写INI文件。
* [joshbetz/config](https://github.com/joshbetz/config)  一个可解析环境变量、JSON文件小巧的配置库，在SIGHUP时会自动重新加载。
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig)  管理来自环境变量的配置数据。
* [mini](https://github.com/sasbury/mini)  用于解析ini类型的配置文件。
* [konfig](https://github.com/lalamove/konfig)  可组合、可观察和高性能的分布式配置管理。
* [gookit/config](https://github.com/gookit/config)  程序配置管理(load,get,set)。支持JSON, YAML, TOML, INI, HCL。支持多文件加载，数据覆盖合并。
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env)  读取环境变量的简单有用的包。
* [onion](http://github.com/goraz/onion)  基于层配置的Go，支持JSON, TOML, YAML，属性，etcd, env，和加密使用PGP。
* [sprbox](https://github.com/oblq/sprbox)  支持构建环境的工具箱工厂和其他不确定的配置解析器(如YAML、TOML、JSON和环境vars)。
* [store](https://github.com/tucnak/store)  轻量级配置管理器。
* [viper](https://github.com/spf13/viper)  配置管理。
* [harvester](https://github.com/beatlabs/harvester)  一个易于使用的静态和动态配置包
* [gofigure](https://github.com/ian-kent/gofigure) **star:58** 让程序配置变得简单。   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  模块化的JSON配置。保持配置结构及其配置的代码，并将解析委托给子模块，而不牺牲配置的完整序列化。
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config)  从AWS SSM(参数存储)加载配置参数的Go实用程序。
* [genv](https://github.com/sakirsensoy/genv)  使用dotenv支持轻松读取环境变量。
* [gcfg](https://github.com/go-gcfg/gcfg)  将ini的配置文件读入 Go structs中;支持用户定义的类型和子选项。
* [envh](https://github.com/antham/envh)  协助管理环境变量的Helpers。
* [envconfig](https://github.com/vrischmann/envconfig)  从环境变量中读取配置。
* [envconf](https://github.com/ian-kent/envconf)  从环境配置中读取配置。
* [envcfg](https://github.com/tomazk/envcfg)  对环境变量进行解析，并赋值到struct。
* [env](https://github.com/caarlos0/env)  解析环境变量并赋值到struct中(默认值)。
* [conflate](https://github.com/the4thamigo-uk/conflate)  合并来自任意url的多个JSON/YAML/TOML文件、针对JSON模式的验证以及模式中定义的默认值的应用程序。
* [confita](https://github.com/heetch/confita)  从多个后端级联加载配置到struct中。
* [configure](https://github.com/paked/configure) **star:54** 通过多个源提供配置，包括JSON、flags和环境变量。   [![最近一年没有更新][Yellow]](https://github.com/paked/configure)   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [config](https://github.com/golobby/config) **star:49** 一个轻量级但功能强大的配置包，用于Go项目。   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:45** 简约的配置阅读器(来自文件、环境，以及你想要的任何地方)。   [![最近一周有更新][Green]](https://github.com/ilyakaznacheev/cleanenv)   [![godoc][GoDoc]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
* [goConfig](https://github.com/crgimenes/goConfig)  将结构体解析为输入，并用来自命令行、环境变量和配置文件的参数填充该结构体的字段。
* [go-up](https://github.com/ufoscout/go-up) **star:26** 一个简单的配置库，具有递归占位符解析功能。   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)

## 持续集成

*用于帮助进行持续集成的工具。 (翻译出错了? 试试 [英文版](README_EN.md#continuous-integration) 吧~)*

* [CDS](https://github.com/ovh/cds)  企业级CI/CD和DevOps自动化开源平台。
* [drone](https://github.com/drone/drone)  Drone 是一个基于 Docker 的持续集成平台，用 Go 编写。
* [duci](https://github.com/duck8823/duci)  一个简单的 ci 服务。
* [gomason](https://github.com/nikogura/gomason)  在一个干净的工作区中对你的 Go 二进制文件进行测试、构建、签名和发布。
* [goveralls](https://github.com/mattn/goveralls)  Coveralls.io 是一个用 Go 编写，可持续对代码覆盖率进行检测的系统。
* [overalls](https://github.com/go-playground/overalls)  针对多package 的 Go 语言项目，可为类似 goveralls 这样的工具生成覆盖率报告。
* [roveralls](https://github.com/LawrenceWoodman/roveralls)  递归覆盖测试工具。

## CSS预处理器

*用于预处理CSS文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#css-preprocessors) 吧~)*

* [gcss](https://github.com/yosssi/gcss)  纯Go编写的 CSS 预处理器。
* [go-libsass](https://github.com/wellington/go-libsass)   采用 Go封装，100% 与 Sass 兼容的 libsass 项目。

## 数据结构

*用 Go 实现的通用的数据结构和算法。 (翻译出错了? 试试 [英文版](README_EN.md#data-structures) 吧~)*

* [algorithms](https://github.com/shady831213/algorithms)  算法和数据结构。来源于CLRS。
* [pipeline](https://github.com/hyfather/pipeline)  实现了 fan-in 和 fan-out 的管道功能。
* [hide](https://github.com/emvi/hide)  ID type with marshalling to/from hash to prevent sending IDs to clients.
* [hilbert](https://github.com/google/hilbert)  用于映射空间填充曲线（例如 Hilbert 曲线和 Peano 曲线）和数值。
* [hyperloglog](https://github.com/axiomhq/hyperloglog)  HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.
* [iter](https://github.com/disksing/iter)  Go实现的c++ STL迭代器和算法。
* [levenshtein](https://github.com/agext/levenshtein)  Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [levenshtein](https://github.com/agnivade/levenshtein)  实现在Go中计算levenshtein距离。
* [mafsa](https://github.com/smartystreets/mafsa)  实现了 MA-FSA ，包含最小完美哈希。
* [merkletree](https://github.com/cbergoon/merkletree)  实现了merkle树，提供了对数据结构内容的有效和安全的验证。
* [mspm](https://github.com/BlackRabbitt/mspm)  用于信息检索的多字符串模式匹配算法。
* [null](https://github.com/emvi/null)  对空的 Go 数据类型也可以进行JSON进行解析/反解析。
* [parsefields](https://github.com/MonaxGT/parsefields)  用于解析类似json的日志的工具，用于收集惟一的字段和事件。
* [ptrie](https://github.com/viant/ptrie)  前缀树的一种实现。
* [gostl](https://github.com/liyue201/gostl)  用于go的数据结构和算法库，旨在提供类似c++ STL的函数。
* [remember-go](https://github.com/rocketlaunchr/remember-go)  用于缓存慢速数据库查询的通用接口(支持redis、memcached、ristretto或in-memory)。
* [ring](https://github.com/TheTannerRyan/ring)  高性能、线程安全的bloom过滤器。
* [roaring](https://github.com/RoaringBitmap/roaring)  实现了压缩 bitsets 的Go包。
* [set](https://github.com/StudioSol/set)  使用LinkedHashMap在Go中实现简单的set数据结构。
* [skiplist](https://github.com/MauriceGit/skiplist)  高性能的 Go 跳表实现。
* [skiplist](https://github.com/gansidui/skiplist)  在Go中实现了跳表。
* [timedmap](https://github.com/zekroTJA/timedmap)  实现了有生命周期的键值对Map。
* [treap](https://github.com/perdata/treap)  使用树堆进行持久、快速有序的映射。
* [trie](https://github.com/derekparker/trie)  在Go中实现Trie。
* [ttlcache](https://github.com/diegobernardes/ttlcache)  基于内存的LRU算法实现。
* [typ](https://github.com/gurukami/typ)  从复杂结构中获取值，支持空类型、安全的类型转换。
* [gota](https://github.com/kniren/gota)  实现了数据帧，序列以及数据噪音。
* [goskiplist](https://github.com/ryszard/goskiplist)  基于 Go 的跳表实现。
* [binpacker](https://github.com/zhuangsirui/binpacker)  帮助用户构建自定义二进制流的二进制封装器和解包器
* [deque](https://github.com/gammazero/deque)  快速环缓冲区deque(双端队列)。
* [bit](https://github.com/yourbasic/bit)  Go 语言集合数据结构。提供了额外的位操作功能。
* [bitset](https://github.com/willf/bitset)  实现了 bitsets 的 Go 包。
* [bloom](https://github.com/zhenjl/bloom)  在Go中实现了Bloom过滤器。
* [bloom](https://github.com/yourbasic/bloom)  Golang Bloom过滤器的实现。
* [boomfilters](https://github.com/tylertreat/BoomFilters)  用于处理连续的概率数据结构。
* [concurrent-writer](https://github.com/free/concurrent-writer)  具备高并发能力，可替换 bufio.Writer。
* [conjungo](https://github.com/InVisionApp/conjungo)  一个小型、强大和灵活的合并库。
* [count-min-log](https://github.com/seiflotfy/count-min-log)  Go实现Count-Min-log sketch的功能 : 使用近似计数器进行近似计数(类似Count-Min sketch，但使用更少内存)。
* [crunch](https://github.com/superwhiskers/crunch)  打包实现缓冲区，以便轻松处理各种数据类型。
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter)  布谷鸟过滤器:一个用Go实现，可替代计数 bloom 过滤器。
* [deque](https://github.com/edwingeng/deque)  高度优化的双端队列。
* [dict](https://github.com/srfrog/dict)  实现类似 python dict的功能(dict)。
* [goset](https://github.com/zoumo/goset)  一个有用的Go集合实现。
* [encoding](https://github.com/zhenjl/encoding)  整形压缩库。
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)  自适应基数树。
* [go-datastructures](https://github.com/Workiva/go-datastructures)  可靠的、高性能的和线程安全的数据结构的集合。
* [go-ef](https://github.com/amallia/go-ef)  实现了 Elias-Fano 编码。
* [go-geoindex](https://github.com/hailocab/go-geoindex)  基于内存的地理索引。
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache)  基于内存的实现了高性能的key:value存储库。指针缓存。
* [go-rquad](https://github.com/aurelien-rainone/go-rquad)  区域四叉树具有高效的点定位和邻域查找功能。
* [gocache](https://github.com/eko/gocache)  一个完整的Go缓存库，具有多个存储(内存，memcache, redis，…)，可链，可加载，指标缓存和更多。
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue)  并行FIFO队列。
* [gods](https://github.com/emirpasic/gods)  数据结构。容器、集合、列表、堆栈、地图、BidiMaps、树、HashSet等。
* [gofal](https://github.com/xxjwxc/gofal)  基于 Go 实现的分数计算。包含分子、分母运算 
* [golang-set](https://github.com/deckarep/golang-set)  线程安全和非线程安全的高性能集。
* [willf/bloom](https://github.com/willf/bloom)  实现Bloom过滤器。

## 数据库

*数据库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [badger](https://github.com/dgraph-io/badger)  快速 K/V 存储。
* [prometheus](https://github.com/prometheus/prometheus)  用于监控系统和时序的数据库。
* [influxdb](https://github.com/influxdb/influxdb)  可伸缩的数据存储，用于指标衡量、事件和实时分析。
* [Kivik](https://github.com/go-kivik/kivik)  Kivik为CouchDB、PouchDB和类似的数据库提供了一个通用的Go和GopherJS客户端库。
* [ledisdb](https://github.com/siddontang/ledisdb)  Ledisdb是一种高性能的NoSQL，类似于基于LevelDB的Redis。
* [levigo](https://github.com/jmhodges/levigo)  实现了对LevelDB封装。
* [moss](https://github.com/couchbase/moss)  Moss是一个用100% Go编写的简单LSM键值存储引擎。
* [nutsdb](https://github.com/xujiajun/nutsdb)  Nutsdb是一个用纯Go编写的简单、快速、可嵌入、持久的键/值存储。它支持完全序列化的事务和许多数据结构，如列表、集合、排序集。
* [piladb](https://github.com/fern4lvarez/piladb)  基于堆栈数据结构的轻量级RESTful数据库引擎。
* [pudge](https://github.com/recoilme/pudge)  使用Go的标准库编写的快速和简单的键/值存储。
* [gostore](https://github.com/twharmon/gostore)  Gostore是一个简单、持久的嵌入式键值存储引擎，用Go编写。
* [rqlite](https://github.com/rqlite/rqlite)  基于SQLite的轻量级分布式关系数据库。
* [Scribble](https://github.com/nanobox-io/golang-scribble)  小型平面文件JSON存储。
* [slowpoke](https://github.com/recoilme/slowpoke)  具有持久性的键值存储。
* [tempdb](https://github.com/rafaeljesus/tempdb)  用于临时数据存放的 K/V 存储。
* [tidb](https://github.com/pingcap/tidb)  TiDB是一个分布式SQL数据库。灵感来自谷歌F1的设计。
* [tiedot](https://github.com/HouzuoGuo/tiedot)  属于你的NoSQL数据库。
* [tracedb](https://github.com/unit-io/tracedb)  快速timeseries数据库物联网，实时消息传递应用程序。使用pubsub通过tcp或websocket访问tracedb，使用github.com/unit-io/trace应用程序。
* [Vasto](https://github.com/chrislusf/vasto)  分布式高性能键值存储。可做磁盘备份。最终一致。高可用。能够在不中断服务的情况下增长或收缩。
* [groupcache](https://github.com/golang/groupcache)  Groupcache是一个缓存和缓存填充库，在许多情况下，它是memcached的替代品。
* [gorocksdb](https://github.com/kapitan-k/gorocksdb)  用 Go 对[RocksDB](https://rocksdb.org)实现了封装。
* [bbolt](https://github.com/etcd-io/bbolt)  一个用于Go的嵌入式键/值数据库。
* [Coffer](https://github.com/claygod/coffer)  支持事务的简单ACID键值数据库。
* [bcache](https://github.com/iwanbk/bcache)  基于内存的最终一致的分布式缓存。
* [BigCache](https://github.com/allegro/bigcache)  高效的键/值缓存为千兆字节的数据。
* [Bitcask](https://github.com/prologic/bitcask)  Bitcask是一个可嵌入的、持久的、快速的键值(KV)数据库，使用纯Go编写，具有可预测的读写性能、低延迟和高吞吐量，这得益于Bitcask的磁盘布局(LSM+WAL)。
* [buntdb](https://github.com/tidwall/buntdb)  基于内存的K/V，快速，可嵌入的数据库，可自定义索引和空间支持。
* [cache](https://github.com/akyoto/cache)  基于内存的 K/V 存储:带生命周期的值存储，0个依赖项，<100 LoC, 100%覆盖率。
* [cache2go](https://github.com/muesli/cache2go)  基于内存的 K/V 缓存，支持超时的自动失效。
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache)  BigCache 支持集群和独立且生命周期存储项。
* [cockroach](https://github.com/cockroachdb/cockroach)  可伸缩、区域备份、事务性数据存储。
* [couchcache](https://github.com/codingsince1985/couchcache)  由 Couchbase服务 支持的RESTful缓存微服务。
* [goleveldb](https://github.com/syndtr/goleveldb)  在Go中实现[LevelDB](https://github.com/google/leveldb) key/value数据库。
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL)  区块链领域的一个SQL数据库。
* [Databunker](https://github.com/paranoidguy/databunker)  个人身份信息(PII)存储服务符合GDPR和CCPA。
* [dgraph](https://github.com/dgraph-io/dgraph)  可伸缩、分布式、低延迟、高吞吐量的图形数据库。
* [diskv](https://github.com/peterbourgon/diskv)  支持磁盘备份的可持久化 K/V 存储。
* [eliasdb](https://github.com/krotik/eliasdb)  无其他依赖项，支持REST API，短语搜索和sql类似的查询语言的事务图数据库。
* [fastcache](https://github.com/VictoriaMetrics/fastcache)  基于内存的快速线程安全的缓存，可缓存大量的条目。最大限度地减少GC开销。
* [GCache](https://github.com/bluele/gcache)  支持过期缓存、LFU、LRU和ARC的缓存库。
* [go-cache](https://github.com/pmylund/go-cache)  基于内存的 K/V 存储/缓存 : (类似于Memcached)，适用于单机应用程序。
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics)  开源，快速，可伸缩的时间序列数据库。支持PromQL。

*数据库迁移。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [avro](https://github.com/khezen/avro)  发现SQL schemas并将其转换为AVRO schemas。
* [darwin](https://github.com/GuiaBolso/darwin)  用于数据库 schema 升级的库。
* [go-fixtures](https://github.com/RichardKnop/go-fixtures)  类似 Django fixture，用于 Go 建立内置数据库/sql库。
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations)  用Go -pg/pg编写的迁移包。
* [gondolier](https://github.com/emvi/gondolier)  使用结构修饰的数据库迁移库。
* [goose](https://github.com/steinbacher/goose)  数据库迁移工具。您可以通过创建增量SQL或Go脚本来管理数据库的升级。
* [gormigrate](https://github.com/go-gormigrate/gormigrate)  面向Gorm ORM的数据库 schema 迁移辅助程序。
* [migrate](https://github.com/golang-migrate/migrate)  基于CLI的数据库迁移库。
* [migrator](https://github.com/lopezator/migrator)  非常简单的 Go 数据库迁移库。
* [pravasan](https://github.com/pravasan/pravasan)  简易的迁移工具-目前只支持MySQL，但计划很快支持Postgres, SQLite, MongoDB等。
* [schema](https://github.com/adlio/schema)  库，用于将数据库/sql兼容数据库的模式迁移嵌入到Go二进制文件中。
* [skeema](https://github.com/skeema/skeema)  用于MySQL的纯sql模式管理系统，支持分片和外部在线模式更改工具。
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [sql-migrate](https://github.com/rubenv/sql-migrate)  数据库迁移工具。允许使用go-bindata将迁移嵌入到应用程序中。

*数据库工具。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [bucket](https://github.com/PumpkinSeed/bucket)  优化的数据结构框架的Couchbase专门针对一个桶的使用。
* [chproxy](https://github.com/Vertamedia/chproxy)  ClickHouse数据库的HTTP代理。
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk)  收集小的 insterts 并向 ClickHouse 服务器发送大请求。
* [datagen](https://github.com/codingconcepts/datagen)  一个快速的数据生成器，支持多表感知和多行DML。
* [dbbench](https://github.com/sj14/dbbench)  数据库基准测试工具，支持多个数据库和脚本。
* [go-mysql](https://github.com/siddontang/go-mysql)   Go 工具集，用于处理MySQL协议和复制。
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch)  自动将MySQL数据同步到Elasticsearch中。
* [kingshard](https://github.com/flike/kingshard)  kingshard 是基于 Golang 的MySQL高性能代理。
* [myreplication](https://github.com/2tvenom/myreplication)  MySql二进制日志复制监听器。支持基于语句和行的复制。
* [octillery](https://github.com/knocknote/octillery)  用于数据库分表(支持每个ORM或原生SQL)。
* [orchestrator](https://github.com/github/orchestrator)  MySQL复制拓扑管理器和可视化工具。
* [pgweb](https://github.com/sosedoff/pgweb)  基于web的PostgreSQL数据库浏览器。
* [prep](https://github.com/hexdigest/prep)  在不更改代码的情况下使用准备好的SQL语句。
* [pREST](https://github.com/nuveo/prest)  基于PostgreSQL database的RESTful API服务。
* [rwdb](https://github.com/andizzle/rwdb)  rwdb为多个数据库服务器的设置提供读取副本功能。
* [vitess](https://github.com/youtube/vitess)  vitess提供了可以为大规模web服务扩展MySQL数据库提供便利的服务和工具。

*SQL查询生成器，用于构建和使用SQL的库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [dbq](https://github.com/rocketlaunchr/dbq)  Zero boilerplate database operations for Go
* [Dotsql](https://github.com/gchaincl/dotsql)  Go library帮助您将sql文件保存在一个地方，并轻松地使用它们。
* [gendry](https://github.com/didi/gendry)  非入侵的SQL构建器和强大的数据绑定器。
* [godbal](https://github.com/xujiajun/godbal)  数据库抽象层(dbal)。支持SQL builder，轻松获取结果。
* [goqu](https://github.com/doug-martin/goqu)  常用的SQL生成器和查询库。
* [igor](https://github.com/galeone/igor)  PostgreSQL的抽象层，支持高级功能和类似gorm的语法。
* [jet](https://github.com/go-jet/jet)  用于在Go中编写类型安全的SQL查询的框架，能够轻松地将数据库查询结果转换为所需的任意对象结构。
* [mpath](https://github.com/spacetab-io/mpath-go)  MPTT(修改前序树遍历)包的SQL记录-物化路径实现。
* [ormlite](https://github.com/pupizoid/ormlite)  包含一些类似orm的特性和sqlite数据库的辅助程序的轻量级包
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)  Powerful data retrieval methods as well as DB-agnostic query building capabilities.
* [qry](https://github.com/HnH/qry)  该工具使用原始SQL查询从文件生成常量。
* [sqlf](https://github.com/leporo/sqlf)  快速SQL查询生成器。
* [sqrl](https://github.com/elgris/sqrl)  SQL查询生成器，从Squirrel fork而来，并再此基础上对性能做了优化。
* [Squalus](https://gitlab.com/qosenergy/squalus)  Go SQL中间层，能使得执行查询更加容易。
* [Squirrel](https://github.com/Masterminds/squirrel)  帮助您构建SQL查询的Go库。
* [xo](https://github.com/knq/xo)  基于现有的schema定义和自定义查询生成 Go 代码，基于支持PostgreSQL、MySQL、SQLite、Oracle和Microsoft SQL Server。

## 数据库驱动程序

*用于连接和操作数据库的库。 (翻译出错了? 试试 [英文版](README_EN.md#database-drivers) 吧~)*

* Relational Databases
    * [avatica](https://github.com/apache/calcite-avatica-go)  Apache Avatica/Phoenix SQL驱动程序。
    * [bgc](https://github.com/viant/bgc)  BigQuery 的数据存储连接。
    * [firebirdsql](https://github.com/nakagami/firebirdsql)  Firebird RDBMS SQL驱动程序。
    * [go-adodb](https://github.com/mattn/go-adodb)  Microsoft ActiveX对象数据库驱动程序。
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb)  微软MSSQL驱动程序。
    * [go-oci8](https://github.com/mattn/go-oci8)  Oracle 驱动程序。
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)  MySQL驱动程序。
    * [go-sqlite3](https://github.com/mattn/go-sqlite3)  SQLite3驱动程序。
    * [gofreetds](https://github.com/minus5/gofreetds)  基于[FreeTDS](http://www.freetds.org)封装的微软MSSQL Go 驱动。
    * [goracle](https://github.com/go-goracle/goracle)  基于 ODPI-C 的 Oracle 驱动程序
    * [pgx](https://github.com/jackc/pgx)  PostgreSQL驱动，支持比现有database/sql更多的特性。
    * [pq](https://github.com/lib/pq)  纯 Go 的Postgres驱动。

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go)  Aerospike 客户端。
    * [gorethink](https://github.com/dancannon/gorethink)  RethinkDB 驱动。
    * [redis](https://github.com/go-redis/redis)  基于 Go 的 Redis 客户端。
    * [redigo](https://github.com/gomodule/redigo)  Redigo 是基于 Go 的Redis 客户端。
    * [redeo](https://github.com/bsm/redeo)  与 redis 协议兼容的 TCP 服务器/服务。
    * [neoism](https://github.com/jmcvetta/neoism)  Golang 的 Neo4j 客户端。
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO)  Neo4j REST 客户端。
    * [neo4j](https://github.com/cihangir/neo4j)  Neo4j Rest API实现。
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)  官方的 MongoDB 驱动。
    * [mgo](https://github.com/globalsign/mgo)  (已停止维护) MongoDB驱动。
    * [mgm](https://github.com/kamva/mgm)  基于MongoDB模型的ODM for Go(基于官方MongoDB驱动)。
    * [goriak](https://github.com/zegl/goriak)   Riak KV 驱动。
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache客户端库。
    * [arangolite](https://github.com/solher/arangolite)  轻量级的 ArangoDB 驱动。
    * [godscache](https://github.com/defcronyke/godscache)  谷歌云平台Go Datastore包的封装，它采用了memcached添加缓存。
    * [godis](https://github.com/piaohao/godis)  由GoLang实现的redis客户端，灵感来自jedis。
    * [gocql](http://gocql.github.io)  Apache Cassandra 的 Go 驱动。
    * [gocb](https://github.com/couchbase/gocb)  官方Couchbase Go SDK。
    * [go-rejson](https://github.com/nitishm/go-rejson)  实现了基于 Redigo 客户端的redislabs' ReJSON 模块。可简单地将结构体存储为JSON对象并对其进行操作。
    * [go-pilosa](https://github.com/pilosa/go-pilosa)   Pilosa客户端。
    * [go-couchbase](https://github.com/couchbase/go-couchbase)  Couchbase客户端。
    * [forestdb](https://github.com/couchbase/goforestdb)  基于 Go 的 ForestDB 接口实现。
    * [dynago](https://github.com/underarmour/dynago)  满足 principle of least surprise 的 DynamoDB 客户端。
    * [asc](https://github.com/viant/asc)  Aerospike 的数据存储连接器。
    * [xredis](https://github.com/shomali11/xredis)  类型安全，可定制，清晰和易用的Redis客户端。

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve)  基于 Go 的现代文本索引库。
    * [elastic](https://github.com/olivere/elastic)  Elasticsearch 客户端。
    * [elasticsql](https://github.com/cch123/elasticsql)  将 SQL 转换为 elasticsearch dsl。
    * [elastigo](https://github.com/mattbaird/elastigo)  Elasticsearch 客户端。
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch)  官方 Elasticsearch 客户端。
    * [goes](https://github.com/OwnLocal/goes)  实现了与 Elasticsearch 交互的库。
    * [riot](https://github.com/go-ego/riot)  基于 Go 的 开源、分布式、简单高效的搜索引擎。
    * [skizze](https://github.com/seiflotfy/skizze)  面向概率数据结构的服务和存储。

* Multiple Backends.
    * [cachego](https://github.com/fabiorphp/cachego)  基于多个驱动程序的缓存组件。
    * [cayley](https://github.com/google/cayley)  图形数据库，支持多个后端。
    * [dsc](https://github.com/viant/dsc)  面向 SQL、NoSQL、结构化文件的数据存储连接。
    * [gokv](https://github.com/philippgille/gokv)  可扩展的简单的 K/V 存储(Redis、Consul、etcd、bbolt、BadgerDB、LevelDB、Memcached、DynamoDB、S3、PostgreSQL、MongoDB、CockroachDB等等)。

## 日期和时间

*用于处理日期和时间的库。 (翻译出错了? 试试 [英文版](README_EN.md#date-and-time) 吧~)*

* [carbon](https://github.com/uniplaces/carbon)  简单的时间扩展，包含了许多使用方法，从 PHP Carbon 库移植的。
* [iso8601](https://github.com/relvacode/iso8601)  不用正则表达式有效解析 ISO8601 日期时间。
* [timeutil](https://github.com/leekchan/timeutil)  面向 Golang 的时间库，集成了很多有用的扩展(Timedelta, Strftime, ...)。
* [timespan](https://github.com/SaidinWoT/timespan)  用于处理时间间隔相关的库，可定义开始时间和持续时间。
* [strftime](https://github.com/awoodbeck/strftime)  C99-compatible strftime formatter。
* [NullTime](https://github.com/kirillDanshin/nulltime)  可以允许 time.Time 为null。
* [now](https://github.com/jinzhu/now)  now 是时间有关的工具类。
* [kair](https://github.com/GuilhermeCaruso/kair)  用于处理日期和时间的 golang 库。
* [go-week](https://github.com/stoewer/go-week)  一个有效的软件包，以配合ISO8601周日期。
* [cronrange](https://github.com/1set/cronrange)  解析cron风格的时间范围表达式，检查给定的时间是否在任何范围内。
* [go-sunrise](https://github.com/nathan-osman/go-sunrise)  计算指定位置的日出和日落时间。
* [go-str2duration](https://github.com/xhit/go-str2duration)  将字符串转换为持续时间。支持的时间。持续时间返回字符串等。
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar)  太阳历实现。
* [feiertage](https://github.com/wlbr/feiertage)  用于计算德国公共假期的函数，比如复活节、感恩节等
* [durafmt](https://github.com/hako/durafmt)  轻量级、可让time.Duration更加易读的库。
* [dateparse](https://github.com/araddon/dateparse)  可以解析很多格式不固定的日期字符串。
* [date](https://github.com/rickb777/date)  增加处理日期、日期范围、时间跨度、时间段和time-of-day。
* [tuesday](https://github.com/osteele/tuesday)  Ruby-compatible Strftime function。

## 分布式系统

*协助构建分布式系统的包。 (翻译出错了? 试试 [英文版](README_EN.md#distributed-systems) 吧~)*

* [celeriac](https://github.com/svcavallar/celeriac.v1)  用于对 Celery worker、任务、事件进行交互和监控的库。
* [raft](https://github.com/hashicorp/raft)  Raft consensus协议的实现。 by HashiCorp。
* [jsonrpc](https://github.com/ybbus/jsonrpc)  JSON-RPC 2.0 HTTP客户端。
* [KrakenD](https://github.com/devopsfaith/krakend)  具有中间件的高性能API网关框架。
* [liftbridge](https://github.com/liftbridge-io/liftbridge)  用于nat的轻量级、容错的消息流。
* [micro](https://github.com/micro/micro)  可插拔的微服务 toolkit 和分布式系统平台。
* [NATS](https://github.com/nats-io/gnatsd)  轻量级、高性能消息传递系统，可用于微服务、物联网(IoT)和云。
* [outboxer](https://github.com/italolelis/outboxer)  实现了 outbox pattern Go 库。
* [pglock](https://cirello.io/pglock)  postgresql 的分布式锁定实现。
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Raft consensus协议的实现。 by CoreOS。
* [hprose](https://github.com/hprose/hprose-golang)  支持25+种语言RPC库。
* [rain](https://github.com/cenkalti/rain)  BitTorrent客户端和库。
* [redis-lock](https://github.com/bsm/redislock)  基于redis的分布式锁简易实现。
* [resgate](https://resgate.io/)  用于构建REST、实时和RPC API的实时API网关，其中所有客户端都是无缝同步的。
* [ringpop-go](https://github.com/uber/ringpop-go)  可伸缩的，容错、应用分层的的Go应用程序。
* [rpcx](https://github.com/smallnest/rpcx)  分布式可插拔的RPC服务框架，如阿里巴巴Dubbo。
* [sleuth](https://github.com/ursiform/sleuth)  用于HTTP服务之间进行无中心p2p自动发现和RPC通信的库(使用[ZeroMQ](https://github.com/zeromq/libzmq))。
* [tendermint](https://github.com/tendermint/tendermint)  一个高性能中间件，可将任何语言的状态机转换为 Byzantine Fault 状态机。使用 Tendermint 一致性及区块链协议。
* [jsonrpc](https://github.com/osamingo/jsonrpc)  jsonrpc 包，实现了 JSON-RPC 2.0。
* [grpc-go](https://github.com/grpc/grpc-go)  gRPC的Go语言实现。
* [consistent](https://github.com/buraksezer/consistent)  Consistent hashing with bounded loads。
* [dynatomic](https://github.com/tylfin/dynatomic)  基于 DynamoDB 的 原子计数器的库。
* [dht](https://github.com/anacrolix/dht)  BitTorrent Kademlia DHT的实现。
* [digota](https://github.com/digota/digota)  基于 grpc 的电子商务微服务。
* [dot](https://github.com/dotchain/dot/)  基于 transformation/OT 的分布式同步。
* [doublejump](https://github.com/edwingeng/doublejump)  实现了谷歌的jump consistent hash。
* [dragonboat](https://github.com/lni/dragonboat)  一个功能齐全，高性能的库集。
* [drmaa](https://github.com/dgruber/drmaa)  符合 DRMAA 标准的集群调度程序作业提交库。
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed 分布式锁定实现。
* [emitter-io](https://github.com/emitter-io/emitter)  高性能、分布式、安全和低延迟的发布-订阅平台，使用MQTT、Websockets和love构建。
* [gorpc](https://github.com/valyala/gorpc)  简单、快速和可伸缩的RPC库。
* [flowgraph](https://github.com/vectaport/flowgraph)  flow-based programming package。
* [gleam](https://github.com/chrislusf/gleam)  使用纯Go和Luajit编写的快速、可伸缩的分布式map/reduce系统，结合了Go的高并发性和Luajit的高性能，可以独立运行或分布式运行。
* [glow](https://github.com/chrislusf/glow)  全部用 Go 实现，易用、可伸缩，可用于分布式大数据处理，Map-Reduce, DAG执行。
* [go-health](https://github.com/InVisionApp/go-health)  用于在服务中启用异步依赖项健康检查的库。
* [go-jump](https://github.com/dgryski/go-jump)  提供了谷歌的 “Jump” 一致哈希函数接口。
* [go-kit](https://github.com/go-kit/kit)  支持服务发现、负载平衡、插件式传输、请求跟踪等功能的Microservice toolkit。
* [go-pdu](https://github.com/pdupub/go-pdu)  一个去中心化的基于身份的社交网络。
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit)  为支持为golang服务定义异步服务健康检查而构建的库。
* [torrent](https://github.com/anacrolix/torrent)  BitTorrent 客户端。

## 动态域名

*更新动态DNS记录的工具。 (翻译出错了? 试试 [英文版](README_EN.md#dynamic-dns) 吧~)*

* [DDNS](https://github.com/skibish/ddns)  个人 DDNS 客户端。
* [dyndns](https://gitlab.com/alcastle/dyndns)  后台去处理定期和自动检查您的IP地址，并作出更新(一个或多个)动态DNS记录，为谷歌域，每当您的地址变化。
* [GoDNS](https://github.com/timothyye/godns)  一个动态DNS客户端工具，支持DNSPod & HE.net。

## 电子邮件

*实现了电子邮件创建和发送。 (翻译出错了? 试试 [英文版](README_EN.md#email) 吧~)*

* [chasquid](https://blitiri.com.ar/p/chasquid)  用Go编写的SMTP服务器。
* [douceur](https://github.com/aymerick/douceur)  在HTML邮件中支持CSS内联。
* [email](https://github.com/jordan-wright/email)  一个强大和灵活的电子邮件库。
* [go-dkim](https://github.com/toorop/go-dkim)  DKIM库，用于签署 & 验证电子邮件。
* [go-imap](https://github.com/emersion/go-imap)  用于客户端和服务器的IMAP库。
* [go-message](https://github.com/emersion/go-message)  用于Internet消息格式化和邮件消息的流媒体库。
* [go-premailer](https://github.com/vanng822/go-premailer)  在HTML邮件中支持CSS内联。
* [go-simple-mail](https://github.com/xhit/go-simple-mail)  非常简单的包发送电子邮件与SMTP保持活动和两个超时:连接和发送。
* [Hectane](https://github.com/hectane/hectane)  轻量级的SMTP客户机，提供了HTTP API。
* [hermes](https://github.com/matcornic/hermes)  可生成干净的、响应式的HTML电子邮件。
* [mailchain](https://github.com/mailchain/mailchain)  发送加密的电子邮件到区块链地址写在Go。
* [mailgun-go](https://github.com/mailgun/mailgun-go)  使用Mailgun API去库发送邮件。
* [MailHog](https://github.com/mailhog/MailHog)  电子邮件和SMTP测试工具，对外提供了 web 和 API 接口。
* [SendGrid](https://github.com/sendgrid/sendgrid-go)  SendGrid 的 Go语言库，用于发送电子邮件。
* [smtp](https://github.com/mailhog/smtp)  SMTP服务器协议状态机。

## 可嵌入的脚本语言

*在go代码中嵌入其他语言。 (翻译出错了? 试试 [英文版](README_EN.md#embeddable-scripting-languages) 吧~)*

* [anko](https://github.com/mattn/anko)  用Go编写的解释器。
* [go-python](https://github.com/sbinet/go-python)  CPython C-API 的 Go 接口。
* [purl](https://github.com/ian-kent/purl)  嵌入 Go 的 Perl 5.18.2。
* [otto](https://github.com/robertkrimen/otto)  用 Go 编写的 JavaScript 解释器。
* [ngaro](https://github.com/db47h/ngaro)  嵌入式 Ngaro VM实现，支持在 Retro 中运行脚本。
* [gval](https://github.com/PaesslerAG/gval)  一种用Go编写的高度可定制的表达式语言。
* [gopher-lua](https://github.com/yuin/gopher-lua)  用 Go 实现的 Lua 5.1 虚拟机和编译器。
* [golua](https://github.com/aarzilli/golua)  Lua C 的 Go 接口。
* [go-php](https://github.com/deuill/go-php)  PHP 的 Go 接口。
* [binder](https://github.com/alexeyco/binder)  Lua接口，基于[gopher-lua](https://github.com/yuin/gopher-lua)。
* [go-lua](https://github.com/Shopify/go-lua)  用 Go 实现的 Lua 5.2 VM接口。
* [go-duktape](https://github.com/olebedev/go-duktape)  支持 Duktape JavaScript 引擎。
* [gisp](https://github.com/jcla1/gisp)  LISP 的 Go 接口。
* [gentee](https://github.com/gentee/gentee)  嵌入式脚本编程语言。
* [expr](https://github.com/antonmedv/expr)  一个可以计算表达式的引擎。
* [cel-go](https://github.com/google/cel-go)  快速，可移植，非图灵完整的表达评估与渐进分型。
* [tengo](https://github.com/d5/tengo)  字节码编译的脚本语言。

## 错误处理

*处理错误的库。 (翻译出错了? 试试 [英文版](README_EN.md#error-handling) 吧~)*

* [emperror](https://github.com/emperror/emperror)  用于Go库和应用程序的错误处理工具和最佳实践。
* [eris](https://github.com/rotisserie/eris)  在Go中处理、跟踪和记录错误的更好方法。兼容标准错误库和github.com/pkg/errors。
* [errlog](https://github.com/snwfdhmp/errlog)  用于定位抛出错误的源代码(以及一些其他快速调试特性)。可插入到任何 logger 的位置。
* [errors](https://github.com/emperror/errors)  替换标准库错误包和github.com/pkg/errors。提供各种错误处理原语。
* [errors](https://github.com/pkg/errors)  可让你很简单的进行错误处理。
* [errors](https://github.com/neuronlabs/errors)  使用分类原语进行简单的golang错误处理。
* [errors](https://github.com/PumpkinSeed/errors)  最简单的错误包装器，具有出色的性能和最小的内存开销。
* [errorx](https://github.com/joomcode/errorx)  一个功能丰富的错误包，可进行堆栈跟踪、组装异常信息以及其他。
* [Falcon](https://github.com/SonicRoshan/falcon)  一个简单但功能强大的错误处理包。
* [go-multierror](https://github.com/hashicorp/go-multierror)  可将一系列的错误作为一个整体来显示。
* [tracerr](https://github.com/ztrue/tracerr)  可展示错误的堆栈跟踪信息和源码片段。
* [werr](https://github.com/txgruppi/werr)  对错误异常进行了捕获封装，封装信息包含了调用它的文件、行和堆栈。

## 文件

*处理文件和文件系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#files) 吧~)*

* [afero](https://github.com/spf13/afero)  文件系统的抽象系统。
* [gut/yos](https://github.com/1set/gut)  简单和可靠的包文件操作，如复制/移动/diff/列表的文件，目录和符号链接。
* [tarfs](https://github.com/posener/tarfs)  tar文件的实现[ FileSystem 接口](https://godoc.org/github.com/kr/fs#FileSystem)。
* [stl](https://gitlab.com/russoj88/stl)  采用并行读取算法的进行读取和写入STL(立体光刻)文件的模块。
* [skywalker](https://github.com/dixonwille/skywalker)  可以轻松地并发地遍历文件系统。
* [pdfcpu](https://github.com/pdfcpu/pdfcpu)  PDF处理器。
* [parquet](https://github.com/parsyl/parquet)  读写[parquet](https://parquet.apache.org)文件。
* [opc](https://github.com/qmuntal/opc)  加载Open Packaging Conventions (OPC)文件。
* [notify](https://github.com/rjeczalik/notify)  文件系统事件通知库，具有类似于os/signal的简单API，。
* [go-gtfs](https://github.com/artonge/go-gtfs)  加载gtfs文件。
* [afs](https://github.com/viant/afs)  用于Go的抽象文件存储(mem、scp、zip、tar、cloud: s3、gs)。
* [go-exiftool](https://github.com/barasher/go-exiftool)  ExifTool 的 Go 实现，这个著名的库用于从文件(图片、PDF、office，…)中提取尽可能多的元数据(EXIF、IPTC，…)。
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy)  拷贝文件。
* [go-csv-tag](https://github.com/artonge/go-csv-tag)  使用 tag 加载 csv 文件。
* [flop](https://github.com/homedepot/flop)  文件操作库，是[GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invoc.html)的镜像。
* [copy](https://github.com/otiai10/copy)  递归地复制目录。
* [checksum](https://github.com/codingsince1985/checksum)  生成大型文件的消息摘要(如 MD5 和 SHA256)。
* [bigfile](https://github.com/bigfile/bigfile)  一个文件传输系统，支持管理文件与http api, rpc调用和ftp客户端。
* [vfs](https://github.com/C2FO/vfs)  一组可插拔的、可扩展的和自定义的文件系统功能，用于跨越许多文件系统类型，如os、S3和GCS。

## 金融

*会计和财务软件包。 (翻译出错了? 试试 [英文版](README_EN.md#financial) 吧~)*

* [accounting](https://github.com/leekchan/accounting)  货币和货币格式。
* [currency](https://github.com/bnkamalesh/currency)  高性能、准确的货币计算软件包。
* [decimal](https://github.com/shopspring/decimal)  任意精度定点的十进制数。
* [go-finance](https://github.com/FlashBoys/go-finance)  综合金融市场数据。
* [go-finance](https://github.com/alpeb/go-finance)  用于货币时间价值(年金)、现金流、利率转换、债券和折旧计算的金融函数库。
* [go-finance](https://github.com/pieterclaerhout/go-finance)  模块获取汇率，通过VIES检查增值税号码，检查IBAN银行账号。
* [go-finnhub](https://github.com/m1/go-finnhub)  来自finnhu .io的股票市场、外汇和密码数据客户。从60多个证券交易所、10个外汇经纪人和15多个密码交易所获取实时金融市场数据。
* [go-money](https://github.com/rhymond/go-money)  Fowler 货币模式的实现。
* [ofxgo](https://github.com/aclindsa/ofxgo)  查询 OFX 服务器和/或解析响应。
* [orderbook](https://github.com/i25959341/orderbook)  限购单匹配引擎。
* [techan](https://github.com/sdcoffey/techan)  拥有先进的市场分析和交易策略的技术分析库。
* [transaction](https://github.com/claygod/transaction)  嵌入式事务数据库，可多线程模式运行。
* [vat](https://github.com/dannyvankooten/vat)  增值税编号验证 & 欧盟增值税税率。

## 表单

*用于处理表单的库。 (翻译出错了? 试试 [英文版](README_EN.md#forms) 吧~)*

* [bind](https://github.com/robfig/bind)  将表单数据与任意 Go 变量进行绑定。
* [binding](https://github.com/mholt/binding)  将来自 net/HTTP 请求的表单、JSON 数据绑定到结构体。
* [conform](https://github.com/leebenson/conform)  控制用户输入。基于struct tags可对数据进行修剪、清理和擦除。
* [form](https://github.com/go-playground/form)   将 url 中的数据解析到 Go 变量中，以及将 Go 语言变量编码进 url。支持 Dual Array 及 Full map。
* [formam](https://github.com/monoculum/formam)  将表单的值解码为 struct。
* [forms](https://github.com/albrow/forms)  与框架无关的库，用于解析和验证支持多部分表单和文件的表单/JSON数据。
* [gorilla/csrf](https://github.com/gorilla/csrf)  用于Go web应用程序和服务的CSRF保护。
* [nosurf](https://github.com/justinas/nosurf)  CSRF保护中间件。
* [queryparam](https://github.com/tomwright/queryparam)  解码的url。值转换为标准或自定义类型的可用结构值。

## 方法

*在Go中支持函数式编程的包。 (翻译出错了? 试试 [英文版](README_EN.md#functional) 吧~)*

* [fpGo](https://github.com/TeaEntityLab/fpGo)  提供函数式编程功能。
* [fuego](https://github.com/seborama/fuego)  Functional Experiment in Go。
* [go-underscore](https://github.com/tobyhede/go-underscore)  常用辅助方法集合。

## 游戏开发

*很棒的游戏开发库。 (翻译出错了? 试试 [英文版](README_EN.md#game-development) 吧~)*

* [Azul3D](https://github.com/azul3d/engine)  用Go编写的 3D 游戏引擎。
* [gonet](https://github.com/xtaci/gonet)  实现了游戏服务器骨架。
* [raylib-go](https://github.com/gen2brain/raylib-go)  实现了 [raylib](http://www.raylib.com/)，一个简单易用的库，用于学习视频游戏编程。
* [Pixel](https://github.com/faiface/pixel)  手工制作的 2D 游戏库。
* [Pitaya](https://github.com/topfreegames/pitaya)  可伸缩的游戏服务器框架，支持集群和客户端库的iOS, Android, Unity。
* [Oak](https://github.com/oakmound/oak)  纯 Go 实现的游戏引擎。
* [nano](https://github.com/lonng/nano)  轻量级、方便、高性能的基于golang的游戏服务器框架。
* [Leaf](https://github.com/name5566/leaf)  轻量级游戏服务器框架。
* [goworld](https://github.com/xiaonanln/goworld)  可伸缩的游戏服务器引擎，具有 space-entity 框架和 hot-swapping 功能。
* [go3d](https://github.com/ungerik/go3d)  以性能为主的2D/3D数学相关包。
* [Ebiten](https://github.com/hajimehoshi/ebiten)  很简单的 2D 游戏库。
* [go-sdl2](https://github.com/veandco/go-sdl2)  实现了[Simple DirectMedia Layer](https://www.libsdl.org/)。
* [go-collada](https://github.com/GlenKelley/go-collada)  处理Collada文件。
* [go-astar](https://github.com/beefsack/go-astar)  实现了A\*路径查找算法。
* [glop](https://github.com/runningwild/glop)  Glop (Game Library Of Power) 是一个相当简单的跨平台游戏库。
* [GarageEngine](https://github.com/vova616/GarageEngine)  用 OpenGL 编写的 2D 游戏引擎。
* [g3n](https://github.com/g3n/engine)   3D游戏引擎。
* [engo](https://github.com/EngoEngine/engo)  开源 2D 游戏引擎。它遵循 Entity-Component-System 范式。
* [termloop](https://github.com/JoelOtter/termloop)  基于终端的 Go 游戏引擎，建立在 Termbox 之上。

## 代码生成与泛型

*增强语言的工具，例如通过代码生成支持泛型。 (翻译出错了? 试试 [英文版](README_EN.md#generation-and-generics) 吧~)*

* [efaceconv](https://github.com/t0pep0/efaceconv)  代码生成工具，可以不通过内存分配就可以高效的将interface{}转换为不可变类型，。
* [gen](https://github.com/clipperhouse/gen)  用于生成泛型等类似方法的功能代码生成工具。
* [generis](https://github.com/senselogic/GENERIS)  提供泛型、free-form 宏、条件编译和HTML模板的代码生成工具。
* [go-enum](https://github.com/abice/go-enum)  从代码注释中生成枚举。
* [go-linq](https://github.com/ahmetalpbalkan/go-linq)  提供类似 .NET LINQ 的查询方法。
* [go-xray](https://github.com/pieterclaerhout/go-xray)  帮助更容易地使用反射。
* [goderive](https://github.com/awalterschulze/goderive)  从输入类型来派生函数。
* [gotype](https://github.com/wzshiming/gotype)  Golang 源码解析，用法类似reflect(反射)。
* [GoWrap](https://github.com/hexdigest/gowrap)  使用简单模板为 Go 接口生成装饰器。
* [interfaces](https://github.com/rjeczalik/interfaces)  用于生成接口定义的命令行工具。
* [jennifer](https://github.com/dave/jennifer)  不使用模板生成任意 Go 代码。
* [pkgreflect](https://github.com/ungerik/pkgreflect)  用于包作用域反射的 Go 预处理器。
* [typeregistry](https://github.com/xiaoxin01/typeregistry)  动态创建类型的库。

## 地理

*地理工具和服务器 (翻译出错了? 试试 [英文版](README_EN.md#geographic) 吧~)*

* [geocache](https://github.com/melihmucuk/geocache)  基于内存缓存的的地理位置的应用程序。
* [geoserver](https://github.com/hishamkaram/geoserver)  基于geoserver REST API的 geoserver 实例。
* [gismanager](https://github.com/hishamkaram/gismanager)  将你的 GIS 数据(矢量数据)发布到 PostGIS 和 Geoserver。
* [mbtileserver](https://github.com/consbio/mbtileserver)  一个简单的基于go的服务器，用于存储mbtiles格式的地图块。
* [osm](https://github.com/paulmach/osm)  用于读取、写入和处理 OpenStreetMap 数据和 APIs。
* [pbf](https://github.com/maguro/pbf)  基于Golang 的 OpenStreetMap PBF 编码器/解码器。
* [S2 geometry](https://github.com/golang/geo)  S2 geometry 库。
* [Tile38](https://github.com/tidwall/tile38)  具有空间索引和实时地理定位功能的地理定位数据库。
* [WGS84](https://github.com/wroge/wgs84)  坐标转换和转换库(ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM)。

## Go 编译器

*可将 Go 转换为其他语言的编译工具。 (翻译出错了? 试试 [英文版](README_EN.md#go-compilers) 吧~)*

* [c4go](https://github.com/Konstantin8105/c4go)  将 C 代码转换为 Go 代码。
* [f4go](https://github.com/Konstantin8105/f4go)  将 FORTRAN 77 转换为 Go代码。
* [gopherjs](https://github.com/gopherjs/gopherjs)  将 Go 编译成 JavaScript。
* [llgo](https://github.com/go-llvm/llgo)  基于 llvm 的编译器。
* [tardisgo](https://github.com/tardisgo/tardisgo)  Golang 转换为 Haxe，再转换为 CPP/CSharp/Java/JavaScript 的编译器.

## Goroutines

*管理和处理 Goroutines 的工具。 (翻译出错了? 试试 [英文版](README_EN.md#goroutines) 吧~)*

* [ants](https://github.com/panjf2000/ants)  一个高性能和低成本的goroutine池在围棋。
* [grpool](https://github.com/ivpusic/grpool)  轻量级协程池。
* [worker-pool](https://github.com/vardius/worker-pool)  一个简单的 Go 异步工作池。
* [tunny](https://github.com/Jeffail/tunny)  golang 的协程池。
* [threadpool](https://github.com/shettyh/threadpool)  Golang 的 threadpool 实现。
* [stl](https://github.com/ssgreg/stl)  基于软件事务内存(STM)并发控制机制的软件事务锁。
* [semaphore](https://github.com/marusama/semaphore)  基于 CAS 的可快速调整的信号量实现(比基于通道的信号量实现更快)。
* [semaphore](https://github.com/kamilsk/semaphore)  信号量模式实现，可根据通道和上下文进行具备超时功能的锁定/解锁操作。
* [routine](https://github.com/x-mod/routine)  go routine control with context, support: Main, Go, Pool and some useful Executors.
* [queue](https://github.com/AnikHasibul/queue)  提供类似队列组可访问性 sync.WaitGroup 的功能。帮助你节流和限制协程、等待所有协程结束以及更多功能。
* [pool](https://github.com/go-playground/pool)  有限消费者协程或无限协程池，可用于更加简单的处理和取消协程
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn)  并行运行函数。
* [oversight](https://cirello.io/oversight)  完整的实现了Erlang supervision trees。
* [kyoo](https://github.com/dirkaholic/kyoo)  提供无限的作业队列和并发工作池。
* [Hunch](https://github.com/AaronJan/Hunch)  Hunch 提供了诸如 All、First、Retry、Waterfall 等功能，这使得异步流控制更加直观。
* [gpool](https://github.com/Sherifabdlnaby/gpool)  manages a resizeable pool of context-aware goroutines to bound concurrency
* [artifex](https://github.com/borderstech/artifex)  简单的内存作业队列。
* [gowp](https://github.com/xxjwxc/gowp)  gowp是高并发性限制异步工作池。
* [goworker](https://github.com/benmanns/goworker)  基于 go 的后台 worker。
* [GoSlaves](https://github.com/themester/GoSlaves)  简单异步的协程池。
* [gollback](https://github.com/vardius/gollback)  异步简单的函数实用程序，用于管理闭包和回调的执行。
* [gohive](https://github.com/loveleshsharma/gohive)  一个高性能和易于使用的Goroutine池去。
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup)  像“同步。有错误处理和并发控制。
* [go-trylock](https://github.com/subchen/go-trylock)  支持 read-write 锁。
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools)  轻量级的协程池库，可以通过简单的API来管理。
* [go-flow](https://github.com/kamildrazkiewicz/go-flow)  控制 goroutines 的执行顺序。
* [go-floc](https://github.com/workanator/go-floc)  轻松编排 goroutines。
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier)  基于 Go 的 CyclicBarrier 实现。
* [conexec](https://github.com/ITcathyh/conexec)  一个并发工具包，帮助以高效和安全的方式并发执行函数。它支持指定总的超时来避免阻塞，并使用goroutine池来提高效率。
* [breaker](https://github.com/kamilsk/breaker)  灵活的机制，可以使执行流可中断。
* [async](https://github.com/studiosol/async)  一种异步执行函数的安全方法，在出现 panic 时恢复它们。
* [workerpool](https://github.com/gammazero/workerpool)  限制任务执行并发数，而不是队列中的任务数量的协程池，。

## GUI

*用于构建GUI应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

*工具包 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [app](https://github.com/murlokswarm/app)  用于创建包含了 GO, HTML 和 CSS 的应用程序。支持 MacOS, Windows 正在开发中。
* [fyne](https://github.com/fyne-io/fyne)  基于材料设计的Go跨平台本机gui设计。支持:Linux, macOS, Windows, BSD, iOS和Android。
* [go-astilectron](https://github.com/asticode/go-astilectron)  使用 GO 和 HTML/JS/CSS (电子驱动)进行构建跨平台 GUI 应用程序。
* [go-gtk](http://mattn.github.io/go-gtk/)  实现了 GTK 的 Go接口。
* [go-sciter](https://github.com/sciter-sdk/go-sciter)  实现了 Sciter 的 Go 接口 : 用于现代桌面 UI 开发的可嵌入HTML/CSS/脚本引擎。可跨平台。
* [gotk3](https://github.com/gotk3/gotk3)  实现了 GTK3 的 Go接口。
* [gowd](https://github.com/dtylman/gowd)  跨平台、快速、简单的桌面UI开发，采用了GO, HTML, CSS和NW.js实现。
* [qt](https://github.com/therecipe/qt)  实现了 Qt 的 Go接口(支持Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi)。
* [ui](https://github.com/andlabs/ui)  跨平台的 Platform-native GUI 库。
* [Wails](https://wails.app)  Mac, Windows, Linux桌面应用程序，主要基于含有内置的OS HTML渲染器的HTML UI。
* [walk](https://github.com/lxn/walk)  Windows应用程序库。
* [webview](https://github.com/zserge/webview)  跨平台webview窗口，具有简单的双向JavaScript绑定(Windows / macOS / Linux)。

*交互 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [go-appindicator](https://github.com/dawidd6/go-appindicator)  实现了 libappindicator3 C库 的 Go接口。
* [gosx-notifier](https://github.com/deckarep/gosx-notifier)  OSX 桌面通知库。
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker)  用于通知计算机上的任何(可插入的)活动的 OSX 库。
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier)  OSX 睡眠/唤醒通知。
* [robotgo](https://github.com/go-vgo/robotgo)  实现跨平台的GUI系统自动化。包含了控制鼠标、键盘等功能。
* [systray](https://github.com/getlantern/systray)  跨平台 Go 库，可在通知区放置图标和菜单。
* [trayhost](https://github.com/shurcooL/trayhost)  跨平台 Go 库，可用于在主机操作系统的任务栏中放置图标。


## 硬件

*硬件交互相关的库、工具和教程。 (翻译出错了? 试试 [英文版](README_EN.md#hardware) 吧~)*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## 图片

*图像处理相关的库。 (翻译出错了? 试试 [英文版](README_EN.md#images) 吧~)*

* [bild](https://github.com/anthonynsimon/bild)  纯Go语言实现的图像处理算法合集。
* [picfit](https://github.com/thoas/picfit)  Go实现的图像调整服务器。
* [imaging](https://github.com/disintegration/imaging)  简单的Go图像处理包。
* [img](https://github.com/hawx/img)  图像处理工具的集合。
* [ln](https://github.com/fogleman/ln)  Go实现的3D线艺术（3D Line Art）渲染。
* [mergi](https://github.com/noelyahan/mergi)  用于图像处理(合并、裁剪、调整大小、水印、动画)的工具和Go库。
* [mort](https://github.com/aldor007/mort)  Go语言实现的图像存储和处理服务器。
* [mpo](https://github.com/donatj/mpo)  MPO三维照片的解码和转换工具。
* [pt](https://github.com/fogleman/pt)  Go实现的路径跟踪（path tracing）引擎。
* [imagick](https://github.com/gographics/imagick)   ImageMagick下MagickWand的C API的Go binding。
* [resize](https://github.com/nfnt/resize)  Go实现的使用常用的插值法（interpolation methods）调整图像大小的库。
* [rez](https://github.com/bamiaux/rez)  纯Go语言和SIMD实现的图像大小调整。
* [smartcrop](https://github.com/muesli/smartcrop)  为任意图片寻找合适的位置进行图片裁剪。
* [steganography](https://github.com/auyer/steganography)  纯Go实现的LSB隐写（LSB steganography）的库。
* [stegify](https://github.com/DimitarPetrov/stegify)   Go实现的LSB隐写（LSB steganography），能够隐藏任何文件到一个图像中。
* [svgo](https://github.com/ajstarks/svgo)   Go实现的SVG生成库。
* [imaginary](https://github.com/h2non/imaginary)  用于图像大小调整的快速、简单的HTTP微服务。
* [image2ascii](https://github.com/qeesung/image2ascii)  将图像转换为ASCII码。
* [bimg](https://github.com/h2non/bimg)  使用libvips实现的快速高效的图像处理包。
* [gltf](https://github.com/qmuntal/gltf)  一个高效、健壮的glTF 2.0文件读取、写入和验证器。
* [cameron](https://github.com/aofei/cameron)  一个Go语言的头像生成器。
* [canvas](https://github.com/tdewolff/canvas)  矢量图形到PDF, SVG或光栅图像。
* [darkroom](https://github.com/gojek/darkroom)  An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.
* [geopattern](https://github.com/pravj/geopattern)  由字符串创建漂亮图案的图片生成器。
* [gg](https://github.com/fogleman/gg)  纯Go语言实现的2D渲染。
* [gift](https://github.com/disintegration/gift)  图像处理包。
* [go-cairo](https://github.com/ungerik/go-cairo)   cairo图形库的Go binding。
* [govatar](https://github.com/o1egl/govatar)  生成有趣头像的库和CMD工具。
* [go-gd](https://github.com/bolknote/go-gd)   GD库的Go binding。
* [go-nude](https://github.com/koyachi/go-nude)  Go语言实现的裸照检测工具。
* [go-opencv](https://github.com/lazywei/go-opencv)  OpenCV库的Go bindings。
* [go-webcolors](https://github.com/jyotiska/go-webcolors)  Python下webcolors库的Go语言调用。
* [gocv](https://github.com/hybridgroup/gocv)  使用OpenCV 3.3+实现的计算机视觉(ComputerVision)的Go语言包。
* [goimagehash](https://github.com/corona10/goimagehash)   图像哈希处理的Go语言包。
* [goimghdr](https://github.com/corona10/goimghdr)  Go语言实现的imghdr模块用于确定文件的图像类型。
* [tga](https://github.com/ftrvxmtrx/tga)  tga包是一个TARGA图像的解码、编码器。

## 物联网

*物联网设备编程库。 (翻译出错了? 试试 [英文版](README_EN.md#iot-(internet-of-things)) 吧~)*

* [connectordb](https://github.com/connectordb/connectordb)  自我量化和物联网的开源平台。
* [devices](https://github.com/goiot/devices)  一套用于物联网设备的库，实验性地用于x/exp/io。
* [eywa](https://github.com/xcodersun/eywa)  Eywa是一个用于跟踪连接的设备连接管理器。
* [flogo](https://github.com/tibcosoftware/flogo)  Flogo是一个面向物联网边缘应用和集成的开源框架。
* [gatt](https://github.com/paypal/gatt)  Gatt是一个用于构建低能耗蓝牙外围设备的Go语言包。
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot是一个用于机器人、物理计算和物联网的框架。
* [huego](https://github.com/amimof/huego)  一个包含广泛的Philips Hue客户端的Go语言库。
* [iot](https://github.com/vaelen/iot/)  IoT是一个实现谷歌物联网核心设备的简单框架。
* [mainflux](https://github.com/Mainflux/mainflux)  工业物联网消息和设备管理服务器。
* [periph](https://periph.io/)  外围设备I/O与低级板(low-level board)设备接口。
* [sensorbee](https://github.com/sensorbee/sensorbee)  轻量级物联网流处理引擎。

## 作业调度器

*用于作业调度的库。 (翻译出错了? 试试 [英文版](README_EN.md#job-scheduler) 吧~)*

* [clockwerk](http://github.com/onatm/clockwerk)  使用简单、流畅的语法调度作业的Go语言库。
* [clockwork](https://github.com/whiteShtef/clockwork)  Go实现的简单直观的作业调度库。
* [go-cron](https://github.com/rk/go-cron)  一个Go实现的简单的定时任务库。可以在不同的时间间隔（每秒一次到在每年在特定的日期执行）执行闭包或函数。主要用于web应用和长时间运行的守护进程。
* [gron](https://github.com/roylee0704/gron)  使用简单的Go API定义基于时间的任务。 之后Gron的调度程序将运行它们。
* [JobRunner](https://github.com/bamzi/jobrunner)  智能和功能丰富的cron作业调度程序（包含任务队列和实时监控）。
* [jobs](https://github.com/albrow/jobs)  持久和灵活的后台作业库。
* [leprechaun](https://github.com/kilgaloon/leprechaun)  支持webhook、crons和经典调度的作业调度程序。
* [scheduler](https://github.com/carlescere/scheduler)  Cronjobs让调度变得很简单。

## JSON

*用于JSON处理的库。 (翻译出错了? 试试 [英文版](README_EN.md#json) 吧~)*

* [ajson](https://github.com/spyzhov/ajson)  为Go语言开发的包含JSONPath支持的抽象JSON。
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  将JSON转换为Go的结构（struct）。
* [mapslice-json](https://github.com/mickep76/mapslice-json)  去MapSlice的有序封送/解封的JSON地图。
* [kazaam](https://github.com/Qntfy/kazaam)  用于任意JSON文档转换的API。
* [jsonhal](https://github.com/RichardKnop/jsonhal)  让自定义结构（struct）转化为JSON兼容的HAL（Hypertext Application Language）返回数据的简单Go包。
* [jsongo](https://github.com/ricardolonga/jsongo)  使用Fluent API来更容易地创建Json对象。
* [jsonf](https://github.com/miolini/jsonf)  用于高亮展示和查询JSON的控制台工具。
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors)  基于JSON API错误引用的Go bindings。
* [json2go](https://github.com/m-zajac/json2go)  高级JSON去结构转换。提供一个包，该包可以解析多个JSON文档并创建结构体以适应所有这些文档。
* [jettison](https://github.com/wI2L/jettison)  高性能，无反射的JSON编码器去。
* [ej](https://github.com/lucassscaravelli/ej)  简洁地编写和读取来自不同来源的JSON。
* [JayDiff](https://github.com/yazgazan/jaydiff)  用Go编写的JSON比对工具。
* [gojson](https://github.com/ChimeraCoder/gojson)  从JSON自动生成Go的结构（struct）定义。
* [gojq](https://github.com/elgs/gojq)  Go语言实现的JSON请求。
* [go-respond](https://github.com/nicklaw5/go-respond)  用于处理常见的HTTP JSON响应的Go语言库。
* [go-jsonerror](https://github.com/ddymko/go-jsonerror)  Go-JsonError允许我们轻松创建符合JsonApi规范的json响应错误。
* [GJSON](https://github.com/tidwall/gjson)  使用一行代码获取JSON的值。
* [gjo](https://github.com/skanehira/gjo)  用于创建JSON对象的小工具。
* [mp](https://github.com/sanbornm/mp)  简单的cli电子邮件解析器。它目前接受stdin并输出JSON。

## 日志记录

*用于生成和处理日志文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#logging) 吧~)*

* [distillog](https://github.com/amoghe/distillog)  distilled日志记录(可以将其视为stdlib +日志)。
* [onelog](https://github.com/francoispqt/onelog)  Onelog是一个非常简单但非常高效的JSON日志程序。它是所有场景中速度最快的JSON日志程序。而且，它是配置要求最低的日志记录器之一。
* [logo](https://github.com/mbndr/logo)  Go的日志工具，可配置的日志写入器。
* [logrus](https://github.com/Sirupsen/logrus)  Go的结构化日志操作 。
* [logrusiowriter](https://github.com/cabify/logrusiowriter)  io。作者的实现使用[logrus](https://github.com/sirupsen/logrus) logger。
* [logrusly](https://github.com/sebest/logrusly)  [logrus](https://github.com/sirupsen/logrus)的插件，将错误信息发送到[Loggly](https://www.loggly.com/)。
* [logutils](https://github.com/hashicorp/logutils)  Go的用于更好地进行日志操作的实用程序，继承了标准日志库。
* [logxi](https://github.com/mgutz/logxi)  12-factor app的日志程序，快速且让人高兴地使用。
* [lumberjack](https://github.com/natefinch/lumberjack)  简单的滚动日志，io.WriteCloser的实现。
* [mlog](https://github.com/jbrodriguez/mlog)  简单的go日志模块，有5个级别，可选循环(rotation)日志文件记录功能和stdout/stderr输出。
* [ozzo-log](https://github.com/go-ozzo/ozzo-log)  支持日志多等级、分类和过滤的高性能日志记录。可以发送过滤后的日志消息到各种目标(如控制台，网络，邮件)。
* [logger](https://github.com/azer/logger)  Go的精简日志库。
* [rollingwriter](https://github.com/arthurkiller/rollingWriter)  RollingWriter是一个自动循环的io.Writer的实现,带有多种策略以提供日志文件循环(rotation)。
* [seelog](https://github.com/cihub/seelog)  具有灵活调度、过滤和格式化的日志功能。
* [spew](https://github.com/davecgh/go-spew)  为Go数据结构实现一个漂亮的printer用于帮助调试。
* [stdlog](https://github.com/alexcesaro/log)  Stdlog是一个面向对象的库，提供了多等级日志记录。它对cron任务非常有用。
* [tail](https://github.com/hpcloud/tail)  努力模拟实现BSD的tail的特性的Go包。
* [xlog](https://github.com/xfxdev/xlog)  插件架构和灵活的日志系统，带有多级别、多日志目标和自定义日志格式。
* [xlog](https://github.com/rs/xlog)  针对'net/context`实现的结构化的记录器，用于HTTP处理程序。
* [zap](https://github.com/uber-go/zap)  快速、结构化、多等级的日志记录。
* [logmatic](https://github.com/borderstech/logmatic)  Go的彩色日志记录器，带有可配置的日志级别。
* [logex](https://github.com/chzyer/logex)  由标准日志库封装的Go日志库，支持跟踪和多等级。
* [glg](https://github.com/kpango/glg)  glg是一个简单而快速的Go日志库。
* [gologger](https://github.com/sadlil/gologger)  为Go提供方便简单的日志操作; 在彩色控制台，简单控制台，文件或Elasticsearch上。
* [glo](https://github.com/lajosbencz/glo)  参照PHP的Monolog实现的具有相同日志等级的Go日志库。
* [glog](https://github.com/golang/glog)  为Go提供了多等级日志记录。
* [go-cronowriter](https://github.com/utahta/go-cronowriter)  基于当前日期和时间的自动日志文件写入工具，类似cronolog。
* [go-log](https://github.com/pieterclaerhout/go-log)  具有strack跟踪、对象转储和可选时间戳的日志记录库。
* [go-log](https://github.com/subchen/go-log)  Go实现的简单且可配置的日志工具，并带有多等级、日志格式化和日志写入工具。
* [go-log](https://github.com/siddontang/go-log)  支持多等级和多处理程序的日志库。
* [go-log](https://github.com/ian-kent/go-log)  Log4j的Go语言。
* [go-logger](https://github.com/apsdehal/go-logger)  简单的日志程序的 Go 程序，与级别处理程序。
* [gomol](https://github.com/aphistic/gomol)  为Go实现可多方式输出、结构化日志, 并可扩展日志输出方式。
* [logdump](https://github.com/ewwwwwqm/logdump)  用于多等级级日志记录的包。
* [gone/log](https://github.com/One-com/gone/tree/master/log)  快速、可扩展、功能齐全、std-lib源兼容的日志库。
* [journald](https://github.com/ssgreg/journald)   Go实现systemd Journal的原生API用于日志记录。
* [log](https://github.com/aerogo/log)  一个O(1)日志系统，允许您将一个日志连接到多个日志写入(例如stdout、文件和TCP连接)。
* [log](https://github.com/apex/log)  Go的结构化日志包。
* [log](https://github.com/go-playground/log)  Go的简单、可配置和可伸缩的结构化日志。
* [log](https://github.com/teris-io/log)  Go的结构化日志接口，清晰地将日志facade与其实现(implementation)分离开来。
* [log-voyage](https://github.com/firstrow/logvoyage)  用Go编写的功能齐全的日志写入saas。
* [log15](https://github.com/inconshreveable/log15)  简单、强大的日志操作。
* [zerolog](https://github.com/rs/zerolog)  Zero-allocation JSON日志记录。

## 机器学习

*机器学习库。 (翻译出错了? 试试 [英文版](README_EN.md#machine-learning) 吧~)*

* [bayesian](https://github.com/jbrukh/bayesian)  Go的朴素贝叶斯分类。
* [neural-go](https://github.com/schuyler/neural-go)  多层感知器网络在Go中的实现，使用反向传播算法进行训练。
* [gorgonia](https://github.com/gorgonia/gorgonia)  基于图形（graph-based）的计算库，如Theano：它为构建各种机器学习和神经网络算法提供了基本框架。
* [gorse](https://github.com/zhenghaoz/gorse)  基于协同过滤的离线推荐系统后端。
* [goscore](https://github.com/asafschers/goscore)   为预言模型标记语言（PMML）实现的评分API。
* [gosseract](https://github.com/otiai10/gosseract)  使用c++的Tesseract库实现的OCR。
* [libsvm](https://github.com/datastream/libsvm)  基于LIBSVM 3.14实现。
* [neat](https://github.com/jinyeom/neat)  即插即用的并行Go框架，用于增强拓扑的神经进化(NeuroEvolution of Augmenting Topologies)。
* [ocrserver](https://github.com/otiai10/ocrserver)  一个简单的OCR API服务器，非常容易地使用Docker和Heroku部署。
* [Goptuna](https://github.com/c-bata/goptuna)  Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.
* [onnx-go](https://github.com/owulveryck/onnx-go)  Go Interface， 用于开放式神经网络交换(Open Neural Network Exchange)。
* [probab](https://github.com/ThePaw/probab)  概率分布函数。贝叶斯推理。使用Go写的。
* [randomforest](https://github.com/malaschitz/randomForest)  容易使用随机森林库去。
* [regommend](https://github.com/muesli/regommend)  推荐和协同过滤引擎。
* [shield](https://github.com/eaigner/shield)  贝叶斯文本分类器，具有灵活的tokenizers和存储后端。
* [tfgo](https://github.com/galeone/tfgo)  易于使用的Tensorflow bindings:简化了官方Tensorflow Go bindings的使用。在Go中定义计算图形，在Python中加载和执行训练的模型。
* [goRecommend](https://github.com/timkaye11/goRecommend)  用Go编写的推荐算法库。
* [gonet](https://github.com/dathoangnd/gonet)  用于围棋的神经网络。
* [CloudForest](https://github.com/ryanbressler/CloudForest)  快速、灵活、多线程集成的决策树，用于机器学习。
* [go-galib](https://github.com/thoj/go-galib)  用Go编写的遗传算法库。
* [eaopt](https://github.com/MaxHalford/eaopt)  一个进化优化（evolutionary optimization）库。
* [evoli](https://github.com/khezen/evoli)  遗传算法（Genetic Algorithm）和粒子群优化（Particle Swarm Optimization）库。
* [fonet](https://github.com/Fontinalis/fonet)  一个用Go编写的深度神经网络库。
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster)  Go实现的k-modes和k-prototypes聚类算法。
* [go-deep](https://github.com/patrikeh/go-deep)  一个功能丰富的神经网络库 。
* [go-fann](https://github.com/white-pony/go-fann)  快速人工神经网络(FANN)库的Go bindings。
* [go-pr](https://github.com/daviddengcn/go-pr)  Go编写的模式识别包。
* [goml](https://github.com/cdipaolo/goml)  在线机器学习。
* [gobrain](https://github.com/goml/gobrain)  用 Go 编写的神经网络库。
* [godist](https://github.com/e-dard/godist)  各种概率分布，以及相关的method。
* [goga](https://github.com/tomcraven/goga)  Go的遗传算法库。
* [GoLearn](https://github.com/sjwhitworth/golearn)  通用机器学习库。
* [golinear](https://github.com/danieldk/golinear)   Go实现的liblinear bindings。
* [GoMind](https://github.com/surenderthakran/gomind)  一个简单的神经网络库。
* [Varis](https://github.com/Xamber/Varis)  Go实现的神经网络。

## 消息

*实现消息传递系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#messaging) 吧~)*

* [ami](https://github.com/kak-tus/ami)  基于Redis集群流的客户端到可靠队列。
* [messagebus](https://github.com/vardius/message-bus)  messagebus是一个Go的简单异步消息总线，非常适合在执行事件sourcing、CQRS和DDD时用作事件总线。
* [guble](https://github.com/smancke/guble)  使用推送通知服务(谷歌Firebase云消息、苹果推送通知服务、SMS)的消息服务器, 支持websockets,REST API，并具有分布式操作和消息持久性。
* [hub](https://github.com/leandro-lugaresi/hub)  适用于Go应用程序的消息/事件中心，使用publish/subscribe模式，并支持别名(类似rabbitMQ exchanges)。
* [jazz](https://github.com/socifi/jazz)  一个简单的RabbitMQ抽象层，用于队列管理和消息的发布和消费。
* [machinery](https://github.com/RichardKnop/machinery)  基于分布式消息传递的异步任务/作业队列。
* [mangos](https://github.com/go-mangos/mangos)  Nanomsg(“可伸缩协议”)的纯go实现,具有传输互操作性。
* [melody](https://github.com/olahol/melody)  处理websocket session的极简框架，包括广播和自动ping/pong处理。
* [Mercure](https://github.com/dunglas/mercure)  使用Mercure协议分派服务器发送(server-sent)更新的服务器和库(构建在服务器发送事件之上)。
* [NATS Go Client](https://github.com/nats-io/nats)  轻量级和高性能的发布-订阅(publish-subscribe)和分布式队列消息传递系统——这是一个Go库。
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster)  gopush-cluster是一个gopush服务器集群。
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus)  一个围绕NSQ topic和channel的小工具。
* [oplog](https://github.com/dailymotion/oplog)  用于REST api的通用oplog/replication系统。
* [pubsub](https://github.com/tuxychandru/pubsub)  简单的pubsub的go包。
* [rabbus](https://github.com/rafaeljesus/rabbus)  amqp exchanges和队列上的一个小工具。
* [rabtap](https://github.com/jandelgado/rabtap)  RabbitMQ的瑞士军刀cli应用。
* [RapidMQ](https://github.com/sybrexsys/RapidMQ)  RapidMQ是用于管理本地消息队列的轻量且可靠的库。
* [redisqueue](https://github.com/robinjoseph08/redisqueue)  基于Redis streams的队列的生产者和消费者。
* [APNs2](https://github.com/sideshow/apns2)  HTTP / 2苹果消息推送provider——发送推送消息到iOS, tvOS, Safari和OSX应用。
* [gorush](https://github.com/appleboy/gorush)  使用[APNs2](https://github.com/sideshow/apns2)和谷歌[GCM](https://github.com/google/go-gcm)推送通知服务器。
* [golongpoll](https://github.com/jcuga/golongpoll)  HTTP longpoll服务器库，使web发布-订阅变得简单。
* [emitter](https://github.com/olebedev/emitter)  使用Go的方式发出事件, 带有通配符、谓词、取消可能性和许多其他优点。
* [Beaver](https://github.com/Clivern/Beaver)  一个实时消息服务器，可用于在web和手机app端构建一个可伸缩的应用内通知，多人游戏，聊天应用。
* [Benthos](https://github.com/Jeffail/benthos)  一系列协议之间的消息流桥接。
* [Bus](https://github.com/mustafaturan/bus)  内部通信的最小消息总线实现。
* [Centrifugo](https://github.com/centrifugal/centrifugo)  实时消息(Websockets或SockJS)服务器。
* [Commander](https://github.com/jeroenrinzema/commander)  高级事件驱动的消费者/生产者(consumer/producer)，支持各种“方言”，如Apache Kafka。
* [dbus](https://github.com/godbus/dbus)  D-Bus的Go bindings。
* [Gollum](https://github.com/trivago/gollum)  一个n:m多路复用器(n:m multiplexer)，收集不同来源的消息并将其广播到一组目的地。
* [drone-line](https://github.com/appleboy/drone-line)  使用二进制、docker或Drone CI发送[Line](https://at.line.me/en)通知。
* [event](https://github.com/agoalofalife/event)  观察者模式的实现。
* [EventBus](https://github.com/asaskevich/EventBus)  具有异步兼容性的轻量级事件总线。
* [gaurun-client](https://github.com/osamingo/gaurun-client)  用Go编写的Gaurun客户端。
* [Glue](https://github.com/desertbit/glue)  健壮的Go和Javascript Socket库(替代Socket.io)。
* [go-notify](https://github.com/TheCreeper/go-notify)  原生的freedesktop通知规范实现。
* [go-nsq](https://github.com/nsqio/go-nsq)  NSQ的官方Go包。
* [go-socket.io](https://github.com/googollee/go-socket.io)  go的socket.io库，一个实时应用程序框架。
* [go-vitotrol](https://github.com/maxatome/go-vitotrol)  用于Viessmann Vitotrol web服务的客户端库。
* [rmqconn](https://github.com/sbabiv/rmqconn) **star:3** RabbitMQ重新连接。amqp.Connection和amqp.Dial之上的Wrapper。允许在强制调用Close()方法关闭连接之前重新连接。   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/rmqconn)
* [sarama](https://github.com/Shopify/sarama)   Apache Kafka的Go库。
* [Uniqush-Push](https://github.com/uniqush/uniqush-push)  Redis支持的统一推送服务, 用于服务端向移动设备的消息通知。
* [zmq4](https://github.com/pebbe/zmq4)  ZeroMQ的Go interface第4版。也可用于[第3版](https://github.com/pebbe/zmq3)和[第2版](https://github.com/pebbe/zmq2)。

## 微软办公软件

* [unioffice](https://github.com/unidoc/unioffice)  用于创建和处理Office Word (.docx)、Excel (.xlsx)和Powerpoint (.pptx)文档的纯go库。

### Microsoft Excel

*用于操作Microsoft Excel的库。 (翻译出错了? 试试 [英文版](README_EN.md#microsoft-excel) 吧~)*

* [xlsx](https://github.com/tealeg/xlsx) **star:3902** 用以简化在Go程序中读取使用最新版本Microsoft Excel的XML格式文件的库。   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx)  在Go程序以快速和安全的方式读取/更新现有的Microsoft Excel文件。
* [excelize](https://github.com/360EntSecGroup-Skylar/excelize)  用于读写Microsoft Excel™(XLSX)文件的Go语言库。
* [go-excel](https://github.com/szyhf/go-excel)  一个简单轻便的阅读器，可以将类关系型数据库(relate-db-like)的excel作为表来读取。
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter)  libxlsxwriter的Go binding, 用于编写XLSX (Microsoft Excel)文件。

## 杂项

### 依赖注入

*用于处理依赖项注入的库。 (翻译出错了? 试试 [英文版](README_EN.md#dependency-injection) 吧~)*

* [alice](https://github.com/magic003/alice)  Go的外挂的依赖注入容器。
* [container](https://github.com/golobby/container)  一个强大的IoC容器，具有流畅且易于使用的接口。
* [dig](https://github.com/uber-go/dig)  一个基于反射的Go依赖注入工具包。
* [dingo](https://github.com/i-love-flamingo/dingo)  基于Guice的Go依赖注入工具包。
* [fx](https://github.com/uber-go/fx)  基于依赖注入的Go应用程序框架(构建在dig之上)。
* [gocontainer](https://github.com/vardius/gocontainer) **star:12** 简单的依赖注入容器。   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)
* [inject](https://github.com/defval/inject)  一个基于反射的依赖注入容器，具有简单的接口。
* [linker](https://github.com/logrange/linker)  A reflection based dependency injection and inversion of control library with components lifecycle support.
* [wire](https://github.com/Fs02/wire)  适用于Go的严格运行时依赖注入(Strict Runtime Dependency Injection)。

### 项目布局

*用于组织项目的非正式模式集。 (翻译出错了? 试试 [英文版](README_EN.md#project-layout) 吧~)*

* [scaffold](https://github.com/catchplay/scaffold) **star:55** 快速生成Go项目布局的脚手架。让您专注于已实现的业务逻辑。   [![最近一年没有更新][Yellow]](https://github.com/catchplay/scaffold)   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)  Go生态系统中历史和新兴的项目布局模式集合。
* [go-sample](https://github.com/zitryss/go-sample) **star:45** 带有实际代码的Go应用程序项目的示例布局。   [![最近一年没有更新][Yellow]](https://github.com/zitryss/go-sample)   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application)  应用程序样板和应用现代实践的例子。
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang)  遵循生产最佳实践快速启动项目的Go应用程序样板模板。

### 字符串

*处理字符串的库。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [strutil](https://github.com/ozgio/strutil)  字符串处理工具。
* [xstrings](https://github.com/huandu/xstrings)  从其他语言移植的有用字符串函数合集。

*这些库之所以放在这里，是因为不适合放在其他分类。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4549** 用于检索进程和系统利用率(CPU、内存、磁盘等)的跨平台的库。   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![最近一周有更新][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [gosh](https://github.com/osamingo/gosh)  提供Go统计处理程序，结构和测量方法。
* [archiver](https://github.com/mholt/archiver) **star:2735** 用于生成和解压.zip和.tar.gz文档的库和命令。   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![最近一周有更新][Green]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [autoflags](https://github.com/artyom/autoflags)  从struct字段自动定义命令行flag的Go包。
* [avgRating](https://github.com/kirillDanshin/avgRating)  根据威尔逊得分排序算法(Wilson Score Equation)计算平均分和评分。
* [banner](https://github.com/dimiro1/banner)  在Go应用程序中添加漂亮的横幅(banner)。
* [gosms](https://github.com/haxpax/gosms) **star:1268** Go编写的私人的本地短信网关，可以用来发送短信。   [![最近一年没有更新][Yellow]](https://github.com/haxpax/gosms)   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [gotoprom](https://github.com/cabify/gotoprom)  为Prometheus客户端提供类型安全的指标(metric)构建工具库。
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:973**  Go语言弹性模式(resiliency pattern)。   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [gommit](https://github.com/antham/gommit)  分析git提交消息，确保它们遵循已定义的格式。
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:825** 用go编写的随机数据生成器。   [![最近一周有更新][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:766** base64Captcha支持数字，字母，算术，音频和混合模式的验证码。   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![包含中文文档][CN]](https://github.com/mojocn/base64Captcha)
* [battery](https://github.com/distatus/battery)  跨平台、标准化的电池信息库。
* [bitio](https://github.com/icza/bitio)  高度优化的位级读写器。
* [shortid](https://github.com/teris-io/shortid) **star:524** 分布式地生成超短、唯一、非顺序、URL友好的id。   [![最近一年没有更新][Yellow]](https://github.com/teris-io/shortid)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:495** 用于在纯Go中与LLVM IR交互的库。   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:380** 易于使用，可扩展的健康检查库。   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [go-openapi](https://github.com/go-openapi)  用于解析和使用开放api模式(open-api schemas)的包的集合。
* [ghorg](https://github.com/gabrie30/ghorg)  快速克隆整个org用户库到一个目录-支持GitHub, GitLab和Bitbucket。
* [ffmt](https://github.com/go-ffmt/ffmt)  美化数据,使其更适合人查看。
* [datacounter](https://github.com/miolini/datacounter)  用于readers/writer/http.ResponseWriter的计数器。
* [conv](https://github.com/cstockton/go-conv) **star:350** conv包提供了跨Go类型(Go types)的快速和直观的转换。   [![最近一年没有更新][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [go-commons-pool](https://github.com/jolestar/go-commons-pool)  Go语言的通用对象池。
* [gountries](https://github.com/pariz/gountries) **star:232** 获取国家和细节数据的包。   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [stateless](https://github.com/qmuntal/stateless) **star:148** 用于创建状态机的流畅库。   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [lk](https://github.com/hyperboloide/lk) **star:143** 一个简单的版权许可证库。   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [stats](https://github.com/go-playground/stats) **star:131** Monitors Go MemStats + 诸如如内存，Swap和CPU的系统状态统计，并通过UDP发送到任何你想记录的地方   [![最近一年没有更新][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:105** 用于RESTful服务的强制的(opinionated)并发健康检查HTTP处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:94** Go的Emojis库。   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:86** 用于RAR、TAR、ZIP和7z文件的解压缩库。   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [anagent](https://github.com/mudler/anagent)  Go语言的最小化，可插入的evloop/timer处理程序, 带有依赖注入。
* [antch](https://github.com/antchfx/antch)  一个快速、强大和可扩展的web爬虫框架。
* [captcha](https://github.com/steambap/captcha) **star:53** captcha包为验证码生成提供了一个易于使用的、unopinionated的API。   [![最近一周有更新][Green]](https://github.com/steambap/captcha)   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [xkg](https://github.com/go-xkg/xkg) **star:48** X Keyboard Grabber(键盘事件捕获)   [![最近一年没有更新][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [persian](https://github.com/mavihq/persian) **star:40** 一些适用于波斯语的Go工具库。   [![最近一年没有更新][Yellow]](https://github.com/mavihq/persian)   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:38** 通过Json请求生成PDF的HTTP服务。   [![最近一年没有更新][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:31** 用于[Browser Capabilities Project](http://browscap.org/)的Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg)规范在Go中的实现。   [![最近一年没有更新][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:19** 一个现代的、强大的、健壮的URL转短链接微服务，带有mysql支持。   [![最近一年没有更新][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  用于生成http输入和输出处理模板。
* [sandid](https://github.com/aofei/sandid) **star:14** 能沟让地球上的每一粒沙子都有自己的ID。   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [indigo](https://github.com/osamingo/indigo)  分布式唯一ID生成器, 使用Sonyflake并由Base58编码。
* [shellwords](https://github.com/Wing924/shellwords) **star:8** 一个Go库，实现了依照UNIX Bourne shell的关键词解析规则进行字符串操纵。   [![最近一年没有更新][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** 一个用于打包和解包FQDNs列表的go语言库。   [![最近一年没有更新][Yellow]](https://github.com/Wing924/hostutils)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [metrics](https://github.com/pascaldekloe/metrics) **star:6** 用于instrumentation和Prometheus exposition的库。   [![最近一周有更新][Green]](https://github.com/pascaldekloe/metrics)   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [morse](https://github.com/alwindoss/morse)  实现字符串与摩尔斯电码转换的库。
* [numa](https://github.com/lrita/numa) **star:2** NUMA是一个用go编写的实用程序库。它可以帮助我们编写一些NUMA-AWARED代码。   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## 自然语言处理

*用于处理人类语言的库。 (翻译出错了? 试试 [英文版](README_EN.md#natural-language-processing) 吧~)*

* [prose](https://github.com/jdkato/prose) **star:2435** 支持标记化、词性标记、命名实体提取等文本处理的库。只说英语。   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1208** 高效的文本分割;支持英语、汉语、日语等。   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![包含中文文档][CN]](https://github.com/go-ego/gse)
* [when](https://github.com/olebedev/when) **star:1060** 带有可插入规则的自然EN和RU语言日期/时间解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:983** 这是一个Go实现的[jieba](https://github.com/fxsjy/jieba)，这是一个中文分词算法。   [![最近一周有更新][Green]](https://github.com/yanyiwu/gojieba)   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![包含中文文档][CN]](https://github.com/yanyiwu/gojieba)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:647** 中文汉字到汉语拼音的转换。   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:472** JP形态学分析仪编写的纯Go。   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:397** Go的自然语言检测包。支持84种语言和24种脚本(书写系统，如拉丁语、西里尔语等)。   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:361** 从字符串中提取值并用nlp填充结构。   [![最近一年没有更新][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:266** 句子标记器:将文本转换为句子列表。   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:241** 支持LSA(潜在语义分析)的Go自然语言处理库。   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  软件包和用于处理本地化文本的附带工具。
* [getlang](https://github.com/rylans/getlang) **star:83** 快速自然语言检测包。   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-nlp](https://github.com/nuance/go-nlp) **star:82** 用于处理离散概率分布的实用程序和用于进行NLP工作的其他工具。   [![最近一年没有更新][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [gounidecode](https://github.com/fiam/gounidecode) **star:68** 用于Go的Unicode音译器(也称为unidecode)。   [![最近一年没有更新][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:67** Unicode文本的ASCII音译。   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [textcat](https://github.com/pebbe/textcat) **star:62** Go package支持基于n-gram的文本分类，支持utf-8和原始文本。   [![最近一年没有更新][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** 这是一个 Go 实现的[MMSEG](http://technology.chtsai.org/mmseg/)，这是一个中文分词算法。   [![最近一年没有更新][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:56** 波特词干算法的实现。   [![最近一年没有更新][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:55** 快速自动关键字提取算法(RAKE)的Go端口。   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [stemmer](https://github.com/dchest/stemmer) **star:49** 用于Go编程语言的Stemmer包。包括英语和德语词根。   [![最近一年没有更新][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [segment](https://github.com/blevesearch/segment) **star:47** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![最近一年没有更新][Yellow]](https://github.com/blevesearch/segment)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [porter2](https://github.com/zhenjl/porter2) **star:37** 非常快的波特2史坦默。   [![最近一年没有更新][Yellow]](https://github.com/zhenjl/porter2)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [go2vec](https://github.com/danieldk/go2vec) **star:33** 用于word2vec嵌入式的阅读器和实用程序函数。   [![最近一年没有更新][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [petrovich](https://github.com/striker2000/petrovich) **star:28** 彼得罗维奇是一个图书馆，它把俄语名字的词形变化成特定的语法格。   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang实现了Paice/外壳阻塞算法。   [![最近一年没有更新][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [snowball](https://github.com/goodsign/snowball) **star:25** 滚雪球柄端口(cgo包装)为 Go 。提供词干提取功能[Snowball native](http://snowball.tartarus.org/)。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/snowball)
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo绑定到Yandex。Mystem -俄罗斯形态学分析仪。   [![最近一年没有更新][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo绑定用于icu4c C库的检测和转换功能。保证与版本50.1兼容。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:17**  Go 绑定斯诺鲍libstemmer库，包括波特2。   [![最近一年没有更新][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [shamoji](https://github.com/osamingo/shamoji) **star:11** shamoji是用Go编写的word过滤包。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** 用于libtextcat C库的Cgo绑定。保证与版本2.2兼容。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [porter](https://github.com/a2800276/porter) **star:8** 这是Martin Porter在C语言中实现的Porter词干分析算法的一个相当简单的移植。   [![最近一年没有更新][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** 一种基于字典和双字母格朗语言模型的记号赋予器。(现在只支持中文分割)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)
* [go-localize](https://github.com/m1/go-localize) **star:3** 简单易用的i18n(国际化和本地化)引擎——用于翻译语言环境字符串。   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-localize)
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:1** 语言检测API Go客户端。支持批量请求，短短语或单字语言检测。   [![godoc][GoDoc]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)

## 网络

*用于处理各种网络层的库。 (翻译出错了? 试试 [英文版](README_EN.md#networking) 吧~)*

* [fasthttp](https://github.com/valyala/fasthttp) **star:11576** 一个快速HTTP实现，比net/http快10倍。   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [gopacket](https://github.com/google/gopacket)  Go library for packet processing with libpcap bindings.
* [golibwireshark](https://github.com/sunwxg/golibwireshark)  用于解码 pcap 文件和分析解剖数据。
* [gobgp](https://github.com/osrg/gobgp) **star:1813** 基于 Go 的 BGP 实现。   [![最近一周有更新][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [gnet](https://github.com/panjf2000/gnet) **star:1692** “gnet”是一个用纯Go编写的高性能、轻量级、非阻塞、事件驱动的网络框架。   [![最近一周有更新][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![包含中文文档][CN]](https://github.com/panjf2000/gnet)
* [ftp](https://github.com/jlaffaye/ftp)  实现了[RFC 959](http://tools.ietf.org/html/rfc959)中描述的ftp客户端。
* [fortio](https://github.com/fortio/fortio) **star:1155** 负载测试库和命令行工具，高级的echo服务器和web UI。允许指定一组每秒查询的负载，并记录延迟直方图和其他有用的统计数据，并将它们作图。支持Tcp、Http、gRPC。   [![最近一周有更新][Green]](https://github.com/fortio/fortio)   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [go-powerdns](https://github.com/joeig/go-powerdns)  Golang的PowerDNS API绑定。
* [go-getter](https://github.com/hashicorp/go-getter) **star:854**  通过URL来下载文件或目录。   [![最近一周有更新][Green]](https://github.com/hashicorp/go-getter)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [grab](https://github.com/cavaliercoder/grab) **star:639**  用于管理文件下载。   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [HTTPLab](https://github.com/gchaincl/httplab)  HTTPLabs 允许你检查 HTTP 请求和伪造响应。
* [graval](https://github.com/koofr/graval)  FTP服务器框架。
* [gotcp](https://github.com/gansidui/gotcp)  用于快速编写 tcp 应用程序。
* [gosocsvr](https://github.com/rakeki/gosocsvr)  Socket服务器变得简单。
* [gosnmp](https://github.com/soniah/gosnmp) **star:491** 用于执行 SNMP 操作的原生 Go 库。   [![最近一周有更新][Green]](https://github.com/soniah/gosnmp)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [dns](https://github.com/miekg/dns)  用于 DNS 的库。
* [cidranger](https://github.com/yl2chen/cidranger) **star:441** 快速得 IP 到 CIDR 查找。   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [dhcp6](https://github.com/mdlayher/dhcp6)  实现了一个DHCPv6服务器，如RFC 3315所述。
* [gopcap](https://github.com/akrennmair/gopcap) **star:376**  用 Go 实现了对 libpcap 的封装。   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [goshark](https://github.com/sunwxg/goshark)  用于解码IP包，创建用于分析的数据结构包。
* [go-stun](https://github.com/ccding/go-stun) **star:344** 实现了 STUN 客户端(RFC 3489和RFC 5389)。   [![最近一年没有更新][Yellow]](https://github.com/ccding/go-stun)   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [webrtc](https://github.com/pions/webrtc)  WebRTC API的纯Go实现。
* [tspool](https://github.com/two/tspool)  TCP库使用工作池来提高性能并保护服务器。
* [water](https://github.com/songgao/water)  简单TUN / TAP图书馆。
* [stun](https://github.com/go-rtc/stun) **star:339** Go实现的RFC 5389 STUN协议。   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server)   Go 图书馆建设tcp服务器更快。
* [winrm](https://github.com/masterzen/winrm)   Go WinRM客户端远程执行Windows机器上的命令。
* [utp](https://github.com/anacrolix/utp)  Go uTP微传输协议的实现。
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:235** 通过TCP传输协议缓冲区数据。   [![最近一年没有更新][Yellow]](https://github.com/stabbycutyou/buffstreams)   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [arp](https://github.com/mdlayher/arp) **star:214** 实现了arp协议，如RFC 826中所述。   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:196** 实现了对IEEE 802.3以太网II帧和IEEE 802.1Q VLAN标签的编组和解组。   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [canopus](https://github.com/zubairhamed/canopus) **star:137** CoAP客户端/服务器实现(RFC 7252)。   [![最近一年没有更新][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:130** Gmqtt是一个灵活、高性能的MQTT代理库，它完全实现了MQTT协议V3.1.1。   [![最近一周有更新][Green]](https://github.com/DrmagicE/gmqtt)   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![包含中文文档][CN]](https://github.com/DrmagicE/gmqtt)
* [sslb](https://github.com/eduardonunesp/sslb) **star:124** 它是一个超级简单的负载平衡器，只是一个实现某种性能的小项目。   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gNxI](https://github.com/google/gnxi) **star:117** 一组基于 gNMI 和 gNOI 协议的网络管理工具。   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [gaio](https://github.com/xtaci/gaio) **star:101** 高性能异步io网络在proactor模式下的Golang。   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gaio)
* [gev](https://github.com/Allenxuxu/gev)  gev是一个基于反应器模式的轻量级、快速的非阻塞TCP网络库。
* [xtcp](https://github.com/xfxdev/xtcp) **star:93** TCP服务器框架具有同时全双工通信，优雅关机，自定义协议。   [![最近一周有更新][Green]](https://github.com/xfxdev/xtcp)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [ether](https://github.com/songgao/ether) **star:65** 一个用于发送和接收以太网帧的跨平台 Go 库。   [![最近一年没有更新][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [quic-go](https://github.com/lucas-clemente/quic-go)  在纯Go中实现了QUIC协议。
* [ssh](https://github.com/gliderlabs/ssh)  用于构建SSH服务器的高级API(封装密码/ SSH)。
* [sftp](https://github.com/pkg/sftp)  Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.
* [raw](https://github.com/mdlayher/raw)  Package raw支持在设备驱动程序级别读取和写入网络接口的数据。
* [publicip](https://github.com/polera/publicip)  包publicip返回面向公共的IPv4地址(internet出口)。
* [iplib](https://github.com/c-robinson/iplib) **star:28** 用于处理IP地址的库(net)。借鉴于python 的 [ipaddress](https://docy-doc.org/3/library/ipaddress.html)和ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [portproxy](https://github.com/aybabtme/portproxy)  Simple TCP proxy which adds CORS support to API's which don't support it.
* [peerdiscovery](https://github.com/schollz/peerdiscovery)  基于UDP组播的跨平台本地对等点发现库。
* [packet](https://github.com/aerogo/packet)  通过TCP和UDP发送数据包。它可以缓冲消息和热交换连接。
* [NFF-Go](https://github.com/intel-go/nff-go)  用于快速开发云计算和裸机网络功能的框架(原YANFF)。
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  Paho Go客户端提供了一个 MQTT 客户端库，用于通过TCP、TLS或WebSockets连接到MQTT代理。
* [mdns](https://github.com/hashicorp/mdns)  简单mDNS(Multicast DNS)客户端/服务器库。
* [llb](https://github.com/kirillDanshin/llb)  一个非常简单、快速的代理服务器后端。可用于快速重定向到预定义域，具有零内存分配和快速响应。
* [linkio](https://github.com/ian-kent/linkio)  网络链路速度模拟，主要用于接口的读/写。
* [lhttp](https://github.com/fanux/lhttp)  强大的websocket框架，可以让你更容易的构建IM服务器。
* [kcptun](https://github.com/xtaci/kcptun)  基于KCP协议的非常简单和快速udp隧道。
* [kcp-go](https://github.com/xtaci/kcp-go)  快速可靠的ARQ协议。
* [jazigo](https://github.com/udhos/jazigo)  Jazigo是一个用Go编写的工具，用于检索多个网络设备的配置。
* [httpproxy](https://github.com/wzshiming/httpproxy) **star:2** HTTP代理处理程序和拨号。   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/httpproxy)

### HTTP客户端

*用于发出HTTP请求的库。 (翻译出错了? 试试 [英文版](README_EN.md#http-clients) 吧~)*

* [gentleman](https://github.com/h2non/gentleman)  功能齐全的插件驱动HTTP客户端库。
* [grequests](https://github.com/levigross/grequests)  一个 Go “克隆”的伟大和著名的请求库。
* [heimdall](https://github.com/gojektech/heimdall)  具有重试和hystrix功能的增强http客户机。
* [httpretry](https://github.com/ybbus/httpretry)  用重试功能丰富默认的go HTTP客户端。
* [pester](https://github.com/sethgrid/pester)  使用重试、后退和并发执行HTTP客户机调用。
* [resty](https://github.com/go-resty/resty)  简单的 HTTP 和 REST 客户端，受到 Ruby rest-client 的启发。
* [rq](https://github.com/ddo/rq)  golang stdlib HTTP客户端更好的接口。
* [sling](https://github.com/dghubble/sling)  Sling是一个用于创建和发送API请求的Go HTTP客户端库。

## OpenGL

*用于在Go中使用OpenGL的库。 (翻译出错了? 试试 [英文版](README_EN.md#opengl) 吧~)*

* [gl](https://github.com/go-gl/gl)  OpenGL 的 Go 接口实现(通过glow生成)。
* [glfw](https://github.com/go-gl/glfw)  GLFW 3 的 Go 接口实现。
* [goxjs/gl](https://github.com/goxjs/gl)  跨平台的OpenGL 接口实现(OS X, Linux, Windows，浏览器，iOS, Android)。
* [goxjs/glfw](https://github.com/goxjs/glfw)  跨平台 glfw 库，可用于创建 OpenGL 上下文并接收事件。
* [mathgl](https://github.com/go-gl/mathgl)  完全基于 Go 实现的数学软件包，专门用于处理三维数学。借鉴于 GLM。

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques. (翻译出错了? 试试 [英文版](README_EN.md#orm) 吧~)*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  强大的orm框架。支持: pq/mysql/sqlite3。
* [Marlow](https://github.com/dadleyy/marlow)  从项目结构生成ORM。
* [Xorm](https://github.com/go-xorm/xorm)  基于 Go 的简单而强大的ORM。
* [upper.io/db](https://github.com/upper/db)  对外提供统一的接口用于访问不同的存储介质，例如PostgreSQL, MySQL, SQLite, MSSQL, QL、MongoDB.。
* [SQLBoiler](https://github.com/volatiletech/sqlboiler)  ORM 生成器。根据数据库 schema 生成一个功能强大且运行速度快的ORM。
* [rel](https://github.com/Fs02/rel)  用于干净(洋葱)架构的Golang SQL存储库层。
* [reform](https://github.com/go-reform/reform)  基于非空接口和代码生成的 ORM。
* [QBS](https://github.com/coocood/qbs)  Stands for Query By Struct. A Go ORM.
* [pop/soda](https://github.com/gobuffalo/pop)  数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [lore](https://github.com/abrahambotros/lore)  Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.
* [go-firestorm](https://github.com/jschoedt/go-firestorm)  一个轻量级的ORM。用于Google/Firebase Cloud Firestore。
* [grimoire](https://github.com/Fs02/grimoire)  基于 golang 的数据库访问层和验证库。(支持: MySQL, PostgreSQL和SQLite3)。
* [gorp](https://github.com/go-gorp/gorp)  基于 Go 的关系持久性 ORM-ish 库。
* [gormt](https://github.com/xxjwxc/gormt)  Mysql数据库到golang gorm结构。
* [GORM](https://github.com/jinzhu/gorm)  一个出色的 ORM 库。主要目标是对开发人员友好。
* [go-store](https://github.com/gosuri/go-store)  简单且快速的 Redis 键值存储库。
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder)  一个灵活而强大的SQL字符串构建器库。
* [go-queryset](https://github.com/jirfag/go-queryset)  基于 GORM 100% 类型安全的 ORM。可支持 MySQL, PostgreSQL, Sqlite3, SQL Server。
* [go-pg](https://github.com/go-pg/pg)  用于 PostgreSQL 的ORM。侧重于 PostgreSQL 的特性和性能。
* [Zoom](https://github.com/albrow/zoom)  基于 Redis 的快速数据存储和查询引擎。

## 包管理

*用于管理依赖和包的官方工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules 是源码的版本控制和交换的单位。go命令直接支持处理模块，包括记录和解决对其他模块的依赖关系。

*包管理的官方实验工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [dep](https://github.com/golang/dep)   Go 的依赖管理工具，需要 Go 1.9+
* [vgo](https://go.googlesource.com/vgo/)  Go 命令版本管理

*用于包和依赖项管理的非官方库。 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [gigo](https://github.com/LyricalSecurity/gigo)  类似pip的golang依赖工具，支持私有存储库和散列。
* [glide](https://github.com/Masterminds/glide)  轻松管理您的 golang 第三方包。受Maven、Bundler和Pip等工具的启发。
* [godep](https://github.com/tools/godep)  godep是go的依赖工具，它通过修复包的依赖关系来帮助构建可重复的包。
* [gom](https://github.com/mattn/gom)  Go Manager - bundle for Go。
* [goop](https://github.com/nitrous-io/goop)  Go 的简单依赖管理器，灵感来自Bundler。
* [gop](https://github.com/lunny/gop)  在GOPATH之外构建和管理Go应用程序。
* [gopm](https://github.com/gpmgo/gopm)  包管理器。
* [govendor](https://github.com/kardianos/govendor)  包管理器。使用 vendor 文件的 Go vendor 工具。
* [gpm](https://github.com/pote/gpm)  基本的 Go 依赖管理器。
* [johnny-deps](https://github.com/VividCortex/johnny-deps)  使用Git的最小依赖版本。
* [mvn-golang](https://github.com/raydac/mvn-golang)  插件，为自动加载Golang SDK，依赖关系管理和启动Maven项目基础设施中的构建环境提供了方法。
* [nut](https://github.com/jingweno/nut)  vendor 依赖。
* [VenGO](https://github.com/DamnWidget/VenGO)  创建和管理可导出的隔离go虚拟环境。

## 性能

* [jaeger](https://github.com/jaegertracing/jaeger)  分布式跟踪系统。
* [profile](https://github.com/pkg/profile)  Go的简单分析支持包。
* [tracer](https://github.com/kamilsk/tracer)  简单、轻量级的跟踪。

## 查询语言

* [gojsonq](https://github.com/thedevsaddam/gojsonq)  一个用来查询JSON数据的Go包。
* [graphql](https://github.com/tmc/graphql)  graphql解析器+工具集
* [graphql](https://github.com/neelance/graphql-go)  关注易用性的GraphQL服务器。
* [graphql-go](https://github.com/graphql-go/graphql)  为Go实现GraphQL。
* [jsonql](https://github.com/elgs/jsonql)  Golang中的JSON查询表达式库。
* [jsonslice](https://github.com/bhmj/jsonslice)  使用高级过滤器查询Jsonpath。
* [rql](https://github.com/a8m/rql)  用于REST API的资源查询语言。
* [straf](https://github.com/SonicRoshan/straf)  轻松地将Golang结构转换为GraphQL对象。

## 嵌入的资源

* [esc](https://github.com/mjibson/esc)  将文件嵌入到Go程序中并提供http文件系统接口。
* [fileb0x](https://github.com/UnnoTed/fileb0x)  一个可定制的工具用来在Go中嵌入文件
* [go-embed](https://github.com/pyros2097/go-embed)  生成go代码，将资源文件嵌入到库或可执行文件中。
* [go-resources](https://github.com/omeid/go-resources)  嵌入到Go中的普通资源。
* [go.rice](https://github.com/GeertJohan/go.rice)  go.rice 是一个Go包，它使处理html、js、css、图像和模板等资源变得非常容易。
* [packr](https://github.com/gobuffalo/packr)  将静态文件嵌入到Go二进制文件中的简单方法。
* [statics](https://github.com/go-playground/statics)  将静态资源嵌入到go文件中，用于单个二进制编译+使用http。文件系统+符号链接。
* [statik](https://github.com/rakyll/statik)  将静态文件嵌入到Go可执行文件中。
* [templify](https://github.com/wlbr/templify)  将外部模板文件嵌入到Go代码中，以创建单个文件二进制文件。
* [vfsgen](https://github.com/shurcooL/vfsgen)  生成一个vfsdata。静态实现给定虚拟文件系统的go文件。

## 科学与数据分析

*用于科学计算和数据分析的库。 (翻译出错了? 试试 [英文版](README_EN.md#science-and-data-analysis) 吧~)*

* [assocentity](https://github.com/ndabAP/assocentity)  assocentity 返回单词到给定实体的平均距离。
* [GoStats](https://github.com/OGFris/GoStats)  GoStats是一个开放源码的GoLang库，主要用于机器学习领域的数学统计，它涵盖了大多数统计度量函数。
* [TextRank](https://github.com/DavidBelicza/TextRank)  TextRank在Golang中的实现，支持扩展特性(摘要、加权、短语提取)和多线程(goroutine)。
* [streamtools](https://github.com/nytlabs/streamtools)  通用图形工具，用于处理数据流。
* [stats](https://github.com/montanaflynn/stats)  包含Golang标准库中缺少的公共函数的统计软件包。
* [sparse](https://github.com/james-bowman/sparse)   Go 稀疏矩阵格式的线性代数支持科学和机器学习应用程序，兼容gonum矩阵库。
* [rootfinding](https://github.com/khezen/rootfinding)  二次函数求根算法库。
* [PiHex](https://github.com/claygod/PiHex)  实现了针对16进制数 Pi 的“bailee - borwein - plouffe”算法。
* [piecewiselinear](https://github.com/sgreben/piecewiselinear)  微型线性插值库。
* [pagerank](https://github.com/alixaxel/pagerank)  加权 PageRank 算法在Go中的实现。
* [orb](https://github.com/paulmach/orb)  2D几何类型，支持剪切、GeoJSON和Mapbox矢量平铺。
* [ode](https://github.com/ChristopherRabotin/ode)  常微分方程(ODE)求解器，支持扩展状态和基于信道的迭代停止条件。
* [graph](https://github.com/yourbasic/graph)  基本图形算法库。
* [gosl](https://github.com/cpmech/gosl)  提供线性代数，FFT，几何，NURBS，数值方法，概率，优化，微分方程，等等。
* [bradleyterry](https://github.com/seanhagen/bradleyterry)  为成对比较提供了一个 Bradley-Terry 模型。
* [goraph](https://github.com/gyuho/goraph)  纯Go图论库(数据结构，算法可视化)。
* [gonum/plot](https://github.com/gonum/plot)  gonum/plot提供了一个API，用于在Go中构建和绘制绘图。
* [gonum](https://github.com/gonum/gonum)  Gonum是一组用于Go编程语言的数字库。它包含用于矩阵、统计、优化等的库。
* [gohistogram](https://github.com/VividCortex/gohistogram)  数据流的近似直方图。
* [goent](https://github.com/kzahedi/goent)   Go 实现熵度量。
* [go-gt](https://github.com/ThePaw/go-gt)  用“Go”语言编写的图论算法。
* [go-dsp](https://github.com/mjibson/go-dsp)  Go数字信号处理。
* [geom](https://github.com/skelterjohn/geom)   Go 的二维几何。
* [ewma](https://github.com/VividCortex/ewma)  提供指数加权移动平均算法。
* [evaler](https://github.com/soniah/evaler)  简单的浮点算术表达式求值器。
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go)  用于机器学习和统计的数据框架(类似于panda)。
* [chart](https://github.com/vdobler/chart)  简单的图表绘制库。支持多种图形类型。
* [triangolatte](https://github.com/tchayen/triangolatte)  二维三角库。允许将线和多边形(都基于点)转换为gpu语言。

## 安全

*用于帮助您的应用程序更安全的库。 (翻译出错了? 试试 [英文版](README_EN.md#security) 吧~)*

* [acmetool](https://github.com/hlandau/acme)  ACME(让我们用自动更新加密)客户端工具。
* [Interpol](https://bitbucket.org/vahidi/interpol)  基于规则的数据生成器，用于模糊和渗透测试。
* [ssh-vault](https://github.com/ssh-vault/ssh-vault)  使用ssh密钥加密/解密。
* [simple-scrypt](https://github.com/elithrar/simple-scrypt)  Scrypt 库，具有简单、易懂的 API，同时具有内置的自动校准功能
* [secure](https://github.com/unrolled/secure)  Go 语言 HTTP 中间件，为 Go 提供了一些安全功能
* [passlib](https://github.com/hlandau/passlib)  不过时的密码哈希库。
* [nacl](https://github.com/kevinburke/nacl)   Go 实现NaCL API的集合。
* [memguard](https://github.com/awnumar/memguard)  一个用于处理内存中敏感值的纯Go库。
* [lego](https://github.com/xenolf/lego)  纯 Go ACME 客户端库及命令行工具
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword)  一个安全哈希和加密密码的偏执包。
* [acra](https://github.com/cossacklabs/acra)  网络加密代理保护基于数据库的应用程序免受数据泄漏:强选择性加密，SQL注入预防，入侵检测系统。
* [goArgonPass](https://github.com/dwin/goArgonPass)  Argon2密码散列和验证设计为与现有Python和PHP实现兼容。
* [go-yara](https://github.com/hillu/go-yara)  YARA的 Go 语言接口，号称是 “对于恶意软件研究者（以及其他人）来说是模式匹配的瑞士军刀”
* [go-generate-password](https://github.com/m1/go-generate-password)  可以在cli上或作为库使用的密码生成器。
* [certificates](https://github.com/mvmaasakkers/certificates)  用于生成tls证书的自定义工具。
* [Cameradar](https://github.com/Ullaakut/cameradar)  工具和库，以远程入侵RTSP流从监控摄像头。
* [BadActor](https://github.com/jaredfolkins/badactor)  一个驻留在内存中的，应用驱动的监控程序，受 fail2ban 的启发
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  让我们加密证书并启动TLS服务器。
* [argon2pw](https://github.com/raja/argon2pw)  使用常量时间密码比较生成Argon2密码散列。
* [sslmgr](https://github.com/adrianosela/sslmgr)  使用围绕acme/autocert的高级包装器，SSL证书变得很容易。

## 序列化

*用于二进制序列化的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#serialization) 吧~)*

* [asn1](https://github.com/PromonLogicalis/asn1)  面向golang的BER和DER编码库。
* [go-codec](https://github.com/ugorji/go)  高性能、多功能、规范化编码解码以及 rpc 库， 用于 msgpack, cbor 和 json，支持基于运行时的 OR 码生成
* [pletter](https://github.com/vimeda/pletter)  A standard way to wrap a proto message for message brokers.
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder)  用于协同 PHP session 格式数据和 PHP 序列化／反序列化函数工作的go语言库
* [mapstructure](https://github.com/mitchellh/mapstructure)  用于对原生键值对进行解码生成 Go 语言结构体
* [jsoniter](https://github.com/json-iterator/go)  高性能，100% 兼容的“encoding/json” 替代品
* [goprotobuf](https://github.com/golang/protobuf)  通过库和协议编译器插件使 Go 语言支持 Google的 protocol buffers.
* [gogoprotobuf](https://github.com/gogo/protobuf)  Go 语言的 Protocol Buffer 库。
* [go-capnproto](https://github.com/glycerine/go-capnproto)  Go 语言用的 Cap'n Proto 库及解析器
* [bambam](https://github.com/glycerine/bambam)  用于 Go 语言生成 Cap'n Proto schemas 的生成器
* [fwencoder](https://github.com/o1egl/fwencoder)  用于Go的固定宽度文件解析器(编码和解码库)。
* [fixedwidth](https://github.com/huydang284/fixedwidth)  固定宽度的文本格式(支持UTF-8)。
* [csvutil](https://github.com/jszwec/csvutil)  高性能、惯用的CSV记录编码和解码到本机Go结构。
* [colfer](https://github.com/pascaldekloe/colfer)  为Colfer二进制格式生成代码。
* [cbor](https://github.com/fxamacker/cbor)  小，安全，容易的CBOR编码和解码库。
* [binstruct](https://github.com/ghostiam/binstruct)  用于将数据映射到结构中的Golang二进制解码器。
* [bel](https://github.com/32leaves/bel)  从Go structs/interface生成TypeScript接口。对JSON RPC很有用。
* [structomap](https://github.com/tuvistavie/structomap)  用于从静态结构体简单、动态的生成键值对的库

## 服务器应用程序

* [algernon](https://github.com/xyproto/algernon)  内置支持Lua、Markdown、GCSS和Amber的HTTP/2 web服务器。
* [minio](https://github.com/minio/minio)  Minio是一个分布式对象存储服务器。
* [Trickster](https://github.com/Comcast/trickster)  HTTP反向代理缓存和时间序列加速器。
* [SFTPGo](https://github.com/drakkan/sftpgo)  功能齐全，高度可配置的SFTP服务器软件。
* [RoadRunner](https://github.com/spiral/roadrunner)  高性能PHP应用服务器，负载平衡器和进程管理器。
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  传递到负载平衡Riemann事件并/或将其转换为 Carbon。
* [psql-streamer](https://github.com/blind-oracle/psql-streamer)  从PostgreSQL到Kafka的流数据库事件。
* [nsq](http://nsq.io/)  一个实时分布式消息平台。
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus)  Nginx日志解析器和Prometheus 导出。
* [jackal](https://github.com/ortuman/jackal)  用Go编写的XMPP服务器。
* [Caddy](https://github.com/mholt/caddy)  Caddy是另一种HTTP/2 web服务器，易于配置和使用。
* [flipt](https://github.com/markphelps/flipt)  一个用Go和Vue.js编写的自包含特性标志解决方案
* [Flagr](https://github.com/checkr/flagr)  Flagr是一个开源特性标记和A/B测试服务。
* [Fider](https://github.com/getfider/fider)  Fider是一个收集和组织客户反馈的开放平台。
* [etcd](https://github.com/coreos/etcd)  为共享配置和服务发现提供高可用的键值存储。
* [dudeldu](https://github.com/krotik/dudeldu)  一个简单的SHOUTcast服务器。
* [discovery](https://github.com/Bilibili/discovery)  用于弹性中间层负载平衡和故障转移的注册表。
* [devd](https://github.com/cortesi/devd)  为开发人员提供本地web服务器。
* [consul](https://www.consul.io/)  Consul 是一个用于服务发现、监控和配置的工具
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  小型化、网络化、基于内存的键值存储


## 流处理

*用于流处理和响应式编程的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#stream-processing) 吧~)*

* [go-streams](https://github.com/reugn/go-streams)  流处理库。

## 模板引擎

*用于模板和词法分析的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#template-engines) 吧~)*

* [ace](https://github.com/yosssi/ace)  Ace 是一个 Go 语言的 HTML 模板引擎，受到了 Slim 和 Jade 的启发。 Ace 是对Gold的一种改进。
* [kasia.go](https://github.com/ziutek/kasia.go)  一个用于HTML 和其他文本文件的模板系统，使用go语言实现
* [Soy](https://github.com/robfig/soy)  Go 语言实现的谷歌闭包模板(也就是 Soy templates) ,遵循[官方规范](https://developer.google.com/closure/templates/)。
* [Razor](https://github.com/sipin/gorazor)  Go 语言的 Razor 视图引擎
* [raymond](https://github.com/aymerick/raymond)  使用 Go 语言实现的完整的 handlebars
* [quicktemplate](https://github.com/valyala/quicktemplate)  快速、强大且易用的模板引擎。将模板转化为 Go 语言并进行编译
* [pongo2](https://github.com/flosch/pongo2)  类似 DjanGo 的模板引擎
* [mustache](https://github.com/hoisie/mustache)  Go 语言实现的 Mustache 模板语言
* [maroto](https://github.com/johnfercher/maroto)  创建pdf文件的maroto方法。Maroto的灵感来自于Bootstrap并使用gofpdf。快速和简单。
* [liquid](https://github.com/osteele/liquid)  Go 语言实现的 Shopify Liquid 模板.
* [jet](https://github.com/CloudyKit/jet)  Jet模板引擎。
* [amber](https://github.com/eknkc/amber)  Amber是一个优雅的Go编程语言模板引擎，它的灵感来自HAML和Jade。
* [hero](https://github.com/shiyanhui/hero)  Hero是一个方便、快速和强大的go模板引擎。
* [goview](https://github.com/foolin/goview)  Goview是一个轻量级、极简的模板库，基于golang html/template构建Go web应用程序。
* [gospin](https://github.com/m1/gospin)  文章纺纱和spintax/纺纱语法引擎，适用于A/B，测试文本/文章片段和创建更多的自然对话。
* [gofpdf](https://github.com/jung-kurt/gofpdf)  PDF 文档生成器，支持文本，绘图和图片
* [fasttemplate](https://github.com/valyala/fasttemplate)  简单快速的模板引擎。进行模板元素替换时，速度是比[text/template](http://golang.org/pkg/text/template/)快10倍。
* [extemplate](https://github.com/dannyvankooten/extemplate)   对 html/template 进行了简单的封装，支持基于文件的模板可以利用其他模板文件进行扩展
* [ego](https://github.com/benbjohnson/ego)  轻量级模板语言，允许您在Go中编写模板。模板被翻译成Go并编译。
* [damsel](https://github.com/dskinner/damsel)  标记语言，通过css选择器实现了 html 框架 ，并可以通过 pkg html/template 等进行扩展
* [velvet](https://github.com/gobuffalo/velvet)  使用 Go 语言实现的完整的 handlebars

## 测试

*用于测试代码库和生成测试数据的库。 (翻译出错了? 试试 [英文版](README_EN.md#testing) 吧~)*

* Testing Frameworks
    * [apitest](https://apitest.dev)  基于REST的服务/HTTP处理程序的简单且可扩展的行为测试库,支持模拟外部HTTP调用和呈现序列图
    * [httpexpect](https://github.com/gavv/httpexpect)  简洁的、声明式的、易用的端到端HTTP 及 REST API 测试
    * [gogiven](https://github.com/corbym/gogiven)  类似于 YATSPEC 的Go BDD测试框架。
    * [gomatch](https://github.com/jfilipczyk/gomatch)  为针对模式测试JSON而创建的库。
    * [gomega](http://onsi.github.io/gomega/)  类似 Rspec 的 matcher/assertion 库
    * [GoSpec](https://github.com/orfjackal/gospec)  用于 Go 编程语言的bdd风格的测试框架。
    * [gospecify](https://github.com/stesla/gospecify)  支持 BDD 语法 。对于任何使用过 rspec 等库的人来说应该非常熟悉。
    * [gosuite](https://github.com/pavlo/gosuite)  轻量级测试套，为 Go1.7's Subtests 带来了setup/teardown 功能
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools)  一组包，用于增强go测试包并支持公共模式。
    * [Hamcrest](https://github.com/rdrdr/hamcrest)  用于声明性 Matcher 对象的连贯框架，当将其应用于输入值时，将产生自描述结果。
    * [jsonassert](https://github.com/kinbiko/jsonassert)  用于验证JSON有效负载已正确序列化的包。
    * [godog](https://github.com/DATA-DOG/godog)  类似 Cucumber 或 Behat 的 BDD 框架
    * [restit](https://github.com/yookoala/restit)  帮助编写 RESTful API 集成测试的 Go 语言微型框架.。
    * [schema](https://github.com/jgroeneveld/schema)  快速、简单地匹配请求和响应中使用的JSON模式的表达式。
    * [testcase](https://github.com/adamluzsi/testcase)  行为驱动开发的惯用测试框架。
    * [testfixtures](https://github.com/go-testfixtures/testfixtures)  类似 Rails 的测试工具，用于测试数据库应用
    * [Testify](https://github.com/stretchr/testify)  对标准测试包的扩展。
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  将markdown代码段转换为可测试的go代码。
    * [testsql](https://github.com/zhulongcheng/testsql)  在测试前从SQL文件生成测试数据，并在测试完成后清除数据。
    * [trial](https://github.com/jgroeneveld/trial)  无需引入太多样板，即可快速轻松地扩展断言。
    * [Tt](https://github.com/vcaesar/tt)  简单而丰富多彩的测试工具。
    * [gofight](https://github.com/appleboy/gofight)  对 Go 语言的路由框架进行 API 测试
    * [gocrest](https://github.com/corbym/gocrest)  用于 Go 断言的可组合的类似 hamcrest 的 matchers。
    * [assert](https://github.com/go-playground/assert)  基础断言库，用于对 Go 语言程序进行测试，提供了一些用于自定义断言的代码块
    * [endly](https://github.com/viant/endly)  声明性端到端功能测试。
    * [badio](https://github.com/cavaliercoder/badio)  Go 语言 testing/iotest 包的扩展。
    * [baloo](https://github.com/h2non/baloo)  表达性强、多功能的、端到端的HTTP API 测试工具
    * [biff](https://github.com/fulldump/biff)  分支测试框架，BDD兼容。
    * [charlatan](https://github.com/percolate/charlatan)  为测试生成假接口实现的工具。
    * [commander](https://github.com/SimonBaeumer/commander)  用于在windows、linux和osx上测试cli应用程序的工具。
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy)  测试框架的简单快照测试插件。
    * [dbcleaner](https://github.com/khaiql/dbcleaner)  清空数据库用于测试，受到database_cleaner 的启发
    * [dsunit](https://github.com/viant/dsunit)  用于SQL、NoSQL、结构化文件的数据存储测试。
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres)  作为另一个Go应用程序或测试的一部分，在Linux、OSX或Windows上本地运行一个真正的Postgres数据库。
    * [flute](https://github.com/suzuki-shunsuke/flute)  HTTP客户端测试框架。
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD 风格的测试框架，具有 web 界面和计时刷新功能
    * [frisby](https://github.com/verdverm/frisby)  REST API测试框架。
    * [ginkgo](http://onsi.github.io/ginkgo/)  Go的BDD测试框架。
    * [go-carpet](https://github.com/msoap/go-carpet)  在终端中查看测试覆盖率的工具。
    * [go-cmp](https://github.com/google/go-cmp)  用于比较测试中的Go值的包。
    * [go-mutesting](https://github.com/zimmski/go-mutesting)  变异测试的Go源代码。
    * [go-testdeep](https://github.com/maxatome/go-testdeep)  极具灵活性的golang深度比较，扩展了go测试包。
    * [go-vcr](https://github.com/dnaeon/go-vcr)  记录并回放HTTP交互，以便进行快速、确定和准确的测试。
    * [goblin](https://github.com/franela/goblin)  类似Mocha的测试框架。
    * [gocheck](http://labix.org/gocheck)  更加高级的测试框架，用于替换 Gotest
    * [wstest](https://github.com/posener/wstest)  用于单元测试Websocket http.Handler的Websocket客户机。

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)  用于生成自包含 mock 对象的工具
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)  Mock SQL ，用于测试数据库交互
    * [go-txdb](https://github.com/DATA-DOG/go-txdb)  基于单事务的数据库驱动，主要用于测试目的
    * [gock](https://github.com/h2non/gock)  多功能、易用 HTTP mock
    * [gomock](https://github.com/golang/mock)  用于Go编程语言的mock框架。
    * [govcr](https://github.com/seborama/govcr)  HTTP mock : 离线测试时记录和重放浏览器的动作
    * [hoverfly](https://github.com/SpectoLabs/hoverfly)  使用可扩展中间件和易于使用的CLI记录和模拟REST/SOAP api的HTTP(S)代理。
    * [httpmock](https://github.com/jarcoal/httpmock)  轻松模拟来自外部资源的HTTP响应。
    * [minimock](https://github.com/gojuno/minimock)  Go接口的模拟生成器。
    * [mockhttp](https://github.com/tv42/mockhttp)  Go http.ResponseWriter的模拟对象。
    * [timex](https://github.com/cabify/timex)  一个测试友好的本地“时间”包的替代品。

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz)  随机测试系统。
    * [gofuzz](https://github.com/google/gofuzz)  用于生成随机值来初始化 Go 语言对象的库
    * [Tavor](https://github.com/zimmski/tavor)  通用模糊测试框架

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp)  用于Chrome调试协议的类型安全绑定，可与实现该协议的浏览器或其他调试目标一起使用。
    * [chromedp](https://github.com/knq/chromedp)  用于驱动和测试 Chrome, Safari, Edge, Android Webviews, 以及其他支持 Chrome 调试协议的产品
    * [ggr](https://github.com/aerokube/ggr)  一个轻量级服务器，可以将 Selenium Wedriver 的请求路由或代理到多个 Selenium hubs
    * [selenoid](https://github.com/aerokube/selenoid)  Selenium hub 服务器的替代品，在容器中启动浏览器

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint)  为Golang实现[failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail)。

## 文本处理

*用于解析和操作文本的库。 (翻译出错了? 试试 [英文版](README_EN.md#text-processing) 吧~)*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align)  对文本进行对齐的通用应用程序。
    * [guesslanguage](https://github.com/endeveit/guesslanguage)  通过一个 unicode 文本来猜测该文本使用的语言
    * [gographviz](https://github.com/awalterschulze/gographviz)  解析Graphviz DOT语言。
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  格式化二进制为字符串。
    * [gonameparts](https://github.com/polera/gonameparts)  将人名解析为单独的名称部分。
    * [goq](https://github.com/andrewstuart/goq)   声明式 HTML 编组，使用结构标签和 jQuery 语法 (使用 GoQuery).
    * [GoQuery](https://github.com/PuerkitoBio/goquery)  GoQuery 为 Go 语言带来了一组类似 jQuery 的语法和功能
    * [goregen](https://github.com/zach-klippenstein/goregen)  根据正则表达式生成随机字符串
    * [goribot](https://github.com/zhshch2002/goribot)  一个简单的golang spider/抓取框架，构建一个三行蜘蛛。
    * [gotext](https://github.com/leonelquinteros/gotext)  GNU gettext 工具
    * [htmlquery](https://github.com/antchfx/htmlquery)  用于HTML的XPath查询包，允许您通过XPath表达式从HTML文档中提取数据或求值。
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width)  用于Go的零宽度字符检测和删除。
    * [inject](https://github.com/facebookgo/inject)  包注入提供了一个基于反射的注入器。
    * [ltsv](https://github.com/Wing924/ltsv)  用于Go的高性能[LTSV(标签为Tab Separeted Value)](http://ltsv.org/)阅读器。
    * [mxj](https://github.com/clbanning/mxj)  将XML编码/解码为JSON或map[string]接口{};使用点符号路径和通配符提取值。替换x2j和j2x包。
    * [podcast](https://github.com/eduncan911/podcast)  在Golang兼容iTunes和RSS 2.0播客生成器
    * [sdp](https://github.com/gortc/sdp)  SDP:会话描述协议[[RFC 4566](https://tools.ietf.org/html/rfc4566)]。
    * [sh](https://github.com/mvdan/sh)  Shell解析器和格式化工具。
    * [slug](https://github.com/gosimple/slug)  URL 友好的 slug 化工具，支持多种语言
    * [Slugify](https://github.com/avelino/slugify)  字符串 slug 化的工具。
    * [syndfeed](https://github.com/zhengchun/syndfeed)  Atom 1.0和RSS 2.0的联合提要。
    * [gofeed](https://github.com/mmcdole/gofeed)  在Go中解析RSS和Atom feeds。
    * [go-vcard](https://github.com/emersion/go-vcard)  解析和格式化vCard。
    * [allot](https://github.com/sbstjn/allot)  用于CLI工具和机器人的占位符和通配符文本解析。
    * [doi](https://github.com/hscells/doi)  文档对象标识符(doi)解析器。
    * [bbConvert](https://github.com/CalebQ42/bbConvert)  将bbCode转换为HTML，使您可以添加对自定义bbCode标记的支持。
    * [blackfriday](https://github.com/russross/blackfriday)  Markdown 解析器
    * [bluemonday](https://github.com/microcosm-cc/bluemonday)  HTML 清理工具
    * [codetree](https://github.com/aerogo/codetree)  解析缩进代码(python、pixy、scarlet等)并返回树结构。
    * [colly](https://github.com/asciimoo/colly)  快速和优雅的 Scraping 框架。
    * [commonregex](https://github.com/mingrammer/commonregex)  一组用于Go的公共正则表达式。
    * [dataflowkit](https://github.com/slotix/dataflowkit)  Web抓取框架将网站转换为结构化数据。
    * [did](https://github.com/ockam-network/did)  DID(分散标识符)解析器和Stringer。
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go)  Editorconfig文件解析器和Go操作器。
    * [go-toml](https://github.com/pelletier/go-toml)  使用带有查询支持和方便的cli工具的TOML格式库。
    * [enca](https://github.com/endeveit/enca)  [libenca](http://cihar.com/software/enca/)的最小cgo绑定。
    * [encdec](https://github.com/mickep76/encdec)  软件包为编码器和解码器提供了通用接口。
    * [genex](https://github.com/alixaxel/genex)  将正则表达式计数并展开为所有匹配的字符串。
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub 风格的 Markdown 渲染器 (使用 blackfriday) ，支持代码块高亮以及可点击的锚点
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth)  固定宽度的文本格式(带反射的编码器/解码器)。
    * [go-humanize](https://github.com/dustin/go-humanize)  格式化程序，用于将时间、数字和内存大小转换为可读格式。
    * [go-nmea](https://github.com/adrianmo/go-nmea)  用于Go语言的NMEA解析器库。
    * [go-runewidth](https://github.com/mattn/go-runewidth)  函数获取字符或字符串的固定宽度。
    * [go-slugify](https://github.com/mozillazg/go-slugify)  生成漂亮的固定链接地址（slug），支持多种语言
    * [toml](https://github.com/BurntSushi/toml)  TOML配置格式(带反射的编码器/解码器)。
* Utility
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself)  一个基于 sanitization 的 Go 敏感词过滤器。
    * [gotabulate](https://github.com/bndr/gotabulate)  使用 Go 语言简单、美观的打印表格数据
    * [kace](https://github.com/codemodus/kace)  通用大小写转换工具
    * [parseargs-go](https://github.com/nproc/parseargs-go)  字符串参数解析器，能够理解引用及反斜杠。
    * [parth](https://github.com/codemodus/parth)  URL路径分割解析。
    * [radix](https://github.com/yourbasic/radix)  快速字符串排序算法。
    * [Tagify](https://github.com/zoomio/tagify)  从给定源生成一组标记。
	* [textwrap](https://github.com/isbm/textwrap)  从Python实现“textwrap”模块。
    * [TySug](https://github.com/Dynom/TySug) **star:4** 关于键盘布局的其他建议。   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
    * [xj2go](https://github.com/stackerzzq/xj2go)  将xml或json转换为struct。
    * [xurls](https://github.com/mvdan/xurls)  从文本中提取url。

## 第三方api

*用于访问第三方api的库。 (翻译出错了? 试试 [英文版](README_EN.md#third-party-apis) 吧~)*

* [go-hacknews](https://github.com/PaulRosset/go-hacknews)  HackerNews API的小型Go客户端。
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans)   SPTrans Olho Vivo API 的客户端。
* [go-here](https://github.com/abdullahselek/go-here)  转到基于这里位置的api的客户端库。
* [go-marathon](https://github.com/gambol99/go-marathon)   用于和 Mesosphere's Marathon PAAS 交互的 Go 语言库
* [go-chronos](https://github.com/axelspringer/go-chronos)  用于与[Chronos](https://mesos.github.io/chronos/)作业调度程序进行交互的Go库
* [githubql](https://github.com/shurcooL/githubql)  访问GitHub GraphQL API v4的库。
* [github](https://github.com/google/go-github)  访问GitHub REST API v3的库。
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:341** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist)  [MyAnimeList API](http://myanimelist.net/modules.php?go=api)的客户端库。
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection)  去模块工作与[邮递员收集](https://learning.getpostman.com/docs/postman/collections/creating-collections/)(兼容失眠)。
* [go-sophos](https://github.com/esurdam/go-sophos)  无任何依赖的[Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/pdfs/documentation/utmonaws/sophos-ut-restful-api.pdf?la=en)客户端
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph 发布平台 API 客户端。
* [go-jira](https://github.com/andygrunwald/go-jira)   Go [Atlassian JIRA](https://www.atlassian.com/software/jira)的客户端库
* [go-trending](https://github.com/andygrunwald/go-trending)  在Github上访问[trends repository](https://github.com/trends)和[developers](https://github.com/trending/developers)的库。
* [go-twitch](https://github.com/knspriggs/go-twitch)   Twitch v3 API 的客户端。
* [go-twitter](https://github.com/dghubble/go-twitter)   Twitter v1.1 api 的客户端.
* [go-unsplash](https://github.com/hbagdi/go-unsplash)  使用[Unsplash.com](https://unsplash.com) API的客户端库。
* [go-xkcd](https://github.com/nishanths/go-xkcd)   xkcd API 的客户端。
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb)  电影数据库API v3的Golang包装器。
* [go-imgur](https://github.com/koffeinsource/go-imgur)  Go [imgur](https://imgur.com)的客户端库
* [wit-go](https://github.com/wit-ai/wit-go) **star:60** wit.ai HTTP API 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [ynab](https://github.com/brunomvsouza/ynab.go)   YNAB API 的 Go 语言封装。
* [anaconda](https://github.com/ChimeraCoder/anaconda)   Twitter 1.1 API 的 go 语言客户端
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:40** [Amazon Product Advertising API](https://program.amazon.com/gp/advertising/api/detail/main.html)的客户端库。   [![最近一年没有更新][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [aws-sdk-go](https://github.com/aws/aws-sdk-go)  AWS 提供的官方go语言 SDK
* [clarifai](https://github.com/samuelcouch/clarifai)  Clarifai API的客户端。
* [circleci](https://github.com/jszwedko/go-circleci)  CircleCI的API的客户端
* [cachet](https://github.com/andygrunwald/cachet)  使用客户端库[Cachet(开源状态页系统)](https://cachethq.io/)。
* [brewerydb](https://github.com/naegelejd/brewerydb)  用于访问 BreweryDB API的 Go 语言库
* [simples3](https://github.com/rhnvrm/simples3) **star:36** 使用REST和用Go编写的V4签名的AWS S3库非常简单。   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [slack](https://github.com/nlopes/slack)  Slack API。
* [smite](https://github.com/sergiotapia/smitego)  对 Smite game API 的封装。
* [spotify](https://github.com/rapito/go-spotify)   用于接入 Spotify WEB API 的 Go 语言库
* [steam](https://github.com/sostronk/go-steam)   用于与Steam服务器进行交互的库。
* [stripe](https://github.com/stripe/stripe-go)   Stripe API 的 Go 语言客户端
* [gostorm](https://github.com/jsgilmore/gostorm)  GoStorm是一个Go库，它实现了编写Storm spout和bolt所需的通信协议，这些协议用于与Storm shell通信。
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api)  [谷歌G Suite Email Audit API](https://developer.google.com/admin-sdk/email-audit/) 的客户端。
* [hipchat](https://github.com/andybons/hipchat)  这个项目为Hipchat API实现了一个golang客户端库。
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat)  一个用于通过XMPP与HipChat通信的golang包。
* [igdb](https://github.com/Henry-Sarabia/igdb)  [Internet Game Database API](https://api.igdb.com/) 客户端。
* [gosip](https://github.com/koltyakov/gosip)  去客户端库SharePoint API。
* [google](https://github.com/google/google-api-go-client)  为Go自动生成谷歌api。
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang)  谷歌云api Go 客户端库。
* [google-analytics](https://github.com/chonthu/go-google-analytics)  简单的包装，方便的谷歌分析报告。
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz)   Go MusicBrainz WS2客户端库。
* [gomalshare](https://github.com/MonaxGT/gomalshare)  Go library MalShare API [malshare.com](http://www.malshare.com/)
* [golyrics](https://github.com/mamal72/golyrics) **star:34** Golyrics是一个从Wikia网站获取音乐歌词数据的Go库。   [![最近一年没有更新][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [gami](https://github.com/bit4bit/gami) **star:27**  Asterisk Manager Interface 的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![归档项目][Archived]](https://github.com/bit4bit/gami)
* [gcm](https://github.com/Aorioli/gcm)   Google Cloud Messaging 库
* [shopify](https://github.com/rapito/go-shopify) **star:19** 一个用于通过 Shopify API 进行增删改查的 Go 语言库。   [![最近一年没有更新][Yellow]](https://github.com/rapito/go-shopify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [codeship-go](https://github.com/codeship/codeship-go) **star:17**  Codeship API v2的客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang)   TripAdvisor API 的 Go 语言封装。
* [translate](https://github.com/poorny/translate)   Go 在线翻译包。
* [webhooks](https://github.com/go-playground/webhooks)  GitHub 和 Bitbucket 的Webhook接收器。
* [vl-go](https://github.com/verifid/vl-go)  使用客户端库中的VerifID身份验证层API。
* [uptimerobot](https://github.com/bitfield/uptimerobot)  运行时Robot v2 API 的 Go 语言封装和命令行客户端。
* [twitter-scraper](https://github.com/n0madic/twitter-scraper)  刮Twitter前端API没有认证和限制。
* [tumblr](https://github.com/mattcunningham/gumblr)   Tumblr v2 API 的 Go 语言封装。
* [textbelt](https://github.com/dietsche/textbelt) **star:16** textbelt.com txt messaging API 的go语言客户端。   [![最近一年没有更新][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [Trello](https://github.com/adlio/trello)   Trello API的 Go 语言封装。
* [ethrpc](https://github.com/onrik/ethrpc)   Ethereum JSON RPC API的客户端。
* [facebook](https://github.com/huandu/facebook)  支持 Facebook Graph API 的库
* [fcm](https://github.com/maddevsio/fcm)   Firebase Cloud Messaging 的 Go 语言库
* [gads](https://github.com/emiddleton/gads)   Google Adwords 非官方 API
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Coinpaprika API的客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [discordgo](https://github.com/bwmarrin/discordgo)   Discord Chat API的客户端。
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:9** 使用[LastPass](https://www.lastpass.com/) API的客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** 用于接入 statdns.com API 的库——RRDA API。通过HTTP协议进行 DNS查询。   [![最近一年没有更新][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [zooz](https://github.com/gojuno/go-zooz) **star:5**  Zooz API 的 Go 语言客户端。   [![最近一年没有更新][Yellow]](https://github.com/gojuno/go-zooz)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [libgoffi](https://github.com/clevabit/libgoffi) **star:1** 库适配器工具箱，用于本地[libffi](http://sourceware.org/libffi/)集成   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)
* [pushover](https://github.com/gregdel/pushover)   Go 包装的 Pushover API。
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk)  Playlyfe Rest API Go SDK。
* [paypal](https://github.com/logpacker/PayPal-Go-SDK)  PayPal支付API的包装器。
* [patreon-go](https://github.com/mxpv/patreon-go)   Go Patreon API库。
* [mixpanel](https://github.com/dukex/mixpanel)  Mixpanel是一个库，用于跟踪事件并将Mixpanel概要文件更新从go应用程序发送到Mixpanel。
* [minio-go](https://github.com/minio/minio-go)  用于Amazon S3兼容云存储的Minio Go库。
* [megos](https://github.com/andygrunwald/megos)  用于访问[Apache Mesos](http://mesos.apache.org/)集群的客户端库。
* [Medium](https://github.com/Medium/medium-sdk-go)  Medium的OAuth2 API 客户端。

## 公用事业公司

*可以让你的生活变得更简单的实用工具.。 (翻译出错了? 试试 [英文版](README_EN.md#utilities) 吧~)*

* [fzf](https://github.com/junegunn/fzf) **star:27197** 用Go编写的命令行模糊查找器。   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![最近一周有更新][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [ctop](https://github.com/bcicen/ctop) **star:9499** [Top-like](http://ctop.sh)接口(例如htop)， 用于容器数据收集。   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [wuzz](https://github.com/asciimoo/wuzz) **star:8550** 用于HTTP检查的交互式cli工具。   [![star > 2000][Awesome]](https://github.com/asciimoo/wuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/wuzz)
* [sqlx](https://github.com/jmoiron/sqlx) **star:7842** 为内建的数据库/sql 软件包提供一组扩展。   [![star > 2000][Awesome]](https://github.com/jmoiron/sqlx)   [![godoc][GoDoc]](https://godoc.org/github.com/jmoiron/sqlx)
* [peco](https://github.com/peco/peco) **star:5766** 简单的交互过滤工具。   [![star > 2000][Awesome]](https://github.com/peco/peco)   [![godoc][GoDoc]](https://godoc.org/github.com/peco/peco)
* [util](https://github.com/shomali11/util)  有用实用函数的集合。(字符串，并发，操作，…)
* [usql](https://github.com/knq/usql) **star:5637** usql 是一个通用的命令行接口，用于操作 sql 数据库。   [![star > 2000][Awesome]](https://github.com/knq/usql)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/usql)
* [goreporter](https://github.com/wgliang/goreporter)  进行代码静态分析，单元测试，代码检视并生成代码质量报告的工具
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:5196** 尽可能快速的发布 Go 语言二进制文件。   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![最近一周有更新][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [godropbox](https://github.com/dropbox/godropbox) **star:3838** 用于编写 Go 语言服务／应用的库，来自 Dropbox.。   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![最近一周有更新][Green]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [gohper](https://github.com/cosiner/gohper)  多种能够帮助你进行软件开发的工具和模块。
* [repeat](https://github.com/ssgreg/repeat)  执行不同的后 backoff 策略，这对重新尝试操作和心跳非常有用。
* [realize](https://github.com/tockins/realize) **star:3520** Go 语言构建系统，可以监控文件变化并重新加载。运行，构建，监控文件并支持自定义路径。   [![star > 2000][Awesome]](https://github.com/tockins/realize)   [![godoc][GoDoc]](https://godoc.org/github.com/tockins/realize)
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2356** 实现 Hystrix 风格的、程序员预定义的 fallback 机制（熔断。   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [Task](https://github.com/go-task/task) **star:2243** 简单的“Go”的选择。   [![star > 2000][Awesome]](https://github.com/go-task/task)   [![最近一周有更新][Green]](https://github.com/go-task/task)   [![godoc][GoDoc]](https://godoc.org/github.com/go-task/task)
* [minify](https://github.com/tdewolff/minify) **star:2022** 用于HTML、CSS、JS、XML、JSON和SVG文件格式的快速缩小器。   [![star > 2000][Awesome]](https://github.com/tdewolff/minify)   [![最近一周有更新][Green]](https://github.com/tdewolff/minify)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/minify)
* [go-health](https://github.com/Talento90/go-health)  健康包简化了向服务中添加健康检查的方式。
* [go-funk](https://github.com/thoas/go-funk) **star:1596** 现代 Go 语言工具库，提供了很多有用的工具 (map, find, contains, filter, chunk, reverse, ...)   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [Storm](https://github.com/asdine/storm) **star:1489** 一个简单又强大的用于 BoltDB 的工具   [![godoc][GoDoc]](https://godoc.org/github.com/asdine/storm)
* [mmake](https://github.com/tj/mmake) **star:1483** 现代 Make 工具   [![godoc][GoDoc]](https://godoc.org/github.com/tj/mmake)
* [mc](https://github.com/minio/mc) **star:1361** Minio Client 提供了一组工具，用于操作 Amazon S3 兼容云存储和文件系统。   [![最近一周有更新][Green]](https://github.com/minio/mc)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/mc)
* [mole](https://github.com/davrodpin/mole) **star:1352** cli应用程序可以轻松创建ssh隧道。   [![godoc][GoDoc]](https://godoc.org/github.com/davrodpin/mole)
* [filetype](https://github.com/h2non/filetype) **star:1050** 通过数字签名来推测文件类型。   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [mergo](https://github.com/imdario/mergo) **star:1023** 用于将结构体和map合并进 Go 语言的工具。对于配置默认值，避免杂乱的if语句很有帮助。   [![godoc][GoDoc]](https://godoc.org/github.com/imdario/mergo)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:857** 接通断路器。   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [jump](https://github.com/gsamokovarov/jump) **star:801** 通过学习你的习惯，可以帮助你更快地导航。   [![godoc][GoDoc]](https://godoc.org/github.com/gsamokovarov/jump)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:774** git-time-metric - 。   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [immortal](https://github.com/immortal/immortal) **star:633** \*nix 跨平台 (与操作系统无关的)监控程序。   [![godoc][GoDoc]](https://godoc.org/github.com/immortal/immortal)
* [htcat](https://github.com/htcat/htcat) **star:507** 并行及流水线的 HTTP GET 工具。   [![最近一年没有更新][Yellow]](https://github.com/htcat/htcat)   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [hub](https://github.com/github/hub)  封装了 git 命令，提供了额外的功能用于在终端中和 Github 进行交互。
* [go-dry](https://github.com/ungerik/go-dry) **star:446** DRY(don't repeat yourself)库。   [![最近一年没有更新][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [gopencils](https://github.com/bndr/gopencils) **star:430** 小而简单的包，可以轻松地使用REST api。   [![最近一年没有更新][Yellow]](https://github.com/bndr/gopencils)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [rerate](https://github.com/abo/rerate)  基于 Redis 的速率计数器和限速器
* [request](https://github.com/mozillazg/request) **star:370** Go 语言版的 HTTP Requests for Humans™.。   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/request)
* [koazee](https://github.com/wesovilabs/koazee) **star:362** 库的灵感来自于延迟计算和函数式编程，从而减少了使用数组的麻烦。   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/koazee)
* [ergo](https://github.com/cristianoliveira/ergo) **star:359** 管理运行在不同端口上的多个本地服务变得很容易。   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [go-rate](https://github.com/beefsack/go-rate) **star:299**  Go 限速器。   [![最近一年没有更新][Yellow]](https://github.com/beefsack/go-rate)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [clockwork](https://github.com/jonboulle/clockwork) **star:254** 一个简单的假 clock 。   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:245** 结构体拷贝   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:209** 用于基于神奇数字的MIME类型检测的包。   [![最近一周有更新][Green]](https://github.com/gabriel-vasile/mimetype)   [![godoc][GoDoc]](https://godoc.org/github.com/gabriel-vasile/mimetype)
* [serve](https://github.com/syntaqx/serve) **star:206** 任何您需要的静态http服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/syntaqx/serve)
* [gubrak](https://github.com/novalagung/gubrak) **star:205** 带有语法糖的Golang实用工具，就像lodash。   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [retry](https://github.com/kamilsk/retry) **star:204** 基于上下文的功能机制，反复执行命令直到成功。   [![最近一周有更新][Green]](https://github.com/kamilsk/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/retry)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:192** Go 语言全局事件触发器，通过 id 和触发器，在程序的任何地方注册事件。   [![最近一年没有更新][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [gotenv](https://github.com/subosito/gotenv) **star:168** 从 `.env` 或者任何 `io.Reader`。   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:163** Go:generate 工具，用于构建 Go 语言插件(1.8 only)，并对导出的符号进行包装。   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [rerun](https://github.com/ivpusic/rerun) **star:155** 当源代码发生更改时，重新编译和重新运行go应用程序。   [![最近一年没有更新][Yellow]](https://github.com/ivpusic/rerun)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/rerun)
* [moldova](https://github.com/StabbyCutyou/moldova) **star:149** 基于输入目标生成随机数据的工具   [![最近一年没有更新][Yellow]](https://github.com/StabbyCutyou/moldova)   [![godoc][GoDoc]](https://godoc.org/github.com/StabbyCutyou/moldova)
* [Death](https://github.com/vrecan/death) **star:148** 利用信号管理应用程序的关闭。   [![最近一年没有更新][Yellow]](https://github.com/vrecan/death)   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [scan](https://github.com/blockloop/scan)  扫描golang的sql。行直接指向结构、片或基本类型。
* [robustly](https://github.com/VividCortex/robustly) **star:140** 有弹性的执行函数，遇到错误时捕获并重新运行。   [![最近一年没有更新][Yellow]](https://github.com/VividCortex/robustly)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/robustly)
* [ugo](https://github.com/alxrm/ugo)  uGo 是一个切片工具箱，有着和 Go 语言一致的语法法。
* [toolbox](https://github.com/viant/toolbox) **star:124** 切片, map, multimap, 结构体, 函数,数据转换工具。服务路由，宏求值和标记器。   [![最近一周有更新][Green]](https://github.com/viant/toolbox)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/toolbox)
* [circuit](https://github.com/cep21/circuit)  一个高效和功能齐全的 类似 Hystrix Go 实现断路器模式。
* [chyle](https://github.com/antham/chyle) **star:117** 使用具有多种配置可能性的git存储库生成变更日志。   [![最近一周有更新][Green]](https://github.com/antham/chyle)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:116** 用Go编写的XML站点地图生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:105** LiveReload 服务器。   [![最近一年没有更新][Yellow]](https://github.com/jaschaephraim/lrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/jaschaephraim/lrserver)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:90** 纯Go bsdiff和bspatch库和CLI工具。   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [pm](https://github.com/VividCortex/pm) **star:77** 进程(即goroutine)管理器与HTTP API。   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/pm)
* [xferspdy](https://github.com/monmohan/xferspdy) **star:72** Xferspdy在golang中提供二进制diff和补丁库。   [![最近一年没有更新][Yellow]](https://github.com/monmohan/xferspdy)   [![godoc][GoDoc]](https://godoc.org/github.com/monmohan/xferspdy)
* [UNIS](https://github.com/esemplastic/unis) **star:69** Go 语言字符串处理函数的通用架构 。   [![最近一年没有更新][Yellow]](https://github.com/esemplastic/unis)   [![godoc][GoDoc]](https://godoc.org/github.com/esemplastic/unis)
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:69** 数据库客户端，用于主-从 (或主-主) 数据库，集成了简单的、轻量级的轮询调度负载均衡。   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/mssqlx)
* [myhttp](https://github.com/inancgumus/myhttp)  简单的API，使HTTP GET请求与超时支持。
* [multitick](https://github.com/VividCortex/multitick) **star:62** 用于 aligned tickers 的多路复用   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/multitick)
* [minquery](https://github.com/icza/minquery) **star:53** MongoDB / mgo.v2, 支持高效分页查询(用于继续列出我们停止的文档的游标)。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/minquery)
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:52** 纯粹 Go 超性能MIME嗅探库/实用程序。   [![最近一年没有更新][Yellow]](https://github.com/zRedShift/mimemagic)   [![godoc][GoDoc]](https://godoc.org/github.com/zRedShift/mimemagic)
* [handy](https://github.com/miguelpragier/handy) **star:50** 许多实用程序和帮助程序，如字符串处理程序/格式化程序和验证器。   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [golog](https://github.com/mlimaloureiro/golog) **star:48** 简单、轻量级的命令后工具，用于对你的计划任务进行跟踪。   [![最近一年没有更新][Yellow]](https://github.com/mlimaloureiro/golog)   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** 解析你 Go 语言代码中的 TODOs（待办事项）。   [![最近一年没有更新][Yellow]](https://github.com/asticode/go-astitodo)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:45** conseilSeaweedFS 客户端，几乎具有全部的特性。   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [gostrutils](https://github.com/ik5/gostrutils)  字符串操作和转换函数的集合。
* [goreadability](https://github.com/philipjkim/goreadability) **star:45** 网页摘要提取器使用Facebook开放图形和arc90的可读性。   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [godaemon](https://github.com/VividCortex/godaemon)  用于编写守护进程的工具
* [goback](https://github.com/carlescere/goback) **star:42**  一个 Go 语言的简单的指数补偿包。   [![最近一年没有更新][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [gaper](https://github.com/maxcnunes/gaper) **star:42** 当Go项目崩溃或一些人看到文件更改时，构建并重新启动该项目。   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:41** 不需要编写任何汇编代码就能使用 x86 SIMD。   [![最近一年没有更新][Yellow]](https://github.com/mengzhuo/intrinsic)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/intrinsic)   [![归档项目][Archived]](https://github.com/mengzhuo/intrinsic)
* [golarm](https://github.com/msempere/golarm) **star:38** 告警（支持系统事件）。   [![最近一年没有更新][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [pgo](https://github.com/arthurkushman/pgo) **star:36** 用于PHP社区的 Convenient 函数。   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/pgo)
* [retry](https://github.com/thedevsaddam/retry) **star:36** 简单易用的重试机制包，为 Go 。   [![最近一年没有更新][Yellow]](https://github.com/thedevsaddam/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/retry)
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:34** 对 Go 来说，重试变得简单而容易。   [![最近一年没有更新][Yellow]](https://github.com/rafaeljesus/retry-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/retry-go)
* [beyond](https://github.com/wesovilabs/beyond) **star:34** Go工具将带你进入AOP的世界!   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [gpath](https://github.com/tenntenn/gpath) **star:33**  用于简化结构体域访问的库。   [![最近一年没有更新][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [dbt](https://github.com/nikogura/dbt) **star:25** 用于从中心可信存储库运行自更新签名二进制文件的框架。   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [generate](https://github.com/go-playground/generate) **star:22** 针对一个路径或环境变量，递归的执行 Go generate，可以通过正则表达式来进行过滤。   [![最近一年没有更新][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:21** 一个小巧的 Go 语言库用于生成占位图片。   [![最近一年没有更新][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [countries](https://github.com/biter777/countries) **star:21** 完整实现ISO-3166-1、ISO-4217、ITU-T E.164、Unicode CLDR和IANA ccTLD标准。   [![godoc][GoDoc]](https://godoc.org/github.com/biter777/countries)
* [evaluator](https://github.com/nullne/evaluator) **star:20** 基于 s-expression。它很简单，很容易扩展。   [![最近一年没有更新][Yellow]](https://github.com/nullne/evaluator)   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [filter](https://github.com/gookit/filter) **star:20** 提供Go数据的过滤、清理和转换。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:19** 用于在osx、windows和linux上执行shell命令的库。   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails)  打包处理问题细节。
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:16**  用于将结构体编码进 http 头的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [slicer](https://github.com/leaanthony/slicer) **star:15** 使处理切片更容易。   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/slicer)
* [filler](https://github.com/yaronsumel/filler) **star:15** 使用“fill”标签填充结构的小工具。   [![最近一年没有更新][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** 编译时控制的日志，让你的 release 包变得更小而不需移除 debug 调用。   [![最近一年没有更新][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [spinner](https://github.com/briandowns/spinner)   一个 Go 语言软件包，提供多种选项，方便在终端中创建加载动画。
* [structs](https://github.com/PumpkinSeed/structs) **star:14** 简单来讲就是 "Make" 的替代品。   [![最近一年没有更新][Yellow]](https://github.com/PumpkinSeed/structs)   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/structs)
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match)  模式匹配图书馆。
* [onecache](https://github.com/adelowo/onecache)  支持多个后端存储(Redis、Memcached、文件系统等)的缓存库。
* [olaf](https://github.com/btnguyen2k/olaf)  Twitter Snowflake 在Go中实现。
* [okrun](https://github.com/xta/okrun) **star:14**  Go 运行错误 steamroller。   [![最近一年没有更新][Yellow]](https://github.com/xta/okrun)   [![godoc][GoDoc]](https://godoc.org/github.com/xta/okrun)
* [panicparse](https://github.com/maruel/panicparse)  将类似的协程分组并对调用栈进行着色
* [ghokin](https://github.com/antham/ghokin) **star:13** 没有外部依赖的gherkin (cucumber, behat…)并行格式化程序。   [![最近一周有更新][Green]](https://github.com/antham/ghokin)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [rest-go](https://github.com/edermanoel94/rest-go) **star:12** 一个包，它为使用rest api提供了许多有用的方法。   [![最近一周有更新][Green]](https://github.com/edermanoel94/rest-go)   [![godoc][GoDoc]](https://godoc.org/github.com/edermanoel94/rest-go)
* [retry](https://github.com/shafreeck/retry) **star:11** 一个相当简单的库，以确保您的工作可以完成。   [![godoc][GoDoc]](https://godoc.org/github.com/shafreeck/retry)
* [ctxutil](https://github.com/posener/ctxutil) **star:11** 上下文工具。   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:11** 一个用于Go的MIME类型嗅探器。   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/mimesniffer)
* [command](https://github.com/txgruppi/command) **star:9** 命令模式，支持线程安全的串行、并行调度。   [![最近一年没有更新][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [backscanner](https://github.com/icza/backscanner) **star:9** 类似 bufio 的扫描器，但它以相反的顺序读取和返回行，从给定的位置开始，然后返回。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [limiters](https://github.com/mennanov/limiters) **star:9** Golang中的分布式应用程序速率限制器，具有可配置的后端和分布式锁。   [![godoc][GoDoc]](https://godoc.org/github.com/mennanov/limiters)
* [copy-pasta](https://github.com/jutkko/copy-pasta)  通用多工作站剪切板，使用类似 S3 的后端作为存储。
* [tome](https://github.com/cyruzin/tome) **star:9** Tome设计用于对简单的RESTful api进行分页。   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/tome)
* [shutdown](https://github.com/ztrue/shutdown) **star:8** ' os的应用程序关闭挂钩。信号的处理。   [![最近一年没有更新][Yellow]](https://github.com/ztrue/shutdown)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/shutdown)
* [silk](https://github.com/chrispassas/silk)  阅读silk netflow文件。
* [slice](https://github.com/psampaz/slice)  用于普通Go切片操作的类型安全函数。
* [jsend](https://github.com/clevergo/jsend) **star:8** 在Go中编写JSend的实现。   [![最近一周有更新][Green]](https://github.com/clevergo/jsend)   [![godoc][GoDoc]](https://godoc.org/github.com/clevergo/jsend)
* [retry](https://github.com/percolate/retry) **star:6** 一个简单但高度可配置的Go重试包。
* [apm](https://github.com/topfreegames/apm)  Go 语言进程管理工具具有HTTP API.。
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:5** 基本类型之间的片转换。   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
* [boilr](https://github.com/tmrts/boilr)  非常快的CLI工具，用于从样板模板创建项目。
* [blank](https://github.com/Henry-Sarabia/blank) **star:4** 验证或删除字符串中的空白。   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [ptr](https://github.com/gotidy/ptr) **star:3** 提供函数的包，用于从基本类型的常量简化指针的创建。   [![godoc][GoDoc]](https://godoc.org/github.com/gotidy/ptr)
* [delve](https://github.com/derekparker/delve) **star:3** Go 语言调试器   [![最近一周有更新][Green]](https://github.com/derekparker/delve)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [rclient](https://github.com/zpatrick/rclient)  可读、灵活、易于使用的REST api客户端。

## UUID

*用于处理uuid的库。 (翻译出错了? 试试 [英文版](README_EN.md#uuid) 吧~)*

* [ulid](https://github.com/oklog/ulid) **star:1828** 实现了ULID(普遍唯一的词典分类标识符)。   [![godoc][GoDoc]](https://godoc.org/github.com/oklog/ulid)
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  没有麻烦，安全，快速的唯一标识符与命令。
* [uuid](https://github.com/agext/uuid)  使用快速或加密质量的随机节点标识符生成、编码和解码UUIDs v1。
* [uuid](https://github.com/gofrs/uuid) **star:654** 通用唯一标识符(UUID)的实现。支持uuid的创建和解析。积极维护satori uuid的fork。   [![godoc][GoDoc]](https://godoc.org/github.com/gofrs/uuid)
* [wuid](https://github.com/edwingeng/wuid) **star:328** 一个非常快的唯一数字生成器，比UUID快10-135倍。   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/wuid)
* [goid](https://github.com/jakehl/goid) **star:27** 生成和解析RFC4122兼容的V4 uuid。   [![最近一年没有更新][Yellow]](https://github.com/jakehl/goid)   [![godoc][GoDoc]](https://godoc.org/github.com/jakehl/goid)
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:16** 一个小而有效的Go唯一字符串ID生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/aidarkhanov/nanoid)
* [sno](https://github.com/muyo/sno)  使用嵌入元数据的紧凑、可排序和快速的惟一id。

## 验证

*库进行验证。 (翻译出错了? 试试 [英文版](README_EN.md#validation) 吧~)*

* [validator](https://github.com/go-playground/validator) **star:4772**  Go 结构体及域验证，包括：跨域、跨结构体, Map, 切片和数组。   [![star > 2000][Awesome]](https://github.com/go-playground/validator)   [![最近一周有更新][Green]](https://github.com/go-playground/validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/validator)
* [govalidator](https://github.com/asaskevich/govalidator) **star:4037** 用于字符串，数字，切片和结构的验证器和sanitizers。   [![star > 2000][Awesome]](https://github.com/asaskevich/govalidator)   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/govalidator)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1283** 支持各种数据类型(结构、字符串、映射、片等)的验证，使用可配置和可扩展的验证规则，这些规则在通常的代码构造中指定，而不是在结构标签中指定。   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
* [terraform-validator](https://github.com/thazelart/terraform-validator)  一种规范和约定验证器。
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:834** 用简单的规则验证Golang请求数据。深受Laravel请求验证的启发。   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/govalidator)
* [checkdigit](https://github.com/osamingo/checkdigit)  提供校验数字算法(Luhn, Verhoeff, Damm)和计算器(ISBN, EAN, JAN, UPC等)。
* [validate](https://github.com/gookit/validate) **star:202**  Go 封装数据验证和过滤。支持验证映射、结构、请求(表单、JSON、url)。值，上载文件)数据和更多特性。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/validate)   [![包含中文文档][CN]](https://github.com/gookit/validate)
* [jio](https://github.com/faceair/jio) **star:30** jio是一个json模式验证器，类似于[joi](https://github.com/hapijs/joi)。   [![godoc][GoDoc]](https://godoc.org/github.com/faceair/jio)   [![包含中文文档][CN]](https://github.com/faceair/jio)
* [validate](https://github.com/gobuffalo/validate) **star:29** 这个包提供了一个框架，用于为Go应用程序编写验证。   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/validate)
* [gody](https://github.com/guiferpa/gody) **star:17** :一个轻量级的结构验证器。   [![godoc][GoDoc]](https://godoc.org/github.com/guiferpa/gody)
* [govalid](https://github.com/twharmon/govalid) **star:13** 快速、基于标签的结构验证。   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/govalid)

## 版本控制

*用于版本控制的库。 (翻译出错了? 试试 [英文版](README_EN.md#version-control) 吧~)*

* [git2go](https://github.com/libgit2/git2go) **star:1449**  libgit2 的 Go 语言接口。   [![最近一周有更新][Green]](https://github.com/libgit2/git2go)   [![godoc][GoDoc]](https://godoc.org/github.com/libgit2/git2go)
* [go-git](https://github.com/src-d/go-git)  纯Go中高度可扩展的Git实现。
* [hercules](https://github.com/src-d/hercules) **star:876** 从Git存储库历史中获得高级见解。   [![最近一周有更新][Green]](https://github.com/src-d/hercules)   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/hercules)
* [gh](https://github.com/rjeczalik/gh) **star:72** 用于GitHub webhook的可编写脚本的服务器和net/http中间件。   [![最近一年没有更新][Yellow]](https://github.com/rjeczalik/gh)   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/gh)
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:72** 在Go中操作和检查VCS存储库。   [![godoc][GoDoc]](https://godoc.org/github.com/sourcegraph/go-vcs)
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo是一组Go包的集合，提供对本地Mercurial存储库的读取访问。   [![最近一年没有更新][Yellow]](https://github.com/beyang/hgo)   [![godoc][GoDoc]](https://godoc.org/github.com/beyang/hgo)

## 视频

*用于操作视频的库。 (翻译出错了? 试试 [英文版](README_EN.md#video) 吧~)*

* [gmf](https://github.com/3d0c/gmf) **star:584**  FFmpeg av\* 库的 Go 语言接口。   [![godoc][GoDoc]](https://godoc.org/github.com/3d0c/gmf)
* [go-astits](https://github.com/asticode/go-astits) **star:294** 在GO中解析和演示MPEG传输流(.ts)。   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astits)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8)  苹果m3u8播放列表的解析器和生成器库。
* [goav](https://github.com/giorgisio/goav)  FFmpeg的Comphrensive。
* [go-astisub](https://github.com/asticode/go-astisub) **star:226** 使用 Go 语言操作字幕(.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.)。   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astisub)
* [gst](https://github.com/ziutek/gst) **star:155**  GStreamer的Go工具。   [![最近一年没有更新][Yellow]](https://github.com/ziutek/gst)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/gst)
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:91** Go绑定libvlc 2.X/3.X/4。X(由VLC媒体播放器使用)。   [![最近一周有更新][Green]](https://github.com/adrg/libvlc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/libvlc-go)
* [m3u8](https://github.com/grafov/m3u8)  苹果HLS的M3U8播放列表解析器和生成器库。
* [v4l](https://github.com/korandiz/v4l) **star:37** 用于Linux的视频捕捉库，用Go编写。   [![最近一年没有更新][Yellow]](https://github.com/korandiz/v4l)   [![godoc][GoDoc]](https://godoc.org/github.com/korandiz/v4l)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:13** 字幕格式支持 .srt、.ttml和.ass。   [![最近一年没有更新][Yellow]](https://github.com/wargarblgarbl/libgosubs)   [![godoc][GoDoc]](https://godoc.org/github.com/wargarblgarbl/libgosubs)

## Web框架

*全栈 web 框架。 (翻译出错了? 试试 [英文版](README_EN.md#web-frameworks) 吧~)*

* [Gin](https://github.com/gin-gonic/gin) **star:35644** Gin是一个用Go编写的web框架!它具有一个类似于martini的API，性能更好，速度快40倍。   [![star > 2000][Awesome]](https://github.com/gin-gonic/gin)   [![最近一周有更新][Green]](https://github.com/gin-gonic/gin)   [![godoc][GoDoc]](https://godoc.org/github.com/gin-gonic/gin)
* [Beego](https://github.com/astaxie/beego) **star:23289** beego是一种用于 Go 编程语言的开源高性能web框架。   [![star > 2000][Awesome]](https://github.com/astaxie/beego)   [![最近一周有更新][Green]](https://github.com/astaxie/beego)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/beego)   [![包含中文文档][CN]](https://github.com/astaxie/beego)
* [Buffalo](http://gobuffalo.io)  为 Go 语言带来堪比 Rails 的高生产效率!
* [Echo](https://github.com/labstack/echo)  高性能、极简的Go web框架。
* [Revel](https://github.com/revel/revel) **star:11580** 用于Go语言的高效web框架。   [![star > 2000][Awesome]](https://github.com/revel/revel)   [![godoc][GoDoc]](https://godoc.org/github.com/revel/revel)
* [goa](https://github.com/goa-go/goa)  goa就像golang中的koajs一样，是一个灵活、轻量级、高性能和可扩展的基于中间件的web框架。
* [Goa](https://github.com/goadesign/goa) **star:3737** Goa为在Go中开发远程api和微服务提供了一种全面的方法。   [![star > 2000][Awesome]](https://github.com/goadesign/goa)   [![最近一周有更新][Green]](https://github.com/goadesign/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goadesign/goa)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3392** 设置RESTful JSON API的快速简便方法。   [![star > 2000][Awesome]](https://github.com/ant0ine/go-json-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ant0ine/go-json-rest)
* [Fiber](https://github.com/gofiber/fiber) **star:2899** 一个灵感来自Express.js的web框架构建在Fasthttp上。   [![star > 2000][Awesome]](https://github.com/gofiber/fiber)   [![最近一周有更新][Green]](https://github.com/gofiber/fiber)   [![godoc][GoDoc]](https://godoc.org/github.com/gofiber/fiber)
* [utron](https://github.com/gernest/utron) **star:2159** Go(Golang)的轻量级MVC框架。   [![star > 2000][Awesome]](https://github.com/gernest/utron)   [![最近一年没有更新][Yellow]](https://github.com/gernest/utron)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/utron)
* [gongular](https://github.com/mustafaakin/gongular) **star:425**  快速 Go web 框架，支持输入映射／验证以及依赖注入。   [![最近一年没有更新][Yellow]](https://github.com/mustafaakin/gongular)   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaakin/gongular)
* [neo](https://github.com/ivpusic/neo) **star:405** Neo是一个非常简单且快速的Web框架API。   [![最近一年没有更新][Yellow]](https://github.com/ivpusic/neo)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/neo)
* [Air](https://github.com/aofei/air) **star:370** 一个理想的精细化的Go web框架。   [![最近一周有更新][Green]](https://github.com/aofei/air)   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/air)
* [Banjo](https://github.com/nsheremet/banjo)  非常简单和快速的网络框架 Go 。
* [mango](https://github.com/paulbellamy/mango) **star:344** ManGo 是一个模块化 web 应用框架，受到 Rack 和 PEP333 的启发。   [![最近一年没有更新][Yellow]](https://github.com/paulbellamy/mango)   [![godoc][GoDoc]](https://godoc.org/github.com/paulbellamy/mango)
* [Aero](https://github.com/aerogo/aero) **star:246** 高性能的Go web框架，在Lighthouse中达到最高分。   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/aero)
* [Golf](https://github.com/dinever/golf) **star:242** Golf 是一个快速、简单、轻量级的 Go 语言微型 web 框架。具有强大的功能且没有标准库以外的依赖。   [![最近一年没有更新][Yellow]](https://github.com/dinever/golf)   [![godoc][GoDoc]](https://godoc.org/github.com/dinever/golf)
* [Gondola](https://github.com/rainycape/gondola)  web框架写的网站越快越好。
* [go-rest](https://github.com/ungerik/go-rest) **star:118** 小型的 REST 框架。   [![最近一年没有更新][Yellow]](https://github.com/ungerik/go-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-rest)
* [hiboot](https://github.com/hidevopsio/hiboot) **star:115** hiboot是一个高性能的web应用程序框架，支持自动配置和依赖注入。   [![最近一周有更新][Green]](https://github.com/hidevopsio/hiboot)   [![godoc][GoDoc]](https://godoc.org/github.com/hidevopsio/hiboot)   [![包含中文文档][CN]](https://github.com/hidevopsio/hiboot)
* [Macaron](https://github.com/go-macaron/macaron)  Macaron 是一个高效的模块化设计的web框架
* [WebGo](https://github.com/bnkamalesh/webgo) **star:89** 构建web应用程序的微框架;处理程序链接、中间件和上下文注入。与标准库兼容的HTTP处理程序(即http.HandlerFunc)。   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/webgo)
* [uAdmin](https://github.com/uadmin/uadmin) **star:79** 受到 Sinatra 启发的 Go 语言 web 框架。   [![最近一周有更新][Green]](https://github.com/uadmin/uadmin)   [![godoc][GoDoc]](https://godoc.org/github.com/uadmin/uadmin)
* [Microservice](https://github.com/claygod/microservice) **star:71** 创建微服务的框架，用Golang编写。   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/microservice)
* [Golax](https://github.com/fulldump/golax) **star:71** 一个非Sinatra快速HTTP框架，支持谷歌自定义方法、深度拦截器、递归等。   [![最近一年没有更新][Yellow]](https://github.com/fulldump/golax)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/golax)
* [Ginrpc](https://github.com/xxjwxc/ginrpc) **star:58** Gin参数自动绑定工具，Gin rpc工具。   [![最近一周有更新][Green]](https://github.com/xxjwxc/ginrpc)   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/ginrpc)   [![包含中文文档][CN]](https://github.com/xxjwxc/ginrpc)
* [Gizmo](https://github.com/NYTimes/gizmo)  《纽约时报》使用的微服务工具包。
* [YARF](https://github.com/yarf-framework/yarf) **star:55** 快速微框架，旨在以快速和简单的方式构建REST api和web服务。   [![godoc][GoDoc]](https://godoc.org/github.com/yarf-framework/yarf)
* [Flamingo](https://github.com/i-love-flamingo/flamingo)  可插拔web项目的框架。包括模块的概念和提供DI、Configareas、i18n、模板引擎、graphql、可观察性、安全性、事件、路由和反向路由等功能。
* [patron](https://github.com/beatlabs/patron) **star:50** Patron是一个遵循最佳云实践的微服务框架，专注于提升开发效率。   [![最近一周有更新][Green]](https://github.com/beatlabs/patron)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/patron)
* [Fireball](https://github.com/zpatrick/fireball) **star:50** 感觉更加自然的 web 框架。   [![最近一年没有更新][Yellow]](https://github.com/zpatrick/fireball)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/fireball)
* [vox](https://github.com/aisk/vox) **star:48** 一个面向人类的golang web框架，深受Koa的启发。   [![godoc][GoDoc]](https://godoc.org/github.com/aisk/vox)
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) **star:43** 提供使用干净的体系结构(如DDD和端口和适配器)的电子商务功能，您可以使用这些功能来构建灵活的电子商务应用程序。   [![最近一周有更新][Green]](https://github.com/i-love-flamingo/flamingo-commerce)   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo-commerce)
* [aah](https://aahframework.org)  可伸缩、高性能、快速开发的Go Web框架。
* [Resoursea](https://github.com/resoursea/api) **star:31** 用于快速编写基于资源的服务的REST框架。   [![最近一年没有更新][Yellow]](https://github.com/resoursea/api)   [![godoc][GoDoc]](https://godoc.org/github.com/resoursea/api)
* [REST Layer](http://rest-layer.io)  框架，用于在数据库之上构建REST/GraphQL API，主要是通过代码进行配置。
* [rex](https://github.com/goanywhere/rex) **star:29**  Rex 是一个用于进行模块化开发的库，基于Gorilla/mux 完全兼容大多数的 net/HTTP.   [![最近一年没有更新][Yellow]](https://github.com/goanywhere/rex)   [![godoc][GoDoc]](https://godoc.org/github.com/goanywhere/rex)
* [rux](https://github.com/gookit/rux) **star:23** 简单而快速的web框架，可用于构建golang HTTP应用程序   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/rux)   [![包含中文文档][CN]](https://github.com/gookit/rux)
* [tango](https://github.com/lunny/tango)  微型的、支持插件的 web 框架。
* [tigertonic](https://github.com/rcrowley/go-tigertonic)  用于构建 JSON web 服务的 Go 语言框架，受到 Dropwizard 的启发。
* [Goyave](https://github.com/System-Glitch/goyave)  功能齐全的web框架旨在干净的代码和快速的开发，具有强大的内置功能。
* [goweb](https://github.com/twharmon/goweb) **star:2** 具有路由、websockets、日志、中间件、静态文件服务器(可选gzip)和自动TLS的Web框架。   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/goweb)

### 中间件

#### 仿真中间件

* [Tollbooth](https://github.com/didip/tollbooth) **star:1488** 限制速率的 HTTP 请求处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/didip/tollbooth)
* [CORS](https://github.com/rs/cors) **star:1387** 轻松地向API添加CORS功能。   [![godoc][GoDoc]](https://godoc.org/github.com/rs/cors)
* [Limiter](https://github.com/ulule/limiter) **star:871** 简单的速度限制中间件。   [![最近一周有更新][Green]](https://github.com/ulule/limiter)   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/limiter)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:772** 添加/解析Server-Timing头。   [![最近一年没有更新][Yellow]](https://github.com/mitchellh/go-server-timing)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/go-server-timing)
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:100** 使用Lightning Network(比特币)实现基于每个请求的api货币化中间件。   [![最近一年没有更新][Yellow]](https://github.com/philippgille/ln-paywall)   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/ln-paywall)
* [XFF](https://github.com/sebest/xff) **star:73** 处理 X-Forwarded-For 头的中间件。   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/xff)
* [client-timing](https://github.com/posener/client-timing)  用于服务器定时报头的HTTP客户机。
* [formjson](https://github.com/rs/formjson) **star:34** 透明地将JSON输入作为标准表单POST处理。   [![最近一年没有更新][Yellow]](https://github.com/rs/formjson)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/formjson)

#### 用于创建HTTP中间件的库

* [render](https://github.com/unrolled/render) **star:1336** Go package用于方便地呈现JSON、XML和HTML模板响应。   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/render)
* [interpose](https://github.com/carbocation/interpose) **star:288** golang的极简网络/http中间件。   [![最近一年没有更新][Yellow]](https://github.com/carbocation/interpose)   [![godoc][GoDoc]](https://godoc.org/github.com/carbocation/interpose)
* [muxchain](https://github.com/stephens2424/muxchain) **star:209** 用于net/http的轻量级中间件。   [![godoc][GoDoc]](https://godoc.org/github.com/stephens2424/muxchain)
* [negroni](https://github.com/urfave/negroni)  符合语言习惯的 HTTP 中间件库。
* [renderer](https://github.com/thedevsaddam/renderer) **star:183** 简单、轻量级和更快的响应(JSON、JSONP、XML、YAML、HTML、文件)。   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/renderer)
* [rye](https://github.com/InVisionApp/rye) **star:93** 支持JWT、CORS、Statsd和Go 1.7上下文的小型Go中间件库(带有罐装中间件)。   [![最近一年没有更新][Yellow]](https://github.com/InVisionApp/rye)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/rye)
* [stats](https://github.com/thoas/stats)  使用中间件来存储关于web应用程序的各种信息。
* [gores](https://github.com/alioygur/gores) **star:88** 处理HTML、JSON、XML等响应的Go包。对于RESTful api非常有用。   [![最近一年没有更新][Yellow]](https://github.com/alioygur/gores)   [![godoc][GoDoc]](https://godoc.org/github.com/alioygur/gores)
* [alice](https://github.com/justinas/alice)  用于连接中间件的库，简单无痛苦。
* [go-wrap](https://github.com/go-on/wrap) **star:59** net/http的小型中间件包。   [![最近一年没有更新][Yellow]](https://github.com/go-on/wrap)   [![godoc][GoDoc]](https://godoc.org/github.com/go-on/wrap)
* [catena](https://github.com/codemodus/catena) **star:7** HTTP.Handler wrapper catenation (和chain具有相同的 API ).。   [![最近一年没有更新][Yellow]](https://github.com/codemodus/catena)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/catena)
* [chain](https://github.com/codemodus/chain)  带有范围数据的处理程序包装器链接(基于网络/上下文的“中间件”)。

### 路由器

* [httprouter](https://github.com/julienschmidt/httprouter) **star:10815** 高性能路由。使用这个库和标准http处理工具可以构建一个非常高性能大web框架。   [![star > 2000][Awesome]](https://github.com/julienschmidt/httprouter)   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/httprouter)
* [chi](https://github.com/go-chi/chi) **star:7047** 小巧、快速、具有丰富表达力的 HTTP 路由，基于net/context.。   [![star > 2000][Awesome]](https://github.com/go-chi/chi)   [![最近一周有更新][Green]](https://github.com/go-chi/chi)   [![godoc][GoDoc]](https://godoc.org/github.com/go-chi/chi)
* [gocraft/web](https://github.com/gocraft/web) **star:1407** Mux和中间件包在Go中。   [![godoc][GoDoc]](https://godoc.org/github.com/gocraft/web)
* [Bone](https://github.com/go-zoo/bone) **star:1248** 闪电快速HTTP多路复用器。   [![godoc][GoDoc]](https://godoc.org/github.com/go-zoo/bone)
* [Bxog](https://github.com/claygod/Bxog)  简单和快速的HTTP路由器 Go 。它可以处理不同难度、长度和嵌套的路径。他还知道如何根据接收到的参数创建URL。
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:823** 高性能路由器分叉从`httprouter`。第一个路由器适合`fasthttp`。   [![godoc][GoDoc]](https://godoc.org/github.com/buaazp/fasthttprouter)
* [Goji](https://github.com/goji/goji) **star:799** 枸杞是一种简约的和灵活的与支持'net/context` HTTP请求多路复用器。   [![godoc][GoDoc]](https://godoc.org/github.com/goji/goji)
* [goroute](https://github.com/goroute/route)  简单但功能强大的HTTP请求多路复用器。
* [GoRouter](https://github.com/vardius/gorouter)  GoRouter 是一个服务器/API 微型框架、HTTP 请求路由 router, 数据分选器，提供了支持net/context的中间件。
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:474** 一个简单和快速的HTTP路由器 Go 。   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gorouter)
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:409** 高速，灵活，基于树的 HTTP 路由。受到了 httprouter 的启发。   [![godoc][GoDoc]](https://godoc.org/github.com/dimfeld/httptreemux)
* [lars](https://github.com/go-playground/lars) **star:380** 是一个轻量级、快速、可扩展、零分配的HTTP路由，用于创建定制化的框架。   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/lars)
* [mux](https://github.com/gorilla/mux)  强大的URL路由器和调度器为golang。
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:373** 一个非常快的Go (golang) HTTP路由器，支持正则表达式路由匹配。完全支持构建RESTful api。   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
* [Siesta](https://github.com/VividCortex/siesta) **star:353** 编写中间件和处理程序的可组合框架。   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/siesta)
* [vestigo](https://github.com/husobee/vestigo) **star:261** 高性能，独立，HTTP兼容的URL路由器的go web应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/husobee/vestigo)
* [gowww/router](https://github.com/gowww/router) **star:159** 超快的HTTP 路由，完全兼容 net/HTTP.Handler 接口.。   [![最近一年没有更新][Yellow]](https://github.com/gowww/router)   [![godoc][GoDoc]](https://godoc.org/github.com/gowww/router)
* [violetear](https://github.com/nbari/violetear) **star:97** HTTP路由器。   [![godoc][GoDoc]](https://godoc.org/github.com/nbari/violetear)
* [bellt](https://github.com/GuilhermeCaruso/bellt)  一个简单的Go HTTP路由器。
* [alien](https://github.com/gernest/alien)  轻量级和快速http路由器从外层空间。
* [pure](https://github.com/go-playground/pure) **star:91** 是一个轻量级HTTP路由器，它坚持net/ HTTP“实现”的std。   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pure)
* [xmux](https://github.com/rs/xmux) **star:88** 高性能mux基于httprouter 'net/context`支持。   [![最近一年没有更新][Yellow]](https://github.com/rs/xmux)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xmux)
* [FastRouter](https://github.com/razonyang/fastrouter) **star:19** 一个快速，灵活的HTTP路由器写在Go。   [![最近一年没有更新][Yellow]](https://github.com/razonyang/fastrouter)   [![godoc][GoDoc]](https://godoc.org/github.com/razonyang/fastrouter)

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:603** 为 Go 语言实现的 Win32 OLE。   [![godoc][GoDoc]](https://godoc.org/github.com/go-ole/go-ole)
* [d3d9](https://github.com/gonutz/d3d9) **star:95**  Direct3D9 的 Go 语言接口。   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/d3d9)
* [gosddl](https://github.com/MonaxGT/gosddl) **star:2** 从SDDL-string到用户友好的JSON的转换器。SDDL由四个部分组成:所有者、主群、DACL、SACL。   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gosddl)

## XML

*用于操作XML的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#xml) 吧~)*

* [zek](https://github.com/miku/zek) **star:325** 从XML生成Go结构。   [![godoc][GoDoc]](https://godoc.org/github.com/miku/zek)
* [xpath](https://github.com/antchfx/xpath) **star:281** Go的XPath包。   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xpath)
* [xquery](https://github.com/antchfx/xquery) **star:154** XQuery允许您使用XPath表达式从HTML/XML文档中提取数据。   [![最近一年没有更新][Yellow]](https://github.com/antchfx/xquery)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xquery)   [![归档项目][Archived]](https://github.com/antchfx/xquery)
* [xml2map](https://github.com/sbabiv/xml2map) **star:19** XML来映射转换器编写的Golang。   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/xml2map)
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** 简单的命令行XML比较器，生成文件夹、文件和标记的差异。   [![最近一年没有更新][Yellow]](https://github.com/xml-comp/xml-comp)   [![godoc][GoDoc]](https://godoc.org/github.com/xml-comp/xml-comp)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:9** 基于libxml2的xmlwriter模块的过程性XML生成API。   [![godoc][GoDoc]](https://godoc.org/github.com/shabbyrobe/xmlwriter)

# 工具

* Go 软件和插件。 (翻译出错了? 试试 [英文版](README_EN.md#tools) 吧~)*

## 代码分析

* [GoLint](https://github.com/golang/lint) **star:3391** Go 源码的 linter。   [![star > 2000][Awesome]](https://github.com/golang/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/lint)
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [errcheck](https://github.com/kisielk/errcheck) **star:1404** Errcheck是一个用于检查Go程序中未检查错误的程序。   [![godoc][GoDoc]](https://godoc.org/github.com/kisielk/errcheck)
* [gcvis](https://github.com/davecheney/gcvis) **star:959** 实时可视化跟踪 GC 数据。   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/gcvis)
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  在源码中寻找没有直接单元测试的函数和方法。
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  用于大量静态分析检查，您可能已经从 c# 的 ReSharper 等工具中习惯了这些检查。
* [php-parser](https://github.com/z7zmey/php-parser) **star:696** 用 Go 编写的 PHP 解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/z7zmey/php-parser)
* [go-critic](https://github.com/go-critic/go-critic) **star:682** 源代码检查工具。   [![最近一周有更新][Green]](https://github.com/go-critic/go-critic)   [![godoc][GoDoc]](https://godoc.org/github.com/go-critic/go-critic)
* [GoCover.io](http://gocover.io/)  GoCover.io 提供了任意 golang 包的代码覆盖率服务。
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  来修复(添加，删除) Go 中自动导入的工具。
* [GolangCI](https://golangci.com/)  GolangCI 是一个针对 GitHub pull 请求的自动代码审查服务。服务是开源的，对开源项目是免费的。
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:427** 基于 Web 的 Golang AST 可视化工具。
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:330** 找出项目中过期的依赖项。   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/go-mod-outdated)
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:314** go-cleanarch 的创建是为了验证 Clean 体系结构规则，比如 Go 项目中的依赖关系。   [![godoc][GoDoc]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
* [unconvert](https://github.com/mdempsky/unconvert) **star:280** 在源码中删除不必要的类型转换。   [![godoc][GoDoc]](https://godoc.org/github.com/mdempsky/unconvert)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  对未使用的常量、变量、函数和类型的代码进行检查。
* [gostatus](https://github.com/shurcooL/gostatus) **star:241** 用于显示包含 Go 包的存储库的状态的命令行工具，。   [![最近一年没有更新][Yellow]](https://github.com/shurcooL/gostatus)   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/gostatus)
* [dupl](https://github.com/mibk/dupl) **star:190** 用于代码克隆检测的工具。   [![最近一年没有更新][Yellow]](https://github.com/mibk/dupl)   [![godoc][GoDoc]](https://godoc.org/github.com/mibk/dupl)
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:172** 检查 Go 项目最近的向下不兼容修改。   [![最近一年没有更新][Yellow]](https://github.com/bradleyfalzon/apicompat)   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyfalzon/apicompat)
* [tickgit](https://github.com/augmentable-dev/tickgit) **star:130** CLI和go包用于显示代码注释TODOs(以任何语言)并应用“git blame”来识别作者。   [![godoc][GoDoc]](https://godoc.org/github.com/augmentable-dev/tickgit)
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:103** checkstyle是一个类似于java checkstyle的检查工具。   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/checkstyle)
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple 是 Go 源代码的linter，专门用于简化代码。
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  添加 zero 返回声明，以匹配 func 返回类型。
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) **star:102** 生成文本plantump类图的库和CLI，其中包含关于结构和接口的信息，以及它们之间的关系。   [![godoc][GoDoc]](https://godoc.org/github.com/jfeliu007/goplantuml)
* [lint](https://github.com/surullabs/lint) **star:64** 将 linters 作为测试的一部分。   [![最近一年没有更新][Yellow]](https://github.com/surullabs/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/surullabs/lint)
* [validate](https://github.com/mccoyst/validate) **star:62** 使用 tags 自动验证结构体字段。   [![最近一年没有更新][Yellow]](https://github.com/mccoyst/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/validate)
* [go-outdated](https://github.com/firstrow/go-outdated) **star:45** 显示过期包的终端应用。   [![最近一年没有更新][Yellow]](https://github.com/firstrow/go-outdated)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/go-outdated)   [![归档项目][Archived]](https://github.com/firstrow/go-outdated)
* [golines](https://github.com/segmentio/golines) **star:35** 格式化程序，它自动缩短Go代码中的长行。   [![godoc][GoDoc]](https://godoc.org/github.com/segmentio/golines)

## 编辑器插件

* [vim-go](https://github.com/fatih/vim-go) **star:11702** Go 开发会用到的 Vim 插件。   [![star > 2000][Awesome]](https://github.com/fatih/vim-go)   [![最近一周有更新][Green]](https://github.com/fatih/vim-go)
* [vscode-go](https://github.com/Microsoft/vscode-go)  Visual Studio代码的扩展(VS代码)，它提供了对Go语言的支持。
* [Watch](https://github.com/eaburns/Watch)  Runs a command in an acme win on file changes.
* [go-mode](https://github.com/dominikh/go-mode.el) **star:1031** 在 GNU/Emacs 支持 GO。
* [go-plus](https://github.com/joefitzgerald/go-plus)  在 Atom 中添加自动完成，格式化，语法检查，高亮和审查。
* [gocode](https://github.com/nsf/gocode)  Autocompletion daemon for the Go programming language.
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  在 VS Code 中支持 Go 的基准分析。
* [GoSublime](https://github.com/DisposaBoy/GoSublime)  包含了可为文本编辑器 SublimeText 3 提供代码自动填充和其他类似IDE的功能的 Golang IDE 插件集合。
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  JetBrains IDEs 的 Go 插件。
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:32** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:19** 基于函数或方法的签名生成Go测试的Vim插件。   [![最近一年没有更新][Yellow]](https://github.com/hexdigest/gounit-vim)
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:13** 在 Theia IDE中支持 Go。
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go)  在保存时突出显示语法错误的 Vim 插件。

## Go 生成工具

* [generic](https://github.com/usk81/generic)  灵活的 Go 数据类型。
* [genny](https://github.com/cheekybits/genny) **star:1133** 优雅的 Go 泛型。   [![godoc][GoDoc]](https://godoc.org/github.com/cheekybits/genny)
* [gocontracts](https://github.com/Parquery/gocontracts) **star:58** 通过同步代码和文档来实现 design-by-contract 设计。   [![最近一年没有更新][Yellow]](https://github.com/Parquery/gocontracts)   [![godoc][GoDoc]](https://godoc.org/github.com/Parquery/gocontracts)
* [gonerics](http://github.com/bouk/gonerics)  Go中的易用的泛型。
* [gotests](https://github.com/cweill/gotests)  从源代码生成测试用例。
* [gounit](https://github.com/hexdigest/gounit)  使用您自己的模板生成Go测试用例。
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:37** 可生成用于切片的 Haskell。   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/hasgo)
* [re2dfa](https://github.com/opennota/re2dfa)  将正则表达式转换为有限状态机，并输出 Go 源代码。
* [TOML-to-Go](https://xuri.me/toml-to-go)  在浏览器中将 TOML 转换为 Go 类型。

## Go 工具

* [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:4194** 借助的 OctoLinker 浏览器扩展，可以高效的地浏览  GitHub go文件。   [![star > 2000][Awesome]](https://github.com/OctoLinker/browser-extension)   [![最近一周有更新][Green]](https://github.com/OctoLinker/browser-extension)
* [go-james](https://github.com/pieterclaerhout/go-james)  去项目骨架创造者，建立和测试您的项目没有手动设置。
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:2384** 使用 dot format 可视化 Go 程序的调用图。   [![star > 2000][Awesome]](https://github.com/TrueFurby/go-callvis)   [![最近一周有更新][Green]](https://github.com/TrueFurby/go-callvis)   [![godoc][GoDoc]](https://godoc.org/github.com/TrueFurby/go-callvis)
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete)  Bash completion for go and wgo。
* [go-swagger](https://github.com/go-swagger/go-swagger)  基于 Go 的Swagger 2.0实现。
* [godbg](https://github.com/tylerwince/godbg)  实现了 Rusts 的 dbg! 宏，可以方便的在开发过程中快速、容易地调试。
* [gomodrun](https://github.com/dustinblackman/gomodrun/)  执行并缓存Go中包含的二进制文件的Go工具。国防部文件。
* [depth](https://github.com/KyleBanks/depth) **star:456** 通过分析导入，将包依赖关系树可视化输出。   [![godoc][GoDoc]](https://godoc.org/github.com/KyleBanks/depth)
* [gb](https://getgb.io/)  一个基于项目的易用的构建工具。
* [richgo](https://github.com/kyoh86/richgo) **star:434** 用文本装饰丰富 go test 的输出。   [![godoc][GoDoc]](https://godoc.org/github.com/kyoh86/richgo)
* [rts](https://github.com/galeone/rts) **star:195** 从服务器响应生成Go结构。   [![最近一年没有更新][Yellow]](https://github.com/galeone/rts)   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/rts)
* [colorgo](https://github.com/songgao/colorgo) **star:103** 将 go 命令包装成彩色的 go build 输出。   [![最近一年没有更新][Yellow]](https://github.com/songgao/colorgo)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/colorgo)
* [gothanks](https://github.com/psampaz/gothanks) **star:89** GoThanks自动帮你开始。mod github依赖，以这种方式发送一些爱给他们的维护者。   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/gothanks)
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:17** 一个[Yeoman](http://yeoman.io)生成器，用于启动新的 Go 项目。
* [gilbert](https://go-gilbert.github.io)  一个为 Go 项目而生的构建系统和任务运行器。

## 软件包

*用Go编写的软件。 (翻译出错了? 试试 [英文版](README_EN.md#software-packages) 吧~)*

### DevOps 工具

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:63482** 来自谷歌的容器集群管理器。   [![star > 2000][Awesome]](https://github.com/kubernetes/kubernetes)   [![最近一周有更新][Green]](https://github.com/kubernetes/kubernetes)   [![godoc][GoDoc]](https://godoc.org/github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby) **star:56432** Collaborative project for the container ecosystem to assemble container-based systems.   [![star > 2000][Awesome]](https://github.com/moby/moby)   [![最近一周有更新][Green]](https://github.com/moby/moby)   [![godoc][GoDoc]](https://godoc.org/github.com/moby/moby)
* [Mora](https://github.com/emicklei/mora)  用于访问 MongoDB 文档和元数据的 REST 服务器。
* [uTask](https://github.com/ovh/utask)  对yaml中声明的业务流程进行建模和执行的自动化引擎。
* [traefik](https://github.com/containous/traefik) **star:27577** 反向代理和负载均衡器，支持多个后端。   [![star > 2000][Awesome]](https://github.com/containous/traefik)   [![最近一周有更新][Green]](https://github.com/containous/traefik)   [![godoc][GoDoc]](https://godoc.org/github.com/containous/traefik)
* [Vegeta](https://github.com/tsenart/vegeta) **star:13926** HTTP负载测试工具和库。超过9000 !   [![star > 2000][Awesome]](https://github.com/tsenart/vegeta)   [![最近一周有更新][Green]](https://github.com/tsenart/vegeta)   [![godoc][GoDoc]](https://godoc.org/github.com/tsenart/vegeta)
* [Packer](https://github.com/mitchellh/packer) **star:9825** 用于从一个源配置为多个平台创建相同的机器图像。   [![star > 2000][Awesome]](https://github.com/mitchellh/packer)   [![最近一周有更新][Green]](https://github.com/mitchellh/packer)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/packer)
* [Hey](https://github.com/rakyll/hey)  压力测试工具，可用来代替 ApacheBench (ab)。
* [GVM](https://github.com/moovweb/gvm) **star:5019** GVM 提供了一个接口来管理 Go 版本。   [![star > 2000][Awesome]](https://github.com/moovweb/gvm)
* [Wide](https://wide.b3log.org/login)  为使用 Golang 的团队提供基于 web 的 IDE。
* [webhook](https://github.com/adnanh/webhook) **star:4923** 允许用户创建在服务器上执行命令的 HTTP hooks。   [![star > 2000][Awesome]](https://github.com/adnanh/webhook)   [![godoc][GoDoc]](https://godoc.org/github.com/adnanh/webhook)
* [gox](https://github.com/mitchellh/gox) **star:3628** 非常简单，没有多余的跨平台编译工具。   [![star > 2000][Awesome]](https://github.com/mitchellh/gox)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/gox)
* [goxc](https://github.com/laher/goxc) **star:1643** 专注于跨平台编译和打包的 Go 构建工具。   [![godoc][GoDoc]](https://godoc.org/github.com/laher/goxc)
* [kala](https://github.com/ajvb/kala) **star:1408** 简单、现代和高性能的作业调度程序。   [![godoc][GoDoc]](https://godoc.org/github.com/ajvb/kala)
* [sg](https://github.com/ChristopherRabotin/sg)  可测试一组HTTP极限(如ab)，可以在每个调用之间使用响应代码和数据，根据之前的响应来确定特定的服务器压力。
* [script](https://github.com/bitfield/script) **star:1271** 让DevOps编写类shell和系统管理任务变得更加容易。   [![最近一周有更新][Green]](https://github.com/bitfield/script)   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/script)
* [Pomerium](https://github.com/pomerium/pomerium) **star:1027** Pomerium是一个可识别身份的访问代理。   [![最近一周有更新][Green]](https://github.com/pomerium/pomerium)   [![godoc][GoDoc]](https://godoc.org/github.com/pomerium/pomerium)
* [Rodent](https://github.com/alouche/rodent)  管理Go版本、项目和跟踪依赖项。
* [Gogs](https://gogs.io/)  自托管的Git服务。
* [govvv](https://github.com/ahmetalpbalkan/govvv)  可轻松地添加版本信息到 Go 二进制文件。
* [gonative](https://github.com/inconshreveable/gonative)  用原生 Go 创建一个跨平台的 Go 工具链。
* [godbg](https://github.com/sirnewton01/godbg)  基于 web 的 gdb 前端应用程序。
* [gobrew](https://github.com/cryptojuice/gobrew)  gobrew 允许您轻松地在 go 的多个版本之间切换。
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:714** 允许你的 Go应用程序 进行自我更新。   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/go-selfupdate)
* [skm](https://github.com/TimothyYe/skm) **star:589** SKM是一个简单而强大的SSH密钥管理器，它可以帮助您轻松地管理多个SSH密钥!   [![godoc][GoDoc]](https://godoc.org/github.com/TimothyYe/skm)
* [StatusOK](https://github.com/sanathp/statusok)  监视您的网站和REST api。当服务器宕机或响应时间超过预期时，通过Slack、电子邮件获得通知。
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi)  Terraform provider plugin，它根据OpenAPI文档(以前称为swagger文件)在运行时动态配置自身，该文档包含了公开的api的定义。
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:569** 从命令行管理 BareMetal 服务器(与使用Docker一样容易)。   [![最近一周有更新][Green]](https://github.com/scaleway/scaleway-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/scaleway/scaleway-cli)
* [lstags](https://github.com/ivanilves/lstags) **star:248** 提供了工具和API，可用来同步不同注册中心的Docker图像。   [![godoc][GoDoc]](https://godoc.org/github.com/ivanilves/lstags)
* [Pewpew](https://github.com/bengadbois/pewpew) **star:218** 灵活的 HTTP 命令行压测工具。   [![godoc][GoDoc]](https://godoc.org/github.com/bengadbois/pewpew)
* [manssh](https://github.com/xwjdsh/manssh) **star:211** manssh是一个命令行工具，可以方便地管理ssh别名配置。   [![最近一年没有更新][Yellow]](https://github.com/xwjdsh/manssh)   [![godoc][GoDoc]](https://godoc.org/github.com/xwjdsh/manssh)
* [aptly](https://github.com/smira/aptly)  Debian存储库管理工具。
* [awsenv](https://github.com/soniah/awsenv)  可以为 profile 加载Amazon (AWS)环境变量的轻量二进制文件。
* [aurora](https://github.com/xuri/aurora)  基于web的跨平台 Beanstalkd 队列服务器控制台。
* [bosun](https://github.com/bosun-monitor/bosun)  按照时间轴发出告警的框架。
* [Blast](https://github.com/dave/blast) **star:180** 一个用于API负载测试和批处理作业的简单工具。   [![最近一年没有更新][Yellow]](https://github.com/dave/blast)   [![godoc][GoDoc]](https://godoc.org/github.com/dave/blast)
* [bombardier](https://github.com/codesenberg/bombardier)  快速跨平台 HTTP 基准测试工具。
* [ostent](https://github.com/ostrost/ostent) **star:165** 收集和显示系统指标，并可选 Graphite and/or fluxdb作为依赖。   [![最近一年没有更新][Yellow]](https://github.com/ostrost/ostent)   [![godoc][GoDoc]](https://godoc.org/github.com/ostrost/ostent)
* [grapes](https://github.com/yaronsumel/grapes) **star:145** 旨在轻松地通过ssh分发命令的轻量级工具。   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/grapes)
* [jcli](https://github.com/jenkins-zh/jenkins-cli) **star:129** 詹金斯CLI允许你管理你的詹金斯作为一个简单的方法。   [![最近一周有更新][Green]](https://github.com/jenkins-zh/jenkins-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/jenkins-zh/jenkins-cli)   [![包含中文文档][CN]](https://github.com/jenkins-zh/jenkins-cli)
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  将所有GitHub repositories、issues、milestones 和 labels 都迁移到 Gitea。
* [Gitea](https://github.com/go-gitea/gitea)  从 Gogs fork，完全由社区驱动。
* [gaia](https://github.com/gaia-pipeline/gaia)  可用于任何编程语言来构建强大的管道。
* [fac](https://github.com/mkchoi212/fac)  修复 git 合并冲突。
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:127** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/easyssh-proxy)
* [kcli](https://github.com/cswank/kcli) **star:106** 用于检查kafka主题/分区/消息的命令行工具。   [![godoc][GoDoc]](https://godoc.org/github.com/cswank/kcli)
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:84** 在Windows机器上远程执行命令的Cli工具。   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm-cli)
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:75** 用Go编写的托管解决方案，可轻松地在AWS、GCP或DigitalOcean上部署应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/go-furnace/go-furnace)
* [dogo](https://github.com/liudng/dogo)  监视源文件中的更改并自动编译和运行(restart)。
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) **star:66** 一个go库和一个可执行文件，它使用各种输入通道生成有效的Dockerfiles。   [![godoc][GoDoc]](https://godoc.org/github.com/ozankasikci/dockerfile-generator)
* [drone-scp](https://github.com/appleboy/drone-scp)  通过 SSH 进行文件拷贝。其中 SSH 通过二进制文件、docker 或 Drone CI触发。
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:27** 使用二进制文件、docker或 Drone CI 来触发下游Jenkins作业。   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-jenkins)
* [Dropship](https://github.com/chrismckenzie/dropship)  通过 cdn 部署代码的工具。
* [DepCharge](https://github.com/centerorbit/depcharge) **star:11** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![godoc][GoDoc]](https://godoc.org/github.com/centerorbit/depcharge)
* [lwc](https://github.com/timdp/lwc) **star:11** UNIX wc命令的实时更新版本。   [![最近一年没有更新][Yellow]](https://github.com/timdp/lwc)   [![godoc][GoDoc]](https://godoc.org/github.com/timdp/lwc)
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) **star:2** 具有GET、PUT和DELETE方法和身份验证(OpenID连接和基本身份验证)的S3代理。   [![godoc][GoDoc]](https://godoc.org/github.com/oxyno-zeta/s3-proxy)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r)  小型实用程序/库，针对大型对象在Amazon S3中的高速传输进行了优化。

### 其他软件

* [restic](https://github.com/restic/restic) **star:9463** 消除重复项备份程序。   [![star > 2000][Awesome]](https://github.com/restic/restic)   [![最近一周有更新][Green]](https://github.com/restic/restic)   [![godoc][GoDoc]](https://godoc.org/github.com/restic/restic)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:9208** 快速、简单、可伸缩的分布式文件系统，采用了O(1)磁盘查找。   [![star > 2000][Awesome]](https://github.com/chrislusf/seaweedfs)   [![最近一周有更新][Green]](https://github.com/chrislusf/seaweedfs)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/seaweedfs)
* [confd](https://github.com/kelseyhightower/confd) **star:6828** 使用 etcd 或 consul 的模板和数据管理本地应用程序配置文件。   [![star > 2000][Awesome]](https://github.com/kelseyhightower/confd)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/confd)
* [Comcast](https://github.com/tylertreat/Comcast) **star:6459** 模拟坏的网络连接。   [![star > 2000][Awesome]](https://github.com/tylertreat/Comcast)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/Comcast)
* [mockingjay](https://github.com/quii/mockingjay-server)  一份配置文件中便可伪造HTTP服务器与用户之间的行为。您还可以使服务器随机宕机，以帮助进行更实际的性能测试。
* [LiteIDE](https://github.com/visualfc/liteide) **star:5870** 简单的、开源的、跨平台的Go IDE。   [![star > 2000][Awesome]](https://github.com/visualfc/liteide)   [![包含中文文档][CN]](https://github.com/visualfc/liteide)
* [nes](https://github.com/fogleman/nes) **star:4362** 任天堂娱乐系统(NES)模拟器。   [![star > 2000][Awesome]](https://github.com/fogleman/nes)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/nes)
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:4344** 为自动化测试模拟网络和系统条件的代理。   [![star > 2000][Awesome]](https://github.com/shopify/toxiproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/shopify/toxiproxy)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:3176** 跨平台网络和云备份工具。   [![star > 2000][Awesome]](https://github.com/gilbertchen/duplicacy)   [![最近一周有更新][Green]](https://github.com/gilbertchen/duplicacy)   [![godoc][GoDoc]](https://godoc.org/github.com/gilbertchen/duplicacy)
* [croc](https://github.com/schollz/croc) **star:2837** 轻松、安全地将文件或文件夹从一台计算机发送到另一台计算机。   [![star > 2000][Awesome]](https://github.com/schollz/croc)   [![最近一周有更新][Green]](https://github.com/schollz/croc)   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/croc)
* [Docker](http://www.docker.com/)  面向开发人员和系统管理员的分布式应用程序的开放平台。
* [Documize](https://github.com/documize/community)  集成了SaaS工具数据的现代wiki软件。
* [myLG](https://github.com/mehrdadrad/mylg) **star:2273** 命令行网络诊断工具。   [![star > 2000][Awesome]](https://github.com/mehrdadrad/mylg)   [![最近一周有更新][Green]](https://github.com/mehrdadrad/mylg)   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/mylg)
* [GoBoy](https://github.com/Humpheh/goboy) **star:2192** 用 Go 编写的任天堂Game Boy彩色模拟器。   [![star > 2000][Awesome]](https://github.com/Humpheh/goboy)   [![godoc][GoDoc]](https://godoc.org/github.com/Humpheh/goboy)
* [Stack Up](https://github.com/pressly/sup) **star:2107** Stack Up 是一个超级简单的部署工具 — 只面向Unix。   [![star > 2000][Awesome]](https://github.com/pressly/sup)   [![最近一周有更新][Green]](https://github.com/pressly/sup)   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/sup)
* [syncthing](https://syncthing.net/)  开放，分散的文件同步工具和协议。
* [lgo](https://github.com/yunabe/lgo) **star:1968** 与 Jupyter 可进行交互 Go 程序。它支持代码完成、代码检查以及与Go 100% 兼容性。   [![godoc][GoDoc]](https://godoc.org/github.com/yunabe/lgo)
* [limetext](http://limetext.org/)  一个强大而优雅的文本编辑器。
* [Circuit](https://github.com/gocircuit/circuit) **star:1817** Circuit 是一个可编程平台即服务(PaaS)和/或基础设施即服务(IaaS)，用于管理、发现、同步和编排包含云应用程序的服务和主机。   [![godoc][GoDoc]](https://godoc.org/github.com/gocircuit/circuit)
* [snap](https://github.com/intelsdi-x/snap) **star:1806** 强大的遥测框架。   [![最近一年没有更新][Yellow]](https://github.com/intelsdi-x/snap)   [![godoc][GoDoc]](https://godoc.org/github.com/intelsdi-x/snap)
* [scc](https://github.com/boyter/scc) **star:1337** 一个非常快速准确的代码计数器，采用了复杂的计算和 COCOMO 预估。   [![godoc][GoDoc]](https://godoc.org/github.com/boyter/scc)
* [peg](https://github.com/pointlander/peg) **star:676** 解析表达式语法，是Packrat解析器生成器的实现。   [![godoc][GoDoc]](https://godoc.org/github.com/pointlander/peg)
* [Leaps](https://github.com/jeffail/leaps) **star:665** 使用操作转换的成对编程服务。   [![godoc][GoDoc]](https://godoc.org/github.com/jeffail/leaps)
* [vFlow](https://github.com/VerizonDigital/vflow) **star:653** 高性能、可伸缩和可靠的 IPFIX、sFlow和 Netflow 收集器。   [![godoc][GoDoc]](https://godoc.org/github.com/VerizonDigital/vflow)
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store)  App that displays updates for the Go packages in your GOPATH.
* [gfile](https://github.com/Antonito/gfile) **star:533** 通过WebRTC在两台计算机之间安全地传输文件，不需要任何第三方依赖。   [![godoc][GoDoc]](https://godoc.org/github.com/Antonito/gfile)
* [shell2http](https://github.com/msoap/shell2http) **star:513** 通过http服务器执行shell命令(用于原型或远程控制)。   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/shell2http)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:395** 视频流 torrent 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
* [gocc](https://github.com/goccmack/gocc) **star:384** Gocc是一个用Go编写的编译器工具包。   [![godoc][GoDoc]](https://godoc.org/github.com/goccmack/gocc)
* [ipe](https://github.com/dimiro1/ipe) **star:298** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   [![最近一年没有更新][Yellow]](https://github.com/dimiro1/ipe)   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/ipe)
* [wellington](https://github.com/wellington/wellington) **star:288** Sass 项目管理工具，使用sprite函数(如Compass)扩展语言。   [![最近一年没有更新][Yellow]](https://github.com/wellington/wellington)   [![godoc][GoDoc]](https://godoc.org/github.com/wellington/wellington)
* [ide](https://github.com/thestrukture/ide) **star:254** 基于浏览器的IDE   [![godoc][GoDoc]](https://godoc.org/github.com/thestrukture/ide)
* [orange-cat](https://github.com/noraesae/orange-cat) **star:184** 用Go编写的Markdown预览器。   [![最近一年没有更新][Yellow]](https://github.com/noraesae/orange-cat)   [![godoc][GoDoc]](https://godoc.org/github.com/noraesae/orange-cat)
* [Orbit](https://github.com/gulien/orbit) **star:134** 一个根据模板来运行命令和生成文件的简单小工具。   [![godoc][GoDoc]](https://godoc.org/github.com/gulien/orbit)
* [Juju](https://jujucharms.com/)  Cloud-agnostic的服务部署和编制 —— 支持EC2、Azure、Openstack、MAAS等。
* [joincap](https://github.com/assafmo/joincap) **star:128** 用于合并多个pcap文件的命令行实用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/assafmo/joincap)
* [borg](https://github.com/crufter/borg)  基于终端的bash代码段搜索引擎。
* [boxed](https://github.com/tejo/boxed) **star:72** 基于Dropbox的博客引擎。   [![最近一年没有更新][Yellow]](https://github.com/tejo/boxed)   [![godoc][GoDoc]](https://godoc.org/github.com/tejo/boxed)
* [Cherry](https://github.com/rafael-santiago/cherry)  微型网络聊天服务器。
* [drive](https://github.com/odeke-em/drive)  基于命令行的谷歌驱动器客户端。
* [dp](https://github.com/scryinfo/dp) **star:57** 通过SDK与区块链进行数据交换，开发人员可以轻松访问DAPP开发。   [![godoc][GoDoc]](https://godoc.org/github.com/scryinfo/dp)
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** 基于加密管的简单的NaCL EC25519工具。   [![最近一年没有更新][Yellow]](https://github.com/unix4fun/naclpipe)   [![godoc][GoDoc]](https://godoc.org/github.com/unix4fun/naclpipe)
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** 测试你的终端。   [![最近一年没有更新][Yellow]](https://github.com/crazcalm/term-quiz)   [![godoc][GoDoc]](https://godoc.org/github.com/crazcalm/term-quiz)
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** 当有人通过 Tsuru 部署任何应用程序时，会通知您的团队以及其他工具。   [![最近一年没有更新][Yellow]](https://github.com/lucasgomide/snitch)   [![godoc][GoDoc]](https://godoc.org/github.com/lucasgomide/snitch)
* [hugo](http://gohugo.io/)  快速、现代的静态网站引擎。
* [Gor](https://github.com/buger/gor)  Http 流量复制工具，用于实时回放从生产环境到阶段/开发环境的流量。
* [GoLand](https://jetbrains.com/go)  功能齐全的跨平台 Go IDE。
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:11** 包含了 Go 使用手册文档的 Chrome 扩展。   [![最近一年没有更新][Yellow]](https://github.com/diankong/GoDocTooltip)

# 资源

*在哪里可以找到新的Go库。 (翻译出错了? 试试 [英文版](README_EN.md#resources) 吧~)*

## 基准

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1344** HTTP请求路由器基准测试和比较。   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1128** web框架基准测试。   [![最近一周有更新][Green]](https://github.com/smallnest/go-web-framework-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
* [skynet](https://github.com/atemerev/skynet) **star:947** 天网 1M 线程微基准测试。
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:946** Go序列化方法的基准测试。   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Go语言常用基本操作的基准测试。
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:189** 对比各种图像大小调整算法性能。   [![godoc][GoDoc]](https://godoc.org/github.com/fawick/speedtest-resize)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:131** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches   [![最近一年没有更新][Yellow]](https://github.com/tylertreat/go-benchmarks)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/go-benchmarks)
* [autobench](https://github.com/davecheney/autobench) **star:90** 用来来比较不同Go版本之间的性能的框架。   [![最近一年没有更新][Yellow]](https://github.com/davecheney/autobench)   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/autobench)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:52** 为流行的 Go 数据库/SQL实用程序收集基准测试。   [![最近一年没有更新][Yellow]](https://github.com/tyler-smith/golang-sql-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/tyler-smith/golang-sql-benchmark)
* [gospeed](https://github.com/feyeleanor/GoSpeed)  计算语言结构的速度的微观基准测试。
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:20** 强大的HTTP基准测试工具，包含了Аb，Wrk，Siege工具。收集统计和各种参数指标，并比较相关结果。   [![最近一年没有更新][Yellow]](https://github.com/mrLSD/go-benchmark-app)   [![godoc][GoDoc]](https://godoc.org/github.com/mrLSD/go-benchmark-app)
* [kvbench](https://github.com/jimrobinson/kvbench) **star:16** K / V 类型数据库基准测试。   [![godoc][GoDoc]](https://godoc.org/github.com/jimrobinson/kvbench)

## 会议

* [Capital Go](http://www.capitalgolang.com)  华盛顿特区。,美国。
* [GopherCon Europe](https://gophercon.is/)  德国柏林。
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
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:13** in Persian.   [![最近一年没有更新][Yellow]](https://github.com/thedevsir/gosuccinctly)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsir/gosuccinctly)
* [GoBooks](https://github.com/dariubs/GoBooks)  一份精选的 Go 书籍清单。
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1804** 由 Maria Letta 提供的与 Gopher 有关的图片包，其中包含了插图,表情文字。   [![godoc][GoDoc]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:71** 可爱的 gopher 标识。   [![最近一年没有更新][Yellow]](https://github.com/GolangUA/gopher-logos)
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:32** 与 Go gopher 相关的媒介数据[。ai . svg)。   [![最近一年没有更新][Yellow]](https://github.com/keygx/Go-gopher-Vector)
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:2061** 阿什莉·麦克纳马拉的歌斐艺术品。   [![star > 2000][Awesome]](https://github.com/ashleymcnamara/gophers)   [![godoc][GoDoc]](https://godoc.org/github.com/ashleymcnamara/gophers)
* [gophers](https://github.com/egonelbre/gophers) **star:1878** Free gophers.   [![godoc][GoDoc]](https://godoc.org/github.com/egonelbre/gophers)
* [gopherize.me](https://github.com/matryer/gopherize.me)  Gopherize自己。
* [gophers](https://github.com/rogeralsing/gophers) **star:51** 随机gopher图形。   [![最近一年没有更新][Yellow]](https://github.com/rogeralsing/gophers)
* [gophers](https://github.com/sillecelik/go-gopher) **star:43** Gopher amigurumi玩具图案。

## 聚会

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
* [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
* [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBandung](https://www.meetup.com/GoBandung/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
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
* [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
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
* [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
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

## 社交媒体
### Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

### Reddit
 * [r/golang](https://www.reddit.com/r/golang/)

## 网站

* [Go Projects](https://github.com/golang/go/wiki/Projects)  wiki上的 Go 社区项目列表。
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:25678** 其他 awesome 系列的列表。   [![star > 2000][Awesome]](https://github.com/bayandin/awesome-awesomeness)   [![最近一周有更新][Green]](https://github.com/bayandin/awesome-awesomeness)
* [CodinGame](https://www.codingame.com/)  以小游戏互动完成任务的形式来学习 Go。
* [Go Blog](http://blog.golang.org)  官方 Go 博客。
* [Go Challenge](http://golang-challenge.org/)  通过解决问题并从 Go 专家那里得到反馈来学习 Go。
* [Go Community on Hashnode](https://hashnode.com/n/go)  Hashnode上的Go社区。
* [Go Forum](https://forum.golangbridge.org)  讨论 Go 的论坛。
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Report Card](https://goreportcard.com)  为你的 Go 包生成一份报告单。
* [go.dev](https://go.dev/)  一个围棋开发者的中心。
* [Gophercises](https://gophercises.com/)  为 Go 初学者提供免费的代码训练。
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  Google+社区 golang爱好者聚集地。
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  加入我们为Gophers设立的全新Slack社区([了解它是如何产生的](https://blog.gopheracademy.com/gophers-slack-community/))。
* [justforfunc](https://www.youtube.com/c/justforfunc)  致力于传授 Go 编程语言技巧和技巧的Youtube 频道，由Francesc Campoy [@francesc](https://twitter.com/francesc)主办。
* [gowalker.org](https://gowalker.org)   Go API 文档。
* [json2go](https://m-zajac.github.io/json2go)  高级JSON去结构转换-在线工具。
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go 邮件列表。
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:142** 收藏的 Go 图像，图形和艺术作品。   [![最近一年没有更新][Yellow]](https://github.com/mholt/golang-graphics)   [![归档项目][Archived]](https://github.com/mholt/golang-graphics)
* [Awesome Go @LibHunt](https://go.libhunt.com)  属于你的 Go 工具箱。
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job)  Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [Golang News](https://golangnews.com)  关于 Go 编程的链接和新闻。
* [Golang Flow](https://golangflow.io)  发布更新、新闻、包等等。
* [Golang Developer Jobs](https://golangjob.xyz)  开发人员的工作专为Golang相关的角色。
* [godoc.org](https://godoc.org/)  开源 Go 包的文档。
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:37** 专门收集需要帮助的Go项目，这是你开启开源之路的好去处。   [![最近一年没有更新][Yellow]](https://github.com/ninedraft/gocryforhelp)
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  与 Go 有关的新闻。
* [studygolang](https://studygolang.com)  Go语言中文网
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  寻找最新的 Go库 的好地方。
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### 教程

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:33808** Golang电子书，主要讲述如何用 Golang 建立一个web应用程序。   [![star > 2000][Awesome]](https://github.com/astaxie/build-web-application-with-golang)   [![最近一周有更新][Green]](https://github.com/astaxie/build-web-application-with-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   [![包含中文文档][CN]](https://github.com/astaxie/build-web-application-with-golang)
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  如何缓存数据库的慢查询。
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  如何取消MySQL查询。
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:554** 一本讲述如何用 Go 进行以太开发的小册。   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/ethereum-development-with-go-book)   [![包含中文文档][CN]](https://github.com/miguelmota/ethereum-development-with-go-book)
* [Games With Go](http://gameswithgo.org/)  关于编程和游戏开发系列教学视频。
* [Go By Example](https://gobyexample.com/)  手把手教你 如何在 Go 应用程序中使用注释。
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)   Go's reference card。
* [Go database/sql tutorial](http://go-database-sql.org/)  数据库概论/ sql。
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  在你的移动设备上编辑你编辑和运行你的 Go 代码。
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Go 初学者经常遇到的陷阱、误区、错误
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  教你如何用 Go 搭建一个电商平台 (包括demo)。
* [A Tour of Go](http://tour.golang.org/)  互动的 Go 之旅。
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns) **star:12156** 策划的清单去设计模式，食谱和习惯用法。   [![star > 2000][Awesome]](https://github.com/tmrts/go-patterns)   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/go-patterns)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:10159** 学习使用测试驱动开发。   [![star > 2000][Awesome]](https://github.com/quii/learn-go-with-tests)   [![godoc][GoDoc]](https://godoc.org/github.com/quii/learn-go-with-tests)   [![包含中文文档][CN]](https://github.com/quii/learn-go-with-tests)
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  面向 Golang 初学者教程。
* [package main](https://www.youtube.com/packagemain)  关于 Go 编程的YouTube频道。
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera的专业学习可以从零开始。
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:1176** 引入示例讲述 Golang 与Node.js在学习上的差异。   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)
* [Golangbot](https://golangbot.com/learn-golang-series/)  Go 编程教程。
* [GolangCode](https://golangcode.com/)  收集代码片段和教程，以帮助处理日常问题。
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Go社区投票选举出来的最好的在线 Go 教程。
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  快速使用Godog —— 一个行为驱动开发的构建和测试应用程序框架。
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1151** 由一个经验丰富的Go程序员群体编写的一系列Go学习范例。
* [Your basic Go](http://yourbasic.org/golang)  如何收集大量的教程。

