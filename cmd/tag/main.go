package main

import (
	"github.com/meesooqa/tg-stat-tag/internal/format"
	"github.com/meesooqa/tg-stat-tag/internal/tag"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	outputBaseDir := "var/output"

	inputPath := os.Args[1]
	/*
		outputPath := createOutputPath(outputBaseDir, inputPath)

		collector := tag.NewTagFileCollector("div.text a")
		tagService := tag.NewService(collector)

		items := tagService.GetStat(inputPath)

		f := format.NewHtmlFileFormatter(outputPath)
		f.Format(items)
	*/
	collector := tag.NewTagFileCollector("div.text a")
	tagService := tag.NewService(collector)

	items := tagService.GetStat(inputPath)

	f := format.NewCsvFileFormatter(outputBaseDir, inputPath)
	f.Format(items)
}

func createOutputPath(baseDir string, inputPath string) string {
	outputPath := getOutputPath(baseDir, inputPath)

	dir := filepath.Dir(outputPath)
	os.MkdirAll(dir, os.ModePerm)

	return outputPath
}

func getOutputPath(baseDir string, inputPath string) string {
	cleanPath := filepath.Clean(inputPath)
	ext := filepath.Ext(cleanPath)
	baseName := strings.TrimSuffix(cleanPath, ext)

	return filepath.Join(baseDir, baseName+".html")
}
