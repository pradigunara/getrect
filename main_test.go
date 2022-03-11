package main

import (
	"github.com/pradigunara/getrect/rectangle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindIntersections(t *testing.T) {
	input, _ := LoadInput("rectinput.json")
	input2, _ := LoadInput("rectinput2.json")

	intersections := FindIntersections(input)
	intersections2 := FindIntersections(input2)

	expectedValue := map[string]rectangle.Rectangle{
		"02":  {140, 160, 210, 20},
		"03":  {160, 140, 190, 40},
		"12":  {140, 200, 230, 60},
		"13":  {160, 200, 210, 130},
		"23":  {160, 160, 230, 100},
		"023": {160, 160, 190, 20},
		"123": {160, 200, 210, 60},
	}

	expectedValue2 := map[string]rectangle.Rectangle{
		"01":  {30, 40, 20, 10},
		"02":  {40, 40, 20, 20},
		"12":  {40, 30, 10, 20},
		"012": {40, 40, 10, 10},
	}

	assert.EqualValues(t, expectedValue, intersections.GetContainer())
	assert.EqualValues(t, expectedValue2, intersections2.GetContainer())
}
