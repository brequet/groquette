package v1

import (
	"github.com/brequet/groquette/internal/usecase"
	"github.com/go-chi/chi"
)

func NewRouter(handler *chi.Mux, chatUseCase *usecase.ChatUseCase) {
	handler.Route("/v1", func(r chi.Router) {
		newChatRoutes(r, chatUseCase)
	})
}
