package main
import "fmt"

func main() {
 	s := make([] string,3,5)
 	parent := [...]int{10,20,30,40,50}
 	child := parent[1:4]
 	fmt.Println("Slice:",s)
 	fmt.Println("len:",len(s))
 	fmt.Println("Cap:", cap(s))
 	fmt.Println("child:",child)
 	child[0] = 99
 	fmt.Println("After modify child,Parent",parent)
}
