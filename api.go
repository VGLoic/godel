package main

import (
	"context"

	"github.com/VGLoic/go-del/backend"
	"github.com/VGLoic/go-del/eventlog"
)

type Api struct {
	b *backend.Backend
}

type PublishEventArgs struct {
	Type    string
	Payload string
	Version float64
}

func NewApi(b *backend.Backend) *Api {
	return &Api{b: b}
}

func (api *Api) PublishEvent(args *PublishEventArgs, reply *eventlog.Event) error {
	req := backend.PublishRequest{
		Type:    args.Type,
		Payload: args.Payload,
		Version: args.Version,
	}
	event, err := api.b.PublishEvent(context.Background(), req)
	if err != nil {
		return err
	}
	reply = &event
	return nil
}

func (api *Api) FindPendingEvents(_ interface{}, reply *[]eventlog.Event) error {
	events, err := api.b.FindPendingEvents()
	if err != nil {
		return err
	}
	reply = &events
	return nil
}

func (api *Api) FindConfirmedEvents(_ interface{}, reply *[]eventlog.Event) error {
	events, err := api.b.FindConfirmedEvents()
	if err != nil {
		return err
	}
	reply = &events
	return nil
}
