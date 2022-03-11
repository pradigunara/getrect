package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pradigunara/getrect/rectangle"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var validate *validator.Validate

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please specify input filename!")
	}

	inputFilename := os.Args[1]

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
					fmt.Sprintf("%d%d", i, j),
					inputRectangle.GetIntersection(otherRectangle),
				)
			}
		}
	}

	deepIntersectionMap := FindIntersectionMultiple(inputRectangles, intersectionMap)
	intersectionMap.Merge(deepIntersectionMap)

	return intersectionMap
}

func FindIntersectionMultiple(
	inputRectangles []rectangle.Rectangle,
	previousMapper rectangle.IntersectionMapper,
) rectangle.IntersectionMapper {
	currentMapper := rectangle.NewIntersectionMap()

	for i, inputRectangle := range inputRectangles {
		for j, intersectionRectangle := range previousMapper.GetContainer() {
			if strings.Contains(j, strconv.Itoa(i)) {
				continue
			}

			if intersectionRectangle.IsColliding(inputRectangle) {
				currentMapper.AddSortedKey(
					fmt.Sprintf("%d%s", i, j),
					intersectionRectangle.GetIntersection(inputRectangle),
				)
			}
		}
	}

	// go deeper levels if current level have intersection
	if currentMapper.Size() > 0 {
		 nextMapper := FindIntersectionMultiple(inputRectangles, currentMapper)
		 currentMapper.Merge(nextMapper)
	}

	return currentMapper
}

func PrintResult(inputRectangles []rectangle.Rectangle, intersections rectangle.IntersectionMapper) {
	formatRectangleNumbers := func(input string) string {
		numbers := strings.Split(input, "")
		lastNumberString := numbers[len(numbers)-1]
		bodyNumberString := numbers[:len(numbers)-1]

		lastNumber, _ := strconv.Atoi(lastNumberString)
		offsetBodyNumber := []string{}
		for _, b := range bodyNumberString {
			num, _ := strconv.Atoi(b)
			offsetBodyNumber = append(offsetBodyNumber, strconv.Itoa(num + 1))
		}

		return fmt.Sprintf("%s and %d", strings.Join(offsetBodyNumber, ", "), lastNumber + 1)
	}

	fmt.Println("Input:")
	for idx, input := range inputRectangles {
		fmt.Printf("\t%d: Rectangle at (%d,%d), w=%d, h=%d.\n", idx+1, input.X, input.Y, input.Width, input.Height)
	}

	fmt.Println("Intersections:")

	for idx, is := range intersections.GetSorted() {
		fmt.Printf(
			"\t%d: Between rectangle %s at (%d,%d), w=%d, h=%d.\n",
			idx+1, formatRectangleNumbers(is.Key), is.Value.X, is.Value.Y, is.Value.Width, is.Value.Height,
		)
	}
}

func LoadInput(inputFilename string) (rects []rectangle.Rectangle, err error) {
	fileContent, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		return rects, fmt.Errorf("error reading file: %w", err)
	}

	parseResult := struct{ Rects []rectangle.Rectangle `validate:"required,dive"` }{}

	if err = json.Unmarshal(fileContent, &parseResult); err != nil {
		return rects, fmt.Errorf("error parsing json: %w", err)
	}

	validate = validator.New()
	err = validate.Struct(parseResult)
	if err != nil {
		return rects, fmt.Errorf("error validating input: %w", err)
	}

	// limit input to 10 rectangles
	if len(parseResult.Rects) > 10 {
		return parseResult.Rects[:10], nil
	}

	return parseResult.Rects, nil
}
