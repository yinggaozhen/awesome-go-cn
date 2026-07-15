# Awesome Go

[Awesome]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.4.1/docs/awesome.svg "star > 2000"
[G]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Green.svg "There was an update last week"
[Y]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Yellow.svg "It hasn't been updated in recent three years"
[CN]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.1/docs/Cn.svg "Contains Chinese documents"
[Archived]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.2.1/docs/archived.svg "The project has been archived"
[D]: https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/DOC.svg "godoc document links"

**This project is [awesome-go](https://awesome-go.com/) Chinese version, last sync time : 2026-07-15 08:10:19(Synchronize every day)**

[![english](https://cdn.jsdelivr.net/gh/yinggaozhen/awesome-go-cn@1.3.0/docs/chinese.svg)](README.md) [![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) financial support to Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python)。

**Icon** :


Icon | State  
:-:|-
![There was an update last week][G] | There was an update last week。You can basically determine that the current library is in an actively maintained state。
![Not updated in recent three year][Y] | Not updated in recent three year。Reflects that the maintenance of the library is not high enthusiasm, use should be careful。
![Archived][Archived] | The project has been archived。
![Contains Chinese documents][CN] | This project contains Chinese documents。
![godoc document][D] | godoc document links。

### Explain

[中文](README.md)  | [English](README_EN.md)
 
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)
[![Last Commit](https://img.shields.io/github/last-commit/avelino/awesome-go)](https://github.com/avelino/awesome-go/commits/main)

We use the _[Golang Bridge](https://github.com/gobridge/about-us/blob/master/README.md)_ community Slack for instant communication, follow the [form here to join](https://invite.slack.golangbridge.org/).

<a href="https://www.producthunt.com/posts/awesome-go?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-go" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=291535&theme=light" alt="awesome-go - Curated list awesome Go frameworks, libraries and software | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

**Sponsorships:**

_Special thanks to_

<div align="center">
<table cellpadding="5">
<tbody align="center">
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-digitalocean">
<img src="https://avelino.run/sponsors/do_logo_horizontal_blue-210.png" width="200" alt="Digital Ocean">
</a>
</td>
</tr>
</tbody>
</table>
</div>

**Awesome Go has no monthly fee**_, but we have employees who **work hard** to keep it running. With money raised, we can repay the effort of each person involved! You can see how we calculate our billing and distribution as it is open to the entire community. Want to be a supporter of the project click [here](mailto:avelinorun+oss@gmail.com?subject=awesome-go%3A%20project%20support)._

> A curated list of awesome Go frameworks, libraries, and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

**Contributing:**

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

> _If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!_

## Contents

<details>
<summary>Expand contents</summary>

- [Awesome Go](#awesome-go)
  - [Contents](#contents)
  - [Actor Model](#actor-model)
  - [Artificial Intelligence](#artificial-intelligence)
  - [Audio and Music](#audio-and-music)
  - [Authentication and Authorization](#authentication-and-authorization)
  - [Blockchain](#blockchain)
  - [Bot Building](#bot-building)
  - [Build Automation](#build-automation)
  - [Command Line](#command-line)
    - [Advanced Console UIs](#advanced-console-uis)
    - [Standard CLI](#standard-cli)
  - [Configuration](#configuration)
  - [Continuous Integration](#continuous-integration)
  - [CSS Preprocessors](#css-preprocessors)
  - [Data Integration Frameworks](#data-integration-frameworks)
  - [Data Structures and Algorithms](#data-structures-and-algorithms)
    - [Bit-packing and Compression](#bit-packing-and-compression)
    - [Bit Sets](#bit-sets)
    - [Bloom and Cuckoo Filters](#bloom-and-cuckoo-filters)
    - [Data Structure and Algorithm Collections](#data-structure-and-algorithm-collections)
    - [Iterators](#iterators)
    - [Maps](#maps)
    - [Miscellaneous Data Structures and Algorithms](#miscellaneous-data-structures-and-algorithms)
    - [Nullable Types](#nullable-types)
    - [Queues](#queues)
    - [Sets](#sets)
    - [Text Analysis](#text-analysis)
    - [Trees](#trees)
    - [Pipes](#pipes)
  - [Database](#database)
    - [Caches](#caches)
    - [Databases Implemented in Go](#databases-implemented-in-go)
    - [Database Schema Migration](#database-schema-migration)
    - [Database Tools](#database-tools)
    - [SQL Query Builders](#sql-query-builders)
  - [Database Drivers](#database-drivers)
    - [Interfaces to Multiple Backends](#interfaces-to-multiple-backends)
    - [Relational Database Drivers](#relational-database-drivers)
    - [NoSQL Database Drivers](#nosql-database-drivers)
    - [Search and Analytic Databases](#search-and-analytic-databases)
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
  - [Generators](#generators)
  - [Geographic](#geographic)
  - [Go Compilers](#go-compilers)
  - [Goroutines](#goroutines)
  - [GUI](#gui)
  - [Hardware](#hardware)
  - [Images](#images)
  - [IoT (Internet of Things)](#iot-internet-of-things)
  - [Job Scheduler](#job-scheduler)
  - [JSON](#json)
  - [Logging](#logging)
  - [Machine Learning](#machine-learning)
  - [Messaging](#messaging)
  - [Microsoft Office](#microsoft-office)
    - [Microsoft Excel](#microsoft-excel)
    - [Microsoft Word](#microsoft-word)
  - [Miscellaneous](#miscellaneous)
    - [Dependency Injection](#dependency-injection)
    - [Project Layout](#project-layout)
    - [Strings](#strings)
    - [Uncategorized](#uncategorized)
  - [Natural Language Processing](#natural-language-processing)
    - [Language Detection](#language-detection)
    - [Morphological Analyzers](#morphological-analyzers)
    - [Slugifiers](#slugifiers)
    - [Tokenizers](#tokenizers)
    - [Translation](#translation)
    - [Transliteration](#transliteration)
  - [Networking](#networking)
    - [HTTP Clients](#http-clients)
  - [OpenGL](#opengl)
  - [ORM](#orm)
  - [Package Management](#package-management)
  - [Performance](#performance)
  - [Query Language](#query-language)
  - [Reflection](#reflection)
  - [Resource Embedding](#resource-embedding)
  - [Science and Data Analysis](#science-and-data-analysis)
  - [Security](#security)
  - [Serialization](#serialization)
  - [Server Applications](#server-applications)
  - [Stream Processing](#stream-processing)
  - [Template Engines](#template-engines)
  - [Testing](#testing)
    - [Testing Frameworks](#testing-frameworks)
    - [Mock](#mock)
    - [Fuzzing and delta-debugging/reducing/shrinking](#fuzzing-and-delta-debuggingreducingshrinking)
    - [Selenium and browser control tools](#selenium-and-browser-control-tools)
    - [Fail injection](#fail-injection)
  - [Text Processing](#text-processing)
    - [Formatters](#formatters)
    - [Markup Languages](#markup-languages)
    - [Parsers/Encoders/Decoders](#parsersencodersdecoders)
    - [Regular Expressions](#regular-expressions)
    - [Sanitation](#sanitation)
    - [Scrapers](#scrapers)
    - [RSS](#rss)
    - [Utility/Miscellaneous](#utilitymiscellaneous)
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
  - [Webhooks Server](#webhooks-server)
  - [Windows](#windows)
  - [Workflow Frameworks](#workflow-frameworks)
  - [XML](#xml)
  - [Zero Trust](#zero-trust)
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
    - [E-books for purchase](#e-books-for-purchase)
    - [Free e-books](#free-e-books)
  - [Gophers](#gophers)
  - [Meetups](#meetups)
  - [Style Guides](#style-guides)
  - [Social Media](#social-media)
    - [Twitter](#twitter)
    - [Reddit](#reddit)
  - [Websites](#websites)
    - [Tutorials](#tutorials)
    - [Guided Learning](#guided-learning)
  - [Contribution](#contribution)
  - [License](#license)

**[⬆ back to top](#contents)**



</details>

## Actor Model

_Libraries for building actor-based programs._

- [ProtoActor](https://github.com/asynkron/protoactor-go) **star:5481** Distributed actors for Go, C#, and Java/Kotlin.   [![godoc][D]](https://godoc.org/github.com/asynkron/protoactor-go)
- [Ergo](https://github.com/ergo-services/ergo) **star:4614** An actor-based Framework with network transparency for creating event-driven architecture in Golang. Inspired by Erlang.   [![godoc][D]](https://godoc.org/github.com/ergo-services/ergo)
- [Hollywood](https://github.com/anthdm/hollywood) **star:2315** Blazingly fast and light-weight Actor engine written in Golang.   [![godoc][D]](https://godoc.org/github.com/anthdm/hollywood)
- [Goakt](https://github.com/Tochemey/goakt) **star:361** Fast and Distributed Actor framework using protocol buffers as message for Golang.   [![There was an update last month][G]](https://github.com/Tochemey/goakt)   [![godoc][D]](https://godoc.org/github.com/Tochemey/goakt)
- [asyncmachine-go/pkg/machine](https://github.com/pancsta/asyncmachine-go/tree/main/pkg/machine)  Graph control flow library (AOP, actor, state-machine).

**[⬆ back to top](#contents)**

## Artificial Intelligence

_Libraries for building programs that leverage AI._

- [Ollama](https://github.com/jmorganca/ollama) **star:176115** Run large language models locally.   [![There was an update last month][G]](https://github.com/jmorganca/ollama)   [![godoc][D]](https://godoc.org/github.com/jmorganca/ollama)
- [LocalAI](https://github.com/mudler/LocalAI) **star:47542** Open Source OpenAI alternative, self-host AI models.   [![There was an update last month][G]](https://github.com/mudler/LocalAI)   [![godoc][D]](https://godoc.org/github.com/mudler/LocalAI)
- [langchaingo](https://github.com/tmc/langchaingo) **star:9538** LangChainGo is a framework for developing applications powered by language models.   [![godoc][D]](https://godoc.org/github.com/tmc/langchaingo)
- [mcp-go](https://github.com/mark3labs/mcp-go) **star:8892** Go implementation of the Model Context Protocol for building MCP servers and clients in Go.   [![There was an update last month][G]](https://github.com/mark3labs/mcp-go)   [![godoc][D]](https://godoc.org/github.com/mark3labs/mcp-go)
- [trpc-agent-go](https://github.com/trpc-group/trpc-agent-go) **star:1542** Framework for building LLM-based multi-agent systems.   [![There was an update last month][G]](https://github.com/trpc-group/trpc-agent-go)   [![godoc][D]](https://godoc.org/github.com/trpc-group/trpc-agent-go)   [![Contains Chinese documents][CN]](https://github.com/trpc-group/trpc-agent-go)
- [chromem-go](https://github.com/philippgille/chromem-go) **star:1029** Embeddable vector database for Go with Chroma-like interface and zero third-party dependencies. In-memory with optional persistence.   [![godoc][D]](https://godoc.org/github.com/philippgille/chromem-go)
- [GoModel](https://github.com/ENTERPILOT/GoModel) **star:998** AI gateway exposing a unified OpenAI-compatible API across OpenAI, Anthropic, Gemini, Groq, xAI, Ollama and other providers, with routing, usage tracking, rate limits, and guardrails.   [![There was an update last month][G]](https://github.com/ENTERPILOT/GoModel)   [![godoc][D]](https://godoc.org/github.com/ENTERPILOT/GoModel)
- [langgraphgo](https://github.com/smallnest/langgraphgo) **star:275** A Go library for building stateful, multi-actor applications with LLMs, built on the concept of LangGraph，with a lot of builtin Agent architectures.   [![There was an update last month][G]](https://github.com/smallnest/langgraphgo)   [![godoc][D]](https://godoc.org/github.com/smallnest/langgraphgo)   [![Contains Chinese documents][CN]](https://github.com/smallnest/langgraphgo)
- [goai](https://github.com/zendev-sh/goai) **star:165** Go SDK for building AI applications. One SDK, 20+ providers. Inspired by Vercel AI SDK.   [![There was an update last month][G]](https://github.com/zendev-sh/goai)   [![godoc][D]](https://godoc.org/github.com/zendev-sh/goai)
- [OllamaFarm](https://github.com/presbrey/ollamafarm) **star:100** Manage, load-balance, and failover packs of Ollamas.   [![There was an update last month][G]](https://github.com/presbrey/ollamafarm)   [![godoc][D]](https://godoc.org/github.com/presbrey/ollamafarm)
- [hotplex](https://github.com/hrygo/hotplex) **star:45** AI Agent runtime engine with long-lived sessions for Claude Code, OpenCode, pi-mono and other CLI AI tools. Provides full-duplex streaming, multi-platform integrations, and secure sandbox.   [![There was an update last month][G]](https://github.com/hrygo/hotplex)   [![godoc][D]](https://godoc.org/github.com/hrygo/hotplex)   [![Contains Chinese documents][CN]](https://github.com/hrygo/hotplex)
- [skillreaper](https://github.com/thousandflowers/skillreaper) **star:43** CLI that scans AI agent session transcripts to identify and safely quarantine unused skills, MCP servers, and agents across Claude Code, Codex CLI, Hermes, OpenCode, Cursor, and OpenClaw.   [![godoc][D]](https://godoc.org/github.com/thousandflowers/skillreaper)
- [web-researcher-mcp](https://github.com/zoharbabin/web-researcher-mcp) **star:41** MCP server providing AI assistants with web search, content extraction, and multi-source research capabilities. Single binary, 5 search providers with circuit-breaker failover, 4-tier scraping pipeline.   [![There was an update last month][G]](https://github.com/zoharbabin/web-researcher-mcp)   [![godoc][D]](https://godoc.org/github.com/zoharbabin/web-researcher-mcp)
- [agent-sdk-go](https://github.com/agenticenv/agent-sdk-go) **star:38** Go SDK for building durable AI agents on Temporal with support for tools, MCP, human approvals, and sub-agent delegation.   [![There was an update last month][G]](https://github.com/agenticenv/agent-sdk-go)   [![godoc][D]](https://godoc.org/github.com/agenticenv/agent-sdk-go)
- [ai](https://github.com/joakimcarlsson/ai) **star:33** A Go toolkit for building AI agents and applications across multiple providers with unified LLM, embeddings, tool calling, and MCP integration.   [![There was an update last month][G]](https://github.com/joakimcarlsson/ai)   [![godoc][D]](https://godoc.org/github.com/joakimcarlsson/ai)
- [otellix](https://github.com/oluwajubelo1/otellix) **star:33** OpenTelemetry-native LLM observability and budget guardrails for cost-constrained production environments.   [![godoc][D]](https://godoc.org/github.com/oluwajubelo1/otellix)
- [zenflow](https://github.com/zendev-sh/zenflow) **star:32** Multi-agent orchestration & workflow engine. Declarative YAML workflows, LLM coordinator with hub-and-spoke mailboxes, race-safe delivery. One YAML file, one Go binary. Runs on any goai-supported provider.   [![godoc][D]](https://godoc.org/github.com/zendev-sh/zenflow)
- [AegisFlow](https://github.com/saivedant169/AegisFlow) **star:25** AI gateway for routing, securing, and monitoring LLM traffic across 10+ providers. OpenAI-compatible API, WASM policy plugins, canary rollouts, real-time dashboard.   [![There was an update last month][G]](https://github.com/saivedant169/AegisFlow)   [![godoc][D]](https://godoc.org/github.com/saivedant169/AegisFlow)
- [jargo](https://github.com/gojargo/jargo) **star:25** Framework for building real-time voice AI agents over WebRTC, wiring speech-to-text, LLMs, and text-to-speech into a streaming pipeline.   [![There was an update last month][G]](https://github.com/gojargo/jargo)   [![godoc][D]](https://godoc.org/github.com/gojargo/jargo)
- [routex](https://github.com/Ad3bay0c/routex) **star:22** YAML-driven multi-agent AI runtime for Go with Erlang-style supervision, MCP tool server support, and a CLI.   [![godoc][D]](https://godoc.org/github.com/Ad3bay0c/routex)
- [Aetheris](https://github.com/Colin4k1024/Aetheris) **star:16** AI Agent execution runtime with event sourcing, checkpoint recovery, and At-Most-Once execution guarantee. Written in Go.   [![godoc][D]](https://godoc.org/github.com/Colin4k1024/Aetheris)
- [localaik](https://github.com/harshaneel/localaik) **star:15** LocalStack-style local emulation of OpenAI and Gemini APIs; single Docker container, llama.cpp + Gemma 3 backend.   [![godoc][D]](https://godoc.org/github.com/harshaneel/localaik)
- [llm-box](https://github.com/alib8b8/llm-box) **star:7** Terminal-based AI workflow engine with YAML-driven pipelines, 20+ LLM providers (DeepSeek, Qwen, GLM, Mistral, etc.), and a TUI for workflow management.   [![There was an update last month][G]](https://github.com/alib8b8/llm-box)   [![godoc][D]](https://godoc.org/github.com/alib8b8/llm-box)
- [dakera-go](https://github.com/dakera-ai/dakera-go) **star:3** Official Go client SDK for the Dakera self-hosted agent memory server, providing typed interfaces for memory store/recall, session management, namespace operations, and decay configuration.   [![godoc][D]](https://godoc.org/github.com/dakera-ai/dakera-go)
- [fun](https://gitlab.com/tozd/go/fun)  The simplest but powerful way to use large language models (LLMs) in Go.

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio and music._

- [Oto](https://github.com/hajimehoshi/oto) **star:1938** A low-level library to play sound on multiple platforms.   [![godoc][D]](https://godoc.org/github.com/hajimehoshi/oto)
- [PortAudio](https://github.com/gordonklaus/portaudio) **star:840** Go bindings for the PortAudio audio I/O library.   [![godoc][D]](https://godoc.org/github.com/gordonklaus/portaudio)
- [beep](https://github.com/gopxl/beep) **star:581** A simple library for playback and audio manipulation.   [![godoc][D]](https://godoc.org/github.com/gopxl/beep)
- [music-theory](https://github.com/go-music-theory/music-theory) **star:461** Music theory models in Go.   [![godoc][D]](https://godoc.org/github.com/go-music-theory/music-theory)
- [GoAudio](https://github.com/DylanMeeus/GoAudio) **star:422** Native Go Audio Processing Library.   [![godoc][D]](https://godoc.org/github.com/DylanMeeus/GoAudio)
- [malgo](https://github.com/gen2brain/malgo) **star:417** Mini audio library.
- [id3v2](https://github.com/bogem/id3v2) **star:370** ID3 decoding and encoding library for Go.   [![godoc][D]](https://godoc.org/github.com/bogem/id3v2)
- [flac](https://github.com/mewkiz/flac) **star:357** Native Go FLAC encoder/decoder with support for FLAC streams.   [![godoc][D]](https://godoc.org/github.com/mewkiz/flac)
- [gaad](https://github.com/Comcast/gaad) **star:137** Native Go AAC bitstream parser.   [![It hasn't been updated in recent three years][Y]](https://github.com/Comcast/gaad)   [![godoc][D]](https://godoc.org/github.com/Comcast/gaad)
- [minimp3](https://github.com/tosone/minimp3) **star:133** Lightweight MP3 decoder library.
- [gosamplerate](https://github.com/dh1tw/gosamplerate) **star:38** libsamplerate bindings for go.   [![godoc][D]](https://godoc.org/github.com/dh1tw/gosamplerate)
- [go-mpris](https://github.com/leberKleber/go-mpris) **star:22** Client for mpris dbus interfaces.   [![godoc][D]](https://godoc.org/github.com/leberKleber/go-mpris)
- [go-resample](https://github.com/gojargo/go-resample)  Pure-Go (no cgo) audio sample-rate converter with sinc, linear, and zero-order-hold converters.
-[voxrai-ai](https://github.com/Voxray-AI/Voxray) - AI voice agents with a JSON configuration,  STT → LLM → TTS pipelines over WebSocket and WebRTC 

**[⬆ back to top](#contents)**

## Authentication and Authorization

_Libraries for implementing authentication and authorization._

- [casbin](https://github.com/casbin/casbin) **star:20241** Authorization library that supports access control models like ACL, RBAC, and ABAC.   [![godoc][D]](https://godoc.org/github.com/casbin/casbin)
- [jwt-go](https://github.com/golang-jwt/jwt) **star:9171** A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs.   [![godoc][D]](https://godoc.org/github.com/golang-jwt/jwt)
- [spicedb](https://github.com/authzed/spicedb) **star:6859** A Zanzibar-inspired database that enables fine-grained authorization.   [![There was an update last month][G]](https://github.com/authzed/spicedb)   [![godoc][D]](https://godoc.org/github.com/authzed/spicedb)   [![Contains Chinese documents][CN]](https://github.com/authzed/spicedb)
- [goth](https://github.com/markbates/goth) **star:6574** provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.   [![godoc][D]](https://godoc.org/github.com/markbates/goth)
- [oauth2](https://github.com/golang/oauth2) **star:5877** Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support.   [![godoc][D]](https://godoc.org/github.com/golang/oauth2)
- [openfga](https://github.com/openfga/openfga) **star:5437** Implementation of fine-grained authorization based on the "Zanzibar: Google's Consistent, Global Authorization System" paper. Backed by [CNCF](https://www.cncf.io/).   [![There was an update last month][G]](https://github.com/openfga/openfga)   [![godoc][D]](https://godoc.org/github.com/openfga/openfga)
- [keto](https://github.com/ory/keto) **star:5370** Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models.   [![There was an update last month][G]](https://github.com/ory/keto)   [![godoc][D]](https://godoc.org/github.com/ory/keto)
- [authboss](https://github.com/volatiletech/authboss) **star:4196** Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time.   [![There was an update last month][G]](https://github.com/volatiletech/authboss)   [![godoc][D]](https://godoc.org/github.com/volatiletech/authboss)
- [authgate](https://github.com/go-authgate/authgate)  A lightweight OAuth 2.0 Authorization Server supporting Device Authorization Grant ([RFC 8628](https://datatracker.ietf.org/doc/html/rfc8628)), Authorization Code Flow with PKCE ([RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749) + [RFC 7636](https://datatracker.ietf.org/doc/html/rfc7636)), and Client Credentials Grant for machine-to-machine authentication.
- [scs](https://github.com/alexedwards/scs) **star:2599** Session Manager for HTTP servers.   [![godoc][D]](https://godoc.org/github.com/alexedwards/scs)
- [jwx](https://github.com/lestrrat-go/jwx) **star:2403** Go module implementing various JWx (JWA/JWE/JWK/JWS/JWT, otherwise known as JOSE) technologies.   [![There was an update last month][G]](https://github.com/lestrrat-go/jwx)   [![godoc][D]](https://godoc.org/github.com/lestrrat-go/jwx)
- [gologin](https://github.com/dghubble/gologin) **star:1958** chainable handlers for login with OAuth1 and OAuth2 authentication providers.   [![There was an update last month][G]](https://github.com/dghubble/gologin)   [![godoc][D]](https://godoc.org/github.com/dghubble/gologin)
- [osin](https://github.com/openshift/osin) **star:1934** Golang OAuth2 server library.   [![godoc][D]](https://godoc.org/github.com/openshift/osin)
- [loginsrv](https://github.com/tarent/loginsrv) **star:1931** JWT login microservice with pluggable backends such as OAuth2 (Github), htpasswd, osiam.   [![It hasn't been updated in recent three years][Y]](https://github.com/tarent/loginsrv)   [![godoc][D]](https://godoc.org/github.com/tarent/loginsrv)
- [oidc](https://github.com/zitadel/oidc) **star:1850** Easy to use OpenID Connect client and server library written for Go and certified by the OpenID Foundation.   [![There was an update last month][G]](https://github.com/zitadel/oidc)   [![godoc][D]](https://godoc.org/github.com/zitadel/oidc)
- [gorbac](https://github.com/mikespook/gorbac) **star:1674** provides a lightweight role-based access control (RBAC) implementation in Golang.   [![godoc][D]](https://godoc.org/github.com/mikespook/gorbac)
- [paseto](https://github.com/o1egl/paseto) **star:939** Golang implementation of Platform-Agnostic Security Tokens (PASETO).   [![It hasn't been updated in recent three years][Y]](https://github.com/o1egl/paseto)   [![godoc][D]](https://godoc.org/github.com/o1egl/paseto)
- [jwt](https://github.com/cristalhq/jwt) **star:690** Safe, simple, and fast JSON Web Tokens for Go.   [![godoc][D]](https://godoc.org/github.com/cristalhq/jwt)
- [go-guardian](https://github.com/shaj13/go-guardian) **star:613** Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token, and Certificate based authentication.   [![godoc][D]](https://godoc.org/github.com/shaj13/go-guardian)
- [go-jose](https://github.com/go-jose/go-jose) **star:520** Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.   [![godoc][D]](https://godoc.org/github.com/go-jose/go-jose)
- [jwt](https://github.com/pascaldekloe/jwt) **star:361** Lightweight JSON Web Token (JWT) library.   [![It hasn't been updated in recent three years][Y]](https://github.com/pascaldekloe/jwt)   [![godoc][D]](https://godoc.org/github.com/pascaldekloe/jwt)
- [jeff](https://github.com/abraithwaite/jeff) **star:272** Simple, flexible, secure, and idiomatic web session management with pluggable backends.   [![godoc][D]](https://godoc.org/github.com/abraithwaite/jeff)
- [gosession](https://github.com/Kwynto/gosession) **star:258** This is quick session for net/http in GoLang. This package is perhaps the best implementation of the session mechanism, or at least it tries to become one.   [![godoc][D]](https://godoc.org/github.com/Kwynto/gosession)
- [jwt-auth](https://github.com/adam-hanna/jwt-auth) **star:238** JWT middleware for Golang http servers with many configuration options.   [![It hasn't been updated in recent three years][Y]](https://github.com/adam-hanna/jwt-auth)   [![godoc][D]](https://godoc.org/github.com/adam-hanna/jwt-auth)
- [goiabada](https://github.com/leodip/goiabada) **star:192** An open-source authentication and authorization server supporting OAuth2 and OpenID Connect.   [![There was an update last month][G]](https://github.com/leodip/goiabada)   [![godoc][D]](https://godoc.org/github.com/leodip/goiabada)
- [otpgen](https://github.com/grijul/otpgen) **star:144** Library to generate TOTP/HOTP codes.   [![godoc][D]](https://godoc.org/github.com/grijul/otpgen)
- [sessionup](https://github.com/swithek/sessionup) **star:131** Simple, yet effective HTTP session management and identification package.   [![godoc][D]](https://godoc.org/github.com/swithek/sessionup)
- [sjwt](https://github.com/brianvoe/sjwt) **star:123** Simple jwt generator and parser.   [![godoc][D]](https://godoc.org/github.com/brianvoe/sjwt)
- [session](https://github.com/icza/session) **star:119** Go session management for web servers (including support for Google App Engine - GAE).   [![godoc][D]](https://godoc.org/github.com/icza/session)
- [branca](https://github.com/essentialkaos/branca) **star:100** branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+.   [![godoc][D]](https://godoc.org/github.com/essentialkaos/branca)
- [securecookie](https://github.com/chmike/securecookie) **star:86** Efficient secure cookie encoding/decoding.   [![It hasn't been updated in recent three years][Y]](https://github.com/chmike/securecookie)   [![godoc][D]](https://godoc.org/github.com/chmike/securecookie)
- [sessions](https://github.com/adam-hanna/sessions) **star:79** Dead simple, highly performant, highly customizable sessions service for go http servers.   [![It hasn't been updated in recent three years][Y]](https://github.com/adam-hanna/sessions)   [![godoc][D]](https://godoc.org/github.com/adam-hanna/sessions)
- [otpgo](https://github.com/jltorresm/otpgo) **star:77** Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/jltorresm/otpgo)   [![godoc][D]](https://godoc.org/github.com/jltorresm/otpgo)
- [go-iam](https://github.com/melvinodsa/go-iam) **star:46** Developer-first Identity and Access Management system with a simple UI.   [![godoc][D]](https://godoc.org/github.com/melvinodsa/go-iam)
- [scope](https://github.com/SonicRoshan/scope) **star:42** Easily Manage OAuth2 Scopes In Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/SonicRoshan/scope)   [![godoc][D]](https://godoc.org/github.com/SonicRoshan/scope)
- [go-githubauth](https://github.com/jferrl/go-githubauth) **star:33** Utilities for GitHub authentication: generate and use GitHub application and installation tokens.   [![godoc][D]](https://godoc.org/github.com/jferrl/go-githubauth)
- [cookiestxt](https://github.com/mengzhuo/cookiestxt) **star:24** provides a parser of cookies.txt file format.   [![godoc][D]](https://godoc.org/github.com/mengzhuo/cookiestxt)
- [go-jwt](https://github.com/pardnchiu/go-jwt) **star:19** JWT authentication package providing access tokens and refresh tokens with fingerprinting, Redis storage, and automatic refresh capabilities.   [![There was an update last month][G]](https://github.com/pardnchiu/go-jwt)   [![godoc][D]](https://godoc.org/github.com/pardnchiu/go-jwt)
- [permissions](https://github.com/xyproto/permissions) **star:14** Library for keeping track of users, login states, and permissions. Uses secure cookies and bcrypt.   [![godoc][D]](https://godoc.org/github.com/xyproto/permissions)
- [x509proxy](https://github.com/vkuznet/x509proxy) **star:9** Library to handle X509 proxy certificates.   [![godoc][D]](https://godoc.org/github.com/vkuznet/x509proxy)
- [go-jwt](https://github.com/deatil/go-jwt) **star:3** A JWT (JSON Web Token) library for Go.   [![There was an update last month][G]](https://github.com/deatil/go-jwt)   [![godoc][D]](https://godoc.org/github.com/deatil/go-jwt)

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [go-ethereum](https://github.com/ethereum/go-ethereum) **star:51351** Official Go implementation of the Ethereum protocol.   [![There was an update last month][G]](https://github.com/ethereum/go-ethereum)   [![godoc][D]](https://godoc.org/github.com/ethereum/go-ethereum)
- [kubo](https://github.com/ipfs/kubo) **star:17083** An IPFS implementation in Go. It provides content-addressable storage which can be used for decentralized storage in DApps. It is based on the IPFS protocol.   [![There was an update last month][G]](https://github.com/ipfs/kubo)   [![godoc][D]](https://godoc.org/github.com/ipfs/kubo)
- [lnd](https://github.com/lightningnetwork/lnd) **star:8161** A complete implementation of a Lightning Network node.   [![There was an update last month][G]](https://github.com/lightningnetwork/lnd)   [![godoc][D]](https://godoc.org/github.com/lightningnetwork/lnd)
- [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) **star:7025** A Framework for Building Public Blockchains in the Cosmos Ecosystem.   [![There was an update last month][G]](https://github.com/cosmos/cosmos-sdk)   [![godoc][D]](https://godoc.org/github.com/cosmos/cosmos-sdk)
- [tendermint](https://github.com/tendermint/tendermint) **star:5865** High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.   [![There was an update last month][G]](https://github.com/tendermint/tendermint)   [![godoc][D]](https://godoc.org/github.com/tendermint/tendermint)
- [solana-go](https://github.com/gagliardetto/solana-go) **star:1567** Go library to interface with Solana JSON RPC and WebSocket interfaces.   [![godoc][D]](https://godoc.org/github.com/gagliardetto/solana-go)
- [gno](https://github.com/gnolang/gno) **star:1075** A comprehensive smart contract suite built with Golang and Gnolang, a deterministic, purpose-built Go variant for blockchains.   [![There was an update last month][G]](https://github.com/gnolang/gno)   [![godoc][D]](https://godoc.org/github.com/gnolang/gno)
- [cometbft](https://github.com/cometbft/cometbft) **star:906** A distributed, Byzantine fault-tolerant, deterministic state machine replication engine. It is a fork of Tendermint Core and implements the Tendermint consensus algorithm.   [![There was an update last month][G]](https://github.com/cometbft/cometbft)   [![godoc][D]](https://godoc.org/github.com/cometbft/cometbft)
- [gossamer](https://github.com/ChainSafe/gossamer) **star:454** A Go implementation of the Polkadot Host.   [![godoc][D]](https://godoc.org/github.com/ChainSafe/gossamer)
- [pactus](https://github.com/pactus-project/pactus) **star:235** A full-node implementation of the Pactus blockchain in Go.   [![There was an update last month][G]](https://github.com/pactus-project/pactus)   [![godoc][D]](https://godoc.org/github.com/pactus-project/pactus)
- [nview](https://github.com/blinklabs-io/nview) **star:40** Local monitoring tool for a Cardano Node. It's a TUI (terminal user interface) designed to fit most screens.   [![There was an update last month][G]](https://github.com/blinklabs-io/nview)   [![godoc][D]](https://godoc.org/github.com/blinklabs-io/nview)
- [gosemble](https://github.com/LimeChain/gosemble) **star:15** A Go-based framework for building Polkadot/Substrate-compatible runtimes.   [![godoc][D]](https://godoc.org/github.com/LimeChain/gosemble)
- [tronlib](https://github.com/kslamph/tronlib) **star:5** A comprehensive, production-ready Go SDK for interacting with the TRON blockchain with TRC20 token support.   [![godoc][D]](https://godoc.org/github.com/kslamph/tronlib)

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) **star:6421** Simple and clean Telegram bot client.   [![godoc][D]](https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api)
- [telebot](https://github.com/tucnak/telebot) **star:4636** Telegram bot framework is written in Go.   [![godoc][D]](https://godoc.org/github.com/tucnak/telebot)
- [wayback](https://github.com/wabarc/wayback) **star:2216** A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages.   [![There was an update last month][G]](https://github.com/wabarc/wayback)   [![godoc][D]](https://godoc.org/github.com/wabarc/wayback)
- [bot](https://github.com/go-telegram/bot) **star:1770** Zero-dependencies Telegram Bot library with additional UI components.   [![godoc][D]](https://godoc.org/github.com/go-telegram/bot)
- [telego](https://github.com/mymmrac/telego) **star:1047** Telegram Bot API library for Golang with full one-to-one API implementation.   [![godoc][D]](https://godoc.org/github.com/mymmrac/telego)
- [arikawa](https://github.com/diamondburned/arikawa) **star:596** A library and framework for the Discord API.   [![godoc][D]](https://godoc.org/github.com/diamondburned/arikawa)
- [echotron](https://github.com/NicoNex/echotron) **star:442** An elegant and concurrent library for Telegram Bots in Go.   [![godoc][D]](https://godoc.org/github.com/NicoNex/echotron)
- [go-joe](https://joe-bot.net)  A general-purpose bot library inspired by Hubot but written in Go.
- [go-twitch-irc](https://github.com/gempir/go-twitch-irc) **star:398** Library to write bots for twitch.tv chat   [![godoc][D]](https://godoc.org/github.com/gempir/go-twitch-irc)
- [go-sarah](https://github.com/oklahomer/go-sarah) **star:271** Framework to build a bot for desired chat services including LINE, Slack, Gitter, and more.   [![godoc][D]](https://godoc.org/github.com/oklahomer/go-sarah)
- [slack-bot](https://github.com/innogames/slack-bot) **star:209** Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github...   [![godoc][D]](https://godoc.org/github.com/innogames/slack-bot)
- [go-tg](https://github.com/mr-linch/go-tg) **star:136** Generated from official docs Go client library for accessing Telegram Bot API, with batteries for building complex bots included.   [![There was an update last month][G]](https://github.com/mr-linch/go-tg)   [![godoc][D]](https://godoc.org/github.com/mr-linch/go-tg)
- [slacker](https://github.com/slack-io/slacker) **star:60** Easy to use framework to create Slack bots.   [![godoc][D]](https://godoc.org/github.com/slack-io/slacker)
- [TG](https://github.com/enetx/tg) **star:52** Telegram Bot Framework for Go.   [![godoc][D]](https://godoc.org/github.com/enetx/tg)
- [micha](https://github.com/onrik/micha) **star:34** Go Library for Telegram bot api.   [![godoc][D]](https://godoc.org/github.com/onrik/micha)
   - [Wisp](https://github.com/wisp-trading/wisp) **star:12** Event-driven trading framework for Go. Spot, perpetual futures, prediction markets. Multi-exchange (Bybit, Hyperliquid, Polymarket).   [![godoc][D]](https://godoc.org/github.com/wisp-trading/wisp)
- [ymsdk](https://github.com/rekurt/ymsdk) **star:7** Go SDK for Yandex Messenger Bot API with type-safe models, automatic retry, and rate-limit handling.   [![godoc][D]](https://godoc.org/github.com/rekurt/ymsdk)

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [air](https://github.com/cosmtrek/air) **star:23788** Air - Live reload for Go apps.   [![There was an update last month][G]](https://github.com/cosmtrek/air)   [![godoc][D]](https://godoc.org/github.com/cosmtrek/air)   [![Contains Chinese documents][CN]](https://github.com/cosmtrek/air)
- [Task](https://github.com/go-task/task) **star:15837** simple "Make" alternative.   [![There was an update last month][G]](https://github.com/go-task/task)   [![godoc][D]](https://godoc.org/github.com/go-task/task)
- [mage](https://github.com/magefile/mage) **star:4681** Mage is a make/rake-like build tool using Go.   [![godoc][D]](https://godoc.org/github.com/magefile/mage)
- [realize](https://github.com/tockins/realize) **star:4439** Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths.   [![It hasn't been updated in recent three years][Y]](https://github.com/tockins/realize)   [![godoc][D]](https://godoc.org/github.com/tockins/realize)
- [mmake](https://github.com/tj/mmake) **star:1741** Modern Make.   [![It hasn't been updated in recent three years][Y]](https://github.com/tj/mmake)   [![godoc][D]](https://godoc.org/github.com/tj/mmake)
- [xc](https://github.com/joerdav/xc) **star:1394** Task runner with README.md defined tasks, executable markdown.   [![godoc][D]](https://godoc.org/github.com/joerdav/xc)
- [goyek](https://github.com/goyek/goyek) **star:693** Create build pipelines in Go.   [![There was an update last month][G]](https://github.com/goyek/goyek)   [![godoc][D]](https://godoc.org/github.com/goyek/goyek)
- [taskctl](https://github.com/taskctl/taskctl) **star:325** Concurrent task runner.   [![godoc][D]](https://godoc.org/github.com/taskctl/taskctl)
- [1build](https://github.com/gopinath-langote/1build) **star:248** Command line tool to frictionlessly manage project-specific commands.   [![godoc][D]](https://godoc.org/github.com/gopinath-langote/1build)
- [gaper](https://github.com/maxclaus/gaper) **star:83** Builds and restarts a Go project when it crashes or some watched file changes.   [![godoc][D]](https://godoc.org/github.com/maxclaus/gaper)
- [gilbert](https://go-gilbert.github.io)  Build system and task runner for Go projects.
- [anko](https://github.com/GuilhermeCaruso/anko) **star:43** Simple application watcher for multiple programming languages.   [![It hasn't been updated in recent three years][Y]](https://github.com/GuilhermeCaruso/anko)   [![godoc][D]](https://godoc.org/github.com/GuilhermeCaruso/anko)
- [gob](https://github.com/kcmvp/gob) **star:15** [Gradle](https://docs.gradle.org/)/[Maven](https://maven.apache.org/) like build tool for Go projects.   [![godoc][D]](https://godoc.org/github.com/kcmvp/gob)
- [rex](https://github.com/rexrun-dev/rex) **star:5** Zero-config universal project runner. Detects your stack (Go, Node, Python, Rust, PHP, Zig, Elixir) and runs the right command.   [![godoc][D]](https://godoc.org/github.com/rexrun-dev/rex)

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [bubbletea](https://github.com/charmbracelet/bubbletea) **star:43743** Go framework to build terminal apps, based on The Elm Architecture.   [![There was an update last month][G]](https://github.com/charmbracelet/bubbletea)   [![godoc][D]](https://godoc.org/github.com/charmbracelet/bubbletea)
- [fx](https://github.com/antonmedv/fx) **star:20534** Terminal JSON viewer & processor.   [![godoc][D]](https://godoc.org/github.com/antonmedv/fx)
- [vhs](https://github.com/charmbracelet/vhs) **star:20353** Your CLI home video recorder - generate terminal GIFs from code for documentation and tutorials.   [![godoc][D]](https://godoc.org/github.com/charmbracelet/vhs)
- [termui](https://github.com/gizak/termui) **star:13575** Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).   [![godoc][D]](https://godoc.org/github.com/gizak/termui)
- [lipgloss](https://github.com/charmbracelet/lipgloss) **star:11564** Declaratively define styles for color, format and layout in the terminal.   [![There was an update last month][G]](https://github.com/charmbracelet/lipgloss)   [![godoc][D]](https://godoc.org/github.com/charmbracelet/lipgloss)
- [gocui](https://github.com/jroimartin/gocui) **star:10595** Minimalist Go library aimed at creating Console User Interfaces.   [![godoc][D]](https://godoc.org/github.com/jroimartin/gocui)
- [gommon/color](https://github.com/labstack/gommon/tree/master/color)  Style terminal text.
- [bubbles](https://github.com/charmbracelet/bubbles) **star:8665** TUI components for bubbletea.   [![There was an update last month][G]](https://github.com/charmbracelet/bubbles)   [![godoc][D]](https://godoc.org/github.com/charmbracelet/bubbles)
- [pterm](https://github.com/pterm/pterm) **star:5501** A library to beautify console output on every platform with many combinable components.   [![There was an update last month][G]](https://github.com/pterm/pterm)   [![godoc][D]](https://godoc.org/github.com/pterm/pterm)
- [go-prompt](https://github.com/c-bata/go-prompt) **star:5488** Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).   [![godoc][D]](https://godoc.org/github.com/c-bata/go-prompt)
- [termbox-go](https://github.com/nsf/termbox-go) **star:4790** Termbox is a library for creating cross-platform text-based interfaces.   [![godoc][D]](https://godoc.org/github.com/nsf/termbox-go)
- [progressbar](https://github.com/schollz/progressbar) **star:4686** Basic thread-safe progress bar that works in every OS.   [![godoc][D]](https://godoc.org/github.com/schollz/progressbar)
- [asciigraph](https://github.com/guptarohit/asciigraph) **star:3071** Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.   [![godoc][D]](https://godoc.org/github.com/guptarohit/asciigraph)
- [termdash](https://github.com/mum4k/termdash) **star:3023** Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).   [![godoc][D]](https://godoc.org/github.com/mum4k/termdash)
- [spinner](https://github.com/briandowns/spinner) **star:2527** Go package to easily provide a terminal spinner with options.   [![godoc][D]](https://godoc.org/github.com/briandowns/spinner)
- [mpb](https://github.com/vbauerster/mpb) **star:2502** Multi progress bar for terminal applications.   [![There was an update last month][G]](https://github.com/vbauerster/mpb)   [![godoc][D]](https://godoc.org/github.com/vbauerster/mpb)
- [uiprogress](https://github.com/gosuri/uiprogress) **star:2141** Flexible library to render progress bars in terminal applications.   [![godoc][D]](https://godoc.org/github.com/gosuri/uiprogress)
- [termenv](https://github.com/muesli/termenv) **star:2014** Advanced ANSI style & color support for your terminal applications.   [![godoc][D]](https://godoc.org/github.com/muesli/termenv)
- [uilive](https://github.com/gosuri/uilive) **star:1727** Library for updating terminal output in real time.   [![It hasn't been updated in recent three years][Y]](https://github.com/gosuri/uilive)   [![godoc][D]](https://godoc.org/github.com/gosuri/uilive)
- [gookit/color](https://github.com/gookit/color) **star:1602** Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.   [![godoc][D]](https://godoc.org/github.com/gookit/color)   [![Contains Chinese documents][CN]](https://github.com/gookit/color)
- [aurora](https://github.com/logrusorgru/aurora) **star:1493** ANSI terminal colors that support fmt.Printf/Sprintf.   [![godoc][D]](https://godoc.org/github.com/logrusorgru/aurora)
- [go-isatty](https://github.com/mattn/go-isatty) **star:920** isatty for golang.   [![There was an update last month][G]](https://github.com/mattn/go-isatty)   [![godoc][D]](https://godoc.org/github.com/mattn/go-isatty)
- [go-colorable](https://github.com/mattn/go-colorable) **star:814** Colorable writer for windows.   [![godoc][D]](https://godoc.org/github.com/mattn/go-colorable)
- [uitable](https://github.com/gosuri/uitable) **star:742** Library to improve readability in terminal apps using tabular data.   [![It hasn't been updated in recent three years][Y]](https://github.com/gosuri/uitable)   [![godoc][D]](https://godoc.org/github.com/gosuri/uitable)
- [box-cli-maker](https://github.com/box-cli-maker/box-cli-maker) **star:645** Render highly customizable boxes in the terminal.   [![godoc][D]](https://godoc.org/github.com/box-cli-maker/box-cli-maker)
- [bubble-table](https://github.com/Evertras/bubble-table) **star:576** An interactive table component for bubbletea.   [![godoc][D]](https://godoc.org/github.com/Evertras/bubble-table)
- [simpletable](https://github.com/alexeyco/simpletable) **star:552** Simple tables in a terminal with Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/alexeyco/simpletable)   [![godoc][D]](https://godoc.org/github.com/alexeyco/simpletable)
- [yacspin](https://github.com/theckman/yacspin) **star:456** Yet Another CLi Spinner package, for working with terminal spinners.   [![It hasn't been updated in recent three years][Y]](https://github.com/theckman/yacspin)   [![godoc][D]](https://godoc.org/github.com/theckman/yacspin)
- [go-tui](https://github.com/grindlemire/go-tui) **star:376** A declarative terminal UI framework with templ-like templates, flexbox layout, and a language server for editor support.   [![There was an update last month][G]](https://github.com/grindlemire/go-tui)   [![godoc][D]](https://godoc.org/github.com/grindlemire/go-tui)
- [tabby](https://github.com/cheynewallace/tabby) **star:371** A tiny library for super simple Golang tables.   [![It hasn't been updated in recent three years][Y]](https://github.com/cheynewallace/tabby)   [![godoc][D]](https://godoc.org/github.com/cheynewallace/tabby)
- [go-colortext](https://github.com/daviddengcn/go-colortext) **star:217** Go library for color output in terminals.   [![It hasn't been updated in recent three years][Y]](https://github.com/daviddengcn/go-colortext)   [![godoc][D]](https://godoc.org/github.com/daviddengcn/go-colortext)
- [loom](https://github.com/loom-go/loom) **star:126** Signal-based reactive components framework for building TUIs.   [![godoc][D]](https://godoc.org/github.com/loom-go/loom)
- [ctc](https://github.com/wzshiming/ctc) **star:55** The non-invasive cross-platform terminal color library does not need to modify the Print method.   [![godoc][D]](https://godoc.org/github.com/wzshiming/ctc)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/ctc)
- [lazyenv](https://github.com/lazynop/lazyenv) **star:55** TUI for browsing, comparing, and editing .env files.   [![godoc][D]](https://godoc.org/github.com/lazynop/lazyenv)
- [marker](https://github.com/cyucelen/marker) **star:54** Easiest way to match and mark strings for colorful terminal outputs.   [![godoc][D]](https://godoc.org/github.com/cyucelen/marker)
- [table](https://github.com/tomlazar/table) **star:52** Small library for terminal color based tables.   [![It hasn't been updated in recent three years][Y]](https://github.com/tomlazar/table)   [![godoc][D]](https://godoc.org/github.com/tomlazar/table)
- [phoenix](https://github.com/phoenix-tui/phoenix) **star:36** High-performance TUI framework with Elm-inspired architecture, perfect Unicode rendering, and zero-allocation event system.   [![godoc][D]](https://godoc.org/github.com/phoenix-tui/phoenix)
- [goscaf](https://github.com/iyashjayesh/goscaf) **star:29** goscaf generates opinionated, production-quality Go project boilerplate via an interactive CLI. Stop copy-pasting skeleton code between projects.   [![godoc][D]](https://godoc.org/github.com/iyashjayesh/goscaf)
- [go-ataman](https://github.com/workanator/go-ataman) **star:20** Go library for rendering ANSI colored text templates in terminals.   [![It hasn't been updated in recent three years][Y]](https://github.com/workanator/go-ataman)   [![godoc][D]](https://godoc.org/github.com/workanator/go-ataman)
- [go-palette](https://github.com/abusomani/go-palette) **star:20** Go library that provides elegant and convenient style definitions using ANSI colors. Fully compatible & wraps the [fmt library](https://pkg.go.dev/fmt) for nice terminal layouts.   [![It hasn't been updated in recent three years][Y]](https://github.com/abusomani/go-palette)   [![godoc][D]](https://godoc.org/github.com/abusomani/go-palette)
- [crab-config-files-templating](https://github.com/alfiankan/crab-config-files-templating) **star:8** Dynamic configuration file templating tool for kubernetes manifest or general configuration files.   [![It hasn't been updated in recent three years][Y]](https://github.com/alfiankan/crab-config-files-templating)   [![godoc][D]](https://godoc.org/github.com/alfiankan/crab-config-files-templating)
- [chroma16](https://github.com/arceus-7/chroma16) **star:5** Generate a harmonious 16-color terminal palette from a single seed color or string.   [![godoc][D]](https://godoc.org/github.com/arceus-7/chroma16)

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [cobra](https://github.com/spf13/cobra) **star:44271** Commander for modern Go CLI interactions.   [![There was an update last month][G]](https://github.com/spf13/cobra)   [![godoc][D]](https://godoc.org/github.com/spf13/cobra)
- [carapace-bin](https://github.com/rsteube/carapace-bin) **star:1888** Multi-shell multi-command argument completer.   [![There was an update last month][G]](https://github.com/rsteube/carapace-bin)   [![godoc][D]](https://godoc.org/github.com/rsteube/carapace-bin)
- [carapace](https://github.com/rsteube/carapace) **star:1368** Command argument completion generator for spf13/cobra.   [![There was an update last month][G]](https://github.com/rsteube/carapace)   [![godoc][D]](https://godoc.org/github.com/rsteube/carapace)
- [complete](https://github.com/posener/complete) **star:956** Write bash completions in Go + Go command bash completion.   [![godoc][D]](https://godoc.org/github.com/posener/complete)
- [argparse](https://github.com/akamensky/argparse) **star:640** Command line argument parser inspired by Python's argparse module.   [![It hasn't been updated in recent three years][Y]](https://github.com/akamensky/argparse)   [![godoc][D]](https://godoc.org/github.com/akamensky/argparse)
- [climax](https://github.com/tucnak/climax) **star:215** Alternative CLI with "human face", in spirit of Go command.   [![It hasn't been updated in recent three years][Y]](https://github.com/tucnak/climax)   [![godoc][D]](https://godoc.org/github.com/tucnak/climax)
- [clîr](https://github.com/leaanthony/clir) **star:201** A Simple and Clear CLI library. Dependency free.   [![godoc][D]](https://godoc.org/github.com/leaanthony/clir)
- [commandeer](https://github.com/jaffee/commandeer) **star:175** Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.   [![It hasn't been updated in recent three years][Y]](https://github.com/jaffee/commandeer)   [![godoc][D]](https://godoc.org/github.com/jaffee/commandeer)
- [acmd](https://github.com/cristalhq/acmd) **star:146** Simple, useful, and opinionated CLI package in Go.   [![godoc][D]](https://godoc.org/github.com/cristalhq/acmd)
- [cmdr](https://github.com/hedzr/cmdr) **star:141** A POSIX/GNU style, getopt-like command-line UI Go library.   [![There was an update last month][G]](https://github.com/hedzr/cmdr)   [![godoc][D]](https://godoc.org/github.com/hedzr/cmdr)
- [command-chain](https://github.com/rainu/go-command-chain) **star:71** A go library for configure and run command chains - such as pipelining in unix shells.   [![godoc][D]](https://godoc.org/github.com/rainu/go-command-chain)
- [boa](https://github.com/GiGurra/boa) **star:48** Declarative flags, env vars, validation, and config files from struct tags. Built on cobra.   [![godoc][D]](https://godoc.org/github.com/GiGurra/boa)
- [cmd](https://github.com/posener/cmd) **star:47** Extends the standard `flag` package to support sub commands and more in idiomatic way.   [![It hasn't been updated in recent three years][Y]](https://github.com/posener/cmd)   [![godoc][D]](https://godoc.org/github.com/posener/cmd)
- [argv](https://github.com/cosiner/argv) **star:43** Go library to split command line string as arguments array using the bash syntax.   [![It hasn't been updated in recent three years][Y]](https://github.com/cosiner/argv)   [![godoc][D]](https://godoc.org/github.com/cosiner/argv)
- [carapace-spec](https://github.com/rsteube/carapace-spec) **star:34** Define simple completions using a spec file.   [![There was an update last month][G]](https://github.com/rsteube/carapace-spec)   [![godoc][D]](https://godoc.org/github.com/rsteube/carapace-spec)
- [console](https://github.com/reeflective/console) Closed-loop application library for Cobra commands, with oh-my-posh prompts, and more.
- [urfave/cli](https://github.com/urfave/cli) **star:24173** Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).   [![godoc][D]](https://godoc.org/github.com/urfave/cli)
- [elvish](https://github.com/elves/elvish) **star:6332** An expressive programming language and a versatile interactive shell.   [![godoc][D]](https://godoc.org/github.com/elves/elvish)
- [kingpin](https://github.com/alecthomas/kingpin) **star:3566** Command line and flag parser supporting sub commands (superseded by `kong`; see below).   [![There was an update last month][G]](https://github.com/alecthomas/kingpin)   [![godoc][D]](https://godoc.org/github.com/alecthomas/kingpin)
- [Dnote](https://github.com/dnote/dnote) **star:3045** A simple command line notebook with multi-device sync.   [![There was an update last month][G]](https://github.com/dnote/dnote)   [![godoc][D]](https://godoc.org/github.com/dnote/dnote)
- [pflag](https://github.com/spf13/pflag) **star:2751** Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.   [![godoc][D]](https://godoc.org/github.com/spf13/pflag)
- [go-flags](https://github.com/jessevdk/go-flags) **star:2698** go command line option parser.   [![godoc][D]](https://godoc.org/github.com/jessevdk/go-flags)
- [go-arg](https://github.com/alexflint/go-arg) **star:2270** Struct-based argument parsing in Go.   [![godoc][D]](https://godoc.org/github.com/alexflint/go-arg)
- [ops](https://github.com/nanovms/ops) **star:1501** Unikernel Builder/Orchestrator.   [![godoc][D]](https://godoc.org/github.com/nanovms/ops)
- [liner](https://github.com/peterh/liner) **star:1094** Go readline-like library for command-line interfaces.   [![It hasn't been updated in recent three years][Y]](https://github.com/peterh/liner)   [![godoc][D]](https://godoc.org/github.com/peterh/liner)
- [flaggy](https://github.com/integrii/flaggy) **star:951** A robust and idiomatic flags package with excellent subcommand support.   [![godoc][D]](https://godoc.org/github.com/integrii/flaggy)
- [mow.cli](https://github.com/jawher/mow.cli) **star:884** Go library for building CLI applications with sophisticated flag and argument parsing and validation.   [![godoc][D]](https://godoc.org/github.com/jawher/mow.cli)
- [mkideal/cli](https://github.com/mkideal/cli) **star:723** Feature-rich and easy to use command-line package based on golang struct tags.   [![godoc][D]](https://godoc.org/github.com/mkideal/cli)
- [wmenu](https://github.com/dixonwille/wmenu) **star:230** Easy to use menu structure for cli applications that prompt users to make choices.   [![godoc][D]](https://godoc.org/github.com/dixonwille/wmenu)
- [sflags](https://github.com/octago/sflags) **star:168** Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin, and other libraries.   [![godoc][D]](https://godoc.org/github.com/octago/sflags)
- [job](https://github.com/liujianping/job) **star:150** JOB, make your short-term command as a long-term job.   [![It hasn't been updated in recent three years][Y]](https://github.com/liujianping/job)
- [readline](https://github.com/reeflective/readline) **star:145** Shell library with modern and easy to use UI features.   [![godoc][D]](https://godoc.org/github.com/reeflective/readline)
- [teris-io/cli](https://github.com/teris-io/cli) **star:132** Simple and complete API for building command line interfaces in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/teris-io/cli)   [![godoc][D]](https://godoc.org/github.com/teris-io/cli)
- [env](https://github.com/codingconcepts/env) **star:125** Tag-based environment configuration for structs.   [![godoc][D]](https://godoc.org/github.com/codingconcepts/env)
- [version](https://github.com/mszostok/version) **star:114** Collects and displays CLI version information in multiple formats along with upgrade notice.   [![It hasn't been updated in recent three years][Y]](https://github.com/mszostok/version)   [![godoc][D]](https://godoc.org/github.com/mszostok/version)
- [gocmd](https://github.com/devfacet/gocmd) **star:67** Go library for building command line applications.   [![It hasn't been updated in recent three years][Y]](https://github.com/devfacet/gocmd)   [![godoc][D]](https://godoc.org/github.com/devfacet/gocmd)
- [wlog](https://github.com/dixonwille/wlog) **star:67** Simple logging interface that supports cross-platform color and concurrency.   [![godoc][D]](https://godoc.org/github.com/dixonwille/wlog)
- [strumt](https://github.com/antham/strumt) **star:63** Library to create prompt chain.   [![godoc][D]](https://godoc.org/github.com/antham/strumt)
- [go-getoptions](https://github.com/DavidGamba/go-getoptions) **star:60** Go option parser inspired by the flexibility of Perl’s GetOpt::Long.   [![godoc][D]](https://godoc.org/github.com/DavidGamba/go-getoptions)
- [flagvar](https://github.com/sgreben/flagvar) **star:48** A collection of flag argument types for Go's standard `flag` package.   [![godoc][D]](https://godoc.org/github.com/sgreben/flagvar)
- [mcli](https://github.com/jxskiss/mcli) **star:47** A minimal but very powerful cli library for Go.   [![godoc][D]](https://godoc.org/github.com/jxskiss/mcli)
- [hashicorp/cli](https://github.com/hashicorp/cli) **star:42** Go library for implementing command-line interfaces.   [![godoc][D]](https://godoc.org/github.com/hashicorp/cli)
- [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli)  cli application framework with auto configuration and dependency injection.
- [go-readline-ny](https://github.com/nyaosorg/go-readline-ny) **star:36** A customizable line-editing library with Emacs keybindings, Unicode support, completion, and syntax highlighting. Used in NYAGOS shell.   [![godoc][D]](https://godoc.org/github.com/nyaosorg/go-readline-ny)
- [orpheus](https://github.com/agilira/orpheus) **star:33** CLI framework with security hardening, plugin storage system, and production observability features.   [![godoc][D]](https://godoc.org/github.com/agilira/orpheus)
- [structcli](https://github.com/leodido/structcli) **star:17** Eliminate Cobra boilerplate: build powerful, feature-rich CLIs declaratively from Go structs.   [![There was an update last month][G]](https://github.com/leodido/structcli)   [![godoc][D]](https://godoc.org/github.com/leodido/structcli)
- [flash-flags](https://github.com/agilira/flash-flags) **star:12** Ultra-fast, zero-dependency, POSIX-compliant flag parsing library that can be used as drop-in stdlib replacement with security hardening.   [![godoc][D]](https://godoc.org/github.com/agilira/flash-flags)
- [memsh](https://github.com/amjadjibon/memsh) **star:11** Virtual bash shell in Go: executes shell commands against an in-memory filesystem (afero), with WASM plugin support and an embeddable HTTP server.   [![There was an update last month][G]](https://github.com/amjadjibon/memsh)   [![godoc][D]](https://godoc.org/github.com/amjadjibon/memsh)
- [subcmd](https://github.com/bobg/subcmd) **star:11** Another approach to parsing and running subcommands. Works alongside the standard `flag` package.   [![godoc][D]](https://godoc.org/github.com/bobg/subcmd)
- [getopt](https://github.com/jon-codes/getopt) **star:10** An accurate Go `getopt`, validated against the GNU libc implementation.   [![godoc][D]](https://godoc.org/github.com/jon-codes/getopt)
- [neuron-cli](https://github.com/steevin/neuron-cli) **star:8** A local-first, Obsidian-compatible terminal knowledge manager.   [![godoc][D]](https://godoc.org/github.com/steevin/neuron-cli)
- [goopt](https://github.com/napalu/goopt) **star:5** A declarative, struct-tag based CLI framework for Go, with a broad feature set such as hierarchical commands/flags, i18n, shell completion, and validation.   [![godoc][D]](https://godoc.org/github.com/napalu/goopt)

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [viper](https://github.com/spf13/viper) **star:30471** Go configuration with fangs.   [![godoc][D]](https://godoc.org/github.com/spf13/viper)
- [godotenv](https://github.com/joho/godotenv) **star:10534** Go port of Ruby's dotenv library (Loads environment variables from `.env`).   [![godoc][D]](https://godoc.org/github.com/joho/godotenv)
- [sonic](https://github.com/bytedance/sonic) **star:9535** A blazingly fast JSON serializing & deserializing library.   [![godoc][D]](https://godoc.org/github.com/bytedance/sonic)   [![Contains Chinese documents][CN]](https://github.com/bytedance/sonic)
- [env](https://github.com/caarlos0/env) **star:6261** Parse environment variables to Go structs (with defaults).   [![godoc][D]](https://godoc.org/github.com/caarlos0/env)
- [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) **star:5457** Go library for managing configuration data from environment variables.   [![godoc][D]](https://godoc.org/github.com/kelseyhightower/envconfig)
- [koanf](https://github.com/knadh/koanf) **star:4114** Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.   [![There was an update last month][G]](https://github.com/knadh/koanf)   [![godoc][D]](https://godoc.org/github.com/knadh/koanf)
- [ini](https://github.com/go-ini/ini) **star:3543** Go package to read and write INI files.   [![godoc][D]](https://godoc.org/github.com/go-ini/ini)
- [kong](https://github.com/alecthomas/kong) **star:3136** Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`).   [![There was an update last month][G]](https://github.com/alecthomas/kong)   [![godoc][D]](https://godoc.org/github.com/alecthomas/kong)
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) **star:2144** Minimalistic configuration reader (from files, ENV, and wherever you want).   [![godoc][D]](https://godoc.org/github.com/ilyakaznacheev/cleanenv)
- [xdg](https://github.com/adrg/xdg) **star:1010** Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).   [![There was an update last month][G]](https://github.com/adrg/xdg)   [![godoc][D]](https://godoc.org/github.com/adrg/xdg)
- [konfig](https://github.com/lalamove/konfig) **star:646** Composable, observable and performant config handling for Go for the distributed processing era.   [![It hasn't been updated in recent three years][Y]](https://github.com/lalamove/konfig)   [![godoc][D]](https://godoc.org/github.com/lalamove/konfig)
- [aconfig](https://github.com/cristalhq/aconfig) **star:637** Simple, useful and opinionated config loader.   [![godoc][D]](https://godoc.org/github.com/cristalhq/aconfig)
- [gookit/config](https://github.com/gookit/config) **star:585** application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.   [![godoc][D]](https://godoc.org/github.com/gookit/config)   [![Contains Chinese documents][CN]](https://github.com/gookit/config)
- [confita](https://github.com/heetch/confita) **star:508** Load configuration in cascade from multiple backends into a struct.   [![godoc][D]](https://godoc.org/github.com/heetch/confita)
- [konf](https://github.com/nil-go/konf) **star:391** The simplest API for reading/watching config from file, env, flag and clouds (e.g. AWS, Azure, GCP).   [![There was an update last month][G]](https://github.com/nil-go/konf)   [![godoc][D]](https://godoc.org/github.com/nil-go/konf)
- [fig](https://github.com/kkyr/fig) **star:383** Tiny library for reading configuration from a file and from environment variables (with validation & defaults).   [![godoc][D]](https://godoc.org/github.com/kkyr/fig)
- [GoLobby/Config](https://github.com/golobby/config) **star:366** GoLobby Config is a lightweight yet powerful configuration manager for the Go programming language.   [![It hasn't been updated in recent three years][Y]](https://github.com/golobby/config)   [![godoc][D]](https://godoc.org/github.com/golobby/config)
- [gone/jconf](https://github.com/One-com/gone/tree/master/jconf)  Modular JSON configuration. Keep your config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
- [hjson](https://github.com/hjson/hjson-go) **star:357** Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.   [![godoc][D]](https://godoc.org/github.com/hjson/hjson-go)
- [config](https://github.com/JeremyLoy/config) **star:339** Cloud native application configuration. Bind ENV to structs in only two lines.   [![It hasn't been updated in recent three years][Y]](https://github.com/JeremyLoy/config)   [![godoc][D]](https://godoc.org/github.com/JeremyLoy/config)
- [envconfig](https://github.com/vrischmann/envconfig) **star:250** Read your configuration from environment variables.   [![godoc][D]](https://godoc.org/github.com/vrischmann/envconfig)
- [zerocfg](https://github.com/chaindead/zerocfg) **star:200** Zero-effort, concise configuration management that avoids boilerplate and repetitive code, supports multiple sources with priority overrides.   [![godoc][D]](https://godoc.org/github.com/chaindead/zerocfg)
- [argus](https://github.com/agilira/argus) **star:170** File watching and configuration management with MPSC ring buffer, adaptive batching strategies, and universal format parsing (JSON, YAML, TOML, INI, HCL, Properties).   [![There was an update last month][G]](https://github.com/agilira/argus)   [![godoc][D]](https://godoc.org/github.com/agilira/argus)
- [harvester](https://github.com/beatlabs/harvester) **star:135** Harvester, a easy to use static and dynamic configuration package supporting seeding, env vars and Consul integration.   [![godoc][D]](https://godoc.org/github.com/beatlabs/harvester)
- [onion](https://github.com/goraz/onion) **star:119** Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.   [![It hasn't been updated in recent three years][Y]](https://github.com/goraz/onion)   [![godoc][D]](https://godoc.org/github.com/goraz/onion)
- [configuration](https://github.com/BoRuDar/configuration) **star:107** Library for initializing configuration structs from env variables, files, flags and 'default' tag.   [![godoc][D]](https://godoc.org/github.com/BoRuDar/configuration)
- [envh](https://github.com/antham/envh) **star:101** Helpers to manage environment variables.   [![godoc][D]](https://godoc.org/github.com/antham/envh)
- [hocon](https://github.com/gurkankaymak/hocon) **star:93** Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files.   [![godoc][D]](https://godoc.org/github.com/gurkankaymak/hocon)
- [configuro](https://github.com/sherifabdlnaby/configuro) **star:92** opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications.   [![There was an update last month][G]](https://github.com/sherifabdlnaby/configuro)   [![godoc][D]](https://godoc.org/github.com/sherifabdlnaby/configuro)
- [env](https://github.com/junk1tm/env) **star:80** A lightweight package for loading environment variables into structs.   [![godoc][D]](https://godoc.org/github.com/junk1tm/env)
- [uConfig](https://github.com/omeid/uconfig) **star:75** Lightweight, zero-dependency, and extendable configuration management.   [![godoc][D]](https://godoc.org/github.com/omeid/uconfig)
- [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) **star:66** Go package that fetches parameters from AWS System Manager - Parameter Store.   [![There was an update last month][G]](https://github.com/PaddleHQ/go-aws-ssm)   [![godoc][D]](https://godoc.org/github.com/PaddleHQ/go-aws-ssm)
- [config](https://github.com/num30/config) **star:61** configure your app using file, environment variables, or flags in two lines of code.   [![godoc][D]](https://godoc.org/github.com/num30/config)
- [config](https://github.com/andreiavrammsd/config)  Struct-based configuration loader with a dedicated config file parser, supporting env vars, flags, defaults, and validation.
- [go-cfg](https://github.com/dsbasko/go-cfg) **star:49** The library provides a unified way to read configuration data into a structure from various sources, such as env, flags, and configuration files (.json, .yaml, .toml, .env).   [![godoc][D]](https://godoc.org/github.com/dsbasko/go-cfg)
- [goconfig](https://github.com/fulldump/goconfig) **star:49** Populate Go structs from flags, environment variables, config.json and defaults with deterministic precedence. No extra dependencies.   [![godoc][D]](https://godoc.org/github.com/fulldump/goconfig)
- [genv](https://github.com/sakirsensoy/genv) **star:45** Read environment variables easily with dotenv support.   [![godoc][D]](https://godoc.org/github.com/sakirsensoy/genv)
- [go-up](https://github.com/ufoscout/go-up) **star:43** A simple configuration library with recursive placeholders resolution and no magic.   [![It hasn't been updated in recent three years][Y]](https://github.com/ufoscout/go-up)   [![godoc][D]](https://godoc.org/github.com/ufoscout/go-up)
- [confiq](https://github.com/greencoda/confiq) **star:41** Structured data format to config struct decoder library for Go - supporting multiple data formats.   [![godoc][D]](https://godoc.org/github.com/greencoda/confiq)
- [enflag](https://github.com/atelpis/enflag) **star:38** Container-oriented, zero-dependency configuration library that unifies Env variable and Flag parsing. Uses generics for type safety, without reflection or struct tags.   [![godoc][D]](https://godoc.org/github.com/atelpis/enflag)
- [conflate](https://github.com/the4thamigo-uk/conflate) **star:36** Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.   [![It hasn't been updated in recent three years][Y]](https://github.com/the4thamigo-uk/conflate)   [![godoc][D]](https://godoc.org/github.com/the4thamigo-uk/conflate)
- [bcl](https://github.com/wkhere/bcl) **star:34** BCL is a configuration language similar to HCL.   [![godoc][D]](https://godoc.org/github.com/wkhere/bcl)
- [go-array](https://github.com/deatil/go-array) **star:23** A Go package that read or set data from map, slice or json.   [![godoc][D]](https://godoc.org/github.com/deatil/go-array)   [![Contains Chinese documents][CN]](https://github.com/deatil/go-array)
- [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) **star:23** Go utility for loading configuration parameters from AWS SSM (Parameter Store).   [![godoc][D]](https://godoc.org/github.com/ianlopshire/go-ssm-config)
- [nasermirzaei89/env](https://github.com/nasermirzaei89/env) **star:23** Simple useful package for read environment variables.   [![godoc][D]](https://godoc.org/github.com/nasermirzaei89/env)
- [envyaml](https://github.com/yuseferi/envyaml) **star:21** Yaml with environment variables reader. it helps to have secrets as environment variable but load them configs as structured Yaml.   [![godoc][D]](https://godoc.org/github.com/yuseferi/envyaml)
- [yamagiconf](https://github.com/romshark/yamagiconf) **star:20** The "safe subset" of YAML for Go configs.   [![There was an update last month][G]](https://github.com/romshark/yamagiconf)   [![godoc][D]](https://godoc.org/github.com/romshark/yamagiconf)
- [go-ini](https://github.com/subpop/go-ini) **star:18** A Go package that marshals and unmarshals INI-files.   [![godoc][D]](https://godoc.org/github.com/subpop/go-ini)
- [ini](https://github.com/wlevene/ini) **star:18** INI Parser & Write Library, Unmarshal to Struct, Marshal to Json, Write File, watch file.   [![godoc][D]](https://godoc.org/github.com/wlevene/ini)
- [go-conf](https://github.com/ThomasObenaus/go-conf) **star:14** Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters.   [![godoc][D]](https://godoc.org/github.com/ThomasObenaus/go-conf)
- [nfigure](https://github.com/muir/nfigure) **star:12** Per-library struct-tag based configuration from command lines (Posix & Go-style); environment, JSON, YAML   [![There was an update last month][G]](https://github.com/muir/nfigure)   [![godoc][D]](https://godoc.org/github.com/muir/nfigure)
- [typenv](https://github.com/diegomarangoni/typenv) **star:12** Minimalistic, zero dependency, typed environment variables library.   [![godoc][D]](https://godoc.org/github.com/diegomarangoni/typenv)
- [gonfig](https://github.com/milad-abbasi/gonfig) **star:10** Tag-based configuration parser which loads values from different providers into typesafe struct.   [![It hasn't been updated in recent three years][Y]](https://github.com/milad-abbasi/gonfig)   [![godoc][D]](https://godoc.org/github.com/milad-abbasi/gonfig)
- [piper](https://github.com/Yiling-J/piper) **star:10** Viper wrapper with config inheritance and key generation.   [![It hasn't been updated in recent three years][Y]](https://github.com/Yiling-J/piper)   [![godoc][D]](https://godoc.org/github.com/Yiling-J/piper)
- [env](https://github.com/syntaqx/env) **star:8** An environment utility package with support for unmarshaling into structs.   [![godoc][D]](https://godoc.org/github.com/syntaqx/env)
- [hedzr/store](https://github.com/hedzr/store) **star:8** Extensible, high-performance configuration management library, optimized for hierarchical data.   [![There was an update last month][G]](https://github.com/hedzr/store)   [![godoc][D]](https://godoc.org/github.com/hedzr/store)
- [swap](https://github.com/oblq/swap) **star:8** Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env).   [![godoc][D]](https://godoc.org/github.com/oblq/swap)
- [azureappconfiguration](https://github.com/Azure/AppConfiguration-GoProvider) **star:6** The configuration provider for consuming data in Azure App Configuration from Go applications.   [![godoc][D]](https://godoc.org/github.com/Azure/AppConfiguration-GoProvider)
- [go-config](https://github.com/MordaTeam/go-config) **star:6** Simple and convenient library for working with app configurations.   [![godoc][D]](https://godoc.org/github.com/MordaTeam/go-config)
- [GoCfg](https://github.com/Jagerente/gocfg) **star:6** Config manager with Struct Tags based contracts, custom value providers, parsers, and documentation generation. Customizable yet simple.   [![godoc][D]](https://godoc.org/github.com/Jagerente/gocfg)
- [go-external-config](https://github.com/go-external-config/go) **star:2** Spring-inspired configuration management library for Go.   [![godoc][D]](https://godoc.org/github.com/go-external-config/go)
- [go-external-config/aws](https://github.com/go-external-config/aws) **star:1** AWS property source support for go-external-config.
- [go-external-config/consul](https://github.com/go-external-config/consul) **star:1** Consul property source support for go-external-config.
- [go-external-config/vault](https://github.com/go-external-config/vault) **star:1** Vault property source support for go-external-config.

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [drone](https://github.com/drone/drone) **star:37305** Drone is a Continuous Integration platform built on Docker, written in Go.   [![There was an update last month][G]](https://github.com/drone/drone)   [![godoc][D]](https://godoc.org/github.com/drone/drone)
- [woodpecker](https://github.com/woodpecker-ci/woodpecker) **star:7439** Woodpecker is a community fork of the Drone CI system.   [![There was an update last month][G]](https://github.com/woodpecker-ci/woodpecker)   [![godoc][D]](https://godoc.org/github.com/woodpecker-ci/woodpecker)
- [CDS](https://github.com/ovh/cds) **star:4832** Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.   [![There was an update last month][G]](https://github.com/ovh/cds)   [![godoc][D]](https://godoc.org/github.com/ovh/cds)
- [muffet](https://github.com/raviqqe/muffet) **star:2612** Fast website link checker in Go, see [alternatives](https://github.com/lycheeverse/lychee#features).   [![There was an update last month][G]](https://github.com/raviqqe/muffet)   [![godoc][D]](https://godoc.org/github.com/raviqqe/muffet)
- [abstruse](https://github.com/bleenco/abstruse) **star:956** Abstruse is a distributed CI platform.   [![godoc][D]](https://godoc.org/github.com/bleenco/abstruse)
- [Bencher](https://bencher.dev/)  A suite of continuous benchmarking tools designed to catch performance regressions in CI.
- [goveralls](https://github.com/mattn/goveralls) **star:796** Go integration for Coveralls.io continuous code coverage tracking system.   [![godoc][D]](https://godoc.org/github.com/mattn/goveralls)
- [gotestfmt](https://github.com/GoTestTools/gotestfmt) **star:597** go test output for humans.   [![It hasn't been updated in recent three years][Y]](https://github.com/GoTestTools/gotestfmt)   [![godoc][D]](https://godoc.org/github.com/GoTestTools/gotestfmt)
- [PikoCI](https://github.com/pikoci/pikoci) **star:156** Self-hosted CI/CD inspired by Concourse. Single binary, any database, any queue. HCL pipelines, pluggable resource types and runners.   [![There was an update last month][G]](https://github.com/pikoci/pikoci)   [![godoc][D]](https://godoc.org/github.com/pikoci/pikoci)
- [overalls](https://github.com/go-playground/overalls) **star:116** Multi-Package go project coverprofile for tools like goveralls.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/overalls)   [![godoc][D]](https://godoc.org/github.com/go-playground/overalls)
- [gomason](https://github.com/nikogura/gomason) **star:69** Test, Build, Sign, and Publish your go binaries from a clean workspace.   [![godoc][D]](https://godoc.org/github.com/nikogura/gomason)
- [dot](https://github.com/opnlabs/dot) **star:36** A minimal, local first continuous integration system that uses Docker to run jobs concurrently in stages.   [![godoc][D]](https://godoc.org/github.com/opnlabs/dot)
- [go-semver-release](https://github.com/s0ders/go-semver-release) **star:26** Automate the semantic versioning of Git repositories.   [![godoc][D]](https://godoc.org/github.com/s0ders/go-semver-release)
- [go-test-coverage](https://github.com/marketplace/actions/go-test-coverage)  A GitHub Action which reports issues when test coverage is below set threshold.
- [go-beautiful-html-coverage](https://github.com/gha-common/go-beautiful-html-coverage) **star:24** A GitHub Action to track code coverage in your pull requests, with a beautiful HTML preview, for free.
- [go-fuzz-action](https://github.com/jidicula/go-fuzz-action) **star:21** Use Go 1.18's built-in fuzz testing in GitHub Actions.
- [roveralls](https://github.com/LawrenceWoodman/roveralls) **star:21** Recursive coverage testing tool.   [![It hasn't been updated in recent three years][Y]](https://github.com/LawrenceWoodman/roveralls)   [![godoc][D]](https://godoc.org/github.com/LawrenceWoodman/roveralls)

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._

- [go-libsass](https://github.com/wellington/go-libsass) **star:216** Go wrapper to the 100% Sass compatible libsass project.   [![Archived][Archived]](https://github.com/wellington/go-libsass)
- [go-css](https://github.com/napsy/go-css) **star:94** A very simple CSS parser, written in Go.   [![godoc][D]](https://godoc.org/github.com/napsy/go-css)

**[⬆ back to top](#contents)**

## Data Integration Frameworks

_Frameworks for performing ELT / ETL_

- [Benthos](https://github.com/benthosdev/benthos) **star:8702** A message streaming bridge between a range of protocols.   [![There was an update last month][G]](https://github.com/benthosdev/benthos)   [![godoc][D]](https://godoc.org/github.com/benthosdev/benthos)
- [CloudQuery](http://github.com/cloudquery/cloudquery)  A high-performance ELT data integration framework with pluggable architecture.
- [omniparser](https://github.com/jf-tech/omniparser) **star:1086** A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema.   [![godoc][D]](https://godoc.org/github.com/jf-tech/omniparser)
- [confluence2md](https://github.com/gkoos/confluence2md) **star:27** Confluence to Markdown crawler and converter.   [![There was an update last month][G]](https://github.com/gkoos/confluence2md)   [![godoc][D]](https://godoc.org/github.com/gkoos/confluence2md)

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [roaring](https://github.com/RoaringBitmap/roaring) **star:2914** Go package implementing compressed bitsets.   [![There was an update last month][G]](https://github.com/RoaringBitmap/roaring)   [![godoc][D]](https://godoc.org/github.com/RoaringBitmap/roaring)
- [binpacker](https://github.com/zhuangsirui/binpacker) **star:234** Binary packer and unpacker helps user build custom binary stream.   [![It hasn't been updated in recent three years][Y]](https://github.com/zhuangsirui/binpacker)   [![godoc][D]](https://godoc.org/github.com/zhuangsirui/binpacker)
- [bit](https://github.com/yourbasic/bit) **star:168** Golang set data structure with bonus bit-twiddling functions.   [![It hasn't been updated in recent three years][Y]](https://github.com/yourbasic/bit)   [![godoc][D]](https://godoc.org/github.com/yourbasic/bit)
- [crunch](https://github.com/superwhiskers/crunch) **star:99** Go package implementing buffers for handling various datatypes easily.   [![It hasn't been updated in recent three years][Y]](https://github.com/superwhiskers/crunch)   [![godoc][D]](https://godoc.org/github.com/superwhiskers/crunch)
- [bingo](https://github.com/iancmcc/bingo) **star:52** Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes.   [![godoc][D]](https://godoc.org/github.com/iancmcc/bingo)
- [go-ef](https://github.com/amallia/go-ef) **star:42** A Go implementation of the Elias-Fano encoding.   [![godoc][D]](https://godoc.org/github.com/amallia/go-ef)

### Bit Sets

- [bitset](https://github.com/bits-and-blooms/bitset) **star:1510** Go package implementing bitsets.   [![There was an update last month][G]](https://github.com/bits-and-blooms/bitset)   [![godoc][D]](https://godoc.org/github.com/bits-and-blooms/bitset)
- [bitmap](https://github.com/kelindar/bitmap) **star:381** Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go.

### Bloom and Cuckoo Filters

- [bloom](https://github.com/bits-and-blooms/bloom) **star:2797** Go package implementing Bloom filters.   [![There was an update last month][G]](https://github.com/bits-and-blooms/bloom)   [![godoc][D]](https://godoc.org/github.com/bits-and-blooms/bloom)
- [boomfilters](https://github.com/tylertreat/BoomFilters) **star:1646** Probabilistic data structures for processing continuous, unbounded streams.   [![godoc][D]](https://godoc.org/github.com/tylertreat/BoomFilters)
- [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) **star:1229** Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.   [![godoc][D]](https://godoc.org/github.com/seiflotfy/cuckoofilter)
- [cuckoo-filter](https://github.com/linvon/cuckoo-filter) **star:306** Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper are available.   [![godoc][D]](https://godoc.org/github.com/linvon/cuckoo-filter)   [![Contains Chinese documents][CN]](https://github.com/linvon/cuckoo-filter)
- [bloom](https://github.com/zhenjl/bloom) **star:147** Bloom filters implemented in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/zhenjl/bloom)   [![godoc][D]](https://godoc.org/github.com/zhenjl/bloom)
- [ring](https://github.com/TheTannerRyan/ring) **star:146** Go implementation of a high performance, thread safe bloom filter.   [![It hasn't been updated in recent three years][Y]](https://github.com/TheTannerRyan/ring)   [![godoc][D]](https://godoc.org/github.com/TheTannerRyan/ring)
- [bloom](https://github.com/yourbasic/bloom) **star:88** Golang Bloom filter implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/yourbasic/bloom)   [![godoc][D]](https://godoc.org/github.com/yourbasic/bloom)
- [bloomfilter](https://github.com/OldPanda/bloomfilter) **star:21** Yet another Bloomfilter implementation in Go, compatible with Java's Guava library.   [![godoc][D]](https://godoc.org/github.com/OldPanda/bloomfilter)
- [ribbonGo](https://github.com/RibbonFilter/ribbonGo) **star:13** First pure Go implementation of Ribbon filters (practically smaller than Bloom and Xor) for space-efficient approximate set membership queries.   [![godoc][D]](https://godoc.org/github.com/RibbonFilter/ribbonGo)

### Data Structure and Algorithm Collections

- [gods](https://github.com/emirpasic/gods) **star:17445** Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.   [![godoc][D]](https://godoc.org/github.com/emirpasic/gods)
- [go-datastructures](https://github.com/Workiva/go-datastructures) **star:7940** Collection of useful, performant, and thread-safe data structures.   [![godoc][D]](https://godoc.org/github.com/Workiva/go-datastructures)
- [gostl](https://github.com/liyue201/gostl) **star:1140** Data structure and algorithm library for go, designed to provide functions similar to C++ STL.   [![godoc][D]](https://godoc.org/github.com/liyue201/gostl)
- [algorithms](https://github.com/shady831213/algorithms) **star:843** Algorithms and data structures.CLRS study.   [![It hasn't been updated in recent three years][Y]](https://github.com/shady831213/algorithms)   [![godoc][D]](https://godoc.org/github.com/shady831213/algorithms)

### Iterators

- [iter](https://github.com/disksing/iter) **star:198** Go implementation of C++ STL iterators and algorithms.   [![It hasn't been updated in recent three years][Y]](https://github.com/disksing/iter)   [![godoc][D]](https://godoc.org/github.com/disksing/iter)   [![Contains Chinese documents][CN]](https://github.com/disksing/iter)
- [gloop](https://github.com/alvii147/gloop) **star:28** Convenient looping using Go's range-over-func feature.   [![godoc][D]](https://godoc.org/github.com/alvii147/gloop)
- [goterator](https://github.com/yaa110/goterator) **star:19** Iterator implementation to provide map and reduce functionalities.   [![godoc][D]](https://godoc.org/github.com/yaa110/goterator)

### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.

- [concurrent-swiss-map](https://github.com/mhmtszr/concurrent-swiss-map) **star:262** A high-performance, thread-safe generic concurrent hash map implementation with Swiss Map.   [![godoc][D]](https://godoc.org/github.com/mhmtszr/concurrent-swiss-map)
- [cmap](https://github.com/lrita/cmap) **star:105** a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards.   [![godoc][D]](https://godoc.org/github.com/lrita/cmap)
- [goradd/maps](https://github.com/goradd/maps) **star:56** Go 1.18+ generic map interface for maps; safe maps; ordered maps; ordered, safe maps; etc.   [![godoc][D]](https://godoc.org/github.com/goradd/maps)
- [dict](https://github.com/srfrog/dict) **star:48** Python-like dictionaries (dict) for Go.   [![godoc][D]](https://godoc.org/github.com/srfrog/dict)
- [go-shelve](https://github.com/lucmq/go-shelve) **star:15** A persistent, map-like object for the Go programming language. Supports multiple embedded key-value stores.   [![godoc][D]](https://godoc.org/github.com/lucmq/go-shelve)
- [hmap](https://github.com/lyonnee/hmap) **star:2** HMap is a concurrent and secure, generic support Map implementation designed to provide an easy-to-use API.   [![godoc][D]](https://godoc.org/github.com/lyonnee/hmap)   [![Contains Chinese documents][CN]](https://github.com/lyonnee/hmap)

### Miscellaneous Data Structures and Algorithms

- [gota](https://github.com/kniren/gota) **star:3271** Implementation of dataframes, series, and data wrangling methods for Go.   [![godoc][D]](https://godoc.org/github.com/kniren/gota)   [![Archived][Archived]](https://github.com/kniren/gota)
- [hyperloglog](https://github.com/axiomhq/hyperloglog) **star:1044** HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.   [![godoc][D]](https://godoc.org/github.com/axiomhq/hyperloglog)
- [go-geoindex](https://github.com/hailocab/go-geoindex) **star:361** In-memory geo index.   [![It hasn't been updated in recent three years][Y]](https://github.com/hailocab/go-geoindex)   [![godoc][D]](https://godoc.org/github.com/hailocab/go-geoindex)
- [go-rquad](https://github.com/aurelien-rainone/go-rquad) **star:141** Region quadtrees with efficient point location and neighbour finding.   [![It hasn't been updated in recent three years][Y]](https://github.com/aurelien-rainone/go-rquad)   [![godoc][D]](https://godoc.org/github.com/aurelien-rainone/go-rquad)
- [gogu](https://github.com/esimov/gogu) **star:111** A comprehensive, reusable and efficient concurrent-safe generics utility functions and data structures library.   [![It hasn't been updated in recent three years][Y]](https://github.com/esimov/gogu)   [![godoc][D]](https://godoc.org/github.com/esimov/gogu)
- [go-rampart](https://github.com/francesconi/go-rampart) **star:107** Determine how intervals relate to each other.   [![godoc][D]](https://godoc.org/github.com/francesconi/go-rampart)
- [go-tuple](https://github.com/barweiss/go-tuple) **star:100** Generic tuple implementation for Go 1.18+.   [![godoc][D]](https://godoc.org/github.com/barweiss/go-tuple)
- [fsm](https://github.com/cocoonspace/fsm) **star:89** Finite-State Machine package.   [![It hasn't been updated in recent three years][Y]](https://github.com/cocoonspace/fsm)   [![godoc][D]](https://godoc.org/github.com/cocoonspace/fsm)
- [go-generics](https://github.com/bobg/go-generics) **star:86** Generic slice, map, set, iterator, and goroutine utilities.   [![godoc][D]](https://godoc.org/github.com/bobg/go-generics)
- [hide](https://github.com/emvi/hide) **star:74** ID type with marshalling to/from hash to prevent sending IDs to clients.   [![It hasn't been updated in recent three years][Y]](https://github.com/emvi/hide)   [![godoc][D]](https://godoc.org/github.com/emvi/hide)
- [count-min-log](https://github.com/seiflotfy/count-min-log) **star:70** Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).   [![godoc][D]](https://godoc.org/github.com/seiflotfy/count-min-log)
- [concurrent-writer](https://github.com/free/concurrent-writer) **star:63** Highly concurrent drop-in replacement for `bufio.Writer`.   [![It hasn't been updated in recent three years][Y]](https://github.com/free/concurrent-writer)   [![godoc][D]](https://godoc.org/github.com/free/concurrent-writer)
- [genfuncs](https://github.com/nwillc/genfuncs) **star:52** Go 1.18+ generics package inspired by Kotlin's Sequence and Map.   [![It hasn't been updated in recent three years][Y]](https://github.com/nwillc/genfuncs)   [![godoc][D]](https://godoc.org/github.com/nwillc/genfuncs)
- [go18ds](https://github.com/daichi-m/go18ds) **star:49** Go Data Structures using Go 1.18 generics.   [![godoc][D]](https://godoc.org/github.com/daichi-m/go18ds)
- [quadtree](https://github.com/s0rg/quadtree) **star:42** Generic, zero-alloc, 100%-test covered quadtree.   [![godoc][D]](https://godoc.org/github.com/s0rg/quadtree)
- [FSM](https://github.com/enetx/fsm) **star:23** FSM for Go.
- [slices](https://github.com/twharmon/slices) **star:23** Pure, generic functions for slices.   [![It hasn't been updated in recent three years][Y]](https://github.com/twharmon/slices)   [![godoc][D]](https://godoc.org/github.com/twharmon/slices)
- [gofal](https://github.com/xxjwxc/gofal) **star:20** fractional api for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/xxjwxc/gofal)   [![godoc][D]](https://godoc.org/github.com/xxjwxc/gofal)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gofal)
- [combo](https://github.com/bobg/combo) **star:4** Combinatorial operations including permutations, combinations, and combinations-with-replacement.   [![godoc][D]](https://godoc.org/github.com/bobg/combo)

### Nullable Types

- [nan](https://github.com/kak-tus/nan) **star:88** Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers.   [![It hasn't been updated in recent three years][Y]](https://github.com/kak-tus/nan)   [![godoc][D]](https://godoc.org/github.com/kak-tus/nan)
- [typ](https://github.com/gurukami/typ) **star:46** Null Types, Safe primitive type conversion and fetching value from complex structures.   [![It hasn't been updated in recent three years][Y]](https://github.com/gurukami/typ)   [![godoc][D]](https://godoc.org/github.com/gurukami/typ)
- [null](https://github.com/emvi/null) **star:40** Nullable Go types that can be marshalled/unmarshalled to/from JSON.   [![It hasn't been updated in recent three years][Y]](https://github.com/emvi/null)   [![godoc][D]](https://godoc.org/github.com/emvi/null)

### Queues

- [hatchet](https://github.com/hatchet-dev/hatchet) **star:7505** Distributed, Fault-tolerant task queue.   [![There was an update last month][G]](https://github.com/hatchet-dev/hatchet)   [![godoc][D]](https://godoc.org/github.com/hatchet-dev/hatchet)
- [deque](https://github.com/gammazero/deque) **star:784** Fast ring-buffer deque (double-ended queue).   [![godoc][D]](https://godoc.org/github.com/gammazero/deque)
- [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) **star:434** Concurrent FIFO queue.   [![It hasn't been updated in recent three years][Y]](https://github.com/enriquebris/goconcurrentqueue)   [![godoc][D]](https://godoc.org/github.com/enriquebris/goconcurrentqueue)
- [queue](https://github.com/adrianbrad/queue) **star:361** Multiple thread-safe, generic queue implementations for Go.   [![There was an update last month][G]](https://github.com/adrianbrad/queue)   [![godoc][D]](https://godoc.org/github.com/adrianbrad/queue)
- [deque](https://github.com/edwingeng/deque) **star:207** A highly optimized double-ended queue.   [![godoc][D]](https://godoc.org/github.com/edwingeng/deque)
- [memlog](https://github.com/embano1/memlog) **star:142** An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka.   [![There was an update last month][G]](https://github.com/embano1/memlog)   [![godoc][D]](https://godoc.org/github.com/embano1/memlog)
- [deheap](https://github.com/aalpar/deheap) **star:19** Doubly-ended heap (min-max heap) with O(log n) access to both minimum and maximum elements.   [![godoc][D]](https://godoc.org/github.com/aalpar/deheap)
- [list](https://github.com/koss-null/list) **star:15** A generic, thread-safe doubly linked list with full iterator support and an intrusive singly linked list for embedded use; a feature-rich replacement for container/list.   [![godoc][D]](https://godoc.org/github.com/koss-null/list)
- [dqueue](https://github.com/vodolaz095/dqueue) **star:5** Simple, in memory, zero dependency and battle tested, thread-safe deferred queue.   [![godoc][D]](https://godoc.org/github.com/vodolaz095/dqueue)

### Sets

- [golang-set](https://github.com/deckarep/golang-set) **star:4704** Thread-Safe and Non-Thread-Safe high-performance sets for Go.   [![godoc][D]](https://godoc.org/github.com/deckarep/golang-set)
- [goset](https://github.com/zoumo/goset) **star:53** A useful Set collection implementation for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/zoumo/goset)   [![godoc][D]](https://godoc.org/github.com/zoumo/goset)
- [set](https://github.com/StudioSol/set) **star:30** Simple set data structure implementation in Go using LinkedHashMap.   [![godoc][D]](https://godoc.org/github.com/StudioSol/set)
- [dsu](https://github.com/ihebu/dsu) **star:19** Disjoint Set data structure implementation in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ihebu/dsu)   [![godoc][D]](https://godoc.org/github.com/ihebu/dsu)

### Text Analysis

- [bleve](https://github.com/blevesearch/bleve) **star:11148** Modern text indexing library for go.   [![There was an update last month][G]](https://github.com/blevesearch/bleve)   [![godoc][D]](https://godoc.org/github.com/blevesearch/bleve)
- [trie](https://github.com/derekparker/trie) **star:791** Trie implementation in Go.   [![godoc][D]](https://godoc.org/github.com/derekparker/trie)
- [go-edlib](https://github.com/hbollon/go-edlib) **star:605** Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode.   [![godoc][D]](https://godoc.org/github.com/hbollon/go-edlib)
- [levenshtein](https://github.com/agnivade/levenshtein) **star:472** Implementation to calculate levenshtein distance in Go.   [![godoc][D]](https://godoc.org/github.com/agnivade/levenshtein)
- [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) **star:414** Go implementation of Adaptive Radix Tree.   [![godoc][D]](https://godoc.org/github.com/plar/go-adaptive-radix-tree)
- [levenshtein](https://github.com/agext/levenshtein) **star:92** Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.   [![It hasn't been updated in recent three years][Y]](https://github.com/agext/levenshtein)   [![godoc][D]](https://godoc.org/github.com/agext/levenshtein)
- [ptrie](https://github.com/viant/ptrie) **star:47** An implementation of prefix tree.   [![godoc][D]](https://godoc.org/github.com/viant/ptrie)
- [radixtree](https://github.com/gammazero/radixtree) **star:42** Adaptive radix tree (prefix-tree or compact-trie).   [![godoc][D]](https://godoc.org/github.com/gammazero/radixtree)
- [mspm](https://github.com/BlackRabbitt/mspm) **star:30** Multi-String Pattern Matching Algorithm for information retrieval.   [![It hasn't been updated in recent three years][Y]](https://github.com/BlackRabbitt/mspm)   [![godoc][D]](https://godoc.org/github.com/BlackRabbitt/mspm)
- [parsefields](https://github.com/MonaxGT/parsefields) **star:9** Tools for parse JSON-like logs for collecting unique fields and events.   [![It hasn't been updated in recent three years][Y]](https://github.com/MonaxGT/parsefields)   [![godoc][D]](https://godoc.org/github.com/MonaxGT/parsefields)

### Trees

- [skiplist](https://github.com/MauriceGit/skiplist) **star:297** Very fast Go Skiplist implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/MauriceGit/skiplist)   [![godoc][D]](https://godoc.org/github.com/MauriceGit/skiplist)
- [skiplist](https://github.com/gansidui/skiplist) **star:84** Skiplist implementation in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/gansidui/skiplist)   [![godoc][D]](https://godoc.org/github.com/gansidui/skiplist)
- [treemap](https://github.com/igrmk/treemap) **star:69** Generic key-sorted map using a red-black tree under the hood.   [![It hasn't been updated in recent three years][Y]](https://github.com/igrmk/treemap)   [![godoc][D]](https://godoc.org/github.com/igrmk/treemap)
- [merkle](https://github.com/bobg/merkle) **star:24** Space-efficient computation of Merkle root hashes and inclusion proofs.   [![godoc][D]](https://godoc.org/github.com/bobg/merkle)
- [graphlib](https://github.com/aio-arch/graphlib) **star:5** Topological sort lib,Sorting and pruning of DAG graphs.   [![godoc][D]](https://godoc.org/github.com/aio-arch/graphlib)
- [hashsplit](http://github.com/bobg/hashsplit)  Split byte streams into chunks, and arrange chunks into trees, with boundaries determined by content, not position.

### Pipes

- [pipeline](https://github.com/hyfather/pipeline) **star:62** An implementation of pipelines with fan-in and fan-out.   [![It hasn't been updated in recent three years][Y]](https://github.com/hyfather/pipeline)   [![godoc][D]](https://godoc.org/github.com/hyfather/pipeline)
- [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) **star:47** Go module that processes work concurrently and returns output in a channel in the order of input.   [![It hasn't been updated in recent three years][Y]](https://github.com/tejzpr/ordered-concurrently)   [![godoc][D]](https://godoc.org/github.com/tejzpr/ordered-concurrently)
- [parapipe](https://github.com/nazar256/parapipe) **star:39** FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results.   [![godoc][D]](https://godoc.org/github.com/nazar256/parapipe)
- [pipelines](https://github.com/nxdir-s/pipelines) **star:20** Generic pipeline functions for concurrent processing.   [![godoc][D]](https://godoc.org/github.com/nxdir-s/pipelines)

**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [BigCache](https://github.com/allegro/bigcache) **star:8148** Efficient key/value cache for gigabytes of data.   [![godoc][D]](https://godoc.org/github.com/allegro/bigcache)
- [GCache](https://github.com/bluele/gcache) **star:2732** Cache library with support for expirable Cache, LFU, LRU and ARC.   [![godoc][D]](https://godoc.org/github.com/bluele/gcache)
- [fastcache](https://github.com/VictoriaMetrics/fastcache) **star:2372** fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.   [![godoc][D]](https://godoc.org/github.com/VictoriaMetrics/fastcache)
- [cache2go](https://github.com/muesli/cache2go) **star:2155** In-memory key:value cache which supports automatic invalidation based on timeouts.   [![godoc][D]](https://godoc.org/github.com/muesli/cache2go)
- [EchoVault](https://github.com/EchoVault/EchoVault) **star:539** Embeddable Distributed in-memory data store compatible with Redis clients.   [![godoc][D]](https://godoc.org/github.com/EchoVault/EchoVault)
- [cachego](https://github.com/faabiosr/cachego) **star:370** Golang Cache component for multiple drivers.   [![godoc][D]](https://godoc.org/github.com/faabiosr/cachego)   [![Archived][Archived]](https://github.com/faabiosr/cachego)
- [bcache](https://github.com/iwanbk/bcache) **star:165** Eventually consistent distributed in-memory cache Go library.   [![It hasn't been updated in recent three years][Y]](https://github.com/iwanbk/bcache)   [![godoc][D]](https://godoc.org/github.com/iwanbk/bcache)
- [go-cache](https://github.com/viney-shih/go-cache) **star:162** A flexible multi-layer Go caching library to deal with in-memory and shared cache by adopting Cache-Aside pattern.   [![godoc][D]](https://godoc.org/github.com/viney-shih/go-cache)
- [couchcache](https://github.com/codingsince1985/couchcache) **star:65** RESTful caching micro-service backed by Couchbase server.   [![godoc][D]](https://godoc.org/github.com/codingsince1985/couchcache)
- [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) **star:45** BigCache with clustering support and individual item expiration.   [![It hasn't been updated in recent three years][Y]](https://github.com/oaStuff/clusteredBigCache)   [![godoc][D]](https://godoc.org/github.com/oaStuff/clusteredBigCache)
- [coherence-go-client](https://github.com/oracle/coherence-go-client) **star:15** Full implementation of Oracle Coherence cache API for Go applications using gRPC as network transport.   [![godoc][D]](https://godoc.org/github.com/oracle/coherence-go-client)
- [gdcache](https://github.com/ulovecode/gdcache) **star:14** A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache.   [![It hasn't been updated in recent three years][Y]](https://github.com/ulovecode/gdcache)   [![godoc][D]](https://godoc.org/github.com/ulovecode/gdcache)   [![Contains Chinese documents][CN]](https://github.com/ulovecode/gdcache)
- [go-freelru](https://github.com/elastic/go-freelru) A GC-less, fast and generic LRU hashmap library with optional locking, sharding, eviction and expiration.
- [groupcache](https://github.com/golang/groupcache) **star:13343** Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.   [![godoc][D]](https://godoc.org/github.com/golang/groupcache)
- [ristretto](https://github.com/dgraph-io/ristretto) **star:6964** A high performance memory-bound Go cache.   [![There was an update last month][G]](https://github.com/dgraph-io/ristretto)   [![godoc][D]](https://godoc.org/github.com/dgraph-io/ristretto)
- [gocache](https://github.com/eko/gocache) **star:2873** A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.   [![godoc][D]](https://godoc.org/github.com/eko/gocache)
- [otter](https://github.com/maypok86/otter) **star:2640** A high performance lockless cache for Go. Many times faster than Ristretto and friends.   [![godoc][D]](https://godoc.org/github.com/maypok86/otter)
- [sturdyc](https://github.com/viccon/sturdyc) **star:1268** A caching library with advanced concurrency features designed to make I/O heavy applications robust and highly performant.   [![godoc][D]](https://godoc.org/github.com/viccon/sturdyc)
- [ttlcache](https://github.com/jellydator/ttlcache) **star:1267** An in-memory cache with item expiration and generics.   [![There was an update last month][G]](https://github.com/jellydator/ttlcache)   [![godoc][D]](https://godoc.org/github.com/jellydator/ttlcache)
- [jetcache-go](https://github.com/mgtv-tech/jetcache-go) **star:504** Unified Go cache library supporting multi-level caching.   [![godoc][D]](https://godoc.org/github.com/mgtv-tech/jetcache-go)   [![Contains Chinese documents][CN]](https://github.com/mgtv-tech/jetcache-go)
- [theine](https://github.com/Yiling-J/theine-go) **star:375** High performance, near optimal in-memory cache with proactive TTL expiration and generics.   [![godoc][D]](https://godoc.org/github.com/Yiling-J/theine-go)
- [pocache](https://github.com/naughtygopher/pocache) **star:236** Pocache is a minimal cache package which focuses on a preemptive optimistic caching strategy.   [![godoc][D]](https://godoc.org/github.com/naughtygopher/pocache)
- [imcache](https://github.com/erni27/imcache) **star:123** A generic in-memory cache Go library. It supports expiration, sliding expiration, max entries limit, eviction callbacks and sharding.   [![godoc][D]](https://godoc.org/github.com/erni27/imcache)
- [go-mcache](https://github.com/OrlovEvgeny/go-mcache) **star:107** Fast in-memory key:value store/cache library. Pointer caches.   [![godoc][D]](https://godoc.org/github.com/OrlovEvgeny/go-mcache)
- [timedmap](https://github.com/zekroTJA/timedmap) **star:75** Map with expiring key-value pairs.   [![godoc][D]](https://godoc.org/github.com/zekroTJA/timedmap)
- [icache](https://github.com/mdaliyan/icache) **star:23** A High Performance, Generic, thread-safe, zero-dependency cache package.   [![godoc][D]](https://godoc.org/github.com/mdaliyan/icache)
- [gocache](https://github.com/yuseferi/gocache) **star:21** A data race free Go ache library with high performance and auto pruge functionality   [![godoc][D]](https://godoc.org/github.com/yuseferi/gocache)
- [nscache](https://github.com/no-src/nscache) **star:15** A Go caching framework that supports multiple data source drivers.   [![godoc][D]](https://godoc.org/github.com/no-src/nscache)
- [go-gcache](https://github.com/szyhf/go-gcache) **star:12** The generic version of `GCache`, cache support for expirable Cache, LFU, LRU and ARC.   [![godoc][D]](https://godoc.org/github.com/szyhf/go-gcache)
- [ttlcache](https://github.com/cheshir/ttlcache) **star:10** In-memory key value storage with TTL for each record.   [![It hasn't been updated in recent three years][Y]](https://github.com/cheshir/ttlcache)   [![godoc][D]](https://godoc.org/github.com/cheshir/ttlcache)

### Databases Implemented in Go

- [cockroach](https://github.com/cockroachdb/cockroach) **star:32296** Scalable, Geo-Replicated, Transactional Datastore.   [![There was an update last month][G]](https://github.com/cockroachdb/cockroach)   [![godoc][D]](https://godoc.org/github.com/cockroachdb/cockroach)
- [dolt](https://github.com/dolthub/dolt) **star:23859** Dolt – It's Git for Data.   [![There was an update last month][G]](https://github.com/dolthub/dolt)   [![godoc][D]](https://godoc.org/github.com/dolthub/dolt)
- [dgraph](https://github.com/dgraph-io/dgraph) **star:21739** Scalable, Distributed, Low Latency, High Throughput Graph Database.   [![There was an update last month][G]](https://github.com/dgraph-io/dgraph)   [![godoc][D]](https://godoc.org/github.com/dgraph-io/dgraph)
- [badger](https://github.com/dgraph-io/badger) **star:15711** Fast key-value store in Go.   [![There was an update last month][G]](https://github.com/dgraph-io/badger)   [![godoc][D]](https://godoc.org/github.com/dgraph-io/badger)
- [DiceDB](https://github.com/DiceDB/dice) **star:10769** An open-source, fast, reactive, in-memory database optimized for modern hardware. Higher throughput and lower median latencies, making it ideal for modern workloads.
- [bbolt](https://github.com/etcd-io/bbolt) **star:9632** An embedded key/value database for Go.   [![There was an update last month][G]](https://github.com/etcd-io/bbolt)   [![godoc][D]](https://godoc.org/github.com/etcd-io/bbolt)
- [Bitcask](https://git.mills.io/prologic/bitcask)  Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).
- [buntdb](https://github.com/tidwall/buntdb) **star:4860** Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.   [![godoc][D]](https://godoc.org/github.com/tidwall/buntdb)
- [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) **star:1528** CovenantSQL is a SQL database on blockchain.   [![It hasn't been updated in recent three years][Y]](https://github.com/CovenantSQL/CovenantSQL)   [![godoc][D]](https://godoc.org/github.com/CovenantSQL/CovenantSQL)
- [column](https://github.com/kelindar/column) **star:1512** High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions.   [![godoc][D]](https://godoc.org/github.com/kelindar/column)
- [Databunker](https://github.com/paranoidguy/databunker) **star:1474** Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.   [![There was an update last month][G]](https://github.com/paranoidguy/databunker)   [![godoc][D]](https://godoc.org/github.com/paranoidguy/databunker)
- [diskv](https://github.com/peterbourgon/diskv) **star:1457** Home-grown disk-backed key-value store.   [![It hasn't been updated in recent three years][Y]](https://github.com/peterbourgon/diskv)   [![godoc][D]](https://godoc.org/github.com/peterbourgon/diskv)
- [eliasdb](https://github.com/krotik/eliasdb) **star:1037** Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.   [![It hasn't been updated in recent three years][Y]](https://github.com/krotik/eliasdb)   [![godoc][D]](https://godoc.org/github.com/krotik/eliasdb)
- [clover](https://github.com/ostafen/clover) **star:825** A lightweight document-oriented NoSQL database written in pure Golang.   [![godoc][D]](https://godoc.org/github.com/ostafen/clover)
- [Coffer](https://github.com/claygod/coffer) **star:41** Simple ACID key-value database that supports transactions.   [![godoc][D]](https://godoc.org/github.com/claygod/coffer)
- [gedb](https://github.com/vinicius-lino-figueiredo/gedb) **star:23** MongoDB-like embedded database, written in pure-go. Supports indexing and complex queries.   [![godoc][D]](https://godoc.org/github.com/vinicius-lino-figueiredo/gedb)
- [go-sqlite](https://github.com/glebarez/go-sqlite) – A Pure Golang implemented SQLite driver without CGO.
- [prometheus](https://github.com/prometheus/prometheus) **star:65217** Monitoring system and time series database.   [![There was an update last month][G]](https://github.com/prometheus/prometheus)   [![godoc][D]](https://godoc.org/github.com/prometheus/prometheus)
- [Milvus](https://github.com/milvus-io/milvus) **star:45224** Milvus is a vector database for embedding management, analytics and search.   [![There was an update last month][G]](https://github.com/milvus-io/milvus)   [![godoc][D]](https://godoc.org/github.com/milvus-io/milvus)
- [tidb](https://github.com/pingcap/tidb) **star:40293** TiDB is a distributed SQL database. Inspired by the design of Google F1.   [![There was an update last month][G]](https://github.com/pingcap/tidb)   [![godoc][D]](https://godoc.org/github.com/pingcap/tidb)   [![Contains Chinese documents][CN]](https://github.com/pingcap/tidb)
- [influxdb](https://github.com/influxdb/influxdb) **star:31642** Scalable datastore for metrics, events, and real-time analytics.   [![There was an update last month][G]](https://github.com/influxdb/influxdb)
- [rqlite](https://github.com/rqlite/rqlite) **star:17627** The lightweight, distributed, relational database built on SQLite.   [![There was an update last month][G]](https://github.com/rqlite/rqlite)   [![godoc][D]](https://godoc.org/github.com/rqlite/rqlite)
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) **star:17330** fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.   [![There was an update last month][G]](https://github.com/VictoriaMetrics/VictoriaMetrics)   [![godoc][D]](https://godoc.org/github.com/VictoriaMetrics/VictoriaMetrics)
- [immudb](https://github.com/codenotary/immudb) **star:9003** immudb is a lightweight, high-speed immutable database for systems and applications written in Go.   [![There was an update last month][G]](https://github.com/codenotary/immudb)   [![godoc][D]](https://godoc.org/github.com/codenotary/immudb)
- [goleveldb](https://github.com/syndtr/goleveldb) **star:6320** Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.   [![godoc][D]](https://godoc.org/github.com/syndtr/goleveldb)
- [pebble](https://github.com/cockroachdb/pebble) **star:5956** RocksDB/LevelDB inspired key-value database in Go.   [![There was an update last month][G]](https://github.com/cockroachdb/pebble)   [![godoc][D]](https://godoc.org/github.com/cockroachdb/pebble)
- [rosedb](https://github.com/roseduan/rosedb) **star:4880** An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset.   [![godoc][D]](https://godoc.org/github.com/roseduan/rosedb)   [![Contains Chinese documents][CN]](https://github.com/roseduan/rosedb)
- [redka](https://github.com/nalgeon/redka) **star:4566** Redis re-implemented with SQLite.   [![godoc][D]](https://godoc.org/github.com/nalgeon/redka)
- [ledisdb](https://github.com/siddontang/ledisdb) **star:4120** Ledisdb is a high performance NoSQL like Redis based on LevelDB.   [![godoc][D]](https://godoc.org/github.com/siddontang/ledisdb)
- [godis](https://github.com/hdt3213/godis) **star:3836** A Golang implemented high-performance Redis server and cluster.   [![godoc][D]](https://godoc.org/github.com/hdt3213/godis)   [![Contains Chinese documents][CN]](https://github.com/hdt3213/godis)
- [nutsdb](https://github.com/xujiajun/nutsdb) **star:3572** Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set.   [![godoc][D]](https://godoc.org/github.com/xujiajun/nutsdb)   [![Contains Chinese documents][CN]](https://github.com/xujiajun/nutsdb)
- [LinDB](https://github.com/lindb/lindb) **star:3069** LinDB is a scalable, high performance, high availability distributed time series database.   [![godoc][D]](https://godoc.org/github.com/lindb/lindb)   [![Contains Chinese documents][CN]](https://github.com/lindb/lindb)
- [tiedot](https://github.com/HouzuoGuo/tiedot) **star:2728** Your NoSQL database powered by Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/HouzuoGuo/tiedot)   [![godoc][D]](https://godoc.org/github.com/HouzuoGuo/tiedot)
- [lotusdb](https://github.com/flower-corp/lotusdb) **star:2253** Fast k/v database compatible with lsm and b+tree.   [![godoc][D]](https://godoc.org/github.com/flower-corp/lotusdb)
- [pogreb](https://github.com/akrylysov/pogreb) **star:1350** Embedded key-value store for read-heavy workloads.   [![godoc][D]](https://godoc.org/github.com/akrylysov/pogreb)
- [objectbox-go](https://github.com/objectbox/objectbox-go) **star:1272** High-performance embedded Object Database (NoSQL) with Go API.   [![godoc][D]](https://godoc.org/github.com/objectbox/objectbox-go)
- [moss](https://github.com/couchbase/moss) **star:1016** Moss is a simple LSM key-value storage engine written in 100% Go.   [![godoc][D]](https://godoc.org/github.com/couchbase/moss)
- [NornicDB](https://github.com/orneryd/NornicDB) **star:826** High performance graph + vector database (Neo4j and qDrant compatible), focused on low latency graph-rag retreival for AI systems.    [![There was an update last month][G]](https://github.com/orneryd/NornicDB)   [![godoc][D]](https://godoc.org/github.com/orneryd/NornicDB)
- [levigo](https://github.com/jmhodges/levigo) **star:420** Levigo is a Go wrapper for LevelDB.   [![It hasn't been updated in recent three years][Y]](https://github.com/jmhodges/levigo)   [![godoc][D]](https://godoc.org/github.com/jmhodges/levigo)
- [pudge](https://github.com/recoilme/pudge) **star:370** Fast and simple key/value store written using Go's standard library.   [![It hasn't been updated in recent three years][Y]](https://github.com/recoilme/pudge)   [![godoc][D]](https://godoc.org/github.com/recoilme/pudge)
- [lynxdb](https://github.com/lynxbase/lynxdb) **star:281** Lightweight columnar log analytics database with a pipe-style query language inspired by SPL.   [![godoc][D]](https://godoc.org/github.com/lynxbase/lynxdb)
- [Vasto](https://github.com/chrislusf/vasto) **star:263** A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.   [![It hasn't been updated in recent three years][Y]](https://github.com/chrislusf/vasto)   [![godoc][D]](https://godoc.org/github.com/chrislusf/vasto)
- [piladb](https://github.com/fern4lvarez/piladb) **star:206** Lightweight RESTful database engine based on stack data structures.   [![godoc][D]](https://godoc.org/github.com/fern4lvarez/piladb)
- [libradb](https://github.com/amit-davidson/LibraDB) **star:203** LibraDB is a simple database with less than 1000 lines of code for learning.   [![godoc][D]](https://godoc.org/github.com/amit-davidson/LibraDB)
- [nanotdb](https://github.com/aymanhs/nanotdb) **star:136** A lightweight, zero-dependency, append-only Time-Series Database and Dashboard optimized for low-power hardware.   [![godoc][D]](https://godoc.org/github.com/aymanhs/nanotdb)
- [unitdb](https://github.com/unit-io/unitdb) **star:124** Fast timeseries database for IoT, realtime messaging applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application.   [![It hasn't been updated in recent three years][Y]](https://github.com/unit-io/unitdb)   [![godoc][D]](https://godoc.org/github.com/unit-io/unitdb)
- [hare](https://github.com/jameycribbs/hare) **star:97** A simple database management system that stores each table as a text file of line-delimited JSON.   [![It hasn't been updated in recent three years][Y]](https://github.com/jameycribbs/hare)   [![godoc][D]](https://godoc.org/github.com/jameycribbs/hare)
- [minisql](https://github.com/RichardKnop/minisql) **star:47** Embedded single file SQL database.   [![godoc][D]](https://godoc.org/github.com/RichardKnop/minisql)
- [rotom](https://github.com/xgzlucario/rotom) **star:42** A tiny Redis server built with Golang, compatible with RESP protocols.   [![godoc][D]](https://godoc.org/github.com/xgzlucario/rotom)
- [tempdb](https://github.com/rafaeljesus/tempdb) **star:19** Key-value store for temporary items.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafaeljesus/tempdb)   [![godoc][D]](https://godoc.org/github.com/rafaeljesus/tempdb)
- [NoKV](https://github.com/feichai0017/NoKV) **star:5** Native metadata service for distributed filesystems, object storage, and AI dataset workloads.
- 
### Database Schema Migration

- [migrate](https://github.com/golang-migrate/migrate) **star:18708** Database migrations. CLI and Golang library.   [![godoc][D]](https://godoc.org/github.com/golang-migrate/migrate)
- [bytebase](https://github.com/bytebase/bytebase) **star:14264** Safe database schema change and version control for DevOps teams.   [![There was an update last month][G]](https://github.com/bytebase/bytebase)   [![godoc][D]](https://godoc.org/github.com/bytebase/bytebase)
- [goose](https://github.com/pressly/goose) **star:11167** Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.   [![There was an update last month][G]](https://github.com/pressly/goose)   [![godoc][D]](https://godoc.org/github.com/pressly/goose)
- [atlas](https://github.com/ariga/atlas) **star:8575** A Database Toolkit. A CLI designed to help companies better work with their data.   [![godoc][D]](https://godoc.org/github.com/ariga/atlas)
- [dbmate](https://github.com/amacneil/dbmate) **star:6996** A lightweight, framework-agnostic database migration tool.   [![There was an update last month][G]](https://github.com/amacneil/dbmate)   [![godoc][D]](https://godoc.org/github.com/amacneil/dbmate)
- [sql-migrate](https://github.com/rubenv/sql-migrate) **star:3414** Database migration tool. Allows embedding migrations into the application using go-bindata.   [![There was an update last month][G]](https://github.com/rubenv/sql-migrate)   [![godoc][D]](https://godoc.org/github.com/rubenv/sql-migrate)
- [skeema](https://github.com/skeema/skeema) **star:1371** Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.   [![godoc][D]](https://godoc.org/github.com/skeema/skeema)
- [soda](https://github.com/gobuffalo/pop/tree/master/soda)  Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
- [gormigrate](https://github.com/go-gormigrate/gormigrate) **star:1173** Database schema migration helper for Gorm ORM.   [![godoc][D]](https://godoc.org/github.com/go-gormigrate/gormigrate)
- [goavro](https://github.com/linkedin/goavro) **star:1065** A Go package that encodes and decodes Avro data.   [![godoc][D]](https://godoc.org/github.com/linkedin/goavro)
- [migrator](https://github.com/lopezator/migrator) **star:178** Dead simple Go database migration library.   [![godoc][D]](https://godoc.org/github.com/lopezator/migrator)
- [darwin](https://github.com/GuiaBolso/darwin) **star:150** Database schema evolution library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/GuiaBolso/darwin)   [![godoc][D]](https://godoc.org/github.com/GuiaBolso/darwin)
- [sqlize](https://github.com/sunary/sqlize) **star:125** Database migration generator. Allows generate sql migration from model and existing sql by differ them.   [![godoc][D]](https://godoc.org/github.com/sunary/sqlize)
- [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) **star:87** A Go package to help write migrations with go-pg/pg.   [![godoc][D]](https://godoc.org/github.com/robinjoseph08/go-pg-migrations)
- [avro](https://github.com/khezen/avro) **star:49** Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.   [![godoc][D]](https://godoc.org/github.com/khezen/avro)
- [schema](https://github.com/adlio/schema) **star:44** Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.   [![godoc][D]](https://godoc.org/github.com/adlio/schema)
- [go-fixtures](https://github.com/RichardKnop/go-fixtures) **star:32** Django style fixtures for Golang's excellent built-in database/sql library.   [![It hasn't been updated in recent three years][Y]](https://github.com/RichardKnop/go-fixtures)   [![godoc][D]](https://godoc.org/github.com/RichardKnop/go-fixtures)
- [migrator](https://github.com/larapulse/migrator) **star:23** MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code.   [![godoc][D]](https://godoc.org/github.com/larapulse/migrator)
- [db-migrator.go](https://github.com/raoptimus/db-migrator.go) **star:19** CLI for versioned database schema migrations with support for PostgreSQL, MySQL, ClickHouse, Tarantool, and Apache Iceberg.   [![There was an update last month][G]](https://github.com/raoptimus/db-migrator.go)   [![godoc][D]](https://godoc.org/github.com/raoptimus/db-migrator.go)
- [libschema](https://github.com/muir/libschema) **star:19** Define your migrations separately in each library. Migrations for open source libraries. MySQL & PostgreSQL.   [![godoc][D]](https://godoc.org/github.com/muir/libschema)
- [gorm-seeder](https://github.com/Kachit/gorm-seeder) **star:18** Simple database seeder for Gorm ORM.   [![It hasn't been updated in recent three years][Y]](https://github.com/Kachit/gorm-seeder)   [![godoc][D]](https://godoc.org/github.com/Kachit/gorm-seeder)
- [godfish](https://github.com/rafaelespinoza/godfish) **star:14** Database migration manager, works with native query language. Support for cassandra, mysql, postgres, sqlite3.   [![There was an update last month][G]](https://github.com/rafaelespinoza/godfish)   [![godoc][D]](https://godoc.org/github.com/rafaelespinoza/godfish)
- [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) **star:11** CLI-friendly package for go-pg migrations management.   [![godoc][D]](https://godoc.org/github.com/lawzava/go-pg-migrate)

### Database Tools

- [vitess](https://github.com/youtube/vitess) **star:21141** vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.   [![There was an update last month][G]](https://github.com/youtube/vitess)   [![godoc][D]](https://godoc.org/github.com/youtube/vitess)
- [pgweb](https://github.com/sosedoff/pgweb) **star:9433** Web-based PostgreSQL database browser.   [![There was an update last month][G]](https://github.com/sosedoff/pgweb)   [![godoc][D]](https://godoc.org/github.com/sosedoff/pgweb)
- [go-mysql](https://github.com/siddontang/go-mysql) **star:4962** Go toolset to handle MySQL protocol and replication.   [![There was an update last month][G]](https://github.com/siddontang/go-mysql)   [![godoc][D]](https://godoc.org/github.com/siddontang/go-mysql)
- [pREST](https://github.com/prest/prest) **star:4589** Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new.   [![There was an update last month][G]](https://github.com/prest/prest)   [![godoc][D]](https://godoc.org/github.com/prest/prest)
- [chproxy](https://github.com/Vertamedia/chproxy) **star:1466** HTTP proxy for ClickHouse database.   [![godoc][D]](https://godoc.org/github.com/Vertamedia/chproxy)
- [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) **star:1378** Advanced scheduling for PostgreSQL.   [![There was an update last month][G]](https://github.com/cybertec-postgresql/pg_timetable)   [![godoc][D]](https://godoc.org/github.com/cybertec-postgresql/pg_timetable)
- [onedump](https://github.com/liweiyi88/onedump) **star:978** Database backup from different drivers to different destinations with one command and configuration.   [![There was an update last month][G]](https://github.com/liweiyi88/onedump)   [![godoc][D]](https://godoc.org/github.com/liweiyi88/onedump)
- [rdb](https://github.com/HDT3213/rdb) **star:618** Redis RDB file parser for secondary development and memory analysis.   [![godoc][D]](https://godoc.org/github.com/HDT3213/rdb)
- [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) **star:510** Collects small inserts and sends big requests to ClickHouse servers.   [![There was an update last month][G]](https://github.com/nikepan/clickhouse-bulk)   [![godoc][D]](https://godoc.org/github.com/nikepan/clickhouse-bulk)
- [wescale](https://github.com/wesql/wescale) **star:316** WeScale is a database proxy designed to enhance the scalability, performance, security, and resilience of your applications.   [![godoc][D]](https://godoc.org/github.com/wesql/wescale)
- [gatewayd](https://github.com/gatewayd-io/gatewayd) **star:286** Cloud-native database gateway and framework for building data-driven applications. Like API gateways, for databases.   [![There was an update last month][G]](https://github.com/gatewayd-io/gatewayd)   [![godoc][D]](https://godoc.org/github.com/gatewayd-io/gatewayd)
- [octillery](https://github.com/knocknote/octillery) **star:202** Go package for sharding databases ( Supports every ORM or raw SQL ).   [![godoc][D]](https://godoc.org/github.com/knocknote/octillery)
- [pgrwl](https://github.com/pgrwl/pgrwl) **star:180** Cloud-native continuous backup for PostgreSQL.   [![There was an update last month][G]](https://github.com/pgrwl/pgrwl)   [![godoc][D]](https://godoc.org/github.com/pgrwl/pgrwl)
- [dbbench](https://github.com/sj14/dbbench) **star:118** Database benchmarking tool with support for several databases and scripts.   [![godoc][D]](https://godoc.org/github.com/sj14/dbbench)
- [GoSQLX](https://github.com/ajitpratap0/GoSQLX) **star:105** High-performance SQL parser, formatter, linter, and security scanner with multi-dialect support and WASM playground.   [![godoc][D]](https://godoc.org/github.com/ajitpratap0/GoSQLX)
- [hasql](https://golang.yandex/hasql)  Library for accessing multi-host SQL database installations.
- [gorm-multitenancy](https://github.com/bartventer/gorm-multitenancy) **star:83** Multi-tenancy support for GORM managed databases.   [![godoc][D]](https://godoc.org/github.com/bartventer/gorm-multitenancy)
- [pgxcli](https://github.com/Balaji01-4D/pgxcli) **star:66** PostgreSQL CLI client written in Go, inspired by pgcli.   [![There was an update last month][G]](https://github.com/Balaji01-4D/pgxcli)   [![godoc][D]](https://godoc.org/github.com/Balaji01-4D/pgxcli)
- [dg](https://github.com/codingconcepts/dg) **star:45** A fast data generator that produces CSV files from generated relational data.   [![godoc][D]](https://godoc.org/github.com/codingconcepts/dg)
- [database-gateway](https://github.com/kazhuravlev/database-gateway) **star:39** Running SQL in production with ACLs, logs, and shared links.   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/database-gateway)
- [prep](https://github.com/hexdigest/prep) **star:36** Use prepared SQL statements without changing your code.   [![It hasn't been updated in recent three years][Y]](https://github.com/hexdigest/prep)   [![godoc][D]](https://godoc.org/github.com/hexdigest/prep)
- [rwdb](https://github.com/andizzle/rwdb) **star:21** rwdb provides read replica capability for multiple database servers setup.   [![It hasn't been updated in recent three years][Y]](https://github.com/andizzle/rwdb)   [![godoc][D]](https://godoc.org/github.com/andizzle/rwdb)
- [go-postgres-s3-backup](https://github.com/nicobistolfi/go-postgres-s3-backup) **star:7** Serverless PostgreSQL backups to S3 using AWS Lambda, with daily, monthly, and yearly rotation.   [![godoc][D]](https://godoc.org/github.com/nicobistolfi/go-postgres-s3-backup)

### SQL Query Builders

_Libraries for building and using SQL._

- [sqlc](https://github.com/kyleconroy/sqlc) **star:18005** Generate type-safe code from SQL.   [![There was an update last month][G]](https://github.com/kyleconroy/sqlc)   [![godoc][D]](https://godoc.org/github.com/kyleconroy/sqlc)
- [Squirrel](https://github.com/Masterminds/squirrel) **star:7965** Go library that helps you build SQL queries.   [![godoc][D]](https://godoc.org/github.com/Masterminds/squirrel)
- [xo](https://github.com/knq/xo) **star:3895** Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.   [![godoc][D]](https://godoc.org/github.com/knq/xo)
- [jet](https://github.com/go-jet/jet) **star:3744** Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.   [![godoc][D]](https://godoc.org/github.com/go-jet/jet)
- [goqu](https://github.com/doug-martin/goqu) **star:2668** Idiomatic SQL builder and query library.   [![godoc][D]](https://godoc.org/github.com/doug-martin/goqu)
- [gendry](https://github.com/didi/gendry) **star:1638** Non-invasive SQL builder and powerful data binder.   [![godoc][D]](https://godoc.org/github.com/didi/gendry)   [![Contains Chinese documents][CN]](https://github.com/didi/gendry)   [![Archived][Archived]](https://github.com/didi/gendry)
- [Dotsql](https://github.com/gchaincl/dotsql) **star:743** Go library that helps you keep sql files in one place and use them with ease.   [![godoc][D]](https://godoc.org/github.com/gchaincl/dotsql)
- [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) **star:669** Powerful data retrieval methods as well as DB-agnostic query building capabilities.   [![There was an update last month][G]](https://github.com/go-ozzo/ozzo-dbx)   [![godoc][D]](https://godoc.org/github.com/go-ozzo/ozzo-dbx)
- [sqlingo](https://github.com/lqs/sqlingo) **star:456** A lightweight DSL to build SQL in Go.   [![godoc][D]](https://godoc.org/github.com/lqs/sqlingo)
- [dbq](https://github.com/rocketlaunchr/dbq) **star:416** Zero boilerplate database operations for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/rocketlaunchr/dbq)   [![godoc][D]](https://godoc.org/github.com/rocketlaunchr/dbq)
- [sqrl](https://github.com/elgris/sqrl) **star:285** SQL query builder, fork of Squirrel with improved performance.   [![It hasn't been updated in recent three years][Y]](https://github.com/elgris/sqrl)   [![godoc][D]](https://godoc.org/github.com/elgris/sqrl)
- [Squalus](https://gitlab.com/qosenergy/squalus)  Thin layer over the Go SQL package that makes it easier to perform queries.
- [sq](https://github.com/bokwoon95/go-structured-query) **star:201** Type-safe SQL builder and struct mapper for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/bokwoon95/go-structured-query)   [![godoc][D]](https://godoc.org/github.com/bokwoon95/go-structured-query)
- [bqb](https://github.com/nullism/bqb) **star:194** Lightweight and easy to learn query builder.   [![godoc][D]](https://godoc.org/github.com/nullism/bqb)
- [sqlf](https://github.com/leporo/sqlf) **star:189** Fast SQL query builder.   [![godoc][D]](https://godoc.org/github.com/leporo/sqlf)
- [buildsqlx](https://github.com/arthurkushman/buildsqlx) **star:187** Go database query builder library for PostgreSQL.   [![godoc][D]](https://godoc.org/github.com/arthurkushman/buildsqlx)
- [igor](https://github.com/galeone/igor) **star:127** Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.   [![godoc][D]](https://godoc.org/github.com/galeone/igor)
- [builq](https://github.com/cristalhq/builq) **star:98** Easily build SQL queries in Go.   [![godoc][D]](https://godoc.org/github.com/cristalhq/builq)
- [godbal](https://github.com/xujiajun/godbal) **star:60** Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.   [![It hasn't been updated in recent three years][Y]](https://github.com/xujiajun/godbal)   [![godoc][D]](https://godoc.org/github.com/xujiajun/godbal)
- [qrafter](https://github.com/SennovE/qrafter) **star:53** Type-safe SQL query builder with dialect-aware rendering, schema introspection, and migration generation.   [![godoc][D]](https://godoc.org/github.com/SennovE/qrafter)
- [gosql](https://github.com/twharmon/gosql) **star:37** SQL Query builder with better null values support.   [![It hasn't been updated in recent three years][Y]](https://github.com/twharmon/gosql)   [![godoc][D]](https://godoc.org/github.com/twharmon/gosql)
- [qry](https://github.com/HnH/qry) **star:35** Tool that generates constants from files with raw SQL queries.   [![godoc][D]](https://godoc.org/github.com/HnH/qry)
- [Hotcoal](https://github.com/motrboat/hotcoal) **star:23** Secure your handcrafted SQL against injection.   [![godoc][D]](https://godoc.org/github.com/motrboat/hotcoal)
- [relica](https://github.com/coregx/relica) **star:17** Type-safe database query builder with zero production dependencies, LRU statement cache, batch operations, and support for JOINs, subqueries, CTEs, and window functions.   [![godoc][D]](https://godoc.org/github.com/coregx/relica)
- [ormlite](https://github.com/pupizoid/ormlite) **star:15** Lightweight package containing some ORM-like features and helpers for sqlite databases.   [![It hasn't been updated in recent three years][Y]](https://github.com/pupizoid/ormlite)   [![godoc][D]](https://godoc.org/github.com/pupizoid/ormlite)
- [obreron](https://github.com/profe-ajedrez/obreron) **star:14** Fast and cheap SQL builder which does only one thing, SQL building.   [![godoc][D]](https://godoc.org/github.com/profe-ajedrez/obreron)
- [sqlh](https://github.com/kirill-scherba/sqlh) **star:14** Zero-boilerplate SQL helper with struct tags and Go generics (CRUD, UPSERT, JOIN, benchmarks).   [![godoc][D]](https://godoc.org/github.com/kirill-scherba/sqlh)
- [patcher](https://github.com/Jacobbrewer1/patcher) **star:9** Powerful SQL Query builder that automatically generates SQL queries from structs.   [![godoc][D]](https://godoc.org/github.com/Jacobbrewer1/patcher)
- [sg](https://github.com/go-the-way/sg) **star:7** A SQL Gen for generating standard SQLs(supports: CRUD) written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-the-way/sg)   [![godoc][D]](https://godoc.org/github.com/go-the-way/sg)

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley](https://github.com/google/cayley) **star:15050** Graph database with support for multiple backends.   [![godoc][D]](https://godoc.org/github.com/google/cayley)
- [gokv](https://github.com/philippgille/gokv) **star:829** Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).   [![godoc][D]](https://godoc.org/github.com/philippgille/gokv)
- [go-transaction-manager](https://github.com/avito-tech/go-transaction-manager) **star:412** Transaction manager with multiple adapters (sql, sqlx, gorm, mongo, ...) controls transaction boundaries.   [![There was an update last month][G]](https://github.com/avito-tech/go-transaction-manager)   [![godoc][D]](https://godoc.org/github.com/avito-tech/go-transaction-manager)
- [dsc](https://github.com/viant/dsc) **star:37** Datastore connectivity for SQL, NoSQL, structured files.   [![godoc][D]](https://godoc.org/github.com/viant/dsc)
- [dynamo](https://github.com/fogfish/dynamo) **star:24** A simple key-value abstraction to store algebraic and linked-data data types at AWS storage services: AWS DynamoDB and AWS S3.   [![godoc][D]](https://godoc.org/github.com/fogfish/dynamo)
- [transactor](https://github.com/metalfm/transactor) **star:6** Type-safe transaction boundary abstraction with adapters for database/sql, sqlx, and pgx.   [![godoc][D]](https://godoc.org/github.com/metalfm/transactor)

### Relational Database Drivers

- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) **star:15272** MySQL driver for Go.   [![There was an update last month][G]](https://github.com/go-sql-driver/mysql)   [![godoc][D]](https://godoc.org/github.com/go-sql-driver/mysql)
- [pgx](https://github.com/jackc/pgx) **star:14034** PostgreSQL driver supporting features beyond those exposed by database/sql.   [![There was an update last month][G]](https://github.com/jackc/pgx)   [![godoc][D]](https://godoc.org/github.com/jackc/pgx)
- [pq](https://github.com/lib/pq) **star:9909** Pure Go Postgres driver for database/sql.   [![There was an update last month][G]](https://github.com/lib/pq)   [![godoc][D]](https://godoc.org/github.com/lib/pq)
- [go-sqlite3](https://github.com/mattn/go-sqlite3) **star:9176** SQLite3 driver for go that uses database/sql.   [![There was an update last month][G]](https://github.com/mattn/go-sqlite3)
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) **star:1883** Microsoft MSSQL driver for Go.   [![godoc][D]](https://godoc.org/github.com/denisenkom/go-mssqldb)
- [go-sqlite3](https://github.com/ncruces/go-sqlite3) **star:1060** This Go module is compatible with the database/sql driver. It allows embedding SQLite into your application, provides direct access to its C API, supports SQLite VFS, and also includes a GORM driver.   [![There was an update last month][G]](https://github.com/ncruces/go-sqlite3)   [![godoc][D]](https://godoc.org/github.com/ncruces/go-sqlite3)
- [sqlhooks](https://github.com/qustavo/sqlhooks) **star:667** Attach hooks to any database/sql driver.   [![godoc][D]](https://godoc.org/github.com/qustavo/sqlhooks)
- [sqlite](https://pkg.go.dev/modernc.org/sqlite)  Package sqlite is a sql/database driver using a CGo-free port of the C SQLite3 library.
- [go-oci8](https://github.com/mattn/go-oci8) **star:631** Oracle driver for go that uses database/sql.   [![There was an update last month][G]](https://github.com/mattn/go-oci8)   [![godoc][D]](https://godoc.org/github.com/mattn/go-oci8)
- [godror](https://github.com/godror/godror) **star:597** Oracle driver for Go, using the ODPI-C driver.   [![There was an update last month][G]](https://github.com/godror/godror)
- [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) **star:532** SQLite with pure Go.   [![godoc][D]](https://godoc.org/github.com/cvilsmeier/sqinn-go)
- [KSQL](https://github.com/VinGarcia/ksql) **star:357** A Simple and Powerful Golang SQL Library.   [![godoc][D]](https://godoc.org/github.com/VinGarcia/ksql)
- [surrealdb.go](https://github.com/surrealdb/surrealdb.go) **star:314** SurrealDB Driver for Go.   [![There was an update last month][G]](https://github.com/surrealdb/surrealdb.go)   [![godoc][D]](https://godoc.org/github.com/surrealdb/surrealdb.go)
- [firebirdsql](https://github.com/nakagami/firebirdsql) **star:265** Firebird RDBMS SQL driver for Go.   [![godoc][D]](https://godoc.org/github.com/nakagami/firebirdsql)
- [go-rqlite](https://github.com/rqlite/gorqlite) **star:188** A Go client for rqlite, providing easy-to-use abstractions for working with the rqlite API.   [![godoc][D]](https://godoc.org/github.com/rqlite/gorqlite)
- [ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk) **star:181** native and database/sql driver YDB (Yandex Database).   [![There was an update last month][G]](https://github.com/ydb-platform/ydb-go-sdk)   [![godoc][D]](https://godoc.org/github.com/ydb-platform/ydb-go-sdk)
- [go-adodb](https://github.com/mattn/go-adodb) **star:153** Microsoft ActiveX Object DataBase driver for go that uses database/sql.   [![There was an update last month][G]](https://github.com/mattn/go-adodb)   [![godoc][D]](https://godoc.org/github.com/mattn/go-adodb)
- [avatica](https://github.com/apache/calcite-avatica-go) **star:126** Apache Avatica/Phoenix SQL driver for database/sql.   [![godoc][D]](https://godoc.org/github.com/apache/calcite-avatica-go)
- [gofreetds](https://github.com/minus5/gofreetds) **star:114** Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org).   [![It hasn't been updated in recent three years][Y]](https://github.com/minus5/gofreetds)   [![godoc][D]](https://godoc.org/github.com/minus5/gofreetds)
- [bgc](https://github.com/viant/bgc) **star:21** Datastore Connectivity for BigQuery for go.   [![godoc][D]](https://godoc.org/github.com/viant/bgc)
- [pig](https://github.com/alexeyco/pig) **star:18** Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily.   [![godoc][D]](https://godoc.org/github.com/alexeyco/pig)

### NoSQL Database Drivers

- [redis](https://github.com/redis/go-redis) **star:22193** Redis client for Golang.   [![There was an update last month][G]](https://github.com/redis/go-redis)   [![godoc][D]](https://godoc.org/github.com/redis/go-redis)   [![Contains Chinese documents][CN]](https://github.com/redis/go-redis)
- [rueidis](http://github.com/rueian/rueidis)  Fast Redis RESP3 client with auto pipelining and server-assisted client side caching.
- [redigo](https://github.com/gomodule/redigo) **star:9857** Redigo is a Go client for the Redis database.   [![godoc][D]](https://godoc.org/github.com/gomodule/redigo)
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) **star:8535** Official MongoDB driver for the Go language.   [![There was an update last month][G]](https://github.com/mongodb/mongo-go-driver)   [![godoc][D]](https://godoc.org/github.com/mongodb/mongo-go-driver)
- [mgo](https://github.com/globalsign/mgo) **star:1965** (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.   [![It hasn't been updated in recent three years][Y]](https://github.com/globalsign/mgo)   [![godoc][D]](https://godoc.org/github.com/globalsign/mgo)
- [gorethink](https://github.com/dancannon/gorethink) **star:1649** Go language driver for RethinkDB.   [![godoc][D]](https://godoc.org/github.com/dancannon/gorethink)
- [qmgo](https://github.com/qiniu/qmgo) **star:1355** The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo.   [![godoc][D]](https://godoc.org/github.com/qiniu/qmgo)   [![Contains Chinese documents][CN]](https://github.com/qiniu/qmgo)
- [mgm](https://github.com/kamva/mgm) **star:763** MongoDB model-based ODM for Go (based on official MongoDB driver).   [![godoc][D]](https://godoc.org/github.com/kamva/mgm)
- [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) **star:458** Aerospike client in Go language.   [![There was an update last month][G]](https://github.com/aerospike/aerospike-client-go)   [![godoc][D]](https://godoc.org/github.com/aerospike/aerospike-client-go)
- [redeo](https://github.com/bsm/redeo) **star:445** Redis-protocol compatible TCP servers/services.   [![It hasn't been updated in recent three years][Y]](https://github.com/bsm/redeo)   [![godoc][D]](https://godoc.org/github.com/bsm/redeo)
- [neoism](https://github.com/jmcvetta/neoism) **star:388** Neo4j client for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/jmcvetta/neoism)   [![godoc][D]](https://godoc.org/github.com/jmcvetta/neoism)
- [gocb](https://github.com/couchbase/gocb) **star:377** Official Couchbase Go SDK.   [![There was an update last month][G]](https://github.com/couchbase/gocb)   [![godoc][D]](https://godoc.org/github.com/couchbase/gocb)
- [Kivik](https://github.com/go-kivik/kivik) **star:344** Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.   [![godoc][D]](https://godoc.org/github.com/go-kivik/kivik)
- [go-rejson](https://github.com/nitishm/go-rejson) **star:341** Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.   [![godoc][D]](https://godoc.org/github.com/nitishm/go-rejson)
- [go-couchbase](https://github.com/couchbase/go-couchbase) **star:323** Couchbase client in Go.   [![godoc][D]](https://godoc.org/github.com/couchbase/go-couchbase)
- [go-mongox](https://github.com/chenmingyong0423/go-mongox) **star:221** A Go Mongo library based on the official driver, featuring streamlined document operations, generic binding of structs to collections, built-in CRUD, aggregation, automated field updates, struct validation, hooks, and plugin-based programming.   [![godoc][D]](https://godoc.org/github.com/chenmingyong0423/go-mongox)   [![Contains Chinese documents][CN]](https://github.com/chenmingyong0423/go-mongox)
- [godis](https://github.com/piaohao/godis) **star:112** redis client implement by golang, inspired by jedis.   [![It hasn't been updated in recent three years][Y]](https://github.com/piaohao/godis)   [![godoc][D]](https://godoc.org/github.com/piaohao/godis)
- [arangolite](https://github.com/solher/arangolite) **star:72** Lightweight golang driver for ArangoDB.   [![It hasn't been updated in recent three years][Y]](https://github.com/solher/arangolite)   [![godoc][D]](https://godoc.org/github.com/solher/arangolite)
- [go-pilosa](https://github.com/pilosa/go-pilosa) **star:58** Go client library for Pilosa.   [![It hasn't been updated in recent three years][Y]](https://github.com/pilosa/go-pilosa)   [![godoc][D]](https://godoc.org/github.com/pilosa/go-pilosa)   [![Archived][Archived]](https://github.com/pilosa/go-pilosa)
- [forestdb](https://github.com/couchbase/goforestdb) **star:36** Go bindings for ForestDB.   [![It hasn't been updated in recent three years][Y]](https://github.com/couchbase/goforestdb)   [![godoc][D]](https://godoc.org/github.com/couchbase/goforestdb)
- [goriak](https://github.com/zegl/goriak) **star:30** Go language driver for Riak KV.   [![It hasn't been updated in recent three years][Y]](https://github.com/zegl/goriak)   [![godoc][D]](https://godoc.org/github.com/zegl/goriak)
- [neo4j](https://github.com/cihangir/neo4j) **star:29** Neo4j Rest API Bindings for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/cihangir/neo4j)   [![godoc][D]](https://godoc.org/github.com/cihangir/neo4j)
- [gomemcached](https://github.com/aliexpressru/gomemcached) **star:23** A binary Memcached client for Go with support for sharding using consistent hashing, along with SASL.   [![godoc][D]](https://godoc.org/github.com/aliexpressru/gomemcached)
- [gocosmos](https://github.com/btnguyen2k/gocosmos) **star:22** REST client and standard `database/sql` driver for Azure Cosmos DB.   [![godoc][D]](https://godoc.org/github.com/btnguyen2k/gocosmos)
- [gocql](https://gocql.github.io)  Go language driver for Apache Cassandra.
- [xredis](https://github.com/shomali11/xredis) **star:19** Typesafe, customizable, clean & easy to use Redis client.   [![It hasn't been updated in recent three years][Y]](https://github.com/shomali11/xredis)   [![godoc][D]](https://godoc.org/github.com/shomali11/xredis)
- [asc](https://github.com/viant/asc) **star:11** Datastore Connectivity for Aerospike for go.   [![godoc][D]](https://godoc.org/github.com/viant/asc)
- [godscache](https://github.com/defcronyke/godscache) **star:11** A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.   [![It hasn't been updated in recent three years][Y]](https://github.com/defcronyke/godscache)   [![godoc][D]](https://godoc.org/github.com/defcronyke/godscache)
- [gomemcache](https://github.com/bradfitz/gomemcache/)  memcache client library for the Go programming language.

### Search and Analytic Databases

- [elastic](https://github.com/olivere/elastic) **star:7450** Elasticsearch client for Go.   [![godoc][D]](https://godoc.org/github.com/olivere/elastic)
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch) **star:6058** Official Elasticsearch client for Go.   [![There was an update last month][G]](https://github.com/elastic/go-elasticsearch)   [![godoc][D]](https://godoc.org/github.com/elastic/go-elasticsearch)
- [zoekt](https://github.com/sourcegraph/zoekt) **star:1768** Fast trigram based code search.   [![There was an update last month][G]](https://github.com/sourcegraph/zoekt)   [![godoc][D]](https://godoc.org/github.com/sourcegraph/zoekt)
- [elasticsql](https://github.com/cch123/elasticsql) **star:1195** Convert sql to elasticsearch dsl in Go.   [![godoc][D]](https://godoc.org/github.com/cch123/elasticsql)
- [elastigo](https://github.com/mattbaird/elastigo) **star:941** Elasticsearch client library.   [![It hasn't been updated in recent three years][Y]](https://github.com/mattbaird/elastigo)   [![godoc][D]](https://godoc.org/github.com/mattbaird/elastigo)
- [skizze](https://github.com/skizzehq/skizze) **star:773** A probabilistic data structure service and storage.   [![It hasn't been updated in recent three years][Y]](https://github.com/skizzehq/skizze)   [![godoc][D]](https://godoc.org/github.com/skizzehq/skizze)
- [effdsl](https://github.com/sdqri/effdsl) **star:35** Elasticsearch query builder for Go.   [![godoc][D]](https://godoc.org/github.com/sdqri/effdsl)
- [goes](https://github.com/OwnLocal/goes) **star:30** Library to interact with Elasticsearch.   [![It hasn't been updated in recent three years][Y]](https://github.com/OwnLocal/goes)   [![godoc][D]](https://godoc.org/github.com/OwnLocal/goes)
- [clickhouse-go](https://github.com/ClickHouse/clickhouse-go/)  ClickHouse SQL client for Go with a `database/sql` compatibility.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [carbon](https://github.com/dromara/carbon) **star:5223** A simple, semantic and developer-friendly time package for golang.   [![godoc][D]](https://godoc.org/github.com/dromara/carbon)   [![Contains Chinese documents][CN]](https://github.com/dromara/carbon)
- [now](https://github.com/jinzhu/now) **star:4696** Now is a time toolkit for golang.   [![godoc][D]](https://godoc.org/github.com/jinzhu/now)
- [dateparse](https://github.com/araddon/dateparse) **star:2147** Parse date's without knowing format in advance.   [![godoc][D]](https://godoc.org/github.com/araddon/dateparse)
- [carbon](https://github.com/uniplaces/carbon) **star:781** Simple Time extension with a lot of util methods, ported from PHP Carbon library.   [![godoc][D]](https://godoc.org/github.com/uniplaces/carbon)
- [durafmt](https://github.com/hako/durafmt) **star:513** Time duration formatting library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/hako/durafmt)   [![godoc][D]](https://godoc.org/github.com/hako/durafmt)
- [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) **star:240** The implementation of the Persian (Solar Hijri) Calendar in Go (golang).   [![godoc][D]](https://godoc.org/github.com/yaa110/go-persian-calendar)
- [gostradamus](https://github.com/bykof/gostradamus) **star:208** A Go package for working with dates.   [![godoc][D]](https://godoc.org/github.com/bykof/gostradamus)
- [timeutil](https://github.com/leekchan/timeutil) **star:193** Useful extensions (Timedelta, Strftime, ...) to the golang's time package.   [![It hasn't been updated in recent three years][Y]](https://github.com/leekchan/timeutil)   [![godoc][D]](https://godoc.org/github.com/leekchan/timeutil)
- [go-sunrise](https://github.com/nathan-osman/go-sunrise) **star:177** Calculate the sunrise and sunset times for a given location.   [![godoc][D]](https://godoc.org/github.com/nathan-osman/go-sunrise)
- [iso8601](https://github.com/relvacode/iso8601) **star:158** Efficiently parse ISO8601 date-times without regex.   [![godoc][D]](https://godoc.org/github.com/relvacode/iso8601)
- [date](https://github.com/rickb777/date) **star:143** Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.   [![godoc][D]](https://godoc.org/github.com/rickb777/date)
- [go-str2duration](https://github.com/xhit/go-str2duration) **star:118** Convert string to duration. Support time.Duration returned string and more.   [![godoc][D]](https://godoc.org/github.com/xhit/go-str2duration)
- [timespan](https://github.com/SaidinWoT/timespan) **star:84** For interacting with intervals of time, defined as a start time and a duration.   [![It hasn't been updated in recent three years][Y]](https://github.com/SaidinWoT/timespan)   [![godoc][D]](https://godoc.org/github.com/SaidinWoT/timespan)
- [feiertage](https://github.com/wlbr/feiertage) **star:53** Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...   [![godoc][D]](https://godoc.org/github.com/wlbr/feiertage)
- [go-anytime](https://github.com/ijt/go-anytime) **star:36** Parse dates/times like "next dec 22nd at 3pm" and ranges like "from today until next thursday" without knowing the format in advance.   [![It hasn't been updated in recent three years][Y]](https://github.com/ijt/go-anytime)   [![godoc][D]](https://godoc.org/github.com/ijt/go-anytime)
- [kair](https://github.com/GuilhermeCaruso/kair) **star:24** Date and Time - Golang Formatting Library.   [![It hasn't been updated in recent three years][Y]](https://github.com/GuilhermeCaruso/kair)   [![godoc][D]](https://godoc.org/github.com/GuilhermeCaruso/kair)
- [approx](https://github.com/goschtalt/approx) **star:18** A Duration extension supporting parsing/printing durations in days, weeks and years.   [![godoc][D]](https://godoc.org/github.com/goschtalt/approx)
- [cronrange](https://github.com/1set/cronrange) **star:18** Parses Cron-style time range expressions, checks if the given time is within any ranges.   [![It hasn't been updated in recent three years][Y]](https://github.com/1set/cronrange)   [![godoc][D]](https://godoc.org/github.com/1set/cronrange)
- [strftime](https://github.com/awoodbeck/strftime) **star:14** C99-compatible strftime formatter.   [![It hasn't been updated in recent three years][Y]](https://github.com/awoodbeck/strftime)   [![godoc][D]](https://godoc.org/github.com/awoodbeck/strftime)
- [tuesday](https://github.com/osteele/tuesday) **star:13** Ruby-compatible Strftime function.   [![godoc][D]](https://godoc.org/github.com/osteele/tuesday)
- [go-week](https://github.com/stoewer/go-week) **star:11** An efficient package to work with ISO8601 week dates.   [![It hasn't been updated in recent three years][Y]](https://github.com/stoewer/go-week)   [![godoc][D]](https://godoc.org/github.com/stoewer/go-week)
- [go-date-fns](https://github.com/chmenegatti/go-date-fns) **star:6** A comprehensive date utility library for Go, inspired by date-fns, with 140+ pure and immutable functions.   [![godoc][D]](https://godoc.org/github.com/chmenegatti/go-date-fns)
- [go-datebin](https://github.com/deatil/go-datebin) **star:5** A simple datetime parse pkg.   [![godoc][D]](https://godoc.org/github.com/deatil/go-datebin)   [![Contains Chinese documents][CN]](https://github.com/deatil/go-datebin)
- [go-faketime](https://github.com/harkaitz/go-faketime) **star:5** A simple `time.Now()` that honors the faketime(1) utility.   [![godoc][D]](https://godoc.org/github.com/harkaitz/go-faketime)

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [go-zero](https://github.com/tal-tech/go-zero) **star:33185** A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity.   [![There was an update last month][G]](https://github.com/tal-tech/go-zero)   [![godoc][D]](https://godoc.org/github.com/tal-tech/go-zero)
- [go-kit](https://github.com/go-kit/kit) **star:27437** Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.   [![godoc][D]](https://godoc.org/github.com/go-kit/kit)
- [Kratos](https://github.com/go-kratos/kratos) **star:25792** A modular-designed and easy-to-use microservices framework in Go.   [![godoc][D]](https://godoc.org/github.com/go-kratos/kratos)   [![Contains Chinese documents][CN]](https://github.com/go-kratos/kratos)
- [grpc-go](https://github.com/grpc/grpc-go) **star:23006** The Go language implementation of gRPC. HTTP/2 based RPC.   [![There was an update last month][G]](https://github.com/grpc/grpc-go)   [![godoc][D]](https://godoc.org/github.com/grpc/grpc-go)
- [go-micro](https://github.com/micro/go-micro) **star:22964** A distributed systems development framework.   [![There was an update last month][G]](https://github.com/micro/go-micro)   [![godoc][D]](https://godoc.org/github.com/micro/go-micro)
- [NATS](https://github.com/nats-io/nats-server) **star:20212** NATS is a simple, secure, and performant communications system for digital systems, services, and devices.   [![There was an update last month][G]](https://github.com/nats-io/nats-server)   [![godoc][D]](https://godoc.org/github.com/nats-io/nats-server)
- [raft](https://github.com/hashicorp/raft) **star:9055** Golang implementation of the Raft consensus protocol, by HashiCorp.   [![godoc][D]](https://godoc.org/github.com/hashicorp/raft)
- [rpcx](https://github.com/smallnest/rpcx) **star:8306** Distributed pluggable RPC service framework like alibaba Dubbo.   [![There was an update last month][G]](https://github.com/smallnest/rpcx)   [![godoc][D]](https://godoc.org/github.com/smallnest/rpcx)
- [Kitex](https://github.com/cloudwego/kitex) **star:7995** A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice.   [![There was an update last month][G]](https://github.com/cloudwego/kitex)   [![godoc][D]](https://godoc.org/github.com/cloudwego/kitex)   [![Contains Chinese documents][CN]](https://github.com/cloudwego/kitex)
- [lura](https://github.com/luraproject/lura) **star:6790** Ultra performant API Gateway framework with middlewares.   [![There was an update last month][G]](https://github.com/luraproject/lura)   [![godoc][D]](https://godoc.org/github.com/luraproject/lura)
- [torrent](https://github.com/anacrolix/torrent) **star:6068** BitTorrent client package.   [![godoc][D]](https://godoc.org/github.com/anacrolix/torrent)
- [dragonboat](https://github.com/lni/dragonboat) **star:5316** A feature complete and high performance multi-group Raft library in Go.   [![godoc][D]](https://godoc.org/github.com/lni/dragonboat)   [![Contains Chinese documents][CN]](https://github.com/lni/dragonboat)
- [evans](https://github.com/ktr0731/evans) **star:4489** Evans: more expressive universal gRPC client.   [![godoc][D]](https://godoc.org/github.com/ktr0731/evans)
- [emitter-io](https://github.com/emitter-io/emitter) **star:4006** High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.   [![godoc][D]](https://godoc.org/github.com/emitter-io/emitter)
- [gleam](https://github.com/chrislusf/gleam) **star:3563** Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.   [![There was an update last month][G]](https://github.com/chrislusf/gleam)   [![godoc][D]](https://godoc.org/github.com/chrislusf/gleam)
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) **star:3259** Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures.   [![There was an update last month][G]](https://github.com/dragonflyoss/Dragonfly2)   [![godoc][D]](https://godoc.org/github.com/dragonflyoss/Dragonfly2)
- [glow](https://github.com/chrislusf/glow) **star:3219** Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/chrislusf/glow)   [![godoc][D]](https://godoc.org/github.com/chrislusf/glow)
- [sponge](https://github.com/zhufuyi/sponge) **star:2854** A distributed development framework that integrates automatic code generation, gin and grpc frameworks, base development frameworks.   [![godoc][D]](https://godoc.org/github.com/zhufuyi/sponge)   [![Contains Chinese documents][CN]](https://github.com/zhufuyi/sponge)
- [liftbridge](https://github.com/liftbridge-io/liftbridge) **star:2802** Lightweight, fault-tolerant message streams for NATS.   [![There was an update last month][G]](https://github.com/liftbridge-io/liftbridge)   [![godoc][D]](https://godoc.org/github.com/liftbridge-io/liftbridge)
- [go-eagle](https://github.com/go-eagle/eagle) **star:2425** A Go framework for the API or Microservice with handy scaffolding tools.   [![godoc][D]](https://godoc.org/github.com/go-eagle/eagle)   [![Contains Chinese documents][CN]](https://github.com/go-eagle/eagle)
- [oras](https://github.com/oras-project/oras) **star:2349** CLI and library for OCI Artifacts in container registries.   [![There was an update last month][G]](https://github.com/oras-project/oras)   [![godoc][D]](https://godoc.org/github.com/oras-project/oras)
- [mochi mqtt](https://github.com/mochi-co/mqtt) **star:1896** Fully spec compliant, embeddable high-performance MQTT v5/v3 broker for IoT, smarthome, and pubsub.   [![godoc][D]](https://godoc.org/github.com/mochi-co/mqtt)   [![Contains Chinese documents][CN]](https://github.com/mochi-co/mqtt)
- [redis-lock](https://github.com/bsm/redislock) **star:1765** Simplified distributed locking implementation using Redis.   [![godoc][D]](https://godoc.org/github.com/bsm/redislock)
- [resgate](https://resgate.io/)  Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
- [hprose](https://github.com/hprose/hprose-golang) **star:1259** Very newbility RPC Library, support 25+ languages now.   [![godoc][D]](https://godoc.org/github.com/hprose/hprose-golang)   [![Contains Chinese documents][CN]](https://github.com/hprose/hprose-golang)
- [K8gb](https://github.com/k8gb-io/k8gb) **star:1190** A cloud native Kubernetes Global Balancer.   [![There was an update last month][G]](https://github.com/k8gb-io/k8gb)   [![godoc][D]](https://godoc.org/github.com/k8gb-io/k8gb)
- [trpc-go](https://github.com/trpc-group/trpc-go) **star:1180** The Go language implementation of tRPC, which is a pluggable, high-performance RPC framework.   [![godoc][D]](https://godoc.org/github.com/trpc-group/trpc-go)   [![Contains Chinese documents][CN]](https://github.com/trpc-group/trpc-go)
- [go-doudou](https://github.com/unionj-cloud/go-doudou) **star:1172** A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity.   [![godoc][D]](https://godoc.org/github.com/unionj-cloud/go-doudou)   [![Contains Chinese documents][CN]](https://github.com/unionj-cloud/go-doudou)
- [rain](https://github.com/cenkalti/rain) **star:1141** BitTorrent client and library.   [![There was an update last month][G]](https://github.com/cenkalti/rain)   [![godoc][D]](https://godoc.org/github.com/cenkalti/rain)
- [arpc](https://github.com/lesismal/arpc) **star:1090** More effective network communication, support two-way-calling, notify, broadcast.   [![godoc][D]](https://godoc.org/github.com/lesismal/arpc)
- [raft](https://github.com/etcd-io/raft) **star:1072** Go implementation of the Raft consensus protocol, by CoreOS.   [![There was an update last month][G]](https://github.com/etcd-io/raft)   [![godoc][D]](https://godoc.org/github.com/etcd-io/raft)
- [Temporal](https://github.com/temporalio/sdk-go) **star:928** Durable execution system for making code fault-tolerant and simple.   [![There was an update last month][G]](https://github.com/temporalio/sdk-go)   [![godoc][D]](https://godoc.org/github.com/temporalio/sdk-go)
- [opentelemetry-go-auto-instrumentation](https://github.com/alibaba/opentelemetry-go-auto-instrumentation) **star:881** OpenTelemetry Compile-Time Instrumentation for Golang.   [![There was an update last month][G]](https://github.com/alibaba/opentelemetry-go-auto-instrumentation)   [![godoc][D]](https://godoc.org/github.com/alibaba/opentelemetry-go-auto-instrumentation)   [![Contains Chinese documents][CN]](https://github.com/alibaba/opentelemetry-go-auto-instrumentation)
- [consistent](https://github.com/buraksezer/consistent) **star:777** Consistent hashing with bounded loads.   [![godoc][D]](https://godoc.org/github.com/buraksezer/consistent)
- [gorpc](https://github.com/valyala/gorpc) **star:709** Simple, fast and scalable RPC library for high load.   [![It hasn't been updated in recent three years][Y]](https://github.com/valyala/gorpc)   [![godoc][D]](https://godoc.org/github.com/valyala/gorpc)
- [go-sundheit](https://github.com/AppsFlyer/go-sundheit) **star:561** A library built to provide support for defining async service health checks for golang services.   [![There was an update last month][G]](https://github.com/AppsFlyer/go-sundheit)   [![godoc][D]](https://godoc.org/github.com/AppsFlyer/go-sundheit)
- [digota](https://github.com/digota/digota) **star:525** grpc ecommerce microservice.   [![It hasn't been updated in recent three years][Y]](https://github.com/digota/digota)   [![godoc][D]](https://godoc.org/github.com/digota/digota)
- [dot](https://github.com/dotchain/dot/)  distributed sync using operational transformation/OT.
- [go-jump](https://github.com/dgryski/go-jump) **star:389** Port of Google's "Jump" Consistent Hash function.   [![It hasn't been updated in recent three years][Y]](https://github.com/dgryski/go-jump)   [![godoc][D]](https://godoc.org/github.com/dgryski/go-jump)
- [sleuth](https://github.com/ursiform/sleuth) **star:388** Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).   [![It hasn't been updated in recent three years][Y]](https://github.com/ursiform/sleuth)   [![godoc][D]](https://godoc.org/github.com/ursiform/sleuth)
- [jsonrpc](https://github.com/ybbus/jsonrpc) **star:370** JSON-RPC 2.0 HTTP client implementation.   [![godoc][D]](https://godoc.org/github.com/ybbus/jsonrpc)
- [dht](https://github.com/anacrolix/dht) **star:358** BitTorrent Kademlia DHT implementation.   [![godoc][D]](https://godoc.org/github.com/anacrolix/dht)
- [Tarmac](https://github.com/tarmac-project/tarmac) **star:344** Framework for writing functions, microservices, or monoliths with WebAssembly   [![There was an update last month][G]](https://github.com/tarmac-project/tarmac)   [![godoc][D]](https://godoc.org/github.com/tarmac-project/tarmac)
- [jsonrpc](https://github.com/osamingo/jsonrpc) **star:193** The jsonrpc package helps implement of JSON-RPC 2.0.   [![godoc][D]](https://godoc.org/github.com/osamingo/jsonrpc)
- [outboxer](https://github.com/italolelis/outboxer) **star:168** Outboxer is a go library that implements the outbox pattern.   [![There was an update last month][G]](https://github.com/italolelis/outboxer)   [![godoc][D]](https://godoc.org/github.com/italolelis/outboxer)
- [pglock](https://cirello.io/pglock)  PostgreSQL-backed distributed locking implementation.
- [pjrpc](https://gitlab.com/pjrpc/pjrpc)  Golang JSON-RPC Server-Client with Protobuf spec.
- [outbox](https://github.com/oagudo/outbox) **star:130** Lightweight library for the transactional outbox pattern in Go, not tied to any specific relational database or broker.   [![godoc][D]](https://godoc.org/github.com/oagudo/outbox)
- [doublejump](https://github.com/edwingeng/doublejump) **star:111** A revamped Google's jump consistent hash.   [![It hasn't been updated in recent three years][Y]](https://github.com/edwingeng/doublejump)   [![godoc][D]](https://godoc.org/github.com/edwingeng/doublejump)
- [Semaphore](https://github.com/jexia/semaphore) **star:97** A straightforward (micro) service orchestrator.   [![It hasn't been updated in recent three years][Y]](https://github.com/jexia/semaphore)   [![godoc][D]](https://godoc.org/github.com/jexia/semaphore)
- [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) **star:95** High-performance distributed command pattern library for Go.   [![godoc][D]](https://godoc.org/github.com/cmd-stream/cmd-stream-go)
- [capillaries](https://github.com/capillariesio/capillaries) **star:73** distributed batch data processing framework.   [![There was an update last month][G]](https://github.com/capillariesio/capillaries)   [![godoc][D]](https://godoc.org/github.com/capillariesio/capillaries)
- [flowgraph](https://github.com/vectaport/flowgraph) **star:66** flow-based programming package.   [![There was an update last month][G]](https://github.com/vectaport/flowgraph)   [![godoc][D]](https://godoc.org/github.com/vectaport/flowgraph)
- [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) **star:66** MySQL based distributed lock.   [![godoc][D]](https://godoc.org/github.com/sanketplus/go-mysql-lock)
- [drmaa](https://github.com/dgruber/drmaa) **star:51** Job submission library for cluster schedulers based on the DRMAA standard.   [![godoc][D]](https://godoc.org/github.com/dgruber/drmaa)
- [dynamolock](https://cirello.io/dynamolock)  DynamoDB-backed distributed locking implementation.
- [go-pdu](https://github.com/pdupub/go-pdu) **star:50** A decentralized identity-based social network.   [![godoc][D]](https://godoc.org/github.com/pdupub/go-pdu)
- [committer](https://github.com/vadiminshakov/committer) **star:44** A distributed transactions management system (2PC/3PC implementation).   [![There was an update last month][G]](https://github.com/vadiminshakov/committer)   [![godoc][D]](https://godoc.org/github.com/vadiminshakov/committer)
- [consistenthash](https://github.com/mbrostami/consistenthash) **star:35** Consistent hashing with configurable replicas.   [![godoc][D]](https://godoc.org/github.com/mbrostami/consistenthash)
- [gmsec](https://github.com/gmsec/micro) **star:27** A Go distributed systems development framework.   [![godoc][D]](https://godoc.org/github.com/gmsec/micro)
- [bedrock](https://github.com/z5labs/bedrock) **star:17** Provides a minimal, modular and composable foundation for quickly developing services and more use case specific frameworks in Go.   [![There was an update last month][G]](https://github.com/z5labs/bedrock)   [![godoc][D]](https://godoc.org/github.com/z5labs/bedrock)
- [dynatomic](https://github.com/tylfin/dynatomic) **star:17** A library for using DynamoDB as an atomic counter.   [![It hasn't been updated in recent three years][Y]](https://github.com/tylfin/dynatomic)   [![godoc][D]](https://godoc.org/github.com/tylfin/dynatomic)
- [failured](https://github.com/andy2046/failured) **star:15** adaptive accrual failure detector for distributed systems.   [![It hasn't been updated in recent three years][Y]](https://github.com/andy2046/failured)   [![godoc][D]](https://godoc.org/github.com/andy2046/failured)
- [health](https://github.com/schigh/health) **star:12** Health checker for Go services with Kubernetes probe support.   [![godoc][D]](https://godoc.org/github.com/schigh/health)
- [circuit](https://github.com/schigh/circuit) **star:7** Circuit breaker with gradual recovery via probabilistic throttling.   [![godoc][D]](https://godoc.org/github.com/schigh/circuit)
- [lock](https://github.com/ubgo/lock) **star:5** Distributed lock family with one Go interface and five backends (filelock, flock, Redis, Postgres, etcd) — fencing tokens, semaphore mode, and observability hooks across all backends.   [![godoc][D]](https://godoc.org/github.com/ubgo/lock)
- [rpcplatform](https://github.com/nexcode/rpcplatform) **star:2** Framework for microservices with service discovery, load balancing, and related features.   [![godoc][D]](https://godoc.org/github.com/nexcode/rpcplatform)

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [GoDNS](https://github.com/timothyye/godns) **star:1749** A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.   [![There was an update last month][G]](https://github.com/timothyye/godns)   [![godoc][D]](https://godoc.org/github.com/timothyye/godns)
- [DDNS](https://github.com/skibish/ddns) **star:266** Personal DDNS client with Digital Ocean Networking DNS as backend.   [![godoc][D]](https://godoc.org/github.com/skibish/ddns)
- [dyndns](https://gitlab.com/alcastle/dyndns)  Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [MailHog](https://github.com/mailhog/MailHog) **star:16083** Email and SMTP testing with web and API interface.   [![godoc][D]](https://godoc.org/github.com/mailhog/MailHog)
- [Mailpit](https://github.com/axllent/mailpit) **star:9866** Email and SMTP testing tool for developers.   [![There was an update last month][G]](https://github.com/axllent/mailpit)   [![godoc][D]](https://godoc.org/github.com/axllent/mailpit)
- [Maddy](https://github.com/foxcpp/maddy) **star:6037** All-in-one (SMTP, IMAP, DKIM, DMARC, MTA-STS, DANE) email server   [![There was an update last month][G]](https://github.com/foxcpp/maddy)   [![godoc][D]](https://godoc.org/github.com/foxcpp/maddy)
- [mox](https://github.com/mjl-/mox) **star:5748** Modern full-featured secure mail server for low-maintenance, self-hosted email.   [![godoc][D]](https://godoc.org/github.com/mjl-/mox)
- [hermes](https://github.com/matcornic/hermes) **star:3019** Golang package that generates clean, responsive HTML e-mails.   [![godoc][D]](https://godoc.org/github.com/matcornic/hermes)
- [email](https://github.com/jordan-wright/email) **star:2799** A robust and flexible email library for Go.   [![godoc][D]](https://godoc.org/github.com/jordan-wright/email)
- [go-imap](https://github.com/emersion/go-imap) **star:2342** IMAP library for clients and servers.   [![godoc][D]](https://godoc.org/github.com/emersion/go-imap)
- [chasquid](https://blitiri.com.ar/p/chasquid)  SMTP server written in Go.
- [email-verifier](https://github.com/AfterShip/email-verifier) **star:1589** A Go library for email verification without sending any emails.   [![godoc][D]](https://godoc.org/github.com/AfterShip/email-verifier)
- [go-mail](https://github.com/wneessen/go-mail) **star:1445** A simple Go library for sending mails in Go.   [![There was an update last month][G]](https://github.com/wneessen/go-mail)   [![godoc][D]](https://godoc.org/github.com/wneessen/go-mail)
- [SendGrid](https://github.com/sendgrid/sendgrid-go) **star:1059** SendGrid's Go library for sending email.   [![godoc][D]](https://godoc.org/github.com/sendgrid/sendgrid-go)
- [mailgun-go](https://github.com/mailgun/mailgun-go) **star:745** Go library for sending mail with the Mailgun API.   [![There was an update last month][G]](https://github.com/mailgun/mailgun-go)   [![godoc][D]](https://godoc.org/github.com/mailgun/mailgun-go)
- [go-simple-mail](https://github.com/xhit/go-simple-mail) **star:695** Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.   [![godoc][D]](https://godoc.org/github.com/xhit/go-simple-mail)
- [go-message](https://github.com/emersion/go-message) **star:457** Streaming library for the Internet Message Format and mail messages.   [![godoc][D]](https://godoc.org/github.com/emersion/go-message)
- [douceur](https://github.com/aymerick/douceur) **star:261** CSS inliner for your HTML emails.   [![It hasn't been updated in recent three years][Y]](https://github.com/aymerick/douceur)   [![godoc][D]](https://godoc.org/github.com/aymerick/douceur)
- [Hectane](https://github.com/hectane/hectane) **star:226** Lightweight SMTP client providing an HTTP API.   [![It hasn't been updated in recent three years][Y]](https://github.com/hectane/hectane)   [![godoc][D]](https://godoc.org/github.com/hectane/hectane)
- [go-premailer](https://github.com/vanng822/go-premailer) **star:203** Inline styling for HTML mail in Go.   [![godoc][D]](https://godoc.org/github.com/vanng822/go-premailer)
- [smtpmock](https://github.com/mocktools/go-smtp-mock) **star:166** Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment.   [![godoc][D]](https://godoc.org/github.com/mocktools/go-smtp-mock)
- [mailchain](https://github.com/mailchain/mailchain) **star:142** Send encrypted emails to blockchain addresses written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/mailchain/mailchain)   [![Archived][Archived]](https://github.com/mailchain/mailchain)
- [truemail-go](https://github.com/truemail-rb/truemail-go) **star:135** Configurable Golang email validator/verifier. Verify email via Regex, DNS, SMTP and even more.   [![godoc][D]](https://godoc.org/github.com/truemail-rb/truemail-go)
- [go-imap](https://github.com/BrianLeishman/go-imap) **star:103** Batteries-included IMAP client with auto-reconnect, OAuth2, IDLE support, and built-in MIME parsing.   [![There was an update last month][G]](https://github.com/BrianLeishman/go-imap)   [![godoc][D]](https://godoc.org/github.com/BrianLeishman/go-imap)
- [go-dkim](https://github.com/toorop/go-dkim) **star:99** DKIM library, to sign & verify email.   [![godoc][D]](https://godoc.org/github.com/toorop/go-dkim)
- [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) **star:79** Golang library for providing a canonical representation of email address.   [![godoc][D]](https://godoc.org/github.com/dimuska139/go-email-normalizer)
- [smtp](https://github.com/mailhog/smtp) **star:77** SMTP server protocol state machine.   [![It hasn't been updated in recent three years][Y]](https://github.com/mailhog/smtp)   [![godoc][D]](https://godoc.org/github.com/mailhog/smtp)
- [mailx](https://github.com/valord577/mailx) **star:23** Mailx is a library that makes it easier to send email via SMTP. It is an enhancement of the golang standard library `net/smtp`.   [![godoc][D]](https://godoc.org/github.com/valord577/mailx)
- [tickstem/verify](https://github.com/tickstem/verify) **star:3** Validate email addresses before they hit your database: syntax, MX lookup, disposable domains, and role-based inboxes.   [![godoc][D]](https://godoc.org/github.com/tickstem/verify)

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [FrankenPHP](https://github.com/dunglas/frankenphp) **star:11213** PHP embedded in Go, with a `net/http` handler.   [![There was an update last month][G]](https://github.com/dunglas/frankenphp)   [![godoc][D]](https://godoc.org/github.com/dunglas/frankenphp)
- [expr](https://github.com/antonmedv/expr) **star:7942** Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing.   [![godoc][D]](https://godoc.org/github.com/antonmedv/expr)
- [goja](https://github.com/dop251/goja) **star:6996** ECMAScript 5.1(+) implementation in Go.   [![godoc][D]](https://godoc.org/github.com/dop251/goja)
- [gopher-lua](https://github.com/yuin/gopher-lua) **star:6952** Lua 5.1 VM and compiler written in Go.   [![godoc][D]](https://godoc.org/github.com/yuin/gopher-lua)
- [tengo](https://github.com/d5/tengo) **star:3830** Bytecode compiled script language for Go.   [![godoc][D]](https://godoc.org/github.com/d5/tengo)
- [go-lua](https://github.com/Shopify/go-lua) **star:3446** Port of the Lua 5.2 VM to pure Go.   [![godoc][D]](https://godoc.org/github.com/Shopify/go-lua)
- [cel-go](https://github.com/google/cel-go) **star:3029** Fast, portable, non-Turing complete expression evaluation with gradual typing.   [![There was an update last month][G]](https://github.com/google/cel-go)   [![godoc][D]](https://godoc.org/github.com/google/cel-go)
- [starlark-go](https://github.com/google/starlark-go) **star:2732** Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution.   [![There was an update last month][G]](https://github.com/google/starlark-go)   [![godoc][D]](https://godoc.org/github.com/google/starlark-go)
- [metacall](https://github.com/metacall/core) **star:1807** Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more.   [![There was an update last month][G]](https://github.com/metacall/core)
- [Wa/凹语言](https://github.com/wa-lang/wa) **star:1771** The Wa Programming Language embedded in Go.   [![godoc][D]](https://godoc.org/github.com/wa-lang/wa)   [![Contains Chinese documents][CN]](https://github.com/wa-lang/wa)
- [anko](https://github.com/mattn/anko) **star:1577** Scriptable interpreter written in Go.   [![There was an update last month][G]](https://github.com/mattn/anko)   [![godoc][D]](https://godoc.org/github.com/mattn/anko)
- [go-php](https://github.com/deuill/go-php) **star:944** PHP bindings for Go.   [![godoc][D]](https://godoc.org/github.com/deuill/go-php)
- [goal](https://codeberg.org/anaseto/goal)  An embeddable scripting array language.
- [gval](https://github.com/PaesslerAG/gval) **star:812** A highly customizable expression language written in Go.   [![godoc][D]](https://godoc.org/github.com/PaesslerAG/gval)
- [prolog](https://github.com/ichiban/prolog) **star:727** Embeddable Prolog.   [![godoc][D]](https://godoc.org/github.com/ichiban/prolog)
- [golua](https://github.com/aarzilli/golua) **star:700** Go bindings for Lua C API.
- [gisp](https://github.com/jcla1/gisp) **star:531** Simple LISP in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/jcla1/gisp)   [![godoc][D]](https://godoc.org/github.com/jcla1/gisp)
- [gentee](https://github.com/gentee/gentee) **star:146** Embeddable scripting programming language.   [![godoc][D]](https://godoc.org/github.com/gentee/gentee)
- [binder](https://github.com/alexeyco/binder) **star:80** Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).   [![It hasn't been updated in recent three years][Y]](https://github.com/alexeyco/binder)   [![godoc][D]](https://godoc.org/github.com/alexeyco/binder)
- [starlet](https://github.com/1set/starlet) **star:48** Go wrapper for [starlark-go](https://github.com/google/starlark-go) that simplifies script execution, offers data conversion, and useful Starlark libraries and extensions.   [![There was an update last month][G]](https://github.com/1set/starlet)   [![godoc][D]](https://godoc.org/github.com/1set/starlet)
- [ecal](https://github.com/krotik/ecal) **star:45** A simple embeddable scripting language which supports concurrent event processing.   [![It hasn't been updated in recent three years][Y]](https://github.com/krotik/ecal)   [![godoc][D]](https://godoc.org/github.com/krotik/ecal)
- [purl](https://github.com/ian-kent/purl) **star:41** Perl 5.18.2 embedded in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ian-kent/purl)   [![godoc][D]](https://godoc.org/github.com/ian-kent/purl)
- [ngaro](https://github.com/db47h/ngaro) **star:30** Embeddable Ngaro VM implementation enabling scripting in Retro.   [![It hasn't been updated in recent three years][Y]](https://github.com/db47h/ngaro)   [![godoc][D]](https://godoc.org/github.com/db47h/ngaro)
- [go-lua](https://github.com/speedata/go-lua) **star:8** Lua 5.4 VM implemented in pure Go.   [![godoc][D]](https://godoc.org/github.com/speedata/go-lua)

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [go-multierror](https://github.com/hashicorp/go-multierror) **star:2573** Go (golang) package for representing a list of errors as a single error.   [![godoc][D]](https://godoc.org/github.com/hashicorp/go-multierror)
- [errors](https://github.com/cockroachdb/errors) **star:2434** Go error library with error portability over the network.   [![There was an update last month][G]](https://github.com/cockroachdb/errors)   [![godoc][D]](https://godoc.org/github.com/cockroachdb/errors)
- [eris](https://github.com/rotisserie/eris) **star:1790** A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.   [![godoc][D]](https://godoc.org/github.com/rotisserie/eris)
- [errorx](https://github.com/joomcode/errorx) **star:1271** A feature rich error package with stack traces, composition of errors and more.   [![godoc][D]](https://godoc.org/github.com/joomcode/errorx)
- [multierr](https://github.com/uber-go/multierr) **star:1186** Package for representing a list of errors as a single error.   [![godoc][D]](https://godoc.org/github.com/uber-go/multierr)
- [tracerr](https://github.com/ztrue/tracerr) **star:1120** Golang errors with stack trace and source fragments.   [![godoc][D]](https://godoc.org/github.com/ztrue/tracerr)
- [oops](https://github.com/samber/oops) **star:970** Error handling with context, stack trace and source fragments.   [![There was an update last month][G]](https://github.com/samber/oops)   [![godoc][D]](https://godoc.org/github.com/samber/oops)
- [errlog](https://github.com/snwfdhmp/errlog) **star:460** Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.   [![It hasn't been updated in recent three years][Y]](https://github.com/snwfdhmp/errlog)   [![godoc][D]](https://godoc.org/github.com/snwfdhmp/errlog)
- [emperror](https://github.com/emperror/emperror) **star:385** Error handling tools and best practices for Go libraries and applications.   [![It hasn't been updated in recent three years][Y]](https://github.com/emperror/emperror)   [![godoc][D]](https://godoc.org/github.com/emperror/emperror)
- [Fault](https://github.com/Southclaws/fault) **star:311** An ergonomic mechanism for wrapping errors in order to facilitate structured metadata and context for error values.   [![godoc][D]](https://godoc.org/github.com/Southclaws/fault)
- [errors](https://github.com/emperror/errors) **star:205** Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.   [![It hasn't been updated in recent three years][Y]](https://github.com/emperror/errors)   [![godoc][D]](https://godoc.org/github.com/emperror/errors)
- [errors](https://github.com/naughtygopher/errors) **star:75** Drop-in replacement for builtin Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions.   [![godoc][D]](https://godoc.org/github.com/naughtygopher/errors)
- [exception](https://github.com/rbrahul/exception) **star:38** A simple utility package for exception handling with try-catch in Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/rbrahul/exception)   [![godoc][D]](https://godoc.org/github.com/rbrahul/exception)
- [Falcon](https://github.com/SonicRoshan/falcon) **star:11** A Simple Yet Highly Powerful Package For Error Handling.   [![It hasn't been updated in recent three years][Y]](https://github.com/SonicRoshan/falcon)   [![godoc][D]](https://godoc.org/github.com/SonicRoshan/falcon)
- [errors](https://github.com/PumpkinSeed/errors) **star:7** The most simple error wrapper with awesome performance and minimal memory overhead.   [![It hasn't been updated in recent three years][Y]](https://github.com/PumpkinSeed/errors)   [![godoc][D]](https://godoc.org/github.com/PumpkinSeed/errors)
- [errors](https://gitlab.com/tozd/go/errors)  Providing errors with a stack trace and optional structured details. Compatible with github.com/pkg/errors API but does not use it internally.
- [errors](https://github.com/neuronlabs/errors) **star:6** Simple golang error handling with classification primitives.   [![It hasn't been updated in recent three years][Y]](https://github.com/neuronlabs/errors)   [![godoc][D]](https://godoc.org/github.com/neuronlabs/errors)
- [metaerr](https://github.com/quantumcycle/metaerr) **star:6** A library to create your custom error builders producing structured errors with metadata from different sources and optional stacktraces.   [![godoc][D]](https://godoc.org/github.com/quantumcycle/metaerr)
- [go-errr](https://github.com/go-errr/go) **star:2** Error handling library with Catch/Recover semantics, wrapped error chains, and stack traces for Go.   [![godoc][D]](https://godoc.org/github.com/go-errr/go)

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [pdfcpu](https://github.com/pdfcpu/pdfcpu) **star:8715** PDF processor.   [![There was an update last month][G]](https://github.com/pdfcpu/pdfcpu)   [![godoc][D]](https://godoc.org/github.com/pdfcpu/pdfcpu)
- [afero](https://github.com/spf13/afero) **star:6680** FileSystem Abstraction System for Go.   [![There was an update last month][G]](https://github.com/spf13/afero)   [![godoc][D]](https://godoc.org/github.com/spf13/afero)
- [gdu](https://github.com/dundee/gdu) **star:5813** Disk usage analyzer with console interface.   [![There was an update last month][G]](https://github.com/dundee/gdu)   [![godoc][D]](https://godoc.org/github.com/dundee/gdu)
- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) **star:1180** A package to convert an HTML template to a PDF file.   [![godoc][D]](https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf)
- [notify](https://github.com/rjeczalik/notify) **star:935** File system event notification library with simple API, similar to os/signal.   [![godoc][D]](https://godoc.org/github.com/rjeczalik/notify)
- [copy](https://github.com/otiai10/copy) **star:772** Copy directory recursively.   [![godoc][D]](https://godoc.org/github.com/otiai10/copy)
- [gofs](https://github.com/no-src/gofs) **star:530** A cross-platform real-time file synchronization tool out of the box.   [![godoc][D]](https://godoc.org/github.com/no-src/gofs)   [![Contains Chinese documents][CN]](https://github.com/no-src/gofs)   [![Archived][Archived]](https://github.com/no-src/gofs)
- [afs](https://github.com/viant/afs) **star:391** Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.   [![godoc][D]](https://godoc.org/github.com/viant/afs)
- [vfs](https://github.com/C2FO/vfs) **star:367** A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.   [![godoc][D]](https://godoc.org/github.com/C2FO/vfs)
- [go-exiftool](https://github.com/barasher/go-exiftool) **star:297** Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).   [![godoc][D]](https://godoc.org/github.com/barasher/go-exiftool)
- [iso9660](https://github.com/kdomanski/iso9660) **star:285** A package for reading and creating ISO9660 disk images   [![godoc][D]](https://godoc.org/github.com/kdomanski/iso9660)
- [fastwalk](https://github.com/charlievieth/fastwalk) **star:144** Fast parallel directory traversal library (used by [fzf](https://github.com/junegunn/fzf)).   [![godoc][D]](https://godoc.org/github.com/charlievieth/fastwalk)
- [go-csv-tag](https://github.com/artonge/go-csv-tag) **star:131** Load csv file using tag.   [![godoc][D]](https://godoc.org/github.com/artonge/go-csv-tag)
- [parquet](https://github.com/parsyl/parquet) **star:127** Read and write [parquet](https://parquet.apache.org) files.   [![godoc][D]](https://godoc.org/github.com/parsyl/parquet)
- [checksum](https://github.com/codingsince1985/checksum) **star:113** Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files.   [![godoc][D]](https://godoc.org/github.com/codingsince1985/checksum)
- [skywalker](https://github.com/dixonwille/skywalker) **star:103** Package to allow one to concurrently go through a filesystem with ease.   [![It hasn't been updated in recent three years][Y]](https://github.com/dixonwille/skywalker)   [![godoc][D]](https://godoc.org/github.com/dixonwille/skywalker)
- [opc](https://github.com/qmuntal/opc) **star:80** Load Open Packaging Conventions (OPC) files for Go.   [![godoc][D]](https://godoc.org/github.com/qmuntal/opc)
- [gulter](https://github.com/adelowo/gulter) **star:72** A simple HTTP middleware to automatically handle all your file upload needs   [![godoc][D]](https://godoc.org/github.com/adelowo/gulter)
- [baraka](https://github.com/xis/baraka) **star:65** A library to process http file uploads easily.   [![It hasn't been updated in recent three years][Y]](https://github.com/xis/baraka)   [![godoc][D]](https://godoc.org/github.com/xis/baraka)
- [go-gtfs](https://github.com/artonge/go-gtfs) **star:47** Load gtfs files in go.   [![godoc][D]](https://godoc.org/github.com/artonge/go-gtfs)
- [gxpdf](https://github.com/coregx/gxpdf) **star:45** Modern full-lifecycle PDF library for Go — parse, extract tables, generate, and sign documents with zero CGO dependencies.   [![godoc][D]](https://godoc.org/github.com/coregx/gxpdf)
- [flop](https://github.com/homedepot/flop) **star:36** File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).   [![It hasn't been updated in recent three years][Y]](https://github.com/homedepot/flop)   [![godoc][D]](https://godoc.org/github.com/homedepot/flop)
- [gut/yos](https://github.com/1set/gut) **star:27** Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.   [![It hasn't been updated in recent three years][Y]](https://github.com/1set/gut)   [![godoc][D]](https://godoc.org/github.com/1set/gut)
- [todotxt](https://github.com/1set/todotxt) **star:25** Go library for Gina Trapani's [_todo.txt_](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [_todo.txt_ format](https://github.com/todotxt/todo.txt).   [![godoc][D]](https://godoc.org/github.com/1set/todotxt)
- [higgs](https://github.com/dastoori/higgs) **star:24** A tiny cross-platform Go library to hide/unhide files and directories.   [![godoc][D]](https://godoc.org/github.com/dastoori/higgs)
- [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) **star:22** Copy files for humans.   [![It hasn't been updated in recent three years][Y]](https://github.com/hugocarreira/go-decent-copy)   [![godoc][D]](https://godoc.org/github.com/hugocarreira/go-decent-copy)
- [pathtype](https://github.com/jonchun/pathtype) **star:14** Treat paths as their own type instead of using strings.   [![It hasn't been updated in recent three years][Y]](https://github.com/jonchun/pathtype)   [![godoc][D]](https://godoc.org/github.com/jonchun/pathtype)
- [gopdfrab](https://github.com/voidrab/gopdfrab) **star:2** PDF/A processing for Go.   [![There was an update last month][G]](https://github.com/voidrab/gopdfrab)   [![godoc][D]](https://godoc.org/github.com/voidrab/gopdfrab)

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [decimal](https://github.com/shopspring/decimal) **star:7445** Arbitrary-precision fixed-point decimal numbers.   [![godoc][D]](https://godoc.org/github.com/shopspring/decimal)
- [ticker](https://github.com/achannarasappa/ticker) **star:6142** Terminal stock watcher and stock position tracker.   [![godoc][D]](https://godoc.org/github.com/achannarasappa/ticker)
- [go-money](https://github.com/rhymond/go-money) **star:1902** Implementation of Fowler's Money pattern.   [![godoc][D]](https://godoc.org/github.com/rhymond/go-money)
- [bbgo](https://github.com/c9s/bbgo) **star:1662** A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies.   [![There was an update last month][G]](https://github.com/c9s/bbgo)   [![godoc][D]](https://godoc.org/github.com/c9s/bbgo)
- [ledger](https://github.com/formancehq/ledger) **star:1302** A programmable financial ledger that provides a foundation for money-moving applications.   [![There was an update last month][G]](https://github.com/formancehq/ledger)   [![godoc][D]](https://godoc.org/github.com/formancehq/ledger)
- [indicator](https://github.com/cinar/indicator) **star:1212** Technical analysis library providing financial indicators, strategies, and backtesting framework.   [![godoc][D]](https://godoc.org/github.com/cinar/indicator)
- [accounting](https://github.com/leekchan/accounting) **star:914** money and currency formatting for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/leekchan/accounting)   [![godoc][D]](https://godoc.org/github.com/leekchan/accounting)
- [techan](https://github.com/sdcoffey/techan) **star:906** Technical analysis library with advanced market analysis and trading strategies.   [![godoc][D]](https://godoc.org/github.com/sdcoffey/techan)
- [currency](https://github.com/bojanz/currency) **star:641** Handles currency amounts, provides currency information and formatting.   [![godoc][D]](https://godoc.org/github.com/bojanz/currency)
- [ach](https://github.com/moov-io/ach) **star:555** A reader, writer, and validator for Automated Clearing House (ACH) files.   [![There was an update last month][G]](https://github.com/moov-io/ach)   [![godoc][D]](https://godoc.org/github.com/moov-io/ach)
- [orderbook](https://github.com/i25959341/orderbook) **star:555** Matching Engine for Limit Order Book in Golang.   [![godoc][D]](https://godoc.org/github.com/i25959341/orderbook)
- [gobl](https://github.com/invopop/gobl) **star:294** Invoice and billing document framework. JSON Schema based. Automates tax calculations and validation, with tooling to convert into global formats.   [![There was an update last month][G]](https://github.com/invopop/gobl)   [![godoc][D]](https://godoc.org/github.com/invopop/gobl)
- [decimal](https://github.com/govalues/decimal) **star:246** Immutable decimal numbers with panic-free arithmetic.   [![godoc][D]](https://godoc.org/github.com/govalues/decimal)
- [go-finance](https://github.com/alpeb/go-finance) **star:193** Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.   [![It hasn't been updated in recent three years][Y]](https://github.com/alpeb/go-finance)   [![godoc][D]](https://godoc.org/github.com/alpeb/go-finance)
- [udecimal](https://github.com/quagmt/udecimal) **star:189** High performance, high precision, zero allocation fixed-point decimal library for financial applications.   [![godoc][D]](https://godoc.org/github.com/quagmt/udecimal)
- [ofxgo](https://github.com/aclindsa/ofxgo) **star:152** Query OFX servers and/or parse the responses (with example command-line client).   [![godoc][D]](https://godoc.org/github.com/aclindsa/ofxgo)
- [transaction](https://github.com/claygod/transaction) **star:139** Embedded transactional database of accounts, running in multithreaded mode.   [![godoc][D]](https://godoc.org/github.com/claygod/transaction)
- [vat](https://github.com/dannyvankooten/vat) **star:119** VAT number validation & EU VAT rates.   [![godoc][D]](https://godoc.org/github.com/dannyvankooten/vat)
- [payme](https://github.com/jovandeginste/payme) **star:91** QR code generator (ASCII & PNG) for SEPA payments.   [![godoc][D]](https://godoc.org/github.com/jovandeginste/payme)
- [currency](https://github.com/naughtygopher/currency) **star:61** High performant & accurate currency computation package.   [![godoc][D]](https://godoc.org/github.com/naughtygopher/currency)
- [money](https://github.com/govalues/money) **star:54** Immutable monetary amounts and exchange rates with panic-free arithmetic.   [![godoc][D]](https://godoc.org/github.com/govalues/money)
- [dec128](https://github.com/jokruger/dec128) **star:47** High performance 128-bit fixed-point decimal numbers.   [![godoc][D]](https://godoc.org/github.com/jokruger/dec128)
- [fpmoney](https://github.com/nikolaydubina/fpmoney) **star:36** Fast and simple ISO4217 fixed-point decimal money.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/fpmoney)
- [fpdecimal](https://github.com/nikolaydubina/fpdecimal) **star:35** Fast and precise serialization and arithmetic for small fixed-point decimals   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/fpdecimal)
- [go-finance](https://github.com/pieterclaerhout/go-finance) **star:32** Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.   [![godoc][D]](https://godoc.org/github.com/pieterclaerhout/go-finance)
- [go-nowpayments](https://github.com/matm/go-nowpayments) **star:5** Library for the crypto NOWPayments API.   [![It hasn't been updated in recent three years][Y]](https://github.com/matm/go-nowpayments)   [![godoc][D]](https://godoc.org/github.com/matm/go-nowpayments)
- [paystack-sdk-go](https://github.com/samaasi/paystack-sdk-go) **star:3** A comprehensive, zero-dependency, and fully typed Go SDK for the Paystack API.   [![godoc][D]](https://godoc.org/github.com/samaasi/paystack-sdk-go)
- [swift](https://code.pfad.fr/swift/)  Offline validity check of IBAN (International Bank Account Number) and retrieval of BIC (for some countries).
- [decimal](https://github.com/aytechnet/decimal) **star:1** High performance 64-bit decimal partially compatible with [shopspring/decimal](https://github.com/shopspring/decimal) and int64, including Weight and Length.   [![godoc][D]](https://godoc.org/github.com/aytechnet/decimal)

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [nosurf](https://github.com/justinas/nosurf) **star:1741** CSRF protection middleware for Go.   [![godoc][D]](https://godoc.org/github.com/justinas/nosurf)
- [gorilla/csrf](https://github.com/gorilla/csrf) **star:1200** CSRF protection for Go web applications & services.   [![godoc][D]](https://godoc.org/github.com/gorilla/csrf)
- [form](https://github.com/go-playground/form) **star:920** Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.   [![godoc][D]](https://godoc.org/github.com/go-playground/form)
- [httpin](https://github.com/ggicci/httpin) **star:388** Decode an HTTP request into a custom struct, including querystring, forms, HTTP headers, etc.   [![godoc][D]](https://godoc.org/github.com/ggicci/httpin)
- [conform](https://github.com/leebenson/conform) **star:325** Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.   [![godoc][D]](https://godoc.org/github.com/leebenson/conform)
- [formam](https://github.com/monoculum/formam) **star:192** decode form's values into a struct.   [![It hasn't been updated in recent three years][Y]](https://github.com/monoculum/formam)   [![godoc][D]](https://godoc.org/github.com/monoculum/formam)
- [forms](https://github.com/albrow/forms) **star:142** Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.   [![It hasn't been updated in recent three years][Y]](https://github.com/albrow/forms)   [![godoc][D]](https://godoc.org/github.com/albrow/forms)
- [qs](https://github.com/sonh/qs) **star:82** Go module for encoding structs into URL query parameters.   [![godoc][D]](https://godoc.org/github.com/sonh/qs)
- [checker](https://github.com/cinar/checker) **star:48** Checker helps validating user input through rules defined in struct tags or directly through functions.   [![godoc][D]](https://godoc.org/github.com/cinar/checker)
- [bind](https://github.com/robfig/bind) **star:32** Bind form data to any Go values.   [![It hasn't been updated in recent three years][Y]](https://github.com/robfig/bind)   [![godoc][D]](https://godoc.org/github.com/robfig/bind)
- [queryparam](https://github.com/tomwright/queryparam) **star:19** Decode `url.Values` into usable struct values of standard or custom types.   [![It hasn't been updated in recent three years][Y]](https://github.com/tomwright/queryparam)   [![godoc][D]](https://godoc.org/github.com/tomwright/queryparam)
- [gbind](https://github.com/bdjimmy/gbind) **star:10** Bind data to any Go value. Can use built-in and custom expression binding capabilities; supports data validation   [![It hasn't been updated in recent three years][Y]](https://github.com/bdjimmy/gbind)   [![godoc][D]](https://godoc.org/github.com/bdjimmy/gbind)   [![Contains Chinese documents][CN]](https://github.com/bdjimmy/gbind)
- [roamer](https://github.com/slipros/roamer) **star:6** Eliminates boilerplate code for parsing HTTP requests by binding cookies, headers, query params, path params, body to structs and more by using simple tags.   [![godoc][D]](https://godoc.org/github.com/slipros/roamer)

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [mo](https://github.com/samber/mo) **star:3383** Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...).   [![godoc][D]](https://godoc.org/github.com/samber/mo)
- [go-underscore](https://github.com/tobyhede/go-underscore) **star:1298** Useful collection of helpfully functional Go collection utilities.   [![It hasn't been updated in recent three years][Y]](https://github.com/tobyhede/go-underscore)   [![godoc][D]](https://godoc.org/github.com/tobyhede/go-underscore)
- [go-functional](https://github.com/BooleanCat/go-functional) **star:537** Functional programming in Go using generics   [![godoc][D]](https://godoc.org/github.com/BooleanCat/go-functional)
- [fpGo](https://github.com/TeaEntityLab/fpGo) **star:355** Monad, Functional Programming features for Golang.   [![godoc][D]](https://godoc.org/github.com/TeaEntityLab/fpGo)
- [fp-go](https://github.com/repeale/fp-go) **star:325** Collection of Functional Programming helpers powered by Golang 1.18+ generics.   [![It hasn't been updated in recent three years][Y]](https://github.com/repeale/fp-go)   [![godoc][D]](https://godoc.org/github.com/repeale/fp-go)
- [FuncFrog](https://github.com/koss-null/FuncFrog) **star:285** Functional helpers library providing Map, Filter, Reduce and other stream operations on generic slices Go1.18+ with lazy evaluation and error handling mechanisms.   [![godoc][D]](https://godoc.org/github.com/koss-null/FuncFrog)
- [gofp](https://github.com/rbrahul/gofp) **star:153** A lodash like powerful utility library for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/rbrahul/gofp)   [![godoc][D]](https://godoc.org/github.com/rbrahul/gofp)
- [fuego](https://github.com/seborama/fuego) **star:145** Functional Experiment in Go.   [![godoc][D]](https://godoc.org/github.com/seborama/fuego)
- [underscore](https://github.com/rjNemo/underscore) **star:118** Functional programming helpers for Go 1.18 and beyond.   [![godoc][D]](https://godoc.org/github.com/rjNemo/underscore)
- [g](https://github.com/enetx/g) **star:61** Functional programming framework for Go.   [![There was an update last month][G]](https://github.com/enetx/g)   [![godoc][D]](https://godoc.org/github.com/enetx/g)
- [valor](https://github.com/phelmkamp/valor) **star:20** Generic option and result types that optionally contain a value.   [![It hasn't been updated in recent three years][Y]](https://github.com/phelmkamp/valor)   [![godoc][D]](https://godoc.org/github.com/phelmkamp/valor)

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Ebitengine](https://github.com/hajimehoshi/ebiten) **star:13335** dead simple 2D game engine in Go.   [![There was an update last month][G]](https://github.com/hajimehoshi/ebiten)   [![godoc][D]](https://godoc.org/github.com/hajimehoshi/ebiten)
- [Leaf](https://github.com/name5566/leaf) **star:5511** Lightweight game server framework.   [![godoc][D]](https://godoc.org/github.com/name5566/leaf)   [![Contains Chinese documents][CN]](https://github.com/name5566/leaf)
- [nano](https://github.com/lonng/nano) **star:3205** Lightweight, facility, high performance golang based game server framework.   [![godoc][D]](https://godoc.org/github.com/lonng/nano)   [![Contains Chinese documents][CN]](https://github.com/lonng/nano)
- [g3n](https://github.com/g3n/engine) **star:3099** Go 3D Game Engine.   [![godoc][D]](https://godoc.org/github.com/g3n/engine)
- [Pitaya](https://github.com/topfreegames/pitaya) **star:2805** Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.   [![godoc][D]](https://godoc.org/github.com/topfreegames/pitaya)
- [goworld](https://github.com/xiaonanln/goworld) **star:2718** Scalable game server engine, featuring space-entity framework and hot-swapping.   [![godoc][D]](https://godoc.org/github.com/xiaonanln/goworld)   [![Contains Chinese documents][CN]](https://github.com/xiaonanln/goworld)
- [raylib-go](https://github.com/gen2brain/raylib-go) **star:2482** Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
- [go-sdl2](https://github.com/veandco/go-sdl2) **star:2324** Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
- [engo](https://github.com/EngoEngine/engo) **star:1817** Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.   [![godoc][D]](https://godoc.org/github.com/EngoEngine/engo)
- [Oak](https://github.com/oakmound/oak) **star:1661** Pure Go game engine.   [![godoc][D]](https://godoc.org/github.com/oakmound/oak)
- [termloop](https://github.com/JoelOtter/termloop) **star:1476** Terminal-based game engine for Go, built on top of Termbox.   [![godoc][D]](https://godoc.org/github.com/JoelOtter/termloop)
- [gonet](https://github.com/xtaci/gonet) **star:1287** Game server skeleton implemented with golang.   [![godoc][D]](https://godoc.org/github.com/xtaci/gonet)   [![Archived][Archived]](https://github.com/xtaci/gonet)
- [go-astar](https://github.com/beefsack/go-astar) **star:629** Go implementation of the A\* path finding algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/beefsack/go-astar)   [![godoc][D]](https://godoc.org/github.com/beefsack/go-astar)
- [Pixel](https://github.com/gopxl/pixel) **star:390** Hand-crafted 2D game library in Go.   [![godoc][D]](https://godoc.org/github.com/gopxl/pixel)
- [go3d](https://github.com/ungerik/go3d) **star:341** Performance oriented 2D/3D math package for Go.   [![godoc][D]](https://godoc.org/github.com/ungerik/go3d)
- [gogpu](https://github.com/gogpu/gogpu) **star:340** GPU application framework with windowing, input, and rendering built on WebGPU — reduces 480+ lines of GPU code to ~20, zero CGO (GoGPU ecosystem: [gg](https://github.com/gogpu/gg), [ui](https://github.com/gogpu/ui), [wgpu](https://github.com/gogpu/wgpu), [naga](https://github.com/gogpu/naga)).   [![There was an update last month][G]](https://github.com/gogpu/gogpu)   [![godoc][D]](https://godoc.org/github.com/gogpu/gogpu)
- [Ark](https://github.com/mlange-42/ark) **star:286** Archetype-based Entity Component System (ECS) for Go.   [![godoc][D]](https://godoc.org/github.com/mlange-42/ark)
- [tile](https://github.com/kelindar/tile) **star:225** Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export.   [![godoc][D]](https://godoc.org/github.com/kelindar/tile)
- [ecs](https://github.com/andygeiss/ecs) **star:176** Build your own Game-Engine based on the Entity Component System concept in Golang.   [![godoc][D]](https://godoc.org/github.com/andygeiss/ecs)
- [gogpu/wgpu](https://github.com/gogpu/wgpu) **star:156** Pure Go WebGPU implementation with Vulkan, DX12, and Metal backends, zero CGO (part of [GoGPU](https://github.com/gogpu) ecosystem).   [![There was an update last month][G]](https://github.com/gogpu/wgpu)   [![godoc][D]](https://godoc.org/github.com/gogpu/wgpu)
- [prototype](https://github.com/gonutz/prototype) **star:108** Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API.
- [GOKe](https://github.com/kjkrol/goke) **star:84** Data-Oriented (DOD), archetype-based ECS engine utilizing an L1 cache-aligned chunked SoA layout for predictable, stepless memory growth and zero-allocation execution paths.   [![There was an update last month][G]](https://github.com/kjkrol/goke)   [![godoc][D]](https://godoc.org/github.com/kjkrol/goke)
- [Pi](https://github.com/elgopher/pi) **star:69** Game engine for creating retro games for modern computers. Inspired by Pico-8 and powered by Ebitengine.   [![godoc][D]](https://godoc.org/github.com/elgopher/pi)
- [fantasyname](https://github.com/s0rg/fantasyname) **star:44** Fantasy names generator.   [![godoc][D]](https://godoc.org/github.com/s0rg/fantasyname)
- [grid](https://github.com/s0rg/grid) **star:27** Generic 2D grid with ray-casting, shadow-casting and path finding.   [![godoc][D]](https://godoc.org/github.com/s0rg/grid)

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [oapi-codegen](https://github.com/deepmap/oapi-codegen) **star:8448** This package contains a set of utilities for generating Go boilerplate code for services based on OpenAPI 3.0 API definitions.   [![There was an update last month][G]](https://github.com/deepmap/oapi-codegen)   [![godoc][D]](https://godoc.org/github.com/deepmap/oapi-codegen)
- [go-linq](https://github.com/ahmetalpbalkan/go-linq) **star:3657** .NET LINQ-like query methods for Go.   [![godoc][D]](https://godoc.org/github.com/ahmetalpbalkan/go-linq)
- [jennifer](https://github.com/dave/jennifer) **star:3623** Generate arbitrary Go code without templates.   [![godoc][D]](https://godoc.org/github.com/dave/jennifer)
- [GoWrap](https://github.com/hexdigest/gowrap) **star:1330** Generate decorators for Go interfaces using simple templates.   [![There was an update last month][G]](https://github.com/hexdigest/gowrap)   [![godoc][D]](https://godoc.org/github.com/hexdigest/gowrap)
- [goderive](https://github.com/awalterschulze/goderive) **star:1264** Derives functions from input types   [![godoc][D]](https://godoc.org/github.com/awalterschulze/goderive)
- [go-enum](https://github.com/abice/go-enum) **star:953** Code generation for enums from code comments.   [![There was an update last month][G]](https://github.com/abice/go-enum)   [![godoc][D]](https://godoc.org/github.com/abice/go-enum)
- [goverter](https://github.com/jmattheis/goverter) **star:861** Generate converters by defining an interface.   [![godoc][D]](https://godoc.org/github.com/jmattheis/goverter)
- [interfaces](https://github.com/rjeczalik/interfaces) **star:431** Command line tool for generating interface definitions.   [![godoc][D]](https://godoc.org/github.com/rjeczalik/interfaces)
- [copygen](https://github.com/switchupcb/copygen) **star:404** Generate any code based on Go types, including type-to-type converters (copy code) without reflection by default.   [![godoc][D]](https://godoc.org/github.com/switchupcb/copygen)
- [convergen](https://github.com/reedom/convergen) **star:50** Feature rich type-to-type copy code generator.   [![godoc][D]](https://godoc.org/github.com/reedom/convergen)
- [generis](https://github.com/senselogic/GENERIS) **star:47** Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.   [![It hasn't been updated in recent three years][Y]](https://github.com/senselogic/GENERIS)
- [go-apispec](https://github.com/antst/go-apispec) **star:31** Generate OpenAPI 3.1 specs from Go source code via static analysis with automatic framework detection.   [![There was an update last month][G]](https://github.com/antst/go-apispec)   [![godoc][D]](https://godoc.org/github.com/antst/go-apispec)
- [protoc-gen-httpgo](https://github.com/MUlt1mate/protoc-gen-httpgo) **star:25** Generate HTTP server and client from protobuf.   [![godoc][D]](https://godoc.org/github.com/MUlt1mate/protoc-gen-httpgo)
- [typeregistry](https://github.com/xiaoxin01/typeregistry) **star:23** A library to create type dynamically.   [![It hasn't been updated in recent three years][Y]](https://github.com/xiaoxin01/typeregistry)   [![godoc][D]](https://godoc.org/github.com/xiaoxin01/typeregistry)
- [go-enum-encoding](https://github.com/nikolaydubina/go-enum-encoding) **star:16** Code generation for enum encoding from code comments.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/go-enum-encoding)

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [Tile38](https://github.com/tidwall/tile38) **star:9689** Geolocation DB with spatial index and realtime geofencing.   [![There was an update last month][G]](https://github.com/tidwall/tile38)   [![godoc][D]](https://godoc.org/github.com/tidwall/tile38)
- [S2 geometry](https://github.com/golang/geo) **star:1845** S2 geometry library in Go.   [![There was an update last month][G]](https://github.com/golang/geo)   [![godoc][D]](https://godoc.org/github.com/golang/geo)
- [mbtileserver](https://github.com/consbio/mbtileserver) **star:787** A simple Go-based server for map tiles stored in mbtiles format.   [![godoc][D]](https://godoc.org/github.com/consbio/mbtileserver)
- [geoos](https://github.com/spatial-go/geoos) **star:530** A library provides spatial data and geometric algorithms.   [![godoc][D]](https://godoc.org/github.com/spatial-go/geoos)
- [osm](https://github.com/paulmach/osm) **star:466** Library for reading, writing and working with OpenStreetMap data and APIs.   [![godoc][D]](https://godoc.org/github.com/paulmach/osm)
- [H3](https://github.com/uber/h3-go) **star:441** Go bindings for H3, a hierarchical hexagonal geospatial indexing system.   [![godoc][D]](https://godoc.org/github.com/uber/h3-go)
- [godal](https://github.com/airbusgeo/godal) **star:179** Go wrapper for GDAL.   [![godoc][D]](https://godoc.org/github.com/airbusgeo/godal)
- [simplefeatures](https://github.com/peterstace/simplefeatures) **star:172** simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them.   [![godoc][D]](https://godoc.org/github.com/peterstace/simplefeatures)
- [geoserver](https://github.com/hishamkaram/geoserver) **star:93** geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.   [![godoc][D]](https://godoc.org/github.com/hishamkaram/geoserver)
- [pbf](https://github.com/maguro/pbf) **star:56** OpenStreetMap PBF golang encoder/decoder.   [![godoc][D]](https://godoc.org/github.com/maguro/pbf)
- [gismanager](https://github.com/hishamkaram/gismanager) **star:55** Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.   [![There was an update last month][G]](https://github.com/hishamkaram/gismanager)   [![godoc][D]](https://godoc.org/github.com/hishamkaram/gismanager)
- [S2 geojson](https://github.com/pantrif/s2-geojson) **star:37** Convert geojson to s2 cells & demonstrating some S2 geometry features on map.   [![godoc][D]](https://godoc.org/github.com/pantrif/s2-geojson)
- [H3 GeoJSON](https://github.com/mmadfox/go-geojson2h3) **star:7** Conversion utilities between H3 indexes and GeoJSON.   [![godoc][D]](https://godoc.org/github.com/mmadfox/go-geojson2h3)
- [H3GeoDist](https://github.com/mmadfox/go-h3geo-dist) **star:2** Distribution of Uber H3geo cells by virtual nodes.   [![It hasn't been updated in recent three years][Y]](https://github.com/mmadfox/go-h3geo-dist)   [![godoc][D]](https://godoc.org/github.com/mmadfox/go-h3geo-dist)
- [borders](https://github.com/kpfaulkner/borders) **star:1** Detects image borders and converts to GeoJSON for GIS operations.   [![godoc][D]](https://godoc.org/github.com/kpfaulkner/borders)
- [Web-Mercator-Projection](https://github.com/jorelosorio/web-mercator-projection) A project to easily use and convert LonLat, Point and Tile to display info, markers, etc, in a map using the Web Mercator Projection.
- [WGS84](https://github.com/wroge/wgs84) **star:142** Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).   [![godoc][D]](https://godoc.org/github.com/wroge/wgs84)

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages and vice-versa._

- [gopherjs](https://github.com/gopherjs/gopherjs) **star:13170** Compiler from Go to JavaScript.   [![godoc][D]](https://godoc.org/github.com/gopherjs/gopherjs)
- [bunster](https://github.com/yassinebenaid/bunster) **star:2674** Compile shell scripts to Go.   [![godoc][D]](https://godoc.org/github.com/yassinebenaid/bunster)
- [cxgo](https://github.com/gotranspile/cxgo) **star:394** Transpile C code to Go code.   [![godoc][D]](https://godoc.org/github.com/gotranspile/cxgo)
- [c4go](https://github.com/Konstantin8105/c4go) **star:376** Transpile C code to Go code.
- [go2hx](https://github.com/go2hx/go2hx) **star:152** Compiler from Go to Haxe to Javascript/C++/Java/C#.
- [esp32](https://github.com/andygeiss/esp32-transpiler) **star:100** Transpile Go into Arduino code.   [![godoc][D]](https://godoc.org/github.com/andygeiss/esp32-transpiler)
- [f4go](https://github.com/Konstantin8105/f4go) **star:51** Transpile FORTRAN 77 code to Go code.   [![godoc][D]](https://godoc.org/github.com/Konstantin8105/f4go)

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [ants](https://github.com/panjf2000/ants) **star:14464** A high-performance and low-cost goroutine pool in Go.   [![godoc][D]](https://godoc.org/github.com/panjf2000/ants)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/ants)
- [conc](https://github.com/sourcegraph/conc) **star:10415** `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer.   [![godoc][D]](https://godoc.org/github.com/sourcegraph/conc)
- [tunny](https://github.com/Jeffail/tunny) **star:4033** Goroutine pool for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/Jeffail/tunny)   [![godoc][D]](https://godoc.org/github.com/Jeffail/tunny)
- [goworker](https://github.com/benmanns/goworker) **star:2847** goworker is a Go-based background worker.   [![godoc][D]](https://godoc.org/github.com/benmanns/goworker)
- [pond](https://github.com/alitto/pond) **star:2172** Minimalistic and High-performance goroutine worker pool written in Go.   [![godoc][D]](https://godoc.org/github.com/alitto/pond)
- [rill](https://github.com/destel/rill) **star:1846** Go toolkit for clean, composable, channel-based concurrency.   [![There was an update last month][G]](https://github.com/destel/rill)   [![godoc][D]](https://godoc.org/github.com/destel/rill)
- [workerpool](https://github.com/gammazero/workerpool) **star:1459** Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.   [![godoc][D]](https://godoc.org/github.com/gammazero/workerpool)
- [grpool](https://github.com/ivpusic/grpool) **star:735** Lightweight Goroutine pool.   [![It hasn't been updated in recent three years][Y]](https://github.com/ivpusic/grpool)   [![godoc][D]](https://godoc.org/github.com/ivpusic/grpool)
- [pool](https://github.com/go-playground/pool) **star:724** Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/pool)   [![godoc][D]](https://godoc.org/github.com/go-playground/pool)
- [gowp](https://github.com/xxjwxc/gowp) **star:516** gowp is concurrency limiting goroutine pool.   [![godoc][D]](https://godoc.org/github.com/xxjwxc/gowp)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/gowp)
- [flowmatic](https://github.com/carlmjohnson/flowmatic) **star:402** Structured concurrency made easy.   [![godoc][D]](https://godoc.org/github.com/carlmjohnson/flowmatic)
- [async](https://github.com/reugn/async) **star:309** An alternative sync library for Go (Future, Promise, Locks).   [![godoc][D]](https://godoc.org/github.com/reugn/async)
- [go-actor](https://github.com/vladopajic/go-actor) **star:301** A tiny library for writing concurrent programs using actor model.   [![There was an update last month][G]](https://github.com/vladopajic/go-actor)   [![godoc][D]](https://godoc.org/github.com/vladopajic/go-actor)
- [routine](https://github.com/timandy/routine) **star:292** `routine` is a `ThreadLocal` for go library. It encapsulates and provides some easy-to-use, non-competitive, high-performance `goroutine` context access interfaces, which can help you access coroutine context information more gracefully.   [![godoc][D]](https://godoc.org/github.com/timandy/routine)   [![Contains Chinese documents][CN]](https://github.com/timandy/routine)
- [go-floc](https://github.com/workanator/go-floc) **star:269** Orchestrate goroutines with ease.   [![It hasn't been updated in recent three years][Y]](https://github.com/workanator/go-floc)   [![godoc][D]](https://godoc.org/github.com/workanator/go-floc)
- [go-flow](https://github.com/kamildrazkiewicz/go-flow) **star:221** Control goroutines execution order.   [![It hasn't been updated in recent three years][Y]](https://github.com/kamildrazkiewicz/go-flow)   [![godoc][D]](https://godoc.org/github.com/kamildrazkiewicz/go-flow)
- [artifex](https://github.com/borderstech/artifex) **star:214** Simple in-memory job queue for Golang using worker-based dispatching.   [![godoc][D]](https://godoc.org/github.com/borderstech/artifex)
- [semaphore](https://github.com/marusama/semaphore) **star:178** Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).   [![It hasn't been updated in recent three years][Y]](https://github.com/marusama/semaphore)   [![godoc][D]](https://godoc.org/github.com/marusama/semaphore)
- [neilotoole/errgroup](https://github.com/neilotoole/errgroup) **star:163** Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines.   [![It hasn't been updated in recent three years][Y]](https://github.com/neilotoole/errgroup)   [![godoc][D]](https://godoc.org/github.com/neilotoole/errgroup)
- [cyclicbarrier](https://github.com/marusama/cyclicbarrier) **star:158** CyclicBarrier for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/marusama/cyclicbarrier)   [![godoc][D]](https://godoc.org/github.com/marusama/cyclicbarrier)
- [async](https://github.com/studiosol/async) **star:140** A safe way to execute functions asynchronously, recovering them in case of panic.   [![It hasn't been updated in recent three years][Y]](https://github.com/studiosol/async)   [![godoc][D]](https://godoc.org/github.com/studiosol/async)
- [gollback](https://github.com/vardius/gollback) **star:123** asynchronous simple function utilities, for managing execution of closures and callbacks.   [![It hasn't been updated in recent three years][Y]](https://github.com/vardius/gollback)   [![godoc][D]](https://godoc.org/github.com/vardius/gollback)
- [Hunch](https://github.com/AaronJan/Hunch) **star:104** Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.   [![It hasn't been updated in recent three years][Y]](https://github.com/AaronJan/Hunch)   [![godoc][D]](https://godoc.org/github.com/AaronJan/Hunch)   [![Archived][Archived]](https://github.com/AaronJan/Hunch)
- [threadpool](https://github.com/shettyh/threadpool) **star:103** Golang threadpool implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/shettyh/threadpool)   [![godoc][D]](https://godoc.org/github.com/shettyh/threadpool)
- [semaphore](https://github.com/kamilsk/semaphore) **star:101** Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.   [![It hasn't been updated in recent three years][Y]](https://github.com/kamilsk/semaphore)   [![godoc][D]](https://godoc.org/github.com/kamilsk/semaphore)
- [worker-pool](https://github.com/vardius/worker-pool) **star:93** goworker is a Go simple async worker pool.   [![It hasn't been updated in recent three years][Y]](https://github.com/vardius/worker-pool)   [![godoc][D]](https://godoc.org/github.com/vardius/worker-pool)
- [gpool](https://github.com/Sherifabdlnaby/gpool) **star:91** manages a resizeable pool of context-aware goroutines to bound concurrency.   [![There was an update last month][G]](https://github.com/Sherifabdlnaby/gpool)   [![godoc][D]](https://godoc.org/github.com/Sherifabdlnaby/gpool)
- [goccm](https://github.com/zenthangplus/goccm) **star:73** Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently.   [![It hasn't been updated in recent three years][Y]](https://github.com/zenthangplus/goccm)   [![godoc][D]](https://godoc.org/github.com/zenthangplus/goccm)
- [gowl](https://github.com/hamed-yousefi/gowl) **star:71** Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status.   [![godoc][D]](https://godoc.org/github.com/hamed-yousefi/gowl)
- [nursery](https://github.com/arunsworld/nursery) **star:71** Structured concurrency in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/arunsworld/nursery)   [![godoc][D]](https://godoc.org/github.com/arunsworld/nursery)
- [oversight](https://pkg.go.dev/cirello.io/oversight)  Oversight is a complete implementation of the Erlang supervision trees.
- [routine](https://github.com/x-mod/routine) **star:63** go routine control with context, support: Main, Go, Pool and some useful Executors.   [![godoc][D]](https://godoc.org/github.com/x-mod/routine)
- [gohive](https://github.com/loveleshsharma/gohive) **star:54** A highly performant and easy to use Goroutine pool for Go.   [![godoc][D]](https://godoc.org/github.com/loveleshsharma/gohive)
- [kyoo](https://github.com/dirkaholic/kyoo) **star:52** Provides an unlimited job queue and concurrent worker pools.   [![godoc][D]](https://godoc.org/github.com/dirkaholic/kyoo)
- [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) **star:49** Like `sync.WaitGroup` with error handling and concurrency control.   [![godoc][D]](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)
- [parallel-fn](https://github.com/rafaeljesus/parallel-fn) **star:37** Run functions in parallel.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafaeljesus/parallel-fn)   [![godoc][D]](https://godoc.org/github.com/rafaeljesus/parallel-fn)
- [go-trylock](https://github.com/subchen/go-trylock) **star:36** TryLock support on read-write lock for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/subchen/go-trylock)   [![godoc][D]](https://godoc.org/github.com/subchen/go-trylock)
- [channelify](https://github.com/ddelizia/channelify) **star:34** Transform your function to return channels for easy and powerful parallel processing.   [![It hasn't been updated in recent three years][Y]](https://github.com/ddelizia/channelify)   [![godoc][D]](https://godoc.org/github.com/ddelizia/channelify)
- [stl](https://github.com/ssgreg/stl) **star:31** Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.   [![It hasn't been updated in recent three years][Y]](https://github.com/ssgreg/stl)   [![godoc][D]](https://godoc.org/github.com/ssgreg/stl)
- [execpool](https://github.com/hexdigest/execpool) **star:29** A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command.   [![It hasn't been updated in recent three years][Y]](https://github.com/hexdigest/execpool)   [![godoc][D]](https://godoc.org/github.com/hexdigest/execpool)
- [breaker](https://github.com/kamilsk/breaker) **star:21** Flexible mechanism to make execution flow interruptible.   [![It hasn't been updated in recent three years][Y]](https://github.com/kamilsk/breaker)
- [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) **star:19** Concurrency limiter with support for timeouts, dynamic priority and context cancellation of goroutines.   [![godoc][D]](https://godoc.org/github.com/vivek-ng/concurrency-limiter)
- [conexec](https://github.com/ITcathyh/conexec) **star:18** A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.   [![godoc][D]](https://godoc.org/github.com/ITcathyh/conexec)
- [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) **star:17** Manage a pool of goroutines using this lightweight library with a simple API.   [![It hasn't been updated in recent three years][Y]](https://github.com/nikhilsaraf/go-tools)   [![godoc][D]](https://godoc.org/github.com/nikhilsaraf/go-tools)
- [async](https://github.com/yaitoo/async) **star:16** An asynchronous task package with async/await style for Go.   [![godoc][D]](https://godoc.org/github.com/yaitoo/async)
- [go-workerpool](https://github.com/zenthangplus/go-workerpool) **star:13** Inspired from Java Thread Pool, Go WorkerPool aims to control heavy Go Routines.   [![It hasn't been updated in recent three years][Y]](https://github.com/zenthangplus/go-workerpool)   [![godoc][D]](https://godoc.org/github.com/zenthangplus/go-workerpool)
- [async-job](https://github.com/lab210-dev/async-job) **star:12** AsyncJob is an asynchronous queue job manager with light code, clear and speed.   [![It hasn't been updated in recent three years][Y]](https://github.com/lab210-dev/async-job)   [![godoc][D]](https://godoc.org/github.com/lab210-dev/async-job)
- [go-accumulator](https://github.com/nar10z/go-accumulator) **star:11** Solution for accumulation of events and their subsequent processing.   [![godoc][D]](https://godoc.org/github.com/nar10z/go-accumulator)
- [hands](https://github.com/duanckham/hands) **star:10** A process controller used to control the execution and return strategies of multiple goroutines.   [![godoc][D]](https://godoc.org/github.com/duanckham/hands)
- [autopool](https://github.com/AshvinBambhaniya/autopool) **star:5** Zero-config, auto-scaling worker pool for Go with priority-aware scheduling.   [![godoc][D]](https://godoc.org/github.com/AshvinBambhaniya/autopool)
- [anchor](https://github.com/kyuff/anchor) **star:4** Library to manage component lifecycle in microservice architectures.   [![There was an update last month][G]](https://github.com/kyuff/anchor)   [![godoc][D]](https://godoc.org/github.com/kyuff/anchor)
- [powerlock](https://github.com/donomii/powerlock) **star:2** Named FIFO mutexes with context cancellation, bounded wait queues, watchdog diagnostics, pprof profiles, and Prometheus metrics.   [![There was an update last month][G]](https://github.com/donomii/powerlock)   [![godoc][D]](https://godoc.org/github.com/donomii/powerlock)

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [fyne](https://github.com/fyne-io/fyne) **star:28469** Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.   [![There was an update last month][G]](https://github.com/fyne-io/fyne)   [![godoc][D]](https://godoc.org/github.com/fyne-io/fyne)
- [gio](https://gioui.org)  Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD, OpenBSD and WebAssembly.
- [go-gtk](https://mattn.github.io/go-gtk/)  Go bindings for GTK.
- [webview](https://github.com/zserge/webview) **star:14155** Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).
- [qt](https://github.com/therecipe/qt) **star:10811** Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).   [![godoc][D]](https://godoc.org/github.com/therecipe/qt)
- [app](https://github.com/murlokswarm/app) **star:8933** Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.   [![godoc][D]](https://godoc.org/github.com/murlokswarm/app)
- [ui](https://github.com/andlabs/ui) **star:8360** Platform-native GUI library for Go. Cross platform.   [![It hasn't been updated in recent three years][Y]](https://github.com/andlabs/ui)   [![godoc][D]](https://godoc.org/github.com/andlabs/ui)
- [walk](https://github.com/lxn/walk) **star:7096** Windows application library kit for Go.   [![godoc][D]](https://godoc.org/github.com/lxn/walk)
- [DarwinKit](https://github.com/progrium/darwinkit) **star:5435** Build native macOS applications using Go.   [![godoc][D]](https://godoc.org/github.com/progrium/darwinkit)
- [go-sciter](https://github.com/sciter-sdk/go-sciter) **star:2621** Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.   [![It hasn't been updated in recent three years][Y]](https://github.com/sciter-sdk/go-sciter)
- [Goey](https://bitbucket.org/rj/goey/src/master/)  Cross platform UI toolkit aggregator for Windows / Linux / Mac. GTK, Cocoa, Windows API
- [Cogent Core](https://github.com/cogentcore/core) **star:2343** A framework for building 2D and 3D apps that run on macOS, Windows, Linux, iOS, Android, and the web.   [![godoc][D]](https://godoc.org/github.com/cogentcore/core)
- [gotk3](https://github.com/gotk3/gotk3) **star:2213** Go bindings for GTK3.   [![godoc][D]](https://godoc.org/github.com/gotk3/gotk3)
- [Spot](https://github.com/roblillack/spot) **star:1256** Reactive, cross-platform desktop GUI toolkit.   [![godoc][D]](https://godoc.org/github.com/roblillack/spot)
- [energy](https://github.com/energye/energy) **star:606** Cross-platform based on LCL(Native System UI Control Library) and CEF(Chromium Embedded Framework) (Windows/ macOS / Linux)   [![There was an update last month][G]](https://github.com/energye/energy)   [![godoc][D]](https://godoc.org/github.com/energye/energy)   [![Contains Chinese documents][CN]](https://github.com/energye/energy)
- [cimgui-go](https://github.com/AllenDang/cimgui-go) **star:528** Auto generated Go wrapper for [Dear ImGui](https://github.com/ocornut/imgui) via [cimgui](https://github.com/cimgui/cimgui).
- [gowd](https://github.com/dtylman/gowd) **star:435** Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.   [![It hasn't been updated in recent three years][Y]](https://github.com/dtylman/gowd)   [![godoc][D]](https://godoc.org/github.com/dtylman/gowd)
- [gogpu/ui](https://github.com/gogpu/ui) **star:387** GPU-accelerated GUI toolkit with 22 widgets, 3 design systems (Material, Fluent, Cupertino), reactive signals, and zero CGO (part of [GoGPU](https://github.com/gogpu) ecosystem).   [![There was an update last month][G]](https://github.com/gogpu/ui)   [![godoc][D]](https://godoc.org/github.com/gogpu/ui)
- [unison](https://github.com/richardwilkes/unison) **star:329** A unified graphical user experience toolkit for Go desktop applications. macOS, Windows, and Linux are supported.   [![There was an update last month][G]](https://github.com/richardwilkes/unison)   [![godoc][D]](https://godoc.org/github.com/richardwilkes/unison)
- [Wails](https://wails.io)  Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
- [proton](https://github.com/CzaxStudio/proton) **star:28** Pure Go immediate-mode GUI framework built on Gio with zero Cgo dependencies.   [![There was an update last month][G]](https://github.com/CzaxStudio/proton)   [![godoc][D]](https://godoc.org/github.com/CzaxStudio/proton)
- [goradd/html5tag](https://github.com/goradd/html5tag) **star:13** Library for outputting HTML5 tags.   [![godoc][D]](https://godoc.org/github.com/goradd/html5tag)

_Interaction_

- [robotgo](https://github.com/go-vgo/robotgo) **star:10747** Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.   [![There was an update last month][G]](https://github.com/go-vgo/robotgo)   [![godoc][D]](https://godoc.org/github.com/go-vgo/robotgo)
- [systray](https://github.com/getlantern/systray) **star:3720** Cross platform Go library to place an icon and menu in the notification area.   [![godoc][D]](https://godoc.org/github.com/getlantern/systray)
- [zenity](https://github.com/ncruces/zenity) **star:915** Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user.   [![There was an update last month][G]](https://github.com/ncruces/zenity)   [![godoc][D]](https://godoc.org/github.com/ncruces/zenity)
- [gosx-notifier](https://github.com/deckarep/gosx-notifier) **star:590** OSX Desktop Notifications library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/deckarep/gosx-notifier)   [![godoc][D]](https://godoc.org/github.com/deckarep/gosx-notifier)
- [trayhost](https://github.com/shurcooL/trayhost) **star:260** Cross-platform Go library to place an icon in the host operating system's taskbar.   [![godoc][D]](https://godoc.org/github.com/shurcooL/trayhost)
- [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) **star:41** OSX Sleep/Wake notifications in golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/prashantgupta24/mac-sleep-notifier)   [![godoc][D]](https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier)
- [gogpu/systray](https://github.com/gogpu/systray) **star:40** Pure Go system tray library for Windows, macOS, and Linux with zero CGO (part of [GoGPU](https://github.com/gogpu) ecosystem).   [![There was an update last month][G]](https://github.com/gogpu/systray)   [![godoc][D]](https://godoc.org/github.com/gogpu/systray)
- [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) **star:32** OSX library to notify about any (pluggable) activity on your machine.   [![godoc][D]](https://godoc.org/github.com/prashantgupta24/activity-tracker)
- [AppIndicator Go](https://github.com/gopherlibs/appindicator) **star:5** Go bindings for libappindicator3 C library.   [![godoc][D]](https://godoc.org/github.com/gopherlibs/appindicator)

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli](https://github.com/arduino/arduino-cli) **star:4965** Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects.   [![There was an update last month][G]](https://github.com/arduino/arduino-cli)   [![godoc][D]](https://godoc.org/github.com/arduino/arduino-cli)
- [go-rpio](https://github.com/stianeikeland/go-rpio) **star:2288** GPIO for Go, doesn't require cgo.   [![godoc][D]](https://godoc.org/github.com/stianeikeland/go-rpio)
- [ghw](https://github.com/jaypipes/ghw) **star:1872** Golang hardware discovery/inspection library.   [![There was an update last month][G]](https://github.com/jaypipes/ghw)   [![godoc][D]](https://godoc.org/github.com/jaypipes/ghw)
- [emgo](https://github.com/ziutek/emgo) **star:1097** Go-like language for programming embedded systems (e.g. STM32 MCU).   [![It hasn't been updated in recent three years][Y]](https://github.com/ziutek/emgo)
- [sysinfo](https://github.com/zcalusic/sysinfo) **star:576** A pure Go library providing Linux OS / kernel / hardware system information.   [![godoc][D]](https://godoc.org/github.com/zcalusic/sysinfo)
- [goroslib](https://github.com/aler9/goroslib) **star:366** Robot Operating System (ROS) library for Go.   [![godoc][D]](https://godoc.org/github.com/aler9/goroslib)   [![Archived][Archived]](https://github.com/aler9/goroslib)
- [go-osc](https://github.com/hypebeast/go-osc) **star:232** Open Sound Control (OSC) bindings for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/hypebeast/go-osc)   [![godoc][D]](https://godoc.org/github.com/hypebeast/go-osc)
- [joystick](https://github.com/0xcafed00d/joystick) **star:76** a polled API to read the state of an attached joystick.   [![godoc][D]](https://godoc.org/github.com/0xcafed00d/joystick)
- [moody](https://github.com/dinakars777/moody) **star:5** Hardware event personality daemon for macOS. Monitors USB, charger, lid, and other hardware events and responds with customizable personalities.   [![godoc][D]](https://godoc.org/github.com/dinakars777/moody)

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [gocv](https://github.com/hybridgroup/gocv) **star:7474** Go package for computer vision using OpenCV 3.3+.   [![godoc][D]](https://godoc.org/github.com/hybridgroup/gocv)
- [imaginary](https://github.com/h2non/imaginary) **star:6070** Fast and simple HTTP microservice for image resizing.   [![godoc][D]](https://godoc.org/github.com/h2non/imaginary)
- [imaging](https://github.com/disintegration/imaging) **star:5738** Simple Go image processing package.   [![godoc][D]](https://godoc.org/github.com/disintegration/imaging)
- [gg](https://github.com/fogleman/gg) **star:4786** 2D rendering in pure Go.   [![godoc][D]](https://godoc.org/github.com/fogleman/gg)
- [gowitness](https://github.com/sensepost/gowitness) **star:4367** Screenshoting webpages using go and headless chrome on command line.   [![godoc][D]](https://godoc.org/github.com/sensepost/gowitness)
- [bild](https://github.com/anthonynsimon/bild) **star:4201** Collection of image processing algorithms in pure Go.   [![There was an update last month][G]](https://github.com/anthonynsimon/bild)   [![godoc][D]](https://godoc.org/github.com/anthonynsimon/bild)
- [imagor](https://github.com/cshum/imagor) **star:3983** Fast, secure image processing server and Go library, using libvips.   [![godoc][D]](https://godoc.org/github.com/cshum/imagor)
- [ln](https://github.com/fogleman/ln) **star:3375** 3D line art rendering in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/fogleman/ln)   [![godoc][D]](https://godoc.org/github.com/fogleman/ln)
- [bimg](https://github.com/h2non/bimg) **star:3019** Small package for fast and efficient image processing using libvips.   [![godoc][D]](https://godoc.org/github.com/h2non/bimg)
- [picfit](https://github.com/thoas/picfit) **star:2338** An image resizing server written in Go.   [![There was an update last month][G]](https://github.com/thoas/picfit)   [![godoc][D]](https://godoc.org/github.com/thoas/picfit)
- [svgo](https://github.com/ajstarks/svgo) **star:2248** Go Language Library for SVG generation.   [![It hasn't been updated in recent three years][Y]](https://github.com/ajstarks/svgo)   [![godoc][D]](https://godoc.org/github.com/ajstarks/svgo)
- [pt](https://github.com/fogleman/pt) **star:2103** Path tracing engine written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/fogleman/pt)   [![godoc][D]](https://godoc.org/github.com/fogleman/pt)
- [imagick](https://github.com/gographics/imagick) **star:1871** Go binding to ImageMagick's MagickWand C API.   [![godoc][D]](https://godoc.org/github.com/gographics/imagick)
- [smartcrop](https://github.com/muesli/smartcrop) **star:1854** Finds good crops for arbitrary images and crop sizes.   [![It hasn't been updated in recent three years][Y]](https://github.com/muesli/smartcrop)   [![godoc][D]](https://godoc.org/github.com/muesli/smartcrop)
- [canvas](https://github.com/tdewolff/canvas) **star:1825** Vector graphics to PDF, SVG or rasterized image.   [![There was an update last month][G]](https://github.com/tdewolff/canvas)   [![godoc][D]](https://godoc.org/github.com/tdewolff/canvas)
- [gift](https://github.com/disintegration/gift) **star:1796** Package of image processing filters.   [![godoc][D]](https://godoc.org/github.com/disintegration/gift)
- [govips](https://github.com/davidbyttow/govips) **star:1633** A lightning fast image processing and resizing library for Go.   [![godoc][D]](https://godoc.org/github.com/davidbyttow/govips)
- [geopattern](https://github.com/pravj/geopattern) **star:1285** Create beautiful generative image patterns from a string.   [![It hasn't been updated in recent three years][Y]](https://github.com/pravj/geopattern)   [![godoc][D]](https://godoc.org/github.com/pravj/geopattern)
- [stegify](https://github.com/DimitarPetrov/stegify) **star:1265** Go tool for LSB steganography, capable of hiding any file within an image.   [![It hasn't been updated in recent three years][Y]](https://github.com/DimitarPetrov/stegify)   [![godoc][D]](https://godoc.org/github.com/DimitarPetrov/stegify)
- [image2ascii](https://github.com/qeesung/image2ascii) **star:980** Convert image to ASCII.   [![It hasn't been updated in recent three years][Y]](https://github.com/qeesung/image2ascii)   [![godoc][D]](https://godoc.org/github.com/qeesung/image2ascii)
- [go-qrcode](https://github.com/yeqown/go-qrcode) **star:855** Generate QR codes with personalized styles, allowing adjustments to color, block size, shape, and icons.   [![godoc][D]](https://godoc.org/github.com/yeqown/go-qrcode)
- [goimagehash](https://github.com/corona10/goimagehash) **star:837** Go Perceptual image hashing package.   [![godoc][D]](https://godoc.org/github.com/corona10/goimagehash)
- [govatar](https://github.com/o1egl/govatar) **star:609** Library and CMD tool for generating funny avatars.   [![It hasn't been updated in recent three years][Y]](https://github.com/o1egl/govatar)   [![godoc][D]](https://godoc.org/github.com/o1egl/govatar)
- [mort](https://github.com/aldor007/mort) **star:522** Storage and image processing server written in Go.   [![godoc][D]](https://godoc.org/github.com/aldor007/mort)
- [nativewebp](https://github.com/HugoSmits86/nativewebp) **star:450** Go native WebP encoder with zero external dependencies.   [![godoc][D]](https://godoc.org/github.com/HugoSmits86/nativewebp)
- [go-nude](https://github.com/koyachi/go-nude) **star:422** Nudity detection with Go.   [![godoc][D]](https://godoc.org/github.com/koyachi/go-nude)
- [steganography](https://github.com/auyer/steganography) **star:356** Pure Go Library for LSB steganography.   [![godoc][D]](https://godoc.org/github.com/auyer/steganography)
- [go-webp](https://github.com/kolesa-team/go-webp) **star:315** Library for encode and decode webp pictures, using libwebp.   [![godoc][D]](https://godoc.org/github.com/kolesa-team/go-webp)
- [transformimgs](https://github.com/Pixboost/transformimgs) **star:292** Transformimgs resizes and optimises images for Web using next-generation formats.   [![godoc][D]](https://godoc.org/github.com/Pixboost/transformimgs)
- [gltf](https://github.com/qmuntal/gltf) **star:286** Efficient and robust glTF 2.0 reader, writer and validator.   [![godoc][D]](https://godoc.org/github.com/qmuntal/gltf)
- [mergi](https://github.com/noelyahan/mergi) **star:243** Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).   [![godoc][D]](https://godoc.org/github.com/noelyahan/mergi)
- [darkroom](https://github.com/gojek/darkroom) **star:235** An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.   [![godoc][D]](https://godoc.org/github.com/gojek/darkroom)
- [img](https://github.com/hawx/img) **star:158** Selection of image manipulation tools.   [![It hasn't been updated in recent three years][Y]](https://github.com/hawx/img)   [![godoc][D]](https://godoc.org/github.com/hawx/img)
- [go-cairo](https://github.com/ungerik/go-cairo) **star:153** Go binding for the cairo graphics library.   [![godoc][D]](https://godoc.org/github.com/ungerik/go-cairo)
- [gogpu/gg](https://github.com/gogpu/gg) **star:137** GPU-accelerated 2D rendering with Canvas-like API, zero CGO (part of [GoGPU](https://github.com/gogpu) pure Go graphics ecosystem).   [![There was an update last month][G]](https://github.com/gogpu/gg)   [![godoc][D]](https://godoc.org/github.com/gogpu/gg)
- [cameron](https://github.com/aofei/cameron) **star:132** An avatar generator for Go.   [![godoc][D]](https://godoc.org/github.com/aofei/cameron)
- [color-extractor](https://github.com/marekm4/color-extractor) **star:131** Dominant color extractor with no external dependencies.   [![It hasn't been updated in recent three years][Y]](https://github.com/marekm4/color-extractor)   [![godoc][D]](https://godoc.org/github.com/marekm4/color-extractor)
- [webp-server](https://github.com/mehdipourfar/webp-server) **star:82** Simple and minimal image server capable of storing, resizing, converting and caching images.   [![It hasn't been updated in recent three years][Y]](https://github.com/mehdipourfar/webp-server)   [![godoc][D]](https://godoc.org/github.com/mehdipourfar/webp-server)
- [gridder](https://github.com/shomali11/gridder) **star:81** A Grid based 2D Graphics library.   [![It hasn't been updated in recent three years][Y]](https://github.com/shomali11/gridder)   [![godoc][D]](https://godoc.org/github.com/shomali11/gridder)
- [go-gd](https://github.com/bolknote/go-gd) **star:60** Go binding for GD library.   [![godoc][D]](https://godoc.org/github.com/bolknote/go-gd)
- [goimghdr](https://github.com/corona10/goimghdr) **star:40** The imghdr module determines the type of image contained in a file for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/corona10/goimghdr)   [![godoc][D]](https://godoc.org/github.com/corona10/goimghdr)
- [scout](https://github.com/jonoton/scout) **star:30** Scout is a standalone open source software solution for DIY video security.   [![godoc][D]](https://godoc.org/github.com/jonoton/scout)
- [go-webcolors](https://github.com/jyotiska/go-webcolors) **star:29** Port of webcolors library from Python to Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/jyotiska/go-webcolors)   [![godoc][D]](https://godoc.org/github.com/jyotiska/go-webcolors)
- [mpo](https://github.com/donatj/mpo) **star:27** Decoder and conversion tool for MPO 3D Photos.   [![There was an update last month][G]](https://github.com/donatj/mpo)   [![godoc][D]](https://godoc.org/github.com/donatj/mpo)

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [flogo](https://github.com/tibcosoftware/flogo) **star:2493** Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
- [ekuiper](https://github.com/lf-edge/ekuiper) **star:1725** Lightweight data stream processing engine for IoT edge.   [![There was an update last month][G]](https://github.com/lf-edge/ekuiper)   [![godoc][D]](https://godoc.org/github.com/lf-edge/ekuiper)   [![Contains Chinese documents][CN]](https://github.com/lf-edge/ekuiper)
- [rulego](https://github.com/rulego/rulego) **star:1563** RuleGo is a lightweight, high-performance, embedded, orchestrable component-based rule engine for IoT edge.   [![There was an update last month][G]](https://github.com/rulego/rulego)   [![godoc][D]](https://godoc.org/github.com/rulego/rulego)   [![Contains Chinese documents][CN]](https://github.com/rulego/rulego)
- [shifu](https://github.com/Edgenesis/shifu) **star:1426** Kubernetes native IoT development framework.   [![There was an update last month][G]](https://github.com/Edgenesis/shifu)   [![godoc][D]](https://godoc.org/github.com/Edgenesis/shifu)   [![Contains Chinese documents][CN]](https://github.com/Edgenesis/shifu)
- [gatt](https://github.com/paypal/gatt) **star:1164** Gatt is a Go package for building Bluetooth Low Energy peripherals.   [![It hasn't been updated in recent three years][Y]](https://github.com/paypal/gatt)   [![godoc][D]](https://godoc.org/github.com/paypal/gatt)
- [gobot](https://github.com/hybridgroup/gobot/)  Gobot is a framework for robotics, physical computing, and the Internet of Things.
- [connectordb](https://github.com/connectordb/connectordb) **star:421** Open-Source Platform for Quantified Self & IoT.   [![It hasn't been updated in recent three years][Y]](https://github.com/connectordb/connectordb)   [![godoc][D]](https://godoc.org/github.com/connectordb/connectordb)
- [devices](https://github.com/goiot/devices) **star:266** Suite of libraries for IoT devices, experimental for x/exp/io.   [![It hasn't been updated in recent three years][Y]](https://github.com/goiot/devices)   [![godoc][D]](https://godoc.org/github.com/goiot/devices)
- [huego](https://github.com/amimof/huego) **star:262** An extensive Philips Hue client library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/amimof/huego)   [![godoc][D]](https://godoc.org/github.com/amimof/huego)
- [iot](https://github.com/vaelen/iot/)  IoT is a simple framework for implementing a Google IoT Core device.
- [periph](https://periph.io/)  Peripherals I/O to interface with low-level board facilities.
- [sensorbee](https://github.com/sensorbee/sensorbee) **star:231** Lightweight stream processing engine for IoT.   [![It hasn't been updated in recent three years][Y]](https://github.com/sensorbee/sensorbee)   [![godoc][D]](https://godoc.org/github.com/sensorbee/sensorbee)
- [smart-home](https://github.com/e154/smart-home) **star:102** Software package for IoT automation.   [![godoc][D]](https://godoc.org/github.com/e154/smart-home)
- [eywa](https://github.com/xcodersun/eywa) **star:66** Project Eywa is essentially a connection manager that keeps track of connected devices.   [![It hasn't been updated in recent three years][Y]](https://github.com/xcodersun/eywa)   [![godoc][D]](https://godoc.org/github.com/xcodersun/eywa)

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._

- [gocron](https://github.com/go-co-op/gocron) **star:7099** Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron).   [![There was an update last month][G]](https://github.com/go-co-op/gocron)   [![godoc][D]](https://godoc.org/github.com/go-co-op/gocron)
- [go-quartz](https://github.com/reugn/go-quartz) **star:2014** Simple, zero-dependency scheduling library for Go.   [![godoc][D]](https://godoc.org/github.com/reugn/go-quartz)
- [JobRunner](https://github.com/bamzi/jobrunner) **star:1090** Smart and featureful cron job scheduler with job queuing and live monitoring built in.   [![It hasn't been updated in recent three years][Y]](https://github.com/bamzi/jobrunner)   [![godoc][D]](https://godoc.org/github.com/bamzi/jobrunner)
- [gron](https://github.com/roylee0704/gron) **star:1034** Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.   [![It hasn't been updated in recent three years][Y]](https://github.com/roylee0704/gron)   [![godoc][D]](https://godoc.org/github.com/roylee0704/gron)
- [gronx](https://github.com/adhocore/gronx) **star:512** Cron expression parser, task runner and daemon consuming crontab like task list.   [![godoc][D]](https://godoc.org/github.com/adhocore/gronx)
- [goflow](https://github.com/fieldryand/goflow) **star:481** A simple but powerful DAG scheduler and dashboard.   [![godoc][D]](https://godoc.org/github.com/fieldryand/goflow)
- [scheduler](https://github.com/carlescere/scheduler) **star:469** Cronjobs scheduling made easy.   [![It hasn't been updated in recent three years][Y]](https://github.com/carlescere/scheduler)   [![godoc][D]](https://godoc.org/github.com/carlescere/scheduler)
- [tasks](https://github.com/madflojo/tasks) **star:330** An easy to use in-process scheduler for recurring tasks in Go.   [![godoc][D]](https://godoc.org/github.com/madflojo/tasks)
- [go-cron](https://github.com/rk/go-cron) **star:239** Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.   [![It hasn't been updated in recent three years][Y]](https://github.com/rk/go-cron)   [![godoc][D]](https://godoc.org/github.com/rk/go-cron)
- [cheek](https://github.com/bart6114/cheek) **star:200** A simple crontab like scheduler that aims to offer a KISS approach to job scheduling.   [![godoc][D]](https://godoc.org/github.com/bart6114/cheek)
- [clockwerk](https://github.com/onatm/clockwerk) **star:182** Go package to schedule periodic jobs using a simple, fluent syntax.   [![godoc][D]](https://godoc.org/github.com/onatm/clockwerk)
- [leprechaun](https://github.com/kilgaloon/leprechaun) **star:106** Job scheduler that supports webhooks, crons and classic scheduling.   [![It hasn't been updated in recent three years][Y]](https://github.com/kilgaloon/leprechaun)   [![godoc][D]](https://godoc.org/github.com/kilgaloon/leprechaun)
- [ofelia](https://github.com/netresearch/ofelia) **star:63** Docker job scheduler (crontab for Docker); fork of mcuadros/ofelia that adds a web UI, job dependencies, retries, and job persistence.   [![There was an update last month][G]](https://github.com/netresearch/ofelia)   [![godoc][D]](https://godoc.org/github.com/netresearch/ofelia)
- [cdule](https://github.com/deepaksinghvi/cdule) **star:61** Job scheduler library with database support   [![godoc][D]](https://godoc.org/github.com/deepaksinghvi/cdule)
- [go-cron](https://github.com/netresearch/go-cron) **star:48** Cron job scheduler with runtime schedule updates, per-entry context, resilience middleware (retry, circuit breaker, rate limiting), and observability hooks; successor to robfig/cron.   [![godoc][D]](https://godoc.org/github.com/netresearch/go-cron)
- [go-scheduler](https://github.com/pardnchiu/go-scheduler) **star:36** Job scheduler supporting standard cron expressions, custom descriptors, intervals, and task dependencies.   [![godoc][D]](https://godoc.org/github.com/pardnchiu/go-scheduler)
- [sched](https://github.com/romshark/sched) **star:31** A job scheduler with the ability to fast-forward time.   [![godoc][D]](https://godoc.org/github.com/romshark/sched)
- [cronticker](https://github.com/krayzpipes/cronticker) **star:20** A ticker implementation to support cron schedules.   [![It hasn't been updated in recent three years][Y]](https://github.com/krayzpipes/cronticker)   [![godoc][D]](https://godoc.org/github.com/krayzpipes/cronticker)
- [pending](https://github.com/kahoon/pending) **star:7** ID-based debounced task scheduler for deferred tasks with cancellation, graceful shutdown, and optional concurrency limits.   [![godoc][D]](https://godoc.org/github.com/kahoon/pending)
- [scheduler](https://github.com/yuseferi/scheduler) **star:6** Go-native distributed job scheduler with delayed tasks, batched Redis coordination, retries, lease-based recovery, and versioned queue partitioning.   [![godoc][D]](https://godoc.org/github.com/yuseferi/scheduler)
- [go-job](https://github.com/cybergarage/go-job) **star:2** A flexible and extensible job scheduling and execution library for Go.   [![godoc][D]](https://godoc.org/github.com/cybergarage/go-job)
- [tickstem/cron](https://github.com/tickstem/cron) **star:1** Go client for scheduling HTTP cron jobs, with execution history, failure alerts, and tsk-local for testing handlers without live credentials.   [![godoc][D]](https://godoc.org/github.com/tickstem/cron)
- [tickstem/heartbeat](https://github.com/tickstem/heartbeat) **star:1** Go client for dead-man's switch heartbeat monitoring: ping a URL after each job run and get alerted by email if pings stop arriving.   [![godoc][D]](https://godoc.org/github.com/tickstem/heartbeat)

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [GJSON](https://github.com/tidwall/gjson) **star:15543** Get a JSON value with one line of code.   [![godoc][D]](https://godoc.org/github.com/tidwall/gjson)
- [gabs](https://github.com/Jeffail/gabs) **star:3531** For parsing, creating and editing unknown or dynamic JSON in Go.   [![godoc][D]](https://godoc.org/github.com/Jeffail/gabs)
- [SJSON](https://github.com/tidwall/sjson) **star:2716** Set a JSON value with one line of code.     [![godoc][D]](https://godoc.org/github.com/tidwall/sjson)
- [gojson](https://github.com/ChimeraCoder/gojson) **star:2688** Automatically generate Go (golang) struct definitions from example JSON.   [![It hasn't been updated in recent three years][Y]](https://github.com/ChimeraCoder/gojson)   [![godoc][D]](https://godoc.org/github.com/ChimeraCoder/gojson)
- [fastjson](https://github.com/valyala/fastjson) **star:2464** Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection.   [![godoc][D]](https://godoc.org/github.com/valyala/fastjson)
- [OjG](https://github.com/ohler55/ojg) **star:952** Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath.   [![godoc][D]](https://godoc.org/github.com/ohler55/ojg)
- [jsondiff](https://github.com/wI2L/jsondiff) **star:629** JSON diff library for Go based on RFC6902 (JSON Patch).   [![godoc][D]](https://godoc.org/github.com/wI2L/jsondiff)
- [marshmallow](https://github.com/PerimeterX/marshmallow) **star:391** Performant JSON unmarshalling for flexible use cases.   [![It hasn't been updated in recent three years][Y]](https://github.com/PerimeterX/marshmallow)   [![godoc][D]](https://godoc.org/github.com/PerimeterX/marshmallow)
- [kazaam](https://github.com/Qntfy/kazaam) **star:293** API for arbitrary transformation of JSON documents.   [![It hasn't been updated in recent three years][Y]](https://github.com/Qntfy/kazaam)   [![godoc][D]](https://godoc.org/github.com/Qntfy/kazaam)
- [ajson](https://github.com/spyzhov/ajson) **star:290** Abstract JSON for golang with JSONPath support.   [![godoc][D]](https://godoc.org/github.com/spyzhov/ajson)
- [jsonvalue](https://github.com/Andrew-M-C/go.jsonvalue) **star:203** A fast and convenient library for unstructured JSON data, replacing `encoding/json`.   [![godoc][D]](https://godoc.org/github.com/Andrew-M-C/go.jsonvalue)   [![Contains Chinese documents][CN]](https://github.com/Andrew-M-C/go.jsonvalue)
- [gojq](https://github.com/elgs/gojq) **star:191** JSON query in Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/elgs/gojq)   [![godoc][D]](https://godoc.org/github.com/elgs/gojq)
- [jettison](https://github.com/wI2L/jettison) **star:180** Fast and flexible JSON encoder for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/wI2L/jettison)   [![godoc][D]](https://godoc.org/github.com/wI2L/jettison)
- [json2go](https://github.com/m-zajac/json2go) **star:142** Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.   [![godoc][D]](https://godoc.org/github.com/m-zajac/json2go)
- [gjo](https://github.com/skanehira/gjo) **star:132** Small utility to create JSON objects.   [![It hasn't been updated in recent three years][Y]](https://github.com/skanehira/gjo)   [![godoc][D]](https://godoc.org/github.com/skanehira/gjo)
- [JayDiff](https://github.com/yazgazan/jaydiff) **star:111** JSON diff utility written in Go.   [![godoc][D]](https://godoc.org/github.com/yazgazan/jaydiff)
- [jsongo](https://github.com/ricardolonga/jsongo) **star:108** Fluent API to make it easier to create Json objects.   [![It hasn't been updated in recent three years][Y]](https://github.com/ricardolonga/jsongo)   [![godoc][D]](https://godoc.org/github.com/ricardolonga/jsongo)
- [jscan](https://github.com/romshark/jscan) **star:101** High performance zero-allocation JSON iterator.   [![godoc][D]](https://godoc.org/github.com/romshark/jscan)
- [JSON-to-Go](https://mholt.github.io/json-to-go/)  Convert JSON to Go struct.
- [JSON-to-Proto](https://json-to-proto.github.io/)  Convert JSON to Protobuf online.
- [ujson](https://github.com/olvrng/ujson) **star:85** Fast and minimal JSON parser and transformer that works on unstructured JSON.   [![godoc][D]](https://godoc.org/github.com/olvrng/ujson)
- [jsonf](https://github.com/miolini/jsonf) **star:64** Console tool for highlighted formatting and struct query fetching JSON.   [![It hasn't been updated in recent three years][Y]](https://github.com/miolini/jsonf)   [![godoc][D]](https://godoc.org/github.com/miolini/jsonf)
- [silentjson](https://github.com/GenshIv/silentjson) **star:60** Zero-allocation JSON boundary scanner and splitter utilizing AVX2 SIMD instructions.   [![There was an update last month][G]](https://github.com/GenshIv/silentjson)   [![godoc][D]](https://godoc.org/github.com/GenshIv/silentjson)   [![Contains Chinese documents][CN]](https://github.com/GenshIv/silentjson)
- [ask](https://github.com/simonnilsson/ask) **star:57** Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types.   [![godoc][D]](https://godoc.org/github.com/simonnilsson/ask)
- [go-respond](https://github.com/nicklaw5/go-respond) **star:55** Go package for handling common HTTP JSON responses.   [![It hasn't been updated in recent three years][Y]](https://github.com/nicklaw5/go-respond)   [![godoc][D]](https://godoc.org/github.com/nicklaw5/go-respond)
- [jsoncolor](https://github.com/neilotoole/jsoncolor) **star:54** Drop-in replacement for `encoding/json` that outputs colorized JSON.   [![There was an update last month][G]](https://github.com/neilotoole/jsoncolor)   [![godoc][D]](https://godoc.org/github.com/neilotoole/jsoncolor)
- [mp](https://github.com/sanbornm/mp) **star:47** Simple cli email parser. It currently takes stdin and outputs JSON.   [![It hasn't been updated in recent three years][Y]](https://github.com/sanbornm/mp)   [![godoc][D]](https://godoc.org/github.com/sanbornm/mp)
- [vjson](https://github.com/miladibra10/vjson) **star:41** Go package for validating JSON objects with declaring a JSON schema with fluent API.   [![godoc][D]](https://godoc.org/github.com/miladibra10/vjson)
- [gojmapr](https://github.com/limiu82214/gojmapr) **star:22** Get simple struct from complex json by json path.   [![It hasn't been updated in recent three years][Y]](https://github.com/limiu82214/gojmapr)   [![godoc][D]](https://godoc.org/github.com/limiu82214/gojmapr)   [![Contains Chinese documents][CN]](https://github.com/limiu82214/gojmapr)
- [mapslice-json](https://github.com/mickep76/mapslice-json) **star:20** Go MapSlice for ordered marshal/ unmarshal of maps in JSON.   [![godoc][D]](https://godoc.org/github.com/mickep76/mapslice-json)
- [epoch](https://github.com/vtopc/epoch) **star:17** Contains primitives for marshaling/unmarshalling Unix timestamp/epoch to/from build-in time.Time type in JSON.   [![godoc][D]](https://godoc.org/github.com/vtopc/epoch)
- [dynjson](https://github.com/cocoonspace/dynjson) **star:16** Client-customizable JSON formats for dynamic APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/cocoonspace/dynjson)   [![godoc][D]](https://godoc.org/github.com/cocoonspace/dynjson)
- [go-jsonerror](https://github.com/ddymko/go-jsonerror) **star:16** Go-JsonError is meant to allow us to easily create json response errors that follow the JsonApi spec.   [![It hasn't been updated in recent three years][Y]](https://github.com/ddymko/go-jsonerror)   [![godoc][D]](https://godoc.org/github.com/ddymko/go-jsonerror)
- [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) **star:13** Go bindings based on the JSON API errors reference.   [![It hasn't been updated in recent three years][Y]](https://github.com/AmuzaTkts/jsonapi-errors)   [![godoc][D]](https://godoc.org/github.com/AmuzaTkts/jsonapi-errors)
- [jsonhal](https://github.com/RichardKnop/jsonhal) **star:13** Simple Go package to make custom structs marshal into HAL compatible JSON responses.   [![It hasn't been updated in recent three years][Y]](https://github.com/RichardKnop/jsonhal)   [![godoc][D]](https://godoc.org/github.com/RichardKnop/jsonhal)
- [jsonic](https://github.com/sinhashubham95/jsonic) **star:11** Utilities to handle and query JSON without defining structs in a type safe manner.   [![It hasn't been updated in recent three years][Y]](https://github.com/sinhashubham95/jsonic)   [![godoc][D]](https://godoc.org/github.com/sinhashubham95/jsonic)
- [jzon](https://github.com/zerosnake0/jzon) **star:11** JSON library with standard compatible API/behavior.   [![godoc][D]](https://godoc.org/github.com/zerosnake0/jzon)
- [ej](https://github.com/lucassscaravelli/ej) **star:10** Write and read JSON from different sources succinctly.   [![It hasn't been updated in recent three years][Y]](https://github.com/lucassscaravelli/ej)   [![godoc][D]](https://godoc.org/github.com/lucassscaravelli/ej)
- [htmljson](https://github.com/nikolaydubina/htmljson) **star:10** Rich rendering of JSON as HTML in Go.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/htmljson)
- [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) **star:6** Simple JSON parser with validation by condition via golang struct fields tags.   [![It hasn't been updated in recent three years][Y]](https://github.com/dedalqq/omg.jsonparser)   [![godoc][D]](https://godoc.org/github.com/dedalqq/omg.jsonparser)
- [jsonhandlers](https://github.com/abusomani/jsonhandlers) **star:2** JSON library to expose simple handlers that lets you easily read and write json from various sources.   [![It hasn't been updated in recent three years][Y]](https://github.com/abusomani/jsonhandlers)   [![godoc][D]](https://godoc.org/github.com/abusomani/jsonhandlers)

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [logrus](https://github.com/Sirupsen/logrus) **star:25753** Structured logger for Go.   [![godoc][D]](https://godoc.org/github.com/Sirupsen/logrus)
- [zap](https://github.com/uber-go/zap) **star:24575** Fast, structured, leveled logging in Go.   [![godoc][D]](https://godoc.org/github.com/uber-go/zap)
- [zerolog](https://github.com/rs/zerolog) **star:12464** Zero-allocation JSON logger.   [![godoc][D]](https://godoc.org/github.com/rs/zerolog)
- [spew](https://github.com/davecgh/go-spew) **star:6399** Implements a deep pretty printer for Go data structures to aid in debugging.   [![godoc][D]](https://godoc.org/github.com/davecgh/go-spew)
- [lumberjack](https://github.com/natefinch/lumberjack) **star:5452** Simple rolling logger, implements io.WriteCloser.   [![godoc][D]](https://godoc.org/github.com/natefinch/lumberjack)
- [glog](https://github.com/golang/glog) **star:3602** Leveled execution logs for Go.   [![godoc][D]](https://godoc.org/github.com/golang/glog)
- [tail](https://github.com/hpcloud/tail) **star:2783** Go package striving to emulate the features of the BSD tail program.   [![It hasn't been updated in recent three years][Y]](https://github.com/hpcloud/tail)   [![godoc][D]](https://godoc.org/github.com/hpcloud/tail)
- [pp](https://github.com/k0kubun/pp) **star:2049** Colored pretty printer for Go language.   [![godoc][D]](https://godoc.org/github.com/k0kubun/pp)
- [seelog](https://github.com/cihub/seelog) **star:1636** Logging functionality with flexible dispatching, filtering, and formatting.   [![It hasn't been updated in recent three years][Y]](https://github.com/cihub/seelog)   [![godoc][D]](https://godoc.org/github.com/cihub/seelog)
- [log](https://github.com/apex/log) **star:1376** Structured logging package for Go.   [![godoc][D]](https://godoc.org/github.com/apex/log)
- [lazyjournal](https://github.com/Lifailon/lazyjournal) **star:1321** A TUI for reading and filtering logs from journalctl, file system, Docker and Podman containers, as well Kubernetes pods.   [![godoc][D]](https://godoc.org/github.com/Lifailon/lazyjournal)
- [tint](https://github.com/lmittmann/tint) **star:1318** A slog.Handler that writes tinted logs.   [![There was an update last month][G]](https://github.com/lmittmann/tint)   [![godoc][D]](https://godoc.org/github.com/lmittmann/tint)
- [log15](https://github.com/inconshreveable/log15) **star:1104** Simple, powerful logging for Go.   [![godoc][D]](https://godoc.org/github.com/inconshreveable/log15)
- [sentry-go](https://github.com/getsentry/sentry-go) **star:1095** Sentry SDK for Go. Helps monitor and track errors with real-time alerts and performance monitoring.   [![There was an update last month][G]](https://github.com/getsentry/sentry-go)   [![godoc][D]](https://godoc.org/github.com/getsentry/sentry-go)
- [phuslu/log](https://github.com/phuslu/log) **star:870** High performance structured logging.   [![There was an update last month][G]](https://github.com/phuslu/log)   [![godoc][D]](https://godoc.org/github.com/phuslu/log)
- [slog-multi](https://github.com/samber/slog-multi) **star:629** Chain of slog.Handler (pipeline, fanout...).   [![godoc][D]](https://godoc.org/github.com/samber/slog-multi)
- [slogor](https://gitlab.com/greyxor/slogor)  A colorful slog handler.
- [slog](https://github.com/gookit/slog) **star:547** Lightweight, configurable, extensible logger for Go.   [![godoc][D]](https://godoc.org/github.com/gookit/slog)   [![Contains Chinese documents][CN]](https://github.com/gookit/slog)
- [httpretty](https://github.com/henvic/httpretty) **star:414** Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest).   [![There was an update last month][G]](https://github.com/henvic/httpretty)   [![godoc][D]](https://godoc.org/github.com/henvic/httpretty)
- [onelog](https://github.com/francoispqt/onelog) **star:414** Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation.   [![It hasn't been updated in recent three years][Y]](https://github.com/francoispqt/onelog)   [![godoc][D]](https://godoc.org/github.com/francoispqt/onelog)
- [sqldb-logger](https://github.com/simukti/sqldb-logger) **star:382** A logger for Go SQL database driver without modify existing \*sql.DB stdlib usage.   [![godoc][D]](https://godoc.org/github.com/simukti/sqldb-logger)
- [logutils](https://github.com/hashicorp/logutils) **star:372** Utilities for slightly better logging in Go (Golang) extending the standard logger.   [![godoc][D]](https://godoc.org/github.com/hashicorp/logutils)
- [logxi](https://github.com/mgutz/logxi) **star:356** 12-factor app logger that is fast and makes you happy.   [![It hasn't been updated in recent three years][Y]](https://github.com/mgutz/logxi)   [![godoc][D]](https://godoc.org/github.com/mgutz/logxi)
- [log](https://github.com/go-playground/log) **star:293** Simple, configurable and scalable Structured Logging for Go.   [![godoc][D]](https://godoc.org/github.com/go-playground/log)
- [rollingwriter](https://github.com/arthurkiller/rollingWriter) **star:292** RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.   [![godoc][D]](https://godoc.org/github.com/arthurkiller/rollingWriter)
- [go-logger](https://github.com/apsdehal/go-logger) **star:288** Simple logger of Go Programs, with level handlers.   [![It hasn't been updated in recent three years][Y]](https://github.com/apsdehal/go-logger)   [![godoc][D]](https://godoc.org/github.com/apsdehal/go-logger)
- [slog-formatter](https://github.com/samber/slog-formatter) **star:220** Common formatters for slog and helpers to build your own.   [![godoc][D]](https://godoc.org/github.com/samber/slog-formatter)
- [glg](https://github.com/kpango/glg) **star:193** glg is simple and fast leveled logging library for Go.   [![godoc][D]](https://godoc.org/github.com/kpango/glg)
- [logger](https://github.com/azer/logger) **star:159** Minimalistic logging library for Go.   [![godoc][D]](https://godoc.org/github.com/azer/logger)
- [timberjack](https://github.com/DeRuina/timberjack) **star:150** Rolling logger with size-based, time-based, and scheduled clock-based rotation, supporting compression and cleanup.   [![godoc][D]](https://godoc.org/github.com/DeRuina/timberjack)
- [xlog](https://github.com/rs/xlog) **star:141** Structured logger for `net/context` aware HTTP handlers with flexible dispatching.   [![godoc][D]](https://godoc.org/github.com/rs/xlog)
- [ozzo-log](https://github.com/go-ozzo/ozzo-log) **star:123** High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).   [![It hasn't been updated in recent three years][Y]](https://github.com/go-ozzo/ozzo-log)   [![godoc][D]](https://godoc.org/github.com/go-ozzo/ozzo-log)   [![Contains Chinese documents][CN]](https://github.com/go-ozzo/ozzo-log)
- [caarlos0/log](https://github.com/caarlos0/log) **star:64** Colorful CLI logger.   [![godoc][D]](https://godoc.org/github.com/caarlos0/log)
- [go-cronowriter](https://github.com/utahta/go-cronowriter) **star:56** Simple writer that rotate log files automatically based on current date and time, like cronolog.   [![It hasn't been updated in recent three years][Y]](https://github.com/utahta/go-cronowriter)   [![godoc][D]](https://godoc.org/github.com/utahta/go-cronowriter)
- [stdlog](https://github.com/alexcesaro/log) **star:48** Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.   [![It hasn't been updated in recent three years][Y]](https://github.com/alexcesaro/log)   [![godoc][D]](https://godoc.org/github.com/alexcesaro/log)
- [journald](https://github.com/ssgreg/journald) **star:46** Go implementation of systemd Journal's native API for logging.   [![It hasn't been updated in recent three years][Y]](https://github.com/ssgreg/journald)   [![godoc][D]](https://godoc.org/github.com/ssgreg/journald)
- [noodlog](https://github.com/gyozatech/noodlog) **star:44** Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings.   [![It hasn't been updated in recent three years][Y]](https://github.com/gyozatech/noodlog)   [![godoc][D]](https://godoc.org/github.com/gyozatech/noodlog)
- [go-log](https://github.com/ian-kent/go-log) **star:42** Log4j implementation in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ian-kent/go-log)   [![godoc][D]](https://godoc.org/github.com/ian-kent/go-log)
- [logex](https://github.com/chzyer/logex) **star:42** Golang log lib, supports tracking and level, wrap by standard log lib.   [![godoc][D]](https://godoc.org/github.com/chzyer/logex)
- [go-log](https://github.com/siddontang/go-log) **star:38** Log lib supports level and multi handlers.   [![It hasn't been updated in recent three years][Y]](https://github.com/siddontang/go-log)   [![godoc][D]](https://godoc.org/github.com/siddontang/go-log)
- [zax](https://github.com/yuseferi/zax) **star:37** Integrate Context with Zap logger, which leads to more flexibility in Go logging.   [![godoc][D]](https://godoc.org/github.com/yuseferi/zax)
- [mlog](https://github.com/jbrodriguez/mlog) **star:33** Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.   [![It hasn't been updated in recent three years][Y]](https://github.com/jbrodriguez/mlog)   [![godoc][D]](https://godoc.org/github.com/jbrodriguez/mlog)
- [distillog](https://github.com/amoghe/distillog) **star:31** distilled levelled logging (think of it as stdlib + log levels).   [![It hasn't been updated in recent three years][Y]](https://github.com/amoghe/distillog)   [![godoc][D]](https://godoc.org/github.com/amoghe/distillog)
- [logrusly](https://github.com/sebest/logrusly) **star:29** [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).   [![It hasn't been updated in recent three years][Y]](https://github.com/sebest/logrusly)   [![godoc][D]](https://godoc.org/github.com/sebest/logrusly)
- [zkits-logger](https://github.com/edoger/zkits-logger) **star:28** A powerful zero-dependency JSON logger.   [![It hasn't been updated in recent three years][Y]](https://github.com/edoger/zkits-logger)   [![godoc][D]](https://godoc.org/github.com/edoger/zkits-logger)
- [log](https://github.com/teris-io/log) **star:25** Structured log interface for Go cleanly separates logging facade from its implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/teris-io/log)   [![godoc][D]](https://godoc.org/github.com/teris-io/log)
- [kemba](https://github.com/clok/kemba) **star:18** A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications.   [![godoc][D]](https://godoc.org/github.com/clok/kemba)
- [log](https://github.com/heartwilltell/log) **star:17** Simple leveled logging wrapper around standard log package.   [![It hasn't been updated in recent three years][Y]](https://github.com/heartwilltell/log)   [![godoc][D]](https://godoc.org/github.com/heartwilltell/log)
- [xylog](https://github.com/xybor-x/xylog) **star:17** Leveled and structured logging, dynamic fields, high performance, zone management, simple configuration, and readable syntax.   [![It hasn't been updated in recent three years][Y]](https://github.com/xybor-x/xylog)   [![godoc][D]](https://godoc.org/github.com/xybor-x/xylog)
- [glo](https://github.com/lajosbencz/glo) **star:16** PHP Monolog inspired logging facility with identical severity levels.   [![It hasn't been updated in recent three years][Y]](https://github.com/lajosbencz/glo)   [![godoc][D]](https://godoc.org/github.com/lajosbencz/glo)
- [logrusiowriter](https://github.com/cabify/logrusiowriter) **star:15** `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.   [![godoc][D]](https://godoc.org/github.com/cabify/logrusiowriter)
- [go-log](https://github.com/subchen/go-log) **star:14** Simple and configurable Logging in Go, with level, formatters and writers.   [![It hasn't been updated in recent three years][Y]](https://github.com/subchen/go-log)   [![godoc][D]](https://godoc.org/github.com/subchen/go-log)
- [logo](https://github.com/mbndr/logo) **star:12** Golang logger to different configurable writers.   [![It hasn't been updated in recent three years][Y]](https://github.com/mbndr/logo)   [![godoc][D]](https://godoc.org/github.com/mbndr/logo)
- [go-log](https://github.com/pieterclaerhout/go-log) **star:11** A logging library with stack traces, object dumping and optional timestamps.   [![godoc][D]](https://godoc.org/github.com/pieterclaerhout/go-log)
- [logdump](https://github.com/ewwwwwqm/logdump) **star:11** Package for multi-level logging.   [![It hasn't been updated in recent three years][Y]](https://github.com/ewwwwwqm/logdump)   [![godoc][D]](https://godoc.org/github.com/ewwwwwqm/logdump)
- [log](https://github.com/aerogo/log) **star:10** An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).   [![It hasn't been updated in recent three years][Y]](https://github.com/aerogo/log)   [![godoc][D]](https://godoc.org/github.com/aerogo/log)
- [zl](https://github.com/nkmr-jp/zl) **star:10** High Developer Experience, zap based logger. It offers rich functionality but is easy to configure.   [![godoc][D]](https://godoc.org/github.com/nkmr-jp/zl)
- [xlog](https://github.com/xfxdev/xlog) **star:8** Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.   [![It hasn't been updated in recent three years][Y]](https://github.com/xfxdev/xlog)   [![godoc][D]](https://godoc.org/github.com/xfxdev/xlog)
- [slf4g](https://github.com/echocat/slf4g) **star:6** Simple Logging Facade for Golang: Simple structured logging; but powerful, extendable and customizable, with huge amount of learnings from decades of past logging frameworks.   [![There was an update last month][G]](https://github.com/echocat/slf4g)   [![godoc][D]](https://godoc.org/github.com/echocat/slf4g)
- [GoLogX](https://github.com/AyoubTadlaoui/GoLogX) **star:5** Append-only, hash-chained, optionally Ed25519-signed slog handler with offline verification of tampering.   [![godoc][D]](https://godoc.org/github.com/AyoubTadlaoui/GoLogX)
- [gone/log](https://github.com/One-com/gone/tree/master/log)  Fast, extendable, full-featured, std-lib source compatible log library.
- [structy/log](https://github.com/structy/log) **star:5** A simple to use log system, minimalist but with features for debugging and differentiation of messages.   [![It hasn't been updated in recent three years][Y]](https://github.com/structy/log)   [![godoc][D]](https://godoc.org/github.com/structy/log)
- [log](https://github.com/no-src/log) **star:4** A simple logging framework out of the box.   [![godoc][D]](https://godoc.org/github.com/no-src/log)
- [yell](https://github.com/jfcg/yell) **star:1** Yet another minimalistic logging library.   [![It hasn't been updated in recent three years][Y]](https://github.com/jfcg/yell)   [![godoc][D]](https://godoc.org/github.com/jfcg/yell)

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [GoLearn](https://github.com/sjwhitworth/golearn) **star:9442** General Machine Learning library for Go.   [![godoc][D]](https://godoc.org/github.com/sjwhitworth/golearn)   [![Contains Chinese documents][CN]](https://github.com/sjwhitworth/golearn)
- [gorgonia](https://github.com/gorgonia/gorgonia) **star:5921** graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.   [![godoc][D]](https://godoc.org/github.com/gorgonia/gorgonia)
- [gorse](https://github.com/zhenghaoz/gorse)  An offline recommender system backend based on collaborative filtering written in Go.
- [gosseract](https://github.com/otiai10/gosseract) **star:3121** Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.   [![godoc][D]](https://godoc.org/github.com/otiai10/gosseract)
- [m2cgen](https://github.com/BayesWitnesses/m2cgen) **star:2991** A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support.
- [tfgo](https://github.com/galeone/tfgo) **star:2492** Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.   [![godoc][D]](https://godoc.org/github.com/galeone/tfgo)
- [goml](https://github.com/cdipaolo/goml) **star:1615** On-line Machine Learning in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/cdipaolo/goml)   [![godoc][D]](https://godoc.org/github.com/cdipaolo/goml)
- [GoMLX](https://github.com/gomlx/gomlx) **star:1481** An accelerated Machine Learning framework for Go.   [![There was an update last month][G]](https://github.com/gomlx/gomlx)   [![godoc][D]](https://godoc.org/github.com/gomlx/gomlx)
- [eaopt](https://github.com/MaxHalford/eaopt) **star:907** An evolutionary optimization library.   [![godoc][D]](https://godoc.org/github.com/MaxHalford/eaopt)
- [onnx-go](https://github.com/owulveryck/onnx-go) **star:902** Go Interface to Open Neural Network Exchange (ONNX).   [![godoc][D]](https://godoc.org/github.com/owulveryck/onnx-go)
- [bayesian](https://github.com/jbrukh/bayesian) **star:814** Naive Bayesian Classification for Golang.   [![godoc][D]](https://godoc.org/github.com/jbrukh/bayesian)
- [ocrserver](https://github.com/otiai10/ocrserver) **star:767** A simple OCR API server, seriously easy to be deployed by Docker and Heroku.   [![It hasn't been updated in recent three years][Y]](https://github.com/otiai10/ocrserver)   [![godoc][D]](https://godoc.org/github.com/otiai10/ocrserver)
- [CloudForest](https://github.com/ryanbressler/CloudForest) **star:747** Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ryanbressler/CloudForest)   [![godoc][D]](https://godoc.org/github.com/ryanbressler/CloudForest)
- [hugot](https://github.com/knights-analytics/hugot) **star:622** Huggingface transformer pipelines for golang with onnxruntime.   [![godoc][D]](https://godoc.org/github.com/knights-analytics/hugot)
- [gobrain](https://github.com/goml/gobrain) **star:568** Neural Networks written in go.   [![It hasn't been updated in recent three years][Y]](https://github.com/goml/gobrain)   [![godoc][D]](https://godoc.org/github.com/goml/gobrain)
- [go-deep](https://github.com/patrikeh/go-deep) **star:558** A feature-rich neural network library in Go.   [![godoc][D]](https://godoc.org/github.com/patrikeh/go-deep)
- [regommend](https://github.com/muesli/regommend) **star:313** Recommendation & collaborative filtering engine.   [![It hasn't been updated in recent three years][Y]](https://github.com/muesli/regommend)   [![godoc][D]](https://godoc.org/github.com/muesli/regommend)
- [Goptuna](https://github.com/c-bata/goptuna) **star:279** Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.   [![godoc][D]](https://godoc.org/github.com/c-bata/goptuna)
- [goga](https://github.com/tomcraven/goga) **star:223** Genetic algorithm library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/tomcraven/goga)   [![godoc][D]](https://godoc.org/github.com/tomcraven/goga)
- [goRecommend](https://github.com/timkaye11/goRecommend) **star:205** Recommendation Algorithms library written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/timkaye11/goRecommend)   [![godoc][D]](https://godoc.org/github.com/timkaye11/goRecommend)
- [go-galib](https://github.com/thoj/go-galib) **star:202** Genetic Algorithms library written in Go / golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/thoj/go-galib)   [![godoc][D]](https://godoc.org/github.com/thoj/go-galib)
- [shield](https://github.com/eaigner/shield) **star:160** Bayesian text classifier with flexible tokenizers and storage backends for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/eaigner/shield)   [![godoc][D]](https://godoc.org/github.com/eaigner/shield)
- [go-fann](https://github.com/white-pony/go-fann) **star:119** Go bindings for Fast Artificial Neural Networks(FANN) library.   [![It hasn't been updated in recent three years][Y]](https://github.com/white-pony/go-fann)   [![godoc][D]](https://godoc.org/github.com/white-pony/go-fann)
- [born](https://github.com/born-ml/born) **star:114** Deep learning framework inspired by Burn (Rust), with autograd, type-safe tensors, and zero-CGO GPU acceleration.   [![There was an update last month][G]](https://github.com/born-ml/born)   [![godoc][D]](https://godoc.org/github.com/born-ml/born)
- [goscore](https://github.com/asafschers/goscore) **star:101** Go Scoring API for PMML.   [![It hasn't been updated in recent three years][Y]](https://github.com/asafschers/goscore)   [![godoc][D]](https://godoc.org/github.com/asafschers/goscore)
- [GoMind](https://github.com/surenderthakran/gomind) **star:98** A simplistic Neural Network Library in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/surenderthakran/gomind)   [![godoc][D]](https://godoc.org/github.com/surenderthakran/gomind)
- [fonet](https://github.com/Fontinalis/fonet) **star:86** A Deep Neural Network library written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/Fontinalis/fonet)   [![godoc][D]](https://godoc.org/github.com/Fontinalis/fonet)
- [gonet](https://github.com/dathoangnd/gonet) **star:82** Neural Network for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/dathoangnd/gonet)   [![godoc][D]](https://godoc.org/github.com/dathoangnd/gonet)
- [neural-go](https://github.com/schuyler/neural-go) **star:73** Multilayer perceptron network implemented in Go, with training via backpropagation.   [![It hasn't been updated in recent three years][Y]](https://github.com/schuyler/neural-go)   [![godoc][D]](https://godoc.org/github.com/schuyler/neural-go)
- [libsvm](https://github.com/datastream/libsvm) **star:72** libsvm golang version derived work based on LIBSVM 3.14.   [![It hasn't been updated in recent three years][Y]](https://github.com/datastream/libsvm)   [![godoc][D]](https://godoc.org/github.com/datastream/libsvm)
- [go-pr](https://github.com/daviddengcn/go-pr) **star:68** Pattern recognition package in Go lang.   [![It hasn't been updated in recent three years][Y]](https://github.com/daviddengcn/go-pr)   [![godoc][D]](https://godoc.org/github.com/daviddengcn/go-pr)
- [randomforest](https://github.com/malaschitz/randomForest) **star:61** Easy to use Random Forest library for Go.   [![godoc][D]](https://godoc.org/github.com/malaschitz/randomForest)
- [Varis](https://github.com/Xamber/Varis) **star:55** Golang Neural Network.   [![It hasn't been updated in recent three years][Y]](https://github.com/Xamber/Varis)   [![godoc][D]](https://godoc.org/github.com/Xamber/Varis)
- [go-cluster](https://github.com/e-XpertSolutions/go-cluster) **star:44** Go implementation of the k-modes and k-prototypes clustering algorithms.   [![It hasn't been updated in recent three years][Y]](https://github.com/e-XpertSolutions/go-cluster)   [![godoc][D]](https://godoc.org/github.com/e-XpertSolutions/go-cluster)
- [ddt](https://github.com/sgrodriguez/ddt) **star:42** Dynamic decision tree, create trees defining customizable rules.   [![It hasn't been updated in recent three years][Y]](https://github.com/sgrodriguez/ddt)   [![godoc][D]](https://godoc.org/github.com/sgrodriguez/ddt)
- [godist](https://github.com/e-dard/godist) **star:42** Various probability distributions, and associated methods.   [![It hasn't been updated in recent three years][Y]](https://github.com/e-dard/godist)   [![godoc][D]](https://godoc.org/github.com/e-dard/godist)
- [evoli](https://github.com/khezen/evoli) **star:33** Genetic Algorithm and Particle Swarm Optimization library.   [![It hasn't been updated in recent three years][Y]](https://github.com/khezen/evoli)   [![godoc][D]](https://godoc.org/github.com/khezen/evoli)
- [Anneal](https://github.com/georgebuilds/anneal) **star:32** Machine learning compiler in Go, a from-scratch tinygrad port with a WebGPU backend.   [![godoc][D]](https://godoc.org/github.com/georgebuilds/anneal)
- [catboost-cgo](https://github.com/mirecl/catboost-cgo) **star:26** Fast, scalable, high performance Gradient Boosting on Decision Trees library. Golang using Cgo for blazing fast inference CatBoost Model.
- [probab](https://github.com/ThePaw/probab) **star:22** Probability distribution functions. Bayesian inference. Written in pure Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ThePaw/probab)   [![godoc][D]](https://godoc.org/github.com/ThePaw/probab)
- [datatrax](https://github.com/rbmuller/datatrax) **star:10** Data engineering and classic ML toolkit with batch processing, type coercion, and 7 algorithms in pure Go with zero dependencies.   [![godoc][D]](https://godoc.org/github.com/rbmuller/datatrax)

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [Asynq](https://github.com/hibiken/asynq) **star:13520** A simple, reliable, and efficient distributed task queue for Go built on top of Redis.   [![godoc][D]](https://godoc.org/github.com/hibiken/asynq)
- [Centrifugo](https://github.com/centrifugal/centrifugo) **star:10509** Real-time messaging (Websockets or SockJS) server in Go.   [![There was an update last month][G]](https://github.com/centrifugal/centrifugo)   [![godoc][D]](https://godoc.org/github.com/centrifugal/centrifugo)
- [gorush](https://github.com/appleboy/gorush) **star:8748** Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).   [![There was an update last month][G]](https://github.com/appleboy/gorush)   [![godoc][D]](https://godoc.org/github.com/appleboy/gorush)
- [machinery](https://github.com/RichardKnop/machinery) **star:7962** Asynchronous task queue/job queue based on distributed message passing.   [![godoc][D]](https://godoc.org/github.com/RichardKnop/machinery)
- [NATS Go Client](https://github.com/nats-io/nats.go) **star:6694** Go client for the NATS   [![There was an update last month][G]](https://github.com/nats-io/nats.go)   [![godoc][D]](https://godoc.org/github.com/nats-io/nats.go)
- [Mercure](https://github.com/dunglas/mercure) **star:5283** Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).   [![There was an update last month][G]](https://github.com/dunglas/mercure)   [![godoc][D]](https://godoc.org/github.com/dunglas/mercure)
- [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) **star:5151** confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform.   [![There was an update last month][G]](https://github.com/confluentinc/confluent-kafka-go)
- [melody](https://github.com/olahol/melody) **star:4080** Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.   [![godoc][D]](https://godoc.org/github.com/olahol/melody)
- [APNs2](https://github.com/sideshow/apns2) **star:3183** HTTP/2 Apple Push Notification provider for Go - Send push notifications to iOS, tvOS, Safari and OSX apps.   [![godoc][D]](https://godoc.org/github.com/sideshow/apns2)
- [go-nsq](https://github.com/nsqio/go-nsq) **star:2654** the official Go package for NSQ.   [![godoc][D]](https://godoc.org/github.com/nsqio/go-nsq)
- [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) **star:2071** gopush-cluster is a go push server cluster.   [![It hasn't been updated in recent three years][Y]](https://github.com/Terry-Mao/gopush-cluster)   [![godoc][D]](https://godoc.org/github.com/Terry-Mao/gopush-cluster)   [![Contains Chinese documents][CN]](https://github.com/Terry-Mao/gopush-cluster)
- [amqp](https://github.com/rabbitmq/amqp091-go) **star:2022** Go RabbitMQ Client Library.   [![There was an update last month][G]](https://github.com/rabbitmq/amqp091-go)   [![godoc][D]](https://godoc.org/github.com/rabbitmq/amqp091-go)
- [EventBus](https://github.com/asaskevich/EventBus) **star:1980** The lightweight event bus with async compatibility.   [![godoc][D]](https://godoc.org/github.com/asaskevich/EventBus)
- [Beaver](https://github.com/Clivern/Beaver) **star:1581** A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.   [![There was an update last month][G]](https://github.com/Clivern/Beaver)   [![godoc][D]](https://godoc.org/github.com/Clivern/Beaver)
- [Chanify](https://github.com/chanify/chanify) **star:1322** A push notification server send message to your iOS devices.   [![It hasn't been updated in recent three years][Y]](https://github.com/chanify/chanify)   [![godoc][D]](https://godoc.org/github.com/chanify/chanify)   [![Contains Chinese documents][CN]](https://github.com/chanify/chanify)
- [dbus](https://github.com/godbus/dbus) **star:1184** Native Go bindings for D-Bus.   [![godoc][D]](https://godoc.org/github.com/godbus/dbus)
- [Gollum](https://github.com/trivago/gollum) **star:940** A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.   [![godoc][D]](https://godoc.org/github.com/trivago/gollum)   [![Archived][Archived]](https://github.com/trivago/gollum)
- [mangos](https://github.com/nanomsg/mangos) **star:757** Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability.   [![godoc][D]](https://godoc.org/github.com/nanomsg/mangos)
- [golongpoll](https://github.com/jcuga/golongpoll) **star:668** HTTP longpoll server library that makes web pub-sub simple.   [![godoc][D]](https://godoc.org/github.com/jcuga/golongpoll)
- [emitter](https://github.com/olebedev/emitter) **star:532** Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.   [![It hasn't been updated in recent three years][Y]](https://github.com/olebedev/emitter)   [![godoc][D]](https://godoc.org/github.com/olebedev/emitter)
- [Glue](https://github.com/desertbit/glue) **star:421** Robust Go and Javascript Socket Library (Alternative to Socket.io).   [![It hasn't been updated in recent three years][Y]](https://github.com/desertbit/glue)   [![godoc][D]](https://godoc.org/github.com/desertbit/glue)
- [Bus](https://github.com/mustafaturan/bus) **star:368** Minimalist message bus implementation for internal communication.   [![It hasn't been updated in recent three years][Y]](https://github.com/mustafaturan/bus)   [![godoc][D]](https://godoc.org/github.com/mustafaturan/bus)
- [messagebus](https://github.com/vardius/message-bus) **star:287** messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.   [![It hasn't been updated in recent three years][Y]](https://github.com/vardius/message-bus)
- [Go-MediatR](https://github.com/mehdihadeli/Go-MediatR) **star:278** A library for handling mediator patterns and simplified CQRS patterns within an event-driven architecture, inspired by csharp MediatR library.   [![godoc][D]](https://godoc.org/github.com/mehdihadeli/Go-MediatR)
- [guble](https://github.com/smancke/guble) **star:161** Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.   [![It hasn't been updated in recent three years][Y]](https://github.com/smancke/guble)   [![godoc][D]](https://godoc.org/github.com/smancke/guble)
- [backlite](https://github.com/mikestefanello/backlite) **star:150** Type-safe, persistent, embedded task queues and background job runner w/ SQLite.   [![godoc][D]](https://godoc.org/github.com/mikestefanello/backlite)
- [hub](https://github.com/leandro-lugaresi/hub) **star:149** A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.   [![It hasn't been updated in recent three years][Y]](https://github.com/leandro-lugaresi/hub)   [![godoc][D]](https://godoc.org/github.com/leandro-lugaresi/hub)
- [go-mq](https://github.com/cheshir/go-mq) **star:91** RabbitMQ client with declarative configuration.   [![godoc][D]](https://godoc.org/github.com/cheshir/go-mq)
- [drone-line](https://github.com/appleboy/drone-line) **star:81** Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.   [![It hasn't been updated in recent three years][Y]](https://github.com/appleboy/drone-line)   [![godoc][D]](https://godoc.org/github.com/appleboy/drone-line)
- [go-notify](https://github.com/TheCreeper/go-notify) **star:71** Native implementation of the freedesktop notification spec.   [![It hasn't been updated in recent three years][Y]](https://github.com/TheCreeper/go-notify)   [![godoc][D]](https://godoc.org/github.com/TheCreeper/go-notify)
- [go-res](https://github.com/jirenius/go-res) **star:69** Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate.   [![godoc][D]](https://godoc.org/github.com/jirenius/go-res)
- [Commander](https://github.com/jeroenrinzema/commander) **star:68** A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.   [![godoc][D]](https://godoc.org/github.com/jeroenrinzema/commander)
- [GoEventBus](https://github.com/Raezil/GoEventBus) **star:68** A blazing‑fast, in‑memory, lock‑free event bus library   [![godoc][D]](https://godoc.org/github.com/Raezil/GoEventBus)
- [event](https://github.com/agoalofalife/event) **star:62** Implementation of the pattern observer.   [![It hasn't been updated in recent three years][Y]](https://github.com/agoalofalife/event)   [![godoc][D]](https://godoc.org/github.com/agoalofalife/event)
- [hare](https://github.com/leozz37/hare) **star:55** A user friendly library for sending messages and listening to TCP sockets.   [![godoc][D]](https://godoc.org/github.com/leozz37/hare)
- [hypermatch](https://github.com/SchwarzIT/hypermatch) **star:36** A very fast and efficient Go library for matching events to a large set of rules   [![godoc][D]](https://godoc.org/github.com/SchwarzIT/hypermatch)
- [ami](https://github.com/kak-tus/ami) **star:34** Go client to reliable queues based on Redis Cluster Streams.   [![It hasn't been updated in recent three years][Y]](https://github.com/kak-tus/ami)   [![godoc][D]](https://godoc.org/github.com/kak-tus/ami)
- [gosd](https://github.com/alexsniffin/gosd) **star:26** A library for scheduling when to dispatch a message to a channel.   [![It hasn't been updated in recent three years][Y]](https://github.com/alexsniffin/gosd)   [![godoc][D]](https://godoc.org/github.com/alexsniffin/gosd)
- [go-vitotrol](https://github.com/maxatome/go-vitotrol) **star:24** Client library to Viessmann Vitotrol web service.   [![godoc][D]](https://godoc.org/github.com/maxatome/go-vitotrol)
- [jazz](https://github.com/socifi/jazz) **star:18** A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.   [![It hasn't been updated in recent three years][Y]](https://github.com/socifi/jazz)   [![godoc][D]](https://godoc.org/github.com/socifi/jazz)
- [broker](https://github.com/qvcloud/broker) **star:11** Production-grade messaging abstraction with a unified API for various brokers and built-in OpenTelemetry integration.   [![godoc][D]](https://godoc.org/github.com/qvcloud/broker)   [![Contains Chinese documents][CN]](https://github.com/qvcloud/broker)
- [gaurun-client](https://github.com/osamingo/gaurun-client) **star:11** Gaurun Client written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/osamingo/gaurun-client)   [![godoc][D]](https://godoc.org/github.com/osamingo/gaurun-client)
- [go-eventbus](https://github.com/stanipetrosyan/go-eventbus) **star:11** Simple Event Bus package for Go.   [![godoc][D]](https://godoc.org/github.com/stanipetrosyan/go-eventbus)
  messaging system.
- [sarama](https://github.com/Shopify/sarama) **star:12505** Go library for Apache Kafka.   [![There was an update last month][G]](https://github.com/Shopify/sarama)   [![godoc][D]](https://godoc.org/github.com/Shopify/sarama)
- [Watermill](https://github.com/ThreeDotsLabs/watermill) **star:9799** Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog.   [![godoc][D]](https://godoc.org/github.com/ThreeDotsLabs/watermill)
- [Uniqush-Push](https://github.com/uniqush/uniqush-push) **star:1563** Redis backed unified push service for server-side notifications to mobile devices.   [![It hasn't been updated in recent three years][Y]](https://github.com/uniqush/uniqush-push)   [![godoc][D]](https://godoc.org/github.com/uniqush/uniqush-push)
- [zmq4](https://github.com/pebbe/zmq4) **star:1259** Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).   [![godoc][D]](https://godoc.org/github.com/pebbe/zmq4)
- [Quamina](https://github.com/timbray/quamina) **star:496** Fast pattern-matching for filtering messages and events.   [![godoc][D]](https://godoc.org/github.com/timbray/quamina)
- [pubsub](https://github.com/tuxychandru/pubsub) **star:454** Simple pubsub package for go.   [![godoc][D]](https://godoc.org/github.com/tuxychandru/pubsub)
- [rabtap](https://github.com/jandelgado/rabtap) **star:287** RabbitMQ swiss army knife cli app.   [![There was an update last month][G]](https://github.com/jandelgado/rabtap)   [![godoc][D]](https://godoc.org/github.com/jandelgado/rabtap)
- [varmq](https://github.com/goptics/varmq) **star:191** A storage-agnostic message queue and worker pool for concurrent Go programs.   [![godoc][D]](https://godoc.org/github.com/goptics/varmq)
- [redisqueue](https://github.com/robinjoseph08/redisqueue) **star:138** redisqueue provides a producer and consumer of a queue that uses Redis streams.   [![godoc][D]](https://godoc.org/github.com/robinjoseph08/redisqueue)
- [Ratus](https://github.com/hyperonym/ratus) **star:124** Ratus is a RESTful asynchronous task queue server.   [![godoc][D]](https://godoc.org/github.com/hyperonym/ratus)
- [rabbitroutine](https://github.com/furdarius/rabbitroutine) **star:115** Lightweight library that handles RabbitMQ auto-reconnect and publishing retries. The library takes into account the need to re-declare entities in RabbitMQ after reconnection.   [![godoc][D]](https://godoc.org/github.com/furdarius/rabbitroutine)
- [oplog](https://github.com/dailymotion/oplog) **star:111** Generic oplog/replication system for REST APIs.   [![godoc][D]](https://godoc.org/github.com/dailymotion/oplog)
- [rabbus](https://github.com/rafaeljesus/rabbus) **star:98** A tiny wrapper over amqp exchanges and queues.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafaeljesus/rabbus)   [![godoc][D]](https://godoc.org/github.com/rafaeljesus/rabbus)
- [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) **star:80** A tiny wrapper around NSQ topic and channel.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafaeljesus/nsq-event-bus)   [![godoc][D]](https://godoc.org/github.com/rafaeljesus/nsq-event-bus)
- [RapidMQ](https://github.com/sybrexsys/RapidMQ) **star:70** RapidMQ is a lightweight and reliable library for managing of the local messages queue.   [![It hasn't been updated in recent three years][Y]](https://github.com/sybrexsys/RapidMQ)   [![godoc][D]](https://godoc.org/github.com/sybrexsys/RapidMQ)
- [rmqconn](https://github.com/sbabiv/rmqconn) **star:23** RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.   [![It hasn't been updated in recent three years][Y]](https://github.com/sbabiv/rmqconn)   [![godoc][D]](https://godoc.org/github.com/sbabiv/rmqconn)

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice](https://github.com/unidoc/unioffice) **star:4893** Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.   [![godoc][D]](https://godoc.org/github.com/unidoc/unioffice)

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [excelize](https://github.com/xuri/excelize) **star:20782** Golang library for reading and writing Microsoft Excel&trade; (XLSX) files.   [![There was an update last month][G]](https://github.com/xuri/excelize)   [![godoc][D]](https://godoc.org/github.com/xuri/excelize)
- [xlsx](https://github.com/tealeg/xlsx) **star:5995** Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.   [![godoc][D]](https://godoc.org/github.com/tealeg/xlsx)   [![Archived][Archived]](https://github.com/tealeg/xlsx)
- [go-excel](https://github.com/szyhf/go-excel) **star:198** A simple and light reader to read a relate-db-like excel as a table.   [![godoc][D]](https://godoc.org/github.com/szyhf/go-excel)
- [xlsx](https://github.com/plandem/xlsx) **star:177** Fast and safe way to read/update your existing Microsoft Excel files in Go programs.   [![It hasn't been updated in recent three years][Y]](https://github.com/plandem/xlsx)   [![godoc][D]](https://godoc.org/github.com/plandem/xlsx)
- [exl](https://github.com/go-the-way/exl) **star:33** Excel binding to struct written in Go.(Only supports Go1.18+)   [![godoc][D]](https://godoc.org/github.com/go-the-way/exl)
- [cellwalker](https://github.com/chonla/cellwalker) **star:3** Virtually traverse Excel cell by cell's name.   [![godoc][D]](https://godoc.org/github.com/chonla/cellwalker)

### Microsoft Word

_Libraries for working with Microsoft Word._

- [godocx](https://github.com/gomutex/godocx) **star:264** Library for reading and writing Microsoft Word (Docx) files.   [![godoc][D]](https://godoc.org/github.com/gomutex/godocx)

**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [fx](https://github.com/uber-go/fx) **star:7600** A dependency injection based application framework for Go (built on top of dig).   [![godoc][D]](https://godoc.org/github.com/uber-go/fx)
- [dig](https://github.com/uber-go/dig) **star:4486** A reflection based dependency injection toolkit for Go.   [![godoc][D]](https://godoc.org/github.com/uber-go/dig)
- [do](https://github.com/samber/do) **star:2767** A dependency injection framework based on Generics.   [![There was an update last month][G]](https://github.com/samber/do)   [![godoc][D]](https://godoc.org/github.com/samber/do)
- [GoLobby/Container](https://github.com/golobby/container) **star:614** GoLobby Container is a lightweight yet powerful IoC dependency injection container for the Go programming language.   [![godoc][D]](https://godoc.org/github.com/golobby/container)
- [goioc/di](https://github.com/goioc/di) **star:379** Spring-inspired Dependency Injection Container.   [![godoc][D]](https://godoc.org/github.com/goioc/di)
- [di](https://github.com/goava/di) **star:242** A dependency injection container for go programming language.   [![godoc][D]](https://godoc.org/github.com/goava/di)
- [kod](https://github.com/go-kod/kod) **star:198** A generics based dependency injection framework for Go.   [![godoc][D]](https://godoc.org/github.com/go-kod/kod)   [![Contains Chinese documents][CN]](https://github.com/go-kod/kod)
- [dingo](https://github.com/i-love-flamingo/dingo) **star:188** A dependency injection toolkit for Go, based on Guice.   [![There was an update last month][G]](https://github.com/i-love-flamingo/dingo)   [![godoc][D]](https://godoc.org/github.com/i-love-flamingo/dingo)
- [gontainer](https://github.com/NVIDIA/gontainer) **star:153** A dependency injection service container for Go projects.   [![godoc][D]](https://godoc.org/github.com/NVIDIA/gontainer)
- [Go-Spring](https://github.com/go-spring/spring-core) **star:84** A high-performance Go framework inspired by Spring Boot, offering DI, auto-configuration, and lifecycle management while maintaining Go's simplicity and efficiency.   [![godoc][D]](https://godoc.org/github.com/go-spring/spring-core)   [![Contains Chinese documents][CN]](https://github.com/go-spring/spring-core)   [![Archived][Archived]](https://github.com/go-spring/spring-core)
- [godi](https://github.com/junioryono/godi) **star:75** Microsoft-style dependency injection for Go with scoped lifetimes and generics.   [![There was an update last month][G]](https://github.com/junioryono/godi)   [![godoc][D]](https://godoc.org/github.com/junioryono/godi)
- [alice](https://github.com/magic003/alice) **star:51** Additive dependency injection container for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/magic003/alice)   [![godoc][D]](https://godoc.org/github.com/magic003/alice)
- [wire](https://github.com/Fs02/wire) **star:40** Strict Runtime Dependency Injection for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/Fs02/wire)   [![godoc][D]](https://godoc.org/github.com/Fs02/wire)
- [linker](https://github.com/logrange/linker) **star:35** A reflection based dependency injection and inversion of control library with components lifecycle support.   [![godoc][D]](https://godoc.org/github.com/logrange/linker)
- [parsley](https://github.com/matzefriedrich/parsley) **star:34** A flexible and modular reflection-based DI library with advanced features like scoped contexts and proxy generation, designed for large-scale Go applications.   [![There was an update last month][G]](https://github.com/matzefriedrich/parsley)   [![godoc][D]](https://godoc.org/github.com/matzefriedrich/parsley)
- [nject](https://github.com/muir/nject) **star:32** A type safe, reflective framework for libraries, tests, http endpoints, and service startup.   [![There was an update last month][G]](https://github.com/muir/nject)   [![godoc][D]](https://godoc.org/github.com/muir/nject)
- [componego](https://github.com/componego/componego) **star:29** A dependency injection framework based on components, allowing dynamic dependency replacement without duplicating code in tests.   [![godoc][D]](https://godoc.org/github.com/componego/componego)
- [cosban/di](https://gitlab.com/cosban/di)  A code generation based dependency injection wiring tool.
- [ore](https://github.com/firasdarwish/ore) **star:27** Lightweight, generic & simple dependency injection (DI) container.   [![godoc][D]](https://godoc.org/github.com/firasdarwish/ore)
- [gocontainer](https://github.com/vardius/gocontainer) **star:21** Simple Dependency Injection Container.   [![It hasn't been updated in recent three years][Y]](https://github.com/vardius/gocontainer)   [![godoc][D]](https://godoc.org/github.com/vardius/gocontainer)
- [gontainer/gontainer](https://github.com/gontainer/gontainer) **star:17** A YAML-based Dependency Injection container for GO. It supports dependencies' scopes, and auto-detection of circular dependencies. Gontainer is concurrent-safe.   [![godoc][D]](https://godoc.org/github.com/gontainer/gontainer)
- [autowire](https://github.com/tiendc/autowire) **star:12** Dependency injection using Generics and reflection.   [![godoc][D]](https://godoc.org/github.com/tiendc/autowire)
- [boot-go](http://github.com/boot-go/boot)  Component-based development with dependency injection using reflections for Go developers.
- [kinit](https://github.com/go-kata/kinit) **star:10** Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-kata/kinit)   [![godoc][D]](https://godoc.org/github.com/go-kata/kinit)
- [HnH/di](https://github.com/HnH/di) **star:9** DI container library that is focused on clean API and flexibility.   [![godoc][D]](https://godoc.org/github.com/HnH/di)
- [go-beans](https://github.com/go-beans/go) **star:1** Spring-inspired dependency injection and application lifecycle framework for Go.   [![godoc][D]](https://godoc.org/github.com/go-beans/go)

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) **star:56300** Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful.
- [go-blueprint](https://github.com/Melkeydev/go-blueprint) **star:8885** Allows users to spin up a quick Go project using a popular framework.   [![godoc][D]](https://godoc.org/github.com/Melkeydev/go-blueprint)
- [ardanlabs/service](https://github.com/ardanlabs/service) **star:4093** A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications.   [![godoc][D]](https://godoc.org/github.com/ardanlabs/service)
- [goxygen](https://github.com/shpota/goxygen) **star:3594** Generate a modern Web project with Go and Angular, React, or Vue in seconds.   [![godoc][D]](https://godoc.org/github.com/shpota/goxygen)
- [pagoda](https://github.com/mikestefanello/pagoda) **star:2944** Rapid, easy full-stack web development starter kit built in Go.   [![godoc][D]](https://godoc.org/github.com/mikestefanello/pagoda)
- [nunu](https://github.com/go-nunu/nunu) **star:2586** Nunu is a scaffolding tool for building Go applications.   [![godoc][D]](https://godoc.org/github.com/go-nunu/nunu)   [![Contains Chinese documents][CN]](https://github.com/go-nunu/nunu)
- [modern-go-application](https://github.com/sagikazarmark/modern-go-application) **star:1945** Go application boilerplate and example applying modern practices.   [![godoc][D]](https://godoc.org/github.com/sagikazarmark/modern-go-application)
- [goapp](https://github.com/naughtygopher/goapp) **star:1067** An opinionated guideline to structure & develop a Go web application/service.   [![godoc][D]](https://godoc.org/github.com/naughtygopher/goapp)
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) **star:737** A Go application boilerplate template for quick starting projects following production best practices.   [![godoc][D]](https://godoc.org/github.com/lacion/cookiecutter-golang)
- [go-starter](https://github.com/allaboutapps/go-starter) **star:617** An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers.   [![godoc][D]](https://godoc.org/github.com/allaboutapps/go-starter)
- [go-ddd](https://github.com/sklinkert/go-ddd) **star:577** Domain-Driven Design template with CQRS, value objects, idempotent commands, and a transactional outbox.   [![There was an update last month][G]](https://github.com/sklinkert/go-ddd)   [![godoc][D]](https://godoc.org/github.com/sklinkert/go-ddd)
- [golang-templates/seed](https://github.com/golang-templates/seed) **star:564** Go application GitHub repository template.   [![There was an update last month][G]](https://github.com/golang-templates/seed)
- [go-todo-backend](https://github.com/Fs02/go-todo-backend) **star:338** Go Todo Backend example using modular project layout for product microservice.   [![godoc][D]](https://godoc.org/github.com/Fs02/go-todo-backend)
- [scaffold](https://github.com/catchplay/scaffold) **star:150** Scaffold generates a starter Go project layout. Lets you focus on business logic implemented.   [![It hasn't been updated in recent three years][Y]](https://github.com/catchplay/scaffold)   [![godoc][D]](https://godoc.org/github.com/catchplay/scaffold)
- [go-sample](https://github.com/zitryss/go-sample) **star:140** A sample layout for Go application projects with the real code.   [![It hasn't been updated in recent three years][Y]](https://github.com/zitryss/go-sample)   [![godoc][D]](https://godoc.org/github.com/zitryss/go-sample)
- [kickstart.go](https://github.com/raeperd/kickstart.go) **star:111** Minimalistic single-file Go HTTP server template without third-party dependencies.   [![godoc][D]](https://godoc.org/github.com/raeperd/kickstart.go)
- [gobase](https://github.com/wajox/gobase) **star:66** A simple skeleton for golang application with basic setup for real golang application.   [![It hasn't been updated in recent three years][Y]](https://github.com/wajox/gobase)   [![godoc][D]](https://godoc.org/github.com/wajox/gobase)
- [go-rest-api-boilerplate](https://github.com/vahiiiid/go-rest-api-boilerplate) **star:60** AI-friendly, production-ready Go REST API boilerplate with clean architecture, JWT authentication, RBAC, PostgreSQL, Docker hot-reload, and Swagger documentation.   [![godoc][D]](https://godoc.org/github.com/vahiiiid/go-rest-api-boilerplate)
- [go-module](https://github.com/octomation/go-module) **star:40** Template for a typical module written on Go.   [![There was an update last month][G]](https://github.com/octomation/go-module)
- [wangyoucao577/go-project-layout](https://github.com/wangyoucao577/go-project-layout) **star:26** Set of practices and discussions on how to structure Go project layout.   [![godoc][D]](https://godoc.org/github.com/wangyoucao577/go-project-layout)
- [insidieux/inizio](https://github.com/insidieux/inizio) **star:19** Golang project layout generator with plugins.   [![It hasn't been updated in recent three years][Y]](https://github.com/insidieux/inizio)   [![godoc][D]](https://godoc.org/github.com/insidieux/inizio)

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [gobeam/Stringy](https://github.com/gobeam/Stringy) **star:248** String manipulation library to convert string to camel case, snake case, kebab case / slugify etc.   [![godoc][D]](https://godoc.org/github.com/gobeam/Stringy)
- [caps](https://github.com/chanced/caps) **star:61** A case conversion library.   [![godoc][D]](https://godoc.org/github.com/chanced/caps)
- [go-formatter](https://gitlab.com/tymonx/go-formatter)  Implements **replacement fields** surrounded by curly braces `{}` format strings.
- [bexp](https://github.com/happy-sdk/happy/tree/main/pkg/strings/bexp)  Go implementation of Brace Expansion mechanism to generate arbitrary strings.
- [strcase](https://github.com/charlievieth/strcase) **star:16** Case-insensitive implementation of the standard library's strings/bytes packages.   [![godoc][D]](https://godoc.org/github.com/charlievieth/strcase)
- [str](https://github.com/schigh/str) **star:10** Pipeline-first string toolkit for composing transformations.   [![godoc][D]](https://godoc.org/github.com/schigh/str)
- [stringFormatter](https://github.com/Wissance/stringFormatter) - String formatting like in Python or C# manner with the additional text formatting features.
- [xstrings](https://github.com/huandu/xstrings) **star:1415** Collection of useful string functions ported from other languages.   [![godoc][D]](https://godoc.org/github.com/huandu/xstrings)
- [sttr](https://github.com/abhimanyu003/sttr) **star:1336** cross-platform, cli app to perform various operations on string.   [![godoc][D]](https://godoc.org/github.com/abhimanyu003/sttr)
- [strutil](https://github.com/ozgio/strutil) **star:206** String utilities.   [![godoc][D]](https://godoc.org/github.com/ozgio/strutil)

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [gopsutil](https://github.com/shirou/gopsutil) **star:11877** Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).   [![There was an update last month][G]](https://github.com/shirou/gopsutil)   [![godoc][D]](https://godoc.org/github.com/shirou/gopsutil)
- [gatus](https://github.com/TwinProduction/gatus) **star:11513** Automated service health dashboard.   [![There was an update last month][G]](https://github.com/TwinProduction/gatus)   [![godoc][D]](https://godoc.org/github.com/TwinProduction/gatus)
- [gofakeit](https://github.com/brianvoe/gofakeit) **star:5372** Random data generator written in go.   [![godoc][D]](https://godoc.org/github.com/brianvoe/gofakeit)
- [purego](https://github.com/ebitengine/purego) **star:3724** A library for calling C functions from Go without Cgo.   [![There was an update last month][G]](https://github.com/ebitengine/purego)
- [base64Captcha](https://github.com/mojocn/base64Captcha) **star:2370** Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.   [![godoc][D]](https://godoc.org/github.com/mojocn/base64Captcha)   [![Contains Chinese documents][CN]](https://github.com/mojocn/base64Captcha)
- [go-resiliency](https://github.com/eapache/go-resiliency) **star:2343** Resiliency patterns for golang.   [![godoc][D]](https://godoc.org/github.com/eapache/go-resiliency)
- [shoutrrr](https://github.com/containrrr/shoutrrr) **star:1628** Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others.   [![godoc][D]](https://godoc.org/github.com/containrrr/shoutrrr)
- [gosms](https://github.com/haxpax/gosms) **star:1469** Your own local SMS gateway in Go that can be used to send SMS.   [![It hasn't been updated in recent three years][Y]](https://github.com/haxpax/gosms)   [![godoc][D]](https://godoc.org/github.com/haxpax/gosms)
- [stateless](https://github.com/qmuntal/stateless) **star:1364** A fluent library for creating state machines.   [![godoc][D]](https://godoc.org/github.com/qmuntal/stateless)
- [llvm](https://github.com/llir/llvm) **star:1274** Library for interacting with LLVM IR in pure Go.   [![godoc][D]](https://godoc.org/github.com/llir/llvm)
- [go-commons-pool](https://github.com/jolestar/go-commons-pool) **star:1239** Generic object pool for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/jolestar/go-commons-pool)   [![godoc][D]](https://godoc.org/github.com/jolestar/go-commons-pool)   [![Contains Chinese documents][CN]](https://github.com/jolestar/go-commons-pool)
- [go-openapi](https://github.com/go-openapi)  Collection of packages to parse and utilize open-api schemas.
- [shortid](https://github.com/teris-io/shortid) **star:960** Distributed generation of super short, unique, non-sequential, URL friendly IDs.   [![It hasn't been updated in recent three years][Y]](https://github.com/teris-io/shortid)   [![godoc][D]](https://godoc.org/github.com/teris-io/shortid)
- [health](https://github.com/alexliesenfeld/health) **star:833** A simple and flexible health check library for Go.   [![godoc][D]](https://godoc.org/github.com/alexliesenfeld/health)
- [xz](https://github.com/ulikunitz/xz) **star:559** Pure golang package for reading and writing xz-compressed files.   [![There was an update last month][G]](https://github.com/ulikunitz/xz)   [![godoc][D]](https://godoc.org/github.com/ulikunitz/xz)
- [banner](https://github.com/dimiro1/banner) **star:462** Add beautiful banners into your Go applications.   [![It hasn't been updated in recent three years][Y]](https://github.com/dimiro1/banner)   [![godoc][D]](https://godoc.org/github.com/dimiro1/banner)
- [health](https://github.com/dimiro1/health) **star:449** Easy to use, extensible health check library.   [![godoc][D]](https://godoc.org/github.com/dimiro1/health)
- [gountries](https://github.com/pariz/gountries) **star:435** Package that exposes country and subdivision data.   [![godoc][D]](https://godoc.org/github.com/pariz/gountries)
- [archives](https://github.com/mholt/archives) **star:430** a cross-platform, multi-format Go library for working with archives and compression formats with a unified API and as virtual file systems compatible with io/fs.   [![godoc][D]](https://godoc.org/github.com/mholt/archives)
- [lk](https://github.com/hyperboloide/lk) **star:419** A simple licensing library for golang.   [![godoc][D]](https://godoc.org/github.com/hyperboloide/lk)
- [conv](https://github.com/cstockton/go-conv) **star:382** Package conv provides fast and intuitive conversions across Go types.   [![It hasn't been updated in recent three years][Y]](https://github.com/cstockton/go-conv)   [![godoc][D]](https://godoc.org/github.com/cstockton/go-conv)
- [gtree](https://github.com/ddddddO/gtree) **star:355** Provide CLI, Package and Web for tree output and directories creation from Markdown or programmatically.   [![There was an update last month][G]](https://github.com/ddddddO/gtree)   [![godoc][D]](https://godoc.org/github.com/ddddddO/gtree)
- [ffmt](https://github.com/go-ffmt/ffmt) **star:316** Beautify data display for Humans.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-ffmt/ffmt)   [![godoc][D]](https://godoc.org/github.com/go-ffmt/ffmt)   [![Contains Chinese documents][CN]](https://github.com/go-ffmt/ffmt)
- [go-unarr](https://github.com/gen2brain/go-unarr) **star:310** Decompression library for RAR, TAR, ZIP and 7z archives.   [![godoc][D]](https://godoc.org/github.com/gen2brain/go-unarr)
- [healthcheck](https://github.com/etherlabsio/healthcheck) **star:275** An opinionated and concurrent health-check HTTP handler for RESTful services.   [![godoc][D]](https://godoc.org/github.com/etherlabsio/healthcheck)
- [battery](https://github.com/distatus/battery) **star:273** Cross-platform, normalized battery information library.   [![godoc][D]](https://godoc.org/github.com/distatus/battery)
- [antch](https://github.com/antchfx/antch) **star:266** A fast, powerful and extensible web crawling & scraping framework.   [![It hasn't been updated in recent three years][Y]](https://github.com/antchfx/antch)   [![godoc][D]](https://godoc.org/github.com/antchfx/antch)   [![Contains Chinese documents][CN]](https://github.com/antchfx/antch)
- [bitio](https://github.com/icza/bitio) **star:257** Highly optimized bit-level Reader and Writer for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/icza/bitio)   [![godoc][D]](https://godoc.org/github.com/icza/bitio)
- [stats](https://github.com/go-playground/stats) **star:172** Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/stats)   [![godoc][D]](https://godoc.org/github.com/go-playground/stats)
- [turtle](https://github.com/hackebrot/turtle) **star:164** Emojis for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/hackebrot/turtle)   [![godoc][D]](https://godoc.org/github.com/hackebrot/turtle)
- [captcha](https://github.com/steambap/captcha) **star:163** Package captcha provides an easy to use, unopinionated API for captcha generation.   [![godoc][D]](https://godoc.org/github.com/steambap/captcha)
- [gommit](https://github.com/antham/gommit) **star:117** Analyze git commit messages to ensure they follow defined patterns.   [![There was an update last month][G]](https://github.com/antham/gommit)   [![godoc][D]](https://godoc.org/github.com/antham/gommit)
- [indigo](https://github.com/osamingo/indigo) **star:112** Distributed unique ID generator of using Sonyflake and encoded by Base58.   [![godoc][D]](https://godoc.org/github.com/osamingo/indigo)
- [gotoprom](https://github.com/cabify/gotoprom) **star:108** Type-safe metrics builder wrapper library for the official Prometheus client.   [![godoc][D]](https://godoc.org/github.com/cabify/gotoprom)
- [faker](https://github.com/pioz/faker) **star:103** Random fake data and struct generator for Go.   [![godoc][D]](https://godoc.org/github.com/pioz/faker)
- [persian](https://github.com/mavihq/persian) **star:94** Some utilities for Persian language in go.   [![godoc][D]](https://godoc.org/github.com/mavihq/persian)
- [morse](https://github.com/alwindoss/morse) **star:87** Library to convert to and from morse code.   [![It hasn't been updated in recent three years][Y]](https://github.com/alwindoss/morse)   [![godoc][D]](https://godoc.org/github.com/alwindoss/morse)
- [pdfgen](https://github.com/hyperboloide/pdfgen) **star:75** HTTP service to generate PDF from Json requests.   [![It hasn't been updated in recent three years][Y]](https://github.com/hyperboloide/pdfgen)   [![godoc][D]](https://godoc.org/github.com/hyperboloide/pdfgen)
- [goffi](https://github.com/go-webgpu/goffi) **star:69** Pure Go FFI with libffi-style typed call interface and structured error handling for calling C libraries without CGO.   [![There was an update last month][G]](https://github.com/go-webgpu/goffi)   [![godoc][D]](https://godoc.org/github.com/go-webgpu/goffi)
- [xkg](https://github.com/go-xkg/xkg) **star:61** X Keyboard Grabber.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-xkg/xkg)   [![godoc][D]](https://godoc.org/github.com/go-xkg/xkg)
- [url-shortener](https://github.com/pantrif/url-shortener) **star:51** A modern, powerful, and robust URL shortener microservice with mysql support.   [![It hasn't been updated in recent three years][Y]](https://github.com/pantrif/url-shortener)   [![godoc][D]](https://godoc.org/github.com/pantrif/url-shortener)
- [VarHandler](https://github.com/azr/generators/tree/master/varhandler)  Generate boilerplate http input and output handling.
- [datacounter](https://github.com/miolini/datacounter) **star:50** Go counters for readers/writer/http.ResponseWriter.   [![It hasn't been updated in recent three years][Y]](https://github.com/miolini/datacounter)   [![godoc][D]](https://godoc.org/github.com/miolini/datacounter)
- [xdg](https://github.com/rkoesters/xdg) **star:50** FreeDesktop.org (xdg) Specs implemented in Go.   [![godoc][D]](https://godoc.org/github.com/rkoesters/xdg)
- [browscap_go](https://github.com/digitalcrab/browscap_go) **star:49** GoLang Library for [Browser Capabilities Project](https://browscap.org/).   [![It hasn't been updated in recent three years][Y]](https://github.com/digitalcrab/browscap_go)   [![godoc][D]](https://godoc.org/github.com/digitalcrab/browscap_go)   [![Archived][Archived]](https://github.com/digitalcrab/browscap_go)
- [sandid](https://github.com/aofei/sandid) **star:46** Every grain of sand on earth has its own ID.   [![godoc][D]](https://godoc.org/github.com/aofei/sandid)   [![Archived][Archived]](https://github.com/aofei/sandid)
- [autoflags](https://github.com/artyom/autoflags) **star:42** Go package to automatically define command line flags from struct fields.   [![It hasn't been updated in recent three years][Y]](https://github.com/artyom/autoflags)   [![godoc][D]](https://godoc.org/github.com/artyom/autoflags)
- [numa](https://github.com/lrita/numa) **star:40** NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.   [![godoc][D]](https://godoc.org/github.com/lrita/numa)
- [gosh](https://github.com/osamingo/gosh) **star:37** Provide Go Statistics Handler, Struct, Measure Method.   [![godoc][D]](https://godoc.org/github.com/osamingo/gosh)
- [shellwords](https://github.com/Wing924/shellwords) **star:28** A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.   [![It hasn't been updated in recent three years][Y]](https://github.com/Wing924/shellwords)   [![godoc][D]](https://godoc.org/github.com/Wing924/shellwords)
- [metrics](https://github.com/pascaldekloe/metrics) **star:27** Library for metrics instrumentation and Prometheus exposition.   [![It hasn't been updated in recent three years][Y]](https://github.com/pascaldekloe/metrics)   [![godoc][D]](https://godoc.org/github.com/pascaldekloe/metrics)
- [fake-useragent](https://github.com/lib4u/fake-useragent) **star:19** Up-to-date simple useragent faker with real world database in Golang   [![godoc][D]](https://godoc.org/github.com/lib4u/fake-useragent)
- [avgRating](https://github.com/kirillDanshin/avgRating) **star:17** Calculate average score and rating based on Wilson Score Equation.   [![It hasn't been updated in recent three years][Y]](https://github.com/kirillDanshin/avgRating)   [![godoc][D]](https://godoc.org/github.com/kirillDanshin/avgRating)
- [anagent](https://github.com/mudler/anagent) **star:16** Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.   [![It hasn't been updated in recent three years][Y]](https://github.com/mudler/anagent)   [![godoc][D]](https://godoc.org/github.com/mudler/anagent)
- [go-commandbus](https://github.com/lana/go-commandbus) **star:15** A slight and pluggable command-bus for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/lana/go-commandbus)   [![godoc][D]](https://godoc.org/github.com/lana/go-commandbus)
- [varint](https://github.com/chmike/varint) **star:15** A faster varying length integer encoder/decoder than the one provided in the standard library.   [![godoc][D]](https://godoc.org/github.com/chmike/varint)
- [hostutils](https://github.com/Wing924/hostutils) **star:13** A golang library for packing and unpacking FQDNs list.   [![godoc][D]](https://godoc.org/github.com/Wing924/hostutils)
- [basexx](https://github.com/bobg/basexx) **star:6** Convert to, from, and between digit strings in various number bases.   [![godoc][D]](https://godoc.org/github.com/bobg/basexx)
- [sitemap-format](https://github.com/mingard/sitemap-format) **star:6** A simple sitemap generator, with a little syntactic sugar.   [![It hasn't been updated in recent three years][Y]](https://github.com/mingard/sitemap-format)   [![godoc][D]](https://godoc.org/github.com/mingard/sitemap-format)
- [common](https://github.com/kubeservice-stack/common) **star:5** A library for server framework.   [![There was an update last month][G]](https://github.com/kubeservice-stack/common)   [![godoc][D]](https://godoc.org/github.com/kubeservice-stack/common)
**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [lingua-go](https://github.com/pemistahl/lingua-go) **star:1358** An accurate natural language detection library, suitable for long and short text alike. Supports detecting multiple languages in mixed-language text.   [![godoc][D]](https://godoc.org/github.com/pemistahl/lingua-go)
- [whatlanggo](https://github.com/abadojack/whatlanggo) **star:689** Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).   [![It hasn't been updated in recent three years][Y]](https://github.com/abadojack/whatlanggo)   [![godoc][D]](https://godoc.org/github.com/abadojack/whatlanggo)
- [getlang](https://github.com/rylans/getlang) **star:175** Fast natural language detection package.   [![It hasn't been updated in recent three years][Y]](https://github.com/rylans/getlang)   [![godoc][D]](https://godoc.org/github.com/rylans/getlang)
- [guesslanguage](https://github.com/endeveit/guesslanguage) **star:58** Functions to determine the natural language of a unicode text.   [![It hasn't been updated in recent three years][Y]](https://github.com/endeveit/guesslanguage)   [![godoc][D]](https://godoc.org/github.com/endeveit/guesslanguage)
- [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) **star:26** Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.   [![godoc][D]](https://godoc.org/github.com/detectlanguage/detectlanguage-go)

### Morphological Analyzers

- [spaGO](https://github.com/nlpodyssey/spago) **star:1850** Self-contained Machine Learning and Natural Language Processing library in Go.   [![godoc][D]](https://godoc.org/github.com/nlpodyssey/spago)
- [kagome](https://github.com/ikawaha/kagome) **star:973** JP morphological analyzer written in pure Go.   [![godoc][D]](https://godoc.org/github.com/ikawaha/kagome)
- [nlp](https://github.com/james-bowman/nlp) **star:475** Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).   [![It hasn't been updated in recent three years][Y]](https://github.com/james-bowman/nlp)   [![godoc][D]](https://godoc.org/github.com/james-bowman/nlp)
- [RAKE.go](https://github.com/afjoseph/RAKE.Go) **star:124** Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).   [![godoc][D]](https://godoc.org/github.com/afjoseph/RAKE.Go)
- [go-stem](https://github.com/agonopol/go-stem) **star:81** Implementation of the porter stemming algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/agonopol/go-stem)   [![godoc][D]](https://godoc.org/github.com/agonopol/go-stem)
- [go2vec](https://github.com/danieldk/go2vec) **star:58** Reader and utility functions for word2vec embeddings.   [![It hasn't been updated in recent three years][Y]](https://github.com/danieldk/go2vec)   [![godoc][D]](https://godoc.org/github.com/danieldk/go2vec)   [![Archived][Archived]](https://github.com/danieldk/go2vec)
- [govader](https://github.com/jonreiter/govader) **star:55** Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment).   [![godoc][D]](https://godoc.org/github.com/jonreiter/govader)
- [porter2](https://github.com/zhenjl/porter2) **star:47** Really fast Porter 2 stemmer.   [![It hasn't been updated in recent three years][Y]](https://github.com/zhenjl/porter2)   [![godoc][D]](https://godoc.org/github.com/zhenjl/porter2)
- [snowball](https://github.com/goodsign/snowball) **star:38** Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).   [![It hasn't been updated in recent three years][Y]](https://github.com/goodsign/snowball)
- [paicehusk](https://github.com/rookii/paicehusk) **star:29** Golang implementation of the Paice/Husk Stemming Algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/rookii/paicehusk)   [![godoc][D]](https://godoc.org/github.com/rookii/paicehusk)
- [golibstemmer](https://github.com/rjohnsondev/golibstemmer) **star:21** Go bindings for the snowball libstemmer library including porter 2.   [![It hasn't been updated in recent three years][Y]](https://github.com/rjohnsondev/golibstemmer)   [![godoc][D]](https://godoc.org/github.com/rjohnsondev/golibstemmer)
- [porter](https://github.com/a2800276/porter) **star:14** This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.   [![godoc][D]](https://godoc.org/github.com/a2800276/porter)
- [libtextcat](https://github.com/goodsign/libtextcat) **star:13** Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.   [![It hasn't been updated in recent three years][Y]](https://github.com/goodsign/libtextcat)   [![godoc][D]](https://godoc.org/github.com/goodsign/libtextcat)
- [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) **star:11** Sentiment analyzer using sentiwordnet lexicon in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/dinopuguh/gosentiwordnet)   [![godoc][D]](https://godoc.org/github.com/dinopuguh/gosentiwordnet)
- [govader-backend](https://github.com/PIMPfiction/govader_backend) **star:6** Microservice implementation of [GoVader](https://github.com/jonreiter/govader).   [![godoc][D]](https://godoc.org/github.com/PIMPfiction/govader_backend)
- [go-propisyu](https://github.com/rekurt/go-propisyu) **star:5** Convert numbers to Russian words with correct grammatical gender and noun declension.   [![godoc][D]](https://godoc.org/github.com/rekurt/go-propisyu)
- [spelling-corrector](https://github.com/jorelosorio/spellingcorrector) **star:2** A spelling corrector for the Spanish language or create your own.   [![It hasn't been updated in recent three years][Y]](https://github.com/jorelosorio/spellingcorrector)   [![godoc][D]](https://godoc.org/github.com/jorelosorio/spellingcorrector)

### Slugifiers

- [slug](https://github.com/gosimple/slug) **star:1329** URL-friendly slugify with multiple languages support.   [![godoc][D]](https://godoc.org/github.com/gosimple/slug)
- [go-slugify](https://github.com/mozillazg/go-slugify) **star:97** Make pretty slug with multiple languages support.   [![It hasn't been updated in recent three years][Y]](https://github.com/mozillazg/go-slugify)   [![godoc][D]](https://godoc.org/github.com/mozillazg/go-slugify)
- [Slugify](https://github.com/avelino/slugify) **star:35** Go slugify application that handles string.   [![It hasn't been updated in recent three years][Y]](https://github.com/avelino/slugify)   [![godoc][D]](https://godoc.org/github.com/avelino/slugify)

### Tokenizers

- [gse](https://github.com/go-ego/gse) **star:2837** Go efficient text segmentation; support english, chinese, japanese and other.   [![godoc][D]](https://godoc.org/github.com/go-ego/gse)   [![Contains Chinese documents][CN]](https://github.com/go-ego/gse)
- [gojieba](https://github.com/yanyiwu/gojieba) **star:2646** This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.   [![godoc][D]](https://godoc.org/github.com/yanyiwu/gojieba)   [![Contains Chinese documents][CN]](https://github.com/yanyiwu/gojieba)
- [sentences](https://github.com/neurosnap/sentences) **star:473** Sentence tokenizer: converts text into a list of sentences.   [![godoc][D]](https://godoc.org/github.com/neurosnap/sentences)
- [segment](https://github.com/blevesearch/segment) **star:89** Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/)   [![It hasn't been updated in recent three years][Y]](https://github.com/blevesearch/segment)   [![godoc][D]](https://godoc.org/github.com/blevesearch/segment)
- [textcat](https://github.com/pebbe/textcat) **star:73** Go package for n-gram based text categorization, with support for utf-8 and raw text.   [![godoc][D]](https://godoc.org/github.com/pebbe/textcat)
- [MMSEGO](https://github.com/awsong/MMSEGO) **star:62** This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/awsong/MMSEGO)   [![godoc][D]](https://godoc.org/github.com/awsong/MMSEGO)
- [stemmer](https://github.com/dchest/stemmer) **star:56** Stemmer packages for Go programming language. Includes English and German stemmers.   [![It hasn't been updated in recent three years][Y]](https://github.com/dchest/stemmer)   [![godoc][D]](https://godoc.org/github.com/dchest/stemmer)
- [gotokenizer](https://github.com/xujiajun/gotokenizer) **star:21** A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)   [![It hasn't been updated in recent three years][Y]](https://github.com/xujiajun/gotokenizer)   [![godoc][D]](https://godoc.org/github.com/xujiajun/gotokenizer)
- [shamoji](https://github.com/osamingo/shamoji) **star:13** The shamoji is word filtering package written in Go.   [![godoc][D]](https://godoc.org/github.com/osamingo/shamoji)

### Translation

- [go-pinyin](https://github.com/mozillazg/go-pinyin) **star:1784** CN Hanzi to Hanyu Pinyin converter.   [![godoc][D]](https://godoc.org/github.com/mozillazg/go-pinyin)
- [gotext](https://github.com/leonelquinteros/gotext) **star:506** GNU gettext utilities for Go.   [![godoc][D]](https://godoc.org/github.com/leonelquinteros/gotext)
- [spreak](https://github.com/vorlif/spreak) **star:94** Flexible translation and humanization library for Go, based on the concepts behind gettext.   [![godoc][D]](https://godoc.org/github.com/vorlif/spreak)
- [ctxi18n](https://github.com/invopop/ctxi18n/)  Context aware i18n with a short and consise API, pluralization, interpolation, and `fs.FS` support. YAML locale definitions are based on [Rails i18n](https://guides.rubyonrails.org/i18n.html).
- [go-i18n](https://github.com/nicksnyder/go-i18n/)  Package and an accompanying tool to work with localized text.
- [iuliia-go](https://github.com/mehanizm/iuliia-go) **star:57** Transliterate Cyrillic → Latin in every possible way.   [![godoc][D]](https://godoc.org/github.com/mehanizm/iuliia-go)
- [go-mystem](https://github.com/dveselov/mystem) **star:36** CGo bindings to Yandex.Mystem - russian morphology analyzer.   [![It hasn't been updated in recent three years][Y]](https://github.com/dveselov/mystem)   [![godoc][D]](https://godoc.org/github.com/dveselov/mystem)
- [t](https://github.com/youthlin/t) **star:22** Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template.   [![godoc][D]](https://godoc.org/github.com/youthlin/t)   [![Contains Chinese documents][CN]](https://github.com/youthlin/t)
- [go-words](https://github.com/saleh-rahimzadeh/go-words) **star:8** A words table and text resource library for Golang projects.   [![godoc][D]](https://godoc.org/github.com/saleh-rahimzadeh/go-words)

### Transliteration

- [go-unidecode](https://github.com/mozillazg/go-unidecode) **star:146** ASCII transliterations of Unicode text.   [![It hasn't been updated in recent three years][Y]](https://github.com/mozillazg/go-unidecode)   [![godoc][D]](https://godoc.org/github.com/mozillazg/go-unidecode)
- [gounidecode](https://github.com/fiam/gounidecode) **star:80** Unicode transliterator (also known as unidecode) for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/fiam/gounidecode)   [![godoc][D]](https://godoc.org/github.com/fiam/gounidecode)
- [transliterator](https://github.com/alexsergivan/transliterator) **star:46** Provides one-way string transliteration with supporting of language-specific transliteration rules.   [![godoc][D]](https://godoc.org/github.com/alexsergivan/transliterator)
- [enca](https://github.com/endeveit/enca) **star:19** Minimal cgo bindings for [libenca](https://cihar.com/software/enca/), which detects character encodings.   [![It hasn't been updated in recent three years][Y]](https://github.com/endeveit/enca)   [![godoc][D]](https://godoc.org/github.com/endeveit/enca)

**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [fasthttp](https://github.com/valyala/fasthttp) **star:23415** Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.   [![There was an update last month][G]](https://github.com/valyala/fasthttp)   [![godoc][D]](https://godoc.org/github.com/valyala/fasthttp)
- [webrtc](https://github.com/pions/webrtc) **star:16636** A pure Go implementation of the WebRTC API.   [![There was an update last month][G]](https://github.com/pions/webrtc)   [![godoc][D]](https://godoc.org/github.com/pions/webrtc)
- [cloudflared](https://github.com/cloudflare/cloudflared) **star:14851** Cloudflare Tunnel client (formerly Argo Tunnel).   [![There was an update last month][G]](https://github.com/cloudflare/cloudflared)   [![godoc][D]](https://godoc.org/github.com/cloudflare/cloudflared)
- [quic-go](https://github.com/lucas-clemente/quic-go) **star:11697** An implementation of the QUIC protocol in pure Go.   [![There was an update last month][G]](https://github.com/lucas-clemente/quic-go)   [![godoc][D]](https://godoc.org/github.com/lucas-clemente/quic-go)
- [gnet](https://github.com/panjf2000/gnet) **star:11209** `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.   [![There was an update last month][G]](https://github.com/panjf2000/gnet)   [![godoc][D]](https://godoc.org/github.com/panjf2000/gnet)   [![Contains Chinese documents][CN]](https://github.com/panjf2000/gnet)
- [dns](https://github.com/miekg/dns) **star:8727** Go library for working with DNS.   [![There was an update last month][G]](https://github.com/miekg/dns)   [![godoc][D]](https://godoc.org/github.com/miekg/dns)
- [gopacket](https://github.com/google/gopacket) **star:6790** Go library for packet processing with libpcap bindings.   [![godoc][D]](https://godoc.org/github.com/google/gopacket)
- [GoProxy](https://github.com/elazarl/goproxy) **star:6708** A library to create a customized HTTP/HTTPS proxy server using Go.   [![There was an update last month][G]](https://github.com/elazarl/goproxy)   [![godoc][D]](https://godoc.org/github.com/elazarl/goproxy)
- [tun2socks](https://github.com/xjasonlyu/tun2socks) **star:5399** A pure go implementation of tun2socks powered by [gVisor](https://gvisor.dev/) TCP/IP stack.   [![There was an update last month][G]](https://github.com/xjasonlyu/tun2socks)   [![godoc][D]](https://godoc.org/github.com/xjasonlyu/tun2socks)
- [netpoll](https://github.com/cloudwego/netpoll) **star:4593** A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance.   [![godoc][D]](https://godoc.org/github.com/cloudwego/netpoll)   [![Contains Chinese documents][CN]](https://github.com/cloudwego/netpoll)
- [kcp-go](https://github.com/xtaci/kcp-go) **star:4528** KCP - Fast and Reliable ARQ Protocol.   [![godoc][D]](https://godoc.org/github.com/xtaci/kcp-go)
- [ssh](https://github.com/gliderlabs/ssh) **star:4156** Higher-level API for building SSH servers (wraps crypto/ssh).   [![godoc][D]](https://godoc.org/github.com/gliderlabs/ssh)
- [HTTPLab](https://github.com/gchaincl/httplab) **star:4150** HTTPLabs let you inspect HTTP requests and forge responses.   [![godoc][D]](https://godoc.org/github.com/gchaincl/httplab)
- [gobgp](https://github.com/osrg/gobgp) **star:4087** BGP implemented in the Go Programming Language.   [![There was an update last month][G]](https://github.com/osrg/gobgp)   [![godoc][D]](https://godoc.org/github.com/osrg/gobgp)
- [fortio](https://github.com/fortio/fortio) **star:3707** Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.   [![There was an update last month][G]](https://github.com/fortio/fortio)   [![godoc][D]](https://godoc.org/github.com/fortio/fortio)
- [nbio](https://github.com/lesismal/nbio) **star:2753** Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use.   [![There was an update last month][G]](https://github.com/lesismal/nbio)   [![godoc][D]](https://godoc.org/github.com/lesismal/nbio)
- [net](https://golang.org/x/net)  This repository holds supplementary Go networking libraries.
- [water](https://github.com/songgao/water) **star:2165** Simple TUN/TAP library.   [![godoc][D]](https://godoc.org/github.com/songgao/water)
- [go-getter](https://github.com/hashicorp/go-getter) **star:1822** Go library for downloading files or directories from various sources using a URL.   [![There was an update last month][G]](https://github.com/hashicorp/go-getter)   [![godoc][D]](https://godoc.org/github.com/hashicorp/go-getter)
- [gws](https://github.com/lxzan/gws) **star:1792** High-Performance WebSocket Server & Client With AsyncIO Supporting .   [![There was an update last month][G]](https://github.com/lxzan/gws)   [![godoc][D]](https://godoc.org/github.com/lxzan/gws)
- [gev](https://github.com/Allenxuxu/gev) **star:1775** gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.   [![godoc][D]](https://godoc.org/github.com/Allenxuxu/gev)
- [sftp](https://github.com/pkg/sftp) **star:1661** Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>.   [![There was an update last month][G]](https://github.com/pkg/sftp)   [![godoc][D]](https://godoc.org/github.com/pkg/sftp)
- [grab](https://github.com/cavaliercoder/grab) **star:1480** Go package for managing file downloads.   [![godoc][D]](https://godoc.org/github.com/cavaliercoder/grab)
- [NFF-Go](https://github.com/intel-go/nff-go) **star:1419** Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).   [![It hasn't been updated in recent three years][Y]](https://github.com/intel-go/nff-go)   [![godoc][D]](https://godoc.org/github.com/intel-go/nff-go)
- [ftp](https://github.com/jlaffaye/ftp) **star:1401** Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959).   [![godoc][D]](https://godoc.org/github.com/jlaffaye/ftp)
- [mdns](https://github.com/hashicorp/mdns) **star:1366** Simple mDNS (Multicast DNS) client/server library in Golang.   [![There was an update last month][G]](https://github.com/hashicorp/mdns)   [![godoc][D]](https://godoc.org/github.com/hashicorp/mdns)
- [mqttPaho](https://eclipse.org/paho/clients/golang/)  The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
- [sdns](https://github.com/semihalev/sdns) **star:1065** A high-performance, recursive DNS resolver server with DNSSEC support, focused on preserving privacy.   [![There was an update last month][G]](https://github.com/semihalev/sdns)   [![godoc][D]](https://godoc.org/github.com/semihalev/sdns)
- [gmqtt](https://github.com/DrmagicE/gmqtt) **star:1046** Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.   [![godoc][D]](https://godoc.org/github.com/DrmagicE/gmqtt)   [![Contains Chinese documents][CN]](https://github.com/DrmagicE/gmqtt)
- [vssh](https://github.com/yahoo/vssh) **star:991** Go library for building network and server automation over SSH protocol.   [![godoc][D]](https://godoc.org/github.com/yahoo/vssh)
- [cidranger](https://github.com/yl2chen/cidranger) **star:970** Fast IP to CIDR lookup for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/yl2chen/cidranger)   [![godoc][D]](https://godoc.org/github.com/yl2chen/cidranger)
- [gaio](https://github.com/xtaci/gaio) **star:930** High performance async-io networking for Golang in proactor mode.   [![godoc][D]](https://godoc.org/github.com/xtaci/gaio)
- [easytcp](https://github.com/DarthPestilane/easytcp) **star:817** A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful.   [![godoc][D]](https://godoc.org/github.com/DarthPestilane/easytcp)
- [go-stun](https://github.com/ccding/go-stun) **star:720** Go implementation of the STUN client (RFC 3489 and RFC 5389).   [![godoc][D]](https://godoc.org/github.com/ccding/go-stun)
- [lhttp](https://github.com/fanux/lhttp) **star:688** Powerful websocket framework, build your IM server more easily.   [![It hasn't been updated in recent three years][Y]](https://github.com/fanux/lhttp)   [![godoc][D]](https://godoc.org/github.com/fanux/lhttp)   [![Contains Chinese documents][CN]](https://github.com/fanux/lhttp)
- [peerdiscovery](https://github.com/schollz/peerdiscovery) **star:671** Pure Go library for cross-platform local peer discovery using UDP multicast.   [![godoc][D]](https://godoc.org/github.com/schollz/peerdiscovery)
- [gotcp](https://github.com/gansidui/gotcp) **star:511** Go package for quickly writing tcp applications.   [![godoc][D]](https://godoc.org/github.com/gansidui/gotcp)
- [stun](https://github.com/go-rtc/stun) **star:495** Go implementation of RFC 5389 STUN protocol.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-rtc/stun)   [![godoc][D]](https://godoc.org/github.com/go-rtc/stun)   [![Archived][Archived]](https://github.com/go-rtc/stun)
- [gopcap](https://github.com/akrennmair/gopcap) **star:491** Go wrapper for libpcap.   [![It hasn't been updated in recent three years][Y]](https://github.com/akrennmair/gopcap)   [![godoc][D]](https://godoc.org/github.com/akrennmair/gopcap)
- [winrm](https://github.com/masterzen/winrm) **star:475** Go WinRM client to remotely execute commands on Windows machines.   [![godoc][D]](https://godoc.org/github.com/masterzen/winrm)
- [ftpserverlib](https://github.com/fclairamb/ftpserverlib) **star:471** Fully featured FTP server library.   [![There was an update last month][G]](https://github.com/fclairamb/ftpserverlib)   [![godoc][D]](https://godoc.org/github.com/fclairamb/ftpserverlib)
- [arp](https://github.com/mdlayher/arp) **star:391** Package arp implements the ARP protocol, as described in RFC 826.   [![godoc][D]](https://godoc.org/github.com/mdlayher/arp)
- [dnsmonster](https://github.com/mosajjal/dnsmonster) **star:357** Passive DNS Capture/Monitoring Framework.   [![There was an update last month][G]](https://github.com/mosajjal/dnsmonster)   [![godoc][D]](https://godoc.org/github.com/mosajjal/dnsmonster)
- [ethernet](https://github.com/mdlayher/ethernet) **star:288** Package ethernet implements marshaling and unmarshalling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.   [![It hasn't been updated in recent three years][Y]](https://github.com/mdlayher/ethernet)   [![godoc][D]](https://godoc.org/github.com/mdlayher/ethernet)
- [gNxI](https://github.com/google/gnxi) **star:287** A collection of tools for Network Management that use the gNMI and gNOI protocols.
- [nodepass](https://github.com/NodePassProject/nodepass) **star:267** A secure, efficient TCP/UDP tunneling solution that delivers fast, reliable access across network restrictions using pre-established TCP/QUIC/WebSocket or HTTP/2 connections.   [![There was an update last month][G]](https://github.com/NodePassProject/nodepass)   [![godoc][D]](https://godoc.org/github.com/NodePassProject/nodepass)
- [buffstreams](https://github.com/stabbycutyou/buffstreams) **star:255** Streaming protocolbuffer data over TCP made easy.   [![It hasn't been updated in recent three years][Y]](https://github.com/stabbycutyou/buffstreams)   [![godoc][D]](https://godoc.org/github.com/stabbycutyou/buffstreams)
- [psql-wire](https://github.com/jeroenrinzema/psql-wire) **star:238** PostgreSQL server wire protocol. Build your own server and start serving connections..   [![godoc][D]](https://godoc.org/github.com/jeroenrinzema/psql-wire)
- [jazigo](https://github.com/udhos/jazigo) **star:230** Jazigo is a tool written in Go for retrieving configuration for multiple network devices.   [![godoc][D]](https://godoc.org/github.com/udhos/jazigo)
- [utp](https://github.com/anacrolix/utp) **star:183** Go uTP micro transport protocol implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/anacrolix/utp)   [![godoc][D]](https://godoc.org/github.com/anacrolix/utp)
- [tcpack](https://github.com/lim-yoona/tcpack) **star:164** tcpack is an application protocol based on TCP to Pack and Unpack bytes stream in go program.   [![godoc][D]](https://godoc.org/github.com/lim-yoona/tcpack)   [![Contains Chinese documents][CN]](https://github.com/lim-yoona/tcpack)
- [xtcp](https://github.com/xfxdev/xtcp) **star:160** TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol.   [![It hasn't been updated in recent three years][Y]](https://github.com/xfxdev/xtcp)   [![godoc][D]](https://godoc.org/github.com/xfxdev/xtcp)
- [canopus](https://github.com/zubairhamed/canopus) **star:156** CoAP Client/Server implementation (RFC 7252).   [![It hasn't been updated in recent three years][Y]](https://github.com/zubairhamed/canopus)   [![godoc][D]](https://godoc.org/github.com/zubairhamed/canopus)
- [sslb](https://github.com/eduardonunesp/sslb) **star:152** It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.   [![godoc][D]](https://godoc.org/github.com/eduardonunesp/sslb)
- [iplib](https://github.com/c-robinson/iplib) **star:151** Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)   [![godoc][D]](https://godoc.org/github.com/c-robinson/iplib)
- [bart](https://github.com/gaissmai/bart) **star:148** Package bart provides a Balanced-Routing-Table (BART) for very fast IP to CIDR lookups and more.   [![There was an update last month][G]](https://github.com/gaissmai/bart)   [![godoc][D]](https://godoc.org/github.com/gaissmai/bart)
- [gldap](https://github.com/jimlambrt/gldap) **star:122** gldap provides an ldap server implementation and you provide handlers for its ldap operations.   [![godoc][D]](https://godoc.org/github.com/jimlambrt/gldap)
- [event](https://github.com/cheng-zhongliang/event) **star:119** Simple I/O event notification library written in Golang.   [![godoc][D]](https://godoc.org/github.com/cheng-zhongliang/event)
- [natiu-mqtt](https://github.com/soypat/natiu-mqtt) **star:106** A dead-simple, non-allocating, low level implementation of MQTT well suited for embedded systems.   [![godoc][D]](https://godoc.org/github.com/soypat/natiu-mqtt)
- [go-powerdns](https://github.com/joeig/go-powerdns) **star:104** PowerDNS API bindings for Golang.   [![godoc][D]](https://godoc.org/github.com/joeig/go-powerdns)
- [ether](https://github.com/songgao/ether) **star:82** Cross-platform Go package for sending and receiving ethernet frames.   [![It hasn't been updated in recent three years][Y]](https://github.com/songgao/ether)   [![godoc][D]](https://godoc.org/github.com/songgao/ether)
- [fullproxy](https://github.com/shoriwe/fullproxy) **star:81** A fully featured scriptable and daemon configurable proxy and pivoting toolkit with SOCKS5, HTTP, raw ports and reverse proxy protocols.   [![godoc][D]](https://godoc.org/github.com/shoriwe/fullproxy)
- [dhcp6](https://github.com/mdlayher/dhcp6) **star:80** Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.   [![It hasn't been updated in recent three years][Y]](https://github.com/mdlayher/dhcp6)   [![godoc][D]](https://godoc.org/github.com/mdlayher/dhcp6)
- [fwdctl](https://github.com/alegrey91/fwdctl) **star:72** A simple and intuitive CLI to manage IPTables forwards in your Linux server.   [![There was an update last month][G]](https://github.com/alegrey91/fwdctl)   [![godoc][D]](https://godoc.org/github.com/alegrey91/fwdctl)
- [portproxy](https://github.com/aybabtme/portproxy) **star:58** Simple TCP proxy which adds CORS support to API's which don't support it.   [![It hasn't been updated in recent three years][Y]](https://github.com/aybabtme/portproxy)   [![godoc][D]](https://godoc.org/github.com/aybabtme/portproxy)
- [linkio](https://github.com/ian-kent/linkio) **star:53** Network link speed simulation for Reader/Writer interfaces.   [![It hasn't been updated in recent three years][Y]](https://github.com/ian-kent/linkio)   [![godoc][D]](https://godoc.org/github.com/ian-kent/linkio)
- [httpproxy](https://github.com/wzshiming/httpproxy) **star:33** HTTP proxy handler and dialer.   [![There was an update last month][G]](https://github.com/wzshiming/httpproxy)   [![godoc][D]](https://godoc.org/github.com/wzshiming/httpproxy)
- [go-multiproxy](https://github.com/presbrey/go-multiproxy) **star:29** Library for making HTTP requests through a pool of proxies offering fault tolerance, load balancing, automatic retries, cookie management, and more, via http.Get/Post replacement or http.Client RoundTripper drop-in   [![godoc][D]](https://godoc.org/github.com/presbrey/go-multiproxy)
- [publicip](https://github.com/polera/publicip) **star:29** Package publicip returns your public facing IPv4 address (internet egress).   [![It hasn't been updated in recent three years][Y]](https://github.com/polera/publicip)   [![godoc][D]](https://godoc.org/github.com/polera/publicip)
- [graval](https://github.com/koofr/graval) **star:28** Experimental FTP server framework.   [![It hasn't been updated in recent three years][Y]](https://github.com/koofr/graval)   [![godoc][D]](https://godoc.org/github.com/koofr/graval)
- [gnet](https://github.com/fish-tennis/gnet) **star:27** `gnet` is a high-performance networking framework,especially for game servers.   [![godoc][D]](https://godoc.org/github.com/fish-tennis/gnet)   [![Contains Chinese documents][CN]](https://github.com/fish-tennis/gnet)
- [goshark](https://github.com/sunwxg/goshark) **star:18** Package goshark use tshark to decode IP packet and create data struct to analyse packet.   [![It hasn't been updated in recent three years][Y]](https://github.com/sunwxg/goshark)   [![godoc][D]](https://godoc.org/github.com/sunwxg/goshark)
- [llb](https://github.com/kirillDanshin/llb) **star:16** It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.   [![It hasn't been updated in recent three years][Y]](https://github.com/kirillDanshin/llb)   [![godoc][D]](https://godoc.org/github.com/kirillDanshin/llb)
- [go-sse](https://github.com/lampctl/go-sse) **star:15** Go client and server implementation of HTML server-sent events.   [![godoc][D]](https://godoc.org/github.com/lampctl/go-sse)
- [tspool](https://github.com/two/tspool) **star:14** A TCP Library use worker pool to improve performance and protect your server.   [![It hasn't been updated in recent three years][Y]](https://github.com/two/tspool)   [![godoc][D]](https://godoc.org/github.com/two/tspool)
- [go-pcaplite](https://github.com/alexcfv/go-pcaplite) **star:9** Lightweight live packet capture library with HTTPS SNI extraction.   [![godoc][D]](https://godoc.org/github.com/alexcfv/go-pcaplite)
- [macwifi](https://github.com/jaisonerick/macwifi) **star:7** Wi-Fi scanning and Keychain password retrieval for macOS 13+.   [![godoc][D]](https://godoc.org/github.com/jaisonerick/macwifi)
- [nethawk](https://github.com/Flowtriq/nethawk) **star:7** Terminal UI for real-time network traffic capture, analysis, and attack detection with JSON output mode.   [![godoc][D]](https://godoc.org/github.com/Flowtriq/nethawk)
- [cdns](https://github.com/junevm/cdns) **star:5** Change DNS servers effortlessly via terminal.   [![godoc][D]](https://godoc.org/github.com/junevm/cdns)
- [fibersse](https://github.com/vinod-morya/fibersse) **star:3** Production-grade Server-Sent Events (SSE) for Fiber v3 with event coalescing, priority lanes, topic wildcards, adaptive throttling, and built-in auth.   [![godoc][D]](https://godoc.org/github.com/vinod-morya/fibersse)
- [gosnmp](https://github.com/soniah/gosnmp) **star:1** Native Go library for performing SNMP actions.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) **star:2338** Retryable HTTP client in Go.   [![godoc][D]](https://godoc.org/github.com/hashicorp/go-retryablehttp)
- [gentleman](https://github.com/h2non/gentleman) **star:1131** Full-featured plugin-driven HTTP client library.   [![godoc][D]](https://godoc.org/github.com/h2non/gentleman)
- [azuretls-client](https://github.com/Noooste/azuretls-client) **star:462** An easy-to-use HTTP client 100% in Go to spoof TLS/JA3 and HTTP2 fingerprint.   [![godoc][D]](https://godoc.org/github.com/Noooste/azuretls-client)
- [go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) **star:442** Get easily stdlib HTTP client, which does not share any state with other clients.   [![godoc][D]](https://godoc.org/github.com/hashicorp/go-cleanhttp)
- [fast-shot](https://github.com/opus-domini/fast-shot) **star:125** Hit your API targets with rapid-fire precision using Go's fastest and simple HTTP Client.   [![There was an update last month][G]](https://github.com/opus-domini/fast-shot)   [![godoc][D]](https://godoc.org/github.com/opus-domini/fast-shot)
- [go-zoox/fetch](https://github.com/go-zoox/fetch) **star:91** A Powerful, Lightweight, Easy Http Client, inspired by Web Fetch API.   [![godoc][D]](https://godoc.org/github.com/go-zoox/fetch)
- [go-otelroundtripper](https://github.com/NdoleStudio/go-otelroundtripper) **star:88** Go http.RoundTripper that emits open telemetry metrics for HTTP requests.   [![godoc][D]](https://godoc.org/github.com/NdoleStudio/go-otelroundtripper)
- [go-http-client](https://github.com/bozd4g/go-http-client) **star:84** Make http calls simply and easily.   [![godoc][D]](https://godoc.org/github.com/bozd4g/go-http-client)
- [axios4go](https://github.com/rezmoss/axios4go) **star:38** A Go HTTP client library inspired by Axios, providing a simple and intuitive API for making HTTP requests.   [![godoc][D]](https://godoc.org/github.com/rezmoss/axios4go)
- [go-ipmux](https://github.com/optimus-hft/go-ipmux) **star:28** A library for Multiplexing HTTP requests based on multiple Source IPs.   [![godoc][D]](https://godoc.org/github.com/optimus-hft/go-ipmux)
- [go-req](https://github.com/wenerme/go-req) **star:24** Declarative golang HTTP client.   [![godoc][D]](https://godoc.org/github.com/wenerme/go-req)
- [Grequest](https://github.com/lib4u/grequest)  - Simple and lightweight golang package for http requests. based on powerful net/http
- [resty](https://github.com/go-resty/resty) **star:11723** Simple HTTP and REST client for Go inspired by Ruby rest-client.   [![There was an update last month][G]](https://github.com/go-resty/resty)   [![godoc][D]](https://godoc.org/github.com/go-resty/resty)
- [req](https://github.com/imroc/req) **star:4820** Simple Go HTTP client with Black Magic (Less code and More efficiency).   [![There was an update last month][G]](https://github.com/imroc/req)   [![godoc][D]](https://godoc.org/github.com/imroc/req)
- [heimdall](https://github.com/gojektech/heimdall) **star:2772** An enhanced http client with retry and hystrix capabilities.   [![godoc][D]](https://godoc.org/github.com/gojektech/heimdall)
- [grequests](https://github.com/levigross/grequests) **star:2187** A Go "clone" of the great and famous Requests library.   [![godoc][D]](https://godoc.org/github.com/levigross/grequests)
- [surf](https://github.com/enetx/surf) **star:1766** Advanced HTTP client with HTTP/1.1, HTTP/2, HTTP/3 (QUIC), SOCKS5 proxy support and browser-grade TLS fingerprinting.   [![godoc][D]](https://godoc.org/github.com/enetx/surf)
- [tls-client](https://github.com/bogdanfinn/tls-client) **star:1728** net/http.Client like HTTP Client with options to select specific client TLS Fingerprints to use for requests.   [![godoc][D]](https://godoc.org/github.com/bogdanfinn/tls-client)
- [sling](https://github.com/dghubble/sling) **star:1720** Sling is a Go HTTP client library for creating and sending API requests.   [![godoc][D]](https://godoc.org/github.com/dghubble/sling)
- [requests](https://github.com/carlmjohnson/requests) **star:1672** HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools.   [![godoc][D]](https://godoc.org/github.com/carlmjohnson/requests)
- [pester](https://github.com/sethgrid/pester) **star:655** Go HTTP client calls with retries, backoff, and concurrency.   [![It hasn't been updated in recent three years][Y]](https://github.com/sethgrid/pester)   [![godoc][D]](https://godoc.org/github.com/sethgrid/pester)
- [request](https://github.com/monaco-io/request) **star:296** HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency.   [![godoc][D]](https://godoc.org/github.com/monaco-io/request)
- [hedge](https://github.com/bhope/hedge) **star:286** Adaptive hedged requests for Go. Cuts p99 latency with zero configuration, based on Google's "The Tail at Scale" paper.   [![godoc][D]](https://godoc.org/github.com/bhope/hedge)
- [httpretry](https://github.com/ybbus/httpretry) **star:56** Enriches the default go HTTP client with retry functionality.   [![godoc][D]](https://godoc.org/github.com/ybbus/httpretry)
- [rq](https://github.com/ddo/rq) **star:52** A nicer interface for golang stdlib HTTP client.   [![It hasn't been updated in recent three years][Y]](https://github.com/ddo/rq)   [![godoc][D]](https://godoc.org/github.com/ddo/rq)
 - [impersonate-http](https://github.com/North-web-dev/impersonate-http) **star:2** Drop-in net/http.Client with a byte-exact browser TLS (JA3/JA4) and HTTP/2 (Akamai) fingerprint.   [![godoc][D]](https://godoc.org/github.com/North-web-dev/impersonate-http)

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [glfw](https://github.com/go-gl/glfw) **star:1689** Go bindings for GLFW 3.
- [gl](https://github.com/go-gl/gl) **star:1200** Go bindings for OpenGL (generated via glow).
- [mathgl](https://github.com/go-gl/mathgl) **star:607** Pure Go math package specialized for 3D math, with inspiration from GLM.   [![godoc][D]](https://godoc.org/github.com/go-gl/mathgl)
- [goxjs/gl](https://github.com/goxjs/gl) **star:174** Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).   [![It hasn't been updated in recent three years][Y]](https://github.com/goxjs/gl)   [![godoc][D]](https://godoc.org/github.com/goxjs/gl)
- [goxjs/glfw](https://github.com/goxjs/glfw) **star:83** Go cross-platform glfw library for creating an OpenGL context and receiving events.   [![It hasn't been updated in recent three years][Y]](https://github.com/goxjs/glfw)   [![godoc][D]](https://godoc.org/github.com/goxjs/glfw)
- [go-glmatrix](https://github.com/technohippy/go-glmatrix) **star:11** Go port of [glMatrix](https://glmatrix.net/) library.   [![It hasn't been updated in recent three years][Y]](https://github.com/technohippy/go-glmatrix)   [![godoc][D]](https://godoc.org/github.com/technohippy/go-glmatrix)

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [GORM](https://github.com/go-gorm/gorm) **star:39863** The fantastic ORM library for Golang, aims to be developer friendly.   [![godoc][D]](https://godoc.org/github.com/go-gorm/gorm)
- [ent](https://github.com/facebook/ent) **star:17135** An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.   [![godoc][D]](https://godoc.org/github.com/facebook/ent)
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) **star:6989** ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.   [![There was an update last month][G]](https://github.com/volatiletech/sqlboiler)   [![godoc][D]](https://godoc.org/github.com/volatiletech/sqlboiler)
- [bun](https://github.com/uptrace/bun) **star:4882** SQL-first Golang ORM. Successor of go-pg.   [![godoc][D]](https://godoc.org/github.com/uptrace/bun)
- [gorp](https://github.com/go-gorp/gorp) **star:3751** Go Relational Persistence, ORM-ish library for Go.   [![godoc][D]](https://godoc.org/github.com/go-gorp/gorp)
- [upper.io/db](https://github.com/upper/db) **star:3662** Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.   [![godoc][D]](https://godoc.org/github.com/upper/db)
- [XORM](https://gitea.com/xorm/xorm)  Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).
- [gormt](https://github.com/xxjwxc/gormt) **star:2427** Mysql database to golang gorm struct.   [![There was an update last month][G]](https://github.com/xxjwxc/gormt)   [![godoc][D]](https://godoc.org/github.com/xxjwxc/gormt)
- [Prisma](https://github.com/prisma/prisma-client-go) **star:2279** Prisma Client Go, Typesafe database access for Go.   [![godoc][D]](https://godoc.org/github.com/prisma/prisma-client-go)   [![Archived][Archived]](https://github.com/prisma/prisma-client-go)
- [bob](https://github.com/stephenafamo/bob) **star:1744** SQL query builder and ORM/Factory generator for Go. Successor of SQLBoiler.   [![There was an update last month][G]](https://github.com/stephenafamo/bob)   [![godoc][D]](https://godoc.org/github.com/stephenafamo/bob)
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) **star:1717** A flexible and powerful SQL string builder library plus a zero-config ORM.   [![godoc][D]](https://godoc.org/github.com/huandu/go-sqlbuilder)
- [pop/soda](https://github.com/gobuffalo/pop) **star:1515** Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.   [![There was an update last month][G]](https://github.com/gobuffalo/pop)   [![godoc][D]](https://godoc.org/github.com/gobuffalo/pop)
- [reform](https://github.com/go-reform/reform) **star:1457** Better ORM for Go, based on non-empty interfaces and code generation.   [![godoc][D]](https://godoc.org/github.com/go-reform/reform)
- [rel](https://github.com/go-rel/rel) **star:785** Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API.   [![godoc][D]](https://godoc.org/github.com/go-rel/rel)
- [Zoom](https://github.com/albrow/zoom) **star:311** Blazing-fast datastore and querying engine built on Redis.   [![It hasn't been updated in recent three years][Y]](https://github.com/albrow/zoom)   [![godoc][D]](https://godoc.org/github.com/albrow/zoom)
- [go-sql](https://github.com/rushteam/gosql) **star:177** A easy ORM for mysql.   [![It hasn't been updated in recent three years][Y]](https://github.com/rushteam/gosql)   [![godoc][D]](https://godoc.org/github.com/rushteam/gosql)
- [golobby/orm](https://github.com/golobby/orm) **star:163** Simple, fast, type-safe, generic orm for developer happiness.   [![godoc][D]](https://godoc.org/github.com/golobby/orm)
- [grimoire](https://github.com/Fs02/grimoire) **star:161** Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).   [![It hasn't been updated in recent three years][Y]](https://github.com/Fs02/grimoire)   [![godoc][D]](https://godoc.org/github.com/Fs02/grimoire)
- [go-store](https://github.com/gosuri/go-store) **star:113** Simple and fast Redis backed key-value store library for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/gosuri/go-store)   [![godoc][D]](https://godoc.org/github.com/gosuri/go-store)
- [go-firestorm](https://github.com/jschoedt/go-firestorm) **star:53** A simple ORM for Google/Firebase Cloud Firestore.   [![It hasn't been updated in recent three years][Y]](https://github.com/jschoedt/go-firestorm)   [![godoc][D]](https://godoc.org/github.com/jschoedt/go-firestorm)
- [cacheme](https://github.com/Yiling-J/cacheme-go) **star:25** Schema based, typed Redis caching/memoize framework for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/Yiling-J/cacheme-go)   [![godoc][D]](https://godoc.org/github.com/Yiling-J/cacheme-go)   [![Contains Chinese documents][CN]](https://github.com/Yiling-J/cacheme-go)
- [CQL](https://github.com/FrancoLiberali/cql) **star:18** Built on top of GORM, adds compile-time verified queries based on auto-generated code.   [![godoc][D]](https://godoc.org/github.com/FrancoLiberali/cql)
- [go-dbw](https://github.com/hashicorp/go-dbw) **star:18** A simple package that encapsulates database operations.   [![There was an update last month][G]](https://github.com/hashicorp/go-dbw)   [![godoc][D]](https://godoc.org/github.com/hashicorp/go-dbw)
- [lore](https://github.com/abrahambotros/lore) **star:14** Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/abrahambotros/lore)   [![godoc][D]](https://godoc.org/github.com/abrahambotros/lore)
- [marlow](https://github.com/marlow/marlow) **star:12** Generated ORM from project structs for compile time safety assurances.   [![It hasn't been updated in recent three years][Y]](https://github.com/marlow/marlow)   [![godoc][D]](https://godoc.org/github.com/marlow/marlow)

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_

- [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)  Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

_Unofficial libraries for package and dependency management._

- [syft](https://github.com/anchore/syft) **star:9235** A CLI tool and Go library for generating a Software Bill of Materials (SBOM) from container images and filesystems.   [![There was an update last month][G]](https://github.com/anchore/syft)   [![godoc][D]](https://godoc.org/github.com/anchore/syft)
- [gup](https://github.com/nao1215/gup) **star:589** Update binaries installed by "go install".   [![godoc][D]](https://godoc.org/github.com/nao1215/gup)
- [modup](https://github.com/chaindead/modup) **star:65** Terminal UI for Go dependency updates with outdated module detection and selective upgrading.   [![godoc][D]](https://godoc.org/github.com/chaindead/modup)

**[⬆ back to top](#contents)**

## Performance

- [jaeger](https://github.com/jaegertracing/jaeger) **star:22997** A distributed tracing system.   [![There was an update last month][G]](https://github.com/jaegertracing/jaeger)   [![godoc][D]](https://godoc.org/github.com/jaegertracing/jaeger)
- [ebpf-go](https://github.com/cilium/ebpf) **star:7860** Provides utilities for loading, compiling, and debugging eBPF programs.   [![There was an update last month][G]](https://github.com/cilium/ebpf)   [![godoc][D]](https://godoc.org/github.com/cilium/ebpf)
- [pixie](https://github.com/pixie-labs/pixie) **star:6489** No instrumentation tracing for Golang applications via eBPF.   [![There was an update last month][G]](https://github.com/pixie-labs/pixie)
- [statsviz](https://github.com/arl/statsviz) **star:3644** Live visualization of your Go application runtime statistics.   [![There was an update last month][G]](https://github.com/arl/statsviz)   [![godoc][D]](https://godoc.org/github.com/arl/statsviz)
- [profile](https://github.com/pkg/profile) **star:2059** Simple profiling support package for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/pkg/profile)   [![godoc][D]](https://godoc.org/github.com/pkg/profile)
- [go-instrument](https://github.com/nikolaydubina/go-instrument) **star:298** Automatically add spans to all methods and functions.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/go-instrument)
- [mm-go](https://github.com/joetifa2003/mm-go) **star:194** Generic manual memory management for golang.   [![godoc][D]](https://godoc.org/github.com/joetifa2003/mm-go)
- [tracer](https://github.com/kamilsk/tracer) **star:90** Simple, lightweight tracing.   [![It hasn't been updated in recent three years][Y]](https://github.com/kamilsk/tracer)   [![godoc][D]](https://godoc.org/github.com/kamilsk/tracer)
- [otelinji](https://github.com/hedhyw/otelinji) **star:28** OpenTelemetry auto-instrumentation tool for adding spans to functions.   [![godoc][D]](https://godoc.org/github.com/hedhyw/otelinji)
- [go-perfstat](https://github.com/go-perfstat/go) **star:1** Lightweight performance statistics and execution time aggregation for Go.   [![godoc][D]](https://godoc.org/github.com/go-perfstat/go)

**[⬆ back to top](#contents)**

## Query Language

- [gqlgen](https://github.com/99designs/gqlgen) **star:10735** go generate based graphql server library.   [![There was an update last month][G]](https://github.com/99designs/gqlgen)   [![godoc][D]](https://godoc.org/github.com/99designs/gqlgen)
- [graphql-go](https://github.com/graphql-go/graphql) **star:10156** Implementation of GraphQL for Go.   [![godoc][D]](https://godoc.org/github.com/graphql-go/graphql)
- [dasel](https://github.com/tomwright/dasel) **star:8004** Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies.   [![godoc][D]](https://godoc.org/github.com/tomwright/dasel)
- [graphql](https://github.com/neelance/graphql-go) **star:4756** GraphQL server with a focus on ease of use.   [![godoc][D]](https://godoc.org/github.com/neelance/graphql-go)
- [gojsonq](https://github.com/thedevsaddam/gojsonq) **star:2227** A simple Go package to Query over JSON Data.   [![It hasn't been updated in recent three years][Y]](https://github.com/thedevsaddam/gojsonq)   [![godoc][D]](https://godoc.org/github.com/thedevsaddam/gojsonq)
- [play](https://github.com/paololazzari/play) **star:576** A TUI playground to experiment with your favorite programs, such as grep, sed, awk, jq and yq.   [![godoc][D]](https://godoc.org/github.com/paololazzari/play)
- [rql](https://github.com/a8m/rql) **star:366** Resource Query Language for REST API.   [![godoc][D]](https://godoc.org/github.com/a8m/rql)
- [jsonql](https://github.com/elgs/jsonql) **star:280** JSON query expression library in Golang.   [![godoc][D]](https://godoc.org/github.com/elgs/jsonql)
- [jsonslice](https://github.com/bhmj/jsonslice) **star:92** Jsonpath queries with advanced filters.   [![godoc][D]](https://godoc.org/github.com/bhmj/jsonslice)
- [rqp](https://github.com/timsolov/rest-query-parser) **star:91** Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query.   [![godoc][D]](https://godoc.org/github.com/timsolov/rest-query-parser)
- [mql](https://github.com/hashicorp/mql) **star:66** Model Query Language (mql) is a query language for your database models.   [![godoc][D]](https://godoc.org/github.com/hashicorp/mql)
- [goven](https://github.com/SeldonIO/goven) **star:61** A drop-in query language for any database schema.   [![It hasn't been updated in recent three years][Y]](https://github.com/SeldonIO/goven)   [![godoc][D]](https://godoc.org/github.com/SeldonIO/goven)
- [api-fu](https://github.com/ccbrown/api-fu) **star:57** Comprehensive GraphQL implementation.   [![godoc][D]](https://godoc.org/github.com/ccbrown/api-fu)
- [straf](https://github.com/SonicRoshan/straf) **star:40** Easily Convert Golang structs to GraphQL objects.   [![It hasn't been updated in recent three years][Y]](https://github.com/SonicRoshan/straf)   [![godoc][D]](https://godoc.org/github.com/SonicRoshan/straf)
- [jsonpath](https://github.com/AsaiYusuke/jsonpath) **star:32** A query library for retrieving part of JSON based on JSONPath syntax.   [![godoc][D]](https://godoc.org/github.com/AsaiYusuke/jsonpath)
- [gws](https://github.com/Zaba505/gws) **star:6** Apollos' "GraphQL over Websocket" client and server implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/Zaba505/gws)   [![godoc][D]](https://godoc.org/github.com/Zaba505/gws)
- [grapher](https://github.com/reaganiwadha/grapher) **star:4** A GraphQL field builder utilizing Go generics with extra utilities and features.   [![godoc][D]](https://godoc.org/github.com/reaganiwadha/grapher)

**[⬆ back to top](#contents)**

## Reflection

- [Deepcopier](https://github.com/ulule/deepcopier) **star:460** Simple struct copying for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ulule/deepcopier)   [![godoc][D]](https://godoc.org/github.com/ulule/deepcopier)
- [go-deepcopy](https://github.com/tiendc/go-deepcopy) **star:130** Fast deep copy library.   [![godoc][D]](https://godoc.org/github.com/tiendc/go-deepcopy)
- [gotype](https://github.com/wzshiming/gotype) **star:66** Golang source code parsing, usage like reflect package.   [![There was an update last month][G]](https://github.com/wzshiming/gotype)   [![godoc][D]](https://godoc.org/github.com/wzshiming/gotype)   [![Contains Chinese documents][CN]](https://github.com/wzshiming/gotype)
- [copy](https://github.com/gotidy/copy) **star:51** Package for fast copying structs of different types.   [![It hasn't been updated in recent three years][Y]](https://github.com/gotidy/copy)   [![godoc][D]](https://godoc.org/github.com/gotidy/copy)
- [gpath](https://github.com/tenntenn/gpath) **star:41** Library to simplify access struct fields with Go's expression in reflection.   [![It hasn't been updated in recent three years][Y]](https://github.com/tenntenn/gpath)   [![godoc][D]](https://godoc.org/github.com/tenntenn/gpath)
- [goenum](https://github.com/lvyahui8/goenum) **star:16** A common enumeration struct based on generics and reflection that allows you to quickly define enumerations and use a set of useful default methods.   [![godoc][D]](https://godoc.org/github.com/lvyahui8/goenum)
- [reflectutils](https://github.com/muir/reflectutils) **star:10** Helpers for working with reflection: struct tag parsing; recursive walking; fill value from string.   [![godoc][D]](https://godoc.org/github.com/muir/reflectutils)
- [reflectpro](https://github.com/gontainer/reflectpro) **star:9** Callers, copiers, getters and setters for go.   [![godoc][D]](https://godoc.org/github.com/gontainer/reflectpro)
- [objwalker](https://github.com/rekby/objwalker) **star:4** Walk by go objects with reflection.   [![It hasn't been updated in recent three years][Y]](https://github.com/rekby/objwalker)   [![godoc][D]](https://godoc.org/github.com/rekby/objwalker)

**[⬆ back to top](#contents)**

## Resource Embedding

- [vfsgen](https://github.com/shurcooL/vfsgen) **star:983** Generates a vfsdata.go file that statically implements the given virtual filesystem.   [![godoc][D]](https://godoc.org/github.com/shurcooL/vfsgen)
- [debme](https://github.com/leaanthony/debme) **star:35** Create an `embed.FS` from an existing `embed.FS` subdirectory.   [![It hasn't been updated in recent three years][Y]](https://github.com/leaanthony/debme)   [![godoc][D]](https://godoc.org/github.com/leaanthony/debme)
- [embed](https://pkg.go.dev/embed)  Package embed provides access to files embedded in the running Go program.
- [rebed](https://github.com/soypat/rebed) **star:30** Recreate folder structures and files from Go 1.16's `embed.FS` type   [![It hasn't been updated in recent three years][Y]](https://github.com/soypat/rebed)   [![godoc][D]](https://godoc.org/github.com/soypat/rebed)

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [gonum](https://github.com/gonum/gonum) **star:8412** Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.   [![godoc][D]](https://godoc.org/github.com/gonum/gonum)
- [stats](https://github.com/montanaflynn/stats) **star:3022** Statistics package with common functions missing from the Golang standard library.   [![There was an update last month][G]](https://github.com/montanaflynn/stats)   [![godoc][D]](https://godoc.org/github.com/montanaflynn/stats)
- [gonum/plot](https://github.com/gonum/plot) **star:2961** gonum/plot provides an API for building and drawing plots in Go.   [![godoc][D]](https://godoc.org/github.com/gonum/plot)
- [gosl](https://github.com/cpmech/gosl) **star:1877** Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.   [![godoc][D]](https://godoc.org/github.com/cpmech/gosl)
- [streamtools](https://github.com/nytlabs/streamtools) **star:1311** general purpose, graphical tool for dealing with streams of data.   [![godoc][D]](https://godoc.org/github.com/nytlabs/streamtools)
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) **star:1288** Dataframes for machine-learning and statistics (similar to pandas).   [![It hasn't been updated in recent three years][Y]](https://github.com/rocketlaunchr/dataframe-go)   [![godoc][D]](https://godoc.org/github.com/rocketlaunchr/dataframe-go)
- [orb](https://github.com/paulmach/orb) **star:1122** 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.   [![godoc][D]](https://godoc.org/github.com/paulmach/orb)
- [go-dsp](https://github.com/mjibson/go-dsp) **star:915** Digital Signal Processing for Go.   [![godoc][D]](https://godoc.org/github.com/mjibson/go-dsp)
- [chart](https://github.com/vdobler/chart) **star:776** Simple Chart Plotting library for Go. Supports many graphs types.   [![It hasn't been updated in recent three years][Y]](https://github.com/vdobler/chart)   [![godoc][D]](https://godoc.org/github.com/vdobler/chart)
- [graph](https://github.com/yourbasic/graph) **star:752** Library of basic graph algorithms.   [![It hasn't been updated in recent three years][Y]](https://github.com/yourbasic/graph)   [![godoc][D]](https://godoc.org/github.com/yourbasic/graph)
- [goraph](https://github.com/gyuho/goraph) **star:748** Pure Go graph theory library(data structure, algorithm visualization).   [![It hasn't been updated in recent three years][Y]](https://github.com/gyuho/goraph)   [![godoc][D]](https://godoc.org/github.com/gyuho/goraph)
- [Poly](https://github.com/bebop/poly) **star:731** A Go package for engineering organisms.   [![godoc][D]](https://godoc.org/github.com/bebop/poly)
- [taxonkit](https://github.com/shenwei356/taxonkit) **star:475** A practical and efficient NCBI taxonomy toolkit; supports querying lineage, reformatting, filtering, and creating custom taxdump files.   [![godoc][D]](https://godoc.org/github.com/shenwei356/taxonkit)   [![Contains Chinese documents][CN]](https://github.com/shenwei356/taxonkit)
- [ewma](https://github.com/VividCortex/ewma) **star:452** Exponentially-weighted moving averages.   [![godoc][D]](https://godoc.org/github.com/VividCortex/ewma)
- [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) **star:410** Calendar heatmap in plain Go inspired by Github contribution activity.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/calendarheatmap)   [![Archived][Archived]](https://github.com/nikolaydubina/calendarheatmap)
- [go-hep](https://github.com/go-hep/hep) **star:253** A set of libraries and tools for performing High Energy Physics analyses with ease.   [![godoc][D]](https://godoc.org/github.com/go-hep/hep)   [![Archived][Archived]](https://github.com/go-hep/hep)
- [TextRank](https://github.com/DavidBelicza/TextRank) **star:225** TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.   [![godoc][D]](https://godoc.org/github.com/DavidBelicza/TextRank)
- [sparse](https://github.com/james-bowman/sparse) **star:169** Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.   [![godoc][D]](https://godoc.org/github.com/james-bowman/sparse)
- [gograph](https://github.com/hmdsefi/gograph) **star:123** A golang generic graph library that provides mathematical graph-theory and algorithms.   [![godoc][D]](https://godoc.org/github.com/hmdsefi/gograph)
- [go-estimate](https://github.com/milosgajdos/go-estimate) **star:122** State estimation and filtering algorithms in Go.   [![godoc][D]](https://godoc.org/github.com/milosgajdos/go-estimate)
- [pagerank](https://github.com/alixaxel/pagerank) **star:87** Weighted PageRank algorithm implemented in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/alixaxel/pagerank)   [![godoc][D]](https://godoc.org/github.com/alixaxel/pagerank)
- [jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) **star:79** Tool to manipulate JSONL graphs with graphviz support.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/jsonl-graph)
- [geom](https://github.com/skelterjohn/geom) **star:55** 2D geometry for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/skelterjohn/geom)   [![godoc][D]](https://godoc.org/github.com/skelterjohn/geom)
- [insyra](https://github.com/HazelnutParadise/insyra) **star:53** Data analysis library with statistics, visualization, Parquet support, and Python integration.   [![There was an update last month][G]](https://github.com/HazelnutParadise/insyra)   [![godoc][D]](https://godoc.org/github.com/HazelnutParadise/insyra)   [![Contains Chinese documents][CN]](https://github.com/HazelnutParadise/insyra)
- [evaler](https://github.com/soniah/evaler) **star:50** Simple floating point arithmetic expression evaluator.   [![It hasn't been updated in recent three years][Y]](https://github.com/soniah/evaler)   [![godoc][D]](https://godoc.org/github.com/soniah/evaler)
- [decimal](https://github.com/db47h/decimal) **star:45** Package decimal implements arbitrary-precision decimal floating-point arithmetic.   [![It hasn't been updated in recent three years][Y]](https://github.com/db47h/decimal)   [![godoc][D]](https://godoc.org/github.com/db47h/decimal)
- [triangolatte](https://github.com/tchayen/triangolatte) **star:37** 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.   [![It hasn't been updated in recent three years][Y]](https://github.com/tchayen/triangolatte)   [![godoc][D]](https://godoc.org/github.com/tchayen/triangolatte)
- [goent](https://github.com/kzahedi/goent) **star:36** GO Implementation of Entropy Measures.   [![It hasn't been updated in recent three years][Y]](https://github.com/kzahedi/goent)   [![godoc][D]](https://godoc.org/github.com/kzahedi/goent)
- [hdf5](https://github.com/scigolib/hdf5) **star:28** Pure Go implementation of the HDF5 file format for scientific data storage and exchange.   [![godoc][D]](https://godoc.org/github.com/scigolib/hdf5)
- [piecewiselinear](https://github.com/sgreben/piecewiselinear) **star:27** Tiny linear interpolation library.   [![godoc][D]](https://godoc.org/github.com/sgreben/piecewiselinear)
- [godesim](https://github.com/soypat/godesim) **star:23** Extended/multivariable ODE solver framework for event-based simulations with simple API.   [![It hasn't been updated in recent three years][Y]](https://github.com/soypat/godesim)   [![godoc][D]](https://godoc.org/github.com/soypat/godesim)
- [GoStats](https://github.com/OGFris/GoStats) **star:22** GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.   [![It hasn't been updated in recent three years][Y]](https://github.com/OGFris/GoStats)   [![godoc][D]](https://godoc.org/github.com/OGFris/GoStats)
- [ode](https://github.com/ChristopherRabotin/ode) **star:22** Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.   [![It hasn't been updated in recent three years][Y]](https://github.com/ChristopherRabotin/ode)   [![godoc][D]](https://godoc.org/github.com/ChristopherRabotin/ode)
- [PiHex](https://github.com/claygod/PiHex) **star:21** Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.   [![godoc][D]](https://godoc.org/github.com/claygod/PiHex)
- [topk](https://github.com/keilerkonzept/topk) **star:19** Sliding-window and regular top-K sketches, based on the HeavyKeeper algorithm.   [![godoc][D]](https://godoc.org/github.com/keilerkonzept/topk)
- [bradleyterry](https://github.com/seanhagen/bradleyterry) **star:12** Provides a Bradley-Terry Model for pairwise comparisons.   [![It hasn't been updated in recent three years][Y]](https://github.com/seanhagen/bradleyterry)   [![godoc][D]](https://godoc.org/github.com/seanhagen/bradleyterry)
- [matlab](https://github.com/scigolib/matlab) **star:12** Pure Go library for reading and writing MATLAB .mat files (v5-v7.3) without CGO.   [![godoc][D]](https://godoc.org/github.com/scigolib/matlab)
- [rootfinding](https://github.com/khezen/rootfinding) **star:12** root-finding algorithms library for finding roots of quadratic functions.   [![It hasn't been updated in recent three years][Y]](https://github.com/khezen/rootfinding)   [![godoc][D]](https://godoc.org/github.com/khezen/rootfinding)
- [go-gt](https://github.com/ThePaw/go-gt) **star:11** Graph theory algorithms written in "Go" language.   [![It hasn't been updated in recent three years][Y]](https://github.com/ThePaw/go-gt)   [![godoc][D]](https://godoc.org/github.com/ThePaw/go-gt)
- [matrix](https://github.com/Arceus-7/matrix) **star:7** A clean, generic, zero-dependency matrix math package for Go with support for arithmetic, decompositions, and linear system solving.   [![godoc][D]](https://godoc.org/github.com/Arceus-7/matrix)
- [MatProInterface.go](https://github.com/MatProGo-dev/MatProInterface.go) **star:3** MatProInterface.go is an open source package for defining mathematical programs (e.g., convex optimization problems) in Go.   [![godoc][D]](https://godoc.org/github.com/MatProGo-dev/MatProInterface.go)
- [entitydebs](https://github.com/ndabAP/entitydebs) **star:2** A social science tool to programmatically analyze entities in non-fictional texts with a built-in dependency parser.   [![godoc][D]](https://godoc.org/github.com/ndabAP/entitydebs)

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [age](https://github.com/FiloSottile/age) **star:22870** A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability.   [![godoc][D]](https://godoc.org/github.com/FiloSottile/age)
- [lego](https://github.com/go-acme/lego) **star:9735** Pure Go ACME client library and CLI tool (for use with Let's Encrypt).   [![There was an update last month][G]](https://github.com/go-acme/lego)   [![godoc][D]](https://godoc.org/github.com/go-acme/lego)
- [CertMagic](https://github.com/caddyserver/certmagic) **star:5581** Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal.   [![There was an update last month][G]](https://github.com/caddyserver/certmagic)   [![godoc][D]](https://godoc.org/github.com/caddyserver/certmagic)
- [Cameradar](https://github.com/Ullaakut/cameradar) **star:5121** Tool and library to remotely hack RTSP streams from surveillance cameras.   [![There was an update last month][G]](https://github.com/Ullaakut/cameradar)   [![godoc][D]](https://godoc.org/github.com/Ullaakut/cameradar)
- [Coraza](https://github.com/corazawaf/coraza) **star:3658** Enterprise-ready, modsecurity and OWASP CRS compatible WAF library.   [![There was an update last month][G]](https://github.com/corazawaf/coraza)   [![godoc][D]](https://godoc.org/github.com/corazawaf/coraza)
- [memguard](https://github.com/awnumar/memguard) **star:2748** A pure Go library for handling sensitive values in memory.   [![godoc][D]](https://godoc.org/github.com/awnumar/memguard)
- [secure](https://github.com/unrolled/secure) **star:2348** HTTP middleware for Go that facilitates some quick security wins.   [![godoc][D]](https://godoc.org/github.com/unrolled/secure)
- [acmetool](https://github.com/hlandau/acme) **star:2091** ACME (Let's Encrypt) client tool with automatic renewal.   [![It hasn't been updated in recent three years][Y]](https://github.com/hlandau/acme)   [![godoc][D]](https://godoc.org/github.com/hlandau/acme)
- [acopw-go](https://sr.ht/~jamesponddotco/acopw-go/)  Small cryptographically secure password generator package for Go.
- [beelzebub](https://github.com/mariocandela/beelzebub) **star:2088** A secure low code honeypot framework, leveraging AI for System Virtualization.   [![There was an update last month][G]](https://github.com/mariocandela/beelzebub)   [![godoc][D]](https://godoc.org/github.com/mariocandela/beelzebub)
- [themis](https://github.com/cossacklabs/themis) **star:1968** high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps.
- [acra](https://github.com/cossacklabs/acra) **star:1484** Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.   [![godoc][D]](https://godoc.org/github.com/cossacklabs/acra)
- [dongle](https://github.com/golang-module/dongle) **star:1108** A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption.   [![godoc][D]](https://godoc.org/github.com/golang-module/dongle)   [![Contains Chinese documents][CN]](https://github.com/golang-module/dongle)
- [SafeDep/vet](https://github.com/safedep/vet) **star:1091** Protect against malicious open source packages.   [![There was an update last month][G]](https://github.com/safedep/vet)   [![godoc][D]](https://godoc.org/github.com/safedep/vet)
- [booster](https://github.com/anatol/booster) **star:657** Fast initramfs generator with full-disk encryption support.   [![There was an update last month][G]](https://github.com/anatol/booster)   [![godoc][D]](https://godoc.org/github.com/anatol/booster)
- [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) **star:583** Password validator based on raw cryptographic entropy values.   [![It hasn't been updated in recent three years][Y]](https://github.com/lane-c-wagner/go-password-validator)   [![godoc][D]](https://godoc.org/github.com/lane-c-wagner/go-password-validator)
- [nacl](https://github.com/kevinburke/nacl) **star:553** Go implementation of the NaCL set of API's.   [![godoc][D]](https://godoc.org/github.com/kevinburke/nacl)
- [ssh-vault](https://github.com/ssh-vault/ssh-vault) **star:507** encrypt/decrypt using ssh keys.   [![There was an update last month][G]](https://github.com/ssh-vault/ssh-vault)
- [teler-waf](https://github.com/kitabisa/teler-waf) **star:402** teler-waf is a Go HTTP middleware that provide teler IDS functionality to protect against web-based attacks and improve the security of Go-based web applications. It is highly configurable and easy to integrate into existing Go applications.   [![godoc][D]](https://godoc.org/github.com/kitabisa/teler-waf)
- [go-yara](https://github.com/hillu/go-yara) **star:389** Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".   [![godoc][D]](https://godoc.org/github.com/hillu/go-yara)
- [optimus-go](https://github.com/pjebs/optimus-go) **star:369** ID hashing and Obfuscation using Knuth's Algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/pjebs/optimus-go)   [![godoc][D]](https://godoc.org/github.com/pjebs/optimus-go)
- [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) **star:344** A rest application to dynamically update firewalld rules on a linux server.   [![godoc][D]](https://godoc.org/github.com/prashantgupta24/firewalld-rest)
- [go-peer](https://github.com/number571/go-peer) **star:328** A software library for creating secure and anonymous decentralized systems.   [![godoc][D]](https://godoc.org/github.com/number571/go-peer)
- [BadActor](https://github.com/jaredfolkins/badactor) **star:326** In-memory, application-driven jailer built in the spirit of fail2ban.   [![It hasn't been updated in recent three years][Y]](https://github.com/jaredfolkins/badactor)   [![godoc][D]](https://godoc.org/github.com/jaredfolkins/badactor)
- [passlib](https://github.com/hlandau/passlib) **star:289** Futureproof password hashing library.   [![It hasn't been updated in recent three years][Y]](https://github.com/hlandau/passlib)   [![godoc][D]](https://godoc.org/github.com/hlandau/passlib)
- [simple-scrypt](https://github.com/elithrar/simple-scrypt) **star:203** Scrypt package with a simple, obvious API and automatic cost calibration built-in.   [![godoc][D]](https://godoc.org/github.com/elithrar/simple-scrypt)
- [urusai](https://github.com/calpa/urusai) **star:180** Urusai ("noisy" in Japanese) is a Go implementation of a random HTTP/DNS traffic noise generator that helps protect privacy by creating digital smokescreens while browsing.   [![godoc][D]](https://godoc.org/github.com/calpa/urusai)
- [pii-shield](https://github.com/pii-shield/pii-shield) **star:161** Zero-code log sanitization sidecar for Kubernetes that redacts PII from logs.   [![There was an update last month][G]](https://github.com/pii-shield/pii-shield)   [![godoc][D]](https://godoc.org/github.com/pii-shield/pii-shield)
- [luks.go](https://github.com/anatol/luks.go) **star:98** Pure Golang library to manage LUKS partitions.   [![godoc][D]](https://godoc.org/github.com/anatol/luks.go)
- [passwap](https://github.com/zitadel/passwap) **star:77** Provides a unified implementation between different password hashing algorithms   [![There was an update last month][G]](https://github.com/zitadel/passwap)   [![godoc][D]](https://godoc.org/github.com/zitadel/passwap)
- [fort](https://github.com/djadmin/fort) **star:72** Audits macOS security settings across 16 checks, reports a score, and fixes issues where it safely can. Single binary, installable via Homebrew.   [![There was an update last month][G]](https://github.com/djadmin/fort)   [![godoc][D]](https://godoc.org/github.com/djadmin/fort)
- [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) **star:60** A probably paranoid package for securely hashing and encrypting passwords.   [![godoc][D]](https://godoc.org/github.com/dwin/goSecretBoxPassword)
- [go-generate-password](https://github.com/m1/go-generate-password) **star:57** Password generator that can be used on the cli or as a library.   [![It hasn't been updated in recent three years][Y]](https://github.com/m1/go-generate-password)   [![godoc][D]](https://godoc.org/github.com/m1/go-generate-password)
- [go-htpasswd](https://github.com/tg123/go-htpasswd) **star:49** Apache htpasswd Parser for Go.   [![There was an update last month][G]](https://github.com/tg123/go-htpasswd)   [![godoc][D]](https://godoc.org/github.com/tg123/go-htpasswd)
- [certificates](https://github.com/mvmaasakkers/certificates) **star:38** An opinionated tool for generating tls certificates.   [![It hasn't been updated in recent three years][Y]](https://github.com/mvmaasakkers/certificates)   [![godoc][D]](https://godoc.org/github.com/mvmaasakkers/certificates)
- [secret](https://github.com/rsjethani/secret) **star:33** Prevent your secrets from leaking into logs, std\* etc.   [![godoc][D]](https://godoc.org/github.com/rsjethani/secret)
- [secureio](https://github.com/xaionaro-go/secureio) **star:33** An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519.   [![godoc][D]](https://godoc.org/github.com/xaionaro-go/secureio)   [![Archived][Archived]](https://github.com/xaionaro-go/secureio)
- [sslmgr](https://github.com/adrianosela/sslmgr) **star:32** SSL certificates made easy with a high level wrapper around acme/autocert.   [![godoc][D]](https://godoc.org/github.com/adrianosela/sslmgr)
- [argon2-hashing](https://github.com/andskur/argon2-hashing) **star:25** light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package.   [![godoc][D]](https://godoc.org/github.com/andskur/argon2-hashing)
- [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)  Auto provision Let's Encrypt certificates and start a TLS server.
- [procscope](https://github.com/Mutasem-mk4/procscope) **star:23** Process-scoped runtime investigator using eBPF to trace process lifecycle, file activity, and network connections.
- [goArgonPass](https://github.com/dwin/goArgonPass) **star:20** Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.   [![It hasn't been updated in recent three years][Y]](https://github.com/dwin/goArgonPass)   [![godoc][D]](https://godoc.org/github.com/dwin/goArgonPass)   [![Archived][Archived]](https://github.com/dwin/goArgonPass)
- [Razify](https://github.com/Hossiy21/razify) **star:19** CLI to scan, validate and audit .env files for leaked secrets and environment drift.   [![godoc][D]](https://godoc.org/github.com/Hossiy21/razify)
- [qrand](https://github.com/bitfield/qrand) **star:18** Client for the ANU Quantum Numbers (AQN) API, providing quantum-mechanically secure random data.   [![godoc][D]](https://godoc.org/github.com/bitfield/qrand)
- [leakhound](https://github.com/nilpoona/leakhound) **star:17** Static analysis tool to detect accidental logging of sensitive struct fields, preventing data leaks in logs.   [![godoc][D]](https://godoc.org/github.com/nilpoona/leakhound)
- [multikey](https://github.com/adrianosela/multikey) **star:13** An n-out-of-N keys encryption/decryption framework based on Shamir's Secret Sharing algorithm.   [![godoc][D]](https://godoc.org/github.com/adrianosela/multikey)
- [aes-ctr-drbg](https://github.com/sixafter/aes-ctr-drbg) **star:8** A Deterministic Random Bit Generator based on AES in Counter mode (AES-CTR-DRBG) as specified in NIST SP 800-90A.   [![godoc][D]](https://godoc.org/github.com/sixafter/aes-ctr-drbg)
- [canery](https://github.com/rluders/canery) **star:8** Minimal, stateless authorization engine with a pluggable evaluation model.   [![godoc][D]](https://godoc.org/github.com/rluders/canery)
- [veil](https://github.com/getveil/veil) **star:8** Local HTTPS proxy that hides API credentials from AI coding agents. OS keychain integration, format-aware placeholders, SQLite audit log.   [![godoc][D]](https://godoc.org/github.com/getveil/veil)
- [encid](https://github.com/bobg/encid) **star:7** Encode and decode encrypted integer IDs.   [![godoc][D]](https://godoc.org/github.com/bobg/encid)
- [entpassgen](https://github.com/andreimerlescu/entpassgen) **star:7** Entropy Password Generator with extensive command line arguments to generate random strings securely including digits, passwords, and passwords built using obscure dictionary words mixed with symbols and digits.   [![godoc][D]](https://godoc.org/github.com/andreimerlescu/entpassgen)
- [Interpol](https://github.com/avahidi/interpol) **star:7** Rule-based data generator for fuzzing and penetration testing.   [![godoc][D]](https://godoc.org/github.com/avahidi/interpol)
- [Crenox](https://github.com/crenoxhq/crenox) **star:5** Zero-dependency pre-commit secret scanner using Aho-Corasick for high-performance credentials leak detection.   [![There was an update last month][G]](https://github.com/crenoxhq/crenox)   [![godoc][D]](https://godoc.org/github.com/crenoxhq/crenox)
- [dotlock](https://github.com/ahmadraza100/dotlock) **star:4** Encrypted .env vault manager with interactive TUI for managing secrets across multiple environments and profiles.   [![godoc][D]](https://godoc.org/github.com/ahmadraza100/dotlock)
- [gost-crypto](https://github.com/rekurt/gost-crypto) **star:2** Go library for Russian GOST cryptographic standards (digital signatures, Streebog hash, Kuznechik cipher, MGM AEAD) backed by OpenSSL gost-engine.   [![godoc][D]](https://godoc.org/github.com/rekurt/gost-crypto)
- [pm](https://github.com/nicola-strappazzon/password-manager) **star:2** Unix-style password manager written in Go to save your data with OpenPGP encryption.   [![There was an update last month][G]](https://github.com/nicola-strappazzon/password-manager)   [![godoc][D]](https://godoc.org/github.com/nicola-strappazzon/password-manager)
- [redact](https://github.com/alesr/redact) **star:2** Redact sensitive information from slog-based logs using a configurable pipeline.   [![godoc][D]](https://godoc.org/github.com/alesr/redact)
- [gspy](https://github.com/Mutasem-mk4/gspy) **star:1** Forensic goroutine-to-syscall inspector for live Go processes.   [![godoc][D]](https://godoc.org/github.com/Mutasem-mk4/gspy)


**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [jsoniter](https://github.com/json-iterator/go) **star:13896** High-performance 100% compatible drop-in replacement of "encoding/json".   [![godoc][D]](https://godoc.org/github.com/json-iterator/go)   [![Archived][Archived]](https://github.com/json-iterator/go)
- [goprotobuf](https://github.com/golang/protobuf) **star:10084** Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.   [![godoc][D]](https://godoc.org/github.com/golang/protobuf)
- [go-codec](https://github.com/ugorji/go) **star:1957** High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.   [![godoc][D]](https://godoc.org/github.com/ugorji/go)
- [cbor](https://github.com/fxamacker/cbor) **star:1068** Small, safe, and easy CBOR encoding and decoding library.   [![There was an update last month][G]](https://github.com/fxamacker/cbor)   [![godoc][D]](https://godoc.org/github.com/fxamacker/cbor)
- [csvutil](https://github.com/jszwec/csvutil) **star:1033** High Performance, idiomatic CSV record encoding and decoding to native Go structures.   [![godoc][D]](https://godoc.org/github.com/jszwec/csvutil)
- [colfer](https://github.com/pascaldekloe/colfer) **star:758** Code generation for the Colfer binary format.
- [proto](https://github.com/emicklei/proto) **star:616** Parser and writer for Google ProtocolBuffers .proto files.   [![godoc][D]](https://godoc.org/github.com/emicklei/proto)
- [go-capnproto](https://github.com/glycerine/go-capnproto) **star:287** Cap'n Proto library and parser for go.   [![It hasn't been updated in recent three years][Y]](https://github.com/glycerine/go-capnproto)   [![godoc][D]](https://godoc.org/github.com/glycerine/go-capnproto)
- [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) **star:170** GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.   [![It hasn't been updated in recent three years][Y]](https://github.com/yvasiyarov/php_session_decoder)   [![godoc][D]](https://godoc.org/github.com/yvasiyarov/php_session_decoder)
- [structomap](https://github.com/tuvistavie/structomap) **star:145** Library to easily and dynamically generate maps from static structures.   [![It hasn't been updated in recent three years][Y]](https://github.com/tuvistavie/structomap)   [![godoc][D]](https://godoc.org/github.com/tuvistavie/structomap)
- [mus-go](https://github.com/mus-format/mus-go) **star:144** MUS format serializer for Go.   [![godoc][D]](https://godoc.org/github.com/mus-format/mus-go)
- [binstruct](https://github.com/ghostiam/binstruct) **star:114** Golang binary decoder for mapping data into the structure.   [![godoc][D]](https://godoc.org/github.com/ghostiam/binstruct)
- [bambam](https://github.com/glycerine/bambam) **star:64** generator for Cap'n Proto schemas from go.   [![It hasn't been updated in recent three years][Y]](https://github.com/glycerine/bambam)   [![godoc][D]](https://godoc.org/github.com/glycerine/bambam)
- [bel](https://github.com/32leaves/bel) **star:46** Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.   [![godoc][D]](https://godoc.org/github.com/32leaves/bel)
- [fwencoder](https://github.com/o1egl/fwencoder) **star:27** Fixed width file parser (encoding and decoding library) for Go.   [![godoc][D]](https://godoc.org/github.com/o1egl/fwencoder)
- [elastic](https://github.com/epiclabs-io/elastic) **star:25** Convert slices, maps or any other unknown value across different types at run-time, no matter what.   [![It hasn't been updated in recent three years][Y]](https://github.com/epiclabs-io/elastic)   [![godoc][D]](https://godoc.org/github.com/epiclabs-io/elastic)
- [gotiny](https://github.com/raszia/gotiny) **star:21** Efficient Go serialization library, gotiny is almost as fast as serialization libraries that generate code.   [![godoc][D]](https://godoc.org/github.com/raszia/gotiny)
- [pletter](https://github.com/vimeda/pletter) **star:20** A standard way to wrap a proto message for message brokers.   [![godoc][D]](https://godoc.org/github.com/vimeda/pletter)
- [go-csvlib](https://github.com/tiendc/go-csvlib) **star:18** High level and rich functionalities CSV serialization/deserialization library.   [![godoc][D]](https://godoc.org/github.com/tiendc/go-csvlib)
- [fixedwidth](https://github.com/huydang284/fixedwidth) **star:9** Fixed-width text formatting (UTF-8 supported).   [![It hasn't been updated in recent three years][Y]](https://github.com/huydang284/fixedwidth)   [![godoc][D]](https://godoc.org/github.com/huydang284/fixedwidth)
- [unitpacking](https://github.com/recolude/unitpacking) **star:7** Library to pack unit vectors into as fewest bytes as possible.   [![It hasn't been updated in recent three years][Y]](https://github.com/recolude/unitpacking)   [![godoc][D]](https://godoc.org/github.com/recolude/unitpacking)

**[⬆ back to top](#contents)**

## Server Applications

- [Caddy](https://github.com/caddyserver/caddy) **star:74027** Caddy is an alternative, HTTP/2 web server that's easy to configure and use.   [![There was an update last month][G]](https://github.com/caddyserver/caddy)   [![godoc][D]](https://godoc.org/github.com/caddyserver/caddy)
- [consul](https://www.consul.io/)  Consul is a tool for service discovery, monitoring and configuration.
- [pocketbase](https://github.com/pocketbase/pocketbase) **star:59676** PocketBase is a realtime backend in 1 file consisting of embedded database (SQLite) with realtime subscriptions, built-in auth management and much more.   [![There was an update last month][G]](https://github.com/pocketbase/pocketbase)   [![godoc][D]](https://godoc.org/github.com/pocketbase/pocketbase)
- [etcd](https://github.com/etcd-io/etcd) **star:51987** Highly-available key value store for shared configuration and service discovery.   [![There was an update last month][G]](https://github.com/etcd-io/etcd)   [![godoc][D]](https://godoc.org/github.com/etcd-io/etcd)
- [SFTPGo](https://github.com/drakkan/sftpgo) **star:12278** Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage.   [![There was an update last month][G]](https://github.com/drakkan/sftpgo)   [![godoc][D]](https://godoc.org/github.com/drakkan/sftpgo)
- [RoadRunner](https://github.com/spiral/roadrunner) **star:8488** High-performance PHP application server, load-balancer and process manager.   [![There was an update last month][G]](https://github.com/spiral/roadrunner)   [![godoc][D]](https://godoc.org/github.com/spiral/roadrunner)
- [Easegress](https://github.com/megaease/easegress) **star:5872** A cloud native high availability/performance traffic orchestration system with observability and extensibility.   [![godoc][D]](https://godoc.org/github.com/megaease/easegress)
- [Engity's Bifröst](https://bifroest.engity.org/)  Highly customizable SSH server with several ways to authorize a user how to execute its session (local or in containers).
- [Wish](https://github.com/charmbracelet/wish) **star:5329** Make SSH apps, just like that!   [![There was an update last month][G]](https://github.com/charmbracelet/wish)   [![godoc][D]](https://godoc.org/github.com/charmbracelet/wish)
- [flipt](https://github.com/markphelps/flipt) **star:4850** A self contained feature flag solution written in Go and Vue.js   [![There was an update last month][G]](https://github.com/markphelps/flipt)   [![godoc][D]](https://godoc.org/github.com/markphelps/flipt)
- [Fider](https://github.com/getfider/fider) **star:4415** Fider is an open platform to collect and organize customer feedback.   [![There was an update last month][G]](https://github.com/getfider/fider)   [![godoc][D]](https://godoc.org/github.com/getfider/fider)
- [devd](https://github.com/cortesi/devd) **star:3471** Local webserver for developers.   [![godoc][D]](https://godoc.org/github.com/cortesi/devd)
- [algernon](https://github.com/xyproto/algernon) **star:3020** HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.   [![There was an update last month][G]](https://github.com/xyproto/algernon)
- [Flagr](https://github.com/checkr/flagr) **star:2601** Flagr is an open-source feature flagging and A/B testing service.   [![There was an update last month][G]](https://github.com/checkr/flagr)   [![godoc][D]](https://godoc.org/github.com/checkr/flagr)
- [Trickster](https://github.com/tricksterproxy/trickster) **star:2083** HTTP reverse proxy cache and time series accelerator.   [![There was an update last month][G]](https://github.com/tricksterproxy/trickster)   [![godoc][D]](https://godoc.org/github.com/tricksterproxy/trickster)
- [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) **star:2061** A simple, complete and lightweight self-hosted feature flag solution 100% Open Source.   [![There was an update last month][G]](https://github.com/thomaspoignant/go-feature-flag)   [![godoc][D]](https://godoc.org/github.com/thomaspoignant/go-feature-flag)
- [minio](https://github.com/pgsty/minio) **star:1832** Community Maintained Fork of minio (Object Storage Service).   [![godoc][D]](https://godoc.org/github.com/pgsty/minio)   [![Contains Chinese documents][CN]](https://github.com/pgsty/minio)
- [discovery](https://github.com/Bilibili/discovery) **star:1807** A registry for resilient mid-tier load balancing and failover.   [![It hasn't been updated in recent three years][Y]](https://github.com/Bilibili/discovery)   [![godoc][D]](https://godoc.org/github.com/Bilibili/discovery)
- [goshs](https://github.com/patrickhener/goshs) **star:930** SimpleHTTPServer replacement with file upload/download, WebDAV, SFTP, SMB, TLS, authentication, and share links.   [![There was an update last month][G]](https://github.com/patrickhener/goshs)   [![godoc][D]](https://godoc.org/github.com/patrickhener/goshs)
- [OpenRun](https://github.com/openrundev/openrun) **star:898** Open-source alternative to Google Cloud Run and AWS App Runner. Easily deploy internal tools across a team.   [![There was an update last month][G]](https://github.com/openrundev/openrun)   [![godoc][D]](https://godoc.org/github.com/openrundev/openrun)
- [Euterpe](https://github.com/ironsmile/euterpe) **star:574** Self-hosted music streaming server with built-in web UI and REST API.   [![godoc][D]](https://godoc.org/github.com/ironsmile/euterpe)
- [wd-41](https://github.com/baalimago/wd-41) **star:153** A (w)eb (d)evelopment server with automatic live-reload on file changes.   [![godoc][D]](https://godoc.org/github.com/baalimago/wd-41)
- [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) **star:150** Simple Reverse Proxy with Caching, written in Go, using Redis.   [![godoc][D]](https://godoc.org/github.com/fabiocicerchia/go-proxy-cache)
- [dudeldu](https://github.com/krotik/dudeldu) **star:148** A simple SHOUTcast server.   [![It hasn't been updated in recent three years][Y]](https://github.com/krotik/dudeldu)   [![godoc][D]](https://godoc.org/github.com/krotik/dudeldu)
- [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) **star:139** Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels.   [![godoc][D]](https://godoc.org/github.com/blind-oracle/cortex-tenant)
- [lets-proxy2](https://github.com/rekby/lets-proxy2) **star:103** Reverse proxy for handle https with issue certificates in fly from lets-encrypt.   [![godoc][D]](https://godoc.org/github.com/rekby/lets-proxy2)
- [whois](https://github.com/KincaidYang/whois) **star:63** Self-hosted WHOIS/RDAP query service and MCP server for domains, IPv4/IPv6 addresses, CIDRs and ASNs.   [![There was an update last month][G]](https://github.com/KincaidYang/whois)   [![godoc][D]](https://godoc.org/github.com/KincaidYang/whois)   [![Contains Chinese documents][CN]](https://github.com/KincaidYang/whois)
- [psql-streamer](https://github.com/blind-oracle/psql-streamer) **star:61** Stream database events from PostgreSQL to Kafka.   [![It hasn't been updated in recent three years][Y]](https://github.com/blind-oracle/psql-streamer)   [![godoc][D]](https://godoc.org/github.com/blind-oracle/psql-streamer)
- [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) **star:43** Nginx log parser and exporter to Prometheus.   [![It hasn't been updated in recent three years][Y]](https://github.com/blind-oracle/nginx-prometheus)   [![godoc][D]](https://godoc.org/github.com/blind-oracle/nginx-prometheus)
- [nsq](https://nsq.io/)  A realtime distributed messaging platform.
- [protoxy](https://github.com/camgraff/protoxy) **star:36** A proxy server that converts JSON request bodies to Protocol Buffers.   [![It hasn't been updated in recent three years][Y]](https://github.com/camgraff/protoxy)   [![godoc][D]](https://godoc.org/github.com/camgraff/protoxy)
- [Kono](https://github.com/starwalkn/kono) **star:18** lightweight extendable API Gateway in Go - parallel fan-out, flexible aggregation, and zero configuration magic.   [![There was an update last month][G]](https://github.com/starwalkn/kono)   [![godoc][D]](https://godoc.org/github.com/starwalkn/kono)
- [Moxy](https://github.com/sinhashubham95/moxy) **star:15** Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint.   [![It hasn't been updated in recent three years][Y]](https://github.com/sinhashubham95/moxy)   [![godoc][D]](https://godoc.org/github.com/sinhashubham95/moxy)
- [gondola](https://github.com/bmf-san/gondola) **star:11** A YAML based golang reverse proxy.   [![There was an update last month][G]](https://github.com/bmf-san/gondola)   [![godoc][D]](https://godoc.org/github.com/bmf-san/gondola)
- [riemann-relay](https://github.com/blind-oracle/riemann-relay) **star:2** Relay to load-balance Riemann events and/or convert them to Carbon.   [![It hasn't been updated in recent three years][Y]](https://github.com/blind-oracle/riemann-relay)   [![godoc][D]](https://godoc.org/github.com/blind-oracle/riemann-relay)

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-streams](https://github.com/reugn/go-streams) **star:2174** Go stream processing library.   [![godoc][D]](https://godoc.org/github.com/reugn/go-streams)
- [ro](https://github.com/samber/ro) **star:672** Reactive Programming: declarative and composable API for event-driven applications.   [![godoc][D]](https://godoc.org/github.com/samber/ro)
- [go-etl](https://github.com/Breeze0806/go-etl) **star:190** A lightweight toolkit for data source extraction, transformation, and loading (ETL).   [![godoc][D]](https://godoc.org/github.com/Breeze0806/go-etl)   [![Contains Chinese documents][CN]](https://github.com/Breeze0806/go-etl)
- [gostream](https://github.com/mariomac/gostream) **star:172** Type-safe stream processing library inspired by the Java Streams API.   [![There was an update last month][G]](https://github.com/mariomac/gostream)   [![godoc][D]](https://godoc.org/github.com/mariomac/gostream)
- [machine](https://github.com/whitaker-io/machine) **star:169** Go library for writing and generating stream workers with built in metrics and traceability.   [![There was an update last month][G]](https://github.com/whitaker-io/machine)   [![godoc][D]](https://godoc.org/github.com/whitaker-io/machine)
- [stream](https://github.com/youthlin/stream) **star:93** Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce...   [![godoc][D]](https://godoc.org/github.com/youthlin/stream)
- [goio](https://github.com/primetalk/goio) **star:90** An implementation of IO, Stream, Fiber for Golang, inspired by awesome Scala libraries cats and fs2.   [![It hasn't been updated in recent three years][Y]](https://github.com/primetalk/goio)   [![godoc][D]](https://godoc.org/github.com/primetalk/goio)
- [StreamSQL](https://github.com/rulego/streamsql) **star:61** A lightweight streaming SQL engine for real-time data processing.   [![There was an update last month][G]](https://github.com/rulego/streamsql)   [![godoc][D]](https://godoc.org/github.com/rulego/streamsql)   [![Contains Chinese documents][CN]](https://github.com/rulego/streamsql)
- [signals](https://github.com/coregx/signals) **star:18** Type-safe reactive state management inspired by Angular Signals with computed values, effects, and dependency tracking.   [![godoc][D]](https://godoc.org/github.com/coregx/signals)
- [nibbler](https://github.com/naughtygopher/nibbler) **star:17** A lightweight package for micro batch processing.   [![godoc][D]](https://godoc.org/github.com/naughtygopher/nibbler)

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [templ](https://github.com/a-h/templ) **star:10402** A HTML templating language that has great developer tooling.   [![There was an update last month][G]](https://github.com/a-h/templ)   [![godoc][D]](https://godoc.org/github.com/a-h/templ)
- [quicktemplate](https://github.com/valyala/quicktemplate) **star:3326** Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.   [![godoc][D]](https://godoc.org/github.com/valyala/quicktemplate)
- [pongo2](https://github.com/flosch/pongo2) **star:3077** Django-like template-engine for Go.   [![godoc][D]](https://godoc.org/github.com/flosch/pongo2)
- [maroto](https://github.com/johnfercher/maroto) **star:2737** A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.   [![There was an update last month][G]](https://github.com/johnfercher/maroto)   [![godoc][D]](https://godoc.org/github.com/johnfercher/maroto)
- [jet](https://github.com/CloudyKit/jet) **star:1403** Jet template engine.   [![godoc][D]](https://godoc.org/github.com/CloudyKit/jet)
- [fasttemplate](https://github.com/valyala/fasttemplate) **star:909** Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/).   [![godoc][D]](https://godoc.org/github.com/valyala/fasttemplate)
- [gomponents](https://www.gomponents.com)  HTML 5 components in pure Go, that look something like this: `func(name string) g.Node { return Div(Class("headline"), g.Textf("Hi %v!", name)) }`.
- [Razor](https://github.com/sipin/gorazor) **star:884** Razor view engine for Golang.   [![godoc][D]](https://godoc.org/github.com/sipin/gorazor)
- [ego](https://github.com/benbjohnson/ego) **star:590** Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.   [![It hasn't been updated in recent three years][Y]](https://github.com/benbjohnson/ego)   [![godoc][D]](https://godoc.org/github.com/benbjohnson/ego)
- [goview](https://github.com/foolin/goview) **star:463** Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.   [![godoc][D]](https://godoc.org/github.com/foolin/goview)
- [liquid](https://github.com/osteele/liquid) **star:354** Go implementation of Shopify Liquid templates.   [![godoc][D]](https://godoc.org/github.com/osteele/liquid)
- [sprout](https://github.com/go-sprout/sprout) **star:225** Useful template functions for Go templates.   [![godoc][D]](https://godoc.org/github.com/go-sprout/sprout)
- [Soy](https://github.com/robfig/soy) **star:177** Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).   [![godoc][D]](https://godoc.org/github.com/robfig/soy)
- [bagme](https://github.com/boxesandglue/bagme) **star:106** HTML/CSS to PDF rendering with TeX-quality typesetting in pure Go.   [![godoc][D]](https://godoc.org/github.com/boxesandglue/bagme)
- [gox](https://github.com/doors-dev/gox) **star:59** HTML templates as first-class Go expressions, with seamless editor support.   [![godoc][D]](https://godoc.org/github.com/doors-dev/gox)
- [htmgo](https://htmgo.dev)  build simple and scalable systems with go + htmx
- [got](https://github.com/goradd/got) **star:38** A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more.   [![godoc][D]](https://godoc.org/github.com/goradd/got)
- [tbd](https://github.com/lucasepe/tbd) **star:27** A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata.   [![It hasn't been updated in recent three years][Y]](https://github.com/lucasepe/tbd)   [![godoc][D]](https://godoc.org/github.com/lucasepe/tbd)
- [templator](https://github.com/alesr/templator) **star:11** A type-safe HTML template rendering engine for Go.   [![godoc][D]](https://godoc.org/github.com/alesr/templator)

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

### Testing Frameworks

- [Testify](https://github.com/stretchr/testify) **star:26178** Sacred extension to the standard go testing package.   [![There was an update last month][G]](https://github.com/stretchr/testify)   [![godoc][D]](https://godoc.org/github.com/stretchr/testify)
- [keploy](https://github.com/keploy/keploy) **star:18120** Generate Testcase and Data Mocks from API calls automatically.   [![There was an update last month][G]](https://github.com/keploy/keploy)   [![godoc][D]](https://godoc.org/github.com/keploy/keploy)
- [testcontainers-go](https://github.com/testcontainers/testcontainers-go) **star:4917** A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.   [![There was an update last month][G]](https://github.com/testcontainers/testcontainers-go)   [![godoc][D]](https://godoc.org/github.com/testcontainers/testcontainers-go)
- [go-cmp](https://github.com/google/go-cmp) **star:4665** Package for comparing Go values in tests.   [![godoc][D]](https://godoc.org/github.com/google/go-cmp)
- [httpexpect](https://github.com/gavv/httpexpect) **star:2724** Concise, declarative, and easy to use end-to-end HTTP and REST API testing.   [![godoc][D]](https://godoc.org/github.com/gavv/httpexpect)
- [godog](https://github.com/cucumber/godog) **star:2641** Cucumber BDD framework for Go.   [![There was an update last month][G]](https://github.com/cucumber/godog)   [![godoc][D]](https://godoc.org/github.com/cucumber/godog)
- [is](https://github.com/matryer/is) **star:1971** Professional lightweight testing mini-framework for Go.   [![godoc][D]](https://godoc.org/github.com/matryer/is)
- [gnomock](https://github.com/orlangure/gnomock) **star:1491** integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks.   [![godoc][D]](https://godoc.org/github.com/orlangure/gnomock)
- [go-vcr](https://github.com/dnaeon/go-vcr) **star:1393** Record and replay your HTTP interactions for fast, deterministic and accurate tests.   [![godoc][D]](https://godoc.org/github.com/dnaeon/go-vcr)
- [tparse](https://github.com/mfridman/tparse) **star:1259** CLI tool for summarizing go test output. Pipe friendly. Compatible with go test flags.   [![godoc][D]](https://godoc.org/github.com/mfridman/tparse)
- [testfixtures](https://github.com/go-testfixtures/testfixtures) **star:1234** A helper for Rails' like test fixtures to test database applications.   [![There was an update last month][G]](https://github.com/go-testfixtures/testfixtures)   [![godoc][D]](https://godoc.org/github.com/go-testfixtures/testfixtures)
- [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) **star:1208** Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.   [![godoc][D]](https://godoc.org/github.com/fergusstrange/embedded-postgres)
- [goblin](https://github.com/franela/goblin) **star:890** Mocha like testing framework of Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/franela/goblin)   [![godoc][D]](https://godoc.org/github.com/franela/goblin)
- [goc](https://github.com/qiniu/goc) **star:873** Goc is a comprehensive coverage testing system for The Go Programming Language.   [![godoc][D]](https://godoc.org/github.com/qiniu/goc)
- [gocheck](https://labix.org/gocheck)  More advanced testing framework alternative to gotest.
- [GoConvey](https://github.com/smartystreets/goconvey/)  BDD-style framework with web UI and live reload.
- [go-httpbin](https://github.com/mccutchen/go-httpbin) **star:862** HTTP testing and debugging tool with various endpoints for client testing.   [![There was an update last month][G]](https://github.com/mccutchen/go-httpbin)   [![godoc][D]](https://godoc.org/github.com/mccutchen/go-httpbin)
- [baloo](https://github.com/h2non/baloo) **star:779** Expressive and versatile end-to-end HTTP API testing made easy.   [![It hasn't been updated in recent three years][Y]](https://github.com/h2non/baloo)   [![godoc][D]](https://godoc.org/github.com/h2non/baloo)
- [gotest.tools](https://github.com/gotestyourself/gotest.tools) **star:579** A collection of packages to augment the go testing package and support common patterns.   [![godoc][D]](https://godoc.org/github.com/gotestyourself/gotest.tools)
- [go-testdeep](https://github.com/maxatome/go-testdeep) **star:464** Extremely flexible golang deep comparison, extends the go testing package.   [![There was an update last month][G]](https://github.com/maxatome/go-testdeep)   [![godoc][D]](https://godoc.org/github.com/maxatome/go-testdeep)
- [gofight](https://github.com/appleboy/gofight) **star:447** API Handler Testing for Golang Router framework.   [![There was an update last month][G]](https://github.com/appleboy/gofight)   [![godoc][D]](https://godoc.org/github.com/appleboy/gofight)
- [testza](https://github.com/MarvinJWendt/testza) **star:416** Full-featured test framework with nice colorized output.   [![godoc][D]](https://godoc.org/github.com/MarvinJWendt/testza)
- [cupaloy](https://github.com/bradleyjkemp/cupaloy) **star:332** Simple snapshot testing addon for your test framework.   [![It hasn't been updated in recent three years][Y]](https://github.com/bradleyjkemp/cupaloy)   [![godoc][D]](https://godoc.org/github.com/bradleyjkemp/cupaloy)
- [frisby](https://github.com/verdverm/frisby) **star:275** REST API testing framework.   [![It hasn't been updated in recent three years][Y]](https://github.com/verdverm/frisby)   [![godoc][D]](https://godoc.org/github.com/verdverm/frisby)
- [endly](https://github.com/viant/endly) **star:268** Declarative end to end functional testing.   [![godoc][D]](https://godoc.org/github.com/viant/endly)
- [arch-go](https://github.com/arch-go/arch-go) **star:267** Architecture testing tool for Go projects.   [![There was an update last month][G]](https://github.com/arch-go/arch-go)   [![godoc][D]](https://godoc.org/github.com/arch-go/arch-go)
- [got](https://github.com/ysmood/got) **star:266** An enjoyable golang test framework.   [![godoc][D]](https://godoc.org/github.com/ysmood/got)
- [go-hit](https://github.com/Eun/go-hit) **star:256** Hit is an http integration test framework written in golang.   [![godoc][D]](https://godoc.org/github.com/Eun/go-hit)
- [go-carpet](https://github.com/msoap/go-carpet) **star:252** Tool for viewing test coverage in terminal.   [![godoc][D]](https://godoc.org/github.com/msoap/go-carpet)
- [go-test-coverage](https://github.com/vladopajic/go-test-coverage) **star:237** Tool that reports coverage of files below set threshold.   [![There was an update last month][G]](https://github.com/vladopajic/go-test-coverage)   [![godoc][D]](https://godoc.org/github.com/vladopajic/go-test-coverage)
- [commander](https://github.com/SimonBaeumer/commander) **star:231** Tool for testing cli applications on windows, linux and osx.   [![godoc][D]](https://godoc.org/github.com/SimonBaeumer/commander)
- [charlatan](https://github.com/percolate/charlatan) **star:205** Tool to generate fake interface implementations for tests.   [![godoc][D]](https://godoc.org/github.com/percolate/charlatan)
- [dbcleaner](https://github.com/khaiql/dbcleaner) **star:163** Clean database for testing purpose, inspired by `database_cleaner` in Ruby.   [![It hasn't been updated in recent three years][Y]](https://github.com/khaiql/dbcleaner)   [![godoc][D]](https://godoc.org/github.com/khaiql/dbcleaner)
- [jsonassert](https://github.com/kinbiko/jsonassert) **star:142** Package for verifying that your JSON payloads are serialized correctly.   [![godoc][D]](https://godoc.org/github.com/kinbiko/jsonassert)
- [Testo](https://github.com/ozontech/testo) **star:135** Plugin-based testing framework with suites, parallel tests, hooks and parametrization. Inspired by Pytest.   [![There was an update last month][G]](https://github.com/ozontech/testo)   [![godoc][D]](https://godoc.org/github.com/ozontech/testo)
- [be](https://github.com/carlmjohnson/be) **star:133** The minimalist generic test assertion library.   [![godoc][D]](https://godoc.org/github.com/carlmjohnson/be)
- [testcase](https://github.com/adamluzsi/testcase) **star:132** Idiomatic testing framework for Behavior Driven Development.   [![There was an update last month][G]](https://github.com/adamluzsi/testcase)   [![godoc][D]](https://godoc.org/github.com/adamluzsi/testcase)
- [gocrest](https://github.com/corbym/gocrest) **star:108** Composable hamcrest-like matchers for Go assertions.   [![godoc][D]](https://godoc.org/github.com/corbym/gocrest)
- [wstest](https://github.com/posener/wstest) **star:102** Websocket client for unit-testing a websocket http.Handler.   [![It hasn't been updated in recent three years][Y]](https://github.com/posener/wstest)   [![godoc][D]](https://godoc.org/github.com/posener/wstest)
- [gherkingen](https://github.com/hedhyw/gherkingen) **star:97** BDD boilerplate generator and framework.   [![godoc][D]](https://godoc.org/github.com/hedhyw/gherkingen)
- [ginkgo](https://onsi.github.io/ginkgo/)  BDD Testing Framework for Go.
- [testcerts](https://github.com/madflojo/testcerts) **star:86** Dynamically generate self-signed certificates and certificate authorities within your test functions.   [![godoc][D]](https://godoc.org/github.com/madflojo/testcerts)
- [assert](https://github.com/go-playground/assert) **star:68** Basic Assertion Library used along side native go testing, with building blocks for custom assertions.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/assert)   [![godoc][D]](https://godoc.org/github.com/go-playground/assert)
- [restit](https://github.com/yookoala/restit) **star:56** Go micro framework to help writing RESTful API integration test.   [![godoc][D]](https://godoc.org/github.com/yookoala/restit)
- [gospecify](https://github.com/stesla/gospecify) **star:52** This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.   [![It hasn't been updated in recent three years][Y]](https://github.com/stesla/gospecify)   [![godoc][D]](https://godoc.org/github.com/stesla/gospecify)
- [gomatch](https://github.com/jfilipczyk/gomatch) **star:47** library created for testing JSON against patterns.   [![It hasn't been updated in recent three years][Y]](https://github.com/jfilipczyk/gomatch)   [![godoc][D]](https://godoc.org/github.com/jfilipczyk/gomatch)
- [gomega](https://onsi.github.io/gomega/)  Rspec like matcher/assertion library.
- [dsunit](https://github.com/viant/dsunit) **star:46** Datastore testing for SQL, NoSQL, structured files.   [![godoc][D]](https://godoc.org/github.com/viant/dsunit)
- [should](https://github.com/Kairum-Labs/should) **star:44** Testing library with zero dependencies, detailed struct diffs and human-readable error messages.   [![godoc][D]](https://godoc.org/github.com/Kairum-Labs/should)
- [fixenv](https://github.com/rekby/fixenv) **star:34** Fixture manage engine, inspired by pytest fixtures.   [![godoc][D]](https://godoc.org/github.com/rekby/fixenv)
- [Hamcrest](https://github.com/rdrdr/hamcrest) **star:30** fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.   [![It hasn't been updated in recent three years][Y]](https://github.com/rdrdr/hamcrest)   [![godoc][D]](https://godoc.org/github.com/rdrdr/hamcrest)
- [flute](https://github.com/suzuki-shunsuke/flute) **star:23** HTTP client testing framework.   [![There was an update last month][G]](https://github.com/suzuki-shunsuke/flute)   [![godoc][D]](https://godoc.org/github.com/suzuki-shunsuke/flute)
- [schema](https://github.com/jgroeneveld/schema) **star:20** Quick and easy expression matching for JSON schemas used in requests and responses.   [![It hasn't been updated in recent three years][Y]](https://github.com/jgroeneveld/schema)   [![godoc][D]](https://godoc.org/github.com/jgroeneveld/schema)
- [dft](https://github.com/abecodes/dft) **star:19** Lightweight, zero dependency docker containers for testing (or more).   [![godoc][D]](https://godoc.org/github.com/abecodes/dft)
- [gogiven](https://github.com/corbym/gogiven) **star:17** YATSPEC-like BDD testing framework for Go.   [![godoc][D]](https://godoc.org/github.com/corbym/gogiven)
- [testsql](https://github.com/zhulongcheng/testsql) **star:17** Generate test data from SQL files before testing and clear it after finished.   [![It hasn't been updated in recent three years][Y]](https://github.com/zhulongcheng/testsql)   [![godoc][D]](https://godoc.org/github.com/zhulongcheng/testsql)
- [biff](https://github.com/fulldump/biff) **star:14** Bifurcation testing framework, BDD compatible.   [![It hasn't been updated in recent three years][Y]](https://github.com/fulldump/biff)   [![godoc][D]](https://godoc.org/github.com/fulldump/biff)
- [stop-and-go](https://github.com/elgohr/stop-and-go) **star:13** Testing helper for concurrency.   [![godoc][D]](https://godoc.org/github.com/elgohr/stop-and-go)
- [gosuite](https://github.com/pavlo/gosuite) **star:12** Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.   [![It hasn't been updated in recent three years][Y]](https://github.com/pavlo/gosuite)   [![godoc][D]](https://godoc.org/github.com/pavlo/gosuite)
- [apitest](https://apitest.dev)  Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
- [Tt](https://github.com/vcaesar/tt) **star:10** Simple and colorful test tools.   [![godoc][D]](https://godoc.org/github.com/vcaesar/tt)
- [go-testing](https://github.com/tkrop/go-testing) **star:6** Go testing extension, that allows a simple setup of strongly isolated unit, component, and integration test providing advanced mock support extending gomock and gock.   [![godoc][D]](https://godoc.org/github.com/tkrop/go-testing)
- [go-testpredicate](https://github.com/maargenton/go-testpredicate) **star:6** Test predicate style assertions library with extensive diagnostics output.   [![godoc][D]](https://godoc.org/github.com/maargenton/go-testpredicate)
- [trial](https://github.com/jgroeneveld/trial) **star:6** Quick and easy extendable assertions without introducing much boilerplate.   [![It hasn't been updated in recent three years][Y]](https://github.com/jgroeneveld/trial)   [![godoc][D]](https://godoc.org/github.com/jgroeneveld/trial)
- [envite](https://github.com/PerimeterX/envite) **star:5** Dev and testing environment management framework.   [![godoc][D]](https://godoc.org/github.com/PerimeterX/envite)
- [assay](https://github.com/tushariitr-19/assay) **star:4** Framework-agnostic evaluation library for testing Go agents and MCP servers with deterministic checks, CI-ready exit codes, and zero-code YAML-based testing.   [![godoc][D]](https://godoc.org/github.com/tushariitr-19/assay)
- [go-mutesting](https://github.com/jonbaldie/go-mutesting) **star:3** Mutation testing for Go with CI quality gates, coverage-aware MSI, baseline tracking, and git-diff filtering.   [![There was an update last month][G]](https://github.com/jonbaldie/go-mutesting)   [![godoc][D]](https://godoc.org/github.com/jonbaldie/go-mutesting)
- [go-mysql-test-container](https://github.com/arikama/go-mysql-test-container) **star:3** Golang MySQL testcontainer to help with MySQL integration testing.   [![It hasn't been updated in recent three years][Y]](https://github.com/arikama/go-mysql-test-container)   [![godoc][D]](https://godoc.org/github.com/arikama/go-mysql-test-container)
- [go-snaps](http://github.com/gkampitakis/go-snaps)  Jest-like snapshot testing in Golang.
- [httper](https://github.com/gustofarbi/httper) **star:1** CLI runner for JetBrains .http files with scripting, assertions, gRPC, and load testing.   [![godoc][D]](https://godoc.org/github.com/gustofarbi/httper)
- [omg.testingtools](https://github.com/dedalqq/omg.testingtools) **star:1** The simple library for change a values of private fields for testing.   [![It hasn't been updated in recent three years][Y]](https://github.com/dedalqq/omg.testingtools)   [![godoc][D]](https://godoc.org/github.com/dedalqq/omg.testingtools)

### Mock

- [mockery](https://github.com/vektra/mockery) **star:7153** Tool to generate Go interfaces.   [![godoc][D]](https://godoc.org/github.com/vektra/mockery)
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) **star:6562** Mock SQL driver for testing database interactions.   [![There was an update last month][G]](https://github.com/DATA-DOG/go-sqlmock)   [![godoc][D]](https://godoc.org/github.com/DATA-DOG/go-sqlmock)
- [gomock](https://github.com/uber-go/mock) **star:3385** Mocking framework for the Go programming language.   [![godoc][D]](https://godoc.org/github.com/uber-go/mock)
- [hoverfly](https://github.com/SpectoLabs/hoverfly) **star:2502** HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.   [![There was an update last month][G]](https://github.com/SpectoLabs/hoverfly)   [![godoc][D]](https://godoc.org/github.com/SpectoLabs/hoverfly)
- [moq](https://github.com/matryer/moq) **star:2202** Utility that generates a struct from any interface. The struct can be used in test code as a mock of the interface.   [![godoc][D]](https://godoc.org/github.com/matryer/moq)
- [moxie](https://lesiw.io/moxie)  Generate mock methods on embedded structs.
- [httpmock](https://github.com/jarcoal/httpmock) **star:2080** Easy mocking of HTTP responses from external resources.   [![godoc][D]](https://godoc.org/github.com/jarcoal/httpmock)
- [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) **star:1137** Tool for generating self-contained mock objects.   [![There was an update last month][G]](https://github.com/maxbrunsfeld/counterfeiter)   [![godoc][D]](https://godoc.org/github.com/maxbrunsfeld/counterfeiter)
- [minimock](https://github.com/gojuno/minimock) **star:753** Mock generator for Go interfaces.   [![There was an update last month][G]](https://github.com/gojuno/minimock)   [![godoc][D]](https://godoc.org/github.com/gojuno/minimock)
- [go-txdb](https://github.com/DATA-DOG/go-txdb) **star:751** Single transaction based database driver mainly for testing purposes.   [![godoc][D]](https://godoc.org/github.com/DATA-DOG/go-txdb)
- [pgxmock](https://github.com/pashagolub/pgxmock) **star:592** A mock library implementing [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx/).   [![godoc][D]](https://godoc.org/github.com/pashagolub/pgxmock)
- [xgo](https://github.com/xhd2015/xgo) **star:432** A general pureposed function mocking library.   [![There was an update last month][G]](https://github.com/xhd2015/xgo)   [![godoc][D]](https://godoc.org/github.com/xhd2015/xgo)   [![Contains Chinese documents][CN]](https://github.com/xhd2015/xgo)
- [govcr](https://github.com/seborama/govcr) **star:200** HTTP mock for Golang: record and replay HTTP interactions for offline testing.   [![godoc][D]](https://godoc.org/github.com/seborama/govcr)
- [go-localstack](https://github.com/elgohr/go-localstack) **star:88** Tool for using localstack in AWS testing.   [![godoc][D]](https://godoc.org/github.com/elgohr/go-localstack)
- [timex](https://github.com/cabify/timex) **star:71** A test-friendly replacement for the native `time` package.   [![It hasn't been updated in recent three years][Y]](https://github.com/cabify/timex)   [![godoc][D]](https://godoc.org/github.com/cabify/timex)
- [fabricator](https://github.com/Goldziher/fabricator) **star:29** Type-safe factories for generating mock and fake data in Go, inspired by factory_boy and interface-forge.   [![There was an update last month][G]](https://github.com/Goldziher/fabricator)   [![godoc][D]](https://godoc.org/github.com/Goldziher/fabricator)
- [genmock](https://gitlab.com/so_literate/genmock)  Go mocking system with code generator for building calls of the interface methods.
- [mockhttp](https://github.com/tv42/mockhttp) **star:23** Mock object for Go http.ResponseWriter.   [![It hasn't been updated in recent three years][Y]](https://github.com/tv42/mockhttp)   [![godoc][D]](https://godoc.org/github.com/tv42/mockhttp)
- [mooncake](https://github.com/GuilhermeCaruso/mooncake) **star:18** A simple way to generate mocks for multiple purposes.   [![It hasn't been updated in recent three years][Y]](https://github.com/GuilhermeCaruso/mooncake)   [![godoc][D]](https://godoc.org/github.com/GuilhermeCaruso/mooncake)
- [mockfs](https://github.com/balinomad/go-mockfs) **star:9** Mock filesystem for Go testing with error injection and latency simulation, built on `testing/fstest.MapFS`.   [![There was an update last month][G]](https://github.com/balinomad/go-mockfs)   [![godoc][D]](https://godoc.org/github.com/balinomad/go-mockfs)
- [gomock](https://github.com/vibridi/gomock) **star:7** CLI tool to generate typed and framework-agnostic interface mocks, with support for generics.   [![godoc][D]](https://godoc.org/github.com/vibridi/gomock)

### Fuzzing and delta-debugging/reducing/shrinking

- [go-fuzz](https://github.com/dvyukov/go-fuzz) **star:4852** Randomized testing system.   [![godoc][D]](https://godoc.org/github.com/dvyukov/go-fuzz)
- [Tavor](https://github.com/zimmski/tavor) **star:246** Generic fuzzing and delta-debugging framework.   [![It hasn't been updated in recent three years][Y]](https://github.com/zimmski/tavor)   [![godoc][D]](https://godoc.org/github.com/zimmski/tavor)

### Selenium and browser control tools

- [chromedp](https://github.com/knq/chromedp) **star:13178** a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.   [![There was an update last month][G]](https://github.com/knq/chromedp)   [![godoc][D]](https://godoc.org/github.com/knq/chromedp)
- [rod](https://github.com/go-rod/rod) **star:7001** A Devtools driver to make web automation and scraping easy.   [![godoc][D]](https://godoc.org/github.com/go-rod/rod)
- [playwright-go](https://github.com/mxschmitt/playwright-go) **star:3418** browser automation library to control Chromium, Firefox and WebKit with a single API.   [![godoc][D]](https://godoc.org/github.com/mxschmitt/playwright-go)
- [cdp](https://github.com/mafredri/cdp) **star:794** Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.   [![godoc][D]](https://godoc.org/github.com/mafredri/cdp)
- [selenosis](https://github.com/alcounit/selenosis) **star:86** Stateless Kubernetes-native hub that routes Selenium, Playwright, and MCP sessions to on-demand browser pods via custom resources.   [![godoc][D]](https://godoc.org/github.com/alcounit/selenosis)
- [bonk](https://github.com/joakimcarlsson/bonk) **star:13** Fast, stealth-first browser automation library using Chrome DevTools Protocol over WebSocket with no external dependencies.   [![There was an update last month][G]](https://github.com/joakimcarlsson/bonk)   [![godoc][D]](https://godoc.org/github.com/joakimcarlsson/bonk)

### Fail injection

- [failpoint](https://github.com/pingcap/failpoint) **star:890** An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.   [![godoc][D]](https://godoc.org/github.com/pingcap/failpoint)

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [go-humanize](https://github.com/dustin/go-humanize) **star:4809** Formatters for time, numbers, and memory size to human readable format.   [![godoc][D]](https://godoc.org/github.com/dustin/go-humanize)
- [sq](https://github.com/neilotoole/sq) **star:2535** Convert data from SQL databases or document formats like CSV or Excel into formats such as JSON, Excel, CSV, HTML, Markdown, XML, and YAML.   [![There was an update last month][G]](https://github.com/neilotoole/sq)   [![godoc][D]](https://godoc.org/github.com/neilotoole/sq)
- [gotabulate](https://github.com/bndr/gotabulate) **star:339** Easily pretty-print your tabular data with Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/bndr/gotabulate)   [![godoc][D]](https://godoc.org/github.com/bndr/gotabulate)
- [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) **star:88** Fixed-width text formatting (encoder/decoder with reflection).   [![godoc][D]](https://godoc.org/github.com/ianlopshire/go-fixedwidth)
- [align](https://github.com/Guitarbum722/align) **star:84** A general purpose application that aligns text.   [![It hasn't been updated in recent three years][Y]](https://github.com/Guitarbum722/align)   [![godoc][D]](https://godoc.org/github.com/Guitarbum722/align)
- [bytes](https://github.com/labstack/gommon/tree/master/bytes)  Formats and parses numeric byte values (10K, 2M, 3G, etc.).
- [address](https://github.com/bojanz/address) **star:83** Handles address representation, validation and formatting.   [![godoc][D]](https://godoc.org/github.com/bojanz/address)
- [textwrap](https://github.com/isbm/textwrap) **star:7** Wraps text at end of lines. Implementation of `textwrap` module from Python.   [![It hasn't been updated in recent three years][Y]](https://github.com/isbm/textwrap)   [![godoc][D]](https://godoc.org/github.com/isbm/textwrap)

### Markup Languages

- [blackfriday](https://github.com/russross/blackfriday) **star:5612** Markdown processor in Go.   [![godoc][D]](https://godoc.org/github.com/russross/blackfriday)
- [toml](https://github.com/BurntSushi/toml) **star:4983** TOML configuration format (encoder/decoder with reflection).   [![godoc][D]](https://godoc.org/github.com/BurntSushi/toml)
- [goldmark](https://github.com/yuin/goldmark) **star:4900** A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured.   [![There was an update last month][G]](https://github.com/yuin/goldmark)   [![godoc][D]](https://godoc.org/github.com/yuin/goldmark)
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) **star:3736** Convert HTML to Markdown. Even works with entire websites and can be extended through rules.   [![There was an update last month][G]](https://github.com/JohannesKaufmann/html-to-markdown)   [![godoc][D]](https://godoc.org/github.com/JohannesKaufmann/html-to-markdown)
- [go-toml](https://github.com/pelletier/go-toml) **star:1960** Go library for the TOML format with query support and handy cli tools.   [![There was an update last month][G]](https://github.com/pelletier/go-toml)   [![godoc][D]](https://godoc.org/github.com/pelletier/go-toml)
- [htmlquery](https://github.com/antchfx/htmlquery) **star:785** An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.   [![godoc][D]](https://godoc.org/github.com/antchfx/htmlquery)
- [mxj](https://github.com/clbanning/mxj) **star:629** Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.   [![godoc][D]](https://godoc.org/github.com/clbanning/mxj)
- [goq](https://github.com/andrewstuart/goq) **star:270** Declarative unmarshalling of HTML using struct tags with jQuery syntax (uses GoQuery).   [![It hasn't been updated in recent three years][Y]](https://github.com/andrewstuart/goq)   [![godoc][D]](https://godoc.org/github.com/andrewstuart/goq)
- [bafi](https://github.com/mmalcek/bafi) **star:116** Universal JSON, BSON, YAML, XML translator to ANY format using templates.   [![godoc][D]](https://godoc.org/github.com/mmalcek/bafi)
- [picoloom](https://github.com/alnah/picoloom) **star:83** Markdown-to-PDF converter with CLI and Go library APIs.   [![godoc][D]](https://godoc.org/github.com/alnah/picoloom)
- [go-output-format](https://github.com/drewstinnett/go-output-format) **star:18** Output go structures into multiple formats (YAML/JSON/etc) in your command line app.   [![godoc][D]](https://godoc.org/github.com/drewstinnett/go-output-format)
- [bbConvert](https://github.com/CalebQ42/bbConvert) **star:12** Converts bbCode to HTML that allows you to add support for custom bbCode tags.   [![godoc][D]](https://godoc.org/github.com/CalebQ42/bbConvert)
- [mdsmith](https://github.com/jeduden/mdsmith) **star:10** fast, auto-fixing Markdown linter and formatter. Checks style, readability, structure, and cross-file integrity.   [![There was an update last month][G]](https://github.com/jeduden/mdsmith)   [![godoc][D]](https://godoc.org/github.com/jeduden/mdsmith)
- [htmlyaml](https://github.com/nikolaydubina/htmlyaml) **star:6** Rich rendering of YAML as HTML in Go.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/htmlyaml)
- [htree](https://github.com/bobg/htree) **star:5** Traverse, navigate, filter, and otherwise process trees of [html.Node](https://pkg.go.dev/golang.org/x/net/html#Node) objects.   [![godoc][D]](https://godoc.org/github.com/bobg/htree)

### Parsers/Encoders/Decoders

- [sh](https://github.com/mvdan/sh) **star:8897** Shell parser and formatter.   [![godoc][D]](https://godoc.org/github.com/mvdan/sh)
- [gofeed](https://github.com/mmcdole/gofeed) **star:2849** Parse RSS and Atom feeds in Go.   [![There was an update last month][G]](https://github.com/mmcdole/gofeed)   [![godoc][D]](https://godoc.org/github.com/mmcdole/gofeed)
- [go-querystring](https://github.com/google/go-querystring) **star:2144** Go library for encoding structs into URL query parameters.   [![godoc][D]](https://godoc.org/github.com/google/go-querystring)
- [godump (goforj)](https://github.com/goforj/godump) **star:1747** Pretty-print Go structs with Laravel/Symfony-style dumps, full type info, colorized CLI output, cycle detection, and private field access.   [![There was an update last month][G]](https://github.com/goforj/godump)   [![godoc][D]](https://godoc.org/github.com/goforj/godump)
- [when](https://github.com/olebedev/when) **star:1462** Natural EN and RU language date/time parser with pluggable rules.   [![godoc][D]](https://godoc.org/github.com/olebedev/when)
- [commonregex](https://github.com/mingrammer/commonregex) **star:898** A collection of common regular expressions for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/mingrammer/commonregex)   [![godoc][D]](https://godoc.org/github.com/mingrammer/commonregex)
- [gographviz](https://github.com/awalterschulze/gographviz) **star:567** Parses the Graphviz DOT language.   [![It hasn't been updated in recent three years][Y]](https://github.com/awalterschulze/gographviz)   [![godoc][D]](https://godoc.org/github.com/awalterschulze/gographviz)
- [go-nmea](https://github.com/adrianmo/go-nmea) **star:264** NMEA parser library for the Go language.   [![godoc][D]](https://godoc.org/github.com/adrianmo/go-nmea)
- [godump](https://github.com/yassinebenaid/godump) **star:223** Pretty print any GO variable with ease, an alternative to Go's `fmt.Printf("%#v")`.   [![godoc][D]](https://godoc.org/github.com/yassinebenaid/godump)
- [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) **star:155** Editorconfig file parser and manipulator for Go.   [![There was an update last month][G]](https://github.com/editorconfig/editorconfig-core-go)   [![godoc][D]](https://godoc.org/github.com/editorconfig/editorconfig-core-go)
- [tokenizer](https://github.com/bzick/tokenizer) **star:141** Parse any string, slice or infinite buffer to any tokens.   [![There was an update last month][G]](https://github.com/bzick/tokenizer)   [![godoc][D]](https://godoc.org/github.com/bzick/tokenizer)
- [go-vcard](https://github.com/emersion/go-vcard) **star:128** Parse and format vCard.   [![godoc][D]](https://godoc.org/github.com/emersion/go-vcard)
- [did](https://github.com/ockam-network/did) **star:95** DID (Decentralized Identifiers) Parser and Stringer in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ockam-network/did)   [![godoc][D]](https://godoc.org/github.com/ockam-network/did)
- [vdf](https://github.com/andygrunwald/vdf) **star:66** A Lexer and Parser for Valves Data Format (known as vdf) written in Go.   [![godoc][D]](https://godoc.org/github.com/andygrunwald/vdf)
- [allot](https://github.com/sbstjn/allot) **star:60** Placeholder and wildcard text parsing for CLI tools and bots.   [![It hasn't been updated in recent three years][Y]](https://github.com/sbstjn/allot)   [![godoc][D]](https://godoc.org/github.com/sbstjn/allot)
- [normalize](https://github.com/avito-tech/normalize) **star:53** Sanitize, normalize and compare fuzzy text.   [![It hasn't been updated in recent three years][Y]](https://github.com/avito-tech/normalize)   [![godoc][D]](https://godoc.org/github.com/avito-tech/normalize)
- [gonameparts](https://github.com/polera/gonameparts) **star:43** Parses human names into individual name parts.   [![godoc][D]](https://godoc.org/github.com/polera/gonameparts)
- [go-fasttld](https://github.com/elliotwutingfeng/go-fasttld) **star:39** High performance effective top level domains (eTLD) extraction module.   [![There was an update last month][G]](https://github.com/elliotwutingfeng/go-fasttld)   [![godoc][D]](https://godoc.org/github.com/elliotwutingfeng/go-fasttld)
- [xj2go](https://github.com/stackerzzq/xj2go) **star:36** Convert xml or json to go struct.   [![It hasn't been updated in recent three years][Y]](https://github.com/stackerzzq/xj2go)   [![godoc][D]](https://godoc.org/github.com/stackerzzq/xj2go)
- [codetree](https://github.com/aerogo/codetree) **star:25** Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.   [![It hasn't been updated in recent three years][Y]](https://github.com/aerogo/codetree)   [![godoc][D]](https://godoc.org/github.com/aerogo/codetree)
- [doi](https://github.com/hscells/doi) **star:11** Document object identifier (doi) parser in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/hscells/doi)   [![godoc][D]](https://godoc.org/github.com/hscells/doi)
- [parseargs-go](https://github.com/nproc/parseargs-go) **star:10** string argument parser that understands quotes and backslashes.   [![It hasn't been updated in recent three years][Y]](https://github.com/nproc/parseargs-go)   [![godoc][D]](https://godoc.org/github.com/nproc/parseargs-go)
- [ltsv](https://github.com/Wing924/ltsv) **star:9** High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go.   [![godoc][D]](https://godoc.org/github.com/Wing924/ltsv)
- [prattle](https://github.com/askeladdk/prattle) **star:9** Scan and parse LL(1) grammars simply and efficiently.   [![godoc][D]](https://godoc.org/github.com/askeladdk/prattle)

### Regular Expressions

- [coregex](https://github.com/coregx/coregex) **star:241** Production regex engine with Rust regex-crate architecture: multi-engine DFA/NFA, SIMD prefilters, drop-in stdlib replacement.   [![godoc][D]](https://godoc.org/github.com/coregx/coregex)
- [rex](https://github.com/hedhyw/rex) **star:214** Regular expressions builder.   [![godoc][D]](https://godoc.org/github.com/hedhyw/rex)
- [regroup](https://github.com/oriser/regroup) **star:150** Match regex expression named groups into go struct using struct tags and automatic parsing.   [![godoc][D]](https://godoc.org/github.com/oriser/regroup)
- [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) **star:103** Simple and lightweight wildcard pattern matching.   [![godoc][D]](https://godoc.org/github.com/IGLOU-EU/go-wildcard)
- [goregen](https://github.com/zach-klippenstein/goregen) **star:92** Library for generating random strings from regular expressions.   [![It hasn't been updated in recent three years][Y]](https://github.com/zach-klippenstein/goregen)   [![godoc][D]](https://godoc.org/github.com/zach-klippenstein/goregen)
- [genex](https://github.com/alixaxel/genex) **star:77** Count and expand Regular Expressions into all matching Strings.   [![It hasn't been updated in recent three years][Y]](https://github.com/alixaxel/genex)   [![godoc][D]](https://godoc.org/github.com/alixaxel/genex)

### Sanitation

- [bluemonday](https://github.com/microcosm-cc/bluemonday) **star:3696** HTML Sanitizer.   [![godoc][D]](https://godoc.org/github.com/microcosm-cc/bluemonday)
- [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) **star:71** A sanitization-based swear filter for Go.   [![godoc][D]](https://godoc.org/github.com/JoshuaDoes/gofuckyourself)

### Scrapers

- [colly](https://github.com/asciimoo/colly) **star:25378** Fast and Elegant Scraping Framework for Gophers.   [![godoc][D]](https://godoc.org/github.com/asciimoo/colly)
- [GoQuery](https://github.com/PuerkitoBio/goquery) **star:14972** GoQuery brings a syntax and a set of features similar to jQuery to the Go language.   [![There was an update last month][G]](https://github.com/PuerkitoBio/goquery)   [![godoc][D]](https://godoc.org/github.com/PuerkitoBio/goquery)
- [xurls](https://github.com/mvdan/xurls) **star:1267** Extract urls from text.   [![godoc][D]](https://godoc.org/github.com/mvdan/xurls)
- [dataflowkit](https://github.com/slotix/dataflowkit) **star:716** Web scraping Framework to turn websites into structured data.   [![It hasn't been updated in recent three years][Y]](https://github.com/slotix/dataflowkit)   [![godoc][D]](https://godoc.org/github.com/slotix/dataflowkit)
- [pagser](https://github.com/foolin/pagser) **star:111** Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler.   [![godoc][D]](https://godoc.org/github.com/foolin/pagser)
- [Tagify](https://github.com/zoomio/tagify) **star:39** Produces a set of tags from given source.
- [go-recipe](https://github.com/kkyr/go-recipe) **star:32** A package for scraping recipes from websites.   [![It hasn't been updated in recent three years][Y]](https://github.com/kkyr/go-recipe)   [![godoc][D]](https://godoc.org/github.com/kkyr/go-recipe)
- [walker](https://github.com/cyucelen/walker) **star:14** Seamlessly fetch paginated data from any source. Simple and high performance API scraping included.   [![It hasn't been updated in recent three years][Y]](https://github.com/cyucelen/walker)   [![godoc][D]](https://godoc.org/github.com/cyucelen/walker)
- [go-sitemap-parser](https://github.com/aafeher/go-sitemap-parser) **star:7** Go language library for parsing Sitemaps.   [![There was an update last month][G]](https://github.com/aafeher/go-sitemap-parser)   [![godoc][D]](https://godoc.org/github.com/aafeher/go-sitemap-parser)

### RSS

- [podcast](https://github.com/eduncan911/podcast) **star:140** iTunes Compliant and RSS 2.0 Podcast Generator in Golang   [![It hasn't been updated in recent three years][Y]](https://github.com/eduncan911/podcast)   [![godoc][D]](https://godoc.org/github.com/eduncan911/podcast)

### Utility/Miscellaneous

- [lancet](https://github.com/duke-git/lancet) **star:5297** A comprehensive, Lodash-like utility library for Go   [![godoc][D]](https://godoc.org/github.com/duke-git/lancet)   [![Contains Chinese documents][CN]](https://github.com/duke-git/lancet)
- [w2vgrep](https://github.com/arunsupe/semantic-grep) **star:1239** A semantic grep tool using word embeddings to find semantically similar matches. For example, searching for "death" will find "dead", "killing", "murder".   [![godoc][D]](https://godoc.org/github.com/arunsupe/semantic-grep)
- [go-runewidth](https://github.com/mattn/go-runewidth) **star:713** Functions to get fixed width of the character or string.   [![godoc][D]](https://godoc.org/github.com/mattn/go-runewidth)
- [radix](https://github.com/yourbasic/radix) **star:197** Fast string sorting algorithm.   [![It hasn't been updated in recent three years][Y]](https://github.com/yourbasic/radix)   [![godoc][D]](https://godoc.org/github.com/yourbasic/radix)
- [petrovich](https://github.com/striker2000/petrovich) **star:51** Petrovich is the library which inflects Russian names to given grammatical case.   [![godoc][D]](https://godoc.org/github.com/striker2000/petrovich)
- [ahocorasick](https://github.com/coregx/ahocorasick) **star:25** High-performance Aho-Corasick multi-pattern string matching with DFA compilation and SIMD prefilter, up to 7 GB/s throughput (part of [coregx](https://github.com/coregx) ecosystem).   [![godoc][D]](https://godoc.org/github.com/coregx/ahocorasick)
- [kace](https://github.com/codemodus/kace) **star:22** Common case conversions covering common initialisms.   [![It hasn't been updated in recent three years][Y]](https://github.com/codemodus/kace)   [![godoc][D]](https://godoc.org/github.com/codemodus/kace)
- [TySug](https://github.com/Dynom/TySug) **star:20** Alternative suggestions with respect to keyboard layouts.   [![godoc][D]](https://godoc.org/github.com/Dynom/TySug)
- [uniwidth](https://github.com/unilibs/uniwidth) **star:8** High-performance Unicode character width calculation with SWAR optimization, O(1) lookup tables, and ZWJ emoji support.   [![godoc][D]](https://godoc.org/github.com/unilibs/uniwidth)

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [github](https://github.com/google/go-github) **star:11271** Go library for accessing the GitHub REST API v3.   [![There was an update last month][G]](https://github.com/google/go-github)   [![godoc][D]](https://godoc.org/github.com/google/go-github)
- [go-openai](https://github.com/sashabaranov/go-openai) **star:10714** OpenAI ChatGPT, DALL·E, Whisper API library for Go.   [![godoc][D]](https://godoc.org/github.com/sashabaranov/go-openai)
- [discordgo](https://github.com/bwmarrin/discordgo) **star:5957** Go bindings for the Discord Chat API.   [![godoc][D]](https://godoc.org/github.com/bwmarrin/discordgo)
- [slack](https://github.com/slack-go/slack) **star:4952** Slack API in Go.   [![There was an update last month][G]](https://github.com/slack-go/slack)   [![godoc][D]](https://godoc.org/github.com/slack-go/slack)
- [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) **star:4474** Google Cloud APIs Go Client Library.   [![There was an update last month][G]](https://github.com/GoogleCloudPlatform/gcloud-golang)   [![godoc][D]](https://godoc.org/github.com/GoogleCloudPlatform/gcloud-golang)
- [google](https://github.com/google/google-api-go-client) **star:4460** Auto-generated Google APIs for Go.   [![There was an update last month][G]](https://github.com/google/google-api-go-client)   [![godoc][D]](https://godoc.org/github.com/google/google-api-go-client)
- [aws-sdk-go](https://github.com/aws/aws-sdk-go-v2) **star:3609** The official AWS SDK for the Go programming language.   [![There was an update last month][G]](https://github.com/aws/aws-sdk-go-v2)   [![godoc][D]](https://godoc.org/github.com/aws/aws-sdk-go-v2)
- [minio-go](https://github.com/minio/minio-go) **star:2974** Minio Go Library for Amazon S3 compatible cloud storage.   [![godoc][D]](https://godoc.org/github.com/minio/minio-go)
- [stripe](https://github.com/stripe/stripe-go) **star:2602** Go client for the Stripe API.   [![There was an update last month][G]](https://github.com/stripe/stripe-go)   [![godoc][D]](https://godoc.org/github.com/stripe/stripe-go)
- [go-jira](https://github.com/andygrunwald/go-jira) **star:1613** Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)   [![godoc][D]](https://godoc.org/github.com/andygrunwald/go-jira)
- [facebook](https://github.com/huandu/facebook) **star:1473** Go Library that supports the Facebook Graph API.   [![godoc][D]](https://godoc.org/github.com/huandu/facebook)
- [githubql](https://github.com/shurcooL/githubql) **star:1193** Go library for accessing the GitHub GraphQL API v4.   [![godoc][D]](https://godoc.org/github.com/shurcooL/githubql)
- [anaconda](https://github.com/ChimeraCoder/anaconda) **star:1137** Go client library for the Twitter 1.1 API.   [![godoc][D]](https://godoc.org/github.com/ChimeraCoder/anaconda)
- [webhooks](https://github.com/go-playground/webhooks) **star:1027** Webhook receiver for GitHub and Bitbucket.   [![godoc][D]](https://godoc.org/github.com/go-playground/webhooks)
- [libopenapi](https://github.com/pb33f/libopenapi) **star:857** Parse, validate, and work with OpenAPI, Swagger, Overlays, and Arazzo specifications.   [![There was an update last month][G]](https://github.com/pb33f/libopenapi)   [![godoc][D]](https://godoc.org/github.com/pb33f/libopenapi)
- [paypal](https://github.com/logpacker/PayPal-Go-SDK) **star:777** Wrapper for PayPal payment API.   [![godoc][D]](https://godoc.org/github.com/logpacker/PayPal-Go-SDK)
- [geo-golang](https://github.com/codingsince1985/geo-golang) **star:546** Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://developer.mapquest.com/documentation/api/geocoding/), [Nominatim](https://nominatim.org/release-docs/latest/api/Overview/), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.   [![godoc][D]](https://godoc.org/github.com/codingsince1985/geo-golang)
- [lark](https://github.com/chyroc/lark) **star:475** [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback.   [![godoc][D]](https://godoc.org/github.com/chyroc/lark)   [![Contains Chinese documents][CN]](https://github.com/chyroc/lark)
- [openaigo](https://github.com/otiai10/openaigo) **star:301** OpenAI GPT3/GPT3.5 ChatGPT API client library for Go.   [![godoc][D]](https://godoc.org/github.com/otiai10/openaigo)   [![Archived][Archived]](https://github.com/otiai10/openaigo)
- [ethrpc](https://github.com/onrik/ethrpc) **star:271** Go bindings for Ethereum JSON RPC API.   [![godoc][D]](https://godoc.org/github.com/onrik/ethrpc)
- [go-lark](https://github.com/go-lark/lark) **star:247** An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform.   [![godoc][D]](https://godoc.org/github.com/go-lark/lark)   [![Contains Chinese documents][CN]](https://github.com/go-lark/lark)
- [Trello](https://github.com/adlio/trello) **star:228** Go wrapper for the Trello API.   [![godoc][D]](https://godoc.org/github.com/adlio/trello)
- [go-atlassian](https://github.com/ctreminiom/go-atlassian) **star:215** Go library for accessing the [Atlassian Cloud](https://www.atlassian.com/enterprise/cloud) services (Jira, Jira Service Management, Jira Agile, Confluence, Admin Cloud)   [![godoc][D]](https://godoc.org/github.com/ctreminiom/go-atlassian)
- [simples3](https://github.com/rhnvrm/simples3) **star:205** Simple no frills AWS S3 Library using REST with V4 Signing written in Go.   [![godoc][D]](https://godoc.org/github.com/rhnvrm/simples3)
- [go-marathon](https://github.com/gambol99/go-marathon) **star:199** Go library for interacting with Mesosphere's Marathon PAAS.   [![It hasn't been updated in recent three years][Y]](https://github.com/gambol99/go-marathon)   [![godoc][D]](https://godoc.org/github.com/gambol99/go-marathon)
- [wit-go](https://github.com/wit-ai/wit-go) **star:170** Go client for wit.ai HTTP API.   [![godoc][D]](https://godoc.org/github.com/wit-ai/wit-go)
- [gosip](https://github.com/koltyakov/gosip) **star:169** Client library for SharePoint.   [![godoc][D]](https://godoc.org/github.com/koltyakov/gosip)
- [golang-tmdb](https://github.com/cyruzin/golang-tmdb) **star:163** Golang wrapper for The Movie Database API v3.   [![godoc][D]](https://godoc.org/github.com/cyruzin/golang-tmdb)
- [pushover](https://github.com/gregdel/pushover) **star:157** Go wrapper for the Pushover API.   [![godoc][D]](https://godoc.org/github.com/gregdel/pushover)
- [go-trending](https://github.com/andygrunwald/go-trending) **star:147** Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.   [![godoc][D]](https://godoc.org/github.com/andygrunwald/go-trending)
- [Medium](https://github.com/Medium/medium-sdk-go) **star:142** Golang SDK for Medium's OAuth2 API.   [![It hasn't been updated in recent three years][Y]](https://github.com/Medium/medium-sdk-go)   [![godoc][D]](https://godoc.org/github.com/Medium/medium-sdk-go)
- [gostorm](https://github.com/jsgilmore/gostorm) **star:129** GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.   [![It hasn't been updated in recent three years][Y]](https://github.com/jsgilmore/gostorm)   [![godoc][D]](https://godoc.org/github.com/jsgilmore/gostorm)
- [spec](https://github.com/oaswrap/spec) **star:115** Lightweight OpenAPI 3.x builder supporting static generation and popular frameworks like chi, echo, gin, fiber, mux and more.   [![godoc][D]](https://godoc.org/github.com/oaswrap/spec)
- [disgo](https://github.com/switchupcb/disgo) **star:114** Go API Wrapper for the Discord API.   [![godoc][D]](https://godoc.org/github.com/switchupcb/disgo)
- [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) **star:111** A golang package to communicate with HipChat over XMPP.   [![It hasn't been updated in recent three years][Y]](https://github.com/daneharrigan/hipchat)   [![godoc][D]](https://godoc.org/github.com/daneharrigan/hipchat)
- [go-gerrit](https://github.com/andygrunwald/go-gerrit) **star:106** Go client library for [Gerrit Code Review](https://www.gerritcodereview.com/).   [![godoc][D]](https://godoc.org/github.com/andygrunwald/go-gerrit)
- [hipchat](https://github.com/andybons/hipchat) **star:104** This project implements a golang client library for the Hipchat API.   [![It hasn't been updated in recent three years][Y]](https://github.com/andybons/hipchat)   [![godoc][D]](https://godoc.org/github.com/andybons/hipchat)
- [go-redoc](https://github.com/mvrilo/go-redoc) **star:95** Embedded OpenAPI/Swagger documentation ui for Go using [ReDoc](https://redocly.com/).   [![godoc][D]](https://godoc.org/github.com/mvrilo/go-redoc)
- [cachet](https://github.com/andygrunwald/cachet) **star:90** Go client library for [Cachet (open source status page system)](https://cachethq.io/).   [![It hasn't been updated in recent three years][Y]](https://github.com/andygrunwald/cachet)   [![godoc][D]](https://godoc.org/github.com/andygrunwald/cachet)
- [GoFreeDB](https://github.com/FreeLeh/GoFreeDB) **star:90** Golang library providing common and simple database abstractions on top of Google Sheets.   [![godoc][D]](https://godoc.org/github.com/FreeLeh/GoFreeDB)
- [gogtrends](https://github.com/groovili/gogtrends) **star:89** Google Trends Unofficial API.   [![It hasn't been updated in recent three years][Y]](https://github.com/groovili/gogtrends)   [![godoc][D]](https://godoc.org/github.com/groovili/gogtrends)
- [airtable](https://github.com/mehanizm/airtable) **star:86** Go client library for the [Airtable API](https://airtable.com/api).   [![godoc][D]](https://godoc.org/github.com/mehanizm/airtable)
- [go-postman-collection](https://github.com/rbretecher/go-postman-collection) **star:86** Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).   [![godoc][D]](https://godoc.org/github.com/rbretecher/go-postman-collection)
- [igdb](https://github.com/Henry-Sarabia/igdb) **star:84** Go client for the [Internet Game Database API](https://api.igdb.com/).   [![It hasn't been updated in recent three years][Y]](https://github.com/Henry-Sarabia/igdb)   [![godoc][D]](https://godoc.org/github.com/Henry-Sarabia/igdb)
- [go-unsplash](https://github.com/hbagdi/go-unsplash) **star:79** Go client library for the [Unsplash.com](https://unsplash.com) API.   [![godoc][D]](https://godoc.org/github.com/hbagdi/go-unsplash)
- [ynab](https://github.com/brunomvsouza/ynab.go) **star:78** Go wrapper for the YNAB API.   [![There was an update last month][G]](https://github.com/brunomvsouza/ynab.go)   [![godoc][D]](https://godoc.org/github.com/brunomvsouza/ynab.go)
- [circleci](https://github.com/jszwedko/go-circleci) **star:65** Go client library for interacting with CircleCI's API.   [![godoc][D]](https://godoc.org/github.com/jszwedko/go-circleci)
- [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) **star:63** Go MusicBrainz WS2 client library.   [![It hasn't been updated in recent three years][Y]](https://github.com/michiwend/gomusicbrainz)   [![godoc][D]](https://godoc.org/github.com/michiwend/gomusicbrainz)
- [mixpanel](https://github.com/dukex/mixpanel) **star:61** Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.   [![godoc][D]](https://godoc.org/github.com/dukex/mixpanel)   [![Archived][Archived]](https://github.com/dukex/mixpanel)
- [uptimerobot](https://github.com/bitfield/uptimerobot) **star:60** Go wrapper and command-line client for the Uptime Robot v2 API.   [![It hasn't been updated in recent three years][Y]](https://github.com/bitfield/uptimerobot)   [![godoc][D]](https://godoc.org/github.com/bitfield/uptimerobot)
- [go-salesforce](https://github.com/k-capehart/go-salesforce) **star:56** Go client library for interacting with the [Salesforce REST API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_list.htm).   [![There was an update last month][G]](https://github.com/k-capehart/go-salesforce)   [![godoc][D]](https://godoc.org/github.com/k-capehart/go-salesforce)
- [megos](https://github.com/andygrunwald/megos) **star:54** Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster.   [![It hasn't been updated in recent three years][Y]](https://github.com/andygrunwald/megos)   [![godoc][D]](https://godoc.org/github.com/andygrunwald/megos)
- [spotify](https://github.com/rapito/go-spotify) **star:53** Go Library to access Spotify WEB API.   [![godoc][D]](https://godoc.org/github.com/rapito/go-spotify)
- [fcm](https://github.com/maddevsio/fcm) **star:52** Go library for Firebase Cloud Messaging.   [![It hasn't been updated in recent three years][Y]](https://github.com/maddevsio/fcm)   [![godoc][D]](https://godoc.org/github.com/maddevsio/fcm)
- [gads](https://github.com/emiddleton/gads) **star:51** Google Adwords Unofficial API.   [![It hasn't been updated in recent three years][Y]](https://github.com/emiddleton/gads)   [![godoc][D]](https://godoc.org/github.com/emiddleton/gads)
- [go-xkcd](https://github.com/nishanths/go-xkcd) **star:51** Go client for the xkcd API.   [![It hasn't been updated in recent three years][Y]](https://github.com/nishanths/go-xkcd)   [![godoc][D]](https://godoc.org/github.com/nishanths/go-xkcd)
- [go-yapla](https://gitlab.com/adrienK/go-yapla)  Go client library for the Yapla v2.0 API.
- [swag](https://github.com/zc2638/swag) **star:49** No comments, simple go wrapper to create swagger 2.0 compatible APIs. Support most routing frameworks, such as built-in, gin, chi, mux, echo, httprouter, fasthttp and more.   [![godoc][D]](https://godoc.org/github.com/zc2638/swag)   [![Contains Chinese documents][CN]](https://github.com/zc2638/swag)
- [patreon-go](https://github.com/mxpv/patreon-go) **star:46** Go library for Patreon API.   [![It hasn't been updated in recent three years][Y]](https://github.com/mxpv/patreon-go)   [![godoc][D]](https://godoc.org/github.com/mxpv/patreon-go)
- [go-myanimelist](https://github.com/nstratos/go-myanimelist) **star:43** Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2).   [![godoc][D]](https://godoc.org/github.com/nstratos/go-myanimelist)
- [golyrics](https://github.com/mamal72/golyrics) **star:40** Golyrics is a Go library to fetch music lyrics data from the Wikia website.   [![It hasn't been updated in recent three years][Y]](https://github.com/mamal72/golyrics)   [![godoc][D]](https://godoc.org/github.com/mamal72/golyrics)
- [steam](https://github.com/sostronk/go-steam) **star:33** Go Library to interact with Steam game servers.   [![godoc][D]](https://godoc.org/github.com/sostronk/go-steam)
- [lastpass-go](https://github.com/ansd/lastpass-go) **star:32** Go client library for the [LastPass](https://www.lastpass.com/) API.   [![It hasn't been updated in recent three years][Y]](https://github.com/ansd/lastpass-go)   [![godoc][D]](https://godoc.org/github.com/ansd/lastpass-go)
- [gcm](https://github.com/Aorioli/gcm) **star:31** Go library for Google Cloud Messaging.   [![It hasn't been updated in recent three years][Y]](https://github.com/Aorioli/gcm)   [![godoc][D]](https://godoc.org/github.com/Aorioli/gcm)
- [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) **star:27** Go client library for interacting with Coinpaprika's API.   [![godoc][D]](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
- [jokeapi-go](https://github.com/icelain/jokeapi) **star:27** Go client for [JokeAPI](https://sv443.net/jokeapi/v2/).   [![godoc][D]](https://godoc.org/github.com/icelain/jokeapi)
- [device-check-go](https://github.com/rinchsan/device-check-go) **star:26** Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1.   [![There was an update last month][G]](https://github.com/rinchsan/device-check-go)   [![godoc][D]](https://godoc.org/github.com/rinchsan/device-check-go)
- [shopify](https://github.com/rapito/go-shopify) **star:25** Go Library to make CRUD request to the Shopify API.   [![It hasn't been updated in recent three years][Y]](https://github.com/rapito/go-shopify)   [![godoc][D]](https://godoc.org/github.com/rapito/go-shopify)
- [aws-encryption-sdk-go](https://github.com/chainifynet/aws-encryption-sdk-go) **star:23** Unofficial Go SDK implementation of the [AWS Encryption SDK](https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/index.html).   [![godoc][D]](https://godoc.org/github.com/chainifynet/aws-encryption-sdk-go)
- [go-imgur](https://github.com/koffeinsource/go-imgur) **star:23** Go client library for [imgur](https://imgur.com)   [![godoc][D]](https://godoc.org/github.com/koffeinsource/go-imgur)
- [goami2](https://github.com/staskobzar/goami2) **star:22** AMI v2 library for Asterisk PBX.   [![godoc][D]](https://godoc.org/github.com/staskobzar/goami2)
- [textbelt](https://github.com/dietsche/textbelt) **star:21** Go client for the textbelt.com txt messaging API.   [![It hasn't been updated in recent three years][Y]](https://github.com/dietsche/textbelt)   [![godoc][D]](https://godoc.org/github.com/dietsche/textbelt)
- [brewerydb](https://github.com/naegelejd/brewerydb) **star:20** Go library for accessing the BreweryDB API.   [![It hasn't been updated in recent three years][Y]](https://github.com/naegelejd/brewerydb)   [![godoc][D]](https://godoc.org/github.com/naegelejd/brewerydb)
- [go-aws-news](https://github.com/circa10a/go-aws-news) **star:19** Go application and library to fetch what's new from AWS.   [![godoc][D]](https://godoc.org/github.com/circa10a/go-aws-news)
- [go-openproject](https://github.com/manuelbcd/go-openproject) **star:19** Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API.   [![It hasn't been updated in recent three years][Y]](https://github.com/manuelbcd/go-openproject)   [![godoc][D]](https://godoc.org/github.com/manuelbcd/go-openproject)
- [codeship-go](https://github.com/codeship/codeship-go) **star:18** Go client library for interacting with Codeship's API v2.   [![It hasn't been updated in recent three years][Y]](https://github.com/codeship/codeship-go)   [![godoc][D]](https://godoc.org/github.com/codeship/codeship-go)
- [bqwriter](https://github.com/OTA-Insight/bqwriter) **star:16** High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout.   [![godoc][D]](https://godoc.org/github.com/OTA-Insight/bqwriter)
- [go-hacknews](https://github.com/PaulRosset/go-hacknews) **star:16** Tiny Go client for HackerNews API.   [![It hasn't been updated in recent three years][Y]](https://github.com/PaulRosset/go-hacknews)   [![godoc][D]](https://godoc.org/github.com/PaulRosset/go-hacknews)
- [gopaapi5](https://github.com/utekaravinash/gopaapi5) **star:16** Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).   [![It hasn't been updated in recent three years][Y]](https://github.com/utekaravinash/gopaapi5)   [![godoc][D]](https://godoc.org/github.com/utekaravinash/gopaapi5)
- [google-analytics](https://github.com/chonthu/go-google-analytics) **star:15** Simple wrapper for easy google analytics reporting.   [![It hasn't been updated in recent three years][Y]](https://github.com/chonthu/go-google-analytics)   [![godoc][D]](https://godoc.org/github.com/chonthu/go-google-analytics)
- [ip2location-io-go](https://github.com/ip2location/ip2location-io-go) **star:15** Go wrapper for the IP2Location.io API [IP2Location.io](https://www.ip2location.io/).   [![godoc][D]](https://godoc.org/github.com/ip2location/ip2location-io-go)
- [go-here](https://github.com/abdullahselek/go-here) **star:13** Go client library around the HERE location based APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/abdullahselek/go-here)   [![godoc][D]](https://godoc.org/github.com/abdullahselek/go-here)
- [go-hibp](https://github.com/wneessen/go-hibp) **star:13** Simple Go binding to the "Have I Been Pwned" APIs.   [![godoc][D]](https://godoc.org/github.com/wneessen/go-hibp)
- [gomalshare](https://github.com/MonaxGT/gomalshare) **star:13** Go library MalShare API [malshare.com](https://www.malshare.com/)   [![It hasn't been updated in recent three years][Y]](https://github.com/MonaxGT/gomalshare)   [![godoc][D]](https://godoc.org/github.com/MonaxGT/gomalshare)
- [go-sophos](https://github.com/esurdam/go-sophos) **star:12** Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.   [![It hasn't been updated in recent three years][Y]](https://github.com/esurdam/go-sophos)   [![godoc][D]](https://godoc.org/github.com/esurdam/go-sophos)
- [goagi](https://github.com/staskobzar/goagi) **star:12** Go library to build Asterisk PBX agi/fastagi applications.   [![godoc][D]](https://godoc.org/github.com/staskobzar/goagi)
- [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) **star:12** Go library for the [RAWG Video Games Database](https://rawg.io/) API   [![It hasn't been updated in recent three years][Y]](https://github.com/dimuska139/rawg-sdk-go)   [![godoc][D]](https://godoc.org/github.com/dimuska139/rawg-sdk-go)
- [go-swagger-ui](https://github.com/esurdam/go-swagger-ui) **star:11** Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json.   [![It hasn't been updated in recent three years][Y]](https://github.com/esurdam/go-swagger-ui)
- [go-telegraph](https://gitlab.com/toby3d/telegraph)  Telegraph publishing platform API client.
- [smite](https://github.com/sergiotapia/smitego) **star:11** Go package to wraps access to the Smite game API.   [![It hasn't been updated in recent three years][Y]](https://github.com/sergiotapia/smitego)   [![godoc][D]](https://godoc.org/github.com/sergiotapia/smitego)
- [libgoffi](https://github.com/clevabit/libgoffi) **star:10** Library adapter toolbox for native [libffi](https://sourceware.org/libffi/) integration   [![It hasn't been updated in recent three years][Y]](https://github.com/clevabit/libgoffi)   [![godoc][D]](https://godoc.org/github.com/clevabit/libgoffi)
- [go-chronos](https://github.com/axelspringer/go-chronos) **star:8** Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler   [![It hasn't been updated in recent three years][Y]](https://github.com/axelspringer/go-chronos)   [![godoc][D]](https://godoc.org/github.com/axelspringer/go-chronos)   [![Archived][Archived]](https://github.com/axelspringer/go-chronos)
- [threads-go](https://github.com/tirthpatell/threads-go) **star:8** Go client library for the Meta Threads API with OAuth 2.0, rate limiting, and type-safe error handling.   [![There was an update last month][G]](https://github.com/tirthpatell/threads-go)   [![godoc][D]](https://godoc.org/github.com/tirthpatell/threads-go)
- [tumblr](https://github.com/mattcunningham/gumblr) **star:8** Go wrapper for the Tumblr v2 API.   [![It hasn't been updated in recent three years][Y]](https://github.com/mattcunningham/gumblr)   [![godoc][D]](https://godoc.org/github.com/mattcunningham/gumblr)
- [newsapi-go](https://github.com/jellydator/newsapi-go) **star:7** Go client for [NewsAPI](https://newsapi.org/).   [![godoc][D]](https://godoc.org/github.com/jellydator/newsapi-go)
- [zooz](https://github.com/gojuno/go-zooz) **star:7** Go client for the Zooz API.   [![godoc][D]](https://godoc.org/github.com/gojuno/go-zooz)
- [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) **star:6** Unofficial Golang SDK for AppStore Connect API.   [![godoc][D]](https://godoc.org/github.com/Kachit/appstore-sdk-go)
- [dusupay-sdk-go](https://github.com/Kachit/dusupay-sdk-go) **star:4** Unofficial Dusupay payment gateway API Client for Go   [![godoc][D]](https://godoc.org/github.com/Kachit/dusupay-sdk-go)
- [gopensky](https://github.com/navidys/gopensky) **star:4** Go client implementation for [OpenSKY Network](https://opensky-network.org/) live's API (airspace ADS-B and Mode S data).   [![godoc][D]](https://godoc.org/github.com/navidys/gopensky)
- [go-restcountries](https://github.com/chriscross0/go-restcountries) **star:3** Go library for the [REST Countries API](https://countrylayer.com/).   [![It hasn't been updated in recent three years][Y]](https://github.com/chriscross0/go-restcountries)   [![godoc][D]](https://godoc.org/github.com/chriscross0/go-restcountries)
- [fasapay-sdk-go](https://github.com/Kachit/fasapay-sdk-go) **star:2** Unofficial Fasapay payment gateway XML API Client for Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/Kachit/fasapay-sdk-go)   [![godoc][D]](https://godoc.org/github.com/Kachit/fasapay-sdk-go)
- [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) **star:2** The Playlyfe Rest API Go SDK.   [![It hasn't been updated in recent three years][Y]](https://github.com/playlyfe/playlyfe-go-sdk)   [![godoc][D]](https://godoc.org/github.com/playlyfe/playlyfe-go-sdk)
- [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) **star:2** Go wrapper for the TripAdvisor API.   [![It hasn't been updated in recent three years][Y]](https://github.com/mrbenosborne/tripadvisor-golang)   [![godoc][D]](https://godoc.org/github.com/mrbenosborne/tripadvisor-golang)
- [vl-go](https://github.com/verifid/vl-go) **star:2** Go client library around the VerifID identity verification layer API.   [![It hasn't been updated in recent three years][Y]](https://github.com/verifid/vl-go)   [![godoc][D]](https://godoc.org/github.com/verifid/vl-go)
- [colony-sdk-go](https://github.com/TheColonyCC/colony-sdk-go) **star:1** Go client library for [The Colony](https://thecolony.cc) — a public social network whose users are AI agents.   [![There was an update last month][G]](https://github.com/TheColonyCC/colony-sdk-go)   [![godoc][D]](https://godoc.org/github.com/TheColonyCC/colony-sdk-go)

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [fzf](https://github.com/junegunn/fzf) **star:81727** Command-line fuzzy finder written in Go.   [![There was an update last month][G]](https://github.com/junegunn/fzf)   [![godoc][D]](https://godoc.org/github.com/junegunn/fzf)
- [dive](https://github.com/wagoodman/dive) **star:54322** A tool for exploring each layer in a Docker image.   [![godoc][D]](https://godoc.org/github.com/wagoodman/dive)
- [hub](https://github.com/github/hub) **star:22950** wrap git commands with additional functionality to interact with github from the terminal.   [![godoc][D]](https://godoc.org/github.com/github/hub)
- [lo](https://github.com/samber/lo) **star:21385** A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...)   [![There was an update last month][G]](https://github.com/samber/lo)   [![godoc][D]](https://godoc.org/github.com/samber/lo)
- [ctop](https://github.com/bcicen/ctop) **star:17787** [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics.   [![godoc][D]](https://godoc.org/github.com/bcicen/ctop)
- [sqlx](https://github.com/jmoiron/sqlx) **star:17699** provides a set of extensions on top of the excellent built-in database/sql package.   [![godoc][D]](https://godoc.org/github.com/jmoiron/sqlx)
- [goreleaser](https://github.com/goreleaser/goreleaser) **star:15925** Deliver Go binaries as fast and easily as possible.   [![There was an update last month][G]](https://github.com/goreleaser/goreleaser)   [![godoc][D]](https://godoc.org/github.com/goreleaser/goreleaser)
- [wuzz](https://github.com/asciimoo/wuzz) **star:10709** Interactive cli tool for HTTP inspection.   [![godoc][D]](https://godoc.org/github.com/asciimoo/wuzz)
- [usql](https://github.com/knq/usql) **star:10036** usql is a universal command-line interface for SQL databases.   [![godoc][D]](https://godoc.org/github.com/knq/usql)
- [peco](https://github.com/peco/peco) **star:7901** Simplistic interactive filtering tool.   [![There was an update last month][G]](https://github.com/peco/peco)   [![godoc][D]](https://godoc.org/github.com/peco/peco)
- [go-funk](https://github.com/thoas/go-funk) **star:4939** Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).   [![godoc][D]](https://godoc.org/github.com/thoas/go-funk)
- [godropbox](https://github.com/dropbox/godropbox) **star:4207** Common libraries for writing Go services/applications from Dropbox.   [![godoc][D]](https://godoc.org/github.com/dropbox/godropbox)
- [minify](https://github.com/tdewolff/minify) **star:4127** Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.   [![There was an update last month][G]](https://github.com/tdewolff/minify)   [![godoc][D]](https://godoc.org/github.com/tdewolff/minify)
- [panicparse](https://github.com/maruel/panicparse) **star:3713** Groups similar goroutines and colorizes stack dump.   [![There was an update last month][G]](https://github.com/maruel/panicparse)   [![godoc][D]](https://godoc.org/github.com/maruel/panicparse)
- [mc](https://github.com/minio/mc) **star:3522** Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.   [![godoc][D]](https://godoc.org/github.com/minio/mc)   [![Archived][Archived]](https://github.com/minio/mc)
- [goreporter](https://github.com/wgliang/goreporter) **star:3121** Golang tool that does static analysis, unit testing, code review and generate code quality report.   [![It hasn't been updated in recent three years][Y]](https://github.com/wgliang/goreporter)   [![godoc][D]](https://godoc.org/github.com/wgliang/goreporter)
- [mergo](https://github.com/imdario/mergo) **star:3103** Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.   [![godoc][D]](https://godoc.org/github.com/imdario/mergo)
- [retry-go](https://github.com/avast/retry-go) **star:2937** Simple library for retry mechanism.   [![godoc][D]](https://godoc.org/github.com/avast/retry-go)
- [create-go-app](https://github.com/create-go-app/cli) **star:2764** A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command.   [![godoc][D]](https://godoc.org/github.com/create-go-app/cli)
- [sesh](https://github.com/joshmedeski/sesh) **star:2681** Sesh is a CLI that helps you create and manage tmux sessions quickly and easily using zoxide.   [![There was an update last month][G]](https://github.com/joshmedeski/sesh)   [![godoc][D]](https://godoc.org/github.com/joshmedeski/sesh)
- [filetype](https://github.com/h2non/filetype) **star:2298** Small package to infer the file type checking the magic numbers signature.   [![godoc][D]](https://godoc.org/github.com/h2non/filetype)
- [EaseProbe](https://github.com/megaease/easeprobe) **star:2297** A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification.   [![There was an update last month][G]](https://github.com/megaease/easeprobe)   [![godoc][D]](https://godoc.org/github.com/megaease/easeprobe)
- [Failsafe-go](https://github.com/failsafe-go/failsafe-go) **star:2230** Fault tolerance and resilience patterns for Go.   [![godoc][D]](https://godoc.org/github.com/failsafe-go/failsafe-go)
- [Storm](https://github.com/asdine/storm) **star:2098** Simple and powerful toolkit for BoltDB.   [![godoc][D]](https://godoc.org/github.com/asdine/storm)
- [mimetype](https://github.com/gabriel-vasile/mimetype) **star:1993** Package for MIME type detection based on magic numbers.   [![There was an update last month][G]](https://github.com/gabriel-vasile/mimetype)   [![godoc][D]](https://godoc.org/github.com/gabriel-vasile/mimetype)
- [jump](https://github.com/gsamokovarov/jump) **star:1934** Jump helps you navigate faster by learning your habits.   [![godoc][D]](https://godoc.org/github.com/gsamokovarov/jump)
- [boilr](https://github.com/tmrts/boilr) **star:1762** Blazingly fast CLI tool for creating projects from boilerplate templates.   [![It hasn't been updated in recent three years][Y]](https://github.com/tmrts/boilr)   [![godoc][D]](https://godoc.org/github.com/tmrts/boilr)
- [mole](https://github.com/davrodpin/mole) **star:1724** cli app to easily create ssh tunnels.   [![godoc][D]](https://godoc.org/github.com/davrodpin/mole)
- [boring](https://github.com/alebeck/boring) **star:1653** Simple command-line SSH tunnel manager.   [![There was an update last month][G]](https://github.com/alebeck/boring)   [![godoc][D]](https://godoc.org/github.com/alebeck/boring)
- [gitbatch](https://github.com/isacikgoz/gitbatch) **star:1563** manage your git repositories in one place.   [![It hasn't been updated in recent three years][Y]](https://github.com/isacikgoz/gitbatch)   [![godoc][D]](https://godoc.org/github.com/isacikgoz/gitbatch)
- [gitcs](https://github.com/knbr13/gitcs/)  Git Commits Visualizer, CLI tool to visualize your Git commits on your local machine.
- [scany](https://github.com/georgysavva/scany) **star:1522** Library for scanning data from a database into Go structs and more.   [![godoc][D]](https://godoc.org/github.com/georgysavva/scany)
- [bed](https://github.com/itchyny/bed) **star:1348** A Vim-like binary editor written in Go.   [![godoc][D]](https://godoc.org/github.com/itchyny/bed)
- [upterm](https://github.com/owenthereal/upterm) **star:1253** A tool for developers to share terminal/tmux sessions securely over the web. It’s perfect for remote pair programming, accessing computers behind NATs/firewalls, remote debugging, and more.   [![There was an update last month][G]](https://github.com/owenthereal/upterm)   [![godoc][D]](https://godoc.org/github.com/owenthereal/upterm)
- [hostctl](https://github.com/guumaster/hostctl) **star:1227** A CLI tool to manage /etc/hosts with easy commands.   [![godoc][D]](https://godoc.org/github.com/guumaster/hostctl)
- [circuitbreaker](https://github.com/rubyist/circuitbreaker) **star:1164** Circuit Breakers in Go.   [![godoc][D]](https://godoc.org/github.com/rubyist/circuitbreaker)
- [git-time-metric](https://github.com/git-time-metric/gtm) **star:1001** Simple, seamless, lightweight time tracking for Git.   [![It hasn't been updated in recent three years][Y]](https://github.com/git-time-metric/gtm)   [![godoc][D]](https://godoc.org/github.com/git-time-metric/gtm)
- [changie](https://github.com/miniscruff/changie) **star:891** Automated changelog tool for preparing releases with lots of customization options.   [![There was an update last month][G]](https://github.com/miniscruff/changie)   [![godoc][D]](https://godoc.org/github.com/miniscruff/changie)
- [clipboard](https://github.com/golang-design/clipboard) **star:842** 📋 cross-platform clipboard package in Go.   [![godoc][D]](https://godoc.org/github.com/golang-design/clipboard)
- [immortal](https://github.com/immortal/immortal) **star:837** \*nix cross-platform (OS agnostic) supervisor.   [![There was an update last month][G]](https://github.com/immortal/immortal)   [![godoc][D]](https://godoc.org/github.com/immortal/immortal)
- [circuit](https://github.com/cep21/circuit) **star:816** An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.   [![godoc][D]](https://godoc.org/github.com/cep21/circuit)
- [mani](https://github.com/alajmo/mani) **star:734** CLI tool to help you manage multiple repositories.   [![godoc][D]](https://godoc.org/github.com/alajmo/mani)
- [clockwork](https://github.com/jonboulle/clockwork) **star:727** A simple fake clock for golang.   [![godoc][D]](https://godoc.org/github.com/jonboulle/clockwork)
- [remote-touchpad](https://github.com/Unrud/remote-touchpad) **star:668** Control mouse and keyboard from a smartphone.   [![There was an update last month][G]](https://github.com/Unrud/remote-touchpad)   [![godoc][D]](https://godoc.org/github.com/Unrud/remote-touchpad)
- [delve](https://github.com/derekparker/delve) **star:661** Go debugger.   [![There was an update last month][G]](https://github.com/derekparker/delve)   [![godoc][D]](https://godoc.org/github.com/derekparker/delve)
- [limiters](https://github.com/mennanov/limiters) **star:650** Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.   [![godoc][D]](https://godoc.org/github.com/mennanov/limiters)
- [ergo](https://github.com/cristianoliveira/ergo) **star:648** The management of multiple local services running over different ports made easy.   [![godoc][D]](https://godoc.org/github.com/cristianoliveira/ergo)
- [scan](https://github.com/blockloop/scan) **star:614** Scan golang `sql.Rows` directly to structs, slices, or primitive types.   [![There was an update last month][G]](https://github.com/blockloop/scan)   [![godoc][D]](https://godoc.org/github.com/blockloop/scan)
- [htcat](https://github.com/htcat/htcat) **star:557** Parallel and Pipelined HTTP GET Utility.   [![godoc][D]](https://godoc.org/github.com/htcat/htcat)
- [koazee](https://github.com/wesovilabs/koazee) **star:529** Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.   [![It hasn't been updated in recent three years][Y]](https://github.com/wesovilabs/koazee)   [![godoc][D]](https://godoc.org/github.com/wesovilabs/koazee)
- [countries](https://github.com/biter777/countries) **star:523** Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standards.   [![godoc][D]](https://godoc.org/github.com/biter777/countries)
- [stacktower](https://github.com/stacktower-io/stacktower) **star:518** Visualize dependency graphs as physical tower structures, inspired by XKCD #2347.   [![godoc][D]](https://godoc.org/github.com/stacktower-io/stacktower)
- [gubrak](https://github.com/novalagung/gubrak) **star:515** Golang utility library with syntactic sugar. It's like lodash, but for golang.   [![godoc][D]](https://godoc.org/github.com/novalagung/gubrak)
- [config-file-validator](https://github.com/Boeing/config-file-validator) **star:510** Cross Platform tool to validate configuration files.   [![There was an update last month][G]](https://github.com/Boeing/config-file-validator)   [![godoc][D]](https://godoc.org/github.com/Boeing/config-file-validator)
- [godaemon](https://github.com/VividCortex/godaemon) **star:494** Utility to write daemons.   [![It hasn't been updated in recent three years][Y]](https://github.com/VividCortex/godaemon)   [![godoc][D]](https://godoc.org/github.com/VividCortex/godaemon)
- [go-dry](https://github.com/ungerik/go-dry) **star:488** DRY (don't repeat yourself) package for Go.   [![godoc][D]](https://godoc.org/github.com/ungerik/go-dry)
- [gopencils](https://github.com/bndr/gopencils) **star:453** Small and simple package to easily consume REST APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/bndr/gopencils)   [![godoc][D]](https://godoc.org/github.com/bndr/gopencils)
- [request](https://github.com/mozillazg/request) **star:421** Go HTTP Requests for Humans™.   [![It hasn't been updated in recent three years][Y]](https://github.com/mozillazg/request)   [![godoc][D]](https://godoc.org/github.com/mozillazg/request)
- [go-rate](https://github.com/beefsack/go-rate) **star:406** Timed rate limiter for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/beefsack/go-rate)   [![godoc][D]](https://godoc.org/github.com/beefsack/go-rate)
- [grofer](https://github.com/pesos/grofer) **star:375** A system and resource monitoring tool written in Golang!   [![It hasn't been updated in recent three years][Y]](https://github.com/pesos/grofer)   [![godoc][D]](https://godoc.org/github.com/pesos/grofer)
- [rospo](https://github.com/ferama/rospo) **star:371** Simple and reliable ssh tunnels with embedded ssh server in Golang.   [![There was an update last month][G]](https://github.com/ferama/rospo)   [![godoc][D]](https://godoc.org/github.com/ferama/rospo)
- [serve](https://github.com/syntaqx/serve) **star:352** A static http server anywhere you need.   [![godoc][D]](https://godoc.org/github.com/syntaqx/serve)
- [retry](https://github.com/kamilsk/retry) **star:344** The most advanced functional mechanism to perform actions repetitively until successful.   [![godoc][D]](https://godoc.org/github.com/kamilsk/retry)
- [gotenv](https://github.com/subosito/gotenv) **star:309** Load environment variables from `.env` or any `io.Reader` in Go.   [![godoc][D]](https://godoc.org/github.com/subosito/gotenv)
- [util](https://github.com/shomali11/util) **star:299** Collection of useful utility functions. (strings, concurrency, manipulations, ...).   [![It hasn't been updated in recent three years][Y]](https://github.com/shomali11/util)   [![godoc][D]](https://godoc.org/github.com/shomali11/util)
- [wifiqr](https://github.com/reugn/wifiqr) **star:287** Wi-Fi QR Code Generator.   [![godoc][D]](https://godoc.org/github.com/reugn/wifiqr)
- [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) **star:251** Pattern matching library.   [![It hasn't been updated in recent three years][Y]](https://github.com/alexpantyukhin/go-pattern-match)   [![godoc][D]](https://godoc.org/github.com/alexpantyukhin/go-pattern-match)
- [go-trigger](https://github.com/sadlil/go-trigger) **star:250** Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.   [![It hasn't been updated in recent three years][Y]](https://github.com/sadlil/go-trigger)   [![godoc][D]](https://godoc.org/github.com/sadlil/go-trigger)
- [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) **star:230** XML Sitemap generator written in Go.   [![godoc][D]](https://godoc.org/github.com/ikeikeikeike/go-sitemap-generator)
- [toolbox](https://github.com/viant/toolbox) **star:230** Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.   [![godoc][D]](https://godoc.org/github.com/viant/toolbox)
- [json-log-viewer](https://github.com/hedhyw/json-log-viewer) **star:229** Interactive viewer for JSON logs.   [![godoc][D]](https://godoc.org/github.com/hedhyw/json-log-viewer)
- [Death](https://github.com/vrecan/death) **star:198** Managing go application shutdown with signals.   [![godoc][D]](https://godoc.org/github.com/vrecan/death)
- [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) **star:181** go:generate tool for wrapping symbols exported by golang plugins (1.8 only).   [![It hasn't been updated in recent three years][Y]](https://github.com/wendigo/go-bind-plugin)   [![godoc][D]](https://godoc.org/github.com/wendigo/go-bind-plugin)
- [goval](https://github.com/maja42/goval) **star:174** Evaluate arbitrary expressions in Go.   [![godoc][D]](https://godoc.org/github.com/maja42/goval)
- [go-bsdiff](https://github.com/gabstv/go-bsdiff) **star:172** Pure Go bsdiff and bspatch libraries and CLI tools.   [![It hasn't been updated in recent three years][Y]](https://github.com/gabstv/go-bsdiff)   [![godoc][D]](https://godoc.org/github.com/gabstv/go-bsdiff)
- [rate](https://github.com/webriots/rate) **star:170** High-performance rate limiting library with token bucket and AIMD strategies.   [![godoc][D]](https://godoc.org/github.com/webriots/rate)
- [apm](https://github.com/topfreegames/apm) **star:169** Process manager for Golang applications with an HTTP API.   [![It hasn't been updated in recent three years][Y]](https://github.com/topfreegames/apm)   [![godoc][D]](https://godoc.org/github.com/topfreegames/apm)
- [moldova](https://github.com/StabbyCutyou/moldova) **star:168** Utility for generating random data based on an input template.   [![It hasn't been updated in recent three years][Y]](https://github.com/StabbyCutyou/moldova)   [![godoc][D]](https://godoc.org/github.com/StabbyCutyou/moldova)
- [rerun](https://github.com/ivpusic/rerun) **star:165** Recompiling and rerunning go apps when source changes.   [![It hasn't been updated in recent three years][Y]](https://github.com/ivpusic/rerun)   [![godoc][D]](https://godoc.org/github.com/ivpusic/rerun)
- [cryptgo](https://github.com/Gituser143/cryptgo) **star:163** Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time!   [![It hasn't been updated in recent three years][Y]](https://github.com/Gituser143/cryptgo)   [![godoc][D]](https://godoc.org/github.com/Gituser143/cryptgo)
- [chyle](https://github.com/antham/chyle) **star:162** Changelog generator using a git repository with multiple configuration possibilities.   [![There was an update last month][G]](https://github.com/antham/chyle)   [![godoc][D]](https://godoc.org/github.com/antham/chyle)
- [cmd](https://github.com/SimonBaeumer/cmd) **star:161** Library for executing shell commands on osx, windows and linux.   [![godoc][D]](https://godoc.org/github.com/SimonBaeumer/cmd)
- [robustly](https://github.com/VividCortex/robustly) **star:159** Runs functions resiliently, catching and restarting panics.   [![godoc][D]](https://godoc.org/github.com/VividCortex/robustly)
- [filter](https://github.com/gookit/filter) **star:151** provide filtering, sanitizing, and conversion of Go data.   [![godoc][D]](https://godoc.org/github.com/gookit/filter)
- [nostromo](https://github.com/pokanop/nostromo) **star:150** CLI for building powerful aliases.   [![godoc][D]](https://godoc.org/github.com/pokanop/nostromo)
- [sorty](https://github.com/jfcg/sorty) **star:145** Fast Concurrent / Parallel Sorting.   [![godoc][D]](https://godoc.org/github.com/jfcg/sorty)
- [gh-image](https://github.com/drogers0/gh-image) **star:142** A gh CLI extension that uploads images to GitHub issues, PRs, and READMEs from the command line, producing user-attachments URLs that respect repository visibility.   [![godoc][D]](https://godoc.org/github.com/drogers0/gh-image)
- [onecache](https://github.com/adelowo/onecache) **star:135** Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).   [![It hasn't been updated in recent three years][Y]](https://github.com/adelowo/onecache)   [![godoc][D]](https://godoc.org/github.com/adelowo/onecache)
- [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) **star:131** Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines.   [![It hasn't been updated in recent three years][Y]](https://github.com/gobeam/mongo-go-pagination)   [![godoc][D]](https://godoc.org/github.com/gobeam/mongo-go-pagination)
- [lrserver](https://github.com/jaschaephraim/lrserver) **star:128** LiveReload server for Go.
- [go-lock](https://github.com/viney-shih/go-lock) **star:125** go-lock is a lock library implementing read-write mutex and read-write trylock without starvation.   [![It hasn't been updated in recent three years][Y]](https://github.com/viney-shih/go-lock)   [![godoc][D]](https://godoc.org/github.com/viney-shih/go-lock)
- [goseaweedfs](https://github.com/linxGnu/goseaweedfs) **star:118** SeaweedFS client library with almost full features.   [![It hasn't been updated in recent three years][Y]](https://github.com/linxGnu/goseaweedfs)   [![godoc][D]](https://godoc.org/github.com/linxGnu/goseaweedfs)
- [cookie](https://github.com/syntaqx/cookie) **star:115** Cookie struct parsing and helper package.   [![godoc][D]](https://godoc.org/github.com/syntaqx/cookie)
- [mssqlx](https://github.com/linxGnu/mssqlx) **star:104** Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.   [![godoc][D]](https://godoc.org/github.com/linxGnu/mssqlx)
- [go-pattern-match](https://github.com/PhakornKiong/go-pattern-match) **star:102** A Pattern matching library inspired by ts-pattern.   [![godoc][D]](https://godoc.org/github.com/PhakornKiong/go-pattern-match)
- [mimemagic](https://github.com/zRedShift/mimemagic) **star:100** Pure Go ultra performant MIME sniffing library/utility.   [![godoc][D]](https://godoc.org/github.com/zRedShift/mimemagic)
- [xferspdy](https://github.com/monmohan/xferspdy) **star:99** Xferspdy provides binary diff and patch library in golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/monmohan/xferspdy)   [![godoc][D]](https://godoc.org/github.com/monmohan/xferspdy)
- [countries](https://github.com/pioz/countries) **star:98** All you need when you are working with countries in Go.   [![godoc][D]](https://godoc.org/github.com/pioz/countries)
- [go-health](https://github.com/Talento90/go-health) **star:97** Health package simplifies the way you add health check to your services.   [![It hasn't been updated in recent three years][Y]](https://github.com/Talento90/go-health)   [![godoc][D]](https://godoc.org/github.com/Talento90/go-health)
- [go-lambda-cleanup](https://github.com/karl-cardenas-coding/go-lambda-cleanup) **star:97** A CLI for removing unused or previous versions of AWS Lambdas.   [![There was an update last month][G]](https://github.com/karl-cardenas-coding/go-lambda-cleanup)   [![godoc][D]](https://godoc.org/github.com/karl-cardenas-coding/go-lambda-cleanup)
- [equalizer](https://github.com/reugn/equalizer) **star:89** Quota manager and rate limiter collection for Go.   [![godoc][D]](https://godoc.org/github.com/reugn/equalizer)
- [go-safecast](https://github.com/ccoVeille/go-safecast) **star:89** Safe number type conversion library that prevents integer overflow and underflow (addresses gosec G115 and CWE-190).   [![godoc][D]](https://godoc.org/github.com/ccoVeille/go-safecast)
- [pgo](https://github.com/arthurkushman/pgo) **star:89** Convenient functions for PHP community.   [![godoc][D]](https://godoc.org/github.com/arthurkushman/pgo)
- [godoclive](https://github.com/syst3mctl/godoclive) **star:84** Generates interactive API documentation from Go HTTP handlers using static analysis of chi, gin, and net/http routers.   [![There was an update last month][G]](https://github.com/syst3mctl/godoclive)   [![godoc][D]](https://godoc.org/github.com/syst3mctl/godoclive)
- [repeat](https://github.com/ssgreg/repeat) **star:84** Go implementation of different backoff strategies useful for retrying operations and heartbeating.   [![It hasn't been updated in recent three years][Y]](https://github.com/ssgreg/repeat)   [![godoc][D]](https://godoc.org/github.com/ssgreg/repeat)
- [handy](https://github.com/miguelpragier/handy) **star:83** Many utilities and helpers like string handlers/formatters and validators.   [![It hasn't been updated in recent three years][Y]](https://github.com/miguelpragier/handy)   [![godoc][D]](https://godoc.org/github.com/miguelpragier/handy)
- [pm](https://github.com/VividCortex/pm) **star:79** Process (i.e. goroutine) manager with an HTTP API.   [![godoc][D]](https://godoc.org/github.com/VividCortex/pm)
- [netbug](https://github.com/e-dard/netbug) **star:72** Easy remote profiling of your services.   [![It hasn't been updated in recent three years][Y]](https://github.com/e-dard/netbug)   [![godoc][D]](https://godoc.org/github.com/e-dard/netbug)
- [multitick](https://github.com/VividCortex/multitick) **star:71** Multiplexor for aligned tickers.   [![godoc][D]](https://godoc.org/github.com/VividCortex/multitick)
- [UNIS](https://github.com/esemplastic/unis) **star:70** Common Architecture™ for String Utilities in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/esemplastic/unis)   [![godoc][D]](https://godoc.org/github.com/esemplastic/unis)
- [backscanner](https://github.com/icza/backscanner) **star:69** A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.   [![godoc][D]](https://godoc.org/github.com/icza/backscanner)
- [dbt](https://github.com/nikogura/dbt) **star:69** A framework for running self-updating signed binaries from a central, trusted repository.   [![godoc][D]](https://godoc.org/github.com/nikogura/dbt)
- [goreadability](https://github.com/philipjkim/goreadability) **star:69** Webpage summary extractor using Facebook Open Graph and arc90's readability.   [![It hasn't been updated in recent three years][Y]](https://github.com/philipjkim/goreadability)   [![godoc][D]](https://godoc.org/github.com/philipjkim/goreadability)
- [scan](https://github.com/wroge/scan) **star:68** Scan sql rows into any type powered by generics.   [![godoc][D]](https://godoc.org/github.com/wroge/scan)
- [retry](https://github.com/thedevsaddam/retry) **star:67** Simple and easy retry mechanism package for Go.   [![godoc][D]](https://godoc.org/github.com/thedevsaddam/retry)
- [go-astitodo](https://github.com/asticode/go-astitodo) **star:65** Parse TODOs in your GO code.   [![godoc][D]](https://godoc.org/github.com/asticode/go-astitodo)
- [golog](https://github.com/mlimaloureiro/golog) **star:63** Easy and lightweight CLI tool to time track your tasks.   [![It hasn't been updated in recent three years][Y]](https://github.com/mlimaloureiro/golog)   [![godoc][D]](https://godoc.org/github.com/mlimaloureiro/golog)
- [minquery](https://github.com/icza/minquery) **star:61** MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).   [![It hasn't been updated in recent three years][Y]](https://github.com/icza/minquery)   [![godoc][D]](https://godoc.org/github.com/icza/minquery)
- [shutdown](https://github.com/ztrue/shutdown) **star:60** App shutdown hooks for `os.Signal` handling.   [![It hasn't been updated in recent three years][Y]](https://github.com/ztrue/shutdown)   [![godoc][D]](https://godoc.org/github.com/ztrue/shutdown)
- [go-utils](https://github.com/Goldziher/go-utils) **star:59** Simple, performant generic utilities for Go inspired by JavaScript and Python (map, filter, reduce, and more).   [![There was an update last month][G]](https://github.com/Goldziher/go-utils)   [![godoc][D]](https://godoc.org/github.com/Goldziher/go-utils)
- [go-qr](https://github.com/piglig/go-qr) **star:57** A native, high-quality and minimalistic QR code generator.   [![godoc][D]](https://godoc.org/github.com/piglig/go-qr)
- [sshman](https://github.com/shoobyban/sshman) **star:57** SSH Manager for authorized_keys files on multiple remote servers.
- [copy-pasta](https://github.com/jutkko/copy-pasta) **star:56** Universal multi-workstation clipboard that uses S3 like backend for the storage.   [![It hasn't been updated in recent three years][Y]](https://github.com/jutkko/copy-pasta)   [![godoc][D]](https://godoc.org/github.com/jutkko/copy-pasta)
- [ghokin](https://github.com/antham/ghokin) **star:56** Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).   [![There was an update last month][G]](https://github.com/antham/ghokin)   [![godoc][D]](https://godoc.org/github.com/antham/ghokin)
- [sqlex](https://github.com/go-sqlex/sqlex) **star:55** Drop-in modernization of jmoiron/sqlx with fixed SQL lexer bugs, automatic IN-clause expansion, pluggable hooks, and unified DB/Tx/Conn interfaces.   [![There was an update last month][G]](https://github.com/go-sqlex/sqlex)   [![godoc][D]](https://godoc.org/github.com/go-sqlex/sqlex)   [![Contains Chinese documents][CN]](https://github.com/go-sqlex/sqlex)
- [cvt](https://github.com/shockerli/cvt) **star:54** Easy and safe convert any value to another type.   [![godoc][D]](https://godoc.org/github.com/shockerli/cvt)   [![Contains Chinese documents][CN]](https://github.com/shockerli/cvt)
- [golarm](https://github.com/msempere/golarm) **star:54** Fire alarms with system events.   [![It hasn't been updated in recent three years][Y]](https://github.com/msempere/golarm)   [![godoc][D]](https://godoc.org/github.com/msempere/golarm)
- [gofn](https://github.com/tiendc/gofn) **star:53** High performance utility functions written using Generics for Go 1.18+.   [![godoc][D]](https://godoc.org/github.com/tiendc/gofn)
- [slice](https://github.com/psampaz/slice) **star:51** Type-safe functions for common Go slice operations.   [![It hasn't been updated in recent three years][Y]](https://github.com/psampaz/slice)   [![godoc][D]](https://godoc.org/github.com/psampaz/slice)
- [goback](https://github.com/carlescere/goback) **star:50** Go simple exponential backoff package.   [![It hasn't been updated in recent three years][Y]](https://github.com/carlescere/goback)   [![godoc][D]](https://godoc.org/github.com/carlescere/goback)
- [retry-go](https://github.com/rafaeljesus/retry-go) **star:49** Retrying made simple and easy for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafaeljesus/retry-go)   [![godoc][D]](https://godoc.org/github.com/rafaeljesus/retry-go)
- [pointer](https://github.com/xorcare/pointer) **star:48** Package pointer contains helper routines for simplifying the creation of optional fields of basic type.   [![godoc][D]](https://godoc.org/github.com/xorcare/pointer)
- [go-httpheader](https://github.com/mozillazg/go-httpheader) **star:47** Go library for encoding structs into Header fields.   [![It hasn't been updated in recent three years][Y]](https://github.com/mozillazg/go-httpheader)   [![godoc][D]](https://godoc.org/github.com/mozillazg/go-httpheader)
- [gostrutils](https://github.com/ik5/gostrutils) **star:47** Collections of string manipulation and conversion functions.   [![godoc][D]](https://godoc.org/github.com/ik5/gostrutils)
- [set](https://github.com/nofeaturesonlybugs/set) **star:47** Performant and flexible struct mapping and loose type conversion.   [![It hasn't been updated in recent three years][Y]](https://github.com/nofeaturesonlybugs/set)   [![godoc][D]](https://godoc.org/github.com/nofeaturesonlybugs/set)
- [slicer](https://github.com/leaanthony/slicer) **star:47** Makes working with slices easier.   [![It hasn't been updated in recent three years][Y]](https://github.com/leaanthony/slicer)   [![godoc][D]](https://godoc.org/github.com/leaanthony/slicer)
- [yogo](https://github.com/antham/yogo) **star:47** Check yopmail mails from command line.   [![There was an update last month][G]](https://github.com/antham/yogo)
- [evaluator](https://github.com/nullne/evaluator) **star:42** Evaluate an expression dynamically based on s-expression. It's simple and easy to extend.   [![It hasn't been updated in recent three years][Y]](https://github.com/nullne/evaluator)   [![godoc][D]](https://godoc.org/github.com/nullne/evaluator)
- [throttle](https://github.com/yudppp/throttle) **star:40** Throttle is an object that will perform exactly one action per duration.   [![It hasn't been updated in recent three years][Y]](https://github.com/yudppp/throttle)   [![godoc][D]](https://godoc.org/github.com/yudppp/throttle)
- [just](https://github.com/kazhuravlev/just) **star:38** Just a collection of useful functions for working with generic data structures.   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/just)
- [debounce](https://github.com/floatdrop/debounce) **star:37** A zero-allocation debouncer written in Go.   [![godoc][D]](https://godoc.org/github.com/floatdrop/debounce)
- [rclient](https://github.com/zpatrick/rclient) **star:36** Readable, flexible, simple-to-use client for REST APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/zpatrick/rclient)   [![godoc][D]](https://godoc.org/github.com/zpatrick/rclient)
- [tome](https://github.com/cyruzin/tome) **star:35** Tome was designed to paginate simple RESTful APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/cyruzin/tome)   [![godoc][D]](https://godoc.org/github.com/cyruzin/tome)
- [watchhttp](https://github.com/nikolaydubina/watchhttp) **star:35** Run command periodically and expose latest STDOUT or its rich delta as HTTP endpoint.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/watchhttp)
- [git-tools](https://github.com/kazhuravlev/git-tools) **star:33** Tool to help manage git tags.   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/git-tools)
- [generate](https://github.com/go-playground/generate) **star:31** runs go generate recursively on a specified path or environment variable and can filter by regex.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/generate)   [![godoc][D]](https://godoc.org/github.com/go-playground/generate)
- [graterm](https://github.com/skovtunenko/graterm) **star:30** Provides primitives to perform ordered (sequential/concurrent) GRAceful TERMination (aka shutdown) in Go application.   [![godoc][D]](https://godoc.org/github.com/skovtunenko/graterm)
- [ptr](https://github.com/gotidy/ptr) **star:30** Package that provide functions for simplified creation of pointers from constants of basic types.   [![It hasn't been updated in recent three years][Y]](https://github.com/gotidy/ptr)   [![godoc][D]](https://godoc.org/github.com/gotidy/ptr)
- [goplaceholder](https://github.com/michiwend/goplaceholder) **star:29** a small golang lib to generate placeholder images.   [![It hasn't been updated in recent three years][Y]](https://github.com/michiwend/goplaceholder)   [![godoc][D]](https://godoc.org/github.com/michiwend/goplaceholder)
- [ctxutil](https://github.com/posener/ctxutil) **star:26** A collection of utility functions for contexts.   [![It hasn't been updated in recent three years][Y]](https://github.com/posener/ctxutil)   [![godoc][D]](https://godoc.org/github.com/posener/ctxutil)
- [go-convert](https://github.com/Eun/go-convert) **star:24** Package go-convert enables you to convert a value into another type.   [![godoc][D]](https://godoc.org/github.com/Eun/go-convert)
- [go-type](https://github.com/mikekonan/go-types) **star:24** Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types.   [![godoc][D]](https://godoc.org/github.com/mikekonan/go-types)
- [healthcheck](https://github.com/kazhuravlev/healthcheck) **star:24** A simple yet powerful readiness test for Kubernetes.   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/healthcheck)
- [structs](https://github.com/PumpkinSeed/structs) **star:24** Implement simple functions to manipulate structs.   [![It hasn't been updated in recent three years][Y]](https://github.com/PumpkinSeed/structs)   [![godoc][D]](https://godoc.org/github.com/PumpkinSeed/structs)
- [jsend](https://github.com/clevergo/jsend) **star:21** JSend's implementation written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/clevergo/jsend)   [![godoc][D]](https://godoc.org/github.com/clevergo/jsend)
- [filler](https://github.com/yaronsumel/filler) **star:18** small utility to fill structs using "fill" tag.   [![It hasn't been updated in recent three years][Y]](https://github.com/yaronsumel/filler)   [![godoc][D]](https://godoc.org/github.com/yaronsumel/filler)
- [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) **star:18** Go package for working with Problem Details.   [![It hasn't been updated in recent three years][Y]](https://github.com/mvmaasakkers/go-problemdetails)   [![godoc][D]](https://godoc.org/github.com/mvmaasakkers/go-problemdetails)
- [dlog](https://github.com/kirillDanshin/dlog) **star:17** Compile-time controlled logger to make your release smaller without removing debug calls.   [![It hasn't been updated in recent three years][Y]](https://github.com/kirillDanshin/dlog)   [![godoc][D]](https://godoc.org/github.com/kirillDanshin/dlog)
- [go-countries](https://github.com/mikekonan/go-countries) **star:16** Lightweight lookup over ISO-3166 codes.   [![It hasn't been updated in recent three years][Y]](https://github.com/mikekonan/go-countries)   [![godoc][D]](https://godoc.org/github.com/mikekonan/go-countries)
- [okrun](https://github.com/xta/okrun) **star:16** go run error steamroller.   [![It hasn't been updated in recent three years][Y]](https://github.com/xta/okrun)   [![godoc][D]](https://godoc.org/github.com/xta/okrun)
- [release](https://github.com/tomodian/release) **star:16** CLI for Keep-a-changelog formatted changelogs.   [![godoc][D]](https://godoc.org/github.com/tomodian/release)
- [go-clip](https://github.com/prashantgupta24/go-clip) **star:15** A minimalistic clipboard manager for Mac.   [![It hasn't been updated in recent three years][Y]](https://github.com/prashantgupta24/go-clip)   [![godoc][D]](https://godoc.org/github.com/prashantgupta24/go-clip)
- [go-tripper](https://github.com/rajnandan1/go-tripper) **star:15** Tripper is a circuit breaker package for Go that allows you to circuit and control the status of circuits.   [![godoc][D]](https://godoc.org/github.com/rajnandan1/go-tripper)
- [jet](https://github.com/NicoNex/jet) **star:15** Just Edit Text: a fast and powerful tool for finding and replacing file content and names using regular expressions.   [![godoc][D]](https://godoc.org/github.com/NicoNex/jet)
- [rest-go](https://github.com/edermanoel94/rest-go) **star:15** A package that provide many helpful methods for working with rest api.   [![It hasn't been updated in recent three years][Y]](https://github.com/edermanoel94/rest-go)   [![godoc][D]](https://godoc.org/github.com/edermanoel94/rest-go)
- [blank](https://github.com/Henry-Sarabia/blank) **star:14** Verify or remove blanks and whitespace from strings.   [![It hasn't been updated in recent three years][Y]](https://github.com/Henry-Sarabia/blank)   [![godoc][D]](https://godoc.org/github.com/Henry-Sarabia/blank)
- [silk](https://github.com/chrispassas/silk) **star:14** Read silk netflow files.   [![It hasn't been updated in recent three years][Y]](https://github.com/chrispassas/silk)   [![godoc][D]](https://godoc.org/github.com/chrispassas/silk)
- [relimpact](https://github.com/hashmap-kz/relimpact) **star:13** Fast API compatibility reports for Go projects.   [![There was an update last month][G]](https://github.com/hashmap-kz/relimpact)   [![godoc][D]](https://godoc.org/github.com/hashmap-kz/relimpact)
- [retry](https://github.com/shafreeck/retry) **star:13** A pretty simple library to ensure your work to be done.   [![It hasn't been updated in recent three years][Y]](https://github.com/shafreeck/retry)   [![godoc][D]](https://godoc.org/github.com/shafreeck/retry)
- [bleep](https://github.com/sinhashubham95/bleep) **star:11** Perform any number of actions on any set of OS signals in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/sinhashubham95/bleep)   [![godoc][D]](https://godoc.org/github.com/sinhashubham95/bleep)
- [goctx](https://github.com/zerosnake0/goctx) **star:11** Get your context value with high performance.   [![It hasn't been updated in recent three years][Y]](https://github.com/zerosnake0/goctx)   [![godoc][D]](https://godoc.org/github.com/zerosnake0/goctx)
- [loncha](https://github.com/kazu/loncha) **star:11** A high-performance slice Utilities.   [![It hasn't been updated in recent three years][Y]](https://github.com/kazu/loncha)   [![godoc][D]](https://godoc.org/github.com/kazu/loncha)
- [nfdump](https://github.com/chrispassas/nfdump) **star:11** Read nfdump netflow files.   [![godoc][D]](https://godoc.org/github.com/chrispassas/nfdump)
- [optional](https://github.com/kazhuravlev/optional) **star:11** Optional struct fields and vars.   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/optional)
- [retry](https://github.com/percolate/retry) **star:11** A simple but highly configurable retry package for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/percolate/retry)
- [statiks](https://github.com/janiltonmaciel/statiks) **star:11** Fast, zero-configuration, static HTTP filer server.   [![godoc][D]](https://godoc.org/github.com/janiltonmaciel/statiks)
- [go-events](https://github.com/deatil/go-events) **star:10** A go event and event'subscribe package, like wordpress hook functions.   [![godoc][D]](https://godoc.org/github.com/deatil/go-events)   [![Contains Chinese documents][CN]](https://github.com/deatil/go-events)
- [sqlz](https://github.com/rfberaldo/sqlz) **star:10** Extension for the database/sql package, adding named queries, struct scanning, and batch operations.   [![There was an update last month][G]](https://github.com/rfberaldo/sqlz)   [![godoc][D]](https://godoc.org/github.com/rfberaldo/sqlz)
- [sliceconv](https://github.com/Henry-Sarabia/sliceconv) **star:9** Slice conversion between primitive types.   [![It hasn't been updated in recent three years][Y]](https://github.com/Henry-Sarabia/sliceconv)   [![godoc][D]](https://godoc.org/github.com/Henry-Sarabia/sliceconv)
- [go-pkg](https://github.com/chenquan/go-pkg) **star:8** A go toolkit.   [![It hasn't been updated in recent three years][Y]](https://github.com/chenquan/go-pkg)   [![godoc][D]](https://godoc.org/github.com/chenquan/go-pkg)   [![Contains Chinese documents][CN]](https://github.com/chenquan/go-pkg)
- [xpool](https://github.com/peczenyj/xpool) **star:8** Yet another golang type safe object pool using generics.   [![godoc][D]](https://godoc.org/github.com/peczenyj/xpool)
- [abstract](https://github.com/maxbolgarin/abstract) **star:7** Abstractions and utilities to get rid of boilerplate code in business logic.   [![godoc][D]](https://godoc.org/github.com/maxbolgarin/abstract)
- [lang](https://github.com/maxbolgarin/lang) **star:7** Generic one-liners to work with variables, slices and maps without boilerplate code.   [![godoc][D]](https://godoc.org/github.com/maxbolgarin/lang)
- [lets-go](https://github.com/aplescia-chwy/lets-go) **star:7** Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities.   [![It hasn't been updated in recent three years][Y]](https://github.com/aplescia-chwy/lets-go)   [![godoc][D]](https://godoc.org/github.com/aplescia-chwy/lets-go)
- [contem](https://github.com/maxbolgarin/contem) **star:6** Drop-in context.Context replacement for graceful shutdown Go applications.   [![godoc][D]](https://godoc.org/github.com/maxbolgarin/contem)
- [go-snk](https://github.com/SharkByteSoftware/go-snk) **star:5** Type-safe generic helpers for slices, maps, strings, errors, JSON, HTTP, and containers, organized as small independently adoptable packages.   [![There was an update last month][G]](https://github.com/SharkByteSoftware/go-snk)   [![godoc][D]](https://godoc.org/github.com/SharkByteSoftware/go-snk)
- [olaf](https://github.com/btnguyen2k/olaf) **star:5** Twitter Snowflake implemented in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/btnguyen2k/olaf)   [![godoc][D]](https://godoc.org/github.com/btnguyen2k/olaf)
- [tik](https://github.com/andy2046/tik) **star:5** Simple and easy timing wheel package for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/andy2046/tik)   [![godoc][D]](https://godoc.org/github.com/andy2046/tik)
- [go-actuator](https://github.com/sinhashubham95/go-actuator) **star:3** Production ready features for Go based web frameworks.   [![godoc][D]](https://godoc.org/github.com/sinhashubham95/go-actuator)
- [retro](https://github.com/goioc/retro) **star:3** Handy retry-on-error library with extensive flexibility (backoff strategies, caps, etc).   [![godoc][D]](https://godoc.org/github.com/goioc/retro)
- [Go-Constant](https://github.com/sajjadrabiee/go-constant) **star:1** Generic typed constant sets with safe string parsing for Go's missing enum type.   [![godoc][D]](https://godoc.org/github.com/sajjadrabiee/go-constant)

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [uuid](https://github.com/google/uuid) **star:6117** Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.   [![godoc][D]](https://godoc.org/github.com/google/uuid)
- [ulid](https://github.com/oklog/ulid) **star:5047** Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).   [![godoc][D]](https://godoc.org/github.com/oklog/ulid)
- [uniq](https://gitlab.com/skilstak/code/go/uniq)  No hassle safe, fast unique identifiers with commands.
- [xid](https://github.com/rs/xid) **star:4280** Xid is a globally unique id generator library, ready to be safely used directly in your server code.   [![godoc][D]](https://godoc.org/github.com/rs/xid)
- [uuid](https://github.com/gofrs/uuid) **star:1814** Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.   [![godoc][D]](https://godoc.org/github.com/gofrs/uuid)
- [wuid](https://github.com/edwingeng/wuid) **star:545** An extremely fast globally unique number generator.   [![godoc][D]](https://godoc.org/github.com/edwingeng/wuid)
- [sno](https://github.com/muyo/sno) **star:94** Compact, sortable and fast unique IDs with embedded metadata.   [![It hasn't been updated in recent three years][Y]](https://github.com/muyo/sno)   [![godoc][D]](https://godoc.org/github.com/muyo/sno)
- [guid](https://github.com/sdrapkin/guid) **star:75** Fast cryptographically safe Guid generator for Go (~10x faster than `uuid`).   [![There was an update last month][G]](https://github.com/sdrapkin/guid)   [![godoc][D]](https://godoc.org/github.com/sdrapkin/guid)
- [nanoid](https://github.com/aidarkhanov/nanoid) **star:62** A tiny and efficient Go unique string ID generator.   [![It hasn't been updated in recent three years][Y]](https://github.com/aidarkhanov/nanoid)   [![godoc][D]](https://godoc.org/github.com/aidarkhanov/nanoid)   [![Archived][Archived]](https://github.com/aidarkhanov/nanoid)
- [goid](https://github.com/jakehl/goid) **star:41** Generate and Parse RFC4122 compliant V4 UUIDs.   [![It hasn't been updated in recent three years][Y]](https://github.com/jakehl/goid)   [![godoc][D]](https://godoc.org/github.com/jakehl/goid)
- [gouid](https://github.com/twharmon/gouid) **star:27** Generate cryptographically secure random string IDs with just one allocation.   [![godoc][D]](https://godoc.org/github.com/twharmon/gouid)
- [uuid](https://github.com/agext/uuid) **star:18** Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.   [![It hasn't been updated in recent three years][Y]](https://github.com/agext/uuid)   [![godoc][D]](https://godoc.org/github.com/agext/uuid)
- [uuidcheck](https://github.com/ashwingopalsamy/uuidcheck) **star:6** A tiny, dependency-free Go library that validates UUIDs against standard RFC 4122 formatting, converts UUIDv7() into UTC timestamps.   [![godoc][D]](https://godoc.org/github.com/ashwingopalsamy/uuidcheck)
- [fastuuid](https://github.com/rekby/fastuuid) **star:1** Fast generate UUIDv4 as string or bytes.   [![It hasn't been updated in recent three years][Y]](https://github.com/rekby/fastuuid)   [![godoc][D]](https://godoc.org/github.com/rekby/fastuuid)

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [govalidator](https://github.com/asaskevich/govalidator) **star:6210** Validators and sanitizers for strings, numerics, slices and structs.   [![godoc][D]](https://godoc.org/github.com/asaskevich/govalidator)
- [govalidator](https://github.com/thedevsaddam/govalidator) **star:1344** Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.   [![godoc][D]](https://godoc.org/github.com/thedevsaddam/govalidator)
- [gody](https://github.com/guiferpa/gody) **star:181** :balloon: A lightweight struct validator for Go.   [![godoc][D]](https://godoc.org/github.com/guiferpa/gody)
- [govalid](https://github.com/twharmon/govalid) **star:119** Fast, tag-based validation for structs.   [![godoc][D]](https://godoc.org/github.com/twharmon/govalid)
- [checkdigit](https://github.com/osamingo/checkdigit) **star:114** Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).   [![godoc][D]](https://godoc.org/github.com/osamingo/checkdigit)
- [govy](https://github.com/nobl9/govy) **star:50** strongly-typed validation rules over functional interface, powered by generics and reflection free with heavy focus on crafting clear and information-rich error messages.   [![There was an update last month][G]](https://github.com/nobl9/govy)   [![godoc][D]](https://godoc.org/github.com/nobl9/govy)
- [go-validator](https://github.com/tiendc/go-validator) **star:31** Validation library using Generics.   [![godoc][D]](https://godoc.org/github.com/tiendc/go-validator)
- [hvalid](https://github.com/lyonnee/hvalid) hvalid is a lightweight validation library written in Go language. It provides a custom validator interface and a series of common validation functions to help developers quickly implement data validation.
- [validator](https://github.com/go-playground/validator) **star:20069** Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.   [![There was an update last month][G]](https://github.com/go-playground/validator)   [![godoc][D]](https://godoc.org/github.com/go-playground/validator)
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) **star:4108** Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.   [![godoc][D]](https://godoc.org/github.com/go-ozzo/ozzo-validation)
- [validate](https://github.com/gookit/validate) **star:1160** Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.   [![godoc][D]](https://godoc.org/github.com/gookit/validate)   [![Contains Chinese documents][CN]](https://github.com/gookit/validate)
- [jio](https://github.com/faceair/jio) **star:126** jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).   [![godoc][D]](https://godoc.org/github.com/faceair/jio)   [![Contains Chinese documents][CN]](https://github.com/faceair/jio)
- [validate](https://github.com/gobuffalo/validate) **star:94** This package provides a framework for writing validations for Go applications.   [![It hasn't been updated in recent three years][Y]](https://github.com/gobuffalo/validate)   [![godoc][D]](https://godoc.org/github.com/gobuffalo/validate)
- [Validator](https://github.com/go-the-way/validator) **star:7** A lightweight model validator written in Go.Contains VFs:Min, Max, MinLength, MaxLength, Length, Enum, Regex.   [![godoc][D]](https://godoc.org/github.com/go-the-way/validator)
- [valix](https://github.com/marrow16/valix) Go package for validating requests
- [Zog](https://github.com/Oudwins/zog) **star:1199** A [Zod](https://github.com/colinhacks/zod) inspired schema builder for runtime value parsing and validation.   [![There was an update last month][G]](https://github.com/Oudwins/zog)   [![godoc][D]](https://godoc.org/github.com/Oudwins/zog)
  **[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [go-git](https://github.com/go-git/go-git) **star:7613** highly extensible Git implementation in pure Go.   [![There was an update last month][G]](https://github.com/go-git/go-git)   [![godoc][D]](https://godoc.org/github.com/go-git/go-git)
- [hercules](https://github.com/src-d/hercules) **star:2799** gaining advanced insights from Git repository history.   [![It hasn't been updated in recent three years][Y]](https://github.com/src-d/hercules)   [![godoc][D]](https://godoc.org/github.com/src-d/hercules)
- [git2go](https://github.com/libgit2/git2go) **star:2008** Go bindings for libgit2.   [![godoc][D]](https://godoc.org/github.com/libgit2/git2go)
- [cli](https://gitlab.com/gitlab-org/cli)  An open-source GitLab command line tool bringing GitLab's cool features to your command line.
- [ggc](https://github.com/bmf-san/ggc) **star:284** A Git CLI tool with both traditional command-line and interactive incremental-search UI, workflow support, and configurable keybindings.   [![There was an update last month][G]](https://github.com/bmf-san/ggc)   [![godoc][D]](https://godoc.org/github.com/bmf-san/ggc)
- [githooks](https://github.com/gabyx/githooks) **star:126** Per-repo and shared Git hooks with version control and auto update.   [![godoc][D]](https://godoc.org/github.com/gabyx/githooks)
- [go-vcs](https://github.com/sourcegraph/go-vcs) **star:81** manipulate and inspect VCS repositories in Go.   [![godoc][D]](https://godoc.org/github.com/sourcegraph/go-vcs)
- [froggit-go](https://github.com/jfrog/froggit-go) **star:54** Froggit-Go is a Go library, allowing to perform actions on VCS providers.   [![godoc][D]](https://godoc.org/github.com/jfrog/froggit-go)
- [git-courer](https://github.com/Alejandro-M-P/git-courer) **star:43** Local MCP server for Git operations using Ollama to save tokens and prevent secret leakage.   [![There was an update last month][G]](https://github.com/Alejandro-M-P/git-courer)   [![godoc][D]](https://godoc.org/github.com/Alejandro-M-P/git-courer)
- [gitty](https://github.com/Omibranch/gitty) **star:27** Single-binary Git/GitHub CLI that replaces add→commit→push with one command; human-readable syntax, no external dependencies.   [![godoc][D]](https://godoc.org/github.com/Omibranch/gitty)   [![Archived][Archived]](https://github.com/Omibranch/gitty)
- [hgo](https://github.com/beyang/hgo) **star:16** Hgo is a collection of Go packages providing read-access to local Mercurial repositories.   [![It hasn't been updated in recent three years][Y]](https://github.com/beyang/hgo)   [![godoc][D]](https://godoc.org/github.com/beyang/hgo)

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [goav](https://github.com/giorgisio/goav) **star:2137** Comprehensive Go bindings for FFmpeg.   [![It hasn't been updated in recent three years][Y]](https://github.com/giorgisio/goav)   [![godoc][D]](https://godoc.org/github.com/giorgisio/goav)
- [gmf](https://github.com/3d0c/gmf) **star:932** Go bindings for FFmpeg av\* libraries.   [![It hasn't been updated in recent three years][Y]](https://github.com/3d0c/gmf)   [![godoc][D]](https://godoc.org/github.com/3d0c/gmf)
- [gortsplib](https://github.com/aler9/gortsplib) **star:929** Pure Go RTSP server and client library.   [![There was an update last month][G]](https://github.com/aler9/gortsplib)   [![godoc][D]](https://godoc.org/github.com/aler9/gortsplib)
- [go-astiav](https://github.com/asticode/go-astiav) **star:723** Better C bindings for ffmpeg in GO.   [![godoc][D]](https://godoc.org/github.com/asticode/go-astiav)
- [go-astisub](https://github.com/asticode/go-astisub) **star:701** Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).   [![There was an update last month][G]](https://github.com/asticode/go-astisub)   [![godoc][D]](https://godoc.org/github.com/asticode/go-astisub)
- [mp4ff](https://github.com/Eyevinn/mp4ff) **star:643** Library and tools for working with MP4 files containing video, audio, subtitles, or metadata.   [![There was an update last month][G]](https://github.com/Eyevinn/mp4ff)   [![godoc][D]](https://godoc.org/github.com/Eyevinn/mp4ff)
- [go-astits](https://github.com/asticode/go-astits) **star:614** Parse and demux MPEG Transport Streams (.ts) natively in GO.   [![godoc][D]](https://godoc.org/github.com/asticode/go-astits)
- [libvlc-go](https://github.com/adrg/libvlc-go) **star:509** Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).   [![godoc][D]](https://godoc.org/github.com/adrg/libvlc-go)
- [v4l](https://github.com/korandiz/v4l) **star:89** Video capture library for Linux, written in Go.   [![godoc][D]](https://godoc.org/github.com/korandiz/v4l)
- [hls-m3u8](https://github.com/Eyevinn/hls-m3u8) **star:67** Parser and generator for HLS (M3U8) playlists; kept up to date with the spec.   [![There was an update last month][G]](https://github.com/Eyevinn/hls-m3u8)   [![godoc][D]](https://godoc.org/github.com/Eyevinn/hls-m3u8)
- [go-mpd](https://github.com/unki2aut/go-mpd) **star:32** Parser and generator library for MPEG-DASH manifest files.   [![godoc][D]](https://godoc.org/github.com/unki2aut/go-mpd)
- [manifestor](https://github.com/alanzng/manifestor) **star:11** Zero-dependency library for parsing, filtering, transforming, and building HLS and DASH manifests.   [![godoc][D]](https://godoc.org/github.com/alanzng/manifestor)

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [Gin](https://github.com/gin-gonic/gin) **star:88913** Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.   [![godoc][D]](https://godoc.org/github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber) **star:39958** An Express.js inspired web framework build on Fasthttp.   [![There was an update last month][G]](https://github.com/gofiber/fiber)   [![godoc][D]](https://godoc.org/github.com/gofiber/fiber)
- [Echo](https://github.com/labstack/echo) **star:32525** High performance, minimalist Go web framework.   [![There was an update last month][G]](https://github.com/labstack/echo)   [![godoc][D]](https://godoc.org/github.com/labstack/echo)
- [Beego](https://github.com/beego/beego) **star:32412** beego is an open-source, high-performance web framework for the Go programming language.   [![There was an update last month][G]](https://github.com/beego/beego)   [![godoc][D]](https://godoc.org/github.com/beego/beego)
- [Confetti Framework](https://confetti-framework.github.io/docs/)  Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.
- [GoFr](https://github.com/gofr-dev/gofr) **star:21167** Gofr is an opinionated microservice development framework.   [![There was an update last month][G]](https://github.com/gofr-dev/gofr)   [![godoc][D]](https://godoc.org/github.com/gofr-dev/gofr)
- [Revel](https://github.com/revel/revel) **star:13224** High-productivity web framework for the Go language.   [![godoc][D]](https://godoc.org/github.com/revel/revel)
- [GoFrame](https://github.com/gogf/gf) **star:13217** GoFrame is a modular, powerful, high-performance and enterprise-class application development framework of Golang.   [![There was an update last month][G]](https://github.com/gogf/gf)   [![godoc][D]](https://godoc.org/github.com/gogf/gf)
- [Hertz](https://github.com/cloudwego/hertz) **star:7302** A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices.   [![godoc][D]](https://godoc.org/github.com/cloudwego/hertz)   [![Contains Chinese documents][CN]](https://github.com/cloudwego/hertz)
- [Goa](https://github.com/goadesign/goa) **star:6090** Goa provides a holistic approach for developing remote APIs and microservices in Go.   [![There was an update last month][G]](https://github.com/goadesign/goa)   [![godoc][D]](https://godoc.org/github.com/goadesign/goa)
- [goravel](https://github.com/goravel/goravel) **star:4781** A Laravel-inspired web framework with ORM, authentication, queue, task scheduling, and more built-in features.   [![godoc][D]](https://godoc.org/github.com/goravel/goravel)   [![Contains Chinese documents][CN]](https://github.com/goravel/goravel)
- [Goyave](https://github.com/go-goyave/goyave) **star:1774** Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities.   [![godoc][D]](https://godoc.org/github.com/go-goyave/goyave)
- [Fuego](https://github.com/go-fuego/fuego) **star:1750** The framework for busy Go developers! Web framework generating OpenAPI 3 spec from source code.   [![There was an update last month][G]](https://github.com/go-fuego/fuego)   [![godoc][D]](https://godoc.org/github.com/go-fuego/fuego)
- [templui](https://github.com/axzilla/templui) **star:1618** Modern UI Components for Go & Templ.   [![There was an update last month][G]](https://github.com/axzilla/templui)
- [Atreugo](https://github.com/savsgio/atreugo) **star:1304** High performance and extensible micro web framework with zero memory allocations in hot paths.   [![godoc][D]](https://godoc.org/github.com/savsgio/atreugo)
- [Yokai](https://github.com/ankorstore/yokai) **star:836** Simple, modular, and observable Go framework for backend applications.   [![There was an update last month][G]](https://github.com/ankorstore/yokai)   [![godoc][D]](https://godoc.org/github.com/ankorstore/yokai)
- [iWF](https://github.com/indeedeng/iwf) **star:649** iWF is an all-in-one platform for developing long-running business processes. It offers a convenient abstraction for utilizing databases, ElasticSearch, message queues, durable timers, and more, with a clean, simple, and user-friendly interface.   [![There was an update last month][G]](https://github.com/indeedeng/iwf)   [![godoc][D]](https://godoc.org/github.com/indeedeng/iwf)
- [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) **star:590** Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.   [![There was an update last month][G]](https://github.com/i-love-flamingo/flamingo-commerce)   [![godoc][D]](https://godoc.org/github.com/i-love-flamingo/flamingo-commerce)
- [rk-boot](https://github.com/rookie-ninja/rk-boot) **star:573** A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily.   [![godoc][D]](https://godoc.org/github.com/rookie-ninja/rk-boot)
- [Fastschema](https://github.com/fastschema/fastschema) **star:563** A flexible Go web framework and Headless CMS.   [![There was an update last month][G]](https://github.com/fastschema/fastschema)   [![godoc][D]](https://godoc.org/github.com/fastschema/fastschema)
- [Flamingo](https://github.com/i-love-flamingo/flamingo) **star:559** Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.   [![There was an update last month][G]](https://github.com/i-love-flamingo/flamingo)   [![godoc][D]](https://godoc.org/github.com/i-love-flamingo/flamingo)
- [uAdmin](https://github.com/uadmin/uadmin) **star:354** Fully featured web framework for Golang, inspired by Django.   [![There was an update last month][G]](https://github.com/uadmin/uadmin)   [![godoc][D]](https://godoc.org/github.com/uadmin/uadmin)
- [WebGo](https://github.com/naughtygopher/webgo) **star:306** A micro-framework to build web apps with handler chaining, middleware, and context injection. With standard library-compliant HTTP handlers (i.e., `http.HandlerFunc`)..   [![godoc][D]](https://godoc.org/github.com/naughtygopher/webgo)
- [Ginrpc](https://github.com/xxjwxc/ginrpc) **star:305** Gin parameter automatic binding tool,gin rpc tools.   [![godoc][D]](https://godoc.org/github.com/xxjwxc/ginrpc)   [![Contains Chinese documents][CN]](https://github.com/xxjwxc/ginrpc)
- [Andurel](https://github.com/mbvlabs/andurel) **star:215** Rails-inspired full-stack Go web framework with scaffolding, database tooling, and server-rendered or Inertia frontends.   [![There was an update last month][G]](https://github.com/mbvlabs/andurel)   [![godoc][D]](https://godoc.org/github.com/mbvlabs/andurel)
- [hiboot](https://github.com/hidevopsio/hiboot) **star:180** hiboot is a high performance web application framework with auto configuration and dependency injection support.   [![godoc][D]](https://godoc.org/github.com/hidevopsio/hiboot)   [![Contains Chinese documents][CN]](https://github.com/hidevopsio/hiboot)
- [Gone](https://github.com/gone-io/gone) **star:131** A lightweight dependency injection and web framework inspired by Spring.   [![godoc][D]](https://godoc.org/github.com/gone-io/gone)   [![Contains Chinese documents][CN]](https://github.com/gone-io/gone)
- [patron](https://github.com/beatlabs/patron) **star:127** Patron is a microservice framework following best cloud practices with a focus on productivity.   [![godoc][D]](https://godoc.org/github.com/beatlabs/patron)
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux)  Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Microservice](https://github.com/claygod/microservice) **star:123** The framework for the creation of microservices, written in Golang.   [![godoc][D]](https://godoc.org/github.com/claygod/microservice)
- [doors](https://github.com/doors-dev/doors) **star:114** Server-driven framework for building stateful, reactive web applications entirely in Go.   [![godoc][D]](https://godoc.org/github.com/doors-dev/doors)
- [Barf](https://github.com/opensaucerer/barf) **star:105** Basically, A Remarkable Framework for building JSON-based web APIs. It is entirely unobtrusive and re-invents no wheel. It is crafted such that getting started is easy and quick while being flexible enough for more complex use cases.   [![godoc][D]](https://godoc.org/github.com/opensaucerer/barf)
- [rux](https://github.com/gookit/rux) **star:100** Simple and fast web framework for build golang HTTP applications.   [![There was an update last month][G]](https://github.com/gookit/rux)   [![godoc][D]](https://godoc.org/github.com/gookit/rux)   [![Contains Chinese documents][CN]](https://github.com/gookit/rux)
- [Xun](https://github.com/yaitoo/xun) **star:92** Web framework built on Go's built-in html/template and net/http package’s router. It is designed to be lightweight, fast, and easy to use while providing a simple and intuitive API for building web applications with advanced features such as middleware, routing, and template rendering.   [![There was an update last month][G]](https://github.com/yaitoo/xun)   [![godoc][D]](https://godoc.org/github.com/yaitoo/xun)
- [Don](https://github.com/abemedia/go-don) **star:60** A highly performant and simple to use API framework.   [![godoc][D]](https://godoc.org/github.com/abemedia/go-don)
- [httpsuite](https://github.com/rluders/httpsuite) **star:43** HTTP request parsing and RFC 9457 problem responses for Go, with a stdlib-only core and optional validation.   [![godoc][D]](https://godoc.org/github.com/rluders/httpsuite)
- [Huma](https://github.com/danielgtaylor/huma/)  Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI.
- [Ronykit](https://github.com/clubpay/ronykit) **star:37** Web framework with pluggable architecture and very performant.   [![There was an update last month][G]](https://github.com/clubpay/ronykit)   [![godoc][D]](https://godoc.org/github.com/clubpay/ronykit)
- [go-api-boot](https://github.com/SaiNageswarS/go-api-boot) **star:35** A gRpc-first micro-service framework. Features include ODM support for Mongo, cloud resource support (AWS/Azure/Google), and a fluent dependency injection which is customized for gRpc. Additionally, grpc-web is supported directly, enabling browser access to all gRpc APIs without a proxy.   [![godoc][D]](https://godoc.org/github.com/SaiNageswarS/go-api-boot)
- [Lit](https://github.com/jvcoutinho/lit) **star:30** Highly performant declarative web framework for Golang, aiming for simplicity and quality of life.   [![godoc][D]](https://godoc.org/github.com/jvcoutinho/lit)
- [NotNet](https://github.com/nottechdm/notnet) **star:3** A lightweight Go framework for building fast, ergonomic RESTful APIs with middleware and flexible routing.   [![godoc][D]](https://godoc.org/github.com/nottechdm/notnet)
- [togo](https://github.com/togo-framework/togo) **star:3** Full-stack framework that ships your Go backend and React frontend as a single binary; a Laravel-artisan-grade CLI.   [![godoc][D]](https://godoc.org/github.com/togo-framework/togo)

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [CORS](https://github.com/rs/cors) **star:2895** Easily add CORS capabilities to your API.   [![godoc][D]](https://godoc.org/github.com/rs/cors)
- [Tollbooth](https://github.com/didip/tollbooth) **star:2859** Rate limit HTTP request handler.   [![godoc][D]](https://godoc.org/github.com/didip/tollbooth)
- [Limiter](https://github.com/ulule/limiter) **star:2340** Dead simple rate limit middleware for Go.   [![godoc][D]](https://godoc.org/github.com/ulule/limiter)
- [go-fault](https://github.com/github/go-fault) **star:555** Fault injection middleware for Go.   [![There was an update last month][G]](https://github.com/github/go-fault)   [![godoc][D]](https://godoc.org/github.com/github/go-fault)
- [ln-paywall](https://github.com/philippgille/ln-paywall) **star:157** Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).   [![It hasn't been updated in recent three years][Y]](https://github.com/philippgille/ln-paywall)   [![godoc][D]](https://godoc.org/github.com/philippgille/ln-paywall)
- [XFF](https://github.com/sebest/xff) **star:100** Handle `X-Forwarded-For` header and friends.   [![It hasn't been updated in recent three years][Y]](https://github.com/sebest/xff)   [![godoc][D]](https://godoc.org/github.com/sebest/xff)
- [rk-grpc](https://github.com/rookie-ninja/rk-grpc) **star:81** Middleware for gRPC with logging, metrics, auth, tracing etc.   [![godoc][D]](https://godoc.org/github.com/rookie-ninja/rk-grpc)
- [rk-gin](https://github.com/rookie-ninja/rk-gin) **star:51** Middleware for Gin framework with logging, metrics, auth, tracing etc.   [![godoc][D]](https://godoc.org/github.com/rookie-ninja/rk-gin)
- [formjson](https://github.com/rs/formjson) **star:38** Transparently handle JSON input as a standard form POST.   [![It hasn't been updated in recent three years][Y]](https://github.com/rs/formjson)   [![godoc][D]](https://godoc.org/github.com/rs/formjson)
- [client-timing](https://github.com/posener/client-timing) **star:25** An HTTP client for Server-Timing header.   [![It hasn't been updated in recent three years][Y]](https://github.com/posener/client-timing)   [![godoc][D]](https://godoc.org/github.com/posener/client-timing)
- [echo-middleware](https://github.com/faabiosr/echo-middleware) **star:16** Middleware for Echo framework with logging and metrics.   [![godoc][D]](https://godoc.org/github.com/faabiosr/echo-middleware)
- [mid](https://github.com/bobg/mid) **star:10** Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more.   [![godoc][D]](https://godoc.org/github.com/bobg/mid)

#### Libraries for creating HTTP middlewares

- [negroni](https://github.com/urfave/negroni) **star:7529** Idiomatic HTTP middleware for Golang.   [![godoc][D]](https://godoc.org/github.com/urfave/negroni)   [![Contains Chinese documents][CN]](https://github.com/urfave/negroni)
- [alice](https://github.com/justinas/alice) **star:3360** Painless middleware chaining for Go.   [![godoc][D]](https://godoc.org/github.com/justinas/alice)
- [render](https://github.com/unrolled/render) **star:1997** Go package for easily rendering JSON, XML, and HTML template responses.   [![godoc][D]](https://godoc.org/github.com/unrolled/render)
- [stats](https://github.com/thoas/stats) **star:594** Go middleware that stores various information about your web application.   [![It hasn't been updated in recent three years][Y]](https://github.com/thoas/stats)   [![godoc][D]](https://godoc.org/github.com/thoas/stats)
- [interpose](https://github.com/carbocation/interpose) **star:291** Minimalist net/http middleware for golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/carbocation/interpose)   [![godoc][D]](https://godoc.org/github.com/carbocation/interpose)
- [renderer](https://github.com/thedevsaddam/renderer) **star:261** Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.   [![godoc][D]](https://godoc.org/github.com/thedevsaddam/renderer)
- [muxchain](https://github.com/stephens2424/muxchain) **star:206** Lightweight middleware for net/http.   [![It hasn't been updated in recent three years][Y]](https://github.com/stephens2424/muxchain)   [![godoc][D]](https://godoc.org/github.com/stephens2424/muxchain)
- [gores](https://github.com/alioygur/gores) **star:106** Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/alioygur/gores)   [![godoc][D]](https://godoc.org/github.com/alioygur/gores)
- [mediary](https://github.com/HereMobilityDevelopers/mediary) **star:88** add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses.   [![It hasn't been updated in recent three years][Y]](https://github.com/HereMobilityDevelopers/mediary)   [![godoc][D]](https://godoc.org/github.com/HereMobilityDevelopers/mediary)
- [chain](https://github.com/codemodus/chain) **star:61** Handler wrapper chaining with scoped data (net/context-based "middleware").   [![It hasn't been updated in recent three years][Y]](https://github.com/codemodus/chain)   [![godoc][D]](https://godoc.org/github.com/codemodus/chain)
- [catena](https://github.com/codemodus/catena) **star:9** http.Handler wrapper catenation (same API as "chain").   [![It hasn't been updated in recent three years][Y]](https://github.com/codemodus/catena)   [![godoc][D]](https://godoc.org/github.com/codemodus/catena)

**[⬆ back to top](#contents)**

### Routers

- [chi](https://github.com/go-chi/chi) **star:22538** Small, fast and expressive HTTP router built on net/context.   [![godoc][D]](https://godoc.org/github.com/go-chi/chi)
- [mux](https://github.com/gorilla/mux) **star:21846** Powerful URL router and dispatcher for golang.   [![godoc][D]](https://godoc.org/github.com/gorilla/mux)
- [httprouter](https://github.com/julienschmidt/httprouter) **star:17135** High performance router. Use this and the standard http handlers to form a very high performance web framework.   [![godoc][D]](https://godoc.org/github.com/julienschmidt/httprouter)
- [gocraft/web](https://github.com/gocraft/web) **star:1523** Mux and middleware package in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/gocraft/web)   [![godoc][D]](https://godoc.org/github.com/gocraft/web)
- [Bone](https://github.com/go-zoo/bone) **star:1284** Lightning Fast HTTP Multiplexer.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-zoo/bone)   [![godoc][D]](https://godoc.org/github.com/go-zoo/bone)
- [Goji](https://github.com/goji/goji) **star:974** Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.   [![It hasn't been updated in recent three years][Y]](https://github.com/goji/goji)   [![godoc][D]](https://godoc.org/github.com/goji/goji)
- [fasthttprouter](https://github.com/buaazp/fasthttprouter) **star:869** High performance router forked from `httprouter`. The first router fit for `fasthttp`.   [![It hasn't been updated in recent three years][Y]](https://github.com/buaazp/fasthttprouter)   [![godoc][D]](https://godoc.org/github.com/buaazp/fasthttprouter)
- [httptreemux](https://github.com/dimfeld/httptreemux) **star:618** High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.   [![godoc][D]](https://godoc.org/github.com/dimfeld/httptreemux)
- [xujiajun/gorouter](https://github.com/xujiajun/gorouter) **star:525** A simple and fast HTTP router for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/xujiajun/gorouter)   [![godoc][D]](https://godoc.org/github.com/xujiajun/gorouter)
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) **star:452** An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-ozzo/ozzo-routing)   [![godoc][D]](https://godoc.org/github.com/go-ozzo/ozzo-routing)
- [lars](https://github.com/go-playground/lars) **star:385** Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/lars)   [![godoc][D]](https://godoc.org/github.com/go-playground/lars)
- [Siesta](https://github.com/VividCortex/siesta) **star:347** Composable framework to write middleware and handlers.   [![godoc][D]](https://godoc.org/github.com/VividCortex/siesta)
- [vestigo](https://github.com/husobee/vestigo) **star:265** Performant, stand-alone, HTTP compliant URL Router for go web applications.   [![It hasn't been updated in recent three years][Y]](https://github.com/husobee/vestigo)   [![godoc][D]](https://godoc.org/github.com/husobee/vestigo)
- [gowww/router](https://github.com/gowww/router) **star:185** Lightning fast HTTP router fully compatible with the net/http.Handler interface.   [![godoc][D]](https://godoc.org/github.com/gowww/router)
- [GoRouter](https://github.com/vardius/gorouter) **star:153** GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.   [![godoc][D]](https://godoc.org/github.com/vardius/gorouter)
- [pure](https://github.com/go-playground/pure) **star:153** Is a lightweight HTTP router that sticks to the std "net/http" implementation.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-playground/pure)   [![godoc][D]](https://godoc.org/github.com/go-playground/pure)
- [alien](https://github.com/gernest/alien) **star:134** Lightweight and fast http router from outer space.   [![godoc][D]](https://godoc.org/github.com/gernest/alien)
- [violetear](https://github.com/nbari/violetear) **star:107** Go HTTP router.   [![It hasn't been updated in recent three years][Y]](https://github.com/nbari/violetear)   [![godoc][D]](https://godoc.org/github.com/nbari/violetear)   [![Archived][Archived]](https://github.com/nbari/violetear)
- [Bxog](https://github.com/claygod/Bxog) **star:104** Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.   [![godoc][D]](https://godoc.org/github.com/claygod/Bxog)
- [xmux](https://github.com/rs/xmux) **star:100** High performance muxer based on `httprouter` with `net/context` support.   [![It hasn't been updated in recent three years][Y]](https://github.com/rs/xmux)   [![godoc][D]](https://godoc.org/github.com/rs/xmux)
- [goblin](https://github.com/bmf-san/goblin) **star:82** A golang http router based on trie tree.   [![godoc][D]](https://godoc.org/github.com/bmf-san/goblin)
- [ngamux](https://github.com/ngamux/ngamux) **star:71** Simple HTTP router for Go.   [![godoc][D]](https://godoc.org/github.com/ngamux/ngamux)
- [bellt](https://github.com/GuilhermeCaruso/bellt) **star:55** A simple Go HTTP router.   [![It hasn't been updated in recent three years][Y]](https://github.com/GuilhermeCaruso/bellt)   [![godoc][D]](https://godoc.org/github.com/GuilhermeCaruso/bellt)
- [FastRouter](https://github.com/razonyang/fastrouter) **star:25** a fast, flexible HTTP router written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/razonyang/fastrouter)   [![godoc][D]](https://godoc.org/github.com/razonyang/fastrouter)
- [GoLobby/Router](https://github.com/golobby/router) **star:22** GoLobby Router is a lightweight yet powerful HTTP router for the Go programming language.   [![It hasn't been updated in recent three years][Y]](https://github.com/golobby/router)   [![godoc][D]](https://godoc.org/github.com/golobby/router)
- [nchi](https://github.com/muir/nchi) **star:19** chi-like router built on httprouter with dependency injection based middleware wrappers   [![godoc][D]](https://godoc.org/github.com/muir/nchi)
- [Fox](https://github.com/fox-toolkit/fox) **star:14** A high-performance HTTP router for building reverse proxies and API gateways, with first-class support for mutating routes at runtime.   [![There was an update last month][G]](https://github.com/fox-toolkit/fox)   [![godoc][D]](https://godoc.org/github.com/fox-toolkit/fox)
- [goroute](https://github.com/goroute/route) **star:9** Simple yet powerful HTTP request multiplexer.   [![It hasn't been updated in recent three years][Y]](https://github.com/goroute/route)   [![godoc][D]](https://godoc.org/github.com/goroute/route)
- [fursy](https://github.com/coregx/fursy) **star:7** HTTP router with type-safe generic handlers, automatic OpenAPI 3.1 generation from code, and RFC 9457 error responses.   [![godoc][D]](https://godoc.org/github.com/coregx/fursy)

**[⬆ back to top](#contents)**

## WebAssembly

- [tinygo](https://github.com/tinygo-org/tinygo) **star:17569** Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.   [![There was an update last month][G]](https://github.com/tinygo-org/tinygo)   [![godoc][D]](https://godoc.org/github.com/tinygo-org/tinygo)
- [dom](https://github.com/dennwc/dom) **star:506** DOM library.   [![It hasn't been updated in recent three years][Y]](https://github.com/dennwc/dom)   [![godoc][D]](https://godoc.org/github.com/dennwc/dom)
- [go-canvas](https://github.com/markfarnan/go-canvas) **star:271** Library to use HTML5 Canvas, with all drawing within go code.   [![It hasn't been updated in recent three years][Y]](https://github.com/markfarnan/go-canvas)   [![godoc][D]](https://godoc.org/github.com/markfarnan/go-canvas)
- [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) **star:211** Run Go WASM tests in your browser.   [![godoc][D]](https://godoc.org/github.com/agnivade/wasmbrowsertest)
- [Extism Go SDK](https://github.com/extism/go-sdk) **star:181** Universal, cross-language WebAssembly framework for building plug-in systems and polyglot apps.   [![godoc][D]](https://godoc.org/github.com/extism/go-sdk)
- [webapi](https://github.com/gowebapi/webapi) **star:181** Bindings for DOM and HTML generated from WebIDL.   [![It hasn't been updated in recent three years][Y]](https://github.com/gowebapi/webapi)   [![godoc][D]](https://godoc.org/github.com/gowebapi/webapi)
- [vert](https://github.com/norunners/vert) **star:108** Interop between Go and JS values.   [![It hasn't been updated in recent three years][Y]](https://github.com/norunners/vert)   [![godoc][D]](https://godoc.org/github.com/norunners/vert)

**[⬆ back to top](#contents)**

## Webhooks Server

- [webhook](https://github.com/adnanh/webhook) **star:11969** Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.   [![godoc][D]](https://godoc.org/github.com/adnanh/webhook)
- [WebhookX](https://github.com/webhookx-io/webhookx) **star:294** A webhooks gateway for message receiving, processing, and reliable delivering.   [![godoc][D]](https://godoc.org/github.com/webhookx-io/webhookx)
- [webhooked](https://github.com/42Atomys/webhooked) **star:43** A webhook receiver on steroids: handle, secure, format and store a Webhook payload has never been easier.   [![godoc][D]](https://godoc.org/github.com/42Atomys/webhooked)
- [HookRun](https://github.com/bluvenr/hookrun) **star:3** Lightweight webhook action engine (~3MB single binary, zero deps) that executes commands and scripts from YAML rules with token/HMAC/IP auth and hot reload.   [![godoc][D]](https://godoc.org/github.com/bluvenr/hookrun)   [![Contains Chinese documents][CN]](https://github.com/bluvenr/hookrun)

**[⬆ back to top](#contents)**

## Windows

- [go-ole](https://github.com/go-ole/go-ole) **star:1316** Win32 OLE implementation for golang.   [![godoc][D]](https://godoc.org/github.com/go-ole/go-ole)
- [d3d9](https://github.com/gonutz/d3d9) **star:164** Go bindings for Direct3D9.   [![godoc][D]](https://godoc.org/github.com/gonutz/d3d9)
- [windowsupdate](https://github.com/ceshihao/windowsupdate) **star:14** A Golang binding for Windows Update Agent API using go-ole.   [![godoc][D]](https://godoc.org/github.com/ceshihao/windowsupdate)
- [gosddl](https://github.com/MonaxGT/gosddl) **star:11** Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.   [![It hasn't been updated in recent three years][Y]](https://github.com/MonaxGT/gosddl)   [![godoc][D]](https://godoc.org/github.com/MonaxGT/gosddl)

**[⬆ back to top](#contents)**

## Workflow Frameworks

_Libraries for creating Workflows._

- [Dagu](https://github.com/dagu-go/dagu) **star:3634** No-code workflow executor. it executes DAGs defined in a simple YAML format.   [![There was an update last month][G]](https://github.com/dagu-go/dagu)   [![godoc][D]](https://godoc.org/github.com/dagu-go/dagu)
- [go-taskflow](https://github.com/noneback/go-taskflow) **star:635** A taskflow-like General-purpose Task-parallel Programming Framework with integrated visualizer and profiler.   [![godoc][D]](https://godoc.org/github.com/noneback/go-taskflow)
- [Cadence-client](https://github.com/uber-go/cadence-client) **star:378** A framework for authoring workflows and activities running on top of the Cadence orchestration engine made by Uber.   [![There was an update last month][G]](https://github.com/uber-go/cadence-client)   [![godoc][D]](https://godoc.org/github.com/uber-go/cadence-client)
- [workflow](https://github.com/luno/workflow) **star:249** A tech stack agnostic Event Driven Workflow framework.   [![godoc][D]](https://godoc.org/github.com/luno/workflow)
- [Flowbaker](https://github.com/flowbaker/flowbaker) **star:203** Self-hosted execution engine for building, connecting, and automating no-code workflows.   [![godoc][D]](https://godoc.org/github.com/flowbaker/flowbaker)
- [go-dag](https://github.com/rhosocial/go-dag) **star:41** A framework developed in Go that manages the execution of workflows described by directed acyclic graphs.   [![godoc][D]](https://godoc.org/github.com/rhosocial/go-dag)

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._

- [zek](https://github.com/miku/zek) **star:821** Generate a Go struct from XML.   [![godoc][D]](https://godoc.org/github.com/miku/zek)
- [xpath](https://github.com/antchfx/xpath) **star:742** XPath package for Go.   [![There was an update last month][G]](https://github.com/antchfx/xpath)   [![godoc][D]](https://godoc.org/github.com/antchfx/xpath)
- [xmlquery](https://github.com/antchfx/xmlquery) **star:490** xmlquery is Golang XPath package for XML query.   [![godoc][D]](https://godoc.org/github.com/antchfx/xmlquery)
- [xml2map](https://github.com/sbabiv/xml2map) **star:65** XML to MAP converter written Golang.   [![It hasn't been updated in recent three years][Y]](https://github.com/sbabiv/xml2map)   [![godoc][D]](https://godoc.org/github.com/sbabiv/xml2map)
- [xmlwriter](https://github.com/shabbyrobe/xmlwriter) **star:30** Procedural XML generation API based on libxml2's xmlwriter module.   [![godoc][D]](https://godoc.org/github.com/shabbyrobe/xmlwriter)   [![Archived][Archived]](https://github.com/shabbyrobe/xmlwriter)
- [XML-Comp](https://github.com/xml-comp/xml-comp) **star:21** Simple command line XML comparer that generates diffs of folders, files and tags.   [![It hasn't been updated in recent three years][Y]](https://github.com/xml-comp/xml-comp)   [![godoc][D]](https://godoc.org/github.com/xml-comp/xml-comp)

## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign](https://github.com/sigstore/cosign) **star:6127** Container Signing, Verification and Storage in an OCI registry.   [![There was an update last month][G]](https://github.com/sigstore/cosign)   [![godoc][D]](https://godoc.org/github.com/sigstore/cosign)
- [OpenZiti](https://github.com/openziti/ziti) **star:4290** A full, open source zero trust overlay network. Including numerous SDKs for numerous languages such as [golang](https://github.com/openziti/sdk-golang) allowing you to embed zero trust principles directly into your applications. The [OpenZiti Test Kitchen](https://github.com/openziti-test-kitchen) has numerous examples to draw inspiration from including a [zero trust ssh client - zssh](https://github.com/openziti-test-kitchen/zssh)   [![There was an update last month][G]](https://github.com/openziti/ziti)   [![godoc][D]](https://godoc.org/github.com/openziti/ziti)
- [Spire](https://github.com/spiffe/spire) **star:2434** SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms.   [![There was an update last month][G]](https://github.com/spiffe/spire)   [![godoc][D]](https://godoc.org/github.com/spiffe/spire)
- [in-toto](https://github.com/in-toto/in-toto-golang) **star:151** Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation.   [![There was an update last month][G]](https://github.com/in-toto/in-toto-golang)   [![godoc][D]](https://godoc.org/github.com/in-toto/in-toto-golang)
- [Spiffe-Vault](https://github.com/philips-labs/spiffe-vault) **star:101** Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication.   [![godoc][D]](https://godoc.org/github.com/philips-labs/spiffe-vault)

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [errcheck](https://github.com/kisielk/errcheck) **star:2525** Errcheck is a program for checking for unchecked errors in Go programs.   [![godoc][D]](https://godoc.org/github.com/kisielk/errcheck)
- [go-critic](https://github.com/go-critic/go-critic) **star:2057** source code linter that brings checks that are currently not implemented in other linters.   [![godoc][D]](https://godoc.org/github.com/go-critic/go-critic)
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) **star:982** go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.   [![It hasn't been updated in recent three years][Y]](https://github.com/roblaszczak/go-cleanarch)   [![godoc][D]](https://godoc.org/github.com/roblaszczak/go-cleanarch)
- [goast-viewer](https://github.com/yuroyoro/goast-viewer) **star:793** Web based Golang AST visualizer.
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)  Tool to fix (add, remove) your Go imports automatically.
- [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) **star:671** An easy way to find outdated dependencies of your Go projects.   [![It hasn't been updated in recent three years][Y]](https://github.com/psampaz/go-mod-outdated)   [![godoc][D]](https://godoc.org/github.com/psampaz/go-mod-outdated)
- [Chronos](https://github.com/amit-davidson/Chronos) **star:441** Detects race conditions statically   [![It hasn't been updated in recent three years][Y]](https://github.com/amit-davidson/Chronos)   [![godoc][D]](https://godoc.org/github.com/amit-davidson/Chronos)
- [dupl](https://github.com/mibk/dupl) **star:368** Tool for code clone detection.   [![godoc][D]](https://godoc.org/github.com/mibk/dupl)
- [apicompat](https://github.com/bradleyfalzon/apicompat) **star:182** Checks recent changes to a Go project for backwards incompatible changes.   [![It hasn't been updated in recent three years][Y]](https://github.com/bradleyfalzon/apicompat)   [![godoc][D]](https://godoc.org/github.com/bradleyfalzon/apicompat)
- [go-checkstyle](https://github.com/qiniu/checkstyle) **star:130** checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.   [![It hasn't been updated in recent three years][Y]](https://github.com/qiniu/checkstyle)   [![godoc][D]](https://godoc.org/github.com/qiniu/checkstyle)
- [asty](https://github.com/asty-org/asty) **star:88** Converts golang AST to JSON and JSON to AST.   [![It hasn't been updated in recent three years][Y]](https://github.com/asty-org/asty)   [![godoc][D]](https://godoc.org/github.com/asty-org/asty)
- [blanket](https://gitlab.com/verygoodsoftwarenotvirus/blanket)  blanket is a tool that helps you catch functions which don't have direct unit tests in your Go packages.
- [fatcontext](https://github.com/Crocmagnon/fatcontext) **star:80** Fatcontext detects nested contexts in loops or function literals.   [![godoc][D]](https://godoc.org/github.com/Crocmagnon/fatcontext)
- [ChainJacking](https://github.com/Checkmarx/chainjacking) **star:64** Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack.
- [golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) **star:13** iFood API SDK.   [![It hasn't been updated in recent three years][Y]](https://github.com/arxdsilva/golang-ifood-sdk)   [![godoc][D]](https://godoc.org/github.com/arxdsilva/golang-ifood-sdk)
- [golangci-lint](https://github.com/golangci/golangci-lint) – A fast Go linters runner. It runs linters in parallel, uses caching, supports `yaml` config, has integrations with all major IDE and has dozens of linters included.
- [GoPlantUML](https://github.com/jfeliu007/goplantuml) **star:2093** Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.   [![godoc][D]](https://godoc.org/github.com/jfeliu007/goplantuml)
- [golines](https://github.com/segmentio/golines) **star:1133** Formatter that automatically shortens long lines in Go code.   [![godoc][D]](https://godoc.org/github.com/segmentio/golines)   [![Archived][Archived]](https://github.com/segmentio/golines)
- [php-parser](https://github.com/z7zmey/php-parser) **star:959** A Parser for PHP written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/z7zmey/php-parser)   [![godoc][D]](https://godoc.org/github.com/z7zmey/php-parser)
- [goreturns](https://github.com/sqs/goreturns) **star:538** Adds zero-value return statements to match the func return types.   [![godoc][D]](https://godoc.org/github.com/sqs/goreturns)
- [gostatus](https://github.com/shurcooL/gostatus) **star:245** Command line tool, shows the status of repositories that contain Go packages.   [![godoc][D]](https://godoc.org/github.com/shurcooL/gostatus)
- [lint](https://github.com/surullabs/lint) **star:66** Run linters as part of go test.   [![It hasn't been updated in recent three years][Y]](https://github.com/surullabs/lint)   [![godoc][D]](https://godoc.org/github.com/surullabs/lint)
- [gomarklint](https://github.com/shinagawa-web/gomarklint) **star:19** Markdown linter with built-in HTTP link validation, single binary, no Node.js required.   [![There was an update last month][G]](https://github.com/shinagawa-web/gomarklint)   [![godoc][D]](https://godoc.org/github.com/shinagawa-web/gomarklint)
- [revive](https://github.com/mgechev/revive) – ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for `golint`.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)  staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
- [stto](https://github.com/mainak55512/stto) **star:47** A light-weight superfast line of code counter written in pure Go.   [![godoc][D]](https://godoc.org/github.com/mainak55512/stto)
- [structalign](https://github.com/peczenyj/structalign) **star:9** Shows how a struct's fields could be reordered to use less memory, printing a diff instead of rewriting files.   [![There was an update last month][G]](https://github.com/peczenyj/structalign)   [![godoc][D]](https://godoc.org/github.com/peczenyj/structalign)
- [testifylint](https://github.com/Antonboom/testifylint) – A linter that checks usage of [github.com/stretchr/testify](https://github.com/stretchr/testify).
- [vacuum](https://github.com/daveshanley/vacuum) **star:1100** An ultra-super-fast, lightweight OpenAPI linter and quality checking tool.   [![There was an update last month][G]](https://github.com/daveshanley/vacuum)   [![godoc][D]](https://godoc.org/github.com/daveshanley/vacuum)
- [todocheck](https://github.com/preslavmihaylov/todocheck) **star:439** Static code analyser which links TODO comments in code with issues in your issue tracker.   [![godoc][D]](https://godoc.org/github.com/preslavmihaylov/todocheck)
- [unconvert](https://github.com/mdempsky/unconvert) **star:388** Remove unnecessary type conversions from Go source.   [![godoc][D]](https://godoc.org/github.com/mdempsky/unconvert)
- [wrapcheck](https://github.com/tomarrell/wrapcheck) **star:375** A linter to check that errors from external packages are wrapped.   [![godoc][D]](https://godoc.org/github.com/tomarrell/wrapcheck)
- [tickgit](https://github.com/augmentable-dev/tickgit) **star:324** CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.   [![godoc][D]](https://godoc.org/github.com/augmentable-dev/tickgit)
- [validate](https://github.com/mccoyst/validate) **star:62** Automatically validates struct fields with tags.   [![godoc][D]](https://godoc.org/github.com/mccoyst/validate)
- [usestdlibvars](https://github.com/sashamelentyev/usestdlibvars) **star:48** A linter that detect the possibility to use variables/constants from the Go standard library.   [![godoc][D]](https://godoc.org/github.com/sashamelentyev/usestdlibvars)

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [vim-go](https://github.com/fatih/vim-go) **star:16226** Go development plugin for Vim.   [![There was an update last month][G]](https://github.com/fatih/vim-go)
- [gocode](https://github.com/nsf/gocode) **star:4995** Autocompletion daemon for the Go programming language.   [![godoc][D]](https://godoc.org/github.com/nsf/gocode)
- [vscode-go](https://github.com/golang/vscode-go) **star:4260** Extension for Visual Studio Code (VS Code) which provides support for the Go language.   [![There was an update last month][G]](https://github.com/golang/vscode-go)
- [GoSublime](https://github.com/DisposaBoy/GoSublime) **star:3404** Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.   [![It hasn't been updated in recent three years][Y]](https://github.com/DisposaBoy/GoSublime)   [![godoc][D]](https://godoc.org/github.com/DisposaBoy/GoSublime)
- [go-mode](https://github.com/dominikh/go-mode.el) **star:1454** Go mode for GNU/Emacs.
- [goimports-reviser](https://github.com/incu6us/goimports-reviser) **star:717** Formatting tool for imports.   [![godoc][D]](https://godoc.org/github.com/incu6us/goimports-reviser)
- [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof)  This extension adds benchmark profiling support for the Go language to VS Code.
- [coc-go language server extension for Vim/Neovim](https://github.com/josa42/coc-go) **star:573** This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim.
- [Watch](https://github.com/eaburns/Watch) **star:201** Runs a command in an acme win on file changes.   [![It hasn't been updated in recent three years][Y]](https://github.com/eaburns/Watch)   [![godoc][D]](https://godoc.org/github.com/eaburns/Watch)
- [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) **star:90** Vim plugin to highlight syntax errors on save.   [![It hasn't been updated in recent three years][Y]](https://github.com/rjohnsondev/vim-compiler-go)
- [gounit-vim](https://github.com/hexdigest/gounit-vim) **star:25** Vim plugin for generating Go tests based on the function's or method's signature.   [![It hasn't been updated in recent three years][Y]](https://github.com/hexdigest/gounit-vim)
- [Go Doc](https://github.com/msyrus/vscode-go-doc) **star:9** A Visual Studio Code extension for showing definition in output and generating go doc.
- [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)  Go plugin for JetBrains IDEs.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [gotests](https://github.com/cweill/gotests) **star:5328** Generate Go tests from your source code.   [![godoc][D]](https://godoc.org/github.com/cweill/gotests)
- [xgen](https://github.com/xuri/xgen) **star:417** XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator.   [![godoc][D]](https://godoc.org/github.com/xuri/xgen)
- [hasgo](https://github.com/DylanMeeus/hasgo) **star:144** Generate Haskell inspired functions for your slices.   [![It hasn't been updated in recent three years][Y]](https://github.com/DylanMeeus/hasgo)   [![godoc][D]](https://godoc.org/github.com/DylanMeeus/hasgo)
- [gocontracts](https://github.com/Parquery/gocontracts) **star:119** brings design-by-contract to Go by synchronizing the code with the documentation.   [![It hasn't been updated in recent three years][Y]](https://github.com/Parquery/gocontracts)   [![godoc][D]](https://godoc.org/github.com/Parquery/gocontracts)
- [gonerics](https://github.com/bouk/gonerics) **star:112** Idiomatic Generics in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/bouk/gonerics)   [![godoc][D]](https://godoc.org/github.com/bouk/gonerics)
- [options-gen](https://github.com/kazhuravlev/options-gen) **star:109** Functional options described by Dave Cheney's post "Functional options for friendly APIs".   [![godoc][D]](https://godoc.org/github.com/kazhuravlev/options-gen)
- [re2dfa](https://gitlab.com/opennota/re2dfa)  Transform regular expressions into finite state machines and output Go source code.
- [envdoc](https://github.com/g4s8/envdoc) **star:98** generate documentation for environment variables from Go source files.   [![There was an update last month][G]](https://github.com/g4s8/envdoc)   [![godoc][D]](https://godoc.org/github.com/g4s8/envdoc)
- [gounit](https://github.com/hexdigest/gounit) **star:87** Generate Go tests using your own templates.   [![It hasn't been updated in recent three years][Y]](https://github.com/hexdigest/gounit)   [![godoc][D]](https://godoc.org/github.com/hexdigest/gounit)
- [sqlgen](https://github.com/anqiansong/sqlgen) **star:85** Generate gorm, xorm, sqlx, bun, sql code from SQL file or DSN.   [![It hasn't been updated in recent three years][Y]](https://github.com/anqiansong/sqlgen)   [![godoc][D]](https://godoc.org/github.com/anqiansong/sqlgen)   [![Contains Chinese documents][CN]](https://github.com/anqiansong/sqlgen)
- [TOML-to-Go](https://xuri.me/toml-to-go)  Translates TOML into a Go type in the browser instantly.
- [generic](https://github.com/usk81/generic) **star:49** flexible data type for Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/usk81/generic)   [![godoc][D]](https://godoc.org/github.com/usk81/generic)
- [godal](https://github.com/mafulong/godal) **star:19** Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm.   [![It hasn't been updated in recent three years][Y]](https://github.com/mafulong/godal)   [![godoc][D]](https://godoc.org/github.com/mafulong/godal)

**[⬆ back to top](#contents)**

## Go Tools

- [go-swagger](https://github.com/go-swagger/go-swagger) **star:9997** Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.   [![There was an update last month][G]](https://github.com/go-swagger/go-swagger)   [![godoc][D]](https://godoc.org/github.com/go-swagger/go-swagger)
- [go-template-playground](https://bartventer.github.io/go-template-playground/)  An interactive environment to create and test Go templates.
- [go-callvis](https://github.com/TrueFurby/go-callvis) **star:6501** Visualize call graph of your Go program using dot format.   [![godoc][D]](https://godoc.org/github.com/TrueFurby/go-callvis)
- [OctoLinker](https://github.com/OctoLinker/browser-extension) **star:5376** Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
- [lensm](https://github.com/loov/lensm) **star:3695** Go assembly and source viewer.   [![godoc][D]](https://godoc.org/github.com/loov/lensm)
- [go-size-analyzer](https://github.com/Zxilly/go-size-analyzer) **star:2152** Analyze and visualize the size of dependencies in compiled Golang binaries, providing insight into their impact on the final build.   [![There was an update last month][G]](https://github.com/Zxilly/go-size-analyzer)   [![godoc][D]](https://godoc.org/github.com/Zxilly/go-size-analyzer)   [![Contains Chinese documents][CN]](https://github.com/Zxilly/go-size-analyzer)
- [richgo](https://github.com/kyoh86/richgo) **star:861** Enrich `go test` outputs with text decorations.   [![godoc][D]](https://godoc.org/github.com/kyoh86/richgo)
- [govisual](https://github.com/doganarif/govisual) **star:686** Zero-config, pure-Go HTTP request visualizer & debugger for local Go web development.   [![godoc][D]](https://godoc.org/github.com/doganarif/govisual)
- [MoniGO](https://github.com/iyashjayesh/monigo) **star:409** A performance monitoring library for Go applications. It provides real-time insights into application performance! 🚀
- [rts](https://github.com/galeone/rts) **star:257** RTS: response to struct. Generates Go structs from server responses.   [![It hasn't been updated in recent three years][Y]](https://github.com/galeone/rts)   [![godoc][D]](https://godoc.org/github.com/galeone/rts)
- [roumon](https://github.com/becheran/roumon) **star:236** Monitor current state of all active goroutines via a command line interface.   [![godoc][D]](https://godoc.org/github.com/becheran/roumon)
- [godbg](https://github.com/tylerwince/godbg) **star:207** Implementation of Rusts `dbg!` macro for quick and easy debugging during development.   [![It hasn't been updated in recent three years][Y]](https://github.com/tylerwince/godbg)   [![godoc][D]](https://godoc.org/github.com/tylerwince/godbg)
- [gomodrun](https://github.com/dustinblackman/gomodrun/)  Go tool that executes and caches binaries included in go.mod files.
- [gotemplate.io](https://gotemplate.io/)  Online tool to preview `text/template` templates live.
- [typex](https://github.com/dtgorski/typex) **star:207** Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration.   [![godoc][D]](https://godoc.org/github.com/dtgorski/typex)
- [gotestdox](https://github.com/bitfield/gotestdox) **star:200** Show Go test results as readable sentences.   [![godoc][D]](https://godoc.org/github.com/bitfield/gotestdox)
- [gothanks](https://github.com/psampaz/gothanks) **star:128** GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.   [![It hasn't been updated in recent three years][Y]](https://github.com/psampaz/gothanks)   [![godoc][D]](https://godoc.org/github.com/psampaz/gothanks)
- [gotutor](https://github.com/ahmedakef/gotutor) **star:83** Online Go Debugger & Visualizer.
- [igo](https://github.com/rocketlaunchr/igo) **star:73** An igo to go transpiler (new language features for Go language!)   [![It hasn't been updated in recent three years][Y]](https://github.com/rocketlaunchr/igo)   [![godoc][D]](https://godoc.org/github.com/rocketlaunchr/igo)
- [docs](https://github.com/go-oas/docs) **star:51** Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-oas/docs)   [![godoc][D]](https://godoc.org/github.com/go-oas/docs)   [![Archived][Archived]](https://github.com/go-oas/docs)
- [decouple](https://github.com/bobg/decouple) **star:37** Find “overspecified” function parameters that could be generalized with interface types.   [![godoc][D]](https://godoc.org/github.com/bobg/decouple)
- [modver](https://github.com/bobg/modver) **star:22** Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules.   [![godoc][D]](https://godoc.org/github.com/bobg/modver)
- [textra](https://github.com/ravsii/textra) **star:7** Extract Go struct field names, types and tags for filtering and exporting.   [![It hasn't been updated in recent three years][Y]](https://github.com/ravsii/textra)   [![godoc][D]](https://godoc.org/github.com/ravsii/textra)

**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [kubernetes](https://github.com/kubernetes/kubernetes) **star:123792** Container Cluster Manager from Google.   [![There was an update last month][G]](https://github.com/kubernetes/kubernetes)   [![godoc][D]](https://godoc.org/github.com/kubernetes/kubernetes)
- [Moby](https://github.com/moby/moby) **star:71919** Collaborative project for the container ecosystem to assemble container-based systems.   [![There was an update last month][G]](https://github.com/moby/moby)   [![godoc][D]](https://godoc.org/github.com/moby/moby)
- [traefik](https://github.com/containous/traefik) **star:63989** Reverse proxy and load balancer with support for multiple backends.   [![There was an update last month][G]](https://github.com/containous/traefik)   [![godoc][D]](https://godoc.org/github.com/containous/traefik)
- [Gitea](https://github.com/go-gitea/gitea) **star:56831** Fork of Gogs, entirely community driven.   [![There was an update last month][G]](https://github.com/go-gitea/gitea)   [![godoc][D]](https://godoc.org/github.com/go-gitea/gitea)   [![Contains Chinese documents][CN]](https://github.com/go-gitea/gitea)
- [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)  Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
- [gitl](https://github.com/akomyagin/gitl)  AI review of git commit ranges with risk scoring (low/medium/high), changelog generation, and multi-repo activity digest. GitHub Action included.
- [k9s](https://github.com/derailed/k9s) **star:34124** Kubernetes CLI to manage your clusters in style.   [![There was an update last month][G]](https://github.com/derailed/k9s)   [![godoc][D]](https://godoc.org/github.com/derailed/k9s)
- [k3s](https://github.com/k3s-io/k3s) **star:33468** Lightweight Kubernetes.   [![There was an update last month][G]](https://github.com/k3s-io/k3s)   [![godoc][D]](https://godoc.org/github.com/k3s-io/k3s)
- [minikube](https://github.com/kubernetes/minikube) **star:31952** Run Kubernetes locally.   [![There was an update last month][G]](https://github.com/kubernetes/minikube)   [![godoc][D]](https://godoc.org/github.com/kubernetes/minikube)
- [k6](https://github.com/grafana/k6) **star:31022** A modern load testing tool, using Go and JavaScript.   [![There was an update last month][G]](https://github.com/grafana/k6)   [![godoc][D]](https://godoc.org/github.com/grafana/k6)
- [colima](https://github.com/abiosoft/colima) **star:29916** Container runtimes on macOS (and Linux) with minimal setup.   [![There was an update last month][G]](https://github.com/abiosoft/colima)   [![godoc][D]](https://godoc.org/github.com/abiosoft/colima)
- [Vegeta](https://github.com/tsenart/vegeta) **star:25110** HTTP load testing tool and library. It's over 9000!   [![godoc][D]](https://godoc.org/github.com/tsenart/vegeta)
- [Hey](https://github.com/rakyll/hey) **star:20178** Hey is a tiny program that sends some load to a web application.   [![godoc][D]](https://godoc.org/github.com/rakyll/hey)
- [chaosmonkey](https://github.com/Netflix/chaosmonkey) **star:17017** A resiliency tool that helps applications tolerate random instance failures.   [![godoc][D]](https://godoc.org/github.com/Netflix/chaosmonkey)
- [Packer](https://github.com/mitchellh/packer) **star:15730** Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.   [![There was an update last month][G]](https://github.com/mitchellh/packer)   [![godoc][D]](https://godoc.org/github.com/mitchellh/packer)
- [kind](https://github.com/kubernetes-sigs/kind) **star:15366** Kubernetes IN Docker - local clusters for testing Kubernetes.   [![There was an update last month][G]](https://github.com/kubernetes-sigs/kind)   [![godoc][D]](https://godoc.org/github.com/kubernetes-sigs/kind)
- [kubeshark](https://github.com/kubeshark/kubeshark) **star:11999** API traffic analyzer for Kubernetes, inspired by Wireshark, purposely built for Kubernetes.   [![There was an update last month][G]](https://github.com/kubeshark/kubeshark)   [![godoc][D]](https://godoc.org/github.com/kubeshark/kubeshark)
- [GVM](https://github.com/moovweb/gvm) **star:11672** GVM provides an interface to manage Go versions.
- [Flannel](https://github.com/flannel-io/flannel) **star:9503** Flannel is a network fabric for containers, designed for Kubernetes.   [![There was an update last month][G]](https://github.com/flannel-io/flannel)   [![godoc][D]](https://godoc.org/github.com/flannel-io/flannel)
- [Ddosify](https://github.com/ddosify/ddosify) **star:8526** High-performance load testing tool, written in Golang.   [![godoc][D]](https://godoc.org/github.com/ddosify/ddosify)
- [ko](https://github.com/google/ko) **star:8475** Command line tool for building and deploying Go applications on Kubernetes   [![There was an update last month][G]](https://github.com/google/ko)   [![godoc][D]](https://godoc.org/github.com/google/ko)
- [KubeVela](https://github.com/kubevela/kubevela) **star:7856** Cloud native application delivery.   [![There was an update last month][G]](https://github.com/kubevela/kubevela)   [![godoc][D]](https://godoc.org/github.com/kubevela/kubevela)
- [script](https://github.com/bitfield/script) **star:7012** Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.   [![godoc][D]](https://godoc.org/github.com/bitfield/script)
- [bombardier](https://github.com/codesenberg/bombardier) **star:6804** Fast cross-platform HTTP benchmarking tool.   [![godoc][D]](https://godoc.org/github.com/codesenberg/bombardier)
- [Fleet device management](https://github.com/fleetdm/fleet) **star:6582** Lightweight, programmable telemetry for servers and workstations.   [![There was an update last month][G]](https://github.com/fleetdm/fleet)   [![godoc][D]](https://godoc.org/github.com/fleetdm/fleet)
- [k3d](https://github.com/k3d-io/k3d) **star:6489** Little helper to run CNCF's k3s in Docker.   [![There was an update last month][G]](https://github.com/k3d-io/k3d)   [![godoc][D]](https://godoc.org/github.com/k3d-io/k3d)
- [k0s](https://github.com/k0sproject/k0s) **star:6335** Zero Friction Kubernetes distribution.   [![There was an update last month][G]](https://github.com/k0sproject/k0s)   [![godoc][D]](https://godoc.org/github.com/k0sproject/k0s)
- [podinfo](https://github.com/stefanprodan/podinfo) **star:5948** Podinfo is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes. Podinfo is used by CNCF projects like Flux and Flagger for end-to-end testing and workshops.   [![There was an update last month][G]](https://github.com/stefanprodan/podinfo)   [![godoc][D]](https://godoc.org/github.com/stefanprodan/podinfo)
- [gaia](https://github.com/gaia-pipeline/gaia) **star:5218** Build powerful pipelines in any programming language.   [![godoc][D]](https://godoc.org/github.com/gaia-pipeline/gaia)   [![Archived][Archived]](https://github.com/gaia-pipeline/gaia)
- [tau](https://github.com/taubyte/tau) **star:5085** Easily build Cloud Computing Platforms with features like Serverless WebAssembly Functions, Frontend Hosting, CI/CD, Object Storage, K/V Database, and Pub-Sub Messaging.   [![There was an update last month][G]](https://github.com/taubyte/tau)   [![godoc][D]](https://godoc.org/github.com/taubyte/tau)
- [Pomerium](https://github.com/pomerium/pomerium) **star:4911** Pomerium is an identity-aware access proxy.   [![There was an update last month][G]](https://github.com/pomerium/pomerium)   [![godoc][D]](https://godoc.org/github.com/pomerium/pomerium)
- [kubefwd](https://github.com/txn2/kubefwd) **star:4139** Bulk Kubernetes port forwarding with unique IPs per service for local development.   [![There was an update last month][G]](https://github.com/txn2/kubefwd)   [![godoc][D]](https://godoc.org/github.com/txn2/kubefwd)
- [s5cmd](https://github.com/peak/s5cmd) **star:4118** Blazing fast S3 and local filesystem execution tool.   [![godoc][D]](https://godoc.org/github.com/peak/s5cmd)
- [docker-volume-backup](https://github.com/offen/docker-volume-backup) **star:3761** Backup Docker volumes locally or to any S3, WebDAV, Azure Blob Storage, Dropbox or SSH compatible storage.   [![There was an update last month][G]](https://github.com/offen/docker-volume-backup)   [![godoc][D]](https://godoc.org/github.com/offen/docker-volume-backup)
- [kubeblocks](https://github.com/apecloud/kubeblocks) **star:3080** KubeBlocks is an open-source control plane that runs and manages databases, message queues and other data infrastructure on K8s.   [![There was an update last month][G]](https://github.com/apecloud/kubeblocks)   [![godoc][D]](https://godoc.org/github.com/apecloud/kubeblocks)
- [aptly](https://github.com/aptly-dev/aptly) **star:2853** aptly is a Debian repository management tool.   [![There was an update last month][G]](https://github.com/aptly-dev/aptly)   [![godoc][D]](https://godoc.org/github.com/aptly-dev/aptly)
- [kala](https://github.com/ajvb/kala) **star:2157** Simplistic, modern, and performant job scheduler.   [![godoc][D]](https://godoc.org/github.com/ajvb/kala)
- [ghorg](https://github.com/gabrie30/ghorg) **star:2088** Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket.   [![There was an update last month][G]](https://github.com/gabrie30/ghorg)   [![godoc][D]](https://godoc.org/github.com/gabrie30/ghorg)
- [fac](https://github.com/mkchoi212/fac) **star:1859** Command-line user interface to fix git merge conflicts.   [![godoc][D]](https://godoc.org/github.com/mkchoi212/fac)
- [go-selfupdate](https://github.com/sanbornm/go-selfupdate) **star:1695** Enable your Go applications to self update.   [![godoc][D]](https://godoc.org/github.com/sanbornm/go-selfupdate)
- [StatusOK](https://github.com/sanathp/statusok) **star:1642** Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.   [![It hasn't been updated in recent three years][Y]](https://github.com/sanathp/statusok)   [![godoc][D]](https://godoc.org/github.com/sanathp/statusok)
- [tlm](https://github.com/yusufcanb/tlm) **star:1489** Local cli copilot, powered by CodeLLaMa   [![godoc][D]](https://godoc.org/github.com/yusufcanb/tlm)
- [uTask](https://github.com/ovh/utask) **star:1391** Automation engine that models and executes business processes declared in yaml.   [![godoc][D]](https://godoc.org/github.com/ovh/utask)
- [KubeVPN](https://github.com/kubenetworks/kubevpn) **star:1350** KubeVPN offers a Cloud-Native Dev Environment that seamlessly connects to your Kubernetes cluster network.   [![There was an update last month][G]](https://github.com/kubenetworks/kubevpn)   [![godoc][D]](https://godoc.org/github.com/kubenetworks/kubevpn)   [![Contains Chinese documents][CN]](https://github.com/kubenetworks/kubevpn)
- [PipeCD](https://github.com/pipe-cd/pipecd) **star:1317** A GitOps-style continuous delivery platform that provides consistent deployment and operations experience for any applications.   [![There was an update last month][G]](https://github.com/pipe-cd/pipecd)   [![godoc][D]](https://godoc.org/github.com/pipe-cd/pipecd)
- [KusionStack](https://github.com/KusionStack/kusion) **star:1314** A unified programmable configuration techstack to deliver modern app in 'platform as code' and 'infra as code' approach.   [![godoc][D]](https://godoc.org/github.com/KusionStack/kusion)   [![Contains Chinese documents][CN]](https://github.com/KusionStack/kusion)
- [podman-tui](https://github.com/containers/podman-tui) **star:1169** Terminal UI for Podman management.   [![There was an update last month][G]](https://github.com/containers/podman-tui)   [![godoc][D]](https://godoc.org/github.com/containers/podman-tui)
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r) **star:1143** Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.   [![It hasn't been updated in recent three years][Y]](https://github.com/rlmcpherson/s3gof3r)   [![godoc][D]](https://godoc.org/github.com/rlmcpherson/s3gof3r)
- [skm](https://github.com/TimothyYe/skm) **star:1070** SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!   [![godoc][D]](https://godoc.org/github.com/TimothyYe/skm)
- [kwatch](https://github.com/abahmed/kwatch) **star:1011** Monitor & detect crashes in your Kubernetes(K8s) cluster instantly.   [![There was an update last month][G]](https://github.com/abahmed/kwatch)   [![godoc][D]](https://godoc.org/github.com/abahmed/kwatch)
- [Scaleway-cli](https://github.com/scaleway/scaleway-cli) **star:987** Manage BareMetal Servers from Command Line (as easily as with Docker).   [![There was an update last month][G]](https://github.com/scaleway/scaleway-cli)   [![godoc][D]](https://godoc.org/github.com/scaleway/scaleway-cli)
- [Updatecli](https://github.com/updatecli/updatecli) **star:939** A universal declarative update policy engine.   [![There was an update last month][G]](https://github.com/updatecli/updatecli)   [![godoc][D]](https://godoc.org/github.com/updatecli/updatecli)
- [cassowary](https://github.com/rogerwelin/cassowary) **star:810** Modern cross-platform HTTP load-testing tool written in Go.   [![godoc][D]](https://godoc.org/github.com/rogerwelin/cassowary)   [![Contains Chinese documents][CN]](https://github.com/rogerwelin/cassowary)
- [kool](https://github.com/kool-dev/kool) **star:722** Command line tool for managing Docker environments as an easy way.   [![godoc][D]](https://godoc.org/github.com/kool-dev/kool)
- [alaz](https://github.com/ddosify/alaz) **star:717** Effortless, Low-Overhead, eBPF-based Kubernetes Monitoring.
- [aurora](https://github.com/xuri/aurora) **star:599** Cross-platform web-based Beanstalkd queue server console.   [![It hasn't been updated in recent three years][Y]](https://github.com/xuri/aurora)
- [govvv](https://github.com/ahmetalpbalkan/govvv) **star:539** “go build” wrapper to easily add version information into Go binaries.   [![It hasn't been updated in recent three years][Y]](https://github.com/ahmetalpbalkan/govvv)   [![godoc][D]](https://godoc.org/github.com/ahmetalpbalkan/govvv)
- [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) **star:478** S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).   [![There was an update last month][G]](https://github.com/oxyno-zeta/s3-proxy)   [![godoc][D]](https://godoc.org/github.com/oxyno-zeta/s3-proxy)
- [Pewpew](https://github.com/bengadbois/pewpew) **star:454** Flexible HTTP command line stress tester.   [![godoc][D]](https://godoc.org/github.com/bengadbois/pewpew)
- [gobrew](https://github.com/kevincobain2000/gobrew) **star:426** Go version manager. Super simple tool to install and manage Go versions. Install go without root. Gobrew doesn't require shell rehash.   [![godoc][D]](https://godoc.org/github.com/kevincobain2000/gobrew)
- [jcli](https://github.com/jenkins-zh/jenkins-cli) **star:424** Jenkins CLI allows you manage your Jenkins as an easy way.   [![godoc][D]](https://godoc.org/github.com/jenkins-zh/jenkins-cli)   [![Contains Chinese documents][CN]](https://github.com/jenkins-zh/jenkins-cli)
- [aws-doctor](https://github.com/elC0mpa/aws-doctor) **star:422** Diagnose AWS costs, detect idle resources, and optimize cloud spending directly from your terminal 🩺 ☁️.   [![godoc][D]](https://godoc.org/github.com/elC0mpa/aws-doctor)
- [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) **star:348** Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.   [![godoc][D]](https://godoc.org/github.com/appleboy/easyssh-proxy)
- [lstags](https://github.com/ivanilves/lstags) **star:340** Tool and API to sync Docker images across different registries.   [![It hasn't been updated in recent three years][Y]](https://github.com/ivanilves/lstags)   [![godoc][D]](https://godoc.org/github.com/ivanilves/lstags)
- [gonative](https://github.com/inconshreveable/gonative) **star:339** Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.   [![It hasn't been updated in recent three years][Y]](https://github.com/inconshreveable/gonative)   [![godoc][D]](https://godoc.org/github.com/inconshreveable/gonative)   [![Archived][Archived]](https://github.com/inconshreveable/gonative)
- [trubka](https://github.com/xitonix/trubka) **star:337** A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka.   [![godoc][D]](https://godoc.org/github.com/xitonix/trubka)
- [Mora](https://github.com/emicklei/mora) **star:314** REST server for accessing MongoDB documents and meta data.   [![godoc][D]](https://godoc.org/github.com/emicklei/mora)
- [Balerter](https://github.com/balerter/balerter) **star:310** A self-hosted script-based alerting manager.   [![godoc][D]](https://godoc.org/github.com/balerter/balerter)
- [manssh](https://github.com/xwjdsh/manssh) **star:310** manssh is a command line tool for managing your ssh alias config easily.   [![It hasn't been updated in recent three years][Y]](https://github.com/xwjdsh/manssh)   [![godoc][D]](https://godoc.org/github.com/xwjdsh/manssh)
- [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) **star:283** Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.   [![godoc][D]](https://godoc.org/github.com/dikhan/terraform-provider-openapi)
- [dish](https://github.com/thevxn/dish) **star:280** A lightweight, remotely configurable monitoring service.   [![godoc][D]](https://godoc.org/github.com/thevxn/dish)
- [Docker](https://www.docker.com/)  Open platform for distributed applications for developers and sysadmins.
- [dogo](https://github.com/liudng/dogo) **star:274** Monitoring changes in the source file and automatically compile and run (restart).   [![It hasn't been updated in recent three years][Y]](https://github.com/liudng/dogo)   [![godoc][D]](https://godoc.org/github.com/liudng/dogo)   [![Contains Chinese documents][CN]](https://github.com/liudng/dogo)
- [abbreviate](https://github.com/dnnrly/abbreviate) **star:224** abbreviate is a tool turning long strings in to shorter ones with configurable separators, for example to embed branch names in to deployment stack IDs.   [![godoc][D]](https://godoc.org/github.com/dnnrly/abbreviate)
- [godbg](https://github.com/sirnewton01/godbg) **star:224** Web-based gdb front-end application.   [![It hasn't been updated in recent three years][Y]](https://github.com/sirnewton01/godbg)
- [Gogs](https://gogs.io/)  A Self Hosted Git Service in the Go Programming Language.
- [kcli](https://github.com/cswank/kcli) **star:223** Command line tool for inspecting kafka topics/partitions/messages.   [![It hasn't been updated in recent three years][Y]](https://github.com/cswank/kcli)   [![godoc][D]](https://godoc.org/github.com/cswank/kcli)
- [kepfi](https://github.com/Knuspii/kepfi)  A smart alternative to rm with a recovery bin and storage tracking.
- [Blast](https://github.com/dave/blast) **star:220** A simple tool for API load testing and batch jobs.   [![It hasn't been updated in recent three years][Y]](https://github.com/dave/blast)   [![godoc][D]](https://godoc.org/github.com/dave/blast)
- [gobrew](https://github.com/cryptojuice/gobrew) **star:191** gobrew lets you easily switch between multiple versions of go.   [![It hasn't been updated in recent three years][Y]](https://github.com/cryptojuice/gobrew)
- [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) **star:187** A go library and an executable that produces valid Dockerfiles using various input channels.   [![It hasn't been updated in recent three years][Y]](https://github.com/ozankasikci/dockerfile-generator)   [![godoc][D]](https://godoc.org/github.com/ozankasikci/dockerfile-generator)
- [goma-gateway](https://github.com/jkaninda/goma-gateway) **star:183** A Lightweight API Gateway and Reverse Proxy with declarative config, robust middleware, and support for REST, GraphQL, TCP, UDP, and gRPC.   [![There was an update last month][G]](https://github.com/jkaninda/goma-gateway)   [![godoc][D]](https://godoc.org/github.com/jkaninda/goma-gateway)
- [sigma](https://github.com/go-sigma/sigma) **star:180** OCI-native container image registry, support OCI-native artifact, scan artifact, image build etc.   [![There was an update last month][G]](https://github.com/go-sigma/sigma)   [![godoc][D]](https://godoc.org/github.com/go-sigma/sigma)
- [ostent](https://github.com/ostrost/ostent) **star:178** collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.   [![It hasn't been updated in recent three years][Y]](https://github.com/ostrost/ostent)   [![godoc][D]](https://godoc.org/github.com/ostrost/ostent)
- [drone-scp](https://github.com/appleboy/drone-scp) **star:173** Copy files and artifacts via SSH using a binary, docker or Drone CI.   [![godoc][D]](https://godoc.org/github.com/appleboy/drone-scp)
- [winrm-cli](https://github.com/masterzen/winrm-cli) **star:173** Cli tool to remotely execute commands on Windows machines.   [![It hasn't been updated in recent three years][Y]](https://github.com/masterzen/winrm-cli)   [![godoc][D]](https://godoc.org/github.com/masterzen/winrm-cli)
- [grapes](https://github.com/yaronsumel/grapes) **star:170** Lightweight tool designed to distribute commands over ssh with ease.   [![godoc][D]](https://godoc.org/github.com/yaronsumel/grapes)
- [tf-profile](https://github.com/datarootsio/tf-profile) **star:163** Profiler for Terraform runs. Generate global stats, resource-level stats or visualizations.   [![godoc][D]](https://godoc.org/github.com/datarootsio/tf-profile)
- [decompose](https://github.com/s0rg/decompose) **star:138** tool to generate and process Docker containers connections graphs.   [![godoc][D]](https://godoc.org/github.com/s0rg/decompose)
- [go-rocket-update](https://github.com/mouuff/go-rocket-update) **star:128** A simple way to make self updating Go applications - Supports Github and Gitlab.   [![godoc][D]](https://godoc.org/github.com/mouuff/go-rocket-update)
- [Mantil](https://github.com/mantil-io/mantil) **star:114** Go specific framework for building serverless applications on AWS that enables you to focus on pure Go code while Mantil takes care of the infrastructure.   [![It hasn't been updated in recent three years][Y]](https://github.com/mantil-io/mantil)   [![godoc][D]](https://godoc.org/github.com/mantil-io/mantil)
- [go-furnace](https://github.com/go-furnace/go-furnace) **star:99** Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.   [![It hasn't been updated in recent three years][Y]](https://github.com/go-furnace/go-furnace)   [![godoc][D]](https://godoc.org/github.com/go-furnace/go-furnace)
- [Dropship](https://github.com/chrismckenzie/dropship) **star:66** Tool for deploying code via cdn.   [![It hasn't been updated in recent three years][Y]](https://github.com/chrismckenzie/dropship)   [![godoc][D]](https://godoc.org/github.com/chrismckenzie/dropship)
- [docker-go-mingw](https://github.com/x1unix/docker-go-mingw) **star:55** Docker image for building Go binaries for Windows with MinGW toolchain.
- [httpref](https://github.com/dnnrly/httpref) **star:44** httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports.   [![godoc][D]](https://godoc.org/github.com/dnnrly/httpref)
- [drone-jenkins](https://github.com/appleboy/drone-jenkins) **star:43** Trigger downstream Jenkins jobs using a binary, docker or Drone CI.   [![godoc][D]](https://godoc.org/github.com/appleboy/drone-jenkins)
- [awsenv](https://github.com/soniah/awsenv) **star:35** Small binary that loads Amazon (AWS) environment variables for a profile.   [![It hasn't been updated in recent three years][Y]](https://github.com/soniah/awsenv)   [![godoc][D]](https://godoc.org/github.com/soniah/awsenv)
- [Rodent](https://github.com/alouche/rodent) **star:33** Rodent helps you manage Go versions, projects and track dependencies.   [![It hasn't been updated in recent three years][Y]](https://github.com/alouche/rodent)
- [lwc](https://github.com/timdp/lwc) **star:32** A live-updating version of the UNIX wc command.   [![It hasn't been updated in recent three years][Y]](https://github.com/timdp/lwc)   [![godoc][D]](https://godoc.org/github.com/timdp/lwc)
- [DepCharge](https://github.com/centerorbit/depcharge) **star:23** Helps orchestrating the execution of commands across the many dependencies in larger projects.   [![It hasn't been updated in recent three years][Y]](https://github.com/centerorbit/depcharge)   [![godoc][D]](https://godoc.org/github.com/centerorbit/depcharge)
- [wait-for](https://github.com/dnnrly/wait-for) **star:17** Wait for something to happen (from the command line) before continuing. Easy orchestration of Docker services and other things.   [![It hasn't been updated in recent three years][Y]](https://github.com/dnnrly/wait-for)   [![godoc][D]](https://godoc.org/github.com/dnnrly/wait-for)
- [Wide](https://wide.b3log.org/login)  Web-based IDE for Teams using Golang.
- [Den](https://github.com/us/den) **star:10** Self-hosted sandbox runtime for AI agents. Open-source E2B alternative.   [![godoc][D]](https://godoc.org/github.com/us/den)   [![Contains Chinese documents][CN]](https://github.com/us/den)
- [sg](https://github.com/ChristopherRabotin/sg) **star:8** Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.   [![It hasn't been updated in recent three years][Y]](https://github.com/ChristopherRabotin/sg)   [![godoc][D]](https://godoc.org/github.com/ChristopherRabotin/sg)
- [pingtower](https://github.com/crleonard/pingtower) **star:7** Lightweight self-hosted uptime monitor for websites and APIs.   [![godoc][D]](https://godoc.org/github.com/crleonard/pingtower)
- [tickstem/uptime](https://github.com/tickstem/uptime) **star:2** Go client for HTTP uptime monitoring with SSL expiry alerts and configurable response assertions.   [![godoc][D]](https://godoc.org/github.com/tickstem/uptime)
- [zerohand](https://github.com/nilpoona/zerohand) **star:1** A simple and efficient load testing tool for Web APIs.   [![godoc][D]](https://godoc.org/github.com/nilpoona/zerohand)

**[⬆ back to top](#contents)**

### Other Software

- [croc](https://github.com/schollz/croc) **star:35543** Easily and securely send files or folders from one computer to another.   [![There was an update last month][G]](https://github.com/schollz/croc)   [![godoc][D]](https://godoc.org/github.com/schollz/croc)
- [restic](https://github.com/restic/restic) **star:35011** De-duplicating backup program.   [![There was an update last month][G]](https://github.com/restic/restic)   [![godoc][D]](https://godoc.org/github.com/restic/restic)
- [Gor](https://github.com/buger/gor) **star:19293** Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.   [![godoc][D]](https://godoc.org/github.com/buger/gor)
- [JuiceFS](https://github.com/juicedata/juicefs) **star:14199** Distributed POSIX file system built on top of Redis and AWS S3.   [![There was an update last month][G]](https://github.com/juicedata/juicefs)   [![godoc][D]](https://godoc.org/github.com/juicedata/juicefs)
- [Juju](https://jujucharms.com/)  Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
- [Layli](https://layli.app)  Draw pretty layout diagrams as code.
- [toxiproxy](https://github.com/shopify/toxiproxy) **star:12155** Proxy to simulate network and system conditions for automated tests.   [![There was an update last month][G]](https://github.com/shopify/toxiproxy)   [![godoc][D]](https://godoc.org/github.com/shopify/toxiproxy)
- [tsuru](https://tsuru.io/)  Extensible and open source Platform as a Service software.
- [Comcast](https://github.com/tylertreat/Comcast) **star:10509** Simulate bad network connections.   [![godoc][D]](https://godoc.org/github.com/tylertreat/Comcast)
- [scc](https://github.com/boyter/scc) **star:8532** Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.   [![There was an update last month][G]](https://github.com/boyter/scc)   [![godoc][D]](https://godoc.org/github.com/boyter/scc)
- [confd](https://github.com/kelseyhightower/confd) **star:8422** Manage local application configuration files using templates and data from etcd or consul.   [![godoc][D]](https://godoc.org/github.com/kelseyhightower/confd)
- [LiteIDE](https://github.com/visualfc/liteide) **star:7763** LiteIDE is a simple, open source, cross-platform Go IDE.   [![Contains Chinese documents][CN]](https://github.com/visualfc/liteide)
- [Backrest](https://github.com/garethgeorge/backrest) **star:6910** Web-based UI and orchestrator for restic backup.   [![There was an update last month][G]](https://github.com/garethgeorge/backrest)
- [Better Go Playground](https://goplay.tools)  Go playground with syntax highlight, code completion and other features.
- [blocky](https://github.com/0xERR0R/blocky) **star:6783** Fast and lightweight DNS proxy as ad-blocker for local network with many features.   [![There was an update last month][G]](https://github.com/0xERR0R/blocky)   [![godoc][D]](https://godoc.org/github.com/0xERR0R/blocky)
- [drive](https://github.com/odeke-em/drive) **star:6731** Google Drive client for the commandline.   [![godoc][D]](https://godoc.org/github.com/odeke-em/drive)
- [Duplicacy](https://github.com/gilbertchen/duplicacy) **star:5659** A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.   [![godoc][D]](https://godoc.org/github.com/gilbertchen/duplicacy)
- [nes](https://github.com/fogleman/nes) **star:5651** Nintendo Entertainment System (NES) emulator written in Go.   [![godoc][D]](https://godoc.org/github.com/fogleman/nes)
- [Gokapi](https://github.com/Forceu/gokapi) **star:2801** Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload.   [![godoc][D]](https://godoc.org/github.com/Forceu/gokapi)
- [GoLand](https://jetbrains.com/go)  Full featured cross-platform Go IDE.
- [myLG](https://github.com/mehrdadrad/mylg) **star:2718** Command Line Network Diagnostic tool written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/mehrdadrad/mylg)   [![godoc][D]](https://godoc.org/github.com/mehrdadrad/mylg)
- [GoBoy](https://github.com/Humpheh/goboy) **star:2638** Nintendo Game Boy Color emulator written in Go.   [![godoc][D]](https://godoc.org/github.com/Humpheh/goboy)
- [Stack Up](https://github.com/pressly/sup) **star:2515** Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.   [![godoc][D]](https://godoc.org/github.com/pressly/sup)
- [lgo](https://github.com/yunabe/lgo) **star:2453** Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.   [![It hasn't been updated in recent three years][Y]](https://github.com/yunabe/lgo)   [![godoc][D]](https://godoc.org/github.com/yunabe/lgo)
- [Documize](https://github.com/documize/community) **star:2413** Modern wiki software that integrates data from SaaS tools.
- [sonic](https://github.com/go-sonic/sonic) **star:2113** Sonic is a Go Blogging Platform. Simple and Powerful.   [![godoc][D]](https://godoc.org/github.com/go-sonic/sonic)   [![Contains Chinese documents][CN]](https://github.com/go-sonic/sonic)
- [Circuit](https://github.com/gocircuit/circuit) **star:1977** Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.   [![godoc][D]](https://godoc.org/github.com/gocircuit/circuit)
- [tldx](https://github.com/brandonyoungdev/tldx) **star:1893** Bulk domain availability checker using RDAP, DNS, and WHOIS fallback with keyword permutation generation.   [![There was an update last month][G]](https://github.com/brandonyoungdev/tldx)   [![godoc][D]](https://godoc.org/github.com/brandonyoungdev/tldx)
- [Plik](https://github.com/root-gg/plik) **star:1795** Plik is a temporary file upload system (Wetransfer like) in Go.   [![There was an update last month][G]](https://github.com/root-gg/plik)   [![godoc][D]](https://godoc.org/github.com/root-gg/plik)
- [portal](https://github.com/SpatiumPortae/portal) **star:1756** Portal is a quick and easy command-line file transfer utility from any computer to another.   [![godoc][D]](https://godoc.org/github.com/SpatiumPortae/portal)
- [borg](https://github.com/crufter/borg) **star:1608** Terminal based search engine for bash snippets.   [![It hasn't been updated in recent three years][Y]](https://github.com/crufter/borg)   [![godoc][D]](https://godoc.org/github.com/crufter/borg)   [![Archived][Archived]](https://github.com/crufter/borg)
- [shell2http](https://github.com/msoap/shell2http) **star:1496** Executing shell commands via http server (for prototyping or remote control).   [![godoc][D]](https://godoc.org/github.com/msoap/shell2http)
- [bluetuith](https://github.com/bluetuith-org/bluetuith) **star:1362** TUI Bluetooth manager for Linux.   [![godoc][D]](https://godoc.org/github.com/bluetuith-org/bluetuith)
- [vFlow](https://github.com/VerizonDigital/vflow) **star:1156** High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.   [![godoc][D]](https://godoc.org/github.com/VerizonDigital/vflow)
- [Wave Terminal](https://waveterm.dev)  Wave is an open-source, AI-native terminal built for seamless developer workflows with inline rendering, a modern UI, and persistent sessions.
- [peg](https://github.com/pointlander/peg) **star:1113** Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.   [![godoc][D]](https://godoc.org/github.com/pointlander/peg)
- [GoNB](https://github.com/janpfeifer/gonb) **star:1032** Interactive Go programming with Jupyter Notebooks (also works in VSCode, Binder and Google's Colab).   [![godoc][D]](https://godoc.org/github.com/janpfeifer/gonb)
- [Go Package Store](https://github.com/shurcooL/Go-Package-Store) **star:897** App that displays updates for the Go packages in your GOPATH.   [![godoc][D]](https://godoc.org/github.com/shurcooL/Go-Package-Store)
- [yai](https://github.com/ekkinox/yai) **star:867** AI powered terminal assistant.   [![godoc][D]](https://godoc.org/github.com/ekkinox/yai)
- [zs](https://git.mills.io/prologic/zs)  an extremely minimal static site generator.
- [gfile](https://github.com/Antonito/gfile) **star:760** Securely transfer files between two computers, without any third party, over WebRTC.   [![There was an update last month][G]](https://github.com/Antonito/gfile)   [![godoc][D]](https://godoc.org/github.com/Antonito/gfile)
- [Leaps](https://github.com/jeffail/leaps) **star:756** Pair programming service using Operational Transforms.   [![It hasn't been updated in recent three years][Y]](https://github.com/jeffail/leaps)   [![godoc][D]](https://godoc.org/github.com/jeffail/leaps)
- [sake](https://github.com/alajmo/sake) **star:748** sake is a command runner for local and remote hosts.   [![godoc][D]](https://godoc.org/github.com/alajmo/sake)
- [Chapar](https://github.com/chapar-rest/chapar) **star:703** Chapar is a cross-platform Postman alternative built with go, aims to help developers to test their api endpoints. it support http and grpc protocols.   [![godoc][D]](https://godoc.org/github.com/chapar-rest/chapar)
- [onWatch](https://github.com/onllm-dev/onWatch) **star:676** Monitor AI API quotas across providers locally with historical tracking, alerts, and a web dashboard to avoid surprise throttling and budget overruns.   [![godoc][D]](https://godoc.org/github.com/onllm-dev/onWatch)
- [Guora](https://github.com/meloalright/guora) **star:672** A self-hosted Quora like web application written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/meloalright/guora)   [![godoc][D]](https://godoc.org/github.com/meloalright/guora)   [![Contains Chinese documents][CN]](https://github.com/meloalright/guora)
- [gocc](https://github.com/goccmack/gocc) **star:663** Gocc is a compiler kit for Go written in Go.   [![godoc][D]](https://godoc.org/github.com/goccmack/gocc)
- [Gebug](https://github.com/moshebe/gebug) **star:632** A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.   [![godoc][D]](https://godoc.org/github.com/moshebe/gebug)
- [mockingjay](https://github.com/quii/mockingjay-server) **star:568** Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.   [![It hasn't been updated in recent three years][Y]](https://github.com/quii/mockingjay-server)   [![godoc][D]](https://godoc.org/github.com/quii/mockingjay-server)
- [woke](https://github.com/get-woke/woke) **star:515** Detect non-inclusive language in your source code.   [![godoc][D]](https://godoc.org/github.com/get-woke/woke)
- [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) **star:477** Video streaming torrent client.   [![It hasn't been updated in recent three years][Y]](https://github.com/Sioro-Neoku/go-peerflix)   [![godoc][D]](https://godoc.org/github.com/Sioro-Neoku/go-peerflix)
- [goblin](https://goblin.run)  Cloud builder for CLI's written in go lang
- [mac-cleanup-go](https://github.com/2ykwang/mac-cleanup-go) **star:435** Preview-first TUI for cleaning macOS caches, logs, and temporary files.   [![There was an update last month][G]](https://github.com/2ykwang/mac-cleanup-go)   [![godoc][D]](https://godoc.org/github.com/2ykwang/mac-cleanup-go)
- [hotswap](https://github.com/edwingeng/hotswap) **star:426** A complete solution to reload your go code without restarting your server, interrupting or blocking any ongoing procedure.   [![godoc][D]](https://godoc.org/github.com/edwingeng/hotswap)   [![Contains Chinese documents][CN]](https://github.com/edwingeng/hotswap)
- [hugo](https://gohugo.io/)  Fast and Modern Static Website Engine.
- [ide](https://github.com/thestrukture/ide) **star:366** Browser accessible IDE. Designed for Go with Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/thestrukture/ide)   [![godoc][D]](https://godoc.org/github.com/thestrukture/ide)
- [stew](https://github.com/marwanhawari/stew) **star:352** An independent package manager for compiled binaries.   [![godoc][D]](https://godoc.org/github.com/marwanhawari/stew)
- [syncthing](https://syncthing.net/)  Open, decentralized file synchronization tool and protocol.
- [Neo-cowsay](https://github.com/Code-Hex/Neo-cowsay) **star:342** 🐮 cowsay is reborn. for a New Era.   [![godoc][D]](https://godoc.org/github.com/Code-Hex/Neo-cowsay)
- [crawley](https://github.com/s0rg/crawley) **star:340** Web scraper/crawler for cli.   [![There was an update last month][G]](https://github.com/s0rg/crawley)   [![godoc][D]](https://godoc.org/github.com/s0rg/crawley)
- [wellington](https://github.com/wellington/wellington) **star:302** Sass project management tool, extends the language with sprite functions (like Compass).   [![It hasn't been updated in recent three years][Y]](https://github.com/wellington/wellington)   [![godoc][D]](https://godoc.org/github.com/wellington/wellington)   [![Archived][Archived]](https://github.com/wellington/wellington)
- [Cherry](https://github.com/rafael-santiago/cherry) **star:299** Tiny webchat server in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/rafael-santiago/cherry)   [![godoc][D]](https://godoc.org/github.com/rafael-santiago/cherry)
- [fjira](https://github.com/mk-5/fjira) **star:271** A fuzzy-search based terminal UI application for Attlasian Jira   [![godoc][D]](https://godoc.org/github.com/mk-5/fjira)
- [tcpdog](https://github.com/mehrdadrad/tcpdog) **star:261** eBPF based TCP observability.   [![It hasn't been updated in recent three years][Y]](https://github.com/mehrdadrad/tcpdog)   [![godoc][D]](https://godoc.org/github.com/mehrdadrad/tcpdog)
- [joincap](https://github.com/assafmo/joincap) **star:220** Command-line utility for merging multiple pcap files together.   [![godoc][D]](https://godoc.org/github.com/assafmo/joincap)
- [Orbit](https://github.com/gulien/orbit) **star:187** A simple tool for running commands and generating files from templates.   [![It hasn't been updated in recent three years][Y]](https://github.com/gulien/orbit)   [![godoc][D]](https://godoc.org/github.com/gulien/orbit)
- [vaku](https://github.com/lingrino/vaku) **star:160** CLI & API for folder-based functions in Vault like copy, move, and search.   [![There was an update last month][G]](https://github.com/lingrino/vaku)   [![godoc][D]](https://godoc.org/github.com/lingrino/vaku)
- [GooseForum](https://github.com/leancodebox/GooseForum) **star:114** Self-hosted forum platform built with Go, Vue, and Tailwind CSS.   [![There was an update last month][G]](https://github.com/leancodebox/GooseForum)   [![godoc][D]](https://godoc.org/github.com/leancodebox/GooseForum)   [![Contains Chinese documents][CN]](https://github.com/leancodebox/GooseForum)
- [CrunchyCleaner](https://github.com/Knuspii/CrunchyCleaner) **star:107** A lightweight, software cache cleanup tool for Windows & Linux.   [![There was an update last month][G]](https://github.com/Knuspii/CrunchyCleaner)   [![godoc][D]](https://godoc.org/github.com/Knuspii/CrunchyCleaner)
- [boxed](https://github.com/tejo/boxed) **star:78** Dropbox based blog engine.   [![It hasn't been updated in recent three years][Y]](https://github.com/tejo/boxed)   [![godoc][D]](https://godoc.org/github.com/tejo/boxed)
- [dp](https://github.com/scryinfo/dp) **star:77** Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.   [![It hasn't been updated in recent three years][Y]](https://github.com/scryinfo/dp)   [![godoc][D]](https://godoc.org/github.com/scryinfo/dp)
- [Seaweed File System](https://github.com/chrislusf/seaweedfs) **star:30** Fast, Simple and Scalable Distributed File System with O(1) disk seek.   [![godoc][D]](https://godoc.org/github.com/chrislusf/seaweedfs)
- [LightCMS](https://github.com/jonradoff/lightcms) **star:23** Self-hosted content management system with static page generation, role-based access control, and an MCP server for agent-driven content operations.   [![godoc][D]](https://godoc.org/github.com/jonradoff/lightcms)
- [limetext](https://limetext.github.io)  Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
- [naclpipe](https://github.com/unix4fun/naclpipe) **star:23** Simple NaCL EC25519 based crypto pipe tool written in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/unix4fun/naclpipe)   [![godoc][D]](https://godoc.org/github.com/unix4fun/naclpipe)
- [tinycare-tui](https://github.com/DMcP89/tinycare-tui) **star:21** Small terminal app that shows git commits from the last 24 hours and week, current weather, some self care advice, a joke, and you current todo list tasks.   [![godoc][D]](https://godoc.org/github.com/DMcP89/tinycare-tui)
- [Snitch](https://github.com/lucasgomide/snitch) **star:16** Simple way to notify your team and many tools when someone has deployed any application via Tsuru.   [![It hasn't been updated in recent three years][Y]](https://github.com/lucasgomide/snitch)   [![godoc][D]](https://godoc.org/github.com/lucasgomide/snitch)
- [GoDocTooltip](https://github.com/diankong/GoDocTooltip) **star:13** Chrome extension for Go Doc sites, which shows function description as tooltip at function list.   [![It hasn't been updated in recent three years][Y]](https://github.com/diankong/GoDocTooltip)
- [hoofli](https://github.com/dnnrly/hoofli) **star:11** Generate PlantUML diagrams from Chrome or Firefox network inspections.   [![godoc][D]](https://godoc.org/github.com/dnnrly/hoofli)
- [mdv](https://github.com/Allra-Fintech/mdv) **star:2** CLI tool that renders Markdown files in the browser with live reload, GFM, syntax highlighting, Mermaid diagrams, and PDF export.   [![godoc][D]](https://godoc.org/github.com/Allra-Fintech/mdv)

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) **star:2138** Go web framework benchmark.   [![There was an update last month][G]](https://github.com/smallnest/go-web-framework-benchmark)   [![godoc][D]](https://godoc.org/github.com/smallnest/go-web-framework-benchmark)
- [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) **star:1659** Go HTTP request router benchmark and comparison.   [![godoc][D]](https://godoc.org/github.com/julienschmidt/go-http-routing-benchmark)
- [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) **star:1627** Benchmarks of Go serialization methods.   [![godoc][D]](https://godoc.org/github.com/alecthomas/go_serialization_benchmarks)
- [skynet](https://github.com/atemerev/skynet) **star:1062** Skynet 1M threads microbenchmark.
- [speedtest-resize](https://github.com/fawick/speedtest-resize) **star:244** Compare various Image resize algorithms for the Go language.   [![It hasn't been updated in recent three years][Y]](https://github.com/fawick/speedtest-resize)   [![godoc][D]](https://godoc.org/github.com/fawick/speedtest-resize)
- [go-benchmarks](https://github.com/tylertreat/go-benchmarks) **star:150** Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.   [![It hasn't been updated in recent three years][Y]](https://github.com/tylertreat/go-benchmarks)   [![godoc][D]](https://godoc.org/github.com/tylertreat/go-benchmarks)
- [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) **star:145** a collection of golang benchmarks.   [![godoc][D]](https://godoc.org/github.com/SimonWaldherr/golang-benchmarks)
- [gospeed](https://github.com/feyeleanor/GoSpeed) **star:128** Go micro-benchmarks for calculating the speed of language constructs.   [![godoc][D]](https://godoc.org/github.com/feyeleanor/GoSpeed)
- [autobench](https://github.com/davecheney/autobench) **star:100** Framework to compare the performance between different Go versions.   [![It hasn't been updated in recent three years][Y]](https://github.com/davecheney/autobench)   [![godoc][D]](https://godoc.org/github.com/davecheney/autobench)
- [vizb](https://github.com/goptics/vizb) **star:77** A CLI tool to visualize Go benchmark data in 4D.   [![There was an update last month][G]](https://github.com/goptics/vizb)   [![godoc][D]](https://godoc.org/github.com/goptics/vizb)
- [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) **star:63** Benchmarks of common basic operations for the Go language.   [![It hasn't been updated in recent three years][Y]](https://github.com/PuerkitoBio/gocostmodel)   [![godoc][D]](https://godoc.org/github.com/PuerkitoBio/gocostmodel)
- [go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks) **star:34** benchmarks for machine learning inference in Go.   [![godoc][D]](https://godoc.org/github.com/nikolaydubina/go-ml-benchmarks)
- [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) **star:27** Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.   [![It hasn't been updated in recent three years][Y]](https://github.com/mrLSD/go-benchmark-app)   [![godoc][D]](https://godoc.org/github.com/mrLSD/go-benchmark-app)
- [kvbench](https://github.com/jimrobinson/kvbench) **star:27** Key/Value database benchmark.   [![It hasn't been updated in recent three years][Y]](https://github.com/jimrobinson/kvbench)   [![godoc][D]](https://godoc.org/github.com/jimrobinson/kvbench)
- [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) **star:11** Go JSON benchmark.   [![It hasn't been updated in recent three years][Y]](https://github.com/zerosnake0/go-json-benchmark)   [![godoc][D]](https://godoc.org/github.com/zerosnake0/go-json-benchmark)

**[⬆ back to top](#contents)**

## Conferences

- [GoCon](https://gocon.connpass.com/)  Tokyo, Japan.
- [GoDays](https://www.godays.io/)  Berlin, Germany.
- [GoLab](https://golab.io/)  Florence, Italy.
- [GopherCon](https://www.gophercon.com/)  Varied Locations Each Year, USA.
- [GopherCon Africa](https://gophercon.africa/)  Nairobi, Kenya.
- [GopherCon Australia](https://gophercon.com.au/)  Sydney, Australia.
- [GopherCon Brazil](https://gopherconbr.org)  Florianópolis, Brazil.
- [GopherCon China](https://gophercon.com.cn)  Shanghai, China.
- [GopherCon Europe](https://gophercon.eu/)  Berlin, Germany.
- [GopherCon India](https://gopherconindia.org/)  Pune, India.
- [GopherCon Israel](https://www.gophercon.org.il/)  Tel Aviv, Israel.
- [GopherCon Russia](https://www.gophercon-russia.ru)  Moscow, Russia.
- [GopherCon Singapore](https://gophercon.sg)  Mapletree Business City, Singapore.
- [GopherCon UK](https://www.gophercon.co.uk/)  London, UK.
- [GopherCon Vietnam](https://gophercon.vn/)  Ho Chi Minh City, Vietnam.
- [GoWest Conference](https://www.gowestconf.com/)  Lehi, USA.

**[⬆ back to top](#contents)**

## E-Books

### E-books for purchase

- [100 Go Mistakes: How to Avoid Them](https://www.manning.com/books/100-go-mistakes-how-to-avoid-them)
- [Black Hat Go](https://nostarch.com/blackhatgo)  Go programming for hackers and pentesters.
- [Build an Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)
- [Continuous Delivery in Go](https://www.manning.com/books/continuous-delivery-in-go)  This practical guide to continuous delivery shows you how to rapidly establish an automated pipeline that will improve your testing, code quality, and final product.
- [Creative DIY Microcontroller Project With TinyGo and WebAssembly](https://www.packtpub.com/product/creative-diy-microcontroller-projects-with-tinygo-and-webassembly/9781800560208)  An introduction into the TinyGo compiler with projects involving Arduino and WebAssembly.
- [Effective Go: Elegant, efficient, and testable code](https://www.manning.com/books/effective-go)  Unlock Go’s unique perspective on program design, and start writing simple, maintainable, and testable Go code.
- [For the Love of Go](https://bitfieldconsulting.com/books/love)  An introductory book for Go beginners.
- [Go in Practice, Second Edition](https://www.manning.com/books/go-in-practice-second-edition)  Your practical guide on the ins-and-outs of Go development, covering the standard library and the most important tools from Go’s powerful ecosystem.
- [Know Go: Generics](https://bitfieldconsulting.com/books/generics)  A guide to understanding and using generics in Go.
- [Lets-Go](https://lets-go.alexedwards.net)  A step-by-step guide to creating fast, secure and maintanable web applications with Go.
- [Lets-Go-Further](https://lets-go-further.alexedwards.net)  Advanced patterns for building APIs and web applications in Go.
- [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests)  A guide to testing in Go.
- [The Power of Go: Tools](https://bitfieldconsulting.com/books/tools)  A guide to writing command-line tools in Go.
- [Writing A Compiler In Go](https://compilerbook.com)
- [Writing An Interpreter In Go](https://interpreterbook.com)  Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.

### Free e-books

- [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
- [An Introduction to Programming in Go](http://www.golang-book.com/)
- [Build a blockchain from scratch in Go with gRPC](https://github.com/volodymyrprokopyuk/go-blockchain) **star:550** The foundational and practical guide for effectively learning and progressively building a blockchain from scratch in Go with gRPC.   [![godoc][D]](https://godoc.org/github.com/volodymyrprokopyuk/go-blockchain)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)
- [Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [GoBooks](https://github.com/dariubs/GoBooks) **star:19553** A curated list of Go books.   [![There was an update last month][G]](https://github.com/dariubs/GoBooks)   [![godoc][D]](https://godoc.org/github.com/dariubs/GoBooks)
- [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook)  A 600 page introduction to Go aimed at first time developers.
- [Go AST Book (Chinese)](https://github.com/chai2010/go-ast-book) **star:5510** A book focusing on Go `go/*` packages.   [![godoc][D]](https://godoc.org/github.com/chai2010/go-ast-book)
- [Go Faster](https://leanpub.com/gofaster)  This book seeks to shorten your learning curve and help you become a proficient Go programmer, faster.
- [Go 101](https://go101.org)  A book focusing on Go syntax/semantics and all kinds of details.
- [Go Succinctly](https://github.com/thedevsir/gosuccinctly) **star:23** in Persian.   [![It hasn't been updated in recent three years][Y]](https://github.com/thedevsir/gosuccinctly)   [![godoc][D]](https://godoc.org/github.com/thedevsir/gosuccinctly)
- [Go with the domain](https://threedots.tech/go-with-the-domain/)  A book showing how to apply DDD, Clean Architecture, and CQRS by practical refactoring.
- [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
- [Network Programming With Go](https://jan.newmarch.name/golang/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
- [The Go Programming Language](https://www.gopl.io/)
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
- [The Little Go Book](https://github.com/karlseguin/the-little-go-book)
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)

**[⬆ back to top](#contents)**

## Gophers

- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) **star:3964** Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.   [![godoc][D]](https://godoc.org/github.com/MariaLetta/free-gophers-pack)
- [gopher-logos](https://github.com/GolangUA/gopher-logos) **star:141** adorable gopher logos.   [![It hasn't been updated in recent three years][Y]](https://github.com/GolangUA/gopher-logos)
- [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) **star:75** Go gopher Vector Data [.ai, .svg].   [![It hasn't been updated in recent three years][Y]](https://github.com/keygx/Go-gopher-Vector)
- [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
- [gophericons](https://github.com/shalakhin/gophericons)
- [gophers](https://github.com/egonelbre/gophers) **star:3820** Free gophers.   [![It hasn't been updated in recent three years][Y]](https://github.com/egonelbre/gophers)   [![godoc][D]](https://godoc.org/github.com/egonelbre/gophers)
- [gophers](https://github.com/ashleymcnamara/gophers) **star:3080** Gopher artworks by Ashley McNamara.   [![godoc][D]](https://godoc.org/github.com/ashleymcnamara/gophers)
- [gopherize.me](https://github.com/matryer/gopherize.me) **star:757** Gopherize yourself.   [![It hasn't been updated in recent three years][Y]](https://github.com/matryer/gopherize.me)
- [gophers](https://github.com/sillecelik/go-gopher) **star:163** Gopher amigurumi toy pattern.
- [gophers](https://github.com/rogeralsing/gophers) **star:58** random gopher graphics.   [![It hasn't been updated in recent three years][Y]](https://github.com/rogeralsing/gophers)
- [gophers](https://github.com/scraly/gophers) **star:37** Gophers by Aurélie Vache.

**[⬆ back to top](#contents)**

## Meetups

- [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
- [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
- [Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)
- [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
- [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
- [Bärner Go Meetup - Berne, Switzerland](https://www.meetup.com/berner-go-meetup/)
- [Go Ireland - Dublin](https://www.meetup.com/goireland/)
- [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
- [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
- [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
- [Go Toronto](https://www.meetup.com/go-toronto/)
- [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
- [GoBandung](https://www.meetup.com/GoBandung/)
- [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
- [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
- [GoJakarta](https://www.meetup.com/GoJakarta/)
- [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
- [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
- [Golang Athens](https://www.meetup.com/Athens-Gophers/)
- [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
- [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
- [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
- [Golang Boston](https://www.meetup.com/bostongo/)
- [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
- [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
- [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
- [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
- [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
- [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
- [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
- [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
- [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
- [Golang Israel](https://www.meetup.com/Go-Israel/)
- [Golang Kathmandu](https://www.meetup.com/Golang-Kathmandu/)
- [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
- [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
- [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
- [Golang Melbourne](https://www.meetup.com/golang-mel/)
- [Golang Milano](https://www.meetup.com/golang-milano/)
- [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
- [Golang Paris](https://www.meetup.com/Golang-Paris/)
- [Golang Poland](https://www.meetup.com/Golang-Poland/)
- [Golang Pune](https://www.meetup.com/Golang-Pune/)
- [Golang Roma](https://www.meetup.com/golangroma/)
- [Golang Rotterdam](https://www.meetup.com/golang-rotterdam/)
- [Golang Singapore](https://www.meetup.com/golangsg/)
- [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
- [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
- [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
- [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
- [Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)
- [Golang Torino](https://www.meetup.com/golang-torino/)
- [Golang Turkey](https://kommunity.com/goturkiye)
- [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
- [Golang Vienna, Austria](https://www.meetup.com/viennago/)
- [Golang Москва](https://www.meetup.com/Golang-Moscow/)
- [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
- [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
- [Lagos Gophers](https://www.meetup.com/GolangNigeria/)
- [Nairobi Gophers](https://www.meetup.com/nairobi-gophers/)
- [Seattle Go Programmers](https://www.meetup.com/golang/)
- [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
- [Utah Go User Group](https://www.meetup.com/utahgophers/)
- [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)
- [Zürich Gophers - Zurich, Switzerland](https://www.meetup.com/zurich-gophers/)

_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md)
- [enra/go-styleguide](https://codeberg.org/enra/go-styleguide)
- [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
- [Google](https://google.github.io/styleguide/go/)
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst)
- [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)
- [Trybe](https://github.com/betrybe/playbook-go/blob/main/README_EN.md)
- [Uber](https://github.com/uber-go/guide/blob/master/style.md)

**[⬆ back to top](#contents)**

## Social Media

### Twitter

- [@GoDiscussions](https://twitter.com/GoDiscussions)
- [@golang](https://twitter.com/golang)
- [@golang_news](https://twitter.com/golang_news)
- [@golangch](https://twitter.com/golangch)
- [@golangweekly](https://twitter.com/golangweekly)

**[⬆ back to top](#contents)**

### Reddit

- [r/golang](https://www.reddit.com/r/golang/)

**[⬆ back to top](#contents)**

## Websites

- [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) **star:46710** Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
- [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) **star:33541** List of other amazingly awesome lists.
- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) **star:524** A curated list of awesome golang workshops.   [![It hasn't been updated in recent three years][Y]](https://github.com/amit-davidson/awesome-golang-workshops)
- [gocryforhelp](https://github.com/ninedraft/gocryforhelp) **star:39** Collection of Go projects that needs help. Good place to start your open-source way in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/ninedraft/gocryforhelp)
- [Golang Developer Jobs](https://golangjob.xyz)  Developer Jobs exclusively for Golang related Roles.
- [Golang News](https://golangnews.com)  Links and news about Go programming.
- [Golang Nugget](https://golangnugget.com)  A weekly roundup of the best Go content, delivered to your inbox every Monday.
- [Golang Weekly](https://discu.eu/weekly/golang/)  Each monday projects, tutorials and articles about Go.
- [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)  Go mailing list.
- [Gopher Community Chat](https://invite.slack.golangbridge.org)  Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
- [Gophercises](https://gophercises.com/)  Free coding exercises for budding gophers.
- [json2go](https://m-zajac.github.io/json2go)  Advanced JSON to Go struct conversion - online tool.
- [justforfunc](https://www.youtube.com/c/justforfunc)  Youtube channel dedicated to Go programming language tips and tricks, hosted by Francesc Campoy [@francesc](https://twitter.com/francesc).
- [Learn Go Programming](https://blog.learngoprogramming.com)  Learn Go concepts with illustrations.
- [Awesome Go @LibHunt](https://go.libhunt.com)  Your go-to Go Toolbox.
- [awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra) **star:27** Parse awesome-go README file and generate a new README file with repo info.   [![It hasn't been updated in recent three years][Y]](https://github.com/xwjdsh/awesome-go-extra)   [![godoc][D]](https://godoc.org/github.com/xwjdsh/awesome-go-extra)   [![Contains Chinese documents][CN]](https://github.com/xwjdsh/awesome-go-extra)
- [Code with Mukesh](https://codewithmukesh.com/categories/golang)  Software Engineer and Blogs @ codewithmukesh.com.
- [Coding Mystery](https://codingmystery.com)  Solve exciting escape-room-inspired programming challenges using Go.
- [CodinGame](https://www.codingame.com/)  Learn Go by solving interactive tasks using small games as practical examples.
- [Go Blog](https://blog.golang.org)  The official Go blog.
- [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3)  A group of Gophers read and discuss a different Go project every week.
- [Go Community on Hashnode](https://hashnode.com/n/go)  Community of Gophers on Hashnode.
- [Go Forum](https://forum.golangbridge.org)  Forum to discuss Go.
- [Go Projects](https://github.com/golang/go/wiki/Projects)  List of projects on the Go community wiki.
- [Go Proverbs](https://go-proverbs.github.io/)  Go Proverbs by Rob Pike.
- [Go Report Card](https://goreportcard.com)  A report card for your Go package.
- [go.dev](https://go.dev/)  A hub for Go developers.
- [Libs.tech](https://libs.tech/go) – Awesome Go libraries and hidden gems
- [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
- [pkg.go.dev](https://pkg.go.dev/)  Documentation for open source Go packages.
- [studygolang](https://studygolang.com)  The community of studygolang in China.
- [Trending Go repositories on GitHub today](https://github.com/trending?l=go)  Good place to find new Go libraries.
- [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

**[⬆ back to top](#contents)**

### Tutorials

- [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) **star:43925** Golang ebook intro how to build a web app with golang.   [![godoc][D]](https://godoc.org/github.com/astaxie/build-web-application-with-golang)   [![Contains Chinese documents][CN]](https://github.com/astaxie/build-web-application-with-golang)
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)  We’ll write an API with the help of the powerful Gorilla Mux.
- [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)  Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
- [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)  How to cache slow database queries.
- [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)  How to cancel MySQL queries.
- [CodeCrafters Golang Track](https://app.codecrafters.io/tracks/go)  Achieve mastery in advanced Go by building your own Redis, Docker, Git, and SQLite. Featuring goroutines, systems programming, file I/O, and more.
- [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) **star:8807** Go's reference card.
- [Go database/sql tutorial](http://go-database-sql.org/)  Introduction to database/sql.
- [Go in 7 days](https://github.com/harrytran103/7_days_of_go) **star:164** Learn everything about Go in 7 days (from a Nodejs developer).   [![godoc][D]](https://godoc.org/github.com/harrytran103/7_days_of_go)
- [Go Language Tutorial](https://www.javatpoint.com/go-tutorial)  Learn Go language Tutorial.
- [Go Tutorial](https://www.tutorialspoint.com/go/index.htm)  Learn Go programming.
- [Design Patterns in Go](https://github.com/shubhamzanwar/design-patterns) **star:132** Collection of programming design patterns implemented in Go.   [![It hasn't been updated in recent three years][Y]](https://github.com/shubhamzanwar/design-patterns)   [![godoc][D]](https://godoc.org/github.com/shubhamzanwar/design-patterns)
- [Games With Go](https://www.youtube.com/watch?v=9D4yH7e_ea8&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x)  A video series teaching programming and game development.
- [Go By Example](https://gobyexample.com/)  Hands-on introduction to Go using annotated example programs.
- [50 Shades of Go](https://golang50shades.github.io/)  Traps, Gotchas, and Common Mistakes for New Golang Devs.
- [A Comprehensive Guide to Structured Logging in Go](https://betterstack.com/community/guides/logging/logging-in-go/)  Delve deep into the world of structured logging in Go with a specific focus on recently accepted slog proposal which aims to bring high performance structured logging with levels to the standard library.
- [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)  Building a Golang site for e-commerce (demo included).
- [A Tour of Go](https://tour.golang.org/)  Interactive tour of Go.
- [Build a Database in 1000 lines of code](https://link.medium.com/O9YQlx89Htb)  Build a NoSQL Database From Zero in 1000 Lines of Code.
- [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [go-patterns](https://github.com/tmrts/go-patterns) **star:28131** Curated list of Go design patterns, recipes and idioms.   [![godoc][D]](https://godoc.org/github.com/tmrts/go-patterns)
- [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) **star:23760** Learn Go with test-driven development.   [![godoc][D]](https://godoc.org/github.com/quii/learn-go-with-tests)   [![Contains Chinese documents][CN]](https://github.com/quii/learn-go-with-tests)
- [Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n)  Series of articles in order to learn Golang language by concrete applications as example.
- [Microservices with Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)  Dive deep into building microservices using Go, including gRPC.
- [package main](https://www.youtube.com/packagemain)  YouTube channel about Programming in Go.
- [Programming with Google Go](https://www.coursera.org/specializations/google-golang)  Coursera Specialization to learn about Go from scratch.
- [Scaling Go Applications](https://betterstack.com/community/guides/scaling-go/)  Everything about building, deploying and scaling Go applications in production.
- [Learn Go with 1000+ Exercises](https://github.com/inancgumus/learngo) **star:20079** Learn Go with thousands of examples, exercises, and quizzes.   [![godoc][D]](https://godoc.org/github.com/inancgumus/learngo)
- [go-clean-template](https://github.com/evrone/go-clean-template) **star:7625** Clean Architecture template for Golang services.   [![There was an update last month][G]](https://github.com/evrone/go-clean-template)   [![godoc][D]](https://godoc.org/github.com/evrone/go-clean-template)
- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) **star:4769** Examples of Golang compared to Node.js for learning.   [![It hasn't been updated in recent three years][Y]](https://github.com/miguelmota/golang-for-nodejs-developers)   [![godoc][D]](https://godoc.org/github.com/miguelmota/golang-for-nodejs-developers)   [![Archived][Archived]](https://github.com/miguelmota/golang-for-nodejs-developers)
- [Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/)  A List of Free Courses to Learn the Go Programming Language.
- [golang-examples](https://github.com/SimonWaldherr/golang-examples) **star:1698** Many examples to learn Golang.   [![godoc][D]](https://godoc.org/github.com/SimonWaldherr/golang-examples)
- [Golangbot](https://golangbot.com/learn-golang-series/)  Tutorials to get started with programming in Go.
- [GopherCoding](https://gophercoding.com/)  Collection of code snippets and tutorials to help tackle every day issues.
- [GopherSnippets](https://gophersnippets.com/)  Code snippets with tests and testable examples for the Go programming language.
- [Gosamples](https://gosamples.dev/)  Collection of code snippets that let you solve everyday code problems.
- [GraphQL with Go](https://hasura.io/learn/graphql/backend-stack/languages/go/)  Learn how to create a Go GraphQL server and client with code generation. Also includes creating REST endpoints.
- [Hackr.io](https://hackr.io/tutorials/learn-golang)  Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
- [Hex Monscape](https://github.com/Haraj-backend/hex-monscape) **star:84** Getting started guidelines in writing maintainable code using Hexagonal Architecture.   [![godoc][D]](https://godoc.org/github.com/Haraj-backend/hex-monscape)
- [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5)  Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
- [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker)  Learn how to use Docker for Go development and how to build production Docker images.
- [How to Implement Role-Based Access Control (RBAC) Authorization in Golang](https://www.permit.io/blog/role-based-access-control-rbac-authorization-in-golang)  A guide to implementing Role-Based Access Control (RBAC) in Golang, including code examples, covering various methods to secure app endpoints with role-based authorization.
- [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)  Get started with Godog - a Behavior-driven development framework for building and testing Go applications.
- [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
- [Understanding Go in a visual way](https://dev.to/aurelievache/series/26234)  Learn Go visually
- [W3basic Go Tutorials](https://www.w3basic.com/golang/)  W3Basic provides an in-depth tutorial and well-organized content to learn Golang programming.
- [Your basic Go](https://yourbasic.org/golang)  Huge collection of tutorials and how to's.

**[⬆ back to top](#contents)**

### Guided Learning

- [The Go Developer Roadmap](https://roadmap.sh/golang)  A visual roadmap that new Go developers can follow through to help them learn Go.
- [The Go Interview Practice](https://github.com/RezaSi/go-interview-practice) **star:2412** A GitHub repository offering coding challenges for Go technical interview preparation.   [![There was an update last month][G]](https://github.com/RezaSi/go-interview-practice)   [![godoc][D]](https://godoc.org/github.com/RezaSi/go-interview-practice)
- [The Go Learning Path](https://tutorialedge.net/paths/golang/)  A guided learning path containing a mix of free and premium resources.
- [The Go Skill Tree](https://labex.io/skilltrees/go)  A structured learning path that combines both free and premium resources.

**[⬆ back to top](#contents)**

## Contribution

We welcome contributions! Please refer to our [CONTRIBUTING.md](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the [MIT License](https://github.com/avelino/awesome-go/blob/main/LICENSE) - see the LICENSE file for details.

