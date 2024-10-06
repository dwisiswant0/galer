package runner

import (
	"fmt"

	"github.com/dwisiswant0/galer/pkg/galer"
	"github.com/remeh/sizedwaitgroup"
)

// New to executes galer
func New(opt *Options) {
	job := make(chan string)
	con := opt.Concurrency
	swg := sizedwaitgroup.New(con)
	cfg = &galer.Config{Timeout: opt.Timeout}
	cfg = galer.New(cfg)

	for i := 0; i < con; i++ {
		swg.Add()
		go func() {
			defer swg.Done()
			for URL := range job {
				run := opt.run(URL, cfg)
				for _, u := range run {
					if opt.Ext != "" {
						if !opt.isOnExt(u) {
							continue
						}
					}

					if opt.InScope {
						if isScope(URL, u) {
							out = u
						}
					} else {
						out = u
					}

					if out != "" {
						fmt.Println(out)

						if opt.File != nil {
							fmt.Fprintf(opt.File, "%s\n", out)
						}
					}
				}
			}
		}()
	}

	for opt.List.Scan() {
		u := opt.List.Text()
		job <- u
	}

	close(job)
	swg.Wait()
	_ = cfg.Close()

	if opt.File != nil {
		opt.File.Close()
	}
}

func (opt *Options) run(URL string, cfg *galer.Config) []string {
	res, err := cfg.Crawl(URL)
	if err != nil && !opt.Silent {
		clog.Error(err, "url", URL)

		return []string{}
	}

	return res
}
