/*
 * @Notice: edit notice here
 * @Author: zhulei
 * @Date: 2022-11-14 11:36:25
 * @LastEditors: zhulei
 * @LastEditTime: 2022-12-26 15:47:54
 */
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (

	DEBUG = zapcore.DebugLevel
	
	INFO = zapcore.InfoLevel
	
	WARN = zapcore.WarnLevel
	
	ERROR = zapcore.ErrorLevel
	
	DPANIC = zapcore.DPanicLevel
	
	PANIC = zapcore.PanicLevel
	
	FATAL = zapcore.FatalLevel

)

var SugarLogger *zap.SugaredLogger

// InitLogger 初始化日志
func InitLogger(logFile string, logLevel zapcore.LevelEnabler, jsonMode bool) {

	writeSyncer, err := getLogWriter(logFile)
	if err != nil {
		fmt.Printf("create or open or rotate log file error: %s", err)
		os.Exit(1)
	}

	encoder := getEncoder(jsonMode)
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	// core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) 
	// 日志级别: 

	// 将调用函数信息记录到日志中, AddCallerSkip增加了调用者注释跳过的调用者数量，日志中显示的是调用封装函数的位置
	// AddStacktrace输出此次调用的堆栈,配置 zapcore.WarnLevel 则 Warn()/Error() 等级别的日志会输出堆栈
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	SugarLogger = logger.Sugar()
}

// getEncoder 配置日志编码
func getEncoder(jsonMode bool) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写字母日志级别
	if !jsonMode {
		// 普通格式日志
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	// json 格式日志
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 创建日志文件
func getLogWriter(logFile string) (zapcore.WriteSyncer, error) {

	os.MkdirAll(filepath.Dir(logFile), os.ModePerm)
	// file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file, err := getWriter(logFile)
	return zapcore.AddSync(file), err
}

// getWriter 日志文件切割, 保存日志30天，每小时分割一次日志
func getWriter(filename string) (io.Writer, error) {

	hook, err := rotatelogs.New(
		filename+"_%Y%m%d%H.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*1),
	)

	//保存日志30天，每1分钟分割一次日志, 测试使用
	// hook, err := rotatelogs.New(
	// 	filename+"_%Y%m%d%H%M.log",
	// 	rotatelogs.WithLinkName(filename),
	// 	rotatelogs.WithMaxAge(time.Hour*24*30),
	// 	rotatelogs.WithRotationTime(time.Minute*1),
	// )

	// if err != nil {
	// 	fmt.Printf("rotate log file error: %s", err)
	// }

	return hook, err

}
