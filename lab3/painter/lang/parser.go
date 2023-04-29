package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/yaryna-bashchak/kpi-architecture-labs/lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	lastBgColor painter.Operation
	lastBgRect  *painter.BgRect
	figures     []painter.Figure
	moveOps     []painter.Operation
	updateOp    painter.Operation
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		err := p.parse(commandLine) // parse the line to get Operation
		if err != nil {
			return nil, err
		}
	}
	return p.finalizeResult(), nil
}

func (p *Parser) finalizeResult() []painter.Operation {
	var res []painter.Operation
	if p.lastBgColor != nil {
		res = append(res, p.lastBgColor)
	}
	if p.lastBgRect != nil {
		res = append(res, p.lastBgRect)
	}
	if len(p.moveOps) != 0 {
		res = append(res, p.moveOps...)
	}
	if len(p.figures) != 0 {
		println(len(p.figures))
		for _, figure := range p.figures {
			res = append(res, &figure)
		}
	}

	if p.updateOp != nil {
		res = append(res, p.updateOp)
	}

	return res
}

func (p *Parser) resetState() {
	p.lastBgColor = nil
	p.lastBgRect = nil
	p.figures = nil
	p.moveOps = nil
	p.updateOp = nil
}

func (p *Parser) parse(commandLine string) error {
	parts := strings.Split(commandLine, " ")
	instruction := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	var intArgs []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			intArgs = append(intArgs, i)
		}
	}

	switch instruction {
	case "white":
		p.lastBgColor = painter.OperationFunc(painter.WhiteFill)
	case "green":
		print()
		p.lastBgColor = painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		p.lastBgRect = &painter.BgRect{XPOS1: intArgs[0], YPOS1: intArgs[1], XPOS2: intArgs[2], YPOS2: intArgs[3]}
	case "figure":
		col := color.RGBA{R: 219, G: 208, B: 48, A: 1}
		figure := painter.Figure{XPOS: intArgs[0], YPOS: intArgs[1], C: col}
		p.figures = append(p.figures, figure)
	case "move":
		moveOp := painter.MoveOp{XPOS: intArgs[0], YPOS: intArgs[1], Figures: p.figures}
		p.moveOps = append(p.moveOps, &moveOp)

	case "reset":
		p.resetState()
		p.lastBgColor = painter.OperationFunc(painter.ResetWindow)
	case "update":
		p.updateOp = painter.UpdateOp
	default:
		return fmt.Errorf("could not parse command %v", commandLine)
	}
	return nil
}
