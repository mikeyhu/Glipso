package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualityNotEqual(t *testing.T) {
	exp := EXP{FunctionName: "=", Arguments: []interfaces.Argument{B(true), B(false)}}
	result := exp.Evaluate()
	assert.Equal(t, B(false), result)
}

func TestEqualityEqual(t *testing.T) {
	exp := EXP{FunctionName: "=", Arguments: []interfaces.Argument{B(true), B(true)}}
	result := exp.Evaluate()
	assert.Equal(t, B(true), result)
}

func TestEqualityPanicsIfTypesNotValid(t *testing.T) {
	exp := EXP{FunctionName: "=", Arguments: []interfaces.Argument{P{}, I(1)}}
	assert.Panics(t, func() {
		exp.Evaluate()
	})
}

func TestConsCreatesPairWithNil(t *testing.T) {
	exp := EXP{FunctionName: "cons", Arguments: []interfaces.Argument{I(1)}}
	result := exp.Evaluate()
	assert.Equal(t, P{I(1), nil}, result)
}

func TestConsCreatesPairWithTailPair(t *testing.T) {
	exp := EXP{FunctionName: "cons", Arguments: []interfaces.Argument{I(1), P{I(2), nil}}}
	result := exp.Evaluate().(P)
	assert.Equal(t, I(1), result.head)
	assert.Equal(t, I(2), result.tail.head)
	assert.Nil(t, result.tail.tail)
}

func TestFirstRetrievesHeadOfPair(t *testing.T) {
	exp := EXP{FunctionName: "first", Arguments: []interfaces.Argument{P{I(3), nil}}}
	result := exp.Evaluate()
	assert.Equal(t, I(3), result)
}

func TestConsRetrievesTailOfPair(t *testing.T) {
	exp := EXP{FunctionName: "tail", Arguments: []interfaces.Argument{P{I(5), &P{I(6), nil}}}}
	result := exp.Evaluate().(P)
	assert.Equal(t, I(6), result.head)
}

func TestApplySendsListToFunction(t *testing.T) {
	exp := EXP{FunctionName: "apply", Arguments: []interfaces.Argument{REF("+"), P{I(2), &P{I(10), nil}}}}
	result := exp.Evaluate()
	assert.Equal(t, I(12), result)
}

func TestIfTrueReturnsSecondArgument(t *testing.T) {
	exp := EXP{FunctionName: "if", Arguments: []interfaces.Argument{B(true), I(1), I(2)}}
	result := exp.Evaluate()
	assert.Equal(t, I(1), result)
}

func TestIfFalseReturnsThirdArgument(t *testing.T) {
	exp := EXP{FunctionName: "if", Arguments: []interfaces.Argument{B(false), I(1), I(2)}}
	result := exp.Evaluate()
	assert.Equal(t, I(2), result)
}

func TestDefRecordsReferences(t *testing.T) {
	exp := EXP{FunctionName: "def", Arguments: []interfaces.Argument{REF("one"), I(1)}}
	exp.Evaluate()
	assert.Equal(t, I(1), GlobalEnvironment.resolveRef(REF("one")))
}

func TestDoReturnsLastArgument(t *testing.T) {
	exp := EXP{FunctionName: "do", Arguments: []interfaces.Argument{I(1), I(2)}}
	result := exp.Evaluate()
	assert.Equal(t, I(2), result)
}

func TestRangeReturnsLazyPair(t *testing.T) {
	exp := EXP{FunctionName: "range", Arguments: []interfaces.Argument{I(1), I(10)}}
	result := exp.Evaluate()
	assert.Equal(t,
		LAZYP{I(1), &EXP{FunctionName: "range", Arguments: []interfaces.Argument{I(2), I(10)}}},
		result)
}
