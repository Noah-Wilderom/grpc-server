package models

const (
	Info  LogType = "info"
	Error LogType = "error"
)

type LogType string

type Log struct {
	Id   string  `bson:"_id,omitempty" json:"id"`
	Type LogType `bson:"type" json:"type"`
	Data string  `bson:"data" json:"data"`
}
