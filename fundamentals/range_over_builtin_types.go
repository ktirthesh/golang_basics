package main

import "fmt"

func range_over_builtin_types_golang() {
	nums := []int{1, 2, 3, 4, 5}
	for _, val := range nums {
		fmt.Println("value in nums is ", val)
	}
	for _, val := range nums {
		if val == 3 {
			fmt.Println("valu is 3", val)
		}
	}

	kvs := map[string]string{"hello": "world", "foo": "bar"}
	for k, v := range kvs {
		fmt.Println("key and value is", k, ":", v)
	}

	for key := range kvs {
		fmt.Println("key", key)
	}

	for i, c := range "gooo" {
		fmt.Println("i,c", i, c) // print unicode points
	}
}
