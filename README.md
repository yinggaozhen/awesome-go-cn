# Awesome Go

**此项目是 [awesome-go](https://awesome-go.com/) 中文版，最后一次同步时间 : 2019-06-27 17:28:31(每隔1h同步一次)**

**:star:项目地址 : [yinggaozhen/awesome-go-cn](https://github.com/yinggaozhen/awesome-go-cn):star:**
**在线访问

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) 为Awesome Go打赏~

精选了一系列很棒的Go框架、库和软件。灵感来自于[awesome-python](https://github.com/vinta/awesome-python)。

### 贡献

你可以快速浏览贡献者名单[contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md). 感觉所有为此项目付出的同学[contributors](https://github.com/avelino/awesome-go/graphs/contributors); 你真棒!

如果您在看到一个包或项目不再维护或不适合，请往awesome-go提交[Pull Request](https://github.com/avelino/awesome-go/pulls)，本项目定时(1小时)与英文文档同步。感谢!

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
    - [形式](#形式)
    - [功能](#功能)
    - [游戏开发](#游戏开发)
    - [一代和泛型](#一代和泛型)
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
    - [消息传递](#消息传递)
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
    - [查询语言](#查询语言)
    - [嵌入的资源](#嵌入的资源)
    - [科学与数据分析](#科学与数据分析)
    - [安全](#安全)
    - [序列化](#序列化)
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
            - [实际仿真中间件](#实际仿真中间件)
            - [用于创建HTTP中间件的库](#用于创建http中间件的库)
        - [路由器](#路由器)
    - [窗户](#窗户)
    - [XML](#xml)

- [工具](#工具)
    - [代码分析](#代码分析)
    - [编辑器插件](#编辑器插件)
    - [Go 生成工具](#go-生成工具)
    - [Go 工具](#go-工具)
    - [软件包](#软件包)
        - [DevOps的工具](#devops的工具)
        - [其他软件](#其他软件)

- [服务器应用程序](#服务器应用程序)

- [资源](#资源)
    - [基准](#基准)
    - [会议](#会议)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [聚会](#聚会)
    - [推特](#推特)
    - [网站](#网站)
        - [教程](#教程)

## 音频和音乐

*用于操作音频的库。*

* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) - **Star : 20** EasyMidi是一个简单可靠的库，用于处理标准midi文件(SMF)。
* [flac](https://github.com/eaburns/flac) - **Star : 84** 简单的本机Go FLAC解码器，将FLAC文件解码为字节片。
* [flac](https://github.com/mewkiz/flac) - **Star : 100** 本机Go FLAC编码器/解码器，支持FLAC流。
* [gaad](https://github.com/Comcast/gaad) - **Star : 55** 本机Go AAC位流解析器。
* [go-sox](https://github.com/krig/go-sox) - **Star : 92** 用于go的libsox绑定。
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) - **Star : 24** 用于go的libmediainfo绑定。
* [gosamplerate](https://github.com/dh1tw/gosamplerate) - **Star : 8** 用于go的libsamplerate绑定。
* [id3v2](https://github.com/bogem/id3v2) - **Star : 104** 快速稳定的ID3解析和编写Go库。
* [malgo](https://github.com/gen2brain/malgo) - **Star : 66** 迷你音频库。
* [minimp3](https://github.com/tosone/minimp3) - **Star : 25** 轻量级MP3解码器库。
* [mix](https://github.com/go-mix/mix) - **Star : 94** 基于顺序的Go-native音乐应用混音器。
* [mp3](https://github.com/tcolgate/mp3) - **Star : 89** 本机 Go MP3解码器。
* [music-theory](https://github.com/go-music-theory/music-theory) - **Star : 250** 音乐理论模型在围棋。
* [Oto](https://github.com/hajimehoshi/oto) - **Star : 381** 用于在多个平台上播放声音的低级库。
* [PortAudio](https://github.com/gordonklaus/portaudio) - **Star : 294** 用于PortAudio audio I/O库的Go绑定。
* [portmidi](https://github.com/rakyll/portmidi) - **Star : 203** PortMidi的Go绑定。
* [taglib](https://github.com/wtolson/go-taglib) - **Star : 65**  Go 绑定taglib。
* [vorbis](https://github.com/mccoyst/vorbis) - **Star : 22** “本机”Go Vorbis解码器(使用CGO，但没有依赖关系)。
* [waveform](https://github.com/mdlayher/waveform) - **Star : 244** Go软件包能够从音频流生成波形图像。

## 身份验证和OAuth

*用于实现验证方案的库。*

* [authboss](https://github.com/volatiletech/authboss) - **Star : 1886** web模块化认证系统。它试图删除尽可能多的样板文件和“硬东西”，以便每次在Go中启动一个新的web项目时，您都可以插入、配置并开始构建您的应用程序，而不必每次都构建一个身份验证系统。
* [branca](https://github.com/hako/branca) - **Star : 66** Golang实现Branca令牌。
* [casbin](https://github.com/hsluoyz/casbin) - 支持ACL、RBAC、ABAC等访问控制模型的授权库。
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) - **Star : 2** 提供cookie .txt文件格式的解析器。
* [go-jose](https://github.com/square/go-jose) - **Star : 1085** 相当完整地实现了JOSE工作组的JSON Web令牌、JSON Web签名和JSON Web加密规范。
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) - **Star : 1241** 独立的、符合规范的、用Golang编写的OAuth2服务器。
* [gologin](https://github.com/dghubble/gologin) - **Star : 1022** 用于使用OAuth1和OAuth2身份验证提供者登录的可链处理程序。
* [gorbac](https://github.com/mikespook/gorbac) - **Star : 894** 在Golang中提供一个轻量级的基于角色的访问控制(RBAC)实现。
* [goth](https://github.com/markbates/goth) - **Star : 2210** 提供使用OAuth和OAuth2的简单、干净和惯用的方法。处理多个提供程序的开箱即用。
* [httpauth](https://github.com/goji/httpauth) - **Star : 175** HTTP身份验证中间件。
* [jwt](https://github.com/robbert229/jwt) - **Star : 68** 干净易用的JSON Web令牌实现(JWT)。
* [jwt](https://github.com/pascaldekloe/jwt) - **Star : 74** 轻量级JSON Web令牌库。
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) - **Star : 151** JWT中间件为Golang http服务器提供了许多配置选项。
* [jwt-go](https://github.com/dgrijalva/jwt-go) - **Star : 5680** JSON Web令牌(JWT)的Golang实现。
* [loginsrv](https://github.com/tarent/loginsrv) - **Star : 789** JWT登录微服务带有可插入的后端，如OAuth2 (Github)、htpasswd、osiam。
* [oauth2](https://github.com/golang/oauth2) - **Star : 2308** goauth2的继任者。通用OAuth 2.0包，附带JWT、谷歌api、计算引擎和应用程序引擎支持。
* [osin](https://github.com/openshift/osin) - **Star : 1533** Golang OAuth2服务器库。
* [paseto](https://github.com/o1egl/paseto) - **Star : 220** Golang实现了与平台无关的安全令牌(PASETO)。
* [permissions2](https://github.com/xyproto/permissions2) - **Star : 346** 库，用于跟踪用户、登录状态和权限。使用安全cookie和bcrypt。
* [rbac](https://github.com/zpatrick/rbac) - **Star : 24** 用于Go应用程序的最小RBAC包。
* [securecookie](https://github.com/chmike/securecookie) - **Star : 31** 高效安全的cookie编码/解码。
* [session](https://github.com/icza/session) - **Star : 87** web服务器会话管理(包括支持谷歌应用程序引擎- GAE)。
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) - **Star : 8** 使用SessionGate Redis模块进行会话管理。
* [sessions](https://github.com/adam-hanna/sessions) - **Star : 45** 死简单，高性能，高定制会话服务的go http服务器。
* [signedvalue](https://github.com/sashka/signedvalue) - **Star : 7** 与[Tornado's]兼容的带签名和时间戳字符串(https://github.com/tornado oweb/tornado)' create_signed_value '， ' decode_signed_value '，因此' set_secure_cookie '和' get_secure_cookie '。

## Bot建设

*用于构建和使用机器人的库。*

* [go-chat-bot](https://github.com/go-chat-bot/bot) - **Star : 450** IRC, Slack和电报机器人用Go编写。
* [go-sarah](https://github.com/oklahomer/go-sarah) - **Star : 127** 框架为所需的聊天服务构建bot，包括LINE、Slack、Gitter等。
* [go-tgbot](https://github.com/olebedev/go-tgbot) - **Star : 82** 由swagger文件、基于会话的路由器和中间件生成的纯Golang Telegram Bot API包装器。
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) - **Star : 202** 一个用于加密货币交易所的基于控制台的交易机器人的golang实现。
* [govkbot](https://github.com/nikepan/govkbot) - **Star : 22** 简单的Go [VK](https://vk.com) bot库。
* [hanu](https://github.com/sbstjn/hanu) - **Star : 106** 编写Slack机器人的框架。
* [Kelp](https://github.com/stellar/kelp) - **Star : 143** 官方交易和做市机器人为[恒星](https://www.stellar.org/) DEX。开箱即用的作品，用Golang编写，兼容集中交易和定制交易策略。
* [margelet](https://github.com/zhulik/margelet) - **Star : 57** 构建电报机器人的框架。
* [micha](https://github.com/onrik/micha) - **Star : 9**  Go 图书馆获取Telegram bot api。
* [slacker](https://github.com/shomali11/slacker) - **Star : 297** 易于使用框架创建Slack机器人。
* [tbot](https://github.com/yanzay/tbot) - **Star : 211** 带有类似于net/http API的Telegram bot服务器。
* [telebot](https://github.com/tucnak/telebot) - **Star : 918** 用Go编写的Telegram bot框架。
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - 简单干净的Telegram bot客户端。
* [Tenyks](https://github.com/kyleterry/tenyks) - **Star : 167** 面向服务的IRC bot，使用Redis和JSON进行消息传递。

## 命令行

### 标准CLI

*用于构建标准或基本命令行应用程序的库。*

* [argparse](https://github.com/akamensky/argparse) - **Star : 99** 命令行参数分析器的灵感来自Python的argparse模块。
* [argv](https://github.com/cosiner/argv) - **Star : 16** 使用bash语法到库中分割命令行字符串作为参数数组。
* [cli](https://github.com/mkideal/cli) - **Star : 470** 功能丰富，易于使用的命令行包基于golang结构标签。
* [cli](https://github.com/teris-io/cli) - **Star : 56** 用于在Go中构建命令行接口的简单而完整的API。
* [cli-init](https://github.com/tcnksm/gcli) - **Star : 866** 开始构建Golang命令行应用程序的简单方法。
* [climax](http://github.com/tucnak/climax) - 另类CLI以“人面”，以精神 Go 指挥。
* [cmdr](https://github.com/hedzr/cmdr) - **Star : 7** 一个POSIX/GNU风格的、类似getopt的命令行UI Go库。
* [cobra](https://github.com/spf13/cobra) - **Star : 12498** 指挥官为现代Go CLI互动。
* [commandeer](https://github.com/jaffee/commandeer) - **Star : 77** 开发友好的CLI应用程序:基于struct字段和标记设置标志、默认值和用法。
* [complete](https://github.com/posener/complete) - **Star : 606** 在Go + Go命令bash completion中编写bash补全。
* [docopt.go](https://github.com/docopt/docopt.go) - **Star : 1131** 命令行参数解析器，让您微笑。
* [env](https://github.com/codingconcepts/env) - **Star : 41** 结构的基于标记的环境配置。
* [flag](https://github.com/cosiner/flag) - **Star : 100** 简单但功能强大的命令行选项解析库，用于支持Go子命令。
* [flaggy](https://github.com/integrii/flaggy) - **Star : 438** 一个健壮的、惯用的标志包，具有出色的子命令支持。
* [flagvar](https://github.com/sgreben/flagvar) - **Star : 31** 一组Go标准“flag”包的标志参数类型。
* [go-arg](https://github.com/alexflint/go-arg) - **Star : 641** 在Go中基于结构的参数解析。
* [go-commander](https://github.com/yitsushi/go-commander) - **Star : 14** 使用库简化CLI工作流。
* [go-flags](https://github.com/jessevdk/go-flags) - **Star : 1486**  Go 命令行选项解析器。
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) - **Star : 5**  Go 选择解析器启发Perl的灵活性的GetOpt::长。
* [gocmd](https://github.com/devfacet/gocmd) - **Star : 33** 使用库构建命令行应用程序。
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - 具有自动配置和依赖注入的cli应用程序框架。
* [job](https://github.com/liujianping/job) - **Star : 41** 工作，把你的短期指令当作长期任务。
* [kingpin](https://github.com/alecthomas/kingpin) - **Star : 2488** 支持子命令的命令行和标志解析器。
* [liner](https://github.com/peterh/liner) - **Star : 568** 为命令行接口使用类似readline的库。
* [mitchellh/cli](https://github.com/mitchellh/cli) - **Star : 981** 用于实现命令行接口的Go库。
* [mow.cli](https://github.com/jawher/mow.cli) - **Star : 619** Go库用于构建具有复杂标志和参数解析和验证的CLI应用程序。
* [pflag](https://github.com/spf13/pflag) - **Star : 723** Drop-in替换Go的标志包，实现POSIX/GNU-style——标志。
* [readline](https://github.com/chzyer/readline) - **Star : 1361** 纯golang实现，在MIT许可下提供了GNU-Readline的大部分特性。
* [sand](https://github.com/Zaba505/sand) - **Star : 5** 用于创建解释器等的简单API。
* [sflags](https://github.com/octago/sflags) - **Star : 84** 基于结构的旗子生成器，用于旗子、urfave/cli、pflag、cobra、kingpin和其他库。
* [strumt](https://github.com/antham/strumt) - **Star : 27** 库创建提示链。
* [ukautz/clif](https://github.com/ukautz/clif) - **Star : 98** 小命令行接口框架。
* [urfave/cli](https://github.com/urfave/cli) - **Star : 11018** 在Go中构建命令行应用程序的简单、快速和有趣的包(以前是codegangsta/cli)。
* [wlog](https://github.com/dixonwille/wlog) - **Star : 33** 支持跨平台颜色和并发的简单日志记录接口。
* [wmenu](https://github.com/dixonwille/wmenu) - **Star : 80** 易于使用的菜单结构为cli应用程序，提示用户作出选择。

### 先进的控制台用户界面

*用于构建控制台应用程序和控制台用户界面的库。*

* [asciigraph](https://github.com/guptarohit/asciigraph) - **Star : 1105**  Go 包使轻量级ASCII线图╭┈╯在命令行应用程序中没有其他依赖项。
* [aurora](https://github.com/logrusorgru/aurora) - **Star : 591** 支持fmt.Printf/Sprintf的ANSI终端颜色。
* [cfmt](https://github.com/mingrammer/cfmt) - **Star : 67** 上下文fmt的灵感来自于引导颜色类。
* [chalk](https://github.com/ttacon/chalk) - **Star : 302** 直观的软件包美化终端/控制台输出。
* [color](https://github.com/fatih/color) - **Star : 2975** 多功能包装，彩色终端输出。
* [colourize](https://github.com/TreyBastian/colourize) - **Star : 15**  Go 图书馆的ANSI彩色文本在终端。
* [ctc](https://github.com/wzshiming/ctc) - **Star : 9** 非侵入性跨平台终端颜色库不需要修改打印方法。
* [go-ataman](https://github.com/workanator/go-ataman) - **Star : 8**  Go 库渲染ANSI彩色文本模板在终端。
* [go-colorable](https://github.com/mattn/go-colorable) - **Star : 366** 适用于windows的可着色写入器。
* [go-colortext](https://github.com/daviddengcn/go-colortext) - **Star : 198** 在终端中使用颜色输出库。
* [go-isatty](https://github.com/mattn/go-isatty) - **Star : 331** isatty golang。
* [go-prompt](https://github.com/c-bata/go-prompt) - **Star : 2259** 库，受[python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit)的启发，构建一个强大的交互式提示。
* [gocui](https://github.com/jroimartin/gocui) - **Star : 4279** 旨在创建控制台用户界面的极简Go库。
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - 样式文本终端。
* [gookit/color](https://github.com/gookit/color) - **Star : 177** 终端显色工具库，支持16种颜色，256种颜色，RGB显色输出，兼容Windows。
* [mpb](https://github.com/vbauerster/mpb) - **Star : 496** 终端应用程序的多进度条。
* [progressbar](https://github.com/schollz/progressbar) - **Star : 549** 基本线程安全的进度条，在每个操作系统工作。
* [simpletable](https://github.com/alexeyco/simpletable) - **Star : 153** 简单的表格在终端与围棋。
* [tabby](https://github.com/cheynewallace/tabby) - **Star : 243** 一个超级简单的Golang表的小库。
* [tabular](https://github.com/InVisionApp/tabular) - **Star : 29** 从命令行实用程序中打印ASCII表，而不需要向API传递大量数据。
* [termbox-go](https://github.com/nsf/termbox-go) - **Star : 3429** Termbox是一个用于创建跨平台的基于文本的接口的库。
* [termdash](https://github.com/mum4k/termdash) - **Star : 766** 基于**termbox-go**的Go terminal dashboard，灵感来自[termui](https://github.com/gizak/termui)。
* [termtables](https://github.com/apcera/termtables) - **Star : 212** 使用Ruby库[terminal-tables](https://github.com/tj/terminal-table)的端口生成简单的ASCII表，并提供标记和HTML输出。
* [termui](https://github.com/gizak/termui) - **Star : 8775** 基于**termbox-go**的Go terminal dashboard，灵感来自[blessed-contrib](https://github.com/yaronn/blessed-contrib)。
* [uilive](https://github.com/gosuri/uilive) - **Star : 813** 用于实时更新终端输出的库。
* [uiprogress](https://github.com/gosuri/uiprogress) - **Star : 1517** 灵活的库在终端应用程序中呈现进度条。
* [uitable](https://github.com/gosuri/uitable) - **Star : 491** 库，用于提高使用表格数据的终端应用程序的可读性。

## 配置

*配置解析的库。*

* [config](https://github.com/JeremyLoy/config) - **Star : 184** 云本地应用程序配置。仅用两行代码将ENV绑定到结构体。
* [config](https://github.com/olebedev/config) - **Star : 209** 带有环境变量和标记解析的JSON或YAML配置包装器。
* [configure](https://github.com/paked/configure) - **Star : 48** 通过多个源提供配置，包括JSON、标志和环境变量。
* [confita](https://github.com/heetch/confita) - **Star : 236** 从多个后端级联加载配置到结构中。
* [conflate](https://github.com/the4thamigo-uk/conflate) - **Star : 8** 库/工具来合并来自任意url的多个JSON/YAML/TOML文件、针对JSON模式的验证以及模式中定义的默认值的应用程序。
* [env](https://github.com/caarlos0/env) - **Star : 831** 解析环境变量以获得struct(默认值)。
* [envcfg](https://github.com/tomazk/envcfg) - **Star : 90** 将环境变量解组到struct。
* [envconf](https://github.com/ian-kent/envconf) - **Star : 7** 从环境配置。
* [envconfig](https://github.com/vrischmann/envconfig) - **Star : 143** 从环境变量中读取配置。
* [envh](https://github.com/antham/envh) - **Star : 93** 帮助管理环境变量。
* [gcfg](https://github.com/go-gcfg/gcfg) - **Star : 115** 将ini风格的配置文件读入Go结构;支持用户定义的类型和子节。
* [go-up](https://github.com/ufoscout/go-up) - **Star : 24** 一个简单的配置库，具有递归占位符解析，没有任何神奇之处。
* [goConfig](https://github.com/crgimenes/goConfig) - **Star : 102** 将结构体解析为输入，并用来自命令行、环境变量和配置文件的参数填充该结构体的字段。
* [godotenv](https://github.com/joho/godotenv) - **Star : 2039** Ruby的dotenv库的Go端口(从' .env '加载环境变量)。
* [gofigure](https://github.com/ian-kent/gofigure) - **Star : 57** 应用程序配置变得简单。
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - 模块化的JSON配置。保持配置结构及其配置的代码，并将解析委托给子模块，而不牺牲配置的完整序列化。
* [gookit/config](https://github.com/gookit/config) - **Star : 69** 应用程序配置管理(负载、获取、设置)。支持JSON, YAML, TOML, INI, HCL。多文件加载，数据覆盖合并。
* [hjson](https://github.com/hjson/hjson-go) - **Star : 172** 人类JSON，一种用于人类的配置文件格式。轻松的语法，更少的错误，更多的注释。
* [ingo](https://github.com/schachmat/ingo) - **Star : 23** 标记保存在类ini的配置文件中。
* [ini](https://github.com/go-ini/ini) - **Star : 1510**  Go 软件包读和写INI文件。
* [joshbetz/config](https://github.com/joshbetz/config) - **Star : 194** 用于Go的小型配置库，它解析环境变量、JSON文件，并在SIGHUP时自动重新加载。
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - **Star : 2333** 使用库管理来自环境变量的配置数据。
* [konfig](https://github.com/lalamove/konfig) - **Star : 506** 面向分布式处理时代的可组合、可观察和高性能配置处理。
* [mini](https://github.com/sasbury/mini) - **Star : 19** 用于解析ini风格的配置文件的Golang包。
* [sprbox](https://github.com/oblq/sprbox) - **Star : 3** 支持构建环境的工具箱工厂和不可知的配置解析器(YAML、TOML、JSON和环境vars)。
* [store](https://github.com/tucnak/store) - **Star : 241** Go的轻量级配置管理器。
* [viper](https://github.com/spf13/viper) - **Star : 8896**  Go 配置尖牙。
* [xdg](https://github.com/OpenPeeDeeP/xdg) - **Star : 31** 遵循[XDG标准]的跨平台包(https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)。

## 持续集成

*用于帮助进行持续集成的工具。*

* [drone](https://github.com/drone/drone) - **Star : 18624** Drone是一个基于Docker的持续集成平台，用Go编写。
* [duci](https://github.com/duck8823/duci) - **Star : 40** 简单的ci服务器不需要特定于域的语言。
* [gomason](https://github.com/nikogura/gomason) - **Star : 27** 在一个干净的工作区中测试、构建、签名和发布go二进制文件。
* [goveralls](https://github.com/mattn/goveralls) - **Star : 567** 把工作服整合起来。io连续代码覆盖跟踪系统。
* [overalls](https://github.com/go-playground/overalls) - **Star : 96** 多包go项目涵盖的工具，如政府。
* [roveralls](https://github.com/LawrenceWoodman/roveralls) - **Star : 12** 递归覆盖测试工具。

## CSS预处理器

*用于预处理CSS文件的库。*

* [gcss](https://github.com/yosssi/gcss) - **Star : 421** 纯Go CSS预处理器。
* [go-libsass](https://github.com/wellington/go-libsass) - **Star : 128**  Go 包装到100% Sass兼容的libsass项目。

## 数据结构

*通用的数据结构和算法在围棋。*
* [algorithms](https://github.com/shady831213/algorithms) - **Star : 227** 算法和数据结构。clr的研究。
* [binpacker](https://github.com/zhuangsirui/binpacker) - **Star : 119** 二进制封隔器和解包器帮助用户构建自定义二进制流。
* [bit](https://github.com/yourbasic/bit) - **Star : 50** 戈朗集数据结构与奖金比特玩弄功能。
* [bitset](https://github.com/willf/bitset) - **Star : 474** 打包实现位集。
* [bloom](https://github.com/zhenjl/bloom) - 在Go中实现了Bloom过滤器。
* [bloom](https://github.com/yourbasic/bloom) - **Star : 36** Golang Bloom过滤器的实现。
* [boomfilters](https://github.com/tylertreat/BoomFilters) - **Star : 1119** 处理连续无界流的概率数据结构。
* [concurrent-writer](https://github.com/free/concurrent-writer) - **Star : 21** 高度并发的drop-in替换“bufio.Writer”。
* [conjungo](https://github.com/InVisionApp/conjungo) - **Star : 72** 一个小型、强大和灵活的合并库。
* [count-min-log](https://github.com/seiflotfy/count-min-log) - **Star : 41** Go实现Count-Min- log sketch:使用近似计数器进行近似计数(类似Count-Min sketch，但使用更少内存)。
* [crunch](https://github.com/superwhiskers/crunch) - **Star : 18** 打包实现缓冲区，以便轻松处理各种数据类型。
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - **Star : 491** 布谷鸟过滤器:一个很好的替代计数布卢姆过滤器实现在围棋。
* [deque](https://github.com/edwingeng/deque) - **Star : 3** 高度优化的双端队列。
* [deque](https://github.com/gammazero/deque) - **Star : 58** 快速环缓冲区deque(双端队列)。
* [dict](https://github.com/srfrog/dict) - **Star : 6** 类python字典(dict)。
* [encoding](https://github.com/zhenjl/encoding) - 用于Go的整数压缩库。
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - **Star : 82** Go实现自适应基数树。
* [go-datastructures](https://github.com/Workiva/go-datastructures) - **Star : 5031** 有用的、高性能的和线程安全的数据结构的集合。
* [go-ef](https://github.com/amallia/go-ef) - **Star : 9** Elias-Fano编码的Go实现。
* [go-geoindex](https://github.com/hailocab/go-geoindex) - **Star : 307** 内存中的地理指数。
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - **Star : 30** 快速内存键:值存储/缓存库。指针缓存。
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) - 区域四叉树具有高效的点定位和邻域查找功能。
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) - **Star : 22** 并行FIFO队列。
* [gods](https://github.com/emirpasic/gods) - **Star : 5997** 数据结构。容器、集合、列表、堆栈、地图、BidiMaps、树、HashSet等。
* [golang-set](https://github.com/deckarep/golang-set) - **Star : 1112** 线程安全和非线程安全的高性能集。
* [goset](https://github.com/zoumo/goset) - **Star : 15** 一个有用的Go集合实现。
* [goskiplist](https://github.com/ryszard/goskiplist) - **Star : 190** 跳转列表实现在Go。
* [gota](https://github.com/kniren/gota) - 实现Go的数据流、系列和数据争用方法。
* [hide](https://github.com/emvi/hide) - **Star : 6** 使用编组到/从散列的ID类型，以防止向客户机发送ID。
* [hilbert](https://github.com/google/hilbert) - **Star : 176** Go软件包用于将值映射到和从空间填充曲线(如Hilbert和Peano曲线)。
* [hyperloglog](https://github.com/axiomhq/hyperloglog) - **Star : 650** HyperLogLog实现稀疏，LogLog-Beta偏差校正和尾部削减空间。
* [levenshtein](https://github.com/agext/levenshtein) - **Star : 31** Levenshtein距离和相似性度量，具有可定制的编辑成本和类似于winkler的通用前缀奖励。
* [levenshtein](https://github.com/agnivade/levenshtein) - **Star : 52** 实现在Go中计算levenshtein距离。
* [mafsa](https://github.com/smartystreets/mafsa) - **Star : 272** 用最少的完美哈希实现MA-FSA。
* [merkletree](https://github.com/cbergoon/merkletree) - **Star : 141** merkle树的实现提供了对数据结构内容的有效和安全的验证。
* [mspm](https://github.com/BlackRabbitt/mspm) - **Star : 6** 用于信息检索的多字符串模式匹配算法。
* [null](https://github.com/emvi/null) - **Star : 4** 可以对JSON进行编组/反编组的可为空的Go类型。
* [parsefields](https://github.com/MonaxGT/parsefields) - **Star : 2** 用于解析类似json的日志的工具，用于收集惟一的字段和事件。
* [pipeline](https://github.com/hyfather/pipeline) - **Star : 14** 具有扇入和扇出的管道的实现。
* [ring](https://github.com/TheTannerRyan/ring) - **Star : 85** Go实现了高性能、线程安全的bloom过滤器。
* [roaring](https://github.com/RoaringBitmap/roaring) - **Star : 637**  Go 包实现压缩位集。
* [set](https://github.com/StudioSol/set) - **Star : 5** 使用LinkedHashMap在Go中实现简单的set数据结构。
* [skiplist](https://github.com/MauriceGit/skiplist) - **Star : 95** 非常快的Go Skiplist实现。
* [skiplist](https://github.com/gansidui/skiplist) - **Star : 61** 在Go中的Skiplist实现。
* [timedmap](https://github.com/zekroTJA/timedmap) - 映射即将到期的键值对。
* [treap](https://github.com/perdata/treap) - **Star : 6** 使用树堆进行持久、快速有序的映射。
* [trie](https://github.com/derekparker/trie) - **Star : 404** 在Go中实现Trie。
* [ttlcache](https://github.com/diegobernardes/ttlcache) - 内存LRU字符串接口{}映射与过期为golang。
* [typ](https://github.com/gurukami/typ) - **Star : 8** 空类型、安全原语类型转换和从复杂结构中获取值。
* [willf/bloom](https://github.com/willf/bloom) - **Star : 642** 打包实现Bloom过滤器。

## 数据库

*数据库在Go中实现。*

* [badger](https://github.com/dgraph-io/badger) - **Star : 6039** 快速键值存储在Go中。
* [bcache](https://github.com/iwanbk/bcache) - **Star : 25** 最终一致的分布式内存缓存Go库。
* [BigCache](https://github.com/allegro/bigcache) - **Star : 2355** 有效的键/值缓存为千兆字节的数据。
* [bolt](https://github.com/boltdb/bolt) - **Star : 9841** Go的低级键/值数据库。
* [buntdb](https://github.com/tidwall/buntdb) - **Star : 2411** 快速，可嵌入，在内存键/值数据库与自定义索引和空间支持 Go 。
* [cache](https://github.com/akyoto/cache) - **Star : 10** 内存键:带过期时间的值存储，0个依赖项，<100 LoC, 100%覆盖率。
* [cache2go](https://github.com/muesli/cache2go) - **Star : 922** 内存内键:值缓存，支持基于超时的自动失效。
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - **Star : 29** BigCache支持集群和单独的条目过期。
* [cockroach](https://github.com/cockroachdb/cockroach) - **Star : 16510** 可伸缩、地理复制、事务性数据存储。
* [couchcache](https://github.com/codingsince1985/couchcache) - **Star : 40** 由Couchbase服务器支持的RESTful缓存微服务。
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - **Star : 1414** CovenantSQL是区块链上的一个SQL数据库。
* [dgraph](https://github.com/dgraph-io/dgraph) - **Star : 9847** 可伸缩、分布式、低延迟、高吞吐量的图形数据库。
* [diskv](https://github.com/peterbourgon/diskv) - **Star : 732** 自定义磁盘支持的键值存储。
* [eliasdb](https://github.com/krotik/eliasdb) - **Star : 531** 不依赖，事务图数据库与REST API，短语搜索和sql类似的查询语言。
* [fastcache](https://github.com/VictoriaMetrics/fastcache) - **Star : 461** 快速线程安全的内存缓存为大量的条目。最大限度地减少GC开销。
* [GCache](https://github.com/bluele/gcache) - **Star : 855** 支持过期缓存、LFU、LRU和ARC的缓存库。
* [go-cache](https://github.com/pmylund/go-cache) - 内存键:Go的值存储/缓存库(类似于Memcached)，适用于单机应用程序。
* [goleveldb](https://github.com/syndtr/goleveldb) - **Star : 3089** 在Go中实现[LevelDB](https://github.com/google/leveldb) key/value数据库。
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) - **Star : 7** Gorocksdb是用Go编写的[RocksDB](https://rocksdb.org)的包装器。
* [groupcache](https://github.com/golang/groupcache) - **Star : 7522** Groupcache是一个缓存和缓存填充库，在许多情况下，它是memcached的替代品。
* [influxdb](https://github.com/influxdb/influxdb) - 可伸缩的数据存储，用于度量、事件和实时分析。
* [ledisdb](https://github.com/siddontang/ledisdb) - **Star : 3032** Ledisdb是一种高性能的NoSQL，类似于基于LevelDB的Redis。
* [levigo](https://github.com/jmhodges/levigo) - **Star : 363** Levigo是LevelDB的Go包装器。
* [moss](https://github.com/couchbase/moss) - **Star : 714** Moss是一个用100% Go编写的简单LSM键值存储引擎。
* [nutsdb](https://github.com/xujiajun/nutsdb) - **Star : 839** Nutsdb是一个用纯Go编写的简单、快速、可嵌入、持久的键/值存储。它支持完全序列化的事务和许多数据结构，如列表、集合、排序集。
* [piladb](https://github.com/fern4lvarez/piladb) - **Star : 169** 基于堆栈数据结构的轻量级RESTful数据库引擎。
* [prometheus](https://github.com/prometheus/prometheus) - **Star : 24775** 监控系统和时序数据库。
* [pudge](https://github.com/recoilme/pudge) - **Star : 212** 快速和简单的键/值存储使用Go的标准库编写。
* [rqlite](https://github.com/rqlite/rqlite) - **Star : 4622** 基于SQLite的轻量级分布式关系数据库。
* [Scribble](https://github.com/nanobox-io/golang-scribble) - **Star : 54** 小平面文件JSON存储。
* [slowpoke](https://github.com/recoilme/slowpoke) - **Star : 84** 具有持久性的键值存储。
* [tempdb](https://github.com/rafaeljesus/tempdb) - **Star : 13** 临时项的键值存储。
* [tidb](https://github.com/pingcap/tidb) - **Star : 19334** TiDB是一个分布式SQL数据库。灵感来自谷歌F1的设计。
* [tiedot](https://github.com/HouzuoGuo/tiedot) - **Star : 2353** 您的NoSQL数据库由Golang提供支持。
* [Vasto](https://github.com/chrislusf/vasto) - **Star : 141** 分布式高性能键值存储。在磁盘上。最终一致。哈哈。能够在不中断服务的情况下增长或收缩。
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - **Star : 837** 快速，资源有效和可伸缩的开放源码时间序列数据库。可作为普罗米修斯的长期远程存储。支持PromQL。

*数据库模式迁移。*

* [avro](https://github.com/khezen/avro) - **Star : 3** 发现SQL模式并将其转换为AVRO模式。查询SQL记录到AVRO字节。
* [darwin](https://github.com/GuiaBolso/darwin) - **Star : 84** Go的数据库模式演化库。
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) - **Star : 19** Django风格的fixture，用于Golang出色的内置数据库/sql库。
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - **Star : 24** 一个帮助用Go -pg/pg编写迁移的Go包。
* [gondolier](https://github.com/emvi/gondolier) - **Star : 26** 使用结构修饰器的数据库迁移库。
* [goose](https://github.com/steinbacher/goose) - **Star : 116** 数据库迁移工具。您可以通过创建增量SQL或Go脚本来管理数据库的演进。
* [gormigrate](https://github.com/go-gormigrate/gormigrate) - **Star : 313** Gorm ORM的数据库模式迁移帮助程序。
* [migrate](https://github.com/golang-migrate/migrate) - **Star : 2389** 数据库迁移。CLI和Golang库。
* [migrator](https://github.com/lopezator/migrator) - **Star : 26** 死简单 Go 数据库迁移库。
* [pravasan](https://github.com/pravasan/pravasan) - **Star : 24** 简单的迁移工具-目前为MySQL，但计划很快支持Postgres, SQLite, MongoDB等。
* [soda](https://github.com/gobuffalo/pop/tree/master/soda) - 数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [sql-migrate](https://github.com/rubenv/sql-migrate) - **Star : 1362** 数据库迁移工具。允许使用go-bindata将迁移嵌入到应用程序中。

*数据库工具。*

* [chproxy](https://github.com/Vertamedia/chproxy) - **Star : 290** ClickHouse数据库的HTTP代理。
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - **Star : 126** 收集小instert并向ClickHouse服务器发送大请求。
* [datagen](https://github.com/codingconcepts/datagen) - **Star : 5** 一个快速的数据生成器，多表感知和支持多行DML。
* [dbbench](https://github.com/sj14/dbbench) - **Star : 30** 数据库基准测试工具，支持多个数据库和脚本。
* [go-mysql](https://github.com/siddontang/go-mysql) - **Star : 1789**  Go 工具集处理MySQL协议和复制。
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - **Star : 2276** 自动将MySQL数据同步到Elasticsearch中。
* [kingshard](https://github.com/flike/kingshard) - **Star : 4513** kingshard是Golang支持的MySQL高性能代理。
* [myreplication](https://github.com/2tvenom/myreplication) - **Star : 141** MySql二进制日志复制监听器。支持基于语句和行的复制。
* [octillery](https://github.com/knocknote/octillery) - **Star : 46** Go package for sharding数据库(支持每个ORM或原始SQL)。
* [orchestrator](https://github.com/github/orchestrator) - **Star : 2936** MySQL复制拓扑管理器和可视化。
* [pgweb](https://github.com/sosedoff/pgweb) - **Star : 5932** 基于web的PostgreSQL数据库浏览器。
* [prep](https://github.com/hexdigest/prep) - **Star : 24** 在不更改代码的情况下使用准备好的SQL语句。
* [pREST](https://github.com/nuveo/prest) - 从任何PostgreSQL数据库中提供RESTful API。
* [rwdb](https://github.com/andizzle/rwdb) - **Star : 10** rwdb为多个数据库服务器的设置提供读取副本功能。
* [vitess](https://github.com/youtube/vitess) - vitess提供了服务器和工具，这些工具可以为大规模web服务扩展MySQL数据库提供便利。

*SQL查询生成器，用于构建和使用SQL的库。*

* [Dotsql](https://github.com/gchaincl/dotsql) - **Star : 433** Go library帮助您将sql文件保存在一个地方，并轻松地使用它们。
* [gendry](https://github.com/didi/gendry) - **Star : 692** 非入侵的SQL builder和强大的数据绑定器。
* [godbal](https://github.com/xujiajun/godbal) - **Star : 50** 数据库抽象层(dbal)。支持SQL builder，轻松获取结果。
* [goqu](https://github.com/doug-martin/goqu) - **Star : 519** 惯用SQL生成器和查询库。
* [igor](https://github.com/galeone/igor) - **Star : 77** PostgreSQL的抽象层，支持高级功能并使用类似gorm的语法。
* [ormlite](https://github.com/pupizoid/ormlite) - 轻量级包，包含一些类似orm的特性和sqlite数据库的帮助程序。
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - **Star : 433** 强大的数据检索方法以及与db无关的查询构建功能。
* [scaneo](https://github.com/variadico/scaneo) - **Star : 149** 生成Go代码将数据库行转换为任意结构。
* [sqrl](https://github.com/elgris/sqrl) - **Star : 171** SQL查询生成器，提高了性能的Squirrel fork。
* [Squalus](https://gitlab.com/qosenergy/squalus) - Go SQL包上的薄层，使得执行查询更加容易。
* [Squirrel](https://github.com/Masterminds/squirrel) - **Star : 2178** 帮助您构建SQL查询的Go库。
* [xo](https://github.com/knq/xo) - 基于支持PostgreSQL、MySQL、SQLite、Oracle和Microsoft SQL Server的现有模式定义或自定义查询，为数据库生成惯用的Go代码。

## 数据库驱动程序

*用于连接和操作数据库的库。*

* Relational Databases
    * [avatica](https://github.com/apache/calcite-avatica-go) - **Star : 31** 用于数据库/ SQL的Apache Avatica/Phoenix SQL驱动程序。
    * [bgc](https://github.com/viant/bgc) - **Star : 12** 用于go的BigQuery的数据存储连接。
    * [firebirdsql](https://github.com/nakagami/firebirdsql) - **Star : 102** Firebird RDBMS SQL驱动程序。
    * [go-adodb](https://github.com/mattn/go-adodb) - **Star : 90** 使用数据库/sql的Microsoft ActiveX对象数据库驱动程序。
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - **Star : 990** 微软MSSQL驱动程序。
    * [go-oci8](https://github.com/mattn/go-oci8) - **Star : 395** 使用数据库/sql的Oracle go驱动程序。
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - **Star : 7824** MySQL驱动程序。
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - **Star : 3338** 使用数据库/sql的SQLite3驱动程序。
    * [gofreetds](https://github.com/minus5/gofreetds) - **Star : 90** 微软该司机。 Go 包装[FreeTDS](http://www.freetds.org)。
    * [goracle](https://github.com/go-goracle/goracle) - **Star : 224** Oracle驱动程序for Go，使用ODPI-C驱动程序。
    * [pgx](https://github.com/jackc/pgx) - **Star : 1849** PostgreSQL驱动程序支持数据库/sql公开之外的特性。
    * [pq](https://github.com/lib/pq) - **Star : 5052** 用于数据库/sql的纯Go Postgres驱动程序。

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - **Star : 300** Aerospike客户端在Go语言中。
    * [arangolite](https://github.com/solher/arangolite) - **Star : 65** 轻量级golang驱动程序为ArangoDB。
    * [asc](https://github.com/viant/asc) - **Star : 4** Aerospike for go的数据存储连接。
    * [dynago](https://github.com/underarmour/dynago) - **Star : 64** Dynago是DynamoDB客户端最不意外的原则。
    * [forestdb](https://github.com/couchbase/goforestdb) - **Star : 30** 用于ForestDB的Go绑定。
    * [go-couchbase](https://github.com/couchbase/go-couchbase) - **Star : 290** Couchbase客户端在Go中。
    * [go-couchdb](https://github.com/fjl/go-couchdb) - **Star : 51** 还有另一个用于Go的CouchDB HTTP API包装器。
    * [go-pilosa](https://github.com/pilosa/go-pilosa) - **Star : 32**  Go Pilosa的客户端库。
    * [go-rejson](https://github.com/nitishm/go-rejson) - **Star : 81** 用于redislabs的ReJSON模块的Golang客户端使用Redigo Golang客户端。在redis中轻松地将结构体存储为JSON对象并对其进行操作。
    * [gocb](https://github.com/couchbase/gocb) - **Star : 288** 官方Couchbase Go SDK。
    * [gocql](http://gocql.github.io) - Apache Cassandra的Go语言驱动程序。
    * [godscache](https://github.com/defcronyke/godscache) - **Star : 6** 谷歌云平台Go Datastore包的包装器，它使用memcached添加缓存。
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - 用于Go编程语言的memcache客户端库。
    * [gorethink](https://github.com/dancannon/gorethink) - RethinkDB的Go语言驱动程序。
    * [goriak](https://github.com/zegl/goriak) - **Star : 24**  Go 语言驱动为Riak千伏。
    * [mgo](https://github.com/globalsign/mgo) - **Star : 1589** (未维护的)用于Go语言的MongoDB驱动程序，它在遵循标准Go习惯用法的非常简单的API下实现了丰富且经过良好测试的特性选择。
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - **Star : 2858** 用于Go语言的官方MongoDB驱动程序。
    * [neo4j](https://github.com/cihangir/neo4j) - **Star : 24** Golang的Neo4j Rest API绑定。
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - **Star : 72** golang中的Neo4j REST客户机。
    * [neoism](https://github.com/jmcvetta/neoism) - **Star : 355** Golang的Neo4j客户机。
    * [redigo](https://github.com/gomodule/redigo) - **Star : 6087** Redigo是Redis数据库的Go客户机。
    * [redis](https://github.com/go-redis/redis) - **Star : 6099** Redis是Golang的客户。
    * [redis](https://github.com/bsm/redeo) - **Star : 258** 与redis协议兼容的TCP服务器/服务。
    * [xredis](https://github.com/shomali11/xredis) - **Star : 9** 类型安全，可定制，干净和易于使用的Redis客户端。

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) - **Star : 5726** 现代文本索引库。
    * [elastic](https://github.com/olivere/elastic) - **Star : 3972** 弹力搜索客户端 Go 。
    * [elasticsql](https://github.com/cch123/elasticsql) - **Star : 366** 在Go中将sql转换为elasticsearch dsl。
    * [elastigo](https://github.com/mattbaird/elastigo) - **Star : 949** Elasticsearch客户机库。
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - **Star : 1444** 官方弹力搜索客户端为Go。
    * [goes](https://github.com/OwnLocal/goes) - **Star : 24** 库与Elasticsearch交互。
    * [riot](https://github.com/go-ego/riot) - **Star : 4613** 走开源、分布式、简单高效的搜索引擎之路。
    * [skizze](https://github.com/seiflotfy/skizze) - **Star : 68** 概率数据结构服务和存储。

* Multiple Backends.
    * [cachego](https://github.com/fabiorphp/cachego) - 用于多个驱动程序的Golang缓存组件。
    * [cayley](https://github.com/google/cayley) - 图形数据库，支持多个后端。
    * [dsc](https://github.com/viant/dsc) - **Star : 12** 用于SQL、NoSQL、结构化文件的数据存储连接。
    * [gokv](https://github.com/philippgille/gokv) - **Star : 74** Go的简单键值存储抽象和实现(Redis、、etcd、bbolt、BadgerDB、LevelDB、Memcached、DynamoDB、S3、PostgreSQL、MongoDB、CockroachDB等等)。

## 日期和时间

*用于处理日期和时间的库。*

* [carbon](https://github.com/uniplaces/carbon) - **Star : 330** 简单的时间扩展与许多util方法，从PHP碳库移植。
* [date](https://github.com/rickb777/date) - **Star : 27** 增加处理日期、日期范围、时间跨度、时间段和一天中的时间的时间。
* [dateparse](https://github.com/araddon/dateparse) - **Star : 875** 在不知道格式的情况下解析日期。
* [durafmt](https://github.com/hako/durafmt) - **Star : 234** 时间长度格式化库。
* [feiertage](https://github.com/wlbr/feiertage) - **Star : 21** 组函数计算公众假期在德国,德国各州及专业化(Bundeslander)。复活节,五旬节,感恩节…
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - **Star : 58** 波斯历法(太阳历)在戈琅的实施。
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) - **Star : 12** 计算给定位置的日出和日落时间。
* [goweek](https://github.com/grsmv/goweek) - **Star : 18** 用于与golang中的week实体一起工作的库。
* [iso8601](https://github.com/relvacode/iso8601) - **Star : 67** 有效地解析没有正则表达式的ISO8601日期时间。
* [kair](https://github.com/GuilhermeCaruso/kair) - **Star : 9** 日期和时间- Golang格式库。
* [now](https://github.com/jinzhu/now) - **Star : 2121** 现在是戈朗的时间工具箱。
* [NullTime](https://github.com/kirillDanshin/nulltime) - **Star : 9** Nullable“time.Time”。
* [strftime](https://github.com/awoodbeck/strftime) - **Star : 5** C99-compatible strftime格式化程序。
* [timespan](https://github.com/SaidinWoT/timespan) - **Star : 60** 用于与时间间隔交互，定义为开始时间和持续时间。
* [timeutil](https://github.com/leekchan/timeutil) - **Star : 168** 对golang时间包的有用扩展(时间增量、Strftime、…)。
* [tuesday](https://github.com/osteele/tuesday) - **Star : 7** Ruby-compatible Strftime函数。

## 分布式系统

*帮助构建分布式系统的包。*

* [celeriac](https://github.com/svcavallar/celeriac.v1) - **Star : 52** 库，用于添加对交互和监视芹菜工人、任务和事件的支持。
* [consistent](https://github.com/buraksezer/consistent) - **Star : 180** 与有界负载一致的散列。
* [dht](https://github.com/anacrolix/dht) - **Star : 121** BitTorrent Kademlia DHT实现。
* [digota](https://github.com/digota/digota) - **Star : 296** grpc电子商务microservice。
* [dot](https://github.com/dotchain/dot/) - 使用操作转换/OT的分布式同步。
* [doublejump](https://github.com/edwingeng/doublejump) - **Star : 37** 修改了谷歌的跳转一致散列。
* [dragonboat](https://github.com/lni/dragonboat) - **Star : 2423** 一个功能齐全，高性能的多组筏库在围棋。
* [drmaa](https://github.com/dgruber/drmaa) - **Star : 24** 基于DRMAA标准的集群调度程序作业提交库。
* [dynamolock](https://cirello.io/dynamolock) - 支持dynamodb的分布式锁定实现。
* [dynatomic](https://github.com/tylfin/dynatomic) - **Star : 8** 使用DynamoDB作为原子计数器的库。
* [emitter-io](https://github.com/emitter-io/emitter) - **Star : 1860** 高性能、分布式、安全和低延迟的发布-订阅平台，使用MQTT、Websockets和love构建。
* [flowgraph](https://github.com/vectaport/flowgraph) - **Star : 16** 基于流程的方法编程方案。
* [gleam](https://github.com/chrislusf/gleam) - **Star : 2026** 使用纯Go和Luajit编写的快速、可伸缩的分布式map/reduce系统，结合了Go的高并发性和Luajit的高性能，可以独立运行或分布式运行。
* [glow](https://github.com/chrislusf/glow) - **Star : 2483** 易于使用的可伸缩分布式大数据处理，Map-Reduce, DAG执行，全部在纯Go。
* [go-health](https://github.com/InVisionApp/go-health) - **Star : 470** 库，用于在服务中启用异步依赖项健康检查。
* [go-jump](https://github.com/dgryski/go-jump) - **Star : 249** 端口谷歌的“跳转”一致哈希函数。
* [go-kit](https://github.com/go-kit/kit) - **Star : 14030** 支持服务发现、负载平衡、可插入传输、请求跟踪等功能的Microservice toolkit。
* [gorpc](https://github.com/valyala/gorpc) - **Star : 546** 简单、快速和可伸缩的RPC库，用于高负载。
* [grpc-go](https://github.com/grpc/grpc-go) - **Star : 8605** gRPC的Go语言实现。HTTP RPC / 2。
* [hprose](https://github.com/hprose/hprose-golang) - **Star : 987** 非常新颖的RPC库，现在支持25+种语言。
* [jaeger](https://github.com/jaegertracing/jaeger) - **Star : 8324** 分布式跟踪系统。
* [jsonrpc](https://github.com/osamingo/jsonrpc) - **Star : 113** jsonrpc包帮助实现JSON-RPC 2.0。
* [jsonrpc](https://github.com/ybbus/jsonrpc) - **Star : 95** JSON-RPC 2.0 HTTP客户端实现。
* [KrakenD](https://github.com/devopsfaith/krakend) - **Star : 1619** 具有中间件的高性能API网关框架。
* [micro](https://github.com/micro/micro) - **Star : 6367** 可插入的微服务工具包和分布式系统平台。
* [NATS](https://github.com/nats-io/gnatsd) - 用于微服务、物联网和云本地系统的轻量级、高性能消息传递系统。
* [outboxer](https://github.com/italolelis/outboxer) - **Star : 1** Outboxer是一个实现发件箱模式的go库。
* [pglock](https://cirello.io/pglock) - postgresql支持的分布式锁定实现。
* [raft](https://github.com/hashicorp/raft) - **Star : 2762** Golang协议由HashiCorp实现。
* [raft](https://github.com/coreos/etcd/tree/master/raft) - Go实现的筏一致协议，由CoreOS。
* [redis-lock](https://github.com/bsm/redis-lock) - **Star : 145** 使用Redis简化分布式锁定实现。
* [resgate](https://resgate.io/) - 用于构建REST、实时和RPC API的实时API网关，其中所有客户端都是无缝同步的。
* [ringpop-go](https://github.com/uber/ringpop-go) - **Star : 562** 可伸缩的，容错的应用程序层分片的Go应用程序。
* [rpcx](https://github.com/smallnest/rpcx) - **Star : 3648** 分布式可插RPC服务框架，如阿里巴巴Dubbo。
* [sleuth](https://github.com/ursiform/sleuth) - **Star : 298** 用于HTTP服务之间的无主p2p自动发现和RPC的库(使用[ZeroMQ](https://github.com/zeromq/libzmq))。
* [tendermint](https://github.com/tendermint/tendermint) - **Star : 3054** 使用Tendermint consensus和区块链协议将任何编程语言编写的状态机转换成拜占庭式容错复制状态机的高性能中间件。
* [torrent](https://github.com/anacrolix/torrent) - **Star : 2752** bt客户端包。

## 电子邮件

*实现电子邮件创建和发送的库和工具。*

* [chasquid](https://blitiri.com.ar/p/chasquid) - 用Go编写的SMTP服务器。
* [douceur](https://github.com/aymerick/douceur) - **Star : 157** CSS内联为您的HTML电子邮件。
* [email](https://github.com/jordan-wright/email) - **Star : 1067** 一个强大和灵活的电子邮件库 Go 。
* [go-dkim](https://github.com/toorop/go-dkim) - **Star : 46** DKIM库，以签署和验证电子邮件。
* [go-imap](https://github.com/emersion/go-imap) - **Star : 699** 用于客户机和服务器的IMAP库。
* [go-message](https://github.com/emersion/go-message) - **Star : 101** 流媒体库，用于Internet消息格式和邮件消息。
* [go-premailer](https://github.com/vanng822/go-premailer) - **Star : 34** 内联样式为HTML邮件在Go。
* [Gomail](https://github.com/go-gomail/gomail/) - Gomail是一个非常简单和强大的邮件发送软件包。
* [Hectane](https://github.com/hectane/hectane) - **Star : 167** 轻量级SMTP客户机，提供HTTP API。
* [hermes](https://github.com/matcornic/hermes) - **Star : 1587** Golang包生成干净的、响应性强的HTML电子邮件。
* [MailHog](https://github.com/mailhog/MailHog) - **Star : 4997** 电子邮件和SMTP测试与web和API接口。
* [SendGrid](https://github.com/sendgrid/sendgrid-go) - **Star : 511** SendGrid的发送电子邮件的Go库。
* [smtp](https://github.com/mailhog/smtp) - **Star : 49** SMTP服务器协议状态机。

## 可嵌入的脚本语言

*在go代码中嵌入其他语言。*

* [agora](https://github.com/PuerkitoBio/agora) - 动态类型，可嵌入的编程语言在围棋。
* [anko](https://github.com/mattn/anko) - **Star : 896** 用Go编写的可编写脚本的解释器。
* [binder](https://github.com/alexeyco/binder) - **Star : 28** 转到基于[gopher-lua]的Lua绑定库(https://github.com/yuin/gopher-lua)。
* [expr](https://github.com/antonmedv/expr) - **Star : 552** 一个可以计算表达式的引擎。
* [gentee](https://github.com/gentee/gentee) - **Star : 23** 嵌入式脚本编程语言。
* [gisp](https://github.com/jcla1/gisp) - **Star : 426** 简单的LISP在围棋。
* [go-duktape](https://github.com/olebedev/go-duktape) - **Star : 648** Duktape JavaScript引擎绑定。
* [go-lua](https://github.com/Shopify/go-lua) - **Star : 1633** 端口的Lua 5.2 VM到纯Go。
* [go-php](https://github.com/deuill/go-php) - **Star : 666** 用于Go的PHP绑定。
* [go-python](https://github.com/sbinet/go-python) - **Star : 882** 简单的go绑定到CPython C-API。
* [golua](https://github.com/aarzilli/golua) - **Star : 434** Lua C API的Go绑定。
* [gopher-lua](https://github.com/yuin/gopher-lua) - **Star : 2897** 用Go编写的Lua 5.1虚拟机和编译器。
* [gval](https://github.com/PaesslerAG/gval) - **Star : 129** 一种用Go编写的高度可定制的表达式语言。
* [ngaro](https://github.com/db47h/ngaro) - **Star : 18** 嵌入式Ngaro VM实现，支持在Retro中编写脚本。
* [otto](https://github.com/robertkrimen/otto) - **Star : 4638** 用Go编写的JavaScript解释器。
* [purl](https://github.com/ian-kent/purl) - **Star : 26** Perl 5.18.2嵌入到Go中。
* [tengo](https://github.com/d5/tengo) - **Star : 1262** 字节码编译的脚本语言。

## 错误处理

*处理错误的库。*

* [errlog](https://github.com/snwfdhmp/errlog) - **Star : 139** 可编程包，用于确定错误的负责源代码(以及一些其他快速调试特性)。可插入到任何记录器的位置。
* [errors](https://github.com/pkg/errors) - **Star : 4667** 提供简单错误处理原语的包。
* [errorx](https://github.com/joomcode/errorx) - **Star : 545** 一个功能丰富的错误包与堆栈跟踪，组成的错误和更多。
* [go-multierror](https://github.com/hashicorp/go-multierror) - **Star : 703** Go (golang)包，用于将错误列表表示为单个错误。
* [tracerr](https://github.com/ztrue/tracerr) - **Star : 484** 堆栈跟踪和源碎片的Golang错误。
* [werr](https://github.com/txgruppi/werr) - **Star : 10** 错误包装器为Go中的错误类型创建一个包装器，该包装器捕获调用它的文件、行和堆栈。

## 文件

*处理文件和文件系统的库。*

* [afero](https://github.com/spf13/afero) - **Star : 2162** 文件系统抽象系统。
* [checksum](https://github.com/codingsince1985/checksum) - **Star : 6** 对于大型文件，计算消息摘要(如MD5和SHA256)。
* [flop](https://github.com/homedepot/flop) - **Star : 8** 文件操作库，它的目标是镜像特性与[GNU cp]的平价(https://www.gnu.org/software/coreutils/manual/html_node/cp- invoc.html)。
* [go-csv-tag](https://github.com/artonge/go-csv-tag) - **Star : 43** 使用标签加载csv文件。
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - **Star : 11** 为人类复制文件。
* [go-exiftool](https://github.com/barasher/go-exiftool) - ExifTool的Go绑定，这个著名的库用于从文件(图片、PDF、office，…)中提取尽可能多的元数据(EXIF、IPTC，…)。
* [go-gtfs](https://github.com/artonge/go-gtfs) - **Star : 15** 在go中加载gtfs文件。
* [notify](https://github.com/rjeczalik/notify) - **Star : 482** 文件系统事件通知库具有简单的API，类似于os/signal。
* [opc](https://github.com/qmuntal/opc) - **Star : 57** 为Go加载Open Packaging Conventions (OPC)文件。
* [pdfcpu](https://github.com/hhrutter/pdfcpu) - **Star : 916** PDF处理器。
* [skywalker](https://github.com/dixonwille/skywalker) - **Star : 46** 包，使您可以轻松地并发地遍历文件系统。
* [stl](https://gitlab.com/russoj88/stl) - 模块读取和写入STL(立体光刻)文件。并行读取算法。
* [tarfs](https://github.com/posener/tarfs) - **Star : 34** 为tar文件实现[' FileSystem '接口](https://godoc.org/github.com/kr/fs#FileSystem)。
* [vfs](https://github.com/C2FO/vfs) - **Star : 18** 一组可插拔的、可扩展的和自定义的文件系统功能，用于跨越许多文件系统类型，如os、S3和GCS。

## 金融

*会计和财务软件包。*

* [accounting](https://github.com/leekchan/accounting) - **Star : 476** golang的货币和货币格式。
* [currency](https://github.com/bnkamalesh/currency) - **Star : 8** 高性能、准确的货币计算软件包。
* [decimal](https://github.com/shopspring/decimal) - **Star : 1532** 任意精度定点十进制数。
* [go-finance](https://github.com/FlashBoys/go-finance) - **Star : 538** 综合金融市场数据在Go。
* [go-finance](https://github.com/alpeb/go-finance) - **Star : 39** 用于货币时间价值(年金)、现金流、利率转换、债券和折旧计算的金融函数库。
* [go-money](https://github.com/rhymond/go-money) - **Star : 604** 福勒货币模式的实现。
* [ofxgo](https://github.com/aclindsa/ofxgo) - **Star : 58** 查询OFX服务器和/或解析响应(使用示例命令行客户机)。
* [orderbook](https://github.com/i25959341/orderbook) - **Star : 60** Golang限购单匹配引擎。
* [techan](https://github.com/sdcoffey/techan) - **Star : 139** 拥有先进的市场分析和交易策略的技术分析库。
* [transaction](https://github.com/claygod/transaction) - **Star : 52** 嵌入式事务数据库的帐户，运行在多线程模式。
* [vat](https://github.com/dannyvankooten/vat) - **Star : 59** 增值税编号验证&欧盟增值税税率。

## 形式

*用于处理表单的库。*

* [bind](https://github.com/robfig/bind) - **Star : 23** 将表单数据绑定到任何Go值。
* [binding](https://github.com/mholt/binding) - **Star : 752** 将表单和JSON数据从net/http请求绑定到struct。
* [conform](https://github.com/leebenson/conform) - **Star : 171** 控制用户输入。基于struct标签对数据进行修剪、清理和擦除。
* [form](https://github.com/go-playground/form) - **Star : 346** 解码的url。将值转换为Go value(s)并将Go value(s)编码为url.Values。双数组和全地图支持。
* [formam](https://github.com/monoculum/formam) - **Star : 123** 将表单的值解码为结构。
* [forms](https://github.com/albrow/forms) - **Star : 103** 与框架无关的库，用于解析和验证支持多部分表单和文件的表单/JSON数据。
* [gorilla/csrf](https://github.com/gorilla/csrf) - **Star : 420** 用于Go web应用程序和服务的CSRF保护。
* [nosurf](https://github.com/justinas/nosurf) - **Star : 957** 用于Go的CSRF保护中间件。

## 功能

*在Go中支持函数式编程的包。*

* [fpGo](https://github.com/TeaEntityLab/fpGo) - **Star : 98** Monad，为Golang提供函数式编程功能。
* [fuego](https://github.com/seborama/fuego) - **Star : 33** 在围棋中进行功能实验。
* [go-underscore](https://github.com/tobyhede/go-underscore) - **Star : 1057** 有用的集合功能齐全的Go集合实用工具。

## 游戏开发

*很棒的游戏开发库。*

* [Azul3D](https://github.com/azul3d/engine) - **Star : 421** 用Go编写的3D游戏引擎。
* [Ebiten](https://github.com/hajimehoshi/ebiten) - **Star : 1742** 死简单的2D游戏库在围棋。
* [engo](https://github.com/EngoEngine/engo) - **Star : 1069** Engo是一个用Go编写的开源2D游戏引擎。它遵循实体-组件-系统范式。
* [g3n](https://github.com/g3n/engine) - **Star : 716**  Go 3D游戏引擎。
* [GarageEngine](https://github.com/vova616/GarageEngine) - **Star : 308** 用OpenGL编写的2d游戏引擎。
* [glop](https://github.com/runningwild/glop) - **Star : 77** Glop (Game Library Of Power)是一个相当简单的跨平台游戏库。
* [go-astar](https://github.com/beefsack/go-astar) - **Star : 320** Go实现了A\*路径查找算法。
* [go-collada](https://github.com/GlenKelley/go-collada) - **Star : 12**  Go 软件包处理Collada文件格式。
* [go-sdl2](https://github.com/veandco/go-sdl2) - **Star : 1129** 用于[Simple DirectMedia层]的Go绑定(https://www.libsdl.org/)。
* [go3d](https://github.com/ungerik/go3d) - **Star : 162** 面向性能的2D/3D数学软件包。
* [gonet](https://github.com/xtaci/gonet) - **Star : 1042** 游戏服务器骨架实现与golang。
* [goworld](https://github.com/xiaonanln/goworld) - **Star : 1124** 可伸缩的游戏服务器引擎，具有空间实体框架和热交换功能。
* [Leaf](https://github.com/name5566/leaf) - **Star : 2977** 轻量级游戏服务器框架。
* [nano](https://github.com/lonng/nano) - **Star : 954** 轻量级、方便、高性能的基于golang的游戏服务器框架。
* [Oak](https://github.com/oakmound/oak) - **Star : 620** 纯围棋游戏引擎。
* [Pitaya](https://github.com/topfreegames/pitaya) - **Star : 279** 可伸缩的游戏服务器框架，支持集群和客户端库的iOS, Android, Unity等通过C SDK。
* [Pixel](https://github.com/faiface/pixel) - **Star : 2362** 手工制作的2D游戏库在围棋。
* [raylib-go](https://github.com/gen2brain/raylib-go) - **Star : 377** Go bindings for [raylib](http://www.raylib.com/)，一个简单易用的库，用于学习视频游戏编程。
* [termloop](https://github.com/JoelOtter/termloop) - **Star : 1017** 基于终端的围棋游戏引擎，建立在Termbox之上。

## 一代和泛型

*通过代码生成用泛型之类的特性增强语言的工具。*

* [efaceconv](https://github.com/t0pep0/efaceconv) - **Star : 42** 代码生成工具，用于从接口{}到不需要分配的不可变类型的高性能转换。
* [gen](https://github.com/clipperhouse/gen) - **Star : 1013** “仿制药”——功能代码生成工具。
* [generis](https://github.com/senselogic/GENERIS) - **Star : 18** 提供泛型、自由形式宏、条件编译和HTML模板的代码生成工具。
* [go-enum](https://github.com/abice/go-enum) - **Star : 80** 从代码注释为枚举生成代码。
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - . net类linq的Go查询方法。
* [goderive](https://github.com/awalterschulze/goderive) - **Star : 716** 从输入类型派生函数。
* [gotype](https://github.com/wzshiming/gotype) - **Star : 19** 说明:Golang源码解析，用法类似反射包。-
* [GoWrap](https://github.com/hexdigest/gowrap) - **Star : 256** 使用简单模板为Go接口生成装饰器。
* [interfaces](https://github.com/rjeczalik/interfaces) - **Star : 183** 用于生成接口定义的命令行工具。
* [jennifer](https://github.com/dave/jennifer) - **Star : 1245** 生成没有模板的任意Go代码。
* [pkgreflect](https://github.com/ungerik/pkgreflect) - **Star : 82**  Go 预处理包范围内的反射。

## 地理

*地理工具和服务器*

* [geocache](https://github.com/melihmucuk/geocache) - **Star : 105** 适用于基于地理位置的应用程序的内存缓存。
* [geoserver](https://github.com/hishamkaram/geoserver) - **Star : 24** geoserver是一个Go包，用于通过geoserver REST API操作geoserver实例。
* [gismanager](https://github.com/hishamkaram/gismanager) - **Star : 17** 将GIS数据(矢量数据)发布到PostGIS和Geoserver。
* [osm](https://github.com/paulmach/osm) - **Star : 58** 用于读取、编写和处理OpenStreetMap数据和api的库。
* [pbf](https://github.com/maguro/pbf) - **Star : 15** OpenStreetMap PBF golang编码器/解码器。
* [S2 geometry](https://github.com/golang/geo) - **Star : 868** S2几何库在围棋。
* [Tile38](https://github.com/tidwall/tile38) - **Star : 6255** 具有空间索引和实时地理定位功能的地理定位数据库。

## Go 编译器

*编译工具转向其他语言。*

* [c4go](https://github.com/Konstantin8105/c4go) - **Star : 142** 转置C代码到Go代码。
* [f4go](https://github.com/Konstantin8105/f4go) - **Star : 11** 转置FORTRAN 77代码到Go代码。
* [gopherjs](https://github.com/gopherjs/gopherjs) - **Star : 8412** 编译器从到JavaScript。
* [llgo](https://github.com/go-llvm/llgo) - **Star : 986** 基于llvm的编译器。
* [tardisgo](https://github.com/tardisgo/tardisgo) - **Star : 391** goang to Haxe to CPP/CSharp/Java/JavaScript转置器。

## Goroutines

*管理和使用Goroutines的工具。*

* [ants](https://github.com/panjf2000/ants) - **Star : 1722** 一个高性能的戈朗goroutine池。
* [artifex](https://github.com/borderstech/artifex) - **Star : 12** 简单的内存作业队列为Golang使用基于工人的调度。
* [async](https://github.com/studiosol/async) - **Star : 18** 一种异步执行函数的安全方法，在出现恐慌时恢复它们。
* [breaker](https://github.com/kamilsk/breaker) - **Star : 23** 🚧灵活的机制,以使代码易碎物品。
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - **Star : 27** CyclicBarrier golang。
* [go-floc](https://github.com/workanator/go-floc) - **Star : 167** 轻松地编排goroutines。
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) - **Star : 103** 控制goroutines的执行顺序。
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) - **Star : 5** 使用这个轻量级库和一个简单的API管理一个goroutine池。
* [go-trylock](https://github.com/subchen/go-trylock) - **Star : 4** 尝试对Golang的读写锁提供支持。
* [gollback](https://github.com/vardius/gollback) - **Star : 24** 异步简单函数实用程序，用于管理闭包和回调的执行。
* [GoSlaves](https://github.com/themester/GoSlaves) - 简单异步Goroutine池库。
* [goworker](https://github.com/benmanns/goworker) - **Star : 2214** goworker是一个基于go的后台工作者。
* [gpool](https://github.com/Sherifabdlnaby/gpool) - **Star : 58** 管理可调整大小的上下文感知goroutine池，以绑定并发性。
* [grpool](https://github.com/ivpusic/grpool) - **Star : 486** 轻量级Goroutine池。
* [Hunch](https://github.com/AaronJan/Hunch) - **Star : 8** Hunch提供了诸如“All”、“First”、“Retry”、“Waterfall”等功能，这使得异步流控制更加直观。
* [oversight](https://cirello.io/oversight) - 监督是Erlang监督树的完整实现。
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - **Star : 24** 并行运行函数。
* [pool](https://github.com/go-playground/pool) - **Star : 471** 有限的消费者goroutine或无限的goroutine池更容易的goroutine处理和取消。
* [queue](https://github.com/AnikHasibul/queue) - **Star : 1** 给你一个“同步”。比如队列组可访问性。帮助您节流和限制goroutine，等待所有goroutine和更多的结束。
* [routine](https://github.com/x-mod/routine) - **Star : 1** 带有上下文、支持:Main、go、Pool和一些有用的执行器的go例程控件。
* [semaphore](https://github.com/kamilsk/semaphore) - **Star : 74** 信号量模式实现，根据通道和上下文超时锁定/解锁操作。
* [semaphore](https://github.com/marusama/semaphore) - **Star : 68** 基于CAS的快速可调整大小的信号量实现(比基于通道的信号量实现更快)。
* [stl](https://github.com/ssgreg/stl) - **Star : 7** 基于软件事务内存(STM)并发控制机制的软件事务锁。
* [threadpool](https://github.com/shettyh/threadpool) - **Star : 17** Golang threadpool实现。
* [tunny](https://github.com/Jeffail/tunny) - **Star : 1268** 戈朗的戈朗游泳池。
* [worker-pool](https://github.com/vardius/worker-pool) - **Star : 43** goworker是一个Go simple异步工作池。
* [workerpool](https://github.com/gammazero/workerpool) - **Star : 116** 限制任务执行并发性的Goroutine池，而不是队列中的任务数量。

## GUI

*用于构建GUI应用程序的库。*

*工具包*

* [app](https://github.com/murlokswarm/app) - 包创建应用程序与GO, HTML和CSS。支持:MacOS, Windows正在开发中。
* [fyne](https://github.com/fyne-io/fyne) - **Star : 5964** 跨平台的本地gui是为围棋而设计的，使用EFL呈现。支持:Linux, macOS, Windows。
* [go-astilectron](https://github.com/asticode/go-astilectron) - **Star : 2582** 使用GO和HTML/JS/CSS(电子驱动)构建跨平台GUI应用程序。
* [go-gtk](http://mattn.github.io/go-gtk/) - GTK的Go绑定。
* [go-sciter](https://github.com/sciter-sdk/go-sciter) - **Star : 1411** Sciter的Go绑定:用于现代桌面UI开发的可嵌入HTML/CSS/脚本引擎。交叉平台。
* [gotk3](https://github.com/gotk3/gotk3) - **Star : 722** 用于GTK3的Go绑定。
* [gowd](https://github.com/dtylman/gowd) - **Star : 200** 快速和简单的桌面UI开发与GO, HTML, CSS和NW.js。交叉平台。
* [qt](https://github.com/therecipe/qt) - **Star : 5775** Qt绑定for Go(支持Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi)。
* [ui](https://github.com/andlabs/ui) - **Star : 6810** 用于Go的平台原生GUI库。交叉平台。
* [Wails](https://wails.app) - Mac, Windows, Linux桌面应用程序的HTML用户界面使用内置的OS HTML渲染器。
* [walk](https://github.com/lxn/walk) - **Star : 3601** Windows应用程序库工具包。
* [webview](https://github.com/zserge/webview) - **Star : 4497** 跨平台webview窗口，具有简单的双向JavaScript绑定(Windows / macOS / Linux)。

*交互*

* [go-appindicator](https://github.com/dawidd6/go-appindicator) - **Star : 1**  Go 绑定libappindicator3 C库。
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) - **Star : 492** 用于Go的OSX桌面通知库。
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) - **Star : 1** 用于通知计算机上的任何(可插入的)活动。
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) - golang中的OSX睡眠/唤醒通知。
* [robotgo](https://github.com/go-vgo/robotgo) - **Star : 4335** 实现跨平台的GUI系统自动化。控制鼠标、键盘等。
* [systray](https://github.com/getlantern/systray) - **Star : 749** 跨平台Go库在通知区放置图标和菜单。
* [trayhost](https://github.com/shurcooL/trayhost) - **Star : 158** 跨平台Go库，用于在主机操作系统的任务栏中放置图标。


## 硬件

*用于与硬件交互的库、工具和教程。*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## 图片

*用于操作图像的库。*

* [bild](https://github.com/anthonynsimon/bild) - **Star : 2016** 集合了纯Go中的图像处理算法。
* [bimg](https://github.com/h2non/bimg) - **Star : 765** 使用libvips快速有效地处理图像的小程序包。
* [cameron](https://github.com/aofei/cameron) - **Star : 30** 一个围棋的化身生成器。
* [geopattern](https://github.com/pravj/geopattern) - **Star : 1008** 从一个字符串创建美丽的生成图像模式。
* [gg](https://github.com/fogleman/gg) - **Star : 1858** 纯粹的围棋2D渲染。
* [gift](https://github.com/disintegration/gift) - **Star : 1201** 图像处理滤波器包。
* [gltf](https://github.com/qmuntal/gltf) - **Star : 34** 高效、健壮的glTF 2.0阅读器、编写器和验证器。
* [go-cairo](https://github.com/ungerik/go-cairo) - **Star : 85**  Go 绑定cairo图形库。
* [go-gd](https://github.com/bolknote/go-gd) - **Star : 48**  Go 绑定GD库。
* [go-nude](https://github.com/koyachi/go-nude) - **Star : 278** 裸照检测与Go。
* [go-opencv](https://github.com/lazywei/go-opencv) - OpenCV的Go绑定。
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - **Star : 24** 端口的webcolors库从Python Go 。
* [gocv](https://github.com/hybridgroup/gocv) - **Star : 2344** 使用OpenCV 3.3+进行计算机视觉打包。
* [goimagehash](https://github.com/corona10/goimagehash) - **Star : 206**  Go 感知图像哈希包。
* [goimghdr](https://github.com/corona10/goimghdr) - **Star : 26** imghdr模块确定Go文件中包含的图像类型。
* [govatar](https://github.com/o1egl/govatar) - **Star : 306** 库和CMD工具生成有趣的化身。
* [image2ascii](https://github.com/qeesung/image2ascii) - **Star : 281** 将图像转换为ASCII码。
* [imagick](https://github.com/gographics/imagick) - **Star : 951**  Go 绑定ImageMagick的MagickWand C API。
* [imaginary](https://github.com/h2non/imaginary) - **Star : 2534** 快速和简单的HTTP微服务，用于图像大小调整。
* [imaging](https://github.com/disintegration/imaging) - **Star : 2453** 简单的 Go 图像处理软件包。
* [img](https://github.com/hawx/img) - **Star : 129** 图像处理工具的选择。
* [ln](https://github.com/fogleman/ln) - **Star : 2435** 3D线艺术渲染在围棋。
* [mergi](https://github.com/noelyahan/mergi) - **Star : 68** 用于图像处理的工具和Go库(合并、裁剪、调整大小、水印、动画)。
* [mort](https://github.com/aldor007/mort) - **Star : 362** 存储和图像处理服务器写在Go。
* [mpo](https://github.com/donatj/mpo) - **Star : 6** 为MPO三维照片解码器和转换工具。
* [picfit](https://github.com/thoas/picfit) - **Star : 1046** 一个用Go编写的图像调整服务器。
* [pt](https://github.com/fogleman/pt) - **Star : 1750** 用Go编写的路径跟踪引擎。
* [resize](https://github.com/nfnt/resize) - **Star : 2104** 使用常用的插值方法调整Go的图像大小。
* [rez](https://github.com/bamiaux/rez) - **Star : 161** 图像大小调整在纯 Go 和SIMD。
* [smartcrop](https://github.com/muesli/smartcrop) - **Star : 1238** 为任意图像和作物大小找到好的作物。
* [steganography](https://github.com/auyer/steganography) - **Star : 25** 用于LSB隐写的纯Go库。
* [stegify](https://github.com/DimitarPetrov/stegify) - **Star : 478**  Go 工具为LSB隐写术，能够隐藏任何文件在一个图像。
* [svgo](https://github.com/ajstarks/svgo) - **Star : 1306**  Go 语言库为SVG生成。
* [tga](https://github.com/ftrvxmtrx/tga) - **Star : 23** 包tga是一个TARGA图像格式解码器/编码器。

## 物联网

*物联网编程设备库。*

* [connectordb](https://github.com/connectordb/connectordb) - **Star : 165** 量化自我和物联网的开源平台。
* [devices](https://github.com/goiot/devices) - **Star : 224** 用于物联网设备的一套库，用于x/exp/io的实验性库。
* [eywa](https://github.com/xcodersun/eywa) - **Star : 35** Eywa项目本质上是一个连接管理器，用于跟踪连接的设备。
* [flogo](https://github.com/tibcosoftware/flogo) - **Star : 1086** Flogo项目是一个面向物联网边缘应用和集成的开源框架。
* [gatt](https://github.com/paypal/gatt) - **Star : 805** Gatt是一个用于构建蓝牙低能耗外围设备的Go软件包。
* [gobot](https://github.com/hybridgroup/gobot/) - Gobot是一个用于机器人、物理计算和物联网的框架。
* [huego](https://github.com/amimof/huego) - **Star : 106** 一个广泛的飞利浦顺化客户端库 Go 。
* [iot](https://github.com/vaelen/iot/) - 物联网是实现谷歌物联网核心设备的简单框架。
* [mainflux](https://github.com/Mainflux/mainflux) - **Star : 550** 工业物联网消息和设备管理服务器。
* [periph](https://periph.io/) - 外围设备I/O与低层电路板设备接口。
* [sensorbee](https://github.com/sensorbee/sensorbee) - **Star : 175** 物联网轻量级流处理引擎。

## 作业调度器

*用于调度作业的库。*

* [clockwerk](http://github.com/onatm/clockwerk) - Go package使用简单、流畅的语法安排周期性的作业。
* [clockwork](https://github.com/whiteShtef/clockwork) - **Star : 74** 简单直观的作业调度库在围棋。
* [go-cron](https://github.com/rk/go-cron) - **Star : 177** 用于go的简单Cron库，可以在不同的时间间隔执行闭包或函数，从每秒一次到每年在特定的日期和时间执行一次。主要用于web应用程序和长时间运行的守护进程。
* [gron](https://github.com/roylee0704/gron) - **Star : 620** 定义基于时间的任务使用一个简单的API, Gron相应的调度程序将运行它们。
* [JobRunner](https://github.com/bamzi/jobrunner) - **Star : 559** 智能和功能丰富的cron作业调度程序与作业排队和实时监控内置。
* [jobs](https://github.com/albrow/jobs) - **Star : 449** 持久和灵活的后台作业库。
* [leprechaun](https://github.com/kilgaloon/leprechaun) - **Star : 36** 支持webhook、crons和经典调度的作业调度程序。
* [scheduler](https://github.com/carlescere/scheduler) - **Star : 290** 克朗乔布斯的日程安排变得很简单。

## JSON

*用于处理JSON的库。*

* [ajson](https://github.com/spyzhov/ajson) - **Star : 10** 抽象JSON为golang与JSONPath支持。
* [gjo](https://github.com/skanehira/gjo) - **Star : 57** 创建JSON对象的小实用程序。
* [GJSON](https://github.com/tidwall/gjson) - **Star : 4673** 用一行代码获取JSON值。
* [go-respond](https://github.com/nicklaw5/go-respond) - **Star : 20** Go package用于处理常见的HTTP JSON响应。
* [gojq](https://github.com/elgs/gojq) - **Star : 138** Golang中的JSON查询。
* [gojson](https://github.com/ChimeraCoder/gojson) - **Star : 2014** 从示例JSON自动生成Go (golang)结构定义。
* [JayDiff](https://github.com/yazgazan/jaydiff) - **Star : 37** 用Go编写的JSON diff实用程序。
* [JSON-to-Go](https://mholt.github.io/json-to-go/) - 将JSON转换为结构体。
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - **Star : 5** 基于JSON API错误引用的Go绑定。
* [jsonf](https://github.com/miolini/jsonf) - **Star : 54** 控制台工具，用于突出显示格式和结构查询获取JSON。
* [jsongo](https://github.com/ricardolonga/jsongo) - **Star : 91** 使用Fluent API更容易地创建Json对象。
* [jsonhal](https://github.com/RichardKnop/jsonhal) - **Star : 9** 简单的Go包，使定制结构元帅成哈尔兼容JSON响应。
* [kazaam](https://github.com/Qntfy/kazaam) - **Star : 123** 用于任意JSON文档转换的API。
* [mp](https://github.com/sanbornm/mp) - **Star : 32** 简单的cli电子邮件解析器。它当前接受stdin并输出JSON。

## 日志记录

*用于生成和处理日志文件的库。*

* [distillog](https://github.com/amoghe/distillog) - **Star : 18** 经过蒸馏的水平日志记录(可以将其视为stdlib +日志级别)。
* [glg](https://github.com/kpango/glg) - **Star : 51** glg是一个简单而快速的Go日志库。
* [glo](https://github.com/lajosbencz/glo) - **Star : 7** PHP独白激发了具有相同严重性级别的日志记录功能。
* [glog](https://github.com/golang/glog) - **Star : 2253** 为Go提供了水平的执行日志。
* [go-cronowriter](https://github.com/utahta/go-cronowriter) - **Star : 19** 基于当前日期和时间自动旋转日志文件的简单编写器，如cronolog。
* [go-log](https://github.com/subchen/go-log) - **Star : 10** 简单且可配置的登录Go，带有level、格式化程序和编写器。
* [go-log](https://github.com/siddontang/go-log) - **Star : 23** 日志库支持级别和多处理程序。
* [go-log](https://github.com/ian-kent/go-log) - **Star : 34** 在Go中实现Log4j。
* [go-logger](https://github.com/apsdehal/go-logger) - **Star : 229** 简单的日志程序的围棋程序，与级别处理程序。
* [gologger](https://github.com/sadlil/gologger) - **Star : 40** 简单易用的日志库，日志在彩色控制台，简单控制台，文件或弹性搜索。
* [gomol](https://github.com/aphistic/gomol) - **Star : 16** 具有可扩展日志输出的多输出、结构化日志记录。
* [gone/log](https://github.com/One-com/gone/tree/master/log) - 快速、可扩展、功能齐全、std-lib源代码兼容的日志库。
* [journald](https://github.com/ssgreg/journald) - **Star : 17**  Go 实现systemd Journal的本地日志API。
* [log](https://github.com/aerogo/log) - **Star : 4** 一个O(1)日志系统，允许您将一个日志连接到多个写入器(例如stdout、文件和TCP连接)。
* [log](https://github.com/apex/log) - **Star : 718** Go的结构化日志包。
* [log](https://github.com/go-playground/log) - **Star : 265** 简单、可配置和可伸缩的Go结构化日志。
* [log](https://github.com/teris-io/log) - **Star : 22** Go的结构化日志接口清晰地将日志外观与其实现分离开来。
* [log-voyage](https://github.com/firstrow/logvoyage) - **Star : 82** 用golang编写的功能齐全的日志saas。
* [log15](https://github.com/inconshreveable/log15) - **Star : 876** 简单、强大的Go日志。
* [logdump](https://github.com/ewwwwwqm/logdump) - **Star : 8** 用于多级日志记录的包。
* [logex](https://github.com/chzyer/logex) - **Star : 32** Golang日志库，支持跟踪和水平，包装由标准日志库。
* [logger](https://github.com/azer/logger) - **Star : 133** Go的最小化日志库。
* [logmatic](https://github.com/borderstech/logmatic) - **Star : 7** 带有动态日志级别配置的Golang彩色日志记录器。
* [logo](https://github.com/mbndr/logo) - **Star : 4** Golang日志程序到不同的可配置作家。
* [logrus](https://github.com/Sirupsen/logrus) - **Star : 11374** 结构化的日志 Go 。
* [logrusly](https://github.com/sebest/logrusly) - **Star : 26** [logrus](https://github.com/sirupsen/logrus)插件将错误发送到[Loggly](https://www.loggly.com/)。
* [logutils](https://github.com/hashicorp/logutils) - **Star : 246** 用于稍微更好地登录的实用程序Go (Golang)扩展了标准日志记录器。
* [logxi](https://github.com/mgutz/logxi) - **Star : 332** 12因素的应用程序日志程序是快速的，让你快乐。
* [lumberjack](https://github.com/natefinch/lumberjack) - **Star : 1368** 简单的滚动日志程序，实现io.WriteCloser。
* [mlog](https://github.com/jbrodriguez/mlog) - **Star : 17** 简单的go日志模块，有5个级别，可选的旋转日志文件功能和stdout/stderr输出。
* [onelog](https://github.com/francoispqt/onelog) - **Star : 324** Onelog是一个非常简单但非常高效的JSON日志程序。它是所有场景中速度最快的JSON日志程序。而且，它是配置最低的日志记录器之一。
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) - **Star : 108** 支持日志严重性、分类和过滤的高性能日志记录。可以发送过滤日志消息到各种目标(如控制台，网络，邮件)。
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) - **Star : 93** RollingWriter是一个自动旋转的io。作者的实现与多个策略，以提供日志文件旋转。
* [seelog](https://github.com/cihub/seelog) - **Star : 1330** 具有灵活调度、过滤和格式化的日志功能。
* [spew](https://github.com/davecgh/go-spew) - **Star : 3218** 为Go数据结构实现一个漂亮的深层打印机，以帮助调试。
* [stdlog](https://github.com/alexcesaro/log) - **Star : 43** Stdlog是一个面向对象的库，提供水平日志记录。它对cron作业非常有用。
* [tail](https://github.com/hpcloud/tail) - **Star : 1487** Go软件包努力模拟BSD tail程序的特性。
* [xlog](https://github.com/xfxdev/xlog) - **Star : 7** 插件架构和灵活的日志系统的Go，与级别ctrl，多日志目标和自定义日志格式。
* [xlog](https://github.com/rs/xlog) - **Star : 129** 结构化记录器'net/context`意识到HTTP处理程序的灵活调度。
* [zap](https://github.com/uber-go/zap) - **Star : 7131** 快速、结构化、水平登录Go。
* [zerolog](https://github.com/rs/zerolog) - **Star : 2072** 零JSON记录器。

## 机器学习

*机器学习库。*

* [bayesian](https://github.com/jbrukh/bayesian) - **Star : 623** Golang的朴素贝叶斯分类。
* [CloudForest](https://github.com/ryanbressler/CloudForest) - **Star : 639** 快速、灵活、多线程的决策树集成，用于纯Go中的机器学习。
* [eaopt](https://github.com/MaxHalford/eaopt) - **Star : 609** 一个进化优化库。
* [evoli](https://github.com/khezen/evoli) - **Star : 6** 遗传算法和粒子群优化库。
* [fonet](https://github.com/Fontinalis/fonet) - **Star : 30** 一个用Go编写的深度神经网络库。
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - **Star : 20** Go实现了k模式和k原型聚类算法。
* [go-deep](https://github.com/patrikeh/go-deep) - **Star : 212** 一个功能丰富的神经网络库在围棋。
* [go-fann](https://github.com/white-pony/go-fann) - **Star : 98** 快速人工神经网络(FANN)库的Go绑定。
* [go-galib](https://github.com/thoj/go-galib) - **Star : 168** 用Go / golang编写的遗传算法库。
* [go-pr](https://github.com/daviddengcn/go-pr) - **Star : 56** 模式识别包在Go lang。
* [gobrain](https://github.com/goml/gobrain) - **Star : 363** 用围棋编写的神经网络。
* [godist](https://github.com/e-dard/godist) - **Star : 23** 各种概率分布，以及相关的方法。
* [goga](https://github.com/tomcraven/goga) - **Star : 77** Go的遗传算法库。
* [GoLearn](https://github.com/sjwhitworth/golearn) - **Star : 6580** 通用机器学习库。
* [golinear](https://github.com/danieldk/golinear) - **Star : 38** 围棋的线性绑定。
* [GoMind](https://github.com/surenderthakran/gomind) - **Star : 5** 一个简单的神经网络库在围棋。
* [goml](https://github.com/cdipaolo/goml) - **Star : 998** 在线机器学习在围棋。
* [goRecommend](https://github.com/timkaye11/goRecommend) - **Star : 140** 推荐算法库用Go编写。
* [gorgonia](https://github.com/chewxy/gorgonia) - 基于图形的计算库，如Theano for Go，它为构建各种机器学习和神经网络算法提供了基本框架。
* [gorse](https://github.com/zhenghaoz/gorse) - **Star : 506** 基于协同过滤的Go高性能推荐系统包。
* [goscore](https://github.com/asafschers/goscore) - **Star : 33**  Go 为PMML评分API。
* [gosseract](https://github.com/otiai10/gosseract) - **Star : 839** 使用Tesseract c++库为OCR(光学字符识别)打包。
* [libsvm](https://github.com/datastream/libsvm) - **Star : 62** 基于libsvm的golang版本派生工作。
* [mlgo](https://github.com/NullHypothesis/mlgo) - **Star : 4** 这个项目的目的是在围棋中提供最小化的机器学习算法。
* [neat](https://github.com/jinyeom/neat) - **Star : 55** 即插即用的并行Go框架，用于增强拓扑的神经进化(整洁)。
* [neural-go](https://github.com/schuyler/neural-go) - **Star : 60** 多层感知器网络在Go中实现，通过反向传播进行训练。
* [ocrserver](https://github.com/otiai10/ocrserver) - **Star : 215** 一个简单的OCR API服务器，非常容易被Docker和Heroku部署。
* [onnx-go](https://github.com/owulveryck/onnx-go) - **Star : 133** Go接口打开神经网络交换(ONNX)。
* [probab](https://github.com/ThePaw/probab) - **Star : 9** 概率分布函数。贝叶斯推理。用纯围棋写的。
* [regommend](https://github.com/muesli/regommend) - **Star : 242** 推荐&协同过滤引擎。
* [shield](https://github.com/eaigner/shield) - **Star : 122** 贝叶斯文本分类器，具有灵活的令牌器和存储后端。
* [tfgo](https://github.com/galeone/tfgo) - **Star : 1154** 易于使用的Tensorflow绑定:简化了官方Tensorflow Go绑定的使用。在Go中定义计算图形，加载和执行Python中训练的模型。
* [Varis](https://github.com/Xamber/Varis) - **Star : 23** Golang神经网络。

## 消息传递

*实现消息传递系统的库。*

* [APNs2](https://github.com/sideshow/apns2) - **Star : 2021** HTTP / 2苹果推送通知供应商——发送推送通知到iOS, tvo, Safari和OSX的应用。
* [Beaver](https://github.com/Clivern/Beaver) - **Star : 711** 一个实时消息服务器，可构建一个可伸缩的应用程序内通知，多人游戏，聊天应用程序在web和移动应用程序。
* [Benthos](https://github.com/Jeffail/benthos) - **Star : 1870** 一系列协议之间的消息流桥。
* [Bus](https://github.com/mustafaturan/bus) - **Star : 109** 内部通信的最小消息总线实现。
* [Centrifugo](https://github.com/centrifugal/centrifugo) - **Star : 3604** 实时消息(Websockets或SockJS)服务器。
* [Commander](https://github.com/jeroenrinzema/commander) - **Star : 19** 高级事件驱动的消费者/生产者，支持各种“方言”，如Apache Kafka。
* [dbus](https://github.com/godbus/dbus) - **Star : 349** D-Bus的本地Go绑定。
* [drone-line](https://github.com/appleboy/drone-line) - **Star : 58** 使用二进制、docker或从属CI发送[Line](https://at.line.me/en)通知。
* [emitter](https://github.com/olebedev/emitter) - **Star : 306** 使用通配符、谓词、取消可能性和许多其他优点，使用Go way发出事件。
* [event](https://github.com/agoalofalife/event) - **Star : 24** 模式观察者的实现。
* [EventBus](https://github.com/asaskevich/EventBus) - **Star : 537** 具有异步兼容性的轻量级事件总线。
* [gaurun-client](https://github.com/osamingo/gaurun-client) - **Star : 8** 用Go编写的Gaurun客户端。
* [Glue](https://github.com/desertbit/glue) - **Star : 312** 健壮的Go和Javascript套接字库(替代Socket.io)。
* [go-notify](https://github.com/TheCreeper/go-notify) - **Star : 47** 本地实现的freedesktop通知规范。
* [go-nsq](https://github.com/nsqio/go-nsq) - **Star : 1431** NSQ的官方Go包。
* [go-socket.io](https://github.com/googollee/go-socket.io) - **Star : 2824** 套接字。面向golang的io库，一个实时应用程序框架。
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) - **Star : 11** 客户端库到Viessmann Vitotrol web服务。
* [Gollum](https://github.com/trivago/gollum) - **Star : 758** n:m多路复用器，它收集来自不同来源的消息并将其广播到一组目的地。
* [golongpoll](https://github.com/jcuga/golongpoll) - **Star : 418** HTTP longpoll服务器库，使web发布-订阅变得简单。
* [goose](https://github.com/ian-kent/goose) - **Star : 37** 服务器在Go中发送事件。
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - **Star : 1818** gopush-cluster是一个gopush服务器集群。
* [gorush](https://github.com/appleboy/gorush) - **Star : 3614** 使用[APNs2](https://github.com/sideshow/apns2)和谷歌[GCM](https://github.com/google/go-gcm)推送通知服务器。
* [guble](https://github.com/smancke/guble) - **Star : 138** 消息服务器使用推送通知(谷歌Firebase云消息、苹果推送通知服务、SMS)以及websockets，一个REST API，具有分布式操作和消息持久性。
* [hub](https://github.com/leandro-lugaresi/hub) - **Star : 24** 用于Go应用程序的消息/事件中心，使用发布/订阅模式，并支持别名(如rabbitMQ交换)。
* [jazz](https://github.com/socifi/jazz) - **Star : 6** 一个简单的RabbitMQ抽象层，用于队列管理和消息的发布和消费。
* [machinery](https://github.com/RichardKnop/machinery) - **Star : 3273** 基于分布式消息传递的异步任务队列/作业队列。
* [mangos](https://github.com/go-mangos/mangos) - 具有传输互操作性的Nanomsg(“可伸缩协议”)的纯go实现。
* [melody](https://github.com/olahol/melody) - **Star : 1499** 处理websocket会话的极简框架，包括广播和自动乒乓球处理。
* [Mercure](https://github.com/dunglas/mercure) - **Star : 1427** 使用Mercure协议(构建在服务器发送事件之上)分派服务器发送的更新的服务器和库。
* [messagebus](https://github.com/vardius/message-bus) - **Star : 62** messagebus是一种Go简单异步消息总线，非常适合在执行事件源、CQRS和DDD时用作事件总线。
* [NATS Go Client](https://github.com/nats-io/nats) - 轻量级和高性能的发布-订阅和分布式队列消息传递系统——这是Go库。
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - **Star : 49** 一个围绕NSQ主题和通道的小包装。
* [oplog](https://github.com/dailymotion/oplog) - **Star : 94** 用于REST api的通用oplog/复制系统。
* [pubsub](https://github.com/tuxychandru/pubsub) - 简单的pubsubpackage for go。
* [rabbus](https://github.com/rafaeljesus/rabbus) - **Star : 60** amqp交换器和队列上的一个小包装。
* [rabtap](https://github.com/jandelgado/rabtap) - **Star : 65** RabbitMQ瑞士军刀cli应用。
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) - **Star : 55** RapidMQ是用于管理本地消息队列的轻量级可靠库。
* [rmqconn](https://github.com/sbabiv/rmqconn) - RabbitMQ重新连接。amqp包装器。连接和amqp.Dial。在强制关闭对Close()方法的调用之前，允许在连接断开时重新连接。
* [sarama](https://github.com/Shopify/sarama) - **Star : 4416**  Go Apache Kafka的库。
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - **Star : 1088** Redis支持面向移动设备的服务器端通知的统一推送服务。
* [zmq4](https://github.com/pebbe/zmq4) - **Star : 767** 转接口到ZeroMQ版本4。也可用于[版本3](https://github.com/pebbe/zmq3)和[版本2](https://github.com/pebbe/zmq2)。

## 微软办公软件

### Microsoft Excel
*用于使用Microsoft Excel的库。*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - **Star : 4184** Golang库读写Microsoft Excel™(XLSX)文件。
* [go-excel](https://github.com/szyhf/go-excel) - **Star : 45** 一个简单而轻便的阅读器，可以将类似于关系数据库的excel作为表来读取。
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - **Star : 12** 用于编写XLSX (Microsoft Excel)文件的libxlsxwriter的Golang绑定。
* [xlsx](https://github.com/tealeg/xlsx) - **Star : 3265** 库，以简化在Go程序中读取Microsoft Excel最新版本使用的XML格式。
* [xlsx](https://github.com/plandem/xlsx) - **Star : 58** 快速和安全的方式读取/更新您现有的Microsoft Excel文件在围棋程序。

## 杂项

### 依赖注入
*用于处理依赖项注入的库。*

* [alice](https://github.com/magic003/alice) - **Star : 34** Golang的添加依赖注入容器。
* [dig](https://github.com/uber-go/dig) - **Star : 852** 一个基于反射的Go依赖注入工具包。
* [fx](https://github.com/uber-go/fx) - **Star : 631** 基于依赖注入的Go应用程序框架(构建在dig之上)。
* [gocontainer](https://github.com/vardius/gocontainer) - **Star : 3** 简单的依赖注入容器。
* [inject](https://github.com/defval/inject) - **Star : 16** 一个基于反射的依赖注入容器，具有简单的接口。
* [wire](https://github.com/Fs02/wire) - **Star : 20** Golang严格的运行时依赖注入。

### 项目布局

*用于组织项目的非正式模式集。*

* [go-sample](https://github.com/zitryss/go-sample) - **Star : 15** 使用实际代码的Go应用程序项目的示例布局。
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - **Star : 8305** Go生态系统中常见的历史和新兴的项目布局模式。
* [scaffold](https://github.com/catchplay/scaffold) - **Star : 19** 脚手架生成starter Go项目布局。让您专注于已实现的业务逻辑。

### 字符串
*处理字符串的库。*
* [strutil](https://github.com/ozgio/strutil) - **Star : 61** 字符串工具。
* [xstrings](https://github.com/huandu/xstrings) - **Star : 604** 从其他语言移植的有用字符串函数的集合。

*这些库之所以放在这里，是因为其他类别似乎都不适合。*

* [anagent](https://github.com/mudler/anagent) - **Star : 11** 最小化，可插入的Golang evloop/计时器处理程序与依赖注入。
* [antch](https://github.com/antchfx/antch) - **Star : 137** 一个快速、强大和可扩展的web爬行和抓取框架。
* [archiver](https://github.com/mholt/archiver) - **Star : 2350** 用于生成和提取.zip和.tar.gz存档的库和命令。
* [autoflags](https://github.com/artyom/autoflags) - **Star : 24** Go package从struct字段自动定义命令行标志。
* [avgRating](https://github.com/kirillDanshin/avgRating) - **Star : 9** 根据Wilson评分方程计算平均分和评分。
* [banner](https://github.com/dimiro1/banner) - **Star : 228** 在Go应用程序中添加漂亮的横幅。
* [base64Captcha](https://github.com/mojocn/base64Captcha) - **Star : 600** Base64captch支持数字，数字，字母，算术，音频和数字-字母验证码。
* [battery](https://github.com/distatus/battery) - **Star : 134** 跨平台、标准化的电池信息库。
* [bitio](https://github.com/icza/bitio) - **Star : 90** 高度优化的位级读写器。
* [browscap_go](https://github.com/digitalcrab/browscap_go) - **Star : 29** 用于[浏览器功能项目]的GoLang库(http://browscap.org/)。
* [captcha](https://github.com/steambap/captcha) - **Star : 40** 软件包captcha为captcha的生成提供了一个易于使用的、未绑定的API。
* [conv](https://github.com/cstockton/go-conv) - **Star : 340** 包conv提供了跨Go类型的快速和直观的转换。
* [datacounter](https://github.com/miolini/datacounter) - **Star : 27** 读取器/写入器/http.ResponseWriter的计数器。
* [ffmt](https://github.com/go-ffmt/ffmt) - **Star : 126** 美化数据显示为人类。
* [ghorg](https://github.com/gabrie30/ghorg) - **Star : 22** 将所有repos从GitHub org复制到一个目录中。
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - **Star : 642** Golang的通用对象池。
* [go-openapi](https://github.com/go-openapi) - 用于解析和利用开放api模式的包的集合。
* [go-resiliency](https://github.com/eapache/go-resiliency) - **Star : 829** 戈朗的弹性模式。
* [go-unarr](https://github.com/gen2brain/go-unarr) - **Star : 66** 用于RAR、TAR、ZIP和7z存档的解压缩库。
* [gofakeit](https://github.com/brianvoe/gofakeit) - **Star : 402** 用go编写的随机数据生成器。
* [gommit](https://github.com/antham/gommit) - **Star : 65** 分析git提交消息，确保它们遵循已定义的模式。
* [gopsutil](https://github.com/shirou/gopsutil) - **Star : 3778** 用于检索进程和系统利用率(CPU、内存、磁盘等)的跨平台库。
* [gosh](https://github.com/osamingo/gosh) - **Star : 15** 提供Go统计处理程序，结构，测量方法。
* [gosms](https://github.com/haxpax/gosms) - **Star : 1220** 您自己的本地短信网关在Go，可以用来发送短信。
* [gotoprom](https://github.com/cabify/gotoprom) - **Star : 15** 为Prometheus客户端提供类型安全的度量构建器包装库。
* [gountries](https://github.com/pariz/gountries) - **Star : 205** 公开国家和细分数据的包。
* [health](https://github.com/dimiro1/health) - **Star : 359** 易于使用，可扩展的健康检查库。
* [healthcheck](https://github.com/etherlabsio/healthcheck) - **Star : 78** 用于RESTful服务的自以为是的并发健康检查HTTP处理程序。
* [hostutils](https://github.com/Wing924/hostutils) - **Star : 7** 一个用于打包和解包FQDNs列表的golang库。
* [indigo](https://github.com/osamingo/indigo) - **Star : 51** 分布式唯一ID生成器使用Sonyflake，并由Base58编码。
* [lk](https://github.com/hyperboloide/lk) - **Star : 111** 一个简单的golang授权库。
* [llvm](https://github.com/llir/llvm) - **Star : 400** 用于在纯Go中与LLVM IR交互的库。
* [metrics](https://github.com/pascaldekloe/metrics) - **Star : 4** 用于度量仪器和普罗米修斯博览会的库。
* [morse](https://github.com/alwindoss/morse) - **Star : 48** 转换成莫尔斯电码和从莫尔斯电码转换成莫尔斯电码的程序库。
* [numa](https://github.com/lrita/numa) - **Star : 2** NUMA是一个用go编写的实用程序库。它帮助我们编写一些NUMA-AWARED代码。
* [pdfgen](https://github.com/hyperboloide/pdfgen) - **Star : 32** HTTP服务从Json请求生成PDF。
* [persian](https://github.com/mavihq/persian) - **Star : 33** 一些实用的波斯语在围棋。
* [sandid](https://github.com/aofei/sandid) - **Star : 12** 地球上的每一粒沙子都有自己的身份。
* [shellwords](https://github.com/Wing924/shellwords) - **Star : 7** 一个Golang库，根据UNIX Bourne shell的单词解析规则操纵字符串。
* [shortid](https://github.com/teris-io/shortid) - **Star : 442** 分布式生成超短、唯一、非顺序、URL友好的id。
* [stats](https://github.com/go-playground/stats) - **Star : 120** 显示器 Go MemStats +系统统计，如内存，交换和CPU，并通过UDP发送到任何地方，你想记录等…
* [turtle](https://github.com/hackebrot/turtle) - **Star : 73** Emojis Go 。
* [url-shortener](https://github.com/pantrif/url-shortener) - **Star : 16** 一个现代的、强大的、健壮的URL缩短器微服务，支持mysql。
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - 生成样板http输入和输出处理。
* [xdg](https://github.com/rkoesters/xdg) - **Star : 19** FreeDesktop.org (xdg)规范在Go中实现。
* [xkg](https://github.com/go-xkg/xkg) - **Star : 39** X键盘打捞工具。

## 自然语言处理

*用于处理人类语言的库。*

* [getlang](https://github.com/rylans/getlang) - **Star : 70** 快速自然语言检测包。
* [go-eco](https://github.com/ThePaw/go-eco) - **Star : 4** 相似、不相似和距离矩阵;多样性、公平性和不平等度量;物种丰富度估计;coenocline模型。
* [go-i18n](https://github.com/nicksnyder/go-i18n/) - 软件包和用于处理本地化文本的附带工具。
* [go-mystem](https://github.com/dveselov/mystem) - **Star : 23** CGo绑定到Yandex。Mystem -俄罗斯形态学分析仪。
* [go-nlp](https://github.com/nuance/go-nlp) - **Star : 79** 用于处理离散概率分布的实用程序和用于进行NLP工作的其他工具。
* [go-pinyin](https://github.com/mozillazg/go-pinyin) - **Star : 509** 中文汉字到汉语拼音的转换。
* [go-stem](https://github.com/agonopol/go-stem) - **Star : 51** 波特词干算法的实现。
* [go-unidecode](https://github.com/mozillazg/go-unidecode) - **Star : 52** Unicode文本的ASCII音译。
* [go2vec](https://github.com/danieldk/go2vec) - **Star : 30** 用于word2vec嵌入式的阅读器和实用程序函数。
* [gojieba](https://github.com/yanyiwu/gojieba) - **Star : 787** 这是一个Go实现的[jieba](https://github.com/fxsjy/jieba)，这是一个中文分词算法。
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - **Star : 15**  Go 绑定斯诺鲍libstemmer库，包括波特2。
* [gotokenizer](https://github.com/xujiajun/gotokenizer) - **Star : 6** 一种基于字典和双字母格朗语言模型的记号赋予器。(现在只支持中文分割)
* [gounidecode](https://github.com/fiam/gounidecode) - **Star : 67** 用于Go的Unicode音译器(也称为unidecode)。
* [gse](https://github.com/go-ego/gse) - **Star : 1034** 高效的文本分割;支持英语、汉语、日语等。
* [icu](https://github.com/goodsign/icu) - **Star : 19** Cgo绑定用于icu4c C库的检测和转换功能。保证与版本50.1兼容。
* [kagome](https://github.com/ikawaha/kagome) - **Star : 412** JP形态学分析仪编写的纯Go。
* [libtextcat](https://github.com/goodsign/libtextcat) - **Star : 10** 用于libtextcat C库的Cgo绑定。保证与版本2.2兼容。
* [MMSEGO](https://github.com/awsong/MMSEGO) - **Star : 59** 这是一个围棋实现的[MMSEG](http://technology.chtsai.org/mmseg/)，这是一个中文分词算法。
* [nlp](https://github.com/Shixzie/nlp) - **Star : 353** 从字符串中提取值并用nlp填充结构。
* [nlp](https://github.com/james-bowman/nlp) - **Star : 210** 支持LSA(潜在语义分析)的Go自然语言处理库。
* [paicehusk](https://github.com/rookii/paicehusk) - **Star : 25** Golang实现了Paice/外壳阻塞算法。
* [petrovich](https://github.com/striker2000/petrovich) - **Star : 21** 彼得罗维奇是一个图书馆，它把俄语名字的词形变化成特定的语法格。
* [porter](https://github.com/a2800276/porter) - **Star : 8** 这是Martin Porter在C语言中实现的Porter词干分析算法的一个相当简单的移植。
* [porter2](https://github.com/zhenjl/porter2) - 非常快的波特2史坦默。
* [prose](https://github.com/jdkato/prose) - **Star : 2016** 用于支持标记化、词性标记、名称实体提取等文本处理的库。
* [RAKE.go](https://github.com/Obaied/RAKE.go) - 快速自动关键字提取算法(RAKE)的Go端口。
* [segment](https://github.com/blevesearch/segment) - **Star : 47** 如[Unicode标准附件#29]所述，执行Unicode文本分割的Go库(http://www.unicode.org/reports/tr29/)
* [sentences](https://github.com/neurosnap/sentences) - **Star : 260** 句子标记器:将文本转换为句子列表。
* [shamoji](https://github.com/osamingo/shamoji) - **Star : 10** shamoji是用Go编写的word过滤包。
* [snowball](https://github.com/goodsign/snowball) - **Star : 24** 滚雪球柄端口(cgo包装)为围棋。提供词干提取功能[Snowball native](http://snowball.tartarus.org/)。
* [stemmer](https://github.com/dchest/stemmer) - **Star : 47** 用于Go编程语言的Stemmer包。包括英语和德语词根。
* [textcat](https://github.com/pebbe/textcat) - **Star : 60** Go package支持基于n-gram的文本分类，支持utf-8和原始文本。
* [whatlanggo](https://github.com/abadojack/whatlanggo) - **Star : 342** Go的自然语言检测包。支持84种语言和24种脚本(书写系统，如拉丁语、西里尔语等)。
* [when](https://github.com/olebedev/when) - **Star : 922** 带有可插入规则的自然EN和RU语言日期/时间解析器。

## 网络

*用于处理网络各层的库。*

* [arp](https://github.com/mdlayher/arp) - **Star : 190** 包arp实现了arp协议，如RFC 826中所述。
* [buffstreams](https://github.com/stabbycutyou/buffstreams) - **Star : 229** 通过TCP传输协议缓冲区数据变得很容易。
* [canopus](https://github.com/zubairhamed/canopus) - **Star : 133** CoAP客户机/服务器实现(RFC 7252)。
* [cidranger](https://github.com/yl2chen/cidranger) - **Star : 372** 快速IP到CIDR查找围棋。
* [dhcp6](https://github.com/mdlayher/dhcp6) - **Star : 59** 包dhcp6实现了一个DHCPv6服务器，如RFC 3315所述。
* [dns](https://github.com/miekg/dns) - **Star : 3703** 使用DNS的库。
* [ether](https://github.com/songgao/ether) - **Star : 60** 用于发送和接收以太网帧的跨平台Go包。
* [ethernet](https://github.com/mdlayher/ethernet) - **Star : 182** 包以太网实现了对IEEE 802.3以太网II帧和IEEE 802.1Q VLAN标签的编组和解组。
* [fasthttp](https://github.com/valyala/fasthttp) - **Star : 8988** 软件包fasthttp是Go的一个快速HTTP实现，比net/http快10倍。
* [fortio](https://github.com/fortio/fortio) - **Star : 811** 负载测试库和命令行工具，先进的echo服务器和web UI。允许指定一组每秒查询的负载，并记录延迟直方图和其他有用的统计数据，并将它们作图。Tcp、Http、gRPC。
* [ftp](https://github.com/jlaffaye/ftp) - **Star : 484** 包ftp实现了[RFC 959](http://tools.ietf.org/html/rfc959)中描述的ftp客户机。
* [gmqtt](https://github.com/DrmagicE/gmqtt) - **Star : 63** Gmqtt是一个灵活、高性能的MQTT代理库，它完全实现了MQTT协议V3.1.1。
* [gNxI](https://github.com/google/gnxi) - **Star : 98** 一组使用gNMI和gNOI协议的网络管理工具。
* [go-getter](https://github.com/hashicorp/go-getter) - **Star : 684**  Go 图书馆下载文件或目录从各种来源使用一个URL。
* [go-stun](https://github.com/ccding/go-stun) - **Star : 320** Go实现了STUN客户机(RFC 3489和RFC 5389)。
* [gobgp](https://github.com/osrg/gobgp) - **Star : 1643** 用围棋编程语言实现的BGP。
* [golibwireshark](https://github.com/sunwxg/golibwireshark) - **Star : 14** golibwireshark包使用libwireshark库解码pcap文件并分析解剖数据。
* [gopacket](https://github.com/google/gopacket) - **Star : 2767** 使用libpcap绑定访问包处理库。
* [gopcap](https://github.com/akrennmair/gopcap) - **Star : 347**  Go 包装libpcap。
* [goshark](https://github.com/sunwxg/goshark) - **Star : 9** goshark包使用tshark来解码IP包，并创建数据结构来分析包。
* [gosnmp](https://github.com/soniah/gosnmp) - **Star : 424** 用于执行SNMP操作的本机Go库。
* [gotcp](https://github.com/gansidui/gotcp) - **Star : 406**  Go 包快速编写tcp应用程序。
* [grab](https://github.com/cavaliercoder/grab) - **Star : 537**  Go 软件包管理文件下载。
* [graval](https://github.com/koofr/graval) - **Star : 24** 实验FTP服务器框架。
* [HTTPLab](https://github.com/gchaincl/httplab) - **Star : 3364** HTTPLabs允许您检查HTTP请求并伪造响应。
* [iplib](https://github.com/c-robinson/iplib) - **Star : 24** 用于处理IP地址的库(net)。受python [ipaddress](https://docy-doc.org/3/library/ipaddress.html)和ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)的启发
* [jazigo](https://github.com/udhos/jazigo) - **Star : 121** Jazigo是一个用Go编写的工具，用于检索多个网络设备的配置。
* [kcp-go](https://github.com/xtaci/kcp-go) - **Star : 2186** 快速可靠的ARQ协议。
* [kcptun](https://github.com/xtaci/kcptun) - **Star : 10496** 非常简单和快速udp隧道基于KCP协议。
* [lhttp](https://github.com/fanux/lhttp) - **Star : 505** 强大的websocket框架，使您的IM服务器更容易构建。
* [linkio](https://github.com/ian-kent/linkio) - **Star : 43** 用于读写器接口的网络链路速度模拟。
* [llb](https://github.com/kirillDanshin/llb) - **Star : 8** 这是一个非常简单但快速的代理服务器后端。可用于快速重定向到预定义域，具有零内存分配和快速响应。
* [mdns](https://github.com/hashicorp/mdns) - **Star : 539** Golang中的简单mDNS(多播DNS)客户机/服务器库。
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - Paho Go客户机提供一个MQTT客户机库，用于通过TCP、TLS或WebSockets连接到MQTT代理。
* [NFF-Go](https://github.com/intel-go/nff-go) - **Star : 642** 快速开发云计算和裸机网络功能的框架(原YANFF)。
* [packet](https://github.com/aerogo/packet) - **Star : 24** 通过TCP和UDP发送数据包。如果需要，它可以缓冲消息和热交换连接。
* [peerdiscovery](https://github.com/schollz/peerdiscovery) - **Star : 356** 使用UDP组播进行跨平台本地对等点发现的纯Go库。
* [portproxy](https://github.com/aybabtme/portproxy) - **Star : 42** 简单的TCP代理将CORS支持添加到不支持它的API中。
* [publicip](https://github.com/polera/publicip) - **Star : 17** 包publicip返回面向公共的IPv4地址(internet出口)。
* [quic-go](https://github.com/lucas-clemente/quic-go) - **Star : 2788** 在纯Go中实现了QUIC协议。
* [raw](https://github.com/mdlayher/raw) - **Star : 295** Package raw支持在设备驱动程序级别读取和写入网络接口的数据。
* [sftp](https://github.com/pkg/sftp) - **Star : 710** 包sftp实现了SSH文件传输协议，如https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt所述。
* [ssh](https://github.com/gliderlabs/ssh) - **Star : 1084** 用于构建SSH服务器的高级API(封装密码/ SSH)。
* [sslb](https://github.com/eduardonunesp/sslb) - **Star : 112** 它是一个超级简单的负载平衡器，只是一个实现某种性能的小项目。
* [stun](https://github.com/go-rtc/stun) - Go实现的RFC 5389 STUN协议。
* [tcp_server](https://github.com/firstrow/tcp_server) - **Star : 273**  Go 图书馆建设tcp服务器更快。
* [tspool](https://github.com/two/tspool) - **Star : 5** TCP库使用工作池来提高性能并保护服务器。
* [utp](https://github.com/anacrolix/utp) - **Star : 149** Go uTP微传输协议的实现。
* [water](https://github.com/songgao/water) - **Star : 819** 简单TUN / TAP图书馆。
* [webrtc](https://github.com/pions/webrtc) - WebRTC API的纯Go实现。
* [winrm](https://github.com/masterzen/winrm) - **Star : 206**  Go WinRM客户端远程执行Windows机器上的命令。
* [xtcp](https://github.com/xfxdev/xtcp) - **Star : 79** TCP服务器框架具有同时全双工通信，优雅关机，自定义协议。

### HTTP客户端

*用于发出HTTP请求的库。*

* [gentleman](https://github.com/h2non/gentleman) - **Star : 669** 功能齐全的插件驱动HTTP客户端库。
* [goreq](https://github.com/smallnest/goreq) - **Star : 98** 基于gorequest的增强简化HTTP客户机。
* [grequests](https://github.com/levigross/grequests) - **Star : 1386** 一个 Go “克隆”的伟大和著名的请求库。
* [heimdall](https://github.com/gojektech/heimdall) - 具有重试和hystrix功能的增强http客户机。
* [pester](https://github.com/sethgrid/pester) - **Star : 318** 使用重试、后退和并发执行HTTP客户机调用。
* [rq](https://github.com/ddo/rq) - **Star : 25** golang stdlib HTTP客户端更好的接口。
* [sling](https://github.com/dghubble/sling) - **Star : 844** Sling是一个用于创建和发送API请求的Go HTTP客户端库。

## OpenGL

*用于在Go中使用OpenGL的库。*

* [gl](https://github.com/go-gl/gl) - **Star : 626** OpenGL的Go绑定(通过glow生成)。
* [glfw](https://github.com/go-gl/glfw) - **Star : 706** Go绑定用于glfw3。
* [goxjs/gl](https://github.com/goxjs/gl) - **Star : 130** 跨平台的OpenGL绑定(OS X, Linux, Windows，浏览器，iOS, Android)。
* [goxjs/glfw](https://github.com/goxjs/glfw) - **Star : 58** 使用跨平台glfw库创建OpenGL上下文并接收事件。
* [mathgl](https://github.com/go-gl/mathgl) - **Star : 286** 纯Go数学软件包专门为三维数学，与灵感来自GLM。

## ORM

*实现对象关系映射或数据映射技术的库。*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - 强大的orm框架。支持:pq / mysql / sqlite3。
* [go-firestorm](https://github.com/jschoedt/go-firestorm) - 一个用于谷歌/Firebase云Firestore的简单ORM。
* [go-pg](https://github.com/go-pg/pg) - **Star : 2854** 关注PostgreSQL的特性和性能。
* [go-queryset](https://github.com/jirfag/go-queryset) - **Star : 441** 100%类型安全ORM与代码生成和MySQL, PostgreSQL, Sqlite3, SQL Server支持基于GORM。
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - **Star : 218** 一个灵活而强大的SQL字符串构建器库加上一个零配置ORM。
* [go-store](https://github.com/gosuri/go-store) - **Star : 92** 简单而快速的Redis支持的键值存储库。
* [GORM](https://github.com/jinzhu/gorm) - **Star : 14057** Golang出色的ORM库的目标是对开发人员友好。
* [gorp](https://github.com/go-gorp/gorp) - **Star : 3057** Go的关系持久性，ORM-ish库。
* [grimoire](https://github.com/Fs02/grimoire) - **Star : 110** Grimoire是golang的数据库访问层和验证。(支持:MySQL, PostgreSQL和SQLite3)。
* [lore](https://github.com/abrahambotros/lore) - **Star : 4** 用于Go的简单轻量级伪orm /伪结构映射环境。
* [Marlow](https://github.com/dadleyy/marlow) - **Star : 60** 从项目结构生成ORM，用于编译时安全保证。
* [pop/soda](https://github.com/gobuffalo/pop) - **Star : 651** 数据库迁移、创建、ORM等。用于MySQL、PostgreSQL和SQLite。
* [QBS](https://github.com/coocood/qbs) - **Star : 536** 表示结构查询。一个ORM。
* [reform](https://github.com/go-reform/reform) - **Star : 783** 更好的ORM for Go，基于非空接口和代码生成。
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) - **Star : 2194** ORM生成器。根据您的数据库模式生成一个功能强大且运行速度快的ORM。
* [upper.io/db](https://github.com/upper/db) - **Star : 1809** 通过使用封装成熟数据库驱动程序的适配器与不同数据源交互的单一接口。
* [Xorm](https://github.com/go-xorm/xorm) - **Star : 4973** 简单而强大的ORM for Go。
* [Zoom](https://github.com/albrow/zoom) - **Star : 237** 基于Redis的快速数据存储和查询引擎。

## 包管理

*用于依赖关系和包管理的官方工具*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - 模块是源代码交换和版本控制的单元。go命令直接支持处理模块，包括记录和解决对其他模块的依赖关系。

*包管理的官方实验工具*

* [dep](https://github.com/golang/dep) - **Star : 12357**  Go 依赖的工具。
* [vgo](https://go.googlesource.com/vgo/) - 版本化 Go 。

*用于包和依赖项管理的非官方库。*

* [gigo](https://github.com/LyricalSecurity/gigo) - **Star : 196** 类似pip的golang依赖工具，支持私有存储库和散列。
* [glide](https://github.com/Masterminds/glide) - **Star : 7732** 轻松管理您的golang供应商和销售包。受Maven、Bundler和Pip等工具的启发。
* [godep](https://github.com/tools/godep) - **Star : 5652** godep是go的依赖工具，它通过修复包的依赖关系来帮助构建可重复的包。
* [gom](https://github.com/mattn/gom) - **Star : 1353** Go Manager - bundle for Go。
* [goop](https://github.com/nitrous-io/goop) - Go (golang)的简单依赖管理器，灵感来自Bundler。
* [gop](https://github.com/lunny/gop) - **Star : 50** 在GOPATH之外构建和管理Go应用程序。
* [gopm](https://github.com/gpmgo/gopm) - **Star : 2322** 包管理器。
* [govendor](https://github.com/kardianos/govendor) - **Star : 4650** 包管理器。使用标准供应商文件的Go vendor工具。
* [gpm](https://github.com/pote/gpm) - **Star : 1204** 基本的Go依赖管理器。
* [johnny-deps](https://github.com/VividCortex/johnny-deps) - **Star : 213** 使用Git的最小依赖版本。
* [mvn-golang](https://github.com/raydac/mvn-golang) - **Star : 85** 插件，为自动加载Golang SDK，依赖关系管理和启动Maven项目基础设施中的构建环境提供了方法。
* [nut](https://github.com/jingweno/nut) - **Star : 245** 供应商的依赖。
* [VenGO](https://github.com/DamnWidget/VenGO) - **Star : 116** 创建和管理可导出的隔离go虚拟环境。

## 查询语言

* [gojsonq](https://github.com/thedevsaddam/gojsonq) - **Star : 827** 一个简单的Go包来查询JSON数据。
* [graphql](https://github.com/tmc/graphql) - **Star : 51** graphql解析器+实用程序。
* [graphql](https://github.com/neelance/graphql-go) - 关注易用性的GraphQL服务器。
* [graphql-go](https://github.com/graphql-go/graphql) - **Star : 5034** 为Go实现GraphQL。
* [jsonql](https://github.com/elgs/jsonql) - **Star : 201** Golang中的JSON查询表达式库。
* [jsonslice](https://github.com/bhmj/jsonslice) - **Star : 22** 使用高级过滤器查询Jsonpath。
* [rql](https://github.com/a8m/rql) - **Star : 108** 用于REST API的资源查询语言。

## 嵌入的资源

* [esc](https://github.com/mjibson/esc) - **Star : 455** 将文件嵌入到Go程序中并提供http。文件系统接口到它们。
* [fileb0x](https://github.com/UnnoTed/fileb0x) - **Star : 413** 简单的工具嵌入文件 Go 与重点“定制”和易于使用。
* [go-embed](https://github.com/pyros2097/go-embed) - **Star : 14** 生成go代码，将资源文件嵌入到库或可执行文件中。
* [go-resources](https://github.com/omeid/go-resources) - **Star : 154** 嵌入到Go中的普通资源。
* [go.rice](https://github.com/GeertJohan/go.rice) - **Star : 1614** 走了。rice是一个Go包，它使处理html、js、css、图像和模板等资源变得非常容易。
* [packr](https://github.com/gobuffalo/packr) - **Star : 1997** 将静态文件嵌入到Go二进制文件中的简单方法。
* [statics](https://github.com/go-playground/statics) - **Star : 53** 将静态资源嵌入到go文件中，用于单个二进制编译+使用http。文件系统+符号链接。
* [statik](https://github.com/rakyll/statik) - **Star : 2012** 将静态文件嵌入到Go可执行文件中。
* [templify](https://github.com/wlbr/templify) - **Star : 18** 将外部模板文件嵌入到Go代码中，以创建单个文件二进制文件。
* [vfsgen](https://github.com/shurcooL/vfsgen) - **Star : 628** 生成一个vfsdata。静态实现给定虚拟文件系统的go文件。

## 科学与数据分析

*用于科学计算和数据分析的库。*

* [assocentity](https://github.com/ndabAP/assocentity) - **Star : 3** 包assocentity返回单词到给定实体的平均距离。
* [bradleyterry](https://github.com/seanhagen/bradleyterry) - 为成对比较提供了一个布莱德利-特里模型。
* [chart](https://github.com/vdobler/chart) - **Star : 572** 简单的图表绘制库。支持多种图形类型。
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - **Star : 56** 用于机器学习和统计的数据模型(类似于熊猫)。
* [evaler](https://github.com/soniah/evaler) - **Star : 40** 简单的浮点算术表达式求值器。
* [ewma](https://github.com/VividCortex/ewma) - **Star : 263** Exponentially-weighted移动平均线。
* [geom](https://github.com/skelterjohn/geom) - **Star : 40** 戈朗的二维几何。
* [go-dsp](https://github.com/mjibson/go-dsp) - **Star : 623** Go数字信号处理。
* [go-fn](https://github.com/ematvey/go-fn) - **Star : 11** 用Go语言编写的数学函数，不包括在math pkg中。
* [go-gt](https://github.com/ThePaw/go-gt) - **Star : 5** 用“Go”语言编写的图论算法。
* [gocomplex](https://github.com/varver/gocomplex) - **Star : 5** 用于围棋编程语言的复数库。
* [goent](https://github.com/kzahedi/goent) - **Star : 13**  Go 实现熵度量。
* [gohistogram](https://github.com/VividCortex/gohistogram) - **Star : 126** 数据流的近似直方图。
* [gonum](https://github.com/gonum/gonum) - **Star : 2830** Gonum是一组用于Go编程语言的数字库。它包含用于矩阵、统计、优化等的库。
* [gonum/plot](https://github.com/gonum/plot) - **Star : 1173** gonum/plot提供了一个API，用于在Go中构建和绘制绘图。
* [goraph](https://github.com/gyuho/goraph) - **Star : 596** 纯Go图论库(数据结构，算法可视化)。
* [gosl](https://github.com/cpmech/gosl) - **Star : 1284**  Go 科学图书馆线性代数，FFT，几何，NURBS，数值方法，概率，优化，微分方程，等等。
* [GoStats](https://github.com/OGFris/GoStats) - **Star : 9** GoStats是一个开放源码的GoLang库，主要用于机器学习领域的数学统计，它涵盖了大多数统计度量函数。
* [graph](https://github.com/yourbasic/graph) - **Star : 223** 基本图形算法库。
* [ode](https://github.com/ChristopherRabotin/ode) - **Star : 10** 常微分方程(ODE)求解器，支持扩展状态和基于信道的迭代停止条件。
* [orb](https://github.com/paulmach/orb) - **Star : 180** 2D几何类型，支持剪切、GeoJSON和Mapbox矢量平铺。
* [pagerank](https://github.com/alixaxel/pagerank) - **Star : 48** 加权PageRank算法在Go中实现。
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) - **Star : 5** 微型线性插值库。
* [PiHex](https://github.com/claygod/PiHex) - **Star : 9** 实现了针对16进制数Pi的“bailee - borwein - plouffe”算法。
* [rootfinding](https://github.com/khezen/rootfinding) - **Star : 3** 二次函数求根算法库。
* [sparse](https://github.com/james-bowman/sparse) - **Star : 63**  Go 稀疏矩阵格式的线性代数支持科学和机器学习应用程序，兼容gonum矩阵库。
* [stats](https://github.com/montanaflynn/stats) - **Star : 1326** 包含Golang标准库中缺少的公共函数的统计软件包。
* [streamtools](https://github.com/nytlabs/streamtools) - **Star : 1315** 通用图形工具，用于处理数据流。
* [TextRank](https://github.com/DavidBelicza/TextRank) - **Star : 65** TextRank在Golang中的实现，支持扩展特性(摘要、加权、短语提取)和多线程(goroutine)。
* [triangolatte](https://github.com/tchayen/triangolatte) - **Star : 11** 二维三角库。允许将线和多边形(都基于点)转换为gpu语言。

## 安全

*用于帮助您的应用程序更安全的库。*

* [acmetool](https://github.com/hlandau/acme) - **Star : 1690** ACME(让我们用自动更新加密)客户端工具。
* [acra](https://github.com/cossacklabs/acra) - **Star : 435** 网络加密代理保护基于数据库的应用程序免受数据泄漏:强选择性加密，SQL注入预防，入侵检测系统。
* [argon2pw](https://github.com/raja/argon2pw) - **Star : 71** 使用常量时间密码比较生成Argon2密码散列。
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - 让我们加密证书并启动TLS服务器。
* [BadActor](https://github.com/jaredfolkins/badactor) - **Star : 241** 内存中，应用程序驱动的jailer基于fail2ban的精神构建。
* [Cameradar](https://github.com/Ullaakut/cameradar) - **Star : 1763** 工具和库，以远程入侵RTSP流从监控摄像头。
* [certificates](https://github.com/mvmaasakkers/certificates) - **Star : 6** 用于生成tls证书的自定义工具。
* [go-yara](https://github.com/hillu/go-yara) - **Star : 132** Go Bindings for [YARA](https://github.com/plusvic/yara)，“用于恶意软件研究人员(和其他人)的模式匹配瑞士刀”。
* [goArgonPass](https://github.com/dwin/goArgonPass) - **Star : 10** Argon2密码散列和验证设计为与现有Python和PHP实现兼容。
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - **Star : 28** 一个安全哈希和加密密码的偏执包。
* [Interpol](https://bitbucket.org/vahidi/interpol) - 基于规则的数据生成器，用于模糊和渗透测试。
* [jwc](https://github.com/khezen/jwc) - **Star : 5** JSON Web加密库。
* [lego](https://github.com/xenolf/lego) - 纯Go ACME客户端库和CLI工具(用于加密)。
* [memguard](https://github.com/awnumar/memguard) - **Star : 992** 一个用于处理内存中敏感值的纯Go库。
* [nacl](https://github.com/kevinburke/nacl) - **Star : 450**  Go 实现NaCL API的集合。
* [passlib](https://github.com/hlandau/passlib) - **Star : 224** 未来证明密码哈希库。
* [secure](https://github.com/unrolled/secure) - **Star : 1182** 用于Go的HTTP中间件促进了一些快速的安全胜利。
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) - **Star : 153** Scrypt包具有简单，明显的API和自动成本校准内置。
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) - **Star : 194** 使用ssh密钥加密/解密。
* [sslmgr](https://github.com/adrianosela/sslmgr) - **Star : 7** 使用围绕acme/autocert的高级包装器，SSL证书变得很容易。

## 序列化

*用于二进制序列化的库和工具。*

* [asn1](https://github.com/PromonLogicalis/asn1) - 面向golang的BER和DER编码库。
* [bambam](https://github.com/glycerine/bambam) - **Star : 61** 为船长原型模式从go生成器。
* [bel](https://github.com/32leaves/bel) - **Star : 4** 从Go structs/interface生成TypeScript接口。对JSON RPC很有用。
* [binstruct](https://github.com/ghostiam/binstruct) - **Star : 7** 用于将数据映射到结构中的Golang二进制解码器。
* [colfer](https://github.com/pascaldekloe/colfer) - **Star : 465** 为Colfer二进制格式生成代码。
* [csvutil](https://github.com/jszwec/csvutil) - **Star : 299** 高性能、惯用的CSV记录编码和解码到本机Go结构。
* [fwencoder](https://github.com/o1egl/fwencoder) - **Star : 6** 用于Go的固定宽度文件解析器(编码和解码库)。
* [go-capnproto](https://github.com/glycerine/go-capnproto) - **Star : 273** 船长的原始库和解析器 Go 。
* [go-codec](https://github.com/ugorji/go) - **Star : 1204** 高性能，功能丰富，惯用的编码，解码和rpc库的msgpack, cbor和json，基于运行时或代码生成的支持。
* [gogoprotobuf](https://github.com/gogo/protobuf) - **Star : 2812** 用于与gadget一起使用的协议缓冲区。
* [goprotobuf](https://github.com/golang/protobuf) - **Star : 4817** Go以库和协议编译器插件的形式支持谷歌的协议缓冲区。
* [jsoniter](https://github.com/json-iterator/go) - **Star : 5222** 高性能100%兼容的drop-in替换“编码/json”。
* [mapstructure](https://github.com/mitchellh/mapstructure) - **Star : 2306** 用于将通用映射值解码为本机Go结构的Go库。
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - **Star : 118** 用于处理PHP会话格式和PHP序列化/反序列化函数的GoLang库。
* [structomap](https://github.com/tuvistavie/structomap) - 库，以方便地从静态结构动态生成映射。

## 服务器应用程序

* [algernon](https://github.com/xyproto/algernon) - **Star : 1570** 内置支持Lua、Markdown、GCSS和Amber的HTTP/2 web服务器。
* [Caddy](https://github.com/mholt/caddy) - **Star : 22407** Caddy是另一种HTTP/2 web服务器，易于配置和使用。
* [consul](https://www.consul.io/) - 领事是用于服务发现、监视和配置的工具。
* [devd](https://github.com/cortesi/devd) - **Star : 2781** 为开发人员提供本地web服务器。
* [discovery](https://github.com/Bilibili/discovery) - **Star : 617** 用于弹性中间层负载平衡和故障转移的注册表。
* [etcd](https://github.com/coreos/etcd) - 为共享配置和服务发现提供高可用的键值存储。
* [Fider](https://github.com/getfider/fider) - **Star : 765** Fider是一个收集和组织客户反馈的开放平台。
* [Flagr](https://github.com/checkr/flagr) - **Star : 774** Flagr是一个开源特性标记和A/B测试服务。
* [flipt](https://github.com/markphelps/flipt) - **Star : 974** 一个用Go和Vue.js编写的自包含特性标志解决方案
* [jackal](https://github.com/ortuman/jackal) - **Star : 703** 用Go编写的XMPP服务器。
* [minio](https://github.com/minio/minio) - **Star : 16748** Minio是一个分布式对象存储服务器。
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) - **Star : 5** Nginx日志解析器和出口到普罗米修斯。
* [nsq](http://nsq.io/) - 一个实时分布式消息平台。
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) - **Star : 4** 从PostgreSQL到Kafka的流数据库事件。
* [riemann-relay](https://github.com/blind-oracle/riemann-relay) - 传递到负载平衡黎曼事件并/或将其转换为碳。
* [RoadRunner](https://github.com/spiral/roadrunner) - **Star : 3115** 高性能PHP应用服务器，负载平衡器和进程管理器。
* [yakvs](https://git.sci4me.com/sci4me/yakvs) - 小型、网络化、内存中的键值存储。

## 模板引擎

*用于模板和词法分析的库和工具。*

* [ace](https://github.com/yosssi/ace) - **Star : 761** Ace是一个用于Go的HTML模板引擎，灵感来自Slim和Jade。Ace是金子的提炼品。
* [amber](https://github.com/eknkc/amber) - **Star : 821** Amber是一个优雅的Go编程语言模板引擎，它的灵感来自HAML和Jade。
* [damsel](https://github.com/dskinner/damsel) - **Star : 20** 标记语言，通过css选择器提供html大纲，通过pkg html/template和其他工具进行扩展。
* [ego](https://github.com/benbjohnson/ego) - **Star : 407** 轻量级模板语言，允许您在Go中编写模板。模板被翻译成Go并编译。
* [extemplate](https://github.com/dannyvankooten/extemplate) - **Star : 12** 围绕html/模板的小包装器，允许简单的基于文件的模板继承。
* [fasttemplate](https://github.com/valyala/fasttemplate) - **Star : 286** 简单而快速的模板引擎。替换模板占位符的速度比[text/template](http://golang.org/pkg/text/template/)快10倍。
* [gofpdf](https://github.com/jung-kurt/gofpdf) - **Star : 2989** PDF文档生成器具有对文本、绘图和图像的高级支持。
* [goview](https://github.com/foolin/goview) - **Star : 36** Goview是一个轻量级、极简和惯用的模板库，基于golang html/template构建Go web应用程序。
* [hero](https://github.com/shiyanhui/hero) - **Star : 1188** Hero是一个方便、快速和强大的go模板引擎。
* [jet](https://github.com/CloudyKit/jet) - **Star : 576** Jet模板引擎。
* [kasia.go](https://github.com/ziutek/kasia.go) - **Star : 70** 模板系统的HTML和其他文本文件- Go 实现。
* [liquid](https://github.com/osteele/liquid) - **Star : 80**  Go 实现Shopify液体模板。
* [mustache](https://github.com/hoisie/mustache) - **Star : 963** Go实现了Mustache模板语言。
* [pongo2](https://github.com/flosch/pongo2) - **Star : 1476** 类似于django的模板引擎。
* [quicktemplate](https://github.com/valyala/quicktemplate) - **Star : 1354** 快速，强大，但易于使用模板引擎。将模板转换为Go代码，然后编译它。
* [raymond](https://github.com/aymerick/raymond) - **Star : 335** 完成手柄在Go中的实现。
* [Razor](https://github.com/sipin/gorazor) - **Star : 670** 戈朗剃刀视图引擎。
* [Soy](https://github.com/robfig/soy) - **Star : 143** Go的闭包模板(又名大豆模板)，遵循[官方规范](https://developer.google.com/closure/templates/)。
* [velvet](https://github.com/gobuffalo/velvet) - **Star : 65** 完成手柄在Go中的实现。

## 测试

*用于测试代码库和生成测试数据的库。*

* Testing Frameworks
    * [assert](https://github.com/go-playground/assert) - **Star : 13** 与本机go测试一起使用的基本断言库，以及用于自定义断言的构建块。
    * [badio](https://github.com/cavaliercoder/badio) - **Star : 8** 扩展到Go的' test /iotest '包。
    * [baloo](https://github.com/h2non/baloo) - **Star : 633** 富有表现力和多用途的端到端HTTP API测试变得很容易。
    * [biff](https://github.com/fulldump/biff) - **Star : 6** 分岔测试框架，BDD兼容。
    * [bro](https://github.com/marioidival/bro) - **Star : 25** 查看目录中的文件并为它们运行测试。
    * [charlatan](https://github.com/percolate/charlatan) - **Star : 188** 为测试生成假接口实现的工具。
    * [commander](https://github.com/SimonBaeumer/commander) - **Star : 31** 用于在windows、linux和osx上测试cli应用程序的工具。
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) - **Star : 77** 测试框架的简单快照测试插件。
    * [dbcleaner](https://github.com/khaiql/dbcleaner) - **Star : 78** Clean database for testing purpose，灵感来自Ruby中的“database_cleaner”。
    * [dsunit](https://github.com/viant/dsunit) - **Star : 24** 用于SQL、NoSQL、结构化文件的数据存储测试。
    * [endly](https://github.com/viant/endly) - **Star : 83** 声明性端到端功能测试。
    * [frisby](https://github.com/verdverm/frisby) - **Star : 245** REST API测试框架。
    * [ginkgo](http://onsi.github.io/ginkgo/) - Go的BDD测试框架。
    * [go-carpet](https://github.com/msoap/go-carpet) - **Star : 194** 在终端中查看测试覆盖率的工具。
    * [go-cmp](https://github.com/google/go-cmp) - **Star : 1079** 用于比较测试中的Go值的包。
    * [go-mutesting](https://github.com/zimmski/go-mutesting) - **Star : 244** 变异测试的Go源代码。
    * [go-testdeep](https://github.com/maxatome/go-testdeep) - **Star : 49** 极具灵活性的golang深度比较，扩展了go测试包。
    * [go-vcr](https://github.com/dnaeon/go-vcr) - **Star : 324** 记录并回放HTTP交互，以便进行快速、确定和准确的测试。
    * [goblin](https://github.com/franela/goblin) - **Star : 613** 摩卡喜欢测试框架fo Go。
    * [gocheck](http://labix.org/gocheck) - 更先进的测试框架替代gotest。
    * [GoConvey](https://github.com/smartystreets/goconvey/) - bdd风格的框架与web UI和实时重载。
    * [gocrest](https://github.com/corbym/gocrest) - **Star : 8** 用于围棋断言的可组合的类仓鼠匹配器。
    * [godog](https://github.com/DATA-DOG/godog) - **Star : 700** Cucumber或Behat类似于Go的BDD框架。
    * [gofight](https://github.com/appleboy/gofight) - **Star : 251** 用于Golang路由器框架的API处理程序测试。
    * [gogiven](https://github.com/corbym/gogiven) - **Star : 7** 类似于yatspec的Go BDD测试框架。
    * [gomatch](https://github.com/jfilipczyk/gomatch) - **Star : 29** 为针对模式测试JSON而创建的库。
    * [gomega](http://onsi.github.io/gomega/) - Rspec类似于匹配器/断言库。
    * [GoSpec](https://github.com/orfjackal/gospec) - 用于围棋编程语言的bdd风格的测试框架。
    * [gospecify](https://github.com/stesla/gospecify) - **Star : 51** 这为测试Go代码提供了一个BDD语法。任何使用过rspec之类库的人都应该熟悉它。
    * [gosuite](https://github.com/pavlo/gosuite) - **Star : 9** 通过利用Go1.7的子测试，为“测试”带来带有设置/拆卸功能的轻量级测试套件。
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) - **Star : 106** 一组包，用于增强go测试包并支持公共模式。
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - **Star : 26** 用于声明性匹配器对象的连贯框架，当将其应用于输入值时，将产生自描述结果。
    * [httpexpect](https://github.com/gavv/httpexpect) - **Star : 1100** 简洁、声明性、易于使用端到端HTTP和REST API测试。
    * [jsonassert](https://github.com/kinbiko/jsonassert) - **Star : 20** 用于验证JSON有效负载已正确序列化的包。
    * [restit](https://github.com/yookoala/restit) - Go micro framework帮助编写RESTful API集成测试。
    * [testcase](https://github.com/adamluzsi/testcase) - **Star : 9** 行为驱动开发的惯用测试框架。
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) - **Star : 313** 一个帮助Rails的测试装置来测试数据库应用程序。
    * [Testify](https://github.com/stretchr/testify) - **Star : 7747** 对标准go测试包的神圣扩展。
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd) - 将markdown代码段转换为可测试的go代码。
    * [testsql](https://github.com/zhulongcheng/testsql) - **Star : 7** 在测试前从SQL文件生成测试数据，并在测试完成后清除数据。
    * [Tt](https://github.com/vcaesar/tt) - **Star : 5** 简单而丰富多彩的测试工具。
    * [wstest](https://github.com/posener/wstest) - **Star : 62** 用于单元测试Websocket http.Handler的Websocket客户机。

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - **Star : 353** 用于生成自包含的模拟对象的工具。
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - **Star : 1604** 用于测试数据库交互的模拟SQL驱动程序。
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - **Star : 152** 基于单事务的数据库驱动程序主要用于测试目的。
    * [gock](https://github.com/h2non/gock) - **Star : 781** 通用HTTP模拟变得容易。
    * [gomock](https://github.com/golang/mock) - **Star : 2656** 用于Go编程语言的mock框架。
    * [govcr](https://github.com/seborama/govcr) - **Star : 78** Golang的HTTP模拟:记录和回放HTTP交互以进行离线测试。
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) - **Star : 1406** 使用可扩展中间件和易于使用的CLI记录和模拟REST/SOAP api的HTTP(S)代理。
    * [httpmock](https://github.com/jarcoal/httpmock) - **Star : 553** 轻松模拟来自外部资源的HTTP响应。
    * [minimock](https://github.com/gojuno/minimock) - **Star : 193** Go接口的模拟生成器。
    * [mockhttp](https://github.com/tv42/mockhttp) - **Star : 22** Go http.ResponseWriter的模拟对象。

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) - **Star : 2833** 随机测试系统。
    * [gofuzz](https://github.com/google/gofuzz) - **Star : 514** 库，用于填充具有随机值的go对象。
    * [Tavor](https://github.com/zimmski/tavor) - **Star : 208** 通用模糊和delta调试框架。

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp) - **Star : 341** 用于Chrome调试协议的类型安全绑定，可与实现该协议的浏览器或其他调试目标一起使用。
    * [chromedp](https://github.com/knq/chromedp) - 一种驱动/测试Chrome、Safari、Edge、Android webview和其他支持Chrome调试协议的浏览器的方法。
    * [ggr](https://github.com/aerokube/ggr) - **Star : 207** 一个轻量级服务器，它将Selenium WebDriver请求路由和代理到多个Selenium集线器。
    * [selenoid](https://github.com/aerokube/selenoid) - **Star : 1165** 在容器中启动浏览器的替代Selenium hub服务器。

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) - **Star : 363** 为Golang实现[failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail)。

## 文本处理

*用于解析和操作文本的库。*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align) - **Star : 57** 对文本进行对齐的通用应用程序。
    * [allot](https://github.com/sbstjn/allot) - **Star : 33** 用于CLI工具和机器人的占位符和通配符文本解析。
    * [bbConvert](https://github.com/CalebQ42/bbConvert) - **Star : 5** 将bbCode转换为HTML，使您可以添加对自定义bbCode标记的支持。
    * [blackfriday](https://github.com/russross/blackfriday) - **Star : 3789** 降价处理器在围棋。
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - **Star : 1210** HTML洗手液。
    * [codetree](https://github.com/aerogo/codetree) - **Star : 6** 解析缩进代码(python、pixy、scarlet等)并返回树结构。
    * [colly](https://github.com/asciimoo/colly) - 快速和优雅的刮地框架为地鼠。
    * [commonregex](https://github.com/mingrammer/commonregex) - **Star : 544** 一组用于Go的公共正则表达式。
    * [dataflowkit](https://github.com/slotix/dataflowkit) - **Star : 276** Web抓取框架将网站转换为结构化数据。
    * [did](https://github.com/ockam-network/did) - **Star : 21** DID(分散标识符)解析器和Stringer。
    * [doi](https://github.com/hscells/doi) - **Star : 4** 文档对象标识符(doi)解析器。
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - **Star : 36** Editorconfig文件解析器和Go操作器。
    * [enca](https://github.com/endeveit/enca) - **Star : 7** [libenca](http://cihar.com/software/enca/)的最小cgo绑定。
    * [encdec](https://github.com/mickep76/encdec) - 软件包为编码器和解码器提供了通用接口。
    * [genex](https://github.com/alixaxel/genex) - **Star : 50** 将正则表达式计数并展开为所有匹配的字符串。
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub风格的Markdown渲染器(使用blackfriday)，带有带保护的代码块高亮显示，可单击头锚链接。
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - **Star : 21** 固定宽度的文本格式(带反射的编码器/解码器)。
    * [go-humanize](https://github.com/dustin/go-humanize) - **Star : 1846** 格式化程序，用于将时间、数字和内存大小转换为可读格式。
    * [go-nmea](https://github.com/adrianmo/go-nmea) - **Star : 90** 用于Go语言的NMEA解析器库。
    * [go-runewidth](https://github.com/mattn/go-runewidth) - **Star : 207** 函数获取字符或字符串的固定宽度。
    * [go-slugify](https://github.com/mozillazg/go-slugify) - **Star : 27** 使多种语言的支持相当鼻涕虫。
    * [go-toml](https://github.com/pelletier/go-toml) - **Star : 588** 使用带有查询支持和方便的cli工具的TOML格式库。
    * [go-vcard](https://github.com/emersion/go-vcard) - **Star : 23** 解析和格式化vCard。
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) - **Star : 41** 用于Go的零宽度字符检测和删除。
    * [gofeed](https://github.com/mmcdole/gofeed) - **Star : 1075** 在Go中解析RSS和Atom提要。
    * [gographviz](https://github.com/awalterschulze/gographviz) - **Star : 282** 解析Graphviz点语言。
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - 将字节格式化为字符串。
    * [gonameparts](https://github.com/polera/gonameparts) - **Star : 30** 将人名解析为单独的名称部分。
    * [goq](https://github.com/andrewstuart/goq) - **Star : 141** 使用带有jQuery语法的struct标记对HTML进行声明式解组(使用GoQuery)。
    * [GoQuery](https://github.com/PuerkitoBio/goquery) - **Star : 7384** GoQuery为Go语言带来了类似于jQuery的语法和一组特性。
    * [goregen](https://github.com/zach-klippenstein/goregen) - **Star : 35** 库，用于从正则表达式生成随机字符串。
    * [gotext](https://github.com/leonelquinteros/gotext) - **Star : 228** 用于Go的GNU gettext实用程序。
    * [guesslanguage](https://github.com/endeveit/guesslanguage) - **Star : 44** 函数确定unicode文本的自然语言。
    * [htmlquery](https://github.com/antchfx/htmlquery) - **Star : 112** 用于HTML的XPath查询包，允许您通过XPath表达式从HTML文档中提取数据或求值。
    * [inject](https://github.com/facebookgo/inject) - 包注入提供了一个基于反射的注入器。
    * [ltsv](https://github.com/Wing924/ltsv) - **Star : 2** 用于Go的高性能[LTSV(标签为Tab Separeted Value)](http://ltsv.org/)阅读器。
    * [mxj](https://github.com/clbanning/mxj) - **Star : 323** 将XML编码/解码为JSON或map[string]接口{};使用点符号路径和通配符提取值。替换x2j和j2x包。
    * [sdp](https://github.com/gortc/sdp) - **Star : 66** SDP:会话描述协议[[RFC 4566](https://tools.ietf.org/html/rfc4566)]。
    * [sh](https://github.com/mvdan/sh) - **Star : 1912** Shell解析器和格式化程序。
    * [slug](https://github.com/gosimple/slug) - **Star : 364** 支持多种语言的url友好的slugify。
    * [Slugify](https://github.com/avelino/slugify) - **Star : 26**  Go slugify处理字符串的应用程序。
    * [syndfeed](https://github.com/zhengchun/syndfeed) - **Star : 4** Atom 1.0和RSS 2.0的联合提要。
    * [toml](https://github.com/BurntSushi/toml) - **Star : 2696** TOML配置格式(带反射的编码器/解码器)。
* Utility
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) - **Star : 15** 一个基于消毒的Go脏话过滤器。
    * [gotabulate](https://github.com/bndr/gotabulate) - **Star : 197** 使用Go轻松漂亮地打印表格数据。
    * [kace](https://github.com/codemodus/kace) - **Star : 12** 覆盖公共初始化的公共用例转换。
    * [parseargs-go](https://github.com/nproc/parseargs-go) - 理解引号和反斜杠的字符串参数分析器。
    * [parth](https://github.com/codemodus/parth) - **Star : 31** URL路径分割解析。
    * [radix](https://github.com/yourbasic/radix) - **Star : 139** 快速字符串排序算法。
    * [Tagify](https://github.com/zoomio/tagify) - 从给定源生成一组标记。
    * [TySug](https://github.com/Dynom/TySug) - **Star : 3** 关于键盘布局的其他建议。
    * [xj2go](https://github.com/stackerzzq/xj2go) - **Star : 17** 将xml或json转换为struct。
    * [xurls](https://github.com/mvdan/xurls) - **Star : 453** 从文本中提取url。

## 第三方api

*用于访问第三方api的库。*

* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - **Star : 39** 使用[Amazon Product Advertising API]的客户端库(https://program.amazon.com/gp/advertising /api/detail/main.html)。
* [anaconda](https://github.com/ChimeraCoder/anaconda) - **Star : 983**  Go 客户端库获取Twitter 1.1 API。
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) - **Star : 4854** Go编程语言的官方AWS SDK。
* [brewerydb](https://github.com/naegelejd/brewerydb) - **Star : 16** 访问BreweryDB API的库。
* [cachet](https://github.com/andygrunwald/cachet) - **Star : 65** 使用客户端库[Cachet(开源状态页系统)](https://cachethq.io/)。
* [circleci](https://github.com/jszwedko/go-circleci) - **Star : 40** 使用客户端库与CircleCI的API进行交互。
* [clarifai](https://github.com/samuelcouch/clarifai) - 使用客户端库与Clarifai API进行接口。
* [codeship-go](https://github.com/codeship/codeship-go) - **Star : 14**  Go 客户端库与Codeship的API v2交互。
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - **Star : 10**  Go 客户端库与Coinpaprika的API交互。
* [discordgo](https://github.com/bwmarrin/discordgo) - **Star : 924**  Go 绑定不和谐聊天API。
* [ethrpc](https://github.com/onrik/ethrpc) - **Star : 164**  Go 绑定Ethereum JSON RPC API。
* [facebook](https://github.com/huandu/facebook) - **Star : 764** 支持Facebook图形API的Go库。
* [fcm](https://github.com/maddevsio/fcm) - **Star : 33**  Go 库中获取Firebase云消息。
* [gads](https://github.com/emiddleton/gads) - **Star : 44** 谷歌Adwords非官方API。
* [gami](https://github.com/bit4bit/gami) - **Star : 26**  Go 星号管理器接口库。
* [gcm](https://github.com/Aorioli/gcm) - **Star : 30**  Go 谷歌云消息库。
* [geo-golang](https://github.com/codingsince1985/geo-golang) - **Star : 300**  Go 图书馆访问谷歌地图(https://developers.google.com/maps/documentation/geocoding/intro), (MapQuest) (http://open.mapquestapi.com/geocoding/), (Nominatim) (https://developer.mapquest.com/documentation/open/nominatim-search), (OpenCage) (http://geocoder.opencagedata.com/api.html), (Bing) (https://msdn.microsoft.com/en-us/library/ff701715.aspx), (Mapbox) (https://www.mapbox.com/developers/api/geocoding/),以及[OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding api。
* [github](https://github.com/google/go-github) - **Star : 4666** 访问GitHub REST API v3的库。
* [githubql](https://github.com/shurcooL/githubql) - 访问GitHub GraphQL API v4的库。
* [go-chronos](https://github.com/axelspringer/go-chronos) - **Star : 3** 用于与[Chronos](https://mesos.github.io/chronos/)作业调度程序进行交互的Go库
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) - **Star : 9** HackerNews API的小型Go客户端。
* [go-imgur](https://github.com/koffeinsource/go-imgur) - **Star : 12**  Go [imgur]的客户端库(https://imgur.com)
* [go-jira](https://github.com/andygrunwald/go-jira) - **Star : 548**  Go [Atlassian JIRA]的客户端库(https://www.atlassian.com/software/jira)
* [go-marathon](https://github.com/gambol99/go-marathon) - **Star : 188**  Go 图书馆与Mesosphere的马拉松PAAS互动。
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) - **Star : 16** 访问[MyAnimeList API]的客户机库(http://myanimelist.net/modules.php?go=api)。
* [go-sophos](https://github.com/esurdam/go-sophos) - **Star : 4** 为[Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/pdfs/documentation/utmonaws/sophos-ut-restful-api.pdf?
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) - **Star : 8**  Go 客户端库的SPTrans Olho Vivo API。
* [go-telegraph](https://gitlab.com/toby3d/telegraph) - 电报发布平台API客户端。
* [go-trending](https://github.com/andygrunwald/go-trending) - **Star : 99** 在Github上访问[trends repository](https://github.com/trends)和[developers](https://github.com/trending/developers)的库。
* [go-twitch](https://github.com/knspriggs/go-twitch) - **Star : 16**  Go 客户端与Twitch v3 API交互。
* [go-twitter](https://github.com/dghubble/go-twitter) - **Star : 687**  Go 客户端库获取Twitter v1.1 api。
* [go-unsplash](https://github.com/hbagdi/go-unsplash) - **Star : 20** 使用[Unsplash.com](https://unsplash.com) API的客户端库。
* [go-xkcd](https://github.com/nishanths/go-xkcd) - **Star : 37**  Go 客户端为xkcd API。
* [golyrics](https://github.com/mamal72/golyrics) - **Star : 29** Golyrics是一个从Wikia网站获取音乐歌词数据的Go库。
* [gomalshare](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](http://www.malshare.com/)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - **Star : 36**  Go MusicBrainz WS2客户端库。
* [google](https://github.com/google/google-api-go-client) - 为Go自动生成谷歌api。
* [google-analytics](https://github.com/chonthu/go-google-analytics) - **Star : 12** 简单的包装，方便的谷歌分析报告。
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - 谷歌云api Go 客户端库。
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - **Star : 6** 使用客户端库[谷歌G Suite Email Audit API](https://developer.google.com/admin-sdk/email-audit/)。
* [gostorm](https://github.com/jsgilmore/gostorm) - **Star : 118** GoStorm是一个Go库，它实现了编写Storm spout和bolt所需的通信协议，这些协议用于与Storm shell通信。
* [hipchat](https://github.com/andybons/hipchat) - **Star : 109** 这个项目为Hipchat API实现了一个golang客户端库。
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - **Star : 114** 一个用于通过XMPP与HipChat通信的golang包。
* [igdb](https://github.com/Henry-Sarabia/igdb) - **Star : 52** 使用client for [Internet Game Database API](https://api.igdb.com/)。
* [Medium](https://github.com/Medium/medium-sdk-go) - **Star : 114** Golang SDK for Medium的OAuth2 API。
* [megos](https://github.com/andygrunwald/megos) - **Star : 57** 用于访问[Apache Mesos](http://mesos.apache.org/)集群的客户机库。
* [minio-go](https://github.com/minio/minio-go) - **Star : 703** 用于Amazon S3兼容云存储的Minio Go库。
* [mixpanel](https://github.com/dukex/mixpanel) - **Star : 28** Mixpanel是一个库，用于跟踪事件并将Mixpanel概要文件更新从go应用程序发送到Mixpanel。
* [patreon-go](https://github.com/mxpv/patreon-go) - **Star : 16**  Go Patreon API库。
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) - **Star : 296** PayPal支付API的包装器。
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - **Star : 1** Playlyfe Rest API Go SDK。
* [pushover](https://github.com/gregdel/pushover) - **Star : 57**  Go 包装的Pushover API。
* [rrdaclient](https://github.com/Omie/rrdaclient) - **Star : 8**  Go 库访问statdns.com API，它是RRDA API。HTTP上的DNS查询。
* [shopify](https://github.com/rapito/go-shopify) - **Star : 20** 到库中对Shopify API发出CRUD请求。
* [simples3](https://github.com/rhnvrm/simples3) - **Star : 9** 使用REST和用Go编写的V4签名的AWS S3库非常简单。
* [slack](https://github.com/nlopes/slack) - **Star : 2324** Slack API in Go。
* [smite](https://github.com/sergiotapia/smitego) - **Star : 10** Go package封装对Smite游戏API的访问。
* [spotify](https://github.com/rapito/go-spotify) - **Star : 16**  Go 图书馆访问Spotify WEB API。
* [steam](https://github.com/sostronk/go-steam) - **Star : 14**  Go 库与Steam游戏服务器进行交互。
* [stripe](https://github.com/stripe/stripe-go) - **Star : 913**  Go 客户端获取Stripe API。
* [textbelt](https://github.com/dietsche/textbelt) -  Go 客户端textbelt.com txt消息传递API。
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) - **Star : 13** 简单的golang包与[themoviedb.org]通信(https://themoviedb.org)。
* [translate](https://github.com/poorny/translate) -  Go 网上翻译包。
* [Trello](https://github.com/adlio/trello) - **Star : 97**  Go 包装Trello API。
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) -  Go 包装TripAdvisor API。
* [tumblr](https://github.com/mattcunningham/gumblr) - **Star : 6**  Go 包装Tumblr v2 API。
* [uptimerobot](https://github.com/bitfield/uptimerobot) - **Star : 34** 运行时Robot v2 API的Go包装器和命令行客户机。
* [webhooks](https://github.com/go-playground/webhooks) - **Star : 334** 用于GitHub和Bitbucket的Webhook接收器。
* [wit-go](https://github.com/wit-ai/wit-go) - **Star : 44** 机智的客户。ai HTTP API。
* [ynab](https://github.com/brunomvsouza/ynab.go) - **Star : 10**  Go 包装YNAB API。
* [zooz](https://github.com/gojuno/go-zooz) - **Star : 5**  Go 客户端获取Zooz API。

## 公用事业公司

*一般实用工具和工具，使您的生活更容易。*

* [abutil](https://github.com/bahlo/abutil) - **Star : 52** 收集常用的戈朗助手。
* [apm](https://github.com/topfreegames/apm) - **Star : 126** 使用HTTP API的Golang应用程序进程管理器。
* [backscanner](https://github.com/icza/backscanner) - **Star : 8** 类似bufio的扫描仪。扫描器，但它以相反的顺序读取和返回行，从给定的位置开始，然后返回。
* [blank](https://github.com/Henry-Sarabia/blank) - **Star : 1** 验证或删除字符串中的空白和空白。
* [boilr](https://github.com/tmrts/boilr) - **Star : 918** 非常快的CLI工具，用于从样板模板创建项目。
* [chyle](https://github.com/antham/chyle) - **Star : 106** 使用具有多种配置可能性的git存储库生成变更日志。
* [circuit](https://github.com/cep21/circuit) - **Star : 313** 一个高效和功能齐全的Hystrix喜欢 Go 实现断路器模式。
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) - **Star : 774** 接通断路器。
* [clockwork](https://github.com/jonboulle/clockwork) - **Star : 213** 一个简单的假时钟给戈朗。
* [command](https://github.com/txgruppi/command) - **Star : 9** 使用线程安全的串行和并行调度程序的Go命令模式。
* [copy-pasta](https://github.com/jutkko/copy-pasta) - **Star : 37** 通用多工作站剪贴板，使用S3作为存储的后端。
* [ctop](https://github.com/bcicen/ctop) - **Star : 8604** [Top-like](http://ctop.sh)接口(例如htop)用于容器度量。
* [ctxutil](https://github.com/posener/ctxutil) - **Star : 6** 上下文的实用程序函数的集合。
* [dbt](https://github.com/nikogura/dbt) - **Star : 10** 用于从中心可信存储库运行自更新签名二进制文件的框架。
* [Death](https://github.com/vrecan/death) - **Star : 131** 使用信号管理go应用程序关闭。
* [Deepcopier](https://github.com/ulule/deepcopier) - **Star : 201** 简单的结构复制为Go。
* [delve](https://github.com/derekparker/delve) - 调试器。
* [dlog](https://github.com/kirillDanshin/dlog) - **Star : 15** 编译时控制的日志程序，使您的版本更小，而不删除调试调用。
* [ergo](https://github.com/cristianoliveira/ergo) - **Star : 306** 管理运行在不同端口上的多个本地服务变得很容易。
* [evaluator](https://github.com/nullne/evaluator) - **Star : 14** 基于s表达式动态计算表达式。它很简单，很容易扩展。
* [fastlz](https://github.com/digitalcrab/fastlz) - **Star : 11** 为GoLang总结[FastLz](http://fastlz.org/)(免费、开源、可移植的实时压缩库)。
* [filetype](https://github.com/h2non/filetype) - **Star : 922** 小程序包来推断文件类型，检查神奇数字签名。
* [filler](https://github.com/yaronsumel/filler) - **Star : 14** 使用“fill”标签填充结构的小工具。
* [filter](https://github.com/gookit/filter) - **Star : 10** 提供Go数据的过滤、清理和转换。
* [fzf](https://github.com/junegunn/fzf) - **Star : 22278** 用Go编写的命令行模糊查找器。
* [gaper](https://github.com/maxcnunes/gaper) - **Star : 37** 当Go项目崩溃或一些人看到文件更改时，构建并重新启动该项目。
* [generate](https://github.com/go-playground/generate) - **Star : 19** 运行go在指定的路径或环境变量上递归生成，并可以通过正则表达式进行筛选。
* [ghokin](https://github.com/antham/ghokin) - **Star : 12** 没有外部依赖的gherkin (cucumber, behat…)并行格式化程序。
* [git-time-metric](https://github.com/git-time-metric/gtm) - **Star : 708** Git的简单、无缝、轻量级时间跟踪。
* [go-astitodo](https://github.com/asticode/go-astitodo) - **Star : 45** 在GO代码中解析TODOs。
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - **Star : 159** go:生成用于包装golang插件导出的符号的工具(仅1.8)。
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) - **Star : 79** 纯Go bsdiff和bspatch库和CLI工具。
* [go-dry](https://github.com/ungerik/go-dry) - **Star : 432** 晾干(不要重复)打包带走。
* [go-funk](https://github.com/thoas/go-funk) - **Star : 945** 现代Go实用程序库，它提供了助手(map, find, contains, filter, chunk, reverse，…)。
* [go-health](https://github.com/Talento90/go-health) - **Star : 63** 健康包简化了向服务中添加健康检查的方式。
* [go-httpheader](https://github.com/mozillazg/go-httpheader) - **Star : 14**  Go 库中编码结构到头字段。
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) - **Star : 2** 打包处理问题细节。
* [go-rate](https://github.com/beefsack/go-rate) - **Star : 290** 围棋限速器。
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - **Star : 103** 用Go编写的XML站点地图生成器。
* [go-torch](https://github.com/uber/go-torch) - 用于Go程序的随机火焰图分析器。
* [go-trigger](https://github.com/sadlil/go-trigger) - **Star : 180** Go-lang全局事件触发器，使用id注册事件并从项目的任何位置触发事件。
* [goback](https://github.com/carlescere/goback) - **Star : 39**  Go 简单的指数回退包。
* [godaemon](https://github.com/VividCortex/godaemon) - **Star : 400** 编写守护进程的实用程序。
* [godropbox](https://github.com/dropbox/godropbox) - **Star : 3727** 用于从Dropbox编写Go服务/应用程序的公共库。
* [gohper](https://github.com/cosiner/gohper) - **Star : 248** 各种工具/模块有助于开发。
* [golarm](https://github.com/msempere/golarm) - **Star : 34** 带有系统事件的火警警报。
* [golog](https://github.com/mlimaloureiro/golog) - **Star : 43** 简单和轻量级的CLI工具，时间跟踪您的任务。
* [gopencils](https://github.com/bndr/gopencils) - **Star : 422** 小而简单的包，可以轻松地使用REST api。
* [goplaceholder](https://github.com/michiwend/goplaceholder) - **Star : 22** 一个用于生成占位符图像的小型golang库。
* [goreadability](https://github.com/philipjkim/goreadability) - **Star : 28** 网页摘要提取器使用Facebook开放图形和arc90的可读性。
* [goreleaser](https://github.com/goreleaser/goreleaser) - **Star : 4307** 交付Go二进制文件的速度越快越容易越好。
* [goreporter](https://github.com/wgliang/goreporter) - Golang工具，做静态分析，单元测试，代码审查和生成代码质量报告。
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - **Star : 24** 具有几乎所有功能的SeaweedFS客户端库。
* [gostrutils](https://github.com/ik5/gostrutils) - **Star : 14** 字符串操作和转换函数的集合。
* [gotenv](https://github.com/subosito/gotenv) - **Star : 136** 从'加载环境变量。env '或any ' io。读者在走。
* [gpath](https://github.com/tenntenn/gpath) - **Star : 25** 库，使用Go在反射中的表达式简化访问结构字段。
* [gubrak](https://github.com/novalagung/gubrak) - **Star : 125** 带有语法糖的Golang实用程序库。就像lodash，但对golang来说。
* [handy](https://github.com/miguelpragier/handy) - **Star : 43** 许多实用程序和帮助程序，如字符串处理程序/格式化程序和验证器。
* [htcat](https://github.com/htcat/htcat) - **Star : 480** 并行和流水线HTTP GET实用程序。
* [hub](https://github.com/github/hub) - **Star : 16500** 使用附加功能包装git命令，以便从终端与github交互。
* [hystrix-go](https://github.com/afex/hystrix-go) - **Star : 1928** 实现程序定义的回退(即断路器)的Hystrix模式。
* [immortal](https://github.com/immortal/immortal) - **Star : 597** \*nix跨平台(OS无关)管理器。
* [intrinsic](https://github.com/mengzhuo/intrinsic) - **Star : 39** 使用x86 SIMD无需编写任何汇编代码。
* [jump](https://github.com/gsamokovarov/jump) - **Star : 638** 通过学习你的习惯，跳跃可以帮助你更快地导航。
* [koazee](https://github.com/wesovilabs/koazee) - **Star : 273** 库的灵感来自于延迟计算和函数式编程，从而减少了使用数组的麻烦。
* [lrserver](https://github.com/jaschaephraim/lrserver) - **Star : 99** lieload服务器为Go。
* [mc](https://github.com/minio/mc) - **Star : 1054** Minio Client提供了与Amazon S3兼容的云存储和文件系统一起工作的最小工具。
* [mergo](https://github.com/imdario/mergo) - **Star : 816** 帮助合并结构和地图在戈朗。对于配置默认值很有用，避免了混乱的if语句。
* [mimemagic](https://github.com/zRedShift/mimemagic) - **Star : 43** 纯粹 Go 超性能MIME嗅探库/实用程序。
* [mimesniffer](https://github.com/aofei/mimesniffer) - **Star : 6** 一个用于Go的MIME类型嗅探器。
* [mimetype](https://github.com/gabriel-vasile/mimetype) - **Star : 97** 用于基于神奇数字的MIME类型检测的包。
* [minify](https://github.com/tdewolff/minify) - **Star : 1832** 用于HTML、CSS、JS、XML、JSON和SVG文件格式的快速缩小器。
* [minquery](https://github.com/icza/minquery) - **Star : 50** MongoDB /分别。支持高效分页的v2查询(用于继续列出我们停止的文档的游标)。
* [mmake](https://github.com/tj/mmake) - **Star : 1449** 现代。
* [moldova](https://github.com/StabbyCutyou/moldova) - **Star : 148** 用于基于输入模板生成随机数据的实用程序。
* [mole](https://github.com/davrodpin/mole) - **Star : 1287** cli应用程序可以轻松创建ssh隧道。
* [mssqlx](https://github.com/linxGnu/mssqlx) - **Star : 56** 数据库客户端库，代理任何主从，主从结构。轻量级和自动平衡的想法。
* [multitick](https://github.com/VividCortex/multitick) - **Star : 58** 多路复用器。
* [myhttp](https://github.com/inancgumus/myhttp) - **Star : 34** 简单的API，使HTTP GET请求与超时支持。
* [netbug](https://github.com/e-dard/netbug) - **Star : 65** 轻松远程分析您的服务。
* [okrun](https://github.com/xta/okrun) - **Star : 14**  Go 运行错误蒸汽压路机。
* [olaf](https://github.com/btnguyen2k/olaf) - **Star : 1** Twitter雪花在Go中实现。
* [onecache](https://github.com/adelowo/onecache) - **Star : 97** 支持多个后端存储(Redis、Memcached、文件系统等)的缓存库。
* [panicparse](https://github.com/maruel/panicparse) - **Star : 2071** 将类似的goroutine分组并为堆栈转储着色。
* [peco](https://github.com/peco/peco) - **Star : 5392** 简单的交互过滤工具。
* [pgo](https://github.com/arthurkushman/pgo) - **Star : 23** 方便的函数为PHP社区。
* [pm](https://github.com/VividCortex/pm) - **Star : 71** 进程(即goroutine)管理器与HTTP API。
* [profile](https://github.com/pkg/profile) - **Star : 965** Go的简单分析支持包。
* [rclient](https://github.com/zpatrick/rclient) - **Star : 26** 可读、灵活、易于使用的REST api客户机。
* [realize](https://github.com/tockins/realize) - 使用文件监视程序构建系统并实时重新加载。使用自定义路径运行、构建和监视文件更改。
* [repeat](https://github.com/ssgreg/repeat) - **Star : 56** 执行不同的后退策略，这对重新尝试操作和心跳非常有用。
* [request](https://github.com/mozillazg/request) - **Star : 353** HTTP请求 Go 人类™。
* [rerate](https://github.com/abo/rerate) - **Star : 12** 基于redis的速率计数器和Go的速率限制器。
* [rerun](https://github.com/ivpusic/rerun) - **Star : 153** 当源代码发生更改时，重新编译和重新运行go应用程序。
* [resty](https://github.com/go-resty/resty) - **Star : 1805** 简单的HTTP和REST客户端Go的灵感来自Ruby REST -client。
* [retry](https://github.com/kamilsk/retry) - **Star : 137** 最先进的重复执行动作的功能机制，直至成功。
* [retry](https://github.com/percolate/retry) - **Star : 2** 一个简单但高度可配置的Go重试包。
* [retry](https://github.com/thedevsaddam/retry) - **Star : 34** 简单易用的重试机制包，为 Go 。
* [retry](https://github.com/shafreeck/retry) - **Star : 9** 一个相当简单的库，以确保您的工作可以完成。
* [retry-go](https://github.com/rafaeljesus/retry-go) - **Star : 26** 对戈朗来说，重试变得简单而容易。
* [robustly](https://github.com/VividCortex/robustly) - **Star : 132** 弹性地运行函数，捕捉并重新启动恐慌。
* [scan](https://github.com/blockloop/scan) - **Star : 11** 扫描golang的sql。行直接指向结构、片或基本类型。
* [serve](https://github.com/syntaqx/serve) - **Star : 190** 任何您需要的静态http服务器。
* [silk](https://github.com/chrispassas/silk) - **Star : 4** 阅读silk netflow文件。
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) - **Star : 2** 基本类型之间的片转换。
* [slicer](https://github.com/leaanthony/slicer) - **Star : 3** 使处理切片更容易。
* [spinner](https://github.com/briandowns/spinner) - **Star : 763**  Go 软件包，以方便地提供一个终端旋转与选项。
* [sqlx](https://github.com/jmoiron/sqlx) - **Star : 6516** 在优秀的内置数据库/sql包之上提供一组扩展。
* [sslice](https://github.com/yaa110/sslice) - **Star : 2** 创建一个总是排序的切片。
* [Storm](https://github.com/asdine/storm) - **Star : 1329** 简单而强大的BoltDB工具包。
* [structs](https://github.com/PumpkinSeed/structs) - **Star : 12** 实现操作结构的简单函数。
* [Task](https://github.com/go-task/task) - **Star : 1855** 简单的“使”的选择。
* [toolbox](https://github.com/viant/toolbox) - **Star : 86** 切片，地图，多ap，结构，功能，数据转换实用工具。服务路由器，宏评估器，令牌器。
* [ugo](https://github.com/alxrm/ugo) - **Star : 20** ugo是一个切片工具箱，具有简洁的Go语法。
* [UNIS](https://github.com/esemplastic/unis) - **Star : 69** 常见的架构™字符串公用事业中 Go 。
* [usql](https://github.com/knq/usql) - usql是一个用于SQL数据库的通用命令行接口。
* [util](https://github.com/shomali11/util) - **Star : 130** 有用实用函数的集合。(字符串，并发，操作，…)
* [wuzz](https://github.com/asciimoo/wuzz) - **Star : 8177** 用于HTTP检查的交互式cli工具。
* [xferspdy](https://github.com/monmohan/xferspdy) - **Star : 68** Xferspdy在golang中提供二进制diff和补丁库。

## UUID

*用于处理uuid的库。*

* [goid](https://github.com/jakehl/goid) - **Star : 20** 生成和解析RFC4122兼容的V4 uuid。
* [sno](https://github.com/muyo/sno) - **Star : 12** 使用嵌入元数据的紧凑、可排序和快速的惟一id。
* [ulid](https://github.com/oklog/ulid) - **Star : 1645** 实现了ULID(普遍唯一的词典分类标识符)。
* [uuid](https://github.com/agext/uuid) - **Star : 10** 使用快速或加密质量的随机节点标识符生成、编码和解码UUIDs v1。
* [uuid](https://github.com/gofrs/uuid) - **Star : 543** 通用唯一标识符(UUID)的实现。支持uuid的创建和解析。积极维护satori uuid的fork。
* [wuid](https://github.com/edwingeng/wuid) - **Star : 265** 一个非常快的唯一数字生成器，比UUID快10-135倍。

## 验证

*库进行验证。*

* [checkdigit](https://github.com/osamingo/checkdigit) - **Star : 43** 提供校验数字算法(Luhn, Verhoeff, Damm)和计算器(ISBN, EAN, JAN, UPC等)。
* [govalidator](https://github.com/asaskevich/govalidator) - **Star : 3456** 验证器和消毒剂的字符串，数字，切片和结构。
* [govalidator](https://github.com/thedevsaddam/govalidator) - **Star : 678** 用简单的规则验证Golang请求数据。深受Laravel请求验证的启发。
* [jio](https://github.com/faceair/jio) - **Star : 20** jio是一个json模式验证器，类似于[joi](https://github.com/hapijs/joi)。
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - **Star : 998** 支持各种数据类型(结构、字符串、映射、片等)的验证，使用可配置和可扩展的验证规则，这些规则在通常的代码构造中指定，而不是在结构标签中指定。
* [validate](https://github.com/gookit/validate) - **Star : 81**  Go 封装数据验证和过滤。支持验证映射、结构、请求(表单、JSON、url)。值，上载文件)数据和更多特性。
* [validate](https://github.com/gobuffalo/validate) - **Star : 19** 这个包提供了一个框架，用于为Go应用程序编写验证。
* [validator](https://github.com/go-playground/validator) - **Star : 3263**  Go 结构和字段验证，包括交叉字段，交叉结构，地图，切片和数组潜水。

## 版本控制

*用于版本控制的库。*

* [gh](https://github.com/rjeczalik/gh) - **Star : 68** 用于GitHub webhook的可编写脚本的服务器和net/http中间件。
* [git2go](https://github.com/libgit2/git2go) - **Star : 1336**  Go 绑定libgit2。
* [go-git](https://github.com/src-d/go-git) - **Star : 4084** 纯Go中高度可扩展的Git实现。
* [go-vcs](https://github.com/sourcegraph/go-vcs) - **Star : 69** 在Go中操作和检查VCS存储库。
* [hercules](https://github.com/src-d/hercules) - **Star : 491** 从Git存储库历史中获得高级见解。
* [hgo](https://github.com/beyang/hgo) - **Star : 12** Hgo是一组Go包的集合，提供对本地Mercurial存储库的读取访问。

## 视频

*用于操作视频的库。*

* [gmf](https://github.com/3d0c/gmf) - **Star : 506**  Go 绑定为FFmpeg av\*库。
* [go-astisub](https://github.com/asticode/go-astisub) - **Star : 161** 在GO(。srt， .stl， .ttml， .webvtt， .ssa/。ass, teletext， .smi等)。
* [go-astits](https://github.com/asticode/go-astits) - **Star : 256** 在GO中解析和演示MPEG传输流(.ts)。
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) - **Star : 36** 苹果m3u8播放列表的解析器和生成器库。
* [goav](https://github.com/giorgisio/goav) - **Star : 743** FFmpeg的Comphrensive Go绑定。
* [gst](https://github.com/ziutek/gst) - **Star : 153**  Go 绑定GStreamer。
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) - **Star : 11** 字幕格式支持 Go 。支持.srt、.ttml和.ass。
* [libvlc-go](https://github.com/adrg/libvlc-go) - **Star : 58** Go绑定libvlc 2.X/3.X/4。X(由VLC媒体播放器使用)。
* [v4l](https://github.com/korandiz/v4l) - **Star : 26** 用于Linux的视频捕捉库，用Go编写。

## Web框架

*完整的堆栈web框架。*

* [aah](https://aahframework.org) - 可伸缩、高性能、快速开发的Go Web框架。
* [Aero](https://github.com/aerogo/aero) - **Star : 152** 高性能的Go web框架，在Lighthouse中达到最高分。
* [Air](https://github.com/aofei/air) - **Star : 511** 一个理想的精细化的Go web框架。
* [Banjo](https://github.com/nsheremet/banjo) - **Star : 7** 非常简单和快速的网络框架 Go 。
* [Beego](https://github.com/astaxie/beego) - **Star : 20950** beego是一种用于围棋编程语言的开源高性能web框架。
* [Buffalo](http://gobuffalo.io) - 提高Rails的生产力!
* [Echo](https://github.com/labstack/echo) - **Star : 14277** 高性能、极简的Go web框架。
* [Fireball](https://github.com/zpatrick/fireball) - **Star : 48** 更“自然”的感觉web框架。
* [Gem](https://github.com/go-gem/gem) - **Star : 153** 简单快速的web框架，对REST API友好。
* [Gin](https://github.com/gin-gonic/gin) - **Star : 28531** Gin是一个用Go编写的web框架!它具有一个类似于martini的API，性能更好，速度快40倍。如果您需要性能和良好的生产力。
* [Gizmo](https://github.com/NYTimes/gizmo) - **Star : 2809** 《纽约时报》使用的微服务工具包。
* [go-json-rest](https://github.com/ant0ine/go-json-rest) - **Star : 3316** 设置RESTful JSON API的快速简便方法。
* [go-rest](https://github.com/ungerik/go-rest) - **Star : 115** 小而恶的休息框架为 Go 。
* [Goa](https://github.com/goadesign/goa) - **Star : 3449** Goa为在Go中开发远程api和微服务提供了一种全面的方法。
* [Golax](https://github.com/fulldump/golax) - **Star : 71** 一个非Sinatra快速HTTP框架，支持谷歌自定义方法、深度拦截器、递归等。
* [Golf](https://github.com/dinever/golf) - **Star : 235** 高尔夫是一个快速、简单、轻量级的围棋微web框架。它具有强大的功能，除了Go标准库之外没有其他依赖项。
* [Gondola](https://github.com/rainycape/gondola) - **Star : 315** web框架写的网站越快越好。
* [gongular](https://github.com/mustafaakin/gongular) - **Star : 415** 带有输入映射/验证和(DI)依赖注入的快速web框架。
* [hiboot](https://github.com/hidevopsio/hiboot) - **Star : 80** hiboot是一个高性能的web应用程序框架，支持自动配置和依赖注入。
* [Macaron](https://github.com/go-macaron/macaron) - **Star : 2780** Macaron是一个高生产力和模块化设计的web框架在Go。
* [mango](https://github.com/paulbellamy/mango) - **Star : 339** 芒果是一个模块化的Go web应用程序框架，灵感来自于Rack和PEP333。
* [Microservice](https://github.com/claygod/microservice) - **Star : 56** 创建微服务的框架，用Golang编写。
* [neo](https://github.com/ivpusic/neo) - **Star : 392** Neo是一个非常简单且快速的Web框架API。
* [nio](https://github.com/go-nio/nio) - **Star : 21** 现代的、最小的和高效的Go HTTP框架。
* [Resoursea](https://github.com/resoursea/api) - **Star : 29** 用于快速编写基于资源的服务的REST框架。
* [REST Layer](http://rest-layer.io) - 框架，用于在数据库之上构建REST/GraphQL API，主要是通过代码进行配置。
* [Revel](https://github.com/revel/revel) - **Star : 11154** 用于Go语言的高效web框架。
* [rex](https://github.com/goanywhere/rex) - **Star : 25** 雷克斯是一个模块化的发展图书馆建立在大猩猩/ mux与'net/http`完全兼容。
* [sawsij](https://github.com/jaybill/sawsij) - **Star : 2** 轻量级、开源的web框架，用于构建高性能、数据驱动的web应用程序。
* [tango](https://github.com/lunny/tango) - **Star : 813** 微和可插入的网络框架 Go 。
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - **Star : 998** 受Dropwizard启发，构建JSON web服务的Go框架。
* [traffic](https://github.com/pilu/traffic) - Sinatra启发了regexp/pattern mux和用于Go的web框架。
* [uAdmin](https://github.com/uadmin/uadmin) - **Star : 46** 完全功能的web框架为Golang，灵感来自Django。
* [utron](https://github.com/gernest/utron) - **Star : 2136** Go(Golang)的轻量级MVC框架。
* [vox](https://github.com/aisk/vox) - **Star : 27** 一个面向人类的golang web框架，深受Koa的启发。
* [WebGo](https://github.com/bnkamalesh/webgo) - **Star : 70** 构建web应用程序的微框架;处理程序链接、中间件和上下文注入。与标准库兼容的HTTP处理程序(即http.HandlerFunc)。
* [YARF](https://github.com/yarf-framework/yarf) - **Star : 49** 快速微框架，旨在以快速和简单的方式构建REST api和web服务。
* [Zerver](https://github.com/cosiner/zerver) - Zerver是一个表现力强、模块化、功能完备的RESTful框架。

### 中间件

#### 实际仿真中间件

* [client-timing](https://github.com/posener/client-timing) - **Star : 11** 用于服务器定时报头的HTTP客户机。
* [CORS](https://github.com/rs/cors) - **Star : 1160** 轻松地向API添加CORS功能。
* [formjson](https://github.com/rs/formjson) - **Star : 33** 透明地将JSON输入作为标准表单POST处理。
* [go-server-timing](https://github.com/mitchellh/go-server-timing) - **Star : 743** 添加/解析Server-Timing头。
* [Limiter](https://github.com/ulule/limiter) - **Star : 758** 死简单的速度限制中间件 Go 。
* [ln-paywall](https://github.com/philippgille/ln-paywall) - **Star : 86** 使用Lightning Network(比特币)实现基于每个请求的api货币化中间件。
* [Tollbooth](https://github.com/didip/tollbooth) - **Star : 1202** 速率限制HTTP请求处理程序。
* [XFF](https://github.com/sebest/xff) - **Star : 71** 处理“x - forwarding - for”头和好友。

#### 用于创建HTTP中间件的库

* [alice](https://github.com/justinas/alice) - **Star : 1801** Go的无痛中间件链接。
* [catena](https://github.com/codemodus/catena) - **Star : 7** http。处理程序包装器连接(与“chain”相同的API)。
* [chain](https://github.com/codemodus/chain) - **Star : 63** 带有范围数据的处理程序包装器链接(基于网络/上下文的“中间件”)。
* [go-wrap](https://github.com/go-on/wrap) - **Star : 56** net/http的小型中间件包。
* [gores](https://github.com/alioygur/gores) - **Star : 82** 处理HTML、JSON、XML等响应的Go包。对于RESTful api非常有用。
* [interpose](https://github.com/carbocation/interpose) - **Star : 290** golang的极简网络/http中间件。
* [muxchain](https://github.com/stephens2424/muxchain) - **Star : 207** 用于net/http的轻量级中间件。
* [negroni](https://github.com/urfave/negroni) - **Star : 6267** Golang的惯用HTTP中间件。
* [render](https://github.com/unrolled/render) - **Star : 1259** Go package用于方便地呈现JSON、XML和HTML模板响应。
* [renderer](https://github.com/thedevsaddam/renderer) - **Star : 165** 简单、轻量级和更快的响应(JSON、JSONP、XML、YAML、HTML、文件)。
* [rye](https://github.com/InVisionApp/rye) - **Star : 92** 支持JWT、CORS、Statsd和Go 1.7上下文的小型Go中间件库(带有罐装中间件)。
* [stats](https://github.com/thoas/stats) - **Star : 536** 使用中间件来存储关于web应用程序的各种信息。

### 路由器

* [alien](https://github.com/gernest/alien) - **Star : 106** 轻量级和快速http路由器从外层空间。
* [bellt](https://github.com/GuilhermeCaruso/bellt) - **Star : 37** 一个简单的Go HTTP路由器。
* [Bone](https://github.com/go-zoo/bone) - **Star : 1219** 闪电快速HTTP多路复用器。
* [Bxog](https://github.com/claygod/Bxog) - **Star : 93** 简单和快速的HTTP路由器 Go 。它可以处理不同难度、长度和嵌套的路径。他还知道如何根据接收到的参数创建URL。
* [chi](https://github.com/go-chi/chi) - **Star : 5813** 小型，快速和表达的HTTP路由器建立在网络/上下文。
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) - **Star : 732** 高性能路由器分叉从' httprouter '。第一个路由器适合“fasthttp”。
* [FastRouter](https://github.com/razonyang/fastrouter) - **Star : 18** 一个快速，灵活的HTTP路由器写在Go。
* [gocraft/web](https://github.com/gocraft/web) - **Star : 1388** Mux和中间件包在Go中。
* [Goji](https://github.com/goji/goji) - **Star : 759** 枸杞是一种简约的和灵活的与支持'net/context` HTTP请求多路复用器。
* [GoRouter](https://github.com/vardius/gorouter) - **Star : 46** GoRouter是一个服务器/ API微framwework HTTP请求路由器,多路复用器,路由器与中间件支持'net/context` mux提供请求。
* [gowww/router](https://github.com/gowww/router) - **Star : 158** 闪电快速HTTP路由器完全兼容网络/ HTTP。处理程序接口。
* [httprouter](https://github.com/julienschmidt/httprouter) - **Star : 9459** 高性能路由器。使用这个和标准http处理程序来形成一个非常高性能的web框架。
* [httptreemux](https://github.com/dimfeld/httptreemux) - **Star : 382** 高速，灵活的基于树的HTTP路由器 Go 。从httprouter灵感。
* [lars](https://github.com/go-playground/lars) - **Star : 375** 是一个轻量级、快速和可扩展的zero allocation HTTP路由器，用于创建可定制框架。
* [mux](https://github.com/gorilla/mux) - **Star : 9119** 强大的URL路由器和调度器为golang。
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - **Star : 357** 一个非常快的Go (golang) HTTP路由器，支持正则表达式路由匹配。完全支持构建RESTful api。
* [pure](https://github.com/go-playground/pure) - **Star : 83** 是一个轻量级HTTP路由器，它坚持net/ HTTP“实现”的std。
* [Siesta](https://github.com/VividCortex/siesta) - **Star : 349** 编写中间件和处理程序的可组合框架。
* [vestigo](https://github.com/husobee/vestigo) - **Star : 250** 高性能，独立，HTTP兼容的URL路由器的go web应用程序。
* [violetear](https://github.com/nbari/violetear) - **Star : 94** HTTP路由器。
* [xmux](https://github.com/rs/xmux) - **Star : 87** 高性能mux基于httprouter 'net/context`支持。
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - **Star : 439** 一个简单和快速的HTTP路由器 Go 。

## 窗户

* [d3d9](https://github.com/gonutz/d3d9) - **Star : 86**  Go 绑定Direct3D9。
* [go-ole](https://github.com/go-ole/go-ole) - **Star : 539** 用于golang的Win32 OLE实现。
* [gosddl](https://github.com/MonaxGT/gosddl) - **Star : 1** 从SDDL-string到用户友好的JSON的转换器。SDDL由四个部分组成:所有者、主群、DACL、SACL。

## XML

*用于操作XML的库和工具。*

* [XML-Comp](https://github.com/xml-comp/xml-comp) - **Star : 15** 简单的命令行XML比较器，生成文件夹、文件和标记的差异。
* [xml2map](https://github.com/sbabiv/xml2map) - **Star : 15** XML来映射转换器编写的Golang。
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - **Star : 6** 基于libxml2的xmlwriter模块的过程性XML生成API。
* [xpath](https://github.com/antchfx/xpath) - **Star : 145** Go的XPath包。
* [xquery](https://github.com/antchfx/xquery) - **Star : 145** XQuery允许您使用XPath表达式从HTML/XML文档中提取数据。
* [zek](https://github.com/miku/zek) - **Star : 239** 从XML生成Go结构。

# 工具

* Go 软件和插件。*

## 代码分析

* [apicompat](https://github.com/bradleyfalzon/apicompat) - **Star : 164** 检查Go项目最近的更改，以获得向后不兼容的更改。
* [dupl](https://github.com/mibk/dupl) - **Star : 166** 用于代码克隆检测的工具。
* [errcheck](https://github.com/kisielk/errcheck) - **Star : 1306** Errcheck是一个用于检查Go程序中未检查错误的程序。
* [gcvis](https://github.com/davecheney/gcvis) - **Star : 908** 实时可视化Go程序GC跟踪数据。
* [go-checkstyle](https://github.com/qiniu/checkstyle) - **Star : 95** checkstyle是一个类似于java checkstyle的样式检查工具。这个工具的灵感来自java checkstyle, golint。该样式引用了Go Code Review注释中的一些要点。
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - **Star : 275** Go -cleanarch的创建是为了验证Clean体系结构规则，比如Go项目中的依赖规则和包之间的交互。
* [go-critic](https://github.com/go-critic/go-critic) - **Star : 548** 源代码linter带来的检查，目前没有实现在其他l。
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) - **Star : 171** 找到Go项目过时依赖项的简单方法。
* [go-outdated](https://github.com/firstrow/go-outdated) - **Star : 45** 显示过期包的控制台应用程序。
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - **Star : 365** 基于Web的Golang AST可视化器。
* [GoCover.io](http://gocover.io/) - GoCover。io提供任何golang包的代码覆盖率。
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - 工具来修复(添加，删除)您的Go自动导入。
* [GolangCI](https://golangci.com/) - GolangCI是一个针对GitHub pull请求的自动Golang代码审查服务。服务是开源的，对开源项目是免费的。
* [GoLint](https://github.com/golang/lint) - **Star : 3086** Golint是Go源代码的linter。
* [Golint online](http://go-lint.appspot.com/) - Lints在线 Go 源文件GitHub, Bitbucket和谷歌项目托管使用golint包。
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - 添加零值返回语句以匹配func返回类型。
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - gosimple是Go源代码的linter，专门用于简化代码。
* [gostatus](https://github.com/shurcooL/gostatus) - **Star : 240** 命令行工具，显示包含Go包的存储库的状态。
* [lint](https://github.com/surullabs/lint) - **Star : 62** 作为go测试的一部分运行l。
* [php-parser](https://github.com/z7zmey/php-parser) - **Star : 614** 用Go编写的PHP解析器。
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck是类固醇上的“go vet”，它应用了大量静态分析检查，您可能已经从c#的ReSharper等工具中习惯了这些检查。
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) - tarp在Go源代码中寻找没有直接单元测试的函数和方法。
* [unconvert](https://github.com/mdempsky/unconvert) - **Star : 257** 从Go源代码中删除不必要的类型转换。
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - 未使用的检查用于未使用的常量、变量、函数和类型的代码。
* [validate](https://github.com/mccoyst/validate) - **Star : 62** 使用标记自动验证结构字段。

## 编辑器插件

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) -  Go JetBrains ide插件。
* [go-language-server](https://github.com/theia-ide/go-language-server) - **Star : 28** 将VSCode go扩展转换为支持语言-服务器-协议的语言服务器的包装器。
* [go-mode](https://github.com/dominikh/go-mode.el) - **Star : 941** GNU/Emacs的Go模式。
* [go-plus](https://github.com/joefitzgerald/go-plus) - **Star : 1479** Go (Golang)包为Atom添加自动完成，格式化，语法检查，Linting和审查。
* [gocode](https://github.com/nsf/gocode) - **Star : 4710** Go编程语言的自动完成守护进程。
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - 这个扩展将对Go语言的基准分析支持添加到VS代码中。
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - **Star : 3211** Golang插件集合为文本编辑器SublimeText 3提供代码完成和其他类似idea的功能。
* [gounit-vim](https://github.com/hexdigest/gounit-vim) - **Star : 17** 用于基于函数或方法的签名生成Go测试的Vim插件。
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) - **Star : 12**  Go 语言支持Theia IDE。
* [velour](https://github.com/velour/velour) - **Star : 16** acme编辑器的IRC客户机。
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - **Star : 80** Vim插件，在保存时突出显示语法错误。
* [vim-go](https://github.com/fatih/vim-go) - **Star : 10631** 为Vim开发插件。
* [vscode-go](https://github.com/Microsoft/vscode-go) - **Star : 4988** Visual Studio代码的扩展(VS代码)，它提供了对Go语言的支持。
* [Watch](https://github.com/eaburns/Watch) - **Star : 165** 在acme win文件更改中运行命令。

## Go 生成工具

* [generic](https://github.com/usk81/generic) - **Star : 28** 灵活的Go数据类型。
* [genny](https://github.com/cheekybits/genny) - **Star : 918** 优雅的Go泛型。
* [gocontracts](https://github.com/Parquery/gocontracts) - **Star : 52** 通过同步代码和文档来实现契约式设计。
* [gonerics](http://github.com/bouk/gonerics) - Go中的惯用泛型。
* [gotests](https://github.com/cweill/gotests) - **Star : 2111** 从源代码生成Go测试。
* [gounit](https://github.com/hexdigest/gounit) - **Star : 28** 使用您自己的模板生成Go测试。
* [hasgo](https://github.com/DylanMeeus/hasgo) - **Star : 11** 为切片生成受Haskell启发的函数。
* [re2dfa](https://github.com/opennota/re2dfa) - **Star : 168** 将正则表达式转换为有限状态机并输出Go源代码。
* [TOML-to-Go](https://xuri.me/toml-to-go) - 立即在浏览器中将TOML转换为Go类型。

## Go 工具

* [colorgo](https://github.com/songgao/colorgo) - **Star : 95** 将“go”命令包装成彩色的“go build”输出。
* [depth](https://github.com/KyleBanks/depth) - **Star : 365** 通过分析导入，可视化任何包的依赖关系树。
* [gb](https://getgb.io/) - 一个易于使用的基于项目的构建工具的围棋编程语言。
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) - **Star : 13** 一个[Yeoman](http://yeoman.io)生成器，用于启动新的Go项目。
* [gilbert](https://go-gilbert.github.io) - 为Go项目构建系统和任务运行器。
* [go-callvis](https://github.com/TrueFurby/go-callvis) - **Star : 1910** 使用点格式可视化Go程序的调用图。
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) - **Star : 37** Bash完成go和wgo。
* [go-swagger](https://github.com/go-swagger/go-swagger) - **Star : 3783** Swagger 2.0实现的go。Swagger是RESTful API的一个简单而强大的表示。
* [godbg](https://github.com/tylerwince/godbg) - **Star : 157** 实现生锈的dbg!宏，以便在开发过程中快速、容易地调试。
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - 使用GitHub的OctoLinker浏览器扩展有效地浏览go文件。
* [richgo](https://github.com/kyoh86/richgo) - **Star : 372** 用文本装饰丰富“go test”输出。
* [rts](https://github.com/galeone/rts) - **Star : 183** RTS:对struct的响应。从服务器响应生成Go结构。

## 软件包

*用Go编写的软件。*

### DevOps的工具

* [aptly](https://github.com/smira/aptly) - 是一个Debian存储库管理工具。
* [aurora](https://github.com/xuri/aurora) - **Star : 386** 基于web的跨平台Beanstalkd队列服务器控制台。
* [awsenv](https://github.com/soniah/awsenv) - **Star : 20** 为概要文件加载Amazon (AWS)环境变量的小二进制文件。
* [Blast](https://github.com/dave/blast) - **Star : 168** 一个用于API负载测试和批处理作业的简单工具。
* [bombardier](https://github.com/codesenberg/bombardier) - **Star : 1669** 快速跨平台HTTP基准测试工具。
* [bosun](https://github.com/bosun-monitor/bosun) - **Star : 2841** 时间序列报警框架。
* [DepCharge](https://github.com/centerorbit/depcharge) - **Star : 9** 帮助在大型项目中的许多依赖项之间编排命令的执行。
* [dogo](https://github.com/liudng/dogo) - **Star : 215** 监视源文件中的更改并自动编译和运行(重新启动)。
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) - **Star : 22** 使用二进制、docker或无人机CI触发下游Jenkins作业。
* [drone-scp](https://github.com/appleboy/drone-scp) - **Star : 54** 使用二进制、docker或从属CI通过SSH复制文件和工件。
* [Dropship](https://github.com/chrismckenzie/dropship) - **Star : 46** 用于通过cdn部署代码的工具。
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - **Star : 93** Golang包，方便远程执行通过SSH和SCP下载通过' ProxyCommand '。
* [fac](https://github.com/mkchoi212/fac) - **Star : 1581** 命令行用户界面修复git合并冲突。
* [gaia](https://github.com/gaia-pipeline/gaia) - **Star : 3681** 用任何编程语言构建强大的管道。
* [Gitea](https://github.com/go-gitea/gitea) - **Star : 14517** 叉的Gogs，完全由社区驱动。
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - 将所有GitHub存储库、问题、里程碑和标签迁移到Gitea实例。
* [go-furnace](https://github.com/go-furnace/go-furnace) - **Star : 62** 用Go编写的托管解决方案。轻松地在AWS、GCP或DigitalOcean上部署应用程序。
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - **Star : 654** 启用Go应用程序进行自我更新。
* [gobrew](https://github.com/cryptojuice/gobrew) - **Star : 175** gobrew允许您轻松地在go的多个版本之间切换。
* [godbg](https://github.com/sirnewton01/godbg) - **Star : 220** 基于web的gdb前端应用程序。
* [Gogs](https://gogs.io/) - 一个在Go编程语言中自托管的Git服务。
* [gonative](https://github.com/inconshreveable/gonative) - **Star : 311** 该工具创建了一个Go构建，可以跨编译到所有平台，同时仍然使用启用了cgi的stdlib包版本。
* [govvv](https://github.com/ahmetalpbalkan/govvv) - “ Go 构建“包装器轻松地添加版本信息到二进制文件。
* [gox](https://github.com/mitchellh/gox) - **Star : 3316** 非常简单，没有多余的交叉编译工具。
* [goxc](https://github.com/laher/goxc) - **Star : 1627** 为Go构建工具，重点是交叉编译和打包。
* [grapes](https://github.com/yaronsumel/grapes) - **Star : 133** 轻量级工具，旨在轻松地通过ssh分发命令。
* [GVM](https://github.com/moovweb/gvm) - **Star : 4394** GVM提供了一个接口来管理Go版本。
* [Hey](https://github.com/rakyll/hey) - **Star : 6019** Hey是一个向web应用程序发送一些负载的小程序。
* [kala](https://github.com/ajvb/kala) - **Star : 1345** 简单、现代和高性能的作业调度程序。
* [kcli](https://github.com/cswank/kcli) - **Star : 65** 命令行工具，用于检查kafka主题/分区/消息。
* [kubernetes](https://github.com/kubernetes/kubernetes) - **Star : 54476** 来自谷歌的容器集群管理器。
* [lstags](https://github.com/ivanilves/lstags) - **Star : 218** 工具和API跨不同注册中心同步Docker图像。
* [lwc](https://github.com/timdp/lwc) - **Star : 8** UNIX wc命令的实时更新版本。
* [manssh](https://github.com/xwjdsh/manssh) - **Star : 202** manssh是一个命令行工具，可以方便地管理ssh别名配置。
* [Moby](https://github.com/moby/moby) - **Star : 53868** 为容器生态系统组装基于容器的系统的协作项目。
* [Mora](https://github.com/emicklei/mora) - **Star : 263** 用于访问MongoDB文档和元数据的REST服务器。
* [ostent](https://github.com/ostrost/ostent) - **Star : 164** 收集和显示系统指标，并可选地中继到石墨和/或fluxdb。
* [Packer](https://github.com/mitchellh/packer) - 封隔器是一种工具，用于从一个源配置为多个平台创建相同的机器图像。
* [Pewpew](https://github.com/bengadbois/pewpew) - **Star : 197** 灵活的HTTP命令行压力测试器。
* [Pomerium](https://github.com/pomerium/pomerium) - **Star : 462** Pomerium是一个可识别身份的访问代理。
* [Rodent](https://github.com/alouche/rodent) - **Star : 30** 啮齿动物帮助您管理Go版本、项目和跟踪依赖项。
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - **Star : 986** 小型实用程序/库，针对大型对象在Amazon S3中的高速传输进行了优化。
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - **Star : 531** 从命令行管理裸金属服务器(与使用Docker一样容易)。
* [script](https://github.com/bitfield/script) - **Star : 599** 使编写用于DevOps和系统管理任务的类shell脚本变得容易。
* [sg](https://github.com/ChristopherRabotin/sg) - **Star : 5** 基准测试一组HTTP端点(如ab)，可以在每个调用之间使用响应代码和数据，根据之前的响应来确定特定的服务器压力。
* [skm](https://github.com/TimothyYe/skm) - **Star : 544** SKM是一个简单而强大的SSH密钥管理器，它可以帮助您轻松地管理多个SSH密钥!
* [StatusOK](https://github.com/sanathp/statusok) - **Star : 1134** 监视您的网站和REST api。当服务器宕机或响应时间超过预期时，通过Slack、电子邮件获得通知。
* [traefik](https://github.com/containous/traefik) - **Star : 23052** 反向代理和负载均衡器，支持多个后端。
* [Vegeta](https://github.com/tsenart/vegeta) - **Star : 11831** HTTP负载测试工具和库。超过9000 !
* [webhook](https://github.com/adnanh/webhook) - **Star : 3965** 允许用户创建在服务器上执行命令的HTTP端点(钩子)的工具。
* [Wide](https://wide.b3log.org/login) - 为使用Golang的团队提供基于web的IDE。
* [winrm-cli](https://github.com/masterzen/winrm-cli) - **Star : 64** 在Windows机器上远程执行命令的Cli工具。

### 其他软件
* [borg](https://github.com/crufter/borg) - 基于终端的bash代码段搜索引擎。
* [boxed](https://github.com/tejo/boxed) - **Star : 72** 基于Dropbox的博客引擎。
* [Cherry](https://github.com/rafael-santiago/cherry) - **Star : 192** 微型网络聊天服务器在围棋。
* [Circuit](https://github.com/gocircuit/circuit) - **Star : 1774** 电路是一个可编程平台即服务(PaaS)和/或基础设施即服务(IaaS)，用于管理、发现、同步和编排包含云应用程序的服务和主机。
* [Comcast](https://github.com/tylertreat/Comcast) - **Star : 6120** 模拟坏的网络连接。
* [confd](https://github.com/kelseyhightower/confd) - **Star : 6298** 使用etcd或领事中的模板和数据管理本地应用程序配置文件。
* [DDNS](https://github.com/skibish/ddns) - **Star : 97** 个人DDNS客户端与数字海洋网络DNS作为后端。
* [Docker](http://www.docker.com/) - 面向开发人员和系统管理员的分布式应用程序的开放平台。
* [Documize](https://github.com/documize/community) - **Star : 781** 集成SaaS工具数据的现代wiki软件。
* [drive](https://github.com/odeke-em/drive) - **Star : 4906** 命令行的谷歌驱动器客户端。
* [Duplicacy](https://github.com/gilbertchen/duplicacy) - **Star : 2657** 基于无锁重复数据删除思想的跨平台网络和云备份工具。
* [gfile](https://github.com/Antonito/gfile) - **Star : 485** 通过WebRTC在两台计算机之间安全地传输文件，不需要任何第三方。
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - **Star : 879** 应用程序，显示更新的Go包在您的GOPATH。
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - **Star : 373** 视频流torrent客户端。
* [GoBoy](https://github.com/Humpheh/goboy) - **Star : 2088** 任天堂Game Boy彩色模拟器编写在围棋。
* [gocc](https://github.com/goccmack/gocc) - **Star : 335** Gocc是一个用Go编写的编译器工具包。
* [GoDNS](https://github.com/timothyye/godns) - **Star : 414** 一个动态DNS客户端工具，支持DNSPod & HE.net，用Go编写。
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - **Star : 12** Chrome扩展的Go Doc网站，其中显示的功能描述作为工具提示在功能列表。
* [GoLand](https://jetbrains.com/go) - 功能齐全的跨平台Go IDE。
* [Gor](https://github.com/buger/gor) - Http流量复制工具，用于实时回放从生产环境到阶段/开发环境的流量。
* [hugo](http://gohugo.io/) - 快速和现代静态网站引擎。
* [ide](https://github.com/thestrukture/ide) - **Star : 249** 浏览器访问的IDE。为Go with Go而设计。
* [ipe](https://github.com/dimiro1/ipe) - **Star : 272** 与用GO编写的Pusher客户机库兼容的开源Pusher服务器实现。
* [joincap](https://github.com/assafmo/joincap) - **Star : 121** 用于合并多个pcap文件的命令行实用程序。
* [Juju](https://jujucharms.com/) - 与云无关的服务部署和编制——支持EC2、Azure、Openstack、MAAS等。
* [Leaps](https://github.com/jeffail/leaps) - **Star : 640** 使用操作转换的成对编程服务。
* [lgo](https://github.com/yunabe/lgo) - **Star : 1761** 与木星互动围棋编程。它支持代码完成、代码检查和100% Go兼容性。
* [limetext](http://limetext.org/) - Lime Text是一个强大而优雅的文本编辑器，最初是在Go中开发的，它的目标是成为卓越文本的免费和开源软件的继承者。
* [LiteIDE](https://github.com/visualfc/liteide) - **Star : 5400** LiteIDE是一个简单的、开源的、跨平台的Go IDE。
* [mockingjay](https://github.com/quii/mockingjay-server) - **Star : 406** 从一个配置文件中伪造HTTP服务器和消费者驱动的契约。您还可以使服务器随机失常，以帮助进行更实际的性能测试。
* [myLG](https://github.com/mehrdadrad/mylg) - **Star : 2175** 用Go编写的命令行网络诊断工具。
* [naclpipe](https://github.com/unix4fun/naclpipe) - **Star : 20** 简单的NaCL EC25519基于加密管工具编写的Go。
* [nes](https://github.com/fogleman/nes) - **Star : 4095** 任天堂娱乐系统(NES)模拟器用围棋编写。
* [orange-cat](https://github.com/noraesae/orange-cat) - 用Go编写的Markdown预览器。
* [Orbit](https://github.com/gulien/orbit) - **Star : 128** 一个运行命令和从模板生成文件的简单工具。
* [peg](https://github.com/pointlander/peg) - **Star : 586** Peg，解析表达式语法，是Packrat解析器生成器的实现。
* [Pipe](https://github.com/b3log/pipe) - **Star : 2579** 一个小巧漂亮的博客平台。
* [restic](https://github.com/restic/restic) - **Star : 7148** 消除重复项备份程序。
* [rkt](https://github.com/coreos/rkt) - 与init系统集成的应用程序容器运行时，与其他容器格式(如Docker)兼容，并支持其他执行引擎(如KVM)。
* [scc](https://github.com/boyter/scc) - **Star : 759** 一个非常快速准确的代码计数器与复杂性计算和COCOMO估计。
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) - **Star : 7993** 快速、简单、可伸缩的分布式文件系统与O(1)磁盘查找。
* [shell2http](https://github.com/msoap/shell2http) - **Star : 395** 通过http服务器执行shell命令(用于原型或远程控制)。
* [snap](https://github.com/intelsdi-x/snap) - **Star : 1801** 强大的遥测框架。
* [Snitch](https://github.com/lucasgomide/snitch) - **Star : 15** 当有人通过Tsuru部署任何应用程序时，通知您的团队和许多工具的简单方法。
* [Stack Up](https://github.com/pressly/sup) - **Star : 1962** Stack Up是一个超级简单的部署工具—Unix—可以把它看作是一个服务器网络的“make”。
* [syncthing](https://syncthing.net/) - 开放，分散的文件同步工具和协议。
* [term-quiz](https://github.com/crazcalm/term-quiz) - **Star : 17** 为您的终端测试。
* [toxiproxy](https://github.com/shopify/toxiproxy) - **Star : 3765** 为自动化测试模拟网络和系统条件的代理。
* [tsuru](https://tsuru.io/) - 作为服务软件的可扩展开源平台。
* [vFlow](https://github.com/VerizonDigital/vflow) - **Star : 583** 高性能、可伸缩和可靠的IPFIX、sFlow和Netflow收集器。
* [wellington](https://github.com/wellington/wellington) - **Star : 288** Sass项目管理工具，使用sprite函数(如Compass)扩展语言。

# 资源

*在哪里可以找到新的Go库。*

## 基准

* [autobench](https://github.com/davecheney/autobench) - **Star : 89** 框架来比较不同Go版本之间的性能。
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - **Star : 19** 强大的HTTP-benchmark工具与Аb混合,Wrk围攻工具。收集统计和各种参数指标和比较结果。
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - **Star : 119** 很少有其他的Go微基准测试。将一些语言特性与其他方法进行比较。
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - **Star : 1253**  Go HTTP请求路由器基准和比较。
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - **Star : 973**  Go web框架基准测试。
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - **Star : 839** Go序列化方法的基准测试。
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Go语言常用基本操作的基准测试。
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) - **Star : 17** 小集合的Go微基准。其目的是将一些语言特性与其他特性进行比较。
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - **Star : 49** 为流行的Go数据库/SQL实用程序收集基准测试。
* [gospeed](https://github.com/feyeleanor/GoSpeed) - **Star : 93** 使用微基准测试来计算语言结构的速度。
* [kvbench](https://github.com/jimrobinson/kvbench) - **Star : 14** 键/值数据库基准。
* [skynet](https://github.com/atemerev/skynet) - **Star : 906** 天网1M线程微基准测试。
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - **Star : 169** 比较各种图像大小调整算法的围棋语言。

## 会议

* [Capital Go](http://www.capitalgolang.com) - 华盛顿特区。,美国。
* [dotGo](http://www.dotgo.eu) - 巴黎,法国。
* [GoCon](http://gocon.connpass.com/) - 东京,日本。
* [GoDays](https://www.godays.io/) - 德国柏林。
* [GoLab](http://golab.io/) - 佛罗伦萨,意大利。
* [GolangUK](http://golanguk.com/) - 伦敦,英国。
* [GopherChina](http://gopherchina.org) - 上海,中国。
* [GopherCon](http://www.gophercon.com/) - 美国丹佛。
* [GopherCon Brazil](https://gopherconbr.org) - 弗洛,BR。
* [GopherCon Europe](https://gophercon.is/) - 雷克雅未克,冰岛。
* [GopherCon India](https://www.gophercon.in/) - 印度浦那。
* [GopherCon Israel](https://www.gophercon.org.il/) - 特拉维夫,以色列。
* [GopherCon Russia](https://www.gophercon-russia.ru) - 莫斯科,俄罗斯。
* [GopherCon Singapore](https://gophercon.sg) - 新加坡枫树商贸城。
* [GopherCon Vietnam](https://gophercon.vn/) - 越南胡志明市。
* [GothamGo](http://gothamgo.com/) - 美国纽约市。
* [GoWayFest](https://goway.io/) - 白俄罗斯明斯克。

## 电子书

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org) - 一本关注围棋语法/语义和各种细节的书。
* [Go Bootcamp](http://golangbootcamp.com)
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) - **Star : 10** 在波斯。
* [GoBooks](https://github.com/dariubs/GoBooks) - **Star : 6653** 一份精选的围棋书籍清单。
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - **Star : 1456** Gopher图形包由玛丽亚莱塔与插图和情感字符在矢量和光栅。
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - **Star : 32**  Go 地鼠媒介数据[。ai . svg)。
* [gopher-logos](https://github.com/GolangUA/gopher-logos) - **Star : 62** 可爱的小田鼠标识。
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gopherize.me](https://github.com/matryer/gopherize.me) - **Star : 312** Gopherize自己。
* [gophers](https://github.com/ashleymcnamara/gophers) - **Star : 1806** 阿什莉·麦克纳马拉的歌斐艺术品。
* [gophers](https://github.com/egonelbre/gophers) - **Star : 1542** 免费打地鼠。
* [gophers](https://github.com/rogeralsing/gophers) - **Star : 49** 随机gopher图形。
* [gophers](https://github.com/sillecelik/go-gopher) - **Star : 41** Gopher amigurumi玩具图案。

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

*在这里添加您所在城市/国家的群组(发送**PR**)*

## 推特

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## 网站

* [Awesome Go @LibHunt](https://go.libhunt.com) - 你的工具箱。
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - **Star : 14338** 精心策划的远程工作列表。很多人都在寻找黑客。
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - **Star : 24442** 其他令人惊叹的列表。
* [CodinGame](https://www.codingame.com/) - 以小游戏为例，通过解决互动任务来学习围棋。
* [Go Blog](http://blog.golang.org) - 官方的Go博客。
* [Go Challenge](http://golang-challenge.org/) - 通过解决问题和从围棋专家那里得到反馈来学习围棋。
* [Go Community on Hashnode](https://hashnode.com/n/go) - 哈希节点上的地鼠社区。
* [Go Forum](https://forum.golangbridge.org) - 论坛讨论 Go 吧。
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5分钟的视频集中在一件事上。
* [Go Projects](https://github.com/golang/go/wiki/Projects) - Go社区wiki上的项目列表。
* [Go Report Card](https://goreportcard.com) - 为您的Go软件包准备一张成绩单。
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - **Star : 35** 收集需要帮助的Go项目。这是你开始开源之路的好地方。
* [godoc.org](https://godoc.org/) - 开源Go包的文档。
* [Golang Flow](https://golangflow.io) - 发布更新、新闻、包等等。
* [Golang News](https://golangnews.com) - 关于Go编程的链接和新闻。
* [golang-graphics](https://github.com/mholt/golang-graphics) - **Star : 141** 收藏的围棋图像，图形和艺术。
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - 邮件列表。
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - 谷歌+社区#golang爱好者。
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - 加入我们为Gophers开发的新Slack社区([了解它是如何产生的](https://blog.gopheracademy.com/gophers-slack-community/))。
* [Gophercises](https://gophercises.com/) - 初露头角的地鼠的免费编码练习。
* [gowalker.org](https://gowalker.org) -  Go 项目API文档。
* [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube频道致力于Go编程语言技巧和技巧，由Francesc Campoy [@francesc](https://twitter.com/francesc)主办。
* [r/Golang](https://www.reddit.com/r/golang) - 的消息。
* [studygolang](https://studygolang.com) - 中国学生群体。
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - 找到新的Go图书馆的好地方。
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### 教程

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - 陷阱、陷阱和新Golang开发者的常见错误。
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - 为电子商务构建一个Golang站点(包括演示)。
* [A Tour of Go](http://tour.golang.org/) - 互动的围棋之旅。
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - **Star : 30834** Golang电子书介绍如何与Golang建立一个web应用程序。
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - 熟悉Gin，了解它如何帮助您减少样板代码并构建请求处理管道。
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - 如何缓存慢速数据库查询。
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - 如何取消MySQL查询。
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) - **Star : 433** 这是一本关于用Go开发以太坊的电子书。
* [Games With Go](http://gameswithgo.org/) - 一个视频系列教学编程和游戏开发。
* [Go By Example](https://gobyexample.com/) - 动手介绍 Go 使用带注释的示例程序。
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - **Star : 3963**  Go 参考卡片。
* [Go database/sql tutorial](http://go-database-sql.org/) - 数据库概论/ sql。
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8) - 互动编辑和播放你的移动设备上的片段。
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - **Star : 656** Golang的例子与Node.js的学习相比较。
* [Golangbot](https://golangbot.com/learn-golang-series/) - 教程开始在围棋编程。
* [Hackr.io](https://hackr.io/tutorials/learn-golang) - 学习从最好的在线戈朗教程提交和戈朗编程社区投票。
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - 开始使用Godog——行为驱动开发框架来构建和测试应用程序。
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - **Star : 3800** 学习使用测试驱动开发。
* [package main](https://www.youtube.com/packagemain) - YouTube频道关于围棋编程。
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) - **Star : 1127** 介绍给有经验的程序员。
* [Your basic Go](http://yourbasic.org/golang) - 大量的教程和如何收集的。


