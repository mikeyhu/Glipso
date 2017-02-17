package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualityNotEqual(t *testing.T) {
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{B(true), B(false)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, B(false), result)
}

func TestEqualityEqual(t *testing.T) {
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{B(true), B(true)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, B(true), result)
}

func TestEqualityPanicsIfTypesNotValid(t *testing.T) {
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{P{}, I(1)}}
	assert.Panics(t, func() {
		exp.Evaluate(GlobalEnvironment)
	})
}

func TestConsCreatesPairWithNil(t *testing.T) {
	exp := EXP{Function: REF("cons"), Arguments: []interfaces.Type{I(1)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, P{I(1), nil}, result)
}

func TestConsCreatesPairWithTailPair(t *testing.T) {
	exp := EXP{Function: REF("cons"), Arguments: []interfaces.Type{I(1), P{I(2), nil}}}
	result := exp.Evaluate(GlobalEnvironment).(P)
	assert.Equal(t, I(1), result.head)
	assert.Equal(t, I(2), result.tail.head)
	assert.Nil(t, result.tail.tail)
}

func TestFirstRetrievesHeadOfPair(t *testing.T) {
	exp := EXP{Function: REF("first"), Arguments: []interfaces.Type{P{I(3), nil}}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(3), result)
}

func TestConsRetrievesTailOfPair(t *testing.T) {
	exp := EXP{Function: REF("tail"), Arguments: []interfaces.Type{P{I(5), &P{I(6), nil}}}}
	result := exp.Evaluate(GlobalEnvironment).(P)
	assert.Equal(t, I(6), result.head)
}

func TestApplySendsListToFunction(t *testing.T) {
	exp := EXP{Function: REF("apply"), Arguments: []interfaces.Type{REF("+"), P{I(2), &P{I(10), nil}}}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(12), result)
}

func TestIfTrueReturnsSecondArgument(t *testing.T) {
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{B(true), I(1), I(2)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(1), result)
}

func TestIfFalseReturnsThirdArgument(t *testing.T) {
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{B(false), I(1), I(2)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(2), result)
}

func TestDefRecordsReferences(t *testing.T) {
	exp := EXP{Function: REF("def"), Arguments: []interfaces.Type{REF("one"), I(1)}}
	exp.Evaluate(GlobalEnvironment)
	resolved, _ := GlobalEnvironment.ResolveRef(REF("one"))
	assert.Equal(t, I(1), resolved)
}

func TestDoReturnsLastArgument(t *testing.T) {
	exp := EXP{Function: REF("do"), Arguments: []interfaces.Type{I(1), I(2)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t, I(2), result)
}

func TestRangeReturnsLazyPair(t *testing.T) {
	exp := EXP{Function: REF("range"), Arguments: []interfaces.Type{I(1), I(10)}}
	result := exp.Evaluate(GlobalEnvironment)
	assert.Equal(t,
		LAZYP{I(1), &EXP{Function: REF("range"), Arguments: []interfaces.Type{I(2), I(10)}}},
		result)
}
