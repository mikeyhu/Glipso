package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

func plusAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func minusAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var all I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(I)
			head = false
		} else {
			all -= v.(I)
		}
	}
	return all
}

func multiplyAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var all I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(I)
			head = false
		} else {
			all *= v.(I)
		}
	}
	return all
}

func mod(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	a, aok := arguments[0].(I)
	b, bok := arguments[1].(I)
	if aok && bok {
		return I(a % b)
	}
	panic("Mod : unsupported type")
}

func equals(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Equalable)
	second, sok := arguments[1].(interfaces.Equalable)

	if fok && sok {
		return first.Equals(second)
	}
	panic(fmt.Sprintf("Equals : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThan(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) < 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThanEqual(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) <= 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThan(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) > 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThanEqual(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) >= 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func cons(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	if len(arguments) == 0 {
		return P{}
	} else if len(arguments) == 1 {
		return P{arguments[0], nil}
	} else if len(arguments) == 2 {
		tail, ok := arguments[1].(P)
		if ok {
			return P{arguments[0], &tail}
		}
	}
	return P{}
}

func first(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(P)
	if ok {
		return pair.head
	}
	fmt.Printf("pair? %v : %v\n", arguments[0], pair)
	panic("Panic - Cannot get head of non Pair type")
}

func tail(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(P)
	if ok {
		return *pair.tail
	}
	panic("Panic - Cannot get tail of non Pair type")
}

func apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	if ap, okEval := arguments[1].(interfaces.Evaluatable); okEval {
		arguments[1] = ap.Evaluate(sco)
	}
	s, okRef := arguments[0].(REF)
	p, okPair := arguments[1].(interfaces.Iterable)
	if !okRef {
		panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
	} else if !okPair {
		panic(fmt.Sprintf("Panic - expected pair, found %v", arguments[1]))
	}
	return &EXP{Function: s, Arguments: p.ToSlice(sco.NewChildScope())}
}

func iff(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var test interfaces.Type
	if exp, ok := arguments[0].(*EXP); ok {
		test = exp.Evaluate(sco)
	} else {
		test = arguments[0]
	}
	if test.(B).Bool() {
		return arguments[1]
	}
	return arguments[2]
}

func def(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var value interfaces.Type
	if eval, ok := arguments[1].(interfaces.Evaluatable); ok {
		value = eval.Evaluate(sco.NewChildScope())
	} else {
		value = arguments[1]
	}
	return GlobalEnvironment.CreateRef(arguments[0].(REF).EvaluateToRef(sco.NewChildScope()), value)
}

func do(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var result interfaces.Type
	for _, a := range arguments {
		if e, ok := a.(interfaces.Evaluatable); ok {
			result = e.Evaluate(sco.NewChildScope())
		} else {
			result = a
		}
	}
	return result
}

func rnge(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	start := arguments[0].(I)
	end := arguments[1].(I)
	if start < end {
		return LAZYP{
			start,
			&EXP{Function: REF("range"), Arguments: []interfaces.Type{
				I(start.Int() + 1),
				end,
			}}}
	}
	return LAZYP{end, nil}

}

func fn(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var argVec VEC
	if args, ok := arguments[0].(REF); ok {
		argVec = args.Evaluate(sco).(VEC)
	} else {
		argVec = arguments[0].(VEC)
	}

	if arg1, ok := arguments[1].(REF); ok {
		return FN{argVec, arg1.Evaluate(sco.NewChildScope()).(*EXP)}
	}
	return FN{argVec, arguments[1].(*EXP)}
}

func filter(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	fn, fnok := arguments[0].(FN)
	pair, pok := arguments[1].(P)

	var flt func(*P) *P
	flt = func(p *P) *P {
		head := p.head
		res := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if include, iok := res.(B); iok {
			if bool(include) {
				if p.tail != nil {
					return &P{head, flt(p.tail)}
				}
				return &P{head, nil}
			} else if p.tail != nil {
				return flt(p.tail)
			}
			return nil
		}
		panic(fmt.Sprintf("filter : expected boolean value, recieved %v", res))
	}

	if fnok && pok {
		return *flt(&pair)
	}
	panic("filter : expected function and list")
}

func mapp(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	fn, fnok := arguments[0].(FN)
	pair, pok := arguments[1].(P)

	var mp func(*P) *P
	mp = func(p *P) *P {
		head := p.head
		res := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if p.tail != nil {
			return &P{res, mp(p.tail)}
		}
		return &P{res, nil}
	}

	if fnok && pok {
		return *mp(&pair)
	}
	panic("map : expected function and list ")
}

func lazypair(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	fmt.Printf("lazypair\n")
	sco.(Environment).DisplayEnvironment()
	head := arguments[0]
	if h, ok := head.(interfaces.Evaluatable); ok {
		head = h.Evaluate(sco)
	}
	if len(arguments) > 1 {
		if tail, ok := arguments[1].(interfaces.Evaluatable); ok {
			fmt.Println("Adding Deferred Eval to tail: ", tail)
			sco.(Environment).DisplayEnvironment()

			return LAZYP{head, BindEvaluation(tail, sco)}
		}
		panic(fmt.Sprintf("lazypair : expected EXP got %v", arguments[1]))
	}
	return LAZYP{head, nil}
}

type evaluator func([]interfaces.Type, interfaces.Scope) interfaces.Type

type FunctionInfo struct {
	function     evaluator
	evaluateArgs bool
}

var inbuilt map[string]FunctionInfo

func init() {
	inbuilt = map[string]FunctionInfo{
		"cons":     {cons, true},
		"first":    {first, true},
		"tail":     {tail, true},
		"=":        {equals, true},
		"+":        {plusAll, true},
		"-":        {minusAll, true},
		"*":        {multiplyAll, true},
		"%":        {mod, true},
		"<":        {lessThan, true},
		">":        {greaterThan, true},
		"<=":       {lessThanEqual, true},
		">=":       {greaterThanEqual, true},
		"apply":    {apply, false},
		"if":       {iff, false},
		"def":      {def, false},
		"do":       {do, false},
		"range":    {rnge, true},
		"fn":       {fn, false},
		"filter":   {filter, true},
		"map":      {mapp, true},
		"lazypair": {lazypair, false},
	}
}
