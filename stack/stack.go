package stack

type stack struct {
	bytes   []byte
	size    int
	nextIdx int
}

type Stack interface {
	Push(value byte)
	Pop() byte
	Size() int
}

func NewStack(size int) Stack {
	bytes := make([]byte, size)
	return &stack{bytes: bytes, size: size, nextIdx: 0}
}

func (s *stack) Push(value byte) {
	if s.nextIdx == s.size {
		return
	}
	s.bytes[s.nextIdx] = value
	s.nextIdx++
}

func (s *stack) Pop() byte {
	if s.nextIdx == 0 {
		return byte(0)
	}
	value := s.bytes[s.nextIdx-1]
	s.nextIdx--
	return value
}

func (s *stack) Size() int {
	return s.size
}
