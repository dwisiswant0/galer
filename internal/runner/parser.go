package runner

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
)

// Options will defines its options
type Options struct {
	Concurrency int
	Timeout     int
	Depth       int
	URL         string
	Ext         string
	Output      string
	InScope     bool
	Silent      bool
	Verbose     bool
	List        *bufio.Scanner
	File        *os.File
}

// Parse user given arguments
func Parse() *Options {
	opt = &Options{}

	flag.StringVar(&opt.URL, "url", "", "")
	flag.StringVar(&opt.URL, "u", "", "")

	flag.IntVar(&opt.Concurrency, "concurrency", 50, "")
	flag.IntVar(&opt.Concurrency, "c", 50, "")

	flag.IntVar(&opt.Timeout, "timeout", 60, "")
	flag.IntVar(&opt.Timeout, "t", 60, "")

	flag.StringVar(&opt.Ext, "e", "", "")
	flag.StringVar(&opt.Ext, "extension", "", "")

	flag.BoolVar(&opt.InScope, "in-scope", false, "")

	flag.StringVar(&opt.Output, "output", "", "")
	flag.StringVar(&opt.Output, "o", "", "")

	flag.BoolVar(&opt.Silent, "silent", false, "")
	flag.BoolVar(&opt.Silent, "s", false, "")

	flag.BoolVar(&opt.Verbose, "v", false, "")
	flag.BoolVar(&opt.Verbose, "verbose", false, "")

	flag.Usage = func() {
		showBanner()
		fmt.Fprint(os.Stderr, help)
	}

	flag.Parse()

	if !opt.Silent {
		showBanner()
	}

	if err := opt.validate(); err != nil {
		gologger.Fatalf("Error! %s", err.Error())
	}

	return opt
}

func showBanner() {
	fmt.Fprint(os.Stderr, aurora.Bold(aurora.Cyan(banner)))
}
