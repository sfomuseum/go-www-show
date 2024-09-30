# go-www-show

Go package for starting a local	webserver and then opening its URL in a target source once it (the web server) is running.

## Usage

This package is meant to be used by other packages that have configured a [http.Mux](https://pkg.go.dev/net/http#ServeMux) instance for serving web requests. The package will start a web server on localhost listening on a randomly chose port number (unless a user-defined value is provided) and then open that URL in a target environment (like a web browser).

```
import (
	"context"
	"net/http"

	"github.com/sfomuseum/go-www-show"
)

func main() {

	ctx := context.Background()
	
	mux := http.NewServeMux()
	
	// Configure handlers for mux here

	browser, _ = show.NewBrowser(ctx, "web://")
	
        show_opts := &www_show.RunOptions{
                Browser: browser,
                Mux:     mux,
        }

        return show.RunWithOptions(ctx, www_show_opts)
}
```

For a complete working example see the [sfomuseum/go-geojson-show](https://github.com/sfomuseum/go-geojson-show/blob/main/show.go) package.

## Browsers

This package defines a `Browser` interface for opening URLs.

### web://

Open URLs in a new window using the operating system's default web browser.

### webkit://

```
$> go build -mod vendor -tags webkit -ldflags="-s -w" -o bin/show cmd/show/main.go
```

```
$> ./bin/show -browser-uri 'webkit://?width=500' -url https://aaronland.info
```
