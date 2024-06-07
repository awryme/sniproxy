package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/awryme/sniproxy/cmd/sniproxy/proxyserver"
	"github.com/awryme/sniproxy/pkg/logging"
)

type App struct {
	ListenAddr      string `help:"network address to listen on" default:"127.0.0.1"`
	ListenPortHTTP  int    `help:"http port to listen on" default:"80"`
	ListenPortHTTPS int    `help:"https port to listen on" default:"443"`
}

func (app *App) Run() error {
	logf := logging.NewLogf(os.Stdout)

	ctx := context.Background()
	// todo: fix cancelation

	errorQueue := make(chan error, 2)
	runServer := func(serverType proxyserver.ProxyType, port int) {
		addr := app.ListenAddr
		err := proxyserver.Start(ctx, logf, addr, port, serverType)
		if err != nil {
			err = fmt.Errorf("running server failed: %w (type = %s, addr = %s, port = %d)", err, serverType, addr, port)
		}
		errorQueue <- err
	}
	go runServer(proxyserver.ProxyTypeHTTP, app.ListenPortHTTP)
	go runServer(proxyserver.ProxyTypeHTTPS, app.ListenPortHTTPS)

	return <-errorQueue
}

func main() {
	var app App
	kctx := kong.Parse(&app, kong.DefaultEnvars("SNIPROXY"))
	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}
