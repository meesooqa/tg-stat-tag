package main

import (
	"os"

	"github.com/meesooqa/tg-stat-tag/internal/format"
	"github.com/meesooqa/tg-stat-tag/internal/tag"
)

func main() {
	outputBaseDir := "var/output"
	inputPath := os.Args[1]

	collector := tag.NewTagFileCollector("div.text a")
	tagService := tag.NewService(collector)

	items := tagService.GetStat(inputPath)

	//f := format.NewCsvFileFormatter(outputBaseDir, inputPath)
	f := format.NewHtmlFileFormatter(outputBaseDir, inputPath)
	f.Format(items)
}
