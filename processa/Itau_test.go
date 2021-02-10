package processa

import (
	"github.com/IgorDePaula/GOBoleto/structs"
	"testing"
)

func TestName(t *testing.T) {
	boleto := structs.Boleto{
		DataVencimento: &structs.Date{Ano:2021, Mes:1, Dia:31},
	}
	itau := Itau{Boleto: boleto}
	if itau.FatorVencimento() != 8517{
		t.Error("Fator Vencimento nao correspondente")
	}
}
