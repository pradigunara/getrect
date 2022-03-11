package rectangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectangle_IsColliding(t *testing.T) {
	type testCase struct {
		firstRect     Rectangle
		secondRect    Rectangle
		shouldCollide bool
	}

	tests := []testCase{
		{
			firstRect:     Rectangle{100, 100, 250, 80},
			secondRect:    Rectangle{140, 160, 250, 100},
			shouldCollide: true,
		}, {
			firstRect:     Rectangle{120, 200, 250, 150},
			secondRect:    Rectangle{160, 140, 350, 190},
			shouldCollide: true,
		}, {
			firstRect:     Rectangle{100, 100, 250, 80},
			secondRect:    Rectangle{120, 200, 250, 150},
			shouldCollide: false,
		},
	}

	for _, tt := range tests {
		result := tt.firstRect.IsColliding(tt.secondRect)

		assert.Equal(t, tt.shouldCollide, result)
	}
}

func TestRectangle_GetIntersection(t *testing.T) {
	type testCase struct {
		firstRect          Rectangle
		secondRect         Rectangle
		intersectionResult Rectangle
	}

	tests := []testCase{
		{
			firstRect:     Rectangle{100, 100, 250, 80},
			secondRect:    Rectangle{140, 160, 250, 100},
			intersectionResult: Rectangle{140, 160, 210, 20},
		}, {
			firstRect:     Rectangle{120, 200, 250, 150},
			secondRect:    Rectangle{160, 140, 350, 190},
			intersectionResult: Rectangle{160, 200, 210, 130},
		},
	}

	for _, tt := range tests {
		result := tt.firstRect.GetIntersection(tt.secondRect)

		assert.Equal(t, tt.intersectionResult.X, result.X)
		assert.Equal(t, tt.intersectionResult.Y, result.Y)
		assert.Equal(t, tt.intersectionResult.Width, result.Width)
		assert.Equal(t, tt.intersectionResult.Height, result.Height)
	}
}
