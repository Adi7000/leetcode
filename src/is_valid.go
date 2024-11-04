package leetcode

func isValid(s string) bool {
	lifo := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			lifo = append(lifo, char)

		} else {
			switch {
			case len(lifo) == 0:
				return false
			case char == ')':
				if lifo[len(lifo)-1] != '(' {
					return false
				}
			case char == '}':
				if lifo[len(lifo)-1] != '{' {
					return false
				}
			case char == ']':
				if lifo[len(lifo)-1] != '[' {
					return false
				}
			}
			lifo = lifo[:len(lifo)-1]

		}
	}
	return len(lifo) == 0
}
