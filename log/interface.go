package log

type Level int8

const (
	Debug Level = iota - 1
	Info
	Warn
	Error
	Fatal
)

type ILogger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

var (
	levelFlags = []string{"DEBG", "INFO", "WARN", "ERRO", "FATL"}
)

func LevelString(lv Level) string {
	return levelFlags[lv+1]
}
