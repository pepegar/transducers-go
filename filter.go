package transducers

func Filter(f func(interface{}) bool) CollToColl {
	return func(collection []interface{}) []interface{} {
		var result []interface{}

		for _, item := range collection {
			if f(item) {
				result = append(result, item)
			}
		}

		return result
	}
}
