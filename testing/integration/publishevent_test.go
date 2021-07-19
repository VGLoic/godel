// +build integration

package integration

import (
	"fmt"
	"net/rpc"
	"testing"

	"github.com/VGLoic/godel/api"
	"github.com/VGLoic/godel/eventlog"
)

func TestPublishEvent(t *testing.T) {
	serverAddress := "localhost:1234"
	rpcClient, dialErr := rpc.DialHTTP("tcp", serverAddress)
	if dialErr != nil {
		t.Errorf("err = %v; want %v", dialErr, nil)
	}
	defer rpcClient.Close()

	args := api.PublishEventArgs{
		Type:        "integration/type",
		Payload:     "{}",
		Version:     0.1,
		Topic:       "integration",
		NewAccounts: []string{"hello"},
	}
	reply := eventlog.Event{}
	callErr := rpcClient.Call("Api.PublishEvent", &args, &reply)
	if callErr != nil {
		t.Errorf("err = %v; want %v", callErr, nil)
	}
	fmt.Println("REPLY: ", reply)
	// reply:
	// {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC {0001-01-01 00:00:00 +0000 UTC false}} 8 integration/type {} 0.1 integration bafyreiags3je4wzlkagusfwh4yufoyymware2lnaq4n7h74eonjcjqmrcq 0xf61B443A155b07D2b2cAeA2d99715dC84E839EEf [hello] 0 0 0 0xcbd4f8f02c48a031ad8e8a35433fcfdf8044563800cea7bc5c514aec87a9e855 2021-07-19 15:17:11.291413076 +0000 +0000 2021-07-19 15:17:11.291413076 +0000 +0000}
}
