# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2020-03-18 12:20:57(Synchronize every day)**

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

* [Oto](https://github.com/hajimehoshi/oto) **star:549** A low-level library to play sound on multiple platforms.   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:339** Go bindings for the PortAudio audio I/O library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gordonklaus/portaudio)   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:286** Music theory models in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:276** Go package capable of generating waveform images from audio streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/waveform)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:227** Go bindings for PortMidi.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:146** Fast and stable ID3 parsing and writing library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [flac](https://github.com/mewkiz/flac) **star:118** Native Go FLAC encoder/decoder with support for FLAC streams.   [![godoc][GoDoc]](https://godoc.org/github.com/mewkiz/flac)
* [mix](https://github.com/go-mix/mix) **star:112** Sequence-based Go-native audio mixer for music apps.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [mp3](https://github.com/tcolgate/mp3) **star:107** Native Go MP3 decoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [go-sox](https://github.com/krig/go-sox) **star:105** libsox bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [malgo](https://github.com/gen2brain/malgo) **star:96** Mini audio library.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/malgo)
* [taglib](https://github.com/wtolson/go-taglib) **star:72** Go bindings for taglib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wtolson/go-taglib)   [![godoc][GoDoc]](https://godoc.org/github.com/wtolson/go-taglib)
* [gaad](https://github.com/Comcast/gaad) **star:70** Native Go AAC bitstream parser.   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:39** Lightweight MP3 decoder library.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) **star:28** libmediainfo bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/go_mediainfo)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/go_mediainfo)
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) **star:28** EasyMidi is a simple and reliable library for working with standard midi file (SMF).   [![It hasn't been updated in the last year][Yellow]](https://github.com/algoGuy/EasyMIDI)   [![godoc][GoDoc]](https://godoc.org/github.com/algoGuy/EasyMIDI)
* [vorbis](https://github.com/mccoyst/vorbis) **star:25** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:9** libsamplerate bindings for go.   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:7233** Golang implementation of JSON Web Tokens (JWT).   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![There was an update last week][Green]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:6260** Authorization library that supports access control models like ACL, RBAC, ABAC.   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![There was an update last week][Green]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [oauth2](https://github.com/golang/oauth2) **star:2771** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:2582** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![There was an update last week][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:2167** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   [![star > 2000][Awesome]](https://github.com/volatiletech/authboss)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [osin](https://github.com/openshift/osin) **star:1585** Golang OAuth2 server library.   [![godoc][GoDoc]](https://godoc.org/github.com/openshift/osin)
* [go-jose](https://github.com/square/go-jose) **star:1489** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.   [![godoc][GoDoc]](https://godoc.org/github.com/square/go-jose)
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1466** Standalone, specification-compliant,  OAuth2 server written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [gologin](https://github.com/dghubble/gologin) **star:1171** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:996** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mikespook/gorbac)   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [loginsrv](https://github.com/tarent/loginsrv) **star:919** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   [![There was an update last week][Green]](https://github.com/tarent/loginsrv)   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [scs](https://github.com/alexedwards/scs) **star:683** Session Manager for HTTP servers.   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [permissions2](https://github.com/xyproto/permissions2) **star:393** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [paseto](https://github.com/o1egl/paseto) **star:319** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![There was an update last week][Green]](https://github.com/o1egl/paseto)   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [httpauth](https://github.com/goji/httpauth) **star:191** HTTP Authentication middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goji/httpauth)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [jeff](https://github.com/abraithwaite/jeff) **star:190** Simple, flexible, secure and idiomatic web session management with pluggable backends.   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:172** JWT middleware for Golang http servers with many configuration options.   [![It hasn't been updated in the last year][Yellow]](https://github.com/adam-hanna/jwt-auth)   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/pascaldekloe/jwt) **star:159** Lightweight JSON Web Token (JWT) library.   [![There was an update last week][Green]](https://github.com/pascaldekloe/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [branca](https://github.com/hako/branca) **star:103** Golang implementation of Branca Tokens.   [![There was an update last week][Green]](https://github.com/hako/branca)   [![godoc][GoDoc]](https://godoc.org/github.com/hako/branca)
* [session](https://github.com/icza/session) **star:99** Go session management for web servers (including support for Google App Engine - GAE).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [sessionup](https://github.com/swithek/sessionup) **star:82** Simple, yet effective HTTP session management and identification package.   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [jwt](https://github.com/robbert229/jwt) **star:78** Clean and easy to use implementation of JSON Web Tokens (JWT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/robbert229/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sjwt](https://github.com/brianvoe/sjwt) **star:61** Simple jwt generator and parser.   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [sessions](https://github.com/adam-hanna/sessions) **star:48** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [rbac](https://github.com/zpatrick/rbac) **star:46** Minimalistic RBAC package for Go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/rbac)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [securecookie](https://github.com/chmike/securecookie) **star:33** Efficient secure cookie encoding/decoding.   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:8** Go session management using the SessionGate Redis module.   [![It hasn't been updated in the last year][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [signedvalue](https://github.com/sashka/signedvalue) **star:8** Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.   [![godoc][GoDoc]](https://godoc.org/github.com/sashka/signedvalue)
* [scope](https://github.com/SonicRoshan/scope) **star:4** Easily Manage OAuth2 Scopes In Go.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:2** provides parser of cookies.txt file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:1977** Simple and clean Telegram bot client.   [![There was an update last week][Green]](https://github.com/Syfaro/telegram-bot-api)   [![godoc][GoDoc]](https://godoc.org/github.com/Syfaro/telegram-bot-api)
* [telebot](https://github.com/tucnak/telebot) **star:1149** Telegram bot framework written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:555** IRC, Slack & Telegram bot written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [go-joe](https://joe-bot.net)  A general-purpose bot library inspired by Hubot but written in Go.
* [slacker](https://github.com/shomali11/slacker) **star:360** Easy to use framework to create Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:290** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   [![It hasn't been updated in the last year][Yellow]](https://github.com/saniales/golang-crypto-trading-bot)   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [Kelp](https://github.com/stellar/kelp) **star:266** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   [![There was an update last week][Green]](https://github.com/stellar/kelp)   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [tbot](https://github.com/yanzay/tbot) **star:251** Telegram bot server with API similar to net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:167** Service oriented IRC bot using Redis and JSON for messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:156** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.   [![godoc][GoDoc]](https://godoc.org/github.com/oklahomer/go-sarah)
* [hanu](https://github.com/sbstjn/hanu) **star:119** Framework for writing Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:96** Libary to write bots for twitch.tv chat   [![There was an update last week][Green]](https://github.com/gempir/go-twitch-irc)   [![godoc][GoDoc]](https://godoc.org/github.com/gempir/go-twitch-irc)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:91** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/olebedev/go-tgbot)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:60** Framework for building Telegram bots.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [govkbot](https://github.com/nikepan/govkbot) **star:30** Simple Go [VK](https://vk.com) bot library.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:24** Another framework for building Slack bots.   [![There was an update last week][Green]](https://github.com/alexandre-normand/slackscot)   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) **star:18** A Discord bot for managing ephemeral roles based upon voice channel member presence.   [![There was an update last week][Green]](https://github.com/ewohltman/ephemeral-roles)   [![godoc][GoDoc]](https://godoc.org/github.com/ewohltman/ephemeral-roles)
* [micha](https://github.com/onrik/micha) **star:12** Go Library for Telegram bot api.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:16135** Commander for modern Go CLI interactions.   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![There was an update last week][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:13333** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![There was an update last week][Green]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2820** Command line and flag parser supporting sub commands.   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1652** go command line option parser.   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [Dnote](https://github.com/dnote/dnote) **star:1621** A simple command line notebook with multi-device sync.   [![There was an update last week][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [readline](https://github.com/chzyer/readline) **star:1472** Pure golang implementation that provides most features in GNU-Readline under MIT license.   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/readline)
* [docopt.go](https://github.com/docopt/docopt.go) **star:1225** Command-line arguments parser that will make you smile.   [![godoc][GoDoc]](https://godoc.org/github.com/docopt/docopt.go)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1100** Go library for implementing command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [pflag](https://github.com/spf13/pflag) **star:1025** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [cli-init](https://github.com/tcnksm/gcli) **star:894** The easy way to start building Golang command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcnksm/gcli)   [![godoc][GoDoc]](https://godoc.org/github.com/tcnksm/gcli)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [go-arg](https://github.com/alexflint/go-arg) **star:846** Struct-based argument parsing in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [complete](https://github.com/posener/complete) **star:676** Write bash completions in Go + Go command bash completion.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [liner](https://github.com/peterh/liner) **star:667** Go readline-like library for command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [mow.cli](https://github.com/jawher/mow.cli) **star:658** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [flaggy](https://github.com/integrii/flaggy) **star:634** A robust and idiomatic flags package with excellent subcommand support.   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:518** Feature-rich and easy to use command-line package based on golang struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:387** Unikernel Builder/Orchestrator.   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:176** Command line argument parser inspired by Python's argparse module.   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [sflags](https://github.com/octago/sflags) **star:110** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [commandeer](https://github.com/jaffee/commandeer) **star:109** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [wmenu](https://github.com/dixonwille/wmenu) **star:102** Easy to use menu structure for cli applications that prompts users to make choices.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [flag](https://github.com/cosiner/flag) **star:102** Simple but powerful command line option parsing library for Go supporting subcommand.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cosiner/flag)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:101** Small command line interface framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ukautz/clif)   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [job](https://github.com/liujianping/job) **star:70** JOB, make your short-term command as a long-term job.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/job)
* [cli](https://github.com/teris-io/cli) **star:66** Simple and complete API for building command line interfaces in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [1build](https://github.com/gopinath-langote/1build) **star:55** Command line tool to frictionlessly manage project-specific commands.   [![There was an update last week][Green]](https://github.com/gopinath-langote/1build)   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [env](https://github.com/codingconcepts/env) **star:48** Tag-based environment configuration for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [wlog](https://github.com/dixonwille/wlog) **star:42** Simple logging interface that supports cross-platform color and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wlog)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [gocmd](https://github.com/devfacet/gocmd) **star:37** Go library for building command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [strumt](https://github.com/antham/strumt) **star:34** Library to create prompt chain.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [flagvar](https://github.com/sgreben/flagvar) **star:33** A collection of flag argument types for Go's standard `flag` package.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [cmdr](https://github.com/hedzr/cmdr) **star:30** A POSIX/GNU style, getopt-like command-line UI Go library.   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [clîr](https://github.com/leaanthony/clir) **star:21** A Simple and Clear CLI library. Dependency free.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:19** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [argv](https://github.com/cosiner/argv) **star:19** Go library to split command line string as arguments array using the bash syntax.   [![There was an update last week][Green]](https://github.com/cosiner/argv)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [go-commander](https://github.com/yitsushi/go-commander) **star:15** Go library to simplify CLI workflow.   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [sand](https://github.com/Zaba505/sand) **star:9** Simple API for creating interpreters and so much more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Zaba505/sand)   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [ts](https://github.com/liujianping/ts) **star:8** Timestamp convert & compare tool.   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/ts)
* [cmd](https://github.com/posener/cmd) **star:6** Extends the standard `flag` package to support sub commands and more in idomatic way.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:9621** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [gocui](https://github.com/jroimartin/gocui) **star:6052** Minimalist Go library aimed at creating Console User Interfaces.   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [termbox-go](https://github.com/nsf/termbox-go) **star:3700** Termbox is a library for creating cross-platform text-based interfaces.   [![star > 2000][Awesome]](https://github.com/nsf/termbox-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/termbox-go)
* [go-prompt](https://github.com/c-bata/go-prompt) **star:2881** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1647** Flexible library to render progress bars in terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1315** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [uilive](https://github.com/gosuri/uilive) **star:1110** Library for updating terminal output in realtime.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [termdash](https://github.com/mum4k/termdash) **star:965** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [progressbar](https://github.com/schollz/progressbar) **star:869** Basic thread-safe progress bar that works in every OS.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/progressbar)
* [mpb](https://github.com/vbauerster/mpb) **star:856** Multi progress bar for terminal applications.   [![There was an update last week][Green]](https://github.com/vbauerster/mpb)   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [aurora](https://github.com/logrusorgru/aurora) **star:790** ANSI terminal colors that supports fmt.Printf/Sprintf.   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [uitable](https://github.com/gosuri/uitable) **star:547** Library to improve readability in terminal apps using tabular data.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:419** Colorable writer for windows.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:405** isatty for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [gookit/color](https://github.com/gookit/color) **star:329** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![Contains Chinese documents][CN]](https://github.com/gookit/color)
* [chalk](https://github.com/ttacon/chalk) **star:329** Intuitive package for prettifying terminal/console output.   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [tabby](https://github.com/cheynewallace/tabby) **star:264** A tiny library for super simple Golang tables.   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [simpletable](https://github.com/alexeyco/simpletable) **star:216** Simple tables in terminal with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:201** Go library for color output in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-colortext)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [cfmt](https://github.com/mingrammer/cfmt) **star:72** Contextual fmt inspired by bootstrap color classes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mingrammer/cfmt)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:36** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/tabular)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [ctc](https://github.com/wzshiming/ctc) **star:17** The non-invasive cross-platform terminal color library does not need to modify the Print method.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/ctc)
* [colourize](https://github.com/TreyBastian/colourize) **star:17** Go library for ANSI colour text in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [go-ataman](https://github.com/workanator/go-ataman) **star:8** Go library for rendering ANSI colored text templates in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/workanator/go-ataman)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:11502** Go configuration with fangs.   [![star > 2000][Awesome]](https://github.com/spf13/viper)   [![There was an update last week][Green]](https://github.com/spf13/viper)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:2822** Go library for managing configuration data from environment variables.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv) **star:2685** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![There was an update last week][Green]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [ini](https://github.com/go-ini/ini) **star:1959** Go package to read and write INI files.   [![There was an update last week][Green]](https://github.com/go-ini/ini)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:1448** Parse environment variables to Go structs (with defaults).   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [konfig](https://github.com/lalamove/konfig) **star:548** Composable, observable and performant config handling for Go for the distributed processing era.   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [confita](https://github.com/heetch/confita) **star:296** Load configuration in cascade from multiple backends into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:248** Lightweight configuration manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tucnak/store)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/JeremyLoy/config) **star:231** Cloud native application configuration. Bind ENV to structs in only two lines.   [![There was an update last week][Green]](https://github.com/JeremyLoy/config)   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [config](https://github.com/olebedev/config) **star:225** JSON or YAML configuration wrapper with environment variables and flags parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/config)
* [joshbetz/config](https://github.com/joshbetz/config) **star:198** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.   [![godoc][GoDoc]](https://godoc.org/github.com/joshbetz/config)
* [hjson](https://github.com/hjson/hjson-go) **star:195** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![There was an update last week][Green]](https://github.com/hjson/hjson-go)   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [envconfig](https://github.com/vrischmann/envconfig) **star:172** Read your configuration from environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/vrischmann/envconfig)
* [koanf](https://github.com/knadh/koanf) **star:170** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [gookit/config](https://github.com/gookit/config) **star:139** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   [![Contains Chinese documents][CN]](https://github.com/gookit/config)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:127** read INI-style configuration files into Go structs; supports user-defined types and subsections.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [goConfig](https://github.com/crgimenes/goConfig) **star:125** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   [![There was an update last week][Green]](https://github.com/crgimenes/goConfig)   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [envh](https://github.com/antham/envh) **star:94** Helpers to manage environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [envcfg](https://github.com/tomazk/envcfg) **star:91** Un-marshaling environment variables to Go structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomazk/envcfg)   [![godoc][GoDoc]](https://godoc.org/github.com/tomazk/envcfg)
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:58** Minimalistic configuration reader (from files, ENV, and wherever you want).   [![godoc][GoDoc]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [gofigure](https://github.com/ian-kent/gofigure) **star:58** Go application configuration made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [configure](https://github.com/paked/configure) **star:54** Provides configuration through multiple sources, including JSON, flags and environment variables.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paked/configure)   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [harvester](https://github.com/beatlabs/harvester) **star:52** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.   [![There was an update last week][Green]](https://github.com/beatlabs/harvester)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/harvester)
* [config](https://github.com/golobby/config) **star:51** A lightweight yet powerful config package for Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:48** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [ingo](https://github.com/schachmat/ingo) **star:28** Flags persisted in an ini-like config file.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schachmat/ingo)   [![godoc][GoDoc]](https://godoc.org/github.com/schachmat/ingo)
* [go-up](https://github.com/ufoscout/go-up) **star:26** A simple configuration library with recursive placeholders resolution and no magic.   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:22** Golang package for parsing ini-style configuration files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sasbury/mini)   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:13** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.   [![godoc][GoDoc]](https://godoc.org/github.com/the4thamigo-uk/conflate)
* [genv](https://github.com/sakirsensoy/genv) **star:9** Read environment variables easily with dotenv support.   [![godoc][GoDoc]](https://godoc.org/github.com/sakirsensoy/genv)
* [envconf](https://github.com/ian-kent/envconf) **star:9** Configuration from environment.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/envconf)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/envconf)
* [sprbox](https://github.com/oblq/sprbox) **star:5** Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/sprbox)
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) **star:3** Go utility for loading configuration parameters from AWS SSM (Parameter Store).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-ssm-config)
* [configuration](https://github.com/BoRuDar/configuration) **star:3** Library for initializing configuration structs from env variables, files, flags and 'default' tag.   [![There was an update last week][Green]](https://github.com/BoRuDar/configuration)   [![godoc][GoDoc]](https://godoc.org/github.com/BoRuDar/configuration)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:2** Simple useful package for read environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/nasermirzaei89/env)
* [onion](http://github.com/goraz/onion)  Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:20682** Drone is a Continuous Integration platform built on Docker, written in Go.   [![star > 2000][Awesome]](https://github.com/drone/drone)   [![There was an update last week][Green]](https://github.com/drone/drone)   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [CDS](https://github.com/ovh/cds) **star:2744** Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.   [![star > 2000][Awesome]](https://github.com/ovh/cds)   [![There was an update last week][Green]](https://github.com/ovh/cds)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/cds)
* [goveralls](https://github.com/mattn/goveralls) **star:634** Go integration for Coveralls.io continuous code coverage tracking system.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/goveralls)
* [overalls](https://github.com/go-playground/overalls) **star:101** Multi-Package go project coverprofile for tools like goveralls.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/overalls)
* [duci](https://github.com/duck8823/duci) **star:53** A simple ci server no needs domain specific languages.   [![There was an update last week][Green]](https://github.com/duck8823/duci)   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:44** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:12** Recursive coverage testing tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:431** Pure Go CSS Preprocessor.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/gcss)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass) **star:155** Go wrapper to the 100% Sass compatible libsass project.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wellington/go-libsass)

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:7918** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5481** Collection of useful, performant, and thread-safe data structures.   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1423** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1236** Probabilistic data structures for processing continuous, unbounded streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/BoomFilters)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [gota](https://github.com/kniren/gota) **star:1107** Implementation of dataframes, series, and data wrangling methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [roaring](https://github.com/RoaringBitmap/roaring) **star:797** Go package implementing compressed bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/RoaringBitmap/roaring)
* [willf/bloom](https://github.com/willf/bloom) **star:759** Go package implementing Bloom filters.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bloom)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:691** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:593** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [bitset](https://github.com/willf/bitset) **star:539** Go package implementing bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [trie](https://github.com/derekparker/trie) **star:463** Trie implementation in Go.   [![There was an update last week][Green]](https://github.com/derekparker/trie)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/trie)
* [gocache](https://github.com/eko/gocache) **star:375** A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.   [![There was an update last week][Green]](https://github.com/eko/gocache)   [![godoc][GoDoc]](https://godoc.org/github.com/eko/gocache)
* [algorithms](https://github.com/shady831213/algorithms) **star:364** Algorithms and data structures.CLRS study.   [![godoc][GoDoc]](https://godoc.org/github.com/shady831213/algorithms)
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:320** In-memory geo index.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hailocab/go-geoindex)   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [mafsa](https://github.com/smartystreets/mafsa) **star:281** MA-FSA implementation with Minimal Perfect Hashing.   [![godoc][GoDoc]](https://godoc.org/github.com/smartystreets/mafsa)   [![Archived][Archived]](https://github.com/smartystreets/mafsa)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:217** Skip list implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [hilbert](https://github.com/google/hilbert) **star:197** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   [![It hasn't been updated in the last year][Yellow]](https://github.com/google/hilbert)   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [merkletree](https://github.com/cbergoon/merkletree) **star:182** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:135** Binary packer and unpacker helps user build custom binary stream.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhuangsirui/binpacker)   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:133** Go implementation of Adaptive Radix Tree.   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [bloom](https://github.com/zhenjl/bloom) **star:131** Bloom filters implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [ttlcache](https://github.com/diegobernardes/ttlcache) **star:129** In-memory LRU string-interface{} map with expiration for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/diegobernardes/ttlcache)
* [iter](https://github.com/disksing/iter) **star:118** Go implementation of C++ STL iterators and algorithms.   [![godoc][GoDoc]](https://godoc.org/github.com/disksing/iter)   [![Contains Chinese documents][CN]](https://github.com/disksing/iter)
* [skiplist](https://github.com/MauriceGit/skiplist) **star:117** Very fast Go Skiplist implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [deque](https://github.com/gammazero/deque) **star:107** Fast ring-buffer deque (double-ended queue).   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [ring](https://github.com/TheTannerRyan/ring) **star:105** Go implementation of a high performance, thread safe bloom filter.   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:101** Region quadtrees with efficient point location and neighbour finding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aurelien-rainone/go-rquad)   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [encoding](https://github.com/zhenjl/encoding) **star:100** Integer Compression Libraries for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/encoding)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:87** A small, powerful and flexible merge library.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [bit](https://github.com/yourbasic/bit) **star:84** Golang set data structure with bonus bit-twiddling functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bit)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [gostl](https://github.com/liyue201/gostl) **star:77** Data structure and algorithm library for go, designed to provide functions similar to C++ STL.   [![godoc][GoDoc]](https://godoc.org/github.com/liyue201/gostl)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:72** Implementation to calculate levenshtein distance in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [skiplist](https://github.com/gansidui/skiplist) **star:69** Skiplist implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:64** Concurrent FIFO queue.   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [bloom](https://github.com/yourbasic/bloom) **star:49** Golang Bloom filter implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bloom)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:49** Fast in-memory key:value store/cache library. Pointer caches.   [![godoc][GoDoc]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:46** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/count-min-log)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/count-min-log)
* [levenshtein](https://github.com/agext/levenshtein) **star:40** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![There was an update last week][Green]](https://github.com/agext/levenshtein)   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [remember-go](https://github.com/rocketlaunchr/remember-go) **star:38** A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/remember-go)
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:27** Highly concurrent drop-in replacement for `bufio.Writer`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/free/concurrent-writer)   [![godoc][GoDoc]](https://godoc.org/github.com/free/concurrent-writer)
* [crunch](https://github.com/superwhiskers/crunch) **star:23** Go package implementing buffers for handling various datatypes easily.   [![There was an update last week][Green]](https://github.com/superwhiskers/crunch)   [![godoc][GoDoc]](https://godoc.org/github.com/superwhiskers/crunch)
* [pipeline](https://github.com/hyfather/pipeline) **star:22** An implementation of pipelines with fan-in and fan-out.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyfather/pipeline)   [![godoc][GoDoc]](https://godoc.org/github.com/hyfather/pipeline)
* [goset](https://github.com/zoumo/goset) **star:20** A useful Set collection implementation for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [typ](https://github.com/gurukami/typ) **star:16** Null Types, Safe primitive type conversion and fetching value from complex structures.   [![godoc][GoDoc]](https://godoc.org/github.com/gurukami/typ)
* [deque](https://github.com/edwingeng/deque) **star:15** A highly optimized double-ended queue.   [![There was an update last week][Green]](https://github.com/edwingeng/deque)   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/deque)
* [hide](https://github.com/emvi/hide) **star:14** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![It hasn't been updated in the last year][Yellow]](https://github.com/emvi/hide)   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [dict](https://github.com/srfrog/dict) **star:11** Python-like dictionaries (dict) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [go-ef](https://github.com/amallia/go-ef) **star:11** A Go implementation of the Elias-Fano encoding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/amallia/go-ef)   [![godoc][GoDoc]](https://godoc.org/github.com/amallia/go-ef)
* [null](https://github.com/emvi/null) **star:9** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [mspm](https://github.com/BlackRabbitt/mspm) **star:9** Multi-String Pattern Matching Algorithm for information retrieval.   [![It hasn't been updated in the last year][Yellow]](https://github.com/BlackRabbitt/mspm)   [![godoc][GoDoc]](https://godoc.org/github.com/BlackRabbitt/mspm)
* [set](https://github.com/StudioSol/set) **star:7** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [ptrie](https://github.com/viant/ptrie) **star:7** An implementation of prefix tree.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/ptrie)
* [treap](https://github.com/perdata/treap) **star:7** Persistent, fast ordered map using tree heaps.   [![godoc][GoDoc]](https://godoc.org/github.com/perdata/treap)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:6** Map with expiring key-value pairs.   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)
* [gofal](https://github.com/xxjwxc/gofal) **star:6** fractional api for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gofal)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gofal)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:3** Tools for parse JSON-like logs for collecting unique fields and events.   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:29652** Monitoring system and time series database.   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   [![There was an update last week][Green]](https://github.com/prometheus/prometheus)   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [tidb](https://github.com/pingcap/tidb) **star:22829** TiDB is a distributed SQL database. Inspired by the design of Google F1.   [![star > 2000][Awesome]](https://github.com/pingcap/tidb)   [![There was an update last week][Green]](https://github.com/pingcap/tidb)   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/tidb)   [![Contains Chinese documents][CN]](https://github.com/pingcap/tidb)
* [influxdb](https://github.com/influxdb/influxdb) **star:18486** Scalable datastore for metrics, events, and real-time analytics.   [![star > 2000][Awesome]](https://github.com/influxdb/influxdb)   [![There was an update last week][Green]](https://github.com/influxdb/influxdb)   [![godoc][GoDoc]](https://godoc.org/github.com/influxdb/influxdb)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:17927** Scalable, Geo-Replicated, Transactional Datastore.   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   [![There was an update last week][Green]](https://github.com/cockroachdb/cockroach)   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [dgraph](https://github.com/dgraph-io/dgraph) **star:12682** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   [![There was an update last week][Green]](https://github.com/dgraph-io/dgraph)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [groupcache](https://github.com/golang/groupcache) **star:8468** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![There was an update last week][Green]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:7399** Fast key-value store in Go.   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   [![There was an update last week][Green]](https://github.com/dgraph-io/badger)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:5726** The lightweight, distributed, relational database built on SQLite.   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [BigCache](https://github.com/allegro/bigcache) **star:3661** Efficient key/value cache for gigabytes of data.   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![There was an update last week][Green]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:3605** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [go-cache](https://github.com/pmylund/go-cache) **star:3591** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![There was an update last week][Green]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3242** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   [![star > 2000][Awesome]](https://github.com/siddontang/ledisdb)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/ledisdb)
* [bbolt](https://github.com/etcd-io/bbolt) **star:3047** An embedded key/value database for Go.   [![star > 2000][Awesome]](https://github.com/etcd-io/bbolt)   [![There was an update last week][Green]](https://github.com/etcd-io/bbolt)   [![godoc][GoDoc]](https://godoc.org/github.com/etcd-io/bbolt)
* [buntdb](https://github.com/tidwall/buntdb) **star:2661** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2464** Your NoSQL database powered by Golang.   [![star > 2000][Awesome]](https://github.com/HouzuoGuo/tiedot)   [![godoc][GoDoc]](https://godoc.org/github.com/HouzuoGuo/tiedot)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:1942** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   [![There was an update last week][Green]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [cache2go](https://github.com/muesli/cache2go) **star:1228** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:1110** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   [![Contains Chinese documents][CN]](https://github.com/xujiajun/nutsdb)
* [GCache](https://github.com/bluele/gcache) **star:1082** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:1003** CovenantSQL is a SQL database on blockchain.   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:860** Home-grown disk-backed key-value store.   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [moss](https://github.com/couchbase/moss) **star:755** Moss is a simple LSM key-value storage engine written in 100% Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/moss)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:723** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
* [eliasdb](https://github.com/krotik/eliasdb) **star:561** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [levigo](https://github.com/jmhodges/levigo) **star:378** Levigo is a Go wrapper for LevelDB.   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [Bitcask](https://github.com/prologic/bitcask) **star:303** Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).   [![There was an update last week][Green]](https://github.com/prologic/bitcask)   [![godoc][GoDoc]](https://godoc.org/github.com/prologic/bitcask)
* [pudge](https://github.com/recoilme/pudge) **star:249** Fast and simple  key/value store written using Go's standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [piladb](https://github.com/fern4lvarez/piladb) **star:173** Lightweight RESTful database engine based on stack data structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fern4lvarez/piladb)   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Vasto](https://github.com/chrislusf/vasto) **star:171** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/vasto)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [Kivik](https://github.com/go-kivik/kivik) **star:155** Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.   [![There was an update last week][Green]](https://github.com/go-kivik/kivik)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kivik/kivik)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:94** Key-value store with persistence.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:89** Tiny flat file JSON store.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nanobox-io/golang-scribble)   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:47** RESTful caching micro-service backed by Couchbase server.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [bcache](https://github.com/iwanbk/bcache) **star:44** Eventually consistent distributed in-memory  cache Go library.   [![godoc][GoDoc]](https://godoc.org/github.com/iwanbk/bcache)
* [cache](https://github.com/akyoto/cache) **star:35** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)
* [Databunker](https://github.com/paranoidguy/databunker) **star:34** Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.   [![There was an update last week][Green]](https://github.com/paranoidguy/databunker)   [![godoc][GoDoc]](https://godoc.org/github.com/paranoidguy/databunker)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:33** BigCache with clustering support and individual item expiration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][GoDoc]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
* [Coffer](https://github.com/claygod/coffer) **star:20** Simple ACID key-value database that supports transactions.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/coffer)
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:14** Key-value store for temporary items.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/tempdb)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/tempdb)
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) **star:12** Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kapitan-k/gorocksdb)   [![godoc][GoDoc]](https://godoc.org/github.com/kapitan-k/gorocksdb)
* [tracedb](https://github.com/unit-io/tracedb) **star:9** Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application.   [![There was an update last week][Green]](https://github.com/unit-io/tracedb)   [![godoc][GoDoc]](https://godoc.org/github.com/unit-io/tracedb)
* [gostore](https://github.com/twharmon/gostore) **star:4** Gostore is a simple, durable, embedded key-value storage engine written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/gostore)

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:3843** Database migrations. CLI and Golang library.   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![There was an update last week][Green]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:1645** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [skeema](https://github.com/skeema/skeema) **star:638** Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.   [![godoc][GoDoc]](https://godoc.org/github.com/skeema/skeema)
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:419** Database schema migration helper for Gorm ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [goose](https://github.com/steinbacher/goose) **star:136** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   [![It hasn't been updated in the last year][Yellow]](https://github.com/steinbacher/goose)   [![godoc][GoDoc]](https://godoc.org/github.com/steinbacher/goose)
* [darwin](https://github.com/GuiaBolso/darwin) **star:99** Database schema evolution library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/GuiaBolso/darwin)
* [migrator](https://github.com/lopezator/migrator) **star:90** Dead simple Go database migration library.   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:45** A Go package to help write migrations with go-pg/pg.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [gondolier](https://github.com/emvi/gondolier) **star:28** Database migration library using struct decorators.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/gondolier)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pravasan/pravasan)
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:23** Django style fixtures for Golang's excellent built-in database/sql library.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-fixtures)
* [avro](https://github.com/khezen/avro) **star:12** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)
* [schema](https://github.com/adlio/schema) **star:4** Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/schema)

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:9675** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   [![There was an update last week][Green]](https://github.com/youtube/vitess)   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [pgweb](https://github.com/sosedoff/pgweb) **star:6315** Web-based PostgreSQL database browser.   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![There was an update last week][Green]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [kingshard](https://github.com/flike/kingshard) **star:5045** kingshard is a high performance proxy for MySQL powered by Golang.   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   [![Contains Chinese documents][CN]](https://github.com/flike/kingshard)
* [orchestrator](https://github.com/github/orchestrator) **star:3422** MySQL replication topology manager & visualizer.   [![star > 2000][Awesome]](https://github.com/github/orchestrator)   [![There was an update last week][Green]](https://github.com/github/orchestrator)   [![godoc][GoDoc]](https://godoc.org/github.com/github/orchestrator)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:2821** Sync your MySQL data into Elasticsearch automatically.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:2264** Go toolset to handle MySQL protocol and replication.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql)   [![There was an update last week][Green]](https://github.com/siddontang/go-mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [pREST](https://github.com/nuveo/prest) **star:2220** Serve a RESTful API from any PostgreSQL database.   [![star > 2000][Awesome]](https://github.com/nuveo/prest)   [![godoc][GoDoc]](https://godoc.org/github.com/nuveo/prest)
* [chproxy](https://github.com/Vertamedia/chproxy) **star:383** HTTP proxy for ClickHouse database.   [![There was an update last week][Green]](https://github.com/Vertamedia/chproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/Vertamedia/chproxy)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:182** Collects small insterts and sends big requests to ClickHouse servers.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [myreplication](https://github.com/2tvenom/myreplication) **star:158** MySql binary log replication listener. Supports statement and row based replication.   [![It hasn't been updated in the last year][Yellow]](https://github.com/2tvenom/myreplication)   [![godoc][GoDoc]](https://godoc.org/github.com/2tvenom/myreplication)
* [octillery](https://github.com/knocknote/octillery) **star:75** Go package for sharding databases ( Supports every ORM or raw SQL ).   [![There was an update last week][Green]](https://github.com/knocknote/octillery)   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [dbbench](https://github.com/sj14/dbbench) **star:35** Database benchmarking tool with support for several databases and scripts.   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [prep](https://github.com/hexdigest/prep) **star:25** Use prepared SQL statements without changing your code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/prep)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/prep)
* [datagen](https://github.com/codingconcepts/datagen) **star:16** A fast data generator that's multi-table aware and supports multi-row DML.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/datagen)
* [rwdb](https://github.com/andizzle/rwdb) **star:11** rwdb provides read replica capability for multiple database servers setup.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andizzle/rwdb)   [![godoc][GoDoc]](https://godoc.org/github.com/andizzle/rwdb)
* [bucket](https://github.com/PumpkinSeed/bucket) **star:6** Optimized data structure framework for Couchbase specialized on one bucket usage.   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/bucket)

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:2797** Go library that helps you build SQL queries.   [![star > 2000][Awesome]](https://github.com/Masterminds/squirrel)   [![There was an update last week][Green]](https://github.com/Masterminds/squirrel)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/squirrel)
* [xo](https://github.com/knq/xo) **star:2394** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   [![star > 2000][Awesome]](https://github.com/knq/xo)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/xo)
* [gendry](https://github.com/didi/gendry) **star:953** Non-invasive SQL builder and powerful data binder.   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   [![Contains Chinese documents][CN]](https://github.com/didi/gendry)
* [goqu](https://github.com/doug-martin/goqu) **star:739** Idiomatic SQL builder and query library.   [![There was an update last week][Green]](https://github.com/doug-martin/goqu)   [![godoc][GoDoc]](https://godoc.org/github.com/doug-martin/goqu)
* [Dotsql](https://github.com/gchaincl/dotsql) **star:507** Go library that helps you keep sql files in one place and use them with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/dotsql)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:462** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
* [jet](https://github.com/go-jet/jet) **star:230** Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.   [![godoc][GoDoc]](https://godoc.org/github.com/go-jet/jet)
* [sqrl](https://github.com/elgris/sqrl) **star:202** SQL query builder, fork of Squirrel with improved performance.   [![godoc][GoDoc]](https://godoc.org/github.com/elgris/sqrl)
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [dbq](https://github.com/rocketlaunchr/dbq) **star:98** Zero boilerplate database operations for Go.   [![There was an update last week][Green]](https://github.com/rocketlaunchr/dbq)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)
* [igor](https://github.com/galeone/igor) **star:81** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [godbal](https://github.com/xujiajun/godbal) **star:50** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xujiajun/godbal)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [sqlf](https://github.com/leporo/sqlf) **star:7** Fast SQL query builder.   [![godoc][GoDoc]](https://godoc.org/github.com/leporo/sqlf)
* [qry](https://github.com/HnH/qry) **star:6** Tool that generates constants from files with raw SQL queries.   [![godoc][GoDoc]](https://godoc.org/github.com/HnH/qry)
* [mpath](https://github.com/spacetab-io/mpath-go) **star:5** MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation.   [![godoc][GoDoc]](https://godoc.org/github.com/spacetab-io/mpath-go)
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:9135** MySQL driver for Go.   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   [![There was an update last week][Green]](https://github.com/go-sql-driver/mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [pq](https://github.com/lib/pq) **star:5722** Pure Go Postgres driver for database/sql.   [![star > 2000][Awesome]](https://github.com/lib/pq)   [![There was an update last week][Green]](https://github.com/lib/pq)   [![godoc][GoDoc]](https://godoc.org/github.com/lib/pq)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:3885** SQLite3 driver for go that uses database/sql.   [![star > 2000][Awesome]](https://github.com/mattn/go-sqlite3)
    * [pgx](https://github.com/jackc/pgx) **star:2434** PostgreSQL driver supporting features beyond those exposed by database/sql.   [![star > 2000][Awesome]](https://github.com/jackc/pgx)   [![There was an update last week][Green]](https://github.com/jackc/pgx)   [![godoc][GoDoc]](https://godoc.org/github.com/jackc/pgx)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1154** Microsoft MSSQL driver for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8) **star:448** Oracle driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-oci8)
    * [goracle](https://github.com/go-goracle/goracle) **star:280** Oracle driver for Go, using the ODPI-C driver.
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:119** Firebird RDBMS SQL driver for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/nakagami/firebirdsql)
    * [go-adodb](https://github.com/mattn/go-adodb) **star:102** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [gofreetds](https://github.com/minus5/gofreetds) **star:98** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).   [![It hasn't been updated in the last year][Yellow]](https://github.com/minus5/gofreetds)   [![godoc][GoDoc]](https://godoc.org/github.com/minus5/gofreetds)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:47** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [bgc](https://github.com/viant/bgc) **star:13** Datastore Connectivity for BigQuery for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:8252** Redis client for Golang.   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   [![There was an update last week][Green]](https://github.com/go-redis/redis)   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo) **star:7077** Redigo is a Go client for the Redis database.   [![star > 2000][Awesome]](https://github.com/gomodule/redigo)   [![godoc][GoDoc]](https://godoc.org/github.com/gomodule/redigo)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:4138** Official MongoDB driver for the Go language.   [![star > 2000][Awesome]](https://github.com/mongodb/mongo-go-driver)   [![There was an update last week][Green]](https://github.com/mongodb/mongo-go-driver)   [![godoc][GoDoc]](https://godoc.org/github.com/mongodb/mongo-go-driver)
    * [mgo](https://github.com/globalsign/mgo) **star:1762** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   [![godoc][GoDoc]](https://godoc.org/github.com/globalsign/mgo)
    * [gorethink](https://github.com/dancannon/gorethink) **star:1504** Go language driver for RethinkDB.   [![There was an update last week][Green]](https://github.com/dancannon/gorethink)   [![godoc][GoDoc]](https://godoc.org/github.com/dancannon/gorethink)
    * [redeo](https://github.com/bsm/redeo) **star:369** Redis-protocol compatible TCP servers/services.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redeo)
    * [neoism](https://github.com/jmcvetta/neoism) **star:363** Neo4j client for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:316** Aerospike client in Go language.   [![There was an update last week][Green]](https://github.com/aerospike/aerospike-client-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [gocb](https://github.com/couchbase/gocb) **star:307** Official Couchbase Go SDK.   [![There was an update last week][Green]](https://github.com/couchbase/gocb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/gocb)
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:298** Couchbase client in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:117** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:74** Neo4j REST Client in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davemeehan/Neo4j-GO)   [![godoc][GoDoc]](https://godoc.org/github.com/davemeehan/Neo4j-GO)
    * [godis](https://github.com/piaohao/godis) **star:72** redis client implement by golang, inspired by jedis.   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [arangolite](https://github.com/solher/arangolite) **star:66** Lightweight golang driver for ArangoDB.   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [dynago](https://github.com/underarmour/dynago) **star:66** Dynago is a principle of least surprise client for DynamoDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/underarmour/dynago)   [![godoc][GoDoc]](https://godoc.org/github.com/underarmour/dynago)
    * [mgm](https://github.com/kamva/mgm) **star:62** MongoDB model-based ODM for Go (based on official MongoDB driver).   [![godoc][GoDoc]](https://godoc.org/github.com/kamva/mgm)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:34** Go client library for Pilosa.   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [forestdb](https://github.com/couchbase/goforestdb) **star:28** Go bindings for ForestDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/couchbase/goforestdb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/goforestdb)
    * [neo4j](https://github.com/cihangir/neo4j) **star:25** Neo4j Rest API Bindings for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cihangir/neo4j)   [![godoc][GoDoc]](https://godoc.org/github.com/cihangir/neo4j)
    * [goriak](https://github.com/zegl/goriak) **star:24** Go language driver for Riak KV.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zegl/goriak)   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [xredis](https://github.com/shomali11/xredis) **star:10** Typesafe, customizable, clean & easy to use Redis client.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)
    * [godscache](https://github.com/defcronyke/godscache) **star:7** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.   [![It hasn't been updated in the last year][Yellow]](https://github.com/defcronyke/godscache)   [![godoc][GoDoc]](https://godoc.org/github.com/defcronyke/godscache)
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [asc](https://github.com/viant/asc) **star:4** Datastore Connectivity for Aerospike for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/asc)

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:6413** Modern text indexing library for go.   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![There was an update last week][Green]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [riot](https://github.com/go-ego/riot) **star:5091** Go Open Source, Distributed, Simple and efficient Search Engine.   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   [![Contains Chinese documents][CN]](https://github.com/go-ego/riot)
    * [elastic](https://github.com/olivere/elastic) **star:4802** Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/olivere/elastic)   [![godoc][GoDoc]](https://godoc.org/github.com/olivere/elastic)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:2250** Official Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/elastic/go-elasticsearch)   [![There was an update last week][Green]](https://github.com/elastic/go-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/elastic/go-elasticsearch)
    * [elastigo](https://github.com/mattbaird/elastigo) **star:952** Elasticsearch client library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattbaird/elastigo)   [![godoc][GoDoc]](https://godoc.org/github.com/mattbaird/elastigo)
    * [elasticsql](https://github.com/cch123/elasticsql) **star:495** Convert sql to elasticsearch dsl in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cch123/elasticsql)   [![godoc][GoDoc]](https://godoc.org/github.com/cch123/elasticsql)
    * [skizze](https://github.com/seiflotfy/skizze) **star:72** probabilistic data-structures service and storage.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/skizze)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OwnLocal/goes)   [![godoc][GoDoc]](https://godoc.org/github.com/OwnLocal/goes)

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:13264** Graph database with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/google/cayley)   [![There was an update last week][Green]](https://github.com/google/cayley)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cayley)
    * [gokv](https://github.com/philippgille/gokv) **star:180** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![There was an update last week][Green]](https://github.com/philippgille/gokv)   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [cachego](https://github.com/fabiorphp/cachego) **star:122** Golang Cache component for multiple drivers.   [![There was an update last week][Green]](https://github.com/fabiorphp/cachego)   [![godoc][GoDoc]](https://godoc.org/github.com/fabiorphp/cachego)
    * [dsc](https://github.com/viant/dsc) **star:18** Datastore connectivity for SQL, NoSQL, structured files.   [![There was an update last week][Green]](https://github.com/viant/dsc)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2440** Now is a time toolkit for golang.   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [dateparse](https://github.com/araddon/dateparse) **star:993** Parse date's without knowing format in advance.   [![godoc][GoDoc]](https://godoc.org/github.com/araddon/dateparse)
* [carbon](https://github.com/uniplaces/carbon) **star:438** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/uniplaces/carbon)   [![godoc][GoDoc]](https://godoc.org/github.com/uniplaces/carbon)
* [durafmt](https://github.com/hako/durafmt) **star:287** Time duration formatting library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [timeutil](https://github.com/leekchan/timeutil) **star:176** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/leekchan/timeutil)   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/timeutil)
* [iso8601](https://github.com/relvacode/iso8601) **star:71** Efficiently parse ISO8601 date-times without regex.   [![It hasn't been updated in the last year][Yellow]](https://github.com/relvacode/iso8601)   [![godoc][GoDoc]](https://godoc.org/github.com/relvacode/iso8601)
* [timespan](https://github.com/SaidinWoT/timespan) **star:68** For interacting with intervals of time, defined as a start time and a duration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/SaidinWoT/timespan)   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:68** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/go-persian-calendar)
* [date](https://github.com/rickb777/date) **star:35** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.   [![godoc][GoDoc]](https://godoc.org/github.com/rickb777/date)
* [feiertage](https://github.com/wlbr/feiertage) **star:25** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:21** Calculate the sunrise and sunset times for a given location.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nathan-osman/go-sunrise)   [![godoc][GoDoc]](https://godoc.org/github.com/nathan-osman/go-sunrise)
* [kair](https://github.com/GuilhermeCaruso/kair) **star:15** Date and Time - Golang Formatting Library.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)
* [cronrange](https://github.com/1set/cronrange) **star:13** Parses Cron-style time range expressions, checks if the given time is within any ranges.   [![godoc][GoDoc]](https://godoc.org/github.com/1set/cronrange)
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:10** Nullable `time.Time`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/nulltime)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/nulltime)
* [tuesday](https://github.com/osteele/tuesday) **star:8** Ruby-compatible Strftime function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/osteele/tuesday)   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/tuesday)
* [strftime](https://github.com/awoodbeck/strftime) **star:4** C99-compatible strftime formatter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awoodbeck/strftime)   [![godoc][GoDoc]](https://godoc.org/github.com/awoodbeck/strftime)
* [go-str2duration](https://github.com/xhit/go-str2duration) **star:2** Convert string to duration. Support time.Duration returned string and more.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-str2duration)
* [go-week](https://github.com/stoewer/go-week) **star:2** An efficient package to work with ISO8601 week dates.   [![godoc][GoDoc]](https://godoc.org/github.com/stoewer/go-week)

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:16445** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [grpc-go](https://github.com/grpc/grpc-go) **star:10821** The Go language implementation of gRPC. HTTP/2 based RPC.   [![star > 2000][Awesome]](https://github.com/grpc/grpc-go)   [![There was an update last week][Green]](https://github.com/grpc/grpc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/grpc/grpc-go)
* [micro](https://github.com/micro/micro) **star:7737** Pluggable microservice toolkit and distributed systems platform.   [![star > 2000][Awesome]](https://github.com/micro/micro)   [![There was an update last week][Green]](https://github.com/micro/micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [NATS](https://github.com/nats-io/gnatsd) **star:7390** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.   [![star > 2000][Awesome]](https://github.com/nats-io/gnatsd)   [![There was an update last week][Green]](https://github.com/nats-io/gnatsd)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/gnatsd)
* [rpcx](https://github.com/smallnest/rpcx) **star:4457** Distributed pluggable RPC service framework like alibaba Dubbo.   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [tendermint](https://github.com/tendermint/tendermint) **star:3540** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   [![There was an update last week][Green]](https://github.com/tendermint/tendermint)   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [torrent](https://github.com/anacrolix/torrent) **star:3376** BitTorrent client package.   [![star > 2000][Awesome]](https://github.com/anacrolix/torrent)   [![There was an update last week][Green]](https://github.com/anacrolix/torrent)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/torrent)
* [raft](https://github.com/hashicorp/raft) **star:3298** Golang implementation of the Raft consensus protocol, by HashiCorp.   [![star > 2000][Awesome]](https://github.com/hashicorp/raft)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/raft)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [dragonboat](https://github.com/lni/dragonboat) **star:2916** A feature complete and high performance multi-group Raft library in Go.   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   [![There was an update last week][Green]](https://github.com/lni/dragonboat)   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   [![Contains Chinese documents][CN]](https://github.com/lni/dragonboat)
* [glow](https://github.com/chrislusf/glow) **star:2757** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [KrakenD](https://github.com/devopsfaith/krakend) **star:2453** Ultra performant API Gateway framework with middlewares.   [![star > 2000][Awesome]](https://github.com/devopsfaith/krakend)   [![godoc][GoDoc]](https://godoc.org/github.com/devopsfaith/krakend)
* [gleam](https://github.com/chrislusf/gleam) **star:2409** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2278** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [liftbridge](https://github.com/liftbridge-io/liftbridge) **star:1483** Lightweight, fault-tolerant message streams for NATS.   [![There was an update last week][Green]](https://github.com/liftbridge-io/liftbridge)   [![godoc][GoDoc]](https://godoc.org/github.com/liftbridge-io/liftbridge)
* [hprose](https://github.com/hprose/hprose-golang) **star:1083** Very newbility RPC Library, support 25+ languages now.   [![There was an update last week][Green]](https://github.com/hprose/hprose-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/hprose/hprose-golang)   [![Contains Chinese documents][CN]](https://github.com/hprose/hprose-golang)
* [ringpop-go](https://github.com/uber/ringpop-go) **star:611** Scalable, fault-tolerant application-layer sharding for Go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/uber/ringpop-go)   [![godoc][GoDoc]](https://godoc.org/github.com/uber/ringpop-go)
* [gorpc](https://github.com/valyala/gorpc) **star:579** Simple, fast and scalable RPC library for high load.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [go-health](https://github.com/InVisionApp/go-health) **star:534** Library for enabling asynchronous dependency health checks in your service.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [rain](https://github.com/cenkalti/rain) **star:482** BitTorrent client and library.   [![godoc][GoDoc]](https://godoc.org/github.com/cenkalti/rain)
* [digota](https://github.com/digota/digota) **star:325** grpc ecommerce microservice.   [![It hasn't been updated in the last year][Yellow]](https://github.com/digota/digota)   [![godoc][GoDoc]](https://godoc.org/github.com/digota/digota)
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [sleuth](https://github.com/ursiform/sleuth) **star:312** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ursiform/sleuth)   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) **star:291** A library built to provide support for defining async service health checks for golang services.   [![godoc][GoDoc]](https://godoc.org/github.com/AppsFlyer/go-sundheit)
* [go-jump](https://github.com/dgryski/go-jump) **star:281** Port of Google's "Jump" Consistent Hash function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgryski/go-jump)   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [consistent](https://github.com/buraksezer/consistent) **star:266** Consistent hashing with bounded loads.   [![godoc][GoDoc]](https://godoc.org/github.com/buraksezer/consistent)
* [dht](https://github.com/anacrolix/dht) **star:150** BitTorrent Kademlia DHT implementation.   [![There was an update last week][Green]](https://github.com/anacrolix/dht)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:125** The jsonrpc package helps implement of JSON-RPC 2.0.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:118** JSON-RPC 2.0 HTTP client implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [redis-lock](https://github.com/bsm/redislock) **star:99** Simplified distributed locking implementation using Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redislock)
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:60** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/svcavallar/celeriac.v1)   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:50** A revamped Google's jump consistent hash.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [drmaa](https://github.com/dgruber/drmaa) **star:29** Job submission library for cluster schedulers based on the DRMAA standard.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgruber/drmaa)   [![godoc][GoDoc]](https://godoc.org/github.com/dgruber/drmaa)
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [outboxer](https://github.com/italolelis/outboxer) **star:27** Outboxer is a go library that implements the outbox pattern.   [![There was an update last week][Green]](https://github.com/italolelis/outboxer)   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)
* [flowgraph](https://github.com/vectaport/flowgraph) **star:26** flow-based programming package.   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [go-pdu](https://github.com/pdupub/go-pdu) **star:12** A decentralized identity-based social network.   [![There was an update last week][Green]](https://github.com/pdupub/go-pdu)   [![godoc][GoDoc]](https://godoc.org/github.com/pdupub/go-pdu)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:10** A library for using DynamoDB as an atomic counter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylfin/dynatomic)   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)

## Dynamic DNS

*Tools for updating dynamic DNS records.*

* [GoDNS](https://github.com/timothyye/godns) **star:555** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)
* [DDNS](https://github.com/skibish/ddns) **star:136** Personal DDNS client with Digital Ocean Networking DNS as backend.   [![godoc][GoDoc]](https://godoc.org/github.com/skibish/ddns)
* [dyndns](https://gitlab.com/alcastle/dyndns)  Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:6084** Email and SMTP testing with web and API interface.   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [hermes](https://github.com/matcornic/hermes) **star:1905** Golang package that generates clean, responsive HTML e-mails.   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1266** A robust and flexible email library for Go.   [![There was an update last week][Green]](https://github.com/jordan-wright/email)   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:921** IMAP library for clients and servers.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:571** SendGrid's Go library for sending email.   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [mailgun-go](https://github.com/mailgun/mailgun-go) **star:443** Go library for sending mail with the Mailgun API.   [![godoc][GoDoc]](https://godoc.org/github.com/mailgun/mailgun-go)
* [Hectane](https://github.com/hectane/hectane) **star:181** Lightweight SMTP client providing an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [douceur](https://github.com/aymerick/douceur) **star:174** CSS inliner for your HTML emails.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/douceur)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [go-message](https://github.com/emersion/go-message) **star:143** Streaming library for the Internet Message Format and mail messages.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-message)
* [smtp](https://github.com/mailhog/smtp) **star:54** SMTP server protocol state machine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mailhog/smtp)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/smtp)
* [go-dkim](https://github.com/toorop/go-dkim) **star:52** DKIM library, to sign & verify email.   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:46** Inline styling for HTML mail in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [mailchain](https://github.com/mailchain/mailchain) **star:42** Send encrypted emails to blockchain addresses written in Go.   [![There was an update last week][Green]](https://github.com/mailchain/mailchain)   [![godoc][GoDoc]](https://godoc.org/github.com/mailchain/mailchain)
* [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:19** Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-simple-mail)

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) **star:5169** JavaScript interpreter written in Go.   [![star > 2000][Awesome]](https://github.com/robertkrimen/otto)   [![godoc][GoDoc]](https://godoc.org/github.com/robertkrimen/otto)
* [gopher-lua](https://github.com/yuin/gopher-lua) **star:3339** Lua 5.1 VM and compiler written in Go.   [![star > 2000][Awesome]](https://github.com/yuin/gopher-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/yuin/gopher-lua)
* [go-lua](https://github.com/Shopify/go-lua) **star:1793** Port of the Lua 5.2 VM to pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [tengo](https://github.com/d5/tengo) **star:1572** Bytecode compiled script language for Go.   [![There was an update last week][Green]](https://github.com/d5/tengo)   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [expr](https://github.com/antonmedv/expr) **star:1152** an engine that can evaluate expressions.   [![There was an update last week][Green]](https://github.com/antonmedv/expr)   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [go-python](https://github.com/sbinet/go-python) **star:1014** naive go bindings to the CPython C-API.   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [anko](https://github.com/mattn/anko) **star:983** Scriptable interpreter written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-php](https://github.com/deuill/go-php) **star:729** PHP bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/deuill/go-php)   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:686** Duktape JavaScript engine bindings for Go.   [![There was an update last week][Green]](https://github.com/olebedev/go-duktape)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [golua](https://github.com/aarzilli/golua) **star:469** Go bindings for Lua C API.
* [gisp](https://github.com/jcla1/gisp) **star:435** Simple LISP in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jcla1/gisp)   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [cel-go](https://github.com/google/cel-go) **star:380** Fast, portable, non-Turing complete expression evaluation with gradual typing.   [![godoc][GoDoc]](https://godoc.org/github.com/google/cel-go)
* [gval](https://github.com/PaesslerAG/gval) **star:184** A highly customizable expression language written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [gentee](https://github.com/gentee/gentee) **star:45** Embeddable scripting programming language.   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)
* [binder](https://github.com/alexeyco/binder) **star:35** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexeyco/binder)   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [purl](https://github.com/ian-kent/purl) **star:29** Perl 5.18.2 embedded in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/purl)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/purl)
* [ngaro](https://github.com/db47h/ngaro) **star:19** Embeddable Ngaro VM implementation enabling scripting in Retro.   [![It hasn't been updated in the last year][Yellow]](https://github.com/db47h/ngaro)   [![godoc][GoDoc]](https://godoc.org/github.com/db47h/ngaro)

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) **star:5597** Package that provides simple error handling primitives.   [![star > 2000][Awesome]](https://github.com/pkg/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/errors)
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:816** Go (golang) package for representing a list of errors as a single error.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [errorx](https://github.com/joomcode/errorx) **star:611** A feature rich error package with stack traces, composition of errors and more.   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:595** Golang errors with stack trace and source fragments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ztrue/tracerr)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [eris](https://github.com/rotisserie/eris) **star:580** A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.   [![There was an update last week][Green]](https://github.com/rotisserie/eris)   [![godoc][GoDoc]](https://godoc.org/github.com/rotisserie/eris)
* [errlog](https://github.com/snwfdhmp/errlog) **star:303** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   [![godoc][GoDoc]](https://godoc.org/github.com/snwfdhmp/errlog)
* [emperror](https://github.com/emperror/emperror) **star:115** Error handling tools and best practices for Go libraries and applications.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/emperror)
* [errors](https://github.com/emperror/errors) **star:43** Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/errors)
* [werr](https://github.com/txgruppi/werr) **star:13** Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/werr)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/werr)
* [errors](https://github.com/neuronlabs/errors) **star:3** Simple golang error handling with classification primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/neuronlabs/errors)
* [errors](https://github.com/PumpkinSeed/errors) **star:2** The most simple error wrapper with awesome performance and minimal memory overhead.   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/errors)
* [Falcon](https://github.com/SonicRoshan/falcon) **star:2** A Simple Yet Highly Powerful Package For Error Handling.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/falcon)

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:2616** FileSystem Abstraction System for Go.   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/pdfcpu/pdfcpu) **star:1349** PDF processor.   [![godoc][GoDoc]](https://godoc.org/github.com/pdfcpu/pdfcpu)
* [notify](https://github.com/rjeczalik/notify) **star:548** File system event notification library with simple API, similar to os/signal.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/notify)
* [copy](https://github.com/otiai10/copy) **star:143** Copy directory recursively.   [![There was an update last week][Green]](https://github.com/otiai10/copy)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/copy)
* [bigfile](https://github.com/bigfile/bigfile) **star:117** A file transfer system, support to manage files with http api, rpc call and ftp client.   [![godoc][GoDoc]](https://godoc.org/github.com/bigfile/bigfile)   [![Contains Chinese documents][CN]](https://github.com/bigfile/bigfile)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:64** Load csv file using tag.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-csv-tag)
* [opc](https://github.com/qmuntal/opc) **star:63** Load Open Packaging Conventions (OPC) files for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [skywalker](https://github.com/dixonwille/skywalker) **star:55** Package to allow one to concurrently go through a filesystem with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/skywalker)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/skywalker)
* [vfs](https://github.com/C2FO/vfs) **star:47** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   [![There was an update last week][Green]](https://github.com/C2FO/vfs)   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [afs](https://github.com/viant/afs) **star:43** Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.   [![There was an update last week][Green]](https://github.com/viant/afs)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/afs)
* [tarfs](https://github.com/posener/tarfs) **star:39** Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.   [![There was an update last week][Green]](https://github.com/posener/tarfs)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/tarfs)
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:23** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:21** Load gtfs files in go.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [gut/yos](https://github.com/1set/gut) **star:20** Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.   [![godoc][GoDoc]](https://godoc.org/github.com/1set/gut)
* [checksum](https://github.com/codingsince1985/checksum) **star:17** Compute message digest, like MD5 and SHA256, for large files.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [flop](https://github.com/homedepot/flop) **star:15** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:14** Copy files for humans.   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
* [parquet](https://github.com/parsyl/parquet) **star:6** Read and write [parquet](https://parquet.apache.org) files.   [![godoc][GoDoc]](https://godoc.org/github.com/parsyl/parquet)

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:1982** Arbitrary-precision fixed-point decimal numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-money](https://github.com/rhymond/go-money) **star:719** Implementation of Fowler's Money pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [accounting](https://github.com/leekchan/accounting) **star:543** money and currency formatting for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [go-finance](https://github.com/FlashBoys/go-finance) **star:539** Comprehensive financial markets data in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/FlashBoys/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/FlashBoys/go-finance)
* [techan](https://github.com/sdcoffey/techan) **star:238** Technical analysis library with advanced market analysis and trading strategies.   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [orderbook](https://github.com/i25959341/orderbook) **star:122** Matching Engine for Limit Order Book in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/i25959341/orderbook)
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:73** Query OFX servers and/or parse the responses (with example command-line client).   [![godoc][GoDoc]](https://godoc.org/github.com/aclindsa/ofxgo)
* [transaction](https://github.com/claygod/transaction) **star:65** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [vat](https://github.com/dannyvankooten/vat) **star:65** VAT number validation & EU VAT rates.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dannyvankooten/vat)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/vat)
* [go-finance](https://github.com/alpeb/go-finance) **star:56** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   [![godoc][GoDoc]](https://godoc.org/github.com/alpeb/go-finance)
* [go-finnhub](https://github.com/m1/go-finnhub) **star:21** Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-finnhub)
* [currency](https://github.com/bnkamalesh/currency) **star:19** High performant & accurate currency computation package.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)
* [go-finance](https://github.com/pieterclaerhout/go-finance) **star:3** Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-finance)

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) **star:1038** CSRF protection middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/nosurf)
* [binding](https://github.com/mholt/binding) **star:768** Binds form and JSON data from net/http Request to struct.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/binding)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:508** CSRF protection for Go web applications & services.   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [form](https://github.com/go-playground/form) **star:398** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [conform](https://github.com/leebenson/conform) **star:184** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/leebenson/conform)
* [formam](https://github.com/monoculum/formam) **star:139** decode form's values into a struct.   [![There was an update last week][Green]](https://github.com/monoculum/formam)   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:104** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/forms)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)
* [bind](https://github.com/robfig/bind) **star:24** Bind form data to any Go values.   [![It hasn't been updated in the last year][Yellow]](https://github.com/robfig/bind)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/bind)
* [queryparam](https://github.com/tomwright/queryparam) **star:2** Decode `url.Values` into usable struct values of standard or custom types.   [![godoc][GoDoc]](https://godoc.org/github.com/tomwright/queryparam)

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1130** Useful collection of helpfully functional Go collection utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tobyhede/go-underscore)   [![godoc][GoDoc]](https://godoc.org/github.com/tobyhede/go-underscore)
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:136** Monad, Functional Programming features for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TeaEntityLab/fpGo)   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:63** Functional Experiment in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)

## Game Development

*Awesome game development libraries.*

* [Leaf](https://github.com/name5566/leaf) **star:3439** Lightweight game server framework.   [![star > 2000][Awesome]](https://github.com/name5566/leaf)   [![godoc][GoDoc]](https://godoc.org/github.com/name5566/leaf)   [![Contains Chinese documents][CN]](https://github.com/name5566/leaf)
* [Pixel](https://github.com/faiface/pixel) **star:2793** Hand-crafted 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/faiface/pixel)   [![godoc][GoDoc]](https://godoc.org/github.com/faiface/pixel)
* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:2569** dead simple 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/hajimehoshi/ebiten)   [![There was an update last week][Green]](https://github.com/hajimehoshi/ebiten)   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/ebiten)
* [goworld](https://github.com/xiaonanln/goworld) **star:1425** Scalable game server engine, featuring space-entity framework and hot-swapping.   [![godoc][GoDoc]](https://godoc.org/github.com/xiaonanln/goworld)   [![Contains Chinese documents][CN]](https://github.com/xiaonanln/goworld)
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1290** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).   [![There was an update last week][Green]](https://github.com/veandco/go-sdl2)
* [nano](https://github.com/lonng/nano) **star:1212** Lightweight, facility, high performance golang based game server framework.   [![godoc][GoDoc]](https://godoc.org/github.com/lonng/nano)   [![Contains Chinese documents][CN]](https://github.com/lonng/nano)
* [engo](https://github.com/EngoEngine/engo) **star:1178** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [termloop](https://github.com/JoelOtter/termloop) **star:1109** Terminal-based game engine for Go, built on top of Termbox.   [![godoc][GoDoc]](https://godoc.org/github.com/JoelOtter/termloop)
* [gonet](https://github.com/xtaci/gonet) **star:1089** Game server skeleton implemented with golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xtaci/gonet)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [g3n](https://github.com/g3n/engine) **star:966** Go 3D Game Engine.   [![godoc][GoDoc]](https://godoc.org/github.com/g3n/engine)
* [Oak](https://github.com/oakmound/oak) **star:715** Pure Go game engine.   [![There was an update last week][Green]](https://github.com/oakmound/oak)   [![godoc][GoDoc]](https://godoc.org/github.com/oakmound/oak)
* [Pitaya](https://github.com/topfreegames/pitaya) **star:494** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   [![There was an update last week][Green]](https://github.com/topfreegames/pitaya)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/pitaya)
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:444** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [Azul3D](https://github.com/azul3d/engine) **star:443** 3D game engine written in Go.
* [go-astar](https://github.com/beefsack/go-astar) **star:358** Go implementation of the A\* path finding algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beefsack/go-astar)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-astar)
* [GarageEngine](https://github.com/vova616/GarageEngine) **star:319** 2d game engine written in Go working on OpenGL.   [![godoc][GoDoc]](https://godoc.org/github.com/vova616/GarageEngine)
* [go3d](https://github.com/ungerik/go3d) **star:174** Performance oriented 2D/3D math package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go3d)
* [glop](https://github.com/runningwild/glop) **star:77** Glop (Game Library Of Power) is a fairly simple cross-platform game library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/runningwild/glop)
* [go-collada](https://github.com/GlenKelley/go-collada) **star:15** Go package for working with the Collada file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GlenKelley/go-collada)   [![godoc][GoDoc]](https://godoc.org/github.com/GlenKelley/go-collada)

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:1997** .NET LINQ-like query methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
* [jennifer](https://github.com/dave/jennifer) **star:1470** Generate arbitrary Go code without templates.   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [gen](https://github.com/clipperhouse/gen) **star:1094** Code generation tool for ‘generics’-like functionality.   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [goderive](https://github.com/awalterschulze/goderive) **star:789** Derives functions from input types.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/goderive)
* [GoWrap](https://github.com/hexdigest/gowrap) **star:331** Generate decorators for Go interfaces using simple templates.   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gowrap)
* [interfaces](https://github.com/rjeczalik/interfaces) **star:208** Command line tool for generating interface definitions.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [go-enum](https://github.com/abice/go-enum) **star:101** Code generation for enums from code comments.   [![godoc][GoDoc]](https://godoc.org/github.com/abice/go-enum)
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:92** Go preprocessor for package scoped reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/pkgreflect)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/pkgreflect)
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:44** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   [![It hasn't been updated in the last year][Yellow]](https://github.com/t0pep0/efaceconv)   [![godoc][GoDoc]](https://godoc.org/github.com/t0pep0/efaceconv)
* [gotype](https://github.com/wzshiming/gotype) **star:30** Golang source code parsing, usage like reflect package.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/gotype)
* [generis](https://github.com/senselogic/GENERIS) **star:21** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
* [go-xray](https://github.com/pieterclaerhout/go-xray) **star:6** Helpers for making the use of reflection easier.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-xray)
* [typeregistry](https://github.com/xiaoxin01/typeregistry) **star:3** A library to create type dynamically.   [![godoc][GoDoc]](https://godoc.org/github.com/xiaoxin01/typeregistry)

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) **star:6729** Geolocation DB with spatial index and realtime geofencing.   [![star > 2000][Awesome]](https://github.com/tidwall/tile38)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/tile38)
* [S2 geometry](https://github.com/golang/geo) **star:1016** S2 geometry library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [mbtileserver](https://github.com/consbio/mbtileserver) **star:132** A simple Go-based server for map tiles stored in mbtiles format.   [![godoc][GoDoc]](https://godoc.org/github.com/consbio/mbtileserver)
* [geocache](https://github.com/melihmucuk/geocache) **star:119** In-memory cache that is suitable for geolocation based applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/melihmucuk/geocache)   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [osm](https://github.com/paulmach/osm) **star:101** Library for reading, writing and working with OpenStreetMap data and APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/osm)
* [WGS84](https://github.com/wroge/wgs84) **star:50** Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).   [![godoc][GoDoc]](https://godoc.org/github.com/wroge/wgs84)
* [geoserver](https://github.com/hishamkaram/geoserver) **star:33** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/geoserver)
* [gismanager](https://github.com/hishamkaram/gismanager) **star:26** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hishamkaram/gismanager)   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)
* [pbf](https://github.com/maguro/pbf) **star:20** OpenStreetMap PBF golang encoder/decoder.   [![godoc][GoDoc]](https://godoc.org/github.com/maguro/pbf)

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:9236** Compiler from Go to JavaScript.   [![star > 2000][Awesome]](https://github.com/gopherjs/gopherjs)   [![godoc][GoDoc]](https://godoc.org/github.com/gopherjs/gopherjs)
* [llgo](https://github.com/go-llvm/llgo) **star:1043** LLVM-based compiler for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-llvm/llgo)   [![godoc][GoDoc]](https://godoc.org/github.com/go-llvm/llgo)
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:396** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tardisgo/tardisgo)   [![godoc][GoDoc]](https://godoc.org/github.com/tardisgo/tardisgo)
* [c4go](https://github.com/Konstantin8105/c4go) **star:203** Transpile C code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:23** Transpile FORTRAN 77 code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) **star:3195** A high-performance and low-cost goroutine pool in Go.   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   [![There was an update last week][Green]](https://github.com/panjf2000/ants)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/ants)
* [goworker](https://github.com/benmanns/goworker) **star:2361** goworker is a Go-based background worker.   [![star > 2000][Awesome]](https://github.com/benmanns/goworker)   [![godoc][GoDoc]](https://godoc.org/github.com/benmanns/goworker)
* [tunny](https://github.com/Jeffail/tunny) **star:1575** Goroutine pool for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/tunny)
* [pool](https://github.com/go-playground/pool) **star:547** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [grpool](https://github.com/ivpusic/grpool) **star:543** Lightweight Goroutine pool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/grpool)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [workerpool](https://github.com/gammazero/workerpool) **star:271** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![There was an update last week][Green]](https://github.com/gammazero/workerpool)   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [go-floc](https://github.com/workanator/go-floc) **star:177** Orchestrate goroutines with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/workanator/go-floc)   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-floc)
* [gowp](https://github.com/xxjwxc/gowp) **star:169** gowp is concurrency limiting goroutine pool.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gowp)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gowp)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:127** Control goroutines execution order.   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [GoSlaves](https://github.com/themester/GoSlaves) **star:91** Simple and Asynchronous Goroutine pool library.   [![godoc][GoDoc]](https://godoc.org/github.com/themester/GoSlaves)
* [semaphore](https://github.com/kamilsk/semaphore) **star:82** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/semaphore)
* [semaphore](https://github.com/marusama/semaphore) **star:81** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).   [![It hasn't been updated in the last year][Yellow]](https://github.com/marusama/semaphore)   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/semaphore)
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:66** manages a resizeable pool of context-aware goroutines to bound concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
* [breaker](https://github.com/kamilsk/breaker) **star:56** Flexible mechanism to make execution flow interruptible.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/breaker)
* [worker-pool](https://github.com/vardius/worker-pool) **star:53** goworker is a Go simple async worker pool.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/worker-pool)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:44** CyclicBarrier for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/marusama/cyclicbarrier)   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/cyclicbarrier)
* [gollback](https://github.com/vardius/gollback) **star:32** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [async](https://github.com/studiosol/async) **star:30** A safe way to execute functions asynchronously, recovering them in case of panic.   [![godoc][GoDoc]](https://godoc.org/github.com/studiosol/async)
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:26** Run functions in parallel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/parallel-fn)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/parallel-fn)
* [threadpool](https://github.com/shettyh/threadpool) **star:26** Golang threadpool implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shettyh/threadpool)   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [artifex](https://github.com/borderstech/artifex) **star:25** Simple in-memory job queue for Golang using worker-based dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/artifex)
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [kyoo](https://github.com/dirkaholic/kyoo) **star:24** Provides an unlimited job queue and concurrent worker pools.   [![godoc][GoDoc]](https://godoc.org/github.com/dirkaholic/kyoo)
* [Hunch](https://github.com/AaronJan/Hunch) **star:23** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   [![godoc][GoDoc]](https://godoc.org/github.com/AaronJan/Hunch)
* [stl](https://github.com/ssgreg/stl) **star:11** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/stl)
* [gohive](https://github.com/loveleshsharma/gohive) **star:11** A highly performant and easy to use Goroutine pool for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/loveleshsharma/gohive)
* [routine](https://github.com/x-mod/routine) **star:9** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![godoc][GoDoc]](https://godoc.org/github.com/x-mod/routine)
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) **star:7** Like `sync.WaitGroup` with error handling and concurrency control.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)
* [conexec](https://github.com/ITcathyh/conexec) **star:5** A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.   [![godoc][GoDoc]](https://godoc.org/github.com/ITcathyh/conexec)
* [go-trylock](https://github.com/subchen/go-trylock) **star:5** TryLock support on read-write lock for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-trylock)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.   [![godoc][GoDoc]](https://godoc.org/github.com/nikhilsaraf/go-tools)
* [queue](https://github.com/AnikHasibul/queue) **star:4** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.   [![godoc][GoDoc]](https://godoc.org/github.com/AnikHasibul/queue)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [fyne](https://github.com/fyne-io/fyne) **star:9005** Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.   [![star > 2000][Awesome]](https://github.com/fyne-io/fyne)   [![There was an update last week][Green]](https://github.com/fyne-io/fyne)   [![godoc][GoDoc]](https://godoc.org/github.com/fyne-io/fyne)
* [ui](https://github.com/andlabs/ui) **star:7397** Platform-native GUI library for Go. Cross platform.   [![star > 2000][Awesome]](https://github.com/andlabs/ui)   [![godoc][GoDoc]](https://godoc.org/github.com/andlabs/ui)
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [qt](https://github.com/therecipe/qt) **star:7089** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   [![star > 2000][Awesome]](https://github.com/therecipe/qt)   [![godoc][GoDoc]](https://godoc.org/github.com/therecipe/qt)
* [webview](https://github.com/zserge/webview) **star:5591** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).   [![star > 2000][Awesome]](https://github.com/zserge/webview)
* [walk](https://github.com/lxn/walk) **star:4311** Windows application library kit for Go.   [![star > 2000][Awesome]](https://github.com/lxn/walk)   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [app](https://github.com/murlokswarm/app) **star:3470** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   [![star > 2000][Awesome]](https://github.com/murlokswarm/app)   [![There was an update last week][Green]](https://github.com/murlokswarm/app)   [![godoc][GoDoc]](https://godoc.org/github.com/murlokswarm/app)
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:3074** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1651** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/sciter-sdk/go-sciter)
* [gotk3](https://github.com/gotk3/gotk3) **star:945** Go bindings for GTK3.   [![There was an update last week][Green]](https://github.com/gotk3/gotk3)   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd) **star:243** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   [![godoc][GoDoc]](https://godoc.org/github.com/dtylman/gowd)

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) **star:4877** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   [![star > 2000][Awesome]](https://github.com/go-vgo/robotgo)   [![There was an update last week][Green]](https://github.com/go-vgo/robotgo)
* [systray](https://github.com/getlantern/systray) **star:1029** Cross platform Go library to place an icon and menu in the notification area.   [![There was an update last week][Green]](https://github.com/getlantern/systray)   [![godoc][GoDoc]](https://godoc.org/github.com/getlantern/systray)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:512** OSX Desktop Notifications library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/gosx-notifier)
* [trayhost](https://github.com/shurcooL/trayhost) **star:175** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:6** Go bindings for libappindicator3 C library.   [![godoc][GoDoc]](https://godoc.org/github.com/dawidd6/go-appindicator)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:5** OSX library to notify about any (pluggable) activity on your machine.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:2** OSX Sleep/Wake notifications in golang.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [gocv](https://github.com/hybridgroup/gocv) **star:3015** Go package for computer vision using OpenCV 3.3+.   [![star > 2000][Awesome]](https://github.com/hybridgroup/gocv)   [![There was an update last week][Green]](https://github.com/hybridgroup/gocv)   [![godoc][GoDoc]](https://godoc.org/github.com/hybridgroup/gocv)
* [imaging](https://github.com/disintegration/imaging) **star:2975** Simple Go image processing package.   [![star > 2000][Awesome]](https://github.com/disintegration/imaging)   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/imaging)
* [imaginary](https://github.com/h2non/imaginary) **star:2908** Fast and simple HTTP microservice for image resizing.   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [bild](https://github.com/anthonynsimon/bild) **star:2799** Collection of image processing algorithms in pure Go.   [![star > 2000][Awesome]](https://github.com/anthonynsimon/bild)   [![godoc][GoDoc]](https://godoc.org/github.com/anthonynsimon/bild)
* [ln](https://github.com/fogleman/ln) **star:2630** 3D line art rendering in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/ln)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/ln)
* [resize](https://github.com/nfnt/resize) **star:2302** Image resizing for Go with common interpolation methods.   [![star > 2000][Awesome]](https://github.com/nfnt/resize)   [![It hasn't been updated in the last year][Yellow]](https://github.com/nfnt/resize)   [![godoc][GoDoc]](https://godoc.org/github.com/nfnt/resize)
* [gg](https://github.com/fogleman/gg) **star:2178** 2D rendering in pure Go.   [![star > 2000][Awesome]](https://github.com/fogleman/gg)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/gg)
* [pt](https://github.com/fogleman/pt) **star:1838** Path tracing engine written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fogleman/pt)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/pt)
* [svgo](https://github.com/ajstarks/svgo) **star:1438** Go Language Library for SVG generation.   [![godoc][GoDoc]](https://godoc.org/github.com/ajstarks/svgo)
* [smartcrop](https://github.com/muesli/smartcrop) **star:1338** Finds good crops for arbitrary images and crop sizes.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/smartcrop)
* [gift](https://github.com/disintegration/gift) **star:1304** Package of image processing filters.   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/gift)
* [picfit](https://github.com/thoas/picfit) **star:1206** An image resizing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/picfit)
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1162** Go bindings for OpenCV.   [![godoc][GoDoc]](https://godoc.org/github.com/lazywei/go-opencv)
* [imagick](https://github.com/gographics/imagick) **star:1126** Go binding to ImageMagick's MagickWand C API.   [![godoc][GoDoc]](https://godoc.org/github.com/gographics/imagick)
* [geopattern](https://github.com/pravj/geopattern) **star:1055** Create beautiful generative image patterns from a string.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pravj/geopattern)   [![godoc][GoDoc]](https://godoc.org/github.com/pravj/geopattern)
* [bimg](https://github.com/h2non/bimg) **star:906** Small package for fast and efficient image processing using libvips.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:797** Go tool for LSB steganography, capable of hiding any file within an image.   [![There was an update last week][Green]](https://github.com/DimitarPetrov/stegify)   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [canvas](https://github.com/tdewolff/canvas) **star:419** Vector graphics to PDF, SVG or rasterized image.   [![There was an update last week][Green]](https://github.com/tdewolff/canvas)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/canvas)
* [mort](https://github.com/aldor007/mort) **star:388** Storage and image processing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aldor007/mort)
* [image2ascii](https://github.com/qeesung/image2ascii) **star:381** Convert image to ASCII.   [![It hasn't been updated in the last year][Yellow]](https://github.com/qeesung/image2ascii)   [![godoc][GoDoc]](https://godoc.org/github.com/qeesung/image2ascii)
* [govatar](https://github.com/o1egl/govatar) **star:347** Library and CMD tool for generating funny avatars.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/govatar)
* [go-nude](https://github.com/koyachi/go-nude) **star:300** Nudity detection with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/koyachi/go-nude)   [![godoc][GoDoc]](https://godoc.org/github.com/koyachi/go-nude)
* [goimagehash](https://github.com/corona10/goimagehash) **star:279** Go Perceptual image hashing package.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [rez](https://github.com/bamiaux/rez) **star:200** Image resizing in pure Go and SIMD.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bamiaux/rez)   [![godoc][GoDoc]](https://godoc.org/github.com/bamiaux/rez)
* [img](https://github.com/hawx/img) **star:133** Selection of image manipulation tools.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hawx/img)   [![godoc][GoDoc]](https://godoc.org/github.com/hawx/img)
* [darkroom](https://github.com/gojek/darkroom) **star:120** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [go-cairo](https://github.com/ungerik/go-cairo) **star:92** Go binding for the cairo graphics library.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [mergi](https://github.com/noelyahan/mergi) **star:89** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).   [![godoc][GoDoc]](https://godoc.org/github.com/noelyahan/mergi)
* [gltf](https://github.com/qmuntal/gltf) **star:62** Efficient and robust glTF 2.0 reader, writer and validator.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [go-gd](https://github.com/bolknote/go-gd) **star:52** Go binding for GD library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bolknote/go-gd)   [![godoc][GoDoc]](https://godoc.org/github.com/bolknote/go-gd)
* [steganography](https://github.com/auyer/steganography) **star:44** Pure Go Library for LSB steganography.   [![godoc][GoDoc]](https://godoc.org/github.com/auyer/steganography)
* [cameron](https://github.com/aofei/cameron) **star:41** An avatar generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)
* [goimghdr](https://github.com/corona10/goimghdr) **star:34** The imghdr module determines the type of image contained in a file for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimghdr)
* [tga](https://github.com/ftrvxmtrx/tga) **star:26** Package tga is a TARGA image format decoder/encoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ftrvxmtrx/tga)   [![godoc][GoDoc]](https://godoc.org/github.com/ftrvxmtrx/tga)
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:24** Port of webcolors library from Python to Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jyotiska/go-webcolors)   [![godoc][GoDoc]](https://godoc.org/github.com/jyotiska/go-webcolors)
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.   [![It hasn't been updated in the last year][Yellow]](https://github.com/donatj/mpo)   [![godoc][GoDoc]](https://godoc.org/github.com/donatj/mpo)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1377** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt) **star:884** Gatt is a Go package for building Bluetooth Low Energy peripherals.   [![godoc][GoDoc]](https://godoc.org/github.com/paypal/gatt)
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [mainflux](https://github.com/Mainflux/mainflux) **star:818** Industrial IoT Messaging and Device Management Server.   [![There was an update last week][Green]](https://github.com/Mainflux/mainflux)   [![godoc][GoDoc]](https://godoc.org/github.com/Mainflux/mainflux)
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [devices](https://github.com/goiot/devices) **star:229** Suite of libraries for IoT devices, experimental for x/exp/io.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goiot/devices)   [![godoc][GoDoc]](https://godoc.org/github.com/goiot/devices)
* [connectordb](https://github.com/connectordb/connectordb) **star:203** Open-Source Platform for Quantified Self & IoT.   [![There was an update last week][Green]](https://github.com/connectordb/connectordb)   [![godoc][GoDoc]](https://godoc.org/github.com/connectordb/connectordb)
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:191** Lightweight stream processing engine for IoT.   [![godoc][GoDoc]](https://godoc.org/github.com/sensorbee/sensorbee)
* [huego](https://github.com/amimof/huego) **star:136** An extensive Philips Hue client library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [eywa](https://github.com/xcodersun/eywa) **star:44** Project Eywa is essentially a connection manager that keeps track of connected devices.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xcodersun/eywa)   [![godoc][GoDoc]](https://godoc.org/github.com/xcodersun/eywa)

## Job Scheduler

*Libraries for scheduling jobs.*

* [gron](https://github.com/roylee0704/gron) **star:696** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   [![There was an update last week][Green]](https://github.com/roylee0704/gron)   [![godoc][GoDoc]](https://godoc.org/github.com/roylee0704/gron)
* [JobRunner](https://github.com/bamzi/jobrunner) **star:666** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   [![godoc][GoDoc]](https://godoc.org/github.com/bamzi/jobrunner)
* [jobs](https://github.com/albrow/jobs) **star:468** Persistent and flexible background jobs library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/jobs)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/jobs)
* [scheduler](https://github.com/carlescere/scheduler) **star:320** Cronjobs scheduling made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/scheduler)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/scheduler)
* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [go-cron](https://github.com/rk/go-cron) **star:185** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   [![godoc][GoDoc]](https://godoc.org/github.com/rk/go-cron)
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:61** Job scheduler that supports webhooks, crons and classic scheduling.   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)
* [clockwork](https://github.com/whiteShtef/clockwork) **star:1** Simple and intuitive job scheduling library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/whiteShtef/clockwork)

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) **star:6000** Get a JSON value with one line of code.   [![star > 2000][Awesome]](https://github.com/tidwall/gjson)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/gjson)
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2155** Automatically generate Go (golang) struct definitions from example JSON.   [![star > 2000][Awesome]](https://github.com/ChimeraCoder/gojson)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/gojson)
* [kazaam](https://github.com/Qntfy/kazaam) **star:146** API for arbitrary transformation of JSON documents.   [![godoc][GoDoc]](https://godoc.org/github.com/Qntfy/kazaam)
* [gojq](https://github.com/elgs/gojq) **star:143** JSON query in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/elgs/gojq)   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/gojq)
* [jsongo](https://github.com/ricardolonga/jsongo) **star:94** Fluent API to make it easier to create Json objects.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ricardolonga/jsongo)   [![godoc][GoDoc]](https://godoc.org/github.com/ricardolonga/jsongo)
* [jettison](https://github.com/wI2L/jettison) **star:74** High performance, reflection-less JSON encoder for Go.   [![There was an update last week][Green]](https://github.com/wI2L/jettison)   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jettison)
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [gjo](https://github.com/skanehira/gjo) **star:71** Small utility to create JSON objects.   [![godoc][GoDoc]](https://godoc.org/github.com/skanehira/gjo)
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:63** JSON diff utility written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [jsonf](https://github.com/miolini/jsonf) **star:55** Console tool for highlighted formatting and struct query fetching JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/miolini/jsonf)   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/jsonf)
* [mp](https://github.com/sanbornm/mp) **star:38** Simple cli email parser. It currently takes stdin and outputs JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sanbornm/mp)   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/mp)
* [json2go](https://github.com/m-zajac/json2go) **star:33** Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.   [![godoc][GoDoc]](https://godoc.org/github.com/m-zajac/json2go)
* [go-respond](https://github.com/nicklaw5/go-respond) **star:30** Go package for handling common HTTP JSON responses.   [![godoc][GoDoc]](https://godoc.org/github.com/nicklaw5/go-respond)
* [ajson](https://github.com/spyzhov/ajson) **star:26** Abstract JSON for golang with JSONPath support.   [![godoc][GoDoc]](https://godoc.org/github.com/spyzhov/ajson)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/RichardKnop/jsonhal)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:9** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.   [![godoc][GoDoc]](https://godoc.org/github.com/ddymko/go-jsonerror)
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:8** Go bindings based on the JSON API errors reference.   [![It hasn't been updated in the last year][Yellow]](https://github.com/AmuzaTkts/jsonapi-errors)   [![godoc][GoDoc]](https://godoc.org/github.com/AmuzaTkts/jsonapi-errors)
* [ej](https://github.com/lucassscaravelli/ej) **star:3** Write and read JSON from different sources succinctly.   [![godoc][GoDoc]](https://godoc.org/github.com/lucassscaravelli/ej)
* [mapslice-json](https://github.com/mickep76/mapslice-json) **star:1** Go MapSlice for ordered marshal/ unmarshal of maps in JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/mapslice-json)

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) **star:14179** Structured logger for Go.   [![star > 2000][Awesome]](https://github.com/Sirupsen/logrus)   [![There was an update last week][Green]](https://github.com/Sirupsen/logrus)   [![godoc][GoDoc]](https://godoc.org/github.com/Sirupsen/logrus)
* [zap](https://github.com/uber-go/zap) **star:9197** Fast, structured, leveled logging in Go.   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![There was an update last week][Green]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [spew](https://github.com/davecgh/go-spew) **star:3770** Implements a deep pretty printer for Go data structures to aid in debugging.   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [zerolog](https://github.com/rs/zerolog) **star:3031** Zero-allocation JSON logger.   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![There was an update last week][Green]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [glog](https://github.com/golang/glog) **star:2488** Leveled execution logs for Go.   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:1814** Simple rolling logger, implements io.WriteCloser.   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [tail](https://github.com/hpcloud/tail) **star:1736** Go package striving to emulate the features of the BSD tail program.   [![There was an update last week][Green]](https://github.com/hpcloud/tail)   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [seelog](https://github.com/cihub/seelog) **star:1443** Logging functionality with flexible dispatching, filtering, and formatting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cihub/seelog)   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log15](https://github.com/inconshreveable/log15) **star:952** Simple, powerful logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/log15)
* [log](https://github.com/apex/log) **star:811** Structured logging package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [onelog](https://github.com/francoispqt/onelog) **star:364** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/francoispqt/onelog)   [![godoc][GoDoc]](https://godoc.org/github.com/francoispqt/onelog)
* [logxi](https://github.com/mgutz/logxi) **star:343** 12-factor app logger that is fast and makes you happy.   [![godoc][GoDoc]](https://godoc.org/github.com/mgutz/logxi)
* [log](https://github.com/go-playground/log) **star:273** Simple, configurable and scalable Structured Logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [logutils](https://github.com/hashicorp/logutils) **star:272** Utilities for slightly better logging in Go (Golang) extending the standard logger.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [go-logger](https://github.com/apsdehal/go-logger) **star:249** Simple logger of Go Programs, with level handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/apsdehal/go-logger)
* [logger](https://github.com/azer/logger) **star:139** Minimalistic logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/azer/logger)
* [xlog](https://github.com/rs/xlog) **star:133** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:132** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:112** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-ozzo/ozzo-log)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-log)   [![Contains Chinese documents][CN]](https://github.com/go-ozzo/ozzo-log)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:84** Full-featured logging saas written in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/logvoyage)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [glg](https://github.com/kpango/glg) **star:68** glg is simple and fast leveled logging library for Go.   [![There was an update last week][Green]](https://github.com/kpango/glg)   [![godoc][GoDoc]](https://godoc.org/github.com/kpango/glg)
* [stdlog](https://github.com/alexcesaro/log) **star:42** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexcesaro/log)   [![godoc][GoDoc]](https://godoc.org/github.com/alexcesaro/log)
* [gologger](https://github.com/sadlil/gologger) **star:38** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/gologger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/gologger)   [![Archived][Archived]](https://github.com/sadlil/gologger)
* [logex](https://github.com/chzyer/logex) **star:35** Golang log lib, supports tracking and level, wrap by standard log lib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chzyer/logex)   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/logex)
* [go-log](https://github.com/ian-kent/go-log) **star:34** Log4j implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/go-log)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:27** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [go-log](https://github.com/siddontang/go-log) **star:26** Log lib supports level and multi handlers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/siddontang/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-log)
* [logrusly](https://github.com/sebest/logrusly) **star:25** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/sebest/logrusly)   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/logrusly)
* [log](https://github.com/teris-io/log) **star:23** Structured log interface for Go cleanly separates logging facade from its implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/log)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/log)
* [distillog](https://github.com/amoghe/distillog) **star:22** distilled levelled logging (think of it as stdlib + log levels).   [![It hasn't been updated in the last year][Yellow]](https://github.com/amoghe/distillog)   [![godoc][GoDoc]](https://godoc.org/github.com/amoghe/distillog)
* [journald](https://github.com/ssgreg/journald) **star:22** Go implementation of systemd Journal's native API for logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ssgreg/journald)   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/journald)
* [mlog](https://github.com/jbrodriguez/mlog) **star:19** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jbrodriguez/mlog)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrodriguez/mlog)
* [gomol](https://github.com/aphistic/gomol) **star:16** Multiple-output, structured logging for Go with extensible logging outputs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aphistic/gomol)   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [go-log](https://github.com/subchen/go-log) **star:10** Simple and configurable Logging in Go, with level, formatters and writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/subchen/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-log)
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ewwwwwqm/logdump)   [![godoc][GoDoc]](https://godoc.org/github.com/ewwwwwqm/logdump)
* [glo](https://github.com/lajosbencz/glo) **star:9** PHP Monolog inspired logging facility with identical severity levels.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lajosbencz/glo)   [![godoc][GoDoc]](https://godoc.org/github.com/lajosbencz/glo)
* [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:7** `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/logrusiowriter)
* [logmatic](https://github.com/borderstech/logmatic) **star:6** Colorized logger for Golang with dynamic log level configuration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/borderstech/logmatic)   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [log](https://github.com/aerogo/log) **star:6** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/log)
* [xlog](https://github.com/xfxdev/xlog) **star:6** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xfxdev/xlog)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [logo](https://github.com/mbndr/logo) **star:5** Golang logger to different configurable writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mbndr/logo)   [![godoc][GoDoc]](https://godoc.org/github.com/mbndr/logo)
* [go-log](https://github.com/pieterclaerhout/go-log) **star:3** A logging library with strack traces, object dumping and optional timestamps.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-log)

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:7097** General Machine Learning library for Go.   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   [![Contains Chinese documents][CN]](https://github.com/sjwhitworth/golearn)
* [gorgonia](https://github.com/gorgonia/gorgonia) **star:3203** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![star > 2000][Awesome]](https://github.com/gorgonia/gorgonia)   [![godoc][GoDoc]](https://godoc.org/github.com/gorgonia/gorgonia)
* [tfgo](https://github.com/galeone/tfgo) **star:1325** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/tfgo)
* [goml](https://github.com/cdipaolo/goml) **star:1084** On-line Machine Learning in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gosseract](https://github.com/otiai10/gosseract) **star:1084** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [gorse](https://github.com/zhenghaoz/gorse) **star:935** An offline recommender system backend based on collaborative filtering written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   [![Contains Chinese documents][CN]](https://github.com/zhenghaoz/gorse)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:667** An evolutionary optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:667** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ryanbressler/CloudForest)   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [bayesian](https://github.com/jbrukh/bayesian) **star:655** Naive Bayesian Classification for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jbrukh/bayesian)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [gobrain](https://github.com/goml/gobrain) **star:433** Neural Networks written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [ocrserver](https://github.com/otiai10/ocrserver) **star:274** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   [![It hasn't been updated in the last year][Yellow]](https://github.com/otiai10/ocrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/ocrserver)
* [regommend](https://github.com/muesli/regommend) **star:266** Recommendation & collaborative filtering engine.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [go-deep](https://github.com/patrikeh/go-deep) **star:256** A feature-rich neural network library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/patrikeh/go-deep)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:250** Go Interface to Open Neural Network Exchange (ONNX).   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-galib](https://github.com/thoj/go-galib) **star:177** Genetic Algorithms library written in Go / golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thoj/go-galib)   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:156** Recommendation Algorithms library written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timkaye11/goRecommend)   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [shield](https://github.com/eaigner/shield) **star:129** Bayesian text classifier with flexible tokenizers and storage backends for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [Goptuna](https://github.com/c-bata/goptuna) **star:119** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![There was an update last week][Green]](https://github.com/c-bata/goptuna)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [go-fann](https://github.com/white-pony/go-fann) **star:103** Go bindings for Fast Artificial Neural Networks(FANN) library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/white-pony/go-fann)   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [goga](https://github.com/tomcraven/goga) **star:86** Genetic algorithm library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomcraven/goga)   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:64** libsvm golang version derived work based on LIBSVM 3.14.   [![It hasn't been updated in the last year][Yellow]](https://github.com/datastream/libsvm)   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [neural-go](https://github.com/schuyler/neural-go) **star:61** Multilayer perceptron network implemented in Go, with training via backpropagation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schuyler/neural-go)   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [go-pr](https://github.com/daviddengcn/go-pr) **star:59** Pattern recognition package in Go lang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-pr)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [gonet](https://github.com/dathoangnd/gonet) **star:58** Neural Network for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/dathoangnd/gonet)
* [neat](https://github.com/jinyeom/neat) **star:56** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/jinyeom/neat)   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   [![Archived][Archived]](https://github.com/jinyeom/neat)
* [goscore](https://github.com/asafschers/goscore) **star:45** Go Scoring API for PMML.   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [golinear](https://github.com/danieldk/golinear) **star:39** liblinear bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/golinear)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [fonet](https://github.com/Fontinalis/fonet) **star:36** A Deep Neural Network library written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Fontinalis/fonet)
* [Varis](https://github.com/Xamber/Varis) **star:28** Golang Neural Network.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Xamber/Varis)   [![godoc][GoDoc]](https://godoc.org/github.com/Xamber/Varis)
* [godist](https://github.com/e-dard/godist) **star:25** Various probability distributions, and associated methods.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-dard/godist)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/godist)
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:22** Go implementation of the k-modes and k-prototypes clustering algorithms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
* [probab](https://github.com/ThePaw/probab) **star:12** Probability distribution functions. Bayesian inference. Written in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/probab)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/probab)
* [evoli](https://github.com/khezen/evoli) **star:9** Genetic Algorithm and Particle Swarm Optimization library.   [![There was an update last week][Green]](https://github.com/khezen/evoli)   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)
* [GoMind](https://github.com/surenderthakran/gomind) **star:7** A simplistic Neural Network Library in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surenderthakran/gomind)   [![godoc][GoDoc]](https://godoc.org/github.com/surenderthakran/gomind)
* [randomforest](https://github.com/malaschitz/randomForest) **star:3** Easy to use Random Forest library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/malaschitz/randomForest)

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) **star:5502** Go library for Apache Kafka.   [![star > 2000][Awesome]](https://github.com/Shopify/sarama)   [![There was an update last week][Green]](https://github.com/Shopify/sarama)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/sarama)
* [gorush](https://github.com/appleboy/gorush) **star:4246** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![There was an update last week][Green]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:4133** Real-time messaging (Websockets or SockJS) server in Go.   [![star > 2000][Awesome]](https://github.com/centrifugal/centrifugo)   [![There was an update last week][Green]](https://github.com/centrifugal/centrifugo)   [![godoc][GoDoc]](https://godoc.org/github.com/centrifugal/centrifugo)
* [machinery](https://github.com/RichardKnop/machinery) **star:3894** Asynchronous task queue/job queue based on distributed message passing.   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:3300** socket.io library for golang, a realtime application framework.   [![star > 2000][Awesome]](https://github.com/googollee/go-socket.io)   [![godoc][GoDoc]](https://godoc.org/github.com/googollee/go-socket.io)
* [NATS Go Client](https://github.com/nats-io/nats) **star:2710** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.   [![star > 2000][Awesome]](https://github.com/nats-io/nats)   [![There was an update last week][Green]](https://github.com/nats-io/nats)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/nats)
* [APNs2](https://github.com/sideshow/apns2) **star:2270** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.   [![star > 2000][Awesome]](https://github.com/sideshow/apns2)   [![godoc][GoDoc]](https://godoc.org/github.com/sideshow/apns2)
* [Benthos](https://github.com/Jeffail/benthos) **star:2251** A message streaming bridge between a range of protocols.   [![star > 2000][Awesome]](https://github.com/Jeffail/benthos)   [![There was an update last week][Green]](https://github.com/Jeffail/benthos)   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/benthos)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1895** gopush-cluster is a go push server cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Terry-Mao/gopush-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   [![Contains Chinese documents][CN]](https://github.com/Terry-Mao/gopush-cluster)
* [Mercure](https://github.com/dunglas/mercure) **star:1890** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [melody](https://github.com/olahol/melody) **star:1763** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1603** the official Go package for NSQ.   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [mangos](https://github.com/go-mangos/mangos) **star:1538** Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mangos/mangos)   [![Archived][Archived]](https://github.com/go-mangos/mangos)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1143** Redis backed unified push service for server-side notifications to mobile devices.   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [zmq4](https://github.com/pebbe/zmq4) **star:837** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/zmq4)
* [Gollum](https://github.com/trivago/gollum) **star:835** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [Beaver](https://github.com/Clivern/Beaver) **star:814** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![There was an update last week][Green]](https://github.com/Clivern/Beaver)   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [EventBus](https://github.com/asaskevich/EventBus) **star:656** The lightweight event bus with async compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [golongpoll](https://github.com/jcuga/golongpoll) **star:447** HTTP longpoll server library that makes web pub-sub simple.   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [dbus](https://github.com/godbus/dbus) **star:430** Native Go bindings for D-Bus.   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [Asynq](https://github.com/hibiken/asynq) **star:420** A simple, reliable, and efficient distributed task queue for Go built on top of Redis.   [![There was an update last week][Green]](https://github.com/hibiken/asynq)   [![godoc][GoDoc]](https://godoc.org/github.com/hibiken/asynq)
* [emitter](https://github.com/olebedev/emitter) **star:344** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [Glue](https://github.com/desertbit/glue) **star:332** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:310** Simple pubsub package for go.   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [guble](https://github.com/smancke/guble) **star:143** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   [![It hasn't been updated in the last year][Yellow]](https://github.com/smancke/guble)   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [Bus](https://github.com/mustafaturan/bus) **star:142** Minimalist message bus implementation for internal communication.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [rabtap](https://github.com/jandelgado/rabtap) **star:111** RabbitMQ swiss army knife cli app.   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [oplog](https://github.com/dailymotion/oplog) **star:102** Generic oplog/replication system for REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dailymotion/oplog)   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [messagebus](https://github.com/vardius/message-bus) **star:99** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.   [![There was an update last week][Green]](https://github.com/vardius/message-bus)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:69** A tiny wrapper over amqp exchanges and queues.   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [drone-line](https://github.com/appleboy/drone-line) **star:67** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-line)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:58** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sybrexsys/RapidMQ)   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:57** A tiny wrapper around NSQ topic and channel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:48** Native implementation of the freedesktop notification spec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TheCreeper/go-notify)   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [hub](https://github.com/leandro-lugaresi/hub) **star:48** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [Commander](https://github.com/jeroenrinzema/commander) **star:40** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [event](https://github.com/agoalofalife/event) **star:33** Implementation of the pattern observer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agoalofalife/event)   [![godoc][GoDoc]](https://godoc.org/github.com/agoalofalife/event)
* [go-res](https://github.com/jirenius/go-res) **star:24** Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate.   [![godoc][GoDoc]](https://godoc.org/github.com/jirenius/go-res)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:13** Client library to Viessmann Vitotrol web service.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:13** redisqueue provides a producer and consumer of a queue that uses Redis streams.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/redisqueue)
* [jazz](https://github.com/socifi/jazz) **star:9** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/socifi/jazz)   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:8** Gaurun Client written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gaurun-client)
* [rmqconn](https://github.com/sbabiv/rmqconn) **star:6** RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/rmqconn)
* [ami](https://github.com/kak-tus/ami) **star:6** Go client to reliable queues based on Redis Cluster Streams.   [![godoc][GoDoc]](https://godoc.org/github.com/kak-tus/ami)

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:2150** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![star > 2000][Awesome]](https://github.com/unidoc/unioffice)   [![There was an update last week][Green]](https://github.com/unidoc/unioffice)   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:5833** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   [![There was an update last week][Green]](https://github.com/360EntSecGroup-Skylar/excelize)   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:3954** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:106** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:71** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:15** Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fterrag/goxlsxwriter)   [![godoc][GoDoc]](https://godoc.org/github.com/fterrag/goxlsxwriter)

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:1304** A reflection based dependency injection toolkit for Go.   [![There was an update last week][Green]](https://github.com/uber-go/dig)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:1195** A dependency injection based application framework for Go (built on top of dig).   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [container](https://github.com/golobby/container) **star:77** A powerful IoC Container with fluent and easy-to-use interface.   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/container)
* [dingo](https://github.com/i-love-flamingo/dingo) **star:41** A dependency injection toolkit for Go, based on Guice.   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/dingo)
* [alice](https://github.com/magic003/alice) **star:38** Additive dependency injection container for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/magic003/alice)   [![godoc][GoDoc]](https://godoc.org/github.com/magic003/alice)
* [wire](https://github.com/Fs02/wire) **star:29** Strict Runtime Dependency Injection for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/wire)
* [linker](https://github.com/logrange/linker) **star:21** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [gocontainer](https://github.com/vardius/gocontainer) **star:12** Simple Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)
* [di](https://github.com/goava/di) **star:7** A dependency injection container for go programming language.   [![There was an update last week][Green]](https://github.com/goava/di)   [![godoc][GoDoc]](https://godoc.org/github.com/goava/di)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:13699** Set of common historical and emerging project layout patterns in the Go ecosystem.   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:566** Go application boilerplate and example applying modern practices.   [![godoc][GoDoc]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) **star:326** A Go application boilerplate template for quick starting projects following production best practices.   [![godoc][GoDoc]](https://godoc.org/github.com/lacion/cookiecutter-golang)
* [scaffold](https://github.com/catchplay/scaffold) **star:58** Scaffold generates starter Go project layout. Lets you focus on business logic implemeted.   [![It hasn't been updated in the last year][Yellow]](https://github.com/catchplay/scaffold)   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [go-sample](https://github.com/zitryss/go-sample) **star:46** A sample layout for Go application projects with the real code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zitryss/go-sample)   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:721** Collection of useful string functions ported from other languages.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:86** String utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:4654** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![There was an update last week][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:2757** Library and command for making and extracting .zip and .tar.gz archives.   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![There was an update last week][Green]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gosms](https://github.com/haxpax/gosms) **star:1273** Your own local SMS gateway in Go that can be used to send SMS.   [![It hasn't been updated in the last year][Yellow]](https://github.com/haxpax/gosms)   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:984** Resiliency patterns for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:845** Random data generator written in go.   [![There was an update last week][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:785** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![There was an update last week][Green]](https://github.com/mojocn/base64Captcha)   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![Contains Chinese documents][CN]](https://github.com/mojocn/base64Captcha)
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:775** Generic object pool for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![Contains Chinese documents][CN]](https://github.com/jolestar/go-commons-pool)
* [shortid](https://github.com/teris-io/shortid) **star:533** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/shortid)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:503** Library for interacting with LLVM IR in pure Go.   [![There was an update last week][Green]](https://github.com/llir/llvm)   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:380** Easy to use, extensible health check library.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:350** Package conv provides fast and intuitive conversions across Go types.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [banner](https://github.com/dimiro1/banner) **star:256** Add beautiful banners into your Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:235** Package that exposes country and subdivision data.   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [antch](https://github.com/antchfx/antch) **star:169** A fast, powerful and extensible web crawling & scraping framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/antch)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   [![Contains Chinese documents][CN]](https://github.com/antchfx/antch)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:155** Beautify data display for Humans.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   [![Contains Chinese documents][CN]](https://github.com/go-ffmt/ffmt)
* [stateless](https://github.com/qmuntal/stateless) **star:148** A fluent library for creating state machines.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [ghorg](https://github.com/gabrie30/ghorg) **star:146** Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [lk](https://github.com/hyperboloide/lk) **star:145** A simple licensing library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [battery](https://github.com/distatus/battery) **star:143** Cross-platform, normalized battery information library.   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [stats](https://github.com/go-playground/stats) **star:131** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [bitio](https://github.com/icza/bitio) **star:116** Highly optimized bit-level Reader and Writer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:107** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [turtle](https://github.com/hackebrot/turtle) **star:96** Emojis for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:87** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [gommit](https://github.com/antham/gommit) **star:85** Analyze git commit messages to ensure they follow defined patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [gotoprom](https://github.com/cabify/gotoprom) **star:80** Type-safe metrics builder wrapper library for the official Prometheus client.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [indigo](https://github.com/osamingo/indigo) **star:58** Distributed unique ID generator of using Sonyflake and encoded by Base58.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [morse](https://github.com/alwindoss/morse) **star:57** Library to convert to and from morse code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alwindoss/morse)   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [captcha](https://github.com/steambap/captcha) **star:54** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [xkg](https://github.com/go-xkg/xkg) **star:48** X Keyboard Grabber.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [persian](https://github.com/mavihq/persian) **star:40** Some utilities for Persian language in go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mavihq/persian)   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:39** HTTP service to generate PDF from Json requests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [datacounter](https://github.com/miolini/datacounter) **star:31** Go counters for readers/writer/http.ResponseWriter.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:31** GoLang Library for [Browser Capabilities Project](http://browscap.org/).   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [autoflags](https://github.com/artyom/autoflags) **star:26** Go package to automatically define command line flags from struct fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/artyom/autoflags)   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [xdg](https://github.com/rkoesters/xdg) **star:21** FreeDesktop.org (xdg) Specs implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [gosh](https://github.com/osamingo/gosh) **star:20** Provide Go Statistics Handler, Struct, Measure Method.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:19** A modern, powerful, and robust URL shortener microservice with mysql support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [sandid](https://github.com/aofei/sandid) **star:15** Every grain of sand on earth has its own ID.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [anagent](https://github.com/mudler/anagent) **star:11** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mudler/anagent)   [![godoc][GoDoc]](https://godoc.org/github.com/mudler/anagent)
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:10** Calculate average score and rating based on Wilson Score Equation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/avgRating)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/avgRating)
* [hostutils](https://github.com/Wing924/hostutils) **star:8** A golang library for packing and unpacking FQDNs list.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/hostutils)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [shellwords](https://github.com/Wing924/shellwords) **star:8** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [metrics](https://github.com/pascaldekloe/metrics) **star:6** Library for metrics instrumentation and Prometheus exposition.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [numa](https://github.com/lrita/numa) **star:2** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2453** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only.   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1221** Go efficient text segmentation; support english, chinese, japanese and other.   [![There was an update last week][Green]](https://github.com/go-ego/gse)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![Contains Chinese documents][CN]](https://github.com/go-ego/gse)
* [when](https://github.com/olebedev/when) **star:1064** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:1000** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![Contains Chinese documents][CN]](https://github.com/yanyiwu/gojieba)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:661** CN Hanzi to Hanyu Pinyin converter.   [![There was an update last week][Green]](https://github.com/mozillazg/go-pinyin)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [kagome](https://github.com/ikawaha/kagome) **star:474** JP morphological analyzer written in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:398** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![It hasn't been updated in the last year][Yellow]](https://github.com/abadojack/whatlanggo)   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:363** Extract values from strings and fill your structs with nlp.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:268** Sentence tokenizer:  converts text into a list of sentences.   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [nlp](https://github.com/james-bowman/nlp) **star:246** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [getlang](https://github.com/rylans/getlang) **star:86** Fast natural language detection package.   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-nlp](https://github.com/nuance/go-nlp) **star:82** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)
* [gounidecode](https://github.com/fiam/gounidecode) **star:68** Unicode transliterator (also known as unidecode) for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:67** ASCII transliterations of Unicode text.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [textcat](https://github.com/pebbe/textcat) **star:61** Go package for n-gram based text categorization, with support for utf-8 and raw text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:58** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [go-stem](https://github.com/agonopol/go-stem) **star:56** Implementation of the porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [RAKE.go](https://github.com/Obaied/RAKE.go) **star:56** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][GoDoc]](https://godoc.org/github.com/Obaied/RAKE.go)
* [stemmer](https://github.com/dchest/stemmer) **star:49** Stemmer packages for Go programming language. Includes English and German stemmers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [segment](https://github.com/blevesearch/segment) **star:48** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![It hasn't been updated in the last year][Yellow]](https://github.com/blevesearch/segment)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [porter2](https://github.com/zhenjl/porter2) **star:37** Really fast Porter 2 stemmer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/porter2)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [go2vec](https://github.com/danieldk/go2vec) **star:34** Reader and utility functions for word2vec embeddings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [petrovich](https://github.com/striker2000/petrovich) **star:29** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang implementation of the Paice/Husk Stemming Algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [snowball](https://github.com/goodsign/snowball) **star:25** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/snowball)
* [go-mystem](https://github.com/dveselov/mystem) **star:23** CGo bindings to Yandex.Mystem - russian morphology analyzer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:17** Go bindings for the snowball libstemmer library including porter 2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [shamoji](https://github.com/osamingo/shamoji) **star:11** The shamoji is word filtering package written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:7** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)
* [go-localize](https://github.com/m1/go-localize) **star:4** Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-localize)
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:1** Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.   [![godoc][GoDoc]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)

## Networking

*Libraries for working with various layers of the network.*

* [fasthttp](https://github.com/valyala/fasthttp) **star:11737** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![There was an update last week][Green]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [kcptun](https://github.com/xtaci/kcptun) **star:11611** Extremely simple & fast udp tunnel based on KCP protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcptun)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcptun)
* [dns](https://github.com/miekg/dns) **star:4389** Go library for working with DNS.   [![star > 2000][Awesome]](https://github.com/miekg/dns)   [![There was an update last week][Green]](https://github.com/miekg/dns)   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:3853** An implementation of the QUIC protocol in pure Go.   [![star > 2000][Awesome]](https://github.com/lucas-clemente/quic-go)   [![There was an update last week][Green]](https://github.com/lucas-clemente/quic-go)   [![godoc][GoDoc]](https://godoc.org/github.com/lucas-clemente/quic-go)
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3535** HTTPLabs let you inspect HTTP requests and forge responses.   [![star > 2000][Awesome]](https://github.com/gchaincl/httplab)   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/httplab)
* [gopacket](https://github.com/google/gopacket) **star:3300** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   [![There was an update last week][Green]](https://github.com/google/gopacket)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [webrtc](https://github.com/pions/webrtc) **star:3239** A pure Go implementation of the WebRTC API.   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   [![There was an update last week][Green]](https://github.com/pions/webrtc)   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2495** KCP - Fast and Reliable ARQ Protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   [![There was an update last week][Green]](https://github.com/xtaci/kcp-go)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [gobgp](https://github.com/osrg/gobgp) **star:1827** BGP implemented in the Go Programming Language.   [![There was an update last week][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [gnet](https://github.com/panjf2000/gnet) **star:1763** `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.   [![There was an update last week][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/gnet)
* [ssh](https://github.com/gliderlabs/ssh) **star:1543** Higher-level API for building SSH servers (wraps crypto/ssh).   [![godoc][GoDoc]](https://godoc.org/github.com/gliderlabs/ssh)
* [fortio](https://github.com/fortio/fortio) **star:1184** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![There was an update last week][Green]](https://github.com/fortio/fortio)   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:981** Simple TUN/TAP library.   [![There was an update last week][Green]](https://github.com/songgao/water)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [go-getter](https://github.com/hashicorp/go-getter) **star:867** Go library for downloading files or directories from various sources using a URL.   [![There was an update last week][Green]](https://github.com/hashicorp/go-getter)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [sftp](https://github.com/pkg/sftp) **star:843** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   [![There was an update last week][Green]](https://github.com/pkg/sftp)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/sftp)
* [gev](https://github.com/Allenxuxu/gev) **star:793** gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.   [![godoc][GoDoc]](https://godoc.org/github.com/Allenxuxu/gev)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:774** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   [![There was an update last week][Green]](https://github.com/intel-go/nff-go)   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [grab](https://github.com/cavaliercoder/grab) **star:646** Go package for managing file downloads.   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [mdns](https://github.com/hashicorp/mdns) **star:624** Simple mDNS (Multicast DNS) client/server library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/mdns)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [ftp](https://github.com/jlaffaye/ftp) **star:622** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   [![There was an update last week][Green]](https://github.com/jlaffaye/ftp)   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [lhttp](https://github.com/fanux/lhttp) **star:549** Powerful websocket framework, build your IM server more easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fanux/lhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   [![Contains Chinese documents][CN]](https://github.com/fanux/lhttp)
* [gosnmp](https://github.com/soniah/gosnmp) **star:498** Native Go library for performing SNMP actions.   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [cidranger](https://github.com/yl2chen/cidranger) **star:455** Fast IP to CIDR lookup for Go.   [![There was an update last week][Green]](https://github.com/yl2chen/cidranger)   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [gotcp](https://github.com/gansidui/gotcp) **star:452** Go package for quickly writing tcp applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/gotcp)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/gotcp)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:405** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [gopcap](https://github.com/akrennmair/gopcap) **star:379** Go wrapper for libpcap.   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [raw](https://github.com/mdlayher/raw) **star:345** Package raw enables reading and writing data at the device driver level for a network interface.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [go-stun](https://github.com/ccding/go-stun) **star:344** Go implementation of the STUN client (RFC 3489 and RFC 5389).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ccding/go-stun)   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [stun](https://github.com/go-rtc/stun) **star:342** Go implementation of RFC 5389 STUN protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:324** Go library for building tcp servers faster.   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [winrm](https://github.com/masterzen/winrm) **star:242** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:234** Streaming protocolbuffer data over TCP made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stabbycutyou/buffstreams)   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [arp](https://github.com/mdlayher/arp) **star:216** Package arp implements the ARP protocol, as described in RFC 826.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:196** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [utp](https://github.com/anacrolix/utp) **star:152** Go uTP micro transport protocol implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/anacrolix/utp)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/utp)
* [jazigo](https://github.com/udhos/jazigo) **star:144** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [canopus](https://github.com/zubairhamed/canopus) **star:138** CoAP Client/Server implementation (RFC 7252).   [![It hasn't been updated in the last year][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:133** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![Contains Chinese documents][CN]](https://github.com/DrmagicE/gmqtt)
* [sslb](https://github.com/eduardonunesp/sslb) **star:124** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [gNxI](https://github.com/google/gnxi) **star:122** A collection of tools for Network Management that use the gNMI and gNOI protocols.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [gaio](https://github.com/xtaci/gaio) **star:113** High performance async-io networking for Golang in proactor mode.   [![There was an update last week][Green]](https://github.com/xtaci/gaio)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gaio)
* [xtcp](https://github.com/xfxdev/xtcp) **star:97** TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:65** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/dhcp6)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/dhcp6)
* [ether](https://github.com/songgao/ether) **star:65** Cross-platform Go package for sending and receiving ethernet frames.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [linkio](https://github.com/ian-kent/linkio) **star:47** Network link speed simulation for Reader/Writer interfaces.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/linkio)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/linkio)
* [packet](https://github.com/aerogo/packet) **star:44** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aybabtme/portproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/aybabtme/portproxy)
* [iplib](https://github.com/c-robinson/iplib) **star:29** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [graval](https://github.com/koofr/graval) **star:25** Experimental FTP server framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/koofr/graval)   [![godoc][GoDoc]](https://godoc.org/github.com/koofr/graval)
* [publicip](https://github.com/polera/publicip) **star:18** Package publicip returns your public facing IPv4 address (internet egress).   [![It hasn't been updated in the last year][Yellow]](https://github.com/polera/publicip)   [![godoc][GoDoc]](https://godoc.org/github.com/polera/publicip)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:18** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/golibwireshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/golibwireshark)
* [go-powerdns](https://github.com/joeig/go-powerdns) **star:12** PowerDNS API bindings for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/joeig/go-powerdns)
* [llb](https://github.com/kirillDanshin/llb) **star:10** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/llb)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/llb)
* [goshark](https://github.com/sunwxg/goshark) **star:9** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/goshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/goshark)
* [tspool](https://github.com/two/tspool) **star:6** A TCP Library use worker pool to improve performance and protect your server.   [![It hasn't been updated in the last year][Yellow]](https://github.com/two/tspool)   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)
* [gosocsvr](https://github.com/rakeki/gosocsvr) **star:5** Socket server made simple.   [![godoc][GoDoc]](https://godoc.org/github.com/rakeki/gosocsvr)
* [httpproxy](https://github.com/wzshiming/httpproxy) **star:2** HTTP proxy handler and dialer.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/httpproxy)

### HTTP Clients

*Libraries for making HTTP requests.*

* [resty](https://github.com/go-resty/resty) **star:2624** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   [![There was an update last week][Green]](https://github.com/go-resty/resty)   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [grequests](https://github.com/levigross/grequests) **star:1532** A Go "clone" of the great and famous Requests library.   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [heimdall](https://github.com/gojektech/heimdall) **star:1259** An enchanced http client with retry and hystrix capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/gojektech/heimdall)
* [sling](https://github.com/dghubble/sling) **star:1108** Sling is a Go HTTP client library for creating and sending API requests.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/sling)
* [gentleman](https://github.com/h2non/gentleman) **star:737** Full-featured plugin-driven HTTP client library.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gentleman)
* [pester](https://github.com/sethgrid/pester) **star:354** Go HTTP client calls with retries, backoff, and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/sethgrid/pester)
* [rq](https://github.com/ddo/rq) **star:34** A nicer interface for golang stdlib HTTP client.   [![godoc][GoDoc]](https://godoc.org/github.com/ddo/rq)
* [httpretry](https://github.com/ybbus/httpretry) **star:4** Enriches the default go HTTP client with retry functionality.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/httpretry)

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) **star:872** Go bindings for GLFW 3.   [![There was an update last week][Green]](https://github.com/go-gl/glfw)
* [gl](https://github.com/go-gl/gl) **star:695** Go bindings for OpenGL (generated via glow).   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-gl/gl)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/gl)
* [mathgl](https://github.com/go-gl/mathgl) **star:322** Pure Go math package specialized for 3D math, with inspiration from GLM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/mathgl)
* [goxjs/gl](https://github.com/goxjs/gl) **star:132** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/gl)
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:61** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) **star:17532** The fantastic ORM library for Golang, aims to be developer friendly.   [![star > 2000][Awesome]](https://github.com/jinzhu/gorm)   [![There was an update last week][Green]](https://github.com/jinzhu/gorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/gorm)
* [Xorm](https://github.com/go-xorm/xorm) **star:5843** Simple and powerful ORM for Go.   [![star > 2000][Awesome]](https://github.com/go-xorm/xorm)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xorm/xorm)   [![Contains Chinese documents][CN]](https://github.com/go-xorm/xorm)
* [go-pg](https://github.com/go-pg/pg) **star:3610** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [gorp](https://github.com/go-gorp/gorp) **star:3309** Go Relational Persistence, ORM-ish library for Go.   [![star > 2000][Awesome]](https://github.com/go-gorp/gorp)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorp/gorp)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:2712** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   [![star > 2000][Awesome]](https://github.com/volatiletech/sqlboiler)   [![There was an update last week][Green]](https://github.com/volatiletech/sqlboiler)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/sqlboiler)
* [upper.io/db](https://github.com/upper/db) **star:2100** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![star > 2000][Awesome]](https://github.com/upper/db)   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [reform](https://github.com/go-reform/reform) **star:848** Better ORM for Go, based on non-empty interfaces and code generation.   [![There was an update last week][Green]](https://github.com/go-reform/reform)   [![godoc][GoDoc]](https://godoc.org/github.com/go-reform/reform)
* [pop/soda](https://github.com/gobuffalo/pop) **star:836** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/pop)
* [QBS](https://github.com/coocood/qbs) **star:549** Stands for Query By Struct. A Go ORM.   [![It hasn't been updated in the last year][Yellow]](https://github.com/coocood/qbs)   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   [![Contains Chinese documents][CN]](https://github.com/coocood/qbs)
* [go-queryset](https://github.com/jirfag/go-queryset) **star:496** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.   [![godoc][GoDoc]](https://godoc.org/github.com/jirfag/go-queryset)
* [gormt](https://github.com/xxjwxc/gormt) **star:396** Mysql database to golang gorm struct.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gormt)
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:361** A flexible and powerful SQL string builder library plus a zero-config ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/go-sqlbuilder)
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [Zoom](https://github.com/albrow/zoom) **star:251** Blazing-fast datastore and querying engine built on Redis.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/zoom)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [grimoire](https://github.com/Fs02/grimoire) **star:126** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/grimoire)
* [go-store](https://github.com/gosuri/go-store) **star:99** Simple and fast Redis backed key-value store library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gosuri/go-store)   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/go-store)
* [Marlow](https://github.com/dadleyy/marlow) **star:76** Generated ORM from project structs for compile time safety assurances.   [![godoc][GoDoc]](https://godoc.org/github.com/dadleyy/marlow)
* [rel](https://github.com/Fs02/rel) **star:61** Golang SQL Repository Layer for Clean (Onion) Architecture.   [![There was an update last week][Green]](https://github.com/Fs02/rel)   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/rel)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:12** A simple ORM for Google/Firebase Cloud Firestore.   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)
* [lore](https://github.com/abrahambotros/lore) **star:5** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abrahambotros/lore)   [![godoc][GoDoc]](https://godoc.org/github.com/abrahambotros/lore)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:13119** Go dependency tool.   [![star > 2000][Awesome]](https://github.com/golang/dep)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/dep)
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) **star:7962** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.   [![star > 2000][Awesome]](https://github.com/Masterminds/glide)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/glide)
* [godep](https://github.com/tools/godep) **star:5646** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   [![star > 2000][Awesome]](https://github.com/tools/godep)   [![It hasn't been updated in the last year][Yellow]](https://github.com/tools/godep)   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   [![Archived][Archived]](https://github.com/tools/godep)
* [govendor](https://github.com/kardianos/govendor) **star:5013** Go Package Manager. Go vendor tool that works with the standard vendor file.   [![star > 2000][Awesome]](https://github.com/kardianos/govendor)   [![godoc][GoDoc]](https://godoc.org/github.com/kardianos/govendor)   [![Archived][Archived]](https://github.com/kardianos/govendor)
* [gopm](https://github.com/gpmgo/gopm) **star:2471** Go Package Manager.   [![star > 2000][Awesome]](https://github.com/gpmgo/gopm)   [![godoc][GoDoc]](https://godoc.org/github.com/gpmgo/gopm)   [![Archived][Archived]](https://github.com/gpmgo/gopm)
* [gom](https://github.com/mattn/gom) **star:1380** Go Manager - bundle for go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [gpm](https://github.com/pote/gpm) **star:1203** Barebones dependency manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pote/gpm)
* [goop](https://github.com/nitrous-io/goop) **star:780** Simple dependency manager for Go (golang), inspired by Bundler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nitrous-io/goop)   [![godoc][GoDoc]](https://godoc.org/github.com/nitrous-io/goop)
* [nut](https://github.com/jingweno/nut) **star:244** Vendor Go dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jingweno/nut)   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:214** Minimal dependency version using Git.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/johnny-deps)
* [gigo](https://github.com/LyricalSecurity/gigo) **star:198** PIP-like dependency tool for golang, with support for private repositories and hashes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LyricalSecurity/gigo)   [![godoc][GoDoc]](https://godoc.org/github.com/LyricalSecurity/gigo)
* [VenGO](https://github.com/DamnWidget/VenGO) **star:117** create and manage exportable isolated go virtual environments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DamnWidget/VenGO)   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:98** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.
* [gop](https://github.com/lunny/gop) **star:49** Build and manage your Go applications out of GOPATH.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lunny/gop)   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/gop)   [![Contains Chinese documents][CN]](https://github.com/lunny/gop)   [![Archived][Archived]](https://github.com/lunny/gop)

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:10427** A distributed tracing system.   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   [![There was an update last week][Green]](https://github.com/jaegertracing/jaeger)   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1203** Simple profiling support package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [tracer](https://github.com/kamilsk/tracer) **star:23** Simple, lightweight tracing.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:6121** Implementation of GraphQL for Go.   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [graphql](https://github.com/neelance/graphql-go) **star:3216** GraphQL server with a focus on ease of use.   [![star > 2000][Awesome]](https://github.com/neelance/graphql-go)   [![godoc][GoDoc]](https://godoc.org/github.com/neelance/graphql-go)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:1324** A simple Go package to Query over JSON Data.   [![There was an update last week][Green]](https://github.com/thedevsaddam/gojsonq)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [jsonql](https://github.com/elgs/jsonql) **star:220** JSON query expression library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/jsonql)
* [rql](https://github.com/a8m/rql) **star:136** Resource Query Language for REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/a8m/rql)
* [graphql](https://github.com/tmc/graphql) **star:52** graphql parser + utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tmc/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [jsonslice](https://github.com/bhmj/jsonslice) **star:39** Jsonpath queries with advanced filters.   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)
* [straf](https://github.com/SonicRoshan/straf) **star:10** Easily Convert Golang structs to GraphQL objects.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/straf)

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) **star:2666** The simple and easy way to embed static files into Go binaries.   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![There was an update last week][Green]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik) **star:2560** Embeds static files into a Go executable.   [![star > 2000][Awesome]](https://github.com/rakyll/statik)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/statik)
* [go.rice](https://github.com/GeertJohan/go.rice) **star:1869** go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy.   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:813** Generates a vfsdata.go file that statically implements the given virtual filesystem.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/vfsgen)
* [esc](https://github.com/mjibson/esc) **star:534** Embeds files into Go programs and provides http.FileSystem interfaces to them.   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:502** Simple tool to embed files in go with focus on "customization" and ease to use.   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-resources](https://github.com/omeid/go-resources) **star:165** Unfancy resources embedding with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/omeid/go-resources)
* [statics](https://github.com/go-playground/statics) **star:54** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/statics)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/statics)
* [templify](https://github.com/wlbr/templify) **star:25** Embed external template files into Go code to create single file binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/templify)
* [go-embed](https://github.com/pyros2097/go-embed) **star:21** Generates go code to embed resource files into your library or executable.   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:3626** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   [![There was an update last week][Green]](https://github.com/gonum/gonum)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [stats](https://github.com/montanaflynn/stats) **star:1647** Statistics package with common functions missing from the Golang standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gonum/plot](https://github.com/gonum/plot) **star:1494** gonum/plot provides an API for building and drawing plots in Go.   [![There was an update last week][Green]](https://github.com/gonum/plot)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [gosl](https://github.com/cpmech/gosl) **star:1401** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![godoc][GoDoc]](https://godoc.org/github.com/cpmech/gosl)
* [streamtools](https://github.com/nytlabs/streamtools) **star:1317** general purpose, graphical tool for dealing with streams of data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nytlabs/streamtools)   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [go-dsp](https://github.com/mjibson/go-dsp) **star:667** Digital Signal Processing for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mjibson/go-dsp)   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/go-dsp)
* [chart](https://github.com/vdobler/chart) **star:627** Simple Chart Plotting library for Go. Supports many graphs types.   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [goraph](https://github.com/gyuho/goraph) **star:619** Pure Go graph theory library(data structure, algorith visualization).   [![It hasn't been updated in the last year][Yellow]](https://github.com/gyuho/goraph)   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [graph](https://github.com/yourbasic/graph) **star:308** Library of basic graph algorithms.   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [ewma](https://github.com/VividCortex/ewma) **star:291** Exponentially-weighted moving averages.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [orb](https://github.com/paulmach/orb) **star:268** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:168** Dataframes for machine-learning and statistics (similar to pandas).   [![There was an update last week][Green]](https://github.com/rocketlaunchr/dataframe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:140** Approximate histograms for data streams.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/gohistogram)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:96** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DavidBelicza/TextRank)   [![godoc][GoDoc]](https://godoc.org/github.com/DavidBelicza/TextRank)
* [sparse](https://github.com/james-bowman/sparse) **star:82** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [pagerank](https://github.com/alixaxel/pagerank) **star:55** Weighted PageRank algorithm implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [geom](https://github.com/skelterjohn/geom) **star:44** 2D geometry for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/geom)   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [evaler](https://github.com/soniah/evaler) **star:39** Simple floating point arithmetic expression evaluator.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/evaler)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:17** GO Implementation of Entropy Measures.   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:13** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![There was an update last week][Green]](https://github.com/tchayen/triangolatte)   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:11** Tiny linear interpolation library.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/piecewiselinear)
* [ode](https://github.com/ChristopherRabotin/ode) **star:11** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/ode)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/ode)
* [GoStats](https://github.com/OGFris/GoStats) **star:10** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OGFris/GoStats)   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [PiHex](https://github.com/claygod/PiHex) **star:8** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [assocentity](https://github.com/ndabAP/assocentity) **star:6** Package assocentity returns the average distance from words to a given entity.   [![There was an update last week][Green]](https://github.com/ndabAP/assocentity)   [![godoc][GoDoc]](https://godoc.org/github.com/ndabAP/assocentity)
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** Graph theory algorithms written in "Go" language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/go-gt)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/go-gt)
* [rootfinding](https://github.com/khezen/rootfinding) **star:3** root-finding algorithms library for finding roots of quadratic functions.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:2** Provides a Bradley-Terry Model for pairwise comparisons.   [![godoc][GoDoc]](https://godoc.org/github.com/seanhagen/bradleyterry)

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/xenolf/lego) **star:3903** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   [![star > 2000][Awesome]](https://github.com/xenolf/lego)   [![There was an update last week][Green]](https://github.com/xenolf/lego)   [![godoc][GoDoc]](https://godoc.org/github.com/xenolf/lego)
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:2154** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![star > 2000][Awesome]](https://github.com/Ullaakut/cameradar)   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [acmetool](https://github.com/hlandau/acme) **star:1747** ACME (Let's Encrypt) client tool with automatic renewal.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/acme)
* [memguard](https://github.com/awnumar/memguard) **star:1727** A pure Go library for handling sensitive values in memory.   [![There was an update last week][Green]](https://github.com/awnumar/memguard)   [![godoc][GoDoc]](https://godoc.org/github.com/awnumar/memguard)
* [secure](https://github.com/unrolled/secure) **star:1430** HTTP middleware for Go that facilitates some quick security wins.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/secure)
* [acra](https://github.com/cossacklabs/acra) **star:547** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   [![There was an update last week][Green]](https://github.com/cossacklabs/acra)   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [nacl](https://github.com/kevinburke/nacl) **star:476** Go implementation of the NaCL set of API's.   [![godoc][GoDoc]](https://godoc.org/github.com/kevinburke/nacl)
* [BadActor](https://github.com/jaredfolkins/badactor) **star:265** In-memory, application-driven jailer built in the spirit of fail2ban.   [![godoc][GoDoc]](https://godoc.org/github.com/jaredfolkins/badactor)
* [passlib](https://github.com/hlandau/passlib) **star:233** Futureproof password hashing library.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/passlib)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:231** encrypt/decrypt using ssh keys.   [![godoc][GoDoc]](https://godoc.org/github.com/ssh-vault/ssh-vault)
* [optimus-go](https://github.com/pjebs/optimus-go) **star:193** ID hashing and Obfuscation using Knuth's Algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/pjebs/optimus-go)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:163** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [go-yara](https://github.com/hillu/go-yara) **star:153** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   [![There was an update last week][Green]](https://github.com/hillu/go-yara)   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [argon2pw](https://github.com/raja/argon2pw) **star:81** Argon2 password hash generation with constant-time password comparison.   [![It hasn't been updated in the last year][Yellow]](https://github.com/raja/argon2pw)   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:35** A probably paranoid package for securely hashing and encrypting passwords.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dwin/goSecretBoxPassword)   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
* [certificates](https://github.com/mvmaasakkers/certificates) **star:13** An opinionated tool for generating tls certificates.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:11** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:8** SSL certificates made easy with a high level wrapper around acme/autocert.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianosela/sslmgr)
* [go-generate-password](https://github.com/m1/go-generate-password) **star:7** Password generator that can be used on the cli or as a library.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-generate-password)
* [secureio](https://github.com/xaionaro-go/secureio) **star:5** An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519.   [![There was an update last week][Green]](https://github.com/xaionaro-go/secureio)   [![godoc][GoDoc]](https://godoc.org/github.com/xaionaro-go/secureio)

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:7296** High-performance 100% compatible drop-in replacement of "encoding/json".   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:6325** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   [![There was an update last week][Green]](https://github.com/golang/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:3563** Protocol Buffers for Go with Gadgets.   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:3131** Go library for decoding generic map values into native Go structures.   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1375** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [colfer](https://github.com/pascaldekloe/colfer) **star:518** Code generation for the Colfer binary format.
* [csvutil](https://github.com/jszwec/csvutil) **star:354** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:277** Cap'n Proto library and parser for go.   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:135** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yvasiyarov/php_session_decoder)   [![godoc][GoDoc]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
* [cbor](https://github.com/fxamacker/cbor) **star:117** Small, safe, and easy CBOR encoding and decoding library.   [![There was an update last week][Green]](https://github.com/fxamacker/cbor)   [![godoc][GoDoc]](https://godoc.org/github.com/fxamacker/cbor)
* [structomap](https://github.com/tuvistavie/structomap) **star:110** Library to easily and dynamically generate maps from static structures.   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [bambam](https://github.com/glycerine/bambam) **star:60** generator for Cap'n Proto schemas from go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/glycerine/bambam)   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/bambam)
* [asn1](https://github.com/PromonLogicalis/asn1) **star:44** Asn.1 BER and DER encoding library for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PromonLogicalis/asn1)   [![godoc][GoDoc]](https://godoc.org/github.com/PromonLogicalis/asn1)   [![Archived][Archived]](https://github.com/PromonLogicalis/asn1)
* [binstruct](https://github.com/ghostiam/binstruct) **star:17** Golang binary decoder for mapping data into the structure.   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [fwencoder](https://github.com/o1egl/fwencoder) **star:10** Fixed width file parser (encoding and decoding library) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/fwencoder)
* [bel](https://github.com/32leaves/bel) **star:8** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.   [![godoc][GoDoc]](https://godoc.org/github.com/32leaves/bel)
* [pletter](https://github.com/vimeda/pletter) **star:7** A standard way to wrap a proto message for message brokers.   [![There was an update last week][Green]](https://github.com/vimeda/pletter)   [![godoc][GoDoc]](https://godoc.org/github.com/vimeda/pletter)
* [elastic](https://github.com/epiclabs-io/elastic) **star:5** Convert slices, maps or any other unknown value across different types at run-time, no matter what.    [![godoc][GoDoc]](https://godoc.org/github.com/epiclabs-io/elastic)
* [fixedwidth](https://github.com/huydang284/fixedwidth) **star:4** Fixed-width text formatting (UTF-8 supported).   [![godoc][GoDoc]](https://godoc.org/github.com/huydang284/fixedwidth)

## Server Applications

* [etcd](https://github.com/coreos/etcd) **star:30049** Highly-available key value store for shared configuration and service discovery.   [![star > 2000][Awesome]](https://github.com/coreos/etcd)   [![There was an update last week][Green]](https://github.com/coreos/etcd)   [![godoc][GoDoc]](https://godoc.org/github.com/coreos/etcd)
* [Caddy](https://github.com/mholt/caddy) **star:26762** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   [![star > 2000][Awesome]](https://github.com/mholt/caddy)   [![There was an update last week][Green]](https://github.com/mholt/caddy)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/caddy)
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [minio](https://github.com/minio/minio) **star:20751** Minio is a distributed object storage server.   [![star > 2000][Awesome]](https://github.com/minio/minio)   [![There was an update last week][Green]](https://github.com/minio/minio)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [RoadRunner](https://github.com/spiral/roadrunner) **star:4098** High-performance PHP application server, load-balancer and process manager.   [![star > 2000][Awesome]](https://github.com/spiral/roadrunner)   [![There was an update last week][Green]](https://github.com/spiral/roadrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/spiral/roadrunner)
* [devd](https://github.com/cortesi/devd) **star:2955** Local webserver for developers.   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [algernon](https://github.com/xyproto/algernon) **star:1675** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/algernon)
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:1487** Full featured and highly configurable SFTP server software.   [![There was an update last week][Green]](https://github.com/drakkan/sftpgo)   [![godoc][GoDoc]](https://godoc.org/github.com/drakkan/sftpgo)
* [flipt](https://github.com/markphelps/flipt) **star:1173** A self contained feature flag solution written in Go and Vue.js   [![There was an update last week][Green]](https://github.com/markphelps/flipt)   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Trickster](https://github.com/Comcast/trickster) **star:1060** HTTP reverse proxy cache and time series accelerator.   [![There was an update last week][Green]](https://github.com/Comcast/trickster)   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/trickster)
* [yakvs](https://git.sci4me.com/sci4me/yakvs)  Small, networked, in-memory key-value store.
* [Fider](https://github.com/getfider/fider) **star:1056** Fider is an open platform to collect and organize customer feedback.   [![There was an update last week][Green]](https://github.com/getfider/fider)   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [Flagr](https://github.com/checkr/flagr) **star:1025** Flagr is an open-source feature flagging and A/B testing service.   [![There was an update last week][Green]](https://github.com/checkr/flagr)   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [discovery](https://github.com/Bilibili/discovery) **star:872** A registry for resilient mid-tier load balancing and failover.   [![There was an update last week][Green]](https://github.com/Bilibili/discovery)   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [jackal](https://github.com/ortuman/jackal) **star:797** An XMPP server written in Go.   [![There was an update last week][Green]](https://github.com/ortuman/jackal)   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [dudeldu](https://github.com/krotik/dudeldu) **star:109** A simple SHOUTcast server.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/dudeldu)
* [lets-proxy2](https://github.com/rekby/lets-proxy2) **star:26** Reverse proxy for handle https with issue certificates in fly from lets-encrypt.    [![There was an update last week][Green]](https://github.com/rekby/lets-proxy2)   [![godoc][GoDoc]](https://godoc.org/github.com/rekby/lets-proxy2)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:15** Stream database events from PostgreSQL to Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:9** Nginx log parser and exporter to Prometheus.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.


## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:287** Go stream processing library.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3495** PDF document generator with high level support for text, drawing and images.   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)   [![Archived][Archived]](https://github.com/jung-kurt/gofpdf)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1654** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [pongo2](https://github.com/flosch/pongo2) **star:1633** Django-like template-engine for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [hero](https://github.com/shiyanhui/hero) **star:1292** Hero is a handy, fast and powerful go template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/shiyanhui/hero)   [![Contains Chinese documents][CN]](https://github.com/shiyanhui/hero)
* [mustache](https://github.com/hoisie/mustache) **star:980** Go implementation of the Mustache template language.   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:840** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eknkc/amber)   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:785** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/ace)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:733** Razor view engine for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [jet](https://github.com/CloudyKit/jet) **star:634** Jet template engine.   [![There was an update last week][Green]](https://github.com/CloudyKit/jet)   [![godoc][GoDoc]](https://godoc.org/github.com/CloudyKit/jet)
* [ego](https://github.com/benbjohnson/ego) **star:429** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:368** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasttemplate)
* [raymond](https://github.com/aymerick/raymond) **star:364** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/raymond)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [Soy](https://github.com/robfig/soy) **star:148** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/robfig/soy)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [maroto](https://github.com/johnfercher/maroto) **star:122** A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.   [![godoc][GoDoc]](https://godoc.org/github.com/johnfercher/maroto)
* [goview](https://github.com/foolin/goview) **star:109** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/goview)
* [liquid](https://github.com/osteele/liquid) **star:101** Go implementation of Shopify Liquid templates.   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/liquid)
* [kasia.go](https://github.com/ziutek/kasia.go) **star:71** Templating system for HTML and other text documents - go implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/kasia.go)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [velvet](https://github.com/gobuffalo/velvet) **star:71** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gobuffalo/velvet)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)
* [damsel](https://github.com/dskinner/damsel) **star:22** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dskinner/damsel)   [![godoc][GoDoc]](https://godoc.org/github.com/dskinner/damsel)
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:19** Tiny wrapper around html/template to allow for easy file-based template inheritance.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dannyvankooten/extemplate)   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/extemplate)
* [gospin](https://github.com/m1/gospin) **star:8** Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/gospin)

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [Testify](https://github.com/stretchr/testify) **star:9884** Sacred extension to the standard go testing package.   [![star > 2000][Awesome]](https://github.com/stretchr/testify)   [![There was an update last week][Green]](https://github.com/stretchr/testify)   [![godoc][GoDoc]](https://godoc.org/github.com/stretchr/testify)
    * [go-cmp](https://github.com/google/go-cmp) **star:1557** Package for comparing Go values in tests.   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1302** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   [![godoc][GoDoc]](https://godoc.org/github.com/gavv/httpexpect)
    * [godog](https://github.com/DATA-DOG/godog) **star:961** Cucumber or Behat like BDD framework for Go.   [![There was an update last week][Green]](https://github.com/DATA-DOG/godog)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/godog)
    * [baloo](https://github.com/h2non/baloo) **star:675** Expressive and versatile end-to-end HTTP API testing made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/h2non/baloo)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [goblin](https://github.com/franela/goblin) **star:663** Mocha like testing framework fo Go.   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:454** A helper for Rails' like test fixtures to test database applications.   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:394** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:330** Mutation testing for Go source code.   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gofight](https://github.com/appleboy/gofight) **star:303** API Handler Testing for Golang Router framework.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [frisby](https://github.com/verdverm/frisby) **star:255** REST API testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:202** Tool for viewing test coverage in terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [charlatan](https://github.com/percolate/charlatan) **star:192** Tool to generate fake interface implementations for tests.   [![godoc][GoDoc]](https://godoc.org/github.com/percolate/charlatan)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:148** A collection of packages to augment the go testing package and support common patterns.   [![There was an update last week][Green]](https://github.com/gotestyourself/gotest.tools)   [![godoc][GoDoc]](https://godoc.org/github.com/gotestyourself/gotest.tools)
    * [endly](https://github.com/viant/endly) **star:133** Declarative end to end functional testing.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [commander](https://github.com/SimonBaeumer/commander) **star:119** Tool for testing cli applications on windows, linux and osx.   [![There was an update last week][Green]](https://github.com/SimonBaeumer/commander)   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:114** BDD-style testing framework for the Go programming language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/orfjackal/gospec)   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:103** Simple snapshot testing addon for your test framework.   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:102** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![There was an update last week][Green]](https://github.com/khaiql/dbcleaner)   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:84** Extremely flexible golang deep comparison, extends the go testing package.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [wstest](https://github.com/posener/wstest) **star:73** Websocket client for unit-testing a websocket http.Handler.   [![There was an update last week][Green]](https://github.com/posener/wstest)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [gospecify](https://github.com/stesla/gospecify) **star:53** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stesla/gospecify)   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [restit](https://github.com/yookoala/restit) **star:50** Go micro framework to help writing RESTful API integration test.   [![godoc][GoDoc]](https://godoc.org/github.com/yookoala/restit)
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:34** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:32** library created for testing JSON against patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [dsunit](https://github.com/viant/dsunit) **star:29** Datastore testing for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:27** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rdrdr/hamcrest)   [![godoc][GoDoc]](https://godoc.org/github.com/rdrdr/hamcrest)
    * [testcase](https://github.com/adamluzsi/testcase) **star:17** Idiomatic testing framework for Behavior Driven Development.   [![There was an update last week][Green]](https://github.com/adamluzsi/testcase)   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [assert](https://github.com/go-playground/assert) **star:17** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/assert)
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) **star:12** Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.   [![godoc][GoDoc]](https://godoc.org/github.com/fergusstrange/embedded-postgres)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:11** HTTP client testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)
    * [badio](https://github.com/cavaliercoder/badio) **star:9** Extensions to Go's `testing/iotest` package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cavaliercoder/badio)   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/badio)
    * [gocrest](https://github.com/corbym/gocrest) **star:9** Composable hamcrest-like matchers for Go assertions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gocrest)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gocrest)
    * [schema](https://github.com/jgroeneveld/schema) **star:9** Quick and easy expression matching for JSON schemas used in requests and responses.   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/schema)
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pavlo/gosuite)   [![godoc][GoDoc]](https://godoc.org/github.com/pavlo/gosuite)
    * [biff](https://github.com/fulldump/biff) **star:8** Bifurcation testing framework, BDD compatible.   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/biff)
    * [gogiven](https://github.com/corbym/gogiven) **star:8** YATSPEC-like BDD testing framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gogiven)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gogiven)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:7** Generate test data from SQL files before testing and clear it after finished.   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [Tt](https://github.com/vcaesar/tt) **star:6** Simple and colorful test tools.   [![godoc][GoDoc]](https://godoc.org/github.com/vcaesar/tt)
    * [trial](https://github.com/jgroeneveld/trial) **star:4** Quick and easy extendable assertions without introducing much boilerplate.   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/trial)

* Mock
    * [gomock](https://github.com/golang/mock) **star:3874** Mocking framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/golang/mock)   [![There was an update last week][Green]](https://github.com/golang/mock)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/mock)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:2315** Mock SQL driver for testing database interactions.   [![star > 2000][Awesome]](https://github.com/DATA-DOG/go-sqlmock)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1531** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   [![There was an update last week][Green]](https://github.com/SpectoLabs/hoverfly)   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:962** Versatile HTTP mocking made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [httpmock](https://github.com/jarcoal/httpmock) **star:719** Easy mocking of HTTP responses from external resources.   [![There was an update last week][Green]](https://github.com/jarcoal/httpmock)   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:402** Tool for generating self-contained mock objects.   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [minimock](https://github.com/gojuno/minimock) **star:294** Mock generator for Go interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:238** Single transaction based database driver mainly for testing purposes.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:88** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)
    * [timex](https://github.com/cabify/timex) **star:26** A test-friendly replacement for the native `time` package.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/timex)
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tv42/mockhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/tv42/mockhttp)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:3340** Randomized testing system.   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![There was an update last week][Green]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz) **star:739** Library for populating go objects with random values.   [![There was an update last week][Green]](https://github.com/google/gofuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gofuzz)
    * [Tavor](https://github.com/zimmski/tavor) **star:219** Generic fuzzing and delta-debugging framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zimmski/tavor)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:4429** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1527** alternative Selenium hub server that launches browsers within containers.   [![There was an update last week][Green]](https://github.com/aerokube/selenoid)   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [cdp](https://github.com/mafredri/cdp) **star:408** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:243** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)
    * [rod](https://github.com/ysmood/rod) **star:13** A chrome devtools controller that is easy and safe to use.   [![There was an update last week][Green]](https://github.com/ysmood/rod)   [![godoc][GoDoc]](https://godoc.org/github.com/ysmood/rod)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:486** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:10305** Fast and Elegant Scraping Framework for Gophers.   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![There was an update last week][Green]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:8501** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [blackfriday](https://github.com/russross/blackfriday) **star:4282** Markdown processor in Go.   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![There was an update last week][Green]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [toml](https://github.com/BurntSushi/toml) **star:3097** TOML configuration format (encoder/decoder with reflection).   [![star > 2000][Awesome]](https://github.com/BurntSushi/toml)   [![godoc][GoDoc]](https://godoc.org/github.com/BurntSushi/toml)
    * [sh](https://github.com/mvdan/sh) **star:2564** Shell parser and formatter.   [![star > 2000][Awesome]](https://github.com/mvdan/sh)   [![There was an update last week][Green]](https://github.com/mvdan/sh)   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/sh)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:2121** Formatters for time, numbers, and memory size to human readable format.   [![star > 2000][Awesome]](https://github.com/dustin/go-humanize)   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1484** HTML Sanitizer.   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [inject](https://github.com/facebookgo/inject) **star:1213** Package inject provides a reflect based injector.   [![It hasn't been updated in the last year][Yellow]](https://github.com/facebookgo/inject)   [![godoc][GoDoc]](https://godoc.org/github.com/facebookgo/inject)   [![Archived][Archived]](https://github.com/facebookgo/inject)
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1206** Parse RSS and Atom feeds in Go.   [![There was an update last week][Green]](https://github.com/mmcdole/gofeed)   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [go-toml](https://github.com/pelletier/go-toml) **star:717** Go library for the TOML format with query support and handy cli tools.   [![There was an update last week][Green]](https://github.com/pelletier/go-toml)   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:586** A collection of common regular expressions for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [slug](https://github.com/gosimple/slug) **star:518** URL-friendly slugify with multiple languages support.   [![godoc][GoDoc]](https://godoc.org/github.com/gosimple/slug)
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:366** Web scraping Framework to turn websites into structured data.   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [mxj](https://github.com/clbanning/mxj) **star:365** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:348** Parses the Graphviz DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [gotext](https://github.com/leonelquinteros/gotext) **star:253** GNU gettext utilities for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/leonelquinteros/gotext)
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:247** Functions to get fixed width of the character or string.   [![There was an update last week][Green]](https://github.com/mattn/go-runewidth)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-runewidth)
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:211** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/htmlquery)
    * [goq](https://github.com/andrewstuart/goq) **star:163** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).   [![godoc][GoDoc]](https://godoc.org/github.com/andrewstuart/goq)
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:114** NMEA parser library for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [sdp](https://github.com/gortc/sdp) **star:87** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:77** Zero-width character detection and removal for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/trubitsyn/go-zero-width)
    * [goribot](https://github.com/zhshch2002/goribot) **star:63** A simple golang spider/scraping framework,build a spider in 3 lines.   [![godoc][GoDoc]](https://godoc.org/github.com/zhshch2002/goribot)   [![Contains Chinese documents][CN]](https://github.com/zhshch2002/goribot)
    * [align](https://github.com/Guitarbum722/align) **star:63** A general purpose application that aligns text.   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [podcast](https://github.com/eduncan911/podcast) **star:61** iTunes Compliant and RSS 2.0 Podcast Generator in Golang   [![godoc][GoDoc]](https://godoc.org/github.com/eduncan911/podcast)
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:58** Make pretty slug with multiple languages support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-slugify)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [genex](https://github.com/alixaxel/genex) **star:58** Count and expand Regular Expressions into all matching Strings.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:57** Editorconfig file parser and manipulator for Go.   [![There was an update last week][Green]](https://github.com/editorconfig/editorconfig-core-go)   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:45** Functions to determine the natural language of a unicode text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/guesslanguage)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/guesslanguage)
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:40** Library for generating random strings from regular expressions.   [![godoc][GoDoc]](https://godoc.org/github.com/zach-klippenstein/goregen)
    * [allot](https://github.com/sbstjn/allot) **star:38** Placeholder and wildcard text parsing for CLI tools and bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [go-vcard](https://github.com/emersion/go-vcard) **star:37** Parse and format vCard.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:33** Fixed-width text formatting (encoder/decoder with reflection).   [![There was an update last week][Green]](https://github.com/ianlopshire/go-fixedwidth)   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [gonameparts](https://github.com/polera/gonameparts) **star:31** Parses human names into individual name parts.   [![godoc][GoDoc]](https://godoc.org/github.com/polera/gonameparts)
    * [did](https://github.com/ockam-network/did) **star:30** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [Slugify](https://github.com/avelino/slugify) **star:27** Go slugify application that handles string.   [![It hasn't been updated in the last year][Yellow]](https://github.com/avelino/slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/avelino/slugify)
    * [codetree](https://github.com/aerogo/codetree) **star:14** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [enca](https://github.com/endeveit/enca) **star:8** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/enca)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/enca)
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:7** A syndication feed for Atom 1.0 and RSS 2.0.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhengchun/syndfeed)   [![godoc][GoDoc]](https://godoc.org/github.com/zhengchun/syndfeed)
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:5** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/CalebQ42/bbConvert)   [![godoc][GoDoc]](https://godoc.org/github.com/CalebQ42/bbConvert)
    * [ltsv](https://github.com/Wing924/ltsv) **star:4** High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/ltsv)
    * [doi](https://github.com/hscells/doi) **star:4** Document object identifier (doi) parser in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hscells/doi)   [![godoc][GoDoc]](https://godoc.org/github.com/hscells/doi)
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:607** Extract urls from text.   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/xurls)
    * [gotabulate](https://github.com/bndr/gotabulate) **star:232** Easily pretty-print your tabular data with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gotabulate)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:157** fast string sorting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/radix)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [parth](https://github.com/codemodus/parth) **star:36** URL path segmentation parsing.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/parth)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/parth)
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:22** A sanitization-based swear filter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/JoshuaDoes/gofuckyourself)   [![godoc][GoDoc]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:18** Convert xml or json to go struct.   [![godoc][GoDoc]](https://godoc.org/github.com/stackerzzq/xj2go)
    * [kace](https://github.com/codemodus/kace) **star:12** Common case conversions covering common initialisms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/kace)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/kace)
    * [Tagify](https://github.com/zoomio/tagify) **star:10** Produces a set of tags from given source.   [![There was an update last week][Green]](https://github.com/zoomio/tagify)   [![godoc][GoDoc]](https://godoc.org/github.com/zoomio/tagify)
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:9** string argument parser that understands quotes and backslashes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nproc/parseargs-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nproc/parseargs-go)
    * [TySug](https://github.com/Dynom/TySug) **star:4** Alternative suggestions with respect to keyboard layouts.   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
	* [textwrap](https://github.com/isbm/textwrap) **star:1** Implementation of `textwrap` module from Python.   [![godoc][GoDoc]](https://godoc.org/github.com/isbm/textwrap)

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:5655** The official AWS SDK for the Go programming language.   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   [![There was an update last week][Green]](https://github.com/aws/aws-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [github](https://github.com/google/go-github) **star:5616** Go library for accessing the GitHub REST API v3.   [![star > 2000][Awesome]](https://github.com/google/go-github)   [![There was an update last week][Green]](https://github.com/google/go-github)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [google](https://github.com/google/google-api-go-client) **star:2163** Auto-generated Google APIs for Go.   [![star > 2000][Awesome]](https://github.com/google/google-api-go-client)   [![There was an update last week][Green]](https://github.com/google/google-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:2061** Google Cloud APIs Go Client Library.   [![star > 2000][Awesome]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![There was an update last week][Green]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
* [discordgo](https://github.com/bwmarrin/discordgo) **star:1159** Go bindings for the Discord Chat API.   [![There was an update last week][Green]](https://github.com/bwmarrin/discordgo)   [![godoc][GoDoc]](https://godoc.org/github.com/bwmarrin/discordgo)
* [stripe](https://github.com/stripe/stripe-go) **star:1098** Go client for the Stripe API.   [![There was an update last week][Green]](https://github.com/stripe/stripe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1020** Go client library for the Twitter 1.1 API.   [![There was an update last week][Green]](https://github.com/ChimeraCoder/anaconda)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [go-twitter](https://github.com/dghubble/go-twitter) **star:922** Go client library for the Twitter v1.1 APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/go-twitter)
* [minio-go](https://github.com/minio/minio-go) **star:897** Minio Go Library for Amazon S3 compatible cloud storage.   [![There was an update last week][Green]](https://github.com/minio/minio-go)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio-go)
* [facebook](https://github.com/huandu/facebook) **star:846** Go Library that supports the Facebook Graph API.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/facebook)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:699** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   [![There was an update last week][Green]](https://github.com/andygrunwald/go-jira)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-jira)
* [githubql](https://github.com/shurcooL/githubql) **star:609** Go library for accessing the GitHub GraphQL API v4.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/githubql)
* [webhooks](https://github.com/go-playground/webhooks) **star:467** Webhook receiver for GitHub and Bitbucket.   [![There was an update last week][Green]](https://github.com/go-playground/webhooks)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/webhooks)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:343** Wrapper for PayPal payment API.   [![godoc][GoDoc]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:342** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [go-marathon](https://github.com/gambol99/go-marathon) **star:192** Go library for interacting with Mesosphere's Marathon PAAS.   [![There was an update last week][Green]](https://github.com/gambol99/go-marathon)   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [ethrpc](https://github.com/onrik/ethrpc) **star:180** Go bindings for Ethereum JSON RPC API.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [Trello](https://github.com/adlio/trello) **star:128** Go wrapper for the Trello API.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [gostorm](https://github.com/jsgilmore/gostorm) **star:124** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jsgilmore/gostorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [Medium](https://github.com/Medium/medium-sdk-go) **star:123** Golang SDK for Medium's OAuth2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Medium/medium-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/Medium/medium-sdk-go)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:112** A golang package to communicate with HipChat over XMPP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daneharrigan/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [hipchat](https://github.com/andybons/hipchat) **star:108** This project implements a golang client library for the Hipchat API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andybons/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/andybons/hipchat)
* [go-trending](https://github.com/andygrunwald/go-trending) **star:108** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-trending)
* [cachet](https://github.com/andygrunwald/cachet) **star:78** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/cachet)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/cachet)
* [pushover](https://github.com/gregdel/pushover) **star:70** Go wrapper for the Pushover API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gregdel/pushover)   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [wit-go](https://github.com/wit-ai/wit-go) **star:65** Go client for wit.ai HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:57** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Go client library for interfacing with the Clarifai API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/samuelcouch/clarifai)   [![godoc][GoDoc]](https://godoc.org/github.com/samuelcouch/clarifai)
* [megos](https://github.com/andygrunwald/megos) **star:56** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/megos)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [circleci](https://github.com/jszwedko/go-circleci) **star:49** Go client library for interacting with CircleCI's API.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [gads](https://github.com/emiddleton/gads) **star:47** Google Adwords Unofficial API.   [![godoc][GoDoc]](https://godoc.org/github.com/emiddleton/gads)
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:40** Go client for the xkcd API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nishanths/go-xkcd)   [![godoc][GoDoc]](https://godoc.org/github.com/nishanths/go-xkcd)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:40** Go MusicBrainz WS2 client library.   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:40** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [simples3](https://github.com/rhnvrm/simples3) **star:36** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:36** Go wrapper and command-line client for the Uptime Robot v2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [fcm](https://github.com/maddevsio/fcm) **star:36** Go library for Firebase Cloud Messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [slack](https://github.com/nlopes/slack)  Slack API in Go.
* [gosip](https://github.com/koltyakov/gosip) **star:35** Go client library SharePoint API.   [![There was an update last week][Green]](https://github.com/koltyakov/gosip)   [![godoc][GoDoc]](https://godoc.org/github.com/koltyakov/gosip)
* [golyrics](https://github.com/mamal72/golyrics) **star:34** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [mixpanel](https://github.com/dukex/mixpanel) **star:32** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [translate](https://github.com/poorny/translate) **star:29** Go online translation package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/poorny/translate)   [![godoc][GoDoc]](https://godoc.org/github.com/poorny/translate)
* [gcm](https://github.com/Aorioli/gcm) **star:29** Go library for Google Cloud Messaging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Aorioli/gcm)   [![godoc][GoDoc]](https://godoc.org/github.com/Aorioli/gcm)
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb) **star:28** Golang wrapper for The Movie Database API v3.   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/golang-tmdb)
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:27** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![There was an update last week][Green]](https://github.com/hbagdi/go-unsplash)   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [gami](https://github.com/bit4bit/gami) **star:27** Go library for Asterisk Manager Interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![Archived][Archived]](https://github.com/bit4bit/gami)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:24** Go wrapper for the YNAB API.   [![godoc][GoDoc]](https://godoc.org/github.com/brunomvsouza/ynab.go)
* [spotify](https://github.com/rapito/go-spotify) **star:21** Go Library to access Spotify WEB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-spotify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-spotify)
* [steam](https://github.com/sostronk/go-steam) **star:19** Go Library to interact with Steam game servers.   [![godoc][GoDoc]](https://godoc.org/github.com/sostronk/go-steam)
* [shopify](https://github.com/rapito/go-shopify) **star:19** Go Library to make CRUD request to the Shopify API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rapito/go-shopify)   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:19** Go client for interacting with the Twitch v3 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/knspriggs/go-twitch)   [![godoc][GoDoc]](https://godoc.org/github.com/knspriggs/go-twitch)
* [patreon-go](https://github.com/mxpv/patreon-go) **star:19** Go library for Patreon API.   [![godoc][GoDoc]](https://godoc.org/github.com/mxpv/patreon-go)
* [codeship-go](https://github.com/codeship/codeship-go) **star:17** Go client library for interacting with Codeship's API v2.   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:17** Go library for accessing the BreweryDB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/naegelejd/brewerydb)   [![godoc][GoDoc]](https://godoc.org/github.com/naegelejd/brewerydb)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:16** Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).   [![It hasn't been updated in the last year][Yellow]](https://github.com/nstratos/go-myanimelist)   [![godoc][GoDoc]](https://godoc.org/github.com/nstratos/go-myanimelist)
* [textbelt](https://github.com/dietsche/textbelt) **star:16** Go client for the textbelt.com txt messaging API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:15** Go client library for [imgur](https://imgur.com)   [![It hasn't been updated in the last year][Yellow]](https://github.com/koffeinsource/go-imgur)   [![godoc][GoDoc]](https://godoc.org/github.com/koffeinsource/go-imgur)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:12** Go client library for interacting with Coinpaprika's API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/coinpaprika/coinpaprika-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:11** Simple wrapper for easy google analytics reporting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chonthu/go-google-analytics)   [![godoc][GoDoc]](https://godoc.org/github.com/chonthu/go-google-analytics)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:10** Tiny Go client for HackerNews API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PaulRosset/go-hacknews)   [![godoc][GoDoc]](https://godoc.org/github.com/PaulRosset/go-hacknews)
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:10** Go client library for the [LastPass](https://www.lastpass.com/) API.   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergiotapia/smitego)   [![godoc][GoDoc]](https://godoc.org/github.com/sergiotapia/smitego)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:7** Go client library for the SPTrans Olho Vivo API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergioaugrod/go-sptrans)   [![godoc][GoDoc]](https://godoc.org/github.com/sergioaugrod/go-sptrans)
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [go-here](https://github.com/abdullahselek/go-here) **star:6** Go client library around the HERE location based APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/abdullahselek/go-here)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:6** Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-google-email-audit-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-google-email-audit-api)
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattcunningham/gumblr)   [![godoc][GoDoc]](https://godoc.org/github.com/mattcunningham/gumblr)
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection) **star:5** Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).   [![There was an update last week][Green]](https://github.com/rbretecher/go-postman-collection)   [![godoc][GoDoc]](https://godoc.org/github.com/rbretecher/go-postman-collection)
* [zooz](https://github.com/gojuno/go-zooz) **star:5** Go client for the Zooz API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gojuno/go-zooz)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [go-sophos](https://github.com/esurdam/go-sophos) **star:5** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/esurdam/go-sophos)
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:5** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)
* [twitter-scraper](https://github.com/n0madic/twitter-scraper) **star:3** Scrape the Twitter Frontend API without authentication and limits.   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/twitter-scraper)
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:3** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   [![It hasn't been updated in the last year][Yellow]](https://github.com/axelspringer/go-chronos)   [![godoc][GoDoc]](https://godoc.org/github.com/axelspringer/go-chronos)
* [google-play-scraper](https://github.com/n0madic/google-play-scraper) **star:2** Get data from Google Play Store.   [![There was an update last week][Green]](https://github.com/n0madic/google-play-scraper)   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/google-play-scraper)
* [gopaapi5](https://github.com/utekaravinash/gopaapi5) **star:2** Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).   [![There was an update last week][Green]](https://github.com/utekaravinash/gopaapi5)   [![godoc][GoDoc]](https://godoc.org/github.com/utekaravinash/gopaapi5)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.   [![godoc][GoDoc]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
* [vl-go](https://github.com/verifid/vl-go) **star:1** Go client library around the VerifID identity verification layer API.   [![godoc][GoDoc]](https://godoc.org/github.com/verifid/vl-go)
* [libgoffi](https://github.com/clevabit/libgoffi) **star:1** Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** The Playlyfe Rest API Go SDK.   [![It hasn't been updated in the last year][Yellow]](https://github.com/playlyfe/playlyfe-go-sdk)   [![godoc][GoDoc]](https://godoc.org/github.com/playlyfe/playlyfe-go-sdk)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:27737** Command-line fuzzy finder written in Go.   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![There was an update last week][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [hub](https://github.com/github/hub) **star:19231** wrap git commands with additional functionality to interact with github from the terminal.   [![star > 2000][Awesome]](https://github.com/github/hub)   [![There was an update last week][Green]](https://github.com/github/hub)   [![godoc][GoDoc]](https://godoc.org/github.com/github/hub)
* [ctop](https://github.com/bcicen/ctop) **star:9603** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:5296** Deliver Go binaries as fast and easily as possible.   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![There was an update last week][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [goreporter](https://github.com/wgliang/goreporter)  Golang tool that does static analysis, unit testing, code review and generate code quality report.
* [godropbox](https://github.com/dropbox/godropbox) **star:3849** Common libraries for writing Go services/applications from Dropbox.   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [xferspdy](https://github.com/monmohan/xferspdy)  Xferspdy provides binary diff and patch library in golang.
* [panicparse](https://github.com/maruel/panicparse)  Groups similar goroutines and colorizes stack dump.
* [onecache](https://github.com/adelowo/onecache)  Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
* [peco](https://github.com/peco/peco)  Simplistic interactive filtering tool.
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match)  Pattern matching libray.
* [okrun](https://github.com/xta/okrun)  go run error steamroller.
* [pgo](https://github.com/arthurkushman/pgo)  Convenient functions for PHP community.
* [pm](https://github.com/VividCortex/pm)  Process (i.e. goroutine) manager with an HTTP API.
* [ptr](https://github.com/gotidy/ptr)  Package that provide functions for simplified creation of pointers from constants of basic types.
* [r](https://github.com/is5/r)  Python-like `range()` experience for Go.
* [rclient](https://github.com/zpatrick/rclient)  Readable, flexible, simple-to-use client for REST APIs.
* [olaf](https://github.com/btnguyen2k/olaf)  Twitter Snowflake implemented in Go.
* [mssqlx](https://github.com/linxGnu/mssqlx)  Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
* [myhttp](https://github.com/inancgumus/myhttp)  Simple API to make HTTP GET requests with timeout support.
* [multitick](https://github.com/VividCortex/multitick)  Multiplexor for aligned tickers.
* [repeat](https://github.com/ssgreg/repeat)  Go implementation of different backoff strategies useful for retrying operations and heartbeating.
* [mole](https://github.com/davrodpin/mole)  cli app to easily create ssh tunnels.
* [moldova](https://github.com/StabbyCutyou/moldova)  Utility for generating random data based on an input template.
* [mmake](https://github.com/tj/mmake)  Modern Make.
* [minquery](https://github.com/icza/minquery)  MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).
* [minify](https://github.com/tdewolff/minify)  Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.
* [mimetype](https://github.com/gabriel-vasile/mimetype)  Package for MIME type detection based on magic numbers.
* [mimesniffer](https://github.com/aofei/mimesniffer)  A MIME type sniffer for Go.
* [mimemagic](https://github.com/zRedShift/mimemagic)  Pure Go ultra performant MIME sniffing library/utility.
* [mergo](https://github.com/imdario/mergo)  Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.
* [mc](https://github.com/minio/mc)  Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.
* [realize](https://github.com/tockins/realize)  Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.
* [rerate](https://github.com/abo/rerate)  Redis-based rate counter and rate limiter for Go.
* [request](https://github.com/mozillazg/request)  Go HTTP Requests for Humans™.
* [slicer](https://github.com/leaanthony/slicer)  Makes working with slices easier.
* [wuzz](https://github.com/asciimoo/wuzz)  Interactive cli tool for HTTP inspection.
* [util](https://github.com/shomali11/util)  Collection of useful utility functions. (strings, concurrency, manipulations, ...).
* [usql](https://github.com/knq/usql)  usql is a universal command-line interface for SQL databases.
* [UNIS](https://github.com/esemplastic/unis)  Common Architecture™ for String Utilities in Go.
* [ugo](https://github.com/alxrm/ugo)  ugo is slice toolbox with concise syntax for Go.
* [toolbox](https://github.com/viant/toolbox)  Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.
* [tome](https://github.com/cyruzin/tome)  Tome was designed to paginate simple RESTful APIs.
* [Task](https://github.com/go-task/task)  simple "Make" alternative.
* [structs](https://github.com/PumpkinSeed/structs)  Implement simple functions to manipulate structs.
* [Storm](https://github.com/asdine/storm)  Simple and powerful toolkit for BoltDB.
* [sqlx](https://github.com/jmoiron/sqlx)  provides a set of extensions on top of the excellent built-in database/sql package.
* [spinner](https://github.com/briandowns/spinner)  Go package to easily provide a terminal spinner with options.
* [sorty](https://github.com/jfcg/sorty)  Fast Concurrent / Parallel Sorting.
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv)  Slice conversion between primitive types.
* [limiters](https://github.com/mennanov/limiters)  Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.
* [slice](https://github.com/psampaz/slice)  Type-safe functions for common Go slice operations.
* [silk](https://github.com/chrispassas/silk)  Read silk netflow files.
* [shutdown](https://github.com/ztrue/shutdown)  App shutdown hooks for `os.Signal` handling.
* [serve](https://github.com/syntaqx/serve)  A static http server anywhere you need.
* [scan](https://github.com/blockloop/scan)  Scan golang `sql.Rows` directly to structs, slices, or primitive types.
* [robustly](https://github.com/VividCortex/robustly)  Runs functions resiliently, catching and restarting panics.
* [retry-go](https://github.com/rafaeljesus/retry-go)  Retrying made simple and easy for golang.
* [retry](https://github.com/shafreeck/retry)  A pretty simple library to ensure your work to be done.
* [retry](https://github.com/thedevsaddam/retry)  Simple and easy retry mechanism package for Go.
* [retry](https://github.com/percolate/retry)  A simple but highly configurable retry package for Go.
* [retry](https://github.com/kamilsk/retry)  The most advanced functional mechanism to perform actions repetitively until successful.
* [rest-go](https://github.com/edermanoel94/rest-go)  A package that provide many helpful methods for working with rest api.
* [rerun](https://github.com/ivpusic/rerun)  Recompiling and rerunning go apps when source changes.
* [lrserver](https://github.com/jaschaephraim/lrserver)  LiveReload server for Go.
* [intrinsic](https://github.com/mengzhuo/intrinsic)  Use x86 SIMD without writing any assembly code.
* [koazee](https://github.com/wesovilabs/koazee)  Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.
* [hystrix-go](https://github.com/afex/hystrix-go) **star:2387** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [jsend](https://github.com/clevergo/jsend)  JSend's implementation writen in Go.
* [jump](https://github.com/gsamokovarov/jump)  Jump helps you navigate faster by learning your habits.
* [immortal](https://github.com/immortal/immortal)  \*nix cross-platform (OS agnostic) supervisor.
* [go-funk](https://github.com/thoas/go-funk) **star:1652** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![There was an update last week][Green]](https://github.com/thoas/go-funk)   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [filetype](https://github.com/h2non/filetype) **star:1059** Small package to infer the file type checking the magic numbers signature.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [boilr](https://github.com/tmrts/boilr) **star:1059** Blazingly fast CLI tool for creating projects from boilerplate templates.   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:859** Circuit Breakers in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:785** Simple, seamless, lightweight time tracking for Git.   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [htcat](https://github.com/htcat/htcat) **star:511** Parallel and Pipelined HTTP GET Utility.   [![It hasn't been updated in the last year][Yellow]](https://github.com/htcat/htcat)   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [go-dry](https://github.com/ungerik/go-dry) **star:446** DRY (don't repeat yourself) package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [circuit](https://github.com/cep21/circuit) **star:441** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/cep21/circuit)
* [gopencils](https://github.com/bndr/gopencils) **star:431** Small and simple package to easily consume REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gopencils)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [godaemon](https://github.com/VividCortex/godaemon) **star:431** Utility to write daemons.   [![It hasn't been updated in the last year][Yellow]](https://github.com/VividCortex/godaemon)   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/godaemon)
* [ergo](https://github.com/cristianoliveira/ergo) **star:363** The management of multiple local services running over different ports made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [go-rate](https://github.com/beefsack/go-rate) **star:301** Timed rate limiter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beefsack/go-rate)   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [clockwork](https://github.com/jonboulle/clockwork) **star:258** A simple fake clock for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [gohper](https://github.com/cosiner/gohper) **star:252** Various tools/modules help for development.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cosiner/gohper)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/gohper)   [![Archived][Archived]](https://github.com/cosiner/gohper)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:246** Simple struct copying for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [gubrak](https://github.com/novalagung/gubrak) **star:209** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:193** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [gotenv](https://github.com/subosito/gotenv) **star:169** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:165** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [Death](https://github.com/vrecan/death) **star:149** Managing go application shutdown with signals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/vrecan/death)   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [apm](https://github.com/topfreegames/apm) **star:137** Process manager for Golang applications with an HTTP API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/topfreegames/apm)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/apm)
* [chyle](https://github.com/antham/chyle) **star:117** Changelog generator using a git repository with multiple configuration possibilities.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:116** XML Sitemap generator written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:91** Pure Go bsdiff and bspatch libraries and CLI tools.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gabstv/go-bsdiff)   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [go-health](https://github.com/Talento90/go-health) **star:74** Health package simplifies the way you add health check to your services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Talento90/go-health)   [![godoc][GoDoc]](https://godoc.org/github.com/Talento90/go-health)
* [handy](https://github.com/miguelpragier/handy) **star:51** Many utilities and helpers like string handlers/formatters and validators.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [golog](https://github.com/mlimaloureiro/golog) **star:48** Easy and lightweight CLI tool to time track your tasks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mlimaloureiro/golog)   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:46** Parse TODOs in your GO code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/asticode/go-astitodo)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [goreadability](https://github.com/philipjkim/goreadability) **star:46** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:45** SeaweedFS client library with almost full features.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [gaper](https://github.com/maxcnunes/gaper) **star:42** Builds and restarts a Go project when it crashes or some watched file changes.   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [goback](https://github.com/carlescere/goback) **star:42** Go simple exponential backoff package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [golarm](https://github.com/msempere/golarm) **star:38** Fire alarms with system events.   [![It hasn't been updated in the last year][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:38** Universal multi-workstation clipboard that uses S3 like backend for the storage.   [![godoc][GoDoc]](https://godoc.org/github.com/jutkko/copy-pasta)
* [beyond](https://github.com/wesovilabs/beyond) **star:34** The Go tool that will drive you to the AOP world!   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [gpath](https://github.com/tenntenn/gpath) **star:33** Library to simplify access struct fields with Go's expression in reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [dbt](https://github.com/nikogura/dbt) **star:26** A framework for running self-updating signed binaries from a central, trusted repository.   [![There was an update last week][Green]](https://github.com/nikogura/dbt)   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [countries](https://github.com/biter777/countries) **star:25** Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts.   [![There was an update last week][Green]](https://github.com/biter777/countries)   [![godoc][GoDoc]](https://godoc.org/github.com/biter777/countries)
* [generate](https://github.com/go-playground/generate) **star:23** runs go generate recursively on a specified path or environment variable and can filter by regex.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [gostrutils](https://github.com/ik5/gostrutils) **star:22** Collections of string manipulation and conversion functions.   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:21** a small golang lib to generate placeholder images.   [![It hasn't been updated in the last year][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:21** Library for executing shell commands on osx, windows and linux.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [delve](https://github.com/derekparker/delve) **star:20** Go debugger.   [![There was an update last week][Green]](https://github.com/derekparker/delve)   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [evaluator](https://github.com/nullne/evaluator) **star:20** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nullne/evaluator)   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [filter](https://github.com/gookit/filter) **star:20** provide filtering, sanitizing, and conversion of Go data.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:16** Go library for encoding structs into Header fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [dlog](https://github.com/kirillDanshin/dlog) **star:15** Compile-time controlled logger to make your release smaller without removing debug calls.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [filler](https://github.com/yaronsumel/filler) **star:15** small utility to fill structs using "fill" tag.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [ghokin](https://github.com/antham/ghokin) **star:13** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [ctxutil](https://github.com/posener/ctxutil) **star:12** A collection of utility functions for contexts.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [command](https://github.com/txgruppi/command) **star:10** Command pattern for Go with thread safe serial and parallel dispatcher.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [backscanner](https://github.com/icza/backscanner) **star:9** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:5** Go package for working with Problem Details.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
* [blank](https://github.com/Henry-Sarabia/blank) **star:4** Verify or remove blanks and whitespace from strings.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [compare](https://github.com/posener/compare) **star:1** Enables more readable and easier comparison tasks.   [![There was an update last week][Green]](https://github.com/posener/compare)   [![godoc][GoDoc]](https://godoc.org/github.com/posener/compare)

## UUID

*Libraries for working with UUIDs.*

* [goid](https://github.com/jakehl/goid)  Generate and Parse RFC4122 compliant V4 UUIDs.
* [nanoid](https://github.com/aidarkhanov/nanoid)  A tiny and efficient Go unique string ID generator.
* [sno](https://github.com/muyo/sno)  Compact, sortable and fast unique IDs with embedded metadata.
* [ulid](https://github.com/oklog/ulid)  Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands.
* [uuid](https://github.com/agext/uuid)  Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
* [uuid](https://github.com/gofrs/uuid)  Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.
* [wuid](https://github.com/edwingeng/wuid)  An extremely fast unique number generator, 10-135 times faster than UUID.

## Validation

*Libraries for validation.*

* [checkdigit](https://github.com/osamingo/checkdigit)  Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
* [gody](https://github.com/guiferpa/gody)  :balloon: A lightweight struct validator for Go.
* [govalid](https://github.com/twharmon/govalid)  Fast, tag-based validation for structs.
* [govalidator](https://github.com/asaskevich/govalidator)  Validators and sanitizers for strings, numerics, slices and structs.
* [govalidator](https://github.com/thedevsaddam/govalidator)  Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.
* [jio](https://github.com/faceair/jio)  jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)  Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.
* [terraform-validator](https://github.com/thazelart/terraform-validator)  A norms and conventions validator for Terraform.
* [validate](https://github.com/gookit/validate)  Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.
* [validate](https://github.com/gobuffalo/validate)  This package provides a framework for writing validations for Go applications.
* [validator](https://github.com/go-playground/validator)  Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.

## Version Control

*Libraries for version control.*

* [gh](https://github.com/rjeczalik/gh)  Scriptable server and net/http middleware for GitHub Webhooks.
* [git2go](https://github.com/libgit2/git2go)  Go bindings for libgit2.
* [go-git](https://github.com/src-d/go-git)  highly extensible Git implementation in pure Go.
* [go-vcs](https://github.com/sourcegraph/go-vcs)  manipulate and inspect VCS repositories in Go.
* [hercules](https://github.com/src-d/hercules)  gaining advanced insights from Git repository history.
* [hgo](https://github.com/beyang/hgo)  Hgo is a collection of Go packages providing read-access to local Mercurial repositories.

## Video

*Libraries for manipulating video.*

* [gmf](https://github.com/3d0c/gmf)  Go bindings for FFmpeg av\* libraries.
* [go-astisub](https://github.com/asticode/go-astisub)  Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).
* [go-astits](https://github.com/asticode/go-astits)  Parse and demux MPEG Transport Streams (.ts) natively in GO.
* [go-m3u8](https://github.com/quangngotan95/go-m3u8)  Parser and generator library for Apple m3u8 playlists.
* [go-mpd](https://github.com/unki2aut/go-mpd)  Parser and generator library for MPEG-DASH manifest files.
* [goav](https://github.com/giorgisio/goav)  Comphrensive Go bindings for FFmpeg.
* [gst](https://github.com/ziutek/gst)  Go bindings for GStreamer.
* [libgosubs](https://github.com/wargarblgarbl/libgosubs)  Subtitle format support for go. Supports .srt, .ttml, and .ass.
* [libvlc-go](https://github.com/adrg/libvlc-go)  Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).
* [m3u8](https://github.com/grafov/m3u8)  Parser and generator library of M3U8 playlists for Apple HLS.
* [v4l](https://github.com/korandiz/v4l)  Video capture library for Linux, written in Go.

## Web Frameworks

*Full stack web frameworks.*

* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [hiboot](https://github.com/hidevopsio/hiboot)  hiboot is a high performance web application framework with auto configuration and dependency injection support.
* [Macaron](https://github.com/go-macaron/macaron)  Macaron is a high productive and modular design web framework in Go.
* [mango](https://github.com/paulbellamy/mango)  Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
* [Microservice](https://github.com/claygod/microservice)  The framework for the creation of microservices, written in Golang.
* [neo](https://github.com/ivpusic/neo)  Neo is minimal and fast Go Web Framework with extremely simple API.
* [patron](https://github.com/beatlabs/patron)  Patron is a microservice framework following best cloud practices with a focus on productivity.
* [Resoursea](https://github.com/resoursea/api)  REST framework for quickly writing resource based services.
* [Revel](https://github.com/revel/revel)  High-productivity web framework for the Go language.
* [goweb](https://github.com/twharmon/goweb)  Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS.
* [rex](https://github.com/goanywhere/rex)  Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.
* [rux](https://github.com/gookit/rux)  Simple and fast web framework for build golang HTTP applications.
* [tango](https://github.com/lunny/tango)  Micro & pluggable web framework for Go.
* [tigertonic](https://github.com/rcrowley/go-tigertonic)  Go framework for building JSON web services inspired by Dropwizard.
* [uAdmin](https://github.com/uadmin/uadmin)  Fully featured web framework for Golang, inspired by Django.
* [utron](https://github.com/gernest/utron)  Lightweight MVC framework for Go(Golang).
* [vox](https://github.com/aisk/vox)  A golang web framework for humans, inspired by Koa heavily.
* [WebGo](https://github.com/bnkamalesh/webgo)  A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).
* [Goyave](https://github.com/System-Glitch/goyave)  Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities.
* [gongular](https://github.com/mustafaakin/gongular)  Fast Go web framework with input mapping/validation and (DI) Dependency Injection.
* [Aero](https://github.com/aerogo/aero)  High-performance web framework for Go, reaches top scores in Lighthouse.
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce)  Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.
* [Air](https://github.com/aofei/air)  An ideally refined web framework for Go.
* [Banjo](https://github.com/nsheremet/banjo)  Very simple and fast web framework for Go.
* [Beego](https://github.com/astaxie/beego)  beego is an open-source, high-performance web framework for the Go programming language.
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [Echo](https://github.com/labstack/echo)  High performance, minimalist Go web framework.
* [Fiber](https://github.com/gofiber/fiber)  An Express.js inspired web framework build on Fasthttp.
* [Fireball](https://github.com/zpatrick/fireball)  More "natural" feeling web framework.
* [Flamingo](https://github.com/i-love-flamingo/flamingo)  Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.
* [Gin](https://github.com/gin-gonic/gin)  Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
* [Gondola](https://github.com/rainycape/gondola)  The web framework for writing faster sites, faster.
* [Ginrpc](https://github.com/xxjwxc/ginrpc)  Gin parameter automatic binding tool,gin rpc tools.
* [Gizmo](https://github.com/NYTimes/gizmo)  Microservice toolkit used by the New York Times.
* [go-json-rest](https://github.com/ant0ine/go-json-rest)  Quick and easy way to setup a RESTful JSON API.
* [go-rest](https://github.com/ungerik/go-rest)  Small and evil REST framework for Go.
* [Goa](https://github.com/goadesign/goa)  Goa provides a holistic approach for developing remote APIs and microservices in Go.
* [goa](https://github.com/goa-go/goa)  goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware.
* [Golax](https://github.com/fulldump/golax)  A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.
* [Golf](https://github.com/dinever/golf)  Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.
* [YARF](https://github.com/yarf-framework/yarf)  Fast micro-framework designed to build REST APIs and web services in a fast and simple way.

### Middlewares

#### Actual middlewares

* [client-timing](https://github.com/posener/client-timing)  An HTTP client for Server-Timing header.
* [CORS](https://github.com/rs/cors)  Easily add CORS capabilities to your API.
* [formjson](https://github.com/rs/formjson)  Transparently handle JSON input as a standard form POST.
* [go-server-timing](https://github.com/mitchellh/go-server-timing)  Add/parse Server-Timing header.
* [Limiter](https://github.com/ulule/limiter)  Dead simple rate limit middleware for Go.
* [ln-paywall](https://github.com/philippgille/ln-paywall)  Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).
* [Tollbooth](https://github.com/didip/tollbooth)  Rate limit HTTP request handler.
* [XFF](https://github.com/sebest/xff)  Handle `X-Forwarded-For` header and friends.

#### Libraries for creating HTTP middlewares

* [alice](https://github.com/justinas/alice)  Painless middleware chaining for Go.
* [catena](https://github.com/codemodus/catena)  http.Handler wrapper catenation (same API as "chain").
* [chain](https://github.com/codemodus/chain)  Handler wrapper chaining with scoped data (net/context-based "middleware").
* [go-wrap](https://github.com/go-on/wrap)  Small middlewares package for net/http.
* [gores](https://github.com/alioygur/gores)  Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
* [interpose](https://github.com/carbocation/interpose)  Minimalist net/http middleware for golang.
* [muxchain](https://github.com/stephens2424/muxchain)  Lightweight middleware for net/http.
* [negroni](https://github.com/urfave/negroni)  Idiomatic HTTP middleware for Golang.
* [render](https://github.com/unrolled/render)  Go package for easily rendering JSON, XML, and HTML template responses.
* [renderer](https://github.com/thedevsaddam/renderer)  Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.
* [rye](https://github.com/InVisionApp/rye)  Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.
* [stats](https://github.com/thoas/stats)  Go middleware that stores various information about your web application.

### Routers

* [alien](https://github.com/gernest/alien)  Lightweight and fast http router from outer space.
* [httprouter](https://github.com/julienschmidt/httprouter)  High performance router. Use this and the standard http handlers to form a very high performance web framework.
* [xmux](https://github.com/rs/xmux)  High performance muxer based on `httprouter` with `net/context` support.
* [violetear](https://github.com/nbari/violetear)  Go HTTP router.
* [vestigo](https://github.com/husobee/vestigo)  Performant, stand-alone, HTTP compliant URL Router for go web applications.
* [Siesta](https://github.com/VividCortex/siesta)  Composable framework to write middleware and handlers.
* [pure](https://github.com/go-playground/pure)  Is a lightweight HTTP router that sticks to the std "net/http" implementation.
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)  An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.
* [mux](https://github.com/gorilla/mux)  Powerful URL router and dispatcher for golang.
* [lars](https://github.com/go-playground/lars)  Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.
* [httptreemux](https://github.com/dimfeld/httptreemux)  High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.
* [gowww/router](https://github.com/gowww/router)  Lightning fast HTTP router fully compatible with the net/http.Handler interface.
* [bellt](https://github.com/GuilhermeCaruso/bellt)  A simple Go HTTP router.
* [GoRouter](https://github.com/vardius/gorouter)  GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
* [goroute](https://github.com/goroute/route)  Simple yet powerful HTTP request multiplexer.
* [Goji](https://github.com/goji/goji)  Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.
* [gocraft/web](https://github.com/gocraft/web)  Mux and middleware package in Go.
* [FastRouter](https://github.com/razonyang/fastrouter)  a fast, flexible HTTP router written in Go.
* [fasthttprouter](https://github.com/buaazp/fasthttprouter)  High performance router forked from `httprouter`. The first router fit for `fasthttp`.
* [chi](https://github.com/go-chi/chi)  Small, fast and expressive HTTP router built on net/context.
* [Bxog](https://github.com/claygod/Bxog)  Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
* [Bone](https://github.com/go-zoo/bone)  Lightning Fast HTTP Multiplexer.
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter)  A simple and fast HTTP router for Go.

## Windows

* [d3d9](https://github.com/gonutz/d3d9)  Go bindings for Direct3D9.
* [go-ole](https://github.com/go-ole/go-ole)  Win32 OLE implementation for golang.
* [gosddl](https://github.com/MonaxGT/gosddl)  Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

## XML

*Libraries and tools for manipulating XML.*

* [XML-Comp](https://github.com/xml-comp/xml-comp)  Simple command line XML comparer that generates diffs of folders, files and tags.
* [xml2map](https://github.com/sbabiv/xml2map)  XML to MAP converter written Golang.
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter)  Procedural XML generation API based on libxml2's xmlwriter module.
* [xpath](https://github.com/antchfx/xpath)  XPath package for Go.
* [xquery](https://github.com/antchfx/xquery)  XQuery lets you extract data from HTML/XML documents using XPath expression.
* [zek](https://github.com/miku/zek)  Generate a Go struct from XML.

# Tools

*Go software and plugins.*

## Code Analysis

* [apicompat](https://github.com/bradleyfalzon/apicompat)  Checks recent changes to a Go project for backwards incompatible changes.
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [unconvert](https://github.com/mdempsky/unconvert)  Remove unnecessary type conversions from Go source.
* [tickgit](https://github.com/augmentable-dev/tickgit)  CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp)  tarp finds functions and methods without direct unit tests in Go source code.
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [php-parser](https://github.com/z7zmey/php-parser)  A Parser for PHP written in Go.
* [lint](https://github.com/surullabs/lint)  Run linters as part of go test.
* [gostatus](https://github.com/shurcooL/gostatus)  Command line tool, shows the status of repositories that contain Go packages.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [GoPlantUML](https://github.com/jfeliu007/goplantuml)  Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.
* [GoLint](https://github.com/golang/lint)  Golint is a linter for Go source code.
* [dupl](https://github.com/mibk/dupl)  Tool for code clone detection.
* [golines](https://github.com/segmentio/golines)  Formatter that automatically shortens long lines in Go code.
* [GolangCI](https://golangci.com/)  GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer)  Web based Golang AST visualizer.
* [go-outdated](https://github.com/firstrow/go-outdated)  Console application that displays outdated packages.
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated)  An easy way to find outdated dependencies of your Go projects.
* [go-critic](https://github.com/go-critic/go-critic)  source code linter that brings checks that are currently not implemented in other linters.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch)  go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.
* [go-checkstyle](https://github.com/qiniu/checkstyle)  checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.
* [gcvis](https://github.com/davecheney/gcvis)  Visualise Go program GC trace data in real time.
* [errcheck](https://github.com/kisielk/errcheck)  Errcheck is a program for checking for unchecked errors in Go programs.
* [validate](https://github.com/mccoyst/validate)  Automatically validates struct fields with tags.

## Editor Plugins

* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server)  A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
* [go-mode](https://github.com/dominikh/go-mode.el)  Go mode for GNU/Emacs.
* [go-plus](https://github.com/joefitzgerald/go-plus)  Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.
* [gocode](https://github.com/nsf/gocode)  Autocompletion daemon for the Go programming language.
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [GoSublime](https://github.com/DisposaBoy/GoSublime)  Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.
* [gounit-vim](https://github.com/hexdigest/gounit-vim)  Vim plugin for generating Go tests based on the function's or method's signature.
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension)  Go language support for the Theia IDE.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go)  Vim plugin to highlight syntax errors on save.
* [vim-go](https://github.com/fatih/vim-go)  Go development plugin for Vim.
* [vscode-go](https://github.com/Microsoft/vscode-go)  Extension for Visual Studio Code (VS Code) which provides support for the Go language.
* [Watch](https://github.com/eaburns/Watch)  Runs a command in an acme win on file changes.

## Go Generate Tools

* [generic](https://github.com/usk81/generic)  flexible data type for Go.
* [genny](https://github.com/cheekybits/genny)  Elegant generics for Go.
* [gocontracts](https://github.com/Parquery/gocontracts)  brings design-by-contract to Go by synchronizing the code with the documentation.
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [gotests](https://github.com/cweill/gotests)  Generate Go tests from your source code.
* [gounit](https://github.com/hexdigest/gounit)  Generate Go tests using your own templates.
* [hasgo](https://github.com/DylanMeeus/hasgo)  Generate Haskell inspired functions for your slices.
* [re2dfa](https://github.com/opennota/re2dfa)  Transform regular expressions into finite state machines and output Go source code.
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.

## Go Tools

* [colorgo](https://github.com/songgao/colorgo)  Wrapper around `go` command for colorized `go build` output.
* [depth](https://github.com/KyleBanks/depth)  Visualize dependency trees of any package by analyzing imports.
* [gb](https://getgb.io/)  An easy to use project based build tool for the Go programming language.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang)  A [Yeoman](http://yeoman.io) generator to get new Go projects started.
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.
* [go-callvis](https://github.com/TrueFurby/go-callvis)  Visualize call graph of your Go program using dot format.
* [go-james](https://github.com/pieterclaerhout/go-james)  Go project skeleton creator, builds and tests your projects without the manual setup.
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete)  Bash completion for go and wgo.
* [go-swagger](https://github.com/go-swagger/go-swagger)  Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.
* [godbg](https://github.com/tylerwince/godbg)  Implementation of Rusts `dbg!` macro for quick and easy debugging during development.
* [gomodrun](https://github.com/dustinblackman/gomodrun/)  Go tool that executes and caches binaries included in go.mod files.
* [gothanks](https://github.com/psampaz/gothanks)  GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.
* [OctoLinker](https://github.com/OctoLinker/browser-extension)  Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
* [richgo](https://github.com/kyoh86/richgo)  Enrich `go test` outputs with text decorations.
* [rts](https://github.com/galeone/rts)  RTS: response to struct. Generates Go structs from server responses.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [aptly](https://github.com/smira/aptly)  aptly is a Debian repository management tool.
* [Rodent](https://github.com/alouche/rodent)  Rodent helps you manage Go versions, projects and track dependencies.
* [kcli](https://github.com/cswank/kcli)  Command line tool for inspecting kafka topics/partitions/messages.
* [kubernetes](https://github.com/kubernetes/kubernetes)  Container Cluster Manager from Google.
* [lstags](https://github.com/ivanilves/lstags)  Tool and API to sync Docker images across different registries.
* [lwc](https://github.com/timdp/lwc)  A live-updating version of the UNIX wc command.
* [manssh](https://github.com/xwjdsh/manssh)  manssh is a command line tool for managing your ssh alias config easily.
* [Moby](https://github.com/moby/moby)  Collaborative project for the container ecosystem to assemble container-based systems.
* [Mora](https://github.com/emicklei/mora)  REST server for accessing MongoDB documents and meta data.
* [ostent](https://github.com/ostrost/ostent)  collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.
* [Packer](https://github.com/mitchellh/packer)  Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
* [Pewpew](https://github.com/bengadbois/pewpew)  Flexible HTTP command line stress tester.
* [Pomerium](https://github.com/pomerium/pomerium)  Pomerium is an identity-aware access proxy.
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy)  S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).
* [jcli](https://github.com/jenkins-zh/jenkins-cli)  Jenkins CLI allows you manage your Jenkins as an easy way.
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r)  Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli)  Manage BareMetal Servers from Command Line (as easily as with Docker).
* [script](https://github.com/bitfield/script)  Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.
* [sg](https://github.com/ChristopherRabotin/sg)  Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.
* [skm](https://github.com/TimothyYe/skm)  SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!
* [StatusOK](https://github.com/sanathp/statusok)  Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi)  Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.
* [traefik](https://github.com/containous/traefik)  Reverse proxy and load balancer with support for multiple backends.
* [uTask](https://github.com/ovh/utask)  Automation engine that models and executes business processes declared in yaml.
* [Vegeta](https://github.com/tsenart/vegeta)  HTTP load testing tool and library. It's over 9000!
* [webhook](https://github.com/adnanh/webhook)  Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [kala](https://github.com/ajvb/kala)  Simplistic, modern, and performant job scheduler.
* [Hey](https://github.com/rakyll/hey)  Hey is a tiny program that sends some load to a web application.
* [aurora](https://github.com/xuri/aurora)  Cross-platform web-based Beanstalkd queue server console.
* [fac](https://github.com/mkchoi212/fac)  Command-line user interface to fix git merge conflicts.
* [awsenv](https://github.com/soniah/awsenv)  Small binary that loads Amazon (AWS) environment variables for a profile.
* [Blast](https://github.com/dave/blast)  A simple tool for API load testing and batch jobs.
* [bombardier](https://github.com/codesenberg/bombardier)  Fast cross-platform HTTP benchmarking tool.
* [bosun](https://github.com/bosun-monitor/bosun)  Time Series Alerting Framework.
* [DepCharge](https://github.com/centerorbit/depcharge)  Helps orchestrating the execution of commands across the many dependencies in larger projects.
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator)  A go library and an executable that produces valid Dockerfiles using various input channels.
* [dogo](https://github.com/liudng/dogo)  Monitoring changes in the source file and automatically compile and run (restart).
* [drone-jenkins](https://github.com/appleboy/drone-jenkins)  Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
* [drone-scp](https://github.com/appleboy/drone-scp)  Copy files and artifacts via SSH using a binary, docker or Drone CI.
* [Dropship](https://github.com/chrismckenzie/dropship)  Tool for deploying code via cdn.
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy)  Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
* [gaia](https://github.com/gaia-pipeline/gaia)  Build powerful pipelines in any programming language.
* [GVM](https://github.com/moovweb/gvm)  GVM provides an interface to manage Go versions.
* [Gitea](https://github.com/go-gitea/gitea)  Fork of Gogs, entirely community driven.
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [go-furnace](https://github.com/go-furnace/go-furnace)  Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate)  Enable your Go applications to self update.
* [gobrew](https://github.com/cryptojuice/gobrew)  gobrew lets you easily switch between multiple versions of go.
* [godbg](https://github.com/sirnewton01/godbg)  Web-based gdb front-end application.
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [gonative](https://github.com/inconshreveable/gonative)  Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
* [govvv](https://github.com/ahmetalpbalkan/govvv)  “go build” wrapper to easily add version information into Go binaries.
* [gox](https://github.com/mitchellh/gox)  Dead simple, no frills Go cross compile tool.
* [goxc](https://github.com/laher/goxc)  build tool for Go, with a focus on cross-compiling and packaging.
* [grapes](https://github.com/yaronsumel/grapes)  Lightweight tool designed to distribute commands over ssh with ease.
* [winrm-cli](https://github.com/masterzen/winrm-cli)  Cli tool to remotely execute commands on Windows machines.

### Other Software

* [borg](https://github.com/crufter/borg)  Terminal based search engine for bash snippets.
* [restic](https://github.com/restic/restic)  De-duplicating backup program.
* [limetext](http://limetext.org/)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [LiteIDE](https://github.com/visualfc/liteide)  LiteIDE is a simple, open source, cross-platform Go IDE.
* [mockingjay](https://github.com/quii/mockingjay-server)  Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.
* [myLG](https://github.com/mehrdadrad/mylg)  Command Line Network Diagnostic tool written in Go.
* [naclpipe](https://github.com/unix4fun/naclpipe)  Simple NaCL EC25519 based crypto pipe tool written in Go.
* [nes](https://github.com/fogleman/nes)  Nintendo Entertainment System (NES) emulator written in Go.
* [orange-cat](https://github.com/noraesae/orange-cat)  Markdown previewer written in Go.
* [Orbit](https://github.com/gulien/orbit)  A simple tool for running commands and generating files from templates.
* [peg](https://github.com/pointlander/peg)  Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
* [scc](https://github.com/boyter/scc)  Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.
* [Leaps](https://github.com/jeffail/leaps)  Pair programming service using Operational Transforms.
* [Seaweed File System](https://github.com/chrislusf/seaweedfs)  Fast, Simple and Scalable Distributed File System with O(1) disk seek.
* [shell2http](https://github.com/msoap/shell2http)  Executing shell commands via http server (for prototyping or remote control).
* [snap](https://github.com/intelsdi-x/snap)  Powerful telemetry framework.
* [Snitch](https://github.com/lucasgomide/snitch)  Simple way to notify your team and many tools when someone has deployed any application via Tsuru.
* [Stack Up](https://github.com/pressly/sup)  Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [term-quiz](https://github.com/crazcalm/term-quiz)  Quizzes for your terminal.
* [toxiproxy](https://github.com/shopify/toxiproxy)  Proxy to simulate network and system conditions for automated tests.
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [vFlow](https://github.com/VerizonDigital/vflow)  High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.
* [lgo](https://github.com/yunabe/lgo)  Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [boxed](https://github.com/tejo/boxed)  Dropbox based blog engine.
* [Duplicacy](https://github.com/gilbertchen/duplicacy)  A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.
* [Cherry](https://github.com/rafael-santiago/cherry)  Tiny webchat server in Go.
* [Circuit](https://github.com/gocircuit/circuit)  Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
* [Comcast](https://github.com/tylertreat/Comcast)  Simulate bad network connections.
* [confd](https://github.com/kelseyhightower/confd)  Manage local application configuration files using templates and data from etcd or consul.
* [croc](https://github.com/schollz/croc)  Easily and securely send files or folders from one computer to another.
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [Documize](https://github.com/documize/community)  Modern wiki software that integrates data from SaaS tools.
* [dp](https://github.com/scryinfo/dp)  Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.
* [drive](https://github.com/odeke-em/drive)  Google Drive client for the commandline.
* [gfile](https://github.com/Antonito/gfile)  Securely transfer files between two computers, without any third party, over WebRTC.
* [joincap](https://github.com/assafmo/joincap)  Command-line utility for merging multiple pcap files together.
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store)  App that displays updates for the Go packages in your GOPATH.
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix)  Video streaming torrent client.
* [GoBoy](https://github.com/Humpheh/goboy)  Nintendo Game Boy Color emulator written in Go.
* [gocc](https://github.com/goccmack/gocc)  Gocc is a compiler kit for Go written in Go.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip)  Chrome extension for Go Doc sites, which shows function description as tooltip at function list.
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
* [Gor](https://github.com/buger/gor)  Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [ide](https://github.com/thestrukture/ide)  Browser accessible IDE. Designed for Go with Go.
* [ipe](https://github.com/dimiro1/ipe)  Open source Pusher server implementation compatible with Pusher client libraries written in GO.
* [wellington](https://github.com/wellington/wellington)  Sass project management tool, extends the language with sprite functions (like Compass).

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [autobench](https://github.com/davecheney/autobench)  Framework to compare the performance between different Go versions.
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app)  Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks)  Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.
* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)  Go HTTP request router benchmark and comparison.
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)  Go web framework benchmark.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)  Benchmarks of Go serialization methods.
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel)  Benchmarks of common basic operations for the Go language.
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark)  Collection of benchmarks for popular Go database/SQL utilities.
* [gospeed](https://github.com/feyeleanor/GoSpeed)  Go micro-benchmarks for calculating the speed of language constructs.
* [kvbench](https://github.com/jimrobinson/kvbench)  Key/Value database benchmark.
* [skynet](https://github.com/atemerev/skynet)  Skynet 1M threads microbenchmark.
* [speedtest-resize](https://github.com/fawick/speedtest-resize)  Compare various Image resize algorithms for the Go language.

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
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly)  in Persian.
* [GoBooks](https://github.com/dariubs/GoBooks)  A curated list of Go books.
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack)  Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector)  Go gopher Vector Data [.ai, .svg].
* [gopher-logos](https://github.com/GolangUA/gopher-logos)  adorable gopher logos.
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gopherize.me](https://github.com/matryer/gopherize.me)  Gopherize yourself.
* [gophers](https://github.com/ashleymcnamara/gophers)  Gopher artworks by Ashley McNamara.
* [gophers](https://github.com/egonelbre/gophers)  Free gophers.
* [gophers](https://github.com/rogeralsing/gophers)  random gopher graphics.
* [gophers](https://github.com/sillecelik/go-gopher)  Gopher amigurumi toy pattern.

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

* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [json2go](https://m-zajac.github.io/json2go)  Advanced JSON to Go struct conversion - online tool.
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [golang-graphics](https://github.com/mholt/golang-graphics)  Collection of Go images, graphics, and art.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Golang Developer Jobs](https://golangjob.xyz)  Developer Jobs exclusivly for Golang related Roles.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp)  Collection of Go projects that needs help. Good place to start your open-source way in Go.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job)  Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
* [go.dev](https://go.dev/)  A hub for Go developers.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness)  List of other amazingly awesome lists.
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang)  Golang ebook intro how to build a web app with golang.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book)  A little e-book on Ethereum Development with Go.
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)  Go's reference card.
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8)  Interactively edit & play Go snippets on your mobile device.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns)  Curated list of Go design patterns, recipes and idioms.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers)  Examples of Golang compared to Node.js for learning.
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/)  Collection of code snippets and tutorials to help tackle every day issues.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests)  Learn Go with test-driven development.
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera Specialization to learn about Go from scratch.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go)  Intro to go for experienced programmers.
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

