package handlers

import (
	"fmt"
	"log"
)

// ListenForWsConn Opens a goroutine that contain channel listen forever for ws connection
func ListenForWsConn(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// Just no payload, do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
		// time.Sleep(time.Millisecond)
	}
}

func ListenToWsChan() {
	log.Println("starting listen for ws chan")
	var res WsJsonResp

	for {
		event := <-wsChan
		switch event.Action {
		case "username":
			// get list of users and broadcast
			clients[event.Conn] = event.Username
			users := getUsers()
			res.Action = "list_users"
			res.ConnectedUsers = users
			broadcastToAll(res)

		case "leave":
			res.Action = "list_users"
			delete(clients, event.Conn)
			users := getUsers()
			res.ConnectedUsers = users
			broadcastToAll(res)

		case "broadcast":
			res.Action = "broadcast"
			res.Message = fmt.Sprintf("<strong>%s</strong>: %s", event.Username, event.Message)
			broadcastToAll(res)
		}

	}
}
