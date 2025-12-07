package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

// HERO STRUCT
type Hero struct {
	Name  string
	Role  string
	Power int
}

// SLICE of HERO
var heroes = []Hero{
	{"Anti-Mage", "Carry", 88},
	{"Juggernaut", "Carry", 90},
	{"Phantom Assassin", "Carry", 92},
	{"Spectre", "Carry", 94},
	{"Terrorblade", "Carry", 95},
	{"Invoker", "Mid", 95},
	{"Shadow Fiend", "Mid", 92},
	{"Queen of Pain", "Mid", 89},
	{"Storm Spirit", "Mid", 89},
	{"Puck", "Mid", 88},
	{"Axe", "Tank", 88},
	{"Pudge", "Tank", 90},
	{"Tidehunter", "Tank", 92},
	{"Centaur Warrunner", "Tank", 90},
	{"Bristleback", "Tank", 91},
	{"Crystal Maiden", "Support", 80},
	{"Witch Doctor", "Support", 82},
	{"Lion", "Support", 85},
	{"Rubick", "Support", 87},
	{"Dazzle", "Support", 82},
	{"Meepo", "Specialist", 97},
	{"Invoker (Quas-Wex)", "Specialist", 94},
	{"Techies", "Specialist", 89},
	{"Earthshaker", "Specialist", 87},
	{"Io", "Specialist", 91},
}

// MAP
var heroMap = make(map[string][]Hero)

var reader = bufio.NewReader(os.Stdin)

// MAIN
func main() {
	rand.Seed(time.Now().UnixNano())
	buildHeroMap()
	welcomeMessage()

	for {
		showMenu()
		choice := readInt()

		switch choice {
		case 1:
			showHeroesByRole()
		case 2:
			addHero()
		case 3:
			showStrongestHero()
		case 4:
			showAveragePower()
		case 5:
			showRoleStats()
		case 6:
			searchHero()
		case 7:
			simulateBattle()
		case 8:
			appInfo()
		case 9:
			exitMessage()
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

// -------- INPUT HELPERS --------
func readLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt() int {
	for {
		input := readLine()
		var num int
		_, err := fmt.Sscan(input, &num)
		if err == nil {
			return num
		}
		fmt.Print("Enter a number: ")
	}
}

// MENU
func showMenu() {
	fmt.Println("\n---- MENU ----")
	fmt.Println("1) Show Heroes By Role")
	fmt.Println("2) Add Hero")
	fmt.Println("3) Strongest Hero")
	fmt.Println("4) Average Power")
	fmt.Println("5) Role Statistics")
	fmt.Println("6) Search Hero By Name")
	fmt.Println("7) Simulate Battle")
	fmt.Println("8) App Info")
	fmt.Println("9) Exit")
	fmt.Print("Select option: ")
}

// MAP
func buildHeroMap() {
	heroMap = make(map[string][]Hero)
	for _, hero := range heroes {
		heroMap[hero.Role] = append(heroMap[hero.Role], hero)
	}
}

// SHOW HERO
func showHeroesByRole() {
	fmt.Println("\nAvailable Roles:")

	roles := []string{}
	for role := range heroMap {
		roles = append(roles, role)
	}
	sort.Strings(roles)

	for i, role := range roles {
		fmt.Println(i+1, ")", role)
	}

	fmt.Print("\nChoose role number: ")
	num := readInt()

	if num < 1 || num > len(roles) {
		fmt.Println("Invalid choice.")
		return
	}

	role := roles[num-1]
	list := heroMap[role]
	sortHeroesByPower(list)

	fmt.Println("\nHeroes in", role, ":")
	for i, hero := range list {
		fmt.Printf("%d) %s | Power: %d | Grade: %s\n",
			i+1,
			hero.Name,
			hero.Power,
			getPowerGrade(hero.Power),
		)
	}
}

// SORT??? PRAVILNO LI???
func sortHeroesByPower(list []Hero) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Power > list[j].Power
	})
}

// ADD HERO
func addHero() {
	fmt.Print("\nHero Name: ")
	name := readLine()

	fmt.Print("Role: ")
	role := readLine()

	fmt.Print("Power: ")
	power := readInt()

	newHero := Hero{name, role, power}
	heroes = append(heroes, newHero)
	buildHeroMap()

	fmt.Println("Hero added successfully!")
}

