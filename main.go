package main

import "fmt"

func main () {
    a := ConstantGene{value: 1.9}
    b := ConstantGene{value: 2.8}
    c := ConstantGene{value: 3.7}
    d := ConstantGene{value: 4.6}
    e := ConstantGene{value: 1.0}
    
    f := AdditionGene{left: a, right: b}
    g := MultiplicationGene{left: c, right: d}
    h := SubtractionGene{left: f, right: e}
    i := DivisionGene{left: h, right: g}
    
    fmt.Println(a.print())
    fmt.Println(b.print())
    fmt.Println(c.print())
    fmt.Println(d.print())
    fmt.Println(e.print())
    fmt.Println(f.print())
    fmt.Println(g.print())
    fmt.Println(h.print())
    fmt.Println(i.print())
    fmt.Println(i.evaluate())
}
