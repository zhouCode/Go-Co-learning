package main
import "fmt"

func main() {
 	a :=[...]string{"USE","CHINA","JAPAN"}
 	b :=a
 	fmt.Println("修改前:")
 	fmt.Println("a:",a)
 	fmt.Println("b",b)
 	b[0] = "Singapore"
 	fmt.Println("修改后:")
 	fmt.Println("a:",a)
 	fmt.Println("b",b)
}
