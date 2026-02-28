package main

import "fmt"

//there is no classes ,inheritance,super or parent in go lang
//as it increases complexity - against idiomaticity

type User struct { //struct name should be upercase public
	name  string  //not public
	class string
	age   int
	Age   int //accesible by other methods
}

func main() {
	kshitija := User{"kshhiti", "a", 24, 24}
	fmt.Println(kshitija)

}
