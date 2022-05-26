package entities

import (
	"errors"
	"strings"
)

const (
	BANHO      = "BANHO"
	TOSA       = "TOSA"
	HIDRATACAO = "HIDRATACAO"
)

type Servico struct {
	Tipo  string
	Ativo bool
}

func (s *Servico) VerificarValorServico() (float64, error) {
	if strings.ToUpper(s.Tipo) == BANHO {

		return 50.0, nil

	} else if strings.ToUpper(s.Tipo) == TOSA {

		return 85.0, nil

	} else if strings.ToUpper(s.Tipo) == HIDRATACAO {

		return 70.0, nil
	}

	return 0, errors.New("serviço não é ofertado")
}

func (s *Servico) AtivarDesativarServico() string {
	if s.Tipo == BANHO {
		s.Ativo = false

		if s.Ativo {
			return "Serviço disponível hoje."
		} else {
			return "Serviço não disponivel hoje."
		}
	}
	if s.Tipo == TOSA {
		s.Ativo = true

		if s.Ativo {
			return "Serviço disponível hoje."
		} else {
			return "Serviço não disponivel hoje."
		}
	}
	if s.Tipo == HIDRATACAO {
		s.Ativo = true

		if s.Ativo {
			return "Serviço disponível hoje."
		} else {
			return "Serviço não disponivel hoje."
		}
	}
	return ""
}
