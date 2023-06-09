package code

import (
	"encoding/hex"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"strconv"
	"strings"
)

var Count int
var IndexesOfSums [][]int
var TextForCoding string
var EncodedBytes []string
var Dict map[string]string

func Encode(win fyne.Window) string {
	discharges := []int{0, 0, 0}
	textBytes := TextToBits(TextForCoding)
	//Encoding
	textBytesArr := strings.Split(textBytes, "")
	Dict = make(map[string]string)
	var regStatus []string
	var regSum []string
	var textBytesArrInt []int
	for key, elem := range textBytesArr {
		number, err1 := strconv.Atoi(elem)
		if err1 != nil {
			dialog.ShowError(err1, win)
		}
		textBytesArrInt = append(textBytesArrInt, number)
		discharges = append([]int{textBytesArrInt[key]}, discharges[:len(discharges)-1]...)

		var dischargesStr []string
		for _, elem2 := range discharges {
			dischargesStr = append(dischargesStr, strconv.Itoa(elem2))
		}

		regStatus = append(regStatus, strings.Join(dischargesStr, ""))

		for _, elem1 := range IndexesOfSums {
			sum := 0
			for _, index := range elem1 {
				sum += discharges[index-1]
			}
			numberToAdd := sum % 2
			EncodedBytes = append(EncodedBytes, strconv.Itoa(numberToAdd))
		}
	}
	fmt.Println(EncodedBytes)
	for i := 0; i < len(EncodedBytes); i += Count {
		regSum = append(regSum, strings.Join(EncodedBytes[i:i+Count], ""))
	}

	for i := 0; i < len(regStatus); i++ {
		Dict[regStatus[i]] = regSum[i]
	}
	fmt.Println(Dict)
	return strings.Join(EncodedBytes, "")
}

func Decode() string {
	var decodedArr []string
	registered := []string{"0", "0", "0"}
	for i := 0; i < len(EncodedBytes); i += Count {
		registered = append([]string{"1"}, registered[:len(registered)-1]...)
		fmt.Println(EncodedBytes[i : i+Count])
		fmt.Println(Dict[strings.Join(registered, "")])

		if strings.Join(EncodedBytes[i:i+Count], "") == Dict[strings.Join(registered, "")] {
			decodedArr = append(decodedArr, "1")
		} else {
			registered[0] = "0"
			if strings.Join(EncodedBytes[i:i+Count], "") == Dict[strings.Join(registered, "")] {
				decodedArr = append(decodedArr, "0")
			}
		}
		fmt.Println(decodedArr)
	}
	return TextFromBits(strings.Join(decodedArr, ""))
}

func TextToBits(text string) string {
	byteSlice := []byte(text)
	hexString := hex.EncodeToString(byteSlice)
	hexInt, _ := strconv.ParseInt(hexString, 16, 64)
	bytes := strconv.FormatInt(hexInt, 2)
	return bytes
}

func TextFromBits(bits string) string {
	hexInt, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		panic(err)
	}
	hexString := fmt.Sprintf("%x", hexInt)
	byteSlice, err1 := hex.DecodeString(hexString)
	if err1 != nil {
		panic(err1)
	}
	return string(byteSlice)
}
