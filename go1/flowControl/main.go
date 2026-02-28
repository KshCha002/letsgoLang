package main

import (
	"math/rand"
	"fmt"		
	"time"
)

func main() {
	i:=10
	if i>0{
		fmt.Println("Bada")
	}else{
	fmt.Println("ho ga")
	
	}
	rand.Seed(time.Now().UnixNano())
	ran:= rand.Intn(3)

	//swaithc case

	switch ran{
	case 0:
		fmt.Println("Ho jaega jldi") 
		fallthrough
	case 1:
		fmt.Println("Nhi hoga") 
	default:
		fmt.Println("Ho jaega not jldi") 

	}
//loops
days := []string{"sun","mo","tu","we","fri","sat"}
// for d:=0;d<len(days);d++{
// fmt.Println(days[d])
// }

// for i:= range days{
// 	fmt.Println(days[i])
// }
las: // Correct label syntax (no comments after labels)
    fmt.Println("Starting loop")

for index,day:=range days{ //for each
	if day =="mo"{
		continue
	}

		fmt.Printf("ind = %d & value %v \n",index,day)
if day =="thur"{
		
		goto las
		break
	}
	}

	
}
