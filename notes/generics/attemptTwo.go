package main

// https://blog.merovius.de/2018/09/05/scrapping_contracts.html
// make typeclasses a first-class type (concat, numeric), specify variable names with name:typeclass
// probably a best practice would be to use 't' for the typeclass variable name.   
// make users define their own typeclass? 
// allow users to make their own objects typeclasses 

func (type T concat) Sum(c ...T) T {
	sum := make(T)
	for _, v := range n {
		sum += v 
	}

	return sum
} 

// how to allow custom types to satisfy typeclasses?
type Foo struct {
	count int 
}

contract (f Foo) numeric {
	return f.count
}

// how to define our own typeclasses? 
// is not allowing that an option?
// it seems you fundamentally need some base types, 
// maybe the builtin typeclasses can be aggregated to make larger ones, 
// but that's about it? Otherwise, seems like you'd just want to use interfaces 

type Number typeclass {
	arith
	equal
	bitwise
}


type Cache(t Type) struct {
	items map[string]T
}

// Example of caller specifying type as first arg to constructor
// this is nice b/c it's very flexible - it will probably be best practice
// we would also want a new keyword like make that creates a new generic object.
// I'll use create here.    
func NewCache(t Type) *Cache {
	c := create(*Cache, t)
	c.items = make(map[string]T)
	return c
}

// problem: T isn't defined here (possibly) -> it's only defined in the scope 
// of Cache(t Type). is that good enough? 
func (c *Cache) Get(key string) T {
	return c.items[key]
}

func (c *Cache) Set(key string, v T) {
	c.items[key] = v
}
