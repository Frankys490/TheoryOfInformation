package code

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func generative() ([][]int, [][]int) {
	inputMatrix := matrixInput()
	searchMatrix := transposeMatrix(inputMatrix)
	for i := 0; i < len(inputMatrix); i++ {
		var searchColumn []int
		for j := 0; j < len(inputMatrix); j++ {
			searchColumn = append(searchColumn, 0)
		}
		searchColumn[i] = 1
		for key, elem := range searchMatrix {
			if equal(elem, searchColumn) {
				searchMatrix[key], searchMatrix[i] = searchMatrix[i], searchMatrix[key]
			}
		}
	}
	matrixGsys := transposeMatrix(searchMatrix)
	matrixHsys := getHsysFromGsys(searchMatrix)
	return matrixGsys, matrixHsys
}

func getHsysFromGsys(matrix [][]int) [][]int {
	matrix = matrix[len(transposeMatrix(matrix)):]
	matrix = transposeMatrix(matrix)
	for i := 0; i < len(matrix[0]); i++ {
		var addColumn []int
		for j := 0; j < len(matrix[0]); j++ {
			addColumn = append(addColumn, 0)
		}
		addColumn[i] = 1
		matrix = append(matrix, addColumn)
	}
	return transposeMatrix(matrix)
}

func verification() ([][]int, [][]int) {
	inputMatrix := matrixInput()
	searchMatrix := transposeMatrix(inputMatrix)
	for i := 0; i < len(inputMatrix); i++ {
		var searchColumn []int
		for j := 0; j < len(inputMatrix); j++ {
			searchColumn = append(searchColumn, 0)
		}
		searchColumn[i] = 1
		for key, elem := range searchMatrix {
			if equal(elem, searchColumn) {
				searchMatrix[key], searchMatrix[i+(len(searchMatrix)-len(inputMatrix))] = searchMatrix[i+(len(searchMatrix)-len(inputMatrix))], searchMatrix[key]
			}
		}
	}
	matrixHsys := transposeMatrix(searchMatrix)
	matrixGsys := getGsysFromHsys(searchMatrix)
	return matrixGsys, matrixHsys
}

func getGsysFromHsys(matrix [][]int) [][]int {
	matrix = matrix[:(len(matrix) - len(transposeMatrix(matrix)))]
	var squadronMatrix [][]int
	for i := 0; i < len(matrix); i++ {
		var addColumn []int
		for j := 0; j < len(matrix); j++ {
			addColumn = append(addColumn, 0)
		}
		addColumn[i] = 1
		squadronMatrix = append(squadronMatrix, addColumn)
	}
	matrix = append(squadronMatrix, transposeMatrix(matrix)...)
	return transposeMatrix(matrix)
}

func informativeTable(matrix [][]int, k int) [][][]int {
	var table [][][]int
	for i := int64(0); i < int64(math.Pow(2, float64(k))); i++ {
		infWord := strconv.FormatInt(i, 2)
		infWordSlice := strings.Split(infWord, "")
		var infWordSliceInt []int
		for _, elem := range infWordSlice {
			number, err := strconv.Atoi(elem)
			if err != nil {
				panic(err)
			}
			infWordSliceInt = append(infWordSliceInt, number)
		}
		if len(infWordSliceInt) != len(matrix) {
			var zeroWord []int
			for j := 0; j <= len(matrix)-len(infWordSliceInt)-1; j++ {
				zeroWord = append(zeroWord, 0)
			}
			infWordSliceInt = append(zeroWord, infWordSliceInt...)
		}
		var row [][]int
		row = append(row, infWordSliceInt)
		infG, weigh := getInfG(matrix, infWordSliceInt)
		row = append(row, infG, weigh)
		table = append(table, row)
	}
	return table
}

func getInfG(matrix [][]int, word []int) ([]int, []int) {
	var infG []int
	var weigh []int
	weigh = append(weigh, 0)
	for i := 0; i < len(matrix[0]); i++ {
		infG = append(infG, 0)
	}
	for key, elem := range word {
		if elem == 1 {
			row := matrix[key]
			for key0, elem0 := range row {
				infG[key0] += elem0
			}
		}
	}
	for key, elem := range infG {
		if elem%2 != 0 {
			infG[key] = 1
		} else {
			infG[key] = 0
		}
		if infG[key] == 1 {
			weigh[0] += 1
		}
	}
	return infG, weigh
}

