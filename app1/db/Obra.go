package db

type Obra struct {
	Id           int
	Name         string
	Date         string
	Materiales   []*Materiales
	Ingresos     []*Ingresos
	ManoObra     []*ManoObra
	Presupuestos []*Presupuestos
}

// adds
func (o *Obra) AddIngresos(i Ingresos) {
	i.Name = o.Name
	i.Id = len(o.Ingresos)
	o.Ingresos = append(o.Ingresos, &i)
}
func (o *Obra) AddMateriales(m Materiales) {
	m.Name = o.Name
	m.Id = len(o.Materiales)
	o.Materiales = append(o.Materiales, &m)
}
func (o *Obra) AddManoObra(m ManoObra) {
	m.Name = o.Name
	m.Id = len(o.ManoObra)
	o.ManoObra = append(o.ManoObra, &m)
}
func (o *Obra) AddPresupuestos(p Presupuestos) {
	p.Name = o.Name
	p.Id = len(o.Presupuestos)
	o.Presupuestos = append(o.Presupuestos, &p)
}

// gets
func (o *Obra) GetIngresos() *[]Ingresos {
	var res []Ingresos
	for _, item := range o.Ingresos {
		if item.Name == o.Name {
			res = append(res, *item)
		}
	}
	return &res
}
func (o *Obra) GetMateriales() *[]Materiales {
	var res []Materiales
	for _, item := range o.Materiales {
		if item.Name == o.Name {
			res = append(res, *item)
		}
	}
	return &res
}
func (o *Obra) GetManoObra() *[]ManoObra {
	var res []ManoObra
	for _, item := range o.ManoObra {
		if item.Name == o.Name {
			res = append(res, *item)
		}
	}
	return &res
}
func (o *Obra) GetPresupuesto() *[]Presupuestos {
	var res []Presupuestos
	for _, item := range o.Presupuestos {
		if item.Name == o.Name {
			res = append(res, *item)
		}
	}
	return &res
}

// removes
func (o *Obra) RemoveIngresos(id int) {
	for index, item := range o.Ingresos {
		if item.Id == id {
			o.Ingresos = append(o.Ingresos[:index], o.Ingresos[index+1:]...)
		}
	}
}
func (o *Obra) RemoveManoObra(id int) {
	for index, item := range o.ManoObra {
		if item.Id == id {
			o.ManoObra = append(o.ManoObra[:index], o.ManoObra[index+1:]...)
		}
	}
}
func (o *Obra) RemoveMateriales(id int) {
	for index, item := range o.Materiales {
		if item.Id == id {
			o.Materiales = append(o.Materiales[:index], o.Materiales[index+1:]...)
		}
	}
}
func (o *Obra) RemovePresupuesto(id int) {
	for index, item := range o.Presupuestos {
		if item.Id == id {
			o.Presupuestos = append(o.Presupuestos[:index], o.Presupuestos[index+1:]...)
		}
	}
}

// update
func (o *Obra) UpDateIngresos(id int, i Ingresos) {
	for index, item := range o.Ingresos {
		if item.Id == id {
			i.Id = index
			*o.Ingresos[index] = i
		}
	}
}
func (o *Obra) UpDateManoObra(id int, m ManoObra) {
	for index, item := range o.ManoObra {
		if item.Id == id {
			m.Id = index
			*o.ManoObra[index] = m
		}
	}
}
func (o *Obra) UpDateMateriales(id int, m Materiales) {
	for index, item := range o.Materiales {
		if item.Id == id {
			m.Id = index
			*o.Materiales[index] = m
		}
	}
}
func (o *Obra) UpDatePresupuesto(id int, p Presupuestos) {
	for index, item := range o.Presupuestos {
		if item.Id == id {
			p.Id = index
			*o.Presupuestos[index] = p
		}
	}
}
