package ipfsshell

import (
	"encoding/json"

	goIpfsShell "github.com/ipfs/go-ipfs-api"
)

type Shell struct {
	goIpfsShell internalShell
}

type ShellConfiguration struct {
	IpfsNodeUrl string
}

type BusinessEvent struct {
	Type    string  `json:"type"`
	Payload string  `json:"payload"`
	Version float64 `json:"version"`
}

func NewShell(shellConfiguration ShellConfiguration) (*Shell, error) {
	internalShell := goIpfsShell.NewShell(shellConfiguration.IpfsNodeUrl)
	ipfsShell := Shell{
		goIpfsShell: internalShell,
	}
	return &ipfsShell, nil
}

func (s *Shell) GetBusinessEvent(cid string) (BusinessEvent, error) {
	businessEvent := BusinessEvent{}
	errGet := s.goIpfsShell.DagGet(cid, &businessEvent)
	if errGet != nil {
		return businessEvent, errGet
	}
	return businessEvent, nil
}

func (s *Shell) PublishBusinessEvent(event BusinessEvent) (string, error) {
	marshalledEvent, marshalErr := json.Marshal(event)
	if marshalErr != nil {
		return "", marshalErr
	}
	cid, errIpfsPut := s.goIpfsShell.DagPut(marshalledEvent, "json", "cbor")
	if errIpfsPut != nil {
		return "", errIpfsPut
	}

	return cid, nil
}
