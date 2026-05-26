package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

type Pet struct {
	Name      string `json:"name"`
	Hunger    int    `json:"hunger"`
	Happiness int    `json:"happiness"`
	Energy    int    `json:"energy"`
	Level     int    `json:"level"`
	XP        int    `json:"xp"`
}

func fixStats(pet *Pet) {
	if pet.Hunger < 0 {
		pet.Hunger = 0
	}
	if pet.Hunger > 100 {
		pet.Hunger = 100
	}

	if pet.Happiness < 0 {
		pet.Happiness = 0
	}
	if pet.Happiness > 100 {
		pet.Happiness = 100
	}

	if pet.Energy < 0 {
		pet.Energy = 0
	}
	if pet.Energy > 100 {
		pet.Energy = 100
	}

	if pet.Level <= 0 {
		pet.Level = 1
	}
}

func addXP(pet *Pet, amount int) {
	pet.XP += amount

	if pet.XP >= 100 {
		pet.Level++
		pet.XP -= 100
		fmt.Println(Green + "🎉 Level up! " + pet.Name + " is now level " + fmt.Sprint(pet.Level) + "!" + Reset)
	}
}

func savePet(pet Pet) {
	data, err := json.MarshalIndent(pet, "", "  ")
	if err != nil {
		fmt.Println(Red+"Error converting pet to JSON:"+Reset, err)
		return
	}

	err = os.WriteFile("pet.json", data, 0644)
	if err != nil {
		fmt.Println(Red+"Error saving pet:"+Reset, err)
		return
	}
}

func createNewPet() Pet {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(Cyan + "Name your pet: " + Reset)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		name = "Chunari"
	}

	pet := Pet{
		Name:      name,
		Hunger:    50,
		Happiness: 70,
		Energy:    80,
		Level:     1,
		XP:        0,
	}

	savePet(pet)
	fmt.Println(Green + "New pet created successfully 🐹" + Reset)

	return pet
}

func loadPet() Pet {
	data, err := os.ReadFile("pet.json")

	if err != nil {
		fmt.Println(Yellow + "No saved pet found. Let's create one." + Reset)
		return createNewPet()
	}

	var pet Pet

	err = json.Unmarshal(data, &pet)
	if err != nil {
		fmt.Println(Red + "Saved file is broken. Creating a new pet." + Reset)
		return createNewPet()
	}

	if pet.Name == "" {
		pet.Name = "Chunari"
	}

	if pet.Level <= 0 {
		pet.Level = 1
	}

	fixStats(&pet)

	fmt.Println(Green + "Loaded saved pet 💾" + Reset)
	return pet
}

func getMoodFace(pet Pet) string {
	if pet.Hunger >= 80 {
		return `
        /\_/\
       ( T.T )
        > ^ <
`
	} else if pet.Energy <= 20 {
		return `
        /\_/\
       ( -.- ) zZ
        > ^ <
`
	} else if pet.Happiness >= 85 {
		return `
        /\_/\
       ( ^.^ )
        > ^ <
`
	} else if pet.Happiness <= 30 {
		return `
        /\_/\
       ( ;_; )
        > ^ <
`
	} else {
		return `
        /\_/\
       ( o.o )
        > ^ <
`
	}
}

func showAsciiPet(pet Pet) {
	fmt.Println(Purple + getMoodFace(pet) + Reset)

	fmt.Println(Cyan + "   ─────────────" + Reset)
	fmt.Println(Bold + "      " + strings.ToUpper(pet.Name) + Reset)
	fmt.Println(Cyan + "   🌸🌸🌸🌸🌸🌸" + Reset)
	fmt.Println(Cyan + "   ─────────────" + Reset)
}

func makeBar(value int) string {
	totalBars := 5
	filledBars := value / 20
	emptyBars := totalBars - filledBars

	bar := ""

	for i := 0; i < filledBars; i++ {
		bar += "●"
	}

	for i := 0; i < emptyBars; i++ {
		bar += "○"
	}

	return bar
}

func showStatus(pet Pet) {
	fmt.Println(Cyan + "-----------------------" + Reset)
	fmt.Println(Bold+"Pet name:"+Reset, pet.Name)
	fmt.Println(Purple+"Level:"+Reset, pet.Level)
	fmt.Println(Purple+"XP:"+Reset, makeBar(pet.XP), pet.XP, "/ 100")
	fmt.Println(Yellow+"Hunger:"+Reset, makeBar(pet.Hunger), pet.Hunger, "/ 100")
	fmt.Println(Green+"Happiness:"+Reset, makeBar(pet.Happiness), pet.Happiness, "/ 100")
	fmt.Println(Blue+"Energy:"+Reset, makeBar(pet.Energy), pet.Energy, "/ 100")

	fmt.Println(Cyan + "Keep " + pet.Name + " alive. No excuses, pet parent. 🐹" + Reset)

	fmt.Println(Cyan + "-----------------------" + Reset)
}

