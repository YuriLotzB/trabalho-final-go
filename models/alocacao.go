package models

type Alocacao struct {
	ID         int    `json:"id"`
	TurmaID    int    `json:"turma_id"`
	SalaID     int    `json:"sala_id"`
	DiaSemana  string `json:"dia_semana"`
	HoraInicio string `json:"hora_inicio"`
	HoraFim    string `json:"hora_fim"`
}
