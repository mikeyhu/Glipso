package main

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(+ 1 2 3 4 5)")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(15), result)
}

func TestApplyAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(apply + (cons 1 (cons 2 (cons 3))))")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestIfEvaluatesSecondExpression(t *testing.T) {
	code := `
	(if (= 1 1) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(4), result)
}

func TestIfEvaluatesThirdExpression(t *testing.T) {
	code := `
	(if (= 1 2) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestCreatesAndUsesVariable(t *testing.T) {
	code := `
	(do
		(def one 1)
		(def two 2)
		(+ one two))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(3), result)
}

func TestSummingRange(t *testing.T) {
	exp, _ := parser.Parse("(apply + (range 1 5))")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(15), result)
}

func TestGlobalAdd1Function(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ 1 a)))
		(add1 5))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestAnonymousAdd1Function(t *testing.T) {
	code := `
	((fn [a] (+ 1 a)) 5)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestEvenFunctionEvaluatesEvenNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 2)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}

func TestEvenFunctionEvaluatesOddNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 1)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(false), result)
}

func TestFilterEvenNumbers(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(apply + (filter even (cons 1 (cons 2 (cons 3 (cons 4))))))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestMapAdd1(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ a 1)))
		(first (map add1 (cons 1)))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(2), result)
}

func TestLazyPairHasAccessToClosure(t *testing.T) {
	code := `
	(do
		(def hasclosure (fn [a b] (lazypair a (lazypair b))))
		(apply + (hasclosure 1 10))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(11), result)
}
