package structs

type Boleto struct {
	Beneficiario Beneficiario
	Banco Banco
	Carteira string
	NomePagador string
	Aceite string
	NossoNumero string
	NumeroDocumento string
	DataProcessamento Date
	DataVencimento Date
	LocalPagamento string
	LocalPagamento2 string
	Instrucao1 string
	Instrucao2 string
	Instrucao3 string
	Instrucao4 string
	Instrucao5 string
	Moeda string
}
