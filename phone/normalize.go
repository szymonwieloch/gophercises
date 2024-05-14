package main

// Normalizes the provided number
func normalize(number string) string {
	result := []rune{}
	for _, ch := range number {
		if ch >= '0' && ch <= '9' {
			result = append(result, ch)
		}
	}
	return string(result)
}
