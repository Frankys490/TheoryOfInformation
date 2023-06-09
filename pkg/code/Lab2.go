package code

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Lab2(expr, fieldStr string) string {
	var answer string

	field, err := strconv.Atoi(fieldStr)
	if err != nil {
		answer = "Ошибка"
	}
	exprAr := strings.Split(expr, "")
	for i := 2; i < field; i++ {
		if field%i == 0 {
			log.Fatal("Поле не простое число")
		}
	}
	bracketsLeft, bracketsRight := checkBrackets(exprAr)
	if len(bracketsLeft) != len(bracketsRight) {
		answer = "Ошибка"
	}

	for {
		bracketsLeft, bracketsRight = checkBrackets(exprAr)
		if len(bracketsRight) == 0 {
			break
		}
		start, end := calcBrackets(exprAr, bracketsLeft, bracketsRight)
		exprToCalc := exprAr[start+1 : end]
		exprToCalc = exprFinish(exprToCalc, field)

		var newExprAr []string
		newExprAr = append(newExprAr, exprAr[:start]...)
		newExprAr = append(newExprAr, exprToCalc...)
		exprAr = append(newExprAr, exprAr[end+1:]...)
	}

	exprAr = exprFinish(exprAr, field)
	answer = exprAr[0]
	answerInt, errConv := strconv.Atoi(answer)
	if errConv != nil {
		panic(errConv)
	}
	for answerInt > field {
		answerInt -= field
	}
	answerStr := strconv.Itoa(answerInt)
	return answerStr
}

func calculate(first string, second string, sign string, field int) int {
	firstNumber, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}
	secondNumber, err := strconv.Atoi(second)
	if firstNumber < 0 {
		for firstNumber < 0 {
			firstNumber += field
		}
	}
	if secondNumber < 0 {
		for secondNumber < 0 {
			secondNumber += field
		}
	}
	if err != nil {
		panic(err)
	}
	var result int
	switch sign {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
		for result < 0 {
			result += field
		}
	case "*":
		result = firstNumber * secondNumber
	case "/":
		if secondNumber == 0 {
			log.Fatal("Деление на ноль")
		}
		_, _, result = gcdex(field, secondNumber)
		for result < 0 {
			result += field
		}
		result *= firstNumber
	}
	return result
}

func calcBrackets(exprAr []string, bracketsLeft []int, bracketsRight []int) (int, int) {
	var start, end int
	for _, elem1 := range bracketsLeft {
		exit := false
		for _, elem2 := range bracketsRight {
			if strings.Index(strings.Join(exprAr[elem1+1:elem2], ""), ")") == -1 &&
				strings.Index(strings.Join(exprAr[elem1+1:elem2], ""), "(") == -1 {
				start = elem1
				end = elem2
				exit = true
			}
		}
		if exit {
			break
		}
	}
	return start, end

}

func exprFinish(exprToCalc []string, field int) []string {
	var sings []string
	var keys []int
	keys = append(keys, 0)
	sings = append(sings, "{}")
	for key, elem := range exprToCalc {
		if !unicode.IsDigit([]rune(elem)[0]) {
			sings = append(sings, elem)
			keys = append(keys, key)
		}
	}
	keys = append(keys, len(exprToCalc))
	sings = append(sings, "{}")
	counter := 0
	for _, elem := range sings {
		if elem == "*" || elem == "/" {
			counter++
		}
	}
	for len(sings) != 2 {
		for counter != 0 {
			for key, elem := range sings {
				if elem == "*" || elem == "/" {
					var firstNumber, secondNumber string
					if key == 1 {
						firstNumber = strings.Join(exprToCalc[:keys[key]], "")
						secondNumber = strings.Join(exprToCalc[keys[key]+1:keys[key+1]], "")
						result := calculate(firstNumber, secondNumber, elem, field)
						resultStr := strconv.Itoa(result)

						var newExprToCalc []string
						newExprToCalc = append(newExprToCalc, resultStr)
						exprToCalc = append(newExprToCalc, exprToCalc[keys[key+1]:]...)
					} else if key == len(keys)-2 {
						firstNumber = strings.Join(exprToCalc[keys[key-1]+1:keys[key]], "")
						secondNumber = strings.Join(exprToCalc[keys[key]+1:], "")
						result := calculate(firstNumber, secondNumber, elem, field)
						resultStr := strconv.Itoa(result)

						exprToCalc = append(exprToCalc[:keys[key-1]+1], resultStr)
					} else {
						firstNumber = strings.Join(exprToCalc[keys[key-1]+1:keys[key]], "")
						secondNumber = strings.Join(exprToCalc[keys[key]+1:keys[key+1]], "")
						result := calculate(firstNumber, secondNumber, elem, field)
						resultStr := strconv.Itoa(result)

						var newExprToCalc []string
						newExprToCalc = append(newExprToCalc, exprToCalc[:keys[key-1]+1]...)
						newExprToCalc = append(newExprToCalc, resultStr)
						exprToCalc = append(newExprToCalc, exprToCalc[keys[key+1]:]...)
					}
					sings = append(sings[:key], sings[key+1:]...)
					keys = append(keys[:0], 0)
					for key1, elem1 := range exprToCalc {
						if !unicode.IsDigit([]rune(elem1)[0]) {
							keys = append(keys, key1)
						}
					}
					keys = append(keys, len(exprToCalc))
					counter--
				}
			}
		}
		for key, elem := range sings {
			if elem == "+" || elem == "-" {
				firstNumber := strings.Join(exprToCalc[:keys[key]], "")
				secondNumber := strings.Join(exprToCalc[keys[key]+1:keys[key+1]], "")
				if firstNumber == "" {
					firstNumber = "0"
				}
				result := calculate(firstNumber, secondNumber, elem, field)
				resultStr := strconv.Itoa(result)

				var newExprToCalc []string
				newExprToCalc = append(newExprToCalc, resultStr)
				exprToCalc = append(newExprToCalc, exprToCalc[keys[key+1]:]...)

				sings = append(sings[:key], sings[key+1:]...)
				keys = append(keys[:0], 0)
				for key1, elem1 := range exprToCalc {
					if !unicode.IsDigit([]rune(elem1)[0]) {
						keys = append(keys, key1)
					}
				}
				keys = append(keys, len(exprToCalc))
			}
		}
	}
	return exprToCalc
}

func checkBrackets(expr []string) ([]int, []int) {
	var bracketsLeft []int
	var bracketsRight []int
	for key, elem := range expr {
		if elem == "(" {
			bracketsLeft = append(bracketsLeft, key)
		} else if elem == ")" {
			bracketsRight = append(bracketsRight, key)
		}
	}
	return bracketsLeft, bracketsRight
}

func gcdex(a, b int) (int, int, int) {
	if b != 0 {
		d, x, y := gcdex(b, a%b)
		return d, y, x - y*(a/b)
	} else {
		return a, 1, 0
	}
}
