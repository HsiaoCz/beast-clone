package logger

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"
)

var lg *zap.Logger

type logConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// InitLogger 初始化Logger
func InitLogger() (err error) {
	maxsize, err := strconv.Atoi(os.Getenv("LOG_MAXSIZE"))
	if err != nil {
		return err
	}
	maxBackups, err := strconv.Atoi(os.Getenv("LOG_MAXBACKUPS"))
	if err != nil {
		return err
	}
	maxAge, err := strconv.Atoi(os.Getenv("LOG_MAXAGE"))
	if err != nil {
		return err
	}
	cfg := logConfig{
		Level:      os.Getenv("LOG_LEVEL"),
		Filename:   os.Getenv("LOG_FILENAME"),
		MaxSize:    maxsize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
