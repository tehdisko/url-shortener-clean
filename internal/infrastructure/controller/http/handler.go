package http

import (
	"encoding/json"
	"io"
	"net/http"
	"url-shortener-clean/internal/entity"
)

type shortener interface {
	Shorten(url string) (entity.URL, error)
}

type expander interface {
	Expand(id string) (entity.URL, error)
}

type Handler struct {
	shortenUseCase shortener
	expandUseCase  expander
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Result string `json:"result"`
}

func NewHandler(s shortener, e expander) *Handler {
	return &Handler{shortenUseCase: s, expandUseCase: e}
}

func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req := shortenRequest{}
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := entity.NewURL("", req.URL)
	_, _ = h.shortenUseCase.Shorten(url)
	resp := shortenResponse{Result: "http://localhost:8080/" + url.ID}
	respBody, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func (h *Handler) ExpandHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	url, _ := h.expandUseCase.Expand(id)
	http.Redirect(w, r, url.OriginalURL, http.StatusTemporaryRedirect)
}
