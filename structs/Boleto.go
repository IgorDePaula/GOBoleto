package structs

import "time"

type Boleto struct {
	Beneficiario Beneficiario
	Banco Banco
	NomePagador string
	Aceite string
	NumeroDocumento string
	DataProcessamento time.Time
	DataVencimento time.Time
	LocalPagamento string
	LocalPagamento2 string
	Instrucao1 string
	Instrucao2 string
	Instrucao3 string
	Instrucao4 string
	Instrucao5 string
}
