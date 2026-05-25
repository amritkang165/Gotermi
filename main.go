package main

import "fmt"

type Pet struct{
	name      string
	hunger    int
	happiness int
	energy    int
}

func showStatus(pet Pet) {
	fmt.Println("-----------------------")
	fmt.Println("Pet name:", pet.name)
	fmt.Println("Hunger:", pet.hunger)
	fmt.Println("Happiness:", pet.happiness)
	fmt.Println("Energy:", pet.energy)
	fmt.Println("------------------------")

}

func feedPet(pet *Pet){
	pet.hunger -= 10
	pet.happiness += 5
	fmt.Println(pet.name , "ate food 😋")
}

func platWithPet(pet *Pet){
	pet.happiness += 10
	pet.energy -= 5
	fmt.Println(pet.name , "played with you 🥰 🐾")
}

func sleepPet(pet *Pet){
	pet.energy += 20
	pet.hunger += 5
	fmt.Println(pet.name , "took a nap 😴")
}

func main() {
	pet := Pet{
		name : "Chunari",
		hunger : 50,
		happiness : 70,
		energy : 80,
	}
	fmt.Println("Welcome to Gotermi 🐹")
	fmt.Println("Your terminal pet is alive.")

	showStatus(pet)

	feedPet(&pet)

	showStatus(pet)

}