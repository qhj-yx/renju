package main

import (
	"gobangIcon"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	//a.Settings().SetTheme(theme.LightTheme())
	win := a.NewWindow("qianhejun")

	//blacker := widget.NewIcon(black)
	//whiter := widget.NewIcon(white)
	//chess := newMyIcon(chessboard)

	//content = fyne.NewContainerWithLayout(NewMyFixedLayout(fyne.NewSize(540, 540), fyne.NewSize(36, 36)), chess)

	// content.AddObject(blacker)
	// content.AddObject(whiter)

	win.SetContent(gobangIcon.NewChessboard())

	win.Resize(fyne.NewSize(540, 540))
	win.SetPadded(false)
	win.SetFixedSize(true)
	win.CenterOnScreen()
	win.ShowAndRun()
}
