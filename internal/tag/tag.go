package tag

import (
	"sort"
)

type Service struct {
	collector TagCollector
}

type StatItem struct {
	Tag   string
	Count int
}

// NewService returns new Service instance
func NewService(c TagCollector) *Service {
	return &Service{
		collector: c,
	}
}

func (s *Service) GetStat(path string) []StatItem {
	tags := s.collector.CollectTags(path)

	// map: tag => count
	tagStatMap := make(map[string]int)
	for _, hashtag := range tags {
		tagStatMap[hashtag]++
	}

	var tagStat []StatItem

	for tag, count := range tagStatMap {
		tagStat = append(tagStat, StatItem{tag, count})
	}
	// sort by name
	sort.Slice(tagStat, func(i, j int) bool {
		return tagStat[i].Tag < tagStat[j].Tag
	})
	// sort by count
	sort.Slice(tagStat, func(i, j int) bool {
		return tagStat[i].Count > tagStat[j].Count
	})

	return tagStat
}
