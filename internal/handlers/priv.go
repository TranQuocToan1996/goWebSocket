package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

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
