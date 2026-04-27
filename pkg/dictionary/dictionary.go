package dictionary

import "errors"

var Dictionary = map[string][]rune{
	"rusLow":  []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
	"rusUp":   []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"),
	"enLow":   []rune("abcdefghijklmnopqrstuvwxyz"),
	"enUp":    []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"numbers": []rune("0123456789"),
	"special": []rune("\n\t !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"),
}

var KeyboardPrintableRunes = []rune(
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ" +
		"абвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
		"0123456789" +
		" !\"#$%&№'()*+,-./:;<=>?@[\\]^_`{|}~\n",
)

var LenPrintableRunes = len(KeyboardPrintableRunes)

var RusMatrix = [][]rune{
	[]rune("АБВГДЕ"), []rune("ЁЖЗИЙК"), []rune("ЛМНОПР"), []rune("СТУФХЦ"), []rune("ЧШЩЪЫЬ"), []rune("ЭЮЯ.,!"),
}

var EngMatrix = [][]rune{
	[]rune("ABCDE"), []rune("FGHIK"), []rune("LMNOP"), []rune("QRSTU"), []rune("VWXYZ"),
}

var RusMap = map[rune][2]int{
	'А': {0, 0}, 'Б': {0, 1}, 'В': {0, 2}, 'Г': {0, 3}, 'Д': {0, 4}, 'Е': {0, 5},
	'Ё': {1, 0}, 'Ж': {1, 1}, 'З': {1, 2}, 'И': {1, 3}, 'Й': {1, 4}, 'К': {1, 5},
	'Л': {2, 0}, 'М': {2, 1}, 'Н': {2, 2}, 'О': {2, 3}, 'П': {2, 4}, 'Р': {2, 5},
	'С': {3, 0}, 'Т': {3, 1}, 'У': {3, 2}, 'Ф': {3, 3}, 'Х': {3, 4}, 'Ц': {3, 5},
	'Ч': {4, 0}, 'Ш': {4, 1}, 'Щ': {4, 2}, 'Ъ': {4, 3}, 'Ы': {4, 4}, 'Ь': {4, 5},
	'Э': {5, 0}, 'Ю': {5, 1}, 'Я': {5, 2}, '.': {5, 3}, ',': {5, 4}, '!': {5, 5},
}

var EngMap = map[rune][2]int{
	'A': {0, 0}, 'B': {0, 1}, 'C': {0, 2}, 'D': {0, 3}, 'E': {0, 4},
	'F': {1, 0}, 'G': {1, 1}, 'H': {1, 2}, 'I': {1, 3}, 'J': {1, 3}, 'K': {1, 4},
	'L': {2, 0}, 'M': {2, 1}, 'N': {2, 2}, 'O': {2, 3}, 'P': {2, 4},
	'Q': {3, 0}, 'R': {3, 1}, 'S': {3, 2}, 'T': {3, 3}, 'U': {3, 4},
	'V': {4, 0}, 'W': {4, 1}, 'X': {4, 2}, 'Y': {4, 3}, 'Z': {4, 4},
}

var FillerRune = '\x00'

func GetIndex(char rune) (int, error) {
	for i, r := range KeyboardPrintableRunes {
		if r == char {
			return i, nil
		}
	}
	return -1, errors.New("непечатный символ")
}

func FindSymbolInfo(char rune) (string, int) {
	for key, alphabet := range Dictionary {
		for i, r := range alphabet {
			if r == char {
				return key, i
			}
		}
	}
	return "", -1
}
