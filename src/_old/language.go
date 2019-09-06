package main

import "math"
import "fmt"

variableSymbols := []string {}
binaryOperatorSymbols := [8]string {"+", "-", "*", "/", "%", "^", "min", "max"}
unaryOperatorSymbols := [8]string {"e^", "ln", "abs", "sin", "cos", "tan", "2^", "N", "P"}

// **********************
// *  BINARY OPERATORS  *
// **********************
func Addition(x1, x2 float64) float64 {
	return x1 + x2
}

func Subtraction(x1, x2 float64) float64 {
	return x1 - x2
}

func Multiplication(x1, x2 float64) float64 {
	return x1 * x2
}

func Division(x1, x2 float64) float64 {
	if x2 == 0.0 {
		if x1 >= 0.0 {
			return math.MaxFloat64
		} else {
			return math.MaxFloat64 * -1.0
		}
	} else {
		return x1 / x2
	}
}

func Mod(x1, x2 float64) float64 {
	return math.Mod(x1, x2)
}

func Pow(x1, x2 float64) float64 {
	return math.Pow(x1, x2)
}

func Max(x1, x2 float64) float64 {
	return math.Max(x1, x2)
}

func Min(x1, x2 float64) float64 {
	return math.Min(x1, x2)
}

// *********************
// *  UNARY OPERATORS  *
// *********************
func Exp(x float64) float64 {
	return math.Exp(x)
}

func Log(x float64) float64 {
	return math.Log(x)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Tan(x float64) float64 {
	return math.Tan(x)
}

func Pow2(x float64) float64 {
	return math.Pow(2, x)
}

func N(x float64) float64 {
	return x - 1.0
}

func P(x float64) float64 {
	return x + 1.0
}

