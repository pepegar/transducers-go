package transducers

func Map(f func(interface{}) interface{}) func([]interface{}) []interface{} {
	return func(collection []interface{}) []interface{} {
		result := make([]interface{}, len(collection))

		for i, item := range collection {
			result[i] = f(item)
		}

		return result
	}
}
