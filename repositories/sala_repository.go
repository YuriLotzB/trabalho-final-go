package repositories

import (
	"sync"

	"api-gin/models"
)

type SalaRepository interface {
	Criar(sala *models.Sala) (*models.Sala, error)
	Listar() ([]models.Sala, error)
	BuscarPorID(id int) (*models.Sala, error)
	Atualizar(sala *models.Sala) error
}

type InMemorySalaRepository struct {
	mu     sync.RWMutex
	itens  map[int]*models.Sala
	nextID int
}

func NewInMemorySalaRepository() *InMemorySalaRepository {
	return &InMemorySalaRepository{itens: make(map[int]*models.Sala), nextID: 1}
}

func (r *InMemorySalaRepository) Criar(sala *models.Sala) (*models.Sala, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	copia := *sala
	copia.ID = r.nextID
	r.nextID++
	copia.Recursos = append([]string(nil), sala.Recursos...)
	r.itens[copia.ID] = &copia
	return cloneSala(&copia), nil
}

func (r *InMemorySalaRepository) Listar() ([]models.Sala, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	resultado := make([]models.Sala, 0, len(r.itens))
	for _, sala := range r.itens {
		resultado = append(resultado, *cloneSala(sala))
	}
	return resultado, nil
}

func (r *InMemorySalaRepository) BuscarPorID(id int) (*models.Sala, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sala, ok := r.itens[id]
	if !ok {
		return nil, ErrNaoEncontrado
	}
	return cloneSala(sala), nil
}

func (r *InMemorySalaRepository) Atualizar(sala *models.Sala) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.itens[sala.ID]; !ok {
		return ErrNaoEncontrado
	}
	copia := *sala
	copia.Recursos = append([]string(nil), sala.Recursos...)
	r.itens[sala.ID] = &copia
	return nil
}

func cloneSala(s *models.Sala) *models.Sala {
	c := *s
	c.Recursos = append([]string(nil), s.Recursos...)
	return &c
}
