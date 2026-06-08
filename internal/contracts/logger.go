package contracts

type Logger interface {
	Info(msg string)
	Infof(msg string, args ...any)
	Error(msg string, err error)
	Errorf(msg string, err error, args ...any)
	Debug(msg string)
	Debugf(msg string, args ...any)
}
