# Dota 2 Hero Manager


A cool command-line app to manage Dota 2 heroes! Built with Go.

## What Does It Do?

- Look at 25 different Dota 2 heroes
- See heroes by their role (Carry, Mid, Tank, Support, Specialist)
- Add your own heroes
- Find the strongest hero
- Search for heroes by name
- Battle simulator - make two heroes fight!
- See stats and averages

## How to Run It

You need Go installed on your computer.

```bash
go run main.go
```

That's it! The menu will show up.

## The Menu

When you start the app, you'll see 9 options:

**1. Show Heroes By Role** - Pick a role and see all heroes in it

**2. Add Hero** - Add a new hero with a name, role, and power number

**3. Strongest Hero** - Shows who has the highest power

**4. Average Power** - Shows the average power of all heroes

**5. Role Statistics** - See stats for each role

**6. Search Hero** - Type a hero name to find them

**7. Simulate Battle** - Pick two heroes and watch them fight! Each hero gets a random bonus or penalty, and whoever has more power wins

**8. App Info** - Shows the version and features

**9. Exit** - Quit the app

## Power Grades

Heroes get letter grades based on their power:
- **S** = 95+ (Amazing!)
- **A** = 90-94 (Really good)
- **B** = 80-89 (Good)
- **C** = 70-79 (Okay)
- **D** = Below 70 (Needs work)

## Heroes Included

The app has 25 heroes already:
- **Carries**: Anti-Mage, Juggernaut, Phantom Assassin, Spectre, Terrorblade
- **Mids**: Invoker, Shadow Fiend, Queen of Pain, Storm Spirit, Puck
- **Tanks**: Axe, Pudge, Tidehunter, Centaur, Bristleback
- **Supports**: Crystal Maiden, Witch Doctor, Lion, Rubick, Dazzle
- **Specialists**: Meepo, Invoker (Quas-Wex), Techies, Earthshaker, Io

## Tech Stack 



**Language**: Go (Golang)






<img align="right" align="center" width="200" height="200" alt="image" src="https://github.com/user-attachments/assets/b6b69664-b68d-4740-87ea-419736ab9adf" />

**What I Used**:
- `structs` - to organize hero data (name, role, power)
- `slices` - to store lists of heroes
- `maps` - to group heroes by their roles
- `sort` package - to sort heroes by power
- `bufio` - to read user input from the terminal
- `rand` - to add randomness in battles


**Why Go?** It's fast, simple, and perfect for CLI apps!

## Made By
samodinn11 on GitHub

Have fun! 🎮
