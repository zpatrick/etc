package main

import "fmt"

func main() {
	m := map[string]string{"key": "val"}
	fmt.Println(m["key"])
	m["key"] = "value"

	// existence check (a.k.a. the "comma ok" idiom)
	if v, ok := m["key"]; ok {
		fmt.Printf("Key exists! Value is: %s\n", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
		delete(m, k)
	}
}
