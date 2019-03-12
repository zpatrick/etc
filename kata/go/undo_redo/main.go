package main

import "fmt"

type Operation func(data map[string]interface{})

type UndoRedo struct {
	data map[string]interface{}
	undo []Operation
	redo []Operation
}

func NewUndoRedo() *UndoRedo {
	return &UndoRedo{
		data: map[string]interface{}{},
		undo: []Operation{},
		redo: []Operation{},
	}
}

func (u *UndoRedo) Get(key string) interface{} {
	return u.data[key]
}

func (u *UndoRedo) Set(key string, val interface{}) {
	prev, existed := u.data[key]
	op := func(data map[string]interface{}) {
		_, exists := data[key]
		if existed {
			data[key] = prev
			return
		}

		if exists {
			delete(data, key)
		}
	}

	u.undo = append(u.undo, op)
	u.data[key] = val
}

func (u *UndoRedo) Del(key string) {

}

func (u *UndoRedo) Undo() {
	if len(u.undo) == 0 {
		return
	}

	op := u.undo[len(u.undo)-1]
	u.undo = u.undo[:len(u.undo)-1]
	op(u.data)
}

func main() {
	m := NewUndoRedo()
	m.Set("a", 1)
	m.Set("a", 2)
	fmt.Println(m.Get("a"))
	m.Undo()
	m.Undo()
	fmt.Println(m.Get("a"))
}
