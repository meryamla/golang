package main

import (
	"fmt"
	"os"
)

func"fmt"() {
	// Define the known colors and their corresponding texts
	colorTexts := map[string]string{
		"Blauw": "zoals de lucht.",
		"Rood":  "met passie.",
		"Geel":  "als de stralen van de zon.",
		"Groen": "van de natuur.",
		// Add three additional colors and their texts
		"Paars":  "met mysterie.",
		"Oranje": "als de gloed van de avondzon.",
		"Roze":   "zoals bloesems in de lente.",
	}

	// Get user input for favorite color
	fmt.Print("Voer je lievelingskleur in: ")
	var favoriteColor string
	fmt.Scanln(&favoriteColor)

	// Check if the chosen color exists in the map
	text, exists := colorTexts[favoriteColor]
	if !exists {
		fmt.Println("Foutmelding: Ongeldige kleur opgegeven. Het programma wordt afgesloten.")
		os.Exit(1)
	}

	// Write the text to a text file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Foutmelding: Kan geen tekstbestand maken. Het programma wordt afgesloten.")
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(favoriteColor + " " + text)
	if err != nil {
		fmt.Println("Foutmelding: Kan geen tekst naar het bestand schrijven. Het programma wordt afgesloten.")
		os.Exit(1)
	}

	fmt.Println("Tekst is succesvol naar het bestand geschreven.")
}
