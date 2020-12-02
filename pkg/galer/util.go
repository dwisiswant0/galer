package galer

import "net/url"

const script = "[...new Set(Array.from(document.querySelectorAll('[src],[href],[url],[action]')).map(i => i.src || i.href || i.url || i.action))]"

// IsURI detect valid URI
func IsURI(s string) bool {
	_, e := url.ParseRequestURI(s)
	if e != nil {
		return false
	}

	u, e := url.Parse(s)
	if e != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
