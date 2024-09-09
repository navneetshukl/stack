package stack

import "errors"

type Service struct {
	Stack []int
}

func NewService(stack []int) *Service {
	return &Service{
		Stack: stack,
	}
}

type StackService interface {
	Push(val int)
	Pop()
	Top() int
}

func (s *Service) Push(value int) {
	s.Stack = append(s.Stack, value)
}

func (s *Service) Pop() error {
	if len(s.Stack) == 0 {
		return errors.New("stack is empty")
	}

	s.Stack = s.Stack[:len(s.Stack)-1]
	return nil
}

func (s *Service) Top() (int, error) {
	if len(s.Stack) == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.Stack[len(s.Stack)-1], nil
}
