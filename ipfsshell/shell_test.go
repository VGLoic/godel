package ipfsshell

import (
	"encoding/json"
	"errors"
	"testing"
)

type mockInternalShell struct {
	returnedCid string
}

var testError = errors.New("test error")

func (s *mockInternalShell) DagGet(cid string, out interface{}) error {
	if cid == "" {
		return testError
	}
	out.(*BusinessEvent).Type = "type"
	out.(*BusinessEvent).Payload = "payload"
	out.(*BusinessEvent).Version = 0.1
	return nil
}
func (s *mockInternalShell) DagPut(data interface{}, ienc string, kind string) (string, error) {
	emptyMarsh, _ := json.Marshal(BusinessEvent{})
	if len(data.([]byte)) == len(emptyMarsh) {
		return "", testError
	}
	return s.returnedCid, nil
}
func (s *mockInternalShell) Pin(path string) error {
	if path == "" {
		return testError
	}
	return nil
}

func newMockShell(returnedCid string) *Shell {
	return &Shell{
		goIpfsShell: &mockInternalShell{returnedCid: returnedCid},
	}
}

func TestGetBusinessEvent(t *testing.T) {
	s := newMockShell("test")
	testBusinessEvent := BusinessEvent{
		Type:    "type",
		Payload: "payload",
		Version: 0.1,
	}

	t.Run("when the business event is found", func(t *testing.T) {
		businessEvent, err := s.GetBusinessEvent("known")
		if err != nil {
			t.Errorf("err = %v; want %v", err, nil)
		}
		if businessEvent != testBusinessEvent {
			t.Errorf("businessEvent = %v; want %v", businessEvent, testBusinessEvent)
		}
	})

	t.Run("when the business event is not found", func(t *testing.T) {
		businessEvent, err := s.GetBusinessEvent("")
		if err != testError {
			t.Errorf("err = %v; want %v", err, testError)
		}
		emptyBusinessEvent := BusinessEvent{}
		if businessEvent != emptyBusinessEvent {
			t.Errorf("businessEvent = %v; want %v", businessEvent, BusinessEvent{})
		}
	})
}

func TestPublishBusinessEvent(t *testing.T) {
	testBusinessEvent := BusinessEvent{
		Type:    "type",
		Payload: "payload",
		Version: 0.1,
	}

	t.Run("when the publish to IPFS is successful", func(t *testing.T) {
		s := newMockShell("test")
		cid, err := s.PublishBusinessEvent(testBusinessEvent)
		if err != nil {
			t.Errorf("err = %v; want %v", err, nil)
		}
		if cid != "test" {
			t.Errorf("cid = %v; want %v", cid, "test")
		}
	})

	t.Run("when the publish to IPFS fails", func(t *testing.T) {
		s := newMockShell("test")
		cid, err := s.PublishBusinessEvent(BusinessEvent{})
		if err != testError {
			t.Errorf("err = %v; want %v", err, testError)
		}
		if cid != "" {
			t.Errorf("cid = %v; want %v", cid, "")
		}
	})

	t.Run("when the pinning of the cid fails", func(t *testing.T) {
		s := newMockShell("")
		cid, err := s.PublishBusinessEvent(BusinessEvent{})
		if err != testError {
			t.Errorf("err = %v; want %v", err, testError)
		}
		if cid != "" {
			t.Errorf("cid = %v; want %v", cid, "")
		}
	})
}
