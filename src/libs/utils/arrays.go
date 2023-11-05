package utils

type newStringArray struct {
	elements []string
}

func (s newStringArray) contains(target string) bool {
	for _, elem := range s.elements {
		if elem == target {
			return true
		}
	}
	return false
}

func (s newStringArray) containsKey(target string) (string, bool) {
	for _, elem := range s.elements {
		if elem == target {
			return elem, true
		}
	}
	return "", false
}
