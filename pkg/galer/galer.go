package galer

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/valyala/fasttemplate"
	"golang.org/x/exp/slices"
	"golang.org/x/net/publicsuffix"
)

// Config declare its configurations
type Config struct {
	Logger   *log.Logger
	SameHost bool
	SameRoot bool
	Template string
	Timeout  int
	Wait     int

	// Headers network.Headers
	ctx      context.Context
	cancel   context.CancelFunc
	template *fasttemplate.Template

	scope struct {
		hostname, root string
	}
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

	// defaulting sleep
	if cfg.Wait <= 0 {
		cfg.Wait = 1
	}

	// defaulting scope (hostname & root)
	if cfg.scope.hostname == "" && cfg.scope.root == "" {
		u, _ := url.Parse(URL)
		cfg.scope.hostname = u.Hostname()
		cfg.scope.root, _ = publicsuffix.EffectiveTLDPlusOne(cfg.scope.hostname)
	}

	var ctxOpts []chromedp.ContextOption
	if cfg.Logger != nil {
		ctxOpts = []chromedp.ContextOption{
			chromedp.WithLogf(cfg.Logger.Printf),
			chromedp.WithDebugf(cfg.Logger.Debugf),
			chromedp.WithErrorf(cfg.Logger.Errorf),
		}
	}

	ctx, cancel := chromedp.NewContext(cfg.ctx, ctxOpts...)
	defer cancel()

	if cfg.Template != "" {
		cfg.template = fasttemplate.New(cfg.Template, "{{", "}}")
	}

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
		chromedp.Sleep(1*time.Second),
		chromedp.Evaluate(script, &res),
	)
	if err != nil {
		return nil, err
	}

	// template eval
	for i, _ := range res {
		res[i] = cfg.eval(res[i])
	}

	for i, _ := range reqs {
		reqs[i] = cfg.eval(reqs[i])
	}

	res = MergeSlices(res, reqs)

	// filters
	switch {
	case cfg.SameRoot:
		for i := 0; i < len(res); i++ {
			r, _ := url.Parse(res[i])
			base, err := publicsuffix.EffectiveTLDPlusOne(r.Hostname())
			if err != nil && cfg.Logger != nil {
				cfg.Logger.Error("could not get eTLD+1", "parsed", r.String())
			}

			if !strings.HasSuffix(cfg.scope.root, base) {
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
	case cfg.SameHost:
		for i := 0; i < len(res); i++ {
			r, _ := url.Parse(res[i])
			if r.Hostname() != cfg.scope.hostname {
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
