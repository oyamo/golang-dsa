package main

import (
	"dsa-golang/hashmap"
	"fmt"
)

func main() {
	h := hashmap.NewHashMap[int](16, .75)

	for i := 0; i < 20; i++ {
		h.Put(fmt.Sprintf("key_%d", i), i)
	}

	for i := 0; i < 16; i++ {
		key := fmt.Sprintf("key_%d", i)
		f, v := h.Get(key)
		fmt.Println(f, v)
	}

	fmt.Println(h.Keys())
	fmt.Println(len(h.Keys()))

}
