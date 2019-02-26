package area

//iniciando com letra minuscula é PRIVADO
//iniciando com letra maiuscula é PUBLICA

// Ponto ldkfalkdfj asçldgjarg
type Ponto struct {
	X float64
	Y float64
}

func catetos(a, b Ponto) (x, y float64) {
	x = a.X
	y = b.Y
	return
}

//Distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return cx + cy
}
