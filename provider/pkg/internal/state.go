package internal

import "log/slog"

type State struct {
	Log *slog.Logger
}

func (s State) WithLogger(logger *slog.Logger) State {
	s.Log = logger
	return s
}
