package main

type numbah int

var x numbah
func main(){
fmt.Println("x")
fmt.Printf("%T\n",x)
x = 42
fmt.Println(x)
}
