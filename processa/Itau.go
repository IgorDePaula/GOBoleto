package processa

import (
	"github.com/IgorDePaula/GOBoleto/structs"
)
type Itau struct{

}
const NUMERO = "341"
/**func DvNossoNumero(campo) int {
	var multiplicador = 2;
	var multiplicacao = 0;
	soma_campo := 0;
	for i := len(campo); i>0; i-- {
		multiplicacao = campo[i-1:1] * multiplicador;
		if multiplicacao >= 10 {
			multiplicacao = multiplicacao[0:1] + multiplicacao[1:1]
		}
		soma_campo += multiplicacao
		multiplicador = (multiplicador % 2) + 1
	}
	var dac = 10 - (soma_campo %10)
	if dac == 10 {
		dac = 0
	}
	return dac

	$multiplicador = 2;
          $multiplicacao = 0;
          $soma_campo = 0;

          for ($i = strlen($campo); $i > 0; $i--) {
              $multiplicacao = substr($campo, $i-1, 1) * $multiplicador;

              if ($multiplicacao >= 10) {
                  $multiplicacao = substr($multiplicacao, 0, 1) + substr($multiplicacao, 1, 1);
              }
              $soma_campo += $multiplicacao;
              // valores assumidos: 212121...
              $multiplicador = ($multiplicador % 2) + 1;
          }
          $dac = 10 - ($soma_campo%10);
          if ($dac == 10)
              $dac = 0;

          return $dac;

} */
// $campo = $this->getNumero() . $this->boleto->getMoeda() . $this->boleto->getCarteira() . substr($this->boleto->getNossoNumero(), 0, 2);

func (Itau)Campo1 (boleto structs.Boleto) string{
	return NUMERO + boleto.Moeda + boleto.Carteira + boleto.NossoNumero[0:2]
}