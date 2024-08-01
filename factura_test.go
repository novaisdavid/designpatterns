package main

import (
	"testing"
)

func TestFactura(t *testing.T) {
	t.Run("Deve criar fatura", func(t *testing.T) {
		//arrange
		f := NewEnvoice()

		//act
		fact := f.CreatEnvoice("", "", "")

		//assert
		if fact != "" {
			t.Errorf(" erro ao criar a factura %s ", fact)
		}
	})

	t.Run("Deve criar fatura com dados neles_deve retornar o nif", func(t *testing.T) {
		//arrange
		nif := "99999999"
		nome := "nnnnnnn"
		data := "hoje"

		f := NewEnvoice()

		//act
		fact := f.CreatEnvoice(nif, nome, data)

		//assert
		if fact == "" {
			t.Errorf(" não deve retornar valor vazio %s ", fact)
		}
	})

	t.Run("Deve verificar se existe factura de um cliente com base ao nif_deve retornar true", func(t *testing.T) {
		//arrange
		nif := "99999999"
		nome := "nnnnnnn"
		data := "hoje"

		f := NewEnvoice()
		f.CreatEnvoice(nif, nome, data)
		//act
		created := f.FindEnvoice(nif)

		//assert
		if !created {
			t.Error("factura não encontrada!", created)
		}
	})

	t.Run("Deve verificar se existe factura de um cliente com base ao nif_deve retornar os dados da factura", func(t *testing.T) {
		//arrange
		nif := "99999999"
		nome := "nnnnnnn"
		data := "hoje"

		f := NewEnvoice()
		f.CreatEnvoice(nif, nome, data)
		//act
		date := f.GetEnvoice(nif)

		//assert
		if (date == Factura{}) {
			t.Errorf("dados da factura não encontrado! %s", date)
		}
	})
}
