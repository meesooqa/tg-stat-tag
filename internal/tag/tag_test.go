package tag

import (
	"reflect"
	"testing"

	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

// TagCollector mock
type MockCollector struct {
	mockTags []string
}

func (m *MockCollector) CollectTags(_ string) []string {
	return m.mockTags
}

func TestService_GetStat(t *testing.T) {
	mockTags := []string{"#golang", "#абвгде", "#golang", "#test", "#гдеё", "#example", "#абвгде"}
	mockCollector := &MockCollector{mockTags: mockTags}

	service := NewService(mockCollector)

	expected := []stat.StatItem{
		{Tag: "#golang", Count: 2},
		{Tag: "#абвгде", Count: 2},
		{Tag: "#example", Count: 1},
		{Tag: "#test", Count: 1},
		{Tag: "#гдеё", Count: 1},
	}

	tagStat := service.GetStat("")
	if !reflect.DeepEqual(tagStat, expected) {
		t.Errorf("Expected %v, got %v", expected, tagStat)
	}
}
