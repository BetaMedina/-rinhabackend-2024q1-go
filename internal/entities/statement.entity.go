package entities

import (
	"time"
)

type Statement struct {
	Client    Client
	ID        string    `json:"id,omitempty" bson:"_id" validate:"required,uuid4"`
	Data      time.Time `json:"realizada_em" bson:"date" validate:"required"`
	Descricao string    `json:"descricao" bson:"description" validate:"required"`
	Tipo      string    `json:"tipo"  bson:"type" validate:"required,string"`
	Valor     float64   `json:"valor" bson:"value" validate:"required,float64"`
}
