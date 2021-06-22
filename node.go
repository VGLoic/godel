package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/VGLoic/go-del/backend"
	"github.com/VGLoic/go-del/config"
	"github.com/VGLoic/go-del/ethshell"
	"github.com/VGLoic/go-del/eventlog"
	"github.com/VGLoic/go-del/ipfsshell"
	"github.com/spf13/viper"
)

type Godel struct {
	b *backend.Backend
}

func NewGodelNode(ctx context.Context) (*Godel, error) {
	fmt.Println("-- Load configuration --")
	configErr := config.LoadConfig()
	if configErr != nil {
		return nil, configErr
	}

	fmt.Println("-- Connect to Database and migrate EventLog --")
	eventLog, eventLogErr := eventlog.NewEventLog(eventlog.EventLogConfiguration{
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresDb:       viper.GetString("POSTGRES_DB"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
	})
	if eventLogErr != nil {
		return nil, eventLogErr
	}

	fmt.Println("-- Connect IPFS shell --")
	ipfsShell, ipfsShellErr := ipfsshell.NewShell(ipfsshell.ShellConfiguration{
		IpfsNodeUrl: viper.GetString("IPFS_DAEMON_URL"),
	})
	if ipfsShellErr != nil {
		return nil, ipfsShellErr
	}

	fmt.Println("-- Connect Eth shell --")
	ethShell, ethShellErr := ethshell.NewShell(ethshell.ShellConfiguration{
		EthNodeUrl:      viper.GetString("ETH_NODE_URL"),
		ContractAddress: viper.GetString("CONTRACT_ADDRESS"),
		PrivateKey:      viper.GetString("PRIVATE_KEY"),
	})
	if ethShellErr != nil {
		return nil, ethShellErr
	}

	godel := Godel{
		b: backend.NewBackend(eventLog, ipfsShell, ethShell, viper.GetString("ACCOUNT_ADDRESS")),
	}

	return &godel, nil
}

func (g *Godel) Start(ctx context.Context) error {
	synchronisationErr := g.b.SynchroniseAllTopics(ctx)
	if synchronisationErr != nil {
		return synchronisationErr
	}

	subscriptionErr := g.b.SubscribeToEvents(ctx)
	if subscriptionErr != nil {
		return subscriptionErr
	}

	return nil
}

func (g *Godel) ServeApi() error {
	api := NewApi(g.b)
	rpc.Register(api)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		return e
	}
	return http.Serve(l, nil)
}
