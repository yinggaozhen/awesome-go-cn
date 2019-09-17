# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4/docs/awesome.svg "awesome!"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2019-09-17 11:53:44(Synchronize every day)**

[![english](https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/chinese.svg)](README.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) financial support to Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python)。

**Icon** :


Icon | State  
:-:|-
![awesome][Awesome] | star > 2000
![There was an update last week][Green] | There was an update last week。You can basically determine that the current library is in an actively maintained state。
![Not updated in last year][Yellow] | Not updated in last year。Reflects that the maintenance of the library is not high enthusiasm, use should be careful。
![Archived][Archived] | The project has been archived。
![Contains Chinese documents][CN] | This project contains Chinese documents。
![godoc document][GoDoc] | godoc document links。

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

* [Oto](https://github.com/hajimehoshi/oto) **star:453** A low-level library to play sound on multiple platforms.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:308** Go bindings for the PortAudio audio I/O library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:260** Music theory models in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:250** Go package capable of generating waveform images from audio streams.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:210** Go bindings for PortMidi.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:114** Fast and stable ID3 parsing and writing library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [flac](https://github.com/mewkiz/flac) **star:103** Native Go FLAC encoder/decoder with support for FLAC streams.   [![godoc][GoDoc]](https://godoc.org/github.com/mewkiz/flac)
* [mix](https://github.com/go-mix/mix) **star:99** Sequence-based Go-native audio mixer for music apps.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:98** Native Go MP3 decoder.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:93** libsox bindings for go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [flac](https://github.com/eaburns/flac) **star:85** No-frills native Go FLAC decoder that decodes FLAC files into byte slices.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/eaburns/flac)
* [malgo](https://github.com/gen2brain/malgo) **star:76** Mini audio library.
* [taglib](https://github.com/wtolson/go-taglib) **star:69** Go bindings for taglib.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:57** Native Go AAC bitstream parser.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:28** Lightweight MP3 decoder library.
* [vorbis](https://github.com/mccoyst/vorbis) **star:22** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:6222** Golang implementation of JSON Web Tokens (JWT).   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:5204** Authorization library that supports access control models like ACL, RBAC, ABAC.   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [oauth2](https://github.com/golang/oauth2) **star:2443** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2327** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:1970** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [osin](https://github.com/openshift/osin) **star:1551** Golang OAuth2 server library.   [![godoc][GoDoc]](https://godoc.org/github.com/openshift/osin)
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1314** Standalone, specification-compliant,  OAuth2 server written in Golang.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [go-jose](https://github.com/square/go-jose) **star:1163** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/square/go-jose)
* [gologin](https://github.com/dghubble/gologin) **star:1070** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:925** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [loginsrv](https://github.com/tarent/loginsrv) **star:828** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [scs](https://github.com/alexedwards/scs) **star:590** Session Manager for HTTP servers.   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [permissions2](https://github.com/xyproto/permissions2) **star:358** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [paseto](https://github.com/o1egl/paseto) **star:247** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [httpauth](https://github.com/goji/httpauth) **star:182** HTTP Authentication middleware.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:165** Simple, flexible, secure and idiomatic web session management with pluggable backends.   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:157** JWT middleware for Golang http servers with many configuration options.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/pascaldekloe/jwt) **star:100** Lightweight JSON Web Token (JWT) library.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [session](https://github.com/icza/session) **star:91** Go session management for web servers (including support for Google App Engine - GAE).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [branca](https://github.com/hako/branca) **star:78** Golang implementation of Branca Tokens.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hako/branca)
* [jwt](https://github.com/robbert229/jwt) **star:72** Clean and easy to use implementation of JSON Web Tokens (JWT).   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sessions](https://github.com/adam-hanna/sessions) **star:47** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [securecookie](https://github.com/chmike/securecookie) **star:34** Efficient secure cookie encoding/decoding.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [sjwt](https://github.com/brianvoe/sjwt) **star:34** Simple jwt generator and parser.   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [rbac](https://github.com/zpatrick/rbac) **star:29** Minimalistic RBAC package for Go applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [signedvalue](https://github.com/sashka/signedvalue) **star:7** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1681** Simple and clean Telegram bot client.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Syfaro/telegram-bot-api)
* [telebot](https://github.com/tucnak/telebot) **star:969** Telegram bot framework written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:485** IRC, Slack & Telegram bot written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [slacker](https://github.com/shomali11/slacker) **star:325** Easy to use framework to create Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:230** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [tbot](https://github.com/yanzay/tbot) **star:227** Telegram bot server with API similar to net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [Kelp](https://github.com/stellar/kelp) **star:183** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:168** Service oriented IRC bot using Redis and JSON for messaging.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:142** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.   [![godoc][GoDoc]](https://godoc.org/github.com/oklahomer/go-sarah)
* [hanu](https://github.com/sbstjn/hanu) **star:109** Framework for writing Slack bots.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:87** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:58** Framework for building Telegram bots.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:27** Simple Go [VK](https://vk.com) bot library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:14** Another framework for building Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:13852** Commander for modern Go CLI interactions.   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:11851** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![star > 2000][Awesome]](https://github.com/urfave/cli)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2628** Command line and flag parser supporting sub commands.   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1530** go command line option parser.   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [readline](https://github.com/chzyer/readline) **star:1388** Pure golang implementation that provides most features in GNU-Readline under MIT license.   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [docopt.go](https://github.com/docopt/docopt.go) **star:1181** Command-line arguments parser that will make you smile.   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1011** Go library for implementing command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [cli-init](https://github.com/tcnksm/gcli) **star:873** The easy way to start building Golang command line applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [pflag](https://github.com/spf13/pflag) **star:790** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [go-arg](https://github.com/alexflint/go-arg) **star:769** Struct-based argument parsing in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:629** Write bash completions in Go + Go command bash completion.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [mow.cli](https://github.com/jawher/mow.cli) **star:624** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [liner](https://github.com/peterh/liner) **star:594** Go readline-like library for command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [cli](https://github.com/mkideal/cli) **star:486** Feature-rich and easy to use command-line package based on golang struct tags.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [flaggy](https://github.com/integrii/flaggy) **star:458** A robust and idiomatic flags package with excellent subcommand support.   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [ops](https://github.com/nanovms/ops) **star:274** Unikernel Builder/Orchestrator.   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:122** Command line argument parser inspired by Python's argparse module.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [flag](https://github.com/cosiner/flag) **star:101** Simple but powerful command line option parsing library for Go supporting subcommand.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:98** Small command line interface framework.   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [sflags](https://github.com/octago/sflags) **star:95** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [commandeer](https://github.com/jaffee/commandeer) **star:93** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [wmenu](https://github.com/dixonwille/wmenu) **star:89** Easy to use menu structure for cli applications that prompts users to make choices.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [cli](https://github.com/teris-io/cli) **star:57** Simple and complete API for building command line interfaces in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [job](https://github.com/liujianping/job) **star:55** JOB, make your short-term command as a long-term job.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [env](https://github.com/codingconcepts/env) **star:42** Tag-based environment configuration for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [gocmd](https://github.com/devfacet/gocmd) **star:34** Go library for building command line applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [flagvar](https://github.com/sgreben/flagvar) **star:31** A collection of flag argument types for Go's standard `flag` package.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [strumt](https://github.com/antham/strumt) **star:31** Library to create prompt chain.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [cmdr](https://github.com/hedzr/cmdr) **star:18** A POSIX/GNU style, getopt-like command-line UI Go library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [argv](https://github.com/cosiner/argv) **star:17** Go library to split command line string as arguments array using the bash syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:10** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [sand](https://github.com/Zaba505/sand) **star:7** Simple API for creating interpreters and so much more.   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [ts](https://github.com/liujianping/ts) **star:7** Timestamp convert & compare tool.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/ts)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:9093** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![star > 2000][Awesome]](https://github.com/gizak/termui)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gocui](https://github.com/jroimartin/gocui) **star:5520** Minimalist Go library aimed at creating Console User Interfaces.   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [termbox-go](https://github.com/nsf/termbox-go) **star:3525** Termbox is a library for creating cross-platform text-based interfaces.   [![star > 2000][Awesome]](https://github.com/nsf/termbox-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/termbox-go)
* [color](https://github.com/fatih/color) **star:3066** Versatile package for colored terminal output.   [![star > 2000][Awesome]](https://github.com/fatih/color)   [![godoc][GoDoc]](https://godoc.org/github.com/fatih/color)   ![Archived][Archived]
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2524** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1553** Flexible library to render progress bars in terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1175** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:855** Library for updating terminal output in realtime.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [termtables](https://github.com/apcera/termtables)  Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.
* [termdash](https://github.com/mum4k/termdash) **star:827** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [mpb](https://github.com/vbauerster/mpb) **star:736** Multi progress bar for terminal applications.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [progressbar](https://github.com/schollz/progressbar) **star:670** Basic thread-safe progress bar that works in every OS.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/progressbar)
* [aurora](https://github.com/logrusorgru/aurora) **star:670** ANSI terminal colors that supports fmt.Printf/Sprintf.   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [uitable](https://github.com/gosuri/uitable) **star:506** Library to improve readability in terminal apps using tabular data.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:388** Colorable writer for windows.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:354** isatty for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [chalk](https://github.com/ttacon/chalk) **star:313** Intuitive package for prettifying terminal/console output.   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [tabby](https://github.com/cheynewallace/tabby) **star:250** A tiny library for super simple Golang tables.   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [gookit/color](https://github.com/gookit/color) **star:225** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   ![Contains Chinese documents][CN]
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:197** Go library for color output in terminals.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [simpletable](https://github.com/alexeyco/simpletable) **star:170** Simple tables in terminal with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [cfmt](https://github.com/mingrammer/cfmt) **star:67** Contextual fmt inspired by bootstrap color classes.   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:30** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [ctc](https://github.com/wzshiming/ctc) **star:10** The non-invasive cross-platform terminal color library does not need to modify the Print method.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   ![Contains Chinese documents][CN]

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:9687** Go configuration with fangs.   [![star > 2000][Awesome]](https://github.com/spf13/viper)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2477** Go library for managing configuration data from environment variables.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv) **star:2227** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [ini](https://github.com/go-ini/ini) **star:1650** Go package to read and write INI files.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:905** Parse environment variables to Go structs (with defaults).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [konfig](https://github.com/lalamove/konfig) **star:516** Composable, observable and performant config handling for Go for the distributed processing era.   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [confita](https://github.com/heetch/confita) **star:249** Load configuration in cascade from multiple backends into a struct.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:243** Lightweight configuration manager for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/olebedev/config) **star:215** JSON or YAML configuration wrapper with environment variables and flags parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/config)
* [joshbetz/config](https://github.com/joshbetz/config) **star:195** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.   [![godoc][GoDoc]](https://godoc.org/github.com/joshbetz/config)
* [config](https://github.com/JeremyLoy/config) **star:193** Cloud native application configuration. Bind ENV to structs in only two lines.   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [hjson](https://github.com/hjson/hjson-go) **star:175** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [envconfig](https://github.com/vrischmann/envconfig) **star:147** Read your configuration from environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/vrischmann/envconfig)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:117** read INI-style configuration files into Go structs; supports user-defined types and subsections.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [goConfig](https://github.com/crgimenes/goConfig) **star:109** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [koanf](https://github.com/knadh/koanf) **star:96** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [gookit/config](https://github.com/gookit/config) **star:93** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   ![Contains Chinese documents][CN]
* [envh](https://github.com/antham/envh) **star:92** Helpers to manage environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [envcfg](https://github.com/tomazk/envcfg) **star:89** Un-marshaling environment variables to Go structs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tomazk/envcfg)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gofigure](https://github.com/ian-kent/gofigure) **star:57** Go application configuration made easy.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [configure](https://github.com/paked/configure) **star:49** Provides configuration through multiple sources, including JSON, flags and environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [harvester](https://github.com/beatlabs/harvester) **star:43** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/harvester)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:39** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [go-up](https://github.com/ufoscout/go-up) **star:25** A simple configuration library with recursive placeholders resolution and no magic.   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:20** Golang package for parsing ini-style configuration files.   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:8** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.   [![godoc][GoDoc]](https://godoc.org/github.com/the4thamigo-uk/conflate)
* [sprbox](https://github.com/oblq/sprbox) **star:4** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/sprbox)

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:19402** Drone is a Continuous Integration platform built on Docker, written in Go.   [![star > 2000][Awesome]](https://github.com/drone/drone)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [goveralls](https://github.com/mattn/goveralls) **star:587** Go integration for Coveralls.io continuous code coverage tracking system.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/goveralls)
* [overalls](https://github.com/go-playground/overalls) **star:98** Multi-Package go project coverprofile for tools like goveralls.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/overalls)
* [duci](https://github.com/duck8823/duci) **star:46** A simple ci server no needs domain specific languages.   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:33** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:423** Pure Go CSS Preprocessor.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass) **star:133** Go wrapper to the 100% Sass compatible libsass project.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:6686** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5215** Collection of useful, performant, and thread-safe data structures.   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1228** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1177** Probabilistic data structures for processing continuous, unbounded streams.   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [gota](https://github.com/kniren/gota) **star:912** Implementation of dataframes, series, and data wrangling methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [roaring](https://github.com/RoaringBitmap/roaring) **star:697** Go package implementing compressed bitsets.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/RoaringBitmap/roaring)
* [willf/bloom](https://github.com/willf/bloom) **star:679** Go package implementing Bloom filters.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bloom)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:665** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:529** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [bitset](https://github.com/willf/bitset) **star:494** Go package implementing bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [trie](https://github.com/derekparker/trie) **star:431** Trie implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/trie)
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:314** In-memory geo index.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [mafsa](https://github.com/smartystreets/mafsa) **star:274** MA-FSA implementation with Minimal Perfect Hashing.   [![godoc][GoDoc]](https://godoc.org/github.com/smartystreets/mafsa)   ![Archived][Archived]
* [algorithms](https://github.com/shady831213/algorithms) **star:261** Algorithms and data structures.CLRS study.   [![godoc][GoDoc]](https://godoc.org/github.com/shady831213/algorithms)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:195** Skip list implementation in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [hilbert](https://github.com/google/hilbert) **star:185** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [merkletree](https://github.com/cbergoon/merkletree) **star:154** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:128** Binary packer and unpacker helps user build custom binary stream.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [bloom](https://github.com/zhenjl/bloom) **star:128** Bloom filters implemented in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:108** Go implementation of Adaptive Radix Tree.   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:102** In-memory LRU string-interface{} map with expiration for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/diegobernardes/ttlcache)
* [skiplist](https://github.com/MauriceGit/skiplist) **star:102** Very fast Go Skiplist implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:101** Region quadtrees with efficient point location and neighbour finding.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [encoding](https://github.com/zhenjl/encoding) **star:96** Integer Compression Libraries for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [ring](https://github.com/TheTannerRyan/ring) **star:91** Go implementation of a high performance, thread safe bloom filter.   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:82** A small, powerful and flexible merge library.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [deque](https://github.com/gammazero/deque) **star:77** Fast ring-buffer deque (double-ended queue).   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [skiplist](https://github.com/gansidui/skiplist) **star:66** Skiplist implementation in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [bit](https://github.com/yourbasic/bit) **star:60** Golang set data structure with bonus bit-twiddling functions.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:57** Implementation to calculate levenshtein distance in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:35** Fast in-memory key:value store/cache library. Pointer caches.   [![godoc][GoDoc]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:35** Concurrent FIFO queue.   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [levenshtein](https://github.com/agext/levenshtein) **star:34** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [crunch](https://github.com/superwhiskers/crunch) **star:19** Go package implementing buffers for handling various datatypes easily.   [![godoc][GoDoc]](https://godoc.org/github.com/superwhiskers/crunch)
* [goset](https://github.com/zoumo/goset) **star:16** A useful Set collection implementation for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [typ](https://github.com/gurukami/typ) **star:11** Null Types, Safe primitive type conversion and fetching value from complex structures.   [![godoc][GoDoc]](https://godoc.org/github.com/gurukami/typ)
* [dict](https://github.com/srfrog/dict) **star:9** Python-like dictionaries (dict) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [hide](https://github.com/emvi/hide) **star:8** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [deque](https://github.com/edwingeng/deque) **star:8** A highly optimized double-ended queue.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/deque)
* [set](https://github.com/StudioSol/set) **star:7** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [null](https://github.com/emvi/null) **star:6** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)
* [ptrie](https://github.com/viant/ptrie) **star:1** An implementation of prefix tree.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/ptrie)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:1** Map with expiring key-value pairs.   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:26351** Monitoring system and time series database.   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [tidb](https://github.com/pingcap/tidb) **star:20570** TiDB is a distributed SQL database. Inspired by the design of Google F1.   [![star > 2000][Awesome]](https://github.com/pingcap/tidb)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/tidb)   ![Contains Chinese documents][CN]
* [influxdb](https://github.com/influxdb/influxdb) **star:17368** Scalable datastore for metrics, events, and real-time analytics.   [![star > 2000][Awesome]](https://github.com/influxdb/influxdb)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/influxdb/influxdb)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:16986** Scalable, Geo-Replicated, Transactional Datastore.   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [dgraph](https://github.com/dgraph-io/dgraph) **star:10894** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [bolt](https://github.com/boltdb/bolt) **star:10096** Low-level key/value database for Go.   [![star > 2000][Awesome]](https://github.com/boltdb/bolt)   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/boltdb/bolt)   ![Archived][Archived]
* [groupcache](https://github.com/golang/groupcache) **star:7770** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:6475** Fast key-value store in Go.   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:4821** The lightweight, distributed, relational database built on SQLite.   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3252** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3101** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   [![star > 2000][Awesome]](https://github.com/siddontang/ledisdb)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/ledisdb)
* [go-cache](https://github.com/pmylund/go-cache) **star:3041** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [BigCache](https://github.com/allegro/bigcache) **star:2532** Efficient key/value cache for gigabytes of data.   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [buntdb](https://github.com/tidwall/buntdb) **star:2470** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2384** Your NoSQL database powered by Golang.   [![star > 2000][Awesome]](https://github.com/HouzuoGuo/tiedot)   [![godoc][GoDoc]](https://godoc.org/github.com/HouzuoGuo/tiedot)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1146** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [cache2go](https://github.com/muesli/cache2go) **star:1085** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [GCache](https://github.com/bluele/gcache) **star:940** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:906** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   ![Contains Chinese documents][CN]
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:894** CovenantSQL is a SQL database on blockchain.   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:772** Home-grown disk-backed key-value store.   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [moss](https://github.com/couchbase/moss) **star:730** Moss is a simple LSM key-value storage engine written in 100% Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/moss)
* [eliasdb](https://github.com/krotik/eliasdb) **star:539** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:526** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
* [levigo](https://github.com/jmhodges/levigo) **star:366** Levigo is a Go wrapper for LevelDB.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [pudge](https://github.com/recoilme/pudge) **star:226** Fast and simple  key/value store written using Go's standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [piladb](https://github.com/fern4lvarez/piladb) **star:170** Lightweight RESTful database engine based on stack data structures.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Vasto](https://github.com/chrislusf/vasto) **star:156** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:87** Key-value store with persistence.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:69** Tiny flat file JSON store.   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:41** RESTful caching micro-service backed by Couchbase server.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [bcache](https://github.com/iwanbk/bcache) **star:32** Eventually consistent distributed in-memory  cache Go library.   [![godoc][GoDoc]](https://godoc.org/github.com/iwanbk/bcache)
* [cache](https://github.com/akyoto/cache) **star:22** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:2940** Database migrations. CLI and Golang library.   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1435** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:344** Database schema migration helper for Gorm ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [goose](https://github.com/steinbacher/goose) **star:119** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/steinbacher/goose)
* [darwin](https://github.com/GuiaBolso/darwin) **star:83** Database schema evolution library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/GuiaBolso/darwin)
* [migrator](https://github.com/lopezator/migrator) **star:76** Dead simple Go database migration library.   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:27** A Go package to help write migrations with go-pg/pg.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [gondolier](https://github.com/emvi/gondolier) **star:26** Database migration library using struct decorators.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/gondolier)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:20** Django style fixtures for Golang's excellent built-in database/sql library.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-fixtures)
* [avro](https://github.com/khezen/avro) **star:7** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:8600** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [pgweb](https://github.com/sosedoff/pgweb) **star:6044** Web-based PostgreSQL database browser.   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [kingshard](https://github.com/flike/kingshard) **star:4716** kingshard is a high performance proxy for MySQL powered by Golang.   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   ![Contains Chinese documents][CN]
* [orchestrator](https://github.com/github/orchestrator) **star:3111** MySQL replication topology manager & visualizer.   [![star > 2000][Awesome]](https://github.com/github/orchestrator)   [![godoc][GoDoc]](https://godoc.org/github.com/github/orchestrator)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2498** Sync your MySQL data into Elasticsearch automatically.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [pREST](https://github.com/nuveo/prest) **star:2102** Serve a RESTful API from any PostgreSQL database.   [![star > 2000][Awesome]](https://github.com/nuveo/prest)   [![godoc][GoDoc]](https://godoc.org/github.com/nuveo/prest)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:1959** Go toolset to handle MySQL protocol and replication.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [chproxy](https://github.com/Vertamedia/chproxy) **star:311** HTTP proxy for ClickHouse database.   [![godoc][GoDoc]](https://godoc.org/github.com/Vertamedia/chproxy)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:143** Collects small insterts and sends big requests to ClickHouse servers.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [myreplication](https://github.com/2tvenom/myreplication) **star:142** MySql binary log replication listener. Supports statement and row based replication.   [![godoc][GoDoc]](https://godoc.org/github.com/2tvenom/myreplication)
* [octillery](https://github.com/knocknote/octillery) **star:60** Go package for sharding databases ( Supports every ORM or raw SQL ).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [dbbench](https://github.com/sj14/dbbench) **star:30** Database benchmarking tool with support for several databases and scripts.   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [datagen](https://github.com/codingconcepts/datagen) **star:10** A fast data generator that's multi-table aware and supports multi-row DML.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/datagen)

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2411** Go library that helps you build SQL queries.   [![star > 2000][Awesome]](https://github.com/Masterminds/squirrel)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/squirrel)
* [xo](https://github.com/knq/xo) **star:2210** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   [![star > 2000][Awesome]](https://github.com/knq/xo)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/knq/xo)
* [gendry](https://github.com/didi/gendry) **star:829** Non-invasive SQL builder and powerful data binder.   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   ![Contains Chinese documents][CN]
* [goqu](https://github.com/doug-martin/goqu) **star:669** Idiomatic SQL builder and query library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/doug-martin/goqu)
* [Dotsql](https://github.com/gchaincl/dotsql) **star:452** Go library that helps you keep sql files in one place and use them with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/dotsql)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:440** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
* [sqrl](https://github.com/elgris/sqrl) **star:184** SQL query builder, fork of Squirrel with improved performance.   [![godoc][GoDoc]](https://godoc.org/github.com/elgris/sqrl)
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [scaneo](https://github.com/variadico/scaneo) **star:149** Generate Go code to convert database rows into arbitrary structs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/variadico/scaneo)   ![Archived][Archived]
* [igor](https://github.com/galeone/igor) **star:78** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [godbal](https://github.com/xujiajun/godbal) **star:48** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [dbq](https://github.com/rocketlaunchr/dbq) **star:46** Zero boilerplate database operations for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:8314** MySQL driver for Go.   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [pq](https://github.com/lib/pq) **star:5269** Pure Go Postgres driver for database/sql.   [![star > 2000][Awesome]](https://github.com/lib/pq)   [![godoc][GoDoc]](https://godoc.org/github.com/lib/pq)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3500** SQLite3 driver for go that uses database/sql.   [![star > 2000][Awesome]](https://github.com/mattn/go-sqlite3)
    * [pgx](https://github.com/jackc/pgx) **star:2020** PostgreSQL driver supporting features beyond those exposed by database/sql.   [![star > 2000][Awesome]](https://github.com/jackc/pgx)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jackc/pgx)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1046** Microsoft MSSQL driver for Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8) **star:410** Oracle driver for go that uses database/sql.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-oci8)
    * [goracle](https://github.com/go-goracle/goracle) **star:251** Oracle driver for Go, using the ODPI-C driver.   ![There was an update last week][Green]
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:105** Firebird RDBMS SQL driver for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/nakagami/firebirdsql)
    * [gofreetds](https://github.com/minus5/gofreetds) **star:91** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).   [![godoc][GoDoc]](https://godoc.org/github.com/minus5/gofreetds)
    * [go-adodb](https://github.com/mattn/go-adodb) **star:90** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:39** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [bgc](https://github.com/viant/bgc) **star:12** Datastore Connectivity for BigQuery for go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:6852** Redis client for Golang.   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo) **star:6456** Redigo is a Go client for the Redis database.   [![star > 2000][Awesome]](https://github.com/gomodule/redigo)   [![godoc][GoDoc]](https://godoc.org/github.com/gomodule/redigo)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:3327** Official MongoDB driver for the Go language.   [![star > 2000][Awesome]](https://github.com/mongodb/mongo-go-driver)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mongodb/mongo-go-driver)
    * [mgo](https://github.com/globalsign/mgo) **star:1672** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   [![godoc][GoDoc]](https://godoc.org/github.com/globalsign/mgo)
    * [gorethink](https://github.com/dancannon/gorethink) **star:1469** Go language driver for RethinkDB.   [![godoc][GoDoc]](https://godoc.org/github.com/dancannon/gorethink)
    * [neoism](https://github.com/jmcvetta/neoism) **star:358** Neo4j client for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [redeo](https://github.com/bsm/redeo) **star:351** Redis-protocol compatible TCP servers/services.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redeo)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:309** Aerospike client in Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [gocb](https://github.com/couchbase/gocb) **star:294** Official Couchbase Go SDK.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/gocb)
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:294** Couchbase client in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:94** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:72** Neo4j REST Client in golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/davemeehan/Neo4j-GO)
    * [arangolite](https://github.com/solher/arangolite) **star:66** Lightweight golang driver for ArangoDB.   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [dynago](https://github.com/underarmour/dynago) **star:66** Dynago is a principle of least surprise client for DynamoDB.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/underarmour/dynago)
    * [godis](https://github.com/piaohao/godis) **star:63** redis client implement by golang, inspired by jedis.   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [go-couchdb](https://github.com/fjl/go-couchdb) **star:51** Yet another CouchDB HTTP API wrapper for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/fjl/go-couchdb)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:32** Go client library for Pilosa.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [xredis](https://github.com/shomali11/xredis) **star:9** Typesafe, customizable, clean & easy to use Redis client.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [godscache](https://github.com/defcronyke/godscache) **star:6** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.   [![godoc][GoDoc]](https://godoc.org/github.com/defcronyke/godscache)
    * [asc](https://github.com/viant/asc) **star:4** Datastore Connectivity for Aerospike for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/asc)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:5935** Modern text indexing library for go.   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [riot](https://github.com/go-ego/riot) **star:4783** Go Open Source, Distributed, Simple and efficient Search Engine.   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   ![Contains Chinese documents][CN]
    * [elastic](https://github.com/olivere/elastic) **star:4299** Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/olivere/elastic)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/olivere/elastic)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:1679** Official Elasticsearch client for Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/elastic/go-elasticsearch)
    * [elastigo](https://github.com/mattbaird/elastigo) **star:952** Elasticsearch client library.   [![godoc][GoDoc]](https://godoc.org/github.com/mattbaird/elastigo)
    * [elasticsql](https://github.com/cch123/elasticsql) **star:426** Convert sql to elasticsearch dsl in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cch123/elasticsql)
    * [skizze](https://github.com/seiflotfy/skizze) **star:68** probabilistic data-structures service and storage.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:12810** Graph database with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/google/cayley)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/google/cayley)
    * [cachego](https://github.com/fabiorphp/cachego) **star:110** Golang Cache component for multiple drivers.   [![godoc][GoDoc]](https://godoc.org/github.com/fabiorphp/cachego)
    * [gokv](https://github.com/philippgille/gokv) **star:99** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [dsc](https://github.com/viant/dsc) **star:14** Datastore connectivity for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2220** Now is a time toolkit for golang.   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [dateparse](https://github.com/araddon/dateparse) **star:927** Parse date's without knowing format in advance.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/araddon/dateparse)
* [carbon](https://github.com/uniplaces/carbon) **star:354** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   [![godoc][GoDoc]](https://godoc.org/github.com/uniplaces/carbon)
* [durafmt](https://github.com/hako/durafmt) **star:245** Time duration formatting library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [timeutil](https://github.com/leekchan/timeutil) **star:170** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/timeutil)
* [iso8601](https://github.com/relvacode/iso8601) **star:68** Efficiently parse ISO8601 date-times without regex.   [![godoc][GoDoc]](https://godoc.org/github.com/relvacode/iso8601)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:65** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/go-persian-calendar)
* [timespan](https://github.com/SaidinWoT/timespan) **star:61** For interacting with intervals of time, defined as a start time and a duration.   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [date](https://github.com/rickb777/date) **star:31** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.   [![godoc][GoDoc]](https://godoc.org/github.com/rickb777/date)
* [feiertage](https://github.com/wlbr/feiertage) **star:23** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [kair](https://github.com/GuilhermeCaruso/kair) **star:11** Date and Time - Golang Formatting Library.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:14773** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [grpc-go](https://github.com/grpc/grpc-go) **star:9436** The Go language implementation of gRPC. HTTP/2 based RPC.   [![star > 2000][Awesome]](https://github.com/grpc/grpc-go)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/grpc/grpc-go)
* [micro](https://github.com/micro/micro) **star:6731** Pluggable microservice toolkit and distributed systems platform.   [![star > 2000][Awesome]](https://github.com/micro/micro)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [NATS](https://github.com/nats-io/gnatsd) **star:6493** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.   [![star > 2000][Awesome]](https://github.com/nats-io/gnatsd)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/gnatsd)
* [rpcx](https://github.com/smallnest/rpcx) **star:3962** Distributed pluggable RPC service framework like alibaba Dubbo.   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [tendermint](https://github.com/tendermint/tendermint) **star:3213** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [torrent](https://github.com/anacrolix/torrent) **star:3043** BitTorrent client package.   [![star > 2000][Awesome]](https://github.com/anacrolix/torrent)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/torrent)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [raft](https://github.com/hashicorp/raft) **star:2923** Golang implementation of the Raft consensus protocol, by HashiCorp.   [![star > 2000][Awesome]](https://github.com/hashicorp/raft)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/raft)
* [dragonboat](https://github.com/lni/dragonboat) **star:2619** A feature complete and high performance multi-group Raft library in Go.   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   ![Contains Chinese documents][CN]
* [glow](https://github.com/chrislusf/glow) **star:2550** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [gleam](https://github.com/chrislusf/gleam) **star:2146** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2004** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [KrakenD](https://github.com/devopsfaith/krakend) **star:1868** Ultra performant API Gateway framework with middlewares.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/devopsfaith/krakend)
* [hprose](https://github.com/hprose/hprose-golang) **star:1022** Very newbility RPC Library, support 25+ languages now.   [![godoc][GoDoc]](https://godoc.org/github.com/hprose/hprose-golang)   ![Contains Chinese documents][CN]
* [ringpop-go](https://github.com/uber/ringpop-go) **star:577** Scalable, fault-tolerant application-layer sharding for Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/uber/ringpop-go)
* [gorpc](https://github.com/valyala/gorpc) **star:559** Simple, fast and scalable RPC library for high load.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [go-health](https://github.com/InVisionApp/go-health) **star:493** Library for enabling asynchronous dependency health checks in your service.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [digota](https://github.com/digota/digota) **star:305** grpc ecommerce microservice.   [![godoc][GoDoc]](https://godoc.org/github.com/digota/digota)
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [sleuth](https://github.com/ursiform/sleuth) **star:303** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [go-jump](https://github.com/dgryski/go-jump) **star:255** Port of Google's "Jump" Consistent Hash function.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [consistent](https://github.com/buraksezer/consistent) **star:208** Consistent hashing with bounded loads.   [![godoc][GoDoc]](https://godoc.org/github.com/buraksezer/consistent)
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [redis-lock](https://github.com/bsm/redis-lock) **star:146** Simplified distributed locking implementation using Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redis-lock)   ![Archived][Archived]
* [dht](https://github.com/anacrolix/dht) **star:131** BitTorrent Kademlia DHT implementation.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:114** The jsonrpc package helps implement of JSON-RPC 2.0.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:102** JSON-RPC 2.0 HTTP client implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:54** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:41** A revamped Google's jump consistent hash.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [drmaa](https://github.com/dgruber/drmaa) **star:25** Job submission library for cluster schedulers based on the DRMAA standard.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dgruber/drmaa)
* [flowgraph](https://github.com/vectaport/flowgraph) **star:18** flow-based programming package.   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:8** A library for using DynamoDB as an atomic counter.   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [outboxer](https://github.com/italolelis/outboxer) **star:5** Outboxer is a go library that implements the outbox pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:5303** Email and SMTP testing with web and API interface.   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [hermes](https://github.com/matcornic/hermes) **star:1643** Golang package that generates clean, responsive HTML e-mails.   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1110** A robust and flexible email library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:752** IMAP library for clients and servers.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:529** SendGrid's Go library for sending email.   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [Hectane](https://github.com/hectane/hectane) **star:169** Lightweight SMTP client providing an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [douceur](https://github.com/aymerick/douceur) **star:161** CSS inliner for your HTML emails.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [go-message](https://github.com/emersion/go-message) **star:120** Streaming library for the Internet Message Format and mail messages.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-message)
* [smtp](https://github.com/mailhog/smtp) **star:51** SMTP server protocol state machine.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/smtp)
* [go-dkim](https://github.com/toorop/go-dkim) **star:47** DKIM library, to sign & verify email.   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:36** Inline styling for HTML mail in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [Gomail](https://github.com/go-gomail/gomail/)  Gomail is a very simple and powerful package to send emails.

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) **star:4810** JavaScript interpreter written in Go.   [![star > 2000][Awesome]](https://github.com/robertkrimen/otto)   [![godoc][GoDoc]](https://godoc.org/github.com/robertkrimen/otto)
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:3008** Lua 5.1 VM and compiler written in Go.   [![star > 2000][Awesome]](https://github.com/yuin/gopher-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/yuin/gopher-lua)
* [go-lua](https://github.com/Shopify/go-lua) **star:1697** Port of the Lua 5.2 VM to pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [tengo](https://github.com/d5/tengo) **star:1344** Bytecode compiled script language for Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [anko](https://github.com/mattn/anko) **star:936** Scriptable interpreter written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-python](https://github.com/sbinet/go-python) **star:929** naive go bindings to the CPython C-API.   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [expr](https://github.com/antonmedv/expr) **star:764** an engine that can evaluate expressions.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [go-php](https://github.com/deuill/go-php) **star:693** PHP bindings for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:658** Duktape JavaScript engine bindings for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [golua](https://github.com/aarzilli/golua) **star:447** Go bindings for Lua C API.
* [gisp](https://github.com/jcla1/gisp) **star:429** Simple LISP in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [agora](https://github.com/PuerkitoBio/agora) **star:324** Dynamically typed, embeddable programming language in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/agora)   ![Archived][Archived]
* [gval](https://github.com/PaesslerAG/gval) **star:144** A highly customizable expression language written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [binder](https://github.com/alexeyco/binder) **star:32** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [gentee](https://github.com/gentee/gentee) **star:32** Embeddable scripting programming language.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) **star:5047** Package that provides simple error handling primitives.   [![star > 2000][Awesome]](https://github.com/pkg/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/errors)
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:750** Go (golang) package for representing a list of errors as a single error.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [errorx](https://github.com/joomcode/errorx) **star:591** A feature rich error package with stack traces, composition of errors and more.   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:506** Golang errors with stack trace and source fragments.   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [errlog](https://github.com/snwfdhmp/errlog) **star:194** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   [![godoc][GoDoc]](https://godoc.org/github.com/snwfdhmp/errlog)

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2303** FileSystem Abstraction System for Go.   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/hhrutter/pdfcpu) **star:1097** PDF processor.   [![godoc][GoDoc]](https://godoc.org/github.com/hhrutter/pdfcpu)
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [notify](https://github.com/rjeczalik/notify) **star:502** File system event notification library with simple API, similar to os/signal.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/notify)
* [opc](https://github.com/qmuntal/opc) **star:60** Load Open Packaging Conventions (OPC) files for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:51** Load csv file using tag.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-csv-tag)
* [vfs](https://github.com/C2FO/vfs) **star:25** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:15** Load gtfs files in go.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [checksum](https://github.com/codingsince1985/checksum) **star:11** Compute message digest, like MD5 and SHA256, for large files.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:11** Copy files for humans.   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
* [flop](https://github.com/homedepot/flop) **star:10** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:6** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:1667** Arbitrary-precision fixed-point decimal numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-money](https://github.com/rhymond/go-money) **star:638** Implementation of Fowler's Money pattern.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [go-finance](https://github.com/FlashBoys/go-finance) **star:537** Comprehensive financial markets data in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/FlashBoys/go-finance)
* [accounting](https://github.com/leekchan/accounting) **star:498** money and currency formatting for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [techan](https://github.com/sdcoffey/techan) **star:171** Technical analysis library with advanced market analysis and trading strategies.   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [orderbook](https://github.com/i25959341/orderbook) **star:81** Matching Engine for Limit Order Book in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/i25959341/orderbook)
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:64** Query OFX servers and/or parse the responses (with example command-line client).   [![godoc][GoDoc]](https://godoc.org/github.com/aclindsa/ofxgo)
* [vat](https://github.com/dannyvankooten/vat) **star:60** VAT number validation & EU VAT rates.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/vat)
* [transaction](https://github.com/claygod/transaction) **star:56** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [go-finance](https://github.com/alpeb/go-finance) **star:44** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/alpeb/go-finance)
* [currency](https://github.com/bnkamalesh/currency) **star:10** High performant & accurate currency computation package.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) **star:983** CSRF protection middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/nosurf)
* [binding](https://github.com/mholt/binding) **star:753** Binds form and JSON data from net/http Request to struct.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:452** CSRF protection for Go web applications & services.   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [form](https://github.com/go-playground/form) **star:357** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [conform](https://github.com/leebenson/conform) **star:175** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/leebenson/conform)
* [formam](https://github.com/monoculum/formam) **star:130** decode form's values into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:106** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1076** Useful collection of helpfully functional Go collection utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/tobyhede/go-underscore)
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:115** Monad, Functional Programming features for Golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:50** Functional Experiment in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)

## Game Development

*Awesome game development libraries.*

* [Leaf](https://github.com/name5566/leaf) **star:3142** Lightweight game server framework.   [![star > 2000][Awesome]](https://github.com/name5566/leaf)   [![godoc][GoDoc]](https://godoc.org/github.com/name5566/leaf)   ![Contains Chinese documents][CN]
* [Pixel](https://github.com/faiface/pixel) **star:2504** Hand-crafted 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/faiface/pixel)   [![godoc][GoDoc]](https://godoc.org/github.com/faiface/pixel)
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:1947** dead simple 2D game library in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/ebiten)
* [goworld](https://github.com/xiaonanln/goworld) **star:1245** Scalable game server engine, featuring space-entity framework and hot-swapping.   [![godoc][GoDoc]](https://godoc.org/github.com/xiaonanln/goworld)   ![Contains Chinese documents][CN]
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1177** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [engo](https://github.com/EngoEngine/engo) **star:1104** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [gonet](https://github.com/xtaci/gonet) **star:1059** Game server skeleton implemented with golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [nano](https://github.com/lonng/nano) **star:1042** Lightweight, facility, high performance golang based game server framework.   [![godoc][GoDoc]](https://godoc.org/github.com/lonng/nano)   ![Contains Chinese documents][CN]
* [termloop](https://github.com/JoelOtter/termloop) **star:1033** Terminal-based game engine for Go, built on top of Termbox.   [![godoc][GoDoc]](https://godoc.org/github.com/JoelOtter/termloop)
* [g3n](https://github.com/g3n/engine) **star:792** Go 3D Game Engine.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/g3n/engine)
* [Oak](https://github.com/oakmound/oak) **star:656** Pure Go game engine.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/oakmound/oak)
* [Azul3D](https://github.com/azul3d/engine) **star:428** 3D game engine written in Go.   ![It hasn't been updated in the last year][Yellow]
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:391** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [go-astar](https://github.com/beefsack/go-astar) **star:333** Go implementation of the A\* path finding algorithm.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-astar)
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:315** 2d game engine written in Go working on OpenGL.   [![godoc][GoDoc]](https://godoc.org/github.com/vova616/GarageEngine)
* [Pitaya](https://github.com/topfreegames/pitaya) **star:314** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/pitaya)
* [go3d](https://github.com/ungerik/go3d) **star:166** Performance oriented 2D/3D math package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go3d)
* [glop](https://github.com/runningwild/glop) **star:78** Glop (Game Library Of Power) is a fairly simple cross-platform game library.   ![It hasn't been updated in the last year][Yellow]

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1824** .NET LINQ-like query methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
* [jennifer](https://github.com/dave/jennifer) **star:1304** Generate arbitrary Go code without templates.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [gen](https://github.com/clipperhouse/gen) **star:1046** Code generation tool for ‘generics’-like functionality.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [goderive](https://github.com/awalterschulze/goderive) **star:758** Derives functions from input types.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/goderive)
* [GoWrap](https://github.com/hexdigest/gowrap) **star:279** Generate decorators for Go interfaces using simple templates.   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gowrap)
* [interfaces](https://github.com/rjeczalik/interfaces) **star:190** Command line tool for generating interface definitions.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [go-enum](https://github.com/abice/go-enum) **star:87** Code generation for enums from code comments.   [![godoc][GoDoc]](https://godoc.org/github.com/abice/go-enum)
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:87** Go preprocessor for package scoped reflection.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/pkgreflect)
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:44** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/t0pep0/efaceconv)
* [gotype](https://github.com/wzshiming/gotype) **star:23** Golang source code parsing, usage like reflect package.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   ![Contains Chinese documents][CN]
* [generis](https://github.com/senselogic/GENERIS) **star:19** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) **star:6396** Geolocation DB with spatial index and realtime geofencing.   [![star > 2000][Awesome]](https://github.com/tidwall/tile38)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/tile38)
* [S2 geometry](https://github.com/golang/geo) **star:907** S2 geometry library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [geocache](https://github.com/melihmucuk/geocache) **star:111** In-memory cache that is suitable for geolocation based applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [osm](https://github.com/paulmach/osm) **star:76** Library for reading, writing and working with OpenStreetMap data and APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/osm)
* [geoserver](https://github.com/hishamkaram/geoserver) **star:25** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/geoserver)
* [gismanager](https://github.com/hishamkaram/gismanager) **star:20** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)
* [pbf](https://github.com/maguro/pbf) **star:17** OpenStreetMap PBF golang encoder/decoder.   [![godoc][GoDoc]](https://godoc.org/github.com/maguro/pbf)

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:8684** Compiler from Go to JavaScript.   [![star > 2000][Awesome]](https://github.com/gopherjs/gopherjs)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gopherjs/gopherjs)
* [llgo](https://github.com/go-llvm/llgo) **star:992** LLVM-based compiler for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-llvm/llgo)
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:392** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tardisgo/tardisgo)
* [c4go](https://github.com/Konstantin8105/c4go) **star:166** Transpile C code to Go code.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:11** Transpile FORTRAN 77 code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)

## Goroutines

*Tools for managing and working with Goroutines.*

* [goworker](https://github.com/benmanns/goworker) **star:2272** goworker is a Go-based background worker.   [![star > 2000][Awesome]](https://github.com/benmanns/goworker)   [![godoc][GoDoc]](https://godoc.org/github.com/benmanns/goworker)
* [ants](https://github.com/panjf2000/ants) **star:2040** A high-performance goroutine pool for golang.   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   ![Contains Chinese documents][CN]
* [tunny](https://github.com/Jeffail/tunny) **star:1389** Goroutine pool for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/tunny)
* [grpool](https://github.com/ivpusic/grpool) **star:515** Lightweight Goroutine pool.   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [pool](https://github.com/go-playground/pool) **star:489** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [go-floc](https://github.com/workanator/go-floc) **star:169** Orchestrate goroutines with ease.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-floc)
* [workerpool](https://github.com/gammazero/workerpool) **star:151** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:110** Control goroutines execution order.   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [GoSlaves](https://github.com/themester/GoSlaves) **star:86** Simple and Asynchronous Goroutine pool library.   [![godoc][GoDoc]](https://godoc.org/github.com/themester/GoSlaves)
* [semaphore](https://github.com/kamilsk/semaphore) **star:76** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/semaphore)
* [semaphore](https://github.com/marusama/semaphore) **star:74** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/semaphore)
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:57** manages a resizeable pool of context-aware goroutines to bound concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
* [worker-pool](https://github.com/vardius/worker-pool) **star:47** goworker is a Go simple async worker pool.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/worker-pool)
* [breaker](https://github.com/kamilsk/breaker) **star:38** Flexible mechanism to make execution flow interruptible.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/breaker)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:36** CyclicBarrier for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/cyclicbarrier)
* [gollback](https://github.com/vardius/gollback) **star:27** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [threadpool](https://github.com/shettyh/threadpool) **star:21** Golang threadpool implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [Hunch](https://github.com/AaronJan/Hunch) **star:16** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   [![godoc][GoDoc]](https://godoc.org/github.com/AaronJan/Hunch)
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [artifex](https://github.com/borderstech/artifex) **star:14** Simple in-memory job queue for Golang using worker-based dispatching.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/artifex)
* [stl](https://github.com/ssgreg/stl) **star:9** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/stl)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.   [![godoc][GoDoc]](https://godoc.org/github.com/nikhilsaraf/go-tools)
* [routine](https://github.com/x-mod/routine) **star:4** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![godoc][GoDoc]](https://godoc.org/github.com/x-mod/routine)
* [queue](https://github.com/AnikHasibul/queue) **star:2** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.   [![godoc][GoDoc]](https://godoc.org/github.com/AnikHasibul/queue)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [ui](https://github.com/andlabs/ui) **star:7078** Platform-native GUI library for Go. Cross platform.   [![star > 2000][Awesome]](https://github.com/andlabs/ui)   [![godoc][GoDoc]](https://godoc.org/github.com/andlabs/ui)
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [fyne](https://github.com/fyne-io/fyne) **star:6537** Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows.   [![star > 2000][Awesome]](https://github.com/fyne-io/fyne)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/fyne-io/fyne)
* [qt](https://github.com/therecipe/qt) **star:6239** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   [![star > 2000][Awesome]](https://github.com/therecipe/qt)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/therecipe/qt)
* [webview](https://github.com/zserge/webview) **star:4782** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).   [![star > 2000][Awesome]](https://github.com/zserge/webview)
* [walk](https://github.com/lxn/walk) **star:3779** Windows application library kit for Go.   [![star > 2000][Awesome]](https://github.com/lxn/walk)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [app](https://github.com/murlokswarm/app) **star:3023** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   [![star > 2000][Awesome]](https://github.com/murlokswarm/app)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/murlokswarm/app)
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:2736** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1480** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/sciter-sdk/go-sciter)
* [gotk3](https://github.com/gotk3/gotk3) **star:800** Go bindings for GTK3.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd) **star:215** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/dtylman/gowd)

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4515** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   [![star > 2000][Awesome]](https://github.com/go-vgo/robotgo)   ![There was an update last week][Green]
* [systray](https://github.com/getlantern/systray) **star:813** Cross platform Go library to place an icon and menu in the notification area.   [![godoc][GoDoc]](https://godoc.org/github.com/getlantern/systray)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:494** OSX Desktop Notifications library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/gosx-notifier)
* [trayhost](https://github.com/shurcooL/trayhost) **star:162** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:4** Go bindings for libappindicator3 C library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/dawidd6/go-appindicator)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:1** OSX library to notify about any (pluggable) activity on your machine.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:1** OSX Sleep/Wake notifications in golang.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [imaging](https://github.com/disintegration/imaging) **star:2649** Simple Go image processing package.   [![star > 2000][Awesome]](https://github.com/disintegration/imaging)   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/imaging)
* [imaginary](https://github.com/h2non/imaginary) **star:2639** Fast and simple HTTP microservice for image resizing.   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [bild](https://github.com/anthonynsimon/bild) **star:2632** Collection of image processing algorithms in pure Go.   [![star > 2000][Awesome]](https://github.com/anthonynsimon/bild)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/anthonynsimon/bild)
* [gocv](https://github.com/hybridgroup/gocv) **star:2563** Go package for computer vision using OpenCV 3.3+.   [![star > 2000][Awesome]](https://github.com/hybridgroup/gocv)   [![godoc][GoDoc]](https://godoc.org/github.com/hybridgroup/gocv)
* [ln](https://github.com/fogleman/ln) **star:2488** 3D line art rendering in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/ln)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/ln)
* [resize](https://github.com/nfnt/resize) **star:2155** Image resizing for Go with common interpolation methods.   [![star > 2000][Awesome]](https://github.com/nfnt/resize)   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/nfnt/resize)
* [gg](https://github.com/fogleman/gg) **star:1954** 2D rendering in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/gg)
* [pt](https://github.com/fogleman/pt) **star:1779** Path tracing engine written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/pt)
* [svgo](https://github.com/ajstarks/svgo) **star:1355** Go Language Library for SVG generation.   [![godoc][GoDoc]](https://godoc.org/github.com/ajstarks/svgo)
* [smartcrop](https://github.com/muesli/smartcrop) **star:1268** Finds good crops for arbitrary images and crop sizes.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/smartcrop)
* [gift](https://github.com/disintegration/gift) **star:1246** Package of image processing filters.   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/gift)
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1105** Go bindings for OpenCV.   [![godoc][GoDoc]](https://godoc.org/github.com/lazywei/go-opencv)
* [picfit](https://github.com/thoas/picfit) **star:1091** An image resizing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/picfit)
* [geopattern](https://github.com/pravj/geopattern) **star:1013** Create beautiful generative image patterns from a string.   [![godoc][GoDoc]](https://godoc.org/github.com/pravj/geopattern)
* [imagick](https://github.com/gographics/imagick) **star:998** Go binding to ImageMagick's MagickWand C API.   [![godoc][GoDoc]](https://godoc.org/github.com/gographics/imagick)
* [bimg](https://github.com/h2non/bimg) **star:817** Small package for fast and efficient image processing using libvips.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:535** Go tool for LSB steganography, capable of hiding any file within an image.   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [mort](https://github.com/aldor007/mort) **star:367** Storage and image processing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aldor007/mort)
* [govatar](https://github.com/o1egl/govatar) **star:312** Library and CMD tool for generating funny avatars.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/govatar)
* [image2ascii](https://github.com/qeesung/image2ascii) **star:303** Convert image to ASCII.   [![godoc][GoDoc]](https://godoc.org/github.com/qeesung/image2ascii)
* [go-nude](https://github.com/koyachi/go-nude) **star:287** Nudity detection with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/koyachi/go-nude)
* [goimagehash](https://github.com/corona10/goimagehash) **star:234** Go Perceptual image hashing package.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [rez](https://github.com/bamiaux/rez) **star:191** Image resizing in pure Go and SIMD.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bamiaux/rez)
* [img](https://github.com/hawx/img) **star:129** Selection of image manipulation tools.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hawx/img)
* [go-cairo](https://github.com/ungerik/go-cairo) **star:87** Go binding for the cairo graphics library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [darkroom](https://github.com/gojek/darkroom) **star:85** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [mergi](https://github.com/noelyahan/mergi) **star:77** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).   [![godoc][GoDoc]](https://godoc.org/github.com/noelyahan/mergi)
* [go-gd](https://github.com/bolknote/go-gd) **star:51** Go binding for GD library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bolknote/go-gd)
* [gltf](https://github.com/qmuntal/gltf) **star:42** Efficient and robust glTF 2.0 reader, writer and validator.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [cameron](https://github.com/aofei/cameron) **star:33** An avatar generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)
* [goimghdr](https://github.com/corona10/goimghdr) **star:31** The imghdr module determines the type of image contained in a file for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimghdr)
* [steganography](https://github.com/auyer/steganography) **star:27** Pure Go Library for LSB steganography.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/auyer/steganography)
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.   [![godoc][GoDoc]](https://godoc.org/github.com/donatj/mpo)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1166** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt) **star:820** Gatt is a Go package for building Bluetooth Low Energy peripherals.   [![godoc][GoDoc]](https://godoc.org/github.com/paypal/gatt)
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [mainflux](https://github.com/Mainflux/mainflux) **star:624** Industrial IoT Messaging and Device Management Server.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Mainflux/mainflux)
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [devices](https://github.com/goiot/devices) **star:226** Suite of libraries for IoT devices, experimental for x/exp/io.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/goiot/devices)
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:181** Lightweight stream processing engine for IoT.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/sensorbee/sensorbee)
* [connectordb](https://github.com/connectordb/connectordb) **star:174** Open-Source Platform for Quantified Self & IoT.   [![godoc][GoDoc]](https://godoc.org/github.com/connectordb/connectordb)
* [huego](https://github.com/amimof/huego) **star:115** An extensive Philips Hue client library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.

## Job Scheduler

*Libraries for scheduling jobs.*

* [gron](https://github.com/roylee0704/gron) **star:647** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   [![godoc][GoDoc]](https://godoc.org/github.com/roylee0704/gron)
* [JobRunner](https://github.com/bamzi/jobrunner) **star:598** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/bamzi/jobrunner)
* [jobs](https://github.com/albrow/jobs) **star:452** Persistent and flexible background jobs library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/jobs)
* [scheduler](https://github.com/carlescere/scheduler) **star:297** Cronjobs scheduling made easy.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/scheduler)
* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [go-cron](https://github.com/rk/go-cron) **star:177** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rk/go-cron)
* [clockwork](https://github.com/whiteShtef/clockwork) **star:85** Simple and intuitive job scheduling library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/whiteShtef/clockwork)
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:36** Job scheduler that supports webhooks, crons and classic scheduling.   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) **star:5096** Get a JSON value with one line of code.   [![star > 2000][Awesome]](https://github.com/tidwall/gjson)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/gjson)
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2063** Automatically generate Go (golang) struct definitions from example JSON.   [![star > 2000][Awesome]](https://github.com/ChimeraCoder/gojson)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/gojson)
* [gojq](https://github.com/elgs/gojq) **star:140** JSON query in Golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/gojq)
* [kazaam](https://github.com/Qntfy/kazaam) **star:133** API for arbitrary transformation of JSON documents.   [![godoc][GoDoc]](https://godoc.org/github.com/Qntfy/kazaam)
* [jsongo](https://github.com/ricardolonga/jsongo) **star:92** Fluent API to make it easier to create Json objects.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ricardolonga/jsongo)
* [gjo](https://github.com/skanehira/gjo) **star:63** Small utility to create JSON objects.   [![godoc][GoDoc]](https://godoc.org/github.com/skanehira/gjo)
* [jsonf](https://github.com/miolini/jsonf) **star:56** Console tool for highlighted formatting and struct query fetching JSON.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/jsonf)
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:53** JSON diff utility written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [go-respond](https://github.com/nicklaw5/go-respond) **star:26** Go package for handling common HTTP JSON responses.   [![godoc][GoDoc]](https://godoc.org/github.com/nicklaw5/go-respond)
* [ajson](https://github.com/spyzhov/ajson) **star:16** Abstract JSON for golang with JSONPath support.   [![godoc][GoDoc]](https://godoc.org/github.com/spyzhov/ajson)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:8** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.   [![godoc][GoDoc]](https://godoc.org/github.com/ddymko/go-jsonerror)

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) **star:12383** Structured logger for Go.   [![star > 2000][Awesome]](https://github.com/Sirupsen/logrus)   [![godoc][GoDoc]](https://godoc.org/github.com/Sirupsen/logrus)
* [zap](https://github.com/uber-go/zap) **star:7716** Fast, structured, leveled logging in Go.   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [spew](https://github.com/davecgh/go-spew) **star:3369** Implements a deep pretty printer for Go data structures to aid in debugging.   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [glog](https://github.com/golang/glog) **star:2370** Leveled execution logs for Go.   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [zerolog](https://github.com/rs/zerolog) **star:2349** Zero-allocation JSON logger.   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [tail](https://github.com/hpcloud/tail) **star:1568** Go package striving to emulate the features of the BSD tail program.   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1504** Simple rolling logger, implements io.WriteCloser.   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [seelog](https://github.com/cihub/seelog) **star:1374** Logging functionality with flexible dispatching, filtering, and formatting.   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log15](https://github.com/inconshreveable/log15) **star:920** Simple, powerful logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/log15)
* [log](https://github.com/apex/log) **star:743** Structured logging package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [logxi](https://github.com/mgutz/logxi) **star:336** 12-factor app logger that is fast and makes you happy.   [![godoc][GoDoc]](https://godoc.org/github.com/mgutz/logxi)
* [onelog](https://github.com/francoispqt/onelog) **star:335** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.   [![godoc][GoDoc]](https://godoc.org/github.com/francoispqt/onelog)
* [log](https://github.com/go-playground/log) **star:269** Simple, configurable and scalable Structured Logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [logutils](https://github.com/hashicorp/logutils) **star:252** Utilities for slightly better logging in Go (Golang) extending the standard logger.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [go-logger](https://github.com/apsdehal/go-logger) **star:238** Simple logger of Go Programs, with level handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/apsdehal/go-logger)
* [logger](https://github.com/azer/logger) **star:135** Minimalistic logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/azer/logger)
* [xlog](https://github.com/rs/xlog) **star:128** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:108** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-log)   ![Contains Chinese documents][CN]
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:101** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:82** Full-featured logging saas written in golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [glg](https://github.com/kpango/glg) **star:52** glg is simple and fast leveled logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kpango/glg)
* [gologger](https://github.com/sadlil/gologger) **star:39** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/gologger)   ![Archived][Archived]
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j implementation in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/go-log)
* [go-log](https://github.com/siddontang/go-log) **star:23** Log lib supports level and multi handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-log)
* [journald](https://github.com/ssgreg/journald) **star:20** Go implementation of systemd Journal's native API for logging.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/journald)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:20** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [gomol](https://github.com/aphistic/gomol) **star:15** Multiple-output, structured logging for Go with extensible logging outputs.   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [glo](https://github.com/lajosbencz/glo) **star:8** PHP Monolog inspired logging facility with identical severity levels.   [![godoc][GoDoc]](https://godoc.org/github.com/lajosbencz/glo)
* [xlog](https://github.com/xfxdev/xlog) **star:7** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [log](https://github.com/aerogo/log) **star:5** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/log)

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:6711** General Machine Learning library for Go.   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   ![Contains Chinese documents][CN]
* [gorgonia](https://github.com/chewxy/gorgonia) **star:2764** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![star > 2000][Awesome]](https://github.com/chewxy/gorgonia)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/chewxy/gorgonia)
* [tfgo](https://github.com/galeone/tfgo) **star:1215** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/tfgo)
* [goml](https://github.com/cdipaolo/goml) **star:1022** On-line Machine Learning in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gosseract](https://github.com/otiai10/gosseract) **star:918** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:649** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [bayesian](https://github.com/jbrukh/bayesian) **star:635** Naive Bayesian Classification for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:630** An evolutionary optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [gorse](https://github.com/zhenghaoz/gorse) **star:551** A High Performance Recommender System Package based on Collaborative Filtering for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   ![Contains Chinese documents][CN]
* [gobrain](https://github.com/goml/gobrain) **star:393** Neural Networks written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [regommend](https://github.com/muesli/regommend) **star:252** Recommendation & collaborative filtering engine.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [ocrserver](https://github.com/otiai10/ocrserver) **star:235** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/ocrserver)
* [go-deep](https://github.com/patrikeh/go-deep) **star:227** A feature-rich neural network library in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/patrikeh/go-deep)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:189** Go Interface to Open Neural Network Exchange (ONNX).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-galib](https://github.com/thoj/go-galib) **star:172** Genetic Algorithms library written in Go / golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:144** Recommendation Algorithms library written in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [shield](https://github.com/eaigner/shield) **star:124** Bayesian text classifier with flexible tokenizers and storage backends for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [go-fann](https://github.com/white-pony/go-fann) **star:100** Go bindings for Fast Artificial Neural Networks(FANN) library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [Goptuna](https://github.com/c-bata/goptuna) **star:89** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [goga](https://github.com/tomcraven/goga) **star:78** Genetic algorithm library for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:63** libsvm golang version derived work based on LIBSVM 3.14.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [neural-go](https://github.com/schuyler/neural-go) **star:61** Multilayer perceptron network implemented in Go, with training via backpropagation.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [go-pr](https://github.com/daviddengcn/go-pr) **star:57** Pattern recognition package in Go lang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [neat](https://github.com/jinyeom/neat) **star:55** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   ![Archived][Archived]
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [goscore](https://github.com/asafschers/goscore) **star:38** Go Scoring API for PMML.   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [fonet](https://github.com/Fontinalis/fonet) **star:33** A Deep Neural Network library written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Fontinalis/fonet)
* [evoli](https://github.com/khezen/evoli) **star:8** Genetic Algorithm and Particle Swarm Optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) **star:4795** Go library for Apache Kafka.   [![star > 2000][Awesome]](https://github.com/Shopify/sarama)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/sarama)
* [gorush](https://github.com/appleboy/gorush) **star:3784** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:3756** Real-time messaging (Websockets or SockJS) server in Go.   [![star > 2000][Awesome]](https://github.com/centrifugal/centrifugo)   [![godoc][GoDoc]](https://godoc.org/github.com/centrifugal/centrifugo)
* [machinery](https://github.com/RichardKnop/machinery) **star:3465** Asynchronous task queue/job queue based on distributed message passing.   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:2962** socket.io library for golang, a realtime application framework.   [![star > 2000][Awesome]](https://github.com/googollee/go-socket.io)
* [NATS Go Client](https://github.com/nats-io/nats) **star:2444** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.   [![star > 2000][Awesome]](https://github.com/nats-io/nats)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/nats)
* [APNs2](https://github.com/sideshow/apns2) **star:2077** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.   [![star > 2000][Awesome]](https://github.com/sideshow/apns2)   [![godoc][GoDoc]](https://godoc.org/github.com/sideshow/apns2)
* [Benthos](https://github.com/Jeffail/benthos) **star:2045** A message streaming bridge between a range of protocols.   [![star > 2000][Awesome]](https://github.com/Jeffail/benthos)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/benthos)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1845** gopush-cluster is a go push server cluster.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   ![Contains Chinese documents][CN]
* [melody](https://github.com/olahol/melody) **star:1601** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [Mercure](https://github.com/dunglas/mercure) **star:1566** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [mangos](https://github.com/go-mangos/mangos) **star:1541** Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mangos/mangos)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1482** the official Go package for NSQ.   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1103** Redis backed unified push service for server-side notifications to mobile devices.   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [zmq4](https://github.com/pebbe/zmq4) **star:789** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/zmq4)
* [Gollum](https://github.com/trivago/gollum) **star:776** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [Beaver](https://github.com/Clivern/Beaver) **star:736** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [EventBus](https://github.com/asaskevich/EventBus) **star:573** The lightweight event bus with async compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [golongpoll](https://github.com/jcuga/golongpoll) **star:432** HTTP longpoll server library that makes web pub-sub simple.   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [dbus](https://github.com/godbus/dbus) **star:368** Native Go bindings for D-Bus.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [Glue](https://github.com/desertbit/glue) **star:321** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [emitter](https://github.com/olebedev/emitter) **star:317** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:288** Simple pubsub package for go.   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [guble](https://github.com/smancke/guble) **star:138** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [Bus](https://github.com/mustafaturan/bus) **star:118** Minimalist message bus implementation for internal communication.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [oplog](https://github.com/dailymotion/oplog) **star:95** Generic oplog/replication system for REST APIs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [rabtap](https://github.com/jandelgado/rabtap) **star:77** RabbitMQ swiss army knife cli app.   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [messagebus](https://github.com/vardius/message-bus) **star:70** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/message-bus)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:63** A tiny wrapper over amqp exchanges and queues.   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [drone-line](https://github.com/appleboy/drone-line) **star:62** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-line)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:56** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:54** A tiny wrapper around NSQ topic and channel.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:48** Native implementation of the freedesktop notification spec.   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [Commander](https://github.com/jeroenrinzema/commander) **star:26** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [hub](https://github.com/leandro-lugaresi/hub) **star:25** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:12** Client library to Viessmann Vitotrol web service.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** Gaurun Client written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gaurun-client)
* [jazz](https://github.com/socifi/jazz) **star:6** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:4** redisqueue provides a producer and consumer of a queue that uses Redis streams.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/redisqueue)
* [rmqconn](https://github.com/sbabiv/rmqconn)  RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:1795** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:4626** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:3496** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:85** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:49** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:947** A reflection based dependency injection toolkit for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:851** A dependency injection based application framework for Go (built on top of dig).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [inject](https://github.com/defval/inject) **star:32** A reflection based dependency injection container with simple interface.   [![godoc][GoDoc]](https://godoc.org/github.com/defval/inject)
* [wire](https://github.com/Fs02/wire) **star:19** Strict Runtime Dependency Injection for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/wire)
* [linker](https://github.com/logrange/linker) **star:10** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [gocontainer](https://github.com/vardius/gocontainer) **star:9** Simple Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:9805** Set of common historical and emerging project layout patterns in the Go ecosystem.   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)
* [go-sample](https://github.com/zitryss/go-sample) **star:31** A sample layout for Go application projects with the real code.   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)
* [scaffold](https://github.com/catchplay/scaffold) **star:27** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:625** Collection of useful string functions ported from other languages.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:67** String utilities.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4058** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:2531** Library and command for making and extracting .zip and .tar.gz archives.   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gosms](https://github.com/haxpax/gosms) **star:1231** Your own local SMS gateway in Go that can be used to send SMS.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:881** Resiliency patterns for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:699** Generic object pool for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   ![Contains Chinese documents][CN]
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:647** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   ![Contains Chinese documents][CN]
* [shortid](https://github.com/teris-io/shortid) **star:464** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:450** Random data generator written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [llvm](https://github.com/llir/llvm) **star:430** Library for interacting with LLVM IR in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:364** Easy to use, extensible health check library.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:342** Package conv provides fast and intuitive conversions across Go types.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [banner](https://github.com/dimiro1/banner) **star:236** Add beautiful banners into your Go applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:216** Package that exposes country and subdivision data.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [antch](https://github.com/antchfx/antch) **star:149** A fast, powerful and extensible web crawling & scraping framework.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   ![Contains Chinese documents][CN]
* [battery](https://github.com/distatus/battery) **star:136** Cross-platform, normalized battery information library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:127** Beautify data display for Humans.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   ![Contains Chinese documents][CN]
* [stats](https://github.com/go-playground/stats) **star:124** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [lk](https://github.com/hyperboloide/lk) **star:123** A simple licensing library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [bitio](https://github.com/icza/bitio) **star:102** Highly optimized bit-level Reader and Writer for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:88** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:78** Emojis for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [gommit](https://github.com/antham/gommit) **star:77** Analyze git commit messages to ensure they follow defined patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:69** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [indigo](https://github.com/osamingo/indigo) **star:53** Distributed unique ID generator of using Sonyflake and encoded by Base58.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [morse](https://github.com/alwindoss/morse) **star:51** Library to convert to and from morse code.   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [captcha](https://github.com/steambap/captcha) **star:44** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [ghorg](https://github.com/gabrie30/ghorg) **star:40** Clone all repos from a GitHub org into a single directory.   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [persian](https://github.com/mavihq/persian) **star:33** Some utilities for Persian language in go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:30** GoLang Library for [Browser Capabilities Project](http://browscap.org/).   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [datacounter](https://github.com/miolini/datacounter) **star:29** Go counters for readers/writer/http.ResponseWriter.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [autoflags](https://github.com/artyom/autoflags) **star:24** Go package to automatically define command line flags from struct fields.   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg) Specs implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [gosh](https://github.com/osamingo/gosh) **star:17** Provide Go Statistics Handler, Struct, Measure Method.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [sandid](https://github.com/aofei/sandid) **star:15** Every grain of sand on earth has its own ID.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [gotoprom](https://github.com/cabify/gotoprom) **star:15** Type-safe metrics builder wrapper library for the official Prometheus client.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [metrics](https://github.com/pascaldekloe/metrics) **star:4** Library for metrics instrumentation and Prometheus exposition.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2300** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more.   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1095** Go efficient text segmentation; support english, chinese, japanese and other.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   ![Contains Chinese documents][CN]
* [when](https://github.com/olebedev/when) **star:933** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:845** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   ![Contains Chinese documents][CN]
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:546** CN Hanzi to Hanyu Pinyin converter.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:432** JP morphological analyzer written in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:362** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:352** Extract values from strings and fill your structs with nlp.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:261** Sentence tokenizer:  converts text into a list of sentences.   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:222** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-nlp](https://github.com/nuance/go-nlp) **star:79** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [getlang](https://github.com/rylans/getlang) **star:77** Fast natural language detection package.   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [gounidecode](https://github.com/fiam/gounidecode) **star:67** Unicode transliterator (also known as unidecode) for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [textcat](https://github.com/pebbe/textcat) **star:61** Go package for n-gram based text categorization, with support for utf-8 and raw text.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:60** ASCII transliterations of Unicode text.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:52** Implementation of the porter stemming algorithm.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:45** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [go2vec](https://github.com/danieldk/go2vec) **star:31** Reader and utility functions for word2vec embeddings.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [petrovich](https://github.com/striker2000/petrovich) **star:23** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [shamoji](https://github.com/osamingo/shamoji) **star:10** The shamoji is word filtering package written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)

## Networking

*Libraries for working with various layers of the network.*

* [kcptun](https://github.com/xtaci/kcptun) **star:10808** Extremely simple & fast udp tunnel based on KCP protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcptun)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcptun)
* [fasthttp](https://github.com/valyala/fasthttp) **star:9685** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [dns](https://github.com/miekg/dns) **star:3909** Go library for working with DNS.   [![star > 2000][Awesome]](https://github.com/miekg/dns)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3433** HTTPLabs let you inspect HTTP requests and forge responses.   [![star > 2000][Awesome]](https://github.com/gchaincl/httplab)   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/httplab)
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:3101** An implementation of the QUIC protocol in pure Go.   [![star > 2000][Awesome]](https://github.com/lucas-clemente/quic-go)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/lucas-clemente/quic-go)
* [gopacket](https://github.com/google/gopacket) **star:2958** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [webrtc](https://github.com/pions/webrtc) **star:2440** A pure Go implementation of the WebRTC API.   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2290** KCP - Fast and Reliable ARQ Protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [gobgp](https://github.com/osrg/gobgp) **star:1713** BGP implemented in the Go Programming Language.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [ssh](https://github.com/gliderlabs/ssh) **star:1136** Higher-level API for building SSH servers (wraps crypto/ssh).   [![godoc][GoDoc]](https://godoc.org/github.com/gliderlabs/ssh)
* [fortio](https://github.com/fortio/fortio) **star:931** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:863** Simple TUN/TAP library.   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [sftp](https://github.com/pkg/sftp) **star:771** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/sftp)
* [go-getter](https://github.com/hashicorp/go-getter) **star:750** Go library for downloading files or directories from various sources using a URL.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:682** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [mdns](https://github.com/hashicorp/mdns) **star:560** Simple mDNS (Multicast DNS) client/server library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/mdns)
* [grab](https://github.com/cavaliercoder/grab) **star:560** Go package for managing file downloads.   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [ftp](https://github.com/jlaffaye/ftp) **star:548** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [lhttp](https://github.com/fanux/lhttp) **star:516** Powerful websocket framework, build your IM server more easily.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   ![Contains Chinese documents][CN]
* [gosnmp](https://github.com/soniah/gosnmp) **star:444** Native Go library for performing SNMP actions.   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [gotcp](https://github.com/gansidui/gotcp) **star:422** Go package for quickly writing tcp applications.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/gotcp)
* [cidranger](https://github.com/yl2chen/cidranger) **star:391** Fast IP to CIDR lookup for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:380** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [gopcap](https://github.com/akrennmair/gopcap) **star:359** Go wrapper for libpcap.   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [go-stun](https://github.com/ccding/go-stun) **star:336** Go implementation of the STUN client (RFC 3489 and RFC 5389).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [raw](https://github.com/mdlayher/raw) **star:314** Package raw enables reading and writing data at the device driver level for a network interface.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [stun](https://github.com/go-rtc/stun) **star:297** Go implementation of RFC 5389 STUN protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:290** Go library for building tcp servers faster.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:234** Streaming protocolbuffer data over TCP made easy.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [winrm](https://github.com/masterzen/winrm) **star:223** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [arp](https://github.com/mdlayher/arp) **star:199** Package arp implements the ARP protocol, as described in RFC 826.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:189** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [utp](https://github.com/anacrolix/utp) **star:151** Go uTP micro transport protocol implementation.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/utp)
* [canopus](https://github.com/zubairhamed/canopus) **star:134** CoAP Client/Server implementation (RFC 7252).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [jazigo](https://github.com/udhos/jazigo) **star:127** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [sslb](https://github.com/eduardonunesp/sslb) **star:115** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gNxI](https://github.com/google/gnxi) **star:106** A collection of tools for Network Management that use the gNMI and gNOI protocols.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:102** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   ![Contains Chinese documents][CN]
* [xtcp](https://github.com/xfxdev/xtcp) **star:89** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:63** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/dhcp6)
* [ether](https://github.com/songgao/ether) **star:62** Cross-platform Go package for sending and receiving ethernet frames.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [packet](https://github.com/aerogo/packet) **star:35** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [iplib](https://github.com/c-robinson/iplib) **star:24** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [tspool](https://github.com/two/tspool) **star:5** A TCP Library use worker pool to improve performance and protect your server.   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)

### HTTP Clients

*Libraries for making HTTP requests.*

* [grequests](https://github.com/levigross/grequests) **star:1447** A Go "clone" of the great and famous Requests library.   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [heimdall](https://github.com/gojektech/heimdall) **star:1118** An enchanced http client with retry and hystrix capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/gojektech/heimdall)
* [sling](https://github.com/dghubble/sling) **star:1026** Sling is a Go HTTP client library for creating and sending API requests.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/sling)
* [gentleman](https://github.com/h2non/gentleman) **star:680** Full-featured plugin-driven HTTP client library.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gentleman)
* [pester](https://github.com/sethgrid/pester) **star:334** Go HTTP client calls with retries, backoff, and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/sethgrid/pester)
* [goreq](https://github.com/smallnest/goreq) **star:100** Enhanced simplified HTTP client based on gorequest.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/goreq)   ![Archived][Archived]
* [rq](https://github.com/ddo/rq) **star:31** A nicer interface for golang stdlib HTTP client.   [![godoc][GoDoc]](https://godoc.org/github.com/ddo/rq)

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) **star:752** Go bindings for GLFW 3.
* [gl](https://github.com/go-gl/gl) **star:649** Go bindings for OpenGL (generated via glow).   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/gl)
* [mathgl](https://github.com/go-gl/mathgl) **star:299** Pure Go math package specialized for 3D math, with inspiration from GLM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/mathgl)
* [goxjs/gl](https://github.com/goxjs/gl) **star:130** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/gl)
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:58** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) **star:15173** The fantastic ORM library for Golang, aims to be developer friendly.   [![star > 2000][Awesome]](https://github.com/jinzhu/gorm)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/gorm)
* [Xorm](https://github.com/go-xorm/xorm) **star:5359** Simple and powerful ORM for Go.   [![star > 2000][Awesome]](https://github.com/go-xorm/xorm)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-xorm/xorm)   ![Contains Chinese documents][CN]
* [go-pg](https://github.com/go-pg/pg) **star:3110** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [gorp](https://github.com/go-gorp/gorp) **star:3099** Go Relational Persistence, ORM-ish library for Go.   [![star > 2000][Awesome]](https://github.com/go-gorp/gorp)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorp/gorp)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2337** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   [![star > 2000][Awesome]](https://github.com/volatiletech/sqlboiler)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/sqlboiler)
* [upper.io/db](https://github.com/upper/db) **star:1923** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [reform](https://github.com/go-reform/reform) **star:815** Better ORM for Go, based on non-empty interfaces and code generation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-reform/reform)
* [pop/soda](https://github.com/gobuffalo/pop) **star:698** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/pop)
* [QBS](https://github.com/coocood/qbs) **star:540** Stands for Query By Struct. A Go ORM.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   ![Contains Chinese documents][CN]
* [go-queryset](https://github.com/jirfag/go-queryset) **star:457** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.   [![godoc][GoDoc]](https://godoc.org/github.com/jirfag/go-queryset)
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:248** A flexible and powerful SQL string builder library plus a zero-config ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/go-sqlbuilder)
* [Zoom](https://github.com/albrow/zoom) **star:244** Blazing-fast datastore and querying engine built on Redis.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [grimoire](https://github.com/Fs02/grimoire) **star:118** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/grimoire)
* [go-store](https://github.com/gosuri/go-store) **star:94** Simple and fast Redis backed key-value store library for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/go-store)
* [Marlow](https://github.com/dadleyy/marlow) **star:67** Generated ORM from project structs for compile time safety assurances.   [![godoc][GoDoc]](https://godoc.org/github.com/dadleyy/marlow)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:2** A simple ORM for Google/Firebase Cloud Firestore.   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:12689** Go dependency tool.   [![star > 2000][Awesome]](https://github.com/golang/dep)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/dep)
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) **star:7808** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.   [![star > 2000][Awesome]](https://github.com/Masterminds/glide)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/glide)
* [godep](https://github.com/tools/godep) **star:5648** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   [![star > 2000][Awesome]](https://github.com/tools/godep)   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   ![Archived][Archived]
* [govendor](https://github.com/kardianos/govendor) **star:4803** Go Package Manager. Go vendor tool that works with the standard vendor file.   [![star > 2000][Awesome]](https://github.com/kardianos/govendor)   [![godoc][GoDoc]](https://godoc.org/github.com/kardianos/govendor)
* [gopm](https://github.com/gpmgo/gopm) **star:2384** Go Package Manager.   [![star > 2000][Awesome]](https://github.com/gpmgo/gopm)   [![godoc][GoDoc]](https://godoc.org/github.com/gpmgo/gopm)   ![Archived][Archived]
* [gom](https://github.com/mattn/gom) **star:1355** Go Manager - bundle for go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [gpm](https://github.com/pote/gpm) **star:1205** Barebones dependency manager for Go.   ![It hasn't been updated in the last year][Yellow]
* [goop](https://github.com/nitrous-io/goop) **star:776** Simple dependency manager for Go (golang), inspired by Bundler.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/nitrous-io/goop)
* [nut](https://github.com/jingweno/nut) **star:245** Vendor Go dependencies.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** Minimal dependency version using Git.   ![It hasn't been updated in the last year][Yellow]
* [gigo](https://github.com/LyricalSecurity/gigo) **star:197** PIP-like dependency tool for golang, with support for private repositories and hashes.   [![godoc][GoDoc]](https://godoc.org/github.com/LyricalSecurity/gigo)
* [VenGO](https://github.com/DamnWidget/VenGO) **star:115** create and manage exportable isolated go virtual environments.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:89** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.
* [gop](https://github.com/lunny/gop) **star:50** Build and manage your Go applications out of GOPATH.   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/gop)   ![Contains Chinese documents][CN]   ![Archived][Archived]

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:8949** A distributed tracing system.   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1063** Simple profiling support package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [tracer](https://github.com/kamilsk/tracer) **star:11** Simple, lightweight tracing.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:5407** Implementation of GraphQL for Go.   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [graphql](https://github.com/neelance/graphql-go) **star:2873** GraphQL server with a focus on ease of use.   [![star > 2000][Awesome]](https://github.com/neelance/graphql-go)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/neelance/graphql-go)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:882** A simple Go package to Query over JSON Data.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [jsonql](https://github.com/elgs/jsonql) **star:205** JSON query expression library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/jsonql)
* [rql](https://github.com/a8m/rql) **star:114** Resource Query Language for REST API.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/a8m/rql)
* [graphql](https://github.com/tmc/graphql) **star:51** graphql parser + utilities.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [jsonslice](https://github.com/bhmj/jsonslice) **star:24** Jsonpath queries with advanced filters.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) **star:2237** The simple and easy way to embed static files into Go binaries.   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik) **star:2223** Embeds static files into a Go executable.   [![star > 2000][Awesome]](https://github.com/rakyll/statik)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/statik)
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1705** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:684** Generates a vfsdata.go file that statically implements the given virtual filesystem.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/vfsgen)
* [esc](https://github.com/mjibson/esc) **star:482** Embeds files into Go programs and provides http.FileSystem interfaces to them.   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:437** Simple tool to embed files in go with focus on "customization" and ease to use.   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-resources](https://github.com/omeid/go-resources) **star:158** Unfancy resources embedding with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/omeid/go-resources)
* [statics](https://github.com/go-playground/statics) **star:53** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/statics)
* [templify](https://github.com/wlbr/templify) **star:22** Embed external template files into Go code to create single file binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/templify)
* [go-embed](https://github.com/pyros2097/go-embed) **star:17** Generates go code to embed resource files into your library or executable.   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:3084** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [stats](https://github.com/montanaflynn/stats) **star:1374** Statistics package with common functions missing from the Golang standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gosl](https://github.com/cpmech/gosl) **star:1327** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![godoc][GoDoc]](https://godoc.org/github.com/cpmech/gosl)
* [streamtools](https://github.com/nytlabs/streamtools) **star:1313** general purpose, graphical tool for dealing with streams of data.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [gonum/plot](https://github.com/gonum/plot) **star:1243** gonum/plot provides an API for building and drawing plots in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [go-dsp](https://github.com/mjibson/go-dsp) **star:633** Digital Signal Processing for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/go-dsp)
* [goraph](https://github.com/gyuho/goraph) **star:601** Pure Go graph theory library(data structure, algorith visualization).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [chart](https://github.com/vdobler/chart) **star:589** Simple Chart Plotting library for Go. Supports many graphs types.   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [ewma](https://github.com/VividCortex/ewma) **star:269** Exponentially-weighted moving averages.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [graph](https://github.com/yourbasic/graph) **star:249** Library of basic graph algorithms.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [orb](https://github.com/paulmach/orb) **star:207** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:130** Approximate histograms for data streams.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:91** Dataframes for Go for machine-learning and statistics (similar to pandas).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
* [sparse](https://github.com/james-bowman/sparse) **star:74** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:69** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/DavidBelicza/TextRank)
* [pagerank](https://github.com/alixaxel/pagerank) **star:50** Weighted PageRank algorithm implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [geom](https://github.com/skelterjohn/geom) **star:41** 2D geometry for golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [evaler](https://github.com/soniah/evaler) **star:40** Simple floating point arithmetic expression evaluator.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:13** GO Implementation of Entropy Measures.   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:11** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [GoStats](https://github.com/OGFris/GoStats) **star:10** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [PiHex](https://github.com/claygod/PiHex) **star:9** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:6** Tiny linear interpolation library.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/piecewiselinear)
* [bradleyterry](https://github.com/seanhagen/bradleyterry)  Provides a Bradley-Terry Model for pairwise comparisons.
* [assocentity](https://github.com/ndabAP/assocentity) **star:4** Package assocentity returns the average distance from words to a given entity.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/ndabAP/assocentity)
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/xenolf/lego) **star:3584** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   [![star > 2000][Awesome]](https://github.com/xenolf/lego)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/xenolf/lego)
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:1867** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [acmetool](https://github.com/hlandau/acme) **star:1699** ACME (Let's Encrypt) client tool with automatic renewal.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/acme)
* [memguard](https://github.com/awnumar/memguard) **star:1570** A pure Go library for handling sensitive values in memory.   [![godoc][GoDoc]](https://godoc.org/github.com/awnumar/memguard)
* [secure](https://github.com/unrolled/secure) **star:1231** HTTP middleware for Go that facilitates some quick security wins.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/secure)
* [acra](https://github.com/cossacklabs/acra) **star:466** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [nacl](https://github.com/kevinburke/nacl) **star:452** Go implementation of the NaCL set of API's.   [![godoc][GoDoc]](https://godoc.org/github.com/kevinburke/nacl)
* [BadActor](https://github.com/jaredfolkins/badactor) **star:250** In-memory, application-driven jailer built in the spirit of fail2ban.   [![godoc][GoDoc]](https://godoc.org/github.com/jaredfolkins/badactor)
* [passlib](https://github.com/hlandau/passlib) **star:229** Futureproof password hashing library.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/passlib)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:197** encrypt/decrypt using ssh keys.   [![godoc][GoDoc]](https://godoc.org/github.com/ssh-vault/ssh-vault)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:157** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [go-yara](https://github.com/hillu/go-yara) **star:135** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [argon2pw](https://github.com/raja/argon2pw) **star:76** Argon2 password hash generation with constant-time password comparison.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:31** A probably paranoid package for securely hashing and encrypting passwords.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [certificates](https://github.com/mvmaasakkers/certificates) **star:10** An opinionated tool for generating tls certificates.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:7** SSL certificates made easy with a high level wrapper around acme/autocert.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianosela/sslmgr)
* [jwc](https://github.com/khezen/jwc) **star:5** JSON Web Cryptography library.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/jwc)   ![Archived][Archived]

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:5786** High-performance 100% compatible drop-in replacement of "encoding/json".   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:5312** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:3083** Protocol Buffers for Go with Gadgets.   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:2591** Go library for decoding generic map values into native Go structures.   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1259** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [colfer](https://github.com/pascaldekloe/colfer) **star:476** Code generation for the Colfer binary format.
* [csvutil](https://github.com/jszwec/csvutil) **star:314** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:273** Cap'n Proto library and parser for go.   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:122** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   [![godoc][GoDoc]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
* [structomap](https://github.com/tuvistavie/structomap) **star:100** Library to easily and dynamically generate maps from static structures.   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [bambam](https://github.com/glycerine/bambam) **star:60** generator for Cap'n Proto schemas from go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/bambam)
* [asn1](https://github.com/PromonLogicalis/asn1) **star:41** Asn.1 BER and DER encoding library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/PromonLogicalis/asn1)   ![Archived][Archived]
* [binstruct](https://github.com/ghostiam/binstruct) **star:9** Golang binary decoder for mapping data into the structure.   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [bel](https://github.com/32leaves/bel) **star:6** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.   [![godoc][GoDoc]](https://godoc.org/github.com/32leaves/bel)
* [pletter](https://github.com/vimeda/pletter) **star:6** A standard way to wrap a proto message for message brokers.   [![godoc][GoDoc]](https://godoc.org/github.com/vimeda/pletter)

## Server Applications

* [etcd](https://github.com/coreos/etcd) **star:27206** Highly-available key value store for shared configuration and service discovery.   [![star > 2000][Awesome]](https://github.com/coreos/etcd)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/coreos/etcd)
* [Caddy](https://github.com/mholt/caddy) **star:23633** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   [![star > 2000][Awesome]](https://github.com/mholt/caddy)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/caddy)
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [minio](https://github.com/minio/minio) **star:18000** Minio is a distributed object storage server.   [![star > 2000][Awesome]](https://github.com/minio/minio)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [RoadRunner](https://github.com/spiral/roadrunner) **star:3433** High-performance PHP application server, load-balancer and process manager.   [![star > 2000][Awesome]](https://github.com/spiral/roadrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/spiral/roadrunner)
* [devd](https://github.com/cortesi/devd) **star:2831** Local webserver for developers.   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [algernon](https://github.com/xyproto/algernon) **star:1604** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/algernon)
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:1042** Full featured and highly configurable SFTP server software.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/drakkan/sftpgo)
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.
* [flipt](https://github.com/markphelps/flipt) **star:999** A self contained feature flag solution written in Go and Vue.js   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Flagr](https://github.com/checkr/flagr) **star:865** Flagr is an open-source feature flagging and A/B testing service.   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [Fider](https://github.com/getfider/fider) **star:827** Fider is an open platform to collect and organize customer feedback.   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [jackal](https://github.com/ortuman/jackal) **star:734** An XMPP server written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [discovery](https://github.com/Bilibili/discovery) **star:705** A registry for resilient mid-tier load balancing and failover.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:8** Stream database events from PostgreSQL to Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:5** Nginx log parser and exporter to Prometheus.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.

## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:202** Go stream processing library.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3159** PDF document generator with high level support for text, drawing and images.   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)
* [pongo2](https://github.com/flosch/pongo2) **star:1509** Django-like template-engine for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1436** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [hero](https://github.com/shiyanhui/hero) **star:1213** Hero is a handy, fast and powerful go template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/shiyanhui/hero)   ![Contains Chinese documents][CN]
* [mustache](https://github.com/hoisie/mustache) **star:967** Go implementation of the Mustache template language.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:825** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:763** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:706** Razor view engine for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [jet](https://github.com/CloudyKit/jet) **star:585** Jet template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/CloudyKit/jet)
* [ego](https://github.com/benbjohnson/ego) **star:417** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [raymond](https://github.com/aymerick/raymond) **star:346** Complete handlebars implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:306** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasttemplate)
* [Soy](https://github.com/robfig/soy) **star:144** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [liquid](https://github.com/osteele/liquid) **star:84** Go implementation of Shopify Liquid templates.   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/liquid)
* [kasia.go](https://github.com/ziutek/kasia.go) **star:70** Templating system for HTML and other text documents - go implementation.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [goview](https://github.com/foolin/goview) **star:67** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/goview)
* [velvet](https://github.com/gobuffalo/velvet) **star:64** Complete handlebars implementation in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [Testify](https://github.com/stretchr/testify) **star:8473** Sacred extension to the standard go testing package.   [![star > 2000][Awesome]](https://github.com/stretchr/testify)   [![godoc][GoDoc]](https://godoc.org/github.com/stretchr/testify)
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [go-cmp](https://github.com/google/go-cmp) **star:1226** Package for comparing Go values in tests.   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1168** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   [![godoc][GoDoc]](https://godoc.org/github.com/gavv/httpexpect)
    * [godog](https://github.com/DATA-DOG/godog) **star:789** Cucumber or Behat like BDD framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/godog)
    * [baloo](https://github.com/h2non/baloo) **star:653** Expressive and versatile end-to-end HTTP API testing made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [goblin](https://github.com/franela/goblin) **star:633** Mocha like testing framework fo Go.   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:344** A helper for Rails' like test fixtures to test database applications.   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:336** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:301** Mutation testing for Go source code.   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gofight](https://github.com/appleboy/gofight) **star:266** API Handler Testing for Golang Router framework.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [frisby](https://github.com/verdverm/frisby) **star:249** REST API testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:198** Tool for viewing test coverage in terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [charlatan](https://github.com/percolate/charlatan) **star:190** Tool to generate fake interface implementations for tests.   [![godoc][GoDoc]](https://godoc.org/github.com/percolate/charlatan)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:128** A collection of packages to augment the go testing package and support common patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/gotestyourself/gotest.tools)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:111** BDD-style testing framework for the Go programming language.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [endly](https://github.com/viant/endly) **star:98** Declarative end to end functional testing.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:90** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:87** Simple snapshot testing addon for your test framework.   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [wstest](https://github.com/posener/wstest) **star:64** Websocket client for unit-testing a websocket http.Handler.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:55** Extremely flexible golang deep comparison, extends the go testing package.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [gospecify](https://github.com/stesla/gospecify) **star:51** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [commander](https://github.com/SimonBaeumer/commander) **star:35** Tool for testing cli applications on windows, linux and osx.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:30** library created for testing JSON against patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [dsunit](https://github.com/viant/dsunit) **star:25** Datastore testing for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:22** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [testcase](https://github.com/adamluzsi/testcase) **star:12** Idiomatic testing framework for Behavior Driven Development.   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [Tt](https://github.com/vcaesar/tt) **star:6** Simple and colorful test tools.   [![godoc][GoDoc]](https://godoc.org/github.com/vcaesar/tt)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:2** HTTP client testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)

* Mock
    * [gomock](https://github.com/golang/mock) **star:2991** Mocking framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/golang/mock)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/mock)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:1858** Mock SQL driver for testing database interactions.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1463** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:857** Versatile HTTP mocking made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [httpmock](https://github.com/jarcoal/httpmock) **star:618** Easy mocking of HTTP responses from external resources.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:369** Tool for generating self-contained mock objects.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [minimock](https://github.com/gojuno/minimock) **star:273** Mock generator for Go interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:175** Single transaction based database driver mainly for testing purposes.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:85** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:2975** Randomized testing system.   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz) **star:545** Library for populating go objects with random values.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gofuzz)
    * [Tavor](https://github.com/zimmski/tavor) **star:214** Generic fuzzing and delta-debugging framework.   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:3723** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1275** alternative Selenium hub server that launches browsers within containers.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [cdp](https://github.com/mafredri/cdp) **star:364** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:213** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:417** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:8789** Fast and Elegant Scraping Framework for Gophers.   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:7766** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [blackfriday](https://github.com/russross/blackfriday) **star:3978** Markdown processor in Go.   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [toml](https://github.com/BurntSushi/toml) **star:2843** TOML configuration format (encoder/decoder with reflection).   [![star > 2000][Awesome]](https://github.com/BurntSushi/toml)   [![godoc][GoDoc]](https://godoc.org/github.com/BurntSushi/toml)
    * [sh](https://github.com/mvdan/sh) **star:2074** Shell parser and formatter.   [![star > 2000][Awesome]](https://github.com/mvdan/sh)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/sh)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:1950** Formatters for time, numbers, and memory size to human readable format.   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1282** HTML Sanitizer.   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [inject](https://github.com/facebookgo/inject) **star:1151** Package inject provides a reflect based injector.   [![godoc][GoDoc]](https://godoc.org/github.com/facebookgo/inject)   ![Archived][Archived]
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1126** Parse RSS and Atom feeds in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [go-toml](https://github.com/pelletier/go-toml) **star:625** Go library for the TOML format with query support and handy cli tools.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:555** A collection of common regular expressions for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [slug](https://github.com/gosimple/slug) **star:394** URL-friendly slugify with multiple languages support.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gosimple/slug)
    * [mxj](https://github.com/clbanning/mxj) **star:338** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:310** Parses the Graphviz DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:300** Web scraping Framework to turn websites into structured data.   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [gotext](https://github.com/leonelquinteros/gotext) **star:233** GNU gettext utilities for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/leonelquinteros/gotext)
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:215** Functions to get fixed width of the character or string.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-runewidth)
    * [goq](https://github.com/andrewstuart/goq) **star:151** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).   [![godoc][GoDoc]](https://godoc.org/github.com/andrewstuart/goq)
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:142** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/htmlquery)
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:101** NMEA parser library for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [sdp](https://github.com/gortc/sdp) **star:73** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [align](https://github.com/Guitarbum722/align) **star:60** A general purpose application that aligns text.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:54** Make pretty slug with multiple languages support.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-slugify)
    * [genex](https://github.com/alixaxel/genex) **star:54** Count and expand Regular Expressions into all matching Strings.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:44** Functions to determine the natural language of a unicode text.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/guesslanguage)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:43** Editorconfig file parser and manipulator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:41** Zero-width character detection and removal for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/trubitsyn/go-zero-width)
    * [allot](https://github.com/sbstjn/allot) **star:34** Placeholder and wildcard text parsing for CLI tools and bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [gonameparts](https://github.com/polera/gonameparts) **star:29** Parses human names into individual name parts.   [![godoc][GoDoc]](https://godoc.org/github.com/polera/gonameparts)
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:27** Fixed-width text formatting (encoder/decoder with reflection).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [Slugify](https://github.com/avelino/slugify) **star:26** Go slugify application that handles string.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/avelino/slugify)
    * [go-vcard](https://github.com/emersion/go-vcard) **star:25** Parse and format vCard.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [did](https://github.com/ockam-network/did) **star:24** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [codetree](https://github.com/aerogo/codetree) **star:10** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
    * [ltsv](https://github.com/Wing924/ltsv) **star:2** High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/ltsv)
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:464** Extract urls from text.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/xurls)
    * [gotabulate](https://github.com/bndr/gotabulate) **star:203** Easily pretty-print your tabular data with Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:145** fast string sorting algorithm.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [parth](https://github.com/codemodus/parth) **star:31** URL path segmentation parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/parth)
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:17** Convert xml or json to go struct.   [![godoc][GoDoc]](https://godoc.org/github.com/stackerzzq/xj2go)
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:15** A sanitization-based swear filter for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
    * [TySug](https://github.com/Dynom/TySug) **star:3** Alternative suggestions with respect to keyboard layouts.   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
    * [Tagify](https://github.com/zoomio/tagify) **star:2** Produces a set of tags from given source.   [![godoc][GoDoc]](https://godoc.org/github.com/zoomio/tagify)

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5142** The official AWS SDK for the Go programming language.   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [github](https://github.com/google/go-github) **star:4967** Go library for accessing the GitHub REST API v3.   [![star > 2000][Awesome]](https://github.com/google/go-github)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [slack](https://github.com/nlopes/slack) **star:2497** Slack API in Go.   [![star > 2000][Awesome]](https://github.com/nlopes/slack)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/nlopes/slack)
* [google](https://github.com/google/google-api-go-client) **star:1962** Auto-generated Google APIs for Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:1846** Google Cloud APIs Go Client Library.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
* [discordgo](https://github.com/bwmarrin/discordgo) **star:1007** Go bindings for the Discord Chat API.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/bwmarrin/discordgo)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:996** Go client library for the Twitter 1.1 API.   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [stripe](https://github.com/stripe/stripe-go) **star:971** Go client for the Stripe API.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [facebook](https://github.com/huandu/facebook) **star:778** Go Library that supports the Facebook Graph API.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/facebook)
* [minio-go](https://github.com/minio/minio-go) **star:743** Minio Go Library for Amazon S3 compatible cloud storage.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio-go)
* [go-twitter](https://github.com/dghubble/go-twitter) **star:732** Go client library for the Twitter v1.1 APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/go-twitter)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:596** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-jira)
* [githubql](https://github.com/shurcooL/githubql) **star:521** Go library for accessing the GitHub GraphQL API v4.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/githubql)
* [webhooks](https://github.com/go-playground/webhooks) **star:375** Webhook receiver for GitHub and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/webhooks)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:310** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:308** Wrapper for PayPal payment API.   [![godoc][GoDoc]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
* [go-marathon](https://github.com/gambol99/go-marathon) **star:191** Go library for interacting with Mesosphere's Marathon PAAS.   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [ethrpc](https://github.com/onrik/ethrpc) **star:175** Go bindings for Ethereum JSON RPC API.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [gostorm](https://github.com/jsgilmore/gostorm) **star:121** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [Medium](https://github.com/Medium/medium-sdk-go) **star:116** Golang SDK for Medium's OAuth2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/Medium/medium-sdk-go)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:113** A golang package to communicate with HipChat over XMPP.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [Trello](https://github.com/adlio/trello) **star:109** Go wrapper for the Trello API.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [hipchat](https://github.com/andybons/hipchat) **star:109** This project implements a golang client library for the Hipchat API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/andybons/hipchat)
* [go-trending](https://github.com/andygrunwald/go-trending) **star:100** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-trending)
* [cachet](https://github.com/andygrunwald/cachet) **star:67** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/cachet)
* [clarifai](https://github.com/samuelcouch/clarifai) **star:58** Go client library for interfacing with the Clarifai API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/samuelcouch/clarifai)
* [pushover](https://github.com/gregdel/pushover) **star:57** Go wrapper for the Pushover API.   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [megos](https://github.com/andygrunwald/megos) **star:57** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:53** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [wit-go](https://github.com/wit-ai/wit-go) **star:49** Go client for wit.ai HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [gads](https://github.com/emiddleton/gads) **star:43** Google Adwords Unofficial API.   [![godoc][GoDoc]](https://godoc.org/github.com/emiddleton/gads)
* [circleci](https://github.com/jszwedko/go-circleci) **star:41** Go client library for interacting with CircleCI's API.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:39** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:37** Go MusicBrainz WS2 client library.   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:35** Go wrapper and command-line client for the Uptime Robot v2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [fcm](https://github.com/maddevsio/fcm) **star:32** Go library for Firebase Cloud Messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [golyrics](https://github.com/mamal72/golyrics) **star:31** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [mixpanel](https://github.com/dukex/mixpanel) **star:30** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [gami](https://github.com/bit4bit/gami) **star:26** Go library for Asterisk Manager Interface.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   ![Archived][Archived]
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:21** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [patreon-go](https://github.com/mxpv/patreon-go) **star:20** Go library for Patreon API.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mxpv/patreon-go)
* [codeship-go](https://github.com/codeship/codeship-go) **star:15** Go client library for interacting with Codeship's API v2.   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:15** Go wrapper for the YNAB API.   [![godoc][GoDoc]](https://godoc.org/github.com/brunomvsouza/ynab.go)
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:13** Go client library for [imgur](https://imgur.com)   [![godoc][GoDoc]](https://godoc.org/github.com/koffeinsource/go-imgur)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [simples3](https://github.com/rhnvrm/simples3) **star:11** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.   [![godoc][GoDoc]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:1** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:23858** Command-line fuzzy finder written in Go.   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [hub](https://github.com/github/hub) **star:17361** wrap git commands with additional functionality to interact with github from the terminal.   [![star > 2000][Awesome]](https://github.com/github/hub)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/github/hub)
* [delve](https://github.com/derekparker/delve) **star:12354** Go debugger.   [![star > 2000][Awesome]](https://github.com/derekparker/delve)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [ctop](https://github.com/bcicen/ctop) **star:8930** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [wuzz](https://github.com/asciimoo/wuzz) **star:8337** Interactive cli tool for HTTP inspection.   [![star > 2000][Awesome]](https://github.com/asciimoo/wuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/wuzz)
* [sqlx](https://github.com/jmoiron/sqlx) **star:6982** provides a set of extensions on top of the excellent built-in database/sql package.   [![star > 2000][Awesome]](https://github.com/jmoiron/sqlx)   [![godoc][GoDoc]](https://godoc.org/github.com/jmoiron/sqlx)
* [peco](https://github.com/peco/peco) **star:5517** Simplistic interactive filtering tool.   [![star > 2000][Awesome]](https://github.com/peco/peco)   [![godoc][GoDoc]](https://godoc.org/github.com/peco/peco)
* [usql](https://github.com/knq/usql) **star:4692** usql is a universal command-line interface for SQL databases.   [![star > 2000][Awesome]](https://github.com/knq/usql)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/usql)
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:4632** Deliver Go binaries as fast and easily as possible.   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [godropbox](https://github.com/dropbox/godropbox) **star:3753** Common libraries for writing Go services/applications from Dropbox.   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [realize](https://github.com/tockins/realize) **star:3223** Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.   [![star > 2000][Awesome]](https://github.com/tockins/realize)   [![godoc][GoDoc]](https://godoc.org/github.com/tockins/realize)
* [goreporter](https://github.com/wgliang/goreporter) **star:2508** Golang tool that does static analysis, unit testing, code review and generate code quality report.   [![star > 2000][Awesome]](https://github.com/wgliang/goreporter)   [![godoc][GoDoc]](https://godoc.org/github.com/wgliang/goreporter)
* [panicparse](https://github.com/maruel/panicparse) **star:2159** Groups similar goroutines and colorizes stack dump.   [![star > 2000][Awesome]](https://github.com/maruel/panicparse)   [![godoc][GoDoc]](https://godoc.org/github.com/maruel/panicparse)
* [resty](https://github.com/go-resty/resty) **star:2114** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2080** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [Task](https://github.com/go-task/task) **star:2001** simple "Make" alternative.   [![star > 2000][Awesome]](https://github.com/go-task/task)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-task/task)
* [minify](https://github.com/tdewolff/minify) **star:1874** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/minify)
* [mmake](https://github.com/tj/mmake) **star:1454** Modern Make.   [![godoc][GoDoc]](https://godoc.org/github.com/tj/mmake)
* [Storm](https://github.com/asdine/storm) **star:1384** Simple and powerful toolkit for BoltDB.   [![godoc][GoDoc]](https://godoc.org/github.com/asdine/storm)
* [mole](https://github.com/davrodpin/mole) **star:1309** cli app to easily create ssh tunnels.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/davrodpin/mole)
* [go-funk](https://github.com/thoas/go-funk) **star:1271** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [mc](https://github.com/minio/mc) **star:1146** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/minio/mc)
* [filetype](https://github.com/h2non/filetype) **star:975** Small package to infer the file type checking the magic numbers signature.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [boilr](https://github.com/tmrts/boilr) **star:954** Blazingly fast CLI tool for creating projects from boilerplate templates.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [mergo](https://github.com/imdario/mergo) **star:878** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   [![godoc][GoDoc]](https://godoc.org/github.com/imdario/mergo)
* [spinner](https://github.com/briandowns/spinner) **star:821** Go package to easily provide a terminal spinner with options.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/briandowns/spinner)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:802** Circuit Breakers in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:729** Simple, seamless, lightweight time tracking for Git.   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [jump](https://github.com/gsamokovarov/jump) **star:681** Jump helps you navigate faster by learning your habits.   [![godoc][GoDoc]](https://godoc.org/github.com/gsamokovarov/jump)
* [immortal](https://github.com/immortal/immortal) **star:612** \*nix cross-platform (OS agnostic) supervisor.   [![godoc][GoDoc]](https://godoc.org/github.com/immortal/immortal)
* [htcat](https://github.com/htcat/htcat) **star:482** Parallel and Pipelined HTTP GET Utility.   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [go-dry](https://github.com/ungerik/go-dry) **star:435** DRY (don't repeat yourself) package for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [gopencils](https://github.com/bndr/gopencils) **star:423** Small and simple package to easily consume REST APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [godaemon](https://github.com/VividCortex/godaemon) **star:406** Utility to write daemons.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/godaemon)
* [circuit](https://github.com/cep21/circuit) **star:372** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/cep21/circuit)
* [request](https://github.com/mozillazg/request) **star:357** Go HTTP Requests for Humans™.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/request)
* [ergo](https://github.com/cristianoliveira/ergo) **star:319** The management of multiple local services running over different ports made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [koazee](https://github.com/wesovilabs/koazee) **star:307** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/koazee)
* [go-rate](https://github.com/beefsack/go-rate) **star:294** Timed rate limiter for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [gohper](https://github.com/cosiner/gohper) **star:248** Various tools/modules help for development.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/gohper)   ![Archived][Archived]
* [clockwork](https://github.com/jonboulle/clockwork) **star:224** A simple fake clock for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:217** Simple struct copying for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [serve](https://github.com/syntaqx/serve) **star:193** A static http server anywhere you need.   [![godoc][GoDoc]](https://godoc.org/github.com/syntaqx/serve)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:186** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [retry](https://github.com/kamilsk/retry) **star:173** The most advanced functional mechanism to perform actions repetitively until successful.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/retry)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:161** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [rerun](https://github.com/ivpusic/rerun) **star:153** Recompiling and rerunning go apps when source changes.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/rerun)
* [gubrak](https://github.com/novalagung/gubrak) **star:148** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [moldova](https://github.com/StabbyCutyou/moldova) **star:146** Utility for generating random data based on an input template.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/StabbyCutyou/moldova)
* [gotenv](https://github.com/subosito/gotenv) **star:145** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:139** Package for MIME type detection based on magic numbers.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gabriel-vasile/mimetype)
* [util](https://github.com/shomali11/util) **star:138** Collection of useful utility functions. (strings, concurrency, manipulations, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/util)
* [Death](https://github.com/vrecan/death) **star:137** Managing go application shutdown with signals.   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [robustly](https://github.com/VividCortex/robustly) **star:137** Runs functions resiliently, catching and restarting panics.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/robustly)
* [apm](https://github.com/topfreegames/apm) **star:130** Process manager for Golang applications with an HTTP API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/apm)
* [chyle](https://github.com/antham/chyle) **star:109** Changelog generator using a git repository with multiple configuration possibilities.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:106** XML Sitemap generator written in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:100** LiveReload server for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/jaschaephraim/lrserver)
* [toolbox](https://github.com/viant/toolbox) **star:100** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/toolbox)
* [onecache](https://github.com/adelowo/onecache) **star:99** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).   [![godoc][GoDoc]](https://godoc.org/github.com/adelowo/onecache)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:87** Pure Go bsdiff and bspatch libraries and CLI tools.   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [pm](https://github.com/VividCortex/pm) **star:72** Process (i.e. goroutine) manager with an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/pm)
* [UNIS](https://github.com/esemplastic/unis) **star:69** Common Architecture™ for String Utilities in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/esemplastic/unis)
* [xferspdy](https://github.com/monmohan/xferspdy) **star:67** Xferspdy provides binary diff and patch library in golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/monmohan/xferspdy)
* [go-health](https://github.com/Talento90/go-health) **star:63** Health package simplifies the way you add health check to your services.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/Talento90/go-health)
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:60** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/mssqlx)
* [multitick](https://github.com/VividCortex/multitick) **star:58** Multiplexor for aligned tickers.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/multitick)
* [repeat](https://github.com/ssgreg/repeat) **star:56** Go implementation of different backoff strategies useful for retrying operations and heartbeating.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/repeat)
* [abutil](https://github.com/bahlo/abutil) **star:51** Collection of often-used Golang helpers.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bahlo/abutil)   ![Archived][Archived]
* [minquery](https://github.com/icza/minquery) **star:50** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/minquery)
* [handy](https://github.com/miguelpragier/handy) **star:47** Many utilities and helpers like string handlers/formatters and validators.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** Parse TODOs in your GO code.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:42** Pure Go ultra performant MIME sniffing library/utility.   [![godoc][GoDoc]](https://godoc.org/github.com/zRedShift/mimemagic)
* [golog](https://github.com/mlimaloureiro/golog) **star:42** Easy and lightweight CLI tool to time track your tasks.   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [goback](https://github.com/carlescere/goback) **star:40** Go simple exponential backoff package.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [gaper](https://github.com/maxcnunes/gaper) **star:40** Builds and restarts a Go project when it crashes or some watched file changes.   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:38** Universal multi-workstation clipboard that uses S3 like backend for the storage.   [![godoc][GoDoc]](https://godoc.org/github.com/jutkko/copy-pasta)
* [goreadability](https://github.com/philipjkim/goreadability) **star:35** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [myhttp](https://github.com/inancgumus/myhttp) **star:33** Simple API to make HTTP GET requests with timeout support.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/inancgumus/myhttp)   ![Archived][Archived]
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:29** Retrying made simple and easy for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/retry-go)
* [rclient](https://github.com/zpatrick/rclient) **star:27** Readable, flexible, simple-to-use client for REST APIs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rclient)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:27** SeaweedFS client library with almost full features.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [pgo](https://github.com/arthurkushman/pgo) **star:24** Convenient functions for PHP community.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/pgo)
* [gostrutils](https://github.com/ik5/gostrutils) **star:17** Collections of string manipulation and conversion functions.   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [dbt](https://github.com/nikogura/dbt) **star:16** A framework for running self-updating signed binaries from a central, trusted repository.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [scan](https://github.com/blockloop/scan) **star:14** Scan golang `sql.Rows` directly to structs, slices, or primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/blockloop/scan)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:14** Go library for encoding structs into Header fields.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [filter](https://github.com/gookit/filter) **star:13** provide filtering, sanitizing, and conversion of Go data.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [ghokin](https://github.com/antham/ghokin) **star:12** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [retry](https://github.com/shafreeck/retry) **star:10** A pretty simple library to ensure your work to be done.   [![godoc][GoDoc]](https://godoc.org/github.com/shafreeck/retry)
* [ctxutil](https://github.com/posener/ctxutil) **star:8** A collection of utility functions for contexts.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [backscanner](https://github.com/icza/backscanner) **star:8** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:7** A MIME type sniffer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/mimesniffer)
* [sslice](https://github.com/yaa110/sslice) **star:5** Create a slice which is always sorted.   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/sslice)
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.   [![godoc][GoDoc]](https://godoc.org/github.com/chrispassas/silk)
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:3** Slice conversion between primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
* [slicer](https://github.com/leaanthony/slicer) **star:3** Makes working with slices easier.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/slicer)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:3** Go package for working with Problem Details.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
* [retry](https://github.com/percolate/retry) **star:2** A simple but highly configurable retry package for Go.
* [olaf](https://github.com/btnguyen2k/olaf) **star:1** Twitter Snowflake implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/btnguyen2k/olaf)
* [blank](https://github.com/Henry-Sarabia/blank) **star:1** Verify or remove blanks and whitespace from strings.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)

## UUID

*Libraries for working with UUIDs.*

* [ulid](https://github.com/oklog/ulid) **star:1708** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   [![godoc][GoDoc]](https://godoc.org/github.com/oklog/ulid)
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands. 
* [uuid](https://github.com/gofrs/uuid) **star:587** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   [![godoc][GoDoc]](https://godoc.org/github.com/gofrs/uuid)
* [wuid](https://github.com/edwingeng/wuid) **star:298** An extremely fast unique number generator, 10-135 times faster than UUID.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/wuid)
* [goid](https://github.com/jakehl/goid) **star:23** Generate and Parse RFC4122 compliant V4 UUIDs.   [![godoc][GoDoc]](https://godoc.org/github.com/jakehl/goid)
* [sno](https://github.com/muyo/sno) **star:16** Compact, sortable and fast unique IDs with embedded metadata.   [![godoc][GoDoc]](https://godoc.org/github.com/muyo/sno)
* [uuid](https://github.com/agext/uuid) **star:10** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/uuid)
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:7** A tiny and efficient Go unique string ID generator.   [![godoc][GoDoc]](https://godoc.org/github.com/aidarkhanov/nanoid)

## Validation

*Libraries for validation.*

* [validator](https://github.com/go-playground/validator) **star:3711** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   [![star > 2000][Awesome]](https://github.com/go-playground/validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/validator)
* [govalidator](https://github.com/asaskevich/govalidator) **star:3678** Validators and sanitizers for strings, numerics, slices and structs.   [![star > 2000][Awesome]](https://github.com/asaskevich/govalidator)   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/govalidator)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1090** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:738** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/govalidator)
* [validate](https://github.com/gookit/validate) **star:119** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/validate)   ![Contains Chinese documents][CN]
* [checkdigit](https://github.com/osamingo/checkdigit) **star:46** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/checkdigit)
* [jio](https://github.com/faceair/jio) **star:24** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   [![godoc][GoDoc]](https://godoc.org/github.com/faceair/jio)   ![Contains Chinese documents][CN]
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:19** A norms and conventions validator for Terraform.   [![godoc][GoDoc]](https://godoc.org/github.com/thazelart/terraform-validator)
* [validate](https://github.com/gobuffalo/validate) **star:19** This package provides a framework for writing validations for Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/validate)

## Version Control

*Libraries for version control.*

* [go-git](https://github.com/src-d/go-git) **star:4381** highly extensible Git implementation in pure Go.   [![star > 2000][Awesome]](https://github.com/src-d/go-git)   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/go-git)
* [git2go](https://github.com/libgit2/git2go) **star:1379** Go bindings for libgit2.   [![godoc][GoDoc]](https://godoc.org/github.com/libgit2/git2go)
* [hercules](https://github.com/src-d/hercules) **star:626** gaining advanced insights from Git repository history.   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/hercules)
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:72** manipulate and inspect VCS repositories in Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/sourcegraph/go-vcs)
* [gh](https://github.com/rjeczalik/gh) **star:70** Scriptable server and net/http middleware for GitHub Webhooks.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/gh)

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) **star:836** Comphrensive Go bindings for FFmpeg.   [![godoc][GoDoc]](https://godoc.org/github.com/giorgisio/goav)
* [gmf](https://github.com/3d0c/gmf) **star:545** Go bindings for FFmpeg av\* libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/3d0c/gmf)
* [go-astits](https://github.com/asticode/go-astits) **star:265** Parse and demux MPEG Transport Streams (.ts) natively in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astits)
* [go-astisub](https://github.com/asticode/go-astisub) **star:188** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astisub)
* [gst](https://github.com/ziutek/gst) **star:153** Go bindings for GStreamer.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/gst)
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:71** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/libvlc-go)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:38** Parser and generator library for Apple m3u8 playlists.   [![godoc][GoDoc]](https://godoc.org/github.com/quangngotan95/go-m3u8)
* [v4l](https://github.com/korandiz/v4l) **star:31** Video capture library for Linux, written in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/korandiz/v4l)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:11** Subtitle format support for go. Supports .srt, .ttml, and .ass.   [![godoc][GoDoc]](https://godoc.org/github.com/wargarblgarbl/libgosubs)

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) **star:31167** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   [![star > 2000][Awesome]](https://github.com/gin-gonic/gin)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gin-gonic/gin)
* [Beego](https://github.com/astaxie/beego) **star:21885** beego is an open-source, high-performance web framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/astaxie/beego)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/beego)   ![Contains Chinese documents][CN]
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) **star:15036** High performance, minimalist Go web framework.   [![star > 2000][Awesome]](https://github.com/labstack/echo)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/labstack/echo)
* [Revel](https://github.com/revel/revel) **star:11316** High-productivity web framework for the Go language.   [![star > 2000][Awesome]](https://github.com/revel/revel)   [![godoc][GoDoc]](https://godoc.org/github.com/revel/revel)
* [Goa](https://github.com/goadesign/goa) **star:3537** Goa provides a holistic approach for developing remote APIs and microservices in Go.   [![star > 2000][Awesome]](https://github.com/goadesign/goa)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/goadesign/goa)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3346** Quick and easy way to setup a RESTful JSON API.   [![star > 2000][Awesome]](https://github.com/ant0ine/go-json-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ant0ine/go-json-rest)
* [Gizmo](https://github.com/NYTimes/gizmo) **star:2877** Microservice toolkit used by the New York Times.   [![star > 2000][Awesome]](https://github.com/NYTimes/gizmo)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/NYTimes/gizmo)
* [Macaron](https://github.com/go-macaron/macaron) **star:2840** Macaron is a high productive and modular design web framework in Go.   [![star > 2000][Awesome]](https://github.com/go-macaron/macaron)   [![godoc][GoDoc]](https://godoc.org/github.com/go-macaron/macaron)
* [utron](https://github.com/gernest/utron) **star:2138** Lightweight MVC framework for Go(Golang).   [![star > 2000][Awesome]](https://github.com/gernest/utron)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/utron)
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:996** Go framework for building JSON web services inspired by Dropwizard.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rcrowley/go-tigertonic)
* [tango](https://github.com/lunny/tango) **star:819** Micro & pluggable web framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/tango)   ![Contains Chinese documents][CN]
* [traffic](https://github.com/pilu/traffic) **star:517** Sinatra inspired regexp/pattern mux and web framework for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/pilu/traffic)
* [gongular](https://github.com/mustafaakin/gongular) **star:418** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaakin/gongular)
* [neo](https://github.com/ivpusic/neo) **star:396** Neo is minimal and fast Go Web Framework with extremely simple API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/neo)
* [Air](https://github.com/aofei/air) **star:341** An ideally refined web framework for Go.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/air)
* [mango](https://github.com/paulbellamy/mango) **star:339** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/paulbellamy/mango)
* [Gondola](https://github.com/rainycape/gondola) **star:314** The web framework for writing faster sites, faster.   [![godoc][GoDoc]](https://godoc.org/github.com/rainycape/gondola)
* [Golf](https://github.com/dinever/golf) **star:238** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dinever/golf)
* [Aero](https://github.com/aerogo/aero) **star:175** High-performance web framework for Go, reaches top scores in Lighthouse.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/aero)
* [Gem](https://github.com/go-gem/gem) **star:153** Simple and fast web framework, friendly to REST API.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-gem/gem)
* [go-rest](https://github.com/ungerik/go-rest) **star:116** Small and evil REST framework for Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-rest)
* [hiboot](https://github.com/hidevopsio/hiboot) **star:93** hiboot is a high performance web application framework with auto configuration and dependency injection support.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/hidevopsio/hiboot)   ![Contains Chinese documents][CN]
* [WebGo](https://github.com/bnkamalesh/webgo) **star:78** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/webgo)
* [Golax](https://github.com/fulldump/golax) **star:72** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/golax)
* [Microservice](https://github.com/claygod/microservice) **star:60** The framework for the creation of microservices, written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/microservice)
* [uAdmin](https://github.com/uadmin/uadmin) **star:58** Fully featured web framework for Golang, inspired by Django.   [![godoc][GoDoc]](https://godoc.org/github.com/uadmin/uadmin)
* [YARF](https://github.com/yarf-framework/yarf) **star:51** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.   [![godoc][GoDoc]](https://godoc.org/github.com/yarf-framework/yarf)
* [Fireball](https://github.com/zpatrick/fireball) **star:49** More "natural" feeling web framework.   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/fireball)
* [vox](https://github.com/aisk/vox) **star:42** A golang web framework for humans, inspired by Koa heavily.   [![godoc][GoDoc]](https://godoc.org/github.com/aisk/vox)
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [patron](https://github.com/beatlabs/patron) **star:36** Patron is a microservice framework following best cloud practices with a focus on productivity.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/patron)
* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [rux](https://github.com/gookit/rux) **star:14** Simple and fast web framework for build golang HTTP applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/rux)   ![Contains Chinese documents][CN]
* [route](https://github.com/goroute/route) **star:4** Simple yet powerful HTTP request multiplexer.   [![godoc][GoDoc]](https://godoc.org/github.com/goroute/route)

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) **star:1301** Rate limit HTTP request handler.   [![godoc][GoDoc]](https://godoc.org/github.com/didip/tollbooth)
* [CORS](https://github.com/rs/cors) **star:1249** Easily add CORS capabilities to your API.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/cors)
* [Limiter](https://github.com/ulule/limiter) **star:806** Dead simple rate limit middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/limiter)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:753** Add/parse Server-Timing header.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/go-server-timing)
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:92** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/ln-paywall)
* [XFF](https://github.com/sebest/xff) **star:72** Handle `X-Forwarded-For` header and friends.   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/xff)
* [client-timing](https://github.com/posener/client-timing) **star:15** An HTTP client for Server-Timing header.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/client-timing)

#### Libraries for creating HTTP middlewares

* [negroni](https://github.com/urfave/negroni) **star:6382** Idiomatic HTTP middleware for Golang.   [![star > 2000][Awesome]](https://github.com/urfave/negroni)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/negroni)   ![Contains Chinese documents][CN]
* [alice](https://github.com/justinas/alice) **star:1849** Painless middleware chaining for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/alice)
* [render](https://github.com/unrolled/render) **star:1286** Go package for easily rendering JSON, XML, and HTML template responses.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/render)
* [stats](https://github.com/thoas/stats) **star:540** Go middleware that stores various information about your web application.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/stats)
* [interpose](https://github.com/carbocation/interpose) **star:288** Minimalist net/http middleware for golang.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/carbocation/interpose)
* [muxchain](https://github.com/stephens2424/muxchain) **star:208** Lightweight middleware for net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/stephens2424/muxchain)
* [renderer](https://github.com/thedevsaddam/renderer) **star:170** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/renderer)
* [rye](https://github.com/InVisionApp/rye) **star:94** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/rye)
* [gores](https://github.com/alioygur/gores) **star:83** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/alioygur/gores)
* [chain](https://github.com/codemodus/chain) **star:64** Handler wrapper chaining with scoped data (net/context-based "middleware").   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/chain)
* [go-wrap](https://github.com/go-on/wrap) **star:59** Small middlewares package for net/http.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/go-on/wrap)

### Routers

* [mux](https://github.com/gorilla/mux) **star:10005** Powerful URL router and dispatcher for golang.   [![star > 2000][Awesome]](https://github.com/gorilla/mux)   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/mux)
* [httprouter](https://github.com/julienschmidt/httprouter) **star:9908** High performance router. Use this and the standard http handlers to form a very high performance web framework.   [![star > 2000][Awesome]](https://github.com/julienschmidt/httprouter)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/httprouter)
* [chi](https://github.com/go-chi/chi) **star:6286** Small, fast and expressive HTTP router built on net/context.   [![star > 2000][Awesome]](https://github.com/go-chi/chi)   [![godoc][GoDoc]](https://godoc.org/github.com/go-chi/chi)
* [gocraft/web](https://github.com/gocraft/web) **star:1397** Mux and middleware package in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gocraft/web)
* [Bone](https://github.com/go-zoo/bone) **star:1225** Lightning Fast HTTP Multiplexer.   [![godoc][GoDoc]](https://godoc.org/github.com/go-zoo/bone)
* [Goji](https://github.com/goji/goji) **star:770** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/goji/goji)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:759** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   [![godoc][GoDoc]](https://godoc.org/github.com/buaazp/fasthttprouter)
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:454** A simple and fast HTTP router for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gorouter)
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:392** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   [![godoc][GoDoc]](https://godoc.org/github.com/dimfeld/httptreemux)
* [lars](https://github.com/go-playground/lars) **star:376** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/lars)
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:361** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
* [Siesta](https://github.com/VividCortex/siesta) **star:351** Composable framework to write middleware and handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/siesta)
* [vestigo](https://github.com/husobee/vestigo) **star:250** Performant, stand-alone, HTTP compliant URL Router for go web applications.   [![godoc][GoDoc]](https://godoc.org/github.com/husobee/vestigo)
* [gowww/router](https://github.com/gowww/router) **star:158** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gowww/router)
* [alien](https://github.com/gernest/alien) **star:107** Lightweight and fast http router from outer space.   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/alien)
* [violetear](https://github.com/nbari/violetear) **star:95** Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/nbari/violetear)
* [Bxog](https://github.com/claygod/Bxog) **star:93** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/Bxog)
* [xmux](https://github.com/rs/xmux) **star:88** High performance muxer based on `httprouter` with `net/context` support.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xmux)
* [pure](https://github.com/go-playground/pure) **star:87** Is a lightweight HTTP router that sticks to the std "net/http" implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pure)
* [GoRouter](https://github.com/vardius/gorouter) **star:51** GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gorouter)
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:39** A simple Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/bellt)

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:561** Win32 OLE implementation for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ole/go-ole)
* [d3d9](https://github.com/gonutz/d3d9) **star:88** Go bindings for Direct3D9.   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/d3d9)
* [gosddl](https://github.com/MonaxGT/gosddl) **star:1** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gosddl)

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) **star:270** Generate a Go struct from XML.   [![godoc][GoDoc]](https://godoc.org/github.com/miku/zek)
* [xpath](https://github.com/antchfx/xpath) **star:180** XPath package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xpath)
* [xquery](https://github.com/antchfx/xquery) **star:147** XQuery lets you extract data from HTML/XML documents using XPath expression.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xquery)
* [xml2map](https://github.com/sbabiv/xml2map) **star:16** XML to MAP converter written Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/xml2map)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:6** Procedural XML generation API based on libxml2's xmlwriter module.   [![godoc][GoDoc]](https://godoc.org/github.com/shabbyrobe/xmlwriter)

# Tools

*Go software and plugins.*

## Code Analysis

* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [GoLint](https://github.com/golang/lint) **star:3202** Golint is a linter for Go source code.   [![star > 2000][Awesome]](https://github.com/golang/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/lint)
* [errcheck](https://github.com/kisielk/errcheck) **star:1335** Errcheck is a program for checking for unchecked errors in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/kisielk/errcheck)
* [gcvis](https://github.com/davecheney/gcvis) **star:923** Visualise Go program GC trace data in real time.   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/gcvis)
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [php-parser](https://github.com/z7zmey/php-parser) **star:661** A Parser for PHP written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/z7zmey/php-parser)
* [go-critic](https://github.com/go-critic/go-critic) **star:598** source code linter that brings checks that are currently not implemented in other linters.   [![godoc][GoDoc]](https://godoc.org/github.com/go-critic/go-critic)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:379** Web based Golang AST visualizer.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:287** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
* [unconvert](https://github.com/mdempsky/unconvert) **star:261** Remove unnecessary type conversions from Go source.   [![godoc][GoDoc]](https://godoc.org/github.com/mdempsky/unconvert)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [gostatus](https://github.com/shurcooL/gostatus) **star:241** Command line tool, shows the status of repositories that contain Go packages.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/gostatus)
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:200** An easy way to find outdated dependencies of your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/go-mod-outdated)
* [dupl](https://github.com/mibk/dupl) **star:175** Tool for code clone detection.   [![godoc][GoDoc]](https://godoc.org/github.com/mibk/dupl)
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:166** Checks recent changes to a Go project for backwards incompatible changes.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyfalzon/apicompat)
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:95** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/checkstyle)
* [lint](https://github.com/surullabs/lint) **star:63** Run linters as part of go test.   [![godoc][GoDoc]](https://godoc.org/github.com/surullabs/lint)
* [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/validate)
* [go-outdated](https://github.com/firstrow/go-outdated) **star:45** Console application that displays outdated packages.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/go-outdated)   ![Archived][Archived]

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) **star:11035** Go development plugin for Vim.   [![star > 2000][Awesome]](https://github.com/fatih/vim-go)   ![There was an update last week][Green]
* [vscode-go](https://github.com/Microsoft/vscode-go) **star:5230** Extension for Visual Studio Code (VS Code) which provides support for the Go language.   [![star > 2000][Awesome]](https://github.com/Microsoft/vscode-go)   ![There was an update last week][Green]
* [gocode](https://github.com/nsf/gocode) **star:4764** Autocompletion daemon for the Go programming language.   [![star > 2000][Awesome]](https://github.com/nsf/gocode)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/gocode)
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3251** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.   [![star > 2000][Awesome]](https://github.com/DisposaBoy/GoSublime)   [![godoc][GoDoc]](https://godoc.org/github.com/DisposaBoy/GoSublime)
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1485** Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.   ![There was an update last week][Green]
* [go-mode](https://github.com/dominikh/go-mode.el) **star:975** Go mode for GNU/Emacs.   ![There was an update last week][Green]
* [Watch](https://github.com/eaburns/Watch) **star:170** Runs a command in an acme win on file changes.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/eaburns/Watch)
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:81** Vim plugin to highlight syntax errors on save.   ![It hasn't been updated in the last year][Yellow]
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:31** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:17** Vim plugin for generating Go tests based on the function's or method's signature.
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:12** Go language support for the Theia IDE.

## Go Generate Tools

* [gotests](https://github.com/cweill/gotests) **star:2275** Generate Go tests from your source code.   [![star > 2000][Awesome]](https://github.com/cweill/gotests)   [![godoc][GoDoc]](https://godoc.org/github.com/cweill/gotests)
* [genny](https://github.com/cheekybits/genny) **star:1012** Elegant generics for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cheekybits/genny)
* [re2dfa](https://github.com/opennota/re2dfa) **star:170** Transform regular expressions into finite state machines and output Go source code.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/opennota/re2dfa)
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.
* [gocontracts](https://github.com/Parquery/gocontracts) **star:53** brings design-by-contract to Go by synchronizing the code with the documentation.   [![godoc][GoDoc]](https://godoc.org/github.com/Parquery/gocontracts)
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [generic](https://github.com/usk81/generic) **star:30** flexible data type for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/usk81/generic)
* [gounit](https://github.com/hexdigest/gounit) **star:30** Generate Go tests using your own templates.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gounit)
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:17** Generate Haskell inspired functions for your slices.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/hasgo)

## Go Tools

* [go-swagger](https://github.com/go-swagger/go-swagger) **star:4138** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.   [![star > 2000][Awesome]](https://github.com/go-swagger/go-swagger)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-swagger/go-swagger)
* [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:3872** Navigate through go files efficiently with the OctoLinker browser extension for GitHub.   [![star > 2000][Awesome]](https://github.com/OctoLinker/browser-extension)   ![There was an update last week][Green]
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:2046** Visualize call graph of your Go program using dot format.   [![star > 2000][Awesome]](https://github.com/TrueFurby/go-callvis)   [![godoc][GoDoc]](https://godoc.org/github.com/TrueFurby/go-callvis)
* [richgo](https://github.com/kyoh86/richgo) **star:398** Enrich `go test` outputs with text decorations.   [![godoc][GoDoc]](https://godoc.org/github.com/kyoh86/richgo)
* [depth](https://github.com/KyleBanks/depth) **star:393** Visualize dependency trees of any package by analyzing imports.   [![godoc][GoDoc]](https://godoc.org/github.com/KyleBanks/depth)
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [rts](https://github.com/galeone/rts) **star:187** RTS: response to struct. Generates Go structs from server responses.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/rts)
* [godbg](https://github.com/tylerwince/godbg) **star:157** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.   [![godoc][GoDoc]](https://godoc.org/github.com/tylerwince/godbg)
* [colorgo](https://github.com/songgao/colorgo) **star:100** Wrapper around `go` command for colorized `go build` output.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/colorgo)
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:38** Bash completion for go and wgo.   ![It hasn't been updated in the last year][Yellow]
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:14** A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:57838** Container Cluster Manager from Google.   [![star > 2000][Awesome]](https://github.com/kubernetes/kubernetes)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby) **star:54915** Collaborative project for the container ecosystem to assemble container-based systems.   [![star > 2000][Awesome]](https://github.com/moby/moby)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/moby/moby)
* [traefik](https://github.com/containous/traefik) **star:24421** Reverse proxy and load balancer with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/containous/traefik)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/containous/traefik)
* [Gitea](https://github.com/go-gitea/gitea) **star:15898** Fork of Gogs, entirely community driven.   [![star > 2000][Awesome]](https://github.com/go-gitea/gitea)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/go-gitea/gitea)   ![Contains Chinese documents][CN]
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [Vegeta](https://github.com/tsenart/vegeta) **star:12509** HTTP load testing tool and library. It's over 9000!   [![star > 2000][Awesome]](https://github.com/tsenart/vegeta)   [![godoc][GoDoc]](https://godoc.org/github.com/tsenart/vegeta)
* [Packer](https://github.com/mitchellh/packer) **star:9355** Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.   [![star > 2000][Awesome]](https://github.com/mitchellh/packer)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/packer)
* [Hey](https://github.com/rakyll/hey) **star:6583** Hey is a tiny program that sends some load to a web application.   [![star > 2000][Awesome]](https://github.com/rakyll/hey)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/hey)
* [GVM](https://github.com/moovweb/gvm) **star:4556** GVM provides an interface to manage Go versions.   [![star > 2000][Awesome]](https://github.com/moovweb/gvm)
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [webhook](https://github.com/adnanh/webhook) **star:4209** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   [![star > 2000][Awesome]](https://github.com/adnanh/webhook)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/adnanh/webhook)
* [gaia](https://github.com/gaia-pipeline/gaia) **star:3784** Build powerful pipelines in any programming language.   [![star > 2000][Awesome]](https://github.com/gaia-pipeline/gaia)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/gaia-pipeline/gaia)
* [gox](https://github.com/mitchellh/gox) **star:3399** Dead simple, no frills Go cross compile tool.   [![star > 2000][Awesome]](https://github.com/mitchellh/gox)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/gox)
* [bosun](https://github.com/bosun-monitor/bosun) **star:2864** Time Series Alerting Framework.   [![star > 2000][Awesome]](https://github.com/bosun-monitor/bosun)   [![godoc][GoDoc]](https://godoc.org/github.com/bosun-monitor/bosun)
* [bombardier](https://github.com/codesenberg/bombardier) **star:1782** Fast cross-platform HTTP benchmarking tool.   [![godoc][GoDoc]](https://godoc.org/github.com/codesenberg/bombardier)
* [goxc](https://github.com/laher/goxc) **star:1628** build tool for Go, with a focus on cross-compiling and packaging.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/laher/goxc)
* [fac](https://github.com/mkchoi212/fac) **star:1614** Command-line user interface to fix git merge conflicts.   [![godoc][GoDoc]](https://godoc.org/github.com/mkchoi212/fac)
* [kala](https://github.com/ajvb/kala) **star:1369** Simplistic, modern, and performant job scheduler.   [![godoc][GoDoc]](https://godoc.org/github.com/ajvb/kala)
* [StatusOK](https://github.com/sanathp/statusok) **star:1182** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.   [![godoc][GoDoc]](https://godoc.org/github.com/sanathp/statusok)
* [script](https://github.com/bitfield/script) **star:1024** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/script)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:1009** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.   [![godoc][GoDoc]](https://godoc.org/github.com/rlmcpherson/s3gof3r)
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:679** Enable your Go applications to self update.   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/go-selfupdate)
* [Pomerium](https://github.com/pomerium/pomerium) **star:594** Pomerium is an identity-aware access proxy.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/pomerium/pomerium)
* [skm](https://github.com/TimothyYe/skm) **star:551** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/TimothyYe/skm)
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:541** Manage BareMetal Servers from Command Line (as easily as with Docker).   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/scaleway/scaleway-cli)
* [aurora](https://github.com/xuri/aurora) **star:414** Cross-platform web-based Beanstalkd queue server console.
* [govvv](https://github.com/ahmetalpbalkan/govvv) **star:395** “go build” wrapper to easily add version information into Go binaries.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/govvv)
* [gonative](https://github.com/inconshreveable/gonative) **star:311** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/gonative)
* [Mora](https://github.com/emicklei/mora) **star:267** REST server for accessing MongoDB documents and meta data.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/emicklei/mora)
* [lstags](https://github.com/ivanilves/lstags) **star:227** Tool and API to sync Docker images across different registries.   [![godoc][GoDoc]](https://godoc.org/github.com/ivanilves/lstags)
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [godbg](https://github.com/sirnewton01/godbg) **star:219** Web-based gdb front-end application.   ![It hasn't been updated in the last year][Yellow]
* [dogo](https://github.com/liudng/dogo) **star:214** Monitoring changes in the source file and automatically compile and run (restart).   [![godoc][GoDoc]](https://godoc.org/github.com/liudng/dogo)   ![Contains Chinese documents][CN]
* [manssh](https://github.com/xwjdsh/manssh) **star:205** manssh is a command line tool for managing your ssh alias config easily.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/xwjdsh/manssh)
* [Pewpew](https://github.com/bengadbois/pewpew) **star:205** Flexible HTTP command line stress tester.   [![godoc][GoDoc]](https://godoc.org/github.com/bengadbois/pewpew)
* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [gobrew](https://github.com/cryptojuice/gobrew) **star:176** gobrew lets you easily switch between multiple versions of go.   ![It hasn't been updated in the last year][Yellow]
* [Blast](https://github.com/dave/blast) **star:170** A simple tool for API load testing and batch jobs.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/dave/blast)
* [ostent](https://github.com/ostrost/ostent) **star:165** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/ostrost/ostent)
* [grapes](https://github.com/yaronsumel/grapes) **star:138** Lightweight tool designed to distribute commands over ssh with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/grapes)
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:103** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/easyssh-proxy)
* [kcli](https://github.com/cswank/kcli) **star:82** Command line tool for inspecting kafka topics/partitions/messages.   [![godoc][GoDoc]](https://godoc.org/github.com/cswank/kcli)
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:71** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   [![godoc][GoDoc]](https://godoc.org/github.com/go-furnace/go-furnace)
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:69** Cli tool to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm-cli)
* [drone-scp](https://github.com/appleboy/drone-scp) **star:58** Copy files and artifacts via SSH using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-scp)
* [Dropship](https://github.com/chrismckenzie/dropship) **star:46** Tool for deploying code via cdn.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/chrismckenzie/dropship)
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:25** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-jenkins)
* [DepCharge](https://github.com/centerorbit/depcharge) **star:9** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![godoc][GoDoc]](https://godoc.org/github.com/centerorbit/depcharge)

### Other Software

* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [Gor](https://github.com/buger/gor) **star:11529** Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.   [![star > 2000][Awesome]](https://github.com/buger/gor)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/buger/gor)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:8440** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   [![star > 2000][Awesome]](https://github.com/chrislusf/seaweedfs)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/seaweedfs)
* [restic](https://github.com/restic/restic) **star:7698** De-duplicating backup program.   [![star > 2000][Awesome]](https://github.com/restic/restic)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/restic/restic)
* [confd](https://github.com/kelseyhightower/confd) **star:6512** Manage local application configuration files using templates and data from etcd or consul.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/confd)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/confd)
* [Comcast](https://github.com/tylertreat/Comcast) **star:6219** Simulate bad network connections.   [![star > 2000][Awesome]](https://github.com/tylertreat/Comcast)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/Comcast)
* [LiteIDE](https://github.com/visualfc/liteide) **star:5587** LiteIDE is a simple, open source, cross-platform Go IDE.   [![star > 2000][Awesome]](https://github.com/visualfc/liteide)   ![There was an update last week][Green]   ![Contains Chinese documents][CN]
* [drive](https://github.com/odeke-em/drive) **star:5032** Google Drive client for the commandline.   [![star > 2000][Awesome]](https://github.com/odeke-em/drive)   [![godoc][GoDoc]](https://godoc.org/github.com/odeke-em/drive)
* [nes](https://github.com/fogleman/nes) **star:4204** Nintendo Entertainment System (NES) emulator written in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/nes)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/nes)
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:4048** Proxy to simulate network and system conditions for automated tests.   [![star > 2000][Awesome]](https://github.com/shopify/toxiproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/shopify/toxiproxy)
* [Pipe](https://github.com/b3log/pipe) **star:3348** A small and beautiful blogging platform.   [![star > 2000][Awesome]](https://github.com/b3log/pipe)   [![godoc][GoDoc]](https://godoc.org/github.com/b3log/pipe)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:2717** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   [![star > 2000][Awesome]](https://github.com/gilbertchen/duplicacy)   [![godoc][GoDoc]](https://godoc.org/github.com/gilbertchen/duplicacy)
* [myLG](https://github.com/mehrdadrad/mylg) **star:2216** Command Line Network Diagnostic tool written in Go.   [![star > 2000][Awesome]](https://github.com/mehrdadrad/mylg)   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/mylg)
* [GoBoy](https://github.com/Humpheh/goboy) **star:2130** Nintendo Game Boy Color emulator written in Go.   [![star > 2000][Awesome]](https://github.com/Humpheh/goboy)   [![godoc][GoDoc]](https://godoc.org/github.com/Humpheh/goboy)
* [Stack Up](https://github.com/pressly/sup) **star:2011** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   [![star > 2000][Awesome]](https://github.com/pressly/sup)   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/sup)
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [lgo](https://github.com/yunabe/lgo) **star:1828** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/yunabe/lgo)
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [snap](https://github.com/intelsdi-x/snap) **star:1804** Powerful telemetry framework.   [![godoc][GoDoc]](https://godoc.org/github.com/intelsdi-x/snap)
* [Circuit](https://github.com/gocircuit/circuit) **star:1795** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gocircuit/circuit)
* [borg](https://github.com/crufter/borg) **star:1425** Terminal based search engine for bash snippets.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/crufter/borg)
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:877** App that displays updates for the Go packages in your GOPATH.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/Go-Package-Store)
* [Documize](https://github.com/documize/community) **star:864** Modern wiki software that integrates data from SaaS tools.   ![There was an update last week][Green]
* [scc](https://github.com/boyter/scc) **star:799** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/boyter/scc)
* [Leaps](https://github.com/jeffail/leaps) **star:643** Pair programming service using Operational Transforms.   [![godoc][GoDoc]](https://godoc.org/github.com/jeffail/leaps)
* [peg](https://github.com/pointlander/peg) **star:623** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   [![godoc][GoDoc]](https://godoc.org/github.com/pointlander/peg)
* [vFlow](https://github.com/VerizonDigital/vflow) **star:609** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   [![godoc][GoDoc]](https://godoc.org/github.com/VerizonDigital/vflow)
* [gfile](https://github.com/Antonito/gfile) **star:505** Securely transfer files between two computers, without any third party, over WebRTC.   [![godoc][GoDoc]](https://godoc.org/github.com/Antonito/gfile)
* [GoDNS](https://github.com/timothyye/godns) **star:453** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [shell2http](https://github.com/msoap/shell2http) **star:439** Executing shell commands via http server (for prototyping or remote control).   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/shell2http)
* [mockingjay](https://github.com/quii/mockingjay-server) **star:422** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   [![godoc][GoDoc]](https://godoc.org/github.com/quii/mockingjay-server)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:377** Video streaming torrent client.   [![godoc][GoDoc]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
* [gocc](https://github.com/goccmack/gocc) **star:349** Gocc is a compiler kit for Go written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/goccmack/gocc)
* [wellington](https://github.com/wellington/wellington) **star:291** Sass project management tool, extends the language with sprite functions (like Compass).   [![godoc][GoDoc]](https://godoc.org/github.com/wellington/wellington)
* [ipe](https://github.com/dimiro1/ipe) **star:282** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/ipe)
* [ide](https://github.com/thestrukture/ide) **star:254** Browser accessible IDE. Designed for Go with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thestrukture/ide)
* [Cherry](https://github.com/rafael-santiago/cherry) **star:194** Tiny webchat server in Go.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/rafael-santiago/cherry)
* [orange-cat](https://github.com/noraesae/orange-cat) **star:178** Markdown previewer written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/noraesae/orange-cat)
* [Orbit](https://github.com/gulien/orbit) **star:130** A simple tool for running commands and generating files from templates.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/gulien/orbit)
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [joincap](https://github.com/assafmo/joincap) **star:123** Command-line utility for merging multiple pcap files together.   [![godoc][GoDoc]](https://godoc.org/github.com/assafmo/joincap)
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [DDNS](https://github.com/skibish/ddns) **star:102** Personal DDNS client with Digital Ocean Networking DNS as backend.   [![godoc][GoDoc]](https://godoc.org/github.com/skibish/ddns)
* [boxed](https://github.com/tejo/boxed) **star:72** Dropbox based blog engine.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tejo/boxed)
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** Simple NaCL EC25519 based crypto pipe tool written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/unix4fun/naclpipe)
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/crazcalm/term-quiz)

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1268** Go HTTP request router benchmark and comparison.   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1024** Go web framework benchmark.   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
* [skynet](https://github.com/atemerev/skynet) **star:926** Skynet 1M threads microbenchmark.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:885** Benchmarks of Go serialization methods.   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:173** Compare various Image resize algorithms for the Go language.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/fawick/speedtest-resize)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:125** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/go-benchmarks)
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:95** Go micro-benchmarks for calculating the speed of language constructs.   [![godoc][GoDoc]](https://godoc.org/github.com/feyeleanor/GoSpeed)
* [autobench](https://github.com/davecheney/autobench) **star:89** Framework to compare the performance between different Go versions.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/autobench)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) **star:52** Benchmarks of common basic operations for the Go language.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/gocostmodel)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:49** Collection of benchmarks for popular Go database/SQL utilities.   ![It hasn't been updated in the last year][Yellow]   [![godoc][GoDoc]](https://godoc.org/github.com/tyler-smith/golang-sql-benchmark)

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
* [GoBooks](https://github.com/dariubs/GoBooks) **star:6926** A curated list of Go books.   [![star > 2000][Awesome]](https://github.com/dariubs/GoBooks)
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:10** in Persian.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsir/gosuccinctly)
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1604** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   [![godoc][GoDoc]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:67** adorable gopher logos.   ![It hasn't been updated in the last year][Yellow]
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:1880** Gopher artworks by Ashley McNamara.   [![godoc][GoDoc]](https://godoc.org/github.com/ashleymcnamara/gophers)
* [gophers](https://github.com/egonelbre/gophers) **star:1644** Free gophers.   [![godoc][GoDoc]](https://godoc.org/github.com/egonelbre/gophers)
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:328** Gopherize yourself.
* [gophers](https://github.com/rogeralsing/gophers) **star:50** random gopher graphics.   ![It hasn't been updated in the last year][Yellow]
* [gophers](https://github.com/sillecelik/go-gopher) **star:41** Gopher amigurumi toy pattern.

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

* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:24870** List of other amazingly awesome lists.   [![star > 2000][Awesome]](https://github.com/bayandin/awesome-awesomeness)
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:14868** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.   [![star > 2000][Awesome]](https://github.com/lukasz-madon/awesome-remote-job)   ![There was an update last week][Green]
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:143** Collection of Go images, graphics, and art.   ![It hasn't been updated in the last year][Yellow]   ![Archived][Archived]
* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:32124** Golang ebook intro how to build a web app with golang.   [![star > 2000][Awesome]](https://github.com/astaxie/build-web-application-with-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   ![Contains Chinese documents][CN]
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:4102** Go's reference card.   [![star > 2000][Awesome]](https://github.com/a8m/go-lang-cheat-sheet)
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:471** A little e-book on Ethereum Development with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/ethereum-development-with-go-book)   ![Contains Chinese documents][CN]
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:7908** Learn Go with test-driven development.   [![star > 2000][Awesome]](https://github.com/quii/learn-go-with-tests)   ![There was an update last week][Green]   [![godoc][GoDoc]](https://godoc.org/github.com/quii/learn-go-with-tests)   ![Contains Chinese documents][CN]
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:728** Examples of Golang compared to Node.js for learning.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1138** Intro to go for experienced programmers.
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

