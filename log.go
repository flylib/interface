package connector

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type ILogger interface {
	Debug(str string, args ...any)
	Info(str string, args ...any)
	Warn(str string, args ...any)
	Error(str string, args ...any)
	Fatal(str string, args ...any)

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
