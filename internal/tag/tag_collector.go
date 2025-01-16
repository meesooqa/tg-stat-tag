package tag

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type TagCollector interface {
	CollectTags(path string) ([]string, error)
}

type TagFileCollector struct {
	htmlSelector string
}

func NewTagFileCollector(htmlSelector string) *TagFileCollector {
	return &TagFileCollector{
		htmlSelector: htmlSelector,
	}
}

func (c *TagFileCollector) CollectTags(path string) []string {
	info, err := os.Stat(path)
	if err != nil {
		log.Printf("[ERROR] can't stat file: %s, error: %v", path, err)
		return nil
	}

	var tags []string

	if info.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("[ERROR] can't read dir: %s, error: %v", path, err)
			return nil
		}
		for _, entry := range dirEntries {
			if !entry.IsDir() {
				filePath := filepath.Join(path, entry.Name())
				fileTags := c.processFile(filePath)
				tags = append(tags, fileTags...)
			}
		}
	} else {
		tags = c.processFile(path)
	}

	return tags
}

func (c *TagFileCollector) processFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("[ERROR] can't open file: %s, error: %v", filePath, err)
		return nil
	}
	defer file.Close()

	var content string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		content += line
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[ERROR] can't read file: %s, error: %v", filePath, err)
			return nil
		}
	}

	return c.extractTags(content)
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
