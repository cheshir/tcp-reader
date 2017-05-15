# About

This app creates server that listen on the specified port and prints all incoming messages to the stdout.

All internal messages is prefixed by `>>>`.

Messages delimiter: `\n`.

# Install

`go get github.com/cheshir/tcp-reader`

# Usage

With default settings:

```bash
tcp-reader
```

Set port:

```bash
tcp-reader -port=9999
```

Filter messages by regexp (supports [golang regexp syntax](https://golang.org/pkg/regexp/syntax/)):

```bash
tcp-reader -filter=mask
```

Matched substrings would be highlighted if TERM environment variable is set.
