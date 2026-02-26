package main

import (
	"fmt"
	"sort"
)

func main(){
	var arr=[3]int {1,21,3}
	fmt.Printf("\narr type =%T,len = %d",arr,len(arr))
	var slice = []int {-11,22,3,13,54}

	fmt.Println("slice=",slice)
	fmt.Printf("\nslice type =%T,len = %d",slice,len(slice))
	var slicenew =append(slice[1:2],5)
	fmt.Println("slicenew=",slicenew)
	fmt.Printf("\nslicenew type =%T,len = %d",slicenew,len(slicenew))

	
	fmt.Println("slice=",slice)

}