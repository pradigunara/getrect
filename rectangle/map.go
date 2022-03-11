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
	GetContainer() map[string]Rectangle
	Merge(mapper IntersectionMapper)
}

type intersectionMap struct {
	container map[string]Rectangle
}

type sortedIntersection struct {
	Key   string
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
			Key:   key,
			Value: value,
		})
	}

	sort.Slice(sortedIntersections, func(i, j int) bool {
		k1, k2 := sortedIntersections[i].Key, sortedIntersections[j].Key
		l1, l2 := len(k1), len(k2)

		if l1 != l2 {
			return l1 < l2
		}

		return k1 < k2
	})

	return sortedIntersections
}

func (c *intersectionMap) GetContainer() map[string]Rectangle {
	return c.container
}

func (c *intersectionMap) Merge(otherImap IntersectionMapper) {
	for key, value := range otherImap.GetContainer() {
		c.Add(key, value)
	}
}
