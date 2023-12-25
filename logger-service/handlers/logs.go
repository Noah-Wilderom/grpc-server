package handlers

import "github.com/Noah-Wilderom/grpc-server/logger-service/database/models"

type LogHandler struct {
	model models.Log
}

func NewLogHandler(m models.Log) *LogHandler {
	return &LogHandler{
		model: m,
	}
}
