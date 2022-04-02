package glog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"teacupapi/config"
)

var (
	zapLog *zap.SugaredLogger // 简易版日志文件
	// Logger *zap.Logger // 这个日志强大一些, 目前还用不到

	logLevel = zap.NewAtomicLevel()
)

// InitLog 初始化日志文件
func InitLog() error {
	logConf := config.GetLogConf()

	loglevel := zapcore.InfoLevel
	switch logConf.LogLevel {
	case "INFO":
		loglevel = zapcore.InfoLevel
	case "ERROR":
		loglevel = zapcore.ErrorLevel
	}
	setLevel(loglevel)

	var core zapcore.Core
	// 打印至文件中
	if logConf.LogType == "file" {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConf.LogPath,
			MaxSize:    128, // MB
			LocalTime:  true,
			Compress:   true,
			MaxBackups: 8, // 最多保留 n 个备份
		})

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			logLevel,
		)
	} else {
		// 打印在控制台
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), logLevel)
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog = logger.Sugar()

	return nil
}

func setLevel(level zapcore.Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

func Info(args ...interface{}) {
	zapLog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLog.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zapLog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLog.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zapLog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLog.Errorf(template, args...)
}
