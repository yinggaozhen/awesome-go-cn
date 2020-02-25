# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2020-02-25 12:00:12(Synchronize every day)**

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
    - [Dynamic DNS](#dynamic-dns)
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
    - [Server Applications](#server-applications)
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

* [Oto](https://github.com/hajimehoshi/oto) **star:537** A low-level library to play sound on multiple platforms.   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:335** Go bindings for the PortAudio audio I/O library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gordonklaus/portaudio)   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:283** Music theory models in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:274** Go package capable of generating waveform images from audio streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/waveform)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:225** Go bindings for PortMidi.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:136** Fast and stable ID3 parsing and writing library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [mix](https://github.com/go-mix/mix) **star:112** Sequence-based Go-native audio mixer for music apps.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:105** Native Go MP3 decoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:104** libsox bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [malgo](https://github.com/gen2brain/malgo) **star:95** Mini audio library.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/malgo)
* [taglib](https://github.com/wtolson/go-taglib) **star:72** Go bindings for taglib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wtolson/go-taglib)   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:67** Native Go AAC bitstream parser.   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:37** Lightweight MP3 decoder library.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:28** libmediainfo bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/go_mediainfo)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/go_mediainfo)
* [flac](https://github.com/mewkiz/flac)  Native Go FLAC encoder/decoder with support for FLAC streams.
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:26** EasyMidi is a simple and reliable library for working with standard midi file (SMF).   [![It hasn't been updated in the last year][Yellow]](https://github.com/algoGuy/EasyMIDI)   [![godoc][GoDoc]](https://godoc.org/github.com/algoGuy/EasyMIDI)
* [vorbis](https://github.com/mccoyst/vorbis) **star:25** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:9** libsamplerate bindings for go.   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:7098** Golang implementation of JSON Web Tokens (JWT).   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![There was an update last week][Green]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:6142** Authorization library that supports access control models like ACL, RBAC, ABAC.   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![There was an update last week][Green]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [osin](https://github.com/openshift/osin)  Golang OAuth2 server library.
* [oauth2](https://github.com/golang/oauth2) **star:2739** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2566** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![There was an update last week][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:2154** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   [![star > 2000][Awesome]](https://github.com/volatiletech/authboss)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [branca](https://github.com/hako/branca)  Golang implementation of Branca Tokens.
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1446** Standalone, specification-compliant,  OAuth2 server written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [gologin](https://github.com/dghubble/gologin) **star:1162** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:986** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [loginsrv](https://github.com/tarent/loginsrv) **star:910** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   [![There was an update last week][Green]](https://github.com/tarent/loginsrv)   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [scs](https://github.com/alexedwards/scs) **star:675** Session Manager for HTTP servers.   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [permissions2](https://github.com/xyproto/permissions2) **star:391** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   [![There was an update last week][Green]](https://github.com/xyproto/permissions2)   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [paseto](https://github.com/o1egl/paseto) **star:315** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [httpauth](https://github.com/goji/httpauth) **star:190** HTTP Authentication middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goji/httpauth)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:189** Simple, flexible, secure and idiomatic web session management with pluggable backends.   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:169** JWT middleware for Golang http servers with many configuration options.   [![It hasn't been updated in the last year][Yellow]](https://github.com/adam-hanna/jwt-auth)   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/pascaldekloe/jwt) **star:153** Lightweight JSON Web Token (JWT) library.   [![There was an update last week][Green]](https://github.com/pascaldekloe/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [session](https://github.com/icza/session) **star:98** Go session management for web servers (including support for Google App Engine - GAE).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [jwt](https://github.com/robbert229/jwt) **star:78** Clean and easy to use implementation of JSON Web Tokens (JWT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/robbert229/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sessionup](https://github.com/swithek/sessionup) **star:77** Simple, yet effective HTTP session management and identification package.   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [sjwt](https://github.com/brianvoe/sjwt) **star:58** Simple jwt generator and parser.   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [sessions](https://github.com/adam-hanna/sessions) **star:48** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [rbac](https://github.com/zpatrick/rbac) **star:46** Minimalistic RBAC package for Go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/rbac)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [securecookie](https://github.com/chmike/securecookie) **star:33** Efficient secure cookie encoding/decoding.   [![There was an update last week][Green]](https://github.com/chmike/securecookie)   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.   [![It hasn't been updated in the last year][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [signedvalue](https://github.com/sashka/signedvalue) **star:8** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)
* [scope](https://github.com/SonicRoshan/scope) **star:4** Easily Manage OAuth2 Scopes In Go.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** provides parser of cookies.txt file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)
* [go-jose](https://github.com/square/go-jose)  Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api)  Simple and clean Telegram bot client.
* [telebot](https://github.com/tucnak/telebot) **star:1125** Telegram bot framework written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:548** IRC, Slack & Telegram bot written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [go-joe](https://joe-bot.net)  A general-purpose bot library inspired by Hubot but written in Go.
* [go-sarah](https://github.com/oklahomer/go-sarah)  Framework to build bot for desired chat services including LINE, Slack, Gitter and more.
* [slacker](https://github.com/shomali11/slacker) **star:356** Easy to use framework to create Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:287** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   [![It hasn't been updated in the last year][Yellow]](https://github.com/saniales/golang-crypto-trading-bot)   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [Kelp](https://github.com/stellar/kelp) **star:265** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   [![There was an update last week][Green]](https://github.com/stellar/kelp)   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [tbot](https://github.com/yanzay/tbot) **star:248** Telegram bot server with API similar to net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:168** Service oriented IRC bot using Redis and JSON for messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [hanu](https://github.com/sbstjn/hanu) **star:118** Framework for writing Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:93** Libary to write bots for twitch.tv chat   [![godoc][GoDoc]](https://godoc.org/github.com/gempir/go-twitch-irc)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:91** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/olebedev/go-tgbot)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:60** Framework for building Telegram bots.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:30** Simple Go [VK](https://vk.com) bot library.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:24** Another framework for building Slack bots.   [![There was an update last week][Green]](https://github.com/alexandre-normand/slackscot)   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [micha](https://github.com/onrik/micha) **star:12** Go Library for Telegram bot api.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:15875** Commander for modern Go CLI interactions.   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![There was an update last week][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:13159** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2811** Command line and flag parser supporting sub commands.   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![There was an update last week][Green]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1636** go command line option parser.   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [Dnote](https://github.com/dnote/dnote) **star:1498** A simple command line notebook with multi-device sync.   [![There was an update last week][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [readline](https://github.com/chzyer/readline) **star:1463** Pure golang implementation that provides most features in GNU-Readline under MIT license.   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [docopt.go](https://github.com/docopt/docopt.go) **star:1223** Command-line arguments parser that will make you smile.   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1092** Go library for implementing command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [pflag](https://github.com/spf13/pflag) **star:997** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   [![There was an update last week][Green]](https://github.com/spf13/pflag)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [cli-init](https://github.com/tcnksm/gcli) **star:893** The easy way to start building Golang command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcnksm/gcli)   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [go-arg](https://github.com/alexflint/go-arg) **star:835** Struct-based argument parsing in Go.   [![There was an update last week][Green]](https://github.com/alexflint/go-arg)   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:665** Write bash completions in Go + Go command bash completion.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [mow.cli](https://github.com/jawher/mow.cli) **star:656** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [liner](https://github.com/peterh/liner) **star:652** Go readline-like library for command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [flaggy](https://github.com/integrii/flaggy) **star:626** A robust and idiomatic flags package with excellent subcommand support.   [![There was an update last week][Green]](https://github.com/integrii/flaggy)   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:513** Feature-rich and easy to use command-line package based on golang struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:380** Unikernel Builder/Orchestrator.   [![There was an update last week][Green]](https://github.com/nanovms/ops)   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:164** Command line argument parser inspired by Python's argparse module.   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [commandeer](https://github.com/jaffee/commandeer) **star:107** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [sflags](https://github.com/octago/sflags) **star:104** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [flag](https://github.com/cosiner/flag) **star:103** Simple but powerful command line option parsing library for Go supporting subcommand.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:101** Small command line interface framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ukautz/clif)   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [wmenu](https://github.com/dixonwille/wmenu) **star:98** Easy to use menu structure for cli applications that prompts users to make choices.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [job](https://github.com/liujianping/job) **star:68** JOB, make your short-term command as a long-term job.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [cli](https://github.com/teris-io/cli) **star:64** Simple and complete API for building command line interfaces in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [1build](https://github.com/gopinath-langote/1build) **star:49** Command line tool to frictionlessly manage project-specific commands.   [![There was an update last week][Green]](https://github.com/gopinath-langote/1build)   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [env](https://github.com/codingconcepts/env) **star:48** Tag-based environment configuration for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [wlog](https://github.com/dixonwille/wlog) **star:40** Simple logging interface that supports cross-platform color and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wlog)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [gocmd](https://github.com/devfacet/gocmd) **star:37** Go library for building command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [strumt](https://github.com/antham/strumt) **star:34** Library to create prompt chain.   [![There was an update last week][Green]](https://github.com/antham/strumt)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [ts](https://github.com/liujianping/ts)  Timestamp convert & compare tool.
* [flagvar](https://github.com/sgreben/flagvar) **star:33** A collection of flag argument types for Go's standard `flag` package.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [cmdr](https://github.com/hedzr/cmdr) **star:30** A POSIX/GNU style, getopt-like command-line UI Go library.   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:19** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [argv](https://github.com/cosiner/argv) **star:18** Go library to split command line string as arguments array using the bash syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [clîr](https://github.com/leaanthony/clir) **star:17** A Simple and Clear CLI library. Dependency free.   [![There was an update last week][Green]](https://github.com/leaanthony/clir)   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [sand](https://github.com/Zaba505/sand) **star:9** Simple API for creating interpreters and so much more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Zaba505/sand)   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [cmd](https://github.com/posener/cmd) **star:3** Extends the standard `flag` package to support sub commands and more in idomatic way.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:9532** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![There was an update last week][Green]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [gocui](https://github.com/jroimartin/gocui) **star:5985** Minimalist Go library aimed at creating Console User Interfaces.   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2842** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   [![There was an update last week][Green]](https://github.com/c-bata/go-prompt)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1640** Flexible library to render progress bars in terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1308** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:1094** Library for updating terminal output in realtime.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [termdash](https://github.com/mum4k/termdash) **star:944** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [progressbar](https://github.com/schollz/progressbar)  Basic thread-safe progress bar that works in every OS.
* [mpb](https://github.com/vbauerster/mpb) **star:839** Multi progress bar for terminal applications.   [![There was an update last week][Green]](https://github.com/vbauerster/mpb)   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [aurora](https://github.com/logrusorgru/aurora) **star:777** ANSI terminal colors that supports fmt.Printf/Sprintf.   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [uitable](https://github.com/gosuri/uitable) **star:544** Library to improve readability in terminal apps using tabular data.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:414** Colorable writer for windows.   [![There was an update last week][Green]](https://github.com/mattn/go-colorable)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:398** isatty for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [chalk](https://github.com/ttacon/chalk) **star:327** Intuitive package for prettifying terminal/console output.   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [gookit/color](https://github.com/gookit/color) **star:319** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![Contains Chinese documents][CN]](https://github.com/gookit/color)
* [tabby](https://github.com/cheynewallace/tabby) **star:261** A tiny library for super simple Golang tables.   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [simpletable](https://github.com/alexeyco/simpletable) **star:207** Simple tables in terminal with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:201** Go library for color output in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-colortext)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [cfmt](https://github.com/mingrammer/cfmt) **star:72** Contextual fmt inspired by bootstrap color classes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mingrammer/cfmt)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:36** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/tabular)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [termbox-go](https://github.com/nsf/termbox-go)  Termbox is a library for creating cross-platform text-based interfaces.
* [colourize](https://github.com/TreyBastian/colourize) **star:17** Go library for ANSI colour text in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [ctc](https://github.com/wzshiming/ctc) **star:16** The non-invasive cross-platform terminal color library does not need to modify the Print method.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/ctc)
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** Go library for rendering ANSI colored text templates in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/workanator/go-ataman)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:11255** Go configuration with fangs.   [![star > 2000][Awesome]](https://github.com/spf13/viper)   [![There was an update last week][Green]](https://github.com/spf13/viper)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2786** Go library for managing configuration data from environment variables.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![There was an update last week][Green]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv) **star:2620** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [joshbetz/config](https://github.com/joshbetz/config)  Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.
* [ini](https://github.com/go-ini/ini) **star:1915** Go package to read and write INI files.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:1431** Parse environment variables to Go structs (with defaults).   [![There was an update last week][Green]](https://github.com/caarlos0/env)   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [envcfg](https://github.com/tomazk/envcfg)  Un-marshaling environment variables to Go structs.
* [konfig](https://github.com/lalamove/konfig) **star:547** Composable, observable and performant config handling for Go for the distributed processing era.   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [conflate](https://github.com/the4thamigo-uk/conflate)  Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
* [confita](https://github.com/heetch/confita) **star:294** Load configuration in cascade from multiple backends into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:248** Lightweight configuration manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tucnak/store)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/olebedev/config)  JSON or YAML configuration wrapper with environment variables and flags parsing.
* [config](https://github.com/JeremyLoy/config) **star:228** Cloud native application configuration. Bind ENV to structs in only two lines.   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [hjson](https://github.com/hjson/hjson-go) **star:193** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [koanf](https://github.com/knadh/koanf) **star:163** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![There was an update last week][Green]](https://github.com/knadh/koanf)   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [harvester](https://github.com/beatlabs/harvester)  Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.
* [gookit/config](https://github.com/gookit/config) **star:134** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   [![Contains Chinese documents][CN]](https://github.com/gookit/config)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:126** read INI-style configuration files into Go structs; supports user-defined types and subsections.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [goConfig](https://github.com/crgimenes/goConfig) **star:123** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [envh](https://github.com/antham/envh) **star:94** Helpers to manage environment variables.   [![There was an update last week][Green]](https://github.com/antham/envh)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gofigure](https://github.com/ian-kent/gofigure) **star:58** Go application configuration made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [configure](https://github.com/paked/configure) **star:54** Provides configuration through multiple sources, including JSON, flags and environment variables.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paked/configure)   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [config](https://github.com/golobby/config) **star:49** A lightweight yet powerful config package for Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:46** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:45** Minimalistic configuration reader (from files, ENV, and wherever you want).   [![There was an update last week][Green]](https://github.com/ilyakaznacheev/cleanenv)   [![godoc][GoDoc]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
* [ingo](https://github.com/schachmat/ingo) **star:27** Flags persisted in an ini-like config file.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schachmat/ingo)   [![godoc][GoDoc]](https://godoc.org/github.com/schachmat/ingo)
* [go-up](https://github.com/ufoscout/go-up) **star:26** A simple configuration library with recursive placeholders resolution and no magic.   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:21** Golang package for parsing ini-style configuration files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sasbury/mini)   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [genv](https://github.com/sakirsensoy/genv) **star:9** Read environment variables easily with dotenv support.   [![godoc][GoDoc]](https://godoc.org/github.com/sakirsensoy/genv)
* [envconfig](https://github.com/vrischmann/envconfig)  Read your configuration from environment variables.
* [envconf](https://github.com/ian-kent/envconf) **star:9** Configuration from environment.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/envconf)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/envconf)
* [sprbox](https://github.com/oblq/sprbox) **star:5** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/sprbox)
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) **star:3** Go utility for loading configuration parameters from AWS SSM (Parameter Store).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-ssm-config)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:2** Simple useful package for read environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/nasermirzaei89/env)
* [onion](http://github.com/goraz/onion)  Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:20538** Drone is a Continuous Integration platform built on Docker, written in Go.   [![star > 2000][Awesome]](https://github.com/drone/drone)   [![There was an update last week][Green]](https://github.com/drone/drone)   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [CDS](https://github.com/ovh/cds) **star:2666** Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.   [![star > 2000][Awesome]](https://github.com/ovh/cds)   [![There was an update last week][Green]](https://github.com/ovh/cds)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/cds)
* [duci](https://github.com/duck8823/duci) **star:53** A simple ci server no needs domain specific languages.   [![There was an update last week][Green]](https://github.com/duck8823/duci)   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:43** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)
* [goveralls](https://github.com/mattn/goveralls)  Go integration for Coveralls.io continuous code coverage tracking system.
* [overalls](https://github.com/go-playground/overalls)  Multi-Package go project coverprofile for tools like goveralls.
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** Recursive coverage testing tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:431** Pure Go CSS Preprocessor.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/gcss)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass)  Go wrapper to the 100% Sass compatible libsass project.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:7801** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [gofal](https://github.com/xxjwxc/gofal)  fractional api for Go.
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5452** Collection of useful, performant, and thread-safe data structures.   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1408** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1231** Probabilistic data structures for processing continuous, unbounded streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/BoomFilters)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [gota](https://github.com/kniren/gota) **star:1091** Implementation of dataframes, series, and data wrangling methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [iter](https://github.com/disksing/iter)  Go implementation of C++ STL iterators and algorithms.
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:686** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:586** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [deque](https://github.com/edwingeng/deque)  A highly optimized double-ended queue.
* [bitset](https://github.com/willf/bitset) **star:537** Go package implementing bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [algorithms](https://github.com/shady831213/algorithms)  Algorithms and data structures.CLRS study.
* [gocache](https://github.com/eko/gocache) **star:362** A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.   [![There was an update last week][Green]](https://github.com/eko/gocache)   [![godoc][GoDoc]](https://godoc.org/github.com/eko/gocache)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache)  Fast in-memory key:value store/cache library. Pointer caches.
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:320** In-memory geo index.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hailocab/go-geoindex)   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [mafsa](https://github.com/smartystreets/mafsa) **star:281** MA-FSA implementation with Minimal Perfect Hashing.   [![godoc][GoDoc]](https://godoc.org/github.com/smartystreets/mafsa)   [![Archived][Archived]](https://github.com/smartystreets/mafsa)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:217** Skip list implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [hilbert](https://github.com/google/hilbert) **star:196** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   [![It hasn't been updated in the last year][Yellow]](https://github.com/google/hilbert)   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [merkletree](https://github.com/cbergoon/merkletree) **star:180** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:134** Binary packer and unpacker helps user build custom binary stream.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhuangsirui/binpacker)   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:132** Go implementation of Adaptive Radix Tree.   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [bloom](https://github.com/zhenjl/bloom) **star:131** Bloom filters implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [typ](https://github.com/gurukami/typ)  Null Types, Safe primitive type conversion and fetching value from complex structures.
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:126** In-memory LRU string-interface{} map with expiration for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/diegobernardes/ttlcache)
* [willf/bloom](https://github.com/willf/bloom)  Go package implementing Bloom filters.
* [skiplist](https://github.com/MauriceGit/skiplist) **star:116** Very fast Go Skiplist implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [deque](https://github.com/gammazero/deque) **star:104** Fast ring-buffer deque (double-ended queue).   [![There was an update last week][Green]](https://github.com/gammazero/deque)   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:101** Region quadtrees with efficient point location and neighbour finding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aurelien-rainone/go-rquad)   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [roaring](https://github.com/RoaringBitmap/roaring)  Go package implementing compressed bitsets.
* [ring](https://github.com/TheTannerRyan/ring) **star:101** Go implementation of a high performance, thread safe bloom filter.   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [encoding](https://github.com/zhenjl/encoding) **star:98** Integer Compression Libraries for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/encoding)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:86** A small, powerful and flexible merge library.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [bit](https://github.com/yourbasic/bit) **star:80** Golang set data structure with bonus bit-twiddling functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bit)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [gostl](https://github.com/liyue201/gostl) **star:75** Data structure and algorithm library for go, designed to provide functions similar to C++ STL.   [![godoc][GoDoc]](https://godoc.org/github.com/liyue201/gostl)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:70** Implementation to calculate levenshtein distance in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [skiplist](https://github.com/gansidui/skiplist) **star:68** Skiplist implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:63** Concurrent FIFO queue.   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [bloom](https://github.com/yourbasic/bloom) **star:48** Golang Bloom filter implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bloom)
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:46** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/count-min-log)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/count-min-log)
* [crunch](https://github.com/superwhiskers/crunch)  Go package implementing buffers for handling various datatypes easily.
* [levenshtein](https://github.com/agext/levenshtein) **star:38** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [remember-go](https://github.com/rocketlaunchr/remember-go) **star:35** A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/remember-go)
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:26** Highly concurrent drop-in replacement for `bufio.Writer`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/free/concurrent-writer)   [![godoc][GoDoc]](https://godoc.org/github.com/free/concurrent-writer)
* [goset](https://github.com/zoumo/goset) **star:20** A useful Set collection implementation for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [pipeline](https://github.com/hyfather/pipeline) **star:20** An implementation of pipelines with fan-in and fan-out.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyfather/pipeline)   [![godoc][GoDoc]](https://godoc.org/github.com/hyfather/pipeline)
* [ptrie](https://github.com/viant/ptrie)  An implementation of prefix tree.
* [hide](https://github.com/emvi/hide) **star:13** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [dict](https://github.com/srfrog/dict) **star:11** Python-like dictionaries (dict) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [go-ef](https://github.com/amallia/go-ef) **star:11** A Go implementation of the Elias-Fano encoding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/amallia/go-ef)   [![godoc][GoDoc]](https://godoc.org/github.com/amallia/go-ef)
* [mspm](https://github.com/BlackRabbitt/mspm) **star:9** Multi-String Pattern Matching Algorithm for information retrieval.   [![It hasn't been updated in the last year][Yellow]](https://github.com/BlackRabbitt/mspm)   [![godoc][GoDoc]](https://godoc.org/github.com/BlackRabbitt/mspm)
* [null](https://github.com/emvi/null) **star:9** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [treap](https://github.com/perdata/treap) **star:7** Persistent, fast ordered map using tree heaps.   [![godoc][GoDoc]](https://godoc.org/github.com/perdata/treap)
* [trie](https://github.com/derekparker/trie)  Trie implementation in Go.
* [set](https://github.com/StudioSol/set) **star:7** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:6** Map with expiring key-value pairs.   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:29286** Monitoring system and time series database.   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   [![There was an update last week][Green]](https://github.com/prometheus/prometheus)   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:17813** Scalable, Geo-Replicated, Transactional Datastore.   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   [![There was an update last week][Green]](https://github.com/cockroachdb/cockroach)   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [Coffer](https://github.com/claygod/coffer)  Simple ACID key-value database that supports transactions.
* [dgraph](https://github.com/dgraph-io/dgraph) **star:12510** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   [![There was an update last week][Green]](https://github.com/dgraph-io/dgraph)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [ledisdb](https://github.com/siddontang/ledisdb)  Ledisdb is a high performance NoSQL like Redis based on LevelDB.
* [Kivik](https://github.com/go-kivik/kivik)  Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.
* [influxdb](https://github.com/influxdb/influxdb)  Scalable datastore for metrics, events, and real-time analytics.
* [groupcache](https://github.com/golang/groupcache) **star:8350** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:7318** Fast key-value store in Go.   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   [![There was an update last week][Green]](https://github.com/dgraph-io/badger)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:5686** The lightweight, distributed, relational database built on SQLite.   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   [![There was an update last week][Green]](https://github.com/rqlite/rqlite)   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [BigCache](https://github.com/allegro/bigcache) **star:3599** Efficient key/value cache for gigabytes of data.   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3557** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![There was an update last week][Green]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [go-cache](https://github.com/pmylund/go-cache) **star:3500** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [bcache](https://github.com/iwanbk/bcache)  Eventually consistent distributed in-memory  cache Go library.
* [bbolt](https://github.com/etcd-io/bbolt) **star:2981** An embedded key/value database for Go.   [![star > 2000][Awesome]](https://github.com/etcd-io/bbolt)   [![godoc][GoDoc]](https://godoc.org/github.com/etcd-io/bbolt)
* [buntdb](https://github.com/tidwall/buntdb) **star:2635** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1831** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   [![There was an update last week][Green]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [cache2go](https://github.com/muesli/cache2go) **star:1210** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:1090** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   [![There was an update last week][Green]](https://github.com/xujiajun/nutsdb)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   [![Contains Chinese documents][CN]](https://github.com/xujiajun/nutsdb)
* [GCache](https://github.com/bluele/gcache) **star:1059** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:993** CovenantSQL is a SQL database on blockchain.   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:855** Home-grown disk-backed key-value store.   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [fastcache](https://github.com/VictoriaMetrics/fastcache)  fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.
* [eliasdb](https://github.com/krotik/eliasdb) **star:558** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [levigo](https://github.com/jmhodges/levigo) **star:378** Levigo is a Go wrapper for LevelDB.   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [moss](https://github.com/couchbase/moss)  Moss is a simple LSM key-value storage engine written in 100% Go.
* [Bitcask](https://github.com/prologic/bitcask) **star:290** Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).   [![There was an update last week][Green]](https://github.com/prologic/bitcask)   [![godoc][GoDoc]](https://godoc.org/github.com/prologic/bitcask)
* [pudge](https://github.com/recoilme/pudge) **star:246** Fast and simple  key/value store written using Go's standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [piladb](https://github.com/fern4lvarez/piladb) **star:173** Lightweight RESTful database engine based on stack data structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fern4lvarez/piladb)   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Vasto](https://github.com/chrislusf/vasto) **star:169** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:93** Key-value store with persistence.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:88** Tiny flat file JSON store.   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:47** RESTful caching micro-service backed by Couchbase server.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:33** BigCache with clustering support and individual item expiration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][GoDoc]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
* [cache](https://github.com/akyoto/cache) **star:33** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.   [![There was an update last week][Green]](https://github.com/akyoto/cache)   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)
* [Databunker](https://github.com/paranoidguy/databunker) **star:28** Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.   [![There was an update last week][Green]](https://github.com/paranoidguy/databunker)   [![godoc][GoDoc]](https://godoc.org/github.com/paranoidguy/databunker)
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:14** Key-value store for temporary items.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/tempdb)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/tempdb)
* [tidb](https://github.com/pingcap/tidb)  TiDB is a distributed SQL database. Inspired by the design of Google F1.
* [tiedot](https://github.com/HouzuoGuo/tiedot)  Your NoSQL database powered by Golang.
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:12** Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kapitan-k/gorocksdb)   [![godoc][GoDoc]](https://godoc.org/github.com/kapitan-k/gorocksdb)
* [tracedb](https://github.com/unit-io/tracedb) **star:8** Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application.   [![There was an update last week][Green]](https://github.com/unit-io/tracedb)   [![godoc][GoDoc]](https://godoc.org/github.com/unit-io/tracedb)
* [gostore](https://github.com/twharmon/gostore) **star:3** Gostore is a simple, durable, embedded key-value storage engine written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/gostore)

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:3726** Database migrations. CLI and Golang library.   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![There was an update last week][Green]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1624** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:411** Database schema migration helper for Gorm ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [goose](https://github.com/steinbacher/goose) **star:134** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   [![It hasn't been updated in the last year][Yellow]](https://github.com/steinbacher/goose)   [![godoc][GoDoc]](https://godoc.org/github.com/steinbacher/goose)
* [migrator](https://github.com/lopezator/migrator) **star:87** Dead simple Go database migration library.   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:44** A Go package to help write migrations with go-pg/pg.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [gondolier](https://github.com/emvi/gondolier) **star:28** Database migration library using struct decorators.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/gondolier)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pravasan/pravasan)
* [avro](https://github.com/khezen/avro) **star:11** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)
* [darwin](https://github.com/GuiaBolso/darwin)  Database schema evolution library for Go.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures)  Django style fixtures for Golang's excellent built-in database/sql library.
* [schema](https://github.com/adlio/schema) **star:4** Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/schema)
* [skeema](https://github.com/skeema/skeema)  Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:9603** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   [![There was an update last week][Green]](https://github.com/youtube/vitess)   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [prep](https://github.com/hexdigest/prep)  Use prepared SQL statements without changing your code.
* [pgweb](https://github.com/sosedoff/pgweb) **star:6283** Web-based PostgreSQL database browser.   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![There was an update last week][Green]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [pREST](https://github.com/nuveo/prest)  Serve a RESTful API from any PostgreSQL database.
* [myreplication](https://github.com/2tvenom/myreplication)  MySql binary log replication listener. Supports statement and row based replication.
* [kingshard](https://github.com/flike/kingshard) **star:5015** kingshard is a high performance proxy for MySQL powered by Golang.   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   [![Contains Chinese documents][CN]](https://github.com/flike/kingshard)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2783** Sync your MySQL data into Elasticsearch automatically.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:2243** Go toolset to handle MySQL protocol and replication.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql)   [![There was an update last week][Green]](https://github.com/siddontang/go-mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:174** Collects small insterts and sends big requests to ClickHouse servers.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [datagen](https://github.com/codingconcepts/datagen)  A fast data generator that's multi-table aware and supports multi-row DML.
* [octillery](https://github.com/knocknote/octillery) **star:73** Go package for sharding databases ( Supports every ORM or raw SQL ).   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [orchestrator](https://github.com/github/orchestrator)  MySQL replication topology manager & visualizer.
* [dbbench](https://github.com/sj14/dbbench) **star:35** Database benchmarking tool with support for several databases and scripts.   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [rwdb](https://github.com/andizzle/rwdb) **star:11** rwdb provides read replica capability for multiple database servers setup.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andizzle/rwdb)   [![godoc][GoDoc]](https://godoc.org/github.com/andizzle/rwdb)
* [bucket](https://github.com/PumpkinSeed/bucket) **star:6** Optimized data structure framework for Couchbase specialized on one bucket usage.   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/bucket)
* [chproxy](https://github.com/Vertamedia/chproxy)  HTTP proxy for ClickHouse database.

*SQL query builder, libraries for building and using SQL.*

* [gendry](https://github.com/didi/gendry) **star:945** Non-invasive SQL builder and powerful data binder.   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   [![Contains Chinese documents][CN]](https://github.com/didi/gendry)
* [jet](https://github.com/go-jet/jet) **star:217** Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.   [![godoc][GoDoc]](https://godoc.org/github.com/go-jet/jet)
* [dbq](https://github.com/rocketlaunchr/dbq) **star:93** Zero boilerplate database operations for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)
* [Dotsql](https://github.com/gchaincl/dotsql)  Go library that helps you keep sql files in one place and use them with ease.
* [igor](https://github.com/galeone/igor) **star:80** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [godbal](https://github.com/xujiajun/godbal) **star:50** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xujiajun/godbal)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [goqu](https://github.com/doug-martin/goqu)  Idiomatic SQL builder and query library.
* [sqlf](https://github.com/leporo/sqlf) **star:7** Fast SQL query builder.   [![godoc][GoDoc]](https://godoc.org/github.com/leporo/sqlf)
* [sqrl](https://github.com/elgris/sqrl)  SQL query builder, fork of Squirrel with improved performance.
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [Squirrel](https://github.com/Masterminds/squirrel)  Go library that helps you build SQL queries.
* [xo](https://github.com/knq/xo)  Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.
* [qry](https://github.com/HnH/qry) **star:6** Tool that generates constants from files with raw SQL queries.   [![godoc][GoDoc]](https://godoc.org/github.com/HnH/qry)
* [mpath](https://github.com/spacetab-io/mpath-go) **star:4** MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation.   [![godoc][GoDoc]](https://godoc.org/github.com/spacetab-io/mpath-go)
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)  Powerful data retrieval methods as well as DB-agnostic query building capabilities.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:9039** MySQL driver for Go.   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   [![There was an update last week][Green]](https://github.com/go-sql-driver/mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3)  SQLite3 driver for go that uses database/sql.
    * [gofreetds](https://github.com/minus5/gofreetds)  Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1146** Microsoft MSSQL driver for Go.   [![There was an update last week][Green]](https://github.com/denisenkom/go-mssqldb)   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8)  Oracle driver for go that uses database/sql.
    * [goracle](https://github.com/go-goracle/goracle) **star:279** Oracle driver for Go, using the ODPI-C driver.
    * [pgx](https://github.com/jackc/pgx)  PostgreSQL driver supporting features beyond those exposed by database/sql.
    * [pq](https://github.com/lib/pq)  Pure Go Postgres driver for database/sql.
    * [go-adodb](https://github.com/mattn/go-adodb) **star:102** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![There was an update last week][Green]](https://github.com/mattn/go-adodb)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:46** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [bgc](https://github.com/viant/bgc) **star:13** Datastore Connectivity for BigQuery for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)
    * [firebirdsql](https://github.com/nakagami/firebirdsql)  Firebird RDBMS SQL driver for Go.

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:8083** Redis client for Golang.   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   [![There was an update last week][Green]](https://github.com/go-redis/redis)   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo)  Redigo is a Go client for the Redis database.
    * [redeo](https://github.com/bsm/redeo)  Redis-protocol compatible TCP servers/services.
    * [neoism](https://github.com/jmcvetta/neoism) **star:361** Neo4j client for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:315** Aerospike client in Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:297** Couchbase client in Go.   [![There was an update last week][Green]](https://github.com/couchbase/go-couchbase)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:115** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [gocb](https://github.com/couchbase/gocb)  Official Couchbase Go SDK.
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [gorethink](https://github.com/dancannon/gorethink)  Go language driver for RethinkDB.
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [godscache](https://github.com/defcronyke/godscache)  A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
    * [godis](https://github.com/piaohao/godis) **star:69** redis client implement by golang, inspired by jedis.   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [arangolite](https://github.com/solher/arangolite) **star:66** Lightweight golang driver for ArangoDB.   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [dynago](https://github.com/underarmour/dynago) **star:66** Dynago is a principle of least surprise client for DynamoDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/underarmour/dynago)   [![godoc][GoDoc]](https://godoc.org/github.com/underarmour/dynago)
    * [asc](https://github.com/viant/asc)  Datastore Connectivity for Aerospike for go.
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:34** Go client library for Pilosa.   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [forestdb](https://github.com/couchbase/goforestdb) **star:28** Go bindings for ForestDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/couchbase/goforestdb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/goforestdb)
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zegl/goriak)   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [mgm](https://github.com/kamva/mgm)  MongoDB model-based ODM for Go (based on official MongoDB driver).
    * [mgo](https://github.com/globalsign/mgo)  (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)  Official MongoDB driver for the Go language.
    * [neo4j](https://github.com/cihangir/neo4j)  Neo4j Rest API Bindings for Golang.
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO)  Neo4j REST Client in golang.
    * [xredis](https://github.com/shomali11/xredis) **star:10** Typesafe, customizable, clean & easy to use Redis client.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:6356** Modern text indexing library for go.   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [elastic](https://github.com/olivere/elastic)  Elasticsearch client for Go.
    * [elasticsql](https://github.com/cch123/elasticsql)  Convert sql to elasticsearch dsl in Go.
    * [elastigo](https://github.com/mattbaird/elastigo)  Elasticsearch client library.
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch)  Official Elasticsearch client for Go.
    * [riot](https://github.com/go-ego/riot) **star:5056** Go Open Source, Distributed, Simple and efficient Search Engine.   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   [![Contains Chinese documents][CN]](https://github.com/go-ego/riot)
    * [skizze](https://github.com/seiflotfy/skizze) **star:72** probabilistic data-structures service and storage.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/skizze)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OwnLocal/goes)   [![godoc][GoDoc]](https://godoc.org/github.com/OwnLocal/goes)

* Multiple Backends.
    * [gokv](https://github.com/philippgille/gokv) **star:169** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [cachego](https://github.com/fabiorphp/cachego)  Golang Cache component for multiple drivers.
    * [cayley](https://github.com/google/cayley)  Graph database with support for multiple backends.
    * [dsc](https://github.com/viant/dsc) **star:17** Datastore connectivity for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2407** Now is a time toolkit for golang.   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [durafmt](https://github.com/hako/durafmt) **star:282** Time duration formatting library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [carbon](https://github.com/uniplaces/carbon)  Simple Time extension with a lot of util methods, ported from PHP Carbon library.
* [timeutil](https://github.com/leekchan/timeutil)  Useful extensions (Timedelta, Strftime, ...) to the golang's time package.
* [timespan](https://github.com/SaidinWoT/timespan) **star:68** For interacting with intervals of time, defined as a start time and a duration.   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [tuesday](https://github.com/osteele/tuesday)  Ruby-compatible Strftime function.
* [feiertage](https://github.com/wlbr/feiertage) **star:25** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar)  The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
* [kair](https://github.com/GuilhermeCaruso/kair) **star:15** Date and Time - Golang Formatting Library.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:10** Nullable `time.Time`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/nulltime)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/nulltime)
* [cronrange](https://github.com/1set/cronrange) **star:8** Parses Cron-style time range expressions, checks if the given time is within any ranges.   [![godoc][GoDoc]](https://godoc.org/github.com/1set/cronrange)
* [dateparse](https://github.com/araddon/dateparse)  Parse date's without knowing format in advance.
* [date](https://github.com/rickb777/date)  Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
* [strftime](https://github.com/awoodbeck/strftime) **star:4** C99-compatible strftime formatter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awoodbeck/strftime)   [![godoc][GoDoc]](https://godoc.org/github.com/awoodbeck/strftime)
* [iso8601](https://github.com/relvacode/iso8601)  Efficiently parse ISO8601 date-times without regex.
* [go-week](https://github.com/stoewer/go-week) **star:2** An efficient package to work with ISO8601 week dates.   [![godoc][GoDoc]](https://godoc.org/github.com/stoewer/go-week)
* [go-str2duration](https://github.com/xhit/go-str2duration) **star:1** Convert string to duration. Support time.Duration returned string and more.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-str2duration)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise)  Calculate the sunrise and sunset times for a given location.

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:16263** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![There was an update last week][Green]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [NATS](https://github.com/nats-io/gnatsd)  Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.
* [micro](https://github.com/micro/micro) **star:7660** Pluggable microservice toolkit and distributed systems platform.   [![star > 2000][Awesome]](https://github.com/micro/micro)   [![There was an update last week][Green]](https://github.com/micro/micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [rpcx](https://github.com/smallnest/rpcx) **star:4381** Distributed pluggable RPC service framework like alibaba Dubbo.   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   [![There was an update last week][Green]](https://github.com/smallnest/rpcx)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [torrent](https://github.com/anacrolix/torrent)  BitTorrent client package.
* [tendermint](https://github.com/tendermint/tendermint) **star:3503** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   [![There was an update last week][Green]](https://github.com/tendermint/tendermint)   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [dragonboat](https://github.com/lni/dragonboat) **star:2884** A feature complete and high performance multi-group Raft library in Go.   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   [![There was an update last week][Green]](https://github.com/lni/dragonboat)   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   [![Contains Chinese documents][CN]](https://github.com/lni/dragonboat)
* [drmaa](https://github.com/dgruber/drmaa)  Job submission library for cluster schedulers based on the DRMAA standard.
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [glow](https://github.com/chrislusf/glow) **star:2750** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [gleam](https://github.com/chrislusf/gleam) **star:2389** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2231** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [hprose](https://github.com/hprose/hprose-golang)  Very newbility RPC Library, support 25+ languages now.
* [gorpc](https://github.com/valyala/gorpc) **star:577** Simple, fast and scalable RPC library for high load.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [grpc-go](https://github.com/grpc/grpc-go)  The Go language implementation of gRPC. HTTP/2 based RPC.
* [go-health](https://github.com/InVisionApp/go-health) **star:529** Library for enabling asynchronous dependency health checks in your service.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [rain](https://github.com/cenkalti/rain) **star:474** BitTorrent client and library.   [![godoc][GoDoc]](https://godoc.org/github.com/cenkalti/rain)
* [sleuth](https://github.com/ursiform/sleuth) **star:312** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ursiform/sleuth)   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [go-jump](https://github.com/dgryski/go-jump) **star:281** Port of Google's "Jump" Consistent Hash function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgryski/go-jump)   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [dht](https://github.com/anacrolix/dht) **star:150** BitTorrent Kademlia DHT implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [digota](https://github.com/digota/digota)  grpc ecommerce microservice.
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:122** The jsonrpc package helps implement of JSON-RPC 2.0.   [![There was an update last week][Green]](https://github.com/osamingo/jsonrpc)   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:116** JSON-RPC 2.0 HTTP client implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [KrakenD](https://github.com/devopsfaith/krakend)  Ultra performant API Gateway framework with middlewares.
* [liftbridge](https://github.com/liftbridge-io/liftbridge)  Lightweight, fault-tolerant message streams for NATS.
* [ringpop-go](https://github.com/uber/ringpop-go)  Scalable, fault-tolerant application-layer sharding for Go applications.
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [redis-lock](https://github.com/bsm/redislock) **star:92** Simplified distributed locking implementation using Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redislock)
* [consistent](https://github.com/buraksezer/consistent)  Consistent hashing with bounded loads.
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:59** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/svcavallar/celeriac.v1)   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:50** A revamped Google's jump consistent hash.   [![There was an update last week][Green]](https://github.com/edwingeng/doublejump)   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [raft](https://github.com/hashicorp/raft)  Golang implementation of the Raft consensus protocol, by HashiCorp.
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [outboxer](https://github.com/italolelis/outboxer) **star:26** Outboxer is a go library that implements the outbox pattern.   [![There was an update last week][Green]](https://github.com/italolelis/outboxer)   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [flowgraph](https://github.com/vectaport/flowgraph) **star:25** flow-based programming package.   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit)  A library built to provide support for defining async service health checks for golang services.
* [go-pdu](https://github.com/pdupub/go-pdu) **star:11** A decentralized identity-based social network.   [![There was an update last week][Green]](https://github.com/pdupub/go-pdu)   [![godoc][GoDoc]](https://godoc.org/github.com/pdupub/go-pdu)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:10** A library for using DynamoDB as an atomic counter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylfin/dynatomic)   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)

## Dynamic DNS

*Tools for updating dynamic DNS records.*

* [DDNS](https://github.com/skibish/ddns)  Personal DDNS client with Digital Ocean Networking DNS as backend.
* [dyndns](https://gitlab.com/alcastle/dyndns)  Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.
* [GoDNS](https://github.com/timothyye/godns) **star:545** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:5989** Email and SMTP testing with web and API interface.   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [hermes](https://github.com/matcornic/hermes) **star:1895** Golang package that generates clean, responsive HTML e-mails.   [![There was an update last week][Green]](https://github.com/matcornic/hermes)   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1248** A robust and flexible email library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:893** IMAP library for clients and servers.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [go-message](https://github.com/emersion/go-message)  Streaming library for the Internet Message Format and mail messages.
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:567** SendGrid's Go library for sending email.   [![There was an update last week][Green]](https://github.com/sendgrid/sendgrid-go)   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [smtp](https://github.com/mailhog/smtp)  SMTP server protocol state machine.
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [Hectane](https://github.com/hectane/hectane) **star:180** Lightweight SMTP client providing an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [douceur](https://github.com/aymerick/douceur) **star:173** CSS inliner for your HTML emails.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/douceur)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [go-dkim](https://github.com/toorop/go-dkim) **star:52** DKIM library, to sign & verify email.   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:44** Inline styling for HTML mail in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [mailchain](https://github.com/mailchain/mailchain) **star:39** Send encrypted emails to blockchain addresses written in Go.   [![There was an update last week][Green]](https://github.com/mailchain/mailchain)   [![godoc][GoDoc]](https://godoc.org/github.com/mailchain/mailchain)
* [mailgun-go](https://github.com/mailgun/mailgun-go)  Go library for sending mail with the Mailgun API.
* [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:18** Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-simple-mail)

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) **star:5133** JavaScript interpreter written in Go.   [![star > 2000][Awesome]](https://github.com/robertkrimen/otto)   [![godoc][GoDoc]](https://godoc.org/github.com/robertkrimen/otto)
* [go-lua](https://github.com/Shopify/go-lua) **star:1784** Port of the Lua 5.2 VM to pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [tengo](https://github.com/d5/tengo) **star:1558** Bytecode compiled script language for Go.   [![There was an update last week][Green]](https://github.com/d5/tengo)   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [expr](https://github.com/antonmedv/expr) **star:1098** an engine that can evaluate expressions.   [![There was an update last week][Green]](https://github.com/antonmedv/expr)   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [go-python](https://github.com/sbinet/go-python) **star:1001** naive go bindings to the CPython C-API.   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [anko](https://github.com/mattn/anko) **star:979** Scriptable interpreter written in Go.   [![There was an update last week][Green]](https://github.com/mattn/anko)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-php](https://github.com/deuill/go-php) **star:727** PHP bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/deuill/go-php)   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:683** Duktape JavaScript engine bindings for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [golua](https://github.com/aarzilli/golua) **star:464** Go bindings for Lua C API.
* [gopher-lua](https://github.com/yuin/gopher-lua)  Lua 5.1 VM and compiler written in Go.
* [gisp](https://github.com/jcla1/gisp) **star:434** Simple LISP in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jcla1/gisp)   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [cel-go](https://github.com/google/cel-go) **star:369** Fast, portable, non-Turing complete expression evaluation with gradual typing.   [![There was an update last week][Green]](https://github.com/google/cel-go)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cel-go)
* [gval](https://github.com/PaesslerAG/gval) **star:180** A highly customizable expression language written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [ngaro](https://github.com/db47h/ngaro)  Embeddable Ngaro VM implementation enabling scripting in Retro.
* [gentee](https://github.com/gentee/gentee) **star:44** Embeddable scripting programming language.   [![There was an update last week][Green]](https://github.com/gentee/gentee)   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)
* [binder](https://github.com/alexeyco/binder) **star:34** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexeyco/binder)   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [purl](https://github.com/ian-kent/purl) **star:29** Perl 5.18.2 embedded in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/purl)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/purl)

## Error Handling

*Libraries for handling errors.*

* [go-multierror](https://github.com/hashicorp/go-multierror) **star:814** Go (golang) package for representing a list of errors as a single error.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [errorx](https://github.com/joomcode/errorx) **star:608** A feature rich error package with stack traces, composition of errors and more.   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:589** Golang errors with stack trace and source fragments.   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [eris](https://github.com/rotisserie/eris) **star:570** A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.   [![There was an update last week][Green]](https://github.com/rotisserie/eris)   [![godoc][GoDoc]](https://godoc.org/github.com/rotisserie/eris)
* [errlog](https://github.com/snwfdhmp/errlog)  Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.
* [emperror](https://github.com/emperror/emperror) **star:110** Error handling tools and best practices for Go libraries and applications.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/emperror)
* [errors](https://github.com/emperror/errors) **star:41** Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/errors)
* [errors](https://github.com/pkg/errors)  Package that provides simple error handling primitives.
* [errors](https://github.com/neuronlabs/errors)  Simple golang error handling with classification primitives.
* [errors](https://github.com/PumpkinSeed/errors)  The most simple error wrapper with awesome performance and minimal memory overhead.
* [werr](https://github.com/txgruppi/werr) **star:13** Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/werr)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/werr)
* [Falcon](https://github.com/SonicRoshan/falcon) **star:2** A Simple Yet Highly Powerful Package For Error Handling.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/falcon)

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2590** FileSystem Abstraction System for Go.   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/pdfcpu/pdfcpu) **star:1322** PDF processor.   [![godoc][GoDoc]](https://godoc.org/github.com/pdfcpu/pdfcpu)
* [copy](https://github.com/otiai10/copy) **star:133** Copy directory recursively.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/copy)
* [bigfile](https://github.com/bigfile/bigfile) **star:111** A file transfer system, support to manage files with http api, rpc call and ftp client.   [![godoc][GoDoc]](https://godoc.org/github.com/bigfile/bigfile)   [![Contains Chinese documents][CN]](https://github.com/bigfile/bigfile)
* [opc](https://github.com/qmuntal/opc) **star:63** Load Open Packaging Conventions (OPC) files for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [parquet](https://github.com/parsyl/parquet)  Read and write [parquet](https://parquet.apache.org) files.
* [tarfs](https://github.com/posener/tarfs)  Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [skywalker](https://github.com/dixonwille/skywalker) **star:55** Package to allow one to concurrently go through a filesystem with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/skywalker)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/skywalker)
* [vfs](https://github.com/C2FO/vfs) **star:46** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [afs](https://github.com/viant/afs) **star:40** Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/afs)
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:22** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)
* [gut/yos](https://github.com/1set/gut)  Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.
* [notify](https://github.com/rjeczalik/notify)  File system event notification library with simple API, similar to os/signal.
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:20** Load gtfs files in go.   [![There was an update last week][Green]](https://github.com/artonge/go-gtfs)   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [checksum](https://github.com/codingsince1985/checksum) **star:16** Compute message digest, like MD5 and SHA256, for large files.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [go-csv-tag](https://github.com/artonge/go-csv-tag)  Load csv file using tag.
* [flop](https://github.com/homedepot/flop) **star:15** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:14** Copy files for humans.   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:1942** Arbitrary-precision fixed-point decimal numbers.   [![There was an update last week][Green]](https://github.com/shopspring/decimal)   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-finance](https://github.com/FlashBoys/go-finance)  Comprehensive financial markets data in Go.
* [go-finance](https://github.com/alpeb/go-finance)  Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.
* [go-money](https://github.com/rhymond/go-money) **star:710** Implementation of Fowler's Money pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [ofxgo](https://github.com/aclindsa/ofxgo)  Query OFX servers and/or parse the responses (with example command-line client).
* [orderbook](https://github.com/i25959341/orderbook)  Matching Engine for Limit Order Book in Golang.
* [accounting](https://github.com/leekchan/accounting) **star:536** money and currency formatting for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [techan](https://github.com/sdcoffey/techan) **star:227** Technical analysis library with advanced market analysis and trading strategies.   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [transaction](https://github.com/claygod/transaction) **star:65** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [vat](https://github.com/dannyvankooten/vat)  VAT number validation & EU VAT rates.
* [currency](https://github.com/bnkamalesh/currency) **star:19** High performant & accurate currency computation package.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)
* [go-finnhub](https://github.com/m1/go-finnhub) **star:17** Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-finnhub)
* [go-finance](https://github.com/pieterclaerhout/go-finance) **star:3** Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-finance)

## Forms

*Libraries for working with forms.*

* [binding](https://github.com/mholt/binding) **star:766** Binds form and JSON data from net/http Request to struct.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/binding)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [conform](https://github.com/leebenson/conform)  Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:501** CSRF protection for Go web applications & services.   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [nosurf](https://github.com/justinas/nosurf)  CSRF protection middleware for Go.
* [form](https://github.com/go-playground/form) **star:395** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [formam](https://github.com/monoculum/formam) **star:137** decode form's values into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:104** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/forms)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)
* [bind](https://github.com/robfig/bind) **star:24** Bind form data to any Go values.   [![It hasn't been updated in the last year][Yellow]](https://github.com/robfig/bind)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/bind)
* [queryparam](https://github.com/tomwright/queryparam) **star:1** Decode `url.Values` into usable struct values of standard or custom types.   [![godoc][GoDoc]](https://godoc.org/github.com/tomwright/queryparam)

## Functional

*Packages to support functional programming in Go.*

* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:136** Monad, Functional Programming features for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TeaEntityLab/fpGo)   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:62** Functional Experiment in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)
* [go-underscore](https://github.com/tobyhede/go-underscore)  Useful collection of helpfully functional Go collection utilities.

## Game Development

*Awesome game development libraries.*

* [go3d](https://github.com/ungerik/go3d)  Performance oriented 2D/3D math package for Go.
* [engo](https://github.com/EngoEngine/engo) **star:1165** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![There was an update last week][Green]](https://github.com/EngoEngine/engo)   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [g3n](https://github.com/g3n/engine)  Go 3D Game Engine.
* [GarageEngine](https://github.com/vova616/GarageEngine)  2d game engine written in Go working on OpenGL.
* [glop](https://github.com/runningwild/glop)  Glop (Game Library Of Power) is a fairly simple cross-platform game library.
* [go-astar](https://github.com/beefsack/go-astar)  Go implementation of the A\* path finding algorithm.
* [go-collada](https://github.com/GlenKelley/go-collada)  Go package for working with the Collada file format.
* [go-sdl2](https://github.com/veandco/go-sdl2)  Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [Pixel](https://github.com/faiface/pixel)  Hand-crafted 2D game library in Go.
* [gonet](https://github.com/xtaci/gonet) **star:1088** Game server skeleton implemented with golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xtaci/gonet)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [goworld](https://github.com/xiaonanln/goworld)  Scalable game server engine, featuring space-entity framework and hot-swapping.
* [Leaf](https://github.com/name5566/leaf)  Lightweight game server framework.
* [nano](https://github.com/lonng/nano)  Lightweight, facility, high performance golang based game server framework.
* [Oak](https://github.com/oakmound/oak)  Pure Go game engine.
* [Pitaya](https://github.com/topfreegames/pitaya)  Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:435** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [termloop](https://github.com/JoelOtter/termloop)  Terminal-based game engine for Go, built on top of Termbox.
* [Ebiten](https://github.com/hajimehoshi/ebiten)  dead simple 2D game library in Go.
* [Azul3D](https://github.com/azul3d/engine)  3D game engine written in Go.

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [jennifer](https://github.com/dave/jennifer) **star:1448** Generate arbitrary Go code without templates.   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [pkgreflect](https://github.com/ungerik/pkgreflect)  Go preprocessor for package scoped reflection.
* [gen](https://github.com/clipperhouse/gen) **star:1091** Code generation tool for ‘generics’-like functionality.   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [generis](https://github.com/senselogic/GENERIS)  Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
* [go-enum](https://github.com/abice/go-enum)  Code generation for enums from code comments.
* [go-linq](https://github.com/ahmetalpbalkan/go-linq)  .NET LINQ-like query methods for Go.
* [go-xray](https://github.com/pieterclaerhout/go-xray)  Helpers for making the use of reflection easier.
* [goderive](https://github.com/awalterschulze/goderive)  Derives functions from input types.
* [efaceconv](https://github.com/t0pep0/efaceconv)  Code generation tool for high performance conversion from interface{} to immutable type without allocations.
* [interfaces](https://github.com/rjeczalik/interfaces) **star:205** Command line tool for generating interface definitions.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [gotype](https://github.com/wzshiming/gotype) **star:29** Golang source code parsing, usage like reflect package.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/gotype)
* [GoWrap](https://github.com/hexdigest/gowrap)  Generate decorators for Go interfaces using simple templates.
* [typeregistry](https://github.com/xiaoxin01/typeregistry) **star:3** A library to create type dynamically.   [![There was an update last week][Green]](https://github.com/xiaoxin01/typeregistry)   [![godoc][GoDoc]](https://godoc.org/github.com/xiaoxin01/typeregistry)

## Geographic

*Geographic tools and servers*

* [S2 geometry](https://github.com/golang/geo) **star:1007** S2 geometry library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [Tile38](https://github.com/tidwall/tile38)  Geolocation DB with spatial index and realtime geofencing.
* [WGS84](https://github.com/wroge/wgs84)  Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).
* [mbtileserver](https://github.com/consbio/mbtileserver) **star:123** A simple Go-based server for map tiles stored in mbtiles format.   [![There was an update last week][Green]](https://github.com/consbio/mbtileserver)   [![godoc][GoDoc]](https://godoc.org/github.com/consbio/mbtileserver)
* [osm](https://github.com/paulmach/osm)  Library for reading, writing and working with OpenStreetMap data and APIs.
* [pbf](https://github.com/maguro/pbf)  OpenStreetMap PBF golang encoder/decoder.
* [geocache](https://github.com/melihmucuk/geocache) **star:119** In-memory cache that is suitable for geolocation based applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/melihmucuk/geocache)   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [geoserver](https://github.com/hishamkaram/geoserver)  geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
* [gismanager](https://github.com/hishamkaram/gismanager) **star:26** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hishamkaram/gismanager)   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)

## Go Compilers

*Tools for compiling Go to other languages.*

* [c4go](https://github.com/Konstantin8105/c4go) **star:199** Transpile C code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:23** Transpile FORTRAN 77 code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)
* [gopherjs](https://github.com/gopherjs/gopherjs)  Compiler from Go to JavaScript.
* [llgo](https://github.com/go-llvm/llgo)  LLVM-based compiler for Go.
* [tardisgo](https://github.com/tardisgo/tardisgo)  Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) **star:3068** A high-performance and low-cost goroutine pool in Go.   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   [![There was an update last week][Green]](https://github.com/panjf2000/ants)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/ants)
* [async](https://github.com/studiosol/async)  A safe way to execute functions asynchronously, recovering them in case of panic.
* [breaker](https://github.com/kamilsk/breaker)  Flexible mechanism to make execution flow interruptible.
* [conexec](https://github.com/ITcathyh/conexec)  A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier)  CyclicBarrier for golang.
* [go-floc](https://github.com/workanator/go-floc)  Orchestrate goroutines with ease.
* [artifex](https://github.com/borderstech/artifex)  Simple in-memory job queue for Golang using worker-based dispatching.
* [queue](https://github.com/AnikHasibul/queue)  Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.
* [pool](https://github.com/go-playground/pool) **star:544** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [semaphore](https://github.com/kamilsk/semaphore)  Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
* [routine](https://github.com/x-mod/routine)  go routine control with context, support: Main, Go, Pool and some useful Executors.
* [semaphore](https://github.com/marusama/semaphore)  Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
* [stl](https://github.com/ssgreg/stl)  Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
* [Hunch](https://github.com/AaronJan/Hunch)  Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn)  Run functions in parallel.
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [kyoo](https://github.com/dirkaholic/kyoo)  Provides an unlimited job queue and concurrent worker pools.
* [grpool](https://github.com/ivpusic/grpool) **star:541** Lightweight Goroutine pool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/grpool)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [workerpool](https://github.com/gammazero/workerpool) **star:262** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:124** Control goroutines execution order.   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools)  Manage a pool of goroutines using this lightweight library with a simple API.
* [gowp](https://github.com/xxjwxc/gowp)  gowp is concurrency limiting goroutine pool.
* [goworker](https://github.com/benmanns/goworker)  goworker is a Go-based background worker.
* [GoSlaves](https://github.com/themester/GoSlaves)  Simple and Asynchronous Goroutine pool library.
* [gollback](https://github.com/vardius/gollback) **star:31** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [gpool](https://github.com/Sherifabdlnaby/gpool)  manages a resizeable pool of context-aware goroutines to bound concurrency.
* [threadpool](https://github.com/shettyh/threadpool) **star:25** Golang threadpool implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [tunny](https://github.com/Jeffail/tunny)  Goroutine pool for golang.
* [worker-pool](https://github.com/vardius/worker-pool)  goworker is a Go simple async worker pool.
* [gohive](https://github.com/loveleshsharma/gohive)  A highly performant and easy to use Goroutine pool for Go.
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup)  Like `sync.WaitGroup` with error handling and concurrency control.
* [go-trylock](https://github.com/subchen/go-trylock) **star:5** TryLock support on read-write lock for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-trylock)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [walk](https://github.com/lxn/walk) **star:4273** Windows application library kit for Go.   [![star > 2000][Awesome]](https://github.com/lxn/walk)   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [webview](https://github.com/zserge/webview)  Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:3025** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   [![There was an update last week][Green]](https://github.com/asticode/go-astilectron)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter)  Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.
* [gotk3](https://github.com/gotk3/gotk3) **star:922** Go bindings for GTK3.   [![There was an update last week][Green]](https://github.com/gotk3/gotk3)   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd)  Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.
* [qt](https://github.com/therecipe/qt)  Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).
* [ui](https://github.com/andlabs/ui)  Platform-native GUI library for Go. Cross platform.
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [app](https://github.com/murlokswarm/app)  Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.
* [fyne](https://github.com/fyne-io/fyne)  Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.

*Interaction*

* [go-appindicator](https://github.com/dawidd6/go-appindicator)  Go bindings for libappindicator3 C library.
* [gosx-notifier](https://github.com/deckarep/gosx-notifier)  OSX Desktop Notifications library for Go.
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker)  OSX library to notify about any (pluggable) activity on your machine.
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier)  OSX Sleep/Wake notifications in golang.
* [robotgo](https://github.com/go-vgo/robotgo)  Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.
* [systray](https://github.com/getlantern/systray)  Cross platform Go library to place an icon and menu in the notification area.
* [trayhost](https://github.com/shurcooL/trayhost) **star:173** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [imaginary](https://github.com/h2non/imaginary) **star:2865** Fast and simple HTTP microservice for image resizing.   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [picfit](https://github.com/thoas/picfit)  An image resizing server written in Go.
* [imaging](https://github.com/disintegration/imaging)  Simple Go image processing package.
* [img](https://github.com/hawx/img)  Selection of image manipulation tools.
* [ln](https://github.com/fogleman/ln)  3D line art rendering in Go.
* [mergi](https://github.com/noelyahan/mergi)  Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
* [mort](https://github.com/aldor007/mort)  Storage and image processing server written in Go.
* [mpo](https://github.com/donatj/mpo)  Decoder and conversion tool for MPO 3D Photos.
* [pt](https://github.com/fogleman/pt)  Path tracing engine written in Go.
* [resize](https://github.com/nfnt/resize)  Image resizing for Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez)  Image resizing in pure Go and SIMD.
* [smartcrop](https://github.com/muesli/smartcrop)  Finds good crops for arbitrary images and crop sizes.
* [steganography](https://github.com/auyer/steganography)  Pure Go Library for LSB steganography.
* [bimg](https://github.com/h2non/bimg) **star:893** Small package for fast and efficient image processing using libvips.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:603** Go tool for LSB steganography, capable of hiding any file within an image.   [![There was an update last week][Green]](https://github.com/DimitarPetrov/stegify)   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [svgo](https://github.com/ajstarks/svgo)  Go Language Library for SVG generation.
* [tga](https://github.com/ftrvxmtrx/tga)  Package tga is a TARGA image format decoder/encoder.
* [canvas](https://github.com/tdewolff/canvas) **star:410** Vector graphics to PDF, SVG or rasterized image.   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/canvas)
* [image2ascii](https://github.com/qeesung/image2ascii)  Convert image to ASCII.
* [imagick](https://github.com/gographics/imagick)  Go binding to ImageMagick's MagickWand C API.
* [govatar](https://github.com/o1egl/govatar)  Library and CMD tool for generating funny avatars.
* [goimghdr](https://github.com/corona10/goimghdr)  The imghdr module determines the type of image contained in a file for Go.
* [goimagehash](https://github.com/corona10/goimagehash) **star:269** Go Perceptual image hashing package.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [bild](https://github.com/anthonynsimon/bild)  Collection of image processing algorithms in pure Go.
* [gift](https://github.com/disintegration/gift)  Package of image processing filters.
* [gg](https://github.com/fogleman/gg)  2D rendering in pure Go.
* [geopattern](https://github.com/pravj/geopattern)  Create beautiful generative image patterns from a string.
* [darkroom](https://github.com/gojek/darkroom) **star:118** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![There was an update last week][Green]](https://github.com/gojek/darkroom)   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [go-webcolors](https://github.com/jyotiska/go-webcolors)  Port of webcolors library from Python to Go.
* [go-opencv](https://github.com/lazywei/go-opencv)  Go bindings for OpenCV.
* [go-nude](https://github.com/koyachi/go-nude)  Nudity detection with Go.
* [gocv](https://github.com/hybridgroup/gocv)  Go package for computer vision using OpenCV 3.3+.
* [go-cairo](https://github.com/ungerik/go-cairo) **star:92** Go binding for the cairo graphics library.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [go-gd](https://github.com/bolknote/go-gd)  Go binding for GD library.
* [gltf](https://github.com/qmuntal/gltf) **star:59** Efficient and robust glTF 2.0 reader, writer and validator.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [cameron](https://github.com/aofei/cameron) **star:40** An avatar generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [connectordb](https://github.com/connectordb/connectordb)  Open-Source Platform for Quantified Self & IoT.
* [devices](https://github.com/goiot/devices)  Suite of libraries for IoT devices, experimental for x/exp/io.
* [eywa](https://github.com/xcodersun/eywa)  Project Eywa is essentially a connection manager that keeps track of connected devices.
* [flogo](https://github.com/tibcosoftware/flogo)  Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt)  Gatt is a Go package for building Bluetooth Low Energy peripherals.
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [huego](https://github.com/amimof/huego) **star:133** An extensive Philips Hue client library for Go.   [![There was an update last week][Green]](https://github.com/amimof/huego)   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [mainflux](https://github.com/Mainflux/mainflux)  Industrial IoT Messaging and Device Management Server.
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [sensorbee](https://github.com/sensorbee/sensorbee)  Lightweight stream processing engine for IoT.

## Job Scheduler

*Libraries for scheduling jobs.*

* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [clockwork](https://github.com/whiteShtef/clockwork)  Simple and intuitive job scheduling library in Go.
* [go-cron](https://github.com/rk/go-cron)  Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
* [gron](https://github.com/roylee0704/gron)  Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.
* [JobRunner](https://github.com/bamzi/jobrunner)  Smart and featureful cron job scheduler with job queuing and live monitoring built in.
* [jobs](https://github.com/albrow/jobs)  Persistent and flexible background jobs library.
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:61** Job scheduler that supports webhooks, crons and classic scheduling.   [![There was an update last week][Green]](https://github.com/kilgaloon/leprechaun)   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)
* [scheduler](https://github.com/carlescere/scheduler)  Cronjobs scheduling made easy.

## JSON

*Libraries for working with JSON.*

* [jettison](https://github.com/wI2L/jettison) **star:69** High performance, reflection-less JSON encoder for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jettison)
* [jsongo](https://github.com/ricardolonga/jsongo)  Fluent API to make it easier to create Json objects.
* [jsonf](https://github.com/miolini/jsonf)  Console tool for highlighted formatting and struct query fetching JSON.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors)  Go bindings based on the JSON API errors reference.
* [json2go](https://github.com/m-zajac/json2go)  Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:63** JSON diff utility written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [ej](https://github.com/lucassscaravelli/ej)  Write and read JSON from different sources succinctly.
* [ajson](https://github.com/spyzhov/ajson)  Abstract JSON for golang with JSONPath support.
* [gojson](https://github.com/ChimeraCoder/gojson)  Automatically generate Go (golang) struct definitions from example JSON.
* [gojq](https://github.com/elgs/gojq)  JSON query in Golang.
* [go-respond](https://github.com/nicklaw5/go-respond)  Go package for handling common HTTP JSON responses.
* [go-jsonerror](https://github.com/ddymko/go-jsonerror)  Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.
* [GJSON](https://github.com/tidwall/gjson)  Get a JSON value with one line of code.
* [gjo](https://github.com/skanehira/gjo)  Small utility to create JSON objects.
* [mp](https://github.com/sanbornm/mp) **star:38** Simple cli email parser. It currently takes stdin and outputs JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sanbornm/mp)   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/mp)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/RichardKnop/jsonhal)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [kazaam](https://github.com/Qntfy/kazaam)  API for arbitrary transformation of JSON documents.
* [mapslice-json](https://github.com/mickep76/mapslice-json)  Go MapSlice for ordered marshal/ unmarshal of maps in JSON.

## Logging

*Libraries for generating and working with log files.*

* [zap](https://github.com/uber-go/zap) **star:9043** Fast, structured, leveled logging in Go.   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![There was an update last week][Green]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [spew](https://github.com/davecgh/go-spew) **star:3740** Implements a deep pretty printer for Go data structures to aid in debugging.   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   [![There was an update last week][Green]](https://github.com/davecgh/go-spew)   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [zerolog](https://github.com/rs/zerolog) **star:2963** Zero-allocation JSON logger.   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [glog](https://github.com/golang/glog) **star:2477** Leveled execution logs for Go.   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1783** Simple rolling logger, implements io.WriteCloser.   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [tail](https://github.com/hpcloud/tail) **star:1713** Go package striving to emulate the features of the BSD tail program.   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [seelog](https://github.com/cihub/seelog) **star:1434** Logging functionality with flexible dispatching, filtering, and formatting.   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log](https://github.com/apex/log) **star:793** Structured logging package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [log](https://github.com/go-playground/log) **star:272** Simple, configurable and scalable Structured Logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [log](https://github.com/teris-io/log)  Structured log interface for Go cleanly separates logging facade from its implementation.
* [logxi](https://github.com/mgutz/logxi)  12-factor app logger that is fast and makes you happy.
* [logutils](https://github.com/hashicorp/logutils) **star:270** Utilities for slightly better logging in Go (Golang) extending the standard logger.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [xlog](https://github.com/rs/xlog) **star:132** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:129** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   [![There was an update last week][Green]](https://github.com/arthurkiller/rollingWriter)   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:84** Full-featured logging saas written in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/logvoyage)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [log15](https://github.com/inconshreveable/log15)  Simple, powerful logging for Go.
* [stdlog](https://github.com/alexcesaro/log) **star:42** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexcesaro/log)   [![godoc][GoDoc]](https://godoc.org/github.com/alexcesaro/log)
* [glg](https://github.com/kpango/glg)  glg is simple and fast leveled logging library for Go.
* [distillog](https://github.com/amoghe/distillog)  distilled levelled logging (think of it as stdlib + log levels).
* [glo](https://github.com/lajosbencz/glo)  PHP Monolog inspired logging facility with identical severity levels.
* [gologger](https://github.com/sadlil/gologger)  Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:26** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [go-log](https://github.com/pieterclaerhout/go-log)  A logging library with strack traces, object dumping and optional timestamps.
* [go-log](https://github.com/subchen/go-log)  Simple and configurable Logging in Go, with level, formatters and writers.
* [go-log](https://github.com/siddontang/go-log)  Log lib supports level and multi handlers.
* [go-log](https://github.com/ian-kent/go-log)  Log4j implementation in Go.
* [go-logger](https://github.com/apsdehal/go-logger)  Simple logger of Go Programs, with level handlers.
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jbrodriguez/mlog)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrodriguez/mlog)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log)  High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
* [onelog](https://github.com/francoispqt/onelog)  Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.
* [log](https://github.com/aerogo/log)  An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).
* [gomol](https://github.com/aphistic/gomol) **star:16** Multiple-output, structured logging for Go with extensible logging outputs.   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [journald](https://github.com/ssgreg/journald)  Go implementation of systemd Journal's native API for logging.
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ewwwwwqm/logdump)   [![godoc][GoDoc]](https://godoc.org/github.com/ewwwwwqm/logdump)
* [logger](https://github.com/azer/logger)  Minimalistic logging library for Go.
* [logex](https://github.com/chzyer/logex)  Golang log lib, supports tracking and level, wrap by standard log lib.
* [logrusly](https://github.com/sebest/logrusly)  [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
* [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:8** `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/logrusiowriter)
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/borderstech/logmatic)   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [xlog](https://github.com/xfxdev/xlog) **star:6** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xfxdev/xlog)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [logrus](https://github.com/Sirupsen/logrus)  Structured logger for Go.
* [logo](https://github.com/mbndr/logo) **star:5** Golang logger to different configurable writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mbndr/logo)   [![godoc][GoDoc]](https://godoc.org/github.com/mbndr/logo)

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:7009** General Machine Learning library for Go.   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   [![Contains Chinese documents][CN]](https://github.com/sjwhitworth/golearn)
* [gorgonia](https://github.com/gorgonia/gorgonia) **star:3160** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![star > 2000][Awesome]](https://github.com/gorgonia/gorgonia)   [![There was an update last week][Green]](https://github.com/gorgonia/gorgonia)   [![godoc][GoDoc]](https://godoc.org/github.com/gorgonia/gorgonia)
* [goml](https://github.com/cdipaolo/goml) **star:1071** On-line Machine Learning in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gosseract](https://github.com/otiai10/gosseract) **star:1053** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   [![There was an update last week][Green]](https://github.com/otiai10/gosseract)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [gorse](https://github.com/zhenghaoz/gorse) **star:920** An offline recommender system backend based on collaborative filtering written in Go.   [![There was an update last week][Green]](https://github.com/zhenghaoz/gorse)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   [![Contains Chinese documents][CN]](https://github.com/zhenghaoz/gorse)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:667** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ryanbressler/CloudForest)   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:663** An evolutionary optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [bayesian](https://github.com/jbrukh/bayesian) **star:653** Naive Bayesian Classification for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [gobrain](https://github.com/goml/gobrain) **star:432** Neural Networks written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [regommend](https://github.com/muesli/regommend) **star:264** Recommendation & collaborative filtering engine.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:241** Go Interface to Open Neural Network Exchange (ONNX).   [![There was an update last week][Green]](https://github.com/owulveryck/onnx-go)   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-galib](https://github.com/thoj/go-galib) **star:178** Genetic Algorithms library written in Go / golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thoj/go-galib)   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:155** Recommendation Algorithms library written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timkaye11/goRecommend)   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [shield](https://github.com/eaigner/shield) **star:129** Bayesian text classifier with flexible tokenizers and storage backends for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eaigner/shield)   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [tfgo](https://github.com/galeone/tfgo)  Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.
* [Goptuna](https://github.com/c-bata/goptuna) **star:116** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [go-fann](https://github.com/white-pony/go-fann) **star:103** Go bindings for Fast Artificial Neural Networks(FANN) library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/white-pony/go-fann)   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [goga](https://github.com/tomcraven/goga) **star:85** Genetic algorithm library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomcraven/goga)   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:64** libsvm golang version derived work based on LIBSVM 3.14.   [![It hasn't been updated in the last year][Yellow]](https://github.com/datastream/libsvm)   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [neural-go](https://github.com/schuyler/neural-go) **star:61** Multilayer perceptron network implemented in Go, with training via backpropagation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schuyler/neural-go)   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [ocrserver](https://github.com/otiai10/ocrserver)  A simple OCR API server, seriously easy to be deployed by Docker and Heroku.
* [go-pr](https://github.com/daviddengcn/go-pr) **star:59** Pattern recognition package in Go lang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-pr)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [gonet](https://github.com/dathoangnd/gonet) **star:57** Neural Network for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/dathoangnd/gonet)
* [neat](https://github.com/jinyeom/neat) **star:56** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/jinyeom/neat)   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   [![Archived][Archived]](https://github.com/jinyeom/neat)
* [goscore](https://github.com/asafschers/goscore) **star:44** Go Scoring API for PMML.   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [GoMind](https://github.com/surenderthakran/gomind)  A simplistic Neural Network Library in Go.
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/golinear)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [Varis](https://github.com/Xamber/Varis) **star:28** Golang Neural Network.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Xamber/Varis)   [![godoc][GoDoc]](https://godoc.org/github.com/Xamber/Varis)
* [godist](https://github.com/e-dard/godist) **star:25** Various probability distributions, and associated methods.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-dard/godist)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/godist)
* [go-deep](https://github.com/patrikeh/go-deep)  A feature-rich neural network library in Go.
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:21** Go implementation of the k-modes and k-prototypes clustering algorithms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
* [probab](https://github.com/ThePaw/probab) **star:11** Probability distribution functions. Bayesian inference. Written in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/probab)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/probab)
* [fonet](https://github.com/Fontinalis/fonet)  A Deep Neural Network library written in Go.
* [evoli](https://github.com/khezen/evoli) **star:9** Genetic Algorithm and Particle Swarm Optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)
* [randomforest](https://github.com/malaschitz/randomForest) **star:3** Easy to use Random Forest library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/malaschitz/randomForest)

## Messaging

*Libraries that implement messaging systems.*

* [gorush](https://github.com/appleboy/gorush) **star:4186** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![There was an update last week][Green]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [mangos](https://github.com/go-mangos/mangos)  Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.
* [machinery](https://github.com/RichardKnop/machinery) **star:3859** Asynchronous task queue/job queue based on distributed message passing.   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   [![There was an update last week][Green]](https://github.com/RichardKnop/machinery)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [NATS Go Client](https://github.com/nats-io/nats)  Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.
* [messagebus](https://github.com/vardius/message-bus)  messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
* [Mercure](https://github.com/dunglas/mercure) **star:1870** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   [![There was an update last week][Green]](https://github.com/dunglas/mercure)   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [melody](https://github.com/olahol/melody) **star:1747** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1591** the official Go package for NSQ.   [![There was an update last week][Green]](https://github.com/nsqio/go-nsq)   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [go-socket.io](https://github.com/googollee/go-socket.io)  socket.io library for golang, a realtime application framework.
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1136** Redis backed unified push service for server-side notifications to mobile devices.   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [zmq4](https://github.com/pebbe/zmq4)  Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).
* [Gollum](https://github.com/trivago/gollum) **star:832** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [Beaver](https://github.com/Clivern/Beaver) **star:808** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![There was an update last week][Green]](https://github.com/Clivern/Beaver)   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [Benthos](https://github.com/Jeffail/benthos)  A message streaming bridge between a range of protocols.
* [EventBus](https://github.com/asaskevich/EventBus) **star:649** The lightweight event bus with async compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [gaurun-client](https://github.com/osamingo/gaurun-client)  Gaurun Client written in Go.
* [golongpoll](https://github.com/jcuga/golongpoll) **star:445** HTTP longpoll server library that makes web pub-sub simple.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jcuga/golongpoll)   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster)  gopush-cluster is a go push server cluster.
* [drone-line](https://github.com/appleboy/drone-line)  Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.
* [dbus](https://github.com/godbus/dbus) **star:424** Native Go bindings for D-Bus.   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [emitter](https://github.com/olebedev/emitter) **star:338** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [Glue](https://github.com/desertbit/glue) **star:331** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:308** Simple pubsub package for go.   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [guble](https://github.com/smancke/guble) **star:142** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   [![It hasn't been updated in the last year][Yellow]](https://github.com/smancke/guble)   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [Centrifugo](https://github.com/centrifugal/centrifugo)  Real-time messaging (Websockets or SockJS) server in Go.
* [Bus](https://github.com/mustafaturan/bus) **star:138** Minimalist message bus implementation for internal communication.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [rabtap](https://github.com/jandelgado/rabtap) **star:105** RabbitMQ swiss army knife cli app.   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [oplog](https://github.com/dailymotion/oplog) **star:102** Generic oplog/replication system for REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dailymotion/oplog)   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:69** A tiny wrapper over amqp exchanges and queues.   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:58** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sybrexsys/RapidMQ)   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [redisqueue](https://github.com/robinjoseph08/redisqueue)  redisqueue provides a producer and consumer of a queue that uses Redis streams.
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:57** A tiny wrapper around NSQ topic and channel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:48** Native implementation of the freedesktop notification spec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TheCreeper/go-notify)   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [hub](https://github.com/leandro-lugaresi/hub) **star:45** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [Commander](https://github.com/jeroenrinzema/commander) **star:39** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [event](https://github.com/agoalofalife/event) **star:33** Implementation of the pattern observer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agoalofalife/event)   [![godoc][GoDoc]](https://godoc.org/github.com/agoalofalife/event)
* [ami](https://github.com/kak-tus/ami)  Go client to reliable queues based on Redis Cluster Streams.
* [APNs2](https://github.com/sideshow/apns2)  HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:12** Client library to Viessmann Vitotrol web service.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [jazz](https://github.com/socifi/jazz) **star:9** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [rmqconn](https://github.com/sbabiv/rmqconn) **star:3** RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/rmqconn)
* [sarama](https://github.com/Shopify/sarama)  Go library for Apache Kafka.

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:2101** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![star > 2000][Awesome]](https://github.com/unidoc/unioffice)   [![There was an update last week][Green]](https://github.com/unidoc/unioffice)   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:5729** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   [![There was an update last week][Green]](https://github.com/360EntSecGroup-Skylar/excelize)   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:3902** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:106** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:66** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter)  Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:1268** A reflection based dependency injection toolkit for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:1173** A dependency injection based application framework for Go (built on top of dig).   [![There was an update last week][Green]](https://github.com/uber-go/fx)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [alice](https://github.com/magic003/alice)  Additive dependency injection container for Golang.
* [container](https://github.com/golobby/container) **star:76** A powerful IoC Container with fluent and easy-to-use interface.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/container)
* [inject](https://github.com/defval/inject) **star:53** A reflection based dependency injection container with simple interface.   [![There was an update last week][Green]](https://github.com/defval/inject)   [![godoc][GoDoc]](https://godoc.org/github.com/defval/inject)
* [dingo](https://github.com/i-love-flamingo/dingo) **star:40** A dependency injection toolkit for Go, based on Guice.   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/dingo)
* [linker](https://github.com/logrange/linker) **star:21** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [wire](https://github.com/Fs02/wire)  Strict Runtime Dependency Injection for Golang.
* [gocontainer](https://github.com/vardius/gocontainer) **star:12** Simple Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:13244** Set of common historical and emerging project layout patterns in the Go ecosystem.   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:538** Go application boilerplate and example applying modern practices.   [![godoc][GoDoc]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) **star:317** A Go application boilerplate template for quick starting projects following production best practices.   [![godoc][GoDoc]](https://godoc.org/github.com/lacion/cookiecutter-golang)
* [scaffold](https://github.com/catchplay/scaffold) **star:55** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.   [![It hasn't been updated in the last year][Yellow]](https://github.com/catchplay/scaffold)   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [go-sample](https://github.com/zitryss/go-sample) **star:45** A sample layout for Go application projects with the real code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zitryss/go-sample)   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:720** Collection of useful string functions ported from other languages.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:84** String utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4549** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![There was an update last week][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:2735** Library and command for making and extracting .zip and .tar.gz archives.   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![There was an update last week][Green]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gosms](https://github.com/haxpax/gosms) **star:1268** Your own local SMS gateway in Go that can be used to send SMS.   [![It hasn't been updated in the last year][Yellow]](https://github.com/haxpax/gosms)   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:973** Resiliency patterns for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:825** Random data generator written in go.   [![There was an update last week][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:769** Generic object pool for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![Contains Chinese documents][CN]](https://github.com/jolestar/go-commons-pool)
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:766** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![Contains Chinese documents][CN]](https://github.com/mojocn/base64Captcha)
* [shortid](https://github.com/teris-io/shortid) **star:524** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/shortid)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:495** Library for interacting with LLVM IR in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:380** Easy to use, extensible health check library.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:350** Package conv provides fast and intuitive conversions across Go types.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [banner](https://github.com/dimiro1/banner) **star:253** Add beautiful banners into your Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:232** Package that exposes country and subdivision data.   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [antch](https://github.com/antchfx/antch) **star:169** A fast, powerful and extensible web crawling & scraping framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/antch)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   [![Contains Chinese documents][CN]](https://github.com/antchfx/antch)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:152** Beautify data display for Humans.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   [![Contains Chinese documents][CN]](https://github.com/go-ffmt/ffmt)
* [stateless](https://github.com/qmuntal/stateless) **star:148** A fluent library for creating state machines.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [battery](https://github.com/distatus/battery) **star:143** Cross-platform, normalized battery information library.   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [lk](https://github.com/hyperboloide/lk) **star:143** A simple licensing library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [stats](https://github.com/go-playground/stats) **star:131** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [ghorg](https://github.com/gabrie30/ghorg) **star:129** Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [bitio](https://github.com/icza/bitio) **star:114** Highly optimized bit-level Reader and Writer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:105** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:94** Emojis for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:86** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [gommit](https://github.com/antham/gommit) **star:85** Analyze git commit messages to ensure they follow defined patterns.   [![There was an update last week][Green]](https://github.com/antham/gommit)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [gotoprom](https://github.com/cabify/gotoprom) **star:81** Type-safe metrics builder wrapper library for the official Prometheus client.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [indigo](https://github.com/osamingo/indigo) **star:58** Distributed unique ID generator of using Sonyflake and encoded by Base58.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [morse](https://github.com/alwindoss/morse) **star:56** Library to convert to and from morse code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alwindoss/morse)   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [captcha](https://github.com/steambap/captcha) **star:53** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![There was an update last week][Green]](https://github.com/steambap/captcha)   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [xkg](https://github.com/go-xkg/xkg) **star:48** X Keyboard Grabber.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [persian](https://github.com/mavihq/persian) **star:40** Some utilities for Persian language in go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mavihq/persian)   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:38** HTTP service to generate PDF from Json requests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [datacounter](https://github.com/miolini/datacounter) **star:31** Go counters for readers/writer/http.ResponseWriter.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:31** GoLang Library for [Browser Capabilities Project](http://browscap.org/).   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [autoflags](https://github.com/artyom/autoflags) **star:26** Go package to automatically define command line flags from struct fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/artyom/autoflags)   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg) Specs implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [gosh](https://github.com/osamingo/gosh) **star:20** Provide Go Statistics Handler, Struct, Measure Method.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:19** A modern, powerful, and robust URL shortener microservice with mysql support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [sandid](https://github.com/aofei/sandid) **star:14** Every grain of sand on earth has its own ID.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [anagent](https://github.com/mudler/anagent) **star:11** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mudler/anagent)   [![godoc][GoDoc]](https://godoc.org/github.com/mudler/anagent)
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:10** Calculate average score and rating based on Wilson Score Equation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/avgRating)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/avgRating)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/hostutils)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [shellwords](https://github.com/Wing924/shellwords) **star:8** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [metrics](https://github.com/pascaldekloe/metrics) **star:6** Library for metrics instrumentation and Prometheus exposition.   [![There was an update last week][Green]](https://github.com/pascaldekloe/metrics)   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2435** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only.   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1208** Go efficient text segmentation; support english, chinese, japanese and other.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![Contains Chinese documents][CN]](https://github.com/go-ego/gse)
* [when](https://github.com/olebedev/when) **star:1060** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:983** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   [![There was an update last week][Green]](https://github.com/yanyiwu/gojieba)   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![Contains Chinese documents][CN]](https://github.com/yanyiwu/gojieba)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:647** CN Hanzi to Hanyu Pinyin converter.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:472** JP morphological analyzer written in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:397** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:361** Extract values from strings and fill your structs with nlp.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:266** Sentence tokenizer:  converts text into a list of sentences.   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:241** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [getlang](https://github.com/rylans/getlang) **star:83** Fast natural language detection package.   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-nlp](https://github.com/nuance/go-nlp) **star:82** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [gounidecode](https://github.com/fiam/gounidecode) **star:68** Unicode transliterator (also known as unidecode) for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:67** ASCII transliterations of Unicode text.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [textcat](https://github.com/pebbe/textcat) **star:62** Go package for n-gram based text categorization, with support for utf-8 and raw text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:56** Implementation of the porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:55** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [stemmer](https://github.com/dchest/stemmer) **star:49** Stemmer packages for Go programming language. Includes English and German stemmers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [segment](https://github.com/blevesearch/segment) **star:47** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![It hasn't been updated in the last year][Yellow]](https://github.com/blevesearch/segment)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [porter2](https://github.com/zhenjl/porter2) **star:37** Really fast Porter 2 stemmer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/porter2)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [go2vec](https://github.com/danieldk/go2vec) **star:33** Reader and utility functions for word2vec embeddings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [petrovich](https://github.com/striker2000/petrovich) **star:28** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang implementation of the Paice/Husk Stemming Algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [snowball](https://github.com/goodsign/snowball) **star:25** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/snowball)
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo bindings to Yandex.Mystem - russian morphology analyzer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:17** Go bindings for the snowball libstemmer library including porter 2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [shamoji](https://github.com/osamingo/shamoji) **star:11** The shamoji is word filtering package written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)
* [go-localize](https://github.com/m1/go-localize) **star:3** Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-localize)
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:1** Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.   [![godoc][GoDoc]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)

## Networking

*Libraries for working with various layers of the network.*

* [fasthttp](https://github.com/valyala/fasthttp) **star:11576** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [dns](https://github.com/miekg/dns) **star:4335** Go library for working with DNS.   [![star > 2000][Awesome]](https://github.com/miekg/dns)   [![There was an update last week][Green]](https://github.com/miekg/dns)   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [gopacket](https://github.com/google/gopacket) **star:3256** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   [![There was an update last week][Green]](https://github.com/google/gopacket)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [webrtc](https://github.com/pions/webrtc) **star:3113** A pure Go implementation of the WebRTC API.   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   [![There was an update last week][Green]](https://github.com/pions/webrtc)   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2463** KCP - Fast and Reliable ARQ Protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   [![There was an update last week][Green]](https://github.com/xtaci/kcp-go)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [kcptun](https://github.com/xtaci/kcptun)  Extremely simple & fast udp tunnel based on KCP protocol.
* [gobgp](https://github.com/osrg/gobgp) **star:1813** BGP implemented in the Go Programming Language.   [![There was an update last week][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [gnet](https://github.com/panjf2000/gnet) **star:1692** `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.   [![There was an update last week][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/gnet)
* [fortio](https://github.com/fortio/fortio) **star:1155** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![There was an update last week][Green]](https://github.com/fortio/fortio)   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:959** Simple TUN/TAP library.   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [go-getter](https://github.com/hashicorp/go-getter) **star:854** Go library for downloading files or directories from various sources using a URL.   [![There was an update last week][Green]](https://github.com/hashicorp/go-getter)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:766** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [graval](https://github.com/koofr/graval)  Experimental FTP server framework.
* [grab](https://github.com/cavaliercoder/grab) **star:639** Go package for managing file downloads.   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [HTTPLab](https://github.com/gchaincl/httplab)  HTTPLabs let you inspect HTTP requests and forge responses.
* [ftp](https://github.com/jlaffaye/ftp) **star:618** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [lhttp](https://github.com/fanux/lhttp) **star:547** Powerful websocket framework, build your IM server more easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fanux/lhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   [![Contains Chinese documents][CN]](https://github.com/fanux/lhttp)
* [gosnmp](https://github.com/soniah/gosnmp) **star:491** Native Go library for performing SNMP actions.   [![There was an update last week][Green]](https://github.com/soniah/gosnmp)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [dhcp6](https://github.com/mdlayher/dhcp6)  Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
* [cidranger](https://github.com/yl2chen/cidranger) **star:441** Fast IP to CIDR lookup for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:402** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [quic-go](https://github.com/lucas-clemente/quic-go)  An implementation of the QUIC protocol in pure Go.
* [publicip](https://github.com/polera/publicip)  Package publicip returns your public facing IPv4 address (internet egress).
* [portproxy](https://github.com/aybabtme/portproxy)  Simple TCP proxy which adds CORS support to API's which don't support it.
* [gopcap](https://github.com/akrennmair/gopcap) **star:376** Go wrapper for libpcap.   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [goshark](https://github.com/sunwxg/goshark)  Package goshark use tshark to decode IP packet and create data struct to analyse packet.
* [go-stun](https://github.com/ccding/go-stun) **star:344** Go implementation of the STUN client (RFC 3489 and RFC 5389).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ccding/go-stun)   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [sftp](https://github.com/pkg/sftp)  Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.
* [ssh](https://github.com/gliderlabs/ssh)  Higher-level API for building SSH servers (wraps crypto/ssh).
* [raw](https://github.com/mdlayher/raw) **star:342** Package raw enables reading and writing data at the device driver level for a network interface.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [stun](https://github.com/go-rtc/stun) **star:339** Go implementation of RFC 5389 STUN protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:315** Go library for building tcp servers faster.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [winrm](https://github.com/masterzen/winrm) **star:239** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:235** Streaming protocolbuffer data over TCP made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stabbycutyou/buffstreams)   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [arp](https://github.com/mdlayher/arp) **star:214** Package arp implements the ARP protocol, as described in RFC 826.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:196** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [jazigo](https://github.com/udhos/jazigo) **star:144** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [canopus](https://github.com/zubairhamed/canopus) **star:137** CoAP Client/Server implementation (RFC 7252).   [![It hasn't been updated in the last year][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:130** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![There was an update last week][Green]](https://github.com/DrmagicE/gmqtt)   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![Contains Chinese documents][CN]](https://github.com/DrmagicE/gmqtt)
* [sslb](https://github.com/eduardonunesp/sslb) **star:124** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gNxI](https://github.com/google/gnxi) **star:117** A collection of tools for Network Management that use the gNMI and gNOI protocols.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [gaio](https://github.com/xtaci/gaio) **star:101** High performance async-io networking for Golang in proactor mode.   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gaio)
* [gev](https://github.com/Allenxuxu/gev)  gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.
* [xtcp](https://github.com/xfxdev/xtcp) **star:93** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.   [![There was an update last week][Green]](https://github.com/xfxdev/xtcp)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [ether](https://github.com/songgao/ether) **star:65** Cross-platform Go package for sending and receiving ethernet frames.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [linkio](https://github.com/ian-kent/linkio) **star:47** Network link speed simulation for Reader/Writer interfaces.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/linkio)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/linkio)
* [packet](https://github.com/aerogo/packet) **star:43** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [iplib](https://github.com/c-robinson/iplib) **star:28** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:17** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/golibwireshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/golibwireshark)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [mdns](https://github.com/hashicorp/mdns)  Simple mDNS (Multicast DNS) client/server library in Golang.
* [llb](https://github.com/kirillDanshin/llb) **star:10** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/llb)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/llb)
* [go-powerdns](https://github.com/joeig/go-powerdns) **star:10** PowerDNS API bindings for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/joeig/go-powerdns)
* [tspool](https://github.com/two/tspool) **star:6** A TCP Library use worker pool to improve performance and protect your server.   [![It hasn't been updated in the last year][Yellow]](https://github.com/two/tspool)   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)
* [utp](https://github.com/anacrolix/utp)  Go uTP micro transport protocol implementation.
* [gosocsvr](https://github.com/rakeki/gosocsvr) **star:5** Socket server made simple.   [![godoc][GoDoc]](https://godoc.org/github.com/rakeki/gosocsvr)
* [gotcp](https://github.com/gansidui/gotcp)  Go package for quickly writing tcp applications.
* [httpproxy](https://github.com/wzshiming/httpproxy) **star:2** HTTP proxy handler and dialer.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/httpproxy)

### HTTP Clients

*Libraries for making HTTP requests.*

* [rq](https://github.com/ddo/rq)  A nicer interface for golang stdlib HTTP client.
* [resty](https://github.com/go-resty/resty) **star:2567** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   [![There was an update last week][Green]](https://github.com/go-resty/resty)   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [sling](https://github.com/dghubble/sling)  Sling is a Go HTTP client library for creating and sending API requests.
* [grequests](https://github.com/levigross/grequests) **star:1523** A Go "clone" of the great and famous Requests library.   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [heimdall](https://github.com/gojektech/heimdall)  An enchanced http client with retry and hystrix capabilities.
* [gentleman](https://github.com/h2non/gentleman)  Full-featured plugin-driven HTTP client library.
* [httpretry](https://github.com/ybbus/httpretry) **star:2** Enriches the default go HTTP client with retry functionality.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/httpretry)
* [pester](https://github.com/sethgrid/pester)  Go HTTP client calls with retries, backoff, and concurrency.

## OpenGL

*Libraries for using OpenGL in Go.*

* [gl](https://github.com/go-gl/gl)  Go bindings for OpenGL (generated via glow).
* [glfw](https://github.com/go-gl/glfw) **star:863** Go bindings for GLFW 3.   [![There was an update last week][Green]](https://github.com/go-gl/glfw)
* [goxjs/gl](https://github.com/goxjs/gl)  Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:60** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)
* [mathgl](https://github.com/go-gl/mathgl)  Pure Go math package specialized for 3D math, with inspiration from GLM.

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) **star:17257** The fantastic ORM library for Golang, aims to be developer friendly.   [![star > 2000][Awesome]](https://github.com/jinzhu/gorm)   [![There was an update last week][Green]](https://github.com/jinzhu/gorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/gorm)
* [go-pg](https://github.com/go-pg/pg) **star:3549** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   [![There was an update last week][Green]](https://github.com/go-pg/pg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [go-queryset](https://github.com/jirfag/go-queryset)  100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder)  A flexible and powerful SQL string builder library plus a zero-config ORM.
* [go-store](https://github.com/gosuri/go-store)  Simple and fast Redis backed key-value store library for Go.
* [upper.io/db](https://github.com/upper/db) **star:2081** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![star > 2000][Awesome]](https://github.com/upper/db)   [![There was an update last week][Green]](https://github.com/upper/db)   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [Xorm](https://github.com/go-xorm/xorm)  Simple and powerful ORM for Go.
* [QBS](https://github.com/coocood/qbs) **star:547** Stands for Query By Struct. A Go ORM.   [![It hasn't been updated in the last year][Yellow]](https://github.com/coocood/qbs)   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   [![Contains Chinese documents][CN]](https://github.com/coocood/qbs)
* [reform](https://github.com/go-reform/reform)  Better ORM for Go, based on non-empty interfaces and code generation.
* [gormt](https://github.com/xxjwxc/gormt) **star:336** Mysql database to golang gorm struct.   [![There was an update last week][Green]](https://github.com/xxjwxc/gormt)   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gormt)
* [gorp](https://github.com/go-gorp/gorp)  Go Relational Persistence, ORM-ish library for Go.
* [grimoire](https://github.com/Fs02/grimoire)  Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).
* [Zoom](https://github.com/albrow/zoom) **star:249** Blazing-fast datastore and querying engine built on Redis.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/zoom)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [Marlow](https://github.com/dadleyy/marlow) **star:75** Generated ORM from project structs for compile time safety assurances.   [![godoc][GoDoc]](https://godoc.org/github.com/dadleyy/marlow)
* [pop/soda](https://github.com/gobuffalo/pop)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [SQLBoiler](https://github.com/volatiletech/sqlboiler)  ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.
* [rel](https://github.com/Fs02/rel) **star:50** Golang SQL Repository Layer for Clean (Onion) Architecture.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/rel)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:10** A simple ORM for Google/Firebase Cloud Firestore.   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)
* [lore](https://github.com/abrahambotros/lore) **star:5** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abrahambotros/lore)   [![godoc][GoDoc]](https://godoc.org/github.com/abrahambotros/lore)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep)  Go dependency tool.
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [godep](https://github.com/tools/godep) **star:5646** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   [![star > 2000][Awesome]](https://github.com/tools/godep)   [![It hasn't been updated in the last year][Yellow]](https://github.com/tools/godep)   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   [![Archived][Archived]](https://github.com/tools/godep)
* [gom](https://github.com/mattn/gom) **star:1377** Go Manager - bundle for go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [goop](https://github.com/nitrous-io/goop)  Simple dependency manager for Go (golang), inspired by Bundler.
* [gop](https://github.com/lunny/gop)  Build and manage your Go applications out of GOPATH.
* [gopm](https://github.com/gpmgo/gopm)  Go Package Manager.
* [govendor](https://github.com/kardianos/govendor)  Go Package Manager. Go vendor tool that works with the standard vendor file.
* [gpm](https://github.com/pote/gpm) **star:1204** Barebones dependency manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pote/gpm)
* [gigo](https://github.com/LyricalSecurity/gigo)  PIP-like dependency tool for golang, with support for private repositories and hashes.
* [glide](https://github.com/Masterminds/glide)  Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.
* [nut](https://github.com/jingweno/nut) **star:244** Vendor Go dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jingweno/nut)   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** Minimal dependency version using Git.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/johnny-deps)
* [VenGO](https://github.com/DamnWidget/VenGO) **star:116** create and manage exportable isolated go virtual environments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DamnWidget/VenGO)   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:99** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:10274** A distributed tracing system.   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   [![There was an update last week][Green]](https://github.com/jaegertracing/jaeger)   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1187** Simple profiling support package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [tracer](https://github.com/kamilsk/tracer) **star:20** Simple, lightweight tracing.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:6054** Implementation of GraphQL for Go.   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [jsonql](https://github.com/elgs/jsonql)  JSON query expression library in Golang.
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:1265** A simple Go package to Query over JSON Data.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [graphql](https://github.com/tmc/graphql) **star:52** graphql parser + utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tmc/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [graphql](https://github.com/neelance/graphql-go)  GraphQL server with a focus on ease of use.
* [jsonslice](https://github.com/bhmj/jsonslice) **star:38** Jsonpath queries with advanced filters.   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)
* [rql](https://github.com/a8m/rql)  Resource Query Language for REST API.
* [straf](https://github.com/SonicRoshan/straf)  Easily Convert Golang structs to GraphQL objects.

## Resource Embedding

* [statics](https://github.com/go-playground/statics)  Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.
* [packr](https://github.com/gobuffalo/packr) **star:2629** The simple and easy way to embed static files into Go binaries.   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik)  Embeds static files into a Go executable.
* [templify](https://github.com/wlbr/templify)  Embed external template files into Go code to create single file binaries.
* [vfsgen](https://github.com/shurcooL/vfsgen)  Generates a vfsdata.go file that statically implements the given virtual filesystem.
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1849** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [esc](https://github.com/mjibson/esc) **star:529** Embeds files into Go programs and provides http.FileSystem interfaces to them.   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:502** Simple tool to embed files in go with focus on "customization" and ease to use.   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-embed](https://github.com/pyros2097/go-embed) **star:20** Generates go code to embed resource files into your library or executable.   [![There was an update last week][Green]](https://github.com/pyros2097/go-embed)   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)
* [go-resources](https://github.com/omeid/go-resources)  Unfancy resources embedding with Go.

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:3540** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   [![There was an update last week][Green]](https://github.com/gonum/gonum)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [assocentity](https://github.com/ndabAP/assocentity)  Package assocentity returns the average distance from words to a given entity.
* [stats](https://github.com/montanaflynn/stats) **star:1626** Statistics package with common functions missing from the Golang standard library.   [![There was an update last week][Green]](https://github.com/montanaflynn/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gonum/plot](https://github.com/gonum/plot) **star:1396** gonum/plot provides an API for building and drawing plots in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [gosl](https://github.com/cpmech/gosl) **star:1389** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![godoc][GoDoc]](https://godoc.org/github.com/cpmech/gosl)
* [TextRank](https://github.com/DavidBelicza/TextRank)  TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.
* [streamtools](https://github.com/nytlabs/streamtools) **star:1317** general purpose, graphical tool for dealing with streams of data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nytlabs/streamtools)   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go)  Dataframes for machine-learning and statistics (similar to pandas).
* [chart](https://github.com/vdobler/chart) **star:623** Simple Chart Plotting library for Go. Supports many graphs types.   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [goraph](https://github.com/gyuho/goraph) **star:619** Pure Go graph theory library(data structure, algorith visualization).   [![It hasn't been updated in the last year][Yellow]](https://github.com/gyuho/goraph)   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [graph](https://github.com/yourbasic/graph) **star:302** Library of basic graph algorithms.   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [ewma](https://github.com/VividCortex/ewma) **star:288** Exponentially-weighted moving averages.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [orb](https://github.com/paulmach/orb) **star:261** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:138** Approximate histograms for data streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/gohistogram)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [sparse](https://github.com/james-bowman/sparse) **star:82** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [pagerank](https://github.com/alixaxel/pagerank) **star:54** Weighted PageRank algorithm implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear)  Tiny linear interpolation library.
* [geom](https://github.com/skelterjohn/geom) **star:43** 2D geometry for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/geom)   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [go-dsp](https://github.com/mjibson/go-dsp)  Digital Signal Processing for Go.
* [evaler](https://github.com/soniah/evaler) **star:39** Simple floating point arithmetic expression evaluator.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/evaler)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:16** GO Implementation of Entropy Measures.   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:12** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [ode](https://github.com/ChristopherRabotin/ode) **star:11** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/ode)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/ode)
* [GoStats](https://github.com/OGFris/GoStats) **star:10** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OGFris/GoStats)   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [PiHex](https://github.com/claygod/PiHex) **star:8** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** Graph theory algorithms written in "Go" language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/go-gt)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/go-gt)
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:2** Provides a Bradley-Terry Model for pairwise comparisons.   [![godoc][GoDoc]](https://godoc.org/github.com/seanhagen/bradleyterry)

## Security

*Libraries that are used to help make your application more secure.*

* [Cameradar](https://github.com/Ullaakut/cameradar) **star:2120** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![star > 2000][Awesome]](https://github.com/Ullaakut/cameradar)   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [acra](https://github.com/cossacklabs/acra) **star:534** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   [![There was an update last week][Green]](https://github.com/cossacklabs/acra)   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [sslmgr](https://github.com/adrianosela/sslmgr)  SSL certificates made easy with a high level wrapper around acme/autocert.
* [ssh-vault](https://github.com/ssh-vault/ssh-vault)  encrypt/decrypt using ssh keys.
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:163** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [go-yara](https://github.com/hillu/go-yara) **star:150** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [argon2pw](https://github.com/raja/argon2pw) **star:80** Argon2 password hash generation with constant-time password comparison.   [![It hasn't been updated in the last year][Yellow]](https://github.com/raja/argon2pw)   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [BadActor](https://github.com/jaredfolkins/badactor)  In-memory, application-driven jailer built in the spirit of fail2ban.
* [nacl](https://github.com/kevinburke/nacl)  Go implementation of the NaCL set of API's.
* [secure](https://github.com/unrolled/secure)  HTTP middleware for Go that facilitates some quick security wins.
* [passlib](https://github.com/hlandau/passlib)  Futureproof password hashing library.
* [acmetool](https://github.com/hlandau/acme)  ACME (Let's Encrypt) client tool with automatic renewal.
* [memguard](https://github.com/awnumar/memguard)  A pure Go library for handling sensitive values in memory.
* [lego](https://github.com/xenolf/lego)  Pure Go ACME client library and CLI tool (for use with Let's Encrypt).
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:12** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword)  A probably paranoid package for securely hashing and encrypting passwords.
* [go-generate-password](https://github.com/m1/go-generate-password)  Password generator that can be used on the cli or as a library.
* [certificates](https://github.com/mvmaasakkers/certificates) **star:11** An opinionated tool for generating tls certificates.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:7198** High-performance 100% compatible drop-in replacement of "encoding/json".   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![There was an update last week][Green]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:6203** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   [![There was an update last week][Green]](https://github.com/golang/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:3522** Protocol Buffers for Go with Gadgets.   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [pletter](https://github.com/vimeda/pletter)  A standard way to wrap a proto message for message brokers.
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder)  GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:3077** Go library for decoding generic map values into native Go structures.   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1358** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [csvutil](https://github.com/jszwec/csvutil) **star:350** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![There was an update last week][Green]](https://github.com/jszwec/csvutil)   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [fixedwidth](https://github.com/huydang284/fixedwidth)  Fixed-width text formatting (UTF-8 supported).
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:276** Cap'n Proto library and parser for go.   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [asn1](https://github.com/PromonLogicalis/asn1)  Asn.1 BER and DER encoding library for golang.
* [bambam](https://github.com/glycerine/bambam)  generator for Cap'n Proto schemas from go.
* [bel](https://github.com/32leaves/bel)  Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.
* [structomap](https://github.com/tuvistavie/structomap) **star:109** Library to easily and dynamically generate maps from static structures.   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [colfer](https://github.com/pascaldekloe/colfer)  Code generation for the Colfer binary format.
* [cbor](https://github.com/fxamacker/cbor)  Small, safe, and easy CBOR encoding and decoding library.
* [binstruct](https://github.com/ghostiam/binstruct) **star:17** Golang binary decoder for mapping data into the structure.   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [fwencoder](https://github.com/o1egl/fwencoder) **star:10** Fixed width file parser (encoding and decoding library) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/fwencoder)

## Server Applications

* [minio](https://github.com/minio/minio) **star:20385** Minio is a distributed object storage server.   [![star > 2000][Awesome]](https://github.com/minio/minio)   [![There was an update last week][Green]](https://github.com/minio/minio)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [devd](https://github.com/cortesi/devd) **star:2947** Local webserver for developers.   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [flipt](https://github.com/markphelps/flipt) **star:1166** A self contained feature flag solution written in Go and Vue.js   [![There was an update last week][Green]](https://github.com/markphelps/flipt)   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Fider](https://github.com/getfider/fider) **star:1018** Fider is an open platform to collect and organize customer feedback.   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [Flagr](https://github.com/checkr/flagr) **star:993** Flagr is an open-source feature flagging and A/B testing service.   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [discovery](https://github.com/Bilibili/discovery) **star:854** A registry for resilient mid-tier load balancing and failover.   [![There was an update last week][Green]](https://github.com/Bilibili/discovery)   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [jackal](https://github.com/ortuman/jackal) **star:790** An XMPP server written in Go.   [![There was an update last week][Green]](https://github.com/ortuman/jackal)   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [Caddy](https://github.com/mholt/caddy)  Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
* [algernon](https://github.com/xyproto/algernon)  HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [etcd](https://github.com/coreos/etcd)  Highly-available key value store for shared configuration and service discovery.
* [dudeldu](https://github.com/krotik/dudeldu) **star:108** A simple SHOUTcast server.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/dudeldu)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:14** Stream database events from PostgreSQL to Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [RoadRunner](https://github.com/spiral/roadrunner)  High-performance PHP application server, load-balancer and process manager.
* [SFTPGo](https://github.com/drakkan/sftpgo)  Full featured and highly configurable SFTP server software.
* [Trickster](https://github.com/Comcast/trickster)  HTTP reverse proxy cache and time series accelerator.
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:9** Nginx log parser and exporter to Prometheus.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.


## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:281** Go stream processing library.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3473** PDF document generator with high level support for text, drawing and images.   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)   [![Archived][Archived]](https://github.com/jung-kurt/gofpdf)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1636** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![There was an update last week][Green]](https://github.com/valyala/quicktemplate)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [pongo2](https://github.com/flosch/pongo2) **star:1614** Django-like template-engine for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [mustache](https://github.com/hoisie/mustache) **star:980** Go implementation of the Mustache template language.   [![There was an update last week][Green]](https://github.com/hoisie/mustache)   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:840** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eknkc/amber)   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:781** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/ace)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:730** Razor view engine for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [ego](https://github.com/benbjohnson/ego) **star:428** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [raymond](https://github.com/aymerick/raymond) **star:360** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/raymond)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [Soy](https://github.com/robfig/soy) **star:147** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [maroto](https://github.com/johnfercher/maroto) **star:116** A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.   [![There was an update last week][Green]](https://github.com/johnfercher/maroto)   [![godoc][GoDoc]](https://godoc.org/github.com/johnfercher/maroto)
* [velvet](https://github.com/gobuffalo/velvet) **star:71** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gobuffalo/velvet)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)
* [liquid](https://github.com/osteele/liquid)  Go implementation of Shopify Liquid templates.
* [kasia.go](https://github.com/ziutek/kasia.go) **star:71** Templating system for HTML and other text documents - go implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/kasia.go)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [damsel](https://github.com/dskinner/damsel) **star:22** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dskinner/damsel)   [![godoc][GoDoc]](https://godoc.org/github.com/dskinner/damsel)
* [fasttemplate](https://github.com/valyala/fasttemplate)  Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:19** Tiny wrapper around html/template to allow for easy file-based template inheritance.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dannyvankooten/extemplate)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/extemplate)
* [hero](https://github.com/shiyanhui/hero)  Hero is a handy, fast and powerful go template engine.
* [goview](https://github.com/foolin/goview)  Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
* [gospin](https://github.com/m1/gospin) **star:7** Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/gospin)
* [jet](https://github.com/CloudyKit/jet)  Jet template engine.

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [go-cmp](https://github.com/google/go-cmp) **star:1523** Package for comparing Go values in tests.   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [baloo](https://github.com/h2non/baloo) **star:674** Expressive and versatile end-to-end HTTP API testing made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/h2non/baloo)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [goblin](https://github.com/franela/goblin) **star:658** Mocha like testing framework fo Go.   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [Testify](https://github.com/stretchr/testify)  Sacred extension to the standard go testing package.
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:442** A helper for Rails' like test fixtures to test database applications.   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:384** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:328** Mutation testing for Go source code.   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gogiven](https://github.com/corbym/gogiven)  YATSPEC-like BDD testing framework for Go.
    * [gofight](https://github.com/appleboy/gofight) **star:300** API Handler Testing for Golang Router framework.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [frisby](https://github.com/verdverm/frisby) **star:255** REST API testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:202** Tool for viewing test coverage in terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [endly](https://github.com/viant/endly) **star:129** Declarative end to end functional testing.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:114** BDD-style testing framework for the Go programming language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/orfjackal/gospec)   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:100** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:99** Simple snapshot testing addon for your test framework.   [![There was an update last week][Green]](https://github.com/bradleyjkemp/cupaloy)   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [commander](https://github.com/SimonBaeumer/commander) **star:87** Tool for testing cli applications on windows, linux and osx.   [![There was an update last week][Green]](https://github.com/SimonBaeumer/commander)   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:80** Extremely flexible golang deep comparison, extends the go testing package.   [![There was an update last week][Green]](https://github.com/maxatome/go-testdeep)   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [wstest](https://github.com/posener/wstest) **star:72** Websocket client for unit-testing a websocket http.Handler.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [assert](https://github.com/go-playground/assert)  Basic Assertion Library used along side native go testing, with building blocks for custom assertions.
    * [gospecify](https://github.com/stesla/gospecify) **star:53** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stesla/gospecify)   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [schema](https://github.com/jgroeneveld/schema)  Quick and easy expression matching for JSON schemas used in requests and responses.
    * [restit](https://github.com/yookoala/restit) **star:50** Go micro framework to help writing RESTful API integration test.   [![godoc][GoDoc]](https://godoc.org/github.com/yookoala/restit)
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:33** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:32** library created for testing JSON against patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres)  Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.
    * [dsunit](https://github.com/viant/dsunit) **star:28** Datastore testing for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [testcase](https://github.com/adamluzsi/testcase) **star:12** Idiomatic testing framework for Behavior Driven Development.   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:11** HTTP client testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Extensions to Go's `testing/iotest` package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cavaliercoder/badio)   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/badio)
    * [httpexpect](https://github.com/gavv/httpexpect)  Concise, declarative, and easy to use end-to-end HTTP and REST API testing.
    * [godog](https://github.com/DATA-DOG/godog)  Cucumber or Behat like BDD framework for Go.
    * [Hamcrest](https://github.com/rdrdr/hamcrest)  fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools)  A collection of packages to augment the go testing package and support common patterns.
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pavlo/gosuite)   [![godoc][GoDoc]](https://godoc.org/github.com/pavlo/gosuite)
    * [gocrest](https://github.com/corbym/gocrest) **star:9** Composable hamcrest-like matchers for Go assertions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gocrest)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gocrest)
    * [charlatan](https://github.com/percolate/charlatan)  Tool to generate fake interface implementations for tests.
    * [biff](https://github.com/fulldump/biff) **star:8** Bifurcation testing framework, BDD compatible.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fulldump/biff)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/biff)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [trial](https://github.com/jgroeneveld/trial) **star:4** Quick and easy extendable assertions without introducing much boilerplate.   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/trial)
    * [Tt](https://github.com/vcaesar/tt)  Simple and colorful test tools.

* Mock
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1524** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:956** Versatile HTTP mocking made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [gomock](https://github.com/golang/mock)  Mocking framework for the Go programming language.
    * [httpmock](https://github.com/jarcoal/httpmock) **star:707** Easy mocking of HTTP responses from external resources.   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:397** Tool for generating self-contained mock objects.   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)  Mock SQL driver for testing database interactions.
    * [minimock](https://github.com/gojuno/minimock) **star:294** Mock generator for Go interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:233** Single transaction based database driver mainly for testing purposes.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:88** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)
    * [timex](https://github.com/cabify/timex) **star:27** A test-friendly replacement for the native `time` package.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/timex)
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tv42/mockhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/tv42/mockhttp)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:3317** Randomized testing system.   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz)  Library for populating go objects with random values.
    * [Tavor](https://github.com/zimmski/tavor) **star:221** Generic fuzzing and delta-debugging framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zimmski/tavor)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:4347** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1500** alternative Selenium hub server that launches browsers within containers.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [cdp](https://github.com/mafredri/cdp) **star:408** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:237** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:480** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:10137** Fast and Elegant Scraping Framework for Gophers.   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![There was an update last week][Green]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:8415** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![There was an update last week][Green]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [ltsv](https://github.com/Wing924/ltsv)  High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.
    * [goribot](https://github.com/zhshch2002/goribot)  A simple golang spider/scraping framework,build a spider in 3 lines.
    * [gotext](https://github.com/leonelquinteros/gotext)  GNU gettext utilities for Go.
    * [guesslanguage](https://github.com/endeveit/guesslanguage)  Functions to determine the natural language of a unicode text.
    * [goregen](https://github.com/zach-klippenstein/goregen)  Library for generating random strings from regular expressions.
    * [inject](https://github.com/facebookgo/inject)  Package inject provides a reflect based injector.
    * [htmlquery](https://github.com/antchfx/htmlquery)  An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.
    * [blackfriday](https://github.com/russross/blackfriday) **star:4253** Markdown processor in Go.   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![There was an update last week][Green]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:2106** Formatters for time, numbers, and memory size to human readable format.   [![star > 2000][Awesome]](https://github.com/dustin/go-humanize)   [![There was an update last week][Green]](https://github.com/dustin/go-humanize)   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1455** HTML Sanitizer.   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1195** Parse RSS and Atom feeds in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [go-toml](https://github.com/pelletier/go-toml) **star:702** Go library for the TOML format with query support and handy cli tools.   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:583** A collection of common regular expressions for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [mxj](https://github.com/clbanning/mxj) **star:362** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:362** Web scraping Framework to turn websites into structured data.   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [goq](https://github.com/andrewstuart/goq)  Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).
    * [gonameparts](https://github.com/polera/gonameparts)  Parses human names into individual name parts.
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:344** Parses the Graphviz DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [go-runewidth](https://github.com/mattn/go-runewidth)  Functions to get fixed width of the character or string.
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:112** NMEA parser library for the Go language.   [![There was an update last week][Green]](https://github.com/adrianmo/go-nmea)   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [go-slugify](https://github.com/mozillazg/go-slugify)  Make pretty slug with multiple languages support.
    * [slug](https://github.com/gosimple/slug)  URL-friendly slugify with multiple languages support.
    * [syndfeed](https://github.com/zhengchun/syndfeed)  A syndication feed for Atom 1.0 and RSS 2.0.
    * [sh](https://github.com/mvdan/sh)  Shell parser and formatter.
    * [Slugify](https://github.com/avelino/slugify)  Go slugify application that handles string.
    * [sdp](https://github.com/gortc/sdp) **star:85** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [toml](https://github.com/BurntSushi/toml)  TOML configuration format (encoder/decoder with reflection).
    * [align](https://github.com/Guitarbum722/align) **star:63** A general purpose application that aligns text.   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [genex](https://github.com/alixaxel/genex) **star:58** Count and expand Regular Expressions into all matching Strings.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [podcast](https://github.com/eduncan911/podcast) **star:58** iTunes Compliant and RSS 2.0 Podcast Generator in Golang   [![godoc][GoDoc]](https://godoc.org/github.com/eduncan911/podcast)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:56** Editorconfig file parser and manipulator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [allot](https://github.com/sbstjn/allot) **star:38** Placeholder and wildcard text parsing for CLI tools and bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [bbConvert](https://github.com/CalebQ42/bbConvert)  Converts bbCode to HTML that allows you to add support for custom bbCode tags.
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width)  Zero-width character detection and removal for Go.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:33** Fixed-width text formatting (encoder/decoder with reflection).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [go-vcard](https://github.com/emersion/go-vcard) **star:33** Parse and format vCard.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [doi](https://github.com/hscells/doi)  Document object identifier (doi) parser in Go.
    * [did](https://github.com/ockam-network/did) **star:28** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [codetree](https://github.com/aerogo/codetree) **star:14** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [enca](https://github.com/endeveit/enca) **star:8** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/enca)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/enca)
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
* Utility
    * [gotabulate](https://github.com/bndr/gotabulate) **star:230** Easily pretty-print your tabular data with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gotabulate)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:157** fast string sorting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/radix)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [Tagify](https://github.com/zoomio/tagify)  Produces a set of tags from given source.
	* [textwrap](https://github.com/isbm/textwrap)  Implementation of `textwrap` module from Python.
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself)  A sanitization-based swear filter for Go.
    * [kace](https://github.com/codemodus/kace) **star:12** Common case conversions covering common initialisms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/kace)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/kace)
    * [parseargs-go](https://github.com/nproc/parseargs-go)  string argument parser that understands quotes and backslashes.
    * [parth](https://github.com/codemodus/parth)  URL path segmentation parsing.
    * [TySug](https://github.com/Dynom/TySug) **star:4** Alternative suggestions with respect to keyboard layouts.   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
    * [xj2go](https://github.com/stackerzzq/xj2go)  Convert xml or json to go struct.
    * [xurls](https://github.com/mvdan/xurls)  Extract urls from text.

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5597** The official AWS SDK for the Go programming language.   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   [![There was an update last week][Green]](https://github.com/aws/aws-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [githubql](https://github.com/shurcooL/githubql)  Go library for accessing the GitHub GraphQL API v4.
* [github](https://github.com/google/go-github) **star:5542** Go library for accessing the GitHub REST API v3.   [![star > 2000][Awesome]](https://github.com/google/go-github)   [![There was an update last week][Green]](https://github.com/google/go-github)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [google](https://github.com/google/google-api-go-client) **star:2133** Auto-generated Google APIs for Go.   [![star > 2000][Awesome]](https://github.com/google/google-api-go-client)   [![There was an update last week][Green]](https://github.com/google/google-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [stripe](https://github.com/stripe/stripe-go) **star:1087** Go client for the Stripe API.   [![There was an update last week][Green]](https://github.com/stripe/stripe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1016** Go client library for the Twitter 1.1 API.   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:341** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist)  Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).
* [go-marathon](https://github.com/gambol99/go-marathon) **star:192** Go library for interacting with Mesosphere's Marathon PAAS.   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection)  Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).
* [facebook](https://github.com/huandu/facebook)  Go Library that supports the Facebook Graph API.
* [ethrpc](https://github.com/onrik/ethrpc) **star:180** Go bindings for Ethereum JSON RPC API.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [Trello](https://github.com/adlio/trello) **star:127** Go wrapper for the Trello API.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang)  Go wrapper for the TripAdvisor API.
* [gostorm](https://github.com/jsgilmore/gostorm) **star:124** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jsgilmore/gostorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [hipchat](https://github.com/andybons/hipchat)  This project implements a golang client library for the Hipchat API.
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:112** A golang package to communicate with HipChat over XMPP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daneharrigan/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [pushover](https://github.com/gregdel/pushover) **star:70** Go wrapper for the Pushover API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gregdel/pushover)   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [ynab](https://github.com/brunomvsouza/ynab.go)  Go wrapper for the YNAB API.
* [wit-go](https://github.com/wit-ai/wit-go) **star:60** Go client for wit.ai HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:57** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [minio-go](https://github.com/minio/minio-go)  Minio Go Library for Amazon S3 compatible cloud storage.
* [megos](https://github.com/andygrunwald/megos) **star:56** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/megos)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [circleci](https://github.com/jszwedko/go-circleci) **star:49** Go client library for interacting with CircleCI's API.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [clarifai](https://github.com/samuelcouch/clarifai)  Go client library for interfacing with the Clarifai API.
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:40** Go MusicBrainz WS2 client library.   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:40** Go client for the xkcd API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nishanths/go-xkcd)   [![godoc][GoDoc]](https://godoc.org/github.com/nishanths/go-xkcd)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:40** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [slack](https://github.com/nlopes/slack)  Slack API in Go.
* [simples3](https://github.com/rhnvrm/simples3) **star:36** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [fcm](https://github.com/maddevsio/fcm) **star:36** Go library for Firebase Cloud Messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [gads](https://github.com/emiddleton/gads)  Google Adwords Unofficial API.
* [webhooks](https://github.com/go-playground/webhooks)  Webhook receiver for GitHub and Bitbucket.
* [vl-go](https://github.com/verifid/vl-go)  Go client library around the VerifID identity verification layer API.
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:35** Go wrapper and command-line client for the Uptime Robot v2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [gosip](https://github.com/koltyakov/gosip) **star:34** Go client library SharePoint API.   [![There was an update last week][Green]](https://github.com/koltyakov/gosip)   [![godoc][GoDoc]](https://godoc.org/github.com/koltyakov/gosip)
* [golyrics](https://github.com/mamal72/golyrics) **star:34** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [mixpanel](https://github.com/dukex/mixpanel) **star:32** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk)  The Playlyfe Rest API Go SDK.
* [paypal](https://github.com/logpacker/PayPal-Go-SDK)  Wrapper for PayPal payment API.
* [patreon-go](https://github.com/mxpv/patreon-go)  Go library for Patreon API.
* [gami](https://github.com/bit4bit/gami) **star:27** Go library for Asterisk Manager Interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![Archived][Archived]](https://github.com/bit4bit/gami)
* [gcm](https://github.com/Aorioli/gcm)  Go library for Google Cloud Messaging.
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:27** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb) **star:26** Golang wrapper for The Movie Database API v3.   [![There was an update last week][Green]](https://github.com/cyruzin/golang-tmdb)   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/golang-tmdb)
* [steam](https://github.com/sostronk/go-steam)  Go Library to interact with Steam game servers.
* [spotify](https://github.com/rapito/go-spotify) **star:21** Go Library to access Spotify WEB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-spotify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-spotify)
* [shopify](https://github.com/rapito/go-shopify) **star:19** Go Library to make CRUD request to the Shopify API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-shopify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:17** Go library for accessing the BreweryDB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/naegelejd/brewerydb)   [![godoc][GoDoc]](https://godoc.org/github.com/naegelejd/brewerydb)
* [codeship-go](https://github.com/codeship/codeship-go) **star:17** Go client library for interacting with Codeship's API v2.   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [cachet](https://github.com/andygrunwald/cachet)  Go client library for [Cachet (open source status page system)](https://cachethq.io/).
* [textbelt](https://github.com/dietsche/textbelt) **star:16** Go client for the textbelt.com txt messaging API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [translate](https://github.com/poorny/translate)  Go online translation package.
* [discordgo](https://github.com/bwmarrin/discordgo)  Go bindings for the Discord Chat API.
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api)  Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:11** Simple wrapper for easy google analytics reporting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chonthu/go-google-analytics)   [![godoc][GoDoc]](https://godoc.org/github.com/chonthu/go-google-analytics)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang)  Google Cloud APIs Go Client Library.
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergiotapia/smitego)   [![godoc][GoDoc]](https://godoc.org/github.com/sergiotapia/smitego)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:10** Tiny Go client for HackerNews API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PaulRosset/go-hacknews)   [![godoc][GoDoc]](https://godoc.org/github.com/PaulRosset/go-hacknews)
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:9** Go client library for the [LastPass](https://www.lastpass.com/) API.   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [go-here](https://github.com/abdullahselek/go-here) **star:6** Go client library around the HERE location based APIs.   [![There was an update last week][Green]](https://github.com/abdullahselek/go-here)   [![godoc][GoDoc]](https://godoc.org/github.com/abdullahselek/go-here)
* [go-imgur](https://github.com/koffeinsource/go-imgur)  Go client library for [imgur](https://imgur.com)
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattcunningham/gumblr)   [![godoc][GoDoc]](https://godoc.org/github.com/mattcunningham/gumblr)
* [twitter-scraper](https://github.com/n0madic/twitter-scraper)  Scrape the Twitter Frontend API without authentication and limits.
* [go-jira](https://github.com/andygrunwald/go-jira)  Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans)  Go client library for the SPTrans Olho Vivo API.
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:5** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/esurdam/go-sophos)
* [go-twitter](https://github.com/dghubble/go-twitter)  Go client library for the Twitter v1.1 APIs.
* [go-twitch](https://github.com/knspriggs/go-twitch)  Go client for interacting with the Twitch v3 API.
* [go-trending](https://github.com/andygrunwald/go-trending)  Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gojuno/go-zooz)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   [![It hasn't been updated in the last year][Yellow]](https://github.com/axelspringer/go-chronos)   [![godoc][GoDoc]](https://godoc.org/github.com/axelspringer/go-chronos)
* [Medium](https://github.com/Medium/medium-sdk-go)  Golang SDK for Medium's OAuth2 API.
* [libgoffi](https://github.com/clevabit/libgoffi) **star:1** Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:27197** Command-line fuzzy finder written in Go.   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![There was an update last week][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [ctop](https://github.com/bcicen/ctop) **star:9499** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [wuzz](https://github.com/asciimoo/wuzz) **star:8550** Interactive cli tool for HTTP inspection.   [![star > 2000][Awesome]](https://github.com/asciimoo/wuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/wuzz)
* [sqlx](https://github.com/jmoiron/sqlx) **star:7842** provides a set of extensions on top of the excellent built-in database/sql package.   [![star > 2000][Awesome]](https://github.com/jmoiron/sqlx)   [![godoc][GoDoc]](https://godoc.org/github.com/jmoiron/sqlx)
* [peco](https://github.com/peco/peco) **star:5766** Simplistic interactive filtering tool.   [![star > 2000][Awesome]](https://github.com/peco/peco)   [![godoc][GoDoc]](https://godoc.org/github.com/peco/peco)
* [usql](https://github.com/knq/usql) **star:5637** usql is a universal command-line interface for SQL databases.   [![star > 2000][Awesome]](https://github.com/knq/usql)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/usql)
* [util](https://github.com/shomali11/util)  Collection of useful utility functions. (strings, concurrency, manipulations, ...).
* [goreporter](https://github.com/wgliang/goreporter)  Golang tool that does static analysis, unit testing, code review and generate code quality report.
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:5196** Deliver Go binaries as fast and easily as possible.   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![There was an update last week][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [godropbox](https://github.com/dropbox/godropbox) **star:3838** Common libraries for writing Go services/applications from Dropbox.   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![There was an update last week][Green]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [gohper](https://github.com/cosiner/gohper)  Various tools/modules help for development.
* [repeat](https://github.com/ssgreg/repeat)  Go implementation of different backoff strategies useful for retrying operations and heartbeating.
* [realize](https://github.com/tockins/realize) **star:3520** Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.   [![star > 2000][Awesome]](https://github.com/tockins/realize)   [![godoc][GoDoc]](https://godoc.org/github.com/tockins/realize)
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2356** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [Task](https://github.com/go-task/task) **star:2243** simple "Make" alternative.   [![star > 2000][Awesome]](https://github.com/go-task/task)   [![There was an update last week][Green]](https://github.com/go-task/task)   [![godoc][GoDoc]](https://godoc.org/github.com/go-task/task)
* [panicparse](https://github.com/maruel/panicparse) **star:2206** Groups similar goroutines and colorizes stack dump.   [![star > 2000][Awesome]](https://github.com/maruel/panicparse)   [![godoc][GoDoc]](https://godoc.org/github.com/maruel/panicparse)
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match)  Pattern matching libray.
* [minify](https://github.com/tdewolff/minify) **star:2022** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   [![star > 2000][Awesome]](https://github.com/tdewolff/minify)   [![There was an update last week][Green]](https://github.com/tdewolff/minify)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/minify)
* [go-health](https://github.com/Talento90/go-health)  Health package simplifies the way you add health check to your services.
* [go-funk](https://github.com/thoas/go-funk) **star:1596** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [Storm](https://github.com/asdine/storm) **star:1489** Simple and powerful toolkit for BoltDB.   [![godoc][GoDoc]](https://godoc.org/github.com/asdine/storm)
* [mmake](https://github.com/tj/mmake) **star:1483** Modern Make.   [![godoc][GoDoc]](https://godoc.org/github.com/tj/mmake)
* [mc](https://github.com/minio/mc) **star:1361** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   [![There was an update last week][Green]](https://github.com/minio/mc)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/mc)
* [mole](https://github.com/davrodpin/mole) **star:1352** cli app to easily create ssh tunnels.   [![godoc][GoDoc]](https://godoc.org/github.com/davrodpin/mole)
* [filetype](https://github.com/h2non/filetype) **star:1050** Small package to infer the file type checking the magic numbers signature.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [boilr](https://github.com/tmrts/boilr) **star:1041** Blazingly fast CLI tool for creating projects from boilerplate templates.   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [mergo](https://github.com/imdario/mergo) **star:1023** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   [![godoc][GoDoc]](https://godoc.org/github.com/imdario/mergo)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:857** Circuit Breakers in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [jump](https://github.com/gsamokovarov/jump) **star:801** Jump helps you navigate faster by learning your habits.   [![godoc][GoDoc]](https://godoc.org/github.com/gsamokovarov/jump)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:774** Simple, seamless, lightweight time tracking for Git.   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [immortal](https://github.com/immortal/immortal) **star:633** \*nix cross-platform (OS agnostic) supervisor.   [![godoc][GoDoc]](https://godoc.org/github.com/immortal/immortal)
* [hub](https://github.com/github/hub)  wrap git commands with additional functionality to interact with github from the terminal.
* [htcat](https://github.com/htcat/htcat) **star:507** Parallel and Pipelined HTTP GET Utility.   [![It hasn't been updated in the last year][Yellow]](https://github.com/htcat/htcat)   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [go-dry](https://github.com/ungerik/go-dry) **star:446** DRY (don't repeat yourself) package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [gopencils](https://github.com/bndr/gopencils) **star:430** Small and simple package to easily consume REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gopencils)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [rerate](https://github.com/abo/rerate)  Redis-based rate counter and rate limiter for Go.
* [request](https://github.com/mozillazg/request) **star:370** Go HTTP Requests for Humans™.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/request)
* [koazee](https://github.com/wesovilabs/koazee) **star:362** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/koazee)
* [ergo](https://github.com/cristianoliveira/ergo) **star:359** The management of multiple local services running over different ports made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [go-rate](https://github.com/beefsack/go-rate) **star:299** Timed rate limiter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beefsack/go-rate)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [clockwork](https://github.com/jonboulle/clockwork) **star:254** A simple fake clock for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:245** Simple struct copying for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:209** Package for MIME type detection based on magic numbers.   [![There was an update last week][Green]](https://github.com/gabriel-vasile/mimetype)   [![godoc][GoDoc]](https://godoc.org/github.com/gabriel-vasile/mimetype)
* [serve](https://github.com/syntaqx/serve) **star:206** A static http server anywhere you need.   [![godoc][GoDoc]](https://godoc.org/github.com/syntaqx/serve)
* [gubrak](https://github.com/novalagung/gubrak) **star:205** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [retry](https://github.com/kamilsk/retry) **star:204** The most advanced functional mechanism to perform actions repetitively until successful.   [![There was an update last week][Green]](https://github.com/kamilsk/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/retry)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:192** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [gotenv](https://github.com/subosito/gotenv) **star:168** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:163** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [rerun](https://github.com/ivpusic/rerun) **star:155** Recompiling and rerunning go apps when source changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/rerun)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/rerun)
* [moldova](https://github.com/StabbyCutyou/moldova) **star:149** Utility for generating random data based on an input template.   [![It hasn't been updated in the last year][Yellow]](https://github.com/StabbyCutyou/moldova)   [![godoc][GoDoc]](https://godoc.org/github.com/StabbyCutyou/moldova)
* [Death](https://github.com/vrecan/death) **star:148** Managing go application shutdown with signals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/vrecan/death)   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [scan](https://github.com/blockloop/scan)  Scan golang `sql.Rows` directly to structs, slices, or primitive types.
* [robustly](https://github.com/VividCortex/robustly) **star:140** Runs functions resiliently, catching and restarting panics.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/robustly)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/robustly)
* [toolbox](https://github.com/viant/toolbox) **star:124** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   [![There was an update last week][Green]](https://github.com/viant/toolbox)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/toolbox)
* [ugo](https://github.com/alxrm/ugo)  ugo is slice toolbox with concise syntax for Go.
* [circuit](https://github.com/cep21/circuit)  An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.
* [chyle](https://github.com/antham/chyle) **star:117** Changelog generator using a git repository with multiple configuration possibilities.   [![There was an update last week][Green]](https://github.com/antham/chyle)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:116** XML Sitemap generator written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:105** LiveReload server for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jaschaephraim/lrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/jaschaephraim/lrserver)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:90** Pure Go bsdiff and bspatch libraries and CLI tools.   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [pm](https://github.com/VividCortex/pm) **star:77** Process (i.e. goroutine) manager with an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/pm)
* [xferspdy](https://github.com/monmohan/xferspdy) **star:72** Xferspdy provides binary diff and patch library in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/monmohan/xferspdy)   [![godoc][GoDoc]](https://godoc.org/github.com/monmohan/xferspdy)
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:69** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/mssqlx)
* [UNIS](https://github.com/esemplastic/unis) **star:69** Common Architecture™ for String Utilities in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/esemplastic/unis)   [![godoc][GoDoc]](https://godoc.org/github.com/esemplastic/unis)
* [myhttp](https://github.com/inancgumus/myhttp)  Simple API to make HTTP GET requests with timeout support.
* [multitick](https://github.com/VividCortex/multitick) **star:62** Multiplexor for aligned tickers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/multitick)
* [minquery](https://github.com/icza/minquery) **star:53** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/minquery)
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:52** Pure Go ultra performant MIME sniffing library/utility.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zRedShift/mimemagic)   [![godoc][GoDoc]](https://godoc.org/github.com/zRedShift/mimemagic)
* [handy](https://github.com/miguelpragier/handy) **star:50** Many utilities and helpers like string handlers/formatters and validators.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [golog](https://github.com/mlimaloureiro/golog) **star:48** Easy and lightweight CLI tool to time track your tasks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mlimaloureiro/golog)   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** Parse TODOs in your GO code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/asticode/go-astitodo)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [goreadability](https://github.com/philipjkim/goreadability) **star:45** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:45** SeaweedFS client library with almost full features.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [goback](https://github.com/carlescere/goback) **star:42** Go simple exponential backoff package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [godaemon](https://github.com/VividCortex/godaemon)  Utility to write daemons.
* [gaper](https://github.com/maxcnunes/gaper) **star:42** Builds and restarts a Go project when it crashes or some watched file changes.   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:41** Use x86 SIMD without writing any assembly code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/intrinsic)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/intrinsic)   [![Archived][Archived]](https://github.com/mengzhuo/intrinsic)
* [golarm](https://github.com/msempere/golarm) **star:38** Fire alarms with system events.   [![It hasn't been updated in the last year][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [retry](https://github.com/thedevsaddam/retry) **star:36** Simple and easy retry mechanism package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thedevsaddam/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/retry)
* [pgo](https://github.com/arthurkushman/pgo) **star:36** Convenient functions for PHP community.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/pgo)
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:34** Retrying made simple and easy for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/retry-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/retry-go)
* [beyond](https://github.com/wesovilabs/beyond) **star:34** The Go tool that will drive you to the AOP world!   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [gpath](https://github.com/tenntenn/gpath) **star:33** Library to simplify access struct fields with Go's expression in reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [dbt](https://github.com/nikogura/dbt) **star:25** A framework for running self-updating signed binaries from a central, trusted repository.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [generate](https://github.com/go-playground/generate) **star:22** runs go generate recursively on a specified path or environment variable and can filter by regex.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [gostrutils](https://github.com/ik5/gostrutils) **star:22** Collections of string manipulation and conversion functions.   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [countries](https://github.com/biter777/countries) **star:21** Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts.   [![godoc][GoDoc]](https://godoc.org/github.com/biter777/countries)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:21** a small golang lib to generate placeholder images.   [![It hasn't been updated in the last year][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [filter](https://github.com/gookit/filter) **star:20** provide filtering, sanitizing, and conversion of Go data.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [evaluator](https://github.com/nullne/evaluator) **star:20** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nullne/evaluator)   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:19** Library for executing shell commands on osx, windows and linux.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:16** Go library for encoding structs into Header fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails)  Go package for working with Problem Details.
* [filler](https://github.com/yaronsumel/filler) **star:15** small utility to fill structs using "fill" tag.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [slicer](https://github.com/leaanthony/slicer) **star:15** Makes working with slices easier.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/slicer)
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** Compile-time controlled logger to make your release smaller without removing debug calls.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [spinner](https://github.com/briandowns/spinner)  Go package to easily provide a terminal spinner with options.
* [okrun](https://github.com/xta/okrun) **star:14** go run error steamroller.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xta/okrun)   [![godoc][GoDoc]](https://godoc.org/github.com/xta/okrun)
* [structs](https://github.com/PumpkinSeed/structs) **star:14** Implement simple functions to manipulate structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PumpkinSeed/structs)   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/structs)
* [olaf](https://github.com/btnguyen2k/olaf)  Twitter Snowflake implemented in Go.
* [onecache](https://github.com/adelowo/onecache)  Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
* [ghokin](https://github.com/antham/ghokin) **star:13** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![There was an update last week][Green]](https://github.com/antham/ghokin)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [rest-go](https://github.com/edermanoel94/rest-go) **star:12** A package that provide many helpful methods for working with rest api.   [![There was an update last week][Green]](https://github.com/edermanoel94/rest-go)   [![godoc][GoDoc]](https://godoc.org/github.com/edermanoel94/rest-go)
* [ctxutil](https://github.com/posener/ctxutil) **star:11** A collection of utility functions for contexts.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [retry](https://github.com/shafreeck/retry) **star:11** A pretty simple library to ensure your work to be done.   [![godoc][GoDoc]](https://godoc.org/github.com/shafreeck/retry)
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:11** A MIME type sniffer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/mimesniffer)
* [copy-pasta](https://github.com/jutkko/copy-pasta)  Universal multi-workstation clipboard that uses S3 like backend for the storage.
* [tome](https://github.com/cyruzin/tome) **star:9** Tome was designed to paginate simple RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/tome)
* [command](https://github.com/txgruppi/command) **star:9** Command pattern for Go with thread safe serial and parallel dispatcher.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [backscanner](https://github.com/icza/backscanner) **star:9** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [limiters](https://github.com/mennanov/limiters) **star:9** Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.   [![godoc][GoDoc]](https://godoc.org/github.com/mennanov/limiters)
* [jsend](https://github.com/clevergo/jsend) **star:8** JSend's implementation writen in Go.   [![There was an update last week][Green]](https://github.com/clevergo/jsend)   [![godoc][GoDoc]](https://godoc.org/github.com/clevergo/jsend)
* [shutdown](https://github.com/ztrue/shutdown) **star:8** App shutdown hooks for `os.Signal` handling.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ztrue/shutdown)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/shutdown)
* [retry](https://github.com/percolate/retry) **star:6** A simple but highly configurable retry package for Go.
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:5** Slice conversion between primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
* [apm](https://github.com/topfreegames/apm)  Process manager for Golang applications with an HTTP API.
* [slice](https://github.com/psampaz/slice)  Type-safe functions for common Go slice operations.
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.   [![godoc][GoDoc]](https://godoc.org/github.com/chrispassas/silk)
* [blank](https://github.com/Henry-Sarabia/blank) **star:4** Verify or remove blanks and whitespace from strings.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [delve](https://github.com/derekparker/delve) **star:3** Go debugger.   [![There was an update last week][Green]](https://github.com/derekparker/delve)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [rclient](https://github.com/zpatrick/rclient)  Readable, flexible, simple-to-use client for REST APIs.
* [ptr](https://github.com/gotidy/ptr) **star:3** Package that provide functions for simplified creation of pointers from constants of basic types.   [![godoc][GoDoc]](https://godoc.org/github.com/gotidy/ptr)

## UUID

*Libraries for working with UUIDs.*

* [ulid](https://github.com/oklog/ulid) **star:1828** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   [![godoc][GoDoc]](https://godoc.org/github.com/oklog/ulid)
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands.
* [uuid](https://github.com/agext/uuid)  Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
* [uuid](https://github.com/gofrs/uuid) **star:654** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   [![godoc][GoDoc]](https://godoc.org/github.com/gofrs/uuid)
* [wuid](https://github.com/edwingeng/wuid) **star:328** An extremely fast unique number generator, 10-135 times faster than UUID.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/wuid)
* [goid](https://github.com/jakehl/goid) **star:27** Generate and Parse RFC4122 compliant V4 UUIDs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jakehl/goid)   [![godoc][GoDoc]](https://godoc.org/github.com/jakehl/goid)
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:16** A tiny and efficient Go unique string ID generator.   [![godoc][GoDoc]](https://godoc.org/github.com/aidarkhanov/nanoid)
* [sno](https://github.com/muyo/sno)  Compact, sortable and fast unique IDs with embedded metadata.

## Validation

*Libraries for validation.*

* [validator](https://github.com/go-playground/validator) **star:4772** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   [![star > 2000][Awesome]](https://github.com/go-playground/validator)   [![There was an update last week][Green]](https://github.com/go-playground/validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/validator)
* [govalidator](https://github.com/asaskevich/govalidator) **star:4037** Validators and sanitizers for strings, numerics, slices and structs.   [![star > 2000][Awesome]](https://github.com/asaskevich/govalidator)   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/govalidator)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1283** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
* [terraform-validator](https://github.com/thazelart/terraform-validator)  A norms and conventions validator for Terraform.
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:834** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/govalidator)
* [checkdigit](https://github.com/osamingo/checkdigit)  Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
* [validate](https://github.com/gookit/validate) **star:202** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/validate)   [![Contains Chinese documents][CN]](https://github.com/gookit/validate)
* [jio](https://github.com/faceair/jio) **star:30** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   [![godoc][GoDoc]](https://godoc.org/github.com/faceair/jio)   [![Contains Chinese documents][CN]](https://github.com/faceair/jio)
* [validate](https://github.com/gobuffalo/validate) **star:29** This package provides a framework for writing validations for Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/validate)
* [gody](https://github.com/guiferpa/gody) **star:17** :balloon: A lightweight struct validator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/guiferpa/gody)
* [govalid](https://github.com/twharmon/govalid) **star:13** Fast, tag-based validation for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/govalid)

## Version Control

*Libraries for version control.*

* [git2go](https://github.com/libgit2/git2go) **star:1449** Go bindings for libgit2.   [![There was an update last week][Green]](https://github.com/libgit2/git2go)   [![godoc][GoDoc]](https://godoc.org/github.com/libgit2/git2go)
* [go-git](https://github.com/src-d/go-git)  highly extensible Git implementation in pure Go.
* [hercules](https://github.com/src-d/hercules) **star:876** gaining advanced insights from Git repository history.   [![There was an update last week][Green]](https://github.com/src-d/hercules)   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/hercules)
* [gh](https://github.com/rjeczalik/gh) **star:72** Scriptable server and net/http middleware for GitHub Webhooks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjeczalik/gh)   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/gh)
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:72** manipulate and inspect VCS repositories in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/sourcegraph/go-vcs)
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beyang/hgo)   [![godoc][GoDoc]](https://godoc.org/github.com/beyang/hgo)

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) **star:1026** Comphrensive Go bindings for FFmpeg.   [![There was an update last week][Green]](https://github.com/giorgisio/goav)   [![godoc][GoDoc]](https://godoc.org/github.com/giorgisio/goav)
* [m3u8](https://github.com/grafov/m3u8) **star:670** Parser and generator library of M3U8 playlists for Apple HLS.   [![There was an update last week][Green]](https://github.com/grafov/m3u8)   [![godoc][GoDoc]](https://godoc.org/github.com/grafov/m3u8)
* [gmf](https://github.com/3d0c/gmf) **star:584** Go bindings for FFmpeg av\* libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/3d0c/gmf)
* [go-astits](https://github.com/asticode/go-astits) **star:294** Parse and demux MPEG Transport Streams (.ts) natively in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astits)
* [go-astisub](https://github.com/asticode/go-astisub) **star:226** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astisub)
* [gst](https://github.com/ziutek/gst) **star:155** Go bindings for GStreamer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/gst)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/gst)
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:91** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).   [![There was an update last week][Green]](https://github.com/adrg/libvlc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/libvlc-go)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:46** Parser and generator library for Apple m3u8 playlists.   [![godoc][GoDoc]](https://godoc.org/github.com/quangngotan95/go-m3u8)
* [v4l](https://github.com/korandiz/v4l) **star:37** Video capture library for Linux, written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/korandiz/v4l)   [![godoc][GoDoc]](https://godoc.org/github.com/korandiz/v4l)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:13** Subtitle format support for go. Supports .srt, .ttml, and .ass.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wargarblgarbl/libgosubs)   [![godoc][GoDoc]](https://godoc.org/github.com/wargarblgarbl/libgosubs)

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) **star:35644** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   [![star > 2000][Awesome]](https://github.com/gin-gonic/gin)   [![There was an update last week][Green]](https://github.com/gin-gonic/gin)   [![godoc][GoDoc]](https://godoc.org/github.com/gin-gonic/gin)
* [Beego](https://github.com/astaxie/beego) **star:23289** beego is an open-source, high-performance web framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/astaxie/beego)   [![There was an update last week][Green]](https://github.com/astaxie/beego)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/beego)   [![Contains Chinese documents][CN]](https://github.com/astaxie/beego)
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo)  High performance, minimalist Go web framework.
* [Revel](https://github.com/revel/revel) **star:11580** High-productivity web framework for the Go language.   [![star > 2000][Awesome]](https://github.com/revel/revel)   [![godoc][GoDoc]](https://godoc.org/github.com/revel/revel)
* [goa](https://github.com/goa-go/goa)  goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware.
* [Goa](https://github.com/goadesign/goa) **star:3737** Goa provides a holistic approach for developing remote APIs and microservices in Go.   [![star > 2000][Awesome]](https://github.com/goadesign/goa)   [![There was an update last week][Green]](https://github.com/goadesign/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goadesign/goa)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3392** Quick and easy way to setup a RESTful JSON API.   [![star > 2000][Awesome]](https://github.com/ant0ine/go-json-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ant0ine/go-json-rest)
* [Gizmo](https://github.com/NYTimes/gizmo) **star:3056** Microservice toolkit used by the New York Times.   [![star > 2000][Awesome]](https://github.com/NYTimes/gizmo)   [![godoc][GoDoc]](https://godoc.org/github.com/NYTimes/gizmo)
* [Fiber](https://github.com/gofiber/fiber) **star:2899** An Express.js inspired web framework build on Fasthttp.   [![star > 2000][Awesome]](https://github.com/gofiber/fiber)   [![There was an update last week][Green]](https://github.com/gofiber/fiber)   [![godoc][GoDoc]](https://godoc.org/github.com/gofiber/fiber)
* [utron](https://github.com/gernest/utron) **star:2159** Lightweight MVC framework for Go(Golang).   [![star > 2000][Awesome]](https://github.com/gernest/utron)   [![It hasn't been updated in the last year][Yellow]](https://github.com/gernest/utron)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/utron)
* [gongular](https://github.com/mustafaakin/gongular) **star:425** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mustafaakin/gongular)   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaakin/gongular)
* [neo](https://github.com/ivpusic/neo) **star:405** Neo is minimal and fast Go Web Framework with extremely simple API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/neo)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/neo)
* [Banjo](https://github.com/nsheremet/banjo)  Very simple and fast web framework for Go.
* [Air](https://github.com/aofei/air) **star:370** An ideally refined web framework for Go.   [![There was an update last week][Green]](https://github.com/aofei/air)   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/air)
* [mango](https://github.com/paulbellamy/mango) **star:344** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paulbellamy/mango)   [![godoc][GoDoc]](https://godoc.org/github.com/paulbellamy/mango)
* [Aero](https://github.com/aerogo/aero) **star:246** High-performance web framework for Go, reaches top scores in Lighthouse.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/aero)
* [Golf](https://github.com/dinever/golf) **star:242** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dinever/golf)   [![godoc][GoDoc]](https://godoc.org/github.com/dinever/golf)
* [Gondola](https://github.com/rainycape/gondola)  The web framework for writing faster sites, faster.
* [Goyave](https://github.com/System-Glitch/goyave) **star:206** Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities.   [![There was an update last week][Green]](https://github.com/System-Glitch/goyave)   [![godoc][GoDoc]](https://godoc.org/github.com/System-Glitch/goyave)
* [go-rest](https://github.com/ungerik/go-rest) **star:118** Small and evil REST framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-rest)
* [hiboot](https://github.com/hidevopsio/hiboot) **star:115** hiboot is a high performance web application framework with auto configuration and dependency injection support.   [![There was an update last week][Green]](https://github.com/hidevopsio/hiboot)   [![godoc][GoDoc]](https://godoc.org/github.com/hidevopsio/hiboot)   [![Contains Chinese documents][CN]](https://github.com/hidevopsio/hiboot)
* [Macaron](https://github.com/go-macaron/macaron)  Macaron is a high productive and modular design web framework in Go.
* [WebGo](https://github.com/bnkamalesh/webgo) **star:89** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/webgo)
* [Flamingo](https://github.com/i-love-flamingo/flamingo) **star:85** Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.   [![There was an update last week][Green]](https://github.com/i-love-flamingo/flamingo)   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo)
* [uAdmin](https://github.com/uadmin/uadmin) **star:79** Fully featured web framework for Golang, inspired by Django.   [![There was an update last week][Green]](https://github.com/uadmin/uadmin)   [![godoc][GoDoc]](https://godoc.org/github.com/uadmin/uadmin)
* [Golax](https://github.com/fulldump/golax) **star:71** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fulldump/golax)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/golax)
* [Microservice](https://github.com/claygod/microservice) **star:71** The framework for the creation of microservices, written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/microservice)
* [Ginrpc](https://github.com/xxjwxc/ginrpc) **star:58** Gin parameter automatic binding tool,gin rpc tools.   [![There was an update last week][Green]](https://github.com/xxjwxc/ginrpc)   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/ginrpc)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/ginrpc)
* [YARF](https://github.com/yarf-framework/yarf) **star:55** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.   [![godoc][GoDoc]](https://godoc.org/github.com/yarf-framework/yarf)
* [Fireball](https://github.com/zpatrick/fireball) **star:50** More "natural" feeling web framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/fireball)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/fireball)
* [patron](https://github.com/beatlabs/patron) **star:50** Patron is a microservice framework following best cloud practices with a focus on productivity.   [![There was an update last week][Green]](https://github.com/beatlabs/patron)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/patron)
* [vox](https://github.com/aisk/vox) **star:48** A golang web framework for humans, inspired by Koa heavily.   [![godoc][GoDoc]](https://godoc.org/github.com/aisk/vox)
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) **star:43** Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.   [![There was an update last week][Green]](https://github.com/i-love-flamingo/flamingo-commerce)   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo-commerce)
* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [Resoursea](https://github.com/resoursea/api) **star:31** REST framework for quickly writing resource based services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/resoursea/api)   [![godoc][GoDoc]](https://godoc.org/github.com/resoursea/api)
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [rex](https://github.com/goanywhere/rex) **star:29** Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goanywhere/rex)   [![godoc][GoDoc]](https://godoc.org/github.com/goanywhere/rex)
* [rux](https://github.com/gookit/rux) **star:23** Simple and fast web framework for build golang HTTP applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/rux)   [![Contains Chinese documents][CN]](https://github.com/gookit/rux)
* [tango](https://github.com/lunny/tango)  Micro & pluggable web framework for Go.
* [tigertonic](https://github.com/rcrowley/go-tigertonic)  Go framework for building JSON web services inspired by Dropwizard.
* [goweb](https://github.com/twharmon/goweb) **star:2** Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/goweb)

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) **star:1488** Rate limit HTTP request handler.   [![godoc][GoDoc]](https://godoc.org/github.com/didip/tollbooth)
* [CORS](https://github.com/rs/cors) **star:1387** Easily add CORS capabilities to your API.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/cors)
* [Limiter](https://github.com/ulule/limiter) **star:871** Dead simple rate limit middleware for Go.   [![There was an update last week][Green]](https://github.com/ulule/limiter)   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/limiter)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:772** Add/parse Server-Timing header.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mitchellh/go-server-timing)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/go-server-timing)
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:100** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).   [![It hasn't been updated in the last year][Yellow]](https://github.com/philippgille/ln-paywall)   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/ln-paywall)
* [XFF](https://github.com/sebest/xff) **star:73** Handle `X-Forwarded-For` header and friends.   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/xff)
* [formjson](https://github.com/rs/formjson) **star:34** Transparently handle JSON input as a standard form POST.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/formjson)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/formjson)
* [client-timing](https://github.com/posener/client-timing) **star:16** An HTTP client for Server-Timing header.   [![It hasn't been updated in the last year][Yellow]](https://github.com/posener/client-timing)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/client-timing)

#### Libraries for creating HTTP middlewares

* [render](https://github.com/unrolled/render) **star:1336** Go package for easily rendering JSON, XML, and HTML template responses.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/render)
* [stats](https://github.com/thoas/stats) **star:552** Go middleware that stores various information about your web application.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/stats)
* [interpose](https://github.com/carbocation/interpose) **star:288** Minimalist net/http middleware for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carbocation/interpose)   [![godoc][GoDoc]](https://godoc.org/github.com/carbocation/interpose)
* [muxchain](https://github.com/stephens2424/muxchain) **star:209** Lightweight middleware for net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/stephens2424/muxchain)
* [negroni](https://github.com/urfave/negroni)  Idiomatic HTTP middleware for Golang.
* [renderer](https://github.com/thedevsaddam/renderer) **star:183** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/renderer)
* [rye](https://github.com/InVisionApp/rye) **star:93** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/rye)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/rye)
* [gores](https://github.com/alioygur/gores) **star:88** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alioygur/gores)   [![godoc][GoDoc]](https://godoc.org/github.com/alioygur/gores)
* [alice](https://github.com/justinas/alice)  Painless middleware chaining for Go.
* [go-wrap](https://github.com/go-on/wrap) **star:59** Small middlewares package for net/http.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-on/wrap)   [![godoc][GoDoc]](https://godoc.org/github.com/go-on/wrap)
* [catena](https://github.com/codemodus/catena) **star:7** http.Handler wrapper catenation (same API as "chain").   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/catena)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/catena)
* [chain](https://github.com/codemodus/chain)  Handler wrapper chaining with scoped data (net/context-based "middleware").

### Routers

* [httprouter](https://github.com/julienschmidt/httprouter) **star:10815** High performance router. Use this and the standard http handlers to form a very high performance web framework.   [![star > 2000][Awesome]](https://github.com/julienschmidt/httprouter)   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/httprouter)
* [chi](https://github.com/go-chi/chi) **star:7047** Small, fast and expressive HTTP router built on net/context.   [![star > 2000][Awesome]](https://github.com/go-chi/chi)   [![There was an update last week][Green]](https://github.com/go-chi/chi)   [![godoc][GoDoc]](https://godoc.org/github.com/go-chi/chi)
* [gocraft/web](https://github.com/gocraft/web) **star:1407** Mux and middleware package in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gocraft/web)
* [Bone](https://github.com/go-zoo/bone) **star:1248** Lightning Fast HTTP Multiplexer.   [![godoc][GoDoc]](https://godoc.org/github.com/go-zoo/bone)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:823** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   [![godoc][GoDoc]](https://godoc.org/github.com/buaazp/fasthttprouter)
* [GoRouter](https://github.com/vardius/gorouter)  GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
* [Goji](https://github.com/goji/goji) **star:799** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/goji/goji)
* [goroute](https://github.com/goroute/route)  Simple yet powerful HTTP request multiplexer.
* [alien](https://github.com/gernest/alien)  Lightweight and fast http router from outer space.
* [bellt](https://github.com/GuilhermeCaruso/bellt)  A simple Go HTTP router.
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:474** A simple and fast HTTP router for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gorouter)
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:409** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   [![godoc][GoDoc]](https://godoc.org/github.com/dimfeld/httptreemux)
* [lars](https://github.com/go-playground/lars) **star:380** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/lars)
* [mux](https://github.com/gorilla/mux)  Powerful URL router and dispatcher for golang.
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:373** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
* [Siesta](https://github.com/VividCortex/siesta) **star:353** Composable framework to write middleware and handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/siesta)
* [vestigo](https://github.com/husobee/vestigo) **star:261** Performant, stand-alone, HTTP compliant URL Router for go web applications.   [![godoc][GoDoc]](https://godoc.org/github.com/husobee/vestigo)
* [gowww/router](https://github.com/gowww/router) **star:159** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gowww/router)   [![godoc][GoDoc]](https://godoc.org/github.com/gowww/router)
* [violetear](https://github.com/nbari/violetear) **star:97** Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/nbari/violetear)
* [Bxog](https://github.com/claygod/Bxog) **star:94** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/Bxog)
* [pure](https://github.com/go-playground/pure) **star:91** Is a lightweight HTTP router that sticks to the std "net/http" implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pure)
* [xmux](https://github.com/rs/xmux) **star:88** High performance muxer based on `httprouter` with `net/context` support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/xmux)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xmux)
* [FastRouter](https://github.com/razonyang/fastrouter) **star:19** a fast, flexible HTTP router written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/razonyang/fastrouter)   [![godoc][GoDoc]](https://godoc.org/github.com/razonyang/fastrouter)

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:603** Win32 OLE implementation for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ole/go-ole)
* [d3d9](https://github.com/gonutz/d3d9) **star:95** Go bindings for Direct3D9.   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/d3d9)
* [gosddl](https://github.com/MonaxGT/gosddl) **star:2** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gosddl)

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) **star:325** Generate a Go struct from XML.   [![godoc][GoDoc]](https://godoc.org/github.com/miku/zek)
* [xpath](https://github.com/antchfx/xpath) **star:281** XPath package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xpath)
* [xquery](https://github.com/antchfx/xquery) **star:154** XQuery lets you extract data from HTML/XML documents using XPath expression.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/xquery)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xquery)   [![Archived][Archived]](https://github.com/antchfx/xquery)
* [xml2map](https://github.com/sbabiv/xml2map) **star:19** XML to MAP converter written Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/xml2map)
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** Simple command line XML comparer that generates diffs of folders, files and tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xml-comp/xml-comp)   [![godoc][GoDoc]](https://godoc.org/github.com/xml-comp/xml-comp)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:9** Procedural XML generation API based on libxml2's xmlwriter module.   [![godoc][GoDoc]](https://godoc.org/github.com/shabbyrobe/xmlwriter)

# Tools

*Go software and plugins.*

## Code Analysis

* [GoLint](https://github.com/golang/lint) **star:3391** Golint is a linter for Go source code.   [![star > 2000][Awesome]](https://github.com/golang/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/lint)
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [errcheck](https://github.com/kisielk/errcheck) **star:1404** Errcheck is a program for checking for unchecked errors in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/kisielk/errcheck)
* [gcvis](https://github.com/davecheney/gcvis) **star:959** Visualise Go program GC trace data in real time.   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/gcvis)
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  tarp finds functions and methods without direct unit tests in Go source code.
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [php-parser](https://github.com/z7zmey/php-parser) **star:696** A Parser for PHP written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/z7zmey/php-parser)
* [go-critic](https://github.com/go-critic/go-critic) **star:682** source code linter that brings checks that are currently not implemented in other linters.   [![There was an update last week][Green]](https://github.com/go-critic/go-critic)   [![godoc][GoDoc]](https://godoc.org/github.com/go-critic/go-critic)
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:427** Web based Golang AST visualizer.
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:330** An easy way to find outdated dependencies of your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/go-mod-outdated)
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:314** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
* [unconvert](https://github.com/mdempsky/unconvert) **star:280** Remove unnecessary type conversions from Go source.   [![godoc][GoDoc]](https://godoc.org/github.com/mdempsky/unconvert)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [gostatus](https://github.com/shurcooL/gostatus) **star:241** Command line tool, shows the status of repositories that contain Go packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shurcooL/gostatus)   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/gostatus)
* [dupl](https://github.com/mibk/dupl) **star:190** Tool for code clone detection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mibk/dupl)   [![godoc][GoDoc]](https://godoc.org/github.com/mibk/dupl)
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:172** Checks recent changes to a Go project for backwards incompatible changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bradleyfalzon/apicompat)   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyfalzon/apicompat)
* [tickgit](https://github.com/augmentable-dev/tickgit) **star:130** CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.   [![godoc][GoDoc]](https://godoc.org/github.com/augmentable-dev/tickgit)
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:103** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/checkstyle)
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) **star:102** Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.   [![godoc][GoDoc]](https://godoc.org/github.com/jfeliu007/goplantuml)
* [lint](https://github.com/surullabs/lint) **star:64** Run linters as part of go test.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surullabs/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/surullabs/lint)
* [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mccoyst/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/validate)
* [go-outdated](https://github.com/firstrow/go-outdated) **star:45** Console application that displays outdated packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/go-outdated)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/go-outdated)   [![Archived][Archived]](https://github.com/firstrow/go-outdated)
* [golines](https://github.com/segmentio/golines) **star:35** Formatter that automatically shortens long lines in Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/segmentio/golines)

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) **star:11702** Go development plugin for Vim.   [![star > 2000][Awesome]](https://github.com/fatih/vim-go)   [![There was an update last week][Green]](https://github.com/fatih/vim-go)
* [vscode-go](https://github.com/Microsoft/vscode-go)  Extension for Visual Studio Code (VS Code) which provides support for the Go language.
* [Watch](https://github.com/eaburns/Watch)  Runs a command in an acme win on file changes.
* [go-mode](https://github.com/dominikh/go-mode.el) **star:1031** Go mode for GNU/Emacs.
* [go-plus](https://github.com/joefitzgerald/go-plus)  Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.
* [gocode](https://github.com/nsf/gocode)  Autocompletion daemon for the Go programming language.
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime)  Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:32** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:19** Vim plugin for generating Go tests based on the function's or method's signature.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/gounit-vim)
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:13** Go language support for the Theia IDE.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go)  Vim plugin to highlight syntax errors on save.

## Go Generate Tools

* [generic](https://github.com/usk81/generic)  flexible data type for Go.
* [genny](https://github.com/cheekybits/genny) **star:1133** Elegant generics for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cheekybits/genny)
* [gocontracts](https://github.com/Parquery/gocontracts) **star:58** brings design-by-contract to Go by synchronizing the code with the documentation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Parquery/gocontracts)   [![godoc][GoDoc]](https://godoc.org/github.com/Parquery/gocontracts)
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests)  Generate Go tests from your source code.
* [gounit](https://github.com/hexdigest/gounit)  Generate Go tests using your own templates.
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:37** Generate Haskell inspired functions for your slices.   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/hasgo)
* [re2dfa](https://github.com/opennota/re2dfa)  Transform regular expressions into finite state machines and output Go source code.
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.

## Go Tools

* [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:4194** Navigate through go files efficiently with the OctoLinker browser extension for GitHub.   [![star > 2000][Awesome]](https://github.com/OctoLinker/browser-extension)   [![There was an update last week][Green]](https://github.com/OctoLinker/browser-extension)
* [go-james](https://github.com/pieterclaerhout/go-james)  Go project skeleton creator, builds and tests your projects without the manual setup.
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:2384** Visualize call graph of your Go program using dot format.   [![star > 2000][Awesome]](https://github.com/TrueFurby/go-callvis)   [![There was an update last week][Green]](https://github.com/TrueFurby/go-callvis)   [![godoc][GoDoc]](https://godoc.org/github.com/TrueFurby/go-callvis)
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete)  Bash completion for go and wgo.
* [go-swagger](https://github.com/go-swagger/go-swagger)  Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.
* [godbg](https://github.com/tylerwince/godbg)  Implementation of Rusts `dbg!` macro for quick and easy debugging during development.
* [gomodrun](https://github.com/dustinblackman/gomodrun/)  Go tool that executes and caches binaries included in go.mod files.
* [depth](https://github.com/KyleBanks/depth) **star:456** Visualize dependency trees of any package by analyzing imports.   [![godoc][GoDoc]](https://godoc.org/github.com/KyleBanks/depth)
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [richgo](https://github.com/kyoh86/richgo) **star:434** Enrich `go test` outputs with text decorations.   [![godoc][GoDoc]](https://godoc.org/github.com/kyoh86/richgo)
* [rts](https://github.com/galeone/rts) **star:195** RTS: response to struct. Generates Go structs from server responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/galeone/rts)   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/rts)
* [colorgo](https://github.com/songgao/colorgo) **star:103** Wrapper around `go` command for colorized `go build` output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/colorgo)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/colorgo)
* [gothanks](https://github.com/psampaz/gothanks) **star:89** GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/gothanks)
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:17** A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:63482** Container Cluster Manager from Google.   [![star > 2000][Awesome]](https://github.com/kubernetes/kubernetes)   [![There was an update last week][Green]](https://github.com/kubernetes/kubernetes)   [![godoc][GoDoc]](https://godoc.org/github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby) **star:56432** Collaborative project for the container ecosystem to assemble container-based systems.   [![star > 2000][Awesome]](https://github.com/moby/moby)   [![There was an update last week][Green]](https://github.com/moby/moby)   [![godoc][GoDoc]](https://godoc.org/github.com/moby/moby)
* [uTask](https://github.com/ovh/utask)  Automation engine that models and executes business processes declared in yaml.
* [traefik](https://github.com/containous/traefik) **star:27577** Reverse proxy and load balancer with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/containous/traefik)   [![There was an update last week][Green]](https://github.com/containous/traefik)   [![godoc][GoDoc]](https://godoc.org/github.com/containous/traefik)
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [Gitea](https://github.com/go-gitea/gitea) **star:18429** Fork of Gogs, entirely community driven.   [![star > 2000][Awesome]](https://github.com/go-gitea/gitea)   [![There was an update last week][Green]](https://github.com/go-gitea/gitea)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gitea/gitea)   [![Contains Chinese documents][CN]](https://github.com/go-gitea/gitea)
* [Vegeta](https://github.com/tsenart/vegeta) **star:13926** HTTP load testing tool and library. It's over 9000!   [![star > 2000][Awesome]](https://github.com/tsenart/vegeta)   [![There was an update last week][Green]](https://github.com/tsenart/vegeta)   [![godoc][GoDoc]](https://godoc.org/github.com/tsenart/vegeta)
* [Packer](https://github.com/mitchellh/packer) **star:9825** Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.   [![star > 2000][Awesome]](https://github.com/mitchellh/packer)   [![There was an update last week][Green]](https://github.com/mitchellh/packer)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/packer)
* [Hey](https://github.com/rakyll/hey)  Hey is a tiny program that sends some load to a web application.
* [GVM](https://github.com/moovweb/gvm) **star:5019** GVM provides an interface to manage Go versions.   [![star > 2000][Awesome]](https://github.com/moovweb/gvm)
* [webhook](https://github.com/adnanh/webhook) **star:4923** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   [![star > 2000][Awesome]](https://github.com/adnanh/webhook)   [![godoc][GoDoc]](https://godoc.org/github.com/adnanh/webhook)
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [gox](https://github.com/mitchellh/gox) **star:3628** Dead simple, no frills Go cross compile tool.   [![star > 2000][Awesome]](https://github.com/mitchellh/gox)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/gox)
* [bosun](https://github.com/bosun-monitor/bosun) **star:2951** Time Series Alerting Framework.   [![star > 2000][Awesome]](https://github.com/bosun-monitor/bosun)   [![godoc][GoDoc]](https://godoc.org/github.com/bosun-monitor/bosun)
* [goxc](https://github.com/laher/goxc) **star:1643** build tool for Go, with a focus on cross-compiling and packaging.   [![godoc][GoDoc]](https://godoc.org/github.com/laher/goxc)
* [kala](https://github.com/ajvb/kala) **star:1408** Simplistic, modern, and performant job scheduler.   [![godoc][GoDoc]](https://godoc.org/github.com/ajvb/kala)
* [script](https://github.com/bitfield/script) **star:1271** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   [![There was an update last week][Green]](https://github.com/bitfield/script)   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/script)
* [sg](https://github.com/ChristopherRabotin/sg)  Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.
* [Rodent](https://github.com/alouche/rodent)  Rodent helps you manage Go versions, projects and track dependencies.
* [Pomerium](https://github.com/pomerium/pomerium) **star:1027** Pomerium is an identity-aware access proxy.   [![There was an update last week][Green]](https://github.com/pomerium/pomerium)   [![godoc][GoDoc]](https://godoc.org/github.com/pomerium/pomerium)
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:714** Enable your Go applications to self update.   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/go-selfupdate)
* [skm](https://github.com/TimothyYe/skm) **star:589** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   [![godoc][GoDoc]](https://godoc.org/github.com/TimothyYe/skm)
* [StatusOK](https://github.com/sanathp/statusok)  Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi)  Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:569** Manage BareMetal Servers from Command Line (as easily as with Docker).   [![There was an update last week][Green]](https://github.com/scaleway/scaleway-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/scaleway/scaleway-cli)
* [govvv](https://github.com/ahmetalpbalkan/govvv) **star:424** “go build” wrapper to easily add version information into Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/govvv)
* [Mora](https://github.com/emicklei/mora) **star:269** REST server for accessing MongoDB documents and meta data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/emicklei/mora)   [![godoc][GoDoc]](https://godoc.org/github.com/emicklei/mora)
* [lstags](https://github.com/ivanilves/lstags) **star:248** Tool and API to sync Docker images across different registries.   [![godoc][GoDoc]](https://godoc.org/github.com/ivanilves/lstags)
* [Pewpew](https://github.com/bengadbois/pewpew) **star:218** Flexible HTTP command line stress tester.   [![godoc][GoDoc]](https://godoc.org/github.com/bengadbois/pewpew)
* [manssh](https://github.com/xwjdsh/manssh) **star:211** manssh is a command line tool for managing your ssh alias config easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xwjdsh/manssh)   [![godoc][GoDoc]](https://godoc.org/github.com/xwjdsh/manssh)
* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [aurora](https://github.com/xuri/aurora)  Cross-platform web-based Beanstalkd queue server console.
* [Blast](https://github.com/dave/blast) **star:180** A simple tool for API load testing and batch jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dave/blast)   [![godoc][GoDoc]](https://godoc.org/github.com/dave/blast)
* [bombardier](https://github.com/codesenberg/bombardier)  Fast cross-platform HTTP benchmarking tool.
* [gonative](https://github.com/inconshreveable/gonative)  Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [godbg](https://github.com/sirnewton01/godbg)  Web-based gdb front-end application.
* [gobrew](https://github.com/cryptojuice/gobrew) **star:177** gobrew lets you easily switch between multiple versions of go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cryptojuice/gobrew)
* [ostent](https://github.com/ostrost/ostent) **star:165** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ostrost/ostent)   [![godoc][GoDoc]](https://godoc.org/github.com/ostrost/ostent)
* [grapes](https://github.com/yaronsumel/grapes) **star:145** Lightweight tool designed to distribute commands over ssh with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/grapes)
* [jcli](https://github.com/jenkins-zh/jenkins-cli) **star:129** Jenkins CLI allows you manage your Jenkins as an easy way.   [![There was an update last week][Green]](https://github.com/jenkins-zh/jenkins-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/jenkins-zh/jenkins-cli)   [![Contains Chinese documents][CN]](https://github.com/jenkins-zh/jenkins-cli)
* [fac](https://github.com/mkchoi212/fac)  Command-line user interface to fix git merge conflicts.
* [gaia](https://github.com/gaia-pipeline/gaia)  Build powerful pipelines in any programming language.
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:127** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/easyssh-proxy)
* [kcli](https://github.com/cswank/kcli) **star:106** Command line tool for inspecting kafka topics/partitions/messages.   [![godoc][GoDoc]](https://godoc.org/github.com/cswank/kcli)
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:84** Cli tool to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm-cli)
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:75** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   [![godoc][GoDoc]](https://godoc.org/github.com/go-furnace/go-furnace)
* [drone-scp](https://github.com/appleboy/drone-scp) **star:68** Copy files and artifacts via SSH using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-scp)
* [dogo](https://github.com/liudng/dogo)  Monitoring changes in the source file and automatically compile and run (restart).
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) **star:66** A go library and an executable that produces valid Dockerfiles using various input channels.   [![godoc][GoDoc]](https://godoc.org/github.com/ozankasikci/dockerfile-generator)
* [Dropship](https://github.com/chrismckenzie/dropship) **star:50** Tool for deploying code via cdn.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrismckenzie/dropship)   [![godoc][GoDoc]](https://godoc.org/github.com/chrismckenzie/dropship)
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:27** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-jenkins)
* [awsenv](https://github.com/soniah/awsenv) **star:23** Small binary that loads Amazon (AWS) environment variables for a profile.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/awsenv)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/awsenv)
* [DepCharge](https://github.com/centerorbit/depcharge) **star:11** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![godoc][GoDoc]](https://godoc.org/github.com/centerorbit/depcharge)
* [lwc](https://github.com/timdp/lwc) **star:11** A live-updating version of the UNIX wc command.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timdp/lwc)   [![godoc][GoDoc]](https://godoc.org/github.com/timdp/lwc)
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) **star:2** S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).   [![godoc][GoDoc]](https://godoc.org/github.com/oxyno-zeta/s3-proxy)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r)  Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.

### Other Software

* [restic](https://github.com/restic/restic) **star:9463** De-duplicating backup program.   [![star > 2000][Awesome]](https://github.com/restic/restic)   [![There was an update last week][Green]](https://github.com/restic/restic)   [![godoc][GoDoc]](https://godoc.org/github.com/restic/restic)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:9208** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   [![star > 2000][Awesome]](https://github.com/chrislusf/seaweedfs)   [![There was an update last week][Green]](https://github.com/chrislusf/seaweedfs)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/seaweedfs)
* [confd](https://github.com/kelseyhightower/confd) **star:6828** Manage local application configuration files using templates and data from etcd or consul.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/confd)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/confd)
* [Comcast](https://github.com/tylertreat/Comcast) **star:6459** Simulate bad network connections.   [![star > 2000][Awesome]](https://github.com/tylertreat/Comcast)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/Comcast)
* [LiteIDE](https://github.com/visualfc/liteide) **star:5870** LiteIDE is a simple, open source, cross-platform Go IDE.   [![star > 2000][Awesome]](https://github.com/visualfc/liteide)   [![Contains Chinese documents][CN]](https://github.com/visualfc/liteide)
* [drive](https://github.com/odeke-em/drive) **star:5291** Google Drive client for the commandline.   [![star > 2000][Awesome]](https://github.com/odeke-em/drive)   [![godoc][GoDoc]](https://godoc.org/github.com/odeke-em/drive)
* [nes](https://github.com/fogleman/nes) **star:4362** Nintendo Entertainment System (NES) emulator written in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/nes)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/nes)
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:4344** Proxy to simulate network and system conditions for automated tests.   [![star > 2000][Awesome]](https://github.com/shopify/toxiproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/shopify/toxiproxy)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:3176** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   [![star > 2000][Awesome]](https://github.com/gilbertchen/duplicacy)   [![There was an update last week][Green]](https://github.com/gilbertchen/duplicacy)   [![godoc][GoDoc]](https://godoc.org/github.com/gilbertchen/duplicacy)
* [croc](https://github.com/schollz/croc) **star:2837** Easily and securely send files or folders from one computer to another.   [![star > 2000][Awesome]](https://github.com/schollz/croc)   [![There was an update last week][Green]](https://github.com/schollz/croc)   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/croc)
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [myLG](https://github.com/mehrdadrad/mylg) **star:2273** Command Line Network Diagnostic tool written in Go.   [![star > 2000][Awesome]](https://github.com/mehrdadrad/mylg)   [![There was an update last week][Green]](https://github.com/mehrdadrad/mylg)   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/mylg)
* [GoBoy](https://github.com/Humpheh/goboy) **star:2192** Nintendo Game Boy Color emulator written in Go.   [![star > 2000][Awesome]](https://github.com/Humpheh/goboy)   [![godoc][GoDoc]](https://godoc.org/github.com/Humpheh/goboy)
* [Stack Up](https://github.com/pressly/sup) **star:2107** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   [![star > 2000][Awesome]](https://github.com/pressly/sup)   [![There was an update last week][Green]](https://github.com/pressly/sup)   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/sup)
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [lgo](https://github.com/yunabe/lgo) **star:1968** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/yunabe/lgo)
* [Circuit](https://github.com/gocircuit/circuit) **star:1817** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gocircuit/circuit)
* [snap](https://github.com/intelsdi-x/snap) **star:1806** Powerful telemetry framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/intelsdi-x/snap)   [![godoc][GoDoc]](https://godoc.org/github.com/intelsdi-x/snap)
* [scc](https://github.com/boyter/scc) **star:1337** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   [![godoc][GoDoc]](https://godoc.org/github.com/boyter/scc)
* [Documize](https://github.com/documize/community) **star:1020** Modern wiki software that integrates data from SaaS tools.
* [peg](https://github.com/pointlander/peg) **star:676** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   [![godoc][GoDoc]](https://godoc.org/github.com/pointlander/peg)
* [Leaps](https://github.com/jeffail/leaps) **star:665** Pair programming service using Operational Transforms.   [![godoc][GoDoc]](https://godoc.org/github.com/jeffail/leaps)
* [vFlow](https://github.com/VerizonDigital/vflow) **star:653** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   [![godoc][GoDoc]](https://godoc.org/github.com/VerizonDigital/vflow)
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store)  App that displays updates for the Go packages in your GOPATH.
* [gfile](https://github.com/Antonito/gfile) **star:533** Securely transfer files between two computers, without any third party, over WebRTC.   [![godoc][GoDoc]](https://godoc.org/github.com/Antonito/gfile)
* [shell2http](https://github.com/msoap/shell2http) **star:513** Executing shell commands via http server (for prototyping or remote control).   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/shell2http)
* [mockingjay](https://github.com/quii/mockingjay-server) **star:442** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   [![godoc][GoDoc]](https://godoc.org/github.com/quii/mockingjay-server)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:395** Video streaming torrent client.   [![godoc][GoDoc]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
* [gocc](https://github.com/goccmack/gocc) **star:384** Gocc is a compiler kit for Go written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/goccmack/gocc)
* [ipe](https://github.com/dimiro1/ipe) **star:298** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dimiro1/ipe)   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/ipe)
* [wellington](https://github.com/wellington/wellington) **star:288** Sass project management tool, extends the language with sprite functions (like Compass).   [![It hasn't been updated in the last year][Yellow]](https://github.com/wellington/wellington)   [![godoc][GoDoc]](https://godoc.org/github.com/wellington/wellington)
* [ide](https://github.com/thestrukture/ide) **star:254** Browser accessible IDE. Designed for Go with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thestrukture/ide)
* [orange-cat](https://github.com/noraesae/orange-cat) **star:184** Markdown previewer written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/noraesae/orange-cat)   [![godoc][GoDoc]](https://godoc.org/github.com/noraesae/orange-cat)
* [Orbit](https://github.com/gulien/orbit) **star:134** A simple tool for running commands and generating files from templates.   [![godoc][GoDoc]](https://godoc.org/github.com/gulien/orbit)
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [joincap](https://github.com/assafmo/joincap) **star:128** Command-line utility for merging multiple pcap files together.   [![godoc][GoDoc]](https://godoc.org/github.com/assafmo/joincap)
* [borg](https://github.com/crufter/borg)  Terminal based search engine for bash snippets.
* [boxed](https://github.com/tejo/boxed) **star:72** Dropbox based blog engine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tejo/boxed)   [![godoc][GoDoc]](https://godoc.org/github.com/tejo/boxed)
* [Cherry](https://github.com/rafael-santiago/cherry)  Tiny webchat server in Go.
* [dp](https://github.com/scryinfo/dp) **star:57** Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.   [![godoc][GoDoc]](https://godoc.org/github.com/scryinfo/dp)
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** Simple NaCL EC25519 based crypto pipe tool written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/unix4fun/naclpipe)   [![godoc][GoDoc]](https://godoc.org/github.com/unix4fun/naclpipe)
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.   [![It hasn't been updated in the last year][Yellow]](https://github.com/crazcalm/term-quiz)   [![godoc][GoDoc]](https://godoc.org/github.com/crazcalm/term-quiz)
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lucasgomide/snitch)   [![godoc][GoDoc]](https://godoc.org/github.com/lucasgomide/snitch)
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [Gor](https://github.com/buger/gor)  Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:11** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.   [![It hasn't been updated in the last year][Yellow]](https://github.com/diankong/GoDocTooltip)

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1344** Go HTTP request router benchmark and comparison.   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1128** Go web framework benchmark.   [![There was an update last week][Green]](https://github.com/smallnest/go-web-framework-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
* [skynet](https://github.com/atemerev/skynet) **star:947** Skynet 1M threads microbenchmark.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:946** Benchmarks of Go serialization methods.   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:189** Compare various Image resize algorithms for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/fawick/speedtest-resize)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:131** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/go-benchmarks)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/go-benchmarks)
* [autobench](https://github.com/davecheney/autobench) **star:90** Framework to compare the performance between different Go versions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davecheney/autobench)   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/autobench)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) **star:55** Benchmarks of common basic operations for the Go language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PuerkitoBio/gocostmodel)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/gocostmodel)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:52** Collection of benchmarks for popular Go database/SQL utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tyler-smith/golang-sql-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/tyler-smith/golang-sql-benchmark)
* [gospeed](https://github.com/feyeleanor/GoSpeed)  Go micro-benchmarks for calculating the speed of language constructs.
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:20** Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mrLSD/go-benchmark-app)   [![godoc][GoDoc]](https://godoc.org/github.com/mrLSD/go-benchmark-app)
* [kvbench](https://github.com/jimrobinson/kvbench) **star:16** Key/Value database benchmark.   [![godoc][GoDoc]](https://godoc.org/github.com/jimrobinson/kvbench)

## Conferences

* [Capital Go](http://www.capitalgolang.com)  Washington, D.C., USA.
* [GopherCon Europe](https://gophercon.is/)  Berlin, Germany.
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
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:13** in Persian.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thedevsir/gosuccinctly)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsir/gosuccinctly)
* [GoBooks](https://github.com/dariubs/GoBooks)  A curated list of Go books.
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1804** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   [![godoc][GoDoc]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:71** adorable gopher logos.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GolangUA/gopher-logos)
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:32** Go gopher Vector Data [.ai, .svg].   [![It hasn't been updated in the last year][Yellow]](https://github.com/keygx/Go-gopher-Vector)
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:2061** Gopher artworks by Ashley McNamara.   [![star > 2000][Awesome]](https://github.com/ashleymcnamara/gophers)   [![godoc][GoDoc]](https://godoc.org/github.com/ashleymcnamara/gophers)
* [gophers](https://github.com/egonelbre/gophers) **star:1878** Free gophers.   [![godoc][GoDoc]](https://godoc.org/github.com/egonelbre/gophers)
* [gopherize.me](https://github.com/matryer/gopherize.me)  Gopherize yourself.
* [gophers](https://github.com/rogeralsing/gophers) **star:51** random gopher graphics.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rogeralsing/gophers)
* [gophers](https://github.com/sillecelik/go-gopher) **star:43** Gopher amigurumi toy pattern.

## Meetups

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

*Add the group of your city/country here (send **PR**)*

## Social Media
### Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

### Reddit
 * [r/golang](https://www.reddit.com/r/golang/)

## Websites

* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:25678** List of other amazingly awesome lists.   [![star > 2000][Awesome]](https://github.com/bayandin/awesome-awesomeness)   [![There was an update last week][Green]](https://github.com/bayandin/awesome-awesomeness)
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [go.dev](https://go.dev/)  A hub for Go developers.
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [json2go](https://m-zajac.github.io/json2go)  Advanced JSON to Go struct conversion - online tool.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:142** Collection of Go images, graphics, and art.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/golang-graphics)   [![Archived][Archived]](https://github.com/mholt/golang-graphics)
* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job)  Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Golang Developer Jobs](https://golangjob.xyz)  Developer Jobs exclusivly for Golang related Roles.
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:37** Collection of Go projects that needs help. Good place to start your open-source way in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ninedraft/gocryforhelp)
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:33808** Golang ebook intro how to build a web app with golang.   [![star > 2000][Awesome]](https://github.com/astaxie/build-web-application-with-golang)   [![There was an update last week][Green]](https://github.com/astaxie/build-web-application-with-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   [![Contains Chinese documents][CN]](https://github.com/astaxie/build-web-application-with-golang)
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:554** A little e-book on Ethereum Development with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/ethereum-development-with-go-book)   [![Contains Chinese documents][CN]](https://github.com/miguelmota/ethereum-development-with-go-book)
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)  Go's reference card.
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns) **star:12156** Curated list of Go design patterns, recipes and idioms.   [![star > 2000][Awesome]](https://github.com/tmrts/go-patterns)   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/go-patterns)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:10159** Learn Go with test-driven development.   [![star > 2000][Awesome]](https://github.com/quii/learn-go-with-tests)   [![godoc][GoDoc]](https://godoc.org/github.com/quii/learn-go-with-tests)   [![Contains Chinese documents][CN]](https://github.com/quii/learn-go-with-tests)
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera Specialization to learn about Go from scratch.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:1176** Examples of Golang compared to Node.js for learning.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/)  Collection of code snippets and tutorials to help tackle every day issues.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1151** Intro to go for experienced programmers.
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

