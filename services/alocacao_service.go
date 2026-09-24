package services

import (
	"fmt"
	"strings"
	"time"

	"api-gin/models"
	"api-gin/repositories"
)

type AlocacaoService struct {
	turmas    repositories.TurmaRepository
	salas     repositories.SalaRepository
	alocacoes repositories.AlocacaoRepository
	alunos    repositories.AlunoRepository
}

func NewAlocacaoService(turmas repositories.TurmaRepository, salas repositories.SalaRepository, alocacoes repositories.AlocacaoRepository, alunos repositories.AlunoRepository) *AlocacaoService {
	return &AlocacaoService{turmas: turmas, salas: salas, alocacoes: alocacoes, alunos: alunos}
}

func (s *AlocacaoService) Alocar(turmaID int, salaID int, diaSemana, inicio, fim string) (*models.Alocacao, error) {
	if strings.TrimSpace(diaSemana) == "" || !horaValida(inicio) || !horaValida(fim) {
		return nil, ErrValidacao
	}
	if !diaValido(diaSemana) {
		return nil, ErrValidacao
	}
	if !horarioInicialAntesDoFinal(inicio, fim) {
		return nil, ErrValidacao
	}
	turma, err := s.turmas.BuscarPorID(turmaID)
	if err != nil {
		if err == repositories.ErrNaoEncontrado {
			return nil, ErrNaoEncontrado
		}
		return nil, err
	}
	if !turma.Ativa {
		return nil, ErrValidacao
	}
	sala, err := s.salas.BuscarPorID(salaID)
	if err != nil {
		if err == repositories.ErrNaoEncontrado {
			return nil, ErrNaoEncontrado
		}
		return nil, err
	}
	if !sala.Ativa {
		return nil, ErrValidacao
	}
	if len(turma.AlunoIDs) > sala.Capacidade {
		return nil, ErrCapacidadeInsuficiente
	}
	existentes, err := s.alocacoes.ListarPorSalaEDia(salaID, diaSemana)
	if err != nil {
		return nil, err
	}
	for _, existente := range existentes {
		if horariosSobrepostos(inicio, fim, existente.HoraInicio, existente.HoraFim) {
			return nil, ErrConflitoAgenda
		}
	}
	if turma.Alocacao != nil {
		return nil, ErrDuplicado
	}
	for _, alunoID := range turma.AlunoIDs {
		if err := s.verificarConflitoAluno(alunoID, turmaID, diaSemana, inicio, fim); err != nil {
			return nil, err
		}
	}
	a := &models.Alocacao{TurmaID: turmaID, SalaID: salaID, DiaSemana: diaSemana, HoraInicio: inicio, HoraFim: fim}
	criada, err := s.alocacoes.Criar(a)
	if err != nil {
		return nil, err
	}
	turma.Alocacao = criada
	if err := s.turmas.Atualizar(turma); err != nil {
		return nil, err
	}
	return criada, nil
}

func (s *AlocacaoService) verificarConflitoAluno(alunoID, turmaID int, dia, inicio, fim string) error {
	turmas, err := s.turmas.Listar()
	if err != nil {
		return err
	}
	for _, outra := range turmas {
		if outra.ID == turmaID || outra.Alocacao == nil || outra.Alocacao.DiaSemana != dia {
			continue
		}
		if !contains(outra.AlunoIDs, alunoID) {
			continue
		}
		if horariosSobrepostos(inicio, fim, outra.Alocacao.HoraInicio, outra.Alocacao.HoraFim) {
			return ErrConflitoAgenda
		}
	}
	return nil
}

func horaValida(valor string) bool {
	_, err := time.Parse("15:04", valor)
	return err == nil
}

func horarioInicialAntesDoFinal(inicio, fim string) bool {
	i, _ := time.Parse("15:04", inicio)
	f, _ := time.Parse("15:04", fim)
	return i.Before(f)
}

func horariosSobrepostos(novoInicio, novoFim, existenteInicio, existenteFim string) bool {
	ni, _ := time.Parse("15:04", novoInicio)
	nf, _ := time.Parse("15:04", novoFim)
	ei, _ := time.Parse("15:04", existenteInicio)
	ef, _ := time.Parse("15:04", existenteFim)
	return ni.Before(ef) && nf.After(ei)
}

func diaValido(dia string) bool {
	switch strings.ToLower(strings.TrimSpace(dia)) {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado", "domingo":
		return true
	default:
		return false
	}
}

func ErroInterno(err error) error { return fmt.Errorf("erro interno: %w", err) }
