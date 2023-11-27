package loges

// Логировать состояние
func (l *Core) Info(format string, v ...any) {
	if l.logInfo != nil {
		l.logInfo.Printf(format, v...)
	}
}

// Логгировать предупреждение
func (l *Core) Warn(format string, v ...any) {
	if l.logWarning != nil {
		l.logWarning.Printf(format, v...)
	}
}

// Логгировать ошибку
func (l *Core) Err(format string, v ...any) {
	if l.logError != nil {
		l.logError.Printf(format, v...)
	}
}
