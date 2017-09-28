package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than second course")
	} else if firstRank < secondRank {
		fmt.Println("\nSecond course is doing better than first course")
	} else {
		fmt.Println("\nThey are about the same.")
	}


	switch tmpNum := random(); tmpNum {
		case 0, 2, 4, 6, 8:
			fmt.Println("We have got an even number: ", tmpNum)
		case 1,3, 5,9:
			fmt.Println("We have got an odd number: ", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}