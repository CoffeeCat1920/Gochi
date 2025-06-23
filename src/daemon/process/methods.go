package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func (process *Process) handler(conn net.Conn) (err error) {
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("close error: %w", closeErr)
			} else {
				fmt.Fprintf(os.Stderr, "warning: close error: %v\n", closeErr)
			}
		}
	}()

	var req Request
	encoder := json.NewEncoder(conn)
	if err = json.NewDecoder(conn).Decode(&req); err != nil {
		err = encoder.Encode(Response{Status: "error", Error: err.Error()})
		return
	}

	switch req.Verb {
	case "list":
		names := process.store.Names()
		err = encoder.Encode(Response{Status: "ok", Data: names})
	case "run":
		runErr := process.store.Run(req.Object)
		if runErr != nil {
			err = encoder.Encode(Response{Status: "error", Error: err.Error()})
		} else {
			err = encoder.Encode(Response{Status: "ok"})
		}
	default:
		err = encoder.Encode(Response{Status: "error", Error: fmt.Sprintf("unknown action, %s", req.Verb)})
		fmt.Print(req.Verb)
	}

	return
}

func (process *Process) Run() error {
	if err := os.RemoveAll(process.connStr); err != nil {
		return fmt.Errorf("could not remove existing socket: %w", err)
	}

	signal.Notify(process.quitChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-process.quitChan
		process.Close()
		os.Exit(0)
	}()

	fmt.Println("Daemon running on", process.connStr)

	for {
		conn, err := process.listner.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			fmt.Fprintln(os.Stderr, "Accept error:", err)
			continue
		}
		go process.handler(conn)
	}

	return nil
}

func (process *Process) Close() {
	fmt.Println("Shutting down daemon...")
	if process.listner != nil {
		process.listner.Close()
	}
	os.Remove(process.connStr)
}
