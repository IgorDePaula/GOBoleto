package processa

import (
	Boleto "github.com/IgorDePaula/GOBoleto/structs/Boleto"
	Date "github.com/IgorDePaula/GOBoleto/structs/Date"
	"testing"
)

func TestName(t *testing.T) {
	boleto :=  Boleto{
		DataVencimento:  Date{Ano:2021, Mes:1, Dia:31},
	}
	itau := Itau{Boleto: boleto}
	if itau.FatorVencimento() != 8517{
		t.Error("Fator Vencimento nao correspondente")
	}
}
