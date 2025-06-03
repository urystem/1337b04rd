package handler

import (
	"log/slog"
	"net/http"
	"text/template"
)

func (h *handler) Catalog(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Usecase.ListOfPosts()
	if err != nil {
		// http.Error(w, "Error parsing template", http.StatusInternalServerError)
		slog.Error("dsfd")
		return
	}

	tmpl, err := template.ParseFiles("templates/catalog.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		slog.Error("exe")
	}
}
