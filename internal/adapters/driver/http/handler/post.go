package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"1337b04rd/internal/domain"
)

func (h *handler) ActivePost(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseUint(r.PathValue("postID"), 10, 64)
	if err != nil {
		slog.Error(err.Error())

		errData := &domain.ErrorPageData{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}

		h.renderError(w, errData)
		return
	}

	post, err := h.use.GetActivePost(r.Context(), postID)
	if err != nil {
		slog.Error(err.Error())

		errData := &domain.ErrorPageData{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

		h.renderError(w, errData)
		return
	}
	gg(post)

	err = h.templates.ExecuteTemplate(w, "post.html", post)
	if err != nil {
		slog.Error(err.Error())
		errData := &domain.ErrorPageData{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		h.renderError(w, errData)
	}
}

func gg(c *domain.ActivePost) {
	file, err := os.Create("comment.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode struct to JSON and write directly to file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: pretty-print
	if err := encoder.Encode(c); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
