package todo

type Todo struct {
  Name    string
  Items []string
}

func (t *Todo) AddItem(item string) []string {
  t.Items = append(t.Items, item)
  return t.Items
}

func New(name string) Todo {
  return Todo{Name: name}
}
