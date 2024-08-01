package main

type Factura struct {
	Nif  string
	Nome string
	Data string
}

func NewEnvoice() Factura {
	return Factura{
		Nif:  "",
		Nome: "",
		Data: "",
	}
}

func (f *Factura) CreatEnvoice(x, y, z string) string {

	f.Nif = x
	f.Nome = y
	f.Data = z
	return f.Nif
}

func (f Factura) FindEnvoice(x string) bool {
	return f.Nif == x
}

func (f Factura) GetEnvoice(x string) Factura {
	if f.Nif == x {
		return f
	}

	return Factura{}
}
