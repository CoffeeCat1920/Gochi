package main

import (
	"daemon/store"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Request struct {
	Action string `json:"action"`
	Name   string `json:"name"`
}

type Response struct {
	Status string   `json:"status"`
	Error  string   `json:"error,omitempty"`
	Data   []string `json:"data,omitempty"`
}

func handle(conn net.Conn, appStore *store.AppStore) {
	defer conn.Close()

	var req Request
	if err := json.NewDecoder(conn).Decode(&req); err != nil {
		json.NewEncoder(conn).Encode(Response{Status: "error", Error: err.Error()})
		return
	}

	switch req.Action {
	case "list":
		names := appStore.Names()
		json.NewEncoder(conn).Encode(Response{Status: "ok", Data: names})
	case "run":
		err := appStore.Run(req.Name)
		if err != nil {
			json.NewEncoder(conn).Encode(Response{Status: "error", Error: err.Error()})
		} else {
			json.NewEncoder(conn).Encode(Response{Status: "ok"})
		}
	default:
		json.NewEncoder(conn).Encode(Response{Status: "error", Error: "unknown action"})
	}

}

func main() {
	socketPath := "/tmp/mochi.sock"
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Auto cleanup on SIGINT/SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		listener.Close()
		os.Remove(socketPath)
		os.Exit(0)
	}()

	appStore := store.New()
	fmt.Println("Daemon running...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn, appStore)
	}

}
