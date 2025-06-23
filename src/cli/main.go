package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Request struct {
	Verb   string `json:"verb"`
	Object string `json:"object,omitempty"`
}

type Response struct {
	Status string   `json:"status"`
	Error  string   `json:"error,omitempty"`
	Data   []string `json:"data,omitempty"`
}

func main() {
	conn, err := net.Dial("unix", "/tmp/mochi.sock")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := Request{Verb: "list"} // or Verb: "run", Object: "AppName"
	json.NewEncoder(conn).Encode(req)

	var resp Response
	if err := json.NewDecoder(conn).Decode(&resp); err != nil {
		panic(err)
	}

	fmt.Printf("Response: %+v\n", resp)
}
