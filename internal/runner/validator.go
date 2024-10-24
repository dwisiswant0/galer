package runner

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func (opt *Options) validate() error {
	var errFile error

	if isStdin() {
		opt.List = bufio.NewScanner(os.Stdin)
	} else if opt.URL != "" {
		if strings.HasPrefix(opt.URL, "http") {
			opt.List = bufio.NewScanner(strings.NewReader(opt.URL))
		} else {
			f, err := os.Open(opt.URL)
			if err != nil {
				return err
			}
			opt.List = bufio.NewScanner(f)
		}
	} else {
		return errors.New("no target inputs provided")
	}

	if opt.Output != "" {
		opt.File, errFile = os.OpenFile(opt.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if errFile != nil {
			return errFile
		}
	}

	return nil
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

// func isScope(target string, URL string) bool {
// 	t, e := url.Parse(target)
// 	if e != nil {
// 		return false
// 	}

// 	u, e := url.Parse(URL)
// 	if e != nil {
// 		return false
// 	}

// 	return t.Host == u.Host
// }

func (opt *Options) isOnExt(URL string) bool {
	for _, e := range strings.Split(opt.Ext, ",") {
		if strings.TrimLeft(filepath.Ext(URL), ".") == e {
			return true
		}
	}

	return false
}
