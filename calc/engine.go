package calc

import (
	"fmt"
	"strings"
)

func New() *Calculator {
	c := &Calculator{
		variables: make(map[string]float64),
	}
	c.Clear()
	return c
}

func (c *Calculator) Clear() {
	c.variables = map[string]float64{
		"pi": 3.141592653589793,
		"e":  2.718281828459045,
	}
}

func (c *Calculator) Calculate(input string) (float64, error) {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return 0, nil
	}
	if strings.Contains(input, "=") && !strings.Contains(input, "==") &&
		!strings.Contains(input, "!=") && !strings.Contains(input, "<=") &&
		!strings.Contains(input, ">=") {
		return c.handleAssignment(input)
	}
	return c.runCalculation(input)
}

func (c *Calculator) handleAssignment(input string) (float64, error) {
	parts := strings.SplitN(input, "=", 2)
	varName := strings.TrimSpace(parts[0])
	res, err := c.runCalculation(parts[1])
	if err != nil {
		return 0, err
	}
	c.variables[varName] = res
	c.variables["ans"] = res
	return res, nil
}

func (c *Calculator) runCalculation(input string) (float64, error) {
	tokens := tokenize(input)
	if len(tokens) == 0 {
		return 0, fmt.Errorf("empty expression")
	}

	postfix := infixToPostfix(tokens)
	res, err := c.evaluatePostfix(postfix)
	if err == nil {
		c.variables["ans"] = res
	}
	return res, err
}
