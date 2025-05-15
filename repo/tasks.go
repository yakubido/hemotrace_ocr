package repo

import (
	"github.com/yakubido/hemotrace_ocr/config"
	"go.uber.org/zap"
)

type Task struct {
}

type Tasks interface {
}

type tasks struct {
	logger *zap.Logger
}

func NewTask(logger *zap.Logger, cfg *config.Config) *tasks {
	t := tasks{
		logger: logger,
	}

	return &t
}
