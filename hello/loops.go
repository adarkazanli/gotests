package main

import(
	"fmt"
	"time"
)

func main() {

	for timer:= 10; timer >= 0; timer-- {
		if timer ==0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
	

	childrensList := []string{"Anwar", "Kareem", "Ahmad", "Julliana"}
	WorkingList := []string{"Anwar", "Kareem", "Laith", "Rakan"}
	for _, i := range childrensList {
		fmt.Println(i)
		for _, j := range WorkingList {
			if i==j {
			fmt.Println("\nSomeone is working: ", i, " is working")
			}
		}
	}

	for timer:= 10; timer >= 0; timer-- {
		if timer % 2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}