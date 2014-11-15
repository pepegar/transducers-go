package transducers

import "testing"

func filtering(item interface{}) bool {
	integer, _ := item.(int)

	return integer%2 == 0
}

func TestFilter(t *testing.T) {
	filterTrans := Filter(filtering)
	coll := []int{1, 2, 3, 4, 5}
	var interfaceSlice []interface{} = make([]interface{}, len(coll))

	for i, item := range coll {
		interfaceSlice[i] = item
	}

	expectedIntSlice := []int{2, 4}
	var expectedInterfaceSlice []interface{} = make([]interface{}, len(expectedIntSlice))
	for i, item := range expectedIntSlice {
		expectedInterfaceSlice[i] = item
	}

	if !equals(expectedInterfaceSlice, filterTrans(interfaceSlice)) {
		t.Error("Filter output doesn't match expected output")
	}
}
