package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// Garden Configuration Constants
const (
	NotesDir     = "Notes"
	FilePerm     = 0644
	DirPerm      = 0755
	TimeLayout   = "15:04:05"
	DateLayout   = "2006-01-02"
)

// Markdown Template for a high-end GitHub appearance
const BeautyTemplate = `# 🌿 Seed: {{.Title}}
> **Metadata**
> - **Date:** {{.Date}}
> - **Status:** Sprouting 🪴
> - **ID:** {{.ID}}

---

### 📝 Garden Log
- *Enter your thoughts here...*

---
[⬅ Back to Garden Index](../README.md)

`

// Seed represents a single digital garden entry
type Seed struct {
	ID    int64
	Title string
	Date  string
	Path  string
}

func main() {
	fmt.Print("\033[H\033[2J") // Clear terminal screen for that "Pro App" feel
	printHeader()

	// Initialize the Garden
	if err := os.MkdirAll(NotesDir, DirPerm); err != nil {
		handleError("Failed to prepare soil (dir creation)", err)
	}

	// Create new Seed data
	newSeed := Seed{
		ID:    time.Now().Unix(),
		Title: "Experimental Growth",
		Date:  time.Now().Format(DateLayout),
	}
	newSeed.Path = filepath.Join(NotesDir, fmt.Sprintf("%s-seed.md", newSeed.Date))

	// Execute planting
	if err := plant(newSeed); err != nil {
		fmt.Printf("\n\033[33mℹ️  The garden already holds a seed for %s.\033[0m\n", newSeed.Date)
	} else {
		fmt.Printf("\n\033[32m✨ Seed successfully synthesized at: %s\033[0m\n", newSeed.Path)
	}
}

func plant(s Seed) error {
	if _, err := os.Stat(s.Path); err == nil {
		return fmt.Errorf("duplicate seed")
	}

	f, err := os.Create(s.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("garden").Parse(BeautyTemplate))
	return tmpl.Execute(f, s)
}

func printHeader() {
	fmt.Println("\033[36m" + `
   ______               __                      
  / ____/___ __________/ /__  ____  ___  _____ 
 / / __/ __ ` + "`" + `/ ___/ __  / _ \/ __ \/ _ \/ ___/ 
/ /_/ / /_/ / /  / /_/ /  __/ / / /  __/ /     
\____/\__,_/_/   \__,_/\___/_/ /_/\___/_/      
          >> DIGITAL GARDEN ENGINE v1.0 <<
	` + "\033[0m")
}

func handleError(msg string, err error) {
	fmt.Printf("\033[31m[FATAL] %s: %v\033[0m\n", msg, err)
	os.Exit(1)
}