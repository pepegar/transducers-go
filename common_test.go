package transducers

import "testing"

func equals(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i, itemA := range a {
		if itemA != b[i] {
			return false
		}
	}

	for i, itemB := range b {
		if itemB != a[i] {
			return false
		}
	}

	return true
}

func TestComp(t *testing.T) {
	mapTrans := Map(mapping)
	filterTrans := Filter(filtering)
	transducer := Comp(mapTrans, filterTrans)

	coll := []int{1, 2, 3, 4, 5}
	var interfaceSlice []interface{} = make([]interface{}, len(coll))

	for i, item := range coll {
		interfaceSlice[i] = item
	}

	expectedIntSlice := []int{6, 12}
	var expectedInterfaceSlice []interface{} = make([]interface{}, len(expectedIntSlice))
	for i, item := range expectedIntSlice {
		expectedInterfaceSlice[i] = item
	}

	if !equals(expectedInterfaceSlice, transducer(interfaceSlice)) {
		t.Error("Comp output doesn't match expected output")
	}
}
