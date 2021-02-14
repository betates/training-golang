package main

import "fmt"

// type Item struct {
// 	Name         string
// 	Price, Stock int
// }

// type Order struct {
// 	CustomerName            string
// 	OrderNumber, TotalOrder int
// }

type Order_Detail struct {
	ItemName                             string
	OrderNumber, Quantity, PriceQuantity int
}

func MainMenu() {
	fmt.Println("===================================")
	fmt.Println("WELCOME TO SIMPLE POS APPLICATION")
	fmt.Println("===================================")
	fmt.Println("Please choose the option :")

	for a := 0; a < 2; a++ {
		if a == 1 {
			fmt.Println("2. Display All Transaction")
		} else {
			fmt.Println("1. New Transaction")
		}
	}
}

func (neworder Order_Detail) NewTransaction(customer string) {
	fmt.Println(customer, "- Item Name :", neworder.ItemName)
	fmt.Println(customer, "- Quantity :", neworder.Quantity)

}

// func Option() {
// 	press := 1
// 	if press == 1 {
// 		fmt.Println(NewTransaction())
// 	} else {
// 		fmt.Println()
// 	}
// }

func main() {
	MainMenu()

	var N int

	fmt.Scan(&N)

	Add := Order_Detail{
		ItemName: "KEYBOARD LOGITECH",
		Quantity: 2,
	}
	Add.NewTransaction("")
}
