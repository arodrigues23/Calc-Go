package main

import "math"

// Print de resultados
type Searchable interface {
	ShowAnswers(n1, n2 float64) float64
}

// Soma
type Sum struct {
	Operation
}

func (s Sum) ShowAnswers(n1, n2 float64) float64 {
	return n1 + n2
}

// Subtração
type Sub struct {
	Operation
}

func (s Sub) ShowAnswers(n1, n2 float64) float64 {
	return n1 - n2
}

// Divisão
type Div struct {
	Operation
}

func (s Div) ShowAnswers(n1, n2 float64) float64 {
	if n2 != 0 {
		return n1 / n2
	}
	return 0
}

// Multiplicação
type Mul struct {
	Operation
}

func (s Mul) ShowAnswers(n1, n2 float64) float64 {
	return n1 * n2
}

// Potência
type Pow struct {
	Operation
}

func (s Pow) ShowAnswers(n1, n2 float64) float64 {
	return math.Pow(n1, n2)
}
