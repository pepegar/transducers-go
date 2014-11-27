package transducers

func Comp(fn1 func([]interface{}) []interface{}, fn2 func([]interface{}) []interface{}) func([]interface{}) []interface{} {
	ret := func(coll []interface{}) []interface{} {
		return fn1(fn2(coll))
	}

	return ret
}
