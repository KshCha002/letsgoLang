package main

import "fmt"


func main(){

	fmt.Println("Pointers")

//var are sorted in some memory 
var ptr *int 
fmt.Println("pte =",ptr)
num:=10
ptr=&num
fmt.Println("pte =",ptr)
fmt.Println("pte =",*ptr)
*ptr=*ptr +1
fmt.Println("pte =",*ptr)
fmt.Println("num=",num)
}

