package main

import "fmt"



func main(){
//array are not used much in go lang

fmt.Println("yooyoyoyoyoy")
var ar[5]int
fmt.Println("arr",ar)
ar[0]=1
fmt.Println("arr",ar)
var arrS[4]string
fmt.Println("arrS",arrS)
var arrN = []string {"a","b","c"}
fmt.Println("arrN",arrN)
fmt.Println("arrN len",len(arrN))
fmt.Printf("arrN type = %T",arrN)
fmt.Printf("\n arrs type = %T",arrS)


}