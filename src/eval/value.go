package eval

import (
	"fmt"
)


type Value interface {
	Children() []Value
}


type BaseValue struct{}

func (*BaseValue) Children() []Value {
	return nil
}


type Int struct {
	BaseValue
	Value int64
}

func NewInt(value int64) Int {
	return Int{Value: value}
}

func (n *Int) String() string {
	return fmt.Sprintf("\"%d\"\n", n.Value)
}


type Float struct {
	BaseValue
	Value float64
}

func NewFloat(value float64) Float {
	return Float{Value: value}
}

func (f *Float) String() string {
	return fmt.Sprintf("\"%f\"\n", f.Value)
}


type String struct {
	BaseValue
	Value string
}

func NewString(value string) String {
	return String{Value: value}
}

func (s *String) String() string {
	return fmt.Sprintf("\"%s\"\n", s.Value)
}
