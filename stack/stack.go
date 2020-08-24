package stack

type Stack interface {
	//copy func name in  github.com/360EntSecGroup-Skylar/excelize/lib.go
	Push(interface{})

	Pop() interface{}

	Peek() interface{}

	Empty() bool

	Len() int
}

type SliceStack struct {
	stack []interface{}
}

func NewArrayStack(size int) *SliceStack {
	return &SliceStack{
		//stack: make([]interface{}, size),
		stack: []interface{}{},
	}
}

func (s *SliceStack) Push(data interface{}) {
	s.stack = append(s.stack, data)
}

func (s *SliceStack) Pop() interface{} {
	pop := s.Peek()
	if s.Len() >= 1 {
		s.stack = s.stack[0 : s.Len()-1]
	}
	return pop
}

func (s *SliceStack) Peek() interface{} {
	if s.Len() < 1 {
		return nil
	}
	return s.stack[s.Len()-1]
}

func (s *SliceStack) Empty() bool {
	return len(s.stack) == 0
}

func (s *SliceStack) Len() int {
	return len(s.stack)
}
