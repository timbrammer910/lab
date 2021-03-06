package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// Bonus delete by val
	m["k3"] = 50
	keyFromVal := func(m map[string]int, val int) (key string) {
		for k, v := range m {
			if v == val {
				key = k
				return
			}
		}
		return
	}

	delete(m, keyFromVal(m, 50))
	fmt.Println("del:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
