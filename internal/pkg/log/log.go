package log

import (
	"io"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

type Field = zap.Field

// function variables for all field types
// in https://github.com/uber-go/zap/blob/master/field.go
var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any

	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
)

func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug
}

type Logger struct {
	instance *zap.Logger
	level    Level
}

var std = New(os.Stdout, InfoLevel)

func Default() *Logger {
	return std
}

func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	cfg := newLoggerConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	return &Logger{
		instance: zap.New(core),
		level:    level,
	}
}

func newLoggerConfig() zap.Config {
	environment := viper.GetString("server.environment")
	if environment == "production" {
		return zap.NewProductionConfig()
	} else {
		return zap.NewDevelopmentConfig()
	}
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.instance.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.instance.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.instance.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.instance.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.instance.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.instance.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.instance.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.instance.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}
