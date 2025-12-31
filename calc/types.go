package calc

type Mode int

const (
	INFIX Mode = iota
	PREFIX
	POSTFIX
)

type Calculator struct {
	variables  map[string]float64
	useDegrees bool
	inputMode  Mode
	lastPrimes []int
}

type UnaryFunc func(float64, *Calculator) (float64, error)
type BinaryFunc func(float64, float64, *Calculator) (float64, error)
type VariadicFunc func([]float64, *Calculator) (float64, error)

const epsilon = 1e-9
