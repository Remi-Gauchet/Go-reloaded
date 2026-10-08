package tolower

func ToLower(s string) string {
	serierunes := []rune(s)
	for i, r := range serierunes {
		if r >= 'A' && r <= 'Z' {
			serierunes[i] = r + ('a' - 'A')
		}
	}
	return string(serierunes)
}
