package api

import (
	"github.com/VGLoic/godel/backend"
	"github.com/VGLoic/godel/eventlog"
)

type Backend interface {
	PublishEvent(publishRequest backend.PublishRequest) (eventlog.Event, error)
	FindPendingEvents() ([]eventlog.Event, error)
	FindConfirmedEvents() ([]eventlog.Event, error)
}
