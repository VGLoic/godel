package ipfsshell

import (
	"encoding/json"
	"net/http"

	goIpfsShell "github.com/ipfs/go-ipfs-api"
)

type Shell struct {
	goIpfsShell internalShell
}

type ShellConfiguration struct {
	IpfsNodeUrl   string
	ProjectId     string
	ProjectSecret string
}

type BusinessEvent struct {
	Type    string  `json:"type"`
	Payload string  `json:"payload"`
	Version float64 `json:"version"`
}

func NewShell(shellConfiguration ShellConfiguration) (*Shell, error) {
	internalShell := goIpfsShell.NewShellWithClient(
		shellConfiguration.IpfsNodeUrl,
		newClient(shellConfiguration.ProjectId, shellConfiguration.ProjectSecret),
	)
	ipfsShell := Shell{
		goIpfsShell: internalShell,
	}
	return &ipfsShell, nil
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func newClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
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
	pinErr := s.goIpfsShell.Pin(cid)
	if pinErr != nil {
		return "", pinErr
	}

	return cid, nil
}

func (s *Shell) PinBusinessEventCid(cid string) error {
	return s.goIpfsShell.Pin(cid)
}
