package tag

import (
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Service struct {
	ServiceParams
}

type ServiceParams struct {
	HtmlSelector string
}

type StatItem struct {
	Tag   string
	Count int
}

// NewService returns new Service instance
func NewService(p ServiceParams) *Service {
	return &Service{
		ServiceParams: p,
	}
}

func (s *Service) GetStat() (tagStat []StatItem) {
	tags := s.collectTags()

	tagStatMap := make(map[string]int)
	for _, hashtag := range tags {
		tagStatMap[hashtag]++
	}

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

func (s *Service) collectTags() (tags []string) {
	//files := []string{
	//	"var/data/2025-01-15-gilticus_gifs/messages.html",
	//	"var/data/2025-01-15-gilticus_gifs/messages2.html",
	//	"var/data/2025-01-15-gilticus_gifs/messages3.html",
	//}

	//path := "var/data/2025-01-15-gilticus_gifs/messages.html"
	//fi, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer func() { err = fi.Close() }()

	// TODO implement the method
	tags = []string{
		"asdf",
		"bsdfasdf",
		"csdfasdfasdf",
		"asdf",
		"dsdfasdfasdfasdf",
		"dsdfasdfasdfasdf",
		"esdfasdfasdfasdfasdf",
		"hsdfasdfasdfasdfasdfasdf",
		"dsdfasdfasdfasdf",
		"ыаы",
		"аыа",
		"ыаы",
		"123",
	}
	return tags
}

// returns list of tags from the Telegram archived HTML
func (s *Service) extractTags(messagesHTML string) (tags []string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(messagesHTML))
	if err != nil {
		log.Printf("[ERROR] can't parse messagesHTML to parse tags: %q, error: %v", messagesHTML, err)
		return nil
	}

	re := regexp.MustCompile(`#[a-zA-Zа-яА-ЯёЁ0-9]+`)
	doc.Find(s.ServiceParams.HtmlSelector).Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		matches := re.FindAllString(text, -1)
		tags = append(tags, matches...)
	})

	return tags
}
