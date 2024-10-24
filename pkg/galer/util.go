package galer

import (
	"errors"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

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

// MergeSlices merges two slices of the same type into a
// single slice, removing duplicates.
func MergeSlices[T1 comparable, T2 []T1](v1, v2 T2) T2 {
	uniq := make(map[T1]struct{})
	for _, v := range v1 {
		uniq[v] = struct{}{}
	}

	for v := range uniq {
		v2 = append(v2, v)
	}

	return v2
}

// SetScope sets the host and root (eTLD+1) for config.
func (cfg *Config) SetScope(s string) {
	if u, err := url.Parse(s); err == nil {
		cfg.scope.hostname = u.Hostname()
		cfg.scope.root, _ = publicsuffix.EffectiveTLDPlusOne(u.Hostname())
	}
}

func (cfg *Config) eval(s string) string {
	u, err := url.Parse(s)
	if err != nil && cfg.Logger != nil {
		cfg.Logger.Errorf("cannot eval %q URL with %q as template: %+v", s, cfg.Template, errors.Unwrap(err))
		return s
	}

	if cfg.template == nil {
		return s
	}

	password, _ := u.User.Password()
	tags := map[string]interface{}{
		"raw_url":      u.String(),
		"scheme":       u.Scheme,
		"user":         u.User.String(),
		"username":     u.User.Username(),
		"password":     password,
		"host":         u.Host,
		"hostname":     u.Hostname(),
		"port":         u.Port(),
		"path":         u.Path,
		"raw_path":     u.RawPath,
		"escaped_path": u.EscapedPath(),
		"raw_query":    u.RawQuery,
		"fragment":     u.Fragment,
		"raw_fragment": u.RawFragment,
	}

	return cfg.template.ExecuteString(tags)
}
