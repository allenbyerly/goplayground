//This is the hello world exaple in go. It displays a hello world prompt.
//In addition it displays prompts to demonstrate some basic functions.

//Author: Allen Byerly
//Name: hello

//The Main Package Definition
package main

//The Imported Libraries
import (
  //import format library
  "fmt"
  //import time library
  "time"
  //import the math library
  "math"
  //import random number functions from the math library
  "math/rand"
  )


//added a basic addition function to perform addition of two numbers
//Shortend parameter inputs to only one int declaration
//note: this is similar to how C declares variables of the same type
func add(x, y int) int {
  return x + y
}

//Main Package Function Definition
func main() {
  //Seed the random number generator based on the current time (UTC Unix Nano)
  // Use
  rand.Seed(time.Now().UTC().UnixNano())
  //Print a Name Prompt
  fmt.Println("Hello World!")
  //Print a Description Prompt
  fmt.Println("This is my Go Playground!")
  //Print a current time Prompt
  fmt.Println("The time is: ", time.Now())
  //Print a random number prompt
  //note: the random number will always return the same number until it is seeded
  fmt.Println("A random number is: ", rand.Intn(100))
  //Print a problems Promt based off the square root of the random number
  //Display the number as a defualt variable
  fmt.Printf("Now you have %v problems.\n", int(math.Sqrt(rand.Float64()*100)))
  //Print a prompt that displays Pi
  fmt.Printf("Pi is: %v\n", math.Pi)
  //add some of my favorite numbers together
  //use variables this time
  favorite_numbers := []int {
    42,
    22,
    13,
    222,
  }
  fmt.Println(favorite_numbers, "are some of my favorite numbers")
  //Print a prompt with the solution to adding some numbers together using the add function
  fmt.Println(favorite_numbers[0], "+", favorite_numbers[2], "=", add(favorite_numbers[0], favorite_numbers[2]))

}
