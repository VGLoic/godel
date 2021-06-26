package api

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VGLoic/godel/backend"
	"github.com/VGLoic/godel/eventlog"
)

type mockBackend struct{}

var testError = errors.New("test error")
var testPendingEvent = eventlog.Event{
	Type: "type",
}
var testConfirmedEvent = eventlog.Event{
	Type:        "type",
	BlockNumber: 123,
	Timestamp:   123,
}

func (b *mockBackend) PublishEvent(req backend.PublishRequest) (eventlog.Event, error) {
	if req.Topic == "" {
		return eventlog.Event{}, testError
	}
	return testPendingEvent, nil
}

func (b *mockBackend) FindPendingEvents() ([]eventlog.Event, error) {
	return []eventlog.Event{testPendingEvent}, nil
}
func (b *mockBackend) FindConfirmedEvents() ([]eventlog.Event, error) {
	return []eventlog.Event{testConfirmedEvent}, nil
}

func newMockApi() *Api {
	return &Api{b: &mockBackend{}}
}

func TestPublishEvent(t *testing.T) {
	api := newMockApi()
	t.Run("when the internal `PublishEvent` fails", func(t *testing.T) {
		var reply eventlog.Event
		err := api.PublishEvent(&PublishEventArgs{Topic: ""}, &reply)
		if err != testError {
			t.Errorf("err = %v; want %v", err, testError)
		}
		var emptyEvent eventlog.Event
		if !reflect.DeepEqual(reply, emptyEvent) {
			t.Errorf("reply = %v; want %v", reply, emptyEvent)
		}
	})
	t.Run("when the internal `PublishEvent` succeeds", func(t *testing.T) {
		var reply eventlog.Event
		err := api.PublishEvent(&PublishEventArgs{Topic: "hello"}, &reply)
		if err != nil {
			t.Errorf("err = %v; want %v", err, nil)
		}
		if !reflect.DeepEqual(reply, testPendingEvent) {
			t.Errorf("reply = %v; want %v", reply, testPendingEvent)
		}
	})
}

func TestFindPendingEvents(t *testing.T) {
	api := newMockApi()
	var arg interface{}
	var reply []eventlog.Event
	err := api.FindPendingEvents(arg, &reply)
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if !reflect.DeepEqual(reply, []eventlog.Event{testPendingEvent}) {
		t.Errorf("reply = %v; want %v", reply, []eventlog.Event{testPendingEvent})
	}
}
func TestFindConfirmedEvents(t *testing.T) {
	api := newMockApi()
	var arg interface{}
	var reply []eventlog.Event
	err := api.FindConfirmedEvents(arg, &reply)
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if !reflect.DeepEqual(reply, []eventlog.Event{testConfirmedEvent}) {
		t.Errorf("reply = %v; want %v", reply, []eventlog.Event{testConfirmedEvent})
	}
}
