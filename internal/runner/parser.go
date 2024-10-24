package runner

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// Options will defines its options
type Options struct {
	Concurrency int
	Depth       int
	Ext         string
	File        *os.File
	List        *bufio.Scanner
	Output      string
	SameHost    bool
	SameRoot    bool
	Silent      bool
	Template    string
	Timeout     int
	URL         string
	Verbose     bool
	Wait        int
}

// Parse user given arguments
func Parse() *Options {
	opt = &Options{}

	flag.StringVar(&opt.URL, "url", "", "")
	flag.StringVar(&opt.URL, "u", "", "")

	flag.IntVar(&opt.Concurrency, "concurrency", 50, "")
	flag.IntVar(&opt.Concurrency, "c", 50, "")

	flag.IntVar(&opt.Wait, "wait", 1, "")
	flag.IntVar(&opt.Wait, "w", 1, "")

	flag.IntVar(&opt.Depth, "depth", 1, "")
	flag.IntVar(&opt.Depth, "d", 1, "")

	flag.IntVar(&opt.Timeout, "timeout", 60, "")
	flag.IntVar(&opt.Timeout, "t", 60, "")

	flag.StringVar(&opt.Ext, "e", "", "")
	flag.StringVar(&opt.Ext, "extension", "", "")

	flag.BoolVar(&opt.SameHost, "same-host", false, "")
	flag.BoolVar(&opt.SameRoot, "same-root", false, "")

	flag.StringVar(&opt.Output, "output", "", "")
	flag.StringVar(&opt.Output, "o", "", "")

	flag.StringVar(&opt.Template, "template", "", "")
	flag.StringVar(&opt.Template, "T", "", "")

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
		clog.Fatal("could not validate options", "err", err)
	}

	return opt
}

func showBanner() {
	fmt.Fprint(os.Stderr, aurora.Bold(aurora.Cyan(banner)))
}
