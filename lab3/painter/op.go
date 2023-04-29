package painter

import (
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRect struct {
	xPos1, yPos1, xPos2, yPos2 int
}

func (op *BgRect) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.xPos1, op.yPos1, op.xPos2, op.yPos2), color.Black, screen.Src)
	return false
}

type Figure struct {
	xPos, yPos int
	color      color.RGBA
}

func (op *Figure) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.xPos-150, op.yPos-100, op.xPos+150, op.yPos), op.color, draw.Src)
	t.Fill(image.Rect(op.xPos-50, op.yPos, op.xPos+50, op.yPos+100), op.color, draw.Src)
	return false
}

type Move struct {
	xPos, yPos int
	Figures    []Figure
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].xPos = op.xPos
		op.Figures[i].yPos = op.yPos
		op.Figures[i].Do(t)
	}
	return false
}

func ResetWindow(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}
