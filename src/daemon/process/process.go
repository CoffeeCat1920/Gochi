package process

import (
	"daemon/store"
	"net"
	"os"
)

type Request struct {
	Verb   string `json:"verb"`
	Object string `json:"object"`
}

type Response struct {
	Status string   `json:"status"`
	Error  string   `json:"error,omitempty"`
	Data   []string `json:"data"`
}

type Process struct {
	connStr  string
	listner  net.Listener
	store    *store.AppStore
	quitChan chan os.Signal
}

var (
	connStr string = "/tmp/mochi.sock"
	process *Process
)

func New() (*Process, error) {
	if process != nil {
		return process, nil
	}

	listner, err := net.Listen("unix", connStr)
	if err != nil {
		return nil, err
	}

	process = &Process{
		connStr:  connStr,
		listner:  listner,
		store:    store.New(),
		quitChan: make(chan os.Signal, 1),
	}

	return process, nil
}
