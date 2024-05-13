package types


type Type interface{
	Equivalent(Type) bool
}


type IntType struct{}

func NewIntType() *IntType {
	return &IntType{}
}

func (t *IntType) Equivalent(other Type) bool {
	_, ok := other.(*IntType)
	return ok
}


type FloatType struct{}

func NewFloatType() *FloatType {
	return &FloatType{}
}

func (t *FloatType) Equivalent(other Type) bool {
	_, ok := other.(*FloatType)
	return ok
}


type StringType struct{}

func NewStringType() *StringType {
	return &StringType{}
}

func (t *StringType) Equivalent(other Type) bool {
	_, ok := other.(*StringType)
	return ok
}


type FuncType struct{}

func NewFuncType() *FuncType {
	return &FuncType{}
}

func (t *FuncType) Equivalent(other Type) bool {
	_, ok := other.(*FuncType)
	return ok
}

