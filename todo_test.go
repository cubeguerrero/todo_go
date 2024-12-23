package todo_test

import (
	"testing"
	"time"

	"cubeguerrero.com/todo"
)

func TestItemStringerImpl(t *testing.T) {
	notDoneItem := todo.Item{
		Title: "Not yet done",
	}

	expected := "❎: Not yet done"

	if notDoneItem.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, notDoneItem.String())
	}

	now := time.Now()
	doneItem := todo.Item{
		Title:  "It is done",
		DoneAt: &now,
	}
	expected = "✅: It is done"

	if doneItem.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, doneItem.String())
	}
}

func TestCreateTodo(t *testing.T) {
	todos := todo.New()

	if len(todos) != 0 {
		t.Errorf("Expected todos length: %d, got: %d", 0, len(todos))
	}
}

func TestAddItem(t *testing.T) {
	todos := todo.New()

	todos.AddItem("Another one")

	if len(todos) != 1 {
		t.Errorf("Expected todos length: %d, got: %d", 1, len(todos))
	}

	item := todos[0]

	if item.Title != "Another one" {
		t.Errorf("Expected item title: Another one, got: %v", item.Title)
	}

	if item.CreatedAt == nil {
		t.Errorf("Expected item CreatedAt not to be nil")
	}

	if item.DoneAt != nil {
		t.Errorf("Expected item DoneAt to be nil")
	}
}

func TestDeleteItem(t *testing.T) {
	todos := todo.New()

	todos.AddItem("1")
	todos.AddItem("2")
	todos.AddItem("3")

	if len(todos) != 3 {
		t.Errorf("Expected todos length to be 3, got: %d", len(todos))
	}

	todos.DeleteItem(2)

	if len(todos) != 2 {
		t.Errorf("Expected todos length to be 2, got: %d", len(todos))
	}
}
