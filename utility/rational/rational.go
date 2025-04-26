package rational

import (
	"euler/utility/integers"
	"fmt"
)

type Rational struct {
	Numerator   int64
	Denominator int64
}

func (r Rational) ToFloat() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r Rational) Plus(other Rational) Rational {
	commonDenominator := integers.LCM(r.Denominator, other.Denominator)

	denominatorGCD := integers.GCD(r.Denominator, other.Denominator)

	numerator := r.Numerator*(other.Denominator/denominatorGCD) + other.Numerator*(r.Denominator/denominatorGCD)

	return Rational{
		Numerator:   numerator,
		Denominator: commonDenominator,
	}.ToLowest()
}

func (r Rational) Mult(other Rational) Rational {
	return Rational{
		Numerator:   r.Numerator * other.Numerator,
		Denominator: r.Denominator * other.Denominator,
	}.ToLowest()
}

func (r Rational) Div(other Rational) Rational {
	return r.Mult(other.Reciprocal())
}

func (r Rational) Reciprocal() Rational {
	return Rational{
		Numerator:   r.Denominator,
		Denominator: r.Numerator,
	}
}

func RationalFromInt(n int64) Rational {
	return Rational{
		Numerator:   n,
		Denominator: 1,
	}
}

func (r Rational) Equals(other Rational) bool {
	return r.Numerator == other.Numerator && r.Denominator == other.Denominator
}

func (r Rational) ToString() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}

func (r Rational) ToLowest() Rational {
	commonFactors := integers.GCD(r.Numerator, r.Denominator)

	return Rational{
		Numerator:   r.Numerator / commonFactors,
		Denominator: r.Denominator / commonFactors,
	}
}
