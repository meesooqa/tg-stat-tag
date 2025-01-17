package format

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

type Formatter interface {
	Format(items []stat.StatItem)
}

type PageData struct {
	Title  string
	Header string
	Items  []stat.StatItem
}

type HtmlFileFormatter struct {
	path string
}

func NewHtmlFileFormatter(path string) *HtmlFileFormatter {
	return &HtmlFileFormatter{
		path: path,
	}
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
	tmpl := template.Must(template.ParseFiles("templates/template.html"))

	data := PageData{
		Title:  "Пример списка структур",
		Header: "Мои элементы:",
		Items:  items,
	}

	tmpl.Execute(w, data)
}
