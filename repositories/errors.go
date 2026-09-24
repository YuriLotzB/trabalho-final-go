package repositories

import "errors"

var (
	ErrNaoEncontrado = errors.New("recurso não encontrado")
	ErrDuplicado     = errors.New("recurso duplicado")
)
