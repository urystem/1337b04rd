package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"1337b04rd/internal/domain"
)

func (h *handler) CreatePostPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "create-post.html", nil)
}

func (h *handler) SubmitPost(w http.ResponseWriter, r *http.Request) {
	// Парсим multipart форму
	// err = r.ParseMultipartForm(10 << 20) // 10 MB
	// if err != nil {
	// 	http.Error(w, "Unable to parse form", http.StatusBadRequest)
	// 	return
	// }
	ctx := r.Context()
	sess, x := h.middleware.FromContext(ctx)
	if !x {
		fmt.Println(x)
		return
	}
	form := &domain.Form{
		// Name:    r.FormValue("name"),
		// Subject: r.FormValue("subject"),
		// Content: r.FormValue("comment"),
	}
	form.Uuid = sess.Uuid
	form.Name = r.FormValue("name")
	form.Subject = r.FormValue("subject")
	form.Content = r.FormValue("content")

	file, header, err := r.FormFile("file")

	if err == nil {
		form.File = new(domain.InPutObject)
		form.File.ReadCloser = file
		form.File.ConType = header.Header.Get("Content-Type")
		form.File.Size = header.Size
	}

	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "File error", http.StatusInternalServerError)
		slog.Error(err.Error())
		errData := &domain.ErrorPageData{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

		h.renderError(w, errData)
		return
	}

	err = h.use.CreatePost(ctx, form)
	if err != nil {
		errData := &domain.ErrorPageData{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		h.renderError(w, errData)
	}
	fmt.Println(form)
}
