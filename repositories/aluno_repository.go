package repositories

import (
	"sync"

	"api-gin/models"
)

type AlunoRepository interface {
	Criar(aluno *models.Aluno) (*models.Aluno, error)
	Listar() ([]models.Aluno, error)
	BuscarPorID(id int) (*models.Aluno, error)
	BuscarPorMatricula(matricula string) (*models.Aluno, error)
}

type InMemoryAlunoRepository struct {
	mu     sync.RWMutex
	itens  map[int]*models.Aluno
	nextID int
}

func NewInMemoryAlunoRepository() *InMemoryAlunoRepository {
	return &InMemoryAlunoRepository{itens: make(map[int]*models.Aluno), nextID: 1}
}

func (r *InMemoryAlunoRepository) Criar(aluno *models.Aluno) (*models.Aluno, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, existente := range r.itens {
		if existente.Matricula == aluno.Matricula {
			return nil, ErrDuplicado
		}
	}
	copia := *aluno
	copia.ID = r.nextID
	r.nextID++
	r.itens[copia.ID] = &copia
	return cloneAluno(&copia), nil
}

func (r *InMemoryAlunoRepository) Listar() ([]models.Aluno, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := make([]models.Aluno, 0, len(r.itens))
	for _, aluno := range r.itens {
		resultado = append(resultado, *cloneAluno(aluno))
	}
	return resultado, nil
}

func (r *InMemoryAlunoRepository) BuscarPorID(id int) (*models.Aluno, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	aluno, ok := r.itens[id]
	if !ok {
		return nil, ErrNaoEncontrado
	}
	return cloneAluno(aluno), nil
}

func (r *InMemoryAlunoRepository) BuscarPorMatricula(matricula string) (*models.Aluno, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, aluno := range r.itens {
		if aluno.Matricula == matricula {
			return cloneAluno(aluno), nil
		}
	}
	return nil, ErrNaoEncontrado
}

func cloneAluno(a *models.Aluno) *models.Aluno {
	c := *a
	return &c
}