func feedPet(pet *Pet) {
	pet.Hunger -= 10
	pet.Happiness += 5
	addXP(pet, 10)

	fmt.Println(Green + pet.Name + " ate food 😋" + Reset)
}

func playWithPet(pet *Pet) {
	pet.Happiness += 10
	pet.Energy -= 5
	pet.Hunger += 5
	addXP(pet, 15)

	fmt.Println(Purple + pet.Name + " played with you 🥰 🐾" + Reset)
}

func sleepPet(pet *Pet) {
	pet.Energy += 20
	pet.Hunger += 5
	addXP(pet, 8)

	fmt.Println(Blue + pet.Name + " took a nap 😴" + Reset)
}

func checkMood(pet Pet) {
	if pet.Hunger >= 80 {
		fmt.Println(Red + pet.Name + " is starving 😭 Feed the poor thing." + Reset)
	} else if pet.Energy <= 20 {
		fmt.Println(Blue + pet.Name + " is exhausted 😴 Let it sleep." + Reset)
	} else if pet.Happiness >= 85 {
		fmt.Println(Green + pet.Name + " is super happy 🥳" + Reset)
	} else if pet.Happiness <= 30 {
		fmt.Println(Yellow + pet.Name + " is sad 😔 Play with it." + Reset)
	} else {
		fmt.Println(Cyan + pet.Name + " is chilling 🐹" + Reset)
	}
}

func randomEvent(pet *Pet) {
	event := rand.Intn(5)

	if event == 0 {
		fmt.Println(Green+Reset, pet.Name, "found a cookie 🍪")
		pet.Happiness += 5
		addXP(pet, 5)
	} else if event == 1 {
		fmt.Println(Yellow+Reset, pet.Name, "got bored 😭")
		pet.Happiness -= 5
	} else if event == 2 {
		fmt.Println(Blue+Reset, pet.Name, "ran around the room 🏃")
		pet.Energy -= 10
		pet.Hunger += 5
		addXP(pet, 5)
	} else if event == 3 {
		fmt.Println(Purple+Reset, pet.Name, "found a shiny terminal bug 🐞")
		pet.Happiness += 8
		addXP(pet, 10)
	} else {
		fmt.Println(Cyan+Reset, pet.Name, "is staring at you silently 👀")
		pet.Hunger += 5
	}
}

func afterAction(pet *Pet) {
	randomEvent(pet)
	fixStats(pet)
	checkMood(*pet)
	showAsciiPet(*pet)
	showStatus(*pet)
	savePet(*pet)
}

func showMenu() {
	fmt.Println("\n" + Bold + "What do you want to do?" + Reset)
	fmt.Println("1. Show status")
	fmt.Println("2. Feed pet")
	fmt.Println("3. Play with pet")
	fmt.Println("4. Let pet sleep")
	fmt.Println("5. Save and exit")
}

func main() {
	pet := loadPet()

	fmt.Println(Bold + Cyan + "Welcome to Gotermi 🐹" + Reset)
	fmt.Println("Your terminal pet is alive.")

	showAsciiPet(pet)
	showStatus(pet)

	for {
		showMenu()

		var choice int
		fmt.Print(Cyan + "Enter choice: " + Reset)
		fmt.Scan(&choice)

		if choice == 1 {
			fixStats(&pet)
			checkMood(pet)
			showAsciiPet(pet)
			showStatus(pet)

		} else if choice == 2 {
			feedPet(&pet)
			afterAction(&pet)

		} else if choice == 3 {
			playWithPet(&pet)
			afterAction(&pet)

		} else if choice == 4 {
			sleepPet(&pet)
			afterAction(&pet)

		} else if choice == 5 {
			fixStats(&pet)
			savePet(pet)
			fmt.Println(Green + "Pet saved successfully 💾" + Reset)
			fmt.Println(Purple + "Bye bye from " + pet.Name + " 🐹" + Reset)
			break

		} else {
			fmt.Println(Red + "Invalid choice. Use brain. Pick 1-5 😭" + Reset)
		}
	}
}