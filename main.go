package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var min, max int
	fmt.Println("-----------------------------Number Guessing Game------------------------------")
	fmt.Println("Enter the Starting Number")
	fmt.Scan(&min)
	fmt.Println("Enter the Last Number")
	fmt.Scan(&max)
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("     Choose numbers between %v and %v and let's see your guessing skill \n", min, max)
	RandomIntegerwithinRange := rand.Intn(max-min) + min
	noofguess := 0
	for {
		var guessed int
		fmt.Println("Enter a Number you are guessing")
		fmt.Scan(&guessed)
		if guessed == RandomIntegerwithinRange {
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Println("                                Correct Answer")
			fmt.Printf("                 You have taken %v chances to guess", noofguess)
			break
		} else if guessed > RandomIntegerwithinRange {
			fmt.Println("Too High")
			noofguess++
		} else {
			fmt.Println("Too Low")
			noofguess++
		}

	}

	// min, max := 100, 500
	// fmt.Print(RandomIntegerwithinRange)

}
