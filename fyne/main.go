package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type customTheme struct {
	fyne.Theme
}

func (t customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameFocus {
		return color.RGBA{225, 225, 225, 255}
	}
	return t.Theme.Color(name, variant)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&customTheme{theme.LightTheme()})

	w := a.NewWindow("£#f4!3r/3")
	w.Resize(fyne.NewSize(410, 330))

	usernameEntry := widget.NewEntry()
	usernameEntry.PlaceHolder = "Username"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Password"

	keyEntry := widget.NewEntry()
	keyEntry.PlaceHolder = "Key"
	keyEntry.Hide()

	var selectAction *widget.Select
	var progressBar *widget.ProgressBar
	var RegisterBtn *widget.Button
	var LoginBtn *widget.Button

	var createLoginScreen func() fyne.CanvasObject
	var createMainScreen func(username string) fyne.CanvasObject

	createMainScreen = func(username string) fyne.CanvasObject {
		title := canvas.NewText("The New Bypass Software", color.NRGBA{R: 0, G: 102, B: 204, A: 255})
		title.TextSize = 24
		title.TextStyle = fyne.TextStyle{Bold: true}
		title.Alignment = fyne.TextAlignCenter

		welcomeLabel := canvas.NewText("Welcome back, "+username, color.NRGBA{R: 60, G: 60, B: 60, A: 255})
		welcomeLabel.TextStyle = fyne.TextStyle{Italic: true}
		welcomeLabel.Alignment = fyne.TextAlignCenter

		versionLabel := widget.NewLabelWithStyle("Version: 0.0.1", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

		statusText := canvas.NewText("Status: Undetected", color.NRGBA{R: 0, G: 180, B: 0, A: 255})
		statusText.TextStyle = fyne.TextStyle{Bold: true}
		statusText.Alignment = fyne.TextAlignCenter

		injectBtn := widget.NewButton("Inject", nil)
		cleanBtn := widget.NewButton("Clean", nil)
		destructBtn := widget.NewButton("Destruct", nil)

		injectBtn.Hide()
		cleanBtn.Hide()
		destructBtn.Hide()

		injectBtn.Resize(fyne.NewSize(400, injectBtn.MinSize().Height))
		cleanBtn.Resize(fyne.NewSize(400, cleanBtn.MinSize().Height))
		destructBtn.Resize(fyne.NewSize(400, destructBtn.MinSize().Height))

		progressBar = widget.NewProgressBar()
		progressBar.Hide()

		executeAction := func(btn *widget.Button) {
			progressBar.Show()
			btn.Disable()
			go func() {
				for i := 0; i < 100; i++ {
					progressBar.SetValue(float64(i) / 100)
					time.Sleep(50 * time.Millisecond)
				}
				progressBar.Hide()
				selectAction.SetSelected("")
				btn.Hide()
				btn.Enable()
			}()
		}

		injectBtn.OnTapped = func() { executeAction(injectBtn) }
		cleanBtn.OnTapped = func() { executeAction(cleanBtn) }
		destructBtn.OnTapped = func() { executeAction(destructBtn) }

		selectAction = widget.NewSelect([]string{"Inject", "Clean", "Destruct", "Back"}, func(s string) {
			injectBtn.Hide()
			cleanBtn.Hide()
			destructBtn.Hide()
			switch s {
			case "Inject":
				injectBtn.Show()
			case "Clean":
				cleanBtn.Show()
			case "Destruct":
				destructBtn.Show()
			case "Back":
				w.SetContent(createLoginScreen())
			}
		})
		selectAction.PlaceHolder = "Select Action"

		content := container.NewVBox(
			container.NewVBox(
				title,
				welcomeLabel,
				versionLabel,
				statusText,
			),
			container.NewCenter(selectAction),
			injectBtn,
			cleanBtn,
			destructBtn,
			progressBar,
		)

		return container.NewBorder(nil, nil, nil, nil, content)
	}

	createLoginScreen = func() fyne.CanvasObject {
		LoginBtn = widget.NewButton("Login", func() {
			w.SetContent(createMainScreen(usernameEntry.Text))
		})

		RegisterBtn = widget.NewButton("Register", func() {
			if keyEntry.Hidden {
				keyEntry.Show()
				RegisterBtn.SetText("Confirm Registration")
			} else {
				keyEntry.Hide()
				RegisterBtn.SetText("Register")
			}
		})

		buttonContainer := container.NewHBox(
			layout.NewSpacer(),
			LoginBtn,
			RegisterBtn,
		)

		bottomContainer := container.NewVBox(
			usernameEntry,
			passwordEntry,
			keyEntry,
			buttonContainer,
		)

		return container.NewVBox(
			layout.NewSpacer(),
			bottomContainer,
		)
	}

	w.SetContent(createLoginScreen())

	w.ShowAndRun()
}
