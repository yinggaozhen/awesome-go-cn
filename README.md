# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "最近一周有更新"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "最近一年没有更新"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "包含中文文档"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "项目已归档"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc文档地址"

**此项目是 [awesome-go](https://awesome-go.com/) 中文版，最后一次同步时间 : 2020-03-18 12:22:08(每隔1天同步一次)**

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

* [Oto](https://github.com/hajimehoshi/oto) **star:549** 多平台的 low-level 声音播放库。   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:339** 基于 Go 的PortAudio audio I/O库。   [![最近一年没有更新][Yellow]](https://github.com/gordonklaus/portaudio)   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:286** 基于 Go 的音乐理论模型。   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:276** 通过音频流生成波形图像的包。   [![最近一年没有更新][Yellow]](https://github.com/mdlayher/waveform)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:227** PortMidi的 Go 语言实现接口.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:146** 快速稳定的 ID3 解析及写入Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [flac](https://github.com/mewkiz/flac) **star:118** 原生 Go FLAC编码器/解码器，支持FLAC流。   [![godoc][GoDoc]](https://godoc.org/github.com/mewkiz/flac)
* [mix](https://github.com/go-mix/mix) **star:112** 基于序列的 Go 原生音乐混音器。   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:107** 原生 Go MP3解码器。   [![最近一年没有更新][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:105** libsox 的 Go 语言实现接口。   [![最近一年没有更新][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [malgo](https://github.com/gen2brain/malgo) **star:96** 迷你音频库。   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/malgo)
* [taglib](https://github.com/wtolson/go-taglib) **star:72** taglib 的 Go 语言实现接口.   [![最近一年没有更新][Yellow]](https://github.com/wtolson/go-taglib)   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:70** 原生 Go AAC位流解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:39** 轻量级MP3解码器库。
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:28** libmediainfo 的 Go 语言实现接口。   [![最近一年没有更新][Yellow]](https://github.com/zhulik/go_mediainfo)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/go_mediainfo)
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:28** EasyMidi是一个简单可靠的库，用于处理标准midi文件(SMF)。   [![最近一年没有更新][Yellow]](https://github.com/algoGuy/EasyMIDI)   [![godoc][GoDoc]](https://godoc.org/github.com/algoGuy/EasyMIDI)
* [vorbis](https://github.com/mccoyst/vorbis) **star:25** “原生” Go Vorbis解码器(使用CGO，但没有依赖关系)。   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:9** libsamplerate 的 Go 语言实现接口。   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)

## 身份验证和OAuth

*用于实现验证方案的库。 (翻译出错了? 试试 [英文版](README_EN.md#authentication-and-oauth) 吧~)*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:7233** JSON Web令牌(JWT)。   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![最近一周有更新][Green]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:6260** 支持ACL、RBAC、ABAC等访问控制模型的授权库。   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![最近一周有更新][Green]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [oauth2](https://github.com/golang/oauth2) **star:2771** goauth2的继任者。通用OAuth 2.0包，附带JWT、谷歌api、计算引擎和应用程序引擎支持。   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2582** 提供了 OAuth 和 OAuth2 的简单清晰易用的方法。可开箱即用处理多个提供程序。   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![最近一周有更新][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:2167** web模块化认证系统。它试图删除尽可能多的模板文件和硬编码，以便每次新建一个新的web项目时，您都可以插入、配置并开始构建您的应用程序，而不必每次都构建一个身份验证系统。   [![star > 2000][Awesome]](https://github.com/volatiletech/authboss)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [osin](https://github.com/openshift/osin) **star:1585** OAuth2服务器库。   [![godoc][GoDoc]](https://godoc.org/github.com/openshift/osin)
* [go-jose](https://github.com/square/go-jose) **star:1489** 相当完整地实现了JOSE工作组的JSON Web令牌、JSON Web签名和JSON Web加密规范。   [![godoc][GoDoc]](https://godoc.org/github.com/square/go-jose)
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1466** 用 Golang 编写的独立且符合规范的OAuth2服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [gologin](https://github.com/dghubble/gologin) **star:1171** 用于使用OAuth1和OAuth2身份验证提供者登录的可链处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:996** 轻量级的基于角色的访问控制(RBAC)实现。   [![最近一年没有更新][Yellow]](https://github.com/mikespook/gorbac)   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [loginsrv](https://github.com/tarent/loginsrv) **star:919** JWT登录微服务带有可插拔的后端服务，如OAuth2 (Github)、htpasswd、osiam。   [![最近一周有更新][Green]](https://github.com/tarent/loginsrv)   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [scs](https://github.com/alexedwards/scs) **star:683** HTTP服务器的会话管理器。   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [permissions2](https://github.com/xyproto/permissions2) **star:393** 用于跟踪用户、登录状态和权限的库。依赖于cookie安全和bcrypt。   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [paseto](https://github.com/o1egl/paseto) **star:319** 平台无关的安全令牌(PASETO)。   [![最近一周有更新][Green]](https://github.com/o1egl/paseto)   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [httpauth](https://github.com/goji/httpauth) **star:191** HTTP身份验证中间件。   [![最近一年没有更新][Yellow]](https://github.com/goji/httpauth)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:190** 简单、灵活、安全和惯用的web会话管理，具有可配置化的后端。   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:172** JWT中间件，可用于Golang http服务器，提供了许多配置选项。   [![最近一年没有更新][Yellow]](https://github.com/adam-hanna/jwt-auth)   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/pascaldekloe/jwt) **star:159** 轻量级JSON Web令牌库。   [![最近一周有更新][Green]](https://github.com/pascaldekloe/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [branca](https://github.com/hako/branca) **star:103** 基于 Go 实现Branca令牌。   [![最近一周有更新][Green]](https://github.com/hako/branca)   [![godoc][GoDoc]](https://godoc.org/github.com/hako/branca)
* [session](https://github.com/icza/session) **star:99** web服务器会话管理(包括支持谷歌应用程序引擎 - GAE)。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [sessionup](https://github.com/swithek/sessionup) **star:82** 简单但有效的HTTP会话管理和标识包。   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [jwt](https://github.com/robbert229/jwt) **star:78** 简单易用的JSON Web令牌实现(JWT)。   [![最近一年没有更新][Yellow]](https://github.com/robbert229/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sjwt](https://github.com/brianvoe/sjwt) **star:61** 简单的jwt生成器和解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [sessions](https://github.com/adam-hanna/sessions) **star:48** 非常简单，高性能，可深度定制的会话服务，主要用于的 go http 服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [rbac](https://github.com/zpatrick/rbac) **star:46** 最小的RBAC包。   [![最近一年没有更新][Yellow]](https://github.com/zpatrick/rbac)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [securecookie](https://github.com/chmike/securecookie) **star:33** 高效安全的cookie编码/解码。   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** 使用SessionGate Redis模块进行会话管理。   [![最近一年没有更新][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [signedvalue](https://github.com/sashka/signedvalue) **star:8** 与[Tornado's](https://github.com/tornadooweb/tornado) 完全兼容的签名和时间戳字符串实现。   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)
* [scope](https://github.com/SonicRoshan/scope) **star:4** 在Go中轻松管理OAuth2范围。   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** 提供cookie .txt文件格式的解析器。   [![最近一年没有更新][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)

## Bot建设

*用于构建和使用机器人的库。 (翻译出错了? 试试 [英文版](README_EN.md#bot-building) 吧~)*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1977** 简单轻量级的Telegram bot客户端。   [![最近一周有更新][Green]](https://github.com/Syfaro/telegram-bot-api)   [![godoc][GoDoc]](https://godoc.org/github.com/Syfaro/telegram-bot-api)
* [telebot](https://github.com/tucnak/telebot) **star:1149** 用Go编写的Telegram bot框架。   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:555** 用 Go 编写的IRC, Slack和电报机器人。   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [go-joe](https://joe-bot.net)  一个通用的机器人库的灵感来自于Hubot，但写在围棋。
* [slacker](https://github.com/shomali11/slacker) **star:360** 可简单创建Slack机器人的框架。   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:290** 基于控制台的，用于加密货币交易所的的交易机器人。   [![最近一年没有更新][Yellow]](https://github.com/saniales/golang-crypto-trading-bot)   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [Kelp](https://github.com/stellar/kelp) **star:266** 官方交易和做市机器人为[Stellar](https://www.stellar.org/) DEX。开箱即用的作品，用 Golang 编写，兼容集中交易和定制交易策略。   [![最近一周有更新][Green]](https://github.com/stellar/kelp)   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [tbot](https://github.com/yanzay/tbot) **star:251** 带有类似于net/http API的Telegram bot服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:167** 面向服务的IRC bot，使用Redis和JSON进行消息传递。   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:156** 此框架提供了聊天机器人相关的服务，包括LINE、Slack、Gitter等。   [![godoc][GoDoc]](https://godoc.org/github.com/oklahomer/go-sarah)
* [hanu](https://github.com/sbstjn/hanu) **star:119** 用于编写Slack机器人的框架。   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:96** Libary为twitch编写机器人程序。电视聊天   [![最近一周有更新][Green]](https://github.com/gempir/go-twitch-irc)   [![godoc][GoDoc]](https://godoc.org/github.com/gempir/go-twitch-irc)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:91** 由swagger文件、基于会话的路由器和中间件生成的纯Golang Telegram Bot API包装器。   [![最近一年没有更新][Yellow]](https://github.com/olebedev/go-tgbot)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:60** 构建电报机器人的框架。   [![最近一年没有更新][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:30** 简单的Go [VK](https://vk.com) bot库。   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:24** 另一个构建Slack机器人的框架。   [![最近一周有更新][Green]](https://github.com/alexandre-normand/slackscot)   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) **star:18** 用于管理基于语音通道成员存在的临时角色的不和机器人。   [![最近一周有更新][Green]](https://github.com/ewohltman/ephemeral-roles)   [![godoc][GoDoc]](https://godoc.org/github.com/ewohltman/ephemeral-roles)
* [micha](https://github.com/onrik/micha) **star:12** 基于 GO 实现的Telegram 机器人API库。   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)

## 命令行

### 标准CLI

*用于构建标准或基本命令行应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#standard-cli) 吧~)*

* [cobra](https://github.com/spf13/cobra) **star:16135** 现代Go CLI命令行交互工具。   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![最近一周有更新][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:13333** 可让你简单、快速和愉快的构建命令行应用(之前是codegangsta/cli)。   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![最近一周有更新][Green]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2820** 支持子命令的命令行和标志解析器。   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1652**  Go 命令行选项解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [Dnote](https://github.com/dnote/dnote) **star:1621** 一个简单的命令行笔记本与多设备同步。   [![最近一周有更新][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [readline](https://github.com/chzyer/readline) **star:1472** 纯golang实现，在MIT许可下提供了GNU-Readline的大部分特性。   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [docopt.go](https://github.com/docopt/docopt.go) **star:1225** 会让你满意的命令行参数解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1100** 用于实现命令行接口的Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [pflag](https://github.com/spf13/pflag) **star:1025** 基于POSIX/GNU-style --flags实现的包，主要用于替换Go的falg包。   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [cli-init](https://github.com/tcnksm/gcli) **star:894** 一个简单就可开启构建Golang命令行的应用程序。   [![最近一年没有更新][Yellow]](https://github.com/tcnksm/gcli)   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [go-arg](https://github.com/alexflint/go-arg) **star:846** 基于结构的参数解析。   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:676** 使用 Go 语言编写的 bash 命令补全工具以及 Go 命令补全工具.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [liner](https://github.com/peterh/liner) **star:667** 类似readline-like的命令行接口库。   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [mow.cli](https://github.com/jawher/mow.cli) **star:658** 用于构建具有复杂标志和参数解析和验证的CLI应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [flaggy](https://github.com/integrii/flaggy) **star:634** 一个健壮的、易用的标志包，具有出色的子命令支持。   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:518** 基于golang结构标签，功能丰富易于使用的命令行包。   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:387** Unikernel 构建器/协调器。   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:176** 命令行参数分析器，灵感来自Python的argparse模块。   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [sflags](https://github.com/octago/sflags) **star:110** 基于结构的flag生成器，用于flag、urfave/cli、pflag、cobra、kingpin和其他库。   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [commandeer](https://github.com/jaffee/commandeer) **star:109** 开发友好的CLI应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [wmenu](https://github.com/dixonwille/wmenu) **star:102** 为cli程序提供了简单上手的菜单，可提示用户作出选择。   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [flag](https://github.com/cosiner/flag) **star:102** 简单但功能强大的命令行选项解析库，用于支持Go子命令。   [![最近一年没有更新][Yellow]](https://github.com/cosiner/flag)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:101** 简小的命令行接口框架。   [![最近一年没有更新][Yellow]](https://github.com/ukautz/clif)   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [job](https://github.com/liujianping/job) **star:70** 工作，把你的短期指令当作长期任务。   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [cli](https://github.com/teris-io/cli) **star:66** 为 Go 构建命令接口提供简单而完整的API。   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [1build](https://github.com/gopinath-langote/1build) **star:55** 命令行工具，以无摩擦地管理项目特定的命令。   [![最近一周有更新][Green]](https://github.com/gopinath-langote/1build)   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [env](https://github.com/codingconcepts/env) **star:48** 基于标记的结构化的环境配置。   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [wlog](https://github.com/dixonwille/wlog) **star:42** 支持跨平台和并发的简单日志记录接口。   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wlog)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  具有自动配置和依赖注入的cli应用程序框架。
* [gocmd](https://github.com/devfacet/gocmd) **star:37** 用于构建命令行应用程序。   [![最近一年没有更新][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [strumt](https://github.com/antham/strumt) **star:34** 用于创建提示链。   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [flagvar](https://github.com/sgreben/flagvar) **star:33** 符合 Go 标准的“flag”标志参数类型包。   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [cmdr](https://github.com/hedzr/cmdr) **star:30** 一个POSIX/GNU风格的、类似getopt的命令行UI Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [clîr](https://github.com/leaanthony/clir) **star:21** 一个简单而清晰的CLI库。依赖免费的。   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:19**  Go 选择解析器，借鉴于Perl灵活性的GetOpt::Long。   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [argv](https://github.com/cosiner/argv) **star:19** 基于Base 语法，用于分隔命令行字符串并将其作为参数的 Go 语言库，   [![最近一周有更新][Green]](https://github.com/cosiner/argv)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** 用于简化CLI工作流的 Go 库。   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [sand](https://github.com/Zaba505/sand) **star:9** 用于创建解释器等的简单API。   [![最近一年没有更新][Yellow]](https://github.com/Zaba505/sand)   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [ts](https://github.com/liujianping/ts) **star:8** 时间戳转换和比较工具。   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/ts)
* [cmd](https://github.com/posener/cmd) **star:6** 扩展标准的' flag '包，以支持子命令和更多的idomatic方式。   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)

### 高级控制台用户界面

*用于构建控制台应用程序和控制台用户界面的库。 (翻译出错了? 试试 [英文版](README_EN.md#advanced-console-uis) 吧~)*

* [termui](https://github.com/gizak/termui) **star:9621** 此库是基于**termbox-go**实现的，借鉴于[blessed-contrib](https://github.com/yaronn/blessed-contrib)。   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  更换终端文本样式。
* [gocui](https://github.com/jroimartin/gocui) **star:6052** 旨在创建控制台用户界面的极简Go库。   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [termbox-go](https://github.com/nsf/termbox-go) **star:3700** 基于文本的跨平台接口库。   [![star > 2000][Awesome]](https://github.com/nsf/termbox-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/termbox-go)
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2881** 构建一个强大的交互式提示，借鉴于[python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit)   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1647** 在终端呈现进度条，可灵活配置的。   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1315** 在命令行中构建轻量级ASCII线图╭┈╯，应用程序中没有其他依赖项。   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:1110** 用于实时更新终端输出的库。   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [termdash](https://github.com/mum4k/termdash) **star:965** 此库是基于**termbox-go**实现的，借鉴于[termui](https://github.com/gizak/termui)。   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [progressbar](https://github.com/schollz/progressbar) **star:869** 基本线程安全的进度条，在每个操作系统工作。   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/progressbar)
* [mpb](https://github.com/vbauerster/mpb) **star:856** 可在终端显示多进度条。   [![最近一周有更新][Green]](https://github.com/vbauerster/mpb)   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [aurora](https://github.com/logrusorgru/aurora) **star:790** 支持fmt.Printf/Sprintf的ANSI终端颜色。   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [uitable](https://github.com/gosuri/uitable) **star:547** 改善终端应用程序中表格数据的可读性。   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:419** 适用于windows的颜色编写器。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:405** Go 实现的 isatty。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [gookit/color](https://github.com/gookit/color) **star:329** 终端显色工具库，支持16种颜色，256种颜色，RGB显色输出，兼容Windows。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![包含中文文档][CN]](https://github.com/gookit/color)
* [chalk](https://github.com/ttacon/chalk) **star:329** 美化终端/控制台输出。   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [tabby](https://github.com/cheynewallace/tabby) **star:264** 一个可在终端生成一个极简Golang表格轻量级库   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [simpletable](https://github.com/alexeyco/simpletable) **star:216** 可在终端显示简易表格。   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:201** 在终端中使用彩色文字。   [![最近一年没有更新][Yellow]](https://github.com/daviddengcn/go-colortext)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [cfmt](https://github.com/mingrammer/cfmt) **star:72** 提供上下文的fmt，灵感来自于bootstrap color classes。   [![最近一年没有更新][Yellow]](https://github.com/mingrammer/cfmt)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:36** 不需要向API传递大量参数就可从命令行实用程序中打印ASCII表。   [![最近一年没有更新][Yellow]](https://github.com/InVisionApp/tabular)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [ctc](https://github.com/wzshiming/ctc) **star:17** 不需要Print方法的非侵入性跨平台终端颜色库。   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   [![包含中文文档][CN]](https://github.com/wzshiming/ctc)
* [colourize](https://github.com/TreyBastian/colourize) **star:17** 在终端提供ANSI彩色文本。   [![最近一年没有更新][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** 在终端提供ANSI彩色文本模板。   [![最近一年没有更新][Yellow]](https://github.com/workanator/go-ataman)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)

## 配置

*配置解析的库。 (翻译出错了? 试试 [英文版](README_EN.md#configuration) 吧~)*

* [viper](https://github.com/spf13/viper) **star:11502** 配置管理。   [![star > 2000][Awesome]](https://github.com/spf13/viper)   [![最近一周有更新][Green]](https://github.com/spf13/viper)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2822** 管理来自环境变量的配置数据。   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv) **star:2685** Ruby 的 dotenv 库的 Go移植版(从.env文件加载环境变量)。   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![最近一周有更新][Green]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [ini](https://github.com/go-ini/ini) **star:1959**  读和写INI文件。   [![最近一周有更新][Green]](https://github.com/go-ini/ini)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:1448** 解析环境变量并赋值到struct中(默认值)。   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [konfig](https://github.com/lalamove/konfig) **star:548** 可组合、可观察和高性能的分布式配置管理。   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [confita](https://github.com/heetch/confita) **star:296** 从多个后端级联加载配置到struct中。   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:248** 轻量级配置管理器。   [![最近一年没有更新][Yellow]](https://github.com/tucnak/store)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/JeremyLoy/config) **star:231** 云本地应用程序配置。将ENV绑定到结构体仅需两行代码。   [![最近一周有更新][Green]](https://github.com/JeremyLoy/config)   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [config](https://github.com/olebedev/config) **star:225** 带有环境变量和标记解析的JSON或YAML配置包装器。   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/config)
* [joshbetz/config](https://github.com/joshbetz/config) **star:198** 一个可解析环境变量、JSON文件小巧的配置库，在SIGHUP时会自动重新加载。   [![godoc][GoDoc]](https://godoc.org/github.com/joshbetz/config)
* [hjson](https://github.com/hjson/hjson-go) **star:195** 更加人性化的JSON配置。轻松的语法，更少的错误，更多的注释。   [![最近一周有更新][Green]](https://github.com/hjson/hjson-go)   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [envconfig](https://github.com/vrischmann/envconfig) **star:172** 从环境变量中读取配置。   [![godoc][GoDoc]](https://godoc.org/github.com/vrischmann/envconfig)
* [koanf](https://github.com/knadh/koanf) **star:170** 轻量级可扩展库，用于读取Go应用程序中的配置。内置支持JSON, TOML, YAML, env，命令行。   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [gookit/config](https://github.com/gookit/config) **star:139** 程序配置管理(load,get,set)。支持JSON, YAML, TOML, INI, HCL。支持多文件加载，数据覆盖合并。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   [![包含中文文档][CN]](https://github.com/gookit/config)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:127** 将ini的配置文件读入 Go structs中;支持用户定义的类型和子选项。   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [goConfig](https://github.com/crgimenes/goConfig) **star:125** 将结构体解析为输入，并用来自命令行、环境变量和配置文件的参数填充该结构体的字段。   [![最近一周有更新][Green]](https://github.com/crgimenes/goConfig)   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [envh](https://github.com/antham/envh) **star:94** 协助管理环境变量的Helpers。   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [envcfg](https://github.com/tomazk/envcfg) **star:91** 对环境变量进行解析，并赋值到struct。   [![最近一年没有更新][Yellow]](https://github.com/tomazk/envcfg)   [![godoc][GoDoc]](https://godoc.org/github.com/tomazk/envcfg)
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:58** 简约的配置阅读器(来自文件、环境，以及你想要的任何地方)。   [![godoc][GoDoc]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  模块化的JSON配置。保持配置结构及其配置的代码，并将解析委托给子模块，而不牺牲配置的完整序列化。
* [gofigure](https://github.com/ian-kent/gofigure) **star:58** 让程序配置变得简单。   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [configure](https://github.com/paked/configure) **star:54** 通过多个源提供配置，包括JSON、flags和环境变量。   [![最近一年没有更新][Yellow]](https://github.com/paked/configure)   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [harvester](https://github.com/beatlabs/harvester) **star:52** 一个易于使用的静态和动态配置包   [![最近一周有更新][Green]](https://github.com/beatlabs/harvester)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/harvester)
* [config](https://github.com/golobby/config) **star:51** 一个轻量级但功能强大的配置包，用于Go项目。   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:48** 遵循[XDG标准](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)的跨平台包。   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [ingo](https://github.com/schachmat/ingo) **star:28** flag保存在类ini的配置文件中。   [![最近一年没有更新][Yellow]](https://github.com/schachmat/ingo)   [![godoc][GoDoc]](https://godoc.org/github.com/schachmat/ingo)
* [go-up](https://github.com/ufoscout/go-up) **star:26** 一个简单的配置库，具有递归占位符解析功能。   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:22** 用于解析ini类型的配置文件。   [![最近一年没有更新][Yellow]](https://github.com/sasbury/mini)   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:13** 合并来自任意url的多个JSON/YAML/TOML文件、针对JSON模式的验证以及模式中定义的默认值的应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/the4thamigo-uk/conflate)
* [genv](https://github.com/sakirsensoy/genv) **star:9** 使用dotenv支持轻松读取环境变量。   [![godoc][GoDoc]](https://godoc.org/github.com/sakirsensoy/genv)
* [envconf](https://github.com/ian-kent/envconf) **star:9** 从环境配置中读取配置。   [![最近一年没有更新][Yellow]](https://github.com/ian-kent/envconf)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/envconf)
* [sprbox](https://github.com/oblq/sprbox) **star:5** 支持构建环境的工具箱工厂和其他不确定的配置解析器(如YAML、TOML、JSON和环境vars)。   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/sprbox)
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) **star:3** 从AWS SSM(参数存储)加载配置参数的Go实用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-ssm-config)
* [configuration](https://github.com/BoRuDar/configuration) **star:3** 库，用于从环境变量、文件、标记和“默认”标记初始化配置结构。   [![最近一周有更新][Green]](https://github.com/BoRuDar/configuration)   [![godoc][GoDoc]](https://godoc.org/github.com/BoRuDar/configuration)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:2** 读取环境变量的简单有用的包。   [![godoc][GoDoc]](https://godoc.org/github.com/nasermirzaei89/env)
* [onion](http://github.com/goraz/onion)  基于层配置的Go，支持JSON, TOML, YAML，属性，etcd, env，和加密使用PGP。

## 持续集成

*用于帮助进行持续集成的工具。 (翻译出错了? 试试 [英文版](README_EN.md#continuous-integration) 吧~)*

* [drone](https://github.com/drone/drone) **star:20682** Drone 是一个基于 Docker 的持续集成平台，用 Go 编写。   [![star > 2000][Awesome]](https://github.com/drone/drone)   [![最近一周有更新][Green]](https://github.com/drone/drone)   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [CDS](https://github.com/ovh/cds) **star:2744** 企业级CI/CD和DevOps自动化开源平台。   [![star > 2000][Awesome]](https://github.com/ovh/cds)   [![最近一周有更新][Green]](https://github.com/ovh/cds)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/cds)
* [goveralls](https://github.com/mattn/goveralls) **star:634** Coveralls.io 是一个用 Go 编写，可持续对代码覆盖率进行检测的系统。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/goveralls)
* [overalls](https://github.com/go-playground/overalls) **star:101** 针对多package 的 Go 语言项目，可为类似 goveralls 这样的工具生成覆盖率报告。   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/overalls)
* [duci](https://github.com/duck8823/duci) **star:53** 一个简单的 ci 服务。   [![最近一周有更新][Green]](https://github.com/duck8823/duci)   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:44** 在一个干净的工作区中对你的 Go 二进制文件进行测试、构建、签名和发布。   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** 递归覆盖测试工具。   [![最近一年没有更新][Yellow]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

## CSS预处理器

*用于预处理CSS文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#css-preprocessors) 吧~)*

* [gcss](https://github.com/yosssi/gcss) **star:431** 纯Go编写的 CSS 预处理器。   [![最近一年没有更新][Yellow]](https://github.com/yosssi/gcss)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass) **star:155**  采用 Go封装，100% 与 Sass 兼容的 libsass 项目。   [![最近一年没有更新][Yellow]](https://github.com/wellington/go-libsass)

## 数据结构

*用 Go 实现的通用的数据结构和算法。 (翻译出错了? 试试 [英文版](README_EN.md#data-structures) 吧~)*

* [gods](https://github.com/emirpasic/gods) **star:7918** 数据结构。容器、集合、列表、堆栈、地图、BidiMaps、树、HashSet等。   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5481** 可靠的、高性能的和线程安全的数据结构的集合。   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1423** 线程安全和非线程安全的高性能集。   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1236** 用于处理连续的概率数据结构。   [![最近一年没有更新][Yellow]](https://github.com/tylertreat/BoomFilters)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [gota](https://github.com/kniren/gota) **star:1107** 实现了数据帧，序列以及数据噪音。   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [roaring](https://github.com/RoaringBitmap/roaring) **star:797** 实现了压缩 bitsets 的Go包。   [![godoc][GoDoc]](https://godoc.org/github.com/RoaringBitmap/roaring)
* [willf/bloom](https://github.com/willf/bloom) **star:759** 实现Bloom过滤器。   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bloom)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:691** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:593** 布谷鸟过滤器:一个用Go实现，可替代计数 bloom 过滤器。   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [bitset](https://github.com/willf/bitset) **star:539** 实现了 bitsets 的 Go 包。   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [trie](https://github.com/derekparker/trie) **star:463** 在Go中实现Trie。   [![最近一周有更新][Green]](https://github.com/derekparker/trie)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/trie)
* [gocache](https://github.com/eko/gocache) **star:375** 一个完整的Go缓存库，具有多个存储(内存，memcache, redis，…)，可链，可加载，指标缓存和更多。   [![最近一周有更新][Green]](https://github.com/eko/gocache)   [![godoc][GoDoc]](https://godoc.org/github.com/eko/gocache)
* [algorithms](https://github.com/shady831213/algorithms) **star:364** 算法和数据结构。来源于CLRS。   [![godoc][GoDoc]](https://godoc.org/github.com/shady831213/algorithms)
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:320** 基于内存的地理索引。   [![最近一年没有更新][Yellow]](https://github.com/hailocab/go-geoindex)   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [mafsa](https://github.com/smartystreets/mafsa) **star:281** 实现了 MA-FSA ，包含最小完美哈希。   [![godoc][GoDoc]](https://godoc.org/github.com/smartystreets/mafsa)   [![归档项目][Archived]](https://github.com/smartystreets/mafsa)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:217** 基于 Go 的跳表实现。   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [hilbert](https://github.com/google/hilbert) **star:197** 用于映射空间填充曲线（例如 Hilbert 曲线和 Peano 曲线）和数值。   [![最近一年没有更新][Yellow]](https://github.com/google/hilbert)   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [merkletree](https://github.com/cbergoon/merkletree) **star:182** 实现了merkle树，提供了对数据结构内容的有效和安全的验证。   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:135** 帮助用户构建自定义二进制流的二进制封装器和解包器   [![最近一年没有更新][Yellow]](https://github.com/zhuangsirui/binpacker)   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:133** 自适应基数树。   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [bloom](https://github.com/zhenjl/bloom) **star:131** 在Go中实现了Bloom过滤器。   [![最近一年没有更新][Yellow]](https://github.com/zhenjl/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:129** 基于内存的LRU算法实现。   [![godoc][GoDoc]](https://godoc.org/github.com/diegobernardes/ttlcache)
* [iter](https://github.com/disksing/iter) **star:118** Go实现的c++ STL迭代器和算法。   [![godoc][GoDoc]](https://godoc.org/github.com/disksing/iter)   [![包含中文文档][CN]](https://github.com/disksing/iter)
* [skiplist](https://github.com/MauriceGit/skiplist) **star:117** 高性能的 Go 跳表实现。   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [deque](https://github.com/gammazero/deque) **star:107** 快速环缓冲区deque(双端队列)。   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [ring](https://github.com/TheTannerRyan/ring) **star:105** 高性能、线程安全的bloom过滤器。   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:101** 区域四叉树具有高效的点定位和邻域查找功能。   [![最近一年没有更新][Yellow]](https://github.com/aurelien-rainone/go-rquad)   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [encoding](https://github.com/zhenjl/encoding) **star:100** 整形压缩库。   [![最近一年没有更新][Yellow]](https://github.com/zhenjl/encoding)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:87** 一个小型、强大和灵活的合并库。   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [bit](https://github.com/yourbasic/bit) **star:84** Go 语言集合数据结构。提供了额外的位操作功能。   [![最近一年没有更新][Yellow]](https://github.com/yourbasic/bit)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [gostl](https://github.com/liyue201/gostl) **star:77** 用于go的数据结构和算法库，旨在提供类似c++ STL的函数。   [![godoc][GoDoc]](https://godoc.org/github.com/liyue201/gostl)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:72** 实现在Go中计算levenshtein距离。   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [skiplist](https://github.com/gansidui/skiplist) **star:69** 在Go中实现了跳表。   [![最近一年没有更新][Yellow]](https://github.com/gansidui/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:64** 并行FIFO队列。   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [bloom](https://github.com/yourbasic/bloom) **star:49** Golang Bloom过滤器的实现。   [![最近一年没有更新][Yellow]](https://github.com/yourbasic/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bloom)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:49** 基于内存的实现了高性能的key:value存储库。指针缓存。   [![godoc][GoDoc]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:46** Go实现Count-Min-log sketch的功能 : 使用近似计数器进行近似计数(类似Count-Min sketch，但使用更少内存)。   [![最近一年没有更新][Yellow]](https://github.com/seiflotfy/count-min-log)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/count-min-log)
* [levenshtein](https://github.com/agext/levenshtein) **star:40** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![最近一周有更新][Green]](https://github.com/agext/levenshtein)   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [remember-go](https://github.com/rocketlaunchr/remember-go) **star:38** 用于缓存慢速数据库查询的通用接口(支持redis、memcached、ristretto或in-memory)。   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/remember-go)
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:27** 具备高并发能力，可替换 bufio.Writer。   [![最近一年没有更新][Yellow]](https://github.com/free/concurrent-writer)   [![godoc][GoDoc]](https://godoc.org/github.com/free/concurrent-writer)
* [crunch](https://github.com/superwhiskers/crunch) **star:23** 打包实现缓冲区，以便轻松处理各种数据类型。   [![最近一周有更新][Green]](https://github.com/superwhiskers/crunch)   [![godoc][GoDoc]](https://godoc.org/github.com/superwhiskers/crunch)
* [pipeline](https://github.com/hyfather/pipeline) **star:22** 实现了 fan-in 和 fan-out 的管道功能。   [![最近一年没有更新][Yellow]](https://github.com/hyfather/pipeline)   [![godoc][GoDoc]](https://godoc.org/github.com/hyfather/pipeline)
* [goset](https://github.com/zoumo/goset) **star:20** 一个有用的Go集合实现。   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [typ](https://github.com/gurukami/typ) **star:16** 从复杂结构中获取值，支持空类型、安全的类型转换。   [![godoc][GoDoc]](https://godoc.org/github.com/gurukami/typ)
* [deque](https://github.com/edwingeng/deque) **star:15** 高度优化的双端队列。   [![最近一周有更新][Green]](https://github.com/edwingeng/deque)   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/deque)
* [hide](https://github.com/emvi/hide) **star:14** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![最近一年没有更新][Yellow]](https://github.com/emvi/hide)   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [dict](https://github.com/srfrog/dict) **star:11** 实现类似 python dict的功能(dict)。   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [go-ef](https://github.com/amallia/go-ef) **star:11** 实现了 Elias-Fano 编码。   [![最近一年没有更新][Yellow]](https://github.com/amallia/go-ef)   [![godoc][GoDoc]](https://godoc.org/github.com/amallia/go-ef)
* [null](https://github.com/emvi/null) **star:9** 对空的 Go 数据类型也可以进行JSON进行解析/反解析。   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [mspm](https://github.com/BlackRabbitt/mspm) **star:9** 用于信息检索的多字符串模式匹配算法。   [![最近一年没有更新][Yellow]](https://github.com/BlackRabbitt/mspm)   [![godoc][GoDoc]](https://godoc.org/github.com/BlackRabbitt/mspm)
* [set](https://github.com/StudioSol/set) **star:7** 使用LinkedHashMap在Go中实现简单的set数据结构。   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [ptrie](https://github.com/viant/ptrie) **star:7** 前缀树的一种实现。   [![godoc][GoDoc]](https://godoc.org/github.com/viant/ptrie)
* [treap](https://github.com/perdata/treap) **star:7** 使用树堆进行持久、快速有序的映射。   [![godoc][GoDoc]](https://godoc.org/github.com/perdata/treap)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:6** 实现了有生命周期的键值对Map。   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)
* [gofal](https://github.com/xxjwxc/gofal) **star:6** 基于 Go 实现的分数计算。包含分子、分母运算    [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gofal)   [![包含中文文档][CN]](https://github.com/xxjwxc/gofal)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** 用于解析类似json的日志的工具，用于收集惟一的字段和事件。   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)

## 数据库

*数据库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [prometheus](https://github.com/prometheus/prometheus) **star:29652** 用于监控系统和时序的数据库。   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   [![最近一周有更新][Green]](https://github.com/prometheus/prometheus)   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [tidb](https://github.com/pingcap/tidb) **star:22829** TiDB是一个分布式SQL数据库。灵感来自谷歌F1的设计。   [![star > 2000][Awesome]](https://github.com/pingcap/tidb)   [![最近一周有更新][Green]](https://github.com/pingcap/tidb)   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/tidb)   [![包含中文文档][CN]](https://github.com/pingcap/tidb)
* [influxdb](https://github.com/influxdb/influxdb) **star:18486** 可伸缩的数据存储，用于指标衡量、事件和实时分析。   [![star > 2000][Awesome]](https://github.com/influxdb/influxdb)   [![最近一周有更新][Green]](https://github.com/influxdb/influxdb)   [![godoc][GoDoc]](https://godoc.org/github.com/influxdb/influxdb)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:17927** 可伸缩、区域备份、事务性数据存储。   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   [![最近一周有更新][Green]](https://github.com/cockroachdb/cockroach)   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [dgraph](https://github.com/dgraph-io/dgraph) **star:12682** 可伸缩、分布式、低延迟、高吞吐量的图形数据库。   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   [![最近一周有更新][Green]](https://github.com/dgraph-io/dgraph)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [groupcache](https://github.com/golang/groupcache) **star:8468** Groupcache是一个缓存和缓存填充库，在许多情况下，它是memcached的替代品。   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![最近一周有更新][Green]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:7399** 快速 K/V 存储。   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   [![最近一周有更新][Green]](https://github.com/dgraph-io/badger)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:5726** 基于SQLite的轻量级分布式关系数据库。   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [BigCache](https://github.com/allegro/bigcache) **star:3661** 高效的键/值缓存为千兆字节的数据。   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![最近一周有更新][Green]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3605** 在Go中实现[LevelDB](https://github.com/google/leveldb) key/value数据库。   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [go-cache](https://github.com/pmylund/go-cache) **star:3591** 基于内存的 K/V 存储/缓存 : (类似于Memcached)，适用于单机应用程序。   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![最近一周有更新][Green]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3242** Ledisdb是一种高性能的NoSQL，类似于基于LevelDB的Redis。   [![star > 2000][Awesome]](https://github.com/siddontang/ledisdb)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/ledisdb)
* [bbolt](https://github.com/etcd-io/bbolt) **star:3047** 一个用于Go的嵌入式键/值数据库。   [![star > 2000][Awesome]](https://github.com/etcd-io/bbolt)   [![最近一周有更新][Green]](https://github.com/etcd-io/bbolt)   [![godoc][GoDoc]](https://godoc.org/github.com/etcd-io/bbolt)
* [buntdb](https://github.com/tidwall/buntdb) **star:2661** 基于内存的K/V，快速，可嵌入的数据库，可自定义索引和空间支持。   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2464** 属于你的NoSQL数据库。   [![star > 2000][Awesome]](https://github.com/HouzuoGuo/tiedot)   [![godoc][GoDoc]](https://godoc.org/github.com/HouzuoGuo/tiedot)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1942** 开源，快速，可伸缩的时间序列数据库。支持PromQL。   [![最近一周有更新][Green]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [cache2go](https://github.com/muesli/cache2go) **star:1228** 基于内存的 K/V 缓存，支持超时的自动失效。   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:1110** Nutsdb是一个用纯Go编写的简单、快速、可嵌入、持久的键/值存储。它支持完全序列化的事务和许多数据结构，如列表、集合、排序集。   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   [![包含中文文档][CN]](https://github.com/xujiajun/nutsdb)
* [GCache](https://github.com/bluele/gcache) **star:1082** 支持过期缓存、LFU、LRU和ARC的缓存库。   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:1003** 区块链领域的一个SQL数据库。   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:860** 支持磁盘备份的可持久化 K/V 存储。   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [moss](https://github.com/couchbase/moss) **star:755** Moss是一个用100% Go编写的简单LSM键值存储引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/moss)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:723** 基于内存的快速线程安全的缓存，可缓存大量的条目。最大限度地减少GC开销。   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
* [eliasdb](https://github.com/krotik/eliasdb) **star:561** 无其他依赖项，支持REST API，短语搜索和sql类似的查询语言的事务图数据库。   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [levigo](https://github.com/jmhodges/levigo) **star:378** 实现了对LevelDB封装。   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [Bitcask](https://github.com/prologic/bitcask) **star:303** Bitcask是一个可嵌入的、持久的、快速的键值(KV)数据库，使用纯Go编写，具有可预测的读写性能、低延迟和高吞吐量，这得益于Bitcask的磁盘布局(LSM+WAL)。   [![最近一周有更新][Green]](https://github.com/prologic/bitcask)   [![godoc][GoDoc]](https://godoc.org/github.com/prologic/bitcask)
* [pudge](https://github.com/recoilme/pudge) **star:249** 使用Go的标准库编写的快速和简单的键/值存储。   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [piladb](https://github.com/fern4lvarez/piladb) **star:173** 基于堆栈数据结构的轻量级RESTful数据库引擎。   [![最近一年没有更新][Yellow]](https://github.com/fern4lvarez/piladb)   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Vasto](https://github.com/chrislusf/vasto) **star:171** 分布式高性能键值存储。可做磁盘备份。最终一致。高可用。能够在不中断服务的情况下增长或收缩。   [![最近一年没有更新][Yellow]](https://github.com/chrislusf/vasto)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [Kivik](https://github.com/go-kivik/kivik) **star:155** Kivik为CouchDB、PouchDB和类似的数据库提供了一个通用的Go和GopherJS客户端库。   [![最近一周有更新][Green]](https://github.com/go-kivik/kivik)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kivik/kivik)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:94** 具有持久性的键值存储。   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:89** 小型平面文件JSON存储。   [![最近一年没有更新][Yellow]](https://github.com/nanobox-io/golang-scribble)   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:47** 由 Couchbase服务 支持的RESTful缓存微服务。   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [bcache](https://github.com/iwanbk/bcache) **star:44** 基于内存的最终一致的分布式缓存。   [![godoc][GoDoc]](https://godoc.org/github.com/iwanbk/bcache)
* [cache](https://github.com/akyoto/cache) **star:35** 基于内存的 K/V 存储:带生命周期的值存储，0个依赖项，<100 LoC, 100%覆盖率。   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)
* [Databunker](https://github.com/paranoidguy/databunker) **star:34** 个人身份信息(PII)存储服务符合GDPR和CCPA。   [![最近一周有更新][Green]](https://github.com/paranoidguy/databunker)   [![godoc][GoDoc]](https://godoc.org/github.com/paranoidguy/databunker)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:33** BigCache 支持集群和独立且生命周期存储项。   [![最近一年没有更新][Yellow]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][GoDoc]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
* [Coffer](https://github.com/claygod/coffer) **star:20** 支持事务的简单ACID键值数据库。   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/coffer)
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:14** 用于临时数据存放的 K/V 存储。   [![最近一年没有更新][Yellow]](https://github.com/rafaeljesus/tempdb)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/tempdb)
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:12** 用 Go 对[RocksDB](https://rocksdb.org)实现了封装。   [![最近一年没有更新][Yellow]](https://github.com/kapitan-k/gorocksdb)   [![godoc][GoDoc]](https://godoc.org/github.com/kapitan-k/gorocksdb)
* [tracedb](https://github.com/unit-io/tracedb) **star:9** 快速timeseries数据库物联网，实时消息传递应用程序。使用pubsub通过tcp或websocket访问tracedb，使用github.com/unit-io/trace应用程序。   [![最近一周有更新][Green]](https://github.com/unit-io/tracedb)   [![godoc][GoDoc]](https://godoc.org/github.com/unit-io/tracedb)
* [gostore](https://github.com/twharmon/gostore) **star:4** Gostore是一个简单、持久的嵌入式键值存储引擎，用Go编写。   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/gostore)

*数据库迁移。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [migrate](https://github.com/golang-migrate/migrate) **star:3843** 基于CLI的数据库迁移库。   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![最近一周有更新][Green]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1645** 数据库迁移工具。允许使用go-bindata将迁移嵌入到应用程序中。   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [skeema](https://github.com/skeema/skeema) **star:638** 用于MySQL的纯sql模式管理系统，支持分片和外部在线模式更改工具。   [![godoc][GoDoc]](https://godoc.org/github.com/skeema/skeema)
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:419** 面向Gorm ORM的数据库 schema 迁移辅助程序。   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [goose](https://github.com/steinbacher/goose) **star:136** 数据库迁移工具。您可以通过创建增量SQL或Go脚本来管理数据库的升级。   [![最近一年没有更新][Yellow]](https://github.com/steinbacher/goose)   [![godoc][GoDoc]](https://godoc.org/github.com/steinbacher/goose)
* [darwin](https://github.com/GuiaBolso/darwin) **star:99** 用于数据库 schema 升级的库。   [![godoc][GoDoc]](https://godoc.org/github.com/GuiaBolso/darwin)
* [migrator](https://github.com/lopezator/migrator) **star:90** 非常简单的 Go 数据库迁移库。   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:45** 用Go -pg/pg编写的迁移包。   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [gondolier](https://github.com/emvi/gondolier) **star:28** 使用结构修饰的数据库迁移库。   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/gondolier)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** 简易的迁移工具-目前只支持MySQL，但计划很快支持Postgres, SQLite, MongoDB等。   [![最近一年没有更新][Yellow]](https://github.com/pravasan/pravasan)
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:23** 类似 Django fixture，用于 Go 建立内置数据库/sql库。   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-fixtures)
* [avro](https://github.com/khezen/avro) **star:12** 发现SQL schemas并将其转换为AVRO schemas。   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)
* [schema](https://github.com/adlio/schema) **star:4** 库，用于将数据库/sql兼容数据库的模式迁移嵌入到Go二进制文件中。   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/schema)

*数据库工具。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [vitess](https://github.com/youtube/vitess) **star:9675** vitess提供了可以为大规模web服务扩展MySQL数据库提供便利的服务和工具。   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   [![最近一周有更新][Green]](https://github.com/youtube/vitess)   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [pgweb](https://github.com/sosedoff/pgweb) **star:6315** 基于web的PostgreSQL数据库浏览器。   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![最近一周有更新][Green]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [kingshard](https://github.com/flike/kingshard) **star:5045** kingshard 是基于 Golang 的MySQL高性能代理。   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   [![包含中文文档][CN]](https://github.com/flike/kingshard)
* [orchestrator](https://github.com/github/orchestrator) **star:3422** MySQL复制拓扑管理器和可视化工具。   [![star > 2000][Awesome]](https://github.com/github/orchestrator)   [![最近一周有更新][Green]](https://github.com/github/orchestrator)   [![godoc][GoDoc]](https://godoc.org/github.com/github/orchestrator)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2821** 自动将MySQL数据同步到Elasticsearch中。   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:2264**  Go 工具集，用于处理MySQL协议和复制。   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql)   [![最近一周有更新][Green]](https://github.com/siddontang/go-mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [pREST](https://github.com/nuveo/prest) **star:2220** 基于PostgreSQL database的RESTful API服务。   [![star > 2000][Awesome]](https://github.com/nuveo/prest)   [![godoc][GoDoc]](https://godoc.org/github.com/nuveo/prest)
* [chproxy](https://github.com/Vertamedia/chproxy) **star:383** ClickHouse数据库的HTTP代理。   [![最近一周有更新][Green]](https://github.com/Vertamedia/chproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/Vertamedia/chproxy)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:182** 收集小的 insterts 并向 ClickHouse 服务器发送大请求。   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [myreplication](https://github.com/2tvenom/myreplication) **star:158** MySql二进制日志复制监听器。支持基于语句和行的复制。   [![最近一年没有更新][Yellow]](https://github.com/2tvenom/myreplication)   [![godoc][GoDoc]](https://godoc.org/github.com/2tvenom/myreplication)
* [octillery](https://github.com/knocknote/octillery) **star:75** 用于数据库分表(支持每个ORM或原生SQL)。   [![最近一周有更新][Green]](https://github.com/knocknote/octillery)   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [dbbench](https://github.com/sj14/dbbench) **star:35** 数据库基准测试工具，支持多个数据库和脚本。   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [prep](https://github.com/hexdigest/prep) **star:25** 在不更改代码的情况下使用准备好的SQL语句。   [![最近一年没有更新][Yellow]](https://github.com/hexdigest/prep)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/prep)
* [datagen](https://github.com/codingconcepts/datagen) **star:16** 一个快速的数据生成器，支持多表感知和多行DML。   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/datagen)
* [rwdb](https://github.com/andizzle/rwdb) **star:11** rwdb为多个数据库服务器的设置提供读取副本功能。   [![最近一年没有更新][Yellow]](https://github.com/andizzle/rwdb)   [![godoc][GoDoc]](https://godoc.org/github.com/andizzle/rwdb)
* [bucket](https://github.com/PumpkinSeed/bucket) **star:6** 优化的数据结构框架的Couchbase专门针对一个桶的使用。   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/bucket)

*SQL查询生成器，用于构建和使用SQL的库。 (翻译出错了? 试试 [英文版](README_EN.md#database) 吧~)*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2797** 帮助您构建SQL查询的Go库。   [![star > 2000][Awesome]](https://github.com/Masterminds/squirrel)   [![最近一周有更新][Green]](https://github.com/Masterminds/squirrel)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/squirrel)
* [xo](https://github.com/knq/xo) **star:2394** 基于现有的schema定义和自定义查询生成 Go 代码，基于支持PostgreSQL、MySQL、SQLite、Oracle和Microsoft SQL Server。   [![star > 2000][Awesome]](https://github.com/knq/xo)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/xo)
* [gendry](https://github.com/didi/gendry) **star:953** 非入侵的SQL构建器和强大的数据绑定器。   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   [![包含中文文档][CN]](https://github.com/didi/gendry)
* [goqu](https://github.com/doug-martin/goqu) **star:739** 常用的SQL生成器和查询库。   [![最近一周有更新][Green]](https://github.com/doug-martin/goqu)   [![godoc][GoDoc]](https://godoc.org/github.com/doug-martin/goqu)
* [Dotsql](https://github.com/gchaincl/dotsql) **star:507** Go library帮助您将sql文件保存在一个地方，并轻松地使用它们。   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/dotsql)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:462** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
* [jet](https://github.com/go-jet/jet) **star:230** 用于在Go中编写类型安全的SQL查询的框架，能够轻松地将数据库查询结果转换为所需的任意对象结构。   [![godoc][GoDoc]](https://godoc.org/github.com/go-jet/jet)
* [sqrl](https://github.com/elgris/sqrl) **star:202** SQL查询生成器，从Squirrel fork而来，并再此基础上对性能做了优化。   [![godoc][GoDoc]](https://godoc.org/github.com/elgris/sqrl)
* [Squalus](https://gitlab.com/qosenergy/squalus)  Go SQL中间层，能使得执行查询更加容易。
* [dbq](https://github.com/rocketlaunchr/dbq) **star:98** Zero boilerplate database operations for Go   [![最近一周有更新][Green]](https://github.com/rocketlaunchr/dbq)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)
* [igor](https://github.com/galeone/igor) **star:81** PostgreSQL的抽象层，支持高级功能和类似gorm的语法。   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [godbal](https://github.com/xujiajun/godbal) **star:50** 数据库抽象层(dbal)。支持SQL builder，轻松获取结果。   [![最近一年没有更新][Yellow]](https://github.com/xujiajun/godbal)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [sqlf](https://github.com/leporo/sqlf) **star:7** 快速SQL查询生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/leporo/sqlf)
* [qry](https://github.com/HnH/qry) **star:6** 该工具使用原始SQL查询从文件生成常量。   [![godoc][GoDoc]](https://godoc.org/github.com/HnH/qry)
* [mpath](https://github.com/spacetab-io/mpath-go) **star:5** MPTT(修改前序树遍历)包的SQL记录-物化路径实现。   [![godoc][GoDoc]](https://godoc.org/github.com/spacetab-io/mpath-go)
* [ormlite](https://github.com/pupizoid/ormlite)  包含一些类似orm的特性和sqlite数据库的辅助程序的轻量级包

## 数据库驱动程序

*用于连接和操作数据库的库。 (翻译出错了? 试试 [英文版](README_EN.md#database-drivers) 吧~)*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:9135** MySQL驱动程序。   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   [![最近一周有更新][Green]](https://github.com/go-sql-driver/mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [pq](https://github.com/lib/pq) **star:5722** 纯 Go 的Postgres驱动。   [![star > 2000][Awesome]](https://github.com/lib/pq)   [![最近一周有更新][Green]](https://github.com/lib/pq)   [![godoc][GoDoc]](https://godoc.org/github.com/lib/pq)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3885** SQLite3驱动程序。   [![star > 2000][Awesome]](https://github.com/mattn/go-sqlite3)
    * [pgx](https://github.com/jackc/pgx) **star:2434** PostgreSQL驱动，支持比现有database/sql更多的特性。   [![star > 2000][Awesome]](https://github.com/jackc/pgx)   [![最近一周有更新][Green]](https://github.com/jackc/pgx)   [![godoc][GoDoc]](https://godoc.org/github.com/jackc/pgx)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1154** 微软MSSQL驱动程序。   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8) **star:448** Oracle 驱动程序。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-oci8)
    * [goracle](https://github.com/go-goracle/goracle) **star:280** 基于 ODPI-C 的 Oracle 驱动程序
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:119** Firebird RDBMS SQL驱动程序。   [![godoc][GoDoc]](https://godoc.org/github.com/nakagami/firebirdsql)
    * [go-adodb](https://github.com/mattn/go-adodb) **star:102** Microsoft ActiveX对象数据库驱动程序。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [gofreetds](https://github.com/minus5/gofreetds) **star:98** 基于[FreeTDS](http://www.freetds.org)封装的微软MSSQL Go 驱动。   [![最近一年没有更新][Yellow]](https://github.com/minus5/gofreetds)   [![godoc][GoDoc]](https://godoc.org/github.com/minus5/gofreetds)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:47** Apache Avatica/Phoenix SQL驱动程序。   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [bgc](https://github.com/viant/bgc) **star:13** BigQuery 的数据存储连接。   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:8252** 基于 Go 的 Redis 客户端。   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   [![最近一周有更新][Green]](https://github.com/go-redis/redis)   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo) **star:7077** Redigo 是基于 Go 的Redis 客户端。   [![star > 2000][Awesome]](https://github.com/gomodule/redigo)   [![godoc][GoDoc]](https://godoc.org/github.com/gomodule/redigo)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:4138** 官方的 MongoDB 驱动。   [![star > 2000][Awesome]](https://github.com/mongodb/mongo-go-driver)   [![最近一周有更新][Green]](https://github.com/mongodb/mongo-go-driver)   [![godoc][GoDoc]](https://godoc.org/github.com/mongodb/mongo-go-driver)
    * [mgo](https://github.com/globalsign/mgo) **star:1762** (已停止维护) MongoDB驱动。   [![godoc][GoDoc]](https://godoc.org/github.com/globalsign/mgo)
    * [gorethink](https://github.com/dancannon/gorethink) **star:1504** RethinkDB 驱动。   [![最近一周有更新][Green]](https://github.com/dancannon/gorethink)   [![godoc][GoDoc]](https://godoc.org/github.com/dancannon/gorethink)
    * [redeo](https://github.com/bsm/redeo) **star:369** 与 redis 协议兼容的 TCP 服务器/服务。   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redeo)
    * [neoism](https://github.com/jmcvetta/neoism) **star:363** Golang 的 Neo4j 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:316** Aerospike 客户端。   [![最近一周有更新][Green]](https://github.com/aerospike/aerospike-client-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [gocb](https://github.com/couchbase/gocb) **star:307** 官方Couchbase Go SDK。   [![最近一周有更新][Green]](https://github.com/couchbase/gocb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/gocb)
    * [gocql](http://gocql.github.io)  Apache Cassandra 的 Go 驱动。
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:298** Couchbase客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:117** 实现了基于 Redigo 客户端的redislabs' ReJSON 模块。可简单地将结构体存储为JSON对象并对其进行操作。   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:74** Neo4j REST 客户端。   [![最近一年没有更新][Yellow]](https://github.com/davemeehan/Neo4j-GO)   [![godoc][GoDoc]](https://godoc.org/github.com/davemeehan/Neo4j-GO)
    * [godis](https://github.com/piaohao/godis) **star:72** 由GoLang实现的redis客户端，灵感来自jedis。   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [arangolite](https://github.com/solher/arangolite) **star:66** 轻量级的 ArangoDB 驱动。   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [dynago](https://github.com/underarmour/dynago) **star:66** 满足 principle of least surprise 的 DynamoDB 客户端。   [![最近一年没有更新][Yellow]](https://github.com/underarmour/dynago)   [![godoc][GoDoc]](https://godoc.org/github.com/underarmour/dynago)
    * [mgm](https://github.com/kamva/mgm) **star:62** 基于MongoDB模型的ODM for Go(基于官方MongoDB驱动)。   [![godoc][GoDoc]](https://godoc.org/github.com/kamva/mgm)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:34**  Pilosa客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [forestdb](https://github.com/couchbase/goforestdb) **star:28** 基于 Go 的 ForestDB 接口实现。   [![最近一年没有更新][Yellow]](https://github.com/couchbase/goforestdb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/goforestdb)
    * [neo4j](https://github.com/cihangir/neo4j) **star:25** Neo4j Rest API实现。   [![最近一年没有更新][Yellow]](https://github.com/cihangir/neo4j)   [![godoc][GoDoc]](https://godoc.org/github.com/cihangir/neo4j)
    * [goriak](https://github.com/zegl/goriak) **star:24**  Riak KV 驱动。   [![最近一年没有更新][Yellow]](https://github.com/zegl/goriak)   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [xredis](https://github.com/shomali11/xredis) **star:10** 类型安全，可定制，清晰和易用的Redis客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)
    * [godscache](https://github.com/defcronyke/godscache) **star:7** 谷歌云平台Go Datastore包的封装，它采用了memcached添加缓存。   [![最近一年没有更新][Yellow]](https://github.com/defcronyke/godscache)   [![godoc][GoDoc]](https://godoc.org/github.com/defcronyke/godscache)
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache客户端库。
    * [asc](https://github.com/viant/asc) **star:4** Aerospike 的数据存储连接器。   [![godoc][GoDoc]](https://godoc.org/github.com/viant/asc)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:6413** 基于 Go 的现代文本索引库。   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![最近一周有更新][Green]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [riot](https://github.com/go-ego/riot) **star:5091** 基于 Go 的 开源、分布式、简单高效的搜索引擎。   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   [![包含中文文档][CN]](https://github.com/go-ego/riot)
    * [elastic](https://github.com/olivere/elastic) **star:4802** Elasticsearch 客户端。   [![star > 2000][Awesome]](https://github.com/olivere/elastic)   [![godoc][GoDoc]](https://godoc.org/github.com/olivere/elastic)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:2250** 官方 Elasticsearch 客户端。   [![star > 2000][Awesome]](https://github.com/elastic/go-elasticsearch)   [![最近一周有更新][Green]](https://github.com/elastic/go-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/elastic/go-elasticsearch)
    * [elastigo](https://github.com/mattbaird/elastigo) **star:952** Elasticsearch 客户端。   [![最近一年没有更新][Yellow]](https://github.com/mattbaird/elastigo)   [![godoc][GoDoc]](https://godoc.org/github.com/mattbaird/elastigo)
    * [elasticsql](https://github.com/cch123/elasticsql) **star:495** 将 SQL 转换为 elasticsearch dsl。   [![最近一年没有更新][Yellow]](https://github.com/cch123/elasticsql)   [![godoc][GoDoc]](https://godoc.org/github.com/cch123/elasticsql)
    * [skizze](https://github.com/seiflotfy/skizze) **star:72** 面向概率数据结构的服务和存储。   [![最近一年没有更新][Yellow]](https://github.com/seiflotfy/skizze)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)
    * [goes](https://github.com/OwnLocal/goes) **star:24** 实现了与 Elasticsearch 交互的库。   [![最近一年没有更新][Yellow]](https://github.com/OwnLocal/goes)   [![godoc][GoDoc]](https://godoc.org/github.com/OwnLocal/goes)

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:13264** 图形数据库，支持多个后端。   [![star > 2000][Awesome]](https://github.com/google/cayley)   [![最近一周有更新][Green]](https://github.com/google/cayley)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cayley)
    * [gokv](https://github.com/philippgille/gokv) **star:180** 可扩展的简单的 K/V 存储(Redis、Consul、etcd、bbolt、BadgerDB、LevelDB、Memcached、DynamoDB、S3、PostgreSQL、MongoDB、CockroachDB等等)。   [![最近一周有更新][Green]](https://github.com/philippgille/gokv)   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [cachego](https://github.com/fabiorphp/cachego) **star:122** 基于多个驱动程序的缓存组件。   [![最近一周有更新][Green]](https://github.com/fabiorphp/cachego)   [![godoc][GoDoc]](https://godoc.org/github.com/fabiorphp/cachego)
    * [dsc](https://github.com/viant/dsc) **star:18** 面向 SQL、NoSQL、结构化文件的数据存储连接。   [![最近一周有更新][Green]](https://github.com/viant/dsc)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## 日期和时间

*用于处理日期和时间的库。 (翻译出错了? 试试 [英文版](README_EN.md#date-and-time) 吧~)*

* [now](https://github.com/jinzhu/now) **star:2440** now 是时间有关的工具类。   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [dateparse](https://github.com/araddon/dateparse) **star:993** 可以解析很多格式不固定的日期字符串。   [![godoc][GoDoc]](https://godoc.org/github.com/araddon/dateparse)
* [carbon](https://github.com/uniplaces/carbon) **star:438** 简单的时间扩展，包含了许多使用方法，从 PHP Carbon 库移植的。   [![最近一年没有更新][Yellow]](https://github.com/uniplaces/carbon)   [![godoc][GoDoc]](https://godoc.org/github.com/uniplaces/carbon)
* [durafmt](https://github.com/hako/durafmt) **star:287** 轻量级、可让time.Duration更加易读的库。   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [timeutil](https://github.com/leekchan/timeutil) **star:176** 面向 Golang 的时间库，集成了很多有用的扩展(Timedelta, Strftime, ...)。   [![最近一年没有更新][Yellow]](https://github.com/leekchan/timeutil)   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/timeutil)
* [iso8601](https://github.com/relvacode/iso8601) **star:71** 不用正则表达式有效解析 ISO8601 日期时间。   [![最近一年没有更新][Yellow]](https://github.com/relvacode/iso8601)   [![godoc][GoDoc]](https://godoc.org/github.com/relvacode/iso8601)
* [timespan](https://github.com/SaidinWoT/timespan) **star:68** 用于处理时间间隔相关的库，可定义开始时间和持续时间。   [![最近一年没有更新][Yellow]](https://github.com/SaidinWoT/timespan)   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:68** 太阳历实现。   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/go-persian-calendar)
* [date](https://github.com/rickb777/date) **star:35** 增加处理日期、日期范围、时间跨度、时间段和time-of-day。   [![godoc][GoDoc]](https://godoc.org/github.com/rickb777/date)
* [feiertage](https://github.com/wlbr/feiertage) **star:25** 用于计算德国公共假期的函数，比如复活节、感恩节等   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:21** 计算指定位置的日出和日落时间。   [![最近一年没有更新][Yellow]](https://github.com/nathan-osman/go-sunrise)   [![godoc][GoDoc]](https://godoc.org/github.com/nathan-osman/go-sunrise)
* [kair](https://github.com/GuilhermeCaruso/kair) **star:15** 用于处理日期和时间的 golang 库。   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)
* [cronrange](https://github.com/1set/cronrange) **star:13** 解析cron风格的时间范围表达式，检查给定的时间是否在任何范围内。   [![godoc][GoDoc]](https://godoc.org/github.com/1set/cronrange)
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:10** 可以允许 time.Time 为null。   [![最近一年没有更新][Yellow]](https://github.com/kirillDanshin/nulltime)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/nulltime)
* [tuesday](https://github.com/osteele/tuesday) **star:8** Ruby-compatible Strftime function。   [![最近一年没有更新][Yellow]](https://github.com/osteele/tuesday)   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/tuesday)
* [strftime](https://github.com/awoodbeck/strftime) **star:4** C99-compatible strftime formatter。   [![最近一年没有更新][Yellow]](https://github.com/awoodbeck/strftime)   [![godoc][GoDoc]](https://godoc.org/github.com/awoodbeck/strftime)
* [go-str2duration](https://github.com/xhit/go-str2duration) **star:2** 将字符串转换为持续时间。支持的时间。持续时间返回字符串等。   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-str2duration)
* [go-week](https://github.com/stoewer/go-week) **star:2** 一个有效的软件包，以配合ISO8601周日期。   [![godoc][GoDoc]](https://godoc.org/github.com/stoewer/go-week)

## 分布式系统

*协助构建分布式系统的包。 (翻译出错了? 试试 [英文版](README_EN.md#distributed-systems) 吧~)*

* [go-kit](https://github.com/go-kit/kit) **star:16445** 支持服务发现、负载平衡、插件式传输、请求跟踪等功能的Microservice toolkit。   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [grpc-go](https://github.com/grpc/grpc-go) **star:10821** gRPC的Go语言实现。   [![star > 2000][Awesome]](https://github.com/grpc/grpc-go)   [![最近一周有更新][Green]](https://github.com/grpc/grpc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/grpc/grpc-go)
* [micro](https://github.com/micro/micro) **star:7737** 可插拔的微服务 toolkit 和分布式系统平台。   [![star > 2000][Awesome]](https://github.com/micro/micro)   [![最近一周有更新][Green]](https://github.com/micro/micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [NATS](https://github.com/nats-io/gnatsd) **star:7390** 轻量级、高性能消息传递系统，可用于微服务、物联网(IoT)和云。   [![star > 2000][Awesome]](https://github.com/nats-io/gnatsd)   [![最近一周有更新][Green]](https://github.com/nats-io/gnatsd)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/gnatsd)
* [rpcx](https://github.com/smallnest/rpcx) **star:4457** 分布式可插拔的RPC服务框架，如阿里巴巴Dubbo。   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [tendermint](https://github.com/tendermint/tendermint) **star:3540** 一个高性能中间件，可将任何语言的状态机转换为 Byzantine Fault 状态机。使用 Tendermint 一致性及区块链协议。   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   [![最近一周有更新][Green]](https://github.com/tendermint/tendermint)   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [torrent](https://github.com/anacrolix/torrent) **star:3376** BitTorrent 客户端。   [![star > 2000][Awesome]](https://github.com/anacrolix/torrent)   [![最近一周有更新][Green]](https://github.com/anacrolix/torrent)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/torrent)
* [raft](https://github.com/hashicorp/raft) **star:3298** Raft consensus协议的实现。 by HashiCorp。   [![star > 2000][Awesome]](https://github.com/hashicorp/raft)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/raft)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Raft consensus协议的实现。 by CoreOS。
* [dragonboat](https://github.com/lni/dragonboat) **star:2916** 一个功能齐全，高性能的库集。   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   [![最近一周有更新][Green]](https://github.com/lni/dragonboat)   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   [![包含中文文档][CN]](https://github.com/lni/dragonboat)
* [glow](https://github.com/chrislusf/glow) **star:2757** 全部用 Go 实现，易用、可伸缩，可用于分布式大数据处理，Map-Reduce, DAG执行。   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![最近一年没有更新][Yellow]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [KrakenD](https://github.com/devopsfaith/krakend) **star:2453** 具有中间件的高性能API网关框架。   [![star > 2000][Awesome]](https://github.com/devopsfaith/krakend)   [![godoc][GoDoc]](https://godoc.org/github.com/devopsfaith/krakend)
* [gleam](https://github.com/chrislusf/gleam) **star:2409** 使用纯Go和Luajit编写的快速、可伸缩的分布式map/reduce系统，结合了Go的高并发性和Luajit的高性能，可以独立运行或分布式运行。   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2278** 高性能、分布式、安全和低延迟的发布-订阅平台，使用MQTT、Websockets和love构建。   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [liftbridge](https://github.com/liftbridge-io/liftbridge) **star:1483** 用于nat的轻量级、容错的消息流。   [![最近一周有更新][Green]](https://github.com/liftbridge-io/liftbridge)   [![godoc][GoDoc]](https://godoc.org/github.com/liftbridge-io/liftbridge)
* [hprose](https://github.com/hprose/hprose-golang) **star:1083** 支持25+种语言RPC库。   [![最近一周有更新][Green]](https://github.com/hprose/hprose-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/hprose/hprose-golang)   [![包含中文文档][CN]](https://github.com/hprose/hprose-golang)
* [ringpop-go](https://github.com/uber/ringpop-go) **star:611** 可伸缩的，容错、应用分层的的Go应用程序。   [![最近一年没有更新][Yellow]](https://github.com/uber/ringpop-go)   [![godoc][GoDoc]](https://godoc.org/github.com/uber/ringpop-go)
* [gorpc](https://github.com/valyala/gorpc) **star:579** 简单、快速和可伸缩的RPC库。   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [go-health](https://github.com/InVisionApp/go-health) **star:534** 用于在服务中启用异步依赖项健康检查的库。   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [rain](https://github.com/cenkalti/rain) **star:482** BitTorrent客户端和库。   [![godoc][GoDoc]](https://godoc.org/github.com/cenkalti/rain)
* [digota](https://github.com/digota/digota) **star:325** 基于 grpc 的电子商务微服务。   [![最近一年没有更新][Yellow]](https://github.com/digota/digota)   [![godoc][GoDoc]](https://godoc.org/github.com/digota/digota)
* [dot](https://github.com/dotchain/dot/)  基于 transformation/OT 的分布式同步。
* [sleuth](https://github.com/ursiform/sleuth) **star:312** 用于HTTP服务之间进行无中心p2p自动发现和RPC通信的库(使用[ZeroMQ](https://github.com/zeromq/libzmq))。   [![最近一年没有更新][Yellow]](https://github.com/ursiform/sleuth)   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) **star:291** 为支持为golang服务定义异步服务健康检查而构建的库。   [![godoc][GoDoc]](https://godoc.org/github.com/AppsFlyer/go-sundheit)
* [go-jump](https://github.com/dgryski/go-jump) **star:281** 提供了谷歌的 “Jump” 一致哈希函数接口。   [![最近一年没有更新][Yellow]](https://github.com/dgryski/go-jump)   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [consistent](https://github.com/buraksezer/consistent) **star:266** Consistent hashing with bounded loads。   [![godoc][GoDoc]](https://godoc.org/github.com/buraksezer/consistent)
* [dht](https://github.com/anacrolix/dht) **star:150** BitTorrent Kademlia DHT的实现。   [![最近一周有更新][Green]](https://github.com/anacrolix/dht)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:125** jsonrpc 包，实现了 JSON-RPC 2.0。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:118** JSON-RPC 2.0 HTTP客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [resgate](https://resgate.io/)  用于构建REST、实时和RPC API的实时API网关，其中所有客户端都是无缝同步的。
* [redis-lock](https://github.com/bsm/redislock) **star:99** 基于redis的分布式锁简易实现。   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redislock)
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:60** 用于对 Celery worker、任务、事件进行交互和监控的库。   [![最近一年没有更新][Yellow]](https://github.com/svcavallar/celeriac.v1)   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:50** 实现了谷歌的jump consistent hash。   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed 分布式锁定实现。
* [drmaa](https://github.com/dgruber/drmaa) **star:29** 符合 DRMAA 标准的集群调度程序作业提交库。   [![最近一年没有更新][Yellow]](https://github.com/dgruber/drmaa)   [![godoc][GoDoc]](https://godoc.org/github.com/dgruber/drmaa)
* [pglock](https://cirello.io/pglock)  postgresql 的分布式锁定实现。
* [outboxer](https://github.com/italolelis/outboxer) **star:27** 实现了 outbox pattern Go 库。   [![最近一周有更新][Green]](https://github.com/italolelis/outboxer)   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)
* [flowgraph](https://github.com/vectaport/flowgraph) **star:26** flow-based programming package。   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [go-pdu](https://github.com/pdupub/go-pdu) **star:12** 一个去中心化的基于身份的社交网络。   [![最近一周有更新][Green]](https://github.com/pdupub/go-pdu)   [![godoc][GoDoc]](https://godoc.org/github.com/pdupub/go-pdu)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:10** 基于 DynamoDB 的 原子计数器的库。   [![最近一年没有更新][Yellow]](https://github.com/tylfin/dynatomic)   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)

## 动态域名

*更新动态DNS记录的工具。 (翻译出错了? 试试 [英文版](README_EN.md#dynamic-dns) 吧~)*

* [GoDNS](https://github.com/timothyye/godns) **star:555** 一个动态DNS客户端工具，支持DNSPod & HE.net。   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)
* [DDNS](https://github.com/skibish/ddns) **star:136** 个人 DDNS 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/skibish/ddns)
* [dyndns](https://gitlab.com/alcastle/dyndns)  后台去处理定期和自动检查您的IP地址，并作出更新(一个或多个)动态DNS记录，为谷歌域，每当您的地址变化。

## 电子邮件

*实现了电子邮件创建和发送。 (翻译出错了? 试试 [英文版](README_EN.md#email) 吧~)*

* [MailHog](https://github.com/mailhog/MailHog) **star:6084** 电子邮件和SMTP测试工具，对外提供了 web 和 API 接口。   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [hermes](https://github.com/matcornic/hermes) **star:1905** 可生成干净的、响应式的HTML电子邮件。   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1266** 一个强大和灵活的电子邮件库。   [![最近一周有更新][Green]](https://github.com/jordan-wright/email)   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:921** 用于客户端和服务器的IMAP库。   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:571** SendGrid 的 Go语言库，用于发送电子邮件。   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [chasquid](https://blitiri.com.ar/p/chasquid)  用Go编写的SMTP服务器。
* [mailgun-go](https://github.com/mailgun/mailgun-go) **star:443** 使用Mailgun API去库发送邮件。   [![godoc][GoDoc]](https://godoc.org/github.com/mailgun/mailgun-go)
* [Hectane](https://github.com/hectane/hectane) **star:181** 轻量级的SMTP客户机，提供了HTTP API。   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [douceur](https://github.com/aymerick/douceur) **star:174** 在HTML邮件中支持CSS内联。   [![最近一年没有更新][Yellow]](https://github.com/aymerick/douceur)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [go-message](https://github.com/emersion/go-message) **star:143** 用于Internet消息格式化和邮件消息的流媒体库。   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-message)
* [smtp](https://github.com/mailhog/smtp) **star:54** SMTP服务器协议状态机。   [![最近一年没有更新][Yellow]](https://github.com/mailhog/smtp)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/smtp)
* [go-dkim](https://github.com/toorop/go-dkim) **star:52** DKIM库，用于签署 & 验证电子邮件。   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:46** 在HTML邮件中支持CSS内联。   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [mailchain](https://github.com/mailchain/mailchain) **star:42** 发送加密的电子邮件到区块链地址写在Go。   [![最近一周有更新][Green]](https://github.com/mailchain/mailchain)   [![godoc][GoDoc]](https://godoc.org/github.com/mailchain/mailchain)
* [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:19** 非常简单的包发送电子邮件与SMTP保持活动和两个超时:连接和发送。   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-simple-mail)

## 可嵌入的脚本语言

*在go代码中嵌入其他语言。 (翻译出错了? 试试 [英文版](README_EN.md#embeddable-scripting-languages) 吧~)*

* [otto](https://github.com/robertkrimen/otto) **star:5169** 用 Go 编写的 JavaScript 解释器。   [![star > 2000][Awesome]](https://github.com/robertkrimen/otto)   [![godoc][GoDoc]](https://godoc.org/github.com/robertkrimen/otto)
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:3339** 用 Go 实现的 Lua 5.1 虚拟机和编译器。   [![star > 2000][Awesome]](https://github.com/yuin/gopher-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/yuin/gopher-lua)
* [go-lua](https://github.com/Shopify/go-lua) **star:1793** 用 Go 实现的 Lua 5.2 VM接口。   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [tengo](https://github.com/d5/tengo) **star:1572** 字节码编译的脚本语言。   [![最近一周有更新][Green]](https://github.com/d5/tengo)   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [expr](https://github.com/antonmedv/expr) **star:1152** 一个可以计算表达式的引擎。   [![最近一周有更新][Green]](https://github.com/antonmedv/expr)   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [go-python](https://github.com/sbinet/go-python) **star:1014** CPython C-API 的 Go 接口。   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [anko](https://github.com/mattn/anko) **star:983** 用Go编写的解释器。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-php](https://github.com/deuill/go-php) **star:729** PHP 的 Go 接口。   [![最近一年没有更新][Yellow]](https://github.com/deuill/go-php)   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:686** 支持 Duktape JavaScript 引擎。   [![最近一周有更新][Green]](https://github.com/olebedev/go-duktape)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [golua](https://github.com/aarzilli/golua) **star:469** Lua C 的 Go 接口。
* [gisp](https://github.com/jcla1/gisp) **star:435** LISP 的 Go 接口。   [![最近一年没有更新][Yellow]](https://github.com/jcla1/gisp)   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [cel-go](https://github.com/google/cel-go) **star:380** 快速，可移植，非图灵完整的表达评估与渐进分型。   [![godoc][GoDoc]](https://godoc.org/github.com/google/cel-go)
* [gval](https://github.com/PaesslerAG/gval) **star:184** 一种用Go编写的高度可定制的表达式语言。   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [gentee](https://github.com/gentee/gentee) **star:45** 嵌入式脚本编程语言。   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)
* [binder](https://github.com/alexeyco/binder) **star:35** Lua接口，基于[gopher-lua](https://github.com/yuin/gopher-lua)。   [![最近一年没有更新][Yellow]](https://github.com/alexeyco/binder)   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [purl](https://github.com/ian-kent/purl) **star:29** 嵌入 Go 的 Perl 5.18.2。   [![最近一年没有更新][Yellow]](https://github.com/ian-kent/purl)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/purl)
* [ngaro](https://github.com/db47h/ngaro) **star:19** 嵌入式 Ngaro VM实现，支持在 Retro 中运行脚本。   [![最近一年没有更新][Yellow]](https://github.com/db47h/ngaro)   [![godoc][GoDoc]](https://godoc.org/github.com/db47h/ngaro)

## 错误处理

*处理错误的库。 (翻译出错了? 试试 [英文版](README_EN.md#error-handling) 吧~)*

* [errors](https://github.com/pkg/errors) **star:5597** 可让你很简单的进行错误处理。   [![star > 2000][Awesome]](https://github.com/pkg/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/errors)
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:816** 可将一系列的错误作为一个整体来显示。   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [errorx](https://github.com/joomcode/errorx) **star:611** 一个功能丰富的错误包，可进行堆栈跟踪、组装异常信息以及其他。   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:595** 可展示错误的堆栈跟踪信息和源码片段。   [![最近一年没有更新][Yellow]](https://github.com/ztrue/tracerr)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [eris](https://github.com/rotisserie/eris) **star:580** 在Go中处理、跟踪和记录错误的更好方法。兼容标准错误库和github.com/pkg/errors。   [![最近一周有更新][Green]](https://github.com/rotisserie/eris)   [![godoc][GoDoc]](https://godoc.org/github.com/rotisserie/eris)
* [errlog](https://github.com/snwfdhmp/errlog) **star:303** 用于定位抛出错误的源代码(以及一些其他快速调试特性)。可插入到任何 logger 的位置。   [![godoc][GoDoc]](https://godoc.org/github.com/snwfdhmp/errlog)
* [emperror](https://github.com/emperror/emperror) **star:115** 用于Go库和应用程序的错误处理工具和最佳实践。   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/emperror)
* [errors](https://github.com/emperror/errors) **star:43** 替换标准库错误包和github.com/pkg/errors。提供各种错误处理原语。   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/errors)
* [werr](https://github.com/txgruppi/werr) **star:13** 对错误异常进行了捕获封装，封装信息包含了调用它的文件、行和堆栈。   [![最近一年没有更新][Yellow]](https://github.com/txgruppi/werr)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/werr)
* [errors](https://github.com/neuronlabs/errors) **star:3** 使用分类原语进行简单的golang错误处理。   [![godoc][GoDoc]](https://godoc.org/github.com/neuronlabs/errors)
* [errors](https://github.com/PumpkinSeed/errors) **star:2** 最简单的错误包装器，具有出色的性能和最小的内存开销。   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/errors)
* [Falcon](https://github.com/SonicRoshan/falcon) **star:2** 一个简单但功能强大的错误处理包。   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/falcon)

## 文件

*处理文件和文件系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#files) 吧~)*

* [afero](https://github.com/spf13/afero) **star:2616** 文件系统的抽象系统。   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/pdfcpu/pdfcpu) **star:1349** PDF处理器。   [![godoc][GoDoc]](https://godoc.org/github.com/pdfcpu/pdfcpu)
* [notify](https://github.com/rjeczalik/notify) **star:548** 文件系统事件通知库，具有类似于os/signal的简单API，。   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/notify)
* [copy](https://github.com/otiai10/copy) **star:143** 递归地复制目录。   [![最近一周有更新][Green]](https://github.com/otiai10/copy)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/copy)
* [bigfile](https://github.com/bigfile/bigfile) **star:117** 一个文件传输系统，支持管理文件与http api, rpc调用和ftp客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/bigfile/bigfile)   [![包含中文文档][CN]](https://github.com/bigfile/bigfile)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:64** 使用 tag 加载 csv 文件。   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-csv-tag)
* [opc](https://github.com/qmuntal/opc) **star:63** 加载Open Packaging Conventions (OPC)文件。   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [stl](https://gitlab.com/russoj88/stl)  采用并行读取算法的进行读取和写入STL(立体光刻)文件的模块。
* [skywalker](https://github.com/dixonwille/skywalker) **star:55** 可以轻松地并发地遍历文件系统。   [![最近一年没有更新][Yellow]](https://github.com/dixonwille/skywalker)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/skywalker)
* [vfs](https://github.com/C2FO/vfs) **star:47** 一组可插拔的、可扩展的和自定义的文件系统功能，用于跨越许多文件系统类型，如os、S3和GCS。   [![最近一周有更新][Green]](https://github.com/C2FO/vfs)   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [afs](https://github.com/viant/afs) **star:43** 用于Go的抽象文件存储(mem、scp、zip、tar、cloud: s3、gs)。   [![最近一周有更新][Green]](https://github.com/viant/afs)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/afs)
* [tarfs](https://github.com/posener/tarfs) **star:39** tar文件的实现[ FileSystem 接口](https://godoc.org/github.com/kr/fs#FileSystem)。   [![最近一周有更新][Green]](https://github.com/posener/tarfs)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/tarfs)
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:23** ExifTool 的 Go 实现，这个著名的库用于从文件(图片、PDF、office，…)中提取尽可能多的元数据(EXIF、IPTC，…)。   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:21** 加载gtfs文件。   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [gut/yos](https://github.com/1set/gut) **star:20** 简单和可靠的包文件操作，如复制/移动/diff/列表的文件，目录和符号链接。   [![godoc][GoDoc]](https://godoc.org/github.com/1set/gut)
* [checksum](https://github.com/codingsince1985/checksum) **star:17** 生成大型文件的消息摘要(如 MD5 和 SHA256)。   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [flop](https://github.com/homedepot/flop) **star:15** 文件操作库，是[GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invoc.html)的镜像。   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:14** 拷贝文件。   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
* [parquet](https://github.com/parsyl/parquet) **star:6** 读写[parquet](https://parquet.apache.org)文件。   [![godoc][GoDoc]](https://godoc.org/github.com/parsyl/parquet)

## 金融

*会计和财务软件包。 (翻译出错了? 试试 [英文版](README_EN.md#financial) 吧~)*

* [decimal](https://github.com/shopspring/decimal) **star:1982** 任意精度定点的十进制数。   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-money](https://github.com/rhymond/go-money) **star:719** Fowler 货币模式的实现。   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [accounting](https://github.com/leekchan/accounting) **star:543** 货币和货币格式。   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [go-finance](https://github.com/FlashBoys/go-finance) **star:539** 综合金融市场数据。   [![最近一年没有更新][Yellow]](https://github.com/FlashBoys/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/FlashBoys/go-finance)
* [techan](https://github.com/sdcoffey/techan) **star:238** 拥有先进的市场分析和交易策略的技术分析库。   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [orderbook](https://github.com/i25959341/orderbook) **star:122** 限购单匹配引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/i25959341/orderbook)
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:73** 查询 OFX 服务器和/或解析响应。   [![godoc][GoDoc]](https://godoc.org/github.com/aclindsa/ofxgo)
* [transaction](https://github.com/claygod/transaction) **star:65** 嵌入式事务数据库，可多线程模式运行。   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [vat](https://github.com/dannyvankooten/vat) **star:65** 增值税编号验证 & 欧盟增值税税率。   [![最近一年没有更新][Yellow]](https://github.com/dannyvankooten/vat)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/vat)
* [go-finance](https://github.com/alpeb/go-finance) **star:56** 用于货币时间价值(年金)、现金流、利率转换、债券和折旧计算的金融函数库。   [![godoc][GoDoc]](https://godoc.org/github.com/alpeb/go-finance)
* [go-finnhub](https://github.com/m1/go-finnhub) **star:21** 来自finnhu .io的股票市场、外汇和密码数据客户。从60多个证券交易所、10个外汇经纪人和15多个密码交易所获取实时金融市场数据。   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-finnhub)
* [currency](https://github.com/bnkamalesh/currency) **star:19** 高性能、准确的货币计算软件包。   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)
* [go-finance](https://github.com/pieterclaerhout/go-finance) **star:3** 模块获取汇率，通过VIES检查增值税号码，检查IBAN银行账号。   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-finance)

## 表单

*用于处理表单的库。 (翻译出错了? 试试 [英文版](README_EN.md#forms) 吧~)*

* [nosurf](https://github.com/justinas/nosurf) **star:1038** CSRF保护中间件。   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/nosurf)
* [binding](https://github.com/mholt/binding) **star:768** 将来自 net/HTTP 请求的表单、JSON 数据绑定到结构体。   [![最近一年没有更新][Yellow]](https://github.com/mholt/binding)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:508** 用于Go web应用程序和服务的CSRF保护。   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [form](https://github.com/go-playground/form) **star:398**  将 url 中的数据解析到 Go 变量中，以及将 Go 语言变量编码进 url。支持 Dual Array 及 Full map。   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [conform](https://github.com/leebenson/conform) **star:184** 控制用户输入。基于struct tags可对数据进行修剪、清理和擦除。   [![godoc][GoDoc]](https://godoc.org/github.com/leebenson/conform)
* [formam](https://github.com/monoculum/formam) **star:139** 将表单的值解码为 struct。   [![最近一周有更新][Green]](https://github.com/monoculum/formam)   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:104** 与框架无关的库，用于解析和验证支持多部分表单和文件的表单/JSON数据。   [![最近一年没有更新][Yellow]](https://github.com/albrow/forms)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)
* [bind](https://github.com/robfig/bind) **star:24** 将表单数据与任意 Go 变量进行绑定。   [![最近一年没有更新][Yellow]](https://github.com/robfig/bind)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/bind)
* [queryparam](https://github.com/tomwright/queryparam) **star:2** 解码的url。值转换为标准或自定义类型的可用结构值。   [![godoc][GoDoc]](https://godoc.org/github.com/tomwright/queryparam)

## 方法

*在Go中支持函数式编程的包。 (翻译出错了? 试试 [英文版](README_EN.md#functional) 吧~)*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1130** 常用辅助方法集合。   [![最近一年没有更新][Yellow]](https://github.com/tobyhede/go-underscore)   [![godoc][GoDoc]](https://godoc.org/github.com/tobyhede/go-underscore)
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:136** 提供函数式编程功能。   [![最近一年没有更新][Yellow]](https://github.com/TeaEntityLab/fpGo)   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:63** Functional Experiment in Go。   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)

## 游戏开发

*很棒的游戏开发库。 (翻译出错了? 试试 [英文版](README_EN.md#game-development) 吧~)*

* [Leaf](https://github.com/name5566/leaf) **star:3439** 轻量级游戏服务器框架。   [![star > 2000][Awesome]](https://github.com/name5566/leaf)   [![godoc][GoDoc]](https://godoc.org/github.com/name5566/leaf)   [![包含中文文档][CN]](https://github.com/name5566/leaf)
* [Pixel](https://github.com/faiface/pixel) **star:2793** 手工制作的 2D 游戏库。   [![star > 2000][Awesome]](https://github.com/faiface/pixel)   [![godoc][GoDoc]](https://godoc.org/github.com/faiface/pixel)
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:2569** 很简单的 2D 游戏库。   [![star > 2000][Awesome]](https://github.com/hajimehoshi/ebiten)   [![最近一周有更新][Green]](https://github.com/hajimehoshi/ebiten)   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/ebiten)
* [goworld](https://github.com/xiaonanln/goworld) **star:1425** 可伸缩的游戏服务器引擎，具有 space-entity 框架和 hot-swapping 功能。   [![godoc][GoDoc]](https://godoc.org/github.com/xiaonanln/goworld)   [![包含中文文档][CN]](https://github.com/xiaonanln/goworld)
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1290** 实现了[Simple DirectMedia Layer](https://www.libsdl.org/)。   [![最近一周有更新][Green]](https://github.com/veandco/go-sdl2)
* [nano](https://github.com/lonng/nano) **star:1212** 轻量级、方便、高性能的基于golang的游戏服务器框架。   [![godoc][GoDoc]](https://godoc.org/github.com/lonng/nano)   [![包含中文文档][CN]](https://github.com/lonng/nano)
* [engo](https://github.com/EngoEngine/engo) **star:1178** 开源 2D 游戏引擎。它遵循 Entity-Component-System 范式。   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [termloop](https://github.com/JoelOtter/termloop) **star:1109** 基于终端的 Go 游戏引擎，建立在 Termbox 之上。   [![godoc][GoDoc]](https://godoc.org/github.com/JoelOtter/termloop)
* [gonet](https://github.com/xtaci/gonet) **star:1089** 实现了游戏服务器骨架。   [![最近一年没有更新][Yellow]](https://github.com/xtaci/gonet)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [g3n](https://github.com/g3n/engine) **star:966**  3D游戏引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/g3n/engine)
* [Oak](https://github.com/oakmound/oak) **star:715** 纯 Go 实现的游戏引擎。   [![最近一周有更新][Green]](https://github.com/oakmound/oak)   [![godoc][GoDoc]](https://godoc.org/github.com/oakmound/oak)
* [Pitaya](https://github.com/topfreegames/pitaya) **star:494** 可伸缩的游戏服务器框架，支持集群和客户端库的iOS, Android, Unity。   [![最近一周有更新][Green]](https://github.com/topfreegames/pitaya)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/pitaya)
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:444** 实现了 [raylib](http://www.raylib.com/)，一个简单易用的库，用于学习视频游戏编程。
* [Azul3D](https://github.com/azul3d/engine) **star:443** 用Go编写的 3D 游戏引擎。
* [go-astar](https://github.com/beefsack/go-astar) **star:358** 实现了A\*路径查找算法。   [![最近一年没有更新][Yellow]](https://github.com/beefsack/go-astar)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-astar)
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:319** 用 OpenGL 编写的 2D 游戏引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/vova616/GarageEngine)
* [go3d](https://github.com/ungerik/go3d) **star:174** 以性能为主的2D/3D数学相关包。   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go3d)
* [glop](https://github.com/runningwild/glop) **star:77** Glop (Game Library Of Power) 是一个相当简单的跨平台游戏库。   [![最近一年没有更新][Yellow]](https://github.com/runningwild/glop)
* [go-collada](https://github.com/GlenKelley/go-collada) **star:15** 处理Collada文件。   [![最近一年没有更新][Yellow]](https://github.com/GlenKelley/go-collada)   [![godoc][GoDoc]](https://godoc.org/github.com/GlenKelley/go-collada)

## 代码生成与泛型

*增强语言的工具，例如通过代码生成支持泛型。 (翻译出错了? 试试 [英文版](README_EN.md#generation-and-generics) 吧~)*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1997** 提供类似 .NET LINQ 的查询方法。   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
* [jennifer](https://github.com/dave/jennifer) **star:1470** 不使用模板生成任意 Go 代码。   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [gen](https://github.com/clipperhouse/gen) **star:1094** 用于生成泛型等类似方法的功能代码生成工具。   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [goderive](https://github.com/awalterschulze/goderive) **star:789** 从输入类型来派生函数。   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/goderive)
* [GoWrap](https://github.com/hexdigest/gowrap) **star:331** 使用简单模板为 Go 接口生成装饰器。   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gowrap)
* [interfaces](https://github.com/rjeczalik/interfaces) **star:208** 用于生成接口定义的命令行工具。   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [go-enum](https://github.com/abice/go-enum) **star:101** 从代码注释中生成枚举。   [![godoc][GoDoc]](https://godoc.org/github.com/abice/go-enum)
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:92** 用于包作用域反射的 Go 预处理器。   [![最近一年没有更新][Yellow]](https://github.com/ungerik/pkgreflect)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/pkgreflect)
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:44** 代码生成工具，可以不通过内存分配就可以高效的将interface{}转换为不可变类型，。   [![最近一年没有更新][Yellow]](https://github.com/t0pep0/efaceconv)   [![godoc][GoDoc]](https://godoc.org/github.com/t0pep0/efaceconv)
* [gotype](https://github.com/wzshiming/gotype) **star:30** Golang 源码解析，用法类似reflect(反射)。   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   [![包含中文文档][CN]](https://github.com/wzshiming/gotype)
* [generis](https://github.com/senselogic/GENERIS) **star:21** 提供泛型、free-form 宏、条件编译和HTML模板的代码生成工具。
* [go-xray](https://github.com/pieterclaerhout/go-xray) **star:6** 帮助更容易地使用反射。   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-xray)
* [typeregistry](https://github.com/xiaoxin01/typeregistry) **star:3** 动态创建类型的库。   [![godoc][GoDoc]](https://godoc.org/github.com/xiaoxin01/typeregistry)

## 地理

*地理工具和服务器 (翻译出错了? 试试 [英文版](README_EN.md#geographic) 吧~)*

* [Tile38](https://github.com/tidwall/tile38) **star:6729** 具有空间索引和实时地理定位功能的地理定位数据库。   [![star > 2000][Awesome]](https://github.com/tidwall/tile38)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/tile38)
* [S2 geometry](https://github.com/golang/geo) **star:1016** S2 geometry 库。   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [mbtileserver](https://github.com/consbio/mbtileserver) **star:132** 一个简单的基于go的服务器，用于存储mbtiles格式的地图块。   [![godoc][GoDoc]](https://godoc.org/github.com/consbio/mbtileserver)
* [geocache](https://github.com/melihmucuk/geocache) **star:119** 基于内存缓存的的地理位置的应用程序。   [![最近一年没有更新][Yellow]](https://github.com/melihmucuk/geocache)   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [osm](https://github.com/paulmach/osm) **star:101** 用于读取、写入和处理 OpenStreetMap 数据和 APIs。   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/osm)
* [WGS84](https://github.com/wroge/wgs84) **star:50** 坐标转换和转换库(ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM)。   [![godoc][GoDoc]](https://godoc.org/github.com/wroge/wgs84)
* [geoserver](https://github.com/hishamkaram/geoserver) **star:33** 基于geoserver REST API的 geoserver 实例。   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/geoserver)
* [gismanager](https://github.com/hishamkaram/gismanager) **star:26** 将你的 GIS 数据(矢量数据)发布到 PostGIS 和 Geoserver。   [![最近一年没有更新][Yellow]](https://github.com/hishamkaram/gismanager)   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)
* [pbf](https://github.com/maguro/pbf) **star:20** 基于Golang 的 OpenStreetMap PBF 编码器/解码器。   [![godoc][GoDoc]](https://godoc.org/github.com/maguro/pbf)

## Go 编译器

*可将 Go 转换为其他语言的编译工具。 (翻译出错了? 试试 [英文版](README_EN.md#go-compilers) 吧~)*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:9236** 将 Go 编译成 JavaScript。   [![star > 2000][Awesome]](https://github.com/gopherjs/gopherjs)   [![godoc][GoDoc]](https://godoc.org/github.com/gopherjs/gopherjs)
* [llgo](https://github.com/go-llvm/llgo) **star:1043** 基于 llvm 的编译器。   [![最近一年没有更新][Yellow]](https://github.com/go-llvm/llgo)   [![godoc][GoDoc]](https://godoc.org/github.com/go-llvm/llgo)
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:396** Golang 转换为 Haxe，再转换为 CPP/CSharp/Java/JavaScript 的编译器.   [![最近一年没有更新][Yellow]](https://github.com/tardisgo/tardisgo)   [![godoc][GoDoc]](https://godoc.org/github.com/tardisgo/tardisgo)
* [c4go](https://github.com/Konstantin8105/c4go) **star:203** 将 C 代码转换为 Go 代码。   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:23** 将 FORTRAN 77 转换为 Go代码。   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)

## Goroutines

*管理和处理 Goroutines 的工具。 (翻译出错了? 试试 [英文版](README_EN.md#goroutines) 吧~)*

* [ants](https://github.com/panjf2000/ants) **star:3195** 一个高性能和低成本的goroutine池在围棋。   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   [![最近一周有更新][Green]](https://github.com/panjf2000/ants)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   [![包含中文文档][CN]](https://github.com/panjf2000/ants)
* [goworker](https://github.com/benmanns/goworker) **star:2361** 基于 go 的后台 worker。   [![star > 2000][Awesome]](https://github.com/benmanns/goworker)   [![godoc][GoDoc]](https://godoc.org/github.com/benmanns/goworker)
* [tunny](https://github.com/Jeffail/tunny) **star:1575** golang 的协程池。   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/tunny)
* [pool](https://github.com/go-playground/pool) **star:547** 有限消费者协程或无限协程池，可用于更加简单的处理和取消协程   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [grpool](https://github.com/ivpusic/grpool) **star:543** 轻量级协程池。   [![最近一年没有更新][Yellow]](https://github.com/ivpusic/grpool)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [workerpool](https://github.com/gammazero/workerpool) **star:271** 限制任务执行并发数，而不是队列中的任务数量的协程池，。   [![最近一周有更新][Green]](https://github.com/gammazero/workerpool)   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [go-floc](https://github.com/workanator/go-floc) **star:177** 轻松编排 goroutines。   [![最近一年没有更新][Yellow]](https://github.com/workanator/go-floc)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-floc)
* [gowp](https://github.com/xxjwxc/gowp) **star:169** gowp是高并发性限制异步工作池。   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gowp)   [![包含中文文档][CN]](https://github.com/xxjwxc/gowp)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:127** 控制 goroutines 的执行顺序。   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [GoSlaves](https://github.com/themester/GoSlaves) **star:91** 简单异步的协程池。   [![godoc][GoDoc]](https://godoc.org/github.com/themester/GoSlaves)
* [semaphore](https://github.com/kamilsk/semaphore) **star:82** 信号量模式实现，可根据通道和上下文进行具备超时功能的锁定/解锁操作。   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/semaphore)
* [semaphore](https://github.com/marusama/semaphore) **star:81** 基于 CAS 的可快速调整的信号量实现(比基于通道的信号量实现更快)。   [![最近一年没有更新][Yellow]](https://github.com/marusama/semaphore)   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/semaphore)
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:66** manages a resizeable pool of context-aware goroutines to bound concurrency   [![godoc][GoDoc]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
* [breaker](https://github.com/kamilsk/breaker) **star:56** 灵活的机制，可以使执行流可中断。   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/breaker)
* [worker-pool](https://github.com/vardius/worker-pool) **star:53** 一个简单的 Go 异步工作池。   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/worker-pool)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:44** 基于 Go 的 CyclicBarrier 实现。   [![最近一年没有更新][Yellow]](https://github.com/marusama/cyclicbarrier)   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/cyclicbarrier)
* [gollback](https://github.com/vardius/gollback) **star:32** 异步简单的函数实用程序，用于管理闭包和回调的执行。   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [async](https://github.com/studiosol/async) **star:30** 一种异步执行函数的安全方法，在出现 panic 时恢复它们。   [![godoc][GoDoc]](https://godoc.org/github.com/studiosol/async)
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:26** 并行运行函数。   [![最近一年没有更新][Yellow]](https://github.com/rafaeljesus/parallel-fn)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/parallel-fn)
* [threadpool](https://github.com/shettyh/threadpool) **star:26** Golang 的 threadpool 实现。   [![最近一年没有更新][Yellow]](https://github.com/shettyh/threadpool)   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [artifex](https://github.com/borderstech/artifex) **star:25** 简单的内存作业队列。   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/artifex)
* [oversight](https://cirello.io/oversight)  完整的实现了Erlang supervision trees。
* [kyoo](https://github.com/dirkaholic/kyoo) **star:24** 提供无限的作业队列和并发工作池。   [![godoc][GoDoc]](https://godoc.org/github.com/dirkaholic/kyoo)
* [Hunch](https://github.com/AaronJan/Hunch) **star:23** Hunch 提供了诸如 All、First、Retry、Waterfall 等功能，这使得异步流控制更加直观。   [![godoc][GoDoc]](https://godoc.org/github.com/AaronJan/Hunch)
* [stl](https://github.com/ssgreg/stl) **star:11** 基于软件事务内存(STM)并发控制机制的软件事务锁。   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/stl)
* [gohive](https://github.com/loveleshsharma/gohive) **star:11** 一个高性能和易于使用的Goroutine池去。   [![godoc][GoDoc]](https://godoc.org/github.com/loveleshsharma/gohive)
* [routine](https://github.com/x-mod/routine) **star:9** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![godoc][GoDoc]](https://godoc.org/github.com/x-mod/routine)
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) **star:7** 像“同步。有错误处理和并发控制。   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)
* [conexec](https://github.com/ITcathyh/conexec) **star:5** 一个并发工具包，帮助以高效和安全的方式并发执行函数。它支持指定总的超时来避免阻塞，并使用goroutine池来提高效率。   [![godoc][GoDoc]](https://godoc.org/github.com/ITcathyh/conexec)
* [go-trylock](https://github.com/subchen/go-trylock) **star:5** 支持 read-write 锁。   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-trylock)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** 轻量级的协程池库，可以通过简单的API来管理。   [![godoc][GoDoc]](https://godoc.org/github.com/nikhilsaraf/go-tools)
* [queue](https://github.com/AnikHasibul/queue) **star:4** 提供类似队列组可访问性 sync.WaitGroup 的功能。帮助你节流和限制协程、等待所有协程结束以及更多功能。   [![godoc][GoDoc]](https://godoc.org/github.com/AnikHasibul/queue)

## GUI

*用于构建GUI应用程序的库。 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

*工具包 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [fyne](https://github.com/fyne-io/fyne) **star:9005** 基于材料设计的Go跨平台本机gui设计。支持:Linux, macOS, Windows, BSD, iOS和Android。   [![star > 2000][Awesome]](https://github.com/fyne-io/fyne)   [![最近一周有更新][Green]](https://github.com/fyne-io/fyne)   [![godoc][GoDoc]](https://godoc.org/github.com/fyne-io/fyne)
* [ui](https://github.com/andlabs/ui) **star:7397** 跨平台的 Platform-native GUI 库。   [![star > 2000][Awesome]](https://github.com/andlabs/ui)   [![godoc][GoDoc]](https://godoc.org/github.com/andlabs/ui)
* [Wails](https://wails.app)  Mac, Windows, Linux桌面应用程序，主要基于含有内置的OS HTML渲染器的HTML UI。
* [qt](https://github.com/therecipe/qt) **star:7089** 实现了 Qt 的 Go接口(支持Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi)。   [![star > 2000][Awesome]](https://github.com/therecipe/qt)   [![godoc][GoDoc]](https://godoc.org/github.com/therecipe/qt)
* [webview](https://github.com/zserge/webview) **star:5591** 跨平台webview窗口，具有简单的双向JavaScript绑定(Windows / macOS / Linux)。   [![star > 2000][Awesome]](https://github.com/zserge/webview)
* [walk](https://github.com/lxn/walk) **star:4311** Windows应用程序库。   [![star > 2000][Awesome]](https://github.com/lxn/walk)   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [app](https://github.com/murlokswarm/app) **star:3470** 用于创建包含了 GO, HTML 和 CSS 的应用程序。支持 MacOS, Windows 正在开发中。   [![star > 2000][Awesome]](https://github.com/murlokswarm/app)   [![最近一周有更新][Green]](https://github.com/murlokswarm/app)   [![godoc][GoDoc]](https://godoc.org/github.com/murlokswarm/app)
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:3074** 使用 GO 和 HTML/JS/CSS (电子驱动)进行构建跨平台 GUI 应用程序。   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  实现了 GTK 的 Go接口。
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1651** 实现了 Sciter 的 Go 接口 : 用于现代桌面 UI 开发的可嵌入HTML/CSS/脚本引擎。可跨平台。   [![godoc][GoDoc]](https://godoc.org/github.com/sciter-sdk/go-sciter)
* [gotk3](https://github.com/gotk3/gotk3) **star:945** 实现了 GTK3 的 Go接口。   [![最近一周有更新][Green]](https://github.com/gotk3/gotk3)   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd) **star:243** 跨平台、快速、简单的桌面UI开发，采用了GO, HTML, CSS和NW.js实现。   [![godoc][GoDoc]](https://godoc.org/github.com/dtylman/gowd)

*交互 (翻译出错了? 试试 [英文版](README_EN.md#gui) 吧~)*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4877** 实现跨平台的GUI系统自动化。包含了控制鼠标、键盘等功能。   [![star > 2000][Awesome]](https://github.com/go-vgo/robotgo)   [![最近一周有更新][Green]](https://github.com/go-vgo/robotgo)
* [systray](https://github.com/getlantern/systray) **star:1029** 跨平台 Go 库，可在通知区放置图标和菜单。   [![最近一周有更新][Green]](https://github.com/getlantern/systray)   [![godoc][GoDoc]](https://godoc.org/github.com/getlantern/systray)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:512** OSX 桌面通知库。   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/gosx-notifier)
* [trayhost](https://github.com/shurcooL/trayhost) **star:175** 跨平台 Go 库，可用于在主机操作系统的任务栏中放置图标。   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:6** 实现了 libappindicator3 C库 的 Go接口。   [![godoc][GoDoc]](https://godoc.org/github.com/dawidd6/go-appindicator)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:5** 用于通知计算机上的任何(可插入的)活动的 OSX 库。   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:2** OSX 睡眠/唤醒通知。   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)


## 硬件

*硬件交互相关的库、工具和教程。 (翻译出错了? 试试 [英文版](README_EN.md#hardware) 吧~)*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## 图片

*图像处理相关的库。 (翻译出错了? 试试 [英文版](README_EN.md#images) 吧~)*

* [gocv](https://github.com/hybridgroup/gocv) **star:3015** 使用OpenCV 3.3+实现的计算机视觉(ComputerVision)的Go语言包。   [![star > 2000][Awesome]](https://github.com/hybridgroup/gocv)   [![最近一周有更新][Green]](https://github.com/hybridgroup/gocv)   [![godoc][GoDoc]](https://godoc.org/github.com/hybridgroup/gocv)
* [imaging](https://github.com/disintegration/imaging) **star:2975** 简单的Go图像处理包。   [![star > 2000][Awesome]](https://github.com/disintegration/imaging)   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/imaging)
* [imaginary](https://github.com/h2non/imaginary) **star:2908** 用于图像大小调整的快速、简单的HTTP微服务。   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [bild](https://github.com/anthonynsimon/bild) **star:2799** 纯Go语言实现的图像处理算法合集。   [![star > 2000][Awesome]](https://github.com/anthonynsimon/bild)   [![godoc][GoDoc]](https://godoc.org/github.com/anthonynsimon/bild)
* [ln](https://github.com/fogleman/ln) **star:2630** Go实现的3D线艺术（3D Line Art）渲染。   [![star > 2000][Awesome]](https://github.com/fogleman/ln)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/ln)
* [resize](https://github.com/nfnt/resize) **star:2302** Go实现的使用常用的插值法（interpolation methods）调整图像大小的库。   [![star > 2000][Awesome]](https://github.com/nfnt/resize)   [![最近一年没有更新][Yellow]](https://github.com/nfnt/resize)   [![godoc][GoDoc]](https://godoc.org/github.com/nfnt/resize)
* [gg](https://github.com/fogleman/gg) **star:2178** 纯Go语言实现的2D渲染。   [![star > 2000][Awesome]](https://github.com/fogleman/gg)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/gg)
* [pt](https://github.com/fogleman/pt) **star:1838** Go实现的路径跟踪（path tracing）引擎。   [![最近一年没有更新][Yellow]](https://github.com/fogleman/pt)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/pt)
* [svgo](https://github.com/ajstarks/svgo) **star:1438**  Go实现的SVG生成库。   [![godoc][GoDoc]](https://godoc.org/github.com/ajstarks/svgo)
* [smartcrop](https://github.com/muesli/smartcrop) **star:1338** 为任意图片寻找合适的位置进行图片裁剪。   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/smartcrop)
* [gift](https://github.com/disintegration/gift) **star:1304** 图像处理包。   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/gift)
* [picfit](https://github.com/thoas/picfit) **star:1206** Go实现的图像调整服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/picfit)
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1162** OpenCV库的Go bindings。   [![godoc][GoDoc]](https://godoc.org/github.com/lazywei/go-opencv)
* [imagick](https://github.com/gographics/imagick) **star:1126**  ImageMagick下MagickWand的C API的Go binding。   [![godoc][GoDoc]](https://godoc.org/github.com/gographics/imagick)
* [geopattern](https://github.com/pravj/geopattern) **star:1055** 由字符串创建漂亮图案的图片生成器。   [![最近一年没有更新][Yellow]](https://github.com/pravj/geopattern)   [![godoc][GoDoc]](https://godoc.org/github.com/pravj/geopattern)
* [bimg](https://github.com/h2non/bimg) **star:906** 使用libvips实现的快速高效的图像处理包。   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:797**  Go实现的LSB隐写（LSB steganography），能够隐藏任何文件到一个图像中。   [![最近一周有更新][Green]](https://github.com/DimitarPetrov/stegify)   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [canvas](https://github.com/tdewolff/canvas) **star:419** 矢量图形到PDF, SVG或光栅图像。   [![最近一周有更新][Green]](https://github.com/tdewolff/canvas)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/canvas)
* [mort](https://github.com/aldor007/mort) **star:388** Go语言实现的图像存储和处理服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/aldor007/mort)
* [image2ascii](https://github.com/qeesung/image2ascii) **star:381** 将图像转换为ASCII码。   [![最近一年没有更新][Yellow]](https://github.com/qeesung/image2ascii)   [![godoc][GoDoc]](https://godoc.org/github.com/qeesung/image2ascii)
* [govatar](https://github.com/o1egl/govatar) **star:347** 生成有趣头像的库和CMD工具。   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/govatar)
* [go-nude](https://github.com/koyachi/go-nude) **star:300** Go语言实现的裸照检测工具。   [![最近一年没有更新][Yellow]](https://github.com/koyachi/go-nude)   [![godoc][GoDoc]](https://godoc.org/github.com/koyachi/go-nude)
* [goimagehash](https://github.com/corona10/goimagehash) **star:279**  图像哈希处理的Go语言包。   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [rez](https://github.com/bamiaux/rez) **star:200** 纯Go语言和SIMD实现的图像大小调整。   [![最近一年没有更新][Yellow]](https://github.com/bamiaux/rez)   [![godoc][GoDoc]](https://godoc.org/github.com/bamiaux/rez)
* [img](https://github.com/hawx/img) **star:133** 图像处理工具的集合。   [![最近一年没有更新][Yellow]](https://github.com/hawx/img)   [![godoc][GoDoc]](https://godoc.org/github.com/hawx/img)
* [darkroom](https://github.com/gojek/darkroom) **star:120** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [go-cairo](https://github.com/ungerik/go-cairo) **star:92**  cairo图形库的Go binding。   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [mergi](https://github.com/noelyahan/mergi) **star:89** 用于图像处理(合并、裁剪、调整大小、水印、动画)的工具和Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/noelyahan/mergi)
* [gltf](https://github.com/qmuntal/gltf) **star:62** 一个高效、健壮的glTF 2.0文件读取、写入和验证器。   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [go-gd](https://github.com/bolknote/go-gd) **star:52**  GD库的Go binding。   [![最近一年没有更新][Yellow]](https://github.com/bolknote/go-gd)   [![godoc][GoDoc]](https://godoc.org/github.com/bolknote/go-gd)
* [steganography](https://github.com/auyer/steganography) **star:44** 纯Go实现的LSB隐写（LSB steganography）的库。   [![godoc][GoDoc]](https://godoc.org/github.com/auyer/steganography)
* [cameron](https://github.com/aofei/cameron) **star:41** 一个Go语言的头像生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)
* [goimghdr](https://github.com/corona10/goimghdr) **star:34** Go语言实现的imghdr模块用于确定文件的图像类型。   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimghdr)
* [tga](https://github.com/ftrvxmtrx/tga) **star:26** tga包是一个TARGA图像的解码、编码器。   [![最近一年没有更新][Yellow]](https://github.com/ftrvxmtrx/tga)   [![godoc][GoDoc]](https://godoc.org/github.com/ftrvxmtrx/tga)
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Python下webcolors库的Go语言调用。   [![最近一年没有更新][Yellow]](https://github.com/jyotiska/go-webcolors)   [![godoc][GoDoc]](https://godoc.org/github.com/jyotiska/go-webcolors)
* [mpo](https://github.com/donatj/mpo) **star:6** MPO三维照片的解码和转换工具。   [![最近一年没有更新][Yellow]](https://github.com/donatj/mpo)   [![godoc][GoDoc]](https://godoc.org/github.com/donatj/mpo)

## 物联网

*物联网设备编程库。 (翻译出错了? 试试 [英文版](README_EN.md#iot-(internet-of-things)) 吧~)*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1377** Flogo是一个面向物联网边缘应用和集成的开源框架。
* [gatt](https://github.com/paypal/gatt) **star:884** Gatt是一个用于构建低能耗蓝牙外围设备的Go语言包。   [![godoc][GoDoc]](https://godoc.org/github.com/paypal/gatt)
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot是一个用于机器人、物理计算和物联网的框架。
* [mainflux](https://github.com/Mainflux/mainflux) **star:818** 工业物联网消息和设备管理服务器。   [![最近一周有更新][Green]](https://github.com/Mainflux/mainflux)   [![godoc][GoDoc]](https://godoc.org/github.com/Mainflux/mainflux)
* [periph](https://periph.io/)  外围设备I/O与低级板(low-level board)设备接口。
* [devices](https://github.com/goiot/devices) **star:229** 一套用于物联网设备的库，实验性地用于x/exp/io。   [![最近一年没有更新][Yellow]](https://github.com/goiot/devices)   [![godoc][GoDoc]](https://godoc.org/github.com/goiot/devices)
* [connectordb](https://github.com/connectordb/connectordb) **star:203** 自我量化和物联网的开源平台。   [![最近一周有更新][Green]](https://github.com/connectordb/connectordb)   [![godoc][GoDoc]](https://godoc.org/github.com/connectordb/connectordb)
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:191** 轻量级物联网流处理引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/sensorbee/sensorbee)
* [huego](https://github.com/amimof/huego) **star:136** 一个包含广泛的Philips Hue客户端的Go语言库。   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT是一个实现谷歌物联网核心设备的简单框架。
* [eywa](https://github.com/xcodersun/eywa) **star:44** Eywa是一个用于跟踪连接的设备连接管理器。   [![最近一年没有更新][Yellow]](https://github.com/xcodersun/eywa)   [![godoc][GoDoc]](https://godoc.org/github.com/xcodersun/eywa)

## 作业调度器

*用于作业调度的库。 (翻译出错了? 试试 [英文版](README_EN.md#job-scheduler) 吧~)*

* [gron](https://github.com/roylee0704/gron) **star:696** 使用简单的Go API定义基于时间的任务。 之后Gron的调度程序将运行它们。   [![最近一周有更新][Green]](https://github.com/roylee0704/gron)   [![godoc][GoDoc]](https://godoc.org/github.com/roylee0704/gron)
* [JobRunner](https://github.com/bamzi/jobrunner) **star:666** 智能和功能丰富的cron作业调度程序（包含任务队列和实时监控）。   [![godoc][GoDoc]](https://godoc.org/github.com/bamzi/jobrunner)
* [jobs](https://github.com/albrow/jobs) **star:468** 持久和灵活的后台作业库。   [![最近一年没有更新][Yellow]](https://github.com/albrow/jobs)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/jobs)
* [scheduler](https://github.com/carlescere/scheduler) **star:320** Cronjobs让调度变得很简单。   [![最近一年没有更新][Yellow]](https://github.com/carlescere/scheduler)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/scheduler)
* [clockwerk](http://github.com/onatm/clockwerk)  使用简单、流畅的语法调度作业的Go语言库。
* [go-cron](https://github.com/rk/go-cron) **star:185** 一个Go实现的简单的定时任务库。可以在不同的时间间隔（每秒一次到在每年在特定的日期执行）执行闭包或函数。主要用于web应用和长时间运行的守护进程。   [![godoc][GoDoc]](https://godoc.org/github.com/rk/go-cron)
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:61** 支持webhook、crons和经典调度的作业调度程序。   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)
* [clockwork](https://github.com/whiteShtef/clockwork) **star:1** Go实现的简单直观的作业调度库。   [![godoc][GoDoc]](https://godoc.org/github.com/whiteShtef/clockwork)

## JSON

*用于JSON处理的库。 (翻译出错了? 试试 [英文版](README_EN.md#json) 吧~)*

* [GJSON](https://github.com/tidwall/gjson) **star:6000** 使用一行代码获取JSON的值。   [![star > 2000][Awesome]](https://github.com/tidwall/gjson)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/gjson)
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2155** 从JSON自动生成Go的结构（struct）定义。   [![star > 2000][Awesome]](https://github.com/ChimeraCoder/gojson)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/gojson)
* [kazaam](https://github.com/Qntfy/kazaam) **star:146** 用于任意JSON文档转换的API。   [![godoc][GoDoc]](https://godoc.org/github.com/Qntfy/kazaam)
* [gojq](https://github.com/elgs/gojq) **star:143** Go语言实现的JSON请求。   [![最近一年没有更新][Yellow]](https://github.com/elgs/gojq)   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/gojq)
* [jsongo](https://github.com/ricardolonga/jsongo) **star:94** 使用Fluent API来更容易地创建Json对象。   [![最近一年没有更新][Yellow]](https://github.com/ricardolonga/jsongo)   [![godoc][GoDoc]](https://godoc.org/github.com/ricardolonga/jsongo)
* [jettison](https://github.com/wI2L/jettison) **star:74** 高性能，无反射的JSON编码器去。   [![最近一周有更新][Green]](https://github.com/wI2L/jettison)   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jettison)
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  将JSON转换为Go的结构（struct）。
* [gjo](https://github.com/skanehira/gjo) **star:71** 用于创建JSON对象的小工具。   [![godoc][GoDoc]](https://godoc.org/github.com/skanehira/gjo)
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:63** 用Go编写的JSON比对工具。   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [jsonf](https://github.com/miolini/jsonf) **star:55** 用于高亮展示和查询JSON的控制台工具。   [![最近一年没有更新][Yellow]](https://github.com/miolini/jsonf)   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/jsonf)
* [mp](https://github.com/sanbornm/mp) **star:38** 简单的cli电子邮件解析器。它目前接受stdin并输出JSON。   [![最近一年没有更新][Yellow]](https://github.com/sanbornm/mp)   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/mp)
* [json2go](https://github.com/m-zajac/json2go) **star:33** 高级JSON去结构转换。提供一个包，该包可以解析多个JSON文档并创建结构体以适应所有这些文档。   [![godoc][GoDoc]](https://godoc.org/github.com/m-zajac/json2go)
* [go-respond](https://github.com/nicklaw5/go-respond) **star:30** 用于处理常见的HTTP JSON响应的Go语言库。   [![godoc][GoDoc]](https://godoc.org/github.com/nicklaw5/go-respond)
* [ajson](https://github.com/spyzhov/ajson) **star:26** 为Go语言开发的包含JSONPath支持的抽象JSON。   [![godoc][GoDoc]](https://godoc.org/github.com/spyzhov/ajson)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** 让自定义结构（struct）转化为JSON兼容的HAL（Hypertext Application Language）返回数据的简单Go包。   [![最近一年没有更新][Yellow]](https://github.com/RichardKnop/jsonhal)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:9** Go-JsonError允许我们轻松创建符合JsonApi规范的json响应错误。   [![godoc][GoDoc]](https://godoc.org/github.com/ddymko/go-jsonerror)
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:8** 基于JSON API错误引用的Go bindings。   [![最近一年没有更新][Yellow]](https://github.com/AmuzaTkts/jsonapi-errors)   [![godoc][GoDoc]](https://godoc.org/github.com/AmuzaTkts/jsonapi-errors)
* [ej](https://github.com/lucassscaravelli/ej) **star:3** 简洁地编写和读取来自不同来源的JSON。   [![godoc][GoDoc]](https://godoc.org/github.com/lucassscaravelli/ej)
* [mapslice-json](https://github.com/mickep76/mapslice-json) **star:1** 去MapSlice的有序封送/解封的JSON地图。   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/mapslice-json)

## 日志记录

*用于生成和处理日志文件的库。 (翻译出错了? 试试 [英文版](README_EN.md#logging) 吧~)*

* [logrus](https://github.com/Sirupsen/logrus) **star:14179** Go的结构化日志操作 。   [![star > 2000][Awesome]](https://github.com/Sirupsen/logrus)   [![最近一周有更新][Green]](https://github.com/Sirupsen/logrus)   [![godoc][GoDoc]](https://godoc.org/github.com/Sirupsen/logrus)
* [zap](https://github.com/uber-go/zap) **star:9197** 快速、结构化、多等级的日志记录。   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![最近一周有更新][Green]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [spew](https://github.com/davecgh/go-spew) **star:3770** 为Go数据结构实现一个漂亮的printer用于帮助调试。   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [zerolog](https://github.com/rs/zerolog) **star:3031** Zero-allocation JSON日志记录。   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![最近一周有更新][Green]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [glog](https://github.com/golang/glog) **star:2488** 为Go提供了多等级日志记录。   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1814** 简单的滚动日志，io.WriteCloser的实现。   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [tail](https://github.com/hpcloud/tail) **star:1736** 努力模拟实现BSD的tail的特性的Go包。   [![最近一周有更新][Green]](https://github.com/hpcloud/tail)   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [seelog](https://github.com/cihub/seelog) **star:1443** 具有灵活调度、过滤和格式化的日志功能。   [![最近一年没有更新][Yellow]](https://github.com/cihub/seelog)   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log15](https://github.com/inconshreveable/log15) **star:952** 简单、强大的日志操作。   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/log15)
* [log](https://github.com/apex/log) **star:811** Go的结构化日志包。   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [onelog](https://github.com/francoispqt/onelog) **star:364** Onelog是一个非常简单但非常高效的JSON日志程序。它是所有场景中速度最快的JSON日志程序。而且，它是配置要求最低的日志记录器之一。   [![最近一年没有更新][Yellow]](https://github.com/francoispqt/onelog)   [![godoc][GoDoc]](https://godoc.org/github.com/francoispqt/onelog)
* [logxi](https://github.com/mgutz/logxi) **star:343** 12-factor app的日志程序，快速且让人高兴地使用。   [![godoc][GoDoc]](https://godoc.org/github.com/mgutz/logxi)
* [log](https://github.com/go-playground/log) **star:273** Go的简单、可配置和可伸缩的结构化日志。   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [logutils](https://github.com/hashicorp/logutils) **star:272** Go的用于更好地进行日志操作的实用程序，继承了标准日志库。   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [go-logger](https://github.com/apsdehal/go-logger) **star:249** 简单的日志程序的 Go 程序，与级别处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/apsdehal/go-logger)
* [logger](https://github.com/azer/logger) **star:139** Go的精简日志库。   [![godoc][GoDoc]](https://godoc.org/github.com/azer/logger)
* [xlog](https://github.com/rs/xlog) **star:133** 针对'net/context`实现的结构化的记录器，用于HTTP处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:132** RollingWriter是一个自动循环的io.Writer的实现,带有多种策略以提供日志文件循环(rotation)。   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:112** 支持日志多等级、分类和过滤的高性能日志记录。可以发送过滤后的日志消息到各种目标(如控制台，网络，邮件)。   [![最近一年没有更新][Yellow]](https://github.com/go-ozzo/ozzo-log)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-log)   [![包含中文文档][CN]](https://github.com/go-ozzo/ozzo-log)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:84** 用Go编写的功能齐全的日志写入saas。   [![最近一年没有更新][Yellow]](https://github.com/firstrow/logvoyage)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [glg](https://github.com/kpango/glg) **star:68** glg是一个简单而快速的Go日志库。   [![最近一周有更新][Green]](https://github.com/kpango/glg)   [![godoc][GoDoc]](https://godoc.org/github.com/kpango/glg)
* [stdlog](https://github.com/alexcesaro/log) **star:42** Stdlog是一个面向对象的库，提供了多等级日志记录。它对cron任务非常有用。   [![最近一年没有更新][Yellow]](https://github.com/alexcesaro/log)   [![godoc][GoDoc]](https://godoc.org/github.com/alexcesaro/log)
* [gologger](https://github.com/sadlil/gologger) **star:38** 为Go提供方便简单的日志操作; 在彩色控制台，简单控制台，文件或Elasticsearch上。   [![最近一年没有更新][Yellow]](https://github.com/sadlil/gologger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/gologger)   [![归档项目][Archived]](https://github.com/sadlil/gologger)
* [logex](https://github.com/chzyer/logex) **star:35** 由标准日志库封装的Go日志库，支持跟踪和多等级。   [![最近一年没有更新][Yellow]](https://github.com/chzyer/logex)   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/logex)
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j的Go语言。   [![最近一年没有更新][Yellow]](https://github.com/ian-kent/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/go-log)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:27** 基于当前日期和时间的自动日志文件写入工具，类似cronolog。   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [go-log](https://github.com/siddontang/go-log) **star:26** 支持多等级和多处理程序的日志库。   [![最近一年没有更新][Yellow]](https://github.com/siddontang/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-log)
* [logrusly](https://github.com/sebest/logrusly) **star:25** [logrus](https://github.com/sirupsen/logrus)的插件，将错误信息发送到[Loggly](https://www.loggly.com/)。   [![最近一年没有更新][Yellow]](https://github.com/sebest/logrusly)   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/logrusly)
* [log](https://github.com/teris-io/log) **star:23** Go的结构化日志接口，清晰地将日志facade与其实现(implementation)分离开来。   [![最近一年没有更新][Yellow]](https://github.com/teris-io/log)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/log)
* [distillog](https://github.com/amoghe/distillog) **star:22** distilled日志记录(可以将其视为stdlib +日志)。   [![最近一年没有更新][Yellow]](https://github.com/amoghe/distillog)   [![godoc][GoDoc]](https://godoc.org/github.com/amoghe/distillog)
* [journald](https://github.com/ssgreg/journald) **star:22**  Go实现systemd Journal的原生API用于日志记录。   [![最近一年没有更新][Yellow]](https://github.com/ssgreg/journald)   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/journald)
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** 简单的go日志模块，有5个级别，可选循环(rotation)日志文件记录功能和stdout/stderr输出。   [![最近一年没有更新][Yellow]](https://github.com/jbrodriguez/mlog)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrodriguez/mlog)
* [gomol](https://github.com/aphistic/gomol) **star:16** 为Go实现可多方式输出、结构化日志, 并可扩展日志输出方式。   [![最近一年没有更新][Yellow]](https://github.com/aphistic/gomol)   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  快速、可扩展、功能齐全、std-lib源兼容的日志库。
* [go-log](https://github.com/subchen/go-log) **star:10** Go实现的简单且可配置的日志工具，并带有多等级、日志格式化和日志写入工具。   [![最近一年没有更新][Yellow]](https://github.com/subchen/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-log)
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** 用于多等级级日志记录的包。   [![最近一年没有更新][Yellow]](https://github.com/ewwwwwqm/logdump)   [![godoc][GoDoc]](https://godoc.org/github.com/ewwwwwqm/logdump)
* [glo](https://github.com/lajosbencz/glo) **star:9** 参照PHP的Monolog实现的具有相同日志等级的Go日志库。   [![最近一年没有更新][Yellow]](https://github.com/lajosbencz/glo)   [![godoc][GoDoc]](https://godoc.org/github.com/lajosbencz/glo)
* [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:7** io。作者的实现使用[logrus](https://github.com/sirupsen/logrus) logger。   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/logrusiowriter)
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Go的彩色日志记录器，带有可配置的日志级别。   [![最近一年没有更新][Yellow]](https://github.com/borderstech/logmatic)   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [log](https://github.com/aerogo/log) **star:6** 一个O(1)日志系统，允许您将一个日志连接到多个日志写入(例如stdout、文件和TCP连接)。   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/log)
* [xlog](https://github.com/xfxdev/xlog) **star:6** 插件架构和灵活的日志系统，带有多级别、多日志目标和自定义日志格式。   [![最近一年没有更新][Yellow]](https://github.com/xfxdev/xlog)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [logo](https://github.com/mbndr/logo) **star:5** Go的日志工具，可配置的日志写入器。   [![最近一年没有更新][Yellow]](https://github.com/mbndr/logo)   [![godoc][GoDoc]](https://godoc.org/github.com/mbndr/logo)
* [go-log](https://github.com/pieterclaerhout/go-log) **star:3** 具有strack跟踪、对象转储和可选时间戳的日志记录库。   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-log)

## 机器学习

*机器学习库。 (翻译出错了? 试试 [英文版](README_EN.md#machine-learning) 吧~)*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:7097** 通用机器学习库。   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   [![包含中文文档][CN]](https://github.com/sjwhitworth/golearn)
* [gorgonia](https://github.com/gorgonia/gorgonia) **star:3203** 基于图形（graph-based）的计算库，如Theano：它为构建各种机器学习和神经网络算法提供了基本框架。   [![star > 2000][Awesome]](https://github.com/gorgonia/gorgonia)   [![godoc][GoDoc]](https://godoc.org/github.com/gorgonia/gorgonia)
* [tfgo](https://github.com/galeone/tfgo) **star:1325** 易于使用的Tensorflow bindings:简化了官方Tensorflow Go bindings的使用。在Go中定义计算图形，在Python中加载和执行训练的模型。   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/tfgo)
* [goml](https://github.com/cdipaolo/goml) **star:1084** 在线机器学习。   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gosseract](https://github.com/otiai10/gosseract) **star:1084** 使用c++的Tesseract库实现的OCR。   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [gorse](https://github.com/zhenghaoz/gorse) **star:935** 基于协同过滤的离线推荐系统后端。   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   [![包含中文文档][CN]](https://github.com/zhenghaoz/gorse)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:667** 一个进化优化（evolutionary optimization）库。   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:667** 快速、灵活、多线程集成的决策树，用于机器学习。   [![最近一年没有更新][Yellow]](https://github.com/ryanbressler/CloudForest)   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [bayesian](https://github.com/jbrukh/bayesian) **star:655** Go的朴素贝叶斯分类。   [![最近一年没有更新][Yellow]](https://github.com/jbrukh/bayesian)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [gobrain](https://github.com/goml/gobrain) **star:433** 用 Go 编写的神经网络库。   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [ocrserver](https://github.com/otiai10/ocrserver) **star:274** 一个简单的OCR API服务器，非常容易地使用Docker和Heroku部署。   [![最近一年没有更新][Yellow]](https://github.com/otiai10/ocrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/ocrserver)
* [regommend](https://github.com/muesli/regommend) **star:266** 推荐和协同过滤引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [go-deep](https://github.com/patrikeh/go-deep) **star:256** 一个功能丰富的神经网络库 。   [![godoc][GoDoc]](https://godoc.org/github.com/patrikeh/go-deep)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:250** Go Interface， 用于开放式神经网络交换(Open Neural Network Exchange)。   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-galib](https://github.com/thoj/go-galib) **star:177** 用Go编写的遗传算法库。   [![最近一年没有更新][Yellow]](https://github.com/thoj/go-galib)   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:156** 用Go编写的推荐算法库。   [![最近一年没有更新][Yellow]](https://github.com/timkaye11/goRecommend)   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [shield](https://github.com/eaigner/shield) **star:129** 贝叶斯文本分类器，具有灵活的tokenizers和存储后端。   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [Goptuna](https://github.com/c-bata/goptuna) **star:119** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![最近一周有更新][Green]](https://github.com/c-bata/goptuna)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [go-fann](https://github.com/white-pony/go-fann) **star:103** 快速人工神经网络(FANN)库的Go bindings。   [![最近一年没有更新][Yellow]](https://github.com/white-pony/go-fann)   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [goga](https://github.com/tomcraven/goga) **star:86** Go的遗传算法库。   [![最近一年没有更新][Yellow]](https://github.com/tomcraven/goga)   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:64** 基于LIBSVM 3.14实现。   [![最近一年没有更新][Yellow]](https://github.com/datastream/libsvm)   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [neural-go](https://github.com/schuyler/neural-go) **star:61** 多层感知器网络在Go中的实现，使用反向传播算法进行训练。   [![最近一年没有更新][Yellow]](https://github.com/schuyler/neural-go)   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [go-pr](https://github.com/daviddengcn/go-pr) **star:59** Go编写的模式识别包。   [![最近一年没有更新][Yellow]](https://github.com/daviddengcn/go-pr)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [gonet](https://github.com/dathoangnd/gonet) **star:58** 用于围棋的神经网络。   [![godoc][GoDoc]](https://godoc.org/github.com/dathoangnd/gonet)
* [neat](https://github.com/jinyeom/neat) **star:56** 即插即用的并行Go框架，用于增强拓扑的神经进化(NeuroEvolution of Augmenting Topologies)。   [![最近一年没有更新][Yellow]](https://github.com/jinyeom/neat)   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   [![归档项目][Archived]](https://github.com/jinyeom/neat)
* [goscore](https://github.com/asafschers/goscore) **star:45**  为预言模型标记语言（PMML）实现的评分API。   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [golinear](https://github.com/danieldk/golinear) **star:39**  Go实现的liblinear bindings。   [![最近一年没有更新][Yellow]](https://github.com/danieldk/golinear)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [fonet](https://github.com/Fontinalis/fonet) **star:36** 一个用Go编写的深度神经网络库。   [![godoc][GoDoc]](https://godoc.org/github.com/Fontinalis/fonet)
* [Varis](https://github.com/Xamber/Varis) **star:28** Go实现的神经网络。   [![最近一年没有更新][Yellow]](https://github.com/Xamber/Varis)   [![godoc][GoDoc]](https://godoc.org/github.com/Xamber/Varis)
* [godist](https://github.com/e-dard/godist) **star:25** 各种概率分布，以及相关的method。   [![最近一年没有更新][Yellow]](https://github.com/e-dard/godist)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/godist)
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:22** Go实现的k-modes和k-prototypes聚类算法。   [![最近一年没有更新][Yellow]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
* [probab](https://github.com/ThePaw/probab) **star:12** 概率分布函数。贝叶斯推理。使用Go写的。   [![最近一年没有更新][Yellow]](https://github.com/ThePaw/probab)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/probab)
* [evoli](https://github.com/khezen/evoli) **star:9** 遗传算法（Genetic Algorithm）和粒子群优化（Particle Swarm Optimization）库。   [![最近一周有更新][Green]](https://github.com/khezen/evoli)   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)
* [GoMind](https://github.com/surenderthakran/gomind) **star:7** 一个简单的神经网络库。   [![最近一年没有更新][Yellow]](https://github.com/surenderthakran/gomind)   [![godoc][GoDoc]](https://godoc.org/github.com/surenderthakran/gomind)
* [randomforest](https://github.com/malaschitz/randomForest) **star:3** 容易使用随机森林库去。   [![godoc][GoDoc]](https://godoc.org/github.com/malaschitz/randomForest)

## 消息

*实现消息传递系统的库。 (翻译出错了? 试试 [英文版](README_EN.md#messaging) 吧~)*

* [sarama](https://github.com/Shopify/sarama) **star:5502**  Apache Kafka的Go库。   [![star > 2000][Awesome]](https://github.com/Shopify/sarama)   [![最近一周有更新][Green]](https://github.com/Shopify/sarama)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/sarama)
* [gorush](https://github.com/appleboy/gorush) **star:4246** 使用[APNs2](https://github.com/sideshow/apns2)和谷歌[GCM](https://github.com/google/go-gcm)推送通知服务器。   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![最近一周有更新][Green]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:4133** 实时消息(Websockets或SockJS)服务器。   [![star > 2000][Awesome]](https://github.com/centrifugal/centrifugo)   [![最近一周有更新][Green]](https://github.com/centrifugal/centrifugo)   [![godoc][GoDoc]](https://godoc.org/github.com/centrifugal/centrifugo)
* [machinery](https://github.com/RichardKnop/machinery) **star:3894** 基于分布式消息传递的异步任务/作业队列。   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:3300** go的socket.io库，一个实时应用程序框架。   [![star > 2000][Awesome]](https://github.com/googollee/go-socket.io)   [![godoc][GoDoc]](https://godoc.org/github.com/googollee/go-socket.io)
* [NATS Go Client](https://github.com/nats-io/nats) **star:2710** 轻量级和高性能的发布-订阅(publish-subscribe)和分布式队列消息传递系统——这是一个Go库。   [![star > 2000][Awesome]](https://github.com/nats-io/nats)   [![最近一周有更新][Green]](https://github.com/nats-io/nats)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/nats)
* [APNs2](https://github.com/sideshow/apns2) **star:2270** HTTP / 2苹果消息推送provider——发送推送消息到iOS, tvOS, Safari和OSX应用。   [![star > 2000][Awesome]](https://github.com/sideshow/apns2)   [![godoc][GoDoc]](https://godoc.org/github.com/sideshow/apns2)
* [Benthos](https://github.com/Jeffail/benthos) **star:2251** 一系列协议之间的消息流桥接。   [![star > 2000][Awesome]](https://github.com/Jeffail/benthos)   [![最近一周有更新][Green]](https://github.com/Jeffail/benthos)   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/benthos)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1895** gopush-cluster是一个gopush服务器集群。   [![最近一年没有更新][Yellow]](https://github.com/Terry-Mao/gopush-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   [![包含中文文档][CN]](https://github.com/Terry-Mao/gopush-cluster)
* [Mercure](https://github.com/dunglas/mercure) **star:1890** 使用Mercure协议分派服务器发送(server-sent)更新的服务器和库(构建在服务器发送事件之上)。   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [melody](https://github.com/olahol/melody) **star:1763** 处理websocket session的极简框架，包括广播和自动ping/pong处理。   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1603** NSQ的官方Go包。   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [mangos](https://github.com/go-mangos/mangos) **star:1538** Nanomsg(“可伸缩协议”)的纯go实现,具有传输互操作性。   [![godoc][GoDoc]](https://godoc.org/github.com/go-mangos/mangos)   [![归档项目][Archived]](https://github.com/go-mangos/mangos)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1143** Redis支持的统一推送服务, 用于服务端向移动设备的消息通知。   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [zmq4](https://github.com/pebbe/zmq4) **star:837** ZeroMQ的Go interface第4版。也可用于[第3版](https://github.com/pebbe/zmq3)和[第2版](https://github.com/pebbe/zmq2)。   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/zmq4)
* [Gollum](https://github.com/trivago/gollum) **star:835** 一个n:m多路复用器(n:m multiplexer)，收集不同来源的消息并将其广播到一组目的地。   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [Beaver](https://github.com/Clivern/Beaver) **star:814** 一个实时消息服务器，可用于在web和手机app端构建一个可伸缩的应用内通知，多人游戏，聊天应用。   [![最近一周有更新][Green]](https://github.com/Clivern/Beaver)   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [EventBus](https://github.com/asaskevich/EventBus) **star:656** 具有异步兼容性的轻量级事件总线。   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [golongpoll](https://github.com/jcuga/golongpoll) **star:447** HTTP longpoll服务器库，使web发布-订阅变得简单。   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [dbus](https://github.com/godbus/dbus) **star:430** D-Bus的Go bindings。   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [Asynq](https://github.com/hibiken/asynq) **star:420** 一个简单、可靠、高效的基于Redis的分布式任务队列。   [![最近一周有更新][Green]](https://github.com/hibiken/asynq)   [![godoc][GoDoc]](https://godoc.org/github.com/hibiken/asynq)
* [emitter](https://github.com/olebedev/emitter) **star:344** 使用Go的方式发出事件, 带有通配符、谓词、取消可能性和许多其他优点。   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [Glue](https://github.com/desertbit/glue) **star:332** 健壮的Go和Javascript Socket库(替代Socket.io)。   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:310** 简单的pubsub的go包。   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [guble](https://github.com/smancke/guble) **star:143** 使用推送通知服务(谷歌Firebase云消息、苹果推送通知服务、SMS)的消息服务器, 支持websockets,REST API，并具有分布式操作和消息持久性。   [![最近一年没有更新][Yellow]](https://github.com/smancke/guble)   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [Bus](https://github.com/mustafaturan/bus) **star:142** 内部通信的最小消息总线实现。   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [rabtap](https://github.com/jandelgado/rabtap) **star:111** RabbitMQ的瑞士军刀cli应用。   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [oplog](https://github.com/dailymotion/oplog) **star:102** 用于REST api的通用oplog/replication系统。   [![最近一年没有更新][Yellow]](https://github.com/dailymotion/oplog)   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [messagebus](https://github.com/vardius/message-bus) **star:99** messagebus是一个Go的简单异步消息总线，非常适合在执行事件sourcing、CQRS和DDD时用作事件总线。   [![最近一周有更新][Green]](https://github.com/vardius/message-bus)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:69** amqp exchanges和队列上的一个小工具。   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [drone-line](https://github.com/appleboy/drone-line) **star:67** 使用二进制、docker或Drone CI发送[Line](https://at.line.me/en)通知。   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-line)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:58** RapidMQ是用于管理本地消息队列的轻量且可靠的库。   [![最近一年没有更新][Yellow]](https://github.com/sybrexsys/RapidMQ)   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:57** 一个围绕NSQ topic和channel的小工具。   [![最近一年没有更新][Yellow]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:48** 原生的freedesktop通知规范实现。   [![最近一年没有更新][Yellow]](https://github.com/TheCreeper/go-notify)   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [hub](https://github.com/leandro-lugaresi/hub) **star:48** 适用于Go应用程序的消息/事件中心，使用publish/subscribe模式，并支持别名(类似rabbitMQ exchanges)。   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [Commander](https://github.com/jeroenrinzema/commander) **star:40** 高级事件驱动的消费者/生产者(consumer/producer)，支持各种“方言”，如Apache Kafka。   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [event](https://github.com/agoalofalife/event) **star:33** 观察者模式的实现。   [![最近一年没有更新][Yellow]](https://github.com/agoalofalife/event)   [![godoc][GoDoc]](https://godoc.org/github.com/agoalofalife/event)
* [go-res](https://github.com/jirenius/go-res) **star:24** 用于构建REST/实时服务的包，其中客户端使用NATS和Resgate进行无缝同步。   [![godoc][GoDoc]](https://godoc.org/github.com/jirenius/go-res)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:13** 用于Viessmann Vitotrol web服务的客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:13** 基于Redis streams的队列的生产者和消费者。   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/redisqueue)
* [jazz](https://github.com/socifi/jazz) **star:9** 一个简单的RabbitMQ抽象层，用于队列管理和消息的发布和消费。   [![最近一年没有更新][Yellow]](https://github.com/socifi/jazz)   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** 用Go编写的Gaurun客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gaurun-client)
* [rmqconn](https://github.com/sbabiv/rmqconn) **star:6** RabbitMQ重新连接。amqp.Connection和amqp.Dial之上的Wrapper。允许在强制调用Close()方法关闭连接之前重新连接。   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/rmqconn)
* [ami](https://github.com/kak-tus/ami) **star:6** 基于Redis集群流的客户端到可靠队列。   [![godoc][GoDoc]](https://godoc.org/github.com/kak-tus/ami)

## 微软办公软件

* [unioffice](https://github.com/unidoc/unioffice) **star:2150** 用于创建和处理Office Word (.docx)、Excel (.xlsx)和Powerpoint (.pptx)文档的纯go库。   [![star > 2000][Awesome]](https://github.com/unidoc/unioffice)   [![最近一周有更新][Green]](https://github.com/unidoc/unioffice)   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*用于操作Microsoft Excel的库。 (翻译出错了? 试试 [英文版](README_EN.md#microsoft-excel) 吧~)*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:5833** 用于读写Microsoft Excel™(XLSX)文件的Go语言库。   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   [![最近一周有更新][Green]](https://github.com/360EntSecGroup-Skylar/excelize)   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:3954** 用以简化在Go程序中读取使用最新版本Microsoft Excel的XML格式文件的库。   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:106** 在Go程序以快速和安全的方式读取/更新现有的Microsoft Excel文件。   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:71** 一个简单轻便的阅读器，可以将类关系型数据库(relate-db-like)的excel作为表来读取。   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:15** libxlsxwriter的Go binding, 用于编写XLSX (Microsoft Excel)文件。   [![最近一年没有更新][Yellow]](https://github.com/fterrag/goxlsxwriter)   [![godoc][GoDoc]](https://godoc.org/github.com/fterrag/goxlsxwriter)

## 杂项

### 依赖注入

*用于处理依赖项注入的库。 (翻译出错了? 试试 [英文版](README_EN.md#dependency-injection) 吧~)*

* [dig](https://github.com/uber-go/dig) **star:1304** 一个基于反射的Go依赖注入工具包。   [![最近一周有更新][Green]](https://github.com/uber-go/dig)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:1195** 基于依赖注入的Go应用程序框架(构建在dig之上)。   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [container](https://github.com/golobby/container) **star:77** 一个强大的IoC容器，具有流畅且易于使用的接口。   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/container)
* [dingo](https://github.com/i-love-flamingo/dingo) **star:41** 基于Guice的Go依赖注入工具包。   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/dingo)
* [alice](https://github.com/magic003/alice) **star:38** Go的外挂的依赖注入容器。   [![最近一年没有更新][Yellow]](https://github.com/magic003/alice)   [![godoc][GoDoc]](https://godoc.org/github.com/magic003/alice)
* [wire](https://github.com/Fs02/wire) **star:29** 适用于Go的严格运行时依赖注入(Strict Runtime Dependency Injection)。   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/wire)
* [linker](https://github.com/logrange/linker) **star:21** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [gocontainer](https://github.com/vardius/gocontainer) **star:12** 简单的依赖注入容器。   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)
* [di](https://github.com/goava/di) **star:7** go编程语言的依赖注入容器。   [![最近一周有更新][Green]](https://github.com/goava/di)   [![godoc][GoDoc]](https://godoc.org/github.com/goava/di)

### 项目布局

*用于组织项目的非正式模式集。 (翻译出错了? 试试 [英文版](README_EN.md#project-layout) 吧~)*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:13699** Go生态系统中历史和新兴的项目布局模式集合。   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:566** 应用程序样板和应用现代实践的例子。   [![godoc][GoDoc]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) **star:326** 遵循生产最佳实践快速启动项目的Go应用程序样板模板。   [![godoc][GoDoc]](https://godoc.org/github.com/lacion/cookiecutter-golang)
* [scaffold](https://github.com/catchplay/scaffold) **star:58** 快速生成Go项目布局的脚手架。让您专注于已实现的业务逻辑。   [![最近一年没有更新][Yellow]](https://github.com/catchplay/scaffold)   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [go-sample](https://github.com/zitryss/go-sample) **star:46** 带有实际代码的Go应用程序项目的示例布局。   [![最近一年没有更新][Yellow]](https://github.com/zitryss/go-sample)   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)

### 字符串

*处理字符串的库。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [xstrings](https://github.com/huandu/xstrings) **star:721** 从其他语言移植的有用字符串函数合集。   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:86** 字符串处理工具。   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)

*这些库之所以放在这里，是因为不适合放在其他分类。 (翻译出错了? 试试 [英文版](README_EN.md#strings) 吧~)*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4654** 用于检索进程和系统利用率(CPU、内存、磁盘等)的跨平台的库。   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![最近一周有更新][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:2757** 用于生成和解压.zip和.tar.gz文档的库和命令。   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![最近一周有更新][Green]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gosms](https://github.com/haxpax/gosms) **star:1273** Go编写的私人的本地短信网关，可以用来发送短信。   [![最近一年没有更新][Yellow]](https://github.com/haxpax/gosms)   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:984**  Go语言弹性模式(resiliency pattern)。   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:845** 用go编写的随机数据生成器。   [![最近一周有更新][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:785** base64Captcha支持数字，字母，算术，音频和混合模式的验证码。   [![最近一周有更新][Green]](https://github.com/mojocn/base64Captcha)   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![包含中文文档][CN]](https://github.com/mojocn/base64Captcha)
* [go-openapi](https://github.com/go-openapi)  用于解析和使用开放api模式(open-api schemas)的包的集合。
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:775** Go语言的通用对象池。   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![包含中文文档][CN]](https://github.com/jolestar/go-commons-pool)
* [shortid](https://github.com/teris-io/shortid) **star:533** 分布式地生成超短、唯一、非顺序、URL友好的id。   [![最近一年没有更新][Yellow]](https://github.com/teris-io/shortid)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:503** 用于在纯Go中与LLVM IR交互的库。   [![最近一周有更新][Green]](https://github.com/llir/llvm)   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:380** 易于使用，可扩展的健康检查库。   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:350** conv包提供了跨Go类型(Go types)的快速和直观的转换。   [![最近一年没有更新][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [banner](https://github.com/dimiro1/banner) **star:256** 在Go应用程序中添加漂亮的横幅(banner)。   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:235** 获取国家和细节数据的包。   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [antch](https://github.com/antchfx/antch) **star:169** 一个快速、强大和可扩展的web爬虫框架。   [![最近一年没有更新][Yellow]](https://github.com/antchfx/antch)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   [![包含中文文档][CN]](https://github.com/antchfx/antch)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:155** 美化数据,使其更适合人查看。   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   [![包含中文文档][CN]](https://github.com/go-ffmt/ffmt)
* [stateless](https://github.com/qmuntal/stateless) **star:148** 用于创建状态机的流畅库。   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [ghorg](https://github.com/gabrie30/ghorg) **star:146** 快速克隆整个org用户库到一个目录-支持GitHub, GitLab和Bitbucket。   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [lk](https://github.com/hyperboloide/lk) **star:145** 一个简单的版权许可证库。   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [battery](https://github.com/distatus/battery) **star:143** 跨平台、标准化的电池信息库。   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [stats](https://github.com/go-playground/stats) **star:131** Monitors Go MemStats + 诸如如内存，Swap和CPU的系统状态统计，并通过UDP发送到任何你想记录的地方   [![最近一年没有更新][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [bitio](https://github.com/icza/bitio) **star:116** 高度优化的位级读写器。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:107** 用于RESTful服务的强制的(opinionated)并发健康检查HTTP处理程序。   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:96** Go的Emojis库。   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:87** 用于RAR、TAR、ZIP和7z文件的解压缩库。   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [gommit](https://github.com/antham/gommit) **star:85** 分析git提交消息，确保它们遵循已定义的格式。   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [gotoprom](https://github.com/cabify/gotoprom) **star:80** 为Prometheus客户端提供类型安全的指标(metric)构建工具库。   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [indigo](https://github.com/osamingo/indigo) **star:58** 分布式唯一ID生成器, 使用Sonyflake并由Base58编码。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [morse](https://github.com/alwindoss/morse) **star:57** 实现字符串与摩尔斯电码转换的库。   [![最近一年没有更新][Yellow]](https://github.com/alwindoss/morse)   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [captcha](https://github.com/steambap/captcha) **star:54** captcha包为验证码生成提供了一个易于使用的、unopinionated的API。   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [xkg](https://github.com/go-xkg/xkg) **star:48** X Keyboard Grabber(键盘事件捕获)   [![最近一年没有更新][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [persian](https://github.com/mavihq/persian) **star:40** 一些适用于波斯语的Go工具库。   [![最近一年没有更新][Yellow]](https://github.com/mavihq/persian)   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:39** 通过Json请求生成PDF的HTTP服务。   [![最近一年没有更新][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [datacounter](https://github.com/miolini/datacounter) **star:31** 用于readers/writer/http.ResponseWriter的计数器。   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:31** 用于[Browser Capabilities Project](http://browscap.org/)的Go库。   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [autoflags](https://github.com/artyom/autoflags) **star:26** 从struct字段自动定义命令行flag的Go包。   [![最近一年没有更新][Yellow]](https://github.com/artyom/autoflags)   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg)规范在Go中的实现。   [![最近一年没有更新][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [gosh](https://github.com/osamingo/gosh) **star:20** 提供Go统计处理程序，结构和测量方法。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:19** 一个现代的、强大的、健壮的URL转短链接微服务，带有mysql支持。   [![最近一年没有更新][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  用于生成http输入和输出处理模板。
* [sandid](https://github.com/aofei/sandid) **star:15** 能沟让地球上的每一粒沙子都有自己的ID。   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [anagent](https://github.com/mudler/anagent) **star:11** Go语言的最小化，可插入的evloop/timer处理程序, 带有依赖注入。   [![最近一年没有更新][Yellow]](https://github.com/mudler/anagent)   [![godoc][GoDoc]](https://godoc.org/github.com/mudler/anagent)
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:10** 根据威尔逊得分排序算法(Wilson Score Equation)计算平均分和评分。   [![最近一年没有更新][Yellow]](https://github.com/kirillDanshin/avgRating)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/avgRating)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** 一个用于打包和解包FQDNs列表的go语言库。   [![最近一年没有更新][Yellow]](https://github.com/Wing924/hostutils)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [shellwords](https://github.com/Wing924/shellwords) **star:8** 一个Go库，实现了依照UNIX Bourne shell的关键词解析规则进行字符串操纵。   [![最近一年没有更新][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [metrics](https://github.com/pascaldekloe/metrics) **star:6** 用于instrumentation和Prometheus exposition的库。   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [numa](https://github.com/lrita/numa) **star:2** NUMA是一个用go编写的实用程序库。它可以帮助我们编写一些NUMA-AWARED代码。   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## 自然语言处理

*用于处理人类语言的库。 (翻译出错了? 试试 [英文版](README_EN.md#natural-language-processing) 吧~)*

* [prose](https://github.com/jdkato/prose) **star:2453** 支持标记化、词性标记、命名实体提取等文本处理的库。只说英语。   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1221** 高效的文本分割;支持英语、汉语、日语等。   [![最近一周有更新][Green]](https://github.com/go-ego/gse)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![包含中文文档][CN]](https://github.com/go-ego/gse)
* [when](https://github.com/olebedev/when) **star:1064** 带有可插入规则的自然EN和RU语言日期/时间解析器。   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:1000** 这是一个Go实现的[jieba](https://github.com/fxsjy/jieba)，这是一个中文分词算法。   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![包含中文文档][CN]](https://github.com/yanyiwu/gojieba)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:661** 中文汉字到汉语拼音的转换。   [![最近一周有更新][Green]](https://github.com/mozillazg/go-pinyin)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:474** JP形态学分析仪编写的纯Go。   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:398** Go的自然语言检测包。支持84种语言和24种脚本(书写系统，如拉丁语、西里尔语等)。   [![最近一年没有更新][Yellow]](https://github.com/abadojack/whatlanggo)   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:363** 从字符串中提取值并用nlp填充结构。   [![最近一年没有更新][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:268** 句子标记器:将文本转换为句子列表。   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:246** 支持LSA(潜在语义分析)的Go自然语言处理库。   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  软件包和用于处理本地化文本的附带工具。
* [getlang](https://github.com/rylans/getlang) **star:86** 快速自然语言检测包。   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-nlp](https://github.com/nuance/go-nlp) **star:82** 用于处理离散概率分布的实用程序和用于进行NLP工作的其他工具。   [![最近一年没有更新][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [gounidecode](https://github.com/fiam/gounidecode) **star:68** 用于Go的Unicode音译器(也称为unidecode)。   [![最近一年没有更新][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:67** Unicode文本的ASCII音译。   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [textcat](https://github.com/pebbe/textcat) **star:61** Go package支持基于n-gram的文本分类，支持utf-8和原始文本。   [![最近一年没有更新][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:58** 这是一个 Go 实现的[MMSEG](http://technology.chtsai.org/mmseg/)，这是一个中文分词算法。   [![最近一年没有更新][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:56** 波特词干算法的实现。   [![最近一年没有更新][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:56** 快速自动关键字提取算法(RAKE)的Go端口。   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [stemmer](https://github.com/dchest/stemmer) **star:49** 用于Go编程语言的Stemmer包。包括英语和德语词根。   [![最近一年没有更新][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [segment](https://github.com/blevesearch/segment) **star:48** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![最近一年没有更新][Yellow]](https://github.com/blevesearch/segment)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [porter2](https://github.com/zhenjl/porter2) **star:37** 非常快的波特2史坦默。   [![最近一年没有更新][Yellow]](https://github.com/zhenjl/porter2)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [go2vec](https://github.com/danieldk/go2vec) **star:34** 用于word2vec嵌入式的阅读器和实用程序函数。   [![最近一年没有更新][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [petrovich](https://github.com/striker2000/petrovich) **star:29** 彼得罗维奇是一个图书馆，它把俄语名字的词形变化成特定的语法格。   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang实现了Paice/外壳阻塞算法。   [![最近一年没有更新][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [snowball](https://github.com/goodsign/snowball) **star:25** 滚雪球柄端口(cgo包装)为 Go 。提供词干提取功能[Snowball native](http://snowball.tartarus.org/)。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/snowball)
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo绑定到Yandex。Mystem -俄罗斯形态学分析仪。   [![最近一年没有更新][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo绑定用于icu4c C库的检测和转换功能。保证与版本50.1兼容。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:17**  Go 绑定斯诺鲍libstemmer库，包括波特2。   [![最近一年没有更新][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [shamoji](https://github.com/osamingo/shamoji) **star:11** shamoji是用Go编写的word过滤包。   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** 用于libtextcat C库的Cgo绑定。保证与版本2.2兼容。   [![最近一年没有更新][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [porter](https://github.com/a2800276/porter) **star:8** 这是Martin Porter在C语言中实现的Porter词干分析算法的一个相当简单的移植。   [![最近一年没有更新][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:7** 一种基于字典和双字母格朗语言模型的记号赋予器。(现在只支持中文分割)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)
* [go-localize](https://github.com/m1/go-localize) **star:4** 简单易用的i18n(国际化和本地化)引擎——用于翻译语言环境字符串。   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-localize)
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:1** 语言检测API Go客户端。支持批量请求，短短语或单字语言检测。   [![godoc][GoDoc]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)

## 网络

*用于处理各种网络层的库。 (翻译出错了? 试试 [英文版](README_EN.md#networking) 吧~)*

* [fasthttp](https://github.com/valyala/fasthttp) **star:11737** 一个快速HTTP实现，比net/http快10倍。   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![最近一周有更新][Green]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [kcptun](https://github.com/xtaci/kcptun) **star:11611** 基于KCP协议的非常简单和快速udp隧道。   [![star > 2000][Awesome]](https://github.com/xtaci/kcptun)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcptun)
* [dns](https://github.com/miekg/dns) **star:4389** 用于 DNS 的库。   [![star > 2000][Awesome]](https://github.com/miekg/dns)   [![最近一周有更新][Green]](https://github.com/miekg/dns)   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:3853** 在纯Go中实现了QUIC协议。   [![star > 2000][Awesome]](https://github.com/lucas-clemente/quic-go)   [![最近一周有更新][Green]](https://github.com/lucas-clemente/quic-go)   [![godoc][GoDoc]](https://godoc.org/github.com/lucas-clemente/quic-go)
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3535** HTTPLabs 允许你检查 HTTP 请求和伪造响应。   [![star > 2000][Awesome]](https://github.com/gchaincl/httplab)   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/httplab)
* [gopacket](https://github.com/google/gopacket) **star:3300** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   [![最近一周有更新][Green]](https://github.com/google/gopacket)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [webrtc](https://github.com/pions/webrtc) **star:3239** WebRTC API的纯Go实现。   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   [![最近一周有更新][Green]](https://github.com/pions/webrtc)   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2495** 快速可靠的ARQ协议。   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   [![最近一周有更新][Green]](https://github.com/xtaci/kcp-go)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [gobgp](https://github.com/osrg/gobgp) **star:1827** 基于 Go 的 BGP 实现。   [![最近一周有更新][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [gnet](https://github.com/panjf2000/gnet) **star:1763** “gnet”是一个用纯Go编写的高性能、轻量级、非阻塞、事件驱动的网络框架。   [![最近一周有更新][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![包含中文文档][CN]](https://github.com/panjf2000/gnet)
* [ssh](https://github.com/gliderlabs/ssh) **star:1543** 用于构建SSH服务器的高级API(封装密码/ SSH)。   [![godoc][GoDoc]](https://godoc.org/github.com/gliderlabs/ssh)
* [fortio](https://github.com/fortio/fortio) **star:1184** 负载测试库和命令行工具，高级的echo服务器和web UI。允许指定一组每秒查询的负载，并记录延迟直方图和其他有用的统计数据，并将它们作图。支持Tcp、Http、gRPC。   [![最近一周有更新][Green]](https://github.com/fortio/fortio)   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:981** 简单TUN / TAP图书馆。   [![最近一周有更新][Green]](https://github.com/songgao/water)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [go-getter](https://github.com/hashicorp/go-getter) **star:867**  通过URL来下载文件或目录。   [![最近一周有更新][Green]](https://github.com/hashicorp/go-getter)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [sftp](https://github.com/pkg/sftp) **star:843** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   [![最近一周有更新][Green]](https://github.com/pkg/sftp)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/sftp)
* [gev](https://github.com/Allenxuxu/gev) **star:793** gev是一个基于反应器模式的轻量级、快速的非阻塞TCP网络库。   [![godoc][GoDoc]](https://godoc.org/github.com/Allenxuxu/gev)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:774** 用于快速开发云计算和裸机网络功能的框架(原YANFF)。   [![最近一周有更新][Green]](https://github.com/intel-go/nff-go)   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [grab](https://github.com/cavaliercoder/grab) **star:646**  用于管理文件下载。   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [mdns](https://github.com/hashicorp/mdns) **star:624** 简单mDNS(Multicast DNS)客户端/服务器库。   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/mdns)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  Paho Go客户端提供了一个 MQTT 客户端库，用于通过TCP、TLS或WebSockets连接到MQTT代理。
* [ftp](https://github.com/jlaffaye/ftp) **star:622** 实现了[RFC 959](http://tools.ietf.org/html/rfc959)中描述的ftp客户端。   [![最近一周有更新][Green]](https://github.com/jlaffaye/ftp)   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [lhttp](https://github.com/fanux/lhttp) **star:549** 强大的websocket框架，可以让你更容易的构建IM服务器。   [![最近一年没有更新][Yellow]](https://github.com/fanux/lhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   [![包含中文文档][CN]](https://github.com/fanux/lhttp)
* [gosnmp](https://github.com/soniah/gosnmp) **star:498** 用于执行 SNMP 操作的原生 Go 库。   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [cidranger](https://github.com/yl2chen/cidranger) **star:455** 快速得 IP 到 CIDR 查找。   [![最近一周有更新][Green]](https://github.com/yl2chen/cidranger)   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [gotcp](https://github.com/gansidui/gotcp) **star:452** 用于快速编写 tcp 应用程序。   [![最近一年没有更新][Yellow]](https://github.com/gansidui/gotcp)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/gotcp)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:405** 基于UDP组播的跨平台本地对等点发现库。   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [gopcap](https://github.com/akrennmair/gopcap) **star:379**  用 Go 实现了对 libpcap 的封装。   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [raw](https://github.com/mdlayher/raw) **star:345** Package raw支持在设备驱动程序级别读取和写入网络接口的数据。   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [go-stun](https://github.com/ccding/go-stun) **star:344** 实现了 STUN 客户端(RFC 3489和RFC 5389)。   [![最近一年没有更新][Yellow]](https://github.com/ccding/go-stun)   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [stun](https://github.com/go-rtc/stun) **star:342** Go实现的RFC 5389 STUN协议。   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:324**  Go 图书馆建设tcp服务器更快。   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [winrm](https://github.com/masterzen/winrm) **star:242**  Go WinRM客户端远程执行Windows机器上的命令。   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:234** 通过TCP传输协议缓冲区数据。   [![最近一年没有更新][Yellow]](https://github.com/stabbycutyou/buffstreams)   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [arp](https://github.com/mdlayher/arp) **star:216** 实现了arp协议，如RFC 826中所述。   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:196** 实现了对IEEE 802.3以太网II帧和IEEE 802.1Q VLAN标签的编组和解组。   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [utp](https://github.com/anacrolix/utp) **star:152** Go uTP微传输协议的实现。   [![最近一年没有更新][Yellow]](https://github.com/anacrolix/utp)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/utp)
* [jazigo](https://github.com/udhos/jazigo) **star:144** Jazigo是一个用Go编写的工具，用于检索多个网络设备的配置。   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [canopus](https://github.com/zubairhamed/canopus) **star:138** CoAP客户端/服务器实现(RFC 7252)。   [![最近一年没有更新][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:133** Gmqtt是一个灵活、高性能的MQTT代理库，它完全实现了MQTT协议V3.1.1。   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![包含中文文档][CN]](https://github.com/DrmagicE/gmqtt)
* [sslb](https://github.com/eduardonunesp/sslb) **star:124** 它是一个超级简单的负载平衡器，只是一个实现某种性能的小项目。   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gNxI](https://github.com/google/gnxi) **star:122** 一组基于 gNMI 和 gNOI 协议的网络管理工具。   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [gaio](https://github.com/xtaci/gaio) **star:113** 高性能异步io网络在proactor模式下的Golang。   [![最近一周有更新][Green]](https://github.com/xtaci/gaio)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gaio)
* [xtcp](https://github.com/xfxdev/xtcp) **star:97** TCP服务器框架具有同时全双工通信，优雅关机，自定义协议。   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:65** 实现了一个DHCPv6服务器，如RFC 3315所述。   [![最近一年没有更新][Yellow]](https://github.com/mdlayher/dhcp6)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/dhcp6)
* [ether](https://github.com/songgao/ether) **star:65** 一个用于发送和接收以太网帧的跨平台 Go 库。   [![最近一年没有更新][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [linkio](https://github.com/ian-kent/linkio) **star:47** 网络链路速度模拟，主要用于接口的读/写。   [![最近一年没有更新][Yellow]](https://github.com/ian-kent/linkio)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/linkio)
* [packet](https://github.com/aerogo/packet) **star:44** 通过TCP和UDP发送数据包。它可以缓冲消息和热交换连接。   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   [![最近一年没有更新][Yellow]](https://github.com/aybabtme/portproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/aybabtme/portproxy)
* [iplib](https://github.com/c-robinson/iplib) **star:29** 用于处理IP地址的库(net)。借鉴于python 的 [ipaddress](https://docy-doc.org/3/library/ipaddress.html)和ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [graval](https://github.com/koofr/graval) **star:25** FTP服务器框架。   [![最近一年没有更新][Yellow]](https://github.com/koofr/graval)   [![godoc][GoDoc]](https://godoc.org/github.com/koofr/graval)
* [publicip](https://github.com/polera/publicip) **star:18** 包publicip返回面向公共的IPv4地址(internet出口)。   [![最近一年没有更新][Yellow]](https://github.com/polera/publicip)   [![godoc][GoDoc]](https://godoc.org/github.com/polera/publicip)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:18** 用于解码 pcap 文件和分析解剖数据。   [![最近一年没有更新][Yellow]](https://github.com/sunwxg/golibwireshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/golibwireshark)
* [go-powerdns](https://github.com/joeig/go-powerdns) **star:12** Golang的PowerDNS API绑定。   [![godoc][GoDoc]](https://godoc.org/github.com/joeig/go-powerdns)
* [llb](https://github.com/kirillDanshin/llb) **star:10** 一个非常简单、快速的代理服务器后端。可用于快速重定向到预定义域，具有零内存分配和快速响应。   [![最近一年没有更新][Yellow]](https://github.com/kirillDanshin/llb)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/llb)
* [goshark](https://github.com/sunwxg/goshark) **star:9** 用于解码IP包，创建用于分析的数据结构包。   [![最近一年没有更新][Yellow]](https://github.com/sunwxg/goshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/goshark)
* [tspool](https://github.com/two/tspool) **star:6** TCP库使用工作池来提高性能并保护服务器。   [![最近一年没有更新][Yellow]](https://github.com/two/tspool)   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)
* [gosocsvr](https://github.com/rakeki/gosocsvr) **star:5** Socket服务器变得简单。   [![godoc][GoDoc]](https://godoc.org/github.com/rakeki/gosocsvr)
* [httpproxy](https://github.com/wzshiming/httpproxy) **star:2** HTTP代理处理程序和拨号。   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/httpproxy)

### HTTP客户端

*用于发出HTTP请求的库。 (翻译出错了? 试试 [英文版](README_EN.md#http-clients) 吧~)*

* [resty](https://github.com/go-resty/resty) **star:2624** 简单的 HTTP 和 REST 客户端，受到 Ruby rest-client 的启发。   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   [![最近一周有更新][Green]](https://github.com/go-resty/resty)   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [grequests](https://github.com/levigross/grequests) **star:1532** 一个 Go “克隆”的伟大和著名的请求库。   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [heimdall](https://github.com/gojektech/heimdall) **star:1259** 具有重试和hystrix功能的增强http客户机。   [![godoc][GoDoc]](https://godoc.org/github.com/gojektech/heimdall)
* [sling](https://github.com/dghubble/sling) **star:1108** Sling是一个用于创建和发送API请求的Go HTTP客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/sling)
* [gentleman](https://github.com/h2non/gentleman) **star:737** 功能齐全的插件驱动HTTP客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gentleman)
* [pester](https://github.com/sethgrid/pester) **star:354** 使用重试、后退和并发执行HTTP客户机调用。   [![godoc][GoDoc]](https://godoc.org/github.com/sethgrid/pester)
* [rq](https://github.com/ddo/rq) **star:34** golang stdlib HTTP客户端更好的接口。   [![godoc][GoDoc]](https://godoc.org/github.com/ddo/rq)
* [httpretry](https://github.com/ybbus/httpretry) **star:4** 用重试功能丰富默认的go HTTP客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/httpretry)

## OpenGL

*用于在Go中使用OpenGL的库。 (翻译出错了? 试试 [英文版](README_EN.md#opengl) 吧~)*

* [glfw](https://github.com/go-gl/glfw) **star:872** GLFW 3 的 Go 接口实现。   [![最近一周有更新][Green]](https://github.com/go-gl/glfw)
* [gl](https://github.com/go-gl/gl) **star:695** OpenGL 的 Go 接口实现(通过glow生成)。   [![最近一年没有更新][Yellow]](https://github.com/go-gl/gl)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/gl)
* [mathgl](https://github.com/go-gl/mathgl) **star:322** 完全基于 Go 实现的数学软件包，专门用于处理三维数学。借鉴于 GLM。   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/mathgl)
* [goxjs/gl](https://github.com/goxjs/gl) **star:132** 跨平台的OpenGL 接口实现(OS X, Linux, Windows，浏览器，iOS, Android)。   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/gl)
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:61** 跨平台 glfw 库，可用于创建 OpenGL 上下文并接收事件。   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques. (翻译出错了? 试试 [英文版](README_EN.md#orm) 吧~)*

* [GORM](https://github.com/jinzhu/gorm) **star:17532** 一个出色的 ORM 库。主要目标是对开发人员友好。   [![star > 2000][Awesome]](https://github.com/jinzhu/gorm)   [![最近一周有更新][Green]](https://github.com/jinzhu/gorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/gorm)
* [Xorm](https://github.com/go-xorm/xorm) **star:5843** 基于 Go 的简单而强大的ORM。   [![star > 2000][Awesome]](https://github.com/go-xorm/xorm)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xorm/xorm)   [![包含中文文档][CN]](https://github.com/go-xorm/xorm)
* [go-pg](https://github.com/go-pg/pg) **star:3610** 用于 PostgreSQL 的ORM。侧重于 PostgreSQL 的特性和性能。   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [gorp](https://github.com/go-gorp/gorp) **star:3309** 基于 Go 的关系持久性 ORM-ish 库。   [![star > 2000][Awesome]](https://github.com/go-gorp/gorp)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorp/gorp)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2712** ORM 生成器。根据数据库 schema 生成一个功能强大且运行速度快的ORM。   [![star > 2000][Awesome]](https://github.com/volatiletech/sqlboiler)   [![最近一周有更新][Green]](https://github.com/volatiletech/sqlboiler)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/sqlboiler)
* [upper.io/db](https://github.com/upper/db) **star:2100** 对外提供统一的接口用于访问不同的存储介质，例如PostgreSQL, MySQL, SQLite, MSSQL, QL、MongoDB.。   [![star > 2000][Awesome]](https://github.com/upper/db)   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [reform](https://github.com/go-reform/reform) **star:848** 基于非空接口和代码生成的 ORM。   [![最近一周有更新][Green]](https://github.com/go-reform/reform)   [![godoc][GoDoc]](https://godoc.org/github.com/go-reform/reform)
* [pop/soda](https://github.com/gobuffalo/pop) **star:836** 数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/pop)
* [QBS](https://github.com/coocood/qbs) **star:549** Stands for Query By Struct. A Go ORM.   [![最近一年没有更新][Yellow]](https://github.com/coocood/qbs)   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   [![包含中文文档][CN]](https://github.com/coocood/qbs)
* [go-queryset](https://github.com/jirfag/go-queryset) **star:496** 基于 GORM 100% 类型安全的 ORM。可支持 MySQL, PostgreSQL, Sqlite3, SQL Server。   [![godoc][GoDoc]](https://godoc.org/github.com/jirfag/go-queryset)
* [gormt](https://github.com/xxjwxc/gormt) **star:396** Mysql数据库到golang gorm结构。   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gormt)
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:361** 一个灵活而强大的SQL字符串构建器库。   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/go-sqlbuilder)
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  强大的orm框架。支持: pq/mysql/sqlite3。
* [Zoom](https://github.com/albrow/zoom) **star:251** 基于 Redis 的快速数据存储和查询引擎。   [![最近一年没有更新][Yellow]](https://github.com/albrow/zoom)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [grimoire](https://github.com/Fs02/grimoire) **star:126** 基于 golang 的数据库访问层和验证库。(支持: MySQL, PostgreSQL和SQLite3)。   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/grimoire)
* [go-store](https://github.com/gosuri/go-store) **star:99** 简单且快速的 Redis 键值存储库。   [![最近一年没有更新][Yellow]](https://github.com/gosuri/go-store)   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/go-store)
* [Marlow](https://github.com/dadleyy/marlow) **star:76** 从项目结构生成ORM。   [![godoc][GoDoc]](https://godoc.org/github.com/dadleyy/marlow)
* [rel](https://github.com/Fs02/rel) **star:61** 用于干净(洋葱)架构的Golang SQL存储库层。   [![最近一周有更新][Green]](https://github.com/Fs02/rel)   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/rel)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:12** 一个轻量级的ORM。用于Google/Firebase Cloud Firestore。   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)
* [lore](https://github.com/abrahambotros/lore) **star:5** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![最近一年没有更新][Yellow]](https://github.com/abrahambotros/lore)   [![godoc][GoDoc]](https://godoc.org/github.com/abrahambotros/lore)

## 包管理

*用于管理依赖和包的官方工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules 是源码的版本控制和交换的单位。go命令直接支持处理模块，包括记录和解决对其他模块的依赖关系。

*包管理的官方实验工具 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [dep](https://github.com/golang/dep) **star:13119**  Go 的依赖管理工具，需要 Go 1.9+   [![star > 2000][Awesome]](https://github.com/golang/dep)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/dep)
* [vgo](https://go.googlesource.com/vgo/)  Go 命令版本管理

*用于包和依赖项管理的非官方库。 (翻译出错了? 试试 [英文版](README_EN.md#package-management) 吧~)*

* [glide](https://github.com/Masterminds/glide) **star:7962** 轻松管理您的 golang 第三方包。受Maven、Bundler和Pip等工具的启发。   [![star > 2000][Awesome]](https://github.com/Masterminds/glide)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/glide)
* [godep](https://github.com/tools/godep) **star:5646** godep是go的依赖工具，它通过修复包的依赖关系来帮助构建可重复的包。   [![star > 2000][Awesome]](https://github.com/tools/godep)   [![最近一年没有更新][Yellow]](https://github.com/tools/godep)   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   [![归档项目][Archived]](https://github.com/tools/godep)
* [govendor](https://github.com/kardianos/govendor) **star:5013** 包管理器。使用 vendor 文件的 Go vendor 工具。   [![star > 2000][Awesome]](https://github.com/kardianos/govendor)   [![godoc][GoDoc]](https://godoc.org/github.com/kardianos/govendor)   [![归档项目][Archived]](https://github.com/kardianos/govendor)
* [gopm](https://github.com/gpmgo/gopm) **star:2471** 包管理器。   [![star > 2000][Awesome]](https://github.com/gpmgo/gopm)   [![godoc][GoDoc]](https://godoc.org/github.com/gpmgo/gopm)   [![归档项目][Archived]](https://github.com/gpmgo/gopm)
* [gom](https://github.com/mattn/gom) **star:1380** Go Manager - bundle for Go。   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [gpm](https://github.com/pote/gpm) **star:1203** 基本的 Go 依赖管理器。   [![最近一年没有更新][Yellow]](https://github.com/pote/gpm)
* [goop](https://github.com/nitrous-io/goop) **star:780** Go 的简单依赖管理器，灵感来自Bundler。   [![最近一年没有更新][Yellow]](https://github.com/nitrous-io/goop)   [![godoc][GoDoc]](https://godoc.org/github.com/nitrous-io/goop)
* [nut](https://github.com/jingweno/nut) **star:244** vendor 依赖。   [![最近一年没有更新][Yellow]](https://github.com/jingweno/nut)   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** 使用Git的最小依赖版本。   [![最近一年没有更新][Yellow]](https://github.com/VividCortex/johnny-deps)
* [gigo](https://github.com/LyricalSecurity/gigo) **star:198** 类似pip的golang依赖工具，支持私有存储库和散列。   [![最近一年没有更新][Yellow]](https://github.com/LyricalSecurity/gigo)   [![godoc][GoDoc]](https://godoc.org/github.com/LyricalSecurity/gigo)
* [VenGO](https://github.com/DamnWidget/VenGO) **star:117** 创建和管理可导出的隔离go虚拟环境。   [![最近一年没有更新][Yellow]](https://github.com/DamnWidget/VenGO)   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:98** 插件，为自动加载Golang SDK，依赖关系管理和启动Maven项目基础设施中的构建环境提供了方法。
* [gop](https://github.com/lunny/gop) **star:49** 在GOPATH之外构建和管理Go应用程序。   [![最近一年没有更新][Yellow]](https://github.com/lunny/gop)   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/gop)   [![包含中文文档][CN]](https://github.com/lunny/gop)   [![归档项目][Archived]](https://github.com/lunny/gop)

## 性能

* [jaeger](https://github.com/jaegertracing/jaeger) **star:10427** 分布式跟踪系统。   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   [![最近一周有更新][Green]](https://github.com/jaegertracing/jaeger)   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1203** Go的简单分析支持包。   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [tracer](https://github.com/kamilsk/tracer) **star:23** 简单、轻量级的跟踪。   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## 查询语言

* [graphql-go](https://github.com/graphql-go/graphql) **star:6121** 为Go实现GraphQL。   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [graphql](https://github.com/neelance/graphql-go) **star:3216** 关注易用性的GraphQL服务器。   [![star > 2000][Awesome]](https://github.com/neelance/graphql-go)   [![godoc][GoDoc]](https://godoc.org/github.com/neelance/graphql-go)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:1324** 一个用来查询JSON数据的Go包。   [![最近一周有更新][Green]](https://github.com/thedevsaddam/gojsonq)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [jsonql](https://github.com/elgs/jsonql) **star:220** Golang中的JSON查询表达式库。   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/jsonql)
* [rql](https://github.com/a8m/rql) **star:136** 用于REST API的资源查询语言。   [![godoc][GoDoc]](https://godoc.org/github.com/a8m/rql)
* [graphql](https://github.com/tmc/graphql) **star:52** graphql解析器+工具集   [![最近一年没有更新][Yellow]](https://github.com/tmc/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [jsonslice](https://github.com/bhmj/jsonslice) **star:39** 使用高级过滤器查询Jsonpath。   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)
* [straf](https://github.com/SonicRoshan/straf) **star:10** 轻松地将Golang结构转换为GraphQL对象。   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/straf)

## 嵌入的资源

* [packr](https://github.com/gobuffalo/packr) **star:2666** 将静态文件嵌入到Go二进制文件中的简单方法。   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![最近一周有更新][Green]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik) **star:2560** 将静态文件嵌入到Go可执行文件中。   [![star > 2000][Awesome]](https://github.com/rakyll/statik)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/statik)
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1869** go.rice 是一个Go包，它使处理html、js、css、图像和模板等资源变得非常容易。   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:813** 生成一个vfsdata。静态实现给定虚拟文件系统的go文件。   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/vfsgen)
* [esc](https://github.com/mjibson/esc) **star:534** 将文件嵌入到Go程序中并提供http文件系统接口。   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:502** 一个可定制的工具用来在Go中嵌入文件   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-resources](https://github.com/omeid/go-resources) **star:165** 嵌入到Go中的普通资源。   [![godoc][GoDoc]](https://godoc.org/github.com/omeid/go-resources)
* [statics](https://github.com/go-playground/statics) **star:54** 将静态资源嵌入到go文件中，用于单个二进制编译+使用http。文件系统+符号链接。   [![最近一年没有更新][Yellow]](https://github.com/go-playground/statics)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/statics)
* [templify](https://github.com/wlbr/templify) **star:25** 将外部模板文件嵌入到Go代码中，以创建单个文件二进制文件。   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/templify)
* [go-embed](https://github.com/pyros2097/go-embed) **star:21** 生成go代码，将资源文件嵌入到库或可执行文件中。   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)

## 科学与数据分析

*用于科学计算和数据分析的库。 (翻译出错了? 试试 [英文版](README_EN.md#science-and-data-analysis) 吧~)*

* [gonum](https://github.com/gonum/gonum) **star:3626** Gonum是一组用于Go编程语言的数字库。它包含用于矩阵、统计、优化等的库。   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   [![最近一周有更新][Green]](https://github.com/gonum/gonum)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [stats](https://github.com/montanaflynn/stats) **star:1647** 包含Golang标准库中缺少的公共函数的统计软件包。   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gonum/plot](https://github.com/gonum/plot) **star:1494** gonum/plot提供了一个API，用于在Go中构建和绘制绘图。   [![最近一周有更新][Green]](https://github.com/gonum/plot)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [gosl](https://github.com/cpmech/gosl) **star:1401** 提供线性代数，FFT，几何，NURBS，数值方法，概率，优化，微分方程，等等。   [![godoc][GoDoc]](https://godoc.org/github.com/cpmech/gosl)
* [streamtools](https://github.com/nytlabs/streamtools) **star:1317** 通用图形工具，用于处理数据流。   [![最近一年没有更新][Yellow]](https://github.com/nytlabs/streamtools)   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [go-dsp](https://github.com/mjibson/go-dsp) **star:667** Go数字信号处理。   [![最近一年没有更新][Yellow]](https://github.com/mjibson/go-dsp)   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/go-dsp)
* [chart](https://github.com/vdobler/chart) **star:627** 简单的图表绘制库。支持多种图形类型。   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [goraph](https://github.com/gyuho/goraph) **star:619** 纯Go图论库(数据结构，算法可视化)。   [![最近一年没有更新][Yellow]](https://github.com/gyuho/goraph)   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [graph](https://github.com/yourbasic/graph) **star:308** 基本图形算法库。   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [ewma](https://github.com/VividCortex/ewma) **star:291** 提供指数加权移动平均算法。   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [orb](https://github.com/paulmach/orb) **star:268** 2D几何类型，支持剪切、GeoJSON和Mapbox矢量平铺。   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:168** 用于机器学习和统计的数据框架(类似于panda)。   [![最近一周有更新][Green]](https://github.com/rocketlaunchr/dataframe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:140** 数据流的近似直方图。   [![最近一年没有更新][Yellow]](https://github.com/VividCortex/gohistogram)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:96** TextRank在Golang中的实现，支持扩展特性(摘要、加权、短语提取)和多线程(goroutine)。   [![最近一年没有更新][Yellow]](https://github.com/DavidBelicza/TextRank)   [![godoc][GoDoc]](https://godoc.org/github.com/DavidBelicza/TextRank)
* [sparse](https://github.com/james-bowman/sparse) **star:82**  Go 稀疏矩阵格式的线性代数支持科学和机器学习应用程序，兼容gonum矩阵库。   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [pagerank](https://github.com/alixaxel/pagerank) **star:55** 加权 PageRank 算法在Go中的实现。   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [geom](https://github.com/skelterjohn/geom) **star:44**  Go 的二维几何。   [![最近一年没有更新][Yellow]](https://github.com/skelterjohn/geom)   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [evaler](https://github.com/soniah/evaler) **star:39** 简单的浮点算术表达式求值器。   [![最近一年没有更新][Yellow]](https://github.com/soniah/evaler)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:17**  Go 实现熵度量。   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:13** 二维三角库。允许将线和多边形(都基于点)转换为gpu语言。   [![最近一周有更新][Green]](https://github.com/tchayen/triangolatte)   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:11** 微型线性插值库。   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/piecewiselinear)
* [ode](https://github.com/ChristopherRabotin/ode) **star:11** 常微分方程(ODE)求解器，支持扩展状态和基于信道的迭代停止条件。   [![最近一年没有更新][Yellow]](https://github.com/ChristopherRabotin/ode)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/ode)
* [GoStats](https://github.com/OGFris/GoStats) **star:10** GoStats是一个开放源码的GoLang库，主要用于机器学习领域的数学统计，它涵盖了大多数统计度量函数。   [![最近一年没有更新][Yellow]](https://github.com/OGFris/GoStats)   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [PiHex](https://github.com/claygod/PiHex) **star:8** 实现了针对16进制数 Pi 的“bailee - borwein - plouffe”算法。   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [assocentity](https://github.com/ndabAP/assocentity) **star:6** assocentity 返回单词到给定实体的平均距离。   [![最近一周有更新][Green]](https://github.com/ndabAP/assocentity)   [![godoc][GoDoc]](https://godoc.org/github.com/ndabAP/assocentity)
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** 用“Go”语言编写的图论算法。   [![最近一年没有更新][Yellow]](https://github.com/ThePaw/go-gt)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/go-gt)
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** 二次函数求根算法库。   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:2** 为成对比较提供了一个 Bradley-Terry 模型。   [![godoc][GoDoc]](https://godoc.org/github.com/seanhagen/bradleyterry)

## 安全

*用于帮助您的应用程序更安全的库。 (翻译出错了? 试试 [英文版](README_EN.md#security) 吧~)*

* [lego](https://github.com/xenolf/lego) **star:3903** 纯 Go ACME 客户端库及命令行工具   [![star > 2000][Awesome]](https://github.com/xenolf/lego)   [![最近一周有更新][Green]](https://github.com/xenolf/lego)   [![godoc][GoDoc]](https://godoc.org/github.com/xenolf/lego)
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:2154** 工具和库，以远程入侵RTSP流从监控摄像头。   [![star > 2000][Awesome]](https://github.com/Ullaakut/cameradar)   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [acmetool](https://github.com/hlandau/acme) **star:1747** ACME(让我们用自动更新加密)客户端工具。   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/acme)
* [memguard](https://github.com/awnumar/memguard) **star:1727** 一个用于处理内存中敏感值的纯Go库。   [![最近一周有更新][Green]](https://github.com/awnumar/memguard)   [![godoc][GoDoc]](https://godoc.org/github.com/awnumar/memguard)
* [secure](https://github.com/unrolled/secure) **star:1430** Go 语言 HTTP 中间件，为 Go 提供了一些安全功能   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/secure)
* [acra](https://github.com/cossacklabs/acra) **star:547** 网络加密代理保护基于数据库的应用程序免受数据泄漏:强选择性加密，SQL注入预防，入侵检测系统。   [![最近一周有更新][Green]](https://github.com/cossacklabs/acra)   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [nacl](https://github.com/kevinburke/nacl) **star:476**  Go 实现NaCL API的集合。   [![godoc][GoDoc]](https://godoc.org/github.com/kevinburke/nacl)
* [BadActor](https://github.com/jaredfolkins/badactor) **star:265** 一个驻留在内存中的，应用驱动的监控程序，受 fail2ban 的启发   [![godoc][GoDoc]](https://godoc.org/github.com/jaredfolkins/badactor)
* [passlib](https://github.com/hlandau/passlib) **star:233** 不过时的密码哈希库。   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/passlib)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:231** 使用ssh密钥加密/解密。   [![godoc][GoDoc]](https://godoc.org/github.com/ssh-vault/ssh-vault)
* [optimus-go](https://github.com/pjebs/optimus-go) **star:193** 使用Knuth的算法进行ID哈希和混淆。   [![godoc][GoDoc]](https://godoc.org/github.com/pjebs/optimus-go)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:163** Scrypt 库，具有简单、易懂的 API，同时具有内置的自动校准功能   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [go-yara](https://github.com/hillu/go-yara) **star:153** YARA的 Go 语言接口，号称是 “对于恶意软件研究者（以及其他人）来说是模式匹配的瑞士军刀”   [![最近一周有更新][Green]](https://github.com/hillu/go-yara)   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [argon2pw](https://github.com/raja/argon2pw) **star:81** 使用常量时间密码比较生成Argon2密码散列。   [![最近一年没有更新][Yellow]](https://github.com/raja/argon2pw)   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  让我们加密证书并启动TLS服务器。
* [Interpol](https://bitbucket.org/vahidi/interpol)  基于规则的数据生成器，用于模糊和渗透测试。
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:35** 一个安全哈希和加密密码的偏执包。   [![最近一年没有更新][Yellow]](https://github.com/dwin/goSecretBoxPassword)   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
* [certificates](https://github.com/mvmaasakkers/certificates) **star:13** 用于生成tls证书的自定义工具。   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2密码散列和验证设计为与现有Python和PHP实现兼容。   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:8** 使用围绕acme/autocert的高级包装器，SSL证书变得很容易。   [![godoc][GoDoc]](https://godoc.org/github.com/adrianosela/sslmgr)
* [go-generate-password](https://github.com/m1/go-generate-password) **star:7** 可以在cli上或作为库使用的密码生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-generate-password)
* [secureio](https://github.com/xaionaro-go/secureio) **star:5** 一个密钥交换+认证+加密的包装和多路复用的' io。读写ecloser '基于XChaCha20-poly1305, ECDH和ED25519。   [![最近一周有更新][Green]](https://github.com/xaionaro-go/secureio)   [![godoc][GoDoc]](https://godoc.org/github.com/xaionaro-go/secureio)

## 序列化

*用于二进制序列化的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#serialization) 吧~)*

* [jsoniter](https://github.com/json-iterator/go) **star:7296** 高性能，100% 兼容的“encoding/json” 替代品   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:6325** 通过库和协议编译器插件使 Go 语言支持 Google的 protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   [![最近一周有更新][Green]](https://github.com/golang/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:3563** Go 语言的 Protocol Buffer 库。   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:3131** 用于对原生键值对进行解码生成 Go 语言结构体   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1375** 高性能、多功能、规范化编码解码以及 rpc 库， 用于 msgpack, cbor 和 json，支持基于运行时的 OR 码生成   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [colfer](https://github.com/pascaldekloe/colfer) **star:518** 为Colfer二进制格式生成代码。
* [csvutil](https://github.com/jszwec/csvutil) **star:354** 高性能、惯用的CSV记录编码和解码到本机Go结构。   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:277** Go 语言用的 Cap'n Proto 库及解析器   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:135** 用于协同 PHP session 格式数据和 PHP 序列化／反序列化函数工作的go语言库   [![最近一年没有更新][Yellow]](https://github.com/yvasiyarov/php_session_decoder)   [![godoc][GoDoc]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
* [cbor](https://github.com/fxamacker/cbor) **star:117** 小，安全，容易的CBOR编码和解码库。   [![最近一周有更新][Green]](https://github.com/fxamacker/cbor)   [![godoc][GoDoc]](https://godoc.org/github.com/fxamacker/cbor)
* [structomap](https://github.com/tuvistavie/structomap) **star:110** 用于从静态结构体简单、动态的生成键值对的库   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [bambam](https://github.com/glycerine/bambam) **star:60** 用于 Go 语言生成 Cap'n Proto schemas 的生成器   [![最近一年没有更新][Yellow]](https://github.com/glycerine/bambam)   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/bambam)
* [asn1](https://github.com/PromonLogicalis/asn1) **star:44** 面向golang的BER和DER编码库。   [![最近一年没有更新][Yellow]](https://github.com/PromonLogicalis/asn1)   [![godoc][GoDoc]](https://godoc.org/github.com/PromonLogicalis/asn1)   [![归档项目][Archived]](https://github.com/PromonLogicalis/asn1)
* [binstruct](https://github.com/ghostiam/binstruct) **star:17** 用于将数据映射到结构中的Golang二进制解码器。   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [fwencoder](https://github.com/o1egl/fwencoder) **star:10** 用于Go的固定宽度文件解析器(编码和解码库)。   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/fwencoder)
* [bel](https://github.com/32leaves/bel) **star:8** 从Go structs/interface生成TypeScript接口。对JSON RPC很有用。   [![godoc][GoDoc]](https://godoc.org/github.com/32leaves/bel)
* [pletter](https://github.com/vimeda/pletter) **star:7** A standard way to wrap a proto message for message brokers.   [![最近一周有更新][Green]](https://github.com/vimeda/pletter)   [![godoc][GoDoc]](https://godoc.org/github.com/vimeda/pletter)
* [elastic](https://github.com/epiclabs-io/elastic) **star:5** 在运行时跨不同类型转换片、映射或任何其他未知值，无论如何。   [![godoc][GoDoc]](https://godoc.org/github.com/epiclabs-io/elastic)
* [fixedwidth](https://github.com/huydang284/fixedwidth) **star:4** 固定宽度的文本格式(支持UTF-8)。   [![godoc][GoDoc]](https://godoc.org/github.com/huydang284/fixedwidth)

## 服务器应用程序

* [etcd](https://github.com/coreos/etcd) **star:30049** 为共享配置和服务发现提供高可用的键值存储。   [![star > 2000][Awesome]](https://github.com/coreos/etcd)   [![最近一周有更新][Green]](https://github.com/coreos/etcd)   [![godoc][GoDoc]](https://godoc.org/github.com/coreos/etcd)
* [Caddy](https://github.com/mholt/caddy) **star:26762** Caddy是另一种HTTP/2 web服务器，易于配置和使用。   [![star > 2000][Awesome]](https://github.com/mholt/caddy)   [![最近一周有更新][Green]](https://github.com/mholt/caddy)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/caddy)
* [consul](https://www.consul.io/)  Consul 是一个用于服务发现、监控和配置的工具
* [minio](https://github.com/minio/minio) **star:20751** Minio是一个分布式对象存储服务器。   [![star > 2000][Awesome]](https://github.com/minio/minio)   [![最近一周有更新][Green]](https://github.com/minio/minio)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [RoadRunner](https://github.com/spiral/roadrunner) **star:4098** 高性能PHP应用服务器，负载平衡器和进程管理器。   [![star > 2000][Awesome]](https://github.com/spiral/roadrunner)   [![最近一周有更新][Green]](https://github.com/spiral/roadrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/spiral/roadrunner)
* [devd](https://github.com/cortesi/devd) **star:2955** 为开发人员提供本地web服务器。   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [algernon](https://github.com/xyproto/algernon) **star:1675** 内置支持Lua、Markdown、GCSS和Amber的HTTP/2 web服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/algernon)
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:1487** 功能齐全，高度可配置的SFTP服务器软件。   [![最近一周有更新][Green]](https://github.com/drakkan/sftpgo)   [![godoc][GoDoc]](https://godoc.org/github.com/drakkan/sftpgo)
* [flipt](https://github.com/markphelps/flipt) **star:1173** 一个用Go和Vue.js编写的自包含特性标志解决方案   [![最近一周有更新][Green]](https://github.com/markphelps/flipt)   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Trickster](https://github.com/Comcast/trickster) **star:1060** HTTP反向代理缓存和时间序列加速器。   [![最近一周有更新][Green]](https://github.com/Comcast/trickster)   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/trickster)
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  小型化、网络化、基于内存的键值存储
* [Fider](https://github.com/getfider/fider) **star:1056** Fider是一个收集和组织客户反馈的开放平台。   [![最近一周有更新][Green]](https://github.com/getfider/fider)   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [Flagr](https://github.com/checkr/flagr) **star:1025** Flagr是一个开源特性标记和A/B测试服务。   [![最近一周有更新][Green]](https://github.com/checkr/flagr)   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [discovery](https://github.com/Bilibili/discovery) **star:872** 用于弹性中间层负载平衡和故障转移的注册表。   [![最近一周有更新][Green]](https://github.com/Bilibili/discovery)   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [jackal](https://github.com/ortuman/jackal) **star:797** 用Go编写的XMPP服务器。   [![最近一周有更新][Green]](https://github.com/ortuman/jackal)   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [dudeldu](https://github.com/krotik/dudeldu) **star:109** 一个简单的SHOUTcast服务器。   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/dudeldu)
* [lets-proxy2](https://github.com/rekby/lets-proxy2) **star:26** 反向代理处理从let -encrypt中发出证书的https。   [![最近一周有更新][Green]](https://github.com/rekby/lets-proxy2)   [![godoc][GoDoc]](https://godoc.org/github.com/rekby/lets-proxy2)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:15** 从PostgreSQL到Kafka的流数据库事件。   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  传递到负载平衡Riemann事件并/或将其转换为 Carbon。
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:9** Nginx日志解析器和Prometheus 导出。   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  一个实时分布式消息平台。


## 流处理

*用于流处理和响应式编程的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#stream-processing) 吧~)*

* [go-streams](https://github.com/reugn/go-streams) **star:287** 流处理库。   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)

## 模板引擎

*用于模板和词法分析的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#template-engines) 吧~)*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3495** PDF 文档生成器，支持文本，绘图和图片   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)   [![归档项目][Archived]](https://github.com/jung-kurt/gofpdf)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1654** 快速、强大且易用的模板引擎。将模板转化为 Go 语言并进行编译   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [pongo2](https://github.com/flosch/pongo2) **star:1633** 类似 DjanGo 的模板引擎   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [hero](https://github.com/shiyanhui/hero) **star:1292** Hero是一个方便、快速和强大的go模板引擎。   [![godoc][GoDoc]](https://godoc.org/github.com/shiyanhui/hero)   [![包含中文文档][CN]](https://github.com/shiyanhui/hero)
* [mustache](https://github.com/hoisie/mustache) **star:980** Go 语言实现的 Mustache 模板语言   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:840** Amber是一个优雅的Go编程语言模板引擎，它的灵感来自HAML和Jade。   [![最近一年没有更新][Yellow]](https://github.com/eknkc/amber)   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:785** Ace 是一个 Go 语言的 HTML 模板引擎，受到了 Slim 和 Jade 的启发。 Ace 是对Gold的一种改进。   [![最近一年没有更新][Yellow]](https://github.com/yosssi/ace)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:733** Go 语言的 Razor 视图引擎   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [jet](https://github.com/CloudyKit/jet) **star:634** Jet模板引擎。   [![最近一周有更新][Green]](https://github.com/CloudyKit/jet)   [![godoc][GoDoc]](https://godoc.org/github.com/CloudyKit/jet)
* [ego](https://github.com/benbjohnson/ego) **star:429** 轻量级模板语言，允许您在Go中编写模板。模板被翻译成Go并编译。   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:368** 简单快速的模板引擎。进行模板元素替换时，速度是比[text/template](http://golang.org/pkg/text/template/)快10倍。   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasttemplate)
* [raymond](https://github.com/aymerick/raymond) **star:364** 使用 Go 语言实现的完整的 handlebars   [![最近一年没有更新][Yellow]](https://github.com/aymerick/raymond)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [Soy](https://github.com/robfig/soy) **star:148** Go 语言实现的谷歌闭包模板(也就是 Soy templates) ,遵循[官方规范](https://developer.google.com/closure/templates/)。   [![最近一年没有更新][Yellow]](https://github.com/robfig/soy)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [maroto](https://github.com/johnfercher/maroto) **star:122** 创建pdf文件的maroto方法。Maroto的灵感来自于Bootstrap并使用gofpdf。快速和简单。   [![godoc][GoDoc]](https://godoc.org/github.com/johnfercher/maroto)
* [goview](https://github.com/foolin/goview) **star:109** Goview是一个轻量级、极简的模板库，基于golang html/template构建Go web应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/goview)
* [liquid](https://github.com/osteele/liquid) **star:101** Go 语言实现的 Shopify Liquid 模板.   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/liquid)
* [kasia.go](https://github.com/ziutek/kasia.go) **star:71** 一个用于HTML 和其他文本文件的模板系统，使用go语言实现   [![最近一年没有更新][Yellow]](https://github.com/ziutek/kasia.go)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [velvet](https://github.com/gobuffalo/velvet) **star:71** 使用 Go 语言实现的完整的 handlebars   [![最近一年没有更新][Yellow]](https://github.com/gobuffalo/velvet)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)
* [damsel](https://github.com/dskinner/damsel) **star:22** 标记语言，通过css选择器实现了 html 框架 ，并可以通过 pkg html/template 等进行扩展   [![最近一年没有更新][Yellow]](https://github.com/dskinner/damsel)   [![godoc][GoDoc]](https://godoc.org/github.com/dskinner/damsel)
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:19**  对 html/template 进行了简单的封装，支持基于文件的模板可以利用其他模板文件进行扩展   [![最近一年没有更新][Yellow]](https://github.com/dannyvankooten/extemplate)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/extemplate)
* [gospin](https://github.com/m1/gospin) **star:8** 文章纺纱和spintax/纺纱语法引擎，适用于A/B，测试文本/文章片段和创建更多的自然对话。   [![godoc][GoDoc]](https://godoc.org/github.com/m1/gospin)

## 测试

*用于测试代码库和生成测试数据的库。 (翻译出错了? 试试 [英文版](README_EN.md#testing) 吧~)*

* Testing Frameworks
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  将markdown代码段转换为可测试的go代码。
    * [Testify](https://github.com/stretchr/testify) **star:9884** 对标准测试包的扩展。   [![star > 2000][Awesome]](https://github.com/stretchr/testify)   [![最近一周有更新][Green]](https://github.com/stretchr/testify)   [![godoc][GoDoc]](https://godoc.org/github.com/stretchr/testify)
    * [go-cmp](https://github.com/google/go-cmp) **star:1557** 用于比较测试中的Go值的包。   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1302** 简洁的、声明式的、易用的端到端HTTP 及 REST API 测试   [![godoc][GoDoc]](https://godoc.org/github.com/gavv/httpexpect)
    * [godog](https://github.com/DATA-DOG/godog) **star:961** 类似 Cucumber 或 Behat 的 BDD 框架   [![最近一周有更新][Green]](https://github.com/DATA-DOG/godog)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/godog)
    * [baloo](https://github.com/h2non/baloo) **star:675** 表达性强、多功能的、端到端的HTTP API 测试工具   [![最近一年没有更新][Yellow]](https://github.com/h2non/baloo)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD 风格的测试框架，具有 web 界面和计时刷新功能
    * [gocheck](http://labix.org/gocheck)  更加高级的测试框架，用于替换 Gotest
    * [goblin](https://github.com/franela/goblin) **star:663** 类似Mocha的测试框架。   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:454** 类似 Rails 的测试工具，用于测试数据库应用   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:394** 记录并回放HTTP交互，以便进行快速、确定和准确的测试。   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:330** 变异测试的Go源代码。   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gofight](https://github.com/appleboy/gofight) **star:303** 对 Go 语言的路由框架进行 API 测试   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [frisby](https://github.com/verdverm/frisby) **star:255** REST API测试框架。   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  Go的BDD测试框架。
    * [go-carpet](https://github.com/msoap/go-carpet) **star:202** 在终端中查看测试覆盖率的工具。   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [charlatan](https://github.com/percolate/charlatan) **star:192** 为测试生成假接口实现的工具。   [![godoc][GoDoc]](https://godoc.org/github.com/percolate/charlatan)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:148** 一组包，用于增强go测试包并支持公共模式。   [![最近一周有更新][Green]](https://github.com/gotestyourself/gotest.tools)   [![godoc][GoDoc]](https://godoc.org/github.com/gotestyourself/gotest.tools)
    * [endly](https://github.com/viant/endly) **star:133** 声明性端到端功能测试。   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [commander](https://github.com/SimonBaeumer/commander) **star:119** 用于在windows、linux和osx上测试cli应用程序的工具。   [![最近一周有更新][Green]](https://github.com/SimonBaeumer/commander)   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:114** 用于 Go 编程语言的bdd风格的测试框架。   [![最近一年没有更新][Yellow]](https://github.com/orfjackal/gospec)   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:103** 测试框架的简单快照测试插件。   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:102** 清空数据库用于测试，受到database_cleaner 的启发   [![最近一周有更新][Green]](https://github.com/khaiql/dbcleaner)   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:84** 极具灵活性的golang深度比较，扩展了go测试包。   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [wstest](https://github.com/posener/wstest) **star:73** 用于单元测试Websocket http.Handler的Websocket客户机。   [![最近一周有更新][Green]](https://github.com/posener/wstest)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [apitest](https://apitest.dev)  基于REST的服务/HTTP处理程序的简单且可扩展的行为测试库,支持模拟外部HTTP调用和呈现序列图
    * [gospecify](https://github.com/stesla/gospecify) **star:53** 支持 BDD 语法 。对于任何使用过 rspec 等库的人来说应该非常熟悉。   [![最近一年没有更新][Yellow]](https://github.com/stesla/gospecify)   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [restit](https://github.com/yookoala/restit) **star:50** 帮助编写 RESTful API 集成测试的 Go 语言微型框架.。   [![godoc][GoDoc]](https://godoc.org/github.com/yookoala/restit)
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:34** 用于验证JSON有效负载已正确序列化的包。   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [gomega](http://onsi.github.io/gomega/)  类似 Rspec 的 matcher/assertion 库
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:32** 为针对模式测试JSON而创建的库。   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [dsunit](https://github.com/viant/dsunit) **star:29** 用于SQL、NoSQL、结构化文件的数据存储测试。   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:27** 用于声明性 Matcher 对象的连贯框架，当将其应用于输入值时，将产生自描述结果。   [![最近一年没有更新][Yellow]](https://github.com/rdrdr/hamcrest)   [![godoc][GoDoc]](https://godoc.org/github.com/rdrdr/hamcrest)
    * [testcase](https://github.com/adamluzsi/testcase) **star:17** 行为驱动开发的惯用测试框架。   [![最近一周有更新][Green]](https://github.com/adamluzsi/testcase)   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [assert](https://github.com/go-playground/assert) **star:17** 基础断言库，用于对 Go 语言程序进行测试，提供了一些用于自定义断言的代码块   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/assert)
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) **star:12** 作为另一个Go应用程序或测试的一部分，在Linux、OSX或Windows上本地运行一个真正的Postgres数据库。   [![godoc][GoDoc]](https://godoc.org/github.com/fergusstrange/embedded-postgres)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:11** HTTP客户端测试框架。   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Go 语言 testing/iotest 包的扩展。   [![最近一年没有更新][Yellow]](https://github.com/cavaliercoder/badio)   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/badio)
    * [gocrest](https://github.com/corbym/gocrest) **star:9** 用于 Go 断言的可组合的类似 hamcrest 的 matchers。   [![最近一年没有更新][Yellow]](https://github.com/corbym/gocrest)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gocrest)
    * [schema](https://github.com/jgroeneveld/schema) **star:9** 快速、简单地匹配请求和响应中使用的JSON模式的表达式。   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/schema)
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** 轻量级测试套，为 Go1.7's Subtests 带来了setup/teardown 功能   [![最近一年没有更新][Yellow]](https://github.com/pavlo/gosuite)   [![godoc][GoDoc]](https://godoc.org/github.com/pavlo/gosuite)
    * [biff](https://github.com/fulldump/biff) **star:8** 分支测试框架，BDD兼容。   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/biff)
    * [gogiven](https://github.com/corbym/gogiven) **star:8** 类似于 YATSPEC 的Go BDD测试框架。   [![最近一年没有更新][Yellow]](https://github.com/corbym/gogiven)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gogiven)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** 在测试前从SQL文件生成测试数据，并在测试完成后清除数据。   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [Tt](https://github.com/vcaesar/tt) **star:6** 简单而丰富多彩的测试工具。   [![godoc][GoDoc]](https://godoc.org/github.com/vcaesar/tt)
    * [trial](https://github.com/jgroeneveld/trial) **star:4** 无需引入太多样板，即可快速轻松地扩展断言。   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/trial)

* Mock
    * [gomock](https://github.com/golang/mock) **star:3874** 用于Go编程语言的mock框架。   [![star > 2000][Awesome]](https://github.com/golang/mock)   [![最近一周有更新][Green]](https://github.com/golang/mock)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/mock)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:2315** Mock SQL ，用于测试数据库交互   [![star > 2000][Awesome]](https://github.com/DATA-DOG/go-sqlmock)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1531** 使用可扩展中间件和易于使用的CLI记录和模拟REST/SOAP api的HTTP(S)代理。   [![最近一周有更新][Green]](https://github.com/SpectoLabs/hoverfly)   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:962** 多功能、易用 HTTP mock   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [httpmock](https://github.com/jarcoal/httpmock) **star:719** 轻松模拟来自外部资源的HTTP响应。   [![最近一周有更新][Green]](https://github.com/jarcoal/httpmock)   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:402** 用于生成自包含 mock 对象的工具   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [minimock](https://github.com/gojuno/minimock) **star:294** Go接口的模拟生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:238** 基于单事务的数据库驱动，主要用于测试目的   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:88** HTTP mock : 离线测试时记录和重放浏览器的动作   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)
    * [timex](https://github.com/cabify/timex) **star:26** 一个测试友好的本地“时间”包的替代品。   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/timex)
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Go http.ResponseWriter的模拟对象。   [![最近一年没有更新][Yellow]](https://github.com/tv42/mockhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/tv42/mockhttp)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:3340** 随机测试系统。   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![最近一周有更新][Green]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz) **star:739** 用于生成随机值来初始化 Go 语言对象的库   [![最近一周有更新][Green]](https://github.com/google/gofuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gofuzz)
    * [Tavor](https://github.com/zimmski/tavor) **star:219** 通用模糊测试框架   [![最近一年没有更新][Yellow]](https://github.com/zimmski/tavor)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:4429** 用于驱动和测试 Chrome, Safari, Edge, Android Webviews, 以及其他支持 Chrome 调试协议的产品   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1527** Selenium hub 服务器的替代品，在容器中启动浏览器   [![最近一周有更新][Green]](https://github.com/aerokube/selenoid)   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [cdp](https://github.com/mafredri/cdp) **star:408** 用于Chrome调试协议的类型安全绑定，可与实现该协议的浏览器或其他调试目标一起使用。   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:243** 一个轻量级服务器，可以将 Selenium Wedriver 的请求路由或代理到多个 Selenium hubs   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)
    * [rod](https://github.com/ysmood/rod) **star:13** 一个chrome开发工具控制器，使用方便，安全。   [![最近一周有更新][Green]](https://github.com/ysmood/rod)   [![godoc][GoDoc]](https://godoc.org/github.com/ysmood/rod)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:486** 为Golang实现[failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail)。   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## 文本处理

*用于解析和操作文本的库。 (翻译出错了? 试试 [英文版](README_EN.md#text-processing) 吧~)*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:10305** 快速和优雅的 Scraping 框架。   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![最近一周有更新][Green]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:8501** GoQuery 为 Go 语言带来了一组类似 jQuery 的语法和功能   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [blackfriday](https://github.com/russross/blackfriday) **star:4282** Markdown 解析器   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![最近一周有更新][Green]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [toml](https://github.com/BurntSushi/toml) **star:3097** TOML配置格式(带反射的编码器/解码器)。   [![star > 2000][Awesome]](https://github.com/BurntSushi/toml)   [![godoc][GoDoc]](https://godoc.org/github.com/BurntSushi/toml)
    * [sh](https://github.com/mvdan/sh) **star:2564** Shell解析器和格式化工具。   [![star > 2000][Awesome]](https://github.com/mvdan/sh)   [![最近一周有更新][Green]](https://github.com/mvdan/sh)   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/sh)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:2121** 格式化程序，用于将时间、数字和内存大小转换为可读格式。   [![star > 2000][Awesome]](https://github.com/dustin/go-humanize)   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1484** HTML 清理工具   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [inject](https://github.com/facebookgo/inject) **star:1213** 包注入提供了一个基于反射的注入器。   [![最近一年没有更新][Yellow]](https://github.com/facebookgo/inject)   [![godoc][GoDoc]](https://godoc.org/github.com/facebookgo/inject)   [![归档项目][Archived]](https://github.com/facebookgo/inject)
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1206** 在Go中解析RSS和Atom feeds。   [![最近一周有更新][Green]](https://github.com/mmcdole/gofeed)   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [go-toml](https://github.com/pelletier/go-toml) **star:717** 使用带有查询支持和方便的cli工具的TOML格式库。   [![最近一周有更新][Green]](https://github.com/pelletier/go-toml)   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:586** 一组用于Go的公共正则表达式。   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [slug](https://github.com/gosimple/slug) **star:518** URL 友好的 slug 化工具，支持多种语言   [![godoc][GoDoc]](https://godoc.org/github.com/gosimple/slug)
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:366** Web抓取框架将网站转换为结构化数据。   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [mxj](https://github.com/clbanning/mxj) **star:365** 将XML编码/解码为JSON或map[string]接口{};使用点符号路径和通配符提取值。替换x2j和j2x包。   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  格式化二进制为字符串。
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:348** 解析Graphviz DOT语言。   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [gotext](https://github.com/leonelquinteros/gotext) **star:253** GNU gettext 工具   [![godoc][GoDoc]](https://godoc.org/github.com/leonelquinteros/gotext)
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:247** 函数获取字符或字符串的固定宽度。   [![最近一周有更新][Green]](https://github.com/mattn/go-runewidth)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-runewidth)
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:211** 用于HTML的XPath查询包，允许您通过XPath表达式从HTML文档中提取数据或求值。   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/htmlquery)
    * [goq](https://github.com/andrewstuart/goq) **star:163**  声明式 HTML 编组，使用结构标签和 jQuery 语法 (使用 GoQuery).   [![godoc][GoDoc]](https://godoc.org/github.com/andrewstuart/goq)
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:114** 用于Go语言的NMEA解析器库。   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [sdp](https://github.com/gortc/sdp) **star:87** SDP:会话描述协议[[RFC 4566](https://tools.ietf.org/html/rfc4566)]。   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:77** 用于Go的零宽度字符检测和删除。   [![godoc][GoDoc]](https://godoc.org/github.com/trubitsyn/go-zero-width)
    * [goribot](https://github.com/zhshch2002/goribot) **star:63** 一个简单的golang spider/抓取框架，构建一个三行蜘蛛。   [![godoc][GoDoc]](https://godoc.org/github.com/zhshch2002/goribot)   [![包含中文文档][CN]](https://github.com/zhshch2002/goribot)
    * [align](https://github.com/Guitarbum722/align) **star:63** 对文本进行对齐的通用应用程序。   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [podcast](https://github.com/eduncan911/podcast) **star:61** 在Golang兼容iTunes和RSS 2.0播客生成器   [![godoc][GoDoc]](https://godoc.org/github.com/eduncan911/podcast)
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:58** 生成漂亮的固定链接地址（slug），支持多种语言   [![最近一年没有更新][Yellow]](https://github.com/mozillazg/go-slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-slugify)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub 风格的 Markdown 渲染器 (使用 blackfriday) ，支持代码块高亮以及可点击的锚点
    * [genex](https://github.com/alixaxel/genex) **star:58** 将正则表达式计数并展开为所有匹配的字符串。   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:57** Editorconfig文件解析器和Go操作器。   [![最近一周有更新][Green]](https://github.com/editorconfig/editorconfig-core-go)   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:45** 通过一个 unicode 文本来猜测该文本使用的语言   [![最近一年没有更新][Yellow]](https://github.com/endeveit/guesslanguage)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/guesslanguage)
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:40** 根据正则表达式生成随机字符串   [![godoc][GoDoc]](https://godoc.org/github.com/zach-klippenstein/goregen)
    * [allot](https://github.com/sbstjn/allot) **star:38** 用于CLI工具和机器人的占位符和通配符文本解析。   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [go-vcard](https://github.com/emersion/go-vcard) **star:37** 解析和格式化vCard。   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:33** 固定宽度的文本格式(带反射的编码器/解码器)。   [![最近一周有更新][Green]](https://github.com/ianlopshire/go-fixedwidth)   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [gonameparts](https://github.com/polera/gonameparts) **star:31** 将人名解析为单独的名称部分。   [![godoc][GoDoc]](https://godoc.org/github.com/polera/gonameparts)
    * [did](https://github.com/ockam-network/did) **star:30** DID(分散标识符)解析器和Stringer。   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [Slugify](https://github.com/avelino/slugify) **star:27** 字符串 slug 化的工具。   [![最近一年没有更新][Yellow]](https://github.com/avelino/slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/avelino/slugify)
    * [codetree](https://github.com/aerogo/codetree) **star:14** 解析缩进代码(python、pixy、scarlet等)并返回树结构。   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [enca](https://github.com/endeveit/enca) **star:8** [libenca](http://cihar.com/software/enca/)的最小cgo绑定。   [![最近一年没有更新][Yellow]](https://github.com/endeveit/enca)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/enca)
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:7** Atom 1.0和RSS 2.0的联合提要。   [![最近一年没有更新][Yellow]](https://github.com/zhengchun/syndfeed)   [![godoc][GoDoc]](https://godoc.org/github.com/zhengchun/syndfeed)
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** 将bbCode转换为HTML，使您可以添加对自定义bbCode标记的支持。   [![最近一年没有更新][Yellow]](https://github.com/CalebQ42/bbConvert)   [![godoc][GoDoc]](https://godoc.org/github.com/CalebQ42/bbConvert)
    * [ltsv](https://github.com/Wing924/ltsv) **star:4** 用于Go的高性能[LTSV(标签为Tab Separeted Value)](http://ltsv.org/)阅读器。   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/ltsv)
    * [doi](https://github.com/hscells/doi) **star:4** 文档对象标识符(doi)解析器。   [![最近一年没有更新][Yellow]](https://github.com/hscells/doi)   [![godoc][GoDoc]](https://godoc.org/github.com/hscells/doi)
    * [encdec](https://github.com/mickep76/encdec) **star:3** 软件包为编码器和解码器提供了通用接口。   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:607** 从文本中提取url。   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/xurls)
    * [gotabulate](https://github.com/bndr/gotabulate) **star:232** 使用 Go 语言简单、美观的打印表格数据   [![最近一年没有更新][Yellow]](https://github.com/bndr/gotabulate)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:157** 快速字符串排序算法。   [![最近一年没有更新][Yellow]](https://github.com/yourbasic/radix)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [parth](https://github.com/codemodus/parth) **star:36** URL路径分割解析。   [![最近一年没有更新][Yellow]](https://github.com/codemodus/parth)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/parth)
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:22** 一个基于 sanitization 的 Go 敏感词过滤器。   [![最近一年没有更新][Yellow]](https://github.com/JoshuaDoes/gofuckyourself)   [![godoc][GoDoc]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:18** 将xml或json转换为struct。   [![godoc][GoDoc]](https://godoc.org/github.com/stackerzzq/xj2go)
    * [kace](https://github.com/codemodus/kace) **star:12** 通用大小写转换工具   [![最近一年没有更新][Yellow]](https://github.com/codemodus/kace)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/kace)
    * [Tagify](https://github.com/zoomio/tagify) **star:10** 从给定源生成一组标记。   [![最近一周有更新][Green]](https://github.com/zoomio/tagify)   [![godoc][GoDoc]](https://godoc.org/github.com/zoomio/tagify)
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:9** 字符串参数解析器，能够理解引用及反斜杠。   [![最近一年没有更新][Yellow]](https://github.com/nproc/parseargs-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nproc/parseargs-go)
    * [TySug](https://github.com/Dynom/TySug) **star:4** 关于键盘布局的其他建议。   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
	* [textwrap](https://github.com/isbm/textwrap) **star:1** 从Python实现“textwrap”模块。   [![godoc][GoDoc]](https://godoc.org/github.com/isbm/textwrap)

## 第三方api

*用于访问第三方api的库。 (翻译出错了? 试试 [英文版](README_EN.md#third-party-apis) 吧~)*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5655** AWS 提供的官方go语言 SDK   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   [![最近一周有更新][Green]](https://github.com/aws/aws-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [github](https://github.com/google/go-github) **star:5616** 访问GitHub REST API v3的库。   [![star > 2000][Awesome]](https://github.com/google/go-github)   [![最近一周有更新][Green]](https://github.com/google/go-github)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [google](https://github.com/google/google-api-go-client) **star:2163** 为Go自动生成谷歌api。   [![star > 2000][Awesome]](https://github.com/google/google-api-go-client)   [![最近一周有更新][Green]](https://github.com/google/google-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:2061** 谷歌云api Go 客户端库。   [![star > 2000][Awesome]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![最近一周有更新][Green]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
* [discordgo](https://github.com/bwmarrin/discordgo) **star:1159**  Discord Chat API的客户端。   [![最近一周有更新][Green]](https://github.com/bwmarrin/discordgo)   [![godoc][GoDoc]](https://godoc.org/github.com/bwmarrin/discordgo)
* [stripe](https://github.com/stripe/stripe-go) **star:1098**  Stripe API 的 Go 语言客户端   [![最近一周有更新][Green]](https://github.com/stripe/stripe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1020**  Twitter 1.1 API 的 go 语言客户端   [![最近一周有更新][Green]](https://github.com/ChimeraCoder/anaconda)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [go-twitter](https://github.com/dghubble/go-twitter) **star:922**  Twitter v1.1 api 的客户端.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/go-twitter)
* [minio-go](https://github.com/minio/minio-go) **star:897** 用于Amazon S3兼容云存储的Minio Go库。   [![最近一周有更新][Green]](https://github.com/minio/minio-go)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio-go)
* [facebook](https://github.com/huandu/facebook) **star:846** 支持 Facebook Graph API 的库   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/facebook)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:699**  Go [Atlassian JIRA](https://www.atlassian.com/software/jira)的客户端库   [![最近一周有更新][Green]](https://github.com/andygrunwald/go-jira)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-jira)
* [githubql](https://github.com/shurcooL/githubql) **star:609** 访问GitHub GraphQL API v4的库。   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/githubql)
* [webhooks](https://github.com/go-playground/webhooks) **star:467** GitHub 和 Bitbucket 的Webhook接收器。   [![最近一周有更新][Green]](https://github.com/go-playground/webhooks)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/webhooks)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:343** PayPal支付API的包装器。   [![godoc][GoDoc]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:342** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [go-marathon](https://github.com/gambol99/go-marathon) **star:192**  用于和 Mesosphere's Marathon PAAS 交互的 Go 语言库   [![最近一周有更新][Green]](https://github.com/gambol99/go-marathon)   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [ethrpc](https://github.com/onrik/ethrpc) **star:180**  Ethereum JSON RPC API的客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [Trello](https://github.com/adlio/trello) **star:128**  Trello API的 Go 语言封装。   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [gostorm](https://github.com/jsgilmore/gostorm) **star:124** GoStorm是一个Go库，它实现了编写Storm spout和bolt所需的通信协议，这些协议用于与Storm shell通信。   [![最近一年没有更新][Yellow]](https://github.com/jsgilmore/gostorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [Medium](https://github.com/Medium/medium-sdk-go) **star:123** Medium的OAuth2 API 客户端。   [![最近一年没有更新][Yellow]](https://github.com/Medium/medium-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/Medium/medium-sdk-go)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:112** 一个用于通过XMPP与HipChat通信的golang包。   [![最近一年没有更新][Yellow]](https://github.com/daneharrigan/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [hipchat](https://github.com/andybons/hipchat) **star:108** 这个项目为Hipchat API实现了一个golang客户端库。   [![最近一年没有更新][Yellow]](https://github.com/andybons/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/andybons/hipchat)
* [go-trending](https://github.com/andygrunwald/go-trending) **star:108** 在Github上访问[trends repository](https://github.com/trends)和[developers](https://github.com/trending/developers)的库。   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-trending)
* [cachet](https://github.com/andygrunwald/cachet) **star:78** 使用客户端库[Cachet(开源状态页系统)](https://cachethq.io/)。   [![最近一年没有更新][Yellow]](https://github.com/andygrunwald/cachet)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/cachet)
* [pushover](https://github.com/gregdel/pushover) **star:70**  Go 包装的 Pushover API。   [![最近一年没有更新][Yellow]](https://github.com/gregdel/pushover)   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [wit-go](https://github.com/wit-ai/wit-go) **star:65** wit.ai HTTP API 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:57** [Internet Game Database API](https://api.igdb.com/) 客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Clarifai API的客户端。   [![最近一年没有更新][Yellow]](https://github.com/samuelcouch/clarifai)   [![godoc][GoDoc]](https://godoc.org/github.com/samuelcouch/clarifai)
* [megos](https://github.com/andygrunwald/megos) **star:56** 用于访问[Apache Mesos](http://mesos.apache.org/)集群的客户端库。   [![最近一年没有更新][Yellow]](https://github.com/andygrunwald/megos)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [circleci](https://github.com/jszwedko/go-circleci) **star:49** CircleCI的API的客户端   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [gads](https://github.com/emiddleton/gads) **star:47**  Google Adwords 非官方 API   [![godoc][GoDoc]](https://godoc.org/github.com/emiddleton/gads)
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:40**  xkcd API 的客户端。   [![最近一年没有更新][Yellow]](https://github.com/nishanths/go-xkcd)   [![godoc][GoDoc]](https://godoc.org/github.com/nishanths/go-xkcd)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:40**  Go MusicBrainz WS2客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:40** [Amazon Product Advertising API](https://program.amazon.com/gp/advertising/api/detail/main.html)的客户端库。   [![最近一年没有更新][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [simples3](https://github.com/rhnvrm/simples3) **star:36** 使用REST和用Go编写的V4签名的AWS S3库非常简单。   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:36** 运行时Robot v2 API 的 Go 语言封装和命令行客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [fcm](https://github.com/maddevsio/fcm) **star:36**  Firebase Cloud Messaging 的 Go 语言库   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [slack](https://github.com/nlopes/slack)  Slack API。
* [gosip](https://github.com/koltyakov/gosip) **star:35** 去客户端库SharePoint API。   [![最近一周有更新][Green]](https://github.com/koltyakov/gosip)   [![godoc][GoDoc]](https://godoc.org/github.com/koltyakov/gosip)
* [golyrics](https://github.com/mamal72/golyrics) **star:34** Golyrics是一个从Wikia网站获取音乐歌词数据的Go库。   [![最近一年没有更新][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [mixpanel](https://github.com/dukex/mixpanel) **star:32** Mixpanel是一个库，用于跟踪事件并将Mixpanel概要文件更新从go应用程序发送到Mixpanel。   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [translate](https://github.com/poorny/translate) **star:29**  Go 在线翻译包。   [![最近一年没有更新][Yellow]](https://github.com/poorny/translate)   [![godoc][GoDoc]](https://godoc.org/github.com/poorny/translate)
* [gcm](https://github.com/Aorioli/gcm) **star:29**  Google Cloud Messaging 库   [![最近一年没有更新][Yellow]](https://github.com/Aorioli/gcm)   [![godoc][GoDoc]](https://godoc.org/github.com/Aorioli/gcm)
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb) **star:28** 电影数据库API v3的Golang包装器。   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/golang-tmdb)
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:27** 使用[Unsplash.com](https://unsplash.com) API的客户端库。   [![最近一周有更新][Green]](https://github.com/hbagdi/go-unsplash)   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [gami](https://github.com/bit4bit/gami) **star:27**  Asterisk Manager Interface 的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![归档项目][Archived]](https://github.com/bit4bit/gami)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:24**  YNAB API 的 Go 语言封装。   [![godoc][GoDoc]](https://godoc.org/github.com/brunomvsouza/ynab.go)
* [spotify](https://github.com/rapito/go-spotify) **star:21**  用于接入 Spotify WEB API 的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/rapito/go-spotify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-spotify)
* [steam](https://github.com/sostronk/go-steam) **star:19**  用于与Steam服务器进行交互的库。   [![godoc][GoDoc]](https://godoc.org/github.com/sostronk/go-steam)
* [shopify](https://github.com/rapito/go-shopify) **star:19** 一个用于通过 Shopify API 进行增删改查的 Go 语言库。   [![最近一年没有更新][Yellow]](https://github.com/rapito/go-shopify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:19**  Twitch v3 API 的客户端。   [![最近一年没有更新][Yellow]](https://github.com/knspriggs/go-twitch)   [![godoc][GoDoc]](https://godoc.org/github.com/knspriggs/go-twitch)
* [patreon-go](https://github.com/mxpv/patreon-go) **star:19**  Go Patreon API库。   [![godoc][GoDoc]](https://godoc.org/github.com/mxpv/patreon-go)
* [codeship-go](https://github.com/codeship/codeship-go) **star:17**  Codeship API v2的客户端。   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:17** 用于访问 BreweryDB API的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/naegelejd/brewerydb)   [![godoc][GoDoc]](https://godoc.org/github.com/naegelejd/brewerydb)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:16** [MyAnimeList API](http://myanimelist.net/modules.php?go=api)的客户端库。   [![最近一年没有更新][Yellow]](https://github.com/nstratos/go-myanimelist)   [![godoc][GoDoc]](https://godoc.org/github.com/nstratos/go-myanimelist)
* [textbelt](https://github.com/dietsche/textbelt) **star:16** textbelt.com txt messaging API 的go语言客户端。   [![最近一年没有更新][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:15** Go [imgur](https://imgur.com)的客户端库   [![最近一年没有更新][Yellow]](https://github.com/koffeinsource/go-imgur)   [![godoc][GoDoc]](https://godoc.org/github.com/koffeinsource/go-imgur)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Coinpaprika API的客户端。   [![最近一年没有更新][Yellow]](https://github.com/coinpaprika/coinpaprika-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:11** 简单的包装，方便的谷歌分析报告。   [![最近一年没有更新][Yellow]](https://github.com/chonthu/go-google-analytics)   [![godoc][GoDoc]](https://godoc.org/github.com/chonthu/go-google-analytics)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:10** HackerNews API的小型Go客户端。   [![最近一年没有更新][Yellow]](https://github.com/PaulRosset/go-hacknews)   [![godoc][GoDoc]](https://godoc.org/github.com/PaulRosset/go-hacknews)
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:10** 使用[LastPass](https://www.lastpass.com/) API的客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [smite](https://github.com/sergiotapia/smitego) **star:10** 对 Smite game API 的封装。   [![最近一年没有更新][Yellow]](https://github.com/sergiotapia/smitego)   [![godoc][GoDoc]](https://godoc.org/github.com/sergiotapia/smitego)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** 用于接入 statdns.com API 的库——RRDA API。通过HTTP协议进行 DNS查询。   [![最近一年没有更新][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:7**  SPTrans Olho Vivo API 的客户端。   [![最近一年没有更新][Yellow]](https://github.com/sergioaugrod/go-sptrans)   [![godoc][GoDoc]](https://godoc.org/github.com/sergioaugrod/go-sptrans)
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph 发布平台 API 客户端。
* [go-here](https://github.com/abdullahselek/go-here) **star:6** 转到基于这里位置的api的客户端库。   [![godoc][GoDoc]](https://godoc.org/github.com/abdullahselek/go-here)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** [谷歌G Suite Email Audit API](https://developer.google.com/admin-sdk/email-audit/) 的客户端。   [![最近一年没有更新][Yellow]](https://github.com/ngs/go-google-email-audit-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-google-email-audit-api)
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6**  Tumblr v2 API 的 Go 语言封装。   [![最近一年没有更新][Yellow]](https://github.com/mattcunningham/gumblr)   [![godoc][GoDoc]](https://godoc.org/github.com/mattcunningham/gumblr)
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection) **star:5** 去模块工作与[邮递员收集](https://learning.getpostman.com/docs/postman/collections/creating-collections/)(兼容失眠)。   [![最近一周有更新][Green]](https://github.com/rbretecher/go-postman-collection)   [![godoc][GoDoc]](https://godoc.org/github.com/rbretecher/go-postman-collection)
* [zooz](https://github.com/gojuno/go-zooz) **star:5**  Zooz API 的 Go 语言客户端。   [![最近一年没有更新][Yellow]](https://github.com/gojuno/go-zooz)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** 无任何依赖的[Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/pdfs/documentation/utmonaws/sophos-ut-restful-api.pdf?la=en)客户端   [![godoc][GoDoc]](https://godoc.org/github.com/esurdam/go-sophos)
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:5** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)
* [twitter-scraper](https://github.com/n0madic/twitter-scraper) **star:3** 刮Twitter前端API没有认证和限制。   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/twitter-scraper)
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** 用于与[Chronos](https://mesos.github.io/chronos/)作业调度程序进行交互的Go库   [![最近一年没有更新][Yellow]](https://github.com/axelspringer/go-chronos)   [![godoc][GoDoc]](https://godoc.org/github.com/axelspringer/go-chronos)
* [google-play-scraper](https://github.com/n0madic/google-play-scraper) **star:2** 从谷歌播放存储获取数据。   [![最近一周有更新][Green]](https://github.com/n0madic/google-play-scraper)   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/google-play-scraper)
* [gopaapi5](https://github.com/utekaravinash/gopaapi5) **star:2** 去客户端库[亚马逊产品广告API 5.0](https://webservices.amazon.com/paapi5/documentation/)。   [![最近一周有更新][Green]](https://github.com/utekaravinash/gopaapi5)   [![godoc][GoDoc]](https://godoc.org/github.com/utekaravinash/gopaapi5)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1**  TripAdvisor API 的 Go 语言封装。   [![godoc][GoDoc]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
* [vl-go](https://github.com/verifid/vl-go) **star:1** 使用客户端库中的VerifID身份验证层API。   [![godoc][GoDoc]](https://godoc.org/github.com/verifid/vl-go)
* [libgoffi](https://github.com/clevabit/libgoffi) **star:1** 库适配器工具箱，用于本地[libffi](http://sourceware.org/libffi/)集成   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** Playlyfe Rest API Go SDK。   [![最近一年没有更新][Yellow]](https://github.com/playlyfe/playlyfe-go-sdk)   [![godoc][GoDoc]](https://godoc.org/github.com/playlyfe/playlyfe-go-sdk)

## 公用事业公司

*可以让你的生活变得更简单的实用工具.。 (翻译出错了? 试试 [英文版](README_EN.md#utilities) 吧~)*

* [fzf](https://github.com/junegunn/fzf) **star:27737** 用Go编写的命令行模糊查找器。   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![最近一周有更新][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [hub](https://github.com/github/hub) **star:19231** 封装了 git 命令，提供了额外的功能用于在终端中和 Github 进行交互。   [![star > 2000][Awesome]](https://github.com/github/hub)   [![最近一周有更新][Green]](https://github.com/github/hub)   [![godoc][GoDoc]](https://godoc.org/github.com/github/hub)
* [ctop](https://github.com/bcicen/ctop) **star:9603** [Top-like](http://ctop.sh)接口(例如htop)， 用于容器数据收集。   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:5296** 尽可能快速的发布 Go 语言二进制文件。   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![最近一周有更新][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [goreporter](https://github.com/wgliang/goreporter)  进行代码静态分析，单元测试，代码检视并生成代码质量报告的工具
* [godropbox](https://github.com/dropbox/godropbox) **star:3849** 用于编写 Go 语言服务／应用的库，来自 Dropbox.。   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [xferspdy](https://github.com/monmohan/xferspdy)  Xferspdy在golang中提供二进制diff和补丁库。
* [panicparse](https://github.com/maruel/panicparse)  将类似的协程分组并对调用栈进行着色
* [onecache](https://github.com/adelowo/onecache)  支持多个后端存储(Redis、Memcached、文件系统等)的缓存库。
* [peco](https://github.com/peco/peco)  简单的交互过滤工具。
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match)  模式匹配图书馆。
* [okrun](https://github.com/xta/okrun)   Go 运行错误 steamroller。
* [pgo](https://github.com/arthurkushman/pgo)  用于PHP社区的 Convenient 函数。
* [pm](https://github.com/VividCortex/pm)  进程(即goroutine)管理器与HTTP API。
* [ptr](https://github.com/gotidy/ptr)  提供函数的包，用于从基本类型的常量简化指针的创建。
* [r](https://github.com/is5/r)  类似于python的“range()”经验。
* [rclient](https://github.com/zpatrick/rclient)  可读、灵活、易于使用的REST api客户端。
* [olaf](https://github.com/btnguyen2k/olaf)  Twitter Snowflake 在Go中实现。
* [mssqlx](https://github.com/linxGnu/mssqlx)  数据库客户端，用于主-从 (或主-主) 数据库，集成了简单的、轻量级的轮询调度负载均衡。
* [myhttp](https://github.com/inancgumus/myhttp)  简单的API，使HTTP GET请求与超时支持。
* [multitick](https://github.com/VividCortex/multitick)  用于 aligned tickers 的多路复用
* [repeat](https://github.com/ssgreg/repeat)  执行不同的后 backoff 策略，这对重新尝试操作和心跳非常有用。
* [mole](https://github.com/davrodpin/mole)  cli应用程序可以轻松创建ssh隧道。
* [moldova](https://github.com/StabbyCutyou/moldova)  基于输入目标生成随机数据的工具
* [mmake](https://github.com/tj/mmake)  现代 Make 工具
* [minquery](https://github.com/icza/minquery)  MongoDB / mgo.v2, 支持高效分页查询(用于继续列出我们停止的文档的游标)。
* [minify](https://github.com/tdewolff/minify)  用于HTML、CSS、JS、XML、JSON和SVG文件格式的快速缩小器。
* [mimetype](https://github.com/gabriel-vasile/mimetype)  用于基于神奇数字的MIME类型检测的包。
* [mimesniffer](https://github.com/aofei/mimesniffer)  一个用于Go的MIME类型嗅探器。
* [mimemagic](https://github.com/zRedShift/mimemagic)  纯粹 Go 超性能MIME嗅探库/实用程序。
* [mergo](https://github.com/imdario/mergo)  用于将结构体和map合并进 Go 语言的工具。对于配置默认值，避免杂乱的if语句很有帮助。
* [mc](https://github.com/minio/mc)  Minio Client 提供了一组工具，用于操作 Amazon S3 兼容云存储和文件系统。
* [realize](https://github.com/tockins/realize)  Go 语言构建系统，可以监控文件变化并重新加载。运行，构建，监控文件并支持自定义路径。
* [rerate](https://github.com/abo/rerate)  基于 Redis 的速率计数器和限速器
* [request](https://github.com/mozillazg/request)  Go 语言版的 HTTP Requests for Humans™.。
* [slicer](https://github.com/leaanthony/slicer)  使处理切片更容易。
* [wuzz](https://github.com/asciimoo/wuzz)  用于HTTP检查的交互式cli工具。
* [util](https://github.com/shomali11/util)  有用实用函数的集合。(字符串，并发，操作，…)
* [usql](https://github.com/knq/usql)  usql 是一个通用的命令行接口，用于操作 sql 数据库。
* [UNIS](https://github.com/esemplastic/unis)  Go 语言字符串处理函数的通用架构 。
* [ugo](https://github.com/alxrm/ugo)  uGo 是一个切片工具箱，有着和 Go 语言一致的语法法。
* [toolbox](https://github.com/viant/toolbox)  切片, map, multimap, 结构体, 函数,数据转换工具。服务路由，宏求值和标记器。
* [tome](https://github.com/cyruzin/tome)  Tome设计用于对简单的RESTful api进行分页。
* [Task](https://github.com/go-task/task)  简单的“Go”的选择。
* [structs](https://github.com/PumpkinSeed/structs)  简单来讲就是 "Make" 的替代品。
* [Storm](https://github.com/asdine/storm)  一个简单又强大的用于 BoltDB 的工具
* [sqlx](https://github.com/jmoiron/sqlx)  为内建的数据库/sql 软件包提供一组扩展。
* [spinner](https://github.com/briandowns/spinner)   一个 Go 语言软件包，提供多种选项，方便在终端中创建加载动画。
* [sorty](https://github.com/jfcg/sorty)  快速并发/并行排序。
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv)  基本类型之间的片转换。
* [limiters](https://github.com/mennanov/limiters)  Golang中的分布式应用程序速率限制器，具有可配置的后端和分布式锁。
* [slice](https://github.com/psampaz/slice)  用于普通Go切片操作的类型安全函数。
* [silk](https://github.com/chrispassas/silk)  阅读silk netflow文件。
* [shutdown](https://github.com/ztrue/shutdown)  ' os的应用程序关闭挂钩。信号的处理。
* [serve](https://github.com/syntaqx/serve)  任何您需要的静态http服务器。
* [scan](https://github.com/blockloop/scan)  扫描golang的sql。行直接指向结构、片或基本类型。
* [robustly](https://github.com/VividCortex/robustly)  有弹性的执行函数，遇到错误时捕获并重新运行。
* [retry-go](https://github.com/rafaeljesus/retry-go)  对 Go 来说，重试变得简单而容易。
* [retry](https://github.com/shafreeck/retry)  一个相当简单的库，以确保您的工作可以完成。
* [retry](https://github.com/thedevsaddam/retry)  简单易用的重试机制包，为 Go 。
* [retry](https://github.com/percolate/retry)  一个简单但高度可配置的Go重试包。
* [retry](https://github.com/kamilsk/retry)  基于上下文的功能机制，反复执行命令直到成功。
* [rest-go](https://github.com/edermanoel94/rest-go)  一个包，它为使用rest api提供了许多有用的方法。
* [rerun](https://github.com/ivpusic/rerun)  当源代码发生更改时，重新编译和重新运行go应用程序。
* [lrserver](https://github.com/jaschaephraim/lrserver)  LiveReload 服务器。
* [intrinsic](https://github.com/mengzhuo/intrinsic)  不需要编写任何汇编代码就能使用 x86 SIMD。
* [koazee](https://github.com/wesovilabs/koazee)  库的灵感来自于延迟计算和函数式编程，从而减少了使用数组的麻烦。
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2387** 实现 Hystrix 风格的、程序员预定义的 fallback 机制（熔断。   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [jsend](https://github.com/clevergo/jsend)  在Go中编写JSend的实现。
* [jump](https://github.com/gsamokovarov/jump)  通过学习你的习惯，可以帮助你更快地导航。
* [immortal](https://github.com/immortal/immortal)  \*nix 跨平台 (与操作系统无关的)监控程序。
* [go-funk](https://github.com/thoas/go-funk) **star:1652** 现代 Go 语言工具库，提供了很多有用的工具 (map, find, contains, filter, chunk, reverse, ...)   [![最近一周有更新][Green]](https://github.com/thoas/go-funk)   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [filetype](https://github.com/h2non/filetype) **star:1059** 通过数字签名来推测文件类型。   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [boilr](https://github.com/tmrts/boilr) **star:1059** 非常快的CLI工具，用于从样板模板创建项目。   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:859** 接通断路器。   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:785** git-time-metric - 。   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [htcat](https://github.com/htcat/htcat) **star:511** 并行及流水线的 HTTP GET 工具。   [![最近一年没有更新][Yellow]](https://github.com/htcat/htcat)   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [go-dry](https://github.com/ungerik/go-dry) **star:446** DRY(don't repeat yourself)库。   [![最近一年没有更新][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [circuit](https://github.com/cep21/circuit) **star:441** 一个高效和功能齐全的 类似 Hystrix Go 实现断路器模式。   [![godoc][GoDoc]](https://godoc.org/github.com/cep21/circuit)
* [gopencils](https://github.com/bndr/gopencils) **star:431** 小而简单的包，可以轻松地使用REST api。   [![最近一年没有更新][Yellow]](https://github.com/bndr/gopencils)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [godaemon](https://github.com/VividCortex/godaemon) **star:431** 用于编写守护进程的工具   [![最近一年没有更新][Yellow]](https://github.com/VividCortex/godaemon)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/godaemon)
* [ergo](https://github.com/cristianoliveira/ergo) **star:363** 管理运行在不同端口上的多个本地服务变得很容易。   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [go-rate](https://github.com/beefsack/go-rate) **star:301**  Go 限速器。   [![最近一年没有更新][Yellow]](https://github.com/beefsack/go-rate)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [clockwork](https://github.com/jonboulle/clockwork) **star:258** 一个简单的假 clock 。   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [gohper](https://github.com/cosiner/gohper) **star:252** 多种能够帮助你进行软件开发的工具和模块。   [![最近一年没有更新][Yellow]](https://github.com/cosiner/gohper)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/gohper)   [![归档项目][Archived]](https://github.com/cosiner/gohper)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:246** 结构体拷贝   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [gubrak](https://github.com/novalagung/gubrak) **star:209** 带有语法糖的Golang实用工具，就像lodash。   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:193** Go 语言全局事件触发器，通过 id 和触发器，在程序的任何地方注册事件。   [![最近一年没有更新][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [gotenv](https://github.com/subosito/gotenv) **star:169** 从 `.env` 或者任何 `io.Reader`。   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:165** Go:generate 工具，用于构建 Go 语言插件(1.8 only)，并对导出的符号进行包装。   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [Death](https://github.com/vrecan/death) **star:149** 利用信号管理应用程序的关闭。   [![最近一年没有更新][Yellow]](https://github.com/vrecan/death)   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [apm](https://github.com/topfreegames/apm) **star:137** Go 语言进程管理工具具有HTTP API.。   [![最近一年没有更新][Yellow]](https://github.com/topfreegames/apm)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/apm)
* [chyle](https://github.com/antham/chyle) **star:117** 使用具有多种配置可能性的git存储库生成变更日志。   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:116** 用Go编写的XML站点地图生成器。   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:91** 纯Go bsdiff和bspatch库和CLI工具。   [![最近一年没有更新][Yellow]](https://github.com/gabstv/go-bsdiff)   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [go-health](https://github.com/Talento90/go-health) **star:74** 健康包简化了向服务中添加健康检查的方式。   [![最近一年没有更新][Yellow]](https://github.com/Talento90/go-health)   [![godoc][GoDoc]](https://godoc.org/github.com/Talento90/go-health)
* [handy](https://github.com/miguelpragier/handy) **star:51** 许多实用程序和帮助程序，如字符串处理程序/格式化程序和验证器。   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [golog](https://github.com/mlimaloureiro/golog) **star:48** 简单、轻量级的命令后工具，用于对你的计划任务进行跟踪。   [![最近一年没有更新][Yellow]](https://github.com/mlimaloureiro/golog)   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** 解析你 Go 语言代码中的 TODOs（待办事项）。   [![最近一年没有更新][Yellow]](https://github.com/asticode/go-astitodo)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [goreadability](https://github.com/philipjkim/goreadability) **star:46** 网页摘要提取器使用Facebook开放图形和arc90的可读性。   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:45** conseilSeaweedFS 客户端，几乎具有全部的特性。   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [gaper](https://github.com/maxcnunes/gaper) **star:42** 当Go项目崩溃或一些人看到文件更改时，构建并重新启动该项目。   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [goback](https://github.com/carlescere/goback) **star:42**  一个 Go 语言的简单的指数补偿包。   [![最近一年没有更新][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [golarm](https://github.com/msempere/golarm) **star:38** 告警（支持系统事件）。   [![最近一年没有更新][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:38** 通用多工作站剪切板，使用类似 S3 的后端作为存储。   [![godoc][GoDoc]](https://godoc.org/github.com/jutkko/copy-pasta)
* [beyond](https://github.com/wesovilabs/beyond) **star:34** Go工具将带你进入AOP的世界!   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [gpath](https://github.com/tenntenn/gpath) **star:33**  用于简化结构体域访问的库。   [![最近一年没有更新][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [dbt](https://github.com/nikogura/dbt) **star:26** 用于从中心可信存储库运行自更新签名二进制文件的框架。   [![最近一周有更新][Green]](https://github.com/nikogura/dbt)   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [countries](https://github.com/biter777/countries) **star:25** 完整实现ISO-3166-1、ISO-4217、ITU-T E.164、Unicode CLDR和IANA ccTLD标准。   [![最近一周有更新][Green]](https://github.com/biter777/countries)   [![godoc][GoDoc]](https://godoc.org/github.com/biter777/countries)
* [generate](https://github.com/go-playground/generate) **star:23** 针对一个路径或环境变量，递归的执行 Go generate，可以通过正则表达式来进行过滤。   [![最近一年没有更新][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [gostrutils](https://github.com/ik5/gostrutils) **star:22** 字符串操作和转换函数的集合。   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:21** 一个小巧的 Go 语言库用于生成占位图片。   [![最近一年没有更新][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:21** 用于在osx、windows和linux上执行shell命令的库。   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [delve](https://github.com/derekparker/delve) **star:20** Go 语言调试器   [![最近一周有更新][Green]](https://github.com/derekparker/delve)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [evaluator](https://github.com/nullne/evaluator) **star:20** 基于 s-expression。它很简单，很容易扩展。   [![最近一年没有更新][Yellow]](https://github.com/nullne/evaluator)   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [filter](https://github.com/gookit/filter) **star:20** 提供Go数据的过滤、清理和转换。   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:16**  用于将结构体编码进 http 头的 Go 语言库   [![最近一年没有更新][Yellow]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** 编译时控制的日志，让你的 release 包变得更小而不需移除 debug 调用。   [![最近一年没有更新][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [filler](https://github.com/yaronsumel/filler) **star:15** 使用“fill”标签填充结构的小工具。   [![最近一年没有更新][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [ghokin](https://github.com/antham/ghokin) **star:13** 没有外部依赖的gherkin (cucumber, behat…)并行格式化程序。   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [ctxutil](https://github.com/posener/ctxutil) **star:12** 上下文工具。   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [command](https://github.com/txgruppi/command) **star:10** 命令模式，支持线程安全的串行、并行调度。   [![最近一年没有更新][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [backscanner](https://github.com/icza/backscanner) **star:9** 类似 bufio 的扫描器，但它以相反的顺序读取和返回行，从给定的位置开始，然后返回。   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:5** 打包处理问题细节。   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
* [blank](https://github.com/Henry-Sarabia/blank) **star:4** 验证或删除字符串中的空白。   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [compare](https://github.com/posener/compare) **star:1** 启用更具可读性和更容易的比较任务。   [![最近一周有更新][Green]](https://github.com/posener/compare)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/compare)

## UUID

*用于处理uuid的库。 (翻译出错了? 试试 [英文版](README_EN.md#uuid) 吧~)*

* [goid](https://github.com/jakehl/goid)  生成和解析RFC4122兼容的V4 uuid。
* [nanoid](https://github.com/aidarkhanov/nanoid)  一个小而有效的Go唯一字符串ID生成器。
* [sno](https://github.com/muyo/sno)  使用嵌入元数据的紧凑、可排序和快速的惟一id。
* [ulid](https://github.com/oklog/ulid)  实现了ULID(普遍唯一的词典分类标识符)。
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  没有麻烦，安全，快速的唯一标识符与命令。
* [uuid](https://github.com/agext/uuid)  使用快速或加密质量的随机节点标识符生成、编码和解码UUIDs v1。
* [uuid](https://github.com/gofrs/uuid)  通用唯一标识符(UUID)的实现。支持uuid的创建和解析。积极维护satori uuid的fork。
* [wuid](https://github.com/edwingeng/wuid)  一个非常快的唯一数字生成器，比UUID快10-135倍。

## 验证

*库进行验证。 (翻译出错了? 试试 [英文版](README_EN.md#validation) 吧~)*

* [checkdigit](https://github.com/osamingo/checkdigit)  提供校验数字算法(Luhn, Verhoeff, Damm)和计算器(ISBN, EAN, JAN, UPC等)。
* [gody](https://github.com/guiferpa/gody)  :一个轻量级的结构验证器。
* [govalid](https://github.com/twharmon/govalid)  快速、基于标签的结构验证。
* [govalidator](https://github.com/asaskevich/govalidator)  用于字符串，数字，切片和结构的验证器和sanitizers。
* [govalidator](https://github.com/thedevsaddam/govalidator)  用简单的规则验证Golang请求数据。深受Laravel请求验证的启发。
* [jio](https://github.com/faceair/jio)  jio是一个json模式验证器，类似于[joi](https://github.com/hapijs/joi)。
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)  支持各种数据类型(结构、字符串、映射、片等)的验证，使用可配置和可扩展的验证规则，这些规则在通常的代码构造中指定，而不是在结构标签中指定。
* [terraform-validator](https://github.com/thazelart/terraform-validator)  一种规范和约定验证器。
* [validate](https://github.com/gookit/validate)   Go 封装数据验证和过滤。支持验证映射、结构、请求(表单、JSON、url)。值，上载文件)数据和更多特性。
* [validate](https://github.com/gobuffalo/validate)  这个包提供了一个框架，用于为Go应用程序编写验证。
* [validator](https://github.com/go-playground/validator)   Go 结构体及域验证，包括：跨域、跨结构体, Map, 切片和数组。

## 版本控制

*用于版本控制的库。 (翻译出错了? 试试 [英文版](README_EN.md#version-control) 吧~)*

* [gh](https://github.com/rjeczalik/gh)  用于GitHub webhook的可编写脚本的服务器和net/http中间件。
* [git2go](https://github.com/libgit2/git2go)   libgit2 的 Go 语言接口。
* [go-git](https://github.com/src-d/go-git)  纯Go中高度可扩展的Git实现。
* [go-vcs](https://github.com/sourcegraph/go-vcs)  在Go中操作和检查VCS存储库。
* [hercules](https://github.com/src-d/hercules)  从Git存储库历史中获得高级见解。
* [hgo](https://github.com/beyang/hgo)  Hgo是一组Go包的集合，提供对本地Mercurial存储库的读取访问。

## 视频

*用于操作视频的库。 (翻译出错了? 试试 [英文版](README_EN.md#video) 吧~)*

* [gmf](https://github.com/3d0c/gmf)   FFmpeg av\* 库的 Go 语言接口。
* [go-astisub](https://github.com/asticode/go-astisub)  使用 Go 语言操作字幕(.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.)。
* [go-astits](https://github.com/asticode/go-astits)  在GO中解析和演示MPEG传输流(.ts)。
* [go-m3u8](https://github.com/quangngotan95/go-m3u8)  苹果m3u8播放列表的解析器和生成器库。
* [go-mpd](https://github.com/unki2aut/go-mpd)  MPEG-DASH清单文件的解析器和生成器库。
* [goav](https://github.com/giorgisio/goav)  FFmpeg的Comphrensive。
* [gst](https://github.com/ziutek/gst)   GStreamer的Go工具。
* [libgosubs](https://github.com/wargarblgarbl/libgosubs)  字幕格式支持 .srt、.ttml和.ass。
* [libvlc-go](https://github.com/adrg/libvlc-go)  Go绑定libvlc 2.X/3.X/4。X(由VLC媒体播放器使用)。
* [m3u8](https://github.com/grafov/m3u8)  苹果HLS的M3U8播放列表解析器和生成器库。
* [v4l](https://github.com/korandiz/v4l)  用于Linux的视频捕捉库，用Go编写。

## Web框架

*全栈 web 框架。 (翻译出错了? 试试 [英文版](README_EN.md#web-frameworks) 吧~)*

* [aah](https://aahframework.org)  可伸缩、高性能、快速开发的Go Web框架。
* [REST Layer](http://rest-layer.io)  框架，用于在数据库之上构建REST/GraphQL API，主要是通过代码进行配置。
* [hiboot](https://github.com/hidevopsio/hiboot)  hiboot是一个高性能的web应用程序框架，支持自动配置和依赖注入。
* [Macaron](https://github.com/go-macaron/macaron)  Macaron 是一个高效的模块化设计的web框架
* [mango](https://github.com/paulbellamy/mango)  ManGo 是一个模块化 web 应用框架，受到 Rack 和 PEP333 的启发。
* [Microservice](https://github.com/claygod/microservice)  创建微服务的框架，用Golang编写。
* [neo](https://github.com/ivpusic/neo)  Neo是一个非常简单且快速的Web框架API。
* [patron](https://github.com/beatlabs/patron)  Patron是一个遵循最佳云实践的微服务框架，专注于提升开发效率。
* [Resoursea](https://github.com/resoursea/api)  用于快速编写基于资源的服务的REST框架。
* [Revel](https://github.com/revel/revel)  用于Go语言的高效web框架。
* [goweb](https://github.com/twharmon/goweb)  具有路由、websockets、日志、中间件、静态文件服务器(可选gzip)和自动TLS的Web框架。
* [rex](https://github.com/goanywhere/rex)   Rex 是一个用于进行模块化开发的库，基于Gorilla/mux 完全兼容大多数的 net/HTTP.
* [rux](https://github.com/gookit/rux)  简单而快速的web框架，可用于构建golang HTTP应用程序
* [tango](https://github.com/lunny/tango)  微型的、支持插件的 web 框架。
* [tigertonic](https://github.com/rcrowley/go-tigertonic)  用于构建 JSON web 服务的 Go 语言框架，受到 Dropwizard 的启发。
* [uAdmin](https://github.com/uadmin/uadmin)  受到 Sinatra 启发的 Go 语言 web 框架。
* [utron](https://github.com/gernest/utron)  Go(Golang)的轻量级MVC框架。
* [vox](https://github.com/aisk/vox)  一个面向人类的golang web框架，深受Koa的启发。
* [WebGo](https://github.com/bnkamalesh/webgo)  构建web应用程序的微框架;处理程序链接、中间件和上下文注入。与标准库兼容的HTTP处理程序(即http.HandlerFunc)。
* [Goyave](https://github.com/System-Glitch/goyave)  功能齐全的web框架旨在干净的代码和快速的开发，具有强大的内置功能。
* [gongular](https://github.com/mustafaakin/gongular)   快速 Go web 框架，支持输入映射／验证以及依赖注入。
* [Aero](https://github.com/aerogo/aero)  高性能的Go web框架，在Lighthouse中达到最高分。
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce)  提供使用干净的体系结构(如DDD和端口和适配器)的电子商务功能，您可以使用这些功能来构建灵活的电子商务应用程序。
* [Air](https://github.com/aofei/air)  一个理想的精细化的Go web框架。
* [Banjo](https://github.com/nsheremet/banjo)  非常简单和快速的网络框架 Go 。
* [Beego](https://github.com/astaxie/beego)  beego是一种用于 Go 编程语言的开源高性能web框架。
* [Buffalo](http://gobuffalo.io)  为 Go 语言带来堪比 Rails 的高生产效率!
* [Echo](https://github.com/labstack/echo)  高性能、极简的Go web框架。
* [Fiber](https://github.com/gofiber/fiber)  一个灵感来自Express.js的web框架构建在Fasthttp上。
* [Fireball](https://github.com/zpatrick/fireball)  感觉更加自然的 web 框架。
* [Flamingo](https://github.com/i-love-flamingo/flamingo)  可插拔web项目的框架。包括模块的概念和提供DI、Configareas、i18n、模板引擎、graphql、可观察性、安全性、事件、路由和反向路由等功能。
* [Gin](https://github.com/gin-gonic/gin)  Gin是一个用Go编写的web框架!它具有一个类似于martini的API，性能更好，速度快40倍。
* [Gondola](https://github.com/rainycape/gondola)  web框架写的网站越快越好。
* [Ginrpc](https://github.com/xxjwxc/ginrpc)  Gin参数自动绑定工具，Gin rpc工具。
* [Gizmo](https://github.com/NYTimes/gizmo)  《纽约时报》使用的微服务工具包。
* [go-json-rest](https://github.com/ant0ine/go-json-rest)  设置RESTful JSON API的快速简便方法。
* [go-rest](https://github.com/ungerik/go-rest)  小型的 REST 框架。
* [Goa](https://github.com/goadesign/goa)  Goa为在Go中开发远程api和微服务提供了一种全面的方法。
* [goa](https://github.com/goa-go/goa)  goa就像golang中的koajs一样，是一个灵活、轻量级、高性能和可扩展的基于中间件的web框架。
* [Golax](https://github.com/fulldump/golax)  一个非Sinatra快速HTTP框架，支持谷歌自定义方法、深度拦截器、递归等。
* [Golf](https://github.com/dinever/golf)  Golf 是一个快速、简单、轻量级的 Go 语言微型 web 框架。具有强大的功能且没有标准库以外的依赖。
* [YARF](https://github.com/yarf-framework/yarf)  快速微框架，旨在以快速和简单的方式构建REST api和web服务。

### 中间件

#### 仿真中间件

* [client-timing](https://github.com/posener/client-timing)  用于服务器定时报头的HTTP客户机。
* [CORS](https://github.com/rs/cors)  轻松地向API添加CORS功能。
* [formjson](https://github.com/rs/formjson)  透明地将JSON输入作为标准表单POST处理。
* [go-server-timing](https://github.com/mitchellh/go-server-timing)  添加/解析Server-Timing头。
* [Limiter](https://github.com/ulule/limiter)  简单的速度限制中间件。
* [ln-paywall](https://github.com/philippgille/ln-paywall)  使用Lightning Network(比特币)实现基于每个请求的api货币化中间件。
* [Tollbooth](https://github.com/didip/tollbooth)  限制速率的 HTTP 请求处理程序。
* [XFF](https://github.com/sebest/xff)  处理 X-Forwarded-For 头的中间件。

#### 用于创建HTTP中间件的库

* [alice](https://github.com/justinas/alice)  用于连接中间件的库，简单无痛苦。
* [catena](https://github.com/codemodus/catena)  HTTP.Handler wrapper catenation (和chain具有相同的 API ).。
* [chain](https://github.com/codemodus/chain)  带有范围数据的处理程序包装器链接(基于网络/上下文的“中间件”)。
* [go-wrap](https://github.com/go-on/wrap)  net/http的小型中间件包。
* [gores](https://github.com/alioygur/gores)  处理HTML、JSON、XML等响应的Go包。对于RESTful api非常有用。
* [interpose](https://github.com/carbocation/interpose)  golang的极简网络/http中间件。
* [muxchain](https://github.com/stephens2424/muxchain)  用于net/http的轻量级中间件。
* [negroni](https://github.com/urfave/negroni)  符合语言习惯的 HTTP 中间件库。
* [render](https://github.com/unrolled/render)  Go package用于方便地呈现JSON、XML和HTML模板响应。
* [renderer](https://github.com/thedevsaddam/renderer)  简单、轻量级和更快的响应(JSON、JSONP、XML、YAML、HTML、文件)。
* [rye](https://github.com/InVisionApp/rye)  支持JWT、CORS、Statsd和Go 1.7上下文的小型Go中间件库(带有罐装中间件)。
* [stats](https://github.com/thoas/stats)  使用中间件来存储关于web应用程序的各种信息。

### 路由器

* [alien](https://github.com/gernest/alien)  轻量级和快速http路由器从外层空间。
* [httprouter](https://github.com/julienschmidt/httprouter)  高性能路由。使用这个库和标准http处理工具可以构建一个非常高性能大web框架。
* [xmux](https://github.com/rs/xmux)  高性能mux基于httprouter 'net/context`支持。
* [violetear](https://github.com/nbari/violetear)  HTTP路由器。
* [vestigo](https://github.com/husobee/vestigo)  高性能，独立，HTTP兼容的URL路由器的go web应用程序。
* [Siesta](https://github.com/VividCortex/siesta)  编写中间件和处理程序的可组合框架。
* [pure](https://github.com/go-playground/pure)  是一个轻量级HTTP路由器，它坚持net/ HTTP“实现”的std。
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)  一个非常快的Go (golang) HTTP路由器，支持正则表达式路由匹配。完全支持构建RESTful api。
* [mux](https://github.com/gorilla/mux)  强大的URL路由器和调度器为golang。
* [lars](https://github.com/go-playground/lars)  是一个轻量级、快速、可扩展、零分配的HTTP路由，用于创建定制化的框架。
* [httptreemux](https://github.com/dimfeld/httptreemux)  高速，灵活，基于树的 HTTP 路由。受到了 httprouter 的启发。
* [gowww/router](https://github.com/gowww/router)  超快的HTTP 路由，完全兼容 net/HTTP.Handler 接口.。
* [bellt](https://github.com/GuilhermeCaruso/bellt)  一个简单的Go HTTP路由器。
* [GoRouter](https://github.com/vardius/gorouter)  GoRouter 是一个服务器/API 微型框架、HTTP 请求路由 router, 数据分选器，提供了支持net/context的中间件。
* [goroute](https://github.com/goroute/route)  简单但功能强大的HTTP请求多路复用器。
* [Goji](https://github.com/goji/goji)  枸杞是一种简约的和灵活的与支持'net/context` HTTP请求多路复用器。
* [gocraft/web](https://github.com/gocraft/web)  Mux和中间件包在Go中。
* [FastRouter](https://github.com/razonyang/fastrouter)  一个快速，灵活的HTTP路由器写在Go。
* [fasthttprouter](https://github.com/buaazp/fasthttprouter)  高性能路由器分叉从`httprouter`。第一个路由器适合`fasthttp`。
* [chi](https://github.com/go-chi/chi)  小巧、快速、具有丰富表达力的 HTTP 路由，基于net/context.。
* [Bxog](https://github.com/claygod/Bxog)  简单和快速的HTTP路由器 Go 。它可以处理不同难度、长度和嵌套的路径。他还知道如何根据接收到的参数创建URL。
* [Bone](https://github.com/go-zoo/bone)  闪电快速HTTP多路复用器。
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter)  一个简单和快速的HTTP路由器 Go 。

## Windows

* [d3d9](https://github.com/gonutz/d3d9)   Direct3D9 的 Go 语言接口。
* [go-ole](https://github.com/go-ole/go-ole)  为 Go 语言实现的 Win32 OLE。
* [gosddl](https://github.com/MonaxGT/gosddl)  从SDDL-string到用户友好的JSON的转换器。SDDL由四个部分组成:所有者、主群、DACL、SACL。

## XML

*用于操作XML的库和工具。 (翻译出错了? 试试 [英文版](README_EN.md#xml) 吧~)*

* [XML-Comp](https://github.com/xml-comp/xml-comp)  简单的命令行XML比较器，生成文件夹、文件和标记的差异。
* [xml2map](https://github.com/sbabiv/xml2map)  XML来映射转换器编写的Golang。
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter)  基于libxml2的xmlwriter模块的过程性XML生成API。
* [xpath](https://github.com/antchfx/xpath)  Go的XPath包。
* [xquery](https://github.com/antchfx/xquery)  XQuery允许您使用XPath表达式从HTML/XML文档中提取数据。
* [zek](https://github.com/miku/zek)  从XML生成Go结构。

# 工具

* Go 软件和插件。 (翻译出错了? 试试 [英文版](README_EN.md#tools) 吧~)*

## 代码分析

* [apicompat](https://github.com/bradleyfalzon/apicompat)  检查 Go 项目最近的向下不兼容修改。
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  对未使用的常量、变量、函数和类型的代码进行检查。
* [unconvert](https://github.com/mdempsky/unconvert)  在源码中删除不必要的类型转换。
* [tickgit](https://github.com/augmentable-dev/tickgit)  CLI和go包用于显示代码注释TODOs(以任何语言)并应用“git blame”来识别作者。
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  在源码中寻找没有直接单元测试的函数和方法。
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  用于大量静态分析检查，您可能已经从 c# 的 ReSharper 等工具中习惯了这些检查。
* [php-parser](https://github.com/z7zmey/php-parser)  用 Go 编写的 PHP 解析器。
* [lint](https://github.com/surullabs/lint)  将 linters 作为测试的一部分。
* [gostatus](https://github.com/shurcooL/gostatus)  用于显示包含 Go 包的存储库的状态的命令行工具，。
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple 是 Go 源代码的linter，专门用于简化代码。
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  添加 zero 返回声明，以匹配 func 返回类型。
* [GoPlantUML](https://github.com/jfeliu007/goplantuml)  生成文本plantump类图的库和CLI，其中包含关于结构和接口的信息，以及它们之间的关系。
* [GoLint](https://github.com/golang/lint)  Go 源码的 linter。
* [dupl](https://github.com/mibk/dupl)  用于代码克隆检测的工具。
* [golines](https://github.com/segmentio/golines)  格式化程序，它自动缩短Go代码中的长行。
* [GolangCI](https://golangci.com/)  GolangCI 是一个针对 GitHub pull 请求的自动代码审查服务。服务是开源的，对开源项目是免费的。
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  来修复(添加，删除) Go 中自动导入的工具。
* [GoCover.io](http://gocover.io/)  GoCover.io 提供了任意 golang 包的代码覆盖率服务。
* [goast-viewer](https://github.com/yuroyoro/goast-viewer)  基于 Web 的 Golang AST 可视化工具。
* [go-outdated](https://github.com/firstrow/go-outdated)  显示过期包的终端应用。
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated)  找出项目中过期的依赖项。
* [go-critic](https://github.com/go-critic/go-critic)  源代码检查工具。
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch)  go-cleanarch 的创建是为了验证 Clean 体系结构规则，比如 Go 项目中的依赖关系。
* [go-checkstyle](https://github.com/qiniu/checkstyle)  checkstyle是一个类似于java checkstyle的检查工具。
* [gcvis](https://github.com/davecheney/gcvis)  实时可视化跟踪 GC 数据。
* [errcheck](https://github.com/kisielk/errcheck)  Errcheck是一个用于检查Go程序中未检查错误的程序。
* [validate](https://github.com/mccoyst/validate)  使用 tags 自动验证结构体字段。

## 编辑器插件

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  JetBrains IDEs 的 Go 插件。
* [go-language-server](https://github.com/theia-ide/go-language-server)  A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [go-mode](https://github.com/dominikh/go-mode.el)  在 GNU/Emacs 支持 GO。
* [go-plus](https://github.com/joefitzgerald/go-plus)  在 Atom 中添加自动完成，格式化，语法检查，高亮和审查。
* [gocode](https://github.com/nsf/gocode)  Autocompletion daemon for the Go programming language.
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  在 VS Code 中支持 Go 的基准分析。
* [GoSublime](https://github.com/DisposaBoy/GoSublime)  包含了可为文本编辑器 SublimeText 3 提供代码自动填充和其他类似IDE的功能的 Golang IDE 插件集合。
* [gounit-vim](https://github.com/hexdigest/gounit-vim)  基于函数或方法的签名生成Go测试的Vim插件。
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension)  在 Theia IDE中支持 Go。
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go)  在保存时突出显示语法错误的 Vim 插件。
* [vim-go](https://github.com/fatih/vim-go)  Go 开发会用到的 Vim 插件。
* [vscode-go](https://github.com/Microsoft/vscode-go)  Visual Studio代码的扩展(VS代码)，它提供了对Go语言的支持。
* [Watch](https://github.com/eaburns/Watch)  Runs a command in an acme win on file changes.

## Go 生成工具

* [generic](https://github.com/usk81/generic)  灵活的 Go 数据类型。
* [genny](https://github.com/cheekybits/genny)  优雅的 Go 泛型。
* [gocontracts](https://github.com/Parquery/gocontracts)  通过同步代码和文档来实现 design-by-contract 设计。
* [gonerics](http://github.com/bouk/gonerics)  Go中的易用的泛型。
* [gotests](https://github.com/cweill/gotests)  从源代码生成测试用例。
* [gounit](https://github.com/hexdigest/gounit)  使用您自己的模板生成Go测试用例。
* [hasgo](https://github.com/DylanMeeus/hasgo)  可生成用于切片的 Haskell。
* [re2dfa](https://github.com/opennota/re2dfa)  将正则表达式转换为有限状态机，并输出 Go 源代码。
* [TOML-to-Go](https://xuri.me/toml-to-go)  在浏览器中将 TOML 转换为 Go 类型。

## Go 工具

* [colorgo](https://github.com/songgao/colorgo)  将 go 命令包装成彩色的 go build 输出。
* [depth](https://github.com/KyleBanks/depth)  通过分析导入，将包依赖关系树可视化输出。
* [gb](https://getgb.io/)  一个基于项目的易用的构建工具。
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang)  一个[Yeoman](http://yeoman.io)生成器，用于启动新的 Go 项目。
* [gilbert](https://go-gilbert.github.io)  一个为 Go 项目而生的构建系统和任务运行器。
* [go-callvis](https://github.com/TrueFurby/go-callvis)  使用 dot format 可视化 Go 程序的调用图。
* [go-james](https://github.com/pieterclaerhout/go-james)  去项目骨架创造者，建立和测试您的项目没有手动设置。
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete)  Bash completion for go and wgo。
* [go-swagger](https://github.com/go-swagger/go-swagger)  基于 Go 的Swagger 2.0实现。
* [godbg](https://github.com/tylerwince/godbg)  实现了 Rusts 的 dbg! 宏，可以方便的在开发过程中快速、容易地调试。
* [gomodrun](https://github.com/dustinblackman/gomodrun/)  执行并缓存Go中包含的二进制文件的Go工具。国防部文件。
* [gothanks](https://github.com/psampaz/gothanks)  GoThanks自动帮你开始。mod github依赖，以这种方式发送一些爱给他们的维护者。
* [OctoLinker](https://github.com/OctoLinker/browser-extension)  借助的 OctoLinker 浏览器扩展，可以高效的地浏览  GitHub go文件。
* [richgo](https://github.com/kyoh86/richgo)  用文本装饰丰富 go test 的输出。
* [rts](https://github.com/galeone/rts)  从服务器响应生成Go结构。

## 软件包

*用Go编写的软件。 (翻译出错了? 试试 [英文版](README_EN.md#software-packages) 吧~)*

### DevOps 工具

* [aptly](https://github.com/smira/aptly)  Debian存储库管理工具。
* [Rodent](https://github.com/alouche/rodent)  管理Go版本、项目和跟踪依赖项。
* [kcli](https://github.com/cswank/kcli)  用于检查kafka主题/分区/消息的命令行工具。
* [kubernetes](https://github.com/kubernetes/kubernetes)  来自谷歌的容器集群管理器。
* [lstags](https://github.com/ivanilves/lstags)  提供了工具和API，可用来同步不同注册中心的Docker图像。
* [lwc](https://github.com/timdp/lwc)  UNIX wc命令的实时更新版本。
* [manssh](https://github.com/xwjdsh/manssh)  manssh是一个命令行工具，可以方便地管理ssh别名配置。
* [Moby](https://github.com/moby/moby)  Collaborative project for the container ecosystem to assemble container-based systems.
* [Mora](https://github.com/emicklei/mora)  用于访问 MongoDB 文档和元数据的 REST 服务器。
* [ostent](https://github.com/ostrost/ostent)  收集和显示系统指标，并可选 Graphite and/or fluxdb作为依赖。
* [Packer](https://github.com/mitchellh/packer)  用于从一个源配置为多个平台创建相同的机器图像。
* [Pewpew](https://github.com/bengadbois/pewpew)  灵活的 HTTP 命令行压测工具。
* [Pomerium](https://github.com/pomerium/pomerium)  Pomerium是一个可识别身份的访问代理。
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy)  具有GET、PUT和DELETE方法和身份验证(OpenID连接和基本身份验证)的S3代理。
* [jcli](https://github.com/jenkins-zh/jenkins-cli)  詹金斯CLI允许你管理你的詹金斯作为一个简单的方法。
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r)  小型实用程序/库，针对大型对象在Amazon S3中的高速传输进行了优化。
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli)  从命令行管理 BareMetal 服务器(与使用Docker一样容易)。
* [script](https://github.com/bitfield/script)  让DevOps编写类shell和系统管理任务变得更加容易。
* [sg](https://github.com/ChristopherRabotin/sg)  可测试一组HTTP极限(如ab)，可以在每个调用之间使用响应代码和数据，根据之前的响应来确定特定的服务器压力。
* [skm](https://github.com/TimothyYe/skm)  SKM是一个简单而强大的SSH密钥管理器，它可以帮助您轻松地管理多个SSH密钥!
* [StatusOK](https://github.com/sanathp/statusok)  监视您的网站和REST api。当服务器宕机或响应时间超过预期时，通过Slack、电子邮件获得通知。
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi)  Terraform provider plugin，它根据OpenAPI文档(以前称为swagger文件)在运行时动态配置自身，该文档包含了公开的api的定义。
* [traefik](https://github.com/containous/traefik)  反向代理和负载均衡器，支持多个后端。
* [uTask](https://github.com/ovh/utask)  对yaml中声明的业务流程进行建模和执行的自动化引擎。
* [Vegeta](https://github.com/tsenart/vegeta)  HTTP负载测试工具和库。超过9000 !
* [webhook](https://github.com/adnanh/webhook)  允许用户创建在服务器上执行命令的 HTTP hooks。
* [Wide](https://wide.b3log.org/login)  为使用 Golang 的团队提供基于 web 的 IDE。
* [kala](https://github.com/ajvb/kala)  简单、现代和高性能的作业调度程序。
* [Hey](https://github.com/rakyll/hey)  压力测试工具，可用来代替 ApacheBench (ab)。
* [aurora](https://github.com/xuri/aurora)  基于web的跨平台 Beanstalkd 队列服务器控制台。
* [fac](https://github.com/mkchoi212/fac)  修复 git 合并冲突。
* [awsenv](https://github.com/soniah/awsenv)  可以为 profile 加载Amazon (AWS)环境变量的轻量二进制文件。
* [Blast](https://github.com/dave/blast)  一个用于API负载测试和批处理作业的简单工具。
* [bombardier](https://github.com/codesenberg/bombardier)  快速跨平台 HTTP 基准测试工具。
* [bosun](https://github.com/bosun-monitor/bosun)  按照时间轴发出告警的框架。
* [DepCharge](https://github.com/centerorbit/depcharge)  Helps orchestrating the execution of commands across the many dependencies in larger projects.
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator)  一个go库和一个可执行文件，它使用各种输入通道生成有效的Dockerfiles。
* [dogo](https://github.com/liudng/dogo)  监视源文件中的更改并自动编译和运行(restart)。
* [drone-jenkins](https://github.com/appleboy/drone-jenkins)  使用二进制文件、docker或 Drone CI 来触发下游Jenkins作业。
* [drone-scp](https://github.com/appleboy/drone-scp)  通过 SSH 进行文件拷贝。其中 SSH 通过二进制文件、docker 或 Drone CI触发。
* [Dropship](https://github.com/chrismckenzie/dropship)  通过 cdn 部署代码的工具。
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy)  Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
* [gaia](https://github.com/gaia-pipeline/gaia)  可用于任何编程语言来构建强大的管道。
* [GVM](https://github.com/moovweb/gvm)  GVM 提供了一个接口来管理 Go 版本。
* [Gitea](https://github.com/go-gitea/gitea)  从 Gogs fork，完全由社区驱动。
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  将所有GitHub repositories、issues、milestones 和 labels 都迁移到 Gitea。
* [go-furnace](https://github.com/go-furnace/go-furnace)  用Go编写的托管解决方案，可轻松地在AWS、GCP或DigitalOcean上部署应用程序。
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate)  允许你的 Go应用程序 进行自我更新。
* [gobrew](https://github.com/cryptojuice/gobrew)  gobrew 允许您轻松地在 go 的多个版本之间切换。
* [godbg](https://github.com/sirnewton01/godbg)  基于 web 的 gdb 前端应用程序。
* [Gogs](https://gogs.io/)  自托管的Git服务。
* [gonative](https://github.com/inconshreveable/gonative)  用原生 Go 创建一个跨平台的 Go 工具链。
* [govvv](https://github.com/ahmetalpbalkan/govvv)  可轻松地添加版本信息到 Go 二进制文件。
* [gox](https://github.com/mitchellh/gox)  非常简单，没有多余的跨平台编译工具。
* [goxc](https://github.com/laher/goxc)  专注于跨平台编译和打包的 Go 构建工具。
* [grapes](https://github.com/yaronsumel/grapes)  旨在轻松地通过ssh分发命令的轻量级工具。
* [winrm-cli](https://github.com/masterzen/winrm-cli)  在Windows机器上远程执行命令的Cli工具。

### 其他软件

* [borg](https://github.com/crufter/borg)  基于终端的bash代码段搜索引擎。
* [restic](https://github.com/restic/restic)  消除重复项备份程序。
* [limetext](http://limetext.org/)  一个强大而优雅的文本编辑器。
* [LiteIDE](https://github.com/visualfc/liteide)  简单的、开源的、跨平台的Go IDE。
* [mockingjay](https://github.com/quii/mockingjay-server)  一份配置文件中便可伪造HTTP服务器与用户之间的行为。您还可以使服务器随机宕机，以帮助进行更实际的性能测试。
* [myLG](https://github.com/mehrdadrad/mylg)  命令行网络诊断工具。
* [naclpipe](https://github.com/unix4fun/naclpipe)  基于加密管的简单的NaCL EC25519工具。
* [nes](https://github.com/fogleman/nes)  任天堂娱乐系统(NES)模拟器。
* [orange-cat](https://github.com/noraesae/orange-cat)  用Go编写的Markdown预览器。
* [Orbit](https://github.com/gulien/orbit)  一个根据模板来运行命令和生成文件的简单小工具。
* [peg](https://github.com/pointlander/peg)  解析表达式语法，是Packrat解析器生成器的实现。
* [scc](https://github.com/boyter/scc)  一个非常快速准确的代码计数器，采用了复杂的计算和 COCOMO 预估。
* [Leaps](https://github.com/jeffail/leaps)  使用操作转换的成对编程服务。
* [Seaweed File System](https://github.com/chrislusf/seaweedfs)  快速、简单、可伸缩的分布式文件系统，采用了O(1)磁盘查找。
* [shell2http](https://github.com/msoap/shell2http)  通过http服务器执行shell命令(用于原型或远程控制)。
* [snap](https://github.com/intelsdi-x/snap)  强大的遥测框架。
* [Snitch](https://github.com/lucasgomide/snitch)  当有人通过 Tsuru 部署任何应用程序时，会通知您的团队以及其他工具。
* [Stack Up](https://github.com/pressly/sup)  Stack Up 是一个超级简单的部署工具 — 只面向Unix。
* [syncthing](https://syncthing.net/)  开放，分散的文件同步工具和协议。
* [term-quiz](https://github.com/crazcalm/term-quiz)  测试你的终端。
* [toxiproxy](https://github.com/shopify/toxiproxy)  为自动化测试模拟网络和系统条件的代理。
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow)  高性能、可伸缩和可靠的 IPFIX、sFlow和 Netflow 收集器。
* [lgo](https://github.com/yunabe/lgo)  与 Jupyter 可进行交互 Go 程序。它支持代码完成、代码检查以及与Go 100% 兼容性。
* [Juju](https://jujucharms.com/)  Cloud-agnostic的服务部署和编制 —— 支持EC2、Azure、Openstack、MAAS等。
* [boxed](https://github.com/tejo/boxed)  基于Dropbox的博客引擎。
* [Duplicacy](https://github.com/gilbertchen/duplicacy)  跨平台网络和云备份工具。
* [Cherry](https://github.com/rafael-santiago/cherry)  微型网络聊天服务器。
* [Circuit](https://github.com/gocircuit/circuit)  Circuit 是一个可编程平台即服务(PaaS)和/或基础设施即服务(IaaS)，用于管理、发现、同步和编排包含云应用程序的服务和主机。
* [Comcast](https://github.com/tylertreat/Comcast)  模拟坏的网络连接。
* [confd](https://github.com/kelseyhightower/confd)  使用 etcd 或 consul 的模板和数据管理本地应用程序配置文件。
* [croc](https://github.com/schollz/croc)  轻松、安全地将文件或文件夹从一台计算机发送到另一台计算机。
* [Docker](http://www.docker.com/)  面向开发人员和系统管理员的分布式应用程序的开放平台。
* [Documize](https://github.com/documize/community)  集成了SaaS工具数据的现代wiki软件。
* [dp](https://github.com/scryinfo/dp)  通过SDK与区块链进行数据交换，开发人员可以轻松访问DAPP开发。
* [drive](https://github.com/odeke-em/drive)  基于命令行的谷歌驱动器客户端。
* [gfile](https://github.com/Antonito/gfile)  通过WebRTC在两台计算机之间安全地传输文件，不需要任何第三方依赖。
* [joincap](https://github.com/assafmo/joincap)  用于合并多个pcap文件的命令行实用程序。
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store)  App that displays updates for the Go packages in your GOPATH.
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix)  视频流 torrent 客户端。
* [GoBoy](https://github.com/Humpheh/goboy)  用 Go 编写的任天堂Game Boy彩色模拟器。
* [gocc](https://github.com/goccmack/gocc)  Gocc是一个用Go编写的编译器工具包。
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip)  包含了 Go 使用手册文档的 Chrome 扩展。
* [GoLand](https://jetbrains.com/go)  功能齐全的跨平台 Go IDE。
* [Gor](https://github.com/buger/gor)  Http 流量复制工具，用于实时回放从生产环境到阶段/开发环境的流量。
* [hugo](http://gohugo.io/)  快速、现代的静态网站引擎。
* [ide](https://github.com/thestrukture/ide)  基于浏览器的IDE
* [ipe](https://github.com/dimiro1/ipe)  Open source Pusher server implementation compatible with Pusher client libraries written in GO.
* [wellington](https://github.com/wellington/wellington)  Sass 项目管理工具，使用sprite函数(如Compass)扩展语言。

# 资源

*在哪里可以找到新的Go库。 (翻译出错了? 试试 [英文版](README_EN.md#resources) 吧~)*

## 基准

* [autobench](https://github.com/davecheney/autobench)  用来来比较不同Go版本之间的性能的框架。
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app)  强大的HTTP基准测试工具，包含了Аb，Wrk，Siege工具。收集统计和各种参数指标，并比较相关结果。
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks)  Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)  HTTP请求路由器基准测试和比较。
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)  web框架基准测试。
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)  Go序列化方法的基准测试。
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Go语言常用基本操作的基准测试。
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark)  为流行的 Go 数据库/SQL实用程序收集基准测试。
* [gospeed](https://github.com/feyeleanor/GoSpeed)  计算语言结构的速度的微观基准测试。
* [kvbench](https://github.com/jimrobinson/kvbench)  K / V 类型数据库基准测试。
* [skynet](https://github.com/atemerev/skynet)  天网 1M 线程微基准测试。
* [speedtest-resize](https://github.com/fawick/speedtest-resize)  对比各种图像大小调整算法性能。

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
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly)  in Persian.
* [GoBooks](https://github.com/dariubs/GoBooks)  一份精选的 Go 书籍清单。
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack)  由 Maria Letta 提供的与 Gopher 有关的图片包，其中包含了插图,表情文字。
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector)  与 Go gopher 相关的媒介数据[。ai . svg)。
* [gopher-logos](https://github.com/GolangUA/gopher-logos)  可爱的 gopher 标识。
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gopherize.me](https://github.com/matryer/gopherize.me)  Gopherize自己。
* [gophers](https://github.com/ashleymcnamara/gophers)  阿什莉·麦克纳马拉的歌斐艺术品。
* [gophers](https://github.com/egonelbre/gophers)  Free gophers.
* [gophers](https://github.com/rogeralsing/gophers)  随机gopher图形。
* [gophers](https://github.com/sillecelik/go-gopher)  Gopher amigurumi玩具图案。

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
* [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
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

* [Awesome Go @LibHunt](https://go.libhunt.com)  属于你的 Go 工具箱。
* [godoc.org](https://godoc.org/)  开源 Go 包的文档。
* [json2go](https://m-zajac.github.io/json2go)  高级JSON去结构转换-在线工具。
* [gowalker.org](https://gowalker.org)   Go API 文档。
* [Gophercises](https://gophercises.com/)  为 Go 初学者提供免费的代码训练。
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  加入我们为Gophers设立的全新Slack社区([了解它是如何产生的](https://blog.gopheracademy.com/gophers-slack-community/))。
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  Google+社区 golang爱好者聚集地。
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go 邮件列表。
* [golang-graphics](https://github.com/mholt/golang-graphics)  收藏的 Go 图像，图形和艺术作品。
* [Golang News](https://golangnews.com)  关于 Go 编程的链接和新闻。
* [Golang Flow](https://golangflow.io)  发布更新、新闻、包等等。
* [Golang Developer Jobs](https://golangjob.xyz)  开发人员的工作专为Golang相关的角色。
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp)  专门收集需要帮助的Go项目，这是你开启开源之路的好去处。
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job)  Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [go.dev](https://go.dev/)  一个围棋开发者的中心。
* [Go Report Card](https://goreportcard.com)  为你的 Go 包生成一份报告单。
* [Go Projects](https://github.com/golang/go/wiki/Projects)  wiki上的 Go 社区项目列表。
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Forum](https://forum.golangbridge.org)  讨论 Go 的论坛。
* [Go Community on Hashnode](https://hashnode.com/n/go)  Hashnode上的Go社区。
* [Go Challenge](http://golang-challenge.org/)  通过解决问题并从 Go 专家那里得到反馈来学习 Go。
* [Go Blog](http://blog.golang.org)  官方 Go 博客。
* [CodinGame](https://www.codingame.com/)  以小游戏互动完成任务的形式来学习 Go。
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness)  其他 awesome 系列的列表。
* [justforfunc](https://www.youtube.com/c/justforfunc)  致力于传授 Go 编程语言技巧和技巧的Youtube 频道，由Francesc Campoy [@francesc](https://twitter.com/francesc)主办。
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  与 Go 有关的新闻。
* [studygolang](https://studygolang.com)  Go语言中文网
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  寻找最新的 Go库 的好地方。
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### 教程

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Go 初学者经常遇到的陷阱、误区、错误
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  教你如何用 Go 搭建一个电商平台 (包括demo)。
* [A Tour of Go](http://tour.golang.org/)  互动的 Go 之旅。
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang)  Golang电子书，主要讲述如何用 Golang 建立一个web应用程序。
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  如何缓存数据库的慢查询。
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  如何取消MySQL查询。
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book)  一本讲述如何用 Go 进行以太开发的小册。
* [Games With Go](http://gameswithgo.org/)  关于编程和游戏开发系列教学视频。
* [Go By Example](https://gobyexample.com/)  手把手教你 如何在 Go 应用程序中使用注释。
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)   Go's reference card。
* [Go database/sql tutorial](http://go-database-sql.org/)  数据库概论/ sql。
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  在你的移动设备上编辑你编辑和运行你的 Go 代码。
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns)  策划的清单去设计模式，食谱和习惯用法。
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers)  引入示例讲述 Golang 与Node.js在学习上的差异。
* [Golangbot](https://golangbot.com/learn-golang-series/)  Go 编程教程。
* [GolangCode](https://golangcode.com/)  收集代码片段和教程，以帮助处理日常问题。
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Go社区投票选举出来的最好的在线 Go 教程。
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  快速使用Godog —— 一个行为驱动开发的构建和测试应用程序框架。
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests)  学习使用测试驱动开发。
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  面向 Golang 初学者教程。
* [package main](https://www.youtube.com/packagemain)  关于 Go 编程的YouTube频道。
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera的专业学习可以从零开始。
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go)  由一个经验丰富的Go程序员群体编写的一系列Go学习范例。
* [Your basic Go](http://yourbasic.org/golang)  如何收集大量的教程。

