package gui

import (
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// Список популярных плат. Можешь дополнить своими FQBN.
var boards = []string{
	"arduino:avr:uno",
	"arduino:avr:nano",
	"arduino:avr:mega",
	"arduino:avr:leonardo",
	"arduino:avr:micro",
	"esp32:esp32:esp32",
	"esp8266:esp8266:generic",
}

// UI — состояние главного окна.
type UI struct {
	win fyne.Window

	hexPath     string
	hexLabel    *widget.Label
	portSelect  *widget.Select
	boardSelect *widget.Select
	logArea     *widget.Entry
	flashBtn    *widget.Button

	mu sync.Mutex
}

// NewUI создаёт UI, привязанный к окну.
func NewUI(w fyne.Window) *UI {
	return &UI{win: w}
}

// Build возвращает корневой контейнер окна.
func (u *UI) Build() fyne.CanvasObject {
	u.hexLabel = widget.NewLabel("No file selected")
	u.hexLabel.Wrapping = fyne.TextWrapWord

	u.portSelect = widget.NewSelect([]string{}, nil)
	u.portSelect.PlaceHolder = "select port"

	u.boardSelect = widget.NewSelect(boards, nil)
	u.boardSelect.SetSelectedIndex(0)

	u.logArea = widget.NewMultiLineEntry()
	u.logArea.Wrapping = fyne.TextWrapWord
	u.logArea.Disable()

	u.flashBtn = widget.NewButton("Flash", u.onFlash)
	u.flashBtn.Importance = widget.HighImportance

	selectFileBtn := widget.NewButton("Select .hex", u.onSelectFile)
	refreshPortsBtn := widget.NewButton("Refresh", func() { u.refreshPorts() })

	form := container.NewVBox(
		widget.NewLabelWithStyle("Hex Loader", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),

		widget.NewLabel("Firmware (.hex):"),
		container.NewBorder(nil, nil, nil, selectFileBtn, u.hexLabel),

		widget.NewLabel("Serial port:"),
		container.NewBorder(nil, nil, nil, refreshPortsBtn, u.portSelect),

		widget.NewLabel("Board (FQBN):"),
		u.boardSelect,

		widget.NewSeparator(),
		u.flashBtn,

		widget.NewLabel("Log:"),
		container.NewScroll(u.logArea),
	)

	u.refreshPorts()
	return container.NewPadded(form)
}

func (u *UI) refreshPorts() {
	ports, err := ListPorts()
	if err != nil {
		u.log("Failed to list ports: " + err.Error())
		return
	}
	if len(ports) == 0 {
		u.log("No serial ports found")
		u.portSelect.Options = []string{}
		u.portSelect.Refresh()
		return
	}
	u.portSelect.Options = ports
	u.portSelect.Refresh()
	u.portSelect.SetSelectedIndex(0)
	u.log("Ports found: " + strings.Join(ports, ", "))
}

func (u *UI) onSelectFile() {
	fd := dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
		if err != nil {
			u.log("File dialog error: " + err.Error())
			return
		}
		if rc == nil {
			return
		}
		defer rc.Close()
		u.hexPath = rc.URI().Path()
		u.hexLabel.SetText(u.hexPath)
		u.log("Selected file: " + u.hexPath)
	}, u.win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".hex"}))
	fd.Show()
}

func (u *UI) onFlash() {
	u.flashBtn.Disable()
	cfg := FlashConfig{
		HexFile: u.hexPath,
		Port:    u.portSelect.Selected,
		FQBN:    u.boardSelect.Selected,
	}
	go func() {
		defer u.flashBtn.Enable()
		if err := Flash(cfg, u.log); err != nil {
			u.log("[ERR] " + err.Error())
		}
	}()
}

func (u *UI) log(s string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	cur := u.logArea.Text
	if cur != "" && !strings.HasSuffix(cur, "\n") {
		cur += "\n"
	}
	u.logArea.SetText(cur + s + "\n")
	u.logArea.CursorRow = strings.Count(u.logArea.Text, "\n")
}
