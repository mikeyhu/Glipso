package interfaces

type Evaluatable interface {
	Evaluate(Scope) Type
}

type Type interface {
	IsType()
}

type Iterable interface {
	Iterate(Scope) Iterable
	ToSlice(Scope) []Type
}

type Scope interface {
	ResolveRef(argument Type) Type
	CreateRef(ref Type, arg Type) Type
	NewChildScope() Scope
}
