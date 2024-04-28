package v1

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/brequet/groquette/internal/entity"
	"github.com/brequet/groquette/internal/usecase"
	"github.com/go-chi/chi"
)

type chatRoutes struct {
	c *usecase.ChatUseCase
}

func newChatRoutes(handler chi.Router, c *usecase.ChatUseCase) {
	cr := &chatRoutes{c}

	handler.Route("/chat", func(r chi.Router) {
		r.Get("/", cr.hello)
		r.Post("/", cr.sendMessage)
	})
}

func (router *chatRoutes) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello !"))
}

type sendMessageRequest struct {
	Content   string `json:"content"`
	ChatModel string `json:"chatModel"`
}

type sendMessageResponse struct {
	Response string `json:"response"`
}

// TODO logging
func (router *chatRoutes) sendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req sendMessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("sending parsing quest body: %s", err)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if !entity.ValidateChatModel(req.ChatModel) {
		slog.Error("wrong chat model '%s': %s", req.ChatModel, err)
		http.Error(w, fmt.Sprintf("chat model wrong: '%s'", req.ChatModel), http.StatusBadRequest)
	}
	messageReponse, err := router.c.SendMessage(req.Content, entity.ChatModel(req.ChatModel))
	if err != nil {
		slog.Error("sending message: %s", err)
		http.Error(w, "Error sending message", http.StatusInternalServerError)
		return
	}

	response := sendMessageResponse{
		Response: messageReponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
