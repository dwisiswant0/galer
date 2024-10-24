package runner

const (
	author  = "dwisiswant0"
	version = "0.2.0"
	banner  = `
             __   v` + version + `
   __    _ _(_ )   __  _ __ 
 /'_ '\/'_' )| | /'__'( '__)
( (_) ( (_| || |(  ___| |
'\__  '\__,_(___'\____(_)
( )_) |
 \___/'  @` + author + `

`
	help = `A fast tool to fetch URLs from HTML attributes by crawl-in

Usage:
  galer -u [URL|URLs.txt] -o [output.txt]

Options:
  -u, --url <URL/FILE>        Target to fetches (single target URL or list)
  -e, --extension <EXT>       Show only certain extensions (comma-separated, e.g. js,php)
  -c, --concurrency <N>       Concurrency level (default: 50)
  -w, --wait <N>              Wait N seconds before evaluate (default: 1)
  -d, --depth <N>             Max. depth for crawling (levels of links to follow)
      --same-host             Same host only
      --same-root             Same root (eTLD+1) only (takes precedence over --same-host)
  -o, --output <FILE>         Save fetched URLs output into file
  -T, --template <string>     Format for output template (e.g., "{{scheme}}://{{host}}{{path}}")
                              Valid variables are: "raw_url", "scheme", "user", "username",
                              "password", "host", "hostname", "port", "path", "raw_path",
                              "escaped_path", "raw_query", "fragment", "raw_fragment".
  -t, --timeout <N>           Max. time (seconds) allowed for connection (default: 60)
  -s, --silent                Silent mode (suppress an errors)
  -v, --verbose               Verbose mode show error details unless you weren't use silent
  -h, --help                  Display its helps

Examples:
  galer -u http://domain.tld
  galer -u urls.txt -o output.txt
  cat urls.txt | galer -o output.txt

`
)
