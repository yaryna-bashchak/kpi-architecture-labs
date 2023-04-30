package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaryna-bashchak/kpi-architecture-labs/lab3/painter"
	"github.com/yaryna-bashchak/kpi-architecture-labs/lab3/painter/lang"
	"image/color"
	"reflect"
	"strings"
	"testing"
)

// Test_parse_struct tests whether the lang.Parser can correctly parse commands
// and return the corresponding painter.Operation structs.
func Test_parse_struct(t *testing.T) {
	// Define test data as a slice of structs
	tests := []struct {
		command string          // Command string to parse
		op      painter.Operation // Expected operation struct
	}{
		{
			command: "bgrect 0 0 80 80",
			op:      &painter.BgRect{XPOS1: 0, YPOS1: 0, XPOS2: 80, YPOS2: 80},
		},
		{
			command: "figure 100 100",
			op:      &painter.Figure{XPOS: 100, YPOS: 100, C: color.RGBA{R: 255, G: 255, B: 0, A: 1}},
		},
		{
			command: "move 200 300",
			op:      &painter.MoveOp{XPOS: 200, YPOS: 300},
		},
		{
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			command: "invalidcommand",
			op:      nil,
		},
	}

	// Iterate over the test data
	for _, tt := range tests {
		// Run subtest with the name of the command string
		t.Run(tt.command, func(t *testing.T) {
			// Create a new parser instance
			parser := &lang.Parser{}

			// Parse the command string and get the parsed operations and error
			op, err := parser.Parse(strings.NewReader(tt.command))

			// If an error occurred, assert that the expected operation is nil
			if err != nil {
				assert.Nil(t, tt.op)
			} else {
				// If no error, assert that the type of the returned operation matches the type of the expected operation
				assert.Equal(t, reflect.TypeOf(tt.op), reflect.TypeOf(op[1]))

				// If the expected operation is not nil, assert that the returned operation is equal to the expected operation
				if tt.op != nil {
					assert.Equal(t, tt.op, op[1])
				}
			}
		})
	}
}

// Test_parse_func tests whether the lang.Parser can correctly parse commands
// and return the corresponding painter.Operation functions.
func Test_parse_func(t *testing.T) {
	// Define test data as a slice of structs
	tests := []struct {
		command string          // Command string to parse
		op      painter.Operation // Expected operation function
	}{
		{
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			command: "reset",
			op:      painter.OperationFunc(painter.ResetWindow),
		},
	}

	// Create a new parser instance
	parser := &lang.Parser{}

	// Iterate over the test data
	for _, tt := range tests {
		// Run subtest with the name of the command string
    	t.Run(tt.command, func(t *testing.T) {
    		op, _ := parser.Parse(strings.NewReader(tt.command))
    		assert.Equal(t, reflect.TypeOf(tt.op), reflect.TypeOf(op[0]))
    	})
    }
}
