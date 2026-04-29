package logger

import "log/slog"

// Error ...
func Error(err error) slog.Attr {
	return slog.Any("error", err)
}

// String ...
func String(key, msg string) slog.Attr {
	return slog.String(key, msg)
}

// Any ...
func Any(key string, data any) slog.Attr {
	return slog.Any(key, data)
}

// Int ...
func Int(key string, data int) slog.Attr {
	return slog.Int(key, data)
}
