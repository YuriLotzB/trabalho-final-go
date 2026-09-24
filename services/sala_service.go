package services

import (
	"strings"

	"api-gin/models"
	"api-gin/repositories"
)

type SalaService struct{ repo repositories.SalaRepository }

func NewSalaService(repo repositories.SalaRepository) *SalaService { return &SalaService{repo: repo} }

func (s *SalaService) Criar(sala models.Sala) (*models.Sala, error) {
	if strings.TrimSpace(sala.Nome) == "" || sala.Capacidade <= 0 {
		return nil, ErrValidacao
	}
	sala.Ativa = true
	return s.repo.Criar(&sala)
}

func (s *SalaService) Listar() ([]models.Sala, error) { return s.repo.Listar() }

func (s *SalaService) BuscarPorID(id int) (*models.Sala, error) {
	sala, err := s.repo.BuscarPorID(id)
	if err == repositories.ErrNaoEncontrado {
		return nil, ErrNaoEncontrado
	}
	return sala, err
}
