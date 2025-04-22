package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//generate a random number generator 
	//between 1 and 100
	secretNumber := rand.Intn(100) + 1

	var guess int
    tries := 0

	fmt.Println("I am thinking of a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		tries++

		if guess < secretNumber{
			fmt.Println("Too low! Try again")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again")
		} else {
			fmt.Printf("Correct! You guessed it in %d tries.\n", tries)
			break
		}
	}

}