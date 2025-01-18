package format

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileFormatter_GetOutputPath(t *testing.T) {
	f := FileFormatter{path: "", ext: "html"}
	baseDir, err := os.MkdirTemp("", "test_html_files")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(baseDir)

	tests := []struct {
		input    string
		expected string
	}{
		{"file.txt", filepath.Join(baseDir, "file.html")},
		{"dir/file.md", filepath.Join(baseDir, "dir/file.html")},
		{"folder/subfolder/file", filepath.Join(baseDir, "folder/subfolder/file.html")},
		{"folder/file.with.multiple.dots", filepath.Join(baseDir, "folder/file.with.multiple.html")},
		{"/absolute/path/to/file", filepath.Join(baseDir, "absolute/path/to/file.html")},
		{"relative/path/to/folder/", filepath.Join(baseDir, "relative/path/to/folder.html")},
		{"simplefile", filepath.Join(baseDir, "simplefile.html")},
	}

	for _, test := range tests {
		result := f.getOutputPath(baseDir, test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestFileFormatter_CreateOutputPath(t *testing.T) {
	f := FileFormatter{path: "", ext: "html"}
	baseDir, err := os.MkdirTemp("", "test_html_files")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(baseDir)

	tests := []struct {
		inputPath string
	}{
		{filepath.Join(baseDir, "subdir", "file.html")},
		{filepath.Join(baseDir, "deeply", "nested", "folder", "file.html")},
		{filepath.Join(baseDir, "anotherdir", "file.html")},
	}

	for _, test := range tests {
		f.createOutputPath("", test.inputPath)

		dir := filepath.Dir(test.inputPath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory '%s' was not created", dir)
		}
	}
}
