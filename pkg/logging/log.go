package logging

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gin-swagger-demo/pkg/setting"
)

var Logger *zap.Logger

func Setup() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	logLevel := zapcore.ErrorLevel
	if setting.ServerSetting.RunMode == "debug"{
		logLevel = zapcore.DebugLevel
	}

	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	Logger = zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))
	Logger.Info( "zap.Logger inited")
}

// Debug output logs at debug level
func Debug(msg string, params ...interface{}) {
	Logger.Debug(msg,formatParams(params...)...)
}

// Info output logs at info level
func Info(msg string, params ...interface{}) {
	Logger.Info(msg,formatParams(params...)...)
}

// Warn output logs at warn level
func Warn(msg string, params ...interface{}) {
	Logger.Warn(msg,formatParams(params...)...)
}

// Error output logs at error level
func Error(msg string, params ...interface{}) {
	fields := formatParams(params...)
	Logger.Error(msg,fields...)
}

// Fatal output logs at fatal level
func Fatal(msg string, params ...interface{}) {
	Logger.Fatal(msg,formatParams(params...)...)
}

func formatParams(params ...interface{}) []zapcore.Field {
	total := len(params)

	//把参数填成偶数个
	if total%2 == 1 {
		params = append(params,"")
		total +=1
	}

	var fields []zapcore.Field
	for i :=0; i < total;i+=2 {
		field := zap.Any(fmt.Sprintf("%+v", params[i]),params[i+1])
		fields = append(fields,field)
	}
	return fields
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

//func getLogWriter() zapcore.WriteSyncer {
//	var err error
//	filePath := getLogFilePath()
//	fileName := getLogFileName()
//	F, err := file.MustOpen(fileName, filePath)
//	if err != nil {
//		log.Fatalf("logging.Setup err: %v", err)
//	}
//
//	return zapcore.AddSync(F)
//}

func getLogWriter() zapcore.WriteSyncer {
	filePath := getLogFilePath()

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./" + filePath + "latest.log",  // 日志输出文件
		MaxSize:    20,  // 日志最大保存20M
		MaxBackups: 50,  // 旧日志保留50个日志备份
		MaxAge:     50,  // 最多保留50天日志
		Compress:   false, // 自导打 gzip包 默认false
	}
	return zapcore.AddSync(lumberJackLogger)
}