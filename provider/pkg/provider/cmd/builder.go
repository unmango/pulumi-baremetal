package cmd

import pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"

type B struct {
	bin   *pb.Bin
	stdin *string
	Args  []string
}

func (b *B) Cmd() *pb.Command {
	return &pb.Command{
		Bin:   *b.bin,
		Args:  b.Args,
		Stdin: b.stdin,
	}
}

func (b *B) Arg(value string) {
	if value != "" {
		b.Args = append(b.Args, value)
	}
}

func (b *B) Bin(bin pb.Bin) {
	b.bin = &bin
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

func (b *B) Stdin(value string) {
	b.stdin = &value
}

func (b *B) add(parts ...string) {
	if len(parts) > 2 {
		panic("don't pass more than 2 parts this function can't handle it")
	}

	// Build backwards so the original args come last
	b.Args = append(parts, b.Args...)
}
