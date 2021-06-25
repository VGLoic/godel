package api

import (
	"github.com/VGLoic/godel/backend"
	"github.com/VGLoic/godel/eventlog"
)

type Api struct {
	b Backend
}

type PublishEventArgs struct {
	Type        string
	Payload     string
	Version     float64
	Topic       string
	NewAccounts []string
}

func NewApi(b Backend) *Api {
	return &Api{b: b}
}

func (api *Api) PublishEvent(args *PublishEventArgs, reply *eventlog.Event) error {
	req := backend.PublishRequest{
		Type:        args.Type,
		Payload:     args.Payload,
		Version:     args.Version,
		Topic:       args.Topic,
		NewAccounts: args.NewAccounts,
	}
	event, err := api.b.PublishEvent(req)
	if err != nil {
		return err
	}
	*reply = event
	return nil
}

func (api *Api) FindPendingEvents(_ interface{}, reply *[]eventlog.Event) error {
	events, err := api.b.FindPendingEvents()
	if err != nil {
		return err
	}
	*reply = events
	return nil
}

func (api *Api) FindConfirmedEvents(_ interface{}, reply *[]eventlog.Event) error {
	events, err := api.b.FindConfirmedEvents()
	if err != nil {
		return err
	}
	*reply = events
	return nil
}
