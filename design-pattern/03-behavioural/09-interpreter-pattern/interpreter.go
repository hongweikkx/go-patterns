package interpreter

type Context struct {
	Input  string
	Output string
}

type AbstractExpression interface {
	Interpret(context Context) string
}

type TerminalExpression struct{}

func (t *TerminalExpression) Interpret(context Context) string {
	return "terminal"
}

type NonTerminalEexpression struct{}

func (nt *NonTerminalEexpression) Interpret(context Context) string {
	return "nonterminal"
}
