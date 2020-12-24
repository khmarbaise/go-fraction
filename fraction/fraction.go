package fraction

import (
	"log"
	"math/big"
)

type Fraction struct {
	numerator   int
	denominator int
}

func calculateGCD(first int, second int) (gcd int) {
	var x, y, a, b *big.Int

	x = big.NewInt(0)
	y = big.NewInt(0)

	a = big.NewInt(int64(first))
	b = big.NewInt(int64(second))

	return int(new(big.Int).GCD(x, y, a, b).Int64())
}

func signum(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return -1
	}
	return +1
}

func (f Fraction) normalize() (result Fraction) {
	if f.denominator == 0 {
		//TODO: What is the best way? via Panic or returning an error?
		log.Panic("denominator is not allowed to be zero.")
	}

	sign := signum(f.denominator) * signum(f.numerator)
	gcd := calculateGCD(f.numerator, f.denominator)

	result.numerator = sign * (signum(f.numerator) * f.numerator) / gcd
	result.denominator = signum(f.denominator) * f.denominator / gcd
	return
}

func (f Fraction) Plus(summand Fraction) (result Fraction) {
	tmp := f.normalize()
	if tmp.denominator == summand.denominator {
		result = Fraction{
			tmp.numerator + summand.numerator,
			tmp.denominator,
		}
	} else {
		denominator := tmp.denominator * summand.denominator
		result = Fraction{
			tmp.numerator*summand.denominator + summand.numerator*tmp.denominator,
			denominator,
		}
	}
	return result.normalize()
}

func (f Fraction) Minus(subtrahend Fraction) (result Fraction) {
	tmp := f.normalize()
	if tmp.denominator == subtrahend.denominator {
		result = Fraction{
			tmp.numerator - subtrahend.numerator,
			tmp.denominator,
		}
	} else {
		denominator := tmp.denominator * subtrahend.denominator
		result = Fraction{
			tmp.numerator*subtrahend.denominator - subtrahend.numerator*tmp.denominator,
			denominator,
		}
	}
	return result.normalize()
}
