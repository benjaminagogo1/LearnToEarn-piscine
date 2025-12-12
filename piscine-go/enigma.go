package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Save original values
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	// Move values around
	***a = tempB     // b -> as
	*b = tempD       // d -> b
	*******c = tempA // a -> c
	****d = tempC    // c -> d
}
