package logger

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type Meta map[string]interface{}

// Logger is the interface for a simple but feature-complete logger.
type Logger interface {
	// Debug logs a debug level entry.
	Debug(arg interface{}, meta ...Meta)
	// Info logs an info level entry.
	Info(arg interface{}, meta ...Meta)
	// Error logs an error level entry.
	Error(arg interface{}, meta ...Meta)
}

// Level represents the logging level.
type Level int

const (
	// LevelDebug will log all log messages.
	LevelDebug Level = iota
	// LevelInfo will log operational messages such as requests.
	LevelInfo
	// LevelError will only log errors.
	LevelError
)

// Output represents the log output.
type Output int

const (
	// OutputStdout will log to the console.
	OutputStdout Output = iota
	// Outputfile will log to a specified file.
	OutputFile
)

type Format int

const (
	// FormatJSON will log in JSON format.
	FormatJSON Format = iota
	// FormatText
	FormatText
)

// Log is a standardised logger.
type Log struct {
	*logrus.Logger
}

// New creates a new logger.
func New(options ...func(l *Log) error) (*Log, error) {
	l := &Log{logrus.New()}

	for _, o := range options {
		if err := o(l); err != nil {
			return nil, err
		}
	}

	return l, nil
}

// LogLevel sets the logging level.
func LogLevel(lv Level) func(*Log) error {
	return func(l *Log) error {
		switch lv {
		case LevelDebug:
			l.Level = logrus.DebugLevel
		case LevelInfo:
			l.Level = logrus.InfoLevel
		case LevelError:
			l.Level = logrus.ErrorLevel
		default:
			return errors.New("Invalid log level")
		}

		return nil
	}
}

// LogOutput sets the log output.
func LogOutput(o Output, file io.Writer) func(*Log) error {
	return func(l *Log) error {
		switch o {
		case OutputStdout:
			l.Out = os.Stdout
		case OutputFile:
			if file == nil {
				return errors.New("You must specify a file to log to.")
			}
			l.Out = file
		default:
			return errors.New("Invalid output type")
		}

		return nil
	}
}

// LogFormat sets the log format.
func LogFormat(f Format) func(*Log) error {
	return func(l *Log) error {
		switch f {
		case FormatJSON:
			l.Formatter = &logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
			}
		case FormatText:
			l.Formatter = &logrus.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339,
			}
		default:
			return errors.New("Invalid log format")
		}

		return nil
	}
}

func (l *Log) Debug(arg interface{}, meta ...Meta) {
	if len(meta) > 0 {
		l.WithFields(logrus.Fields(meta[0])).Debug(arg)
		return
	}

	l.Logger.Debug(arg)
}

func (l *Log) Info(arg interface{}, meta ...Meta) {
	if len(meta) > 0 {
		l.WithFields(logrus.Fields(meta[0])).Info(arg)
		return
	}

	l.Logger.Info(arg)
}

func (l *Log) Error(arg interface{}, meta ...Meta) {
	if len(meta) > 0 {
		l.WithFields(logrus.Fields(meta[0])).Error(arg)
		return
	}

	l.Logger.Error(arg)
}
