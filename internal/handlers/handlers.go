package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

const (
	routeHTML = "./html"
	routeHome = "home.html"
)

var (
	// Channel work with payload that send into websocket
	wsChan = make(chan WsPayload, 1000)

	// Keep track on who online on the chatbox
	clients = make(map[WebSocketConnection]string)

	views = jet.NewSet(
		jet.NewOSFileSystemLoader(routeHTML), // Load template files from path
		jet.InDevelopmentMode(),              // Bypass cache in dev mode
	)

	upgradeConn = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
	}
)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsJsonResp struct {
	Action         string   `json:"action"`
	MessageType    string   `json:"message_type"`
	Message        string   `json:"message"`
	ConnectedUsers []string `json:"connected_users"`
	Code           string   `json:"code"`
}

type WsPayload struct {
	Username string              `json:"username"`
	Action   string              `json:"action"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, routeHome, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConn.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("ws upgraded")
	resp := WsJsonResp{
		Message: `<em><small>Connected to server</small></em>`,
	}

	// Add new client to map tracking
	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(resp)
	if err != nil {
		log.Println(err.Error())
	}

	go ListenForWsConn(&conn)
}
