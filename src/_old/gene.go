package main

import "fmt"

// LOOK AT JAMES STUFF
// RANDOMLY SELECT A REFERENCE VALUE IN ARRAY LESS THAN CURRENT POSITION

// Set up Gene struct and associated methods
type Gene struct {
	ref1 int          // Chromosome array index value 1, used by unary and binary operators
	ref2 int          // Chromosome array index value 2, only used by binary operator
	style string      // Can be u (unary operator), b (binary operator), v (variable), or c (constant)
	operator string   // Operator character, only used if the gene is an operator
	constant float64  // Constant value, only used if the gene is a constant
    variable int      // Variable array index value, only used if the gene is a variable
}

// Print out gene information
func (g Gene) Info() {
	fmt.Println(g.left)
	fmt.Println(g.right)
	fmt.Println(g.style)
	fmt.Println(g.opperator)
	fmt.Println()
}

func (g Gene) Copy(inGene Gene) {
	g.ref1 = inGene.ref1
	g.ref2 = inGene.ref2
	g.style = inGene.style
	g.operator = inGene.operator
	g.constant = inGene.constant
	g.variable = inGene.variable
}

func (g Gene) ToString() string  {
	s := fmt.Sprintf("%s; %s; %v; %d; %d; %d", g.style, g.operator, g.constant, g.variable, g.ref1, g.ref2)

	return s
}
