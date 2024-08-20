package show

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	port := opts.Port

	if port == 0 {

		listener, err := net.Listen("tcp", "localhost:0")

		if err != nil {
			return fmt.Errorf("Failed to determine next available port, %w", err)
		}

		port = listener.Addr().(*net.TCPAddr).Port
		err = listener.Close()

		if err != nil {
			return fmt.Errorf("Failed to close listener used to derive port, %w", err)
		}
	}

	//

	addr := fmt.Sprintf("localhost:%d", port)
	url := fmt.Sprintf("http://%s", addr)

	http_server := http.Server{
		Addr: addr,
	}

	http_server.Handler = opts.Mux

	sigint := make(chan bool)
	err_ch := make(chan error)

	go func() {

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		slog.Info("Shutting server down")
		err := http_server.Shutdown(ctx)

		if err != nil {
			slog.Error("HTTP server shutdown error", "error", err)
		}

	}()

	go func() {

		err := http_server.ListenAndServe()

		if err != nil {
			err_ch <- fmt.Errorf("Failed to start server, %w", err)
		}
	}()

	server_ready := false

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case err := <-err_ch:
			log.Fatalf("Received error starting server, %v", err)
		case <-ticker.C:

			rsp, err := http.Head(url)

			if err != nil {
				slog.Warn("HEAD request failed", "url", url, "error", err)
			} else {

				defer rsp.Body.Close()

				if rsp.StatusCode != 200 {
					slog.Warn("HEAD request did not return expected status code", "url", url, "code", rsp.StatusCode)
				} else {
					slog.Debug("HEAD request succeeded", "url", url)
					server_ready = true
				}
			}
		}

		if server_ready {
			break
		}
	}

	err := opts.Browser.OpenURL(ctx, url)

	if err != nil {
		log.Fatalf("Failed to open URL %s, %v", url, err)
	}

	log.Printf("Features are viewable at %s\n", url)
	<-sigint

	return nil
}
