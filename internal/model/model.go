package model

import (
	"log"
	"strings"
)

// import "fmt"

type calculation struct {
	line   string
	result float64
}

type Model struct {
	calcs []calculation
}

func NewModel() *Model {
	return &Model{
		calcs: []calculation{},
	}
}

func (m *Model) Calculate(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	calc := calculation{
		line:   line,
		result: 0,
	}

	m.calcs = append(m.calcs, calc)
	log.Printf("len(calcs)=%v", len(m.calcs))
	return "reesssuult for " + line
}

func addition(a, b float64) float64 {
	return a + b
}

func subtracion(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	return a / b
}

func pow(a, b float64) float64 {
	pow := float64(0)
	for i := 0; i < int(b); i++ {
		pow *= a
	}
	return pow
}

func remainderDiv(a, b float64) float64 {
	return 0.0
}
