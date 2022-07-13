package figuras

import "fmt"

type figura interface {
	area() float64
	perimetro() float64
}

func Imprint(figura figura) {
	fmt.Println("El area de la figura es: ", figura.area())
	fmt.Println("El perimetro de la figura es: ", figura.perimetro())
}
