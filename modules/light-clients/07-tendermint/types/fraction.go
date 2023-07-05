package types

import (
	ocmath "github.com/PikeEcosystem/tendermint/libs/math"
	"github.com/PikeEcosystem/tendermint/light"
)

// DefaultTrustLevel is the tendermint light client default trust level
var DefaultTrustLevel = NewFractionFromOc(light.DefaultTrustLevel)

// NewFractionFromOc returns a new Fraction instance from a ocmath.Fraction
func NewFractionFromOc(f ocmath.Fraction) Fraction {
	return Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}

// ToTendermint converts Fraction to ocmath.Fraction
func (f Fraction) ToTendermint() ocmath.Fraction {
	return ocmath.Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}
