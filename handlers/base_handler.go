package handlers

import (
	"html/template"
	"net/http"
	"path"
)

type BaseHandler struct{}

func (h *BaseHandler) RenderText(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func (h *BaseHandler) RenderHTML(w http.ResponseWriter, status int, templateFile string, data interface{}) {
	filepath := path.Join("templates", templateFile)
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		h.RenderText(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		h.RenderText(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(status)
}
