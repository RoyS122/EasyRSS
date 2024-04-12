package main

func Split(s string, sep rune) (res []string) {
	var word string
	for _, k := range s {
		if k == sep {
			res = append(res, word)
			word = ""
		} else {
			word += string(k)
		}
	}
	if len(word) != 0 {
		res = append(res, word)
	}
	return res
}

func Atoi(str string) (res int) {
	var neg int = 1

	if str[0] == '-' {
		str, neg = str[1:], -1
	}

	for i, k := range RevertString(str) {

		res += PowInt(10, i) * int(rune(k)-'0')

	}

	return res * neg
}

func RevertString(in_str string) (res string) {
	for i := len(in_str); i > 0; i-- {
		res += string(in_str[i-1])
	}
	return res
}

func PowInt(a, b int) (res int) {
	if b == 0 {
		return 1
	}
	res = a
	for i := 1; i < b; i++ {
		res = res * a
	}

	return res
}

func removeChar(c rune, s string) (r string) {
	for _, k := range s {
		if k != c {
			r += string(k)
		}
	}
	return r
}
