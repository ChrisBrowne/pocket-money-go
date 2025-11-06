package server

type ChildStore interface {
	GetAllChildren() []Child
}

type InMemoryChildStore struct {
	data map[string]Child
}

func NewInMemoryChildStore() *InMemoryChildStore {
	return &InMemoryChildStore{
		data: make(map[string]Child),
	}
}

func (s *InMemoryChildStore) GetAllChildren() []Child {
	children := make([]Child, 0, len(s.data))

	for _, v := range s.data {
		children = append(children, v)
	}
	return children
}
func (s *InMemoryChildStore) AddChild(name string) {
	s.data[name] = Child{Name: name}
}
