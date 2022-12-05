package main

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	} else {
		idx := len(*s) - 1
		elm := (*s)[idx]
		*s = (*s)[:idx]

		return elm
	}
}
