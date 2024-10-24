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
