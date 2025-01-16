package main

import (
	"log"

	"github.com/meesooqa/tg-stat-tag/internal/tag"
)

func main() {
	path := "var/data/2025-01-15-gilticus_gifs/"
	//path := "var/data/2025-01-15-gilticus_gifs/messages.html"

	collector := tag.NewTagFileCollector("div.text a")
	tagService := tag.NewService(collector)

	tgStat := tagService.GetStat(path)
	log.Println("Hashtag counts:")
	for _, item := range tgStat {
		log.Printf("%s: %d\n", item.Tag, item.Count)
	}
}
