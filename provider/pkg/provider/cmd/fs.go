package cmd

type FsManipulator interface {
	ExpectCreated() []string
	ExpectMoved() map[string]string
}
