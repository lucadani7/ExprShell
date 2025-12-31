package main

import (
	"ExprShell/calc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	myCalc := calc.New()
	fmt.Println("=== ExprShell v1.0 ===")
	for {
		fmt.Print("exprshell> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "exit" {
			break
		}
		formatted := myCalc.Format(line)
		fmt.Printf("   -> %s\n", formatted)
		result, err := myCalc.Calculate(line)
		if err != nil {
			fmt.Printf("Error calculating expression: %s\n", err)
		} else {
			if result == float64(int64(result)) {
				fmt.Printf("Result: %.0f\n", result)
			} else {
				fmt.Printf("Result: %.4f\n", result)
			}
		}
	}
}
