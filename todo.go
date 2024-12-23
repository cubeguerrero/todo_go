package todo

import (
	"fmt"
	"time"
)

type Item struct {
	Title     string
	DoneAt    *time.Time
	CreatedAt *time.Time
}

func (i *Item) String() string {
	state := "❎"
	if i.DoneAt != nil {
		state = "✅"
	}
	return fmt.Sprintf("%s: %s", state, i.Title)
}

type Todo []Item

func New() Todo {
	return Todo{}
}

func (t *Todo) AddItem(title string) {
	now := time.Now()
	item := Item{
		Title:     title,
		CreatedAt: &now,
	}

	*t = append(*t, item)
}

func (t *Todo) DeleteItem(index int) {
	if index < 0 || index >= len(*t) {
		return
	}

	// Remove the item at the given index
	*t = append((*t)[:index], (*t)[index+1:]...)
}
