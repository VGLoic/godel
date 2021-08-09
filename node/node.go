package node

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	"github.com/VGLoic/godel/api"
	"github.com/VGLoic/godel/backend"
	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
)

type Godel struct {
	b      *backend.Backend
	cancel context.CancelFunc
}

func NewGodelNode(
	eventLogConfig eventlog.EventLogConfiguration,
	ipfsShellConfig ipfsshell.ShellConfiguration,
	ethShellConfig ethshell.ShellConfiguration,
	accountAddress string,
) (*Godel, error) {

	eventLog, err := eventlog.NewEventLog(eventLogConfig)
	if err != nil {
		return nil, err
	}

	ipfsShell, err := ipfsshell.NewShell(ipfsShellConfig)
	if err != nil {
		return nil, err
	}

	ethShell, err := ethshell.NewShell(ethShellConfig)
	if err != nil {
		return nil, err
	}

	godel := Godel{
		b: backend.NewBackend(eventLog, ipfsShell, ethShell, accountAddress),
	}

	return &godel, nil
}

func (g *Godel) Start(parentCtx context.Context) error {

	ctx, cancel := context.WithCancel(parentCtx)
	g.cancel = cancel

	go g.b.MakeLocalDataAvailable(ctx)

	err := g.b.SynchroniseAllTopics(ctx)
	if err != nil {
		return err
	}

	err = g.b.SubscribeToEvents(ctx)
	if err != nil {
		return err
	}

	return g.serveApi(ctx)
}

func (g *Godel) Stop() error {
	g.cancel()
	return nil
}

func (g *Godel) serveApi(ctx context.Context) error {
	api := api.NewApi(g.b)
	rpc.Register(api)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		return e
	}
	errChan := make(chan error)
	go func() {
		err := http.Serve(l, nil)
		errChan <- err
	}()
	select {
	case <-ctx.Done():
		err := l.Close()
		if err != nil {
			return err
		}
		return nil
	case err := <-errChan:
		return err
	}
}
