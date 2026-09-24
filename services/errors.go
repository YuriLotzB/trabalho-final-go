package services

import "errors"

var (
	ErrValidacao              = errors.New("erro de validação")
	ErrNaoEncontrado          = errors.New("recurso não encontrado")
	ErrDuplicado              = errors.New("recurso duplicado")
	ErrCapacidadeInsuficiente = errors.New("capacidade insuficiente")
	ErrConflitoAgenda         = errors.New("conflito de agenda")
)
