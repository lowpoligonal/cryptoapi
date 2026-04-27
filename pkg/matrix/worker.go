package matrix

import (
	"cryptoapi/pkg/dictionary"
	"errors"
)

func determinant(matrix [][]int) (int, error) {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0], nil
	}

	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
	}

	det := 0
	for j := 0; j < n; j++ {
		minor := make([][]int, n-1)
		for i := 1; i < n; i++ {
			row := make([]int, 0, n-1)
			for k := 0; k < n; k++ {
				if k != j {
					row = append(row, matrix[i][k])
				}
			}
			minor[i-1] = row
		}
		minorDet, err := determinant(minor)
		if err != nil {
			return 0, err
		}
		sign := 1
		if j%2 == 1 {
			sign = -1
		}
		det += sign * matrix[0][j] * minorDet
	}
	return det, nil
}

func DeterminantMod(matrix [][]int, m int) (int, error) {
	det, err := determinant(matrix)
	if err != nil {
		return 0, err
	}
	det %= m
	if det < 0 {
		det += m
	}
	return det, nil
}

func ExtendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func modInverse(a, m int) (int, error) {
	gcd, x, _ := ExtendedGCD(a, m)
	if gcd != 1 {
		return 0, errors.New("обратный элемент не существует")
	}
	result := (x%m + m) % m
	return result, nil
}

func MultiplyMatrixVectorMod(matrix [][]int, vector []int, m int) ([]int, error) {
	n := len(matrix)
	if len(vector) != n {
		return nil, errors.New("размерность матрицы и вектора не совпадает")
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum = (sum + matrix[i][j]*vector[j]) % m
		}
		result[i] = sum
	}
	return result, nil
}

func InverseMatrixMod(matrix [][]int, m int) ([][]int, error) {
	n := len(matrix)

	det, err := DeterminantMod(matrix, m)
	if err != nil {
		return nil, err
	}

	invDet, err := modInverse(det, m)
	if err != nil {
		return nil, errors.New("матрица необратима по модулю")
	}

	inverse := make([][]int, n)

	if n == 1 {
		inverse[0] = []int{invDet}
		return inverse, nil
	}

	for i := 0; i < n; i++ {
		inverse[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minor := make([][]int, n-1)
			for k := 0; k < n-1; k++ {
				minor[k] = make([]int, n-1)
			}

			minorRow := 0
			for row := 0; row < n; row++ {
				if row == j {
					continue
				}
				minorCol := 0
				for col := 0; col < n; col++ {
					if col == i {
						continue
					}
					minor[minorRow][minorCol] = matrix[row][col]
					minorCol++
				}
				minorRow++
			}

			minorDet, _ := determinant(minor)

			cofactor := minorDet
			if (i+j)%2 != 0 {
				cofactor = -cofactor
			}

			val := (cofactor * invDet) % m
			if val < 0 {
				val += m
			}
			inverse[i][j] = val
		}
	}
	return inverse, nil
}

func ConvertToMatrix(keyRune []rune) ([][]int, int, error) {
	var vLen int

	for i := 1; ; i++ {
		if i*i >= len(keyRune) {
			vLen = i
			break
		}
	}

	matrix := make([][]int, vLen)

	keyPos := 0
	for i := 0; i < vLen; i++ {
		matrix[i] = make([]int, vLen)
		for j := 0; j < vLen; j++ {
			keyChar := keyRune[keyPos%len(keyRune)]
			idx, err := dictionary.GetIndex(keyChar)
			if err != nil {
				return matrix, 0, err
			}
			matrix[i][j] = idx
			keyPos++
		}
	}

	return matrix, vLen, nil
}

func ConvertToVectors(strRune []rune, vectorLen int) ([][]int, error) {
	var vectors [][]int

	for i := 0; i < len(strRune); {
		end := i + vectorLen
		if end > len(strRune) {
			end = len(strRune)
		}

		vRune := make([]rune, end-i)
		copy(vRune, strRune[i:end])

		for len(vRune) < vectorLen {
			vRune = append(vRune, ' ')
		}

		var vector []int
		for _, char := range vRune {
			idx, err := dictionary.GetIndex(char)
			if err != nil {
				return nil, err
			}
			vector = append(vector, idx)
		}
		vectors = append(vectors, vector)

		i = i + vectorLen
	}

	return vectors, nil
}
