package main

func inputValid(s *string) bool {

	for _, val := range *s {
		if val < '0' || '9' < val {
			return false
		}
	}
	return true
}
