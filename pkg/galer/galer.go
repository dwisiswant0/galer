package galer

import (
	"context"
	"errors"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/exp/slices"
)

// Config declare its configurations
type Config struct {
	Timeout int
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
	var res []string

	if !IsURI(URL) {
		return nil, errors.New("cannot parse URL")
	}

	ctx, cancel := chromedp.NewContext(cfg.ctx)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent: // Outgoing requests
			url := ev.Request.URL
			if !IsURI(url) {
				break
			}

			if !slices.Contains(res, url) {
				res = append(res, url)
			}
		}
	})

	err := chromedp.Run(ctx, chromedp.Navigate(URL), chromedp.Evaluate(script, &res))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cfg *Config) Close() error {
	cfg.cancel()

	return chromedp.Cancel(cfg.ctx)
}
