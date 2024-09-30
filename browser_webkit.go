//go:build webkit

package show

import (
	"context"
	"fmt"
	_ "log/slog"
	"net/url"
	"strconv"

	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/webkit"
	"github.com/progrium/darwinkit/objc"
)

// WebKitBrowser implements the `Browser` interface for loading URLs in a WebKit-based WebView window.
type WebKitBrowser struct {
	Browser
	width  float64
	height float64
}

func init() {

	ctx := context.Background()

	err := RegisterBrowser(ctx, "webkit", NewWebKitBrowser)

	if err != nil {
		panic(err)
	}
}

// NewWebKitBrowser retuns a new `WebKitBrowser` instance configured by 'uri' which is expected to take the form of:
//
//	webkit://?{PARAMETERS}
//
// Where valid parameters are:
// * width={INT} – Define the width of the window to open. Default is 1024.
// * height={INT} – Define the height of the window to open. Default is 800.
func NewWebKitBrowser(ctx context.Context, uri string) (Browser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	width := 1024.0
	height := 800.0

	if q.Has("width") {

		v, err := strconv.ParseFloat(q.Get("width"), 64)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?width= parameter, %w", err)
		}

		width = v
	}

	if q.Has("height") {

		v, err := strconv.ParseFloat(q.Get("height"), 64)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?height= parameter, %w", err)
		}

		height = v
	}

	br := &WebKitBrowser{
		width:  width,
		height: height,
	}

	return br, nil
}

// OpenURL opens 'url' in a WebKit-based WebView window.
func (br *WebKitBrowser) OpenURL(ctx context.Context, url string) error {

	if url == "" {
		return fmt.Errorf("URL is empty")
	}

	launch_url := func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		w := appkit.NewWindowWithSize(br.width, br.height)
		objc.Retain(&w)

		configuration := webkit.NewWebViewConfiguration()

		view := webkit.NewWebViewWithFrameConfiguration(foundation.Rect{}, configuration)

		/*
		webkit.AddScriptMessageHandlerWithReply(view, "greet", func(message objc.Object) (objc.Object, error) {
			param := message.Description()
			fmt.Println("greet handled")
			return foundation.NewStringWithString("hello: " + param).Object, nil
		})
		*/
		
		w.SetContentView(view)
		w.MakeKeyAndOrderFront(nil)
		w.Center()

		webkit.LoadURL(view, url)

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

	}

	macos.RunApp(launch_url)
	return nil

	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {

		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		f_url := foundation.URL_URLWithString(url)
		req := foundation.NewURLRequestWithURL(f_url)

		frame := foundation.Rect{Size: foundation.Size{br.width, br.height}}

		config := webkit.NewWebViewConfiguration()
		wv := webkit.NewWebViewWithFrameConfiguration(frame, config)
		wv.LoadRequest(req)

		w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)

		objc.Retain(&w)
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})

	return nil
}
