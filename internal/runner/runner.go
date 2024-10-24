package runner

import (
	"fmt"
	"io"
	"os"

	"github.com/dwisiswant0/galer/pkg/galer"
	"github.com/remeh/sizedwaitgroup"
)

type Runner struct {
	opt   *Options
	swg   sizedwaitgroup.SizedWaitGroup
	urls  map[string]bool
	galer *galer.Config
}

// New initialize [Runner]
func New(opt *Options) *Runner {
	return &Runner{
		opt:  opt,
		swg:  sizedwaitgroup.New(opt.Concurrency),
		urls: make(map[string]bool),
		galer: &galer.Config{
			Logger:   clog,
			SameHost: opt.SameHost,
			SameRoot: opt.SameRoot,
			Template: opt.Template,
			Timeout:  opt.Timeout,
			Wait:     opt.Wait,
		},
	}
}

// Do runs crawling
func (r *Runner) Do() {
	jobs := make(chan string)

	for i := 0; i < r.opt.Concurrency; i++ {
		r.swg.Add()
		go func() {
			defer r.swg.Done()
			for job := range jobs {
				r.galer.SetScope(job)
				r.run(job, 1)
			}
		}()
	}

	for r.opt.List.Scan() {
		u := r.opt.List.Text()
		jobs <- u
	}

	close(jobs)
	r.swg.Wait()
	r.galer.Close()

	if r.opt.File != nil {
		r.opt.File.Close()
	}
}

func (r *Runner) run(URL string, counter int) {
	cfg := galer.New(r.galer)

	var writer io.Writer = os.Stdout
	if r.opt.File != nil {
		writer = io.MultiWriter(os.Stdout, r.opt.File)
	}

	for counter <= r.opt.Depth {
		crawl := r.crawl(URL, cfg)
		if len(crawl) == 0 {
			break
		}
		counter++

		var batches []string
		for _, u := range crawl {
			if !r.urls[u] {
				fmt.Fprintf(writer, "%s\n", u)
				batches = append(batches, u)
				r.urls[u] = true
			}
		}

		for _, u := range batches {
			if r.opt.Ext != "" {
				if !r.opt.isOnExt(u) {
					continue
				}
			}

			if counter <= r.opt.Depth {
				r.run(u, counter+1)
			}
		}
	}
}

func (r *Runner) crawl(URL string, cfg *galer.Config) []string {
	res, err := cfg.Crawl(URL)
	if err != nil && opt.Verbose {
		clog.Error(err, "url", URL)

		return []string{}
	}

	return res
}
