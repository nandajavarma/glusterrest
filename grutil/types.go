package grutil

import (
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type EventMsg struct {
	Node    string `json:"node"`
	Message string `json:"message"`
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type Lockfile struct {
	path string
	fh   *os.File
	sync.Mutex
}

type WsClients struct {
	clients []*websocket.Conn
	sync.Mutex
}

type Resp struct {
	Ok         bool   `json:"ok"`
	Msg        string `json:"msg"`
	ReturnCode string `json:"return_code"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Config struct {
	Port  int  `json:"port"`
	Https bool `json:"https"`
}
