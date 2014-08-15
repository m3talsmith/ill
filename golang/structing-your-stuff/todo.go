package todo

type Todo struct {
  Name string
  Items []string
}

func New(name string) Todo {
  t := Todo{Name: name}
  return t
}
