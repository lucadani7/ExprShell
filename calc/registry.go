package calc

import (
	"fmt"
	"math"
)

var unaryFunctions = map[string]UnaryFunc{
	// Trigonometrie
	"sin": func(v float64, c *Calculator) (float64, error) { return math.Sin(c.toRad(v)), nil },
	"cos": func(v float64, c *Calculator) (float64, error) { return math.Cos(c.toRad(v)), nil },
	"tan": func(v float64, c *Calculator) (float64, error) { return math.Tan(c.toRad(v)), nil },
	"abs": func(v float64, c *Calculator) (float64, error) { return math.Abs(v), nil },

	// Teoria Numerelor
	"is_prime": func(v float64, c *Calculator) (float64, error) {
		return boolToFloat(c.IsPrime(int(v))), nil
	},
	"phi": func(v float64, c *Calculator) (float64, error) {
		return c.Phi(v)
	},
	"next_prime": func(v float64, c *Calculator) (float64, error) {
		return c.NextPrime(v)
	},
	"factorial": func(v float64, c *Calculator) (float64, error) {
		res := 1.0
		for i := 2.0; i <= v; i++ {
			res *= i
		}
		return res, nil
	},
	"!": func(v float64, c *Calculator) (float64, error) {
		if v == 0 {
			return 1, nil
		}
		return 0, nil
	},
}

var binaryFunctions = map[string]BinaryFunc{
	"pow": func(v1, v2 float64, c *Calculator) (float64, error) {
		return math.Pow(v1, v2), nil
	},
	"root": func(v1, v2 float64, c *Calculator) (float64, error) {
		if v2 == 0 {
			return 0, fmt.Errorf("root of 0th")
		}
		return math.Pow(v1, 1/v2), nil
	},
	"mod": func(v1, v2 float64, c *Calculator) (float64, error) {
		return float64(int64(v1) % int64(v2)), nil
	},
}

var variadicFunctions = map[string]VariadicFunc{
	"sum": func(args []float64, c *Calculator) (float64, error) {
		res := 0.0
		for _, v := range args {
			res += v
		}
		return res, nil
	},
	"avg": func(args []float64, c *Calculator) (float64, error) {
		if len(args) == 0 {
			return 0, nil
		}
		s, _ := variadicFunctions["sum"](args, c)
		return s / float64(len(args)), nil
	},
	"max": func(args []float64, c *Calculator) (float64, error) {
		if len(args) == 0 {
			return 0, nil
		}
		m := args[0]
		for _, v := range args {
			if v > m {
				m = v
			}
		}
		return m, nil
	},
	"gcd": func(args []float64, c *Calculator) (float64, error) {
		if len(args) < 2 {
			return 0, fmt.Errorf("gcd needs at least 2 arguments")
		}
		res := int64(math.Abs(args[0]))
		for i := 1; i < len(args); i++ {
			res = gcdInt(res, int64(math.Abs(args[i])))
		}
		return float64(res), nil
	},
	"lcm": func(args []float64, c *Calculator) (float64, error) {
		if len(args) < 2 {
			return 0, fmt.Errorf("lcm needs at least 2 arguments")
		}
		res := int64(math.Abs(args[0]))
		for i := 1; i < len(args); i++ {
			res = lcmInt(res, int64(math.Abs(args[i])))
		}
		return float64(res), nil
	},
	"generate_primes": func(args []float64, c *Calculator) (float64, error) {
		if len(args) != 1 {
			return 0, fmt.Errorf("generate_primes needs 1 argument")
		}
		primes := c.Sieve(int(args[0]))
		return float64(len(primes)), nil
	},
}

func isVariadicFunc(name string) bool {
	_, ok := variadicFunctions[name]
	return ok
}

func isBinaryFunc(name string) bool {
	_, ok := binaryFunctions[name]
	return ok
}
