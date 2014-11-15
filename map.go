package transducers

func Map(f func(interface{}) interface{}) CollToColl {
	return func(collection []interface{}) []interface{} {
		result := make([]interface{}, len(collection))

		for i, item := range collection {
			result[i] = f(item)
		}

		return result
	}
}
