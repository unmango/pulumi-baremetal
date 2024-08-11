package cmd

type State[T Builder] struct {
	CommandArgs[T]

	ExitCode     int               `pulumi:"exitCode"`
	Stderr       string            `pulumi:"stderr"`
	Stdout       string            `pulumi:"stdout"`
	CreatedFiles []string          `pulumi:"createdFiles"`
	MovedFiles   map[string]string `pulumi:"movedFiles"`
}

func (s *State[T]) Copy() State[T] {
	return State[T]{
		CommandArgs:  s.CommandArgs,
		ExitCode:     s.ExitCode,
		Stderr:       s.Stderr,
		Stdout:       s.Stdout,
		CreatedFiles: s.CreatedFiles,
		MovedFiles:   s.MovedFiles,
	}
}
