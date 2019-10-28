package ipam

import (
	"github.com/giantswarm/microerror"
)

var (
	maskTooBigError                  = microerror.New("mask too big")
	nilIPError                       = microerror.New("nil IP")
	ipNotContainedError              = microerror.New("IP not contained")
	maskIncorrectSizeError           = microerror.New("mask incorrect size")
	spaceExhaustedError              = microerror.New("space exhausted")
	incorrectNumberOfBoundariesError = microerror.New("incorrect number of boundaries")
	incorrectNumberOfFreeRangesError = microerror.New("incorrect number of free ranges")
	invalidConfigError               = microerror.New("invalid config")
)

// IsMaskTooBig asserts maskTooBigError.
func IsMaskTooBig(err error) bool {
	return microerror.Cause(err) == maskTooBigError
}

// IsNilIP asserts nilIPError.
func IsNilIP(err error) bool {
	return microerror.Cause(err) == nilIPError
}

// IsIPNotContained asserts ipNotContainedError.
func IsIPNotContained(err error) bool {
	return microerror.Cause(err) == ipNotContainedError
}

// IsMaskIncorrectSize asserts maskIncorrectSizeError.
func IsMaskIncorrectSize(err error) bool {
	return microerror.Cause(err) == maskIncorrectSizeError
}

// IsSpaceExhausted asserts spaceExhaustedError.
func IsSpaceExhausted(err error) bool {
	return microerror.Cause(err) == spaceExhaustedError
}

// IsIncorrectNumberOfBoundaries asserts incorrectNumberOfBoundariesError.
func IsIncorrectNumberOfBoundaries(err error) bool {
	return microerror.Cause(err) == incorrectNumberOfBoundariesError
}

// IsIncorrectNumberOfFreeRangesError asserts incorrectNumberOfFreeRangesError.
func IsIncorrectNumberOfFreeRanges(err error) bool {
	return microerror.Cause(err) == incorrectNumberOfFreeRangesError
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidParameterError = &microerror.Error{
	Kind: "invalid parameter",
}

// IsInvalidParameter asserts invalidParameterError.
func IsInvalidParameter(err error) bool {
	return microerror.Cause(err) == invalidParameterError
}
