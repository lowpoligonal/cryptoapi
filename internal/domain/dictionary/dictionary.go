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

func MakeAlphMatrix(alph []rune) [][]rune {
	alphLen := len(alph)
	var matrixSize int
	for i := 0; i < (alphLen / 2); i++ {
		if i*i > alphLen {
			matrixSize = i
		}
	}
	alphMatrix := make([][]rune, 0, matrixSize)
	for i := 0; i < alphLen; i += matrixSize {
		end := min(i+matrixSize, alphLen)
		endStr := alph[i:end]
		if len(endStr) < matrixSize {
			for range matrixSize - len(endStr) {
				endStr = append(endStr, '*')
			}
		}
		alphMatrix = append(alphMatrix, endStr)
	}
	return alphMatrix
}
