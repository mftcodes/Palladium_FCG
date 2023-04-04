package main

import (
	"fmt"
	"os"
	"strconv"
)

func reviewer(choice string) {
	dbConnect()

	if choice == "1" {
		characters, err := getCharacterNames()
		if err != nil {
			fmt.Printf("error getting character names: %v \n", err)
			fmt.Errorf("characterNames: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Here is your list of characters:\n")
		for _, character := range characters {
			fmt.Printf("Press %d to see stats for %s\n", character.Id, character.Name)
		}
		fmt.Scanln(&choice)
		charId, _ := strconv.Atoi(choice)
		character, err := getCharacterById(charId)
		if err != nil {
			fmt.Printf("Character of ID %s was not found.", choice)
			// fmt.Printf("error getting character by id: %v \n", err)
			// fmt.Errorf("characterById: %v", err)
			// os.Exit(1)
		} else {
			fmt.Printf("\nHere are your characters attributes:\n")
			printCharacter(character)
		}
	} else {
		character, err := getCharacterByName(choice)
		if err != nil {
			fmt.Printf("%s was not found.", choice)
			// fmt.Printf("error getting character by name: %v \n", err)
			// fmt.Errorf("characterByName: %v", err)
			// os.Exit(1)
		} else {
			fmt.Printf("\nHere are your characters attributes:\n")
			printCharacter(character)
		}
	}
}