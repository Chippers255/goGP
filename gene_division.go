package main

type DivisionGene struct {
    left Gene
    right Gene
}

func (g DivisionGene) evaluate() float64 {
    return g.left.evaluate() / g.right.evaluate()
}

func (g DivisionGene) print() string {
    return "(" + g.left.print() + "/" + g.right.print() + ")"
}

func (g DivisionGene) gtype() string {
    return "B"
}