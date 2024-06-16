package main

func checkMap(s *string) bool {
	if len(*s) != m {
		return false
	}
	for i := 0; i < m; i++ {
		if (*s)[i] != '.' && (*s)[i] != '*' {
			return false
		}
	}
	return true
}
