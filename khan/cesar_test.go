package khan

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const MaxWordLen = 10

func getDictWords() (map[string]int, error) {
	dat, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		return nil, err
	}
	words := strings.ToLower(string(dat))
	wordsSlice := strings.Split(words, "\n")
	dic := make(map[string]int)
	for _, word := range wordsSlice {
		dic[word] = 1
	}
	return dic, nil
}
func TestCesar(t *testing.T) {
	words, err := getDictWords()
	if err != nil {
		t.Fatal(err)
	}
	text1 := "vwduwljudeehghyhubwklqjlfrxogilqgsohdvhuhwxuqdqbeoxhsulqwviruydxowdqgdodupghvljqedvhgrqzklfkedqnbrxghflghrqldpvhwwlqjxsvdihkrxvhfr"
	var res string
	for i := 0; i < 25; i++ {
		res = ""
		decodeRes := decode(text1, byte(i))
		runes := []rune(decodeRes)
		var lastIndex = 0
		for lastIndex < len(runes) {
			var word string
			for i := lastIndex + MaxWordLen; i > lastIndex; i-- {
				if i >= len(decodeRes) {
					break
				}
				word = ""
				if words[string(runes[lastIndex:i])] == 1 {
					word = string(runes[lastIndex:i])
					lastIndex = i
					break
				}
			}
			if word == "" {
				lastIndex = i
				break
			}
			res += word + " "
		}
		fmt.Printf("decode: %v\n", decodeRes)
		fmt.Printf("shift %v: %v\n\n", i, res)
	}

}

func TestCesar2(t *testing.T) {
	//text1 := "gluhtlishjrvbadvyyplkaohavbyjpwolypzavvdlhrvuuleatlzzhnlzdpajoavcpnlulyljpwolyrlfdvykpzaolopkkluzftivsvmklhaoputfmhcvypalovsilpuluk"
	text2 := "vwduwljudeehghyhubwklqjlfrxogilqgsohdvhuhwxuqdqbeoxhsulqwviruydxowdqgdodupghvljqedvhgrqzklfkedqnbrxghflghrqldpvhwwlqjxsvdihkrxvhfr"
	for i := 0; i < 25; i++ {
		res := decode(text2, byte(i))
		t.Logf("shift %v: %v", i, res)

	}
	//TestCesar2:

}
