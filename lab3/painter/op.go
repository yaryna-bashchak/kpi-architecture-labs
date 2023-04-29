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
	XPOS1, YPOS1, XPOS2, YPOS2 int
}

func (op *BgRect) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.XPOS1, op.YPOS1, op.XPOS2, op.YPOS2), color.Black, screen.Src)
	return false
}

type Figure struct {
	XPOS, YPOS int
	C          color.RGBA
}

func (op *Figure) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.XPOS-150, op.YPOS-100, op.XPOS+150, op.YPOS), op.C, draw.Src)
	t.Fill(image.Rect(op.XPOS-50, op.YPOS, op.XPOS+50, op.YPOS+100), op.C, draw.Src)
	return false
}

type MoveOp struct {
	XPOS, YPOS int
	Figures    []Figure
}

func (op *MoveOp) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].XPOS += op.XPOS
		op.Figures[i].YPOS += op.YPOS
		op.Figures[i].Do(t)
	}
	return false
}

func ResetWindow(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}
