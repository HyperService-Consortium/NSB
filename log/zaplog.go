package log

import (
	"fmt"
	"io"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	kitzaplog "github.com/go-kit/kit/log/zap"
	colorful "github.com/gookit/color"
	tendermintlog "github.com/tendermint/tendermint/libs/log"
	zap "go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

type TendermintLogger = tendermintlog.Logger

type splitingIO struct {
	IO   io.Writer
	next *splitingIO
}

func (sio *splitingIO) Write(p []byte) (int, error) {
	n, err := sio.IO.Write(p)
	if err != nil {
		return n, err
	}
	if sio.next != nil {
		nn, err := sio.IO.Write(p)
		if err != nil {
			return nn, err
		}
		if n < nn {
			return n, err
		} else {
			return nn, err
		}
	} else {
		return n, err
	}
}

const (
	msgKey    = "_msg" // "_" prefixed to avoid collisions
	moduleKey = "module"
)

// Info logs a message at level Info.
func (l *zapLogger) Info(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Info(l.srcLogger)
	if err := kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...); err != nil {
		errLogger := kitlevel.Error(l.srcLogger)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

type zapLogger struct {
	srcLogger kitlog.Logger
}

// Debug logs a message at level Debug.
func (l *zapLogger) Debug(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Debug(l.srcLogger)
	if err := kitlog.With(lWithLevel, msgKey, msg).Log(keyvals...); err != nil {
		errLogger := kitlevel.Error(l.srcLogger)
		kitlog.With(errLogger, msgKey, msg).Log("err", err)
	}
}

// Error logs a message at level Error.
func (l *zapLogger) Error(msg string, keyvals ...interface{}) {
	lWithLevel := kitlevel.Error(l.srcLogger)
	lWithMsg := kitlog.With(lWithLevel, msgKey, msg)
	if err := lWithMsg.Log(keyvals...); err != nil {
		lWithMsg.Log("err", err)
	}
}

// With returns a new contextual logger with keyvals prepended to those passed
// to calls to Info, Debug or Error.
func (l *zapLogger) With(keyvals ...interface{}) tendermintlog.Logger {
	return &zapLogger{kitlog.With(l.srcLogger, keyvals...)}
}

const (
	// 	lightSkyBlue      = "87CEFA"
	lightlightSkyBlue = "B0E2FF"
	skyBlue           = "7EC0EE"
	pathColorPrefix   = "\x1b[38;5;230m "
	pathColorSuffix   = " \x1b[0m"
)

var (
	// colorLightSkyBlue      = colorful.HEX(lightSkyBlue)
	colorLightLightSkyBlue = colorful.HEX(lightlightSkyBlue)
	colorSkyBlue           = colorful.HEX(skyBlue)
	colorInfo              = colorLightLightSkyBlue.Sprintf("Info")
	colorDebug             = colorSkyBlue.Sprintf("Debug")
	colorWarn              = colorful.Yellow.Sprintf("Warn")
	colorPanic             = colorful.Red.Sprintf("Panic")
	colorError             = colorful.Red.Sprintf("Error")
	colorFatal             = colorful.Red.Sprintf("Fatal")
	colorDPanic            = colorful.Red.Sprintf("DPanic")
)

func zapColorfulLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString(colorDebug)
	case zapcore.InfoLevel:
		enc.AppendString(colorInfo)
	case zapcore.WarnLevel:
		enc.AppendString(colorWarn)
	case zapcore.ErrorLevel:
		enc.AppendString(colorError)
	case zapcore.DPanicLevel:
		enc.AppendString(colorDPanic)
	case zapcore.PanicLevel:
		enc.AppendString(colorPanic)
	case zapcore.FatalLevel:
		enc.AppendString(colorFatal)
	default:
		enc.AppendString(fmt.Sprintf("LEVEL(%d)", level))
	}
}

// ShortColorfulCallerEncoder serializes a caller in package/file:line format, trimming
// all but the final directory from the full path.
func ShortColorfulCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	if !caller.Defined {
		enc.AppendString("undefined")
		return
	}
	// nb. To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	//
	// Find the last separator.
	//
	idx := strings.LastIndexByte(caller.File, '/')
	if idx == -1 {
		enc.AppendString(caller.FullPath())
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(caller.File[:idx], '/')
	if idx == -1 {
		enc.AppendString(caller.FullPath())
	}
	buf := bufferpool.Get()
	buf.AppendString(pathColorPrefix)
	// Keep everything after the penultimate separator.
	buf.AppendString(caller.File[idx+1:])
	buf.AppendByte(':')
	buf.AppendInt(int64(caller.Line))
	buf.AppendString(pathColorSuffix)
	enc.AppendString(buf.String())
	buf.Free()
}

// FullColorfulCallerEncoder serializes a caller in package/file:line format, trimming
// all but the final directory from the full path.
func FullColorfulCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	if !caller.Defined {
		enc.AppendString("undefined")
		return
	}
	// nb. To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	//
	// Find the last separator.
	//
	buf := bufferpool.Get()
	buf.AppendString(pathColorPrefix)
	// Keep everything after the penultimate separator.
	buf.AppendString(caller.File)
	buf.AppendByte(':')
	buf.AppendInt(int64(caller.Line))
	buf.AppendString(pathColorSuffix)
	enc.AppendString(buf.String())
	buf.Free()
}

// // ColorfulISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
// // with millisecond precision.
// func ColorfulISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	buf := bufferpool.Get()
// 	buf.AppendString("\x1b[38;5;195m")
// 	buf.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
// 	buf.AppendString("\x1b[0m")
// 	enc.AppendString(buf.String())
// 	buf.Free()
// }

func NewZapColorfulDevelopmentSugarLogger(options ...zap.Option) (tendermintlog.Logger, error) {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapColorfulLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   ShortColorfulCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(append(options, zap.AddCallerSkip(1))...)

	// zap.NewDevelopment(options...)
	if err != nil {
		return nil, err
	}
	return &zapLogger{kitzaplog.NewZapSugarLogger(logger, zapcore.DebugLevel)}, nil
}
