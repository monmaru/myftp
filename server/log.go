package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newFileRotateLogger(path string) *zap.Logger {
	zapConf := zap.NewProductionEncoderConfig()
	zapConf.EncodeTime = zapcore.ISO8601TimeEncoder
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: path,
		MaxSize:  10,
		MaxAge:   7,
		Compress: false,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConf),
		w,
		zapcore.DebugLevel)
	return zap.New(core, zap.AddCaller())
}
