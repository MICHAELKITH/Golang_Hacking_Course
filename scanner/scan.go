package main

  import "fmt"
func main(){
    c := make(chan int)

    go strl("Golang", c)


    y := <- c

    fmt.Println(" The string length is ", y)
}


func strl(s string , c chan int){

  c <- len(s)

}