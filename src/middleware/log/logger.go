package log

type Logger interface {
	Tracef(format string, params ...interface{})

	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{}) error

	Errorf(format string, params ...interface{}) error

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{}) error

	Error(v ...interface{}) error
}

type RPLogger struct {
	logger Logger
}

func (rpLogger *RPLogger) Tracef(format string, params ...interface{}) {

}

func (rpLogger *RPLogger) Debugf(format string, params ...interface{}) {

}
func (rpLogger *RPLogger) Infof(format string, params ...interface{}) {

}
func (rpLogger *RPLogger) Warnf(format string, params ...interface{}) error {
	return rpLogger.logger.Warnf(format, params)
}
func (rpLogger *RPLogger) Errorf(format string, params ...interface{}) error {
	return rpLogger.logger.Errorf(format, params)
}

func (rpLogger *RPLogger) Debug(v ...interface{}) {

}
func (rpLogger *RPLogger) Info(v ...interface{}) {

}

func (rpLogger *RPLogger) Warn(v ...interface{}) error {
	return rpLogger.logger.Warn(v)
}

func (rpLogger *RPLogger) Error(v ...interface{}) error {
	return rpLogger.logger.Error(v)
}
