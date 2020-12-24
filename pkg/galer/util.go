package galer

import (
	"github.com/projectdiscovery/gologger"
	"io/ioutil"
	"net/url"
	"strings"
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

// LoadUrlsFromFile is method to load url form text file split by new line
// path is string path can be absolute path or relative path
func LoadUrlsFromFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		gologger.Errorf("Unable to load file: %s", err)
	}
	return strings.Split(string(content), "\n")
}
