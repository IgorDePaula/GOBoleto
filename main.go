package main

func main(){
/**
	   $this->boleto = new Boleto(new Itau());
          # Trazer uma lista de doadores do banco processar os itens abaixo:
          $this->boleto->setNomePagador(utf8_decode($this->nome));

          if (!empty($this->cpf)) {
              //$this->boleto->setCpfPagador(mask($this->cpf, '###.###.###-##'));
              $this->boleto->setCpfPagador($this->cpf);
          }

          if (!empty($this->cnpj)) {
              $this->boleto->setCnpjPagador($this->cnpj);
          }

          $this->boleto->setCarteira($this->carteira);
          $this->boleto->setDataDocumento($this->dataDocumento);
          $this->boleto->setDataProcessamento($this->dataProcessamento);
          $this->boleto->setDataVencimento($this->dataVencimento);
          $this->boleto->setNossoNumero(adicionarZeros($this->nossoNumero, 8)); #formatar e passar o número do Boleto com 8 posições
          $this->boleto->setLocalPagamento2(utf8_decode($this->localPagamento2));
          $this->boleto->setEnderecoPagador(($this->enderecoPagador));

          if ($this->boleto->getInstrucao4() === NULL || empty($this->boleto->getInstrucao4()) || $this->boleto->getInstrucao4() === 'instrucao4') {
              $this->boleto->setInstrucao4("");
          }

          if ($this->boleto->getInstrucao5() === NULL || empty($this->boleto->getInstrucao5()) || $this->boleto->getInstrucao5() === 'instrucao5') {
              $this->boleto->setInstrucao5("");
          }

          $this->boleto->setNoDocumento($this->numBoleto);

          $ADigitoNossoNumero =
              $this->boleto->getDvNossoNumeroItau($this->boleto->getAgencia() .
                  $this->boleto->getContaCorrente() .
                  $this->boleto->getCarteira() .
                  str_pad($this->numBoleto, 8, '0', STR_PAD_LEFT));

          if ($ADigitoNossoNumero == "10") {
              $this->boleto->setDvNossoNumero("X");
          } else {
              $this->boleto->setDvNossoNumero($ADigitoNossoNumero);
          }

          $valor = $this->valor;
          $this->boleto->setValorBoleto($valor); # Formatar para moeda R$

          return $this->boleto;
 */
}
