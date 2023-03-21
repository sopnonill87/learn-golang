package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%.2f please.\n", 0.23*5)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Values: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------------------------------")

}
