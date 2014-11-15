[![Build Status](https://travis-ci.org/pepegar/transducers-go.svg)](https://travis-ci.org/pepegar/transducers-go)
transducers-go
==============

Clojure's transducers library port to Go.

usage:

```go
package main

import (
	"fmt"

	t "github.com/pepegar/transducers-go"
)

func mapping(item interface{}) interface{} {
	integer, _ := item.(int)

	return integer * 3
}

func filtering(item interface{}) bool {
	integer, _ := item.(int)

	return integer%2 == 0
}

func main() {
	mapTrans := t.Map(mapping)
	filterTrans := t.Filter(filtering)
	transducer := t.Comp(mapTrans, filterTrans)

	coll := []int{1, 2, 3, 4, 5}
	var interfaceColl []interface{} = make([]interface{}, len(coll))

	for i, item := range coll {
		interfaceColl[i] = item
	}

	fmt.Println(transducer(interfaceColl))
}
```
