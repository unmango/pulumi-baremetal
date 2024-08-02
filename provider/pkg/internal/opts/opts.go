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

func NoOp[V O[T], T any]() V {
	return func(*T) error {
		return nil
	}
}

func If[V O[T], T any](predicate bool, o V) V {
	if predicate {
		return o
	} else {
		return NoOp[V]()
	}
}
