package set

import (
	"fmt"
)

//MakeSet initialize the set
func MakeIntegerSet() *integerSet {
	return &integerSet{
		container: make(map[int]struct{}),
	}
}

type integerSet struct {
	container map[int]struct{}
}

func (c *integerSet) Exists(key int) bool {
	_, exists := c.container[key]
	return exists
}

func (c *integerSet) Add(key int) {
	c.container[key] = struct{}{}
}

func (c *integerSet) Remove(key int) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *integerSet) Size() int {
	return len(c.container)
}
