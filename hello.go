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
  //import random number functions from the math library
  "math/rand"
  )

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
  fmt.Println("MA random number is: ", rand.Intn(100))
}
