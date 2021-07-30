// +build integration

package integration

import (
	"net/rpc"
	"regexp"
	"testing"
	"time"

	"github.com/VGLoic/godel/api"
	"github.com/VGLoic/godel/eventlog"
	_require "github.com/stretchr/testify/require"
)

func TestPublishEvent(t *testing.T) {
	require := _require.New(t)

	serverAddress := "localhost:1234"
	rpcClient, err := rpc.DialHTTP("tcp", serverAddress)
	require.Nil(err)
	defer rpcClient.Close()

	publishEventArgs := api.PublishEventArgs{
		Type:        "integration/type",
		Payload:     "{}",
		Version:     0.1,
		Topic:       "integration",
		NewAccounts: []string{},
	}
	publishedEvent := publishEvent(rpcClient, require, publishEventArgs)

	waitForConfirmation(rpcClient, require, publishEventArgs, publishedEvent)
}

func publishEvent(rpcClient *rpc.Client, require *_require.Assertions, publishEventArgs api.PublishEventArgs) eventlog.Event {
	publishEventReply := eventlog.Event{}
	err := rpcClient.Call("Api.PublishEvent", &publishEventArgs, &publishEventReply)
	require.Nil(err)

	checkEvent(require, publishEventArgs, publishEventReply)
	require.Regexp(regexp.MustCompile("^0x[0-9a-fA-F]{64}$"), publishEventReply.TxHash, "TxHash should have proper format")

	return publishEventReply
}

func waitForConfirmation(rpcClient *rpc.Client, require *_require.Assertions, publishEventArgs api.PublishEventArgs, publishedEvent eventlog.Event) eventlog.Event {
	retryCount := 0
	isConfirmed := false
	var confirmedEvent eventlog.Event
	for retryCount < 10 && !isConfirmed {
		var queryArgs interface{}
		var confirmedQueryReply []eventlog.Event
		err := rpcClient.Call("Api.FindConfirmedEvents", &queryArgs, &confirmedQueryReply)
		require.Nil(err)
		isContained, event := contains(confirmedQueryReply, publishedEvent.ID)

		if isContained {
			confirmedEvent = event
			isConfirmed = true
		} else {
			var pendingQueryReply []eventlog.Event
			err = rpcClient.Call("Api.FindPendingEvents", &queryArgs, &pendingQueryReply)
			require.Nil(err)
			isContained, event := contains(pendingQueryReply, publishedEvent.ID)
			require.Equal(isContained, true, "Event must be pending")
			checkEvent(require, publishEventArgs, event)
			require.Equal(event.Timestamp, uint64(0), "Timestamp should be equal to 0")
			require.Equal(event.BlockNumber, uint64(0), "Blocknumber should be equal to 0")
			require.Equal(event.Depth, uint64(0), "Depth should be equal to 0")

			retryCount += 1
			time.Sleep(1 * time.Second)
		}

	}

	checkEvent(require, publishEventArgs, confirmedEvent)
	require.Greater(confirmedEvent.Timestamp, uint64(0), "Timestamp should be greater than 0")
	require.Greater(confirmedEvent.BlockNumber, uint64(0), "Blocknumber should be greater than 0")
	require.Greater(confirmedEvent.Depth, uint64(0), "Depth should be greater than 0")

	return confirmedEvent
}

func checkEvent(require *_require.Assertions, publishEventArgs api.PublishEventArgs, event eventlog.Event) {
	require.Equal(event.Type, publishEventArgs.Type, "Type should match")
	require.Equal(event.Payload, publishEventArgs.Payload, "Payload should match")
	require.Equal(event.Version, publishEventArgs.Version, "Version should match")
	require.Equal(event.Topic, publishEventArgs.Topic, "Topic should match")
	require.ElementsMatch(event.NewAccounts, publishEventArgs.NewAccounts, "New accounts should match")
	require.Equal(event.Cid, "bafyreiags3je4wzlkagusfwh4yufoyymware2lnaq4n7h74eonjcjqmrcq", "Cid should be the expected one")
	require.Greater(event.ID, uint(0), "ID should be greater or equal to one")
}

func contains(arr []eventlog.Event, id uint) (bool, eventlog.Event) {
	for _, e := range arr {
		if e.ID == id {
			return true, e
		}
	}
	return false, eventlog.Event{}
}
