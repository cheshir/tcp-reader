[![Build Status](https://travis-ci.com/cheshir/tcp-reader.svg?branch=master)](https://travis-ci.com/cheshir/tcp-reader)
[![Go Report Card](https://goreportcard.com/badge/github.com/cheshir/tcp-reader)](https://goreportcard.com/report/github.com/cheshir/tcp-reader)

# About

This app creates server that listen on the specified port and prints all incoming messages to the stdout.

Default configuration: 

```bash
tcp-reader -host=localhost -port=9999 -prefix=">>> "
```

All internal messages is prefixed by `>>>`.

Messages delimiter: `\n`.

# Install

For everyone:

`curl -sf https://gobinaries.com/cheshir/tcp-reader | sh`

For go users:

`go get -u github.com/cheshir/tcp-reader`

# Usage

With default settings:

```bash
tcp-reader
```

For example we're going to read error and panic messages only:

```bash
tcp-reader -host=localhost -port=9999 -prefix=">>> " -filter="(error|panic)"
```

Filter supports [golang regexp syntax](https://golang.org/pkg/regexp/syntax/).

Matched substrings would be highlighted if TERM environment variable is set.
