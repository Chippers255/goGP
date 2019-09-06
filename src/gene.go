package main

type Gene interface {
    evaluate() float64
    print() string
    gtype() string
}
