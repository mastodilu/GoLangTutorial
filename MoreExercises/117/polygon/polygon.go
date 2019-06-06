package polygon

import (
	"fmt"
	"math"
)

// Polygon è un'interfaccia che rappresenta un poligono con metodi per restituire area e perimetro
type Polygon interface {
	Area() float64
	Perimeter() float64
}

// Circle rappresenta una circonferenza
type Circle struct {
	Radius float64
}

// Square rappresenta un quadrato
type Square struct {
	Side float64
}

// Info stampa le informazioni sul poligono passato
func Info(p Polygon) (string, error) {
	var s string
	switch p.(type) {
	case Circle:
		s = fmt.Sprintf("Circle(%v)-Area:%v,Perimeter:%v", p.(Circle).Radius, p.(Circle).Area(), p.(Circle).Perimeter())
	case Square:
		s = fmt.Sprintf("Square(%v)-Area:%v,Perimeter:%v", p.(Square).Side, p.(Square).Area(), p.(Square).Perimeter())
	default:
		return "", fmt.Errorf("unkown polygon")
	}
	return s, nil
}

// Area calcola l'area del cerchio
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter calcola il perimetro del cerchio
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi // 2 pi R
}

// Area calcola l'area del quadrato
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter calcola il perimetro del quadrato
func (s Square) Perimeter() float64 {
	return s.Side * 4
}
