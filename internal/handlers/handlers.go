package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

const (
	routeHTML = "./html"
	routeHome = "home.jet"
)

var (
	views = jet.NewSet(
		jet.NewOSFileSystemLoader(routeHTML), // Load template files from path
		jet.InDevelopmentMode(),              // Bypass cache in dev mode
	)

	upgradeConn = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WsJsonResp struct {
	Action      string `json:"action"`
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
	Code        string `json:"code"`
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

	err = ws.WriteJSON(resp)
	if err != nil {
		log.Println(err.Error())
	}
}
