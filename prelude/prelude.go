package prelude

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/mikeyhu/glipso/parser"
)

// ParsePrelude loads a number of definitions such as functions into global scope
func ParsePrelude(scope interfaces.Scope) {
	prelude := `
	(do
		(def defmacro (macro [n a e] (def n (macro a e))))
		(defmacro defn [nn aa ee] (def nn (fn aa ee)))
	)
	`
	exp, err := parser.Parse(prelude)
	if err != nil {
		panic(fmt.Sprintf("Error parsing prelude, error %v", err))
	}
	exp.Evaluate(scope)
}
