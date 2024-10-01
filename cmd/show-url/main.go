// show-url is a command-line tool for opening a URL in a sfomuseum/go-www-show/v2.Browser environment.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sfomuseum/go-www-show/v2"
)

func main() {

	var browser_uri string
	var url string

	browser_schemes := show.BrowserSchemes()
	str_schemes := strings.Join(browser_schemes, ",")

	browser_desc := fmt.Sprintf("A valid sfomuseum/go-www-show/v2.Browser URI. Valid options are: %s", str_schemes)

	flag.StringVar(&browser_uri, "browser-uri", "web://", browser_desc)
	flag.StringVar(&url, "url", "", "The URL to open")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Open a URL in a sfomuseum/go-www-show/v2.Browser environment.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	br, err := show.NewBrowser(ctx, browser_uri)

	if err != nil {
		log.Fatalf("Failed to create new browser, %w", err)
	}

	done_ch := make(chan bool)
	err = br.OpenURL(ctx, url, done_ch)

	if err != nil {
		log.Fatalf("Failed to open URL, %w", err)
	}
}
