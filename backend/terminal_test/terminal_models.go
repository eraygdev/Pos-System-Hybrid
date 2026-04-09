package terminal_test

type MenuOption struct {
	ID     int
	Label  string
	Action func() error
}
