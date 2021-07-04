package ipfsshell

type internalShell interface {
	DagGet(cid string, out interface{}) error
	DagPut(data interface{}, ienc string, kind string) (string, error)
	Pin(path string) error
}
