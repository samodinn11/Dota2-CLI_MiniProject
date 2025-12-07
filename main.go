package main

import "fmt"

// HERO STRUCT
type Hero struct {
	Name  string
	Role  string
	Power int
}

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
	fmt.Println("1) Show Hero")
	fmt.Println("2) Show Version")
	fmt.Println("3) Exit")
}

// INFO ABOUT HERO (NOW USING STRUCT & STRUCT LITERAL)
func heroInfo() {
	hero := Hero{
		Name:  "Witch Doctor",
		Role:  "Support",
		Power: 0,
	}

	fmt.Println("Hero:", hero.Name)
	fmt.Println("Role:", hero.Role)
	fmt.Println("Power:", hero.Power)
}

// INFO ABOUT APP
func appInfo() {
	version := "1.0"
	patch := "Test Patch"
	fmt.Println("---------------------")
	fmt.Println("Version:", version)
	fmt.Println("Patch:", patch)
	fmt.Println("---------------------")
}

// WELCOME MESSAGE
func welcomeMessage() {
	fmt.Println("WELCOME TO DOTA2 CLI!")
	fmt.Println("---------------------")
}

// EXIT MESSAGE
func exitMessage() {
	fmt.Println("---------------------")
	fmt.Println("Thank you commander!")
	fmt.Println("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒")
	fmt.Println("▒██▄▒▒▒▒▒▒▒▀██▒▒")
	fmt.Println("▒▒▀██▄▒▒▒▒▒▒▒▀▒▒")
	fmt.Println("▒▒▒▒▀██▄▒▒▒▒▒▒▒▒")
	fmt.Println("▒▒▒▒▒▒▀██▄▄▒▒▒▒▒")
	fmt.Println("▒▒▒▒▒▒▒▀████▄▒▒▒")
	fmt.Println("▒▒█▄▒▒▒▒▒▀███▒▒▒")
	fmt.Println("▒▒▀▀▀▒▒▒▒▒▒▒▒▒▒▒")
	fmt.Println("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒")
}
