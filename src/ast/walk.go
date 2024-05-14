package ast


func Walk(ast Node, l Listener) {
	(&walker{}).walk(ast, l)
}


type walker struct{}

func (w *walker) walk(n Node, l Listener) {
	l.EnterEveryNode(n)
	n.EnterNode(l)
	for _, child := range n.Children() {
		w.walk(child, l)
	}
	n.ExitNode(l)
	l.ExitEveryNode(n)
}
