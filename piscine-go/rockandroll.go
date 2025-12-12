package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	div2 := n%2 == 0
	div3 := n%3 == 0

	if div2 && div3 {
		return "rock and roll\n"
	}
	if div2 {
		return "rock\n"
	}
	if div3 {
		return "roll\n"
	}

	return "error: non divisible\n"
}
