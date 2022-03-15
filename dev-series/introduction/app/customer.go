package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {
	anna := Customer{FirstName: "Anna", LastName: "Bolika"}
	bob := Customer{FirstName: "Bob", LastName: "Smith"}
	john := Customer{FirstName: "John", LastName: "Doe"}

	customers = append(customers, anna, bob, john,
		Customer{FirstName: "Hans", LastName: "Nötig"},
		Customer{FirstName: "Peter", LastName: "Fox"},
		Customer{FirstName: "Annenmay", LastName: "Kantereit"},
		Customer{FirstName: "Test", LastName: "User1"},
		Customer{FirstName: "Test", LastName: "User2"},
		Customer{FirstName: "Test", LastName: "User3"},
		Customer{FirstName: "Test", LastName: "User4"},
	)

	return customers
}
