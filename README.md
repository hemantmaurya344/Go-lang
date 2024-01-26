// package name
package main

// import statement for package fmt
import "fmt"

// main function
func main() {
	//var
	var confName = "Go Confrence"
	const confTickets = 50
	var remainingTickets uint = 50

	// to concatinate variable use ", var. Name , example down"
	// fmt.Println("Welcome to", confName, "booking application")
	// fmt.Println("TotalTickets are:", confTickets, "and remaining tickets are", remainingTickets)

	// for statement like above use printf statements which allow us to place a placeholder in place of varable but it removes new line we need to add explicitly on line using \n

	fmt.Printf("Welcome to %vbooking application.\n", confName)
	fmt.Printf("TotalTickets are: %v and remaining tickets are %v.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")

	// we can create a empty var but they need to specify datatype it is not mandatory to specify when we are providing values to var.
	var userName string
	var userTickets int

	userName = "name"
	userTickets = 2
	fmt.Printf("User %v booked %v Tickets\n", userName, userTickets)

	// Go detects the error on writing time before even compiling them while
	// we can print the datatype of variable using printf var using %T as placeholder
	fmt.Printf("Type of confTicket var is \"%T\", while confName is \"%T\" \n", confTickets, confName)

	// we know int type can hold -ve values also but when we don't want our int var to go -ve or hold -ve values we use uint, example on remainingTickets variable above

}
