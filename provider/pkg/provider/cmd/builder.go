package cmd

type B struct {
	Args []string
}

func (b *B) Arg(value string) {
	if value != "" {
		b.Args = append(b.Args, value)
	}
}

func (b *B) Op(input bool, name string) {
	if input {
		b.add(name)
	}
}

func (b *B) Opv(value, name string) {
	if value != "" {
		b.add(name, value)
	}
}

func (b *B) add(parts ...string) {
	if len(parts) > 2 {
		panic("don't pass more than 2 parts this function can't handle it")
	}

	// Build backwards so the original args come last
	b.Args = append(parts, b.Args...)
}
