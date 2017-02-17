package common

import (
	"github.com/mikeyhu/glipso/interfaces"
)

// Environment Provides a mechanism for creating and resolving variables
type Environment struct {
	variables map[string]interfaces.Type
	parent    *Environment
}

// ResolveRef will try to resolve a provided reference to a value in this or parent scope
func (env Environment) ResolveRef(ref interfaces.Type) (interfaces.Type, bool) {
	if result, ok := env.variables[ref.(REF).String()]; ok {
		return result, true
	}
	if env.parent != nil {
		return env.parent.ResolveRef(ref)
	}
	return nil, false
}

// CreateRef will create a variable in this scope
func (env Environment) CreateRef(name interfaces.Type, arg interfaces.Type) interfaces.Type {
	env.variables[name.(REF).String()] = arg
	return name
}

// NewChildScope creates new scope that inherits from this one
func (env Environment) NewChildScope() interfaces.Scope {
	return Environment{
		map[string]interfaces.Type{},
		&env,
	}
}

// GlobalEnvironment acts as the global scope for variables
var GlobalEnvironment Environment

func init() {
	GlobalEnvironment = Environment{
		variables: map[string]interfaces.Type{},
	}
}
