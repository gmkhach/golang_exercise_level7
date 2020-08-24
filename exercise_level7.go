package main

import "fmt"

type person struct {
    first   string
    last    string
    age     int
}

func main() {
    /*
        Exercise 1
        1. Create a value and assign it to a variable.
        2. Print the address of that value.
    */
    a := 42
    fmt.Println(&a)

    /*
        Exercise 2
        1. Create a person struct.
        2. Create a func called "changeMe" with a *person as a parameter
        3. Inside "changeMe" change the value stored at the *person address
        4. Create a value of type person
        5. Print out the value
        6. Call "changeMe"
    */
    p := person {
        first: "James",
        last: "Bond",
        age: 32,
    }
    fmt.Println(p)
    changeMe(&p) 
    fmt.Println(p)

}

func changeMe(p *person) {
    (*p).first = "Miss"
    (*p).last  = "Moneypenny"
    (*p).age = 27
}