package main

import (
	"fmt"
	lru "github.com/tetleytea/cachein/lru"
)

type Employee struct {
	id int
	name string
}

func main() {
	fmt.Println("This is a start to this project")
	cache, err := lru.NewCache(5)
	if cache != nil {
		fmt.Println("Created cache")
		fmt.Println(err)
		
		employee := &Employee {
			123,
			"Peter",
		}

		cache.Add("test1",employee)
		
		value, ok := cache.Get("test1")

		if value != nil && ok == true {
			fmt.Println(ok)
			fmt.Println(value)
		}


		cache.Add("test2", &Employee{
			123,
			"James",
		})

		value2, ok2 := cache.Get("test2")

		if(value2 != nil && ok2 == true) {
			value, ok = cache.Get("test1")

			if(value == nil && ok == false) {
				fmt.Println(value)
				fmt.Println(ok)
			}
		}

	}
}