// STRONGEST HERO
func showStrongestHero() {
	if len(heroes) == 0 {
		fmt.Println("No heroes available.")
		return
	}

	strongest := heroes[0]
	for _, hero := range heroes {
		if hero.Power > strongest.Power {
			strongest = hero
		}
	}

	fmt.Println("\n⚔ STRONGEST HERO ⚔")
	fmt.Println("Name :", strongest.Name)
	fmt.Println("Role :", strongest.Role)
	fmt.Println("Power:", strongest.Power)
	fmt.Println("Grade:", getPowerGrade(strongest.Power))
}

// AVG PWR
func showAveragePower() {
	if len(heroes) == 0 {
		fmt.Println("No heroes available.")
		return
	}

	total := 0
	for _, hero := range heroes {
		total += hero.Power
	}
	avg := float64(total) / float64(len(heroes))
	fmt.Printf("\nAverage Power: %.2f\n", avg)
}

// ROLE STAT
func showRoleStats() {
	fmt.Println("\nROLE STATISTICS")
	for role, list := range heroMap {
		total := 0
		for _, hero := range list {
			total += hero.Power
		}
		avg := float64(total) / float64(len(list))
		fmt.Println("\nRole:", role)
		fmt.Println("Heroes:", len(list))
		fmt.Printf("Average Power: %.2f\n", avg)
	}
}

// SEARCHH
func searchHero() {
	fmt.Print("\nEnter hero name: ")
	name := strings.ToLower(readLine())

	for _, hero := range heroes {
		if strings.ToLower(hero.Name) == name {
			fmt.Println("\nHERO FOUND")
			fmt.Println("Name :", hero.Name)
			fmt.Println("Role :", hero.Role)
			fmt.Println("Power:", hero.Power)
			fmt.Println("Grade:", getPowerGrade(hero.Power))
			return
		}
	}
	fmt.Println("Hero not found.")
}

// SIM
func simulateBattle() {
	fmt.Print("\nHero 1: ")
	n1 := readLine()
	fmt.Print("Hero 2: ")
	n2 := readLine()

	h1 := findHero(n1)
	h2 := findHero(n2)

	if h1 == nil || h2 == nil {
		fmt.Println("One or both heroes not found.")
		return
	}

	boost1 := rand.Intn(21) - 10
	boost2 := rand.Intn(21) - 10

	p1 := h1.Power + boost1
	p2 := h2.Power + boost2

	fmt.Println("\n🧪 BATTLE RESULT")
	fmt.Printf("%s: %d (%+d)\n", h1.Name, p1, boost1)
	fmt.Printf("%s: %d (%+d)\n", h2.Name, p2, boost2)

	if p1 > p2 {
		fmt.Println("Winner:", h1.Name)
	} else if p2 > p1 {
		fmt.Println("Winner:", h2.Name)
	} else {
		fmt.Println("Draw!")
	}
}

// FIND
func findHero(name string) *Hero {
	name = strings.ToLower(name)
	for i := range heroes {
		if strings.ToLower(heroes[i].Name) == name {
			return &heroes[i]
		}
	}
	return nil
}

// PWR GRADE
func getPowerGrade(power int) string {
	switch {
	case power >= 95:
		return "S"
	case power >= 90:
		return "A"
	case power >= 80:
		return "B"
	case power >= 70:
		return "C"
	default:
		return "D"
	}
}

// APP INFO
func appInfo() {
	fmt.Println("\nVersion 3.1")
	fmt.Println("Features:")
	fmt.Println("- Power grades")
	fmt.Println("- Search by name")
	fmt.Println("- Battle simulator")
	fmt.Println("- Stats & averages")
	fmt.Println("- Role system with maps")
	fmt.Println("- Input validation")
	fmt.Println("MADE BY samodinn11 ON GITHUB ")
}

// MSGS
func welcomeMessage() {
	fmt.Println("WELCOME TO DOTA2 CLI!")
}

func exitMessage() {
	fmt.Println("Goodbye Commander!")
	fmt.Println("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n▒██▄▒▒▒▒▒▒▒▀██▒▒\n▒▒▀██▄▒▒▒▒▒▒▒▀▒▒\n▒▒▒▒▀██▄▒▒▒▒▒▒▒▒\n▒▒▒▒▒▒▀██▄▄▒▒▒▒▒\n▒▒▒▒▒▒▒▀████▄▒▒▒\n▒▒█▄▒▒▒▒▒▀███▒▒▒\n▒▒▀▀▀▒▒▒▒▒▒▒▒▒▒▒\n▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒")
}
