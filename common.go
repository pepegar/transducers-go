package transducers

type CollToColl func([]interface{}) []interface{}

func Comp(fn1 CollToColl, fn2 CollToColl) CollToColl {
	ret := func(coll []interface{}) []interface{} {
		return fn1(fn2(coll))
	}

	return ret
}
