package handlers

import (
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
)

// renderPage renders template with jet lib
func renderPage(w http.ResponseWriter, template string, dataMapToTemplate jet.VarMap) error {
	// Find parsed template first
	// If can't try to load existing temp
	// if cant load, try find inside path
	// At the end, if cant, return err
	view, err := views.GetTemplate(template)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Executive temp
	err = view.Execute(w, dataMapToTemplate, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Sent to all users when a user sending message
func broadcastToAll(res WsJsonResp) {
	for client := range clients {
		err := client.WriteJSON(res)
		if err != nil {
			log.Println("ws err")
			err = client.Close()
			if err != nil {
				// try to close one more time
				defer client.Close()
			}

			// delete user from tracking map
			delete(clients, client)
		}
	}
}

func getUsers() []string {
	list := []string{}
	for _, client := range clients {
		if client != "" {
			list = append(list, client)
		}
	}
	sort.Strings(list)
	return list
}
