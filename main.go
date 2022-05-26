/*Crie um programa em Go que permita registrar serviços a serem feitos em pets, com as seguintes características: valor, tipo do serviço e se o serviço está ativo ou não.
O sistema deve possuir a funcionalidade de ativar e desativar um serviço (de acordo com o que o dono do estabelecimento tiver disponível naquele dia).
E o valor do serviço já está pré-definido em uma tabela de valores informada abaixo:
tipo valor
banho 50.0
tosa 85.0
hidratação 70.0
*/

package main

import (
	"fmt"
	"strings"

	"github.com/Erikaa81/pet-shop/entities"
)

func main() {

	Servico := "banho"

	servico := entities.Servico{
		Tipo: strings.ToUpper(Servico),
	}

	fmt.Println(servico.Tipo)
	fmt.Println(servico.VerificarValorServico())
	fmt.Println(servico.AtivarDesativarServico())
}
