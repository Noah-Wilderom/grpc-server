package handlers

import (
	"context"
	"github.com/Noah-Wilderom/grpc-server/logger-service/database/models"
)

type LogHandler struct {
	model models.LogModel
}

func NewLogHandler(m models.LogModel) *LogHandler {
	return &LogHandler{
		model: m,
	}
}

func (h *LogHandler) SaveLog(l *models.Log) error {
	return h.model.Insert(context.Background(), l)
}
