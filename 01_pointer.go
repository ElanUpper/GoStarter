package main

import "fmt"

func change_val(numb *int){
	*numb += 10;
}


func main(){
	numb_a := 10 ;
	change_val(&numb_a)
	fmt.Printf("%d \n ", numb_a)
}
