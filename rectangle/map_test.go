package rectangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionMap_GetSorted(t *testing.T) {
	type testCase struct {
		input    intersectionMap
		expected []sortedIntersection
	}

	tests := []testCase{
		{
			input: intersectionMap{container: map[string]Rectangle{
				"123": {},
				"45":  {},
				"9":   {},
				"0":   {},
			}},
			expected: []sortedIntersection{{Key: "0"}, {Key: "9"}, {Key: "45"}, {Key: "123"}},
		},
		{
			input: intersectionMap{container: map[string]Rectangle{
				"567": {},
				"135": {},
				"678": {},
				"456": {},
			}},
			expected: []sortedIntersection{{Key: "135"}, {Key: "456"}, {Key: "567"}, {Key: "678"}},
		},
	}

	for _, tt := range tests {
		assert.EqualValues(t, tt.expected, tt.input.GetSorted())
	}
}

func TestIntersectionMap_Merge(t *testing.T) {
	type testCase struct {
		mapper1    intersectionMap
		mapper2    intersectionMap
		expected   intersectionMap
	}

	tests := []testCase{
		{
			mapper1: intersectionMap{container: map[string]Rectangle{
				"01": {100, 100, 250, 80},
				"12":  {140, 160, 250, 100},
			}},
			mapper2: intersectionMap{container: map[string]Rectangle{
				"123": {100, 100, 250, 80},
				"456":  {140, 160, 250, 100},
			}},
			expected: intersectionMap{container: map[string]Rectangle{
				"01": {100, 100, 250, 80},
				"12":  {140, 160, 250, 100},
				"123": {100, 100, 250, 80},
				"456":  {140, 160, 250, 100},
			}},
		},
		{
			mapper1: intersectionMap{container: map[string]Rectangle{
				"4567": {100, 100, 250, 80},
				"7890":  {140, 160, 250, 100},
				"3456":  {120, 200, 250, 150},
			}},
			mapper2: intersectionMap{container: map[string]Rectangle{
				"1234": {100, 100, 250, 80},
				"5678":  {140, 160, 250, 100},
			}},
			expected: intersectionMap{container: map[string]Rectangle{
				"4567": {100, 100, 250, 80},
				"7890":  {140, 160, 250, 100},
				"3456":  {120, 200, 250, 150},
				"1234": {100, 100, 250, 80},
				"5678":  {140, 160, 250, 100},
			}},
		},
	}

	for _, tt := range tests {
		tt.mapper1.Merge(&tt.mapper2)
		assert.EqualValues(t, tt.expected, tt.mapper1)
	}
}