func getMinWeigh(table [][][]int) int {
	var minWeigh int
	for key, elem := range table {
		switch key {
		case 0:
			continue
		case 1:
			minWeigh = elem[2][0]
		default:
			if minWeigh > elem[2][0] {
				minWeigh = elem[2][0]
			}
		}
	}
	return minWeigh
}

func getSynTable(matrix [][]int) [][][]int {
	var table [][][]int
	var errorVectors [][]int
	matrix = transposeMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		var errorVector []int
		for j := 0; j < len(matrix); j++ {
			errorVector = append(errorVector, 0)
		}
		errorVector[i] = 1
		errorVectors = append(errorVectors, errorVector)
	}
	for i := 0; i < len(matrix); i++ {
		var row [][]int
		row = append(row, matrix[i], errorVectors[i])
		table = append(table, row)
	}
	return table
}

func findError(table [][][]int, vector []int) []int {
	var vectorErrPlace []int
	for i := 0; i < len(table[0][0]); i++ {
		vectorErrPlace = append(vectorErrPlace, 0)
	}
	for key, elem := range vector {
		if elem == 1 {
			addSyn := table[key][0]
			for keySyn, elemSyn := range addSyn {
				vectorErrPlace[keySyn] += elemSyn
			}
		}
	}
	for key, elem := range vectorErrPlace {
		if elem%2 != 0 {
			vectorErrPlace[key] = 1
		} else {
			vectorErrPlace[key] = 0
		}
	}
	for key, elem := range table {
		if equal(elem[0], vectorErrPlace) {
			if vector[key] == 0 {
				vector[key] = 1
			} else {
				vector[key] = 0
			}
		}
	}

	return vector
}

func transposeMatrix(matrix [][]int) [][]int {
	var matrixTransposed [][]int
	for i := 0; i < len(matrix[0]); i++ {
		var row []int
		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[j][i])
		}
		matrixTransposed = append(matrixTransposed, row)
	}
	return matrixTransposed
}

func matrixInput() [][]int {
	inputData := strings.Split(scanFile("src/inputMatrix.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	var matrix [][]int
	for _, elem := range inputData {
		var row []int
		stringRow := strings.Split(elem, " ")
		for _, elemSecond := range stringRow {
			number, err := strconv.Atoi(elemSecond)
			if err != nil {
				panic(err)
			}
			row = append(row, number)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func scanFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var str string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
		str += ";"
	}
	return str
}

func equal(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func matrixOut(matrix [][]int) {
	for _, elem := range matrix {
		fmt.Println(elem)
	}
}

func Lab1() {
	fmt.Println("Какая матрица подается на вход?\nПорождающая или проверочная? (G/H)")
	var matrixGsys, matrixHsys [][]int
	var matrixType string
	_, err := fmt.Scan(&matrixType)
	if err != nil {
		panic(err)
	}
	switch matrixType {
	case "G":
		matrixGsys, matrixHsys = generative()
	case "H":
		matrixGsys, matrixHsys = verification()
	default:
		log.Fatal("Неверные входные данные")
	}

	fmt.Println("Порождающая систематическая матрица:")
	matrixOut(matrixGsys)
	fmt.Println("Проверочная систематическая матрица:")
	matrixOut(matrixHsys)

	n := len(matrixGsys[0])
	k := len(matrixGsys)
	fmt.Println("n =", n)
	fmt.Println("k =", k)
	fmt.Println("Таблица информационных слов:")
	infTable := informativeTable(matrixGsys, k)
	for _, elem := range infTable {
		fmt.Println(elem)
	}
	minWeigh := getMinWeigh(infTable)
	fmt.Println("d(min) =", minWeigh)
	t := (minWeigh - 1) / 2
	fmt.Println("t =", t)

	fmt.Println("Введите вектор:")
	var vectorStr string
	fmt.Scan(&vectorStr)
	vectorSlice := strings.Split(vectorStr, "")
	var vectorSliceInt []int
	for _, elem := range vectorSlice {
		number, err1 := strconv.Atoi(elem)
		if err1 != nil {
			panic(err)
		}
		vectorSliceInt = append(vectorSliceInt, number)
	}

	fmt.Println("Таблица синдромов:")
	synTable := getSynTable(matrixHsys)
	for _, elem := range synTable {
		fmt.Println(elem)
	}

	correctVector := findError(synTable, vectorSliceInt)
	fmt.Println("Исправленный вектор:")
	fmt.Println(correctVector)

	for _, elem := range infTable {
		if equal(elem[1], correctVector) {
			fmt.Println("Зашифрованное слово:")
			fmt.Println(elem[0])
		}
	}
}
