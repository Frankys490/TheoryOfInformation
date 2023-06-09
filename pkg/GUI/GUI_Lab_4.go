package GUI

import (
	Code "TheoryOfInformation/pkg/code"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func GuiLab4() {
	a := app.New()
	win := a.NewWindow("Четвертая лаба")
	win.Resize(fyne.NewSize(720, 300))

	labelSum := widget.NewLabel("Введите количество сумматоров")
	entrySum := widget.NewEntry()
	sum1Entry := widget.NewEntry()
	sum2Entry := widget.NewEntry()
	sum3Entry := widget.NewEntry()
	sum4Entry := widget.NewEntry()
	sum1 := container.NewHBox(sum1Entry)
	sum2 := container.NewHBox(sum1Entry, sum2Entry)
	sum3 := container.NewHBox(sum1Entry, sum2Entry, sum3Entry)
	sum4 := container.NewHBox(sum1Entry, sum2Entry, sum3Entry, sum4Entry)
	indexesLabel := widget.NewLabel("Введите индексы для сумматоров через запятую")
	indexesBtnSum1 := widget.NewButton("Подтвердить", func() {
		countPriv, err := strconv.Atoi(entrySum.Text)
		if err != nil {
			dialog.ShowError(err, win)
		}
		var indexesOfSumsPriv [][]int
		idx := strings.Split(sum1Entry.Text, ",")
		var idxInt []int
		for _, elem := range idx {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		Code.Count = countPriv
		Code.IndexesOfSums = indexesOfSumsPriv
	})
	indexesBtnSum2 := widget.NewButton("Подтвердить", func() {
		countPriv, err := strconv.Atoi(entrySum.Text)
		if err != nil {
			dialog.ShowError(err, win)
		}
		var indexesOfSumsPriv [][]int
		idx1 := strings.Split(sum1Entry.Text, ",")
		idx2 := strings.Split(sum2Entry.Text, ",")
		var idxInt []int
		for _, elem := range idx1 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx2 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		Code.Count = countPriv
		Code.IndexesOfSums = indexesOfSumsPriv
	})
	indexesBtnSum3 := widget.NewButton("Подтвердить", func() {
		countPriv, err := strconv.Atoi(entrySum.Text)
		if err != nil {
			dialog.ShowError(err, win)
		}
		var indexesOfSumsPriv [][]int
		idx1 := strings.Split(sum1Entry.Text, ",")
		idx2 := strings.Split(sum2Entry.Text, ",")
		idx3 := strings.Split(sum3Entry.Text, ",")
		var idxInt []int
		for _, elem := range idx1 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx2 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx3 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		Code.Count = countPriv
		Code.IndexesOfSums = indexesOfSumsPriv
	})
	indexesBtnSum4 := widget.NewButton("Подтвердить", func() {
		countPriv, err := strconv.Atoi(entrySum.Text)
		if err != nil {
			dialog.ShowError(err, win)
		}
		var indexesOfSumsPriv [][]int
		idx1 := strings.Split(sum1Entry.Text, ",")
		idx2 := strings.Split(sum2Entry.Text, ",")
		idx3 := strings.Split(sum3Entry.Text, ",")
		idx4 := strings.Split(sum4Entry.Text, ",")
		var idxInt []int
		for _, elem := range idx1 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx2 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx3 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		idxInt = []int{}
		for _, elem := range idx4 {
			elemInt, err1 := strconv.Atoi(elem)
			if err1 != nil {
				dialog.ShowError(err1, win)
			}
			idxInt = append(idxInt, elemInt)
		}
		indexesOfSumsPriv = append(indexesOfSumsPriv, idxInt)
		Code.Count = countPriv
		Code.IndexesOfSums = indexesOfSumsPriv
	})

	labelWord := widget.NewLabel("Введите символы для кодирования")
	entryWord := widget.NewEntry()

	labelEncoded := widget.NewLabel("Закодированное слово")
	labelEncodedEnter := widget.NewLabel("")
	entryWordBtn := widget.NewButton("Закодировать", func() {
		Code.TextForCoding = entryWord.Text
		labelEncodedEnter.SetText(Code.Encode(win))
	})
	labelDecoded := widget.NewLabel("Декодированное слово")
	labelDecodedEnter := widget.NewLabel("")
	decodeBtn := widget.NewButton("Декодировать", func() {
		labelDecodedEnter.SetText(Code.Decode())
	})

	entrySumBtn := widget.NewButton("Подтвердить", func() {
		countPriv, err := strconv.Atoi(entrySum.Text)
		if err != nil {
			dialog.ShowError(err, win)
		}
		switch countPriv {
		case 1:
			win.Resize(fyne.NewSize(720, 350))
			win.SetContent(container.NewVBox(
				labelSum,
				entrySum,
				indexesLabel,
				sum1,
				indexesBtnSum1,
				labelWord,
				entryWord, entryWordBtn,
				labelEncoded,
				labelEncodedEnter,
				decodeBtn,
				labelDecoded,
				labelDecodedEnter,
			))
		case 2:
			win.Resize(fyne.NewSize(720, 350))
			win.SetContent(container.NewVBox(
				labelSum,
				entrySum,
				indexesLabel,
				sum2,
				indexesBtnSum2,
				labelWord,
				entryWord, entryWordBtn,
				labelEncoded,
				labelEncodedEnter,
				decodeBtn,
				labelDecoded,
				labelDecodedEnter,
			))
		case 3:
			win.Resize(fyne.NewSize(720, 350))
			win.SetContent(container.NewVBox(
				labelSum,
				entrySum,
				indexesLabel,
				sum3,
				indexesBtnSum3,
				labelWord,
				entryWord, entryWordBtn,
				labelEncoded,
				labelEncodedEnter,
				decodeBtn,
				labelDecoded,
				labelDecodedEnter,
			))
		case 4:
			win.Resize(fyne.NewSize(720, 350))
			win.SetContent(container.NewVBox(
				labelSum,
				entrySum,
				indexesLabel,
				sum4,
				indexesBtnSum4,
				labelWord,
				entryWord, entryWordBtn,
				labelEncoded,
				labelEncodedEnter,
				decodeBtn,
				labelDecoded,
				labelDecodedEnter,
			))
		}
	})

	content := container.NewVBox(
		labelSum,
		entrySum, entrySumBtn,
		labelWord,
		entryWord, entryWordBtn,
		labelEncoded,
		labelEncodedEnter,
		decodeBtn,
		labelDecoded,
		labelDecodedEnter,
	)
	win.SetContent(content)
	win.ShowAndRun()
}
