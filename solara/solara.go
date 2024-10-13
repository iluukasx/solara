package solara

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := app.New()
	myApp.Settings().SetTheme(newSolaraTheme())
	myWindow := myApp.NewWindow("Solara")

	// Barra de título personalizada
	titleBar := canvas.NewRectangle(color.NRGBA{R: 75, G: 0, B: 130, A: 255})
	titleBar.SetMinSize(fyne.NewSize(0, 30))
	titleLabel := canvas.NewText("Solara", color.White)
	titleLabel.TextSize = 20
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewStack(titleBar, container.NewCenter(titleLabel))

	// Área de código
	codeEntry := widget.NewMultiLineEntry()
	codeEntry.SetText("// Tu código aquí")
	codeEntry.TextStyle = fyne.TextStyle{Monospace: true}
	codeScroll := container.NewScroll(codeEntry)

	// Barra de herramientas inferior
	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {})
	folderButton := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {})
	clearButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})
	clipboardButton := widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {})

	toolbar := container.NewHBox(
		layout.NewSpacer(),
		playButton,
		folderButton,
		clearButton,
		clipboardButton,
	)

	toolbarBg := canvas.NewRectangle(color.NRGBA{R: 40, G: 0, B: 70, A: 255})
	toolbarContainer := container.NewStack(toolbarBg, toolbar)

	// Contenedor principal
	content := container.NewBorder(
		titleContainer,
		toolbarContainer,
		nil,
		nil,
		codeScroll,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}

// Tema personalizado para Solara
type solaraTheme struct {
	fyne.Theme
}

func newSolaraTheme() fyne.Theme {
	return &solaraTheme{Theme: theme.DarkTheme()}
}

func (t *solaraTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 30, G: 30, B: 40, A: 255}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 200, G: 200, B: 255, A: 255}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 100, G: 0, B: 150, A: 255}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 120, G: 20, B: 180, A: 255}
	default:
		return t.Theme.Color(name, variant)
	}
}

func (t *solaraTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 12
	case theme.SizeNamePadding:
		return 6
	default:
		return t.Theme.Size(name)
	}
}
