package main

import "fmt"

func main() {
	var name string = "Akash"
	fmt.Println(name)
	fmt.Printf("The type of variable is %T \n", name)

	var isbool bool = true
	fmt.Println(isbool)
	fmt.Printf("The type of variable is %T \n", isbool)

	var smallVal int = 8
	fmt.Println(smallVal)
	fmt.Printf("The type of variable is %T \n", smallVal)

	var small int
	fmt.Println(small)
	fmt.Printf("The type of variable is %T \n", small)

	var str string
	fmt.Println(str)
	fmt.Printf("The type of variable is %T \n", str)
}
