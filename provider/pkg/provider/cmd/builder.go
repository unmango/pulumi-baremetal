package cmd

type builder struct {
	args []string
}

func (b *builder) arg(value string) {
	b.args = append(b.args, value)
}

func (b *builder) op(input bool, name string) {
	if input {
		b.add(name)
	}
}

func (b *builder) opv(value, name string) {
	if value != "" {
		b.add(name, value)
	}
}

func (b *builder) add(parts ...string) {
	if len(parts) > 2 {
		panic("don't pass more than 2 parts this function can't handle it")
	}

	// Build backwards so the original args come last
	b.args = append(parts, b.args...)
}
