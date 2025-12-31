package calc

import (
	"fmt"
	"math"
	"sort"

	"github.com/bits-and-blooms/bitset"
)

func gcdInt(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmInt(a, b int64) int64 {
	return a * b / gcdInt(a, b)
}

func (c *Calculator) IsPrime(target int) bool {
	if target < 2 {
		return false
	}
	if target == 2 {
		return true
	}
	if target%2 == 0 {
		return false
	}
	if len(c.lastPrimes) > 0 {
		maxInCache := c.lastPrimes[len(c.lastPrimes)-1]
		if target <= maxInCache {
			index := sort.SearchInts(c.lastPrimes, target)
			return index < len(c.lastPrimes) && c.lastPrimes[index] == target
		}
	}
	limit := int(math.Sqrt(float64(target)))
	for i := 3; i <= limit; i += 2 {
		if target%i == 0 {
			return false
		}
	}
	return true
}

func (c *Calculator) NextPrime(n float64) (float64, error) {
	candidate := int(n) + 1
	if candidate < 3 {
		return 2, nil
	}
	if candidate%2 == 0 {
		candidate++
	}
	for {
		if c.IsPrime(candidate) {
			return float64(candidate), nil
		}
		candidate += 2
	}
}

func (c *Calculator) Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	result := []int{2}
	if limit == 2 {
		c.lastPrimes = result
		return result
	}
	numOdd := (limit-3)/2 + 1
	if numOdd <= 0 {
		c.lastPrimes = result
		return result
	}
	b := bitset.New(uint(numOdd))
	for pIdx := 0; ; pIdx++ {
		p := 2*pIdx + 3
		if p*p > limit {
			break
		}
		if !b.Test(uint(pIdx)) {
			start := (p*p - 3) / 2
			step := p
			for i := start; i < numOdd; i += step {
				b.Set(uint(i))
			}
		}
	}
	for i := 0; i < numOdd; i++ {
		if !b.Test(uint(i)) {
			result = append(result, 2*i+3)
		}
	}
	c.lastPrimes = result
	return result
}

func (c *Calculator) Phi(n float64) (float64, error) {
	if n == 0 || n == 1 {
		return n, nil
	}
	factors, err := c.Factorize(n)
	if err != nil {
		return 0, err
	}
	res := int64(math.Abs(n))
	for p := range factors {
		res = res / p * (p - 1)
	}
	return float64(res), nil
}

func (c *Calculator) Factorize(n float64) (map[int64]int64, error) {
	val := int64(math.Abs(n))
	if val < 2 {
		return nil, fmt.Errorf("factorization needs a number bigger than 1")
	}
	factors := make(map[int64]int64)
	d := val
	if len(c.lastPrimes) > 0 {
		for _, p := range c.lastPrimes {
			p64 := int64(p)
			if p64*p64 > d {
				break
			}
			for d%p64 == 0 {
				factors[p64]++
				d /= p64
			}
		}
	}
	if d > 1 {
		start := int64(3)
		if len(c.lastPrimes) > 0 {
			start = int64(c.lastPrimes[len(c.lastPrimes)-1]) + 2
		}
		for d%2 == 0 {
			factors[2]++
			d /= 2
		}

		for i := start; i*i <= d; i += 2 {
			for d%i == 0 {
				factors[i]++
				d /= i
			}
		}
		if d > 1 {
			factors[d]++
			d = 1
		}
	}
	return factors, nil
}
