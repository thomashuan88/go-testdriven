package factory

type student struct {
	Name  string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}
