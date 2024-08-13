package opts

import "log/slog"

type HasLogger interface {
	SetLogger(*slog.Logger)
}

func WithLogger[V O[T], T HasLogger](logger *slog.Logger) V {
	return Safe[V](func(t *T) {
		(*t).SetLogger(logger)
	})
}
