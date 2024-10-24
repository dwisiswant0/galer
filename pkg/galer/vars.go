package galer

import "github.com/chromedp/chromedp"

var execAllocOpts = append(
	chromedp.DefaultExecAllocatorOptions[:],
	chromedp.DisableGPU,
	chromedp.IgnoreCertErrors,
	// chromedp.Flag("headless", false),
)
