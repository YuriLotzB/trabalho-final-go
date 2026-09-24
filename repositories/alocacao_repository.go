package repositories

import (
	"sync"

	"api-gin/models"
)

type AlocacaoRepository interface {
	Criar(alocacao *models.Alocacao) (*models.Alocacao, error)
	Listar() ([]models.Alocacao, error)
	BuscarPorID(id int) (*models.Alocacao, error)
	ListarPorSalaEDia(salaID int, diaSemana string) ([]models.Alocacao, error)
	ListarPorTurma(turmaID int) ([]models.Alocacao, error)
}

type InMemoryAlocacaoRepository struct {
	mu     sync.RWMutex
	itens  map[int]*models.Alocacao
	nextID int
}

func NewInMemoryAlocacaoRepository() *InMemoryAlocacaoRepository {
	return &InMemoryAlocacaoRepository{itens: make(map[int]*models.Alocacao), nextID: 1}
}

func (r *InMemoryAlocacaoRepository) Criar(a *models.Alocacao) (*models.Alocacao, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	c := *a
	c.ID = r.nextID
	r.nextID++
	r.itens[c.ID] = &c
	return cloneAlocacao(&c), nil
}

func (r *InMemoryAlocacaoRepository) Listar() ([]models.Alocacao, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := make([]models.Alocacao, 0, len(r.itens))
	for _, a := range r.itens {
		resultado = append(resultado, *cloneAlocacao(a))
	}
	return resultado, nil
}

func (r *InMemoryAlocacaoRepository) BuscarPorID(id int) (*models.Alocacao, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	a, ok := r.itens[id]
	if !ok {
		return nil, ErrNaoEncontrado
	}
	return cloneAlocacao(a), nil
}

func (r *InMemoryAlocacaoRepository) ListarPorSalaEDia(salaID int, diaSemana string) ([]models.Alocacao, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := []models.Alocacao{}
	for _, a := range r.itens {
		if a.SalaID == salaID && a.DiaSemana == diaSemana {
			resultado = append(resultado, *cloneAlocacao(a))
		}
	}
	return resultado, nil
}

func (r *InMemoryAlocacaoRepository) ListarPorTurma(turmaID int) ([]models.Alocacao, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := []models.Alocacao{}
	for _, a := range r.itens {
		if a.TurmaID == turmaID {
			resultado = append(resultado, *cloneAlocacao(a))
		}
	}
	return resultado, nil
}

func cloneAlocacao(a *models.Alocacao) *models.Alocacao {
	c := *a
	return &c
}
