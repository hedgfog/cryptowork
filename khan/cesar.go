package khan

func decode(code string, shift byte) string {
	var res string
	for i := 0; i < len(code); i++ {
		letter := ((byte(code[i]) + shift) - 97) % 25
		res += string(letter + 97)
	}

	return res
}

func PolyalphabeticDecode() {

}
