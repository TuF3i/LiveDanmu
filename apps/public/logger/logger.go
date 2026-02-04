package logger

import (
	"errors"

	"go.uber.org/zap"
)

type NewLogger struct {
	Logger   *zap.Logger
	LokiHook *lokiHook
}

func GetLogger(v interface{}) (*NewLogger, error) {
	l := NewLogger{}
	conf, ok := v.(*LokiConfig)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	l.Logger, l.LokiHook = initZapWithLoki(*conf)
	return &l, nil
}
