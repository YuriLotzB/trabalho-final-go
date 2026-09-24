package services

import (
	"strings"

	"api-gin/models"
	"api-gin/repositories"
)

type AlunoService struct{ repo repositories.AlunoRepository }

func NewAlunoService(repo repositories.AlunoRepository) *AlunoService {
	return &AlunoService{repo: repo}
}

func (s *AlunoService) Criar(aluno models.Aluno) (*models.Aluno, error) {
	if strings.TrimSpace(aluno.Matricula) == "" || strings.TrimSpace(aluno.Nome) == "" || strings.TrimSpace(aluno.Email) == "" {
		return nil, ErrValidacao
	}
	criado, err := s.repo.Criar(&aluno)
	if err == repositories.ErrDuplicado {
		return nil, ErrDuplicado
	}
	return criado, err
}
func (s *AlunoService) Listar() ([]models.Aluno, error) { return s.repo.Listar() }
func (s *AlunoService) BuscarPorID(id int) (*models.Aluno, error) {
	aluno, err := s.repo.BuscarPorID(id)
	if err == repositories.ErrNaoEncontrado {
		return nil, ErrNaoEncontrado
	}
	return aluno, err
}
