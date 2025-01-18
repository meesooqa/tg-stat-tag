package format

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

type CsvFileFormatter struct {
	FileFormatter
}

func NewCsvFileFormatter(outputBaseDir string, inputPath string) *CsvFileFormatter {
	f := CsvFileFormatter{
		FileFormatter: FileFormatter{
			ext: "csv",
		},
	}
	f.path = f.createOutputPath(outputBaseDir, inputPath)

	return &f
}

func (f *CsvFileFormatter) Format(items []stat.StatItem) {
	file, err := os.Create(f.path)
	if err != nil {
		log.Fatalf("[ERROR] file creation: %v", err)
	}
	defer file.Close()

	f.handler(file, items)
}

func (f *CsvFileFormatter) handler(w io.Writer, items []stat.StatItem) {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	err := writer.Write([]string{"Tag", "Count"})
	if err != nil {
		log.Fatalf("[ERROR] header writing: %v", err)
	}

	for _, item := range items {
		record := []string{item.Tag, strconv.Itoa(item.Count)}
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("[ERROR] item writing: %v", err)
		}
	}
}
