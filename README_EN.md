# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[Green]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Yellow]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in the last year"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[GoDoc]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2021-01-27 07:21:28(Synchronize every day)**

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
 
**We have no monthly cost**_, but we have employees **working hard** to maintain the Awesome Go, with money raised we can repay the effort of each person involved! All billing and distribution will be open to the entire community._

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Bot Building](#bot-building)
    - [Build Automation](#build-automation)
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
    - [File Handling](#file-handling)
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
        - [Uncategorized](#uncategorized)
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
    - [WebAssembly](#webassembly)
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
    - [Style Guides](#style-guides)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [Oto](https://github.com/hajimehoshi/oto) **star:766** A low-level library to play sound on multiple platforms.   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/oto)
* [PortAudio](https://github.com/gordonklaus/portaudio) **star:428** Go bindings for the PortAudio audio I/O library.   [![godoc][GoDoc]](https://godoc.org/github.com/gordonklaus/portaudio)
* [music-theory](https://github.com/go-music-theory/music-theory) **star:328** Music theory models in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-music-theory/music-theory)
* [waveform](https://github.com/mdlayher/waveform) **star:322** Go package capable of generating waveform images from audio streams.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/waveform)
* [portmidi](https://github.com/rakyll/portmidi) **star:252** Go bindings for PortMidi.   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/portmidi)
* [id3v2](https://github.com/bogem/id3v2) **star:176** ID3 decoding and encoding library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bogem/id3v2)
* [flac](https://github.com/mewkiz/flac) **star:142** Native Go FLAC encoder/decoder with support for FLAC streams.   [![There was an update last week][Green]](https://github.com/mewkiz/flac)   [![godoc][GoDoc]](https://godoc.org/github.com/mewkiz/flac)
* [mix](https://github.com/go-mix/mix) **star:136** Sequence-based Go-native audio mixer for music apps.   [![godoc][GoDoc]](https://godoc.org/github.com/go-mix/mix)
* [malgo](https://github.com/gen2brain/malgo) **star:134** Mini audio library.
* [go-sox](https://github.com/krig/go-sox) **star:117** libsox bindings for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/krig/go-sox)   [![godoc][GoDoc]](https://godoc.org/github.com/krig/go-sox)
* [mp3](https://github.com/tcolgate/mp3) **star:113** Native Go MP3 decoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tcolgate/mp3)   [![godoc][GoDoc]](https://godoc.org/github.com/tcolgate/mp3)
* [GoAudio](https://github.com/DylanMeeus/GoAudio) **star:107** Native Go Audio Processing Library.   [![There was an update last week][Green]](https://github.com/DylanMeeus/GoAudio)   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/GoAudio)
* [gaad](https://github.com/Comcast/gaad) **star:76** Native Go AAC bitstream parser.   [![godoc][GoDoc]](https://godoc.org/github.com/Comcast/gaad)
* [minimp3](https://github.com/tosone/minimp3) **star:52** Lightweight MP3 decoder library.
* [vorbis](https://github.com/mccoyst/vorbis) **star:28** "Native" Go Vorbis decoder (uses CGO, but has no dependencies).   [![It hasn't been updated in the last year][Yellow]](https://github.com/mccoyst/vorbis)   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/vorbis)
* [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:10** libsamplerate bindings for go.   [![godoc][GoDoc]](https://godoc.org/github.com/dh1tw/gosamplerate)

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) **star:8961** Golang implementation of JSON Web Tokens (JWT).   [![star > 2000][Awesome]](https://github.com/dgrijalva/jwt-go)   [![godoc][GoDoc]](https://godoc.org/github.com/dgrijalva/jwt-go)
* [casbin](https://github.com/hsluoyz/casbin) **star:8510** Authorization library that supports access control models like ACL, RBAC, ABAC.   [![star > 2000][Awesome]](https://github.com/hsluoyz/casbin)   [![There was an update last week][Green]](https://github.com/hsluoyz/casbin)   [![godoc][GoDoc]](https://godoc.org/github.com/hsluoyz/casbin)
* [oauth2](https://github.com/golang/oauth2) **star:3458** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.   [![star > 2000][Awesome]](https://github.com/golang/oauth2)   [![There was an update last week][Green]](https://github.com/golang/oauth2)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/oauth2)
* [goth](https://github.com/markbates/goth) **star:3035** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![star > 2000][Awesome]](https://github.com/markbates/goth)   [![There was an update last week][Green]](https://github.com/markbates/goth)   [![godoc][GoDoc]](https://godoc.org/github.com/markbates/goth)
* [authboss](https://github.com/volatiletech/authboss) **star:2532** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.   [![star > 2000][Awesome]](https://github.com/volatiletech/authboss)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/authboss)
* [loginsrv](https://github.com/tarent/loginsrv) **star:1731** JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.   [![There was an update last week][Green]](https://github.com/tarent/loginsrv)   [![godoc][GoDoc]](https://godoc.org/github.com/tarent/loginsrv)
* [go-jose](https://github.com/square/go-jose) **star:1700** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) **star:1690** Standalone, specification-compliant,  OAuth2 server written in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-oauth2-server)
* [osin](https://github.com/openshift/osin) **star:1644** Golang OAuth2 server library.   [![godoc][GoDoc]](https://godoc.org/github.com/openshift/osin)
* [gologin](https://github.com/dghubble/gologin) **star:1325** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/gologin)
* [gorbac](https://github.com/mikespook/gorbac) **star:1127** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mikespook/gorbac)   [![godoc][GoDoc]](https://godoc.org/github.com/mikespook/gorbac)
* [scs](https://github.com/alexedwards/scs) **star:826** Session Manager for HTTP servers.   [![godoc][GoDoc]](https://godoc.org/github.com/alexedwards/scs)
* [paseto](https://github.com/o1egl/paseto) **star:425** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/paseto)
* [permissions2](https://github.com/xyproto/permissions2) **star:418** Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/permissions2)
* [jwt](https://github.com/pascaldekloe/jwt) **star:236** Lightweight JSON Web Token (JWT) library.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/jwt)
* [jeff](https://github.com/abraithwaite/jeff) **star:222** Simple, flexible, secure and idiomatic web session management with pluggable backends.   [![godoc][GoDoc]](https://godoc.org/github.com/abraithwaite/jeff)
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:203** JWT middleware for Golang http servers with many configuration options.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/jwt-auth)
* [jwt](https://github.com/cristalhq/jwt) **star:201** Safe, simple and fast JSON Web Tokens for Go.   [![There was an update last week][Green]](https://github.com/cristalhq/jwt)   [![godoc][GoDoc]](https://godoc.org/github.com/cristalhq/jwt)
* [httpauth](https://github.com/goji/httpauth) **star:198** HTTP Authentication middleware.   [![godoc][GoDoc]](https://godoc.org/github.com/goji/httpauth)
* [go-guardian](https://github.com/shaj13/go-guardian) **star:170** Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token and Certificate based authentication.   [![godoc][GoDoc]](https://godoc.org/github.com/shaj13/go-guardian)
* [branca](https://github.com/hako/branca) **star:144** Golang implementation of Branca Tokens.   [![godoc][GoDoc]](https://godoc.org/github.com/hako/branca)
* [session](https://github.com/icza/session) **star:104** Go session management for web servers (including support for Google App Engine - GAE).   [![It hasn't been updated in the last year][Yellow]](https://github.com/icza/session)   [![godoc][GoDoc]](https://godoc.org/github.com/icza/session)
* [sessionup](https://github.com/swithek/sessionup) **star:104** Simple, yet effective HTTP session management and identification package.   [![godoc][GoDoc]](https://godoc.org/github.com/swithek/sessionup)
* [jwt](https://github.com/robbert229/jwt) **star:89** Clean and easy to use implementation of JSON Web Tokens (JWT).   [![godoc][GoDoc]](https://godoc.org/github.com/robbert229/jwt)
* [sjwt](https://github.com/brianvoe/sjwt) **star:81** Simple jwt generator and parser.   [![It hasn't been updated in the last year][Yellow]](https://github.com/brianvoe/sjwt)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/sjwt)
* [rbac](https://github.com/zpatrick/rbac) **star:70** Minimalistic RBAC package for Go applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/rbac)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rbac)
* [sessions](https://github.com/adam-hanna/sessions) **star:56** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![godoc][GoDoc]](https://godoc.org/github.com/adam-hanna/sessions)
* [securecookie](https://github.com/chmike/securecookie) **star:43** Efficient secure cookie encoding/decoding.   [![godoc][GoDoc]](https://godoc.org/github.com/chmike/securecookie)
* [otpgo](https://github.com/jltorresm/otpgo) **star:14** Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/jltorresm/otpgo)
* [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) **star:10** Golang library for providing a canonical representation of email address.   [![godoc][GoDoc]](https://godoc.org/github.com/dimuska139/go-email-normalizer)
* [scope](https://github.com/SonicRoshan/scope) **star:10** Easily Manage OAuth2 Scopes In Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/SonicRoshan/scope)   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/scope)
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) **star:9** Go session management using the SessionGate Redis module.   [![It hasn't been updated in the last year][Yellow]](https://github.com/f0rmiga/sessiongate-go)   [![godoc][GoDoc]](https://godoc.org/github.com/f0rmiga/sessiongate-go)
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:6** provides parser of cookies.txt file format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/cookiestxt)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/cookiestxt)

## Bot Building

*Libraries for building and working with bots.*

* [olivia](https://github.com/olivia-ai/olivia) **star:2819** A chatbot built with an artificial neural network.   [![star > 2000][Awesome]](https://github.com/olivia-ai/olivia)   [![godoc][GoDoc]](https://godoc.org/github.com/olivia-ai/olivia)
* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) **star:2610** Simple and clean Telegram bot client.   [![star > 2000][Awesome]](https://github.com/Syfaro/telegram-bot-api)   [![There was an update last week][Green]](https://github.com/Syfaro/telegram-bot-api)   [![godoc][GoDoc]](https://godoc.org/github.com/Syfaro/telegram-bot-api)
* [telebot](https://github.com/tucnak/telebot) **star:1632** Telegram bot framework written in Go.   [![There was an update last week][Green]](https://github.com/tucnak/telebot)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/telebot)
* [go-chat-bot](https://github.com/go-chat-bot/bot) **star:647** IRC, Slack & Telegram bot written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-chat-bot/bot)
* [go-joe](https://joe-bot.net)  A general-purpose bot library inspired by Hubot but written in Go.
* [Kelp](https://github.com/stellar/kelp) **star:489** official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.   [![There was an update last week][Green]](https://github.com/stellar/kelp)   [![godoc][GoDoc]](https://godoc.org/github.com/stellar/kelp)
* [slacker](https://github.com/shomali11/slacker) **star:487** Easy to use framework to create Slack bots.   [![There was an update last week][Green]](https://github.com/shomali11/slacker)   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/slacker)
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) **star:406** A golang implementation of a console-based trading bot for cryptocurrency exchanges.   [![There was an update last week][Green]](https://github.com/saniales/golang-crypto-trading-bot)   [![godoc][GoDoc]](https://godoc.org/github.com/saniales/golang-crypto-trading-bot)
* [tbot](https://github.com/yanzay/tbot) **star:297** Telegram bot server with API similar to net/http.   [![godoc][GoDoc]](https://godoc.org/github.com/yanzay/tbot)
* [go-sarah](https://github.com/oklahomer/go-sarah) **star:188** Framework to build bot for desired chat services including LINE, Slack, Gitter and more.   [![godoc][GoDoc]](https://godoc.org/github.com/oklahomer/go-sarah)
* [Tenyks](https://github.com/kyleterry/tenyks) **star:169** Service oriented IRC bot using Redis and JSON for messaging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kyleterry/tenyks)   [![godoc][GoDoc]](https://godoc.org/github.com/kyleterry/tenyks)
* [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:145** Libary to write bots for twitch.tv chat   [![godoc][GoDoc]](https://godoc.org/github.com/gempir/go-twitch-irc)
* [hanu](https://github.com/sbstjn/hanu) **star:127** Framework for writing Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/hanu)
* [go-tgbot](https://github.com/olebedev/go-tgbot) **star:102** Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/olebedev/go-tgbot)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-tgbot)
* [margelet](https://github.com/zhulik/margelet) **star:64** Framework for building Telegram bots.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulik/margelet)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulik/margelet)
* [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) **star:38** A Discord bot for managing ephemeral roles based upon voice channel member presence.   [![There was an update last week][Green]](https://github.com/ewohltman/ephemeral-roles)   [![godoc][GoDoc]](https://godoc.org/github.com/ewohltman/ephemeral-roles)
* [slackscot](https://github.com/alexandre-normand/slackscot) **star:37** Another framework for building Slack bots.   [![godoc][GoDoc]](https://godoc.org/github.com/alexandre-normand/slackscot)
* [govkbot](https://github.com/nikepan/govkbot) **star:35** Simple Go [VK](https://vk.com) bot library.   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/govkbot)
* [micha](https://github.com/onrik/micha) **star:18** Go Library for Telegram bot api.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/micha)

## Build Automation

*Libraries and tools helping with build automation.*

* [realize](https://github.com/tockins/realize) **star:3904** Go build system with file watchers and live reload. Run, build and watch file changes with custom paths.   [![star > 2000][Awesome]](https://github.com/tockins/realize)   [![godoc][GoDoc]](https://godoc.org/github.com/tockins/realize)
* [Task](https://github.com/go-task/task) **star:3071** simple "Make" alternative.   [![star > 2000][Awesome]](https://github.com/go-task/task)   [![There was an update last week][Green]](https://github.com/go-task/task)   [![godoc][GoDoc]](https://godoc.org/github.com/go-task/task)
* [mmake](https://github.com/tj/mmake) **star:1534** Modern Make.   [![godoc][GoDoc]](https://godoc.org/github.com/tj/mmake)
* [1build](https://github.com/gopinath-langote/1build) **star:85** Command line tool to frictionlessly manage project-specific commands.   [![godoc][GoDoc]](https://godoc.org/github.com/gopinath-langote/1build)
* [taskctl](https://github.com/taskctl/taskctl) **star:78** Concurrent task runner.   [![godoc][GoDoc]](https://godoc.org/github.com/taskctl/taskctl)
* [gaper](https://github.com/maxcnunes/gaper) **star:45** Builds and restarts a Go project when it crashes or some watched file changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/maxcnunes/gaper)   [![godoc][GoDoc]](https://godoc.org/github.com/maxcnunes/gaper)
* [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.
* [taskflow](https://github.com/pellared/taskflow) **star:41** Create build pipelines in Go.   [![There was an update last week][Green]](https://github.com/pellared/taskflow)   [![godoc][GoDoc]](https://godoc.org/github.com/pellared/taskflow)

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) **star:20164** Commander for modern Go CLI interactions.   [![star > 2000][Awesome]](https://github.com/spf13/cobra)   [![There was an update last week][Green]](https://github.com/spf13/cobra)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/cobra)
* [urfave/cli](https://github.com/urfave/cli) **star:15168** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![star > 2000][Awesome]](https://github.com/urfave/cli)   [![There was an update last week][Green]](https://github.com/urfave/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/cli)
* [kingpin](https://github.com/alecthomas/kingpin) **star:2993** Command line and flag parser supporting sub commands.   [![star > 2000][Awesome]](https://github.com/alecthomas/kingpin)   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/kingpin)
* [Dnote](https://github.com/dnote/dnote) **star:1990** A simple command line notebook with multi-device sync.   [![There was an update last week][Green]](https://github.com/dnote/dnote)   [![godoc][GoDoc]](https://godoc.org/github.com/dnote/dnote)
* [go-flags](https://github.com/jessevdk/go-flags) **star:1851** go command line option parser.   [![godoc][GoDoc]](https://godoc.org/github.com/jessevdk/go-flags)
* [pflag](https://github.com/spf13/pflag) **star:1347** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/pflag)
* [mitchellh/cli](https://github.com/mitchellh/cli) **star:1262** Go library for implementing command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/cli)
* [go-arg](https://github.com/alexflint/go-arg) **star:1028** Struct-based argument parsing in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexflint/go-arg)
* [liner](https://github.com/peterh/liner) **star:749** Go readline-like library for command-line interfaces.   [![godoc][GoDoc]](https://godoc.org/github.com/peterh/liner)
* [complete](https://github.com/posener/complete) **star:747** Write bash completions in Go + Go command bash completion.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/complete)
* [mow.cli](https://github.com/jawher/mow.cli) **star:716** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   [![godoc][GoDoc]](https://godoc.org/github.com/jawher/mow.cli)
* [flaggy](https://github.com/integrii/flaggy) **star:709** A robust and idiomatic flags package with excellent subcommand support.   [![godoc][GoDoc]](https://godoc.org/github.com/integrii/flaggy)
* [cli](https://github.com/mkideal/cli) **star:551** Feature-rich and easy to use command-line package based on golang struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/mkideal/cli)
* [ops](https://github.com/nanovms/ops) **star:540** Unikernel Builder/Orchestrator.   [![There was an update last week][Green]](https://github.com/nanovms/ops)   [![godoc][GoDoc]](https://godoc.org/github.com/nanovms/ops)
* [argparse](https://github.com/akamensky/argparse) **star:276** Command line argument parser inspired by Python's argparse module.   [![godoc][GoDoc]](https://godoc.org/github.com/akamensky/argparse)
* [commandeer](https://github.com/jaffee/commandeer) **star:142** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   [![godoc][GoDoc]](https://godoc.org/github.com/jaffee/commandeer)
* [wmenu](https://github.com/dixonwille/wmenu) **star:126** Easy to use menu structure for cli applications that prompts users to make choices.   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wmenu)
* [sflags](https://github.com/octago/sflags) **star:122** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/octago/sflags)
* [flag](https://github.com/cosiner/flag) **star:111** Simple but powerful command line option parsing library for Go supporting subcommand.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/flag)
* [ukautz/clif](https://github.com/ukautz/clif) **star:106** Small command line interface framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ukautz/clif)   [![godoc][GoDoc]](https://godoc.org/github.com/ukautz/clif)
* [job](https://github.com/liujianping/job) **star:94** JOB, make your short-term command as a long-term job.
* [cli](https://github.com/teris-io/cli) **star:77** Simple and complete API for building command line interfaces in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/cli)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/cli)
* [climax](http://github.com/tucnak/climax)  Alternative CLI with "human face", in spirit of Go command.
* [env](https://github.com/codingconcepts/env) **star:71** Tag-based environment configuration for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/env)
* [clîr](https://github.com/leaanthony/clir) **star:58** A Simple and Clear CLI library. Dependency free.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/clir)
* [cmdr](https://github.com/hedzr/cmdr) **star:55** A POSIX/GNU style, getopt-like command-line UI Go library.   [![There was an update last week][Green]](https://github.com/hedzr/cmdr)   [![godoc][GoDoc]](https://godoc.org/github.com/hedzr/cmdr)
* [gocmd](https://github.com/devfacet/gocmd) **star:49** Go library for building command line applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/devfacet/gocmd)   [![godoc][GoDoc]](https://godoc.org/github.com/devfacet/gocmd)
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
* [wlog](https://github.com/dixonwille/wlog) **star:46** Simple logging interface that supports cross-platform color and concurrency.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/wlog)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/wlog)
* [strumt](https://github.com/antham/strumt) **star:40** Library to create prompt chain.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/strumt)
* [flagvar](https://github.com/sgreben/flagvar) **star:36** A collection of flag argument types for Go's standard `flag` package.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/flagvar)
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:31** Go option parser inspired on the flexibility of Perl’s GetOpt::Long.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidGamba/go-getoptions)
* [cmd](https://github.com/posener/cmd) **star:30** Extends the standard `flag` package to support sub commands and more in idomatic way.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/cmd)
* [argv](https://github.com/cosiner/argv) **star:27** Go library to split command line string as arguments array using the bash syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/argv)
* [go-commander](https://github.com/yitsushi/go-commander) **star:20** Go library to simplify CLI workflow.   [![godoc][GoDoc]](https://godoc.org/github.com/yitsushi/go-commander)
* [sand](https://github.com/Zaba505/sand) **star:12** Simple API for creating interpreters and so much more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Zaba505/sand)   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/sand)
* [ts](https://github.com/liujianping/ts) **star:12** Timestamp convert & compare tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/liujianping/ts)   [![godoc][GoDoc]](https://godoc.org/github.com/liujianping/ts)

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) **star:10565** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![star > 2000][Awesome]](https://github.com/gizak/termui)   [![There was an update last week][Green]](https://github.com/gizak/termui)   [![godoc][GoDoc]](https://godoc.org/github.com/gizak/termui)
* [gocui](https://github.com/jroimartin/gocui) **star:6905** Minimalist Go library aimed at creating Console User Interfaces.   [![star > 2000][Awesome]](https://github.com/jroimartin/gocui)   [![godoc][GoDoc]](https://godoc.org/github.com/jroimartin/gocui)
* [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
* [termbox-go](https://github.com/nsf/termbox-go) **star:3944** Termbox is a library for creating cross-platform text-based interfaces.   [![star > 2000][Awesome]](https://github.com/nsf/termbox-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/termbox-go)
* [go-prompt](https://github.com/c-bata/go-prompt) **star:3679** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![star > 2000][Awesome]](https://github.com/c-bata/go-prompt)   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/go-prompt)
* [uiprogress](https://github.com/gosuri/uiprogress) **star:1746** Flexible library to render progress bars in terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uiprogress)
* [asciigraph](https://github.com/guptarohit/asciigraph) **star:1532** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/guptarohit/asciigraph)
* [termdash](https://github.com/mum4k/termdash) **star:1525** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][GoDoc]](https://godoc.org/github.com/mum4k/termdash)
* [progressbar](https://github.com/schollz/progressbar) **star:1518** Basic thread-safe progress bar that works in every OS.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/progressbar)
* [uilive](https://github.com/gosuri/uilive) **star:1259** Library for updating terminal output in realtime.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uilive)
* [mpb](https://github.com/vbauerster/mpb) **star:1184** Multi progress bar for terminal applications.   [![godoc][GoDoc]](https://godoc.org/github.com/vbauerster/mpb)
* [aurora](https://github.com/logrusorgru/aurora) **star:1020** ANSI terminal colors that supports fmt.Printf/Sprintf.   [![godoc][GoDoc]](https://godoc.org/github.com/logrusorgru/aurora)
* [gookit/color](https://github.com/gookit/color) **star:755** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/color)   [![Contains Chinese documents][CN]](https://github.com/gookit/color)
* [uitable](https://github.com/gosuri/uitable) **star:610** Library to improve readability in terminal apps using tabular data.   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/uitable)
* [go-colorable](https://github.com/mattn/go-colorable) **star:498** Colorable writer for windows.   [![There was an update last week][Green]](https://github.com/mattn/go-colorable)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-colorable)
* [go-isatty](https://github.com/mattn/go-isatty) **star:481** isatty for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-isatty)
* [pterm](https://github.com/pterm/pterm) **star:389** A library to beautify console output on every platform with many combinable components.   [![godoc][GoDoc]](https://godoc.org/github.com/pterm/pterm)
* [chalk](https://github.com/ttacon/chalk) **star:356** Intuitive package for prettifying terminal/console output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ttacon/chalk)   [![godoc][GoDoc]](https://godoc.org/github.com/ttacon/chalk)
* [simpletable](https://github.com/alexeyco/simpletable) **star:281** Simple tables in terminal with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/simpletable)
* [tabby](https://github.com/cheynewallace/tabby) **star:277** A tiny library for super simple Golang tables.   [![godoc][GoDoc]](https://godoc.org/github.com/cheynewallace/tabby)
* [go-colortext](https://github.com/daviddengcn/go-colortext) **star:207** Go library for color output in terminals.   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-colortext)
* [yacspin](https://github.com/theckman/yacspin) **star:133** Yet Another CLi Spinner package, for working with terminal spinners.   [![godoc][GoDoc]](https://godoc.org/github.com/theckman/yacspin)
* [cfmt](https://github.com/mingrammer/cfmt) **star:77** Contextual fmt inspired by bootstrap color classes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mingrammer/cfmt)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/cfmt)
* [tabular](https://github.com/InVisionApp/tabular) **star:49** Print ASCII tables from command line utilities without the need to pass large sets of data to the API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/tabular)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/tabular)
* [ctc](https://github.com/wzshiming/ctc) **star:28** The non-invasive cross-platform terminal color library does not need to modify the Print method.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/ctc)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/ctc)
* [colourize](https://github.com/TreyBastian/colourize) **star:23** Go library for ANSI colour text in terminals.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TreyBastian/colourize)   [![godoc][GoDoc]](https://godoc.org/github.com/TreyBastian/colourize)
* [cfmt](https://github.com/i582/cfmt) **star:19** Simple and convenient formatted stylized output fully compatible with fmt library.   [![godoc][GoDoc]](https://godoc.org/github.com/i582/cfmt)
* [go-ataman](https://github.com/workanator/go-ataman) **star:9** Go library for rendering ANSI colored text templates in terminals.   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-ataman)
* [table](https://github.com/tomlazar/table) **star:1** Small library for terminal color based tables .   [![godoc][GoDoc]](https://godoc.org/github.com/tomlazar/table)

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) **star:14557** Go configuration with fangs.   [![star > 2000][Awesome]](https://github.com/spf13/viper)   [![There was an update last week][Green]](https://github.com/spf13/viper)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/viper)
* [godotenv](https://github.com/joho/godotenv) **star:3534** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![star > 2000][Awesome]](https://github.com/joho/godotenv)   [![godoc][GoDoc]](https://godoc.org/github.com/joho/godotenv)
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:3424** Go library for managing configuration data from environment variables.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/envconfig)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/envconfig)
* [ini](https://github.com/go-ini/ini) **star:2346** Go package to read and write INI files.   [![star > 2000][Awesome]](https://github.com/go-ini/ini)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ini/ini)
* [env](https://github.com/caarlos0/env) **star:1760** Parse environment variables to Go structs (with defaults).   [![There was an update last week][Green]](https://github.com/caarlos0/env)   [![godoc][GoDoc]](https://godoc.org/github.com/caarlos0/env)
* [konfig](https://github.com/lalamove/konfig) **star:592** Composable, observable and performant config handling for Go for the distributed processing era.   [![godoc][GoDoc]](https://godoc.org/github.com/lalamove/konfig)
* [koanf](https://github.com/knadh/koanf) **star:406** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![There was an update last week][Green]](https://github.com/knadh/koanf)   [![godoc][GoDoc]](https://godoc.org/github.com/knadh/koanf)
* [confita](https://github.com/heetch/confita) **star:338** Load configuration in cascade from multiple backends into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/heetch/confita)
* [store](https://github.com/tucnak/store) **star:255** Lightweight configuration manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tucnak/store)   [![godoc][GoDoc]](https://godoc.org/github.com/tucnak/store)
* [config](https://github.com/JeremyLoy/config) **star:251** Cloud native application configuration. Bind ENV to structs in only two lines.   [![godoc][GoDoc]](https://godoc.org/github.com/JeremyLoy/config)
* [config](https://github.com/olebedev/config) **star:235** JSON or YAML configuration wrapper with environment variables and flags parsing.   [![It hasn't been updated in the last year][Yellow]](https://github.com/olebedev/config)   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/config)
* [hjson](https://github.com/hjson/hjson-go) **star:228** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![godoc][GoDoc]](https://godoc.org/github.com/hjson/hjson-go)
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:217** Minimalistic configuration reader (from files, ENV, and wherever you want).   [![godoc][GoDoc]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
* [gookit/config](https://github.com/gookit/config) **star:205** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![There was an update last week][Green]](https://github.com/gookit/config)   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/config)   [![Contains Chinese documents][CN]](https://github.com/gookit/config)
* [joshbetz/config](https://github.com/joshbetz/config) **star:203** Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/joshbetz/config)   [![godoc][GoDoc]](https://godoc.org/github.com/joshbetz/config)
* [envconfig](https://github.com/vrischmann/envconfig) **star:192** Read your configuration from environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/vrischmann/envconfig)
* [aconfig](https://github.com/cristalhq/aconfig) **star:171** Simple, useful and opinionated config loader.   [![godoc][GoDoc]](https://godoc.org/github.com/cristalhq/aconfig)
* [goConfig](https://github.com/crgimenes/goConfig) **star:145** Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.   [![godoc][GoDoc]](https://godoc.org/github.com/crgimenes/goConfig)
* [fig](https://github.com/kkyr/fig) **star:139** Tiny library for reading configuration from a file and from environment variables (with validation & defaults).   [![godoc][GoDoc]](https://godoc.org/github.com/kkyr/fig)
* [gcfg](https://github.com/go-gcfg/gcfg) **star:139** read INI-style configuration files into Go structs; supports user-defined types and subsections.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gcfg/gcfg)
* [config](https://github.com/golobby/config) **star:112** A lightweight yet powerful config package for Go projects.   [![There was an update last week][Green]](https://github.com/golobby/config)   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/config)
* [envh](https://github.com/antham/envh) **star:94** Helpers to manage environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/envh)
* [envcfg](https://github.com/tomazk/envcfg) **star:93** Un-marshaling environment variables to Go structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomazk/envcfg)   [![godoc][GoDoc]](https://godoc.org/github.com/tomazk/envcfg)
* [harvester](https://github.com/beatlabs/harvester) **star:77** Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration.   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/harvester)
* [xdg](https://github.com/adrg/xdg) **star:73** Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).   [![There was an update last week][Green]](https://github.com/adrg/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/xdg)
* [configuro](https://github.com/sherifabdlnaby/configuro) **star:68** opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications.   [![godoc][GoDoc]](https://godoc.org/github.com/sherifabdlnaby/configuro)
* [xdg](https://github.com/OpenPeeDeeP/xdg) **star:61** Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).   [![godoc][GoDoc]](https://godoc.org/github.com/OpenPeeDeeP/xdg)
* [gofigure](https://github.com/ian-kent/gofigure) **star:59** Go application configuration made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/gofigure)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/gofigure)
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
* [configure](https://github.com/paked/configure) **star:55** Provides configuration through multiple sources, including JSON, flags and environment variables.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paked/configure)   [![godoc][GoDoc]](https://godoc.org/github.com/paked/configure)
* [configuration](https://github.com/BoRuDar/configuration) **star:37** Library for initializing configuration structs from env variables, files, flags and 'default' tag.   [![godoc][GoDoc]](https://godoc.org/github.com/BoRuDar/configuration)
* [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) **star:37** Go package that fetches parameters from AWS System Manager - Parameter Store.   [![godoc][GoDoc]](https://godoc.org/github.com/PaddleHQ/go-aws-ssm)
* [ingo](https://github.com/schachmat/ingo) **star:32** Flags persisted in an ini-like config file.   [![It hasn't been updated in the last year][Yellow]](https://github.com/schachmat/ingo)   [![godoc][GoDoc]](https://godoc.org/github.com/schachmat/ingo)
* [go-up](https://github.com/ufoscout/go-up) **star:29** A simple configuration library with recursive placeholders resolution and no magic.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ufoscout/go-up)   [![godoc][GoDoc]](https://godoc.org/github.com/ufoscout/go-up)
* [mini](https://github.com/sasbury/mini) **star:28** Golang package for parsing ini-style configuration files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sasbury/mini)   [![godoc][GoDoc]](https://godoc.org/github.com/sasbury/mini)
* [hocon](https://github.com/gurkankaymak/hocon) **star:20** Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files.   [![godoc][GoDoc]](https://godoc.org/github.com/gurkankaymak/hocon)
* [conflate](https://github.com/the4thamigo-uk/conflate) **star:19** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.   [![godoc][GoDoc]](https://godoc.org/github.com/the4thamigo-uk/conflate)
* [genv](https://github.com/sakirsensoy/genv) **star:18** Read environment variables easily with dotenv support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sakirsensoy/genv)   [![godoc][GoDoc]](https://godoc.org/github.com/sakirsensoy/genv)
* [envconf](https://github.com/ian-kent/envconf) **star:10** Configuration from environment.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/envconf)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/envconf)
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) **star:8** Go utility for loading configuration parameters from AWS SSM (Parameter Store).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-ssm-config)
* [swap](https://github.com/oblq/swap) **star:4** Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env).   [![godoc][GoDoc]](https://godoc.org/github.com/oblq/swap)
* [go-ini](https://github.com/subpop/go-ini) **star:3** A Go package that marshals and unmarshals INI-files.   [![godoc][GoDoc]](https://godoc.org/github.com/subpop/go-ini)
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:3** Simple useful package for read environment variables.   [![godoc][GoDoc]](https://godoc.org/github.com/nasermirzaei89/env)
* [onion](http://github.com/goraz/onion)  Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.
* [typenv](https://github.com/diegomarangoni/typenv) **star:3** Minimalistic, zero dependency, typed environment variables library.   [![godoc][GoDoc]](https://godoc.org/github.com/diegomarangoni/typenv)

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) **star:22442** Drone is a Continuous Integration platform built on Docker, written in Go.   [![star > 2000][Awesome]](https://github.com/drone/drone)   [![There was an update last week][Green]](https://github.com/drone/drone)   [![godoc][GoDoc]](https://godoc.org/github.com/drone/drone)
* [CDS](https://github.com/ovh/cds) **star:3243** Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.   [![star > 2000][Awesome]](https://github.com/ovh/cds)   [![There was an update last week][Green]](https://github.com/ovh/cds)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/cds)
* [goveralls](https://github.com/mattn/goveralls) **star:676** Go integration for Coveralls.io continuous code coverage tracking system.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/goveralls)
* [overalls](https://github.com/go-playground/overalls) **star:105** Multi-Package go project coverprofile for tools like goveralls.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/overalls)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/overalls)
* [duci](https://github.com/duck8823/duci) **star:62** A simple ci server no needs domain specific languages.   [![There was an update last week][Green]](https://github.com/duck8823/duci)   [![godoc][GoDoc]](https://godoc.org/github.com/duck8823/duci)
* [gomason](https://github.com/nikogura/gomason) **star:50** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/gomason)
* [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:14** Recursive coverage testing tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][GoDoc]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) **star:438** Pure Go CSS Preprocessor.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/gcss)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/gcss)
* [go-libsass](https://github.com/wellington/go-libsass) **star:167** Go wrapper to the 100% Sass compatible libsass project.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) **star:9529** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![star > 2000][Awesome]](https://github.com/emirpasic/gods)   [![There was an update last week][Green]](https://github.com/emirpasic/gods)   [![godoc][GoDoc]](https://godoc.org/github.com/emirpasic/gods)
* [go-datastructures](https://github.com/Workiva/go-datastructures) **star:5892** Collection of useful, performant, and thread-safe data structures.   [![star > 2000][Awesome]](https://github.com/Workiva/go-datastructures)   [![godoc][GoDoc]](https://godoc.org/github.com/Workiva/go-datastructures)
* [golang-set](https://github.com/deckarep/golang-set) **star:1795** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/golang-set)
* [gota](https://github.com/kniren/gota) **star:1495** Implementation of dataframes, series, and data wrangling methods for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/kniren/gota)
* [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1323** Probabilistic data structures for processing continuous, unbounded streams.   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/BoomFilters)
* [roaring](https://github.com/RoaringBitmap/roaring) **star:1084** Go package implementing compressed bitsets.   [![There was an update last week][Green]](https://github.com/RoaringBitmap/roaring)   [![godoc][GoDoc]](https://godoc.org/github.com/RoaringBitmap/roaring)
* [willf/bloom](https://github.com/willf/bloom) **star:950** Go package implementing Bloom filters.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bloom)
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:748** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
* [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:736** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![It hasn't been updated in the last year][Yellow]](https://github.com/axiomhq/hyperloglog)   [![godoc][GoDoc]](https://godoc.org/github.com/axiomhq/hyperloglog)
* [gocache](https://github.com/eko/gocache) **star:730** A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.   [![There was an update last week][Green]](https://github.com/eko/gocache)   [![godoc][GoDoc]](https://godoc.org/github.com/eko/gocache)
* [bitset](https://github.com/willf/bitset) **star:645** Go package implementing bitsets.   [![godoc][GoDoc]](https://godoc.org/github.com/willf/bitset)
* [trie](https://github.com/derekparker/trie) **star:517** Trie implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/trie)
* [algorithms](https://github.com/shady831213/algorithms) **star:465** Algorithms and data structures.CLRS study.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shady831213/algorithms)   [![godoc][GoDoc]](https://godoc.org/github.com/shady831213/algorithms)
* [go-geoindex](https://github.com/hailocab/go-geoindex) **star:327** In-memory geo index.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hailocab/go-geoindex)   [![godoc][GoDoc]](https://godoc.org/github.com/hailocab/go-geoindex)
* [go-edlib](https://github.com/hbollon/go-edlib) **star:245** Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode.   [![godoc][GoDoc]](https://godoc.org/github.com/hbollon/go-edlib)
* [gostl](https://github.com/liyue201/gostl) **star:232** Data structure and algorithm library for go, designed to provide functions similar to C++ STL.   [![godoc][GoDoc]](https://godoc.org/github.com/liyue201/gostl)
* [merkletree](https://github.com/cbergoon/merkletree) **star:231** Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cbergoon/merkletree)   [![godoc][GoDoc]](https://godoc.org/github.com/cbergoon/merkletree)
* [hilbert](https://github.com/google/hilbert) **star:225** Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.   [![It hasn't been updated in the last year][Yellow]](https://github.com/google/hilbert)   [![godoc][GoDoc]](https://godoc.org/github.com/google/hilbert)
* [goskiplist](https://github.com/ryszard/goskiplist) **star:222** Skip list implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ryszard/goskiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/ryszard/goskiplist)
* [ttlcache](https://github.com/ReneKroon/ttlcache) **star:198** In-memory string-interface{} cache with various time-based expiration options and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/ReneKroon/ttlcache)
* [deque](https://github.com/gammazero/deque) **star:196** Fast ring-buffer deque (double-ended queue).   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/deque)
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:171** Go implementation of Adaptive Radix Tree.   [![godoc][GoDoc]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
* [binpacker](https://github.com/zhuangsirui/binpacker) **star:162** Binary packer and unpacker helps user build custom binary stream.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhuangsirui/binpacker)   [![godoc][GoDoc]](https://godoc.org/github.com/zhuangsirui/binpacker)
* [skiplist](https://github.com/MauriceGit/skiplist) **star:138** Very fast Go Skiplist implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/MauriceGit/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/MauriceGit/skiplist)
* [bloom](https://github.com/zhenjl/bloom) **star:137** Bloom filters implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/bloom)
* [iter](https://github.com/disksing/iter) **star:130** Go implementation of C++ STL iterators and algorithms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/disksing/iter)   [![godoc][GoDoc]](https://godoc.org/github.com/disksing/iter)   [![Contains Chinese documents][CN]](https://github.com/disksing/iter)
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:124** Concurrent FIFO queue.   [![godoc][GoDoc]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
* [levenshtein](https://github.com/agnivade/levenshtein) **star:116** Implementation to calculate levenshtein distance in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/levenshtein)
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:112** Region quadtrees with efficient point location and neighbour finding.   [![godoc][GoDoc]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
* [ring](https://github.com/TheTannerRyan/ring) **star:112** Go implementation of a high performance, thread safe bloom filter.   [![godoc][GoDoc]](https://godoc.org/github.com/TheTannerRyan/ring)
* [encoding](https://github.com/zhenjl/encoding) **star:104** Integer Compression Libraries for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhenjl/encoding)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/encoding)
* [bit](https://github.com/yourbasic/bit) **star:98** Golang set data structure with bonus bit-twiddling functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bit)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bit)
* [conjungo](https://github.com/InVisionApp/conjungo) **star:92** A small, powerful and flexible merge library.   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/conjungo)
* [remember-go](https://github.com/rocketlaunchr/remember-go) **star:72** A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory).   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/remember-go)
* [skiplist](https://github.com/gansidui/skiplist) **star:71** Skiplist implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/skiplist)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/skiplist)
* [bloom](https://github.com/yourbasic/bloom) **star:58** Golang Bloom filter implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/bloom)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/bloom)
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:58** Fast in-memory key:value store/cache library. Pointer caches.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OrlovEvgeny/go-mcache)   [![godoc][GoDoc]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
* [levenshtein](https://github.com/agext/levenshtein) **star:55** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/levenshtein)
* [count-min-log](https://github.com/seiflotfy/count-min-log) **star:52** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/count-min-log)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/count-min-log)
* [crunch](https://github.com/superwhiskers/crunch) **star:42** Go package implementing buffers for handling various datatypes easily.   [![There was an update last week][Green]](https://github.com/superwhiskers/crunch)   [![godoc][GoDoc]](https://godoc.org/github.com/superwhiskers/crunch)
* [nan](https://github.com/kak-tus/nan) **star:33** Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers.   [![godoc][GoDoc]](https://godoc.org/github.com/kak-tus/nan)
* [concurrent-writer](https://github.com/free/concurrent-writer) **star:30** Highly concurrent drop-in replacement for `bufio.Writer`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/free/concurrent-writer)   [![godoc][GoDoc]](https://godoc.org/github.com/free/concurrent-writer)
* [goset](https://github.com/zoumo/goset) **star:29** A useful Set collection implementation for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/zoumo/goset)
* [hide](https://github.com/emvi/hide) **star:29** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/hide)
* [pipeline](https://github.com/hyfather/pipeline) **star:28** An implementation of pipelines with fan-in and fan-out.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyfather/pipeline)   [![godoc][GoDoc]](https://godoc.org/github.com/hyfather/pipeline)
* [deque](https://github.com/edwingeng/deque) **star:23** A highly optimized double-ended queue.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/deque)
* [timedmap](https://github.com/zekroTJA/timedmap) **star:22** Map with expiring key-value pairs.   [![godoc][GoDoc]](https://godoc.org/github.com/zekroTJA/timedmap)
* [typ](https://github.com/gurukami/typ) **star:22** Null Types, Safe primitive type conversion and fetching value from complex structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gurukami/typ)   [![godoc][GoDoc]](https://godoc.org/github.com/gurukami/typ)
* [dict](https://github.com/srfrog/dict) **star:18** Python-like dictionaries (dict) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/dict)
* [go-ef](https://github.com/amallia/go-ef) **star:17** A Go implementation of the Elias-Fano encoding.   [![It hasn't been updated in the last year][Yellow]](https://github.com/amallia/go-ef)   [![godoc][GoDoc]](https://godoc.org/github.com/amallia/go-ef)
* [null](https://github.com/emvi/null) **star:15** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/emvi/null)
* [mspm](https://github.com/BlackRabbitt/mspm) **star:13** Multi-String Pattern Matching Algorithm for information retrieval.   [![It hasn't been updated in the last year][Yellow]](https://github.com/BlackRabbitt/mspm)   [![godoc][GoDoc]](https://godoc.org/github.com/BlackRabbitt/mspm)
* [set](https://github.com/StudioSol/set) **star:12** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][GoDoc]](https://godoc.org/github.com/StudioSol/set)
* [ptrie](https://github.com/viant/ptrie) **star:11** An implementation of prefix tree.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/ptrie)
* [treap](https://github.com/perdata/treap) **star:10** Persistent, fast ordered map using tree heaps.   [![It hasn't been updated in the last year][Yellow]](https://github.com/perdata/treap)   [![godoc][GoDoc]](https://godoc.org/github.com/perdata/treap)
* [cmap](https://github.com/lrita/cmap) **star:9** a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/cmap)
* [gofal](https://github.com/xxjwxc/gofal) **star:9** fractional api for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xxjwxc/gofal)   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gofal)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gofal)
* [parsefields](https://github.com/MonaxGT/parsefields) **star:5** Tools for parse JSON-like logs for collecting unique fields and events.   [![It hasn't been updated in the last year][Yellow]](https://github.com/MonaxGT/parsefields)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/parsefields)
* [goterator](https://github.com/yaa110/goterator) **star:3** Iterator implementation to provide map and reduce functionalities.   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/goterator)
* [slices](https://github.com/srfrog/slices) **star:2** Functions that operate on slices; like `package strings` but adapted to work with slices.   [![godoc][GoDoc]](https://godoc.org/github.com/srfrog/slices)

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) **star:34961** Monitoring system and time series database.   [![star > 2000][Awesome]](https://github.com/prometheus/prometheus)   [![There was an update last week][Green]](https://github.com/prometheus/prometheus)   [![godoc][GoDoc]](https://godoc.org/github.com/prometheus/prometheus)
* [tidb](https://github.com/pingcap/tidb) **star:26659** TiDB is a distributed SQL database. Inspired by the design of Google F1.   [![star > 2000][Awesome]](https://github.com/pingcap/tidb)   [![There was an update last week][Green]](https://github.com/pingcap/tidb)   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/tidb)   [![Contains Chinese documents][CN]](https://github.com/pingcap/tidb)
* [influxdb](https://github.com/influxdb/influxdb) **star:20374** Scalable datastore for metrics, events, and real-time analytics.   [![star > 2000][Awesome]](https://github.com/influxdb/influxdb)   [![There was an update last week][Green]](https://github.com/influxdb/influxdb)   [![godoc][GoDoc]](https://godoc.org/github.com/influxdb/influxdb)
* [cockroach](https://github.com/cockroachdb/cockroach) **star:19800** Scalable, Geo-Replicated, Transactional Datastore.   [![star > 2000][Awesome]](https://github.com/cockroachdb/cockroach)   [![There was an update last week][Green]](https://github.com/cockroachdb/cockroach)   [![godoc][GoDoc]](https://godoc.org/github.com/cockroachdb/cockroach)
* [dgraph](https://github.com/dgraph-io/dgraph) **star:14716** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![star > 2000][Awesome]](https://github.com/dgraph-io/dgraph)   [![There was an update last week][Green]](https://github.com/dgraph-io/dgraph)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/dgraph)
* [groupcache](https://github.com/golang/groupcache) **star:9729** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![star > 2000][Awesome]](https://github.com/golang/groupcache)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/groupcache)
* [badger](https://github.com/dgraph-io/badger) **star:8743** Fast key-value store in Go.   [![star > 2000][Awesome]](https://github.com/dgraph-io/badger)   [![There was an update last week][Green]](https://github.com/dgraph-io/badger)   [![godoc][GoDoc]](https://godoc.org/github.com/dgraph-io/badger)
* [rqlite](https://github.com/rqlite/rqlite) **star:7370** The lightweight, distributed, relational database built on SQLite.   [![star > 2000][Awesome]](https://github.com/rqlite/rqlite)   [![There was an update last week][Green]](https://github.com/rqlite/rqlite)   [![godoc][GoDoc]](https://godoc.org/github.com/rqlite/rqlite)
* [go-cache](https://github.com/pmylund/go-cache) **star:4638** In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.   [![star > 2000][Awesome]](https://github.com/pmylund/go-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/pmylund/go-cache)
* [BigCache](https://github.com/allegro/bigcache) **star:4602** Efficient key/value cache for gigabytes of data.   [![star > 2000][Awesome]](https://github.com/allegro/bigcache)   [![godoc][GoDoc]](https://godoc.org/github.com/allegro/bigcache)
* [goleveldb](https://github.com/syndtr/goleveldb) **star:4102** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![star > 2000][Awesome]](https://github.com/syndtr/goleveldb)   [![godoc][GoDoc]](https://godoc.org/github.com/syndtr/goleveldb)
* [bbolt](https://github.com/etcd-io/bbolt) **star:4014** An embedded key/value database for Go.   [![star > 2000][Awesome]](https://github.com/etcd-io/bbolt)   [![There was an update last week][Green]](https://github.com/etcd-io/bbolt)   [![godoc][GoDoc]](https://godoc.org/github.com/etcd-io/bbolt)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:3667** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   [![star > 2000][Awesome]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![There was an update last week][Green]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
* [ledisdb](https://github.com/siddontang/ledisdb) **star:3513** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   [![star > 2000][Awesome]](https://github.com/siddontang/ledisdb)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/ledisdb)
* [buntdb](https://github.com/tidwall/buntdb) **star:3076** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![star > 2000][Awesome]](https://github.com/tidwall/buntdb)   [![There was an update last week][Green]](https://github.com/tidwall/buntdb)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/buntdb)
* [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2567** Your NoSQL database powered by Golang.   [![star > 2000][Awesome]](https://github.com/HouzuoGuo/tiedot)   [![godoc][GoDoc]](https://godoc.org/github.com/HouzuoGuo/tiedot)
* [immudb](https://github.com/codenotary/immudb) **star:1712** immudb is a lightweight, high-speed immutable database for systems and applications written in Go.   [![There was an update last week][Green]](https://github.com/codenotary/immudb)   [![godoc][GoDoc]](https://godoc.org/github.com/codenotary/immudb)
* [nutsdb](https://github.com/xujiajun/nutsdb) **star:1479** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/nutsdb)   [![Contains Chinese documents][CN]](https://github.com/xujiajun/nutsdb)
* [cache2go](https://github.com/muesli/cache2go) **star:1457** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/cache2go)
* [GCache](https://github.com/bluele/gcache) **star:1352** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![There was an update last week][Green]](https://github.com/bluele/gcache)   [![godoc][GoDoc]](https://godoc.org/github.com/bluele/gcache)
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:1133** CovenantSQL is a SQL database on blockchain.   [![godoc][GoDoc]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
* [diskv](https://github.com/peterbourgon/diskv) **star:998** Home-grown disk-backed key-value store.   [![There was an update last week][Green]](https://github.com/peterbourgon/diskv)   [![godoc][GoDoc]](https://godoc.org/github.com/peterbourgon/diskv)
* [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:979** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   [![godoc][GoDoc]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
* [moss](https://github.com/couchbase/moss) **star:788** Moss is a simple LSM key-value storage engine written in 100% Go.   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/moss)
* [pogreb](https://github.com/akrylysov/pogreb) **star:686** Embedded key-value store for read-heavy workloads.   [![godoc][GoDoc]](https://godoc.org/github.com/akrylysov/pogreb)
* [eliasdb](https://github.com/krotik/eliasdb) **star:604** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/eliasdb)
* [Bitcask](https://github.com/prologic/bitcask) **star:590** Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).   [![godoc][GoDoc]](https://godoc.org/github.com/prologic/bitcask)
* [levigo](https://github.com/jmhodges/levigo) **star:391** Levigo is a Go wrapper for LevelDB.   [![godoc][GoDoc]](https://godoc.org/github.com/jmhodges/levigo)
* [pudge](https://github.com/recoilme/pudge) **star:280** Fast and simple  key/value store written using Go's standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/pudge)
* [Vasto](https://github.com/chrislusf/vasto) **star:205** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/vasto)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/vasto)
* [Kivik](https://github.com/go-kivik/kivik) **star:201** Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.   [![godoc][GoDoc]](https://godoc.org/github.com/go-kivik/kivik)
* [piladb](https://github.com/fern4lvarez/piladb) **star:181** Lightweight RESTful database engine based on stack data structures.   [![godoc][GoDoc]](https://godoc.org/github.com/fern4lvarez/piladb)
* [Scribble](https://github.com/nanobox-io/golang-scribble) **star:114** Tiny flat file JSON store.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nanobox-io/golang-scribble)   [![godoc][GoDoc]](https://godoc.org/github.com/nanobox-io/golang-scribble)
* [Databunker](https://github.com/paranoidguy/databunker) **star:107** Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.   [![There was an update last week][Green]](https://github.com/paranoidguy/databunker)   [![godoc][GoDoc]](https://godoc.org/github.com/paranoidguy/databunker)
* [slowpoke](https://github.com/recoilme/slowpoke) **star:98** Key-value store with persistence.   [![It hasn't been updated in the last year][Yellow]](https://github.com/recoilme/slowpoke)   [![godoc][GoDoc]](https://godoc.org/github.com/recoilme/slowpoke)
* [cache](https://github.com/akyoto/cache) **star:76** In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.   [![godoc][GoDoc]](https://godoc.org/github.com/akyoto/cache)
* [bcache](https://github.com/iwanbk/bcache) **star:61** Eventually consistent distributed in-memory  cache Go library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/iwanbk/bcache)   [![godoc][GoDoc]](https://godoc.org/github.com/iwanbk/bcache)
* [couchcache](https://github.com/codingsince1985/couchcache) **star:51** RESTful caching micro-service backed by Couchbase server.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/couchcache)
* [unitdb](https://github.com/unit-io/unitdb) **star:50** Fast timeseries database for IoT, realtime messaging  applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application.   [![There was an update last week][Green]](https://github.com/unit-io/unitdb)   [![godoc][GoDoc]](https://godoc.org/github.com/unit-io/unitdb)
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:36** BigCache with clustering support and individual item expiration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][GoDoc]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
* [Coffer](https://github.com/claygod/coffer) **star:23** Simple ACID key-value database that supports transactions.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/coffer)
* [tempdb](https://github.com/rafaeljesus/tempdb) **star:14** Key-value store for temporary items.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/tempdb)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/tempdb)
* [ttlcache](https://github.com/cheshir/ttlcache) **star:1** In-memory key value storage with TTL for each record.   [![godoc][GoDoc]](https://godoc.org/github.com/cheshir/ttlcache)

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) **star:5687** Database migrations. CLI and Golang library.   [![star > 2000][Awesome]](https://github.com/golang-migrate/migrate)   [![There was an update last week][Green]](https://github.com/golang-migrate/migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/golang-migrate/migrate)
* [sql-migrate](https://github.com/rubenv/sql-migrate) **star:2016** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![star > 2000][Awesome]](https://github.com/rubenv/sql-migrate)   [![There was an update last week][Green]](https://github.com/rubenv/sql-migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/rubenv/sql-migrate)
* [goose](https://github.com/pressly/goose) **star:1538** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/goose)
* [skeema](https://github.com/skeema/skeema) **star:827** Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.   [![godoc][GoDoc]](https://godoc.org/github.com/skeema/skeema)
* [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:551** Database schema migration helper for Gorm ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gormigrate/gormigrate)
* [darwin](https://github.com/GuiaBolso/darwin) **star:110** Database schema evolution library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GuiaBolso/darwin)   [![godoc][GoDoc]](https://godoc.org/github.com/GuiaBolso/darwin)
* [migrator](https://github.com/lopezator/migrator) **star:106** Dead simple Go database migration library.   [![godoc][GoDoc]](https://godoc.org/github.com/lopezator/migrator)
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:69** A Go package to help write migrations with go-pg/pg.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:25** Django style fixtures for Golang's excellent built-in database/sql library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/RichardKnop/go-fixtures)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/go-fixtures)
* [pravasan](https://github.com/pravasan/pravasan) **star:24** Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pravasan/pravasan)
* [avro](https://github.com/khezen/avro) **star:21** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/avro)
* [schema](https://github.com/adlio/schema) **star:7** Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/schema)
* [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) **star:2** CLI-friendly package for go-pg migrations management.   [![There was an update last week][Green]](https://github.com/lawzava/go-pg-migrate)   [![godoc][GoDoc]](https://godoc.org/github.com/lawzava/go-pg-migrate)

*Database tools.*

* [vitess](https://github.com/youtube/vitess) **star:11306** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![star > 2000][Awesome]](https://github.com/youtube/vitess)   [![There was an update last week][Green]](https://github.com/youtube/vitess)   [![godoc][GoDoc]](https://godoc.org/github.com/youtube/vitess)
* [pgweb](https://github.com/sosedoff/pgweb) **star:6740** Web-based PostgreSQL database browser.   [![star > 2000][Awesome]](https://github.com/sosedoff/pgweb)   [![godoc][GoDoc]](https://godoc.org/github.com/sosedoff/pgweb)
* [kingshard](https://github.com/flike/kingshard) **star:5532** kingshard is a high performance proxy for MySQL powered by Golang.   [![star > 2000][Awesome]](https://github.com/flike/kingshard)   [![godoc][GoDoc]](https://godoc.org/github.com/flike/kingshard)   [![Contains Chinese documents][CN]](https://github.com/flike/kingshard)
* [orchestrator](https://github.com/github/orchestrator) **star:3920** MySQL replication topology manager & visualizer.   [![star > 2000][Awesome]](https://github.com/github/orchestrator)   [![There was an update last week][Green]](https://github.com/github/orchestrator)   [![godoc][GoDoc]](https://godoc.org/github.com/github/orchestrator)
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) **star:3324** Sync your MySQL data into Elasticsearch automatically.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql-elasticsearch)
* [go-mysql](https://github.com/siddontang/go-mysql) **star:2811** Go toolset to handle MySQL protocol and replication.   [![star > 2000][Awesome]](https://github.com/siddontang/go-mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-mysql)
* [pREST](https://github.com/prest/prest) **star:2544** Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new.   [![star > 2000][Awesome]](https://github.com/prest/prest)   [![There was an update last week][Green]](https://github.com/prest/prest)   [![godoc][GoDoc]](https://godoc.org/github.com/prest/prest)
* [chproxy](https://github.com/Vertamedia/chproxy) **star:543** HTTP proxy for ClickHouse database.   [![godoc][GoDoc]](https://godoc.org/github.com/Vertamedia/chproxy)
* [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) **star:356** Advanced scheduling for PostgreSQL.   [![There was an update last week][Green]](https://github.com/cybertec-postgresql/pg_timetable)   [![godoc][GoDoc]](https://godoc.org/github.com/cybertec-postgresql/pg_timetable)
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:249** Collects small insterts and sends big requests to ClickHouse servers.   [![There was an update last week][Green]](https://github.com/nikepan/clickhouse-bulk)   [![godoc][GoDoc]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
* [myreplication](https://github.com/2tvenom/myreplication) **star:171** MySql binary log replication listener. Supports statement and row based replication.   [![It hasn't been updated in the last year][Yellow]](https://github.com/2tvenom/myreplication)   [![godoc][GoDoc]](https://godoc.org/github.com/2tvenom/myreplication)
* [octillery](https://github.com/knocknote/octillery) **star:118** Go package for sharding databases ( Supports every ORM or raw SQL ).   [![godoc][GoDoc]](https://godoc.org/github.com/knocknote/octillery)
* [dbbench](https://github.com/sj14/dbbench) **star:51** Database benchmarking tool with support for several databases and scripts.   [![There was an update last week][Green]](https://github.com/sj14/dbbench)   [![godoc][GoDoc]](https://godoc.org/github.com/sj14/dbbench)
* [datagen](https://github.com/codingconcepts/datagen) **star:29** A fast data generator that's multi-table aware and supports multi-row DML.   [![godoc][GoDoc]](https://godoc.org/github.com/codingconcepts/datagen)
* [prep](https://github.com/hexdigest/prep) **star:27** Use prepared SQL statements without changing your code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/prep)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/prep)
* [rwdb](https://github.com/andizzle/rwdb) **star:11** rwdb provides read replica capability for multiple database servers setup.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andizzle/rwdb)   [![godoc][GoDoc]](https://godoc.org/github.com/andizzle/rwdb)

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) **star:3612** Go library that helps you build SQL queries.   [![star > 2000][Awesome]](https://github.com/Masterminds/squirrel)   [![There was an update last week][Green]](https://github.com/Masterminds/squirrel)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/squirrel)
* [xo](https://github.com/knq/xo) **star:2690** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   [![star > 2000][Awesome]](https://github.com/knq/xo)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/xo)
* [gendry](https://github.com/didi/gendry) **star:1148** Non-invasive SQL builder and powerful data binder.   [![godoc][GoDoc]](https://godoc.org/github.com/didi/gendry)   [![Contains Chinese documents][CN]](https://github.com/didi/gendry)
* [goqu](https://github.com/doug-martin/goqu) **star:946** Idiomatic SQL builder and query library.   [![godoc][GoDoc]](https://godoc.org/github.com/doug-martin/goqu)
* [Dotsql](https://github.com/gchaincl/dotsql) **star:570** Go library that helps you keep sql files in one place and use them with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/dotsql)
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:515** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
* [jet](https://github.com/go-jet/jet) **star:350** Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.   [![There was an update last week][Green]](https://github.com/go-jet/jet)   [![godoc][GoDoc]](https://godoc.org/github.com/go-jet/jet)
* [dbq](https://github.com/rocketlaunchr/dbq) **star:259** Zero boilerplate database operations for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dbq)
* [sqrl](https://github.com/elgris/sqrl) **star:216** SQL query builder, fork of Squirrel with improved performance.   [![godoc][GoDoc]](https://godoc.org/github.com/elgris/sqrl)
* [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
* [sqlingo](https://github.com/lqs/sqlingo) **star:136** A lightweight DSL to build SQL in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/lqs/sqlingo)
* [sq](https://github.com/bokwoon95/go-structured-query) **star:95** Type-safe SQL builder and struct mapper for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bokwoon95/go-structured-query)
* [igor](https://github.com/galeone/igor) **star:83** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/igor)
* [godbal](https://github.com/xujiajun/godbal) **star:51** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xujiajun/godbal)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/godbal)
* [buildsqlx](https://github.com/arthurkushman/buildsqlx) **star:26** Go database query builder library for PostgreSQL.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/buildsqlx)
* [qry](https://github.com/HnH/qry) **star:19** Tool that generates constants from files with raw SQL queries.   [![godoc][GoDoc]](https://godoc.org/github.com/HnH/qry)
* [sqlf](https://github.com/leporo/sqlf) **star:16** Fast SQL query builder.   [![godoc][GoDoc]](https://godoc.org/github.com/leporo/sqlf)
* [gosql](https://github.com/twharmon/gosql) **star:15** SQL Query builder with better null values support.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/gosql)
* [hasql](https://golang.yandex/hasql)  Library for accessing multi-host SQL database installations.
* [mpath](https://github.com/spacetab-io/mpath-go) **star:7** MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/spacetab-io/mpath-go)   [![godoc][GoDoc]](https://godoc.org/github.com/spacetab-io/mpath-go)
* [ormlite](https://github.com/pupizoid/ormlite)  Lightweight package containing some ORM-like features and helpers for sqlite databases.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:10434** MySQL driver for Go.   [![star > 2000][Awesome]](https://github.com/go-sql-driver/mysql)   [![There was an update last week][Green]](https://github.com/go-sql-driver/mysql)   [![godoc][GoDoc]](https://godoc.org/github.com/go-sql-driver/mysql)
    * [pq](https://github.com/lib/pq) **star:6328** Pure Go Postgres driver for database/sql.   [![star > 2000][Awesome]](https://github.com/lib/pq)   [![godoc][GoDoc]](https://godoc.org/github.com/lib/pq)
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:4576** SQLite3 driver for go that uses database/sql.   [![star > 2000][Awesome]](https://github.com/mattn/go-sqlite3)   [![There was an update last week][Green]](https://github.com/mattn/go-sqlite3)
    * [pgx](https://github.com/jackc/pgx) **star:3642** PostgreSQL driver supporting features beyond those exposed by database/sql.   [![star > 2000][Awesome]](https://github.com/jackc/pgx)   [![There was an update last week][Green]](https://github.com/jackc/pgx)   [![godoc][GoDoc]](https://godoc.org/github.com/jackc/pgx)
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1333** Microsoft MSSQL driver for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/denisenkom/go-mssqldb)
    * [go-oci8](https://github.com/mattn/go-oci8) **star:529** Oracle driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-oci8)
    * [godror](https://github.com/godror/godror) **star:217** Oracle driver for Go, using the ODPI-C driver.   [![There was an update last week][Green]](https://github.com/godror/godror)
    * [firebirdsql](https://github.com/nakagami/firebirdsql) **star:137** Firebird RDBMS SQL driver for Go.   [![There was an update last week][Green]](https://github.com/nakagami/firebirdsql)   [![godoc][GoDoc]](https://godoc.org/github.com/nakagami/firebirdsql)
    * [go-adodb](https://github.com/mattn/go-adodb) **star:116** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-adodb)
    * [gofreetds](https://github.com/minus5/gofreetds) **star:105** Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).   [![godoc][GoDoc]](https://godoc.org/github.com/minus5/gofreetds)
    * [avatica](https://github.com/apache/calcite-avatica-go) **star:61** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][GoDoc]](https://godoc.org/github.com/apache/calcite-avatica-go)
    * [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) **star:30** SQLite with pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cvilsmeier/sqinn-go)
    * [bgc](https://github.com/viant/bgc) **star:15** Datastore Connectivity for BigQuery for go.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/bgc)

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) **star:10618** Redis client for Golang.   [![star > 2000][Awesome]](https://github.com/go-redis/redis)   [![There was an update last week][Green]](https://github.com/go-redis/redis)   [![godoc][GoDoc]](https://godoc.org/github.com/go-redis/redis)
    * [redigo](https://github.com/gomodule/redigo) **star:8140** Redigo is a Go client for the Redis database.   [![star > 2000][Awesome]](https://github.com/gomodule/redigo)   [![godoc][GoDoc]](https://godoc.org/github.com/gomodule/redigo)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:5265** Official MongoDB driver for the Go language.   [![star > 2000][Awesome]](https://github.com/mongodb/mongo-go-driver)   [![There was an update last week][Green]](https://github.com/mongodb/mongo-go-driver)   [![godoc][GoDoc]](https://godoc.org/github.com/mongodb/mongo-go-driver)
    * [mgo](https://github.com/globalsign/mgo) **star:1867** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   [![godoc][GoDoc]](https://godoc.org/github.com/globalsign/mgo)
    * [gorethink](https://github.com/dancannon/gorethink) **star:1551** Go language driver for RethinkDB.   [![godoc][GoDoc]](https://godoc.org/github.com/dancannon/gorethink)
    * [qmgo](https://github.com/qiniu/qmgo) **star:391** The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo.   [![There was an update last week][Green]](https://github.com/qiniu/qmgo)   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/qmgo)   [![Contains Chinese documents][CN]](https://github.com/qiniu/qmgo)
    * [redeo](https://github.com/bsm/redeo) **star:390** Redis-protocol compatible TCP servers/services.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redeo)
    * [neoism](https://github.com/jmcvetta/neoism) **star:373** Neo4j client for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jmcvetta/neoism)
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:351** Aerospike client in Go language.   [![There was an update last week][Green]](https://github.com/aerospike/aerospike-client-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aerospike/aerospike-client-go)
    * [gocb](https://github.com/couchbase/gocb) **star:321** Official Couchbase Go SDK.   [![There was an update last week][Green]](https://github.com/couchbase/gocb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/gocb)
    * [go-couchbase](https://github.com/couchbase/go-couchbase) **star:308** Couchbase client in Go.   [![There was an update last week][Green]](https://github.com/couchbase/go-couchbase)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/go-couchbase)
    * [mgm](https://github.com/kamva/mgm) **star:234** MongoDB model-based ODM for Go (based on official MongoDB driver).   [![godoc][GoDoc]](https://godoc.org/github.com/kamva/mgm)
    * [go-rejson](https://github.com/nitishm/go-rejson) **star:152** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/nitishm/go-rejson)
    * [godis](https://github.com/piaohao/godis) **star:86** redis client implement by golang, inspired by jedis.   [![godoc][GoDoc]](https://godoc.org/github.com/piaohao/godis)
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) **star:76** Neo4j REST Client in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davemeehan/Neo4j-GO)   [![godoc][GoDoc]](https://godoc.org/github.com/davemeehan/Neo4j-GO)
    * [arangolite](https://github.com/solher/arangolite) **star:69** Lightweight golang driver for ArangoDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/solher/arangolite)   [![godoc][GoDoc]](https://godoc.org/github.com/solher/arangolite)
    * [go-pilosa](https://github.com/pilosa/go-pilosa) **star:43** Go client library for Pilosa.   [![godoc][GoDoc]](https://godoc.org/github.com/pilosa/go-pilosa)
    * [forestdb](https://github.com/couchbase/goforestdb) **star:29** Go bindings for ForestDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/couchbase/goforestdb)   [![godoc][GoDoc]](https://godoc.org/github.com/couchbase/goforestdb)
    * [goriak](https://github.com/zegl/goriak) **star:26** Go language driver for Riak KV.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zegl/goriak)   [![godoc][GoDoc]](https://godoc.org/github.com/zegl/goriak)
    * [neo4j](https://github.com/cihangir/neo4j) **star:26** Neo4j Rest API Bindings for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cihangir/neo4j)   [![godoc][GoDoc]](https://godoc.org/github.com/cihangir/neo4j)
    * [xredis](https://github.com/shomali11/xredis) **star:13** Typesafe, customizable, clean & easy to use Redis client.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shomali11/xredis)   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/xredis)
    * [godscache](https://github.com/defcronyke/godscache) **star:7** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.   [![It hasn't been updated in the last year][Yellow]](https://github.com/defcronyke/godscache)   [![godoc][GoDoc]](https://godoc.org/github.com/defcronyke/godscache)
    * [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.
    * [asc](https://github.com/viant/asc) **star:6** Datastore Connectivity for Aerospike for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/viant/asc)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/asc)
    * [gocosmos](https://github.com/btnguyen2k/gocosmos) **star:2** REST client and standard `database/sql` driver for Azure Cosmos DB.   [![godoc][GoDoc]](https://godoc.org/github.com/btnguyen2k/gocosmos)
    * [gocql](http://gocql.github.io)  Go language driver for Apache Cassandra.

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) **star:7347** Modern text indexing library for go.   [![star > 2000][Awesome]](https://github.com/blevesearch/bleve)   [![There was an update last week][Green]](https://github.com/blevesearch/bleve)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/bleve)
    * [elastic](https://github.com/olivere/elastic) **star:5672** Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/olivere/elastic)   [![godoc][GoDoc]](https://godoc.org/github.com/olivere/elastic)
    * [riot](https://github.com/go-ego/riot) **star:5589** Go Open Source, Distributed, Simple and efficient Search Engine.   [![star > 2000][Awesome]](https://github.com/go-ego/riot)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/riot)   [![Contains Chinese documents][CN]](https://github.com/go-ego/riot)
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:3077** Official Elasticsearch client for Go.   [![star > 2000][Awesome]](https://github.com/elastic/go-elasticsearch)   [![godoc][GoDoc]](https://godoc.org/github.com/elastic/go-elasticsearch)
    * [elastigo](https://github.com/mattbaird/elastigo) **star:946** Elasticsearch client library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattbaird/elastigo)   [![godoc][GoDoc]](https://godoc.org/github.com/mattbaird/elastigo)
    * [elasticsql](https://github.com/cch123/elasticsql) **star:661** Convert sql to elasticsearch dsl in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cch123/elasticsql)
    * [skizze](https://github.com/seiflotfy/skizze) **star:75** probabilistic data-structures service and storage.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seiflotfy/skizze)   [![godoc][GoDoc]](https://godoc.org/github.com/seiflotfy/skizze)
    * [goes](https://github.com/OwnLocal/goes) **star:24** Library to interact with Elasticsearch.   [![godoc][GoDoc]](https://godoc.org/github.com/OwnLocal/goes)

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) **star:13727** Graph database with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/google/cayley)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cayley)
    * [gokv](https://github.com/philippgille/gokv) **star:302** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/gokv)
    * [cachego](https://github.com/fabiorphp/cachego) **star:144** Golang Cache component for multiple drivers.   [![godoc][GoDoc]](https://godoc.org/github.com/fabiorphp/cachego)
    * [dsc](https://github.com/viant/dsc) **star:21** Datastore connectivity for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsc)

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) **star:2893** Now is a time toolkit for golang.   [![star > 2000][Awesome]](https://github.com/jinzhu/now)   [![godoc][GoDoc]](https://godoc.org/github.com/jinzhu/now)
* [dateparse](https://github.com/araddon/dateparse) **star:1267** Parse date's without knowing format in advance.   [![godoc][GoDoc]](https://godoc.org/github.com/araddon/dateparse)
* [carbon](https://github.com/uniplaces/carbon) **star:534** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   [![godoc][GoDoc]](https://godoc.org/github.com/uniplaces/carbon)
* [durafmt](https://github.com/hako/durafmt) **star:339** Time duration formatting library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hako/durafmt)
* [timeutil](https://github.com/leekchan/timeutil) **star:183** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/leekchan/timeutil)   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/timeutil)
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:90** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).   [![godoc][GoDoc]](https://godoc.org/github.com/yaa110/go-persian-calendar)
* [iso8601](https://github.com/relvacode/iso8601) **star:84** Efficiently parse ISO8601 date-times without regex.   [![godoc][GoDoc]](https://godoc.org/github.com/relvacode/iso8601)
* [timespan](https://github.com/SaidinWoT/timespan) **star:74** For interacting with intervals of time, defined as a start time and a duration.   [![It hasn't been updated in the last year][Yellow]](https://github.com/SaidinWoT/timespan)   [![godoc][GoDoc]](https://godoc.org/github.com/SaidinWoT/timespan)
* [date](https://github.com/rickb777/date) **star:58** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.   [![godoc][GoDoc]](https://godoc.org/github.com/rickb777/date)
* [feiertage](https://github.com/wlbr/feiertage) **star:35** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/feiertage)
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:31** Calculate the sunrise and sunset times for a given location.   [![godoc][GoDoc]](https://godoc.org/github.com/nathan-osman/go-sunrise)
* [go-str2duration](https://github.com/xhit/go-str2duration) **star:18** Convert string to duration. Support time.Duration returned string and more.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-str2duration)
* [kair](https://github.com/GuilhermeCaruso/kair) **star:18** Date and Time - Golang Formatting Library.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/kair)
* [NullTime](https://github.com/kirillDanshin/nulltime) **star:11** Nullable `time.Time`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/nulltime)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/nulltime)
* [tuesday](https://github.com/osteele/tuesday) **star:9** Ruby-compatible Strftime function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/osteele/tuesday)   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/tuesday)
* [cronrange](https://github.com/1set/cronrange) **star:8** Parses Cron-style time range expressions, checks if the given time is within any ranges.   [![It hasn't been updated in the last year][Yellow]](https://github.com/1set/cronrange)   [![godoc][GoDoc]](https://godoc.org/github.com/1set/cronrange)
* [strftime](https://github.com/awoodbeck/strftime) **star:7** C99-compatible strftime formatter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awoodbeck/strftime)   [![godoc][GoDoc]](https://godoc.org/github.com/awoodbeck/strftime)
* [go-week](https://github.com/stoewer/go-week) **star:4** An efficient package to work with ISO8601 week dates.   [![godoc][GoDoc]](https://godoc.org/github.com/stoewer/go-week)

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) **star:19215** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![star > 2000][Awesome]](https://github.com/go-kit/kit)   [![There was an update last week][Green]](https://github.com/go-kit/kit)   [![godoc][GoDoc]](https://godoc.org/github.com/go-kit/kit)
* [go-micro](https://github.com/micro/go-micro) **star:15279** A distributed systems development framework.   [![star > 2000][Awesome]](https://github.com/micro/go-micro)   [![There was an update last week][Green]](https://github.com/micro/go-micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/go-micro)
* [grpc-go](https://github.com/grpc/grpc-go) **star:13006** The Go language implementation of gRPC. HTTP/2 based RPC.   [![star > 2000][Awesome]](https://github.com/grpc/grpc-go)   [![There was an update last week][Green]](https://github.com/grpc/grpc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/grpc/grpc-go)
* [micro](https://github.com/micro/micro) **star:9618** A distributed systems runtime for the cloud and beyond.   [![star > 2000][Awesome]](https://github.com/micro/micro)   [![There was an update last week][Green]](https://github.com/micro/micro)   [![godoc][GoDoc]](https://godoc.org/github.com/micro/micro)
* [NATS](https://github.com/nats-io/gnatsd) **star:8788** Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.   [![star > 2000][Awesome]](https://github.com/nats-io/gnatsd)   [![There was an update last week][Green]](https://github.com/nats-io/gnatsd)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/gnatsd)
* [rpcx](https://github.com/smallnest/rpcx) **star:5392** Distributed pluggable RPC service framework like alibaba Dubbo.   [![star > 2000][Awesome]](https://github.com/smallnest/rpcx)   [![There was an update last week][Green]](https://github.com/smallnest/rpcx)   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/rpcx)
* [raft](https://github.com/hashicorp/raft) **star:4261** Golang implementation of the Raft consensus protocol, by HashiCorp.   [![star > 2000][Awesome]](https://github.com/hashicorp/raft)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/raft)
* [raft](https://github.com/coreos/etcd/tree/master/raft)  Go implementation of the Raft consensus protocol, by CoreOS.
* [tendermint](https://github.com/tendermint/tendermint) **star:3874** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![star > 2000][Awesome]](https://github.com/tendermint/tendermint)   [![There was an update last week][Green]](https://github.com/tendermint/tendermint)   [![godoc][GoDoc]](https://godoc.org/github.com/tendermint/tendermint)
* [torrent](https://github.com/anacrolix/torrent) **star:3767** BitTorrent client package.   [![star > 2000][Awesome]](https://github.com/anacrolix/torrent)   [![There was an update last week][Green]](https://github.com/anacrolix/torrent)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/torrent)
* [KrakenD](https://github.com/devopsfaith/krakend) **star:3486** Ultra performant API Gateway framework with middlewares.   [![star > 2000][Awesome]](https://github.com/devopsfaith/krakend)   [![There was an update last week][Green]](https://github.com/devopsfaith/krakend)   [![godoc][GoDoc]](https://godoc.org/github.com/devopsfaith/krakend)
* [dragonboat](https://github.com/lni/dragonboat) **star:3454** A feature complete and high performance multi-group Raft library in Go.   [![star > 2000][Awesome]](https://github.com/lni/dragonboat)   [![There was an update last week][Green]](https://github.com/lni/dragonboat)   [![godoc][GoDoc]](https://godoc.org/github.com/lni/dragonboat)   [![Contains Chinese documents][CN]](https://github.com/lni/dragonboat)
* [glow](https://github.com/chrislusf/glow) **star:2915** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![star > 2000][Awesome]](https://github.com/chrislusf/glow)   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrislusf/glow)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/glow)
* [gleam](https://github.com/chrislusf/gleam) **star:2731** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![star > 2000][Awesome]](https://github.com/chrislusf/gleam)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/gleam)
* [emitter-io](https://github.com/emitter-io/emitter) **star:2707** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![star > 2000][Awesome]](https://github.com/emitter-io/emitter)   [![godoc][GoDoc]](https://godoc.org/github.com/emitter-io/emitter)
* [liftbridge](https://github.com/liftbridge-io/liftbridge) **star:1980** Lightweight, fault-tolerant message streams for NATS.   [![godoc][GoDoc]](https://godoc.org/github.com/liftbridge-io/liftbridge)
* [hprose](https://github.com/hprose/hprose-golang) **star:1130** Very newbility RPC Library, support 25+ languages now.   [![There was an update last week][Green]](https://github.com/hprose/hprose-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/hprose/hprose-golang)   [![Contains Chinese documents][CN]](https://github.com/hprose/hprose-golang)
* [ringpop-go](https://github.com/uber/ringpop-go) **star:656** Scalable, fault-tolerant application-layer sharding for Go applications.   [![There was an update last week][Green]](https://github.com/uber/ringpop-go)   [![godoc][GoDoc]](https://godoc.org/github.com/uber/ringpop-go)
* [gorpc](https://github.com/valyala/gorpc) **star:627** Simple, fast and scalable RPC library for high load.   [![It hasn't been updated in the last year][Yellow]](https://github.com/valyala/gorpc)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/gorpc)
* [go-health](https://github.com/InVisionApp/go-health) **star:580** Library for enabling asynchronous dependency health checks in your service.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/go-health)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/go-health)
* [rain](https://github.com/cenkalti/rain) **star:559** BitTorrent client and library.   [![godoc][GoDoc]](https://godoc.org/github.com/cenkalti/rain)
* [digota](https://github.com/digota/digota) **star:376** grpc ecommerce microservice.   [![It hasn't been updated in the last year][Yellow]](https://github.com/digota/digota)   [![godoc][GoDoc]](https://godoc.org/github.com/digota/digota)
* [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
* [consistent](https://github.com/buraksezer/consistent) **star:355** Consistent hashing with bounded loads.   [![It hasn't been updated in the last year][Yellow]](https://github.com/buraksezer/consistent)   [![godoc][GoDoc]](https://godoc.org/github.com/buraksezer/consistent)
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) **star:344** A library built to provide support for defining async service health checks for golang services.   [![There was an update last week][Green]](https://github.com/AppsFlyer/go-sundheit)   [![godoc][GoDoc]](https://godoc.org/github.com/AppsFlyer/go-sundheit)
* [sleuth](https://github.com/ursiform/sleuth) **star:330** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ursiform/sleuth)   [![godoc][GoDoc]](https://godoc.org/github.com/ursiform/sleuth)
* [redis-lock](https://github.com/bsm/redislock) **star:317** Simplified distributed locking implementation using Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/bsm/redislock)
* [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [go-jump](https://github.com/dgryski/go-jump) **star:308** Port of Google's "Jump" Consistent Hash function.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dgryski/go-jump)   [![godoc][GoDoc]](https://godoc.org/github.com/dgryski/go-jump)
* [dht](https://github.com/anacrolix/dht) **star:181** BitTorrent Kademlia DHT implementation.   [![There was an update last week][Green]](https://github.com/anacrolix/dht)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/dht)
* [jsonrpc](https://github.com/ybbus/jsonrpc) **star:160** JSON-RPC 2.0 HTTP client implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/jsonrpc)
* [arpc](https://github.com/lesismal/arpc) **star:147** More effective network communication, support two-way-calling, notify, broadcast.   [![godoc][GoDoc]](https://godoc.org/github.com/lesismal/arpc)
* [jsonrpc](https://github.com/osamingo/jsonrpc) **star:139** The jsonrpc package helps implement of JSON-RPC 2.0.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/jsonrpc)
* [celeriac](https://github.com/svcavallar/celeriac.v1) **star:62** Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/svcavallar/celeriac.v1)
* [doublejump](https://github.com/edwingeng/doublejump) **star:57** A revamped Google's jump consistent hash.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/doublejump)
* [Semaphore](https://github.com/jexia/semaphore) **star:51** A straightforward (micro) service orchestrator.   [![godoc][GoDoc]](https://godoc.org/github.com/jexia/semaphore)
* [outboxer](https://github.com/italolelis/outboxer) **star:47** Outboxer is a go library that implements the outbox pattern.   [![There was an update last week][Green]](https://github.com/italolelis/outboxer)   [![godoc][GoDoc]](https://godoc.org/github.com/italolelis/outboxer)
* [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
* [pjrpc](https://gitlab.com/pjrpc/pjrpc)  Golang JSON-RPC Server-Client with Protobuf spec.
* [flowgraph](https://github.com/vectaport/flowgraph) **star:38** flow-based programming package.   [![godoc][GoDoc]](https://godoc.org/github.com/vectaport/flowgraph)
* [drmaa](https://github.com/dgruber/drmaa) **star:29** Job submission library for cluster schedulers based on the DRMAA standard.   [![godoc][GoDoc]](https://godoc.org/github.com/dgruber/drmaa)
* [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
* [go-pdu](https://github.com/pdupub/go-pdu) **star:24** A decentralized identity-based social network.   [![godoc][GoDoc]](https://godoc.org/github.com/pdupub/go-pdu)
* [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) **star:17** MySQL based distributed lock.   [![godoc][GoDoc]](https://godoc.org/github.com/sanketplus/go-mysql-lock)
* [dynatomic](https://github.com/tylfin/dynatomic) **star:14** A library for using DynamoDB as an atomic counter.   [![godoc][GoDoc]](https://godoc.org/github.com/tylfin/dynatomic)
* [gmsec](https://github.com/gmsec/micro) **star:9** A Go distributed systems development framework.   [![godoc][GoDoc]](https://godoc.org/github.com/gmsec/micro)
* [consistenthash](https://github.com/mbrostami/consistenthash) **star:5** Consistent hashing with configurable replicas.   [![godoc][GoDoc]](https://godoc.org/github.com/mbrostami/consistenthash)

## Dynamic DNS

*Tools for updating dynamic DNS records.*

* [GoDNS](https://github.com/timothyye/godns) **star:744** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![There was an update last week][Green]](https://github.com/timothyye/godns)   [![godoc][GoDoc]](https://godoc.org/github.com/timothyye/godns)
* [DDNS](https://github.com/skibish/ddns) **star:172** Personal DDNS client with Digital Ocean Networking DNS as backend.   [![There was an update last week][Green]](https://github.com/skibish/ddns)   [![godoc][GoDoc]](https://godoc.org/github.com/skibish/ddns)
* [dyndns](https://gitlab.com/alcastle/dyndns)  Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) **star:7767** Email and SMTP testing with web and API interface.   [![star > 2000][Awesome]](https://github.com/mailhog/MailHog)   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/MailHog)
* [hermes](https://github.com/matcornic/hermes) **star:2163** Golang package that generates clean, responsive HTML e-mails.   [![star > 2000][Awesome]](https://github.com/matcornic/hermes)   [![godoc][GoDoc]](https://godoc.org/github.com/matcornic/hermes)
* [email](https://github.com/jordan-wright/email) **star:1630** A robust and flexible email library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/jordan-wright/email)
* [go-imap](https://github.com/emersion/go-imap) **star:1186** IMAP library for clients and servers.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-imap)
* [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
* [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:699** SendGrid's Go library for sending email.   [![There was an update last week][Green]](https://github.com/sendgrid/sendgrid-go)   [![godoc][GoDoc]](https://godoc.org/github.com/sendgrid/sendgrid-go)
* [mailgun-go](https://github.com/mailgun/mailgun-go) **star:500** Go library for sending mail with the Mailgun API.   [![There was an update last week][Green]](https://github.com/mailgun/mailgun-go)   [![godoc][GoDoc]](https://godoc.org/github.com/mailgun/mailgun-go)
* [Hectane](https://github.com/hectane/hectane) **star:195** Lightweight SMTP client providing an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/hectane/hectane)
* [go-message](https://github.com/emersion/go-message) **star:189** Streaming library for the Internet Message Format and mail messages.   [![There was an update last week][Green]](https://github.com/emersion/go-message)   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-message)
* [douceur](https://github.com/aymerick/douceur) **star:187** CSS inliner for your HTML emails.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aymerick/douceur)   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/douceur)
* [email-verifier](https://github.com/AfterShip/email-verifier) **star:147** A Go library for email verification without sending any emails.   [![godoc][GoDoc]](https://godoc.org/github.com/AfterShip/email-verifier)
* [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:89** Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.   [![godoc][GoDoc]](https://godoc.org/github.com/xhit/go-simple-mail)
* [mailchain](https://github.com/mailchain/mailchain) **star:68** Send encrypted emails to blockchain addresses written in Go.   [![There was an update last week][Green]](https://github.com/mailchain/mailchain)   [![godoc][GoDoc]](https://godoc.org/github.com/mailchain/mailchain)
* [go-premailer](https://github.com/vanng822/go-premailer) **star:67** Inline styling for HTML mail in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/vanng822/go-premailer)
* [go-dkim](https://github.com/toorop/go-dkim) **star:64** DKIM library, to sign & verify email.   [![godoc][GoDoc]](https://godoc.org/github.com/toorop/go-dkim)
* [smtp](https://github.com/mailhog/smtp) **star:59** SMTP server protocol state machine.   [![godoc][GoDoc]](https://godoc.org/github.com/mailhog/smtp)
* [go-email-validator](https://github.com/go-email-validator/go-email-validator) **star:7** Modular email validator for syntax, disposable, smtp, etc... checking.   [![There was an update last week][Green]](https://github.com/go-email-validator/go-email-validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-email-validator/go-email-validator)

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [gopher-lua](https://github.com/yuin/gopher-lua) **star:3903** Lua 5.1 VM and compiler written in Go.   [![star > 2000][Awesome]](https://github.com/yuin/gopher-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/yuin/gopher-lua)
* [tengo](https://github.com/d5/tengo) **star:2145** Bytecode compiled script language for Go.   [![star > 2000][Awesome]](https://github.com/d5/tengo)   [![godoc][GoDoc]](https://godoc.org/github.com/d5/tengo)
* [go-lua](https://github.com/Shopify/go-lua) **star:2010** Port of the Lua 5.2 VM to pure Go.   [![star > 2000][Awesome]](https://github.com/Shopify/go-lua)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/go-lua)
* [goja](https://github.com/dop251/goja) **star:1938** ECMAScript 5.1(+) implementation in Go.   [![There was an update last week][Green]](https://github.com/dop251/goja)   [![godoc][GoDoc]](https://godoc.org/github.com/dop251/goja)
* [expr](https://github.com/antonmedv/expr) **star:1597** Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing.   [![godoc][GoDoc]](https://godoc.org/github.com/antonmedv/expr)
* [go-python](https://github.com/sbinet/go-python) **star:1225** naive go bindings to the CPython C-API.   [![godoc][GoDoc]](https://godoc.org/github.com/sbinet/go-python)
* [anko](https://github.com/mattn/anko) **star:1101** Scriptable interpreter written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/anko)
* [go-php](https://github.com/deuill/go-php) **star:777** PHP bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/deuill/go-php)   [![godoc][GoDoc]](https://godoc.org/github.com/deuill/go-php)
* [go-duktape](https://github.com/olebedev/go-duktape) **star:741** Duktape JavaScript engine bindings for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/go-duktape)
* [cel-go](https://github.com/google/cel-go) **star:679** Fast, portable, non-Turing complete expression evaluation with gradual typing.   [![There was an update last week][Green]](https://github.com/google/cel-go)   [![godoc][GoDoc]](https://godoc.org/github.com/google/cel-go)
* [golua](https://github.com/aarzilli/golua) **star:507** Go bindings for Lua C API.
* [gisp](https://github.com/jcla1/gisp) **star:446** Simple LISP in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jcla1/gisp)   [![godoc][GoDoc]](https://godoc.org/github.com/jcla1/gisp)
* [gval](https://github.com/PaesslerAG/gval) **star:282** A highly customizable expression language written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/PaesslerAG/gval)
* [gentee](https://github.com/gentee/gentee) **star:69** Embeddable scripting programming language.   [![godoc][GoDoc]](https://godoc.org/github.com/gentee/gentee)
* [binder](https://github.com/alexeyco/binder) **star:48** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexeyco/binder)   [![godoc][GoDoc]](https://godoc.org/github.com/alexeyco/binder)
* [purl](https://github.com/ian-kent/purl) **star:30** Perl 5.18.2 embedded in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/purl)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/purl)
* [ngaro](https://github.com/db47h/ngaro) **star:20** Embeddable Ngaro VM implementation enabling scripting in Retro.   [![It hasn't been updated in the last year][Yellow]](https://github.com/db47h/ngaro)   [![godoc][GoDoc]](https://godoc.org/github.com/db47h/ngaro)
* [ecal](https://github.com/krotik/ecal) **star:5** A simple embeddable scripting language which supports concurrent event processing.   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/ecal)

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) **star:6484** Package that provides simple error handling primitives.   [![star > 2000][Awesome]](https://github.com/pkg/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/errors)
* [go-multierror](https://github.com/hashicorp/go-multierror) **star:1113** Go (golang) package for representing a list of errors as a single error.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-multierror)
* [eris](https://github.com/rotisserie/eris) **star:729** A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.   [![There was an update last week][Green]](https://github.com/rotisserie/eris)   [![godoc][GoDoc]](https://godoc.org/github.com/rotisserie/eris)
* [errorx](https://github.com/joomcode/errorx) **star:663** A feature rich error package with stack traces, composition of errors and more.   [![godoc][GoDoc]](https://godoc.org/github.com/joomcode/errorx)
* [tracerr](https://github.com/ztrue/tracerr) **star:639** Golang errors with stack trace and source fragments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ztrue/tracerr)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/tracerr)
* [errlog](https://github.com/snwfdhmp/errlog) **star:391** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   [![godoc][GoDoc]](https://godoc.org/github.com/snwfdhmp/errlog)
* [emperror](https://github.com/emperror/emperror) **star:196** Error handling tools and best practices for Go libraries and applications.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/emperror)
* [errors](https://github.com/emperror/errors) **star:87** Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.   [![godoc][GoDoc]](https://godoc.org/github.com/emperror/errors)
* [errors](https://github.com/bnkamalesh/errors) **star:15** Drop-in replacement for builting Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/errors)
* [Falcon](https://github.com/SonicRoshan/falcon) **star:6** A Simple Yet Highly Powerful Package For Error Handling.   [![It hasn't been updated in the last year][Yellow]](https://github.com/SonicRoshan/falcon)   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/falcon)
* [errors](https://github.com/neuronlabs/errors) **star:3** Simple golang error handling with classification primitives.   [![It hasn't been updated in the last year][Yellow]](https://github.com/neuronlabs/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/neuronlabs/errors)
* [errors](https://github.com/PumpkinSeed/errors) **star:2** The most simple error wrapper with awesome performance and minimal memory overhead.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PumpkinSeed/errors)   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/errors)

## File Handling

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) **star:3475** FileSystem Abstraction System for Go.   [![star > 2000][Awesome]](https://github.com/spf13/afero)   [![godoc][GoDoc]](https://godoc.org/github.com/spf13/afero)
* [pdfcpu](https://github.com/pdfcpu/pdfcpu) **star:1884** PDF processor.   [![There was an update last week][Green]](https://github.com/pdfcpu/pdfcpu)   [![godoc][GoDoc]](https://godoc.org/github.com/pdfcpu/pdfcpu)
* [notify](https://github.com/rjeczalik/notify) **star:608** File system event notification library with simple API, similar to os/signal.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/notify)
* [copy](https://github.com/otiai10/copy) **star:235** Copy directory recursively.   [![There was an update last week][Green]](https://github.com/otiai10/copy)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/copy)
* [bigfile](https://github.com/bigfile/bigfile) **star:179** A file transfer system, support to manage files with http api, rpc call and ftp client.   [![godoc][GoDoc]](https://godoc.org/github.com/bigfile/bigfile)   [![Contains Chinese documents][CN]](https://github.com/bigfile/bigfile)
* [afs](https://github.com/viant/afs) **star:114** Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.   [![There was an update last week][Green]](https://github.com/viant/afs)   [![godoc][GoDoc]](https://godoc.org/github.com/viant/afs)
* [vfs](https://github.com/C2FO/vfs) **star:84** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   [![godoc][GoDoc]](https://godoc.org/github.com/C2FO/vfs)
* [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:78** Load csv file using tag.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-csv-tag)
* [opc](https://github.com/qmuntal/opc) **star:65** Load Open Packaging Conventions (OPC) files for Go.   [![There was an update last week][Green]](https://github.com/qmuntal/opc)   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/opc)
* [skywalker](https://github.com/dixonwille/skywalker) **star:61** Package to allow one to concurrently go through a filesystem with ease.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dixonwille/skywalker)   [![godoc][GoDoc]](https://godoc.org/github.com/dixonwille/skywalker)
* [stl](https://gitlab.com/russoj88/stl)  Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.
* [go-exiftool](https://github.com/barasher/go-exiftool) **star:45** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/barasher/go-exiftool)
* [tarfs](https://github.com/posener/tarfs) **star:45** Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/tarfs)
* [checksum](https://github.com/codingsince1985/checksum) **star:31** Compute message digest, like MD5 and SHA256, for large files.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/checksum)
* [go-gtfs](https://github.com/artonge/go-gtfs) **star:24** Load gtfs files in go.   [![godoc][GoDoc]](https://godoc.org/github.com/artonge/go-gtfs)
* [flop](https://github.com/homedepot/flop) **star:22** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![godoc][GoDoc]](https://godoc.org/github.com/homedepot/flop)
* [baraka](https://github.com/xis/baraka) **star:17** A library to process http file uploads easily.   [![There was an update last week][Green]](https://github.com/xis/baraka)   [![godoc][GoDoc]](https://godoc.org/github.com/xis/baraka)
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:14** Copy files for humans.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hugocarreira/go-decent-copy)   [![godoc][GoDoc]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
* [gut/yos](https://github.com/1set/gut) **star:14** Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.   [![godoc][GoDoc]](https://godoc.org/github.com/1set/gut)
* [parquet](https://github.com/parsyl/parquet) **star:9** Read and write [parquet](https://parquet.apache.org) files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/parsyl/parquet)   [![godoc][GoDoc]](https://godoc.org/github.com/parsyl/parquet)
* [todotxt](https://github.com/1set/todotxt) **star:7** Go library for Gina Trapani's [*todo.txt*](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [*todo.txt* format](https://github.com/todotxt/todo.txt).   [![godoc][GoDoc]](https://godoc.org/github.com/1set/todotxt)

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) **star:2735** Arbitrary-precision fixed-point decimal numbers.   [![star > 2000][Awesome]](https://github.com/shopspring/decimal)   [![godoc][GoDoc]](https://godoc.org/github.com/shopspring/decimal)
* [go-money](https://github.com/rhymond/go-money) **star:865** Implementation of Fowler's Money pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/rhymond/go-money)
* [accounting](https://github.com/leekchan/accounting) **star:615** money and currency formatting for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/leekchan/accounting)
* [go-finance](https://github.com/FlashBoys/go-finance) **star:535** Comprehensive financial markets data in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/FlashBoys/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/FlashBoys/go-finance)
* [techan](https://github.com/sdcoffey/techan) **star:354** Technical analysis library with advanced market analysis and trading strategies.   [![godoc][GoDoc]](https://godoc.org/github.com/sdcoffey/techan)
* [currency](https://github.com/bojanz/currency) **star:231** Handles currency amounts, provides currency information and formatting.   [![godoc][GoDoc]](https://godoc.org/github.com/bojanz/currency)
* [orderbook](https://github.com/i25959341/orderbook) **star:176** Matching Engine for Limit Order Book in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/i25959341/orderbook)   [![godoc][GoDoc]](https://godoc.org/github.com/i25959341/orderbook)
* [go-finance](https://github.com/alpeb/go-finance) **star:88** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   [![godoc][GoDoc]](https://godoc.org/github.com/alpeb/go-finance)
* [transaction](https://github.com/claygod/transaction) **star:82** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/transaction)
* [vat](https://github.com/dannyvankooten/vat) **star:81** VAT number validation & EU VAT rates.   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/vat)
* [ofxgo](https://github.com/aclindsa/ofxgo) **star:80** Query OFX servers and/or parse the responses (with example command-line client).   [![godoc][GoDoc]](https://godoc.org/github.com/aclindsa/ofxgo)
* [go-finnhub](https://github.com/m1/go-finnhub) **star:50** Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges.   [![It hasn't been updated in the last year][Yellow]](https://github.com/m1/go-finnhub)   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-finnhub)   [![Archived][Archived]](https://github.com/m1/go-finnhub)
* [currency](https://github.com/bnkamalesh/currency) **star:36** High performant & accurate currency computation package.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/currency)
* [sleet](https://github.com/BoltApp/sleet) **star:27** One unified interface for multiple Payment Service Providers (PsP) to process online payment.   [![godoc][GoDoc]](https://godoc.org/github.com/BoltApp/sleet)
* [fastme](https://github.com/newity/fastme) **star:19** Fast extensible matching engine Go implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/newity/fastme)
* [go-finance](https://github.com/pieterclaerhout/go-finance) **star:4** Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pieterclaerhout/go-finance)   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-finance)

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) **star:1114** CSRF protection middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/nosurf)
* [binding](https://github.com/mholt/binding) **star:783** Binds form and JSON data from net/http Request to struct.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/binding)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/binding)
* [gorilla/csrf](https://github.com/gorilla/csrf) **star:609** CSRF protection for Go web applications & services.   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/csrf)
* [form](https://github.com/go-playground/form) **star:442** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/form)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/form)
* [conform](https://github.com/leebenson/conform) **star:211** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   [![godoc][GoDoc]](https://godoc.org/github.com/leebenson/conform)
* [formam](https://github.com/monoculum/formam) **star:155** decode form's values into a struct.   [![godoc][GoDoc]](https://godoc.org/github.com/monoculum/formam)
* [forms](https://github.com/albrow/forms) **star:115** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/forms)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/forms)
* [qs](https://github.com/sonh/qs) **star:55** Go module for encoding structs into URL query parameters.   [![godoc][GoDoc]](https://godoc.org/github.com/sonh/qs)
* [bind](https://github.com/robfig/bind) **star:25** Bind form data to any Go values.   [![It hasn't been updated in the last year][Yellow]](https://github.com/robfig/bind)   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/bind)
* [queryparam](https://github.com/tomwright/queryparam) **star:7** Decode `url.Values` into usable struct values of standard or custom types.   [![godoc][GoDoc]](https://godoc.org/github.com/tomwright/queryparam)

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) **star:1180** Useful collection of helpfully functional Go collection utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tobyhede/go-underscore)   [![godoc][GoDoc]](https://godoc.org/github.com/tobyhede/go-underscore)
* [fpGo](https://github.com/TeaEntityLab/fpGo) **star:162** Monad, Functional Programming features for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/TeaEntityLab/fpGo)   [![godoc][GoDoc]](https://godoc.org/github.com/TeaEntityLab/fpGo)
* [fuego](https://github.com/seborama/fuego) **star:82** Functional Experiment in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/fuego)

## Game Development

*Awesome game development libraries.*

* [Ebiten](https://github.com/hajimehoshi/ebiten) **star:4012** dead simple 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/hajimehoshi/ebiten)   [![There was an update last week][Green]](https://github.com/hajimehoshi/ebiten)   [![godoc][GoDoc]](https://godoc.org/github.com/hajimehoshi/ebiten)
* [Leaf](https://github.com/name5566/leaf) **star:3836** Lightweight game server framework.   [![star > 2000][Awesome]](https://github.com/name5566/leaf)   [![godoc][GoDoc]](https://godoc.org/github.com/name5566/leaf)   [![Contains Chinese documents][CN]](https://github.com/name5566/leaf)
* [Pixel](https://github.com/faiface/pixel) **star:3324** Hand-crafted 2D game library in Go.   [![star > 2000][Awesome]](https://github.com/faiface/pixel)   [![godoc][GoDoc]](https://godoc.org/github.com/faiface/pixel)
* [goworld](https://github.com/xiaonanln/goworld) **star:1742** Scalable game server engine, featuring space-entity framework and hot-swapping.   [![godoc][GoDoc]](https://godoc.org/github.com/xiaonanln/goworld)   [![Contains Chinese documents][CN]](https://github.com/xiaonanln/goworld)
* [nano](https://github.com/lonng/nano) **star:1560** Lightweight, facility, high performance golang based game server framework.   [![godoc][GoDoc]](https://godoc.org/github.com/lonng/nano)   [![Contains Chinese documents][CN]](https://github.com/lonng/nano)
* [go-sdl2](https://github.com/veandco/go-sdl2) **star:1475** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [engo](https://github.com/EngoEngine/engo) **star:1324** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![godoc][GoDoc]](https://godoc.org/github.com/EngoEngine/engo)
* [g3n](https://github.com/g3n/engine) **star:1296** Go 3D Game Engine.   [![godoc][GoDoc]](https://godoc.org/github.com/g3n/engine)
* [termloop](https://github.com/JoelOtter/termloop) **star:1181** Terminal-based game engine for Go, built on top of Termbox.   [![godoc][GoDoc]](https://godoc.org/github.com/JoelOtter/termloop)
* [gonet](https://github.com/xtaci/gonet) **star:1121** Game server skeleton implemented with golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xtaci/gonet)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gonet)
* [Pitaya](https://github.com/topfreegames/pitaya) **star:865** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/pitaya)
* [Oak](https://github.com/oakmound/oak) **star:833** Pure Go game engine.   [![There was an update last week][Green]](https://github.com/oakmound/oak)   [![godoc][GoDoc]](https://godoc.org/github.com/oakmound/oak)
* [raylib-go](https://github.com/gen2brain/raylib-go) **star:567** Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [Azul3D](https://github.com/azul3d/engine) **star:470** 3D game engine written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/azul3d/engine)
* [go-astar](https://github.com/beefsack/go-astar) **star:406** Go implementation of the A\* path finding algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-astar)
* [go3d](https://github.com/ungerik/go3d) **star:189** Performance oriented 2D/3D math package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go3d)
* [prototype](https://github.com/gonutz/prototype) **star:42** Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API.   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/prototype)
* [tile](https://github.com/kelindar/tile) **star:15** Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export.   [![godoc][GoDoc]](https://godoc.org/github.com/kelindar/tile)

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:2369** .NET LINQ-like query methods for Go.   [![star > 2000][Awesome]](https://github.com/ahmetalpbalkan/go-linq)   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
* [jennifer](https://github.com/dave/jennifer) **star:1824** Generate arbitrary Go code without templates.   [![godoc][GoDoc]](https://godoc.org/github.com/dave/jennifer)
* [gen](https://github.com/clipperhouse/gen) **star:1224** Code generation tool for ‘generics’-like functionality.   [![It hasn't been updated in the last year][Yellow]](https://github.com/clipperhouse/gen)   [![godoc][GoDoc]](https://godoc.org/github.com/clipperhouse/gen)
* [goderive](https://github.com/awalterschulze/goderive) **star:840** Derives functions from input types.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/goderive)
* [GoWrap](https://github.com/hexdigest/gowrap) **star:442** Generate decorators for Go interfaces using simple templates.   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gowrap)
* [interfaces](https://github.com/rjeczalik/interfaces) **star:270** Command line tool for generating interface definitions.   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/interfaces)
* [go-enum](https://github.com/abice/go-enum) **star:157** Code generation for enums from code comments.   [![godoc][GoDoc]](https://godoc.org/github.com/abice/go-enum)
* [pkgreflect](https://github.com/ungerik/pkgreflect) **star:96** Go preprocessor for package scoped reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/pkgreflect)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/pkgreflect)
* [efaceconv](https://github.com/t0pep0/efaceconv) **star:47** Code generation tool for high performance conversion from interface{} to immutable type without allocations.   [![It hasn't been updated in the last year][Yellow]](https://github.com/t0pep0/efaceconv)   [![godoc][GoDoc]](https://godoc.org/github.com/t0pep0/efaceconv)
* [gotype](https://github.com/wzshiming/gotype) **star:33** Golang source code parsing, usage like reflect package.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/gotype)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/gotype)
* [generis](https://github.com/senselogic/GENERIS) **star:22** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
* [go-xray](https://github.com/pieterclaerhout/go-xray) **star:16** Helpers for making the use of reflection easier.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pieterclaerhout/go-xray)   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-xray)
* [typeregistry](https://github.com/xiaoxin01/typeregistry) **star:8** A library to create type dynamically.   [![godoc][GoDoc]](https://godoc.org/github.com/xiaoxin01/typeregistry)

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) **star:7225** Geolocation DB with spatial index and realtime geofencing.   [![star > 2000][Awesome]](https://github.com/tidwall/tile38)   [![There was an update last week][Green]](https://github.com/tidwall/tile38)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/tile38)
* [S2 geometry](https://github.com/golang/geo) **star:1160** S2 geometry library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/golang/geo)
* [mbtileserver](https://github.com/consbio/mbtileserver) **star:203** A simple Go-based server for map tiles stored in mbtiles format.   [![godoc][GoDoc]](https://godoc.org/github.com/consbio/mbtileserver)
* [osm](https://github.com/paulmach/osm) **star:148** Library for reading, writing and working with OpenStreetMap data and APIs.   [![There was an update last week][Green]](https://github.com/paulmach/osm)   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/osm)
* [geocache](https://github.com/melihmucuk/geocache) **star:128** In-memory cache that is suitable for geolocation based applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/melihmucuk/geocache)   [![godoc][GoDoc]](https://godoc.org/github.com/melihmucuk/geocache)
* [WGS84](https://github.com/wroge/wgs84) **star:59** Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).   [![godoc][GoDoc]](https://godoc.org/github.com/wroge/wgs84)
* [geoserver](https://github.com/hishamkaram/geoserver) **star:44** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/geoserver)
* [gismanager](https://github.com/hishamkaram/gismanager) **star:36** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hishamkaram/gismanager)   [![godoc][GoDoc]](https://godoc.org/github.com/hishamkaram/gismanager)
* [pbf](https://github.com/maguro/pbf) **star:23** OpenStreetMap PBF golang encoder/decoder.   [![godoc][GoDoc]](https://godoc.org/github.com/maguro/pbf)
* [S2 geojson](https://github.com/pantrif/s2-geojson) **star:10** Convert geojson to s2 cells & demonstrating some S2 geometry features on map.   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/s2-geojson)

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) **star:9887** Compiler from Go to JavaScript.   [![star > 2000][Awesome]](https://github.com/gopherjs/gopherjs)   [![godoc][GoDoc]](https://godoc.org/github.com/gopherjs/gopherjs)
* [tardisgo](https://github.com/tardisgo/tardisgo) **star:404** Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tardisgo/tardisgo)   [![godoc][GoDoc]](https://godoc.org/github.com/tardisgo/tardisgo)
* [c4go](https://github.com/Konstantin8105/c4go) **star:246** Transpile C code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/c4go)
* [f4go](https://github.com/Konstantin8105/f4go) **star:23** Transpile FORTRAN 77 code to Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/Konstantin8105/f4go)

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) **star:4875** A high-performance and low-cost goroutine pool in Go.   [![star > 2000][Awesome]](https://github.com/panjf2000/ants)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/ants)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/ants)
* [goworker](https://github.com/benmanns/goworker) **star:2504** goworker is a Go-based background worker.   [![star > 2000][Awesome]](https://github.com/benmanns/goworker)   [![godoc][GoDoc]](https://godoc.org/github.com/benmanns/goworker)
* [tunny](https://github.com/Jeffail/tunny) **star:1832** Goroutine pool for golang.   [![There was an update last week][Green]](https://github.com/Jeffail/tunny)   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/tunny)
* [grpool](https://github.com/ivpusic/grpool) **star:605** Lightweight Goroutine pool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/grpool)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/grpool)
* [pool](https://github.com/go-playground/pool) **star:588** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/pool)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pool)
* [workerpool](https://github.com/gammazero/workerpool) **star:450** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![godoc][GoDoc]](https://godoc.org/github.com/gammazero/workerpool)
* [gowp](https://github.com/xxjwxc/gowp) **star:253** gowp is concurrency limiting goroutine pool.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gowp)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gowp)
* [go-floc](https://github.com/workanator/go-floc) **star:184** Orchestrate goroutines with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/workanator/go-floc)
* [pond](https://github.com/alitto/pond) **star:167** Minimalistic and High-performance goroutine worker pool written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alitto/pond)
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:150** Control goroutines execution order.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kamildrazkiewicz/go-flow)   [![godoc][GoDoc]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
* [semaphore](https://github.com/marusama/semaphore) **star:106** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/semaphore)
* [artifex](https://github.com/borderstech/artifex) **star:96** Simple in-memory job queue for Golang using worker-based dispatching.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/artifex)
* [breaker](https://github.com/kamilsk/breaker) **star:90** Flexible mechanism to make execution flow interruptible.   [![There was an update last week][Green]](https://github.com/kamilsk/breaker)   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/breaker)
* [semaphore](https://github.com/kamilsk/semaphore) **star:83** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/semaphore)
* [worker-pool](https://github.com/vardius/worker-pool) **star:73** goworker is a Go simple async worker pool.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/worker-pool)
* [gpool](https://github.com/Sherifabdlnaby/gpool) **star:72** manages a resizeable pool of context-aware goroutines to bound concurrency.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Sherifabdlnaby/gpool)   [![godoc][GoDoc]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
* [async](https://github.com/studiosol/async) **star:69** A safe way to execute functions asynchronously, recovering them in case of panic.   [![godoc][GoDoc]](https://godoc.org/github.com/studiosol/async)
* [neilotoole/errgroup](https://github.com/neilotoole/errgroup) **star:65** Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines.   [![godoc][GoDoc]](https://godoc.org/github.com/neilotoole/errgroup)
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:57** CyclicBarrier for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/marusama/cyclicbarrier)
* [threadpool](https://github.com/shettyh/threadpool) **star:51** Golang threadpool implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/shettyh/threadpool)
* [gollback](https://github.com/vardius/gollback) **star:50** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gollback)
* [Hunch](https://github.com/AaronJan/Hunch) **star:41** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   [![godoc][GoDoc]](https://godoc.org/github.com/AaronJan/Hunch)
* [routine](https://github.com/x-mod/routine) **star:39** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![godoc][GoDoc]](https://godoc.org/github.com/x-mod/routine)
* [kyoo](https://github.com/dirkaholic/kyoo) **star:31** Provides an unlimited job queue and concurrent worker pools.   [![godoc][GoDoc]](https://godoc.org/github.com/dirkaholic/kyoo)
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:29** Run functions in parallel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/parallel-fn)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/parallel-fn)
* [async](https://github.com/reugn/async) **star:19** An alternative sync library for Go (Future, Promise, Locks).   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/async)
* [nursery](https://github.com/arunsworld/nursery) **star:19** Structured concurrency in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/arunsworld/nursery)
* [oversight](https://cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
* [goccm](https://github.com/zenthangplus/goccm) **star:18** Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently.   [![godoc][GoDoc]](https://godoc.org/github.com/zenthangplus/goccm)
* [stl](https://github.com/ssgreg/stl) **star:18** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/stl)
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) **star:15** Like `sync.WaitGroup` with error handling and concurrency control.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)
* [go-trylock](https://github.com/subchen/go-trylock) **star:13** TryLock support on read-write lock for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-trylock)
* [gohive](https://github.com/loveleshsharma/gohive) **star:13** A highly performant and easy to use Goroutine pool for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/loveleshsharma/gohive)   [![godoc][GoDoc]](https://godoc.org/github.com/loveleshsharma/gohive)
* [conexec](https://github.com/ITcathyh/conexec) **star:10** A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.   [![godoc][GoDoc]](https://godoc.org/github.com/ITcathyh/conexec)
* [queue](https://github.com/AnikHasibul/queue) **star:9** Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/AnikHasibul/queue)   [![godoc][GoDoc]](https://godoc.org/github.com/AnikHasibul/queue)
* [channelify](https://github.com/ddelizia/channelify) **star:6** Transform your function to return channels for easy and powerful parallel processing.   [![godoc][GoDoc]](https://godoc.org/github.com/ddelizia/channelify)
* [hands](https://github.com/duanckham/hands) **star:6** A process controller used to control the execution and return strategies of multiple goroutines.   [![godoc][GoDoc]](https://godoc.org/github.com/duanckham/hands)
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:5** Manage a pool of goroutines using this lightweight library with a simple API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nikhilsaraf/go-tools)   [![godoc][GoDoc]](https://godoc.org/github.com/nikhilsaraf/go-tools)
* [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) **star:4** Concurrency limiter with support for timeouts , dynamic priority and context cancellation of goroutines.   [![godoc][GoDoc]](https://godoc.org/github.com/vivek-ng/concurrency-limiter)

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [fyne](https://github.com/fyne-io/fyne) **star:12098** Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.   [![star > 2000][Awesome]](https://github.com/fyne-io/fyne)   [![There was an update last week][Green]](https://github.com/fyne-io/fyne)   [![godoc][GoDoc]](https://godoc.org/github.com/fyne-io/fyne)
* [qt](https://github.com/therecipe/qt) **star:8229** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   [![star > 2000][Awesome]](https://github.com/therecipe/qt)   [![godoc][GoDoc]](https://godoc.org/github.com/therecipe/qt)
* [ui](https://github.com/andlabs/ui) **star:7775** Platform-native GUI library for Go. Cross platform.   [![star > 2000][Awesome]](https://github.com/andlabs/ui)   [![godoc][GoDoc]](https://godoc.org/github.com/andlabs/ui)
* [Wails](https://wails.app)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
* [webview](https://github.com/zserge/webview) **star:7437** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).   [![star > 2000][Awesome]](https://github.com/zserge/webview)   [![There was an update last week][Green]](https://github.com/zserge/webview)
* [walk](https://github.com/lxn/walk) **star:5153** Windows application library kit for Go.   [![star > 2000][Awesome]](https://github.com/lxn/walk)   [![There was an update last week][Green]](https://github.com/lxn/walk)   [![godoc][GoDoc]](https://godoc.org/github.com/lxn/walk)
* [app](https://github.com/murlokswarm/app) **star:4357** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   [![star > 2000][Awesome]](https://github.com/murlokswarm/app)   [![There was an update last week][Green]](https://github.com/murlokswarm/app)   [![godoc][GoDoc]](https://godoc.org/github.com/murlokswarm/app)
* [go-astilectron](https://github.com/asticode/go-astilectron) **star:3624** Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).   [![star > 2000][Awesome]](https://github.com/asticode/go-astilectron)   [![There was an update last week][Green]](https://github.com/asticode/go-astilectron)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astilectron)
* [go-gtk](http://mattn.github.io/go-gtk/)  Go bindings for GTK.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:1962** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   [![There was an update last week][Green]](https://github.com/sciter-sdk/go-sciter)
* [gotk3](https://github.com/gotk3/gotk3) **star:1376** Go bindings for GTK3.   [![There was an update last week][Green]](https://github.com/gotk3/gotk3)   [![godoc][GoDoc]](https://godoc.org/github.com/gotk3/gotk3)
* [gowd](https://github.com/dtylman/gowd) **star:299** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dtylman/gowd)   [![godoc][GoDoc]](https://godoc.org/github.com/dtylman/gowd)

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) **star:6410** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   [![star > 2000][Awesome]](https://github.com/go-vgo/robotgo)   [![There was an update last week][Green]](https://github.com/go-vgo/robotgo)
* [systray](https://github.com/getlantern/systray) **star:1629** Cross platform Go library to place an icon and menu in the notification area.   [![godoc][GoDoc]](https://godoc.org/github.com/getlantern/systray)
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:529** OSX Desktop Notifications library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/deckarep/gosx-notifier)
* [trayhost](https://github.com/shurcooL/trayhost) **star:200** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shurcooL/trayhost)   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/trayhost)
* [go-appindicator](https://github.com/dawidd6/go-appindicator) **star:13** Go bindings for libappindicator3 C library.   [![godoc][GoDoc]](https://godoc.org/github.com/dawidd6/go-appindicator)
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:8** OSX library to notify about any (pluggable) activity on your machine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/prashantgupta24/activity-tracker)   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:7** OSX Sleep/Wake notifications in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/prashantgupta24/mac-sleep-notifier)   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [gocv](https://github.com/hybridgroup/gocv) **star:3770** Go package for computer vision using OpenCV 3.3+.   [![star > 2000][Awesome]](https://github.com/hybridgroup/gocv)   [![There was an update last week][Green]](https://github.com/hybridgroup/gocv)   [![godoc][GoDoc]](https://godoc.org/github.com/hybridgroup/gocv)
* [imaging](https://github.com/disintegration/imaging) **star:3473** Simple Go image processing package.   [![star > 2000][Awesome]](https://github.com/disintegration/imaging)   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/imaging)
* [imaginary](https://github.com/h2non/imaginary) **star:3446** Fast and simple HTTP microservice for image resizing.   [![star > 2000][Awesome]](https://github.com/h2non/imaginary)   [![There was an update last week][Green]](https://github.com/h2non/imaginary)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/imaginary)
* [bild](https://github.com/anthonynsimon/bild) **star:3022** Collection of image processing algorithms in pure Go.   [![star > 2000][Awesome]](https://github.com/anthonynsimon/bild)   [![godoc][GoDoc]](https://godoc.org/github.com/anthonynsimon/bild)
* [ln](https://github.com/fogleman/ln) **star:2830** 3D line art rendering in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/ln)   [![It hasn't been updated in the last year][Yellow]](https://github.com/fogleman/ln)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/ln)
* [gg](https://github.com/fogleman/gg) **star:2630** 2D rendering in pure Go.   [![star > 2000][Awesome]](https://github.com/fogleman/gg)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/gg)
* [resize](https://github.com/nfnt/resize) **star:2564** Image resizing for Go with common interpolation methods.   [![star > 2000][Awesome]](https://github.com/nfnt/resize)   [![godoc][GoDoc]](https://godoc.org/github.com/nfnt/resize)
* [pt](https://github.com/fogleman/pt) **star:1913** Path tracing engine written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fogleman/pt)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/pt)
* [svgo](https://github.com/ajstarks/svgo) **star:1589** Go Language Library for SVG generation.   [![godoc][GoDoc]](https://godoc.org/github.com/ajstarks/svgo)
* [smartcrop](https://github.com/muesli/smartcrop) **star:1490** Finds good crops for arbitrary images and crop sizes.   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/smartcrop)
* [gift](https://github.com/disintegration/gift) **star:1406** Package of image processing filters.   [![godoc][GoDoc]](https://godoc.org/github.com/disintegration/gift)
* [picfit](https://github.com/thoas/picfit) **star:1402** An image resizing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/picfit)
* [bimg](https://github.com/h2non/bimg) **star:1328** Small package for fast and efficient image processing using libvips.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/bimg)
* [imagick](https://github.com/gographics/imagick) **star:1252** Go binding to ImageMagick's MagickWand C API.   [![godoc][GoDoc]](https://godoc.org/github.com/gographics/imagick)
* [go-opencv](https://github.com/lazywei/go-opencv) **star:1218** Go bindings for OpenCV.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lazywei/go-opencv)   [![godoc][GoDoc]](https://godoc.org/github.com/lazywei/go-opencv)
* [geopattern](https://github.com/pravj/geopattern) **star:1105** Create beautiful generative image patterns from a string.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pravj/geopattern)   [![godoc][GoDoc]](https://godoc.org/github.com/pravj/geopattern)
* [stegify](https://github.com/DimitarPetrov/stegify) **star:916** Go tool for LSB steganography, capable of hiding any file within an image.   [![godoc][GoDoc]](https://godoc.org/github.com/DimitarPetrov/stegify)
* [canvas](https://github.com/tdewolff/canvas) **star:602** Vector graphics to PDF, SVG or rasterized image.   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/canvas)
* [draft](https://github.com/lucasepe/draft) **star:476** Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax.   [![godoc][GoDoc]](https://godoc.org/github.com/lucasepe/draft)
* [image2ascii](https://github.com/qeesung/image2ascii) **star:476** Convert image to ASCII.   [![It hasn't been updated in the last year][Yellow]](https://github.com/qeesung/image2ascii)   [![godoc][GoDoc]](https://godoc.org/github.com/qeesung/image2ascii)
* [govips](https://github.com/davidbyttow/govips) **star:421** A lightning fast image processing and resizing library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/davidbyttow/govips)
* [mort](https://github.com/aldor007/mort) **star:418** Storage and image processing server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aldor007/mort)
* [govatar](https://github.com/o1egl/govatar) **star:407** Library and CMD tool for generating funny avatars.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/govatar)
* [goimagehash](https://github.com/corona10/goimagehash) **star:380** Go Perceptual image hashing package.   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimagehash)
* [go-nude](https://github.com/koyachi/go-nude) **star:320** Nudity detection with Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/koyachi/go-nude)   [![godoc][GoDoc]](https://godoc.org/github.com/koyachi/go-nude)
* [rez](https://github.com/bamiaux/rez) **star:200** Image resizing in pure Go and SIMD.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bamiaux/rez)   [![godoc][GoDoc]](https://godoc.org/github.com/bamiaux/rez)
* [darkroom](https://github.com/gojek/darkroom) **star:158** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![There was an update last week][Green]](https://github.com/gojek/darkroom)   [![godoc][GoDoc]](https://godoc.org/github.com/gojek/darkroom)
* [img](https://github.com/hawx/img) **star:137** Selection of image manipulation tools.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hawx/img)   [![godoc][GoDoc]](https://godoc.org/github.com/hawx/img)
* [mergi](https://github.com/noelyahan/mergi) **star:127** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).   [![godoc][GoDoc]](https://godoc.org/github.com/noelyahan/mergi)
* [go-cairo](https://github.com/ungerik/go-cairo) **star:97** Go binding for the cairo graphics library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-cairo)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-cairo)
* [steganography](https://github.com/auyer/steganography) **star:90** Pure Go Library for LSB steganography.   [![godoc][GoDoc]](https://godoc.org/github.com/auyer/steganography)
* [gltf](https://github.com/qmuntal/gltf) **star:88** Efficient and robust glTF 2.0 reader, writer and validator.   [![There was an update last week][Green]](https://github.com/qmuntal/gltf)   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/gltf)
* [cameron](https://github.com/aofei/cameron) **star:66** An avatar generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/cameron)
* [go-gd](https://github.com/bolknote/go-gd) **star:52** Go binding for GD library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bolknote/go-gd)   [![godoc][GoDoc]](https://godoc.org/github.com/bolknote/go-gd)
* [goimghdr](https://github.com/corona10/goimghdr) **star:36** The imghdr module determines the type of image contained in a file for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corona10/goimghdr)   [![godoc][GoDoc]](https://godoc.org/github.com/corona10/goimghdr)
* [gridder](https://github.com/shomali11/gridder) **star:36** A Grid based 2D Graphics library.   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/gridder)
* [tga](https://github.com/ftrvxmtrx/tga) **star:26** Package tga is a TARGA image format decoder/encoder.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ftrvxmtrx/tga)   [![godoc][GoDoc]](https://godoc.org/github.com/ftrvxmtrx/tga)   [![Archived][Archived]](https://github.com/ftrvxmtrx/tga)
* [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:25** Port of webcolors library from Python to Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jyotiska/go-webcolors)   [![godoc][GoDoc]](https://godoc.org/github.com/jyotiska/go-webcolors)
* [mpo](https://github.com/donatj/mpo) **star:6** Decoder and conversion tool for MPO 3D Photos.   [![godoc][GoDoc]](https://godoc.org/github.com/donatj/mpo)
* [webp-server](https://github.com/mehdipourfar/webp-server) **star:4** Simple and minimal image server capable of storing, resizing, converting and caching images.   [![godoc][GoDoc]](https://godoc.org/github.com/mehdipourfar/webp-server)

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) **star:1664** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [mainflux](https://github.com/Mainflux/mainflux) **star:1279** Industrial IoT Messaging and Device Management Server.   [![There was an update last week][Green]](https://github.com/Mainflux/mainflux)   [![godoc][GoDoc]](https://godoc.org/github.com/Mainflux/mainflux)
* [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
* [gatt](https://github.com/paypal/gatt) **star:951** Gatt is a Go package for building Bluetooth Low Energy peripherals.   [![godoc][GoDoc]](https://godoc.org/github.com/paypal/gatt)
* [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [connectordb](https://github.com/connectordb/connectordb) **star:250** Open-Source Platform for Quantified Self & IoT.   [![There was an update last week][Green]](https://github.com/connectordb/connectordb)   [![godoc][GoDoc]](https://godoc.org/github.com/connectordb/connectordb)
* [devices](https://github.com/goiot/devices) **star:239** Suite of libraries for IoT devices, experimental for x/exp/io.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goiot/devices)   [![godoc][GoDoc]](https://godoc.org/github.com/goiot/devices)
* [sensorbee](https://github.com/sensorbee/sensorbee) **star:200** Lightweight stream processing engine for IoT.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sensorbee/sensorbee)   [![godoc][GoDoc]](https://godoc.org/github.com/sensorbee/sensorbee)
* [huego](https://github.com/amimof/huego) **star:175** An extensive Philips Hue client library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/amimof/huego)
* [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
* [eywa](https://github.com/xcodersun/eywa) **star:46** Project Eywa is essentially a connection manager that keeps track of connected devices.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xcodersun/eywa)   [![godoc][GoDoc]](https://godoc.org/github.com/xcodersun/eywa)

## Job Scheduler

*Libraries for scheduling jobs.*

* [JobRunner](https://github.com/bamzi/jobrunner) **star:825** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   [![godoc][GoDoc]](https://godoc.org/github.com/bamzi/jobrunner)
* [gron](https://github.com/roylee0704/gron) **star:822** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   [![godoc][GoDoc]](https://godoc.org/github.com/roylee0704/gron)
* [gocron](https://github.com/go-co-op/gocron) **star:500** Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron).   [![There was an update last week][Green]](https://github.com/go-co-op/gocron)   [![godoc][GoDoc]](https://godoc.org/github.com/go-co-op/gocron)
* [jobs](https://github.com/albrow/jobs) **star:482** Persistent and flexible background jobs library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/albrow/jobs)   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/jobs)
* [scheduler](https://github.com/carlescere/scheduler) **star:355** Cronjobs scheduling made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/scheduler)
* [go-cron](https://github.com/rk/go-cron) **star:203** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   [![godoc][GoDoc]](https://godoc.org/github.com/rk/go-cron)
* [clockwerk](http://github.com/onatm/clockwerk)  Go package to schedule periodic jobs using a simple, fluent syntax.
* [go-quartz](https://github.com/reugn/go-quartz) **star:104** Simple, zero-dependency scheduling library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-quartz)
* [leprechaun](https://github.com/kilgaloon/leprechaun) **star:78** Job scheduler that supports webhooks, crons and classic scheduling.   [![godoc][GoDoc]](https://godoc.org/github.com/kilgaloon/leprechaun)
* [clockwork](https://github.com/whiteShtef/clockwork) **star:24** Simple and intuitive job scheduling library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/whiteShtef/clockwork)
* [cronticker](https://github.com/krayzpipes/cronticker) **star:1** A ticker implementation to support cron schedules.   [![godoc][GoDoc]](https://godoc.org/github.com/krayzpipes/cronticker)

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) **star:7662** Get a JSON value with one line of code.   [![star > 2000][Awesome]](https://github.com/tidwall/gjson)   [![godoc][GoDoc]](https://godoc.org/github.com/tidwall/gjson)
* [gojson](https://github.com/ChimeraCoder/gojson) **star:2315** Automatically generate Go (golang) struct definitions from example JSON.   [![star > 2000][Awesome]](https://github.com/ChimeraCoder/gojson)   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/gojson)
* [fastjson](https://github.com/valyala/fastjson) **star:1116** Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fastjson)
* [kazaam](https://github.com/Qntfy/kazaam) **star:176** API for arbitrary transformation of JSON documents.   [![godoc][GoDoc]](https://godoc.org/github.com/Qntfy/kazaam)
* [gojq](https://github.com/elgs/gojq) **star:161** JSON query in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/gojq)
* [jsondiff](https://github.com/wI2L/jsondiff) **star:101** JSON diff library for Go based on RFC6902 (JSON Patch).   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jsondiff)
* [jettison](https://github.com/wI2L/jettison) **star:96** Fast and flexible JSON encoder for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/wI2L/jettison)
* [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
* [jsongo](https://github.com/ricardolonga/jsongo) **star:95** Fluent API to make it easier to create Json objects.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ricardolonga/jsongo)   [![godoc][GoDoc]](https://godoc.org/github.com/ricardolonga/jsongo)
* [gjo](https://github.com/skanehira/gjo) **star:86** Small utility to create JSON objects.   [![godoc][GoDoc]](https://godoc.org/github.com/skanehira/gjo)
* [JayDiff](https://github.com/yazgazan/jaydiff) **star:80** JSON diff utility written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yazgazan/jaydiff)   [![godoc][GoDoc]](https://godoc.org/github.com/yazgazan/jaydiff)
* [json2go](https://github.com/m-zajac/json2go) **star:72** Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.   [![godoc][GoDoc]](https://godoc.org/github.com/m-zajac/json2go)
* [jsonf](https://github.com/miolini/jsonf) **star:60** Console tool for highlighted formatting and struct query fetching JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/jsonf)
* [ajson](https://github.com/spyzhov/ajson) **star:53** Abstract JSON for golang with JSONPath support.   [![godoc][GoDoc]](https://godoc.org/github.com/spyzhov/ajson)
* [mp](https://github.com/sanbornm/mp) **star:44** Simple cli email parser. It currently takes stdin and outputs JSON.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sanbornm/mp)   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/mp)
* [go-respond](https://github.com/nicklaw5/go-respond) **star:40** Go package for handling common HTTP JSON responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nicklaw5/go-respond)   [![godoc][GoDoc]](https://godoc.org/github.com/nicklaw5/go-respond)
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:10** Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ddymko/go-jsonerror)   [![godoc][GoDoc]](https://godoc.org/github.com/ddymko/go-jsonerror)
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:9** Go bindings based on the JSON API errors reference.   [![It hasn't been updated in the last year][Yellow]](https://github.com/AmuzaTkts/jsonapi-errors)   [![godoc][GoDoc]](https://godoc.org/github.com/AmuzaTkts/jsonapi-errors)
* [jsonhal](https://github.com/RichardKnop/jsonhal) **star:9** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/jsonhal)
* [dynjson](https://github.com/cocoonspace/dynjson) **star:6** Client-customizable JSON formats for dynamic APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/cocoonspace/dynjson)
* [ej](https://github.com/lucassscaravelli/ej) **star:5** Write and read JSON from different sources succinctly.   [![godoc][GoDoc]](https://godoc.org/github.com/lucassscaravelli/ej)
* [epoch](https://github.com/vtopc/epoch) **star:4** Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from build-in time.Time type in JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/vtopc/epoch)
* [jzon](https://github.com/zerosnake0/jzon) **star:3** JSON library with standard compatible API/behavior.   [![godoc][GoDoc]](https://godoc.org/github.com/zerosnake0/jzon)
* [mapslice-json](https://github.com/mickep76/mapslice-json) **star:3** Go MapSlice for ordered marshal/ unmarshal of maps in JSON.   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/mapslice-json)
* [jsonic](https://github.com/sinhashubham95/jsonic) **star:1** Utilities to handle and query JSON without defining structs in a type safe manner.   [![godoc][GoDoc]](https://godoc.org/github.com/sinhashubham95/jsonic)

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) **star:16995** Structured logger for Go.   [![star > 2000][Awesome]](https://github.com/Sirupsen/logrus)   [![There was an update last week][Green]](https://github.com/Sirupsen/logrus)   [![godoc][GoDoc]](https://godoc.org/github.com/Sirupsen/logrus)
* [zap](https://github.com/uber-go/zap) **star:11747** Fast, structured, leveled logging in Go.   [![star > 2000][Awesome]](https://github.com/uber-go/zap)   [![There was an update last week][Green]](https://github.com/uber-go/zap)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/zap)
* [zerolog](https://github.com/rs/zerolog) **star:4255** Zero-allocation JSON logger.   [![star > 2000][Awesome]](https://github.com/rs/zerolog)   [![There was an update last week][Green]](https://github.com/rs/zerolog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/zerolog)
* [spew](https://github.com/davecgh/go-spew) **star:4189** Implements a deep pretty printer for Go data structures to aid in debugging.   [![star > 2000][Awesome]](https://github.com/davecgh/go-spew)   [![godoc][GoDoc]](https://godoc.org/github.com/davecgh/go-spew)
* [glog](https://github.com/golang/glog) **star:2645** Leveled execution logs for Go.   [![star > 2000][Awesome]](https://github.com/golang/glog)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/glog)
* [lumberjack](https://github.com/natefinch/lumberjack) **star:2344** Simple rolling logger, implements io.WriteCloser.   [![star > 2000][Awesome]](https://github.com/natefinch/lumberjack)   [![There was an update last week][Green]](https://github.com/natefinch/lumberjack)   [![godoc][GoDoc]](https://godoc.org/github.com/natefinch/lumberjack)
* [tail](https://github.com/hpcloud/tail) **star:1992** Go package striving to emulate the features of the BSD tail program.   [![godoc][GoDoc]](https://godoc.org/github.com/hpcloud/tail)
* [seelog](https://github.com/cihub/seelog) **star:1530** Logging functionality with flexible dispatching, filtering, and formatting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cihub/seelog)   [![godoc][GoDoc]](https://godoc.org/github.com/cihub/seelog)
* [log](https://github.com/apex/log) **star:1067** Structured logging package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/apex/log)
* [log15](https://github.com/inconshreveable/log15) **star:993** Simple, powerful logging for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/log15)
* [onelog](https://github.com/francoispqt/onelog) **star:383** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/francoispqt/onelog)   [![godoc][GoDoc]](https://godoc.org/github.com/francoispqt/onelog)
* [logxi](https://github.com/mgutz/logxi) **star:346** 12-factor app logger that is fast and makes you happy.   [![godoc][GoDoc]](https://godoc.org/github.com/mgutz/logxi)
* [phuslu/log](https://github.com/phuslu/log) **star:292** Structured Logging Made Easy.   [![godoc][GoDoc]](https://godoc.org/github.com/phuslu/log)
* [logutils](https://github.com/hashicorp/logutils) **star:289** Utilities for slightly better logging in Go (Golang) extending the standard logger.   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/logutils)
* [log](https://github.com/go-playground/log) **star:276** Simple, configurable and scalable Structured Logging for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/log)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/log)
* [go-logger](https://github.com/apsdehal/go-logger) **star:261** Simple logger of Go Programs, with level handlers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/apsdehal/go-logger)   [![godoc][GoDoc]](https://godoc.org/github.com/apsdehal/go-logger)
* [httpretty](https://github.com/henvic/httpretty) **star:206** Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest).   [![godoc][GoDoc]](https://godoc.org/github.com/henvic/httpretty)
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:164** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkiller/rollingWriter)
* [sqldb-logger](https://github.com/simukti/sqldb-logger) **star:150** A logger for Go SQL database driver without modify existing *sql.DB stdlib usage.   [![godoc][GoDoc]](https://godoc.org/github.com/simukti/sqldb-logger)
* [logger](https://github.com/azer/logger) **star:148** Minimalistic logging library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/azer/logger)
* [xlog](https://github.com/rs/xlog) **star:135** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/xlog)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xlog)
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:115** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-log)   [![Contains Chinese documents][CN]](https://github.com/go-ozzo/ozzo-log)
* [logur](https://github.com/logur/logur) **star:108** An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc).   [![godoc][GoDoc]](https://godoc.org/github.com/logur/logur)
* [glg](https://github.com/kpango/glg) **star:97** glg is simple and fast leveled logging library for Go.   [![There was an update last week][Green]](https://github.com/kpango/glg)   [![godoc][GoDoc]](https://godoc.org/github.com/kpango/glg)
* [log-voyage](https://github.com/firstrow/logvoyage) **star:88** Full-featured logging saas written in golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/logvoyage)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/logvoyage)
* [stdlog](https://github.com/alexcesaro/log) **star:45** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alexcesaro/log)   [![godoc][GoDoc]](https://godoc.org/github.com/alexcesaro/log)
* [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:40** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![godoc][GoDoc]](https://godoc.org/github.com/utahta/go-cronowriter)
* [gologger](https://github.com/sadlil/gologger) **star:39** Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/gologger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/gologger)   [![Archived][Archived]](https://github.com/sadlil/gologger)
* [go-log](https://github.com/ian-kent/go-log) **star:37** Log4j implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/go-log)
* [logex](https://github.com/chzyer/logex) **star:36** Golang log lib, supports tracking and level, wrap by standard log lib.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chzyer/logex)   [![godoc][GoDoc]](https://godoc.org/github.com/chzyer/logex)
* [go-log](https://github.com/siddontang/go-log) **star:28** Log lib supports level and multi handlers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/siddontang/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/siddontang/go-log)
* [distillog](https://github.com/amoghe/distillog) **star:26** distilled levelled logging (think of it as stdlib + log levels).   [![It hasn't been updated in the last year][Yellow]](https://github.com/amoghe/distillog)   [![godoc][GoDoc]](https://godoc.org/github.com/amoghe/distillog)
* [journald](https://github.com/ssgreg/journald) **star:26** Go implementation of systemd Journal's native API for logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ssgreg/journald)   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/journald)
* [logrusly](https://github.com/sebest/logrusly) **star:26** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/logrusly)
* [log](https://github.com/teris-io/log) **star:24** Structured log interface for Go cleanly separates logging facade from its implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/teris-io/log)   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/log)
* [mlog](https://github.com/jbrodriguez/mlog) **star:22** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jbrodriguez/mlog)   [![godoc][GoDoc]](https://godoc.org/github.com/jbrodriguez/mlog)
* [gomol](https://github.com/aphistic/gomol) **star:16** Multiple-output, structured logging for Go with extensible logging outputs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aphistic/gomol)   [![godoc][GoDoc]](https://godoc.org/github.com/aphistic/gomol)
* [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
* [zkits-logger](https://github.com/edoger/zkits-logger) **star:15** A powerful zero-dependency JSON logger.   [![godoc][GoDoc]](https://godoc.org/github.com/edoger/zkits-logger)
* [glo](https://github.com/lajosbencz/glo) **star:13** PHP Monolog inspired logging facility with identical severity levels.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lajosbencz/glo)   [![godoc][GoDoc]](https://godoc.org/github.com/lajosbencz/glo)
* [go-log](https://github.com/subchen/go-log) **star:11** Simple and configurable Logging in Go, with level, formatters and writers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/subchen/go-log)   [![godoc][GoDoc]](https://godoc.org/github.com/subchen/go-log)
* [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:10** `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/logrusiowriter)
* [logdump](https://github.com/ewwwwwqm/logdump) **star:9** Package for multi-level logging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ewwwwwqm/logdump)   [![godoc][GoDoc]](https://godoc.org/github.com/ewwwwwqm/logdump)
* [logmatic](https://github.com/borderstech/logmatic) **star:9** Colorized logger for Golang with dynamic log level configuration.   [![godoc][GoDoc]](https://godoc.org/github.com/borderstech/logmatic)
* [log](https://github.com/aerogo/log) **star:8** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).   [![It hasn't been updated in the last year][Yellow]](https://github.com/aerogo/log)   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/log)
* [go-log](https://github.com/pieterclaerhout/go-log) **star:7** A logging library with strack traces, object dumping and optional timestamps.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-log)
* [logo](https://github.com/mbndr/logo) **star:7** Golang logger to different configurable writers.   [![godoc][GoDoc]](https://godoc.org/github.com/mbndr/logo)
* [xlog](https://github.com/xfxdev/xlog) **star:6** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xfxdev/xlog)   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xlog)
* [kemba](https://github.com/clok/kemba) **star:4** A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications.   [![There was an update last week][Green]](https://github.com/clok/kemba)   [![godoc][GoDoc]](https://godoc.org/github.com/clok/kemba)

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) **star:7626** General Machine Learning library for Go.   [![star > 2000][Awesome]](https://github.com/sjwhitworth/golearn)   [![godoc][GoDoc]](https://godoc.org/github.com/sjwhitworth/golearn)   [![Contains Chinese documents][CN]](https://github.com/sjwhitworth/golearn)
* [gorgonia](https://github.com/gorgonia/gorgonia) **star:3812** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![star > 2000][Awesome]](https://github.com/gorgonia/gorgonia)   [![There was an update last week][Green]](https://github.com/gorgonia/gorgonia)   [![godoc][GoDoc]](https://godoc.org/github.com/gorgonia/gorgonia)
* [tfgo](https://github.com/galeone/tfgo) **star:1621** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   [![There was an update last week][Green]](https://github.com/galeone/tfgo)   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/tfgo)
* [gosseract](https://github.com/otiai10/gosseract) **star:1384** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   [![There was an update last week][Green]](https://github.com/otiai10/gosseract)   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/gosseract)
* [goml](https://github.com/cdipaolo/goml) **star:1156** On-line Machine Learning in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cdipaolo/goml)   [![godoc][GoDoc]](https://godoc.org/github.com/cdipaolo/goml)
* [gorse](https://github.com/zhenghaoz/gorse) **star:1114** An offline recommender system backend based on collaborative filtering written in Go.   [![There was an update last week][Green]](https://github.com/zhenghaoz/gorse)   [![godoc][GoDoc]](https://godoc.org/github.com/zhenghaoz/gorse)   [![Contains Chinese documents][CN]](https://github.com/zhenghaoz/gorse)
* [eaopt](https://github.com/MaxHalford/eaopt) **star:710** An evolutionary optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/MaxHalford/eaopt)
* [bayesian](https://github.com/jbrukh/bayesian) **star:699** Naive Bayesian Classification for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jbrukh/bayesian)
* [CloudForest](https://github.com/ryanbressler/CloudForest) **star:678** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ryanbressler/CloudForest)
* [gobrain](https://github.com/goml/gobrain) **star:461** Neural Networks written in go.   [![godoc][GoDoc]](https://godoc.org/github.com/goml/gobrain)
* [ocrserver](https://github.com/otiai10/ocrserver) **star:351** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   [![godoc][GoDoc]](https://godoc.org/github.com/otiai10/ocrserver)
* [onnx-go](https://github.com/owulveryck/onnx-go) **star:300** Go Interface to Open Neural Network Exchange (ONNX).   [![godoc][GoDoc]](https://godoc.org/github.com/owulveryck/onnx-go)
* [go-deep](https://github.com/patrikeh/go-deep) **star:297** A feature-rich neural network library in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/patrikeh/go-deep)
* [regommend](https://github.com/muesli/regommend) **star:288** Recommendation & collaborative filtering engine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/muesli/regommend)   [![godoc][GoDoc]](https://godoc.org/github.com/muesli/regommend)
* [go-galib](https://github.com/thoj/go-galib) **star:183** Genetic Algorithms library written in Go / golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thoj/go-galib)   [![godoc][GoDoc]](https://godoc.org/github.com/thoj/go-galib)
* [goRecommend](https://github.com/timkaye11/goRecommend) **star:168** Recommendation Algorithms library written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/timkaye11/goRecommend)   [![godoc][GoDoc]](https://godoc.org/github.com/timkaye11/goRecommend)
* [Goptuna](https://github.com/c-bata/goptuna) **star:165** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![godoc][GoDoc]](https://godoc.org/github.com/c-bata/goptuna)
* [shield](https://github.com/eaigner/shield) **star:137** Bayesian text classifier with flexible tokenizers and storage backends for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/eaigner/shield)
* [go-fann](https://github.com/white-pony/go-fann) **star:104** Go bindings for Fast Artificial Neural Networks(FANN) library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/white-pony/go-fann)   [![godoc][GoDoc]](https://godoc.org/github.com/white-pony/go-fann)
* [goga](https://github.com/tomcraven/goga) **star:100** Genetic algorithm library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tomcraven/goga)   [![godoc][GoDoc]](https://godoc.org/github.com/tomcraven/goga)
* [libsvm](https://github.com/datastream/libsvm) **star:67** libsvm golang version derived work based on LIBSVM 3.14.   [![It hasn't been updated in the last year][Yellow]](https://github.com/datastream/libsvm)   [![godoc][GoDoc]](https://godoc.org/github.com/datastream/libsvm)
* [gonet](https://github.com/dathoangnd/gonet) **star:63** Neural Network for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/dathoangnd/gonet)
* [goscore](https://github.com/asafschers/goscore) **star:63** Go Scoring API for PMML.   [![It hasn't been updated in the last year][Yellow]](https://github.com/asafschers/goscore)   [![godoc][GoDoc]](https://godoc.org/github.com/asafschers/goscore)
* [neural-go](https://github.com/schuyler/neural-go) **star:62** Multilayer perceptron network implemented in Go, with training via backpropagation.   [![godoc][GoDoc]](https://godoc.org/github.com/schuyler/neural-go)
* [go-pr](https://github.com/daviddengcn/go-pr) **star:58** Pattern recognition package in Go lang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daviddengcn/go-pr)   [![godoc][GoDoc]](https://godoc.org/github.com/daviddengcn/go-pr)
* [neat](https://github.com/jinyeom/neat) **star:58** Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).   [![It hasn't been updated in the last year][Yellow]](https://github.com/jinyeom/neat)   [![godoc][GoDoc]](https://godoc.org/github.com/jinyeom/neat)   [![Archived][Archived]](https://github.com/jinyeom/neat)
* [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) **star:51** Fast and convenient feature processing for low latency machine leraning in Go.   [![There was an update last week][Green]](https://github.com/nikolaydubina/go-featureprocessing)   [![godoc][GoDoc]](https://godoc.org/github.com/nikolaydubina/go-featureprocessing)
* [fonet](https://github.com/Fontinalis/fonet) **star:43** A Deep Neural Network library written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/Fontinalis/fonet)
* [golinear](https://github.com/danieldk/golinear) **star:41** liblinear bindings for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/golinear)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/golinear)
* [Varis](https://github.com/Xamber/Varis) **star:33** Golang Neural Network.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Xamber/Varis)   [![godoc][GoDoc]](https://godoc.org/github.com/Xamber/Varis)
* [godist](https://github.com/e-dard/godist) **star:27** Various probability distributions, and associated methods.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-dard/godist)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/godist)
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:26** Go implementation of the k-modes and k-prototypes clustering algorithms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
* [evoli](https://github.com/khezen/evoli) **star:15** Genetic Algorithm and Particle Swarm Optimization library.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/evoli)
* [probab](https://github.com/ThePaw/probab) **star:14** Probability distribution functions. Bayesian inference. Written in pure Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/probab)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/probab)
* [GoMind](https://github.com/surenderthakran/gomind) **star:10** A simplistic Neural Network Library in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surenderthakran/gomind)   [![godoc][GoDoc]](https://godoc.org/github.com/surenderthakran/gomind)
* [randomforest](https://github.com/malaschitz/randomForest) **star:7** Easy to use Random Forest library for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/malaschitz/randomForest)

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) **star:6758** Go library for Apache Kafka.   [![star > 2000][Awesome]](https://github.com/Shopify/sarama)   [![There was an update last week][Green]](https://github.com/Shopify/sarama)   [![godoc][GoDoc]](https://godoc.org/github.com/Shopify/sarama)
* [gorush](https://github.com/appleboy/gorush) **star:5203** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![star > 2000][Awesome]](https://github.com/appleboy/gorush)   [![There was an update last week][Green]](https://github.com/appleboy/gorush)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gorush)
* [machinery](https://github.com/RichardKnop/machinery) **star:4882** Asynchronous task queue/job queue based on distributed message passing.   [![star > 2000][Awesome]](https://github.com/RichardKnop/machinery)   [![There was an update last week][Green]](https://github.com/RichardKnop/machinery)   [![godoc][GoDoc]](https://godoc.org/github.com/RichardKnop/machinery)
* [Centrifugo](https://github.com/centrifugal/centrifugo) **star:4833** Real-time messaging (Websockets or SockJS) server in Go.   [![star > 2000][Awesome]](https://github.com/centrifugal/centrifugo)   [![There was an update last week][Green]](https://github.com/centrifugal/centrifugo)   [![godoc][GoDoc]](https://godoc.org/github.com/centrifugal/centrifugo)
* [go-socket.io](https://github.com/googollee/go-socket.io) **star:3900** socket.io library for golang, a realtime application framework.   [![star > 2000][Awesome]](https://github.com/googollee/go-socket.io)   [![godoc][GoDoc]](https://godoc.org/github.com/googollee/go-socket.io)
* [NATS Go Client](https://github.com/nats-io/nats) **star:3186** Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.   [![star > 2000][Awesome]](https://github.com/nats-io/nats)   [![There was an update last week][Green]](https://github.com/nats-io/nats)   [![godoc][GoDoc]](https://godoc.org/github.com/nats-io/nats)
* [Benthos](https://github.com/Jeffail/benthos) **star:2809** A message streaming bridge between a range of protocols.   [![star > 2000][Awesome]](https://github.com/Jeffail/benthos)   [![There was an update last week][Green]](https://github.com/Jeffail/benthos)   [![godoc][GoDoc]](https://godoc.org/github.com/Jeffail/benthos)
* [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) **star:2477** confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform.   [![star > 2000][Awesome]](https://github.com/confluentinc/confluent-kafka-go)   [![There was an update last week][Green]](https://github.com/confluentinc/confluent-kafka-go)   [![godoc][GoDoc]](https://godoc.org/github.com/confluentinc/confluent-kafka-go)
* [APNs2](https://github.com/sideshow/apns2) **star:2425** HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.   [![star > 2000][Awesome]](https://github.com/sideshow/apns2)   [![There was an update last week][Green]](https://github.com/sideshow/apns2)   [![godoc][GoDoc]](https://godoc.org/github.com/sideshow/apns2)
* [Mercure](https://github.com/dunglas/mercure) **star:2277** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   [![star > 2000][Awesome]](https://github.com/dunglas/mercure)   [![There was an update last week][Green]](https://github.com/dunglas/mercure)   [![godoc][GoDoc]](https://godoc.org/github.com/dunglas/mercure)
* [melody](https://github.com/olahol/melody) **star:2045** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![star > 2000][Awesome]](https://github.com/olahol/melody)   [![It hasn't been updated in the last year][Yellow]](https://github.com/olahol/melody)   [![godoc][GoDoc]](https://godoc.org/github.com/olahol/melody)
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:1959** gopush-cluster is a go push server cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Terry-Mao/gopush-cluster)   [![godoc][GoDoc]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   [![Contains Chinese documents][CN]](https://github.com/Terry-Mao/gopush-cluster)
* [go-nsq](https://github.com/nsqio/go-nsq) **star:1831** the official Go package for NSQ.   [![godoc][GoDoc]](https://godoc.org/github.com/nsqio/go-nsq)
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1230** Redis backed unified push service for server-side notifications to mobile devices.   [![godoc][GoDoc]](https://godoc.org/github.com/uniqush/uniqush-push)
* [Beaver](https://github.com/Clivern/Beaver) **star:1030** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![There was an update last week][Green]](https://github.com/Clivern/Beaver)   [![godoc][GoDoc]](https://godoc.org/github.com/Clivern/Beaver)
* [zmq4](https://github.com/pebbe/zmq4) **star:887** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/zmq4)
* [Gollum](https://github.com/trivago/gollum) **star:883** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   [![It hasn't been updated in the last year][Yellow]](https://github.com/trivago/gollum)   [![godoc][GoDoc]](https://godoc.org/github.com/trivago/gollum)
* [EventBus](https://github.com/asaskevich/EventBus) **star:858** The lightweight event bus with async compatibility.   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/EventBus)
* [Asynq](https://github.com/hibiken/asynq) **star:849** A simple, reliable, and efficient distributed task queue for Go built on top of Redis.   [![There was an update last week][Green]](https://github.com/hibiken/asynq)   [![godoc][GoDoc]](https://godoc.org/github.com/hibiken/asynq)
* [dbus](https://github.com/godbus/dbus) **star:541** Native Go bindings for D-Bus.   [![godoc][GoDoc]](https://godoc.org/github.com/godbus/dbus)
* [golongpoll](https://github.com/jcuga/golongpoll) **star:501** HTTP longpoll server library that makes web pub-sub simple.   [![godoc][GoDoc]](https://godoc.org/github.com/jcuga/golongpoll)
* [emitter](https://github.com/olebedev/emitter) **star:389** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/emitter)
* [mangos](https://github.com/nanomsg/mangos) **star:365** Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability.   [![godoc][GoDoc]](https://godoc.org/github.com/nanomsg/mangos)
* [Glue](https://github.com/desertbit/glue) **star:359** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![godoc][GoDoc]](https://godoc.org/github.com/desertbit/glue)
* [pubsub](https://github.com/tuxychandru/pubsub) **star:339** Simple pubsub package for go.   [![godoc][GoDoc]](https://godoc.org/github.com/tuxychandru/pubsub)
* [Bus](https://github.com/mustafaturan/bus) **star:178** Minimalist message bus implementation for internal communication.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaturan/bus)
* [rabtap](https://github.com/jandelgado/rabtap) **star:164** RabbitMQ swiss army knife cli app.   [![godoc][GoDoc]](https://godoc.org/github.com/jandelgado/rabtap)
* [messagebus](https://github.com/vardius/message-bus) **star:159** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
* [guble](https://github.com/smancke/guble) **star:145** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   [![It hasn't been updated in the last year][Yellow]](https://github.com/smancke/guble)   [![godoc][GoDoc]](https://godoc.org/github.com/smancke/guble)
* [oplog](https://github.com/dailymotion/oplog) **star:105** Generic oplog/replication system for REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dailymotion/oplog)   [![godoc][GoDoc]](https://godoc.org/github.com/dailymotion/oplog)
* [hub](https://github.com/leandro-lugaresi/hub) **star:90** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   [![godoc][GoDoc]](https://godoc.org/github.com/leandro-lugaresi/hub)
* [rabbus](https://github.com/rafaeljesus/rabbus) **star:85** A tiny wrapper over amqp exchanges and queues.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/rabbus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/rabbus)
* [drone-line](https://github.com/appleboy/drone-line) **star:73** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-line)
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:66** A tiny wrapper around NSQ topic and channel.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:59** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sybrexsys/RapidMQ)   [![godoc][GoDoc]](https://godoc.org/github.com/sybrexsys/RapidMQ)
* [go-mq](https://github.com/cheshir/go-mq) **star:53** RabbitMQ client with declarative configuration.   [![godoc][GoDoc]](https://godoc.org/github.com/cheshir/go-mq)
* [Commander](https://github.com/jeroenrinzema/commander) **star:52** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/jeroenrinzema/commander)
* [go-notify](https://github.com/TheCreeper/go-notify) **star:52** Native implementation of the freedesktop notification spec.   [![godoc][GoDoc]](https://godoc.org/github.com/TheCreeper/go-notify)
* [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:49** redisqueue provides a producer and consumer of a queue that uses Redis streams.   [![godoc][GoDoc]](https://godoc.org/github.com/robinjoseph08/redisqueue)
* [go-res](https://github.com/jirenius/go-res) **star:45** Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate.   [![godoc][GoDoc]](https://godoc.org/github.com/jirenius/go-res)
* [event](https://github.com/agoalofalife/event) **star:36** Implementation of the pattern observer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agoalofalife/event)   [![godoc][GoDoc]](https://godoc.org/github.com/agoalofalife/event)
* [gosd](https://github.com/alexsniffin/gosd) **star:17** A library for scheduling when to dispatch a message to a channel.   [![godoc][GoDoc]](https://godoc.org/github.com/alexsniffin/gosd)
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:16** Client library to Viessmann Vitotrol web service.   [![It hasn't been updated in the last year][Yellow]](https://github.com/maxatome/go-vitotrol)   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-vitotrol)
* [ami](https://github.com/kak-tus/ami) **star:13** Go client to reliable queues based on Redis Cluster Streams.   [![godoc][GoDoc]](https://godoc.org/github.com/kak-tus/ami)
* [jazz](https://github.com/socifi/jazz) **star:13** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/socifi/jazz)   [![godoc][GoDoc]](https://godoc.org/github.com/socifi/jazz)
* [rmqconn](https://github.com/sbabiv/rmqconn) **star:13** RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sbabiv/rmqconn)   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/rmqconn)
* [hare](https://github.com/leozz37/hare) **star:11** A user friendly library for sending messages and listening to TCP sockets.   [![godoc][GoDoc]](https://godoc.org/github.com/leozz37/hare)
* [gaurun-client](https://github.com/osamingo/gaurun-client) **star:9** Gaurun Client written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/osamingo/gaurun-client)   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gaurun-client)

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) **star:2709** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![star > 2000][Awesome]](https://github.com/unidoc/unioffice)   [![godoc][GoDoc]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) **star:7940** Golang library for reading and writing Microsoft Excel™ (XLSX) files.   [![star > 2000][Awesome]](https://github.com/360EntSecGroup-Skylar/excelize)   [![There was an update last week][Green]](https://github.com/360EntSecGroup-Skylar/excelize)   [![godoc][GoDoc]](https://godoc.org/github.com/360EntSecGroup-Skylar/excelize)
* [xlsx](https://github.com/tealeg/xlsx) **star:4748** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![star > 2000][Awesome]](https://github.com/tealeg/xlsx)   [![godoc][GoDoc]](https://godoc.org/github.com/tealeg/xlsx)
* [xlsx](https://github.com/plandem/xlsx) **star:129** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/plandem/xlsx)
* [go-excel](https://github.com/szyhf/go-excel) **star:113** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][GoDoc]](https://godoc.org/github.com/szyhf/go-excel)
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) **star:15** Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fterrag/goxlsxwriter)   [![godoc][GoDoc]](https://godoc.org/github.com/fterrag/goxlsxwriter)

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) **star:1760** A reflection based dependency injection toolkit for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/dig)
* [fx](https://github.com/uber-go/fx) **star:1731** A dependency injection based application framework for Go (built on top of dig).   [![There was an update last week][Green]](https://github.com/uber-go/fx)   [![godoc][GoDoc]](https://godoc.org/github.com/uber-go/fx)
* [container](https://github.com/golobby/container) **star:136** A powerful IoC Container with fluent and easy-to-use interface.   [![There was an update last week][Green]](https://github.com/golobby/container)   [![godoc][GoDoc]](https://godoc.org/github.com/golobby/container)
* [dingo](https://github.com/i-love-flamingo/dingo) **star:96** A dependency injection toolkit for Go, based on Guice.   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/dingo)
* [goioc/di](https://github.com/goioc/di) **star:64** Spring-inspired Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/goioc/di)
* [di](https://github.com/goava/di) **star:63** A dependency injection container for go programming language.   [![godoc][GoDoc]](https://godoc.org/github.com/goava/di)
* [alice](https://github.com/magic003/alice) **star:44** Additive dependency injection container for Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/magic003/alice)   [![godoc][GoDoc]](https://godoc.org/github.com/magic003/alice)
* [wire](https://github.com/Fs02/wire) **star:33** Strict Runtime Dependency Injection for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/wire)
* [linker](https://github.com/logrange/linker) **star:28** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][GoDoc]](https://godoc.org/github.com/logrange/linker)
* [gocontainer](https://github.com/vardius/gocontainer) **star:14** Simple Dependency Injection Container.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gocontainer)

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:20717** Set of common historical and emerging project layout patterns in the Go ecosystem.   [![star > 2000][Awesome]](https://github.com/golang-standards/project-layout)
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:860** Go application boilerplate and example applying modern practices.   [![godoc][GoDoc]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) **star:426** A Go application boilerplate template for quick starting projects following production best practices.   [![godoc][GoDoc]](https://godoc.org/github.com/lacion/cookiecutter-golang)
* [scaffold](https://github.com/catchplay/scaffold) **star:94** Scaffold generates a starter Go project layout. Lets you focus on business logic implemented.   [![It hasn't been updated in the last year][Yellow]](https://github.com/catchplay/scaffold)   [![godoc][GoDoc]](https://godoc.org/github.com/catchplay/scaffold)
* [go-sample](https://github.com/zitryss/go-sample) **star:81** A sample layout for Go application projects with the real code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zitryss/go-sample)   [![godoc][GoDoc]](https://godoc.org/github.com/zitryss/go-sample)
* [golang-templates/seed](https://github.com/golang-templates/seed) **star:78** Go application GitHub repository template.   [![There was an update last week][Green]](https://github.com/golang-templates/seed)
* [go-todo-backend](https://github.com/Fs02/go-todo-backend) **star:48** Go Todo Backend example using modular project layout for product microservice.   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/go-todo-backend)
* [gobase](https://github.com/wajox/gobase) **star:1** A simple skeleton for golang application with basic setup for real golang application.   [![godoc][GoDoc]](https://godoc.org/github.com/wajox/gobase)

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) **star:871** Collection of useful string functions ported from other languages.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/xstrings)
* [strutil](https://github.com/ozgio/strutil) **star:109** String utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ozgio/strutil)   [![godoc][GoDoc]](https://godoc.org/github.com/ozgio/strutil)
* [go-formatter](https://gitlab.com/tymonx/go-formatter)  Implements **replacement fields** surrounded by curly braces `{}` format strings.
* [gobeam/Stringy](https://github.com/gobeam/Stringy) **star:43** String manipulation library to convert string to camel case, snake case, kebab case / slugify etc.   [![godoc][GoDoc]](https://godoc.org/github.com/gobeam/Stringy)

### Uncategorized

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) **star:5821** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![star > 2000][Awesome]](https://github.com/shirou/gopsutil)   [![There was an update last week][Green]](https://github.com/shirou/gopsutil)   [![godoc][GoDoc]](https://godoc.org/github.com/shirou/gopsutil)
* [archiver](https://github.com/mholt/archiver) **star:3062** Library and command for making and extracting .zip and .tar.gz archives.   [![star > 2000][Awesome]](https://github.com/mholt/archiver)   [![godoc][GoDoc]](https://godoc.org/github.com/mholt/archiver)
* [gofakeit](https://github.com/brianvoe/gofakeit) **star:1611** Random data generator written in go.   [![There was an update last week][Green]](https://github.com/brianvoe/gofakeit)   [![godoc][GoDoc]](https://godoc.org/github.com/brianvoe/gofakeit)
* [gosms](https://github.com/haxpax/gosms) **star:1322** Your own local SMS gateway in Go that can be used to send SMS.   [![godoc][GoDoc]](https://godoc.org/github.com/haxpax/gosms)
* [go-resiliency](https://github.com/eapache/go-resiliency) **star:1140** Resiliency patterns for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/eapache/go-resiliency)
* [gatus](https://github.com/TwinProduction/gatus) **star:1118** Automated service health dashboard.   [![There was an update last week][Green]](https://github.com/TwinProduction/gatus)   [![godoc][GoDoc]](https://godoc.org/github.com/TwinProduction/gatus)
* [base64Captcha](https://github.com/mojocn/base64Captcha) **star:1028** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![godoc][GoDoc]](https://godoc.org/github.com/mojocn/base64Captcha)   [![Contains Chinese documents][CN]](https://github.com/mojocn/base64Captcha)
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:899** Generic object pool for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![Contains Chinese documents][CN]](https://github.com/jolestar/go-commons-pool)
* [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
* [shortid](https://github.com/teris-io/shortid) **star:641** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   [![godoc][GoDoc]](https://godoc.org/github.com/teris-io/shortid)
* [llvm](https://github.com/llir/llvm) **star:635** Library for interacting with LLVM IR in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/llir/llvm)
* [health](https://github.com/dimiro1/health) **star:409** Easy to use, extensible health check library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dimiro1/health)   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/health)
* [conv](https://github.com/cstockton/go-conv) **star:365** Package conv provides fast and intuitive conversions across Go types.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cstockton/go-conv)   [![godoc][GoDoc]](https://godoc.org/github.com/cstockton/go-conv)
* [ghorg](https://github.com/gabrie30/ghorg) **star:340** Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket.   [![godoc][GoDoc]](https://godoc.org/github.com/gabrie30/ghorg)
* [banner](https://github.com/dimiro1/banner) **star:324** Add beautiful banners into your Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/banner)
* [gountries](https://github.com/pariz/gountries) **star:297** Package that exposes country and subdivision data.   [![godoc][GoDoc]](https://godoc.org/github.com/pariz/gountries)
* [stateless](https://github.com/qmuntal/stateless) **star:235** A fluent library for creating state machines.   [![godoc][GoDoc]](https://godoc.org/github.com/qmuntal/stateless)
* [ffmt](https://github.com/go-ffmt/ffmt) **star:215** Beautify data display for Humans.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ffmt/ffmt)   [![Contains Chinese documents][CN]](https://github.com/go-ffmt/ffmt)
* [antch](https://github.com/antchfx/antch) **star:193** A fast, powerful and extensible web crawling & scraping framework.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/antch)   [![Contains Chinese documents][CN]](https://github.com/antchfx/antch)
* [lk](https://github.com/hyperboloide/lk) **star:182** A simple licensing library for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/lk)
* [battery](https://github.com/distatus/battery) **star:167** Cross-platform, normalized battery information library.   [![godoc][GoDoc]](https://godoc.org/github.com/distatus/battery)
* [healthcheck](https://github.com/etherlabsio/healthcheck) **star:157** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/etherlabsio/healthcheck)   [![godoc][GoDoc]](https://godoc.org/github.com/etherlabsio/healthcheck)
* [shoutrrr](https://github.com/containrrr/shoutrrr) **star:151** Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others.   [![There was an update last week][Green]](https://github.com/containrrr/shoutrrr)   [![godoc][GoDoc]](https://godoc.org/github.com/containrrr/shoutrrr)
* [stats](https://github.com/go-playground/stats) **star:145** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/stats)
* [bitio](https://github.com/icza/bitio) **star:141** Highly optimized bit-level Reader and Writer for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/icza/bitio)   [![godoc][GoDoc]](https://godoc.org/github.com/icza/bitio)
* [go-unarr](https://github.com/gen2brain/go-unarr) **star:120** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][GoDoc]](https://godoc.org/github.com/gen2brain/go-unarr)
* [turtle](https://github.com/hackebrot/turtle) **star:116** Emojis for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/hackebrot/turtle)
* [gommit](https://github.com/antham/gommit) **star:92** Analyze git commit messages to ensure they follow defined patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/gommit)
* [gotoprom](https://github.com/cabify/gotoprom) **star:88** Type-safe metrics builder wrapper library for the official Prometheus client.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cabify/gotoprom)   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/gotoprom)
* [indigo](https://github.com/osamingo/indigo) **star:77** Distributed unique ID generator of using Sonyflake and encoded by Base58.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/indigo)
* [captcha](https://github.com/steambap/captcha) **star:71** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![godoc][GoDoc]](https://godoc.org/github.com/steambap/captcha)
* [morse](https://github.com/alwindoss/morse) **star:63** Library to convert to and from morse code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alwindoss/morse)   [![godoc][GoDoc]](https://godoc.org/github.com/alwindoss/morse)
* [xkg](https://github.com/go-xkg/xkg) **star:50** X Keyboard Grabber.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-xkg/xkg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-xkg/xkg)
* [pdfgen](https://github.com/hyperboloide/pdfgen) **star:49** HTTP service to generate PDF from Json requests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hyperboloide/pdfgen)   [![godoc][GoDoc]](https://godoc.org/github.com/hyperboloide/pdfgen)
* [persian](https://github.com/mavihq/persian) **star:48** Some utilities for Persian language in go.   [![godoc][GoDoc]](https://godoc.org/github.com/mavihq/persian)
* [browscap_go](https://github.com/digitalcrab/browscap_go) **star:36** GoLang Library for [Browser Capabilities Project](http://browscap.org/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/digitalcrab/browscap_go)   [![godoc][GoDoc]](https://godoc.org/github.com/digitalcrab/browscap_go)
* [datacounter](https://github.com/miolini/datacounter) **star:36** Go counters for readers/writer/http.ResponseWriter.   [![godoc][GoDoc]](https://godoc.org/github.com/miolini/datacounter)
* [faker](https://github.com/pioz/faker) **star:33** Random fake data and struct generator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pioz/faker)
* [autoflags](https://github.com/artyom/autoflags) **star:32** Go package to automatically define command line flags from struct fields.   [![It hasn't been updated in the last year][Yellow]](https://github.com/artyom/autoflags)   [![godoc][GoDoc]](https://godoc.org/github.com/artyom/autoflags)
* [sandid](https://github.com/aofei/sandid) **star:26** Every grain of sand on earth has its own ID.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/sandid)
* [url-shortener](https://github.com/pantrif/url-shortener) **star:25** A modern, powerful, and robust URL shortener microservice with mysql support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pantrif/url-shortener)   [![godoc][GoDoc]](https://godoc.org/github.com/pantrif/url-shortener)
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
* [gosh](https://github.com/osamingo/gosh) **star:24** Provide Go Statistics Handler, Struct, Measure Method.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/gosh)
* [xdg](https://github.com/rkoesters/xdg) **star:24** FreeDesktop.org (xdg) Specs implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rkoesters/xdg)   [![godoc][GoDoc]](https://godoc.org/github.com/rkoesters/xdg)
* [metrics](https://github.com/pascaldekloe/metrics) **star:18** Library for metrics instrumentation and Prometheus exposition.   [![godoc][GoDoc]](https://godoc.org/github.com/pascaldekloe/metrics)
* [shellwords](https://github.com/Wing924/shellwords) **star:12** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/shellwords)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/shellwords)
* [anagent](https://github.com/mudler/anagent) **star:11** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mudler/anagent)   [![godoc][GoDoc]](https://godoc.org/github.com/mudler/anagent)
* [avgRating](https://github.com/kirillDanshin/avgRating) **star:10** Calculate average score and rating based on Wilson Score Equation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/avgRating)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/avgRating)
* [hostutils](https://github.com/Wing924/hostutils) **star:9** A golang library for packing and unpacking FQDNs list.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/hostutils)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/hostutils)
* [numa](https://github.com/lrita/numa) **star:4** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][GoDoc]](https://godoc.org/github.com/lrita/numa)

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) **star:2669** Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only.   [![star > 2000][Awesome]](https://github.com/jdkato/prose)   [![There was an update last week][Green]](https://github.com/jdkato/prose)   [![godoc][GoDoc]](https://godoc.org/github.com/jdkato/prose)
* [gse](https://github.com/go-ego/gse) **star:1513** Go efficient text segmentation; support english, chinese, japanese and other.   [![There was an update last week][Green]](https://github.com/go-ego/gse)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ego/gse)   [![Contains Chinese documents][CN]](https://github.com/go-ego/gse)
* [gojieba](https://github.com/yanyiwu/gojieba) **star:1384** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/yanyiwu/gojieba)   [![Contains Chinese documents][CN]](https://github.com/yanyiwu/gojieba)
* [when](https://github.com/olebedev/when) **star:1109** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][GoDoc]](https://godoc.org/github.com/olebedev/when)
* [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:886** CN Hanzi to Hanyu Pinyin converter.   [![There was an update last week][Green]](https://github.com/mozillazg/go-pinyin)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-pinyin)
* [spaGO](https://github.com/nlpodyssey/spago) **star:813** Self-contained Machine Learning and Natural Language Processing library in Go.   [![There was an update last week][Green]](https://github.com/nlpodyssey/spago)   [![godoc][GoDoc]](https://godoc.org/github.com/nlpodyssey/spago)
* [kagome](https://github.com/ikawaha/kagome) **star:548** JP morphological analyzer written in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ikawaha/kagome)
* [whatlanggo](https://github.com/abadojack/whatlanggo) **star:473** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![godoc][GoDoc]](https://godoc.org/github.com/abadojack/whatlanggo)
* [nlp](https://github.com/Shixzie/nlp) **star:368** Extract values from strings and fill your structs with nlp.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Shixzie/nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/Shixzie/nlp)
* [nlp](https://github.com/james-bowman/nlp) **star:298** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/nlp)
* [sentences](https://github.com/neurosnap/sentences) **star:292** Sentence tokenizer:  converts text into a list of sentences.   [![It hasn't been updated in the last year][Yellow]](https://github.com/neurosnap/sentences)   [![godoc][GoDoc]](https://godoc.org/github.com/neurosnap/sentences)
* [getlang](https://github.com/rylans/getlang) **star:108** Fast natural language detection package.   [![godoc][GoDoc]](https://godoc.org/github.com/rylans/getlang)
* [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
* [go-nlp](https://github.com/nuance/go-nlp) **star:88** Utilities for working with discrete probability distributions and other tools useful for doing NLP work.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nuance/go-nlp)   [![godoc][GoDoc]](https://godoc.org/github.com/nuance/go-nlp)   [![Archived][Archived]](https://github.com/nuance/go-nlp)
* [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:82** ASCII transliterations of Unicode text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/go-unidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-unidecode)
* [RAKE.go](https://github.com/afjoseph/RAKE.Go) **star:75** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][GoDoc]](https://godoc.org/github.com/afjoseph/RAKE.Go)
* [gounidecode](https://github.com/fiam/gounidecode) **star:72** Unicode transliterator (also known as unidecode) for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fiam/gounidecode)   [![godoc][GoDoc]](https://godoc.org/github.com/fiam/gounidecode)
* [textcat](https://github.com/pebbe/textcat) **star:64** Go package for n-gram based text categorization, with support for utf-8 and raw text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pebbe/textcat)   [![godoc][GoDoc]](https://godoc.org/github.com/pebbe/textcat)
* [go-stem](https://github.com/agonopol/go-stem) **star:60** Implementation of the porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/agonopol/go-stem)   [![godoc][GoDoc]](https://godoc.org/github.com/agonopol/go-stem)
* [MMSEGO](https://github.com/awsong/MMSEGO) **star:59** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/awsong/MMSEGO)   [![godoc][GoDoc]](https://godoc.org/github.com/awsong/MMSEGO)
* [segment](https://github.com/blevesearch/segment) **star:56** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/)   [![godoc][GoDoc]](https://godoc.org/github.com/blevesearch/segment)
* [stemmer](https://github.com/dchest/stemmer) **star:49** Stemmer packages for Go programming language. Includes English and German stemmers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dchest/stemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/dchest/stemmer)
* [go2vec](https://github.com/danieldk/go2vec) **star:41** Reader and utility functions for word2vec embeddings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/danieldk/go2vec)   [![godoc][GoDoc]](https://godoc.org/github.com/danieldk/go2vec)
* [porter2](https://github.com/zhenjl/porter2) **star:39** Really fast Porter 2 stemmer.   [![godoc][GoDoc]](https://godoc.org/github.com/zhenjl/porter2)
* [petrovich](https://github.com/striker2000/petrovich) **star:33** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][GoDoc]](https://godoc.org/github.com/striker2000/petrovich)
* [snowball](https://github.com/goodsign/snowball) **star:28** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/snowball)
* [paicehusk](https://github.com/rookii/paicehusk) **star:26** Golang implementation of the Paice/Husk Stemming Algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rookii/paicehusk)   [![godoc][GoDoc]](https://godoc.org/github.com/rookii/paicehusk)
* [address](https://github.com/bojanz/address) **star:25** Handles address representation, validation and formatting.   [![godoc][GoDoc]](https://godoc.org/github.com/bojanz/address)
* [go-mystem](https://github.com/dveselov/mystem) **star:25** CGo bindings to Yandex.Mystem - russian morphology analyzer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dveselov/mystem)   [![godoc][GoDoc]](https://godoc.org/github.com/dveselov/mystem)
* [icu](https://github.com/goodsign/icu) **star:19** Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/icu)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/icu)
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:18** Go bindings for the snowball libstemmer library including porter 2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][GoDoc]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
* [iuliia-go](https://github.com/mehanizm/iuliia-go) **star:16** Transliterate Cyrillic → Latin in every possible way.   [![godoc][GoDoc]](https://godoc.org/github.com/mehanizm/iuliia-go)
* [go-localize](https://github.com/m1/go-localize) **star:12** Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-localize)
* [shamoji](https://github.com/osamingo/shamoji) **star:12** The shamoji is word filtering package written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/shamoji)
* [libtextcat](https://github.com/goodsign/libtextcat) **star:10** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goodsign/libtextcat)   [![godoc][GoDoc]](https://godoc.org/github.com/goodsign/libtextcat)
* [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:9** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![It hasn't been updated in the last year][Yellow]](https://github.com/xujiajun/gotokenizer)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gotokenizer)
* [govader](https://github.com/jonreiter/govader) **star:9** Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment).   [![godoc][GoDoc]](https://godoc.org/github.com/jonreiter/govader)
* [porter](https://github.com/a2800276/porter) **star:8** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/a2800276/porter)   [![godoc][GoDoc]](https://godoc.org/github.com/a2800276/porter)
* [transliterator](https://github.com/alexsergivan/transliterator) **star:8** Provides one-way string transliteration with supporting of language-specific transliteration rules.   [![godoc][GoDoc]](https://godoc.org/github.com/alexsergivan/transliterator)
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:7** Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.   [![godoc][GoDoc]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)
* [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) **star:4** Sentiment analyzer using sentiwordnet lexicon in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/dinopuguh/gosentiwordnet)

## Networking

*Libraries for working with various layers of the network.*

* [fasthttp](https://github.com/valyala/fasthttp) **star:14318** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![star > 2000][Awesome]](https://github.com/valyala/fasthttp)   [![There was an update last week][Green]](https://github.com/valyala/fasthttp)   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasthttp)
* [kcptun](https://github.com/xtaci/kcptun) **star:12267** Extremely simple & fast udp tunnel based on KCP protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcptun)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcptun)
* [webrtc](https://github.com/pions/webrtc) **star:6347** A pure Go implementation of the WebRTC API.   [![star > 2000][Awesome]](https://github.com/pions/webrtc)   [![There was an update last week][Green]](https://github.com/pions/webrtc)   [![godoc][GoDoc]](https://godoc.org/github.com/pions/webrtc)
* [dns](https://github.com/miekg/dns) **star:5137** Go library for working with DNS.   [![star > 2000][Awesome]](https://github.com/miekg/dns)   [![There was an update last week][Green]](https://github.com/miekg/dns)   [![godoc][GoDoc]](https://godoc.org/github.com/miekg/dns)
* [quic-go](https://github.com/lucas-clemente/quic-go) **star:4834** An implementation of the QUIC protocol in pure Go.   [![star > 2000][Awesome]](https://github.com/lucas-clemente/quic-go)   [![There was an update last week][Green]](https://github.com/lucas-clemente/quic-go)   [![godoc][GoDoc]](https://godoc.org/github.com/lucas-clemente/quic-go)
* [gopacket](https://github.com/google/gopacket) **star:3885** Go library for packet processing with libpcap bindings.   [![star > 2000][Awesome]](https://github.com/google/gopacket)   [![There was an update last week][Green]](https://github.com/google/gopacket)   [![godoc][GoDoc]](https://godoc.org/github.com/google/gopacket)
* [HTTPLab](https://github.com/gchaincl/httplab) **star:3646** HTTPLabs let you inspect HTTP requests and forge responses.   [![star > 2000][Awesome]](https://github.com/gchaincl/httplab)   [![It hasn't been updated in the last year][Yellow]](https://github.com/gchaincl/httplab)   [![godoc][GoDoc]](https://godoc.org/github.com/gchaincl/httplab)
* [gnet](https://github.com/panjf2000/gnet) **star:3553** `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.   [![star > 2000][Awesome]](https://github.com/panjf2000/gnet)   [![There was an update last week][Green]](https://github.com/panjf2000/gnet)   [![godoc][GoDoc]](https://godoc.org/github.com/panjf2000/gnet)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/gnet)
* [kcp-go](https://github.com/xtaci/kcp-go) **star:2849** KCP - Fast and Reliable ARQ Protocol.   [![star > 2000][Awesome]](https://github.com/xtaci/kcp-go)   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/kcp-go)
* [gobgp](https://github.com/osrg/gobgp) **star:2125** BGP implemented in the Go Programming Language.   [![star > 2000][Awesome]](https://github.com/osrg/gobgp)   [![There was an update last week][Green]](https://github.com/osrg/gobgp)   [![godoc][GoDoc]](https://godoc.org/github.com/osrg/gobgp)
* [ssh](https://github.com/gliderlabs/ssh) **star:1836** Higher-level API for building SSH servers (wraps crypto/ssh).   [![godoc][GoDoc]](https://godoc.org/github.com/gliderlabs/ssh)
* [fortio](https://github.com/fortio/fortio) **star:1724** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![There was an update last week][Green]](https://github.com/fortio/fortio)   [![godoc][GoDoc]](https://godoc.org/github.com/fortio/fortio)
* [water](https://github.com/songgao/water) **star:1166** Simple TUN/TAP library.   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/water)
* [go-getter](https://github.com/hashicorp/go-getter) **star:1079** Go library for downloading files or directories from various sources using a URL.   [![There was an update last week][Green]](https://github.com/hashicorp/go-getter)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/go-getter)
* [NFF-Go](https://github.com/intel-go/nff-go) **star:1026** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   [![godoc][GoDoc]](https://godoc.org/github.com/intel-go/nff-go)
* [gev](https://github.com/Allenxuxu/gev) **star:1023** gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.   [![godoc][GoDoc]](https://godoc.org/github.com/Allenxuxu/gev)
* [sftp](https://github.com/pkg/sftp) **star:951** Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt.   [![There was an update last week][Green]](https://github.com/pkg/sftp)   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/sftp)
* [grab](https://github.com/cavaliercoder/grab) **star:793** Go package for managing file downloads.   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/grab)
* [ftp](https://github.com/jlaffaye/ftp) **star:772** Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959).   [![godoc][GoDoc]](https://godoc.org/github.com/jlaffaye/ftp)
* [mdns](https://github.com/hashicorp/mdns) **star:731** Simple mDNS (Multicast DNS) client/server library in Golang.   [![There was an update last week][Green]](https://github.com/hashicorp/mdns)   [![godoc][GoDoc]](https://godoc.org/github.com/hashicorp/mdns)
* [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
* [vssh](https://github.com/yahoo/vssh) **star:699** Go library for building network and server automation over SSH protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/yahoo/vssh)
* [gosnmp](https://github.com/soniah/gosnmp) **star:671** Native Go library for performing SNMP actions.   [![There was an update last week][Green]](https://github.com/soniah/gosnmp)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/gosnmp)
* [lhttp](https://github.com/fanux/lhttp) **star:600** Powerful websocket framework, build your IM server more easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fanux/lhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/fanux/lhttp)   [![Contains Chinese documents][CN]](https://github.com/fanux/lhttp)
* [cidranger](https://github.com/yl2chen/cidranger) **star:566** Fast IP to CIDR lookup for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/yl2chen/cidranger)
* [gotcp](https://github.com/gansidui/gotcp) **star:478** Go package for quickly writing tcp applications.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gansidui/gotcp)   [![godoc][GoDoc]](https://godoc.org/github.com/gansidui/gotcp)
* [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:460** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/peerdiscovery)
* [stun](https://github.com/go-rtc/stun) **star:439** Go implementation of RFC 5389 STUN protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rtc/stun)
* [gopcap](https://github.com/akrennmair/gopcap) **star:407** Go wrapper for libpcap.   [![godoc][GoDoc]](https://godoc.org/github.com/akrennmair/gopcap)
* [go-stun](https://github.com/ccding/go-stun) **star:405** Go implementation of the STUN client (RFC 3489 and RFC 5389).   [![godoc][GoDoc]](https://godoc.org/github.com/ccding/go-stun)
* [raw](https://github.com/mdlayher/raw) **star:379** Package raw enables reading and writing data at the device driver level for a network interface.   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/raw)
* [tcp_server](https://github.com/firstrow/tcp_server) **star:350** Go library for building tcp servers faster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/tcp_server)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/tcp_server)
* [winrm](https://github.com/masterzen/winrm) **star:290** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm)
* [gaio](https://github.com/xtaci/gaio) **star:255** High performance async-io networking for Golang in proactor mode.   [![godoc][GoDoc]](https://godoc.org/github.com/xtaci/gaio)
* [gmqtt](https://github.com/DrmagicE/gmqtt) **star:248** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![There was an update last week][Green]](https://github.com/DrmagicE/gmqtt)   [![godoc][GoDoc]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![Contains Chinese documents][CN]](https://github.com/DrmagicE/gmqtt)
* [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:242** Streaming protocolbuffer data over TCP made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/stabbycutyou/buffstreams)
* [arp](https://github.com/mdlayher/arp) **star:238** Package arp implements the ARP protocol, as described in RFC 826.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/arp)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/arp)
* [ethernet](https://github.com/mdlayher/ethernet) **star:212** Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/ethernet)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/ethernet)
* [gNxI](https://github.com/google/gnxi) **star:170** A collection of tools for Network Management that use the gNMI and gNOI protocols.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gnxi)
* [jazigo](https://github.com/udhos/jazigo) **star:163** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   [![It hasn't been updated in the last year][Yellow]](https://github.com/udhos/jazigo)   [![godoc][GoDoc]](https://godoc.org/github.com/udhos/jazigo)
* [utp](https://github.com/anacrolix/utp) **star:154** Go uTP micro transport protocol implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/anacrolix/utp)   [![godoc][GoDoc]](https://godoc.org/github.com/anacrolix/utp)
* [canopus](https://github.com/zubairhamed/canopus) **star:142** CoAP Client/Server implementation (RFC 7252).   [![It hasn't been updated in the last year][Yellow]](https://github.com/zubairhamed/canopus)   [![godoc][GoDoc]](https://godoc.org/github.com/zubairhamed/canopus)
* [sslb](https://github.com/eduardonunesp/sslb) **star:131** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eduardonunesp/sslb)   [![godoc][GoDoc]](https://godoc.org/github.com/eduardonunesp/sslb)
* [xtcp](https://github.com/xfxdev/xtcp) **star:115** TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol.   [![godoc][GoDoc]](https://godoc.org/github.com/xfxdev/xtcp)
* [dhcp6](https://github.com/mdlayher/dhcp6) **star:68** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mdlayher/dhcp6)   [![godoc][GoDoc]](https://godoc.org/github.com/mdlayher/dhcp6)
* [ether](https://github.com/songgao/ether) **star:68** Cross-platform Go package for sending and receiving ethernet frames.   [![It hasn't been updated in the last year][Yellow]](https://github.com/songgao/ether)   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/ether)
* [packet](https://github.com/aerogo/packet) **star:55** Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aerogo/packet)   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/packet)
* [linkio](https://github.com/ian-kent/linkio) **star:50** Network link speed simulation for Reader/Writer interfaces.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ian-kent/linkio)   [![godoc][GoDoc]](https://godoc.org/github.com/ian-kent/linkio)
* [portproxy](https://github.com/aybabtme/portproxy) **star:43** Simple TCP proxy which adds CORS support to API's which don't support it.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aybabtme/portproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/aybabtme/portproxy)
* [iplib](https://github.com/c-robinson/iplib) **star:42** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![There was an update last week][Green]](https://github.com/c-robinson/iplib)   [![godoc][GoDoc]](https://godoc.org/github.com/c-robinson/iplib)
* [graval](https://github.com/koofr/graval) **star:26** Experimental FTP server framework.   [![godoc][GoDoc]](https://godoc.org/github.com/koofr/graval)
* [go-powerdns](https://github.com/joeig/go-powerdns) **star:23** PowerDNS API bindings for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/joeig/go-powerdns)
* [publicip](https://github.com/polera/publicip) **star:21** Package publicip returns your public facing IPv4 address (internet egress).   [![It hasn't been updated in the last year][Yellow]](https://github.com/polera/publicip)   [![godoc][GoDoc]](https://godoc.org/github.com/polera/publicip)
* [golibwireshark](https://github.com/sunwxg/golibwireshark) **star:19** Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/golibwireshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/golibwireshark)
* [panoptes-stream](https://github.com/yahoo/panoptes-stream) **star:17** A cloud native distributed streaming network telemetry (gNMI, Juniper JTI and Cisco MDT).   [![There was an update last week][Green]](https://github.com/yahoo/panoptes-stream)   [![godoc][GoDoc]](https://godoc.org/github.com/yahoo/panoptes-stream)
* [gohooks](https://github.com/averageflow/gohooks) **star:11** GoHooks make it easy to send and consume secured web-hooks from a Go application. Inspired by Spatie's Laravel Webhook Client and Server.   [![godoc][GoDoc]](https://godoc.org/github.com/averageflow/gohooks)
* [llb](https://github.com/kirillDanshin/llb) **star:11** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/llb)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/llb)
* [goshark](https://github.com/sunwxg/goshark) **star:10** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sunwxg/goshark)   [![godoc][GoDoc]](https://godoc.org/github.com/sunwxg/goshark)
* [tspool](https://github.com/two/tspool) **star:8** A TCP Library use worker pool to improve performance and protect your server.   [![It hasn't been updated in the last year][Yellow]](https://github.com/two/tspool)   [![godoc][GoDoc]](https://godoc.org/github.com/two/tspool)
* [httpproxy](https://github.com/wzshiming/httpproxy) **star:7** HTTP proxy handler and dialer.   [![godoc][GoDoc]](https://godoc.org/github.com/wzshiming/httpproxy)

### HTTP Clients

*Libraries for making HTTP requests.*

* [resty](https://github.com/go-resty/resty) **star:3794** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![star > 2000][Awesome]](https://github.com/go-resty/resty)   [![There was an update last week][Green]](https://github.com/go-resty/resty)   [![godoc][GoDoc]](https://godoc.org/github.com/go-resty/resty)
* [heimdall](https://github.com/gojektech/heimdall) **star:1738** An enchanced http client with retry and hystrix capabilities.   [![godoc][GoDoc]](https://godoc.org/github.com/gojektech/heimdall)
* [grequests](https://github.com/levigross/grequests) **star:1725** A Go "clone" of the great and famous Requests library.   [![godoc][GoDoc]](https://godoc.org/github.com/levigross/grequests)
* [sling](https://github.com/dghubble/sling) **star:1233** Sling is a Go HTTP client library for creating and sending API requests.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/sling)
* [gentleman](https://github.com/h2non/gentleman) **star:864** Full-featured plugin-driven HTTP client library.   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gentleman)
* [pester](https://github.com/sethgrid/pester) **star:535** Go HTTP client calls with retries, backoff, and concurrency.   [![godoc][GoDoc]](https://godoc.org/github.com/sethgrid/pester)
* [request](https://github.com/monaco-io/request) **star:90** HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency.   [![There was an update last week][Green]](https://github.com/monaco-io/request)   [![godoc][GoDoc]](https://godoc.org/github.com/monaco-io/request)
* [rq](https://github.com/ddo/rq) **star:37** A nicer interface for golang stdlib HTTP client.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ddo/rq)   [![godoc][GoDoc]](https://godoc.org/github.com/ddo/rq)
* [go-http-client](https://github.com/bozd4g/go-http-client) **star:16** Make http calls simply and easily.   [![godoc][GoDoc]](https://godoc.org/github.com/bozd4g/go-http-client)
* [httpretry](https://github.com/ybbus/httpretry) **star:11** Enriches the default go HTTP client with retry functionality.   [![godoc][GoDoc]](https://godoc.org/github.com/ybbus/httpretry)

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) **star:1043** Go bindings for GLFW 3.
* [gl](https://github.com/go-gl/gl) **star:776** Go bindings for OpenGL (generated via glow).   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-gl/gl)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/gl)
* [mathgl](https://github.com/go-gl/mathgl) **star:357** Pure Go math package specialized for 3D math, with inspiration from GLM.   [![godoc][GoDoc]](https://godoc.org/github.com/go-gl/mathgl)
* [goxjs/gl](https://github.com/goxjs/gl) **star:146** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/gl)
* [goxjs/glfw](https://github.com/goxjs/glfw) **star:68** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![godoc][GoDoc]](https://godoc.org/github.com/goxjs/glfw)
* [go-glmatrix](https://github.com/technohippy/go-glmatrix) **star:1** Go port of [glMatrix](http://glmatrix.net/) library.   [![godoc][GoDoc]](https://godoc.org/github.com/technohippy/go-glmatrix)

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/go-gorm/gorm) **star:22517** The fantastic ORM library for Golang, aims to be developer friendly.   [![star > 2000][Awesome]](https://github.com/go-gorm/gorm)   [![There was an update last week][Green]](https://github.com/go-gorm/gorm)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorm/gorm)
* [ent](https://github.com/facebook/ent) **star:6412** An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.   [![star > 2000][Awesome]](https://github.com/facebook/ent)   [![There was an update last week][Green]](https://github.com/facebook/ent)   [![godoc][GoDoc]](https://godoc.org/github.com/facebook/ent)
* [go-pg](https://github.com/go-pg/pg) **star:4334** PostgreSQL ORM with focus on PostgreSQL specific features and performance.   [![star > 2000][Awesome]](https://github.com/go-pg/pg)   [![There was an update last week][Green]](https://github.com/go-pg/pg)   [![godoc][GoDoc]](https://godoc.org/github.com/go-pg/pg)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:3613** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   [![star > 2000][Awesome]](https://github.com/volatiletech/sqlboiler)   [![There was an update last week][Green]](https://github.com/volatiletech/sqlboiler)   [![godoc][GoDoc]](https://godoc.org/github.com/volatiletech/sqlboiler)
* [gorp](https://github.com/go-gorp/gorp) **star:3465** Go Relational Persistence, ORM-ish library for Go.   [![star > 2000][Awesome]](https://github.com/go-gorp/gorp)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gorp/gorp)
* [upper.io/db](https://github.com/upper/db) **star:2449** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![star > 2000][Awesome]](https://github.com/upper/db)   [![godoc][GoDoc]](https://godoc.org/github.com/upper/db)
* [XORM](https://gitea.com/xorm/xorm)  Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).
* [reform](https://github.com/go-reform/reform) **star:1060** Better ORM for Go, based on non-empty interfaces and code generation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-reform/reform)
* [pop/soda](https://github.com/gobuffalo/pop) **star:1012** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   [![There was an update last week][Green]](https://github.com/gobuffalo/pop)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/pop)
* [gormt](https://github.com/xxjwxc/gormt) **star:964** Mysql database to golang gorm struct.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/gormt)
* [go-queryset](https://github.com/jirfag/go-queryset) **star:580** 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM.   [![godoc][GoDoc]](https://godoc.org/github.com/jirfag/go-queryset)
* [QBS](https://github.com/coocood/qbs) **star:553** Stands for Query By Struct. A Go ORM.   [![It hasn't been updated in the last year][Yellow]](https://github.com/coocood/qbs)   [![godoc][GoDoc]](https://godoc.org/github.com/coocood/qbs)   [![Contains Chinese documents][CN]](https://github.com/coocood/qbs)
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:514** A flexible and powerful SQL string builder library plus a zero-config ORM.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/go-sqlbuilder)
* [beego orm](https://github.com/astaxie/beego/tree/master/orm)  Powerful orm framework for go. Support: pq/mysql/sqlite3.
* [rel](https://github.com/go-rel/rel) **star:300** Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API.   [![godoc][GoDoc]](https://godoc.org/github.com/go-rel/rel)
* [Zoom](https://github.com/albrow/zoom) **star:271** Blazing-fast datastore and querying engine built on Redis.   [![godoc][GoDoc]](https://godoc.org/github.com/albrow/zoom)
* [grimoire](https://github.com/Fs02/grimoire) **star:149** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   [![godoc][GoDoc]](https://godoc.org/github.com/Fs02/grimoire)
* [go-sql](https://github.com/rushteam/gosql) **star:136** A easy ORM for mysql.   [![godoc][GoDoc]](https://godoc.org/github.com/rushteam/gosql)
* [go-store](https://github.com/gosuri/go-store) **star:100** Simple and fast Redis backed key-value store library for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gosuri/go-store)   [![godoc][GoDoc]](https://godoc.org/github.com/gosuri/go-store)
* [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:22** A simple ORM for Google/Firebase Cloud Firestore.   [![godoc][GoDoc]](https://godoc.org/github.com/jschoedt/go-firestorm)
* [lore](https://github.com/abrahambotros/lore) **star:6** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abrahambotros/lore)   [![godoc][GoDoc]](https://godoc.org/github.com/abrahambotros/lore)
* [marlow](https://github.com/marlow/marlow) **star:4** Generated ORM from project structs for compile time safety assurances.   [![godoc][GoDoc]](https://godoc.org/github.com/marlow/marlow)

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) **star:13226** Go dependency tool.   [![star > 2000][Awesome]](https://github.com/golang/dep)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/dep)   [![Archived][Archived]](https://github.com/golang/dep)
* [vgo](https://go.googlesource.com/vgo/)  Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) **star:8075** Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.   [![star > 2000][Awesome]](https://github.com/Masterminds/glide)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/glide)
* [godep](https://github.com/tools/godep) **star:5628** dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.   [![star > 2000][Awesome]](https://github.com/tools/godep)   [![It hasn't been updated in the last year][Yellow]](https://github.com/tools/godep)   [![godoc][GoDoc]](https://godoc.org/github.com/tools/godep)   [![Archived][Archived]](https://github.com/tools/godep)
* [govendor](https://github.com/kardianos/govendor) **star:5010** Go Package Manager. Go vendor tool that works with the standard vendor file.   [![star > 2000][Awesome]](https://github.com/kardianos/govendor)   [![godoc][GoDoc]](https://godoc.org/github.com/kardianos/govendor)   [![Archived][Archived]](https://github.com/kardianos/govendor)
* [gopm](https://github.com/gpmgo/gopm) **star:2499** Go Package Manager.   [![star > 2000][Awesome]](https://github.com/gpmgo/gopm)   [![It hasn't been updated in the last year][Yellow]](https://github.com/gpmgo/gopm)   [![godoc][GoDoc]](https://godoc.org/github.com/gpmgo/gopm)   [![Archived][Archived]](https://github.com/gpmgo/gopm)
* [gom](https://github.com/mattn/gom) **star:1390** Go Manager - bundle for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattn/gom)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/gom)
* [gpm](https://github.com/pote/gpm) **star:1201** Barebones dependency manager for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pote/gpm)
* [goop](https://github.com/nitrous-io/goop) **star:780** Simple dependency manager for Go (golang), inspired by Bundler.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nitrous-io/goop)   [![godoc][GoDoc]](https://godoc.org/github.com/nitrous-io/goop)
* [modgv](https://github.com/lucasepe/modgv) **star:356** Converts 'go mod graph' output into Graphviz's DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/lucasepe/modgv)
* [nut](https://github.com/jingweno/nut) **star:242** Vendor Go dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jingweno/nut)   [![godoc][GoDoc]](https://godoc.org/github.com/jingweno/nut)
* [johnny-deps](https://github.com/VividCortex/johnny-deps) **star:215** Minimal dependency version using Git.
* [VenGO](https://github.com/DamnWidget/VenGO) **star:121** create and manage exportable isolated go virtual environments.   [![It hasn't been updated in the last year][Yellow]](https://github.com/DamnWidget/VenGO)   [![godoc][GoDoc]](https://godoc.org/github.com/DamnWidget/VenGO)
* [mvn-golang](https://github.com/raydac/mvn-golang) **star:114** plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.
* [gop](https://github.com/lunny/gop) **star:49** Build and manage your Go applications out of GOPATH.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lunny/gop)   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/gop)   [![Contains Chinese documents][CN]](https://github.com/lunny/gop)   [![Archived][Archived]](https://github.com/lunny/gop)

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) **star:12693** A distributed tracing system.   [![star > 2000][Awesome]](https://github.com/jaegertracing/jaeger)   [![There was an update last week][Green]](https://github.com/jaegertracing/jaeger)   [![godoc][GoDoc]](https://godoc.org/github.com/jaegertracing/jaeger)
* [profile](https://github.com/pkg/profile) **star:1376** Simple profiling support package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/pkg/profile)
* [statsviz](https://github.com/arl/statsviz) **star:982** Live visualization of your Go application runtime statistics.   [![There was an update last week][Green]](https://github.com/arl/statsviz)
* [pixie](https://github.com/pixie-labs/pixie) **star:519** No instrumentation tracing for Golang applications via eBPF.
* [tracer](https://github.com/kamilsk/tracer) **star:41** Simple, lightweight tracing.   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/tracer)

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) **star:7278** Implementation of GraphQL for Go.   [![star > 2000][Awesome]](https://github.com/graphql-go/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/graphql-go/graphql)
* [graphql](https://github.com/neelance/graphql-go) **star:3588** GraphQL server with a focus on ease of use.   [![star > 2000][Awesome]](https://github.com/neelance/graphql-go)   [![There was an update last week][Green]](https://github.com/neelance/graphql-go)   [![godoc][GoDoc]](https://godoc.org/github.com/neelance/graphql-go)
* [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:1610** A simple Go package to Query over JSON Data.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/gojsonq)
* [dasel](https://github.com/tomwright/dasel) **star:699** Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/tomwright/dasel)
* [jsonql](https://github.com/elgs/jsonql) **star:242** JSON query expression library in Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/elgs/jsonql)
* [rql](https://github.com/a8m/rql) **star:183** Resource Query Language for REST API.   [![godoc][GoDoc]](https://godoc.org/github.com/a8m/rql)
* [graphql](https://github.com/tmc/graphql) **star:52** graphql parser + utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tmc/graphql)   [![godoc][GoDoc]](https://godoc.org/github.com/tmc/graphql)
* [jsonslice](https://github.com/bhmj/jsonslice) **star:49** Jsonpath queries with advanced filters.   [![godoc][GoDoc]](https://godoc.org/github.com/bhmj/jsonslice)
* [api-fu](https://github.com/ccbrown/api-fu) **star:23** Comprehensive GraphQL implementation.   [![There was an update last week][Green]](https://github.com/ccbrown/api-fu)   [![godoc][GoDoc]](https://godoc.org/github.com/ccbrown/api-fu)
* [straf](https://github.com/SonicRoshan/straf) **star:22** Easily Convert Golang structs to GraphQL objects.   [![godoc][GoDoc]](https://godoc.org/github.com/SonicRoshan/straf)
* [rqp](https://github.com/timsolov/rest-query-parser) **star:15** Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query.   [![godoc][GoDoc]](https://godoc.org/github.com/timsolov/rest-query-parser)
* [gws](https://github.com/Zaba505/gws) **star:4** Apollos' "GraphQL over Websocket" client and server implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/Zaba505/gws)
* [jsonpath](https://github.com/AsaiYusuke/jsonpath) **star:2** A query library for retrieving part of JSON based on JSONPath syntax.   [![There was an update last week][Green]](https://github.com/AsaiYusuke/jsonpath)   [![godoc][GoDoc]](https://godoc.org/github.com/AsaiYusuke/jsonpath)

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) **star:3141** The simple and easy way to embed static files into Go binaries.   [![star > 2000][Awesome]](https://github.com/gobuffalo/packr)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/packr)
* [statik](https://github.com/rakyll/statik) **star:3094** Embeds static files into a Go executable.   [![star > 2000][Awesome]](https://github.com/rakyll/statik)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/statik)
* [go.rice](https://github.com/GeertJohan/go.rice) **star:2155** go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy.   [![star > 2000][Awesome]](https://github.com/GeertJohan/go.rice)   [![godoc][GoDoc]](https://godoc.org/github.com/GeertJohan/go.rice)
* [vfsgen](https://github.com/shurcooL/vfsgen) **star:920** Generates a vfsdata.go file that statically implements the given virtual filesystem.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/vfsgen)
* [esc](https://github.com/mjibson/esc) **star:587** Embeds files into Go programs and provides http.FileSystem interfaces to them.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mjibson/esc)   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/esc)
* [fileb0x](https://github.com/UnnoTed/fileb0x) **star:580** Simple tool to embed files in go with focus on "customization" and ease to use.   [![godoc][GoDoc]](https://godoc.org/github.com/UnnoTed/fileb0x)
* [go-resources](https://github.com/omeid/go-resources) **star:172** Unfancy resources embedding with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/omeid/go-resources)
* [statics](https://github.com/go-playground/statics) **star:62** Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/statics)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/statics)
* [go-embed](https://github.com/pyros2097/go-embed) **star:28** Generates go code to embed resource files into your library or executable.   [![godoc][GoDoc]](https://godoc.org/github.com/pyros2097/go-embed)
* [templify](https://github.com/wlbr/templify) **star:25** Embed external template files into Go code to create single file binaries.   [![It hasn't been updated in the last year][Yellow]](https://github.com/wlbr/templify)   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/templify)
* [mule](https://github.com/wlbr/mule) **star:9** Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focussed on simplicity.   [![godoc][GoDoc]](https://godoc.org/github.com/wlbr/mule)

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) **star:4551** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![star > 2000][Awesome]](https://github.com/gonum/gonum)   [![There was an update last week][Green]](https://github.com/gonum/gonum)   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/gonum)
* [stats](https://github.com/montanaflynn/stats) **star:1898** Statistics package with common functions missing from the Golang standard library.   [![godoc][GoDoc]](https://godoc.org/github.com/montanaflynn/stats)
* [gonum/plot](https://github.com/gonum/plot) **star:1806** gonum/plot provides an API for building and drawing plots in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gonum/plot)
* [gosl](https://github.com/cpmech/gosl) **star:1551** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![There was an update last week][Green]](https://github.com/cpmech/gosl)
* [streamtools](https://github.com/nytlabs/streamtools) **star:1314** general purpose, graphical tool for dealing with streams of data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nytlabs/streamtools)   [![godoc][GoDoc]](https://godoc.org/github.com/nytlabs/streamtools)
* [go-dsp](https://github.com/mjibson/go-dsp) **star:707** Digital Signal Processing for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mjibson/go-dsp)   [![godoc][GoDoc]](https://godoc.org/github.com/mjibson/go-dsp)
* [chart](https://github.com/vdobler/chart) **star:665** Simple Chart Plotting library for Go. Supports many graphs types.   [![godoc][GoDoc]](https://godoc.org/github.com/vdobler/chart)
* [goraph](https://github.com/gyuho/goraph) **star:636** Pure Go graph theory library(data structure, algorith visualization).   [![It hasn't been updated in the last year][Yellow]](https://github.com/gyuho/goraph)   [![godoc][GoDoc]](https://godoc.org/github.com/gyuho/goraph)
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:434** Dataframes for machine-learning and statistics (similar to pandas).   [![There was an update last week][Green]](https://github.com/rocketlaunchr/dataframe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
* [graph](https://github.com/yourbasic/graph) **star:412** Library of basic graph algorithms.   [![There was an update last week][Green]](https://github.com/yourbasic/graph)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/graph)
* [orb](https://github.com/paulmach/orb) **star:359** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][GoDoc]](https://godoc.org/github.com/paulmach/orb)
* [ewma](https://github.com/VividCortex/ewma) **star:320** Exponentially-weighted moving averages.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/ewma)
* [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) **star:250** Calendar heatmap in plain Go inspired by Github contribution activity.   [![godoc][GoDoc]](https://godoc.org/github.com/nikolaydubina/calendarheatmap)
* [gohistogram](https://github.com/VividCortex/gohistogram) **star:149** Approximate histograms for data streams.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/gohistogram)
* [TextRank](https://github.com/DavidBelicza/TextRank) **star:121** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.   [![godoc][GoDoc]](https://godoc.org/github.com/DavidBelicza/TextRank)
* [sparse](https://github.com/james-bowman/sparse) **star:108** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/james-bowman/sparse)
* [pagerank](https://github.com/alixaxel/pagerank) **star:67** Weighted PageRank algorithm implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/pagerank)
* [geom](https://github.com/skelterjohn/geom) **star:46** 2D geometry for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/geom)   [![godoc][GoDoc]](https://godoc.org/github.com/skelterjohn/geom)
* [evaler](https://github.com/soniah/evaler) **star:45** Simple floating point arithmetic expression evaluator.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/evaler)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/evaler)
* [goent](https://github.com/kzahedi/goent) **star:21** GO Implementation of Entropy Measures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kzahedi/goent)   [![godoc][GoDoc]](https://godoc.org/github.com/kzahedi/goent)
* [decimal](https://github.com/db47h/decimal) **star:19** Package decimal implements arbitrary-precision decimal floating-point arithmetic.   [![godoc][GoDoc]](https://godoc.org/github.com/db47h/decimal)
* [triangolatte](https://github.com/tchayen/triangolatte) **star:19** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![godoc][GoDoc]](https://godoc.org/github.com/tchayen/triangolatte)
* [GoStats](https://github.com/OGFris/GoStats) **star:14** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/OGFris/GoStats)   [![godoc][GoDoc]](https://godoc.org/github.com/OGFris/GoStats)
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:14** Tiny linear interpolation library.   [![godoc][GoDoc]](https://godoc.org/github.com/sgreben/piecewiselinear)
* [ode](https://github.com/ChristopherRabotin/ode) **star:13** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/ode)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/ode)
* [PiHex](https://github.com/claygod/PiHex) **star:13** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/PiHex)
* [assocentity](https://github.com/ndabAP/assocentity) **star:6** Package assocentity returns the average distance from words to a given entity.   [![godoc][GoDoc]](https://godoc.org/github.com/ndabAP/assocentity)
* [go-gt](https://github.com/ThePaw/go-gt) **star:5** Graph theory algorithms written in "Go" language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ThePaw/go-gt)   [![godoc][GoDoc]](https://godoc.org/github.com/ThePaw/go-gt)
* [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:4** Provides a Bradley-Terry Model for pairwise comparisons.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seanhagen/bradleyterry)   [![godoc][GoDoc]](https://godoc.org/github.com/seanhagen/bradleyterry)
* [rootfinding](https://github.com/khezen/rootfinding) **star:4** root-finding algorithms library for finding roots of quadratic functions.   [![godoc][GoDoc]](https://godoc.org/github.com/khezen/rootfinding)

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/go-acme/lego) **star:4422** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   [![star > 2000][Awesome]](https://github.com/go-acme/lego)   [![There was an update last week][Green]](https://github.com/go-acme/lego)   [![godoc][GoDoc]](https://godoc.org/github.com/go-acme/lego)
* [Cameradar](https://github.com/Ullaakut/cameradar) **star:2455** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![star > 2000][Awesome]](https://github.com/Ullaakut/cameradar)   [![godoc][GoDoc]](https://godoc.org/github.com/Ullaakut/cameradar)
* [memguard](https://github.com/awnumar/memguard) **star:1844** A pure Go library for handling sensitive values in memory.   [![godoc][GoDoc]](https://godoc.org/github.com/awnumar/memguard)
* [acmetool](https://github.com/hlandau/acme) **star:1829** ACME (Let's Encrypt) client tool with automatic renewal.   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/acme)
* [secure](https://github.com/unrolled/secure) **star:1609** HTTP middleware for Go that facilitates some quick security wins.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/secure)
* [themis](https://github.com/cossacklabs/themis) **star:1197** high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps.   [![There was an update last week][Green]](https://github.com/cossacklabs/themis)
* [acra](https://github.com/cossacklabs/acra) **star:700** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   [![There was an update last week][Green]](https://github.com/cossacklabs/acra)   [![godoc][GoDoc]](https://godoc.org/github.com/cossacklabs/acra)
* [nacl](https://github.com/kevinburke/nacl) **star:492** Go implementation of the NaCL set of API's.   [![godoc][GoDoc]](https://godoc.org/github.com/kevinburke/nacl)
* [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) **star:298** A rest application to dynamically update firewalld rules on a linux server.   [![godoc][GoDoc]](https://godoc.org/github.com/prashantgupta24/firewalld-rest)
* [BadActor](https://github.com/jaredfolkins/badactor) **star:291** In-memory, application-driven jailer built in the spirit of fail2ban.   [![godoc][GoDoc]](https://godoc.org/github.com/jaredfolkins/badactor)
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:269** encrypt/decrypt using ssh keys.   [![godoc][GoDoc]](https://godoc.org/github.com/ssh-vault/ssh-vault)
* [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) **star:255** Password validator based on raw cryptographic entropy values.   [![godoc][GoDoc]](https://godoc.org/github.com/lane-c-wagner/go-password-validator)
* [optimus-go](https://github.com/pjebs/optimus-go) **star:253** ID hashing and Obfuscation using Knuth's Algorithm.   [![godoc][GoDoc]](https://godoc.org/github.com/pjebs/optimus-go)
* [passlib](https://github.com/hlandau/passlib) **star:239** Futureproof password hashing library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hlandau/passlib)   [![godoc][GoDoc]](https://godoc.org/github.com/hlandau/passlib)
* [go-yara](https://github.com/hillu/go-yara) **star:196** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   [![godoc][GoDoc]](https://godoc.org/github.com/hillu/go-yara)
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:165** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][GoDoc]](https://godoc.org/github.com/elithrar/simple-scrypt)
* [argon2pw](https://github.com/raja/argon2pw) **star:85** Argon2 password hash generation with constant-time password comparison.   [![It hasn't been updated in the last year][Yellow]](https://github.com/raja/argon2pw)   [![godoc][GoDoc]](https://godoc.org/github.com/raja/argon2pw)
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:39** A probably paranoid package for securely hashing and encrypting passwords.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
* [Interpol](https://bitbucket.org/vahidi/interpol)  Rule-based data generator for fuzzing and penetration testing.
* [certificates](https://github.com/mvmaasakkers/certificates) **star:20** An opinionated tool for generating tls certificates.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/certificates)
* [go-generate-password](https://github.com/m1/go-generate-password) **star:16** Password generator that can be used on the cli or as a library.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/go-generate-password)
* [goArgonPass](https://github.com/dwin/goArgonPass) **star:13** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![godoc][GoDoc]](https://godoc.org/github.com/dwin/goArgonPass)
* [secureio](https://github.com/xaionaro-go/secureio) **star:13** An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519.   [![godoc][GoDoc]](https://godoc.org/github.com/xaionaro-go/secureio)
* [argon2-hashing](https://github.com/andskur/argon2-hashing) **star:12** light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package.   [![godoc][GoDoc]](https://godoc.org/github.com/andskur/argon2-hashing)
* [sslmgr](https://github.com/adrianosela/sslmgr) **star:10** SSL certificates made easy with a high level wrapper around acme/autocert.   [![It hasn't been updated in the last year][Yellow]](https://github.com/adrianosela/sslmgr)   [![godoc][GoDoc]](https://godoc.org/github.com/adrianosela/sslmgr)


## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) **star:8805** High-performance 100% compatible drop-in replacement of "encoding/json".   [![star > 2000][Awesome]](https://github.com/json-iterator/go)   [![There was an update last week][Green]](https://github.com/json-iterator/go)   [![godoc][GoDoc]](https://godoc.org/github.com/json-iterator/go)
* [goprotobuf](https://github.com/golang/protobuf) **star:7392** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![star > 2000][Awesome]](https://github.com/golang/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/protobuf)
* [gogoprotobuf](https://github.com/gogo/protobuf) **star:4318** Protocol Buffers for Go with Gadgets.   [![star > 2000][Awesome]](https://github.com/gogo/protobuf)   [![godoc][GoDoc]](https://godoc.org/github.com/gogo/protobuf)
* [mapstructure](https://github.com/mitchellh/mapstructure) **star:4134** Go library for decoding generic map values into native Go structures.   [![star > 2000][Awesome]](https://github.com/mitchellh/mapstructure)   [![There was an update last week][Green]](https://github.com/mitchellh/mapstructure)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/mapstructure)
* [go-codec](https://github.com/ugorji/go) **star:1497** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   [![There was an update last week][Green]](https://github.com/ugorji/go)   [![godoc][GoDoc]](https://godoc.org/github.com/ugorji/go)
* [colfer](https://github.com/pascaldekloe/colfer) **star:589** Code generation for the Colfer binary format.
* [csvutil](https://github.com/jszwec/csvutil) **star:453** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![godoc][GoDoc]](https://godoc.org/github.com/jszwec/csvutil)
* [go-capnproto](https://github.com/glycerine/go-capnproto) **star:277** Cap'n Proto library and parser for go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/glycerine/go-capnproto)   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/go-capnproto)
* [cbor](https://github.com/fxamacker/cbor) **star:232** Small, safe, and easy CBOR encoding and decoding library.   [![godoc][GoDoc]](https://godoc.org/github.com/fxamacker/cbor)
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:145** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yvasiyarov/php_session_decoder)   [![godoc][GoDoc]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
* [structomap](https://github.com/tuvistavie/structomap) **star:117** Library to easily and dynamically generate maps from static structures.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tuvistavie/structomap)   [![godoc][GoDoc]](https://godoc.org/github.com/tuvistavie/structomap)
* [bambam](https://github.com/glycerine/bambam) **star:61** generator for Cap'n Proto schemas from go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/glycerine/bambam)   [![godoc][GoDoc]](https://godoc.org/github.com/glycerine/bambam)
* [asn1](https://github.com/PromonLogicalis/asn1) **star:47** Asn.1 BER and DER encoding library for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PromonLogicalis/asn1)   [![godoc][GoDoc]](https://godoc.org/github.com/PromonLogicalis/asn1)   [![Archived][Archived]](https://github.com/PromonLogicalis/asn1)
* [binstruct](https://github.com/ghostiam/binstruct) **star:28** Golang binary decoder for mapping data into the structure.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ghostiam/binstruct)   [![godoc][GoDoc]](https://godoc.org/github.com/ghostiam/binstruct)
* [fwencoder](https://github.com/o1egl/fwencoder) **star:14** Fixed width file parser (encoding and decoding library) for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/o1egl/fwencoder)
* [elastic](https://github.com/epiclabs-io/elastic) **star:13** Convert slices, maps or any other unknown value across different types at run-time, no matter what.   [![godoc][GoDoc]](https://godoc.org/github.com/epiclabs-io/elastic)
* [pletter](https://github.com/vimeda/pletter) **star:12** A standard way to wrap a proto message for message brokers.   [![godoc][GoDoc]](https://godoc.org/github.com/vimeda/pletter)
* [bel](https://github.com/32leaves/bel) **star:11** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.   [![godoc][GoDoc]](https://godoc.org/github.com/32leaves/bel)
* [fixedwidth](https://github.com/huydang284/fixedwidth) **star:5** Fixed-width text formatting (UTF-8 supported).   [![It hasn't been updated in the last year][Yellow]](https://github.com/huydang284/fixedwidth)   [![godoc][GoDoc]](https://godoc.org/github.com/huydang284/fixedwidth)
* [go-lctree](https://github.com/sbourlon/go-lctree) **star:2** Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation).   [![godoc][GoDoc]](https://godoc.org/github.com/sbourlon/go-lctree)

## Server Applications

* [etcd](https://github.com/coreos/etcd) **star:34513** Highly-available key value store for shared configuration and service discovery.   [![star > 2000][Awesome]](https://github.com/coreos/etcd)   [![There was an update last week][Green]](https://github.com/coreos/etcd)   [![godoc][GoDoc]](https://godoc.org/github.com/coreos/etcd)
* [Caddy](https://github.com/caddyserver/caddy) **star:31675** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   [![star > 2000][Awesome]](https://github.com/caddyserver/caddy)   [![There was an update last week][Green]](https://github.com/caddyserver/caddy)   [![godoc][GoDoc]](https://godoc.org/github.com/caddyserver/caddy)
* [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
* [minio](https://github.com/minio/minio) **star:25819** Minio is a distributed object storage server.   [![star > 2000][Awesome]](https://github.com/minio/minio)   [![There was an update last week][Green]](https://github.com/minio/minio)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio)
* [RoadRunner](https://github.com/spiral/roadrunner) **star:4781** High-performance PHP application server, load-balancer and process manager.   [![star > 2000][Awesome]](https://github.com/spiral/roadrunner)   [![There was an update last week][Green]](https://github.com/spiral/roadrunner)   [![godoc][GoDoc]](https://godoc.org/github.com/spiral/roadrunner)
* [devd](https://github.com/cortesi/devd) **star:3086** Local webserver for developers.   [![star > 2000][Awesome]](https://github.com/cortesi/devd)   [![godoc][GoDoc]](https://godoc.org/github.com/cortesi/devd)
* [SFTPGo](https://github.com/drakkan/sftpgo) **star:2213** Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage.   [![star > 2000][Awesome]](https://github.com/drakkan/sftpgo)   [![There was an update last week][Green]](https://github.com/drakkan/sftpgo)   [![godoc][GoDoc]](https://godoc.org/github.com/drakkan/sftpgo)
* [algernon](https://github.com/xyproto/algernon) **star:1774** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   [![godoc][GoDoc]](https://godoc.org/github.com/xyproto/algernon)
* [Flagr](https://github.com/checkr/flagr) **star:1460** Flagr is an open-source feature flagging and A/B testing service.   [![There was an update last week][Green]](https://github.com/checkr/flagr)   [![godoc][GoDoc]](https://godoc.org/github.com/checkr/flagr)
* [Fider](https://github.com/getfider/fider) **star:1427** Fider is an open platform to collect and organize customer feedback.   [![godoc][GoDoc]](https://godoc.org/github.com/getfider/fider)
* [flipt](https://github.com/markphelps/flipt) **star:1383** A self contained feature flag solution written in Go and Vue.js   [![There was an update last week][Green]](https://github.com/markphelps/flipt)   [![godoc][GoDoc]](https://godoc.org/github.com/markphelps/flipt)
* [Trickster](https://github.com/tricksterproxy/trickster) **star:1277** HTTP reverse proxy cache and time series accelerator.   [![There was an update last week][Green]](https://github.com/tricksterproxy/trickster)   [![godoc][GoDoc]](https://godoc.org/github.com/tricksterproxy/trickster)
* [discovery](https://github.com/Bilibili/discovery) **star:1266** A registry for resilient mid-tier load balancing and failover.   [![godoc][GoDoc]](https://godoc.org/github.com/Bilibili/discovery)
* [jackal](https://github.com/ortuman/jackal) **star:873** An XMPP server written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ortuman/jackal)
* [dudeldu](https://github.com/krotik/dudeldu) **star:119** A simple SHOUTcast server.   [![It hasn't been updated in the last year][Yellow]](https://github.com/krotik/dudeldu)   [![godoc][GoDoc]](https://godoc.org/github.com/krotik/dudeldu)
* [lets-proxy2](https://github.com/rekby/lets-proxy2) **star:47** Reverse proxy for handle https with issue certificates in fly from lets-encrypt.   [![godoc][GoDoc]](https://godoc.org/github.com/rekby/lets-proxy2)
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:26** Stream database events from PostgreSQL to Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/psql-streamer)
* [riemann-relay](https://github.com/blind-oracle/riemann-relay)  Relay to load-balance Riemann events and/or convert them to Carbon.
* [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) **star:25** A feature flag solution, with only a YAML file in the backend (S3, GitHub, HTTP, local file ...), no server to install, just add a file in a central system and refer to it.   [![There was an update last week][Green]](https://github.com/thomaspoignant/go-feature-flag)   [![godoc][GoDoc]](https://godoc.org/github.com/thomaspoignant/go-feature-flag)
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:16** Nginx log parser and exporter to Prometheus.   [![godoc][GoDoc]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
* [nsq](http://nsq.io/)  A realtime distributed messaging platform.
* [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) **star:11** Simple and lightweight provider which exhibits JWTs, supports login, password-reset (via mail) and user management.   [![There was an update last week][Green]](https://github.com/leberKleber/simple-jwt-provider)   [![godoc][GoDoc]](https://godoc.org/github.com/leberKleber/simple-jwt-provider)
* [protoxy](https://github.com/camgraff/protoxy) **star:10** A proxy server that converts JSON request bodies to Protocol Buffers.   [![godoc][GoDoc]](https://godoc.org/github.com/camgraff/protoxy)
* [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) **star:6** Simple Reverse Proxy with Caching, written in Go, using Redis.   [![There was an update last week][Green]](https://github.com/fabiocicerchia/go-proxy-cache)   [![godoc][GoDoc]](https://godoc.org/github.com/fabiocicerchia/go-proxy-cache)


## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) **star:585** Go stream processing library.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/go-streams)
* [machine](https://github.com/whitaker-io/machine) **star:61** Go library for writing and generating stream workers with built in metrics and traceability.   [![godoc][GoDoc]](https://godoc.org/github.com/whitaker-io/machine)
* [stream](https://github.com/youthlin/stream) **star:22** Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce...   [![godoc][GoDoc]](https://godoc.org/github.com/youthlin/stream)

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) **star:3733** PDF document generator with high level support for text, drawing and images.   [![star > 2000][Awesome]](https://github.com/jung-kurt/gofpdf)   [![It hasn't been updated in the last year][Yellow]](https://github.com/jung-kurt/gofpdf)   [![godoc][GoDoc]](https://godoc.org/github.com/jung-kurt/gofpdf)   [![Archived][Archived]](https://github.com/jung-kurt/gofpdf)
* [sprig](https://github.com/Masterminds/sprig) **star:2121** Useful template functions for Go templates.   [![star > 2000][Awesome]](https://github.com/Masterminds/sprig)   [![godoc][GoDoc]](https://godoc.org/github.com/Masterminds/sprig)
* [quicktemplate](https://github.com/valyala/quicktemplate) **star:1969** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/quicktemplate)
* [pongo2](https://github.com/flosch/pongo2) **star:1870** Django-like template-engine for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/flosch/pongo2)
* [hero](https://github.com/shiyanhui/hero) **star:1430** Hero is a handy, fast and powerful go template engine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shiyanhui/hero)   [![godoc][GoDoc]](https://godoc.org/github.com/shiyanhui/hero)   [![Contains Chinese documents][CN]](https://github.com/shiyanhui/hero)
* [mustache](https://github.com/hoisie/mustache) **star:1003** Go implementation of the Mustache template language.   [![godoc][GoDoc]](https://godoc.org/github.com/hoisie/mustache)
* [amber](https://github.com/eknkc/amber) **star:865** Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade.   [![godoc][GoDoc]](https://godoc.org/github.com/eknkc/amber)
* [ace](https://github.com/yosssi/ace) **star:797** Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yosssi/ace)   [![godoc][GoDoc]](https://godoc.org/github.com/yosssi/ace)
* [Razor](https://github.com/sipin/gorazor) **star:770** Razor view engine for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sipin/gorazor)
* [jet](https://github.com/CloudyKit/jet) **star:744** Jet template engine.   [![godoc][GoDoc]](https://godoc.org/github.com/CloudyKit/jet)
* [ego](https://github.com/benbjohnson/ego) **star:464** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![There was an update last week][Green]](https://github.com/benbjohnson/ego)   [![godoc][GoDoc]](https://godoc.org/github.com/benbjohnson/ego)
* [fasttemplate](https://github.com/valyala/fasttemplate) **star:462** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/).   [![godoc][GoDoc]](https://godoc.org/github.com/valyala/fasttemplate)
* [raymond](https://github.com/aymerick/raymond) **star:411** Complete handlebars implementation in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aymerick/raymond)
* [maroto](https://github.com/johnfercher/maroto) **star:307** A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.   [![There was an update last week][Green]](https://github.com/johnfercher/maroto)   [![godoc][GoDoc]](https://godoc.org/github.com/johnfercher/maroto)
* [goview](https://github.com/foolin/goview) **star:200** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/goview)
* [Soy](https://github.com/robfig/soy) **star:154** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![godoc][GoDoc]](https://godoc.org/github.com/robfig/soy)
* [liquid](https://github.com/osteele/liquid) **star:115** Go implementation of Shopify Liquid templates.   [![godoc][GoDoc]](https://godoc.org/github.com/osteele/liquid)
* [velvet](https://github.com/gobuffalo/velvet) **star:74** Complete handlebars implementation in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gobuffalo/velvet)   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/velvet)
* [kasia.go](https://github.com/ziutek/kasia.go) **star:72** Templating system for HTML and other text documents - go implementation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ziutek/kasia.go)   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/kasia.go)
* [extemplate](https://github.com/dannyvankooten/extemplate) **star:30** Tiny wrapper around html/template to allow for easy file-based template inheritance.   [![godoc][GoDoc]](https://godoc.org/github.com/dannyvankooten/extemplate)
* [damsel](https://github.com/dskinner/damsel) **star:24** Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dskinner/damsel)   [![godoc][GoDoc]](https://godoc.org/github.com/dskinner/damsel)
* [gospin](https://github.com/m1/gospin) **star:21** Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations.   [![godoc][GoDoc]](https://godoc.org/github.com/m1/gospin)

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [Testify](https://github.com/stretchr/testify) **star:12373** Sacred extension to the standard go testing package.   [![star > 2000][Awesome]](https://github.com/stretchr/testify)   [![There was an update last week][Green]](https://github.com/stretchr/testify)   [![godoc][GoDoc]](https://godoc.org/github.com/stretchr/testify)
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd)  Convert markdown snippets into testable go code.
    * [go-cmp](https://github.com/google/go-cmp) **star:2113** Package for comparing Go values in tests.   [![star > 2000][Awesome]](https://github.com/google/go-cmp)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-cmp)
    * [httpexpect](https://github.com/gavv/httpexpect) **star:1588** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   [![There was an update last week][Green]](https://github.com/gavv/httpexpect)   [![godoc][GoDoc]](https://godoc.org/github.com/gavv/httpexpect)
    * [godog](https://github.com/DATA-DOG/godog) **star:1243** Cucumber or Behat like BDD framework for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/godog)
    * [goblin](https://github.com/franela/goblin) **star:747** Mocha like testing framework fo Go.   [![godoc][GoDoc]](https://godoc.org/github.com/franela/goblin)
    * [baloo](https://github.com/h2non/baloo) **star:698** Expressive and versatile end-to-end HTTP API testing made easy.   [![It hasn't been updated in the last year][Yellow]](https://github.com/h2non/baloo)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/baloo)
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:628** A helper for Rails' like test fixtures to test database applications.   [![godoc][GoDoc]](https://godoc.org/github.com/go-testfixtures/testfixtures)
    * [go-vcr](https://github.com/dnaeon/go-vcr) **star:483** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![godoc][GoDoc]](https://godoc.org/github.com/dnaeon/go-vcr)
    * [go-mutesting](https://github.com/zimmski/go-mutesting) **star:379** Mutation testing for Go source code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zimmski/go-mutesting)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/go-mutesting)
    * [gnomock](https://github.com/orlangure/gnomock) **star:370** integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks.   [![godoc][GoDoc]](https://godoc.org/github.com/orlangure/gnomock)
    * [gofight](https://github.com/appleboy/gofight) **star:357** API Handler Testing for Golang Router framework.   [![There was an update last week][Green]](https://github.com/appleboy/gofight)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/gofight)
    * [goc](https://github.com/qiniu/goc) **star:272** Goc is a comprehensive coverage testing system for The Go Programming Language.   [![There was an update last week][Green]](https://github.com/qiniu/goc)   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/goc)
    * [gocheck](http://labix.org/gocheck)  More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
    * [frisby](https://github.com/verdverm/frisby) **star:264** REST API testing framework.   [![godoc][GoDoc]](https://godoc.org/github.com/verdverm/frisby)
    * [ginkgo](http://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
    * [go-carpet](https://github.com/msoap/go-carpet) **star:211** Tool for viewing test coverage in terminal.   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/go-carpet)
    * [charlatan](https://github.com/percolate/charlatan) **star:194** Tool to generate fake interface implementations for tests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/percolate/charlatan)   [![godoc][GoDoc]](https://godoc.org/github.com/percolate/charlatan)
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:189** A collection of packages to augment the go testing package and support common patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/gotestyourself/gotest.tools)
    * [commander](https://github.com/SimonBaeumer/commander) **star:173** Tool for testing cli applications on windows, linux and osx.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/commander)
    * [endly](https://github.com/viant/endly) **star:172** Declarative end to end functional testing.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/endly)
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:127** Simple snapshot testing addon for your test framework.   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
    * [dbcleaner](https://github.com/khaiql/dbcleaner) **star:120** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![godoc][GoDoc]](https://godoc.org/github.com/khaiql/dbcleaner)
    * [go-testdeep](https://github.com/maxatome/go-testdeep) **star:115** Extremely flexible golang deep comparison, extends the go testing package.   [![godoc][GoDoc]](https://godoc.org/github.com/maxatome/go-testdeep)
    * [GoSpec](https://github.com/orfjackal/gospec) **star:115** BDD-style testing framework for the Go programming language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/orfjackal/gospec)   [![godoc][GoDoc]](https://godoc.org/github.com/orfjackal/gospec)
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) **star:88** Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.   [![There was an update last week][Green]](https://github.com/fergusstrange/embedded-postgres)   [![godoc][GoDoc]](https://godoc.org/github.com/fergusstrange/embedded-postgres)
    * [wstest](https://github.com/posener/wstest) **star:82** Websocket client for unit-testing a websocket http.Handler.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/wstest)
    * [gocrest](https://github.com/corbym/gocrest) **star:77** Composable hamcrest-like matchers for Go assertions.   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gocrest)
    * [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [jsonassert](https://github.com/kinbiko/jsonassert) **star:57** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][GoDoc]](https://godoc.org/github.com/kinbiko/jsonassert)
    * [testcase](https://github.com/adamluzsi/testcase) **star:55** Idiomatic testing framework for Behavior Driven Development.   [![There was an update last week][Green]](https://github.com/adamluzsi/testcase)   [![godoc][GoDoc]](https://godoc.org/github.com/adamluzsi/testcase)
    * [gospecify](https://github.com/stesla/gospecify) **star:54** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stesla/gospecify)   [![godoc][GoDoc]](https://godoc.org/github.com/stesla/gospecify)
    * [restit](https://github.com/yookoala/restit) **star:52** Go micro framework to help writing RESTful API integration test.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yookoala/restit)   [![godoc][GoDoc]](https://godoc.org/github.com/yookoala/restit)
    * [gomatch](https://github.com/jfilipczyk/gomatch) **star:40** library created for testing JSON against patterns.   [![godoc][GoDoc]](https://godoc.org/github.com/jfilipczyk/gomatch)
    * [gomega](http://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
    * [dsunit](https://github.com/viant/dsunit) **star:36** Datastore testing for SQL, NoSQL, structured files.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/dsunit)
    * [covergates](https://github.com/covergates/covergates) **star:35** Self-hosted code coverage report review and management service.   [![godoc][GoDoc]](https://godoc.org/github.com/covergates/covergates)
    * [assert](https://github.com/go-playground/assert) **star:28** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/assert)
    * [Hamcrest](https://github.com/rdrdr/hamcrest) **star:27** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   [![godoc][GoDoc]](https://godoc.org/github.com/rdrdr/hamcrest)
    * [go-hit](https://github.com/Eun/go-hit) **star:25** Hit is an http integration test framework written in golang.   [![godoc][GoDoc]](https://godoc.org/github.com/Eun/go-hit)
    * [flute](https://github.com/suzuki-shunsuke/flute) **star:15** HTTP client testing framework.   [![There was an update last week][Green]](https://github.com/suzuki-shunsuke/flute)   [![godoc][GoDoc]](https://godoc.org/github.com/suzuki-shunsuke/flute)
    * [schema](https://github.com/jgroeneveld/schema) **star:14** Quick and easy expression matching for JSON schemas used in requests and responses.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jgroeneveld/schema)   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/schema)
    * [badio](https://github.com/cavaliercoder/badio) **star:10** Extensions to Go's `testing/iotest` package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cavaliercoder/badio)   [![godoc][GoDoc]](https://godoc.org/github.com/cavaliercoder/badio)
    * [biff](https://github.com/fulldump/biff) **star:10** Bifurcation testing framework, BDD compatible.   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/biff)
    * [testsql](https://github.com/zhulongcheng/testsql) **star:10** Generate test data from SQL files before testing and clear it after finished.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhulongcheng/testsql)   [![godoc][GoDoc]](https://godoc.org/github.com/zhulongcheng/testsql)
    * [gogiven](https://github.com/corbym/gogiven) **star:9** YATSPEC-like BDD testing framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/corbym/gogiven)   [![godoc][GoDoc]](https://godoc.org/github.com/corbym/gogiven)
    * [gosuite](https://github.com/pavlo/gosuite) **star:9** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   [![It hasn't been updated in the last year][Yellow]](https://github.com/pavlo/gosuite)   [![godoc][GoDoc]](https://godoc.org/github.com/pavlo/gosuite)
    * [stop-and-go](https://github.com/elgohr/stop-and-go) **star:5** Testing helper for concurrency.   [![There was an update last week][Green]](https://github.com/elgohr/stop-and-go)   [![godoc][GoDoc]](https://godoc.org/github.com/elgohr/stop-and-go)
    * [Tt](https://github.com/vcaesar/tt) **star:5** Simple and colorful test tools.   [![godoc][GoDoc]](https://godoc.org/github.com/vcaesar/tt)
    * [trial](https://github.com/jgroeneveld/trial) **star:4** Quick and easy extendable assertions without introducing much boilerplate.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jgroeneveld/trial)   [![godoc][GoDoc]](https://godoc.org/github.com/jgroeneveld/trial)

* Mock
    * [gomock](https://github.com/golang/mock) **star:5108** Mocking framework for the Go programming language.   [![star > 2000][Awesome]](https://github.com/golang/mock)   [![There was an update last week][Green]](https://github.com/golang/mock)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/mock)
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:3114** Mock SQL driver for testing database interactions.   [![star > 2000][Awesome]](https://github.com/DATA-DOG/go-sqlmock)   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:1688** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   [![There was an update last week][Green]](https://github.com/SpectoLabs/hoverfly)   [![godoc][GoDoc]](https://godoc.org/github.com/SpectoLabs/hoverfly)
    * [gock](https://github.com/h2non/gock) **star:1136** Versatile HTTP mocking made easy.   [![There was an update last week][Green]](https://github.com/h2non/gock)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/gock)
    * [httpmock](https://github.com/jarcoal/httpmock) **star:962** Easy mocking of HTTP responses from external resources.   [![godoc][GoDoc]](https://godoc.org/github.com/jarcoal/httpmock)
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:499** Tool for generating self-contained mock objects.   [![There was an update last week][Green]](https://github.com/maxbrunsfeld/counterfeiter)   [![godoc][GoDoc]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
    * [minimock](https://github.com/gojuno/minimock) **star:336** Mock generator for Go interfaces.   [![There was an update last week][Green]](https://github.com/gojuno/minimock)   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/minimock)
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:330** Single transaction based database driver mainly for testing purposes.   [![godoc][GoDoc]](https://godoc.org/github.com/DATA-DOG/go-txdb)
    * [govcr](https://github.com/seborama/govcr) **star:92** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![It hasn't been updated in the last year][Yellow]](https://github.com/seborama/govcr)   [![godoc][GoDoc]](https://godoc.org/github.com/seborama/govcr)
    * [timex](https://github.com/cabify/timex) **star:52** A test-friendly replacement for the native `time` package.   [![godoc][GoDoc]](https://godoc.org/github.com/cabify/timex)
    * [mockhttp](https://github.com/tv42/mockhttp) **star:22** Mock object for Go http.ResponseWriter.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tv42/mockhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/tv42/mockhttp)
    * [go-localstack](https://github.com/elgohr/go-localstack) **star:7** Tool for using localstack in AWS testing.   [![There was an update last week][Green]](https://github.com/elgohr/go-localstack)   [![godoc][GoDoc]](https://godoc.org/github.com/elgohr/go-localstack)
    * [mockit](https://github.com/pasdam/mockit) **star:5** Allows functions and method easy mocking, without defining new types; it's similar to Mockito for Java.   [![There was an update last week][Green]](https://github.com/pasdam/mockit)   [![godoc][GoDoc]](https://godoc.org/github.com/pasdam/mockit)

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:3799** Randomized testing system.   [![star > 2000][Awesome]](https://github.com/dvyukov/go-fuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/dvyukov/go-fuzz)
    * [gofuzz](https://github.com/google/gofuzz) **star:962** Library for populating go objects with random values.   [![godoc][GoDoc]](https://godoc.org/github.com/google/gofuzz)
    * [Tavor](https://github.com/zimmski/tavor) **star:225** Generic fuzzing and delta-debugging framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zimmski/tavor)   [![godoc][GoDoc]](https://godoc.org/github.com/zimmski/tavor)

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) **star:5821** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![star > 2000][Awesome]](https://github.com/knq/chromedp)   [![There was an update last week][Green]](https://github.com/knq/chromedp)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/chromedp)
    * [selenoid](https://github.com/aerokube/selenoid) **star:1801** alternative Selenium hub server that launches browsers within containers.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/selenoid)
    * [rod](https://github.com/go-rod/rod) **star:1314** A Devtools driver to make web automation and scraping easy.   [![There was an update last week][Green]](https://github.com/go-rod/rod)   [![godoc][GoDoc]](https://godoc.org/github.com/go-rod/rod)
    * [cdp](https://github.com/mafredri/cdp) **star:511** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][GoDoc]](https://godoc.org/github.com/mafredri/cdp)
    * [ggr](https://github.com/aerokube/ggr) **star:266** a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.   [![godoc][GoDoc]](https://godoc.org/github.com/aerokube/ggr)
    * [playwright-go](https://github.com/mxschmitt/playwright-go) **star:237** browser automation library to control Chromium, Firefox and WebKit with a single API.   [![There was an update last week][Green]](https://github.com/mxschmitt/playwright-go)   [![godoc][GoDoc]](https://godoc.org/github.com/mxschmitt/playwright-go)

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) **star:558** An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/pingcap/failpoint)

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) **star:12951** Fast and Elegant Scraping Framework for Gophers.   [![star > 2000][Awesome]](https://github.com/asciimoo/colly)   [![There was an update last week][Green]](https://github.com/asciimoo/colly)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/colly)
    * [GoQuery](https://github.com/PuerkitoBio/goquery) **star:9717** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![star > 2000][Awesome]](https://github.com/PuerkitoBio/goquery)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/goquery)
    * [blackfriday](https://github.com/russross/blackfriday) **star:4626** Markdown processor in Go.   [![star > 2000][Awesome]](https://github.com/russross/blackfriday)   [![godoc][GoDoc]](https://godoc.org/github.com/russross/blackfriday)
    * [sh](https://github.com/mvdan/sh) **star:3433** Shell parser and formatter.   [![star > 2000][Awesome]](https://github.com/mvdan/sh)   [![There was an update last week][Green]](https://github.com/mvdan/sh)   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/sh)
    * [toml](https://github.com/BurntSushi/toml) **star:3405** TOML configuration format (encoder/decoder with reflection).   [![star > 2000][Awesome]](https://github.com/BurntSushi/toml)   [![godoc][GoDoc]](https://godoc.org/github.com/BurntSushi/toml)
    * [go-humanize](https://github.com/dustin/go-humanize) **star:2493** Formatters for time, numbers, and memory size to human readable format.   [![star > 2000][Awesome]](https://github.com/dustin/go-humanize)   [![godoc][GoDoc]](https://godoc.org/github.com/dustin/go-humanize)
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:1759** HTML Sanitizer.   [![godoc][GoDoc]](https://godoc.org/github.com/microcosm-cc/bluemonday)
    * [gofeed](https://github.com/mmcdole/gofeed) **star:1532** Parse RSS and Atom feeds in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/mmcdole/gofeed)
    * [inject](https://github.com/facebookgo/inject) **star:1300** Package inject provides a reflect based injector.   [![It hasn't been updated in the last year][Yellow]](https://github.com/facebookgo/inject)   [![godoc][GoDoc]](https://godoc.org/github.com/facebookgo/inject)   [![Archived][Archived]](https://github.com/facebookgo/inject)
    * [go-toml](https://github.com/pelletier/go-toml) **star:910** Go library for the TOML format with query support and handy cli tools.   [![There was an update last week][Green]](https://github.com/pelletier/go-toml)   [![godoc][GoDoc]](https://godoc.org/github.com/pelletier/go-toml)
    * [commonregex](https://github.com/mingrammer/commonregex) **star:717** A collection of common regular expressions for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mingrammer/commonregex)   [![godoc][GoDoc]](https://godoc.org/github.com/mingrammer/commonregex)
    * [slug](https://github.com/gosimple/slug) **star:623** URL-friendly slugify with multiple languages support.   [![godoc][GoDoc]](https://godoc.org/github.com/gosimple/slug)
    * [dataflowkit](https://github.com/slotix/dataflowkit) **star:447** Web scraping Framework to turn websites into structured data.   [![godoc][GoDoc]](https://godoc.org/github.com/slotix/dataflowkit)
    * [mxj](https://github.com/clbanning/mxj) **star:435** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![There was an update last week][Green]](https://github.com/clbanning/mxj)   [![godoc][GoDoc]](https://godoc.org/github.com/clbanning/mxj)
    * [gographviz](https://github.com/awalterschulze/gographviz) **star:411** Parses the Graphviz DOT language.   [![godoc][GoDoc]](https://godoc.org/github.com/awalterschulze/gographviz)
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes)  Format bytes to string.
    * [go-runewidth](https://github.com/mattn/go-runewidth) **star:326** Functions to get fixed width of the character or string.   [![There was an update last week][Green]](https://github.com/mattn/go-runewidth)   [![godoc][GoDoc]](https://godoc.org/github.com/mattn/go-runewidth)
    * [htmlquery](https://github.com/antchfx/htmlquery) **star:322** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/htmlquery)
    * [gotext](https://github.com/leonelquinteros/gotext) **star:289** GNU gettext utilities for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/leonelquinteros/gotext)
    * [goq](https://github.com/andrewstuart/goq) **star:186** Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery).   [![It hasn't been updated in the last year][Yellow]](https://github.com/andrewstuart/goq)   [![godoc][GoDoc]](https://godoc.org/github.com/andrewstuart/goq)
    * [goribot](https://github.com/zhshch2002/goribot) **star:181** A simple golang spider/scraping framework,build a spider in 3 lines.   [![godoc][GoDoc]](https://godoc.org/github.com/zhshch2002/goribot)   [![Contains Chinese documents][CN]](https://github.com/zhshch2002/goribot)   [![Archived][Archived]](https://github.com/zhshch2002/goribot)
    * [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) **star:137** Convert HTML to Markdown. Even works with entire websites and can be extended through rules.   [![godoc][GoDoc]](https://godoc.org/github.com/JohannesKaufmann/html-to-markdown)
    * [go-nmea](https://github.com/adrianmo/go-nmea) **star:135** NMEA parser library for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/adrianmo/go-nmea)
    * [omniparser](https://github.com/jf-tech/omniparser) **star:132** A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema.   [![godoc][GoDoc]](https://godoc.org/github.com/jf-tech/omniparser)
    * [sdp](https://github.com/gortc/sdp) **star:107** SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].   [![godoc][GoDoc]](https://godoc.org/github.com/gortc/sdp)
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) **star:93** Zero-width character detection and removal for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/trubitsyn/go-zero-width)
    * [podcast](https://github.com/eduncan911/podcast) **star:87** iTunes Compliant and RSS 2.0 Podcast Generator in Golang   [![godoc][GoDoc]](https://godoc.org/github.com/eduncan911/podcast)
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:79** Editorconfig file parser and manipulator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
    * [align](https://github.com/Guitarbum722/align) **star:68** A general purpose application that aligns text.   [![godoc][GoDoc]](https://godoc.org/github.com/Guitarbum722/align)
    * [go-slugify](https://github.com/mozillazg/go-slugify) **star:65** Make pretty slug with multiple languages support.   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-slugify)
    * [genex](https://github.com/alixaxel/genex) **star:60** Count and expand Regular Expressions into all matching Strings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alixaxel/genex)   [![godoc][GoDoc]](https://godoc.org/github.com/alixaxel/genex)
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown)  GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
    * [go-vcard](https://github.com/emersion/go-vcard) **star:52** Parse and format vCard.   [![godoc][GoDoc]](https://godoc.org/github.com/emersion/go-vcard)
    * [goregen](https://github.com/zach-klippenstein/goregen) **star:51** Library for generating random strings from regular expressions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zach-klippenstein/goregen)   [![godoc][GoDoc]](https://godoc.org/github.com/zach-klippenstein/goregen)
    * [guesslanguage](https://github.com/endeveit/guesslanguage) **star:51** Functions to determine the natural language of a unicode text.   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/guesslanguage)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/guesslanguage)
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:50** Fixed-width text formatting (encoder/decoder with reflection).   [![godoc][GoDoc]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
    * [allot](https://github.com/sbstjn/allot) **star:46** Placeholder and wildcard text parsing for CLI tools and bots.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sbstjn/allot)   [![godoc][GoDoc]](https://godoc.org/github.com/sbstjn/allot)
    * [did](https://github.com/ockam-network/did) **star:44** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ockam-network/did)
    * [gonameparts](https://github.com/polera/gonameparts) **star:33** Parses human names into individual name parts.   [![It hasn't been updated in the last year][Yellow]](https://github.com/polera/gonameparts)   [![godoc][GoDoc]](https://godoc.org/github.com/polera/gonameparts)
    * [Slugify](https://github.com/avelino/slugify) **star:28** Go slugify application that handles string.   [![It hasn't been updated in the last year][Yellow]](https://github.com/avelino/slugify)   [![godoc][GoDoc]](https://godoc.org/github.com/avelino/slugify)
    * [pagser](https://github.com/foolin/pagser) **star:22** Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler.   [![godoc][GoDoc]](https://godoc.org/github.com/foolin/pagser)
    * [codetree](https://github.com/aerogo/codetree) **star:16** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![It hasn't been updated in the last year][Yellow]](https://github.com/aerogo/codetree)   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/codetree)
    * [enca](https://github.com/endeveit/enca) **star:11** Minimal cgo bindings for [libenca](http://cihar.com/software/enca/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/endeveit/enca)   [![godoc][GoDoc]](https://godoc.org/github.com/endeveit/enca)
    * [syndfeed](https://github.com/zhengchun/syndfeed) **star:7** A syndication feed for Atom 1.0 and RSS 2.0.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zhengchun/syndfeed)   [![godoc][GoDoc]](https://godoc.org/github.com/zhengchun/syndfeed)
    * [bbConvert](https://github.com/CalebQ42/bbConvert) **star:6** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/CalebQ42/bbConvert)   [![godoc][GoDoc]](https://godoc.org/github.com/CalebQ42/bbConvert)
    * [doi](https://github.com/hscells/doi) **star:6** Document object identifier (doi) parser in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hscells/doi)   [![godoc][GoDoc]](https://godoc.org/github.com/hscells/doi)
    * [ltsv](https://github.com/Wing924/ltsv) **star:5** High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Wing924/ltsv)   [![godoc][GoDoc]](https://godoc.org/github.com/Wing924/ltsv)
    * [encdec](https://github.com/mickep76/encdec) **star:3** Package provides a generic interface to encoders and decodersa.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mickep76/encdec)   [![godoc][GoDoc]](https://godoc.org/github.com/mickep76/encdec)
* Utility
    * [xurls](https://github.com/mvdan/xurls) **star:753** Extract urls from text.   [![godoc][GoDoc]](https://godoc.org/github.com/mvdan/xurls)
    * [gotabulate](https://github.com/bndr/gotabulate) **star:257** Easily pretty-print your tabular data with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gotabulate)
    * [radix](https://github.com/yourbasic/radix) **star:166** fast string sorting algorithm.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yourbasic/radix)   [![godoc][GoDoc]](https://godoc.org/github.com/yourbasic/radix)
    * [regroup](https://github.com/oriser/regroup) **star:69** Match regex expression named groups into go struct using struct tags and automatic parsing.   [![godoc][GoDoc]](https://godoc.org/github.com/oriser/regroup)
    * [parth](https://github.com/codemodus/parth) **star:39** URL path segmentation parsing.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/parth)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/parth)
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:30** A sanitization-based swear filter for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)
    * [xj2go](https://github.com/stackerzzq/xj2go) **star:20** Convert xml or json to go struct.   [![godoc][GoDoc]](https://godoc.org/github.com/stackerzzq/xj2go)
    * [kace](https://github.com/codemodus/kace) **star:16** Common case conversions covering common initialisms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/kace)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/kace)
    * [Tagify](https://github.com/zoomio/tagify) **star:13** Produces a set of tags from given source.   [![godoc][GoDoc]](https://godoc.org/github.com/zoomio/tagify)
    * [TySug](https://github.com/Dynom/TySug) **star:9** Alternative suggestions with respect to keyboard layouts.   [![godoc][GoDoc]](https://godoc.org/github.com/Dynom/TySug)
    * [parseargs-go](https://github.com/nproc/parseargs-go) **star:8** string argument parser that understands quotes and backslashes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nproc/parseargs-go)   [![godoc][GoDoc]](https://godoc.org/github.com/nproc/parseargs-go)
	* [textwrap](https://github.com/isbm/textwrap) **star:2** Implementation of `textwrap` module from Python.   [![It hasn't been updated in the last year][Yellow]](https://github.com/isbm/textwrap)   [![godoc][GoDoc]](https://godoc.org/github.com/isbm/textwrap)

## Third-party APIs

*Libraries for accessing third party APIs.*

* [github](https://github.com/google/go-github) **star:7129** Go library for accessing the GitHub REST API v3.   [![star > 2000][Awesome]](https://github.com/google/go-github)   [![There was an update last week][Green]](https://github.com/google/go-github)   [![godoc][GoDoc]](https://godoc.org/github.com/google/go-github)
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) **star:6444** The official AWS SDK for the Go programming language.   [![star > 2000][Awesome]](https://github.com/aws/aws-sdk-go)   [![There was an update last week][Green]](https://github.com/aws/aws-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/aws/aws-sdk-go)
* [google](https://github.com/google/google-api-go-client) **star:2537** Auto-generated Google APIs for Go.   [![star > 2000][Awesome]](https://github.com/google/google-api-go-client)   [![There was an update last week][Green]](https://github.com/google/google-api-go-client)   [![godoc][GoDoc]](https://godoc.org/github.com/google/google-api-go-client)
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:2438** Google Cloud APIs Go Client Library.   [![star > 2000][Awesome]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![There was an update last week][Green]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
* [discordgo](https://github.com/bwmarrin/discordgo) **star:1770** Go bindings for the Discord Chat API.   [![There was an update last week][Green]](https://github.com/bwmarrin/discordgo)   [![godoc][GoDoc]](https://godoc.org/github.com/bwmarrin/discordgo)
* [stripe](https://github.com/stripe/stripe-go) **star:1294** Go client for the Stripe API.   [![There was an update last week][Green]](https://github.com/stripe/stripe-go)   [![godoc][GoDoc]](https://godoc.org/github.com/stripe/stripe-go)
* [minio-go](https://github.com/minio/minio-go) **star:1183** Minio Go Library for Amazon S3 compatible cloud storage.   [![godoc][GoDoc]](https://godoc.org/github.com/minio/minio-go)
* [go-twitter](https://github.com/dghubble/go-twitter) **star:1141** Go client library for the Twitter v1.1 APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/dghubble/go-twitter)
* [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1072** Go client library for the Twitter 1.1 API.   [![godoc][GoDoc]](https://godoc.org/github.com/ChimeraCoder/anaconda)
* [facebook](https://github.com/huandu/facebook) **star:936** Go Library that supports the Facebook Graph API.   [![godoc][GoDoc]](https://godoc.org/github.com/huandu/facebook)
* [go-jira](https://github.com/andygrunwald/go-jira) **star:847** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   [![There was an update last week][Green]](https://github.com/andygrunwald/go-jira)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-jira)
* [githubql](https://github.com/shurcooL/githubql) **star:742** Go library for accessing the GitHub GraphQL API v4.   [![There was an update last week][Green]](https://github.com/shurcooL/githubql)   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/githubql)
* [webhooks](https://github.com/go-playground/webhooks) **star:581** Webhook receiver for GitHub and Bitbucket.   [![There was an update last week][Green]](https://github.com/go-playground/webhooks)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/webhooks)
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:401** Wrapper for PayPal payment API.   [![There was an update last week][Green]](https://github.com/logpacker/PayPal-Go-SDK)   [![godoc][GoDoc]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
* [geo-golang](https://github.com/codingsince1985/geo-golang) **star:386** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/codingsince1985/geo-golang)
* [go-marathon](https://github.com/gambol99/go-marathon) **star:196** Go library for interacting with Mesosphere's Marathon PAAS.   [![godoc][GoDoc]](https://godoc.org/github.com/gambol99/go-marathon)
* [ethrpc](https://github.com/onrik/ethrpc) **star:190** Go bindings for Ethereum JSON RPC API.   [![godoc][GoDoc]](https://godoc.org/github.com/onrik/ethrpc)
* [Trello](https://github.com/adlio/trello) **star:159** Go wrapper for the Trello API.   [![godoc][GoDoc]](https://godoc.org/github.com/adlio/trello)
* [Medium](https://github.com/Medium/medium-sdk-go) **star:131** Golang SDK for Medium's OAuth2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Medium/medium-sdk-go)   [![godoc][GoDoc]](https://godoc.org/github.com/Medium/medium-sdk-go)
* [gostorm](https://github.com/jsgilmore/gostorm) **star:127** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jsgilmore/gostorm)   [![godoc][GoDoc]](https://godoc.org/github.com/jsgilmore/gostorm)
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:112** A golang package to communicate with HipChat over XMPP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/daneharrigan/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/daneharrigan/hipchat)
* [go-trending](https://github.com/andygrunwald/go-trending) **star:111** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/go-trending)
* [hipchat](https://github.com/andybons/hipchat) **star:106** This project implements a golang client library for the Hipchat API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andybons/hipchat)   [![godoc][GoDoc]](https://godoc.org/github.com/andybons/hipchat)
* [wit-go](https://github.com/wit-ai/wit-go) **star:99** Go client for wit.ai HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/wit-ai/wit-go)
* [pushover](https://github.com/gregdel/pushover) **star:88** Go wrapper for the Pushover API.   [![godoc][GoDoc]](https://godoc.org/github.com/gregdel/pushover)
* [cachet](https://github.com/andygrunwald/cachet) **star:84** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/cachet)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/cachet)
* [twitter-scraper](https://github.com/n0madic/twitter-scraper) **star:80** Scrape the Twitter Frontend API without authentication and limits.   [![There was an update last week][Green]](https://github.com/n0madic/twitter-scraper)   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/twitter-scraper)
* [igdb](https://github.com/Henry-Sarabia/igdb) **star:65** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/igdb)
* [clarifai](https://github.com/samuelcouch/clarifai) **star:57** Go client library for interfacing with the Clarifai API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/samuelcouch/clarifai)   [![godoc][GoDoc]](https://godoc.org/github.com/samuelcouch/clarifai)
* [megos](https://github.com/andygrunwald/megos) **star:56** Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/andygrunwald/megos)   [![godoc][GoDoc]](https://godoc.org/github.com/andygrunwald/megos)
* [circleci](https://github.com/jszwedko/go-circleci) **star:54** Go client library for interacting with CircleCI's API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jszwedko/go-circleci)   [![godoc][GoDoc]](https://godoc.org/github.com/jszwedko/go-circleci)
* [gosip](https://github.com/koltyakov/gosip) **star:52** Go client library SharePoint API.   [![There was an update last week][Green]](https://github.com/koltyakov/gosip)   [![godoc][GoDoc]](https://godoc.org/github.com/koltyakov/gosip)
* [simples3](https://github.com/rhnvrm/simples3) **star:50** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![There was an update last week][Green]](https://github.com/rhnvrm/simples3)   [![godoc][GoDoc]](https://godoc.org/github.com/rhnvrm/simples3)
* [gads](https://github.com/emiddleton/gads) **star:49** Google Adwords Unofficial API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/emiddleton/gads)   [![godoc][GoDoc]](https://godoc.org/github.com/emiddleton/gads)
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) **star:46** Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-amazon-product-advertising-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api)
* [ynab](https://github.com/brunomvsouza/ynab.go) **star:45** Go wrapper for the YNAB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/brunomvsouza/ynab.go)   [![godoc][GoDoc]](https://godoc.org/github.com/brunomvsouza/ynab.go)
* [gogtrends](https://github.com/groovili/gogtrends) **star:44** Google Trends Unofficial API.   [![godoc][GoDoc]](https://godoc.org/github.com/groovili/gogtrends)
* [go-xkcd](https://github.com/nishanths/go-xkcd) **star:43** Go client for the xkcd API.   [![godoc][GoDoc]](https://godoc.org/github.com/nishanths/go-xkcd)
* [uptimerobot](https://github.com/bitfield/uptimerobot) **star:43** Go wrapper and command-line client for the Uptime Robot v2 API.   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/uptimerobot)
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:42** Go MusicBrainz WS2 client library.   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/gomusicbrainz)
* [fcm](https://github.com/maddevsio/fcm) **star:40** Go library for Firebase Cloud Messaging.   [![godoc][GoDoc]](https://godoc.org/github.com/maddevsio/fcm)
* [golang-tmdb](https://github.com/cyruzin/golang-tmdb) **star:37** Golang wrapper for The Movie Database API v3.   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/golang-tmdb)
* [golyrics](https://github.com/mamal72/golyrics) **star:36** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mamal72/golyrics)   [![godoc][GoDoc]](https://godoc.org/github.com/mamal72/golyrics)
* [mixpanel](https://github.com/dukex/mixpanel) **star:36** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/dukex/mixpanel)
* [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:34** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![godoc][GoDoc]](https://godoc.org/github.com/hbagdi/go-unsplash)
* [spotify](https://github.com/rapito/go-spotify) **star:32** Go Library to access Spotify WEB API.   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-spotify)
* [translate](https://github.com/poorny/translate) **star:31** Go online translation package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/poorny/translate)   [![godoc][GoDoc]](https://godoc.org/github.com/poorny/translate)
* [gami](https://github.com/bit4bit/gami) **star:29** Go library for Asterisk Manager Interface.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bit4bit/gami)   [![godoc][GoDoc]](https://godoc.org/github.com/bit4bit/gami)   [![Archived][Archived]](https://github.com/bit4bit/gami)
* [gcm](https://github.com/Aorioli/gcm) **star:29** Go library for Google Cloud Messaging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Aorioli/gcm)   [![godoc][GoDoc]](https://godoc.org/github.com/Aorioli/gcm)
* [patreon-go](https://github.com/mxpv/patreon-go) **star:24** Go library for Patreon API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mxpv/patreon-go)   [![godoc][GoDoc]](https://godoc.org/github.com/mxpv/patreon-go)
* [steam](https://github.com/sostronk/go-steam) **star:22** Go Library to interact with Steam game servers.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sostronk/go-steam)   [![godoc][GoDoc]](https://godoc.org/github.com/sostronk/go-steam)
* [go-imgur](https://github.com/koffeinsource/go-imgur) **star:21** Go client library for [imgur](https://imgur.com)   [![godoc][GoDoc]](https://godoc.org/github.com/koffeinsource/go-imgur)
* [shopify](https://github.com/rapito/go-shopify) **star:21** Go Library to make CRUD request to the Shopify API.   [![godoc][GoDoc]](https://godoc.org/github.com/rapito/go-shopify)
* [go-twitch](https://github.com/knspriggs/go-twitch) **star:20** Go client for interacting with the Twitch v3 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/knspriggs/go-twitch)   [![godoc][GoDoc]](https://godoc.org/github.com/knspriggs/go-twitch)
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:19** Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api).   [![godoc][GoDoc]](https://godoc.org/github.com/nstratos/go-myanimelist)
* [lastpass-go](https://github.com/ansd/lastpass-go) **star:19** Go client library for the [LastPass](https://www.lastpass.com/) API.   [![There was an update last week][Green]](https://github.com/ansd/lastpass-go)   [![godoc][GoDoc]](https://godoc.org/github.com/ansd/lastpass-go)
* [brewerydb](https://github.com/naegelejd/brewerydb) **star:17** Go library for accessing the BreweryDB API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/naegelejd/brewerydb)   [![godoc][GoDoc]](https://godoc.org/github.com/naegelejd/brewerydb)
* [textbelt](https://github.com/dietsche/textbelt) **star:17** Go client for the textbelt.com txt messaging API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dietsche/textbelt)   [![godoc][GoDoc]](https://godoc.org/github.com/dietsche/textbelt)
* [codeship-go](https://github.com/codeship/codeship-go) **star:16** Go client library for interacting with Codeship's API v2.   [![godoc][GoDoc]](https://godoc.org/github.com/codeship/codeship-go)
* [airtable](https://github.com/mehanizm/airtable) **star:14** Go client library for the [Airtable API](https://airtable.com/api).   [![godoc][GoDoc]](https://godoc.org/github.com/mehanizm/airtable)
* [go-postman-collection](https://github.com/rbretecher/go-postman-collection) **star:14** Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).   [![godoc][GoDoc]](https://godoc.org/github.com/rbretecher/go-postman-collection)
* [google-analytics](https://github.com/chonthu/go-google-analytics) **star:12** Simple wrapper for easy google analytics reporting.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chonthu/go-google-analytics)   [![godoc][GoDoc]](https://godoc.org/github.com/chonthu/go-google-analytics)
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:11** Go client library for interacting with Coinpaprika's API.   [![godoc][GoDoc]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:11** Tiny Go client for HackerNews API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PaulRosset/go-hacknews)   [![godoc][GoDoc]](https://godoc.org/github.com/PaulRosset/go-hacknews)
* [go-aws-news](https://github.com/circa10a/go-aws-news) **star:10** Go application and library to fetch what's new from AWS.   [![godoc][GoDoc]](https://godoc.org/github.com/circa10a/go-aws-news)
* [google-play-scraper](https://github.com/n0madic/google-play-scraper) **star:10** Get data from Google Play Store.   [![godoc][GoDoc]](https://godoc.org/github.com/n0madic/google-play-scraper)
* [smite](https://github.com/sergiotapia/smitego) **star:10** Go package to wraps access to the Smite game API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sergiotapia/smitego)   [![godoc][GoDoc]](https://godoc.org/github.com/sergiotapia/smitego)
* [device-check-go](https://github.com/rinchsan/device-check-go) **star:9** Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1.   [![There was an update last week][Green]](https://github.com/rinchsan/device-check-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rinchsan/device-check-go)
* [go-here](https://github.com/abdullahselek/go-here) **star:8** Go client library around the HERE location based APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/abdullahselek/go-here)
* [go-sophos](https://github.com/esurdam/go-sophos) **star:8** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.   [![godoc][GoDoc]](https://godoc.org/github.com/esurdam/go-sophos)
* [gomalshare](https://github.com/MonaxGT/gomalshare) **star:8** Go library MalShare API [malshare.com](http://www.malshare.com/)   [![It hasn't been updated in the last year][Yellow]](https://github.com/MonaxGT/gomalshare)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gomalshare)
* [rrdaclient](https://github.com/Omie/rrdaclient) **star:8** Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Omie/rrdaclient)   [![godoc][GoDoc]](https://godoc.org/github.com/Omie/rrdaclient)
* [slack](https://github.com/nlopes/slack) **star:8** Slack API in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/nlopes/slack)
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) **star:7** Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).   [![It hasn't been updated in the last year][Yellow]](https://github.com/ngs/go-google-email-audit-api)   [![godoc][GoDoc]](https://godoc.org/github.com/ngs/go-google-email-audit-api)
* [gopaapi5](https://github.com/utekaravinash/gopaapi5) **star:7** Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).   [![godoc][GoDoc]](https://godoc.org/github.com/utekaravinash/gopaapi5)
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) **star:6** Go client library for the SPTrans Olho Vivo API.   [![godoc][GoDoc]](https://godoc.org/github.com/sergioaugrod/go-sptrans)
* [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
* [tumblr](https://github.com/mattcunningham/gumblr) **star:6** Go wrapper for the Tumblr v2 API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mattcunningham/gumblr)   [![godoc][GoDoc]](https://godoc.org/github.com/mattcunningham/gumblr)
* [zooz](https://github.com/gojuno/go-zooz) **star:6** Go client for the Zooz API.   [![godoc][GoDoc]](https://godoc.org/github.com/gojuno/go-zooz)
* [go-chronos](https://github.com/axelspringer/go-chronos) **star:4** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   [![It hasn't been updated in the last year][Yellow]](https://github.com/axelspringer/go-chronos)   [![godoc][GoDoc]](https://godoc.org/github.com/axelspringer/go-chronos)
* [libgoffi](https://github.com/clevabit/libgoffi) **star:3** Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration   [![godoc][GoDoc]](https://godoc.org/github.com/clevabit/libgoffi)
* [kanka](https://github.com/Henry-Sarabia/kanka) **star:2** Go client for the [Kanka API](https://kanka.io/en-US/docs/1.0).   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/kanka)
* [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) **star:2** Go library for the [RAWG Video Games Database](https://rawg.io/) API   [![godoc][GoDoc]](https://godoc.org/github.com/dimuska139/rawg-sdk-go)
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:1** The Playlyfe Rest API Go SDK.   [![It hasn't been updated in the last year][Yellow]](https://github.com/playlyfe/playlyfe-go-sdk)   [![godoc][GoDoc]](https://godoc.org/github.com/playlyfe/playlyfe-go-sdk)
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:1** Go wrapper for the TripAdvisor API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mrbenosborne/tripadvisor-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
* [vl-go](https://github.com/verifid/vl-go) **star:1** Go client library around the VerifID identity verification layer API.   [![godoc][GoDoc]](https://godoc.org/github.com/verifid/vl-go)

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) **star:34408** Command-line fuzzy finder written in Go.   [![star > 2000][Awesome]](https://github.com/junegunn/fzf)   [![There was an update last week][Green]](https://github.com/junegunn/fzf)   [![godoc][GoDoc]](https://godoc.org/github.com/junegunn/fzf)
* [hub](https://github.com/github/hub) **star:20680** wrap git commands with additional functionality to interact with github from the terminal.   [![star > 2000][Awesome]](https://github.com/github/hub)   [![There was an update last week][Green]](https://github.com/github/hub)   [![godoc][GoDoc]](https://godoc.org/github.com/github/hub)
* [ctop](https://github.com/bcicen/ctop) **star:11002** [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics.   [![star > 2000][Awesome]](https://github.com/bcicen/ctop)   [![godoc][GoDoc]](https://godoc.org/github.com/bcicen/ctop)
* [sqlx](https://github.com/jmoiron/sqlx) **star:9616** provides a set of extensions on top of the excellent built-in database/sql package.   [![star > 2000][Awesome]](https://github.com/jmoiron/sqlx)   [![There was an update last week][Green]](https://github.com/jmoiron/sqlx)   [![godoc][GoDoc]](https://godoc.org/github.com/jmoiron/sqlx)
* [wuzz](https://github.com/asciimoo/wuzz) **star:9340** Interactive cli tool for HTTP inspection.   [![star > 2000][Awesome]](https://github.com/asciimoo/wuzz)   [![There was an update last week][Green]](https://github.com/asciimoo/wuzz)   [![godoc][GoDoc]](https://godoc.org/github.com/asciimoo/wuzz)
* [goreleaser](https://github.com/goreleaser/goreleaser) **star:6858** Deliver Go binaries as fast and easily as possible.   [![star > 2000][Awesome]](https://github.com/goreleaser/goreleaser)   [![There was an update last week][Green]](https://github.com/goreleaser/goreleaser)   [![godoc][GoDoc]](https://godoc.org/github.com/goreleaser/goreleaser)
* [peco](https://github.com/peco/peco) **star:6182** Simplistic interactive filtering tool.   [![star > 2000][Awesome]](https://github.com/peco/peco)   [![godoc][GoDoc]](https://godoc.org/github.com/peco/peco)
* [usql](https://github.com/knq/usql) **star:6081** usql is a universal command-line interface for SQL databases.   [![star > 2000][Awesome]](https://github.com/knq/usql)   [![There was an update last week][Green]](https://github.com/knq/usql)   [![godoc][GoDoc]](https://godoc.org/github.com/knq/usql)
* [godropbox](https://github.com/dropbox/godropbox) **star:3924** Common libraries for writing Go services/applications from Dropbox.   [![star > 2000][Awesome]](https://github.com/dropbox/godropbox)   [![godoc][GoDoc]](https://godoc.org/github.com/dropbox/godropbox)
* [hystrix-go](https://github.com/afex/hystrix-go) **star:3034** Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.   [![star > 2000][Awesome]](https://github.com/afex/hystrix-go)   [![godoc][GoDoc]](https://godoc.org/github.com/afex/hystrix-go)
* [goreporter](https://github.com/wgliang/goreporter) **star:2804** Golang tool that does static analysis, unit testing, code review and generate code quality report.   [![star > 2000][Awesome]](https://github.com/wgliang/goreporter)   [![It hasn't been updated in the last year][Yellow]](https://github.com/wgliang/goreporter)   [![godoc][GoDoc]](https://godoc.org/github.com/wgliang/goreporter)
* [minify](https://github.com/tdewolff/minify) **star:2555** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   [![star > 2000][Awesome]](https://github.com/tdewolff/minify)   [![There was an update last week][Green]](https://github.com/tdewolff/minify)   [![godoc][GoDoc]](https://godoc.org/github.com/tdewolff/minify)
* [panicparse](https://github.com/maruel/panicparse) **star:2512** Groups similar goroutines and colorizes stack dump.   [![star > 2000][Awesome]](https://github.com/maruel/panicparse)   [![godoc][GoDoc]](https://godoc.org/github.com/maruel/panicparse)
* [go-funk](https://github.com/thoas/go-funk) **star:2315** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![star > 2000][Awesome]](https://github.com/thoas/go-funk)   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/go-funk)
* [mc](https://github.com/minio/mc) **star:1718** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   [![There was an update last week][Green]](https://github.com/minio/mc)   [![godoc][GoDoc]](https://godoc.org/github.com/minio/mc)
* [Storm](https://github.com/asdine/storm) **star:1672** Simple and powerful toolkit for BoltDB.   [![godoc][GoDoc]](https://godoc.org/github.com/asdine/storm)
* [mole](https://github.com/davrodpin/mole) **star:1429** cli app to easily create ssh tunnels.   [![godoc][GoDoc]](https://godoc.org/github.com/davrodpin/mole)
* [mergo](https://github.com/imdario/mergo) **star:1372** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   [![godoc][GoDoc]](https://godoc.org/github.com/imdario/mergo)
* [spinner](https://github.com/briandowns/spinner) **star:1353** Go package to easily provide a terminal spinner with options.   [![godoc][GoDoc]](https://godoc.org/github.com/briandowns/spinner)
* [filetype](https://github.com/h2non/filetype) **star:1239** Small package to infer the file type checking the magic numbers signature.   [![There was an update last week][Green]](https://github.com/h2non/filetype)   [![godoc][GoDoc]](https://godoc.org/github.com/h2non/filetype)
* [boilr](https://github.com/tmrts/boilr) **star:1235** Blazingly fast CLI tool for creating projects from boilerplate templates.   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/boilr)
* [jump](https://github.com/gsamokovarov/jump) **star:990** Jump helps you navigate faster by learning your habits.   [![There was an update last week][Green]](https://github.com/gsamokovarov/jump)   [![godoc][GoDoc]](https://godoc.org/github.com/gsamokovarov/jump)
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:921** Circuit Breakers in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rubyist/circuitbreaker)   [![godoc][GoDoc]](https://godoc.org/github.com/rubyist/circuitbreaker)
* [git-time-metric](https://github.com/git-time-metric/gtm) **star:853** Simple, seamless, lightweight time tracking for Git.   [![godoc][GoDoc]](https://godoc.org/github.com/git-time-metric/gtm)
* [immortal](https://github.com/immortal/immortal) **star:696** \*nix cross-platform (OS agnostic) supervisor.   [![godoc][GoDoc]](https://godoc.org/github.com/immortal/immortal)
* [hostctl](https://github.com/guumaster/hostctl) **star:638** A CLI tool to manage /etc/hosts with easy commands.   [![godoc][GoDoc]](https://godoc.org/github.com/guumaster/hostctl)
* [circuit](https://github.com/cep21/circuit) **star:534** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   [![godoc][GoDoc]](https://godoc.org/github.com/cep21/circuit)
* [htcat](https://github.com/htcat/htcat) **star:529** Parallel and Pipelined HTTP GET Utility.   [![It hasn't been updated in the last year][Yellow]](https://github.com/htcat/htcat)   [![godoc][GoDoc]](https://godoc.org/github.com/htcat/htcat)
* [godaemon](https://github.com/VividCortex/godaemon) **star:466** Utility to write daemons.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/godaemon)
* [go-dry](https://github.com/ungerik/go-dry) **star:459** DRY (don't repeat yourself) package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-dry)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-dry)
* [ergo](https://github.com/cristianoliveira/ergo) **star:444** The management of multiple local services running over different ports made easy.   [![godoc][GoDoc]](https://godoc.org/github.com/cristianoliveira/ergo)
* [koazee](https://github.com/wesovilabs/koazee) **star:441** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/koazee)
* [gopencils](https://github.com/bndr/gopencils) **star:435** Small and simple package to easily consume REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bndr/gopencils)   [![godoc][GoDoc]](https://godoc.org/github.com/bndr/gopencils)
* [mimetype](https://github.com/gabriel-vasile/mimetype) **star:423** Package for MIME type detection based on magic numbers.   [![There was an update last week][Green]](https://github.com/gabriel-vasile/mimetype)   [![godoc][GoDoc]](https://godoc.org/github.com/gabriel-vasile/mimetype)
* [request](https://github.com/mozillazg/request) **star:392** Go HTTP Requests for Humans™.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mozillazg/request)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/request)
* [Deepcopier](https://github.com/ulule/deepcopier) **star:328** Simple struct copying for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/deepcopier)
* [go-rate](https://github.com/beefsack/go-rate) **star:319** Timed rate limiter for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/beefsack/go-rate)
* [clockwork](https://github.com/jonboulle/clockwork) **star:316** A simple fake clock for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/jonboulle/clockwork)
* [gubrak](https://github.com/novalagung/gubrak) **star:310** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/novalagung/gubrak)
* [retry](https://github.com/kamilsk/retry) **star:277** The most advanced functional mechanism to perform actions repetitively until successful.   [![There was an update last week][Green]](https://github.com/kamilsk/retry)   [![godoc][GoDoc]](https://godoc.org/github.com/kamilsk/retry)
* [delve](https://github.com/derekparker/delve) **star:267** Go debugger.   [![godoc][GoDoc]](https://godoc.org/github.com/derekparker/delve)
* [create-go-app](https://github.com/create-go-app/cli) **star:258** A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command.   [![godoc][GoDoc]](https://godoc.org/github.com/create-go-app/cli)
* [gohper](https://github.com/cosiner/gohper) **star:253** Various tools/modules help for development.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cosiner/gohper)   [![godoc][GoDoc]](https://godoc.org/github.com/cosiner/gohper)   [![Archived][Archived]](https://github.com/cosiner/gohper)
* [serve](https://github.com/syntaqx/serve) **star:232** A static http server anywhere you need.   [![godoc][GoDoc]](https://godoc.org/github.com/syntaqx/serve)
* [go-trigger](https://github.com/sadlil/go-trigger) **star:210** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sadlil/go-trigger)   [![godoc][GoDoc]](https://godoc.org/github.com/sadlil/go-trigger)
* [scany](https://github.com/georgysavva/scany) **star:200** Library for scanning data from a database into Go structs and more.   [![godoc][GoDoc]](https://godoc.org/github.com/georgysavva/scany)
* [util](https://github.com/shomali11/util) **star:194** Collection of useful utility functions. (strings, concurrency, manipulations, ...).   [![godoc][GoDoc]](https://godoc.org/github.com/shomali11/util)
* [gotenv](https://github.com/subosito/gotenv) **star:188** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/subosito/gotenv)
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:169** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![It hasn't been updated in the last year][Yellow]](https://github.com/wendigo/go-bind-plugin)   [![godoc][GoDoc]](https://godoc.org/github.com/wendigo/go-bind-plugin)
* [Death](https://github.com/vrecan/death) **star:164** Managing go application shutdown with signals.   [![godoc][GoDoc]](https://godoc.org/github.com/vrecan/death)
* [rerun](https://github.com/ivpusic/rerun) **star:161** Recompiling and rerunning go apps when source changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/rerun)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/rerun)
* [moldova](https://github.com/StabbyCutyou/moldova) **star:158** Utility for generating random data based on an input template.   [![It hasn't been updated in the last year][Yellow]](https://github.com/StabbyCutyou/moldova)   [![godoc][GoDoc]](https://godoc.org/github.com/StabbyCutyou/moldova)
* [toolbox](https://github.com/viant/toolbox) **star:155** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   [![godoc][GoDoc]](https://godoc.org/github.com/viant/toolbox)
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:150** XML Sitemap generator written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ikeikeikeike/go-sitemap-generator)   [![godoc][GoDoc]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
* [apm](https://github.com/topfreegames/apm) **star:149** Process manager for Golang applications with an HTTP API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/topfreegames/apm)   [![godoc][GoDoc]](https://godoc.org/github.com/topfreegames/apm)
* [robustly](https://github.com/VividCortex/robustly) **star:149** Runs functions resiliently, catching and restarting panics.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/robustly)
* [chyle](https://github.com/antham/chyle) **star:135** Changelog generator using a git repository with multiple configuration possibilities.   [![godoc][GoDoc]](https://godoc.org/github.com/antham/chyle)
* [onecache](https://github.com/adelowo/onecache) **star:115** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).   [![godoc][GoDoc]](https://godoc.org/github.com/adelowo/onecache)
* [lrserver](https://github.com/jaschaephraim/lrserver) **star:112** LiveReload server for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jaschaephraim/lrserver)   [![godoc][GoDoc]](https://godoc.org/github.com/jaschaephraim/lrserver)
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:109** Pure Go bsdiff and bspatch libraries and CLI tools.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gabstv/go-bsdiff)   [![godoc][GoDoc]](https://godoc.org/github.com/gabstv/go-bsdiff)
* [xferspdy](https://github.com/monmohan/xferspdy) **star:85** Xferspdy provides binary diff and patch library in golang.   [![godoc][GoDoc]](https://godoc.org/github.com/monmohan/xferspdy)
* [mssqlx](https://github.com/linxGnu/mssqlx) **star:82** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/mssqlx)
* [nostromo](https://github.com/pokanop/nostromo) **star:81** CLI for building powerful aliases.   [![godoc][GoDoc]](https://godoc.org/github.com/pokanop/nostromo)
* [countries](https://github.com/biter777/countries) **star:80** Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts.   [![godoc][GoDoc]](https://godoc.org/github.com/biter777/countries)
* [pm](https://github.com/VividCortex/pm) **star:79** Process (i.e. goroutine) manager with an HTTP API.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/pm)
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:77** SeaweedFS client library with almost full features.   [![godoc][GoDoc]](https://godoc.org/github.com/linxGnu/goseaweedfs)
* [go-health](https://github.com/Talento90/go-health) **star:76** Health package simplifies the way you add health check to your services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Talento90/go-health)   [![godoc][GoDoc]](https://godoc.org/github.com/Talento90/go-health)
* [repeat](https://github.com/ssgreg/repeat) **star:72** Go implementation of different backoff strategies useful for retrying operations and heartbeating.   [![godoc][GoDoc]](https://godoc.org/github.com/ssgreg/repeat)
* [sorty](https://github.com/jfcg/sorty) **star:71** Fast Concurrent / Parallel Sorting.   [![There was an update last week][Green]](https://github.com/jfcg/sorty)   [![godoc][GoDoc]](https://godoc.org/github.com/jfcg/sorty)
* [scan](https://github.com/blockloop/scan) **star:69** Scan golang `sql.Rows` directly to structs, slices, or primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/blockloop/scan)
* [netbug](https://github.com/e-dard/netbug) **star:68** Easy remote profiling of your services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/e-dard/netbug)   [![godoc][GoDoc]](https://godoc.org/github.com/e-dard/netbug)
* [UNIS](https://github.com/esemplastic/unis) **star:68** Common Architecture™ for String Utilities in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/esemplastic/unis)   [![godoc][GoDoc]](https://godoc.org/github.com/esemplastic/unis)
* [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) **star:67** Pattern matching libray.   [![godoc][GoDoc]](https://godoc.org/github.com/alexpantyukhin/go-pattern-match)
* [multitick](https://github.com/VividCortex/multitick) **star:64** Multiplexor for aligned tickers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/multitick)
* [handy](https://github.com/miguelpragier/handy) **star:62** Many utilities and helpers like string handlers/formatters and validators.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelpragier/handy)
* [mimemagic](https://github.com/zRedShift/mimemagic) **star:58** Pure Go ultra performant MIME sniffing library/utility.   [![godoc][GoDoc]](https://godoc.org/github.com/zRedShift/mimemagic)
* [go-astitodo](https://github.com/asticode/go-astitodo) **star:56** Parse TODOs in your GO code.   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astitodo)
* [minquery](https://github.com/icza/minquery) **star:56** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).   [![godoc][GoDoc]](https://godoc.org/github.com/icza/minquery)
* [goreadability](https://github.com/philipjkim/goreadability) **star:55** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![It hasn't been updated in the last year][Yellow]](https://github.com/philipjkim/goreadability)   [![godoc][GoDoc]](https://godoc.org/github.com/philipjkim/goreadability)
* [pgo](https://github.com/arthurkushman/pgo) **star:52** Convenient functions for PHP community.   [![godoc][GoDoc]](https://godoc.org/github.com/arthurkushman/pgo)
* [golog](https://github.com/mlimaloureiro/golog) **star:50** Easy and lightweight CLI tool to time track your tasks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mlimaloureiro/golog)   [![godoc][GoDoc]](https://godoc.org/github.com/mlimaloureiro/golog)
* [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) **star:50** Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines.   [![godoc][GoDoc]](https://godoc.org/github.com/gobeam/mongo-go-pagination)
* [copy-pasta](https://github.com/jutkko/copy-pasta) **star:47** Universal multi-workstation clipboard that uses S3 like backend for the storage.   [![godoc][GoDoc]](https://godoc.org/github.com/jutkko/copy-pasta)
* [cmd](https://github.com/SimonBaeumer/cmd) **star:46** Library for executing shell commands on osx, windows and linux.   [![godoc][GoDoc]](https://godoc.org/github.com/SimonBaeumer/cmd)
* [retry](https://github.com/thedevsaddam/retry) **star:46** Simple and easy retry mechanism package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/retry)
* [goback](https://github.com/carlescere/goback) **star:43** Go simple exponential backoff package.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carlescere/goback)   [![godoc][GoDoc]](https://godoc.org/github.com/carlescere/goback)
* [golarm](https://github.com/msempere/golarm) **star:43** Fire alarms with system events.   [![It hasn't been updated in the last year][Yellow]](https://github.com/msempere/golarm)   [![godoc][GoDoc]](https://godoc.org/github.com/msempere/golarm)
* [intrinsic](https://github.com/mengzhuo/intrinsic) **star:42** Use x86 SIMD without writing any assembly code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mengzhuo/intrinsic)   [![godoc][GoDoc]](https://godoc.org/github.com/mengzhuo/intrinsic)   [![Archived][Archived]](https://github.com/mengzhuo/intrinsic)
* [filter](https://github.com/gookit/filter) **star:40** provide filtering, sanitizing, and conversion of Go data.   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/filter)
* [gpath](https://github.com/tenntenn/gpath) **star:40** Library to simplify access struct fields with Go's expression in reflection.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tenntenn/gpath)   [![godoc][GoDoc]](https://godoc.org/github.com/tenntenn/gpath)
* [retry-go](https://github.com/rafaeljesus/retry-go) **star:39** Retrying made simple and easy for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafaeljesus/retry-go)   [![godoc][GoDoc]](https://godoc.org/github.com/rafaeljesus/retry-go)
* [beyond](https://github.com/wesovilabs/beyond) **star:38** The Go tool that will drive you to the AOP world!   [![It hasn't been updated in the last year][Yellow]](https://github.com/wesovilabs/beyond)   [![godoc][GoDoc]](https://godoc.org/github.com/wesovilabs/beyond)
* [dbt](https://github.com/nikogura/dbt) **star:38** A framework for running self-updating signed binaries from a central, trusted repository.   [![godoc][GoDoc]](https://godoc.org/github.com/nikogura/dbt)
* [slice](https://github.com/psampaz/slice) **star:38** Type-safe functions for common Go slice operations.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/slice)
* [myhttp](https://github.com/inancgumus/myhttp) **star:33** Simple API to make HTTP GET requests with timeout support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/inancgumus/myhttp)   [![godoc][GoDoc]](https://godoc.org/github.com/inancgumus/myhttp)   [![Archived][Archived]](https://github.com/inancgumus/myhttp)
* [go-lock](https://github.com/viney-shih/go-lock) **star:32** go-lock is a lock library implementing read-write mutex and read-write trylock without starvation.   [![godoc][GoDoc]](https://godoc.org/github.com/viney-shih/go-lock)
* [rclient](https://github.com/zpatrick/rclient) **star:32** Readable, flexible, simple-to-use client for REST APIs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/rclient)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/rclient)
* [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:30** Go library for encoding structs into Header fields.   [![There was an update last week][Green]](https://github.com/mozillazg/go-httpheader)   [![godoc][GoDoc]](https://godoc.org/github.com/mozillazg/go-httpheader)
* [evaluator](https://github.com/nullne/evaluator) **star:29** Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend.   [![godoc][GoDoc]](https://godoc.org/github.com/nullne/evaluator)
* [gostrutils](https://github.com/ik5/gostrutils) **star:29** Collections of string manipulation and conversion functions.   [![godoc][GoDoc]](https://godoc.org/github.com/ik5/gostrutils)
* [changie](https://github.com/miniscruff/changie) **star:28** Automated changelog tool for preparing releases with lots of customization options.   [![godoc][GoDoc]](https://godoc.org/github.com/miniscruff/changie)
* [tome](https://github.com/cyruzin/tome) **star:27** Tome was designed to paginate simple RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/cyruzin/tome)
* [equalizer](https://github.com/reugn/equalizer) **star:24** Quota manager and rate limiter collection for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/reugn/equalizer)
* [generate](https://github.com/go-playground/generate) **star:24** runs go generate recursively on a specified path or environment variable and can filter by regex.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/generate)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/generate)
* [slicer](https://github.com/leaanthony/slicer) **star:24** Makes working with slices easier.   [![godoc][GoDoc]](https://godoc.org/github.com/leaanthony/slicer)
* [ugo](https://github.com/alxrm/ugo) **star:24** ugo is slice toolbox with concise syntax for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alxrm/ugo)   [![godoc][GoDoc]](https://godoc.org/github.com/alxrm/ugo)
* [limiters](https://github.com/mennanov/limiters) **star:23** Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.   [![godoc][GoDoc]](https://godoc.org/github.com/mennanov/limiters)
* [goplaceholder](https://github.com/michiwend/goplaceholder) **star:22** a small golang lib to generate placeholder images.   [![It hasn't been updated in the last year][Yellow]](https://github.com/michiwend/goplaceholder)   [![godoc][GoDoc]](https://godoc.org/github.com/michiwend/goplaceholder)
* [rerate](https://github.com/abo/rerate) **star:18** Redis-based rate counter and rate limiter for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/abo/rerate)   [![godoc][GoDoc]](https://godoc.org/github.com/abo/rerate)
* [ghokin](https://github.com/antham/ghokin) **star:17** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![godoc][GoDoc]](https://godoc.org/github.com/antham/ghokin)
* [ctxutil](https://github.com/posener/ctxutil) **star:16** A collection of utility functions for contexts.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/ctxutil)
* [dlog](https://github.com/kirillDanshin/dlog) **star:16** Compile-time controlled logger to make your release smaller without removing debug calls.   [![It hasn't been updated in the last year][Yellow]](https://github.com/kirillDanshin/dlog)   [![godoc][GoDoc]](https://godoc.org/github.com/kirillDanshin/dlog)
* [filler](https://github.com/yaronsumel/filler) **star:16** small utility to fill structs using "fill" tag.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yaronsumel/filler)   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/filler)
* [structs](https://github.com/PumpkinSeed/structs) **star:16** Implement simple functions to manipulate structs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PumpkinSeed/structs)   [![godoc][GoDoc]](https://godoc.org/github.com/PumpkinSeed/structs)
* [okrun](https://github.com/xta/okrun) **star:15** go run error steamroller.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xta/okrun)   [![godoc][GoDoc]](https://godoc.org/github.com/xta/okrun)
* [shutdown](https://github.com/ztrue/shutdown) **star:15** App shutdown hooks for `os.Signal` handling.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ztrue/shutdown)   [![godoc][GoDoc]](https://godoc.org/github.com/ztrue/shutdown)
* [mimesniffer](https://github.com/aofei/mimesniffer) **star:14** A MIME type sniffer for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/mimesniffer)
* [rest-go](https://github.com/edermanoel94/rest-go) **star:14** A package that provide many helpful methods for working with rest api.   [![godoc][GoDoc]](https://godoc.org/github.com/edermanoel94/rest-go)
* [command](https://github.com/txgruppi/command) **star:13** Command pattern for Go with thread safe serial and parallel dispatcher.   [![It hasn't been updated in the last year][Yellow]](https://github.com/txgruppi/command)   [![godoc][GoDoc]](https://godoc.org/github.com/txgruppi/command)
* [jsend](https://github.com/clevergo/jsend) **star:13** JSend's implementation writen in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/clevergo/jsend)
* [backscanner](https://github.com/icza/backscanner) **star:12** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![godoc][GoDoc]](https://godoc.org/github.com/icza/backscanner)
* [retry](https://github.com/shafreeck/retry) **star:11** A pretty simple library to ensure your work to be done.   [![godoc][GoDoc]](https://godoc.org/github.com/shafreeck/retry)
* [go-convert](https://github.com/Eun/go-convert) **star:10** Package go-convert enbles you to convert a value into another type.   [![godoc][GoDoc]](https://godoc.org/github.com/Eun/go-convert)
* [go-countries](https://github.com/mikekonan/go-countries) **star:8** Lightweight lookup over ISO-3166 codes.   [![godoc][GoDoc]](https://godoc.org/github.com/mikekonan/go-countries)
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:7** Go package for working with Problem Details.   [![godoc][GoDoc]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
* [retry](https://github.com/percolate/retry) **star:7** A simple but highly configurable retry package for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/percolate/retry)
* [statiks](https://github.com/janiltonmaciel/statiks) **star:7** Fast, zero-configuration, static HTTP filer server.   [![godoc][GoDoc]](https://godoc.org/github.com/janiltonmaciel/statiks)
* [ptr](https://github.com/gotidy/ptr) **star:6** Package that provide functions for simplified creation of pointers from constants of basic types.   [![godoc][GoDoc]](https://godoc.org/github.com/gotidy/ptr)
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:6** Slice conversion between primitive types.   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
* [blank](https://github.com/Henry-Sarabia/blank) **star:5** Verify or remove blanks and whitespace from strings.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Henry-Sarabia/blank)   [![godoc][GoDoc]](https://godoc.org/github.com/Henry-Sarabia/blank)
* [nfdump](https://github.com/chrispassas/nfdump) **star:4** Read nfdump netflow files.   [![godoc][GoDoc]](https://godoc.org/github.com/chrispassas/nfdump)
* [silk](https://github.com/chrispassas/silk) **star:4** Read silk netflow files.   [![godoc][GoDoc]](https://godoc.org/github.com/chrispassas/silk)
* [copy](https://github.com/gotidy/copy) **star:2** Package for fast copying structs of different types.   [![godoc][GoDoc]](https://godoc.org/github.com/gotidy/copy)
* [go-safe](https://github.com/kenkyu392/go-safe) **star:2** Panic-safe sandbox.   [![godoc][GoDoc]](https://godoc.org/github.com/kenkyu392/go-safe)
* [goctx](https://github.com/zerosnake0/goctx) **star:2** Get your context value with high performance.   [![godoc][GoDoc]](https://godoc.org/github.com/zerosnake0/goctx)
* [lets-go](https://github.com/aplescia-chwy/lets-go) **star:2** Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities.   [![godoc][GoDoc]](https://godoc.org/github.com/aplescia-chwy/lets-go)
* [olaf](https://github.com/btnguyen2k/olaf) **star:2** Twitter Snowflake implemented in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/btnguyen2k/olaf)   [![godoc][GoDoc]](https://godoc.org/github.com/btnguyen2k/olaf)
* [tik](https://github.com/andy2046/tik) **star:2** Simple and easy timing wheel package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/andy2046/tik)
* [bleep](https://github.com/sinhashubham95/bleep) **star:1** Perform any number of actions on any set of OS signals in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/sinhashubham95/bleep)

## UUID

*Libraries for working with UUIDs.*

* [uuid](https://github.com/google/uuid) **star:2397** Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.   [![star > 2000][Awesome]](https://github.com/google/uuid)   [![There was an update last week][Green]](https://github.com/google/uuid)   [![godoc][GoDoc]](https://godoc.org/github.com/google/uuid)
* [ulid](https://github.com/oklog/ulid) **star:2129** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   [![star > 2000][Awesome]](https://github.com/oklog/ulid)   [![godoc][GoDoc]](https://godoc.org/github.com/oklog/ulid)
* [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands.
* [uuid](https://github.com/gofrs/uuid) **star:876** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   [![godoc][GoDoc]](https://godoc.org/github.com/gofrs/uuid)
* [wuid](https://github.com/edwingeng/wuid) **star:386** An extremely fast unique number generator, 10-135 times faster than UUID.   [![godoc][GoDoc]](https://godoc.org/github.com/edwingeng/wuid)
* [sno](https://github.com/muyo/sno) **star:42** Compact, sortable and fast unique IDs with embedded metadata.   [![godoc][GoDoc]](https://godoc.org/github.com/muyo/sno)
* [goid](https://github.com/jakehl/goid) **star:29** Generate and Parse RFC4122 compliant V4 UUIDs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jakehl/goid)   [![godoc][GoDoc]](https://godoc.org/github.com/jakehl/goid)
* [nanoid](https://github.com/aidarkhanov/nanoid) **star:29** A tiny and efficient Go unique string ID generator.   [![godoc][GoDoc]](https://godoc.org/github.com/aidarkhanov/nanoid)
* [uuid](https://github.com/agext/uuid) **star:12** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.   [![godoc][GoDoc]](https://godoc.org/github.com/agext/uuid)
* [gouid](https://github.com/twharmon/gouid) **star:5** Generate cryptographically secure random string IDs with just one allocation.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/gouid)

## Validation

*Libraries for validation.*

* [validator](https://github.com/go-playground/validator) **star:6948** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   [![star > 2000][Awesome]](https://github.com/go-playground/validator)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/validator)
* [govalidator](https://github.com/asaskevich/govalidator) **star:4593** Validators and sanitizers for strings, numerics, slices and structs.   [![star > 2000][Awesome]](https://github.com/asaskevich/govalidator)   [![There was an update last week][Green]](https://github.com/asaskevich/govalidator)   [![godoc][GoDoc]](https://godoc.org/github.com/asaskevich/govalidator)
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:1867** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   [![There was an update last week][Green]](https://github.com/go-ozzo/ozzo-validation)   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
* [govalidator](https://github.com/thedevsaddam/govalidator) **star:957** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/govalidator)
* [validate](https://github.com/gookit/validate) **star:358** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   [![There was an update last week][Green]](https://github.com/gookit/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/validate)   [![Contains Chinese documents][CN]](https://github.com/gookit/validate)
* [checkdigit](https://github.com/osamingo/checkdigit) **star:73** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).   [![godoc][GoDoc]](https://godoc.org/github.com/osamingo/checkdigit)
* [jio](https://github.com/faceair/jio) **star:55** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   [![godoc][GoDoc]](https://godoc.org/github.com/faceair/jio)   [![Contains Chinese documents][CN]](https://github.com/faceair/jio)
* [validate](https://github.com/gobuffalo/validate) **star:55** This package provides a framework for writing validations for Go applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gobuffalo/validate)
* [gody](https://github.com/guiferpa/gody) **star:47** :balloon: A lightweight struct validator for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/guiferpa/gody)
* [terraform-validator](https://github.com/thazelart/terraform-validator) **star:40** A norms and conventions validator for Terraform.   [![godoc][GoDoc]](https://godoc.org/github.com/thazelart/terraform-validator)
* [govalid](https://github.com/twharmon/govalid) **star:23** Fast, tag-based validation for structs.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/govalid)

## Version Control

*Libraries for version control.*

* [go-git](https://github.com/go-git/go-git) **star:1941** highly extensible Git implementation in pure Go.   [![godoc][GoDoc]](https://godoc.org/github.com/go-git/go-git)
* [git2go](https://github.com/libgit2/git2go) **star:1582** Go bindings for libgit2.   [![godoc][GoDoc]](https://godoc.org/github.com/libgit2/git2go)
* [hercules](https://github.com/src-d/hercules) **star:1236** gaining advanced insights from Git repository history.   [![godoc][GoDoc]](https://godoc.org/github.com/src-d/hercules)
* [gh](https://github.com/rjeczalik/gh) **star:75** Scriptable server and net/http middleware for GitHub Webhooks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjeczalik/gh)   [![godoc][GoDoc]](https://godoc.org/github.com/rjeczalik/gh)
* [go-vcs](https://github.com/sourcegraph/go-vcs) **star:75** manipulate and inspect VCS repositories in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/sourcegraph/go-vcs)
* [hgo](https://github.com/beyang/hgo) **star:13** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   [![It hasn't been updated in the last year][Yellow]](https://github.com/beyang/hgo)   [![godoc][GoDoc]](https://godoc.org/github.com/beyang/hgo)

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) **star:1493** Comprehensive Go bindings for FFmpeg.   [![godoc][GoDoc]](https://godoc.org/github.com/giorgisio/goav)
* [m3u8](https://github.com/grafov/m3u8) **star:784** Parser and generator library of M3U8 playlists for Apple HLS.   [![godoc][GoDoc]](https://godoc.org/github.com/grafov/m3u8)
* [gmf](https://github.com/3d0c/gmf) **star:653** Go bindings for FFmpeg av\* libraries.   [![godoc][GoDoc]](https://godoc.org/github.com/3d0c/gmf)
* [go-astits](https://github.com/asticode/go-astits) **star:337** Parse and demux MPEG Transport Streams (.ts) natively in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astits)
* [go-astisub](https://github.com/asticode/go-astisub) **star:297** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   [![There was an update last week][Green]](https://github.com/asticode/go-astisub)   [![godoc][GoDoc]](https://godoc.org/github.com/asticode/go-astisub)
* [libvlc-go](https://github.com/adrg/libvlc-go) **star:178** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).   [![There was an update last week][Green]](https://github.com/adrg/libvlc-go)   [![godoc][GoDoc]](https://godoc.org/github.com/adrg/libvlc-go)
* [gst](https://github.com/ziutek/gst) **star:159** Go bindings for GStreamer.   [![godoc][GoDoc]](https://godoc.org/github.com/ziutek/gst)
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) **star:69** Parser and generator library for Apple m3u8 playlists.   [![godoc][GoDoc]](https://godoc.org/github.com/quangngotan95/go-m3u8)
* [v4l](https://github.com/korandiz/v4l) **star:49** Video capture library for Linux, written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/korandiz/v4l)
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) **star:13** Subtitle format support for go. Supports .srt, .ttml, and .ass.   [![godoc][GoDoc]](https://godoc.org/github.com/wargarblgarbl/libgosubs)
* [go-mpd](https://github.com/unki2aut/go-mpd) **star:5** Parser and generator library for MPEG-DASH manifest files.   [![godoc][GoDoc]](https://godoc.org/github.com/unki2aut/go-mpd)

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) **star:45147** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   [![star > 2000][Awesome]](https://github.com/gin-gonic/gin)   [![There was an update last week][Green]](https://github.com/gin-gonic/gin)   [![godoc][GoDoc]](https://godoc.org/github.com/gin-gonic/gin)
* [Echo](https://github.com/labstack/echo) **star:19027** High performance, minimalist Go web framework.   [![star > 2000][Awesome]](https://github.com/labstack/echo)   [![godoc][GoDoc]](https://godoc.org/github.com/labstack/echo)
* [Revel](https://github.com/revel/revel) **star:12115** High-productivity web framework for the Go language.   [![star > 2000][Awesome]](https://github.com/revel/revel)   [![godoc][GoDoc]](https://godoc.org/github.com/revel/revel)
* [Fiber](https://github.com/gofiber/fiber) **star:11343** An Express.js inspired web framework build on Fasthttp.   [![star > 2000][Awesome]](https://github.com/gofiber/fiber)   [![There was an update last week][Green]](https://github.com/gofiber/fiber)   [![godoc][GoDoc]](https://godoc.org/github.com/gofiber/fiber)
* [Goa](https://github.com/goadesign/goa) **star:4111** Goa provides a holistic approach for developing remote APIs and microservices in Go.   [![star > 2000][Awesome]](https://github.com/goadesign/goa)   [![There was an update last week][Green]](https://github.com/goadesign/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goadesign/goa)
* [go-json-rest](https://github.com/ant0ine/go-json-rest) **star:3452** Quick and easy way to setup a RESTful JSON API.   [![star > 2000][Awesome]](https://github.com/ant0ine/go-json-rest)   [![There was an update last week][Green]](https://github.com/ant0ine/go-json-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ant0ine/go-json-rest)
* [Gizmo](https://github.com/NYTimes/gizmo) **star:3337** Microservice toolkit used by the New York Times.   [![star > 2000][Awesome]](https://github.com/NYTimes/gizmo)   [![There was an update last week][Green]](https://github.com/NYTimes/gizmo)   [![godoc][GoDoc]](https://godoc.org/github.com/NYTimes/gizmo)
* [Macaron](https://github.com/go-macaron/macaron) **star:3117** Macaron is a high productive and modular design web framework in Go.   [![star > 2000][Awesome]](https://github.com/go-macaron/macaron)   [![godoc][GoDoc]](https://godoc.org/github.com/go-macaron/macaron)
* [utron](https://github.com/gernest/utron) **star:2180** Lightweight MVC framework for Go(Golang).   [![star > 2000][Awesome]](https://github.com/gernest/utron)   [![It hasn't been updated in the last year][Yellow]](https://github.com/gernest/utron)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/utron)
* [tigertonic](https://github.com/rcrowley/go-tigertonic) **star:1001** Go framework for building JSON web services inspired by Dropwizard.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rcrowley/go-tigertonic)   [![godoc][GoDoc]](https://godoc.org/github.com/rcrowley/go-tigertonic)
* [tango](https://github.com/lunny/tango) **star:837** Micro & pluggable web framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lunny/tango)   [![godoc][GoDoc]](https://godoc.org/github.com/lunny/tango)   [![Contains Chinese documents][CN]](https://github.com/lunny/tango)
* [Goyave](https://github.com/System-Glitch/goyave) **star:788** Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities.   [![There was an update last week][Green]](https://github.com/System-Glitch/goyave)   [![godoc][GoDoc]](https://godoc.org/github.com/System-Glitch/goyave)
* [gongular](https://github.com/mustafaakin/gongular) **star:437** Fast Go web framework with input mapping/validation and (DI) Dependency Injection.   [![godoc][GoDoc]](https://godoc.org/github.com/mustafaakin/gongular)
* [Gearbox](https://github.com/abahmed/gearbox) **star:415** A web framework written in Go with a focus on high performance and memory optimization.   [![godoc][GoDoc]](https://godoc.org/github.com/abahmed/gearbox)
* [neo](https://github.com/ivpusic/neo) **star:410** Neo is minimal and fast Go Web Framework with extremely simple API.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ivpusic/neo)   [![godoc][GoDoc]](https://godoc.org/github.com/ivpusic/neo)
* [Air](https://github.com/aofei/air) **star:399** An ideally refined web framework for Go.   [![There was an update last week][Green]](https://github.com/aofei/air)   [![godoc][GoDoc]](https://godoc.org/github.com/aofei/air)
* [mango](https://github.com/paulbellamy/mango) **star:354** Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.   [![It hasn't been updated in the last year][Yellow]](https://github.com/paulbellamy/mango)   [![godoc][GoDoc]](https://godoc.org/github.com/paulbellamy/mango)
* [Aero](https://github.com/aerogo/aero) **star:346** High-performance web framework for Go, reaches top scores in Lighthouse.   [![godoc][GoDoc]](https://godoc.org/github.com/aerogo/aero)
* [Gondola](https://github.com/rainycape/gondola) **star:311** The web framework for writing faster sites, faster.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rainycape/gondola)   [![godoc][GoDoc]](https://godoc.org/github.com/rainycape/gondola)
* [Golf](https://github.com/dinever/golf) **star:246** Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dinever/golf)   [![godoc][GoDoc]](https://godoc.org/github.com/dinever/golf)
* [Flamingo](https://github.com/i-love-flamingo/flamingo) **star:187** Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo)
* [WebGo](https://github.com/bnkamalesh/webgo) **star:154** A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/webgo)
* [Ginrpc](https://github.com/xxjwxc/ginrpc) **star:150** Gin parameter automatic binding tool,gin rpc tools.   [![godoc][GoDoc]](https://godoc.org/github.com/xxjwxc/ginrpc)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/ginrpc)
* [hiboot](https://github.com/hidevopsio/hiboot) **star:148** hiboot is a high performance web application framework with auto configuration and dependency injection support.   [![There was an update last week][Green]](https://github.com/hidevopsio/hiboot)   [![godoc][GoDoc]](https://godoc.org/github.com/hidevopsio/hiboot)   [![Contains Chinese documents][CN]](https://github.com/hidevopsio/hiboot)
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) **star:143** Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.   [![There was an update last week][Green]](https://github.com/i-love-flamingo/flamingo-commerce)   [![godoc][GoDoc]](https://godoc.org/github.com/i-love-flamingo/flamingo-commerce)
* [go-rest](https://github.com/ungerik/go-rest) **star:126** Small and evil REST framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ungerik/go-rest)   [![godoc][GoDoc]](https://godoc.org/github.com/ungerik/go-rest)
* [uAdmin](https://github.com/uadmin/uadmin) **star:121** Fully featured web framework for Golang, inspired by Django.   [![godoc][GoDoc]](https://godoc.org/github.com/uadmin/uadmin)
* [appy](https://github.com/appist/appy) **star:79** An opinionated productive web framework that helps scaling business easier.   [![There was an update last week][Green]](https://github.com/appist/appy)   [![godoc][GoDoc]](https://godoc.org/github.com/appist/appy)
* [vox](https://github.com/aisk/vox) **star:74** A golang web framework for humans, inspired by Koa heavily.   [![There was an update last week][Green]](https://github.com/aisk/vox)   [![godoc][GoDoc]](https://godoc.org/github.com/aisk/vox)
* [Golax](https://github.com/fulldump/golax) **star:73** A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.   [![It hasn't been updated in the last year][Yellow]](https://github.com/fulldump/golax)   [![godoc][GoDoc]](https://godoc.org/github.com/fulldump/golax)
* [Microservice](https://github.com/claygod/microservice) **star:73** The framework for the creation of microservices, written in Golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/claygod/microservice)   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/microservice)
* [patron](https://github.com/beatlabs/patron) **star:69** Patron is a microservice framework following best cloud practices with a focus on productivity.   [![There was an update last week][Green]](https://github.com/beatlabs/patron)   [![godoc][GoDoc]](https://godoc.org/github.com/beatlabs/patron)
* [YARF](https://github.com/yarf-framework/yarf) **star:61** Fast micro-framework designed to build REST APIs and web services in a fast and simple way.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yarf-framework/yarf)   [![godoc][GoDoc]](https://godoc.org/github.com/yarf-framework/yarf)
* [Fireball](https://github.com/zpatrick/fireball) **star:56** More "natural" feeling web framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/zpatrick/fireball)   [![godoc][GoDoc]](https://godoc.org/github.com/zpatrick/fireball)
* [rux](https://github.com/gookit/rux) **star:56** Simple and fast web framework for build golang HTTP applications.   [![There was an update last week][Green]](https://github.com/gookit/rux)   [![godoc][GoDoc]](https://godoc.org/github.com/gookit/rux)   [![Contains Chinese documents][CN]](https://github.com/gookit/rux)
* [Beego](https://github.com/astaxie/beego) **star:53** beego is an open-source, high-performance web framework for the Go programming language.   [![There was an update last week][Green]](https://github.com/astaxie/beego)   [![Contains Chinese documents][CN]](https://github.com/astaxie/beego)
* [Buffalo](http://gobuffalo.io)  Bringing the productivity of Rails to Go!
* [aah](https://aahframework.org)  Scalable, performant, rapid development Web framework for Go.
* [goa](https://github.com/goa-go/goa) **star:44** goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goa-go/goa)   [![godoc][GoDoc]](https://godoc.org/github.com/goa-go/goa)
* [Resoursea](https://github.com/resoursea/api) **star:31** REST framework for quickly writing resource based services.   [![It hasn't been updated in the last year][Yellow]](https://github.com/resoursea/api)   [![godoc][GoDoc]](https://godoc.org/github.com/resoursea/api)
* [REST Layer](http://rest-layer.io)  Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [rex](https://github.com/goanywhere/rex) **star:31** Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goanywhere/rex)   [![godoc][GoDoc]](https://godoc.org/github.com/goanywhere/rex)
* [goweb](https://github.com/twharmon/goweb) **star:20** Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS.   [![godoc][GoDoc]](https://godoc.org/github.com/twharmon/goweb)
* [Banjo](https://github.com/nsheremet/banjo) **star:17** Very simple and fast web framework for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nsheremet/banjo)   [![godoc][GoDoc]](https://godoc.org/github.com/nsheremet/banjo)

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) **star:1842** Rate limit HTTP request handler.   [![godoc][GoDoc]](https://godoc.org/github.com/didip/tollbooth)
* [CORS](https://github.com/rs/cors) **star:1703** Easily add CORS capabilities to your API.   [![godoc][GoDoc]](https://godoc.org/github.com/rs/cors)
* [Limiter](https://github.com/ulule/limiter) **star:1212** Dead simple rate limit middleware for Go.   [![There was an update last week][Green]](https://github.com/ulule/limiter)   [![godoc][GoDoc]](https://godoc.org/github.com/ulule/limiter)
* [go-server-timing](https://github.com/mitchellh/go-server-timing) **star:791** Add/parse Server-Timing header.   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/go-server-timing)
* [go-fault](https://github.com/github/go-fault) **star:383** Fault injection middleware for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/github/go-fault)
* [ln-paywall](https://github.com/philippgille/ln-paywall) **star:109** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).   [![It hasn't been updated in the last year][Yellow]](https://github.com/philippgille/ln-paywall)   [![godoc][GoDoc]](https://godoc.org/github.com/philippgille/ln-paywall)
* [XFF](https://github.com/sebest/xff) **star:77** Handle `X-Forwarded-For` header and friends.   [![godoc][GoDoc]](https://godoc.org/github.com/sebest/xff)
* [formjson](https://github.com/rs/formjson) **star:36** Transparently handle JSON input as a standard form POST.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/formjson)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/formjson)
* [client-timing](https://github.com/posener/client-timing) **star:18** An HTTP client for Server-Timing header.   [![godoc][GoDoc]](https://godoc.org/github.com/posener/client-timing)

#### Libraries for creating HTTP middlewares

* [negroni](https://github.com/urfave/negroni) **star:6887** Idiomatic HTTP middleware for Golang.   [![star > 2000][Awesome]](https://github.com/urfave/negroni)   [![There was an update last week][Green]](https://github.com/urfave/negroni)   [![godoc][GoDoc]](https://godoc.org/github.com/urfave/negroni)   [![Contains Chinese documents][CN]](https://github.com/urfave/negroni)
* [alice](https://github.com/justinas/alice) **star:2191** Painless middleware chaining for Go.   [![star > 2000][Awesome]](https://github.com/justinas/alice)   [![There was an update last week][Green]](https://github.com/justinas/alice)   [![godoc][GoDoc]](https://godoc.org/github.com/justinas/alice)
* [render](https://github.com/unrolled/render) **star:1404** Go package for easily rendering JSON, XML, and HTML template responses.   [![godoc][GoDoc]](https://godoc.org/github.com/unrolled/render)
* [stats](https://github.com/thoas/stats) **star:574** Go middleware that stores various information about your web application.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thoas/stats)   [![godoc][GoDoc]](https://godoc.org/github.com/thoas/stats)
* [interpose](https://github.com/carbocation/interpose) **star:291** Minimalist net/http middleware for golang.   [![It hasn't been updated in the last year][Yellow]](https://github.com/carbocation/interpose)   [![godoc][GoDoc]](https://godoc.org/github.com/carbocation/interpose)
* [renderer](https://github.com/thedevsaddam/renderer) **star:217** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsaddam/renderer)
* [muxchain](https://github.com/stephens2424/muxchain) **star:209** Lightweight middleware for net/http.   [![It hasn't been updated in the last year][Yellow]](https://github.com/stephens2424/muxchain)   [![godoc][GoDoc]](https://godoc.org/github.com/stephens2424/muxchain)
* [rye](https://github.com/InVisionApp/rye) **star:95** Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.   [![It hasn't been updated in the last year][Yellow]](https://github.com/InVisionApp/rye)   [![godoc][GoDoc]](https://godoc.org/github.com/InVisionApp/rye)
* [gores](https://github.com/alioygur/gores) **star:94** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/alioygur/gores)
* [mediary](https://github.com/HereMobilityDevelopers/mediary) **star:69** add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses.   [![godoc][GoDoc]](https://godoc.org/github.com/HereMobilityDevelopers/mediary)
* [chain](https://github.com/codemodus/chain) **star:64** Handler wrapper chaining with scoped data (net/context-based "middleware").   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/chain)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/chain)
* [go-wrap](https://github.com/go-on/wrap) **star:59** Small middlewares package for net/http.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-on/wrap)   [![godoc][GoDoc]](https://godoc.org/github.com/go-on/wrap)   [![Archived][Archived]](https://github.com/go-on/wrap)
* [catena](https://github.com/codemodus/catena) **star:7** http.Handler wrapper catenation (same API as "chain").   [![It hasn't been updated in the last year][Yellow]](https://github.com/codemodus/catena)   [![godoc][GoDoc]](https://godoc.org/github.com/codemodus/catena)

### Routers

* [mux](https://github.com/gorilla/mux) **star:13586** Powerful URL router and dispatcher for golang.   [![star > 2000][Awesome]](https://github.com/gorilla/mux)   [![There was an update last week][Green]](https://github.com/gorilla/mux)   [![godoc][GoDoc]](https://godoc.org/github.com/gorilla/mux)
* [httprouter](https://github.com/julienschmidt/httprouter) **star:12274** High performance router. Use this and the standard http handlers to form a very high performance web framework.   [![star > 2000][Awesome]](https://github.com/julienschmidt/httprouter)   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/httprouter)
* [chi](https://github.com/go-chi/chi) **star:8770** Small, fast and expressive HTTP router built on net/context.   [![star > 2000][Awesome]](https://github.com/go-chi/chi)   [![godoc][GoDoc]](https://godoc.org/github.com/go-chi/chi)
* [gocraft/web](https://github.com/gocraft/web) **star:1438** Mux and middleware package in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/gocraft/web)
* [Bone](https://github.com/go-zoo/bone) **star:1264** Lightning Fast HTTP Multiplexer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-zoo/bone)   [![godoc][GoDoc]](https://godoc.org/github.com/go-zoo/bone)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:868** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/buaazp/fasthttprouter)   [![godoc][GoDoc]](https://godoc.org/github.com/buaazp/fasthttprouter)
* [Goji](https://github.com/goji/goji) **star:843** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goji/goji)   [![godoc][GoDoc]](https://godoc.org/github.com/goji/goji)
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:491** A simple and fast HTTP router for Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xujiajun/gorouter)   [![godoc][GoDoc]](https://godoc.org/github.com/xujiajun/gorouter)
* [httptreemux](https://github.com/dimfeld/httptreemux) **star:472** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   [![godoc][GoDoc]](https://godoc.org/github.com/dimfeld/httptreemux)
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:391** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
* [lars](https://github.com/go-playground/lars) **star:382** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-playground/lars)   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/lars)
* [Siesta](https://github.com/VividCortex/siesta) **star:350** Composable framework to write middleware and handlers.   [![godoc][GoDoc]](https://godoc.org/github.com/VividCortex/siesta)
* [vestigo](https://github.com/husobee/vestigo) **star:265** Performant, stand-alone, HTTP compliant URL Router for go web applications.   [![godoc][GoDoc]](https://godoc.org/github.com/husobee/vestigo)
* [gowww/router](https://github.com/gowww/router) **star:157** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   [![godoc][GoDoc]](https://godoc.org/github.com/gowww/router)
* [alien](https://github.com/gernest/alien) **star:115** Lightweight and fast http router from outer space.   [![It hasn't been updated in the last year][Yellow]](https://github.com/gernest/alien)   [![godoc][GoDoc]](https://godoc.org/github.com/gernest/alien)
* [pure](https://github.com/go-playground/pure) **star:111** Is a lightweight HTTP router that sticks to the std "net/http" implementation.   [![godoc][GoDoc]](https://godoc.org/github.com/go-playground/pure)
* [violetear](https://github.com/nbari/violetear) **star:100** Go HTTP router.   [![It hasn't been updated in the last year][Yellow]](https://github.com/nbari/violetear)   [![godoc][GoDoc]](https://godoc.org/github.com/nbari/violetear)
* [Bxog](https://github.com/claygod/Bxog) **star:98** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.   [![godoc][GoDoc]](https://godoc.org/github.com/claygod/Bxog)
* [GoRouter](https://github.com/vardius/gorouter) **star:88** GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.   [![godoc][GoDoc]](https://godoc.org/github.com/vardius/gorouter)
* [xmux](https://github.com/rs/xmux) **star:88** High performance muxer based on `httprouter` with `net/context` support.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rs/xmux)   [![godoc][GoDoc]](https://godoc.org/github.com/rs/xmux)
* [bellt](https://github.com/GuilhermeCaruso/bellt) **star:48** A simple Go HTTP router.   [![godoc][GoDoc]](https://godoc.org/github.com/GuilhermeCaruso/bellt)
* [FastRouter](https://github.com/razonyang/fastrouter) **star:19** a fast, flexible HTTP router written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/razonyang/fastrouter)   [![godoc][GoDoc]](https://godoc.org/github.com/razonyang/fastrouter)
* [goroute](https://github.com/goroute/route) **star:7** Simple yet powerful HTTP request multiplexer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/goroute/route)   [![godoc][GoDoc]](https://godoc.org/github.com/goroute/route)

## WebAssembly

* [tinygo](https://github.com/tinygo-org/tinygo) **star:7392** Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.   [![star > 2000][Awesome]](https://github.com/tinygo-org/tinygo)   [![There was an update last week][Green]](https://github.com/tinygo-org/tinygo)   [![godoc][GoDoc]](https://godoc.org/github.com/tinygo-org/tinygo)
* [dom](https://github.com/dennwc/dom) **star:422** DOM library.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dennwc/dom)   [![godoc][GoDoc]](https://godoc.org/github.com/dennwc/dom)
* [go-canvas](https://github.com/markfarnan/go-canvas) **star:111** Library to use HTML5 Canvas, with all drawing within go code.   [![godoc][GoDoc]](https://godoc.org/github.com/markfarnan/go-canvas)
* [webapi](https://github.com/gowebapi/webapi) **star:75** Bindings for DOM and HTML generated from WebIDL.   [![There was an update last week][Green]](https://github.com/gowebapi/webapi)   [![godoc][GoDoc]](https://godoc.org/github.com/gowebapi/webapi)
* [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) **star:71** Run Go WASM tests in your browser.   [![godoc][GoDoc]](https://godoc.org/github.com/agnivade/wasmbrowsertest)
* [vert](https://github.com/norunners/vert) **star:45** Interop between Go and JS values.   [![godoc][GoDoc]](https://godoc.org/github.com/norunners/vert)

## Windows

* [go-ole](https://github.com/go-ole/go-ole) **star:710** Win32 OLE implementation for golang.   [![godoc][GoDoc]](https://godoc.org/github.com/go-ole/go-ole)
* [d3d9](https://github.com/gonutz/d3d9) **star:115** Go bindings for Direct3D9.   [![godoc][GoDoc]](https://godoc.org/github.com/gonutz/d3d9)
* [gosddl](https://github.com/MonaxGT/gosddl) **star:4** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.   [![It hasn't been updated in the last year][Yellow]](https://github.com/MonaxGT/gosddl)   [![godoc][GoDoc]](https://godoc.org/github.com/MonaxGT/gosddl)

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) **star:442** Generate a Go struct from XML.   [![godoc][GoDoc]](https://godoc.org/github.com/miku/zek)
* [xpath](https://github.com/antchfx/xpath) **star:367** XPath package for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xpath)
* [xquery](https://github.com/antchfx/xquery) **star:155** XQuery lets you extract data from HTML/XML documents using XPath expression.   [![It hasn't been updated in the last year][Yellow]](https://github.com/antchfx/xquery)   [![godoc][GoDoc]](https://godoc.org/github.com/antchfx/xquery)   [![Archived][Archived]](https://github.com/antchfx/xquery)
* [xml2map](https://github.com/sbabiv/xml2map) **star:27** XML to MAP converter written Golang.   [![godoc][GoDoc]](https://godoc.org/github.com/sbabiv/xml2map)
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:17** Procedural XML generation API based on libxml2's xmlwriter module.   [![godoc][GoDoc]](https://godoc.org/github.com/shabbyrobe/xmlwriter)
* [XML-Comp](https://github.com/xml-comp/xml-comp) **star:16** Simple command line XML comparer that generates diffs of folders, files and tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xml-comp/xml-comp)   [![godoc][GoDoc]](https://godoc.org/github.com/xml-comp/xml-comp)

# Tools

*Go software and plugins.*

## Code Analysis

* [GoLint](https://github.com/golang/lint) **star:3797** Golint is a linter for Go source code.   [![star > 2000][Awesome]](https://github.com/golang/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/golang/lint)
* [Golint online](http://go-lint.appspot.com/)  Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [errcheck](https://github.com/kisielk/errcheck) **star:1637** Errcheck is a program for checking for unchecked errors in Go programs.   [![godoc][GoDoc]](https://godoc.org/github.com/kisielk/errcheck)
* [gcvis](https://github.com/davecheney/gcvis) **star:1003** Visualise Go program GC trace data in real time.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davecheney/gcvis)   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/gcvis)
* [go-critic](https://github.com/go-critic/go-critic) **star:845** source code linter that brings checks that are currently not implemented in other linters.   [![There was an update last week][Green]](https://github.com/go-critic/go-critic)   [![godoc][GoDoc]](https://godoc.org/github.com/go-critic/go-critic)
* [php-parser](https://github.com/z7zmey/php-parser) **star:782** A Parser for PHP written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/z7zmey/php-parser)
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:488** Web based Golang AST visualizer.   [![It hasn't been updated in the last year][Yellow]](https://github.com/yuroyoro/goast-viewer)
* [GoCover.io](http://gocover.io/)  GoCover.io offers the code coverage of any golang package as a service.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:457** An easy way to find outdated dependencies of your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/go-mod-outdated)
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:410** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   [![godoc][GoDoc]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) **star:315** Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.   [![There was an update last week][Green]](https://github.com/jfeliu007/goplantuml)   [![godoc][GoDoc]](https://godoc.org/github.com/jfeliu007/goplantuml)
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns)  Adds zero-value return statements to match the func return types.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)  gosimple is a linter for Go source code that specialises on simplifying code.
* [unconvert](https://github.com/mdempsky/unconvert) **star:306** Remove unnecessary type conversions from Go source.   [![godoc][GoDoc]](https://godoc.org/github.com/mdempsky/unconvert)
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused)  unused checks Go code for unused constants, variables, functions and types.
* [gostatus](https://github.com/shurcooL/gostatus) **star:244** Command line tool, shows the status of repositories that contain Go packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/shurcooL/gostatus)   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/gostatus)
* [tickgit](https://github.com/augmentable-dev/tickgit) **star:241** CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.   [![godoc][GoDoc]](https://godoc.org/github.com/augmentable-dev/tickgit)
* [dupl](https://github.com/mibk/dupl) **star:223** Tool for code clone detection.   [![godoc][GoDoc]](https://godoc.org/github.com/mibk/dupl)
* [apicompat](https://github.com/bradleyfalzon/apicompat) **star:170** Checks recent changes to a Go project for backwards incompatible changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/bradleyfalzon/apicompat)   [![godoc][GoDoc]](https://godoc.org/github.com/bradleyfalzon/apicompat)
* [go-checkstyle](https://github.com/qiniu/checkstyle) **star:111** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.   [![godoc][GoDoc]](https://godoc.org/github.com/qiniu/checkstyle)
* [golines](https://github.com/segmentio/golines) **star:92** Formatter that automatically shortens long lines in Go code.   [![godoc][GoDoc]](https://godoc.org/github.com/segmentio/golines)
* [lint](https://github.com/surullabs/lint) **star:66** Run linters as part of go test.   [![It hasn't been updated in the last year][Yellow]](https://github.com/surullabs/lint)   [![godoc][GoDoc]](https://godoc.org/github.com/surullabs/lint)
* [validate](https://github.com/mccoyst/validate) **star:60** Automatically validates struct fields with tags.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mccoyst/validate)   [![godoc][GoDoc]](https://godoc.org/github.com/mccoyst/validate)
* [go-outdated](https://github.com/firstrow/go-outdated) **star:44** Console application that displays outdated packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/firstrow/go-outdated)   [![godoc][GoDoc]](https://godoc.org/github.com/firstrow/go-outdated)   [![Archived][Archived]](https://github.com/firstrow/go-outdated)
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) **star:14** tarp finds functions and methods without direct unit tests in Go source code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/verygoodsoftwarenotvirus/tarp)   [![godoc][GoDoc]](https://godoc.org/github.com/verygoodsoftwarenotvirus/tarp)   [![Archived][Archived]](https://github.com/verygoodsoftwarenotvirus/tarp)

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) **star:13057** Go development plugin for Vim.   [![star > 2000][Awesome]](https://github.com/fatih/vim-go)   [![There was an update last week][Green]](https://github.com/fatih/vim-go)
* [gocode](https://github.com/nsf/gocode) **star:4907** Autocompletion daemon for the Go programming language.   [![star > 2000][Awesome]](https://github.com/nsf/gocode)   [![godoc][GoDoc]](https://godoc.org/github.com/nsf/gocode)
* [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3392** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.   [![star > 2000][Awesome]](https://github.com/DisposaBoy/GoSublime)   [![godoc][GoDoc]](https://godoc.org/github.com/DisposaBoy/GoSublime)
* [go-plus](https://github.com/joefitzgerald/go-plus) **star:1515** Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.   [![There was an update last week][Green]](https://github.com/joefitzgerald/go-plus)
* [vscode-go](https://github.com/golang/vscode-go) **star:1347** Extension for Visual Studio Code (VS Code) which provides support for the Go language.   [![There was an update last week][Green]](https://github.com/golang/vscode-go)
* [go-mode](https://github.com/dominikh/go-mode.el) **star:1136** Go mode for GNU/Emacs.
* [Watch](https://github.com/eaburns/Watch) **star:178** Runs a command in an acme win on file changes.   [![It hasn't been updated in the last year][Yellow]](https://github.com/eaburns/Watch)   [![godoc][GoDoc]](https://godoc.org/github.com/eaburns/Watch)
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:88** Vim plugin to highlight syntax errors on save.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rjohnsondev/vim-compiler-go)
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.
* [go-language-server](https://github.com/theia-ide/go-language-server) **star:33** A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.   [![It hasn't been updated in the last year][Yellow]](https://github.com/theia-ide/go-language-server)
* [goimports-reviser](https://github.com/incu6us/goimports-reviser) **star:24** Formatting tool for imports.   [![godoc][GoDoc]](https://godoc.org/github.com/incu6us/goimports-reviser)
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:19** Vim plugin for generating Go tests based on the function's or method's signature.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/gounit-vim)
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) **star:15** Go language support for the Theia IDE.   [![It hasn't been updated in the last year][Yellow]](https://github.com/theia-ide/theia-go-extension)

## Go Generate Tools

* [gotests](https://github.com/cweill/gotests) **star:3006** Generate Go tests from your source code.   [![star > 2000][Awesome]](https://github.com/cweill/gotests)   [![godoc][GoDoc]](https://godoc.org/github.com/cweill/gotests)
* [genny](https://github.com/cheekybits/genny) **star:1383** Elegant generics for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/cheekybits/genny)
* [re2dfa](https://github.com/opennota/re2dfa) **star:181** Transform regular expressions into finite state machines and output Go source code.   [![It hasn't been updated in the last year][Yellow]](https://github.com/opennota/re2dfa)   [![godoc][GoDoc]](https://godoc.org/github.com/opennota/re2dfa)
* [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.
* [hasgo](https://github.com/DylanMeeus/hasgo) **star:79** Generate Haskell inspired functions for your slices.   [![godoc][GoDoc]](https://godoc.org/github.com/DylanMeeus/hasgo)
* [gocontracts](https://github.com/Parquery/gocontracts) **star:65** brings design-by-contract to Go by synchronizing the code with the documentation.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Parquery/gocontracts)   [![godoc][GoDoc]](https://godoc.org/github.com/Parquery/gocontracts)
* [gonerics](http://github.com/bouk/gonerics)  Idiomatic Generics in Go.
* [xgen](https://github.com/xuri/xgen) **star:65** XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator.   [![There was an update last week][Green]](https://github.com/xuri/xgen)   [![godoc][GoDoc]](https://godoc.org/github.com/xuri/xgen)
* [gounit](https://github.com/hexdigest/gounit) **star:45** Generate Go tests using your own templates.   [![It hasn't been updated in the last year][Yellow]](https://github.com/hexdigest/gounit)   [![godoc][GoDoc]](https://godoc.org/github.com/hexdigest/gounit)
* [generic](https://github.com/usk81/generic) **star:35** flexible data type for Go.   [![godoc][GoDoc]](https://godoc.org/github.com/usk81/generic)

## Go Tools

* [go-swagger](https://github.com/go-swagger/go-swagger) **star:6039** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.   [![star > 2000][Awesome]](https://github.com/go-swagger/go-swagger)   [![There was an update last week][Green]](https://github.com/go-swagger/go-swagger)   [![godoc][GoDoc]](https://godoc.org/github.com/go-swagger/go-swagger)
* [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:4613** Navigate through go files efficiently with the OctoLinker browser extension for GitHub.   [![star > 2000][Awesome]](https://github.com/OctoLinker/browser-extension)   [![There was an update last week][Green]](https://github.com/OctoLinker/browser-extension)
* [go-callvis](https://github.com/TrueFurby/go-callvis) **star:3011** Visualize call graph of your Go program using dot format.   [![star > 2000][Awesome]](https://github.com/TrueFurby/go-callvis)   [![godoc][GoDoc]](https://godoc.org/github.com/TrueFurby/go-callvis)
* [depth](https://github.com/KyleBanks/depth) **star:577** Visualize dependency trees of any package by analyzing imports.   [![godoc][GoDoc]](https://godoc.org/github.com/KyleBanks/depth)
* [richgo](https://github.com/kyoh86/richgo) **star:526** Enrich `go test` outputs with text decorations.   [![There was an update last week][Green]](https://github.com/kyoh86/richgo)   [![godoc][GoDoc]](https://godoc.org/github.com/kyoh86/richgo)
* [rts](https://github.com/galeone/rts) **star:207** RTS: response to struct. Generates Go structs from server responses.   [![godoc][GoDoc]](https://godoc.org/github.com/galeone/rts)
* [godbg](https://github.com/tylerwince/godbg) **star:172** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylerwince/godbg)   [![godoc][GoDoc]](https://godoc.org/github.com/tylerwince/godbg)
* [gomodrun](https://github.com/dustinblackman/gomodrun/)  Go tool that executes and caches binaries included in go.mod files.
* [typex](https://github.com/dtgorski/typex) **star:128** Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration.   [![godoc][GoDoc]](https://godoc.org/github.com/dtgorski/typex)
* [colorgo](https://github.com/songgao/colorgo) **star:108** Wrapper around `go` command for colorized `go build` output.   [![godoc][GoDoc]](https://godoc.org/github.com/songgao/colorgo)
* [gothanks](https://github.com/psampaz/gothanks) **star:96** GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.   [![godoc][GoDoc]](https://godoc.org/github.com/psampaz/gothanks)
* [igo](https://github.com/rocketlaunchr/igo) **star:40** An igo to go transpiler (new language features for Go language!)   [![godoc][GoDoc]](https://godoc.org/github.com/rocketlaunchr/igo)
* [go-james](https://github.com/pieterclaerhout/go-james) **star:38** Go project skeleton creator, builds and tests your projects without the manual setup.   [![godoc][GoDoc]](https://godoc.org/github.com/pieterclaerhout/go-james)
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) **star:37** Bash completion for go and wgo.   [![It hasn't been updated in the last year][Yellow]](https://github.com/skelterjohn/go-pkg-complete)
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) **star:22** A [Yeoman](http://yeoman.io) generator to get new Go projects started.   [![Archived][Archived]](https://github.com/axelspringer/generator-go-lang)

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) **star:73842** Container Cluster Manager from Google.   [![star > 2000][Awesome]](https://github.com/kubernetes/kubernetes)   [![There was an update last week][Green]](https://github.com/kubernetes/kubernetes)   [![godoc][GoDoc]](https://godoc.org/github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby) **star:59492** Collaborative project for the container ecosystem to assemble container-based systems.   [![star > 2000][Awesome]](https://github.com/moby/moby)   [![There was an update last week][Green]](https://github.com/moby/moby)   [![godoc][GoDoc]](https://godoc.org/github.com/moby/moby)
* [traefik](https://github.com/containous/traefik) **star:32346** Reverse proxy and load balancer with support for multiple backends.   [![star > 2000][Awesome]](https://github.com/containous/traefik)   [![There was an update last week][Green]](https://github.com/containous/traefik)   [![godoc][GoDoc]](https://godoc.org/github.com/containous/traefik)
* [Gitea](https://github.com/go-gitea/gitea) **star:23286** Fork of Gogs, entirely community driven.   [![star > 2000][Awesome]](https://github.com/go-gitea/gitea)   [![There was an update last week][Green]](https://github.com/go-gitea/gitea)   [![godoc][GoDoc]](https://godoc.org/github.com/go-gitea/gitea)   [![Contains Chinese documents][CN]](https://github.com/go-gitea/gitea)
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
* [Vegeta](https://github.com/tsenart/vegeta) **star:16223** HTTP load testing tool and library. It's over 9000!   [![star > 2000][Awesome]](https://github.com/tsenart/vegeta)   [![There was an update last week][Green]](https://github.com/tsenart/vegeta)   [![godoc][GoDoc]](https://godoc.org/github.com/tsenart/vegeta)
* [Packer](https://github.com/mitchellh/packer) **star:12496** Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.   [![star > 2000][Awesome]](https://github.com/mitchellh/packer)   [![There was an update last week][Green]](https://github.com/mitchellh/packer)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/packer)
* [Hey](https://github.com/rakyll/hey) **star:10149** Hey is a tiny program that sends some load to a web application.   [![star > 2000][Awesome]](https://github.com/rakyll/hey)   [![godoc][GoDoc]](https://godoc.org/github.com/rakyll/hey)
* [webhook](https://github.com/adnanh/webhook) **star:6145** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   [![star > 2000][Awesome]](https://github.com/adnanh/webhook)   [![There was an update last week][Green]](https://github.com/adnanh/webhook)   [![godoc][GoDoc]](https://godoc.org/github.com/adnanh/webhook)
* [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
* [GVM](https://github.com/moovweb/gvm) **star:5931** GVM provides an interface to manage Go versions.   [![star > 2000][Awesome]](https://github.com/moovweb/gvm)
* [gaia](https://github.com/gaia-pipeline/gaia) **star:4309** Build powerful pipelines in any programming language.   [![star > 2000][Awesome]](https://github.com/gaia-pipeline/gaia)   [![godoc][GoDoc]](https://godoc.org/github.com/gaia-pipeline/gaia)
* [gox](https://github.com/mitchellh/gox) **star:3959** Dead simple, no frills Go cross compile tool.   [![star > 2000][Awesome]](https://github.com/mitchellh/gox)   [![godoc][GoDoc]](https://godoc.org/github.com/mitchellh/gox)
* [bosun](https://github.com/bosun-monitor/bosun) **star:3094** Time Series Alerting Framework.   [![star > 2000][Awesome]](https://github.com/bosun-monitor/bosun)   [![godoc][GoDoc]](https://godoc.org/github.com/bosun-monitor/bosun)
* [bombardier](https://github.com/codesenberg/bombardier) **star:2462** Fast cross-platform HTTP benchmarking tool.   [![star > 2000][Awesome]](https://github.com/codesenberg/bombardier)   [![godoc][GoDoc]](https://godoc.org/github.com/codesenberg/bombardier)
* [Pomerium](https://github.com/pomerium/pomerium) **star:2285** Pomerium is an identity-aware access proxy.   [![star > 2000][Awesome]](https://github.com/pomerium/pomerium)   [![There was an update last week][Green]](https://github.com/pomerium/pomerium)   [![godoc][GoDoc]](https://godoc.org/github.com/pomerium/pomerium)
* [fac](https://github.com/mkchoi212/fac) **star:1694** Command-line user interface to fix git merge conflicts.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mkchoi212/fac)   [![godoc][GoDoc]](https://godoc.org/github.com/mkchoi212/fac)
* [script](https://github.com/bitfield/script) **star:1687** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   [![There was an update last week][Green]](https://github.com/bitfield/script)   [![godoc][GoDoc]](https://godoc.org/github.com/bitfield/script)
* [goxc](https://github.com/laher/goxc) **star:1662** build tool for Go, with a focus on cross-compiling and packaging.   [![It hasn't been updated in the last year][Yellow]](https://github.com/laher/goxc)   [![godoc][GoDoc]](https://godoc.org/github.com/laher/goxc)
* [kala](https://github.com/ajvb/kala) **star:1609** Simplistic, modern, and performant job scheduler.   [![There was an update last week][Green]](https://github.com/ajvb/kala)   [![godoc][GoDoc]](https://godoc.org/github.com/ajvb/kala)
* [StatusOK](https://github.com/sanathp/statusok) **star:1446** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.   [![There was an update last week][Green]](https://github.com/sanathp/statusok)   [![godoc][GoDoc]](https://godoc.org/github.com/sanathp/statusok)
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:1082** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.   [![godoc][GoDoc]](https://godoc.org/github.com/rlmcpherson/s3gof3r)
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:793** Enable your Go applications to self update.   [![godoc][GoDoc]](https://godoc.org/github.com/sanbornm/go-selfupdate)
* [skm](https://github.com/TimothyYe/skm) **star:667** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   [![godoc][GoDoc]](https://godoc.org/github.com/TimothyYe/skm)
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:647** Manage BareMetal Servers from Command Line (as easily as with Docker).   [![There was an update last week][Green]](https://github.com/scaleway/scaleway-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/scaleway/scaleway-cli)
* [s5cmd](https://github.com/peak/s5cmd) **star:522** Blazing fast S3 and local filesystem execution tool.   [![There was an update last week][Green]](https://github.com/peak/s5cmd)   [![godoc][GoDoc]](https://godoc.org/github.com/peak/s5cmd)
* [aurora](https://github.com/xuri/aurora) **star:499** Cross-platform web-based Beanstalkd queue server console.
* [cassowary](https://github.com/rogerwelin/cassowary) **star:468** Modern cross-platform HTTP load-testing tool written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/rogerwelin/cassowary)   [![Contains Chinese documents][CN]](https://github.com/rogerwelin/cassowary)
* [govvv](https://github.com/ahmetalpbalkan/govvv) **star:468** “go build” wrapper to easily add version information into Go binaries.   [![godoc][GoDoc]](https://godoc.org/github.com/ahmetalpbalkan/govvv)
* [uTask](https://github.com/ovh/utask) **star:352** Automation engine that models and executes business processes declared in yaml.   [![There was an update last week][Green]](https://github.com/ovh/utask)   [![godoc][GoDoc]](https://godoc.org/github.com/ovh/utask)
* [gonative](https://github.com/inconshreveable/gonative) **star:323** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/inconshreveable/gonative)   [![godoc][GoDoc]](https://godoc.org/github.com/inconshreveable/gonative)
* [trubka](https://github.com/xitonix/trubka) **star:294** A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka.   [![godoc][GoDoc]](https://godoc.org/github.com/xitonix/trubka)
* [Mora](https://github.com/emicklei/mora) **star:282** REST server for accessing MongoDB documents and meta data.   [![It hasn't been updated in the last year][Yellow]](https://github.com/emicklei/mora)   [![godoc][GoDoc]](https://godoc.org/github.com/emicklei/mora)
* [lstags](https://github.com/ivanilves/lstags) **star:274** Tool and API to sync Docker images across different registries.   [![godoc][GoDoc]](https://godoc.org/github.com/ivanilves/lstags)
* [Pewpew](https://github.com/bengadbois/pewpew) **star:263** Flexible HTTP command line stress tester.   [![There was an update last week][Green]](https://github.com/bengadbois/pewpew)   [![godoc][GoDoc]](https://godoc.org/github.com/bengadbois/pewpew)
* [jcli](https://github.com/jenkins-zh/jenkins-cli) **star:233** Jenkins CLI allows you manage your Jenkins as an easy way.   [![There was an update last week][Green]](https://github.com/jenkins-zh/jenkins-cli)   [![godoc][GoDoc]](https://godoc.org/github.com/jenkins-zh/jenkins-cli)   [![Contains Chinese documents][CN]](https://github.com/jenkins-zh/jenkins-cli)
* [dogo](https://github.com/liudng/dogo) **star:232** Monitoring changes in the source file and automatically compile and run (restart).   [![It hasn't been updated in the last year][Yellow]](https://github.com/liudng/dogo)   [![godoc][GoDoc]](https://godoc.org/github.com/liudng/dogo)   [![Contains Chinese documents][CN]](https://github.com/liudng/dogo)
* [godbg](https://github.com/sirnewton01/godbg) **star:225** Web-based gdb front-end application.   [![It hasn't been updated in the last year][Yellow]](https://github.com/sirnewton01/godbg)
* [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
* [manssh](https://github.com/xwjdsh/manssh) **star:222** manssh is a command line tool for managing your ssh alias config easily.   [![It hasn't been updated in the last year][Yellow]](https://github.com/xwjdsh/manssh)   [![godoc][GoDoc]](https://godoc.org/github.com/xwjdsh/manssh)
* [Blast](https://github.com/dave/blast) **star:192** A simple tool for API load testing and batch jobs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/dave/blast)   [![godoc][GoDoc]](https://godoc.org/github.com/dave/blast)
* [gobrew](https://github.com/cryptojuice/gobrew) **star:177** gobrew lets you easily switch between multiple versions of go.
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:170** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![There was an update last week][Green]](https://github.com/appleboy/easyssh-proxy)   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/easyssh-proxy)
* [ostent](https://github.com/ostrost/ostent) **star:170** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ostrost/ostent)   [![godoc][GoDoc]](https://godoc.org/github.com/ostrost/ostent)
* [abbreviate](https://github.com/dnnrly/abbreviate) **star:160** abbreviate is a tool turning long strings in to shorter ones with configurable seperaters, for example to embed branch names in to deployment stack IDs.   [![godoc][GoDoc]](https://godoc.org/github.com/dnnrly/abbreviate)
* [grapes](https://github.com/yaronsumel/grapes) **star:153** Lightweight tool designed to distribute commands over ssh with ease.   [![godoc][GoDoc]](https://godoc.org/github.com/yaronsumel/grapes)
* [kcli](https://github.com/cswank/kcli) **star:143** Command line tool for inspecting kafka topics/partitions/messages.   [![It hasn't been updated in the last year][Yellow]](https://github.com/cswank/kcli)   [![godoc][GoDoc]](https://godoc.org/github.com/cswank/kcli)
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) **star:132** Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.   [![godoc][GoDoc]](https://godoc.org/github.com/dikhan/terraform-provider-openapi)
* [winrm-cli](https://github.com/masterzen/winrm-cli) **star:108** Cli tool to remotely execute commands on Windows machines.   [![godoc][GoDoc]](https://godoc.org/github.com/masterzen/winrm-cli)
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) **star:92** A go library and an executable that produces valid Dockerfiles using various input channels.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ozankasikci/dockerfile-generator)   [![godoc][GoDoc]](https://godoc.org/github.com/ozankasikci/dockerfile-generator)
* [drone-scp](https://github.com/appleboy/drone-scp) **star:79** Copy files and artifacts via SSH using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-scp)
* [go-furnace](https://github.com/go-furnace/go-furnace) **star:79** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   [![It hasn't been updated in the last year][Yellow]](https://github.com/go-furnace/go-furnace)   [![godoc][GoDoc]](https://godoc.org/github.com/go-furnace/go-furnace)
* [Dropship](https://github.com/chrismckenzie/dropship) **star:54** Tool for deploying code via cdn.   [![It hasn't been updated in the last year][Yellow]](https://github.com/chrismckenzie/dropship)   [![godoc][GoDoc]](https://godoc.org/github.com/chrismckenzie/dropship)
* [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) **star:32** S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).   [![There was an update last week][Green]](https://github.com/oxyno-zeta/s3-proxy)   [![godoc][GoDoc]](https://godoc.org/github.com/oxyno-zeta/s3-proxy)
* [Rodent](https://github.com/alouche/rodent) **star:31** Rodent helps you manage Go versions, projects and track dependencies.   [![It hasn't been updated in the last year][Yellow]](https://github.com/alouche/rodent)
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:30** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.   [![godoc][GoDoc]](https://godoc.org/github.com/appleboy/drone-jenkins)
* [awsenv](https://github.com/soniah/awsenv) **star:26** Small binary that loads Amazon (AWS) environment variables for a profile.   [![It hasn't been updated in the last year][Yellow]](https://github.com/soniah/awsenv)   [![godoc][GoDoc]](https://godoc.org/github.com/soniah/awsenv)
* [lwc](https://github.com/timdp/lwc) **star:24** A live-updating version of the UNIX wc command.   [![godoc][GoDoc]](https://godoc.org/github.com/timdp/lwc)
* [DepCharge](https://github.com/centerorbit/depcharge) **star:14** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![godoc][GoDoc]](https://godoc.org/github.com/centerorbit/depcharge)
* [httpref](https://github.com/dnnrly/httpref) **star:13** httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports.   [![There was an update last week][Green]](https://github.com/dnnrly/httpref)   [![godoc][GoDoc]](https://godoc.org/github.com/dnnrly/httpref)
* [sg](https://github.com/ChristopherRabotin/sg) **star:6** Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ChristopherRabotin/sg)   [![godoc][GoDoc]](https://godoc.org/github.com/ChristopherRabotin/sg)
* [aptly](https://github.com/smira/aptly) **star:3** aptly is a Debian repository management tool.   [![It hasn't been updated in the last year][Yellow]](https://github.com/smira/aptly)   [![godoc][GoDoc]](https://godoc.org/github.com/smira/aptly)

### Other Software

* [Gor](https://github.com/buger/gor) **star:13802** Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.   [![star > 2000][Awesome]](https://github.com/buger/gor)   [![There was an update last week][Green]](https://github.com/buger/gor)   [![godoc][GoDoc]](https://godoc.org/github.com/buger/gor)
* [restic](https://github.com/restic/restic) **star:11761** De-duplicating backup program.   [![star > 2000][Awesome]](https://github.com/restic/restic)   [![There was an update last week][Green]](https://github.com/restic/restic)   [![godoc][GoDoc]](https://godoc.org/github.com/restic/restic)
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:11356** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   [![star > 2000][Awesome]](https://github.com/chrislusf/seaweedfs)   [![There was an update last week][Green]](https://github.com/chrislusf/seaweedfs)   [![godoc][GoDoc]](https://godoc.org/github.com/chrislusf/seaweedfs)
* [croc](https://github.com/schollz/croc) **star:11336** Easily and securely send files or folders from one computer to another.   [![star > 2000][Awesome]](https://github.com/schollz/croc)   [![There was an update last week][Green]](https://github.com/schollz/croc)   [![godoc][GoDoc]](https://godoc.org/github.com/schollz/croc)
* [Docker](http://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
* [confd](https://github.com/kelseyhightower/confd) **star:7308** Manage local application configuration files using templates and data from etcd or consul.   [![star > 2000][Awesome]](https://github.com/kelseyhightower/confd)   [![godoc][GoDoc]](https://godoc.org/github.com/kelseyhightower/confd)
* [Comcast](https://github.com/tylertreat/Comcast) **star:6823** Simulate bad network connections.   [![star > 2000][Awesome]](https://github.com/tylertreat/Comcast)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/Comcast)
* [LiteIDE](https://github.com/visualfc/liteide) **star:6301** LiteIDE is a simple, open source, cross-platform Go IDE.   [![star > 2000][Awesome]](https://github.com/visualfc/liteide)   [![Contains Chinese documents][CN]](https://github.com/visualfc/liteide)
* [drive](https://github.com/odeke-em/drive) **star:5821** Google Drive client for the commandline.   [![star > 2000][Awesome]](https://github.com/odeke-em/drive)   [![godoc][GoDoc]](https://godoc.org/github.com/odeke-em/drive)
* [toxiproxy](https://github.com/shopify/toxiproxy) **star:5117** Proxy to simulate network and system conditions for automated tests.   [![star > 2000][Awesome]](https://github.com/shopify/toxiproxy)   [![godoc][GoDoc]](https://godoc.org/github.com/shopify/toxiproxy)
* [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
* [nes](https://github.com/fogleman/nes) **star:4668** Nintendo Entertainment System (NES) emulator written in Go.   [![star > 2000][Awesome]](https://github.com/fogleman/nes)   [![godoc][GoDoc]](https://godoc.org/github.com/fogleman/nes)
* [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:3561** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   [![star > 2000][Awesome]](https://github.com/gilbertchen/duplicacy)   [![godoc][GoDoc]](https://godoc.org/github.com/gilbertchen/duplicacy)
* [myLG](https://github.com/mehrdadrad/mylg) **star:2410** Command Line Network Diagnostic tool written in Go.   [![star > 2000][Awesome]](https://github.com/mehrdadrad/mylg)   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/mylg)
* [GoBoy](https://github.com/Humpheh/goboy) **star:2282** Nintendo Game Boy Color emulator written in Go.   [![star > 2000][Awesome]](https://github.com/Humpheh/goboy)   [![godoc][GoDoc]](https://godoc.org/github.com/Humpheh/goboy)
* [Stack Up](https://github.com/pressly/sup) **star:2233** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   [![star > 2000][Awesome]](https://github.com/pressly/sup)   [![godoc][GoDoc]](https://godoc.org/github.com/pressly/sup)
* [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
* [lgo](https://github.com/yunabe/lgo) **star:2107** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   [![star > 2000][Awesome]](https://github.com/yunabe/lgo)   [![godoc][GoDoc]](https://godoc.org/github.com/yunabe/lgo)
* [limetext](https://limetext.github.io)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [scc](https://github.com/boyter/scc) **star:2025** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   [![star > 2000][Awesome]](https://github.com/boyter/scc)   [![There was an update last week][Green]](https://github.com/boyter/scc)   [![godoc][GoDoc]](https://godoc.org/github.com/boyter/scc)
* [Circuit](https://github.com/gocircuit/circuit) **star:1877** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   [![godoc][GoDoc]](https://godoc.org/github.com/gocircuit/circuit)
* [snap](https://github.com/intelsdi-x/snap) **star:1797** Powerful telemetry framework.   [![It hasn't been updated in the last year][Yellow]](https://github.com/intelsdi-x/snap)   [![godoc][GoDoc]](https://godoc.org/github.com/intelsdi-x/snap)
* [borg](https://github.com/crufter/borg) **star:1490** Terminal based search engine for bash snippets.   [![It hasn't been updated in the last year][Yellow]](https://github.com/crufter/borg)   [![godoc][GoDoc]](https://godoc.org/github.com/crufter/borg)
* [Documize](https://github.com/documize/community) **star:1263** Modern wiki software that integrates data from SaaS tools.
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:889** App that displays updates for the Go packages in your GOPATH.   [![godoc][GoDoc]](https://godoc.org/github.com/shurcooL/Go-Package-Store)
* [vFlow](https://github.com/VerizonDigital/vflow) **star:762** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   [![godoc][GoDoc]](https://godoc.org/github.com/VerizonDigital/vflow)
* [peg](https://github.com/pointlander/peg) **star:760** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   [![godoc][GoDoc]](https://godoc.org/github.com/pointlander/peg)
* [Leaps](https://github.com/jeffail/leaps) **star:696** Pair programming service using Operational Transforms.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jeffail/leaps)   [![godoc][GoDoc]](https://godoc.org/github.com/jeffail/leaps)
* [shell2http](https://github.com/msoap/shell2http) **star:691** Executing shell commands via http server (for prototyping or remote control).   [![godoc][GoDoc]](https://godoc.org/github.com/msoap/shell2http)
* [gfile](https://github.com/Antonito/gfile) **star:593** Securely transfer files between two computers, without any third party, over WebRTC.   [![It hasn't been updated in the last year][Yellow]](https://github.com/Antonito/gfile)   [![godoc][GoDoc]](https://godoc.org/github.com/Antonito/gfile)
* [Gebug](https://github.com/moshebe/gebug) **star:516** A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.   [![There was an update last week][Green]](https://github.com/moshebe/gebug)   [![godoc][GoDoc]](https://godoc.org/github.com/moshebe/gebug)
* [Guora](https://github.com/meloalright/guora) **star:509** A self-hosted Quora like web application written in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/meloalright/guora)   [![Contains Chinese documents][CN]](https://github.com/meloalright/guora)
* [hugo](http://gohugo.io/)  Fast and Modern Static Website Engine.
* [mockingjay](https://github.com/quii/mockingjay-server) **star:471** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   [![godoc][GoDoc]](https://godoc.org/github.com/quii/mockingjay-server)
* [gocc](https://github.com/goccmack/gocc) **star:446** Gocc is a compiler kit for Go written in Go.   [![There was an update last week][Green]](https://github.com/goccmack/gocc)   [![godoc][GoDoc]](https://godoc.org/github.com/goccmack/gocc)
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:423** Video streaming torrent client.   [![godoc][GoDoc]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
* [ipe](https://github.com/dimiro1/ipe) **star:312** Open source Pusher server implementation compatible with Pusher client libraries written in GO.   [![godoc][GoDoc]](https://godoc.org/github.com/dimiro1/ipe)
* [ide](https://github.com/thestrukture/ide) **star:307** Browser accessible IDE. Designed for Go with Go.   [![godoc][GoDoc]](https://godoc.org/github.com/thestrukture/ide)
* [wellington](https://github.com/wellington/wellington) **star:296** Sass project management tool, extends the language with sprite functions (like Compass).   [![godoc][GoDoc]](https://godoc.org/github.com/wellington/wellington)
* [Cherry](https://github.com/rafael-santiago/cherry) **star:237** Tiny webchat server in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/rafael-santiago/cherry)   [![godoc][GoDoc]](https://godoc.org/github.com/rafael-santiago/cherry)
* [tcpprobe](https://github.com/mehrdadrad/tcpprobe) **star:178** TCP tool for network performance and path monitoring, including socket statistics.   [![godoc][GoDoc]](https://godoc.org/github.com/mehrdadrad/tcpprobe)
* [woke](https://github.com/get-woke/woke) **star:175** Detect non-inclusive language in your source code.   [![godoc][GoDoc]](https://godoc.org/github.com/get-woke/woke)
* [joincap](https://github.com/assafmo/joincap) **star:157** Command-line utility for merging multiple pcap files together.   [![godoc][GoDoc]](https://godoc.org/github.com/assafmo/joincap)
* [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [Orbit](https://github.com/gulien/orbit) **star:146** A simple tool for running commands and generating files from templates.   [![godoc][GoDoc]](https://godoc.org/github.com/gulien/orbit)
* [Better Go Playground](https://goplay.tools)  Go playground with syntax highlight, code completion and other features.
* [vaku](https://github.com/lingrino/vaku) **star:101** CLI & API for folder-based functions in Vault like copy, move, and search.   [![There was an update last week][Green]](https://github.com/lingrino/vaku)   [![godoc][GoDoc]](https://godoc.org/github.com/lingrino/vaku)
* [boxed](https://github.com/tejo/boxed) **star:74** Dropbox based blog engine.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tejo/boxed)   [![godoc][GoDoc]](https://godoc.org/github.com/tejo/boxed)
* [dp](https://github.com/scryinfo/dp) **star:60** Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.   [![godoc][GoDoc]](https://godoc.org/github.com/scryinfo/dp)
* [naclpipe](https://github.com/unix4fun/naclpipe) **star:21** Simple NaCL EC25519 based crypto pipe tool written in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/unix4fun/naclpipe)   [![godoc][GoDoc]](https://godoc.org/github.com/unix4fun/naclpipe)
* [term-quiz](https://github.com/crazcalm/term-quiz) **star:17** Quizzes for your terminal.   [![It hasn't been updated in the last year][Yellow]](https://github.com/crazcalm/term-quiz)   [![godoc][GoDoc]](https://godoc.org/github.com/crazcalm/term-quiz)
* [Snitch](https://github.com/lucasgomide/snitch) **star:15** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   [![It hasn't been updated in the last year][Yellow]](https://github.com/lucasgomide/snitch)   [![godoc][GoDoc]](https://godoc.org/github.com/lucasgomide/snitch)
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:13** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.
* [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1466** Go HTTP request router benchmark and comparison.   [![godoc][GoDoc]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:1395** Go web framework benchmark.   [![godoc][GoDoc]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:1091** Benchmarks of Go serialization methods.   [![godoc][GoDoc]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
* [skynet](https://github.com/atemerev/skynet) **star:975** Skynet 1M threads microbenchmark.   [![There was an update last week][Green]](https://github.com/atemerev/skynet)
* [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:207** Compare various Image resize algorithms for the Go language.   [![godoc][GoDoc]](https://godoc.org/github.com/fawick/speedtest-resize)
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:136** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tylertreat/go-benchmarks)   [![godoc][GoDoc]](https://godoc.org/github.com/tylertreat/go-benchmarks)
* [gospeed](https://github.com/feyeleanor/GoSpeed) **star:106** Go micro-benchmarks for calculating the speed of language constructs.   [![It hasn't been updated in the last year][Yellow]](https://github.com/feyeleanor/GoSpeed)   [![godoc][GoDoc]](https://godoc.org/github.com/feyeleanor/GoSpeed)
* [autobench](https://github.com/davecheney/autobench) **star:89** Framework to compare the performance between different Go versions.   [![It hasn't been updated in the last year][Yellow]](https://github.com/davecheney/autobench)   [![godoc][GoDoc]](https://godoc.org/github.com/davecheney/autobench)
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) **star:55** Benchmarks of common basic operations for the Go language.   [![It hasn't been updated in the last year][Yellow]](https://github.com/PuerkitoBio/gocostmodel)   [![godoc][GoDoc]](https://godoc.org/github.com/PuerkitoBio/gocostmodel)
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) **star:54** Collection of benchmarks for popular Go database/SQL utilities.   [![It hasn't been updated in the last year][Yellow]](https://github.com/tyler-smith/golang-sql-benchmark)   [![godoc][GoDoc]](https://godoc.org/github.com/tyler-smith/golang-sql-benchmark)
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:21** Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mrLSD/go-benchmark-app)   [![godoc][GoDoc]](https://godoc.org/github.com/mrLSD/go-benchmark-app)
* [kvbench](https://github.com/jimrobinson/kvbench) **star:19** Key/Value database benchmark.   [![It hasn't been updated in the last year][Yellow]](https://github.com/jimrobinson/kvbench)   [![godoc][GoDoc]](https://godoc.org/github.com/jimrobinson/kvbench)
* [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) **star:4** Go JSON benchmark.   [![godoc][GoDoc]](https://godoc.org/github.com/zerosnake0/go-json-benchmark)

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
* [GopherCon Europe](https://gophercon.is/)  Berlin, Germany.
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
* [For the Love of Go](https://bitfieldconsulting.com/books)  A series of introductory books for Go beginners.
* [Go 101](https://go101.org)  A book focusing on Go syntax/semantics and all kinds of details.
* [Go Bootcamp](http://golangbootcamp.com)
* [GoBooks](https://github.com/dariubs/GoBooks) **star:8815** A curated list of Go books.   [![star > 2000][Awesome]](https://github.com/dariubs/GoBooks)   [![There was an update last week][Green]](https://github.com/dariubs/GoBooks)
* [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook)  A 600 page introduction to Go aimed at first time developers.
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:18** in Persian.   [![It hasn't been updated in the last year][Yellow]](https://github.com/thedevsir/gosuccinctly)   [![godoc][GoDoc]](https://godoc.org/github.com/thedevsir/gosuccinctly)
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
* [The Go Programming Language](http://www.gopl.io/)
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:2135** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   [![star > 2000][Awesome]](https://github.com/MariaLetta/free-gophers-pack)   [![godoc][GoDoc]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
* [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:83** adorable gopher logos.   [![It hasn't been updated in the last year][Yellow]](https://github.com/GolangUA/gopher-logos)
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:37** Go gopher Vector Data [.ai, .svg].   [![It hasn't been updated in the last year][Yellow]](https://github.com/keygx/Go-gopher-Vector)
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
* [gopher-vector](https://github.com/golang-samples/gopher-vector)
* [gophericons](https://github.com/shalakhin/gophericons)
* [gophers](https://github.com/ashleymcnamara/gophers) **star:2364** Gopher artworks by Ashley McNamara.   [![star > 2000][Awesome]](https://github.com/ashleymcnamara/gophers)   [![It hasn't been updated in the last year][Yellow]](https://github.com/ashleymcnamara/gophers)   [![godoc][GoDoc]](https://godoc.org/github.com/ashleymcnamara/gophers)
* [gophers](https://github.com/egonelbre/gophers) **star:2258** Free gophers.   [![star > 2000][Awesome]](https://github.com/egonelbre/gophers)   [![godoc][GoDoc]](https://godoc.org/github.com/egonelbre/gophers)
* [gopherize.me](https://github.com/matryer/gopherize.me) **star:473** Gopherize yourself.
* [gophers](https://github.com/sillecelik/go-gopher) **star:79** Gopher amigurumi toy pattern.
* [gophers](https://github.com/rogeralsing/gophers) **star:54** random gopher graphics.

## Meetups

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
* [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
* [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
* [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
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
* [Golang Turkey](https://kommunity.com/goturkiye)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Vienna, Austria](https://www.meetup.com/viennago/)
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

## Style Guides

* [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide)
* [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md)
* [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
* [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst)
* [Magnetico](https://github.com/boramalper/magnetico/wiki/magnetico-Design-Specification)
* [Sourcegraph](https://about.sourcegraph.com/handbook/engineering/go_style_guide)
* [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)
* [Uber](https://github.com/uber-go/guide/blob/master/style.md)

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

* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:27140** List of other amazingly awesome lists.   [![star > 2000][Awesome]](https://github.com/bayandin/awesome-awesomeness)   [![There was an update last week][Green]](https://github.com/bayandin/awesome-awesomeness)
* [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
* [Go Blog](http://blog.golang.org)  The official Go blog.
* [Go Challenge](http://golang-challenge.org/)  Learn Go by solving problems and getting feedback from Go experts.
* [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3)  A group of Gophers read and discuss a different Go project every week.
* [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/)  5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com)  A report card for your Go package.
* [go.dev](https://go.dev/)  A hub for Go developers.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:19860** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.   [![star > 2000][Awesome]](https://github.com/lukasz-madon/awesome-remote-job)   [![There was an update last week][Green]](https://github.com/lukasz-madon/awesome-remote-job)
* [golang-graphics](https://github.com/mholt/golang-graphics) **star:142** Collection of Go images, graphics, and art.   [![It hasn't been updated in the last year][Yellow]](https://github.com/mholt/golang-graphics)   [![Archived][Archived]](https://github.com/mholt/golang-graphics)
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571)  The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
* [gowalker.org](https://gowalker.org)  Go Project API documentation.
* [json2go](https://m-zajac.github.io/json2go)  Advanced JSON to Go struct conversion - online tool.
* [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [Lille Gophers](https://lille-gophers.loscrackitos.codes/)  Golang talks community in Lille, France ([@LilleGophers](https://twitter.com/LilleGophers)).
* [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:38** Collection of Go projects that needs help. Good place to start your open-source way in Go.   [![It hasn't been updated in the last year][Yellow]](https://github.com/ninedraft/gocryforhelp)
* [godoc.org](https://godoc.org/)  Documentation for open source Go packages.
* [Golang Developer Jobs](https://golangjob.xyz)  Developer Jobs exclusivly for Golang related Roles.
* [Golang Flow](https://golangflow.io)  Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com)  Links and news about Go programming.
* [Golang Resources](https://golangresources.com)  A curation of the best articles, exercises, talks and videos to learn Go.
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang)  News about Go.
* [studygolang](https://studygolang.com)  The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:36785** Golang ebook intro how to build a web app with golang.   [![star > 2000][Awesome]](https://github.com/astaxie/build-web-application-with-golang)   [![godoc][GoDoc]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   [![Contains Chinese documents][CN]](https://github.com/astaxie/build-web-application-with-golang)
* [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)  We’ll write an API with the help of the powerful Gorilla Mux.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:5090** Go's reference card.   [![star > 2000][Awesome]](https://github.com/a8m/go-lang-cheat-sheet)
* [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
* [Go Playground for iOS](https://codeplayground.app)  Interactively edit & play Go snippets on your mobile device.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) **star:723** A little e-book on Ethereum Development with Go.   [![There was an update last week][Green]](https://github.com/miguelmota/ethereum-development-with-go-book)   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/ethereum-development-with-go-book)   [![Contains Chinese documents][CN]](https://github.com/miguelmota/ethereum-development-with-go-book)
* [Games With Go](http://gameswithgo.org/)  A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
* [A Tour of Go](http://tour.golang.org/)  Interactive tour of Go.
* [Design Patterns in Go](https://github.com/shubhamzanwar/design-patterns) **star:42** Collection of programming design patterns implemented in Go.   [![godoc][GoDoc]](https://godoc.org/github.com/shubhamzanwar/design-patterns)
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [go-patterns](https://github.com/tmrts/go-patterns) **star:14465** Curated list of Go design patterns, recipes and idioms.   [![star > 2000][Awesome]](https://github.com/tmrts/go-patterns)   [![There was an update last week][Green]](https://github.com/tmrts/go-patterns)   [![godoc][GoDoc]](https://godoc.org/github.com/tmrts/go-patterns)
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:12690** Learn Go with test-driven development.   [![star > 2000][Awesome]](https://github.com/quii/learn-go-with-tests)   [![There was an update last week][Green]](https://github.com/quii/learn-go-with-tests)   [![godoc][GoDoc]](https://godoc.org/github.com/quii/learn-go-with-tests)   [![Contains Chinese documents][CN]](https://github.com/quii/learn-go-with-tests)
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)  Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera Specialization to learn about Go from scratch.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:1566** Examples of Golang compared to Node.js for learning.   [![godoc][GoDoc]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)
* [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/)  Collection of code snippets and tutorials to help tackle every day issues.
* [GopherSnippets](https://gophersnippets.com/)  Code snippets with tests and testable examples for the Go programming language.
* [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5)  Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
* [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker)  Learn how to use Docker for Go development and how to build production Docker images.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [goapp](https://github.com/bnkamalesh/goapp) **star:219** An opinionated guideline to structure & develop a Go web application/service.   [![godoc][GoDoc]](https://godoc.org/github.com/bnkamalesh/goapp)
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [Working with Go](https://github.com/mkaz/working-with-go) **star:1162** Intro to go for experienced programmers.   [![Archived][Archived]](https://github.com/mkaz/working-with-go)
* [Your basic Go](http://yourbasic.org/golang)  Huge collection of tutorials and how to's.

