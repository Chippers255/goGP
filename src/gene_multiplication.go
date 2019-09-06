package main

type MultiplicationGene struct {
    left Gene
    right Gene
}

func (g MultiplicationGene) evaluate() float64 {
    return g.left.evaluate() * g.right.evaluate()
}

func (g MultiplicationGene) print() string {
    return "(" + g.left.print() + "*" + g.right.print() + ")"
}

func (g MultiplicationGene) gtype() string {
    return "B"
}