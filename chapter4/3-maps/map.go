package main

import (
	"fmt"
	"sort"
)

func main() {
	// ages := map[string]int{
	// 	"alice":   31,
	// 	"charlie": 34,
	// }
	// // or
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34
	// or
	ages := map[string]int{}
	ages["alice"] = 31
	ages["charlie"] = 34

	ages["alice"] = 32
	fmt.Println(ages["alice"])
	// delete(ages, "alice")
	fmt.Println(ages["alice"])

	ages["bob"]++
	fmt.Println("bob:", ages["bob"])

	for name, age := range ages {
		fmt.Println(name, age)
	}

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, ages[name])
	}

	var ages2 map[string]int
	fmt.Println(ages2, ages)
	fmt.Println(ages2 == nil)
	fmt.Println(len(ages2) == 0)

	// error, nil
	// ages2["sss"] = 1

	// age, ok := ages["fred"]
	// if !ok {
	// 	fmt.Print(ok, " ")
	// }
	// fmt.Println("fred", age)
	// or
	if age, ok := ages["fred"]; !ok {
		fmt.Println("fred", ok, age)
	}

	fmt.Println("eql", equal(map[string]int{"B": 1}, map[string]int{"B": 1}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
