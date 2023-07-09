package models

// Calculation structure containing the input data, intermediate and final calculation results
type Calculation struct {
	// Base calculation input
	// required: true
	// min: 0
	// example: 1
	Base uint64 `json:"base"`

	// Result calculation result
	// required: true
	// min: 0
	// example: 1
	Result uint64 `json:"result"`
}

// CalculationCallback function for processing calculation steps (fill structure fields)
type CalculationCallback func(input Calculation) (Calculation, error)

func DefaultCalculation(input Calculation) (Calculation, error) {
	input.Result *= input.Base
	return input, nil
}

type CalculatorExternal interface {
	Calculate(input Calculation) (Calculation, error)
}
