package main

import "fmt"

func main() {
	var userChoice int

	welcomeMessage()

	for {
		showMenu()
		fmt.Scan(&userChoice)
		fmt.Println("You Chose Option:", userChoice)

		switch userChoice {
		case 1:
			heroInfo()
		case 2:
			appInfo()
		case 3:
			exitMessage()
			return
		default:
			fmt.Println("Invalid option")
		}
	}

}

func showMenu() {
	fmt.Println("1) Show Hero: ")
	fmt.Println("2) Show Version: ")
	fmt.Println("3) Exit: ")
}

// INFO ABT HERO
func heroInfo() {
	name := "Witch Doctor"
	role := "Support"
	power := 0

	fmt.Println("Hero: ", name)
	fmt.Println("Role: ", role)
	fmt.Println("Power: ", power)
}

// INFO ABOUT APP
func appInfo() {
	version := "1.0"
	patch := "Test Patch"
	fmt.Println("---------------------")
	fmt.Println("Version: ", version, ", Patch: ", patch)
	fmt.Println("---------------------")
}

// WELCOME & EXIT MESSAGE
func welcomeMessage() {
	fmt.Println("WELCOME TO DOTA2 CLI!")
	fmt.Println("---------------------")

}
func exitMessage() {
	fmt.Println("---------------------")
	fmt.Println("Thank you commander!")
	fmt.Println("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n▒██▄▒▒▒▒▒▒▒▀██▒▒\n▒▒▀██▄▒▒▒▒▒▒▒▀▒▒\n▒▒▒▒▀██▄▒▒▒▒▒▒▒▒\n▒▒▒▒▒▒▀██▄▄▒▒▒▒▒\n▒▒▒▒▒▒▒▀████▄▒▒▒\n▒▒█▄▒▒▒▒▒▀███▒▒▒\n▒▒▀▀▀▒▒▒▒▒▒▒▒▒▒▒\n▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
}
