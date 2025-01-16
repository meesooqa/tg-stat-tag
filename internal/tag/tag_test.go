package tag

import (
	"reflect"
	"testing"
)

// TagCollector mock
type MockCollector struct {
	mockTags []string
	err      error
}

func (m *MockCollector) CollectTags(_ string) ([]string, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.mockTags, nil
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

	stat, err := service.GetStat("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(stat, expected) {
		t.Errorf("Expected %v, got %v", expected, stat)
	}
}
