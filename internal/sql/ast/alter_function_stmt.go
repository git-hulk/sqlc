package ast

import ()

type AlterFunctionStmt struct {
	Func    *ObjectWithArgs
	Actions *List
}

func (n *AlterFunctionStmt) Pos() int {
	return 0
}
