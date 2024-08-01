package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Deve inicializar a pilha__retorna_vazio", func(t *testing.T) {
		// arrange
		s := NewStack()

		//act
		e := s.EmpetyStack()

		// assert

		if e != true {
			t.Error("a pilha não pode estar cheia ao inicializar")
		}
	})

	t.Run("Deve inserir um elemento na pilha", func(t *testing.T) {
		// arrange
		s := NewStack()

		//act
		s.Push(0)

		//assert
		if s.EmpetyStack() {
			t.Error("deve inserir um elemento na pilha")
		}
	})

	t.Run("Deve retirar um elemento na pilha", func(t *testing.T) {
		// arrange
		s := NewStack()

		//act
		s.Pop()

		//assert
		if !s.EmpetyStack() {
			t.Error("a pilha deve estar vazia")
		}

	})

	t.Run("Deve inserir um elemento na pilha e retirar um_nenhum_elemento", func(t *testing.T) {
		// arrange
		s := NewStack()

		//act
		s.Push(0)
		s.Pop()

		//assert
		if !s.EmpetyStack() {
			t.Error("a pilha deve estar vazia")
		}

	})

	t.Run("Deve inserir dois elemento na pilha e retirar um_um_elemento", func(t *testing.T) {
		// arrange
		s := NewStack()

		//act
		s.Push(0)
		s.Push(1)
		s.Pop()

		//assert
		if s.EmpetyStack() {
			t.Error("a pilha deve não estar vazia")
		}

	})
}
