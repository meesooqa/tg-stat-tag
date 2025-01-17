package stat

type StatItem struct {
	Tag   string
	Count int
}

func NewStatItem(tag string, count int) *StatItem {
	return &StatItem{
		Tag:   tag,
		Count: count,
	}
}
