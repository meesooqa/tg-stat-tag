package tag

import (
	"os"
	"reflect"
	"testing"
)

func TestTagFileCollector_CollectTags_File(t *testing.T) {
	collector := &TagFileCollector{htmlSelector: "div"}

	tempFile, err := os.CreateTemp("", "testfile*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() { err = os.Remove(tempFile.Name()) }()

	content := `<div class="message">#golang #test</div>`
	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	tags := collector.CollectTags(tempFile.Name())

	expected := []string{"#golang", "#test"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Expected %v, got %v", expected, tags)
	}
}

func TestTagFileCollector_processFile(t *testing.T) {
	collector := &TagFileCollector{htmlSelector: "div"}

	tempFile, err := os.CreateTemp("", "testfile*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() { err = os.Remove(tempFile.Name()) }()

	content := `<div class="message">#golang #test</div>`
	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	tags := collector.processFile(tempFile.Name())

	expected := []string{"#golang", "#test"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Expected %v, got %v", expected, tags)
	}
}

func TestTagFileCollector_extractTags(t *testing.T) {
	tests := []struct {
		name         string
		messagesHTML string
		expectedTags []string
		expectError  bool
	}{
		{
			name: "Valid HTML with tags",
			messagesHTML: `
				<div>some text #tag1 more text #tag2</div>
				<div>other text <a href="#">#tag3</a></div>
				<p>more text #tag4</p>
			`,
			expectedTags: []string{"#tag1", "#tag2", "#tag3"},
			expectError:  false,
		},
		{
			name: "HTML with no tags",
			messagesHTML: `
				<div>some text here</div>
				<p>other content</p>
			`,
			expectedTags: []string{},
			expectError:  false,
		},
		//{
		//	name: "HTML with invalid tag format",
		//	messagesHTML: `
		//		<div>some text #123abc</div>
		//		<div>#invalidTag!@#</div>
		//	`,
		//	expectedTags: []string{"#123abc"},
		//	expectError:  false,
		//},
		{
			name: "HTML with Russian tags",
			messagesHTML: `
				<div>some text #тег1</div>
				<div>#тег2 more text</div>
			`,
			expectedTags: []string{"#тег1", "#тег2"},
			expectError:  false,
		},
		{
			name:         "Invalid HTML",
			messagesHTML: `invalid html <div>#tag1</div>`,
			expectedTags: []string{"#tag1"},
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := &TagFileCollector{htmlSelector: "div"}
			tags := collector.extractTags(tt.messagesHTML)
			if len(tags) != len(tt.expectedTags) {
				t.Errorf("len: expected %v, got %v", tt.expectedTags, tags)
			}
			for i, tag := range tags {
				if tag != tt.expectedTags[i] {
					t.Errorf("list: expected tag %v, got %v", tt.expectedTags[i], tag)
				}
			}
		})
	}
}
