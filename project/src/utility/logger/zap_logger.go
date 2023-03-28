package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	day := time.Now()
	logFile := &lumberjack.Logger{
		Filename:   fmt.Sprintf("/go/src/log/app_%s.log", day.Format("2006-01-02")),
		MaxSize:    100, // MB
		MaxBackups: 5,   // 保存するバックアップ数
		MaxAge:     28,  // 日数
		LocalTime:  true,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)),
		zap.DebugLevel,
	)

	l := zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())
	s := l.Sugar()
	return &ZapLogger{logger: l, sugaredLogger: s}
}

func (l *ZapLogger) Debug(msg string) {
	msg = fmt.Sprintf("[DEBUG] %s", msg)
	l.logger.Debug(msg)
}

func (l *ZapLogger) Info(msg string) {
	msg = fmt.Sprintf("[INFO] %s", msg)
	l.logger.Info(msg)
}

func (l *ZapLogger) Warn(msg string) {
	msg = fmt.Sprintf("[WARN] %s", msg)
	l.logger.Warn(msg)
}

func (l *ZapLogger) Error(msg string) {
	msg = fmt.Sprintf("[ERROR] %s", msg)
	l.logger.Error(msg)
}

func (l *ZapLogger) Fatal(msg string) {
	msg = fmt.Sprintf("[FATAL] %s", msg)
	l.logger.Fatal(msg)
}

func (l *ZapLogger) Debugf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[DEBUG] %s", msg)
	l.logger.Debug(msg)
}

func (l *ZapLogger) Infof(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[INFO] %s", msg)
	l.logger.Info(msg)
}

func (l *ZapLogger) Warnf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[WARN] %s", msg)
	l.logger.Warn(msg)
}

func (l *ZapLogger) Errorf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[ERROR] %s", msg)
	l.logger.Error(msg)
}

func (l *ZapLogger) Fatalf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[FATAL] %s", msg)
	l.logger.Fatal(msg)
}
