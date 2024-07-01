package logger

import (
	"io"
	"log/slog"
	"os"
)

var Logger *logger

type logger struct {
	infof  io.Writer
	debugf io.Writer
	errf   io.Writer
	warnf  io.Writer
	log    *slog.Logger
}

func InitLogger(infoFileName string, debugFileName string, errFileName string, warnFileName string) error {
	infoFile, err := os.OpenFile(infoFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	debugFile, err := os.OpenFile(debugFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	errFile, err := os.OpenFile(errFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	warnFile, err := os.OpenFile(warnFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logger := &logger{
		infof:  infoFile,
		debugf: debugFile,
		errf:   errFile,
		warnf:  warnFile,
	}
	Logger = logger
	return nil
}

func (l *logger) Info(msg string, args ...any) {
	l.log = slog.New(slog.NewTextHandler(io.MultiWriter(l.infof, os.Stdout), &slog.HandlerOptions{}))
	l.log.Info(msg, args...)
}

func (l *logger) Debug(msg string, args ...any) {
	l.log = slog.New(slog.NewTextHandler(io.MultiWriter(l.debugf, os.Stdout), &slog.HandlerOptions{}))
	l.log.Debug(msg, args...)
}
func (l *logger) Error(msg string, args ...any) {
	l.log = slog.New(slog.NewTextHandler(io.MultiWriter(l.errf, os.Stdout), &slog.HandlerOptions{}))
	l.log.Error(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log = slog.New(slog.NewTextHandler(io.MultiWriter(l.warnf, os.Stdout), &slog.HandlerOptions{}))
	l.log.Warn(msg, args...)
}
