package handlers

import "github.com/Noah-Wilderom/grpc-server/logger-service/database/models"

type LogHandler struct {
	model models.LogModel
}

func NewLogHandler(m models.LogModel) *LogHandler {
	return &LogHandler{
		model: m,
	}
}
