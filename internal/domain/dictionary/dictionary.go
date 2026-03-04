package dictionary

var Dictionary = map[string][]rune{
	"rusLow":  []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
	"rusUp":   []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"),
	"enLow":   []rune("abcdefghijklmnopqrstuvwxyz"),
	"enUp":    []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"numbers": []rune("0123456789"),
	"special": []rune(" !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"),
}

func FindSymbolInfo(char rune) (category string, index int) {
	for key, alphabet := range Dictionary {
		for i, r := range alphabet {
			if r == char {
				return key, i
			}
		}
	}
	return "unknown", -1
}
