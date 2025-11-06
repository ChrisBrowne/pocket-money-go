package server

type ChildStore interface {
	GetAllChildren() []Child
	GetChild(name string) Child
	SetChild(child Child)
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

func (s *InMemoryChildStore) GetChild(name string) Child {
	return s.data[name]
}

func (s *InMemoryChildStore) SetChild(child Child) {
	s.data[child.Name] = child
}
