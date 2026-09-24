package repositories

import (
	"sync"

	"api-gin/models"
)

type TurmaRepository interface {
	Criar(turma *models.Turma) (*models.Turma, error)
	Listar() ([]models.Turma, error)
	BuscarPorID(id int) (*models.Turma, error)
	Atualizar(turma *models.Turma) error
}

type InMemoryTurmaRepository struct {
	mu     sync.RWMutex
	itens  map[int]*models.Turma
	nextID int
}

func NewInMemoryTurmaRepository() *InMemoryTurmaRepository {
	return &InMemoryTurmaRepository{itens: make(map[int]*models.Turma), nextID: 1}
}

func (r *InMemoryTurmaRepository) Criar(turma *models.Turma) (*models.Turma, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	copia := cloneTurma(turma)
	copia.ID = r.nextID
	r.nextID++
	r.itens[copia.ID] = copia
	return cloneTurma(copia), nil
}

func (r *InMemoryTurmaRepository) Listar() ([]models.Turma, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := make([]models.Turma, 0, len(r.itens))
	for _, turma := range r.itens {
		resultado = append(resultado, *cloneTurma(turma))
	}
	return resultado, nil
}

func (r *InMemoryTurmaRepository) BuscarPorID(id int) (*models.Turma, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	turma, ok := r.itens[id]
	if !ok {
		return nil, ErrNaoEncontrado
	}
	return cloneTurma(turma), nil
}

func (r *InMemoryTurmaRepository) Atualizar(turma *models.Turma) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.itens[turma.ID]; !ok {
		return ErrNaoEncontrado
	}
	r.itens[turma.ID] = cloneTurma(turma)
	return nil
}

func cloneTurma(t *models.Turma) *models.Turma {
	c := *t
	c.AlunoIDs = append([]int(nil), t.AlunoIDs...)
	if t.Alocacao != nil {
		a := *t.Alocacao
		c.Alocacao = &a
	}
	return &c
}
