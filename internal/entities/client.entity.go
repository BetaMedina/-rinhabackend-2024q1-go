package entities

type Client struct {
	ID         string  `json:"_id" bson:"_id" validate:"required,uuid4"`
	FriendlyId int     `json:"id" bson:"id" validate:"required,uuid4"`
	Limite     float64 `json:"limit" bson:"limit" validate:"required,float64"`
	Saldo      float64 `json:"amount" bson:"amount" validate:"required,float64"`
}
