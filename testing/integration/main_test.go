// +build integration

package integration

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/ethshell/contract"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
	"github.com/VGLoic/godel/node"
	"github.com/VGLoic/godel/testing/dockerclient"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	viper.SetEnvPrefix("INTEGRATION")
	viper.AutomaticEnv()

	cli, err := dockerclient.NewDockerCli()
	if err != nil {
		log.Fatal(err)
	}

	mnemonic := viper.GetString("MNEMONIC")

	_, err = createPostgresContainer(ctx, cli)
	if err != nil {
		log.Fatal(err)
	}

	_, err = createGanacheContainer(ctx, cli, mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex := viper.GetString("PRIVATE_KEY")
	address, contractAddress, err := setupContract(ctx, privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	eventLogConfig := eventlog.EventLogConfiguration{
		PostgresHost:     "localhost",
		PostgresUser:     "admin",
		PostgresPassword: "admin",
		PostgresDb:       "event",
		PostgresPort:     "5432",
	}
	ipfsShellConfig := ipfsshell.ShellConfiguration{
		IpfsNodeUrl:   viper.GetString("IPFS_DAEMON_URL"),
		ProjectId:     viper.GetString("IPFS_PROJECT_ID"),
		ProjectSecret: viper.GetString("IPFS_PROJECT_SECRET"),
	}
	ethShellConfig := ethshell.ShellConfiguration{
		EthNodeUrl:      "ws://localhost:7545",
		ContractAddress: contractAddress,
		PrivateKey:      privateKeyHex,
	}
	godelNode, err := node.NewGodelNode(
		eventLogConfig,
		ipfsShellConfig,
		ethShellConfig,
		address,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = godelNode.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(godelNode.ServeApi())
	}()

	exitCode := m.Run()

	godelNode.Stop()
	cli.StopContainers(ctx)

	os.Exit(exitCode)
}

func createPostgresContainer(ctx context.Context, cli *dockerclient.DockerCli) (container.ContainerCreateCreatedBody, error) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "5432",
	}

	containerPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	hostConfig := container.HostConfig{
		PortBindings: portBinding,
	}
	env := []string{
		"POSTGRES_PORT=5432",
		"POSTGRES_DB=event",
		"POSTGRES_PASSWORD=admin",
		"POSTGRES_USER=admin",
	}
	containerConfig := container.Config{
		Image: "postgres",
		Env:   env,
	}
	return cli.CreateContainer(ctx, &containerConfig, &hostConfig, "db")
}
func createGanacheContainer(ctx context.Context, cli *dockerclient.DockerCli, mnemonic string) (container.ContainerCreateCreatedBody, error) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "7545",
	}

	containerPort, err := nat.NewPort("tcp", "8545")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	hostConfig := container.HostConfig{
		PortBindings: portBinding,
	}
	containerConfig := container.Config{
		Image: "trufflesuite/ganache-cli",
		Cmd:   strslice.StrSlice{"-m " + mnemonic},
	}
	return cli.CreateContainer(ctx, &containerConfig, &hostConfig, "ganache")
}

func setupContract(ctx context.Context, privateKeyHex string) (string, string, error) {
	retryCount := 0
	isConnected := false
	var client *ethclient.Client
	var err error
	for retryCount < 10 && !isConnected {
		client, err = ethclient.Dial("http://localhost:7545")
		if err == nil {
			_, err = client.BlockNumber(ctx)
			if err == nil {
				isConnected = true
				break
			}
		}

		retryCount += 1
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return "", "", err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", "", err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	auth.Context = ctx

	address, _, _, err := contract.DeployContract(auth, client)
	if err != nil {
		return "", "", err
	}
	return fromAddress.String(), address.String(), nil
}
