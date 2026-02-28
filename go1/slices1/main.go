package main

import (
	"fmt"

)

func main() {

	var arr = [3]int{1, 21, 3}
	fmt.Printf("\narr type =%T,len = %d", arr, len(arr))
	var slice = []int{-11, 22, 3, 13, 54}

	fmt.Println("slice=", slice)
	fmt.Printf("\nslice type =%T,len = %d", slice, len(slice))
	var slicenew = append(slice[1:2], 5)
	fmt.Println("slicenew=", slicenew)
	fmt.Printf("\nslicenew type =%T,len = %d", slicenew, len(slicenew))

	fmt.Println("slice=", slice)
	tryMake := make([]int, 3)
	tryMake[0] = 2
	tryMake[1] = 3
	tryMake[2] = 4
	fmt.Printf("\n tryMake type =%T,len = %d", tryMake, len(tryMake))

	tryMake = append(tryMake, 2, 3, 4, 5, 6, 7)
	fmt.Println("tryMake ", tryMake)
	fmt.Printf("\n tryMake type =%T,len = %d", tryMake, len(tryMake)) //reallocates memory here
	//tryMake[10]=4// will give out of mem as trymake is allocated only 9 size

	//to remove index from a slice

	var inx = 2
	slice = append(slice[:inx], slice[inx+1:]...)
	fmt.Println("SLICE", slice)

}
