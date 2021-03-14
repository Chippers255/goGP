package main

type SubtractionGene struct {
    left Gene
    right Gene
}

func (g SubtractionGene) evaluate() float64 {
    return g.left.evaluate() - g.right.evaluate()
}

func (g SubtractionGene) print() string {
    return "(" + g.left.print() + "-" + g.right.print() + ")"
}

func (g SubtractionGene) gtype() string {
    return "B"
}