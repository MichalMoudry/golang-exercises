package main

import "fmt"

func main() {
	testMap := map[int]string{1: "Test value", 2: "test value", 0: "zero"}
	testMap[3] = "Test value 3"
	println(testMap[3])

	_, exists := testMap[4]
	println("Does testMap[4] exist?", exists)
	testMap[4] = "four"
	_, exists = testMap[4]
	println("Does testMap[4] exist?", exists)

	println("All items in the testMap:")
	printAllItemsInMap(testMap)
	delete(testMap, 4)
	println("All items in the testMap:")
	printAllItemsInMap(testMap)

	// Assignment
	phonebook := map[string]string{"555-123": "Alice", "555-124": "Bob", "555-125": "Jean"}
	println("Welcome to your phonebook.")
	var response string
	var contactName string
	var contactNumber string
	for {
		print("Command> ")
		fmt.Scan(&response)
		if response == "quit" {
			break
		} else if response == "list" {
			for key, value := range phonebook {
				println(value, key)
			}
		} else if response == "store" {
			print("Enter contact: ")
			fmt.Scan(&contactName, &contactNumber)
			phonebook[contactNumber] = contactName
		} else if response == "lookup" {
			print("Enter a number: ")
			fmt.Scan(&contactNumber)
			var result = phonebook[contactNumber]
			if result != "" {
				println(result)
			} else {
				println("Number is not in this phonebook.")
			}
		} else {
			println("Unknown command.")
		}
	}
}

func printAllItemsInMap(input map[int]string) {
	for key, value := range input {
		println(key, value)
	}
}
