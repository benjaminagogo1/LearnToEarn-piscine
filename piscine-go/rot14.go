package piscine

func Rot14(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			// Shift forward by 14
			newRune := r + 14
			// Wrap around if it goes past 'z'
			if newRune > 'z' {
				newRune = newRune - 26
			}
			result += string(newRune)
		} else if r >= 'A' && r <= 'Z' {
			// Shift forward by 14
			newRune := r + 14
			// Wrap around if it goes past 'Z'
			if newRune > 'Z' {
				newRune = newRune - 26
			}
			result += string(newRune)
		} else {
			// Non-alphabet characters remain unchanged
			result += string(r)
		}
	}
	return result
}
