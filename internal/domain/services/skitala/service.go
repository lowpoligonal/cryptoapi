package skitala

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) Encode(mKey int, str string) (string, error) {
	if mKey <= 0 {
		return str, errors.New("zero key")
	}

	strRune := []rune(str)
	lenStr := len(strRune)

	encoded := make([]rune, 0, lenStr)

	nKey := (lenStr-1)/mKey + 1

	subStr := make([][]rune, 0, nKey)
	for i := 0; i < lenStr; i += nKey {
		end := min(i+nKey, lenStr)
		endStr := strRune[i:end]
		if len(endStr) < nKey {
			for range nKey - len(endStr) {
				endStr = append(endStr, ' ')
			}
		}
		subStr = append(subStr, endStr)
	}

	for i := range nKey {
		for _, sub := range subStr {
			if i < len(sub) {
				encoded = append(encoded, sub[i])
			}
		}
	}

	return string(encoded), nil
}

func (srv *Service) Decode(mKey int, str string) (string, error) {
	if mKey <= 0 {
		return str, errors.New("zero key")
	}

	strRune := []rune(str)
	lenStr := len(strRune)

	var decoded []rune

	nKey := (lenStr-1)/mKey + 1

	subStr := make([][]rune, 0, mKey)
	for i := 0; i < lenStr; i += mKey {
		end := min(i+mKey, lenStr)
		subStr = append(subStr, strRune[i:end])
	}

	for j := range nKey {
		for i := range len(subStr) {
			if j < len(subStr[i]) {
				decoded = append(decoded, subStr[i][j])
			}
		}
	}

	for len(decoded) > 0 && decoded[len(decoded)-1] == ' ' {
		decoded = decoded[:len(decoded)-1]
	}

	return string(decoded), nil
}
