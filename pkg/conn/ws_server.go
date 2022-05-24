package conn

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WsServerOptions struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type WsServer struct {
	options  *WsServerOptions
	upgrader websocket.Upgrader
	handler  ConnectionHandler
}

// NewWsServer options can be nil, use default value when nil.
func NewWsServer(options *WsServerOptions) Server {

	if options == nil {
		options = &WsServerOptions{
			ReadTimeout:  8 * time.Minute,
			WriteTimeout: 8 * time.Minute,
		}
	}
	ws := new(WsServer)
	ws.options = options
	ws.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 65536,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return ws
}

func (ws *WsServer) handleWebSocketRequest(writer http.ResponseWriter, request *http.Request) {

	conn, err := ws.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		// logger.E("upgrade http to ws error", err)
		return
	}

	proxy := ConnectionProxy{
		conn: NewWsConnection(conn, ws.options),
	}
	ws.handler(proxy)
}

func (ws *WsServer) SetConnHandler(handler ConnectionHandler) {
	ws.handler = handler
}

func (ws *WsServer) Run(host string, port int) error {

	http.HandleFunc("/ws", ws.handleWebSocketRequest)

	addr := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
