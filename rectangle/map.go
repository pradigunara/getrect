package rectangle

import (
	"fmt"
	"sort"
	"strings"
)

type IntersectionMapper interface {
	Exists(key string) bool
	Add(key string, value Rectangle)
	AddSortedKey(key string, value Rectangle)
	Remove(key string) error
	Size() int
	GetSorted() []sortedIntersection
}

type intersectionMap struct {
	container map[string]Rectangle
}

type sortedIntersection struct {
	Key string
	Value Rectangle
}

func NewIntersectionMap() *intersectionMap {
	return &intersectionMap{
		container: map[string]Rectangle{},
	}
}

func (c *intersectionMap) Exists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *intersectionMap) Add(key string, value Rectangle) {
	c.container[key] = value
}

func (c *intersectionMap) AddSortedKey(key string, value Rectangle) {
	s := strings.Split(key, "")
	sort.Strings(s)
	c.Add(strings.Join(s, ""), value)
}

func (c *intersectionMap) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *intersectionMap) Size() int {
	return len(c.container)
}

func (c *intersectionMap) GetSorted() []sortedIntersection {
	sortedIntersections := []sortedIntersection{}

	for key, value := range c.container {
		sortedIntersections = append(sortedIntersections, sortedIntersection{
			Key:  key,
			Value: value,
		})
	}

	sort.Slice(sortedIntersections, func(i, j int) bool {
		return sortedIntersections[i].Key < sortedIntersections[j].Key
	})

	return sortedIntersections
}
