package main

import "fmt"

func main() {
	welcomeMessage()
	heroInfo()
	appInfo()
	exitMessage()

}

// WELCOME MESSAGE
func welcomeMessage() {
	fmt.Println("WELCOME TO DOTA2 CLI!")
	fmt.Println("---------------------")

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

func appInfo() {
	version := "1.0"
	patch := "Test Patch"
	fmt.Println("---------------------")
	fmt.Println("Version: ", version, ", Patch: ", patch)
}
func exitMessage() {
	fmt.Println("---------------------")
	fmt.Println("Thank you commander!")
}
