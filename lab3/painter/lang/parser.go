package lang

import (
	"io"
	"bufio"
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/yaryna-bashchak/kpi-architecture-labs/blob/main/lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	var res []painter.Operation

	for scanner.Scan() {
		commandLine := scanner.Text()
		op := parse(commandLine) // parse the line to get Operation
		if op == nil {
			return nil, fmt.Errorf("failed to parse command: %s", commandLine)
		}
		if bgRect, ok := op.(*painter.BgRect); ok {
			for i, oldOp := range res {
				if _, ok := oldOp.(*painter.BgRect); ok {
					res[i] = bgRect
					break
				}
			}
		}
		res = append(res, op)
	}
	return res, nil
}

func parse(commandLine string) painter.Operation {
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

	var figureOps []painter.Figure

	switch instruction {
	case "white":
		return painter.OperationFunc(painter.WhiteFill)
	case "green":
		return painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		return &painter.BgRect{xPos1: intArgs[0], yPos1: intArgs[1], xPos2: intArgs[2], yPos2: intArgs[3]}
	case "figure":
		col := color.RGBA{R: 255, G: 255, B: 0, A: 1}
		figure := painter.Figure{xPos: intArgs[0], yPos: intArgs[1], color: col}
		figureOps = append(figureOps, figure)
		return &figure
	case "move":
		return &painter.Move{xPos: intArgs[0], yPos: intArgs[1], figure: figureOps}
	case "reset":
		figureOps = figureOps[0:0]
		return painter.OperationFunc(painter.ResetWindow)
	case "update":
		return painter.UpdateOp
	}
	return nil
}