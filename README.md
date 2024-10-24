# galer

[![made-with-Go](https://img.shields.io/badge/made%20with-Go-blue.svg)](http://golang.org)
[![issues](https://img.shields.io/github/issues/dwisiswant0/galer?color=blue)](https://github.com/dwisiswant0/galer/issues)

```txt
             __
   __    _ _(_ )   __  _ __ 
 /'_ '\/'_' )| | /'__'( '__)
( (_) ( (_| || |(  ___| |
'\__  '\__,_(___'\____(_)
( )_) |
 \___/'  @dwisiswant0
```

A fast tool to fetch URLs from HTML attributes by crawl-in. Inspired by the [@omespino Tweet](https://twitter.com/omespino/status/1318605084989837312), which is possible to extract `src`, `href`, `url` and `action` values by evaluating JavaScript through Chrome DevTools Protocol.

---

## Resources

- [Installation](#installation)
	- [from Binary](#from-binary)
	- [from Source](#from-source)
	- [from GitHub](#from-github)
- [Usage](#usage)
	- [Basic Usage](#basic-usage)
	- [Flags](#flags)
	- [Examples](#examples)
		- [Single URL](#single-url)
		- [URLs from list](#urls-from-list)
		- [from Stdin](#from-stdin)
	- [Library](#library)
- [TODOs](#todos)
- [Help & Bugs](#help--bugs)
- [License](#license)
- [Version](#version)
- [Acknowledgement](#acknowledgement)

## Installation

### from Binary

The installation is easy. You can download a prebuilt binary from [releases page](https://github.com/dwisiswant0/galer/releases), unpack and run! or with

```bash
▶ (sudo) curl -sSfL https://git.io/galer | sh -s -- -b /usr/local/bin
```

### from Source

If you have go1.22+ compiler installed and configured:

```bash
▶ go install -v github.com/dwisiswant0/galer@latest
```

### from GitHub

```bash
▶ git clone https://github.com/dwisiswant0/galer
▶ cd galer
▶ go build .
▶ (sudo) install galer /usr/local/bin
```

## Usage

### Basic Usage

Simply, galer can be run with:

```bash
▶ galer -u "http://domain.tld"
```

### Flags

![galer](https://user-images.githubusercontent.com/25837540/100824601-0ee53b80-3489-11eb-878d-a58d1ec3489d.jpg)

This will display help for the tool. Here are all the options it supports.

```console
$ galer -h

             __   v0.2.0
   __    _ _(_ )   __  _ __
 /'_ '\/'_' )| | /'__'( '__)
( (_) ( (_| || |(  ___| |
'\__  '\__,_(___'\____(_)
( )_) |
 \___/'  @dwisiswant0

A fast tool to fetch URLs from HTML attributes by crawl-in

Usage:
  galer -u [URL|URLs.txt] -o [output.txt]

Options:
  -u, --url <URL/FILE>        Target to fetches (single target URL or list)
  -e, --extension <EXT>       Show only certain extensions (comma-separated, e.g. js,php)
  -c, --concurrency <N>       Concurrency level (default: 50)
  -w, --wait <N>              Wait N seconds before evaluate (default: 1)
  -d, --depth <N>             Max. depth for crawling (levels of links to follow)
      --same-host             Same host only
      --same-root             Same root (eTLD+1) only (takes precedence over --same-host)
  -o, --output <FILE>         Save fetched URLs output into file
  -T, --template <string>     Format for output template (e.g., "{{scheme}}://{{host}}{{path}}")
                              Valid variables are: "raw_url", "scheme", "user", "username",
                              "password", "host", "hostname", "port", "path", "raw_path",
                              "escaped_path", "raw_query", "fragment", "raw_fragment".
  -t, --timeout <N>           Max. time (seconds) allowed for connection (default: 60)
  -s, --silent                Silent mode (suppress an errors)
  -v, --verbose               Verbose mode show error details unless you weren't use silent
  -h, --help                  Display its helps
```

### Examples

#### Single URL

```bash
▶ galer -u "http://domain.tld"
```

#### URLs from list

```bash
▶ galer -u /path/to/urls.txt
```

#### from Stdin

```bash
▶ cat urls.txt | galer
```

In case you want to chained with other tools:

```bash
▶ subfinder -d domain.tld -silent | httpx -silent | galer
```

### Library

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/dwisiswant0/galer/pkg/galer)

You can use **galer** as library.

```
▶ go get github.com/dwisiswant0/galer/pkg/galer@latest
```

For example:

```go
package main

import (
	"fmt"

	"github.com/dwisiswant0/galer/pkg/galer"
)

func main() {
	cfg := &galer.Config{
		Timeout: 60,
	}
	cfg = galer.New(cfg)

	run, err := cfg.Crawl("https://twitter.com")
	if err != nil {
		panic(err)
	}

	for _, url := range run {
		fmt.Println(url)
	}
}
```

## TODOs

- [ ] Enable to set extra HTTP headers
- [ ] Provide randomly User-Agent
- [ ] Bypass headless browser
- [ ] Add exception for specific extensions

## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/dwisiswant0/galer/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/dwisiswant0/galer/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.

## Status

> [!CAUTION]
> galer has NOT reached 1.0 yet. Therefore, this library is currently not supported and does not offer a stable API; use at your own risk.

There are no guarantees of stability for the APIs in this library, and while they are not expected to change dramatically. API tweaks and bug fixes may occur.

## Pronunciation

`id_ID` • **/gäˈlər/** — kalau _galer_ jangan dicium baunya, langsung cuci tangan, _bego_!

## Acknowledgement

- [Omar Espino](https://twitter.com/omespino) for the idea, that's why this tool was made!

### License

`sebel` is released by **@dwisiswant0** under the MIT license. See [LICENSE](/LICENSE).