package runner

import (
	"fmt"
	"github.com/dwisiswant0/galer/pkg/galer"
	"github.com/projectdiscovery/gologger"
	"strings"
)

// Run to executes galer
func Run(opt *Options) {
	opt.ListURLS = []string{opt.URL}
	// check if load from text file
	if opt.Path != "" {
		opt.ListURLS = galer.LoadUrlsFromFile(opt.Path)
	}
	var ch = make(chan bool)
	var t = make([]interface{}, len(opt.ListURLS))
	cfg = &galer.Config{
		Timeout: opt.Timeout,
	}

	// fetch all url in list urls
	for i, url := range opt.ListURLS {
		cfg = galer.New(cfg)
		go opt.fetch(url, cfg, &t[i], ch)
	}

	// collect data
	for i := 0; i < len(opt.ListURLS); i++ {
		<-ch
	}

	// result
	var res = opt.mapValue(t)

	// close channel
	close(ch)
	cfg.Cancel()

	// if wants to collect result in text
	if opt.File != nil {
		fmt.Fprintf(opt.File, "%s\n", strings.Join(res, "\n"))
		opt.File.Close()
	} else {
		fmt.Println(strings.Join(res, "\n"))
	}
}

func (opt *Options) mapValue(results []interface{}) []string {
	var t []string
	for i, v := range results {
		if c, ok := v.([]string); ok {
			for _, j := range c {
				if opt.Ext != "" && !opt.isOnExt(j) {
					continue
				}
				if opt.InScope && !isScope(opt.ListURLS[i], j) {
					continue
				}
				t = append(t, j)
			}
		}
	}
	return t
}

func (opt *Options) fetch(URL string, cfg *galer.Config, result *interface{}, ch chan<- bool) {
	var resp []string
	res, err := cfg.Crawl(URL)
	if err != nil && !opt.Silent {
		msg := "cannot fetch URL"
		if opt.Verbose {
			msg = err.Error()
		}
		gologger.Errorf("Error '%s': %s.", URL, msg)
		*result = resp
	}
	*result = res
	ch <- true
}
