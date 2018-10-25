/*
- Use interfaces anywhere you would want to use generics
- Introduce a new builtin type named "generic"
- Create contracts by creating functions on the "generic" type
*/

package main

// SCENARIO: single sum function
// SOLUTION: use inferred types
func Sum(values ...infer) (ret infer) {
	for _, v := range values {
		ret += v
	}
	return ret
}

// SCENARIO: a single sort function
// SOLUTION: use generic contracts
func (g generic) Len() int {
	return len(g)
}

func (g generic) Less(i, j int) bool {
	return g[i] < g[j]
}

func (g generic) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// SOLUTION: use infer type
func Sort(s []infer) {
	// normal sort body here, don't need Swap or Len
	// PROBLEM: how to do Equal operations?
}

// SCENARIO: single cache struct
// SOLUTION: use inferred types
type Cache struct {
	items map[string]infer
}

// PROBLEMS: How would the compiler know what type use during initialization, Set, and Get?
// SOLUTION: Maybe it all goes back to initialization?
// That may not work with Set.
func NewCache() *Cache {
	return &Cache{
		items: new(map[string]infer),
	}
}

func (c *Cache) Set(key string, v infer) {
	c.items[key] = v
}

// how would the compiler know that these match?
func (c *Cache) Get(key string) infer {
	return c.items[key]
}
