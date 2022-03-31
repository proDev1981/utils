package db

type Base struct {
	Obras []*Obra
}

func Open() *Base {
	return new(Base)
}
func (b *Base) NewObra(n string) *Obra {
	obra := &Obra{Id: len(b.Obras), Name: n}
	b.Obras = append(b.Obras, obra)
	return obra
}
