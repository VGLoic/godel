package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VGLoic/godel/config"
	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
	"github.com/VGLoic/godel/node"
	"github.com/spf13/viper"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configuration: %s \n", err))
	}

	fmt.Println("-- Creating Godel Node --")
	eventLogConfig := eventlog.EventLogConfiguration{
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresDb:       viper.GetString("POSTGRES_DB"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
	}
	ipfsShellConfig := ipfsshell.ShellConfiguration{
		IpfsNodeUrl:   viper.GetString("IPFS_DAEMON_URL"),
		ProjectId:     viper.GetString("IPFS_PROJECT_ID"),
		ProjectSecret: viper.GetString("IPFS_PROJECT_SECRET"),
	}
	ethShellConfig := ethshell.ShellConfiguration{
		EthNodeUrl:      viper.GetString("ETH_NODE_URL"),
		ContractAddress: viper.GetString("CONTRACT_ADDRESS"),
		PrivateKey:      viper.GetString("PRIVATE_KEY"),
	}
	godelNode, err := node.NewGodelNode(
		eventLogConfig,
		ipfsShellConfig,
		ethShellConfig,
		viper.GetString("ACCOUNT_ADDRESS"),
	)
	if err != nil {
		panic(fmt.Errorf("Fatal error godel node: %s \n", err))
	}

	fmt.Println("-- Starting Godel Node --")
	err = godelNode.Start(ctx)
	defer godelNode.Stop()
	if err != nil {
		panic(fmt.Errorf("Fatal error godel node start: %s \n", err))
	}

	fmt.Println("-- Serving the API --")
	log.Fatal(godelNode.ServeApi())
}
