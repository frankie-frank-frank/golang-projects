package main 

import "fmt"

type intSlice []int

func (i intSlice) EvenOrOdd() {
	for _, item := range i{
		if item%2 == 0 {
			fmt.Printf("%v is even\n", item)
		} else {
			fmt.Printf("%v is odd\n", item)
		}	
	}
}

func main(){
	integerSlice := intSlice{0,1,2,3,4,5,6,7,8,9,10}
	integerSlice.EvenOrOdd()
}