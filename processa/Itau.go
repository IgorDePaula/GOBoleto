package processa

import (
	"fmt"
	"github.com/IgorDePaula/GOBoleto/structs"
	"strconv"
	"strings"
	"time"
)

type Itau struct {
	Boleto structs.Boleto
}

const NUMERO = "341"

func (itau Itau) FatorVencimento() int {
	database := ToDate(1997, 10, 7)
	maturity := ToDate(itau.Boleto.DataVencimento.Ano, itau.Boleto.DataVencimento.Mes, itau.Boleto.DataVencimento.Dia)

	result := maturity.Sub(database).Hours() / 24
	return int(result)
}

func (itau Itau) reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (itau Itau) Dac(numero int) string {
	members := strings.Split(itau.dacFirstStep(numero), "")
	result := 0
	for i := len(members); i > 0; i-- {
		par, _ := strconv.Atoi(strings.Join(members[i-1:i], ""))
		result += par
	}
	var mod = (result % 10)
	var dac = 10 - mod
	if dac == 10 {
		return strconv.Itoa(0)
	}
	return strconv.Itoa(dac)
}

func (itau Itau) dacFirstStep(numero int) string {
	count := len(strconv.Itoa(numero))
	result := []int{}
	lista := strings.Split(strconv.Itoa(numero), "")
	for i := count; i > 0; i-- {
		var fator int
		if i%2 != 0 {
			fator = 2
		} else {
			fator = 1
		}
		op, _ := strconv.Atoi(strings.Join(lista[i-1:i], ""))
		result = append(result, op*fator)
	}
	r := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ""), "[]")
	return itau.reverse(r)
}
func ToDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (itau Itau) Campo1() string {
	var num = NUMERO
	var moeda = "9"
	var carteira = itau.Boleto.Carteira
	var nn = itau.Boleto.NossoNumero[0:2]
	var firstNumber = num + moeda + carteira + nn
	dacNumber, _ := strconv.Atoi(firstNumber)
	var dac = itau.Dac(dacNumber)
	return num + moeda + carteira + nn + dac
}
func (itau Itau) Campo2() string {
	var rnn = itau.Boleto.NossoNumero[2:]
	var agencia = itau.Boleto.Agencia
	var conta = itau.Boleto.Conta
	var carteira = itau.Boleto.Carteira
	var nn = itau.Boleto.NossoNumero
	var dacArgs, _ = strconv.Atoi(rnn+agencia+conta+carteira+nn)
	var dac1 = itau.Dac(dacArgs)
	var agen3 = itau.Boleto.Agencia[0:3]
	dacArgs, _ = strconv.Atoi(agen3+dac1)
	var dac = itau.Dac(dacArgs)
	return dac
}