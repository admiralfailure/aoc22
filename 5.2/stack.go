package main

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) PushMany(strs []string) {
	*s = append(*s, strs...)
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

func (s *Stack) PopMany(count int) []string {
	if len(*s) == 0 {
		return make([]string, 0)
	} else {
		idx := len(*s) - count
		elms := (*s)[idx : idx+count]
		*s = (*s)[:idx]

		return elms
	}
}
