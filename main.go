package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const (
	NotesDir   = "Notes"
	ReadmeFile = "README.md"
	HeroGIF    = "https://mir-s3-cdn-cf.behance.net/project_modules/fs/b6374879531965.5cc649b7adaf6.gif"
)

// This template builds the look of your GitHub front page
const ReadmeTemplate = `<div align="center">
  <img src="{{.Hero}}" width="160" />
  <h1>DIGITAL GARDEN</h1>
  <p>Last sync: {{.Today}} | Entries: {{.Count}}</p>
</div>

---

### 📂 Archive
| Date | Note |
| :--- | :--- |
{{range .Links}}| {{.Date}} | [View Entry]({{.Path}}) |
{{end}}

---
<p align="center">Built with Go</p>`

type NoteLink struct {
	Date string
	Path string
}

func main() {
	// 1. Setup the folder
	os.MkdirAll(NotesDir, 0755)

	// 2. Create today's note automatically
	today := time.Now().Format("2006-01-02")
	newNotePath := filepath.Join(NotesDir, today+"-seed.md")
	
	if _, err := os.Stat(newNotePath); os.IsNotExist(err) {
		content := fmt.Sprintf("# Note: %s\n\nStart writing here...", today)
		os.WriteFile(newNotePath, []byte(content), 0644)
		fmt.Println("🌱 Created today's note.")
	}

	// 3. Scan all notes to update README
	files, _ := os.ReadDir(NotesDir)
	var links []NoteLink

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".md") {
			date := strings.TrimSuffix(f.Name(), "-seed.md")
			links = append(links, NoteLink{
				Date: date,
				Path: "./Notes/" + f.Name(),
			})
		}
	}

	// 4. Update the README
	data := struct {
		Hero        string
		Today       string
		Count       int
		Links       []NoteLink
	}{
		Hero:        HeroGIF,
		Today:       today,
		Count:       len(links),
		Links:       links,
	}

	f, err := os.Create(ReadmeFile)
	if err != nil {
		fmt.Println("Error creating README:", err)
		return
	}
	defer f.Close()

	tmpl := template.Must(template.New("readme").Parse(ReadmeTemplate))
	tmpl.Execute(f, data)

	fmt.Println("✅ README updated. Garden is ready.")
}