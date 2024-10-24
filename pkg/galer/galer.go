package galer

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/exp/slices"
	"golang.org/x/net/publicsuffix"
)

// Config declare its configurations
type Config struct {
	Timeout  int
	SameHost bool
	SameRoot bool

	// Headers network.Headers
	ctx    context.Context
	cancel context.CancelFunc
}

// New defines context for the configurations
func New(cfg *Config) *Config {
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), execAllocOpts...)
	cfg.ctx, _ = chromedp.NewContext(allocCtx)
	cfg.ctx, cfg.cancel = context.WithTimeout(cfg.ctx, time.Duration(cfg.Timeout)*time.Second)

	return cfg
}

// Crawl to navigate to the URL & dump URLs on it
func (cfg *Config) Crawl(URL string) ([]string, error) {
	var res, reqs []string

	if !IsURI(URL) {
		return nil, errors.New("cannot parse URL")
	}
	u, _ := url.Parse(URL)

	ctx, cancel := chromedp.NewContext(cfg.ctx)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent: // Outgoing requests
			url := ev.Request.URL
			if !IsURI(url) {
				break
			}

			if url == URL {
				break
			}

			if !slices.Contains(reqs, url) {
				reqs = append(reqs, url)
			}
		}
	})

	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.Evaluate(script, &res),
	)
	if err != nil {
		return nil, err
	}

	res = MergeSlices(res, reqs)

	// filters
	switch {
	case cfg.SameRoot:
		for i := 0; i < len(res); i++ {
			r, _ := url.Parse(res[i])
			base, _ := publicsuffix.EffectiveTLDPlusOne(r.Host)
			if base != u.Host {
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
	case cfg.SameHost:
		for i := 0; i < len(res); i++ {
			r, _ := url.Parse(res[i])
			if r.Host != u.Host {
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
	}

	return res, nil
}

func (cfg *Config) Close() error {
	cfg.cancel()

	return chromedp.Cancel(cfg.ctx)
}
