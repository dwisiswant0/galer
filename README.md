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

If you have go1.15+ compiler installed and configured:

```bash
▶ GO111MODULE=on go get github.com/dwisiswant0/galer
```

### from GitHub

```bash
▶ git clone https://github.com/dwisiswant0/galer
▶ cd galer
▶ go build .
▶ (sudo) mv galer /usr/local/bin
```

## Usage

### Basic Usage

Simply, galer can be run with:

```bash
▶ galer -u "http://domain.tld"
```

### Flags

```bash
▶ galer -h
```

![galer](https://user-images.githubusercontent.com/25837540/100824601-0ee53b80-3489-11eb-878d-a58d1ec3489d.jpg)

This will display help for the tool. Here are all the switches it supports.

| **Flag**          	| **Description**                                                 	|
|-------------------	|-----------------------------------------------------------------	|
| -u, --url         	| Target to fetches _(single target URL or list)_                 	|
| -e, --extension   	| Show only certain extensions _(comma-separated, e.g. js,php)_   	|
| -c, --concurrency 	| Concurrency level _(default: 50)_                               	|
|     --in-scope    	| Show in-scope URLs/same host only                               	|
| -o, --output      	| Save fetched URLs output into file                              	|
| -t, --timeout     	| Maximum time _(seconds)_ allowed for connection _(default: 60)_ 	|
| -s, --silent      	| Silent mode _(suppress an errors)_                              	|
| -v, --verbose     	| Verbose mode show error details unless you weren't use silent   	|
| -h, --help        	| Display its helps                                               	|

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
▶ go get github.com/dwisiswant0/galer/pkg/galer
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

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**galer** released under MIT. See `LICENSE` for more details.

## Version

**Current version is 0.0.2** and still development.

## Pronunciation

`id_ID` • **/gäˈlər/** — kalau _galer_ jangan dicium baunya, langsung cuci tangan, _bego_!

## Acknowledgement

- [Omar Espino](https://twitter.com/omespino) for the idea, that's why this tool was made!
