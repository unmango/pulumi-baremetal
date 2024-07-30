package opts

type O[T any] interface {
	~func(*T) error
}

func Apply[V O[T], T any](x *T, opts ...V) error {
	for _, opt := range opts {
		if err := opt(x); err != nil {
			return err
		}
	}

	return nil
}
