package tag

import (
	"reflect"
	"testing"
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

	expected := []StatItem{
		{Tag: "#golang", Count: 2},
		{Tag: "#абвгде", Count: 2},
		{Tag: "#example", Count: 1},
		{Tag: "#test", Count: 1},
		{Tag: "#гдеё", Count: 1},
	}

	stat := service.GetStat("")
	if !reflect.DeepEqual(stat, expected) {
		t.Errorf("Expected %v, got %v", expected, stat)
	}
}
