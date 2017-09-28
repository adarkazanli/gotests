package main

import (
	"fmt"
	"runtime"
	"reflect"
	"os"
)

var (
	name = "Ammar" // Name of the subscriber
	course = "my Courses" // Course attended
	module = 3.2 // Version
)

const speedOfLightMph = 186000

func main() {
    	fmt.Println("hello, world from: ", runtime.GOOS, "\n\n\n")

	fmt.Println("Name is set to: ", name, " and is of type: ", reflect.TypeOf(name))
	fmt.Println("Course is set to: ", course,  " and is of type: ", reflect.TypeOf(course))
	fmt.Println("Module: ", module,  " and is of type: ", reflect.TypeOf(module), "\n\n\n")

	a := 10.7
	b := 3
	fmt.Println("\nA has value of: ", a, " and is type: ", reflect.TypeOf(a))
	fmt.Println("\nB has value of: ", b, " and is type: ", reflect.TypeOf(b))
	
	c := int(a) + b
	fmt.Println("\nC has value: ", c, " and is of type:", reflect.TypeOf(c))
	fmt.Println("Memory location for A:", &a)
	fmt.Println("Memory location for B:", &b)
	fmt.Println(" AND Memory locationf or C:", &c)

	ptr := &c
	fmt.Println(" and the value of c is:", *ptr)

	changeMyC(a, b, &c)
	fmt.Println("\nYou have changed c to be:", c)
	fmt.Println("The speed of light is:", speedOfLightMph)
	fmt.Println("Environment Variables:", os.Environ())
	fmt.Println("-------------------------")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

func changeMyC(a float64, b int, c *int) int {
	*c = int(a) * b
	fmt.Println("\nTrying to change c:", *c)
	return *c
}
