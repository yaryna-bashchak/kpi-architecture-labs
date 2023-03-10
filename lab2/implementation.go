package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	if err == nil {
		return true
	} else {
		return false
	}
}

func PostfixToInfix(postfix string) (string, error) {
	tokens := strings.Split(postfix, " ")
	array := []string{}
	for _, token := range tokens {
		if isOperator(token) && len(array) >= 2 {
			operand1 := array[len(array)-2]
			operand2 := array[len(array)-1]
			array = array[:len(array)-2]

			if token == "*" || token == "/" || token == "^" {
				if !isNumber(operand1) {
					operand1 = fmt.Sprintf("(%s)", operand1)
				}
				if !isNumber(operand2) {
					operand2 = fmt.Sprintf("(%s)", operand2)
				}
			}
			array = append(array, fmt.Sprintf("%s %s %s", operand1, token, operand2))
		} else if isNumber(token) {
			array = append(array, token)
		} else if len(array) < 2 {
			return "", fmt.Errorf("Invalid postfix expression: there are too many arithmetic operands")
		} else {
			return "", fmt.Errorf(fmt.Sprintf("Invalid symbol: %s", token))
		}
	}

	if len(array) != 1 {
		return "", fmt.Errorf("Invalid postfix expression: not the whole expression is related by an arithmetic operand")
	}

	return array[0], nil
}
