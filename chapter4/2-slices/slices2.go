package main

import "fmt"

func main() {
	//nonempty
	// data := []string{"one", "", "three"}
	// fmt.Println(nonempty(data))
	// fmt.Println(data)

	//nonempty2
	// data := []string{"one", "", "three"}
	// fmt.Println(nonempty2(data))
	// fmt.Println(data)

	//remove
	// sl := []int{1, 2, 3, 4, 5}
	// sl = remove(sl, 3)
	// fmt.Println(sl)

	// remove2
	sl2 := []int{1, 2, 3, 4, 5}
	sl2 = remove2(sl2, 0)
	fmt.Println(sl2)
}

//nonempty возвращает сред, содержащий только непустые строки
//содержание базового массива изменяется
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
