package main

  import "fmt"
func main(){
    c := make(chan int)

    go strl("Michael", c)


    x := <- c

    fmt.Println(" The string length is ", x)
}


func strl(s string , c chan int){

  c <- len(s)

}