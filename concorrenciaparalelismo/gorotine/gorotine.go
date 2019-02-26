package gorotine

// Pessoa a
type Pessoa struct {
	Nome  string
	Idade int
}

// Fale f
func (p *Pessoa) Fale(qtds int, c chan Pessoa) {
	for i := 0; i < qtds; i++ {
		p.Idade = i
		c <- *p
	}
	close(c)
}
