// key value pairs

package main

import (
//	"crypto/rand"
	"fmt"
)


func main(){
	//fmt.Println("Maps")
	language1 := make(map[int]string)
	//lang2 := new(map[int]int)
	//fmt.Println("language",language1);
	//fmt.Println(lang2);
	language1[0]= "a"
	language1[2]="b"
	language1[3]="c"
	language1[4]="e"
	
	map1:= map[int]int{}
	map1[21]=0
	fmt.Println("map1",map1);

//fmt.Println("0 is ",language1[0]);

//delete(language1,2)

fmt.Println("l is ",language1);

for key,value := range language1{
	fmt.Printf("YAYAYAYAY %d = %s \n",key,value)
}
}