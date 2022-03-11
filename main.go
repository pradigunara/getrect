package main

import (
	"encoding/json"
	"fmt"
	"github.com/pradigunara/getrect/rectangle"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//if len(os.Args) < 2 {
	//	log.Fatal("please specify input filename!")
	//}
	//
	//inputFilename := os.Args[1]
	inputFilename := "rectinput.json"

	inputRectangles, err := LoadInput(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	if len(inputRectangles) == 0 {
		log.Fatal("input is empty")
	}

	intersections := FindIntersections(inputRectangles)

	PrintResult(inputRectangles, intersections)
}

func FindIntersections(inputRectangles []rectangle.Rectangle) rectangle.IntersectionMapper {
	intersectionMap := rectangle.NewIntersectionMap()

	for i, inputRectangle := range inputRectangles {
		for j, otherRectangle := range inputRectangles {
			if i == j {
				continue
			}

			if inputRectangle.IsColliding(otherRectangle) {
				intersectionMap.AddSortedKey(
					fmt.Sprintf("%d%d", i+1, j+1),
					inputRectangle.GetIntersection(otherRectangle),
				)
			}
		}
	}

	return intersectionMap
}

func PrintResult(inputRectangles []rectangle.Rectangle, intersections rectangle.IntersectionMapper) {
	formatRectangleNumbers := func(input string) string {
		numbers := strings.Split(input, "")
		last := numbers[len(numbers)-1]
		body := numbers[:len(numbers)-1]

		return fmt.Sprintf("%s and %s", strings.Join(body, ", "), last)
	}

	fmt.Println("Input:")
	for idx, input := range inputRectangles {
		fmt.Printf("\t%d: Rectangle at (%d,%d), w=%d, h=%d.\n", idx+1, input.X, input.Y, input.Width, input.Height)
	}

	fmt.Println("Intersections:")

	for idx, is := range intersections.GetSorted() {
		fmt.Printf(
			"\t%d: Between rectangle %s at (%d,%d), w=%d, h=%d.\n",
			idx + 1, formatRectangleNumbers(is.Key), is.Value.X, is.Value.Y, is.Value.Width, is.Value.Height,
		)
	}
}

func LoadInput(inputFilename string) (rects []rectangle.Rectangle, err error) {
	fileContent, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		return rects, fmt.Errorf("error reading file: %w", err)
	}

	parseResult := struct{ Rects []rectangle.Rectangle }{}

	if err = json.Unmarshal(fileContent, &parseResult); err != nil {
		return rects, fmt.Errorf("error parsing json: %w", err)
	}

	// limit input to 10 rectangles
	if len(parseResult.Rects) > 10 {
		return parseResult.Rects[:10], nil
	}

	return parseResult.Rects, nil
}
