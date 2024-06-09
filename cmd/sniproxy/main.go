package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/awryme/slogf"
	"github.com/awryme/sniproxy/cmd/sniproxy/proxyserver"
	"github.com/oklog/run"
)

type App struct {
	ListenAddr      string `help:"network address to listen on" default:"127.0.0.1"`
	ListenPortHTTP  int    `help:"http port to listen on" default:"80"`
	ListenPortHTTPS int    `help:"https port to listen on" default:"443"`
}

func (app *App) Run() error {
	logf := slogf.New(slogf.DefaultHandler(os.Stdout))
	printBuildInfo(logf)

	ctx := context.Background()
	// todo: fix cancelation

	group := new(run.Group)
	onInterrupt := func(error) {}

	makeServer := func(serverType proxyserver.ProxyType, port int) func() error {
		return func() error {
			addr := app.ListenAddr
			err := proxyserver.Start(ctx, logf, addr, port, serverType)
			if err != nil {
				return fmt.Errorf("running server failed: %w (type = %s, addr = %s, port = %d)", err, serverType, addr, port)
			}
			return nil
		}
	}
	// group.Add(run.SignalHandler(ctx, syscall.SIGINT))

	group.Add(makeServer(proxyserver.ProxyTypeHTTP, app.ListenPortHTTP), onInterrupt)
	group.Add(makeServer(proxyserver.ProxyTypeHTTPS, app.ListenPortHTTPS), onInterrupt)

	return group.Run()
}

func main() {
	var app App
	kctx := kong.Parse(&app, kong.DefaultEnvars("SNIPROXY"))
	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}
