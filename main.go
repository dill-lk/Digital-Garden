package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. Get today's date in a clean format (YYYY-MM-DD)
	// Fun fact: Go uses a specific date "2006-01-02" as the layout string!
	today := time.Now().Format("2006-01-02")
	fileName := today + "-seed.md"

	// 2. Check if the file already exists
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("⚠️  The seed for %s is already planted!\n", today)
		return
	}

	// 3. Create and write to the file
	content := "# Digital Garden: " + today + "\n\nLearning Go today... 🌱"
	err := os.WriteFile(fileName, []byte(content), 0644)
	
	if err != nil {
		fmt.Println("❌ Error planting seed:", err)
		return
	}

	fmt.Printf("🌱 Success! Created your new note: %s\n", fileName)
}