//go:build webview

package show

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"

	"github.com/webview/webview_go"
)

// WebviewBrowser implements the `Browser` interface for loading URLs in a Webview-based WebView window.
type WebviewBrowser struct {
	Browser
	width  int
	height int
}

func init() {

	ctx := context.Background()

	err := RegisterBrowser(ctx, "webview", NewWebviewBrowser)

	if err != nil {
		panic(err)
	}
}

// NewWebviewBrowser retuns a new `WebviewBrowser` instance configured by 'uri' which is expected to take the form of:
//
//	webview://?{PARAMETERS}
//
// Where valid parameters are:
// * width={INT} – Define the width of the window to open. Default is 1024.
// * height={INT} – Define the height of the window to open. Default is 800.
func NewWebviewBrowser(ctx context.Context, uri string) (Browser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	width := 1024
	height := 800

	if q.Has("width") {

		v, err := strconv.Atoi(q.Get("width"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?width= parameter, %w", err)
		}

		width = v
	}

	if q.Has("height") {

		v, err := strconv.Atoi(q.Get("height"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?height= parameter, %w", err)
		}

		height = v
	}

	br := &WebviewBrowser{
		width:  width,
		height: height,
	}

	return br, nil
}

// OpenURL opens 'url' in a Webview-based WebView window.
func (br *WebviewBrowser) OpenURL(ctx context.Context, url string) error {

	slog.Info("Open", "url", url)
	
	w := webview.New(false)
	defer w.Destroy()

	w.SetSize(br.width, br.height, webview.HintNone)
	w.Navigate(url)
	w.Run()
	
	return nil
}
