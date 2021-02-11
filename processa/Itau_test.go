package processa

import (
	"github.com/IgorDePaula/GOBoleto/processa"
	"github.com/IgorDePaula/GOBoleto/structs"
	"testing"
)

func TestFatorVencimento(t *testing.T) {
	boleto :=  structs.Boleto{
		DataVencimento:  structs.Date{Ano:2021, Mes:1, Dia:31},
	}

	itau := processa.Itau{Boleto: boleto}
	if itau.FatorVencimento() != 8517{
		t.Error("Fator Vencimento nao correspondente")
	}
}
func TestDac(t *testing.T) {
	itau := processa.Itau{}
	dac := itau.Dac(341911012)
	if dac != "1" {
		t.Error("DAC nao correspondente: "+dac)
	}
}

func TestCampo1(t *testing.T) {
	itau := processa.Itau{
		Boleto: structs.Boleto{
			NossoNumero: "22345",
			Carteira:    "109",
		},
	}
	campo1 := itau.Campo1()
	if campo1 != "3419109222" {
		t.Error("Campo1 nao correspondente: "+campo1)
	}
}