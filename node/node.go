package node

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	"github.com/VGLoic/godel/api"
	"github.com/VGLoic/godel/backend"
	"github.com/VGLoic/godel/config"
	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
	"github.com/spf13/viper"
)

type Godel struct {
	b *backend.Backend
}

func NewGodelNode(ctx context.Context) (*Godel, error) {
	configErr := config.LoadConfig()
	if configErr != nil {
		return nil, configErr
	}

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

	ipfsShell, ipfsShellErr := ipfsshell.NewShell(ipfsshell.ShellConfiguration{
		IpfsNodeUrl:   viper.GetString("IPFS_DAEMON_URL"),
		ProjectId:     viper.GetString("IPFS_PROJECT_ID"),
		ProjectSecret: viper.GetString("IPFS_PROJECT_SECRET"),
	})
	if ipfsShellErr != nil {
		return nil, ipfsShellErr
	}

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

	go g.b.MakeLocalDataAvailable(ctx)

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
	api := api.NewApi(g.b)
	rpc.Register(api)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		return e
	}
	return http.Serve(l, nil)
}
