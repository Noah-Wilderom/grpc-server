package models

import "context"

type LogModel interface {
	GetAll(context.Context) ([]*Log, error)
	GetById(context.Context, string) (*Log, error)
	Insert(context.Context, *Log) error
	Delete(context.Context, *Log) bool
}
