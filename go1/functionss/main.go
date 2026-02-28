package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("mkkkkkkk")
	v:=[4]int{2,3,4,4}
	r:=kki(v)
	
	fmt.Println(r)
	fmt.Println(v)
	// var v int32 = 10
	// var f float32 = 3.2
	// var add bool = bool(v) + bool(f)
	//fmt.Printf("add %T",add)

	k := struct1{"AA", 10, false}
	bol := k.Getstuct()
	fmt.Println("C for struct is ", k.C)
	fmt.Println("C for struct is ", bol)
	checkdefer()

}

func k(v1 int, v2 int, v3 int) (int, string) {
	fmt.Printf("kkkk\n")
	return (v1 + v2 + v3), "bruuh"
}

func kki(v1 [4]int) int {
	//fmt.Printf("v1 is %T\n", v1)
	// var ar [4]int
	// fmt.Printf("ar is %T\n",ar)
	ans := 0
	for _, v := range v1 {
		ans = ans + v
	}
	v1[0]=44
	fmt.Println("ar is \n",v1)
	return (ans)
}

// methods -function in structs
type struct1 struct {
	A string
	B int
	C bool
}

func (s *struct1) Getstuct() bool { //pass a xopy of s struct1 ...use pointers to pass by reference
	defer fmt.Println("Prrinted?????")
	s.C = true
	fmt.Println("Pr")
	return s.C
}

//defer
// if defer statement is in a function then it will be exceuted before function clsing '}' and 
// if thre are multiple defer statements then defer will executed in LIFO order eg below


func checkdefer(){
	read:= bufio.NewReader(os.Stdin)
	s,_:= read.ReadString('\n') 
	defer fmt.Println("1") 
	defer fmt.Println("2")
	defer fmt.Println("3")
	 fmt.Println("normalllll flow")
	 defer fmt.Println("4")
	 //const s string ="abcdjdd"
	a,_:= strconv.ParseInt(strings.TrimSpace(s),0,32)
	fmt.Printf("A IS of type %T and value %v \n",a,a)

}//defered stack could be assumed like 1,2,3,4
//output normalllll flow,4,3,2,1 LIFO Mcdonalds queue
