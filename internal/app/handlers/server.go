package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aristat/http-go-kit/internal/app/api"

	"github.com/go-kit/kit/log"
)

type ServerHandler struct {
	logger log.Logger
}

func NewServerHandler(logger log.Logger) *ServerHandler {
	return &ServerHandler{logger: logger}
}

func (s *ServerHandler) GetProduct(w http.ResponseWriter, r *http.Request, id string) {
	intId, _ := strconv.ParseInt(id, 10, 64)

	product := api.Product{Id: intId, Name: "test"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}
