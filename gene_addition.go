package main

type AdditionGene struct {
    left Gene
    right Gene
}

func (g AdditionGene) evaluate() float64 {
    return g.left.evaluate() + g.right.evaluate()
}

func (g AdditionGene) print() string {
    return "(" + g.left.print() + "+" + g.right.print() + ")"
}

func (g AdditionGene) gtype() string {
    return "B"
}