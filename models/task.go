package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	PENDENTE  TaskStatus = "pendente"
	ANDAMENTO TaskStatus = "andamento"
	CONCLUIDA TaskStatus = "concluida"
	ATRASADA  TaskStatus = "atrasada"
	URGENTE   TaskStatus = "urgente"
)

type Task struct {
	gorm.Model

	UserID uint `gorm:"not null" json:"user_id"`

	Titulo           string     `gorm:"not null" json:"titulo"`
	Descricao        string     `json:"descricao"`
	PrazoEntrega     *time.Time `json:"prazo_entrega"`
	TempoEstimado    *float64   `json:"tempo_estimado"`
	Repetitiva       bool       `gorm:"default:false" json:"repetitiva"`
	Status           TaskStatus `gorm:"type:varchar(20);default:'pendente'" json:"status"`
	RecompensaMoedas int        `gorm:"default:10;not null" json:"recompensa_moedas"`
	RecompensaXp     int        `gorm:"default:20;not null" json:"recompensa_xp"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}

type TaskResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserID           uint       `json:"user_id"`
	Titulo           string     `json:"titulo"`
	Descricao        string     `json:"descricao,omitempty"`
	PrazoEntrega     *time.Time `json:"prazo_entrega,omitempty", format:"date-time"`
	TempoEstimado    *float64   `json:"tempo_estimado,omitempty"`
	Repetitiva       bool       `json:"repetitiva"`
	Status           TaskStatus `json:"status"`
	RecompensaMoedas int        `json:"recompensa_moedas"`
	RecompensaXp     int        `json:"recompensa_xp"`
}
