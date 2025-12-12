package main

import (
	"fmt"
)

/*func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	} else {
		return "error: non divisible"
	}
}
func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}*/

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }

/*func Digitlen(n, base int) int {
	if base < 2 || base < 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n = n / base
		count++
	}
	return count
}
func main() {
	fmt.Println(Digitlen(100, 10))
	fmt.Println(Digitlen(100, 2))
	fmt.Println(Digitlen(-100, 16))
	fmt.Println(Digitlen(100, -1))
}*/

/*func FirstWord(s string) string {
	if len(s) == 0 {
		return "\n"
	}
	start := 0
	end := 0
	end = start
	for end < len(s) && s[end] != ' ' {
		end++
	}
	return s[start:end] + "\n"
}
func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}*/

/*
	func PrintIf(s string) string {
		if len(s) >= 3 {
			return "G\n"
		}
		return "invalid input"
	}

	func main() {
		fmt.Print(PrintIf("abcdefz"))
		fmt.Print(PrintIf("abc"))
		fmt.Print(PrintIf(""))
		fmt.Print(PrintIf("14"))
	}
*/
/*func HashCode(dec string) string {
	res := " "
	for _, r := range dec {
		if h := (int(r) + len(dec)) % 127; h < 33 {
			res += string(h + 33)
		} else {
			res += string(h)
		}
	}
	return res

res := " "
for _, r := range dec {
	if h := (int(r) + len(dec)) %127; h < 33 {
		res += string(h + 33)
	} else {
	 	res += string(h)
	}
}



}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}*/

/*
	func PrintIfNot(s string) string {
		if len(s) < 3 {
			return "G\n"
		}
		return "invalid input\n"
	}

	func main() {
		fmt.Print(PrintIfNot("abcdefz"))
		fmt.Print(PrintIfNot("abc"))
		fmt.Print(PrintIfNot(""))
		fmt.Print(PrintIfNot("14"))
	}
*/
// func PrintIf(s string) string {
// 	if s == "" {
// 		return "G\n"
// 	}
// 	if len(s) >= 3 {
// 		return "G\n"
// 	}
// 	return "invalid input\n"

// }
// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }

// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}

// 	s := os.Args[1]
// 	from := os.Args[2]
// 	to := os.Args[3]

// 	if len(from) != 1 || len(to) != 1 {
// 		fmt.Println(s)
// 		return
// 	}

// 	result := ""
// 	for _, char := range s {
// 		if string(char) == from {
// 			result += to
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	fmt.Println(result)
// }

// func HashCode(dec string) string {
// 	size := len(dec)
// 	result := ""

// 	for i := 0; i < size; i++ {
// 		c := dec[i]
// 		hashed := (int(c) + size) % 127

// 		if hashed < 33 {
// 			hashed += 33
// 		}
// 		result += string(rune(hashed))

// 	}

// 	return result
// }
// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }

// func LastWord(s string) string {
// 	end := len(s) - 1

// 	for end >= 0 && (s[end] == ' ' || s[end] == '\n' || s[end] == '\t') {
// 		end--
// 	}
// 	if end < 0 {
// 		return "\n"
// 	}

// 	start := end

// 	for end >= 0 && s[start] != ' ' && s[start] != '\n' && s[start] != '\t' {
// 		start--
// 	}
// 	return s[start+1:end+1] + "\n"
// }

// func main() {
// 	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(LastWord(" lorem,ipsum "))
// 	fmt.Print(LastWord(" "))
// }

// func FirstWord(s string) string {
// 	if len(s) == 0 {
// 		return "\n"
//
// 	}start := 0
// 	end := 0

// 	for end < len(s) && s[end] != ' ' {
// 		end++
// 	}
// 	return s[start:end] + "\n"
// }
// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("hello   .........  bye"))
// }

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= 'A' && c <= 'Z') && !(c >= 'a' && c <= 'z') {
			return s
		}
	}

	if len(s) > 1 && s[0] >= 'A' && s[0] <= 'Z' && s[1] >= 'A' && s[1] <= 'Z' {
		return s
	}

	res := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'A' && c <= 'Z' && i != 0 {
			res += "_"
		}

		res += string(c)
	}

	return res
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
