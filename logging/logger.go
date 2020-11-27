package logging

func Debug(msg string) {
	Logger.Debug(msg)
}

func Debugf(template string, args ...interface{}) {
	SugarLogger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	SugarLogger.Debugw(msg, keysAndValues...)
}

func Info(msg string) {
	Logger.Info(msg)
}

func Infof(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	SugarLogger.Infow(msg, keysAndValues...)
}

func Warn(msg string) {
	Logger.Warn(msg)
}

func Warnf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	SugarLogger.Warnw(msg, keysAndValues...)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Errorf(template string, args ...interface{}) {
	SugarLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	SugarLogger.Errorw(msg, keysAndValues...)
}

func Panic(msg string) {
	Logger.Panic(msg)
}

func Panicf(template string, args ...interface{}) {
	SugarLogger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	SugarLogger.Panicw(msg, keysAndValues...)
}
