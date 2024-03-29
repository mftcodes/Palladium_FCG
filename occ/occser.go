package occ

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pfcg/dbservice"
	"pfcg/types"
)

func Occser(raceId int, characterId int64) int64 {
	var choice string
	var characterIdUpdated int64

	fmt.Println("Now that the base character is set up, let's see what OCC/PCC is available to choose.")

	occs, err := dbservice.GetOccsByRace(raceId)
	if err != nil {
		fmt.Printf("error getting OCCs by Race: %v \n", err)
		fmt.Errorf("occsByRace: %v", err)
		os.Exit(1)
	}

ChooseOcc:
	for _, occ := range occs {
		fmt.Printf("Press %d to see stats for %s: %s\n", occ.Id, occ.Type, occ.Desc)
	}
	fmt.Scanln(&choice)
	occId, _ := strconv.Atoi(choice)
	var occ types.Occ
	for i := range occs {
		if occs[i].Id == occId {
			occ = occs[i]
		}
	}
	choice = "y"
	fmt.Printf("You chose OCC '%s: %s', is this correct?  (Y/n): ", occ.Type, occ.Desc)
	fmt.Scanln(&choice)
	if strings.ToLower(choice) == "y" {
		characterIdUpdated, err = dbservice.SaveOcc(characterId, occId)
		if err != nil {
			fmt.Printf("error saving occ: %v \n", err)
			fmt.Errorf("occSave: %v", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("\nLet's try that again.")
		fmt.Println()
		goto ChooseOcc
	}
	return characterIdUpdated
}
