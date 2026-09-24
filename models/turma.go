package models

type Turma struct {
	ID         int       `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	Professor  string    `json:"professor"`
	AlunoIDs   []int     `json:"aluno_ids"`
	Alocacao   *Alocacao `json:"alocacao,omitempty"`
	Ativa      bool      `json:"ativa"`
}

type TurmaResumo struct {
	ID               int    `json:"id"`
	Nome             string `json:"nome"`
	Disciplina       string `json:"disciplina"`
	Professor        string `json:"professor"`
	QuantidadeAlunos int    `json:"quantidade_alunos"`
	Alocada          bool   `json:"alocada"`
}
