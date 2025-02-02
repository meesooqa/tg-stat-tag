package format

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

type PageData struct {
	Title    string
	Header   string
	TagSum   int
	CountSum int
	Items    []stat.StatItem
}

type HtmlFileFormatter struct {
	FileFormatter
}

func NewHtmlFileFormatter(outputBaseDir string, inputPath string) *HtmlFileFormatter {
	f := HtmlFileFormatter{
		FileFormatter: FileFormatter{
			ext: "html",
		},
	}
	f.path = f.createOutputPath(outputBaseDir, inputPath)

	return &f
}

func (f *HtmlFileFormatter) Format(items []stat.StatItem) {
	file, err := os.Create(f.path)
	if err != nil {
		log.Fatalf("[ERROR] file creation: %v", err)
	}
	defer file.Close()

	f.handler(file, items)
}

func (f *HtmlFileFormatter) handler(w io.Writer, items []stat.StatItem) {
	tmpl := template.Must(template.ParseFiles(
		filepath.Join("templates", "template.html"),
		filepath.Join("templates", "table.html"),
	))

	var tagSum int
	var countSum int
	for _, item := range items {
		countSum += item.Count
		tagSum++
	}

	data := PageData{
		Title:    "Hashtag list",
		Header:   "Hashtags:",
		TagSum:   tagSum,
		CountSum: countSum,
		Items:    items,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
