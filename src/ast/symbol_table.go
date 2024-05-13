package ast

import (
	"fmt"

	_ "github.com/SuneelFreimuth/vinland/src/ast/types"
)

type SymbolTable struct {
	Scopes []Scope
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{}
	table.Put("print")
	return table
}

func (st *SymbolTable) Empty() bool {
	return len(st.Scopes) == 0
}

func (st *SymbolTable) Enter() {
	st.Scopes = append(st.Scopes, make(Scope))
}

func (st *SymbolTable) Exit() {
	st.Scopes = st.Scopes[:len(st.Scopes)-1]
}

func (st *SymbolTable) Put(name string) (s Symbol, exists bool) {
	last := len(st.Scopes) - 1
	s, exists = st.Scopes[last][name]
	if exists {
		return
	}

	s = Symbol{Name: name}
	st.Scopes[last][name] = s
	return
}

func (st *SymbolTable) Lookup(name string) (s Symbol, exists bool) {
	for i := len(st.Scopes) - 1; i >= 0; i-- {
		s, exists = st.Scopes[i][name]
		if exists {
			return
		}
	}
	return
}


type Scope map[string]Symbol


type Symbol struct {
	Name  string
	Error SymbolError
	// Type types.Type
}

type SymbolError interface {
	String() string
}

type Undefined struct {
	symbol *Symbol
}

func (u *Undefined) String() string {
	return fmt.Sprintf("%s was not declared in this scope.", u.symbol.Name)
}
