package project

// Stack identifies which framework/CMS a project targets.
type Stack string

// Supported stacks.
const (
	StackSymfony   Stack = "symfony"
	StackWordPress Stack = "wordpress"
)

// String returns the human-readable label for s (e.g. "Symfony").
func (s Stack) String() string {
	switch s {
	case StackSymfony:
		return "Symfony"
	case StackWordPress:
		return "WordPress"
	default:
		return string(s)
	}
}
