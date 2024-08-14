package main


import "fmt"



func main (){

	x:=20
	 p(&x)
	fmt.Println("Learning", x)
}



func p (p *int){

	*p++
}