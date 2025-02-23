package utility

import "fmt"

type Rational struct {
	Numerator   int64
	Denominator int64
}

func (r Rational) ToFloat() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r Rational) Plus(other Rational) Rational {
	commonDenominator := LCM(r.Denominator, other.Denominator)

	denominatorGCD := GCD(r.Denominator, other.Denominator)

	numerator := r.Numerator*(other.Denominator/denominatorGCD) + other.Numerator*(r.Denominator/denominatorGCD)

	return Rational{
		Numerator:   numerator,
		Denominator: commonDenominator,
	}.ToLowest()
}

func (r Rational) Equals(other Rational) bool {
	return r.Numerator == other.Numerator && r.Denominator == other.Denominator
}

func (r Rational) ToString() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}

func (r Rational) ToLowest() Rational {
	commonFactors := GCD(r.Numerator, r.Denominator)

	return Rational{
		Numerator:   r.Numerator / commonFactors,
		Denominator: r.Denominator / commonFactors,
	}
}
