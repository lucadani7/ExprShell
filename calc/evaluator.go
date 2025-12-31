package calc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Calculator) evaluatePostfix(tokens []string) (float64, error) {
	var stack []float64
	for _, token := range tokens {
		if val, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, val)
			continue
		}
		if val, ok := c.variables[token]; ok {
			stack = append(stack, val)
			continue
		}
		if strings.Contains(token, "|") {
			parts := strings.Split(token, "|")
			name := parts[0]
			count, _ := strconv.Atoi(parts[1])
			if len(stack) < count {
				return 0, fmt.Errorf("insufficient arguments for %s", name)
			}
			args := make([]float64, count)
			for i := count - 1; i >= 0; i-- {
				args[i] = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			fn, ok := variadicFunctions[name]
			if !ok {
				return 0, fmt.Errorf("unknown variadic function: %s", name)
			}
			res, err := fn(args, c)
			if err != nil {
				return 0, err
			}
			stack = append(stack, res)
			continue
		}
		if fn, ok := unaryFunctions[token]; ok {
			if len(stack) < 1 {
				return 0, fmt.Errorf("lack argument for %s", token)
			}
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res, err := fn(val, c)
			if err != nil {
				return 0, err
			}
			stack = append(stack, res)
			continue
		}
		if isOperatorStr(token) || isBinaryFunc(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("operanzi insuficienți pentru %s", token)
			}
			v2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			v1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var res float64
			var err error

			if fn, ok := binaryFunctions[token]; ok {
				res, err = fn(v1, v2, c)
			} else {
				res, err = c.execBinaryOp(token, v1, v2)
			}

			if err != nil {
				return 0, err
			}
			stack = append(stack, res)
			continue
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("evaluation error: invalid stack at the end")
	}
	return stack[0], nil
}
func (c *Calculator) execBinaryOp(op string, v1, v2 float64) (float64, error) {
	const epsilon = 1e-9
	switch op {
	case "+":
		return v1 + v2, nil
	case "-":
		return v1 - v2, nil
	case "*":
		return v1 * v2, nil
	case "/":
		if v2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return v1 / v2, nil
	case "^":
		return math.Pow(v1, v2), nil
	case "==":
		return boolToFloat(math.Abs(v1-v2) < epsilon), nil
	case "!=":
		return boolToFloat(math.Abs(v1-v2) >= epsilon), nil
	case "<":
		return boolToFloat(v1 < v2), nil
	case ">":
		return boolToFloat(v1 > v2), nil
	case "<=":
		return boolToFloat(v1 <= v2), nil
	case ">=":
		return boolToFloat(v1 >= v2), nil
	case "&&":
		return boolToFloat(v1 != 0 && v2 != 0), nil
	case "||":
		return boolToFloat(v1 != 0 || v2 != 0), nil
	}
	return 0, fmt.Errorf("unknown operator: %s", op)
}

func boolToFloat(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}
