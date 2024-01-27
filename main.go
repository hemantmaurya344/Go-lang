// package name
// package main

// // import statement for package fmt
// import "fmt"   // fmt stands for format
// GO has multiple predefined packages having different functions for different working specified on github of GO
// To import multiple packages we need to use
// import(
// 	"pack1"
// 	"pack2"
// )

//   	Basic Syntax of Hello world

// package main
// import "fmt"
// func main() {
//		fmt.println("Hello World")
// }

//  function definition
// func main() {

// 	var definiton
// 	var confName = "Go Confrence"

// constant definiton
// 	const confTickets = 50

// 	var remainingTickets uint = 50

//  }

//  we have another syntax for creating var is name := "value" this is not valid for constants and where we specify the type, here we don't need to specify the var
// eg name := "value"

// 	to concatinate variable use ", var. Name , example down"
// 	fmt.Println("Welcome to", confName, "booking application")
// 	fmt.Println("TotalTickets are:", confTickets, "and", remainingTickets)

// 	for statement like above use printf statements which allow us to place a placeholder in place of varable but it removes new line we need to add explicitly on line using \n

// 	fmt.Printf("Welcome to %vbooking application.\n", confName)
// 	fmt.Printf("TotalTickets are: %v and remaining tickets are %v.\n", confTickets, remainingTickets)

// 	fmt.Println("Get your tickets to attend.")

// 	 we can create a empty var but they need to specify datatype it is not mandatory to specify when we are providing values to var.

// 	var userName string
// 	var userTickets int

// 	userName = "name"
// 	userTickets = 2
// 	fmt.Printf("User %v booked %v Tickets\n", userName, userTickets)

// 	Go detects the error on writing time before even compiling them while
// 	we can print the datatype of variable using printf var using %T as placeholder

// eg: 	fmt.Printf("Type of confTicket var is \"%T\", while confName is \"%T\" \n", confTickets, confName)

// 	 we know int type can hold -ve values also but when we don't want our int var to go -ve or hold -ve values we use "uint", example on remainingTickets variable above

// to read user input we use "scan" function from same fmt package, just same as print function Scan also takes message, parameter for scan function is user input
// eg fmt.Scan(userName)

// we can create var and take refrence on that variable but when code run it doesn't wait for input for user to fix that we use "pointer"

// pointer: Pointers is variable that points to memory address of another variable that refrences the actual value
// pointers in go lang also called special variable
// syntax  &pointer
// eg
// fmt.Println(remainingTickets)   50
// fmt.Println(&remainingTickets)   0xc000012000 (memory address)

// Scan() can't process after space in input and exits the program

// Typecasting in GO is done by T(v) where T is type and v is value/variable

// 	Functions
//	functions are defined using func keyword in go
//  syntax:   func <functionName>() {  }
// 		eg:  func <Hello>() {  }

// when functions take some var and return some var
// syntax: func <function-name>(<var-name> typeOfVar) return-type { return var/something }
//  eg: func Hello(name string) string { return name }

// "go mod tidy" in terminal will sync the import packages
// eg: file dir (greeting/greetings.go)
// package greetings
// import "fmt"
// func Hello(name string) string {
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message
// }

//  file dir (hello/main.go)
// package main
// import (
//     "fmt"
//     "example.com/greetings"
// )
// func main() {
//     message := greetings.Hello("Gladys")
//     fmt.Println(message)
// }

// if you had pushed you module to repository GO directly downloads from there but it is in local setup we need to adapt the hello module to adapt greeting module for this
// "go mod edit -replace example.com/greetings/greeting=../greetings"
// then use go mod tidy to add the dependency

//	Arrays & Slices
// Arrays used to fixed size while slice is for dynamic size and increase size according to inputs.
// Slices uses Arrays under the hood.

// Array eq: var name = [50] int{2,3,4,5.....}
// we can have empty array like {} or we can specify values eq var name = [50]string{}
// for empty array we can use syntax like this "var name [40]int"

// for slices we don't use index we use append() to append and for retrieving we use the indexes.
// eg: var name = []string  //slice
// or var name []string{}
// or name := []string{}
//  eg: name = append(name, value)

// Loops
// In go we don't have while, do while, for each loop we have only for loop
// for loop has different types in GO
// 1. Infinite:  syntax:  for{ .... }
// 2. for index, var-Name := targetArray {  ...  }
//  inplace of index we can use '_' in go _ used to make unused variables

// Conditionals
// syntax: if conditon { ... }

// Swtich
// syntax : switch var {
// 	case "sad":  do something
// 	case "sac":  do something
// 	case "sacd":   do something
// }

// To export user created function of a package first import the package where you want to use with 'moduleName/packageName' then make first letter of function capital

// we can export variables,constants also by the same way making first letter capital

// Maps
// It maps unique keys to values
// syntax: var mapName = make(map[typeOfKey]typeOfValue)  make here give us empty map
// access the map using: mapName["key"]=value

// Struct
// Struct is like classes in java
// syntax:
// type nameOfStruct struct {
// 	var string
//  var bool
// }

// To Implement concurrency in Go we know that some task take time to process so we break out that task from current thread and execute it on its own thread.
// to make the new thread we need to use 'go' keyword in front of task which execute it on new thread
// Specifing the 'go' keyword is only takes to spin a new thread while GOLANG do its own thing to spin new thread on background
// eg: printf(asdfasd)
// 	   go run(afdsd) { ..... }

// when we need to wait for any thread to complete its execution before exiting the main application thread we need to tell main thread to wait for completion of new thread
// we can do this using 'waitgroup' which is from sync package
// eg: import "sync"
// 	   var wg = sync.WaitGroup{}
//  now we can use wg var befor where we are spining new thread
// eg:  wg.Add(2)
//      go run(Asdsa){...}
//      go dotast(){ ... }

// we can wait for thread to complete using wait function also from sync package which we need to use at end of the main thread
// eg:
//    func main(){
// 	      ..
// 	      ..
// 	      ..
// 	      wg.Wait()
//    }

// we need to add Done function at end of new thread we are creating  'wg.Done()'  done function remove the thread from waiting list which we define above

package main

import "fmt"

func main() {

	confName := "Go Confrence"
	const confTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %vbooking application.\n", confName)
	fmt.Printf("TotalTickets are: %v and remaining tickets are %v.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")

	bookings := []string{} //slice

	for {
		var userName string
		var userTickets uint
		var userEmail string

		fmt.Printf("Enter your Name: ")
		fmt.Scan(&userName)

		fmt.Printf("Enter your email:")
		fmt.Scan(&userEmail)

		fmt.Printf("Enter No. of Tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, userName)

			fmt.Printf("Thank You %v for booking %v Tickets. you will recieve confirmation and copy of ticket at %v\n", userName, userTickets, userEmail)
			fmt.Printf("Remaining %v tickets for confrence\n", remainingTickets)

			fmt.Printf("Users who booked tickets are %v\n", bookings)

			if remainingTickets == 0 {
				fmt.Println("All tickets are booked")
				break
			}
		} else {
			fmt.Printf("we have %v remaining tickets your ticket amount %v cannot exceed it \n", remainingTickets, userTickets)
		}
	}
}
