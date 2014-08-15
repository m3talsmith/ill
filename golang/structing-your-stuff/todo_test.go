package todo

import "testing"

var todo Todo = New("Test Todo")

func TestNew (t *testing.T) {
  if todo.Name != "Test Todo" {
    t.Errorf("Expected 'Test Todo', Got '%s'", todo.Name)
  }
}

func TestAddItem (t *testing.T) {}
func TestEditItem (t *testing.T) {}
func TestDeleteItem (t *testing.T) {}

func TestGenerateSlug (t *testing.T) {}
func TestTodoExists (t *testing.T) {}

func TestSave (t *testing.T) {}
func TestLoad (t *testing.T) {}
func TestDelete (t *testing.T) {}
