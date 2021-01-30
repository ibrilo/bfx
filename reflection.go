package bfx

import "fmt"

type number interface {
	Int() int
	Int64() int64
	Float32() float32
	Float64() float64
}

type Number struct {
	val float64
}

func parseNumber(v interface{}) Number {
	val, ok := v.(float64)
	if !ok {
		iVal, ok := v.(int)
		if !ok {
			fmt.Printf("%v is not a number. failed to convert to float64.\n", val)
		}
		val = float64(iVal)
	}
	return Number{val}
}

func (n Number) Int() int {
	return int(n.val)
}

func (n Number) Int64() int64 {
	return int64(n.val)
}

func (n Number) Float32() float32 {
	return float32(n.val)
}

func (n Number) Float64() float64 {
	return n.val
}

type String struct {
	val string
}

func parseString(v interface{}) String {
	val, ok := v.(string)
	if !ok {
		fmt.Printf("%s is not a string. failed to convert to string.\n", val)
	}
	return String{val}
}

func (s String) String() string {
	return s.val
}

type Bool struct {
	val bool
}

func parseBool(v interface{}) Bool {
	val, ok := v.(bool)
	if !ok {
		fmt.Printf("%t is not a bool. failed to convert to bool.\n", val)
	}
	return Bool{val}
}

func (b Bool) Bool() bool {
	return b.val
}
