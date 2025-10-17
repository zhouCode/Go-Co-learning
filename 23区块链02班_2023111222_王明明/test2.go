package code

import "fmt"

func main() {
	a := 1;
	switch a {
	case 1:
		fmt.Println("1")
	
	case 2:
		fmt.Println("2")
	
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	}
	fmt.Println(2)
}