package format

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

type Formatter interface {
	Format(items []stat.StatItem)
}

type FileFormatter struct {
	path string
	ext  string
}

func (f *FileFormatter) createOutputPath(baseDir string, inputPath string) string {
	outputPath := f.getOutputPath(baseDir, inputPath)

	dir := filepath.Dir(outputPath)
	os.MkdirAll(dir, os.ModePerm)

	return outputPath
}

func (f *FileFormatter) getOutputPath(baseDir string, inputPath string) string {
	cleanPath := filepath.Clean(inputPath)
	ext := filepath.Ext(cleanPath)
	baseName := strings.TrimSuffix(cleanPath, ext)

	return filepath.Join(baseDir, baseName+"."+f.ext)
}
