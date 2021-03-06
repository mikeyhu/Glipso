package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type numericCombiner func(interfaces.Numeric, interfaces.Numeric) interfaces.Numeric

func numericFlatten(args []interfaces.Value, combiner numericCombiner) (interfaces.Value, error) {
	var all interfaces.Numeric
	head := true
	for i, v := range args {
		vAsN, ok := v.(interfaces.Numeric)
		if !ok {
			return NILL, fmt.Errorf("numericFlatten : expected Numeric but argument %v was %v", i+1, v)
		}
		if head {
			all = vAsN
			head = false
		} else {
			all = combiner(all, vAsN)
		}
	}
	return all, nil
}

// I (Integer)
type I int

// IsType for I
func (i I) IsType() {}

// IsValue for I
func (i I) IsValue() {}

// String representation of I
func (i I) String() string {
	return fmt.Sprintf("%d", i.Int())
}

// Int unboxes an int from Int
func (i I) Int() int {
	return int(i)
}

// Int converts an int to a float64
func (i I) float() float64 {
	return float64(i)
}

func (i I) asF() F {
	return F(i)
}

// Equals checks equality with another item of type Type
func (i I) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(I); ok {
		return B(i == other)
	}
	if other, ok := o.(F); ok {
		return B(i.asF() == other)
	}
	return B(false)
}

// CompareTo compares one I to another I and returns -1, 0 or 1
func (i I) CompareTo(o interfaces.Comparable) (int, error) {
	if other, ok := o.(I); ok {
		if i < other {
			return -1, nil
		} else if i == other {
			return 0, nil
		}
		return 1, nil
	}
	if other, ok := o.(F); ok {
		f := F(i)
		return f.CompareTo(other)
	}
	return 0, fmt.Errorf("CompareTo : Cannot compare %v to %v", i, o)
}

// Add for I
func (i I) Add(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(I); ok {
		return i + other
	}
	if other, ok := n.(F); ok {
		return i.asF() + other
	}
	panic("not implemented")
}

// Subtract for I
func (i I) Subtract(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(I); ok {
		return i - other
	}
	if other, ok := n.(F); ok {
		return i.asF() - other
	}
	panic("not implemented")
}

// Multiply for I
func (i I) Multiply(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(I); ok {
		return i * other
	}
	if other, ok := n.(F); ok {
		return i.asF() * other
	}
	panic("not implemented")
}

// Divide for I
func (i I) Divide(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(I); ok {
		return i / other
	}
	if other, ok := n.(F); ok {
		return i.asF() / other
	}
	panic("not implemented")
}

// Mod for I
func (i I) Mod(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(I); ok {
		return i % other
	}
	panic("not implemented")
}

// F (Float)
type F float64

// IsType for F
func (f F) IsType() {}

// IsValue for F
func (f F) IsValue() {}

// String representation of F
func (f F) String() string {
	return fmt.Sprintf("%f", f.float())
}

// float unboxes a float from F
func (f F) float() float64 {
	return float64(f)
}

// Equals checks equality with another item of type Type
func (f F) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(F); ok {
		return B(f == other)
	}
	if other, ok := o.(I); ok {
		return B(f == other.asF())
	}
	return B(false)
}

// CompareTo compares one F to another F and returns -1, 0 or 1
func (f F) CompareTo(o interfaces.Comparable) (int, error) {
	if other, ok := o.(F); ok {
		if f < other {
			return -1, nil
		} else if f == other {
			return 0, nil
		}
		return 1, nil
	}
	if other, ok := o.(I); ok {
		return f.CompareTo(F(other))
	}
	return 0, fmt.Errorf("CompareTo : Cannot compare %v to %v", f, o)
}

// Add for F
func (f F) Add(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(F); ok {
		return f + other
	}
	if other, ok := n.(I); ok {
		return f + other.asF()
	}
	panic("not implemented")
}

// Subtract for F
func (f F) Subtract(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(F); ok {
		return f - other
	}
	if other, ok := n.(I); ok {
		return f - other.asF()
	}
	panic("not implemented")
}

// Multiply for F
func (f F) Multiply(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(F); ok {
		return f * other
	}
	if other, ok := n.(I); ok {
		return f * other.asF()
	}
	panic("not implemented")
}

// Divide for F
func (f F) Divide(n interfaces.Numeric) interfaces.Numeric {
	if other, ok := n.(F); ok {
		return f / other
	}
	if other, ok := n.(I); ok {
		return f / other.asF()
	}
	panic("not implemented")
}
