package handler

import (
	"fmt"
	"net/http"

	"1337b04rd/internal/domain"
)

func (h *handler) CreatePostPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "create-post.html", nil)
}

func (h *handler) SubmitPost(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("web/templates/create-post.html")
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	http.Error(w, "Error parsing template", http.StatusInternalServerError)
	// 	return
	// }
	// Парсим multipart форму
	// err = r.ParseMultipartForm(10 << 20) // 10 MB
	// if err != nil {
	// 	http.Error(w, "Unable to parse form", http.StatusBadRequest)
	// 	return
	// }

	form := &domain.InputPost{
		Name:    r.FormValue("name"),
		Subject: r.FormValue("subject"),
		Content: r.FormValue("comment"),
	}
	file, header, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	form.File = file
	con := header.Header.Get("Content-Type")
	fmt.Println(form)
	fmt.Println(con)
}
