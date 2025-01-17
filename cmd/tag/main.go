package main

import (
	"github.com/meesooqa/tg-stat-tag/internal/format"
	"os"

	"github.com/meesooqa/tg-stat-tag/internal/tag"
)

func main() {
	path := os.Args[1]

	collector := tag.NewTagFileCollector("div.text a")
	tagService := tag.NewService(collector)

	items := tagService.GetStat(path)

	f := format.NewHtmlFileFormatter("var/output/sdafsa.html")
	f.Format(items)
}
