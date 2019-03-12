package nav

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type Tree struct {
	Name     string
	Value    interface{}
	parent   *Tree
	children []*Tree
}

func NewTree(name string, value interface{}, parent *Tree, children ...*Tree) *Tree {
	return &Tree{
		Name:     name,
		Value:    value,
		parent:   parent,
		children: children,
	}
}

func (t *Tree) Parent() *Tree {
	return t.parent
}

func (t *Tree) SetParent(p *Tree) {
	t.parent = p
}

func (t *Tree) Children() []*Tree {
	return t.children
}

func (t *Tree) Child(name string) *Tree {
	for _, child := range t.children {
		if child.Name == name {
			return child
		}
	}

	return nil
}

func (t *Tree) AddChild(c *Tree) {
	if t.children == nil {
		t.children = []*Tree{}
	}

	t.children = append(t.children, c)
}

func LoadJSONTree(data []byte) (*Tree, error) {
	var blob map[string]interface{}
	if err := json.Unmarshal(data, &blob); err != nil {
		return nil, err
	}

	root := NewTree("root", nil, nil)
	for key, val := range blob {
		buildTree(root, key, val)
	}

	return root, nil
}

func buildTree(parent *Tree, name string, value interface{}) {
	child := NewTree(name, value, parent)
	parent.AddChild(child)

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			buildTree(child, key.String(), v.MapIndex(key))
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			buildTree(child, strconv.Itoa(i), v.Index(i).Interface())
		}
	}
}
