package main

func convert(s *string) int {
	res := 0
	for i := 0; i < len(*s); i++ {
		res += int((*s)[i] - '0')
		res *= 10
	}
	return res / 10
}
