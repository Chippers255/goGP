package main

import "fmt"

type ConstantGene struct {
    value float64
}

func (g ConstantGene) evaluate() float64 {
    return g.value
}

func (g ConstantGene) print() string {
    return fmt.Sprintf("%.2f", g.value)
}

func (g ConstantGene) gtype() string {
    return "C"
}