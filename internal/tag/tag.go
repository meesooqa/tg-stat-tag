package tag

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TagCollector interface {
	CollectTags(path string) ([]string, error)
}

type TagFileCollector struct {
	htmlSelector string
}

type Service struct {
	collector TagCollector
}

type StatItem struct {
	Tag   string
	Count int
}

func NewTagFileCollector(htmlSelector string) *TagFileCollector {
	return &TagFileCollector{
		htmlSelector: htmlSelector,
	}
}

// NewService returns new Service instance
func NewService(c TagCollector) *Service {
	return &Service{
		collector: c,
	}
}

func (s *Service) GetStat(path string) ([]StatItem, error) {
	tags, err := s.collector.CollectTags(path)
	if err != nil {
		return nil, err
	}

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

	return tagStat, nil
}

func (c *TagFileCollector) collectTags(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var tags []string

	if info.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, entry := range dirEntries {
			if !entry.IsDir() {
				filePath := filepath.Join(path, entry.Name())
				fileTags, err := c.processFile(filePath)
				if err != nil {
					return nil, err
				}
				tags = append(tags, fileTags...)
			}
		}
	} else {
		tags, err = c.processFile(path)
		if err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func (c *TagFileCollector) processFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() { err = file.Close() }()

	var content string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		content += line
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return c.extractTags(content), nil
}

// returns list of tags from the Telegram archived HTML
func (c *TagFileCollector) extractTags(messagesHTML string) (tags []string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(messagesHTML))
	if err != nil {
		log.Printf("[ERROR] can't parse messagesHTML to parse tags: %q, error: %v", messagesHTML, err)
		return nil
	}

	re := regexp.MustCompile(`#[a-zA-Zа-яА-ЯёЁ0-9]+`)
	doc.Find(c.htmlSelector).Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		matches := re.FindAllString(text, -1)
		tags = append(tags, matches...)
	})

	return tags
}
