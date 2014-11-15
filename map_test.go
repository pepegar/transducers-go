package transducers

import "testing"

func mapping(item interface{}) interface{} {
	integer, _ := item.(int)

	return integer * 3
}

func TestMap(t *testing.T) {
	mapTrans := Map(mapping)

	coll := []int{1, 2, 3, 4, 5}
	var interfaceSlice []interface{} = make([]interface{}, len(coll))

	for i, item := range coll {
		interfaceSlice[i] = item
	}

	expectedIntSlice := []int{3, 6, 9, 12, 15}
	var expectedInterfaceSlice []interface{} = make([]interface{}, len(expectedIntSlice))
	for i, item := range expectedIntSlice {
		expectedInterfaceSlice[i] = item
	}

	if !equals(expectedInterfaceSlice, mapTrans(interfaceSlice)) {
		t.Error("Comp output doesn't match expected output")
	}
}
