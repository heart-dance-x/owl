package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var logLevel = zap.NewAtomicLevel()

func init() {
	//filePath := getFilePath()

	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:  filePath,
	//	MaxSize:   1024, //MB
	//	LocalTime: true,
	//	Compress:  true,
	//})

	//config := zap.NewProductionEncoderConfig()
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(config),
		zapcore.NewConsoleEncoder(config),
		os.Stdout,
		logLevel,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	//log = logger.Sugar()
}

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel
)

func SetLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

//func getCurrentDirectory() string {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	if err != nil {
//		log.Info(err)
//	}
//	return strings.Replace(dir, "\\", "/", -1)
//}
//
//func getFilePath() string {
//	logfile := getCurrentDirectory() + "/" + getAppname() + ".log"
//	return logfile
//}

//func getAppname() string {
//	full := os.Args[0]
//	full = strings.Replace(full, "\\", "/", -1)
//	splits := strings.Split(full, "/")
//	if len(splits) >= 1 {
//		name := splits[len(splits)-1]
//		name = strings.TrimSuffix(name, ".exe")
//		return name
//	}
//
//	return ""
//}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

//
//func Infof(template string, args ...interface{}) {
//	log.Infof(template, args...)
//}
//
//func Warn(args ...interface{}) {
//	log.Warn(args...)
//}
//
//func Warnf(template string, args ...interface{}) {
//	log.Warnf(template, args...)
//}
//
//func Error(args ... interface{}) {
//	log.Error(args...)
//}
//
//func Errorf(template string, args ...interface{}) {
//	log.Errorf(template, args...)
//}
//
//func Panic(args ...interface{}) {
//	log.Panic(args...)
//}
//
//func Panicf(template string, args ...interface{}) {
//	log.Panicf(template, args...)
//}
