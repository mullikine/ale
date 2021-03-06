package arity

import (
	"errors"
	"fmt"

	"github.com/kode4food/ale/data"
)

// Error messages
const (
	BadFixedArity   = "expected %d arguments, got %d"
	BadMinimumArity = "expected at least %d arguments, got %d"
	BadRangedArity  = "expected between %d and %d arguments, got %d"
)

// OrMore is the constant used when you want to tell MakeChecker
// to generate a minimum arity checker
const OrMore = -1

// MakeChecker produces an arity checker based on its parameters
func MakeChecker(arity ...int) data.ArityChecker {
	al := len(arity)
	switch {
	case al == 0:
		return nil
	case al > 2:
		panic(errors.New("too many arity check arguments"))
	case al == 1 || arity[0] == arity[1]:
		return MakeFixedChecker(arity[0])
	case al == 2 && arity[1] == OrMore:
		return MakeMinimumChecker(arity[0])
	default:
		return MakeRangedChecker(arity[0], arity[1])
	}
}

// AssertFixed explodes if a fixed arity check fails
func AssertFixed(fixed, count int) int {
	if err := MakeFixedChecker(fixed)(count); err != nil {
		panic(err)
	}
	return count
}

// MakeFixedChecker generates a fixed arity checker
func MakeFixedChecker(fixed int) data.ArityChecker {
	return func(count int) error {
		if count != fixed {
			return fmt.Errorf(BadFixedArity, fixed, count)
		}
		return nil
	}
}

// AssertMinimum explodes if a fixed arity check fails
func AssertMinimum(min, count int) int {
	if err := MakeMinimumChecker(min)(count); err != nil {
		panic(err)
	}
	return count
}

// MakeMinimumChecker generates a minimum arity checker
func MakeMinimumChecker(min int) data.ArityChecker {
	return func(count int) error {
		if count < min {
			return fmt.Errorf(BadMinimumArity, min, count)
		}
		return nil
	}
}

// AssertRanged explodes if a fixed arity check fails
func AssertRanged(min, max, count int) int {
	if err := MakeRangedChecker(min, max)(count); err != nil {
		panic(err)
	}
	return count
}

// MakeRangedChecker generates a ranged arity checker
func MakeRangedChecker(min, max int) data.ArityChecker {
	return func(count int) error {
		if count < min || count > max {
			return fmt.Errorf(BadRangedArity, min, max, count)
		}
		return nil
	}
}
