package GUI

import (
	"TheoryOfInformation/pkg/code"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GuiLab2() {
	a := app.New()
	win := a.NewWindow("Вторая лаба")
	win.Resize(fyne.NewSize(500, 50))

	label := widget.NewLabel("")
	entry := widget.NewEntry()
	entry.SetText("Введите поле")
	btnRBracket := widget.NewButton(")", func() {
		label.SetText(label.Text + ")")
	})
	btnRBracket.Resize(fyne.NewSize(30, 30))
	btnLBracket := widget.NewButton("(", func() {
		label.SetText(label.Text + "(")
	})
	btnLBracket.Resize(fyne.NewSize(30, 30))
	btnOne := widget.NewButton("1", func() {
		label.SetText(label.Text + "1")
	})
	btnTwo := widget.NewButton("2", func() {
		label.SetText(label.Text + "2")
	})
	btnThree := widget.NewButton("3", func() {
		label.SetText(label.Text + "3")
	})
	btnFour := widget.NewButton("4", func() {
		label.SetText(label.Text + "4")
	})
	btnFive := widget.NewButton("5", func() {
		label.SetText(label.Text + "5")
	})
	btnSix := widget.NewButton("6", func() {
		label.SetText(label.Text + "6")
	})
	btnSeven := widget.NewButton("7", func() {
		label.SetText(label.Text + "7")
	})
	btnEight := widget.NewButton("8", func() {
		label.SetText(label.Text + "8")
	})
	btnNine := widget.NewButton("9", func() {
		label.SetText(label.Text + "9")
	})
	btnZero := widget.NewButton("0", func() {
		label.SetText(label.Text + "0")
	})
	btnPlus := widget.NewButton("+", func() {
		label.SetText(label.Text + "+")
	})
	btnMinus := widget.NewButton("-", func() {
		label.SetText(label.Text + "-")
	})
	btnMultiply := widget.NewButton("*", func() {
		label.SetText(label.Text + "*")
	})
	btnDivide := widget.NewButton("/", func() {
		label.SetText(label.Text + "/")
	})
	btnEqual := widget.NewButton("=", func() {
		answer := code.Lab2(label.Text, entry.Text)
		label.SetText(answer)
	})
	btnClear := widget.NewButton("C", func() {
		label.SetText("")
	})
	entryBox := container.NewVBox(
		label,
		entry,
	)
	btnBox := container.NewHBox(
		btnOne,
		btnTwo,
		btnThree,
		btnFour,
		btnFive,
		btnSix,
		btnSeven,
		btnEight,
		btnNine,
		btnZero,
		btnPlus,
		btnMinus,
		btnMultiply,
		btnDivide,
		btnLBracket,
		btnRBracket,
		btnEqual,
		btnClear,
	)

	winBox := container.NewVBox(entryBox, btnBox)
	win.SetContent(winBox)
	win.ShowAndRun()
}
