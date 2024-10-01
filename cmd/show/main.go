package main

import (
	"context"
	"flag"
	"log"

	"github.com/sfomuseum/go-www-show/v2"
)

func main() {

	var browser_uri string
	var url string

	flag.StringVar(&browser_uri, "browser-uri", "web://", "")
	flag.StringVar(&url, "url", "", "The URL to open")

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
