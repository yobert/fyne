package widget

import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

const (
	textAreaSpaceSymbol   = '·'
	textAreaTabSymbol     = '→'
	textAreaNewLineSymbol = '↵'
)

// TextGrid is a monospaced grid of characters.
// This is designed to be used by a text editor or advanced test presentation.
type TextGrid struct {
	baseWidget
	textHandler

	LineNumbers bool
	Whitespace  bool
}

// Resize sets a new size for a widget.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (t *TextGrid) Resize(size fyne.Size) {
	t.resize(size, t)
}

// Move the widget to a new position, relative to it's parent.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (t *TextGrid) Move(pos fyne.Position) {
	t.move(pos, t)
}

// MinSize returns the smallest size this widget can shrink to
func (t *TextGrid) MinSize() fyne.Size {
	return t.minSize(t)
}

// Show this widget, if it was previously hidden
func (t *TextGrid) Show() {
	t.show(t)
}

// Hide this widget, if it was previously visible
func (t *TextGrid) Hide() {
	t.hide(t)
}

// CreateRenderer is a private method to Fyne which links this widget to it's renderer
func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer {
	render := &textGridRender{text: t}
	t.updated = render.update
	t.updateRowBounds()
	render.update()

	cell := canvas.NewText("M", color.White)
	cell.TextStyle.Monospace = true
	render.cellSize = cell.MinSize()

	return render
}

// NewTextGrid creates a new textgrid widget with the specified string content.
func NewTextGrid(content string) *TextGrid {
	handler := textHandler{buffer: []rune(content)}
	grid := &TextGrid{textHandler: handler}
	return grid
}

type textGridRender struct {
	text *TextGrid

	cols, rows int

	cellSize fyne.Size
	objects  []fyne.CanvasObject
}

func (t *textGridRender) appendTextCell(str rune) {
	text := canvas.NewText(string(str), theme.TextColor())
	text.TextStyle.Monospace = true

	t.objects = append(t.objects, text)
}

func (t *textGridRender) setCellRune(str rune, pos int) {
	text := t.objects[pos].(*canvas.Text)
	text.Text = string(str)

	if str == textAreaSpaceSymbol || str == textAreaTabSymbol || str == textAreaNewLineSymbol {
		text.Color = theme.PlaceHolderColor()
	}
}

func (t *textGridRender) update() {
	t.ensureGrid()
	t.refreshGrid()
}

func (t *textGridRender) ensureGrid() {
	t.cols = t.text.maxCols
	if t.text.Whitespace {
		t.cols++ // 1 more for newline option
	}
	if t.text.LineNumbers {
		t.cols += t.lineCountWidth() + 1
	}
	t.rows = t.text.rows()

	cellCount := t.cols * t.rows
	if len(t.objects) == cellCount {
		return
	}
	for i := len(t.objects); i < cellCount; i++ {
		t.appendTextCell(' ')
	}
}

func (t *textGridRender) refreshGrid() {
	line := 1
	x := 0

	for _, bound := range t.text.rowBounds {
		i := 0
		if t.text.LineNumbers {
			lineStr := []rune(fmt.Sprintf("%d", line))
			for c := 0; c < len(lineStr); c++ {
				t.setCellRune(lineStr[c], x)
				i++
				x++
			}
			for ; i < t.lineCountWidth(); i++ {
				t.setCellRune(' ', x)
				x++
			}

			t.setCellRune(' ', x)
			i++
			x++
		}
		for j := bound[0]; j < bound[1]; j++ {
			r := t.text.buffer[j]
			if t.text.Whitespace && r == ' ' {
				r = textAreaSpaceSymbol
			}
			t.setCellRune(r, x)
			i++
			x++
		}
		if t.text.Whitespace {
			t.setCellRune(textAreaNewLineSymbol, x)
			i++
			x++
		}
		for ; i < t.cols; i++ {
			t.setCellRune(' ', x)
			x++
		}

		line++
	}
}

func (t *textGridRender) lineCountWidth() int {
	return len(fmt.Sprintf("%d", t.text.rows()+1))
}

func (t *textGridRender) Layout(size fyne.Size) {
	i := 0
	cellPos := fyne.NewPos(0, 0)
	for y := 0; y < t.rows; y++ {
		for x := 0; x < t.cols; x++ {
			t.objects[i].Move(cellPos)

			cellPos.X += t.cellSize.Width
			i++
		}

		cellPos.X = 0
		cellPos.Y += t.cellSize.Height
	}
}

func (t *textGridRender) MinSize() fyne.Size {
	return fyne.NewSize(t.cellSize.Width*t.cols, t.cellSize.Height*t.rows)
}

func (t *textGridRender) Refresh() {
	t.refreshGrid()
}

func (t *textGridRender) ApplyTheme() {
}

func (t *textGridRender) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (t *textGridRender) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *textGridRender) Destroy() {
}
