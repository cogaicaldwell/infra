package cli

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"coginfra/storage"
	"coginfra/utils"
)

func StartMenu(s storage.Storage) {
	utils.Logger.Print("\033[H\033[2J")
	utils.Logger.Println("Application Started Successfully.")
	genAPIKey := "Generate a new API Key"
	clearScreen := "Clear the logs"

	for {
		prompt := promptui.Select{
			Label: "Choose an opts",
			Items: []string{
				genAPIKey, clearScreen,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			utils.Logger.Printf("Prompt failed %v\n", err)
			return
		}

		if result == genAPIKey {
			handleGenAPIKey(s)
		} else {
			handleCLS()
		}

	}
}

func handleGenAPIKey(store storage.Storage) {
	utils.Logger.Println("[info] Generating a new api key ... ")
	key, err := store.CreateAPIKey()
	if err != nil {
		utils.Logger.Println("[error] Failed to create a new api key. Details: \n", err)
	}
	// utils.Logger.Println("[info] Here is your API Key. KEEP IT SECRET \n\t ", key)

	blue := color.New(color.FgHiBlue)
	boldRed := blue.Add(color.Bold)
	boldRed.Println("[info] API Key generated successfully \n \t ", key)
}

func handleCLS() {
	utils.Logger.Print("\033[H\033[2J")
}
