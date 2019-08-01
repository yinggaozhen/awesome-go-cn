# Awesome Go

[Gold]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Gold.svg "star > 5000"
[Silver]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Silver.svg "star > 1000"
[Bronze]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Bronze.svg "star > 100"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2019-08-01 21:18:49(Synchronize every day)**

[![english](https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/chinese.svg)](README.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) 为Awesome Go打赏~

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python)。

**Icon** :


Icon | State  
:-:|-
![star > 5000][Gold] | star  > 5000 
![star > 1000][Silver] | star > 1000
![star > 100][Bronze] | star > 100
![There was an update last week][Green] | There was an update last week。You can basically determine that the current library is in an actively maintained state。
![Not updated in last year][Yellow] | Not updated in last year。Reflects that the maintenance of the library is not high enthusiasm, use should be careful。
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
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
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

* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:22** EasyMidi is a simple and reliable library for working with standard midi file (SMF).   ![It hasn't been updated in the last year][Yellow]
* [flac](https://github.com/eaburns/flac) **star:83** No-frills native Go FLAC decoder that decodes FLAC files into byte slices.   ![It hasn't been updated in the last year][Yellow]
* [flac](https://github.com/mewkiz/flac) **star:102** Native Go FLAC encoder/decoder with support for FLAC streams.![star > 100][Bronze]
* [gaad](https://github.com/Comcast/gaad) **star:56** Native Go AAC bitstream parser.   ![It hasn't been updated in the last year][Yellow]
* [go-sox](https://github.com/krig/go-sox) **star:93** libsox bindings for go.   ![It hasn't been updated in the last year][Yellow]
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:24** libmediainfo bindings for go.   ![It hasn't been updated in the last year][Yellow]
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:8** libsamplerate bindings for go.   ![It hasn't been updated in the last year][Yellow]
* [id3v2](https://github.com/bogem/id3v2) **star:109** Fast and stable ID3 parsing and writing library for Go.![star > 100][Bronze]
* [malgo](https://github.com/gen2brain/malgo) **star:70** Mini audio library.
* [minimp3](https://github.com/tosone/minimp3) **star:25** Lightweight MP3 decoder library.
* [mix](https://github.com/go-mix/mix) **star:98** Sequence-based Go-native audio mixer for music apps.   ![It hasn't been updated in the last year][Yellow]
* [mp3](https://github.com/tcolgate/mp3) **star:89** Native Go MP3 decoder.   ![It hasn't been updated in the last year][Yellow]
* [music-theory](https://github.com/go-music-theory/music-theory) **star:253** Music theory models in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Oto](https://github.com/hajimehoshi/oto) **star:430** A low-level library to play sound on multiple platforms.![star > 100][Bronze]   ![There was an update last week][Green]
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:300** Go bindings for the PortAudio audio I/O library.![star > 100][Bronze]
* [portmidi](https://github.com/rakyll/portmidi) **star:205** Go bindings for PortMidi.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [taglib](https://github.com/wtolson/go-taglib) **star:66** Go bindings for taglib.   ![It hasn't been updated in the last year][Yellow]
* [vorbis](https://github.com/mccoyst/vorbis) **star:22** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).
* [waveform](https://github.com/mdlayher/waveform) **star:248** Go package capable of generating waveform images from audio streams.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [authboss](https://github.com/volatiletech/authboss) **star:1927** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.![star > 1000][Silver]
* [branca](https://github.com/hako/branca) **star:76** Golang implementation of Branca Tokens.
* [casbin](https://github.com/hsluoyz/casbin) **star:4907** Authorization library that supports access control models like ACL, RBAC, ABAC.![star > 1000][Silver]   ![There was an update last week][Green]
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** provides parser of cookies.txt file format.   ![It hasn't been updated in the last year][Yellow]
* [go-jose](https://github.com/square/go-jose) **star:1120** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.![star > 1000][Silver]
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1274** Standalone, specification-compliant,  OAuth2 server written in Golang.![star > 1000][Silver]
* [gologin](https://github.com/dghubble/gologin) **star:1039** chainable handlers for login with OAuth1 and OAuth2 authentication providers.![star > 1000][Silver]
* [gorbac](https://github.com/mikespook/gorbac) **star:906** provides a lightweight role-based access control (RBAC) implementation in Golang.![star > 100][Bronze]
* [goth](https://github.com/markbates/goth) **star:2259** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.![star > 1000][Silver]
* [httpauth](https://github.com/goji/httpauth) **star:178** HTTP Authentication middleware.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [jwt](https://github.com/robbert229/jwt) **star:70** Clean and easy to use implementation of JSON Web Tokens (JWT).
* [jwt](https://github.com/pascaldekloe/jwt) **star:91** Lightweight JSON Web Token (JWT) library.
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:151** JWT middleware for Golang http servers with many configuration options.![star > 100][Bronze]
* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:5939** Golang implementation of JSON Web Tokens (JWT).![star > 5000][Gold]   ![There was an update last week][Green]
* [loginsrv](https://github.com/tarent/loginsrv) **star:807** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.![star > 100][Bronze]
* [oauth2](https://github.com/golang/oauth2) **star:2376** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.![star > 1000][Silver]
* [osin](https://github.com/openshift/osin) **star:1539** Golang OAuth2 server library.![star > 1000][Silver]
* [paseto](https://github.com/o1egl/paseto) **star:241** Golang implementation of Platform-Agnostic Security Tokens (PASETO).![star > 100][Bronze]
* [permissions2](https://github.com/xyproto/permissions2) **star:352** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.![star > 100][Bronze]   ![There was an update last week][Green]
* [rbac](https://github.com/zpatrick/rbac) **star:27** Minimalistic RBAC package for Go applications.
* [scs](https://github.com/alexedwards/scs) **star:523** Session Manager for HTTP servers.![star > 100][Bronze]
* [securecookie](https://github.com/chmike/securecookie) **star:32** Efficient secure cookie encoding/decoding.
* [session](https://github.com/icza/session) **star:89** Go session management for web servers (including support for Google App Engine - GAE).
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.
* [sessions](https://github.com/adam-hanna/sessions) **star:47** Dead simple, highly performant, highly customizable sessions service for go http servers.
* [signedvalue](https://github.com/sashka/signedvalue) **star:7** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.
* [sjwt](https://github.com/brianvoe/sjwt) **star:31** Simple jwt generator and parser.

## Bot Building

*Libraries for building and working with bots.*

* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:468** IRC, Slack & Telegram bot written in Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:137** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.![star > 100][Bronze]
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:85** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   ![It hasn't been updated in the last year][Yellow]
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:220** A golang implementation of a console-based trading bot for cryptocurrency exchanges.![star > 100][Bronze]
* [govkbot](https://github.com/nikepan/govkbot) **star:24** Simple Go [VK](https://vk.com) bot library.
* [hanu](https://github.com/sbstjn/hanu) **star:108** Framework for writing Slack bots.![star > 100][Bronze]
* [Kelp](https://github.com/stellar/kelp) **star:157** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.![star > 100][Bronze]   ![There was an update last week][Green]
* [margelet](https://github.com/zhulik/margelet) **star:56** Framework for building Telegram bots.   ![It hasn't been updated in the last year][Yellow]
* [micha](https://github.com/onrik/micha) **star:10** Go Library for Telegram bot api.   ![It hasn't been updated in the last year][Yellow]
* [slacker](https://github.com/shomali11/slacker) **star:306** Easy to use framework to create Slack bots.![star > 100][Bronze]
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:10** Another framework for building Slack bots.
* [tbot](https://github.com/yanzay/tbot) **star:218** Telegram bot server with API similar to net/http.![star > 100][Bronze]   ![There was an update last week][Green]
* [telebot](https://github.com/tucnak/telebot) **star:945** Telegram bot framework written in Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1620** Simple and clean Telegram bot client.![star > 1000][Silver]   ![There was an update last week][Green]
* [Tenyks](https://github.com/kyleterry/tenyks) **star:167** Service oriented IRC bot using Redis and JSON for messaging.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [argparse](https://github.com/akamensky/argparse) **star:105** Command line argument parser inspired by Python's argparse module.![star > 100][Bronze]
* [argv](https://github.com/cosiner/argv) **star:17** Go library to split command line string as arguments array using the bash syntax.
* [cli](https://github.com/mkideal/cli) **star:475** Feature-rich and easy to use command-line package based on golang struct tags.![star > 100][Bronze]
* [cli](https://github.com/teris-io/cli) **star:57** Simple and complete API for building command line interfaces in Go.
* [cli-init](https://github.com/tcnksm/gcli) **star:868** The easy way to start building Golang command line applications.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [cmdr](https://github.com/hedzr/cmdr) **star:12** A POSIX/GNU style, getopt-like command-line UI Go library.   ![There was an update last week][Green]
* [cobra](https://github.com/spf13/cobra) **star:13224** Commander for modern Go CLI interactions.![star > 5000][Gold]   ![There was an update last week][Green]
* [commandeer](https://github.com/jaffee/commandeer) **star:90** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   ![There was an update last week][Green]
* [complete](https://github.com/posener/complete) **star:611** Write bash completions in Go + Go command bash completion.![star > 100][Bronze]
* [docopt.go](https://github.com/docopt/docopt.go) **star:1172** Command-line arguments parser that will make you smile.![star > 1000][Silver]
* [env](https://github.com/codingconcepts/env) **star:42** Tag-based environment configuration for structs.
* [flag](https://github.com/cosiner/flag) **star:100** Simple but powerful command line option parsing library for Go supporting subcommand.![star > 100][Bronze]
* [flaggy](https://github.com/integrii/flaggy) **star:448** A robust and idiomatic flags package with excellent subcommand support.![star > 100][Bronze]
* [flagvar](https://github.com/sgreben/flagvar) **star:31** A collection of flag argument types for Go's standard `flag` package.
* [go-arg](https://github.com/alexflint/go-arg) **star:656** Struct-based argument parsing in Go.![star > 100][Bronze]
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.
* [go-flags](https://github.com/jessevdk/go-flags) **star:1504** go command line option parser.![star > 1000][Silver]
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:7** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   ![There was an update last week][Green]
* [gocmd](https://github.com/devfacet/gocmd) **star:33** Go library for building command line applications.
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [job](https://github.com/liujianping/job) **star:49** JOB, make your short-term command as a long-term job.
* [kingpin](https://github.com/alecthomas/kingpin) **star:2526** Command line and flag parser supporting sub commands.![star > 1000][Silver]
* [liner](https://github.com/peterh/liner) **star:582** Go readline-like library for command-line interfaces.![star > 100][Bronze]
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:996** Go library for implementing command-line interfaces.![star > 100][Bronze]
* [mow.cli](https://github.com/jawher/mow.cli) **star:623** Go library for building CLI applications with sophisticated flag and argument parsing and validation.![star > 100][Bronze]
* [ops](https://github.com/nanovms/ops) **star:261** Unikernel Builder/Orchestrator.![star > 100][Bronze]
* [pflag](https://github.com/spf13/pflag) **star:748** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.![star > 100][Bronze]
* [readline](https://github.com/chzyer/readline) **star:1366** Pure golang implementation that provides most features in GNU-Readline under MIT license.![star > 1000][Silver]
* [sand](https://github.com/Zaba505/sand) **star:5** Simple API for creating interpreters and so much more.
* [sflags](https://github.com/octago/sflags) **star:89** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   ![There was an update last week][Green]
* [strumt](https://github.com/antham/strumt) **star:27** Library to create prompt chain.
* [ts](https://github.com/liujianping/ts) **star:4** Timestamp convert & compare tool.
* [ukautz/clif](https://github.com/ukautz/clif) **star:97** Small command line interface framework.
* [urfave/cli](https://github.com/urfave/cli) **star:11322** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).![star > 5000][Gold]   ![There was an update last week][Green]
* [wlog](https://github.com/dixonwille/wlog) **star:35** Simple logging interface that supports cross-platform color and concurrency.   ![It hasn't been updated in the last year][Yellow]
* [wmenu](https://github.com/dixonwille/wmenu) **star:84** Easy to use menu structure for cli applications that prompts users to make choices.

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1137** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.![star > 1000][Silver]
* [aurora](https://github.com/logrusorgru/aurora) **star:627** ANSI terminal colors that supports fmt.Printf/Sprintf.![star > 100][Bronze]
* [cfmt](https://github.com/mingrammer/cfmt) **star:67** Contextual fmt inspired by bootstrap color classes.
* [chalk](https://github.com/ttacon/chalk) **star:304** Intuitive package for prettifying terminal/console output.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [color](https://github.com/fatih/color) **star:3009** Versatile package for colored terminal output.![star > 1000][Silver]
* [colourize](https://github.com/TreyBastian/colourize) **star:16** Go library for ANSI colour text in terminals.   ![It hasn't been updated in the last year][Yellow]
* [ctc](https://github.com/wzshiming/ctc) **star:9** The non-invasive cross-platform terminal color library does not need to modify the Print method.   ![Contains Chinese documents][CN]
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** Go library for rendering ANSI colored text templates in terminals.   ![It hasn't been updated in the last year][Yellow]
* [go-colorable](https://github.com/mattn/go-colorable) **star:373** Colorable writer for windows.![star > 100][Bronze]
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:197** Go library for color output in terminals.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-isatty](https://github.com/mattn/go-isatty) **star:344** isatty for golang.![star > 100][Bronze]
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2324** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).![star > 1000][Silver]   ![There was an update last week][Green]
* [gocui](https://github.com/jroimartin/gocui) **star:5332** Minimalist Go library aimed at creating Console User Interfaces.![star > 5000][Gold]
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [gookit/color](https://github.com/gookit/color) **star:199** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [mpb](https://github.com/vbauerster/mpb) **star:680** Multi progress bar for terminal applications.![star > 100][Bronze]   ![There was an update last week][Green]
* [progressbar](https://github.com/schollz/progressbar) **star:567** Basic thread-safe progress bar that works in every OS.![star > 100][Bronze]
* [simpletable](https://github.com/alexeyco/simpletable) **star:159** Simple tables in terminal with Go.![star > 100][Bronze]
* [tabby](https://github.com/cheynewallace/tabby) **star:246** A tiny library for super simple Golang tables.![star > 100][Bronze]
* [tabular](https://github.com/InVisionApp/tabular) **star:29** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   ![It hasn't been updated in the last year][Yellow]
* [termbox-go](https://github.com/nsf/termbox-go) **star:3471** Termbox is a library for creating cross-platform text-based interfaces.![star > 1000][Silver]
* [termdash](https://github.com/mum4k/termdash) **star:804** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).![star > 100][Bronze]
* [termtables](https://github.com/apcera/termtables) **star:214** Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [termui](https://github.com/gizak/termui) **star:8891** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).![star > 5000][Gold]   ![There was an update last week][Green]
* [uilive](https://github.com/gosuri/uilive) **star:831** Library for updating terminal output in realtime.![star > 100][Bronze]
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1533** Flexible library to render progress bars in terminal applications.![star > 1000][Silver]
* [uitable](https://github.com/gosuri/uitable) **star:502** Library to improve readability in terminal apps using tabular data.![star > 100][Bronze]

## Configuration

*Libraries for configuration parsing.*

* [config](https://github.com/JeremyLoy/config) **star:189** Cloud native application configuration. Bind ENV to structs in only two lines.![star > 100][Bronze]
* [config](https://github.com/olebedev/config) **star:210** JSON or YAML configuration wrapper with environment variables and flags parsing.![star > 100][Bronze]
* [configure](https://github.com/paked/configure) **star:48** Provides configuration through multiple sources, including JSON, flags and environment variables.
* [confita](https://github.com/heetch/confita) **star:238** Load configuration in cascade from multiple backends into a struct.![star > 100][Bronze]
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:8** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
* [env](https://github.com/caarlos0/env) **star:857** Parse environment variables to Go structs (with defaults).![star > 100][Bronze]
* [envcfg](https://github.com/tomazk/envcfg) **star:89** Un-marshaling environment variables to Go structs.   ![It hasn't been updated in the last year][Yellow]
* [envconf](https://github.com/ian-kent/envconf) **star:7** Configuration from environment.   ![It hasn't been updated in the last year][Yellow]
* [envconfig](https://github.com/vrischmann/envconfig) **star:145** Read your configuration from environment variables.![star > 100][Bronze]
* [envh](https://github.com/antham/envh) **star:92** Helpers to manage environment variables.
* [gcfg](https://github.com/go-gcfg/gcfg) **star:116** read INI-style configuration files into Go structs; supports user-defined types and subsections.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-up](https://github.com/ufoscout/go-up) **star:24** A simple configuration library with recursive placeholders resolution and no magic.
* [goConfig](https://github.com/crgimenes/goConfig) **star:109** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.![star > 100][Bronze]
* [godotenv](https://github.com/joho/godotenv) **star:2114** Go port of Ruby's dotenv library (Loads environment variables from `.env`).![star > 1000][Silver]
* [gofigure](https://github.com/ian-kent/gofigure) **star:57** Go application configuration made easy.   ![It hasn't been updated in the last year][Yellow]
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gookit/config](https://github.com/gookit/config) **star:81** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [harvester](https://github.com/beatlabs/harvester) **star:39** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.   ![There was an update last week][Green]
* [hjson](https://github.com/hjson/hjson-go) **star:171** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.![star > 100][Bronze]
* [ingo](https://github.com/schachmat/ingo) **star:24** Flags persisted in an ini-like config file.   ![It hasn't been updated in the last year][Yellow]
* [ini](https://github.com/go-ini/ini) **star:1587** Go package to read and write INI files.![star > 1000][Silver]
* [joshbetz/config](https://github.com/joshbetz/config) **star:193** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2400** Go library for managing configuration data from environment variables.![star > 1000][Silver]
* [koanf](https://github.com/knadh/koanf) **star:83** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.
* [konfig](https://github.com/lalamove/konfig) **star:513** Composable, observable and performant config handling for Go for the distributed processing era.![star > 100][Bronze]
* [mini](https://github.com/sasbury/mini) **star:19** Golang package for parsing ini-style configuration files.
* [sprbox](https://github.com/oblq/sprbox) **star:3** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).   ![There was an update last week][Green]
* [store](https://github.com/tucnak/store) **star:241** Lightweight configuration manager for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [viper](https://github.com/spf13/viper) **star:9260** Go configuration with fangs.![star > 5000][Gold]   ![There was an update last week][Green]
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:35** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:18957** Drone is a Continuous Integration platform built on Docker, written in Go.![star > 5000][Gold]   ![There was an update last week][Green]
* [duci](https://github.com/duck8823/duci) **star:44** A simple ci server no needs domain specific languages.   ![There was an update last week][Green]
* [gomason](https://github.com/nikogura/gomason) **star:30** Test, Build, Sign, and Publish your go binaries from a clean workspace.
* [goveralls](https://github.com/mattn/goveralls) **star:579** Go integration for Coveralls.io continuous code coverage tracking system.![star > 100][Bronze]
* [overalls](https://github.com/go-playground/overalls) **star:98** Multi-Package go project coverprofile for tools like goveralls.
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** Recursive coverage testing tool.   ![It hasn't been updated in the last year][Yellow]

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:423** Pure Go CSS Preprocessor.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-libsass](https://github.com/wellington/go-libsass) **star:130** Go wrapper to the 100% Sass compatible libsass project.![star > 100][Bronze]

## Data Structures

*Generic datastructures and algorithms in Go.*

* [algorithms](https://github.com/shady831213/algorithms) **star:240** Algorithms and data structures.CLRS study.![star > 100][Bronze]
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:123** Binary packer and unpacker helps user build custom binary stream.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [bit](https://github.com/yourbasic/bit) **star:55** Golang set data structure with bonus bit-twiddling functions.   ![It hasn't been updated in the last year][Yellow]
* [bitset](https://github.com/willf/bitset) **star:479** Go package implementing bitsets.![star > 100][Bronze]
* [bloom](https://github.com/zhenjl/bloom) **star:128** Bloom filters implemented in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [bloom](https://github.com/yourbasic/bloom) **star:39** Golang Bloom filter implementation.   ![It hasn't been updated in the last year][Yellow]
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1132** Probabilistic data structures for processing continuous, unbounded streams.![star > 1000][Silver]
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:24** Highly concurrent drop-in replacement for `bufio.Writer`.   ![It hasn't been updated in the last year][Yellow]
* [conjungo](https://github.com/InVisionApp/conjungo) **star:76** A small, powerful and flexible merge library.
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:43** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   ![It hasn't been updated in the last year][Yellow]
* [crunch](https://github.com/superwhiskers/crunch) **star:19** Go package implementing buffers for handling various datatypes easily.
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:511** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.![star > 100][Bronze]
* [deque](https://github.com/edwingeng/deque) **star:7** A highly optimized double-ended queue.
* [deque](https://github.com/gammazero/deque) **star:64** Fast ring-buffer deque (double-ended queue).
* [dict](https://github.com/srfrog/dict) **star:9** Python-like dictionaries (dict) for Go.
* [encoding](https://github.com/zhenjl/encoding) **star:94** Integer Compression Libraries for Go.   ![It hasn't been updated in the last year][Yellow]
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:88** Go implementation of Adaptive Radix Tree.
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5124** Collection of useful, performant, and thread-safe data structures.![star > 5000][Gold]
* [go-ef](https://github.com/amallia/go-ef) **star:11** A Go implementation of the Elias-Fano encoding.   ![It hasn't been updated in the last year][Yellow]
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:311** In-memory geo index.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:34** Fast in-memory key:value store/cache library. Pointer caches.
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:99** Region quadtrees with efficient point location and neighbour finding.   ![It hasn't been updated in the last year][Yellow]
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:26** Concurrent FIFO queue.   ![There was an update last week][Green]
* [gods](https://github.com/emirpasic/gods) **star:6327** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.![star > 5000][Gold]
* [golang-set](https://github.com/deckarep/golang-set) **star:1160** Thread-Safe and Non-Thread-Safe high-performance sets for Go.![star > 1000][Silver]
* [goset](https://github.com/zoumo/goset) **star:16** A useful Set collection implementation for Go.   ![It hasn't been updated in the last year][Yellow]
* [goskiplist](https://github.com/ryszard/goskiplist) **star:193** Skip list implementation in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gota](https://github.com/kniren/gota) **star:875** Implementation of dataframes, series, and data wrangling methods for Go.![star > 100][Bronze]
* [hide](https://github.com/emvi/hide) **star:7** ID type with marshalling to/from hash to prevent sending IDs to clients.
* [hilbert](https://github.com/google/hilbert) **star:181** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.![star > 100][Bronze]
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:660** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.![star > 100][Bronze]
* [levenshtein](https://github.com/agext/levenshtein) **star:32** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [levenshtein](https://github.com/agnivade/levenshtein) **star:55** Implementation to calculate levenshtein distance in Go.
* [mafsa](https://github.com/smartystreets/mafsa) **star:273** MA-FSA implementation with Minimal Perfect Hashing.![star > 100][Bronze]
* [merkletree](https://github.com/cbergoon/merkletree) **star:144** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.![star > 100][Bronze]
* [mspm](https://github.com/BlackRabbitt/mspm) **star:7** Multi-String Pattern Matching Algorithm for information retrieval.   ![It hasn't been updated in the last year][Yellow]
* [null](https://github.com/emvi/null) **star:5** Nullable Go types that can be marshalled/unmarshalled to/from JSON.
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.
* [pipeline](https://github.com/hyfather/pipeline) **star:15** An implementation of pipelines with fan-in and fan-out.
* [ptrie](https://github.com/viant/ptrie)  An implementation of prefix tree.
* [ring](https://github.com/TheTannerRyan/ring) **star:86** Go implementation of a high performance, thread safe bloom filter.
* [roaring](https://github.com/RoaringBitmap/roaring) **star:655** Go package implementing compressed bitsets.![star > 100][Bronze]   ![There was an update last week][Green]
* [set](https://github.com/StudioSol/set) **star:6** Simple set data structure implementation in Go using LinkedHashMap.
* [skiplist](https://github.com/MauriceGit/skiplist) **star:100** Very fast Go Skiplist implementation.![star > 100][Bronze]
* [skiplist](https://github.com/gansidui/skiplist) **star:64** Skiplist implementation in Go.   ![It hasn't been updated in the last year][Yellow]
* [timedmap](https://github.com/zekroTJA/timedmap) **star:1** Map with expiring key-value pairs.
* [treap](https://github.com/perdata/treap) **star:7** Persistent, fast ordered map using tree heaps.
* [trie](https://github.com/derekparker/trie) **star:420** Trie implementation in Go.![star > 100][Bronze]
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:96** In-memory LRU string-interface{} map with expiration for golang.
* [typ](https://github.com/gurukami/typ) **star:9** Null Types, Safe primitive type conversion and fetching value from complex structures.
* [willf/bloom](https://github.com/willf/bloom) **star:665** Go package implementing Bloom filters.![star > 100][Bronze]

## Database

*Databases implemented in Go.*

* [badger](https://github.com/dgraph-io/badger) **star:6250** Fast key-value store in Go.![star > 5000][Gold]   ![There was an update last week][Green]
* [bcache](https://github.com/iwanbk/bcache) **star:28** Eventually consistent distributed in-memory  cache Go library.
* [BigCache](https://github.com/allegro/bigcache) **star:2441** Efficient key/value cache for gigabytes of data.![star > 1000][Silver]
* [bolt](https://github.com/boltdb/bolt) **star:9957** Low-level key/value database for Go.![star > 5000][Gold]   ![It hasn't been updated in the last year][Yellow]
* [buntdb](https://github.com/tidwall/buntdb) **star:2433** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.![star > 1000][Silver]
* [cache](https://github.com/akyoto/cache) **star:15** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.
* [cache2go](https://github.com/muesli/cache2go) **star:1028** In-memory key:value cache which supports automatic invalidation based on timeouts.![star > 1000][Silver]   ![There was an update last week][Green]
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:29** BigCache with clustering support and individual item expiration.   ![It hasn't been updated in the last year][Yellow]
* [cockroach](https://github.com/cockroachdb/cockroach) **star:16742** Scalable, Geo-Replicated, Transactional Datastore.![star > 5000][Gold]   ![There was an update last week][Green]
* [couchcache](https://github.com/codingsince1985/couchcache) **star:40** RESTful caching micro-service backed by Couchbase server.
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:872** CovenantSQL is a SQL database on blockchain.![star > 100][Bronze]   ![There was an update last week][Green]
* [dgraph](https://github.com/dgraph-io/dgraph) **star:10189** Scalable, Distributed, Low Latency, High Throughput Graph Database.![star > 5000][Gold]   ![There was an update last week][Green]
* [diskv](https://github.com/peterbourgon/diskv) **star:745** Home-grown disk-backed key-value store.![star > 100][Bronze]
* [eliasdb](https://github.com/krotik/eliasdb) **star:532** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.![star > 100][Bronze]
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:483** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.![star > 100][Bronze]
* [GCache](https://github.com/bluele/gcache) **star:893** Cache library with support for expirable Cache, LFU, LRU and ARC.![star > 100][Bronze]
* [go-cache](https://github.com/pmylund/go-cache) **star:2886** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.![star > 1000][Silver]
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3164** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.![star > 1000][Silver]
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:8** Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.   ![It hasn't been updated in the last year][Yellow]
* [groupcache](https://github.com/golang/groupcache) **star:7643** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.![star > 5000][Gold]
* [influxdb](https://github.com/influxdb/influxdb) **star:17063** Scalable datastore for metrics, events, and real-time analytics.![star > 5000][Gold]   ![There was an update last week][Green]
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3065** Ledisdb is a high performance NoSQL like Redis based on LevelDB.![star > 1000][Silver]
* [levigo](https://github.com/jmhodges/levigo) **star:364** Levigo is a Go wrapper for LevelDB.![star > 100][Bronze]
* [moss](https://github.com/couchbase/moss) **star:716** Moss is a simple LSM key-value storage engine written in 100% Go.![star > 100][Bronze]
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:885** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [piladb](https://github.com/fern4lvarez/piladb) **star:171** Lightweight RESTful database engine based on stack data structures.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [prometheus](https://github.com/prometheus/prometheus) **star:25446** Monitoring system and time series database.![star > 5000][Gold]   ![There was an update last week][Green]
* [pudge](https://github.com/recoilme/pudge) **star:218** Fast and simple  key/value store written using Go's standard library.![star > 100][Bronze]
* [rqlite](https://github.com/rqlite/rqlite) **star:4687** The lightweight, distributed, relational database built on SQLite.![star > 1000][Silver]   ![There was an update last week][Green]
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:59** Tiny flat file JSON store.
* [slowpoke](https://github.com/recoilme/slowpoke) **star:86** Key-value store with persistence.
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:13** Key-value store for temporary items.   ![It hasn't been updated in the last year][Yellow]
* [tidb](https://github.com/pingcap/tidb) **star:19978** TiDB is a distributed SQL database. Inspired by the design of Google F1.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2364** Your NoSQL database powered by Golang.![star > 1000][Silver]
* [Vasto](https://github.com/chrislusf/vasto) **star:146** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.![star > 100][Bronze]
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:987** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.![star > 100][Bronze]   ![There was an update last week][Green]

*Database schema migration.*

* [avro](https://github.com/khezen/avro) **star:5** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   ![There was an update last week][Green]
* [darwin](https://github.com/GuiaBolso/darwin) **star:83** Database schema evolution library for Go.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:20** Django style fixtures for Golang's excellent built-in database/sql library.
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:23** A Go package to help write migrations with go-pg/pg.
* [gondolier](https://github.com/emvi/gondolier) **star:26** Database migration library using struct decorators.
* [goose](https://github.com/steinbacher/goose) **star:119** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:327** Database schema migration helper for Gorm ORM.![star > 100][Bronze]
* [migrate](https://github.com/golang-migrate/migrate) **star:2590** Database migrations. CLI and Golang library.![star > 1000][Silver]   ![There was an update last week][Green]
* [migrator](https://github.com/lopezator/migrator) **star:31** Dead simple Go database migration library.   ![There was an update last week][Green]
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1404** Database migration tool. Allows embedding migrations into the application using go-bindata.![star > 1000][Silver]

*Database tools.*

* [chproxy](https://github.com/Vertamedia/chproxy) **star:303** HTTP proxy for ClickHouse database.![star > 100][Bronze]
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:135** Collects small insterts and sends big requests to ClickHouse servers.![star > 100][Bronze]
* [datagen](https://github.com/codingconcepts/datagen) **star:7** A fast data generator that's multi-table aware and supports multi-row DML.
* [dbbench](https://github.com/sj14/dbbench) **star:30** Database benchmarking tool with support for several databases and scripts.
* [go-mysql](https://github.com/siddontang/go-mysql) **star:1873** Go toolset to handle MySQL protocol and replication.![star > 1000][Silver]
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2383** Sync your MySQL data into Elasticsearch automatically.![star > 1000][Silver]
* [kingshard](https://github.com/flike/kingshard) **star:4593** kingshard is a high performance proxy for MySQL powered by Golang.![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [myreplication](https://github.com/2tvenom/myreplication) **star:141** MySql binary log replication listener. Supports statement and row based replication.![star > 100][Bronze]
* [octillery](https://github.com/knocknote/octillery) **star:53** Go package for sharding databases ( Supports every ORM or raw SQL ).
* [orchestrator](https://github.com/github/orchestrator) **star:3019** MySQL replication topology manager & visualizer.![star > 1000][Silver]   ![There was an update last week][Green]
* [pgweb](https://github.com/sosedoff/pgweb) **star:5975** Web-based PostgreSQL database browser.![star > 5000][Gold]   ![There was an update last week][Green]
* [prep](https://github.com/hexdigest/prep) **star:24** Use prepared SQL statements without changing your code.   ![It hasn't been updated in the last year][Yellow]
* [pREST](https://github.com/nuveo/prest) **star:2082** Serve a RESTful API from any PostgreSQL database.![star > 1000][Silver]
* [rwdb](https://github.com/andizzle/rwdb) **star:10** rwdb provides read replica capability for multiple database servers setup.   ![It hasn't been updated in the last year][Yellow]
* [vitess](https://github.com/youtube/vitess) **star:8427** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.![star > 5000][Gold]   ![There was an update last week][Green]

*SQL query builder, libraries for building and using SQL.*

* [Dotsql](https://github.com/gchaincl/dotsql) **star:437** Go library that helps you keep sql files in one place and use them with ease.![star > 100][Bronze]
* [gendry](https://github.com/didi/gendry) **star:754** Non-invasive SQL builder and powerful data binder.![star > 100][Bronze]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [godbal](https://github.com/xujiajun/godbal) **star:50** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.
* [goqu](https://github.com/doug-martin/goqu) **star:632** Idiomatic SQL builder and query library.![star > 100][Bronze]   ![There was an update last week][Green]
* [igor](https://github.com/galeone/igor) **star:78** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   ![It hasn't been updated in the last year][Yellow]
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:435** Powerful data retrieval methods as well as DB-agnostic query building capabilities.![star > 100][Bronze]
* [scaneo](https://github.com/variadico/scaneo) **star:149** Generate Go code to convert database rows into arbitrary structs.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [sqrl](https://github.com/elgris/sqrl) **star:177** SQL query builder, fork of Squirrel with improved performance.![star > 100][Bronze]
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [Squirrel](https://github.com/Masterminds/squirrel) **star:2299** Go library that helps you build SQL queries.![star > 1000][Silver]   ![There was an update last week][Green]
* [xo](https://github.com/knq/xo) **star:2176** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.![star > 1000][Silver]

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:34** Apache Avatica/Phoenix SQL driver for database/sql.
    * [bgc](https://github.com/viant/bgc) **star:12** Datastore Connectivity for BigQuery for go.
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:103** Firebird RDBMS SQL driver for Go.![star > 100][Bronze]
    * [go-adodb](https://github.com/mattn/go-adodb) **star:91** Microsoft ActiveX Object DataBase driver for go that uses database/sql.
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1017** Microsoft MSSQL driver for Go.![star > 1000][Silver]   ![There was an update last week][Green]
    * [go-oci8](https://github.com/mattn/go-oci8) **star:404** Oracle driver for go that uses database/sql.![star > 100][Bronze]
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:8076** MySQL driver for Go.![star > 5000][Gold]   ![There was an update last week][Green]
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3422** SQLite3 driver for go that uses database/sql.![star > 1000][Silver]
    * [gofreetds](https://github.com/minus5/gofreetds) **star:90** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [goracle](https://github.com/go-goracle/goracle) **star:235** Oracle driver for Go, using the ODPI-C driver.![star > 100][Bronze]   ![There was an update last week][Green]
    * [pgx](https://github.com/jackc/pgx) **star:1897** PostgreSQL driver supporting features beyond those exposed by database/sql.![star > 1000][Silver]   ![There was an update last week][Green]
    * [pq](https://github.com/lib/pq) **star:5173** Pure Go Postgres driver for database/sql.![star > 5000][Gold]   ![There was an update last week][Green]

* NoSQL Databases
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:304** Aerospike client in Go language.![star > 100][Bronze]
    * [arangolite](https://github.com/solher/arangolite) **star:65** Lightweight golang driver for ArangoDB.
    * [asc](https://github.com/viant/asc) **star:4** Datastore Connectivity for Aerospike for go.
    * [dynago](https://github.com/underarmour/dynago) **star:64** Dynago is a principle of least surprise client for DynamoDB.   ![It hasn't been updated in the last year][Yellow]
    * [forestdb](https://github.com/couchbase/goforestdb) **star:29** Go bindings for ForestDB.   ![It hasn't been updated in the last year][Yellow]
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:292** Couchbase client in Go.![star > 100][Bronze]
    * [go-couchdb](https://github.com/fjl/go-couchdb) **star:51** Yet another CouchDB HTTP API wrapper for Go.
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:31** Go client library for Pilosa.
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:90** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   ![There was an update last week][Green]
    * [gocb](https://github.com/couchbase/gocb) **star:290** Official Couchbase Go SDK.![star > 100][Bronze]   ![There was an update last week][Green]
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [godis](https://github.com/piaohao/godis) **star:60** redis client implement by golang, inspired by jedis.
    * [godscache](https://github.com/defcronyke/godscache) **star:6** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [gorethink](https://github.com/dancannon/gorethink) **star:1456** Go language driver for RethinkDB.![star > 1000][Silver]
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.
    * [mgo](https://github.com/globalsign/mgo) **star:1629** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.![star > 1000][Silver]   ![There was an update last week][Green]
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:3071** Official MongoDB driver for the Go language.![star > 1000][Silver]   ![There was an update last week][Green]
    * [neo4j](https://github.com/cihangir/neo4j) **star:24** Neo4j Rest API Bindings for Golang.   ![It hasn't been updated in the last year][Yellow]
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:72** Neo4j REST Client in golang.   ![It hasn't been updated in the last year][Yellow]
    * [neoism](https://github.com/jmcvetta/neoism) **star:357** Neo4j client for Golang.![star > 100][Bronze]
    * [redeo](https://github.com/bsm/redeo) **star:259** Redis-protocol compatible TCP servers/services.![star > 100][Bronze]
    * [redigo](https://github.com/gomodule/redigo) **star:6252** Redigo is a Go client for the Redis database.![star > 5000][Gold]   ![There was an update last week][Green]
    * [redis](https://github.com/go-redis/redis) **star:6465** Redis client for Golang.![star > 5000][Gold]   ![There was an update last week][Green]
    * [xredis](https://github.com/shomali11/xredis) **star:9** Typesafe, customizable, clean & easy to use Redis client.

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:5821** Modern text indexing library for go.![star > 5000][Gold]   ![There was an update last week][Green]
    * [elastic](https://github.com/olivere/elastic) **star:4117** Elasticsearch client for Go.![star > 1000][Silver]   ![There was an update last week][Green]
    * [elasticsql](https://github.com/cch123/elasticsql) **star:391** Convert sql to elasticsearch dsl in Go.![star > 100][Bronze]
    * [elastigo](https://github.com/mattbaird/elastigo) **star:951** Elasticsearch client library.![star > 100][Bronze]
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:1557** Official Elasticsearch client for Go.![star > 1000][Silver]   ![There was an update last week][Green]
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   ![It hasn't been updated in the last year][Yellow]
    * [riot](https://github.com/go-ego/riot) **star:4681** Go Open Source, Distributed, Simple and efficient Search Engine.![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
    * [skizze](https://github.com/seiflotfy/skizze) **star:68** probabilistic data-structures service and storage.   ![It hasn't been updated in the last year][Yellow]

* Multiple Backends.
    * [cachego](https://github.com/fabiorphp/cachego) **star:111** Golang Cache component for multiple drivers.![star > 100][Bronze]
    * [cayley](https://github.com/google/cayley) **star:12673** Graph database with support for multiple backends.![star > 5000][Gold]   ![There was an update last week][Green]
    * [dsc](https://github.com/viant/dsc) **star:13** Datastore connectivity for SQL, NoSQL, structured files.
    * [gokv](https://github.com/philippgille/gokv) **star:81** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).

## Date and Time

*Libraries for working with dates and times.*

* [carbon](https://github.com/uniplaces/carbon) **star:334** Simple Time extension with a lot of util methods, ported from PHP Carbon library.![star > 100][Bronze]
* [date](https://github.com/rickb777/date) **star:27** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
* [dateparse](https://github.com/araddon/dateparse) **star:891** Parse date's without knowing format in advance.![star > 100][Bronze]
* [durafmt](https://github.com/hako/durafmt) **star:239** Time duration formatting library for Go.![star > 100][Bronze]
* [feiertage](https://github.com/wlbr/feiertage) **star:22** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:64** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:13** Calculate the sunrise and sunset times for a given location.   ![It hasn't been updated in the last year][Yellow]
* [goweek](https://github.com/grsmv/goweek) **star:18** Library for working with week entity in golang.   ![It hasn't been updated in the last year][Yellow]
* [iso8601](https://github.com/relvacode/iso8601) **star:68** Efficiently parse ISO8601 date-times without regex.
* [kair](https://github.com/GuilhermeCaruso/kair) **star:10** Date and Time - Golang Formatting Library.
* [now](https://github.com/jinzhu/now) **star:2172** Now is a time toolkit for golang.![star > 1000][Silver]
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:9** Nullable `time.Time`.   ![It hasn't been updated in the last year][Yellow]
* [strftime](https://github.com/awoodbeck/strftime) **star:5** C99-compatible strftime formatter.   ![It hasn't been updated in the last year][Yellow]
* [timespan](https://github.com/SaidinWoT/timespan) **star:60** For interacting with intervals of time, defined as a start time and a duration.
* [timeutil](https://github.com/leekchan/timeutil) **star:168** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.![star > 100][Bronze]
* [tuesday](https://github.com/osteele/tuesday) **star:7** Ruby-compatible Strftime function.   ![It hasn't been updated in the last year][Yellow]

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:52** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
* [consistent](https://github.com/buraksezer/consistent) **star:188** Consistent hashing with bounded loads.![star > 100][Bronze]
* [dht](https://github.com/anacrolix/dht) **star:128** BitTorrent Kademlia DHT implementation.![star > 100][Bronze]
* [digota](https://github.com/digota/digota) **star:298** grpc ecommerce microservice.![star > 100][Bronze]
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [doublejump](https://github.com/edwingeng/doublejump) **star:39** A revamped Google's jump consistent hash.
* [dragonboat](https://github.com/lni/dragonboat) **star:2515** A feature complete and high performance multi-group Raft library in Go.![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [drmaa](https://github.com/dgruber/drmaa) **star:24** Job submission library for cluster schedulers based on the DRMAA standard.   ![It hasn't been updated in the last year][Yellow]
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [dynatomic](https://github.com/tylfin/dynatomic) **star:8** A library for using DynamoDB as an atomic counter.
* [emitter-io](https://github.com/emitter-io/emitter) **star:1924** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.![star > 1000][Silver]
* [flowgraph](https://github.com/vectaport/flowgraph) **star:18** flow-based programming package.
* [gleam](https://github.com/chrislusf/gleam) **star:2080** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.![star > 1000][Silver]
* [glow](https://github.com/chrislusf/glow) **star:2519** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.![star > 1000][Silver]
* [go-health](https://github.com/InVisionApp/go-health) **star:475** Library for enabling asynchronous dependency health checks in your service.![star > 100][Bronze]
* [go-jump](https://github.com/dgryski/go-jump) **star:253** Port of Google's "Jump" Consistent Hash function.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-kit](https://github.com/go-kit/kit) **star:14375** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.![star > 5000][Gold]
* [gorpc](https://github.com/valyala/gorpc) **star:551** Simple, fast and scalable RPC library for high load.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [grpc-go](https://github.com/grpc/grpc-go) **star:9035** The Go language implementation of gRPC. HTTP/2 based RPC.![star > 5000][Gold]   ![There was an update last week][Green]
* [hprose](https://github.com/hprose/hprose-golang) **star:1000** Very newbility RPC Library, support 25+ languages now.![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [jaeger](https://github.com/jaegertracing/jaeger) **star:8615** A distributed tracing system.![star > 5000][Gold]   ![There was an update last week][Green]
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:113** The jsonrpc package helps implement of JSON-RPC 2.0.![star > 100][Bronze]
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:100** JSON-RPC 2.0 HTTP client implementation.![star > 100][Bronze]
* [KrakenD](https://github.com/devopsfaith/krakend) **star:1735** Ultra performant API Gateway framework with middlewares.![star > 1000][Silver]
* [micro](https://github.com/micro/micro) **star:6545** Pluggable microservice toolkit and distributed systems platform.![star > 5000][Gold]   ![There was an update last week][Green]
* [NATS](https://github.com/nats-io/gnatsd) **star:6306** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.![star > 5000][Gold]   ![There was an update last week][Green]
* [outboxer](https://github.com/italolelis/outboxer) **star:5** Outboxer is a go library that implements the outbox pattern.
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [raft](https://github.com/hashicorp/raft) **star:2831** Golang implementation of the Raft consensus protocol, by HashiCorp.![star > 1000][Silver]
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [redis-lock](https://github.com/bsm/redis-lock) **star:147** Simplified distributed locking implementation using Redis.![star > 100][Bronze]
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [ringpop-go](https://github.com/uber/ringpop-go) **star:571** Scalable, fault-tolerant application-layer sharding for Go applications.![star > 100][Bronze]
* [rpcx](https://github.com/smallnest/rpcx) **star:3765** Distributed pluggable RPC service framework like alibaba Dubbo.![star > 1000][Silver]
* [sleuth](https://github.com/ursiform/sleuth) **star:300** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tendermint](https://github.com/tendermint/tendermint) **star:3119** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.![star > 1000][Silver]   ![There was an update last week][Green]
* [torrent](https://github.com/anacrolix/torrent) **star:2833** BitTorrent client package.![star > 1000][Silver]   ![There was an update last week][Green]

## Email

*Libraries and tools that implement email creation and sending.*

* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [douceur](https://github.com/aymerick/douceur) **star:158** CSS inliner for your HTML emails.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [email](https://github.com/jordan-wright/email) **star:1092** A robust and flexible email library for Go.![star > 1000][Silver]
* [go-dkim](https://github.com/toorop/go-dkim) **star:46** DKIM library, to sign & verify email.
* [go-imap](https://github.com/emersion/go-imap) **star:729** IMAP library for clients and servers.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-message](https://github.com/emersion/go-message) **star:108** Streaming library for the Internet Message Format and mail messages.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-premailer](https://github.com/vanng822/go-premailer) **star:35** Inline styling for HTML mail in Go.
* [Gomail](https://github.com/go-gomail/gomail/)  Gomail is a very simple and powerful package to send emails.
* [Hectane](https://github.com/hectane/hectane) **star:168** Lightweight SMTP client providing an HTTP API.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hermes](https://github.com/matcornic/hermes) **star:1613** Golang package that generates clean, responsive HTML e-mails.![star > 1000][Silver]
* [MailHog](https://github.com/mailhog/MailHog) **star:5144** Email and SMTP testing with web and API interface.![star > 5000][Gold]   ![There was an update last week][Green]
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:517** SendGrid's Go library for sending email.![star > 100][Bronze]
* [smtp](https://github.com/mailhog/smtp) **star:51** SMTP server protocol state machine.   ![It hasn't been updated in the last year][Yellow]

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [agora](https://github.com/PuerkitoBio/agora) **star:321** Dynamically typed, embeddable programming language in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [anko](https://github.com/mattn/anko) **star:921** Scriptable interpreter written in Go.![star > 100][Bronze]
* [binder](https://github.com/alexeyco/binder) **star:29** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   ![It hasn't been updated in the last year][Yellow]
* [expr](https://github.com/antonmedv/expr) **star:693** an engine that can evaluate expressions.![star > 100][Bronze]
* [gentee](https://github.com/gentee/gentee) **star:27** Embeddable scripting programming language.   ![There was an update last week][Green]
* [gisp](https://github.com/jcla1/gisp) **star:429** Simple LISP in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-duktape](https://github.com/olebedev/go-duktape) **star:653** Duktape JavaScript engine bindings for Go.![star > 100][Bronze]
* [go-lua](https://github.com/Shopify/go-lua) **star:1663** Port of the Lua 5.2 VM to pure Go.![star > 1000][Silver]
* [go-php](https://github.com/deuill/go-php) **star:677** PHP bindings for Go.![star > 100][Bronze]
* [go-python](https://github.com/sbinet/go-python) **star:906** naive go bindings to the CPython C-API.![star > 100][Bronze]
* [golua](https://github.com/aarzilli/golua) **star:441** Go bindings for Lua C API.![star > 100][Bronze]
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:2948** Lua 5.1 VM and compiler written in Go.![star > 1000][Silver]
* [gval](https://github.com/PaesslerAG/gval) **star:135** A highly customizable expression language written in Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [ngaro](https://github.com/db47h/ngaro) **star:19** Embeddable Ngaro VM implementation enabling scripting in Retro.   ![It hasn't been updated in the last year][Yellow]
* [otto](https://github.com/robertkrimen/otto) **star:4716** JavaScript interpreter written in Go.![star > 1000][Silver]
* [purl](https://github.com/ian-kent/purl) **star:27** Perl 5.18.2 embedded in Go.   ![It hasn't been updated in the last year][Yellow]
* [tengo](https://github.com/d5/tengo) **star:1292** Bytecode compiled script language for Go.![star > 1000][Silver]

## Error Handling

*Libraries for handling errors.*

* [errlog](https://github.com/snwfdhmp/errlog) **star:150** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.![star > 100][Bronze]   ![There was an update last week][Green]
* [errors](https://github.com/pkg/errors) **star:4859** Package that provides simple error handling primitives.![star > 1000][Silver]
* [errorx](https://github.com/joomcode/errorx) **star:552** A feature rich error package with stack traces, composition of errors and more.![star > 100][Bronze]
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:723** Go (golang) package for representing a list of errors as a single error.![star > 100][Bronze]
* [tracerr](https://github.com/ztrue/tracerr) **star:496** Golang errors with stack trace and source fragments.![star > 100][Bronze]
* [werr](https://github.com/txgruppi/werr) **star:11** Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.   ![It hasn't been updated in the last year][Yellow]

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2218** FileSystem Abstraction System for Go.![star > 1000][Silver]
* [checksum](https://github.com/codingsince1985/checksum) **star:6** Compute message digest, like MD5 and SHA256, for large files.
* [flop](https://github.com/homedepot/flop) **star:9** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:48** Load csv file using tag.
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:11** Copy files for humans.
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:2** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:15** Load gtfs files in go.
* [notify](https://github.com/rjeczalik/notify) **star:492** File system event notification library with simple API, similar to os/signal.![star > 100][Bronze]
* [opc](https://github.com/qmuntal/opc) **star:57** Load Open Packaging Conventions (OPC) files for Go.
* [pdfcpu](https://github.com/hhrutter/pdfcpu) **star:944** PDF processor.![star > 100][Bronze]   ![There was an update last week][Green]
* [skywalker](https://github.com/dixonwille/skywalker) **star:48** Package to allow one to concurrently go through a filesystem with ease.   ![It hasn't been updated in the last year][Yellow]
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [tarfs](https://github.com/posener/tarfs) **star:35** Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.   ![It hasn't been updated in the last year][Yellow]
* [vfs](https://github.com/C2FO/vfs) **star:22** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   ![There was an update last week][Green]

## Financial

*Packages for accounting and finance.*

* [accounting](https://github.com/leekchan/accounting) **star:486** money and currency formatting for golang.![star > 100][Bronze]
* [currency](https://github.com/bnkamalesh/currency) **star:8** High performant & accurate currency computation package.
* [decimal](https://github.com/shopspring/decimal) **star:1590** Arbitrary-precision fixed-point decimal numbers.![star > 1000][Silver]
* [go-finance](https://github.com/FlashBoys/go-finance) **star:536** Comprehensive financial markets data in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-finance](https://github.com/alpeb/go-finance) **star:42** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   ![It hasn't been updated in the last year][Yellow]
* [go-money](https://github.com/rhymond/go-money) **star:615** Implementation of Fowler's Money pattern.![star > 100][Bronze]   ![There was an update last week][Green]
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:61** Query OFX servers and/or parse the responses (with example command-line client).
* [orderbook](https://github.com/i25959341/orderbook) **star:68** Matching Engine for Limit Order Book in Golang.
* [techan](https://github.com/sdcoffey/techan) **star:157** Technical analysis library with advanced market analysis and trading strategies.![star > 100][Bronze]
* [transaction](https://github.com/claygod/transaction) **star:55** Embedded transactional database of accounts, running in multithreaded mode.
* [vat](https://github.com/dannyvankooten/vat) **star:61** VAT number validation & EU VAT rates.

## Forms

*Libraries for working with forms.*

* [bind](https://github.com/robfig/bind) **star:23** Bind form data to any Go values.   ![It hasn't been updated in the last year][Yellow]
* [binding](https://github.com/mholt/binding) **star:754** Binds form and JSON data from net/http Request to struct.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [conform](https://github.com/leebenson/conform) **star:173** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [form](https://github.com/go-playground/form) **star:348** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.![star > 100][Bronze]
* [formam](https://github.com/monoculum/formam) **star:122** decode form's values into a struct.![star > 100][Bronze]   ![There was an update last week][Green]
* [forms](https://github.com/albrow/forms) **star:105** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:432** CSRF protection for Go web applications & services.![star > 100][Bronze]
* [nosurf](https://github.com/justinas/nosurf) **star:971** CSRF protection middleware for Go.![star > 100][Bronze]

## Functional

*Packages to support functional programming in Go.*

* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:107** Monad, Functional Programming features for Golang.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [fuego](https://github.com/seborama/fuego) **star:38** Functional Experiment in Go.
* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1067** Useful collection of helpfully functional Go collection utilities.![star > 1000][Silver]

## Game Development

*Awesome game development libraries.*

* [Azul3D](https://github.com/azul3d/engine) **star:426** 3D game engine written in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:1846** dead simple 2D game library in Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [engo](https://github.com/EngoEngine/engo) **star:1079** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.![star > 1000][Silver]
* [g3n](https://github.com/g3n/engine) **star:744** Go 3D Game Engine.![star > 100][Bronze]
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:308** 2d game engine written in Go working on OpenGL.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [glop](https://github.com/runningwild/glop) **star:77** Glop (Game Library Of Power) is a fairly simple cross-platform game library.   ![It hasn't been updated in the last year][Yellow]
* [go-astar](https://github.com/beefsack/go-astar) **star:326** Go implementation of the A\* path finding algorithm.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-collada](https://github.com/GlenKelley/go-collada) **star:12** Go package for working with the Collada file format.   ![It hasn't been updated in the last year][Yellow]
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1153** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).![star > 1000][Silver]   ![There was an update last week][Green]
* [go3d](https://github.com/ungerik/go3d) **star:164** Performance oriented 2D/3D math package for Go.![star > 100][Bronze]
* [gonet](https://github.com/xtaci/gonet) **star:1048** Game server skeleton implemented with golang.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [goworld](https://github.com/xiaonanln/goworld) **star:1186** Scalable game server engine, featuring space-entity framework and hot-swapping.![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [Leaf](https://github.com/name5566/leaf) **star:3047** Lightweight game server framework.![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [nano](https://github.com/lonng/nano) **star:985** Lightweight, facility, high performance golang based game server framework.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [Oak](https://github.com/oakmound/oak) **star:625** Pure Go game engine.![star > 100][Bronze]   ![There was an update last week][Green]
* [Pitaya](https://github.com/topfreegames/pitaya) **star:299** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.![star > 100][Bronze]
* [Pixel](https://github.com/faiface/pixel) **star:2427** Hand-crafted 2D game library in Go.![star > 1000][Silver]
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:380** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.![star > 100][Bronze]   ![There was an update last week][Green]
* [termloop](https://github.com/JoelOtter/termloop) **star:1022** Terminal-based game engine for Go, built on top of Termbox.![star > 1000][Silver]

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [efaceconv](https://github.com/t0pep0/efaceconv) **star:43** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   ![It hasn't been updated in the last year][Yellow]
* [gen](https://github.com/clipperhouse/gen) **star:1034** Code generation tool for ‘generics’-like functionality.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [generis](https://github.com/senselogic/GENERIS) **star:18** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
* [go-enum](https://github.com/abice/go-enum) **star:85** Code generation for enums from code comments.
* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1795** .NET LINQ-like query methods for Go.![star > 1000][Silver]
* [goderive](https://github.com/awalterschulze/goderive) **star:738** Derives functions from input types.![star > 100][Bronze]   ![There was an update last week][Green]
* [gotype](https://github.com/wzshiming/gotype) **star:20** Golang source code parsing, usage like reflect package.   ![Contains Chinese documents][CN]
* [GoWrap](https://github.com/hexdigest/gowrap) **star:262** Generate decorators for Go interfaces using simple templates.![star > 100][Bronze]
* [interfaces](https://github.com/rjeczalik/interfaces) **star:187** Command line tool for generating interface definitions.![star > 100][Bronze]
* [jennifer](https://github.com/dave/jennifer) **star:1264** Generate arbitrary Go code without templates.![star > 1000][Silver]
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:85** Go preprocessor for package scoped reflection.   ![It hasn't been updated in the last year][Yellow]

## Geographic

*Geographic tools and servers*

* [geocache](https://github.com/melihmucuk/geocache) **star:111** In-memory cache that is suitable for geolocation based applications.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [geoserver](https://github.com/hishamkaram/geoserver) **star:25** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
* [gismanager](https://github.com/hishamkaram/gismanager) **star:19** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.
* [osm](https://github.com/paulmach/osm) **star:67** Library for reading, writing and working with OpenStreetMap data and APIs.
* [pbf](https://github.com/maguro/pbf) **star:15** OpenStreetMap PBF golang encoder/decoder.
* [S2 geometry](https://github.com/golang/geo) **star:884** S2 geometry library in Go.![star > 100][Bronze]
* [Tile38](https://github.com/tidwall/tile38) **star:6314** Geolocation DB with spatial index and realtime geofencing.![star > 5000][Gold]   ![There was an update last week][Green]

## Go Compilers

*Tools for compiling Go to other languages.*

* [c4go](https://github.com/Konstantin8105/c4go) **star:150** Transpile C code to Go code.![star > 100][Bronze]
* [f4go](https://github.com/Konstantin8105/f4go) **star:11** Transpile FORTRAN 77 code to Go code.
* [gopherjs](https://github.com/gopherjs/gopherjs) **star:8527** Compiler from Go to JavaScript.![star > 5000][Gold]
* [llgo](https://github.com/go-llvm/llgo) **star:989** LLVM-based compiler for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:393** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) **star:1862** A high-performance goroutine pool for golang.![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [artifex](https://github.com/borderstech/artifex) **star:12** Simple in-memory job queue for Golang using worker-based dispatching.
* [async](https://github.com/studiosol/async) **star:22** A safe way to execute functions asynchronously, recovering them in case of panic.
* [breaker](https://github.com/kamilsk/breaker) **star:34** Flexible mechanism to make execution flow interruptible.
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:36** CyclicBarrier for golang.
* [go-floc](https://github.com/workanator/go-floc) **star:167** Orchestrate goroutines with ease.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:103** Control goroutines execution order.![star > 100][Bronze]
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.
* [go-trylock](https://github.com/subchen/go-trylock) **star:4** TryLock support on read-write lock for Golang.   ![It hasn't been updated in the last year][Yellow]
* [gollback](https://github.com/vardius/gollback) **star:27** asynchronous simple function utilities, for managing execution of closures and callbacks.
* [GoSlaves](https://github.com/themester/GoSlaves) **star:75** Simple and Asynchronous Goroutine pool library.
* [goworker](https://github.com/benmanns/goworker) **star:2243** goworker is a Go-based background worker.![star > 1000][Silver]
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:56** manages a resizeable pool of context-aware goroutines to bound concurrency.
* [grpool](https://github.com/ivpusic/grpool) **star:496** Lightweight Goroutine pool.![star > 100][Bronze]
* [Hunch](https://github.com/AaronJan/Hunch) **star:11** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   ![There was an update last week][Green]
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:24** Run functions in parallel.   ![It hasn't been updated in the last year][Yellow]
* [pool](https://github.com/go-playground/pool) **star:479** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [queue](https://github.com/AnikHasibul/queue) **star:1** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.
* [routine](https://github.com/x-mod/routine) **star:3** go routine control with context, support: Main, Go, Pool and some useful Executors.   ![There was an update last week][Green]
* [semaphore](https://github.com/kamilsk/semaphore) **star:75** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
* [semaphore](https://github.com/marusama/semaphore) **star:72** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
* [stl](https://github.com/ssgreg/stl) **star:8** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
* [threadpool](https://github.com/shettyh/threadpool) **star:18** Golang threadpool implementation.
* [tunny](https://github.com/Jeffail/tunny) **star:1338** Goroutine pool for golang.![star > 1000][Silver]
* [worker-pool](https://github.com/vardius/worker-pool) **star:44** goworker is a Go simple async worker pool.
* [workerpool](https://github.com/gammazero/workerpool) **star:134** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.![star > 100][Bronze]

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [app](https://github.com/murlokswarm/app) **star:2946** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.![star > 1000][Silver]   ![There was an update last week][Green]
* [fyne](https://github.com/fyne-io/fyne) **star:6246** Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows.![star > 5000][Gold]   ![There was an update last week][Green]
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:2645** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).![star > 1000][Silver]
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1446** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.![star > 1000][Silver]   ![There was an update last week][Green]
* [gotk3](https://github.com/gotk3/gotk3) **star:765** Go bindings for GTK3.![star > 100][Bronze]   ![There was an update last week][Green]
* [gowd](https://github.com/dtylman/gowd) **star:208** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.![star > 100][Bronze]
* [qt](https://github.com/therecipe/qt) **star:6016** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).![star > 5000][Gold]
* [ui](https://github.com/andlabs/ui) **star:6928** Platform-native GUI library for Go. Cross platform.![star > 5000][Gold]
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [walk](https://github.com/lxn/walk) **star:3681** Windows application library kit for Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [webview](https://github.com/zserge/webview) **star:4634** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).![star > 1000][Silver]   ![There was an update last week][Green]

*Interaction*

* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:1** Go bindings for libappindicator3 C library.
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:494** OSX Desktop Notifications library for Go.![star > 100][Bronze]
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:1** OSX library to notify about any (pluggable) activity on your machine.
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier)  OSX Sleep/Wake notifications in golang.
* [robotgo](https://github.com/go-vgo/robotgo) **star:4417** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.![star > 1000][Silver]   ![There was an update last week][Green]
* [systray](https://github.com/getlantern/systray) **star:779** Cross platform Go library to place an icon and menu in the notification area.![star > 100][Bronze]   ![There was an update last week][Green]
* [trayhost](https://github.com/shurcooL/trayhost) **star:159** Cross-platform Go library to place an icon in the host operating system's taskbar.![star > 100][Bronze]


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [bild](https://github.com/anthonynsimon/bild) **star:2441** Collection of image processing algorithms in pure Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [bimg](https://github.com/h2non/bimg) **star:786** Small package for fast and efficient image processing using libvips.![star > 100][Bronze]
* [cameron](https://github.com/aofei/cameron) **star:31** An avatar generator for Go.
* [darkroom](https://github.com/gojek/darkroom) **star:20** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   ![There was an update last week][Green]
* [geopattern](https://github.com/pravj/geopattern) **star:1011** Create beautiful generative image patterns from a string.![star > 1000][Silver]
* [gg](https://github.com/fogleman/gg) **star:1901** 2D rendering in pure Go.![star > 1000][Silver]
* [gift](https://github.com/disintegration/gift) **star:1207** Package of image processing filters.![star > 1000][Silver]
* [gltf](https://github.com/qmuntal/gltf) **star:37** Efficient and robust glTF 2.0 reader, writer and validator.   ![There was an update last week][Green]
* [go-cairo](https://github.com/ungerik/go-cairo) **star:85** Go binding for the cairo graphics library.
* [go-gd](https://github.com/bolknote/go-gd) **star:50** Go binding for GD library.   ![It hasn't been updated in the last year][Yellow]
* [go-nude](https://github.com/koyachi/go-nude) **star:286** Nudity detection with Go.![star > 100][Bronze]
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1092** Go bindings for OpenCV.![star > 1000][Silver]
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Port of webcolors library from Python to Go.   ![It hasn't been updated in the last year][Yellow]
* [gocv](https://github.com/hybridgroup/gocv) **star:2436** Go package for computer vision using OpenCV 3.3+.![star > 1000][Silver]   ![There was an update last week][Green]
* [goimagehash](https://github.com/corona10/goimagehash) **star:212** Go Perceptual image hashing package.![star > 100][Bronze]
* [goimghdr](https://github.com/corona10/goimghdr) **star:26** The imghdr module determines the type of image contained in a file for Go.
* [govatar](https://github.com/o1egl/govatar) **star:310** Library and CMD tool for generating funny avatars.![star > 100][Bronze]
* [image2ascii](https://github.com/qeesung/image2ascii) **star:289** Convert image to ASCII.![star > 100][Bronze]
* [imagick](https://github.com/gographics/imagick) **star:972** Go binding to ImageMagick's MagickWand C API.![star > 100][Bronze]
* [imaginary](https://github.com/h2non/imaginary) **star:2587** Fast and simple HTTP microservice for image resizing.![star > 1000][Silver]
* [imaging](https://github.com/disintegration/imaging) **star:2511** Simple Go image processing package.![star > 1000][Silver]
* [img](https://github.com/hawx/img) **star:129** Selection of image manipulation tools.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [ln](https://github.com/fogleman/ln) **star:2458** 3D line art rendering in Go.![star > 1000][Silver]
* [mergi](https://github.com/noelyahan/mergi) **star:70** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
* [mort](https://github.com/aldor007/mort) **star:365** Storage and image processing server written in Go.![star > 100][Bronze]
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.
* [picfit](https://github.com/thoas/picfit) **star:1064** An image resizing server written in Go.![star > 1000][Silver]
* [pt](https://github.com/fogleman/pt) **star:1768** Path tracing engine written in Go.![star > 1000][Silver]
* [resize](https://github.com/nfnt/resize) **star:2123** Image resizing for Go with common interpolation methods.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [rez](https://github.com/bamiaux/rez) **star:189** Image resizing in pure Go and SIMD.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [smartcrop](https://github.com/muesli/smartcrop) **star:1253** Finds good crops for arbitrary images and crop sizes.![star > 1000][Silver]
* [steganography](https://github.com/auyer/steganography) **star:25** Pure Go Library for LSB steganography.   ![There was an update last week][Green]
* [stegify](https://github.com/DimitarPetrov/stegify) **star:487** Go tool for LSB steganography, capable of hiding any file within an image.![star > 100][Bronze]
* [svgo](https://github.com/ajstarks/svgo) **star:1327** Go Language Library for SVG generation.![star > 1000][Silver]
* [tga](https://github.com/ftrvxmtrx/tga) **star:23** Package tga is a TARGA image format decoder/encoder.   ![It hasn't been updated in the last year][Yellow]

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb) **star:168** Open-Source Platform for Quantified Self & IoT.![star > 100][Bronze]
* [devices](https://github.com/goiot/devices) **star:225** Suite of libraries for IoT devices, experimental for x/exp/io.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [eywa](https://github.com/xcodersun/eywa) **star:36** Project Eywa is essentially a connection manager that keeps track of connected devices.   ![It hasn't been updated in the last year][Yellow]
* [flogo](https://github.com/tibcosoftware/flogo) **star:1116** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.![star > 1000][Silver]   ![There was an update last week][Green]
* [gatt](https://github.com/paypal/gatt) **star:812** Gatt is a Go package for building Bluetooth Low Energy peripherals.![star > 100][Bronze]
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [huego](https://github.com/amimof/huego) **star:109** An extensive Philips Hue client library for Go.![star > 100][Bronze]
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [mainflux](https://github.com/Mainflux/mainflux) **star:585** Industrial IoT Messaging and Device Management Server.![star > 100][Bronze]   ![There was an update last week][Green]
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:179** Lightweight stream processing engine for IoT.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Job Scheduler

*Libraries for scheduling jobs.*

* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [clockwork](https://github.com/whiteShtef/clockwork) **star:76** Simple and intuitive job scheduling library in Go.
* [go-cron](https://github.com/rk/go-cron) **star:177** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gron](https://github.com/roylee0704/gron) **star:629** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.![star > 100][Bronze]
* [JobRunner](https://github.com/bamzi/jobrunner) **star:570** Smart and featureful cron job scheduler with job queuing and live monitoring built in.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [jobs](https://github.com/albrow/jobs) **star:451** Persistent and flexible background jobs library.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:37** Job scheduler that supports webhooks, crons and classic scheduling.   ![There was an update last week][Green]
* [scheduler](https://github.com/carlescere/scheduler) **star:293** Cronjobs scheduling made easy.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## JSON

*Libraries for working with JSON.*

* [ajson](https://github.com/spyzhov/ajson) **star:12** Abstract JSON for golang with JSONPath support.
* [gjo](https://github.com/skanehira/gjo) **star:60** Small utility to create JSON objects.
* [GJSON](https://github.com/tidwall/gjson) **star:4854** Get a JSON value with one line of code.![star > 1000][Silver]
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:7** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.
* [go-respond](https://github.com/nicklaw5/go-respond) **star:22** Go package for handling common HTTP JSON responses.
* [gojq](https://github.com/elgs/gojq) **star:140** JSON query in Golang.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2025** Automatically generate Go (golang) struct definitions from example JSON.![star > 1000][Silver]
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:40** JSON diff utility written in Go.
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:5** Go bindings based on the JSON API errors reference.   ![It hasn't been updated in the last year][Yellow]
* [jsonf](https://github.com/miolini/jsonf) **star:55** Console tool for highlighted formatting and struct query fetching JSON.   ![It hasn't been updated in the last year][Yellow]
* [jsongo](https://github.com/ricardolonga/jsongo) **star:92** Fluent API to make it easier to create Json objects.   ![It hasn't been updated in the last year][Yellow]
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.
* [kazaam](https://github.com/Qntfy/kazaam) **star:131** API for arbitrary transformation of JSON documents.![star > 100][Bronze]
* [mp](https://github.com/sanbornm/mp) **star:33** Simple cli email parser. It currently takes stdin and outputs JSON.   ![It hasn't been updated in the last year][Yellow]

## Logging

*Libraries for generating and working with log files.*

* [distillog](https://github.com/amoghe/distillog) **star:19** distilled levelled logging (think of it as stdlib + log levels).   ![It hasn't been updated in the last year][Yellow]
* [glg](https://github.com/kpango/glg) **star:51** glg is simple and fast leveled logging library for Go.
* [glo](https://github.com/lajosbencz/glo) **star:8** PHP Monolog inspired logging facility with identical severity levels.
* [glog](https://github.com/golang/glog) **star:2300** Leveled execution logs for Go.![star > 1000][Silver]
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:19** Simple writer that rotate log files automatically based on current date and time, like cronolog.
* [go-log](https://github.com/subchen/go-log) **star:10** Simple and configurable Logging in Go, with level, formatters and writers.   ![It hasn't been updated in the last year][Yellow]
* [go-log](https://github.com/siddontang/go-log) **star:23** Log lib supports level and multi handlers.
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j implementation in Go.   ![It hasn't been updated in the last year][Yellow]
* [go-logger](https://github.com/apsdehal/go-logger) **star:233** Simple logger of Go Programs, with level handlers.![star > 100][Bronze]
* [gologger](https://github.com/sadlil/gologger) **star:39** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   ![It hasn't been updated in the last year][Yellow]
* [gomol](https://github.com/aphistic/gomol) **star:15** Multiple-output, structured logging for Go with extensible logging outputs.
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [journald](https://github.com/ssgreg/journald) **star:19** Go implementation of systemd Journal's native API for logging.
* [log](https://github.com/aerogo/log) **star:4** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).
* [log](https://github.com/apex/log) **star:729** Structured logging package for Go.![star > 100][Bronze]
* [log](https://github.com/go-playground/log) **star:266** Simple, configurable and scalable Structured Logging for Go.![star > 100][Bronze]
* [log](https://github.com/teris-io/log) **star:22** Structured log interface for Go cleanly separates logging facade from its implementation.   ![It hasn't been updated in the last year][Yellow]
* [log-voyage](https://github.com/firstrow/logvoyage) **star:82** Full-featured logging saas written in golang.   ![It hasn't been updated in the last year][Yellow]
* [log15](https://github.com/inconshreveable/log15) **star:883** Simple, powerful logging for Go.![star > 100][Bronze]
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   ![It hasn't been updated in the last year][Yellow]
* [logex](https://github.com/chzyer/logex) **star:32** Golang log lib, supports tracking and level, wrap by standard log lib.   ![It hasn't been updated in the last year][Yellow]
* [logger](https://github.com/azer/logger) **star:135** Minimalistic logging library for Go.![star > 100][Bronze]
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.
* [logo](https://github.com/mbndr/logo) **star:4** Golang logger to different configurable writers.   ![It hasn't been updated in the last year][Yellow]
* [logrus](https://github.com/Sirupsen/logrus) **star:11859** Structured logger for Go.![star > 5000][Gold]
* [logrusly](https://github.com/sebest/logrusly) **star:26** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   ![It hasn't been updated in the last year][Yellow]
* [logutils](https://github.com/hashicorp/logutils) **star:247** Utilities for slightly better logging in Go (Golang) extending the standard logger.![star > 100][Bronze]
* [logxi](https://github.com/mgutz/logxi) **star:334** 12-factor app logger that is fast and makes you happy.![star > 100][Bronze]
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1433** Simple rolling logger, implements io.WriteCloser.![star > 1000][Silver]
* [mlog](https://github.com/jbrodriguez/mlog) **star:18** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   ![It hasn't been updated in the last year][Yellow]
* [onelog](https://github.com/francoispqt/onelog) **star:330** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.![star > 100][Bronze]
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:110** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:96** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.
* [seelog](https://github.com/cihub/seelog) **star:1348** Logging functionality with flexible dispatching, filtering, and formatting.![star > 1000][Silver]
* [spew](https://github.com/davecgh/go-spew) **star:3291** Implements a deep pretty printer for Go data structures to aid in debugging.![star > 1000][Silver]
* [stdlog](https://github.com/alexcesaro/log) **star:43** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   ![It hasn't been updated in the last year][Yellow]
* [tail](https://github.com/hpcloud/tail) **star:1534** Go package striving to emulate the features of the BSD tail program.![star > 1000][Silver]
* [xlog](https://github.com/xfxdev/xlog) **star:7** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
* [xlog](https://github.com/rs/xlog) **star:128** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.![star > 100][Bronze]
* [zap](https://github.com/uber-go/zap) **star:7388** Fast, structured, leveled logging in Go.![star > 5000][Gold]
* [zerolog](https://github.com/rs/zerolog) **star:2196** Zero-allocation JSON logger.![star > 1000][Silver]

## Machine Learning

*Libraries for Machine Learning.*

* [bayesian](https://github.com/jbrukh/bayesian) **star:630** Naive Bayesian Classification for Golang.![star > 100][Bronze]
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:644** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.![star > 100][Bronze]
* [eaopt](https://github.com/MaxHalford/eaopt) **star:619** An evolutionary optimization library.![star > 100][Bronze]   ![There was an update last week][Green]
* [evoli](https://github.com/khezen/evoli) **star:8** Genetic Algorithm and Particle Swarm Optimization library.
* [fonet](https://github.com/Fontinalis/fonet) **star:33** A Deep Neural Network library written in Go.   ![It hasn't been updated in the last year][Yellow]
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:21** Go implementation of the k-modes and k-prototypes clustering algorithms.   ![It hasn't been updated in the last year][Yellow]
* [go-deep](https://github.com/patrikeh/go-deep) **star:216** A feature-rich neural network library in Go.![star > 100][Bronze]
* [go-fann](https://github.com/white-pony/go-fann) **star:99** Go bindings for Fast Artificial Neural Networks(FANN) library.   ![It hasn't been updated in the last year][Yellow]
* [go-galib](https://github.com/thoj/go-galib) **star:171** Genetic Algorithms library written in Go / golang.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-pr](https://github.com/daviddengcn/go-pr) **star:57** Pattern recognition package in Go lang.   ![It hasn't been updated in the last year][Yellow]
* [gobrain](https://github.com/goml/gobrain) **star:383** Neural Networks written in go.![star > 100][Bronze]
* [godist](https://github.com/e-dard/godist) **star:24** Various probability distributions, and associated methods.   ![It hasn't been updated in the last year][Yellow]
* [goga](https://github.com/tomcraven/goga) **star:78** Genetic algorithm library for Go.   ![It hasn't been updated in the last year][Yellow]
* [GoLearn](https://github.com/sjwhitworth/golearn) **star:6643** General Machine Learning library for Go.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.
* [GoMind](https://github.com/surenderthakran/gomind) **star:6** A simplistic Neural Network Library in Go.   ![It hasn't been updated in the last year][Yellow]
* [goml](https://github.com/cdipaolo/goml) **star:1013** On-line Machine Learning in Go.![star > 1000][Silver]
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:143** Recommendation Algorithms library written in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gorgonia](https://github.com/chewxy/gorgonia) **star:2682** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.![star > 1000][Silver]   ![There was an update last week][Green]
* [gorse](https://github.com/zhenghaoz/gorse) **star:530** A High Performance Recommender System Package based on Collaborative Filtering for Go.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [goscore](https://github.com/asafschers/goscore) **star:35** Go Scoring API for PMML.
* [gosseract](https://github.com/otiai10/gosseract) **star:866** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.![star > 100][Bronze]
* [libsvm](https://github.com/datastream/libsvm) **star:63** libsvm golang version derived work based on LIBSVM 3.14.   ![It hasn't been updated in the last year][Yellow]
* [mlgo](https://github.com/NullHypothesis/mlgo) **star:5** This project aims to provide minimalistic machine learning algorithms in Go.   ![It hasn't been updated in the last year][Yellow]
* [neat](https://github.com/jinyeom/neat) **star:55** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   ![It hasn't been updated in the last year][Yellow]
* [neural-go](https://github.com/schuyler/neural-go) **star:61** Multilayer perceptron network implemented in Go, with training via backpropagation.   ![It hasn't been updated in the last year][Yellow]
* [ocrserver](https://github.com/otiai10/ocrserver) **star:224** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.![star > 100][Bronze]
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:149** Go Interface to Open Neural Network Exchange (ONNX).![star > 100][Bronze]
* [probab](https://github.com/ThePaw/probab) **star:10** Probability distribution functions. Bayesian inference. Written in pure Go.   ![It hasn't been updated in the last year][Yellow]
* [regommend](https://github.com/muesli/regommend) **star:246** Recommendation & collaborative filtering engine.![star > 100][Bronze]
* [shield](https://github.com/eaigner/shield) **star:124** Bayesian text classifier with flexible tokenizers and storage backends for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [tfgo](https://github.com/galeone/tfgo) **star:1180** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.![star > 1000][Silver]
* [Varis](https://github.com/Xamber/Varis) **star:24** Golang Neural Network.   ![It hasn't been updated in the last year][Yellow]

## Messaging

*Libraries that implement messaging systems.*

* [APNs2](https://github.com/sideshow/apns2) **star:2038** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.![star > 1000][Silver]
* [Beaver](https://github.com/Clivern/Beaver) **star:721** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.![star > 100][Bronze]   ![There was an update last week][Green]
* [Benthos](https://github.com/Jeffail/benthos) **star:1999** A message streaming bridge between a range of protocols.![star > 1000][Silver]   ![There was an update last week][Green]
* [Bus](https://github.com/mustafaturan/bus) **star:114** Minimalist message bus implementation for internal communication.![star > 100][Bronze]
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:3668** Real-time messaging (Websockets or SockJS) server in Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [Commander](https://github.com/jeroenrinzema/commander) **star:21** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   ![There was an update last week][Green]
* [dbus](https://github.com/godbus/dbus) **star:355** Native Go bindings for D-Bus.![star > 100][Bronze]   ![There was an update last week][Green]
* [drone-line](https://github.com/appleboy/drone-line) **star:60** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.
* [emitter](https://github.com/olebedev/emitter) **star:309** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.![star > 100][Bronze]
* [event](https://github.com/agoalofalife/event) **star:27** Implementation of the pattern observer.   ![It hasn't been updated in the last year][Yellow]
* [EventBus](https://github.com/asaskevich/EventBus) **star:554** The lightweight event bus with async compatibility.![star > 100][Bronze]
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** Gaurun Client written in Go.
* [Glue](https://github.com/desertbit/glue) **star:317** Robust Go and Javascript Socket Library (Alternative to Socket.io).![star > 100][Bronze]
* [go-notify](https://github.com/TheCreeper/go-notify) **star:47** Native implementation of the freedesktop notification spec.
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1453** the official Go package for NSQ.![star > 1000][Silver]
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:2875** socket.io library for golang, a realtime application framework.![star > 1000][Silver]
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:11** Client library to Viessmann Vitotrol web service.
* [Gollum](https://github.com/trivago/gollum) **star:767** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.![star > 100][Bronze]   ![There was an update last week][Green]
* [golongpoll](https://github.com/jcuga/golongpoll) **star:426** HTTP longpoll server library that makes web pub-sub simple.![star > 100][Bronze]
* [goose](https://github.com/ian-kent/goose) **star:36** Server Sent Events in Go.   ![It hasn't been updated in the last year][Yellow]
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1834** gopush-cluster is a go push server cluster.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [gorush](https://github.com/appleboy/gorush) **star:3697** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).![star > 1000][Silver]
* [guble](https://github.com/smancke/guble) **star:138** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hub](https://github.com/leandro-lugaresi/hub) **star:24** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   ![It hasn't been updated in the last year][Yellow]
* [jazz](https://github.com/socifi/jazz) **star:6** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.
* [machinery](https://github.com/RichardKnop/machinery) **star:3355** Asynchronous task queue/job queue based on distributed message passing.![star > 1000][Silver]   ![There was an update last week][Green]
* [mangos](https://github.com/go-mangos/mangos) **star:1534** Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.![star > 1000][Silver]
* [melody](https://github.com/olahol/melody) **star:1553** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.![star > 1000][Silver]
* [Mercure](https://github.com/dunglas/mercure) **star:1494** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).![star > 1000][Silver]   ![There was an update last week][Green]
* [messagebus](https://github.com/vardius/message-bus) **star:65** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
* [NATS Go Client](https://github.com/nats-io/nats) **star:2390** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.![star > 1000][Silver]   ![There was an update last week][Green]
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:50** A tiny wrapper around NSQ topic and channel.   ![It hasn't been updated in the last year][Yellow]
* [oplog](https://github.com/dailymotion/oplog) **star:94** Generic oplog/replication system for REST APIs.   ![It hasn't been updated in the last year][Yellow]
* [pubsub](https://github.com/tuxychandru/pubsub) **star:275** Simple pubsub package for go.![star > 100][Bronze]
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:61** A tiny wrapper over amqp exchanges and queues.
* [rabtap](https://github.com/jandelgado/rabtap) **star:72** RabbitMQ swiss army knife cli app.
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:54** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   ![It hasn't been updated in the last year][Yellow]
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:1** redisqueue provides a producer and consumer of a queue that uses Redis streams.
* [rmqconn](https://github.com/sbabiv/rmqconn)  RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.
* [sarama](https://github.com/Shopify/sarama) **star:4592** Go library for Apache Kafka.![star > 1000][Silver]   ![There was an update last week][Green]
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1097** Redis backed unified push service for server-side notifications to mobile devices.![star > 1000][Silver]
* [zmq4](https://github.com/pebbe/zmq4) **star:773** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).![star > 100][Bronze]

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:1693** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.![star > 1000][Silver]   ![There was an update last week][Green]

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:4427** Golang library for reading and writing Microsoft Excel™ (XLSX) files.![star > 1000][Silver]   ![There was an update last week][Green]
* [go-excel](https://github.com/szyhf/go-excel) **star:46** A simple and light reader to read a relate-db-like excel as a table.
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:12** Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.   ![It hasn't been updated in the last year][Yellow]
* [xlsx](https://github.com/tealeg/xlsx) **star:3351** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.![star > 1000][Silver]   ![There was an update last week][Green]
* [xlsx](https://github.com/plandem/xlsx) **star:68** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   ![There was an update last week][Green]

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [alice](https://github.com/magic003/alice) **star:34** Additive dependency injection container for Golang.   ![It hasn't been updated in the last year][Yellow]
* [dig](https://github.com/uber-go/dig) **star:902** A reflection based dependency injection toolkit for Go.![star > 100][Bronze]
* [fx](https://github.com/uber-go/fx) **star:741** A dependency injection based application framework for Go (built on top of dig).![star > 100][Bronze]
* [gocontainer](https://github.com/vardius/gocontainer) **star:9** Simple Dependency Injection Container.
* [inject](https://github.com/defval/inject) **star:25** A reflection based dependency injection container with simple interface.
* [linker](https://github.com/logrange/linker) **star:5** A reflection based dependency injection and inversion of control library with components lifecycle support.
* [wire](https://github.com/Fs02/wire) **star:19** Strict Runtime Dependency Injection for Golang.

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [go-sample](https://github.com/zitryss/go-sample) **star:23** A sample layout for Go application projects with the real code.
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:8993** Set of common historical and emerging project layout patterns in the Go ecosystem.![star > 5000][Gold]
* [scaffold](https://github.com/catchplay/scaffold) **star:24** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.

### Strings

*Libraries for working with strings.*

* [strutil](https://github.com/ozgio/strutil) **star:62** String utilities.
* [xstrings](https://github.com/huandu/xstrings) **star:617** Collection of useful string functions ported from other languages.![star > 100][Bronze]

*These libraries were placed here because none of the other categories seemed to fit.*

* [anagent](https://github.com/mudler/anagent) **star:11** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.
* [antch](https://github.com/antchfx/antch) **star:140** A fast, powerful and extensible web crawling & scraping framework.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [archiver](https://github.com/mholt/archiver) **star:2482** Library and command for making and extracting .zip and .tar.gz archives.![star > 1000][Silver]
* [autoflags](https://github.com/artyom/autoflags) **star:24** Go package to automatically define command line flags from struct fields.
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:9** Calculate average score and rating based on Wilson Score Equation.   ![It hasn't been updated in the last year][Yellow]
* [banner](https://github.com/dimiro1/banner) **star:231** Add beautiful banners into your Go applications.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:624** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [battery](https://github.com/distatus/battery) **star:135** Cross-platform, normalized battery information library.![star > 100][Bronze]
* [bitio](https://github.com/icza/bitio) **star:92** Highly optimized bit-level Reader and Writer for Go.   ![It hasn't been updated in the last year][Yellow]
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:29** GoLang Library for [Browser Capabilities Project](http://browscap.org/).
* [captcha](https://github.com/steambap/captcha) **star:43** Package captcha provides an easy to use, unopinionated API for captcha generation.
* [conv](https://github.com/cstockton/go-conv) **star:341** Package conv provides fast and intuitive conversions across Go types.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [datacounter](https://github.com/miolini/datacounter) **star:27** Go counters for readers/writer/http.ResponseWriter.
* [ffmt](https://github.com/go-ffmt/ffmt) **star:126** Beautify data display for Humans.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [ghorg](https://github.com/gabrie30/ghorg) **star:24** Clone all repos from a GitHub org into a single directory.   ![There was an update last week][Green]
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:666** Generic object pool for Golang.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:849** Resiliency patterns for golang.![star > 100][Bronze]
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:67** Decompression library for RAR, TAR, ZIP and 7z archives.
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:413** Random data generator written in go.![star > 100][Bronze]
* [gommit](https://github.com/antham/gommit) **star:72** Analyze git commit messages to ensure they follow defined patterns.
* [gopsutil](https://github.com/shirou/gopsutil) **star:3904** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).![star > 1000][Silver]   ![There was an update last week][Green]
* [gosh](https://github.com/osamingo/gosh) **star:17** Provide Go Statistics Handler, Struct, Measure Method.
* [gosms](https://github.com/haxpax/gosms) **star:1226** Your own local SMS gateway in Go that can be used to send SMS.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [gotoprom](https://github.com/cabify/gotoprom) **star:15** Type-safe metrics builder wrapper library for the official Prometheus client.
* [gountries](https://github.com/pariz/gountries) **star:209** Package that exposes country and subdivision data.![star > 100][Bronze]
* [health](https://github.com/dimiro1/health) **star:360** Easy to use, extensible health check library.![star > 100][Bronze]
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:82** An opinionated and concurrent health-check HTTP handler for RESTful services.
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.
* [indigo](https://github.com/osamingo/indigo) **star:51** Distributed unique ID generator of using Sonyflake and encoded by Base58.
* [lk](https://github.com/hyperboloide/lk) **star:121** A simple licensing library for golang.![star > 100][Bronze]
* [llvm](https://github.com/llir/llvm) **star:410** Library for interacting with LLVM IR in pure Go.![star > 100][Bronze]
* [metrics](https://github.com/pascaldekloe/metrics) **star:4** Library for metrics instrumentation and Prometheus exposition.
* [morse](https://github.com/alwindoss/morse) **star:50** Library to convert to and from morse code.
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:33** HTTP service to generate PDF from Json requests.   ![It hasn't been updated in the last year][Yellow]
* [persian](https://github.com/mavihq/persian) **star:32** Some utilities for Persian language in go.   ![It hasn't been updated in the last year][Yellow]
* [sandid](https://github.com/aofei/sandid) **star:12** Every grain of sand on earth has its own ID.
* [shellwords](https://github.com/Wing924/shellwords) **star:7** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   ![It hasn't been updated in the last year][Yellow]
* [shortid](https://github.com/teris-io/shortid) **star:448** Distributed generation of super short, unique, non-sequential, URL friendly IDs.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [stats](https://github.com/go-playground/stats) **star:121** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [turtle](https://github.com/hackebrot/turtle) **star:73** Emojis for Go.   ![It hasn't been updated in the last year][Yellow]
* [url-shortener](https://github.com/pantrif/url-shortener) **star:17** A modern, powerful, and robust URL shortener microservice with mysql support.   ![It hasn't been updated in the last year][Yellow]
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [xdg](https://github.com/rkoesters/xdg) **star:20** FreeDesktop.org (xdg) Specs implemented in Go.
* [xkg](https://github.com/go-xkg/xkg) **star:40** X Keyboard Grabber.   ![It hasn't been updated in the last year][Yellow]

## Natural Language Processing

*Libraries for working with human languages.*

* [getlang](https://github.com/rylans/getlang) **star:74** Fast natural language detection package.
* [go-eco](https://github.com/ThePaw/go-eco) **star:4** Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models.   ![It hasn't been updated in the last year][Yellow]
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo bindings to Yandex.Mystem - russian morphology analyzer.   ![It hasn't been updated in the last year][Yellow]
* [go-nlp](https://github.com/nuance/go-nlp) **star:79** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   ![It hasn't been updated in the last year][Yellow]
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:521** CN Hanzi to Hanyu Pinyin converter.![star > 100][Bronze]
* [go-stem](https://github.com/agonopol/go-stem) **star:52** Implementation of the porter stemming algorithm.   ![It hasn't been updated in the last year][Yellow]
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:56** ASCII transliterations of Unicode text.
* [go2vec](https://github.com/danieldk/go2vec) **star:30** Reader and utility functions for word2vec embeddings.
* [gojieba](https://github.com/yanyiwu/gojieba) **star:816** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:15** Go bindings for the snowball libstemmer library including porter 2.   ![It hasn't been updated in the last year][Yellow]
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)
* [gounidecode](https://github.com/fiam/gounidecode) **star:67** Unicode transliterator (also known as unidecode) for Go.   ![It hasn't been updated in the last year][Yellow]
* [gse](https://github.com/go-ego/gse) **star:1068** Go efficient text segmentation; support english, chinese, japanese and other.![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   ![It hasn't been updated in the last year][Yellow]
* [kagome](https://github.com/ikawaha/kagome) **star:413** JP morphological analyzer written in pure Go.![star > 100][Bronze]
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   ![It hasn't been updated in the last year][Yellow]
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   ![It hasn't been updated in the last year][Yellow]
* [nlp](https://github.com/Shixzie/nlp) **star:353** Extract values from strings and fill your structs with nlp.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [nlp](https://github.com/james-bowman/nlp) **star:216** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).![star > 100][Bronze]
* [paicehusk](https://github.com/rookii/paicehusk) **star:25** Golang implementation of the Paice/Husk Stemming Algorithm.   ![It hasn't been updated in the last year][Yellow]
* [petrovich](https://github.com/striker2000/petrovich) **star:22** Petrovich is the library which inflects Russian names to given grammatical case.
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   ![It hasn't been updated in the last year][Yellow]
* [porter2](https://github.com/zhenjl/porter2) **star:33** Really fast Porter 2 stemmer.   ![It hasn't been updated in the last year][Yellow]
* [prose](https://github.com/jdkato/prose) **star:2040** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more.![star > 1000][Silver]
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:45** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
* [segment](https://github.com/blevesearch/segment) **star:46** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   ![It hasn't been updated in the last year][Yellow]
* [sentences](https://github.com/neurosnap/sentences) **star:260** Sentence tokenizer:  converts text into a list of sentences.![star > 100][Bronze]
* [shamoji](https://github.com/osamingo/shamoji) **star:10** The shamoji is word filtering package written in Go.
* [snowball](https://github.com/goodsign/snowball) **star:24** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   ![It hasn't been updated in the last year][Yellow]
* [stemmer](https://github.com/dchest/stemmer) **star:47** Stemmer packages for Go programming language. Includes English and German stemmers.   ![It hasn't been updated in the last year][Yellow]
* [textcat](https://github.com/pebbe/textcat) **star:60** Go package for n-gram based text categorization, with support for utf-8 and raw text.   ![It hasn't been updated in the last year][Yellow]
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:350** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).![star > 100][Bronze]
* [when](https://github.com/olebedev/when) **star:926** Natural EN and RU language date/time parser with pluggable rules.![star > 100][Bronze]

## Networking

*Libraries for working with various layers of the network.*

* [arp](https://github.com/mdlayher/arp) **star:195** Package arp implements the ARP protocol, as described in RFC 826.![star > 100][Bronze]
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:232** Streaming protocolbuffer data over TCP made easy.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [canopus](https://github.com/zubairhamed/canopus) **star:133** CoAP Client/Server implementation (RFC 7252).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [cidranger](https://github.com/yl2chen/cidranger) **star:387** Fast IP to CIDR lookup for Go.![star > 100][Bronze]
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:62** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
* [dns](https://github.com/miekg/dns) **star:3784** Go library for working with DNS.![star > 1000][Silver]
* [ether](https://github.com/songgao/ether) **star:62** Cross-platform Go package for sending and receiving ethernet frames.   ![It hasn't been updated in the last year][Yellow]
* [ethernet](https://github.com/mdlayher/ethernet) **star:185** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.![star > 100][Bronze]
* [fasthttp](https://github.com/valyala/fasthttp) **star:9279** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.![star > 5000][Gold]   ![There was an update last week][Green]
* [fortio](https://github.com/fortio/fortio) **star:868** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.![star > 100][Bronze]
* [ftp](https://github.com/jlaffaye/ftp) **star:522** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).![star > 100][Bronze]   ![There was an update last week][Green]
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:84** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [gNxI](https://github.com/google/gnxi) **star:98** A collection of tools for Network Management that use the gNMI and gNOI protocols.   ![There was an update last week][Green]
* [go-getter](https://github.com/hashicorp/go-getter) **star:723** Go library for downloading files or directories from various sources using a URL.![star > 100][Bronze]
* [go-stun](https://github.com/ccding/go-stun) **star:332** Go implementation of the STUN client (RFC 3489 and RFC 5389).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gobgp](https://github.com/osrg/gobgp) **star:1683** BGP implemented in the Go Programming Language.![star > 1000][Silver]   ![There was an update last week][Green]
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:15** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   ![It hasn't been updated in the last year][Yellow]
* [gopacket](https://github.com/google/gopacket) **star:2863** Go library for packet processing with libpcap bindings.![star > 1000][Silver]   ![There was an update last week][Green]
* [gopcap](https://github.com/akrennmair/gopcap) **star:353** Go wrapper for libpcap.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [goshark](https://github.com/sunwxg/goshark) **star:9** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   ![It hasn't been updated in the last year][Yellow]
* [gosnmp](https://github.com/soniah/gosnmp) **star:433** Native Go library for performing SNMP actions.![star > 100][Bronze]
* [gotcp](https://github.com/gansidui/gotcp) **star:414** Go package for quickly writing tcp applications.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [grab](https://github.com/cavaliercoder/grab) **star:545** Go package for managing file downloads.![star > 100][Bronze]
* [graval](https://github.com/koofr/graval) **star:24** Experimental FTP server framework.   ![It hasn't been updated in the last year][Yellow]
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3400** HTTPLabs let you inspect HTTP requests and forge responses.![star > 1000][Silver]
* [iplib](https://github.com/c-robinson/iplib) **star:24** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)
* [jazigo](https://github.com/udhos/jazigo) **star:123** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.![star > 100][Bronze]
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2237** KCP - Fast and Reliable ARQ Protocol.![star > 1000][Silver]
* [kcptun](https://github.com/xtaci/kcptun) **star:10638** Extremely simple & fast udp tunnel based on KCP protocol.![star > 5000][Gold]
* [lhttp](https://github.com/fanux/lhttp) **star:513** Powerful websocket framework, build your IM server more easily.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [linkio](https://github.com/ian-kent/linkio) **star:44** Network link speed simulation for Reader/Writer interfaces.   ![It hasn't been updated in the last year][Yellow]
* [llb](https://github.com/kirillDanshin/llb) **star:8** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   ![It hasn't been updated in the last year][Yellow]
* [mdns](https://github.com/hashicorp/mdns) **star:547** Simple mDNS (Multicast DNS) client/server library in Golang.![star > 100][Bronze]
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [NFF-Go](https://github.com/intel-go/nff-go) **star:659** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).![star > 100][Bronze]
* [packet](https://github.com/aerogo/packet) **star:27** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:361** Pure Go library for cross-platform local peer discovery using UDP multicast.![star > 100][Bronze]
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   ![It hasn't been updated in the last year][Yellow]
* [publicip](https://github.com/polera/publicip) **star:18** Package publicip returns your public facing IPv4 address (internet egress).   ![It hasn't been updated in the last year][Yellow]
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:2880** An implementation of the QUIC protocol in pure Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [raw](https://github.com/mdlayher/raw) **star:303** Package raw enables reading and writing data at the device driver level for a network interface.![star > 100][Bronze]
* [sftp](https://github.com/pkg/sftp) **star:733** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.![star > 100][Bronze]
* [ssh](https://github.com/gliderlabs/ssh) **star:1109** Higher-level API for building SSH servers (wraps crypto/ssh).![star > 1000][Silver]
* [sslb](https://github.com/eduardonunesp/sslb) **star:113** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.![star > 100][Bronze]
* [stun](https://github.com/go-rtc/stun) **star:273** Go implementation of RFC 5389 STUN protocol.![star > 100][Bronze]
* [tcp_server](https://github.com/firstrow/tcp_server) **star:284** Go library for building tcp servers faster.![star > 100][Bronze]
* [tspool](https://github.com/two/tspool) **star:5** A TCP Library use worker pool to improve performance and protect your server.
* [utp](https://github.com/anacrolix/utp) **star:149** Go uTP micro transport protocol implementation.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [water](https://github.com/songgao/water) **star:836** Simple TUN/TAP library.![star > 100][Bronze]   ![There was an update last week][Green]
* [webrtc](https://github.com/pions/webrtc) **star:2188** A pure Go implementation of the WebRTC API.![star > 1000][Silver]   ![There was an update last week][Green]
* [winrm](https://github.com/masterzen/winrm) **star:210** Go WinRM client to remotely execute commands on Windows machines.![star > 100][Bronze]
* [xtcp](https://github.com/xfxdev/xtcp) **star:82** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.

### HTTP Clients

*Libraries for making HTTP requests.*

* [gentleman](https://github.com/h2non/gentleman) **star:678** Full-featured plugin-driven HTTP client library.![star > 100][Bronze]
* [goreq](https://github.com/smallnest/goreq) **star:98** Enhanced simplified HTTP client based on gorequest.   ![It hasn't been updated in the last year][Yellow]
* [grequests](https://github.com/levigross/grequests) **star:1410** A Go "clone" of the great and famous Requests library.![star > 1000][Silver]
* [heimdall](https://github.com/gojektech/heimdall) **star:1068** An enchanced http client with retry and hystrix capabilities.![star > 1000][Silver]
* [pester](https://github.com/sethgrid/pester) **star:328** Go HTTP client calls with retries, backoff, and concurrency.![star > 100][Bronze]
* [rq](https://github.com/ddo/rq) **star:26** A nicer interface for golang stdlib HTTP client.   ![There was an update last week][Green]
* [sling](https://github.com/dghubble/sling) **star:995** Sling is a Go HTTP client library for creating and sending API requests.![star > 100][Bronze]

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl) **star:640** Go bindings for OpenGL (generated via glow).![star > 100][Bronze]
* [glfw](https://github.com/go-gl/glfw) **star:729** Go bindings for GLFW 3.![star > 100][Bronze]
* [goxjs/gl](https://github.com/goxjs/gl) **star:130** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).![star > 100][Bronze]
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:58** Go cross-platform glfw library for creating an OpenGL context and receiving events.
* [mathgl](https://github.com/go-gl/mathgl) **star:289** Pure Go math package specialized for 3D math, with inspiration from GLM.![star > 100][Bronze]

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:1** A simple ORM for Google/Firebase Cloud Firestore.
* [go-pg](https://github.com/go-pg/pg) **star:2958** PostgreSQL ORM with focus on PostgreSQL specific features and performance.![star > 1000][Silver]   ![There was an update last week][Green]
* [go-queryset](https://github.com/jirfag/go-queryset) **star:447** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:233** A flexible and powerful SQL string builder library plus a zero-config ORM.![star > 100][Bronze]
* [go-store](https://github.com/gosuri/go-store) **star:93** Simple and fast Redis backed key-value store library for Go.   ![It hasn't been updated in the last year][Yellow]
* [GORM](https://github.com/jinzhu/gorm) **star:14583** The fantastic ORM library for Golang, aims to be developer friendly.![star > 5000][Gold]   ![There was an update last week][Green]
* [gorp](https://github.com/go-gorp/gorp) **star:3073** Go Relational Persistence, ORM-ish library for Go.![star > 1000][Silver]
* [grimoire](https://github.com/Fs02/grimoire) **star:112** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).![star > 100][Bronze]
* [lore](https://github.com/abrahambotros/lore) **star:4** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   ![It hasn't been updated in the last year][Yellow]
* [Marlow](https://github.com/dadleyy/marlow) **star:64** Generated ORM from project structs for compile time safety assurances.
* [pop/soda](https://github.com/gobuffalo/pop) **star:674** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.![star > 100][Bronze]   ![There was an update last week][Green]
* [QBS](https://github.com/coocood/qbs) **star:539** Stands for Query By Struct. A Go ORM.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]   ![Contains Chinese documents][CN]
* [reform](https://github.com/go-reform/reform) **star:796** Better ORM for Go, based on non-empty interfaces and code generation.![star > 100][Bronze]   ![There was an update last week][Green]
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2261** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.![star > 1000][Silver]   ![There was an update last week][Green]
* [upper.io/db](https://github.com/upper/db) **star:1846** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.![star > 1000][Silver]   ![There was an update last week][Green]
* [Xorm](https://github.com/go-xorm/xorm) **star:5171** Simple and powerful ORM for Go.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Zoom](https://github.com/albrow/zoom) **star:239** Blazing-fast datastore and querying engine built on Redis.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:12546** Go dependency tool.![star > 5000][Gold]
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [gigo](https://github.com/LyricalSecurity/gigo) **star:197** PIP-like dependency tool for golang, with support for private repositories and hashes.![star > 100][Bronze]
* [glide](https://github.com/Masterminds/glide) **star:7777** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.![star > 5000][Gold]
* [godep](https://github.com/tools/godep) **star:5650** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.![star > 5000][Gold]   ![It hasn't been updated in the last year][Yellow]
* [gom](https://github.com/mattn/gom) **star:1350** Go Manager - bundle for go.![star > 1000][Silver]   ![There was an update last week][Green]
* [goop](https://github.com/nitrous-io/goop) **star:777** Simple dependency manager for Go (golang), inspired by Bundler.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gop](https://github.com/lunny/gop) **star:50** Build and manage your Go applications out of GOPATH.   ![Contains Chinese documents][CN]
* [gopm](https://github.com/gpmgo/gopm) **star:2352** Go Package Manager.![star > 1000][Silver]   ![There was an update last week][Green]
* [govendor](https://github.com/kardianos/govendor) **star:4727** Go Package Manager. Go vendor tool that works with the standard vendor file.![star > 1000][Silver]
* [gpm](https://github.com/pote/gpm) **star:1205** Barebones dependency manager for Go.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** Minimal dependency version using Git.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:87** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.   ![There was an update last week][Green]
* [nut](https://github.com/jingweno/nut) **star:245** Vendor Go dependencies.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [VenGO](https://github.com/DamnWidget/VenGO) **star:115** create and manage exportable isolated go virtual environments.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Query Language

* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:858** A simple Go package to Query over JSON Data.![star > 100][Bronze]
* [graphql](https://github.com/tmc/graphql) **star:51** graphql parser + utilities.   ![It hasn't been updated in the last year][Yellow]
* [graphql](https://github.com/neelance/graphql-go) **star:2757** GraphQL server with a focus on ease of use.![star > 1000][Silver]
* [graphql-go](https://github.com/graphql-go/graphql) **star:5194** Implementation of GraphQL for Go.![star > 5000][Gold]   ![There was an update last week][Green]
* [jsonql](https://github.com/elgs/jsonql) **star:201** JSON query expression library in Golang.![star > 100][Bronze]
* [jsonslice](https://github.com/bhmj/jsonslice) **star:23** Jsonpath queries with advanced filters.
* [rql](https://github.com/a8m/rql) **star:111** Resource Query Language for REST API.![star > 100][Bronze]   ![There was an update last week][Green]

## Resource Embedding

* [esc](https://github.com/mjibson/esc) **star:467** Embeds files into Go programs and provides http.FileSystem interfaces to them.![star > 100][Bronze]
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:421** Simple tool to embed files in go with focus on "customization" and ease to use.![star > 100][Bronze]
* [go-embed](https://github.com/pyros2097/go-embed) **star:14** Generates go code to embed resource files into your library or executable.   ![It hasn't been updated in the last year][Yellow]
* [go-resources](https://github.com/omeid/go-resources) **star:154** Unfancy resources embedding with Go.![star > 100][Bronze]
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1642** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.![star > 1000][Silver]
* [packr](https://github.com/gobuffalo/packr) **star:2097** The simple and easy way to embed static files into Go binaries.![star > 1000][Silver]
* [statics](https://github.com/go-playground/statics) **star:53** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   ![It hasn't been updated in the last year][Yellow]
* [statik](https://github.com/rakyll/statik) **star:2089** Embeds static files into a Go executable.![star > 1000][Silver]   ![There was an update last week][Green]
* [templify](https://github.com/wlbr/templify) **star:19** Embed external template files into Go code to create single file binaries.
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:650** Generates a vfsdata.go file that statically implements the given virtual filesystem.![star > 100][Bronze]

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [assocentity](https://github.com/ndabAP/assocentity) **star:3** Package assocentity returns the average distance from words to a given entity.   ![There was an update last week][Green]
* [bradleyterry](https://github.com/seanhagen/bradleyterry)  Provides a Bradley-Terry Model for pairwise comparisons.
* [chart](https://github.com/vdobler/chart) **star:580** Simple Chart Plotting library for Go. Supports many graphs types.![star > 100][Bronze]
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:75** Dataframes for Go for machine-learning and statistics (similar to pandas).   ![There was an update last week][Green]
* [evaler](https://github.com/soniah/evaler) **star:40** Simple floating point arithmetic expression evaluator.   ![It hasn't been updated in the last year][Yellow]
* [ewma](https://github.com/VividCortex/ewma) **star:266** Exponentially-weighted moving averages.![star > 100][Bronze]
* [geom](https://github.com/skelterjohn/geom) **star:40** 2D geometry for golang.   ![It hasn't been updated in the last year][Yellow]
* [go-dsp](https://github.com/mjibson/go-dsp) **star:626** Digital Signal Processing for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-fn](https://github.com/ematvey/go-fn) **star:11** Mathematical functions written in Go language, that are not covered by math pkg.   ![It hasn't been updated in the last year][Yellow]
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** Graph theory algorithms written in "Go" language.   ![It hasn't been updated in the last year][Yellow]
* [gocomplex](https://github.com/varver/gocomplex) **star:5** Complex number library for the Go programming language.   ![It hasn't been updated in the last year][Yellow]
* [goent](https://github.com/kzahedi/goent) **star:13** GO Implementation of Entropy Measures.
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:126** Approximate histograms for data streams.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gonum](https://github.com/gonum/gonum) **star:2936** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.![star > 1000][Silver]   ![There was an update last week][Green]
* [gonum/plot](https://github.com/gonum/plot) **star:1210** gonum/plot provides an API for building and drawing plots in Go.![star > 1000][Silver]
* [goraph](https://github.com/gyuho/goraph) **star:597** Pure Go graph theory library(data structure, algorith visualization).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gosl](https://github.com/cpmech/gosl) **star:1304** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.![star > 1000][Silver]
* [GoStats](https://github.com/OGFris/GoStats) **star:9** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.
* [graph](https://github.com/yourbasic/graph) **star:232** Library of basic graph algorithms.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [ode](https://github.com/ChristopherRabotin/ode) **star:10** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   ![It hasn't been updated in the last year][Yellow]
* [orb](https://github.com/paulmach/orb) **star:196** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.![star > 100][Bronze]
* [pagerank](https://github.com/alixaxel/pagerank) **star:48** Weighted PageRank algorithm implemented in Go.
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:5** Tiny linear interpolation library.
* [PiHex](https://github.com/claygod/PiHex) **star:9** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.
* [sparse](https://github.com/james-bowman/sparse) **star:68** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.
* [stats](https://github.com/montanaflynn/stats) **star:1338** Statistics package with common functions missing from the Golang standard library.![star > 1000][Silver]
* [streamtools](https://github.com/nytlabs/streamtools) **star:1313** general purpose, graphical tool for dealing with streams of data.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:65** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.
* [triangolatte](https://github.com/tchayen/triangolatte) **star:11** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.

## Security

*Libraries that are used to help make your application more secure.*

* [acmetool](https://github.com/hlandau/acme) **star:1694** ACME (Let's Encrypt) client tool with automatic renewal.![star > 1000][Silver]
* [acra](https://github.com/cossacklabs/acra) **star:450** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.![star > 100][Bronze]
* [argon2pw](https://github.com/raja/argon2pw) **star:73** Argon2 password hash generation with constant-time password comparison.
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [BadActor](https://github.com/jaredfolkins/badactor) **star:246** In-memory, application-driven jailer built in the spirit of fail2ban.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:1821** Tool and library to remotely hack RTSP streams from surveillance cameras.![star > 1000][Silver]
* [certificates](https://github.com/mvmaasakkers/certificates) **star:6** An opinionated tool for generating tls certificates.
* [go-yara](https://github.com/hillu/go-yara) **star:132** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".![star > 100][Bronze]
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:29** A probably paranoid package for securely hashing and encrypting passwords.
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [jwc](https://github.com/khezen/jwc) **star:5** JSON Web Cryptography library.
* [lego](https://github.com/xenolf/lego) **star:3471** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).![star > 1000][Silver]   ![There was an update last week][Green]
* [memguard](https://github.com/awnumar/memguard) **star:1455** A pure Go library for handling sensitive values in memory.![star > 1000][Silver]   ![There was an update last week][Green]
* [nacl](https://github.com/kevinburke/nacl) **star:450** Go implementation of the NaCL set of API's.![star > 100][Bronze]
* [passlib](https://github.com/hlandau/passlib) **star:226** Futureproof password hashing library.![star > 100][Bronze]
* [secure](https://github.com/unrolled/secure) **star:1200** HTTP middleware for Go that facilitates some quick security wins.![star > 1000][Silver]
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:156** Scrypt package with a simple, obvious API and automatic cost calibration built-in.![star > 100][Bronze]
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:195** encrypt/decrypt using ssh keys.![star > 100][Bronze]
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:7** SSL certificates made easy with a high level wrapper around acme/autocert.   ![There was an update last week][Green]

## Serialization

*Libraries and tools for binary serialization.*

* [asn1](https://github.com/PromonLogicalis/asn1) **star:40** Asn.1 BER and DER encoding library for golang.
* [bambam](https://github.com/glycerine/bambam) **star:60** generator for Cap'n Proto schemas from go.   ![It hasn't been updated in the last year][Yellow]
* [bel](https://github.com/32leaves/bel) **star:5** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.
* [binstruct](https://github.com/ghostiam/binstruct) **star:7** Golang binary decoder for mapping data into the structure.
* [colfer](https://github.com/pascaldekloe/colfer) **star:471** Code generation for the Colfer binary format.![star > 100][Bronze]
* [csvutil](https://github.com/jszwec/csvutil) **star:302** High Performance, idiomatic CSV record encoding and decoding to native Go structures.![star > 100][Bronze]
* [fwencoder](https://github.com/o1egl/fwencoder) **star:6** Fixed width file parser (encoding and decoding library) for Go.   ![It hasn't been updated in the last year][Yellow]
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:272** Cap'n Proto library and parser for go.![star > 100][Bronze]
* [go-codec](https://github.com/ugorji/go) **star:1229** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.![star > 1000][Silver]
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:2935** Protocol Buffers for Go with Gadgets.![star > 1000][Silver]   ![There was an update last week][Green]
* [goprotobuf](https://github.com/golang/protobuf) **star:5055** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.![star > 5000][Gold]   ![There was an update last week][Green]
* [jsoniter](https://github.com/json-iterator/go) **star:5510** High-performance 100% compatible drop-in replacement of "encoding/json".![star > 5000][Gold]   ![There was an update last week][Green]
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:2400** Go library for decoding generic map values into native Go structures.![star > 1000][Silver]   ![There was an update last week][Green]
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:118** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.![star > 100][Bronze]
* [structomap](https://github.com/tuvistavie/structomap) **star:95** Library to easily and dynamically generate maps from static structures.

## Server Applications

* [algernon](https://github.com/xyproto/algernon) **star:1586** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.![star > 1000][Silver]
* [Caddy](https://github.com/mholt/caddy) **star:23104** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.![star > 5000][Gold]   ![There was an update last week][Green]
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [devd](https://github.com/cortesi/devd) **star:2802** Local webserver for developers.![star > 1000][Silver]
* [discovery](https://github.com/Bilibili/discovery) **star:670** A registry for resilient mid-tier load balancing and failover.![star > 100][Bronze]
* [etcd](https://github.com/coreos/etcd) **star:26350** Highly-available key value store for shared configuration and service discovery.![star > 5000][Gold]   ![There was an update last week][Green]
* [Fider](https://github.com/getfider/fider) **star:791** Fider is an open platform to collect and organize customer feedback.![star > 100][Bronze]
* [Flagr](https://github.com/checkr/flagr) **star:821** Flagr is an open-source feature flagging and A/B testing service.![star > 100][Bronze]   ![There was an update last week][Green]
* [flipt](https://github.com/markphelps/flipt) **star:985** A self contained feature flag solution written in Go and Vue.js![star > 100][Bronze]   ![There was an update last week][Green]
* [jackal](https://github.com/ortuman/jackal) **star:711** An XMPP server written in Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [minio](https://github.com/minio/minio) **star:17238** Minio is a distributed object storage server.![star > 5000][Gold]   ![There was an update last week][Green]
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:5** Nginx log parser and exporter to Prometheus.
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:5** Stream database events from PostgreSQL to Kafka.
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [RoadRunner](https://github.com/spiral/roadrunner) **star:3255** High-performance PHP application server, load-balancer and process manager.![star > 1000][Silver]   ![There was an update last week][Green]
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.

## Template Engines

*Libraries and tools for templating and lexing.*

* [ace](https://github.com/yosssi/ace) **star:761** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [amber](https://github.com/eknkc/amber) **star:822** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.![star > 100][Bronze]
* [damsel](https://github.com/dskinner/damsel) **star:20** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   ![It hasn't been updated in the last year][Yellow]
* [ego](https://github.com/benbjohnson/ego) **star:416** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.![star > 100][Bronze]
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:13** Tiny wrapper around html/template to allow for easy file-based template inheritance.
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:294** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).![star > 100][Bronze]
* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3070** PDF document generator with high level support for text, drawing and images.![star > 1000][Silver]
* [goview](https://github.com/foolin/goview) **star:43** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
* [hero](https://github.com/shiyanhui/hero) **star:1200** Hero is a handy, fast and powerful go template engine.![star > 1000][Silver]   ![Contains Chinese documents][CN]
* [jet](https://github.com/CloudyKit/jet) **star:581** Jet template engine.![star > 100][Bronze]   ![There was an update last week][Green]
* [kasia.go](https://github.com/ziutek/kasia.go) **star:70** Templating system for HTML and other text documents - go implementation.   ![It hasn't been updated in the last year][Yellow]
* [liquid](https://github.com/osteele/liquid) **star:83** Go implementation of Shopify Liquid templates.
* [mustache](https://github.com/hoisie/mustache) **star:967** Go implementation of the Mustache template language.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [pongo2](https://github.com/flosch/pongo2) **star:1485** Django-like template-engine for Go.![star > 1000][Silver]
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1396** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.![star > 1000][Silver]
* [raymond](https://github.com/aymerick/raymond) **star:338** Complete handlebars implementation in Go.![star > 100][Bronze]
* [Razor](https://github.com/sipin/gorazor) **star:696** Razor view engine for Golang.![star > 100][Bronze]
* [Soy](https://github.com/robfig/soy) **star:144** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).![star > 100][Bronze]
* [velvet](https://github.com/gobuffalo/velvet) **star:64** Complete handlebars implementation in Go.   ![It hasn't been updated in the last year][Yellow]

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [assert](https://github.com/go-playground/assert) **star:13** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   ![It hasn't been updated in the last year][Yellow]
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Extensions to Go's `testing/iotest` package.   ![It hasn't been updated in the last year][Yellow]
    * [baloo](https://github.com/h2non/baloo) **star:646** Expressive and versatile end-to-end HTTP API testing made easy.![star > 100][Bronze]
    * [biff](https://github.com/fulldump/biff) **star:6** Bifurcation testing framework, BDD compatible.   ![It hasn't been updated in the last year][Yellow]
    * [bro](https://github.com/marioidival/bro) **star:26** Watch files in directory and run tests for them.   ![It hasn't been updated in the last year][Yellow]
    * [charlatan](https://github.com/percolate/charlatan) **star:189** Tool to generate fake interface implementations for tests.![star > 100][Bronze]
    * [commander](https://github.com/SimonBaeumer/commander) **star:32** Tool for testing cli applications on windows, linux and osx.
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:83** Simple snapshot testing addon for your test framework.   ![There was an update last week][Green]
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:85** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.
    * [dsunit](https://github.com/viant/dsunit) **star:25** Datastore testing for SQL, NoSQL, structured files.   ![There was an update last week][Green]
    * [endly](https://github.com/viant/endly) **star:91** Declarative end to end functional testing.   ![There was an update last week][Green]
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:1** HTTP client testing framework.
    * [frisby](https://github.com/verdverm/frisby) **star:250** REST API testing framework.![star > 100][Bronze]
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:196** Tool for viewing test coverage in terminal.![star > 100][Bronze]
    * [go-cmp](https://github.com/google/go-cmp) **star:1139** Package for comparing Go values in tests.![star > 1000][Silver]
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:291** Mutation testing for Go source code.![star > 100][Bronze]   ![There was an update last week][Green]
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:55** Extremely flexible golang deep comparison, extends the go testing package.   ![There was an update last week][Green]
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:329** Record and replay your HTTP interactions for fast, deterministic and accurate tests.![star > 100][Bronze]
    * [goblin](https://github.com/franela/goblin) **star:624** Mocha like testing framework fo Go.![star > 100][Bronze]
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [gocrest](https://github.com/corbym/gocrest) **star:8** Composable hamcrest-like matchers for Go assertions.   ![It hasn't been updated in the last year][Yellow]
    * [godog](https://github.com/DATA-DOG/godog) **star:748** Cucumber or Behat like BDD framework for Go.![star > 100][Bronze]
    * [gofight](https://github.com/appleboy/gofight) **star:255** API Handler Testing for Golang Router framework.![star > 100][Bronze]
    * [gogiven](https://github.com/corbym/gogiven) **star:7** YATSPEC-like BDD testing framework for Go.   ![It hasn't been updated in the last year][Yellow]
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:30** library created for testing JSON against patterns.
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [GoSpec](https://github.com/orfjackal/gospec) **star:111** BDD-style testing framework for the Go programming language.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [gospecify](https://github.com/stesla/gospecify) **star:51** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   ![It hasn't been updated in the last year][Yellow]
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   ![It hasn't been updated in the last year][Yellow]
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:118** A collection of packages to augment the go testing package and support common patterns.![star > 100][Bronze]
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:26** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   ![It hasn't been updated in the last year][Yellow]
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1138** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.![star > 1000][Silver]   ![There was an update last week][Green]
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:21** Package for verifying that your JSON payloads are serialized correctly.
    * [restit](https://github.com/yookoala/restit) **star:49** Go micro framework to help writing RESTful API integration test.   ![It hasn't been updated in the last year][Yellow]
    * [testcase](https://github.com/adamluzsi/testcase) **star:8** Idiomatic testing framework for Behavior Driven Development.
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:323** A helper for Rails' like test fixtures to test database applications.![star > 100][Bronze]
    * [Testify](https://github.com/stretchr/testify) **star:8088** Sacred extension to the standard go testing package.![star > 5000][Gold]   ![There was an update last week][Green]
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.
    * [Tt](https://github.com/vcaesar/tt) **star:5** Simple and colorful test tools.
    * [wstest](https://github.com/posener/wstest) **star:62** Websocket client for unit-testing a websocket http.Handler.   ![It hasn't been updated in the last year][Yellow]

* Mock
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:360** Tool for generating self-contained mock objects.![star > 100][Bronze]   ![There was an update last week][Green]
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:1728** Mock SQL driver for testing database interactions.![star > 1000][Silver]
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:160** Single transaction based database driver mainly for testing purposes.![star > 100][Bronze]   ![There was an update last week][Green]
    * [gock](https://github.com/h2non/gock) **star:815** Versatile HTTP mocking made easy.![star > 100][Bronze]
    * [gomock](https://github.com/golang/mock) **star:2821** Mocking framework for the Go programming language.![star > 1000][Silver]
    * [govcr](https://github.com/seborama/govcr) **star:82** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   ![There was an update last week][Green]
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1435** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.![star > 1000][Silver]
    * [httpmock](https://github.com/jarcoal/httpmock) **star:577** Easy mocking of HTTP responses from external resources.![star > 100][Bronze]
    * [minimock](https://github.com/gojuno/minimock) **star:259** Mock generator for Go interfaces.![star > 100][Bronze]
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   ![It hasn't been updated in the last year][Yellow]

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:2874** Randomized testing system.![star > 1000][Silver]   ![There was an update last week][Green]
    * [gofuzz](https://github.com/google/gofuzz) **star:530** Library for populating go objects with random values.![star > 100][Bronze]
    * [Tavor](https://github.com/zimmski/tavor) **star:212** Generic fuzzing and delta-debugging framework.![star > 100][Bronze]

* Selenium and browser control tools.
    * [cdp](https://github.com/mafredri/cdp) **star:352** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.![star > 100][Bronze]
    * [chromedp](https://github.com/knq/chromedp) **star:3558** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.![star > 1000][Silver]
    * [ggr](https://github.com/aerokube/ggr) **star:208** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.![star > 100][Bronze]
    * [selenoid](https://github.com/aerokube/selenoid) **star:1219** alternative Selenium hub server that launches browsers within containers.![star > 1000][Silver]   ![There was an update last week][Green]

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:383** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.![star > 100][Bronze]

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [align](https://github.com/Guitarbum722/align) **star:58** A general purpose application that aligns text.   ![It hasn't been updated in the last year][Yellow]
    * [allot](https://github.com/sbstjn/allot) **star:34** Placeholder and wildcard text parsing for CLI tools and bots.
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   ![It hasn't been updated in the last year][Yellow]
    * [blackfriday](https://github.com/russross/blackfriday) **star:3877** Markdown processor in Go.![star > 1000][Silver]
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1242** HTML Sanitizer.![star > 1000][Silver]
    * [codetree](https://github.com/aerogo/codetree) **star:7** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.
    * [colly](https://github.com/asciimoo/colly) **star:8325** Fast and Elegant Scraping Framework for Gophers.![star > 5000][Gold]   ![There was an update last week][Green]
    * [commonregex](https://github.com/mingrammer/commonregex) **star:552** A collection of common regular expressions for Go.![star > 100][Bronze]
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:289** Web scraping Framework to turn websites into structured data.![star > 100][Bronze]
    * [did](https://github.com/ockam-network/did) **star:23** DID (Decentralized Identifiers) Parser and Stringer in Go.
    * [doi](https://github.com/hscells/doi) **star:4** Document object identifier (doi) parser in Go.   ![It hasn't been updated in the last year][Yellow]
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:37** Editorconfig file parser and manipulator for Go.   ![There was an update last week][Green]
    * [enca](https://github.com/endeveit/enca) **star:7** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   ![It hasn't been updated in the last year][Yellow]
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.
    * [genex](https://github.com/alixaxel/genex) **star:51** Count and expand Regular Expressions into all matching Strings.
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:25** Fixed-width text formatting (encoder/decoder with reflection).
    * [go-humanize](https://github.com/dustin/go-humanize) **star:1892** Formatters for time, numbers, and memory size to human readable format.![star > 1000][Silver]
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:94** NMEA parser library for the Go language.   ![There was an update last week][Green]
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:210** Functions to get fixed width of the character or string.![star > 100][Bronze]
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:29** Make pretty slug with multiple languages support.   ![It hasn't been updated in the last year][Yellow]
    * [go-toml](https://github.com/pelletier/go-toml) **star:606** Go library for the TOML format with query support and handy cli tools.![star > 100][Bronze]
    * [go-vcard](https://github.com/emersion/go-vcard) **star:24** Parse and format vCard.
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:41** Zero-width character detection and removal for Go.
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1094** Parse RSS and Atom feeds in Go.![star > 1000][Silver]
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:296** Parses the Graphviz DOT language.![star > 100][Bronze]
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [gonameparts](https://github.com/polera/gonameparts) **star:29** Parses human names into individual name parts.   ![It hasn't been updated in the last year][Yellow]
    * [goq](https://github.com/andrewstuart/goq) **star:146** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).![star > 100][Bronze]
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:7536** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.![star > 5000][Gold]
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:36** Library for generating random strings from regular expressions.   ![It hasn't been updated in the last year][Yellow]
    * [gotext](https://github.com/leonelquinteros/gotext) **star:230** GNU gettext utilities for Go.![star > 100][Bronze]
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:44** Functions to determine the natural language of a unicode text.   ![It hasn't been updated in the last year][Yellow]
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:129** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.![star > 100][Bronze]
    * [inject](https://github.com/facebookgo/inject) **star:1139** Package inject provides a reflect based injector.![star > 1000][Silver]
    * [ltsv](https://github.com/Wing924/ltsv) **star:2** High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.
    * [mxj](https://github.com/clbanning/mxj) **star:330** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.![star > 100][Bronze]
    * [sdp](https://github.com/gortc/sdp) **star:69** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   ![There was an update last week][Green]
    * [sh](https://github.com/mvdan/sh) **star:1984** Shell parser and formatter.![star > 1000][Silver]
    * [slug](https://github.com/gosimple/slug) **star:370** URL-friendly slugify with multiple languages support.![star > 100][Bronze]
    * [Slugify](https://github.com/avelino/slugify) **star:26** Go slugify application that handles string.   ![It hasn't been updated in the last year][Yellow]
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:5** A syndication feed for Atom 1.0 and RSS 2.0.   ![It hasn't been updated in the last year][Yellow]
    * [toml](https://github.com/BurntSushi/toml) **star:2767** TOML configuration format (encoder/decoder with reflection).![star > 1000][Silver]
* Utility
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:15** A sanitization-based swear filter for Go.
    * [gotabulate](https://github.com/bndr/gotabulate) **star:198** Easily pretty-print your tabular data with Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [kace](https://github.com/codemodus/kace) **star:12** Common case conversions covering common initialisms.
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:6** string argument parser that understands quotes and backslashes.   ![It hasn't been updated in the last year][Yellow]
    * [parth](https://github.com/codemodus/parth) **star:31** URL path segmentation parsing.
    * [radix](https://github.com/yourbasic/radix) **star:144** fast string sorting algorithm.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
    * [Tagify](https://github.com/zoomio/tagify) **star:1** Produces a set of tags from given source.
    * [TySug](https://github.com/Dynom/TySug) **star:3** Alternative suggestions with respect to keyboard layouts.
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:17** Convert xml or json to go struct.
    * [xurls](https://github.com/mvdan/xurls) **star:459** Extract urls from text.![star > 100][Bronze]   ![There was an update last week][Green]

## Third-party APIs

*Libraries for accessing third party APIs.*

* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:39** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   ![It hasn't been updated in the last year][Yellow]
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:985** Go client library for the Twitter 1.1 API.![star > 100][Bronze]
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5006** The official AWS SDK for the Go programming language.![star > 5000][Gold]   ![There was an update last week][Green]
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:16** Go library for accessing the BreweryDB API.   ![It hasn't been updated in the last year][Yellow]
* [cachet](https://github.com/andygrunwald/cachet) **star:66** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   ![It hasn't been updated in the last year][Yellow]
* [circleci](https://github.com/jszwedko/go-circleci) **star:41** Go client library for interacting with CircleCI's API.
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Go client library for interfacing with the Clarifai API.   ![It hasn't been updated in the last year][Yellow]
* [codeship-go](https://github.com/codeship/codeship-go) **star:14** Go client library for interacting with Codeship's API v2.
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.
* [discordgo](https://github.com/bwmarrin/discordgo) **star:967** Go bindings for the Discord Chat API.![star > 100][Bronze]
* [ethrpc](https://github.com/onrik/ethrpc) **star:169** Go bindings for Ethereum JSON RPC API.![star > 100][Bronze]
* [facebook](https://github.com/huandu/facebook) **star:768** Go Library that supports the Facebook Graph API.![star > 100][Bronze]
* [fcm](https://github.com/maddevsio/fcm) **star:32** Go library for Firebase Cloud Messaging.
* [gads](https://github.com/emiddleton/gads) **star:43** Google Adwords Unofficial API.
* [gami](https://github.com/bit4bit/gami) **star:26** Go library for Asterisk Manager Interface.   ![It hasn't been updated in the last year][Yellow]
* [gcm](https://github.com/Aorioli/gcm) **star:29** Go library for Google Cloud Messaging.   ![It hasn't been updated in the last year][Yellow]
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:306** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.![star > 100][Bronze]
* [github](https://github.com/google/go-github) **star:4749** Go library for accessing the GitHub REST API v3.![star > 1000][Silver]   ![There was an update last week][Green]
* [githubql](https://github.com/shurcooL/githubql) **star:500** Go library for accessing the GitHub GraphQL API v4.![star > 100][Bronze]
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   ![It hasn't been updated in the last year][Yellow]
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:9** Tiny Go client for HackerNews API.   ![It hasn't been updated in the last year][Yellow]
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:13** Go client library for [imgur](https://imgur.com)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:570** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)![star > 100][Bronze]
* [go-marathon](https://github.com/gambol99/go-marathon) **star:189** Go library for interacting with Mesosphere's Marathon PAAS.![star > 100][Bronze]
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:16** Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).   ![It hasn't been updated in the last year][Yellow]
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:8** Go client library for the SPTrans Olho Vivo API.   ![It hasn't been updated in the last year][Yellow]
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [go-trending](https://github.com/andygrunwald/go-trending) **star:100** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.![star > 100][Bronze]
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:17** Go client for interacting with the Twitch v3 API.   ![It hasn't been updated in the last year][Yellow]
* [go-twitter](https://github.com/dghubble/go-twitter) **star:709** Go client library for the Twitter v1.1 APIs.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:20** Go client library for the [Unsplash.com](https://unsplash.com) API.   ![It hasn't been updated in the last year][Yellow]
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:38** Go client for the xkcd API.   ![It hasn't been updated in the last year][Yellow]
* [golyrics](https://github.com/mamal72/golyrics) **star:30** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   ![It hasn't been updated in the last year][Yellow]
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:1** Go library MalShare API [malshare.com](http://www.malshare.com/)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:36** Go MusicBrainz WS2 client library.
* [google](https://github.com/google/google-api-go-client) **star:1912** Auto-generated Google APIs for Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:12** Simple wrapper for easy google analytics reporting.   ![It hasn't been updated in the last year][Yellow]
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:1770** Google Cloud APIs Go Client Library.![star > 1000][Silver]   ![There was an update last week][Green]
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).   ![It hasn't been updated in the last year][Yellow]
* [gostorm](https://github.com/jsgilmore/gostorm) **star:118** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hipchat](https://github.com/andybons/hipchat) **star:110** This project implements a golang client library for the Hipchat API.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:113** A golang package to communicate with HipChat over XMPP.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:52** Go client for the [Internet Game Database API](https://api.igdb.com/).
* [Medium](https://github.com/Medium/medium-sdk-go) **star:114** Golang SDK for Medium's OAuth2 API.![star > 100][Bronze]
* [megos](https://github.com/andygrunwald/megos) **star:57** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   ![It hasn't been updated in the last year][Yellow]
* [minio-go](https://github.com/minio/minio-go) **star:716** Minio Go Library for Amazon S3 compatible cloud storage.![star > 100][Bronze]
* [mixpanel](https://github.com/dukex/mixpanel) **star:29** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
* [patreon-go](https://github.com/mxpv/patreon-go) **star:18** Go library for Patreon API.
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:301** Wrapper for PayPal payment API.![star > 100][Bronze]   ![There was an update last week][Green]
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** The Playlyfe Rest API Go SDK.   ![It hasn't been updated in the last year][Yellow]
* [pushover](https://github.com/gregdel/pushover) **star:57** Go wrapper for the Pushover API.
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   ![It hasn't been updated in the last year][Yellow]
* [shopify](https://github.com/rapito/go-shopify) **star:19** Go Library to make CRUD request to the Shopify API.   ![It hasn't been updated in the last year][Yellow]
* [simples3](https://github.com/rhnvrm/simples3) **star:9** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.
* [slack](https://github.com/nlopes/slack) **star:2403** Slack API in Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   ![It hasn't been updated in the last year][Yellow]
* [spotify](https://github.com/rapito/go-spotify) **star:17** Go Library to access Spotify WEB API.   ![It hasn't been updated in the last year][Yellow]
* [steam](https://github.com/sostronk/go-steam) **star:14** Go Library to interact with Steam game servers.   ![It hasn't been updated in the last year][Yellow]
* [stripe](https://github.com/stripe/stripe-go) **star:938** Go client for the Stripe API.![star > 100][Bronze]   ![There was an update last week][Green]
* [textbelt](https://github.com/dietsche/textbelt) **star:14** Go client for the textbelt.com txt messaging API.   ![It hasn't been updated in the last year][Yellow]
* [TheMovieDb](https://github.com/jbrodriguez/go-tmdb) **star:12** Simple golang package to communicate with [themoviedb.org](https://themoviedb.org).   ![It hasn't been updated in the last year][Yellow]
* [translate](https://github.com/poorny/translate) **star:27** Go online translation package.   ![It hasn't been updated in the last year][Yellow]
* [Trello](https://github.com/adlio/trello) **star:101** Go wrapper for the Trello API.![star > 100][Bronze]
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   ![It hasn't been updated in the last year][Yellow]
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:35** Go wrapper and command-line client for the Uptime Robot v2 API.   ![There was an update last week][Green]
* [webhooks](https://github.com/go-playground/webhooks) **star:348** Webhook receiver for GitHub and Bitbucket.![star > 100][Bronze]
* [wit-go](https://github.com/wit-ai/wit-go) **star:47** Go client for wit.ai HTTP API.
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:12** Go wrapper for the YNAB API.   ![There was an update last week][Green]
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.

## Utilities

*General utilities and tools to make your life easier.*

* [abutil](https://github.com/bahlo/abutil) **star:51** Collection of often-used Golang helpers.   ![It hasn't been updated in the last year][Yellow]
* [apm](https://github.com/topfreegames/apm) **star:129** Process manager for Golang applications with an HTTP API.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [backscanner](https://github.com/icza/backscanner) **star:8** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.
* [blank](https://github.com/Henry-Sarabia/blank) **star:1** Verify or remove blanks and whitespace from strings.   ![There was an update last week][Green]
* [boilr](https://github.com/tmrts/boilr) **star:935** Blazingly fast CLI tool for creating projects from boilerplate templates.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [chyle](https://github.com/antham/chyle) **star:106** Changelog generator using a git repository with multiple configuration possibilities.![star > 100][Bronze]
* [circuit](https://github.com/cep21/circuit) **star:328** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.![star > 100][Bronze]
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:787** Circuit Breakers in Go.![star > 100][Bronze]
* [clockwork](https://github.com/jonboulle/clockwork) **star:219** A simple fake clock for golang.![star > 100][Bronze]
* [command](https://github.com/txgruppi/command) **star:9** Command pattern for Go with thread safe serial and parallel dispatcher.   ![It hasn't been updated in the last year][Yellow]
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:37** Universal multi-workstation clipboard that uses S3 like backend for the storage.
* [ctop](https://github.com/bcicen/ctop) **star:8803** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.![star > 5000][Gold]
* [ctxutil](https://github.com/posener/ctxutil) **star:7** A collection of utility functions for contexts.
* [dbt](https://github.com/nikogura/dbt) **star:13** A framework for running self-updating signed binaries from a central, trusted repository.   ![There was an update last week][Green]
* [Death](https://github.com/vrecan/death) **star:132** Managing go application shutdown with signals.![star > 100][Bronze]
* [Deepcopier](https://github.com/ulule/deepcopier) **star:209** Simple struct copying for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [delve](https://github.com/derekparker/delve) **star:11976** Go debugger.![star > 5000][Gold]   ![There was an update last week][Green]
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** Compile-time controlled logger to make your release smaller without removing debug calls.   ![It hasn't been updated in the last year][Yellow]
* [ergo](https://github.com/cristianoliveira/ergo) **star:309** The management of multiple local services running over different ports made easy.![star > 100][Bronze]
* [evaluator](https://github.com/nullne/evaluator) **star:15** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   ![It hasn't been updated in the last year][Yellow]
* [fastlz](https://github.com/digitalcrab/fastlz) **star:12** Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang.   ![It hasn't been updated in the last year][Yellow]
* [filetype](https://github.com/h2non/filetype) **star:943** Small package to infer the file type checking the magic numbers signature.![star > 100][Bronze]
* [filler](https://github.com/yaronsumel/filler) **star:14** small utility to fill structs using "fill" tag.   ![It hasn't been updated in the last year][Yellow]
* [filter](https://github.com/gookit/filter) **star:12** provide filtering, sanitizing, and conversion of Go data.
* [fzf](https://github.com/junegunn/fzf) **star:23051** Command-line fuzzy finder written in Go.![star > 5000][Gold]   ![There was an update last week][Green]
* [gaper](https://github.com/maxcnunes/gaper) **star:39** Builds and restarts a Go project when it crashes or some watched file changes.
* [generate](https://github.com/go-playground/generate) **star:19** runs go generate recursively on a specified path or environment variable and can filter by regex.   ![It hasn't been updated in the last year][Yellow]
* [ghokin](https://github.com/antham/ghokin) **star:12** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:715** Simple, seamless, lightweight time tracking for Git.![star > 100][Bronze]
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:45** Parse TODOs in your GO code.
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:159** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:83** Pure Go bsdiff and bspatch libraries and CLI tools.
* [go-dry](https://github.com/ungerik/go-dry) **star:433** DRY (don't repeat yourself) package for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-funk](https://github.com/thoas/go-funk) **star:1168** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).![star > 1000][Silver]   ![There was an update last week][Green]
* [go-health](https://github.com/Talento90/go-health) **star:63** Health package simplifies the way you add health check to your services.   ![It hasn't been updated in the last year][Yellow]
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:14** Go library for encoding structs into Header fields.
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:2** Go package for working with Problem Details.
* [go-rate](https://github.com/beefsack/go-rate) **star:292** Timed rate limiter for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:104** XML Sitemap generator written in Go.![star > 100][Bronze]
* [go-torch](https://github.com/uber/go-torch) **star:3636** Stochastic flame graph profiler for Go programs.![star > 1000][Silver]
* [go-trigger](https://github.com/sadlil/go-trigger) **star:181** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [goback](https://github.com/carlescere/goback) **star:40** Go simple exponential backoff package.   ![It hasn't been updated in the last year][Yellow]
* [godaemon](https://github.com/VividCortex/godaemon) **star:403** Utility to write daemons.![star > 100][Bronze]
* [godropbox](https://github.com/dropbox/godropbox) **star:3741** Common libraries for writing Go services/applications from Dropbox.![star > 1000][Silver]
* [gohper](https://github.com/cosiner/gohper) **star:249** Various tools/modules help for development.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [golarm](https://github.com/msempere/golarm) **star:34** Fire alarms with system events.   ![It hasn't been updated in the last year][Yellow]
* [golog](https://github.com/mlimaloureiro/golog) **star:43** Easy and lightweight CLI tool to time track your tasks.
* [gopencils](https://github.com/bndr/gopencils) **star:423** Small and simple package to easily consume REST APIs.![star > 100][Bronze]
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:22** a small golang lib to generate placeholder images.   ![It hasn't been updated in the last year][Yellow]
* [goreadability](https://github.com/philipjkim/goreadability) **star:32** Webpage summary extractor using Facebook Open Graph and arc90's readability.
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:4482** Deliver Go binaries as fast and easily as possible.![star > 1000][Silver]   ![There was an update last week][Green]
* [goreporter](https://github.com/wgliang/goreporter) **star:2481** Golang tool that does static analysis, unit testing, code review and generate code quality report.![star > 1000][Silver]
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:27** SeaweedFS client library with almost full features.
* [gostrutils](https://github.com/ik5/gostrutils) **star:16** Collections of string manipulation and conversion functions.
* [gotenv](https://github.com/subosito/gotenv) **star:140** Load environment variables from `.env` or any `io.Reader` in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gpath](https://github.com/tenntenn/gpath) **star:30** Library to simplify access struct fields with Go's expression in reflection.   ![It hasn't been updated in the last year][Yellow]
* [gubrak](https://github.com/novalagung/gubrak) **star:138** Golang utility library with syntactic sugar. It's like lodash, but for golang.![star > 100][Bronze]
* [handy](https://github.com/miguelpragier/handy) **star:45** Many utilities and helpers like string handlers/formatters and validators.   ![There was an update last week][Green]
* [htcat](https://github.com/htcat/htcat) **star:481** Parallel and Pipelined HTTP GET Utility.![star > 100][Bronze]
* [hub](https://github.com/github/hub) **star:16898** wrap git commands with additional functionality to interact with github from the terminal.![star > 5000][Gold]   ![There was an update last week][Green]
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2003** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.![star > 1000][Silver]
* [immortal](https://github.com/immortal/immortal) **star:603** \*nix cross-platform (OS agnostic) supervisor.![star > 100][Bronze]
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:39** Use x86 SIMD without writing any assembly code.   ![It hasn't been updated in the last year][Yellow]
* [jump](https://github.com/gsamokovarov/jump) **star:655** Jump helps you navigate faster by learning your habits.![star > 100][Bronze]
* [koazee](https://github.com/wesovilabs/koazee) **star:294** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.![star > 100][Bronze]
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:100** LiveReload server for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mc](https://github.com/minio/mc) **star:1106** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.![star > 1000][Silver]   ![There was an update last week][Green]
* [mergo](https://github.com/imdario/mergo) **star:848** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.![star > 100][Bronze]
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:42** Pure Go ultra performant MIME sniffing library/utility.
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:7** A MIME type sniffer for Go.
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:106** Package for MIME type detection based on magic numbers.![star > 100][Bronze]   ![There was an update last week][Green]
* [minify](https://github.com/tdewolff/minify) **star:1855** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.![star > 1000][Silver]   ![There was an update last week][Green]
* [minquery](https://github.com/icza/minquery) **star:50** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).
* [mmake](https://github.com/tj/mmake) **star:1446** Modern Make.![star > 1000][Silver]
* [moldova](https://github.com/StabbyCutyou/moldova) **star:148** Utility for generating random data based on an input template.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [mole](https://github.com/davrodpin/mole) **star:1300** cli app to easily create ssh tunnels.![star > 1000][Silver]
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:58** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
* [multitick](https://github.com/VividCortex/multitick) **star:58** Multiplexor for aligned tickers.   ![It hasn't been updated in the last year][Yellow]
* [myhttp](https://github.com/inancgumus/myhttp) **star:34** Simple API to make HTTP GET requests with timeout support.   ![It hasn't been updated in the last year][Yellow]
* [netbug](https://github.com/e-dard/netbug) **star:65** Easy remote profiling of your services.   ![It hasn't been updated in the last year][Yellow]
* [okrun](https://github.com/xta/okrun) **star:14** go run error steamroller.   ![It hasn't been updated in the last year][Yellow]
* [olaf](https://github.com/btnguyen2k/olaf) **star:1** Twitter Snowflake implemented in Go.
* [onecache](https://github.com/adelowo/onecache) **star:99** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
* [panicparse](https://github.com/maruel/panicparse) **star:2139** Groups similar goroutines and colorizes stack dump.![star > 1000][Silver]
* [peco](https://github.com/peco/peco) **star:5445** Simplistic interactive filtering tool.![star > 5000][Gold]
* [pgo](https://github.com/arthurkushman/pgo) **star:24** Convenient functions for PHP community.
* [pm](https://github.com/VividCortex/pm) **star:71** Process (i.e. goroutine) manager with an HTTP API.
* [profile](https://github.com/pkg/profile) **star:1002** Simple profiling support package for Go.![star > 1000][Silver]
* [rclient](https://github.com/zpatrick/rclient) **star:27** Readable, flexible, simple-to-use client for REST APIs.
* [realize](https://github.com/tockins/realize) **star:3154** Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.![star > 1000][Silver]   ![There was an update last week][Green]
* [repeat](https://github.com/ssgreg/repeat) **star:56** Go implementation of different backoff strategies useful for retrying operations and heartbeating.
* [request](https://github.com/mozillazg/request) **star:356** Go HTTP Requests for Humans™.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [rerate](https://github.com/abo/rerate) **star:12** Redis-based rate counter and rate limiter for Go.   ![It hasn't been updated in the last year][Yellow]
* [rerun](https://github.com/ivpusic/rerun) **star:154** Recompiling and rerunning go apps when source changes.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [resty](https://github.com/go-resty/resty) **star:1950** Simple HTTP and REST client for Go inspired by Ruby rest-client.![star > 1000][Silver]
* [retry](https://github.com/kamilsk/retry) **star:154** The most advanced functional mechanism to perform actions repetitively until successful.![star > 100][Bronze]   ![There was an update last week][Green]
* [retry](https://github.com/percolate/retry) **star:2** A simple but highly configurable retry package for Go.
* [retry](https://github.com/thedevsaddam/retry) **star:34** Simple and easy retry mechanism package for Go.   ![It hasn't been updated in the last year][Yellow]
* [retry](https://github.com/shafreeck/retry) **star:10** A pretty simple library to ensure your work to be done.
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:28** Retrying made simple and easy for golang.
* [robustly](https://github.com/VividCortex/robustly) **star:134** Runs functions resiliently, catching and restarting panics.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [scan](https://github.com/blockloop/scan) **star:12** Scan golang `sql.Rows` directly to structs, slices, or primitive types.
* [serve](https://github.com/syntaqx/serve) **star:190** A static http server anywhere you need.![star > 100][Bronze]
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:3** Slice conversion between primitive types.   ![There was an update last week][Green]
* [slicer](https://github.com/leaanthony/slicer) **star:3** Makes working with slices easier.
* [spinner](https://github.com/briandowns/spinner) **star:779** Go package to easily provide a terminal spinner with options.![star > 100][Bronze]
* [sqlx](https://github.com/jmoiron/sqlx) **star:6753** provides a set of extensions on top of the excellent built-in database/sql package.![star > 5000][Gold]   ![There was an update last week][Green]
* [sslice](https://github.com/yaa110/sslice) **star:5** Create a slice which is always sorted.
* [Storm](https://github.com/asdine/storm) **star:1348** Simple and powerful toolkit for BoltDB.![star > 1000][Silver]
* [structs](https://github.com/PumpkinSeed/structs) **star:12** Implement simple functions to manipulate structs.   ![It hasn't been updated in the last year][Yellow]
* [Task](https://github.com/go-task/task) **star:1928** simple "Make" alternative.![star > 1000][Silver]   ![There was an update last week][Green]
* [toolbox](https://github.com/viant/toolbox) **star:94** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   ![There was an update last week][Green]
* [tracer](https://github.com/kamilsk/tracer) **star:8** Simple, lightweight tracing.
* [ugo](https://github.com/alxrm/ugo) **star:20** ugo is slice toolbox with concise syntax for Go.   ![It hasn't been updated in the last year][Yellow]
* [UNIS](https://github.com/esemplastic/unis) **star:69** Common Architecture™ for String Utilities in Go.   ![It hasn't been updated in the last year][Yellow]
* [usql](https://github.com/knq/usql) **star:4667** usql is a universal command-line interface for SQL databases.![star > 1000][Silver]   ![There was an update last week][Green]
* [util](https://github.com/shomali11/util) **star:133** Collection of useful utility functions. (strings, concurrency, manipulations, ...).![star > 100][Bronze]
* [wuzz](https://github.com/asciimoo/wuzz) **star:8253** Interactive cli tool for HTTP inspection.![star > 5000][Gold]
* [xferspdy](https://github.com/monmohan/xferspdy) **star:68** Xferspdy provides binary diff and patch library in golang.   ![It hasn't been updated in the last year][Yellow]

## UUID

*Libraries for working with UUIDs.*

* [goid](https://github.com/jakehl/goid) **star:20** Generate and Parse RFC4122 compliant V4 UUIDs.
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:2** A tiny and efficient Go unique string ID generator.
* [sno](https://github.com/muyo/sno) **star:15** Compact, sortable and fast unique IDs with embedded metadata.
* [ulid](https://github.com/oklog/ulid) **star:1671** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).![star > 1000][Silver]
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands. 
* [uuid](https://github.com/agext/uuid) **star:10** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
* [uuid](https://github.com/gofrs/uuid) **star:564** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.![star > 100][Bronze]
* [wuid](https://github.com/edwingeng/wuid) **star:285** An extremely fast unique number generator, 10-135 times faster than UUID.![star > 100][Bronze]

## Validation

*Libraries for validation.*

* [checkdigit](https://github.com/osamingo/checkdigit) **star:44** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
* [govalidator](https://github.com/asaskevich/govalidator) **star:3537** Validators and sanitizers for strings, numerics, slices and structs.![star > 1000][Silver]
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:713** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.![star > 100][Bronze]
* [jio](https://github.com/faceair/jio) **star:21** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   ![Contains Chinese documents][CN]
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1031** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.![star > 1000][Silver]   ![There was an update last week][Green]
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:6** A norms and conventions validator for Terraform.   ![There was an update last week][Green]
* [validate](https://github.com/gookit/validate) **star:92** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [validate](https://github.com/gobuffalo/validate) **star:19** This package provides a framework for writing validations for Go applications.
* [validator](https://github.com/go-playground/validator) **star:3485** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.![star > 1000][Silver]

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh) **star:70** Scriptable server and net/http middleware for GitHub Webhooks.
* [git2go](https://github.com/libgit2/git2go) **star:1356** Go bindings for libgit2.![star > 1000][Silver]
* [go-git](https://github.com/src-d/go-git) **star:4185** highly extensible Git implementation in pure Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:69** manipulate and inspect VCS repositories in Go.   ![There was an update last week][Green]
* [hercules](https://github.com/src-d/hercules) **star:523** gaining advanced insights from Git repository history.![star > 100][Bronze]
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   ![It hasn't been updated in the last year][Yellow]

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf) **star:524** Go bindings for FFmpeg av\* libraries.![star > 100][Bronze]
* [go-astisub](https://github.com/asticode/go-astisub) **star:165** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).![star > 100][Bronze]
* [go-astits](https://github.com/asticode/go-astits) **star:259** Parse and demux MPEG Transport Streams (.ts) natively in GO.![star > 100][Bronze]
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:37** Parser and generator library for Apple m3u8 playlists.
* [goav](https://github.com/giorgisio/goav) **star:776** Comphrensive Go bindings for FFmpeg.![star > 100][Bronze]
* [gst](https://github.com/ziutek/gst) **star:152** Go bindings for GStreamer.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:11** Subtitle format support for go. Supports .srt, .ttml, and .ass.
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:62** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).
* [v4l](https://github.com/korandiz/v4l) **star:28** Video capture library for Linux, written in Go.   ![It hasn't been updated in the last year][Yellow]

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [Aero](https://github.com/aerogo/aero) **star:157** High-performance web framework for Go, reaches top scores in Lighthouse.![star > 100][Bronze]
* [Air](https://github.com/aofei/air) **star:333** An ideally refined web framework for Go.![star > 100][Bronze]
* [Banjo](https://github.com/nsheremet/banjo) **star:7** Very simple and fast web framework for Go.   ![It hasn't been updated in the last year][Yellow]
* [Beego](https://github.com/astaxie/beego) **star:21403** beego is an open-source, high-performance web framework for the Go programming language.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) **star:14628** High performance, minimalist Go web framework.![star > 5000][Gold]   ![There was an update last week][Green]
* [Fireball](https://github.com/zpatrick/fireball) **star:49** More "natural" feeling web framework.
* [Gem](https://github.com/go-gem/gem) **star:153** Simple and fast web framework, friendly to REST API.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Gin](https://github.com/gin-gonic/gin) **star:29787** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.![star > 5000][Gold]   ![There was an update last week][Green]
* [Gizmo](https://github.com/NYTimes/gizmo) **star:2840** Microservice toolkit used by the New York Times.![star > 1000][Silver]   ![There was an update last week][Green]
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3331** Quick and easy way to setup a RESTful JSON API.![star > 1000][Silver]
* [go-rest](https://github.com/ungerik/go-rest) **star:116** Small and evil REST framework for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Goa](https://github.com/goadesign/goa) **star:3502** Goa provides a holistic approach for developing remote APIs and microservices in Go.![star > 1000][Silver]   ![There was an update last week][Green]
* [Golax](https://github.com/fulldump/golax) **star:71** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   ![It hasn't been updated in the last year][Yellow]
* [Golf](https://github.com/dinever/golf) **star:235** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Gondola](https://github.com/rainycape/gondola) **star:314** The web framework for writing faster sites, faster.![star > 100][Bronze]
* [gongular](https://github.com/mustafaakin/gongular) **star:415** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.![star > 100][Bronze]
* [hiboot](https://github.com/hidevopsio/hiboot) **star:86** hiboot is a high performance web application framework with auto configuration and dependency injection support.   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Macaron](https://github.com/go-macaron/macaron) **star:2814** Macaron is a high productive and modular design web framework in Go.![star > 1000][Silver]
* [mango](https://github.com/paulbellamy/mango) **star:339** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Microservice](https://github.com/claygod/microservice) **star:57** The framework for the creation of microservices, written in Golang.
* [neo](https://github.com/ivpusic/neo) **star:392** Neo is minimal and fast Go Web Framework with extremely simple API.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [nio](https://github.com/go-nio/nio) **star:21** Modern, minimal and productive Go HTTP framework.
* [patron](https://github.com/beatlabs/patron) **star:33** Patron is a microservice framework following best cloud practices with a focus on productivity.   ![There was an update last week][Green]
* [Resoursea](https://github.com/resoursea/api) **star:29** REST framework for quickly writing resource based services.   ![It hasn't been updated in the last year][Yellow]
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Revel](https://github.com/revel/revel) **star:11233** High-productivity web framework for the Go language.![star > 5000][Gold]
* [rex](https://github.com/goanywhere/rex) **star:27** Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.   ![It hasn't been updated in the last year][Yellow]
* [sawsij](https://github.com/jaybill/sawsij) **star:2** lightweight, open-source web framework for building high-performance, data-driven web applications.   ![It hasn't been updated in the last year][Yellow]
* [tango](https://github.com/lunny/tango) **star:817** Micro & pluggable web framework for Go.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:995** Go framework for building JSON web services inspired by Dropwizard.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [traffic](https://github.com/pilu/traffic) **star:518** Sinatra inspired regexp/pattern mux and web framework for Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [uAdmin](https://github.com/uadmin/uadmin) **star:52** Fully featured web framework for Golang, inspired by Django.
* [utron](https://github.com/gernest/utron) **star:2135** Lightweight MVC framework for Go(Golang).![star > 1000][Silver]
* [vox](https://github.com/aisk/vox) **star:39** A golang web framework for humans, inspired by Koa heavily.
* [WebGo](https://github.com/bnkamalesh/webgo) **star:73** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).
* [YARF](https://github.com/yarf-framework/yarf) **star:50** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.
* [Zerver](https://github.com/cosiner/zerver)  Zerver is an expressive, modular, feature completed RESTful framework.

### Middlewares

#### Actual middlewares

* [client-timing](https://github.com/posener/client-timing) **star:13** An HTTP client for Server-Timing header.
* [CORS](https://github.com/rs/cors) **star:1215** Easily add CORS capabilities to your API.![star > 1000][Silver]
* [formjson](https://github.com/rs/formjson) **star:33** Transparently handle JSON input as a standard form POST.   ![It hasn't been updated in the last year][Yellow]
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:744** Add/parse Server-Timing header.![star > 100][Bronze]
* [Limiter](https://github.com/ulule/limiter) **star:789** Dead simple rate limit middleware for Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:89** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).
* [Tollbooth](https://github.com/didip/tollbooth) **star:1242** Rate limit HTTP request handler.![star > 1000][Silver]
* [XFF](https://github.com/sebest/xff) **star:72** Handle `X-Forwarded-For` header and friends.

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice) **star:1822** Painless middleware chaining for Go.![star > 1000][Silver]
* [catena](https://github.com/codemodus/catena) **star:7** http.Handler wrapper catenation (same API as "chain").
* [chain](https://github.com/codemodus/chain) **star:63** Handler wrapper chaining with scoped data (net/context-based "middleware").
* [go-wrap](https://github.com/go-on/wrap) **star:59** Small middlewares package for net/http.
* [gores](https://github.com/alioygur/gores) **star:81** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
* [interpose](https://github.com/carbocation/interpose) **star:289** Minimalist net/http middleware for golang.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [muxchain](https://github.com/stephens2424/muxchain) **star:208** Lightweight middleware for net/http.![star > 100][Bronze]
* [negroni](https://github.com/urfave/negroni) **star:6321** Idiomatic HTTP middleware for Golang.![star > 5000][Gold]   ![Contains Chinese documents][CN]
* [render](https://github.com/unrolled/render) **star:1268** Go package for easily rendering JSON, XML, and HTML template responses.![star > 1000][Silver]
* [renderer](https://github.com/thedevsaddam/renderer) **star:169** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.![star > 100][Bronze]
* [rye](https://github.com/InVisionApp/rye) **star:94** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.
* [stats](https://github.com/thoas/stats) **star:536** Go middleware that stores various information about your web application.![star > 100][Bronze]

### Routers

* [alien](https://github.com/gernest/alien) **star:106** Lightweight and fast http router from outer space.![star > 100][Bronze]
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:39** A simple Go HTTP router.
* [Bone](https://github.com/go-zoo/bone) **star:1219** Lightning Fast HTTP Multiplexer.![star > 1000][Silver]
* [Bxog](https://github.com/claygod/Bxog) **star:94** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
* [chi](https://github.com/go-chi/chi) **star:6011** Small, fast and expressive HTTP router built on net/context.![star > 5000][Gold]   ![There was an update last week][Green]
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:745** High performance router forked from `httprouter`. The first router fit for `fasthttp`.![star > 100][Bronze]
* [FastRouter](https://github.com/razonyang/fastrouter) **star:19** a fast, flexible HTTP router written in Go.   ![It hasn't been updated in the last year][Yellow]
* [gocraft/web](https://github.com/gocraft/web) **star:1395** Mux and middleware package in Go.![star > 1000][Silver]
* [Goji](https://github.com/goji/goji) **star:768** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.![star > 100][Bronze]
* [GoRouter](https://github.com/vardius/gorouter) **star:49** GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
* [gowww/router](https://github.com/gowww/router) **star:157** Lightning fast HTTP router fully compatible with the net/http.Handler interface.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [httprouter](https://github.com/julienschmidt/httprouter) **star:9680** High performance router. Use this and the standard http handlers to form a very high performance web framework.![star > 5000][Gold]
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:385** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.![star > 100][Bronze]
* [lars](https://github.com/go-playground/lars) **star:375** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.![star > 100][Bronze]
* [mux](https://github.com/gorilla/mux) **star:9630** Powerful URL router and dispatcher for golang.![star > 5000][Gold]
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:359** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.![star > 100][Bronze]
* [pure](https://github.com/go-playground/pure) **star:85** Is a lightweight HTTP router that sticks to the std "net/http" implementation.
* [Siesta](https://github.com/VividCortex/siesta) **star:349** Composable framework to write middleware and handlers.![star > 100][Bronze]
* [vestigo](https://github.com/husobee/vestigo) **star:251** Performant, stand-alone, HTTP compliant URL Router for go web applications.![star > 100][Bronze]
* [violetear](https://github.com/nbari/violetear) **star:95** Go HTTP router.
* [xmux](https://github.com/rs/xmux) **star:88** High performance muxer based on `httprouter` with `net/context` support.   ![It hasn't been updated in the last year][Yellow]
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:451** A simple and fast HTTP router for Go.![star > 100][Bronze]

## Windows

* [d3d9](https://github.com/gonutz/d3d9) **star:85** Go bindings for Direct3D9.
* [go-ole](https://github.com/go-ole/go-ole) **star:551** Win32 OLE implementation for golang.![star > 100][Bronze]
* [gosddl](https://github.com/MonaxGT/gosddl) **star:1** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

## XML

*Libraries and tools for manipulating XML.*

* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** Simple command line XML comparer that generates diffs of folders, files and tags.   ![It hasn't been updated in the last year][Yellow]
* [xml2map](https://github.com/sbabiv/xml2map) **star:15** XML to MAP converter written Golang.
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:6** Procedural XML generation API based on libxml2's xmlwriter module.
* [xpath](https://github.com/antchfx/xpath) **star:164** XPath package for Go.![star > 100][Bronze]
* [xquery](https://github.com/antchfx/xquery) **star:145** XQuery lets you extract data from HTML/XML documents using XPath expression.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [zek](https://github.com/miku/zek) **star:254** Generate a Go struct from XML.![star > 100][Bronze]

# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:165** Checks recent changes to a Go project for backwards incompatible changes.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [dupl](https://github.com/mibk/dupl) **star:171** Tool for code clone detection.![star > 100][Bronze]
* [errcheck](https://github.com/kisielk/errcheck) **star:1319** Errcheck is a program for checking for unchecked errors in Go programs.![star > 1000][Silver]
* [gcvis](https://github.com/davecheney/gcvis) **star:918** Visualise Go program GC trace data in real time.![star > 100][Bronze]
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:95** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:281** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-critic](https://github.com/go-critic/go-critic) **star:573** source code linter that brings checks that are currently not implemented in other linters.![star > 100][Bronze]
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:182** An easy way to find outdated dependencies of your Go projects.![star > 100][Bronze]
* [go-outdated](https://github.com/firstrow/go-outdated) **star:45** Console application that displays outdated packages.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:372** Web based Golang AST visualizer.![star > 100][Bronze]
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [GoLint](https://github.com/golang/lint) **star:3144** Golint is a linter for Go source code.![star > 1000][Silver]
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [gostatus](https://github.com/shurcooL/gostatus) **star:241** Command line tool, shows the status of repositories that contain Go packages.![star > 100][Bronze]
* [lint](https://github.com/surullabs/lint) **star:63** Run linters as part of go test.
* [php-parser](https://github.com/z7zmey/php-parser) **star:633** A Parser for PHP written in Go.![star > 100][Bronze]
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) **star:14** tarp finds functions and methods without direct unit tests in Go source code.   ![It hasn't been updated in the last year][Yellow]
* [unconvert](https://github.com/mdempsky/unconvert) **star:258** Remove unnecessary type conversions from Go source.![star > 100][Bronze]
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   ![It hasn't been updated in the last year][Yellow]

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:29** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [go-mode](https://github.com/dominikh/go-mode.el) **star:950** Go mode for GNU/Emacs.![star > 100][Bronze]   ![There was an update last week][Green]
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1480** Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.![star > 1000][Silver]   ![There was an update last week][Green]
* [gocode](https://github.com/nsf/gocode) **star:4736** Autocompletion daemon for the Go programming language.![star > 1000][Silver]
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3232** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.![star > 1000][Silver]
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:17** Vim plugin for generating Go tests based on the function's or method's signature.
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:12** Go language support for the Theia IDE.
* [velour](https://github.com/velour/velour) **star:16** IRC client for the acme editor.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:81** Vim plugin to highlight syntax errors on save.   ![It hasn't been updated in the last year][Yellow]
* [vim-go](https://github.com/fatih/vim-go) **star:10802** Go development plugin for Vim.![star > 5000][Gold]   ![There was an update last week][Green]
* [vscode-go](https://github.com/Microsoft/vscode-go) **star:5101** Extension for Visual Studio Code (VS Code) which provides support for the Go language.![star > 5000][Gold]   ![There was an update last week][Green]
* [Watch](https://github.com/eaburns/Watch) **star:167** Runs a command in an acme win on file changes.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Go Generate Tools

* [generic](https://github.com/usk81/generic) **star:28** flexible data type for Go.
* [genny](https://github.com/cheekybits/genny) **star:960** Elegant generics for Go.![star > 100][Bronze]
* [gocontracts](https://github.com/Parquery/gocontracts) **star:51** brings design-by-contract to Go by synchronizing the code with the documentation.
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests) **star:2182** Generate Go tests from your source code.![star > 1000][Silver]
* [gounit](https://github.com/hexdigest/gounit) **star:29** Generate Go tests using your own templates.
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:12** Generate Haskell inspired functions for your slices.
* [re2dfa](https://github.com/opennota/re2dfa) **star:168** Transform regular expressions into finite state machines and output Go source code.![star > 100][Bronze]
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.

## Go Tools

* [colorgo](https://github.com/songgao/colorgo) **star:97** Wrapper around `go` command for colorized `go build` output.   ![It hasn't been updated in the last year][Yellow]
* [depth](https://github.com/KyleBanks/depth) **star:375** Visualize dependency trees of any package by analyzing imports.![star > 100][Bronze]
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:13** A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:1986** Visualize call graph of your Go program using dot format.![star > 1000][Silver]
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:37** Bash completion for go and wgo.   ![It hasn't been updated in the last year][Yellow]
* [go-swagger](https://github.com/go-swagger/go-swagger) **star:3961** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.![star > 1000][Silver]   ![There was an update last week][Green]
* [godbg](https://github.com/tylerwince/godbg) **star:156** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.![star > 100][Bronze]
* [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:3764** Navigate through go files efficiently with the OctoLinker browser extension for GitHub.![star > 1000][Silver]   ![There was an update last week][Green]
* [richgo](https://github.com/kyoh86/richgo) **star:387** Enrich `go test` outputs with text decorations.![star > 100][Bronze]
* [rts](https://github.com/galeone/rts) **star:184** RTS: response to struct. Generates Go structs from server responses.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [aurora](https://github.com/xuri/aurora) **star:400** Cross-platform web-based Beanstalkd queue server console.![star > 100][Bronze]
* [awsenv](https://github.com/soniah/awsenv) **star:21** Small binary that loads Amazon (AWS) environment variables for a profile.   ![It hasn't been updated in the last year][Yellow]
* [Blast](https://github.com/dave/blast) **star:167** A simple tool for API load testing and batch jobs.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [bombardier](https://github.com/codesenberg/bombardier) **star:1737** Fast cross-platform HTTP benchmarking tool.![star > 1000][Silver]
* [bosun](https://github.com/bosun-monitor/bosun) **star:2851** Time Series Alerting Framework.![star > 1000][Silver]   ![There was an update last week][Green]
* [DepCharge](https://github.com/centerorbit/depcharge) **star:9** Helps orchestrating the execution of commands across the many dependencies in larger projects.
* [dogo](https://github.com/liudng/dogo) **star:216** Monitoring changes in the source file and automatically compile and run (restart).![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:24** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
* [drone-scp](https://github.com/appleboy/drone-scp) **star:55** Copy files and artifacts via SSH using a binary, docker or Drone CI.
* [Dropship](https://github.com/chrismckenzie/dropship) **star:46** Tool for deploying code via cdn.   ![It hasn't been updated in the last year][Yellow]
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:99** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
* [fac](https://github.com/mkchoi212/fac) **star:1614** Command-line user interface to fix git merge conflicts.![star > 1000][Silver]
* [gaia](https://github.com/gaia-pipeline/gaia) **star:3737** Build powerful pipelines in any programming language.![star > 1000][Silver]   ![There was an update last week][Green]
* [Gitea](https://github.com/go-gitea/gitea) **star:15147** Fork of Gogs, entirely community driven.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:64** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:667** Enable your Go applications to self update.![star > 100][Bronze]
* [gobrew](https://github.com/cryptojuice/gobrew) **star:175** gobrew lets you easily switch between multiple versions of go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [godbg](https://github.com/sirnewton01/godbg) **star:219** Web-based gdb front-end application.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative) **star:312** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [govvv](https://github.com/ahmetalpbalkan/govvv) **star:386** “go build” wrapper to easily add version information into Go binaries.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [gox](https://github.com/mitchellh/gox) **star:3352** Dead simple, no frills Go cross compile tool.![star > 1000][Silver]
* [goxc](https://github.com/laher/goxc) **star:1624** build tool for Go, with a focus on cross-compiling and packaging.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [grapes](https://github.com/yaronsumel/grapes) **star:135** Lightweight tool designed to distribute commands over ssh with ease.![star > 100][Bronze]
* [GVM](https://github.com/moovweb/gvm) **star:4465** GVM provides an interface to manage Go versions.![star > 1000][Silver]
* [Hey](https://github.com/rakyll/hey) **star:6266** Hey is a tiny program that sends some load to a web application.![star > 5000][Gold]
* [kala](https://github.com/ajvb/kala) **star:1357** Simplistic, modern, and performant job scheduler.![star > 1000][Silver]
* [kcli](https://github.com/cswank/kcli) **star:78** Command line tool for inspecting kafka topics/partitions/messages.
* [kubernetes](https://github.com/kubernetes/kubernetes) **star:56020** Container Cluster Manager from Google.![star > 5000][Gold]   ![There was an update last week][Green]
* [lstags](https://github.com/ivanilves/lstags) **star:219** Tool and API to sync Docker images across different registries.![star > 100][Bronze]
* [lwc](https://github.com/timdp/lwc) **star:8** A live-updating version of the UNIX wc command.   ![It hasn't been updated in the last year][Yellow]
* [manssh](https://github.com/xwjdsh/manssh) **star:204** manssh is a command line tool for managing your ssh alias config easily.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Moby](https://github.com/moby/moby) **star:54356** Collaborative project for the container ecosystem to assemble container-based systems.![star > 5000][Gold]   ![There was an update last week][Green]
* [Mora](https://github.com/emicklei/mora) **star:266** REST server for accessing MongoDB documents and meta data.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [ostent](https://github.com/ostrost/ostent) **star:164** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Packer](https://github.com/mitchellh/packer) **star:9223** Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.![star > 5000][Gold]   ![There was an update last week][Green]
* [Pewpew](https://github.com/bengadbois/pewpew) **star:202** Flexible HTTP command line stress tester.![star > 100][Bronze]   ![There was an update last week][Green]
* [Pomerium](https://github.com/pomerium/pomerium) **star:494** Pomerium is an identity-aware access proxy.![star > 100][Bronze]   ![There was an update last week][Green]
* [Rodent](https://github.com/alouche/rodent) **star:30** Rodent helps you manage Go versions, projects and track dependencies.   ![It hasn't been updated in the last year][Yellow]
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:994** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.![star > 100][Bronze]
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:537** Manage BareMetal Servers from Command Line (as easily as with Docker).![star > 100][Bronze]
* [script](https://github.com/bitfield/script) **star:914** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.![star > 100][Bronze]
* [sg](https://github.com/ChristopherRabotin/sg) **star:5** Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.   ![It hasn't been updated in the last year][Yellow]
* [skm](https://github.com/TimothyYe/skm) **star:549** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!![star > 100][Bronze]   ![There was an update last week][Green]
* [StatusOK](https://github.com/sanathp/statusok) **star:1160** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.![star > 1000][Silver]
* [traefik](https://github.com/containous/traefik) **star:23640** Reverse proxy and load balancer with support for multiple backends.![star > 5000][Gold]   ![There was an update last week][Green]
* [Vegeta](https://github.com/tsenart/vegeta) **star:12115** HTTP load testing tool and library. It's over 9000!![star > 5000][Gold]
* [webhook](https://github.com/adnanh/webhook) **star:4073** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.![star > 1000][Silver]
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:67** Cli tool to remotely execute commands on Windows machines.

### Other Software

* [borg](https://github.com/crufter/borg) **star:1417** Terminal based search engine for bash snippets.![star > 1000][Silver]   ![It hasn't been updated in the last year][Yellow]
* [boxed](https://github.com/tejo/boxed) **star:72** Dropbox based blog engine.
* [Cherry](https://github.com/rafael-santiago/cherry) **star:193** Tiny webchat server in Go.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [Circuit](https://github.com/gocircuit/circuit) **star:1786** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.![star > 1000][Silver]
* [Comcast](https://github.com/tylertreat/Comcast) **star:6173** Simulate bad network connections.![star > 5000][Gold]
* [confd](https://github.com/kelseyhightower/confd) **star:6406** Manage local application configuration files using templates and data from etcd or consul.![star > 5000][Gold]
* [DDNS](https://github.com/skibish/ddns) **star:97** Personal DDNS client with Digital Ocean Networking DNS as backend.
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community) **star:824** Modern wiki software that integrates data from SaaS tools.![star > 100][Bronze]
* [drive](https://github.com/odeke-em/drive) **star:4954** Google Drive client for the commandline.![star > 1000][Silver]
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:2686** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.![star > 1000][Silver]
* [gfile](https://github.com/Antonito/gfile) **star:497** Securely transfer files between two computers, without any third party, over WebRTC.![star > 100][Bronze]
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:877** App that displays updates for the Go packages in your GOPATH.![star > 100][Bronze]
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:374** Video streaming torrent client.![star > 100][Bronze]
* [GoBoy](https://github.com/Humpheh/goboy) **star:2101** Nintendo Game Boy Color emulator written in Go.![star > 1000][Silver]
* [gocc](https://github.com/goccmack/gocc) **star:343** Gocc is a compiler kit for Go written in Go.![star > 100][Bronze]
* [GoDNS](https://github.com/timothyye/godns) **star:427** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.![star > 100][Bronze]   ![There was an update last week][Green]
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:12** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.   ![It hasn't been updated in the last year][Yellow]
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor) **star:11330** Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.![star > 5000][Gold]
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide) **star:250** Browser accessible IDE. Designed for Go with Go.![star > 100][Bronze]
* [ipe](https://github.com/dimiro1/ipe) **star:275** Open source Pusher server implementation compatible with Pusher client libraries written in GO.![star > 100][Bronze]
* [joincap](https://github.com/assafmo/joincap) **star:122** Command-line utility for merging multiple pcap files together.![star > 100][Bronze]
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Leaps](https://github.com/jeffail/leaps) **star:640** Pair programming service using Operational Transforms.![star > 100][Bronze]
* [lgo](https://github.com/yunabe/lgo) **star:1798** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.![star > 1000][Silver]
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide) **star:5480** LiteIDE is a simple, open source, cross-platform Go IDE.![star > 5000][Gold]   ![Contains Chinese documents][CN]
* [mockingjay](https://github.com/quii/mockingjay-server) **star:414** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.![star > 100][Bronze]   ![There was an update last week][Green]
* [myLG](https://github.com/mehrdadrad/mylg) **star:2196** Command Line Network Diagnostic tool written in Go.![star > 1000][Silver]
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** Simple NaCL EC25519 based crypto pipe tool written in Go.
* [nes](https://github.com/fogleman/nes) **star:4126** Nintendo Entertainment System (NES) emulator written in Go.![star > 1000][Silver]
* [orange-cat](https://github.com/noraesae/orange-cat) **star:177** Markdown previewer written in Go.![star > 100][Bronze]
* [Orbit](https://github.com/gulien/orbit) **star:128** A simple tool for running commands and generating files from templates.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [peg](https://github.com/pointlander/peg) **star:607** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.![star > 100][Bronze]
* [Pipe](https://github.com/b3log/pipe) **star:2904** A small and beautiful blogging platform.![star > 1000][Silver]
* [restic](https://github.com/restic/restic) **star:7416** De-duplicating backup program.![star > 5000][Gold]   ![There was an update last week][Green]
* [rkt](https://github.com/coreos/rkt) **star:8738** App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM.![star > 5000][Gold]   ![There was an update last week][Green]
* [scc](https://github.com/boyter/scc) **star:768** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.![star > 100][Bronze]   ![There was an update last week][Green]
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:8217** Fast, Simple and Scalable Distributed File System with O(1) disk seek.![star > 5000][Gold]   ![There was an update last week][Green]
* [shell2http](https://github.com/msoap/shell2http) **star:406** Executing shell commands via http server (for prototyping or remote control).![star > 100][Bronze]   ![There was an update last week][Green]
* [snap](https://github.com/intelsdi-x/snap) **star:1802** Powerful telemetry framework.![star > 1000][Silver]
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   ![It hasn't been updated in the last year][Yellow]
* [Stack Up](https://github.com/pressly/sup) **star:1990** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.![star > 1000][Silver]
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:3929** Proxy to simulate network and system conditions for automated tests.![star > 1000][Silver]
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow) **star:598** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.![star > 100][Bronze]
* [wellington](https://github.com/wellington/wellington) **star:289** Sass project management tool, extends the language with sprite functions (like Compass).![star > 100][Bronze]

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench) **star:89** Framework to compare the performance between different Go versions.   ![It hasn't been updated in the last year][Yellow]
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:19** Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.   ![It hasn't been updated in the last year][Yellow]
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:121** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1259** Go HTTP request router benchmark and comparison.![star > 1000][Silver]
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:997** Go web framework benchmark.![star > 100][Bronze]   ![There was an update last week][Green]
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:865** Benchmarks of Go serialization methods.![star > 100][Bronze]   ![There was an update last week][Green]
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) **star:52** Benchmarks of common basic operations for the Go language.   ![It hasn't been updated in the last year][Yellow]
* [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) **star:17** Tiny collection of Go micro benchmarks. The intent is to compare some language features to others.
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:48** Collection of benchmarks for popular Go database/SQL utilities.   ![It hasn't been updated in the last year][Yellow]
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:93** Go micro-benchmarks for calculating the speed of language constructs.
* [kvbench](https://github.com/jimrobinson/kvbench) **star:14** Key/Value database benchmark.   ![It hasn't been updated in the last year][Yellow]
* [skynet](https://github.com/atemerev/skynet) **star:913** Skynet 1M threads microbenchmark.![star > 100][Bronze]
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:172** Compare various Image resize algorithms for the Go language.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]

## Conferences

* [Capital Go](http://www.capitalgolang.com)  Washington, D.C., USA.
* [dotGo](http://www.dotgo.eu)  Paris, France.
* [GoCon](http://gocon.connpass.com/)  Tokyo, Japan.
* [GoDays](https://www.godays.io/)  Berlin, Germany.
* [GoLab](http://golab.io/)  Florence, Italy.
* [GolangUK](http://golanguk.com/)  London, UK.
* [GopherChina](http://gopherchina.org)  Shanghai, China.
* [GopherCon](http://www.gophercon.com/)  Denver, USA.
* [GopherCon Australia](https://gophercon.com.au/)  Sydney, Australia.
* [GopherCon Brazil](https://gopherconbr.org)  Florianópolis, BR.
* [GopherCon Europe](https://gophercon.is/)  Reykjavik, Iceland.
* [GopherCon India](https://www.gophercon.in/)  Pune, India.
* [GopherCon Israel](https://www.gophercon.org.il/)  Tel Aviv, Israel.
* [GopherCon Russia](https://www.gophercon-russia.ru)  Moscow, Russia.
* [GopherCon Singapore](https://gophercon.sg)  Mapletree Business City, Singapore.
* [GopherCon Vietnam](https://gophercon.vn/)  Ho Chi Minh City, Vietnam.
* [GothamGo](http://gothamgo.com/)  New York City, USA.
* [GoWayFest](https://goway.io/)  Minsk, Belarus.

## E-Books

* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Go 101](https://go101.org)  A book focusing on Go syntax/semantics and all kinds of details.
* [Go Bootcamp](http://golangbootcamp.com)
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:10** in Persian.
* [GoBooks](https://github.com/dariubs/GoBooks) **star:6770** A curated list of Go books.![star > 5000][Gold]
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1522** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.![star > 1000][Silver]
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:31** Go gopher Vector Data [.ai, .svg].   ![It hasn't been updated in the last year][Yellow]
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:64** adorable gopher logos.   ![It hasn't been updated in the last year][Yellow]
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:311** Gopherize yourself.![star > 100][Bronze]
* [gophers](https://github.com/ashleymcnamara/gophers) **star:1850** Gopher artworks by Ashley McNamara.![star > 1000][Silver]
* [gophers](https://github.com/egonelbre/gophers) **star:1579** Free gophers.![star > 1000][Silver]
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

* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:14535** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.![star > 5000][Gold]
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:24609** List of other amazingly awesome lists.![star > 5000][Gold]
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:36** Collection of Go projects that needs help. Good place to start your open-source way in Go.   ![It hasn't been updated in the last year][Yellow]
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:141** Collection of Go images, graphics, and art.![star > 100][Bronze]   ![It hasn't been updated in the last year][Yellow]
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:31452** Golang ebook intro how to build a web app with golang.![star > 5000][Gold]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:452** A little e-book on Ethereum Development with Go.![star > 100][Bronze]   ![Contains Chinese documents][CN]
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:4029** Go's reference card.![star > 1000][Silver]
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:689** Examples of Golang compared to Node.js for learning.![star > 100][Bronze]
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:4440** Learn Go with test-driven development.![star > 1000][Silver]   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1131** Intro to go for experienced programmers.![star > 1000][Silver]
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

