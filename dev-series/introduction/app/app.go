package main

import (
	"fmt"
	// "time"
)

// global vars (must be with var! and not := )
// var test = "blah"

func main() {

	customers := GetCustomers()

	for _, customer := range customers {
		// we can access the "customer" varaible in this approach
		fmt.Println(customer)
	}

	// fmt.Println("Hello World!")

	/*
		customers := getData()
		fmt.Println(customers)

		fmt.Println(len(customers))

	*/
	/*
		customer = getData(2)
		fmt.Println(customer)

		customer = getData(3)
		fmt.Println(customer)
	*/

	// fmt.Println(test)
}

func getData() (customers []string) {

	// Arrays always need to be a fixed size

	// long version
	//var firstName = "Anna"
	// short version
	// lastName := "Bolika"

	/*
		var firstName string
		var lastName string

		if customerId == 1 {
			firstName = "Anna"
			lastName = "Bolika"
		} else if customerId == 2 {
			firstName = "Hans"
			lastName = "Nötig"
		}

		return firstName + " " + lastName
	*/

	/*
		customer := "Anna Bolika"

		customers[0] = customer
		customers[1] = "Hans Nötig"

	*/

	/*
		// instance and append slice item
		customers = []string
		customers = append(customers, "Anna Bolika")
	*/

	// faster way
	customers = []string{"Anna Bolika", "Hans Nötig", "Max VonHinten"}

	customers = append(customers, "Max Muster")
	customers = append(customers, "Test User1")
	customers = append(customers, "Test User2")
	customers = append(customers, "Test User3")
	customers = append(customers, "Test User4")
	customers = append(customers, "Test User5")
	customers = append(customers, "Test User6")

	/*
		for {
			fmt.Println("Infinite Loop 1")
			time.Sleep(time.Second)
			// reak
		}
	*/

	/*
		for x := 0; x < 10; x++ {
			// any code in here will run 10 times! (unless we break!)
			fmt.Println(customers[x])
		}

	*/

	/*
		for x, customer := range customers {
			customer = customers[x]
			fmt.Println(customer)
		}
	*/

	// silence x and make it more efficient
	for _, customer := range customers {
		fmt.Println(customer)
	}

	return customers
}
