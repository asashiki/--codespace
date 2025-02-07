package main

import (
	"fmt"
)

// func main() {
// 	var name = "David"
// 	name = "42" //只能字符串类型,42报错 "42"✓

// 	fmt.Println(name)
// }


// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a number")
// 		return
// 	}

// 	input := os.Args[1]
// 	number, err := strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("Invalid number:", input)
// 		return
// 	}

// 	if number > 35 {
// 		fmt.Println("Too old")
// 	}else{
// 		fmt.Println("OK")
// 	}
// }

func main() {

	var age = 36
	if tooOld(age) {
		fmt.Println("You are too old")
	}
}

func tooOld(age int) bool {
	return age > 35
}

/*
package main

import (
	"fmt"
)

type Person struct {
	age int
}

func main() {
//	var age = 36
	var p = Person{age: 36}
	var m = make(map[string]int)
	m["age"] = 36

	change_p(&p)
	change_m(m)

	fmt.Println(p.age)
	fmt.Println(m["age"])
}

func change_p(p *Person) int {
	p.age = 20

	return p.age
}

func change_m(m map[string]int) {
	m["age"] = 20

}

*/