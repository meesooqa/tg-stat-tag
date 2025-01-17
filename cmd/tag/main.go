package main

import (
	"log"
	"os"

	"github.com/meesooqa/tg-stat-tag/internal/tag"
)

func main() {
	path := os.Args[1]

	collector := tag.NewTagFileCollector("div.text a")
	tagService := tag.NewService(collector)

	tgStat := tagService.GetStat(path)
	log.Println("Hashtag counts:")
	for _, item := range tgStat {
		log.Printf("%s: %d\n", item.Tag, item.Count)
	}
}
