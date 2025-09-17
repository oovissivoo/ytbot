package link

import (
	"encoding/json"
	"net/http"
	"ytbot/pkg/req"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error, http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdLink)
	}
}
