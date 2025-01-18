package helpers

import (
	"reflect"
	"testing"
)

func TestFlattenMaps(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]string
		expected map[string]string
	}{
		{
			name: "Single map",
			input: []map[string]string{
				{"key1": "value1", "key2": "value2"},
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Multiple maps with unique keys",
			input: []map[string]string{
				{"key1": "value1"},
				{"key2": "value2"},
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Overlapping keys with override",
			input: []map[string]string{
				{"key1": "value1", "key2": "value2"},
				{"key2": "override_value2", "key3": "value3"},
			},
			expected: map[string]string{
				"key1": "value1",
				"key2": "override_value2",
				"key3": "value3",
			},
		},
		{
			name: "Empty input",
			input: []map[string]string{
				{},
			},
			expected: map[string]string{},
		},
		{
			name:     "No maps in slice",
			input:    []map[string]string{},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlattenMaps(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("flattenMaps() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
