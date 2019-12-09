# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2019-12-09 10:52:02(Synchronize every day)**

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

* [Oto](https://github.com/hajimehoshi/oto) **star:485** A low-level library to play sound on multiple platforms.   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:324** Go bindings for the PortAudio audio I/O library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gordonklaus/portaudio)   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:269** Music theory models in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:263** Go package capable of generating waveform images from audio streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/waveform)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:217** Go bindings for PortMidi.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:125** Fast and stable ID3 parsing and writing library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [flac](https://github.com/mewkiz/flac) **star:108** Native Go FLAC encoder/decoder with support for FLAC streams.   [![godoc][GoDoc]](https://godoc.org/github.com/mewkiz/flac)
* [mix](https://github.com/go-mix/mix) **star:103** Sequence-based Go-native audio mixer for music apps.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-mix/mix)   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:100** Native Go MP3 decoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:96** libsox bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [malgo](https://github.com/gen2brain/malgo) **star:86** Mini audio library.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/malgo)
* [taglib](https://github.com/wtolson/go-taglib) **star:70** Go bindings for taglib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wtolson/go-taglib)   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:61** Native Go AAC bitstream parser.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Comcast/gaad)   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:33** Lightweight MP3 decoder library.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:27** libmediainfo bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/go_mediainfo)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/go_mediainfo)
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:24** EasyMidi is a simple and reliable library for working with standard midi file (SMF).   [![It hasn't been updated in the last year][Yellow]](https://github.com/algoGuy/EasyMIDI)   [![godoc][GoDoc]](https://godoc.org/github.com/algoGuy/EasyMIDI)
* [vorbis](https://github.com/mccoyst/vorbis) **star:23** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:8** libsamplerate bindings for go.   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:6662** Golang implementation of JSON Web Tokens (JWT).   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:5681** Authorization library that supports access control models like ACL, RBAC, ABAC.   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![There was an update last week][Green]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [oauth2](https://github.com/golang/oauth2) **star:2598** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![There was an update last week][Green]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2429** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![There was an update last week][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:2047** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   [![star > 2000][Awesome]](https://github.com/volatiletech/authboss)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [osin](https://github.com/openshift/osin) **star:1566** Golang OAuth2 server library.   [![godoc][GoDoc]](https://godoc.org/github.com/openshift/osin)
* [go-jose](https://github.com/square/go-jose) **star:1406** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.   [![godoc][GoDoc]](https://godoc.org/github.com/square/go-jose)
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1386** Standalone, specification-compliant,  OAuth2 server written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [gologin](https://github.com/dghubble/gologin) **star:1123** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:960** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [loginsrv](https://github.com/tarent/loginsrv) **star:864** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   [![There was an update last week][Green]](https://github.com/tarent/loginsrv)   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [scs](https://github.com/alexedwards/scs) **star:626** Session Manager for HTTP servers.   [![There was an update last week][Green]](https://github.com/alexedwards/scs)   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [permissions2](https://github.com/xyproto/permissions2) **star:375** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [paseto](https://github.com/o1egl/paseto) **star:264** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![There was an update last week][Green]](https://github.com/o1egl/paseto)   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [httpauth](https://github.com/goji/httpauth) **star:187** HTTP Authentication middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goji/httpauth)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:182** Simple, flexible, secure and idiomatic web session management with pluggable backends.   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:162** JWT middleware for Golang http servers with many configuration options.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/pascaldekloe/jwt) **star:133** Lightweight JSON Web Token (JWT) library.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [session](https://github.com/icza/session) **star:95** Go session management for web servers (including support for Google App Engine - GAE).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [branca](https://github.com/hako/branca) **star:86** Golang implementation of Branca Tokens.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hako/branca)   [![godoc][GoDoc]](https://godoc.org/github.com/hako/branca)
* [jwt](https://github.com/robbert229/jwt) **star:76** Clean and easy to use implementation of JSON Web Tokens (JWT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/robbert229/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sessionup](https://github.com/swithek/sessionup) **star:75** Simple, yet effective HTTP session management and identification package.   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [sessions](https://github.com/adam-hanna/sessions) **star:48** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [sjwt](https://github.com/brianvoe/sjwt) **star:44** Simple jwt generator and parser.   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [rbac](https://github.com/zpatrick/rbac) **star:37** Minimalistic RBAC package for Go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/rbac)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [securecookie](https://github.com/chmike/securecookie) **star:34** Efficient secure cookie encoding/decoding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chmike/securecookie)   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.   [![It hasn't been updated in the last year][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [signedvalue](https://github.com/sashka/signedvalue) **star:8** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)
* [scope](https://github.com/SonicRoshan/scope) **star:2** Easily Manage OAuth2 Scopes In Go.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** provides parser of cookies.txt file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1805** Simple and clean Telegram bot client.   [![There was an update last week][Green]](https://github.com/Syfaro/telegram-bot-api)   [![godoc][GoDoc]](https://godoc.org/github.com/Syfaro/telegram-bot-api)
* [telebot](https://github.com/tucnak/telebot) **star:1051** Telegram bot framework written in Go.   [![There was an update last week][Green]](https://github.com/tucnak/telebot)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:524** IRC, Slack & Telegram bot written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [slacker](https://github.com/shomali11/slacker) **star:343** Easy to use framework to create Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:260** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [tbot](https://github.com/yanzay/tbot) **star:239** Telegram bot server with API similar to net/http.   [![There was an update last week][Green]](https://github.com/yanzay/tbot)   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [Kelp](https://github.com/stellar/kelp) **star:237** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:168** Service oriented IRC bot using Redis and JSON for messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:150** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.   [![godoc][GoDoc]](https://godoc.org/github.com/oklahomer/go-sarah)
* [hanu](https://github.com/sbstjn/hanu) **star:116** Framework for writing Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:90** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/olebedev/go-tgbot)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:60** Framework for building Telegram bots.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:29** Simple Go [VK](https://vk.com) bot library.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:18** Another framework for building Slack bots.   [![There was an update last week][Green]](https://github.com/alexandre-normand/slackscot)   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [micha](https://github.com/onrik/micha) **star:12** Go Library for Telegram bot api.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:14918** Commander for modern Go CLI interactions.   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![There was an update last week][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:12575** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![There was an update last week][Green]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2737** Command line and flag parser supporting sub commands.   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![There was an update last week][Green]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1579** go command line option parser.   [![There was an update last week][Green]](https://github.com/jessevdk/go-flags)   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [readline](https://github.com/chzyer/readline) **star:1412** Pure golang implementation that provides most features in GNU-Readline under MIT license.   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [Dnote](https://github.com/dnote/dnote) **star:1348** A simple and end-to-end encrypted notebook for developers.   [![There was an update last week][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [docopt.go](https://github.com/docopt/docopt.go) **star:1201** Command-line arguments parser that will make you smile.   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1055** Go library for implementing command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [pflag](https://github.com/spf13/pflag) **star:920** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   [![There was an update last week][Green]](https://github.com/spf13/pflag)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [cli-init](https://github.com/tcnksm/gcli) **star:884** The easy way to start building Golang command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcnksm/gcli)   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [go-arg](https://github.com/alexflint/go-arg) **star:803** Struct-based argument parsing in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:657** Write bash completions in Go + Go command bash completion.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [mow.cli](https://github.com/jawher/mow.cli) **star:637** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [liner](https://github.com/peterh/liner) **star:614** Go readline-like library for command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [flaggy](https://github.com/integrii/flaggy) **star:523** A robust and idiomatic flags package with excellent subcommand support.   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:497** Feature-rich and easy to use command-line package based on golang struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:308** Unikernel Builder/Orchestrator.   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:139** Command line argument parser inspired by Python's argparse module.   [![There was an update last week][Green]](https://github.com/akamensky/argparse)   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [commandeer](https://github.com/jaffee/commandeer) **star:104** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [flag](https://github.com/cosiner/flag) **star:102** Simple but powerful command line option parsing library for Go supporting subcommand.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [sflags](https://github.com/octago/sflags) **star:99** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [ukautz/clif](https://github.com/ukautz/clif) **star:97** Small command line interface framework.   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [wmenu](https://github.com/dixonwille/wmenu) **star:95** Easy to use menu structure for cli applications that prompts users to make choices.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [job](https://github.com/liujianping/job) **star:63** JOB, make your short-term command as a long-term job.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [cli](https://github.com/teris-io/cli) **star:58** Simple and complete API for building command line interfaces in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [env](https://github.com/codingconcepts/env) **star:45** Tag-based environment configuration for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [1build](https://github.com/gopinath-langote/1build) **star:42** Command line tool to frictionlessly manage project-specific commands.   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [wlog](https://github.com/dixonwille/wlog) **star:40** Simple logging interface that supports cross-platform color and concurrency.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/wlog)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wlog)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [gocmd](https://github.com/devfacet/gocmd) **star:36** Go library for building command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [flagvar](https://github.com/sgreben/flagvar) **star:31** A collection of flag argument types for Go's standard `flag` package.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [strumt](https://github.com/antham/strumt) **star:31** Library to create prompt chain.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [cmdr](https://github.com/hedzr/cmdr) **star:21** A POSIX/GNU style, getopt-like command-line UI Go library.   [![There was an update last week][Green]](https://github.com/hedzr/cmdr)   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [argv](https://github.com/cosiner/argv) **star:18** Go library to split command line string as arguments array using the bash syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:16** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [sand](https://github.com/Zaba505/sand) **star:9** Simple API for creating interpreters and so much more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Zaba505/sand)   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [ts](https://github.com/liujianping/ts) **star:7** Timestamp convert & compare tool.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/ts)
* [cmd](https://github.com/posener/cmd) **star:2** Extends the standard `flag` package to support sub commands and more in idomatic way.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)
* [clîr](https://github.com/leaanthony/clir) **star:2** A Simple and Clear CLI library. Dependency free.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:9324** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [gocui](https://github.com/jroimartin/gocui) **star:5756** Minimalist Go library aimed at creating Console User Interfaces.   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [termbox-go](https://github.com/nsf/termbox-go) **star:3605** Termbox is a library for creating cross-platform text-based interfaces.   [![star > 2000][Awesome]](https://github.com/nsf/termbox-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/termbox-go)
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2688** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1599** Flexible library to render progress bars in terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1227** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:1036** Library for updating terminal output in realtime.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [termdash](https://github.com/mum4k/termdash) **star:894** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [progressbar](https://github.com/schollz/progressbar) **star:814** Basic thread-safe progress bar that works in every OS.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/progressbar)
* [mpb](https://github.com/vbauerster/mpb) **star:790** Multi progress bar for terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [aurora](https://github.com/logrusorgru/aurora) **star:730** ANSI terminal colors that supports fmt.Printf/Sprintf.   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [uitable](https://github.com/gosuri/uitable) **star:527** Library to improve readability in terminal apps using tabular data.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:398** Colorable writer for windows.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:379** isatty for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [chalk](https://github.com/ttacon/chalk) **star:325** Intuitive package for prettifying terminal/console output.   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [gookit/color](https://github.com/gookit/color) **star:278** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![Contains Chinese documents][CN]](https://github.com/gookit/color)
* [tabby](https://github.com/cheynewallace/tabby) **star:255** A tiny library for super simple Golang tables.   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:201** Go library for color output in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-colortext)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [simpletable](https://github.com/alexeyco/simpletable) **star:180** Simple tables in terminal with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [cfmt](https://github.com/mingrammer/cfmt) **star:69** Contextual fmt inspired by bootstrap color classes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mingrammer/cfmt)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:35** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/tabular)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [ctc](https://github.com/wzshiming/ctc)  The non-invasive cross-platform terminal color library does not need to modify the Print method.
* [colourize](https://github.com/TreyBastian/colourize) **star:17** Go library for ANSI colour text in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** Go library for rendering ANSI colored text templates in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/workanator/go-ataman)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:10551** Go configuration with fangs.   [![star > 2000][Awesome]](https://github.com/spf13/viper)   [![There was an update last week][Green]](https://github.com/spf13/viper)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2643** Go library for managing configuration data from environment variables.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![There was an update last week][Green]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv) **star:2434** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [ini](https://github.com/go-ini/ini) **star:1780** Go package to read and write INI files.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:1354** Parse environment variables to Go structs (with defaults).   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [konfig](https://github.com/lalamove/konfig) **star:540** Composable, observable and performant config handling for Go for the distributed processing era.   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [confita](https://github.com/heetch/confita) **star:279** Load configuration in cascade from multiple backends into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:246** Lightweight configuration manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tucnak/store)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/olebedev/config) **star:221** JSON or YAML configuration wrapper with environment variables and flags parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/config)
* [config](https://github.com/JeremyLoy/config) **star:209** Cloud native application configuration. Bind ENV to structs in only two lines.   [![There was an update last week][Green]](https://github.com/JeremyLoy/config)   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [joshbetz/config](https://github.com/joshbetz/config) **star:199** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.   [![godoc][GoDoc]](https://godoc.org/github.com/joshbetz/config)
* [hjson](https://github.com/hjson/hjson-go) **star:182** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [envconfig](https://github.com/vrischmann/envconfig) **star:169** Read your configuration from environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/vrischmann/envconfig)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:120** read INI-style configuration files into Go structs; supports user-defined types and subsections.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [goConfig](https://github.com/crgimenes/goConfig) **star:118** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [gookit/config](https://github.com/gookit/config) **star:114** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   [![Contains Chinese documents][CN]](https://github.com/gookit/config)
* [koanf](https://github.com/knadh/koanf) **star:113** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [envh](https://github.com/antham/envh) **star:92** Helpers to manage environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [envcfg](https://github.com/tomazk/envcfg) **star:90** Un-marshaling environment variables to Go structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomazk/envcfg)   [![godoc][GoDoc]](https://godoc.org/github.com/tomazk/envcfg)
* [gofigure](https://github.com/ian-kent/gofigure) **star:58** Go application configuration made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [configure](https://github.com/paked/configure) **star:52** Provides configuration through multiple sources, including JSON, flags and environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [harvester](https://github.com/beatlabs/harvester) **star:45** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/harvester)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:42** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [config](https://github.com/golobby/config) **star:35** A lightweight yet powerful config package for Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [ingo](https://github.com/schachmat/ingo) **star:26** Flags persisted in an ini-like config file.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schachmat/ingo)   [![godoc][GoDoc]](https://godoc.org/github.com/schachmat/ingo)
* [go-up](https://github.com/ufoscout/go-up) **star:26** A simple configuration library with recursive placeholders resolution and no magic.   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:20** Golang package for parsing ini-style configuration files.   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:10** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.   [![godoc][GoDoc]](https://godoc.org/github.com/the4thamigo-uk/conflate)
* [genv](https://github.com/sakirsensoy/genv) **star:8** Read environment variables easily with dotenv support.   [![godoc][GoDoc]](https://godoc.org/github.com/sakirsensoy/genv)
* [envconf](https://github.com/ian-kent/envconf) **star:8** Configuration from environment.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/envconf)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/envconf)
* [sprbox](https://github.com/oblq/sprbox) **star:4** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/sprbox)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:2** Simple useful package for read environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/nasermirzaei89/env)
* [onion](http://github.com/goraz/onion)  Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:20027** Drone is a Continuous Integration platform built on Docker, written in Go.   [![star > 2000][Awesome]](https://github.com/drone/drone)   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [CDS](https://github.com/ovh/cds) **star:2515** Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.   [![star > 2000][Awesome]](https://github.com/ovh/cds)   [![There was an update last week][Green]](https://github.com/ovh/cds)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/cds)
* [goveralls](https://github.com/mattn/goveralls) **star:615** Go integration for Coveralls.io continuous code coverage tracking system.   [![There was an update last week][Green]](https://github.com/mattn/goveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/goveralls)
* [overalls](https://github.com/go-playground/overalls) **star:98** Multi-Package go project coverprofile for tools like goveralls.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/overalls)
* [duci](https://github.com/duck8823/duci) **star:49** A simple ci server no needs domain specific languages.   [![There was an update last week][Green]](https://github.com/duck8823/duci)   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:38** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![There was an update last week][Green]](https://github.com/nikogura/gomason)   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** Recursive coverage testing tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:425** Pure Go CSS Preprocessor.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/gcss)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass) **star:142** Go wrapper to the 100% Sass compatible libsass project.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:7233** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5342** Collection of useful, performant, and thread-safe data structures.   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1331** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![There was an update last week][Green]](https://github.com/deckarep/golang-set)   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1200** Probabilistic data structures for processing continuous, unbounded streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/BoomFilters)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [gota](https://github.com/kniren/gota) **star:1008** Implementation of dataframes, series, and data wrangling methods for Go.   [![There was an update last week][Green]](https://github.com/kniren/gota)   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [roaring](https://github.com/RoaringBitmap/roaring) **star:750** Go package implementing compressed bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/RoaringBitmap/roaring)
* [willf/bloom](https://github.com/willf/bloom) **star:720** Go package implementing Bloom filters.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bloom)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:672** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:557** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [bitset](https://github.com/willf/bitset) **star:522** Go package implementing bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [trie](https://github.com/derekparker/trie) **star:443** Trie implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/trie)
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:317** In-memory geo index.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hailocab/go-geoindex)   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [algorithms](https://github.com/shady831213/algorithms) **star:304** Algorithms and data structures.CLRS study.   [![godoc][GoDoc]](https://godoc.org/github.com/shady831213/algorithms)
* [gocache](https://github.com/eko/gocache) **star:298** A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.   [![There was an update last week][Green]](https://github.com/eko/gocache)   [![godoc][GoDoc]](https://godoc.org/github.com/eko/gocache)
* [mafsa](https://github.com/smartystreets/mafsa) **star:279** MA-FSA implementation with Minimal Perfect Hashing.   [![godoc][GoDoc]](https://godoc.org/github.com/smartystreets/mafsa)   [![Archived][Archived]](https://github.com/smartystreets/mafsa)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:211** Skip list implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [hilbert](https://github.com/google/hilbert) **star:190** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   [![It hasn't been updated in the last year][Yellow]](https://github.com/google/hilbert)   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [merkletree](https://github.com/cbergoon/merkletree) **star:172** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [bloom](https://github.com/zhenjl/bloom) **star:131** Bloom filters implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:130** Binary packer and unpacker helps user build custom binary stream.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhuangsirui/binpacker)   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:117** Go implementation of Adaptive Radix Tree.   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:113** In-memory LRU string-interface{} map with expiration for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/diegobernardes/ttlcache)
* [skiplist](https://github.com/MauriceGit/skiplist) **star:108** Very fast Go Skiplist implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:100** Region quadtrees with efficient point location and neighbour finding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aurelien-rainone/go-rquad)   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [encoding](https://github.com/zhenjl/encoding) **star:97** Integer Compression Libraries for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/encoding)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [ring](https://github.com/TheTannerRyan/ring) **star:94** Go implementation of a high performance, thread safe bloom filter.   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [deque](https://github.com/gammazero/deque) **star:90** Fast ring-buffer deque (double-ended queue).   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:84** A small, powerful and flexible merge library.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [bit](https://github.com/yourbasic/bit) **star:74** Golang set data structure with bonus bit-twiddling functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bit)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [skiplist](https://github.com/gansidui/skiplist) **star:66** Skiplist implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:59** Implementation to calculate levenshtein distance in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [bloom](https://github.com/yourbasic/bloom) **star:46** Golang Bloom filter implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bloom)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:45** Concurrent FIFO queue.   [![There was an update last week][Green]](https://github.com/enriquebris/goconcurrentqueue)   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:43** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/count-min-log)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/count-min-log)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:40** Fast in-memory key:value store/cache library. Pointer caches.   [![godoc][GoDoc]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
* [levenshtein](https://github.com/agext/levenshtein) **star:35** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [remember-go](https://github.com/rocketlaunchr/remember-go) **star:30** A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/remember-go)
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:26** Highly concurrent drop-in replacement for `bufio.Writer`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/free/concurrent-writer)   [![godoc][GoDoc]](https://godoc.org/github.com/free/concurrent-writer)
* [crunch](https://github.com/superwhiskers/crunch) **star:20** Go package implementing buffers for handling various datatypes easily.   [![godoc][GoDoc]](https://godoc.org/github.com/superwhiskers/crunch)
* [pipeline](https://github.com/hyfather/pipeline) **star:18** An implementation of pipelines with fan-in and fan-out.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyfather/pipeline)   [![godoc][GoDoc]](https://godoc.org/github.com/hyfather/pipeline)
* [goset](https://github.com/zoumo/goset) **star:17** A useful Set collection implementation for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [typ](https://github.com/gurukami/typ) **star:15** Null Types, Safe primitive type conversion and fetching value from complex structures.   [![godoc][GoDoc]](https://godoc.org/github.com/gurukami/typ)
* [deque](https://github.com/edwingeng/deque) **star:11** A highly optimized double-ended queue.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/deque)
* [go-ef](https://github.com/amallia/go-ef) **star:11** A Go implementation of the Elias-Fano encoding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/amallia/go-ef)   [![godoc][GoDoc]](https://godoc.org/github.com/amallia/go-ef)
* [hide](https://github.com/emvi/hide) **star:11** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [dict](https://github.com/srfrog/dict) **star:10** Python-like dictionaries (dict) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [mspm](https://github.com/BlackRabbitt/mspm) **star:8** Multi-String Pattern Matching Algorithm for information retrieval.   [![It hasn't been updated in the last year][Yellow]](https://github.com/BlackRabbitt/mspm)   [![godoc][GoDoc]](https://godoc.org/github.com/BlackRabbitt/mspm)
* [treap](https://github.com/perdata/treap) **star:7** Persistent, fast ordered map using tree heaps.   [![It hasn't been updated in the last year][Yellow]](https://github.com/perdata/treap)   [![godoc][GoDoc]](https://godoc.org/github.com/perdata/treap)
* [null](https://github.com/emvi/null) **star:7** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [set](https://github.com/StudioSol/set) **star:7** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:5** Map with expiring key-value pairs.   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)
* [gofal](https://github.com/xxjwxc/gofal) **star:3** fractional api for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gofal)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gofal)
* [ptrie](https://github.com/viant/ptrie) **star:3** An implementation of prefix tree.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/ptrie)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:28005** Monitoring system and time series database.   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   [![There was an update last week][Green]](https://github.com/prometheus/prometheus)   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [tidb](https://github.com/pingcap/tidb) **star:21749** TiDB is a distributed SQL database. Inspired by the design of Google F1.   [![star > 2000][Awesome]](https://github.com/pingcap/tidb)   [![There was an update last week][Green]](https://github.com/pingcap/tidb)   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/tidb)   [![Contains Chinese documents][CN]](https://github.com/pingcap/tidb)
* [influxdb](https://github.com/influxdb/influxdb) **star:17931** Scalable datastore for metrics, events, and real-time analytics.   [![star > 2000][Awesome]](https://github.com/influxdb/influxdb)   [![There was an update last week][Green]](https://github.com/influxdb/influxdb)   [![godoc][GoDoc]](https://godoc.org/github.com/influxdb/influxdb)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:17455** Scalable, Geo-Replicated, Transactional Datastore.   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   [![There was an update last week][Green]](https://github.com/cockroachdb/cockroach)   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [dgraph](https://github.com/dgraph-io/dgraph) **star:11902** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   [![There was an update last week][Green]](https://github.com/dgraph-io/dgraph)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [bolt](https://github.com/boltdb/bolt) **star:10323** Low-level key/value database for Go.   [![star > 2000][Awesome]](https://github.com/boltdb/bolt)   [![It hasn't been updated in the last year][Yellow]](https://github.com/boltdb/bolt)   [![godoc][GoDoc]](https://godoc.org/github.com/boltdb/bolt)   [![Archived][Archived]](https://github.com/boltdb/bolt)
* [groupcache](https://github.com/golang/groupcache) **star:8045** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:6994** Fast key-value store in Go.   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   [![There was an update last week][Green]](https://github.com/dgraph-io/badger)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:5110** The lightweight, distributed, relational database built on SQLite.   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   [![There was an update last week][Green]](https://github.com/rqlite/rqlite)   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3418** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [go-cache](https://github.com/pmylund/go-cache) **star:3304** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3172** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   [![star > 2000][Awesome]](https://github.com/siddontang/ledisdb)   [![There was an update last week][Green]](https://github.com/siddontang/ledisdb)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/ledisdb)
* [BigCache](https://github.com/allegro/bigcache) **star:2985** Efficient key/value cache for gigabytes of data.   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![There was an update last week][Green]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [buntdb](https://github.com/tidwall/buntdb) **star:2558** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2423** Your NoSQL database powered by Golang.   [![star > 2000][Awesome]](https://github.com/HouzuoGuo/tiedot)   [![godoc][GoDoc]](https://godoc.org/github.com/HouzuoGuo/tiedot)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1478** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   [![There was an update last week][Green]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [cache2go](https://github.com/muesli/cache2go) **star:1162** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:1038** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   [![Contains Chinese documents][CN]](https://github.com/xujiajun/nutsdb)
* [GCache](https://github.com/bluele/gcache) **star:1002** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:956** CovenantSQL is a SQL database on blockchain.   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:822** Home-grown disk-backed key-value store.   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [moss](https://github.com/couchbase/moss) **star:738** Moss is a simple LSM key-value storage engine written in 100% Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/moss)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:602** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
* [eliasdb](https://github.com/krotik/eliasdb) **star:555** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [levigo](https://github.com/jmhodges/levigo) **star:375** Levigo is a Go wrapper for LevelDB.   [![There was an update last week][Green]](https://github.com/jmhodges/levigo)   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [pudge](https://github.com/recoilme/pudge) **star:236** Fast and simple  key/value store written using Go's standard library.   [![There was an update last week][Green]](https://github.com/recoilme/pudge)   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [Bitcask](https://github.com/prologic/bitcask) **star:221** Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).   [![godoc][GoDoc]](https://godoc.org/github.com/prologic/bitcask)
* [piladb](https://github.com/fern4lvarez/piladb) **star:170** Lightweight RESTful database engine based on stack data structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fern4lvarez/piladb)   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Vasto](https://github.com/chrislusf/vasto) **star:164** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [Kivik](https://github.com/go-kivik/kivik) **star:131** Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.   [![There was an update last week][Green]](https://github.com/go-kivik/kivik)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kivik/kivik)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:94** Key-value store with persistence.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:81** Tiny flat file JSON store.   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:46** RESTful caching micro-service backed by Couchbase server.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [bcache](https://github.com/iwanbk/bcache) **star:37** Eventually consistent distributed in-memory  cache Go library.   [![godoc][GoDoc]](https://godoc.org/github.com/iwanbk/bcache)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:31** BigCache with clustering support and individual item expiration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][GoDoc]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
* [cache](https://github.com/akyoto/cache) **star:27** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)
* [Coffer](https://github.com/claygod/coffer) **star:14** Simple ACID key-value database that supports transactions.   [![There was an update last week][Green]](https://github.com/claygod/coffer)   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/coffer)
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:13** Key-value store for temporary items.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/tempdb)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/tempdb)
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:10** Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kapitan-k/gorocksdb)   [![godoc][GoDoc]](https://godoc.org/github.com/kapitan-k/gorocksdb)

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:3333** Database migrations. CLI and Golang library.   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![There was an update last week][Green]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1537** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [skeema](https://github.com/skeema/skeema) **star:484** Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.   [![There was an update last week][Green]](https://github.com/skeema/skeema)   [![godoc][GoDoc]](https://godoc.org/github.com/skeema/skeema)
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:379** Database schema migration helper for Gorm ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [goose](https://github.com/steinbacher/goose) **star:125** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   [![It hasn't been updated in the last year][Yellow]](https://github.com/steinbacher/goose)   [![godoc][GoDoc]](https://godoc.org/github.com/steinbacher/goose)
* [darwin](https://github.com/GuiaBolso/darwin) **star:95** Database schema evolution library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/GuiaBolso/darwin)
* [migrator](https://github.com/lopezator/migrator) **star:82** Dead simple Go database migration library.   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:36** A Go package to help write migrations with go-pg/pg.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [gondolier](https://github.com/emvi/gondolier) **star:26** Database migration library using struct decorators.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/gondolier)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:20** Django style fixtures for Golang's excellent built-in database/sql library.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-fixtures)
* [avro](https://github.com/khezen/avro) **star:8** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)
* [schema](https://github.com/adlio/schema) **star:4** Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/schema)

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:9162** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   [![There was an update last week][Green]](https://github.com/youtube/vitess)   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [pgweb](https://github.com/sosedoff/pgweb) **star:6175** Web-based PostgreSQL database browser.   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![There was an update last week][Green]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [kingshard](https://github.com/flike/kingshard) **star:4900** kingshard is a high performance proxy for MySQL powered by Golang.   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![There was an update last week][Green]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   [![Contains Chinese documents][CN]](https://github.com/flike/kingshard)
* [orchestrator](https://github.com/github/orchestrator) **star:3249** MySQL replication topology manager & visualizer.   [![star > 2000][Awesome]](https://github.com/github/orchestrator)   [![There was an update last week][Green]](https://github.com/github/orchestrator)   [![godoc][GoDoc]](https://godoc.org/github.com/github/orchestrator)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2679** Sync your MySQL data into Elasticsearch automatically.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [pREST](https://github.com/nuveo/prest) **star:2158** Serve a RESTful API from any PostgreSQL database.   [![star > 2000][Awesome]](https://github.com/nuveo/prest)   [![godoc][GoDoc]](https://godoc.org/github.com/nuveo/prest)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:2106** Go toolset to handle MySQL protocol and replication.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql)   [![There was an update last week][Green]](https://github.com/siddontang/go-mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [chproxy](https://github.com/Vertamedia/chproxy) **star:347** HTTP proxy for ClickHouse database.   [![There was an update last week][Green]](https://github.com/Vertamedia/chproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/Vertamedia/chproxy)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:162** Collects small insterts and sends big requests to ClickHouse servers.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [myreplication](https://github.com/2tvenom/myreplication) **star:151** MySql binary log replication listener. Supports statement and row based replication.   [![It hasn't been updated in the last year][Yellow]](https://github.com/2tvenom/myreplication)   [![godoc][GoDoc]](https://godoc.org/github.com/2tvenom/myreplication)
* [octillery](https://github.com/knocknote/octillery) **star:68** Go package for sharding databases ( Supports every ORM or raw SQL ).   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [dbbench](https://github.com/sj14/dbbench) **star:30** Database benchmarking tool with support for several databases and scripts.   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [prep](https://github.com/hexdigest/prep) **star:25** Use prepared SQL statements without changing your code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/prep)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/prep)
* [datagen](https://github.com/codingconcepts/datagen) **star:13** A fast data generator that's multi-table aware and supports multi-row DML.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/datagen)
* [rwdb](https://github.com/andizzle/rwdb) **star:10** rwdb provides read replica capability for multiple database servers setup.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andizzle/rwdb)   [![godoc][GoDoc]](https://godoc.org/github.com/andizzle/rwdb)
* [bucket](https://github.com/PumpkinSeed/bucket) **star:6** Optimized data structure framework for Couchbase specialized on one bucket usage.   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/bucket)

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2567** Go library that helps you build SQL queries.   [![star > 2000][Awesome]](https://github.com/Masterminds/squirrel)   [![There was an update last week][Green]](https://github.com/Masterminds/squirrel)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/squirrel)
* [xo](https://github.com/knq/xo) **star:2306** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   [![star > 2000][Awesome]](https://github.com/knq/xo)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/xo)
* [gendry](https://github.com/didi/gendry) **star:910** Non-invasive SQL builder and powerful data binder.   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   [![Contains Chinese documents][CN]](https://github.com/didi/gendry)
* [goqu](https://github.com/doug-martin/goqu) **star:702** Idiomatic SQL builder and query library.   [![There was an update last week][Green]](https://github.com/doug-martin/goqu)   [![godoc][GoDoc]](https://godoc.org/github.com/doug-martin/goqu)
* [Dotsql](https://github.com/gchaincl/dotsql) **star:471** Go library that helps you keep sql files in one place and use them with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/dotsql)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:444** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![There was an update last week][Green]](https://github.com/go-ozzo/ozzo-dbx)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
* [sqrl](https://github.com/elgris/sqrl) **star:193** SQL query builder, fork of Squirrel with improved performance.   [![godoc][GoDoc]](https://godoc.org/github.com/elgris/sqrl)
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [jet](https://github.com/go-jet/jet) **star:179** Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.   [![There was an update last week][Green]](https://github.com/go-jet/jet)   [![godoc][GoDoc]](https://godoc.org/github.com/go-jet/jet)
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [igor](https://github.com/galeone/igor) **star:78** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [dbq](https://github.com/rocketlaunchr/dbq) **star:75** Zero boilerplate database operations for Go.   [![There was an update last week][Green]](https://github.com/rocketlaunchr/dbq)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)
* [godbal](https://github.com/xujiajun/godbal) **star:50** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [sqlf](https://github.com/leporo/sqlf) **star:5** Fast SQL query builder.   [![godoc][GoDoc]](https://godoc.org/github.com/leporo/sqlf)
* [qry](https://github.com/HnH/qry) **star:3** Tool that generates constants from files with raw SQL queries.   [![godoc][GoDoc]](https://godoc.org/github.com/HnH/qry)
* [scaneo](https://github.com/variadico/scaneo)  Generate Go code to convert database rows into arbitrary structs.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:8716** MySQL driver for Go.   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [pq](https://github.com/lib/pq) **star:5495** Pure Go Postgres driver for database/sql.   [![star > 2000][Awesome]](https://github.com/lib/pq)   [![There was an update last week][Green]](https://github.com/lib/pq)   [![godoc][GoDoc]](https://godoc.org/github.com/lib/pq)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3656** SQLite3 driver for go that uses database/sql.   [![star > 2000][Awesome]](https://github.com/mattn/go-sqlite3)   [![There was an update last week][Green]](https://github.com/mattn/go-sqlite3)
    * [pgx](https://github.com/jackc/pgx) **star:2215** PostgreSQL driver supporting features beyond those exposed by database/sql.   [![star > 2000][Awesome]](https://github.com/jackc/pgx)   [![There was an update last week][Green]](https://github.com/jackc/pgx)   [![godoc][GoDoc]](https://godoc.org/github.com/jackc/pgx)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1108** Microsoft MSSQL driver for Go.   [![There was an update last week][Green]](https://github.com/denisenkom/go-mssqldb)   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8) **star:425** Oracle driver for go that uses database/sql.   [![There was an update last week][Green]](https://github.com/mattn/go-oci8)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-oci8)
    * [goracle](https://github.com/go-goracle/goracle) **star:281** Oracle driver for Go, using the ODPI-C driver.   [![There was an update last week][Green]](https://github.com/go-goracle/goracle)
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:112** Firebird RDBMS SQL driver for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/nakagami/firebirdsql)
    * [gofreetds](https://github.com/minus5/gofreetds) **star:94** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).   [![godoc][GoDoc]](https://godoc.org/github.com/minus5/gofreetds)
    * [go-adodb](https://github.com/mattn/go-adodb) **star:93** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:43** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [bgc](https://github.com/viant/bgc) **star:12** Datastore Connectivity for BigQuery for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:7536** Redis client for Golang.   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   [![There was an update last week][Green]](https://github.com/go-redis/redis)   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo) **star:6779** Redigo is a Go client for the Redis database.   [![star > 2000][Awesome]](https://github.com/gomodule/redigo)   [![godoc][GoDoc]](https://godoc.org/github.com/gomodule/redigo)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:3719** Official MongoDB driver for the Go language.   [![star > 2000][Awesome]](https://github.com/mongodb/mongo-go-driver)   [![There was an update last week][Green]](https://github.com/mongodb/mongo-go-driver)   [![godoc][GoDoc]](https://godoc.org/github.com/mongodb/mongo-go-driver)
    * [mgo](https://github.com/globalsign/mgo) **star:1713** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   [![godoc][GoDoc]](https://godoc.org/github.com/globalsign/mgo)
    * [gorethink](https://github.com/dancannon/gorethink) **star:1489** Go language driver for RethinkDB.   [![godoc][GoDoc]](https://godoc.org/github.com/dancannon/gorethink)
    * [neoism](https://github.com/jmcvetta/neoism) **star:363** Neo4j client for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [redeo](https://github.com/bsm/redeo) **star:362** Redis-protocol compatible TCP servers/services.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redeo)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:313** Aerospike client in Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [gocb](https://github.com/couchbase/gocb) **star:307** Official Couchbase Go SDK.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/gocb)
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:297** Couchbase client in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:106** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:73** Neo4j REST Client in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davemeehan/Neo4j-GO)   [![godoc][GoDoc]](https://godoc.org/github.com/davemeehan/Neo4j-GO)
    * [godis](https://github.com/piaohao/godis) **star:67** redis client implement by golang, inspired by jedis.   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [arangolite](https://github.com/solher/arangolite) **star:66** Lightweight golang driver for ArangoDB.   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [dynago](https://github.com/underarmour/dynago) **star:66** Dynago is a principle of least surprise client for DynamoDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/underarmour/dynago)   [![godoc][GoDoc]](https://godoc.org/github.com/underarmour/dynago)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:33** Go client library for Pilosa.   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [forestdb](https://github.com/couchbase/goforestdb) **star:28** Go bindings for ForestDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/couchbase/goforestdb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/goforestdb)
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [neo4j](https://github.com/cihangir/neo4j) **star:24** Neo4j Rest API Bindings for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cihangir/neo4j)   [![godoc][GoDoc]](https://godoc.org/github.com/cihangir/neo4j)
    * [xredis](https://github.com/shomali11/xredis) **star:9** Typesafe, customizable, clean & easy to use Redis client.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)
    * [godscache](https://github.com/defcronyke/godscache) **star:6** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.   [![godoc][GoDoc]](https://godoc.org/github.com/defcronyke/godscache)
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [asc](https://github.com/viant/asc) **star:4** Datastore Connectivity for Aerospike for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/asc)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:6185** Modern text indexing library for go.   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [riot](https://github.com/go-ego/riot) **star:4933** Go Open Source, Distributed, Simple and efficient Search Engine.   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   [![Contains Chinese documents][CN]](https://github.com/go-ego/riot)
    * [elastic](https://github.com/olivere/elastic) **star:4565** Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/olivere/elastic)   [![There was an update last week][Green]](https://github.com/olivere/elastic)   [![godoc][GoDoc]](https://godoc.org/github.com/olivere/elastic)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:1998** Official Elasticsearch client for Go.   [![There was an update last week][Green]](https://github.com/elastic/go-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/elastic/go-elasticsearch)
    * [elastigo](https://github.com/mattbaird/elastigo) **star:950** Elasticsearch client library.   [![godoc][GoDoc]](https://godoc.org/github.com/mattbaird/elastigo)
    * [elasticsql](https://github.com/cch123/elasticsql) **star:464** Convert sql to elasticsearch dsl in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cch123/elasticsql)
    * [skizze](https://github.com/seiflotfy/skizze) **star:68** probabilistic data-structures service and storage.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/skizze)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OwnLocal/goes)   [![godoc][GoDoc]](https://godoc.org/github.com/OwnLocal/goes)

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:13023** Graph database with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/google/cayley)   [![There was an update last week][Green]](https://github.com/google/cayley)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cayley)
    * [gokv](https://github.com/philippgille/gokv) **star:143** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [cachego](https://github.com/fabiorphp/cachego) **star:111** Golang Cache component for multiple drivers.   [![godoc][GoDoc]](https://godoc.org/github.com/fabiorphp/cachego)
    * [dsc](https://github.com/viant/dsc) **star:13** Datastore connectivity for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2324** Now is a time toolkit for golang.   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [dateparse](https://github.com/araddon/dateparse) **star:950** Parse date's without knowing format in advance.   [![godoc][GoDoc]](https://godoc.org/github.com/araddon/dateparse)
* [carbon](https://github.com/uniplaces/carbon) **star:372** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   [![godoc][GoDoc]](https://godoc.org/github.com/uniplaces/carbon)
* [durafmt](https://github.com/hako/durafmt) **star:270** Time duration formatting library for Go.   [![There was an update last week][Green]](https://github.com/hako/durafmt)   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [timeutil](https://github.com/leekchan/timeutil) **star:171** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/timeutil)
* [iso8601](https://github.com/relvacode/iso8601) **star:68** Efficiently parse ISO8601 date-times without regex.   [![godoc][GoDoc]](https://godoc.org/github.com/relvacode/iso8601)
* [timespan](https://github.com/SaidinWoT/timespan) **star:66** For interacting with intervals of time, defined as a start time and a duration.   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:66** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/go-persian-calendar)
* [date](https://github.com/rickb777/date) **star:33** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.   [![godoc][GoDoc]](https://godoc.org/github.com/rickb777/date)
* [feiertage](https://github.com/wlbr/feiertage) **star:24** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:16** Calculate the sunrise and sunset times for a given location.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nathan-osman/go-sunrise)   [![godoc][GoDoc]](https://godoc.org/github.com/nathan-osman/go-sunrise)
* [kair](https://github.com/GuilhermeCaruso/kair) **star:12** Date and Time - Golang Formatting Library.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:9** Nullable `time.Time`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/nulltime)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/nulltime)
* [tuesday](https://github.com/osteele/tuesday) **star:7** Ruby-compatible Strftime function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/osteele/tuesday)   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/tuesday)
* [cronrange](https://github.com/1set/cronrange) **star:5** Parses Cron-style time range expressions, checks if the given time is within any ranges.   [![godoc][GoDoc]](https://godoc.org/github.com/1set/cronrange)
* [strftime](https://github.com/awoodbeck/strftime) **star:4** C99-compatible strftime formatter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awoodbeck/strftime)   [![godoc][GoDoc]](https://godoc.org/github.com/awoodbeck/strftime)
* [go-week](https://github.com/stoewer/go-week) **star:2** An efficient package to work with ISO8601 week dates.   [![godoc][GoDoc]](https://godoc.org/github.com/stoewer/go-week)

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:15625** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [grpc-go](https://github.com/grpc/grpc-go) **star:10187** The Go language implementation of gRPC. HTTP/2 based RPC.   [![star > 2000][Awesome]](https://github.com/grpc/grpc-go)   [![There was an update last week][Green]](https://github.com/grpc/grpc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/grpc/grpc-go)
* [micro](https://github.com/micro/micro) **star:7247** Pluggable microservice toolkit and distributed systems platform.   [![star > 2000][Awesome]](https://github.com/micro/micro)   [![There was an update last week][Green]](https://github.com/micro/micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [NATS](https://github.com/nats-io/gnatsd) **star:6976** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.   [![star > 2000][Awesome]](https://github.com/nats-io/gnatsd)   [![There was an update last week][Green]](https://github.com/nats-io/gnatsd)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/gnatsd)
* [rpcx](https://github.com/smallnest/rpcx) **star:4151** Distributed pluggable RPC service framework like alibaba Dubbo.   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   [![There was an update last week][Green]](https://github.com/smallnest/rpcx)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [tendermint](https://github.com/tendermint/tendermint) **star:3390** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   [![There was an update last week][Green]](https://github.com/tendermint/tendermint)   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [torrent](https://github.com/anacrolix/torrent) **star:3157** BitTorrent client package.   [![star > 2000][Awesome]](https://github.com/anacrolix/torrent)   [![There was an update last week][Green]](https://github.com/anacrolix/torrent)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/torrent)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [raft](https://github.com/hashicorp/raft) **star:3072** Golang implementation of the Raft consensus protocol, by HashiCorp.   [![star > 2000][Awesome]](https://github.com/hashicorp/raft)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/raft)
* [dragonboat](https://github.com/lni/dragonboat) **star:2792** A feature complete and high performance multi-group Raft library in Go.   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   [![There was an update last week][Green]](https://github.com/lni/dragonboat)   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   [![Contains Chinese documents][CN]](https://github.com/lni/dragonboat)
* [glow](https://github.com/chrislusf/glow) **star:2695** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [gleam](https://github.com/chrislusf/gleam) **star:2295** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [KrakenD](https://github.com/devopsfaith/krakend) **star:2138** Ultra performant API Gateway framework with middlewares.   [![star > 2000][Awesome]](https://github.com/devopsfaith/krakend)   [![There was an update last week][Green]](https://github.com/devopsfaith/krakend)   [![godoc][GoDoc]](https://godoc.org/github.com/devopsfaith/krakend)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2127** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   [![There was an update last week][Green]](https://github.com/emitter-io/emitter)   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [liftbridge](https://github.com/liftbridge-io/liftbridge) **star:1376** Lightweight, fault-tolerant message streams for NATS.   [![godoc][GoDoc]](https://godoc.org/github.com/liftbridge-io/liftbridge)
* [hprose](https://github.com/hprose/hprose-golang) **star:1053** Very newbility RPC Library, support 25+ languages now.   [![godoc][GoDoc]](https://godoc.org/github.com/hprose/hprose-golang)   [![Contains Chinese documents][CN]](https://github.com/hprose/hprose-golang)
* [ringpop-go](https://github.com/uber/ringpop-go) **star:592** Scalable, fault-tolerant application-layer sharding for Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/uber/ringpop-go)
* [gorpc](https://github.com/valyala/gorpc) **star:570** Simple, fast and scalable RPC library for high load.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [go-health](https://github.com/InVisionApp/go-health) **star:514** Library for enabling asynchronous dependency health checks in your service.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [rain](https://github.com/cenkalti/rain) **star:367** BitTorrent client and library.   [![godoc][GoDoc]](https://godoc.org/github.com/cenkalti/rain)
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [digota](https://github.com/digota/digota) **star:321** grpc ecommerce microservice.   [![It hasn't been updated in the last year][Yellow]](https://github.com/digota/digota)   [![godoc][GoDoc]](https://godoc.org/github.com/digota/digota)
* [sleuth](https://github.com/ursiform/sleuth) **star:305** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ursiform/sleuth)   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) **star:274** A library built to provide support for defining async service health checks for golang services.   [![godoc][GoDoc]](https://godoc.org/github.com/AppsFlyer/go-sundheit)
* [go-jump](https://github.com/dgryski/go-jump) **star:270** Port of Google's "Jump" Consistent Hash function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgryski/go-jump)   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [consistent](https://github.com/buraksezer/consistent) **star:246** Consistent hashing with bounded loads.   [![godoc][GoDoc]](https://godoc.org/github.com/buraksezer/consistent)
* [dht](https://github.com/anacrolix/dht) **star:140** BitTorrent Kademlia DHT implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:116** The jsonrpc package helps implement of JSON-RPC 2.0.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:110** JSON-RPC 2.0 HTTP client implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [redis-lock](https://github.com/bsm/redislock) **star:73** Simplified distributed locking implementation using Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redislock)
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:58** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/svcavallar/celeriac.v1)   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:46** A revamped Google's jump consistent hash.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [drmaa](https://github.com/dgruber/drmaa) **star:28** Job submission library for cluster schedulers based on the DRMAA standard.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgruber/drmaa)   [![godoc][GoDoc]](https://godoc.org/github.com/dgruber/drmaa)
* [flowgraph](https://github.com/vectaport/flowgraph) **star:24** flow-based programming package.   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [outboxer](https://github.com/italolelis/outboxer) **star:20** Outboxer is a go library that implements the outbox pattern.   [![There was an update last week][Green]](https://github.com/italolelis/outboxer)   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:10** A library for using DynamoDB as an atomic counter.   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:5656** Email and SMTP testing with web and API interface.   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [hermes](https://github.com/matcornic/hermes) **star:1713** Golang package that generates clean, responsive HTML e-mails.   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1167** A robust and flexible email library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:837** IMAP library for clients and servers.   [![There was an update last week][Green]](https://github.com/emersion/go-imap)   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:544** SendGrid's Go library for sending email.   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [mailgun-go](https://github.com/mailgun/mailgun-go) **star:422** Go library for sending mail with the Mailgun API.   [![There was an update last week][Green]](https://github.com/mailgun/mailgun-go)   [![godoc][GoDoc]](https://godoc.org/github.com/mailgun/mailgun-go)
* [Hectane](https://github.com/hectane/hectane) **star:171** Lightweight SMTP client providing an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [douceur](https://github.com/aymerick/douceur) **star:165** CSS inliner for your HTML emails.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/douceur)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [go-message](https://github.com/emersion/go-message) **star:126** Streaming library for the Internet Message Format and mail messages.   [![There was an update last week][Green]](https://github.com/emersion/go-message)   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-message)
* [smtp](https://github.com/mailhog/smtp) **star:53** SMTP server protocol state machine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mailhog/smtp)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/smtp)
* [go-dkim](https://github.com/toorop/go-dkim) **star:52** DKIM library, to sign & verify email.   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:41** Inline styling for HTML mail in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [mailchain](https://github.com/mailchain/mailchain) **star:31** Send encrypted emails to blockchain addresses written in Go.   [![There was an update last week][Green]](https://github.com/mailchain/mailchain)   [![godoc][GoDoc]](https://godoc.org/github.com/mailchain/mailchain)
* [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:9** Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-simple-mail)

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) **star:4971** JavaScript interpreter written in Go.   [![star > 2000][Awesome]](https://github.com/robertkrimen/otto)   [![godoc][GoDoc]](https://godoc.org/github.com/robertkrimen/otto)
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:3157** Lua 5.1 VM and compiler written in Go.   [![star > 2000][Awesome]](https://github.com/yuin/gopher-lua)   [![There was an update last week][Green]](https://github.com/yuin/gopher-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/yuin/gopher-lua)
* [go-lua](https://github.com/Shopify/go-lua) **star:1742** Port of the Lua 5.2 VM to pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [tengo](https://github.com/d5/tengo) **star:1468** Bytecode compiled script language for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [expr](https://github.com/antonmedv/expr) **star:997** an engine that can evaluate expressions.   [![There was an update last week][Green]](https://github.com/antonmedv/expr)   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [anko](https://github.com/mattn/anko) **star:971** Scriptable interpreter written in Go.   [![There was an update last week][Green]](https://github.com/mattn/anko)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-python](https://github.com/sbinet/go-python) **star:967** naive go bindings to the CPython C-API.   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [go-php](https://github.com/deuill/go-php) **star:717** PHP bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/deuill/go-php)   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:673** Duktape JavaScript engine bindings for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [golua](https://github.com/aarzilli/golua) **star:455** Go bindings for Lua C API.
* [gisp](https://github.com/jcla1/gisp) **star:432** Simple LISP in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jcla1/gisp)   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [cel-go](https://github.com/google/cel-go) **star:329** Fast, portable, non-Turing complete expression evaluation with gradual typing.   [![There was an update last week][Green]](https://github.com/google/cel-go)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cel-go)
* [gval](https://github.com/PaesslerAG/gval) **star:167** A highly customizable expression language written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [gentee](https://github.com/gentee/gentee) **star:40** Embeddable scripting programming language.   [![There was an update last week][Green]](https://github.com/gentee/gentee)   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)
* [binder](https://github.com/alexeyco/binder) **star:32** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexeyco/binder)   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [purl](https://github.com/ian-kent/purl) **star:29** Perl 5.18.2 embedded in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/purl)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/purl)
* [ngaro](https://github.com/db47h/ngaro) **star:19** Embeddable Ngaro VM implementation enabling scripting in Retro.   [![It hasn't been updated in the last year][Yellow]](https://github.com/db47h/ngaro)   [![godoc][GoDoc]](https://godoc.org/github.com/db47h/ngaro)

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) **star:5318** Package that provides simple error handling primitives.   [![star > 2000][Awesome]](https://github.com/pkg/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/errors)
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:792** Go (golang) package for representing a list of errors as a single error.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [errorx](https://github.com/joomcode/errorx) **star:596** A feature rich error package with stack traces, composition of errors and more.   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:528** Golang errors with stack trace and source fragments.   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [errlog](https://github.com/snwfdhmp/errlog) **star:254** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   [![There was an update last week][Green]](https://github.com/snwfdhmp/errlog)   [![godoc][GoDoc]](https://godoc.org/github.com/snwfdhmp/errlog)
* [emperror](https://github.com/emperror/emperror) **star:89** Error handling tools and best practices for Go libraries and applications.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/emperror)
* [errors](https://github.com/emperror/errors) **star:31** Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/errors)
* [werr](https://github.com/txgruppi/werr) **star:12** Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/werr)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/werr)
* [errors](https://github.com/neuronlabs/errors) **star:3** Simple golang error handling with classification primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/neuronlabs/errors)
* [Falcon](https://github.com/SonicRoshan/falcon) **star:2** A Simple Yet Highly Powerful Package For Error Handling.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/falcon)

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2428** FileSystem Abstraction System for Go.   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/hhrutter/pdfcpu) **star:1217** PDF processor.   [![There was an update last week][Green]](https://github.com/hhrutter/pdfcpu)   [![godoc][GoDoc]](https://godoc.org/github.com/hhrutter/pdfcpu)
* [notify](https://github.com/rjeczalik/notify) **star:514** File system event notification library with simple API, similar to os/signal.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/notify)
* [bigfile](https://github.com/bigfile/bigfile) **star:91** A file transfer system, support to manage files with http api, rpc call and ftp client.   [![godoc][GoDoc]](https://godoc.org/github.com/bigfile/bigfile)   [![Contains Chinese documents][CN]](https://github.com/bigfile/bigfile)
* [opc](https://github.com/qmuntal/opc) **star:63** Load Open Packaging Conventions (OPC) files for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:53** Load csv file using tag.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-csv-tag)
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [skywalker](https://github.com/dixonwille/skywalker) **star:52** Package to allow one to concurrently go through a filesystem with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/skywalker)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/skywalker)
* [tarfs](https://github.com/posener/tarfs) **star:38** Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/posener/tarfs)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/tarfs)
* [vfs](https://github.com/C2FO/vfs) **star:37** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [afs](https://github.com/viant/afs) **star:19** Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.   [![There was an update last week][Green]](https://github.com/viant/afs)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/afs)
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:17** Load gtfs files in go.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:15** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![There was an update last week][Green]](https://github.com/barasher/go-exiftool)   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)
* [checksum](https://github.com/codingsince1985/checksum) **star:14** Compute message digest, like MD5 and SHA256, for large files.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:13** Copy files for humans.   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
* [flop](https://github.com/homedepot/flop) **star:12** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [parquet](https://github.com/parsyl/parquet) **star:5** Read and write [parquet](https://parquet.apache.org) files.   [![godoc][GoDoc]](https://godoc.org/github.com/parsyl/parquet)

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:1828** Arbitrary-precision fixed-point decimal numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-money](https://github.com/rhymond/go-money) **star:669** Implementation of Fowler's Money pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [go-finance](https://github.com/FlashBoys/go-finance) **star:536** Comprehensive financial markets data in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/FlashBoys/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/FlashBoys/go-finance)
* [accounting](https://github.com/leekchan/accounting) **star:521** money and currency formatting for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [techan](https://github.com/sdcoffey/techan) **star:196** Technical analysis library with advanced market analysis and trading strategies.   [![There was an update last week][Green]](https://github.com/sdcoffey/techan)   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [orderbook](https://github.com/i25959341/orderbook) **star:105** Matching Engine for Limit Order Book in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/i25959341/orderbook)
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:70** Query OFX servers and/or parse the responses (with example command-line client).   [![godoc][GoDoc]](https://godoc.org/github.com/aclindsa/ofxgo)
* [vat](https://github.com/dannyvankooten/vat) **star:62** VAT number validation & EU VAT rates.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dannyvankooten/vat)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/vat)
* [transaction](https://github.com/claygod/transaction) **star:61** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [go-finance](https://github.com/alpeb/go-finance) **star:48** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alpeb/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/alpeb/go-finance)
* [currency](https://github.com/bnkamalesh/currency) **star:14** High performant & accurate currency computation package.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)
* [go-finance](https://github.com/pieterclaerhout/go-finance) **star:2** Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-finance)

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) **star:1007** CSRF protection middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/nosurf)
* [binding](https://github.com/mholt/binding) **star:759** Binds form and JSON data from net/http Request to struct.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/binding)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:484** CSRF protection for Go web applications & services.   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [form](https://github.com/go-playground/form) **star:377** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [conform](https://github.com/leebenson/conform) **star:181** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/leebenson/conform)
* [formam](https://github.com/monoculum/formam) **star:134** decode form's values into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:106** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/forms)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)
* [bind](https://github.com/robfig/bind) **star:24** Bind form data to any Go values.   [![It hasn't been updated in the last year][Yellow]](https://github.com/robfig/bind)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/bind)

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1103** Useful collection of helpfully functional Go collection utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/tobyhede/go-underscore)
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:126** Monad, Functional Programming features for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TeaEntityLab/fpGo)   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:56** Functional Experiment in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)

## Game Development

*Awesome game development libraries.*

* [Leaf](https://github.com/name5566/leaf) **star:3275** Lightweight game server framework.   [![star > 2000][Awesome]](https://github.com/name5566/leaf)   [![godoc][GoDoc]](https://godoc.org/github.com/name5566/leaf)   [![Contains Chinese documents][CN]](https://github.com/name5566/leaf)
* [Pixel](https://github.com/faiface/pixel) **star:2620** Hand-crafted 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/faiface/pixel)   [![godoc][GoDoc]](https://godoc.org/github.com/faiface/pixel)
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:2307** dead simple 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/hajimehoshi/ebiten)   [![There was an update last week][Green]](https://github.com/hajimehoshi/ebiten)   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/ebiten)
* [goworld](https://github.com/xiaonanln/goworld) **star:1319** Scalable game server engine, featuring space-entity framework and hot-swapping.   [![There was an update last week][Green]](https://github.com/xiaonanln/goworld)   [![godoc][GoDoc]](https://godoc.org/github.com/xiaonanln/goworld)   [![Contains Chinese documents][CN]](https://github.com/xiaonanln/goworld)
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1233** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).   [![There was an update last week][Green]](https://github.com/veandco/go-sdl2)
* [engo](https://github.com/EngoEngine/engo) **star:1132** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [nano](https://github.com/lonng/nano) **star:1118** Lightweight, facility, high performance golang based game server framework.   [![godoc][GoDoc]](https://godoc.org/github.com/lonng/nano)   [![Contains Chinese documents][CN]](https://github.com/lonng/nano)
* [gonet](https://github.com/xtaci/gonet) **star:1071** Game server skeleton implemented with golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xtaci/gonet)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [termloop](https://github.com/JoelOtter/termloop) **star:1059** Terminal-based game engine for Go, built on top of Termbox.   [![godoc][GoDoc]](https://godoc.org/github.com/JoelOtter/termloop)
* [g3n](https://github.com/g3n/engine) **star:871** Go 3D Game Engine.   [![There was an update last week][Green]](https://github.com/g3n/engine)   [![godoc][GoDoc]](https://godoc.org/github.com/g3n/engine)
* [Oak](https://github.com/oakmound/oak) **star:690** Pure Go game engine.   [![godoc][GoDoc]](https://godoc.org/github.com/oakmound/oak)
* [Azul3D](https://github.com/azul3d/engine) **star:438** 3D game engine written in Go.
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:413** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.   [![There was an update last week][Green]](https://github.com/gen2brain/raylib-go)
* [Pitaya](https://github.com/topfreegames/pitaya) **star:395** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/pitaya)
* [go-astar](https://github.com/beefsack/go-astar) **star:340** Go implementation of the A\* path finding algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beefsack/go-astar)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-astar)
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:315** 2d game engine written in Go working on OpenGL.   [![godoc][GoDoc]](https://godoc.org/github.com/vova616/GarageEngine)
* [go3d](https://github.com/ungerik/go3d) **star:170** Performance oriented 2D/3D math package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go3d)
* [glop](https://github.com/runningwild/glop) **star:76** Glop (Game Library Of Power) is a fairly simple cross-platform game library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/runningwild/glop)
* [go-collada](https://github.com/GlenKelley/go-collada) **star:14** Go package for working with the Collada file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GlenKelley/go-collada)   [![godoc][GoDoc]](https://godoc.org/github.com/GlenKelley/go-collada)

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1899** .NET LINQ-like query methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
* [jennifer](https://github.com/dave/jennifer) **star:1392** Generate arbitrary Go code without templates.   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [gen](https://github.com/clipperhouse/gen) **star:1060** Code generation tool for ‘generics’-like functionality.   [![It hasn't been updated in the last year][Yellow]](https://github.com/clipperhouse/gen)   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [goderive](https://github.com/awalterschulze/goderive) **star:771** Derives functions from input types.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/goderive)
* [GoWrap](https://github.com/hexdigest/gowrap) **star:301** Generate decorators for Go interfaces using simple templates.   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gowrap)
* [interfaces](https://github.com/rjeczalik/interfaces) **star:198** Command line tool for generating interface definitions.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:90** Go preprocessor for package scoped reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/pkgreflect)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/pkgreflect)
* [go-enum](https://github.com/abice/go-enum) **star:89** Code generation for enums from code comments.   [![There was an update last week][Green]](https://github.com/abice/go-enum)   [![godoc][GoDoc]](https://godoc.org/github.com/abice/go-enum)
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:44** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   [![It hasn't been updated in the last year][Yellow]](https://github.com/t0pep0/efaceconv)   [![godoc][GoDoc]](https://godoc.org/github.com/t0pep0/efaceconv)
* [gotype](https://github.com/wzshiming/gotype) **star:26** Golang source code parsing, usage like reflect package.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/gotype)
* [generis](https://github.com/senselogic/GENERIS) **star:20** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
* [go-xray](https://github.com/pieterclaerhout/go-xray) **star:2** Helpers for making the use of reflection easier.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-xray)

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) **star:6549** Geolocation DB with spatial index and realtime geofencing.   [![star > 2000][Awesome]](https://github.com/tidwall/tile38)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/tile38)
* [S2 geometry](https://github.com/golang/geo) **star:967** S2 geometry library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [geocache](https://github.com/melihmucuk/geocache) **star:116** In-memory cache that is suitable for geolocation based applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/melihmucuk/geocache)   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [osm](https://github.com/paulmach/osm) **star:86** Library for reading, writing and working with OpenStreetMap data and APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/osm)
* [WGS84](https://github.com/wroge/wgs84) **star:43** Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).   [![godoc][GoDoc]](https://godoc.org/github.com/wroge/wgs84)
* [geoserver](https://github.com/hishamkaram/geoserver) **star:27** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/geoserver)
* [gismanager](https://github.com/hishamkaram/gismanager) **star:21** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hishamkaram/gismanager)   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)
* [pbf](https://github.com/maguro/pbf) **star:18** OpenStreetMap PBF golang encoder/decoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/maguro/pbf)   [![godoc][GoDoc]](https://godoc.org/github.com/maguro/pbf)

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:8906** Compiler from Go to JavaScript.   [![star > 2000][Awesome]](https://github.com/gopherjs/gopherjs)   [![godoc][GoDoc]](https://godoc.org/github.com/gopherjs/gopherjs)
* [llgo](https://github.com/go-llvm/llgo) **star:1010** LLVM-based compiler for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-llvm/llgo)   [![godoc][GoDoc]](https://godoc.org/github.com/go-llvm/llgo)
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:391** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tardisgo/tardisgo)   [![godoc][GoDoc]](https://godoc.org/github.com/tardisgo/tardisgo)
* [c4go](https://github.com/Konstantin8105/c4go) **star:186** Transpile C code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:22** Transpile FORTRAN 77 code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) **star:2655** A high-performance goroutine pool for golang.   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   [![There was an update last week][Green]](https://github.com/panjf2000/ants)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/ants)
* [goworker](https://github.com/benmanns/goworker) **star:2311** goworker is a Go-based background worker.   [![star > 2000][Awesome]](https://github.com/benmanns/goworker)   [![godoc][GoDoc]](https://godoc.org/github.com/benmanns/goworker)
* [tunny](https://github.com/Jeffail/tunny) **star:1464** Goroutine pool for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/tunny)
* [grpool](https://github.com/ivpusic/grpool) **star:531** Lightweight Goroutine pool.   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [pool](https://github.com/go-playground/pool) **star:519** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [workerpool](https://github.com/gammazero/workerpool) **star:219** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![There was an update last week][Green]](https://github.com/gammazero/workerpool)   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [go-floc](https://github.com/workanator/go-floc) **star:172** Orchestrate goroutines with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/workanator/go-floc)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-floc)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:117** Control goroutines execution order.   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [GoSlaves](https://github.com/themester/GoSlaves) **star:88** Simple and Asynchronous Goroutine pool library.   [![godoc][GoDoc]](https://godoc.org/github.com/themester/GoSlaves)
* [semaphore](https://github.com/kamilsk/semaphore) **star:80** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/semaphore)
* [semaphore](https://github.com/marusama/semaphore) **star:78** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/semaphore)
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:57** manages a resizeable pool of context-aware goroutines to bound concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
* [worker-pool](https://github.com/vardius/worker-pool) **star:51** goworker is a Go simple async worker pool.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/worker-pool)
* [gowp](https://github.com/xxjwxc/gowp) **star:46** gowp is concurrency limiting goroutine pool.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gowp)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gowp)
* [breaker](https://github.com/kamilsk/breaker) **star:45** Flexible mechanism to make execution flow interruptible.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/breaker)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:42** CyclicBarrier for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/marusama/cyclicbarrier)   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/cyclicbarrier)
* [gollback](https://github.com/vardius/gollback) **star:30** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [async](https://github.com/studiosol/async) **star:26** A safe way to execute functions asynchronously, recovering them in case of panic.   [![godoc][GoDoc]](https://godoc.org/github.com/studiosol/async)
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:26** Run functions in parallel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/parallel-fn)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/parallel-fn)
* [threadpool](https://github.com/shettyh/threadpool) **star:23** Golang threadpool implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [artifex](https://github.com/borderstech/artifex) **star:19** Simple in-memory job queue for Golang using worker-based dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/artifex)
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [Hunch](https://github.com/AaronJan/Hunch) **star:17** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   [![godoc][GoDoc]](https://godoc.org/github.com/AaronJan/Hunch)
* [stl](https://github.com/ssgreg/stl) **star:11** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/stl)
* [gohive](https://github.com/loveleshsharma/gohive) **star:7** A highly performant and easy to use Goroutine pool for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/loveleshsharma/gohive)
* [go-trylock](https://github.com/subchen/go-trylock) **star:5** TryLock support on read-write lock for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-trylock)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.   [![godoc][GoDoc]](https://godoc.org/github.com/nikhilsaraf/go-tools)
* [routine](https://github.com/x-mod/routine) **star:4** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![There was an update last week][Green]](https://github.com/x-mod/routine)   [![godoc][GoDoc]](https://godoc.org/github.com/x-mod/routine)
* [queue](https://github.com/AnikHasibul/queue) **star:3** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.   [![godoc][GoDoc]](https://godoc.org/github.com/AnikHasibul/queue)
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) **star:2** Like `sync.WaitGroup` with error handling and concurrency control.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [ui](https://github.com/andlabs/ui) **star:7223** Platform-native GUI library for Go. Cross platform.   [![star > 2000][Awesome]](https://github.com/andlabs/ui)   [![godoc][GoDoc]](https://godoc.org/github.com/andlabs/ui)
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [fyne](https://github.com/fyne-io/fyne) **star:7038** Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows.   [![star > 2000][Awesome]](https://github.com/fyne-io/fyne)   [![There was an update last week][Green]](https://github.com/fyne-io/fyne)   [![godoc][GoDoc]](https://godoc.org/github.com/fyne-io/fyne)
* [qt](https://github.com/therecipe/qt) **star:6686** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   [![star > 2000][Awesome]](https://github.com/therecipe/qt)   [![godoc][GoDoc]](https://godoc.org/github.com/therecipe/qt)
* [webview](https://github.com/zserge/webview) **star:5112** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).   [![star > 2000][Awesome]](https://github.com/zserge/webview)   [![There was an update last week][Green]](https://github.com/zserge/webview)
* [walk](https://github.com/lxn/walk) **star:4030** Windows application library kit for Go.   [![star > 2000][Awesome]](https://github.com/lxn/walk)   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [app](https://github.com/murlokswarm/app) **star:3178** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   [![star > 2000][Awesome]](https://github.com/murlokswarm/app)   [![godoc][GoDoc]](https://godoc.org/github.com/murlokswarm/app)
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:2888** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   [![There was an update last week][Green]](https://github.com/asticode/go-astilectron)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1559** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/sciter-sdk/go-sciter)
* [gotk3](https://github.com/gotk3/gotk3) **star:861** Go bindings for GTK3.   [![There was an update last week][Green]](https://github.com/gotk3/gotk3)   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd) **star:225** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/dtylman/gowd)

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4713** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   [![star > 2000][Awesome]](https://github.com/go-vgo/robotgo)   [![There was an update last week][Green]](https://github.com/go-vgo/robotgo)
* [systray](https://github.com/getlantern/systray) **star:900** Cross platform Go library to place an icon and menu in the notification area.   [![There was an update last week][Green]](https://github.com/getlantern/systray)   [![godoc][GoDoc]](https://godoc.org/github.com/getlantern/systray)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:497** OSX Desktop Notifications library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/gosx-notifier)
* [trayhost](https://github.com/shurcooL/trayhost) **star:165** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:4** Go bindings for libappindicator3 C library.   [![godoc][GoDoc]](https://godoc.org/github.com/dawidd6/go-appindicator)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:2** OSX library to notify about any (pluggable) activity on your machine.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:1** OSX Sleep/Wake notifications in golang.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [imaging](https://github.com/disintegration/imaging) **star:2791** Simple Go image processing package.   [![star > 2000][Awesome]](https://github.com/disintegration/imaging)   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/imaging)
* [gocv](https://github.com/hybridgroup/gocv) **star:2774** Go package for computer vision using OpenCV 3.3+.   [![star > 2000][Awesome]](https://github.com/hybridgroup/gocv)   [![There was an update last week][Green]](https://github.com/hybridgroup/gocv)   [![godoc][GoDoc]](https://godoc.org/github.com/hybridgroup/gocv)
* [imaginary](https://github.com/h2non/imaginary) **star:2747** Fast and simple HTTP microservice for image resizing.   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [bild](https://github.com/anthonynsimon/bild) **star:2693** Collection of image processing algorithms in pure Go.   [![star > 2000][Awesome]](https://github.com/anthonynsimon/bild)   [![godoc][GoDoc]](https://godoc.org/github.com/anthonynsimon/bild)
* [ln](https://github.com/fogleman/ln) **star:2566** 3D line art rendering in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/ln)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/ln)
* [resize](https://github.com/nfnt/resize) **star:2220** Image resizing for Go with common interpolation methods.   [![star > 2000][Awesome]](https://github.com/nfnt/resize)   [![It hasn't been updated in the last year][Yellow]](https://github.com/nfnt/resize)   [![godoc][GoDoc]](https://godoc.org/github.com/nfnt/resize)
* [gg](https://github.com/fogleman/gg) **star:2052** 2D rendering in pure Go.   [![star > 2000][Awesome]](https://github.com/fogleman/gg)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/gg)
* [pt](https://github.com/fogleman/pt) **star:1808** Path tracing engine written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/pt)
* [svgo](https://github.com/ajstarks/svgo) **star:1399** Go Language Library for SVG generation.   [![godoc][GoDoc]](https://godoc.org/github.com/ajstarks/svgo)
* [smartcrop](https://github.com/muesli/smartcrop) **star:1295** Finds good crops for arbitrary images and crop sizes.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/smartcrop)
* [gift](https://github.com/disintegration/gift) **star:1274** Package of image processing filters.   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/gift)
* [picfit](https://github.com/thoas/picfit) **star:1151** An image resizing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/picfit)
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1134** Go bindings for OpenCV.   [![godoc][GoDoc]](https://godoc.org/github.com/lazywei/go-opencv)
* [imagick](https://github.com/gographics/imagick) **star:1037** Go binding to ImageMagick's MagickWand C API.   [![godoc][GoDoc]](https://godoc.org/github.com/gographics/imagick)
* [geopattern](https://github.com/pravj/geopattern) **star:1033** Create beautiful generative image patterns from a string.   [![godoc][GoDoc]](https://godoc.org/github.com/pravj/geopattern)
* [bimg](https://github.com/h2non/bimg) **star:855** Small package for fast and efficient image processing using libvips.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:578** Go tool for LSB steganography, capable of hiding any file within an image.   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [mort](https://github.com/aldor007/mort) **star:377** Storage and image processing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aldor007/mort)
* [canvas](https://github.com/tdewolff/canvas) **star:350** Vector graphics to PDF, SVG or rasterized image.   [![There was an update last week][Green]](https://github.com/tdewolff/canvas)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/canvas)
* [image2ascii](https://github.com/qeesung/image2ascii) **star:335** Convert image to ASCII.   [![It hasn't been updated in the last year][Yellow]](https://github.com/qeesung/image2ascii)   [![godoc][GoDoc]](https://godoc.org/github.com/qeesung/image2ascii)
* [govatar](https://github.com/o1egl/govatar) **star:329** Library and CMD tool for generating funny avatars.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/govatar)
* [go-nude](https://github.com/koyachi/go-nude) **star:294** Nudity detection with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/koyachi/go-nude)   [![godoc][GoDoc]](https://godoc.org/github.com/koyachi/go-nude)
* [goimagehash](https://github.com/corona10/goimagehash) **star:250** Go Perceptual image hashing package.   [![There was an update last week][Green]](https://github.com/corona10/goimagehash)   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [rez](https://github.com/bamiaux/rez) **star:193** Image resizing in pure Go and SIMD.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bamiaux/rez)   [![godoc][GoDoc]](https://godoc.org/github.com/bamiaux/rez)
* [img](https://github.com/hawx/img) **star:131** Selection of image manipulation tools.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hawx/img)   [![godoc][GoDoc]](https://godoc.org/github.com/hawx/img)
* [darkroom](https://github.com/gojek/darkroom) **star:99** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [go-cairo](https://github.com/ungerik/go-cairo) **star:90** Go binding for the cairo graphics library.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [mergi](https://github.com/noelyahan/mergi) **star:85** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).   [![godoc][GoDoc]](https://godoc.org/github.com/noelyahan/mergi)
* [gltf](https://github.com/qmuntal/gltf) **star:55** Efficient and robust glTF 2.0 reader, writer and validator.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [go-gd](https://github.com/bolknote/go-gd) **star:51** Go binding for GD library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bolknote/go-gd)   [![godoc][GoDoc]](https://godoc.org/github.com/bolknote/go-gd)
* [cameron](https://github.com/aofei/cameron) **star:36** An avatar generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)
* [steganography](https://github.com/auyer/steganography) **star:32** Pure Go Library for LSB steganography.   [![godoc][GoDoc]](https://godoc.org/github.com/auyer/steganography)
* [goimghdr](https://github.com/corona10/goimghdr) **star:31** The imghdr module determines the type of image contained in a file for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimghdr)
* [tga](https://github.com/ftrvxmtrx/tga) **star:25** Package tga is a TARGA image format decoder/encoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ftrvxmtrx/tga)   [![godoc][GoDoc]](https://godoc.org/github.com/ftrvxmtrx/tga)
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Port of webcolors library from Python to Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jyotiska/go-webcolors)   [![godoc][GoDoc]](https://godoc.org/github.com/jyotiska/go-webcolors)
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.   [![godoc][GoDoc]](https://godoc.org/github.com/donatj/mpo)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1262** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt) **star:856** Gatt is a Go package for building Bluetooth Low Energy peripherals.   [![godoc][GoDoc]](https://godoc.org/github.com/paypal/gatt)
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [mainflux](https://github.com/Mainflux/mainflux) **star:711** Industrial IoT Messaging and Device Management Server.   [![There was an update last week][Green]](https://github.com/Mainflux/mainflux)   [![godoc][GoDoc]](https://godoc.org/github.com/Mainflux/mainflux)
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [devices](https://github.com/goiot/devices) **star:229** Suite of libraries for IoT devices, experimental for x/exp/io.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goiot/devices)   [![godoc][GoDoc]](https://godoc.org/github.com/goiot/devices)
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:187** Lightweight stream processing engine for IoT.   [![godoc][GoDoc]](https://godoc.org/github.com/sensorbee/sensorbee)
* [connectordb](https://github.com/connectordb/connectordb) **star:186** Open-Source Platform for Quantified Self & IoT.   [![There was an update last week][Green]](https://github.com/connectordb/connectordb)   [![godoc][GoDoc]](https://godoc.org/github.com/connectordb/connectordb)
* [huego](https://github.com/amimof/huego) **star:125** An extensive Philips Hue client library for Go.   [![There was an update last week][Green]](https://github.com/amimof/huego)   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [eywa](https://github.com/xcodersun/eywa) **star:42** Project Eywa is essentially a connection manager that keeps track of connected devices.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xcodersun/eywa)   [![godoc][GoDoc]](https://godoc.org/github.com/xcodersun/eywa)

## Job Scheduler

*Libraries for scheduling jobs.*

* [gron](https://github.com/roylee0704/gron) **star:669** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   [![godoc][GoDoc]](https://godoc.org/github.com/roylee0704/gron)
* [JobRunner](https://github.com/bamzi/jobrunner) **star:623** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   [![There was an update last week][Green]](https://github.com/bamzi/jobrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/bamzi/jobrunner)
* [jobs](https://github.com/albrow/jobs) **star:459** Persistent and flexible background jobs library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/jobs)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/jobs)
* [scheduler](https://github.com/carlescere/scheduler) **star:309** Cronjobs scheduling made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/scheduler)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/scheduler)
* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [go-cron](https://github.com/rk/go-cron) **star:177** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rk/go-cron)   [![godoc][GoDoc]](https://godoc.org/github.com/rk/go-cron)
* [clockwork](https://github.com/whiteShtef/clockwork) **star:94** Simple and intuitive job scheduling library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/whiteShtef/clockwork)
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:49** Job scheduler that supports webhooks, crons and classic scheduling.   [![There was an update last week][Green]](https://github.com/kilgaloon/leprechaun)   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) **star:5510** Get a JSON value with one line of code.   [![star > 2000][Awesome]](https://github.com/tidwall/gjson)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/gjson)
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2111** Automatically generate Go (golang) struct definitions from example JSON.   [![star > 2000][Awesome]](https://github.com/ChimeraCoder/gojson)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/gojson)
* [kazaam](https://github.com/Qntfy/kazaam) **star:143** API for arbitrary transformation of JSON documents.   [![godoc][GoDoc]](https://godoc.org/github.com/Qntfy/kazaam)
* [gojq](https://github.com/elgs/gojq) **star:142** JSON query in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/elgs/gojq)   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/gojq)
* [jsongo](https://github.com/ricardolonga/jsongo) **star:94** Fluent API to make it easier to create Json objects.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ricardolonga/jsongo)   [![godoc][GoDoc]](https://godoc.org/github.com/ricardolonga/jsongo)
* [gjo](https://github.com/skanehira/gjo) **star:67** Small utility to create JSON objects.   [![There was an update last week][Green]](https://github.com/skanehira/gjo)   [![godoc][GoDoc]](https://godoc.org/github.com/skanehira/gjo)
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:59** JSON diff utility written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [jsonf](https://github.com/miolini/jsonf) **star:55** Console tool for highlighted formatting and struct query fetching JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/miolini/jsonf)   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/jsonf)
* [jettison](https://github.com/wI2L/jettison) **star:54** High performance, reflection-less JSON encoder for Go.   [![There was an update last week][Green]](https://github.com/wI2L/jettison)   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jettison)
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [mp](https://github.com/sanbornm/mp) **star:35** Simple cli email parser. It currently takes stdin and outputs JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sanbornm/mp)   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/mp)
* [go-respond](https://github.com/nicklaw5/go-respond) **star:27** Go package for handling common HTTP JSON responses.   [![godoc][GoDoc]](https://godoc.org/github.com/nicklaw5/go-respond)
* [json2go](https://github.com/m-zajac/json2go) **star:26** Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.   [![godoc][GoDoc]](https://godoc.org/github.com/m-zajac/json2go)
* [ajson](https://github.com/spyzhov/ajson) **star:19** Abstract JSON for golang with JSONPath support.   [![godoc][GoDoc]](https://godoc.org/github.com/spyzhov/ajson)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/RichardKnop/jsonhal)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:9** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.   [![godoc][GoDoc]](https://godoc.org/github.com/ddymko/go-jsonerror)
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:6** Go bindings based on the JSON API errors reference.   [![It hasn't been updated in the last year][Yellow]](https://github.com/AmuzaTkts/jsonapi-errors)   [![godoc][GoDoc]](https://godoc.org/github.com/AmuzaTkts/jsonapi-errors)

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) **star:13225** Structured logger for Go.   [![star > 2000][Awesome]](https://github.com/Sirupsen/logrus)   [![There was an update last week][Green]](https://github.com/Sirupsen/logrus)   [![godoc][GoDoc]](https://godoc.org/github.com/Sirupsen/logrus)
* [zap](https://github.com/uber-go/zap) **star:8488** Fast, structured, leveled logging in Go.   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [spew](https://github.com/davecgh/go-spew) **star:3522** Implements a deep pretty printer for Go data structures to aid in debugging.   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   [![There was an update last week][Green]](https://github.com/davecgh/go-spew)   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [zerolog](https://github.com/rs/zerolog) **star:2679** Zero-allocation JSON logger.   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![There was an update last week][Green]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [glog](https://github.com/golang/glog) **star:2428** Leveled execution logs for Go.   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [tail](https://github.com/hpcloud/tail) **star:1652** Go package striving to emulate the features of the BSD tail program.   [![There was an update last week][Green]](https://github.com/hpcloud/tail)   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1650** Simple rolling logger, implements io.WriteCloser.   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [seelog](https://github.com/cihub/seelog) **star:1398** Logging functionality with flexible dispatching, filtering, and formatting.   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log15](https://github.com/inconshreveable/log15) **star:940** Simple, powerful logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/log15)
* [log](https://github.com/apex/log) **star:766** Structured logging package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [onelog](https://github.com/francoispqt/onelog) **star:353** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.   [![godoc][GoDoc]](https://godoc.org/github.com/francoispqt/onelog)
* [logxi](https://github.com/mgutz/logxi) **star:338** 12-factor app logger that is fast and makes you happy.   [![godoc][GoDoc]](https://godoc.org/github.com/mgutz/logxi)
* [log](https://github.com/go-playground/log) **star:270** Simple, configurable and scalable Structured Logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [logutils](https://github.com/hashicorp/logutils) **star:261** Utilities for slightly better logging in Go (Golang) extending the standard logger.   [![There was an update last week][Green]](https://github.com/hashicorp/logutils)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [go-logger](https://github.com/apsdehal/go-logger) **star:245** Simple logger of Go Programs, with level handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/apsdehal/go-logger)
* [logger](https://github.com/azer/logger) **star:136** Minimalistic logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/azer/logger)
* [xlog](https://github.com/rs/xlog) **star:130** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:120** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   [![There was an update last week][Green]](https://github.com/arthurkiller/rollingWriter)   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:109** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-ozzo/ozzo-log)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-log)   [![Contains Chinese documents][CN]](https://github.com/go-ozzo/ozzo-log)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:83** Full-featured logging saas written in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/logvoyage)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [glg](https://github.com/kpango/glg) **star:56** glg is simple and fast leveled logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kpango/glg)
* [stdlog](https://github.com/alexcesaro/log) **star:42** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexcesaro/log)   [![godoc][GoDoc]](https://godoc.org/github.com/alexcesaro/log)
* [gologger](https://github.com/sadlil/gologger) **star:38** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/gologger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/gologger)   [![Archived][Archived]](https://github.com/sadlil/gologger)
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/go-log)
* [logex](https://github.com/chzyer/logex) **star:33** Golang log lib, supports tracking and level, wrap by standard log lib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chzyer/logex)   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/logex)
* [logrusly](https://github.com/sebest/logrusly) **star:25** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/sebest/logrusly)   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/logrusly)
* [go-log](https://github.com/siddontang/go-log) **star:25** Log lib supports level and multi handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-log)
* [log](https://github.com/teris-io/log) **star:23** Structured log interface for Go cleanly separates logging facade from its implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/log)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/log)
* [distillog](https://github.com/amoghe/distillog) **star:21** distilled levelled logging (think of it as stdlib + log levels).   [![It hasn't been updated in the last year][Yellow]](https://github.com/amoghe/distillog)   [![godoc][GoDoc]](https://godoc.org/github.com/amoghe/distillog)
* [journald](https://github.com/ssgreg/journald) **star:21** Go implementation of systemd Journal's native API for logging.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/journald)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:20** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jbrodriguez/mlog)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrodriguez/mlog)
* [gomol](https://github.com/aphistic/gomol) **star:15** Multiple-output, structured logging for Go with extensible logging outputs.   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ewwwwwqm/logdump)   [![godoc][GoDoc]](https://godoc.org/github.com/ewwwwwqm/logdump)
* [go-log](https://github.com/subchen/go-log) **star:9** Simple and configurable Logging in Go, with level, formatters and writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/subchen/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-log)
* [glo](https://github.com/lajosbencz/glo) **star:8** PHP Monolog inspired logging facility with identical severity levels.   [![godoc][GoDoc]](https://godoc.org/github.com/lajosbencz/glo)
* [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:7** `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/logrusiowriter)
* [log](https://github.com/aerogo/log) **star:6** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/log)
* [xlog](https://github.com/xfxdev/xlog) **star:6** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [logo](https://github.com/mbndr/logo) **star:5** Golang logger to different configurable writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mbndr/logo)   [![godoc][GoDoc]](https://godoc.org/github.com/mbndr/logo)
* [go-log](https://github.com/pieterclaerhout/go-log) **star:1** A logging library with strack traces, object dumping and optional timestamps.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-log)

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:6858** General Machine Learning library for Go.   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   [![Contains Chinese documents][CN]](https://github.com/sjwhitworth/golearn)
* [gorgonia](https://github.com/gorgonia/gorgonia) **star:2995** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![star > 2000][Awesome]](https://github.com/gorgonia/gorgonia)   [![There was an update last week][Green]](https://github.com/gorgonia/gorgonia)   [![godoc][GoDoc]](https://godoc.org/github.com/gorgonia/gorgonia)
* [tfgo](https://github.com/galeone/tfgo) **star:1261** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/tfgo)
* [goml](https://github.com/cdipaolo/goml) **star:1041** On-line Machine Learning in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gosseract](https://github.com/otiai10/gosseract) **star:979** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [gorse](https://github.com/zhenghaoz/gorse) **star:849** A High Performance Recommender System Package based on Collaborative Filtering for Go.   [![There was an update last week][Green]](https://github.com/zhenghaoz/gorse)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   [![Contains Chinese documents][CN]](https://github.com/zhenghaoz/gorse)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:658** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ryanbressler/CloudForest)   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:650** An evolutionary optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [bayesian](https://github.com/jbrukh/bayesian) **star:642** Naive Bayesian Classification for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [gobrain](https://github.com/goml/gobrain) **star:412** Neural Networks written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [regommend](https://github.com/muesli/regommend) **star:260** Recommendation & collaborative filtering engine.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [ocrserver](https://github.com/otiai10/ocrserver) **star:252** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/ocrserver)
* [go-deep](https://github.com/patrikeh/go-deep) **star:241** A feature-rich neural network library in Go.   [![There was an update last week][Green]](https://github.com/patrikeh/go-deep)   [![godoc][GoDoc]](https://godoc.org/github.com/patrikeh/go-deep)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:225** Go Interface to Open Neural Network Exchange (ONNX).   [![There was an update last week][Green]](https://github.com/owulveryck/onnx-go)   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-galib](https://github.com/thoj/go-galib) **star:173** Genetic Algorithms library written in Go / golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thoj/go-galib)   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:149** Recommendation Algorithms library written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timkaye11/goRecommend)   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [shield](https://github.com/eaigner/shield) **star:125** Bayesian text classifier with flexible tokenizers and storage backends for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eaigner/shield)   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [go-fann](https://github.com/white-pony/go-fann) **star:102** Go bindings for Fast Artificial Neural Networks(FANN) library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/white-pony/go-fann)   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [Goptuna](https://github.com/c-bata/goptuna) **star:97** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![There was an update last week][Green]](https://github.com/c-bata/goptuna)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [goga](https://github.com/tomcraven/goga) **star:79** Genetic algorithm library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomcraven/goga)   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:63** libsvm golang version derived work based on LIBSVM 3.14.   [![It hasn't been updated in the last year][Yellow]](https://github.com/datastream/libsvm)   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [neural-go](https://github.com/schuyler/neural-go) **star:61** Multilayer perceptron network implemented in Go, with training via backpropagation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schuyler/neural-go)   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [go-pr](https://github.com/daviddengcn/go-pr) **star:57** Pattern recognition package in Go lang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-pr)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [neat](https://github.com/jinyeom/neat) **star:56** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/jinyeom/neat)   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   [![Archived][Archived]](https://github.com/jinyeom/neat)
* [goscore](https://github.com/asafschers/goscore) **star:41** Go Scoring API for PMML.   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/golinear)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [fonet](https://github.com/Fontinalis/fonet) **star:35** A Deep Neural Network library written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Fontinalis/fonet)
* [godist](https://github.com/e-dard/godist) **star:25** Various probability distributions, and associated methods.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-dard/godist)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/godist)
* [Varis](https://github.com/Xamber/Varis) **star:25** Golang Neural Network.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Xamber/Varis)   [![godoc][GoDoc]](https://godoc.org/github.com/Xamber/Varis)
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:21** Go implementation of the k-modes and k-prototypes clustering algorithms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
* [probab](https://github.com/ThePaw/probab) **star:11** Probability distribution functions. Bayesian inference. Written in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/probab)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/probab)
* [evoli](https://github.com/khezen/evoli) **star:8** Genetic Algorithm and Particle Swarm Optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)
* [GoMind](https://github.com/surenderthakran/gomind) **star:6** A simplistic Neural Network Library in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surenderthakran/gomind)   [![godoc][GoDoc]](https://godoc.org/github.com/surenderthakran/gomind)

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) **star:5134** Go library for Apache Kafka.   [![star > 2000][Awesome]](https://github.com/Shopify/sarama)   [![There was an update last week][Green]](https://github.com/Shopify/sarama)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/sarama)
* [gorush](https://github.com/appleboy/gorush) **star:3979** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![There was an update last week][Green]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:3940** Real-time messaging (Websockets or SockJS) server in Go.   [![star > 2000][Awesome]](https://github.com/centrifugal/centrifugo)   [![There was an update last week][Green]](https://github.com/centrifugal/centrifugo)   [![godoc][GoDoc]](https://godoc.org/github.com/centrifugal/centrifugo)
* [machinery](https://github.com/RichardKnop/machinery) **star:3700** Asynchronous task queue/job queue based on distributed message passing.   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:3120** socket.io library for golang, a realtime application framework.   [![star > 2000][Awesome]](https://github.com/googollee/go-socket.io)   [![There was an update last week][Green]](https://github.com/googollee/go-socket.io)   [![godoc][GoDoc]](https://godoc.org/github.com/googollee/go-socket.io)
* [NATS Go Client](https://github.com/nats-io/nats) **star:2583** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.   [![star > 2000][Awesome]](https://github.com/nats-io/nats)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/nats)
* [APNs2](https://github.com/sideshow/apns2) **star:2190** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.   [![star > 2000][Awesome]](https://github.com/sideshow/apns2)   [![There was an update last week][Green]](https://github.com/sideshow/apns2)   [![godoc][GoDoc]](https://godoc.org/github.com/sideshow/apns2)
* [Benthos](https://github.com/Jeffail/benthos) **star:2136** A message streaming bridge between a range of protocols.   [![star > 2000][Awesome]](https://github.com/Jeffail/benthos)   [![There was an update last week][Green]](https://github.com/Jeffail/benthos)   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/benthos)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1863** gopush-cluster is a go push server cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Terry-Mao/gopush-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   [![Contains Chinese documents][CN]](https://github.com/Terry-Mao/gopush-cluster)
* [Mercure](https://github.com/dunglas/mercure) **star:1772** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   [![There was an update last week][Green]](https://github.com/dunglas/mercure)   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [melody](https://github.com/olahol/melody) **star:1674** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [mangos](https://github.com/go-mangos/mangos) **star:1548** Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mangos/mangos)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1547** the official Go package for NSQ.   [![There was an update last week][Green]](https://github.com/nsqio/go-nsq)   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1122** Redis backed unified push service for server-side notifications to mobile devices.   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [zmq4](https://github.com/pebbe/zmq4) **star:812** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/zmq4)
* [Gollum](https://github.com/trivago/gollum) **star:794** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [Beaver](https://github.com/Clivern/Beaver) **star:766** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![There was an update last week][Green]](https://github.com/Clivern/Beaver)   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [EventBus](https://github.com/asaskevich/EventBus) **star:610** The lightweight event bus with async compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [golongpoll](https://github.com/jcuga/golongpoll) **star:436** HTTP longpoll server library that makes web pub-sub simple.   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [dbus](https://github.com/godbus/dbus) **star:389** Native Go bindings for D-Bus.   [![There was an update last week][Green]](https://github.com/godbus/dbus)   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [emitter](https://github.com/olebedev/emitter) **star:327** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [Glue](https://github.com/desertbit/glue) **star:323** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:299** Simple pubsub package for go.   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [guble](https://github.com/smancke/guble) **star:141** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   [![It hasn't been updated in the last year][Yellow]](https://github.com/smancke/guble)   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [Bus](https://github.com/mustafaturan/bus) **star:131** Minimalist message bus implementation for internal communication.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [oplog](https://github.com/dailymotion/oplog) **star:100** Generic oplog/replication system for REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dailymotion/oplog)   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [rabtap](https://github.com/jandelgado/rabtap) **star:93** RabbitMQ swiss army knife cli app.   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [messagebus](https://github.com/vardius/message-bus) **star:76** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/message-bus)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:67** A tiny wrapper over amqp exchanges and queues.   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [drone-line](https://github.com/appleboy/drone-line) **star:63** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-line)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:57** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sybrexsys/RapidMQ)   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:56** A tiny wrapper around NSQ topic and channel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:47** Native implementation of the freedesktop notification spec.   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [Commander](https://github.com/jeroenrinzema/commander) **star:34** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   [![There was an update last week][Green]](https://github.com/jeroenrinzema/commander)   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [event](https://github.com/agoalofalife/event) **star:31** Implementation of the pattern observer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agoalofalife/event)   [![godoc][GoDoc]](https://godoc.org/github.com/agoalofalife/event)
* [hub](https://github.com/leandro-lugaresi/hub) **star:27** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   [![It hasn't been updated in the last year][Yellow]](https://github.com/leandro-lugaresi/hub)   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:12** Client library to Viessmann Vitotrol web service.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** Gaurun Client written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gaurun-client)
* [jazz](https://github.com/socifi/jazz) **star:7** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:7** redisqueue provides a producer and consumer of a queue that uses Redis streams.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/redisqueue)
* [rmqconn](https://github.com/sbabiv/rmqconn)  RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:1951** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![There was an update last week][Green]](https://github.com/unidoc/unioffice)   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:5077** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   [![There was an update last week][Green]](https://github.com/360EntSecGroup-Skylar/excelize)   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:3724** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:93** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:53** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:14** Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fterrag/goxlsxwriter)   [![godoc][GoDoc]](https://godoc.org/github.com/fterrag/goxlsxwriter)

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:1072** A reflection based dependency injection toolkit for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:1051** A dependency injection based application framework for Go (built on top of dig).   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [container](https://github.com/golobby/container) **star:53** A powerful IoC Container with fluent and easy-to-use interface.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/container)
* [inject](https://github.com/defval/inject) **star:45** A reflection based dependency injection container with simple interface.   [![There was an update last week][Green]](https://github.com/defval/inject)   [![godoc][GoDoc]](https://godoc.org/github.com/defval/inject)
* [alice](https://github.com/magic003/alice) **star:35** Additive dependency injection container for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/magic003/alice)   [![godoc][GoDoc]](https://godoc.org/github.com/magic003/alice)
* [wire](https://github.com/Fs02/wire) **star:20** Strict Runtime Dependency Injection for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/wire)
* [linker](https://github.com/logrange/linker) **star:17** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [gocontainer](https://github.com/vardius/gocontainer) **star:11** Simple Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:11833** Set of common historical and emerging project layout patterns in the Go ecosystem.   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)   [![There was an update last week][Green]](https://github.com/golang-standards/project-layout)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:421** Go application boilerplate and example applying modern practices.   [![There was an update last week][Green]](https://github.com/sagikazarmark/modern-go-application)   [![godoc][GoDoc]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
* [scaffold](https://github.com/catchplay/scaffold) **star:44** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [go-sample](https://github.com/zitryss/go-sample) **star:39** A sample layout for Go application projects with the real code.   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:667** Collection of useful string functions ported from other languages.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:80** String utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4341** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![There was an update last week][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:2631** Library and command for making and extracting .zip and .tar.gz archives.   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![There was an update last week][Green]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gosms](https://github.com/haxpax/gosms) **star:1253** Your own local SMS gateway in Go that can be used to send SMS.   [![It hasn't been updated in the last year][Yellow]](https://github.com/haxpax/gosms)   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:935** Resiliency patterns for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:740** Generic object pool for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![Contains Chinese documents][CN]](https://github.com/jolestar/go-commons-pool)
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:700** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![There was an update last week][Green]](https://github.com/mojocn/base64Captcha)   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![Contains Chinese documents][CN]](https://github.com/mojocn/base64Captcha)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:653** Random data generator written in go.   [![There was an update last week][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [shortid](https://github.com/teris-io/shortid) **star:496** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/shortid)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:462** Library for interacting with LLVM IR in pure Go.   [![There was an update last week][Green]](https://github.com/llir/llvm)   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:372** Easy to use, extensible health check library.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:344** Package conv provides fast and intuitive conversions across Go types.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [banner](https://github.com/dimiro1/banner) **star:242** Add beautiful banners into your Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:222** Package that exposes country and subdivision data.   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [antch](https://github.com/antchfx/antch) **star:161** A fast, powerful and extensible web crawling & scraping framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/antch)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   [![Contains Chinese documents][CN]](https://github.com/antchfx/antch)
* [stateless](https://github.com/qmuntal/stateless) **star:139** A fluent library for creating state machines.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:138** Beautify data display for Humans.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   [![Contains Chinese documents][CN]](https://github.com/go-ffmt/ffmt)
* [lk](https://github.com/hyperboloide/lk) **star:137** A simple licensing library for golang.   [![There was an update last week][Green]](https://github.com/hyperboloide/lk)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [battery](https://github.com/distatus/battery) **star:137** Cross-platform, normalized battery information library.   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [stats](https://github.com/go-playground/stats) **star:128** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [bitio](https://github.com/icza/bitio) **star:105** Highly optimized bit-level Reader and Writer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:93** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:87** Emojis for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [gommit](https://github.com/antham/gommit) **star:79** Analyze git commit messages to ensure they follow defined patterns.   [![There was an update last week][Green]](https://github.com/antham/gommit)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [ghorg](https://github.com/gabrie30/ghorg) **star:77** Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:76** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [indigo](https://github.com/osamingo/indigo) **star:57** Distributed unique ID generator of using Sonyflake and encoded by Base58.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [morse](https://github.com/alwindoss/morse) **star:53** Library to convert to and from morse code.   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [captcha](https://github.com/steambap/captcha) **star:47** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [xkg](https://github.com/go-xkg/xkg) **star:45** X Keyboard Grabber.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [persian](https://github.com/mavihq/persian) **star:35** Some utilities for Persian language in go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mavihq/persian)   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:35** HTTP service to generate PDF from Json requests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:31** GoLang Library for [Browser Capabilities Project](http://browscap.org/).   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [datacounter](https://github.com/miolini/datacounter) **star:30** Go counters for readers/writer/http.ResponseWriter.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [gotoprom](https://github.com/cabify/gotoprom) **star:25** Type-safe metrics builder wrapper library for the official Prometheus client.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [autoflags](https://github.com/artyom/autoflags) **star:25** Go package to automatically define command line flags from struct fields.   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [xdg](https://github.com/rkoesters/xdg) **star:20** FreeDesktop.org (xdg) Specs implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [gosh](https://github.com/osamingo/gosh) **star:19** Provide Go Statistics Handler, Struct, Measure Method.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:18** A modern, powerful, and robust URL shortener microservice with mysql support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [sandid](https://github.com/aofei/sandid) **star:13** Every grain of sand on earth has its own ID.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [anagent](https://github.com/mudler/anagent) **star:10** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mudler/anagent)   [![godoc][GoDoc]](https://godoc.org/github.com/mudler/anagent)
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:9** Calculate average score and rating based on Wilson Score Equation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/avgRating)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/avgRating)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [shellwords](https://github.com/Wing924/shellwords) **star:7** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [metrics](https://github.com/pascaldekloe/metrics) **star:4** Library for metrics instrumentation and Prometheus exposition.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2373** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only.   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![There was an update last week][Green]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1146** Go efficient text segmentation; support english, chinese, japanese and other.   [![There was an update last week][Green]](https://github.com/go-ego/gse)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![Contains Chinese documents][CN]](https://github.com/go-ego/gse)
* [when](https://github.com/olebedev/when) **star:1042** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:924** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![Contains Chinese documents][CN]](https://github.com/yanyiwu/gojieba)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:614** CN Hanzi to Hanyu Pinyin converter.   [![There was an update last week][Green]](https://github.com/mozillazg/go-pinyin)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:458** JP morphological analyzer written in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:379** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:357** Extract values from strings and fill your structs with nlp.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:264** Sentence tokenizer:  converts text into a list of sentences.   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:231** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![There was an update last week][Green]](https://github.com/james-bowman/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-nlp](https://github.com/nuance/go-nlp) **star:81** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [getlang](https://github.com/rylans/getlang) **star:78** Fast natural language detection package.   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [gounidecode](https://github.com/fiam/gounidecode) **star:68** Unicode transliterator (also known as unidecode) for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:64** ASCII transliterations of Unicode text.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [textcat](https://github.com/pebbe/textcat) **star:62** Go package for n-gram based text categorization, with support for utf-8 and raw text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:58** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:52** Implementation of the porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:50** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [stemmer](https://github.com/dchest/stemmer) **star:48** Stemmer packages for Go programming language. Includes English and German stemmers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [segment](https://github.com/blevesearch/segment) **star:46** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![It hasn't been updated in the last year][Yellow]](https://github.com/blevesearch/segment)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [porter2](https://github.com/zhenjl/porter2) **star:35** Really fast Porter 2 stemmer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/porter2)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [go2vec](https://github.com/danieldk/go2vec) **star:32** Reader and utility functions for word2vec embeddings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang implementation of the Paice/Husk Stemming Algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [petrovich](https://github.com/striker2000/petrovich) **star:24** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [snowball](https://github.com/goodsign/snowball) **star:24** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/snowball)
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo bindings to Yandex.Mystem - russian morphology analyzer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:15** Go bindings for the snowball libstemmer library including porter 2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [shamoji](https://github.com/osamingo/shamoji) **star:11** The shamoji is word filtering package written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:6** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)

## Networking

*Libraries for working with various layers of the network.*

* [kcptun](https://github.com/xtaci/kcptun) **star:11253** Extremely simple & fast udp tunnel based on KCP protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcptun)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcptun)
* [fasthttp](https://github.com/valyala/fasthttp) **star:10932** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![There was an update last week][Green]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [dns](https://github.com/miekg/dns) **star:4146** Go library for working with DNS.   [![star > 2000][Awesome]](https://github.com/miekg/dns)   [![There was an update last week][Green]](https://github.com/miekg/dns)   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:3562** An implementation of the QUIC protocol in pure Go.   [![star > 2000][Awesome]](https://github.com/lucas-clemente/quic-go)   [![There was an update last week][Green]](https://github.com/lucas-clemente/quic-go)   [![godoc][GoDoc]](https://godoc.org/github.com/lucas-clemente/quic-go)
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3494** HTTPLabs let you inspect HTTP requests and forge responses.   [![star > 2000][Awesome]](https://github.com/gchaincl/httplab)   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/httplab)
* [gopacket](https://github.com/google/gopacket) **star:3134** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [webrtc](https://github.com/pions/webrtc) **star:2759** A pure Go implementation of the WebRTC API.   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   [![There was an update last week][Green]](https://github.com/pions/webrtc)   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2392** KCP - Fast and Reliable ARQ Protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [gobgp](https://github.com/osrg/gobgp) **star:1774** BGP implemented in the Go Programming Language.   [![There was an update last week][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [gnet](https://github.com/panjf2000/gnet) **star:1426** `gnet` is a high-performance, lightweight, nonblocking, event-loop networking library written in pure Go.   [![There was an update last week][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/gnet)
* [ssh](https://github.com/gliderlabs/ssh) **star:1241** Higher-level API for building SSH servers (wraps crypto/ssh).   [![godoc][GoDoc]](https://godoc.org/github.com/gliderlabs/ssh)
* [fortio](https://github.com/fortio/fortio) **star:1036** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:923** Simple TUN/TAP library.   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [go-getter](https://github.com/hashicorp/go-getter) **star:810** Go library for downloading files or directories from various sources using a URL.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [sftp](https://github.com/pkg/sftp) **star:807** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/sftp)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:738** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   [![There was an update last week][Green]](https://github.com/intel-go/nff-go)   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [gev](https://github.com/Allenxuxu/gev) **star:716** gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.   [![There was an update last week][Green]](https://github.com/Allenxuxu/gev)   [![godoc][GoDoc]](https://godoc.org/github.com/Allenxuxu/gev)
* [grab](https://github.com/cavaliercoder/grab) **star:605** Go package for managing file downloads.   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [mdns](https://github.com/hashicorp/mdns) **star:590** Simple mDNS (Multicast DNS) client/server library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/mdns)
* [ftp](https://github.com/jlaffaye/ftp) **star:581** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [lhttp](https://github.com/fanux/lhttp) **star:535** Powerful websocket framework, build your IM server more easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fanux/lhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   [![Contains Chinese documents][CN]](https://github.com/fanux/lhttp)
* [gosnmp](https://github.com/soniah/gosnmp) **star:472** Native Go library for performing SNMP actions.   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [gotcp](https://github.com/gansidui/gotcp) **star:440** Go package for quickly writing tcp applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/gotcp)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/gotcp)
* [cidranger](https://github.com/yl2chen/cidranger) **star:421** Fast IP to CIDR lookup for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:388** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [gopcap](https://github.com/akrennmair/gopcap) **star:369** Go wrapper for libpcap.   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [go-stun](https://github.com/ccding/go-stun) **star:339** Go implementation of the STUN client (RFC 3489 and RFC 5389).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ccding/go-stun)   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [raw](https://github.com/mdlayher/raw) **star:333** Package raw enables reading and writing data at the device driver level for a network interface.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [stun](https://github.com/go-rtc/stun) **star:319** Go implementation of RFC 5389 STUN protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:303** Go library for building tcp servers faster.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:234** Streaming protocolbuffer data over TCP made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stabbycutyou/buffstreams)   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [winrm](https://github.com/masterzen/winrm) **star:232** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [arp](https://github.com/mdlayher/arp) **star:205** Package arp implements the ARP protocol, as described in RFC 826.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:191** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [utp](https://github.com/anacrolix/utp) **star:152** Go uTP micro transport protocol implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/anacrolix/utp)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/utp)
* [canopus](https://github.com/zubairhamed/canopus) **star:137** CoAP Client/Server implementation (RFC 7252).   [![It hasn't been updated in the last year][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [jazigo](https://github.com/udhos/jazigo) **star:137** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [sslb](https://github.com/eduardonunesp/sslb) **star:122** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:116** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![Contains Chinese documents][CN]](https://github.com/DrmagicE/gmqtt)
* [gNxI](https://github.com/google/gnxi) **star:112** A collection of tools for Network Management that use the gNMI and gNOI protocols.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [xtcp](https://github.com/xfxdev/xtcp) **star:91** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:64** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/dhcp6)
* [ether](https://github.com/songgao/ether) **star:61** Cross-platform Go package for sending and receiving ethernet frames.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [linkio](https://github.com/ian-kent/linkio) **star:44** Network link speed simulation for Reader/Writer interfaces.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/linkio)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/linkio)
* [portproxy](https://github.com/aybabtme/portproxy) **star:42** Simple TCP proxy which adds CORS support to API's which don't support it.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aybabtme/portproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/aybabtme/portproxy)
* [packet](https://github.com/aerogo/packet) **star:40** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [iplib](https://github.com/c-robinson/iplib) **star:26** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [graval](https://github.com/koofr/graval) **star:24** Experimental FTP server framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/koofr/graval)   [![godoc][GoDoc]](https://godoc.org/github.com/koofr/graval)
* [publicip](https://github.com/polera/publicip) **star:18** Package publicip returns your public facing IPv4 address (internet egress).   [![It hasn't been updated in the last year][Yellow]](https://github.com/polera/publicip)   [![godoc][GoDoc]](https://godoc.org/github.com/polera/publicip)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:16** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/golibwireshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/golibwireshark)
* [llb](https://github.com/kirillDanshin/llb) **star:9** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/llb)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/llb)
* [go-powerdns](https://github.com/joeig/go-powerdns) **star:9** PowerDNS API bindings for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/joeig/go-powerdns)
* [goshark](https://github.com/sunwxg/goshark) **star:9** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/goshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/goshark)
* [tspool](https://github.com/two/tspool) **star:6** A TCP Library use worker pool to improve performance and protect your server.   [![It hasn't been updated in the last year][Yellow]](https://github.com/two/tspool)   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)

### HTTP Clients

*Libraries for making HTTP requests.*

* [resty](https://github.com/go-resty/resty) **star:2342** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [grequests](https://github.com/levigross/grequests) **star:1494** A Go "clone" of the great and famous Requests library.   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [heimdall](https://github.com/gojektech/heimdall)  An enchanced http client with retry and hystrix capabilities.
* [sling](https://github.com/dghubble/sling) **star:1072** Sling is a Go HTTP client library for creating and sending API requests.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/sling)
* [gentleman](https://github.com/h2non/gentleman) **star:712** Full-featured plugin-driven HTTP client library.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gentleman)
* [pester](https://github.com/sethgrid/pester) **star:343** Go HTTP client calls with retries, backoff, and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/sethgrid/pester)
* [rq](https://github.com/ddo/rq) **star:33** A nicer interface for golang stdlib HTTP client.   [![godoc][GoDoc]](https://godoc.org/github.com/ddo/rq)
* [sreq](https://github.com/winterssy/sreq) **star:27** A simple, user-friendly and concurrent safe HTTP request library for Go.   [![There was an update last week][Green]](https://github.com/winterssy/sreq)   [![godoc][GoDoc]](https://godoc.org/github.com/winterssy/sreq)   [![Contains Chinese documents][CN]](https://github.com/winterssy/sreq)

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) **star:809** Go bindings for GLFW 3.
* [gl](https://github.com/go-gl/gl) **star:669** Go bindings for OpenGL (generated via glow).   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/gl)
* [mathgl](https://github.com/go-gl/mathgl) **star:310** Pure Go math package specialized for 3D math, with inspiration from GLM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/mathgl)
* [goxjs/gl](https://github.com/goxjs/gl) **star:130** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   [![It hasn't been updated in the last year][Yellow]](https://github.com/goxjs/gl)   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/gl)
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:59** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) **star:16350** The fantastic ORM library for Golang, aims to be developer friendly.   [![star > 2000][Awesome]](https://github.com/jinzhu/gorm)   [![There was an update last week][Green]](https://github.com/jinzhu/gorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/gorm)
* [Xorm](https://github.com/go-xorm/xorm) **star:5615** Simple and powerful ORM for Go.   [![star > 2000][Awesome]](https://github.com/go-xorm/xorm)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xorm/xorm)   [![Contains Chinese documents][CN]](https://github.com/go-xorm/xorm)
* [go-pg](https://github.com/go-pg/pg) **star:3333** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   [![There was an update last week][Green]](https://github.com/go-pg/pg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [gorp](https://github.com/go-gorp/gorp) **star:3246** Go Relational Persistence, ORM-ish library for Go.   [![star > 2000][Awesome]](https://github.com/go-gorp/gorp)   [![There was an update last week][Green]](https://github.com/go-gorp/gorp)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorp/gorp)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2532** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   [![star > 2000][Awesome]](https://github.com/volatiletech/sqlboiler)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/sqlboiler)
* [upper.io/db](https://github.com/upper/db) **star:2004** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![star > 2000][Awesome]](https://github.com/upper/db)   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [reform](https://github.com/go-reform/reform) **star:828** Better ORM for Go, based on non-empty interfaces and code generation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-reform/reform)
* [pop/soda](https://github.com/gobuffalo/pop) **star:767** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   [![There was an update last week][Green]](https://github.com/gobuffalo/pop)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/pop)
* [QBS](https://github.com/coocood/qbs) **star:539** Stands for Query By Struct. A Go ORM.   [![It hasn't been updated in the last year][Yellow]](https://github.com/coocood/qbs)   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   [![Contains Chinese documents][CN]](https://github.com/coocood/qbs)
* [go-queryset](https://github.com/jirfag/go-queryset) **star:476** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.   [![godoc][GoDoc]](https://godoc.org/github.com/jirfag/go-queryset)
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:317** A flexible and powerful SQL string builder library plus a zero-config ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/go-sqlbuilder)
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [Zoom](https://github.com/albrow/zoom) **star:245** Blazing-fast datastore and querying engine built on Redis.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/zoom)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [grimoire](https://github.com/Fs02/grimoire) **star:121** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   [![There was an update last week][Green]](https://github.com/Fs02/grimoire)   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/grimoire)
* [go-store](https://github.com/gosuri/go-store) **star:98** Simple and fast Redis backed key-value store library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gosuri/go-store)   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/go-store)
* [Marlow](https://github.com/dadleyy/marlow) **star:68** Generated ORM from project structs for compile time safety assurances.   [![godoc][GoDoc]](https://godoc.org/github.com/dadleyy/marlow)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:6** A simple ORM for Google/Firebase Cloud Firestore.   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)
* [lore](https://github.com/abrahambotros/lore) **star:5** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abrahambotros/lore)   [![godoc][GoDoc]](https://godoc.org/github.com/abrahambotros/lore)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:12961** Go dependency tool.   [![star > 2000][Awesome]](https://github.com/golang/dep)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/dep)
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) **star:7884** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.   [![star > 2000][Awesome]](https://github.com/Masterminds/glide)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/glide)
* [godep](https://github.com/tools/godep) **star:5643** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   [![star > 2000][Awesome]](https://github.com/tools/godep)   [![It hasn't been updated in the last year][Yellow]](https://github.com/tools/godep)   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   [![Archived][Archived]](https://github.com/tools/godep)
* [govendor](https://github.com/kardianos/govendor) **star:4930** Go Package Manager. Go vendor tool that works with the standard vendor file.   [![star > 2000][Awesome]](https://github.com/kardianos/govendor)   [![godoc][GoDoc]](https://godoc.org/github.com/kardianos/govendor)
* [gopm](https://github.com/gpmgo/gopm) **star:2430** Go Package Manager.   [![star > 2000][Awesome]](https://github.com/gpmgo/gopm)   [![godoc][GoDoc]](https://godoc.org/github.com/gpmgo/gopm)   [![Archived][Archived]](https://github.com/gpmgo/gopm)
* [gom](https://github.com/mattn/gom) **star:1361** Go Manager - bundle for go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [gpm](https://github.com/pote/gpm) **star:1204** Barebones dependency manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pote/gpm)
* [goop](https://github.com/nitrous-io/goop) **star:779** Simple dependency manager for Go (golang), inspired by Bundler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nitrous-io/goop)   [![godoc][GoDoc]](https://godoc.org/github.com/nitrous-io/goop)
* [nut](https://github.com/jingweno/nut) **star:245** Vendor Go dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jingweno/nut)   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** Minimal dependency version using Git.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/johnny-deps)
* [gigo](https://github.com/LyricalSecurity/gigo) **star:197** PIP-like dependency tool for golang, with support for private repositories and hashes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LyricalSecurity/gigo)   [![godoc][GoDoc]](https://godoc.org/github.com/LyricalSecurity/gigo)
* [VenGO](https://github.com/DamnWidget/VenGO) **star:116** create and manage exportable isolated go virtual environments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DamnWidget/VenGO)   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:94** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.   [![There was an update last week][Green]](https://github.com/raydac/mvn-golang)
* [gop](https://github.com/lunny/gop) **star:50** Build and manage your Go applications out of GOPATH.   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/gop)   [![Contains Chinese documents][CN]](https://github.com/lunny/gop)   [![Archived][Archived]](https://github.com/lunny/gop)

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:9726** A distributed tracing system.   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   [![There was an update last week][Green]](https://github.com/jaegertracing/jaeger)   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1136** Simple profiling support package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [tracer](https://github.com/kamilsk/tracer) **star:16** Simple, lightweight tracing.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:5726** Implementation of GraphQL for Go.   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   [![There was an update last week][Green]](https://github.com/graphql-go/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [graphql](https://github.com/neelance/graphql-go) **star:3048** GraphQL server with a focus on ease of use.   [![star > 2000][Awesome]](https://github.com/neelance/graphql-go)   [![godoc][GoDoc]](https://godoc.org/github.com/neelance/graphql-go)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:1078** A simple Go package to Query over JSON Data.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [jsonql](https://github.com/elgs/jsonql) **star:216** JSON query expression library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/jsonql)
* [rql](https://github.com/a8m/rql) **star:128** Resource Query Language for REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/a8m/rql)
* [graphql](https://github.com/tmc/graphql) **star:51** graphql parser + utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tmc/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [jsonslice](https://github.com/bhmj/jsonslice) **star:29** Jsonpath queries with advanced filters.   [![There was an update last week][Green]](https://github.com/bhmj/jsonslice)   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)
* [straf](https://github.com/SonicRoshan/straf) **star:5** Easily Convert Golang structs to GraphQL objects.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/straf)

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) **star:2506** The simple and easy way to embed static files into Go binaries.   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![There was an update last week][Green]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik) **star:2375** Embeds static files into a Go executable.   [![star > 2000][Awesome]](https://github.com/rakyll/statik)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/statik)
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1782** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:751** Generates a vfsdata.go file that statically implements the given virtual filesystem.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/vfsgen)
* [esc](https://github.com/mjibson/esc) **star:505** Embeds files into Go programs and provides http.FileSystem interfaces to them.   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:483** Simple tool to embed files in go with focus on "customization" and ease to use.   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-resources](https://github.com/omeid/go-resources) **star:158** Unfancy resources embedding with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/omeid/go-resources)
* [statics](https://github.com/go-playground/statics) **star:54** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/statics)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/statics)
* [templify](https://github.com/wlbr/templify) **star:22** Embed external template files into Go code to create single file binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/templify)
* [go-embed](https://github.com/pyros2097/go-embed) **star:18** Generates go code to embed resource files into your library or executable.   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:3322** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   [![There was an update last week][Green]](https://github.com/gonum/gonum)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [stats](https://github.com/montanaflynn/stats) **star:1408** Statistics package with common functions missing from the Golang standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gosl](https://github.com/cpmech/gosl) **star:1359** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![There was an update last week][Green]](https://github.com/cpmech/gosl)   [![godoc][GoDoc]](https://godoc.org/github.com/cpmech/gosl)
* [streamtools](https://github.com/nytlabs/streamtools) **star:1318** general purpose, graphical tool for dealing with streams of data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nytlabs/streamtools)   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [gonum/plot](https://github.com/gonum/plot) **star:1307** gonum/plot provides an API for building and drawing plots in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [go-dsp](https://github.com/mjibson/go-dsp) **star:651** Digital Signal Processing for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mjibson/go-dsp)   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/go-dsp)
* [goraph](https://github.com/gyuho/goraph) **star:614** Pure Go graph theory library(data structure, algorith visualization).   [![It hasn't been updated in the last year][Yellow]](https://github.com/gyuho/goraph)   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [chart](https://github.com/vdobler/chart) **star:607** Simple Chart Plotting library for Go. Supports many graphs types.   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [graph](https://github.com/yourbasic/graph) **star:279** Library of basic graph algorithms.   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [ewma](https://github.com/VividCortex/ewma) **star:277** Exponentially-weighted moving averages.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [orb](https://github.com/paulmach/orb) **star:234** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:136** Approximate histograms for data streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/gohistogram)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:121** Dataframes for machine-learning and statistics (similar to pandas).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:86** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DavidBelicza/TextRank)   [![godoc][GoDoc]](https://godoc.org/github.com/DavidBelicza/TextRank)
* [sparse](https://github.com/james-bowman/sparse) **star:78** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![There was an update last week][Green]](https://github.com/james-bowman/sparse)   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [pagerank](https://github.com/alixaxel/pagerank) **star:53** Weighted PageRank algorithm implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [geom](https://github.com/skelterjohn/geom) **star:43** 2D geometry for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/geom)   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [evaler](https://github.com/soniah/evaler) **star:40** Simple floating point arithmetic expression evaluator.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/evaler)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:14** GO Implementation of Entropy Measures.   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:12** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [ode](https://github.com/ChristopherRabotin/ode) **star:11** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/ode)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/ode)
* [GoStats](https://github.com/OGFris/GoStats) **star:11** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:10** Tiny linear interpolation library.   [![There was an update last week][Green]](https://github.com/sgreben/piecewiselinear)   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/piecewiselinear)
* [PiHex](https://github.com/claygod/PiHex) **star:8** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [assocentity](https://github.com/ndabAP/assocentity) **star:6** Package assocentity returns the average distance from words to a given entity.   [![godoc][GoDoc]](https://godoc.org/github.com/ndabAP/assocentity)
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** Graph theory algorithms written in "Go" language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/go-gt)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/go-gt)
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:2** Provides a Bradley-Terry Model for pairwise comparisons.   [![godoc][GoDoc]](https://godoc.org/github.com/seanhagen/bradleyterry)

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/xenolf/lego) **star:3724** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   [![star > 2000][Awesome]](https://github.com/xenolf/lego)   [![There was an update last week][Green]](https://github.com/xenolf/lego)   [![godoc][GoDoc]](https://godoc.org/github.com/xenolf/lego)
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:1964** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![There was an update last week][Green]](https://github.com/Ullaakut/cameradar)   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [acmetool](https://github.com/hlandau/acme) **star:1721** ACME (Let's Encrypt) client tool with automatic renewal.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/acme)
* [memguard](https://github.com/awnumar/memguard) **star:1633** A pure Go library for handling sensitive values in memory.   [![godoc][GoDoc]](https://godoc.org/github.com/awnumar/memguard)
* [secure](https://github.com/unrolled/secure) **star:1372** HTTP middleware for Go that facilitates some quick security wins.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/secure)
* [acra](https://github.com/cossacklabs/acra) **star:494** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   [![There was an update last week][Green]](https://github.com/cossacklabs/acra)   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [nacl](https://github.com/kevinburke/nacl) **star:460** Go implementation of the NaCL set of API's.   [![godoc][GoDoc]](https://godoc.org/github.com/kevinburke/nacl)
* [BadActor](https://github.com/jaredfolkins/badactor) **star:260** In-memory, application-driven jailer built in the spirit of fail2ban.   [![godoc][GoDoc]](https://godoc.org/github.com/jaredfolkins/badactor)
* [passlib](https://github.com/hlandau/passlib) **star:230** Futureproof password hashing library.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/passlib)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:209** encrypt/decrypt using ssh keys.   [![godoc][GoDoc]](https://godoc.org/github.com/ssh-vault/ssh-vault)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:160** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [go-yara](https://github.com/hillu/go-yara) **star:144** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [argon2pw](https://github.com/raja/argon2pw) **star:78** Argon2 password hash generation with constant-time password comparison.   [![It hasn't been updated in the last year][Yellow]](https://github.com/raja/argon2pw)   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:34** A probably paranoid package for securely hashing and encrypting passwords.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [certificates](https://github.com/mvmaasakkers/certificates) **star:10** An opinionated tool for generating tls certificates.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:7** SSL certificates made easy with a high level wrapper around acme/autocert.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianosela/sslmgr)

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:6686** High-performance 100% compatible drop-in replacement of "encoding/json".   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![There was an update last week][Green]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:5827** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:3335** Protocol Buffers for Go with Gadgets.   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:2845** Go library for decoding generic map values into native Go structures.   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1315** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   [![There was an update last week][Green]](https://github.com/ugorji/go)   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [colfer](https://github.com/pascaldekloe/colfer) **star:496** Code generation for the Colfer binary format.
* [csvutil](https://github.com/jszwec/csvutil) **star:332** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![There was an update last week][Green]](https://github.com/jszwec/csvutil)   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:274** Cap'n Proto library and parser for go.   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:130** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yvasiyarov/php_session_decoder)   [![godoc][GoDoc]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
* [structomap](https://github.com/tuvistavie/structomap) **star:103** Library to easily and dynamically generate maps from static structures.   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [cbor](https://github.com/fxamacker/cbor) **star:83** Small, safe, and easy CBOR encoding and decoding library.   [![There was an update last week][Green]](https://github.com/fxamacker/cbor)   [![godoc][GoDoc]](https://godoc.org/github.com/fxamacker/cbor)
* [bambam](https://github.com/glycerine/bambam) **star:60** generator for Cap'n Proto schemas from go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/glycerine/bambam)   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/bambam)
* [asn1](https://github.com/PromonLogicalis/asn1) **star:43** Asn.1 BER and DER encoding library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/PromonLogicalis/asn1)   [![Archived][Archived]](https://github.com/PromonLogicalis/asn1)
* [binstruct](https://github.com/ghostiam/binstruct) **star:13** Golang binary decoder for mapping data into the structure.   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [fwencoder](https://github.com/o1egl/fwencoder) **star:10** Fixed width file parser (encoding and decoding library) for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/o1egl/fwencoder)   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/fwencoder)
* [bel](https://github.com/32leaves/bel) **star:7** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.   [![godoc][GoDoc]](https://godoc.org/github.com/32leaves/bel)
* [pletter](https://github.com/vimeda/pletter) **star:7** A standard way to wrap a proto message for message brokers.   [![godoc][GoDoc]](https://godoc.org/github.com/vimeda/pletter)

## Server Applications

* [etcd](https://github.com/coreos/etcd) **star:28563** Highly-available key value store for shared configuration and service discovery.   [![star > 2000][Awesome]](https://github.com/coreos/etcd)   [![There was an update last week][Green]](https://github.com/coreos/etcd)   [![godoc][GoDoc]](https://godoc.org/github.com/coreos/etcd)
* [Caddy](https://github.com/mholt/caddy) **star:25363** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   [![star > 2000][Awesome]](https://github.com/mholt/caddy)   [![There was an update last week][Green]](https://github.com/mholt/caddy)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/caddy)
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [minio](https://github.com/minio/minio) **star:19163** Minio is a distributed object storage server.   [![star > 2000][Awesome]](https://github.com/minio/minio)   [![There was an update last week][Green]](https://github.com/minio/minio)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [RoadRunner](https://github.com/spiral/roadrunner) **star:3750** High-performance PHP application server, load-balancer and process manager.   [![star > 2000][Awesome]](https://github.com/spiral/roadrunner)   [![There was an update last week][Green]](https://github.com/spiral/roadrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/spiral/roadrunner)
* [devd](https://github.com/cortesi/devd) **star:2888** Local webserver for developers.   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [algernon](https://github.com/xyproto/algernon) **star:1636** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   [![There was an update last week][Green]](https://github.com/xyproto/algernon)   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/algernon)
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:1155** Full featured and highly configurable SFTP server software.   [![There was an update last week][Green]](https://github.com/drakkan/sftpgo)   [![godoc][GoDoc]](https://godoc.org/github.com/drakkan/sftpgo)
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.
* [flipt](https://github.com/markphelps/flipt) **star:1120** A self contained feature flag solution written in Go and Vue.js   [![There was an update last week][Green]](https://github.com/markphelps/flipt)   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Flagr](https://github.com/checkr/flagr) **star:951** Flagr is an open-source feature flagging and A/B testing service.   [![There was an update last week][Green]](https://github.com/checkr/flagr)   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [Fider](https://github.com/getfider/fider) **star:901** Fider is an open platform to collect and organize customer feedback.   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [discovery](https://github.com/Bilibili/discovery) **star:764** A registry for resilient mid-tier load balancing and failover.   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [jackal](https://github.com/ortuman/jackal) **star:755** An XMPP server written in Go.   [![There was an update last week][Green]](https://github.com/ortuman/jackal)   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [dudeldu](https://github.com/krotik/dudeldu) **star:103** A simple SHOUTcast server.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/dudeldu)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:10** Stream database events from PostgreSQL to Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:8** Nginx log parser and exporter to Prometheus.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.

## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:244** Go stream processing library.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3418** PDF document generator with high level support for text, drawing and images.   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)   [![Archived][Archived]](https://github.com/jung-kurt/gofpdf)
* [pongo2](https://github.com/flosch/pongo2) **star:1567** Django-like template-engine for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1544** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [hero](https://github.com/shiyanhui/hero) **star:1251** Hero is a handy, fast and powerful go template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/shiyanhui/hero)   [![Contains Chinese documents][CN]](https://github.com/shiyanhui/hero)
* [mustache](https://github.com/hoisie/mustache) **star:972** Go implementation of the Mustache template language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hoisie/mustache)   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:835** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eknkc/amber)   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:772** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/ace)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:722** Razor view engine for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [jet](https://github.com/CloudyKit/jet) **star:604** Jet template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/CloudyKit/jet)
* [ego](https://github.com/benbjohnson/ego) **star:422** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [raymond](https://github.com/aymerick/raymond) **star:353** Complete handlebars implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:330** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasttemplate)
* [Soy](https://github.com/robfig/soy) **star:145** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [liquid](https://github.com/osteele/liquid) **star:92** Go implementation of Shopify Liquid templates.   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/liquid)
* [goview](https://github.com/foolin/goview) **star:84** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/goview)
* [maroto](https://github.com/johnfercher/maroto) **star:80** A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.   [![godoc][GoDoc]](https://godoc.org/github.com/johnfercher/maroto)
* [kasia.go](https://github.com/ziutek/kasia.go) **star:71** Templating system for HTML and other text documents - go implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/kasia.go)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [velvet](https://github.com/gobuffalo/velvet) **star:68** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gobuffalo/velvet)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)
* [damsel](https://github.com/dskinner/damsel) **star:21** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dskinner/damsel)   [![godoc][GoDoc]](https://godoc.org/github.com/dskinner/damsel)
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:16** Tiny wrapper around html/template to allow for easy file-based template inheritance.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dannyvankooten/extemplate)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/extemplate)
* [gospin](https://github.com/m1/gospin) **star:5** Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations.   [![There was an update last week][Green]](https://github.com/m1/gospin)   [![godoc][GoDoc]](https://godoc.org/github.com/m1/gospin)

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [Testify](https://github.com/stretchr/testify) **star:9065** Sacred extension to the standard go testing package.   [![star > 2000][Awesome]](https://github.com/stretchr/testify)   [![There was an update last week][Green]](https://github.com/stretchr/testify)   [![godoc][GoDoc]](https://godoc.org/github.com/stretchr/testify)
    * [go-cmp](https://github.com/google/go-cmp) **star:1438** Package for comparing Go values in tests.   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1213** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   [![godoc][GoDoc]](https://godoc.org/github.com/gavv/httpexpect)
    * [godog](https://github.com/DATA-DOG/godog) **star:852** Cucumber or Behat like BDD framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/godog)
    * [baloo](https://github.com/h2non/baloo) **star:661** Expressive and versatile end-to-end HTTP API testing made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [goblin](https://github.com/franela/goblin) **star:648** Mocha like testing framework fo Go.   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:369** A helper for Rails' like test fixtures to test database applications.   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:357** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![There was an update last week][Green]](https://github.com/dnaeon/go-vcr)   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:319** Mutation testing for Go source code.   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gofight](https://github.com/appleboy/gofight) **star:279** API Handler Testing for Golang Router framework.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [frisby](https://github.com/verdverm/frisby) **star:255** REST API testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:200** Tool for viewing test coverage in terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [charlatan](https://github.com/percolate/charlatan) **star:192** Tool to generate fake interface implementations for tests.   [![godoc][GoDoc]](https://godoc.org/github.com/percolate/charlatan)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:138** A collection of packages to augment the go testing package and support common patterns.   [![There was an update last week][Green]](https://github.com/gotestyourself/gotest.tools)   [![godoc][GoDoc]](https://godoc.org/github.com/gotestyourself/gotest.tools)
    * [endly](https://github.com/viant/endly) **star:114** Declarative end to end functional testing.   [![There was an update last week][Green]](https://github.com/viant/endly)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:112** BDD-style testing framework for the Go programming language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/orfjackal/gospec)   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:95** Simple snapshot testing addon for your test framework.   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:92** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:75** Extremely flexible golang deep comparison, extends the go testing package.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [wstest](https://github.com/posener/wstest) **star:68** Websocket client for unit-testing a websocket http.Handler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/posener/wstest)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [gospecify](https://github.com/stesla/gospecify) **star:53** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stesla/gospecify)   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [restit](https://github.com/yookoala/restit) **star:49** Go micro framework to help writing RESTful API integration test.   [![godoc][GoDoc]](https://godoc.org/github.com/yookoala/restit)
    * [commander](https://github.com/SimonBaeumer/commander) **star:40** Tool for testing cli applications on windows, linux and osx.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:32** library created for testing JSON against patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:28** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:27** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rdrdr/hamcrest)   [![godoc][GoDoc]](https://godoc.org/github.com/rdrdr/hamcrest)
    * [dsunit](https://github.com/viant/dsunit) **star:26** Datastore testing for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [assert](https://github.com/go-playground/assert) **star:16** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/assert)
    * [testcase](https://github.com/adamluzsi/testcase) **star:12** Idiomatic testing framework for Behavior Driven Development.   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:11** HTTP client testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Extensions to Go's `testing/iotest` package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cavaliercoder/badio)   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/badio)
    * [gocrest](https://github.com/corbym/gocrest) **star:9** Composable hamcrest-like matchers for Go assertions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gocrest)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gocrest)
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pavlo/gosuite)   [![godoc][GoDoc]](https://godoc.org/github.com/pavlo/gosuite)
    * [gogiven](https://github.com/corbym/gogiven) **star:8** YATSPEC-like BDD testing framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gogiven)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gogiven)
    * [schema](https://github.com/jgroeneveld/schema) **star:7** Quick and easy expression matching for JSON schemas used in requests and responses.   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/schema)
    * [biff](https://github.com/fulldump/biff) **star:7** Bifurcation testing framework, BDD compatible.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fulldump/biff)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/biff)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [Tt](https://github.com/vcaesar/tt) **star:6** Simple and colorful test tools.   [![godoc][GoDoc]](https://godoc.org/github.com/vcaesar/tt)
    * [trial](https://github.com/jgroeneveld/trial) **star:4** Quick and easy extendable assertions without introducing much boilerplate.   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/trial)

* Mock
    * [gomock](https://github.com/golang/mock) **star:3344** Mocking framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/golang/mock)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/mock)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:2086** Mock SQL driver for testing database interactions.   [![star > 2000][Awesome]](https://github.com/DATA-DOG/go-sqlmock)   [![There was an update last week][Green]](https://github.com/DATA-DOG/go-sqlmock)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1493** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   [![There was an update last week][Green]](https://github.com/SpectoLabs/hoverfly)   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:921** Versatile HTTP mocking made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [httpmock](https://github.com/jarcoal/httpmock) **star:667** Easy mocking of HTTP responses from external resources.   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:386** Tool for generating self-contained mock objects.   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [minimock](https://github.com/gojuno/minimock) **star:286** Mock generator for Go interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:199** Single transaction based database driver mainly for testing purposes.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:86** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tv42/mockhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/tv42/mockhttp)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:3220** Randomized testing system.   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![There was an update last week][Green]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz) **star:676** Library for populating go objects with random values.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gofuzz)
    * [Tavor](https://github.com/zimmski/tavor) **star:213** Generic fuzzing and delta-debugging framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zimmski/tavor)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:4026** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1412** alternative Selenium hub server that launches browsers within containers.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [cdp](https://github.com/mafredri/cdp) **star:391** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:230** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   [![There was an update last week][Green]](https://github.com/aerokube/ggr)   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:457** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:9445** Fast and Elegant Scraping Framework for Gophers.   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![There was an update last week][Green]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:8090** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [blackfriday](https://github.com/russross/blackfriday) **star:4119** Markdown processor in Go.   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![There was an update last week][Green]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [toml](https://github.com/BurntSushi/toml) **star:2955** TOML configuration format (encoder/decoder with reflection).   [![star > 2000][Awesome]](https://github.com/BurntSushi/toml)   [![godoc][GoDoc]](https://godoc.org/github.com/BurntSushi/toml)
    * [sh](https://github.com/mvdan/sh) **star:2215** Shell parser and formatter.   [![star > 2000][Awesome]](https://github.com/mvdan/sh)   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/sh)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:2022** Formatters for time, numbers, and memory size to human readable format.   [![star > 2000][Awesome]](https://github.com/dustin/go-humanize)   [![There was an update last week][Green]](https://github.com/dustin/go-humanize)   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1342** HTML Sanitizer.   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [inject](https://github.com/facebookgo/inject) **star:1179** Package inject provides a reflect based injector.   [![godoc][GoDoc]](https://godoc.org/github.com/facebookgo/inject)   [![Archived][Archived]](https://github.com/facebookgo/inject)
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1163** Parse RSS and Atom feeds in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [go-toml](https://github.com/pelletier/go-toml) **star:670** Go library for the TOML format with query support and handy cli tools.   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:566** A collection of common regular expressions for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [slug](https://github.com/gosimple/slug) **star:470** URL-friendly slugify with multiple languages support.   [![godoc][GoDoc]](https://godoc.org/github.com/gosimple/slug)
    * [mxj](https://github.com/clbanning/mxj) **star:348** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:334** Web scraping Framework to turn websites into structured data.   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:331** Parses the Graphviz DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [gotext](https://github.com/leonelquinteros/gotext) **star:244** GNU gettext utilities for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/leonelquinteros/gotext)
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:234** Functions to get fixed width of the character or string.   [![There was an update last week][Green]](https://github.com/mattn/go-runewidth)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-runewidth)
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:172** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/htmlquery)
    * [goq](https://github.com/andrewstuart/goq) **star:157** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).   [![godoc][GoDoc]](https://godoc.org/github.com/andrewstuart/goq)
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:108** NMEA parser library for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [sdp](https://github.com/gortc/sdp) **star:77** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [align](https://github.com/Guitarbum722/align) **star:60** A general purpose application that aligns text.   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:57** Make pretty slug with multiple languages support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-slugify)
    * [genex](https://github.com/alixaxel/genex) **star:56** Count and expand Regular Expressions into all matching Strings.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:53** Editorconfig file parser and manipulator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [goribot](https://github.com/zhshch2002/goribot) **star:50** A simple golang spider/scraping framework,build a spider in 3 lines.   [![godoc][GoDoc]](https://godoc.org/github.com/zhshch2002/goribot)   [![Contains Chinese documents][CN]](https://github.com/zhshch2002/goribot)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:45** Zero-width character detection and removal for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/trubitsyn/go-zero-width)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:45** Functions to determine the natural language of a unicode text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/guesslanguage)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/guesslanguage)
    * [allot](https://github.com/sbstjn/allot) **star:38** Placeholder and wildcard text parsing for CLI tools and bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:37** Library for generating random strings from regular expressions.   [![godoc][GoDoc]](https://godoc.org/github.com/zach-klippenstein/goregen)
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:33** Fixed-width text formatting (encoder/decoder with reflection).   [![There was an update last week][Green]](https://github.com/ianlopshire/go-fixedwidth)   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [gonameparts](https://github.com/polera/gonameparts) **star:29** Parses human names into individual name parts.   [![godoc][GoDoc]](https://godoc.org/github.com/polera/gonameparts)
    * [Slugify](https://github.com/avelino/slugify) **star:27** Go slugify application that handles string.   [![It hasn't been updated in the last year][Yellow]](https://github.com/avelino/slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/avelino/slugify)
    * [did](https://github.com/ockam-network/did) **star:26** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [go-vcard](https://github.com/emersion/go-vcard) **star:26** Parse and format vCard.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [codetree](https://github.com/aerogo/codetree) **star:13** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [enca](https://github.com/endeveit/enca) **star:8** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/enca)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/enca)
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:7** A syndication feed for Atom 1.0 and RSS 2.0.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhengchun/syndfeed)   [![godoc][GoDoc]](https://godoc.org/github.com/zhengchun/syndfeed)
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/CalebQ42/bbConvert)   [![godoc][GoDoc]](https://godoc.org/github.com/CalebQ42/bbConvert)
    * [doi](https://github.com/hscells/doi) **star:4** Document object identifier (doi) parser in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hscells/doi)   [![godoc][GoDoc]](https://godoc.org/github.com/hscells/doi)
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
    * [ltsv](https://github.com/Wing924/ltsv) **star:3** High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/ltsv)
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:548** Extract urls from text.   [![There was an update last week][Green]](https://github.com/mvdan/xurls)   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/xurls)
    * [gotabulate](https://github.com/bndr/gotabulate) **star:218** Easily pretty-print your tabular data with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gotabulate)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:149** fast string sorting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/radix)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [parth](https://github.com/codemodus/parth) **star:34** URL path segmentation parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/parth)
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:18** Convert xml or json to go struct.   [![godoc][GoDoc]](https://godoc.org/github.com/stackerzzq/xj2go)
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:17** A sanitization-based swear filter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/JoshuaDoes/gofuckyourself)   [![godoc][GoDoc]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
    * [kace](https://github.com/codemodus/kace) **star:12** Common case conversions covering common initialisms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/kace)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/kace)
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:7** string argument parser that understands quotes and backslashes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nproc/parseargs-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nproc/parseargs-go)
    * [Tagify](https://github.com/zoomio/tagify) **star:3** Produces a set of tags from given source.   [![There was an update last week][Green]](https://github.com/zoomio/tagify)   [![godoc][GoDoc]](https://godoc.org/github.com/zoomio/tagify)
    * [TySug](https://github.com/Dynom/TySug) **star:3** Alternative suggestions with respect to keyboard layouts.   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
	* [textwrap](https://github.com/isbm/textwrap) **star:1** Implementation of `textwrap` module from Python.   [![godoc][GoDoc]](https://godoc.org/github.com/isbm/textwrap)

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5386** The official AWS SDK for the Go programming language.   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   [![There was an update last week][Green]](https://github.com/aws/aws-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [github](https://github.com/google/go-github) **star:5301** Go library for accessing the GitHub REST API v3.   [![star > 2000][Awesome]](https://github.com/google/go-github)   [![There was an update last week][Green]](https://github.com/google/go-github)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [slack](https://github.com/nlopes/slack) **star:2642** Slack API in Go.   [![star > 2000][Awesome]](https://github.com/nlopes/slack)   [![godoc][GoDoc]](https://godoc.org/github.com/nlopes/slack)
* [google](https://github.com/google/google-api-go-client) **star:2051** Auto-generated Google APIs for Go.   [![star > 2000][Awesome]](https://github.com/google/google-api-go-client)   [![There was an update last week][Green]](https://github.com/google/google-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:1939** Google Cloud APIs Go Client Library.   [![There was an update last week][Green]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
* [discordgo](https://github.com/bwmarrin/discordgo) **star:1069** Go bindings for the Discord Chat API.   [![godoc][GoDoc]](https://godoc.org/github.com/bwmarrin/discordgo)
* [stripe](https://github.com/stripe/stripe-go) **star:1022** Go client for the Stripe API.   [![There was an update last week][Green]](https://github.com/stripe/stripe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1007** Go client library for the Twitter 1.1 API.   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [go-twitter](https://github.com/dghubble/go-twitter) **star:862** Go client library for the Twitter v1.1 APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/go-twitter)
* [facebook](https://github.com/huandu/facebook) **star:815** Go Library that supports the Facebook Graph API.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/facebook)
* [minio-go](https://github.com/minio/minio-go) **star:813** Minio Go Library for Amazon S3 compatible cloud storage.   [![There was an update last week][Green]](https://github.com/minio/minio-go)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio-go)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:643** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   [![There was an update last week][Green]](https://github.com/andygrunwald/go-jira)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-jira)
* [githubql](https://github.com/shurcooL/githubql) **star:568** Go library for accessing the GitHub GraphQL API v4.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/githubql)
* [webhooks](https://github.com/go-playground/webhooks) **star:419** Webhook receiver for GitHub and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/webhooks)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:331** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:328** Wrapper for PayPal payment API.   [![There was an update last week][Green]](https://github.com/logpacker/PayPal-Go-SDK)   [![godoc][GoDoc]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
* [go-marathon](https://github.com/gambol99/go-marathon) **star:192** Go library for interacting with Mesosphere's Marathon PAAS.   [![There was an update last week][Green]](https://github.com/gambol99/go-marathon)   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [ethrpc](https://github.com/onrik/ethrpc) **star:179** Go bindings for Ethereum JSON RPC API.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [gostorm](https://github.com/jsgilmore/gostorm) **star:123** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jsgilmore/gostorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [Trello](https://github.com/adlio/trello) **star:121** Go wrapper for the Trello API.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [Medium](https://github.com/Medium/medium-sdk-go) **star:120** Golang SDK for Medium's OAuth2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Medium/medium-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/Medium/medium-sdk-go)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:112** A golang package to communicate with HipChat over XMPP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daneharrigan/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [hipchat](https://github.com/andybons/hipchat) **star:108** This project implements a golang client library for the Hipchat API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andybons/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/andybons/hipchat)
* [go-trending](https://github.com/andygrunwald/go-trending) **star:104** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-trending)
* [cachet](https://github.com/andygrunwald/cachet) **star:73** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/cachet)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/cachet)
* [pushover](https://github.com/gregdel/pushover) **star:63** Go wrapper for the Pushover API.   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [megos](https://github.com/andygrunwald/megos) **star:57** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/megos)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Go client library for interfacing with the Clarifai API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/samuelcouch/clarifai)   [![godoc][GoDoc]](https://godoc.org/github.com/samuelcouch/clarifai)
* [wit-go](https://github.com/wit-ai/wit-go) **star:57** Go client for wit.ai HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:53** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [gads](https://github.com/emiddleton/gads) **star:46** Google Adwords Unofficial API.   [![godoc][GoDoc]](https://godoc.org/github.com/emiddleton/gads)
* [circleci](https://github.com/jszwedko/go-circleci) **star:46** Go client library for interacting with CircleCI's API.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:40** Go client for the xkcd API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nishanths/go-xkcd)   [![godoc][GoDoc]](https://godoc.org/github.com/nishanths/go-xkcd)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:39** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:37** Go MusicBrainz WS2 client library.   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:36** Go wrapper and command-line client for the Uptime Robot v2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [golyrics](https://github.com/mamal72/golyrics) **star:33** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [fcm](https://github.com/maddevsio/fcm) **star:33** Go library for Firebase Cloud Messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [mixpanel](https://github.com/dukex/mixpanel) **star:30** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dukex/mixpanel)   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [translate](https://github.com/poorny/translate) **star:29** Go online translation package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/poorny/translate)   [![godoc][GoDoc]](https://godoc.org/github.com/poorny/translate)
* [gcm](https://github.com/Aorioli/gcm) **star:29** Go library for Google Cloud Messaging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Aorioli/gcm)   [![godoc][GoDoc]](https://godoc.org/github.com/Aorioli/gcm)
* [gami](https://github.com/bit4bit/gami) **star:27** Go library for Asterisk Manager Interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![Archived][Archived]](https://github.com/bit4bit/gami)
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:24** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![There was an update last week][Green]](https://github.com/hbagdi/go-unsplash)   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [spotify](https://github.com/rapito/go-spotify) **star:20** Go Library to access Spotify WEB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-spotify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-spotify)
* [shopify](https://github.com/rapito/go-shopify) **star:19** Go Library to make CRUD request to the Shopify API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-shopify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:19** Go wrapper for the YNAB API.   [![godoc][GoDoc]](https://godoc.org/github.com/brunomvsouza/ynab.go)
* [patreon-go](https://github.com/mxpv/patreon-go) **star:19** Go library for Patreon API.   [![godoc][GoDoc]](https://godoc.org/github.com/mxpv/patreon-go)
* [steam](https://github.com/sostronk/go-steam) **star:18** Go Library to interact with Steam game servers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sostronk/go-steam)   [![godoc][GoDoc]](https://godoc.org/github.com/sostronk/go-steam)
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:18** Go client for interacting with the Twitch v3 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/knspriggs/go-twitch)   [![godoc][GoDoc]](https://godoc.org/github.com/knspriggs/go-twitch)
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:16** Go library for accessing the BreweryDB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/naegelejd/brewerydb)   [![godoc][GoDoc]](https://godoc.org/github.com/naegelejd/brewerydb)
* [codeship-go](https://github.com/codeship/codeship-go) **star:16** Go client library for interacting with Codeship's API v2.   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [textbelt](https://github.com/dietsche/textbelt) **star:16** Go client for the textbelt.com txt messaging API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:15** Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).   [![It hasn't been updated in the last year][Yellow]](https://github.com/nstratos/go-myanimelist)   [![godoc][GoDoc]](https://godoc.org/github.com/nstratos/go-myanimelist)
* [simples3](https://github.com/rhnvrm/simples3) **star:14** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:14** Go client library for [imgur](https://imgur.com)   [![It hasn't been updated in the last year][Yellow]](https://github.com/koffeinsource/go-imgur)   [![godoc][GoDoc]](https://godoc.org/github.com/koffeinsource/go-imgur)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:11** Simple wrapper for easy google analytics reporting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chonthu/go-google-analytics)   [![godoc][GoDoc]](https://godoc.org/github.com/chonthu/go-google-analytics)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:10** Tiny Go client for HackerNews API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PaulRosset/go-hacknews)   [![godoc][GoDoc]](https://godoc.org/github.com/PaulRosset/go-hacknews)
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergiotapia/smitego)   [![godoc][GoDoc]](https://godoc.org/github.com/sergiotapia/smitego)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:7** Go client library for the [LastPass](https://www.lastpass.com/) API.   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:7** Go client library for the SPTrans Olho Vivo API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergioaugrod/go-sptrans)   [![godoc][GoDoc]](https://godoc.org/github.com/sergioaugrod/go-sptrans)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-google-email-audit-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-google-email-audit-api)
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattcunningham/gumblr)   [![godoc][GoDoc]](https://godoc.org/github.com/mattcunningham/gumblr)
* [go-here](https://github.com/abdullahselek/go-here) **star:5** Go client library around the HERE location based APIs.   [![There was an update last week][Green]](https://github.com/abdullahselek/go-here)   [![godoc][GoDoc]](https://godoc.org/github.com/abdullahselek/go-here)
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/esurdam/go-sophos)
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gojuno/go-zooz)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:3** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   [![It hasn't been updated in the last year][Yellow]](https://github.com/axelspringer/go-chronos)   [![godoc][GoDoc]](https://godoc.org/github.com/axelspringer/go-chronos)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** The Playlyfe Rest API Go SDK.   [![It hasn't been updated in the last year][Yellow]](https://github.com/playlyfe/playlyfe-go-sdk)   [![godoc][GoDoc]](https://godoc.org/github.com/playlyfe/playlyfe-go-sdk)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.   [![godoc][GoDoc]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
* [vl-go](https://github.com/verifid/vl-go) **star:1** Go client library around the VerifID identity verification layer API.   [![godoc][GoDoc]](https://godoc.org/github.com/verifid/vl-go)
* [libgoffi](https://github.com/clevabit/libgoffi) **star:1** Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:25647** Command-line fuzzy finder written in Go.   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![There was an update last week][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [hub](https://github.com/github/hub) **star:18091** wrap git commands with additional functionality to interact with github from the terminal.   [![star > 2000][Awesome]](https://github.com/github/hub)   [![There was an update last week][Green]](https://github.com/github/hub)   [![godoc][GoDoc]](https://godoc.org/github.com/github/hub)
* [delve](https://github.com/derekparker/delve) **star:12909** Go debugger.   [![star > 2000][Awesome]](https://github.com/derekparker/delve)   [![There was an update last week][Green]](https://github.com/derekparker/delve)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [ctop](https://github.com/bcicen/ctop) **star:9209** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [wuzz](https://github.com/asciimoo/wuzz) **star:8460** Interactive cli tool for HTTP inspection.   [![star > 2000][Awesome]](https://github.com/asciimoo/wuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/wuzz)
* [sqlx](https://github.com/jmoiron/sqlx) **star:7420** provides a set of extensions on top of the excellent built-in database/sql package.   [![star > 2000][Awesome]](https://github.com/jmoiron/sqlx)   [![There was an update last week][Green]](https://github.com/jmoiron/sqlx)   [![godoc][GoDoc]](https://godoc.org/github.com/jmoiron/sqlx)
* [peco](https://github.com/peco/peco) **star:5640** Simplistic interactive filtering tool.   [![star > 2000][Awesome]](https://github.com/peco/peco)   [![godoc][GoDoc]](https://godoc.org/github.com/peco/peco)
* [goreporter](https://github.com/wgliang/goreporter)  Golang tool that does static analysis, unit testing, code review and generate code quality report.
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:4895** Deliver Go binaries as fast and easily as possible.   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![There was an update last week][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [godropbox](https://github.com/dropbox/godropbox) **star:3793** Common libraries for writing Go services/applications from Dropbox.   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2236** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [panicparse](https://github.com/maruel/panicparse) **star:2173** Groups similar goroutines and colorizes stack dump.   [![star > 2000][Awesome]](https://github.com/maruel/panicparse)   [![godoc][GoDoc]](https://godoc.org/github.com/maruel/panicparse)
* [Task](https://github.com/go-task/task) **star:2116** simple "Make" alternative.   [![star > 2000][Awesome]](https://github.com/go-task/task)   [![There was an update last week][Green]](https://github.com/go-task/task)   [![godoc][GoDoc]](https://godoc.org/github.com/go-task/task)
* [minify](https://github.com/tdewolff/minify) **star:1953** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   [![There was an update last week][Green]](https://github.com/tdewolff/minify)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/minify)
* [mmake](https://github.com/tj/mmake) **star:1468** Modern Make.   [![godoc][GoDoc]](https://godoc.org/github.com/tj/mmake)
* [Storm](https://github.com/asdine/storm) **star:1449** Simple and powerful toolkit for BoltDB.   [![godoc][GoDoc]](https://godoc.org/github.com/asdine/storm)
* [go-funk](https://github.com/thoas/go-funk) **star:1437** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [mole](https://github.com/davrodpin/mole) **star:1342** cli app to easily create ssh tunnels.   [![godoc][GoDoc]](https://godoc.org/github.com/davrodpin/mole)
* [mc](https://github.com/minio/mc) **star:1249** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   [![There was an update last week][Green]](https://github.com/minio/mc)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/mc)
* [boilr](https://github.com/tmrts/boilr) **star:1009** Blazingly fast CLI tool for creating projects from boilerplate templates.   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [filetype](https://github.com/h2non/filetype) **star:1009** Small package to infer the file type checking the magic numbers signature.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [mergo](https://github.com/imdario/mergo) **star:962** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   [![godoc][GoDoc]](https://godoc.org/github.com/imdario/mergo)
* [spinner](https://github.com/briandowns/spinner) **star:879** Go package to easily provide a terminal spinner with options.   [![godoc][GoDoc]](https://godoc.org/github.com/briandowns/spinner)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:826** Circuit Breakers in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:759** Simple, seamless, lightweight time tracking for Git.   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [jump](https://github.com/gsamokovarov/jump) **star:747** Jump helps you navigate faster by learning your habits.   [![There was an update last week][Green]](https://github.com/gsamokovarov/jump)   [![godoc][GoDoc]](https://godoc.org/github.com/gsamokovarov/jump)
* [immortal](https://github.com/immortal/immortal) **star:628** \*nix cross-platform (OS agnostic) supervisor.   [![There was an update last week][Green]](https://github.com/immortal/immortal)   [![godoc][GoDoc]](https://godoc.org/github.com/immortal/immortal)
* [htcat](https://github.com/htcat/htcat) **star:496** Parallel and Pipelined HTTP GET Utility.   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [go-dry](https://github.com/ungerik/go-dry) **star:440** DRY (don't repeat yourself) package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [gopencils](https://github.com/bndr/gopencils) **star:426** Small and simple package to easily consume REST APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [godaemon](https://github.com/VividCortex/godaemon) **star:416** Utility to write daemons.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/godaemon)
* [circuit](https://github.com/cep21/circuit) **star:408** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/cep21/circuit)
* [request](https://github.com/mozillazg/request) **star:367** Go HTTP Requests for Humans™.   [![There was an update last week][Green]](https://github.com/mozillazg/request)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/request)
* [ergo](https://github.com/cristianoliveira/ergo) **star:342** The management of multiple local services running over different ports made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [koazee](https://github.com/wesovilabs/koazee) **star:342** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/koazee)
* [go-rate](https://github.com/beefsack/go-rate) **star:294** Timed rate limiter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beefsack/go-rate)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [gohper](https://github.com/cosiner/gohper) **star:247** Various tools/modules help for development.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cosiner/gohper)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/gohper)   [![Archived][Archived]](https://github.com/cosiner/gohper)
* [clockwork](https://github.com/jonboulle/clockwork) **star:239** A simple fake clock for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:229** Simple struct copying for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ulule/deepcopier)   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [serve](https://github.com/syntaqx/serve) **star:198** A static http server anywhere you need.   [![godoc][GoDoc]](https://godoc.org/github.com/syntaqx/serve)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:189** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [gubrak](https://github.com/novalagung/gubrak) **star:187** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [retry](https://github.com/kamilsk/retry) **star:179** The most advanced functional mechanism to perform actions repetitively until successful.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/retry)
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:171** Package for MIME type detection based on magic numbers.   [![There was an update last week][Green]](https://github.com/gabriel-vasile/mimetype)   [![godoc][GoDoc]](https://godoc.org/github.com/gabriel-vasile/mimetype)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:163** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [gotenv](https://github.com/subosito/gotenv) **star:157** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [rerun](https://github.com/ivpusic/rerun) **star:155** Recompiling and rerunning go apps when source changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/rerun)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/rerun)
* [moldova](https://github.com/StabbyCutyou/moldova) **star:148** Utility for generating random data based on an input template.   [![It hasn't been updated in the last year][Yellow]](https://github.com/StabbyCutyou/moldova)   [![godoc][GoDoc]](https://godoc.org/github.com/StabbyCutyou/moldova)
* [util](https://github.com/shomali11/util) **star:145** Collection of useful utility functions. (strings, concurrency, manipulations, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/util)
* [Death](https://github.com/vrecan/death) **star:144** Managing go application shutdown with signals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/vrecan/death)   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [robustly](https://github.com/VividCortex/robustly) **star:140** Runs functions resiliently, catching and restarting panics.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/robustly)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/robustly)
* [apm](https://github.com/topfreegames/apm) **star:131** Process manager for Golang applications with an HTTP API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/topfreegames/apm)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/apm)
* [toolbox](https://github.com/viant/toolbox) **star:113** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/toolbox)
* [chyle](https://github.com/antham/chyle) **star:113** Changelog generator using a git repository with multiple configuration possibilities.   [![There was an update last week][Green]](https://github.com/antham/chyle)   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:111** XML Sitemap generator written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:105** LiveReload server for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jaschaephraim/lrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/jaschaephraim/lrserver)
* [onecache](https://github.com/adelowo/onecache) **star:104** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).   [![godoc][GoDoc]](https://godoc.org/github.com/adelowo/onecache)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:90** Pure Go bsdiff and bspatch libraries and CLI tools.   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [pm](https://github.com/VividCortex/pm) **star:75** Process (i.e. goroutine) manager with an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/pm)
* [usql](https://github.com/knq/usql)  usql is a universal command-line interface for SQL databases.
* [UNIS](https://github.com/esemplastic/unis) **star:69** Common Architecture™ for String Utilities in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/esemplastic/unis)   [![godoc][GoDoc]](https://godoc.org/github.com/esemplastic/unis)
* [xferspdy](https://github.com/monmohan/xferspdy) **star:69** Xferspdy provides binary diff and patch library in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/monmohan/xferspdy)   [![godoc][GoDoc]](https://godoc.org/github.com/monmohan/xferspdy)
* [go-health](https://github.com/Talento90/go-health) **star:66** Health package simplifies the way you add health check to your services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Talento90/go-health)   [![godoc][GoDoc]](https://godoc.org/github.com/Talento90/go-health)
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:66** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/mssqlx)
* [repeat](https://github.com/ssgreg/repeat) **star:60** Go implementation of different backoff strategies useful for retrying operations and heartbeating.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/repeat)
* [multitick](https://github.com/VividCortex/multitick) **star:59** Multiplexor for aligned tickers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/multitick)
* [minquery](https://github.com/icza/minquery) **star:51** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/minquery)
* [handy](https://github.com/miguelpragier/handy) **star:50** Many utilities and helpers like string handlers/formatters and validators.   [![There was an update last week][Green]](https://github.com/miguelpragier/handy)   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:48** Pure Go ultra performant MIME sniffing library/utility.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zRedShift/mimemagic)   [![godoc][GoDoc]](https://godoc.org/github.com/zRedShift/mimemagic)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** Parse TODOs in your GO code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/asticode/go-astitodo)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [golog](https://github.com/mlimaloureiro/golog) **star:44** Easy and lightweight CLI tool to time track your tasks.   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [goback](https://github.com/carlescere/goback) **star:42** Go simple exponential backoff package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [goreadability](https://github.com/philipjkim/goreadability) **star:42** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [gaper](https://github.com/maxcnunes/gaper) **star:41** Builds and restarts a Go project when it crashes or some watched file changes.   [![There was an update last week][Green]](https://github.com/maxcnunes/gaper)   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:41** SeaweedFS client library with almost full features.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:41** Use x86 SIMD without writing any assembly code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/intrinsic)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/intrinsic)   [![Archived][Archived]](https://github.com/mengzhuo/intrinsic)
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:38** Universal multi-workstation clipboard that uses S3 like backend for the storage.   [![godoc][GoDoc]](https://godoc.org/github.com/jutkko/copy-pasta)
* [golarm](https://github.com/msempere/golarm) **star:36** Fire alarms with system events.   [![It hasn't been updated in the last year][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [retry](https://github.com/thedevsaddam/retry) **star:34** Simple and easy retry mechanism package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thedevsaddam/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/retry)
* [myhttp](https://github.com/inancgumus/myhttp) **star:33** Simple API to make HTTP GET requests with timeout support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/inancgumus/myhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/inancgumus/myhttp)   [![Archived][Archived]](https://github.com/inancgumus/myhttp)
* [gpath](https://github.com/tenntenn/gpath) **star:32** Library to simplify access struct fields with Go's expression in reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [realize](https://github.com/tockins/realize)  Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.
* [rclient](https://github.com/zpatrick/rclient) **star:31** Readable, flexible, simple-to-use client for REST APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rclient)
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:31** Retrying made simple and easy for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/retry-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/retry-go)
* [pgo](https://github.com/arthurkushman/pgo) **star:31** Convenient functions for PHP community.   [![There was an update last week][Green]](https://github.com/arthurkushman/pgo)   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/pgo)
* [beyond](https://github.com/wesovilabs/beyond) **star:27** The Go tool that will drive you to the AOP world!   [![There was an update last week][Green]](https://github.com/wesovilabs/beyond)   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [scan](https://github.com/blockloop/scan) **star:22** Scan golang `sql.Rows` directly to structs, slices, or primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/blockloop/scan)
* [ugo](https://github.com/alxrm/ugo) **star:22** ugo is slice toolbox with concise syntax for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alxrm/ugo)   [![godoc][GoDoc]](https://godoc.org/github.com/alxrm/ugo)
* [generate](https://github.com/go-playground/generate) **star:21** runs go generate recursively on a specified path or environment variable and can filter by regex.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [dbt](https://github.com/nikogura/dbt) **star:21** A framework for running self-updating signed binaries from a central, trusted repository.   [![There was an update last week][Green]](https://github.com/nikogura/dbt)   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:21** a small golang lib to generate placeholder images.   [![It hasn't been updated in the last year][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [evaluator](https://github.com/nullne/evaluator) **star:20** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nullne/evaluator)   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [filter](https://github.com/gookit/filter) **star:19** provide filtering, sanitizing, and conversion of Go data.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [gostrutils](https://github.com/ik5/gostrutils) **star:18** Collections of string manipulation and conversion functions.   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:15** Go library for encoding structs into Header fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** Compile-time controlled logger to make your release smaller without removing debug calls.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [rerate](https://github.com/abo/rerate) **star:14** Redis-based rate counter and rate limiter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abo/rerate)   [![godoc][GoDoc]](https://godoc.org/github.com/abo/rerate)
* [slice](https://github.com/psampaz/slice) **star:14** Type-safe functions for common Go slice operations.   [![There was an update last week][Green]](https://github.com/psampaz/slice)   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/slice)
* [okrun](https://github.com/xta/okrun) **star:14** go run error steamroller.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xta/okrun)   [![godoc][GoDoc]](https://godoc.org/github.com/xta/okrun)
* [filler](https://github.com/yaronsumel/filler) **star:14** small utility to fill structs using "fill" tag.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [structs](https://github.com/PumpkinSeed/structs) **star:13** Implement simple functions to manipulate structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PumpkinSeed/structs)   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/structs)
* [ghokin](https://github.com/antham/ghokin) **star:13** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:12** Library for executing shell commands on osx, windows and linux.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [rest-go](https://github.com/edermanoel94/rest-go) **star:12** A package that provide many helpful methods for working with rest api.   [![godoc][GoDoc]](https://godoc.org/github.com/edermanoel94/rest-go)
* [ctxutil](https://github.com/posener/ctxutil) **star:11** A collection of utility functions for contexts.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:11** A MIME type sniffer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/mimesniffer)
* [retry](https://github.com/shafreeck/retry) **star:10** A pretty simple library to ensure your work to be done.   [![godoc][GoDoc]](https://godoc.org/github.com/shafreeck/retry)
* [command](https://github.com/txgruppi/command) **star:9** Command pattern for Go with thread safe serial and parallel dispatcher.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [backscanner](https://github.com/icza/backscanner) **star:9** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![It hasn't been updated in the last year][Yellow]](https://github.com/icza/backscanner)   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [limiters](https://github.com/mennanov/limiters) **star:7** Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.   [![godoc][GoDoc]](https://godoc.org/github.com/mennanov/limiters)
* [slicer](https://github.com/leaanthony/slicer) **star:6** Makes working with slices easier.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/slicer)
* [shutdown](https://github.com/ztrue/shutdown) **star:5** App shutdown hooks for `os.Signal` handling.   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/shutdown)
* [sslice](https://github.com/yaa110/sslice) **star:5** Create a slice which is always sorted.   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/sslice)
* [blank](https://github.com/Henry-Sarabia/blank) **star:4** Verify or remove blanks and whitespace from strings.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [retry](https://github.com/percolate/retry) **star:4** A simple but highly configurable retry package for Go.
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.   [![godoc][GoDoc]](https://godoc.org/github.com/chrispassas/silk)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:3** Go package for working with Problem Details.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:3** Slice conversion between primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
* [olaf](https://github.com/btnguyen2k/olaf) **star:1** Twitter Snowflake implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/btnguyen2k/olaf)

## UUID

*Libraries for working with UUIDs.*

* [ulid](https://github.com/oklog/ulid) **star:1762** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   [![godoc][GoDoc]](https://godoc.org/github.com/oklog/ulid)
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands.
* [uuid](https://github.com/gofrs/uuid) **star:619** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   [![godoc][GoDoc]](https://godoc.org/github.com/gofrs/uuid)
* [wuid](https://github.com/edwingeng/wuid) **star:317** An extremely fast unique number generator, 10-135 times faster than UUID.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/wuid)
* [goid](https://github.com/jakehl/goid) **star:25** Generate and Parse RFC4122 compliant V4 UUIDs.   [![godoc][GoDoc]](https://godoc.org/github.com/jakehl/goid)
* [sno](https://github.com/muyo/sno) **star:20** Compact, sortable and fast unique IDs with embedded metadata.   [![godoc][GoDoc]](https://godoc.org/github.com/muyo/sno)
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:13** A tiny and efficient Go unique string ID generator.   [![godoc][GoDoc]](https://godoc.org/github.com/aidarkhanov/nanoid)
* [uuid](https://github.com/agext/uuid) **star:10** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/uuid)

## Validation

*Libraries for validation.*

* [validator](https://github.com/go-playground/validator) **star:4393** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   [![star > 2000][Awesome]](https://github.com/go-playground/validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/validator)
* [govalidator](https://github.com/asaskevich/govalidator) **star:3880** Validators and sanitizers for strings, numerics, slices and structs.   [![star > 2000][Awesome]](https://github.com/asaskevich/govalidator)   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/govalidator)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1179** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   [![There was an update last week][Green]](https://github.com/go-ozzo/ozzo-validation)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:796** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/govalidator)
* [validate](https://github.com/gookit/validate) **star:165** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/validate)   [![Contains Chinese documents][CN]](https://github.com/gookit/validate)
* [checkdigit](https://github.com/osamingo/checkdigit) **star:48** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/checkdigit)
* [jio](https://github.com/faceair/jio) **star:29** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   [![godoc][GoDoc]](https://godoc.org/github.com/faceair/jio)   [![Contains Chinese documents][CN]](https://github.com/faceair/jio)
* [validate](https://github.com/gobuffalo/validate) **star:26** This package provides a framework for writing validations for Go applications.   [![There was an update last week][Green]](https://github.com/gobuffalo/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/validate)
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:24** A norms and conventions validator for Terraform.   [![godoc][GoDoc]](https://godoc.org/github.com/thazelart/terraform-validator)

## Version Control

*Libraries for version control.*

* [go-git](https://github.com/src-d/go-git) **star:4706** highly extensible Git implementation in pure Go.   [![star > 2000][Awesome]](https://github.com/src-d/go-git)   [![There was an update last week][Green]](https://github.com/src-d/go-git)   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/go-git)
* [git2go](https://github.com/libgit2/git2go) **star:1412** Go bindings for libgit2.   [![godoc][GoDoc]](https://godoc.org/github.com/libgit2/git2go)
* [hercules](https://github.com/src-d/hercules) **star:749** gaining advanced insights from Git repository history.   [![There was an update last week][Green]](https://github.com/src-d/hercules)   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/hercules)
* [gh](https://github.com/rjeczalik/gh) **star:71** Scriptable server and net/http middleware for GitHub Webhooks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjeczalik/gh)   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/gh)
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:71** manipulate and inspect VCS repositories in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/sourcegraph/go-vcs)
* [hgo](https://github.com/beyang/hgo) **star:12** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beyang/hgo)   [![godoc][GoDoc]](https://godoc.org/github.com/beyang/hgo)

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) **star:942** Comphrensive Go bindings for FFmpeg.   [![godoc][GoDoc]](https://godoc.org/github.com/giorgisio/goav)
* [m3u8](https://github.com/grafov/m3u8) **star:626** Parser and generator library of M3U8 playlists for Apple HLS.   [![godoc][GoDoc]](https://godoc.org/github.com/grafov/m3u8)
* [gmf](https://github.com/3d0c/gmf) **star:570** Go bindings for FFmpeg av\* libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/3d0c/gmf)
* [go-astits](https://github.com/asticode/go-astits) **star:277** Parse and demux MPEG Transport Streams (.ts) natively in GO.   [![There was an update last week][Green]](https://github.com/asticode/go-astits)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astits)
* [go-astisub](https://github.com/asticode/go-astisub) **star:210** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   [![There was an update last week][Green]](https://github.com/asticode/go-astisub)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astisub)
* [gst](https://github.com/ziutek/gst) **star:154** Go bindings for GStreamer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/gst)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/gst)
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:77** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/libvlc-go)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:43** Parser and generator library for Apple m3u8 playlists.   [![godoc][GoDoc]](https://godoc.org/github.com/quangngotan95/go-m3u8)
* [v4l](https://github.com/korandiz/v4l) **star:34** Video capture library for Linux, written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/korandiz/v4l)   [![godoc][GoDoc]](https://godoc.org/github.com/korandiz/v4l)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:11** Subtitle format support for go. Supports .srt, .ttml, and .ass.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wargarblgarbl/libgosubs)   [![godoc][GoDoc]](https://godoc.org/github.com/wargarblgarbl/libgosubs)

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) **star:33713** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   [![star > 2000][Awesome]](https://github.com/gin-gonic/gin)   [![There was an update last week][Green]](https://github.com/gin-gonic/gin)   [![godoc][GoDoc]](https://godoc.org/github.com/gin-gonic/gin)
* [Beego](https://github.com/astaxie/beego) **star:22734** beego is an open-source, high-performance web framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/astaxie/beego)   [![There was an update last week][Green]](https://github.com/astaxie/beego)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/beego)   [![Contains Chinese documents][CN]](https://github.com/astaxie/beego)
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo) **star:15917** High performance, minimalist Go web framework.   [![star > 2000][Awesome]](https://github.com/labstack/echo)   [![There was an update last week][Green]](https://github.com/labstack/echo)   [![godoc][GoDoc]](https://godoc.org/github.com/labstack/echo)
* [Revel](https://github.com/revel/revel) **star:11452** High-productivity web framework for the Go language.   [![star > 2000][Awesome]](https://github.com/revel/revel)   [![godoc][GoDoc]](https://godoc.org/github.com/revel/revel)
* [Goa](https://github.com/goadesign/goa) **star:3650** Goa provides a holistic approach for developing remote APIs and microservices in Go.   [![star > 2000][Awesome]](https://github.com/goadesign/goa)   [![There was an update last week][Green]](https://github.com/goadesign/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goadesign/goa)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3369** Quick and easy way to setup a RESTful JSON API.   [![star > 2000][Awesome]](https://github.com/ant0ine/go-json-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ant0ine/go-json-rest)
* [Gizmo](https://github.com/NYTimes/gizmo) **star:2976** Microservice toolkit used by the New York Times.   [![star > 2000][Awesome]](https://github.com/NYTimes/gizmo)   [![godoc][GoDoc]](https://godoc.org/github.com/NYTimes/gizmo)
* [Macaron](https://github.com/go-macaron/macaron) **star:2897** Macaron is a high productive and modular design web framework in Go.   [![star > 2000][Awesome]](https://github.com/go-macaron/macaron)   [![godoc][GoDoc]](https://godoc.org/github.com/go-macaron/macaron)
* [utron](https://github.com/gernest/utron) **star:2148** Lightweight MVC framework for Go(Golang).   [![star > 2000][Awesome]](https://github.com/gernest/utron)   [![It hasn't been updated in the last year][Yellow]](https://github.com/gernest/utron)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/utron)
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:996** Go framework for building JSON web services inspired by Dropwizard.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rcrowley/go-tigertonic)   [![godoc][GoDoc]](https://godoc.org/github.com/rcrowley/go-tigertonic)
* [tango](https://github.com/lunny/tango) **star:828** Micro & pluggable web framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/tango)   [![Contains Chinese documents][CN]](https://github.com/lunny/tango)
* [gongular](https://github.com/mustafaakin/gongular) **star:421** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaakin/gongular)
* [neo](https://github.com/ivpusic/neo) **star:399** Neo is minimal and fast Go Web Framework with extremely simple API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/neo)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/neo)
* [Air](https://github.com/aofei/air) **star:350** An ideally refined web framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/air)
* [mango](https://github.com/paulbellamy/mango) **star:340** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paulbellamy/mango)   [![godoc][GoDoc]](https://godoc.org/github.com/paulbellamy/mango)
* [Gondola](https://github.com/rainycape/gondola) **star:315** The web framework for writing faster sites, faster.   [![godoc][GoDoc]](https://godoc.org/github.com/rainycape/gondola)
* [Golf](https://github.com/dinever/golf) **star:238** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dinever/golf)   [![godoc][GoDoc]](https://godoc.org/github.com/dinever/golf)
* [Aero](https://github.com/aerogo/aero) **star:208** High-performance web framework for Go, reaches top scores in Lighthouse.   [![There was an update last week][Green]](https://github.com/aerogo/aero)   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/aero)
* [go-rest](https://github.com/ungerik/go-rest) **star:116** Small and evil REST framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-rest)
* [hiboot](https://github.com/hidevopsio/hiboot) **star:105** hiboot is a high performance web application framework with auto configuration and dependency injection support.   [![godoc][GoDoc]](https://godoc.org/github.com/hidevopsio/hiboot)   [![Contains Chinese documents][CN]](https://github.com/hidevopsio/hiboot)
* [WebGo](https://github.com/bnkamalesh/webgo) **star:83** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/webgo)
* [Golax](https://github.com/fulldump/golax) **star:71** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fulldump/golax)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/golax)
* [Flamingo](https://github.com/i-love-flamingo/flamingo) **star:70** Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.   [![There was an update last week][Green]](https://github.com/i-love-flamingo/flamingo)   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo)
* [Microservice](https://github.com/claygod/microservice) **star:68** The framework for the creation of microservices, written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/microservice)
* [uAdmin](https://github.com/uadmin/uadmin) **star:66** Fully featured web framework for Golang, inspired by Django.   [![godoc][GoDoc]](https://godoc.org/github.com/uadmin/uadmin)
* [YARF](https://github.com/yarf-framework/yarf) **star:52** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.   [![godoc][GoDoc]](https://godoc.org/github.com/yarf-framework/yarf)
* [Fireball](https://github.com/zpatrick/fireball) **star:49** More "natural" feeling web framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/fireball)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/fireball)
* [vox](https://github.com/aisk/vox) **star:45** A golang web framework for humans, inspired by Koa heavily.   [![godoc][GoDoc]](https://godoc.org/github.com/aisk/vox)
* [patron](https://github.com/beatlabs/patron) **star:43** Patron is a microservice framework following best cloud practices with a focus on productivity.   [![There was an update last week][Green]](https://github.com/beatlabs/patron)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/patron)
* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Resoursea](https://github.com/resoursea/api) **star:29** REST framework for quickly writing resource based services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/resoursea/api)   [![godoc][GoDoc]](https://godoc.org/github.com/resoursea/api)
* [rex](https://github.com/goanywhere/rex) **star:27** Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goanywhere/rex)   [![godoc][GoDoc]](https://godoc.org/github.com/goanywhere/rex)
* [goa](https://github.com/goa-go/goa) **star:26** goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware.   [![There was an update last week][Green]](https://github.com/goa-go/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goa-go/goa)
* [rux](https://github.com/gookit/rux) **star:18** Simple and fast web framework for build golang HTTP applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/rux)   [![Contains Chinese documents][CN]](https://github.com/gookit/rux)
* [Ginrpc](https://github.com/xxjwxc/ginrpc) **star:15** Gin parameter automatic binding tool,gin rpc tools.   [![There was an update last week][Green]](https://github.com/xxjwxc/ginrpc)   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/ginrpc)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/ginrpc)
* [Banjo](https://github.com/nsheremet/banjo) **star:11** Very simple and fast web framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nsheremet/banjo)   [![godoc][GoDoc]](https://godoc.org/github.com/nsheremet/banjo)

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) **star:1411** Rate limit HTTP request handler.   [![godoc][GoDoc]](https://godoc.org/github.com/didip/tollbooth)
* [CORS](https://github.com/rs/cors) **star:1318** Easily add CORS capabilities to your API.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/cors)
* [Limiter](https://github.com/ulule/limiter) **star:851** Dead simple rate limit middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/limiter)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:766** Add/parse Server-Timing header.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mitchellh/go-server-timing)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/go-server-timing)
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:98** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/ln-paywall)
* [XFF](https://github.com/sebest/xff) **star:72** Handle `X-Forwarded-For` header and friends.   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/xff)
* [formjson](https://github.com/rs/formjson) **star:34** Transparently handle JSON input as a standard form POST.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/formjson)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/formjson)
* [client-timing](https://github.com/posener/client-timing) **star:16** An HTTP client for Server-Timing header.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/client-timing)

#### Libraries for creating HTTP middlewares

* [negroni](https://github.com/urfave/negroni) **star:6455** Idiomatic HTTP middleware for Golang.   [![star > 2000][Awesome]](https://github.com/urfave/negroni)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/negroni)   [![Contains Chinese documents][CN]](https://github.com/urfave/negroni)
* [alice](https://github.com/justinas/alice) **star:1908** Painless middleware chaining for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/alice)
* [render](https://github.com/unrolled/render) **star:1312** Go package for easily rendering JSON, XML, and HTML template responses.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/render)
* [stats](https://github.com/thoas/stats) **star:547** Go middleware that stores various information about your web application.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/stats)
* [interpose](https://github.com/carbocation/interpose) **star:290** Minimalist net/http middleware for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carbocation/interpose)   [![godoc][GoDoc]](https://godoc.org/github.com/carbocation/interpose)
* [muxchain](https://github.com/stephens2424/muxchain) **star:209** Lightweight middleware for net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/stephens2424/muxchain)
* [renderer](https://github.com/thedevsaddam/renderer) **star:177** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/renderer)
* [rye](https://github.com/InVisionApp/rye) **star:92** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/rye)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/rye)
* [gores](https://github.com/alioygur/gores) **star:88** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alioygur/gores)   [![godoc][GoDoc]](https://godoc.org/github.com/alioygur/gores)
* [chain](https://github.com/codemodus/chain) **star:64** Handler wrapper chaining with scoped data (net/context-based "middleware").   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/chain)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/chain)
* [go-wrap](https://github.com/go-on/wrap) **star:59** Small middlewares package for net/http.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-on/wrap)   [![godoc][GoDoc]](https://godoc.org/github.com/go-on/wrap)
* [catena](https://github.com/codemodus/catena) **star:7** http.Handler wrapper catenation (same API as "chain").   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/catena)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/catena)

### Routers

* [mux](https://github.com/gorilla/mux) **star:10688** Powerful URL router and dispatcher for golang.   [![star > 2000][Awesome]](https://github.com/gorilla/mux)   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/mux)
* [httprouter](https://github.com/julienschmidt/httprouter) **star:10435** High performance router. Use this and the standard http handlers to form a very high performance web framework.   [![star > 2000][Awesome]](https://github.com/julienschmidt/httprouter)   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/httprouter)
* [chi](https://github.com/go-chi/chi) **star:6718** Small, fast and expressive HTTP router built on net/context.   [![star > 2000][Awesome]](https://github.com/go-chi/chi)   [![There was an update last week][Green]](https://github.com/go-chi/chi)   [![godoc][GoDoc]](https://godoc.org/github.com/go-chi/chi)
* [gocraft/web](https://github.com/gocraft/web) **star:1401** Mux and middleware package in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gocraft/web)
* [Bone](https://github.com/go-zoo/bone) **star:1243** Lightning Fast HTTP Multiplexer.   [![godoc][GoDoc]](https://godoc.org/github.com/go-zoo/bone)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:798** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   [![godoc][GoDoc]](https://godoc.org/github.com/buaazp/fasthttprouter)
* [Goji](https://github.com/goji/goji) **star:782** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/goji/goji)
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:467** A simple and fast HTTP router for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gorouter)
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:400** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   [![godoc][GoDoc]](https://godoc.org/github.com/dimfeld/httptreemux)
* [lars](https://github.com/go-playground/lars) **star:377** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/lars)
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:366** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   [![There was an update last week][Green]](https://github.com/go-ozzo/ozzo-routing)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
* [Siesta](https://github.com/VividCortex/siesta) **star:352** Composable framework to write middleware and handlers.   [![There was an update last week][Green]](https://github.com/VividCortex/siesta)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/siesta)
* [vestigo](https://github.com/husobee/vestigo) **star:254** Performant, stand-alone, HTTP compliant URL Router for go web applications.   [![godoc][GoDoc]](https://godoc.org/github.com/husobee/vestigo)
* [gowww/router](https://github.com/gowww/router) **star:158** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gowww/router)   [![godoc][GoDoc]](https://godoc.org/github.com/gowww/router)
* [alien](https://github.com/gernest/alien) **star:108** Lightweight and fast http router from outer space.   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/alien)
* [violetear](https://github.com/nbari/violetear) **star:95** Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/nbari/violetear)
* [Bxog](https://github.com/claygod/Bxog) **star:92** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/Bxog)
* [pure](https://github.com/go-playground/pure) **star:90** Is a lightweight HTTP router that sticks to the std "net/http" implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pure)
* [xmux](https://github.com/rs/xmux) **star:85** High performance muxer based on `httprouter` with `net/context` support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/xmux)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xmux)
* [GoRouter](https://github.com/vardius/gorouter) **star:52** GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gorouter)
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:40** A simple Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/bellt)
* [FastRouter](https://github.com/razonyang/fastrouter) **star:18** a fast, flexible HTTP router written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/razonyang/fastrouter)   [![godoc][GoDoc]](https://godoc.org/github.com/razonyang/fastrouter)
* [goroute](https://github.com/goroute/route) **star:6** Simple yet powerful HTTP request multiplexer.   [![godoc][GoDoc]](https://godoc.org/github.com/goroute/route)

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:585** Win32 OLE implementation for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ole/go-ole)
* [d3d9](https://github.com/gonutz/d3d9) **star:94** Go bindings for Direct3D9.   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/d3d9)
* [gosddl](https://github.com/MonaxGT/gosddl) **star:1** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gosddl)

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) **star:307** Generate a Go struct from XML.   [![godoc][GoDoc]](https://godoc.org/github.com/miku/zek)
* [xpath](https://github.com/antchfx/xpath) **star:200** XPath package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xpath)
* [xquery](https://github.com/antchfx/xquery) **star:151** XQuery lets you extract data from HTML/XML documents using XPath expression.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/xquery)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xquery)   [![Archived][Archived]](https://github.com/antchfx/xquery)
* [xml2map](https://github.com/sbabiv/xml2map) **star:18** XML to MAP converter written Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/xml2map)
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** Simple command line XML comparer that generates diffs of folders, files and tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xml-comp/xml-comp)   [![godoc][GoDoc]](https://godoc.org/github.com/xml-comp/xml-comp)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:7** Procedural XML generation API based on libxml2's xmlwriter module.   [![godoc][GoDoc]](https://godoc.org/github.com/shabbyrobe/xmlwriter)

# Tools

*Go software and plugins.*

## Code Analysis

* [GoLint](https://github.com/golang/lint) **star:3298** Golint is a linter for Go source code.   [![star > 2000][Awesome]](https://github.com/golang/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/lint)
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [errcheck](https://github.com/kisielk/errcheck) **star:1371** Errcheck is a program for checking for unchecked errors in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/kisielk/errcheck)
* [gcvis](https://github.com/davecheney/gcvis) **star:940** Visualise Go program GC trace data in real time.   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/gcvis)
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  tarp finds functions and methods without direct unit tests in Go source code.
* [php-parser](https://github.com/z7zmey/php-parser) **star:681** A Parser for PHP written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/z7zmey/php-parser)
* [go-critic](https://github.com/go-critic/go-critic) **star:637** source code linter that brings checks that are currently not implemented in other linters.   [![There was an update last week][Green]](https://github.com/go-critic/go-critic)   [![godoc][GoDoc]](https://godoc.org/github.com/go-critic/go-critic)
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:408** Web based Golang AST visualizer.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:296** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   [![It hasn't been updated in the last year][Yellow]](https://github.com/roblaszczak/go-cleanarch)   [![godoc][GoDoc]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
* [unconvert](https://github.com/mdempsky/unconvert) **star:272** Remove unnecessary type conversions from Go source.   [![godoc][GoDoc]](https://godoc.org/github.com/mdempsky/unconvert)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [gostatus](https://github.com/shurcooL/gostatus) **star:242** Command line tool, shows the status of repositories that contain Go packages.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/gostatus)
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:239** An easy way to find outdated dependencies of your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/go-mod-outdated)
* [dupl](https://github.com/mibk/dupl) **star:182** Tool for code clone detection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mibk/dupl)   [![godoc][GoDoc]](https://godoc.org/github.com/mibk/dupl)
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:169** Checks recent changes to a Go project for backwards incompatible changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bradleyfalzon/apicompat)   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyfalzon/apicompat)
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:96** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/checkstyle)
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) **star:66** Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.   [![godoc][GoDoc]](https://godoc.org/github.com/jfeliu007/goplantuml)
* [lint](https://github.com/surullabs/lint) **star:64** Run linters as part of go test.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surullabs/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/surullabs/lint)
* [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mccoyst/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/validate)
* [go-outdated](https://github.com/firstrow/go-outdated) **star:45** Console application that displays outdated packages.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/go-outdated)   [![Archived][Archived]](https://github.com/firstrow/go-outdated)

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) **star:11337** Go development plugin for Vim.   [![star > 2000][Awesome]](https://github.com/fatih/vim-go)   [![There was an update last week][Green]](https://github.com/fatih/vim-go)
* [vscode-go](https://github.com/Microsoft/vscode-go) **star:5483** Extension for Visual Studio Code (VS Code) which provides support for the Go language.   [![star > 2000][Awesome]](https://github.com/Microsoft/vscode-go)   [![There was an update last week][Green]](https://github.com/Microsoft/vscode-go)
* [gocode](https://github.com/nsf/gocode) **star:4800** Autocompletion daemon for the Go programming language.   [![star > 2000][Awesome]](https://github.com/nsf/gocode)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/gocode)
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3286** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.   [![star > 2000][Awesome]](https://github.com/DisposaBoy/GoSublime)   [![godoc][GoDoc]](https://godoc.org/github.com/DisposaBoy/GoSublime)
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1493** Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.   [![There was an update last week][Green]](https://github.com/joefitzgerald/go-plus)
* [go-mode](https://github.com/dominikh/go-mode.el) **star:1006** Go mode for GNU/Emacs.
* [Watch](https://github.com/eaburns/Watch) **star:170** Runs a command in an acme win on file changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eaburns/Watch)   [![godoc][GoDoc]](https://godoc.org/github.com/eaburns/Watch)
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:82** Vim plugin to highlight syntax errors on save.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/vim-compiler-go)
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:33** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:17** Vim plugin for generating Go tests based on the function's or method's signature.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/gounit-vim)
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:13** Go language support for the Theia IDE.

## Go Generate Tools

* [gotests](https://github.com/cweill/gotests) **star:2427** Generate Go tests from your source code.   [![star > 2000][Awesome]](https://github.com/cweill/gotests)   [![godoc][GoDoc]](https://godoc.org/github.com/cweill/gotests)
* [genny](https://github.com/cheekybits/genny) **star:1071** Elegant generics for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cheekybits/genny)
* [re2dfa](https://github.com/opennota/re2dfa) **star:174** Transform regular expressions into finite state machines and output Go source code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/opennota/re2dfa)   [![godoc][GoDoc]](https://godoc.org/github.com/opennota/re2dfa)
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.
* [gocontracts](https://github.com/Parquery/gocontracts) **star:55** brings design-by-contract to Go by synchronizing the code with the documentation.   [![godoc][GoDoc]](https://godoc.org/github.com/Parquery/gocontracts)
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [gounit](https://github.com/hexdigest/gounit) **star:34** Generate Go tests using your own templates.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/gounit)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gounit)
* [generic](https://github.com/usk81/generic) **star:32** flexible data type for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/usk81/generic)
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:29** Generate Haskell inspired functions for your slices.   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/hasgo)

## Go Tools

* [go-swagger](https://github.com/go-swagger/go-swagger) **star:4546** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.   [![star > 2000][Awesome]](https://github.com/go-swagger/go-swagger)   [![godoc][GoDoc]](https://godoc.org/github.com/go-swagger/go-swagger)
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:2229** Visualize call graph of your Go program using dot format.   [![star > 2000][Awesome]](https://github.com/TrueFurby/go-callvis)   [![godoc][GoDoc]](https://godoc.org/github.com/TrueFurby/go-callvis)
* [depth](https://github.com/KyleBanks/depth) **star:418** Visualize dependency trees of any package by analyzing imports.   [![godoc][GoDoc]](https://godoc.org/github.com/KyleBanks/depth)
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [richgo](https://github.com/kyoh86/richgo) **star:417** Enrich `go test` outputs with text decorations.   [![There was an update last week][Green]](https://github.com/kyoh86/richgo)   [![godoc][GoDoc]](https://godoc.org/github.com/kyoh86/richgo)
* [rts](https://github.com/galeone/rts) **star:190** RTS: response to struct. Generates Go structs from server responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/galeone/rts)   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/rts)
* [godbg](https://github.com/tylerwince/godbg) **star:162** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.   [![godoc][GoDoc]](https://godoc.org/github.com/tylerwince/godbg)
* [colorgo](https://github.com/songgao/colorgo) **star:100** Wrapper around `go` command for colorized `go build` output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/colorgo)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/colorgo)
* [gothanks](https://github.com/psampaz/gothanks) **star:78** GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/gothanks)
* [OctoLinker](https://github.com/OctoLinker/browser-extension)  Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:36** Bash completion for go and wgo.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/go-pkg-complete)
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:17** A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.
* [go-james](https://github.com/pieterclaerhout/go-james) **star:2** Go project skeleton creator, builds and tests your projects without the manual setup.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-james)

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:60916** Container Cluster Manager from Google.   [![star > 2000][Awesome]](https://github.com/kubernetes/kubernetes)   [![There was an update last week][Green]](https://github.com/kubernetes/kubernetes)   [![godoc][GoDoc]](https://godoc.org/github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby) **star:55774** Collaborative project for the container ecosystem to assemble container-based systems.   [![star > 2000][Awesome]](https://github.com/moby/moby)   [![There was an update last week][Green]](https://github.com/moby/moby)   [![godoc][GoDoc]](https://godoc.org/github.com/moby/moby)
* [traefik](https://github.com/containous/traefik) **star:26244** Reverse proxy and load balancer with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/containous/traefik)   [![There was an update last week][Green]](https://github.com/containous/traefik)   [![godoc][GoDoc]](https://godoc.org/github.com/containous/traefik)
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [Gitea](https://github.com/go-gitea/gitea) **star:17375** Fork of Gogs, entirely community driven.   [![star > 2000][Awesome]](https://github.com/go-gitea/gitea)   [![There was an update last week][Green]](https://github.com/go-gitea/gitea)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gitea/gitea)   [![Contains Chinese documents][CN]](https://github.com/go-gitea/gitea)
* [Vegeta](https://github.com/tsenart/vegeta) **star:13371** HTTP load testing tool and library. It's over 9000!   [![star > 2000][Awesome]](https://github.com/tsenart/vegeta)   [![There was an update last week][Green]](https://github.com/tsenart/vegeta)   [![godoc][GoDoc]](https://godoc.org/github.com/tsenart/vegeta)
* [Hey](https://github.com/rakyll/hey) **star:7190** Hey is a tiny program that sends some load to a web application.   [![star > 2000][Awesome]](https://github.com/rakyll/hey)   [![There was an update last week][Green]](https://github.com/rakyll/hey)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/hey)
* [GVM](https://github.com/moovweb/gvm) **star:4812** GVM provides an interface to manage Go versions.   [![star > 2000][Awesome]](https://github.com/moovweb/gvm)
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [webhook](https://github.com/adnanh/webhook) **star:4632** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   [![star > 2000][Awesome]](https://github.com/adnanh/webhook)   [![There was an update last week][Green]](https://github.com/adnanh/webhook)   [![godoc][GoDoc]](https://godoc.org/github.com/adnanh/webhook)
* [gaia](https://github.com/gaia-pipeline/gaia) **star:3899** Build powerful pipelines in any programming language.   [![star > 2000][Awesome]](https://github.com/gaia-pipeline/gaia)   [![There was an update last week][Green]](https://github.com/gaia-pipeline/gaia)   [![godoc][GoDoc]](https://godoc.org/github.com/gaia-pipeline/gaia)
* [gox](https://github.com/mitchellh/gox) **star:3542** Dead simple, no frills Go cross compile tool.   [![star > 2000][Awesome]](https://github.com/mitchellh/gox)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/gox)
* [bosun](https://github.com/bosun-monitor/bosun) **star:2902** Time Series Alerting Framework.   [![star > 2000][Awesome]](https://github.com/bosun-monitor/bosun)   [![There was an update last week][Green]](https://github.com/bosun-monitor/bosun)   [![godoc][GoDoc]](https://godoc.org/github.com/bosun-monitor/bosun)
* [bombardier](https://github.com/codesenberg/bombardier) **star:1890** Fast cross-platform HTTP benchmarking tool.   [![godoc][GoDoc]](https://godoc.org/github.com/codesenberg/bombardier)
* [goxc](https://github.com/laher/goxc) **star:1634** build tool for Go, with a focus on cross-compiling and packaging.   [![godoc][GoDoc]](https://godoc.org/github.com/laher/goxc)
* [fac](https://github.com/mkchoi212/fac) **star:1629** Command-line user interface to fix git merge conflicts.   [![godoc][GoDoc]](https://godoc.org/github.com/mkchoi212/fac)
* [kala](https://github.com/ajvb/kala) **star:1380** Simplistic, modern, and performant job scheduler.   [![There was an update last week][Green]](https://github.com/ajvb/kala)   [![godoc][GoDoc]](https://godoc.org/github.com/ajvb/kala)
* [StatusOK](https://github.com/sanathp/statusok) **star:1238** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.   [![godoc][GoDoc]](https://godoc.org/github.com/sanathp/statusok)
* [script](https://github.com/bitfield/script) **star:1233** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   [![There was an update last week][Green]](https://github.com/bitfield/script)   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/script)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:1024** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.   [![godoc][GoDoc]](https://godoc.org/github.com/rlmcpherson/s3gof3r)
* [Pomerium](https://github.com/pomerium/pomerium) **star:706** Pomerium is an identity-aware access proxy.   [![There was an update last week][Green]](https://github.com/pomerium/pomerium)   [![godoc][GoDoc]](https://godoc.org/github.com/pomerium/pomerium)
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:696** Enable your Go applications to self update.   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/go-selfupdate)
* [skm](https://github.com/TimothyYe/skm) **star:572** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   [![godoc][GoDoc]](https://godoc.org/github.com/TimothyYe/skm)
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:553** Manage BareMetal Servers from Command Line (as easily as with Docker).   [![godoc][GoDoc]](https://godoc.org/github.com/scaleway/scaleway-cli)
* [aurora](https://github.com/xuri/aurora) **star:433** Cross-platform web-based Beanstalkd queue server console.
* [gonative](https://github.com/inconshreveable/gonative) **star:313** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/inconshreveable/gonative)   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/gonative)
* [govvv](https://github.com/ahmetalpbalkan/govvv)  “go build” wrapper to easily add version information into Go binaries.
* [Mora](https://github.com/emicklei/mora) **star:268** REST server for accessing MongoDB documents and meta data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/emicklei/mora)   [![godoc][GoDoc]](https://godoc.org/github.com/emicklei/mora)
* [lstags](https://github.com/ivanilves/lstags) **star:240** Tool and API to sync Docker images across different registries.   [![There was an update last week][Green]](https://github.com/ivanilves/lstags)   [![godoc][GoDoc]](https://godoc.org/github.com/ivanilves/lstags)
* [dogo](https://github.com/liudng/dogo) **star:221** Monitoring changes in the source file and automatically compile and run (restart).   [![godoc][GoDoc]](https://godoc.org/github.com/liudng/dogo)   [![Contains Chinese documents][CN]](https://github.com/liudng/dogo)
* [godbg](https://github.com/sirnewton01/godbg) **star:219** Web-based gdb front-end application.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sirnewton01/godbg)
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [Pewpew](https://github.com/bengadbois/pewpew) **star:216** Flexible HTTP command line stress tester.   [![godoc][GoDoc]](https://godoc.org/github.com/bengadbois/pewpew)
* [manssh](https://github.com/xwjdsh/manssh) **star:208** manssh is a command line tool for managing your ssh alias config easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xwjdsh/manssh)   [![godoc][GoDoc]](https://godoc.org/github.com/xwjdsh/manssh)
* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [gobrew](https://github.com/cryptojuice/gobrew) **star:176** gobrew lets you easily switch between multiple versions of go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cryptojuice/gobrew)
* [Blast](https://github.com/dave/blast) **star:171** A simple tool for API load testing and batch jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dave/blast)   [![godoc][GoDoc]](https://godoc.org/github.com/dave/blast)
* [ostent](https://github.com/ostrost/ostent) **star:164** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ostrost/ostent)   [![godoc][GoDoc]](https://godoc.org/github.com/ostrost/ostent)
* [Packer](https://github.com/mitchellh/packer)  Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
* [grapes](https://github.com/yaronsumel/grapes) **star:141** Lightweight tool designed to distribute commands over ssh with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/grapes)
* [uTask](https://github.com/ovh/utask) **star:115** Automation engine that models and executes business processes declared in yaml.   [![There was an update last week][Green]](https://github.com/ovh/utask)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/utask)
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:114** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/easyssh-proxy)
* [jcli](https://github.com/jenkins-zh/jenkins-cli) **star:101** Jenkins CLI allows you manage your Jenkins as an easy way.   [![There was an update last week][Green]](https://github.com/jenkins-zh/jenkins-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/jenkins-zh/jenkins-cli)   [![Contains Chinese documents][CN]](https://github.com/jenkins-zh/jenkins-cli)
* [kcli](https://github.com/cswank/kcli) **star:97** Command line tool for inspecting kafka topics/partitions/messages.   [![godoc][GoDoc]](https://godoc.org/github.com/cswank/kcli)
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) **star:86** Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.   [![There was an update last week][Green]](https://github.com/dikhan/terraform-provider-openapi)   [![godoc][GoDoc]](https://godoc.org/github.com/dikhan/terraform-provider-openapi)
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:81** Cli tool to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm-cli)
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:73** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   [![godoc][GoDoc]](https://godoc.org/github.com/go-furnace/go-furnace)
* [drone-scp](https://github.com/appleboy/drone-scp) **star:65** Copy files and artifacts via SSH using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-scp)
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) **star:56** A go library and an executable that produces valid Dockerfiles using various input channels.   [![godoc][GoDoc]](https://godoc.org/github.com/ozankasikci/dockerfile-generator)
* [Dropship](https://github.com/chrismckenzie/dropship) **star:46** Tool for deploying code via cdn.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrismckenzie/dropship)   [![godoc][GoDoc]](https://godoc.org/github.com/chrismckenzie/dropship)
* [Rodent](https://github.com/alouche/rodent) **star:30** Rodent helps you manage Go versions, projects and track dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alouche/rodent)
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:26** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-jenkins)
* [awsenv](https://github.com/soniah/awsenv) **star:22** Small binary that loads Amazon (AWS) environment variables for a profile.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/awsenv)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/awsenv)
* [DepCharge](https://github.com/centerorbit/depcharge) **star:9** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![godoc][GoDoc]](https://godoc.org/github.com/centerorbit/depcharge)
* [lwc](https://github.com/timdp/lwc) **star:9** A live-updating version of the UNIX wc command.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timdp/lwc)   [![godoc][GoDoc]](https://godoc.org/github.com/timdp/lwc)
* [sg](https://github.com/ChristopherRabotin/sg) **star:5** Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/sg)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/sg)

### Other Software

* [restic](https://github.com/restic/restic) **star:8893** De-duplicating backup program.   [![star > 2000][Awesome]](https://github.com/restic/restic)   [![There was an update last week][Green]](https://github.com/restic/restic)   [![godoc][GoDoc]](https://godoc.org/github.com/restic/restic)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:8847** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   [![star > 2000][Awesome]](https://github.com/chrislusf/seaweedfs)   [![There was an update last week][Green]](https://github.com/chrislusf/seaweedfs)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/seaweedfs)
* [confd](https://github.com/kelseyhightower/confd) **star:6706** Manage local application configuration files using templates and data from etcd or consul.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/confd)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/confd)
* [Comcast](https://github.com/tylertreat/Comcast) **star:6306** Simulate bad network connections.   [![star > 2000][Awesome]](https://github.com/tylertreat/Comcast)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/Comcast)
* [LiteIDE](https://github.com/visualfc/liteide) **star:5746** LiteIDE is a simple, open source, cross-platform Go IDE.   [![star > 2000][Awesome]](https://github.com/visualfc/liteide)   [![There was an update last week][Green]](https://github.com/visualfc/liteide)   [![Contains Chinese documents][CN]](https://github.com/visualfc/liteide)
* [drive](https://github.com/odeke-em/drive) **star:5156** Google Drive client for the commandline.   [![star > 2000][Awesome]](https://github.com/odeke-em/drive)   [![godoc][GoDoc]](https://godoc.org/github.com/odeke-em/drive)
* [nes](https://github.com/fogleman/nes) **star:4278** Nintendo Entertainment System (NES) emulator written in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/nes)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/nes)
* [orange-cat](https://github.com/noraesae/orange-cat)  Markdown previewer written in Go.
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:4212** Proxy to simulate network and system conditions for automated tests.   [![star > 2000][Awesome]](https://github.com/shopify/toxiproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/shopify/toxiproxy)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:2820** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   [![star > 2000][Awesome]](https://github.com/gilbertchen/duplicacy)   [![godoc][GoDoc]](https://godoc.org/github.com/gilbertchen/duplicacy)
* [croc](https://github.com/schollz/croc) **star:2636** Easily and securely send files or folders from one computer to another.   [![star > 2000][Awesome]](https://github.com/schollz/croc)   [![There was an update last week][Green]](https://github.com/schollz/croc)   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/croc)
* [myLG](https://github.com/mehrdadrad/mylg) **star:2245** Command Line Network Diagnostic tool written in Go.   [![star > 2000][Awesome]](https://github.com/mehrdadrad/mylg)   [![It hasn't been updated in the last year][Yellow]](https://github.com/mehrdadrad/mylg)   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/mylg)
* [GoBoy](https://github.com/Humpheh/goboy) **star:2162** Nintendo Game Boy Color emulator written in Go.   [![star > 2000][Awesome]](https://github.com/Humpheh/goboy)   [![godoc][GoDoc]](https://godoc.org/github.com/Humpheh/goboy)
* [Stack Up](https://github.com/pressly/sup) **star:2039** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   [![star > 2000][Awesome]](https://github.com/pressly/sup)   [![There was an update last week][Green]](https://github.com/pressly/sup)   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/sup)
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [lgo](https://github.com/yunabe/lgo) **star:1913** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/yunabe/lgo)
* [Circuit](https://github.com/gocircuit/circuit) **star:1803** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gocircuit/circuit)
* [snap](https://github.com/intelsdi-x/snap) **star:1801** Powerful telemetry framework.   [![godoc][GoDoc]](https://godoc.org/github.com/intelsdi-x/snap)
* [scc](https://github.com/boyter/scc) **star:1161** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   [![There was an update last week][Green]](https://github.com/boyter/scc)   [![godoc][GoDoc]](https://godoc.org/github.com/boyter/scc)
* [Documize](https://github.com/documize/community) **star:949** Modern wiki software that integrates data from SaaS tools.   [![There was an update last week][Green]](https://github.com/documize/community)
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:882** App that displays updates for the Go packages in your GOPATH.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/Go-Package-Store)
* [peg](https://github.com/pointlander/peg) **star:651** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   [![godoc][GoDoc]](https://godoc.org/github.com/pointlander/peg)
* [Pipe](https://github.com/b3log/pipe)  A small and beautiful blogging platform.
* [Leaps](https://github.com/jeffail/leaps) **star:648** Pair programming service using Operational Transforms.   [![godoc][GoDoc]](https://godoc.org/github.com/jeffail/leaps)
* [vFlow](https://github.com/VerizonDigital/vflow) **star:629** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   [![godoc][GoDoc]](https://godoc.org/github.com/VerizonDigital/vflow)
* [gfile](https://github.com/Antonito/gfile) **star:517** Securely transfer files between two computers, without any third party, over WebRTC.   [![godoc][GoDoc]](https://godoc.org/github.com/Antonito/gfile)
* [GoDNS](https://github.com/timothyye/godns) **star:504** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)
* [shell2http](https://github.com/msoap/shell2http) **star:474** Executing shell commands via http server (for prototyping or remote control).   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/shell2http)
* [mockingjay](https://github.com/quii/mockingjay-server) **star:431** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   [![godoc][GoDoc]](https://godoc.org/github.com/quii/mockingjay-server)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:386** Video streaming torrent client.   [![There was an update last week][Green]](https://github.com/Sioro-Neoku/go-peerflix)   [![godoc][GoDoc]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
* [gocc](https://github.com/goccmack/gocc) **star:369** Gocc is a compiler kit for Go written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/goccmack/gocc)
* [wellington](https://github.com/wellington/wellington) **star:292** Sass project management tool, extends the language with sprite functions (like Compass).   [![It hasn't been updated in the last year][Yellow]](https://github.com/wellington/wellington)   [![godoc][GoDoc]](https://godoc.org/github.com/wellington/wellington)
* [ipe](https://github.com/dimiro1/ipe) **star:289** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/ipe)
* [ide](https://github.com/thestrukture/ide) **star:252** Browser accessible IDE. Designed for Go with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thestrukture/ide)
* [Cherry](https://github.com/rafael-santiago/cherry) **star:205** Tiny webchat server in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafael-santiago/cherry)   [![godoc][GoDoc]](https://godoc.org/github.com/rafael-santiago/cherry)
* [Orbit](https://github.com/gulien/orbit) **star:132** A simple tool for running commands and generating files from templates.   [![godoc][GoDoc]](https://godoc.org/github.com/gulien/orbit)
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [joincap](https://github.com/assafmo/joincap) **star:127** Command-line utility for merging multiple pcap files together.   [![godoc][GoDoc]](https://godoc.org/github.com/assafmo/joincap)
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [DDNS](https://github.com/skibish/ddns) **star:118** Personal DDNS client with Digital Ocean Networking DNS as backend.   [![godoc][GoDoc]](https://godoc.org/github.com/skibish/ddns)
* [borg](https://github.com/crufter/borg)  Terminal based search engine for bash snippets.
* [boxed](https://github.com/tejo/boxed) **star:72** Dropbox based blog engine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tejo/boxed)   [![godoc][GoDoc]](https://godoc.org/github.com/tejo/boxed)
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:20** Simple NaCL EC25519 based crypto pipe tool written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/unix4fun/naclpipe)   [![godoc][GoDoc]](https://godoc.org/github.com/unix4fun/naclpipe)
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.   [![It hasn't been updated in the last year][Yellow]](https://github.com/crazcalm/term-quiz)   [![godoc][GoDoc]](https://godoc.org/github.com/crazcalm/term-quiz)
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lucasgomide/snitch)   [![godoc][GoDoc]](https://godoc.org/github.com/lucasgomide/snitch)
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [Gor](https://github.com/buger/gor)  Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:12** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.   [![It hasn't been updated in the last year][Yellow]](https://github.com/diankong/GoDocTooltip)

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1309** Go HTTP request router benchmark and comparison.   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1073** Go web framework benchmark.   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
* [skynet](https://github.com/atemerev/skynet) **star:941** Skynet 1M threads microbenchmark.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:923** Benchmarks of Go serialization methods.   [![There was an update last week][Green]](https://github.com/alecthomas/go_serialization_benchmarks)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Benchmarks of common basic operations for the Go language.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:183** Compare various Image resize algorithms for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/fawick/speedtest-resize)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:129** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/go-benchmarks)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/go-benchmarks)
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:97** Go micro-benchmarks for calculating the speed of language constructs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/feyeleanor/GoSpeed)   [![godoc][GoDoc]](https://godoc.org/github.com/feyeleanor/GoSpeed)
* [autobench](https://github.com/davecheney/autobench) **star:89** Framework to compare the performance between different Go versions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davecheney/autobench)   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/autobench)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:49** Collection of benchmarks for popular Go database/SQL utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tyler-smith/golang-sql-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/tyler-smith/golang-sql-benchmark)
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
* [GoBooks](https://github.com/dariubs/GoBooks) **star:7243** A curated list of Go books.   [![star > 2000][Awesome]](https://github.com/dariubs/GoBooks)
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:13** in Persian.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thedevsir/gosuccinctly)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsir/gosuccinctly)
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:1737** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   [![godoc][GoDoc]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:68** adorable gopher logos.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GolangUA/gopher-logos)
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:30** Go gopher Vector Data [.ai, .svg].   [![It hasn't been updated in the last year][Yellow]](https://github.com/keygx/Go-gopher-Vector)
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:1977** Gopher artworks by Ashley McNamara.   [![godoc][GoDoc]](https://godoc.org/github.com/ashleymcnamara/gophers)
* [gophers](https://github.com/egonelbre/gophers) **star:1762** Free gophers.   [![godoc][GoDoc]](https://godoc.org/github.com/egonelbre/gophers)
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:347** Gopherize yourself.   [![It hasn't been updated in the last year][Yellow]](https://github.com/matryer/gopherize.me)
* [gophers](https://github.com/rogeralsing/gophers) **star:50** random gopher graphics.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rogeralsing/gophers)
* [gophers](https://github.com/sillecelik/go-gopher) **star:41** Gopher amigurumi toy pattern.

## Meetups

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
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

* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:25295** List of other amazingly awesome lists.   [![star > 2000][Awesome]](https://github.com/bayandin/awesome-awesomeness)   [![There was an update last week][Green]](https://github.com/bayandin/awesome-awesomeness)
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [go.dev](https://go.dev/)  A hub for Go developers.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:15363** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.   [![star > 2000][Awesome]](https://github.com/lukasz-madon/awesome-remote-job)
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:142** Collection of Go images, graphics, and art.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/golang-graphics)   [![Archived][Archived]](https://github.com/mholt/golang-graphics)
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [json2go](https://m-zajac.github.io/json2go)  Advanced JSON to Go struct conversion - online tool.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Golang Developer Jobs](https://golangjob.xyz)  Developer Jobs exclusivly for Golang related Roles.
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:36** Collection of Go projects that needs help. Good place to start your open-source way in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ninedraft/gocryforhelp)
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:33113** Golang ebook intro how to build a web app with golang.   [![star > 2000][Awesome]](https://github.com/astaxie/build-web-application-with-golang)   [![There was an update last week][Green]](https://github.com/astaxie/build-web-application-with-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   [![Contains Chinese documents][CN]](https://github.com/astaxie/build-web-application-with-golang)
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:521** A little e-book on Ethereum Development with Go.   [![There was an update last week][Green]](https://github.com/miguelmota/ethereum-development-with-go-book)   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/ethereum-development-with-go-book)   [![Contains Chinese documents][CN]](https://github.com/miguelmota/ethereum-development-with-go-book)
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)  Go's reference card.
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns) **star:11579** Curated list of Go design patterns, recipes and idioms.   [![star > 2000][Awesome]](https://github.com/tmrts/go-patterns)   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/go-patterns)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:8840** Learn Go with test-driven development.   [![star > 2000][Awesome]](https://github.com/quii/learn-go-with-tests)   [![godoc][GoDoc]](https://godoc.org/github.com/quii/learn-go-with-tests)   [![Contains Chinese documents][CN]](https://github.com/quii/learn-go-with-tests)
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera Specialization to learn about Go from scratch.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:879** Examples of Golang compared to Node.js for learning.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/)  Collection of code snippets and tutorials to help tackle every day issues.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1138** Intro to go for experienced programmers.   [![There was an update last week][Green]](https://github.com/mkaz/working-with-go)
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

