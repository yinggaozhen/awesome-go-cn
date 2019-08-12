# Awesome Go

[Gold]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Gold.svg "star > 5000"
[Silver]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Silver.svg "star > 1000"
[Bronze]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Bronze.svg "star > 100"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2019-08-12 10:37:49(Synchronize every day)**

[![english](https://yinggaozhen.github.io/docs/chinese.svg)](README.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) financial support to Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python)。

**Icon** :


Icon | State  
:-:|-
![star > 5000][Gold] | star  > 5000 
![star > 1000][Silver] | star > 1000
![star > 100][Bronze] | star > 100
![There was an update last week][Green] | There was an update last week。You can basically determine that the current library is in an actively maintained state。
![Not updated in last year][Yellow] | Not updated in last year。Reflects that the maintenance of the library is not high enthusiasm, use should be careful。
![Archived][Archived] | The project has been archived。
![Contains Chinese documents][CN] | This project contains Chinese documents。

### Explain

[中文](README.md)  | [English](README_EN.md)
 

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Bot Building](#bot-building)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Error Handling](#error-handling)
    - [Files](#files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Functional](#functional)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Geographic](#geographic)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot)
    - [Job Scheduler](#job-scheduler)
    - [JSON](#json)
    - [Logging](#logging)
    - [Machine Learning](#machine-learning)
    - [Messaging](#messaging)
    - [Microsoft Office](#microsoft-office)
        - [Microsoft Excel](#microsoft-excel)
    - [Miscellaneous](#miscellaneous)
        - [Dependency Injection](#dependency-injection)
        - [Project Layout](#project-layout)
        - [Strings](#strings)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
        - [HTTP Clients](#http-clients)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Performance](#performance)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Stream Processing](#stream-processing)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
    - [Text Processing](#text-processing)
    - [Third-party APIs](#third-party-apis)
    - [Utilities](#utilities)
    - [UUID](#uuid)
    - [Validation](#validation)
    - [Version Control](#version-control)
    - [Video](#video)
    - [Web Frameworks](#web-frameworks)
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
        - [Routers](#routers)
    - [Windows](#windows)
    - [XML](#xml)

- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)

- [Server Applications](#server-applications)

- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [Oto](https://github.com/hajimehoshi/oto) **star:438** A low-level library to play sound on multiple platforms.   ![star > 100][Bronze]
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:303** Go bindings for the PortAudio audio I/O library.   ![star > 100][Bronze]
* [music-theory](https://github.com/go-music-theory/music-theory) **star:257** Music theory models in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [waveform](https://github.com/mdlayher/waveform) **star:248** Go package capable of generating waveform images from audio streams.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [portmidi](https://github.com/rakyll/portmidi) **star:207** Go bindings for PortMidi.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [id3v2](https://github.com/bogem/id3v2) **star:110** Fast and stable ID3 parsing and writing library for Go.   ![star > 100][Bronze]
* [flac](https://github.com/mewkiz/flac) **star:102** Native Go FLAC encoder/decoder with support for FLAC streams.   ![star > 100][Bronze]
* [mix](https://github.com/go-mix/mix) **star:98** Sequence-based Go-native audio mixer for music apps.   ![It hasn't been updated in the last year][Yellow]
* [mp3](https://github.com/tcolgate/mp3) **star:94** Native Go MP3 decoder.   ![It hasn't been updated in the last year][Yellow]
* [go-sox](https://github.com/krig/go-sox) **star:93** libsox bindings for go.   ![It hasn't been updated in the last year][Yellow]
* [flac](https://github.com/eaburns/flac) **star:83** No-frills native Go FLAC decoder that decodes FLAC files into byte slices.   ![It hasn't been updated in the last year][Yellow]
* [malgo](https://github.com/gen2brain/malgo) **star:71** Mini audio library.
* [taglib](https://github.com/wtolson/go-taglib) **star:65** Go bindings for taglib.   ![It hasn't been updated in the last year][Yellow]
* [gaad](https://github.com/Comcast/gaad) **star:56** Native Go AAC bitstream parser.   ![It hasn't been updated in the last year][Yellow]
* [minimp3](https://github.com/tosone/minimp3) **star:25** Lightweight MP3 decoder library.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:24** libmediainfo bindings for go.   ![It hasn't been updated in the last year][Yellow]
* [vorbis](https://github.com/mccoyst/vorbis) **star:22** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:22** EasyMidi is a simple and reliable library for working with standard midi file (SMF).   ![It hasn't been updated in the last year][Yellow]
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:8** libsamplerate bindings for go.   ![It hasn't been updated in the last year][Yellow]

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:6026** Golang implementation of JSON Web Tokens (JWT).   ![star > 5000][Gold]
* [casbin](https://github.com/hsluoyz/casbin) **star:4983** Authorization library that supports access control models like ACL, RBAC, ABAC.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [oauth2](https://github.com/golang/oauth2) **star:2386** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   ![star > 1000][Silver]
* [goth](https://github.com/markbates/goth) **star:2279** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [authboss](https://github.com/volatiletech/authboss) **star:1936** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   ![star > 1000][Silver]
* [osin](https://github.com/openshift/osin) **star:1546** Golang OAuth2 server library.   ![star > 1000][Silver]
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1285** Standalone, specification-compliant,  OAuth2 server written in Golang.   ![star > 1000][Silver]
* [go-jose](https://github.com/square/go-jose) **star:1126** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.   ![star > 1000][Silver]
* [gologin](https://github.com/dghubble/gologin) **star:1042** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gorbac](https://github.com/mikespook/gorbac) **star:912** provides a lightweight role-based access control (RBAC) implementation in Golang.   ![star > 100][Bronze]
* [loginsrv](https://github.com/tarent/loginsrv) **star:820** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   ![star > 100][Bronze]
* [scs](https://github.com/alexedwards/scs) **star:529** Session Manager for HTTP servers.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [permissions2](https://github.com/xyproto/permissions2) **star:352** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   ![star > 100][Bronze]
* [paseto](https://github.com/o1egl/paseto) **star:242** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   ![star > 100][Bronze]
* [httpauth](https://github.com/goji/httpauth) **star:180** HTTP Authentication middleware.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:153** JWT middleware for Golang http servers with many configuration options.   ![star > 100][Bronze]
* [jwt](https://github.com/pascaldekloe/jwt) **star:94** Lightweight JSON Web Token (JWT) library.
* [session](https://github.com/icza/session) **star:91** Go session management for web servers (including support for Google App Engine - GAE).
* [branca](https://github.com/hako/branca) **star:76** Golang implementation of Branca Tokens.   ![It hasn't been updated in the last year][Yellow]
* [jwt](https://github.com/robbert229/jwt) **star:71** Clean and easy to use implementation of JSON Web Tokens (JWT).
* [sessions](https://github.com/adam-hanna/sessions) **star:47** Dead simple, highly performant, highly customizable sessions service for go http servers.
* [securecookie](https://github.com/chmike/securecookie) **star:32** Efficient secure cookie encoding/decoding.
* [sjwt](https://github.com/brianvoe/sjwt) **star:32** Simple jwt generator and parser.
* [rbac](https://github.com/zpatrick/rbac) **star:27** Minimalistic RBAC package for Go applications.
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.
* [signedvalue](https://github.com/sashka/signedvalue) **star:7** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** provides parser of cookies.txt file format.   ![It hasn't been updated in the last year][Yellow]

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1634** Simple and clean Telegram bot client.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [telebot](https://github.com/tucnak/telebot) **star:953** Telegram bot framework written in Go.   ![star > 100][Bronze]
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:469** IRC, Slack & Telegram bot written in Go.   ![star > 100][Bronze]
* [slacker](https://github.com/shomali11/slacker) **star:316** Easy to use framework to create Slack bots.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:223** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   ![star > 100][Bronze]
* [tbot](https://github.com/yanzay/tbot) **star:220** Telegram bot server with API similar to net/http.   ![star > 100][Bronze]
* [Tenyks](https://github.com/kyleterry/tenyks) **star:167** Service oriented IRC bot using Redis and JSON for messaging.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Kelp](https://github.com/stellar/kelp) **star:163** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   ![star > 100][Bronze]
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:137** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [hanu](https://github.com/sbstjn/hanu) **star:109** Framework for writing Slack bots.   ![star > 100][Bronze]
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:85** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   ![It hasn't been updated in the last year][Yellow]
* [margelet](https://github.com/zhulik/margelet) **star:56** Framework for building Telegram bots.   ![It hasn't been updated in the last year][Yellow]
* [govkbot](https://github.com/nikepan/govkbot) **star:25** Simple Go [VK](https://vk.com) bot library.
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:11** Another framework for building Slack bots.
* [micha](https://github.com/onrik/micha) **star:10** Go Library for Telegram bot api.   ![It hasn't been updated in the last year][Yellow]

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:13389** Commander for modern Go CLI interactions.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [urfave/cli](https://github.com/urfave/cli) **star:11466** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   ![star > 5000][Gold]   ![There was an update last week][Green]
* [kingpin](https://github.com/alecthomas/kingpin) **star:2536** Command line and flag parser supporting sub commands.   ![star > 1000][Silver]
* [go-flags](https://github.com/jessevdk/go-flags) **star:1508** go command line option parser.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [readline](https://github.com/chzyer/readline) **star:1371** Pure golang implementation that provides most features in GNU-Readline under MIT license.   ![star > 1000][Silver]
* [docopt.go](https://github.com/docopt/docopt.go) **star:1174** Command-line arguments parser that will make you smile.   ![star > 1000][Silver]
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1000** Go library for implementing command-line interfaces.   ![star > 1000][Silver]
* [cli-init](https://github.com/tcnksm/gcli) **star:869** The easy way to start building Golang command line applications.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [pflag](https://github.com/spf13/pflag) **star:753** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   ![star > 100][Bronze]
* [go-arg](https://github.com/alexflint/go-arg) **star:730** Struct-based argument parsing in Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [mow.cli](https://github.com/jawher/mow.cli) **star:623** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   ![star > 100][Bronze]
* [complete](https://github.com/posener/complete) **star:620** Write bash completions in Go + Go command bash completion.   ![star > 100][Bronze]
* [liner](https://github.com/peterh/liner) **star:588** Go readline-like library for command-line interfaces.   ![star > 100][Bronze]
* [cli](https://github.com/mkideal/cli) **star:479** Feature-rich and easy to use command-line package based on golang struct tags.   ![star > 100][Bronze]
* [flaggy](https://github.com/integrii/flaggy) **star:453** A robust and idiomatic flags package with excellent subcommand support.   ![star > 100][Bronze]
* [ops](https://github.com/nanovms/ops) **star:264** Unikernel Builder/Orchestrator.   ![star > 100][Bronze]
* [argparse](https://github.com/akamensky/argparse) **star:109** Command line argument parser inspired by Python's argparse module.   ![star > 100][Bronze]
* [flag](https://github.com/cosiner/flag) **star:101** Simple but powerful command line option parsing library for Go supporting subcommand.   ![star > 100][Bronze]
* [ukautz/clif](https://github.com/ukautz/clif) **star:97** Small command line interface framework.
* [commandeer](https://github.com/jaffee/commandeer) **star:93** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.
* [sflags](https://github.com/octago/sflags) **star:91** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.
* [wmenu](https://github.com/dixonwille/wmenu) **star:85** Easy to use menu structure for cli applications that prompts users to make choices.
* [cli](https://github.com/teris-io/cli) **star:57** Simple and complete API for building command line interfaces in Go.
* [job](https://github.com/liujianping/job) **star:50** JOB, make your short-term command as a long-term job.
* [env](https://github.com/codingconcepts/env) **star:42** Tag-based environment configuration for structs.
* [wlog](https://github.com/dixonwille/wlog) **star:35** Simple logging interface that supports cross-platform color and concurrency.   ![It hasn't been updated in the last year][Yellow]
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [gocmd](https://github.com/devfacet/gocmd) **star:33** Go library for building command line applications.
* [flagvar](https://github.com/sgreben/flagvar) **star:31** A collection of flag argument types for Go's standard `flag` package.
* [strumt](https://github.com/antham/strumt) **star:27** Library to create prompt chain.
* [argv](https://github.com/cosiner/argv) **star:17** Go library to split command line string as arguments array using the bash syntax.
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.
* [cmdr](https://github.com/hedzr/cmdr) **star:15** A POSIX/GNU style, getopt-like command-line UI Go library.   ![There was an update last week][Green]
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:8** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   ![There was an update last week][Green]
* [sand](https://github.com/Zaba505/sand) **star:6** Simple API for creating interpreters and so much more.
* [ts](https://github.com/liujianping/ts) **star:4** Timestamp convert & compare tool.

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:8942** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   ![star > 5000][Gold]
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [gocui](https://github.com/jroimartin/gocui) **star:5390** Minimalist Go library aimed at creating Console User Interfaces.   ![star > 5000][Gold]
* [termbox-go](https://github.com/nsf/termbox-go) **star:3492** Termbox is a library for creating cross-platform text-based interfaces.   ![star > 1000][Silver]
* [color](https://github.com/fatih/color) **star:3023** Versatile package for colored terminal output.   ![star > 1000][Silver]   ![Archived][Archived]
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2341** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   ![star > 1000][Silver]
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1543** Flexible library to render progress bars in terminal applications.   ![star > 1000][Silver]
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1146** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   ![star > 1000][Silver]
* [uilive](https://github.com/gosuri/uilive) **star:833** Library for updating terminal output in realtime.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [termdash](https://github.com/mum4k/termdash) **star:812** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   ![star > 100][Bronze]
* [mpb](https://github.com/vbauerster/mpb) **star:705** Multi progress bar for terminal applications.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [aurora](https://github.com/logrusorgru/aurora) **star:637** ANSI terminal colors that supports fmt.Printf/Sprintf.   ![star > 100][Bronze]
* [progressbar](https://github.com/schollz/progressbar) **star:578** Basic thread-safe progress bar that works in every OS.   ![star > 100][Bronze]
* [uitable](https://github.com/gosuri/uitable) **star:504** Library to improve readability in terminal apps using tabular data.   ![star > 100][Bronze]
* [go-colorable](https://github.com/mattn/go-colorable) **star:377** Colorable writer for windows.   ![star > 100][Bronze]
* [go-isatty](https://github.com/mattn/go-isatty) **star:345** isatty for golang.   ![star > 100][Bronze]
* [chalk](https://github.com/ttacon/chalk) **star:304** Intuitive package for prettifying terminal/console output.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tabby](https://github.com/cheynewallace/tabby) **star:248** A tiny library for super simple Golang tables.   ![star > 100][Bronze]
* [termtables](https://github.com/apcera/termtables) **star:214** Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gookit/color](https://github.com/gookit/color) **star:206** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   ![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:196** Go library for color output in terminals.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [simpletable](https://github.com/alexeyco/simpletable) **star:168** Simple tables in terminal with Go.   ![star > 100][Bronze]
* [cfmt](https://github.com/mingrammer/cfmt) **star:67** Contextual fmt inspired by bootstrap color classes.
* [tabular](https://github.com/InVisionApp/tabular) **star:29** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   ![It hasn't been updated in the last year][Yellow]
* [colourize](https://github.com/TreyBastian/colourize) **star:16** Go library for ANSI colour text in terminals.   ![It hasn't been updated in the last year][Yellow]
* [ctc](https://github.com/wzshiming/ctc) **star:10** The non-invasive cross-platform terminal color library does not need to modify the Print method.   ![Contains Chinese documents][CN]
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** Go library for rendering ANSI colored text templates in terminals.   ![It hasn't been updated in the last year][Yellow]

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:9379** Go configuration with fangs.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2413** Go library for managing configuration data from environment variables.   ![star > 1000][Silver]
* [godotenv](https://github.com/joho/godotenv) **star:2135** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   ![star > 1000][Silver]
* [ini](https://github.com/go-ini/ini) **star:1598** Go package to read and write INI files.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [env](https://github.com/caarlos0/env) **star:874** Parse environment variables to Go structs (with defaults).   ![star > 100][Bronze]
* [konfig](https://github.com/lalamove/konfig) **star:515** Composable, observable and performant config handling for Go for the distributed processing era.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [store](https://github.com/tucnak/store) **star:240** Lightweight configuration manager for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [confita](https://github.com/heetch/confita) **star:238** Load configuration in cascade from multiple backends into a struct.   ![star > 100][Bronze]
* [config](https://github.com/olebedev/config) **star:210** JSON or YAML configuration wrapper with environment variables and flags parsing.   ![star > 100][Bronze]
* [joshbetz/config](https://github.com/joshbetz/config) **star:192** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [config](https://github.com/JeremyLoy/config) **star:190** Cloud native application configuration. Bind ENV to structs in only two lines.   ![star > 100][Bronze]
* [hjson](https://github.com/hjson/hjson-go) **star:172** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   ![star > 100][Bronze]
* [envconfig](https://github.com/vrischmann/envconfig) **star:144** Read your configuration from environment variables.   ![star > 100][Bronze]
* [gcfg](https://github.com/go-gcfg/gcfg) **star:116** read INI-style configuration files into Go structs; supports user-defined types and subsections.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [goConfig](https://github.com/crgimenes/goConfig) **star:110** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   ![star > 100][Bronze]
* [envh](https://github.com/antham/envh) **star:92** Helpers to manage environment variables.
* [envcfg](https://github.com/tomazk/envcfg) **star:89** Un-marshaling environment variables to Go structs.   ![It hasn't been updated in the last year][Yellow]
* [koanf](https://github.com/knadh/koanf) **star:86** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.
* [gookit/config](https://github.com/gookit/config) **star:85** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   ![Contains Chinese documents][CN]
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gofigure](https://github.com/ian-kent/gofigure) **star:57** Go application configuration made easy.   ![It hasn't been updated in the last year][Yellow]
* [configure](https://github.com/paked/configure) **star:49** Provides configuration through multiple sources, including JSON, flags and environment variables.
* [harvester](https://github.com/beatlabs/harvester) **star:40** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:37** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).
* [ingo](https://github.com/schachmat/ingo) **star:24** Flags persisted in an ini-like config file.   ![It hasn't been updated in the last year][Yellow]
* [go-up](https://github.com/ufoscout/go-up) **star:24** A simple configuration library with recursive placeholders resolution and no magic.
* [mini](https://github.com/sasbury/mini) **star:19** Golang package for parsing ini-style configuration files.
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:8** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
* [envconf](https://github.com/ian-kent/envconf) **star:7** Configuration from environment.   ![It hasn't been updated in the last year][Yellow]
* [sprbox](https://github.com/oblq/sprbox) **star:3** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:19078** Drone is a Continuous Integration platform built on Docker, written in Go.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [goveralls](https://github.com/mattn/goveralls) **star:578** Go integration for Coveralls.io continuous code coverage tracking system.   ![star > 100][Bronze]
* [overalls](https://github.com/go-playground/overalls) **star:98** Multi-Package go project coverprofile for tools like goveralls.
* [duci](https://github.com/duck8823/duci) **star:44** A simple ci server no needs domain specific languages.   ![There was an update last week][Green]
* [gomason](https://github.com/nikogura/gomason) **star:33** Test, Build, Sign, and Publish your go binaries from a clean workspace.
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** Recursive coverage testing tool.   ![It hasn't been updated in the last year][Yellow]

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:422** Pure Go CSS Preprocessor.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-libsass](https://github.com/wellington/go-libsass) **star:130** Go wrapper to the 100% Sass compatible libsass project.   ![star > 100][Bronze]

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:6417** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   ![star > 5000][Gold]
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5136** Collection of useful, performant, and thread-safe data structures.   ![star > 5000][Gold]
* [golang-set](https://github.com/deckarep/golang-set) **star:1172** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   ![star > 1000][Silver]
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1161** Probabilistic data structures for processing continuous, unbounded streams.   ![star > 1000][Silver]
* [gota](https://github.com/kniren/gota) **star:883** Implementation of dataframes, series, and data wrangling methods for Go.   ![star > 100][Bronze]
* [willf/bloom](https://github.com/willf/bloom) **star:667** Go package implementing Bloom filters.   ![star > 100][Bronze]
* [roaring](https://github.com/RoaringBitmap/roaring) **star:667** Go package implementing compressed bitsets.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:661** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   ![star > 100][Bronze]
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:515** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   ![star > 100][Bronze]
* [bitset](https://github.com/willf/bitset) **star:480** Go package implementing bitsets.   ![star > 100][Bronze]
* [trie](https://github.com/derekparker/trie) **star:422** Trie implementation in Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:312** In-memory geo index.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mafsa](https://github.com/smartystreets/mafsa) **star:272** MA-FSA implementation with Minimal Perfect Hashing.   ![star > 100][Bronze]   ![Archived][Archived]
* [algorithms](https://github.com/shady831213/algorithms) **star:241** Algorithms and data structures.CLRS study.   ![star > 100][Bronze]
* [goskiplist](https://github.com/ryszard/goskiplist) **star:193** Skip list implementation in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hilbert](https://github.com/google/hilbert) **star:182** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   ![star > 100][Bronze]
* [merkletree](https://github.com/cbergoon/merkletree) **star:146** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   ![star > 100][Bronze]
* [bloom](https://github.com/zhenjl/bloom) **star:127** Bloom filters implemented in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:125** Binary packer and unpacker helps user build custom binary stream.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [skiplist](https://github.com/MauriceGit/skiplist) **star:101** Very fast Go Skiplist implementation.   ![star > 100][Bronze]
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:100** Region quadtrees with efficient point location and neighbour finding.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:98** In-memory LRU string-interface{} map with expiration for golang.
* [encoding](https://github.com/zhenjl/encoding) **star:94** Integer Compression Libraries for Go.   ![It hasn't been updated in the last year][Yellow]
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:89** Go implementation of Adaptive Radix Tree.
* [ring](https://github.com/TheTannerRyan/ring) **star:89** Go implementation of a high performance, thread safe bloom filter.
* [conjungo](https://github.com/InVisionApp/conjungo) **star:79** A small, powerful and flexible merge library.
* [deque](https://github.com/gammazero/deque) **star:66** Fast ring-buffer deque (double-ended queue).
* [skiplist](https://github.com/gansidui/skiplist) **star:64** Skiplist implementation in Go.   ![It hasn't been updated in the last year][Yellow]
* [levenshtein](https://github.com/agnivade/levenshtein) **star:55** Implementation to calculate levenshtein distance in Go.
* [bit](https://github.com/yourbasic/bit) **star:54** Golang set data structure with bonus bit-twiddling functions.   ![It hasn't been updated in the last year][Yellow]
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:43** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   ![It hasn't been updated in the last year][Yellow]
* [bloom](https://github.com/yourbasic/bloom) **star:39** Golang Bloom filter implementation.   ![It hasn't been updated in the last year][Yellow]
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:35** Fast in-memory key:value store/cache library. Pointer caches.
* [levenshtein](https://github.com/agext/levenshtein) **star:33** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:27** Concurrent FIFO queue.
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:24** Highly concurrent drop-in replacement for `bufio.Writer`.   ![It hasn't been updated in the last year][Yellow]
* [crunch](https://github.com/superwhiskers/crunch) **star:19** Go package implementing buffers for handling various datatypes easily.   ![There was an update last week][Green]
* [goset](https://github.com/zoumo/goset) **star:16** A useful Set collection implementation for Go.   ![There was an update last week][Green]
* [pipeline](https://github.com/hyfather/pipeline) **star:15** An implementation of pipelines with fan-in and fan-out.
* [ptrie](https://github.com/viant/ptrie)  An implementation of prefix tree.
* [go-ef](https://github.com/amallia/go-ef) **star:11** A Go implementation of the Elias-Fano encoding.   ![It hasn't been updated in the last year][Yellow]
* [typ](https://github.com/gurukami/typ) **star:9** Null Types, Safe primitive type conversion and fetching value from complex structures.
* [dict](https://github.com/srfrog/dict) **star:9** Python-like dictionaries (dict) for Go.
* [mspm](https://github.com/BlackRabbitt/mspm) **star:7** Multi-String Pattern Matching Algorithm for information retrieval.   ![It hasn't been updated in the last year][Yellow]
* [hide](https://github.com/emvi/hide) **star:7** ID type with marshalling to/from hash to prevent sending IDs to clients.
* [treap](https://github.com/perdata/treap) **star:7** Persistent, fast ordered map using tree heaps.
* [deque](https://github.com/edwingeng/deque) **star:7** A highly optimized double-ended queue.
* [set](https://github.com/StudioSol/set) **star:6** Simple set data structure implementation in Go using LinkedHashMap.
* [null](https://github.com/emvi/null) **star:5** Nullable Go types that can be marshalled/unmarshalled to/from JSON.
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.
* [timedmap](https://github.com/zekroTJA/timedmap) **star:1** Map with expiring key-value pairs.

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:25638** Monitoring system and time series database.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [tidb](https://github.com/pingcap/tidb) **star:20119** TiDB is a distributed SQL database. Inspired by the design of Google F1.   ![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [influxdb](https://github.com/influxdb/influxdb) **star:17138** Scalable datastore for metrics, events, and real-time analytics.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [cockroach](https://github.com/cockroachdb/cockroach) **star:16803** Scalable, Geo-Replicated, Transactional Datastore.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [dgraph](https://github.com/dgraph-io/dgraph) **star:10480** Scalable, Distributed, Low Latency, High Throughput Graph Database.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [bolt](https://github.com/boltdb/bolt) **star:9993** Low-level key/value database for Go.   ![star > 5000][Gold]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [groupcache](https://github.com/golang/groupcache) **star:7669** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   ![star > 5000][Gold]
* [badger](https://github.com/dgraph-io/badger) **star:6295** Fast key-value store in Go.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [rqlite](https://github.com/rqlite/rqlite) **star:4701** The lightweight, distributed, relational database built on SQLite.   ![star > 1000][Silver]
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3172** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3079** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   ![star > 1000][Silver]
* [go-cache](https://github.com/pmylund/go-cache) **star:2916** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   ![star > 1000][Silver]
* [BigCache](https://github.com/allegro/bigcache) **star:2463** Efficient key/value cache for gigabytes of data.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [buntdb](https://github.com/tidwall/buntdb) **star:2443** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   ![star > 1000][Silver]
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2367** Your NoSQL database powered by Golang.   ![star > 1000][Silver]
* [cache2go](https://github.com/muesli/cache2go) **star:1052** In-memory key:value cache which supports automatic invalidation based on timeouts.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1019** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [GCache](https://github.com/bluele/gcache) **star:900** Cache library with support for expirable Cache, LFU, LRU and ARC.   ![star > 100][Bronze]
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:891** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:877** CovenantSQL is a SQL database on blockchain.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [diskv](https://github.com/peterbourgon/diskv) **star:751** Home-grown disk-backed key-value store.   ![star > 100][Bronze]
* [moss](https://github.com/couchbase/moss) **star:718** Moss is a simple LSM key-value storage engine written in 100% Go.   ![star > 100][Bronze]
* [eliasdb](https://github.com/krotik/eliasdb) **star:532** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   ![star > 100][Bronze]
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:489** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   ![star > 100][Bronze]
* [levigo](https://github.com/jmhodges/levigo) **star:363** Levigo is a Go wrapper for LevelDB.   ![star > 100][Bronze]
* [pudge](https://github.com/recoilme/pudge) **star:218** Fast and simple  key/value store written using Go's standard library.   ![star > 100][Bronze]
* [piladb](https://github.com/fern4lvarez/piladb) **star:171** Lightweight RESTful database engine based on stack data structures.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Vasto](https://github.com/chrislusf/vasto) **star:149** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   ![star > 100][Bronze]
* [slowpoke](https://github.com/recoilme/slowpoke) **star:86** Key-value store with persistence.
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:61** Tiny flat file JSON store.
* [couchcache](https://github.com/codingsince1985/couchcache) **star:40** RESTful caching micro-service backed by Couchbase server.
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:30** BigCache with clustering support and individual item expiration.   ![It hasn't been updated in the last year][Yellow]
* [bcache](https://github.com/iwanbk/bcache) **star:29** Eventually consistent distributed in-memory  cache Go library.
* [cache](https://github.com/akyoto/cache) **star:16** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:13** Key-value store for temporary items.   ![It hasn't been updated in the last year][Yellow]
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:8** Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.   ![It hasn't been updated in the last year][Yellow]

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:2762** Database migrations. CLI and Golang library.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1415** Database migration tool. Allows embedding migrations into the application using go-bindata.   ![star > 1000][Silver]
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:335** Database schema migration helper for Gorm ORM.   ![star > 100][Bronze]
* [goose](https://github.com/steinbacher/goose) **star:119** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [darwin](https://github.com/GuiaBolso/darwin) **star:83** Database schema evolution library for Go.
* [migrator](https://github.com/lopezator/migrator) **star:31** Dead simple Go database migration library.   ![There was an update last week][Green]
* [gondolier](https://github.com/emvi/gondolier) **star:26** Database migration library using struct decorators.
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:23** A Go package to help write migrations with go-pg/pg.   ![There was an update last week][Green]
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:20** Django style fixtures for Golang's excellent built-in database/sql library.
* [avro](https://github.com/khezen/avro) **star:6** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   ![There was an update last week][Green]

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:8466** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [pgweb](https://github.com/sosedoff/pgweb) **star:5996** Web-based PostgreSQL database browser.   ![star > 5000][Gold]
* [kingshard](https://github.com/flike/kingshard) **star:4620** kingshard is a high performance proxy for MySQL powered by Golang.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [orchestrator](https://github.com/github/orchestrator) **star:3044** MySQL replication topology manager & visualizer.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2402** Sync your MySQL data into Elasticsearch automatically.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [pREST](https://github.com/nuveo/prest) **star:2090** Serve a RESTful API from any PostgreSQL database.   ![star > 1000][Silver]
* [go-mysql](https://github.com/siddontang/go-mysql) **star:1894** Go toolset to handle MySQL protocol and replication.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [chproxy](https://github.com/Vertamedia/chproxy) **star:308** HTTP proxy for ClickHouse database.   ![star > 100][Bronze]
* [myreplication](https://github.com/2tvenom/myreplication) **star:141** MySql binary log replication listener. Supports statement and row based replication.   ![star > 100][Bronze]
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:139** Collects small insterts and sends big requests to ClickHouse servers.   ![star > 100][Bronze]
* [octillery](https://github.com/knocknote/octillery) **star:54** Go package for sharding databases ( Supports every ORM or raw SQL ).
* [dbbench](https://github.com/sj14/dbbench) **star:30** Database benchmarking tool with support for several databases and scripts.
* [prep](https://github.com/hexdigest/prep) **star:24** Use prepared SQL statements without changing your code.   ![It hasn't been updated in the last year][Yellow]
* [rwdb](https://github.com/andizzle/rwdb) **star:10** rwdb provides read replica capability for multiple database servers setup.   ![It hasn't been updated in the last year][Yellow]
* [datagen](https://github.com/codingconcepts/datagen) **star:8** A fast data generator that's multi-table aware and supports multi-row DML.

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2327** Go library that helps you build SQL queries.   ![star > 1000][Silver]
* [xo](https://github.com/knq/xo) **star:2189** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   ![star > 1000][Silver]
* [gendry](https://github.com/didi/gendry) **star:761** Non-invasive SQL builder and powerful data binder.   ![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [goqu](https://github.com/doug-martin/goqu) **star:638** Idiomatic SQL builder and query library.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Dotsql](https://github.com/gchaincl/dotsql) **star:440** Go library that helps you keep sql files in one place and use them with ease.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:436** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   ![star > 100][Bronze]
* [sqrl](https://github.com/elgris/sqrl) **star:178** SQL query builder, fork of Squirrel with improved performance.   ![star > 100][Bronze]
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [scaneo](https://github.com/variadico/scaneo) **star:149** Generate Go code to convert database rows into arbitrary structs.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [igor](https://github.com/galeone/igor) **star:78** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   ![It hasn't been updated in the last year][Yellow]
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [godbal](https://github.com/xujiajun/godbal) **star:50** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:8115** MySQL driver for Go.   ![star > 5000][Gold]
    * [pq](https://github.com/lib/pq) **star:5198** Pure Go Postgres driver for database/sql.   ![star > 5000][Gold]
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3432** SQLite3 driver for go that uses database/sql.   ![star > 1000][Silver]
    * [pgx](https://github.com/jackc/pgx) **star:1968** PostgreSQL driver supporting features beyond those exposed by database/sql.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1022** Microsoft MSSQL driver for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [go-oci8](https://github.com/mattn/go-oci8) **star:405** Oracle driver for go that uses database/sql.   ![star > 100][Bronze]
    * [goracle](https://github.com/go-goracle/goracle) **star:239** Oracle driver for Go, using the ODPI-C driver.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:104** Firebird RDBMS SQL driver for Go.   ![star > 100][Bronze]
    * [go-adodb](https://github.com/mattn/go-adodb) **star:90** Microsoft ActiveX Object DataBase driver for go that uses database/sql.
    * [gofreetds](https://github.com/minus5/gofreetds) **star:90** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:35** Apache Avatica/Phoenix SQL driver for database/sql.
    * [bgc](https://github.com/viant/bgc) **star:12** Datastore Connectivity for BigQuery for go.   ![There was an update last week][Green]

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:6529** Redis client for Golang.   ![star > 5000][Gold]   ![There was an update last week][Green]
    * [redigo](https://github.com/gomodule/redigo) **star:6300** Redigo is a Go client for the Redis database.   ![star > 5000][Gold]
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:3131** Official MongoDB driver for the Go language.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [mgo](https://github.com/globalsign/mgo) **star:1646** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [gorethink](https://github.com/dancannon/gorethink) **star:1462** Go language driver for RethinkDB.   ![star > 1000][Silver]
    * [neoism](https://github.com/jmcvetta/neoism) **star:356** Neo4j client for Golang.   ![star > 100][Bronze]
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:304** Aerospike client in Go language.   ![star > 100][Bronze]
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:292** Couchbase client in Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [gocb](https://github.com/couchbase/gocb) **star:291** Official Couchbase Go SDK.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [redeo](https://github.com/bsm/redeo) **star:260** Redis-protocol compatible TCP servers/services.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:90** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   ![There was an update last week][Green]
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:72** Neo4j REST Client in golang.   ![It hasn't been updated in the last year][Yellow]
    * [arangolite](https://github.com/solher/arangolite) **star:65** Lightweight golang driver for ArangoDB.
    * [dynago](https://github.com/underarmour/dynago) **star:65** Dynago is a principle of least surprise client for DynamoDB.   ![It hasn't been updated in the last year][Yellow]
    * [godis](https://github.com/piaohao/godis) **star:60** redis client implement by golang, inspired by jedis.
    * [go-couchdb](https://github.com/fjl/go-couchdb) **star:51** Yet another CouchDB HTTP API wrapper for Go.
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:31** Go client library for Pilosa.
    * [forestdb](https://github.com/couchbase/goforestdb) **star:29** Go bindings for ForestDB.   ![It hasn't been updated in the last year][Yellow]
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.
    * [neo4j](https://github.com/cihangir/neo4j) **star:24** Neo4j Rest API Bindings for Golang.   ![It hasn't been updated in the last year][Yellow]
    * [xredis](https://github.com/shomali11/xredis) **star:9** Typesafe, customizable, clean & easy to use Redis client.
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [godscache](https://github.com/defcronyke/godscache) **star:6** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
    * [asc](https://github.com/viant/asc) **star:4** Datastore Connectivity for Aerospike for go.

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:5848** Modern text indexing library for go.   ![star > 5000][Gold]   ![There was an update last week][Green]
    * [riot](https://github.com/go-ego/riot) **star:4702** Go Open Source, Distributed, Simple and efficient Search Engine.   ![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
    * [elastic](https://github.com/olivere/elastic) **star:4159** Elasticsearch client for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:1588** Official Elasticsearch client for Go.   ![star > 1000][Silver]
    * [elastigo](https://github.com/mattbaird/elastigo) **star:950** Elasticsearch client library.   ![star > 100][Bronze]
    * [elasticsql](https://github.com/cch123/elasticsql) **star:397** Convert sql to elasticsearch dsl in Go.   ![star > 100][Bronze]
    * [skizze](https://github.com/seiflotfy/skizze) **star:68** probabilistic data-structures service and storage.   ![It hasn't been updated in the last year][Yellow]
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   ![It hasn't been updated in the last year][Yellow]

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:12691** Graph database with support for multiple backends.   ![star > 5000][Gold]
    * [cachego](https://github.com/fabiorphp/cachego) **star:110** Golang Cache component for multiple drivers.   ![star > 100][Bronze]
    * [gokv](https://github.com/philippgille/gokv) **star:85** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).
    * [dsc](https://github.com/viant/dsc) **star:14** Datastore connectivity for SQL, NoSQL, structured files.   ![There was an update last week][Green]

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2186** Now is a time toolkit for golang.   ![star > 1000][Silver]
* [dateparse](https://github.com/araddon/dateparse) **star:896** Parse date's without knowing format in advance.   ![star > 100][Bronze]
* [carbon](https://github.com/uniplaces/carbon) **star:339** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   ![star > 100][Bronze]
* [durafmt](https://github.com/hako/durafmt) **star:241** Time duration formatting library for Go.   ![star > 100][Bronze]
* [timeutil](https://github.com/leekchan/timeutil) **star:169** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   ![star > 100][Bronze]
* [iso8601](https://github.com/relvacode/iso8601) **star:68** Efficiently parse ISO8601 date-times without regex.
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:64** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
* [timespan](https://github.com/SaidinWoT/timespan) **star:60** For interacting with intervals of time, defined as a start time and a duration.
* [date](https://github.com/rickb777/date) **star:28** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
* [feiertage](https://github.com/wlbr/feiertage) **star:22** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
* [goweek](https://github.com/grsmv/goweek) **star:18** Library for working with week entity in golang.   ![It hasn't been updated in the last year][Yellow]
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:13** Calculate the sunrise and sunset times for a given location.   ![It hasn't been updated in the last year][Yellow]
* [kair](https://github.com/GuilhermeCaruso/kair) **star:11** Date and Time - Golang Formatting Library.
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:9** Nullable `time.Time`.   ![It hasn't been updated in the last year][Yellow]
* [tuesday](https://github.com/osteele/tuesday) **star:7** Ruby-compatible Strftime function.   ![It hasn't been updated in the last year][Yellow]
* [strftime](https://github.com/awoodbeck/strftime) **star:5** C99-compatible strftime formatter.   ![It hasn't been updated in the last year][Yellow]

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:14484** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [grpc-go](https://github.com/grpc/grpc-go) **star:9142** The Go language implementation of gRPC. HTTP/2 based RPC.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [micro](https://github.com/micro/micro) **star:6604** Pluggable microservice toolkit and distributed systems platform.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [NATS](https://github.com/nats-io/gnatsd) **star:6358** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [rpcx](https://github.com/smallnest/rpcx) **star:3814** Distributed pluggable RPC service framework like alibaba Dubbo.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [tendermint](https://github.com/tendermint/tendermint) **star:3156** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [torrent](https://github.com/anacrolix/torrent) **star:2854** BitTorrent client package.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [raft](https://github.com/hashicorp/raft) **star:2847** Golang implementation of the Raft consensus protocol, by HashiCorp.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [dragonboat](https://github.com/lni/dragonboat) **star:2544** A feature complete and high performance multi-group Raft library in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [glow](https://github.com/chrislusf/glow) **star:2532** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   ![star > 1000][Silver]
* [gleam](https://github.com/chrislusf/gleam) **star:2103** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   ![star > 1000][Silver]
* [emitter-io](https://github.com/emitter-io/emitter) **star:1937** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [KrakenD](https://github.com/devopsfaith/krakend) **star:1762** Ultra performant API Gateway framework with middlewares.   ![star > 1000][Silver]
* [hprose](https://github.com/hprose/hprose-golang) **star:1007** Very newbility RPC Library, support 25+ languages now.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [ringpop-go](https://github.com/uber/ringpop-go) **star:573** Scalable, fault-tolerant application-layer sharding for Go applications.   ![star > 100][Bronze]
* [gorpc](https://github.com/valyala/gorpc) **star:553** Simple, fast and scalable RPC library for high load.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-health](https://github.com/InVisionApp/go-health) **star:479** Library for enabling asynchronous dependency health checks in your service.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [digota](https://github.com/digota/digota) **star:301** grpc ecommerce microservice.   ![star > 100][Bronze]
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [sleuth](https://github.com/ursiform/sleuth) **star:300** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-jump](https://github.com/dgryski/go-jump) **star:254** Port of Google's "Jump" Consistent Hash function.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [consistent](https://github.com/buraksezer/consistent) **star:193** Consistent hashing with bounded loads.   ![star > 100][Bronze]
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [redis-lock](https://github.com/bsm/redis-lock) **star:147** Simplified distributed locking implementation using Redis.   ![star > 100][Bronze]   ![Archived][Archived]
* [dht](https://github.com/anacrolix/dht) **star:128** BitTorrent Kademlia DHT implementation.   ![star > 100][Bronze]
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:113** The jsonrpc package helps implement of JSON-RPC 2.0.   ![star > 100][Bronze]
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:102** JSON-RPC 2.0 HTTP client implementation.   ![star > 100][Bronze]
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:52** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
* [doublejump](https://github.com/edwingeng/doublejump) **star:40** A revamped Google's jump consistent hash.
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [drmaa](https://github.com/dgruber/drmaa) **star:25** Job submission library for cluster schedulers based on the DRMAA standard.   ![It hasn't been updated in the last year][Yellow]
* [flowgraph](https://github.com/vectaport/flowgraph) **star:18** flow-based programming package.
* [dynatomic](https://github.com/tylfin/dynatomic) **star:8** A library for using DynamoDB as an atomic counter.
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [outboxer](https://github.com/italolelis/outboxer) **star:5** Outboxer is a go library that implements the outbox pattern.

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:5180** Email and SMTP testing with web and API interface.   ![star > 5000][Gold]
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [hermes](https://github.com/matcornic/hermes) **star:1619** Golang package that generates clean, responsive HTML e-mails.   ![star > 1000][Silver]
* [email](https://github.com/jordan-wright/email) **star:1094** A robust and flexible email library for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-imap](https://github.com/emersion/go-imap) **star:734** IMAP library for clients and servers.   ![star > 100][Bronze]
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:521** SendGrid's Go library for sending email.   ![star > 100][Bronze]
* [Hectane](https://github.com/hectane/hectane) **star:167** Lightweight SMTP client providing an HTTP API.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [douceur](https://github.com/aymerick/douceur) **star:162** CSS inliner for your HTML emails.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-message](https://github.com/emersion/go-message) **star:113** Streaming library for the Internet Message Format and mail messages.   ![star > 100][Bronze]
* [smtp](https://github.com/mailhog/smtp) **star:51** SMTP server protocol state machine.   ![It hasn't been updated in the last year][Yellow]
* [go-dkim](https://github.com/toorop/go-dkim) **star:47** DKIM library, to sign & verify email.
* [go-premailer](https://github.com/vanng822/go-premailer) **star:35** Inline styling for HTML mail in Go.
* [Gomail](https://github.com/go-gomail/gomail/)  Gomail is a very simple and powerful package to send emails.

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) **star:4736** JavaScript interpreter written in Go.   ![star > 1000][Silver]
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:2966** Lua 5.1 VM and compiler written in Go.   ![star > 1000][Silver]
* [go-lua](https://github.com/Shopify/go-lua) **star:1671** Port of the Lua 5.2 VM to pure Go.   ![star > 1000][Silver]
* [tengo](https://github.com/d5/tengo) **star:1309** Bytecode compiled script language for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [anko](https://github.com/mattn/anko) **star:921** Scriptable interpreter written in Go.   ![star > 100][Bronze]
* [go-python](https://github.com/sbinet/go-python) **star:905** naive go bindings to the CPython C-API.   ![star > 100][Bronze]
* [expr](https://github.com/antonmedv/expr) **star:705** an engine that can evaluate expressions.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-php](https://github.com/deuill/go-php) **star:684** PHP bindings for Go.   ![star > 100][Bronze]
* [go-duktape](https://github.com/olebedev/go-duktape) **star:652** Duktape JavaScript engine bindings for Go.   ![star > 100][Bronze]
* [golua](https://github.com/aarzilli/golua) **star:440** Go bindings for Lua C API.   ![star > 100][Bronze]
* [gisp](https://github.com/jcla1/gisp) **star:428** Simple LISP in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [agora](https://github.com/PuerkitoBio/agora) **star:320** Dynamically typed, embeddable programming language in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [gval](https://github.com/PaesslerAG/gval) **star:139** A highly customizable expression language written in Go.   ![star > 100][Bronze]
* [binder](https://github.com/alexeyco/binder) **star:31** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   ![It hasn't been updated in the last year][Yellow]
* [gentee](https://github.com/gentee/gentee) **star:27** Embeddable scripting programming language.   ![There was an update last week][Green]
* [purl](https://github.com/ian-kent/purl) **star:27** Perl 5.18.2 embedded in Go.   ![It hasn't been updated in the last year][Yellow]
* [ngaro](https://github.com/db47h/ngaro) **star:19** Embeddable Ngaro VM implementation enabling scripting in Retro.   ![It hasn't been updated in the last year][Yellow]

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) **star:4908** Package that provides simple error handling primitives.   ![star > 1000][Silver]
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:731** Go (golang) package for representing a list of errors as a single error.   ![star > 100][Bronze]
* [errorx](https://github.com/joomcode/errorx) **star:556** A feature rich error package with stack traces, composition of errors and more.   ![star > 100][Bronze]
* [tracerr](https://github.com/ztrue/tracerr) **star:496** Golang errors with stack trace and source fragments.   ![star > 100][Bronze]
* [errlog](https://github.com/snwfdhmp/errlog) **star:181** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [werr](https://github.com/txgruppi/werr) **star:11** Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.   ![It hasn't been updated in the last year][Yellow]

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2242** FileSystem Abstraction System for Go.   ![star > 1000][Silver]
* [pdfcpu](https://github.com/hhrutter/pdfcpu) **star:949** PDF processor.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [notify](https://github.com/rjeczalik/notify) **star:497** File system event notification library with simple API, similar to os/signal.   ![star > 100][Bronze]
* [opc](https://github.com/qmuntal/opc) **star:57** Load Open Packaging Conventions (OPC) files for Go.
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:48** Load csv file using tag.
* [skywalker](https://github.com/dixonwille/skywalker) **star:48** Package to allow one to concurrently go through a filesystem with ease.   ![It hasn't been updated in the last year][Yellow]
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [tarfs](https://github.com/posener/tarfs) **star:35** Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.   ![It hasn't been updated in the last year][Yellow]
* [vfs](https://github.com/C2FO/vfs) **star:22** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   ![There was an update last week][Green]
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:15** Load gtfs files in go.
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:11** Copy files for humans.
* [flop](https://github.com/homedepot/flop) **star:9** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).
* [checksum](https://github.com/codingsince1985/checksum) **star:6** Compute message digest, like MD5 and SHA256, for large files.
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:3** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:1603** Arbitrary-precision fixed-point decimal numbers.   ![star > 1000][Silver]
* [go-money](https://github.com/rhymond/go-money) **star:622** Implementation of Fowler's Money pattern.   ![star > 100][Bronze]
* [go-finance](https://github.com/FlashBoys/go-finance) **star:536** Comprehensive financial markets data in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [accounting](https://github.com/leekchan/accounting) **star:488** money and currency formatting for golang.   ![star > 100][Bronze]
* [techan](https://github.com/sdcoffey/techan) **star:161** Technical analysis library with advanced market analysis and trading strategies.   ![star > 100][Bronze]
* [orderbook](https://github.com/i25959341/orderbook) **star:70** Matching Engine for Limit Order Book in Golang.
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:62** Query OFX servers and/or parse the responses (with example command-line client).
* [vat](https://github.com/dannyvankooten/vat) **star:60** VAT number validation & EU VAT rates.
* [transaction](https://github.com/claygod/transaction) **star:55** Embedded transactional database of accounts, running in multithreaded mode.
* [go-finance](https://github.com/alpeb/go-finance) **star:42** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   ![It hasn't been updated in the last year][Yellow]
* [currency](https://github.com/bnkamalesh/currency) **star:9** High performant & accurate currency computation package.

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) **star:971** CSRF protection middleware for Go.   ![star > 100][Bronze]
* [binding](https://github.com/mholt/binding) **star:753** Binds form and JSON data from net/http Request to struct.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:440** CSRF protection for Go web applications & services.   ![star > 100][Bronze]
* [form](https://github.com/go-playground/form) **star:349** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   ![star > 100][Bronze]
* [conform](https://github.com/leebenson/conform) **star:173** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [formam](https://github.com/monoculum/formam) **star:127** decode form's values into a struct.   ![star > 100][Bronze]
* [forms](https://github.com/albrow/forms) **star:105** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [bind](https://github.com/robfig/bind) **star:23** Bind form data to any Go values.   ![It hasn't been updated in the last year][Yellow]

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1067** Useful collection of helpfully functional Go collection utilities.   ![star > 1000][Silver]
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:107** Monad, Functional Programming features for Golang.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [fuego](https://github.com/seborama/fuego) **star:39** Functional Experiment in Go.

## Game Development

*Awesome game development libraries.*

* [Leaf](https://github.com/name5566/leaf) **star:3064** Lightweight game server framework.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [Pixel](https://github.com/faiface/pixel) **star:2445** Hand-crafted 2D game library in Go.   ![star > 1000][Silver]
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:1857** dead simple 2D game library in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [goworld](https://github.com/xiaonanln/goworld) **star:1198** Scalable game server engine, featuring space-entity framework and hot-swapping.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1158** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).   ![star > 1000][Silver]   ![There was an update last week][Green]
* [engo](https://github.com/EngoEngine/engo) **star:1083** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   ![star > 1000][Silver]
* [gonet](https://github.com/xtaci/gonet) **star:1050** Game server skeleton implemented with golang.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [termloop](https://github.com/JoelOtter/termloop) **star:1025** Terminal-based game engine for Go, built on top of Termbox.   ![star > 1000][Silver]
* [nano](https://github.com/lonng/nano) **star:1000** Lightweight, facility, high performance golang based game server framework.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [g3n](https://github.com/g3n/engine) **star:757** Go 3D Game Engine.   ![star > 100][Bronze]
* [Oak](https://github.com/oakmound/oak) **star:628** Pure Go game engine.   ![star > 100][Bronze]
* [Azul3D](https://github.com/azul3d/engine) **star:426** 3D game engine written in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:382** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.   ![star > 100][Bronze]
* [go-astar](https://github.com/beefsack/go-astar) **star:326** Go implementation of the A\* path finding algorithm.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:312** 2d game engine written in Go working on OpenGL.   ![star > 100][Bronze]
* [Pitaya](https://github.com/topfreegames/pitaya) **star:300** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go3d](https://github.com/ungerik/go3d) **star:164** Performance oriented 2D/3D math package for Go.   ![star > 100][Bronze]
* [glop](https://github.com/runningwild/glop) **star:76** Glop (Game Library Of Power) is a fairly simple cross-platform game library.   ![It hasn't been updated in the last year][Yellow]
* [go-collada](https://github.com/GlenKelley/go-collada) **star:12** Go package for working with the Collada file format.   ![It hasn't been updated in the last year][Yellow]

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1798** .NET LINQ-like query methods for Go.   ![star > 1000][Silver]
* [jennifer](https://github.com/dave/jennifer) **star:1277** Generate arbitrary Go code without templates.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gen](https://github.com/clipperhouse/gen) **star:1037** Code generation tool for ‘generics’-like functionality.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [goderive](https://github.com/awalterschulze/goderive) **star:747** Derives functions from input types.   ![star > 100][Bronze]
* [GoWrap](https://github.com/hexdigest/gowrap) **star:266** Generate decorators for Go interfaces using simple templates.   ![star > 100][Bronze]
* [interfaces](https://github.com/rjeczalik/interfaces) **star:188** Command line tool for generating interface definitions.   ![star > 100][Bronze]
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:87** Go preprocessor for package scoped reflection.   ![It hasn't been updated in the last year][Yellow]
* [go-enum](https://github.com/abice/go-enum) **star:86** Code generation for enums from code comments.
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:43** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   ![It hasn't been updated in the last year][Yellow]
* [gotype](https://github.com/wzshiming/gotype) **star:23** Golang source code parsing, usage like reflect package.   ![Contains Chinese documents][CN]
* [generis](https://github.com/senselogic/GENERIS) **star:18** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) **star:6339** Geolocation DB with spatial index and realtime geofencing.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [S2 geometry](https://github.com/golang/geo) **star:887** S2 geometry library in Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [geocache](https://github.com/melihmucuk/geocache) **star:111** In-memory cache that is suitable for geolocation based applications.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [osm](https://github.com/paulmach/osm) **star:68** Library for reading, writing and working with OpenStreetMap data and APIs.
* [geoserver](https://github.com/hishamkaram/geoserver) **star:25** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
* [gismanager](https://github.com/hishamkaram/gismanager) **star:19** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.
* [pbf](https://github.com/maguro/pbf) **star:16** OpenStreetMap PBF golang encoder/decoder.

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:8546** Compiler from Go to JavaScript.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [llgo](https://github.com/go-llvm/llgo) **star:989** LLVM-based compiler for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:393** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [c4go](https://github.com/Konstantin8105/c4go) **star:157** Transpile C code to Go code.   ![star > 100][Bronze]
* [f4go](https://github.com/Konstantin8105/f4go) **star:11** Transpile FORTRAN 77 code to Go code.

## Goroutines

*Tools for managing and working with Goroutines.*

* [goworker](https://github.com/benmanns/goworker) **star:2253** goworker is a Go-based background worker.   ![star > 1000][Silver]
* [ants](https://github.com/panjf2000/ants) **star:1895** A high-performance goroutine pool for golang.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [tunny](https://github.com/Jeffail/tunny) **star:1344** Goroutine pool for golang.   ![star > 1000][Silver]
* [grpool](https://github.com/ivpusic/grpool) **star:499** Lightweight Goroutine pool.   ![star > 100][Bronze]
* [pool](https://github.com/go-playground/pool) **star:482** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-floc](https://github.com/workanator/go-floc) **star:167** Orchestrate goroutines with ease.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [workerpool](https://github.com/gammazero/workerpool) **star:138** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   ![star > 100][Bronze]
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:104** Control goroutines execution order.   ![star > 100][Bronze]
* [semaphore](https://github.com/kamilsk/semaphore) **star:75** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   ![There was an update last week][Green]
* [GoSlaves](https://github.com/themester/GoSlaves) **star:75** Simple and Asynchronous Goroutine pool library.
* [semaphore](https://github.com/marusama/semaphore) **star:72** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:56** manages a resizeable pool of context-aware goroutines to bound concurrency.
* [worker-pool](https://github.com/vardius/worker-pool) **star:45** goworker is a Go simple async worker pool.
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:36** CyclicBarrier for golang.
* [breaker](https://github.com/kamilsk/breaker) **star:34** Flexible mechanism to make execution flow interruptible.   ![There was an update last week][Green]
* [gollback](https://github.com/vardius/gollback) **star:27** asynchronous simple function utilities, for managing execution of closures and callbacks.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:25** Run functions in parallel.   ![It hasn't been updated in the last year][Yellow]
* [async](https://github.com/studiosol/async) **star:23** A safe way to execute functions asynchronously, recovering them in case of panic.
* [threadpool](https://github.com/shettyh/threadpool) **star:19** Golang threadpool implementation.
* [Hunch](https://github.com/AaronJan/Hunch) **star:13** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [artifex](https://github.com/borderstech/artifex) **star:12** Simple in-memory job queue for Golang using worker-based dispatching.
* [stl](https://github.com/ssgreg/stl) **star:8** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.
* [go-trylock](https://github.com/subchen/go-trylock) **star:4** TryLock support on read-write lock for Golang.   ![It hasn't been updated in the last year][Yellow]
* [routine](https://github.com/x-mod/routine) **star:3** go routine control with context, support: Main, Go, Pool and some useful Executors.
* [queue](https://github.com/AnikHasibul/queue) **star:2** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [ui](https://github.com/andlabs/ui) **star:6966** Platform-native GUI library for Go. Cross platform.   ![star > 5000][Gold]
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [fyne](https://github.com/fyne-io/fyne) **star:6304** Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [qt](https://github.com/therecipe/qt) **star:6074** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   ![star > 5000][Gold]   ![There was an update last week][Green]
* [webview](https://github.com/zserge/webview) **star:4669** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).   ![star > 1000][Silver]
* [walk](https://github.com/lxn/walk) **star:3698** Windows application library kit for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [app](https://github.com/murlokswarm/app) **star:2958** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   ![star > 1000][Silver]
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:2674** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1449** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   ![star > 1000][Silver]
* [gotk3](https://github.com/gotk3/gotk3) **star:773** Go bindings for GTK3.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [gowd](https://github.com/dtylman/gowd) **star:211** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   ![star > 100][Bronze]

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4435** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [systray](https://github.com/getlantern/systray) **star:789** Cross platform Go library to place an icon and menu in the notification area.   ![star > 100][Bronze]
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:493** OSX Desktop Notifications library for Go.   ![star > 100][Bronze]
* [trayhost](https://github.com/shurcooL/trayhost) **star:160** Cross-platform Go library to place an icon in the host operating system's taskbar.   ![star > 100][Bronze]
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:1** Go bindings for libappindicator3 C library.
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:1** OSX library to notify about any (pluggable) activity on your machine.
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier)  OSX Sleep/Wake notifications in golang.


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [imaginary](https://github.com/h2non/imaginary) **star:2593** Fast and simple HTTP microservice for image resizing.   ![star > 1000][Silver]
* [bild](https://github.com/anthonynsimon/bild) **star:2548** Collection of image processing algorithms in pure Go.   ![star > 1000][Silver]
* [imaging](https://github.com/disintegration/imaging) **star:2534** Simple Go image processing package.   ![star > 1000][Silver]
* [ln](https://github.com/fogleman/ln) **star:2462** 3D line art rendering in Go.   ![star > 1000][Silver]
* [gocv](https://github.com/hybridgroup/gocv) **star:2461** Go package for computer vision using OpenCV 3.3+.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [resize](https://github.com/nfnt/resize) **star:2130** Image resizing for Go with common interpolation methods.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [gg](https://github.com/fogleman/gg) **star:1918** 2D rendering in pure Go.   ![star > 1000][Silver]
* [pt](https://github.com/fogleman/pt) **star:1774** Path tracing engine written in Go.   ![star > 1000][Silver]
* [svgo](https://github.com/ajstarks/svgo) **star:1333** Go Language Library for SVG generation.   ![star > 1000][Silver]
* [smartcrop](https://github.com/muesli/smartcrop) **star:1261** Finds good crops for arbitrary images and crop sizes.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gift](https://github.com/disintegration/gift) **star:1210** Package of image processing filters.   ![star > 1000][Silver]
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1093** Go bindings for OpenCV.   ![star > 1000][Silver]
* [picfit](https://github.com/thoas/picfit) **star:1069** An image resizing server written in Go.   ![star > 1000][Silver]
* [geopattern](https://github.com/pravj/geopattern) **star:1011** Create beautiful generative image patterns from a string.   ![star > 1000][Silver]
* [imagick](https://github.com/gographics/imagick) **star:976** Go binding to ImageMagick's MagickWand C API.   ![star > 100][Bronze]
* [bimg](https://github.com/h2non/bimg) **star:797** Small package for fast and efficient image processing using libvips.   ![star > 100][Bronze]
* [stegify](https://github.com/DimitarPetrov/stegify) **star:506** Go tool for LSB steganography, capable of hiding any file within an image.   ![star > 100][Bronze]
* [mort](https://github.com/aldor007/mort) **star:365** Storage and image processing server written in Go.   ![star > 100][Bronze]
* [govatar](https://github.com/o1egl/govatar) **star:310** Library and CMD tool for generating funny avatars.   ![star > 100][Bronze]
* [image2ascii](https://github.com/qeesung/image2ascii) **star:290** Convert image to ASCII.   ![star > 100][Bronze]
* [go-nude](https://github.com/koyachi/go-nude) **star:285** Nudity detection with Go.   ![star > 100][Bronze]
* [goimagehash](https://github.com/corona10/goimagehash) **star:215** Go Perceptual image hashing package.   ![star > 100][Bronze]
* [rez](https://github.com/bamiaux/rez) **star:188** Image resizing in pure Go and SIMD.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [img](https://github.com/hawx/img) **star:129** Selection of image manipulation tools.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-cairo](https://github.com/ungerik/go-cairo) **star:85** Go binding for the cairo graphics library.
* [mergi](https://github.com/noelyahan/mergi) **star:71** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
* [darkroom](https://github.com/gojek/darkroom) **star:57** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   ![There was an update last week][Green]
* [go-gd](https://github.com/bolknote/go-gd) **star:50** Go binding for GD library.   ![It hasn't been updated in the last year][Yellow]
* [gltf](https://github.com/qmuntal/gltf) **star:37** Efficient and robust glTF 2.0 reader, writer and validator.
* [cameron](https://github.com/aofei/cameron) **star:31** An avatar generator for Go.
* [goimghdr](https://github.com/corona10/goimghdr) **star:26** The imghdr module determines the type of image contained in a file for Go.
* [steganography](https://github.com/auyer/steganography) **star:26** Pure Go Library for LSB steganography.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Port of webcolors library from Python to Go.   ![It hasn't been updated in the last year][Yellow]
* [tga](https://github.com/ftrvxmtrx/tga) **star:24** Package tga is a TARGA image format decoder/encoder.   ![It hasn't been updated in the last year][Yellow]
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1130** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.   ![star > 1000][Silver]
* [gatt](https://github.com/paypal/gatt) **star:814** Gatt is a Go package for building Bluetooth Low Energy peripherals.   ![star > 100][Bronze]
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [mainflux](https://github.com/Mainflux/mainflux) **star:592** Industrial IoT Messaging and Device Management Server.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [devices](https://github.com/goiot/devices) **star:225** Suite of libraries for IoT devices, experimental for x/exp/io.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:180** Lightweight stream processing engine for IoT.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [connectordb](https://github.com/connectordb/connectordb) **star:168** Open-Source Platform for Quantified Self & IoT.   ![star > 100][Bronze]
* [huego](https://github.com/amimof/huego) **star:110** An extensive Philips Hue client library for Go.   ![star > 100][Bronze]
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [eywa](https://github.com/xcodersun/eywa) **star:36** Project Eywa is essentially a connection manager that keeps track of connected devices.   ![It hasn't been updated in the last year][Yellow]

## Job Scheduler

*Libraries for scheduling jobs.*

* [gron](https://github.com/roylee0704/gron) **star:635** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   ![star > 100][Bronze]
* [JobRunner](https://github.com/bamzi/jobrunner) **star:569** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [jobs](https://github.com/albrow/jobs) **star:450** Persistent and flexible background jobs library.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [scheduler](https://github.com/carlescere/scheduler) **star:294** Cronjobs scheduling made easy.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [go-cron](https://github.com/rk/go-cron) **star:176** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [clockwork](https://github.com/whiteShtef/clockwork) **star:77** Simple and intuitive job scheduling library in Go.
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:36** Job scheduler that supports webhooks, crons and classic scheduling.

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) **star:4914** Get a JSON value with one line of code.   ![star > 1000][Silver]
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2031** Automatically generate Go (golang) struct definitions from example JSON.   ![star > 1000][Silver]
* [gojq](https://github.com/elgs/gojq) **star:140** JSON query in Golang.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [kazaam](https://github.com/Qntfy/kazaam) **star:133** API for arbitrary transformation of JSON documents.   ![star > 100][Bronze]
* [jsongo](https://github.com/ricardolonga/jsongo) **star:92** Fluent API to make it easier to create Json objects.   ![It hasn't been updated in the last year][Yellow]
* [gjo](https://github.com/skanehira/gjo) **star:62** Small utility to create JSON objects.
* [jsonf](https://github.com/miolini/jsonf) **star:55** Console tool for highlighted formatting and struct query fetching JSON.   ![It hasn't been updated in the last year][Yellow]
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:40** JSON diff utility written in Go.
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [mp](https://github.com/sanbornm/mp) **star:33** Simple cli email parser. It currently takes stdin and outputs JSON.   ![It hasn't been updated in the last year][Yellow]
* [go-respond](https://github.com/nicklaw5/go-respond) **star:23** Go package for handling common HTTP JSON responses.
* [ajson](https://github.com/spyzhov/ajson) **star:13** Abstract JSON for golang with JSONPath support.
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:7** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:5** Go bindings based on the JSON API errors reference.   ![It hasn't been updated in the last year][Yellow]

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) **star:11980** Structured logger for Go.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [zap](https://github.com/uber-go/zap) **star:7469** Fast, structured, leveled logging in Go.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [spew](https://github.com/davecgh/go-spew) **star:3310** Implements a deep pretty printer for Go data structures to aid in debugging.   ![star > 1000][Silver]
* [glog](https://github.com/golang/glog) **star:2301** Leveled execution logs for Go.   ![star > 1000][Silver]
* [zerolog](https://github.com/rs/zerolog) **star:2229** Zero-allocation JSON logger.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [tail](https://github.com/hpcloud/tail) **star:1541** Go package striving to emulate the features of the BSD tail program.   ![star > 1000][Silver]
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1450** Simple rolling logger, implements io.WriteCloser.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [seelog](https://github.com/cihub/seelog) **star:1360** Logging functionality with flexible dispatching, filtering, and formatting.   ![star > 1000][Silver]
* [log15](https://github.com/inconshreveable/log15) **star:911** Simple, powerful logging for Go.   ![star > 100][Bronze]
* [log](https://github.com/apex/log) **star:733** Structured logging package for Go.   ![star > 100][Bronze]
* [logxi](https://github.com/mgutz/logxi) **star:333** 12-factor app logger that is fast and makes you happy.   ![star > 100][Bronze]
* [onelog](https://github.com/francoispqt/onelog) **star:331** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.   ![star > 100][Bronze]
* [log](https://github.com/go-playground/log) **star:267** Simple, configurable and scalable Structured Logging for Go.   ![star > 100][Bronze]
* [logutils](https://github.com/hashicorp/logutils) **star:247** Utilities for slightly better logging in Go (Golang) extending the standard logger.   ![star > 100][Bronze]
* [go-logger](https://github.com/apsdehal/go-logger) **star:233** Simple logger of Go Programs, with level handlers.   ![star > 100][Bronze]
* [logger](https://github.com/azer/logger) **star:134** Minimalistic logging library for Go.   ![star > 100][Bronze]
* [xlog](https://github.com/rs/xlog) **star:128** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   ![star > 100][Bronze]
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:109** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:97** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.
* [log-voyage](https://github.com/firstrow/logvoyage) **star:82** Full-featured logging saas written in golang.   ![It hasn't been updated in the last year][Yellow]
* [glg](https://github.com/kpango/glg) **star:51** glg is simple and fast leveled logging library for Go.
* [stdlog](https://github.com/alexcesaro/log) **star:43** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   ![It hasn't been updated in the last year][Yellow]
* [gologger](https://github.com/sadlil/gologger) **star:38** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j implementation in Go.   ![It hasn't been updated in the last year][Yellow]
* [logex](https://github.com/chzyer/logex) **star:32** Golang log lib, supports tracking and level, wrap by standard log lib.   ![It hasn't been updated in the last year][Yellow]
* [logrusly](https://github.com/sebest/logrusly) **star:26** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   ![It hasn't been updated in the last year][Yellow]
* [go-log](https://github.com/siddontang/go-log) **star:23** Log lib supports level and multi handlers.
* [log](https://github.com/teris-io/log) **star:22** Structured log interface for Go cleanly separates logging facade from its implementation.   ![It hasn't been updated in the last year][Yellow]
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   ![It hasn't been updated in the last year][Yellow]
* [journald](https://github.com/ssgreg/journald) **star:19** Go implementation of systemd Journal's native API for logging.
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:19** Simple writer that rotate log files automatically based on current date and time, like cronolog.
* [distillog](https://github.com/amoghe/distillog) **star:19** distilled levelled logging (think of it as stdlib + log levels).   ![It hasn't been updated in the last year][Yellow]
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [gomol](https://github.com/aphistic/gomol) **star:15** Multiple-output, structured logging for Go with extensible logging outputs.
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   ![It hasn't been updated in the last year][Yellow]
* [go-log](https://github.com/subchen/go-log) **star:9** Simple and configurable Logging in Go, with level, formatters and writers.   ![It hasn't been updated in the last year][Yellow]
* [glo](https://github.com/lajosbencz/glo) **star:8** PHP Monolog inspired logging facility with identical severity levels.
* [xlog](https://github.com/xfxdev/xlog) **star:7** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.
* [logo](https://github.com/mbndr/logo) **star:5** Golang logger to different configurable writers.   ![It hasn't been updated in the last year][Yellow]
* [log](https://github.com/aerogo/log) **star:4** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:6650** General Machine Learning library for Go.   ![star > 5000][Gold]   ![Contains Chinese documents][CN]
* [gorgonia](https://github.com/chewxy/gorgonia) **star:2697** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   ![star > 1000][Silver]
* [tfgo](https://github.com/galeone/tfgo) **star:1190** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   ![star > 1000][Silver]
* [goml](https://github.com/cdipaolo/goml) **star:1010** On-line Machine Learning in Go.   ![star > 1000][Silver]
* [gosseract](https://github.com/otiai10/gosseract) **star:876** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:643** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   ![star > 100][Bronze]
* [bayesian](https://github.com/jbrukh/bayesian) **star:629** Naive Bayesian Classification for Golang.   ![star > 100][Bronze]
* [eaopt](https://github.com/MaxHalford/eaopt) **star:623** An evolutionary optimization library.   ![star > 100][Bronze]
* [gorse](https://github.com/zhenghaoz/gorse) **star:537** A High Performance Recommender System Package based on Collaborative Filtering for Go.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [gobrain](https://github.com/goml/gobrain) **star:386** Neural Networks written in go.   ![star > 100][Bronze]
* [regommend](https://github.com/muesli/regommend) **star:249** Recommendation & collaborative filtering engine.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [ocrserver](https://github.com/otiai10/ocrserver) **star:225** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   ![star > 100][Bronze]
* [go-deep](https://github.com/patrikeh/go-deep) **star:219** A feature-rich neural network library in Go.   ![star > 100][Bronze]
* [go-galib](https://github.com/thoj/go-galib) **star:171** Genetic Algorithms library written in Go / golang.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:148** Go Interface to Open Neural Network Exchange (ONNX).   ![star > 100][Bronze]   ![There was an update last week][Green]
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:144** Recommendation Algorithms library written in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [shield](https://github.com/eaigner/shield) **star:123** Bayesian text classifier with flexible tokenizers and storage backends for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-fann](https://github.com/white-pony/go-fann) **star:99** Go bindings for Fast Artificial Neural Networks(FANN) library.   ![It hasn't been updated in the last year][Yellow]
* [Goptuna](https://github.com/c-bata/goptuna) **star:78** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   ![There was an update last week][Green]
* [goga](https://github.com/tomcraven/goga) **star:78** Genetic algorithm library for Go.   ![It hasn't been updated in the last year][Yellow]
* [libsvm](https://github.com/datastream/libsvm) **star:63** libsvm golang version derived work based on LIBSVM 3.14.   ![It hasn't been updated in the last year][Yellow]
* [neural-go](https://github.com/schuyler/neural-go) **star:60** Multilayer perceptron network implemented in Go, with training via backpropagation.   ![It hasn't been updated in the last year][Yellow]
* [go-pr](https://github.com/daviddengcn/go-pr) **star:56** Pattern recognition package in Go lang.   ![It hasn't been updated in the last year][Yellow]
* [neat](https://github.com/jinyeom/neat) **star:55** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.
* [goscore](https://github.com/asafschers/goscore) **star:37** Go Scoring API for PMML.
* [fonet](https://github.com/Fontinalis/fonet) **star:33** A Deep Neural Network library written in Go.   ![There was an update last week][Green]
* [godist](https://github.com/e-dard/godist) **star:24** Various probability distributions, and associated methods.   ![It hasn't been updated in the last year][Yellow]
* [Varis](https://github.com/Xamber/Varis) **star:23** Golang Neural Network.   ![It hasn't been updated in the last year][Yellow]
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:21** Go implementation of the k-modes and k-prototypes clustering algorithms.   ![It hasn't been updated in the last year][Yellow]
* [probab](https://github.com/ThePaw/probab) **star:10** Probability distribution functions. Bayesian inference. Written in pure Go.   ![It hasn't been updated in the last year][Yellow]
* [evoli](https://github.com/khezen/evoli) **star:8** Genetic Algorithm and Particle Swarm Optimization library.
* [GoMind](https://github.com/surenderthakran/gomind) **star:6** A simplistic Neural Network Library in Go.   ![It hasn't been updated in the last year][Yellow]

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) **star:4644** Go library for Apache Kafka.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gorush](https://github.com/appleboy/gorush) **star:3717** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   ![star > 1000][Silver]
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:3687** Real-time messaging (Websockets or SockJS) server in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [machinery](https://github.com/RichardKnop/machinery) **star:3370** Asynchronous task queue/job queue based on distributed message passing.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:2888** socket.io library for golang, a realtime application framework.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [NATS Go Client](https://github.com/nats-io/nats) **star:2401** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [APNs2](https://github.com/sideshow/apns2) **star:2043** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Benthos](https://github.com/Jeffail/benthos) **star:2014** A message streaming bridge between a range of protocols.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1837** gopush-cluster is a go push server cluster.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [melody](https://github.com/olahol/melody) **star:1565** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   ![star > 1000][Silver]
* [mangos](https://github.com/go-mangos/mangos) **star:1534** Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.   ![star > 1000][Silver]
* [Mercure](https://github.com/dunglas/mercure) **star:1511** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1458** the official Go package for NSQ.   ![star > 1000][Silver]
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1096** Redis backed unified push service for server-side notifications to mobile devices.   ![star > 1000][Silver]
* [zmq4](https://github.com/pebbe/zmq4) **star:774** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   ![star > 100][Bronze]
* [Gollum](https://github.com/trivago/gollum) **star:767** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   ![star > 100][Bronze]
* [Beaver](https://github.com/Clivern/Beaver) **star:727** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   ![star > 100][Bronze]
* [EventBus](https://github.com/asaskevich/EventBus) **star:561** The lightweight event bus with async compatibility.   ![star > 100][Bronze]
* [golongpoll](https://github.com/jcuga/golongpoll) **star:428** HTTP longpoll server library that makes web pub-sub simple.   ![star > 100][Bronze]
* [dbus](https://github.com/godbus/dbus) **star:358** Native Go bindings for D-Bus.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Glue](https://github.com/desertbit/glue) **star:320** Robust Go and Javascript Socket Library (Alternative to Socket.io).   ![star > 100][Bronze]
* [emitter](https://github.com/olebedev/emitter) **star:309** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   ![star > 100][Bronze]
* [pubsub](https://github.com/tuxychandru/pubsub) **star:278** Simple pubsub package for go.   ![star > 100][Bronze]
* [guble](https://github.com/smancke/guble) **star:138** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Bus](https://github.com/mustafaturan/bus) **star:115** Minimalist message bus implementation for internal communication.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [oplog](https://github.com/dailymotion/oplog) **star:94** Generic oplog/replication system for REST APIs.   ![It hasn't been updated in the last year][Yellow]
* [rabtap](https://github.com/jandelgado/rabtap) **star:72** RabbitMQ swiss army knife cli app.   ![There was an update last week][Green]
* [messagebus](https://github.com/vardius/message-bus) **star:65** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:62** A tiny wrapper over amqp exchanges and queues.
* [drone-line](https://github.com/appleboy/drone-line) **star:60** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:55** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   ![It hasn't been updated in the last year][Yellow]
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:51** A tiny wrapper around NSQ topic and channel.   ![It hasn't been updated in the last year][Yellow]
* [go-notify](https://github.com/TheCreeper/go-notify) **star:47** Native implementation of the freedesktop notification spec.
* [goose](https://github.com/ian-kent/goose) **star:36** Server Sent Events in Go.   ![It hasn't been updated in the last year][Yellow]
* [event](https://github.com/agoalofalife/event) **star:27** Implementation of the pattern observer.   ![It hasn't been updated in the last year][Yellow]
* [hub](https://github.com/leandro-lugaresi/hub) **star:25** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   ![It hasn't been updated in the last year][Yellow]
* [Commander](https://github.com/jeroenrinzema/commander) **star:21** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   ![There was an update last week][Green]
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:11** Client library to Viessmann Vitotrol web service.
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** Gaurun Client written in Go.
* [jazz](https://github.com/socifi/jazz) **star:6** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:1** redisqueue provides a producer and consumer of a queue that uses Redis streams.
* [rmqconn](https://github.com/sbabiv/rmqconn)  RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:1729** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   ![star > 1000][Silver]   ![There was an update last week][Green]

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:4489** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [xlsx](https://github.com/tealeg/xlsx) **star:3433** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [xlsx](https://github.com/plandem/xlsx) **star:73** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.
* [go-excel](https://github.com/szyhf/go-excel) **star:46** A simple and light reader to read a relate-db-like excel as a table.
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:12** Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.   ![It hasn't been updated in the last year][Yellow]

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:915** A reflection based dependency injection toolkit for Go.   ![star > 100][Bronze]
* [fx](https://github.com/uber-go/fx) **star:753** A dependency injection based application framework for Go (built on top of dig).   ![star > 100][Bronze]   ![There was an update last week][Green]
* [alice](https://github.com/magic003/alice) **star:34** Additive dependency injection container for Golang.   ![It hasn't been updated in the last year][Yellow]
* [inject](https://github.com/defval/inject) **star:27** A reflection based dependency injection container with simple interface.
* [wire](https://github.com/Fs02/wire) **star:19** Strict Runtime Dependency Injection for Golang.
* [gocontainer](https://github.com/vardius/gocontainer) **star:9** Simple Dependency Injection Container.
* [linker](https://github.com/logrange/linker) **star:8** A reflection based dependency injection and inversion of control library with components lifecycle support.

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:9192** Set of common historical and emerging project layout patterns in the Go ecosystem.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [scaffold](https://github.com/catchplay/scaffold) **star:26** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.
* [go-sample](https://github.com/zitryss/go-sample) **star:24** A sample layout for Go application projects with the real code.

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:619** Collection of useful string functions ported from other languages.   ![star > 100][Bronze]
* [strutil](https://github.com/ozgio/strutil) **star:63** String utilities.

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:3936** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   ![star > 1000][Silver]   ![There was an update last week][Green]
* [archiver](https://github.com/mholt/archiver) **star:2491** Library and command for making and extracting .zip and .tar.gz archives.   ![star > 1000][Silver]
* [gosms](https://github.com/haxpax/gosms) **star:1225** Your own local SMS gateway in Go that can be used to send SMS.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:857** Resiliency patterns for golang.   ![star > 100][Bronze]
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:672** Generic object pool for Golang.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:629** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   ![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [shortid](https://github.com/teris-io/shortid) **star:451** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:423** Random data generator written in go.   ![star > 100][Bronze]
* [llvm](https://github.com/llir/llvm) **star:412** Library for interacting with LLVM IR in pure Go.   ![star > 100][Bronze]
* [health](https://github.com/dimiro1/health) **star:362** Easy to use, extensible health check library.   ![star > 100][Bronze]
* [conv](https://github.com/cstockton/go-conv) **star:342** Package conv provides fast and intuitive conversions across Go types.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [banner](https://github.com/dimiro1/banner) **star:233** Add beautiful banners into your Go applications.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gountries](https://github.com/pariz/gountries) **star:210** Package that exposes country and subdivision data.   ![star > 100][Bronze]
* [antch](https://github.com/antchfx/antch) **star:142** A fast, powerful and extensible web crawling & scraping framework.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [battery](https://github.com/distatus/battery) **star:135** Cross-platform, normalized battery information library.   ![star > 100][Bronze]
* [ffmt](https://github.com/go-ffmt/ffmt) **star:127** Beautify data display for Humans.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [stats](https://github.com/go-playground/stats) **star:121** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [lk](https://github.com/hyperboloide/lk) **star:121** A simple licensing library for golang.   ![star > 100][Bronze]
* [bitio](https://github.com/icza/bitio) **star:97** Highly optimized bit-level Reader and Writer for Go.   ![It hasn't been updated in the last year][Yellow]
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:83** An opinionated and concurrent health-check HTTP handler for RESTful services.
* [turtle](https://github.com/hackebrot/turtle) **star:76** Emojis for Go.   ![It hasn't been updated in the last year][Yellow]
* [gommit](https://github.com/antham/gommit) **star:75** Analyze git commit messages to ensure they follow defined patterns.   ![There was an update last week][Green]
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:68** Decompression library for RAR, TAR, ZIP and 7z archives.
* [indigo](https://github.com/osamingo/indigo) **star:51** Distributed unique ID generator of using Sonyflake and encoded by Base58.
* [morse](https://github.com/alwindoss/morse) **star:50** Library to convert to and from morse code.
* [captcha](https://github.com/steambap/captcha) **star:43** Package captcha provides an easy to use, unopinionated API for captcha generation.
* [xkg](https://github.com/go-xkg/xkg) **star:40** X Keyboard Grabber.   ![It hasn't been updated in the last year][Yellow]
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:34** HTTP service to generate PDF from Json requests.   ![It hasn't been updated in the last year][Yellow]
* [persian](https://github.com/mavihq/persian) **star:32** Some utilities for Persian language in go.   ![It hasn't been updated in the last year][Yellow]
* [ghorg](https://github.com/gabrie30/ghorg) **star:31** Clone all repos from a GitHub org into a single directory.   ![There was an update last week][Green]
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:30** GoLang Library for [Browser Capabilities Project](http://browscap.org/).
* [datacounter](https://github.com/miolini/datacounter) **star:28** Go counters for readers/writer/http.ResponseWriter.
* [autoflags](https://github.com/artyom/autoflags) **star:24** Go package to automatically define command line flags from struct fields.
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg) Specs implemented in Go.
* [gosh](https://github.com/osamingo/gosh) **star:17** Provide Go Statistics Handler, Struct, Measure Method.
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [url-shortener](https://github.com/pantrif/url-shortener) **star:17** A modern, powerful, and robust URL shortener microservice with mysql support.   ![It hasn't been updated in the last year][Yellow]
* [gotoprom](https://github.com/cabify/gotoprom) **star:14** Type-safe metrics builder wrapper library for the official Prometheus client.
* [sandid](https://github.com/aofei/sandid) **star:12** Every grain of sand on earth has its own ID.
* [anagent](https://github.com/mudler/anagent) **star:11** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   ![It hasn't been updated in the last year][Yellow]
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:9** Calculate average score and rating based on Wilson Score Equation.   ![It hasn't been updated in the last year][Yellow]
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.
* [shellwords](https://github.com/Wing924/shellwords) **star:7** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   ![It hasn't been updated in the last year][Yellow]
* [metrics](https://github.com/pascaldekloe/metrics) **star:4** Library for metrics instrumentation and Prometheus exposition.
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2047** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more.   ![star > 1000][Silver]
* [gse](https://github.com/go-ego/gse) **star:1075** Go efficient text segmentation; support english, chinese, japanese and other.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [when](https://github.com/olebedev/when) **star:927** Natural EN and RU language date/time parser with pluggable rules.   ![star > 100][Bronze]
* [gojieba](https://github.com/yanyiwu/gojieba) **star:820** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:523** CN Hanzi to Hanyu Pinyin converter.   ![star > 100][Bronze]
* [kagome](https://github.com/ikawaha/kagome) **star:415** JP morphological analyzer written in pure Go.   ![star > 100][Bronze]
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:352** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   ![star > 100][Bronze]
* [nlp](https://github.com/Shixzie/nlp) **star:352** Extract values from strings and fill your structs with nlp.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [sentences](https://github.com/neurosnap/sentences) **star:261** Sentence tokenizer:  converts text into a list of sentences.   ![star > 100][Bronze]
* [nlp](https://github.com/james-bowman/nlp) **star:216** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   ![star > 100][Bronze]
* [go-nlp](https://github.com/nuance/go-nlp) **star:79** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   ![It hasn't been updated in the last year][Yellow]
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [getlang](https://github.com/rylans/getlang) **star:74** Fast natural language detection package.
* [gounidecode](https://github.com/fiam/gounidecode) **star:67** Unicode transliterator (also known as unidecode) for Go.   ![It hasn't been updated in the last year][Yellow]
* [textcat](https://github.com/pebbe/textcat) **star:60** Go package for n-gram based text categorization, with support for utf-8 and raw text.   ![It hasn't been updated in the last year][Yellow]
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:58** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   ![It hasn't been updated in the last year][Yellow]
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:56** ASCII transliterations of Unicode text.
* [go-stem](https://github.com/agonopol/go-stem) **star:52** Implementation of the porter stemming algorithm.   ![It hasn't been updated in the last year][Yellow]
* [stemmer](https://github.com/dchest/stemmer) **star:47** Stemmer packages for Go programming language. Includes English and German stemmers.   ![It hasn't been updated in the last year][Yellow]
* [segment](https://github.com/blevesearch/segment) **star:46** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   ![It hasn't been updated in the last year][Yellow]
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:45** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
* [porter2](https://github.com/zhenjl/porter2) **star:33** Really fast Porter 2 stemmer.   ![It hasn't been updated in the last year][Yellow]
* [go2vec](https://github.com/danieldk/go2vec) **star:30** Reader and utility functions for word2vec embeddings.
* [paicehusk](https://github.com/rookii/paicehusk) **star:25** Golang implementation of the Paice/Husk Stemming Algorithm.   ![It hasn't been updated in the last year][Yellow]
* [snowball](https://github.com/goodsign/snowball) **star:24** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   ![It hasn't been updated in the last year][Yellow]
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo bindings to Yandex.Mystem - russian morphology analyzer.   ![It hasn't been updated in the last year][Yellow]
* [petrovich](https://github.com/striker2000/petrovich) **star:22** Petrovich is the library which inflects Russian names to given grammatical case.
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   ![It hasn't been updated in the last year][Yellow]
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:15** Go bindings for the snowball libstemmer library including porter 2.   ![It hasn't been updated in the last year][Yellow]
* [shamoji](https://github.com/osamingo/shamoji) **star:10** The shamoji is word filtering package written in Go.
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   ![It hasn't been updated in the last year][Yellow]
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   ![It hasn't been updated in the last year][Yellow]
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)

## Networking

*Libraries for working with various layers of the network.*

* [kcptun](https://github.com/xtaci/kcptun) **star:10688** Extremely simple & fast udp tunnel based on KCP protocol.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [fasthttp](https://github.com/valyala/fasthttp) **star:9412** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [dns](https://github.com/miekg/dns) **star:3813** Go library for working with DNS.   ![star > 1000][Silver]
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3410** HTTPLabs let you inspect HTTP requests and forge responses.   ![star > 1000][Silver]
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:2928** An implementation of the QUIC protocol in pure Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gopacket](https://github.com/google/gopacket) **star:2883** Go library for packet processing with libpcap bindings.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [webrtc](https://github.com/pions/webrtc) **star:2280** A pure Go implementation of the WebRTC API.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2248** KCP - Fast and Reliable ARQ Protocol.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gobgp](https://github.com/osrg/gobgp) **star:1693** BGP implemented in the Go Programming Language.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [ssh](https://github.com/gliderlabs/ssh) **star:1113** Higher-level API for building SSH servers (wraps crypto/ssh).   ![star > 1000][Silver]
* [fortio](https://github.com/fortio/fortio) **star:876** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   ![star > 100][Bronze]
* [water](https://github.com/songgao/water) **star:842** Simple TUN/TAP library.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [sftp](https://github.com/pkg/sftp) **star:739** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-getter](https://github.com/hashicorp/go-getter) **star:725** Go library for downloading files or directories from various sources using a URL.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [NFF-Go](https://github.com/intel-go/nff-go) **star:665** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   ![star > 100][Bronze]   ![There was an update last week][Green]
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [mdns](https://github.com/hashicorp/mdns) **star:551** Simple mDNS (Multicast DNS) client/server library in Golang.   ![star > 100][Bronze]
* [grab](https://github.com/cavaliercoder/grab) **star:546** Go package for managing file downloads.   ![star > 100][Bronze]
* [ftp](https://github.com/jlaffaye/ftp) **star:529** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   ![star > 100][Bronze]
* [lhttp](https://github.com/fanux/lhttp) **star:513** Powerful websocket framework, build your IM server more easily.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [gosnmp](https://github.com/soniah/gosnmp) **star:438** Native Go library for performing SNMP actions.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [gotcp](https://github.com/gansidui/gotcp) **star:418** Go package for quickly writing tcp applications.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [cidranger](https://github.com/yl2chen/cidranger) **star:387** Fast IP to CIDR lookup for Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:364** Pure Go library for cross-platform local peer discovery using UDP multicast.   ![star > 100][Bronze]
* [gopcap](https://github.com/akrennmair/gopcap) **star:353** Go wrapper for libpcap.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-stun](https://github.com/ccding/go-stun) **star:334** Go implementation of the STUN client (RFC 3489 and RFC 5389).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [raw](https://github.com/mdlayher/raw) **star:303** Package raw enables reading and writing data at the device driver level for a network interface.   ![star > 100][Bronze]
* [stun](https://github.com/go-rtc/stun) **star:286** Go implementation of RFC 5389 STUN protocol.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [tcp_server](https://github.com/firstrow/tcp_server) **star:285** Go library for building tcp servers faster.   ![star > 100][Bronze]
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:232** Streaming protocolbuffer data over TCP made easy.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [winrm](https://github.com/masterzen/winrm) **star:213** Go WinRM client to remotely execute commands on Windows machines.   ![star > 100][Bronze]
* [arp](https://github.com/mdlayher/arp) **star:196** Package arp implements the ARP protocol, as described in RFC 826.   ![star > 100][Bronze]
* [ethernet](https://github.com/mdlayher/ethernet) **star:185** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   ![star > 100][Bronze]
* [utp](https://github.com/anacrolix/utp) **star:149** Go uTP micro transport protocol implementation.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [canopus](https://github.com/zubairhamed/canopus) **star:135** CoAP Client/Server implementation (RFC 7252).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [jazigo](https://github.com/udhos/jazigo) **star:125** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   ![star > 100][Bronze]
* [sslb](https://github.com/eduardonunesp/sslb) **star:112** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   ![star > 100][Bronze]
* [gNxI](https://github.com/google/gnxi) **star:103** A collection of tools for Network Management that use the gNMI and gNOI protocols.   ![star > 100][Bronze]
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:90** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [xtcp](https://github.com/xfxdev/xtcp) **star:83** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.
* [ether](https://github.com/songgao/ether) **star:62** Cross-platform Go package for sending and receiving ethernet frames.   ![It hasn't been updated in the last year][Yellow]
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:62** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
* [linkio](https://github.com/ian-kent/linkio) **star:44** Network link speed simulation for Reader/Writer interfaces.   ![It hasn't been updated in the last year][Yellow]
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   ![It hasn't been updated in the last year][Yellow]
* [packet](https://github.com/aerogo/packet) **star:28** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.
* [iplib](https://github.com/c-robinson/iplib) **star:24** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)
* [graval](https://github.com/koofr/graval) **star:24** Experimental FTP server framework.   ![It hasn't been updated in the last year][Yellow]
* [publicip](https://github.com/polera/publicip) **star:18** Package publicip returns your public facing IPv4 address (internet egress).   ![It hasn't been updated in the last year][Yellow]
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:15** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   ![It hasn't been updated in the last year][Yellow]
* [goshark](https://github.com/sunwxg/goshark) **star:9** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   ![It hasn't been updated in the last year][Yellow]
* [llb](https://github.com/kirillDanshin/llb) **star:8** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   ![It hasn't been updated in the last year][Yellow]
* [tspool](https://github.com/two/tspool) **star:5** A TCP Library use worker pool to improve performance and protect your server.

### HTTP Clients

*Libraries for making HTTP requests.*

* [grequests](https://github.com/levigross/grequests) **star:1416** A Go "clone" of the great and famous Requests library.   ![star > 1000][Silver]
* [heimdall](https://github.com/gojektech/heimdall) **star:1082** An enchanced http client with retry and hystrix capabilities.   ![star > 1000][Silver]
* [sling](https://github.com/dghubble/sling) **star:1006** Sling is a Go HTTP client library for creating and sending API requests.   ![star > 1000][Silver]
* [gentleman](https://github.com/h2non/gentleman) **star:678** Full-featured plugin-driven HTTP client library.   ![star > 100][Bronze]
* [pester](https://github.com/sethgrid/pester) **star:332** Go HTTP client calls with retries, backoff, and concurrency.   ![star > 100][Bronze]
* [goreq](https://github.com/smallnest/goreq) **star:98** Enhanced simplified HTTP client based on gorequest.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [rq](https://github.com/ddo/rq) **star:27** A nicer interface for golang stdlib HTTP client.   ![There was an update last week][Green]

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) **star:735** Go bindings for GLFW 3.   ![star > 100][Bronze]
* [gl](https://github.com/go-gl/gl) **star:641** Go bindings for OpenGL (generated via glow).   ![star > 100][Bronze]
* [mathgl](https://github.com/go-gl/mathgl) **star:291** Pure Go math package specialized for 3D math, with inspiration from GLM.   ![star > 100][Bronze]
* [goxjs/gl](https://github.com/goxjs/gl) **star:129** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   ![star > 100][Bronze]
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:57** Go cross-platform glfw library for creating an OpenGL context and receiving events.

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) **star:14728** The fantastic ORM library for Golang, aims to be developer friendly.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Xorm](https://github.com/go-xorm/xorm) **star:5220** Simple and powerful ORM for Go.   ![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [gorp](https://github.com/go-gorp/gorp) **star:3080** Go Relational Persistence, ORM-ish library for Go.   ![star > 1000][Silver]
* [go-pg](https://github.com/go-pg/pg) **star:3012** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2277** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   ![star > 1000][Silver]
* [upper.io/db](https://github.com/upper/db) **star:1852** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [reform](https://github.com/go-reform/reform) **star:804** Better ORM for Go, based on non-empty interfaces and code generation.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [pop/soda](https://github.com/gobuffalo/pop) **star:680** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [QBS](https://github.com/coocood/qbs) **star:538** Stands for Query By Struct. A Go ORM.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [go-queryset](https://github.com/jirfag/go-queryset) **star:448** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.   ![star > 100][Bronze]
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [Zoom](https://github.com/albrow/zoom) **star:239** Blazing-fast datastore and querying engine built on Redis.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:236** A flexible and powerful SQL string builder library plus a zero-config ORM.   ![star > 100][Bronze]
* [grimoire](https://github.com/Fs02/grimoire) **star:112** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-store](https://github.com/gosuri/go-store) **star:92** Simple and fast Redis backed key-value store library for Go.   ![It hasn't been updated in the last year][Yellow]
* [Marlow](https://github.com/dadleyy/marlow) **star:67** Generated ORM from project structs for compile time safety assurances.
* [lore](https://github.com/abrahambotros/lore) **star:4** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   ![It hasn't been updated in the last year][Yellow]
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:1** A simple ORM for Google/Firebase Cloud Firestore.

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:12586** Go dependency tool.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) **star:7785** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.   ![star > 5000][Gold]
* [godep](https://github.com/tools/godep) **star:5650** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   ![star > 5000][Gold]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [govendor](https://github.com/kardianos/govendor) **star:4746** Go Package Manager. Go vendor tool that works with the standard vendor file.   ![star > 1000][Silver]
* [gopm](https://github.com/gpmgo/gopm) **star:2364** Go Package Manager.   ![star > 1000][Silver]   ![Archived][Archived]
* [gom](https://github.com/mattn/gom) **star:1351** Go Manager - bundle for go.   ![star > 1000][Silver]
* [gpm](https://github.com/pote/gpm) **star:1204** Barebones dependency manager for Go.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [goop](https://github.com/nitrous-io/goop) **star:776** Simple dependency manager for Go (golang), inspired by Bundler.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [nut](https://github.com/jingweno/nut) **star:244** Vendor Go dependencies.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:213** Minimal dependency version using Git.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gigo](https://github.com/LyricalSecurity/gigo) **star:196** PIP-like dependency tool for golang, with support for private repositories and hashes.   ![star > 100][Bronze]
* [VenGO](https://github.com/DamnWidget/VenGO) **star:114** create and manage exportable isolated go virtual environments.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:87** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.   ![There was an update last week][Green]
* [gop](https://github.com/lunny/gop) **star:50** Build and manage your Go applications out of GOPATH.   ![Archived][Archived]   ![Contains Chinese documents][CN]

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:8694** A distributed tracing system.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [profile](https://github.com/pkg/profile) **star:1011** Simple profiling support package for Go.   ![star > 1000][Silver]
* [tracer](https://github.com/kamilsk/tracer) **star:8** Simple, lightweight tracing.   ![There was an update last week][Green]

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:5247** Implementation of GraphQL for Go.   ![star > 5000][Gold]
* [graphql](https://github.com/neelance/graphql-go) **star:2797** GraphQL server with a focus on ease of use.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:864** A simple Go package to Query over JSON Data.   ![star > 100][Bronze]
* [jsonql](https://github.com/elgs/jsonql) **star:201** JSON query expression library in Golang.   ![star > 100][Bronze]
* [rql](https://github.com/a8m/rql) **star:111** Resource Query Language for REST API.   ![star > 100][Bronze]
* [graphql](https://github.com/tmc/graphql) **star:51** graphql parser + utilities.   ![It hasn't been updated in the last year][Yellow]
* [jsonslice](https://github.com/bhmj/jsonslice) **star:24** Jsonpath queries with advanced filters.

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) **star:2134** The simple and easy way to embed static files into Go binaries.   ![star > 1000][Silver]
* [statik](https://github.com/rakyll/statik) **star:2115** Embeds static files into a Go executable.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1653** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.   ![star > 1000][Silver]
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:660** Generates a vfsdata.go file that statically implements the given virtual filesystem.   ![star > 100][Bronze]
* [esc](https://github.com/mjibson/esc) **star:472** Embeds files into Go programs and provides http.FileSystem interfaces to them.   ![star > 100][Bronze]
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:425** Simple tool to embed files in go with focus on "customization" and ease to use.   ![star > 100][Bronze]
* [go-resources](https://github.com/omeid/go-resources) **star:154** Unfancy resources embedding with Go.   ![star > 100][Bronze]
* [statics](https://github.com/go-playground/statics) **star:53** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   ![It hasn't been updated in the last year][Yellow]
* [templify](https://github.com/wlbr/templify) **star:20** Embed external template files into Go code to create single file binaries.
* [go-embed](https://github.com/pyros2097/go-embed) **star:16** Generates go code to embed resource files into your library or executable.

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:2969** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [stats](https://github.com/montanaflynn/stats) **star:1344** Statistics package with common functions missing from the Golang standard library.   ![star > 1000][Silver]
* [streamtools](https://github.com/nytlabs/streamtools) **star:1313** general purpose, graphical tool for dealing with streams of data.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [gosl](https://github.com/cpmech/gosl) **star:1309** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gonum/plot](https://github.com/gonum/plot) **star:1218** gonum/plot provides an API for building and drawing plots in Go.   ![star > 1000][Silver]
* [go-dsp](https://github.com/mjibson/go-dsp) **star:626** Digital Signal Processing for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [goraph](https://github.com/gyuho/goraph) **star:598** Pure Go graph theory library(data structure, algorith visualization).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [chart](https://github.com/vdobler/chart) **star:580** Simple Chart Plotting library for Go. Supports many graphs types.   ![star > 100][Bronze]
* [ewma](https://github.com/VividCortex/ewma) **star:265** Exponentially-weighted moving averages.   ![star > 100][Bronze]
* [graph](https://github.com/yourbasic/graph) **star:236** Library of basic graph algorithms.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [orb](https://github.com/paulmach/orb) **star:199** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:127** Approximate histograms for data streams.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:78** Dataframes for Go for machine-learning and statistics (similar to pandas).   ![There was an update last week][Green]
* [sparse](https://github.com/james-bowman/sparse) **star:69** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:66** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.
* [pagerank](https://github.com/alixaxel/pagerank) **star:49** Weighted PageRank algorithm implemented in Go.
* [geom](https://github.com/skelterjohn/geom) **star:40** 2D geometry for golang.   ![It hasn't been updated in the last year][Yellow]
* [evaler](https://github.com/soniah/evaler) **star:40** Simple floating point arithmetic expression evaluator.   ![It hasn't been updated in the last year][Yellow]
* [goent](https://github.com/kzahedi/goent) **star:13** GO Implementation of Entropy Measures.
* [triangolatte](https://github.com/tchayen/triangolatte) **star:11** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.
* [ode](https://github.com/ChristopherRabotin/ode) **star:10** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   ![It hasn't been updated in the last year][Yellow]
* [PiHex](https://github.com/claygod/PiHex) **star:9** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.
* [GoStats](https://github.com/OGFris/GoStats) **star:9** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.
* [go-gt](https://github.com/ThePaw/go-gt) **star:6** Graph theory algorithms written in "Go" language.   ![It hasn't been updated in the last year][Yellow]
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:5** Tiny linear interpolation library.
* [bradleyterry](https://github.com/seanhagen/bradleyterry)  Provides a Bradley-Terry Model for pairwise comparisons.
* [assocentity](https://github.com/ndabAP/assocentity) **star:4** Package assocentity returns the average distance from words to a given entity.   ![There was an update last week][Green]
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/xenolf/lego) **star:3525** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:1824** Tool and library to remotely hack RTSP streams from surveillance cameras.   ![star > 1000][Silver]
* [acmetool](https://github.com/hlandau/acme) **star:1694** ACME (Let's Encrypt) client tool with automatic renewal.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [memguard](https://github.com/awnumar/memguard) **star:1522** A pure Go library for handling sensitive values in memory.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [secure](https://github.com/unrolled/secure) **star:1210** HTTP middleware for Go that facilitates some quick security wins.   ![star > 1000][Silver]
* [acra](https://github.com/cossacklabs/acra) **star:452** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [nacl](https://github.com/kevinburke/nacl) **star:451** Go implementation of the NaCL set of API's.   ![star > 100][Bronze]
* [BadActor](https://github.com/jaredfolkins/badactor) **star:246** In-memory, application-driven jailer built in the spirit of fail2ban.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [passlib](https://github.com/hlandau/passlib) **star:225** Futureproof password hashing library.   ![star > 100][Bronze]
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:196** encrypt/decrypt using ssh keys.   ![star > 100][Bronze]
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:154** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   ![star > 100][Bronze]
* [go-yara](https://github.com/hillu/go-yara) **star:133** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   ![star > 100][Bronze]
* [argon2pw](https://github.com/raja/argon2pw) **star:75** Argon2 password hash generation with constant-time password comparison.
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:30** A probably paranoid package for securely hashing and encrypting passwords.
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.
* [certificates](https://github.com/mvmaasakkers/certificates) **star:9** An opinionated tool for generating tls certificates.
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:7** SSL certificates made easy with a high level wrapper around acme/autocert.
* [jwc](https://github.com/khezen/jwc) **star:5** JSON Web Cryptography library.

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:5559** High-performance 100% compatible drop-in replacement of "encoding/json".   ![star > 5000][Gold]
* [goprotobuf](https://github.com/golang/protobuf) **star:5105** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:2965** Protocol Buffers for Go with Gadgets.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:2442** Go library for decoding generic map values into native Go structures.   ![star > 1000][Silver]
* [go-codec](https://github.com/ugorji/go) **star:1236** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   ![star > 1000][Silver]
* [colfer](https://github.com/pascaldekloe/colfer) **star:471** Code generation for the Colfer binary format.   ![star > 100][Bronze]
* [csvutil](https://github.com/jszwec/csvutil) **star:306** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   ![star > 100][Bronze]
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:272** Cap'n Proto library and parser for go.   ![star > 100][Bronze]
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:119** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   ![star > 100][Bronze]
* [structomap](https://github.com/tuvistavie/structomap) **star:96** Library to easily and dynamically generate maps from static structures.
* [bambam](https://github.com/glycerine/bambam) **star:59** generator for Cap'n Proto schemas from go.   ![It hasn't been updated in the last year][Yellow]
* [asn1](https://github.com/PromonLogicalis/asn1) **star:40** Asn.1 BER and DER encoding library for golang.   ![Archived][Archived]
* [binstruct](https://github.com/ghostiam/binstruct) **star:7** Golang binary decoder for mapping data into the structure.
* [fwencoder](https://github.com/o1egl/fwencoder) **star:6** Fixed width file parser (encoding and decoding library) for Go.   ![It hasn't been updated in the last year][Yellow]
* [bel](https://github.com/32leaves/bel) **star:5** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.

## Server Applications

* [etcd](https://github.com/coreos/etcd) **star:26583** Highly-available key value store for shared configuration and service discovery.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Caddy](https://github.com/mholt/caddy) **star:23205** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [minio](https://github.com/minio/minio) **star:17545** Minio is a distributed object storage server.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.
* [RoadRunner](https://github.com/spiral/roadrunner) **star:3301** High-performance PHP application server, load-balancer and process manager.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [devd](https://github.com/cortesi/devd) **star:2806** Local webserver for developers.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [algernon](https://github.com/xyproto/algernon) **star:1590** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   ![star > 1000][Silver]
* [flipt](https://github.com/markphelps/flipt) **star:991** A self contained feature flag solution written in Go and Vue.js   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Flagr](https://github.com/checkr/flagr) **star:828** Flagr is an open-source feature flagging and A/B testing service.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Fider](https://github.com/getfider/fider) **star:798** Fider is an open platform to collect and organize customer feedback.   ![star > 100][Bronze]
* [jackal](https://github.com/ortuman/jackal) **star:718** An XMPP server written in Go.   ![star > 100][Bronze]
* [discovery](https://github.com/Bilibili/discovery) **star:677** A registry for resilient mid-tier load balancing and failover.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:6** Stream database events from PostgreSQL to Kafka.
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:5** Nginx log parser and exporter to Prometheus.
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.

## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:187** Go stream processing library.   ![star > 100][Bronze]

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3081** PDF document generator with high level support for text, drawing and images.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [pongo2](https://github.com/flosch/pongo2) **star:1490** Django-like template-engine for Go.   ![star > 1000][Silver]
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1409** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   ![star > 1000][Silver]
* [hero](https://github.com/shiyanhui/hero) **star:1205** Hero is a handy, fast and powerful go template engine.   ![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [mustache](https://github.com/hoisie/mustache) **star:968** Go implementation of the Mustache template language.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [amber](https://github.com/eknkc/amber) **star:822** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   ![star > 100][Bronze]
* [ace](https://github.com/yosssi/ace) **star:761** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Razor](https://github.com/sipin/gorazor) **star:701** Razor view engine for Golang.   ![star > 100][Bronze]
* [jet](https://github.com/CloudyKit/jet) **star:582** Jet template engine.   ![star > 100][Bronze]
* [ego](https://github.com/benbjohnson/ego) **star:415** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   ![star > 100][Bronze]
* [raymond](https://github.com/aymerick/raymond) **star:342** Complete handlebars implementation in Go.   ![star > 100][Bronze]
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:299** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).   ![star > 100][Bronze]
* [Soy](https://github.com/robfig/soy) **star:143** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   ![star > 100][Bronze]
* [liquid](https://github.com/osteele/liquid) **star:84** Go implementation of Shopify Liquid templates.
* [kasia.go](https://github.com/ziutek/kasia.go) **star:70** Templating system for HTML and other text documents - go implementation.   ![It hasn't been updated in the last year][Yellow]
* [velvet](https://github.com/gobuffalo/velvet) **star:64** Complete handlebars implementation in Go.   ![It hasn't been updated in the last year][Yellow]
* [goview](https://github.com/foolin/goview) **star:46** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
* [damsel](https://github.com/dskinner/damsel) **star:20** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   ![It hasn't been updated in the last year][Yellow]
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:13** Tiny wrapper around html/template to allow for easy file-based template inheritance.

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [Testify](https://github.com/stretchr/testify) **star:8168** Sacred extension to the standard go testing package.   ![star > 5000][Gold]
    * [go-cmp](https://github.com/google/go-cmp) **star:1175** Package for comparing Go values in tests.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1141** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   ![star > 1000][Silver]
    * [godog](https://github.com/DATA-DOG/godog) **star:760** Cucumber or Behat like BDD framework for Go.   ![star > 100][Bronze]
    * [baloo](https://github.com/h2non/baloo) **star:648** Expressive and versatile end-to-end HTTP API testing made easy.   ![star > 100][Bronze]
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [goblin](https://github.com/franela/goblin) **star:625** Mocha like testing framework fo Go.   ![star > 100][Bronze]
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:332** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   ![star > 100][Bronze]
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:329** A helper for Rails' like test fixtures to test database applications.   ![star > 100][Bronze]
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:293** Mutation testing for Go source code.   ![star > 100][Bronze]
    * [gofight](https://github.com/appleboy/gofight) **star:258** API Handler Testing for Golang Router framework.   ![star > 100][Bronze]
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [frisby](https://github.com/verdverm/frisby) **star:248** REST API testing framework.   ![star > 100][Bronze]
    * [go-carpet](https://github.com/msoap/go-carpet) **star:195** Tool for viewing test coverage in terminal.   ![star > 100][Bronze]
    * [charlatan](https://github.com/percolate/charlatan) **star:190** Tool to generate fake interface implementations for tests.   ![star > 100][Bronze]
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:122** A collection of packages to augment the go testing package and support common patterns.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [GoSpec](https://github.com/orfjackal/gospec) **star:110** BDD-style testing framework for the Go programming language.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [endly](https://github.com/viant/endly) **star:94** Declarative end to end functional testing.   ![There was an update last week][Green]
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:87** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:83** Simple snapshot testing addon for your test framework.
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [wstest](https://github.com/posener/wstest) **star:62** Websocket client for unit-testing a websocket http.Handler.   ![It hasn't been updated in the last year][Yellow]
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:55** Extremely flexible golang deep comparison, extends the go testing package.
    * [gospecify](https://github.com/stesla/gospecify) **star:50** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   ![It hasn't been updated in the last year][Yellow]
    * [restit](https://github.com/yookoala/restit) **star:48** Go micro framework to help writing RESTful API integration test.   ![It hasn't been updated in the last year][Yellow]
    * [commander](https://github.com/SimonBaeumer/commander) **star:32** Tool for testing cli applications on windows, linux and osx.   ![There was an update last week][Green]
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:30** library created for testing JSON against patterns.
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:26** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   ![It hasn't been updated in the last year][Yellow]
    * [dsunit](https://github.com/viant/dsunit) **star:25** Datastore testing for SQL, NoSQL, structured files.
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:22** Package for verifying that your JSON payloads are serialized correctly.
    * [assert](https://github.com/go-playground/assert) **star:13** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   ![It hasn't been updated in the last year][Yellow]
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   ![It hasn't been updated in the last year][Yellow]
    * [testcase](https://github.com/adamluzsi/testcase) **star:9** Idiomatic testing framework for Behavior Driven Development.
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Extensions to Go's `testing/iotest` package.   ![It hasn't been updated in the last year][Yellow]
    * [gocrest](https://github.com/corbym/gocrest) **star:8** Composable hamcrest-like matchers for Go assertions.   ![It hasn't been updated in the last year][Yellow]
    * [gogiven](https://github.com/corbym/gogiven) **star:7** YATSPEC-like BDD testing framework for Go.   ![It hasn't been updated in the last year][Yellow]
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.
    * [biff](https://github.com/fulldump/biff) **star:6** Bifurcation testing framework, BDD compatible.   ![It hasn't been updated in the last year][Yellow]
    * [Tt](https://github.com/vcaesar/tt) **star:5** Simple and colorful test tools.   ![There was an update last week][Green]
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:1** HTTP client testing framework.

* Mock
    * [gomock](https://github.com/golang/mock) **star:2862** Mocking framework for the Go programming language.   ![star > 1000][Silver]
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:1750** Mock SQL driver for testing database interactions.   ![star > 1000][Silver]
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1441** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   ![star > 1000][Silver]
    * [gock](https://github.com/h2non/gock) **star:822** Versatile HTTP mocking made easy.   ![star > 100][Bronze]
    * [httpmock](https://github.com/jarcoal/httpmock) **star:583** Easy mocking of HTTP responses from external resources.   ![star > 100][Bronze]
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:361** Tool for generating self-contained mock objects.   ![star > 100][Bronze]
    * [minimock](https://github.com/gojuno/minimock) **star:264** Mock generator for Go interfaces.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:164** Single transaction based database driver mainly for testing purposes.   ![star > 100][Bronze]
    * [govcr](https://github.com/seborama/govcr) **star:82** HTTP mock for Golang: record and replay HTTP interactions for offline testing.
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   ![It hasn't been updated in the last year][Yellow]

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:2890** Randomized testing system.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [gofuzz](https://github.com/google/gofuzz) **star:532** Library for populating go objects with random values.   ![star > 100][Bronze]
    * [Tavor](https://github.com/zimmski/tavor) **star:212** Generic fuzzing and delta-debugging framework.   ![star > 100][Bronze]

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:3602** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [selenoid](https://github.com/aerokube/selenoid) **star:1234** alternative Selenium hub server that launches browsers within containers.   ![star > 1000][Silver]
    * [cdp](https://github.com/mafredri/cdp) **star:354** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   ![star > 100][Bronze]
    * [ggr](https://github.com/aerokube/ggr) **star:209** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   ![star > 100][Bronze]

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:391** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   ![star > 100][Bronze]

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:8452** Fast and Elegant Scraping Framework for Gophers.   ![star > 5000][Gold]
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:7608** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   ![star > 5000][Gold]
    * [blackfriday](https://github.com/russross/blackfriday) **star:3900** Markdown processor in Go.   ![star > 1000][Silver]
    * [toml](https://github.com/BurntSushi/toml) **star:2792** TOML configuration format (encoder/decoder with reflection).   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [sh](https://github.com/mvdan/sh) **star:2005** Shell parser and formatter.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [go-humanize](https://github.com/dustin/go-humanize) **star:1904** Formatters for time, numbers, and memory size to human readable format.   ![star > 1000][Silver]
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1251** HTML Sanitizer.   ![star > 1000][Silver]   ![There was an update last week][Green]
    * [inject](https://github.com/facebookgo/inject) **star:1140** Package inject provides a reflect based injector.   ![star > 1000][Silver]   ![Archived][Archived]
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1100** Parse RSS and Atom feeds in Go.   ![star > 1000][Silver]
    * [go-toml](https://github.com/pelletier/go-toml) **star:609** Go library for the TOML format with query support and handy cli tools.   ![star > 100][Bronze]
    * [commonregex](https://github.com/mingrammer/commonregex) **star:553** A collection of common regular expressions for Go.   ![star > 100][Bronze]
    * [slug](https://github.com/gosimple/slug) **star:373** URL-friendly slugify with multiple languages support.   ![star > 100][Bronze]
    * [mxj](https://github.com/clbanning/mxj) **star:332** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:301** Parses the Graphviz DOT language.   ![star > 100][Bronze]
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:292** Web scraping Framework to turn websites into structured data.   ![star > 100][Bronze]
    * [gotext](https://github.com/leonelquinteros/gotext) **star:231** GNU gettext utilities for Go.   ![star > 100][Bronze]
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:209** Functions to get fixed width of the character or string.   ![star > 100][Bronze]
    * [goq](https://github.com/andrewstuart/goq) **star:146** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).   ![star > 100][Bronze]
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:130** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   ![star > 100][Bronze]   ![There was an update last week][Green]
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:98** NMEA parser library for the Go language.   ![There was an update last week][Green]
    * [sdp](https://github.com/gortc/sdp) **star:69** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].
    * [align](https://github.com/Guitarbum722/align) **star:59** A general purpose application that aligns text.   ![It hasn't been updated in the last year][Yellow]
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [genex](https://github.com/alixaxel/genex) **star:53** Count and expand Regular Expressions into all matching Strings.
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:50** Make pretty slug with multiple languages support.   ![It hasn't been updated in the last year][Yellow]
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:44** Functions to determine the natural language of a unicode text.   ![It hasn't been updated in the last year][Yellow]
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:41** Zero-width character detection and removal for Go.
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:37** Editorconfig file parser and manipulator for Go.   ![There was an update last week][Green]
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:36** Library for generating random strings from regular expressions.   ![It hasn't been updated in the last year][Yellow]
    * [allot](https://github.com/sbstjn/allot) **star:34** Placeholder and wildcard text parsing for CLI tools and bots.
    * [gonameparts](https://github.com/polera/gonameparts) **star:29** Parses human names into individual name parts.   ![There was an update last week][Green]
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:26** Fixed-width text formatting (encoder/decoder with reflection).
    * [Slugify](https://github.com/avelino/slugify) **star:26** Go slugify application that handles string.   ![It hasn't been updated in the last year][Yellow]
    * [go-vcard](https://github.com/emersion/go-vcard) **star:25** Parse and format vCard.
    * [did](https://github.com/ockam-network/did) **star:23** DID (Decentralized Identifiers) Parser and Stringer in Go.
    * [enca](https://github.com/endeveit/enca) **star:8** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   ![It hasn't been updated in the last year][Yellow]
    * [codetree](https://github.com/aerogo/codetree) **star:7** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   ![It hasn't been updated in the last year][Yellow]
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:5** A syndication feed for Atom 1.0 and RSS 2.0.   ![It hasn't been updated in the last year][Yellow]
    * [doi](https://github.com/hscells/doi) **star:4** Document object identifier (doi) parser in Go.   ![It hasn't been updated in the last year][Yellow]
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.
    * [ltsv](https://github.com/Wing924/ltsv) **star:2** High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:460** Extract urls from text.   ![star > 100][Bronze]
    * [gotabulate](https://github.com/bndr/gotabulate) **star:197** Easily pretty-print your tabular data with Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [radix](https://github.com/yourbasic/radix) **star:144** fast string sorting algorithm.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [parth](https://github.com/codemodus/parth) **star:31** URL path segmentation parsing.
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:17** Convert xml or json to go struct.
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:15** A sanitization-based swear filter for Go.
    * [kace](https://github.com/codemodus/kace) **star:12** Common case conversions covering common initialisms.
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:6** string argument parser that understands quotes and backslashes.   ![It hasn't been updated in the last year][Yellow]
    * [TySug](https://github.com/Dynom/TySug) **star:3** Alternative suggestions with respect to keyboard layouts.
    * [Tagify](https://github.com/zoomio/tagify) **star:1** Produces a set of tags from given source.

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5037** The official AWS SDK for the Go programming language.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [github](https://github.com/google/go-github) **star:4766** Go library for accessing the GitHub REST API v3.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [slack](https://github.com/nlopes/slack) **star:2424** Slack API in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [google](https://github.com/google/google-api-go-client) **star:1926** Auto-generated Google APIs for Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:1789** Google Cloud APIs Go Client Library.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:989** Go client library for the Twitter 1.1 API.   ![star > 100][Bronze]
* [discordgo](https://github.com/bwmarrin/discordgo) **star:970** Go bindings for the Discord Chat API.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [stripe](https://github.com/stripe/stripe-go) **star:944** Go client for the Stripe API.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [facebook](https://github.com/huandu/facebook) **star:771** Go Library that supports the Facebook Graph API.   ![star > 100][Bronze]
* [minio-go](https://github.com/minio/minio-go) **star:722** Minio Go Library for Amazon S3 compatible cloud storage.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-twitter](https://github.com/dghubble/go-twitter) **star:713** Go client library for the Twitter v1.1 APIs.   ![star > 100][Bronze]
* [go-jira](https://github.com/andygrunwald/go-jira) **star:576** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   ![star > 100][Bronze]
* [githubql](https://github.com/shurcooL/githubql) **star:503** Go library for accessing the GitHub GraphQL API v4.   ![star > 100][Bronze]
* [webhooks](https://github.com/go-playground/webhooks) **star:357** Webhook receiver for GitHub and Bitbucket.   ![star > 100][Bronze]
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:306** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   ![star > 100][Bronze]
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:300** Wrapper for PayPal payment API.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-marathon](https://github.com/gambol99/go-marathon) **star:189** Go library for interacting with Mesosphere's Marathon PAAS.   ![star > 100][Bronze]
* [ethrpc](https://github.com/onrik/ethrpc) **star:169** Go bindings for Ethereum JSON RPC API.   ![star > 100][Bronze]
* [gostorm](https://github.com/jsgilmore/gostorm) **star:117** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Medium](https://github.com/Medium/medium-sdk-go) **star:116** Golang SDK for Medium's OAuth2 API.   ![star > 100][Bronze]
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:113** A golang package to communicate with HipChat over XMPP.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hipchat](https://github.com/andybons/hipchat) **star:109** This project implements a golang client library for the Hipchat API.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Trello](https://github.com/adlio/trello) **star:102** Go wrapper for the Trello API.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-trending](https://github.com/andygrunwald/go-trending) **star:100** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   ![star > 100][Bronze]
* [cachet](https://github.com/andygrunwald/cachet) **star:66** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   ![It hasn't been updated in the last year][Yellow]
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Go client library for interfacing with the Clarifai API.   ![It hasn't been updated in the last year][Yellow]
* [pushover](https://github.com/gregdel/pushover) **star:57** Go wrapper for the Pushover API.
* [megos](https://github.com/andygrunwald/megos) **star:56** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   ![It hasn't been updated in the last year][Yellow]
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:52** Go client for the [Internet Game Database API](https://api.igdb.com/).
* [wit-go](https://github.com/wit-ai/wit-go) **star:47** Go client for wit.ai HTTP API.
* [gads](https://github.com/emiddleton/gads) **star:43** Google Adwords Unofficial API.
* [circleci](https://github.com/jszwedko/go-circleci) **star:41** Go client library for interacting with CircleCI's API.
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:39** Go client for the xkcd API.   ![It hasn't been updated in the last year][Yellow]
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:39** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   ![It hasn't been updated in the last year][Yellow]
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:36** Go MusicBrainz WS2 client library.
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:35** Go wrapper and command-line client for the Uptime Robot v2 API.   ![There was an update last week][Green]
* [fcm](https://github.com/maddevsio/fcm) **star:32** Go library for Firebase Cloud Messaging.
* [golyrics](https://github.com/mamal72/golyrics) **star:31** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   ![It hasn't been updated in the last year][Yellow]
* [mixpanel](https://github.com/dukex/mixpanel) **star:30** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
* [gcm](https://github.com/Aorioli/gcm) **star:29** Go library for Google Cloud Messaging.   ![It hasn't been updated in the last year][Yellow]
* [translate](https://github.com/poorny/translate) **star:26** Go online translation package.   ![It hasn't been updated in the last year][Yellow]
* [gami](https://github.com/bit4bit/gami) **star:26** Go library for Asterisk Manager Interface.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:20** Go client library for the [Unsplash.com](https://unsplash.com) API.
* [shopify](https://github.com/rapito/go-shopify) **star:19** Go Library to make CRUD request to the Shopify API.   ![It hasn't been updated in the last year][Yellow]
* [patreon-go](https://github.com/mxpv/patreon-go) **star:18** Go library for Patreon API.
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:17** Go client for interacting with the Twitch v3 API.   ![It hasn't been updated in the last year][Yellow]
* [spotify](https://github.com/rapito/go-spotify) **star:17** Go Library to access Spotify WEB API.   ![It hasn't been updated in the last year][Yellow]
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:16** Go library for accessing the BreweryDB API.   ![It hasn't been updated in the last year][Yellow]
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:16** Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).   ![It hasn't been updated in the last year][Yellow]
* [codeship-go](https://github.com/codeship/codeship-go) **star:15** Go client library for interacting with Codeship's API v2.
* [steam](https://github.com/sostronk/go-steam) **star:14** Go Library to interact with Steam game servers.   ![It hasn't been updated in the last year][Yellow]
* [textbelt](https://github.com/dietsche/textbelt) **star:14** Go client for the textbelt.com txt messaging API.   ![It hasn't been updated in the last year][Yellow]
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:13** Go client library for [imgur](https://imgur.com)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:13** Go wrapper for the YNAB API.
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:12** Simple wrapper for easy google analytics reporting.   ![It hasn't been updated in the last year][Yellow]
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   ![It hasn't been updated in the last year][Yellow]
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:9** Tiny Go client for HackerNews API.   ![It hasn't been updated in the last year][Yellow]
* [simples3](https://github.com/rhnvrm/simples3) **star:9** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:8** Go client library for the SPTrans Olho Vivo API.   ![It hasn't been updated in the last year][Yellow]
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   ![It hasn't been updated in the last year][Yellow]
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   ![It hasn't been updated in the last year][Yellow]
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).   ![It hasn't been updated in the last year][Yellow]
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   ![It hasn't been updated in the last year][Yellow]
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** The Playlyfe Rest API Go SDK.   ![It hasn't been updated in the last year][Yellow]
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:1** Go library MalShare API [malshare.com](http://www.malshare.com/)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:23254** Command-line fuzzy finder written in Go.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [hub](https://github.com/github/hub) **star:17052** wrap git commands with additional functionality to interact with github from the terminal.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [delve](https://github.com/derekparker/delve) **star:12064** Go debugger.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [ctop](https://github.com/bcicen/ctop) **star:8830** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   ![star > 5000][Gold]
* [wuzz](https://github.com/asciimoo/wuzz) **star:8265** Interactive cli tool for HTTP inspection.   ![star > 5000][Gold]
* [sqlx](https://github.com/jmoiron/sqlx) **star:6824** provides a set of extensions on top of the excellent built-in database/sql package.   ![star > 5000][Gold]
* [peco](https://github.com/peco/peco) **star:5472** Simplistic interactive filtering tool.   ![star > 5000][Gold]
* [usql](https://github.com/knq/usql) **star:4679** usql is a universal command-line interface for SQL databases.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:4512** Deliver Go binaries as fast and easily as possible.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [godropbox](https://github.com/dropbox/godropbox) **star:3744** Common libraries for writing Go services/applications from Dropbox.   ![star > 1000][Silver]
* [realize](https://github.com/tockins/realize) **star:3170** Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.   ![star > 1000][Silver]
* [goreporter](https://github.com/wgliang/goreporter) **star:2482** Golang tool that does static analysis, unit testing, code review and generate code quality report.   ![star > 1000][Silver]
* [panicparse](https://github.com/maruel/panicparse) **star:2144** Groups similar goroutines and colorizes stack dump.   ![star > 1000][Silver]
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2021** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   ![star > 1000][Silver]
* [resty](https://github.com/go-resty/resty) **star:1993** Simple HTTP and REST client for Go inspired by Ruby rest-client.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Task](https://github.com/go-task/task) **star:1940** simple "Make" alternative.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [minify](https://github.com/tdewolff/minify) **star:1860** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   ![star > 1000][Silver]
* [mmake](https://github.com/tj/mmake) **star:1446** Modern Make.   ![star > 1000][Silver]
* [Storm](https://github.com/asdine/storm) **star:1358** Simple and powerful toolkit for BoltDB.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [mole](https://github.com/davrodpin/mole) **star:1301** cli app to easily create ssh tunnels.   ![star > 1000][Silver]
* [go-funk](https://github.com/thoas/go-funk) **star:1196** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   ![star > 1000][Silver]
* [mc](https://github.com/minio/mc) **star:1120** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [filetype](https://github.com/h2non/filetype) **star:952** Small package to infer the file type checking the magic numbers signature.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [boilr](https://github.com/tmrts/boilr) **star:943** Blazingly fast CLI tool for creating projects from boilerplate templates.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mergo](https://github.com/imdario/mergo) **star:850** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   ![star > 100][Bronze]
* [spinner](https://github.com/briandowns/spinner) **star:798** Go package to easily provide a terminal spinner with options.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:795** Circuit Breakers in Go.   ![star > 100][Bronze]
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:718** Simple, seamless, lightweight time tracking for Git.   ![star > 100][Bronze]
* [jump](https://github.com/gsamokovarov/jump) **star:661** Jump helps you navigate faster by learning your habits.   ![star > 100][Bronze]
* [immortal](https://github.com/immortal/immortal) **star:603** \*nix cross-platform (OS agnostic) supervisor.   ![star > 100][Bronze]
* [htcat](https://github.com/htcat/htcat) **star:481** Parallel and Pipelined HTTP GET Utility.   ![star > 100][Bronze]
* [go-dry](https://github.com/ungerik/go-dry) **star:432** DRY (don't repeat yourself) package for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gopencils](https://github.com/bndr/gopencils) **star:422** Small and simple package to easily consume REST APIs.   ![star > 100][Bronze]
* [godaemon](https://github.com/VividCortex/godaemon) **star:404** Utility to write daemons.   ![star > 100][Bronze]
* [request](https://github.com/mozillazg/request) **star:355** Go HTTP Requests for Humans™.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [circuit](https://github.com/cep21/circuit) **star:329** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   ![star > 100][Bronze]
* [ergo](https://github.com/cristianoliveira/ergo) **star:312** The management of multiple local services running over different ports made easy.   ![star > 100][Bronze]
* [koazee](https://github.com/wesovilabs/koazee) **star:296** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   ![star > 100][Bronze]
* [go-rate](https://github.com/beefsack/go-rate) **star:291** Timed rate limiter for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gohper](https://github.com/cosiner/gohper) **star:247** Various tools/modules help for development.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [clockwork](https://github.com/jonboulle/clockwork) **star:220** A simple fake clock for golang.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Deepcopier](https://github.com/ulule/deepcopier) **star:209** Simple struct copying for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [serve](https://github.com/syntaqx/serve) **star:191** A static http server anywhere you need.   ![star > 100][Bronze]
* [go-trigger](https://github.com/sadlil/go-trigger) **star:182** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [retry](https://github.com/kamilsk/retry) **star:161** The most advanced functional mechanism to perform actions repetitively until successful.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:160** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [rerun](https://github.com/ivpusic/rerun) **star:153** Recompiling and rerunning go apps when source changes.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [moldova](https://github.com/StabbyCutyou/moldova) **star:148** Utility for generating random data based on an input template.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gotenv](https://github.com/subosito/gotenv) **star:142** Load environment variables from `.env` or any `io.Reader` in Go.   ![star > 100][Bronze]
* [gubrak](https://github.com/novalagung/gubrak) **star:141** Golang utility library with syntactic sugar. It's like lodash, but for golang.   ![star > 100][Bronze]
* [robustly](https://github.com/VividCortex/robustly) **star:134** Runs functions resiliently, catching and restarting panics.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Death](https://github.com/vrecan/death) **star:133** Managing go application shutdown with signals.   ![star > 100][Bronze]
* [util](https://github.com/shomali11/util) **star:133** Collection of useful utility functions. (strings, concurrency, manipulations, ...).   ![star > 100][Bronze]
* [apm](https://github.com/topfreegames/apm) **star:129** Process manager for Golang applications with an HTTP API.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:122** Package for MIME type detection based on magic numbers.   ![star > 100][Bronze]
* [chyle](https://github.com/antham/chyle) **star:109** Changelog generator using a git repository with multiple configuration possibilities.   ![star > 100][Bronze]
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:105** XML Sitemap generator written in Go.   ![star > 100][Bronze]
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:100** LiveReload server for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [onecache](https://github.com/adelowo/onecache) **star:98** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
* [toolbox](https://github.com/viant/toolbox) **star:96** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   ![There was an update last week][Green]
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:83** Pure Go bsdiff and bspatch libraries and CLI tools.
* [pm](https://github.com/VividCortex/pm) **star:71** Process (i.e. goroutine) manager with an HTTP API.
* [UNIS](https://github.com/esemplastic/unis) **star:69** Common Architecture™ for String Utilities in Go.   ![It hasn't been updated in the last year][Yellow]
* [xferspdy](https://github.com/monmohan/xferspdy) **star:68** Xferspdy provides binary diff and patch library in golang.   ![It hasn't been updated in the last year][Yellow]
* [go-health](https://github.com/Talento90/go-health) **star:63** Health package simplifies the way you add health check to your services.   ![It hasn't been updated in the last year][Yellow]
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:58** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
* [multitick](https://github.com/VividCortex/multitick) **star:57** Multiplexor for aligned tickers.   ![It hasn't been updated in the last year][Yellow]
* [repeat](https://github.com/ssgreg/repeat) **star:56** Go implementation of different backoff strategies useful for retrying operations and heartbeating.
* [abutil](https://github.com/bahlo/abutil) **star:51** Collection of often-used Golang helpers.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [minquery](https://github.com/icza/minquery) **star:50** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).
* [handy](https://github.com/miguelpragier/handy) **star:45** Many utilities and helpers like string handlers/formatters and validators.
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:45** Parse TODOs in your GO code.
* [golog](https://github.com/mlimaloureiro/golog) **star:42** Easy and lightweight CLI tool to time track your tasks.
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:42** Pure Go ultra performant MIME sniffing library/utility.
* [goback](https://github.com/carlescere/goback) **star:40** Go simple exponential backoff package.   ![It hasn't been updated in the last year][Yellow]
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:39** Use x86 SIMD without writing any assembly code.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [gaper](https://github.com/maxcnunes/gaper) **star:39** Builds and restarts a Go project when it crashes or some watched file changes.
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:37** Universal multi-workstation clipboard that uses S3 like backend for the storage.   ![There was an update last week][Green]
* [golarm](https://github.com/msempere/golarm) **star:34** Fire alarms with system events.   ![It hasn't been updated in the last year][Yellow]
* [retry](https://github.com/thedevsaddam/retry) **star:34** Simple and easy retry mechanism package for Go.   ![It hasn't been updated in the last year][Yellow]
* [myhttp](https://github.com/inancgumus/myhttp) **star:34** Simple API to make HTTP GET requests with timeout support.   ![It hasn't been updated in the last year][Yellow]
* [goreadability](https://github.com/philipjkim/goreadability) **star:32** Webpage summary extractor using Facebook Open Graph and arc90's readability.
* [gpath](https://github.com/tenntenn/gpath) **star:31** Library to simplify access struct fields with Go's expression in reflection.   ![It hasn't been updated in the last year][Yellow]
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:28** Retrying made simple and easy for golang.
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:27** SeaweedFS client library with almost full features.
* [rclient](https://github.com/zpatrick/rclient) **star:27** Readable, flexible, simple-to-use client for REST APIs.
* [pgo](https://github.com/arthurkushman/pgo) **star:24** Convenient functions for PHP community.
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:22** a small golang lib to generate placeholder images.   ![It hasn't been updated in the last year][Yellow]
* [ugo](https://github.com/alxrm/ugo) **star:20** ugo is slice toolbox with concise syntax for Go.   ![It hasn't been updated in the last year][Yellow]
* [generate](https://github.com/go-playground/generate) **star:19** runs go generate recursively on a specified path or environment variable and can filter by regex.   ![It hasn't been updated in the last year][Yellow]
* [evaluator](https://github.com/nullne/evaluator) **star:16** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   ![It hasn't been updated in the last year][Yellow]
* [gostrutils](https://github.com/ik5/gostrutils) **star:16** Collections of string manipulation and conversion functions.
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** Compile-time controlled logger to make your release smaller without removing debug calls.   ![It hasn't been updated in the last year][Yellow]
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:14** Go library for encoding structs into Header fields.
* [filler](https://github.com/yaronsumel/filler) **star:14** small utility to fill structs using "fill" tag.   ![It hasn't been updated in the last year][Yellow]
* [okrun](https://github.com/xta/okrun) **star:14** go run error steamroller.   ![It hasn't been updated in the last year][Yellow]
* [dbt](https://github.com/nikogura/dbt) **star:13** A framework for running self-updating signed binaries from a central, trusted repository.
* [structs](https://github.com/PumpkinSeed/structs) **star:12** Implement simple functions to manipulate structs.   ![It hasn't been updated in the last year][Yellow]
* [scan](https://github.com/blockloop/scan) **star:12** Scan golang `sql.Rows` directly to structs, slices, or primitive types.
* [filter](https://github.com/gookit/filter) **star:12** provide filtering, sanitizing, and conversion of Go data.
* [ghokin](https://github.com/antham/ghokin) **star:12** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).
* [rerate](https://github.com/abo/rerate) **star:12** Redis-based rate counter and rate limiter for Go.   ![It hasn't been updated in the last year][Yellow]
* [retry](https://github.com/shafreeck/retry) **star:10** A pretty simple library to ensure your work to be done.
* [command](https://github.com/txgruppi/command) **star:9** Command pattern for Go with thread safe serial and parallel dispatcher.   ![It hasn't been updated in the last year][Yellow]
* [backscanner](https://github.com/icza/backscanner) **star:8** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.
* [ctxutil](https://github.com/posener/ctxutil) **star:7** A collection of utility functions for contexts.
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:7** A MIME type sniffer for Go.
* [sslice](https://github.com/yaa110/sslice) **star:5** Create a slice which is always sorted.
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.   ![There was an update last week][Green]
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:3** Slice conversion between primitive types.
* [slicer](https://github.com/leaanthony/slicer) **star:3** Makes working with slices easier.
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:3** Go package for working with Problem Details.
* [retry](https://github.com/percolate/retry) **star:2** A simple but highly configurable retry package for Go.
* [blank](https://github.com/Henry-Sarabia/blank) **star:1** Verify or remove blanks and whitespace from strings.
* [olaf](https://github.com/btnguyen2k/olaf) **star:1** Twitter Snowflake implemented in Go.

## UUID

*Libraries for working with UUIDs.*

* [ulid](https://github.com/oklog/ulid) **star:1681** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   ![star > 1000][Silver]
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands. 
* [uuid](https://github.com/gofrs/uuid) **star:569** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   ![star > 100][Bronze]
* [wuid](https://github.com/edwingeng/wuid) **star:288** An extremely fast unique number generator, 10-135 times faster than UUID.   ![star > 100][Bronze]
* [goid](https://github.com/jakehl/goid) **star:21** Generate and Parse RFC4122 compliant V4 UUIDs.
* [sno](https://github.com/muyo/sno) **star:15** Compact, sortable and fast unique IDs with embedded metadata.
* [uuid](https://github.com/agext/uuid) **star:10** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:3** A tiny and efficient Go unique string ID generator.

## Validation

*Libraries for validation.*

* [govalidator](https://github.com/asaskevich/govalidator) **star:3557** Validators and sanitizers for strings, numerics, slices and structs.   ![star > 1000][Silver]
* [validator](https://github.com/go-playground/validator) **star:3530** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   ![star > 1000][Silver]
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1053** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:717** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   ![star > 100][Bronze]
* [validate](https://github.com/gookit/validate) **star:97** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   ![Contains Chinese documents][CN]
* [checkdigit](https://github.com/osamingo/checkdigit) **star:44** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
* [jio](https://github.com/faceair/jio) **star:21** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   ![Contains Chinese documents][CN]
* [validate](https://github.com/gobuffalo/validate) **star:19** This package provides a framework for writing validations for Go applications.
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:11** A norms and conventions validator for Terraform.   ![There was an update last week][Green]

## Version Control

*Libraries for version control.*

* [go-git](https://github.com/src-d/go-git) **star:4294** highly extensible Git implementation in pure Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [git2go](https://github.com/libgit2/git2go) **star:1362** Go bindings for libgit2.   ![star > 1000][Silver]
* [hercules](https://github.com/src-d/hercules) **star:540** gaining advanced insights from Git repository history.   ![star > 100][Bronze]
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:71** manipulate and inspect VCS repositories in Go.   ![There was an update last week][Green]
* [gh](https://github.com/rjeczalik/gh) **star:69** Scriptable server and net/http middleware for GitHub Webhooks.
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   ![It hasn't been updated in the last year][Yellow]

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) **star:790** Comphrensive Go bindings for FFmpeg.   ![star > 100][Bronze]
* [gmf](https://github.com/3d0c/gmf) **star:531** Go bindings for FFmpeg av\* libraries.   ![star > 100][Bronze]
* [go-astits](https://github.com/asticode/go-astits) **star:259** Parse and demux MPEG Transport Streams (.ts) natively in GO.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [go-astisub](https://github.com/asticode/go-astisub) **star:168** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   ![star > 100][Bronze]
* [gst](https://github.com/ziutek/gst) **star:151** Go bindings for GStreamer.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:64** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:38** Parser and generator library for Apple m3u8 playlists.
* [v4l](https://github.com/korandiz/v4l) **star:28** Video capture library for Linux, written in Go.   ![It hasn't been updated in the last year][Yellow]
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:11** Subtitle format support for go. Supports .srt, .ttml, and .ass.

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) **star:30109** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Beego](https://github.com/astaxie/beego) **star:21515** beego is an open-source, high-performance web framework for the Go programming language.   ![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) **star:14723** High performance, minimalist Go web framework.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Revel](https://github.com/revel/revel) **star:11253** High-productivity web framework for the Go language.   ![star > 5000][Gold]
* [Goa](https://github.com/goadesign/goa) **star:3511** Goa provides a holistic approach for developing remote APIs and microservices in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3335** Quick and easy way to setup a RESTful JSON API.   ![star > 1000][Silver]
* [Gizmo](https://github.com/NYTimes/gizmo) **star:2850** Microservice toolkit used by the New York Times.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Macaron](https://github.com/go-macaron/macaron) **star:2817** Macaron is a high productive and modular design web framework in Go.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [utron](https://github.com/gernest/utron) **star:2134** Lightweight MVC framework for Go(Golang).   ![star > 1000][Silver]
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:995** Go framework for building JSON web services inspired by Dropwizard.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tango](https://github.com/lunny/tango) **star:818** Micro & pluggable web framework for Go.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [traffic](https://github.com/pilu/traffic) **star:517** Sinatra inspired regexp/pattern mux and web framework for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gongular](https://github.com/mustafaakin/gongular) **star:415** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.   ![star > 100][Bronze]
* [neo](https://github.com/ivpusic/neo) **star:391** Neo is minimal and fast Go Web Framework with extremely simple API.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mango](https://github.com/paulbellamy/mango) **star:339** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Air](https://github.com/aofei/air) **star:332** An ideally refined web framework for Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Gondola](https://github.com/rainycape/gondola) **star:314** The web framework for writing faster sites, faster.   ![star > 100][Bronze]
* [Golf](https://github.com/dinever/golf) **star:235** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Aero](https://github.com/aerogo/aero) **star:158** High-performance web framework for Go, reaches top scores in Lighthouse.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Gem](https://github.com/go-gem/gem) **star:153** Simple and fast web framework, friendly to REST API.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-rest](https://github.com/ungerik/go-rest) **star:115** Small and evil REST framework for Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hiboot](https://github.com/hidevopsio/hiboot) **star:87** hiboot is a high performance web application framework with auto configuration and dependency injection support.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [WebGo](https://github.com/bnkamalesh/webgo) **star:74** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).
* [Golax](https://github.com/fulldump/golax) **star:71** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   ![It hasn't been updated in the last year][Yellow]
* [Microservice](https://github.com/claygod/microservice) **star:57** The framework for the creation of microservices, written in Golang.
* [uAdmin](https://github.com/uadmin/uadmin) **star:54** Fully featured web framework for Golang, inspired by Django.
* [YARF](https://github.com/yarf-framework/yarf) **star:50** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.
* [Fireball](https://github.com/zpatrick/fireball) **star:49** More "natural" feeling web framework.
* [vox](https://github.com/aisk/vox) **star:39** A golang web framework for humans, inspired by Koa heavily.
* [patron](https://github.com/beatlabs/patron) **star:34** Patron is a microservice framework following best cloud practices with a focus on productivity.   ![There was an update last week][Green]
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Resoursea](https://github.com/resoursea/api) **star:29** REST framework for quickly writing resource based services.   ![It hasn't been updated in the last year][Yellow]
* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [rex](https://github.com/goanywhere/rex) **star:27** Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.   ![It hasn't been updated in the last year][Yellow]
* [nio](https://github.com/go-nio/nio) **star:21** Modern, minimal and productive Go HTTP framework.   ![There was an update last week][Green]   ![Archived][Archived]
* [Banjo](https://github.com/nsheremet/banjo) **star:8** Very simple and fast web framework for Go.   ![It hasn't been updated in the last year][Yellow]
* [rux](https://github.com/gookit/rux) **star:7** Simple and fast web framework for build golang HTTP applications.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) **star:1252** Rate limit HTTP request handler.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [CORS](https://github.com/rs/cors) **star:1223** Easily add CORS capabilities to your API.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Limiter](https://github.com/ulule/limiter) **star:792** Dead simple rate limit middleware for Go.   ![star > 100][Bronze]
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:746** Add/parse Server-Timing header.   ![star > 100][Bronze]
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:89** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).
* [XFF](https://github.com/sebest/xff) **star:71** Handle `X-Forwarded-For` header and friends.
* [formjson](https://github.com/rs/formjson) **star:33** Transparently handle JSON input as a standard form POST.   ![It hasn't been updated in the last year][Yellow]
* [client-timing](https://github.com/posener/client-timing) **star:13** An HTTP client for Server-Timing header.

#### Libraries for creating HTTP middlewares

* [negroni](https://github.com/urfave/negroni) **star:6340** Idiomatic HTTP middleware for Golang.   ![star > 5000][Gold]   ![Contains Chinese documents][CN]
* [alice](https://github.com/justinas/alice) **star:1823** Painless middleware chaining for Go.   ![star > 1000][Silver]
* [render](https://github.com/unrolled/render) **star:1267** Go package for easily rendering JSON, XML, and HTML template responses.   ![star > 1000][Silver]
* [stats](https://github.com/thoas/stats) **star:536** Go middleware that stores various information about your web application.   ![star > 100][Bronze]
* [interpose](https://github.com/carbocation/interpose) **star:288** Minimalist net/http middleware for golang.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [muxchain](https://github.com/stephens2424/muxchain) **star:207** Lightweight middleware for net/http.   ![star > 100][Bronze]
* [renderer](https://github.com/thedevsaddam/renderer) **star:169** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   ![star > 100][Bronze]
* [rye](https://github.com/InVisionApp/rye) **star:94** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.
* [gores](https://github.com/alioygur/gores) **star:81** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
* [chain](https://github.com/codemodus/chain) **star:63** Handler wrapper chaining with scoped data (net/context-based "middleware").
* [go-wrap](https://github.com/go-on/wrap) **star:58** Small middlewares package for net/http.   ![It hasn't been updated in the last year][Yellow]
* [catena](https://github.com/codemodus/catena) **star:7** http.Handler wrapper catenation (same API as "chain").

### Routers

* [httprouter](https://github.com/julienschmidt/httprouter) **star:9736** High performance router. Use this and the standard http handlers to form a very high performance web framework.   ![star > 5000][Gold]
* [mux](https://github.com/gorilla/mux) **star:9703** Powerful URL router and dispatcher for golang.   ![star > 5000][Gold]
* [chi](https://github.com/go-chi/chi) **star:6054** Small, fast and expressive HTTP router built on net/context.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [gocraft/web](https://github.com/gocraft/web) **star:1396** Mux and middleware package in Go.   ![star > 1000][Silver]
* [Bone](https://github.com/go-zoo/bone) **star:1218** Lightning Fast HTTP Multiplexer.   ![star > 1000][Silver]
* [Goji](https://github.com/goji/goji) **star:768** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   ![star > 100][Bronze]
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:750** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   ![star > 100][Bronze]
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:452** A simple and fast HTTP router for Go.   ![star > 100][Bronze]
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:385** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   ![star > 100][Bronze]
* [lars](https://github.com/go-playground/lars) **star:374** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   ![star > 100][Bronze]
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:359** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   ![star > 100][Bronze]
* [Siesta](https://github.com/VividCortex/siesta) **star:349** Composable framework to write middleware and handlers.   ![star > 100][Bronze]
* [vestigo](https://github.com/husobee/vestigo) **star:250** Performant, stand-alone, HTTP compliant URL Router for go web applications.   ![star > 100][Bronze]
* [gowww/router](https://github.com/gowww/router) **star:157** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [alien](https://github.com/gernest/alien) **star:106** Lightweight and fast http router from outer space.   ![star > 100][Bronze]
* [violetear](https://github.com/nbari/violetear) **star:95** Go HTTP router.
* [Bxog](https://github.com/claygod/Bxog) **star:94** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
* [xmux](https://github.com/rs/xmux) **star:88** High performance muxer based on `httprouter` with `net/context` support.   ![It hasn't been updated in the last year][Yellow]
* [pure](https://github.com/go-playground/pure) **star:86** Is a lightweight HTTP router that sticks to the std "net/http" implementation.
* [GoRouter](https://github.com/vardius/gorouter) **star:49** GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:38** A simple Go HTTP router.
* [FastRouter](https://github.com/razonyang/fastrouter) **star:18** a fast, flexible HTTP router written in Go.   ![It hasn't been updated in the last year][Yellow]

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:550** Win32 OLE implementation for golang.   ![star > 100][Bronze]
* [d3d9](https://github.com/gonutz/d3d9) **star:85** Go bindings for Direct3D9.
* [gosddl](https://github.com/MonaxGT/gosddl) **star:1** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) **star:256** Generate a Go struct from XML.   ![star > 100][Bronze]
* [xpath](https://github.com/antchfx/xpath) **star:169** XPath package for Go.   ![star > 100][Bronze]
* [xquery](https://github.com/antchfx/xquery) **star:145** XQuery lets you extract data from HTML/XML documents using XPath expression.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** Simple command line XML comparer that generates diffs of folders, files and tags.   ![It hasn't been updated in the last year][Yellow]
* [xml2map](https://github.com/sbabiv/xml2map) **star:15** XML to MAP converter written Golang.
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:6** Procedural XML generation API based on libxml2's xmlwriter module.

# Tools

*Go software and plugins.*

## Code Analysis

* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [GoLint](https://github.com/golang/lint) **star:3158** Golint is a linter for Go source code.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [errcheck](https://github.com/kisielk/errcheck) **star:1319** Errcheck is a program for checking for unchecked errors in Go programs.   ![star > 1000][Silver]
* [gcvis](https://github.com/davecheney/gcvis) **star:918** Visualise Go program GC trace data in real time.   ![star > 100][Bronze]
* [php-parser](https://github.com/z7zmey/php-parser) **star:638** A Parser for PHP written in Go.   ![star > 100][Bronze]
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  tarp finds functions and methods without direct unit tests in Go source code.
* [go-critic](https://github.com/go-critic/go-critic) **star:579** source code linter that brings checks that are currently not implemented in other linters.   ![star > 100][Bronze]
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:375** Web based Golang AST visualizer.   ![star > 100][Bronze]
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:282** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [unconvert](https://github.com/mdempsky/unconvert) **star:258** Remove unnecessary type conversions from Go source.   ![star > 100][Bronze]
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [gostatus](https://github.com/shurcooL/gostatus) **star:240** Command line tool, shows the status of repositories that contain Go packages.   ![star > 100][Bronze]
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:185** An easy way to find outdated dependencies of your Go projects.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [dupl](https://github.com/mibk/dupl) **star:176** Tool for code clone detection.   ![star > 100][Bronze]
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:166** Checks recent changes to a Go project for backwards incompatible changes.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:95** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.
* [lint](https://github.com/surullabs/lint) **star:63** Run linters as part of go test.
* [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   ![It hasn't been updated in the last year][Yellow]
* [go-outdated](https://github.com/firstrow/go-outdated) **star:44** Console application that displays outdated packages.   ![Archived][Archived]

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) **star:10868** Go development plugin for Vim.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [vscode-go](https://github.com/Microsoft/vscode-go) **star:5136** Extension for Visual Studio Code (VS Code) which provides support for the Go language.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [gocode](https://github.com/nsf/gocode) **star:4743** Autocompletion daemon for the Go programming language.   ![star > 1000][Silver]
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3240** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.   ![star > 1000][Silver]
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1483** Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-mode](https://github.com/dominikh/go-mode.el) **star:958** Go mode for GNU/Emacs.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Watch](https://github.com/eaburns/Watch) **star:168** Runs a command in an acme win on file changes.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:81** Vim plugin to highlight syntax errors on save.   ![It hasn't been updated in the last year][Yellow]
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:30** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:17** Vim plugin for generating Go tests based on the function's or method's signature.
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:12** Go language support for the Theia IDE.

## Go Generate Tools

* [gotests](https://github.com/cweill/gotests) **star:2208** Generate Go tests from your source code.   ![star > 1000][Silver]
* [genny](https://github.com/cheekybits/genny) **star:973** Elegant generics for Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [re2dfa](https://github.com/opennota/re2dfa) **star:168** Transform regular expressions into finite state machines and output Go source code.   ![star > 100][Bronze]
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.
* [gocontracts](https://github.com/Parquery/gocontracts) **star:51** brings design-by-contract to Go by synchronizing the code with the documentation.
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [gounit](https://github.com/hexdigest/gounit) **star:29** Generate Go tests using your own templates.
* [generic](https://github.com/usk81/generic) **star:28** flexible data type for Go.
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:12** Generate Haskell inspired functions for your slices.

## Go Tools

* [go-swagger](https://github.com/go-swagger/go-swagger) **star:4004** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:1993** Visualize call graph of your Go program using dot format.   ![star > 1000][Silver]
* [richgo](https://github.com/kyoh86/richgo) **star:389** Enrich `go test` outputs with text decorations.   ![star > 100][Bronze]
* [depth](https://github.com/KyleBanks/depth) **star:380** Visualize dependency trees of any package by analyzing imports.   ![star > 100][Bronze]
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [rts](https://github.com/galeone/rts) **star:184** RTS: response to struct. Generates Go structs from server responses.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [godbg](https://github.com/tylerwince/godbg) **star:156** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.   ![star > 100][Bronze]
* [OctoLinker](https://github.com/OctoLinker/browser-extension)  Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
* [colorgo](https://github.com/songgao/colorgo) **star:97** Wrapper around `go` command for colorized `go build` output.   ![It hasn't been updated in the last year][Yellow]
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:37** Bash completion for go and wgo.   ![It hasn't been updated in the last year][Yellow]
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:13** A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:56492** Container Cluster Manager from Google.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Moby](https://github.com/moby/moby) **star:54479** Collaborative project for the container ecosystem to assemble container-based systems.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [traefik](https://github.com/containous/traefik) **star:23822** Reverse proxy and load balancer with support for multiple backends.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [Gitea](https://github.com/go-gitea/gitea) **star:15374** Fork of Gogs, entirely community driven.   ![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [Vegeta](https://github.com/tsenart/vegeta) **star:12216** HTTP load testing tool and library. It's over 9000!   ![star > 5000][Gold]
* [Hey](https://github.com/rakyll/hey) **star:6350** Hey is a tiny program that sends some load to a web application.   ![star > 5000][Gold]
* [GVM](https://github.com/moovweb/gvm) **star:4480** GVM provides an interface to manage Go versions.   ![star > 1000][Silver]
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [webhook](https://github.com/adnanh/webhook) **star:4097** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   ![star > 1000][Silver]
* [gaia](https://github.com/gaia-pipeline/gaia) **star:3750** Build powerful pipelines in any programming language.   ![star > 1000][Silver]
* [gox](https://github.com/mitchellh/gox) **star:3365** Dead simple, no frills Go cross compile tool.   ![star > 1000][Silver]
* [bosun](https://github.com/bosun-monitor/bosun) **star:2854** Time Series Alerting Framework.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [bombardier](https://github.com/codesenberg/bombardier) **star:1753** Fast cross-platform HTTP benchmarking tool.   ![star > 1000][Silver]
* [goxc](https://github.com/laher/goxc) **star:1623** build tool for Go, with a focus on cross-compiling and packaging.   ![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [fac](https://github.com/mkchoi212/fac) **star:1614** Command-line user interface to fix git merge conflicts.   ![star > 1000][Silver]
* [kala](https://github.com/ajvb/kala) **star:1358** Simplistic, modern, and performant job scheduler.   ![star > 1000][Silver]
* [StatusOK](https://github.com/sanathp/statusok) **star:1164** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.   ![star > 1000][Silver]
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:998** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.   ![star > 100][Bronze]
* [script](https://github.com/bitfield/script) **star:924** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   ![star > 100][Bronze]
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:673** Enable your Go applications to self update.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [skm](https://github.com/TimothyYe/skm) **star:550** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   ![star > 100][Bronze]
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:538** Manage BareMetal Servers from Command Line (as easily as with Docker).   ![star > 100][Bronze]
* [Pomerium](https://github.com/pomerium/pomerium) **star:518** Pomerium is an identity-aware access proxy.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [aurora](https://github.com/xuri/aurora) **star:404** Cross-platform web-based Beanstalkd queue server console.   ![star > 100][Bronze]
* [gonative](https://github.com/inconshreveable/gonative) **star:311** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [govvv](https://github.com/ahmetalpbalkan/govvv)  “go build” wrapper to easily add version information into Go binaries.
* [Mora](https://github.com/emicklei/mora) **star:266** REST server for accessing MongoDB documents and meta data.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [lstags](https://github.com/ivanilves/lstags) **star:219** Tool and API to sync Docker images across different registries.   ![star > 100][Bronze]
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [godbg](https://github.com/sirnewton01/godbg) **star:218** Web-based gdb front-end application.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [dogo](https://github.com/liudng/dogo) **star:216** Monitoring changes in the source file and automatically compile and run (restart).   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [manssh](https://github.com/xwjdsh/manssh) **star:205** manssh is a command line tool for managing your ssh alias config easily.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Pewpew](https://github.com/bengadbois/pewpew) **star:203** Flexible HTTP command line stress tester.   ![star > 100][Bronze]
* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [gobrew](https://github.com/cryptojuice/gobrew) **star:174** gobrew lets you easily switch between multiple versions of go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Blast](https://github.com/dave/blast) **star:167** A simple tool for API load testing and batch jobs.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [ostent](https://github.com/ostrost/ostent) **star:164** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Packer](https://github.com/mitchellh/packer)  Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
* [grapes](https://github.com/yaronsumel/grapes) **star:135** Lightweight tool designed to distribute commands over ssh with ease.   ![star > 100][Bronze]
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:101** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   ![star > 100][Bronze]
* [kcli](https://github.com/cswank/kcli) **star:78** Command line tool for inspecting kafka topics/partitions/messages.
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:67** Cli tool to remotely execute commands on Windows machines.
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:67** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   ![There was an update last week][Green]
* [drone-scp](https://github.com/appleboy/drone-scp) **star:55** Copy files and artifacts via SSH using a binary, docker or Drone CI.
* [Dropship](https://github.com/chrismckenzie/dropship) **star:46** Tool for deploying code via cdn.   ![It hasn't been updated in the last year][Yellow]
* [Rodent](https://github.com/alouche/rodent) **star:29** Rodent helps you manage Go versions, projects and track dependencies.   ![It hasn't been updated in the last year][Yellow]
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:24** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
* [awsenv](https://github.com/soniah/awsenv) **star:21** Small binary that loads Amazon (AWS) environment variables for a profile.   ![It hasn't been updated in the last year][Yellow]
* [DepCharge](https://github.com/centerorbit/depcharge) **star:9** Helps orchestrating the execution of commands across the many dependencies in larger projects.
* [lwc](https://github.com/timdp/lwc) **star:8** A live-updating version of the UNIX wc command.   ![It hasn't been updated in the last year][Yellow]
* [sg](https://github.com/ChristopherRabotin/sg) **star:5** Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.   ![It hasn't been updated in the last year][Yellow]

### Other Software

* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:8274** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [rkt](https://github.com/coreos/rkt)  App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM.
* [restic](https://github.com/restic/restic) **star:7518** De-duplicating backup program.   ![star > 5000][Gold]   ![There was an update last week][Green]
* [confd](https://github.com/kelseyhightower/confd) **star:6428** Manage local application configuration files using templates and data from etcd or consul.   ![star > 5000][Gold]
* [Comcast](https://github.com/tylertreat/Comcast) **star:6179** Simulate bad network connections.   ![star > 5000][Gold]
* [LiteIDE](https://github.com/visualfc/liteide) **star:5503** LiteIDE is a simple, open source, cross-platform Go IDE.   ![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [drive](https://github.com/odeke-em/drive) **star:4971** Google Drive client for the commandline.   ![star > 1000][Silver]
* [orange-cat](https://github.com/noraesae/orange-cat)  Markdown previewer written in Go.
* [nes](https://github.com/fogleman/nes) **star:4131** Nintendo Entertainment System (NES) emulator written in Go.   ![star > 1000][Silver]
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:3947** Proxy to simulate network and system conditions for automated tests.   ![star > 1000][Silver]
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [Pipe](https://github.com/b3log/pipe) **star:3000** A small and beautiful blogging platform.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:2695** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   ![star > 1000][Silver]
* [myLG](https://github.com/mehrdadrad/mylg) **star:2200** Command Line Network Diagnostic tool written in Go.   ![star > 1000][Silver]
* [GoBoy](https://github.com/Humpheh/goboy) **star:2108** Nintendo Game Boy Color emulator written in Go.   ![star > 1000][Silver]
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [Stack Up](https://github.com/pressly/sup) **star:1994** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   ![star > 1000][Silver]
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [lgo](https://github.com/yunabe/lgo) **star:1805** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   ![star > 1000][Silver]
* [snap](https://github.com/intelsdi-x/snap) **star:1803** Powerful telemetry framework.   ![star > 1000][Silver]
* [Circuit](https://github.com/gocircuit/circuit) **star:1786** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   ![star > 1000][Silver]
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:876** App that displays updates for the Go packages in your GOPATH.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Documize](https://github.com/documize/community) **star:831** Modern wiki software that integrates data from SaaS tools.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [scc](https://github.com/boyter/scc) **star:773** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [Leaps](https://github.com/jeffail/leaps) **star:639** Pair programming service using Operational Transforms.   ![star > 100][Bronze]
* [peg](https://github.com/pointlander/peg) **star:611** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   ![star > 100][Bronze]
* [vFlow](https://github.com/VerizonDigital/vflow) **star:600** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   ![star > 100][Bronze]
* [gfile](https://github.com/Antonito/gfile) **star:502** Securely transfer files between two computers, without any third party, over WebRTC.   ![star > 100][Bronze]
* [GoDNS](https://github.com/timothyye/godns) **star:436** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [shell2http](https://github.com/msoap/shell2http) **star:417** Executing shell commands via http server (for prototyping or remote control).   ![star > 100][Bronze]
* [mockingjay](https://github.com/quii/mockingjay-server) **star:414** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   ![star > 100][Bronze]
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:374** Video streaming torrent client.   ![star > 100][Bronze]
* [gocc](https://github.com/goccmack/gocc) **star:343** Gocc is a compiler kit for Go written in Go.   ![star > 100][Bronze]
* [wellington](https://github.com/wellington/wellington) **star:289** Sass project management tool, extends the language with sprite functions (like Compass).   ![star > 100][Bronze]
* [ipe](https://github.com/dimiro1/ipe) **star:276** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   ![star > 100][Bronze]
* [ide](https://github.com/thestrukture/ide) **star:252** Browser accessible IDE. Designed for Go with Go.   ![star > 100][Bronze]
* [Cherry](https://github.com/rafael-santiago/cherry) **star:192** Tiny webchat server in Go.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Orbit](https://github.com/gulien/orbit) **star:128** A simple tool for running commands and generating files from templates.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [joincap](https://github.com/assafmo/joincap) **star:122** Command-line utility for merging multiple pcap files together.   ![star > 100][Bronze]
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [DDNS](https://github.com/skibish/ddns) **star:98** Personal DDNS client with Digital Ocean Networking DNS as backend.
* [boxed](https://github.com/tejo/boxed) **star:72** Dropbox based blog engine.   ![It hasn't been updated in the last year][Yellow]
* [borg](https://github.com/crufter/borg)  Terminal based search engine for bash snippets.
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** Simple NaCL EC25519 based crypto pipe tool written in Go.
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   ![It hasn't been updated in the last year][Yellow]
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:12** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.   ![It hasn't been updated in the last year][Yellow]
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [Gor](https://github.com/buger/gor)  Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1260** Go HTTP request router benchmark and comparison.   ![star > 1000][Silver]
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1007** Go web framework benchmark.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [skynet](https://github.com/atemerev/skynet) **star:915** Skynet 1M threads microbenchmark.   ![star > 100][Bronze]
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:873** Benchmarks of Go serialization methods.   ![star > 100][Bronze]   ![There was an update last week][Green]
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Benchmarks of common basic operations for the Go language.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:172** Compare various Image resize algorithms for the Go language.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:123** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:93** Go micro-benchmarks for calculating the speed of language constructs.
* [autobench](https://github.com/davecheney/autobench) **star:89** Framework to compare the performance between different Go versions.   ![It hasn't been updated in the last year][Yellow]
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:48** Collection of benchmarks for popular Go database/SQL utilities.   ![It hasn't been updated in the last year][Yellow]
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:19** Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.   ![It hasn't been updated in the last year][Yellow]
* [kvbench](https://github.com/jimrobinson/kvbench) **star:14** Key/Value database benchmark.   ![It hasn't been updated in the last year][Yellow]

## Conferences

* [Capital Go](http://www.capitalgolang.com)  Washington, D.C., USA.
* [GopherCon Europe](https://gophercon.is/)  Reykjavik, Iceland.
* [GothamGo](http://gothamgo.com/)  New York City, USA.
* [GopherCon Vietnam](https://gophercon.vn/)  Ho Chi Minh City, Vietnam.
* [GopherCon Singapore](https://gophercon.sg)  Mapletree Business City, Singapore.
* [GopherCon Russia](https://www.gophercon-russia.ru)  Moscow, Russia.
* [GopherCon Israel](https://www.gophercon.org.il/)  Tel Aviv, Israel.
* [GopherCon India](https://www.gophercon.in/)  Pune, India.
* [GopherCon Brazil](https://gopherconbr.org)  Florianópolis, BR.
* [dotGo](http://www.dotgo.eu)  Paris, France.
* [GopherCon Australia](https://gophercon.com.au/)  Sydney, Australia.
* [GopherCon](http://www.gophercon.com/)  Denver, USA.
* [GopherChina](http://gopherchina.org)  Shanghai, China.
* [GolangUK](http://golanguk.com/)  London, UK.
* [GoLab](http://golab.io/)  Florence, Italy.
* [GoDays](https://www.godays.io/)  Berlin, Germany.
* [GoCon](http://gocon.connpass.com/)  Tokyo, Japan.
* [GoWayFest](https://goway.io/)  Minsk, Belarus.

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org)  A book focusing on Go syntax/semantics and all kinds of details.
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) **star:6812** A curated list of Go books.   ![star > 5000][Gold]
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:10** in Persian.
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1543** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   ![star > 1000][Silver]   ![There was an update last week][Green]
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:65** adorable gopher logos.   ![It hasn't been updated in the last year][Yellow]
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:31** Go gopher Vector Data [.ai, .svg].   ![It hasn't been updated in the last year][Yellow]
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:1859** Gopher artworks by Ashley McNamara.   ![star > 1000][Silver]
* [gophers](https://github.com/egonelbre/gophers) **star:1603** Free gophers.   ![star > 1000][Silver]
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:315** Gopherize yourself.   ![star > 100][Bronze]
* [gophers](https://github.com/rogeralsing/gophers) **star:50** random gopher graphics.   ![It hasn't been updated in the last year][Yellow]
* [gophers](https://github.com/sillecelik/go-gopher) **star:40** Gopher amigurumi toy pattern.

## Meetups

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

*Add the group of your city/country here (send **PR**)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## Websites

* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:24665** List of other amazingly awesome lists.   ![star > 5000][Gold]
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:14568** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.   ![star > 5000][Gold]
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:140** Collection of Go images, graphics, and art.   ![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:36** Collection of Go projects that needs help. Good place to start your open-source way in Go.   ![It hasn't been updated in the last year][Yellow]
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:31602** Golang ebook intro how to build a web app with golang.   ![star > 5000][Gold]   ![Contains Chinese documents][CN]
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:4044** Go's reference card.   ![star > 1000][Silver]
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:455** A little e-book on Ethereum Development with Go.   ![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:4575** Learn Go with test-driven development.   ![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:698** Examples of Golang compared to Node.js for learning.   ![star > 100][Bronze]
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1131** Intro to go for experienced programmers.   ![star > 1000][Silver]
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

