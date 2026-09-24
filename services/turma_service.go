package services

import (
	"strings"

	"api-gin/models"
	"api-gin/repositories"
)

type TurmaService struct {
	repo      repositories.TurmaRepository
	alunos    repositories.AlunoRepository
	salas     repositories.SalaRepository
	alocacoes repositories.AlocacaoRepository
}

func NewTurmaService(repo repositories.TurmaRepository, alunos repositories.AlunoRepository, salas repositories.SalaRepository, alocacoes repositories.AlocacaoRepository) *TurmaService {
	return &TurmaService{repo: repo, alunos: alunos, salas: salas, alocacoes: alocacoes}
}

func (s *TurmaService) Criar(turma models.Turma) (*models.Turma, error) {
	if strings.TrimSpace(turma.Nome) == "" || strings.TrimSpace(turma.Disciplina) == "" || strings.TrimSpace(turma.Professor) == "" {
		return nil, ErrValidacao
	}
	turma.Ativa = true
	turma.AlunoIDs = []int{}
	return s.repo.Criar(&turma)
}

func (s *TurmaService) Listar() ([]models.TurmaResumo, error) {
	turmas, err := s.repo.Listar()
	if err != nil {
		return nil, err
	}
	result := make([]models.TurmaResumo, 0, len(turmas))
	for _, t := range turmas {
		result = append(result, models.TurmaResumo{ID: t.ID, Nome: t.Nome, Disciplina: t.Disciplina, Professor: t.Professor, QuantidadeAlunos: len(t.AlunoIDs), Alocada: t.Alocacao != nil})
	}
	return result, nil
}

func (s *TurmaService) BuscarPorID(id int) (*models.Turma, error) {
	turma, err := s.repo.BuscarPorID(id)
	if err == repositories.ErrNaoEncontrado {
		return nil, ErrNaoEncontrado
	}
	return turma, err
}

func (s *TurmaService) MatricularAluno(turmaID, alunoID int) (*models.Turma, error) {
	turma, err := s.BuscarPorID(turmaID)
	if err != nil {
		return nil, err
	}
	if !turma.Ativa {
		return nil, ErrValidacao
	}
	if _, err := s.alunos.BuscarPorID(alunoID); err != nil {
		if err == repositories.ErrNaoEncontrado {
			return nil, ErrNaoEncontrado
		}
		return nil, err
	}
	for _, id := range turma.AlunoIDs {
		if id == alunoID {
			return nil, ErrDuplicado
		}
	}
	if turma.Alocacao != nil {
		sala, err := s.salas.BuscarPorID(turma.Alocacao.SalaID)
		if err != nil {
			return nil, ErrNaoEncontrado
		}
		if len(turma.AlunoIDs)+1 > sala.Capacidade {
			return nil, ErrCapacidadeInsuficiente
		}
		if err := s.verificarConflitoAluno(alunoID, turma); err != nil {
			return nil, err
		}
	}
	turma.AlunoIDs = append(turma.AlunoIDs, alunoID)
	if err := s.repo.Atualizar(turma); err != nil {
		return nil, err
	}
	return turma, nil
}

func (s *TurmaService) ListarAlunos(turmaID int) ([]models.Aluno, error) {
	turma, err := s.BuscarPorID(turmaID)
	if err != nil {
		return nil, err
	}
	result := make([]models.Aluno, 0, len(turma.AlunoIDs))
	for _, id := range turma.AlunoIDs {
		aluno, err := s.alunos.BuscarPorID(id)
		if err != nil {
			continue
		}
		result = append(result, *aluno)
	}
	return result, nil
}

func (s *TurmaService) verificarConflitoAluno(alunoID int, novaTurma *models.Turma) error {
	if novaTurma.Alocacao == nil {
		return nil
	}
	turmas, err := s.repo.Listar()
	if err != nil {
		return err
	}
	for _, outra := range turmas {
		if outra.ID == novaTurma.ID || outra.Alocacao == nil {
			continue
		}
		if !contains(outra.AlunoIDs, alunoID) {
			continue
		}
		if outra.Alocacao.DiaSemana != novaTurma.Alocacao.DiaSemana {
			continue
		}
		if horariosSobrepostos(novaTurma.Alocacao.HoraInicio, novaTurma.Alocacao.HoraFim, outra.Alocacao.HoraInicio, outra.Alocacao.HoraFim) {
			return ErrConflitoAgenda
		}
	}
	return nil
}

func contains(ids []int, alvo int) bool {
	for _, id := range ids {
		if id == alvo {
			return true
		}
	}
	return false
}
